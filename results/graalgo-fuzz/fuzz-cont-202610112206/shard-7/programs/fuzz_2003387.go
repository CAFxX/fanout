package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint16) bool {
    v4 := []int{}
    var v5 int32 = -1
    _ = p0
    _ = v4
    _ = v5
    return ((1e10 * ((1e10 * (0.0 / 3.14159)) - 0.0)) != 0.0)
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]string)
    var v3 uint64 = 42
    v3 = 42
    v6 := []int{9223372036854775807, -999, 2147483647, 4294967295}
    v7 := GenS1{F0: 24, F1: (v3 / (((v3 - (v3 >> ((1) & 63)))) | 1))}
    var v8 string = (("a" + "hello") + "a")
    v8 = v8
    v3 = v3
    v9 := GenS1{F0: 238, F1: (v3 >> ((2) & 63))}
    v3 = 18446744073709551615
    _ = len(v6)
    var v10 float64 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
