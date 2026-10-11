package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint16; F2 string }

func main() {
    var v1 uint32 = 534
    var v2 uint64 = 2
    var v3 bool = (!true)
    var v4 float64 = 0.0
    v5 := []uint32{v1}
    var v6 uint64 = (v2 - (v2 | (^(v2))))
    v7 := GenS1{F0: 1, F1: v1}
    fmt.Println(v6)
    v5 = append(v5, uint32(v2))
    fmt.Println(v6, v3)
    if (false && (255 != 2)) {
        for fi8 := 0; fi8 < 8; fi8++ {
            fmt.Println(v1, v4)
            _ = fi8
        }
        v2 = (v2 - v6)
    }
    var v9 int32 = 433
    if false {
        v9 -= 2147483647
    } else {
        var v10 uint8 = 88
        switch "a" {
        case "\n\t":
            var v11 float64 = -2.25
            _ = v11
        }
        switch "" {
        case "a":
            var v12 float64 = float64((-2.25 - float32((v4 - v4))))
            _ = v12
        default:
            var v13 float32 = 3.14159
            _ = v13
        }
        _ = v10
    }
    var v14 uint8 = 127
    if (!v3) {
        var v15 uint16 = uint16(v6)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v15
    } else {
        if len(v5) > 0 {
            v5[0] = 0
        }
        v2 -= (v2 | v6)
    }
    fmt.Println(v3, v2, v4)
    v16 := GenS2{F0: -1}
    v14 -= v14
    v4 = v4
    fmt.Println(v14)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (v3 && ((v1 ^ v1) != v1)) {
        v17 := GenS3{F0: "hello", F1: 2, F2: ""}
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
