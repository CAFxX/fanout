package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 string; F1 int64; F2 int }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int8; F1 bool }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 32767
    var v3 int = -9223372036854775808
    var v4 float64 = (1.5 + float64((1e10 * 3.14159)))
    v5 := GenS1{F0: 9223372036854775807, F1: false}
    v4 += (1e10 + v4)
    var v6 int8 = -104
    var v7 uint32 = 4294967295
    v8 := v1[(v3 % (((v3 * (-(v3)))) | 1))]
    fmt.Println(v8, v2)
    v7 = (v7 - v7)
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS2{F0: v8, F1: -1, F2: (v3 + v3)}
        _ = fi9
        _ = v10
    }
    var v11 uint16 = (^(v2))
    v2 += (v2 - v2)
    v12 := []int{(v3 + 4294967295), v3, (v3 - v3), v3}
    v13 := GenS3{F0: 152, F1: (v7 % ((v7) | 1))}
    v14 := GenS3{F0: 0, F1: 844}
    v15 := v1[(v3 - v3)]
    v16 := GenS4{F0: v6, F1: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
