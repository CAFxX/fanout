package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint32; F2 int32 }

func main() {
    var v1 float64 = float64(((0.1 * 0.1) / 3.14159))
    v2 := []int{}
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    var v5 int16 = -32768
    var v6 uint32 = 785
    v7 := GenS1{F0: 20}
    fmt.Println(v1, v5)
    v1 = v1
    v6 -= v6
    v1 = (v1 / 1.5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v8 := GenS2{F0: v6, F1: v6, F2: -1}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 += (((1e10 + 1e10) - v1) / (v1 - 0.0))
    v3[("" + "foo bar")] = -648
    if true {
        var v9 string = "\n\t"
        _ = v9
    }
    v10 := GenS2{F0: (v6 << ((65535) & 63)), F1: (v6 + v6), F2: 0}
    switch 127 {
    case 32767:
        if true {
            fmt.Println(v5, v6, v6)
        }
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
