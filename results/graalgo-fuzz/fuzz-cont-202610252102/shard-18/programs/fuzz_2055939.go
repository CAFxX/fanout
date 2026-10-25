package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int8; F1 uint32; F2 uint16 }

func main() {
    var v1 int = 1
    var v2 int64 = 0
    var v3 int = v1
    v1 -= 490
    var v4 uint16 = 2
    switch "a" {
    case "foo bar":
        v5 := GenS1{F0: 42, F1: "hello"}
        _ = v5
    case "a":
        v6 := GenS1{F0: 565, F1: (("" + "") + "hello")}
        _ = v6
    }
    v7 := []int{v1, (v3 ^ (v1 / (((v3 * v1)) | 1))), -834, (v1 % (((v3 % ((v1) | 1))) | 1))}
    var v8 uint = 2
    v3 = (v3 / ((-795) | 1))
    fmt.Println(v4, v8, v2)
    v2 = 0
    var v9 uint = uint(v4)
    if len(v7) > 0 {
        v7[0] = 127
    }
    v9 = 18446744073709551615
    if false {
        var v10 uint32 = 42
        if len(v7) > 0 {
            v7[0] = (v1 / ((v1) | 1))
        }
        _ = v10
    }
    v2 = v2
    v7 = append(v7, 2147483647)
    v11 := GenS2{F0: 100, F1: 2, F2: (^(v4))}
    v9 -= v8
    fmt.Println(v4, v9, v8)
    if false {
        v7 = append(v7, -1)
    } else {
        var v12 float32 = -2.25
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
