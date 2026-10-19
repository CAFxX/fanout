package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int16; F2 uint }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint }
func genFn0() uint16 {
    var v4 uint8 = 86
    fmt.Println(v4)
    fmt.Println(v4)
    _ = v4
    return 78
}

func main() {
    var v1 uint32 = 42
    var v2 uint32 = (v1 & v1)
    v3 := make(map[uint32]string)
    switch "" {
    case "a":
        switch "a" {
        case "":
            v1 -= 289
        case "a":
            for fi5 := 0; fi5 < 6; fi5++ {
                var v6 int8 = 0
                _ = fi5
                _ = v6
            }
        default:
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := GenS1{F0: -369}
                v2 += v1
                if true {
                    _ = len(v3)
                }
                _ = fi7
                _ = v8
            }
        }
    case "hello":
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 uint = 0
            _ = fi9
            _ = v10
        }
    case "":
        v1 += (v1 << ((2) & 63))
    default:
        var v11 int64 = 42
        _ = v11
    }
    fmt.Println(v1, v1)
    v1 = (v2 / ((v2) | 1))
    var v12 int32 = 647
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := []int{}
        var v15 uint = 2
        for fi16 := 0; fi16 < 17; fi16++ {
            switch "foo bar" {
            case "":
                v1 = (v2 * v2)
            default:
                var v17 uint = v15
                _ = v17
            }
            _ = fi16
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v2 = (v1 % ((v1) | 1))
    switch "a" {
    case "a":
        v18 := GenS1{F0: v12}
        _ = v18
    case "hello":
        v12 -= -2147483648
    default:
        v12 -= (v12 | 593)
    }
    v19 := GenS2{F0: -32, F1: -265, F2: 1}
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            v22 := v3[v1]
            v23 := GenS3{F0: 4294967295, F1: 42, F2: 18446744073709551615}
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
}
