package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
func genFn0(p0 uint64, p1 int16, p2 string) uint32 {
    v5 := GenS1{F0: (p0 >> ((0) & 63)), F1: 4294967295}
    var v6 uint16 = 1
    v7 := []int{0, 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 1
    var v3 int16 = int16(v2)
    v4 := make(map[int]uint32)
    v3 -= (v3 / ((v3) | 1))
    v2 += -1
    fmt.Println(v2, v2)
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(fi8, v3)
        _ = fi8
    }
    v9 := GenS1{F0: 1, F1: 255}
    v10 := GenS1{F0: 943, F1: 274}
    for fi11 := 0; fi11 < 6; fi11++ {
        v3 -= (v3 << ((4294967295) & 63))
        var v12 int16 = v3
        _ = fi11
        _ = v12
    }
    var v13 uint32 = 2147483648
    v2 -= (v2 | v2)
    var v14 string = "foo bar"
    var v15 uint8 = 127
    fmt.Println(v13, v3)
    var v16 uint8 = 128
    var v17 int32 = v2
    v17 = 1
    v18 := []int{2147483647, 951}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
