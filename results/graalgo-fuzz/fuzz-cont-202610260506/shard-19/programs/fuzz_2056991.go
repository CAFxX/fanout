package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 uint32; F1 uint64; F2 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int16; F1 uint8 }
type GenS10 struct { F0 uint32; F1 int32 }
type GenS11 struct { F0 int; F1 int }
type GenS12 struct { F0 int64; F1 int8; F2 uint64 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 bool; F1 string; F2 int64 }
type GenS15 struct { F0 uint64; F1 int8; F2 string }
type GenS16 struct { F0 uint; F1 bool; F2 bool }
type GenS17 struct { F0 uint8; F1 int8; F2 uint }
type GenS18 struct { F0 string; F1 uint32; F2 uint64 }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 uint32; F1 int64; F2 int16 }
type GenS21 struct { F0 uint }
func genFn0() uint16 {
    var v7 int16 = -32768
    var v8 int8 = 100
    if true {
        var v9 uint32 = 4294967295
        _ = v9
    }
    _ = v7
    _ = v8
    return 120
}
func genFn1(p0 int, p1 int64, p2 uint16) uint8 {
    p0 += (p0 / (((p0 << ((65535) & 63))) | 1))
    v10 := GenS1{F0: 0, F1: 18446744073709551615}
    if false {
        if (42 >= 94) {
            if (1.5 >= (((1.5 / (1.5 / 3.14159)) + (0.0 - 1e10)) / (float32((0.0 - -2.25)) - 3.14159))) {
                if ((("a" >= "\n\t") || ((!((717 <= 0) || (("foo bar" + "\n\t") <= "a"))) && (-1 == 0))) || false) {
                    switch "hello" {
                    case "":
                        p2 = genFn0()
                    case "foo bar":
                        v11 := GenS1{F0: p1, F1: 4294967295}
                        _ = v11
                    case "a":
                        for fi12 := 0; fi12 < 6; fi12++ {
                            var v13 uint32 = 2147483648
                            for fi14 := 0; fi14 < 1; fi14++ {
                                fmt.Println(fi12, p0)
                                p0 = (^(fi14))
                                _ = fi14
                            }
                            _ = fi12
                            _ = v13
                        }
                    }
                }
                for fi15 := 0; fi15 < 10; fi15++ {
                    v16 := []int{(p0 & (-(p0))), (fi15 + (fi15 >> ((91) & 63))), p0, -291}
                    p2 = (p2 * (p2 & p2))
                    _ = fi15
                    _ = v16
                }
                for fi17 := 0; fi17 < 15; fi17++ {
                    v18 := GenS1{F0: (p1 >> ((0) & 63)), F1: 184}
                    p2 -= (^(p2))
                    var v19 int16 = 120
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
        } else {
            v20 := []int{p0}
            _ = v20
        }
        var v21 bool = (256 < 2147483648)
        _ = v21
    }
    switch "a" {
    case "a":
        for fi22 := 0; fi22 < 5; fi22++ {
            if false {
                for fi23 := 0; fi23 < 7; fi23++ {
                    if false {
                        if true {
                            if (!true) {
                                if true {
                                    switch "foo bar" {
                                    case "hello":
                                        v24 := GenS1{F0: (p1 * 0), F1: 18446744073709551615}
                                        _ = v24
                                    case "\n\t":
                                        var v25 float64 = ((1e10 * 3.14159) + ((0.1 - 3.14159) / float64(float32(-2.25))))
                                        _ = v25
                                    default:
                                        if (!(true || (false && (false && true)))) {
                                            if true {
                                                if false {
                                                    p0 = (fi22 + (fi23 & fi23))
                                                } else {
                                                    v26 := []int{(p0 | (p0 & (fi22 ^ fi23))), 1}
                                                    p0 += (p0 + -1)
                                                    v27 := GenS2{F0: 42, F1: fi23}
                                                    _ = v26
                                                    _ = v27
                                                }
                                                var v28 string = "\n\t"
                                                if (42 == 65535) {
                                                    if false {
                                                        v29 := GenS3{F0: 18446744073709551615}
                                                        if (true && (!true)) {
                                                            if (294 <= 2) {
                                                                if false {
                                                                    var v30 uint16 = p2
                                                                    _ = v30
                                                                } else {
                                                                    switch (v28 + v28) {
                                                                    case "\n\t":
                                                                        v31 := []int{(fi22 ^ (fi22 - -1)), 127, (p0 - p0), (fi22 * (p0 - (fi22 ^ p0)))}
                                                                        _ = v31
                                                                    }
                                                                    v32 := GenS4{F0: 42, F1: 1}
                                                                    fmt.Println(p2)
                                                                    _ = v32
                                                                }
                                                            } else {
                                                                v33 := GenS5{F0: 0, F1: 439}
                                                                _ = v33
                                                            }
                                                        }
                                                        _ = v29
                                                    } else {
                                                        var v34 bool = (!false)
                                                        _ = v34
                                                    }
                                                    if (!false) {
                                                        v35 := GenS6{F0: 65535, F1: 719, F2: (false || (-928 <= (p0 | fi22)))}
                                                        _ = v35
                                                    }
                                                    v28 = v28
                                                }
                                                _ = v28
                                            } else {
                                                p0 += p0
                                            }
                                            if true {
                                                switch p0 {
                                                case 32768:
                                                    var v36 uint64 = 2
                                                    _ = v36
                                                case 127:
                                                    fmt.Println(p1)
                                                default:
                                                    if (127 <= 127) {
                                                        fmt.Println(p1, p1)
                                                        var v37 int = (p0 - fi23)
                                                        _ = v37
                                                    }
                                                }
                                            }
                                            p1 -= p1
                                            v38 := GenS1{F0: (-(p1)), F1: 298}
                                            _ = v38
                                        } else {
                                            var v39 uint = 1
                                            fmt.Println(fi22)
                                            _ = v39
                                        }
                                    }
                                    var v40 int = fi23
                                    _ = v40
                                }
                                var v41 int8 = -128
                                v42 := []int{}
                                switch "foo bar" {
                                case "":
                                    v43 := GenS6{F0: 508, F1: 18446744073709551615, F2: false}
                                    _ = v43
                                default:
                                    var v44 int8 = 79
                                    _ = v44
                                }
                                _ = v41
                                _ = v42
                            }
                            p1 += p1
                            p2 += 906
                        }
                        if (42 < 48) {
                            fmt.Println(fi22, fi23, fi23)
                            if (((127 <= 13) && (true || true)) || true) {
                                p1 = p1
                                var v45 float32 = 0.0
                                _ = v45
                            } else {
                                var v46 string = ("\n\t" + (("foo bar" + "hello") + "a"))
                                v47 := []int{128}
                                p0 = -2147483648
                                _ = v46
                                _ = v47
                            }
                            fmt.Println(fi23, p2, fi23)
                        }
                    }
                    var v48 int16 = 684
                    switch ("hello" + "hello") {
                    case "hello":
                        fmt.Println(fi22, p0)
                    case "foo bar":
                        v49 := GenS3{F0: 4294967295}
                        _ = v49
                    }
                    _ = fi23
                    _ = v48
                }
            } else {
                switch fi22 {
                case 569:
                    for fi50 := 0; fi50 < 8; fi50++ {
                        if true {
                            fmt.Println(p2)
                            var v51 uint16 = p2
                            p2 -= 0
                            _ = v51
                        }
                        _ = fi50
                    }
                default:
                    var v52 bool = false
                    _ = v52
                }
            }
            p2 -= genFn0()
            if (!false) {
                if true {
                    v53 := []int{1, (fi22 / ((-2147483648) | 1))}
                    p0 += p0
                    var v54 float32 = ((1.5 * float32((-2.25 * ((0.0 - 0.1) * -2.25)))) + 0.0)
                    _ = v53
                    _ = v54
                } else {
                    switch ("foo bar" + ("" + (("a" + ("\n\t" + "foo bar")) + "a"))) {
                    case "a":
                        var v55 int = 257
                        _ = v55
                    default:
                        if (("a" == ("" + "a")) && false) {
                            v56 := GenS5{F0: 42, F1: 42}
                            v57 := GenS4{F0: -2147483648, F1: 0}
                            fmt.Println(p2, p2)
                            _ = v56
                            _ = v57
                        } else {
                            var v58 float32 = -2.25
                            switch "a" {
                            case "a":
                                var v59 bool = (!((p1 & 1) != p1))
                                _ = v59
                            case "foo bar":
                                fmt.Println(p2, p0, v58)
                            default:
                                for fi60 := 0; fi60 < 14; fi60++ {
                                    switch (fi22 ^ fi22) {
                                    case -949:
                                        fmt.Println(p0)
                                    default:
                                        if ((("\n\t" + "") + "a") > "") {
                                            v61 := []int{}
                                            _ = v61
                                        } else {
                                            fmt.Println(fi22, fi60, p1)
                                        }
                                    }
                                    if (42 < -205) {
                                        v62 := GenS1{F0: (-(p1)), F1: 65535}
                                        _ = v62
                                    } else {
                                        if true {
                                            p2 += 555
                                        } else {
                                            fmt.Println(fi22, fi22, v58)
                                            if ((fi22 - fi60) > fi22) {
                                                switch fi22 {
                                                case -1:
                                                    v63 := []int{fi22}
                                                    _ = v63
                                                default:
                                                    v64 := []int{127}
                                                    _ = v64
                                                }
                                            } else {
                                                switch (p0 << ((42) & 63)) {
                                                case 128:
                                                    v65 := GenS7{F0: 2}
                                                    _ = v65
                                                case -2147483648:
                                                    v58 = (3.14159 + v58)
                                                }
                                                p0 = p0
                                            }
                                            v66 := []int{fi60, (^(fi60))}
                                            var v67 int16 = 1
                                            fmt.Println(p2)
                                            _ = v66
                                            _ = v67
                                        }
                                        if true {
                                            switch (fi22 + fi22) {
                                            case -2147483648:
                                                v68 := GenS2{F0: 373, F1: (^(fi60))}
                                                var v69 int = (-(fi60))
                                                var v70 int = v69
                                                _ = v68
                                                _ = v69
                                                _ = v70
                                            case 4294967295:
                                                p2 = p2
                                            case 175:
                                                if false {
                                                    v58 = 1.5
                                                    var v71 int8 = 100
                                                    _ = v71
                                                }
                                            default:
                                                v72 := GenS1{F0: (p1 >> ((4294967295) & 63)), F1: 262}
                                                _ = v72
                                            }
                                            if (true && (false && true)) {
                                                v58 = 0.0
                                                v73 := GenS6{F0: 349, F1: 42, F2: false}
                                                if true {
                                                    v74 := GenS8{F0: (p2 / ((p2) | 1))}
                                                    _ = v74
                                                }
                                                var v75 int = (fi22 + (-(fi22)))
                                                switch ("\n\t" + "hello") {
                                                case "hello":
                                                    if (true || (!(true || (("hello" != "") || (100 >= int8(fi22)))))) {
                                                        v76 := []int{(p0 << ((1) & 63)), fi60}
                                                        _ = v76
                                                    }
                                                default:
                                                    var v77 uint = 39
                                                    _ = v77
                                                }
                                                fmt.Println(fi22, p2, v58)
                                                var v78 int16 = -189
                                                _ = v73
                                                _ = v75
                                                _ = v78
                                            }
                                        } else {
                                            p1 -= 0
                                            var v79 int16 = 1000
                                            _ = v79
                                        }
                                        var v80 int32 = 42
                                        _ = v80
                                    }
                                    v81 := GenS9{F0: int16(p2), F1: 1}
                                    _ = fi60
                                    _ = v81
                                }
                            }
                            fmt.Println(p0, p2, p1)
                            _ = v58
                        }
                    }
                    switch "\n\t" {
                    case "\n\t":
                        p0 = p0
                    }
                }
                var v82 int64 = (p1 << ((481) & 63))
                v83 := GenS4{F0: -482, F1: 0}
                _ = v82
                _ = v83
            } else {
                var v84 int = p0
                _ = v84
            }
            _ = fi22
        }
    case "":
        var v85 int8 = 100
        _ = v85
    default:
        for fi86 := 0; fi86 < 9; fi86++ {
            p1 = (p1 - p1)
            _ = fi86
        }
    }
    var v87 uint = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v87
    return 128
}

func main() {
    var v1 uint = 145
    v2 := []string{("hello" + ("a" + (("a" + "\n\t") + "hello"))), "hello", ("foo bar" + "hello"), "foo bar"}
    v3 := []int{42, 128, -716}
    var v4 int32 = 42
    var v5 int16 = 1000
    var v6 uint32 = 88
    for fi88 := 0; fi88 < 5; fi88++ {
        switch "hello" {
        case "hello":
            var v89 uint64 = 18446744073709551615
            _ = v89
        case "\n\t":
            switch 9223372036854775807 {
            case 128:
                if false {
                    var v90 uint16 = 0
                    _ = v90
                } else {
                    if true {
                        for fi91 := 0; fi91 < 19; fi91++ {
                            v92 := GenS9{F0: v5, F1: 42}
                            _ = fi91
                            _ = v92
                        }
                        if true {
                            var v93 float32 = float32((0.0 - (((3.14159 * -2.25) - -2.25) - 0.1)))
                            _ = v93
                        } else {
                            v6 = (^(v6))
                        }
                        v94 := GenS8{F0: 32767}
                        _ = v94
                    } else {
                        fmt.Println(v5)
                        if false {
                            v4 = (v4 ^ (v4 >> ((v1) & 63)))
                            v1 = v1
                            switch (("" + "foo bar") + "") {
                            case "hello":
                                v3 = append(v3, fi88)
                            case "a":
                                for fi95 := 0; fi95 < 19; fi95++ {
                                    fmt.Println(v4, fi88)
                                    _ = fi95
                                }
                            }
                        }
                        switch (fi88 * 32767) {
                        case 2147483647:
                            v5 += (v5 & 1)
                        }
                    }
                    if true {
                        var v96 uint = (v1 / ((288) | 1))
                        var v97 float32 = 0.1
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v96
                        _ = v97
                    }
                }
            default:
                switch (fi88 ^ (fi88 | (-(fi88)))) {
                case -429:
                    if true {
                        if true {
                            if (-9223372036854775808 >= 9223372036854775807) {
                                if len(v3) > 0 {
                                    v3[0] = 42
                                }
                            }
                            v98 := GenS5{F0: 41, F1: -536}
                            switch "\n\t" {
                            case "a":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "\n\t":
                                if (v1 < 255) {
                                    switch ("\n\t" + "a") {
                                    case "foo bar":
                                        v99 := GenS4{F0: (v4 - v4), F1: (^(v6))}
                                        _ = v99
                                    case "\n\t":
                                        var v100 uint16 = 32767
                                        _ = v100
                                    default:
                                        if false {
                                            if false {
                                                if true {
                                                    if false {
                                                        if (2147483647 >= (v6 - (v6 % ((v6) | 1)))) {
                                                            if (v1 == v1) {
                                                                v6 += v6
                                                                for fi101 := 0; fi101 < 2; fi101++ {
                                                                    v4 += (v4 & (v4 | v4))
                                                                    fmt.Println(fi101, v5)
                                                                    switch ((((("hello" + (("foo bar" + "a") + "hello")) + ("" + "hello")) + "foo bar") + "a") + "\n\t") {
                                                                    case "hello":
                                                                        if false {
                                                                            v4 += (v4 ^ (^(v4)))
                                                                            if ((int8(fi88) >= 42) || false) {
                                                                                if ((v4 >> (((v1 ^ (^(v1)))) & 63)) == v4) {
                                                                                    switch fi101 {
                                                                                    case 32768:
                                                                                        switch fi88 {
                                                                                        case 1:
                                                                                            if true {
                                                                                                var v102 int = (fi101 / ((2147483647) | 1))
                                                                                                var v103 uint64 = 242
                                                                                                var v104 uint32 = 316
                                                                                                _ = v102
                                                                                                _ = v103
                                                                                                _ = v104
                                                                                            } else {
                                                                                                var v105 float32 = (1.5 / (-2.25 / 1.0))
                                                                                                _ = v105
                                                                                            }
                                                                                        case 32767:
                                                                                            var v106 int = fi101
                                                                                            v107 := GenS7{F0: 1}
                                                                                            v4 += (v4 / ((-1) | 1))
                                                                                            _ = v106
                                                                                            _ = v107
                                                                                        case -1:
                                                                                            switch 9223372036854775807 {
                                                                                            case 0:
                                                                                                fmt.Println(v6, v5, fi88)
                                                                                            case -123:
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = fi88
                                                                                                }
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = fi88
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            v108 := GenS10{F0: (^(v6)), F1: 42}
                                                                                            _ = v108
                                                                                        }
                                                                                    default:
                                                                                        var v109 int64 = -1
                                                                                        _ = v109
                                                                                    }
                                                                                    _ = len(v3)
                                                                                } else {
                                                                                    if (132 == v1) {
                                                                                        v5 = 32767
                                                                                        fmt.Println(v4, fi88, v6)
                                                                                    } else {
                                                                                        v6 = 1
                                                                                        v110 := GenS11{F0: (fi101 - 4294967295), F1: fi88}
                                                                                        _ = v110
                                                                                    }
                                                                                    if ((v5 | v5) == v5) {
                                                                                        switch fi88 {
                                                                                        case 1:
                                                                                            var v111 int8 = 42
                                                                                            _ = v111
                                                                                        case 32767:
                                                                                            if false {
                                                                                                v3 = append(v3, (fi88 - fi101))
                                                                                            }
                                                                                        case 32768:
                                                                                            v112 := GenS12{F0: 9223372036854775807, F1: -128, F2: 1}
                                                                                            _ = v112
                                                                                        default:
                                                                                            v6 -= v6
                                                                                        }
                                                                                        var v113 uint = v1
                                                                                        v114 := GenS10{F0: v6, F1: v4}
                                                                                        _ = v113
                                                                                        _ = v114
                                                                                    }
                                                                                    v6 = 324
                                                                                }
                                                                                var v115 uint64 = 2
                                                                                v6 = (v6 >> (((v1 >> (((^(v1))) & 63))) & 63))
                                                                                _ = v115
                                                                            }
                                                                        }
                                                                    case "foo bar":
                                                                        var v116 int16 = v5
                                                                        _ = v116
                                                                    case "\n\t":
                                                                        switch fi101 {
                                                                        case -2147483648:
                                                                            v4 = (v4 - (v4 >> ((v1) & 63)))
                                                                        default:
                                                                            v117 := GenS9{F0: v5, F1: genFn1(2147483647, -1, 32767)}
                                                                            _ = v117
                                                                        }
                                                                        if len(v3) > 0 {
                                                                            v3[0] = fi88
                                                                        }
                                                                        v4 = v4
                                                                    default:
                                                                        var v118 int8 = 0
                                                                        _ = v118
                                                                    }
                                                                    v119 := GenS9{F0: (v5 | 42), F1: 184}
                                                                    v6 += 65535
                                                                    _ = fi101
                                                                    _ = v119
                                                                }
                                                            }
                                                        }
                                                        for fi120 := 0; fi120 < 18; fi120++ {
                                                            v121 := GenS3{F0: 2}
                                                            if true {
                                                                var v122 bool = true
                                                                _ = v122
                                                            }
                                                            _ = fi120
                                                            _ = v121
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = ("\n\t" + "")
                                                        }
                                                    } else {
                                                        var v123 uint8 = 42
                                                        switch (-(fi88)) {
                                                        case 840:
                                                            var v124 int = (fi88 % ((32768) | 1))
                                                            _ = v124
                                                        case 0:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        case 1:
                                                            fmt.Println(v123, v5)
                                                        }
                                                        for fi125 := 0; fi125 < 4; fi125++ {
                                                            v4 = v4
                                                            if (84 != int16(v123)) {
                                                                var v126 float64 = -2.25
                                                                _ = v126
                                                            }
                                                            _ = fi125
                                                        }
                                                        _ = v123
                                                    }
                                                    _ = len(v3)
                                                    switch "a" {
                                                    case "\n\t":
                                                        for fi127 := 0; fi127 < 17; fi127++ {
                                                            var v128 uint = v1
                                                            switch "" {
                                                            case "":
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            case "foo bar":
                                                                switch (fi127 % ((fi88) | 1)) {
                                                                case -9223372036854775808:
                                                                    v129 := GenS13{F0: (v4 / (((v4 << (((v128 << ((v1) & 63))) & 63))) | 1))}
                                                                    _ = v129
                                                                case 128:
                                                                    var v130 uint8 = 0
                                                                    _ = v130
                                                                case -900:
                                                                    if false {
                                                                        v131 := GenS14{F0: (!(0.0 < 3.14159)), F1: "foo bar", F2: 0}
                                                                        _ = v131
                                                                    }
                                                                default:
                                                                    var v132 string = "\n\t"
                                                                    _ = v132
                                                                }
                                                            case "a":
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    v133 := GenS7{F0: 42}
                                                                    _ = v133
                                                                }
                                                            }
                                                            if false {
                                                                v128 = v1
                                                            }
                                                            v3 = append(v3, fi88)
                                                            if false {
                                                                if false {
                                                                    v5 += 445
                                                                    v5 = v5
                                                                    var v134 uint8 = 2
                                                                    _ = v134
                                                                }
                                                            }
                                                            _ = len(v2)
                                                            _ = fi127
                                                            _ = v128
                                                        }
                                                    case "foo bar":
                                                        var v135 int32 = (v4 << (((v1 + v1)) & 63))
                                                        _ = v135
                                                    default:
                                                        var v136 uint64 = uint64(v1)
                                                        _ = v136
                                                    }
                                                }
                                                for fi137 := 0; fi137 < 16; fi137++ {
                                                    if ("hello" >= "hello") {
                                                        v4 = (v4 - (v4 ^ v4))
                                                    }
                                                    var v138 int = fi88
                                                    switch ("" + "\n\t") {
                                                    case "a":
                                                        var v139 int = -1
                                                        _ = v139
                                                    case "\n\t":
                                                        fmt.Println(fi88, fi137, v4)
                                                    case "foo bar":
                                                        var v140 int8 = -128
                                                        _ = v140
                                                    default:
                                                        var v141 string = "\n\t"
                                                        _ = v141
                                                    }
                                                    var v142 int = (v138 << ((v1) & 63))
                                                    v1 -= v1
                                                    var v143 uint32 = (^(v6))
                                                    _ = fi137
                                                    _ = v138
                                                    _ = v142
                                                    _ = v143
                                                }
                                                var v144 uint16 = 42
                                                _ = v144
                                            } else {
                                                switch -641 {
                                                case 1:
                                                    for fi145 := 0; fi145 < 7; fi145++ {
                                                        fmt.Println(fi88, v4, fi88)
                                                        var v146 int32 = (v4 | v4)
                                                        v2 = append(v2, ("hello" + "\n\t"))
                                                        _ = fi145
                                                        _ = v146
                                                    }
                                                case 4294967295:
                                                    v5 = -149
                                                case -9223372036854775808:
                                                    var v147 float64 = 0.1
                                                    _ = v147
                                                }
                                                var v148 uint16 = 1
                                                v3 = append(v3, (fi88 << ((uint(v148)) & 63)))
                                                _ = v148
                                            }
                                        }
                                    }
                                }
                            }
                            _ = v98
                        }
                    }
                case 0:
                    switch 128 {
                    case -120:
                        v4 = v4
                    case 32767:
                        for fi149 := 0; fi149 < 17; fi149++ {
                            switch 32768 {
                            case 32768:
                                if (!false) {
                                    switch ((((("a" + (("a" + "foo bar") + "\n\t")) + (("\n\t" + ("a" + "")) + "")) + "\n\t") + (((("a" + "\n\t") + ("hello" + "")) + ("" + "a")) + ("" + ("foo bar" + "")))) + "\n\t") {
                                    case "foo bar":
                                        if false {
                                            v6 -= 627
                                            v6 = 2147483648
                                            var v150 int32 = v4
                                            _ = v150
                                        }
                                    case "hello":
                                        if (-1 >= (-(v4))) {
                                            if (false && true) {
                                                v151 := GenS15{F0: 4294967295, F1: 1, F2: "a"}
                                                var v152 uint = 2
                                                v5 += (^(v5))
                                                var v153 uint64 = 18446744073709551615
                                                _ = v151
                                                _ = v152
                                                _ = v153
                                            }
                                            switch fi88 {
                                            case 4294967295:
                                                var v154 string = "foo bar"
                                                _ = v154
                                            case 0:
                                                var v155 int64 = -1
                                                _ = v155
                                            case -2147483648:
                                                if (!true) {
                                                    v156 := GenS16{F0: (^(v1)), F1: true, F2: true}
                                                    switch (("" + "") + (("a" + "") + "foo bar")) {
                                                    case "a":
                                                        v3 = append(v3, fi88)
                                                    }
                                                    v3 = append(v3, (fi88 / ((-9223372036854775808) | 1)))
                                                    v1 += (v1 ^ v1)
                                                    if true {
                                                        var v157 float32 = ((((0.0 - (0.1 - 1.5)) * 1e10) - -2.25) - ((0.1 * 1e10) / 1.5))
                                                        v6 += (^(v6))
                                                        _ = v157
                                                    }
                                                    _ = v156
                                                } else {
                                                    switch 32768 {
                                                    case -581:
                                                        if (2 == 4294967295) {
                                                            fmt.Println(fi149)
                                                            var v158 int64 = -1
                                                            if (!(("foo bar" < ("hello" + "")) || (((v1 >> ((v1) & 63)) <= (v1 >> ((v1) & 63))) && (!false)))) {
                                                                var v159 bool = ((!(false || ((1 > 42) || true))) && false)
                                                                if (100 > 1) {
                                                                    var v160 string = "a"
                                                                    var v161 int8 = 42
                                                                    v6 = (v6 << ((4294967295) & 63))
                                                                    _ = v160
                                                                    _ = v161
                                                                } else {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    if ((v4 & (v4 ^ 768)) != v4) {
                                                                        var v162 int32 = (v4 ^ (v4 * -2147483648))
                                                                        _ = v162
                                                                    }
                                                                }
                                                                _ = v159
                                                            } else {
                                                                if true {
                                                                    if false {
                                                                        if false {
                                                                            if (-1 <= 23) {
                                                                                v163 := GenS15{F0: 2, F1: 103, F2: "\n\t"}
                                                                                fmt.Println(fi88)
                                                                                var v164 int64 = 1
                                                                                _ = v163
                                                                                _ = v164
                                                                            }
                                                                            var v165 int32 = (v4 * -1)
                                                                            fmt.Println(v1, v165)
                                                                            if (((0.1 + (3.14159 - float32(1.5))) - -2.25) > 3.14159) {
                                                                                var v166 int64 = int64(fi149)
                                                                                _ = v166
                                                                            }
                                                                            fmt.Println(v165, v165, v165)
                                                                            _ = v165
                                                                        } else {
                                                                            v167 := GenS8{F0: 2}
                                                                            v168 := GenS9{F0: (^(v5)), F1: uint8(v1)}
                                                                            fmt.Println(v5, v6)
                                                                            var v169 int32 = v4
                                                                            _ = v167
                                                                            _ = v168
                                                                            _ = v169
                                                                        }
                                                                    } else {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (fi149 & fi149)
                                                                        }
                                                                        v5 -= 42
                                                                        v158 = (v158 * v158)
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        if false {
                                                                            v2 = append(v2, "")
                                                                            var v170 int32 = v4
                                                                            var v171 float32 = 3.14159
                                                                            var v172 bool = (!true)
                                                                            var v173 bool = ((v6 - (v6 - v6)) >= v6)
                                                                            _ = v170
                                                                            _ = v171
                                                                            _ = v172
                                                                            _ = v173
                                                                        } else {
                                                                            var v174 uint = 18446744073709551615
                                                                            _ = v174
                                                                        }
                                                                        var v175 int64 = v158
                                                                        var v176 uint32 = v6
                                                                        _ = v175
                                                                        _ = v176
                                                                    } else {
                                                                        var v177 float64 = 1.5
                                                                        v5 = (v5 / ((1) | 1))
                                                                        var v178 int64 = (-(v158))
                                                                        _ = v177
                                                                        _ = v178
                                                                    }
                                                                    _ = len(v2)
                                                                }
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                fmt.Println(v5, v158, v5)
                                                            }
                                                            _ = v158
                                                        }
                                                        var v179 int32 = -883
                                                        v179 -= -826
                                                        v6 -= (v6 >> ((1) & 63))
                                                        v5 = (^(v5))
                                                        var v180 int16 = (v5 / ((42) | 1))
                                                        _ = v179
                                                        _ = v180
                                                    }
                                                }
                                            default:
                                                v181 := GenS14{F0: (v6 > (^(v6))), F1: "", F2: 0}
                                                var v182 uint64 = 4294967295
                                                _ = v181
                                                _ = v182
                                            }
                                            v5 = -881
                                            v183 := GenS17{F0: 166, F1: 127, F2: 4294967295}
                                            _ = v183
                                        } else {
                                            var v184 float64 = (1e10 - 1e10)
                                            var v185 uint8 = 1
                                            _ = v184
                                            _ = v185
                                        }
                                    case "a":
                                        var v186 int16 = (v5 >> ((255) & 63))
                                        _ = v186
                                    default:
                                        v1 -= 188
                                    }
                                    if true {
                                        v6 += v6
                                        v187 := GenS18{F0: ((("a" + "foo bar") + (("hello" + "foo bar") + ("\n\t" + "a"))) + "\n\t"), F1: v6, F2: 2}
                                        _ = v187
                                    }
                                    v5 -= (^(v5))
                                } else {
                                    if false {
                                        _ = len(v3)
                                        fmt.Println(fi88)
                                        var v188 int64 = 966
                                        _ = v188
                                    } else {
                                        var v189 float64 = -2.25
                                        v190 := GenS9{F0: 32767, F1: genFn1((fi149 / ((-1) | 1)), 42, genFn0())}
                                        _ = v189
                                        _ = v190
                                    }
                                    var v191 int64 = 911
                                    v191 = -1
                                    var v192 uint8 = 234
                                    _ = v191
                                    _ = v192
                                }
                            }
                            if false {
                                v193 := GenS19{F0: v6}
                                if len(v2) > 0 {
                                    v2[0] = ""
                                }
                                _ = v193
                            }
                            v1 -= (^(v1))
                            _ = fi149
                        }
                    }
                default:
                    fmt.Println(fi88, fi88)
                }
            }
        default:
            v2 = append(v2, "\n\t")
        }
        _ = fi88
    }
    v1 = (v1 / ((449) | 1))
    for fi194 := 0; fi194 < 20; fi194++ {
        switch fi194 {
        case 385:
            var v195 uint = v1
            _ = v195
        case 32768:
            v196 := GenS20{F0: (^(v6)), F1: 714, F2: v5}
            _ = v196
        case 0:
            v4 = 455
        }
        _ = fi194
    }
    fmt.Println(v4, v5, v4)
    v197 := GenS21{F0: v1}
    for fi198 := 0; fi198 < 19; fi198++ {
        v199 := GenS18{F0: "hello", F1: (^(v6)), F2: 18446744073709551615}
        _ = fi198
        _ = v199
    }
    var v200 uint64 = 732
    v5 -= -598
    if ((-926 >= 9223372036854775807) || (float32(3.14159) <= 0.0)) {
        var v201 int16 = int16(v4)
        v1 = 340
        v201 += -32768
        _ = v201
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v200:", v200)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v197
    _ = v2
    _ = v200
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
