package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint16, p1 uint8) bool {
    p0 -= 42
    p1 += 0
    if ((((2 == 18446744073709551615) || false) && false) && (p0 == (p0 >> ((2) & 63)))) {
        if (1 != uint64(p0)) {
            for fi5 := 0; fi5 < 9; fi5++ {
                v6 := GenS1{F0: 65535, F1: p0}
                fmt.Println(p1)
                _ = fi5
                _ = v6
            }
        }
        p1 += p1
        p0 += p0
    }
    v7 := []int{}
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 uint16 = 42
        v10 := GenS2{F0: 42}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return (false && (true && true))
}

func main() {
    var v1 int8 = -128
    var v2 uint32 = 4294967295
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]uint32)
    var v11 float64 = (-2.25 / (1e10 / 0.1))
    v12 := GenS2{F0: 42}
    var v13 bool = false
    v14 := GenS3{F0: 42}
    v1 = (v1 % ((v1) | 1))
    v15 := GenS2{F0: -1}
    var v16 int32 = 2147483647
    if v13 {
        v3["a"] = v2
        if v13 {
            v17 := []int{1}
            v18 := GenS2{F0: 1}
            _ = v17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 int8 = (v1 << ((255) & 63))
        for fi21 := 0; fi21 < 19; fi21++ {
            v11 += (((-2.25 / v11) / 1e10) - v11)
            v22 := GenS3{F0: 85}
            _ = fi21
            _ = v22
        }
        if (((((v11 * v11) * v11) / v11) - v11) > v11) {
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 uint64 = 1
                v13 = v13
                _ = fi23
                _ = v24
            }
            var v25 uint8 = 255
            _ = v25
        }
        _ = fi19
        _ = v20
    }
    v1 = v1
    var v26 bool = v13
    v16 = v16
    var v27 int = -485
    for fi28 := 0; fi28 < 6; fi28++ {
        v11 += v11
        _ = fi28
    }
    var v29 uint32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
}
