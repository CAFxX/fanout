package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint32 }

func main() {
    v1 := []uint32{2147483648, 42, 2147483647}
    var v2 uint = 65535
    v3 := make(map[int]int)
    v4 := []uint32{987, 2147483647, 1, 453}
    var v5 int8 = -128
    v6 := make(map[string]string)
    if ((false || (0.0 == 0.0)) && false) {
        if (-1 == int16(v5)) {
            fmt.Println(v5, v5)
            for fi7 := 0; fi7 < 12; fi7++ {
                v2 -= v2
                var v8 uint32 = 528
                _ = fi7
                _ = v8
            }
        } else {
            var v9 float64 = ((0.0 * 3.14159) - (1.5 / 1.5))
            _ = v9
        }
    }
    fmt.Println(v2, v5)
    v5 = 0
    v5 += (v5 | v5)
    if len(v1) > 0 {
        v1[0] = 1
    }
    switch "hello" {
    case "":
        v10 := GenS1{F0: "a", F1: 42, F2: 212}
        _ = v10
    }
    _ = len(v3)
    if true {
        var v11 float32 = -2.25
        for fi12 := 0; fi12 < 14; fi12++ {
            fmt.Println(v11, fi12)
            _ = fi12
        }
        _ = v11
    }
    v2 += v2
    var v13 uint16 = 65535
    fmt.Println(v13, v5)
    v5 += 125
    if true {
        v14 := GenS1{F0: "hello", F1: 0, F2: 46}
        _ = v14
    } else {
        fmt.Println(v5, v5, v13)
        v13 -= v13
        fmt.Println(v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
