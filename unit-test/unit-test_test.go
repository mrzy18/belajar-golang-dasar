package main

import "testing"

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestHitungVolume(t *testing.T) {
	t.Logf("Volume : %2.f\n", kubus.Volume())
	if kubus.Volume() != volumeSeharusnya {
		t.Errorf("Salah!, seharusnya %2.f\n", volumeSeharusnya)
	}
}
func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %2.f\n", kubus.Luas())
	if kubus.Luas() != luasSeharusnya {
		t.Errorf("Salah!, seharusnya %2.f\n", luasSeharusnya)
	}
}
func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling : %2.f\n", kubus.Keliling())
	if kubus.Keliling() != kelilingSeharusnya {
		t.Errorf("Salah!, seharusnya %2.f\n", kelilingSeharusnya)
	}
}

func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.Luas()
	}
}
