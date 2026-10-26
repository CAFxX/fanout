package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }

func main() {
    var v1 int64 = -827
    var v2 string = "\n\t"
    v3 := make(map[int]uint32)
    var v4 int32 = 1
    var v5 string = "hello"
    var v6 string = (v5 + (v2 + (((v5 + v5) + "foo bar") + v5)))
    var v7 uint8 = 130
    v8 := []int{0, 128, 42}
    v9 := v3[-412]
    v10 := GenS1{F0: 11, F1: -128}
    var v11 uint64 = 443
    v6 = ("hello" + ("hello" + (("\n\t" + v6) + ("\n\t" + v6))))
    v5 = v6
    v12 := GenS1{F0: v11, F1: -79}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
