package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int16, p1 int16, p2 uint) int16 {
    var v6 bool = (!true)
    v7 := []int{2147483647, 250, 4294967295}
    v6 = v6
    p0 = (p1 & (-(p1)))
    p0 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p0 >> (((p2 - (p2 - (p2 % ((p2) | 1))))) & 63))
}

func main() {
    v1 := []int{102}
    var v2 int64 = 1
    var v3 bool = (0.0 <= 1.5)
    v4 := []int{-279, 32767}
    v5 := []uint32{65535, 1}
    v8 := GenS1{F0: 1, F1: "\n\t"}
    fmt.Println(v2)
    v3 = v3
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 4; fi10++ {
            v2 = (v2 - v2)
            v11 := GenS1{F0: 0, F1: "a"}
            v2 -= v2
            _ = fi10
            _ = v11
        }
        _ = len(v5)
        _ = fi9
    }
    v3 = true
    fmt.Println(v2, v2)
    var v12 int32 = 1
    v1 = append(v1, 470)
    var v13 float64 = (((0.0 / 1.5) / 1.0) + (0.1 - 0.0))
    v14 := GenS2{F0: 193}
    var v15 int = 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
