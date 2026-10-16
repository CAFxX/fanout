package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint; F1 int16; F2 int }
func genFn0(p0 int, p1 uint32, p2 int64) int {
    var v5 uint8 = 177
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: 42, F1: 675, F2: 1}
        v8 := []int{127, (p0 % ((fi6) | 1))}
        _ = len(v8)
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -393
}

func main() {
    var v1 bool = (false && true)
    v2 := make(map[uint32]bool)
    var v3 int8 = -59
    v4 := make(map[int]string)
    var v9 int16 = -1
    switch "a" {
    case "a":
        v3 = (^(v3))
    case "foo bar":
        v9 = (v9 << ((2) & 63))
    case "hello":
        v1 = (0 != 240)
    default:
        v9 += (v9 << ((1) & 63))
    }
    var v10 int64 = -1
    v11 := []int{42, 345, 4294967295, 2147483647}
    v9 -= v9
    var v12 float32 = ((0.1 / 1e10) * 0.1)
    v11 = append(v11, -9223372036854775808)
    fmt.Println(v1, v3)
    v4[-653] = "foo bar"
    switch 4294967295 {
    case -36:
        for fi13 := 0; fi13 < 14; fi13++ {
            v3 = 64
            v3 -= (-(v3))
            _ = fi13
        }
    case 32767:
        fmt.Println(v1, v9, v3)
    default:
        v12 += float32(3.14159)
    }
    var v14 uint64 = 2
    v1 = v1
    var v15 int8 = v3
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 int64 = v10
        v14 = 18446744073709551615
        _ = fi16
        _ = v17
    }
    v10 += int64(v9)
    for fi18 := 0; fi18 < 1; fi18++ {
        v19 := GenS2{F0: 1, F1: -1, F2: genFn0(4294967295, 1, 573)}
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
