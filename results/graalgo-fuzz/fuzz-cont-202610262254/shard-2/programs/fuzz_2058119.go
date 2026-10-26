package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int32 }
func genFn0(p0 int64, p1 bool) int16 {
    var v6 uint64 = 42
    switch -478 {
    case 313:
        p0 -= (p0 - (p0 ^ p0))
    case 854:
        var v7 int32 = 1
        _ = v7
    case 32768:
        switch 550 {
        case 0:
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 float64 = ((-2.25 / 1.5) / 0.1)
                var v10 int8 = -23
                var v11 bool = (float32(v9) == -2.25)
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        default:
            var v12 uint32 = 575
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return -909
}
func genFn1(p0 uint64) int8 {
    var v13 bool = false
    p0 -= p0
    var v14 uint8 = 127
    v14 += 255
    _ = p0
    _ = v13
    _ = v14
    return 127
}

func main() {
    var v1 float64 = float64(3.14159)
    var v2 int8 = 116
    var v3 int = 32768
    var v4 int64 = 1
    var v5 uint = 475
    v15 := GenS1{F0: v4, F1: false, F2: 2147483647}
    var v16 uint16 = 427
    if true {
        var v17 uint = (v5 << ((v5) & 63))
        var v18 string = ("" + "")
        for fi19 := 0; fi19 < 1; fi19++ {
            fmt.Println(v4, v16, v18)
            v2 = v2
            if (672 > -1) {
                v20 := []int{(fi19 + fi19), (fi19 / (((fi19 + int(v4))) | 1))}
                _ = v20
            } else {
                var v21 uint = v17
                _ = v21
            }
            _ = fi19
        }
        _ = v17
        _ = v18
    }
    var v22 int16 = 0
    fmt.Println(v5, v5)
    v2 -= v2
    v22 += int16(v4)
    if false {
        for fi23 := 0; fi23 < 11; fi23++ {
            v24 := GenS1{F0: v4, F1: (1 <= 460), F2: 661}
            if (v2 <= (^(v2))) {
                if true {
                    v25 := GenS1{F0: (v4 * int64(v5)), F1: true, F2: 222}
                    _ = v25
                }
            }
            _ = fi23
            _ = v24
        }
    } else {
        v22 -= (v22 >> ((4294967295) & 63))
    }
    v26 := []int{1}
    v16 -= v16
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
