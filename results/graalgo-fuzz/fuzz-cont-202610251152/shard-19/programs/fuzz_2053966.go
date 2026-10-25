package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint64; F1 uint }
func genFn0(p0 int16, p1 uint32, p2 string) int {
    v5 := GenS1{F0: 255, F1: 605, F2: (p2 + p2)}
    if false {
        var v6 bool = (!(0 < 0))
        switch 32767 {
        case 473:
            var v7 uint16 = 91
            _ = v7
        default:
            if (((^(p1)) == p1) || v6) {
                var v8 uint64 = 0
                if false {
                    fmt.Println(p1, v8)
                    v9 := []int{127}
                    p1 = p1
                    _ = v9
                }
                _ = v8
            } else {
                var v10 uint8 = 0
                var v11 uint16 = 65535
                var v12 uint8 = 29
                _ = v10
                _ = v11
                _ = v12
            }
        }
        _ = v6
    } else {
        var v13 float32 = float32(float64(-2.25))
        var v14 uint = 1
        fmt.Println(p0, v14)
        _ = v13
        _ = v14
    }
    p0 -= (p0 << ((591) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 128
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 0
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    for fi15 := 0; fi15 < 4; fi15++ {
        switch genFn0(-1, v2, "hello") {
        case -447:
            v16 := GenS2{F0: 983}
            _ = v16
        }
        v2 = 0
        switch 9223372036854775807 {
        case 32767:
            v2 = 42
        default:
            v17 := GenS3{F0: -266, F1: ((-1 >= 100) || false), F2: 157}
            _ = v17
        }
        _ = fi15
    }
    v2 = v2
    switch -231 {
    case -9223372036854775808:
        v18 := GenS4{F0: 297, F1: 426}
        _ = v18
    default:
        v2 += (^(v2))
    }
    var v19 int8 = 100
    var v20 uint64 = 18446744073709551615
    _ = len(v1)
    v21 := []int{genFn0(820, (v2 * (v2 & v2)), "foo bar"), -929, 127, 375}
    var v22 int32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
