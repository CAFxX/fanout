package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int }
func genFn0(p0 string, p1 uint32) uint {
    var v4 int = 128
    var v5 uint = 402
    fmt.Println(v5, v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (v5 ^ uint(p1))
}

func main() {
    var v1 uint32 = 89
    var v2 uint32 = (v1 & (v1 - (v1 * v1)))
    var v3 bool = (0 < -956)
    v1 += v2
    v2 = (v2 << ((genFn0("\n\t", (v2 << ((2) & 63)))) & 63))
    v6 := GenS1{F0: 208, F1: 255, F2: (v2 ^ v1)}
    v7 := []int{0, -1, 762, -102}
    v3 = v3
    var v8 uint32 = 2147483648
    v7 = append(v7, 2147483647)
    v3 = (-9223372036854775808 >= 32768)
    var v9 uint = 1
    v8 = v8
    switch 32768 {
    case -2147483648:
        if true {
            if v3 {
                var v10 uint16 = 510
                var v11 uint8 = 42
                v11 = 2
                _ = v10
                _ = v11
            }
            var v12 float32 = 1.5
            for fi13 := 0; fi13 < 9; fi13++ {
                fmt.Println(fi13)
                switch (fi13 % ((42) | 1)) {
                case 9223372036854775807:
                    v14 := GenS2{F0: 0}
                    _ = v14
                case 1:
                    var v15 uint8 = 2
                    _ = v15
                default:
                    var v16 int16 = 0
                    _ = v16
                }
                _ = fi13
            }
            _ = v12
        } else {
            if len(v7) > 0 {
                v7[0] = 42
            }
        }
    case -891:
        var v17 uint32 = v8
        _ = v17
    case 570:
        if len(v7) > 0 {
            v7[0] = -185
        }
    }
    if (!(!v3)) {
        v18 := GenS2{F0: 0}
        _ = len(v7)
        for fi19 := 0; fi19 < 15; fi19++ {
            var v20 string = ("a" + "a")
            _ = fi19
            _ = v20
        }
        _ = v18
    } else {
        v21 := GenS3{F0: uint32(v9), F1: 4294967295, F2: 32768}
        v8 = v8
        if len(v7) > 0 {
            v7[0] = -2147483648
        }
        _ = v21
    }
    fmt.Println(v9, v9)
    var v22 int32 = -1
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
