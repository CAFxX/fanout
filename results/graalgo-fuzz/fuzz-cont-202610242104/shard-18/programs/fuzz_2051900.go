package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
func genFn0(p0 int8) uint16 {
    var v6 uint16 = 65535
    var v7 uint8 = 255
    v8 := []int{-369, 0, -2147483648, 1}
    var v9 int32 = -1
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: (v7 >> ((255) & 63)), F1: 583}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return v6
}
func genFn1() uint32 {
    var v12 uint16 = genFn0(127)
    var v13 float32 = 0.1
    _ = v12
    _ = v13
    return 2
}

func main() {
    var v1 uint32 = 1
    v2 := []int{32768, -9223372036854775808, 4294967295, 32767}
    var v3 int64 = -9223372036854775808
    v4 := []int{9223372036854775807, 984}
    var v5 uint8 = 255
    switch 0 {
    case -419:
        v5 -= v5
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v14 bool = true
    if len(v2) > 0 {
        v2[0] = 128
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS1{F0: (v5 - v5), F1: 9223372036854775807}
        _ = fi15
        _ = v16
    }
    _ = len(v4)
    var v17 float32 = 0.1
    v18 := GenS1{F0: v5, F1: (v3 + (v3 | v3))}
    v19 := GenS1{F0: (v5 + v5), F1: -9223372036854775808}
    v1 = genFn1()
    var v20 string = ""
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if v14 {
        _ = len(v4)
        fmt.Println(v20, v1, v20)
        if (v14 && (float64(v17) > 0.0)) {
            fmt.Println(v14)
        }
    }
    if false {
        v21 := GenS1{F0: 255, F1: v3}
        var v22 int = 1
        _ = v21
        _ = v22
    } else {
        fmt.Println(v1)
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 int16 = 475
            if v14 {
                var v25 string = ("a" + (v20 + v20))
                v26 := GenS1{F0: 255, F1: (v3 & (v3 | -9223372036854775808))}
                _ = v25
                _ = v26
            }
            var v27 uint = 0
            _ = fi23
            _ = v24
            _ = v27
        }
        if len(v2) > 0 {
            v2[0] = 9223372036854775807
        }
    }
    if false {
        v28 := GenS1{F0: (v5 * v5), F1: 9223372036854775807}
        v2 = append(v2, 2147483647)
        _ = v28
    } else {
        var v29 uint = 2
        v30 := GenS1{F0: 1, F1: (v3 - v3)}
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
