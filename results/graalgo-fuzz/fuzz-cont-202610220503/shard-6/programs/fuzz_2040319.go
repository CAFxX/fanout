package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS6 struct { F0 int32; F1 int; F2 bool }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 bool; F1 uint64; F2 int32 }
func genFn0(p0 int32) uint8 {
    p0 = (p0 / ((p0) | 1))
    v6 := []int{-517, -1}
    if (("" + "") != "foo bar") {
        if false {
            switch "foo bar" {
            case "":
                p0 = (^(p0))
            case "foo bar":
                v7 := GenS1{F0: 18446744073709551615, F1: 2}
                _ = v7
            case "a":
                switch -2147483648 {
                case 1:
                    fmt.Println(p0)
                case -459:
                    v8 := GenS2{F0: -128}
                    _ = v8
                }
            default:
                fmt.Println(p0)
            }
            var v9 int32 = (p0 * p0)
            switch "\n\t" {
            case "":
                if len(v6) > 0 {
                    v6[0] = 2147483647
                }
            case "a":
                var v10 uint16 = 451
                _ = v10
            case "hello":
                fmt.Println(p0)
            default:
                switch 2147483647 {
                case 42:
                    switch -1 {
                    case -2147483648:
                        switch 32768 {
                        case 128:
                            p0 -= v9
                        case 121:
                            var v11 int64 = -9223372036854775808
                            _ = v11
                        case -425:
                            switch ("foo bar" + ("foo bar" + "foo bar")) {
                            case "hello":
                                p0 = (^(v9))
                            }
                        default:
                            p0 += p0
                        }
                    case 0:
                        fmt.Println(p0, v9)
                    default:
                        for fi12 := 0; fi12 < 4; fi12++ {
                            var v13 uint32 = 0
                            var v14 uint8 = 127
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                    }
                }
            }
            _ = v9
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v15 uint32 = 2147483647
            v15 -= 2
            _ = v15
        }
    }
    _ = p0
    _ = v6
    return 2
}
func genFn1(p0 int16, p1 int32, p2 int32) int {
    if ((42 >= -1) && (42 > 1)) {
        for fi16 := 0; fi16 < 6; fi16++ {
            v17 := GenS1{F0: 0, F1: 88}
            v18 := []int{}
            if len(v18) > 0 {
                v18[0] = fi16
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        if ((3.14159 <= 1.5) && (0 == 128)) {
            switch int(p0) {
            case 2147483647:
                if true {
                    var v19 float32 = 0.1
                    var v20 int16 = p0
                    _ = v19
                    _ = v20
                }
            case -745:
                switch "" {
                case "foo bar":
                    fmt.Println(p2, p0)
                case "a":
                    v21 := GenS2{F0: -128}
                    _ = v21
                }
            case -2147483648:
                var v22 uint32 = 196
                _ = v22
            default:
                v23 := []int{-434, -265, 127, -1}
                _ = v23
            }
        }
    }
    switch 128 {
    case 618:
        for fi24 := 0; fi24 < 19; fi24++ {
            v25 := GenS3{F0: 18446744073709551615, F1: 2147483647, F2: 0}
            switch "a" {
            case "hello":
                if (1000 != 1000) {
                    var v26 int32 = p2
                    _ = v26
                } else {
                    var v27 string = "foo bar"
                    _ = v27
                }
            case "":
                p2 += p1
            case "\n\t":
                if false {
                    for fi28 := 0; fi28 < 1; fi28++ {
                        switch "a" {
                        case "\n\t":
                            switch (("a" + "foo bar") + "foo bar") {
                            case "foo bar":
                                p2 = (p2 / (((p2 - (^(p1)))) | 1))
                            case "hello":
                                fmt.Println(fi24, fi28, fi28)
                                v29 := []int{9223372036854775807, fi28}
                                p2 -= p1
                                _ = v29
                            case "a":
                                v30 := []int{(fi24 >> ((18446744073709551615) & 63))}
                                var v31 uint8 = 127
                                _ = v30
                                _ = v31
                            }
                        }
                        if true {
                            fmt.Println(p1)
                            p1 = 308
                            p1 = 1
                        } else {
                            var v32 int16 = 0
                            p1 += (p2 << ((65535) & 63))
                            _ = v32
                        }
                        var v33 int64 = -1
                        _ = fi28
                        _ = v33
                    }
                }
            default:
                var v34 string = ((("hello" + "foo bar") + (("foo bar" + "foo bar") + "hello")) + ("" + "a"))
                _ = v34
            }
            _ = fi24
            _ = v25
        }
    }
    p0 = 42
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint16 = 0
    v2 := []int{0}
    var v3 int = -2147483648
    var v4 uint8 = 0
    var v5 int8 = 0
    for fi35 := 0; fi35 < 17; fi35++ {
        switch int(v5) {
        case -713:
            if false {
                fmt.Println(v5)
                var v36 float32 = 1.5
                for fi37 := 0; fi37 < 18; fi37++ {
                    var v38 uint8 = (v4 / ((genFn0(-1)) | 1))
                    var v39 uint8 = v38
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                _ = v36
            }
        case -2147483648:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        if false {
            if (2 < 0) {
                for fi40 := 0; fi40 < 20; fi40++ {
                    var v41 int16 = -443
                    v42 := GenS2{F0: (v5 >> ((18446744073709551615) & 63))}
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                v3 = -1
            }
            if true {
                fmt.Println(v1)
                for fi43 := 0; fi43 < 2; fi43++ {
                    v1 -= 255
                    v44 := GenS1{F0: 4294967295, F1: (v1 - v1)}
                    v45 := GenS3{F0: 42, F1: 2147483647, F2: 73}
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                fmt.Println(v1)
                v1 -= 728
                var v46 uint8 = v4
                _ = v46
            }
        } else {
            if false {
                var v47 string = "\n\t"
                for fi48 := 0; fi48 < 10; fi48++ {
                    v49 := GenS1{F0: uint64(v4), F1: (v1 & 178)}
                    v50 := GenS3{F0: 18446744073709551615, F1: 858, F2: 42}
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                _ = v47
            }
        }
        _ = fi35
    }
    var v51 int = v3
    v1 = 158
    for fi52 := 0; fi52 < 7; fi52++ {
        var v53 int64 = -823
        v51 = -807
        v54 := GenS1{F0: 462, F1: v1}
        _ = fi52
        _ = v53
        _ = v54
    }
    if (!(!true)) {
        var v55 string = "hello"
        switch v55 {
        case "foo bar":
            v51 = genFn1(1, 914, -705)
        default:
            fmt.Println(v5, v55, v5)
        }
        var v56 uint64 = 1
        _ = v55
        _ = v56
    }
    if false {
        for fi57 := 0; fi57 < 9; fi57++ {
            v51 += v51
            _ = fi57
        }
        if len(v2) > 0 {
            v2[0] = 32767
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v51 += v51
        if (1e10 == 1e10) {
            switch ("hello" + "foo bar") {
            case "foo bar":
                var v58 float64 = float64(0.1)
                _ = v58
            case "a":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "hello":
                var v59 uint64 = 4294967295
                _ = v59
            default:
                for fi60 := 0; fi60 < 12; fi60++ {
                    fmt.Println(v3)
                    var v61 string = "a"
                    v3 = v3
                    _ = fi60
                    _ = v61
                }
            }
        } else {
            if (((0.0 / -2.25) >= (0.0 + 3.14159)) || false) {
                switch ("foo bar" + (((("foo bar" + ("hello" + "\n\t")) + (("" + "") + ("hello" + ("hello" + "a")))) + "a") + ("" + ""))) {
                case "a":
                    v4 += v4
                }
            } else {
                var v62 int = (v51 | (v3 & -238))
                for fi63 := 0; fi63 < 20; fi63++ {
                    if ((9223372036854775807 != int64(v3)) && true) {
                        v64 := GenS2{F0: v5}
                        if true {
                            switch ("\n\t" + ("\n\t" + "")) {
                            case "a":
                                if true {
                                    var v65 int64 = 9223372036854775807
                                    _ = v65
                                }
                            case "\n\t":
                                var v66 uint8 = (v4 / (((v4 - 183)) | 1))
                                _ = v66
                            case "hello":
                                fmt.Println(v5, fi63, v62)
                            default:
                                fmt.Println(v3, v51)
                            }
                            v67 := GenS4{F0: 18446744073709551615, F1: 0}
                            _ = v67
                        } else {
                            for fi68 := 0; fi68 < 15; fi68++ {
                                switch "\n\t" {
                                case "a":
                                    v69 := GenS5{F0: (^(v4)), F1: v5, F2: v1}
                                    _ = v69
                                case "hello":
                                    switch 196 {
                                    case 2147483647:
                                        switch "a" {
                                        case "hello":
                                            _ = len(v2)
                                        case "a":
                                            switch v3 {
                                            case 42:
                                                switch 32767 {
                                                case 4294967295:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case 32767:
                                                    _ = len(v2)
                                                default:
                                                    v51 -= v51
                                                }
                                            case -557:
                                                var v70 uint32 = 42
                                                _ = v70
                                            default:
                                                v62 = (v62 % ((genFn1(972, 1, 42)) | 1))
                                            }
                                            if true {
                                                var v71 int32 = -2147483648
                                                _ = v71
                                            } else {
                                                var v72 float32 = (0.1 * 0.1)
                                                v1 -= uint16(v4)
                                                _ = v72
                                            }
                                            v1 = 0
                                        case "foo bar":
                                            if true {
                                                var v73 int16 = 1000
                                                v74 := GenS6{F0: 997, F1: v51, F2: true}
                                                _ = v73
                                                _ = v74
                                            }
                                        }
                                    case 896:
                                        v75 := GenS2{F0: (v5 & v5)}
                                        _ = v75
                                    default:
                                        v5 = v5
                                    }
                                    if (0 > 42) {
                                        fmt.Println(v5, v51)
                                    } else {
                                        if (((((3.14159 + -2.25) + float32(float64((float32((-2.25 / -2.25)) / (0.1 / 3.14159))))) + ((3.14159 * float32((-2.25 / 1.0))) - 0.0)) / -2.25) == -2.25) {
                                            v76 := GenS6{F0: 314, F1: v3, F2: true}
                                            if (668 < 42) {
                                                if true {
                                                    var v77 bool = ((!true) || true)
                                                    if (42 >= uint64(v4)) {
                                                        var v78 int16 = 42
                                                        var v79 uint16 = v1
                                                        _ = v78
                                                        _ = v79
                                                    } else {
                                                        switch "" {
                                                        case "":
                                                            var v80 int16 = 32767
                                                            var v81 bool = v77
                                                            if (false && (0.1 == float64(1e10))) {
                                                                v62 += fi68
                                                            }
                                                            _ = v80
                                                            _ = v81
                                                        case "foo bar":
                                                            switch "hello" {
                                                            case "foo bar":
                                                                if len(v2) > 0 {
                                                                    v2[0] = 668
                                                                }
                                                            default:
                                                                var v82 uint32 = 0
                                                                _ = v82
                                                            }
                                                            if (!true) {
                                                                var v83 float32 = -2.25
                                                                if v77 {
                                                                    var v84 float64 = -2.25
                                                                    _ = v84
                                                                }
                                                                _ = v83
                                                            }
                                                            var v85 uint64 = 376
                                                            _ = v85
                                                        }
                                                        if (214 != 255) {
                                                            var v86 uint16 = (v1 % ((v1) | 1))
                                                            _ = v86
                                                        } else {
                                                            v5 = v5
                                                            v4 -= v4
                                                        }
                                                    }
                                                    fmt.Println(v3)
                                                    _ = v77
                                                } else {
                                                    if false {
                                                        v5 += -128
                                                        v3 -= v51
                                                        switch fi68 {
                                                        case 225:
                                                            var v87 uint64 = 2
                                                            _ = v87
                                                        }
                                                    } else {
                                                        v2 = append(v2, (fi68 + 500))
                                                    }
                                                    if false {
                                                        v62 = genFn1(127, 1, 1)
                                                        fmt.Println(v62, v5, v1)
                                                        var v88 bool = (!(1000 == -32768))
                                                        _ = v88
                                                    } else {
                                                        var v89 uint8 = 255
                                                        _ = v89
                                                    }
                                                    var v90 uint32 = 2147483647
                                                    _ = v90
                                                }
                                                v4 += genFn0(42)
                                                var v91 float64 = ((3.14159 - (1.5 + -2.25)) + 3.14159)
                                                if len(v2) > 0 {
                                                    v2[0] = 9223372036854775807
                                                }
                                                fmt.Println(v3, v1, v4)
                                                fmt.Println(fi63, v3)
                                                _ = v91
                                            }
                                            _ = v76
                                        } else {
                                            if ((v5 >> ((65535) & 63)) == v5) {
                                                if true {
                                                    var v92 int32 = 1
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = v92
                                                }
                                                var v93 float64 = -2.25
                                                if len(v2) > 0 {
                                                    v2[0] = genFn1(32767, 767, 0)
                                                }
                                                _ = v93
                                            } else {
                                                v94 := GenS2{F0: v5}
                                                v4 = 2
                                                _ = v94
                                            }
                                        }
                                    }
                                    var v95 uint32 = uint32(v1)
                                    _ = v95
                                }
                                if (1 < 42) {
                                    _ = len(v2)
                                } else {
                                    if (155 == 42) {
                                        if false {
                                            fmt.Println(v5, v62)
                                        }
                                        v96 := GenS7{F0: "a", F1: v1}
                                        _ = v96
                                    }
                                }
                                _ = fi68
                            }
                            v97 := GenS1{F0: 81, F1: (v1 & (^(v1)))}
                            v5 = 0
                            _ = v97
                        }
                        _ = v64
                    } else {
                        for fi98 := 0; fi98 < 14; fi98++ {
                            v1 = (^(v1))
                            _ = fi98
                        }
                    }
                    _ = fi63
                }
                for fi99 := 0; fi99 < 16; fi99++ {
                    v100 := GenS4{F0: 2, F1: v51}
                    for fi101 := 0; fi101 < 9; fi101++ {
                        v5 = v5
                        switch ((("foo bar" + ("" + "foo bar")) + (("foo bar" + ("hello" + "\n\t")) + ((("" + "a") + "foo bar") + "foo bar"))) + "") {
                        case "a":
                            if len(v2) > 0 {
                                v2[0] = (-(fi101))
                            }
                        }
                        _ = fi101
                    }
                    _ = fi99
                    _ = v100
                }
                _ = v62
            }
        }
        v102 := GenS8{F0: true, F1: 875, F2: -2147483648}
        _ = v102
    }
    fmt.Println(v3, v1, v51)
    var v103 string = "a"
    fmt.Println(v1, v4)
    for fi104 := 0; fi104 < 17; fi104++ {
        var v105 string = (v103 + v103)
        var v106 float32 = (1.5 + 1.5)
        _ = fi104
        _ = v105
        _ = v106
    }
    fmt.Println(v3, v103, v1)
    var v107 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v107:", v107)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v103
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
