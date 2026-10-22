package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 bool }
func genFn0() uint {
    for fi5 := 0; fi5 < 10; fi5++ {
        fmt.Println(fi5)
        v6 := GenS1{F0: 594}
        _ = fi5
        _ = v6
    }
    switch 127 {
    case 2147483647:
        switch "foo bar" {
        case "":
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := GenS2{F0: 54, F1: 128, F2: 127}
                var v9 bool = true
                _ = fi7
                _ = v8
                _ = v9
            }
        default:
            var v10 uint32 = 4294967295
            _ = v10
        }
    default:
        v11 := GenS3{F0: 4294967295, F1: 42}
        _ = v11
    }
    var v12 uint16 = 2
    v12 = (v12 - (^(v12)))
    _ = v12
    return 654
}
func genFn1(p0 uint8, p1 int32) int32 {
    var v13 int64 = 1
    if true {
        p1 -= -7
        p1 += -457
        if false {
            p0 = (^(p0))
            switch -2147483648 {
            case 1:
                v14 := []int{-9223372036854775808, 778, 4294967295, 127}
                _ = v14
            default:
                for fi15 := 0; fi15 < 19; fi15++ {
                    for fi16 := 0; fi16 < 12; fi16++ {
                        p0 = 2
                        var v17 float64 = (-2.25 / (((0.1 / 1e10) * -2.25) / float64((3.14159 * 0.1))))
                        var v18 int8 = 127
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    switch "foo bar" {
                    case "foo bar":
                        var v19 int8 = 40
                        _ = v19
                    case "hello":
                        var v20 uint8 = p0
                        _ = v20
                    default:
                        v13 = (v13 * v13)
                    }
                    _ = fi15
                }
            }
        }
    }
    v21 := GenS2{F0: p0, F1: (p0 + p0), F2: (p0 % (((^(p0))) | 1))}
    var v22 int8 = 1
    _ = p0
    _ = p1
    _ = v13
    _ = v21
    _ = v22
    return (-(p1))
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[uint32]int)
    var v3 uint16 = 32767
    v4 := make(map[int]uint32)
    for fi23 := 0; fi23 < 2; fi23++ {
        v24 := []int{2147483647}
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        v26 := v2[v1]
        var v27 int8 = 1
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint16 = (^(v3))
    fmt.Println(v3, v1, v1)
    v3 -= (v3 + (v3 ^ v3))
    if (-2147483648 <= 863) {
        if false {
            _ = len(v4)
        } else {
            var v29 int32 = -1
            v29 += 2147483647
            if (197 > uint(v1)) {
                v30 := GenS3{F0: 946, F1: 65535}
                _ = v30
            }
            _ = v29
        }
        v31 := []int{467, -2147483648, 0, 108}
        _ = v31
    }
    if (-9223372036854775808 == 1) {
        v32 := []int{}
        _ = v32
    }
    if (-9223372036854775808 == -9223372036854775808) {
        v33 := []int{-1, 1, -9223372036854775808}
        _ = v33
    }
    var v34 int32 = 42
    var v35 uint32 = 170
    v36 := []int{32768, 42, -317}
    if (658 >= 0) {
        v37 := GenS4{F0: true}
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
