package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint8, p1 uint8, p2 uint64) uint {
    if ((!false) && true) {
        v4 := GenS1{F0: 65535, F1: 42}
        var v5 uint64 = p2
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 int = fi6
            _ = fi6
            _ = v7
        }
        _ = v4
        _ = v5
    }
    fmt.Println(p0, p2)
    fmt.Println(p2)
    var v8 bool = (0.0 <= 0.1)
    v9 := []int{1, 32768, 313}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 1
    var v3 int = (v2 ^ v2)
    var v10 int8 = -128
    if false {
        _ = len(v1)
        var v11 uint8 = 42
        var v12 string = "a"
        _ = v11
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 float32 = 0.0
            _ = fi13
            _ = v14
        }
        v15 := GenS2{F0: 0, F1: 758}
        v2 += (v3 << ((809) & 63))
        _ = v15
    }
    var v16 uint16 = 124
    v17 := GenS2{F0: 127, F1: 329}
    var v18 uint16 = 42
    switch v2 {
    case 128:
        v19 := []int{v2, -1, (v2 % ((v3) | 1))}
        _ = v19
    default:
        fmt.Println(v10, v10, v3)
    }
    var v20 int16 = 1000
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
}
