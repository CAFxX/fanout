package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int; F1 uint64 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int32; F1 int32; F2 uint16 }
type GenS13 struct { F0 int; F1 bool }
type GenS14 struct { F0 string; F1 uint16; F2 int8 }
type GenS15 struct { F0 bool; F1 int32; F2 string }
func genFn0(p0 uint32, p1 uint64) int32 {
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 int16 = 32767
        _ = fi5
        _ = v6
    }
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1(p0 int8) uint64 {
    var v8 uint8 = 1
    p0 -= p0
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            switch fi10 {
            case 4294967295:
                fmt.Println(v8)
            case -2147483648:
                var v11 int16 = 32767
                v12 := GenS1{F0: "hello", F1: v8}
                _ = v11
                _ = v12
            default:
                var v13 uint32 = 2
                _ = v13
            }
            p0 += (p0 << ((42) & 63))
            switch ("" + "hello") {
            case "foo bar":
                if true {
                    v8 -= 42
                    var v14 bool = (float64(-2.25) >= 0.1)
                    _ = v14
                }
            default:
                v15 := GenS1{F0: ("" + ""), F1: (v8 + (v8 + v8))}
                _ = v15
            }
            _ = fi10
        }
        _ = fi9
    }
    p0 = p0
    _ = p0
    _ = v8
    return 68
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    var v4 uint = 65535
    var v16 int32 = 42
    var v17 uint8 = 1
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 float64 = v1
        for fi20 := 0; fi20 < 2; fi20++ {
            switch fi18 {
            case -2147483648:
                v4 -= 65535
                fmt.Println(v19, v19)
                switch "" {
                case "":
                    var v21 uint32 = 275
                    _ = v21
                case "foo bar":
                    switch ("" + "a") {
                    case "hello":
                        switch "a" {
                        case "\n\t":
                            if (true || ((fi18 + fi18) >= fi20)) {
                                switch "" {
                                case "foo bar":
                                    if true {
                                        fmt.Println(v4, fi18)
                                    } else {
                                        if (!(-32768 < 42)) {
                                            v19 += float64(((float32((3.14159 + (0.0 + v1))) - 1.5) - (float32(v1) - 1e10)))
                                            switch "a" {
                                            case "\n\t":
                                                v1 += v19
                                            case "a":
                                                fmt.Println(v4, v4, v1)
                                            default:
                                                if false {
                                                    fmt.Println(fi18, v16, fi20)
                                                }
                                            }
                                            if true {
                                                v17 = v17
                                                v17 -= (v17 - (v17 - (v17 >> ((v4) & 63))))
                                                v19 = 3.14159
                                                switch fi18 {
                                                case 4294967295:
                                                    switch (fi20 + (^(fi18))) {
                                                    case 128:
                                                        v22 := []int{(^(fi20)), (fi20 >> ((712) & 63)), (-(fi20)), fi20}
                                                        var v23 int8 = -128
                                                        if (true && ((v16 | genFn0(0, uint64(v17))) > -653)) {
                                                            fmt.Println(v17, fi20, v23)
                                                            if true {
                                                                v24 := GenS1{F0: "a", F1: 121}
                                                                if true {
                                                                    v4 -= 872
                                                                    fmt.Println(v19)
                                                                    if false {
                                                                        v25 := GenS2{F0: v23}
                                                                        v26 := GenS1{F0: "", F1: (v17 & v17)}
                                                                        var v27 uint16 = 2
                                                                        _ = v25
                                                                        _ = v26
                                                                        _ = v27
                                                                    }
                                                                }
                                                                _ = v24
                                                            } else {
                                                                var v28 uint8 = (v17 % ((v17) | 1))
                                                                v29 := GenS1{F0: (("a" + "a") + "\n\t"), F1: (v17 / (((v28 + (v28 & (v17 / ((v28) | 1))))) | 1))}
                                                                if true {
                                                                    var v30 uint64 = 42
                                                                    fmt.Println(v1, v16)
                                                                    _ = v30
                                                                }
                                                                if len(v22) > 0 {
                                                                    v22[0] = 1
                                                                }
                                                                v23 -= v23
                                                                var v31 int64 = -189
                                                                v32 := GenS2{F0: v23}
                                                                switch (-(fi18)) {
                                                                case -2147483648:
                                                                    if (!(false && (true || true))) {
                                                                        var v33 int64 = (v31 & v31)
                                                                        var v34 int16 = 0
                                                                        _ = v33
                                                                        _ = v34
                                                                    } else {
                                                                        v16 = v16
                                                                        var v35 bool = (v28 >= 255)
                                                                        _ = v35
                                                                    }
                                                                case -1:
                                                                    fmt.Println(fi18, v28)
                                                                default:
                                                                    v1 = 3.14159
                                                                }
                                                                v36 := GenS2{F0: v23}
                                                                _ = v28
                                                                _ = v29
                                                                _ = v31
                                                                _ = v32
                                                                _ = v36
                                                            }
                                                            var v37 float64 = v1
                                                            _ = v37
                                                        } else {
                                                            v38 := GenS3{F0: (fi20 << ((0) & 63)), F1: "a"}
                                                            v4 -= v4
                                                            var v39 uint32 = 81
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                        _ = v22
                                                        _ = v23
                                                    default:
                                                        delete(v2, ("hello" + "\n\t"))
                                                    }
                                                case 127:
                                                    v40 := []int{9223372036854775807, (fi18 / ((fi20) | 1)), (fi20 | (fi18 << (((v4 / ((v4) | 1))) & 63)))}
                                                    _ = v40
                                                case -9223372036854775808:
                                                    v16 += v16
                                                default:
                                                    var v41 int16 = 749
                                                    _ = v41
                                                }
                                                v4 += 42
                                            }
                                        }
                                    }
                                    switch fi20 {
                                    case 127:
                                        var v42 bool = true
                                        _ = v42
                                    case 32768:
                                        fmt.Println(fi18)
                                    case 467:
                                        var v43 uint = (v4 >> ((v4) & 63))
                                        _ = v43
                                    default:
                                        v44 := []int{}
                                        _ = v44
                                    }
                                    if (737 < 2147483647) {
                                        if true {
                                            v16 += (^(v16))
                                            var v45 uint64 = 4294967295
                                            _ = v45
                                        }
                                        var v46 int = (-(fi20))
                                        if false {
                                            v1 -= v19
                                            var v47 float32 = (0.0 - (-2.25 * 3.14159))
                                            v48 := v2["hello"]
                                            _ = v47
                                            _ = v48
                                        } else {
                                            v46 = 42
                                            v17 = (v17 ^ v17)
                                            v49 := GenS3{F0: fi20, F1: (("hello" + "a") + "\n\t")}
                                            _ = v49
                                        }
                                        _ = v46
                                    } else {
                                        delete(v2, "foo bar")
                                        switch "a" {
                                        case "hello":
                                            v50 := GenS4{F0: -1}
                                            _ = v50
                                        case "foo bar":
                                            var v51 bool = (0 == 624)
                                            _ = v51
                                        }
                                    }
                                    if true {
                                        var v52 int64 = -1
                                        v1 -= 1e10
                                        if true {
                                            var v53 int32 = -1
                                            var v54 uint = v4
                                            var v55 uint64 = 2
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        } else {
                                            v1 = v1
                                        }
                                        _ = v52
                                    }
                                case "\n\t":
                                    if (-80 > 127) {
                                        if false {
                                            fmt.Println(fi20, v17, v16)
                                            v56 := GenS4{F0: 9223372036854775807}
                                            switch fi20 {
                                            case 4294967295:
                                                switch (^(fi18)) {
                                                case -1:
                                                    var v57 int16 = 42
                                                    _ = v57
                                                }
                                            case 32768:
                                                if (2147483648 > 2147483647) {
                                                    if true {
                                                        switch (fi20 - -355) {
                                                        case 9223372036854775807:
                                                            switch fi18 {
                                                            case 42:
                                                                v58 := GenS3{F0: (fi20 << ((950) & 63)), F1: ""}
                                                                _ = v58
                                                            }
                                                            v59 := []int{fi20, fi18, (fi20 * fi20)}
                                                            switch (("foo bar" + "hello") + ("\n\t" + ((("\n\t" + "a") + "hello") + ""))) {
                                                            case "":
                                                                v17 -= v17
                                                            case "hello":
                                                                var v60 uint16 = 1
                                                                _ = v60
                                                            }
                                                            _ = v59
                                                        case -9223372036854775808:
                                                            switch ("\n\t" + "foo bar") {
                                                            case "a":
                                                                delete(v3, (fi20 ^ -2147483648))
                                                                v16 += (v16 << ((230) & 63))
                                                            case "hello":
                                                                fmt.Println(fi20)
                                                            case "\n\t":
                                                                if false {
                                                                    if (!false) {
                                                                        v19 -= -2.25
                                                                    }
                                                                    delete(v3, -653)
                                                                } else {
                                                                    v19 = 0.0
                                                                    var v61 bool = true
                                                                    var v62 uint64 = 4294967295
                                                                    _ = v61
                                                                    _ = v62
                                                                }
                                                            }
                                                        default:
                                                            fmt.Println(v16, v16)
                                                        }
                                                    }
                                                    v16 -= 0
                                                    v63 := []int{(fi18 + int(v16)), (-(fi20)), int(v17)}
                                                    _ = v63
                                                }
                                            case 9223372036854775807:
                                                v64 := GenS5{F0: 2147483647, F1: 0}
                                                _ = v64
                                            default:
                                                switch -832 {
                                                case 67:
                                                    v65 := []int{(fi20 - fi20), fi20, (fi18 - fi20), fi18}
                                                    _ = v65
                                                case -1:
                                                    v66 := []int{fi20}
                                                    _ = v66
                                                case 127:
                                                    switch fi18 {
                                                    case 2147483647:
                                                        v67 := GenS3{F0: int(v16), F1: ("\n\t" + ("\n\t" + ""))}
                                                        _ = v67
                                                    }
                                                default:
                                                    v19 = 0.0
                                                }
                                                fmt.Println(v16, v17, v17)
                                                var v68 int64 = int64(fi20)
                                                _ = v68
                                            }
                                            fmt.Println(fi18, fi20, v19)
                                            _ = v56
                                        }
                                    }
                                case "":
                                    if (42 < 1) {
                                        v1 += v1
                                        var v69 int8 = -1
                                        v70 := GenS2{F0: int8(fi18)}
                                        _ = v69
                                        _ = v70
                                    } else {
                                        var v71 uint16 = 0
                                        v72 := GenS3{F0: int(v16), F1: "a"}
                                        _ = v71
                                        _ = v72
                                    }
                                default:
                                    switch "hello" {
                                    case "":
                                        switch fi18 {
                                        case 287:
                                            var v73 uint8 = (v17 + (v17 * (v17 | v17)))
                                            _ = v73
                                        case 128:
                                            var v74 uint8 = v17
                                            _ = v74
                                        case 42:
                                            v75 := GenS6{F0: -8, F1: 100}
                                            _ = v75
                                        }
                                    }
                                }
                            } else {
                                fmt.Println(v1)
                            }
                        case "hello":
                            switch (^(fi20)) {
                            case -569:
                                switch ((("a" + "\n\t") + "") + "") {
                                case "foo bar":
                                    var v76 int64 = -283
                                    _ = v76
                                default:
                                    switch "hello" {
                                    case "a":
                                        v77 := v3[(fi18 % (((fi20 + (fi20 / ((-1) | 1)))) | 1))]
                                        _ = v77
                                    case "":
                                        v78 := GenS7{F0: ("" + "foo bar")}
                                        _ = v78
                                    default:
                                        if (0 > 0) {
                                            switch "" {
                                            case "\n\t":
                                                v79 := GenS8{F0: v4}
                                                _ = v79
                                            }
                                            var v80 int64 = -9223372036854775808
                                            v19 = ((0.1 + 0.1) * (((3.14159 - ((-2.25 + float64(0.0)) * 1e10)) * v1) + (((v1 / v1) / v19) + 3.14159)))
                                            var v81 uint64 = 4294967295
                                            v4 += v4
                                            fmt.Println(v1)
                                            _ = v80
                                            _ = v81
                                        } else {
                                            v19 = (((v1 + ((v1 - v1) + ((v19 + ((v1 - v19) + v1)) - v1))) + ((v19 / v19) / v1)) / float64(0.1))
                                            v82 := GenS7{F0: ("" + "hello")}
                                            _ = v82
                                        }
                                    }
                                }
                            }
                        case "foo bar":
                            v1 += (v1 + v1)
                        }
                    case "\n\t":
                        v83 := GenS9{F0: true}
                        _ = v83
                    default:
                        v17 += (v17 % (((v17 % ((v17) | 1))) | 1))
                    }
                }
            case 32768:
                var v84 uint32 = 42
                _ = v84
            }
            v85 := v2[("a" + "\n\t")]
            switch -2147483648 {
            case 4294967295:
                var v86 uint64 = genFn1(0)
                _ = v86
            case 2147483647:
                if true {
                    fmt.Println(fi20, fi18, v1)
                } else {
                    var v87 uint32 = 2147483648
                    v1 += (((v1 - v1) / (v19 + float64(1.5))) * (v19 + ((v1 + (v1 / 1e10)) - 1e10)))
                    if true {
                        v4 += 18446744073709551615
                        v16 += (v16 ^ (-(v16)))
                    } else {
                        switch (((("" + (("hello" + "\n\t") + "\n\t")) + "a") + "foo bar") + ("a" + ("foo bar" + ("hello" + "\n\t")))) {
                        case "a":
                            fmt.Println(fi20, v4, v19)
                        case "foo bar":
                            var v88 int64 = -1
                            _ = v88
                        case "\n\t":
                            var v89 float64 = 3.14159
                            _ = v89
                        default:
                            fmt.Println(v16, fi20, v85)
                        }
                    }
                    _ = v87
                }
            case 32768:
                v1 = (v19 - -2.25)
            default:
                switch fi18 {
                case 0:
                    var v90 uint64 = 18446744073709551615
                    _ = v90
                default:
                    v91 := GenS4{F0: -9223372036854775808}
                    _ = v91
                }
            }
            if true {
                v16 = (v16 + (v16 << ((v4) & 63)))
                if (-176 > 42) {
                    var v92 uint16 = 255
                    if true {
                        var v93 float64 = ((v1 / (v19 - v19)) - v19)
                        _ = v93
                    }
                    _ = v92
                } else {
                    var v94 string = ""
                    var v95 int32 = v16
                    if (!(false || false)) {
                        v96 := GenS10{F0: (fi20 >> ((0) & 63)), F1: uint64(v17)}
                        switch fi18 {
                        case 159:
                            delete(v2, v94)
                        case 32768:
                            fmt.Println(v17)
                        case 9223372036854775807:
                            var v97 int64 = 42
                            _ = v97
                        }
                        _ = v96
                    }
                    _ = v94
                    _ = v95
                }
            }
            _ = fi20
            _ = v85
        }
        if ((-18 < 120) && false) {
            var v98 uint8 = (v17 + v17)
            v17 += v98
            _ = v98
        } else {
            v1 = v1
            var v99 int = (fi18 >> ((uint(v17)) & 63))
            _ = v99
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v16)
    for fi100 := 0; fi100 < 14; fi100++ {
        if (v17 != 42) {
            v101 := GenS2{F0: 127}
            for fi102 := 0; fi102 < 3; fi102++ {
                v103 := GenS2{F0: 107}
                var v104 int32 = v16
                var v105 int32 = (^(v104))
                var v106 float32 = 0.1
                v106 += 0.1
                _ = fi102
                _ = v103
                _ = v104
                _ = v105
                _ = v106
            }
            v17 = (v17 >> ((18446744073709551615) & 63))
            _ = v101
        }
        v17 -= v17
        _ = fi100
    }
    v107 := GenS11{F0: (v4 + (v4 + v4))}
    if false {
        v108 := []int{-463}
        for fi109 := 0; fi109 < 6; fi109++ {
            v110 := GenS11{F0: 1}
            _ = fi109
            _ = v110
        }
        for fi111 := 0; fi111 < 1; fi111++ {
            var v112 string = "foo bar"
            v113 := GenS8{F0: 131}
            v114 := v3[(^(fi111))]
            _ = fi111
            _ = v112
            _ = v113
            _ = v114
        }
        _ = v108
    }
    if (1 > 508) {
        v115 := GenS12{F0: -635, F1: (v16 ^ (v16 / (((^(v16))) | 1))), F2: 42}
        _ = v115
    } else {
        for fi116 := 0; fi116 < 6; fi116++ {
            if false {
                var v117 uint16 = 319
                v118 := GenS7{F0: ""}
                var v119 float64 = (0.1 / v1)
                _ = v117
                _ = v118
                _ = v119
            } else {
                v17 = v17
                if (100 >= 79) {
                    for fi120 := 0; fi120 < 14; fi120++ {
                        v1 = v1
                        if false {
                            v121 := []int{(fi120 ^ -2147483648), 42, int(v4)}
                            switch "foo bar" {
                            case "foo bar":
                                if (false || (!(!true))) {
                                    _ = len(v3)
                                    var v122 float32 = (1.5 + ((1e10 - (0.0 * (-2.25 / 1.0))) - float32(((3.14159 / (((v1 / (v1 - (0.1 / v1))) - 1.5) + v1)) / float64((0.1 - (1.5 - 1.5)))))))
                                    _ = v122
                                }
                            case "a":
                                var v123 int64 = 9223372036854775807
                                _ = v123
                            }
                            _ = v121
                        }
                        fmt.Println(v17, fi120)
                        _ = fi120
                    }
                    v124 := []int{4294967295, (fi116 & (fi116 & fi116)), 127}
                    v125 := GenS6{F0: 83, F1: -80}
                    _ = v124
                    _ = v125
                } else {
                    v2[("\n\t" + ("" + ("" + ("foo bar" + "\n\t"))))] = -1
                }
                if true {
                    var v126 float32 = 3.14159
                    var v127 uint8 = v17
                    var v128 int8 = int8(v4)
                    _ = v126
                    _ = v127
                    _ = v128
                } else {
                    v129 := GenS3{F0: (fi116 - 458), F1: ("hello" + ("\n\t" + ("foo bar" + "foo bar")))}
                    v16 = (v16 << (((v4 + v4)) & 63))
                    var v130 float32 = (3.14159 + 1.5)
                    _ = v129
                    _ = v130
                }
            }
            _ = fi116
        }
        var v131 bool = (float64(-2.25) <= v1)
        _ = v131
    }
    if true {
        var v132 int = 42
        if false {
            switch -9223372036854775808 {
            case 60:
                for fi133 := 0; fi133 < 1; fi133++ {
                    if true {
                        if (-9223372036854775808 <= 698) {
                            var v134 uint = v4
                            v135 := GenS5{F0: v132, F1: 565}
                            var v136 uint8 = v17
                            _ = v134
                            _ = v135
                            _ = v136
                        }
                        var v137 int8 = -1
                        _ = v137
                    } else {
                        v138 := GenS9{F0: false}
                        _ = v138
                    }
                    _ = fi133
                }
            case 182:
                if true {
                    v17 = v17
                } else {
                    for fi139 := 0; fi139 < 5; fi139++ {
                        for fi140 := 0; fi140 < 4; fi140++ {
                            if (100 <= 100) {
                                if true {
                                    var v141 float32 = -2.25
                                    if false {
                                        v142 := GenS4{F0: -9223372036854775808}
                                        _ = v142
                                    } else {
                                        switch "\n\t" {
                                        case "\n\t":
                                            v143 := []int{73, (fi139 | (fi139 - fi139)), fi140}
                                            if true {
                                                var v144 int32 = (v16 * (v16 / ((0) | 1)))
                                                var v145 bool = (v17 <= (v17 * (^(v17))))
                                                fmt.Println(v4, v132, v144)
                                                _ = v144
                                                _ = v145
                                            }
                                            v1 = 1.5
                                            _ = v143
                                        default:
                                            if false {
                                                v146 := GenS2{F0: 1}
                                                if (true && ((v132 >> (((^(v4))) & 63)) < (v132 & 32767))) {
                                                    if true {
                                                        v4 += (v4 % (((v4 << ((396) & 63))) | 1))
                                                        _ = len(v2)
                                                        if true {
                                                            v147 := GenS13{F0: fi139, F1: true}
                                                            if false {
                                                                v148 := GenS13{F0: 9223372036854775807, F1: (v132 != fi140)}
                                                                _ = v148
                                                            }
                                                            v4 -= (v4 >> ((873) & 63))
                                                            _ = v147
                                                        }
                                                    }
                                                    v149 := GenS14{F0: "hello", F1: 1, F2: int8(v17)}
                                                    _ = v149
                                                } else {
                                                    var v150 float32 = (((float32(v1) + v141) - v141) + float32(3.14159))
                                                    v132 = int(v16)
                                                    _ = v150
                                                }
                                                v151 := GenS14{F0: "", F1: 367, F2: 127}
                                                _ = v146
                                                _ = v151
                                            }
                                        }
                                        var v152 uint64 = 355
                                        _ = v152
                                    }
                                    v1 = v1
                                    _ = v141
                                } else {
                                    v153 := GenS5{F0: fi139, F1: 50}
                                    var v154 int64 = -1
                                    _ = v153
                                    _ = v154
                                }
                                switch fi140 {
                                case 4294967295:
                                    v16 += (v16 ^ v16)
                                case 1:
                                    fmt.Println(fi139, v1, v16)
                                case -1:
                                    v4 += v4
                                default:
                                    var v155 string = "foo bar"
                                    _ = v155
                                }
                            }
                            var v156 float32 = 0.1
                            var v157 uint64 = 18446744073709551615
                            _ = fi140
                            _ = v156
                            _ = v157
                        }
                        _ = fi139
                    }
                }
            default:
                if (!true) {
                    v132 += -343
                } else {
                    if false {
                        var v158 uint32 = 1
                        switch "\n\t" {
                        case "\n\t":
                            if (-164 == 0) {
                                for fi159 := 0; fi159 < 10; fi159++ {
                                    for fi160 := 0; fi160 < 2; fi160++ {
                                        v161 := GenS12{F0: (v16 ^ genFn0((^(v158)), 18446744073709551615)), F1: v16, F2: 65535}
                                        _ = fi160
                                        _ = v161
                                    }
                                    v132 = (v132 & (-(v132)))
                                    for fi162 := 0; fi162 < 7; fi162++ {
                                        var v163 int8 = 100
                                        _ = fi162
                                        _ = v163
                                    }
                                    _ = fi159
                                }
                            } else {
                                fmt.Println(v16, v4)
                                v4 = (v4 >> (((v4 + (v4 % (((v4 / ((v4) | 1))) | 1)))) & 63))
                            }
                        case "hello":
                            v164 := []int{v132, (v132 >> ((255) & 63))}
                            _ = v164
                        }
                        _ = v158
                    } else {
                        v132 = v132
                    }
                    v165 := GenS15{F0: (!(v4 > (v4 % (((^(v4))) | 1)))), F1: v16, F2: "\n\t"}
                    _ = v165
                }
            }
        }
        _ = v132
    } else {
        v166 := []int{32768}
        _ = v166
    }
    var v167 string = "\n\t"
    if (525 == 9) {
        if true {
            var v168 int32 = (-(v16))
            v169 := GenS8{F0: (v4 % ((18) | 1))}
            v167 = (v167 + v167)
            _ = v168
            _ = v169
        }
        var v170 uint16 = 255
        _ = v170
    } else {
        var v171 uint64 = genFn1(0)
        v172 := GenS7{F0: (v167 + v167)}
        _ = v171
        _ = v172
    }
    fmt.Println(v4, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v167:", v167)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v107
    _ = v16
    _ = v167
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
