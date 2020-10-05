package main
	
	import "fmt"
	
	type Siswa struct{
	    Nama string;
	    Nourut int;
	    Sekolah string;
	}
	
	func main(){
		var mrd Siswa
		fmt.Println("MEMASUKAN INPUTAN NAMA SISWA")
	
		for i := 0; i < 6; i++{
		
		fmt.Println("Masukan Nama :")
		fmt.Scan(&mrd.Nama)
		fmt.Println("Masukan Nourut :")
		fmt.Scan(&mrd.Nourut)
		fmt.Println("Masukan Sekolah :")
		fmt.Scan(&mrd.Sekolah)
		}
	
	}
