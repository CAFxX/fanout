package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint, p1 uint16) int {
    var v4 int16 = -32768
    v5 := []int{-9223372036854775808, -1}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -173
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    var v3 float64 = 1e10
    var v6 string = "hello"
    v7 := []int{128}
    _ = len(v2)
    var v8 int64 = 1
    delete(v2, "foo bar")
    fmt.Println(v6, v8, v6)
    v9 := GenS1{F0: 32767, F1: 128}
    delete(v2, v6)
    var v10 bool = (127 >= 128)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
