package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int8; F1 int64; F2 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 169
    v3 := []int{1}
    v4 := []string{"a", "a", ""}
    v5 := make(map[uint32]int)
    v6 := []string{""}
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint16 = 2
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 1, F1: (("hello" + "hello") + "a")}
    var v10 int32 = -2147483648
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(v10)
        v10 = 2147483647
        _ = fi11
    }
    delete(v5, 457)
    switch "a" {
    case "\n\t":
        if ((v10 % (((^(v10))) | 1)) <= -18) {
            fmt.Println(v2)
        } else {
            var v12 uint = 4294967295
            _ = v12
        }
    }
    v2 += 0
    switch 128 {
    case 9223372036854775807:
        v13 := GenS2{F0: v2}
        _ = v13
    case 32768:
        var v14 int32 = v10
        _ = v14
    default:
        switch "hello" {
        case "hello":
            var v15 int16 = 42
            _ = v15
        default:
            var v16 float32 = 0.1
            _ = v16
        }
    }
    var v17 int8 = 42
    var v18 uint = 422
    v19 := v5[845]
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        fmt.Println(fi20)
        _ = fi20
    }
    v2 += v2
    v21 := GenS3{F0: v10, F1: 267}
    var v22 int = (-(v19))
    v19 -= (v22 ^ (v22 * 32767))
    v23 := GenS4{F0: -32, F1: (v2 / ((9223372036854775807) | 1)), F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
