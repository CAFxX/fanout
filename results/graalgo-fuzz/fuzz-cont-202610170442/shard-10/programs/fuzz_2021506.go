package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint16, p1 string) int {
    if false {
        fmt.Println(p1, p1)
        v6 := GenS1{F0: 381, F1: 42, F2: 563}
        _ = v6
    } else {
        fmt.Println(p0, p1)
        var v7 int8 = 42
        _ = v7
    }
    var v8 uint16 = p0
    if false {
        p0 -= v8
    } else {
        v9 := GenS1{F0: -2147483648, F1: 1, F2: 2}
        v10 := []int{-396, 0, -243}
        v11 := GenS1{F0: 2147483647, F1: 1, F2: 4294967295}
        _ = v10
        _ = v11
        _ = v9
    }
    p0 = p0
    if (32767 == 42) {
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 uint16 = (p0 % ((p0) | 1))
            p1 = "hello"
            switch -138 {
            case 52:
                var v14 int8 = 1
                _ = v14
            default:
                var v15 int16 = 1
                _ = v15
            }
            _ = fi12
            _ = v13
        }
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    _ = v8
    return 32768
}
func genFn1(p0 uint64, p1 uint64, p2 uint8) uint32 {
    var v16 int8 = 62
    var v17 uint = 15
    var v18 uint8 = p2
    v19 := GenS1{F0: -2147483648, F1: 624, F2: v17}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    return 4294967295
}

func main() {
    var v1 uint16 = 32767
    var v2 uint16 = 32767
    var v3 int32 = 2147483647
    var v4 uint64 = 42
    v5 := make(map[int]string)
    var v20 float64 = float64((-2.25 / 1.0))
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 bool = (-1 <= v3)
        v22 = (int(v1) < (fi21 - (fi21 << ((469) & 63))))
        _ = fi21
        _ = v22
    }
    v20 -= float64((0.1 + 0.1))
    for fi23 := 0; fi23 < 5; fi23++ {
        v24 := GenS1{F0: v3, F1: -933, F2: 255}
        if false {
            var v25 float64 = v20
            v20 = ((v25 + 0.1) / v20)
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    var v26 float32 = (3.14159 * -2.25)
    v27 := GenS2{F0: v2}
    var v28 int32 = v3
    switch 2147483647 {
    case 4294967295:
        fmt.Println(v20, v4, v4)
    case -281:
        for fi29 := 0; fi29 < 14; fi29++ {
            v5[127] = ("hello" + "\n\t")
            var v30 string = "foo bar"
            _ = fi29
            _ = v30
        }
    default:
        v31 := []int{-971}
        _ = v31
    }
    var v32 uint = 555
    v5[2147483647] = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
