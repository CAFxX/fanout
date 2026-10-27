package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 string }
func genFn0(p0 int, p1 int) bool {
    v7 := []int{p0, p0, (p0 + (-(p1))), p0}
    p1 = (^(p0))
    fmt.Println(p1, p0)
    p1 = (p1 - (p1 * 127))
    _ = p0
    _ = p1
    _ = v7
    return false
}

func main() {
    var v1 int64 = -403
    var v2 uint64 = 830
    v3 := make(map[int]bool)
    var v4 int = 32767
    v5 := []uint32{4294967295, 0, 640}
    var v6 float64 = float64((-2.25 - (-2.25 - (float32(1e10) / 1e10))))
    var v8 int8 = -34
    v9 := GenS1{F0: 255}
    v1 += 9223372036854775807
    v10 := GenS2{F0: v8, F1: "a"}
    var v11 uint32 = 42
    v2 += v2
    var v12 uint = 42
    var v13 int = v4
    var v14 int8 = v8
    var v15 uint = v12
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
