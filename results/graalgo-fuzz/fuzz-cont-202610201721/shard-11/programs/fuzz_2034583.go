package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 string; F2 int }
func genFn0(p0 int64) int {
    fmt.Println(p0)
    if (926 == 255) {
        if ((-2147483648 > 42) && (false && true)) {
            p0 = p0
            fmt.Println(p0)
        }
        v4 := GenS1{F0: -2147483648, F1: ((p0 / ((p0) | 1)) < (p0 % (((-(p0))) | 1))), F2: 496}
        _ = v4
    } else {
        var v5 uint8 = 218
        switch "" {
        case "foo bar":
            switch 81 {
            case 352:
                p0 = (p0 * (p0 | (-(p0))))
            case 86:
                v6 := []int{int(p0), -580, -356, -431}
                _ = v6
            case -1:
                for fi7 := 0; fi7 < 16; fi7++ {
                    v8 := []int{fi7, (fi7 * (^(fi7))), 127, fi7}
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    p0 -= 0
                    _ = fi7
                    _ = v8
                }
            }
        case "a":
            p0 -= (p0 * (^(p0)))
        case "\n\t":
            fmt.Println(p0, p0)
        }
        _ = v5
    }
    fmt.Println(p0)
    p0 -= -16
    _ = p0
    return 42
}
func genFn1() uint32 {
    if (2 == 368) {
        v9 := GenS2{F0: 130}
        var v10 uint32 = 312
        _ = v10
        _ = v9
    }
    var v11 int32 = 42
    _ = v11
    return 185
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    delete(v2, 625)
    v1 -= v1
    if ((!false) || (v1 > v1)) {
        var v12 int16 = -552
        v13 := GenS1{F0: -994, F1: true, F2: 1}
        _ = v12
        _ = v13
    }
    v1 = (-(v1))
    var v14 uint8 = 128
    v14 = (v14 ^ (v14 ^ (v14 >> ((249) & 63))))
    switch "" {
    case "hello":
        var v15 float64 = (1.5 / (0.1 - (0.1 - 1e10)))
        _ = v15
    }
    var v16 uint32 = 2147483648
    switch -559 {
    case 32767:
        v17 := []int{-111}
        _ = v17
    case -9223372036854775808:
        v1 = v1
    default:
        var v18 int64 = 0
        _ = v18
    }
    var v19 float64 = 1e10
    v20 := v2[605]
    var v21 float64 = (v19 + v19)
    v20 = ((v20 + v20) + ("foo bar" + (v20 + v20)))
    var v22 float64 = v21
    switch v20 {
    case "hello":
        v23 := GenS3{F0: 255, F1: "", F2: 2147483647}
        _ = v23
    case "":
        var v24 float32 = 3.14159
        _ = v24
    }
    v25 := []int{128}
    v19 += v19
    for fi26 := 0; fi26 < 14; fi26++ {
        v22 += 3.14159
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
}
