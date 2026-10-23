package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int32; F1 int8 }
func genFn0() uint16 {
    var v7 float64 = ((float64(-2.25) * (((1e10 / float64(3.14159)) + 3.14159) * 1e10)) + 0.0)
    v7 -= 1.5
    v8 := []int{-2147483648, 952, -2147483648}
    _ = v7
    _ = v8
    return 476
}

func main() {
    var v1 uint16 = 519
    var v2 uint32 = 4294967295
    v3 := make(map[string]bool)
    var v4 uint64 = 1
    var v5 float32 = 1e10
    var v6 int64 = 1
    v5 -= v5
    v6 -= v6
    switch 81 {
    case 128:
        v9 := GenS1{F0: 1, F1: v4}
        _ = v9
    }
    if (0 >= uint(v2)) {
        if true {
            fmt.Println(v4)
        } else {
            v5 -= 0.0
        }
        v6 -= 9223372036854775807
        switch int(v4) {
        case 406:
            fmt.Println(v5, v2, v6)
        case -2147483648:
            var v10 uint8 = 42
            _ = v10
        default:
            v11 := v3[""]
            _ = v11
        }
    } else {
        v12 := GenS1{F0: -155, F1: 2}
        var v13 string = "\n\t"
        if (82 > int8(v2)) {
            var v14 bool = false
            var v15 int64 = 42
            var v16 int = int(v4)
            _ = v14
            _ = v15
            _ = v16
        } else {
            v6 += (-(v6))
            v17 := []int{42, -287, -1}
            var v18 uint8 = 127
            _ = v17
            _ = v18
        }
        _ = v12
        _ = v13
    }
    v19 := GenS1{F0: int32(v1), F1: v4}
    var v20 int8 = 100
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 uint16 = v1
        _ = fi21
        _ = v22
    }
    var v23 float64 = float64(v5)
    v24 := GenS2{F0: 203, F1: 127}
    v25 := GenS2{F0: 2147483647, F1: v20}
    if true {
        switch 4294967295 {
        case 4294967295:
            v26 := v3[""]
            _ = v26
        default:
            var v27 string = ""
            _ = v27
        }
        v4 = 151
    }
    var v28 int = -317
    v29 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
