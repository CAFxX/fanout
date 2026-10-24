package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 uint64; F2 bool }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    v4 := []int{-535}
    v5 := make(map[uint32]int)
    if (!true) {
        fmt.Println(v1)
        switch -1 {
        case 0:
            v1 = (v1 << ((1) & 63))
        default:
            var v6 int16 = 0
            _ = v6
        }
    }
    v1 = v1
    if true {
        switch 4294967295 {
        case -458:
            var v7 int64 = -9223372036854775808
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 6; fi8++ {
            switch "a" {
            case "foo bar":
                fmt.Println(fi8, v1)
            case "":
                fmt.Println(v1)
            }
            _ = fi8
        }
    }
    v1 -= 0
    if true {
        if (-2.25 < (((1e10 - -2.25) - float32(0.1)) + 0.1)) {
            v9 := GenS1{F0: "\n\t", F1: 255, F2: true}
            if len(v4) > 0 {
                v4[0] = 32768
            }
            _ = v9
        } else {
            v1 += (v1 & (v1 | v1))
            fmt.Println(v1)
            var v10 int8 = 105
            _ = v10
        }
        v11 := GenS2{F0: int(v1)}
        _ = v11
    }
    v1 += (v1 & (^(v1)))
    v12 := GenS3{F0: 574, F1: 42, F2: (1 > 259)}
    v1 = v1
    switch 1 {
    case 2147483647:
        fmt.Println(v1)
    default:
        var v13 bool = true
        _ = v13
    }
    _ = len(v2)
    for fi14 := 0; fi14 < 2; fi14++ {
        fmt.Println(fi14, fi14)
        v3[("foo bar" + "foo bar")] = (-(fi14))
        _ = fi14
    }
    v1 = (v1 & (v1 % ((v1) | 1)))
    if true {
        var v15 int8 = 83
        _ = v15
    }
    fmt.Println(v1)
    v1 += (-(v1))
    var v16 uint32 = 482
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
