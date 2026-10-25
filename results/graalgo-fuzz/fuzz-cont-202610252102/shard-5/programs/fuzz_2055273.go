package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 bool }
type GenS2 struct { F0 int32; F1 int64; F2 uint16 }

func main() {
    v1 := []int{-1, 4294967295, -451, 32767}
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]string)
    v5 := []string{"", "\n\t", "hello", "foo bar"}
    v6 := make(map[int]bool)
    var v7 float64 = 0.0
    var v8 int = -9223372036854775808
    fmt.Println(v7, v7)
    var v9 uint32 = 4294967295
    v8 = 32768
    var v10 float64 = 0.0
    var v11 float32 = float32(float64((3.14159 + 3.14159)))
    fmt.Println(v10, v8)
    v3[(v9 - (v9 | 2147483647))] = v9
    fmt.Println(v11, v9, v9)
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 bool = ("hello" != "")
        _ = fi12
        _ = v13
    }
    var v14 int32 = 302
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS1{F0: "hello", F1: 2, F2: false}
        v17 := GenS2{F0: (v14 | v14), F1: 835, F2: 65535}
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
