package main
	
	import "fmt"
	
	type Film struct{
	    Judul string;
	    Genre string;
	    Tahunrilis int;
	}
	
	func main(){
	 var kumpulan []Film
	
	    kumpulan = []Film{
	     Film{
	         Judul : "ironman",
	         Genre : "action",
	         Tahunrilis : 2008,
	     },
	     Film{
	         Judul : "thor",
	         Genre : "action",
	         Tahunrilis : 2011,
	     },
	     Film{
	         Judul : "train to busan",
	         Genre : "horror",
	         Tahunrilis : 2016,
	     },
	     Film{
	         Judul : "peninsula",
	         Genre : "action",
	         Tahunrilis : 2020,
	     },
	     Film{
	         Judul : "Mulan",
	         Genre : "action",
	         Tahunrilis : 2020,
	     },
	     Film{
	         Judul : "toy story",
	         Genre : "adventure",
	         Tahunrilis : 1995,
	     },
	     Film{
	         Judul : "cars",
	         Genre : "animation",
	         Tahunrilis : 2006,
	     },
	     Film{
	         Judul : "frozen II",
	         Genre : "horror",
	         Tahunrilis : 2016,
	     }, 
	     Film{
	         Judul : "aladdin",
	         Genre : "adventure",
	         Tahunrilis : 2019,
	     }, 
	 }
	
	 fmt.Println("-------Data Film--------")
	 fmt.Println(kumpulan[0])
	 fmt.Println(kumpulan[1])
	 fmt.Println(kumpulan[2])
	 fmt.Println(kumpulan[3])
	 fmt.Println(kumpulan[4])
	 fmt.Println(kumpulan[5])
	 fmt.Println(kumpulan[6])
	 fmt.Println(kumpulan[7])
	 fmt.Println(kumpulan[8])
	 fmt.Println(kumpulan[9])
	 
	}

