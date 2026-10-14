package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint64, p1 int16) uint {
    v6 := GenS1{F0: 321, F1: 42}
    p1 = (p1 - (^(p1)))
    p0 = (p0 & p0)
    _ = p0
    _ = p1
    _ = v6
    return 325
}
func genFn1(p0 uint64, p1 int, p2 int32) uint8 {
    fmt.Println(p1, p0, p1)
    v7 := []int{}
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            _ = len(v7)
            v7 = append(v7, 32767)
            _ = fi9
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 127
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 811
    v3 := []uint32{0}
    var v4 float64 = 3.14159
    var v5 int64 = v2
    for fi10 := 0; fi10 < 8; fi10++ {
        v5 = (^(v2))
        delete(v1, fi10)
        _ = len(v1)
        _ = fi10
    }
    var v11 float32 = ((0.1 * 1e10) * (3.14159 - float32(float64(float32(3.14159)))))
    v12 := GenS2{F0: "a", F1: 257}
    v4 += 3.14159
    v5 += v2
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi13 := 0; fi13 < 1; fi13++ {
            for fi14 := 0; fi14 < 15; fi14++ {
                fmt.Println(v5, fi13)
                var v15 uint16 = 88
                _ = fi14
                _ = v15
            }
            v16 := GenS3{F0: v2}
            _ = fi13
            _ = v16
        }
        for fi17 := 0; fi17 < 16; fi17++ {
            fmt.Println(v11, v5, v11)
            _ = fi17
        }
    }
    fmt.Println(v11, v11)
    v11 = (1.5 / (((v11 - (-2.25 / (v11 / (v11 - 3.14159)))) / 1.5) * v11))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
