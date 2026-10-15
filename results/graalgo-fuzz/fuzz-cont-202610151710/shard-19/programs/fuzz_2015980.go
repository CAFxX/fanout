package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int64; F1 bool; F2 int8 }

func main() {
    var v1 int8 = 1
    var v2 int64 = 0
    var v3 uint16 = 42
    var v4 int64 = 42
    var v5 uint64 = 42
    if false {
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := []int{}
            if true {
                if true {
                    for fi8 := 0; fi8 < 5; fi8++ {
                        switch ("foo bar" + "foo bar") {
                        case "\n\t":
                            v9 := GenS1{F0: v5, F1: -197, F2: 2}
                            _ = v9
                        case "hello":
                            var v10 uint32 = 2147483647
                            _ = v10
                        case "a":
                            v11 := GenS1{F0: v5, F1: 32767, F2: 1}
                            _ = v11
                        default:
                            var v12 bool = (0 > 339)
                            _ = v12
                        }
                        v1 = -128
                        _ = fi8
                    }
                    v3 = v3
                    for fi13 := 0; fi13 < 1; fi13++ {
                        v3 = (v3 ^ v3)
                        v14 := GenS1{F0: v5, F1: 634, F2: 1}
                        _ = fi13
                        _ = v14
                    }
                } else {
                    if true {
                        v1 -= (v1 * v1)
                    } else {
                        if len(v7) > 0 {
                            v7[0] = fi6
                        }
                        var v15 uint16 = (v3 | (v3 | (v3 - 2)))
                        _ = v15
                    }
                }
                var v16 uint8 = 0
                if (false || true) {
                    if len(v7) > 0 {
                        v7[0] = (fi6 * (fi6 & (fi6 ^ 1)))
                    }
                } else {
                    var v17 int32 = -1
                    _ = v17
                }
                _ = v16
            }
            _ = fi6
            _ = v7
        }
        fmt.Println(v2, v5)
    }
    v18 := []int{}
    var v19 uint64 = 18446744073709551615
    v18 = append(v18, 42)
    v3 = (v3 | (v3 / ((v3) | 1)))
    for fi20 := 0; fi20 < 18; fi20++ {
        if (false || true) {
            v5 = v19
        }
        _ = fi20
    }
    var v21 uint32 = 555
    v22 := GenS2{F0: v3}
    var v23 int64 = v2
    v24 := GenS1{F0: 697, F1: 1000, F2: 42}
    fmt.Println(v21, v3, v1)
    for fi25 := 0; fi25 < 13; fi25++ {
        v4 = v2
        v23 = v23
        v26 := GenS1{F0: v5, F1: 42, F2: uint8(v5)}
        _ = fi25
        _ = v26
    }
    switch 32768 {
    case 32767:
        v27 := GenS3{F0: 2147483647, F1: (v4 << ((255) & 63))}
        _ = v27
    case 9223372036854775807:
        v28 := GenS4{F0: v2, F1: (710 == 1), F2: v1}
        _ = v28
    case -9223372036854775808:
        var v29 uint64 = 42
        _ = v29
    default:
        var v30 string = ("a" + ("foo bar" + "foo bar"))
        _ = v30
    }
    fmt.Println(v2, v19, v5)
    v31 := GenS2{F0: (v3 & (v3 ^ v3))}
    var v32 uint8 = 84
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
