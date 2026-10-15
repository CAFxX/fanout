package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS2 struct { F0 bool; F1 int16 }

func main() {
    var v1 uint = 1
    v2 := []uint32{53, 2147483647, 0}
    var v3 int16 = -650
    var v4 int64 = 227
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 458, F1: 545, F2: 1}
        if ((v4 ^ v4) == v4) {
            var v7 int = 1
            _ = v7
        }
        var v8 string = ("a" + "hello")
        _ = fi5
        _ = v6
        _ = v8
    }
    v3 = (v3 + v3)
    var v9 float32 = 3.14159
    var v10 float64 = 0.0
    switch ("" + "a") {
    case "hello":
        var v11 uint8 = 42
        _ = v11
    case "\n\t":
        switch -747 {
        case 32768:
            var v12 uint16 = 1
            _ = v12
        case -9223372036854775808:
            fmt.Println(v3, v10)
        }
    case "":
        v13 := GenS2{F0: false, F1: v3}
        _ = v13
    default:
        v10 -= v10
    }
    var v14 uint = (v1 & 760)
    if ((((v4 & v4) <= (v4 % ((v4) | 1))) || ((1 <= 18446744073709551615) || false)) || (((v1 | 401) <= v14) || false)) {
        var v15 int32 = -715
        var v16 float32 = v9
        var v17 float32 = 0.1
        _ = v15
        _ = v16
        _ = v17
    } else {
        if ((v4 / ((0) | 1)) < (v4 + -9223372036854775808)) {
            if (v10 != v10) {
                if true {
                    fmt.Println(v4, v10, v10)
                }
            }
            for fi18 := 0; fi18 < 8; fi18++ {
                v2 = append(v2, 461)
                v4 = v4
                v1 = v14
                _ = fi18
            }
            _ = len(v2)
        } else {
            v1 = (^(v1))
        }
        var v19 uint = (v1 / ((1) | 1))
        v2 = append(v2, 2147483647)
        _ = v19
    }
    v1 += v14
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
