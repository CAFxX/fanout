package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint16 }
func genFn0(p0 int, p1 uint64, p2 int64) int64 {
    if (!false) {
        for fi5 := 0; fi5 < 2; fi5++ {
            for fi6 := 0; fi6 < 8; fi6++ {
                switch fi6 {
                case 2147483647:
                    fmt.Println(fi6, fi5, p2)
                case 1:
                    var v7 int = -1
                    _ = v7
                }
                _ = fi6
            }
            for fi8 := 0; fi8 < 2; fi8++ {
                p0 = (fi5 >> ((464) & 63))
                var v9 bool = true
                fmt.Println(fi5, fi8, p2)
                _ = fi8
                _ = v9
            }
            p1 += (p1 % ((p1) | 1))
            _ = fi5
        }
    } else {
        switch p0 {
        case 0:
            v10 := []int{}
            _ = v10
        case 4294967295:
            fmt.Println(p2, p2)
        case -594:
            var v11 uint8 = 1
            _ = v11
        }
        fmt.Println(p0, p2)
    }
    p1 = p1
    v12 := []int{p0, (^(p0)), (-(p0)), (p0 ^ p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return (p2 | 0)
}
func genFn1(p0 bool, p1 uint, p2 int32) uint32 {
    fmt.Println(p1)
    v13 := GenS1{F0: 127, F1: 2, F2: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 859
}

func main() {
    var v1 uint8 = 122
    var v2 int32 = 2147483647
    var v3 float64 = -2.25
    v4 := make(map[uint32]bool)
    fmt.Println(v2, v1)
    v3 -= (v3 * v3)
    v14 := GenS1{F0: 1, F1: 787, F2: 255}
    fmt.Println(v3, v3, v2)
    var v15 float32 = -2.25
    var v16 string = "\n\t"
    v16 = v16
    v2 += (v2 & v2)
    v17 := GenS1{F0: -128, F1: 255, F2: 1}
    v2 = (v2 ^ 567)
    var v18 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
