package pertemuan_test

import (
	"testing"

	"its.id/akademik/presensi/domain/pertemuan"
)

func Test_panjang_deskripsi_topik_perkuliahan_sesuai(t *testing.T) {

	di := "1"
	de := "1"

	_, err := pertemuan.NewTopikPerkuliahan(di, de)

	if err == nil {
		t.Fatal("seharusnya muncul error")
	}

	t.Log(err.Error())
}

func Test_deskripsi_topik_perkuliahan_sesuai(t *testing.T) {

	di := "deskripsi topik perkuliahan bahasa Indonesia"
	de := "course topic description in English"

	tp, _ := pertemuan.NewTopikPerkuliahan(di, de)

	if tp.Deskripsi() != di {
		t.Fatalf("deskripsi %s seharusnya %s", tp.Deskripsi(), di)
	}

	if tp.DeskripsiEn() != de {
		t.Fatalf("deskripsi %s seharusnya %s", tp.Deskripsi(), di)
	}

}
