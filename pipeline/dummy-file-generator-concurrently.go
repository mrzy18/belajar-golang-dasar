package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const totalFile = 1500
const contentLength = 2500

var tempPath = filepath.Join(os.Getenv("TEMP"), "dummy-concurrent")

type FileInfo struct {
	Index       int
	FileName    string
	WorkerIndex int
	Err         error
}

func main() {
	log.Println("start")
	start := time.Now()

	generateFiles()

	duration := time.Since(start)
	fmt.Println("done in", duration.Seconds(), "seconds")
}

func randomString(length int) string {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}

	return string(b)
}

func generateFiles() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	// pipiline 1: job distribution
	chanFileIndex := generateFilesIndexes()

	// pipeline 2: creating files
	createFileWorker := 100
	chanFileResult := createFiles(chanFileIndex, createFileWorker)

	// track and print output
	counterTotal := 0
	counterSucces := 0
	for fileResult := range chanFileResult {
		if fileResult.Err != nil {
			log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
		} else {
			counterSucces++
		}
		counterTotal++
	}

	log.Printf("%d/%d of total files created", counterSucces, counterTotal)

}

func generateFilesIndexes() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for i := 0; i < totalFile; i++ {
			chanOut <- FileInfo{
				Index:    i,
				FileName: fmt.Sprintf("file-%d.txt", i),
			}
		}
		close(chanOut)
	}()
	return chanOut
}

func createFiles(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	wg := new(sync.WaitGroup)

	wg.Add(numberOfWorkers)

	go func() {
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {
				for job := range chanIn {
					filePath := filepath.Join(tempPath, job.FileName)
					content := randomString(contentLength)
					err := os.WriteFile(filePath, []byte(content), os.ModePerm)

					log.Println("worker", workerIndex, "working on", job.FileName)

					chanOut <- FileInfo{
						FileName:    job.FileName,
						WorkerIndex: workerIndex,
						Err:         err,
					}
				}
				wg.Done()
			}(workerIndex)
		}
	}()

	go func() {
		wg.Wait()
		close(chanOut)
	}()
	return chanOut
}
