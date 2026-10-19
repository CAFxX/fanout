package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 string }
func genFn0(p0 int16) int32 {
    switch "" {
    case "a":
        p0 = 685
    default:
        switch 127 {
        case -984:
            for fi6 := 0; fi6 < 14; fi6++ {
                p0 = (-(p0))
                for fi7 := 0; fi7 < 19; fi7++ {
                    p0 += (^(p0))
                    p0 += p0
                    p0 -= p0
                    _ = fi7
                }
                _ = fi6
            }
        case 32768:
            for fi8 := 0; fi8 < 14; fi8++ {
                switch (fi8 & fi8) {
                case 2147483647:
                    p0 = p0
                }
                v9 := []int{-2147483648, 439}
                p0 += (p0 & (p0 | (p0 % ((p0) | 1))))
                _ = fi8
                _ = v9
            }
        case -2147483648:
            if true {
                v10 := []int{1, 42, 9223372036854775807, 0}
                _ = v10
            } else {
                var v11 uint = 53
                v12 := GenS1{F0: v11, F1: 1, F2: (v11 * (v11 << ((18446744073709551615) & 63)))}
                _ = v11
                _ = v12
            }
        }
    }
    v13 := []int{9223372036854775807}
    var v14 int16 = p0
    if true {
        _ = len(v13)
    }
    _ = p0
    _ = v13
    _ = v14
    return 2147483647
}

func main() {
    var v1 uint64 = 107
    v2 := make(map[uint32]uint32)
    var v3 int16 = -32768
    var v4 int16 = v3
    v5 := make(map[uint32]uint32)
    v15 := GenS1{F0: 4294967295, F1: -1, F2: 227}
    switch "foo bar" {
    case "":
        v16 := GenS2{F0: 1000}
        _ = v16
    case "\n\t":
        if (32767 <= 0) {
            v17 := []int{-2147483648, 9223372036854775807}
            v4 -= v3
            if false {
                v5[2147483648] = 241
                var v18 int8 = -1
                _ = v18
            } else {
                for fi19 := 0; fi19 < 7; fi19++ {
                    if ((v1 % (((v1 - (v1 / (((v1 ^ v1)) | 1)))) | 1)) <= v1) {
                        v20 := GenS3{F0: 4294967295, F1: -128}
                        _ = v20
                    }
                    _ = fi19
                }
            }
            _ = v17
        }
    }
    fmt.Println(v4, v4)
    var v21 int64 = 42
    var v22 float32 = 0.1
    v23 := GenS3{F0: 42, F1: 1}
    v1 = 42
    v24 := GenS1{F0: uint(v1), F1: v21, F2: 65535}
    v25 := GenS4{F0: -1, F1: -74}
    v21 = v21
    v26 := GenS5{F0: "a"}
    for fi27 := 0; fi27 < 11; fi27++ {
        fmt.Println(v22)
        _ = fi27
    }
    v5[603] = 2147483648
    fmt.Println(v4, v22)
    for fi28 := 0; fi28 < 11; fi28++ {
        var v29 string = ("" + "")
        var v30 float32 = 0.1
        switch (fi28 * fi28) {
        case 32767:
            v5[895] = 496
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
