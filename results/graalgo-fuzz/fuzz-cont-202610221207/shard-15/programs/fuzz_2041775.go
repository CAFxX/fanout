package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int16 }
type GenS2 struct { F0 int8; F1 uint; F2 int32 }
func genFn0() int {
    v7 := []int{127}
    if false {
        var v8 int16 = -224
        _ = len(v7)
        var v9 int = 590
        _ = v8
        _ = v9
    } else {
        var v10 int = -2147483648
        if (1 <= 1) {
            fmt.Println(v10)
        } else {
            var v11 int64 = 88
            v12 := GenS1{F0: 1, F1: v10, F2: -688}
            fmt.Println(v11, v11)
            _ = v11
            _ = v12
        }
        var v13 int64 = 608
        _ = v10
        _ = v13
    }
    _ = v7
    return -1
}
func genFn1(p0 string, p1 uint, p2 bool) int32 {
    var v14 string = p0
    var v15 int8 = 22
    v15 = v15
    p0 = p0
    v16 := []int{-2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    return 0
}

func main() {
    v1 := []string{((("a" + "\n\t") + ("a" + "a")) + "a")}
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]int)
    var v4 int16 = -492
    v5 := make(map[string]bool)
    var v6 uint32 = 65535
    var v17 float64 = 1e10
    v4 += v4
    v18 := GenS1{F0: 121, F1: 32767, F2: (^(v4))}
    var v19 float64 = -2.25
    v20 := GenS1{F0: 211, F1: 369, F2: v4}
    switch "foo bar" {
    case "":
        for fi21 := 0; fi21 < 19; fi21++ {
            switch "hello" {
            case "foo bar":
                var v22 string = "a"
                _ = v22
            case "hello":
                for fi23 := 0; fi23 < 3; fi23++ {
                    if (128 < 128) {
                        v24 := GenS1{F0: 1, F1: (fi23 | 1), F2: v4}
                        v25 := GenS2{F0: 42, F1: 179, F2: 0}
                        fmt.Println(fi21, v2)
                        _ = v24
                        _ = v25
                    }
                    _ = fi23
                }
            case "\n\t":
                for fi26 := 0; fi26 < 19; fi26++ {
                    v27 := GenS1{F0: 218, F1: int(v2), F2: (v4 | 0)}
                    _ = fi26
                    _ = v27
                }
            }
            var v28 string = "hello"
            _ = fi21
            _ = v28
        }
    default:
        if (2147483647 != 172) {
            var v29 float64 = v19
            var v30 bool = (true || (!(0.1 >= 0.0)))
            _ = v29
            _ = v30
        }
    }
    var v31 uint = 4294967295
    var v32 float32 = 1e10
    for fi33 := 0; fi33 < 18; fi33++ {
        switch "a" {
        case "":
            var v34 int = 262
            _ = v34
        case "a":
            fmt.Println(v31)
        }
        v6 -= (v6 & (v6 / ((uint32(v31)) | 1)))
        _ = fi33
    }
    v19 -= v17
    switch 9223372036854775807 {
    case 32768:
        fmt.Println(v17)
    case 9223372036854775807:
        var v35 uint8 = 42
        _ = v35
    default:
        fmt.Println(v17)
    }
    v2 += uint64(v6)
    var v36 float64 = v19
    if true {
        v37 := GenS2{F0: -1, F1: (v31 - v31), F2: 1}
        v4 = (v4 * (v4 * 0))
        var v38 int = 0
        _ = v37
        _ = v38
    } else {
        var v39 int = -9223372036854775808
        v4 = v4
        v32 -= v32
        _ = v39
    }
    for fi40 := 0; fi40 < 11; fi40++ {
        switch 936 {
        case 2147483647:
            v17 += (1e10 / ((3.14159 - (float64(float32((((((float64(1.5) + v36) * 1.5) - 1e10) + (((1e10 + v19) * (0.1 + -2.25)) * v17)) * v36))) / v19)) * v17))
        }
        v19 = (v19 / -2.25)
        _ = fi40
    }
    var v41 uint8 = 128
    delete(v3, 372)
    var v42 uint64 = uint64(v31)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
