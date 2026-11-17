package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }

func main() {
    v1 := []int{32767, 32767, -488, 9223372036854775807}
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    if len(v1) > 0 {
        v1[0] = -13
    }
    if (0 >= 994) {
        v4 := GenS1{F0: ((0 < 0) || (42 > 18446744073709551615)), F1: 127}
        _ = v4
    }
    var v5 uint8 = 0
    v5 += 125
    v1 = append(v1, 42)
    var v6 int32 = 854
    v7 := GenS1{F0: true, F1: 42}
    fmt.Println(v6, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
