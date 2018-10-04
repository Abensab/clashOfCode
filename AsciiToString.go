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

    var charCount int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&charCount)
    
    scanner.Scan()
    inputs := strings.Split(scanner.Text()," ")
    var cadena  string
    for i := 0; i < charCount; i++ {
        charCode,_ := strconv.ParseInt(inputs[i],10,32)
        character := string(charCode)
        cadena = cadena + string(character)
    } 
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    fmt.Println(cadena)// Write answer to stdout
}
