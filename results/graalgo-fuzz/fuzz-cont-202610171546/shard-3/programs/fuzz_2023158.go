package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }

func main() {
    var v1 uint16 = 937
    var v2 int16 = 978
    var v3 bool = (int64(v1) >= 9223372036854775807)
    v4 := make(map[uint32]int)
    v5 := []int{4294967295, 32767}
    var v6 uint8 = 128
    v2 -= v2
    v6 = (v6 << ((0) & 63))
    var v7 uint64 = 4294967295
    var v8 string = ("\n\t" + "a")
    fmt.Println(v6, v1, v7)
    v9 := GenS1{F0: false, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
