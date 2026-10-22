package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }

func main() {
    var v1 float64 = ((1.5 + 1.5) - 0.0)
    v2 := make(map[int]uint32)
    var v3 uint8 = 101
    var v4 int64 = -685
    fmt.Println(v3, v1, v4)
    switch "foo bar" {
    case "foo bar":
        if false {
            v3 = (v3 & v3)
            var v5 string = "a"
            var v6 float64 = v1
            _ = v5
            _ = v6
        }
    case "\n\t":
        switch -9223372036854775808 {
        case 850:
            for fi7 := 0; fi7 < 1; fi7++ {
                var v8 int32 = 42
                var v9 float32 = 0.1
                _ = fi7
                _ = v8
                _ = v9
            }
        default:
            v3 = v3
        }
    }
    var v10 int32 = -472
    v11 := []int{2147483647, 32768}
    var v12 int64 = v4
    var v13 int8 = 1
    v14 := GenS1{F0: v10, F1: (-32768 == 0)}
    v4 = v12
    var v15 uint8 = (v3 + v3)
    v10 -= v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
