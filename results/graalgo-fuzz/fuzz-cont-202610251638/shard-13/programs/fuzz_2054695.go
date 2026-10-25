package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int32 }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint64, p1 uint, p2 int32) uint16 {
    var v4 int32 = (p2 ^ p2)
    v5 := GenS1{F0: 65535, F1: -645, F2: 42}
    if false {
        var v6 uint16 = 323
        p2 += -303
        _ = v6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        p2 -= p2
        fmt.Println(p1, v4, p2)
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 32767
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]uint32)
    var v8 uint = 65535
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(v8, v1)
        _ = fi9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        v1 = 1
        var v11 int64 = -1
        switch "" {
        case "\n\t":
            v12 := GenS2{F0: 1, F1: 18446744073709551615}
            _ = v12
        case "foo bar":
            var v13 int8 = -23
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    v1 = (v1 * (v1 - -2147483648))
    v8 += (v8 + (^(v8)))
    var v14 bool = true
    fmt.Println(v8, v1, v14)
    var v15 bool = (v8 == v8)
    v16 := []int{-648, 9223372036854775807, 2147483647, -804}
    v15 = ((0.1 - (((1.5 - 1e10) + (0.0 + 1e10)) - -2.25)) >= -2.25)
    v17 := v2[232]
    v18 := v3["hello"]
    v18 = 824
    var v19 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
}
