package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int8) int64 {
    v5 := GenS1{F0: 32767}
    v6 := []int{9223372036854775807, 9223372036854775807, 42, -9223372036854775808}
    _ = p0
    _ = v5
    _ = v6
    return -92
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -128
    var v3 uint16 = 255
    v4 := []uint32{606, 840, 810, 0}
    v1[32] = false
    var v7 int8 = v2
    var v8 uint = 255
    var v9 string = ((("" + (("a" + "") + "\n\t")) + "foo bar") + ("hello" + "a"))
    v10 := GenS1{F0: v3}
    v11 := GenS2{F0: 255}
    v1[4294967295] = true
    v8 += (v8 * (v8 + 0))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
