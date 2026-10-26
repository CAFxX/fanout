package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int32; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 bool; F2 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 uint32 }
type GenS8 struct { F0 int; F1 uint32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int16; F1 uint; F2 uint32 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 string; F1 uint64 }
type GenS14 struct { F0 uint16; F1 uint16; F2 uint64 }
func genFn0() uint8 {
    v6 := []int{-9223372036854775808, 32767, 32768, 793}
    var v7 int16 = 1000
    _ = v6
    _ = v7
    return 1
}
func genFn1() int8 {
    var v8 uint = 42
    v8 = 889
    fmt.Println(v8)
    v9 := []int{}
    v8 = 295
    _ = v8
    _ = v9
    return -128
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint = 344
    var v3 float64 = (1.5 - (1.5 - 3.14159))
    var v4 int32 = 834
    var v5 int = int(v4)
    for fi10 := 0; fi10 < 9; fi10++ {
        if true {
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 int8 = 0
                _ = fi11
                _ = v12
            }
        } else {
            fmt.Println(fi10)
        }
        v13 := GenS1{F0: 0, F1: fi10}
        _ = fi10
        _ = v13
    }
    v14 := GenS2{F0: v5, F1: 255}
    var v15 uint16 = 32767
    v4 += (v4 ^ v4)
    for fi16 := 0; fi16 < 20; fi16++ {
        v17 := GenS1{F0: 2, F1: fi16}
        _ = fi16
        _ = v17
    }
    switch "hello" {
    case "\n\t":
        if false {
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 string = "foo bar"
                switch v19 {
                case "foo bar":
                    if false {
                        switch (v19 + "a") {
                        case "foo bar":
                            var v20 float32 = (v1 / float32((v3 * v3)))
                            _ = v20
                        case "a":
                            var v21 int = 2147483647
                            _ = v21
                        }
                        fmt.Println(v4, v1)
                    }
                case "hello":
                    var v22 string = v19
                    _ = v22
                case "":
                    if true {
                        for fi23 := 0; fi23 < 12; fi23++ {
                            if true {
                                switch v19 {
                                case "foo bar":
                                    var v24 string = "hello"
                                    _ = v24
                                case "hello":
                                    switch v19 {
                                    case "foo bar":
                                        if (635 <= v2) {
                                            fmt.Println(v1)
                                            v25 := []int{v5}
                                            _ = v25
                                        } else {
                                            v15 = v15
                                            v1 = 0.1
                                            v26 := GenS3{F0: 128, F1: (-(v4)), F2: (v19 + v19)}
                                            _ = v26
                                        }
                                    case "":
                                        switch (v5 | 32768) {
                                        case -2147483648:
                                            switch 127 {
                                            case 361:
                                                fmt.Println(fi18)
                                            default:
                                                v3 += 1e10
                                                v27 := GenS3{F0: 128, F1: 2147483647, F2: v19}
                                                v19 = "hello"
                                                v28 := GenS3{F0: genFn0(), F1: (v4 | v4), F2: v19}
                                                _ = v27
                                                _ = v28
                                            }
                                        case 2147483647:
                                            var v29 int64 = -1
                                            _ = v29
                                        }
                                        if (false && (int32(v15) >= v4)) {
                                            v19 = ("foo bar" + "")
                                            var v30 string = v19
                                            _ = v30
                                        } else {
                                            if true {
                                                fmt.Println(fi18, v15, v5)
                                            }
                                        }
                                        v5 = int(v4)
                                        v19 = "hello"
                                        var v31 uint16 = 65535
                                        v32 := GenS4{F0: "foo bar"}
                                        if false {
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v33 uint16 = v15
                                                _ = v33
                                            case "hello":
                                                v34 := []int{}
                                                _ = v34
                                            case "foo bar":
                                                var v35 int16 = 0
                                                _ = v35
                                            default:
                                                v36 := GenS4{F0: "\n\t"}
                                                _ = v36
                                            }
                                            if false {
                                                var v37 uint16 = v31
                                                v15 -= (^(v37))
                                                if true {
                                                    var v38 uint8 = 2
                                                    _ = v38
                                                } else {
                                                    v39 := []int{(v5 * 9223372036854775807), fi18}
                                                    var v40 int16 = 1000
                                                    _ = v39
                                                    _ = v40
                                                }
                                                _ = v37
                                            } else {
                                                v3 -= 1.5
                                                if false {
                                                    v3 += 0.0
                                                    var v41 uint8 = 2
                                                    fmt.Println(v1, v4, fi23)
                                                    var v42 int32 = v4
                                                    fmt.Println(v4, v1, v5)
                                                    v1 = v1
                                                    v43 := GenS5{F0: -9223372036854775808, F1: false, F2: (v15 | v15)}
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                }
                                                v44 := []int{fi18}
                                                if len(v44) >= 2 {
                                                    v44 = v44[1:len(v44)]
                                                }
                                                _ = v44
                                            }
                                            v45 := GenS1{F0: 413, F1: 42}
                                            _ = v45
                                        }
                                        _ = v31
                                        _ = v32
                                    }
                                case "\n\t":
                                    if true {
                                        if true {
                                            if ((^(v15)) < v15) {
                                                v46 := []int{fi23, v5}
                                                v46 = append(v46, (fi23 | 0))
                                                _ = v46
                                            }
                                            var v47 int64 = 0
                                            v48 := GenS4{F0: ((v19 + (v19 + "foo bar")) + (("" + "hello") + (((v19 + v19) + "\n\t") + "foo bar")))}
                                            var v49 uint16 = (v15 * v15)
                                            var v50 int32 = (v4 - v4)
                                            var v51 bool = true
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                            _ = v50
                                            _ = v51
                                        }
                                        var v52 string = (v19 + v19)
                                        if true {
                                            v1 -= (v1 + (1.5 - v1))
                                            v2 -= 871
                                            var v53 int = 4294967295
                                            _ = v53
                                        }
                                        _ = v52
                                    }
                                default:
                                    var v54 int = v5
                                    _ = v54
                                }
                                var v55 bool = false
                                v56 := GenS6{F0: false}
                                _ = v55
                                _ = v56
                            } else {
                                v4 += (v4 | (v4 >> (((^(v2))) & 63)))
                                v19 = v19
                            }
                            _ = fi23
                        }
                        v57 := []int{int(v2), v5, (v5 / (((fi18 << ((v2) & 63))) | 1))}
                        _ = v57
                    } else {
                        for fi58 := 0; fi58 < 17; fi58++ {
                            v4 += v4
                            var v59 string = v19
                            _ = fi58
                            _ = v59
                        }
                        for fi60 := 0; fi60 < 12; fi60++ {
                            var v61 int16 = -445
                            _ = fi60
                            _ = v61
                        }
                    }
                default:
                    var v62 uint32 = 2147483647
                    _ = v62
                }
                v63 := GenS1{F0: 39, F1: v5}
                _ = fi18
                _ = v19
                _ = v63
            }
        } else {
            v64 := []int{(v5 | (v5 ^ (v5 / (((v5 << ((v2) & 63))) | 1)))), -58}
            if true {
                v65 := GenS7{F0: 42, F1: 2}
                v66 := GenS2{F0: (v5 >> (((v2 + 65535)) & 63)), F1: v15}
                _ = v65
                _ = v66
            } else {
                var v67 float64 = v3
                _ = v67
            }
            fmt.Println(v15)
            _ = v64
        }
    case "a":
        fmt.Println(v5, v5, v15)
    case "":
        v68 := GenS8{F0: (v5 % ((0) | 1)), F1: 0}
        _ = v68
    default:
        for fi69 := 0; fi69 < 19; fi69++ {
            v4 = 1
            _ = fi69
        }
    }
    if true {
        fmt.Println(v15, v5)
        for fi70 := 0; fi70 < 13; fi70++ {
            v1 -= float32(v3)
            v3 = (v3 - 1.5)
            _ = fi70
        }
        if true {
            v71 := []int{v5}
            _ = v71
        } else {
            if false {
                var v72 int = 32767
                _ = v72
            } else {
                switch v5 {
                case -728:
                    for fi73 := 0; fi73 < 19; fi73++ {
                        if false {
                            v74 := GenS8{F0: (fi73 ^ (v5 - fi73)), F1: 2147483648}
                            _ = v74
                        } else {
                            v3 -= v3
                            v75 := []int{-821, (v5 << ((v2) & 63))}
                            var v76 uint8 = 255
                            _ = v75
                            _ = v76
                        }
                        _ = fi73
                    }
                default:
                    for fi77 := 0; fi77 < 10; fi77++ {
                        if true {
                            if ((0.0 < (v3 / (((3.14159 - (v3 / 1.0)) + float64(v1)) - v3))) && (float32(v3) < float32((v3 + 1.5)))) {
                                switch v5 {
                                case 171:
                                    switch (v5 % ((v5) | 1)) {
                                    case 0:
                                        switch "\n\t" {
                                        case "":
                                            v78 := GenS6{F0: false}
                                            _ = v78
                                        case "\n\t":
                                            fmt.Println(v3, v15)
                                        default:
                                            if false {
                                                switch v5 {
                                                case 49:
                                                    var v79 string = ("foo bar" + "")
                                                    _ = v79
                                                case 968:
                                                    for fi80 := 0; fi80 < 19; fi80++ {
                                                        if false {
                                                            switch ((("" + "hello") + "a") + "hello") {
                                                            case "":
                                                                fmt.Println(v1)
                                                            case "\n\t":
                                                                v4 = (v4 * v4)
                                                            case "foo bar":
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    v81 := []int{(fi77 % (((fi77 / ((-937) | 1))) | 1)), v5, (fi77 >> (((v2 - 18446744073709551615)) & 63))}
                                                                    _ = v81
                                                                default:
                                                                    var v82 float32 = (v1 * v1)
                                                                    _ = v82
                                                                }
                                                                if (2 != (v2 * (^(v2)))) {
                                                                    if (!false) {
                                                                        if (!(-478 >= int16(fi77))) {
                                                                            v83 := GenS6{F0: false}
                                                                            if ((!(v5 < (fi77 / ((127) | 1)))) || (!(!(!false)))) {
                                                                                var v84 int32 = int32(v15)
                                                                                var v85 int16 = 1000
                                                                                switch v5 {
                                                                                case 973:
                                                                                    v3 = 1e10
                                                                                case 9223372036854775807:
                                                                                    v86 := GenS6{F0: true}
                                                                                    _ = v86
                                                                                case 32767:
                                                                                    v87 := []int{(fi80 ^ (^(fi77))), 128, (v5 + fi80)}
                                                                                    if len(v87) > 0 {
                                                                                        v87[0] = -230
                                                                                    }
                                                                                    _ = v87
                                                                                default:
                                                                                    if true {
                                                                                        v88 := GenS9{F0: (fi77 + v5)}
                                                                                        _ = v88
                                                                                    }
                                                                                }
                                                                                _ = v84
                                                                                _ = v85
                                                                            } else {
                                                                                v4 = 2147483647
                                                                                switch (v5 ^ (fi80 & fi80)) {
                                                                                case -1:
                                                                                    v4 -= 2147483647
                                                                                }
                                                                            }
                                                                            var v89 int16 = 42
                                                                            _ = v83
                                                                            _ = v89
                                                                        }
                                                                        var v90 int8 = 41
                                                                        v91 := GenS3{F0: 1, F1: (v4 ^ (v4 / ((0) | 1))), F2: "\n\t"}
                                                                        v2 += (v2 | v2)
                                                                        v2 += v2
                                                                        _ = v90
                                                                        _ = v91
                                                                    }
                                                                    if (!(!(0 <= 32767))) {
                                                                        v92 := []int{-2147483648, fi80}
                                                                        switch v5 {
                                                                        case 127:
                                                                            var v93 float64 = 0.0
                                                                            _ = v93
                                                                        case 9223372036854775807:
                                                                            if len(v92) >= 2 {
                                                                                v92 = v92[1:len(v92)]
                                                                            }
                                                                        case -489:
                                                                            _ = len(v92)
                                                                        }
                                                                        var v94 float64 = v3
                                                                        var v95 int16 = 0
                                                                        var v96 int8 = 1
                                                                        if len(v92) >= 2 {
                                                                            v92 = v92[1:len(v92)]
                                                                        }
                                                                        _ = v92
                                                                        _ = v94
                                                                        _ = v95
                                                                        _ = v96
                                                                    } else {
                                                                        switch fi77 {
                                                                        case -517:
                                                                            v3 = v3
                                                                        case 127:
                                                                            var v97 uint = 65535
                                                                            _ = v97
                                                                        default:
                                                                            v15 = v15
                                                                        }
                                                                        if true {
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                v98 := []int{fi80, (fi80 % ((fi77) | 1))}
                                                                                _ = v98
                                                                            case "a":
                                                                                v4 = v4
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v3, v4)
                                                                            if (2 < 1) {
                                                                                var v99 uint64 = 18446744073709551615
                                                                                v100 := []int{fi77, fi77, fi77}
                                                                                _ = v100
                                                                                _ = v99
                                                                            } else {
                                                                                switch (("hello" + ("hello" + ("hello" + "foo bar"))) + "") {
                                                                                case "a":
                                                                                    switch "foo bar" {
                                                                                    case "hello":
                                                                                        var v101 int8 = -128
                                                                                        if false {
                                                                                            v102 := []int{fi77, (fi80 >> (((v2 % ((v2) | 1))) & 63)), (fi80 % ((102) | 1))}
                                                                                            var v103 uint64 = 0
                                                                                            if (genFn1() > 0) {
                                                                                                fmt.Println(v103, v4, fi80)
                                                                                                switch ("foo bar" + ("foo bar" + (("hello" + "foo bar") + ""))) {
                                                                                                case "foo bar":
                                                                                                    var v104 uint32 = 2147483648
                                                                                                    _ = v104
                                                                                                }
                                                                                                v15 = (v15 / ((v15) | 1))
                                                                                                if (81 > genFn0()) {
                                                                                                    if false {
                                                                                                        v102 = append(v102, int(v101))
                                                                                                    }
                                                                                                    fmt.Println(v15)
                                                                                                }
                                                                                                var v105 uint16 = v15
                                                                                                _ = v105
                                                                                            } else {
                                                                                                var v106 uint8 = 54
                                                                                                var v107 uint64 = 4294967295
                                                                                                if (v3 < (-2.25 + float64((0.0 + v1)))) {
                                                                                                    v108 := GenS10{F0: genFn0()}
                                                                                                    _ = v108
                                                                                                }
                                                                                                v4 += int32(v106)
                                                                                                _ = v106
                                                                                                _ = v107
                                                                                            }
                                                                                            _ = v102
                                                                                            _ = v103
                                                                                        } else {
                                                                                            var v109 int8 = (^(v101))
                                                                                            _ = v109
                                                                                        }
                                                                                        var v110 int = (fi80 | fi80)
                                                                                        _ = v101
                                                                                        _ = v110
                                                                                    case "\n\t":
                                                                                        v111 := []int{(v5 & (fi77 - (fi77 | fi80))), v5}
                                                                                        _ = v111
                                                                                    default:
                                                                                        var v112 uint8 = genFn0()
                                                                                        _ = v112
                                                                                    }
                                                                                case "":
                                                                                    fmt.Println(v3)
                                                                                case "hello":
                                                                                    switch "hello" {
                                                                                    case "":
                                                                                        switch ("a" + ("" + "a")) {
                                                                                        case "foo bar":
                                                                                            v15 = (v15 & v15)
                                                                                        default:
                                                                                            v5 += (fi77 | (fi80 - (fi77 + (fi77 / ((v5) | 1)))))
                                                                                        }
                                                                                    default:
                                                                                        v4 = 219
                                                                                    }
                                                                                }
                                                                                switch "a" {
                                                                                case "a":
                                                                                    if (("foo bar" + ("foo bar" + ("a" + "\n\t"))) >= "") {
                                                                                        v113 := GenS11{F0: 0, F1: v2, F2: 294}
                                                                                        _ = v113
                                                                                    } else {
                                                                                        v114 := GenS6{F0: (!true)}
                                                                                        _ = v114
                                                                                    }
                                                                                case "hello":
                                                                                    var v115 float32 = (v1 + ((v1 - v1) / 1.5))
                                                                                    _ = v115
                                                                                }
                                                                                v2 = v2
                                                                            }
                                                                            fmt.Println(v1, v3, v2)
                                                                        }
                                                                        if false {
                                                                            v15 = uint16(v2)
                                                                            var v116 float32 = float32(v3)
                                                                            _ = v116
                                                                        }
                                                                    }
                                                                    var v117 int64 = 1
                                                                    _ = v117
                                                                }
                                                                v118 := GenS1{F0: 42, F1: (v5 & (v5 | fi80))}
                                                                _ = v118
                                                            }
                                                            var v119 int = (v5 % (((^(fi77))) | 1))
                                                            _ = v119
                                                        } else {
                                                            var v120 float32 = (float32(((1.5 - -2.25) * (v3 - 0.0))) - v1)
                                                            fmt.Println(v15)
                                                            v121 := GenS2{F0: -38, F1: (v15 >> ((v2) & 63))}
                                                            _ = v120
                                                            _ = v121
                                                        }
                                                        _ = fi80
                                                    }
                                                case -799:
                                                    var v122 float32 = (float32(v3) / float32(1.5))
                                                    _ = v122
                                                }
                                                v123 := []int{(fi77 ^ fi77), fi77, int(v4)}
                                                v123 = append(v123, v5)
                                                _ = v123
                                            } else {
                                                if true {
                                                    if true {
                                                        v1 += (v1 + (0.0 / v1))
                                                        if false {
                                                            v124 := []int{fi77}
                                                            v4 = (^(v4))
                                                            v125 := GenS8{F0: (fi77 | 127), F1: 635}
                                                            _ = v124
                                                            _ = v125
                                                        } else {
                                                            v126 := []int{v5, v5, (^(fi77))}
                                                            v15 = v15
                                                            var v127 uint = (^(v2))
                                                            _ = v126
                                                            _ = v127
                                                        }
                                                    } else {
                                                        if (42 <= 2147483647) {
                                                            v3 += v3
                                                            if (-13 >= 100) {
                                                                v1 += (0.1 * float32(((0.0 / ((0.0 - 0.0) / v3)) / v3)))
                                                                for fi128 := 0; fi128 < 18; fi128++ {
                                                                    if ((true && ((v4 - (v4 / ((934) | 1))) > 42)) && (1e10 >= (v3 + float64((0.0 + ((v1 * ((v1 / ((3.14159 + float32(v3)) * v1)) + -2.25)) - ((1.5 + v1) * ((3.14159 + ((v1 * ((v1 + (float32(v3) + (v1 - v1))) * (1.5 * v1))) + (v1 - v1))) + v1)))))))) {
                                                                        var v129 float64 = float64(v1)
                                                                        if ("" < "hello") {
                                                                            v1 += ((v1 + (v1 / v1)) * v1)
                                                                            var v130 int16 = 1
                                                                            _ = v130
                                                                        } else {
                                                                            var v131 uint = 255
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                v15 = (v15 + 920)
                                                                                v132 := []int{(-(fi77)), (-(fi128)), -9223372036854775808}
                                                                                v133 := GenS5{F0: 42, F1: ("hello" > (("\n\t" + (("" + (("" + "foo bar") + "foo bar")) + "foo bar")) + "")), F2: v15}
                                                                                _ = v132
                                                                                _ = v133
                                                                            case "foo bar":
                                                                                fmt.Println(v5, v131)
                                                                            case "\n\t":
                                                                                v134 := GenS6{F0: false}
                                                                                _ = v134
                                                                            default:
                                                                                switch -419 {
                                                                                case 140:
                                                                                    if (!(!false)) {
                                                                                        var v135 uint8 = genFn0()
                                                                                        v4 = (-(v4))
                                                                                        v136 := []int{fi128, 678, fi128, fi128}
                                                                                        _ = v135
                                                                                        _ = v136
                                                                                    } else {
                                                                                        v15 += (^(v15))
                                                                                        v137 := GenS6{F0: true}
                                                                                        var v138 uint16 = v15
                                                                                        _ = v137
                                                                                        _ = v138
                                                                                    }
                                                                                case 9223372036854775807:
                                                                                    if (4294967295 == 4294967295) {
                                                                                        v139 := GenS6{F0: false}
                                                                                        v140 := GenS3{F0: genFn0(), F1: (v4 >> ((v2) & 63)), F2: ""}
                                                                                        v141 := GenS2{F0: (v5 & (fi77 % ((v5) | 1))), F1: (v15 << ((v131) & 63))}
                                                                                        _ = v139
                                                                                        _ = v140
                                                                                        _ = v141
                                                                                    }
                                                                                case 347:
                                                                                    v129 = (v129 * (v3 / v3))
                                                                                    var v142 int = 1
                                                                                    var v143 float64 = float64(((1.5 + float32(1.5)) * 1.5))
                                                                                    v144 := []int{-906, (fi77 ^ (fi77 + int(v15))), (v5 ^ fi128)}
                                                                                    _ = v142
                                                                                    _ = v143
                                                                                    _ = v144
                                                                                default:
                                                                                    var v145 uint = 316
                                                                                    _ = v145
                                                                                }
                                                                            }
                                                                            var v146 int32 = (v4 | v4)
                                                                            _ = v131
                                                                            _ = v146
                                                                        }
                                                                        _ = v129
                                                                    } else {
                                                                        var v147 uint8 = 55
                                                                        v148 := GenS12{F0: (((v3 - ((3.14159 / 1.0) / -2.25)) <= ((float64(-2.25) / (float64((0.1 + 1e10)) - float64(float32((3.14159 * ((v3 / (v3 + (v3 / v3))) / 1e10)))))) / float64(3.14159))) && (false && false))}
                                                                        _ = v147
                                                                        _ = v148
                                                                    }
                                                                    fmt.Println(v1, v4, v5)
                                                                    _ = fi128
                                                                }
                                                                v4 -= (-(v4))
                                                            }
                                                        } else {
                                                            v4 = (v4 << (((v2 - (v2 >> ((65535) & 63)))) & 63))
                                                            if (!(1 > 100)) {
                                                                var v149 uint = v2
                                                                _ = v149
                                                            }
                                                        }
                                                        v1 = v1
                                                    }
                                                } else {
                                                    var v150 uint16 = 32767
                                                    if true {
                                                        v151 := []int{v5, fi77, v5}
                                                        _ = len(v151)
                                                        var v152 int = (v5 * v5)
                                                        _ = v151
                                                        _ = v152
                                                    } else {
                                                        v153 := []int{v5, (v5 * -9223372036854775808), v5}
                                                        var v154 int = v5
                                                        for fi155 := 0; fi155 < 18; fi155++ {
                                                            if true {
                                                                v156 := GenS12{F0: false}
                                                                switch "" {
                                                                case "":
                                                                    switch fi77 {
                                                                    case 4294967295:
                                                                        if false {
                                                                            v154 -= v5
                                                                            fmt.Println(v15, fi155)
                                                                            v4 += 274
                                                                            switch 0 {
                                                                            case 0:
                                                                                if (0 != 42) {
                                                                                    v4 -= (v4 + v4)
                                                                                    v154 -= fi77
                                                                                    v1 -= ((-2.25 / v1) + v1)
                                                                                } else {
                                                                                    if len(v153) > 0 {
                                                                                        v153[0] = v5
                                                                                    }
                                                                                    v157 := GenS2{F0: 760, F1: v150}
                                                                                    switch (fi155 >> (((v2 ^ uint(v15))) & 63)) {
                                                                                    case 0:
                                                                                        _ = len(v153)
                                                                                        var v158 uint64 = 42
                                                                                        _ = v158
                                                                                    default:
                                                                                        var v159 int64 = 1
                                                                                        _ = v159
                                                                                    }
                                                                                    _ = v157
                                                                                }
                                                                                switch (fi155 / (((v5 % ((203) | 1))) | 1)) {
                                                                                case 0:
                                                                                    var v160 string = ("hello" + "hello")
                                                                                    _ = v160
                                                                                case 349:
                                                                                    v3 += v3
                                                                                default:
                                                                                    fmt.Println(fi155, v4)
                                                                                }
                                                                            case 9223372036854775807:
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    switch "foo bar" {
                                                                                    case "\n\t":
                                                                                        var v161 uint16 = v150
                                                                                        _ = v161
                                                                                    case "":
                                                                                        fmt.Println(v1, fi77)
                                                                                    case "hello":
                                                                                        _ = len(v153)
                                                                                    default:
                                                                                        if (true && (32767 == v15)) {
                                                                                            if len(v153) >= 2 {
                                                                                                v153 = v153[1:len(v153)]
                                                                                            }
                                                                                            var v162 uint16 = v15
                                                                                            _ = v162
                                                                                        } else {
                                                                                            var v163 int = -91
                                                                                            fmt.Println(v15)
                                                                                            if len(v153) >= 2 {
                                                                                                v153 = v153[1:len(v153)]
                                                                                            }
                                                                                            v2 = 860
                                                                                            v3 = v3
                                                                                            switch "hello" {
                                                                                            case "\n\t":
                                                                                                v154 += fi155
                                                                                            case "":
                                                                                                if true {
                                                                                                    if true {
                                                                                                        if false {
                                                                                                            var v164 int16 = 209
                                                                                                            v1 = (v1 / v1)
                                                                                                            if (0 >= 42) {
                                                                                                                var v165 uint32 = 2147483647
                                                                                                                v153 = append(v153, (fi155 >> ((v2) & 63)))
                                                                                                                if len(v153) >= 2 {
                                                                                                                    v153 = v153[1:len(v153)]
                                                                                                                }
                                                                                                                _ = v165
                                                                                                            }
                                                                                                            _ = v164
                                                                                                        }
                                                                                                        v150 -= (^(v150))
                                                                                                    } else {
                                                                                                        var v166 uint16 = 32767
                                                                                                        if (true && true) {
                                                                                                            fmt.Println(v2, v154)
                                                                                                        }
                                                                                                        switch (fi77 << (((^(v2))) & 63)) {
                                                                                                        case 427:
                                                                                                            v163 += (v163 | 1)
                                                                                                        default:
                                                                                                            switch fi155 {
                                                                                                            case -437:
                                                                                                                _ = len(v153)
                                                                                                            case 32767:
                                                                                                                v153 = append(v153, fi155)
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v166
                                                                                                    }
                                                                                                    fmt.Println(v15)
                                                                                                    v154 += (fi155 + -2147483648)
                                                                                                } else {
                                                                                                    if false {
                                                                                                        var v167 float64 = v3
                                                                                                        _ = v167
                                                                                                    } else {
                                                                                                        switch -2147483648 {
                                                                                                        case 128:
                                                                                                            if false {
                                                                                                                var v168 uint8 = 127
                                                                                                                if len(v153) > 0 {
                                                                                                                    v153[0] = (^(fi77))
                                                                                                                }
                                                                                                                _ = v168
                                                                                                            }
                                                                                                        case 127:
                                                                                                            var v169 float64 = v3
                                                                                                            _ = v169
                                                                                                        default:
                                                                                                            v15 += 1
                                                                                                        }
                                                                                                        var v170 string = ("" + "\n\t")
                                                                                                        _ = v170
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                var v171 int = v5
                                                                                                _ = v171
                                                                                            }
                                                                                            _ = len(v153)
                                                                                            _ = v163
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case 1:
                                                                                if len(v153) > 0 {
                                                                                    v153[0] = fi77
                                                                                }
                                                                            default:
                                                                                v1 = 0.0
                                                                            }
                                                                        }
                                                                    case 128:
                                                                        _ = len(v153)
                                                                    case -279:
                                                                        var v172 string = (((("\n\t" + "foo bar") + "hello") + "foo bar") + "a")
                                                                        _ = v172
                                                                    }
                                                                }
                                                                var v173 uint32 = 162
                                                                fmt.Println(v150, v154, fi77)
                                                                var v174 uint = (v2 - v2)
                                                                _ = v156
                                                                _ = v173
                                                                _ = v174
                                                            }
                                                            _ = fi155
                                                        }
                                                        _ = v153
                                                        _ = v154
                                                    }
                                                    var v175 string = "\n\t"
                                                    _ = v150
                                                    _ = v175
                                                }
                                            }
                                        }
                                    default:
                                        switch ("foo bar" + "\n\t") {
                                        case "a":
                                            v4 = (v4 / (((v4 / ((v4) | 1))) | 1))
                                        case "":
                                            if (0.0 != (v1 / (1.5 * ((0.1 / float32(v3)) - -2.25)))) {
                                                for fi176 := 0; fi176 < 17; fi176++ {
                                                    fmt.Println(v15, fi77, fi176)
                                                    if (!false) {
                                                        v177 := []int{(fi176 << ((42) & 63)), int(v15)}
                                                        if (((fi176 / ((fi77) | 1)) > fi77) || (-128 < -1)) {
                                                            if (18446744073709551615 <= 0) {
                                                                var v178 string = ("hello" + ("\n\t" + (("\n\t" + "foo bar") + "hello")))
                                                                fmt.Println(v1)
                                                                _ = v178
                                                            }
                                                            v2 += (v2 + v2)
                                                            var v179 int16 = 140
                                                            var v180 int32 = v4
                                                            var v181 uint = (v2 << ((437) & 63))
                                                            _ = v179
                                                            _ = v180
                                                            _ = v181
                                                        }
                                                        _ = v177
                                                    } else {
                                                        v182 := GenS8{F0: 2147483647, F1: 0}
                                                        fmt.Println(v2, v5)
                                                        _ = v182
                                                    }
                                                    var v183 string = "\n\t"
                                                    v184 := GenS13{F0: "hello", F1: 1}
                                                    _ = fi176
                                                    _ = v183
                                                    _ = v184
                                                }
                                            } else {
                                                var v185 int = fi77
                                                var v186 int = (v5 + v185)
                                                _ = v185
                                                _ = v186
                                            }
                                        }
                                    }
                                default:
                                    var v187 int8 = -1
                                    _ = v187
                                }
                                var v188 int64 = -939
                                v15 = v15
                                _ = v188
                            } else {
                                for fi189 := 0; fi189 < 18; fi189++ {
                                    var v190 float32 = 1.5
                                    switch ("hello" + ("a" + "hello")) {
                                    case "":
                                        v3 -= 3.14159
                                    case "hello":
                                        v2 = v2
                                    case "foo bar":
                                        fmt.Println(fi189)
                                        if (false && ((2147483648 < 42) && false)) {
                                            fmt.Println(v5)
                                        }
                                    }
                                    v15 = (v15 << (((v2 ^ v2)) & 63))
                                    _ = fi189
                                    _ = v190
                                }
                            }
                            fmt.Println(v1)
                        } else {
                            v1 -= ((float32(0.1) * float32(v3)) / v1)
                            v1 = v1
                            if ((3.14159 * v1) > float32((v3 - 1e10))) {
                                v2 -= 2
                            }
                        }
                        var v191 float32 = -2.25
                        v5 = -1
                        _ = fi77
                        _ = v191
                    }
                }
            }
            if false {
                v192 := []int{}
                fmt.Println(v15, v2, v2)
                _ = v192
            }
        }
    }
    var v193 int16 = 0
    if true {
        if true {
            fmt.Println(v5, v15)
            if (9223372036854775807 > -9223372036854775808) {
                if (v1 == (0.1 * v1)) {
                    var v194 int64 = 42
                    var v195 float64 = (v3 * (((0.0 + (v3 - v3)) / v3) + float64(((v1 * float32(((v3 + float64(((float32(v3) + (((3.14159 / ((v1 / -2.25) * float32((v3 * (v3 + v3))))) * v1) + v1)) + (v1 - 3.14159)))) - v3))) * v1))))
                    _ = v194
                    _ = v195
                }
                var v196 int16 = (v193 ^ v193)
                for fi197 := 0; fi197 < 1; fi197++ {
                    for fi198 := 0; fi198 < 12; fi198++ {
                        if true {
                            if false {
                                if (!((-(v193)) < v193)) {
                                    v199 := GenS14{F0: 65535, F1: v15, F2: 864}
                                    var v200 int64 = -1
                                    var v201 int8 = 42
                                    _ = v199
                                    _ = v200
                                    _ = v201
                                }
                                v193 += -1
                            }
                        }
                        v202 := GenS8{F0: fi198, F1: 2}
                        _ = fi198
                        _ = v202
                    }
                    var v203 int8 = 127
                    _ = fi197
                    _ = v203
                }
                _ = v196
            }
            switch "\n\t" {
            case "hello":
                fmt.Println(v1)
            case "foo bar":
                switch -1 {
                case 9223372036854775807:
                    v15 = 399
                }
            case "\n\t":
                var v204 uint16 = v15
                _ = v204
            }
        } else {
            var v205 int16 = (v193 ^ v193)
            _ = v205
        }
    }
    v206 := GenS8{F0: v5, F1: 4294967295}
    var v207 uint = (^(v2))
    var v208 uint64 = 1
    v209 := GenS9{F0: (v5 - (v5 << (((v207 + (v207 * 952))) & 63)))}
    var v210 uint16 = (v15 >> ((570) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v193:", v193)
    fmt.Println("v2:", v2)
    fmt.Println("v207:", v207)
    fmt.Println("v208:", v208)
    fmt.Println("v210:", v210)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v193
    _ = v2
    _ = v206
    _ = v207
    _ = v208
    _ = v209
    _ = v210
    _ = v3
    _ = v4
    _ = v5
}
