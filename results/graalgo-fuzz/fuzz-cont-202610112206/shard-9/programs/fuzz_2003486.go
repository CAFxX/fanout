package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint64 }
func genFn0() int16 {
    if (false && ((0.1 - (1.5 * -2.25)) == 0.0)) {
        var v6 uint8 = 127
        switch 1 {
        case -390:
            v7 := []int{-1, 481, 379}
            _ = v7
        case -110:
            if false {
                v6 += v6
                if (797 != 65535) {
                    v6 = (v6 ^ 89)
                    v8 := GenS1{F0: 65535, F1: true}
                    v6 = v6
                    _ = v8
                }
                var v9 float64 = float64(0.1)
                _ = v9
            }
        default:
            if (0 != 32767) {
                fmt.Println(v6)
                v10 := []int{-1, 127, 127}
                var v11 uint = 2
                _ = v10
                _ = v11
            } else {
                v12 := GenS1{F0: 396, F1: (42 > 2)}
                v6 = 2
                v13 := []int{-2147483648, 42}
                _ = v12
                _ = v13
            }
        }
        _ = v6
    } else {
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 int16 = 0
            var v16 bool = ((true || (1.5 < ((0.1 - 3.14159) - (float32(1.5) * -2.25)))) || (false && true))
            v16 = v16
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    var v17 float32 = 3.14159
    _ = v17
    return 714
}

func main() {
    v1 := []string{}
    var v2 int32 = -442
    v3 := []string{"hello", "hello", "foo bar", ("" + (("" + "") + "\n\t"))}
    var v4 int16 = 660
    var v5 int16 = v4
    v18 := GenS2{F0: 192}
    for fi19 := 0; fi19 < 20; fi19++ {
        v5 = (v5 / (((v4 >> ((2) & 63))) | 1))
        _ = fi19
    }
    fmt.Println(v4, v5)
    fmt.Println(v2, v4)
    for fi20 := 0; fi20 < 8; fi20++ {
        if true {
            fmt.Println(v5, v4)
            v21 := GenS2{F0: 255}
            v5 -= (v5 | v5)
            _ = v21
        }
        var v22 uint16 = 65535
        _ = fi20
        _ = v22
    }
    var v23 uint16 = 750
    for fi24 := 0; fi24 < 10; fi24++ {
        var v25 int32 = -891
        _ = fi24
        _ = v25
    }
    v26 := GenS3{F0: 42, F1: v2}
    if false {
        v27 := GenS4{F0: (v2 + (^(v2))), F1: 2, F2: 557}
        v28 := GenS3{F0: -8, F1: (v2 - v2)}
        _ = v27
        _ = v28
    }
    switch -268 {
    case 127:
        var v29 uint16 = (v23 / ((v23) | 1))
        _ = v29
    default:
        v2 = (v2 - -826)
    }
    if (!(!(3.14159 <= 0.1))) {
        v23 -= v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
