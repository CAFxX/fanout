package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int32 }
func genFn0(p0 int) uint32 {
    var v6 int64 = 703
    v6 += (v6 * (v6 & 979))
    v7 := GenS1{F0: -465, F1: 42, F2: -1}
    _ = p0
    _ = v6
    _ = v7
    return 32
}
func genFn1(p0 string, p1 uint8, p2 uint16) uint64 {
    v8 := GenS1{F0: -720, F1: 4294967295, F2: -47}
    var v9 uint16 = (p2 - 255)
    if false {
        fmt.Println(p0, p1)
        var v10 int16 = 42
        _ = v10
    }
    v11 := []int{0, 756, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint32 = 835
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    v4 := []string{("hello" + "a"), ((("hello" + "hello") + "hello") + "hello")}
    var v5 int32 = -818
    if false {
        v1 = (v1 + 42)
    } else {
        v12 := GenS1{F0: 128, F1: 930, F2: (v5 << ((943) & 63))}
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 int = fi13
            v14 = (fi13 >> ((736) & 63))
            var v15 uint16 = 2
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v12
    }
    v5 = v5
    fmt.Println(v1)
    v16 := GenS1{F0: 32768, F1: 0, F2: (v5 & (v5 ^ v5))}
    v1 -= (v1 % (((^(v1))) | 1))
    var v17 uint32 = (v1 & v1)
    v18 := GenS1{F0: 32768, F1: 866, F2: v5}
    v1 = (v17 >> ((4294967295) & 63))
    fmt.Println(v1, v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
