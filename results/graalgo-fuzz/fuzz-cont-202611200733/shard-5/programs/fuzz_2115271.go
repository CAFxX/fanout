package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 int16 = 576
    fmt.Println(v3)
    v4 := GenS1{F0: 474}
    if (2147483648 != 298) {
        v5 := []int{-531, 426}
        if len(v5) > 0 {
            v5[0] = 30
        }
        if (0.1 <= 1.5) {
            var v6 int8 = -105
            var v7 float64 = 0.1
            _ = v6
            _ = v7
        }
        _ = v5
    } else {
        v3 = (^(v3))
    }
    var v8 uint32 = 0
    for fi9 := 0; fi9 < 15; fi9++ {
        fmt.Println(v8)
        v3 -= v3
        _ = fi9
    }
    v3 += 42
    var v10 uint64 = 895
    switch "foo bar" {
    case "\n\t":
        v10 += (^(v10))
    case "hello":
        v3 = v3
    }
    v8 = (v8 / (((v8 - v8)) | 1))
    for fi11 := 0; fi11 < 10; fi11++ {
        v1[fi11] = (v8 * (v8 / ((2147483648) | 1)))
        _ = fi11
    }
    var v12 int32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
