package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint8 }
func genFn0(p0 int32) int16 {
    if (42 == 42) {
        p0 = (^(p0))
        var v4 uint32 = 4294967295
        _ = v4
    }
    var v5 int32 = p0
    fmt.Println(p0, p0)
    _ = p0
    _ = v5
    return 1
}
func genFn1(p0 uint16, p1 uint16, p2 uint8) bool {
    switch "hello" {
    case "foo bar":
        v6 := GenS1{F0: 77, F1: -32768, F2: (p2 / ((p2) | 1))}
        _ = v6
    case "":
        var v7 uint32 = 627
        _ = v7
    case "hello":
        v8 := []int{431}
        _ = v8
    default:
        fmt.Println(p1, p0)
    }
    var v9 float32 = (1.5 / 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return false
}

func main() {
    var v1 uint16 = 2
    var v2 uint64 = 0
    var v3 int64 = 111
    if (!(!false)) {
        var v10 uint = 1
        _ = v10
    }
    v11 := GenS1{F0: 1, F1: 1, F2: 128}
    v1 = 32767
    var v12 int8 = 102
    var v13 uint32 = 42
    v2 += v2
    var v14 string = "hello"
    fmt.Println(v13, v1, v2)
    fmt.Println(v13, v2)
    v15 := []int{838, 127, 4294967295}
    var v16 int32 = -2147483648
    var v17 int32 = (^(v16))
    var v18 uint64 = 415
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    v2 = (v18 & (v18 ^ 4294967295))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
