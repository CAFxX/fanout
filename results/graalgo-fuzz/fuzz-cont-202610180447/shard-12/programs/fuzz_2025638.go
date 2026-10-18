package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 float32 = ((((3.14159 / 1.0) * ((1.5 - 0.1) + 0.1)) + 1.5) + (1.5 + float32(3.14159)))
    var v3 uint = 358
    var v4 uint16 = 0
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 uint32 = 601
        fmt.Println(v4, v6)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: -1}
    if (-1 >= 111) {
        var v8 float64 = 1.5
        var v9 uint32 = 42
        _ = v8
        _ = v9
    } else {
        var v10 float64 = 3.14159
        v11 := GenS2{F0: 2147483647, F1: 42}
        _ = v10
        _ = v11
    }
    v12 := []int{2147483647, -356}
    for fi13 := 0; fi13 < 1; fi13++ {
        v4 = (v4 | 1)
        switch fi13 {
        case 0:
            var v14 uint = v3
            _ = v14
        case 42:
            v4 += v4
        default:
            if ((!true) || (0 < 2)) {
                for fi15 := 0; fi15 < 2; fi15++ {
                    _ = len(v1)
                    _ = fi15
                }
                var v16 int8 = -128
                var v17 bool = ((float32(-2.25) + float32(1e10)) != (float32(-2.25) / (float32(((3.14159 * float64(v2)) / 1.0)) - v2)))
                _ = v16
                _ = v17
            }
        }
        var v18 uint = v3
        _ = fi13
        _ = v18
    }
    switch "foo bar" {
    case "hello":
        v3 += v3
    case "\n\t":
        v4 = 255
    }
    fmt.Println(v2, v2, v2)
    var v19 int = 2147483647
    for fi20 := 0; fi20 < 6; fi20++ {
        v2 -= v2
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
