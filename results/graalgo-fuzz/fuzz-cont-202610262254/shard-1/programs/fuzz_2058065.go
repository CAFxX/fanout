package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int }
type GenS2 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 uint = 433
    v2 := make(map[int]bool)
    var v3 uint8 = 2
    var v4 int8 = -128
    v5 := []int{}
    v6 := GenS1{F0: (v3 << (((v1 / ((905) | 1))) & 63)), F1: uint16(v1), F2: 9223372036854775807}
    v3 += 30
    switch "" {
    case "a":
        fmt.Println(v3, v3)
    case "\n\t":
        var v7 uint8 = 66
        _ = v7
    default:
        var v8 float64 = 1.5
        _ = v8
    }
    if true {
        v3 += v3
    } else {
        var v9 uint8 = 246
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: 56, F1: 255, F2: (^(fi10))}
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            if (3.14159 > -2.25) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                fmt.Println(fi12)
                v1 += (v1 >> ((42) & 63))
            }
            switch "foo bar" {
            case "":
                if true {
                    v3 = (v3 % (((v3 * v3)) | 1))
                }
            case "a":
                fmt.Println(v3, v4, v1)
            default:
                v13 := GenS2{F0: 736, F1: 2147483648}
                _ = v13
            }
            _ = fi12
        }
        _ = v9
    }
    var v14 uint64 = 2
    v4 = v4
    var v15 int64 = 645
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
