package main

import "fmt"
import "strconv"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    var a, b int
    fmt.Scan(&a, &b)
    
    print(180-a-b)
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println(strconv.Itoa(180 -b-a))
}
