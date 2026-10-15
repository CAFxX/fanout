package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 bool; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 uint16 }
func genFn0() uint32 {
    var v6 int16 = -721
    v7 := GenS1{F0: 1}
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS2{F0: "a", F1: true, F2: 255}
        if (!true) {
            var v10 float32 = 3.14159
            _ = v10
        } else {
            var v11 int16 = (v6 % ((v6) | 1))
            v12 := []int{}
            var v13 int8 = 100
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi8
        _ = v9
    }
    var v14 uint64 = 2
    _ = v14
    _ = v6
    _ = v7
    return 110
}
func genFn1(p0 uint32) int {
    var v15 int = -2147483648
    switch "foo bar" {
    case "\n\t":
        v15 -= (v15 % (((v15 | (v15 / (((v15 * 4294967295)) | 1)))) | 1))
    }
    _ = p0
    _ = v15
    return (v15 + (v15 | v15))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 2
    var v3 float32 = 1.5
    var v4 float64 = -2.25
    var v5 int8 = -99
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 uint64 = 4294967295
        switch "foo bar" {
        case "hello":
            v17 = (^(v17))
        case "\n\t":
            v18 := GenS3{F0: 127, F1: 578}
            _ = v18
        case "":
            var v19 float64 = v4
            _ = v19
        default:
            v20 := GenS4{F0: -9223372036854775808}
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    v2 = v2
    var v21 uint16 = 693
    v5 = v5
    if false {
        v2 -= (v2 / ((4294967295) | 1))
        v22 := GenS1{F0: 140}
        v4 = (v4 - v4)
        _ = v22
    }
    v21 -= (v21 << ((1) & 63))
    if (!(534 == 0)) {
        switch "hello" {
        case "hello":
            switch 127 {
            case 32768:
                for fi23 := 0; fi23 < 7; fi23++ {
                    var v24 int16 = 245
                    _ = fi23
                    _ = v24
                }
            case -2147483648:
                v21 -= 854
            case -542:
                if (!false) {
                    var v25 uint16 = (v21 & v21)
                    for fi26 := 0; fi26 < 9; fi26++ {
                        v5 += v5
                        _ = len(v1)
                        _ = fi26
                    }
                    _ = v25
                } else {
                    var v27 uint64 = 2
                    v2 = genFn0()
                    _ = v27
                }
            }
        }
        v28 := GenS3{F0: v5, F1: 2}
        _ = v28
    } else {
        switch -922 {
        case 128:
            v29 := GenS5{F0: v2, F1: 1}
            _ = v29
        }
        fmt.Println(v2, v4)
        v1[42] = (^(v2))
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        if (-2147483648 <= -1) {
            if (false && (v5 != v5)) {
                var v31 int8 = v5
                _ = v31
            }
        }
        fmt.Println(fi30, v2, v5)
        _ = fi30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
