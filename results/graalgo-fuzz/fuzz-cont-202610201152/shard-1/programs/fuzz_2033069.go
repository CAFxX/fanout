package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 int64 = -655
    var v2 float64 = 1.5
    var v3 uint8 = 4
    var v4 int64 = (v1 & v1)
    v5 := make(map[uint32]string)
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint8 = (v3 << ((824) & 63))
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: 42}
            if true {
                v10 := []int{fi6}
                if true {
                    fmt.Println(v1)
                    var v11 uint = 1
                    v7 += v3
                    _ = v11
                }
                _ = v10
            } else {
                v12 := GenS1{F0: 489}
                _ = v12
            }
            v1 = v4
            v4 += (v1 ^ v4)
            v7 += (v3 / ((36) | 1))
            v1 = (v1 * (v1 | 82))
            _ = fi8
            _ = v9
        }
        v13 := GenS2{F0: 42, F1: 5}
        _ = fi6
        _ = v13
        _ = v7
    }
    _ = len(v5)
    var v14 int32 = 0
    v15 := []int{9223372036854775807, 32768}
    v1 += 1
    var v16 uint16 = 2
    switch 2147483647 {
    case -496:
        fmt.Println(v4)
    case -574:
        v17 := GenS3{F0: 0}
        _ = v17
    case -1:
        var v18 bool = false
        _ = v18
    default:
        var v19 uint64 = 42
        _ = v19
    }
    if false {
        v20 := GenS1{F0: (v16 % (((^(v16))) | 1))}
        _ = v20
    }
    if (42 < 791) {
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    } else {
        for fi21 := 0; fi21 < 3; fi21++ {
            if ((v14 & v14) <= 1) {
                fmt.Println(v14, v1, v3)
            } else {
                if true {
                    var v22 int32 = int32(v3)
                    _ = v22
                }
            }
            _ = fi21
        }
        if len(v15) > 0 {
            v15[0] = 13
        }
        _ = len(v5)
    }
    if false {
        var v23 uint32 = 242
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
