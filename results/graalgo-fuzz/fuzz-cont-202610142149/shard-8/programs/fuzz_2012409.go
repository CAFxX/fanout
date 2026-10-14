package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int32; F1 int64; F2 int64 }
type GenS3 struct { F0 int64; F1 int8 }
func genFn0(p0 uint) bool {
    var v6 int16 = 256
    v6 = v6
    v7 := []int{-2147483648, 32767, 2147483647}
    if (!(false && (!false))) {
        p0 = p0
        v7 = append(v7, 4294967295)
        fmt.Println(p0)
    } else {
        p0 = 42
        for fi8 := 0; fi8 < 11; fi8++ {
            if true {
                if false {
                    v9 := GenS1{F0: p0, F1: 18446744073709551615, F2: 4294967295}
                    var v10 uint = p0
                    _ = v10
                    _ = v9
                } else {
                    v6 = (^(v6))
                    var v11 float32 = -2.25
                    var v12 uint16 = 489
                    _ = v11
                    _ = v12
                }
            }
            fmt.Println(p0, p0, p0)
            _ = fi8
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return (((("\n\t" + ("foo bar" + ("\n\t" + ""))) + "foo bar") + "") > ("" + "\n\t"))
}

func main() {
    var v1 uint = 304
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    var v4 int64 = 474
    v5 := make(map[int]bool)
    v13 := GenS1{F0: v1, F1: 684, F2: 4294967295}
    v14 := GenS1{F0: v1, F1: 1, F2: 1}
    v15 := []int{}
    if len(v15) > 0 {
        v15[0] = 4294967295
    }
    v1 += (^(v1))
    v16 := GenS2{F0: 2147483647, F1: (v4 - (v4 / ((1) | 1))), F2: (v4 & 9223372036854775807)}
    _ = len(v2)
    var v17 int8 = 1
    v18 := GenS3{F0: 1, F1: v17}
    switch 591 {
    case 42:
        var v19 uint = (^(v1))
        _ = v19
    case 942:
        for fi20 := 0; fi20 < 13; fi20++ {
            switch "foo bar" {
            case "\n\t":
                switch -2147483648 {
                case -187:
                    var v21 float64 = (1.5 + (((1e10 - float64((0.1 + 0.0))) * (0.0 - (3.14159 / 1e10))) / 3.14159))
                    _ = v21
                case -189:
                    var v22 uint = v1
                    _ = v22
                default:
                    for fi23 := 0; fi23 < 15; fi23++ {
                        var v24 uint8 = uint8(v1)
                        _ = fi23
                        _ = v24
                    }
                }
            default:
                if (v1 != v1) {
                    var v25 int16 = 652
                    v26 := GenS3{F0: v4, F1: (v17 * (v17 * v17))}
                    v25 -= v25
                    _ = v25
                    _ = v26
                }
            }
            fmt.Println(v4, v4)
            v4 = v4
            _ = fi20
        }
    default:
        var v27 uint8 = 255
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
