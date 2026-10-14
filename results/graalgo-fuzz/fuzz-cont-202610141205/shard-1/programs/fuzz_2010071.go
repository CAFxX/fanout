package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int8; F1 uint64; F2 int32 }
type GenS4 struct { F0 int8; F1 int8; F2 uint8 }
func genFn0() int64 {
    if (!(!false)) {
        var v6 uint16 = 2
        _ = v6
    } else {
        v7 := []int{}
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 uint8 = 2
            _ = fi8
            _ = v9
        }
        var v10 int16 = 0
        _ = v10
        _ = v7
    }
    var v11 float32 = 1.5
    _ = v11
    return -794
}

func main() {
    var v1 bool = (!(((!false) || false) && (255 < 42)))
    var v2 uint64 = 42
    v3 := make(map[string]string)
    var v4 int8 = -128
    var v5 float64 = float64(float32(0.0))
    switch -709 {
    case 1:
        v12 := GenS1{F0: v4, F1: 0}
        _ = v12
    default:
        v1 = v1
    }
    v13 := GenS1{F0: 42, F1: 65535}
    fmt.Println(v2)
    v2 = (v2 ^ (v2 >> ((4294967295) & 63)))
    v5 -= v5
    var v14 uint8 = 127
    if false {
        if v1 {
            v15 := GenS2{F0: int16(v14), F1: 0}
            v14 -= (v14 << ((0) & 63))
            if (!((!(750 > 65535)) && v1)) {
                v16 := v3["foo bar"]
                v17 := []int{-1, 2147483647}
                _ = v16
                _ = v17
            } else {
                for fi18 := 0; fi18 < 20; fi18++ {
                    v14 = (v14 + (v14 << ((110) & 63)))
                    _ = fi18
                }
            }
            _ = v15
        } else {
            v19 := GenS3{F0: (^(v4)), F1: v2, F2: 0}
            v20 := GenS4{F0: v4, F1: v4, F2: v14}
            var v21 uint32 = 1
            _ = v19
            _ = v20
            _ = v21
        }
        fmt.Println(v14, v5, v2)
        v1 = ((988 <= 2147483647) && (!(v5 >= 0.0)))
    }
    fmt.Println(v2, v2)
    v2 = (v2 * (^(v2)))
    var v22 int16 = 736
    fmt.Println(v22, v14, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
