package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (1e10 - (-2.25 + -2.25))
    var v3 uint16 = 1
    v4 := make(map[string]uint32)
    var v5 int8 = 100
    var v6 int16 = int16(v5)
    v3 += (^(v3))
    fmt.Println(v2, v2)
    var v7 int16 = (v6 + (v6 + -1))
    v8 := []int{940, 0}
    v3 = (v3 | v3)
    v6 = (v6 & 32767)
    v3 -= (v3 % (((v3 % ((922) | 1))) | 1))
    switch 32768 {
    case 661:
        v9 := GenS1{F0: 18446744073709551615, F1: v3, F2: -616}
        _ = v9
    case 275:
        v7 = int16(v5)
    }
    v7 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
