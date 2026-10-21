package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint8; F2 bool }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 string = ""
    fmt.Println(v3)
    var v4 float64 = 3.14159
    var v5 float64 = 1.5
    v6 := GenS1{F0: 4294967295, F1: 580}
    var v7 int = 127
    for fi8 := 0; fi8 < 1; fi8++ {
        v7 += (fi8 & v7)
        v7 += v7
        _ = fi8
    }
    var v9 bool = (2 <= 18446744073709551615)
    v10 := GenS2{F0: -1, F1: 42, F2: (255 >= 128)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
