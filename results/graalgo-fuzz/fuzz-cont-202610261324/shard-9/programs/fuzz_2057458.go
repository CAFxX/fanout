package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int16 = -1
    var v2 float64 = (0.0 - 1e10)
    v3 := []int{587}
    var v4 int16 = (v1 / (((v1 / ((1000) | 1))) | 1))
    fmt.Println(v1)
    v4 += (v4 << ((18446744073709551615) & 63))
    var v5 float32 = ((1e10 / 3.14159) * 3.14159)
    v5 = (-2.25 * (float32((v2 * ((1e10 - (float64((0.0 + (v5 * 0.1))) + 3.14159)) - v2))) * v5))
    v6 := GenS1{F0: 0}
    v5 = 0.0
    v1 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
