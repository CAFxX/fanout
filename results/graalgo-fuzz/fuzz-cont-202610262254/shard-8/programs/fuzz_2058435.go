package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int }

func main() {
    var v1 uint16 = 783
    var v2 uint16 = v1
    v3 := []string{"hello", "hello", (("" + ("foo bar" + "hello")) + (("\n\t" + "\n\t") + "\n\t"))}
    var v4 uint16 = v1
    v5 := make(map[uint32]int)
    v6 := GenS1{F0: (2147483647 < 2147483647), F1: 968, F2: 63}
    v4 += (v1 ^ 255)
    v3 = append(v3, "")
    v1 = (v4 * v1)
    v2 = 919
    var v7 int64 = 388
    v8 := v5[42]
    var v9 uint32 = 2
    fmt.Println(v7, v7, v2)
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
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
