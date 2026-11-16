package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int }

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]uint32)
    v3 := []uint32{2}
    var v4 int = 128
    v5 := []int{128, (v4 | (v4 - v4)), v4, 127}
    var v6 int32 = -701
    var v7 uint64 = 750
    v4 += (v4 % ((128) | 1))
    fmt.Println(v1, v6, v7)
    for fi8 := 0; fi8 < 19; fi8++ {
        switch 212 {
        case 0:
            fmt.Println(v6, v4, fi8)
        case 127:
            v7 = v7
        case -644:
            var v9 uint = 2
            _ = v9
        default:
            var v10 float64 = 1e10
            _ = v10
        }
        _ = fi8
    }
    fmt.Println(v7, v4)
    v4 = (v4 + 32767)
    v11 := GenS1{F0: v4, F1: 65535, F2: v4}
    var v12 int8 = 42
    for fi13 := 0; fi13 < 14; fi13++ {
        if false {
            if true {
                var v14 int32 = v6
                fmt.Println(v4)
                _ = v14
            } else {
                switch v4 {
                case 864:
                    fmt.Println(v4)
                case -87:
                    var v15 uint16 = 0
                    _ = v15
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            for fi16 := 0; fi16 < 17; fi16++ {
                fmt.Println(v4, v4, v7)
                v6 = v6
                _ = fi16
            }
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            v1 -= 128
            _ = fi17
        }
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
