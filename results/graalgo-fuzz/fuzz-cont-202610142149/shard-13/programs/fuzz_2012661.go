package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int16 }
func genFn0(p0 int16, p1 int, p2 bool) int {
    var v6 int8 = 127
    p2 = ("\n\t" != "hello")
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (^(p1))
}
func genFn1(p0 int, p1 uint, p2 uint64) uint16 {
    p1 = (p1 & (p1 % ((p1) | 1)))
    var v7 uint32 = 1
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 bool = (2 >= 18446744073709551615)
    v4 := make(map[string]string)
    v5 := make(map[string]int)
    fmt.Println(v3)
    v3 = false
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int64 = -1
        _ = fi8
        _ = v9
    }
    var v10 int32 = 2147483647
    fmt.Println(v3)
    for fi11 := 0; fi11 < 12; fi11++ {
        fmt.Println(v3)
        v10 = (v10 >> ((4294967295) & 63))
        _ = fi11
    }
    v12 := []int{-2147483648, -2147483648, 471, 2147483647}
    v13 := GenS1{F0: 186, F1: v3, F2: 0}
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 int64 = -9223372036854775808
        _ = fi14
        _ = v15
    }
    v3 = (0 != 949)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
