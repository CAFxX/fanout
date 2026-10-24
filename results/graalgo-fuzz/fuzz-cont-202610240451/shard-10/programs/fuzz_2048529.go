package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint16 }
func genFn0(p0 int8, p1 string) uint16 {
    fmt.Println(p0)
    p0 += p0
    var v7 uint = 514
    for fi8 := 0; fi8 < 14; fi8++ {
        v7 = v7
        v9 := []int{fi8, (fi8 * (^(fi8))), 0}
        var v10 uint8 = uint8(v7)
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 450
}
func genFn1(p0 uint) int64 {
    p0 += (p0 | (^(p0)))
    p0 -= (p0 ^ p0)
    _ = p0
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{26, 32767, 0}
    var v3 int32 = -614
    v4 := make(map[int]int)
    var v5 int16 = 0
    var v6 int = 0
    v3 = (^(v3))
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 int = (fi11 + (fi11 * 25))
        v5 = v5
        v13 := GenS1{F0: 65535, F1: 193, F2: genFn0(114, ("foo bar" + (("a" + "\n\t") + "hello")))}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 float32 = 0.0
    v5 = 32767
    v3 = (^(v3))
    var v15 uint32 = 42
    var v16 bool = true
    v6 += (v6 >> ((18446744073709551615) & 63))
    fmt.Println(v3, v14, v14)
    var v17 float32 = (v14 + 0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
