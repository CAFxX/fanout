package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int }

func main() {
    var v1 int = 223
    v2 := make(map[string]int)
    var v3 float32 = 3.14159
    v4 := make(map[int]bool)
    v5 := GenS1{F0: 0}
    fmt.Println(v3, v1)
    var v6 float64 = -2.25
    fmt.Println(v6, v1, v1)
    switch "" {
    case "a":
        fmt.Println(v6, v6, v6)
    case "":
        v1 = v1
    }
    v7 := GenS2{F0: 29, F1: v1}
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            fmt.Println(fi9)
            _ = fi9
        }
        for fi10 := 0; fi10 < 15; fi10++ {
            switch "\n\t" {
            case "hello":
                v6 = v6
            }
            v3 -= (((-2.25 - float32(v6)) * v3) + v3)
            _ = fi10
        }
        fmt.Println(v1, fi8)
        _ = fi8
    }
    if false {
        var v11 uint = 65535
        var v12 float32 = (float32((((v6 - v6) / 1.0) - v6)) * (v3 / 1.0))
        _ = v11
        _ = v12
    }
    var v13 uint = 2
    fmt.Println(v3)
    fmt.Println(v3)
    var v14 uint16 = 32767
    v13 += v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
