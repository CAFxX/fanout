package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
func genFn0() int8 {
    var v5 int32 = 2147483647
    switch 127 {
    case 9223372036854775807:
        v6 := []int{128, 32767, 0, 635}
        _ = v6
    }
    _ = v5
    return 93
}
func genFn1(p0 uint64, p1 string, p2 uint64) string {
    var v7 float64 = 3.14159
    v8 := GenS1{F0: 734, F1: (p1 + p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ""
}

func main() {
    var v1 uint64 = 265
    var v2 int = -9223372036854775808
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    v9 := GenS1{F0: 326, F1: (("foo bar" + (("a" + "a") + "")) + "\n\t")}
    var v10 int64 = 9223372036854775807
    switch "a" {
    case "\n\t":
        v11 := v4[494]
        _ = v11
    case "foo bar":
        v12 := []int{42, (v2 / ((v2) | 1)), v2, (v2 / ((v2) | 1))}
        _ = v12
    case "hello":
        v13 := []int{-2147483648, int(v10)}
        _ = v13
    }
    v1 = v1
    var v14 uint = 0
    var v15 uint32 = 4294967295
    v16 := GenS1{F0: 1, F1: "hello"}
    _ = len(v4)
    v17 := []int{v2}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
