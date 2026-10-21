package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
func genFn0() int {
    var v6 uint16 = 255
    v6 -= (v6 % (((v6 << ((255) & 63))) | 1))
    _ = v6
    return 42
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 string = "foo bar"
    var v3 int16 = 1
    var v4 uint32 = 1
    var v5 int64 = v1
    var v7 float64 = ((3.14159 - 0.1) * 1.5)
    switch 2147483647 {
    case -9223372036854775808:
        v4 = (^(v4))
    case 32768:
        v2 = v2
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        v7 = (float64(0.0) / 1.0)
        v3 += v3
        v5 -= (v5 ^ 70)
        _ = fi8
    }
    v9 := []int{}
    fmt.Println(v7, v4)
    switch v2 {
    case "foo bar":
        if len(v9) > 0 {
            v9[0] = 127
        }
    case "\n\t":
        var v10 int32 = 193
        _ = v10
    case "":
        var v11 uint16 = 806
        _ = v11
    default:
        var v12 int32 = -135
        _ = v12
    }
    if (42 == 18446744073709551615) {
        _ = len(v9)
        var v13 int = 871
        if false {
            fmt.Println(v13, v1)
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        } else {
            v5 += v5
        }
        _ = v13
    } else {
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 float32 = 0.1
            _ = fi14
            _ = v15
        }
        if (int8(v4) <= 127) {
            fmt.Println(v4, v1, v7)
        } else {
            var v16 uint16 = 65535
            _ = v16
        }
        var v17 uint64 = 0
        _ = v17
    }
    var v18 string = ("\n\t" + v2)
    v19 := GenS1{F0: (((v2 + v18) + v2) + v2), F1: -1, F2: 255}
    switch v18 {
    case "":
        if true {
            switch "hello" {
            case "a":
                v9 = append(v9, genFn0())
            case "hello":
                v20 := GenS2{F0: 18446744073709551615, F1: (v1 << ((18446744073709551615) & 63))}
                _ = v20
            case "":
                var v21 int32 = 1
                _ = v21
            }
            var v22 uint64 = 749
            if len(v9) > 0 {
                v9[0] = 127
            }
            _ = v22
        }
    default:
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    v18 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
