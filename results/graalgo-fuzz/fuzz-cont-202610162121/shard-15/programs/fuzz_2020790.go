package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int64; F1 int64; F2 uint }
type GenS3 struct { F0 int64; F1 uint8; F2 int }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 int64; F1 string; F2 int }
type GenS6 struct { F0 uint8; F1 int8; F2 bool }
type GenS7 struct { F0 uint; F1 int32 }
type GenS8 struct { F0 int; F1 uint }
type GenS9 struct { F0 int; F1 uint32; F2 int64 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int64; F1 bool; F2 uint64 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint8; F1 int32 }
type GenS15 struct { F0 int16; F1 uint32 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 uint8 }
type GenS19 struct { F0 int32; F1 int; F2 string }
type GenS20 struct { F0 uint16; F1 uint32 }
type GenS21 struct { F0 uint16; F1 int; F2 uint64 }
type GenS22 struct { F0 bool; F1 string; F2 int64 }
type GenS23 struct { F0 uint16; F1 string }
type GenS24 struct { F0 int32; F1 string }
type GenS25 struct { F0 string; F1 uint32; F2 uint32 }
type GenS26 struct { F0 int16 }
type GenS27 struct { F0 int32; F1 uint16; F2 uint }
type GenS28 struct { F0 bool; F1 uint64; F2 int8 }
type GenS29 struct { F0 int; F1 int64; F2 int }
type GenS30 struct { F0 int; F1 int64 }
func genFn0() int {
    v4 := GenS1{F0: 1, F1: false}
    var v5 int16 = 1000
    if (false && (!(0 != 784))) {
        v6 := GenS1{F0: 9223372036854775807, F1: (!false)}
        var v7 int8 = -128
        _ = v6
        _ = v7
    } else {
        var v8 float32 = 1e10
        _ = v8
    }
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int, p1 int32) int {
    fmt.Println(p1, p0)
    for fi9 := 0; fi9 < 15; fi9++ {
        p1 += 42
        v10 := GenS1{F0: -372, F1: false}
        for fi11 := 0; fi11 < 2; fi11++ {
            p1 -= p1
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    switch 0 {
    case -297:
        switch p0 {
        case 0:
            v12 := GenS1{F0: 42, F1: false}
            _ = v12
        }
    case 1:
        var v13 float64 = 3.14159
        _ = v13
    case 128:
        v14 := []int{}
        _ = v14
    }
    p1 = int32(p0)
    fmt.Println(p1)
    _ = p0
    _ = p1
    return (^(p0))
}

func main() {
    var v1 uint = 255
    var v2 int64 = -182
    var v3 float32 = -2.25
    v2 += v2
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := []int{32768, genFn1((fi15 / ((genFn1(9223372036854775807, -2147483648)) | 1)), 42)}
        v16 = append(v16, fi15)
        _ = fi15
        _ = v16
    }
    fmt.Println(v2, v1, v1)
    v2 = (v2 % (((v2 / ((9223372036854775807) | 1))) | 1))
    v17 := GenS1{F0: v2, F1: true}
    var v18 int64 = (-(v2))
    switch 32768 {
    case -1:
        v19 := []int{}
        _ = v19
    case 32768:
        v20 := []int{-331, 688}
        _ = v20
    case 127:
        if ((4294967295 == 545) && (1 > 0)) {
            var v21 uint = 2
            fmt.Println(v21, v2, v3)
            _ = v21
        }
    }
    var v22 uint8 = 255
    for fi23 := 0; fi23 < 20; fi23++ {
        if ((v1 >> ((20) & 63)) == 18446744073709551615) {
            switch (fi23 % (((^(fi23))) | 1)) {
            case 1:
                for fi24 := 0; fi24 < 8; fi24++ {
                    if true {
                        switch "a" {
                        case "":
                            var v25 uint8 = (^(v22))
                            _ = v25
                        default:
                            if true {
                                var v26 int32 = -718
                                _ = v26
                            } else {
                                v27 := GenS2{F0: (v18 - (v18 - (-(v18)))), F1: (v18 ^ v2), F2: v1}
                                _ = v27
                            }
                            if (1 >= v1) {
                                v3 += ((3.14159 * v3) * 0.0)
                            } else {
                                v28 := []int{(fi24 + (fi24 * (fi23 ^ fi24))), (fi24 % (((fi23 - (fi23 >> ((37) & 63)))) | 1))}
                                var v29 uint16 = uint16(v22)
                                v30 := GenS1{F0: v2, F1: (v22 >= uint8(v29))}
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                        }
                        v31 := GenS3{F0: v18, F1: (v22 >> ((v1) & 63)), F2: -887}
                        _ = v31
                    } else {
                        v18 = (v2 << ((1) & 63))
                        if (42 != 32767) {
                            fmt.Println(v18, fi23, v3)
                            v22 = (v22 % ((v22) | 1))
                            fmt.Println(fi23, v1)
                        } else {
                            switch 128 {
                            case 32767:
                                if (25 != 100) {
                                    v32 := GenS1{F0: -378, F1: false}
                                    if (!(462 > 1)) {
                                        if false {
                                            if false {
                                                v33 := []int{2147483647, (fi23 << ((v1) & 63))}
                                                v22 = v22
                                                _ = v33
                                            }
                                            if (v1 > uint(v22)) {
                                                v34 := []int{(fi23 % (((fi24 << ((v1) & 63))) | 1)), (fi24 / ((fi24) | 1))}
                                                v35 := GenS3{F0: v2, F1: (^(v22)), F2: int(v22)}
                                                _ = v34
                                                _ = v35
                                            } else {
                                                v36 := GenS4{F0: v22, F1: 65535}
                                                if true {
                                                    var v37 float64 = 0.0
                                                    _ = v37
                                                }
                                                v22 = (v22 / (((v22 - (v22 ^ (v22 / ((v22) | 1))))) | 1))
                                                switch ("hello" + "") {
                                                case "hello":
                                                    v38 := []int{}
                                                    _ = v38
                                                case "":
                                                    if true {
                                                        var v39 int16 = 1
                                                        _ = v39
                                                    }
                                                default:
                                                    if (1 < (^(v1))) {
                                                        if true {
                                                            fmt.Println(v3)
                                                            if (!(!false)) {
                                                                if true {
                                                                    v22 -= 127
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        var v40 int8 = 0
                                                                        _ = v40
                                                                    case "foo bar":
                                                                        v41 := []int{}
                                                                        _ = v41
                                                                    default:
                                                                        switch ("a" + ("" + "")) {
                                                                        case "hello":
                                                                            v3 += 3.14159
                                                                        case "\n\t":
                                                                            v42 := []int{(fi23 << ((v1) & 63)), fi23}
                                                                            _ = v42
                                                                        case "foo bar":
                                                                            v43 := []int{}
                                                                            _ = v43
                                                                        default:
                                                                            switch 127 {
                                                                            case 2147483647:
                                                                                var v44 int8 = 100
                                                                                _ = v44
                                                                            case -1:
                                                                                v45 := GenS5{F0: v2, F1: "foo bar", F2: -2147483648}
                                                                                _ = v45
                                                                            }
                                                                        }
                                                                    }
                                                                    if (166 == uint64(v1)) {
                                                                        v2 = v2
                                                                        v46 := GenS6{F0: v22, F1: 31, F2: true}
                                                                        if (((0 <= 2147483647) || false) && true) {
                                                                            if true {
                                                                                v47 := []int{}
                                                                                v22 = 127
                                                                                _ = v47
                                                                            } else {
                                                                                v48 := GenS3{F0: (v2 + 918), F1: (v22 * 55), F2: 32767}
                                                                                _ = v48
                                                                            }
                                                                        }
                                                                        _ = v46
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            var v49 int = genFn1(genFn1((fi24 | fi24), 2147483647), 0)
                                                            fmt.Println(v2)
                                                            _ = v49
                                                        }
                                                    } else {
                                                        var v50 bool = false
                                                        var v51 float32 = (v3 - v3)
                                                        var v52 bool = ((1.5 + -2.25) != 1e10)
                                                        _ = v50
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                }
                                                var v53 uint32 = 1
                                                var v54 int64 = v18
                                                switch ("foo bar" + ("hello" + ("foo bar" + "foo bar"))) {
                                                case "\n\t":
                                                    v55 := []int{(fi24 >> (((v1 + (v1 ^ v1))) & 63)), (fi24 >> ((v1) & 63))}
                                                    _ = v55
                                                case "":
                                                    v22 = (^(v22))
                                                }
                                                _ = v36
                                                _ = v53
                                                _ = v54
                                            }
                                            v56 := GenS6{F0: (v22 + v22), F1: 42, F2: true}
                                            v57 := GenS7{F0: uint(v22), F1: 0}
                                            _ = v56
                                            _ = v57
                                        }
                                        switch (fi24 >> (((^(v1))) & 63)) {
                                        case 9223372036854775807:
                                            v58 := GenS6{F0: v22, F1: 1, F2: (!(!(779 >= 42)))}
                                            _ = v58
                                        case 128:
                                            v18 -= v18
                                        }
                                    } else {
                                        v59 := GenS8{F0: fi24, F1: (v1 % ((v1) | 1))}
                                        if false {
                                            if true {
                                                var v60 int = (fi24 >> (((v1 * (v1 - v1))) & 63))
                                                if true {
                                                    v60 -= -9223372036854775808
                                                    if false {
                                                        var v61 string = "hello"
                                                        v22 += v22
                                                        _ = v61
                                                    } else {
                                                        var v62 float64 = 0.1
                                                        _ = v62
                                                    }
                                                }
                                                fmt.Println(v1)
                                                _ = v60
                                            } else {
                                                if ("foo bar" >= "\n\t") {
                                                    v22 -= (v22 / ((127) | 1))
                                                    v63 := GenS4{F0: v22, F1: 65535}
                                                    v64 := GenS7{F0: 2, F1: 42}
                                                    _ = v63
                                                    _ = v64
                                                } else {
                                                    var v65 float64 = ((1e10 + (3.14159 + 0.1)) - 1.5)
                                                    if (!(452 != 42)) {
                                                        v66 := GenS8{F0: (fi24 / ((-781) | 1)), F1: v1}
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            var v67 uint32 = 86
                                                            _ = v67
                                                        case "hello":
                                                            v1 -= (v1 - v1)
                                                        default:
                                                            v68 := GenS4{F0: v22, F1: 488}
                                                            _ = v68
                                                        }
                                                        _ = v66
                                                    } else {
                                                        var v69 uint8 = 27
                                                        v70 := []int{(-(fi23))}
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                    _ = v65
                                                }
                                                var v71 int64 = v2
                                                _ = v71
                                            }
                                        }
                                        _ = v59
                                    }
                                    v18 = v2
                                    _ = v32
                                } else {
                                    v72 := GenS8{F0: (fi23 - genFn0()), F1: (v1 << ((42) & 63))}
                                    v3 += 0.1
                                    fmt.Println(v1)
                                    _ = v72
                                }
                            }
                            var v73 bool = false
                            _ = v73
                        }
                        v3 -= (-2.25 * float32((1.5 - 0.0)))
                    }
                    _ = fi24
                }
            case 127:
                fmt.Println(v3)
            default:
                v1 = (^(v1))
            }
        } else {
            v74 := GenS9{F0: 9223372036854775807, F1: 987, F2: v18}
            _ = v74
        }
        _ = fi23
    }
    for fi75 := 0; fi75 < 7; fi75++ {
        if (!false) {
            fmt.Println(v22, v1, v3)
        }
        _ = fi75
    }
    for fi76 := 0; fi76 < 5; fi76++ {
        for fi77 := 0; fi77 < 10; fi77++ {
            v3 -= v3
            var v78 float32 = v3
            if true {
                v79 := GenS10{F0: "a"}
                if true {
                    var v80 bool = (((v78 * (float32(float64(float32(0.1))) / v3)) + v3) != (0.0 / (v3 * v78)))
                    v81 := GenS11{F0: (v18 - (v18 & 9223372036854775807)), F1: ((-(v18)) <= (v18 - v2)), F2: 4294967295}
                    _ = v80
                    _ = v81
                } else {
                    v2 += (^(v2))
                    v22 -= 171
                    if (42 > 42) {
                        var v82 uint16 = 158
                        _ = v82
                    } else {
                        var v83 uint = (^(v1))
                        _ = v83
                    }
                }
                _ = v79
            }
            _ = fi77
            _ = v78
        }
        for fi84 := 0; fi84 < 18; fi84++ {
            if (false && (100 > 100)) {
                fmt.Println(fi76, v1, v18)
                fmt.Println(fi76, v18, v18)
            } else {
                v85 := []int{(fi84 + fi84)}
                _ = v85
            }
            _ = fi84
        }
        for fi86 := 0; fi86 < 15; fi86++ {
            switch "a" {
            case "\n\t":
                var v87 uint8 = v22
                _ = v87
            case "":
                var v88 int64 = v18
                _ = v88
            case "hello":
                v89 := GenS10{F0: ("hello" + ("\n\t" + "a"))}
                _ = v89
            }
            _ = fi86
        }
        _ = fi76
    }
    if (("a" + "a") < "foo bar") {
        v90 := []int{}
        _ = v90
    } else {
        var v91 int = 0
        fmt.Println(v3, v1, v3)
        switch int(v22) {
        case -9223372036854775808:
            switch ("\n\t" + ("foo bar" + "a")) {
            case "\n\t":
                if false {
                    var v92 bool = true
                    switch ("hello" + ("hello" + "foo bar")) {
                    case "\n\t":
                        v18 = v18
                    case "":
                        v1 -= (^(v1))
                    case "hello":
                        v93 := []int{v91, 9223372036854775807, (-(v91)), (v91 ^ -2147483648)}
                        _ = v93
                    default:
                        if (!(genFn0() > -9223372036854775808)) {
                            v2 -= v2
                            if v92 {
                                if true {
                                    v94 := GenS12{F0: 18446744073709551615}
                                    _ = v94
                                }
                                var v95 float32 = v3
                                _ = v95
                            } else {
                                if (2 < (v22 % (((v22 ^ v22)) | 1))) {
                                    v96 := []int{v91}
                                    for fi97 := 0; fi97 < 20; fi97++ {
                                        for fi98 := 0; fi98 < 15; fi98++ {
                                            if (!(3.14159 >= 3.14159)) {
                                                var v99 int32 = 2147483647
                                                fmt.Println(v92, fi98, fi98)
                                                var v100 uint64 = 2
                                                fmt.Println(v91)
                                                _ = v100
                                                _ = v99
                                            } else {
                                                _ = len(v96)
                                                v22 = v22
                                                v3 -= 0.0
                                            }
                                            switch fi97 {
                                            case 1:
                                                var v101 uint = v1
                                                _ = v101
                                            case -126:
                                                if (("hello" + "\n\t") < ("hello" + "")) {
                                                    var v102 string = "\n\t"
                                                    v22 = (^(v22))
                                                    if ((float64((v3 * (v3 * v3))) - -2.25) <= (0.1 + float64(float32((3.14159 + 1.5))))) {
                                                        v96 = append(v96, fi97)
                                                        switch (fi97 % (((v91 / ((32767) | 1))) | 1)) {
                                                        case 128:
                                                            v103 := GenS13{F0: (v22 << ((255) & 63))}
                                                            _ = v103
                                                        default:
                                                            v104 := GenS14{F0: v22, F1: -977}
                                                            _ = v104
                                                        }
                                                        _ = len(v96)
                                                        v92 = (!((v22 * v22) >= 2))
                                                        var v105 float32 = 0.1
                                                        if (false && ((406 >= -32768) || true)) {
                                                            if (!(true && (!(1.5 > 1.5)))) {
                                                                if false {
                                                                    v106 := GenS14{F0: (v22 >> ((v1) & 63)), F1: 1}
                                                                    _ = v106
                                                                } else {
                                                                    var v107 uint8 = (v22 % ((1) | 1))
                                                                    _ = v107
                                                                }
                                                            } else {
                                                                var v108 int16 = 610
                                                                _ = v108
                                                            }
                                                            if len(v96) > 0 {
                                                                v96[0] = (v91 + fi97)
                                                            }
                                                        } else {
                                                            switch genFn0() {
                                                            case -427:
                                                                _ = len(v96)
                                                            case -2147483648:
                                                                var v109 int64 = v18
                                                                _ = v109
                                                            case -876:
                                                                var v110 int32 = 1
                                                                _ = v110
                                                            default:
                                                                v111 := GenS1{F0: (v18 >> ((v1) & 63)), F1: (255 >= v1)}
                                                                _ = v111
                                                            }
                                                            fmt.Println(v105, v92, v2)
                                                            if true {
                                                                v92 = true
                                                            }
                                                        }
                                                        v112 := GenS15{F0: 42, F1: uint32(v22)}
                                                        _ = v105
                                                        _ = v112
                                                    }
                                                    var v113 uint32 = 268
                                                    _ = v102
                                                    _ = v113
                                                } else {
                                                    switch genFn0() {
                                                    case -274:
                                                        v91 -= fi98
                                                    default:
                                                        switch (-(fi98)) {
                                                        case 40:
                                                            var v114 float64 = 0.1
                                                            _ = v114
                                                        case -622:
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v115 int16 = 1
                                                                _ = v115
                                                            default:
                                                                switch (fi98 % (((fi97 & -801)) | 1)) {
                                                                case 9223372036854775807:
                                                                    v92 = (false || true)
                                                                case 32767:
                                                                    if false {
                                                                        if (0.1 <= 1e10) {
                                                                            var v116 uint = (v1 | (v1 * v1))
                                                                            v2 = (-(v2))
                                                                            _ = v116
                                                                        } else {
                                                                            if false {
                                                                                fmt.Println(v18)
                                                                                _ = len(v96)
                                                                            }
                                                                            _ = len(v96)
                                                                            var v117 uint = v1
                                                                            _ = v117
                                                                        }
                                                                        if (float32(float64((v3 * 3.14159))) < v3) {
                                                                            v18 = (^(v18))
                                                                            if false {
                                                                                if (-2147483648 > (fi97 * fi97)) {
                                                                                    v118 := GenS15{F0: 1, F1: 161}
                                                                                    var v119 bool = true
                                                                                    _ = v118
                                                                                    _ = v119
                                                                                } else {
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        v120 := GenS5{F0: (v18 + (v18 >> ((v1) & 63))), F1: "", F2: fi98}
                                                                                        _ = v120
                                                                                    case "a":
                                                                                        if (32767 != 65535) {
                                                                                            v3 += v3
                                                                                            if (-2.25 == v3) {
                                                                                                if ((0.0 * (0.0 + 1e10)) >= float64(v3)) {
                                                                                                    v121 := GenS8{F0: (fi97 | fi97), F1: v1}
                                                                                                    var v122 int32 = 42
                                                                                                    _ = v121
                                                                                                    _ = v122
                                                                                                }
                                                                                                switch "\n\t" {
                                                                                                case "a":
                                                                                                    v123 := GenS16{F0: fi97}
                                                                                                    _ = v123
                                                                                                }
                                                                                            }
                                                                                            var v124 float32 = v3
                                                                                            _ = v124
                                                                                        }
                                                                                    case "foo bar":
                                                                                        fmt.Println(v2, v3, fi98)
                                                                                    }
                                                                                }
                                                                                v91 = (-(fi97))
                                                                                v2 = v18
                                                                            } else {
                                                                                if len(v96) >= 2 {
                                                                                    v96 = v96[1:len(v96)]
                                                                                }
                                                                                v125 := GenS10{F0: "hello"}
                                                                                _ = v125
                                                                            }
                                                                            switch -885 {
                                                                            case 32767:
                                                                                fmt.Println(v2, fi98)
                                                                                switch ("a" + "") {
                                                                                case "\n\t":
                                                                                    if len(v96) >= 2 {
                                                                                        v96 = v96[1:len(v96)]
                                                                                    }
                                                                                case "":
                                                                                    if len(v96) > 0 {
                                                                                        v96[0] = 1
                                                                                    }
                                                                                default:
                                                                                    if v92 {
                                                                                        _ = len(v96)
                                                                                        if v92 {
                                                                                            if (false && v92) {
                                                                                                var v126 float64 = -2.25
                                                                                                var v127 float64 = ((float64(v3) + 0.1) - (1e10 + -2.25))
                                                                                                fmt.Println(v92, v92)
                                                                                                _ = v126
                                                                                                _ = v127
                                                                                            }
                                                                                            v92 = (v92 || v92)
                                                                                            var v128 int8 = 0
                                                                                            _ = v128
                                                                                        }
                                                                                    } else {
                                                                                        v92 = (113 < (v22 + (v22 - (v22 + 42))))
                                                                                        switch fi97 {
                                                                                        case 42:
                                                                                            if false {
                                                                                                var v129 int16 = 0
                                                                                                _ = v129
                                                                                            }
                                                                                        case 609:
                                                                                            v3 = v3
                                                                                        default:
                                                                                            v96 = append(v96, fi97)
                                                                                            v92 = ((("foo bar" + "a") + ("foo bar" + "\n\t")) <= ("foo bar" + "foo bar"))
                                                                                            _ = len(v96)
                                                                                        }
                                                                                        if len(v96) > 0 {
                                                                                            v96[0] = (fi97 / ((genFn1(fi98, 0)) | 1))
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if false {
                                                                                    fmt.Println(v91, v91, v1)
                                                                                    if v92 {
                                                                                        v1 += 441
                                                                                    } else {
                                                                                        var v130 int8 = 9
                                                                                        _ = v130
                                                                                    }
                                                                                } else {
                                                                                    var v131 bool = v92
                                                                                    _ = v131
                                                                                }
                                                                            case 127:
                                                                                fmt.Println(fi98)
                                                                            default:
                                                                                if len(v96) >= 2 {
                                                                                    v96 = v96[1:len(v96)]
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v132 uint32 = 42
                                                                            _ = v132
                                                                        }
                                                                    } else {
                                                                        fmt.Println(fi97, v92)
                                                                        if v92 {
                                                                            if v92 {
                                                                                var v133 int8 = -1
                                                                                if v92 {
                                                                                    if true {
                                                                                        if len(v96) >= 2 {
                                                                                            v96 = v96[1:len(v96)]
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    var v134 int16 = 42
                                                                                    v135 := GenS16{F0: (fi98 + (-(fi97)))}
                                                                                    var v136 int = v91
                                                                                    v137 := GenS17{F0: 728}
                                                                                    _ = v134
                                                                                    _ = v135
                                                                                    _ = v136
                                                                                    _ = v137
                                                                                }
                                                                                _ = len(v96)
                                                                                _ = v133
                                                                            }
                                                                            var v138 int32 = 42
                                                                            _ = v138
                                                                        } else {
                                                                            fmt.Println(fi97, v22)
                                                                        }
                                                                    }
                                                                default:
                                                                    v91 = (fi97 << ((v1) & 63))
                                                                }
                                                            }
                                                            v139 := GenS11{F0: (v2 ^ 42), F1: (v22 == v22), F2: 2}
                                                            v3 = (((v3 + 0.0) / v3) / (v3 / ((0.1 + ((0.0 - v3) - v3)) - float32(-2.25))))
                                                            _ = v139
                                                        default:
                                                            var v140 int8 = -1
                                                            _ = v140
                                                        }
                                                    }
                                                }
                                                if (4294967295 > 18446744073709551615) {
                                                    v91 = fi97
                                                }
                                                switch (fi98 - genFn0()) {
                                                case 32767:
                                                    v141 := GenS12{F0: 18446744073709551615}
                                                    _ = v141
                                                case 519:
                                                    if v92 {
                                                        v142 := GenS18{F0: (v22 * v22)}
                                                        var v143 int = (v91 >> ((v1) & 63))
                                                        _ = v142
                                                        _ = v143
                                                    }
                                                    if (1e10 < 3.14159) {
                                                        fmt.Println(v18)
                                                        var v144 int64 = -1
                                                        _ = v144
                                                    } else {
                                                        if (1 != 0) {
                                                            if len(v96) >= 2 {
                                                                v96 = v96[1:len(v96)]
                                                            }
                                                            switch (fi97 * (fi97 & 69)) {
                                                            case -144:
                                                                v145 := GenS19{F0: -1, F1: 127, F2: ""}
                                                                _ = v145
                                                            case 823:
                                                                if len(v96) >= 2 {
                                                                    v96 = v96[1:len(v96)]
                                                                }
                                                            default:
                                                                fmt.Println(v22)
                                                            }
                                                        } else {
                                                            if false {
                                                                v1 -= 20
                                                            } else {
                                                                if (v3 > ((1.5 - 3.14159) - -2.25)) {
                                                                    if v92 {
                                                                        switch 4294967295 {
                                                                        case 42:
                                                                            _ = len(v96)
                                                                        case 1:
                                                                            var v146 string = ""
                                                                            _ = v146
                                                                        default:
                                                                            if (0.0 != v3) {
                                                                                v18 = v2
                                                                                switch (("hello" + ("" + "a")) + "hello") {
                                                                                case "foo bar":
                                                                                    fmt.Println(fi98, v22, v1)
                                                                                case "":
                                                                                    v147 := GenS20{F0: 468, F1: 42}
                                                                                    _ = v147
                                                                                }
                                                                                v22 -= 84
                                                                                fmt.Println(v2, fi98, v1)
                                                                                v22 = v22
                                                                                if v92 {
                                                                                    if (0 != 477) {
                                                                                        if v92 {
                                                                                            v148 := GenS5{F0: v2, F1: "a", F2: (v91 << ((v1) & 63))}
                                                                                            if len(v96) > 0 {
                                                                                                v96[0] = 32767
                                                                                            }
                                                                                            _ = v148
                                                                                        }
                                                                                        var v149 float32 = v3
                                                                                        switch (v91 >> (((v1 & v1)) & 63)) {
                                                                                        case 127:
                                                                                            if (!(v92 && (42 == (v22 << ((v1) & 63))))) {
                                                                                                if len(v96) >= 2 {
                                                                                                    v96 = v96[1:len(v96)]
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(v18)
                                                                                                if (112 != 42) {
                                                                                                    switch (("hello" + "a") + "foo bar") {
                                                                                                    case "a":
                                                                                                        switch 899 {
                                                                                                        case 32767:
                                                                                                            v91 -= genFn1(fi98, 1)
                                                                                                        case 2147483647:
                                                                                                            var v150 uint = v1
                                                                                                            _ = v150
                                                                                                        case 4294967295:
                                                                                                            v2 = v18
                                                                                                        }
                                                                                                    default:
                                                                                                        v3 = (v149 * float32((3.14159 / 3.14159)))
                                                                                                    }
                                                                                                    if (!v92) {
                                                                                                        v18 = (v18 * (v18 / ((v2) | 1)))
                                                                                                        switch (^(fi98)) {
                                                                                                        case 128:
                                                                                                            fmt.Println(v18, v22)
                                                                                                        case -28:
                                                                                                            var v151 float64 = 0.0
                                                                                                            _ = v151
                                                                                                        case 0:
                                                                                                            switch fi98 {
                                                                                                            case 32767:
                                                                                                                v91 += int(v18)
                                                                                                            case -725:
                                                                                                                v91 = (fi98 & 127)
                                                                                                            case 128:
                                                                                                                if (42 == 2147483647) {
                                                                                                                    v152 := GenS1{F0: 42, F1: v92}
                                                                                                                    var v153 float64 = 3.14159
                                                                                                                    switch (v91 & genFn1(v91, -195)) {
                                                                                                                    case -161:
                                                                                                                        var v154 uint8 = v22
                                                                                                                        _ = v154
                                                                                                                    case 127:
                                                                                                                        var v155 int32 = -36
                                                                                                                        _ = v155
                                                                                                                    case 364:
                                                                                                                        if len(v96) > 0 {
                                                                                                                            v96[0] = (v91 / ((-9223372036854775808) | 1))
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v152
                                                                                                                    _ = v153
                                                                                                                } else {
                                                                                                                    var v156 int = genFn1((fi98 * (fi98 | (fi97 - (v91 >> (((v1 - 42)) & 63))))), 843)
                                                                                                                    _ = v156
                                                                                                                }
                                                                                                            default:
                                                                                                                fmt.Println(v92, v149, fi98)
                                                                                                            }
                                                                                                        }
                                                                                                        if len(v96) > 0 {
                                                                                                            v96[0] = (fi98 + 128)
                                                                                                        }
                                                                                                    } else {
                                                                                                        v92 = (!v92)
                                                                                                        _ = len(v96)
                                                                                                        v22 -= 42
                                                                                                    }
                                                                                                }
                                                                                                var v157 uint32 = uint32(v22)
                                                                                                _ = v157
                                                                                            }
                                                                                        case 9223372036854775807:
                                                                                            v1 = v1
                                                                                        case -524:
                                                                                            var v158 int16 = 0
                                                                                            _ = v158
                                                                                        }
                                                                                        v2 -= (v18 * (v18 + v2))
                                                                                        var v159 int = (v91 ^ fi97)
                                                                                        _ = v149
                                                                                        _ = v159
                                                                                    } else {
                                                                                        v18 += (v2 >> (((v1 & v1)) & 63))
                                                                                    }
                                                                                    var v160 int32 = -2147483648
                                                                                    _ = v160
                                                                                }
                                                                                if v92 {
                                                                                    if len(v96) >= 2 {
                                                                                        v96 = v96[1:len(v96)]
                                                                                    }
                                                                                    var v161 int16 = 0
                                                                                    _ = v161
                                                                                } else {
                                                                                    fmt.Println(v18, v3, v2)
                                                                                    if len(v96) >= 2 {
                                                                                        v96 = v96[1:len(v96)]
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            v91 = fi97
                                                                            switch genFn1(genFn0(), 1) {
                                                                            case 128:
                                                                                fmt.Println(v18, fi97)
                                                                            }
                                                                        } else {
                                                                            v162 := GenS16{F0: (fi98 ^ (fi98 % (((fi97 >> ((258) & 63))) | 1)))}
                                                                            v163 := GenS16{F0: fi97}
                                                                            _ = v162
                                                                            _ = v163
                                                                        }
                                                                        if len(v96) >= 2 {
                                                                            v96 = v96[1:len(v96)]
                                                                        }
                                                                        if v92 {
                                                                            if (!(v3 >= v3)) {
                                                                                v96 = append(v96, v91)
                                                                            } else {
                                                                                if (int32(fi98) <= 2147483647) {
                                                                                    v164 := GenS4{F0: (^(v22)), F1: 65535}
                                                                                    fmt.Println(v92)
                                                                                    v92 = (127 > -1)
                                                                                    _ = v164
                                                                                }
                                                                            }
                                                                            v165 := GenS21{F0: 42, F1: (fi98 + fi98), F2: 18446744073709551615}
                                                                            v166 := GenS22{F0: v92, F1: ("a" + "foo bar"), F2: v2}
                                                                            _ = v165
                                                                            _ = v166
                                                                        }
                                                                    }
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        if len(v96) >= 2 {
                                                                            v96 = v96[1:len(v96)]
                                                                        }
                                                                    case "a":
                                                                        v91 = fi97
                                                                    }
                                                                    v167 := GenS23{F0: uint16(v22), F1: ("hello" + "\n\t")}
                                                                    _ = v167
                                                                } else {
                                                                    var v168 uint64 = uint64(v1)
                                                                    _ = v168
                                                                }
                                                                switch "" {
                                                                case "a":
                                                                    if v92 {
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            var v169 int16 = -1
                                                                            var v170 string = ("foo bar" + "\n\t")
                                                                            fmt.Println(fi98, fi97)
                                                                            _ = v169
                                                                            _ = v170
                                                                        }
                                                                    }
                                                                default:
                                                                    v91 -= 9223372036854775807
                                                                }
                                                                v96 = append(v96, fi97)
                                                            }
                                                            if v92 {
                                                                fmt.Println(fi97, fi98, v91)
                                                                if true {
                                                                    if false {
                                                                        v92 = v92
                                                                        var v171 uint = (v1 / ((v1) | 1))
                                                                        _ = v171
                                                                    }
                                                                }
                                                                var v172 uint8 = (v22 ^ 127)
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    _ = len(v96)
                                                                default:
                                                                    var v173 int64 = -1
                                                                    _ = v173
                                                                }
                                                                _ = v172
                                                            } else {
                                                                switch (^(v91)) {
                                                                case -2147483648:
                                                                    var v174 int16 = 32767
                                                                    _ = v174
                                                                case 128:
                                                                    if v92 {
                                                                        v175 := GenS24{F0: -2147483648, F1: ("" + "a")}
                                                                        v92 = true
                                                                        _ = v175
                                                                    } else {
                                                                        var v176 uint64 = 42
                                                                        v177 := GenS22{F0: (!false), F1: "a", F2: (v18 % (((-(v2))) | 1))}
                                                                        v178 := GenS24{F0: 0, F1: "\n\t"}
                                                                        _ = v176
                                                                        _ = v177
                                                                        _ = v178
                                                                    }
                                                                }
                                                                v96 = append(v96, 219)
                                                                fmt.Println(fi97)
                                                                v96 = append(v96, (v91 | fi98))
                                                                v179 := GenS25{F0: "foo bar", F1: 65535, F2: 218}
                                                                fmt.Println(v1, v1, fi97)
                                                                _ = v179
                                                            }
                                                        }
                                                    }
                                                }
                                            case 784:
                                                var v180 bool = v92
                                                _ = v180
                                            default:
                                                if v92 {
                                                    switch fi97 {
                                                    case 127:
                                                        _ = len(v96)
                                                    case 9223372036854775807:
                                                        var v181 uint8 = v22
                                                        _ = v181
                                                    case 42:
                                                        var v182 uint16 = 2
                                                        _ = v182
                                                    }
                                                    var v183 int = fi98
                                                    if (112 != 1) {
                                                        if v92 {
                                                            fmt.Println(fi98)
                                                            switch (fi98 & fi98) {
                                                            case 128:
                                                                v184 := GenS8{F0: 951, F1: v1}
                                                                _ = v184
                                                            default:
                                                                v185 := GenS21{F0: 813, F1: (^(fi97)), F2: 42}
                                                                _ = v185
                                                            }
                                                            if v92 {
                                                                var v186 uint8 = v22
                                                                _ = v186
                                                            } else {
                                                                fmt.Println(v2, v22, v183)
                                                            }
                                                        } else {
                                                            if len(v96) >= 2 {
                                                                v96 = v96[1:len(v96)]
                                                            }
                                                            v187 := GenS26{F0: -32768}
                                                            _ = v187
                                                        }
                                                    }
                                                    var v188 uint32 = 2147483647
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        var v189 bool = true
                                                        _ = v189
                                                    case "":
                                                        if len(v96) > 0 {
                                                            v96[0] = fi97
                                                        }
                                                    case "a":
                                                        v22 += v22
                                                    }
                                                    _ = v183
                                                    _ = v188
                                                }
                                            }
                                            _ = fi98
                                        }
                                        v92 = (v92 || (!v92))
                                        var v190 uint8 = 2
                                        _ = fi97
                                        _ = v190
                                    }
                                    _ = v96
                                }
                                v191 := GenS27{F0: 0, F1: 1, F2: (v1 / (((v1 / ((0) | 1))) | 1))}
                                v18 = -9223372036854775808
                                _ = v191
                            }
                        }
                    }
                    if v92 {
                        if (((-789 >= 1000) && v92) || (v1 <= v1)) {
                            v18 = 56
                            if v92 {
                                v192 := GenS28{F0: (0.0 >= -2.25), F1: 18446744073709551615, F2: 127}
                                var v193 uint = (v1 / ((v1) | 1))
                                if (32767 < 1000) {
                                    for fi194 := 0; fi194 < 4; fi194++ {
                                        if v92 {
                                            var v195 int8 = 91
                                            v2 += v2
                                            var v196 int16 = 32767
                                            _ = v195
                                            _ = v196
                                        }
                                        _ = fi194
                                    }
                                    var v197 uint32 = 188
                                    v91 -= (v91 & genFn1(genFn1((v91 & v91), 960), 0))
                                    _ = v197
                                } else {
                                    v198 := []int{genFn1(42, -1), (v91 >> (((v1 % (((^(v193))) | 1))) & 63)), (v91 ^ (v91 + v91))}
                                    _ = v198
                                }
                                _ = v192
                                _ = v193
                            } else {
                                v199 := []int{(^(v91)), v91, -9223372036854775808}
                                v200 := GenS21{F0: 255, F1: v91, F2: uint64(v1)}
                                if (false && (v92 || v92)) {
                                    var v201 string = "hello"
                                    _ = v201
                                } else {
                                    if (("\n\t" + "a") > ("\n\t" + "hello")) {
                                        fmt.Println(v2, v2)
                                        var v202 string = "a"
                                        _ = v202
                                    } else {
                                        if len(v199) > 0 {
                                            v199[0] = (v91 & 496)
                                        }
                                    }
                                    switch ("\n\t" + "") {
                                    case "hello":
                                        switch 350 {
                                        case 32767:
                                            fmt.Println(v1, v3, v18)
                                        }
                                    case "\n\t":
                                        switch (v91 & genFn1(v91, 42)) {
                                        case -1:
                                            var v203 int32 = 0
                                            _ = v203
                                        case 42:
                                            switch v91 {
                                            case 0:
                                                switch (("a" + ("hello" + ("\n\t" + ""))) + "") {
                                                case "a":
                                                    if len(v199) >= 2 {
                                                        v199 = v199[1:len(v199)]
                                                    }
                                                case "\n\t":
                                                    var v204 int = v91
                                                    _ = v204
                                                case "foo bar":
                                                    v205 := GenS5{F0: (v2 + (v18 | (v2 & (v18 % ((v2) | 1))))), F1: ("foo bar" + ""), F2: v91}
                                                    _ = v205
                                                }
                                            case 128:
                                                var v206 int8 = 1
                                                _ = v206
                                            case 42:
                                                v91 -= 1
                                            }
                                        default:
                                            switch 0 {
                                            case 2147483647:
                                                v207 := GenS29{F0: (-(v91)), F1: v18, F2: v91}
                                                _ = v207
                                            case -145:
                                                var v208 int8 = -78
                                                _ = v208
                                            }
                                        }
                                    }
                                }
                                _ = v199
                                _ = v200
                            }
                        } else {
                            var v209 uint64 = 51
                            switch "hello" {
                            case "hello":
                                if v92 {
                                    switch v91 {
                                    case -161:
                                        v210 := []int{genFn1((v91 >> (((v1 >> (((v1 + v1)) & 63))) & 63)), 42), genFn0(), v91}
                                        _ = v210
                                    }
                                }
                            case "a":
                                var v211 bool = (1 >= 1)
                                _ = v211
                            default:
                                v1 -= v1
                            }
                            v212 := []int{(v91 - 1)}
                            _ = v209
                            _ = v212
                        }
                    }
                    _ = v92
                }
            }
        case 0:
            switch (v91 >> ((v1) & 63)) {
            case 42:
                var v213 float32 = v3
                _ = v213
            case 656:
                for fi214 := 0; fi214 < 20; fi214++ {
                    if (true || true) {
                        v3 += 1e10
                    }
                    _ = fi214
                }
            }
        case 1:
            switch v91 {
            case 32767:
                fmt.Println(v22, v1, v22)
            case 4294967295:
                if (127 < 0) {
                    v215 := GenS30{F0: genFn0(), F1: v18}
                    _ = v215
                }
            default:
                if true {
                    v216 := GenS10{F0: ""}
                    _ = v216
                } else {
                    for fi217 := 0; fi217 < 18; fi217++ {
                        if ("" > ("" + "hello")) {
                            fmt.Println(v22, v2)
                        }
                        v18 = v18
                        v218 := []int{(fi217 - 32768), v91, 720}
                        _ = fi217
                        _ = v218
                    }
                    fmt.Println(v18)
                    if true {
                        if false {
                            var v219 int = 32767
                            _ = v219
                        } else {
                            v220 := GenS16{F0: -718}
                            _ = v220
                        }
                        if false {
                            v221 := []int{v91}
                            _ = v221
                        }
                        switch genFn0() {
                        case 32767:
                            v222 := GenS7{F0: 42, F1: 2147483647}
                            _ = v222
                        case 4294967295:
                            v223 := []int{v91, v91, v91, v91}
                            _ = v223
                        case 42:
                            switch v91 {
                            case -1:
                                v224 := []int{}
                                _ = v224
                            case 2147483647:
                                var v225 uint64 = 42
                                _ = v225
                            }
                        default:
                            var v226 int16 = -32768
                            _ = v226
                        }
                    }
                }
            }
        }
        _ = v91
    }
    for fi227 := 0; fi227 < 3; fi227++ {
        var v228 bool = false
        if true {
            v1 = 738
        }
        _ = fi227
        _ = v228
    }
    fmt.Println(v1)
    v22 += 0
    fmt.Println(v22, v18)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
}
