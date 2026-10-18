package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint16, p1 uint32) int8 {
    v7 := []int{4294967295, -963, 42}
    p0 = (p0 * 1)
    p0 += (^(p0))
    v7 = append(v7, 466)
    var v8 float64 = (-2.25 - (0.0 - (0.1 + float64((1.5 + 1.5)))))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -81
}
func genFn1(p0 uint) int8 {
    if false {
        var v9 int = -2147483648
        _ = v9
    } else {
        switch -1 {
        case 127:
            v10 := []int{}
            _ = v10
        case -2147483648:
            v11 := GenS1{F0: 2147483647, F1: 0}
            _ = v11
        }
        var v12 int8 = 59
        _ = v12
    }
    p0 = p0
    fmt.Println(p0)
    v13 := []int{128, -706, -1}
    fmt.Println(p0)
    _ = p0
    _ = v13
    return -128
}

func main() {
    var v1 float64 = 0.1
    var v2 uint32 = 4294967295
    var v3 uint16 = 307
    var v4 float64 = v1
    var v5 float32 = 1e10
    v6 := make(map[string]bool)
    var v14 uint = 18446744073709551615
    var v15 float64 = (((v1 - v4) * v1) + 0.1)
    if false {
        if (4294967295 < v14) {
            switch (("hello" + "\n\t") + "foo bar") {
            case "foo bar":
                v16 := []int{42}
                _ = v16
            case "a":
                if false {
                    var v17 int16 = 1
                    _ = v17
                }
            case "":
                v18 := GenS1{F0: 432, F1: 42}
                _ = v18
            default:
                fmt.Println(v5, v15)
            }
        }
    }
    fmt.Println(v1, v2)
    v2 = v2
    v5 += 0.0
    v2 = (v2 / ((v2) | 1))
    var v19 float64 = (1.5 * float64((((v5 - v5) - float32(0.0)) + (v5 * v5))))
    v1 = v4
    v20 := GenS1{F0: 292, F1: 127}
    for fi21 := 0; fi21 < 6; fi21++ {
        switch (fi21 & (^(fi21))) {
        case 128:
            fmt.Println(v5, v14, v14)
        default:
            v22 := []int{1, (^(fi21)), fi21, fi21}
            _ = v22
        }
        var v23 uint16 = (v3 / (((v3 | v3)) | 1))
        _ = fi21
        _ = v23
    }
    _ = len(v6)
    v1 -= -2.25
    v24 := []int{-2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
