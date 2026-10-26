package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    var v3 uint16 = 1
    var v4 uint = 42
    v5 := []int{-9223372036854775808, 897, -940}
    var v6 float64 = (1e10 / 1.0)
    v7 := GenS1{F0: 9223372036854775807, F1: v3}
    var v8 float64 = float64(0.0)
    switch -103 {
    case 9223372036854775807:
        v1 = (640 < uint32(v4))
    case 42:
        var v9 uint8 = 84
        _ = v9
    case 736:
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 int64 = 1
            var v12 int8 = 83
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    var v13 float32 = (float32(v8) / 1.0)
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    if (v1 || (!v1)) {
        if v1 {
            fmt.Println(v3, v8)
            switch -568 {
            case 9223372036854775807:
                fmt.Println(v4)
            case 842:
                v14 := GenS2{F0: 823}
                _ = v14
            default:
                v15 := GenS2{F0: 0}
                _ = v15
            }
            v1 = (!(!(!(127 != 2))))
        }
        fmt.Println(v4, v6)
    }
    v8 = float64(0.1)
    if (1 != 104) {
        if (v1 || (v13 == ((v13 * (v13 / v13)) + v13))) {
            for fi16 := 0; fi16 < 6; fi16++ {
                var v17 uint8 = 177
                if v1 {
                    if v1 {
                        if (v4 < (^(v4))) {
                            v13 += float32(v8)
                        }
                    }
                    v3 -= v3
                }
                v2[32768] = 868
                _ = fi16
                _ = v17
            }
        }
        var v18 uint8 = 54
        _ = v18
    } else {
        v3 -= (v3 - v3)
    }
    var v19 bool = ((255 >= 128) && (v1 || (!(!false))))
    v19 = true
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
