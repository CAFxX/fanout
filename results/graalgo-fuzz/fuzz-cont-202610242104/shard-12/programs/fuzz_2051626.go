package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint) bool {
    if ((!(2147483648 < 290)) && false) {
        if (false && true) {
            var v4 int = -9223372036854775808
            for fi5 := 0; fi5 < 14; fi5++ {
                if true {
                    var v6 uint64 = 1
                    var v7 uint = (p0 % ((p0) | 1))
                    v8 := GenS1{F0: (v7 >> ((p0) & 63))}
                    _ = v6
                    _ = v7
                    _ = v8
                }
                p0 += (p0 << ((2) & 63))
                _ = fi5
            }
            _ = v4
        }
    }
    p0 -= (^(p0))
    p0 += 65535
    _ = p0
    return false
}
func genFn1(p0 uint8, p1 string) int {
    v9 := GenS2{F0: 255}
    fmt.Println(p0, p1)
    var v10 uint64 = 1
    var v11 uint32 = 0
    var v12 uint8 = (p0 | p0)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 0
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[uint32]uint32)
    var v3 float32 = 3.14159
    v2[42] = 65535
    switch 128 {
    case -1:
        for fi13 := 0; fi13 < 6; fi13++ {
            for fi14 := 0; fi14 < 4; fi14++ {
                v15 := GenS3{F0: 101, F1: v1}
                v16 := []int{-1, fi14}
                _ = fi14
                _ = v15
                _ = v16
            }
            var v17 int = fi13
            var v18 uint = 681
            _ = fi13
            _ = v17
            _ = v18
        }
    case 4294967295:
        var v19 float64 = 1.5
        _ = v19
    }
    var v20 int8 = -128
    delete(v2, 666)
    v21 := []int{178, genFn1(127, "\n\t"), -752}
    switch 127 {
    case 9223372036854775807:
        v1 = (v1 + v1)
    case 911:
        fmt.Println(v1)
    }
    var v22 float32 = v3
    v21 = append(v21, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
