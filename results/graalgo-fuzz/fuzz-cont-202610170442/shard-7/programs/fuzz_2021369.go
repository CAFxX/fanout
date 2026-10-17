package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int64; F1 bool; F2 int64 }
type GenS3 struct { F0 int; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint8; F1 uint16; F2 uint }
func genFn0(p0 uint64) bool {
    if ((18446744073709551615 != 255) || false) {
        p0 -= p0
        var v5 int8 = -1
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 int64 = 9223372036854775807
            v8 := GenS1{F0: 0, F1: p0}
            switch 128 {
            case -9223372036854775808:
                v7 -= int64(v5)
            case 9223372036854775807:
                v9 := []int{128, fi6, fi6}
                _ = v9
            case 1:
                fmt.Println(v7)
            default:
                if (!false) {
                    v5 += v5
                    v5 = 1
                    v10 := []int{fi6}
                    _ = v10
                }
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = v5
    } else {
        v11 := []int{42, 127, 1, 275}
        _ = v11
    }
    p0 -= p0
    _ = p0
    return ("\n\t" != (("" + ("hello" + "foo bar")) + "a"))
}
func genFn1() string {
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS1{F0: 2, F1: 18446744073709551615}
        var v14 float32 = 3.14159
        v15 := GenS1{F0: 974, F1: 18446744073709551615}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 int64 = 666
    v17 := []int{28, 396}
    _ = v16
    _ = v17
    return (("" + "foo bar") + "foo bar")
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 231
    v3 := []uint32{175, 2147483647, 526, 1}
    v4 := []string{("a" + (("hello" + ("\n\t" + (("hello" + "hello") + "foo bar"))) + ""))}
    if (0.0 >= 0.1) {
        var v18 int = -9223372036854775808
        if len(v4) > 0 {
            v4[0] = ""
        }
        _ = v18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        if true {
            for fi20 := 0; fi20 < 10; fi20++ {
                if true {
                    var v21 float32 = float32(0.0)
                    _ = v21
                }
                var v22 uint64 = 563
                delete(v1, (fi20 | (fi20 & fi20)))
                _ = fi20
                _ = v22
            }
        }
        if (v2 == (v2 * 32767)) {
            v23 := GenS2{F0: 1, F1: true, F2: 1}
            delete(v1, -536)
            _ = v23
        } else {
            switch 9223372036854775807 {
            case 127:
                v24 := GenS2{F0: 9223372036854775807, F1: false, F2: 0}
                _ = v24
            case -9223372036854775808:
                var v25 int = fi19
                _ = v25
            case 9223372036854775807:
                v1[9223372036854775807] = fi19
            default:
                fmt.Println(v2, v2)
            }
            v2 = (v2 - v2)
        }
        _ = fi19
    }
    if true {
        v26 := GenS3{F0: 0, F1: 289, F2: 0}
        _ = v26
    } else {
        v2 -= (v2 % ((42) | 1))
    }
    if true {
        var v27 float64 = float64(3.14159)
        _ = v27
    }
    v28 := GenS4{F0: 124, F1: 2, F2: 646}
    fmt.Println(v2)
    switch -1 {
    case 188:
        switch "hello" {
        case "\n\t":
            v29 := GenS1{F0: 1, F1: 18446744073709551615}
            _ = v29
        case "foo bar":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            var v30 int = 128
            _ = v30
        }
    }
    v31 := GenS2{F0: -1, F1: (true || false), F2: 1}
    for fi32 := 0; fi32 < 4; fi32++ {
        for fi33 := 0; fi33 < 8; fi33++ {
            fmt.Println(fi33, fi32, fi32)
            _ = fi33
        }
        var v34 int8 = -128
        _ = fi32
        _ = v34
    }
    var v35 int16 = v2
    var v36 int32 = 546
    if (880 >= 65535) {
        for fi37 := 0; fi37 < 13; fi37++ {
            var v38 int16 = (v35 % (((-(v2))) | 1))
            _ = fi37
            _ = v38
        }
        var v39 float64 = ((((float64(((-2.25 + 0.0) / 3.14159)) - (0.0 / float64(3.14159))) - (((0.1 - (3.14159 / -2.25)) - -2.25) / 1.5)) + ((0.1 / float64((0.0 - (0.1 * (-2.25 / 1.5))))) + -2.25)) - (float64(1e10) - float64((((-2.25 / 1e10) / 1.0) + 3.14159))))
        _ = v39
    } else {
        var v40 uint16 = 1
        var v41 bool = (65535 < 2)
        _ = v40
        _ = v41
    }
    fmt.Println(v35)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v4
}
