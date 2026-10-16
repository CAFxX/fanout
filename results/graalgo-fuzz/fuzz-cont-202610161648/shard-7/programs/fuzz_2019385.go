package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int16 }
func genFn0() int {
    if ((2147483647 >= -778) || true) {
        v7 := []int{}
        _ = v7
    } else {
        if false {
            switch 1 {
            case 257:
                v8 := GenS1{F0: 1, F1: 4294967295}
                _ = v8
            case -2147483648:
                var v9 float32 = 3.14159
                _ = v9
            }
            var v10 int = 2147483647
            var v11 int32 = 2147483647
            _ = v10
            _ = v11
        }
        var v12 uint64 = 0
        switch 81 {
        case 128:
            var v13 float64 = float64(-2.25)
            _ = v13
        case 4294967295:
            v14 := []int{-1, int(v12), 0}
            _ = v14
        case 127:
            var v15 int16 = 42
            _ = v15
        default:
            var v16 int64 = 771
            _ = v16
        }
        _ = v12
    }
    v17 := []int{2147483647, 2147483647}
    var v18 bool = (616 > 509)
    _ = v17
    _ = v18
    return 128
}
func genFn1(p0 int8, p1 int64) uint {
    fmt.Println(p1)
    p0 += p0
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint32 = 118
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v4 bool = (255 <= 255)
    v5 := make(map[int]uint32)
    var v6 int8 = 42
    if v4 {
        var v19 int8 = 1
        var v20 uint16 = 2
        fmt.Println(v6)
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: 2, F1: 4294967295}
    switch 42 {
    case 32767:
        for fi22 := 0; fi22 < 7; fi22++ {
            fmt.Println(v1, v4)
            _ = fi22
        }
    case 0:
        switch -9223372036854775808 {
        case -1:
            var v23 int16 = 1
            _ = v23
        case 454:
            var v24 float32 = 1.5
            _ = v24
        default:
            var v25 int16 = 0
            _ = v25
        }
    case -663:
        v4 = true
    default:
        if (1e10 != 3.14159) {
            var v26 float64 = 1e10
            v27 := []int{}
            _ = v26
            _ = v27
        } else {
            var v28 string = "\n\t"
            v28 = v28
            v29 := []int{genFn0(), 2147483647}
            _ = v28
            _ = v29
        }
    }
    fmt.Println(v1)
    v30 := []int{genFn0()}
    if v4 {
        var v31 string = ""
        v31 = ("foo bar" + v31)
        var v32 int8 = -1
        _ = v31
        _ = v32
    }
    var v33 int8 = (v6 + (-(v6)))
    v34 := GenS2{F0: 154}
    if false {
        fmt.Println(v4)
        if true {
            v30 = append(v30, -892)
            v1 -= v1
            _ = len(v5)
        }
    }
    delete(v5, -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
