package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint8; F2 bool }
func genFn0(p0 int, p1 uint64) int32 {
    var v7 uint16 = 721
    switch (p0 << ((1) & 63)) {
    case -1:
        var v8 int32 = 2147483647
        _ = v8
    }
    v9 := []int{}
    switch "" {
    case "\n\t":
        fmt.Println(p0)
    default:
        var v10 int16 = 0
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return 2147483647
}
func genFn1(p0 int16) string {
    if true {
        var v11 string = "a"
        switch -264 {
        case 32767:
            if true {
                if (float32(3.14159) <= 0.1) {
                    var v12 int16 = (-(p0))
                    _ = v12
                }
                p0 = p0
            }
        }
        _ = v11
    }
    if true {
        v13 := GenS1{F0: -1, F1: 2147483647}
        p0 += (p0 & 1000)
        _ = v13
    }
    _ = p0
    return "hello"
}

func main() {
    var v1 int = -782
    var v2 string = "hello"
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    v5 := make(map[string]uint32)
    var v6 float32 = 0.0
    var v14 uint64 = 42
    v2 = v2
    v15 := []int{2147483647, -199, v1}
    var v16 float32 = (v6 + (-2.25 * v6))
    var v17 uint64 = 0
    var v18 uint = 255
    _ = len(v15)
    v6 -= float32(0.1)
    v16 += float32(float64(0.1))
    v19 := GenS2{F0: v17, F1: 0, F2: (v18 >= (v18 % (((v18 >> ((v18) & 63))) | 1)))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
