package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint64; F2 bool }
func genFn0(p0 uint, p1 uint, p2 int8) uint64 {
    p1 += 18446744073709551615
    var v4 int64 = 0
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 int32 = -2147483648
        var v7 int = fi5
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}

func main() {
    v1 := []int{128, -764, 42}
    var v2 bool = true
    v3 := make(map[int]uint32)
    v2 = true
    fmt.Println(v2)
    var v8 float32 = 3.14159
    switch "a" {
    case "foo bar":
        v8 = (-2.25 * v8)
    }
    v9 := GenS1{F0: 759}
    var v10 int32 = 2147483647
    v10 = v10
    var v11 uint = 65535
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 uint64 = 42
        _ = fi12
        _ = v13
    }
    v8 = 1e10
    var v14 uint8 = 1
    v15 := GenS1{F0: -307}
    v16 := GenS2{F0: 0, F1: 2, F2: (false && (!(!((-2.25 + (3.14159 - (3.14159 / (0.0 + float64(v8))))) <= (0.1 / ((0.0 / 3.14159) + (1.5 / (1.5 / 0.1))))))))}
    switch "" {
    case "\n\t":
        var v17 uint64 = 384
        _ = v17
    case "hello":
        var v18 int8 = 0
        _ = v18
    default:
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 bool = true
            fmt.Println(v11)
            _ = fi19
            _ = v20
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
