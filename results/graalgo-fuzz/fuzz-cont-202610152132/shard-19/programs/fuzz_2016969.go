package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }

func main() {
    var v1 int8 = 100
    v2 := make(map[int]string)
    var v3 float32 = float32(1.5)
    v4 := []uint32{518, 1, 267}
    v5 := make(map[uint32]string)
    v6 := []uint32{65535, 2, 1}
    v4 = append(v4, 406)
    v7 := GenS1{F0: -1, F1: 2147483647}
    v3 = ((v3 / 3.14159) / float32(-2.25))
    v8 := GenS1{F0: 9223372036854775807, F1: -198}
    v2[32767] = ("foo bar" + (("foo bar" + "hello") + "hello"))
    _ = len(v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
