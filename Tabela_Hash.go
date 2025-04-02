package main

import "fmt"

func main() {
    hashTable := make(map[string]int)
    hashTable["João"] = 25
    hashTable["Maria"] = 30

    fmt.Println(hashTable["João"]) // Saída: 25
    fmt.Println(hashTable["Maria"]) // Saída: 30
}
