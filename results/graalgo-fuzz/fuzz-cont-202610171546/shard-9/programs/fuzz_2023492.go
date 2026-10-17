package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0(p0 int, p1 uint16, p2 uint64) uint8 {
    fmt.Println(p2, p1, p0)
    v6 := GenS1{F0: 32767, F1: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 197
    var v3 float32 = (3.14159 / 3.14159)
    var v4 uint32 = 65535
    var v5 uint8 = 0
    var v7 int32 = 2147483647
    var v8 uint64 = (v2 / ((v2) | 1))
    fmt.Println(v3, v5, v3)
    fmt.Println(v2, v2)
    v9 := GenS1{F0: -79, F1: 0}
    if false {
        var v10 uint = 0
        _ = v10
    } else {
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 int64 = -662
            fmt.Println(v2, v4, v7)
            v13 := GenS2{F0: v2, F1: 42}
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 int8 = 84
        _ = v14
    }
    if (!((3.14159 * (3.14159 / -2.25)) > -2.25)) {
        fmt.Println(v3, v7)
        _ = len(v1)
    } else {
        v15 := []int{-2147483648}
        v8 = 2
        _ = v15
    }
    var v16 int32 = 474
    for fi17 := 0; fi17 < 13; fi17++ {
        var v18 uint16 = 42
        _ = fi17
        _ = v18
    }
    v19 := v1[("" + (("a" + "a") + ""))]
    v19 = "a"
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 int = 768
        _ = fi20
        _ = v21
    }
    v7 = (v16 + (-(v16)))
    switch "a" {
    case "\n\t":
        fmt.Println(v4, v8)
    default:
        for fi22 := 0; fi22 < 17; fi22++ {
            v3 = (1e10 * 3.14159)
            v23 := []int{}
            _ = fi22
            _ = v23
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
