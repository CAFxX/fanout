package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int64; F1 int; F2 string }
type GenS4 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS7 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 bool; F1 bool }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int64; F1 uint32; F2 bool }
type GenS12 struct { F0 int64; F1 string }
type GenS13 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS14 struct { F0 uint8; F1 uint }
func genFn0(p0 uint, p1 int32) int {
    v6 := GenS1{F0: 1, F1: p1}
    v7 := []int{-577}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -989
}
func genFn1() string {
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS2{F0: 246, F1: -32768}
        v10 := GenS3{F0: 1, F1: fi8, F2: "a"}
        var v11 uint32 = 0
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 string = ("\n\t" + "foo bar")
        v13 = v13
        if true {
            var v14 int8 = 100
            if true {
                if (-128 >= v14) {
                    switch "" {
                    case "foo bar":
                        var v15 uint8 = 255
                        _ = v15
                    }
                    if false {
                        v16 := []int{(fi12 % ((fi12) | 1)), (-(fi12))}
                        _ = v16
                    } else {
                        switch v13 {
                        case "hello":
                            v17 := GenS2{F0: 1000, F1: 42}
                            _ = v17
                        }
                    }
                    v18 := GenS4{F0: 255, F1: 81, F2: 1000}
                    _ = v18
                }
            }
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    v19 := []int{1, 4294967295, genFn0(766, -2147483648), genFn0(0, -232)}
    var v20 uint32 = 1
    _ = v19
    _ = v20
    return "\n\t"
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    var v3 float32 = 1e10
    var v4 bool = (42 == 0)
    var v5 int64 = -127
    for fi21 := 0; fi21 < 15; fi21++ {
        fmt.Println(v3)
        for fi22 := 0; fi22 < 11; fi22++ {
            v4 = v4
            switch "a" {
            case "\n\t":
                if false {
                    v4 = (27 <= 115)
                    v5 = (v5 >> ((65535) & 63))
                } else {
                    var v23 float64 = (((0.1 - 1e10) / (1e10 + ((-2.25 - float64(v3)) - -2.25))) + (float64(((v3 * float32(3.14159)) - v3)) * 0.1))
                    var v24 int64 = (^(v5))
                    _ = v23
                    _ = v24
                }
            case "foo bar":
                v25 := []int{fi22, genFn0(2, 42), 2147483647}
                _ = v25
            case "a":
                fmt.Println(fi21, fi22)
            default:
                v26 := GenS5{F0: 0, F1: 65535}
                _ = v26
            }
            fmt.Println(v5, v4)
            var v27 float64 = ((0.0 * 1e10) + 3.14159)
            v3 -= v3
            _ = fi22
            _ = v27
        }
        v3 += 0.1
        _ = fi21
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        if (false || false) {
            switch 1 {
            case 156:
                if v4 {
                    for fi29 := 0; fi29 < 19; fi29++ {
                        v30 := GenS6{F0: 127, F1: 65535, F2: 883}
                        v5 -= (v5 & (v5 % (((v5 - (v5 & 0))) | 1)))
                        if v4 {
                            fmt.Println(fi28)
                            v31 := []int{genFn0(65535, 1)}
                            _ = v31
                        }
                        _ = fi29
                        _ = v30
                    }
                    v4 = (false || v4)
                    if false {
                        v4 = false
                    }
                } else {
                    var v32 uint8 = 255
                    for fi33 := 0; fi33 < 12; fi33++ {
                        switch (fi28 & fi33) {
                        case -545:
                            v34 := GenS1{F0: (-(v5)), F1: -234}
                            _ = v34
                        case 128:
                            var v35 int8 = 1
                            _ = v35
                        default:
                            if v4 {
                                var v36 bool = (false || (("\n\t" <= ("hello" + "\n\t")) || v4))
                                _ = v36
                            }
                        }
                        v37 := GenS1{F0: (v5 & 42), F1: -2147483648}
                        if ((v32 / ((v32) | 1)) < (^(v32))) {
                            v38 := GenS1{F0: (v5 + 1), F1: int32(v5)}
                            v39 := GenS7{F0: 870, F1: 4294967295, F2: (v5 << ((255) & 63))}
                            _ = v38
                            _ = v39
                        }
                        _ = fi33
                        _ = v37
                    }
                    _ = v32
                }
            }
            for fi40 := 0; fi40 < 9; fi40++ {
                v41 := []int{fi40, fi28}
                if (!(65535 != 42)) {
                    switch "hello" {
                    case "hello":
                        delete(v1, "\n\t")
                    }
                    var v42 uint16 = 0
                    v43 := GenS8{F0: 0}
                    var v44 int32 = 2147483647
                    v45 := GenS8{F0: (v5 & (v5 / ((v5) | 1)))}
                    _ = v42
                    _ = v43
                    _ = v44
                    _ = v45
                } else {
                    v46 := GenS9{F0: (128 <= 255), F1: false}
                    _ = len(v41)
                    v4 = v4
                    _ = v46
                }
                v47 := GenS10{F0: -639}
                _ = fi40
                _ = v41
                _ = v47
            }
            switch ("hello" + (("a" + (("" + "\n\t") + "hello")) + "foo bar")) {
            case "foo bar":
                if v4 {
                    v5 = v5
                    if (42 == 538) {
                        for fi48 := 0; fi48 < 12; fi48++ {
                            var v49 int = fi48
                            if v4 {
                                if ((1e10 - ((((((1e10 / float64(0.1)) - float64(v3)) / ((1.5 - -2.25) * 1e10)) / 1e10) - -2.25) * (-2.25 / (3.14159 / float64(3.14159))))) > 0.1) {
                                    var v50 int8 = -128
                                    v50 = (^(v50))
                                    _ = v50
                                }
                                v51 := GenS7{F0: 42, F1: 2, F2: (v5 ^ -1)}
                                _ = v51
                            } else {
                                v5 += v5
                                if true {
                                    switch "a" {
                                    case "a":
                                        switch "a" {
                                        case "":
                                            switch "a" {
                                            case "foo bar":
                                                if (!true) {
                                                    v52 := []int{(fi48 & (fi48 << ((18446744073709551615) & 63))), (fi28 | -74), 869, 32768}
                                                    _ = v52
                                                }
                                            case "a":
                                                switch (-(fi28)) {
                                                case 9223372036854775807:
                                                    var v53 bool = v4
                                                    _ = v53
                                                default:
                                                    var v54 uint16 = 371
                                                    _ = v54
                                                }
                                            default:
                                                if false {
                                                    v55 := GenS7{F0: 2, F1: 42, F2: v5}
                                                    if v4 {
                                                        v56 := []int{(fi28 | -1)}
                                                        _ = v56
                                                    }
                                                    var v57 uint16 = 1
                                                    _ = v55
                                                    _ = v57
                                                } else {
                                                    if (false || true) {
                                                        fmt.Println(v5, v4, v49)
                                                        var v58 string = ("hello" + "foo bar")
                                                        if (2147483648 > 42) {
                                                            var v59 uint16 = 0
                                                            if (77 == 2) {
                                                                v5 += (v5 % ((-1) | 1))
                                                            }
                                                            fmt.Println(fi28)
                                                            var v60 uint = 4294967295
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                        _ = v58
                                                    }
                                                }
                                            }
                                        case "a":
                                            switch "hello" {
                                            case "hello":
                                                if true {
                                                    if false {
                                                        switch (v49 + fi48) {
                                                        case 127:
                                                            fmt.Println(v49)
                                                        case 0:
                                                            var v61 int64 = (v5 ^ (v5 | v5))
                                                            _ = v61
                                                        }
                                                        var v62 uint16 = 810
                                                        var v63 uint64 = 4294967295
                                                        v64 := GenS7{F0: v62, F1: (v63 & (v63 & 1)), F2: v5}
                                                        var v65 bool = (216 != uint8(v62))
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                        _ = v65
                                                    } else {
                                                        switch v49 {
                                                        case 2147483647:
                                                            if true {
                                                                var v66 uint8 = 127
                                                                v67 := []int{42, genFn0(2, -295), -9223372036854775808, (v49 + genFn0(42, 0))}
                                                                switch ("" + "foo bar") {
                                                                case "":
                                                                    v68 := GenS1{F0: (v5 | v5), F1: 42}
                                                                    _ = v68
                                                                case "foo bar":
                                                                    v69 := GenS9{F0: ((!(!((v5 + v5) < (^(v5))))) || ((v4 && true) && (!v4))), F1: v4}
                                                                    _ = v69
                                                                default:
                                                                    v70 := GenS9{F0: ((((float32(0.0) - v3) / 1e10) > v3) || ((!(2147483647 > 42)) || true)), F1: false}
                                                                    _ = v70
                                                                }
                                                                _ = v66
                                                                _ = v67
                                                            } else {
                                                                switch fi28 {
                                                                case 1:
                                                                    var v71 int8 = 100
                                                                    _ = v71
                                                                case 32767:
                                                                    fmt.Println(v5, v49, fi48)
                                                                case 0:
                                                                    v72 := GenS1{F0: v5, F1: -2147483648}
                                                                    _ = v72
                                                                }
                                                                fmt.Println(v3, fi28, v49)
                                                                v73 := []int{}
                                                                var v74 float64 = 0.1
                                                                _ = v73
                                                                _ = v74
                                                            }
                                                        }
                                                        var v75 uint32 = 2147483647
                                                        switch "" {
                                                        case "hello":
                                                            var v76 int = (v49 & (v49 * (^(fi48))))
                                                            _ = v76
                                                        case "":
                                                            switch ("foo bar" + "foo bar") {
                                                            case "a":
                                                                v77 := GenS10{F0: -1}
                                                                _ = v77
                                                            case "foo bar":
                                                                fmt.Println(v3)
                                                            }
                                                        case "\n\t":
                                                            v3 += 3.14159
                                                        default:
                                                            v78 := v2[(v75 - 2147483647)]
                                                            _ = v78
                                                        }
                                                        fmt.Println(v49, v4)
                                                        v75 = (v75 % ((v75) | 1))
                                                        _ = len(v2)
                                                        _ = v75
                                                    }
                                                    if v4 {
                                                        v4 = true
                                                        v3 = (-2.25 - (3.14159 + 0.0))
                                                        v79 := []int{genFn0(0, 0)}
                                                        var v80 uint = 0
                                                        _ = v79
                                                        _ = v80
                                                    } else {
                                                        fmt.Println(fi48, fi48, v49)
                                                        v81 := GenS11{F0: -9223372036854775808, F1: 2147483648, F2: v4}
                                                        v82 := []int{v49, fi28, -1, fi28}
                                                        v83 := GenS10{F0: 0}
                                                        _ = v81
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                    v84 := GenS4{F0: 1, F1: 778, F2: -491}
                                                    _ = v84
                                                }
                                            case "\n\t":
                                                v85 := GenS2{F0: 909, F1: -1}
                                                _ = v85
                                            default:
                                                var v86 int8 = 100
                                                _ = v86
                                            }
                                        }
                                    case "":
                                        if ((v4 || (2 > 1)) && ((255 > 802) || true)) {
                                            if ("" < "a") {
                                                if ("hello" > "\n\t") {
                                                    if (v5 <= 9223372036854775807) {
                                                        v5 += (v5 | (v5 & (v5 - 9223372036854775807)))
                                                        v87 := GenS12{F0: v5, F1: ""}
                                                        _ = v87
                                                    } else {
                                                        var v88 float64 = 1e10
                                                        _ = v88
                                                    }
                                                } else {
                                                    var v89 int8 = -128
                                                    v90 := GenS11{F0: v5, F1: 2, F2: (!(1 == -416))}
                                                    _ = v89
                                                    _ = v90
                                                }
                                            }
                                        }
                                    default:
                                        switch v49 {
                                        case 113:
                                            if (!((true && ((1e10 + (v3 - ((float32((0.1 + -2.25)) + 1e10) * v3))) != ((3.14159 / (1.5 / v3)) / ((((v3 + float32(-2.25)) * v3) - ((v3 - 0.1) - v3)) + v3)))) && (18446744073709551615 != 42))) {
                                                v4 = false
                                            } else {
                                                fmt.Println(fi48, fi48)
                                            }
                                        case 4294967295:
                                            if (!false) {
                                                v5 = (v5 + -1)
                                                var v91 int8 = -128
                                                _ = v91
                                            } else {
                                                v49 += (fi28 - (^(fi28)))
                                            }
                                        case 419:
                                            v49 -= v49
                                        default:
                                            if ((1.5 >= ((0.1 / 1.0) - 1.5)) && (!(!((18446744073709551615 > 2) || v4)))) {
                                                v5 += -9223372036854775808
                                                fmt.Println(v49, v49)
                                            } else {
                                                v3 = 1.5
                                            }
                                        }
                                        fmt.Println(v4, v5)
                                    }
                                    v4 = v4
                                    var v92 int8 = -47
                                    _ = v92
                                } else {
                                    fmt.Println(v49, v49)
                                    var v93 uint64 = 42
                                    var v94 uint8 = 66
                                    _ = v93
                                    _ = v94
                                }
                            }
                            switch (("foo bar" + "") + "") {
                            case "a":
                                switch "a" {
                                case "hello":
                                    if v4 {
                                        if v4 {
                                            var v95 uint32 = 2
                                            _ = v95
                                        }
                                        fmt.Println(v49)
                                    }
                                }
                            }
                            v4 = true
                            _ = len(v1)
                            var v96 string = ("" + "")
                            v49 = -490
                            _ = fi48
                            _ = v49
                            _ = v96
                        }
                    }
                    var v97 int8 = -1
                    _ = v97
                }
            case "":
                if (1 == 42) {
                    var v98 int8 = 100
                    _ = v98
                } else {
                    switch ("a" + ("" + "a")) {
                    case "":
                        var v99 uint8 = 53
                        _ = v99
                    default:
                        v100 := []int{genFn0(807, 42), fi28}
                        _ = v100
                    }
                    var v101 uint = 65535
                    fmt.Println(v5)
                    _ = v101
                }
            case "\n\t":
                v102 := []int{fi28, fi28, fi28}
                _ = v102
            }
        } else {
            v1["a"] = true
        }
        _ = fi28
    }
    switch -246 {
    case -260:
        if (((v3 * (-2.25 / ((0.0 + 0.1) - (v3 + -2.25)))) + (v3 + v3)) == ((v3 - float32(-2.25)) / 0.1)) {
            if (v4 || v4) {
                v103 := GenS13{F0: 2147483647, F1: 291, F2: 2147483647}
                _ = v103
            }
            for fi104 := 0; fi104 < 2; fi104++ {
                switch fi104 {
                case 267:
                    v105 := GenS9{F0: (v4 || v4), F1: (!true)}
                    _ = v105
                case -2147483648:
                    if v4 {
                        v3 += 3.14159
                        if ("a" < "") {
                            v4 = (false || true)
                        } else {
                            if false {
                                fmt.Println(fi104, fi104, v3)
                            } else {
                                v4 = (!(1 == 648))
                                var v106 int64 = v5
                                _ = v106
                            }
                            var v107 string = ("" + ("\n\t" + "a"))
                            _ = v107
                        }
                    }
                case 128:
                    for fi108 := 0; fi108 < 17; fi108++ {
                        _ = len(v1)
                        v109 := GenS14{F0: 127, F1: 882}
                        _ = fi108
                        _ = v109
                    }
                default:
                    if true {
                        fmt.Println(v4)
                        _ = len(v2)
                        v5 += 875
                    }
                }
                _ = fi104
            }
        } else {
            var v110 int64 = v5
            if (-32768 != -32768) {
                fmt.Println(v3)
                v111 := []int{}
                v112 := GenS7{F0: 65535, F1: 2, F2: -1}
                _ = v111
                _ = v112
            }
            _ = v110
        }
    case 127:
        delete(v1, "foo bar")
    case 128:
        var v113 uint64 = 0
        _ = v113
    default:
        fmt.Println(v4, v4, v3)
    }
    v114 := []int{-9223372036854775808, 42}
    fmt.Println(v3, v5)
    v4 = v4
    v4 = (v4 && (42 != 42))
    for fi115 := 0; fi115 < 10; fi115++ {
        fmt.Println(fi115, v3, v4)
        _ = fi115
    }
    for fi116 := 0; fi116 < 5; fi116++ {
        var v117 int64 = 0
        v114 = append(v114, (fi116 & (fi116 / ((genFn0(454, 0)) | 1))))
        v117 = 9223372036854775807
        _ = fi116
        _ = v117
    }
    var v118 uint16 = 1
    for fi119 := 0; fi119 < 17; fi119++ {
        _ = len(v2)
        if len(v114) > 0 {
            v114[0] = (fi119 - (fi119 % (((fi119 - (fi119 - fi119))) | 1)))
        }
        _ = fi119
    }
    var v120 uint8 = 46
    switch "" {
    case "a":
        v3 += ((1e10 + float32(-2.25)) / 1.5)
    }
    v121 := GenS4{F0: 2, F1: 933, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v114:", len(v114), v114)
    fmt.Println("v118:", v118)
    fmt.Println("v120:", v120)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v114
    _ = v118
    _ = v120
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
