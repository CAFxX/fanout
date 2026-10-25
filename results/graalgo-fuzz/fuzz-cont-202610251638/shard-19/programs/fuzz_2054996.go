package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS4 struct { F0 int; F1 uint8; F2 int8 }
func genFn0(p0 int8, p1 bool, p2 int16) int32 {
    p1 = (p1 && p1)
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1(p0 uint, p1 uint8) uint {
    v6 := GenS1{F0: (0 > -32768), F1: 1, F2: 4294967295}
    p0 = 502
    p1 = (p1 + p1)
    fmt.Println(p1, p1)
    var v7 int16 = -311
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p0 * (p0 << (((p0 % ((p0) | 1))) & 63)))
}

func main() {
    v1 := []int{812, 1, 47}
    v2 := []int{9223372036854775807, 0}
    var v3 bool = (267 <= -948)
    var v4 int16 = -32768
    v5 := make(map[int]bool)
    v8 := GenS1{F0: v3, F1: -751, F2: 2147483647}
    var v9 int64 = 968
    if (-2.25 > 1e10) {
        fmt.Println(v4, v9)
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch -9223372036854775808 {
        case 600:
            v9 -= (v9 << ((4294967295) & 63))
        case -2147483648:
            v10 := GenS2{F0: (((((v9 >> ((508) & 63)) == v9) || (128 != 32767)) || (-1 < 157)) && (true || (2147483647 < 9223372036854775807)))}
            _ = v10
        default:
            for fi11 := 0; fi11 < 9; fi11++ {
                fmt.Println(v3)
                _ = fi11
            }
        }
        if true {
            if (!((true || false) && v3)) {
                v12 := GenS2{F0: false}
                _ = v12
            } else {
                v13 := GenS1{F0: (-2.25 >= 3.14159), F1: -1, F2: 2147483648}
                _ = v13
            }
        }
    }
    switch "a" {
    case "hello":
        switch -582 {
        case 9223372036854775807:
            v9 -= v9
        case 2147483647:
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 bool = ("foo bar" >= "\n\t")
                _ = fi14
                _ = v15
            }
        case 127:
            fmt.Println(v4, v9)
        default:
            v16 := GenS2{F0: v3}
            _ = v16
        }
    case "foo bar":
        for fi17 := 0; fi17 < 7; fi17++ {
            v18 := GenS1{F0: (!true), F1: -564, F2: 0}
            if len(v1) > 0 {
                v1[0] = fi17
            }
            _ = fi17
            _ = v18
        }
    case "\n\t":
        var v19 uint64 = 266
        _ = v19
    default:
        var v20 uint8 = 255
        _ = v20
    }
    var v21 int = -1
    v9 += v9
    fmt.Println(v4, v4, v21)
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 int8 = 127
        v3 = (v3 && (644 != 42))
        v3 = v3
        _ = fi22
        _ = v23
    }
    if false {
        v4 = 502
        v4 += v4
    } else {
        var v24 int8 = 116
        v21 = -2147483648
        if (((float32(float64((-2.25 + (0.1 + (3.14159 - 0.1))))) >= float32(3.14159)) && (!v3)) || false) {
            var v25 bool = (165 < 127)
            v4 -= (v4 ^ (-(v4)))
            _ = v25
        }
        _ = v24
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi26
    }
    switch -9223372036854775808 {
    case -2147483648:
        fmt.Println(v3)
    case 2147483647:
        v27 := GenS1{F0: false, F1: -372, F2: 1}
        _ = v27
    case 0:
        v28 := GenS3{F0: 32767, F1: 412, F2: -899}
        _ = v28
    default:
        v29 := GenS2{F0: v3}
        _ = v29
    }
    v30 := GenS4{F0: v21, F1: 0, F2: -66}
    if ((v4 - v4) == (v4 | (v4 + v4))) {
        for fi31 := 0; fi31 < 5; fi31++ {
            if v3 {
                var v32 string = ""
                _ = v32
            } else {
                v33 := v5[v21]
                _ = v33
            }
            _ = fi31
        }
        _ = len(v5)
        _ = len(v1)
    }
    v4 = v4
    var v34 uint = 649
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
