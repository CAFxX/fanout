package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint }
type GenS2 struct { F0 int16; F1 int64; F2 int8 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0() int32 {
    var v6 bool = true
    var v7 float32 = (0.1 * float32(3.14159))
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int64, p1 uint8) uint {
    var v8 float32 = 0.1
    v9 := GenS1{F0: 633, F1: 1, F2: 635}
    v10 := []int{131, 9223372036854775807, -407}
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 914
}

func main() {
    var v1 float32 = (float32((0.1 / (-2.25 - float64(3.14159)))) / ((0.1 + 3.14159) + 0.0))
    v2 := make(map[string]string)
    v3 := []uint32{685}
    v4 := make(map[uint32]uint32)
    var v5 int16 = 480
    delete(v2, ("a" + ("foo bar" + ("foo bar" + "\n\t"))))
    _ = len(v2)
    v11 := GenS2{F0: v5, F1: 2, F2: -1}
    v5 += v5
    fmt.Println(v5, v5)
    v12 := GenS1{F0: 2, F1: 1, F2: 255}
    v13 := GenS3{F0: 195, F1: -1}
    fmt.Println(v1)
    var v14 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
