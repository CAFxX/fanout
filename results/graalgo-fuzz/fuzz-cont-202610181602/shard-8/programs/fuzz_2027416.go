package main

import (
    "fmt"
)


func main() {
    var v1 string = "a"
    var v2 int16 = -974
    var v3 uint8 = 42
    var v4 int8 = 26
    v3 = 179
    v3 = (v3 % ((v3) | 1))
    v4 -= (v4 % ((127) | 1))
    v3 = 128
    fmt.Println(v1)
    fmt.Println(v3, v3)
    v2 -= (v2 / (((v2 / (((v2 & v2)) | 1))) | 1))
    v2 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
