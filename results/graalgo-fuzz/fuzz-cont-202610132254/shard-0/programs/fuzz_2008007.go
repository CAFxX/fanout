package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    var v3 int8 = 100
    var v4 uint8 = 162
    v3 = (v3 * v3)
    v4 += (v4 >> ((413) & 63))
    v3 = (v3 / (((v3 & (v3 * 42))) | 1))
    v5 := GenS1{F0: 646, F1: 114}
    var v6 float32 = 1.5
    v3 -= (v3 | (^(v3)))
    var v7 int8 = (v3 & (v3 / (((v3 & 42)) | 1)))
    var v8 bool = true
    v9 := []int{32768, -509}
    var v10 int16 = 32767
    var v11 string = "hello"
    var v12 int32 = 288
    fmt.Println(v7)
    var v13 uint32 = 2147483648
    v6 = v6
    v3 -= v3
    var v14 string = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
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
