package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	kullaniciAdi string
	parola       string
}

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetFlags(0)
	if err != nil {
		fmt.Println("Log dosyası açılamadı.")
		os.Exit(0)
	}
	log.SetOutput(file)

	var giris int

	fmt.Println("Hoş geldiniz! Lütfen giriş türünüzü seçiniz:")
	fmt.Println("0 - Admin\n1 - Öğrenci")
	fmt.Scanln(&giris)

	if giris == 0 {
		admin()
	} else if giris == 1 {
		ogrenci()
	} else {
		fmt.Println("Lütfen geçerli bir giriş türü seçiniz!")
		fmt.Scanln(&giris)
	}
}

func admin() {
	admin := User{
		kullaniciAdi: "admin",
		parola:       "admin",
	}

	fmt.Println("Admin Giriş Paneli")
	fmt.Println("-------------------")

	girisBasarili := false

	for i := 5; i > 0 && girisBasarili == false; i-- {
		fmt.Print("Kullanıcı adı: ")
		fmt.Scanln(&admin.kullaniciAdi)
		fmt.Print("Parola: ")
		fmt.Scanln(&admin.parola)

		if admin.kullaniciAdi == "admin" && admin.parola == "admin" {
			fmt.Println("Admin girişi başarılı! Hoşgeldin,", admin.kullaniciAdi)

			log.Println("Kullanıcı Adı: " + admin.kullaniciAdi)
			log.Printf("Giriş Tarihi ve Saati: %s", time.Now().Format("2006-01-02 15:04:05"))
			log.Println("Giriş Durumu: " + "\033[32m" + "Başarılı" + "\033[0m")
			log.Println("------------------------------------------")

			girisBasarili = true

			var secim int

			fmt.Println("Seçiminizi yapın:")
			fmt.Println("0 - Logları görüntüle")
			fmt.Println("1 - Çıkış yap")
			fmt.Scanln(&secim)

			if secim == 0 {
				fmt.Println("Loglar:")
				fmt.Println("--------")
				readLogs()
			} else if secim == 1 {
				fmt.Println("Çıkış yapılıyor...")
				return
			} else {
				fmt.Println("Geçersiz seçim yapıldı. Çıkış yapılıyor...")
				return
			}
		} else {
			fmt.Printf("\033[35m"+"Hatalı kullanıcı adı veya parola! Kalan giriş hakkı: %d\n"+"\033[0m", i-1)

			log.Println("Kullanıcı Adı: " + admin.kullaniciAdi)
			log.Printf("Giriş Tarihi ve Saati: %s", time.Now().Format("2006-01-02 15:04:05"))
			log.Println("Giriş Durumu: " + "\033[31m" + "Başarısız" + "\033[0m")
			log.Println("------------------------------------------")
		}
	}
	if girisBasarili == false {
		fmt.Println("Giriş hakkınız bitti. Çıkış yapılıyor...")
	}
}

func ogrenci() {
	ogrenci := User{
		kullaniciAdi: "yavuz",
		parola:       "yavuz",
	}

	fmt.Println("Öğrenci Giriş Paneli:")
	fmt.Println("---------------------")

	girisBasarili := false

	for i := 5; i > 0 && girisBasarili == false; i-- {
		fmt.Print("Kullanıcı adı: ")
		fmt.Scanln(&ogrenci.kullaniciAdi)
		fmt.Print("Parola: ")
		fmt.Scanln(&ogrenci.parola)

		if ogrenci.kullaniciAdi == "yavuz" && ogrenci.parola == "yavuz" {
			fmt.Println("Öğrenci girişi başarılı! Hoşgeldin,", ogrenci.kullaniciAdi)

			log.Println("Kullanıcı Adı: " + ogrenci.kullaniciAdi)
			log.Printf("Giriş Tarihi ve Saati: %s", time.Now().Format("2006-01-02 15:04:05"))
			log.Println("Giriş Durumu: " + "\033[32m" + "Başarılı" + "\033[0m")
			log.Println("------------------------------------------")

			girisBasarili = true
		} else {
			fmt.Printf("\033[35m"+"Hatalı kullanıcı adı veya parola! Kalan giriş hakkı: %d\n"+"\033[0m", i-1)

			log.Println("Kullanıcı Adı: " + ogrenci.kullaniciAdi)
			log.Printf("Giriş Tarihi ve Saati: %s", time.Now().Format("2006-01-02 15:04:05"))
			log.Println("Giriş Durumu: " + "\033[31m" + "Başarısız" + "\033[0m")
			log.Println("------------------------------------------")
		}
	}
	if girisBasarili == false {
		fmt.Println("Giriş hakkınız bitti. Çıkış yapılıyor...")
	}
}

func readLogs() {
	fileContent, err := os.ReadFile("logs.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(fileContent))
}
