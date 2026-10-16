package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 bool; F1 string; F2 uint8 }
func genFn0(p0 uint16, p1 int) int8 {
    var v7 uint64 = 0
    var v8 int16 = 1
    v7 = (v7 & (v7 << ((442) & 63)))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -108
}
func genFn1() int8 {
    v9 := []int{4294967295, 2147483647}
    var v10 float32 = (-2.25 + 0.1)
    var v11 int64 = 0
    _ = v10
    _ = v11
    _ = v9
    return -128
}

func main() {
    var v1 float32 = 0.1
    var v2 int16 = 840
    var v3 uint64 = 338
    v4 := make(map[int]uint32)
    var v5 int64 = 0
    var v6 int8 = -84
    v12 := v4[-110]
    var v13 uint8 = 128
    v14 := GenS1{F0: v12, F1: 637}
    var v15 uint16 = 65535
    v13 = v13
    v1 = v1
    v12 = (v12 - 65535)
    v16 := v4[0]
    v1 += 1e10
    for fi17 := 0; fi17 < 1; fi17++ {
        if false {
            switch fi17 {
            case -2147483648:
                v1 = 0.0
            default:
                var v18 uint64 = v3
                _ = v18
            }
        }
        fmt.Println(v13, fi17)
        _ = fi17
    }
    v19 := GenS1{F0: v12, F1: 121}
    v1 = v1
    if (genFn1() < v6) {
        if (!false) {
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 float64 = 0.1
                _ = fi20
                _ = v21
            }
            if (true && true) {
                v22 := GenS2{F0: 888, F1: v2}
                switch ("" + "a") {
                case "\n\t":
                    if true {
                        var v23 float64 = (3.14159 / 3.14159)
                        _ = v23
                    }
                default:
                    if false {
                        var v24 int32 = -2147483648
                        var v25 int = 4294967295
                        _ = v24
                        _ = v25
                    }
                }
                _ = v22
            }
            v26 := GenS2{F0: 255, F1: v2}
            _ = v26
        } else {
            fmt.Println(v5, v15)
        }
        v27 := []int{806}
        _ = v27
    }
    v28 := GenS3{F0: (((!(-1 < -188)) && (v3 != v3)) || ((!((-2147483648 > -538) || (!false))) && (42 <= 2147483647))), F1: "", F2: (^(v13))}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
