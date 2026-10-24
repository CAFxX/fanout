package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int16) uint32 {
    v4 := GenS1{F0: 92}
    fmt.Println(p0)
    if true {
        p0 -= p0
    } else {
        var v5 uint64 = 4294967295
        if (!true) {
            if (p0 != p0) {
                var v6 uint = 42
                v7 := []int{}
                _ = v6
                _ = v7
            } else {
                for fi8 := 0; fi8 < 13; fi8++ {
                    if true {
                        for fi9 := 0; fi9 < 7; fi9++ {
                            var v10 bool = (((!true) || false) || true)
                            _ = fi9
                            _ = v10
                        }
                    } else {
                        v11 := GenS2{F0: true, F1: (p0 + p0)}
                        v12 := GenS3{F0: 32767, F1: (v5 << ((327) & 63))}
                        _ = v11
                        _ = v12
                    }
                    _ = fi8
                }
                switch 128 {
                case 2147483647:
                    for fi13 := 0; fi13 < 15; fi13++ {
                        v14 := GenS2{F0: ((true || (2147483647 < 4294967295)) && (false && true)), F1: 23}
                        _ = fi13
                        _ = v14
                    }
                case 0:
                    var v15 int64 = 801
                    _ = v15
                default:
                    v16 := GenS4{F0: 0}
                    _ = v16
                }
            }
        }
        _ = v5
    }
    _ = p0
    _ = v4
    return 1
}
func genFn1(p0 uint64, p1 bool, p2 uint64) int {
    switch 270 {
    case 128:
        var v17 int = -1
        _ = v17
    case 9223372036854775807:
        p1 = (-128 >= 1)
    case 4294967295:
        var v18 int8 = 42
        _ = v18
    }
    p0 += (p0 << ((2) & 63))
    v19 := GenS3{F0: 32767, F1: 42}
    var v20 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return 32768
}

func main() {
    var v1 uint = 0
    var v2 uint32 = 42
    var v3 float64 = ((-2.25 - float64(0.0)) / (float64(3.14159) * 1e10))
    var v21 float64 = 1.5
    switch -2147483648 {
    case -350:
        var v22 float64 = 0.0
        _ = v22
    default:
        v3 += float64(((-2.25 * 1e10) / -2.25))
    }
    v21 = ((1.5 / v21) / v3)
    v2 = uint32(v1)
    if true {
        v23 := GenS5{F0: 128}
        _ = v23
    } else {
        v24 := []int{-149, 128, -1}
        for fi25 := 0; fi25 < 9; fi25++ {
            switch ("\n\t" + ("" + ("foo bar" + ("foo bar" + "foo bar")))) {
            case "a":
                fmt.Println(v1)
            case "\n\t":
                v1 += v1
            case "hello":
                fmt.Println(v3, v21, v1)
            }
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
            _ = fi25
        }
        v21 = v21
        _ = v24
    }
    v26 := GenS6{F0: genFn1(4294967295, false, 4294967295)}
    v27 := []int{4294967295, -9223372036854775808, -1, 446}
    v28 := GenS2{F0: false, F1: 383}
    switch 1 {
    case -2147483648:
        if (0 >= -9223372036854775808) {
            v29 := GenS7{F0: 2, F1: 65535}
            _ = v29
        } else {
            var v30 float32 = (1.5 / (float32((v3 / (float64(3.14159) / 1.0))) * 1.5))
            _ = v30
        }
    case -9223372036854775808:
        if (!(-788 >= 843)) {
            v2 += v2
        } else {
            v21 += v3
        }
    case 1:
        var v31 uint64 = 1
        _ = v31
    default:
        v32 := GenS1{F0: 42}
        _ = v32
    }
    fmt.Println(v1, v2, v1)
    for fi33 := 0; fi33 < 1; fi33++ {
        fmt.Println(fi33, v2, v21)
        v3 += -2.25
        switch ("foo bar" + "foo bar") {
        case "a":
            v3 = (v21 - (0.1 * v3))
        default:
            v3 = v21
        }
        _ = fi33
    }
    var v34 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v34
}
