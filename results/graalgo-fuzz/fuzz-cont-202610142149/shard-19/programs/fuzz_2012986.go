package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int32) uint32 {
    var v6 string = "hello"
    switch 32767 {
    case 433:
        fmt.Println(p0)
    case 9223372036854775807:
        p0 += p0
    case -1:
        v7 := GenS1{F0: 2, F1: 822, F2: 448}
        _ = v7
    }
    fmt.Println(p0, v6)
    var v8 bool = (false && false)
    _ = p0
    _ = v6
    _ = v8
    return 1
}
func genFn1(p0 uint32, p1 int8, p2 bool) uint16 {
    var v9 uint16 = 65535
    p2 = false
    v10 := GenS1{F0: 42, F1: 4294967295, F2: v9}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return (v9 % ((42) | 1))
}

func main() {
    var v1 uint64 = 167
    v2 := make(map[int]int)
    var v3 bool = (0 < -293)
    var v4 bool = v3
    v5 := []uint32{0, 342}
    v11 := GenS1{F0: 42, F1: 42, F2: 2}
    fmt.Println(v4, v3)
    v4 = true
    if v4 {
        var v12 bool = (-1 < -128)
        var v13 float64 = 1e10
        v13 = v13
        _ = v12
        _ = v13
    } else {
        switch "" {
        case "\n\t":
            v3 = v3
        case "a":
            var v14 float64 = 0.1
            _ = v14
        default:
            fmt.Println(v4, v3)
        }
    }
    fmt.Println(v1, v4)
    if (-2.25 != 0.1) {
        fmt.Println(v3, v4, v1)
        _ = len(v5)
    } else {
        v15 := GenS1{F0: (v1 % (((^(v1))) | 1)), F1: 0, F2: 2}
        var v16 uint = 1
        v17 := GenS2{F0: 42}
        _ = v15
        _ = v16
        _ = v17
    }
    var v18 string = "\n\t"
    var v19 uint64 = (v1 | (v1 * (v1 + v1)))
    delete(v2, 127)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
