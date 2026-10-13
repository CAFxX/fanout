package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint16 }
func genFn0(p0 bool, p1 bool, p2 uint) uint64 {
    for fi5 := 0; fi5 < 11; fi5++ {
        p2 = (p2 * (p2 | p2))
        _ = fi5
    }
    switch -2147483648 {
    case -350:
        switch 9223372036854775807 {
        case 9223372036854775807:
            var v6 bool = (p0 && ((p0 || ((^(p2)) >= (^(p2)))) || true))
            _ = v6
        case -568:
            fmt.Println(p1)
        default:
            var v7 uint8 = 0
            _ = v7
        }
    case -661:
        if (!false) {
            var v8 float32 = 0.1
            _ = v8
        }
    default:
        if p0 {
            v9 := []int{-290, 1}
            p2 = (p2 | p2)
            _ = v9
        }
    }
    v10 := []int{-139, 335}
    p0 = false
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 2
}
func genFn1(p0 int16, p1 int8, p2 int8) int64 {
    var v11 string = "hello"
    var v12 int8 = (p2 - p2)
    if true {
        var v13 uint8 = 190
        fmt.Println(v11, v13, v12)
        var v14 uint64 = 1
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return -9223372036854775808
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint16 = 1
    var v3 uint = 0
    v4 := make(map[uint32]bool)
    v2 += v2
    v1 -= (v1 >> (((v3 / ((v3) | 1))) & 63))
    var v15 int16 = -502
    v1 = v1
    v16 := GenS1{F0: -652, F1: "foo bar", F2: v2}
    var v17 int8 = 1
    for fi18 := 0; fi18 < 12; fi18++ {
        v3 = v3
        v17 = v17
        _ = fi18
    }
    v17 -= (v17 / ((0) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
