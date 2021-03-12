package main

import (
	"curd/model"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/golangatu?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Koneksi GAGAL")
	}
	fmt.Println("Koneksi BERHASIL")

	// mhs:= model.Mahasiswa{
	// 	Nama: "Anggraeni",
	// 	Nim : "200602069",
	// 	Alamat: "Sakra",
	// }
	//db.Create(&mahasiswa)

	dosen := model.Dosen{
		Nama:   "YAHYA,S.Kom,M.Kom",
		Nidn:   "0811068606",
		Alamat: "Lendang Bedurik",
	}
	db.Create(&dosen)
	// // Berfungsi untuk menyimpan data ke database

	//dosen

	// db.Model(&dosen).Where("id =?",1).Update("nama","purwarupa")
	db.Where("id = ?", 2).Delete(&model.Dosen{})
	// //Untuk Hapus Data
	// db.Where("id = ?",2).Delete(&model.Mahasiswa{})

	// mahasiswa
	// var mahasiswa []model.Mahasiswa
	// db.Find(&mahasiswa)

	// for _,m := range mahasiswa {
	// 	fmt.Println("Nama = " + m.Nama)
	// 	fmt.Println("Nim = "+ m.Nim)
	// 	fmt.Println("Alamat = "+ m.Alamat)
	// 	fmt.Println("============")
	// }

	var Dosen []model.Dosen
	db.Find(&dosen)

	for _, m := range Dosen {
		fmt.Println("Nama = " + m.Nama)
		fmt.Println("NIDN = " + m.Nidn)
		fmt.Println("Alamat = " + m.Alamat)
		fmt.Println("============")
	}

	// jalanin server(mhs)
	// r := gin.Default()

	// r.GET("/getmahasiswa", func(c *gin.Context) {
	// var mahasiswa []model.Mahasiswa
	// db.Find(&mahasiswa)
	// c.JSON(http.StatusOK,&mahasiswa)
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// dosen
	r := gin.Default()

	r.GET("/getdosen", func(c *gin.Context) {
		var dosen []model.Dosen
		db.Find(&dosen)
		c.JSON(http.StatusOK, &dosen)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
