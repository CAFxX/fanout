package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int16 }

func main() {
    v1 := []uint32{}
    v2 := make(map[int]bool)
    var v3 int32 = 768
    v4 := make(map[string]uint32)
    v5 := []string{(("foo bar" + (("" + "\n\t") + "foo bar")) + "foo bar")}
    for fi6 := 0; fi6 < 10; fi6++ {
        v3 = -540
        var v7 float32 = 1e10
        _ = fi6
        _ = v7
    }
    if true {
        var v8 bool = false
        var v9 int64 = 42
        _ = v8
        _ = v9
    } else {
        if (489 >= -871) {
            var v10 int64 = -739
            _ = v10
        }
    }
    var v11 uint64 = 18446744073709551615
    var v12 int16 = -169
    var v13 uint16 = 2
    switch "" {
    case "\n\t":
        for fi14 := 0; fi14 < 3; fi14++ {
            if false {
                _ = len(v2)
            } else {
                v15 := GenS1{F0: v3}
                _ = v15
            }
            _ = fi14
        }
    case "a":
        var v16 float64 = 3.14159
        _ = v16
    case "foo bar":
        v12 = (v12 & (v12 ^ int16(v11)))
    default:
        var v17 uint64 = (v11 * (^(v11)))
        _ = v17
    }
    if false {
        v18 := v2[int(v12)]
        var v19 uint = 604
        _ = v18
        _ = v19
    } else {
        v20 := GenS1{F0: 875}
        _ = v20
    }
    v11 = 643
    fmt.Println(v12)
    var v21 string = "\n\t"
    fmt.Println(v3)
    if false {
        var v22 int = 4294967295
        fmt.Println(v22, v3, v22)
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        _ = v22
    }
    v3 = (^(v3))
    if ((0.1 * 3.14159) <= -2.25) {
        v23 := GenS2{F0: 42, F1: (v12 << ((42) & 63))}
        _ = v23
    } else {
        v5 = append(v5, v21)
        var v24 int32 = 1
        _ = v24
    }
    v25 := GenS2{F0: 2, F1: v12}
    v13 -= 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
