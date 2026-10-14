package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int32; F1 bool; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int8; F2 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 2
    var v3 uint32 = 554
    var v4 uint32 = (v3 % (((v3 >> (((v2 / (((v2 * v2)) | 1))) & 63))) | 1))
    var v5 uint = v2
    var v6 uint8 = uint8(v2)
    for fi7 := 0; fi7 < 17; fi7++ {
        v5 += v5
        var v8 float32 = -2.25
        _ = fi7
        _ = v8
    }
    switch "a" {
    case "foo bar":
        v6 -= (v6 + 1)
    case "":
        for fi9 := 0; fi9 < 10; fi9++ {
            switch "\n\t" {
            case "hello":
                fmt.Println(v3, v4, v4)
            case "a":
                v10 := GenS1{F0: 42, F1: -2147483648}
                _ = v10
            case "\n\t":
                v6 = 42
            }
            switch 128 {
            case 1:
                if false {
                    if false {
                        fmt.Println(v6, fi9, v3)
                    }
                }
            case 97:
                for fi11 := 0; fi11 < 17; fi11++ {
                    v5 += v5
                    _ = fi11
                }
            default:
                for fi12 := 0; fi12 < 8; fi12++ {
                    var v13 uint16 = 1
                    fmt.Println(v6, fi9)
                    v6 += (^(v6))
                    _ = fi12
                    _ = v13
                }
            }
            v4 += v3
            _ = fi9
        }
    case "\n\t":
        v14 := v1[-331]
        _ = v14
    }
    fmt.Println(v2)
    for fi15 := 0; fi15 < 11; fi15++ {
        v16 := GenS1{F0: 2, F1: 0}
        v17 := GenS2{F0: 0, F1: (0.0 > -2.25), F2: (v6 * 42)}
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS3{F0: 1, F1: int8(v2), F2: 27}
    v6 -= 2
    fmt.Println(v6, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
