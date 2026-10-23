package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint }
func genFn0(p0 uint32, p1 int16, p2 int16) uint8 {
    var v5 string = (("" + "") + ("" + (("\n\t" + "a") + "hello")))
    p2 = 0
    if false {
        var v6 uint8 = 42
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 127
}

func main() {
    var v1 int32 = -2147483648
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]uint32)
    var v4 float64 = 0.1
    v7 := GenS1{F0: 127, F1: "hello", F2: 1}
    if true {
        if false {
            v8 := GenS1{F0: 2, F1: "", F2: 642}
            fmt.Println(v1, v1)
            _ = v8
        }
        delete(v3, -203)
        v9 := []int{-2147483648}
        _ = v9
    }
    var v10 float32 = (float32(v4) * float32(3.14159))
    fmt.Println(v1, v4, v1)
    v11 := []int{4294967295, -438, 1}
    v2 = (-(v2))
    v4 = v4
    v12 := GenS1{F0: 100, F1: "foo bar", F2: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
