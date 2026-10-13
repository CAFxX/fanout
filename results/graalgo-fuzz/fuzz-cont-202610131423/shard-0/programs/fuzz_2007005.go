package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 string }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = ((3.14159 / ((((((0.0 - 1e10) * (float64(0.0) / (1e10 / 1.0))) / 1e10) * 0.1) - ((1e10 * 3.14159) * 0.0)) + 0.1)) + 3.14159)
    var v3 int8 = 33
    var v4 int = int(v3)
    var v5 int16 = 0
    v6 := []uint32{0}
    var v7 bool = false
    v2 -= (v2 / float64(0.1))
    v8 := GenS1{F0: 9223372036854775807, F1: 0, F2: "a"}
    switch v4 {
    case -9223372036854775808:
        if (!(-903 == 980)) {
            v1[(v4 % ((-1) | 1))] = (v4 * int(v3))
            var v9 int8 = (v3 ^ v3)
            _ = v9
        }
    }
    if v7 {
        fmt.Println(v5)
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int8 = 64
            fmt.Println(v5)
            var v12 float64 = 0.1
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    if v7 {
        fmt.Println(v5, v5)
    }
    if v7 {
        v2 = (1e10 + float64(0.1))
        var v13 uint64 = 126
        if (0.0 != -2.25) {
            var v14 uint32 = 4294967295
            var v15 bool = v7
            _ = v14
            _ = v15
        } else {
            delete(v1, v4)
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 string = "hello"
                var v18 uint16 = 32767
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        _ = v13
    } else {
        switch "foo bar" {
        case "foo bar":
            if (float32(0.1) == 1.5) {
                fmt.Println(v4, v2, v5)
                var v19 int = 518
                fmt.Println(v4)
                _ = v19
            }
        case "\n\t":
            var v20 int8 = v3
            _ = v20
        default:
            v5 -= -1
        }
        v2 = ((v2 + float64(0.0)) + (1e10 * v2))
    }
    v1[v4] = (v4 / ((v4) | 1))
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 int16 = (v5 & (-(v5)))
        _ = fi21
        _ = v22
    }
    var v23 int = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
