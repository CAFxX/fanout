package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int; F2 uint8 }
func genFn0(p0 uint32, p1 uint8, p2 int32) uint16 {
    p0 = (p0 / (((p0 ^ (p0 & (^(p0))))) | 1))
    var v6 bool = true
    var v7 uint8 = (p1 | 127)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int64, p1 int32, p2 int64) uint32 {
    v8 := GenS1{F0: 218}
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 int = (fi9 * (fi9 & fi9))
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 156
}

func main() {
    v1 := []uint32{0, 2147483648}
    v2 := make(map[uint32]int)
    v3 := []int{127, -156, 0}
    v4 := make(map[int]bool)
    v5 := make(map[string]string)
    var v11 int64 = -905
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 bool = true
        _ = fi12
        _ = v13
    }
    v14 := GenS2{F0: 0, F1: 42, F2: 1}
    v11 = 1
    fmt.Println(v11)
    fmt.Println(v11)
    switch "foo bar" {
    case "hello":
        switch 0 {
        case 32767:
            fmt.Println(v11)
        default:
            fmt.Println(v11)
        }
    case "a":
        if false {
            v11 -= 9223372036854775807
        }
    }
    var v15 int16 = 1000
    var v16 uint16 = 391
    var v17 uint32 = 2147483648
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 bool = true
        v11 += (v11 >> ((0) & 63))
        _ = fi18
        _ = v19
    }
    switch "hello" {
    case "foo bar":
        if (genFn0((v17 & 2147483648), 104, -2147483648) <= v16) {
            var v20 uint16 = v16
            var v21 uint64 = 2
            v15 = -1
            _ = v20
            _ = v21
        } else {
            var v22 uint32 = v17
            if (0.1 >= 3.14159) {
                v22 -= genFn1(-1, 42, v11)
                if true {
                    v22 -= (v22 * v17)
                    v3 = append(v3, -9223372036854775808)
                    var v23 int16 = (v15 << ((273) & 63))
                    _ = v23
                }
                v16 = v16
            }
            fmt.Println(v15)
            _ = v22
        }
    default:
        var v24 string = (("a" + ("" + "foo bar")) + "foo bar")
        _ = v24
    }
    if (42 != 127) {
        if (true || (!false)) {
            var v25 float64 = 0.0
            _ = v25
        }
        for fi26 := 0; fi26 < 1; fi26++ {
            var v27 int8 = 59
            v28 := GenS1{F0: 65535}
            _ = fi26
            _ = v27
            _ = v28
        }
    } else {
        var v29 int64 = 9223372036854775807
        var v30 int64 = (v11 & (v11 - (v11 ^ (v29 + v11))))
        _ = v29
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
