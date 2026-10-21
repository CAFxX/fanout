package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 bool }
func genFn0(p0 uint8) uint8 {
    v6 := []int{0, 32768}
    p0 = 0
    _ = p0
    _ = v6
    return (p0 | (p0 * 2))
}
func genFn1(p0 bool) string {
    v7 := []int{127, 2147483647, 128}
    var v8 int = 492
    v8 -= (-(v8))
    _ = p0
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    var v4 string = "foo bar"
    var v5 uint16 = 98
    fmt.Println(v4, v5, v5)
    var v9 int16 = 1000
    switch ("\n\t" + ("foo bar" + v4)) {
    case "\n\t":
        v5 += (v5 + (v5 * 2))
    default:
        var v10 uint8 = 128
        _ = v10
    }
    v11 := []int{}
    _ = len(v3)
    var v12 uint64 = 18446744073709551615
    v13 := GenS1{F0: (v5 << ((2) & 63)), F1: 773, F2: (-2.25 >= 1e10)}
    var v14 int32 = 42
    var v15 uint16 = v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
