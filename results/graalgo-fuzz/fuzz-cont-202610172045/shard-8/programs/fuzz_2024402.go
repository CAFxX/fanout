package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int64 }
type GenS2 struct { F0 string; F1 uint32; F2 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 uint; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string; F1 uint8 }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint; F1 int16 }
type GenS11 struct { F0 int; F1 int8; F2 uint8 }
type GenS12 struct { F0 uint; F1 int16; F2 uint8 }
func genFn0() int32 {
    v7 := GenS1{F0: 0, F1: 623, F2: -873}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint64 = 999
        var v10 string = ("a" + "\n\t")
        for fi11 := 0; fi11 < 13; fi11++ {
            switch (((("foo bar" + v10) + v10) + "foo bar") + v10) {
            case "hello":
                v10 = (v10 + "")
            case "\n\t":
                if (0.1 <= 3.14159) {
                    if (-379 > -1) {
                        var v12 int16 = -378
                        if (false && (-32768 <= (v12 | (v12 ^ (^(v12)))))) {
                            if ((true || ((int16(v9) >= v12) && (0 < 57))) || (v9 == (v9 - 649))) {
                                v10 = "foo bar"
                                var v13 float64 = ((float64((float32(1.5) * 0.0)) - 1e10) - 1e10)
                                _ = v13
                            } else {
                                switch v10 {
                                case "a":
                                    var v14 string = "\n\t"
                                    _ = v14
                                default:
                                    switch (("\n\t" + (v10 + v10)) + (("a" + v10) + v10)) {
                                    case "a":
                                        if (((float64(3.14159) * 0.1) + (0.1 - (-2.25 + float64(0.1)))) > float64(0.1)) {
                                            switch (^(fi11)) {
                                            case -1:
                                                var v15 int32 = 0
                                                _ = v15
                                            }
                                            var v16 int16 = v12
                                            var v17 int16 = 1
                                            _ = v16
                                            _ = v17
                                        }
                                        v9 = v9
                                        fmt.Println(fi8, v9, v9)
                                    case "foo bar":
                                        v18 := GenS1{F0: 42, F1: v12, F2: 9223372036854775807}
                                        _ = v18
                                    case "hello":
                                        if (42 == -36) {
                                            v10 = (((v10 + ((("a" + "") + "") + "\n\t")) + "\n\t") + ((v10 + (v10 + ("" + (v10 + v10)))) + (v10 + v10)))
                                            v19 := []int{128, (fi8 - fi8), fi8}
                                            var v20 string = (("" + "\n\t") + v10)
                                            _ = v19
                                            _ = v20
                                        } else {
                                            v12 -= 1
                                            switch v10 {
                                            case "":
                                                var v21 bool = (2 < 864)
                                                _ = v21
                                            default:
                                                switch (v10 + v10) {
                                                case "foo bar":
                                                    var v22 uint = 65535
                                                    _ = v22
                                                default:
                                                    fmt.Println(v9, v12, v12)
                                                    var v23 uint8 = 128
                                                    var v24 uint32 = 2147483647
                                                    _ = v23
                                                    _ = v24
                                                }
                                            }
                                            if true {
                                                fmt.Println(v10, fi11, v10)
                                                switch fi11 {
                                                case 32768:
                                                    if (v10 > "\n\t") {
                                                        var v25 int = fi8
                                                        fmt.Println(v10)
                                                        _ = v25
                                                    } else {
                                                        v26 := []int{fi11, fi11, 1}
                                                        var v27 bool = (2 != 1)
                                                        if false {
                                                            if false {
                                                                var v28 uint64 = 42
                                                                _ = v28
                                                            }
                                                            v9 = 0
                                                        } else {
                                                            if false {
                                                                switch (fi11 / ((fi11) | 1)) {
                                                                case 4294967295:
                                                                    switch fi8 {
                                                                    case 100:
                                                                        v9 = v9
                                                                    default:
                                                                        if true {
                                                                            switch v10 {
                                                                            case "a":
                                                                                var v29 int8 = 127
                                                                                _ = v29
                                                                            }
                                                                        }
                                                                    }
                                                                case 2147483647:
                                                                    if len(v26) >= 2 {
                                                                        v26 = v26[1:len(v26)]
                                                                    }
                                                                }
                                                            } else {
                                                                var v30 uint16 = 217
                                                                _ = v30
                                                            }
                                                        }
                                                        _ = v26
                                                        _ = v27
                                                    }
                                                default:
                                                    v31 := GenS2{F0: v10, F1: 2147483648, F2: -90}
                                                    _ = v31
                                                }
                                                fmt.Println(fi11)
                                            }
                                        }
                                    default:
                                        var v32 string = "a"
                                        switch v32 {
                                        case "a":
                                            fmt.Println(v12, fi8, v10)
                                        default:
                                            v33 := GenS2{F0: v32, F1: 2147483647, F2: 127}
                                            _ = v33
                                        }
                                        _ = v32
                                    }
                                }
                            }
                            v34 := []int{fi8}
                            if true {
                                fmt.Println(fi8, v9)
                                var v35 float64 = 1.5
                                v10 = ("hello" + v10)
                                var v36 uint16 = 32767
                                v37 := GenS2{F0: (v10 + v10), F1: 0, F2: -110}
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                            _ = v34
                        }
                        _ = v12
                    }
                    v9 = (v9 + 0)
                    v38 := GenS1{F0: 999, F1: -32768, F2: 1}
                    _ = v38
                } else {
                    switch (-(fi11)) {
                    case -2147483648:
                        var v39 uint64 = v9
                        _ = v39
                    case 128:
                        v40 := []int{fi11}
                        _ = v40
                    default:
                        var v41 int = (fi8 ^ (fi11 << ((1) & 63)))
                        _ = v41
                    }
                    v10 = ("foo bar" + ("\n\t" + v10))
                }
            default:
                if false {
                    v42 := GenS1{F0: 63, F1: 360, F2: 1}
                    _ = v42
                } else {
                    var v43 int64 = -1
                    var v44 bool = ((v10 != "\n\t") || (!(false && ((!((v10 >= ("" + "")) && true)) && (false || (!(!(2 < 42))))))))
                    _ = v43
                    _ = v44
                }
            }
            _ = fi11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi45 := 0; fi45 < 20; fi45++ {
        fmt.Println(fi45)
        var v46 uint64 = 115
        _ = fi45
        _ = v46
    }
    v47 := GenS3{F0: (!true)}
    for fi48 := 0; fi48 < 10; fi48++ {
        for fi49 := 0; fi49 < 13; fi49++ {
            v50 := GenS4{F0: 441, F1: 42, F2: 42}
            _ = fi49
            _ = v50
        }
        var v51 float32 = 0.1
        v51 = float32(((3.14159 * ((-2.25 + ((3.14159 / 1.0) - -2.25)) / 1.0)) / -2.25))
        _ = fi48
        _ = v51
    }
    _ = v47
    _ = v7
    return -960
}

func main() {
    v1 := []int{}
    var v2 string = ""
    v3 := make(map[int]string)
    v4 := make(map[string]bool)
    var v5 uint64 = 589
    var v6 uint64 = (^(v5))
    for fi52 := 0; fi52 < 2; fi52++ {
        v2 = "a"
        _ = fi52
    }
    for fi53 := 0; fi53 < 5; fi53++ {
        if (-32768 > 42) {
            v6 -= (v6 / ((0) | 1))
            var v54 float32 = 3.14159
            _ = v54
        } else {
            delete(v3, (^(fi53)))
            var v55 uint = 0
            switch fi53 {
            case 42:
                fmt.Println(v5)
            case 2147483647:
                if false {
                    v56 := GenS4{F0: 2, F1: 42, F2: -128}
                    var v57 int = fi53
                    _ = v56
                    _ = v57
                } else {
                    v58 := GenS1{F0: 0, F1: 0, F2: -1}
                    for fi59 := 0; fi59 < 14; fi59++ {
                        if false {
                            fmt.Println(fi59)
                            v2 = v2
                            var v60 uint8 = 2
                            _ = v60
                        } else {
                            var v61 uint16 = 42
                            var v62 int32 = 2147483647
                            _ = v61
                            _ = v62
                        }
                        fmt.Println(v6, v2)
                        _ = fi59
                    }
                    _ = v58
                }
            case 127:
                switch "\n\t" {
                case "":
                    fmt.Println(v55, v5, v2)
                case "\n\t":
                    switch "" {
                    case "":
                        v5 = v5
                    case "a":
                        if true {
                            switch fi53 {
                            case 42:
                                var v63 uint = v55
                                _ = v63
                            }
                            for fi64 := 0; fi64 < 13; fi64++ {
                                if (!(0 <= 255)) {
                                    if (1.5 >= ((1.5 * 0.0) / 3.14159)) {
                                        var v65 uint64 = (v6 / ((v5) | 1))
                                        v1 = append(v1, -359)
                                        v1 = append(v1, fi53)
                                        v66 := v4["hello"]
                                        if v66 {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v67 bool = (!(!true))
                                            if (1 >= 0) {
                                                if (!(255 <= (^(v55)))) {
                                                    fmt.Println(v5, v2)
                                                    var v68 string = ("a" + (v2 + (v2 + "a")))
                                                    _ = v68
                                                }
                                                if true {
                                                    var v69 float32 = 0.0
                                                    var v70 int64 = -44
                                                    var v71 uint16 = 1
                                                    _ = v69
                                                    _ = v70
                                                    _ = v71
                                                } else {
                                                    var v72 int16 = -225
                                                    _ = v72
                                                }
                                            } else {
                                                _ = len(v4)
                                                v66 = (!((false && (100 > -128)) || v67))
                                                switch (fi64 - (-(fi53))) {
                                                case 32767:
                                                    v66 = (!v67)
                                                case -9223372036854775808:
                                                    v73 := GenS5{F0: 2}
                                                    _ = v73
                                                case 0:
                                                    if (!(1 > 0)) {
                                                        v5 += (v6 | 74)
                                                    } else {
                                                        v74 := GenS5{F0: 255}
                                                        var v75 bool = v66
                                                        delete(v3, (fi53 % (((-(fi53))) | 1)))
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                default:
                                                    fmt.Println(v5)
                                                }
                                                switch (v2 + "\n\t") {
                                                case "foo bar":
                                                    switch v2 {
                                                    case "foo bar":
                                                        if v67 {
                                                            var v76 float32 = 3.14159
                                                            switch v2 {
                                                            case "\n\t":
                                                                v77 := GenS5{F0: 32767}
                                                                _ = v77
                                                            case "foo bar":
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            case "":
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            default:
                                                                v4["\n\t"] = v67
                                                            }
                                                            if (int8(v6) > 100) {
                                                                fmt.Println(v6, fi64, v5)
                                                            } else {
                                                                _ = len(v4)
                                                                v78 := GenS1{F0: (^(v55)), F1: 42, F2: -9223372036854775808}
                                                                _ = v78
                                                            }
                                                            _ = v76
                                                        }
                                                    }
                                                case "a":
                                                    var v79 int32 = 2147483647
                                                    _ = v79
                                                default:
                                                    _ = len(v1)
                                                }
                                            }
                                            _ = v67
                                        }
                                        if true {
                                            v65 -= v5
                                            var v80 uint = v55
                                            _ = v80
                                        } else {
                                            delete(v3, fi64)
                                            v81 := GenS4{F0: 666, F1: (v55 * (v55 * 18446744073709551615)), F2: 100}
                                            switch v2 {
                                            case "":
                                                v1 = append(v1, fi53)
                                            }
                                            _ = v81
                                        }
                                        var v82 int64 = 612
                                        if len(v1) > 0 {
                                            v1[0] = fi64
                                        }
                                        _ = v65
                                        _ = v66
                                        _ = v82
                                    }
                                    var v83 uint8 = 2
                                    _ = v83
                                }
                                v84 := v3[(fi64 + fi53)]
                                switch (fi53 / ((9223372036854775807) | 1)) {
                                case 774:
                                    fmt.Println(v6, v2)
                                case 1:
                                    if false {
                                        v5 = v5
                                    }
                                default:
                                    v85 := GenS3{F0: ("\n\t" != ("hello" + v2))}
                                    _ = v85
                                }
                                _ = fi64
                                _ = v84
                            }
                            for fi86 := 0; fi86 < 4; fi86++ {
                                var v87 uint = (v55 - (v55 & (v55 & v55)))
                                v55 += v87
                                fmt.Println(fi86, v87, v6)
                                _ = fi86
                                _ = v87
                            }
                        } else {
                            if false {
                                fmt.Println(v55, v55)
                            } else {
                                v2 = v2
                                v88 := GenS1{F0: 981, F1: 1, F2: -974}
                                for fi89 := 0; fi89 < 9; fi89++ {
                                    v90 := GenS1{F0: 717, F1: 551, F2: 42}
                                    _ = fi89
                                    _ = v90
                                }
                                _ = v88
                            }
                            if len(v1) > 0 {
                                v1[0] = (fi53 + (fi53 << ((255) & 63)))
                            }
                            for fi91 := 0; fi91 < 12; fi91++ {
                                var v92 float64 = -2.25
                                _ = fi91
                                _ = v92
                            }
                        }
                    default:
                        for fi93 := 0; fi93 < 20; fi93++ {
                            var v94 float32 = (0.1 / -2.25)
                            var v95 uint8 = 165
                            _ = len(v1)
                            v96 := GenS6{F0: "\n\t"}
                            fmt.Println(fi53)
                            _ = fi93
                            _ = v94
                            _ = v95
                            _ = v96
                        }
                    }
                default:
                    for fi97 := 0; fi97 < 13; fi97++ {
                        if ((1000 != 32767) || true) {
                            v98 := GenS7{F0: v2, F1: 42}
                            _ = v98
                        }
                        _ = fi97
                    }
                }
            default:
                var v99 string = (v2 + v2)
                _ = v99
            }
            _ = v55
        }
        _ = fi53
    }
    if true {
        if true {
            if false {
                switch -1 {
                case 4294967295:
                    v6 += 2
                case -128:
                    for fi100 := 0; fi100 < 3; fi100++ {
                        v101 := GenS8{F0: 0, F1: true}
                        for fi102 := 0; fi102 < 15; fi102++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v103 int32 = -569
                            _ = fi102
                            _ = v103
                        }
                        if ("hello" >= v2) {
                            if true {
                                var v104 int = fi100
                                _ = v104
                            }
                            v105 := v3[(fi100 & (^(fi100)))]
                            _ = v105
                        } else {
                            v106 := GenS7{F0: "\n\t", F1: 81}
                            _ = v106
                        }
                        _ = fi100
                        _ = v101
                    }
                default:
                    _ = len(v4)
                }
                if true {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if len(v1) > 0 {
                        v1[0] = 288
                    }
                    if false {
                        for fi107 := 0; fi107 < 16; fi107++ {
                            v2 = (v2 + v2)
                            fmt.Println(v5)
                            _ = fi107
                        }
                    } else {
                        if true {
                            fmt.Println(v2, v5, v5)
                            switch "" {
                            case "\n\t":
                                if ((-1 <= -235) || false) {
                                    var v108 int8 = -1
                                    var v109 float64 = (3.14159 - ((0.0 - 1e10) + 0.0))
                                    if false {
                                        var v110 float32 = ((0.0 / 0.1) + 3.14159)
                                        v2 = v2
                                        _ = v110
                                    }
                                    _ = v108
                                    _ = v109
                                }
                            case "foo bar":
                                _ = len(v1)
                            case "":
                                v111 := GenS9{F0: 0}
                                _ = v111
                            }
                        } else {
                            v2 = v2
                        }
                        for fi112 := 0; fi112 < 7; fi112++ {
                            var v113 int64 = 0
                            v1 = append(v1, 128)
                            var v114 int = 128
                            _ = fi112
                            _ = v113
                            _ = v114
                        }
                    }
                } else {
                    var v115 uint = 0
                    for fi116 := 0; fi116 < 6; fi116++ {
                        var v117 uint16 = 65535
                        for fi118 := 0; fi118 < 1; fi118++ {
                            switch (fi116 ^ 114) {
                            case 340:
                                v6 = 4294967295
                            default:
                                var v119 uint8 = 255
                                _ = v119
                            }
                            _ = fi118
                        }
                        switch (^(fi116)) {
                        case 831:
                            v5 = uint64(v117)
                        default:
                            v120 := GenS7{F0: (v2 + v2), F1: 1}
                            _ = v120
                        }
                        _ = fi116
                        _ = v117
                    }
                    delete(v4, v2)
                    _ = v115
                }
                var v121 int = 4294967295
                _ = v121
            }
            var v122 int8 = 78
            for fi123 := 0; fi123 < 13; fi123++ {
                var v124 uint16 = 2
                _ = fi123
                _ = v124
            }
            _ = v122
        }
        v5 -= 0
    }
    v5 -= v6
    v125 := GenS10{F0: 255, F1: -186}
    if (false && (1 != -1)) {
        v5 += (^(v6))
        var v126 int32 = 0
        v126 = 1
        _ = v126
    } else {
        for fi127 := 0; fi127 < 19; fi127++ {
            v6 = 4294967295
            var v128 int8 = 42
            _ = fi127
            _ = v128
        }
        for fi129 := 0; fi129 < 9; fi129++ {
            if true {
                switch "foo bar" {
                case "":
                    fmt.Println(fi129)
                case "hello":
                    if true {
                        fmt.Println(v5, fi129, v2)
                        var v130 float64 = ((-2.25 + (float64(-2.25) * ((-2.25 / 1.5) / (-2.25 - (1.5 * 0.1))))) * -2.25)
                        if (32767 <= -32768) {
                            switch fi129 {
                            case 4294967295:
                                for fi131 := 0; fi131 < 20; fi131++ {
                                    var v132 uint32 = 2147483648
                                    v133 := GenS7{F0: (v2 + v2), F1: 255}
                                    _ = fi131
                                    _ = v132
                                    _ = v133
                                }
                            case 0:
                                for fi134 := 0; fi134 < 17; fi134++ {
                                    var v135 int64 = -267
                                    if true {
                                        _ = len(v1)
                                        v135 = (v135 * (^(v135)))
                                        var v136 uint = 255
                                        _ = v136
                                    }
                                    _ = fi134
                                    _ = v135
                                }
                            case -2147483648:
                                if (3.14159 <= 1.5) {
                                    for fi137 := 0; fi137 < 17; fi137++ {
                                        var v138 int32 = 1
                                        var v139 int16 = 42
                                        _ = fi137
                                        _ = v138
                                        _ = v139
                                    }
                                    if false {
                                        v1 = append(v1, fi129)
                                        v130 += (v130 / 3.14159)
                                    }
                                } else {
                                    var v140 float64 = (3.14159 / (v130 + (v130 * 0.1)))
                                    v141 := v3[fi129]
                                    var v142 float32 = 1.5
                                    _ = v140
                                    _ = v141
                                    _ = v142
                                }
                            default:
                                v143 := GenS9{F0: 42}
                                _ = v143
                            }
                            var v144 float32 = -2.25
                            v145 := GenS11{F0: 32768, F1: -1, F2: 0}
                            _ = v144
                            _ = v145
                        }
                        _ = v130
                    } else {
                        var v146 float32 = 3.14159
                        v147 := GenS12{F0: 2, F1: -32768, F2: 2}
                        _ = v146
                        _ = v147
                    }
                case "foo bar":
                    v2 = "\n\t"
                }
            }
            switch fi129 {
            case 128:
                v2 = "foo bar"
            case -2147483648:
                if (986 == 42) {
                    fmt.Println(v5, v6)
                    v148 := GenS6{F0: "\n\t"}
                    var v149 uint16 = 1
                    _ = v148
                    _ = v149
                }
            }
            var v150 float32 = (-2.25 * 1e10)
            _ = fi129
            _ = v150
        }
    }
    v151 := GenS4{F0: 2, F1: 865, F2: 42}
    if (3.14159 >= 0.1) {
        v152 := v4[((v2 + ((v2 + "\n\t") + "foo bar")) + ((v2 + "hello") + ((v2 + (v2 + v2)) + ("hello" + v2))))]
        v153 := GenS3{F0: false}
        _ = v152
        _ = v153
    }
    for fi154 := 0; fi154 < 6; fi154++ {
        v155 := v4["a"]
        var v156 uint = 255
        _ = fi154
        _ = v155
        _ = v156
    }
    delete(v3, -9223372036854775808)
    var v157 int32 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v157:", v157)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v125
    _ = v151
    _ = v157
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
