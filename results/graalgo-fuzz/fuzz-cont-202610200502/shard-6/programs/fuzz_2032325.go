package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint8) int16 {
    var v6 int16 = -32768
    var v7 uint16 = 356
    v8 := []int{-2147483648, -1, 4294967295, 0}
    p0 += p0
    fmt.Println(p0, v6)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return (v6 | v6)
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]bool)
    var v3 float64 = (v1 * (((1e10 * float64(1e10)) - (v1 * 1e10)) / v1))
    v4 := make(map[int]uint32)
    v5 := make(map[string]uint32)
    var v9 int64 = 42
    v3 += float64((0.1 - (0.1 - 1.5)))
    fmt.Println(v1, v9)
    v9 = 1
    v9 = 9223372036854775807
    fmt.Println(v3, v9)
    var v10 uint64 = 1
    var v11 int = 0
    v12 := GenS1{F0: v10}
    switch (v11 - v11) {
    case 2147483647:
        v9 -= -762
    case 42:
        switch ("hello" + "foo bar") {
        case "\n\t":
            v3 += (v1 / v3)
        case "hello":
            switch 2147483647 {
            case 32767:
                v13 := []int{(v11 / ((v11) | 1)), (v11 * (v11 - v11)), (v11 + v11), 4294967295}
                _ = v13
            case 1:
                v14 := GenS1{F0: v10}
                _ = v14
            default:
                var v15 uint64 = v10
                _ = v15
            }
        default:
            var v16 int16 = -1
            _ = v16
        }
    case -1:
        v17 := []int{(^(v11))}
        _ = v17
    default:
        if false {
            fmt.Println(v1)
        } else {
            if (-578 >= -1) {
                v11 -= v11
                var v18 float64 = v3
                _ = v18
            } else {
                v10 += (v10 ^ v10)
            }
            v10 -= (v10 + (v10 | v10))
        }
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
