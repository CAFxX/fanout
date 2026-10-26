package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    var v1 uint64 = 1
    var v2 uint64 = v1
    var v3 string = ""
    v4 := []string{v3, (("a" + v3) + ""), v3, "a"}
    var v5 uint16 = 1
    v6 := GenS1{F0: -163, F1: 127}
    v3 = ""
    v7 := GenS1{F0: 9223372036854775807, F1: 606}
    var v8 float64 = ((((0.1 / (-2.25 / ((1.5 / (((((0.0 + (((1e10 * 0.1) + 0.1) - (0.0 / 0.1))) + (0.0 * (1.5 - 0.0))) - -2.25) + 1e10) - 1e10)) + 0.1))) / 1.5) / ((0.1 + float64(float32(3.14159))) + ((float64(float32(0.1)) * (0.0 + (1e10 * (float64(((0.1 - (1e10 - 0.0)) / 1e10)) * 0.0)))) - float64((-2.25 + (3.14159 - (((1e10 / ((1.5 / 0.1) / 3.14159)) / 1.0) / (1.5 / 1.0)))))))) * float64(3.14159))
    var v9 int32 = -2147483648
    var v10 uint64 = v2
    _ = len(v4)
    var v11 float64 = v8
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
