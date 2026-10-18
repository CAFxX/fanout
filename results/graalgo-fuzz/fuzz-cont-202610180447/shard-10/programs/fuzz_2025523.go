package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint8; F1 int; F2 string }

func main() {
    v1 := []int{32767, 151}
    v2 := []string{"hello", "\n\t", "\n\t", "hello"}
    var v3 float64 = (float64((3.14159 * 0.1)) + (1.5 / float64(1.5)))
    v4 := make(map[int]string)
    v5 := []int{-538, 4294967295, -9223372036854775808}
    var v6 float64 = 0.1
    if false {
        var v7 bool = false
        _ = v7
    } else {
        v3 -= (-2.25 / float64(1.5))
        var v8 float64 = (v3 * 3.14159)
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := GenS1{F0: 42, F1: 0}
        fmt.Println(v6, fi9, fi9)
        _ = fi9
        _ = v10
    }
    switch "a" {
    case "":
        var v11 int32 = 42
        _ = v11
    case "foo bar":
        v3 -= 0.1
    case "\n\t":
        if (921 == 0) {
            v12 := GenS1{F0: 127, F1: 42}
            var v13 float64 = v3
            _ = v12
            _ = v13
        }
    default:
        v14 := GenS2{F0: 2, F1: 32768, F2: ((("hello" + "\n\t") + "\n\t") + "hello")}
        _ = v14
    }
    var v15 int64 = -1
    fmt.Println(v3)
    if (!false) {
        var v16 uint32 = 42
        var v17 int16 = -32768
        _ = v16
        _ = v17
    } else {
        var v18 uint32 = 0
        var v19 float64 = 0.1
        _ = v18
        _ = v19
    }
    v6 = -2.25
    fmt.Println(v6)
    var v20 uint8 = 212
    switch "hello" {
    case "\n\t":
        var v21 string = ("a" + "a")
        _ = v21
    default:
        fmt.Println(v15, v20)
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 uint64 = 18446744073709551615
        _ = fi22
        _ = v23
    }
    v15 = -1
    v3 = (((float64(1.5) + (3.14159 * v6)) + 0.1) + v6)
    var v24 int16 = 1000
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
