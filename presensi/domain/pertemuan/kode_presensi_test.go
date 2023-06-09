package pertemuan_test

import (
	"regexp"
	"testing"
	"time"

	"its.id/akademik/presensi/domain/pertemuan"
)

func Test_panjang_kode_presensi_tidak_sesuai(t *testing.T) {

	_, err := pertemuan.NewKodePresensi("1", time.Time{})

	if err == nil {
		t.Fatal("seharusnya muncul error")
	}

	t.Log(err.Error())
}

func Test_waktu_berlaku_kode_presensi_tidak_sesuai(t *testing.T) {

	_, err := pertemuan.NewKodePresensi("123456", time.Time{})

	if err == nil {
		t.Fatal("seharusnya muncul error")
	}

	t.Log(err.Error())
}

func Test_kode_presensi_sesuai(t *testing.T) {

	kode := "123456"
	time := time.Date(2023, 5, 13, 1, 0, 0, 0, time.Local)

	kp, err := pertemuan.NewKodePresensi(kode, time)

	if err != nil {
		t.Log(err.Error())
		t.Fatal("seharusnya tidak muncul error")
	}

	if kp.Kode() != kode {
		t.Fatalf("kode presensi seharusnya %s tetapi %s.", kp.Kode(), kode)
	}

	if kp.BerlakuSampai() != time {
		t.Fatalf("waktu berlaku kode presensi seharusnya %s tetapi %s.", kp.BerlakuSampai(), time)
	}

}

func Test_buat_kode_presensi_random(t *testing.T) {

	masaBerlaku := time.Date(2023, 05, 16, 11, 0, 0, 0, time.Local)
	pattern := `^[0-9]{6}$`
	regex := regexp.MustCompile(pattern)

	kodePresensi, err := pertemuan.GenerateRandomKodePresensi(masaBerlaku)

	if err != nil {
		t.Fatalf("seharusnya tidak muncul error")
	}

	t.Logf("Kode presensi: %s", kodePresensi.Kode())

	if !regex.MatchString(kodePresensi.Kode()) {
		t.Fatal("kode presensi tidak sesuai format")
	}
}
