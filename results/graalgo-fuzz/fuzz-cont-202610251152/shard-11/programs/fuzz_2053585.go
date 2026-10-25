package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 2
    var v3 uint8 = 42
    var v4 float32 = 0.0
    v5 := []uint32{582, 42, (^(v2))}
    var v6 int8 = 127
    _ = len(v5)
    if (1 >= -1) {
        v7 := GenS1{F0: 209, F1: 42, F2: ""}
        fmt.Println(v2, v6, v4)
        if (true && true) {
            for fi8 := 0; fi8 < 18; fi8++ {
                fmt.Println(v4)
                _ = fi8
            }
        }
        _ = v7
    }
    var v9 float32 = (v4 / v4)
    if false {
        v10 := GenS1{F0: (v3 * (v3 / ((v3) | 1))), F1: 449, F2: "hello"}
        if true {
            fmt.Println(v3, v3, v4)
        } else {
            if len(v5) > 0 {
                v5[0] = 2147483648
            }
            switch 127 {
            case 128:
                v2 += (v2 ^ (v2 % ((v2) | 1)))
            case -1:
                v11 := GenS1{F0: v3, F1: 32767, F2: ""}
                _ = v11
            default:
                if (false || false) {
                    v4 += 0.0
                }
            }
        }
        _ = v10
    }
    var v12 bool = (-2.25 > v4)
    v13 := GenS1{F0: (v3 + (^(v3))), F1: 42, F2: "hello"}
    var v14 int64 = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
