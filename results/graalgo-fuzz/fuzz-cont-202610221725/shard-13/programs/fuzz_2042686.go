package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint64, p1 uint) uint {
    p1 = (p1 + 0)
    v4 := []int{2147483647, 32768, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v4
    return (p1 >> ((p1) & 63))
}

func main() {
    var v1 int8 = -1
    var v2 int64 = 149
    var v3 int8 = -1
    v5 := GenS1{F0: 718}
    v6 := []int{85, -9223372036854775808, 0, 32768}
    fmt.Println(v2, v3, v2)
    if len(v6) > 0 {
        v6[0] = 32768
    }
    var v7 bool = true
    fmt.Println(v7, v2)
    v8 := GenS1{F0: 9223372036854775807}
    var v9 float64 = (0.0 + 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
