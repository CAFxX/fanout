package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
func genFn0(p0 int8, p1 uint64) int8 {
    v6 := []int{-669, -1, -878}
    var v7 int8 = (p0 - (p0 | -86))
    v8 := GenS1{F0: -714, F1: v7}
    fmt.Println(p1, p0)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 uint16 = 1
    v2 := []uint32{uint32(v1), 65535}
    v3 := make(map[uint32]string)
    v4 := []uint32{181}
    v5 := make(map[uint32]int)
    delete(v3, 4294967295)
    v1 += (v1 << ((18446744073709551615) & 63))
    var v9 uint64 = 18446744073709551615
    fmt.Println(v9, v9)
    var v10 uint64 = 18446744073709551615
    v11 := GenS1{F0: 2147483647, F1: 127}
    v12 := GenS1{F0: 0, F1: genFn0(1, (v9 / ((v10) | 1)))}
    if false {
        v13 := GenS1{F0: -2147483648, F1: 99}
        var v14 bool = true
        var v15 uint16 = v1
        _ = v13
        _ = v14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
