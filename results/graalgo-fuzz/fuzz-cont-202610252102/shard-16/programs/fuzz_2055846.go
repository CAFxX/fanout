package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int; F1 uint16; F2 int32 }

func main() {
    var v1 float64 = ((-2.25 / 0.1) / -2.25)
    var v2 int8 = 100
    var v3 int64 = 1
    v4 := make(map[string]int)
    var v5 int64 = (v3 % ((0) | 1))
    if ((v1 / 1e10) < v1) {
        var v6 uint8 = 218
        v5 = (v5 / (((-(v5))) | 1))
        fmt.Println(v6)
        _ = v6
    } else {
        for fi7 := 0; fi7 < 18; fi7++ {
            v5 = -994
            _ = fi7
        }
    }
    v8 := GenS1{F0: 0, F1: 42}
    var v9 int64 = (^(v5))
    fmt.Println(v9, v1)
    v1 = (0.0 * (v1 - (((v1 + float64(1.5)) * 0.0) / 0.1)))
    v9 += (v9 + (v5 % (((v5 + v5)) | 1)))
    var v10 int16 = 42
    switch 32768 {
    case 32767:
        var v11 int8 = -1
        _ = v11
    case -347:
        v12 := GenS2{F0: 0, F1: 110, F2: 0}
        _ = v12
    }
    fmt.Println(v1, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
