package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int32, p1 int) int8 {
    switch "" {
    case "hello":
        switch "\n\t" {
        case "foo bar":
            v7 := []int{128, p1}
            _ = v7
        case "":
            for fi8 := 0; fi8 < 14; fi8++ {
                for fi9 := 0; fi9 < 8; fi9++ {
                    v10 := []int{}
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = fi9
                    _ = v10
                }
                var v11 uint = 255
                fmt.Println(p0)
                _ = fi8
                _ = v11
            }
        default:
            var v12 uint = 2
            _ = v12
        }
    default:
        if (!(!false)) {
            var v13 uint64 = 2
            _ = v13
        }
    }
    p0 -= (p0 % ((-300) | 1))
    v14 := GenS1{F0: 42, F1: 1}
    _ = p0
    _ = p1
    _ = v14
    return 1
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]uint32)
    var v3 uint16 = 511
    var v4 uint32 = 0
    var v5 uint = 356
    var v6 float64 = (1e10 / 1e10)
    v2[("hello" + "\n\t")] = v4
    var v15 int = 127
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 uint32 = 387
        fmt.Println(v15, v15, v4)
        v6 += 1.5
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: v3, F1: v1}
    if true {
        var v19 bool = true
        _ = v19
    } else {
        v3 -= 255
        fmt.Println(v5, v15, v4)
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 float32 = (3.14159 * float32((1e10 * 0.1)))
        _ = fi20
        _ = v21
    }
    if true {
        switch -2147483648 {
        case 288:
            v15 = (v15 / ((v15) | 1))
        default:
            if true {
                var v22 uint32 = v4
                v23 := GenS1{F0: v3, F1: v1}
                if (1.5 > (-2.25 * ((0.1 + v6) * v6))) {
                    v6 = 1.5
                    v24 := GenS1{F0: (v3 % (((^(v3))) | 1)), F1: (^(v1))}
                    for fi25 := 0; fi25 < 14; fi25++ {
                        v15 = (-(v15))
                        _ = fi25
                    }
                    _ = v24
                }
                _ = v22
                _ = v23
            }
        }
    }
    fmt.Println(v5, v3, v4)
    var v26 int16 = 1000
    v27 := []int{(v15 * v15)}
    v28 := GenS1{F0: v3, F1: (v1 + v1)}
    v6 = (((v6 * v6) + v6) / 3.14159)
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    v1 += v1
    v3 = (v3 + (v3 % (((v3 & (v3 + 255))) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
