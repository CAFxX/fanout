package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint16; F1 int }

func main() {
    var v1 uint = 65535
    var v2 float64 = ((-2.25 / (((1.5 + 1.5) * 1.5) - (float64(3.14159) * -2.25))) / 1.5)
    var v3 int64 = -1
    v4 := []uint32{}
    var v5 int16 = 0
    var v6 int32 = 42
    switch "\n\t" {
    case "a":
        var v7 int16 = v5
        _ = v7
    case "foo bar":
        var v8 float64 = (((v2 - v2) / v2) - (float64(3.14159) + ((1e10 + (((float64((0.0 / 1e10)) + 0.1) - 1.5) + v2)) * 3.14159)))
        _ = v8
    case "":
        v1 += v1
    }
    v9 := GenS1{F0: 2147483647, F1: 4294967295, F2: (v6 & (v6 / (((v6 + v6)) | 1)))}
    v1 -= (v1 & (v1 / ((65535) | 1)))
    if true {
        var v10 int = -1
        v6 = v6
        _ = v10
    }
    var v11 uint8 = 128
    v4 = append(v4, 996)
    for fi12 := 0; fi12 < 4; fi12++ {
        v2 = (0.1 / v2)
        _ = fi12
    }
    if true {
        switch int(v3) {
        case 797:
            var v13 uint64 = 759
            _ = v13
        case -2147483648:
            v6 -= 42
        case 695:
            var v14 int16 = v5
            _ = v14
        }
        v11 += 98
    } else {
        if (v5 <= (v5 % ((v5) | 1))) {
            if (0 < v6) {
                fmt.Println(v5, v5, v11)
            }
        } else {
            v4 = append(v4, 65535)
            v3 = -806
            for fi15 := 0; fi15 < 14; fi15++ {
                _ = len(v4)
                v16 := GenS2{F0: 0, F1: (fi15 >> ((v1) & 63))}
                var v17 int8 = 127
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    }
    v6 = v6
    var v18 uint8 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
