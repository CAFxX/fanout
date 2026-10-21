package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint }
func genFn0(p0 uint64) uint16 {
    p0 -= 18446744073709551615
    v5 := GenS1{F0: (!((127 >= 42) || (false && true))), F1: 42, F2: 0}
    p0 -= p0
    p0 = (p0 - p0)
    _ = p0
    _ = v5
    return 0
}

func main() {
    v1 := []int{1}
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]string)
    var v4 uint8 = 0
    switch 937 {
    case 32768:
        _ = len(v1)
    }
    var v6 uint8 = (^(v4))
    switch ("\n\t" + "") {
    case "":
        fmt.Println(v2)
    }
    var v7 int16 = 150
    fmt.Println(v2, v2)
    fmt.Println(v7, v7, v7)
    var v8 int16 = (-(v7))
    fmt.Println(v7, v8, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
