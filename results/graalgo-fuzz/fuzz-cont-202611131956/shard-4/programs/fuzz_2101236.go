package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int8 }
func genFn0(p0 uint32) int {
    var v7 int64 = 42
    v7 -= (v7 & v7)
    if true {
        v7 = v7
        var v8 uint64 = 456
        if true {
            v9 := []int{}
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = v9
        }
        _ = v8
    } else {
        var v10 string = (("foo bar" + ("foo bar" + "foo bar")) + "hello")
        _ = v10
    }
    v11 := []int{2147483647, 42, 1}
    v12 := GenS1{F0: 42, F1: "hello", F2: -1}
    _ = p0
    _ = v11
    _ = v12
    _ = v7
    return 32768
}
func genFn1(p0 uint, p1 uint) bool {
    var v13 uint32 = 42
    p0 -= (p1 >> ((p0) & 63))
    v14 := GenS2{F0: ""}
    var v15 float32 = (0.0 + (-2.25 * 0.0))
    v13 += v13
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return true
}

func main() {
    var v1 bool = (-74 <= 0)
    v2 := []int{-9223372036854775808, 9223372036854775807, 998, 127}
    var v3 float64 = (float64(0.0) + 1e10)
    var v4 int8 = 88
    var v5 uint16 = 42
    v6 := make(map[string]string)
    var v16 float32 = 1.5
    v17 := GenS3{F0: int64(v4), F1: 1}
    fmt.Println(v4, v16, v4)
    var v18 int8 = (v4 - (v4 + (v4 << ((1) & 63))))
    var v19 uint8 = 1
    var v20 float32 = v16
    v21 := GenS2{F0: "a"}
    var v22 float32 = (v20 * v16)
    fmt.Println(v16)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
