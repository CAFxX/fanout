package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int64 }
func genFn0(p0 bool) bool {
    v4 := GenS1{F0: 4, F1: 42, F2: 964}
    var v5 int8 = -49
    _ = p0
    _ = v4
    _ = v5
    return (p0 || (("\n\t" + "hello") > "hello"))
}
func genFn1(p0 int) int16 {
    v6 := []int{(p0 + p0), p0}
    fmt.Println(p0)
    _ = p0
    _ = v6
    return -1
}

func main() {
    v1 := []string{}
    var v2 uint32 = 42
    var v3 int = 268
    var v7 uint16 = 874
    switch "a" {
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "\n\t":
        var v8 int32 = 1
        _ = v8
    case "foo bar":
        v9 := GenS1{F0: 65535, F1: 2147483647, F2: -846}
        _ = v9
    default:
        v10 := GenS1{F0: 255, F1: -1, F2: 9223372036854775807}
        _ = v10
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    var v11 float32 = 1.5
    v7 -= v7
    var v12 int64 = 0
    fmt.Println(v3, v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
}
