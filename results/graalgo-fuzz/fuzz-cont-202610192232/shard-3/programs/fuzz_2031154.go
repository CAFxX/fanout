package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int8; F1 bool; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint }
func genFn0(p0 string, p1 uint, p2 uint64) int32 {
    if (false && false) {
        v5 := []int{32767, 9223372036854775807}
        _ = v5
    } else {
        var v6 uint32 = 1
        v6 += (^(v6))
        _ = v6
    }
    p0 = p0
    fmt.Println(p1)
    fmt.Println(p1)
    var v7 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -562
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]uint32)
    var v4 int64 = 0
    v4 = v4
    v4 = v4
    if ((!(1 != 2)) || (!false)) {
        v8 := GenS1{F0: 18446744073709551615, F1: -506}
        switch 4294967295 {
        case -2147483648:
            var v9 uint16 = 255
            _ = v9
        case 32767:
            v4 += v4
        default:
            v10 := []int{1}
            _ = v10
        }
        var v11 int32 = -2147483648
        _ = v11
        _ = v8
    }
    v12 := []int{}
    v13 := GenS1{F0: 18446744073709551615, F1: 32767}
    if (true || (false || (false || (2 <= 42)))) {
        switch 32768 {
        case 128:
            for fi14 := 0; fi14 < 9; fi14++ {
                fmt.Println(v4)
                v4 -= (v4 ^ v4)
                _ = fi14
            }
        case 615:
            v15 := GenS2{F0: 127, F1: (!false), F2: 10}
            _ = v15
        default:
            fmt.Println(v4)
        }
        if true {
            v16 := GenS3{F0: 32767}
            _ = v16
        } else {
            _ = len(v12)
            if (0 < 255) {
                var v17 uint32 = 312
                var v18 uint64 = 18446744073709551615
                v4 += (^(v4))
                _ = v17
                _ = v18
            }
            var v19 string = ""
            _ = v19
        }
    }
    var v20 int16 = 42
    for fi21 := 0; fi21 < 13; fi21++ {
        fmt.Println(v20, fi21, v20)
        v22 := GenS4{F0: 18446744073709551615, F1: 0}
        _ = fi21
        _ = v22
    }
    v4 = (v4 % ((1) | 1))
    v1[2147483647] = -642
    v20 = v20
    v20 += (v20 % ((v20) | 1))
    v20 = v20
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
