package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint16) bool {
    switch "foo bar" {
    case "\n\t":
        v7 := []int{656, 638, 32767, 1}
        _ = v7
    default:
        p0 = 1
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: 1, F1: 31}
        var v10 uint32 = 247
        v11 := GenS2{F0: 1}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    _ = p0
    return (65535 == 0)
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 0
    v3 := make(map[uint32]string)
    var v4 int64 = 660
    var v5 int8 = 100
    var v6 float32 = 3.14159
    v6 = float32(float64(1.5))
    v12 := []int{}
    for fi13 := 0; fi13 < 12; fi13++ {
        fmt.Println(v6, v4)
        fmt.Println(v6, v2, v6)
        var v14 int32 = 2147483647
        _ = fi13
        _ = v14
    }
    fmt.Println(v2, v5)
    fmt.Println(v5, v2, v5)
    v5 -= -40
    fmt.Println(v5, v4)
    fmt.Println(v5, v6)
    v5 = v5
    if true {
        _ = len(v12)
        var v15 int16 = 73
        _ = v15
    }
    v4 -= (v4 * v4)
    v2 = (-(v2))
    switch "" {
    case "a":
        var v16 int = 32767
        _ = v16
    case "\n\t":
        fmt.Println(v4, v6)
    case "hello":
        var v17 uint32 = 65535
        _ = v17
    default:
        var v18 int32 = 0
        _ = v18
    }
    v19 := GenS1{F0: 2147483647, F1: -88}
    if (false || true) {
        if false {
            v6 = (v6 + (v6 * ((0.0 + (float32((0.1 - (float64(float32(float64(v6))) + 1e10))) + 3.14159)) / 1e10)))
        }
    } else {
        v6 += (-2.25 * -2.25)
    }
    var v20 uint = 531
    var v21 uint16 = 32767
    fmt.Println(v2, v20, v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
