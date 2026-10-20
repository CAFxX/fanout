package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 uint }
type GenS2 struct { F0 int16 }
func genFn0() uint {
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int64 = -688
        fmt.Println(v6, fi5)
        v7 := []int{}
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := []int{(fi8 ^ (^(fi8))), -473, (fi8 | (fi8 << ((4294967295) & 63)))}
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 int32 = -889
            var v12 int32 = (v11 * v11)
            v11 -= (v12 % ((v11) | 1))
            _ = fi10
            _ = v11
            _ = v12
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi8
        _ = v9
    }
    v13 := GenS1{F0: -2147483648, F1: 0, F2: 2}
    var v14 uint64 = 4294967295
    _ = v13
    _ = v14
    return 18446744073709551615
}
func genFn1(p0 string, p1 int64) uint64 {
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS2{F0: 199}
        _ = fi15
        _ = v16
    }
    p0 = (p0 + p0)
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 42
    v3 := []int{4294967295}
    var v4 uint64 = v2
    v2 -= (^(v4))
    switch 9223372036854775807 {
    case -168:
        v2 = 540
    case -9223372036854775808:
        v17 := GenS2{F0: -32768}
        _ = v17
    default:
        for fi18 := 0; fi18 < 14; fi18++ {
            if true {
                if (127 < 239) {
                    fmt.Println(fi18, fi18)
                } else {
                    v19 := GenS2{F0: -32768}
                    _ = v19
                }
                var v20 string = ""
                var v21 uint8 = 0
                _ = v20
                _ = v21
            }
            v4 = 0
            _ = fi18
        }
    }
    v4 = genFn1((("hello" + "\n\t") + "\n\t"), 236)
    v22 := v1[0]
    fmt.Println(v22)
    var v23 int32 = -305
    if len(v3) > 0 {
        v3[0] = 128
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "\n\t" {
    case "":
        v22 = (!(v22 || v22))
    case "\n\t":
        v23 -= v23
    default:
        var v24 int = 2147483647
        _ = v24
    }
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
