package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint16, p1 int64, p2 int) uint8 {
    v6 := []int{p2, -2147483648}
    v7 := GenS1{F0: (^(p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    var v4 int = -1
    var v5 int16 = 1
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := []int{fi8, v4, v4, -99}
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: true, F1: 520}
    v11 := GenS1{F0: 0}
    v12 := GenS1{F0: 255}
    var v13 int16 = v5
    if (128 <= 248) {
        v14 := []int{v4, v4, 32768}
        _ = v14
    } else {
        var v15 int8 = 1
        _ = v15
    }
    v1 += (^(v1))
    var v16 uint = 18446744073709551615
    for fi17 := 0; fi17 < 20; fi17++ {
        delete(v3, 2)
        var v18 int8 = 1
        var v19 string = "foo bar"
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v1 += v1
        _ = fi20
    }
    v21 := []int{127, 975}
    var v22 uint8 = 67
    for fi23 := 0; fi23 < 4; fi23++ {
        v13 = (v13 >> ((263) & 63))
        v21 = append(v21, -1)
        v1 -= v1
        _ = fi23
    }
    v16 += v16
    v4 -= (v4 % (((v4 ^ (v4 & (v4 * v4)))) | 1))
    v2["a"] = true
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
