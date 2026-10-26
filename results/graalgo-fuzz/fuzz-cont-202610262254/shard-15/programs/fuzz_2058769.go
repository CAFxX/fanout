package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }

func main() {
    var v1 uint16 = 65535
    var v2 float64 = 1.5
    v3 := make(map[int]int)
    v2 += float64((1e10 * 1e10))
    delete(v3, 1)
    var v4 bool = true
    var v5 string = "hello"
    var v6 string = v5
    v4 = true
    switch 32767 {
    case -109:
        v2 -= (0.1 + v2)
    case 642:
        for fi7 := 0; fi7 < 6; fi7++ {
            v3[fi7] = (fi7 ^ (fi7 * (^(fi7))))
            v4 = (-68 <= 100)
            fmt.Println(v4, v1, v6)
            _ = fi7
        }
    }
    v8 := GenS1{F0: 236, F1: 127}
    var v9 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
