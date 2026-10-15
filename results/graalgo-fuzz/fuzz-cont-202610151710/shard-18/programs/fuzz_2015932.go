package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 uint32; F2 int }
type GenS8 struct { F0 int64; F1 uint; F2 uint8 }
type GenS9 struct { F0 uint64; F1 uint64 }
type GenS10 struct { F0 uint8; F1 bool }
type GenS11 struct { F0 int; F1 uint32 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS14 struct { F0 uint; F1 uint }
type GenS15 struct { F0 bool; F1 uint }
func genFn0(p0 bool, p1 uint16) int16 {
    p0 = (p0 && p0)
    p0 = false
    v4 := GenS1{F0: "\n\t", F1: 42, F2: 1000}
    var v5 string = "a"
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 bool) uint32 {
    p0 = false
    var v6 bool = (577 != 324)
    fmt.Println(v6)
    for fi7 := 0; fi7 < 19; fi7++ {
        v6 = p0
        if ("" >= "hello") {
            var v8 uint16 = 1
            _ = v8
        } else {
            fmt.Println(fi7)
        }
        p0 = (255 > 32767)
        _ = fi7
    }
    _ = p0
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 42
    var v3 uint16 = 996
    var v9 uint64 = v2
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            switch fi11 {
            case 1:
                v9 = (v2 << ((65535) & 63))
            case 0:
                v1[2] = fi10
            case 286:
                if true {
                    if false {
                        v12 := []int{-2147483648}
                        if true {
                            switch fi10 {
                            case 127:
                                var v13 uint64 = 0
                                v13 = (v2 - (v9 / ((v13) | 1)))
                                v14 := GenS1{F0: (("\n\t" + "foo bar") + "foo bar"), F1: 65535, F2: -758}
                                _ = v13
                                _ = v14
                            case -561:
                                v15 := GenS2{F0: fi10}
                                _ = v15
                            }
                            var v16 string = "hello"
                            if true {
                                v17 := GenS2{F0: fi10}
                                _ = v17
                            } else {
                                v2 = (v2 >> ((1) & 63))
                                if (103 < 42) {
                                    var v18 float32 = 1e10
                                    fmt.Println(v2)
                                    switch fi10 {
                                    case 32768:
                                        fmt.Println(fi11, fi11)
                                    case 4294967295:
                                        v1[0] = 1
                                    case 127:
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v19 uint16 = 65535
                                            _ = v19
                                        }
                                    }
                                    _ = v18
                                } else {
                                    fmt.Println(v3, v9, v16)
                                    if true {
                                        var v20 int16 = 75
                                        if false {
                                            var v21 int32 = 911
                                            var v22 uint64 = 1
                                            _ = v21
                                            _ = v22
                                        }
                                        _ = v20
                                    }
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                }
                            }
                            _ = v16
                        } else {
                            v23 := GenS2{F0: (fi10 >> ((41) & 63))}
                            var v24 bool = (!false)
                            var v25 uint8 = 56
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        if (((-2.25 + 1.5) != ((float64(1.5) * (1e10 + (3.14159 + -2.25))) / 1.5)) || (false && (-2.25 < float32(float64(((0.1 * (0.0 * 0.1)) + 1.5)))))) {
                            v3 = v3
                        } else {
                            if (-96 >= 0) {
                                fmt.Println(v3)
                                v26 := GenS2{F0: fi10}
                                _ = v26
                            }
                            var v27 uint8 = 255
                            _ = v27
                        }
                        _ = v12
                    } else {
                        var v28 uint = 255
                        _ = v28
                    }
                    v29 := []int{fi11}
                    _ = len(v29)
                    _ = v29
                } else {
                    v30 := GenS2{F0: (fi10 - fi10)}
                    _ = v30
                }
            }
            v2 = (^(v2))
            v2 = v9
            _ = fi11
        }
        _ = fi10
    }
    v9 += (v9 | (v9 / ((1) | 1)))
    for fi31 := 0; fi31 < 13; fi31++ {
        var v32 string = "a"
        _ = len(v1)
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        if (((false || (fi33 != 1)) && ((^(v2)) == (v2 | v9))) && (false || true)) {
            if false {
                v34 := GenS3{F0: 636}
                for fi35 := 0; fi35 < 15; fi35++ {
                    var v36 int8 = 39
                    _ = fi35
                    _ = v36
                }
                _ = v34
            } else {
                for fi37 := 0; fi37 < 9; fi37++ {
                    v2 -= v2
                    _ = fi37
                }
            }
            if false {
                v9 -= (v2 * v2)
                var v38 uint8 = 127
                for fi39 := 0; fi39 < 8; fi39++ {
                    fmt.Println(v2, fi39)
                    _ = fi39
                }
                _ = v38
            } else {
                var v40 int16 = genFn0(true, v3)
                v40 = (v40 - v40)
                _ = v40
            }
            v41 := []int{fi33, fi33, fi33}
            _ = v41
        }
        _ = fi33
    }
    v42 := GenS4{F0: v3, F1: 1}
    for fi43 := 0; fi43 < 7; fi43++ {
        fmt.Println(v2, fi43)
        var v44 int64 = 0
        _ = fi43
        _ = v44
    }
    v45 := GenS5{F0: 9223372036854775807, F1: 2, F2: 1}
    for fi46 := 0; fi46 < 18; fi46++ {
        v1[417] = fi46
        var v47 uint8 = 128
        _ = fi46
        _ = v47
    }
    v3 = (^(v3))
    switch "hello" {
    case "":
        if true {
            for fi48 := 0; fi48 < 19; fi48++ {
                switch (^(fi48)) {
                case 1:
                    if true {
                        fmt.Println(v2, v9)
                        if (18446744073709551615 != 255) {
                            var v49 uint16 = v3
                            if true {
                                v50 := []int{-9223372036854775808, fi48}
                                for fi51 := 0; fi51 < 3; fi51++ {
                                    v52 := GenS5{F0: 42, F1: 51, F2: (v3 << ((42) & 63))}
                                    var v53 int16 = 0
                                    _ = fi51
                                    _ = v52
                                    _ = v53
                                }
                                for fi54 := 0; fi54 < 18; fi54++ {
                                    switch (-(fi48)) {
                                    case -919:
                                        if true {
                                            if ((36 <= 0) && ("foo bar" != "hello")) {
                                                fmt.Println(fi48)
                                                v55 := GenS6{F0: 4294967295}
                                                _ = v55
                                            } else {
                                                v1[42] = (-(fi54))
                                                fmt.Println(fi48)
                                            }
                                        } else {
                                            delete(v1, 2147483647)
                                            var v56 float64 = 1.5
                                            _ = v56
                                        }
                                    }
                                    v9 = 0
                                    v57 := GenS5{F0: 1, F1: 41, F2: v49}
                                    v50 = append(v50, fi54)
                                    if ((3.14159 + -2.25) <= 0.1) {
                                        if false {
                                            switch 1 {
                                            case 4294967295:
                                                if (!true) {
                                                    fmt.Println(v9, fi48)
                                                    v49 = (v49 ^ v3)
                                                    fmt.Println(fi54, fi54, fi48)
                                                    switch (fi48 | fi48) {
                                                    case 0:
                                                        var v58 int8 = -112
                                                        _ = v58
                                                    case 32767:
                                                        v50 = append(v50, fi54)
                                                    case 128:
                                                        v59 := v1[genFn1(("hello" > "a"))]
                                                        _ = v59
                                                    }
                                                    delete(v1, 2147483647)
                                                }
                                            case -347:
                                                switch (("\n\t" + "hello") + "hello") {
                                                case "hello":
                                                    v60 := v1[508]
                                                    _ = v60
                                                }
                                            case 559:
                                                v49 += (v49 >> ((2) & 63))
                                            default:
                                                v1[0] = fi48
                                            }
                                            v9 = (v9 ^ (v2 | (v2 % (((v9 - (v9 * 158))) | 1))))
                                            v2 -= 42
                                        }
                                        switch fi48 {
                                        case 1:
                                            v61 := GenS7{F0: "foo bar", F1: 0, F2: (fi54 ^ fi54)}
                                            _ = v61
                                        case 695:
                                            switch (fi54 + 4294967295) {
                                            case -426:
                                                if (false && (genFn0(false, v49) >= genFn0(true, 626))) {
                                                    var v62 bool = false
                                                    _ = v62
                                                } else {
                                                    v49 += (v3 >> ((205) & 63))
                                                }
                                            default:
                                                if false {
                                                    v9 = (v2 ^ (^(v9)))
                                                    fmt.Println(v49, fi54)
                                                } else {
                                                    if (0 > 0) {
                                                        v63 := GenS3{F0: 158}
                                                        v64 := GenS8{F0: -9223372036854775808, F1: 1, F2: 0}
                                                        _ = v63
                                                        _ = v64
                                                    } else {
                                                        v9 += v2
                                                    }
                                                }
                                                v65 := GenS8{F0: 9223372036854775807, F1: 65535, F2: uint8(v9)}
                                                var v66 uint16 = v3
                                                v49 += (v3 ^ 2)
                                                _ = v65
                                                _ = v66
                                            }
                                        default:
                                            if (2147483647 > 2147483647) {
                                                switch fi48 {
                                                case 9223372036854775807:
                                                    fmt.Println(v9, v9, v49)
                                                case 32768:
                                                    v67 := GenS7{F0: "a", F1: 2147483647, F2: (fi54 & fi48)}
                                                    _ = v67
                                                case 32767:
                                                    var v68 int16 = 1000
                                                    _ = v68
                                                default:
                                                    if false {
                                                        _ = len(v50)
                                                        if false {
                                                            if len(v50) >= 2 {
                                                                v50 = v50[1:len(v50)]
                                                            }
                                                            v69 := GenS3{F0: 916}
                                                            _ = v69
                                                        }
                                                        switch (fi48 ^ (fi54 - fi48)) {
                                                        case 128:
                                                            switch fi54 {
                                                            case 9223372036854775807:
                                                                fmt.Println(v3, fi48)
                                                            case 32768:
                                                                v2 -= v2
                                                            }
                                                        }
                                                        v1[genFn1(false)] = fi54
                                                        fmt.Println(fi48)
                                                        var v70 bool = (int16(v49) == -1)
                                                        if (42 == 65535) {
                                                            fmt.Println(fi54, v2)
                                                            _ = len(v1)
                                                        }
                                                        fmt.Println(fi54, v9, v49)
                                                        _ = v70
                                                    }
                                                }
                                                var v71 int = fi54
                                                _ = v71
                                            } else {
                                                if (!true) {
                                                    fmt.Println(fi48, fi54)
                                                    v3 += (v3 % ((v49) | 1))
                                                    v3 = (v3 % ((2) | 1))
                                                }
                                                v2 += 2
                                                var v72 string = "hello"
                                                _ = v72
                                            }
                                        }
                                    } else {
                                        var v73 uint = 81
                                        _ = len(v50)
                                        var v74 float32 = 1.5
                                        _ = v73
                                        _ = v74
                                    }
                                    _ = fi54
                                    _ = v57
                                }
                                _ = v50
                            }
                            if false {
                                var v75 int8 = 0
                                _ = v75
                            }
                            _ = v49
                        }
                        v76 := GenS9{F0: v9, F1: v2}
                        _ = v76
                    } else {
                        switch -161 {
                        case -2147483648:
                            var v77 int16 = 32767
                            _ = v77
                        case 128:
                            v3 = (^(v3))
                        case 398:
                            v78 := GenS10{F0: 0, F1: (v3 > v3)}
                            _ = v78
                        default:
                            if (!false) {
                                v79 := GenS8{F0: 9223372036854775807, F1: 529, F2: 255}
                                var v80 uint = 4294967295
                                if true {
                                    if (65535 != v3) {
                                        for fi81 := 0; fi81 < 19; fi81++ {
                                            v3 -= 32767
                                            v80 = v80
                                            var v82 int16 = -1
                                            _ = fi81
                                            _ = v82
                                        }
                                        v83 := GenS1{F0: "a", F1: v80, F2: 42}
                                        if false {
                                            v84 := GenS11{F0: (-(fi48)), F1: 324}
                                            v85 := GenS3{F0: (v80 | 343)}
                                            switch (fi48 - -1) {
                                            case -2147483648:
                                                for fi86 := 0; fi86 < 7; fi86++ {
                                                    fmt.Println(v3)
                                                    _ = fi86
                                                }
                                            default:
                                                var v87 uint = 736
                                                _ = v87
                                            }
                                            _ = v84
                                            _ = v85
                                        } else {
                                            var v88 bool = (0 == 150)
                                            if v88 {
                                                v9 = v2
                                                if (877 > 1) {
                                                    switch fi48 {
                                                    case 2147483647:
                                                        v88 = ((2147483648 <= 720) || false)
                                                    default:
                                                        if (1e10 <= ((float32(((((-2.25 * ((0.0 - 1.5) + 1e10)) + 0.1) * 1.5) + (1e10 - 0.0))) - (-2.25 + 0.1)) - 0.0)) {
                                                            v89 := []int{(fi48 ^ 9223372036854775807)}
                                                            v90 := GenS5{F0: -9223372036854775808, F1: 255, F2: 255}
                                                            v88 = (v9 >= 687)
                                                            _ = v89
                                                            _ = v90
                                                        } else {
                                                            for fi91 := 0; fi91 < 1; fi91++ {
                                                                v2 = v2
                                                                fmt.Println(fi91)
                                                                v92 := []int{}
                                                                _ = fi91
                                                                _ = v92
                                                            }
                                                            v88 = ((v88 && true) && v88)
                                                        }
                                                    }
                                                    var v93 int64 = 1
                                                    if (-1 >= 2147483647) {
                                                        v80 = (v80 + v80)
                                                        switch (fi48 >> (((v80 + 65535)) & 63)) {
                                                        case 128:
                                                            v94 := GenS4{F0: v3, F1: 1000}
                                                            _ = v94
                                                        }
                                                    }
                                                    _ = v93
                                                } else {
                                                    v80 -= v80
                                                    if (v88 && v88) {
                                                        for fi95 := 0; fi95 < 16; fi95++ {
                                                            switch (fi95 * fi48) {
                                                            case -2147483648:
                                                                delete(v1, 2)
                                                            case -9223372036854775808:
                                                                var v96 int64 = -486
                                                                _ = v96
                                                            case 0:
                                                                switch fi48 {
                                                                case 9223372036854775807:
                                                                    var v97 int64 = -9223372036854775808
                                                                    _ = v97
                                                                case 32767:
                                                                    v98 := GenS2{F0: fi48}
                                                                    _ = v98
                                                                default:
                                                                    v99 := []int{fi95}
                                                                    _ = v99
                                                                }
                                                                fmt.Println(v88, v88, v2)
                                                            default:
                                                                fmt.Println(fi95, v80, fi95)
                                                            }
                                                            v3 = 1
                                                            v3 = v3
                                                            _ = fi95
                                                        }
                                                    }
                                                    for fi100 := 0; fi100 < 11; fi100++ {
                                                        v88 = false
                                                        v88 = v88
                                                        v2 -= (^(v9))
                                                        _ = fi100
                                                    }
                                                }
                                            }
                                            var v101 bool = v88
                                            _ = v101
                                            _ = v88
                                        }
                                        _ = v83
                                    } else {
                                        var v102 bool = false
                                        switch 356 {
                                        case -2147483648:
                                            switch "foo bar" {
                                            case "hello":
                                                if true {
                                                    var v103 float64 = -2.25
                                                    if v102 {
                                                        fmt.Println(v103, v3)
                                                        v104 := GenS12{F0: (v3 + (^(v3)))}
                                                        _ = v104
                                                    } else {
                                                        v3 = (v3 - v3)
                                                    }
                                                    _ = v103
                                                }
                                            case "\n\t":
                                                v3 -= v3
                                            case "foo bar":
                                                if (float64(-2.25) != 1.5) {
                                                    switch ("\n\t" + "") {
                                                    case "foo bar":
                                                        v105 := GenS13{F0: v2, F1: 240, F2: 102}
                                                        _ = v105
                                                    case "a":
                                                        v106 := GenS5{F0: 9223372036854775807, F1: 0, F2: 255}
                                                        _ = v106
                                                    default:
                                                        var v107 uint = (v80 | (^(v80)))
                                                        _ = v107
                                                    }
                                                    for fi108 := 0; fi108 < 13; fi108++ {
                                                        var v109 float64 = 0.1
                                                        v3 -= (^(v3))
                                                        fmt.Println(v102, v109)
                                                        _ = fi108
                                                        _ = v109
                                                    }
                                                }
                                            }
                                        case 9223372036854775807:
                                            for fi110 := 0; fi110 < 12; fi110++ {
                                                fmt.Println(v9, v102)
                                                var v111 uint64 = 4294967295
                                                _ = fi110
                                                _ = v111
                                            }
                                        }
                                        _ = v102
                                    }
                                    v112 := []int{42}
                                    for fi113 := 0; fi113 < 12; fi113++ {
                                        if false {
                                            v1[0] = fi48
                                            if true {
                                                switch fi113 {
                                                case 462:
                                                    v3 = v3
                                                default:
                                                    v2 = (v9 + v9)
                                                }
                                                switch "a" {
                                                case "a":
                                                    var v114 uint16 = (v3 >> (((v80 + (^(v80)))) & 63))
                                                    _ = v114
                                                default:
                                                    v3 = v3
                                                }
                                                fmt.Println(fi48, v9, v9)
                                            }
                                        } else {
                                            v80 -= 2
                                            v80 = 65535
                                            if (!(0.0 >= float64(1e10))) {
                                                fmt.Println(fi48, v80, v9)
                                            } else {
                                                if (!(true || (("\n\t" + ("a" + "foo bar")) < "foo bar"))) {
                                                    var v115 int16 = 1
                                                    fmt.Println(v3, fi48)
                                                    _ = v115
                                                } else {
                                                    fmt.Println(fi113, v2, fi48)
                                                    if (true || false) {
                                                        switch "hello" {
                                                        case "hello":
                                                            var v116 float64 = 1.5
                                                            _ = v116
                                                        case "":
                                                            if len(v112) > 0 {
                                                                v112[0] = fi48
                                                            }
                                                        }
                                                        v117 := GenS13{F0: (v9 - (v9 / ((v2) | 1))), F1: 0, F2: 127}
                                                        _ = v117
                                                    }
                                                    var v118 uint32 = genFn1(false)
                                                    _ = v118
                                                }
                                                var v119 uint32 = 1
                                                v1[(v119 * uint32(v3))] = fi48
                                                _ = v119
                                            }
                                        }
                                        _ = fi113
                                    }
                                    _ = v112
                                }
                                _ = v79
                                _ = v80
                            }
                        }
                        for fi120 := 0; fi120 < 17; fi120++ {
                            v2 = v2
                            fmt.Println(v3, fi120)
                            switch "\n\t" {
                            case "\n\t":
                                v121 := []int{fi120, (fi48 + (fi120 / ((fi48) | 1))), (fi48 & (fi48 >> ((893) & 63))), fi120}
                                if true {
                                    if (299 < 553) {
                                        var v122 int16 = 32767
                                        var v123 int32 = 197
                                        _ = v122
                                        _ = v123
                                    }
                                }
                                _ = v121
                            case "a":
                                v124 := GenS14{F0: 42, F1: 403}
                                v125 := GenS8{F0: -1, F1: 182, F2: 58}
                                _ = v124
                                _ = v125
                            }
                            _ = fi120
                        }
                    }
                default:
                    v2 -= 18446744073709551615
                }
                _ = fi48
            }
        } else {
            v9 -= 2
            v1[genFn1((v9 >= 18446744073709551615))] = -1
        }
    case "\n\t":
        var v126 float32 = float32(1e10)
        _ = v126
    case "foo bar":
        switch "foo bar" {
        case "a":
            fmt.Println(v3)
        case "foo bar":
            v127 := GenS15{F0: (genFn1((false && true)) <= 255), F1: 1}
            _ = v127
        case "hello":
            var v128 string = ("" + "hello")
            _ = v128
        default:
            v2 += v9
        }
    }
    v9 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v45
    _ = v9
}
