package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint) string {
    v5 := []int{128, -917, 197, 1}
    var v6 uint = p0
    _ = p0
    _ = v5
    _ = v6
    return "a"
}

func main() {
    var v1 float64 = ((float64(((((((0.0 * float32(1e10)) / 3.14159) * (0.0 - float32(0.0))) - (((3.14159 - -2.25) - 1.5) + 0.1)) - 0.0) + (3.14159 / 1.0))) + 1e10) / 3.14159)
    var v2 bool = (0 < 127)
    var v3 int = 9223372036854775807
    v4 := make(map[int]bool)
    v7 := GenS1{F0: 221, F1: 54}
    if (!v2) {
        v1 += v1
    } else {
        var v8 int32 = -211
        _ = v8
    }
    v2 = true
    switch 9223372036854775807 {
    case -1:
        var v9 int16 = 0
        _ = v9
    case 4294967295:
        v10 := GenS1{F0: -253, F1: 1}
        _ = v10
    }
    if true {
        fmt.Println(v3, v2, v3)
        fmt.Println(v3, v1, v1)
        var v11 uint16 = 42
        _ = v11
    }
    var v12 int = v3
    v13 := []int{v12, (v3 % ((v3) | 1)), v3}
    v14 := GenS2{F0: 2}
    if ((2 <= 127) && (v12 < v3)) {
        var v15 int64 = 9223372036854775807
        _ = v15
    }
    v3 = 4294967295
    var v16 float32 = 1.5
    var v17 uint = 1
    v16 = v16
    switch "a" {
    case "a":
        fmt.Println(v16, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
