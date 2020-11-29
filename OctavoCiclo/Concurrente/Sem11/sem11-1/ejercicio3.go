package main
/*
import("fmt")

func envioNumeros(ch chan int){
	for i:=0;i<=10;i++{
		ch<-i
	}

	close(ch)
}


func main(){

	ch:=make(chan int)

	go envioNumeros(ch)

	for num:=range ch{
		fmt.Println(num)
	}

}
*/