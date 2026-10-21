package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[string]int)
    v3 := []int{-2147483648, 9223372036854775807, 128, 720}
    var v4 uint32 = 0
    var v5 uint8 = 255
    v6 := GenS1{F0: v5}
    var v7 uint8 = (v5 - v5)
    fmt.Println(v7)
    var v8 bool = (1 < 901)
    v2["\n\t"] = 528
    v7 += v5
    v3 = append(v3, 2147483647)
    v9 := GenS1{F0: 127}
    var v10 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
