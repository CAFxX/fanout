package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int32 }
func genFn0(p0 int32, p1 string, p2 int32) uint {
    p0 = (p2 / ((p2) | 1))
    var v6 int8 = -128
    fmt.Println(v6, p2)
    fmt.Println(p0, v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 666
}
func genFn1() int {
    var v7 uint16 = 667
    fmt.Println(v7)
    v7 = (v7 & 65535)
    _ = v7
    return -700
}

func main() {
    var v1 int16 = 42
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    var v4 float32 = ((0.0 / 1.5) / (0.1 * (((((-2.25 * 1.5) + float32(1e10)) - ((-2.25 - 1.5) * 3.14159)) - (3.14159 / 3.14159)) - 1e10)))
    var v5 uint8 = 206
    var v8 uint32 = 267
    v5 = 255
    var v9 bool = (!true)
    v3[("hello" + ("a" + "\n\t"))] = uint32(v5)
    switch 2147483647 {
    case -168:
        v10 := GenS1{F0: 2147483647, F1: -930}
        _ = v10
    }
    v11 := GenS2{F0: 124}
    v12 := []int{42, -1, genFn1()}
    fmt.Println(v5)
    if v9 {
        var v13 float32 = v4
        _ = len(v12)
        _ = v13
    }
    var v14 float32 = (float32(((3.14159 / 1.0) + (3.14159 / 1.5))) - 0.1)
    var v15 uint8 = v5
    v16 := GenS3{F0: 127, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
