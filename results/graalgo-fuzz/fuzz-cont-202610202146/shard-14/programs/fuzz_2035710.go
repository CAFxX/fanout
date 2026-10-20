package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint }
func genFn0(p0 string) bool {
    var v7 uint = 65535
    if true {
        for fi8 := 0; fi8 < 4; fi8++ {
            for fi9 := 0; fi9 < 19; fi9++ {
                if true {
                    fmt.Println(v7)
                    var v10 uint16 = 229
                    v10 = (v10 / (((^(v10))) | 1))
                    v11 := GenS1{F0: uint64(v7)}
                    _ = v10
                    _ = v11
                } else {
                    v12 := GenS1{F0: 1}
                    v13 := []int{(fi8 >> ((v7) & 63)), fi9, fi8}
                    _ = len(v13)
                    _ = v12
                    _ = v13
                }
                _ = fi9
            }
            _ = fi8
        }
    }
    p0 = "\n\t"
    _ = p0
    _ = v7
    return (-189 < 542)
}
func genFn1(p0 uint, p1 uint8) uint64 {
    v14 := []int{2147483647, 32767, -227, -284}
    var v15 float64 = (0.0 + 3.14159)
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    var v4 string = ((("" + "\n\t") + "foo bar") + "")
    v5 := make(map[int]bool)
    v6 := []uint32{0, 1, 4294967295}
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 uint32 = 61
        v18 := GenS2{F0: fi16, F1: 0}
        v19 := GenS2{F0: -608, F1: 9223372036854775807}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    switch 128 {
    case 2147483647:
        for fi20 := 0; fi20 < 3; fi20++ {
            fmt.Println(fi20)
            _ = fi20
        }
    case -497:
        v4 = (v4 + v4)
    }
    var v21 uint32 = 137
    if (127 >= 1) {
        v21 = (v21 * (v21 - (v21 & (^(v21)))))
        v22 := GenS3{F0: 255, F1: 18446744073709551615}
        _ = v22
    }
    v4 = v4
    v21 -= (v21 * v21)
    for fi23 := 0; fi23 < 3; fi23++ {
        v4 = v4
        var v24 uint8 = 127
        v4 = ((v4 + "hello") + "")
        _ = fi23
        _ = v24
    }
    v21 = (v21 | (v21 ^ v21))
    if true {
        v21 = v21
    }
    v25 := GenS2{F0: 42, F1: 9223372036854775807}
    var v26 string = (v4 + "hello")
    fmt.Println(v21)
    if true {
        var v27 uint16 = 12
        _ = v27
    }
    v28 := GenS1{F0: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
