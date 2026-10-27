package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint }

func main() {
    var v1 int16 = -1
    var v2 int64 = 1
    var v3 int64 = v2
    v4 := make(map[int]bool)
    var v5 uint = 818
    var v6 uint8 = 248
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v6, v5)
    default:
        var v7 uint16 = 371
        _ = v7
    }
    v3 -= -1
    fmt.Println(v2, v3)
    var v8 int64 = (v3 * 9223372036854775807)
    v9 := GenS1{F0: v8, F1: v1, F2: (v5 - 42)}
    var v10 bool = false
    var v11 uint64 = 4294967295
    if v10 {
        v8 -= -250
    }
    v10 = v10
    var v12 float64 = float64(0.0)
    var v13 int64 = (v2 >> ((4294967295) & 63))
    var v14 float64 = (((((v12 / v12) / v12) * (3.14159 * 1e10)) * (v12 * float64(1e10))) * (v12 + float64(1.5)))
    delete(v4, int(v1))
    if ("a" < "\n\t") {
        v15 := GenS1{F0: (v3 | 42), F1: v1, F2: v5}
        v16 := []int{-764, -1}
        var v17 int16 = int16(v5)
        _ = v15
        _ = v16
        _ = v17
    } else {
        var v18 uint32 = 2147483648
        _ = v18
    }
    v19 := GenS1{F0: (v2 * -427), F1: 42, F2: (v5 - v5)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
