package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 1
    v3 := make(map[int]string)
    var v4 int32 = 1
    if false {
        if (42 >= 248) {
            var v5 float64 = (0.1 + ((float64((((0.0 * (0.0 + (0.1 * (-2.25 + 1.5)))) + 0.1) + (float32(float64(0.1)) - (3.14159 / (float32((0.0 - 1.5)) / 1.5))))) / (0.1 / -2.25)) + 1e10))
            _ = v5
        } else {
            v4 += v4
            var v6 uint = 4294967295
            v4 = (v4 % (((v4 << (((v6 ^ 717)) & 63))) | 1))
            _ = v6
        }
        var v7 int8 = 127
        _ = v7
    }
    v8 := GenS1{F0: 42}
    v2 -= v2
    var v9 bool = true
    fmt.Println(v4)
    v10 := GenS2{F0: 255}
    v9 = (2 > 0)
    fmt.Println(v4, v9, v9)
    switch "" {
    case "":
        if false {
            fmt.Println(v4, v2, v2)
        } else {
            delete(v1, "")
        }
    case "\n\t":
        var v11 int = -9223372036854775808
        _ = v11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
