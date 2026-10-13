package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 bool = (-1 == 0)
    var v2 float64 = (0.1 * 3.14159)
    var v3 int64 = 861
    v4 := []uint32{}
    var v5 uint8 = 42
    var v6 int32 = -2147483648
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 uint16 = 1
        _ = fi9
        _ = v10
    }
    var v11 int16 = 32767
    var v12 int64 = v3
    _ = len(v4)
    if v1 {
        var v13 string = "a"
        _ = v13
    }
    fmt.Println(v12, v6, v2)
    switch 1 {
    case -724:
        fmt.Println(v6, v12, v12)
    case -1:
        fmt.Println(v5, v12, v3)
    case 877:
        for fi14 := 0; fi14 < 8; fi14++ {
            switch (fi14 >> ((0) & 63)) {
            case 32767:
                v1 = (65535 >= 0)
            case 42:
                v3 = (v12 | (v3 | v3))
            default:
                switch -9223372036854775808 {
                case 42:
                    v5 += v5
                case 2147483647:
                    v11 += (v11 | v11)
                default:
                    fmt.Println(v6, v12, fi14)
                }
            }
            var v15 uint16 = 24
            _ = fi14
            _ = v15
        }
    default:
        v11 = (v11 & (v11 % ((v11) | 1)))
    }
    v5 += (v5 + (v5 | v5))
    _ = len(v4)
    var v16 int = int(v6)
    var v17 int8 = 121
    v18 := GenS1{F0: 46, F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
