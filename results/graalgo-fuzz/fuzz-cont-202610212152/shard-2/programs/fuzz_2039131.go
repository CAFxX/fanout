package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int; F1 int; F2 uint64 }
type GenS3 struct { F0 bool; F1 bool; F2 int }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64; F1 uint16; F2 int32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int32; F1 uint16; F2 int }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 bool; F1 int8 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 bool; F1 int32 }
type GenS18 struct { F0 int32; F1 uint8; F2 int8 }
type GenS19 struct { F0 uint8 }
func genFn0(p0 int64, p1 string) bool {
    fmt.Println(p1, p0)
    fmt.Println(p1, p1)
    v5 := []int{-2147483648, 9223372036854775807}
    if (p1 != ("hello" + "a")) {
        p1 = (("foo bar" + "a") + (p1 + p1))
        v6 := GenS1{F0: 32767, F1: p1}
        for fi7 := 0; fi7 < 9; fi7++ {
            p0 -= -1
            if (fi7 != int(p0)) {
                if (36 != 0) {
                    var v8 float64 = 3.14159
                    for fi9 := 0; fi9 < 2; fi9++ {
                        switch (p1 + "hello") {
                        case "\n\t":
                            var v10 int16 = 536
                            _ = v10
                        case "foo bar":
                            var v11 uint8 = 0
                            _ = v11
                        case "hello":
                            if (!(18446744073709551615 == 4294967295)) {
                                _ = len(v5)
                                if false {
                                    if (!true) {
                                        switch (p1 + p1) {
                                        case "a":
                                            if true {
                                                v12 := GenS2{F0: fi9, F1: 42, F2: 1}
                                                v13 := GenS2{F0: -9223372036854775808, F1: 4294967295, F2: 18446744073709551615}
                                                fmt.Println(p0)
                                                _ = v12
                                                _ = v13
                                            } else {
                                                if ("" == p1) {
                                                    if (false && false) {
                                                        var v14 uint = 0
                                                        p1 = p1
                                                        _ = v14
                                                    }
                                                    fmt.Println(fi7, p1)
                                                }
                                            }
                                        case "hello":
                                            var v15 float32 = 0.1
                                            _ = v15
                                        case "foo bar":
                                            switch p1 {
                                            case "foo bar":
                                                switch (fi7 << ((1) & 63)) {
                                                case 4294967295:
                                                    if len(v5) > 0 {
                                                        v5[0] = (fi7 + (fi9 * (fi7 << ((2) & 63))))
                                                    }
                                                case -1:
                                                    var v16 int64 = 0
                                                    _ = v16
                                                case 162:
                                                    v17 := GenS3{F0: (!true), F1: false, F2: 322}
                                                    _ = v17
                                                default:
                                                    var v18 bool = false
                                                    _ = v18
                                                }
                                                if (0 > 32767) {
                                                    switch fi7 {
                                                    case 42:
                                                        if false {
                                                            v5 = append(v5, (fi9 ^ (fi7 % ((fi9) | 1))))
                                                        } else {
                                                            if true {
                                                                var v19 string = ("hello" + (p1 + p1))
                                                                if false {
                                                                    switch (fi9 & fi7) {
                                                                    case -356:
                                                                        v20 := GenS4{F0: 255, F1: 0}
                                                                        _ = v20
                                                                    case -567:
                                                                        var v21 uint8 = 127
                                                                        _ = v21
                                                                    case 965:
                                                                        if (false && (1000 <= -32768)) {
                                                                            fmt.Println(p0)
                                                                            if false {
                                                                                p0 = 796
                                                                            } else {
                                                                                if true {
                                                                                    v22 := GenS4{F0: 0, F1: 18446744073709551615}
                                                                                    p1 = v19
                                                                                    v19 = "a"
                                                                                    if false {
                                                                                        p0 -= 812
                                                                                    } else {
                                                                                        var v23 int = -1
                                                                                        _ = v23
                                                                                    }
                                                                                    p1 = v19
                                                                                    _ = v22
                                                                                } else {
                                                                                    p0 = (p0 & 9223372036854775807)
                                                                                }
                                                                                v24 := GenS5{F0: 32767, F1: 857}
                                                                                _ = v24
                                                                            }
                                                                            fmt.Println(p1, v8)
                                                                        } else {
                                                                            if (true && true) {
                                                                                p1 = (p1 + v19)
                                                                            } else {
                                                                                var v25 int8 = -128
                                                                                if false {
                                                                                    v25 = (v25 | (-(v25)))
                                                                                } else {
                                                                                    var v26 int32 = -145
                                                                                    p1 = (((("\n\t" + p1) + "") + "hello") + (("foo bar" + (((("foo bar" + "") + "hello") + p1) + p1)) + v19))
                                                                                    _ = v26
                                                                                }
                                                                                var v27 int8 = int8(p0)
                                                                                _ = v25
                                                                                _ = v27
                                                                            }
                                                                            switch ((v19 + (((v19 + (v19 + "foo bar")) + p1) + v19)) + p1) {
                                                                            case "foo bar":
                                                                                var v28 uint8 = 127
                                                                                _ = v28
                                                                            }
                                                                        }
                                                                    default:
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    }
                                                                    p0 -= p0
                                                                }
                                                                _ = v19
                                                            }
                                                            if len(v5) > 0 {
                                                                v5[0] = fi7
                                                            }
                                                        }
                                                    case 0:
                                                        var v29 uint16 = 32767
                                                        _ = v29
                                                    }
                                                } else {
                                                    var v30 int = -608
                                                    p0 += (p0 % ((p0) | 1))
                                                    var v31 uint8 = 2
                                                    v32 := GenS5{F0: 36, F1: 65535}
                                                    v33 := GenS6{F0: 2147483647, F1: (!true)}
                                                    var v34 int64 = p0
                                                    _ = v30
                                                    _ = v31
                                                    _ = v32
                                                    _ = v33
                                                    _ = v34
                                                }
                                            }
                                        default:
                                            if true {
                                                if false {
                                                    p0 = 695
                                                }
                                            } else {
                                                var v35 float64 = -2.25
                                                var v36 bool = (-2.25 == 0.0)
                                                _ = v35
                                                _ = v36
                                            }
                                        }
                                    } else {
                                        v37 := GenS7{F0: 461}
                                        fmt.Println(p0, fi9, p1)
                                        fmt.Println(p0, fi9)
                                        _ = v37
                                    }
                                }
                                v38 := GenS8{F0: (p0 & (p0 << ((1) & 63))), F1: 42, F2: 42}
                                _ = v38
                            } else {
                                p0 -= p0
                            }
                            var v39 int32 = 595
                            _ = v39
                        }
                        switch fi7 {
                        case 32767:
                            if (false || (true || ((2 >= 0) && false))) {
                                p0 -= p0
                                switch (("a" + p1) + "\n\t") {
                                case "hello":
                                    if false {
                                        v40 := GenS9{F0: 9223372036854775807}
                                        switch p1 {
                                        case "\n\t":
                                            var v41 bool = true
                                            _ = v41
                                        }
                                        v8 -= v8
                                        v42 := GenS2{F0: fi7, F1: fi7, F2: 42}
                                        fmt.Println(fi7, fi9)
                                        _ = v40
                                        _ = v42
                                    }
                                case "foo bar":
                                    switch -312 {
                                    case 128:
                                        var v43 string = p1
                                        _ = v43
                                    default:
                                        if len(v5) > 0 {
                                            v5[0] = (fi9 << ((384) & 63))
                                        }
                                    }
                                default:
                                    p0 += (p0 / ((-9223372036854775808) | 1))
                                }
                                switch -164 {
                                case 839:
                                    p1 = "foo bar"
                                case 1:
                                    v8 = (float64(float32(v8)) + v8)
                                    switch (p1 + "foo bar") {
                                    case "\n\t":
                                        var v44 uint32 = 0
                                        _ = v44
                                    case "foo bar":
                                        p0 += p0
                                    default:
                                        if true {
                                            switch "" {
                                            case "":
                                                switch p1 {
                                                case "foo bar":
                                                    v8 = v8
                                                case "\n\t":
                                                    p0 += p0
                                                case "a":
                                                    fmt.Println(p0, v8)
                                                default:
                                                    if len(v5) > 0 {
                                                        v5[0] = (^(fi9))
                                                    }
                                                    var v45 uint8 = 127
                                                    _ = v45
                                                }
                                            case "\n\t":
                                                v8 = 1.5
                                            default:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                            v46 := GenS10{F0: false}
                                            v47 := GenS10{F0: true}
                                            _ = v46
                                            _ = v47
                                        }
                                    }
                                case -233:
                                    fmt.Println(p1, p1, v8)
                                    p0 += (p0 + 131)
                                }
                            }
                        }
                        v5 = append(v5, fi7)
                        if len(v5) > 0 {
                            v5[0] = fi7
                        }
                        _ = fi9
                    }
                    _ = v8
                }
                if ((p0 ^ p0) > -9223372036854775808) {
                    switch ("hello" + "\n\t") {
                    case "":
                        v48 := GenS8{F0: p0, F1: 0, F2: 1}
                        _ = v48
                    case "hello":
                        if true {
                            if false {
                                var v49 uint64 = 977
                                var v50 int8 = -128
                                _ = v49
                                _ = v50
                            } else {
                                if (!(false && false)) {
                                    p0 = -9223372036854775808
                                    p0 = 1
                                } else {
                                    v51 := GenS8{F0: (p0 << ((0) & 63)), F1: 42, F2: 1}
                                    _ = v51
                                }
                                for fi52 := 0; fi52 < 5; fi52++ {
                                    var v53 uint8 = 70
                                    v54 := GenS11{F0: ((p1 + (p1 + p1)) + p1)}
                                    _ = fi52
                                    _ = v53
                                    _ = v54
                                }
                            }
                            switch (fi7 >> ((0) & 63)) {
                            case -2147483648:
                                for fi55 := 0; fi55 < 18; fi55++ {
                                    if true {
                                        var v56 bool = (fi55 != (fi55 + 127))
                                        p0 = (p0 / (((p0 / ((p0) | 1))) | 1))
                                        v57 := GenS4{F0: 32767, F1: 2}
                                        _ = v56
                                        _ = v57
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        var v58 int = (fi55 & (fi7 | (-(fi55))))
                                        v59 := GenS8{F0: (p0 ^ p0), F1: 42, F2: 1}
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = fi55
                                }
                            case -9223372036854775808:
                                fmt.Println(p1, p1, p1)
                            default:
                                for fi60 := 0; fi60 < 3; fi60++ {
                                    var v61 uint16 = 255
                                    _ = len(v5)
                                    if len(v5) > 0 {
                                        v5[0] = fi60
                                    }
                                    p1 = "a"
                                    if len(v5) > 0 {
                                        v5[0] = (fi60 ^ (fi7 ^ 0))
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = -1
                                    }
                                    _ = fi60
                                    _ = v61
                                }
                            }
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v62 string = "foo bar"
                            var v63 int32 = -1
                            _ = v62
                            _ = v63
                        }
                    case "\n\t":
                        for fi64 := 0; fi64 < 1; fi64++ {
                            p1 = ("" + ("" + ""))
                            switch p1 {
                            case "foo bar":
                                var v65 uint8 = 193
                                _ = v65
                            case "\n\t":
                                if ((2 != 18446744073709551615) || (!(2147483647 == 0))) {
                                    _ = len(v5)
                                    p1 = p1
                                } else {
                                    p1 = p1
                                    switch p1 {
                                    case "a":
                                        fmt.Println(fi7, fi7, p1)
                                    default:
                                        v66 := GenS9{F0: (fi7 - (^(fi7)))}
                                        fmt.Println(p1)
                                        var v67 uint16 = 255
                                        _ = v66
                                        _ = v67
                                    }
                                }
                            }
                            _ = fi64
                        }
                    default:
                        for fi68 := 0; fi68 < 16; fi68++ {
                            var v69 uint16 = 520
                            _ = fi68
                            _ = v69
                        }
                    }
                    switch "\n\t" {
                    case "foo bar":
                        for fi70 := 0; fi70 < 1; fi70++ {
                            var v71 string = "hello"
                            v71 = ((p1 + v71) + ((p1 + "foo bar") + "a"))
                            _ = fi70
                            _ = v71
                        }
                    case "\n\t":
                        var v72 string = "hello"
                        _ = v72
                    case "a":
                        if (495 > 2147483647) {
                            var v73 uint32 = 820
                            _ = v73
                        }
                    }
                    for fi74 := 0; fi74 < 19; fi74++ {
                        p0 -= (p0 / ((42) | 1))
                        _ = fi74
                    }
                } else {
                    if false {
                        for fi75 := 0; fi75 < 11; fi75++ {
                            if true {
                                var v76 int32 = -286
                                _ = v76
                            }
                            var v77 bool = false
                            p1 = (p1 + (p1 + ((p1 + p1) + p1)))
                            _ = fi75
                            _ = v77
                        }
                        switch 127 {
                        case -2147483648:
                            v5 = append(v5, (fi7 + -9223372036854775808))
                        case -308:
                            v78 := GenS12{F0: 0}
                            _ = v78
                        }
                        var v79 int = (fi7 | fi7)
                        _ = v79
                    } else {
                        switch -2147483648 {
                        case 4294967295:
                            if false {
                                var v80 uint32 = 1
                                var v81 uint8 = 42
                                v80 = v80
                                _ = v80
                                _ = v81
                            }
                        case -2147483648:
                            for fi82 := 0; fi82 < 11; fi82++ {
                                var v83 string = ("" + "hello")
                                _ = fi82
                                _ = v83
                            }
                        }
                        var v84 int32 = 989
                        var v85 float64 = -2.25
                        _ = v84
                        _ = v85
                    }
                }
            }
            var v86 uint = 65535
            _ = fi7
            _ = v86
        }
        _ = v6
    } else {
        for fi87 := 0; fi87 < 1; fi87++ {
            fmt.Println(p0)
            _ = fi87
        }
    }
    p0 -= (-(p0))
    _ = p0
    _ = p1
    _ = v5
    return false
}
func genFn1(p0 int8, p1 string, p2 string) uint32 {
    v88 := GenS6{F0: 2147483647, F1: (!false)}
    p0 = -1
    v89 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v88
    _ = v89
    return 42
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int = 42
    var v3 uint8 = 4
    var v4 float64 = 3.14159
    v1 += (v1 * 780)
    switch "\n\t" {
    case "\n\t":
        v1 = v1
    }
    for fi90 := 0; fi90 < 9; fi90++ {
        v3 -= (v3 | (v3 - 0))
        v91 := GenS5{F0: 255, F1: 2}
        var v92 uint16 = 841
        _ = fi90
        _ = v91
        _ = v92
    }
    switch (v2 & 4294967295) {
    case -659:
        v2 = (^(v2))
    case 32768:
        v93 := []int{}
        _ = v93
    case 70:
        if true {
            v1 += v1
            for fi94 := 0; fi94 < 8; fi94++ {
                for fi95 := 0; fi95 < 15; fi95++ {
                    v96 := GenS10{F0: false}
                    v1 -= v1
                    v97 := []int{}
                    switch ("" + "hello") {
                    case "\n\t":
                        if len(v97) >= 2 {
                            v97 = v97[1:len(v97)]
                        }
                    case "hello":
                        v1 = (v1 << ((2) & 63))
                    case "":
                        var v98 uint32 = 261
                        _ = v98
                    default:
                        switch fi95 {
                        case -2147483648:
                            if len(v97) >= 2 {
                                v97 = v97[1:len(v97)]
                            }
                            v1 += v1
                            v99 := GenS7{F0: 1}
                            v2 -= 32768
                            _ = v99
                        }
                    }
                    fmt.Println(v3, fi95, fi94)
                    v100 := GenS5{F0: 97, F1: 179}
                    if len(v97) >= 2 {
                        v97 = v97[1:len(v97)]
                    }
                    if false {
                        v97 = append(v97, (fi95 % (((^(fi94))) | 1)))
                        if false {
                            v97 = append(v97, (fi94 + 4294967295))
                            v101 := GenS13{F0: 0, F1: 1, F2: (-(v2))}
                            if (v4 < v4) {
                                switch "foo bar" {
                                case "hello":
                                    var v102 uint = 65535
                                    _ = v102
                                case "foo bar":
                                    v103 := GenS14{F0: int64(fi95)}
                                    _ = v103
                                case "a":
                                    if true {
                                        v4 = 1.5
                                        fmt.Println(v2, fi95)
                                        var v104 float32 = 0.1
                                        var v105 int = 1
                                        if true {
                                            v106 := GenS3{F0: true, F1: ((96 > 566) || (1e10 >= float32(1e10))), F2: -941}
                                            switch (fi95 * (v105 >> ((513) & 63))) {
                                            case 0:
                                                switch (v2 / ((-910) | 1)) {
                                                case 0:
                                                    var v107 int32 = 725
                                                    _ = v107
                                                case 127:
                                                    switch ((("hello" + ("" + ("a" + "\n\t"))) + (("" + ("hello" + "\n\t")) + "")) + "") {
                                                    case "foo bar":
                                                        var v108 uint64 = (v1 >> ((65535) & 63))
                                                        var v109 int = (^(v105))
                                                        var v110 uint = 0
                                                        v1 -= (v1 / ((v1) | 1))
                                                        v109 += (^(v2))
                                                        _ = v108
                                                        _ = v109
                                                        _ = v110
                                                    case "hello":
                                                        v97 = append(v97, -230)
                                                    }
                                                default:
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v1 = v1
                                                        default:
                                                            fmt.Println(v3)
                                                        }
                                                    case "hello":
                                                        if false {
                                                            var v111 uint = 0
                                                            if (true && true) {
                                                                var v112 uint8 = 42
                                                                _ = v112
                                                            }
                                                            if (-1 >= -1) {
                                                                v97 = append(v97, fi95)
                                                                if true {
                                                                    if true {
                                                                        v97 = append(v97, (fi95 - v2))
                                                                    } else {
                                                                        var v113 int16 = 874
                                                                        _ = v113
                                                                    }
                                                                    v114 := GenS15{F0: true, F1: 31}
                                                                    _ = v114
                                                                }
                                                            } else {
                                                                var v115 uint32 = 42
                                                                switch ("" + "foo bar") {
                                                                case "hello":
                                                                    var v116 int8 = 42
                                                                    _ = v116
                                                                case "a":
                                                                    if (false || true) {
                                                                        v111 -= 670
                                                                        fmt.Println(v4, v2, v4)
                                                                        v3 += 0
                                                                    }
                                                                }
                                                                _ = v115
                                                            }
                                                            _ = v111
                                                        }
                                                    }
                                                }
                                            case -9223372036854775808:
                                                var v117 uint8 = uint8(v1)
                                                _ = v117
                                            default:
                                                v118 := GenS16{F0: v1}
                                                _ = v118
                                            }
                                            _ = v106
                                        } else {
                                            v4 = v4
                                            var v119 int64 = -9223372036854775808
                                            _ = v119
                                        }
                                        v1 = (v1 % ((v1) | 1))
                                        _ = v104
                                        _ = v105
                                    }
                                default:
                                    v2 = v2
                                }
                                if false {
                                    v2 -= fi95
                                }
                                v120 := GenS17{F0: (!((^(fi94)) <= 127)), F1: -1}
                                v121 := GenS4{F0: 65535, F1: (v1 + (v1 - (v1 >> ((4294967295) & 63))))}
                                v2 += (-(v2))
                                _ = v120
                                _ = v121
                            } else {
                                if true {
                                    v122 := GenS4{F0: 65535, F1: (v1 & v1)}
                                    _ = v122
                                }
                            }
                            _ = v101
                        } else {
                            if len(v97) > 0 {
                                v97[0] = (v2 + 516)
                            }
                            var v123 uint8 = 127
                            var v124 uint8 = v123
                            v97 = append(v97, fi95)
                            var v125 uint = uint(v1)
                            var v126 float64 = (-2.25 / v4)
                            _ = v123
                            _ = v124
                            _ = v125
                            _ = v126
                        }
                        var v127 float32 = (3.14159 / ((0.1 * float32((v4 * v4))) * (0.0 / 3.14159)))
                        _ = v127
                    }
                    v3 = (^(v3))
                    _ = fi95
                    _ = v100
                    _ = v96
                    _ = v97
                }
                var v128 uint = 599
                for fi129 := 0; fi129 < 18; fi129++ {
                    v130 := []int{}
                    _ = fi129
                    _ = v130
                }
                _ = fi94
                _ = v128
            }
            v131 := GenS8{F0: 0, F1: 0, F2: -2147483648}
            _ = v131
        } else {
            v132 := []int{v2}
            fmt.Println(v4)
            _ = v132
        }
    }
    var v133 uint16 = 577
    v1 = 65
    if (-32768 == 42) {
        for fi134 := 0; fi134 < 17; fi134++ {
            v135 := GenS13{F0: -1, F1: 2, F2: (^(fi134))}
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v3, v4)
            case "hello":
                v136 := []int{(-(v2)), int(v1), (v2 % (((-(v2))) | 1))}
                _ = v136
            case "a":
                var v137 int = (fi134 | (v2 & (v2 / ((fi134) | 1))))
                _ = v137
            default:
                var v138 int64 = -9223372036854775808
                _ = v138
            }
            _ = fi134
            _ = v135
        }
    } else {
        if false {
            var v139 bool = true
            var v140 int64 = 0
            var v141 int16 = 0
            _ = v139
            _ = v140
            _ = v141
        } else {
            if false {
                v142 := GenS17{F0: false, F1: 499}
                v3 -= v3
                _ = v142
            }
            v2 = (v2 >> ((65535) & 63))
            v143 := GenS18{F0: 2147483647, F1: (v3 & 2), F2: -1}
            _ = v143
        }
        v133 = (v133 >> ((2) & 63))
    }
    for fi144 := 0; fi144 < 14; fi144++ {
        for fi145 := 0; fi145 < 17; fi145++ {
            fmt.Println(fi144, v4, v3)
            var v146 uint64 = v1
            _ = fi145
            _ = v146
        }
        v133 -= v133
        if true {
            switch "" {
            case "a":
                v147 := GenS7{F0: 0}
                _ = v147
            case "\n\t":
                switch v2 {
                case 0:
                    if (v133 == v133) {
                        v148 := GenS19{F0: (v3 & 255)}
                        _ = v148
                    } else {
                        if (515 < uint(v1)) {
                            for fi149 := 0; fi149 < 9; fi149++ {
                                switch (-(v2)) {
                                case -638:
                                    var v150 int32 = 0
                                    _ = v150
                                case 128:
                                    fmt.Println(v1)
                                default:
                                    if false {
                                        v151 := []int{(fi149 & 444)}
                                        var v152 uint8 = (v3 ^ (v3 ^ (^(v3))))
                                        _ = v151
                                        _ = v152
                                    }
                                }
                                switch 1 {
                                case 9223372036854775807:
                                    if ((true && (42 >= -611)) || true) {
                                        v153 := GenS6{F0: 42, F1: ((((110 != -128) || ((fi149 + (v2 | (fi149 ^ 9223372036854775807))) == (fi144 / ((fi149) | 1)))) && (84 != 224)) && (!(-53 > 0)))}
                                        _ = v153
                                    } else {
                                        var v154 uint8 = v3
                                        fmt.Println(v2, fi149)
                                        if true {
                                            if (v1 != v1) {
                                                var v155 int32 = 2147483647
                                                v1 = 1
                                                _ = v155
                                            }
                                            v156 := []int{(v2 >> ((255) & 63)), (fi144 * (fi149 >> ((0) & 63))), 4294967295}
                                            _ = v156
                                        } else {
                                            if (!(0 != 100)) {
                                                var v157 bool = (-48 > 21)
                                                _ = v157
                                            } else {
                                                v2 = fi144
                                                var v158 uint8 = (v154 * v154)
                                                _ = v158
                                            }
                                            fmt.Println(v154, v1)
                                        }
                                        _ = v154
                                    }
                                case 468:
                                    if (true || (293 == 2)) {
                                        v159 := []int{}
                                        switch (fi149 << ((255) & 63)) {
                                        case 2147483647:
                                            var v160 uint16 = uint16(v3)
                                            _ = v160
                                        }
                                        var v161 float64 = v4
                                        _ = v159
                                        _ = v161
                                    }
                                case -342:
                                    var v162 uint32 = 4294967295
                                    _ = v162
                                default:
                                    if false {
                                        v163 := []int{(v2 * (fi144 & fi149)), 689, (v2 - 154), (fi144 | v2)}
                                        v163 = append(v163, 770)
                                        _ = v163
                                    }
                                }
                                var v164 int64 = 635
                                var v165 float64 = v4
                                _ = fi149
                                _ = v164
                                _ = v165
                            }
                            v3 = 220
                        }
                    }
                case -159:
                    var v166 uint64 = uint64(v133)
                    _ = v166
                case 42:
                    v167 := GenS18{F0: 281, F1: 255, F2: 0}
                    _ = v167
                default:
                    var v168 uint8 = (v3 & (v3 * (v3 - v3)))
                    _ = v168
                }
            }
        } else {
            for fi169 := 0; fi169 < 16; fi169++ {
                v170 := []int{(-(v2)), fi169, v2, int(v3)}
                _ = fi169
                _ = v170
            }
        }
        _ = fi144
    }
    var v171 uint32 = genFn1(1, "a", "a")
    v172 := []int{32767, (^(v2)), (v2 - v2), (-(v2))}
    if ((v3 >> ((72) & 63)) > 0) {
        var v173 int16 = 349
        v173 += v173
        if true {
            var v174 uint8 = 207
            if (!((v3 + (v3 + (^(v3)))) >= uint8(v133))) {
                v174 -= 2
                var v175 int8 = 1
                _ = v175
            }
            _ = v174
        }
        _ = v173
    } else {
        v3 = v3
    }
    fmt.Println(v2, v4, v1)
    switch "" {
    case "":
        if len(v172) >= 2 {
            v172 = v172[1:len(v172)]
        }
    case "\n\t":
        switch 2147483647 {
        case -412:
            fmt.Println(v4)
        case 127:
            for fi176 := 0; fi176 < 6; fi176++ {
                var v177 float32 = 0.1
                v1 = v1
                var v178 int32 = 1
                _ = fi176
                _ = v177
                _ = v178
            }
        case 32768:
            if true {
                if (float64((0.1 / -2.25)) > (v4 + float64(1e10))) {
                    for fi179 := 0; fi179 < 18; fi179++ {
                        v1 += v1
                        _ = fi179
                    }
                    for fi180 := 0; fi180 < 4; fi180++ {
                        fmt.Println(v3, v171)
                        var v181 uint32 = (^(v171))
                        _ = fi180
                        _ = v181
                    }
                }
                if false {
                    v3 -= uint8(v133)
                    var v182 int32 = 2147483647
                    v1 -= v1
                    _ = v182
                }
            }
        }
    case "a":
        v172 = append(v172, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v133:", v133)
    fmt.Println("v171:", v171)
    fmt.Println("slicelen v172:", len(v172), v172)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v133
    _ = v171
    _ = v172
    _ = v2
    _ = v3
    _ = v4
}
