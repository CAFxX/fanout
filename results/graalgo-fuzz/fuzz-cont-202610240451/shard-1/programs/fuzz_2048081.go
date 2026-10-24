package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int16 }
func genFn0(p0 string) uint32 {
    var v5 uint64 = 1
    switch "" {
    case "a":
        switch 128 {
        case 9223372036854775807:
            var v6 int = 4294967295
            _ = v6
        case 586:
            fmt.Println(v5, v5)
        case 934:
            var v7 float32 = float32(((0.1 - -2.25) / (1.5 + (0.0 - 0.0))))
            _ = v7
        default:
            v8 := GenS1{F0: -507, F1: 0, F2: 1000}
            _ = v8
        }
    case "hello":
        var v9 int32 = 42
        _ = v9
    case "foo bar":
        v10 := []int{2147483647, -9223372036854775808}
        _ = v10
    }
    var v11 int32 = 377
    p0 = p0
    _ = p0
    _ = v11
    _ = v5
    return 2147483648
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = -36
    var v3 uint8 = 42
    v4 := make(map[uint32]uint32)
    var v12 uint64 = 1
    var v13 uint32 = 2147483648
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 uint8 = v3
        _ = fi14
        _ = v15
    }
    v12 += v12
    _ = len(v4)
    v16 := GenS1{F0: 1, F1: 805, F2: 992}
    var v17 float64 = 1.5
    v13 = genFn0("a")
    var v18 int16 = 0
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
