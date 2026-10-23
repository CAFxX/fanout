package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 4294967295
    v3 := []string{}
    var v4 float32 = 1.5
    var v5 float64 = 1.5
    fmt.Println(v2)
    if true {
        var v6 float32 = (1e10 / 1e10)
        var v7 float32 = 1.5
        if false {
            var v8 uint16 = 1
            v2 = (v2 | (v2 / ((v2) | 1)))
            _ = v8
        }
        _ = v6
        _ = v7
    } else {
        v4 = (((((v4 + ((v4 / 1e10) * 0.1)) - (v4 + 0.0)) / v4) / v4) / float32(float64(float32(3.14159))))
    }
    switch "foo bar" {
    case "foo bar":
        _ = len(v3)
    case "":
        v2 = v2
    default:
        var v9 bool = (((false && (((1.5 - (v5 - ((v5 * ((float64((v4 / (0.1 - (float32((v5 - v5)) + (((float32(v5) + (v4 - -2.25)) + (v4 / float32((0.1 + (0.0 + (1e10 * (-2.25 * float64((v4 - -2.25))))))))) - (0.0 * v4)))))) + 0.0) / float64(((-2.25 + v4) * float32(v5))))) / (v5 * float64(1e10))))) > 1e10) && (-93 <= 127))) && (42 < -976)) && (float64(v4) > (v5 - (v5 - 3.14159))))
        _ = v9
    }
    var v10 int8 = -104
    v10 = (v10 % ((v10) | 1))
    v10 -= v10
    v4 -= 0.1
    switch "" {
    case "":
        v10 = v10
    }
    v1[-612] = 2
    fmt.Println(v2, v10, v5)
    var v11 float64 = (float64(float32(v5)) + 0.1)
    v12 := GenS1{F0: 315, F1: -645, F2: 227}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
