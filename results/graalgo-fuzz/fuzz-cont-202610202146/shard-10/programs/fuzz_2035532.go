package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int8; F1 int64; F2 uint64 }
type GenS5 struct { F0 int16; F1 uint32; F2 int64 }
type GenS6 struct { F0 uint; F1 bool; F2 int }
type GenS7 struct { F0 uint; F1 uint }
type GenS8 struct { F0 uint16; F1 uint; F2 int64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int32; F1 int8 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int32; F1 int64 }
type GenS15 struct { F0 int32; F1 int32 }
type GenS16 struct { F0 uint; F1 int64; F2 uint32 }
type GenS17 struct { F0 uint }
func genFn0(p0 int16, p1 int8) uint32 {
    p1 = p1
    var v5 uint = 42
    p1 += 31
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1(p0 string) int {
    fmt.Println(p0)
    fmt.Println(p0)
    switch p0 {
    case "foo bar":
        v6 := []int{}
        _ = v6
    case "hello":
        v7 := GenS1{F0: p0, F1: 1, F2: 2147483648}
        _ = v7
    case "a":
        v8 := []int{1, -2147483648}
        _ = v8
    }
    if (530 > 1) {
        p0 = p0
        v9 := GenS1{F0: "", F1: 42, F2: 647}
        _ = v9
    }
    if false {
        p0 = "hello"
        p0 = p0
    } else {
        p0 = ((((p0 + p0) + "foo bar") + "hello") + p0)
        v10 := []int{127, -2147483648}
        switch 42 {
        case 2147483647:
            var v11 float64 = 3.14159
            _ = v11
        case 922:
            v12 := GenS1{F0: (p0 + p0), F1: 2147483647, F2: 2}
            _ = v12
        }
        _ = v10
    }
    _ = p0
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -93
    v3 := make(map[int]string)
    v4 := []uint32{}
    v2 = -117
    fmt.Println(v2)
    switch 2147483647 {
    case 128:
        var v13 uint16 = 2
        _ = v13
    default:
        switch 704 {
        case 42:
            var v14 int = 4294967295
            _ = v14
        case 32767:
            if ((v2 == 127) && (!(("hello" + (("\n\t" + "") + "foo bar")) > ("\n\t" + "\n\t")))) {
                var v15 uint64 = 1
                v16 := GenS1{F0: ("hello" + ("a" + ("\n\t" + ""))), F1: 909, F2: 42}
                if ("\n\t" > "hello") {
                    var v17 uint64 = 42
                    v18 := GenS2{F0: 406}
                    var v19 int64 = 546
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    v15 = v15
                    v20 := GenS3{F0: 176, F1: -1}
                    var v21 float64 = 3.14159
                    _ = v20
                    _ = v21
                }
                _ = v15
                _ = v16
            }
        default:
            if false {
                var v22 uint8 = 255
                v2 = (-(v2))
                if (4294967295 >= genFn0(481, v2)) {
                    var v23 float32 = (1.5 - (1e10 + 0.0))
                    if true {
                        for fi24 := 0; fi24 < 10; fi24++ {
                            if true {
                                fmt.Println(fi24)
                                v23 = 0.1
                            } else {
                                for fi25 := 0; fi25 < 14; fi25++ {
                                    v26 := GenS3{F0: 0, F1: 0}
                                    var v27 uint = uint(v22)
                                    v27 = v27
                                    var v28 uint32 = 71
                                    if false {
                                        v29 := GenS3{F0: 378, F1: 592}
                                        if true {
                                            switch genFn1(("foo bar" + "a")) {
                                            case -2147483648:
                                                if true {
                                                    if (v28 < (v28 % ((v28) | 1))) {
                                                        v23 = (v23 * v23)
                                                        v2 = v2
                                                    } else {
                                                        v30 := GenS3{F0: 788, F1: -689}
                                                        _ = v30
                                                    }
                                                }
                                            case 1:
                                                var v31 int32 = int32(v2)
                                                _ = v31
                                            default:
                                                var v32 int = (fi25 - (fi24 & -2147483648))
                                                _ = v32
                                            }
                                            v28 -= (v28 * (v28 | (v28 << ((2) & 63))))
                                        }
                                        _ = v29
                                    }
                                    var v33 int64 = 9223372036854775807
                                    if (true && ((((v23 * 1e10) * float32(1.5)) == 1e10) && (!((fi25 - (fi25 | int(v27))) < (fi25 ^ (fi24 / (((fi25 - (fi25 % ((int(v2)) | 1)))) | 1))))))) {
                                        if false {
                                            v34 := GenS2{F0: (v28 * v28)}
                                            var v35 int32 = int32(fi25)
                                            _ = v34
                                            _ = v35
                                        } else {
                                            switch "\n\t" {
                                            case "hello":
                                                var v36 uint32 = (v28 / ((genFn0(-32768, v2)) | 1))
                                                _ = v36
                                            case "foo bar":
                                                v22 -= v22
                                            case "":
                                                var v37 uint32 = (v28 & (v28 >> (((v27 / ((v27) | 1))) & 63)))
                                                _ = v37
                                            }
                                            switch "" {
                                            case "foo bar":
                                                var v38 int = fi24
                                                _ = v38
                                            case "hello":
                                                v2 += 100
                                            default:
                                                v27 += (v27 | v27)
                                            }
                                            var v39 float64 = 1e10
                                            _ = v39
                                        }
                                        v40 := GenS3{F0: 42, F1: -9223372036854775808}
                                        fmt.Println(v28)
                                        _ = v40
                                    }
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                    _ = v33
                                }
                                var v41 int32 = 1
                                for fi42 := 0; fi42 < 15; fi42++ {
                                    _ = len(v1)
                                    v43 := GenS4{F0: (v2 * (v2 & 127)), F1: 0, F2: 42}
                                    if ((3.14159 - (1.5 + 1.5)) < 1e10) {
                                        var v44 int8 = (-(v2))
                                        var v45 uint = 255
                                        _ = v44
                                        _ = v45
                                    }
                                    switch (("" + ("" + "hello")) + "a") {
                                    case "\n\t":
                                        var v46 string = "\n\t"
                                        _ = v46
                                    default:
                                        switch (fi42 % ((2147483647) | 1)) {
                                        case -112:
                                            var v47 int8 = (v2 * (v2 * v2))
                                            _ = v47
                                        case -1:
                                            var v48 int32 = v41
                                            _ = v48
                                        }
                                    }
                                    fmt.Println(v41, v22, v2)
                                    _ = fi42
                                    _ = v43
                                }
                                _ = v41
                            }
                            if false {
                                var v49 int8 = -1
                                if true {
                                    v2 -= (v2 % (((v49 << ((1) & 63))) | 1))
                                    var v50 int32 = -132
                                    _ = v50
                                } else {
                                    if false {
                                        if ((((0.1 / ((0.1 + 0.0) * -2.25)) * 1.5) * ((float64((v23 + 1e10)) * (float64(0.0) / 3.14159)) + 1.5)) > (0.0 + 0.1)) {
                                            var v51 float64 = (3.14159 + 0.0)
                                            if false {
                                                v2 += (v49 << ((18446744073709551615) & 63))
                                            }
                                            v51 = v51
                                            _ = v51
                                        }
                                        var v52 int = genFn1("\n\t")
                                        _ = v52
                                    } else {
                                        v53 := GenS5{F0: 566, F1: 65535, F2: -9223372036854775808}
                                        _ = v53
                                    }
                                }
                                if true {
                                    v49 = (v49 | (v49 - v2))
                                    if true {
                                        var v54 uint64 = 39
                                        _ = v54
                                    }
                                    v2 += 127
                                }
                                _ = v49
                            }
                            if (!(!true)) {
                                var v55 int16 = 0
                                v2 = (v2 << ((1) & 63))
                                var v56 int64 = 9223372036854775807
                                _ = v55
                                _ = v56
                            } else {
                                if false {
                                    if true {
                                        v22 = v22
                                        fmt.Println(v2, fi24)
                                    } else {
                                        if (((fi24 | (fi24 << ((945) & 63))) != fi24) || false) {
                                            v23 += v23
                                            if true {
                                                var v57 float32 = v23
                                                _ = v57
                                            } else {
                                                v22 += (v22 - v22)
                                                v22 -= (v22 << ((705) & 63))
                                            }
                                        } else {
                                            if false {
                                                v23 -= v23
                                                for fi58 := 0; fi58 < 10; fi58++ {
                                                    if len(v4) > 0 {
                                                        v4[0] = 0
                                                    }
                                                    _ = fi58
                                                }
                                                var v59 uint16 = 32767
                                                _ = v59
                                            }
                                            for fi60 := 0; fi60 < 6; fi60++ {
                                                v23 += v23
                                                if len(v4) > 0 {
                                                    v4[0] = 65535
                                                }
                                                _ = fi60
                                            }
                                            if (false || ("hello" > ("a" + ("hello" + (("hello" + "") + (("a" + "foo bar") + ("\n\t" + "a"))))))) {
                                                for fi61 := 0; fi61 < 6; fi61++ {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v23 = v23
                                                    _ = fi61
                                                }
                                                switch (fi24 >> ((1) & 63)) {
                                                case 32767:
                                                    v62 := GenS4{F0: (v2 * v2), F1: int64(v22), F2: 1}
                                                    _ = v62
                                                case -79:
                                                    if false {
                                                        for fi63 := 0; fi63 < 12; fi63++ {
                                                            fmt.Println(v2)
                                                            _ = fi63
                                                        }
                                                    } else {
                                                        delete(v1, 65535)
                                                        v22 -= v22
                                                    }
                                                }
                                            }
                                        }
                                        var v64 float64 = 1e10
                                        v22 += 255
                                        _ = v64
                                    }
                                }
                                v22 = v22
                                var v65 float64 = -2.25
                                _ = v65
                            }
                            _ = fi24
                        }
                        fmt.Println(v23, v23, v2)
                        for fi66 := 0; fi66 < 4; fi66++ {
                            v23 = (float32(0.0) * v23)
                            _ = fi66
                        }
                    } else {
                        if true {
                            v67 := GenS2{F0: 2147483647}
                            if true {
                                v68 := GenS6{F0: 255, F1: false, F2: genFn1("foo bar")}
                                for fi69 := 0; fi69 < 1; fi69++ {
                                    fmt.Println(fi69, fi69, v22)
                                    v70 := v1[2]
                                    switch "" {
                                    case "":
                                        for fi71 := 0; fi71 < 2; fi71++ {
                                            var v72 bool = (!(-481 > 687))
                                            if v70 {
                                                _ = len(v4)
                                            } else {
                                                v3[(fi69 % (((fi69 % ((fi71) | 1))) | 1))] = ("" + ("" + ""))
                                                var v73 string = "hello"
                                                var v74 int32 = 0
                                                _ = v73
                                                _ = v74
                                            }
                                            _ = len(v4)
                                            switch (fi69 * fi71) {
                                            case 128:
                                                v75 := GenS5{F0: 32767, F1: 2147483647, F2: 0}
                                                _ = v75
                                            case 0:
                                                v4 = append(v4, 22)
                                            }
                                            _ = fi71
                                            _ = v72
                                        }
                                    case "foo bar":
                                        switch genFn1("") {
                                        case -910:
                                            v76 := GenS7{F0: 2, F1: 565}
                                            _ = v76
                                        case -997:
                                            for fi77 := 0; fi77 < 12; fi77++ {
                                                var v78 uint32 = 42
                                                v78 -= v78
                                                var v79 uint32 = v78
                                                v23 -= v23
                                                _ = fi77
                                                _ = v78
                                                _ = v79
                                            }
                                        default:
                                            fmt.Println(v2)
                                        }
                                    }
                                    _ = fi69
                                    _ = v70
                                }
                                _ = v68
                            } else {
                                fmt.Println(v22)
                                for fi80 := 0; fi80 < 10; fi80++ {
                                    v81 := v3[(^(fi80))]
                                    if ((0.1 * 0.1) > v23) {
                                        for fi82 := 0; fi82 < 13; fi82++ {
                                            v83 := GenS2{F0: 2147483647}
                                            v84 := GenS3{F0: 42, F1: -9223372036854775808}
                                            v22 = v22
                                            _ = fi82
                                            _ = v83
                                            _ = v84
                                        }
                                    }
                                    fmt.Println(v81, v22, fi80)
                                    _ = fi80
                                    _ = v81
                                }
                            }
                            var v85 uint = 2
                            _ = v67
                            _ = v85
                        } else {
                            var v86 uint16 = uint16(v22)
                            v87 := GenS5{F0: -32768, F1: 240, F2: -9223372036854775808}
                            for fi88 := 0; fi88 < 7; fi88++ {
                                v89 := GenS6{F0: 255, F1: true, F2: (fi88 << ((255) & 63))}
                                for fi90 := 0; fi90 < 6; fi90++ {
                                    switch fi88 {
                                    case 9223372036854775807:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v22 = 42
                                    case -2147483648:
                                        var v91 int64 = 300
                                        var v92 uint = 1
                                        var v93 uint = (v92 / ((v92) | 1))
                                        _ = v91
                                        _ = v92
                                        _ = v93
                                    case -1:
                                        v94 := GenS4{F0: -58, F1: 42, F2: 0}
                                        _ = v94
                                    default:
                                        fmt.Println(v86)
                                        fmt.Println(v2, v86, v86)
                                    }
                                    _ = fi90
                                }
                                v2 = v2
                                _ = fi88
                                _ = v89
                            }
                            _ = v86
                            _ = v87
                        }
                        var v95 int = -1
                        _ = v95
                    }
                    _ = v23
                } else {
                    if true {
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                        for fi96 := 0; fi96 < 16; fi96++ {
                            fmt.Println(v22)
                            _ = fi96
                        }
                    }
                    var v97 int8 = (v2 + 0)
                    fmt.Println(v22, v97)
                    _ = v97
                }
                _ = v22
            }
        }
    }
    var v98 float64 = 0.1
    v99 := GenS8{F0: 42, F1: 121, F2: -1}
    var v100 float64 = 1.5
    fmt.Println(v2, v100)
    var v101 uint16 = 1
    var v102 uint32 = 2147483648
    var v103 float64 = v100
    switch 42 {
    case 637:
        var v104 uint32 = 42
        _ = v104
    default:
        v105 := GenS5{F0: 1000, F1: v102, F2: 9223372036854775807}
        _ = v105
    }
    for fi106 := 0; fi106 < 1; fi106++ {
        if (458 <= 0) {
            v107 := GenS9{F0: 111}
            if false {
                if ((-9223372036854775808 != 954) && ((v103 + v98) == (1.5 * float64(1.5)))) {
                    if true {
                        fmt.Println(fi106)
                        v108 := GenS10{F0: fi106}
                        _ = v108
                    } else {
                        if false {
                            switch "hello" {
                            case "foo bar":
                                if (1 != v101) {
                                    for fi109 := 0; fi109 < 14; fi109++ {
                                        v110 := GenS1{F0: "a", F1: 2147483647, F2: v102}
                                        if len(v4) > 0 {
                                            v4[0] = (v102 << ((255) & 63))
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = fi109
                                        _ = v110
                                    }
                                    for fi111 := 0; fi111 < 2; fi111++ {
                                        var v112 float32 = (((1.5 * 0.0) - (1.5 - float32((1e10 * v98)))) + (float32(1.5) - float32(((v98 + v103) - (v98 * 1e10)))))
                                        switch "" {
                                        case "hello":
                                            if (4294967295 >= 255) {
                                                v4 = append(v4, genFn0(1000, (v2 ^ int8(fi106))))
                                                if (false && false) {
                                                    v3[(fi111 & genFn1(""))] = "\n\t"
                                                }
                                            } else {
                                                var v113 int32 = 0
                                                _ = v113
                                            }
                                        default:
                                            v114 := GenS9{F0: 2}
                                            _ = v114
                                        }
                                        _ = fi111
                                        _ = v112
                                    }
                                    if (0.1 < 1.5) {
                                        var v115 float64 = float64(1e10)
                                        _ = len(v4)
                                        fmt.Println(v100, v98)
                                        _ = v115
                                    }
                                }
                            case "":
                                if true {
                                    v102 += (^(v102))
                                    for fi116 := 0; fi116 < 7; fi116++ {
                                        v117 := GenS3{F0: 479, F1: 0}
                                        _ = fi116
                                        _ = v117
                                    }
                                    var v118 int16 = 32767
                                    _ = v118
                                }
                            case "\n\t":
                                for fi119 := 0; fi119 < 10; fi119++ {
                                    var v120 int8 = (v2 | (v2 >> ((4294967295) & 63)))
                                    v102 = 627
                                    _ = fi119
                                    _ = v120
                                }
                            default:
                                var v121 int = 32768
                                _ = v121
                            }
                        } else {
                            if (v2 >= -45) {
                                var v122 uint32 = (^(v102))
                                _ = v122
                            }
                            v123 := GenS11{F0: -1}
                            if true {
                                _ = len(v4)
                                var v124 uint8 = 2
                                _ = v124
                            } else {
                                var v125 uint32 = (^(v102))
                                v4 = append(v4, (v102 | (v125 + v125)))
                                switch (fi106 / ((42) | 1)) {
                                case 32767:
                                    if (0 < 181) {
                                        v125 -= uint32(v101)
                                        if len(v4) > 0 {
                                            v4[0] = 241
                                        }
                                        v126 := GenS1{F0: "", F1: -1, F2: (^(v102))}
                                        _ = v126
                                    } else {
                                        fmt.Println(v100)
                                        var v127 float32 = ((-2.25 - 1e10) * 3.14159)
                                        var v128 bool = true
                                        _ = v127
                                        _ = v128
                                    }
                                case 128:
                                    fmt.Println(v98)
                                case -874:
                                    v129 := GenS11{F0: 246}
                                    _ = v129
                                }
                                _ = v125
                            }
                            _ = v123
                        }
                        var v130 bool = true
                        _ = v130
                    }
                    for fi131 := 0; fi131 < 15; fi131++ {
                        v132 := v3[fi131]
                        var v133 int = -2147483648
                        if false {
                            var v134 int = genFn1((("hello" + "hello") + v132))
                            var v135 int16 = -255
                            v135 = v135
                            var v136 float64 = float64((1.5 - 0.0))
                            var v137 bool = (-968 == -1)
                            v98 -= ((v100 - (v100 * v100)) + (v100 + ((v100 + float64(0.0)) * (float64(((1e10 * (1.5 + 1.5)) * 0.1)) / ((0.0 * 0.1) + ((1.5 + (-2.25 + float64(3.14159))) - v103))))))
                            _ = v134
                            _ = v135
                            _ = v136
                            _ = v137
                        }
                        _ = fi131
                        _ = v132
                        _ = v133
                    }
                }
                switch fi106 {
                case 42:
                    var v138 uint8 = 0
                    _ = v138
                case 32768:
                    for fi139 := 0; fi139 < 20; fi139++ {
                        v140 := GenS6{F0: 266, F1: (((1 == 169) && (0 > -1)) && (!false)), F2: fi139}
                        v98 = (v100 / v100)
                        switch "\n\t" {
                        case "":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            fmt.Println(v102, v102, fi106)
                        }
                        _ = fi139
                        _ = v140
                    }
                case -322:
                    v98 = 0.1
                default:
                    switch "hello" {
                    case "":
                        switch "a" {
                        case "hello":
                            if len(v4) > 0 {
                                v4[0] = (v102 % (((v102 * v102)) | 1))
                            }
                        case "":
                            v101 = (v101 + (v101 * v101))
                        default:
                            _ = len(v1)
                        }
                    case "a":
                        if true {
                            var v141 float32 = -2.25
                            fmt.Println(v100)
                            _ = v141
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            switch genFn1("hello") {
                            case 32768:
                                fmt.Println(v100, v2, v98)
                            case 42:
                                v142 := GenS4{F0: (v2 ^ (v2 * (v2 + int8(v102)))), F1: 494, F2: 424}
                                _ = v142
                            }
                            for fi143 := 0; fi143 < 20; fi143++ {
                                fmt.Println(fi143, fi143, v103)
                                var v144 float32 = 1.5
                                _ = fi143
                                _ = v144
                            }
                        }
                    case "foo bar":
                        fmt.Println(v101, v98, v98)
                    }
                }
            } else {
                v98 -= ((v103 - 1e10) - (float64((1e10 + (1e10 - 1.5))) - ((((0.1 * (3.14159 / 1.0)) / (1e10 + float64(3.14159))) - 3.14159) * v100)))
            }
            var v145 float32 = 1.5
            _ = v107
            _ = v145
        } else {
            if false {
                var v146 int64 = -1
                v103 -= (v98 / ((v100 + 1e10) * (v98 - ((3.14159 - 0.0) + (v100 + v103)))))
                _ = v146
            } else {
                var v147 uint64 = 2
                v100 -= -2.25
                if false {
                    if len(v4) > 0 {
                        v4[0] = 4294967295
                    }
                    v103 += (-2.25 - 1.5)
                } else {
                    v148 := GenS11{F0: -2147483648}
                    for fi149 := 0; fi149 < 4; fi149++ {
                        v2 += v2
                        switch ("foo bar" + (("foo bar" + ("a" + "foo bar")) + "\n\t")) {
                        case "":
                            fmt.Println(v101, fi149)
                        case "\n\t":
                            v101 = v101
                        default:
                            v103 = (float64(0.0) + 0.1)
                        }
                        _ = fi149
                    }
                    _ = v148
                }
                _ = v147
            }
            if false {
                v150 := GenS9{F0: 244}
                v100 = v103
                var v151 uint = 65535
                _ = v150
                _ = v151
            }
            if true {
                switch fi106 {
                case 2147483647:
                    fmt.Println(v102, v2)
                case 0:
                    var v152 float32 = (float32(v103) - 1.5)
                    _ = v152
                }
                v153 := GenS12{F0: -2147483648, F1: v2}
                v154 := GenS13{F0: 4294967295}
                _ = v153
                _ = v154
            }
        }
        for fi155 := 0; fi155 < 15; fi155++ {
            v156 := GenS14{F0: 345, F1: -9223372036854775808}
            var v157 int16 = -1
            _ = fi155
            _ = v156
            _ = v157
        }
        for fi158 := 0; fi158 < 7; fi158++ {
            fmt.Println(v100)
            _ = len(v3)
            if ((57 > v2) || ((false && ((670 == 42) || (4294967295 > 4294967295))) && false)) {
                fmt.Println(v103, fi106)
                var v159 uint8 = 42
                _ = v159
            }
            _ = fi158
        }
        _ = fi106
    }
    var v160 int32 = -1
    var v161 uint8 = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (255 >= 18446744073709551615) {
        var v162 uint16 = v101
        switch 32767 {
        case 32768:
            for fi163 := 0; fi163 < 13; fi163++ {
                v2 = int8(v102)
                _ = fi163
            }
        }
        var v164 uint8 = v161
        _ = v162
        _ = v164
    }
    for fi165 := 0; fi165 < 16; fi165++ {
        var v166 int64 = 350
        v3[4294967295] = "a"
        if false {
            fmt.Println(v160)
            if false {
                var v167 int32 = int32(fi165)
                var v168 int64 = (-(v166))
                _ = v167
                _ = v168
            } else {
                var v169 int16 = 1000
                _ = v169
            }
            v166 -= int64(v101)
        } else {
            switch 381 {
            case 128:
                for fi170 := 0; fi170 < 12; fi170++ {
                    fmt.Println(fi165, v166)
                    if len(v4) > 0 {
                        v4[0] = v102
                    }
                    _ = fi170
                }
            case -989:
                v171 := v3[790]
                _ = v171
            case -892:
                switch fi165 {
                case 32768:
                    v172 := v3[2147483647]
                    _ = v172
                case -818:
                    var v173 int32 = int32(v161)
                    _ = v173
                case -696:
                    var v174 uint16 = (v101 & (v101 | uint16(v161)))
                    _ = v174
                }
            }
            if false {
                for fi175 := 0; fi175 < 20; fi175++ {
                    if (!(!true)) {
                        var v176 int8 = (v2 << ((1) & 63))
                        v177 := GenS15{F0: (v160 + 2147483647), F1: v160}
                        if (!false) {
                            switch int(v102) {
                            case -9223372036854775808:
                                var v178 int16 = 1
                                _ = v178
                            default:
                                if true {
                                    var v179 uint64 = 1
                                    var v180 int = (fi165 ^ genFn1(("\n\t" + "")))
                                    var v181 int64 = (v166 ^ (v166 + v166))
                                    _ = v179
                                    _ = v180
                                    _ = v181
                                } else {
                                    var v182 uint64 = 18446744073709551615
                                    v2 -= 23
                                    var v183 int64 = (v166 >> ((504) & 63))
                                    v2 = (v176 << ((255) & 63))
                                    _ = v182
                                    _ = v183
                                }
                            }
                            fmt.Println(v103)
                        } else {
                            v184 := GenS9{F0: v161}
                            _ = v184
                        }
                        _ = v176
                        _ = v177
                    } else {
                        if (fi175 < fi165) {
                            var v185 uint64 = 1
                            _ = v185
                        }
                        switch "hello" {
                        case "":
                            v160 += -2147483648
                            switch (fi165 | fi175) {
                            case 817:
                                switch ("hello" + "") {
                                case "a":
                                    v103 += float64(((1e10 / 3.14159) / 3.14159))
                                    fmt.Println(v102, v2)
                                    switch 32768 {
                                    case 32768:
                                        v186 := GenS16{F0: uint(v161), F1: 264, F2: v102}
                                        _ = v186
                                    case -2147483648:
                                        var v187 float32 = 1.5
                                        v102 += 2147483647
                                        _ = v187
                                    default:
                                        var v188 int32 = 42
                                        _ = v188
                                    }
                                case "\n\t":
                                    var v189 uint8 = uint8(v101)
                                    _ = v189
                                case "":
                                    v100 -= v98
                                default:
                                    var v190 float64 = v103
                                    _ = v190
                                }
                            case 1:
                                v161 += uint8(v102)
                            case 0:
                                fmt.Println(v103, fi175)
                            default:
                                if ((v161 >= 255) || (!false)) {
                                    var v191 uint32 = v102
                                    _ = v191
                                }
                            }
                            if (-440 >= int(v166)) {
                                switch ((("" + ("\n\t" + "hello")) + ("\n\t" + (("hello" + "") + ("foo bar" + "foo bar")))) + "hello") {
                                case "":
                                    v4 = append(v4, 2147483647)
                                case "a":
                                    var v192 string = (("" + "a") + "")
                                    _ = v192
                                case "foo bar":
                                    if true {
                                        v193 := GenS17{F0: 4294967295}
                                        _ = v193
                                    } else {
                                        v102 = (v102 ^ v102)
                                    }
                                default:
                                    var v194 uint = 2
                                    _ = v194
                                }
                                fmt.Println(v102)
                                var v195 bool = true
                                _ = v195
                            }
                            v101 = (v101 - 0)
                            v98 = v103
                            var v196 float64 = 1e10
                            switch fi165 {
                            case 663:
                                if false {
                                    v197 := GenS12{F0: (v160 % (((-(v160))) | 1)), F1: (v2 / (((v2 | 42)) | 1))}
                                    v4 = append(v4, genFn0(32767, v2))
                                    _ = v197
                                } else {
                                    v100 += v196
                                }
                            default:
                                if ((v161 + v161) > (v161 & v161)) {
                                    var v198 uint64 = 806
                                    _ = v198
                                } else {
                                    v199 := GenS9{F0: v161}
                                    _ = v199
                                }
                            }
                            fmt.Println(v102)
                            if (!(genFn1(("a" + ("hello" + "hello"))) == fi165)) {
                                v102 = 43
                            } else {
                                var v200 int8 = v2
                                var v201 uint64 = 0
                                _ = v200
                                _ = v201
                            }
                            _ = v196
                        default:
                            if len(v4) > 0 {
                                v4[0] = (^(v102))
                            }
                        }
                    }
                    var v202 uint16 = v101
                    v2 = (v2 % ((v2) | 1))
                    _ = fi175
                    _ = v202
                }
                var v203 int64 = (v166 % ((v166) | 1))
                _ = v203
            }
        }
        _ = fi165
        _ = v166
    }
    var v204 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v160:", v160)
    fmt.Println("v161:", v161)
    fmt.Println("v2:", v2)
    fmt.Println("v204:", v204)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v100
    _ = v101
    _ = v102
    _ = v103
    _ = v160
    _ = v161
    _ = v2
    _ = v204
    _ = v3
    _ = v4
    _ = v98
    _ = v99
}
