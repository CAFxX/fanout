package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int16 }
func genFn0(p0 uint64, p1 uint16, p2 uint) uint64 {
    v6 := []int{-1, -2147483648, 127}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p0 & p0)
}
func genFn1(p0 int32, p1 uint16, p2 bool) int8 {
    p1 = 0
    p1 = 1
    _ = p0
    _ = p1
    _ = p2
    return 100
}

func main() {
    v1 := []string{("" + "foo bar"), ("hello" + ""), "foo bar", ("\n\t" + "\n\t")}
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    v5 := []int{-1, 1, 32768, -1}
    var v7 int = 2147483647
    var v8 uint64 = 0
    var v9 uint64 = genFn0(v8, 255, 255)
    var v10 bool = false
    fmt.Println(v9)
    v11 := GenS1{F0: 0, F1: 65535, F2: 1000}
    var v12 int8 = -1
    v8 = genFn0(v9, 65535, 2)
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := GenS1{F0: -2147483648, F1: uint(v9), F2: 1}
        var v15 int8 = 100
        for fi16 := 0; fi16 < 7; fi16++ {
            _ = len(v5)
            v17 := GenS1{F0: -2147483648, F1: 0, F2: 0}
            v12 = v15
            fmt.Println(v7, v15, v7)
            _ = fi16
            _ = v17
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v8 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
