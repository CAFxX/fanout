package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint32 }

func main() {
    var v1 int = 2147483647
    v2 := []uint32{4294967295, 4294967295, 2147483647, 2147483647}
    v3 := make(map[uint32]int)
    var v4 string = ("a" + (("\n\t" + ((("" + "foo bar") + ("hello" + "")) + ("foo bar" + "a"))) + ("foo bar" + ("hello" + "\n\t"))))
    v5 := make(map[int]uint32)
    v1 = v1
    var v6 int64 = -9223372036854775808
    v4 = v4
    v7 := GenS1{F0: v6, F1: 128, F2: 2147483648}
    _ = len(v2)
    var v8 int64 = v6
    _ = len(v2)
    v5[v1] = 4294967295
    var v9 uint16 = 255
    v6 = (^(v8))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
