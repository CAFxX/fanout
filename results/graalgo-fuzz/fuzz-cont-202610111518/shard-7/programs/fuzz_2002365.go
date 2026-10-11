package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int8; F2 uint }
func genFn0(p0 int64, p1 bool) uint32 {
    switch "\n\t" {
    case "a":
        switch "\n\t" {
        case "":
            p0 = (p0 >> ((42) & 63))
        default:
            v4 := GenS1{F0: -379}
            _ = v4
        }
    default:
        v5 := GenS1{F0: 2147483647}
        _ = v5
    }
    switch -305 {
    case 127:
        fmt.Println(p0)
    case 1:
        var v6 uint64 = 18446744073709551615
        _ = v6
    case -746:
        v7 := GenS1{F0: -2147483648}
        _ = v7
    }
    if (!p1) {
        var v8 string = ("hello" + "hello")
        _ = v8
    } else {
        fmt.Println(p0, p0)
        v9 := []int{127, 32768, 128, -9223372036854775808}
        var v10 uint16 = 32767
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 32768}
    _ = p0
    _ = p1
    _ = v11
    return 2147483648
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    var v3 uint64 = 4294967295
    v3 += (^(v3))
    var v12 int16 = -748
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := GenS2{F0: (false || false)}
        var v15 uint64 = 1
        v2[fi13] = 256
        _ = fi13
        _ = v14
        _ = v15
    }
    v3 -= (v3 | (v3 - 0))
    var v16 uint16 = 1
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 float32 = 1.5
        _ = fi17
        _ = v18
    }
    switch 127 {
    case 2147483647:
        var v19 int64 = -9223372036854775808
        _ = v19
    case 128:
        v20 := GenS3{F0: 42, F1: 110, F2: 1}
        _ = v20
    default:
        v16 += 32767
    }
    var v21 int16 = 42
    var v22 int8 = -128
    for fi23 := 0; fi23 < 8; fi23++ {
        v2[fi23] = 4294967295
        switch fi23 {
        case 661:
            for fi24 := 0; fi24 < 14; fi24++ {
                var v25 uint16 = v16
                v16 += (v16 << ((4294967295) & 63))
                v22 = v22
                _ = fi24
                _ = v25
            }
        case 32767:
            v26 := GenS3{F0: (v22 >> ((690) & 63)), F1: v22, F2: 42}
            _ = v26
        case 2147483647:
            var v27 int16 = v12
            _ = v27
        default:
            fmt.Println(v3, v16)
        }
        var v28 string = "a"
        _ = fi23
        _ = v28
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        var v30 uint8 = 217
        if ((0 >= 18446744073709551615) || false) {
            delete(v1, -9223372036854775808)
            v3 = (^(v3))
            v12 -= (v12 / ((42) | 1))
        } else {
            var v31 int64 = 984
            for fi32 := 0; fi32 < 10; fi32++ {
                if ((v3 % ((v3) | 1)) < (v3 >> ((75) & 63))) {
                    v33 := []int{}
                    var v34 int32 = -2147483648
                    _ = v33
                    _ = v34
                }
                v16 += v16
                v12 = (^(v12))
                _ = fi32
            }
            _ = v31
        }
        _ = fi29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
