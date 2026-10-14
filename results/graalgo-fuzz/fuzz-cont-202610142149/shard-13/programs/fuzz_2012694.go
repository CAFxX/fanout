package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int8; F1 int64; F2 int16 }
func genFn0() uint8 {
    v7 := []int{-1, 416, 42}
    var v8 uint16 = 0
    var v9 int32 = int32(v8)
    v10 := GenS1{F0: -9223372036854775808, F1: -948}
    v11 := GenS1{F0: 0, F1: -637}
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]int)
    var v5 int8 = 42
    var v6 uint8 = 42
    fmt.Println(v5, v5, v6)
    v1 += 2
    v6 -= (v6 % ((v6) | 1))
    v1 = (v1 << ((18446744073709551615) & 63))
    v1 += (v1 / (((v1 - (v1 % ((18446744073709551615) | 1)))) | 1))
    v4[2147483648] = 32767
    switch 309 {
    case 32768:
        switch "hello" {
        case "foo bar":
            v12 := []int{9223372036854775807, 127}
            _ = v12
        default:
            fmt.Println(v6)
        }
    }
    if (288 >= 1) {
        v6 = v6
    }
    switch "" {
    case "":
        _ = len(v2)
    case "a":
        v6 = v6
    case "\n\t":
        var v13 int64 = -634
        _ = v13
    default:
        v1 -= 724
    }
    v14 := []int{}
    delete(v3, 32768)
    delete(v3, 127)
    for fi15 := 0; fi15 < 11; fi15++ {
        _ = len(v14)
        _ = fi15
    }
    v16 := GenS2{F0: v5, F1: 848, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
