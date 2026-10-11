package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint64; F2 int16 }
type GenS4 struct { F0 int16; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int16; F1 uint64; F2 int32 }
func genFn0() uint {
    v6 := GenS1{F0: 0}
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS2{F0: 42}
        _ = fi7
        _ = v8
    }
    _ = v6
    return 278
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[uint32]bool)
    var v3 float64 = 0.0
    v4 := make(map[uint32]uint32)
    v5 := []int{0}
    var v9 int = 9223372036854775807
    for fi10 := 0; fi10 < 2; fi10++ {
        switch "a" {
        case "a":
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 int = 0
                fmt.Println(fi11)
                _ = fi11
                _ = v12
            }
        case "\n\t":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        default:
            v9 -= (^(fi10))
        }
        v13 := GenS1{F0: 2}
        _ = fi10
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        if (!true) {
            fmt.Println(v9, v1, v9)
        } else {
            fmt.Println(v9, fi14, fi14)
            v1 = 0
            if ((!true) && (156 != 0)) {
                for fi15 := 0; fi15 < 10; fi15++ {
                    v4[42] = 2147483648
                    _ = fi15
                }
                if true {
                    if (-1 == -9223372036854775808) {
                        var v16 int8 = 0
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 6; fi17++ {
                            var v18 uint16 = 0
                            v19 := GenS2{F0: 9223372036854775807}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    }
                    v5 = append(v5, v9)
                }
                var v20 uint = 834
                _ = v20
            }
        }
        if (false || false) {
            v21 := GenS2{F0: -1}
            v22 := GenS1{F0: 32767}
            var v23 float64 = float64(0.0)
            _ = v21
            _ = v22
            _ = v23
        }
        _ = fi14
    }
    if (42 < 598) {
        if (v9 > -944) {
            if ((1.5 < -2.25) || ((776 != -447) && (((v9 << ((355) & 63)) < (v9 >> ((18446744073709551615) & 63))) || false))) {
                var v24 uint64 = 18446744073709551615
                v3 = (v3 - 1e10)
                v1 = (v1 | 127)
                _ = v24
            } else {
                var v25 int32 = 42
                var v26 bool = false
                _ = v25
                _ = v26
            }
            if (false || true) {
                for fi27 := 0; fi27 < 18; fi27++ {
                    var v28 uint16 = 1
                    _ = len(v4)
                    _ = fi27
                    _ = v28
                }
            }
        } else {
            switch "foo bar" {
            case "hello":
                v29 := GenS2{F0: 9223372036854775807}
                _ = v29
            case "a":
                if (42 >= 9223372036854775807) {
                    if (-88 >= -2147483648) {
                        for fi30 := 0; fi30 < 11; fi30++ {
                            v9 += v9
                            _ = fi30
                        }
                        if (!false) {
                            var v31 float64 = (3.14159 + float64(1.5))
                            _ = v31
                        } else {
                            for fi32 := 0; fi32 < 6; fi32++ {
                                v3 = 0.0
                                _ = fi32
                            }
                            for fi33 := 0; fi33 < 3; fi33++ {
                                if (!(4294967295 != genFn0())) {
                                    v34 := GenS1{F0: 32767}
                                    v35 := GenS2{F0: 244}
                                    _ = v34
                                    _ = v35
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = fi33
                            }
                            fmt.Println(v1)
                        }
                        switch "foo bar" {
                        case "\n\t":
                            switch "\n\t" {
                            case "\n\t":
                                v3 += 1.5
                            case "hello":
                                var v36 int16 = 944
                                _ = v36
                            default:
                                v37 := GenS3{F0: 2, F1: 0, F2: -642}
                                _ = v37
                            }
                        case "hello":
                            if true {
                                fmt.Println(v9)
                            } else {
                                var v38 uint32 = 2
                                var v39 int8 = 0
                                if (v3 != v3) {
                                    var v40 int8 = v39
                                    _ = v40
                                }
                                _ = v38
                                _ = v39
                            }
                        case "foo bar":
                            v3 = v3
                        default:
                            v9 = (-(v9))
                        }
                    } else {
                        fmt.Println(v1)
                        fmt.Println(v9, v3)
                    }
                } else {
                    var v41 string = "foo bar"
                    _ = v41
                }
            case "":
                v3 -= 3.14159
            default:
                v9 = 9223372036854775807
            }
            v1 += (v1 & v1)
            v42 := GenS2{F0: 9223372036854775807}
            _ = v42
        }
    } else {
        if true {
            v3 = 1e10
            var v43 uint16 = 991
            _ = v43
        } else {
            for fi44 := 0; fi44 < 16; fi44++ {
                fmt.Println(v3)
                switch "a" {
                case "":
                    v2[50] = true
                case "a":
                    v45 := GenS1{F0: 382}
                    _ = v45
                case "\n\t":
                    fmt.Println(fi44, fi44, v1)
                default:
                    v3 = 0.1
                }
                if false {
                    for fi46 := 0; fi46 < 12; fi46++ {
                        if ("a" < "\n\t") {
                            fmt.Println(fi44, v1, fi44)
                        }
                        _ = fi46
                    }
                    var v47 uint32 = 547
                    _ = v47
                }
                _ = fi44
            }
            for fi48 := 0; fi48 < 10; fi48++ {
                fmt.Println(v9)
                _ = fi48
            }
            if false {
                v3 -= v3
                var v49 float32 = (-2.25 + -2.25)
                _ = v49
            } else {
                for fi50 := 0; fi50 < 5; fi50++ {
                    v51 := GenS4{F0: 1000, F1: true, F2: 2147483648}
                    v3 = v3
                    _ = fi50
                    _ = v51
                }
                if false {
                    var v52 int64 = -9223372036854775808
                    for fi53 := 0; fi53 < 3; fi53++ {
                        v54 := GenS1{F0: 2}
                        v52 += 198
                        v1 -= 1
                        _ = fi53
                        _ = v54
                    }
                    _ = v52
                } else {
                    switch "foo bar" {
                    case "":
                        if false {
                            v55 := GenS1{F0: 671}
                            for fi56 := 0; fi56 < 5; fi56++ {
                                if true {
                                    if (!false) {
                                        if (2 > 2) {
                                            v1 += 2
                                        }
                                        var v57 int16 = 983
                                        var v58 uint = 65535
                                        _ = v57
                                        _ = v58
                                    } else {
                                        if (!("\n\t" <= "")) {
                                            if len(v5) > 0 {
                                                v5[0] = 661
                                            }
                                        } else {
                                            var v59 float32 = 1e10
                                            _ = v59
                                        }
                                        fmt.Println(v3)
                                    }
                                } else {
                                    switch ((("\n\t" + "\n\t") + "hello") + "") {
                                    case "foo bar":
                                        v9 -= -2147483648
                                    }
                                }
                                if false {
                                    var v60 bool = (4294967295 == 18446744073709551615)
                                    _ = v60
                                } else {
                                    var v61 bool = (((false || ((false && true) && false)) && true) && (!(0 > 754)))
                                    fmt.Println(v61, v3)
                                    v1 += (v1 + 128)
                                    _ = v61
                                }
                                switch ((("\n\t" + ("a" + "hello")) + "a") + "hello") {
                                case "foo bar":
                                    if (993 >= 387) {
                                        if false {
                                            switch "\n\t" {
                                            case "hello":
                                                var v62 uint16 = 65535
                                                _ = v62
                                            }
                                            switch v9 {
                                            case -963:
                                                v63 := GenS5{F0: 227, F1: 127}
                                                _ = v63
                                            }
                                            if (0 > 32767) {
                                                _ = len(v2)
                                                v64 := GenS1{F0: 184}
                                                var v65 bool = true
                                                _ = v64
                                                _ = v65
                                            } else {
                                                switch (v9 | (fi56 & v9)) {
                                                case 1:
                                                    for fi66 := 0; fi66 < 16; fi66++ {
                                                        fmt.Println(fi56, fi66, v9)
                                                        _ = fi66
                                                    }
                                                case -2147483648:
                                                    for fi67 := 0; fi67 < 2; fi67++ {
                                                        fmt.Println(fi56, fi67, v9)
                                                        var v68 uint32 = 2147483647
                                                        _ = fi67
                                                        _ = v68
                                                    }
                                                }
                                            }
                                        } else {
                                            fmt.Println(v3, v9)
                                            v1 -= (^(v1))
                                            fmt.Println(v9)
                                        }
                                        switch ("foo bar" + ("hello" + "hello")) {
                                        case "\n\t":
                                            v69 := GenS1{F0: 540}
                                            _ = v69
                                        case "a":
                                            v9 = (^(v9))
                                        default:
                                            v70 := GenS6{F0: 823}
                                            _ = v70
                                        }
                                    } else {
                                        for fi71 := 0; fi71 < 14; fi71++ {
                                            if true {
                                                _ = len(v5)
                                            } else {
                                                v1 = 42
                                                var v72 uint16 = 65535
                                                _ = v72
                                            }
                                            var v73 float64 = (float64((((3.14159 / float32((v3 * (v3 + 1e10)))) * (3.14159 + 1.5)) - 1e10)) * v3)
                                            v74 := GenS7{F0: -53}
                                            _ = fi71
                                            _ = v73
                                            _ = v74
                                        }
                                        for fi75 := 0; fi75 < 6; fi75++ {
                                            switch "" {
                                            case "a":
                                                v1 += 196
                                            default:
                                                v76 := GenS3{F0: 65535, F1: 815, F2: 42}
                                                _ = v76
                                            }
                                            var v77 uint8 = v1
                                            v4[42] = 937
                                            _ = fi75
                                            _ = v77
                                        }
                                        v4[2147483648] = 42
                                    }
                                case "":
                                    var v78 string = "\n\t"
                                    _ = v78
                                case "a":
                                    v3 = v3
                                }
                                _ = fi56
                            }
                            _ = v55
                        } else {
                            var v79 float32 = (((0.0 * (-2.25 - (-2.25 * 0.0))) * 0.0) + (1.5 + 0.0))
                            v80 := GenS8{F0: (v1 >= (^(v1)))}
                            _ = v79
                            _ = v80
                        }
                    case "hello":
                        if (!true) {
                            v9 = (v9 << ((42) & 63))
                        } else {
                            switch v9 {
                            case 32767:
                                v81 := GenS7{F0: 127}
                                _ = v81
                            }
                            var v82 uint = 473
                            _ = v82
                        }
                    default:
                        _ = len(v4)
                    }
                    v83 := GenS6{F0: 759}
                    _ = v83
                }
                v84 := GenS5{F0: 65535, F1: -128}
                _ = v84
            }
        }
        var v85 bool = false
        _ = v85
    }
    var v86 int8 = 42
    v3 -= (((v3 + v3) + -2.25) / ((1e10 / v3) * v3))
    for fi87 := 0; fi87 < 4; fi87++ {
        var v88 string = "hello"
        switch 0 {
        case -1:
            if (-2147483648 == -2147483648) {
                var v89 string = "\n\t"
                _ = v89
            }
        case 914:
            v88 = (("\n\t" + "") + (("hello" + "a") + "hello"))
        }
        for fi90 := 0; fi90 < 2; fi90++ {
            _ = len(v2)
            _ = len(v4)
            v91 := GenS5{F0: 255, F1: v86}
            _ = fi90
            _ = v91
        }
        _ = fi87
        _ = v88
    }
    if true {
        v92 := GenS1{F0: 32767}
        _ = v92
    } else {
        if true {
            var v93 int16 = 42
            _ = v93
        } else {
            v94 := GenS6{F0: 560}
            _ = v94
        }
        delete(v2, 2147483648)
    }
    if true {
        if false {
            v95 := GenS9{F0: 55}
            var v96 int32 = 1
            _ = v95
            _ = v96
        } else {
            for fi97 := 0; fi97 < 2; fi97++ {
                if false {
                    if true {
                        if (float64((0.0 - 0.0)) > 1e10) {
                            v3 = (-2.25 - 1e10)
                        }
                        if (false || (false && true)) {
                            if false {
                                var v98 int64 = -1
                                var v99 int32 = -541
                                for fi100 := 0; fi100 < 3; fi100++ {
                                    switch ("foo bar" + "\n\t") {
                                    case "":
                                        v98 = -1
                                    case "\n\t":
                                        fmt.Println(fi97, fi97, v86)
                                    }
                                    _ = fi100
                                }
                                _ = v98
                                _ = v99
                            } else {
                                var v101 int = (fi97 * -530)
                                switch (v9 >> ((998) & 63)) {
                                case -464:
                                    if (42 == 1) {
                                        if true {
                                            fmt.Println(v3, fi97, v1)
                                        }
                                        var v102 uint64 = 932
                                        var v103 int8 = (v86 * v86)
                                        _ = v102
                                        _ = v103
                                    }
                                default:
                                    var v104 int32 = -2147483648
                                    _ = v104
                                }
                                _ = v101
                            }
                        }
                        switch (v9 >> ((230) & 63)) {
                        case -473:
                            fmt.Println(v9)
                        case -1:
                            var v105 uint8 = (v1 - 1)
                            _ = v105
                        case 871:
                            v86 = -100
                        }
                    } else {
                        fmt.Println(v86, fi97)
                        v1 += 127
                        v106 := GenS7{F0: (-(v86))}
                        _ = v106
                    }
                }
                switch (fi97 * v9) {
                case 32767:
                    for fi107 := 0; fi107 < 12; fi107++ {
                        fmt.Println(v9, v86)
                        fmt.Println(v1, v1)
                        var v108 float64 = -2.25
                        _ = fi107
                        _ = v108
                    }
                }
                _ = fi97
            }
            fmt.Println(v1, v86)
            for fi109 := 0; fi109 < 7; fi109++ {
                v110 := GenS10{F0: -370, F1: 568, F2: 2147483647}
                if (!false) {
                    var v111 float64 = 3.14159
                    fmt.Println(v3, v86)
                    _ = len(v4)
                    _ = v111
                } else {
                    var v112 uint = 65535
                    _ = v112
                }
                for fi113 := 0; fi113 < 18; fi113++ {
                    var v114 int = (fi113 - fi113)
                    fmt.Println(v86, v3, v9)
                    _ = fi113
                    _ = v114
                }
                _ = fi109
                _ = v110
            }
        }
    } else {
        v9 -= (v9 ^ v9)
        v86 -= v86
        v1 = 255
    }
    v9 -= 2147483647
    if true {
        for fi115 := 0; fi115 < 16; fi115++ {
            if (v3 > (v3 + 1.5)) {
                switch ((("\n\t" + "hello") + ("hello" + "\n\t")) + "foo bar") {
                case "a":
                    v116 := GenS7{F0: (v86 * v86)}
                    _ = v116
                case "hello":
                    var v117 float32 = ((3.14159 - -2.25) * (0.1 - 0.1))
                    _ = v117
                }
                var v118 uint8 = (v1 + 127)
                v119 := GenS3{F0: 4294967295, F1: 4294967295, F2: 319}
                _ = v118
                _ = v119
            } else {
                v3 = v3
                var v120 int32 = -942
                _ = v120
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            delete(v4, 65535)
            _ = fi115
        }
        if true {
            var v121 int64 = -1
            v9 = v9
            v122 := GenS6{F0: 65535}
            _ = v121
            _ = v122
        }
    } else {
        var v123 uint32 = 852
        v124 := GenS8{F0: false}
        _ = v123
        _ = v124
    }
    for fi125 := 0; fi125 < 20; fi125++ {
        for fi126 := 0; fi126 < 20; fi126++ {
            fmt.Println(fi126, v9)
            _ = fi126
        }
        _ = fi125
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v86:", v86)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v86
    _ = v9
}
