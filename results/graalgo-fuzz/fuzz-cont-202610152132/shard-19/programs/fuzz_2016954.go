package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 string }
func genFn0(p0 int64, p1 uint32) uint {
    v7 := GenS1{F0: -128, F1: true}
    for fi8 := 0; fi8 < 19; fi8++ {
        if true {
            v9 := []int{(fi8 / (((fi8 & (fi8 + fi8))) | 1))}
            v10 := GenS1{F0: int8(p0), F1: true}
            if false {
                p0 = p0
                var v11 bool = (true && (p1 != p1))
                if (51 <= 4294967295) {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    fmt.Println(p0, p1, fi8)
                }
                _ = v11
            } else {
                p1 = 889
                var v12 int16 = 221
                var v13 int32 = -2147483648
                _ = v12
                _ = v13
            }
            _ = v10
            _ = v9
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]string)
    var v3 int8 = -128
    var v4 uint = 2
    v5 := make(map[uint32]bool)
    v6 := make(map[int]string)
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 uint = v4
        fmt.Println(fi14, v3)
        _ = fi14
        _ = v15
    }
    var v16 uint16 = 32767
    switch -9223372036854775808 {
    case -771:
        if len(v1) > 0 {
            v1[0] = 1
        }
    }
    v17 := GenS1{F0: 1, F1: (2147483647 != 0)}
    v5[1] = false
    for fi18 := 0; fi18 < 20; fi18++ {
        v19 := GenS1{F0: 50, F1: true}
        v20 := GenS1{F0: -1, F1: true}
        _ = fi18
        _ = v19
        _ = v20
    }
    v3 -= (^(v3))
    switch "\n\t" {
    case "a":
        var v21 int8 = v3
        _ = v21
    case "\n\t":
        var v22 uint16 = v16
        _ = v22
    default:
        fmt.Println(v3)
    }
    v23 := GenS1{F0: v3, F1: false}
    if (32767 == 9223372036854775807) {
        v24 := GenS2{F0: ("foo bar" + ("" + ("foo bar" + "hello")))}
        _ = v24
    }
    var v25 int16 = 32767
    v26 := GenS2{F0: ""}
    v27 := GenS1{F0: -128, F1: (426 <= 1)}
    fmt.Println(v16, v4)
    v16 = v16
    if (false && false) {
        v16 -= v16
        v5[544] = (!(false && false))
        var v28 bool = (((0.1 + float64(((3.14159 * 3.14159) / (0.0 - (0.1 + ((((1.5 * 3.14159) - float32(-2.25)) / 1.0) * -2.25)))))) * 3.14159) == 0.0)
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
