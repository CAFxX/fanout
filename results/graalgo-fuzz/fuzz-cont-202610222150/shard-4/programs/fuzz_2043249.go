package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
func genFn0() bool {
    var v7 uint = 0
    v8 := GenS1{F0: "hello", F1: 42}
    _ = v7
    _ = v8
    return (1 == 1)
}
func genFn1(p0 uint, p1 string, p2 int64) uint8 {
    v9 := []int{-400, 127}
    p2 = (-(p2))
    v10 := GenS1{F0: "hello", F1: 21}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 0
}

func main() {
    v1 := []string{"\n\t"}
    var v2 uint64 = 4294967295
    v3 := []uint32{2147483647, 743, 65535, 65535}
    v4 := []uint32{2147483647}
    var v5 float32 = 0.0
    v6 := []int{}
    var v11 float32 = ((v5 / (float32(0.0) + (v5 + -2.25))) / v5)
    v11 = v5
    if ((1000 <= -32768) || (v2 != v2)) {
        var v12 int = 127
        var v13 string = ("" + ("" + "a"))
        v12 = (^(v12))
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: "hello", F1: -128}
    fmt.Println(v2, v11)
    var v15 uint32 = 65535
    fmt.Println(v15, v2)
    fmt.Println(v11, v15)
    var v16 int32 = -1
    var v17 float32 = v11
    if len(v1) > 0 {
        v1[0] = ("\n\t" + "a")
    }
    v16 -= 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
