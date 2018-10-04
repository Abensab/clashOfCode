package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)
    for i := 0; i < N; i++ {
        var X, Y, Z int
        fmt.Scan(&X, &Y, &Z)
        if (X<=Y && X>=Z) || (X>=Y &&X<=Z) {
            fmt.Println(X)
        }else if (Y<=X && Y>=Z) || (Y>=X &&Y<=Z) {
            fmt.Println(Y)
        }else{
            fmt.Println(Z)
        }
    }
}
