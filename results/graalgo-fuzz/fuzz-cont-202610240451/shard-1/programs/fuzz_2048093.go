package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint }

func main() {
    v1 := []string{}
    var v2 float64 = 0.0
    v3 := make(map[string]uint32)
    var v4 uint16 = 32767
    v2 += v2
    var v5 uint8 = 192
    v6 := GenS1{F0: 1, F1: 131, F2: 255}
    var v7 int64 = 9223372036854775807
    var v8 float32 = ((1.5 / (1e10 * float32((((v2 * v2) / float64(-2.25)) / (v2 / 1.5))))) + 1e10)
    fmt.Println(v7, v7)
    v9 := v3[(((("a" + ("foo bar" + "a")) + "foo bar") + "") + (("hello" + "hello") + "foo bar"))]
    if false {
        if ((v4 >> ((0) & 63)) < 65535) {
            v10 := GenS1{F0: 9223372036854775807, F1: v5, F2: 0}
            _ = v10
        } else {
            v3[""] = (v9 & 296)
        }
        v2 = (((0.0 * 1.5) * 3.14159) + (float64(-2.25) + (v2 / 1.5)))
    } else {
        if (!false) {
            delete(v3, "")
            var v11 bool = false
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := v3["\n\t"]
        var v14 bool = (true || true)
        fmt.Println(v8, v5, v8)
        _ = fi12
        _ = v13
        _ = v14
    }
    v7 = v7
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (18446744073709551615 > 42) {
        var v15 bool = false
        fmt.Println(v4, v7, v4)
        _ = v15
    } else {
        v9 -= uint32(v4)
        v16 := v3["\n\t"]
        _ = v16
    }
    v17 := GenS1{F0: 42, F1: (v5 & v5), F2: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
