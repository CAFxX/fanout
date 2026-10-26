package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int8; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint; F2 int32 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 int64; F1 int32; F2 int16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int64; F1 string; F2 int64 }
type GenS11 struct { F0 int8; F1 int64 }
type GenS12 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS13 struct { F0 uint; F1 uint8; F2 int64 }
type GenS14 struct { F0 uint16; F1 int8; F2 uint }
type GenS15 struct { F0 int; F1 uint32; F2 int8 }
type GenS16 struct { F0 int64; F1 uint32 }
type GenS17 struct { F0 uint16; F1 uint }
type GenS18 struct { F0 int64; F1 uint32 }
type GenS19 struct { F0 string; F1 string; F2 uint }
type GenS20 struct { F0 bool; F1 uint8; F2 uint }
type GenS21 struct { F0 int; F1 uint64; F2 int8 }
func genFn0(p0 string, p1 uint16, p2 uint16) int {
    fmt.Println(p1)
    var v5 float32 = (0.0 / 1.0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -806
}
func genFn1() int64 {
    switch "hello" {
    case "foo bar":
        v6 := GenS1{F0: 1}
        _ = v6
    case "hello":
        if (-1 < 9) {
            if (false && true) {
                v7 := GenS1{F0: 992}
                var v8 int16 = 1000
                if (42 != -701) {
                    for fi9 := 0; fi9 < 2; fi9++ {
                        switch "" {
                        case "\n\t":
                            if true {
                                v8 = 562
                                switch ("" + (("foo bar" + "foo bar") + "")) {
                                case "":
                                    v8 -= (v8 - (^(v8)))
                                default:
                                    v8 += v8
                                }
                                var v10 uint64 = 0
                                _ = v10
                            } else {
                                v8 -= v8
                                v11 := GenS1{F0: 416}
                                _ = v11
                            }
                        case "a":
                            v8 = v8
                        case "foo bar":
                            v8 += (v8 * 32767)
                        }
                        fmt.Println(v8)
                        v12 := GenS1{F0: -1}
                        _ = fi9
                        _ = v12
                    }
                    switch (("" + ("hello" + "\n\t")) + "") {
                    case "hello":
                        var v13 uint16 = 65535
                        _ = v13
                    default:
                        var v14 uint64 = 42
                        _ = v14
                    }
                    v15 := []int{0, -338}
                    _ = v15
                } else {
                    var v16 int = -579
                    v17 := GenS1{F0: -1}
                    v18 := GenS2{F0: v16, F1: -128, F2: 18446744073709551615}
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = v7
                _ = v8
            } else {
                var v19 int = -390
                v20 := []int{}
                _ = v19
                _ = v20
            }
            v21 := GenS3{F0: -32768, F1: 1, F2: 42}
            _ = v21
        } else {
            v22 := []int{42, 32768, genFn0(("\n\t" + ("a" + "a")), 2, 1), genFn0("\n\t", 65535, 0)}
            var v23 int32 = -760
            var v24 int32 = (v23 << ((293) & 63))
            _ = v22
            _ = v23
            _ = v24
        }
    default:
        var v25 uint16 = 65535
        _ = v25
    }
    v26 := []int{527, -787, 642}
    var v27 int16 = -351
    if false {
        switch ("\n\t" + "a") {
        case "hello":
            if ((((!(-1 <= 99)) && (!((("foo bar" + "\n\t") + (("hello" + "") + ("" + "hello"))) != "a"))) || (1e10 >= ((1.5 * 1e10) * 1.5))) || false) {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v27)
                case "hello":
                    if len(v26) > 0 {
                        v26[0] = 4294967295
                    }
                case "a":
                    fmt.Println(v27)
                default:
                    v27 = v27
                }
                v26 = append(v26, 748)
            }
        default:
            if true {
                v27 = (v27 * v27)
            } else {
                if (true && (42 > 1)) {
                    v27 = v27
                    if false {
                        var v28 int32 = -2147483648
                        for fi29 := 0; fi29 < 10; fi29++ {
                            v30 := GenS2{F0: fi29, F1: -39, F2: 1}
                            _ = fi29
                            _ = v30
                        }
                        _ = v28
                    }
                }
                var v31 uint32 = 372
                if (88 != 127) {
                    if (-2.25 == (-2.25 / ((0.0 - 3.14159) / 0.1))) {
                        fmt.Println(v27, v27)
                        v27 = (-(v27))
                    }
                    for fi32 := 0; fi32 < 15; fi32++ {
                        v27 = v27
                        switch fi32 {
                        case 0:
                            switch "\n\t" {
                            case "foo bar":
                                switch "\n\t" {
                                case "":
                                    var v33 float64 = float64(float32(0.1))
                                    _ = v33
                                case "foo bar":
                                    v27 -= v27
                                default:
                                    v34 := GenS3{F0: v27, F1: 18446744073709551615, F2: 0}
                                    _ = v34
                                }
                            case "":
                                var v35 uint32 = (v31 ^ (v31 - (^(v31))))
                                _ = v35
                            case "a":
                                if len(v26) >= 2 {
                                    v26 = v26[1:len(v26)]
                                }
                            }
                        case -2147483648:
                            v27 += v27
                        default:
                            fmt.Println(v31, v27, v27)
                        }
                        _ = fi32
                    }
                } else {
                    switch genFn0("foo bar", 65535, 907) {
                    case 28:
                        v26 = append(v26, -217)
                    case 127:
                        fmt.Println(v27, v27)
                    case 4294967295:
                        if false {
                            if false {
                                v36 := GenS3{F0: v27, F1: 255, F2: 0}
                                _ = v36
                            }
                            if (float32(3.14159) != float32(-2.25)) {
                                v37 := GenS4{F0: 1, F1: 42}
                                v27 += (v27 % (((v27 * v27)) | 1))
                                if false {
                                    var v38 uint16 = 851
                                    var v39 float32 = 0.1
                                    var v40 uint = 263
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                }
                                _ = v37
                            }
                            var v41 uint64 = 18446744073709551615
                            _ = v41
                        } else {
                            if true {
                                var v42 int = -641
                                v43 := GenS5{F0: 98, F1: -91}
                                _ = v42
                                _ = v43
                            }
                        }
                    default:
                        switch ("foo bar" + "hello") {
                        case "hello":
                            var v44 uint32 = 435
                            _ = v44
                        }
                    }
                    v31 += (^(v31))
                }
                _ = v31
            }
        }
        var v45 bool = true
        _ = v45
    }
    _ = v26
    _ = v27
    return 1
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    var v4 int8 = 100
    var v46 uint32 = 2147483648
    switch 128 {
    case 32767:
        v47 := v3[""]
        _ = v47
    case -1:
        v4 = 110
    }
    for fi48 := 0; fi48 < 11; fi48++ {
        _ = len(v2)
        delete(v3, ("foo bar" + ("" + "a")))
        _ = fi48
    }
    if false {
        var v49 int32 = 0
        for fi50 := 0; fi50 < 18; fi50++ {
            fmt.Println(v4)
            v51 := GenS6{F0: -9223372036854775808, F1: (v49 * -938), F2: (v1 | 1)}
            v52 := v3["hello"]
            _ = fi50
            _ = v51
            _ = v52
        }
        switch "\n\t" {
        case "hello":
            for fi53 := 0; fi53 < 16; fi53++ {
                v3[("a" + (("\n\t" + "hello") + "\n\t"))] = (!false)
                if true {
                    if ((0.1 >= (0.1 - 3.14159)) && (!false)) {
                        for fi54 := 0; fi54 < 18; fi54++ {
                            if false {
                                var v55 float32 = 0.0
                                var v56 string = "a"
                                v49 += (v49 % ((42) | 1))
                                _ = v55
                                _ = v56
                            } else {
                                if (0.0 != float64((1.5 * (3.14159 - 0.0)))) {
                                    if (696 == -1) {
                                        var v57 uint32 = (v46 - v46)
                                        switch (("a" + "a") + "") {
                                        case "a":
                                            fmt.Println(v46)
                                        case "":
                                            v58 := GenS3{F0: 32767, F1: 18446744073709551615, F2: (v49 | (^(v49)))}
                                            _ = v58
                                        case "hello":
                                            var v59 uint8 = 127
                                            _ = v59
                                        }
                                        _ = v57
                                    } else {
                                        if (255 != 2) {
                                            v60 := GenS7{F0: genFn0(("hello" + ""), 42, 0)}
                                            var v61 float32 = (0.1 + 1.5)
                                            v46 -= (v46 & v46)
                                            _ = v60
                                            _ = v61
                                        }
                                        if (!true) {
                                            var v62 string = ("a" + "hello")
                                            _ = v62
                                        } else {
                                            switch 1 {
                                            case -9223372036854775808:
                                                if true {
                                                    v63 := GenS3{F0: (v1 % (((v1 ^ 294)) | 1)), F1: 745, F2: (v49 << ((582) & 63))}
                                                    _ = v63
                                                }
                                            case 556:
                                                fmt.Println(v4, fi53, v1)
                                            }
                                        }
                                        delete(v2, (v46 * 1))
                                        var v64 int8 = (v4 & -97)
                                        if (4294967295 < 2) {
                                            if (!(978 > 255)) {
                                                v46 += (v46 - (v46 - v46))
                                            } else {
                                                v64 = v4
                                            }
                                            v65 := GenS8{F0: 1}
                                            _ = v65
                                        } else {
                                            v66 := []int{}
                                            _ = v66
                                        }
                                        _ = v64
                                    }
                                }
                                switch fi54 {
                                case 32767:
                                    v67 := []int{914, genFn0("\n\t", 1, 720)}
                                    var v68 uint32 = (v46 * (v46 & 2147483647))
                                    _ = v67
                                    _ = v68
                                case 933:
                                    var v69 float64 = 1.5
                                    _ = v69
                                case -675:
                                    v70 := []int{(fi54 ^ genFn0((("foo bar" + "hello") + "\n\t"), 32767, 255)), fi53, (fi53 / ((9223372036854775807) | 1)), genFn0(("a" + ("foo bar" + "a")), 42, 42)}
                                    _ = v70
                                default:
                                    switch -721 {
                                    case 2147483647:
                                        _ = len(v3)
                                    case 128:
                                        v71 := GenS3{F0: (^(v1)), F1: 478, F2: -2147483648}
                                        v2[v46] = false
                                        _ = v71
                                    case -1:
                                        var v72 bool = (!(!false))
                                        var v73 uint64 = 42
                                        _ = v72
                                        _ = v73
                                    }
                                }
                                if (-793 <= 1) {
                                    fmt.Println(v46, v1, v49)
                                    if (1 == 2) {
                                        v49 = (v49 - -2147483648)
                                    } else {
                                        fmt.Println(v46)
                                    }
                                    var v74 uint8 = 0
                                    v75 := []int{(fi53 ^ (-(fi54)))}
                                    v76 := v3[("foo bar" + "")]
                                    _ = v74
                                    _ = v75
                                    _ = v76
                                }
                            }
                            v49 -= v49
                            var v77 float64 = 1e10
                            if false {
                                switch 2147483647 {
                                case 146:
                                    v78 := GenS4{F0: v49, F1: 0}
                                    _ = v78
                                }
                                if true {
                                    v4 += (v4 - v4)
                                    v77 = float64(0.0)
                                } else {
                                    if false {
                                        v79 := GenS1{F0: 979}
                                        _ = v79
                                    } else {
                                        fmt.Println(v46)
                                        if (0 == 42) {
                                            var v80 uint32 = 42
                                            v81 := GenS9{F0: v46}
                                            _ = v80
                                            _ = v81
                                        } else {
                                            v46 = (v46 % (((v46 & (v46 ^ 42))) | 1))
                                            if true {
                                                var v82 int32 = v49
                                                v83 := []int{(fi54 >> ((2) & 63)), (fi53 | int(v4))}
                                                var v84 uint64 = 18446744073709551615
                                                _ = v82
                                                _ = v83
                                                _ = v84
                                            }
                                        }
                                    }
                                    var v85 uint64 = 1
                                    switch fi54 {
                                    case -1:
                                        v4 = (^(v4))
                                    case -383:
                                        fmt.Println(v85, v1, v46)
                                    default:
                                        v1 -= int16(fi53)
                                        if (true || (!false)) {
                                            var v86 uint16 = 844
                                            _ = v86
                                        } else {
                                            v77 = ((1.5 / (1e10 * (v77 * 3.14159))) - float64(1e10))
                                            if (-679 != (v49 * (v49 % ((-2147483648) | 1)))) {
                                                var v87 float64 = (-2.25 * (1e10 + float64(1e10)))
                                                if (int64(v85) == 9223372036854775807) {
                                                    switch genFn0(("foo bar" + ""), 65535, 255) {
                                                    case 32767:
                                                        var v88 int64 = 1
                                                        _ = v88
                                                    }
                                                }
                                                _ = v87
                                            } else {
                                                var v89 float32 = float32((v77 / 1.5))
                                                v89 = v89
                                                _ = v89
                                            }
                                        }
                                        switch "\n\t" {
                                        case "":
                                            if (!(!(!false))) {
                                                switch (^(fi53)) {
                                                case 128:
                                                    v49 = v49
                                                case 1:
                                                    if true {
                                                        if true {
                                                            v1 += (v1 / ((232) | 1))
                                                            switch "" {
                                                            case "hello":
                                                                v90 := GenS10{F0: -594, F1: "a", F2: 9223372036854775807}
                                                                _ = v90
                                                            default:
                                                                switch (fi54 % (((fi54 * fi54)) | 1)) {
                                                                case 231:
                                                                    var v91 uint16 = 42
                                                                    v92 := []int{0, genFn0("\n\t", (v91 & (v91 - v91)), v91), fi54}
                                                                    _ = len(v92)
                                                                    _ = v91
                                                                    _ = v92
                                                                case 127:
                                                                    v49 -= v49
                                                                default:
                                                                    v77 = v77
                                                                }
                                                            }
                                                            delete(v2, 656)
                                                        } else {
                                                            v93 := GenS11{F0: -1, F1: 42}
                                                            if ((float64((float32(v77) + 1e10)) + v77) != (float64((0.1 * 0.1)) * (v77 / (-2.25 + 1.5)))) {
                                                                switch 42 {
                                                                case 32768:
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        var v94 uint16 = 877
                                                                        _ = v94
                                                                    default:
                                                                        var v95 float64 = v77
                                                                        _ = v95
                                                                    }
                                                                default:
                                                                    var v96 int16 = (v1 / ((-1) | 1))
                                                                    _ = v96
                                                                }
                                                            }
                                                            _ = v93
                                                        }
                                                        v97 := []int{(fi54 ^ (fi54 % (((fi54 << ((0) & 63))) | 1))), fi53}
                                                        v1 += v1
                                                        var v98 int32 = v49
                                                        var v99 int = (fi53 << ((4294967295) & 63))
                                                        _ = v97
                                                        _ = v98
                                                        _ = v99
                                                    } else {
                                                        switch (("a" + "a") + ((((("\n\t" + "foo bar") + "hello") + "\n\t") + ("a" + ("a" + ("hello" + "a")))) + "")) {
                                                        case "\n\t":
                                                            switch int(v49) {
                                                            case 1:
                                                                var v100 float64 = v77
                                                                _ = v100
                                                            default:
                                                                v101 := GenS12{F0: (v49 / (((-(v49))) | 1)), F1: 168, F2: 915}
                                                                _ = v101
                                                            }
                                                        case "hello":
                                                            if false {
                                                                v102 := []int{}
                                                                v2[1] = false
                                                                _ = v102
                                                            } else {
                                                                if (!false) {
                                                                    if true {
                                                                        v46 = 831
                                                                    } else {
                                                                        if false {
                                                                            var v103 uint = 255
                                                                            if true {
                                                                                var v104 int32 = (v49 * -1)
                                                                                switch (fi53 - (fi53 & 2147483647)) {
                                                                                case 128:
                                                                                    var v105 uint16 = 75
                                                                                    _ = v105
                                                                                default:
                                                                                    v46 = (^(v46))
                                                                                }
                                                                                _ = v104
                                                                            } else {
                                                                                if (-1 < fi53) {
                                                                                    var v106 int64 = -1
                                                                                    delete(v2, v46)
                                                                                    _ = v106
                                                                                } else {
                                                                                    var v107 bool = ((-(v49)) > (-(v49)))
                                                                                    if (!v107) {
                                                                                        var v108 int64 = -1
                                                                                        v109 := []int{(fi53 + (fi53 % ((genFn0("a", 2, 857)) | 1))), fi53, -994}
                                                                                        var v110 int64 = (v108 << ((uint(v46)) & 63))
                                                                                        _ = v108
                                                                                        _ = v109
                                                                                        _ = v110
                                                                                    } else {
                                                                                        var v111 float64 = v77
                                                                                        var v112 bool = v107
                                                                                        delete(v2, v46)
                                                                                        _ = v111
                                                                                        _ = v112
                                                                                    }
                                                                                    _ = v107
                                                                                }
                                                                                v85 -= 1
                                                                                switch (fi54 + fi53) {
                                                                                case 42:
                                                                                    var v113 float64 = 1e10
                                                                                    _ = v113
                                                                                default:
                                                                                    var v114 int16 = v1
                                                                                    _ = v114
                                                                                }
                                                                            }
                                                                            _ = v103
                                                                        } else {
                                                                            v115 := GenS11{F0: int8(v1), F1: 944}
                                                                            if true {
                                                                                if ("a" > "") {
                                                                                    v4 -= (v4 >> ((576) & 63))
                                                                                }
                                                                            }
                                                                            _ = v115
                                                                        }
                                                                        fmt.Println(fi53, v1)
                                                                        v4 += (v4 / (((v4 * (v4 - v4))) | 1))
                                                                        v116 := GenS13{F0: 426, F1: 2, F2: 42}
                                                                        _ = v116
                                                                    }
                                                                    v117 := GenS5{F0: 127, F1: (v4 / ((v4) | 1))}
                                                                    _ = len(v2)
                                                                    _ = v117
                                                                }
                                                                _ = len(v2)
                                                                v77 -= -2.25
                                                            }
                                                        default:
                                                            v4 = v4
                                                        }
                                                        if false {
                                                            var v118 bool = (v1 != v1)
                                                            v4 = (^(v4))
                                                            _ = v118
                                                        }
                                                        delete(v3, (("foo bar" + ("hello" + "\n\t")) + (("\n\t" + (("a" + "a") + "a")) + "hello")))
                                                    }
                                                    v4 = (v4 * (v4 + 53))
                                                }
                                                switch "a" {
                                                case "":
                                                    v77 = (v77 / v77)
                                                default:
                                                    if (!true) {
                                                        fmt.Println(v1, v49)
                                                    }
                                                }
                                            } else {
                                                var v119 float64 = 3.14159
                                                v49 -= v49
                                                var v120 int8 = v4
                                                _ = v119
                                                _ = v120
                                            }
                                        case "foo bar":
                                            if (0 >= (fi54 / ((fi53) | 1))) {
                                                v121 := []int{}
                                                var v122 int16 = (v1 / (((v1 / ((v1) | 1))) | 1))
                                                _ = v121
                                                _ = v122
                                            } else {
                                                fmt.Println(v77, fi54)
                                            }
                                        case "hello":
                                            if true {
                                                fmt.Println(v46)
                                            } else {
                                                v77 -= ((v77 / float64(1e10)) * (v77 + -2.25))
                                                var v123 string = ("foo bar" + "a")
                                                v77 -= v77
                                                _ = v123
                                            }
                                        }
                                        var v124 int = genFn0((("\n\t" + "") + "hello"), 1, 0)
                                        _ = v124
                                    }
                                    _ = v85
                                }
                                var v125 uint32 = 234
                                _ = v125
                            } else {
                                switch "a" {
                                case "hello":
                                    _ = len(v2)
                                case "\n\t":
                                    if (true || ((1 > 1) && true)) {
                                        v49 = v49
                                    }
                                case "a":
                                    v126 := GenS13{F0: 18446744073709551615, F1: 1, F2: genFn1()}
                                    _ = v126
                                }
                            }
                            switch "a" {
                            case "\n\t":
                                switch (fi54 ^ genFn0(("" + ((("a" + (("\n\t" + ("a" + ("" + (("a" + "foo bar") + "")))) + "hello")) + (("\n\t" + ("a" + "hello")) + ("foo bar" + "\n\t"))) + "foo bar")), 1, 846)) {
                                case -9223372036854775808:
                                    var v127 int64 = -9223372036854775808
                                    _ = v127
                                case -2147483648:
                                    v3["a"] = ((-(v49)) < 2147483647)
                                }
                            case "a":
                                fmt.Println(fi53)
                            case "hello":
                                v128 := []int{(fi53 / (((-(fi53))) | 1)), genFn0("\n\t", 255, 420), -785}
                                _ = v128
                            default:
                                v129 := GenS14{F0: 510, F1: (v4 * (^(v4))), F2: 593}
                                _ = v129
                            }
                            v130 := GenS15{F0: (^(fi54)), F1: (v46 % ((v46) | 1)), F2: v4}
                            _ = fi54
                            _ = v130
                            _ = v77
                        }
                    } else {
                        if true {
                            if (!true) {
                                v49 += (v49 % ((2147483647) | 1))
                                v131 := []int{9223372036854775807, fi53, genFn0("", 440, 42)}
                                _ = v131
                            } else {
                                if ((((("hello" + "hello") + "hello") + "a") >= (("hello" + "hello") + "a")) || (!false)) {
                                    switch ("\n\t" + "a") {
                                    case "\n\t":
                                        var v132 int32 = (v49 * (v49 | v49))
                                        _ = v132
                                    default:
                                        for fi133 := 0; fi133 < 5; fi133++ {
                                            var v134 int32 = 880
                                            _ = fi133
                                            _ = v134
                                        }
                                    }
                                    v135 := GenS16{F0: 1, F1: 42}
                                    if false {
                                        var v136 int32 = v49
                                        switch "a" {
                                        case "":
                                            for fi137 := 0; fi137 < 4; fi137++ {
                                                v138 := GenS5{F0: 72, F1: -1}
                                                var v139 int16 = (^(v1))
                                                _ = fi137
                                                _ = v138
                                                _ = v139
                                            }
                                        case "foo bar":
                                            v1 -= (v1 - v1)
                                        case "\n\t":
                                            fmt.Println(v1, v49)
                                        }
                                        _ = v136
                                    }
                                    _ = v135
                                } else {
                                    v49 = v49
                                }
                            }
                            v49 -= (-(v49))
                            if ((v49 / ((v49) | 1)) < v49) {
                                v140 := GenS17{F0: 321, F1: 4294967295}
                                for fi141 := 0; fi141 < 14; fi141++ {
                                    var v142 int = (^(fi53))
                                    fmt.Println(fi53, v1, v49)
                                    _ = fi141
                                    _ = v142
                                }
                                if true {
                                    for fi143 := 0; fi143 < 19; fi143++ {
                                        v144 := []int{fi53, fi143, genFn0("foo bar", 758, 2)}
                                        _ = fi143
                                        _ = v144
                                    }
                                    v3["foo bar"] = true
                                }
                                _ = v140
                            }
                        } else {
                            switch (("\n\t" + "") + "foo bar") {
                            case "\n\t":
                                fmt.Println(v49, v4, v49)
                            case "hello":
                                fmt.Println(v49, v46, v1)
                            }
                            fmt.Println(v49, fi53, v1)
                            if false {
                                v46 += (v46 & v46)
                                v4 = (-(v4))
                                v1 = -32768
                            }
                        }
                    }
                }
                _ = fi53
            }
        case "":
            if true {
                if (1 != 830) {
                    for fi145 := 0; fi145 < 8; fi145++ {
                        var v146 float32 = 0.1
                        v4 = v4
                        _ = fi145
                        _ = v146
                    }
                } else {
                    var v147 int64 = -880
                    v147 -= (v147 % (((v147 & v147)) | 1))
                    _ = v147
                }
                v148 := GenS7{F0: -2147483648}
                var v149 float64 = 0.0
                _ = v148
                _ = v149
            }
        case "foo bar":
            for fi150 := 0; fi150 < 7; fi150++ {
                if false {
                    for fi151 := 0; fi151 < 2; fi151++ {
                        fmt.Println(v46, fi151, v49)
                        _ = fi151
                    }
                } else {
                    switch fi150 {
                    case 4294967295:
                        fmt.Println(v1)
                    case 32768:
                        v152 := []int{fi150, (fi150 % ((0) | 1))}
                        _ = v152
                    default:
                        for fi153 := 0; fi153 < 9; fi153++ {
                            if true {
                                switch ((("foo bar" + "foo bar") + "a") + ("a" + ("foo bar" + "hello"))) {
                                case "hello":
                                    if (uint64(v46) < 1) {
                                        fmt.Println(v49, fi150, fi150)
                                        v154 := []int{fi150, (fi153 - 42)}
                                        v2[(^(v46))] = (42 == 42)
                                        _ = v154
                                    } else {
                                        var v155 bool = (1 == genFn1())
                                        if v155 {
                                            v156 := GenS18{F0: -9223372036854775808, F1: v46}
                                            fmt.Println(v46, v155)
                                            fmt.Println(v155, v46, v155)
                                            _ = v156
                                        } else {
                                            var v157 float64 = (0.1 - 1.5)
                                            _ = v157
                                        }
                                        v158 := []int{}
                                        _ = len(v3)
                                        var v159 string = "foo bar"
                                        switch ((v159 + "\n\t") + v159) {
                                        case "":
                                            fmt.Println(fi150, fi150)
                                        case "foo bar":
                                            switch ((v159 + v159) + ("a" + v159)) {
                                            case "foo bar":
                                                switch fi150 {
                                                case 2147483647:
                                                    fmt.Println(fi150)
                                                default:
                                                    switch v159 {
                                                    case "\n\t":
                                                        switch (v159 + v159) {
                                                        case "a":
                                                            v155 = (2 < 42)
                                                        }
                                                    case "hello":
                                                        v158 = append(v158, (-(fi150)))
                                                    }
                                                }
                                            case "hello":
                                                var v160 float64 = 3.14159
                                                _ = v160
                                            }
                                        case "\n\t":
                                            switch v159 {
                                            case "a":
                                                v158 = append(v158, genFn0(v159, 32767, 255))
                                            case "foo bar":
                                                var v161 int16 = (v1 ^ 32767)
                                                _ = v161
                                            case "":
                                                v158 = append(v158, (fi153 - fi150))
                                            }
                                        default:
                                            v1 += v1
                                        }
                                        var v162 uint = 2
                                        v4 -= 100
                                        var v163 uint = 683
                                        _ = v155
                                        _ = v158
                                        _ = v159
                                        _ = v162
                                        _ = v163
                                    }
                                case "a":
                                    v1 -= v1
                                default:
                                    switch fi153 {
                                    case 127:
                                        v164 := GenS19{F0: "hello", F1: ("" + (("" + "\n\t") + (("hello" + "\n\t") + "hello"))), F2: 42}
                                        _ = v164
                                    default:
                                        v1 = v1
                                    }
                                }
                            } else {
                                if (((float64(0.1) - 1e10) * 1.5) != ((-2.25 * 0.0) + 0.1)) {
                                    v165 := GenS4{F0: v49, F1: -51}
                                    v1 = (-(v1))
                                    _ = v165
                                }
                                v2[(v46 | v46)] = true
                            }
                            _ = fi153
                        }
                    }
                }
                _ = fi150
            }
        }
        _ = v49
    }
    v1 += (v1 >> ((1) & 63))
    for fi166 := 0; fi166 < 7; fi166++ {
        for fi167 := 0; fi167 < 17; fi167++ {
            fmt.Println(fi166, v4)
            _ = fi167
        }
        v4 -= v4
        var v168 uint32 = (v46 % (((v46 ^ 1)) | 1))
        _ = fi166
        _ = v168
    }
    var v169 uint32 = v46
    v170 := []int{}
    v46 -= 2147483647
    if false {
        if len(v170) > 0 {
            v170[0] = 128
        }
        if (!true) {
            v169 += v46
            v46 = v169
            v46 = (^(v46))
        }
    } else {
        v46 = v46
    }
    for fi171 := 0; fi171 < 11; fi171++ {
        var v172 int = (^(fi171))
        switch v172 {
        case 0:
            if false {
                v173 := GenS1{F0: 9223372036854775807}
                fmt.Println(fi171, v169, v1)
                if true {
                    switch genFn0("hello", 65535, 0) {
                    case -666:
                        var v174 uint32 = (v169 % ((42) | 1))
                        _ = v174
                    case 127:
                        switch fi171 {
                        case 32768:
                            for fi175 := 0; fi175 < 7; fi175++ {
                                switch "a" {
                                case "a":
                                    v170 = append(v170, -197)
                                default:
                                    var v176 int64 = 586
                                    _ = v176
                                }
                                _ = fi175
                            }
                        case 128:
                            var v177 uint16 = 32767
                            _ = v177
                        case -1:
                            v4 = 74
                        default:
                            if true {
                                v178 := GenS5{F0: v4, F1: (v4 * v4)}
                                _ = v178
                            } else {
                                var v179 uint = 4294967295
                                _ = v179
                            }
                        }
                    case 0:
                        fmt.Println(v169, v172)
                    default:
                        if (v172 != (fi171 + (fi171 * (^(fi171))))) {
                            if false {
                                var v180 int64 = 0
                                _ = v180
                            } else {
                                if false {
                                    if (!(true && true)) {
                                        for fi181 := 0; fi181 < 10; fi181++ {
                                            var v182 uint = 887
                                            var v183 int8 = 127
                                            if false {
                                                v184 := GenS1{F0: -9223372036854775808}
                                                if ((v169 % ((v46) | 1)) > (^(v169))) {
                                                    v185 := v2[uint32(v182)]
                                                    v182 -= v182
                                                    _ = v185
                                                } else {
                                                    v183 = (v183 | (v183 & 43))
                                                }
                                                v46 = v46
                                                v186 := GenS17{F0: 209, F1: 322}
                                                v46 -= v169
                                                var v187 int64 = -43
                                                var v188 uint64 = 4294967295
                                                switch ("\n\t" + (("hello" + ("" + "\n\t")) + "a")) {
                                                case "a":
                                                    var v189 uint8 = 255
                                                    _ = v189
                                                case "foo bar":
                                                    switch "a" {
                                                    case "":
                                                        var v190 float32 = float32(float64((float32(3.14159) * (0.1 / float32(3.14159)))))
                                                        if true {
                                                            var v191 int64 = -9223372036854775808
                                                            _ = v191
                                                        } else {
                                                            if true {
                                                                v192 := GenS19{F0: ("" + "hello"), F1: "", F2: 2}
                                                                if len(v170) > 0 {
                                                                    v170[0] = -9223372036854775808
                                                                }
                                                                v187 -= (-(v187))
                                                                _ = v192
                                                            } else {
                                                                v187 = genFn1()
                                                                var v193 int64 = -9223372036854775808
                                                                _ = v193
                                                            }
                                                            if (!(v183 == -22)) {
                                                                var v194 float64 = 0.1
                                                                v195 := GenS20{F0: (v190 == v190), F1: 127, F2: v182}
                                                                switch (^(fi171)) {
                                                                case 2147483647:
                                                                    var v196 float64 = (v194 / (((0.0 * v194) + float64(-2.25)) / v194))
                                                                    _ = v196
                                                                case 128:
                                                                    v197 := v2[638]
                                                                    _ = v197
                                                                case 9223372036854775807:
                                                                    fmt.Println(v172)
                                                                default:
                                                                    var v198 uint32 = 65535
                                                                    _ = v198
                                                                }
                                                                v199 := GenS21{F0: fi171, F1: uint64(v182), F2: v4}
                                                                _ = v194
                                                                _ = v195
                                                                _ = v199
                                                            }
                                                            if len(v170) >= 2 {
                                                                v170 = v170[1:len(v170)]
                                                            }
                                                        }
                                                        var v200 string = "foo bar"
                                                        _ = v190
                                                        _ = v200
                                                    }
                                                }
                                                _ = v184
                                                _ = v186
                                                _ = v187
                                                _ = v188
                                            }
                                            v201 := GenS10{F0: 1, F1: "a", F2: 227}
                                            v183 += (-(v4))
                                            if true {
                                                if len(v170) > 0 {
                                                    v170[0] = 32767
                                                }
                                            }
                                            _ = fi181
                                            _ = v182
                                            _ = v183
                                            _ = v201
                                        }
                                        v202 := GenS1{F0: genFn1()}
                                        _ = v202
                                    }
                                } else {
                                    for fi203 := 0; fi203 < 12; fi203++ {
                                        var v204 int64 = 1
                                        _ = fi203
                                        _ = v204
                                    }
                                    if (int8(v46) >= (-(v4))) {
                                        var v205 uint64 = 2
                                        if (v169 == v46) {
                                            v172 -= 738
                                            fmt.Println(v205, v4)
                                        } else {
                                            v172 -= fi171
                                            var v206 float32 = 3.14159
                                            _ = v206
                                        }
                                        fmt.Println(v46)
                                        _ = v205
                                    } else {
                                        var v207 uint64 = 800
                                        v170 = append(v170, (-(v172)))
                                        if false {
                                            for fi208 := 0; fi208 < 19; fi208++ {
                                                v172 += (fi208 * int(v207))
                                                var v209 uint8 = 127
                                                if (v1 != 32767) {
                                                    var v210 uint = 4294967295
                                                    var v211 float32 = 1.5
                                                    fmt.Println(v4)
                                                    _ = v210
                                                    _ = v211
                                                } else {
                                                    _ = len(v170)
                                                }
                                                _ = fi208
                                                _ = v209
                                            }
                                            v4 -= v4
                                            if true {
                                                fmt.Println(fi171)
                                                v4 = (v4 * (v4 & v4))
                                                fmt.Println(v46, v169)
                                            } else {
                                                switch (v172 - (v172 * (v172 - (v172 & genFn0("\n\t", 42, 42))))) {
                                                case 32767:
                                                    var v212 int64 = 0
                                                    _ = v212
                                                }
                                                if true {
                                                    v46 += v169
                                                    v213 := GenS18{F0: -271, F1: (v46 + (v46 % ((34) | 1)))}
                                                    v172 -= genFn0("a", uint16(v169), 1)
                                                    _ = v213
                                                } else {
                                                    fmt.Println(v172, fi171, v169)
                                                }
                                            }
                                        } else {
                                            fmt.Println(v1, fi171, v4)
                                            v3[(("" + "a") + ("foo bar" + "a"))] = false
                                            v169 += (v46 % (((^(v169))) | 1))
                                        }
                                        _ = v207
                                    }
                                }
                            }
                            v1 = (-(v1))
                        }
                    }
                } else {
                    var v214 uint8 = 42
                    _ = v214
                }
                _ = v173
            }
        default:
            v1 = v1
        }
        _ = fi171
        _ = v172
    }
    switch -9223372036854775808 {
    case 791:
        var v215 bool = true
        _ = v215
    }
    fmt.Println("v1:", v1)
    fmt.Println("v169:", v169)
    fmt.Println("slicelen v170:", len(v170), v170)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v169
    _ = v170
    _ = v2
    _ = v3
    _ = v4
    _ = v46
}
