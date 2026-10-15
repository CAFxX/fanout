package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 bool; F1 int16; F2 uint32 }
func genFn0(p0 int16, p1 int, p2 uint16) uint16 {
    v7 := GenS1{F0: p1, F1: 1}
    fmt.Println(p0)
    if true {
        p1 -= (^(p1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}
func genFn1(p0 uint, p1 uint64, p2 int16) uint {
    var v8 uint32 = 273
    var v9 int64 = -9223372036854775808
    var v10 int = 2147483647
    var v11 float32 = -2.25
    var v12 uint8 = uint8(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v8
    _ = v9
    return 412
}

func main() {
    var v1 bool = true
    var v2 float32 = (((1.5 / 0.1) / 1.5) - -2.25)
    v3 := []int{1}
    v4 := make(map[uint32]string)
    v5 := make(map[int]string)
    var v6 bool = (536 >= 65535)
    if (!true) {
        var v13 uint = 57
        _ = v13
    } else {
        v14 := v5[-9223372036854775808]
        fmt.Println(v14, v2)
        _ = v14
    }
    var v15 int = -2147483648
    var v16 float64 = (0.1 * 0.0)
    v15 = v15
    v17 := GenS1{F0: (v15 - v15), F1: -2147483648}
    if false {
        var v18 int8 = 100
        switch "foo bar" {
        case "":
            var v19 uint64 = 18446744073709551615
            _ = v19
        case "a":
            v20 := GenS1{F0: v15, F1: 1}
            _ = v20
        default:
            v21 := GenS2{F0: v1, F1: 0, F2: 1}
            _ = v21
        }
        if (v6 && v1) {
            fmt.Println(v2, v18, v6)
            delete(v5, (v15 >> ((531) & 63)))
        }
        _ = v18
    }
    var v22 float32 = ((v2 + float32((v16 / 3.14159))) - -2.25)
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
