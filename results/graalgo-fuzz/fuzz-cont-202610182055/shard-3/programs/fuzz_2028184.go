package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int }
func genFn0(p0 int16, p1 uint32, p2 uint32) int16 {
    p1 -= p1
    for fi4 := 0; fi4 < 15; fi4++ {
        v5 := GenS1{F0: -109, F1: p1}
        v6 := GenS1{F0: -1, F1: 65535}
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint16 = 1
        _ = fi7
        _ = v8
    }
    var v9 uint16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p0
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{}
    var v3 uint16 = 792
    fmt.Println(v3)
    v3 = v3
    for fi10 := 0; fi10 < 11; fi10++ {
        v3 += (v3 << ((2) & 63))
        fmt.Println(v3)
        _ = fi10
    }
    switch 32767 {
    case -9223372036854775808:
        var v11 bool = (55 >= 271)
        _ = v11
    case 4294967295:
        fmt.Println(v3)
    default:
        v3 += (v3 - v3)
    }
    if true {
        var v12 int64 = 9223372036854775807
        v3 = v3
        _ = v12
    } else {
        var v13 int = -9223372036854775808
        _ = v13
    }
    v14 := GenS2{F0: 449}
    fmt.Println(v3)
    for fi15 := 0; fi15 < 17; fi15++ {
        if (1 >= -1) {
            for fi16 := 0; fi16 < 4; fi16++ {
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v3 = v3
                    v17 := GenS2{F0: fi16}
                    _ = v17
                }
                v18 := GenS3{F0: 2147483648, F1: (fi16 * fi15)}
                _ = fi16
                _ = v18
            }
            var v19 int8 = -128
            _ = v19
        } else {
            v3 -= 1
            fmt.Println(v3, fi15)
            var v20 int8 = 42
            _ = v20
        }
        v3 -= v3
        _ = fi15
    }
    var v21 float32 = 1e10
    delete(v1, 953)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
}
