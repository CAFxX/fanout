package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int16; F1 int32 }
func genFn0() int16 {
    var v4 int64 = 0
    var v5 int8 = 42
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 string = "hello"
    var v2 bool = false
    var v3 int = -991
    v1 = (v1 + v1)
    var v6 uint32 = 1
    fmt.Println(v1, v3)
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(v6)
        for fi8 := 0; fi8 < 6; fi8++ {
            v1 = v1
            _ = fi8
        }
        var v9 uint8 = 98
        _ = fi7
        _ = v9
    }
    var v10 string = ""
    v11 := GenS1{F0: -39, F1: 42, F2: (^(v6))}
    v10 = v10
    switch v3 {
    case 32768:
        v12 := []int{-9223372036854775808, v3, 0, (v3 ^ 0)}
        _ = v12
    case 2147483647:
        switch (v3 * (v3 >> ((399) & 63))) {
        case 127:
            if (3.14159 > (-2.25 + -2.25)) {
                v13 := GenS1{F0: 42, F1: 1, F2: 189}
                _ = v13
            }
        case -478:
            if false {
                v14 := GenS1{F0: -128, F1: 255, F2: (^(v6))}
                var v15 int16 = 0
                _ = v14
                _ = v15
            } else {
                var v16 int32 = -394
                var v17 uint8 = 2
                v18 := []int{v3, (v3 * 32767), 1}
                _ = v16
                _ = v17
                _ = v18
            }
        case -1:
            fmt.Println(v2)
        }
    case 0:
        v19 := GenS2{F0: 128, F1: 649}
        _ = v19
    default:
        for fi20 := 0; fi20 < 10; fi20++ {
            if v2 {
                switch fi20 {
                case 630:
                    var v21 uint16 = 317
                    _ = v21
                default:
                    var v22 int16 = genFn0()
                    _ = v22
                }
                var v23 bool = false
                var v24 uint8 = 1
                _ = v23
                _ = v24
            } else {
                v25 := []int{235, v3, 653, (v3 >> ((916) & 63))}
                v26 := GenS3{F0: 1000, F1: -1}
                if (float64(((1.5 - 1.5) - 0.0)) < ((3.14159 - ((0.1 + (-2.25 / 1.5)) + -2.25)) * 1e10)) {
                    var v27 float64 = (1e10 * 3.14159)
                    _ = v27
                }
                _ = v25
                _ = v26
            }
            v2 = (v2 || v2)
            v28 := []int{}
            _ = fi20
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        if false {
            for fi30 := 0; fi30 < 19; fi30++ {
                v31 := GenS2{F0: uint8(v6), F1: 0}
                v32 := []int{(fi29 - fi30), (v3 / ((4294967295) | 1)), fi30}
                v2 = (236 == 1)
                _ = fi30
                _ = v31
                _ = v32
            }
        }
        v6 = v6
        _ = fi29
    }
    var v33 string = (v1 + "a")
    for fi34 := 0; fi34 < 11; fi34++ {
        for fi35 := 0; fi35 < 3; fi35++ {
            v36 := GenS1{F0: 0, F1: 525, F2: 0}
            v37 := GenS1{F0: -1, F1: 2, F2: 326}
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi34
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        var v39 string = ((("" + v33) + "\n\t") + v33)
        _ = fi38
        _ = v39
    }
    if (float32(0.1) <= 0.1) {
        v3 = 9223372036854775807
        var v40 float32 = 1e10
        _ = v40
    }
    v41 := []int{(v3 ^ (v3 ^ 336)), (v3 << ((65535) & 63))}
    var v42 float64 = 0.0
    v42 = ((0.0 / v42) / v42)
    for fi43 := 0; fi43 < 19; fi43++ {
        v41 = append(v41, (v3 * v3))
        _ = fi43
    }
    v6 += v6
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v41
    _ = v42
    _ = v6
}
