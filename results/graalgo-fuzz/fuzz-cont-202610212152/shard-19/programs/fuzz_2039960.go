package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 bool }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 int64 }
type GenS8 struct { F0 int64; F1 uint16; F2 int8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int16; F1 int16 }
type GenS11 struct { F0 bool; F1 uint16; F2 int8 }
type GenS12 struct { F0 int64; F1 string; F2 string }
type GenS13 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS14 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 bool }
func genFn0(p0 uint32) int {
    var v6 uint8 = 128
    if (!((false && false) && true)) {
        var v7 int16 = 42
        _ = v7
    } else {
        p0 = p0
        v6 += (^(v6))
        v6 = v6
    }
    var v8 float64 = 0.0
    var v9 uint8 = 42
    _ = p0
    _ = v6
    _ = v8
    _ = v9
    return -872
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint8 = 127
    v3 := []uint32{4294967295, 4294967295, 805, 0}
    var v4 uint8 = 0
    v5 := []string{}
    if (false || ("hello" == "hello")) {
        _ = len(v3)
    }
    if true {
        for fi10 := 0; fi10 < 12; fi10++ {
            switch "foo bar" {
            case "foo bar":
                switch "" {
                case "foo bar":
                    var v11 string = "foo bar"
                    _ = v11
                default:
                    var v12 string = "foo bar"
                    _ = v12
                }
            case "\n\t":
                if len(v3) > 0 {
                    v3[0] = 2
                }
            case "hello":
                for fi13 := 0; fi13 < 5; fi13++ {
                    v1 = v1
                    v14 := GenS1{F0: 562, F1: 18446744073709551615, F2: false}
                    var v15 float64 = 1.5
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            default:
                for fi16 := 0; fi16 < 7; fi16++ {
                    _ = len(v3)
                    _ = fi16
                }
            }
            _ = fi10
        }
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 uint64 = 0
            _ = fi17
            _ = v18
        }
        v1 += (v1 ^ v1)
    } else {
        var v19 float64 = 1e10
        switch 42 {
        case 48:
            fmt.Println(v2)
        case 32767:
            if (255 >= 89) {
                for fi20 := 0; fi20 < 12; fi20++ {
                    v19 = ((v19 - v19) + 0.0)
                    switch "a" {
                    case "\n\t":
                        v21 := GenS2{F0: -1, F1: true}
                        _ = v21
                    case "foo bar":
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(v19, v4)
                        case "\n\t":
                            var v22 int8 = -128
                            _ = v22
                        case "":
                            var v23 float32 = (3.14159 + 0.0)
                            _ = v23
                        }
                    case "a":
                        if (1 >= 100) {
                            if true {
                                v1 = 2147483647
                            } else {
                                v24 := GenS3{F0: 0}
                                if len(v5) > 0 {
                                    v5[0] = "foo bar"
                                }
                                _ = v24
                            }
                            v25 := GenS4{F0: false, F1: -9223372036854775808}
                            _ = v25
                        }
                    default:
                        for fi26 := 0; fi26 < 11; fi26++ {
                            v19 += -2.25
                            if ((v1 >> ((65535) & 63)) > v1) {
                                _ = len(v3)
                            } else {
                                var v27 float64 = v19
                                switch "a" {
                                case "hello":
                                    var v28 int32 = v1
                                    _ = v28
                                case "":
                                    if (!("hello" > "")) {
                                        var v29 int = 128
                                        if ((-1 >= -1) && (712 == 1)) {
                                            if false {
                                                var v30 string = ("hello" + "\n\t")
                                                var v31 int8 = 1
                                                var v32 int8 = v31
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            }
                                            switch 939 {
                                            case -662:
                                                v29 = (fi26 % ((32767) | 1))
                                            case -862:
                                                if false {
                                                    v1 = (-(v1))
                                                    v27 += v19
                                                } else {
                                                    v27 = (float64((((0.0 + 1e10) / 1e10) + 1.5)) + (v19 + v27))
                                                    v33 := GenS2{F0: 42, F1: true}
                                                    _ = v33
                                                }
                                                if true {
                                                    _ = len(v5)
                                                    var v34 int64 = 9223372036854775807
                                                    _ = v34
                                                } else {
                                                    v35 := GenS5{F0: 2, F1: 2}
                                                    _ = v35
                                                }
                                                var v36 bool = (!true)
                                                v2 -= v4
                                                fmt.Println(v1)
                                                v1 = (^(v1))
                                                var v37 uint = 18446744073709551615
                                                _ = v36
                                                _ = v37
                                            case 32767:
                                                var v38 float32 = -2.25
                                                _ = v38
                                            default:
                                                var v39 int32 = (v1 >> ((1) & 63))
                                                _ = v39
                                            }
                                            if (1 > 65535) {
                                                v1 = (v1 << ((123) & 63))
                                                switch fi26 {
                                                case 2147483647:
                                                    switch (fi20 << ((0) & 63)) {
                                                    case -938:
                                                        v2 -= 128
                                                    default:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                }
                                                if true {
                                                    v40 := GenS1{F0: 42, F1: 916, F2: false}
                                                    fmt.Println(v2, v2, v4)
                                                    _ = v40
                                                }
                                                switch ("hello" + ("" + "hello")) {
                                                case "a":
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case "foo bar":
                                                    if (false || false) {
                                                        switch fi26 {
                                                        case 948:
                                                            if false {
                                                                v41 := GenS6{F0: 826}
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    v2 = v4
                                                                }
                                                                _ = v41
                                                            } else {
                                                                v2 -= 2
                                                                if (127 == v4) {
                                                                    v1 += (v1 - (v1 + (v1 >> ((88) & 63))))
                                                                } else {
                                                                    v4 = (v4 / (((v4 & (v2 >> ((924) & 63)))) | 1))
                                                                }
                                                                v27 = 1.5
                                                            }
                                                        case -499:
                                                            v42 := GenS7{F0: 4294967295, F1: 1}
                                                            _ = v42
                                                        }
                                                    }
                                                case "hello":
                                                    if (958 > 42) {
                                                        if ((!((-191 != v29) && true)) || (1 == 901)) {
                                                            var v43 bool = true
                                                            _ = v43
                                                        } else {
                                                            var v44 uint = 4294967295
                                                            if false {
                                                                var v45 uint = (v44 << (((v44 >> ((359) & 63))) & 63))
                                                                _ = v45
                                                            } else {
                                                                var v46 int16 = 1
                                                                var v47 uint32 = 394
                                                                v48 := GenS5{F0: 565, F1: uint64(v44)}
                                                                _ = v46
                                                                _ = v47
                                                                _ = v48
                                                            }
                                                            _ = v44
                                                        }
                                                        var v49 uint32 = 2147483647
                                                        v27 = 1.5
                                                        _ = v49
                                                    }
                                                }
                                                v29 -= fi26
                                            }
                                        } else {
                                            v5 = append(v5, "foo bar")
                                            v50 := GenS5{F0: 0, F1: 1}
                                            var v51 uint32 = 682
                                            if true {
                                                switch v29 {
                                                case 32767:
                                                    v52 := GenS8{F0: -1, F1: 1, F2: 0}
                                                    _ = v52
                                                case 145:
                                                    var v53 uint16 = 852
                                                    _ = v53
                                                }
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = v51
                                            }
                                            _ = v50
                                            _ = v51
                                        }
                                        if (1 <= 732) {
                                            switch "" {
                                            case "a":
                                                switch "" {
                                                case "":
                                                    v27 = v27
                                                case "foo bar":
                                                    var v54 uint32 = 65535
                                                    _ = v54
                                                case "a":
                                                    switch (v29 * 0) {
                                                    case -1:
                                                        v55 := GenS7{F0: 341, F1: 9223372036854775807}
                                                        _ = v55
                                                    case 128:
                                                        v1 = (v1 - int32(v29))
                                                    default:
                                                        v1 += (v1 / ((v1) | 1))
                                                    }
                                                }
                                            case "\n\t":
                                                if false {
                                                    var v56 float32 = 3.14159
                                                    var v57 uint16 = 32767
                                                    v57 += (v57 ^ 2)
                                                    _ = v56
                                                    _ = v57
                                                } else {
                                                    v5 = append(v5, (("a" + "foo bar") + ((("hello" + (("a" + "\n\t") + "")) + "hello") + "hello")))
                                                    var v58 int32 = v1
                                                    var v59 uint = 42
                                                    _ = v58
                                                    _ = v59
                                                }
                                            default:
                                                fmt.Println(fi20, v4, fi26)
                                            }
                                        }
                                        _ = v29
                                    }
                                case "\n\t":
                                    v60 := GenS8{F0: 988, F1: 557, F2: -17}
                                    var v61 uint8 = v2
                                    switch "" {
                                    case "a":
                                        if len(v5) > 0 {
                                            v5[0] = ("hello" + "\n\t")
                                        }
                                    case "":
                                        var v62 int16 = 1
                                        _ = v62
                                    case "\n\t":
                                        var v63 uint = 2
                                        _ = v63
                                    }
                                    _ = v60
                                    _ = v61
                                default:
                                    v19 += v19
                                }
                                switch fi20 {
                                case -586:
                                    v27 = v19
                                }
                                v27 = (3.14159 * v27)
                                _ = len(v3)
                                var v64 int = 32767
                                _ = v27
                                _ = v64
                            }
                            _ = fi26
                        }
                    }
                    for fi65 := 0; fi65 < 9; fi65++ {
                        v66 := GenS4{F0: true, F1: 9223372036854775807}
                        var v67 uint32 = 42
                        v1 = (v1 << ((0) & 63))
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                    _ = fi20
                }
                fmt.Println(v19)
                v68 := GenS2{F0: -9223372036854775808, F1: true}
                _ = v68
            }
        case 336:
            if len(v3) > 0 {
                v3[0] = 581
            }
        default:
            for fi69 := 0; fi69 < 17; fi69++ {
                for fi70 := 0; fi70 < 4; fi70++ {
                    var v71 bool = (42 != 9223372036854775807)
                    _ = fi70
                    _ = v71
                }
                var v72 int16 = -979
                v1 -= (v1 * (v1 + v1))
                _ = fi69
                _ = v72
            }
        }
        if (((true && true) || (!false)) && (1 != 127)) {
            for fi73 := 0; fi73 < 12; fi73++ {
                v19 += (v19 * v19)
                for fi74 := 0; fi74 < 11; fi74++ {
                    var v75 uint8 = (v2 | v4)
                    v75 = 162
                    v76 := GenS5{F0: 18446744073709551615, F1: 350}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi74
                    _ = v75
                    _ = v76
                }
                v2 -= v2
                _ = fi73
            }
            v19 += v19
            for fi77 := 0; fi77 < 6; fi77++ {
                v4 += (v4 << ((4294967295) & 63))
                var v78 string = ("foo bar" + (("" + ("a" + ("foo bar" + "hello"))) + "a"))
                if false {
                    var v79 uint64 = 2
                    v80 := GenS4{F0: (true && false), F1: 0}
                    for fi81 := 0; fi81 < 7; fi81++ {
                        if (-128 < 0) {
                            var v82 uint8 = (v2 + v2)
                            var v83 bool = (true && true)
                            _ = v82
                            _ = v83
                        }
                        var v84 int64 = -828
                        switch -824 {
                        case 128:
                            var v85 int = 202
                            switch "foo bar" {
                            case "\n\t":
                                var v86 uint32 = 42
                                _ = v86
                            case "":
                                fmt.Println(fi81)
                            case "a":
                                v79 = v79
                            default:
                                fmt.Println(v1, v4, v1)
                            }
                            _ = v85
                        default:
                            v84 -= (v84 & int64(v1))
                            v3 = append(v3, 262)
                        }
                        _ = fi81
                        _ = v84
                    }
                    _ = v79
                    _ = v80
                } else {
                    v87 := GenS5{F0: 4294967295, F1: 4294967295}
                    fmt.Println(v1)
                    for fi88 := 0; fi88 < 9; fi88++ {
                        var v89 int = (^(fi77))
                        v78 = "\n\t"
                        var v90 float64 = v19
                        v91 := GenS3{F0: v2}
                        v92 := GenS2{F0: -345, F1: false}
                        _ = fi88
                        _ = v89
                        _ = v90
                        _ = v91
                        _ = v92
                    }
                    _ = v87
                }
                _ = fi77
                _ = v78
            }
        }
        _ = v19
    }
    if (true && (1e10 != (float32(3.14159) - -2.25))) {
        for fi93 := 0; fi93 < 14; fi93++ {
            for fi94 := 0; fi94 < 8; fi94++ {
                if true {
                    v2 -= v4
                    var v95 int64 = 42
                    if true {
                        v1 -= 894
                        v4 += 127
                        v4 = 128
                    }
                    _ = v95
                } else {
                    v96 := GenS1{F0: 2, F1: 74, F2: false}
                    _ = v96
                }
                _ = fi94
            }
            var v97 uint = 157
            _ = fi93
            _ = v97
        }
        v98 := GenS9{F0: 0}
        _ = v98
    }
    switch -2147483648 {
    case 128:
        v2 = 255
    case -992:
        fmt.Println(v4, v4)
    case -162:
        switch "" {
        case "":
            if true {
                v2 = 255
            } else {
                v4 += 123
                var v99 int8 = -128
                _ = v99
            }
        case "a":
            var v100 int16 = 1
            _ = v100
        case "hello":
            fmt.Println(v1, v2, v4)
        default:
            v101 := GenS6{F0: 118}
            _ = v101
        }
    }
    var v102 int8 = 100
    if true {
        switch 1 {
        case -9223372036854775808:
            v2 += v4
        case 42:
            v103 := GenS10{F0: -655, F1: 0}
            _ = v103
        }
        if true {
            v4 += (^(v4))
            if (("a" + (("foo bar" + "hello") + "hello")) >= "foo bar") {
                fmt.Println(v102, v2, v4)
                if (1 == (-(v102))) {
                    switch 290 {
                    case -9223372036854775808:
                        if true {
                            fmt.Println(v102, v102, v2)
                            fmt.Println(v1, v4, v4)
                            v4 = (v4 / ((68) | 1))
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                }
                v1 = v1
            } else {
                v104 := GenS8{F0: 9223372036854775807, F1: 255, F2: 127}
                v105 := GenS1{F0: 1, F1: 141, F2: false}
                _ = v104
                _ = v105
            }
        }
    } else {
        v1 = 42
        fmt.Println(v1, v4, v102)
    }
    if len(v5) > 0 {
        v5[0] = ""
    }
    if (("\n\t" + "") <= "a") {
        v2 += (^(v4))
    }
    var v106 int64 = -666
    var v107 int16 = 32767
    v107 -= 1000
    if (4294967295 >= 65535) {
        if true {
            v2 = (v4 * (v4 * (^(v2))))
            for fi108 := 0; fi108 < 8; fi108++ {
                for fi109 := 0; fi109 < 7; fi109++ {
                    if true {
                        _ = len(v5)
                    } else {
                        if (491 <= 32767) {
                            if false {
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = 65535
                                    }
                                    var v110 int = int(v102)
                                    v1 += v1
                                    _ = v110
                                } else {
                                    v111 := GenS7{F0: 17, F1: 0}
                                    var v112 uint32 = 2
                                    _ = v111
                                    _ = v112
                                }
                            } else {
                                if ("a" >= "\n\t") {
                                    var v113 uint = 42
                                    v4 += 128
                                    _ = v113
                                } else {
                                    var v114 bool = ((901 < 2) && true)
                                    if v114 {
                                        var v115 int16 = (v107 * 32767)
                                        fmt.Println(v102, fi108)
                                        v116 := GenS11{F0: true, F1: 0, F2: (v102 & 127)}
                                        _ = len(v3)
                                        var v117 uint64 = 42
                                        var v118 int16 = v115
                                        v114 = ((v102 ^ (v102 << ((0) & 63))) >= v102)
                                        _ = v115
                                        _ = v116
                                        _ = v117
                                        _ = v118
                                    } else {
                                        if false {
                                            v119 := GenS3{F0: 187}
                                            var v120 uint = 20
                                            _ = v119
                                            _ = v120
                                        }
                                        fmt.Println(fi109, v2, fi109)
                                        v4 += 63
                                    }
                                    switch "" {
                                    case "a":
                                        if ((!true) && (!("hello" > "\n\t"))) {
                                            if (0 > 255) {
                                                fmt.Println(v1, v106, v114)
                                            }
                                            if len(v5) > 0 {
                                                v5[0] = "foo bar"
                                            }
                                        }
                                    case "":
                                        v102 -= (v102 | 42)
                                    case "foo bar":
                                        fmt.Println(v107)
                                    default:
                                        fmt.Println(v102)
                                    }
                                    _ = v114
                                }
                                fmt.Println(v2, fi109)
                                fmt.Println(v2)
                            }
                            switch fi108 {
                            case -96:
                                fmt.Println(v102, v4, v2)
                            case 4294967295:
                                var v121 uint16 = 65535
                                _ = v121
                            }
                            var v122 uint64 = 42
                            _ = v122
                        } else {
                            if true {
                                if false {
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch "foo bar" {
                                        case "hello":
                                            fmt.Println(v102)
                                            switch (-(fi109)) {
                                            case 128:
                                                if false {
                                                    var v123 int32 = -18
                                                    v1 += (v1 % ((v123) | 1))
                                                    _ = v123
                                                } else {
                                                    v124 := GenS12{F0: (v106 % ((v106) | 1)), F1: "", F2: ""}
                                                    v125 := GenS11{F0: (v2 == (^(v2))), F1: 321, F2: (v102 + v102)}
                                                    _ = v124
                                                    _ = v125
                                                }
                                            default:
                                                v106 -= v106
                                                v102 = (v102 << ((42) & 63))
                                            }
                                        case "foo bar":
                                            _ = len(v5)
                                        default:
                                            fmt.Println(v107, fi108, v102)
                                        }
                                        if ((v1 % ((v1) | 1)) < (v1 * v1)) {
                                            var v126 uint32 = 512
                                            _ = v126
                                        } else {
                                            switch (("hello" + "hello") + (("" + ("" + "")) + "")) {
                                            case "":
                                                var v127 uint8 = (v4 % ((v2) | 1))
                                                _ = v127
                                            case "\n\t":
                                                if len(v5) > 0 {
                                                    v5[0] = ""
                                                }
                                            }
                                            v106 += (v106 - (v106 << ((4294967295) & 63)))
                                        }
                                    case "a":
                                        v128 := GenS13{F0: (v1 % (((v1 & 0)) | 1)), F1: 0, F2: 42}
                                        switch ("" + "a") {
                                        case "foo bar":
                                            v3 = append(v3, 0)
                                        default:
                                            v129 := GenS5{F0: 0, F1: 208}
                                            _ = v129
                                        }
                                        _ = v128
                                    case "foo bar":
                                        var v130 int32 = (-(v1))
                                        _ = v130
                                    default:
                                        v131 := GenS7{F0: 98, F1: 9223372036854775807}
                                        _ = v131
                                    }
                                    v132 := GenS3{F0: 128}
                                    _ = v132
                                } else {
                                    if false {
                                        v1 = (v1 * v1)
                                        _ = len(v3)
                                    } else {
                                        var v133 uint8 = 42
                                        var v134 float64 = 3.14159
                                        _ = v133
                                        _ = v134
                                    }
                                    var v135 float64 = ((1e10 - (1e10 - (float64(0.1) / 1.5))) - 0.0)
                                    _ = len(v5)
                                    _ = v135
                                }
                            } else {
                                switch (^(fi108)) {
                                case 4294967295:
                                    v102 = v102
                                case 32768:
                                    if false {
                                        switch genFn0(2) {
                                        case -254:
                                            if (((-(v102)) >= 100) && (true && true)) {
                                                var v136 float32 = 0.1
                                                if true {
                                                    v1 = v1
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = ""
                                                }
                                                var v137 uint16 = 0
                                                _ = v136
                                                _ = v137
                                            } else {
                                                fmt.Println(v107, v106, v4)
                                                var v138 uint64 = 2
                                                v107 -= (v107 | (v107 ^ (v107 % ((v107) | 1))))
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                if (2 >= 42) {
                                                    fmt.Println(v106, v2, v4)
                                                    v102 = (v102 ^ -128)
                                                }
                                                _ = v138
                                            }
                                        default:
                                            if ((!true) && (18446744073709551615 == 4294967295)) {
                                                fmt.Println(v102, v1, v106)
                                            }
                                            v102 += 82
                                            v139 := GenS14{F0: 32767, F1: 0, F2: 2}
                                            _ = v139
                                        }
                                        v2 += (v4 * v4)
                                        v140 := GenS15{F0: 1000}
                                        _ = v140
                                    }
                                }
                                v102 = (-(v102))
                                var v141 int8 = 1
                                var v142 int32 = (v1 << ((2) & 63))
                                _ = v141
                                _ = v142
                            }
                            v143 := GenS1{F0: 749, F1: 136, F2: (!((v4 < 0) || false))}
                            v144 := GenS15{F0: v107}
                            _ = v143
                            _ = v144
                        }
                        v4 = v2
                    }
                    _ = fi109
                }
                _ = fi108
            }
            for fi145 := 0; fi145 < 9; fi145++ {
                if (-32768 > v107) {
                    var v146 float32 = (0.0 / 0.1)
                    switch 150 {
                    case 1:
                        var v147 int32 = (^(v1))
                        _ = v147
                    case 42:
                        v2 += v4
                    default:
                        fmt.Println(fi145, v102)
                    }
                    _ = v146
                } else {
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                }
                _ = fi145
            }
        } else {
            v148 := GenS3{F0: v2}
            _ = v148
        }
        var v149 int64 = 9223372036854775807
        v4 = (^(v2))
        _ = v149
    } else {
        v150 := GenS7{F0: 42, F1: 42}
        _ = v150
    }
    v151 := GenS5{F0: 18446744073709551615, F1: 0}
    for fi152 := 0; fi152 < 7; fi152++ {
        v107 = (v107 ^ (v107 ^ v107))
        _ = fi152
    }
    var v153 float64 = (3.14159 / (float64((3.14159 - ((-2.25 + 1.5) / 1.0))) / 0.1))
    v154 := GenS5{F0: 966, F1: 401}
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    v155 := GenS16{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v106:", v106)
    fmt.Println("v107:", v107)
    fmt.Println("v153:", v153)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v102
    _ = v106
    _ = v107
    _ = v151
    _ = v153
    _ = v154
    _ = v155
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
