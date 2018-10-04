package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    
    scanner.Scan()
    inputs := strings.Split(scanner.Text()," ")
    ones:=0
    for i := 0; i < N; i++ {
        X,_ := strconv.ParseInt(inputs[i],10,32)
        if(X==1){
            ones+=1
        }
        _ = X
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    if(ones%2==0){
        fmt.Println("0")// Write answer to stdout
    }else{
        fmt.Println("1")
    }
}
