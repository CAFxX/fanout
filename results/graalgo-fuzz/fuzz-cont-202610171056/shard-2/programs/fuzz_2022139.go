package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint16, p1 bool, p2 uint16) uint32 {
    var v5 int64 = 42
    switch "a" {
    case "":
        v6 := GenS1{F0: int(v5), F1: p0}
        _ = v6
    default:
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int64 = (v5 / (((v5 << ((1) & 63))) | 1))
            v9 := []int{}
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    if (1 <= 65535) {
        p2 = p0
    }
    v5 = v5
    fmt.Println(v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}
func genFn1(p0 uint32) int64 {
    v10 := []int{0, 127, 0}
    var v11 uint64 = 51
    v11 = 42
    _ = p0
    _ = v10
    _ = v11
    return -824
}

func main() {
    var v1 uint = 42
    var v2 string = ("foo bar" + "")
    var v3 uint8 = 35
    v4 := make(map[int]uint32)
    for fi12 := 0; fi12 < 19; fi12++ {
        switch 127 {
        case 313:
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 bool = (147 != 18446744073709551615)
                v15 := GenS2{F0: 42, F1: 2, F2: 2}
                _ = fi13
                _ = v14
                _ = v15
            }
        case 32768:
            v3 = 167
        default:
            var v16 int8 = 42
            _ = v16
        }
        v1 += v1
        _ = fi12
    }
    v17 := GenS3{F0: 42, F1: -1}
    v1 += (v1 + v1)
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 uint32 = genFn0(65535, (!(-93 > -29)), 1)
        v1 = 18446744073709551615
        _ = fi18
        _ = v19
    }
    v1 -= (^(v1))
    v1 = v1
    switch v2 {
    case "\n\t":
        v3 += (v3 << (((v1 - (v1 / (((v1 + (v1 + v1))) | 1)))) & 63))
    case "a":
        fmt.Println(v1, v1, v2)
    case "":
        var v20 int32 = 42
        _ = v20
    }
    var v21 bool = true
    v1 -= v1
    v22 := GenS1{F0: 127, F1: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
