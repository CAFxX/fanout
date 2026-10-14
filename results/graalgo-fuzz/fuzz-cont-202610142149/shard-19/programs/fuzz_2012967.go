package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 string }
func genFn0(p0 int16, p1 uint, p2 bool) int16 {
    p0 += p0
    if true {
        p2 = (!false)
    }
    v7 := []int{1, 32767, 225, -855}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}
func genFn1(p0 int, p1 uint) int8 {
    var v8 int64 = -304
    v9 := GenS1{F0: -85}
    fmt.Println(p1, p1, v8)
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -103
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    v3 := []int{-9223372036854775808, -2147483648, 32768, 9223372036854775807}
    v4 := []int{127}
    var v5 uint = 2
    var v6 float64 = 3.14159
    var v10 int16 = 1
    v11 := GenS2{F0: 42, F1: "\n\t"}
    var v12 uint8 = 2
    if len(v4) > 0 {
        v4[0] = int(v12)
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        v10 = v10
        _ = fi13
    }
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    switch "a" {
    case "foo bar":
        v12 = (v12 * v12)
    case "a":
        v5 -= (^(v5))
    }
    var v14 int64 = 9223372036854775807
    v12 += (v12 << ((v5) & 63))
    var v15 int8 = 1
    v14 = 1
    var v16 int = 2147483647
    v15 += (v15 << (((v5 - v5)) & 63))
    fmt.Println(v12, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
