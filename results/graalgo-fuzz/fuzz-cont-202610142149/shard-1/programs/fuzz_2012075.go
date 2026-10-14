package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int8; F1 int16; F2 uint64 }
type GenS3 struct { F0 int; F1 uint16; F2 uint }
type GenS4 struct { F0 int64; F1 int8; F2 int64 }
type GenS5 struct { F0 int8; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint16; F1 int; F2 uint64 }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 string; F1 uint8 }
func genFn0(p0 uint32, p1 uint16, p2 uint16) int {
    v6 := GenS1{F0: 42, F1: 2}
    p0 -= (p0 / (((p0 - p0)) | 1))
    if false {
        var v7 uint8 = 88
        var v8 int64 = 0
        _ = v7
        _ = v8
    } else {
        var v9 uint8 = 127
        var v10 uint32 = 342
        var v11 uint16 = 255
        _ = v10
        _ = v11
        _ = v9
    }
    p1 = (p1 + p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1(p0 int32, p1 int64, p2 int) int8 {
    v12 := []int{-71}
    if false {
        for fi13 := 0; fi13 < 9; fi13++ {
            p0 -= (p0 >> ((65535) & 63))
            _ = fi13
        }
    } else {
        if ((false && false) || ((!(float32((float64(-2.25) + 1e10)) != (1.5 + (3.14159 * -2.25)))) && ((^(p1)) < (p1 * (p1 | 9223372036854775807))))) {
            v14 := GenS2{F0: 1, F1: 42, F2: 2}
            switch (-(p2)) {
            case -2147483648:
                p2 -= genFn0(100, 2, 1)
            }
            _ = v14
        }
        var v15 int64 = 32
        _ = v15
    }
    if (1 >= -32768) {
        p1 -= 9223372036854775807
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := GenS3{F0: 371, F1: 32767, F2: 255}
        v18 := GenS1{F0: 19, F1: 42}
        for fi19 := 0; fi19 < 4; fi19++ {
            v20 := GenS3{F0: p2, F1: 460, F2: 553}
            _ = len(v12)
            v21 := GenS3{F0: -518, F1: 42, F2: 2}
            if (65535 <= 42) {
                var v22 uint32 = 42
                v23 := GenS4{F0: (-(p1)), F1: -28, F2: p1}
                v12 = append(v12, p2)
                _ = v22
                _ = v23
            } else {
                var v24 int32 = p0
                var v25 int64 = (p1 + p1)
                _ = v24
                _ = v25
            }
            if len(v12) > 0 {
                v12[0] = (fi16 * p2)
            }
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 78
}

func main() {
    v1 := []uint32{2147483647, 337}
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    v5 := []int{}
    _ = len(v5)
    var v26 uint = 268
    switch "\n\t" {
    case "":
        v27 := GenS5{F0: 1, F1: 281, F2: 725}
        _ = v27
    case "hello":
        v28 := GenS5{F0: 1, F1: 343, F2: 4294967295}
        _ = v28
    default:
        var v29 float32 = 0.0
        _ = v29
    }
    fmt.Println(v26)
    for fi30 := 0; fi30 < 4; fi30++ {
        if len(v5) > 0 {
            v5[0] = (fi30 >> (((v26 % ((65535) | 1))) & 63))
        }
        var v31 int8 = 67
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        fmt.Println(v26, v26)
        _ = fi32
    }
    if (1 > genFn1(42, 64, 32768)) {
        v26 += v26
        switch 32767 {
        case -302:
            v33 := GenS6{F0: 42, F1: 42, F2: 1}
            _ = v33
        case 0:
            switch "a" {
            case "a":
                switch "a" {
                case "":
                    v34 := v2[910]
                    _ = v34
                case "\n\t":
                    if (false || (-2147483648 <= 445)) {
                        delete(v2, 2147483648)
                        if false {
                            fmt.Println(v26)
                            fmt.Println(v26)
                        } else {
                            v35 := GenS7{F0: -58, F1: -32768}
                            for fi36 := 0; fi36 < 20; fi36++ {
                                if (0.0 == -2.25) {
                                    switch ("foo bar" + ("\n\t" + "a")) {
                                    case "\n\t":
                                        v26 = (v26 % (((^(v26))) | 1))
                                    case "":
                                        if len(v1) > 0 {
                                            v1[0] = 324
                                        }
                                    default:
                                        if (2147483647 == 1) {
                                            var v37 int8 = 100
                                            _ = v37
                                        }
                                    }
                                    switch fi36 {
                                    case -9223372036854775808:
                                        switch (-(fi36)) {
                                        case -158:
                                            switch (fi36 >> ((1) & 63)) {
                                            case 42:
                                                switch ("\n\t" + ("a" + "foo bar")) {
                                                case "a":
                                                    v26 -= (^(v26))
                                                default:
                                                    var v38 string = ("a" + "")
                                                    _ = v38
                                                }
                                            }
                                        default:
                                            _ = len(v5)
                                        }
                                    case 32767:
                                        if true {
                                            for fi39 := 0; fi39 < 3; fi39++ {
                                                v26 = v26
                                                var v40 uint8 = 218
                                                _ = fi39
                                                _ = v40
                                            }
                                            var v41 string = ("hello" + "a")
                                            _ = v41
                                        } else {
                                            v26 = v26
                                            switch (("\n\t" + "foo bar") + (("" + ("a" + "foo bar")) + "")) {
                                            case "foo bar":
                                                fmt.Println(v26, v26)
                                            case "\n\t":
                                                fmt.Println(v26, v26)
                                            case "":
                                                fmt.Println(v26, fi36)
                                            default:
                                                v26 += v26
                                            }
                                        }
                                    case 4294967295:
                                        switch ("foo bar" + "") {
                                        case "\n\t":
                                            fmt.Println(v26, fi36)
                                        case "hello":
                                            v42 := GenS4{F0: -9223372036854775808, F1: 1, F2: 956}
                                            _ = v42
                                        default:
                                            if len(v5) > 0 {
                                                v5[0] = (fi36 << ((v26) & 63))
                                            }
                                        }
                                    }
                                }
                                _ = fi36
                            }
                            _ = v35
                        }
                        v26 -= (v26 % ((4294967295) | 1))
                    }
                default:
                    var v43 float32 = 1e10
                    _ = v43
                }
            case "":
                if false {
                    for fi44 := 0; fi44 < 2; fi44++ {
                        for fi45 := 0; fi45 < 16; fi45++ {
                            v26 = v26
                            _ = fi45
                        }
                        if ((((0.0 / 0.1) + 0.1) <= float64(float32((0.0 - 0.0)))) || (-1 >= genFn1(618, 9223372036854775807, fi44))) {
                            var v46 float32 = 0.1
                            _ = v46
                        }
                        var v47 string = "a"
                        _ = fi44
                        _ = v47
                    }
                    var v48 float32 = 3.14159
                    _ = v48
                }
            }
        }
    }
    var v49 uint16 = 617
    for fi50 := 0; fi50 < 9; fi50++ {
        switch fi50 {
        case -9223372036854775808:
            for fi51 := 0; fi51 < 9; fi51++ {
                switch fi51 {
                case -445:
                    v52 := GenS2{F0: -76, F1: 42, F2: 823}
                    _ = v52
                case -497:
                    if true {
                        if (true && true) {
                            v53 := GenS4{F0: 0, F1: genFn1(42, 1, 32767), F2: -1}
                            v49 = v49
                            v49 -= v49
                            _ = v53
                        }
                    }
                default:
                    fmt.Println(fi51, fi50, fi50)
                }
                if (570 > 4294967295) {
                    switch 873 {
                    case -218:
                        v3[42] = 0
                    default:
                        switch "hello" {
                        case "a":
                            fmt.Println(fi51, v26)
                            _ = len(v4)
                            fmt.Println(fi51, v49)
                        case "\n\t":
                            v5 = append(v5, (-(fi50)))
                        case "hello":
                            if len(v1) > 0 {
                                v1[0] = 989
                            }
                        }
                    }
                } else {
                    if (32767 <= v49) {
                        v26 = v26
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                delete(v2, 238)
                v26 += 277
                _ = fi51
            }
        }
        var v54 bool = (!(1 != 93))
        _ = fi50
        _ = v54
    }
    var v55 int32 = -1
    var v56 float32 = -2.25
    fmt.Println(v49, v55, v26)
    fmt.Println(v26)
    v57 := v2[194]
    v58 := GenS8{F0: "hello", F1: 82}
    var v59 int64 = 9223372036854775807
    v57 -= 2147483648
    v57 -= 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}
