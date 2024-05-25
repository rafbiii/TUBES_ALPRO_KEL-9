package main

import "fmt"

const NMAX int = 2024

var Play_list PL
var banyak_lagu_di_setiap_PL [NMAX]int

type PL struct {
	nama_PL   [NMAX]string
	PL_ke     [NMAX]Lagu
	banyak_PL int
}

type sum struct {
	sumListener [NMAX]int
	sumRate     [NMAX]float64
	sumLagu     int
	sumArtist   [NMAX]int
	sumLikes    [NMAX]int
	sumDislikes [NMAX]int
	sumViews    [NMAX]int
}

type Lagu struct {
	artist [NMAX]string
	judul  [NMAX]string
	genre  [NMAX]string
	banyak sum
}
type Akun struct {
	username, password string
}

func login(acc *Akun) {
	var lanjut string
	fmt.Println("========================================================")
	fmt.Println("===================   MUSIC EXPLORER    ================")
	fmt.Println("========================================================")
	fmt.Println()
	fmt.Print("masukkan username: ")
	fmt.Scan(&acc.username)
	fmt.Print("masukkan password: ")
	fmt.Scan(&acc.password)
	fmt.Print("\033[H\033[2J")
	fmt.Println("========================================================")
	fmt.Printf("==       Selamat datang %s di Music Explorer       ==\n", acc.username)
	fmt.Println("========================================================")
	fmt.Print("masukkan angka apapun untuk melanjutkan : ")
	fmt.Scan(&lanjut)
	fmt.Print("\033[H\033[2J")
}

func MenuUtama() {
	fmt.Println("========================================================")
	fmt.Println("==                    Menu Utama                      ==")
	fmt.Println("========================================================")
	fmt.Println("==  1. Upload                                         ==")
	fmt.Println("==  2. Explore Lagu                                   ==")
	fmt.Println("==  3. Putar Lagu                                     ==")
	fmt.Println("==  4. Exit                                           ==")
	fmt.Println("========================================================")
	fmt.Print("Menu Yang Dipilih: ")
}

func PlayLagu() {
	fmt.Println("========================================================")
	fmt.Println("==                   Putar Lagu                       ==")
	fmt.Println("========================================================")
	fmt.Println("==  1. Cari Judul                                     ==")
	fmt.Println("==  2. Putar Dari Playlist                            ==")
	fmt.Println("==  3. Top 5 Hits                                     ==")
	fmt.Println("==  4. Top Rating                                     ==")
	fmt.Println("==  5. Kembali                                        ==")
	fmt.Println("========================================================")
	fmt.Print("Menu Yang Dipilih: ")
}

func pem_menu_utama(L *Lagu) {
	var pilihan int
	for pilihan != 4 {
		MenuUtama()
		fmt.Scan(&pilihan)
		fmt.Print("\033[H\033[2J")
		if pilihan == 1 {
			upload(L)
		} else if pilihan == 2 {
			explore(L)
		} else if pilihan == 3 {
			pem_play_lagu(L)
		} else if pilihan != 4 {
			fmt.Println("masukkan tidak valid. ulangi")
		}
	}
}
func pem_menu_upload() {
	fmt.Println("========================================================")
	fmt.Println("==                       Upload                       ==")
	fmt.Println("========================================================")
	fmt.Println("==  1. Upload Lagu                                    ==")
	fmt.Println("==  2. Edit Data Lagu                                 ==")
	fmt.Println("==  3. Kembali                                        ==")
	fmt.Println("========================================================")
}
func upload(L *Lagu) {
	var i, b, j, pilihan, ketemu int
	var pil, lanjut, tempjudul, tempArtist, tempGenre string
	ketemu = -1
	for pil != "3" {
		pem_menu_upload()
		fmt.Print("Pilih Opsi Yang Ingin Dilakukan : ")
		fmt.Scan(&pil)
		fmt.Print("\033[H\033[2J")
		if pil == "1" {
			fmt.Print("Berapa Banyak Lagu yang Ingin di Inputkan : ")
			fmt.Scan(&b)

			b = b + L.banyak.sumLagu
			for i = 0 + L.banyak.sumLagu; i < b; i++ {
				fmt.Print("Masukkan Judul Lagu   : ")
				fmt.Scan(&tempjudul)
				fmt.Print("Masukkan Nama Artist  : ")
				fmt.Scan(&tempArtist)
				fmt.Print("Masukkan Genre Lagu   : ")
				fmt.Scan(&tempGenre)
				fmt.Println()
				fmt.Println()
				for j = 0 + L.banyak.sumLagu; j < b; j++ {
					if L.artist[i] == tempArtist && L.judul[i] == tempjudul && L.genre[i] == tempGenre {
						ketemu++
					}
				}
				if ketemu < 0 {
					L.artist[i] = tempArtist
					L.judul[i] = tempjudul
					L.genre[i] = tempGenre
					L.banyak.sumLagu++
				} else {
					fmt.Print("Lagu Sudah Ada , Silahkan Upload lagu lain.")
				}
			}
			fmt.Print("\033[H\033[2J")
		} else if pil == "2" {
			if L.banyak.sumLagu > 0 {
				fmt.Println("========================================================")
				for i := 0; i < L.banyak.sumLagu; i++ {
					fmt.Printf("%-3d %-23s %-20s %-10s\n", i+1, L.judul[i], L.artist[i], L.genre[i])
				}
				fmt.Println("========================================================")
				fmt.Print("Pilih Lagu Yang ingin di Edit (1/2/3...): ")
				fmt.Scan(&pilihan)
				if pilihan <= L.banyak.sumLagu {
					fmt.Print("Masukkan Judul Lagu   : ")
					fmt.Scan(&tempjudul)
					fmt.Print("Masukkan Nama Artist  : ")
					fmt.Scan(&tempArtist)
					fmt.Print("Masukkan Genre Lagu   : ")
					fmt.Scan(&tempGenre)
					fmt.Print("\033[H\033[2J")
					for j = 0 + L.banyak.sumLagu; j < b; j++ {
						if L.artist[i] == tempArtist && L.judul[i] == tempjudul && L.genre[i] == tempGenre {
							ketemu++
						}
					}
					if ketemu < 0 {
						L.artist[pilihan-1] = tempArtist
						L.judul[pilihan-1] = tempjudul
						L.genre[pilihan-1] = tempGenre
						fmt.Println("========================================================")
						fmt.Println("==                   Data Lagu Saat ini               ==")
						fmt.Println("========================================================")
						for i := 0; i < L.banyak.sumLagu; i++ {
							fmt.Printf("%-3d %-23s %-20s %-10s\n", i+1, L.judul[i], L.artist[i], L.genre[i])
						}
					} else {
						fmt.Print("Lagu Sudah Ada , Silahkan Upload lagu lain.")
					}
				}
			} else {
				fmt.Println("========================================================")
				fmt.Println("==            Anda Belum Mengupload Lagu Apapun       ==")
				fmt.Println("========================================================")
				fmt.Print("Masukkan Angka berapapun untuk melanjutkan : ")
				fmt.Scan(&lanjut)
				fmt.Print("\033[H\033[2J")
			}
		} else if pil != "3" {
			fmt.Println("========================================================")
			fmt.Println("==               Eror Silahkan Coba Lagi              ==")
			fmt.Println("========================================================")
		}
	}
}

func pem_play_lagu(L *Lagu) {
	var pilihan int
	var lanjut string
	for pilihan != 5 {
		PlayLagu()
		fmt.Scan(&pilihan)
		fmt.Print("\033[H\033[2J")
		if L.banyak.sumLagu != 0 {
			if pilihan == 1 {
				cari_judul(L)
			} else if pilihan == 2 {
				pem_berdasar_PL(L)
			} else if pilihan == 3 {
				top_5_hits(L)
			} else if pilihan == 4 {
				top_rating(L)
			} else if pilihan != 5 {
				fmt.Println("Masukkan tidak valid silahkan coba lagi")
			}
		} else if pilihan != 5 {
			fmt.Println("=========================================================")
			fmt.Println("== Silahkan Upload Lagu Untuk Masuk ke Menu Putar Lagu ==")
			fmt.Println("=========================================================")
			fmt.Print("Masukkan angka berapa saja untuk melanjutkan : ")
			fmt.Scan(&lanjut)
			fmt.Print("\033[H\033[2J")
		}
	}
}

func top_rating(L *Lagu) {
	var pilihan int

	// Sorting berdasarkan rating
	sorting_rate(L)

	fmt.Println("========================================================")
	fmt.Println("==                    Top 5 Rating                    ==")
	fmt.Println("========================================================")
	for i := 0; i < 5 && i < L.banyak.sumLagu; i++ {
		fmt.Printf("%d. %s - %s : %.2f Rating\n", i+1, L.judul[i], L.artist[i], L.banyak.sumRate[i])
	}
	fmt.Println("Pilih lagu yang ingin didengar (1/2/3/4/5)")
	fmt.Println("========================================================")
	fmt.Scan(&pilihan)
	if pilihan >= 1 && pilihan <= 5 && pilihan-1 < L.banyak.sumLagu {
		sedang_memainkan_lagu(L, pilihan-1)
	}

}

func pem_berdasar_PL(L *Lagu) {
	var pilihan, pil2 int
	print_entire_PL(L)
	fmt.Print("Pilih Play List (1/2/3/4/5): ")
	fmt.Scan(&pilihan)
	fmt.Print("\033[H\033[2J")
	print_Lagu_playlist(L, pilihan-1)
	fmt.Print("Pilih Lagu untuk dimainkan (1/2/3/..): ")
	fmt.Scan(&pil2)
	fmt.Print("\033[H\033[2J")
	sedang_memainkan_lagu_playList(L, pilihan-1, pil2-1)
}

func sedang_memainkan_lagu_playList(L *Lagu, pil1, pil2 int) {
	var pilihan, i int
	var pilihan_s string
	fmt.Println("========================================================")
	fmt.Println("                  Lagu Sedang Dimainkan                 ")
	fmt.Println("========================================================")
	fmt.Println(" Judul  :", Play_list.PL_ke[pil1].judul[pil2])
	fmt.Println(" Artist :", Play_list.PL_ke[pil1].artist[pil2])
	fmt.Println(" Genre  :", Play_list.PL_ke[pil1].genre[pil2])
	fmt.Println("========================================================")
	fmt.Print(" Masukkan Huruf Apapun untuk Berhenti ")
	fmt.Scan(&pilihan_s)
	fmt.Println("========================================================")
	fmt.Println("Apakah anda menyukai lagunya ?")
	fmt.Println("1. Ya")
	fmt.Println("2. Tidak")
	fmt.Println("========================================================")
	fmt.Scan(&pilihan)
	for i = 0; i < L.banyak.sumLagu; i++ {
		if L.judul[i] == Play_list.PL_ke[pil1].judul[pil2] && L.artist[i] == Play_list.PL_ke[pil1].artist[pil2] {
			if pilihan == 1 {
				L.banyak.sumLikes[i]++
			} else if pilihan == 2 {
				L.banyak.sumDislikes[i]++
			}
			L.banyak.sumRate[i] = float64(L.banyak.sumLikes[i]*5) / float64(L.banyak.sumLikes[i]+L.banyak.sumDislikes[i])
			L.banyak.sumListener[i] = L.banyak.sumListener[i] + 1000
		}
	}
}

func cari_judul(L *Lagu) {
	var judul, artis string
	var pilihan, ketemu int
	var left, right, mid int

	fmt.Println("Masukkan judul lagu yang ingin didengar")
	fmt.Scan(&judul)
	fmt.Println("Masukkan nama artis dari musik yang ingin didengar")
	fmt.Scan(&artis)
	fmt.Print("\033[H\033[2J")
	ketemu = 0

	left = 0
	right = L.banyak.sumLagu
	for left <= right && ketemu == 0 {
		mid = (left + right) / 2
		if L.judul[mid] == judul && L.artist[mid] == artis {
			sedang_memainkan_lagu(L, mid)
			ketemu = 1
		}
		if L.judul[mid] < judul || (L.judul[mid] == judul && L.artist[mid] < artis) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if ketemu == 0 {
		fmt.Println("========================================================")
		fmt.Println("==                 Judul Tidak Ditemukan              ==")
		fmt.Println("==          Silahkan Upload Lagu Terlebih Dahulu      ==")
		fmt.Println("========================================================")
		fmt.Println("Apakah Anda Ingin Mengupload Lagunya ?")
		fmt.Println("1. Ya")
		fmt.Println("2. Tidak")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			upload(L)
		}
	}
}

func explore(L *Lagu) {
	var stop bool = false
	var pily_n, lanjut string
	var pilihan int
	for pilihan != 3 {
		fmt.Println("========================================================")
		fmt.Println("==                     Playlist                       ==")
		fmt.Println("========================================================")
		fmt.Println("==  1. Buat Daftar Putar                              ==")
		fmt.Println("==  2. Lihat Daftar Putar                             ==")
		fmt.Println("==  3. Kembali                                        ==")
		fmt.Println("========================================================")
		fmt.Print("Pilih menu (1/2/3/4): ")
		fmt.Scan(&pilihan)
		fmt.Print("\033[H\033[2J")
		if L.banyak.sumLagu != 0 {
			if pilihan == 1 {
				buat_dafput(L)
				pily_n = ""
				stop = false
				for !stop {
					fmt.Print("Ingin Tambahkan Lagu lain ke dalam Play List ini? (Y/N):")
					fmt.Scan(&pily_n)
					fmt.Print("\033[H\033[2J")
					if pily_n == "Y" || pily_n == "y" {
						tambah_lagu(L)
					}
					stop = pily_n == "N" || pily_n == "n"
				}
				Play_list.banyak_PL++
			} else if pilihan == 2 {
				lihat_dafput(L)
			} else if pilihan != 3 {
				fmt.Println("Masukkan Tidak Valid. Coba Lagi")
			}
		} else if pilihan != 3 {
			fmt.Println("========================================================")
			fmt.Println("== Silahkan Upload Lagu Untuk Masuk ke Menu Play List ==")
			fmt.Println("========================================================")
			fmt.Print("masukkan angka apapun untuk melanjutkan : ")
			fmt.Scan(&lanjut)
			fmt.Print("\033[H\033[2J")
		}
	}
}

func buat_dafput(L *Lagu) {
	var pilihan int
	fmt.Println("========================================================")
	fmt.Println("==  1. Pilih Berdasarkan Nama Artist                  ==")
	fmt.Println("==  2. Pilih Berdasarkan Genre                        ==")
	fmt.Println("==  3. Pilih Berdasarkan TOP 5 hits                   ==")
	fmt.Println("==  4. Pilih Berdasarkan Best Rating                  ==")
	fmt.Println("==  5. Kembali                                        ==")
	fmt.Println("========================================================")
	fmt.Print("Pilih Menu (1/2/3/4/5): ")
	fmt.Scan(&pilihan)
	fmt.Print("\033[H\033[2J")
	if pilihan == 1 {
		berdasar_artist(L)
	} else if pilihan == 2 {
		berdasar_genre(L)
	} else if pilihan == 3 {
		berdasar_hits(L)
	} else if pilihan == 4 {
		berdasar_rate(L)
	}
}

func sorting_rate(L *Lagu) {
	var i, pass int
	var tempRate float64
	var tempLikes, temp_listener, tempDislike, tempViews int
	var temp_artis, temp_genre, temp_judul string
	pass = 1
	for pass <= L.banyak.sumLagu-1 {
		i = pass
		temp_artis = L.artist[i]
		temp_genre = L.genre[i]
		temp_judul = L.judul[i]
		temp_listener = L.banyak.sumListener[i]
		tempLikes = L.banyak.sumLikes[i]
		tempDislike = L.banyak.sumDislikes[i]
		tempRate = L.banyak.sumRate[i]
		tempViews = L.banyak.sumViews[i]
		for (i > 0 && tempRate > L.banyak.sumRate[i-1]) || (i > 0 && tempRate == L.banyak.sumRate[i-1] && temp_judul < L.judul[i-1]) {
			L.artist[i] = L.artist[i-1]
			L.genre[i] = L.genre[i-1]
			L.judul[i] = L.judul[i-1]
			L.banyak.sumListener[i] = L.banyak.sumListener[i-1]
			L.banyak.sumLikes[i] = L.banyak.sumLikes[i-1]
			L.banyak.sumDislikes[i] = L.banyak.sumDislikes[i-1]
			L.banyak.sumRate[i] = L.banyak.sumRate[i-1]
			L.banyak.sumViews[i] = L.banyak.sumViews[i-1]
			i--
		}
		L.artist[i] = temp_artis
		L.genre[i] = temp_genre
		L.judul[i] = temp_judul
		L.banyak.sumListener[i] = temp_listener
		L.banyak.sumLikes[i] = tempLikes
		L.banyak.sumDislikes[i] = tempDislike
		L.banyak.sumRate[i] = tempRate
		L.banyak.sumViews[i] = tempViews
		pass++
	}
}

func berdasar_rate(L *Lagu) {
	var pilihan int
	sorting_rate(L)
	fmt.Print("Masukkan Nama Play List : ")
	fmt.Scan(&Play_list.nama_PL[Play_list.banyak_PL])
	fmt.Print("\033[H\033[2J")
	fmt.Println("========================================================")
	fmt.Println("==                     Top 5 Rate                     ==")
	for i := 0; i < 5; i++ {
		fmt.Printf("%-3d. %-23s %-17s Rating %-10.2f\n", i+1, L.judul[i], L.artist[i], L.banyak.sumRate[i])
	}

	fmt.Print("Pilih lagu yang ingin dipilih (1/2/3/4/5/6) 6.Untuk Kembali ")
	fmt.Scan(&pilihan)
	if pilihan <= 5 && pilihan > 0 {
		Play_list.PL_ke[Play_list.banyak_PL].artist[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.artist[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].genre[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.genre[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].judul[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.judul[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumLikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumLikes[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumDislikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumDislikes[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumViews[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumViews[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumRate[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumRate[pilihan-1]
		fmt.Println()
		fmt.Println("========================================================")
		fmt.Println("==             ~Lagu Berhasil ditambahkan~            ==")
		fmt.Println("========================================================")
		banyak_lagu_di_setiap_PL[Play_list.banyak_PL]++
	} else if pilihan != 6 {
		fmt.Println("========================================================")
		fmt.Println("==             Eror Gagal Menambahkan Lagu            ==")
		fmt.Println("========================================================")
	}

}

func berdasar_hits(L *Lagu) {
	var pilihan int
	sorting_hits(L)
	fmt.Print("Masukkan Nama Play List : ")
	fmt.Scan(&Play_list.nama_PL[Play_list.banyak_PL])
	fmt.Print("\033[H\033[2J")
	fmt.Println("========================================================")
	fmt.Println("==                     Top 5 Hits                     ==")
	fmt.Println("========================================================")
	fmt.Println("1. ", L.judul[0], L.artist[0], L.banyak.sumListener[0], "Views")
	fmt.Println("2. ", L.judul[1], L.artist[1], L.banyak.sumListener[1], "Views")
	fmt.Println("3. ", L.judul[2], L.artist[2], L.banyak.sumListener[2], "Views")
	fmt.Println("4. ", L.judul[3], L.artist[3], L.banyak.sumListener[3], "Views")
	fmt.Println("5. ", L.judul[4], L.artist[4], L.banyak.sumListener[4], "Views")
	fmt.Println("========================================================")
	fmt.Print("Pilih lagu yang ingin dipilih (1/2/3/4/5/6) 6.Untuk Kembali ")
	fmt.Scan(&pilihan)
	if pilihan <= 5 && pilihan > 0 {
		Play_list.PL_ke[Play_list.banyak_PL].artist[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.artist[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].genre[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.genre[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].judul[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.judul[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumLikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumLikes[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumDislikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumDislikes[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumViews[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumViews[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumRate[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumRate[pilihan-1]
		fmt.Println("========================================================")
		fmt.Println("==              ~Lagu Berhasil ditambahkan~           ==")
		fmt.Println("========================================================")
		banyak_lagu_di_setiap_PL[Play_list.banyak_PL]++
	} else if pilihan != 6 {
		fmt.Println("========================================================")
		fmt.Println("==             Eror Gagal Menambahkan Lagu            ==")
		fmt.Println("========================================================")
	}

}
func sorting_hits(L *Lagu) {
	var i, pass int
	var tempRate float64
	var tempLikes, temp_listener, tempDislike, tempViews int
	var temp_artis, temp_genre, temp_judul string
	pass = 1
	for pass <= L.banyak.sumLagu-1 {
		i = pass
		temp_artis = L.artist[i]
		temp_genre = L.genre[i]
		temp_judul = L.judul[i]
		temp_listener = L.banyak.sumListener[i]
		tempLikes = L.banyak.sumLikes[i]
		tempDislike = L.banyak.sumDislikes[i]
		tempRate = L.banyak.sumRate[i]
		tempViews = L.banyak.sumViews[i]
		for (i > 0 && temp_listener > L.banyak.sumListener[i-1]) || (i > 0 && temp_listener == L.banyak.sumListener[i-1] && temp_judul < L.judul[i-1]) {
			L.artist[i] = L.artist[i-1]
			L.genre[i] = L.genre[i-1]
			L.judul[i] = L.judul[i-1]
			L.banyak.sumListener[i] = L.banyak.sumListener[i-1]
			L.banyak.sumLikes[i] = L.banyak.sumLikes[i-1]
			L.banyak.sumDislikes[i] = L.banyak.sumDislikes[i-1]
			L.banyak.sumRate[i] = L.banyak.sumRate[i-1]
			L.banyak.sumViews[i] = L.banyak.sumViews[i-1]
			i--
		}
		L.artist[i] = temp_artis
		L.genre[i] = temp_genre
		L.judul[i] = temp_judul
		L.banyak.sumListener[i] = temp_listener
		L.banyak.sumLikes[i] = tempLikes
		L.banyak.sumDislikes[i] = tempDislike
		L.banyak.sumRate[i] = tempRate
		L.banyak.sumViews[i] = tempViews
		pass++
	}
}
func berdasar_artist(L *Lagu) {
	var pilihan int
	sorting_artist(L)
	fmt.Print("Masukkan Nama Play List : ")
	fmt.Scan(&Play_list.nama_PL[Play_list.banyak_PL])
	fmt.Print("\033[H\033[2J")
	fmt.Println("========================================================")
	fmt.Println("==                Silahkan Pilih Lagu                 ==")
	fmt.Println("==                     Artist A-Z                     ==")
	fmt.Println("========================================================")

	for i := 0; i < L.banyak.sumLagu; i++ {
		fmt.Printf("%d. \t %s \t %s \n", i+1, L.artist[i], L.judul[i])
	}
	fmt.Println("========================================================")
	fmt.Print("Lagu yang dipilih (masukkan nomor): ")
	fmt.Scan(&pilihan)
	fmt.Print("\033[H\033[2J")
	if pilihan <= L.banyak.sumLagu && pilihan > 0 {
		Play_list.PL_ke[Play_list.banyak_PL].artist[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.artist[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].genre[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.genre[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].judul[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.judul[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumLikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumLikes[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumDislikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumDislikes[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumViews[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumViews[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumRate[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumRate[pilihan-1]
		fmt.Println("========================================================")
		fmt.Println("==             ~Lagu Berhasil ditambahkan~            ==")
		fmt.Println("========================================================")
		banyak_lagu_di_setiap_PL[Play_list.banyak_PL]++
	} else {
		fmt.Println("========================================================")
		fmt.Println("==             Eror Gagal Menambahkan Lagu            ==")
		fmt.Println("========================================================")
	}
}

func sorting_artist(L *Lagu) {
	var idx, i int
	for i = 0; i < L.banyak.sumLagu; i++ {
		idx = findmin_artist(*L, L.banyak.sumLagu, i)
		swap_Lagu(L, idx, i)
	}
}

func findmin_artist(L Lagu, n int, i int) int {
	var Hurufmin string
	var min int
	Hurufmin = L.artist[i]
	min = i
	for i := i + 1; i < n; i++ {
		if L.artist[i] < Hurufmin {
			Hurufmin = L.artist[i]
			min = i
		}
	}
	return min
}

func swap_Lagu(L *Lagu, idx int, i int) {
	var temp_artis, temp_judul, temp_genre string
	var temp_hits, tempLikes, tempDislike, tempViews int
	var tempRate float64

	temp_artis = (L.artist)[idx]
	(L.artist)[idx] = (L.artist)[i]
	(L.artist)[i] = temp_artis

	temp_judul = (L.judul)[idx]
	(L.judul)[idx] = (L.judul)[i]
	(L.judul)[i] = temp_judul

	temp_genre = (L.genre)[idx]
	(L.genre)[idx] = (L.genre)[i]
	(L.genre)[i] = temp_genre

	temp_hits = (L.banyak.sumListener)[idx]
	(L.banyak.sumListener)[idx] = (L.banyak.sumListener)[i]
	(L.banyak.sumListener)[i] = temp_hits

	tempLikes = (L.banyak.sumLikes)[idx]
	(L.banyak.sumLikes)[idx] = (L.banyak.sumLikes)[i]
	(L.banyak.sumLikes)[i] = tempLikes

	tempDislike = (L.banyak.sumDislikes)[idx]
	(L.banyak.sumDislikes)[idx] = (L.banyak.sumDislikes)[i]
	(L.banyak.sumDislikes)[i] = tempDislike

	tempRate = (L.banyak.sumRate)[idx]
	(L.banyak.sumRate)[idx] = (L.banyak.sumRate)[i]
	(L.banyak.sumRate)[i] = tempRate

	tempViews = (L.banyak.sumViews)[idx]
	(L.banyak.sumViews)[idx] = (L.banyak.sumViews)[i]
	(L.banyak.sumViews)[i] = tempViews
}

func sorting_genre(L *Lagu) {
	var idx, i int
	for i = 0; i < L.banyak.sumLagu; i++ {
		idx = findmin_genre(*L, L.banyak.sumLagu, i)
		swap_Lagu(L, idx, i)
	}
}

func findmin_genre(L Lagu, n int, i int) int {
	var Hurufmin string
	var min int
	Hurufmin = L.genre[i]
	min = i
	for i := i + 1; i < n; i++ {
		if L.genre[i] < Hurufmin {
			Hurufmin = L.genre[i]
			min = i
		}
	}
	return min
}

func berdasar_genre(L *Lagu) {
	var pilihan int
	sorting_genre(L)
	fmt.Print("Masukkan Nama Play List : ")
	fmt.Scan(&Play_list.nama_PL[Play_list.banyak_PL])
	fmt.Print("\033[H\033[2J")
	fmt.Println("========================================================")
	fmt.Println("==                  Silahkan Pilih Lagu               ==")
	fmt.Println("==                       Genre A-Z                    ==")
	fmt.Println("========================================================")
	for i := 0; i < L.banyak.sumLagu; i++ {
		fmt.Printf("%d. \t %s \t  %s \n", i+1, L.genre[i], L.judul[i])
	}
	fmt.Println("========================================================")
	fmt.Print("Lagu yang dipilih (masukkan nomor): ")
	fmt.Scan(&pilihan)
	if pilihan <= L.banyak.sumLagu && pilihan > 0 {
		Play_list.PL_ke[Play_list.banyak_PL].artist[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.artist[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].genre[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.genre[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].judul[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.judul[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumLikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumLikes[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumDislikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumDislikes[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumViews[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumViews[pilihan-1]
		Play_list.PL_ke[Play_list.banyak_PL].banyak.sumRate[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumRate[pilihan-1]
		fmt.Println("========================================================")
		fmt.Println("==              ~Lagu Berhasil ditambahkan~           ==")
		fmt.Println("========================================================")
		banyak_lagu_di_setiap_PL[Play_list.banyak_PL]++
	} else {
		fmt.Println("========================================================")
		fmt.Println("==             Eror Gagal Menambahkan Lagu            ==")
		fmt.Println("========================================================")
	}
}

func print_entire_PL(L *Lagu) {
	var i int
	fmt.Println("========================================================")
	fmt.Println("==               Silahkan Pilih Play List             ==")
	fmt.Println("========================================================")
	for i = 0; i < L.banyak.sumLagu; i++ {
		fmt.Printf("%d. %s\n", i+1, Play_list.nama_PL[i])
	}
}
func lihat_dafput(L *Lagu) {
	var stop bool = false
	var pilihan int
	var pil2 string
	print_entire_PL(L)
	fmt.Print("Play List yang dipilih (1/2/3/4..): 6. Untuk kembali ")
	fmt.Scan(&pilihan)
	fmt.Print("\033[H\033[2J")
	if pilihan < 6 {
		print_Lagu_playlist(L, pilihan-1)
	}
	if pilihan < 6 {
		if Play_list.nama_PL[0] != "" {
			for !stop {
				fmt.Print("Pilih Opsi Lanjut : a(tambah Lagu), e(edit),d(hapus),b(kembali)  ")
				fmt.Scan(&pil2)
				if pil2 == "a" || pil2 == "A" {
					edit_tambah_PL(L, pilihan-1)
				} else if pil2 == "e" || pil2 == "E" {
					edit_PL(L, pilihan-1)
				} else if pil2 == "d" || pil2 == "D" {
					delete_PL(L, pilihan-1)
				} else if pil2 != "B" && pil2 != "b" {
					fmt.Println("Masukkan error coba lagi.")
				}
				stop = pil2 == "B" || pil2 == "b"
			}
		}
	}
}
func delete_PL(L *Lagu, p int) {
	var pilihan, simpan_PL int
	print_Lagu_playlist(L, p)
	fmt.Print("Masukkan Lagu Yang Ingin Dihapus : ")
	fmt.Scan(&pilihan)

	simpan_PL = banyak_lagu_di_setiap_PL[p]
	for j := pilihan - 1; j <= simpan_PL; j++ {
		Play_list.PL_ke[p].artist[j] = Play_list.PL_ke[p].artist[j+1]
		Play_list.PL_ke[p].judul[j] = Play_list.PL_ke[p].judul[j+1]
		Play_list.PL_ke[p].genre[j] = Play_list.PL_ke[p].genre[j+1]
		Play_list.PL_ke[p].banyak.sumListener[j] = Play_list.PL_ke[p].banyak.sumListener[j+1]
		Play_list.PL_ke[p].banyak.sumRate[j] = Play_list.PL_ke[p].banyak.sumRate[j+1]
		Play_list.PL_ke[p].banyak.sumLikes[j] = Play_list.PL_ke[p].banyak.sumLikes[j+1]
		Play_list.PL_ke[p].banyak.sumDislikes[j] = Play_list.PL_ke[p].banyak.sumDislikes[j+1]
		Play_list.PL_ke[p].banyak.sumViews[j] = Play_list.PL_ke[p].banyak.sumViews[j+1]
		Play_list.PL_ke[p].banyak.sumLikes[j] = Play_list.PL_ke[p].banyak.sumLikes[j+1]

	}
	banyak_lagu_di_setiap_PL[p]--
}

func edit_tambah_PL(L *Lagu, p int) {
	var pil2 int
	pil_tambah()
	fmt.Scan(&pil2)
	if pil2 == 1 {
		edit_tambah_berdasar_artist(L, p)
	} else if pil2 == 2 {
		edit_tambah_berdasar_genre(L, p)
	} else if pil2 == 3 {
		edit_tambah_berdasar_hits(L, p)
	} else if pil2 == 4 {
		edit_tambah_berdasar_rate(L, p)
	}
}
func edit_tambah_berdasar_rate(L *Lagu, pil1 int) {
	var ketemu bool = false
	var pilihan int
	sorting_rate(L)
	fmt.Println("========================================================")
	fmt.Println("==                  Silahkan Pilih Lagu               ==")
	fmt.Println("==                      TOP 5 Rate                    ==")
	fmt.Println("========================================================")
	for i := 0; i < 5; i++ {
		fmt.Printf("%-3d. %-23s %-17s Rating %-10.2f\n", i+1, L.judul[i], L.artist[i], L.banyak.sumRate[i])
	}
	fmt.Println("========================================================")
	fmt.Print("Lagu yang dipilih (masukkan nomor): ")
	fmt.Scan(&pilihan)
	fmt.Print("\033[H\033[2J")
	for j := 0; j < banyak_lagu_di_setiap_PL[pil1] && pilihan > 0 && pilihan <= 5; j++ {
		if Play_list.PL_ke[pil1].artist[j] == L.artist[pilihan-1] && Play_list.PL_ke[pil1].genre[j] == L.genre[pilihan-1] && Play_list.PL_ke[pil1].judul[j] == L.judul[pilihan-1] {
			ketemu = true
		}
	}
	if ketemu {
		fmt.Println("========================================================")
		fmt.Println("==          Lagu Sudah ada di dalam Play List         ==")
		fmt.Println("========================================================")
	} else {
		if pilihan > 0 && pilihan <= 5 {
			Play_list.PL_ke[pil1].artist[banyak_lagu_di_setiap_PL[pil1]] = L.artist[pilihan-1]
			Play_list.PL_ke[pil1].genre[banyak_lagu_di_setiap_PL[pil1]] = L.genre[pilihan-1]
			Play_list.PL_ke[pil1].judul[banyak_lagu_di_setiap_PL[pil1]] = L.judul[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumListener[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumListener[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumLikes[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumLikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumDislikes[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumDislikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumViews[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumViews[pilihan-1]
			fmt.Println("========================================================")
			fmt.Println("==             ~Lagu Berhasil ditambahkan~            ==")
			fmt.Println("========================================================")
			banyak_lagu_di_setiap_PL[pil1]++
		} else {
			fmt.Println("========================================================")
			fmt.Println("==             Eror Gagal Menambahkan Lagu            ==")
			fmt.Println("========================================================")
		}
	}
}
func Pil_edit() {
	fmt.Println("========================================================")
	fmt.Println("==                Edit Lagu Playlist                  ==")
	fmt.Println("========================================================")
	fmt.Println("==  1. Edit Berdasar Artist                           ==")
	fmt.Println("==  2. Edit Berdasar Genre                            ==")
	fmt.Println("==  3. Edit dari Top 5 Hits                           ==")
	fmt.Println("==  4. Edit dari Top Rating                           ==")
	fmt.Println("==  5. Kembali                                        ==")
	fmt.Println("========================================================")
}
func pil_tambah() {
	fmt.Println("========================================================")
	fmt.Println("==                Tambah Lagu Playlist                ==")
	fmt.Println("========================================================")
	fmt.Println("==  1. Tambah Berdasar Artist                         ==")
	fmt.Println("==  2. Tambah Berdasar Genre                          ==")
	fmt.Println("==  3. Tambah dari Top 5 Hits                         ==")
	fmt.Println("==  4. Tambah dari Top Rating                         ==")
	fmt.Println("==  5. Kembali                                        ==")
	fmt.Println("========================================================")
}
func edit_PL(L *Lagu, p int) {
	var pilihan, pil2 int
	print_Lagu_playlist(L, p)
	fmt.Println("Pilih Lagu yang akan di edit (1/2/3/...): ")
	fmt.Scan(&pilihan)
	if pilihan <= banyak_lagu_di_setiap_PL[p] && pilihan > 0 {
		Pil_edit()
		fmt.Scan(&pil2)
		if pil2 == 1 {
			edit_ubah_berdasar_artist(L, p, pilihan-1)
		} else if pil2 == 2 {
			edit_ubah_berdasar_genre(L, p, pilihan-1)
		} else if pil2 == 3 {
			edit_ubah_berdasar_hits(L, p, pilihan-1)
		} else if pil2 == 4 {
			edit_ubah_berdasar_rate(L, p, pilihan-1)
		}
	} else {
		fmt.Println("========================================================")
		fmt.Println("==                Eror Silahkan Coba Lagi             ==")
		fmt.Println("========================================================")
	}
}

func edit_ubah_berdasar_rate(L *Lagu, pil1 int, pil2 int) {
	var ketemu bool = false
	var pilihan int
	sorting_rate(L)
	fmt.Println("========================================================")
	fmt.Println("==                  Silahkan Pilih Lagu               ==")
	fmt.Println("==                      TOP 5 RATE                    ==")
	fmt.Println("========================================================")
	for i := 0; i < 5; i++ {
		fmt.Printf("%-3d. %-23s %-17s Rating %-10.2f\n", i+1, L.judul[i], L.artist[i], L.banyak.sumRate[i])
	}
	fmt.Println("========================================================")
	fmt.Print("Lagu yang dipilih (masukkan nomor): ")
	fmt.Scan(&pilihan)
	for j := 0; j < banyak_lagu_di_setiap_PL[pil1] && pilihan > 0 && pilihan <= 5; j++ {
		if Play_list.PL_ke[pil1].artist[j] == L.artist[pilihan-1] && Play_list.PL_ke[pil1].genre[j] == L.genre[pilihan-1] && Play_list.PL_ke[pil1].judul[j] == L.judul[pilihan-1] {
			ketemu = true
		}
	}
	if ketemu {
		fmt.Println("========================================================")
		fmt.Println("==                        Gagal                       ==")
		fmt.Println("==          Lagu Sudah ada di dalam Play List         ==")
		fmt.Println("========================================================")
	} else {
		if pilihan > 0 && pilihan <= 5 {
			Play_list.PL_ke[pil1].artist[pil2] = L.artist[pilihan-1]
			Play_list.PL_ke[pil1].genre[pil2] = L.genre[pilihan-1]
			Play_list.PL_ke[pil1].judul[pil2] = L.judul[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumListener[pil2] = L.banyak.sumListener[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumLikes[pil2] = L.banyak.sumLikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumDislikes[pil2] = L.banyak.sumDislikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumViews[pil2] = L.banyak.sumViews[pilihan-1]
			fmt.Println()
			fmt.Println("========================================================")
			fmt.Println("==                ~Lagu Berhasil Diubah~              ==")
			fmt.Println("========================================================")
		} else {
			fmt.Println("========================================================")
			fmt.Println("==                Eror Gagal Mengubah Lagu            ==")
			fmt.Println("========================================================")
		}
	}
}

func edit_ubah_berdasar_hits(L *Lagu, pil1 int, pil2 int) {
	var ketemu bool = false
	var pilihan int
	sorting_hits(L)
	fmt.Println("========================================================")
	fmt.Println("==                  Silahkan Pilih Lagu               ==")
	fmt.Println("==                      TOP 5 HITS                    ==")
	fmt.Println("========================================================")
	for i := 0; i < 5; i++ {
		fmt.Printf("%-3d. %-23s  %-17s %-10d Views\n", i+1, L.artist[i], L.judul[i], L.banyak.sumListener[i])
	}
	fmt.Println("========================================================")
	fmt.Print("Lagu yang dipilih (masukkan nomor): ")
	fmt.Scan(&pilihan)
	for j := 0; j < banyak_lagu_di_setiap_PL[pil1] && pilihan > 0 && pilihan <= 5; j++ {
		if Play_list.PL_ke[pil1].artist[j] == L.artist[pilihan-1] && Play_list.PL_ke[pil1].genre[j] == L.genre[pilihan-1] && Play_list.PL_ke[pil1].judul[j] == L.judul[pilihan-1] {
			ketemu = true
		}
	}
	if ketemu {
		fmt.Println("========================================================")
		fmt.Println("==                        Gagal                       ==")
		fmt.Println("==          Lagu Sudah ada di dalam Play List         ==")
		fmt.Println("========================================================")
	} else {
		if pilihan > 0 && pilihan <= 5 {
			Play_list.PL_ke[pil1].artist[pil2] = L.artist[pilihan-1]
			Play_list.PL_ke[pil1].genre[pil2] = L.genre[pilihan-1]
			Play_list.PL_ke[pil1].judul[pil2] = L.judul[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumListener[pil2] = L.banyak.sumListener[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumLikes[pil2] = L.banyak.sumLikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumDislikes[pil2] = L.banyak.sumDislikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumViews[pil2] = L.banyak.sumViews[pilihan-1]
			fmt.Println()
			fmt.Println("========================================================")
			fmt.Println("==                ~Lagu Berhasil Diubah~              ==")
			fmt.Println("========================================================")
		} else {
			fmt.Println("========================================================")
			fmt.Println("==                Eror Gagal Mengubah Lagu            ==")
			fmt.Println("========================================================")
		}
	}
}
func edit_ubah_berdasar_genre(L *Lagu, pil1 int, pil2 int) {
	var ketemu bool = false
	var pilihan int
	sorting_artist(L)
	fmt.Println("========================================================")
	fmt.Println("==                  Silahkan Pilih Lagu               ==")
	fmt.Println("==                      Genre A-Z                    ==")
	fmt.Println("========================================================")
	for i := 0; i < L.banyak.sumLagu; i++ {
		fmt.Printf("%-3d. %-27s %-17s \n", i+1, L.genre[i], L.judul[i])
	}
	fmt.Println("========================================================")
	fmt.Print("Lagu yang dipilih (masukkan nomor): ")
	fmt.Scan(&pilihan)
	for j := 0; j < banyak_lagu_di_setiap_PL[Play_list.banyak_PL] && pilihan > 0; j++ {
		if Play_list.PL_ke[pil1].artist[j] == L.artist[pilihan-1] && Play_list.PL_ke[pil1].genre[j] == L.genre[pilihan-1] && Play_list.PL_ke[pil1].judul[j] == L.judul[pilihan-1] {
			ketemu = true
		}
	}
	if ketemu {
		fmt.Println("========================================================")
		fmt.Println("==                        Gagal                       ==")
		fmt.Println("==          Lagu Sudah ada di dalam Play List         ==")
		fmt.Println("========================================================")
	} else {
		if pilihan > 0 && pilihan <= L.banyak.sumLagu {
			Play_list.PL_ke[pil1].artist[pil2] = L.artist[pilihan-1]
			Play_list.PL_ke[pil1].genre[pil2] = L.genre[pilihan-1]
			Play_list.PL_ke[pil1].judul[pil2] = L.judul[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumListener[pil2] = L.banyak.sumListener[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumLikes[pil2] = L.banyak.sumLikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumDislikes[pil2] = L.banyak.sumDislikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumViews[pil2] = L.banyak.sumViews[pilihan-1]
			fmt.Println()
			fmt.Println("========================================================")
			fmt.Println("==                ~Lagu Berhasil Diubah~              ==")
			fmt.Println("========================================================")
		} else {
			fmt.Println("========================================================")
			fmt.Println("==                Eror Gagal Mengubah Lagu            ==")
			fmt.Println("========================================================")
		}
	}
}

func edit_ubah_berdasar_artist(L *Lagu, pil1 int, pil2 int) {
	var ketemu bool = false
	var pilihan int
	sorting_artist(L)
	fmt.Println("========================================================")
	fmt.Println("==                  Silahkan Pilih Lagu               ==")
	fmt.Println("==                      Artist A-Z                    ==")
	fmt.Println("========================================================")
	for i := 0; i < L.banyak.sumLagu; i++ {
		fmt.Printf("%d. \t %s \t %s \n", i+1, L.artist[i], L.judul[i])
	}
	fmt.Println("========================================================")
	fmt.Print("Lagu yang dipilih (masukkan nomor): ")
	fmt.Scan(&pilihan)
	for j := 0; j < banyak_lagu_di_setiap_PL[pil1] && pilihan > 0; j++ {
		if Play_list.PL_ke[pil1].artist[j] == L.artist[pilihan-1] && Play_list.PL_ke[pil1].genre[j] == L.genre[pilihan-1] && Play_list.PL_ke[pil1].judul[j] == L.judul[pilihan-1] {
			ketemu = true
		}
	}
	if !ketemu {
		if pilihan > 0 && pilihan <= L.banyak.sumLagu {
			Play_list.PL_ke[pil1].artist[pil2] = L.artist[pilihan-1]
			Play_list.PL_ke[pil1].genre[pil2] = L.genre[pilihan-1]
			Play_list.PL_ke[pil1].judul[pil2] = L.judul[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumListener[pil2] = L.banyak.sumListener[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumLikes[pil2] = L.banyak.sumLikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumDislikes[pil2] = L.banyak.sumDislikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumViews[pil2] = L.banyak.sumViews[pilihan-1]
			fmt.Println()
			fmt.Println("========================================================")
			fmt.Println("==                ~Lagu Berhasil Diubah~              ==")
			fmt.Println("========================================================")
		} else {
			fmt.Println("========================================================")
			fmt.Println("==                  Gagal Mengubah Lagu               ==")
			fmt.Println("========================================================")
		}
	} else {
		fmt.Println("========================================================")
		fmt.Println("==                        Gagal                       ==")
		fmt.Println("==          Lagu Sudah ada di dalam Play List         ==")
		fmt.Println("========================================================")
	}
}

func edit_tambah_berdasar_hits(L *Lagu, pil1 int) {
	var ketemu bool = false
	var pilihan int
	sorting_hits(L)
	fmt.Println("========================================================")
	fmt.Println("==                  Silahkan Pilih Lagu               ==")
	fmt.Println("==                      TOP 5 HITS                    ==")
	fmt.Println("========================================================")
	for i := 0; i < 5; i++ {
		fmt.Printf("%-3d. %-23s  %-17s %-10d Views\n", i+1, L.artist[i], L.judul[i], L.banyak.sumListener[i])
	}
	fmt.Println("========================================================")
	fmt.Print("Lagu yang dipilih (masukkan nomor): ")
	fmt.Scan(&pilihan)
	fmt.Print("\033[H\033[2J")
	for j := 0; j < banyak_lagu_di_setiap_PL[pil1] && pilihan > 0; j++ {
		if Play_list.PL_ke[pil1].artist[j] == L.artist[pilihan-1] && Play_list.PL_ke[pil1].genre[j] == L.genre[pilihan-1] && Play_list.PL_ke[pil1].judul[j] == L.judul[pilihan-1] {
			ketemu = true
		}
	}
	if ketemu {
		fmt.Println("========================================================")
		fmt.Println("==          Lagu Sudah ada di dalam Play List         ==")
		fmt.Println("==============================================s==========")
	} else {
		if pilihan > 0 && pilihan <= 5 {
			Play_list.PL_ke[pil1].artist[banyak_lagu_di_setiap_PL[pil1]] = L.artist[pilihan-1]
			Play_list.PL_ke[pil1].genre[banyak_lagu_di_setiap_PL[pil1]] = L.genre[pilihan-1]
			Play_list.PL_ke[pil1].judul[banyak_lagu_di_setiap_PL[pil1]] = L.judul[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumListener[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumListener[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumLikes[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumLikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumDislikes[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumDislikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumViews[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumViews[pilihan-1]
			fmt.Println("========================================================")
			fmt.Println("==             ~Lagu Berhasil ditambahkan~            ==")
			fmt.Println("========================================================")
			banyak_lagu_di_setiap_PL[pil1]++
		} else {
			fmt.Println("========================================================")
			fmt.Println("==             Eror Gagal Menambahkan Lagu            ==")
			fmt.Println("========================================================")
		}
	}
}
func edit_tambah_berdasar_artist(L *Lagu, pil1 int) {
	var ketemu bool = false
	var pilihan int
	sorting_artist(L)
	fmt.Println("========================================================")
	fmt.Println("==                  Silahkan Pilih Lagu               ==")
	fmt.Println("==                      Artist A-Z                    ==")
	fmt.Println("========================================================")
	for i := 0; i < L.banyak.sumLagu; i++ {
		fmt.Printf("%d. \t %s \t %s \n", i+1, L.artist[i], L.judul[i])
	}
	fmt.Println("========================================================")
	fmt.Print("Lagu yang dipilih (masukkan nomor): ")
	fmt.Scan(&pilihan)
	for j := 0; j < banyak_lagu_di_setiap_PL[pil1] && pilihan > 0; j++ {
		if Play_list.PL_ke[pil1].artist[j] == L.artist[pilihan-1] && Play_list.PL_ke[pil1].genre[j] == L.genre[pilihan-1] && Play_list.PL_ke[pil1].judul[j] == L.judul[pilihan-1] {
			ketemu = true
		}
	}
	if ketemu {
		fmt.Println("========================================================")
		fmt.Println("==          Lagu Sudah ada di dalam Play List         ==")
		fmt.Println("========================================================")
	} else {
		if pilihan > 0 && pilihan <= L.banyak.sumLagu {
			Play_list.PL_ke[pil1].artist[banyak_lagu_di_setiap_PL[pil1]] = L.artist[pilihan-1]
			Play_list.PL_ke[pil1].genre[banyak_lagu_di_setiap_PL[pil1]] = L.genre[pilihan-1]
			Play_list.PL_ke[pil1].judul[banyak_lagu_di_setiap_PL[pil1]] = L.judul[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumListener[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumListener[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumLikes[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumLikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumDislikes[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumDislikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumViews[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumViews[pilihan-1]
			fmt.Println()
			fmt.Println("========================================================")
			fmt.Println("==             ~Lagu Berhasil ditambahkan~            ==")
			fmt.Println("========================================================")
			banyak_lagu_di_setiap_PL[pil1]++
		} else {
			fmt.Println("========================================================")
			fmt.Println("==             Eror Gagal Menambahkan Lagu            ==")
			fmt.Println("========================================================")
		}
	}
}
func edit_tambah_berdasar_genre(L *Lagu, pil1 int) {
	var ketemu bool = false
	var pilihan int
	sorting_genre(L)
	fmt.Println("========================================================")
	fmt.Println("==                  Silahkan Pilih Lagu               ==")
	fmt.Println("==                       Genre A-Z                    ==")
	fmt.Println("========================================================")
	for i := 0; i < L.banyak.sumLagu; i++ {
		fmt.Printf("%d. \t %s \t %s \n", i+1, L.artist[i], L.judul[i])
	}
	fmt.Println("========================================================")
	fmt.Print("Lagu yang dipilih (masukkan nomor): ")
	fmt.Scan(&pilihan)
	for j := 0; j < banyak_lagu_di_setiap_PL[pil1] && pilihan > 0; j++ {
		if Play_list.PL_ke[pil1].artist[j] == L.artist[pilihan-1] && Play_list.PL_ke[pil1].genre[j] == L.genre[pilihan-1] && Play_list.PL_ke[pil1].judul[j] == L.judul[pilihan-1] {
			ketemu = true
		}
	}
	if ketemu {
		fmt.Println("========================================================")
		fmt.Println("==          Lagu Sudah ada di dalam Play List         ==")
		fmt.Println("========================================================")
	} else {
		if pilihan > 0 && pilihan <= L.banyak.sumLagu {
			Play_list.PL_ke[pil1].artist[banyak_lagu_di_setiap_PL[pil1]] = L.artist[pilihan-1]
			Play_list.PL_ke[pil1].genre[banyak_lagu_di_setiap_PL[pil1]] = L.genre[pilihan-1]
			Play_list.PL_ke[pil1].judul[banyak_lagu_di_setiap_PL[pil1]] = L.judul[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumListener[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumListener[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumLikes[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumLikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumDislikes[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumDislikes[pilihan-1]
			Play_list.PL_ke[pil1].banyak.sumViews[banyak_lagu_di_setiap_PL[pil1]] = L.banyak.sumViews[pilihan-1]
			fmt.Println()
			fmt.Println("========================================================")
			fmt.Println("==             ~Lagu Berhasil ditambahkan~            ==")
			fmt.Println("========================================================")
			banyak_lagu_di_setiap_PL[pil1]++
		} else {
			fmt.Println("========================================================")
			fmt.Println("==             Eror Gagal Menambahkan Lagu            ==")
			fmt.Println("========================================================")
		}
	}
}
func print_Lagu_playlist(L *Lagu, p int) {
	var i int
	fmt.Println("========================================================")
	fmt.Println("Play List : ", Play_list.nama_PL[p])
	fmt.Println("========================================================")
	for i = 0; i < banyak_lagu_di_setiap_PL[p]; i++ {
		fmt.Printf("%d. %s %30.s %30.s\n", i+1, Play_list.PL_ke[p].judul[i], Play_list.PL_ke[p].artist[i], Play_list.PL_ke[p].genre[i])
	}
}
func tambah_lagu(L *Lagu) {
	var pilihan int
	fmt.Println("========================================================")
	fmt.Println("==                Tambah Lagu Playlist                ==")
	fmt.Println("========================================================")
	fmt.Println("==  1. Pilih Berdasarkan Nama Artist                  ==")
	fmt.Println("==  2. Pilih Berdasarkan Genre                        ==")
	fmt.Println("==  3. Pilih Berdasarkan TOP 5 hits                   ==")
	fmt.Println("==  4. Pilih Berdasarkan Best Rating                  ==")
	fmt.Println("==  5. Kembali                                        ==")
	fmt.Println("========================================================")
	fmt.Print("Pilih Menu (1/2/3/4/5): ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		tambah_berdasar_artist(L)
	} else if pilihan == 2 {
		tambah_berdasar_genre(L)
	} else if pilihan == 3 {
		tambah_berdasar_hits(L)
	} else if pilihan == 4 {
		tambah_berdasar_rate(L)
	}
}
func tambah_berdasar_genre(L *Lagu) {
	var ketemu bool = false
	var pilihan int
	sorting_genre(L)
	fmt.Println("========================================================")
	fmt.Println("==                  Silahkan Pilih Lagu               ==")
	fmt.Println("==                       Genre A-Z                    ==")
	fmt.Println("========================================================")
	for i := 0; i < L.banyak.sumLagu; i++ {
		fmt.Printf("%d. \t %s \t %s \n", i+1, L.genre[i], L.judul[i])
	}
	fmt.Println("========================================================")
	fmt.Print("Lagu yang dipilih (masukkan nomor): ")
	fmt.Scan(&pilihan)
	for j := 0; j < L.banyak.sumLagu && pilihan != 0; j++ {
		if Play_list.PL_ke[Play_list.banyak_PL].artist[j] == L.artist[pilihan-1] && Play_list.PL_ke[Play_list.banyak_PL].genre[j] == L.genre[pilihan-1] && Play_list.PL_ke[Play_list.banyak_PL].judul[j] == L.judul[pilihan-1] {
			ketemu = true
		}
	}
	if ketemu {
		fmt.Println("========================================================")
		fmt.Println("==          Lagu Sudah ada di dalam Play List         ==")
		fmt.Println("========================================================")
	} else {
		if pilihan > 0 && pilihan <= L.banyak.sumLagu {
			Play_list.PL_ke[Play_list.banyak_PL].artist[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.artist[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].genre[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.genre[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].judul[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.judul[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].banyak.sumListener[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumListener[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].banyak.sumLikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumLikes[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].banyak.sumDislikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumDislikes[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].banyak.sumViews[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumViews[pilihan-1]
			fmt.Println()
			fmt.Println("========================================================")
			fmt.Println("==             ~Lagu Berhasil ditambahkan~            ==")
			fmt.Println("========================================================")
			banyak_lagu_di_setiap_PL[Play_list.banyak_PL]++
		} else {
			fmt.Println("========================================================")
			fmt.Println("==             Eror Gagal Menambahkan Lagu            ==")
			fmt.Println("========================================================")
		}
	}
}

func tambah_berdasar_artist(L *Lagu) {
	var ketemu bool = false
	var pilihan int
	sorting_artist(L)
	fmt.Println("========================================================")
	fmt.Println("==                Silahkan Pilih Lagu                 ==")
	fmt.Println("==                     Artist A-Z                     ==")
	fmt.Println("========================================================")

	for i := 0; i < L.banyak.sumLagu; i++ {
		fmt.Printf("%d. \t %s \t %s \n", i+1, L.artist[i], L.judul[i])
	}
	fmt.Println("========================================================")
	fmt.Print("Lagu yang dipilih (masukkan nomor): ")
	fmt.Scan(&pilihan)
	for j := 0; j < banyak_lagu_di_setiap_PL[Play_list.banyak_PL] && pilihan != 0; j++ {
		if Play_list.PL_ke[Play_list.banyak_PL].artist[j] == L.artist[pilihan-1] && Play_list.PL_ke[Play_list.banyak_PL].genre[j] == L.genre[pilihan-1] && Play_list.PL_ke[Play_list.banyak_PL].judul[j] == L.judul[pilihan-1] {
			ketemu = true
		}
	}
	if ketemu {
		fmt.Println("========================================================")
		fmt.Println("==          Lagu Sudah ada di dalam Play List         ==")
		fmt.Println("========================================================")
	} else {
		if pilihan > 0 && pilihan <= L.banyak.sumLagu {
			Play_list.PL_ke[Play_list.banyak_PL].artist[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.artist[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].genre[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.genre[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].judul[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.judul[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].banyak.sumListener[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumListener[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].banyak.sumLikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumLikes[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].banyak.sumDislikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumDislikes[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].banyak.sumViews[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumViews[pilihan-1]
			fmt.Println()
			fmt.Println("========================================================")
			fmt.Println("==             ~Lagu Berhasil ditambahkan~            ==")
			fmt.Println("========================================================")
			banyak_lagu_di_setiap_PL[Play_list.banyak_PL]++
		} else {
			fmt.Println("========================================================")
			fmt.Println("==             Eror Gagal Menambahkan Lagu            ==")
			fmt.Println("========================================================")
		}
	}
}

func sedang_memainkan_lagu(L *Lagu, i int) {
	var pilihan int
	var pilihan_s string
	fmt.Println("========================================================")
	fmt.Println("                  Lagu Sedang Dimainkan                 ")
	fmt.Println("========================================================")
	fmt.Println(" Judul  :", L.judul[i])
	fmt.Println(" Artist :", L.artist[i])
	fmt.Println(" Genre  :", L.genre[i])
	fmt.Println(" Masukkan huruf Apapun untuk keluar ")
	fmt.Println("========================================================")
	fmt.Scan(&pilihan_s)
	fmt.Println("========================================================")
	fmt.Println("Apakah anda menyukai lagunya ?")
	fmt.Println("1. Ya")
	fmt.Println("2. Tidak")
	fmt.Println("========================================================")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		L.banyak.sumLikes[i]++
	} else if pilihan == 2 {
		L.banyak.sumDislikes[i]++
	}
	L.banyak.sumRate[i] = float64(L.banyak.sumLikes[i]*5) / float64(L.banyak.sumLikes[i]+L.banyak.sumDislikes[i])
	L.banyak.sumListener[i] = L.banyak.sumListener[i] + 1000
}
func top_5_hits(L *Lagu) {
	var pilihan int

	sorting_hits(L)

	fmt.Println("========================================================")
	fmt.Println("==                     Top 5 Hits                     ==")
	fmt.Println("========================================================")
	fmt.Println("1. ", L.judul[0], L.artist[0], L.banyak.sumListener[0], "Views")
	fmt.Println("2. ", L.judul[1], L.artist[1], L.banyak.sumListener[1], "Views")
	fmt.Println("3. ", L.judul[2], L.artist[2], L.banyak.sumListener[2], "Views")
	fmt.Println("4. ", L.judul[3], L.artist[3], L.banyak.sumListener[3], "Views")
	fmt.Println("5. ", L.judul[4], L.artist[4], L.banyak.sumListener[4], "Views")
	fmt.Println("========================================================")
	fmt.Print("Pilih lagu yang ingin didengar (1/2/3/4/5)")
	fmt.Scan(&pilihan)
	fmt.Print("\033[H\033[2J")
	if pilihan == 1 {
		sedang_memainkan_lagu(L, 0)
	} else if pilihan == 2 {
		sedang_memainkan_lagu(L, 1)
	} else if pilihan == 3 {
		sedang_memainkan_lagu(L, 2)
	} else if pilihan == 4 {
		sedang_memainkan_lagu(L, 3)
	} else if pilihan == 5 {
		sedang_memainkan_lagu(L, 4)
	}
}
func tambah_berdasar_hits(L *Lagu) {
	var pilihan int
	var ketemu bool
	sorting_hits(L)
	fmt.Println("========================================================")
	fmt.Println("==                     Top 5 Hits                     ==")
	fmt.Println("========================================================")
	for i := 0; i < 5; i++ {
		fmt.Printf("%-3d. %-23s  %-17s %-10d Views\n", i+1, L.artist[i], L.judul[i], L.banyak.sumListener[i])
	}
	fmt.Println("========================================================")
	fmt.Print("Pilih lagu yang ingin dipilih (1/2/3/4/5/6) 6.Untuk Kembali ")
	fmt.Scan(&pilihan)
	for j := 0; j < banyak_lagu_di_setiap_PL[Play_list.banyak_PL] && pilihan > 0 && pilihan <= 5; j++ {
		if Play_list.PL_ke[Play_list.banyak_PL].artist[j] == L.artist[pilihan-1] && Play_list.PL_ke[Play_list.banyak_PL].genre[j] == L.genre[pilihan-1] && Play_list.PL_ke[Play_list.banyak_PL].judul[j] == L.judul[pilihan-1] {
			ketemu = true
		}
	}
	if ketemu {
		fmt.Println("========================================================")
		fmt.Println("==          Lagu Sudah ada di dalam Play List         ==")
		fmt.Println("========================================================")
	} else {
		if pilihan > 0 && pilihan <= 5 {
			Play_list.PL_ke[Play_list.banyak_PL].artist[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.artist[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].genre[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.genre[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].judul[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.judul[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].banyak.sumListener[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumListener[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].banyak.sumLikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumLikes[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].banyak.sumDislikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumDislikes[pilihan-1]
			Play_list.PL_ke[Play_list.banyak_PL].banyak.sumViews[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumViews[pilihan-1]
			fmt.Println()
			fmt.Println("========================================================")
			fmt.Println("==             ~Lagu Berhasil ditambahkan~            ==")
			fmt.Println("========================================================")
			banyak_lagu_di_setiap_PL[Play_list.banyak_PL]++
		} else {
			fmt.Println("========================================================")
			fmt.Println("==             Eror Gagal Menambahkan Lagu            ==")
			fmt.Println("========================================================")
		}
	}
}
func tambah_berdasar_rate(L *Lagu) {
	var ketemu bool
	var pilihan int
	sorting_rate(L)
	fmt.Println("========================================================")
	fmt.Println("==                     Top 5 Rate                     ==")
	fmt.Println("========================================================")
	for i := 0; i < 5; i++ {
		fmt.Printf("%-3d. %-23s %-17s Rating %-10.2f\n", i+1, L.judul[i], L.artist[i], L.banyak.sumRate[i])
	}
	fmt.Print("Pilih lagu yang ingin dipilih (1/2/3/4/5/6) 6.Untuk Kembali ")
	fmt.Scan(&pilihan)
	if pilihan <= 5 {
		for j := 0; j < banyak_lagu_di_setiap_PL[Play_list.banyak_PL] && pilihan > 0 && pilihan <= 5; j++ {
			if Play_list.PL_ke[Play_list.banyak_PL].artist[j] == L.artist[pilihan-1] && Play_list.PL_ke[Play_list.banyak_PL].genre[j] == L.genre[pilihan-1] && Play_list.PL_ke[Play_list.banyak_PL].judul[j] == L.judul[pilihan-1] {
				ketemu = true
			}
		}
		if ketemu {
			fmt.Println("========================================================")
			fmt.Println("==          Lagu Sudah ada di dalam Play List         ==")
			fmt.Println("========================================================")
		} else {
			if pilihan > 0 && pilihan <= 5 {
				Play_list.PL_ke[Play_list.banyak_PL].artist[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.artist[pilihan-1]
				Play_list.PL_ke[Play_list.banyak_PL].genre[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.genre[pilihan-1]
				Play_list.PL_ke[Play_list.banyak_PL].judul[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.judul[pilihan-1]
				Play_list.PL_ke[Play_list.banyak_PL].banyak.sumListener[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumListener[pilihan-1]
				Play_list.PL_ke[Play_list.banyak_PL].banyak.sumLikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumLikes[pilihan-1]
				Play_list.PL_ke[Play_list.banyak_PL].banyak.sumDislikes[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumDislikes[pilihan-1]
				Play_list.PL_ke[Play_list.banyak_PL].banyak.sumViews[banyak_lagu_di_setiap_PL[Play_list.banyak_PL]] = L.banyak.sumViews[pilihan-1]
				fmt.Println("========================================================")
				fmt.Println("==             ~Lagu Berhasil ditambahkan~            ==")
				fmt.Println("========================================================")
				banyak_lagu_di_setiap_PL[Play_list.banyak_PL]++
			}
		}
	} else if pilihan != 6 {
		fmt.Println("========================================================")
		fmt.Println("==               Eror Silahkan Coba Lagi              ==")
		fmt.Println("========================================================")
	}
}
func main() {
	var account Akun
	var song Lagu
	login(&account)
	pem_menu_utama(&song)
}
