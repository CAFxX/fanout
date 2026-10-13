package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }

func main() {
    v1 := []int{128}
    v2 := make(map[int]string)
    v3 := []uint32{1, 987, 680}
    var v4 uint = 395
    v4 -= v4
    v4 -= 2
    var v5 float32 = (-2.25 * ((-2.25 - float32(3.14159)) / float32((1.5 - (-2.25 - (0.0 - float64(((((0.1 / float32(-2.25)) * 0.0) * -2.25) + float32(1e10)))))))))
    v6 := GenS1{F0: -1, F1: 1}
    var v7 string = ("a" + ("" + "hello"))
    v4 = (v4 + (^(v4)))
    if false {
        v8 := GenS1{F0: 565, F1: 42}
        v4 = v4
        _ = v8
    } else {
        v7 = v7
        v9 := GenS1{F0: -524, F1: 677}
        _ = len(v2)
        _ = v9
    }
    var v10 int8 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
