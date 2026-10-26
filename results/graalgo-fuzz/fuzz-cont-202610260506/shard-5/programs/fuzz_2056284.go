package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{}
    var v3 int16 = -1
    v4 := make(map[uint32]bool)
    var v5 bool = false
    fmt.Println(v3, v3)
    var v6 bool = (v5 || (!(!(1 <= -122))))
    v6 = v6
    v3 -= v3
    v6 = ((126 >= 42) && (42 < 1))
    var v7 int = 1
    v8 := GenS1{F0: 42, F1: 0, F2: 42}
    fmt.Println(v6, v6, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
