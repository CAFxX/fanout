package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS7 struct { F0 uint64; F1 int64; F2 int8 }
type GenS8 struct { F0 uint64; F1 int; F2 uint8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int16; F1 int64 }
type GenS11 struct { F0 uint8; F1 uint8 }
type GenS12 struct { F0 uint8; F1 uint64 }
type GenS13 struct { F0 bool; F1 int32; F2 uint }
type GenS14 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS15 struct { F0 int; F1 uint32 }
type GenS16 struct { F0 string; F1 int32 }
type GenS17 struct { F0 uint32 }
func genFn0(p0 int, p1 int32) uint16 {
    var v4 int8 = 36
    v5 := []int{(p0 & 9223372036854775807), (p0 | (p0 ^ (p0 / (((^(p0))) | 1))))}
    var v6 int64 = 1
    for fi7 := 0; fi7 < 11; fi7++ {
        _ = len(v5)
        fmt.Println(p0, v4)
        v6 = (-(v6))
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 755
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 4
    var v3 float32 = (1.5 / (0.1 / 0.1))
    if (-1 <= 42) {
        var v8 int8 = 0
        fmt.Println(v2, v3)
        _ = v8
    }
    if (-1 <= 1) {
        if (0 < 1) {
            var v9 uint16 = 32767
            for fi10 := 0; fi10 < 13; fi10++ {
                for fi11 := 0; fi11 < 19; fi11++ {
                    v9 = (v9 * genFn0(fi11, 1))
                    _ = len(v1)
                    switch "foo bar" {
                    case "\n\t":
                        v3 = float32(0.0)
                    }
                    _ = fi11
                }
                _ = fi10
            }
            _ = v9
        }
        var v12 uint32 = 65535
        if true {
            for fi13 := 0; fi13 < 19; fi13++ {
                v14 := []int{}
                _ = fi13
                _ = v14
            }
            switch "hello" {
            case "a":
                if (2 == 42) {
                    var v15 uint64 = 2
                    _ = v15
                }
            case "hello":
                v16 := []int{127, 32768, -906}
                _ = v16
            case "":
                v2 = 0
            default:
                delete(v1, -474)
            }
            v3 = ((float32(3.14159) / 0.1) / 3.14159)
        } else {
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 uint64 = 622
                _ = fi17
                _ = v18
            }
            v19 := []int{}
            _ = v19
        }
        _ = v12
    } else {
        _ = len(v1)
    }
    v2 += 42
    var v20 int16 = 1000
    switch "hello" {
    case "hello":
        switch ("a" + "hello") {
        case "a":
            v21 := []int{-301}
            _ = v21
        case "\n\t":
            var v22 int32 = 2147483647
            _ = v22
        case "foo bar":
            fmt.Println(v2, v3)
        default:
            v23 := GenS1{F0: -950, F1: 28}
            _ = v23
        }
    case "a":
        var v24 uint = 2
        _ = v24
    default:
        var v25 float32 = (0.0 + v3)
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        if false {
            var v27 float64 = 1.5
            _ = v27
        }
        switch 32767 {
        case 1:
            v28 := GenS2{F0: -128, F1: genFn0(fi26, v2)}
            _ = v28
        case 4294967295:
            if true {
                fmt.Println(v2)
            } else {
                delete(v1, (fi26 & (-(fi26))))
            }
        case 9223372036854775807:
            v2 = (v2 / ((v2) | 1))
        }
        _ = fi26
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        v3 = v3
        switch "a" {
        case "a":
            for fi30 := 0; fi30 < 9; fi30++ {
                if false {
                    v31 := []int{(-(fi29)), (^(fi30))}
                    if (!false) {
                        var v32 float64 = float64(float32(-2.25))
                        if false {
                            fmt.Println(v20, fi29, fi30)
                        } else {
                            if ("a" != (("hello" + "a") + "a")) {
                                fmt.Println(v2, v3, fi29)
                            } else {
                                var v33 int = (fi29 - fi30)
                                v34 := GenS3{F0: 255}
                                _ = v33
                                _ = v34
                            }
                            v2 += v2
                            var v35 int8 = 8
                            _ = v35
                        }
                        _ = v32
                    }
                    _ = v31
                } else {
                    var v36 string = ""
                    var v37 uint64 = 42
                    v37 += (^(v37))
                    _ = v36
                    _ = v37
                }
                _ = fi30
            }
        default:
            for fi38 := 0; fi38 < 20; fi38++ {
                switch (fi38 ^ (fi38 | 9223372036854775807)) {
                case 127:
                    if (!true) {
                        var v39 int8 = -1
                        if (1 < 81) {
                            fmt.Println(fi29, v39, fi29)
                            if (42 != 209) {
                                switch "a" {
                                case "foo bar":
                                    var v40 int32 = (v2 ^ (v2 / ((v2) | 1)))
                                    _ = v40
                                case "hello":
                                    var v41 bool = (("\n\t" >= (("foo bar" + ("hello" + (("" + ("" + ("\n\t" + (("a" + ("a" + "a")) + "hello")))) + (("" + "foo bar") + "a")))) + "a")) && (2147483648 < 2))
                                    _ = v41
                                default:
                                    v2 = int32(fi38)
                                }
                                var v42 bool = true
                                _ = v42
                            }
                        }
                        if (("foo bar" == "hello") || (false || false)) {
                            switch 873 {
                            case 777:
                                if true {
                                    var v43 float32 = v3
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v44 int = fi38
                                        v44 += (v44 | 32768)
                                        _ = v44
                                    }
                                    var v45 uint8 = 127
                                    if ((!((fi29 * fi38) != fi38)) && (!false)) {
                                        v46 := GenS1{F0: (v20 | 1), F1: -9223372036854775808}
                                        v20 = (v20 | (v20 % ((v20) | 1)))
                                        var v47 int32 = v2
                                        _ = v46
                                        _ = v47
                                    }
                                    _ = len(v1)
                                    _ = v43
                                    _ = v45
                                }
                            case -1:
                                switch (fi38 / ((966) | 1)) {
                                case -2147483648:
                                    v20 -= 379
                                default:
                                    v48 := GenS2{F0: v39, F1: 415}
                                    _ = v48
                                }
                            }
                            v1[fi38] = fi38
                        }
                        var v49 string = ("hello" + ("" + "hello"))
                        _ = v39
                        _ = v49
                    } else {
                        v50 := GenS4{F0: 2}
                        v2 = 1
                        v51 := GenS2{F0: 66, F1: 1}
                        _ = v50
                        _ = v51
                    }
                    fmt.Println(fi29)
                case 4294967295:
                    switch -51 {
                    case -2147483648:
                        v52 := GenS5{F0: genFn0((-(fi38)), (v2 % ((v2) | 1)))}
                        _ = v52
                    }
                case 0:
                    v20 = (-(v20))
                }
                var v53 int16 = (v20 * (v20 ^ (v20 + (^(v20)))))
                delete(v1, 2147483647)
                v53 = (-(v20))
                var v54 int64 = 42
                _ = fi38
                _ = v53
                _ = v54
            }
        }
        _ = fi29
    }
    for fi55 := 0; fi55 < 14; fi55++ {
        if false {
            if false {
                v56 := []int{(fi55 & (fi55 & -1)), fi55}
                switch (fi55 << ((561) & 63)) {
                case 1:
                    for fi57 := 0; fi57 < 14; fi57++ {
                        var v58 int64 = 42
                        if ((v20 < v20) && false) {
                            switch (fi57 & (fi55 ^ int(v58))) {
                            case 2147483647:
                                if (255 <= 554) {
                                    v59 := GenS6{F0: v2, F1: 224, F2: 2147483647}
                                    var v60 string = ("hello" + "\n\t")
                                    _ = v59
                                    _ = v60
                                } else {
                                    v2 += (v2 / ((-1) | 1))
                                    switch 32767 {
                                    case -514:
                                        var v61 uint = 620
                                        _ = v61
                                    }
                                    v2 += -38
                                }
                            default:
                                v62 := GenS4{F0: 1}
                                _ = v62
                            }
                            var v63 float32 = (-2.25 + 0.0)
                            fmt.Println(fi55, fi57, v3)
                            _ = v63
                        } else {
                            var v64 uint16 = 42
                            if ((v2 + v2) < 42) {
                                var v65 bool = (!(!false))
                                fmt.Println(v2, v65)
                                fmt.Println(fi55)
                                _ = v65
                            }
                            fmt.Println(v20)
                            v66 := GenS7{F0: 2, F1: v58, F2: -128}
                            _ = v64
                            _ = v66
                        }
                        _ = fi57
                        _ = v58
                    }
                case -854:
                    if (127 != -11) {
                        v67 := GenS1{F0: v20, F1: -1}
                        _ = v67
                    } else {
                        if len(v56) >= 2 {
                            v56 = v56[1:len(v56)]
                        }
                        if (!(!false)) {
                            switch (-(fi55)) {
                            case 1:
                                v68 := GenS8{F0: 2, F1: fi55, F2: 42}
                                _ = v68
                            }
                            switch (fi55 << ((976) & 63)) {
                            case 32767:
                                fmt.Println(fi55, v2)
                            case -1:
                                for fi69 := 0; fi69 < 16; fi69++ {
                                    if ((v20 * (^(v20))) == (v20 ^ v20)) {
                                        v3 -= (v3 * 0.0)
                                        if true {
                                            switch 128 {
                                            case 973:
                                                v20 = 32767
                                                _ = len(v56)
                                            case -729:
                                                v2 = (-(v2))
                                            }
                                            if false {
                                                if true {
                                                    switch ("" + "") {
                                                    case "":
                                                        v2 = (v2 + v2)
                                                    }
                                                    var v70 uint16 = 42
                                                    v71 := GenS9{F0: -109}
                                                    var v72 uint32 = 278
                                                    _ = v70
                                                    _ = v71
                                                    _ = v72
                                                } else {
                                                    v3 -= v3
                                                    v56 = append(v56, (fi55 + (fi69 + (^(fi55)))))
                                                }
                                            }
                                        } else {
                                            v73 := GenS10{F0: (-(v20)), F1: 717}
                                            if true {
                                                v20 = v20
                                            } else {
                                                switch ("a" + "") {
                                                case "a":
                                                    var v74 bool = false
                                                    _ = v74
                                                case "":
                                                    if (127 < 2147483647) {
                                                        var v75 uint64 = 18446744073709551615
                                                        fmt.Println(v20)
                                                        _ = v75
                                                    } else {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            var v76 uint = 2
                                                            _ = v76
                                                        case "hello":
                                                            delete(v1, fi69)
                                                        }
                                                        v3 += v3
                                                    }
                                                    v77 := GenS3{F0: 2}
                                                    _ = v77
                                                default:
                                                    var v78 uint8 = 42
                                                    _ = v78
                                                }
                                                var v79 int64 = -136
                                                v80 := GenS11{F0: 209, F1: 42}
                                                _ = v79
                                                _ = v80
                                            }
                                            _ = v73
                                        }
                                        v20 -= 42
                                        var v81 int32 = v2
                                        if ((0 > v81) && (!(v3 != v3))) {
                                            var v82 int64 = 9223372036854775807
                                            _ = v82
                                        }
                                        _ = v81
                                    } else {
                                        var v83 int16 = int16(v2)
                                        switch 9223372036854775807 {
                                        case -438:
                                            v84 := v1[(fi69 | 128)]
                                            _ = v84
                                        case -1:
                                            v56 = append(v56, (fi55 & (fi69 % ((fi55) | 1))))
                                        case 32768:
                                            v20 += (v20 & v20)
                                        }
                                        _ = v83
                                    }
                                    v85 := GenS11{F0: 48, F1: 127}
                                    v86 := v1[(fi55 * (-(fi69)))]
                                    _ = fi69
                                    _ = v85
                                    _ = v86
                                }
                            case 9223372036854775807:
                                var v87 string = "foo bar"
                                _ = v87
                            default:
                                switch (("\n\t" + "hello") + "\n\t") {
                                case "":
                                    var v88 int8 = 100
                                    _ = v88
                                case "hello":
                                    for fi89 := 0; fi89 < 7; fi89++ {
                                        v90 := GenS7{F0: 18446744073709551615, F1: -599, F2: -128}
                                        var v91 uint16 = 0
                                        _ = fi89
                                        _ = v90
                                        _ = v91
                                    }
                                case "a":
                                    delete(v1, fi55)
                                default:
                                    v20 = int16(v2)
                                }
                            }
                            var v92 int32 = (v2 & (v2 & (^(v2))))
                            _ = v92
                        } else {
                            _ = len(v56)
                        }
                    }
                default:
                    v93 := GenS6{F0: (^(v2)), F1: 39, F2: 98}
                    _ = v93
                }
                var v94 uint8 = 207
                _ = v56
                _ = v94
            } else {
                v95 := GenS5{F0: 531}
                v20 = (v20 % ((v20) | 1))
                _ = v95
            }
            v2 += (v2 - -1)
        } else {
            v1[(fi55 + fi55)] = fi55
            for fi96 := 0; fi96 < 18; fi96++ {
                switch "a" {
                case "hello":
                    if false {
                        if ((!("hello" < "foo bar")) && true) {
                            if false {
                                var v97 uint32 = 65535
                                _ = v97
                            } else {
                                var v98 uint32 = 2147483647
                                _ = v98
                            }
                            var v99 int32 = -2147483648
                            var v100 uint16 = 735
                            _ = v100
                            _ = v99
                        } else {
                            v20 = (v20 / ((42) | 1))
                            v20 = (^(v20))
                        }
                        v2 = v2
                    } else {
                        v101 := v1[fi96]
                        if (!(!false)) {
                            switch "hello" {
                            case "\n\t":
                                var v102 uint = 42
                                var v103 int16 = v20
                                var v104 int64 = -1
                                _ = v102
                                _ = v103
                                _ = v104
                            }
                        }
                        _ = v101
                    }
                case "\n\t":
                    fmt.Println(v3, v3)
                case "a":
                    _ = len(v1)
                }
                var v105 bool = (!((false && false) || true))
                if (true && v105) {
                    var v106 uint32 = 292
                    if (!false) {
                        if ((((v20 & v20) < (v20 + (v20 / (((v20 & (v20 ^ v20))) | 1)))) && v105) && v105) {
                            _ = len(v1)
                            var v107 float64 = 0.0
                            v105 = (42 < 722)
                            var v108 float32 = (v3 + (v3 + (1.5 - 0.0)))
                            _ = v107
                            _ = v108
                        } else {
                            v2 += (^(v2))
                            v109 := GenS12{F0: 0, F1: 18446744073709551615}
                            var v110 int8 = -104
                            _ = v109
                            _ = v110
                        }
                        var v111 float64 = 1.5
                        var v112 string = (("\n\t" + "a") + "")
                        _ = v111
                        _ = v112
                    } else {
                        v113 := []int{(fi55 >> ((18446744073709551615) & 63)), (fi55 & -9223372036854775808)}
                        v114 := GenS13{F0: v105, F1: -470, F2: 65535}
                        var v115 float64 = 0.0
                        if v105 {
                            v3 -= v3
                            fmt.Println(v105, v115, fi55)
                        }
                        _ = v113
                        _ = v114
                        _ = v115
                    }
                    _ = v106
                }
                v105 = (true || (((("\n\t" + "foo bar") + ("" + "\n\t")) >= "a") && true))
                v116 := []int{(-(fi96)), (fi55 >> ((255) & 63)), 128, (-(fi96))}
                _ = fi96
                _ = v105
                _ = v116
            }
        }
        v117 := []int{fi55}
        for fi118 := 0; fi118 < 17; fi118++ {
            var v119 float64 = 1.5
            if (255 >= 18446744073709551615) {
                v119 = float64(-2.25)
                if false {
                    if (v3 < v3) {
                        v120 := GenS14{F0: (false && false), F1: 65535, F2: 753}
                        v2 = v2
                        v121 := v1[fi118]
                        _ = v120
                        _ = v121
                    }
                    fmt.Println(fi118, fi55)
                }
            } else {
                v119 += (v119 - ((-2.25 / v119) * v119))
                if true {
                    delete(v1, fi118)
                    if (false || ("" == "")) {
                        if len(v117) > 0 {
                            v117[0] = -2147483648
                        }
                        v122 := GenS9{F0: 42}
                        if ((v20 - -13) <= v20) {
                            fmt.Println(fi118)
                            v3 -= v3
                        } else {
                            var v123 float64 = v119
                            var v124 uint16 = 32767
                            _ = v123
                            _ = v124
                        }
                        _ = v122
                    } else {
                        switch (fi55 * fi118) {
                        case -1:
                            if len(v117) >= 2 {
                                v117 = v117[1:len(v117)]
                            }
                        }
                        var v125 int64 = 1
                        _ = v125
                    }
                } else {
                    var v126 int = fi55
                    var v127 int16 = (v20 & (^(v20)))
                    fmt.Println(v20, fi55)
                    var v128 int8 = 1
                    if true {
                        var v129 float32 = v3
                        switch "" {
                        case "hello":
                            v130 := GenS7{F0: 18446744073709551615, F1: int64(v2), F2: (-(v128))}
                            _ = v130
                        default:
                            v126 = 9223372036854775807
                        }
                        v128 = v128
                        _ = v129
                    }
                    var v131 uint32 = 325
                    if (!(true || false)) {
                        v132 := v1[fi118]
                        if ((^(v2)) >= (v2 + (-(v2)))) {
                            var v133 uint16 = 255
                            v134 := GenS2{F0: -28, F1: v133}
                            fmt.Println(v126, fi55, v133)
                            _ = v133
                            _ = v134
                        }
                        var v135 int8 = (v128 % ((v128) | 1))
                        _ = v132
                        _ = v135
                    }
                    _ = v126
                    _ = v127
                    _ = v128
                    _ = v131
                }
                if true {
                    v136 := GenS14{F0: false, F1: 42, F2: 4294967295}
                    var v137 bool = false
                    _ = v136
                    _ = v137
                }
            }
            _ = fi118
            _ = v119
        }
        _ = fi55
        _ = v117
    }
    for fi138 := 0; fi138 < 9; fi138++ {
        v139 := GenS15{F0: (fi138 + fi138), F1: 26}
        var v140 uint = 1
        _ = fi138
        _ = v139
        _ = v140
    }
    v2 = (v2 >> ((18446744073709551615) & 63))
    v20 = (v20 | 762)
    switch 32768 {
    case 127:
        switch "hello" {
        case "a":
            fmt.Println(v2)
        case "\n\t":
            if false {
                if true {
                    v20 += v20
                    v3 += 1e10
                }
                var v141 float32 = v3
                if true {
                    v141 = v141
                    for fi142 := 0; fi142 < 16; fi142++ {
                        if true {
                            v2 -= v2
                        }
                        v143 := v1[fi142]
                        _ = fi142
                        _ = v143
                    }
                    v20 = (v20 ^ v20)
                } else {
                    if true {
                        for fi144 := 0; fi144 < 4; fi144++ {
                            v2 = (-(v2))
                            for fi145 := 0; fi145 < 10; fi145++ {
                                var v146 string = ""
                                fmt.Println(v3)
                                _ = fi145
                                _ = v146
                            }
                            _ = fi144
                        }
                        for fi147 := 0; fi147 < 5; fi147++ {
                            for fi148 := 0; fi148 < 15; fi148++ {
                                switch (-(fi147)) {
                                case 1:
                                    var v149 int16 = v20
                                    _ = v149
                                case -488:
                                    delete(v1, fi147)
                                default:
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v150 uint8 = 52
                                        _ = v150
                                    case "a":
                                        v151 := GenS7{F0: 18446744073709551615, F1: 9223372036854775807, F2: int8(fi147)}
                                        switch (fi148 + fi148) {
                                        case 32768:
                                            v3 -= 3.14159
                                        case -1:
                                            var v152 uint32 = 967
                                            _ = v152
                                        }
                                        v153 := []int{}
                                        _ = v151
                                        _ = v153
                                    default:
                                        v154 := GenS16{F0: "a", F1: v2}
                                        _ = v154
                                    }
                                }
                                if false {
                                    fmt.Println(v141)
                                }
                                var v155 int32 = (v2 ^ (v2 << ((982) & 63)))
                                _ = fi148
                                _ = v155
                            }
                            _ = fi147
                        }
                    } else {
                        var v156 bool = (1e10 == float64(((float32((0.1 / 1e10)) * float32(float64((0.1 / v141)))) / v3)))
                        v157 := GenS5{F0: 613}
                        var v158 uint32 = 494
                        _ = v156
                        _ = v157
                        _ = v158
                    }
                    var v159 uint16 = 202
                    _ = v159
                }
                _ = v141
            } else {
                v20 = v20
                fmt.Println(v20, v3)
            }
        }
    case -864:
        var v160 int32 = v2
        _ = v160
    default:
        for fi161 := 0; fi161 < 3; fi161++ {
            for fi162 := 0; fi162 < 18; fi162++ {
                v163 := GenS17{F0: 304}
                _ = fi162
                _ = v163
            }
            for fi164 := 0; fi164 < 11; fi164++ {
                v1[fi161] = fi161
                _ = fi164
            }
            if true {
                v20 += (v20 & (v20 ^ v20))
            } else {
                v20 = v20
                if true {
                    if false {
                        var v165 string = ("foo bar" + "a")
                        switch (("a" + v165) + (v165 + ((v165 + "hello") + ("foo bar" + v165)))) {
                        case "hello":
                            v166 := []int{fi161, fi161}
                            _ = v166
                        }
                        var v167 int32 = (-(v2))
                        _ = v165
                        _ = v167
                    } else {
                        var v168 int16 = 936
                        for fi169 := 0; fi169 < 14; fi169++ {
                            delete(v1, (^(fi169)))
                            v170 := v1[(fi169 << ((401) & 63))]
                            if (v20 > (v168 + int16(v2))) {
                                fmt.Println(v170)
                                v3 = 1e10
                            } else {
                                var v171 int = fi169
                                v172 := GenS1{F0: (v168 | (v20 - 0)), F1: 0}
                                _ = v171
                                _ = v172
                            }
                            v168 = 1
                            _ = fi169
                            _ = v170
                        }
                        v3 = (1.5 + v3)
                        _ = v168
                    }
                } else {
                    if (!((v3 * float32(1e10)) > v3)) {
                        switch 0 {
                        case 95:
                            v173 := GenS15{F0: (fi161 - (-(fi161))), F1: 2147483648}
                            _ = v173
                        case 127:
                            v3 = v3
                        case -645:
                            var v174 int16 = 1000
                            _ = v174
                        }
                    }
                    v175 := []int{fi161, (fi161 * fi161)}
                    _ = v175
                }
                v20 = v20
            }
            _ = fi161
        }
    }
    _ = len(v1)
    var v176 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v176:", v176)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v176
    _ = v2
    _ = v20
    _ = v3
}
