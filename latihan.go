package main
import "fmt"
func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
    var hasil int
    hasil = arrivalTime+delayedTime
    if hasil == 24{
        return 0
    }else if hasil > 24{
        hasil -=24
        return hasil
    }else{
        return hasil
    }
}
func main(){
	var arrivalTime, delayedTime int
	fmt.Scan(&arrivalTime, &delayedTime)
	hasil:= findDelayedArrivalTime(arrivalTime, delayedTime)
	fmt.Print(hasil)
}