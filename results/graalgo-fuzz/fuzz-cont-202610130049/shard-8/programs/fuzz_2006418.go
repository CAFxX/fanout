package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint64; F1 string }
func genFn0(p0 uint, p1 uint16, p2 uint8) uint64 {
    var v6 int16 = -32768
    p1 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 898
}
func genFn1(p0 uint) int32 {
    var v7 uint8 = 223
    v7 = v7
    for fi8 := 0; fi8 < 13; fi8++ {
        p0 += p0
        _ = fi8
    }
    v9 := GenS1{F0: (p0 ^ p0), F1: int32(p0)}
    p0 -= (p0 & (p0 & 2))
    _ = p0
    _ = v7
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{2, 108, 4294967295, 65535}
    v2 := []uint32{42, 65535}
    v3 := make(map[int]uint32)
    var v4 string = ""
    v5 := []string{((v4 + "hello") + v4), "foo bar"}
    fmt.Println(v4)
    v10 := GenS2{F0: 2, F1: "foo bar"}
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(fi11, fi11)
        _ = fi11
    }
    fmt.Println(v4)
    var v12 float32 = 1e10
    var v13 int32 = genFn1(42)
    _ = len(v3)
    v12 = v12
    v12 += 0.1
    v4 = "\n\t"
    v14 := GenS2{F0: 1, F1: v4}
    for fi15 := 0; fi15 < 17; fi15++ {
        v16 := GenS2{F0: 838, F1: ""}
        _ = fi15
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
