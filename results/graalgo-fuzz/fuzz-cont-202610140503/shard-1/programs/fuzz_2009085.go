package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }

func main() {
    v1 := []uint32{4294967295, 2, 1}
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    v6 := []int{128, 32767}
    var v7 uint8 = 94
    var v8 uint16 = 1
    v8 += 32767
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 string = "foo bar"
        _ = fi9
        _ = v10
    }
    fmt.Println(v8)
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 float64 = (0.0 + 0.0)
        _ = fi11
        _ = v12
    }
    fmt.Println(v8)
    fmt.Println(v8, v7)
    fmt.Println(v8, v7)
    v8 = (v8 % (((v8 - (v8 >> ((780) & 63)))) | 1))
    v7 -= v7
    var v13 uint32 = 0
    var v14 int64 = 9223372036854775807
    v15 := GenS1{F0: 199, F1: (v8 >> ((0) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
