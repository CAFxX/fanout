package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS2 struct { F0 int64; F1 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 0
    v3 := make(map[int]bool)
    v4 := make(map[int]int)
    if (1 <= 0) {
        v2 += (v2 >> ((65535) & 63))
        switch "" {
        case "":
            if (100 == 0) {
                if false {
                    if (!true) {
                        v5 := GenS1{F0: 255, F1: 0, F2: 2147483648}
                        _ = v5
                    }
                    var v6 int16 = 42
                    _ = v6
                }
            } else {
                var v7 uint32 = 154
                _ = v7
            }
        case "a":
            v8 := v4[9223372036854775807]
            _ = v8
        case "\n\t":
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 float64 = 3.14159
                _ = fi9
                _ = v10
            }
        }
    }
    v2 -= v2
    switch 128 {
    case -1:
        fmt.Println(v2)
    case 2147483647:
        var v11 uint8 = 1
        _ = v11
    default:
        if (-1 < 1) {
            v2 = v2
            var v12 int = int(v2)
            var v13 bool = (1e10 >= 0.1)
            _ = v12
            _ = v13
        }
    }
    v2 = 2
    if true {
        v14 := GenS2{F0: -9223372036854775808, F1: 72}
        _ = v14
    }
    v2 -= v2
    v2 -= v2
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := []int{}
        _ = fi15
        _ = v16
    }
    fmt.Println(v2)
    fmt.Println(v2)
    v17 := GenS2{F0: 42, F1: 9223372036854775807}
    v2 = 940
    v18 := GenS1{F0: 128, F1: 52, F2: 1}
    for fi19 := 0; fi19 < 14; fi19++ {
        fmt.Println(fi19, fi19)
        v2 -= (v2 | 576)
        _ = fi19
    }
    fmt.Println(v2)
    for fi20 := 0; fi20 < 5; fi20++ {
        v21 := []int{}
        var v22 string = ""
        _ = fi20
        _ = v21
        _ = v22
    }
    v2 = (^(v2))
    var v23 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
