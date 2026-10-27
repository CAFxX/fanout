package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint16) uint8 {
    var v7 float32 = 1e10
    p0 = 0
    _ = p0
    _ = v7
    return 1
}
func genFn1(p0 int64, p1 uint64) uint8 {
    p0 -= p0
    var v8 int16 = int16(p1)
    var v9 int32 = 42
    switch (("" + "\n\t") + "foo bar") {
    case "\n\t":
        v10 := []int{246}
        _ = v10
    case "foo bar":
        var v11 uint = 0
        _ = v11
    default:
        var v12 float64 = (1e10 * ((0.1 / 3.14159) * 3.14159))
        _ = v12
    }
    var v13 int8 = -128
    _ = p0
    _ = p1
    _ = v13
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := []string{""}
    var v2 bool = false
    var v3 float64 = 1.5
    v4 := make(map[int]bool)
    v5 := make(map[uint32]string)
    var v6 uint = 65535
    if (false && (-619 >= 42)) {
        v2 = v2
        var v14 bool = (uint64(v6) <= 132)
        var v15 int32 = 268
        _ = v14
        _ = v15
    }
    if false {
        var v16 int16 = -32768
        v17 := GenS1{F0: v2, F1: (v16 * (-(v16)))}
        switch 42 {
        case 128:
            v3 -= float64(0.0)
        case 32768:
            var v18 int8 = -128
            _ = v18
        default:
            if ("" == "\n\t") {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if v2 {
                    delete(v5, 809)
                    _ = len(v1)
                }
                for fi19 := 0; fi19 < 17; fi19++ {
                    v6 -= v6
                    for fi20 := 0; fi20 < 1; fi20++ {
                        v2 = ((!false) && true)
                        var v21 uint8 = 171
                        v16 += (v16 & (v16 % ((1) | 1)))
                        _ = fi20
                        _ = v21
                    }
                    _ = fi19
                }
            } else {
                var v22 float32 = 0.0
                _ = v22
            }
        }
        _ = v16
        _ = v17
    } else {
        if (v2 || (1 >= 9223372036854775807)) {
            if (0.1 < float32((0.0 * (v3 + v3)))) {
                v3 = v3
                var v23 uint32 = 55
                _ = v23
            }
            v24 := v5[2147483647]
            _ = v24
        }
        fmt.Println(v2)
    }
    if (genFn0(693) >= 128) {
        var v25 int16 = 1
        v26 := GenS1{F0: v2, F1: v25}
        fmt.Println(v25, v25, v3)
        _ = v25
        _ = v26
    }
    if v2 {
        var v27 float32 = float32((v3 - ((1.5 * v3) * 0.0)))
        v2 = (genFn1(-710, 4294967295) > 255)
        v28 := GenS2{F0: 127}
        _ = v27
        _ = v28
    } else {
        v6 = (v6 + v6)
        var v29 string = "\n\t"
        var v30 int64 = 42
        _ = v29
        _ = v30
    }
    v1 = append(v1, "\n\t")
    if v2 {
        delete(v4, 0)
        if ("" > ("hello" + "hello")) {
            v6 += v6
            var v31 int32 = -1
            v2 = true
            _ = v31
        } else {
            var v32 bool = (1 < -1)
            _ = v32
        }
        v3 -= (float64((0.0 - 0.0)) / v3)
    }
    v3 = 1.5
    v33 := GenS1{F0: (4294967295 < 18446744073709551615), F1: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
