package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint32; F1 string }

func main() {
    var v1 uint64 = 2
    var v2 int8 = 100
    var v3 float32 = (1.5 * ((1.5 / ((0.1 * (1e10 * 0.1)) - ((0.1 * (((0.1 * 1e10) + 1e10) - 1e10)) - (0.1 + 0.0)))) + 0.1))
    v3 = v3
    v4 := []int{}
    _ = len(v4)
    v5 := GenS1{F0: 1, F1: (v2 << ((65535) & 63)), F2: 42}
    fmt.Println(v1, v3, v3)
    var v6 uint32 = 2147483648
    var v7 int16 = 1
    _ = len(v4)
    fmt.Println(v1, v7)
    v7 += v7
    for fi8 := 0; fi8 < 13; fi8++ {
        if (127 > (v2 >> ((65535) & 63))) {
            v9 := GenS2{F0: (v6 - v6), F1: (("foo bar" + "hello") + "\n\t")}
            _ = v9
        } else {
            v10 := GenS1{F0: -2147483648, F1: (v2 / ((v2) | 1)), F2: 171}
            var v11 int32 = int32(v6)
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    v1 -= (^(v1))
    if (42 >= -757) {
        var v12 uint8 = 0
        _ = len(v4)
        if false {
            v7 -= v7
            fmt.Println(v12, v2, v12)
            v13 := GenS2{F0: (^(v6)), F1: (("\n\t" + ("hello" + "hello")) + "")}
            _ = v13
        }
        _ = v12
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v14 := GenS2{F0: v6, F1: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
