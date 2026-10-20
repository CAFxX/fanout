package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int16, p1 uint64, p2 uint32) int8 {
    var v5 int8 = 42
    v6 := []int{128, -578}
    p1 += p1
    v5 = (v5 * (v5 * (v5 & 41)))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 int8) uint8 {
    v7 := GenS1{F0: -1, F1: p0}
    v8 := GenS1{F0: 1, F1: 1}
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 42
    var v3 int16 = -791
    v4 := []int{1, 4294967295, -1}
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    var v9 int64 = 9223372036854775807
    v10 := GenS2{F0: (^(v3))}
    switch "hello" {
    case "foo bar":
        var v11 int = -836
        _ = v11
    }
    v9 = 0
    var v12 uint32 = 113
    if (0.1 != (((0.0 * 1e10) * 1.5) + (((0.0 + 0.1) / 1.0) + float64(float32(-2.25))))) {
        var v13 int16 = v3
        _ = v13
    } else {
        switch -659 {
        case -1:
            fmt.Println(v9, v9)
        case 127:
            v14 := GenS2{F0: -1}
            _ = v14
        }
        v12 -= (^(v12))
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 uint64 = 18446744073709551615
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := v1[("hello" + (("" + ("a" + "hello")) + ("\n\t" + "hello")))]
                if len(v4) > 0 {
                    v4[0] = fi15
                }
                _ = fi17
                _ = v18
            }
            _ = fi15
            _ = v16
        }
    }
    var v19 uint32 = 2
    var v20 uint64 = 455
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
