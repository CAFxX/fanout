package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int32 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint }

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 0.0
    v3 := make(map[int]bool)
    v4 := []int{584, -9223372036854775808}
    v5 := []uint32{}
    var v6 int32 = 0
    v3[2147483647] = (!((54 < 127) || false))
    var v7 uint = 102
    var v8 float32 = 1.5
    v9 := GenS1{F0: -556, F1: 253, F2: v6}
    v10 := GenS2{F0: -1, F1: 671, F2: (v7 & v7)}
    if false {
        if true {
            v4 = append(v4, -1)
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            delete(v1, 2147483648)
            _ = fi11
        }
    }
    var v12 float64 = (3.14159 - v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
