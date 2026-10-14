package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -1
    var v3 string = ""
    v4 := make(map[int]bool)
    if (float64((0.1 + (1e10 - float32(((1.5 / (((float64(0.1) * 3.14159) * (3.14159 / (1.5 / 1.0))) / 1.0)) - 0.0))))) <= (0.1 - -2.25)) {
        var v5 float32 = 0.0
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 float32 = 3.14159
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    var v8 float64 = (0.0 / -2.25)
    var v9 bool = (v3 == (v3 + v3))
    var v10 int8 = (v2 | (v2 % ((v2) | 1)))
    v11 := v4[650]
    var v12 bool = (255 <= 2)
    v13 := GenS1{F0: 251}
    if v9 {
        v8 = (v8 * (v8 * (v8 * (v8 - 0.1))))
    }
    switch v3 {
    case "foo bar":
        var v14 float64 = v8
        _ = v14
    default:
        if v9 {
            v15 := GenS2{F0: 18446744073709551615, F1: (((("\n\t" + v3) + v3) + "hello") + v3)}
            fmt.Println(v3, v8)
            v16 := GenS1{F0: 363}
            _ = v15
            _ = v16
        }
    }
    switch "\n\t" {
    case "\n\t":
        v10 = 127
    }
    _ = len(v1)
    var v17 uint16 = 438
    _ = len(v1)
    v18 := []int{127}
    var v19 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
