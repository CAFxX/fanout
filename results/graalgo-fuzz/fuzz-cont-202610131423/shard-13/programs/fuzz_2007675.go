package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }

func main() {
    var v1 float64 = 1.5
    var v2 int = -676
    var v3 uint32 = 2
    v4 := []int{}
    _ = len(v4)
    if len(v4) > 0 {
        v4[0] = (v2 ^ v2)
    }
    for fi5 := 0; fi5 < 16; fi5++ {
        switch fi5 {
        case 32768:
            var v6 uint8 = 24
            _ = v6
        default:
            var v7 float64 = v1
            _ = v7
        }
        _ = fi5
    }
    v8 := GenS1{F0: 2, F1: (v2 & v2), F2: false}
    var v9 uint64 = 493
    v3 = (^(v3))
    if (1 <= -887) {
        v4 = append(v4, v2)
        v4 = append(v4, v2)
        fmt.Println(v2, v3, v3)
    }
    if len(v4) > 0 {
        v4[0] = v2
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        if true {
            if (128 >= 1) {
                v11 := GenS2{F0: 42}
                v12 := GenS1{F0: (v9 + v9), F1: -9223372036854775808, F2: true}
                _ = v11
                _ = v12
            } else {
                v13 := GenS3{F0: 255}
                _ = v13
            }
            if false {
                v3 = (v3 ^ 42)
                for fi14 := 0; fi14 < 9; fi14++ {
                    v9 += (v9 / ((v9) | 1))
                    _ = fi14
                }
                v15 := GenS3{F0: 853}
                _ = v15
            } else {
                v9 = v9
            }
        }
        var v16 uint8 = 219
        _ = fi10
        _ = v16
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v17 bool = false
    var v18 uint32 = v3
    fmt.Println(v17, v1, v17)
    switch "\n\t" {
    case "":
        v1 = float64(1.5)
    case "hello":
        var v19 int8 = -95
        _ = v19
    default:
        if true {
            if (v17 && ((v18 | v3) != 2)) {
                switch -532 {
                case -605:
                    v1 += (0.1 - 3.14159)
                }
            }
            var v20 uint32 = v18
            _ = v20
        } else {
            if len(v4) > 0 {
                v4[0] = (v2 | v2)
            }
            var v21 bool = (-9223372036854775808 == 684)
            v2 -= (v2 - v2)
            _ = v21
        }
    }
    v3 = 2
    v22 := GenS1{F0: v9, F1: -2147483648, F2: (v9 <= 4294967295)}
    for fi23 := 0; fi23 < 14; fi23++ {
        v1 -= (0.0 * 1e10)
        v2 -= (fi23 | fi23)
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
