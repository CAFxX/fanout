package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 1e10
    v3 := []uint32{0}
    var v4 int64 = 0
    if (-1 < -32768) {
        v5 := GenS1{F0: 4294967295, F1: 18446744073709551615}
        v2 = v2
        var v6 uint = 864
        _ = v5
        _ = v6
    } else {
        fmt.Println(v4)
        var v7 uint64 = 42
        _ = v7
    }
    v1[2147483647] = 244
    if ((3.14159 / 3.14159) <= 1.5) {
        delete(v1, -174)
        v2 = (v2 - -2.25)
    } else {
        if (4294967295 != 183) {
            v8 := v1[127]
            _ = v8
        } else {
            fmt.Println(v2, v4)
            v4 = 9223372036854775807
            switch "\n\t" {
            case "foo bar":
                if len(v3) > 0 {
                    v3[0] = 250
                }
            case "":
                var v9 float32 = (0.1 / ((1e10 / 1.0) * ((1.5 - (1.5 / (1.5 - float32(v2)))) * 0.0)))
                _ = v9
            }
        }
    }
    var v10 uint64 = 289
    var v11 float64 = 0.1
    for fi12 := 0; fi12 < 9; fi12++ {
        v4 -= (v4 / ((v4) | 1))
        delete(v1, fi12)
        _ = fi12
    }
    var v13 uint16 = 359
    v2 = v11
    var v14 bool = true
    switch 128 {
    case 431:
        v15 := GenS2{F0: -586}
        _ = v15
    case 32768:
        var v16 int32 = -577
        _ = v16
    }
    v4 = v4
    var v17 uint16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
