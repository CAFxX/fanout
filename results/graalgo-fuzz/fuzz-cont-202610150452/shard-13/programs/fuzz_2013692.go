package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 int64; F1 uint; F2 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool; F1 string; F2 uint8 }
type GenS8 struct { F0 int8; F1 int; F2 uint32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint32; F1 bool }
type GenS12 struct { F0 uint; F1 uint32 }
type GenS13 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS14 struct { F0 string; F1 uint; F2 uint32 }
type GenS15 struct { F0 bool; F1 bool }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 int16; F1 uint8 }
type GenS18 struct { F0 bool; F1 uint32 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 uint16; F1 uint8 }
type GenS21 struct { F0 uint8; F1 bool }
type GenS22 struct { F0 string; F1 string; F2 string }
type GenS23 struct { F0 int8 }
type GenS24 struct { F0 int8 }
type GenS25 struct { F0 int64 }
type GenS26 struct { F0 int64; F1 uint16 }
type GenS27 struct { F0 uint64; F1 int32 }
type GenS28 struct { F0 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 333
    var v3 int64 = 0
    var v4 int = 32768
    var v5 int16 = (v2 << ((2) & 63))
    var v6 string = "\n\t"
    var v7 int32 = -2147483648
    var v8 float64 = -2.25
    if (542 < 1) {
        if false {
            for fi9 := 0; fi9 < 13; fi9++ {
                fmt.Println(v5, fi9)
                v4 -= int(v5)
                fmt.Println(v2)
                _ = fi9
            }
            var v10 int32 = v7
            _ = v10
        } else {
            for fi11 := 0; fi11 < 6; fi11++ {
                v2 = v5
                v4 = fi11
                _ = fi11
            }
            switch v4 {
            case -654:
                v12 := v1[""]
                _ = v12
            }
            if (!(4294967295 == 2147483647)) {
                v13 := GenS1{F0: ("hello" + (("hello" + "a") + (("a" + (v6 + v6)) + ("foo bar" + ("hello" + "a"))))), F1: 665}
                if false {
                    v1[v6] = 999
                    v14 := GenS1{F0: v6, F1: 65535}
                    _ = v14
                } else {
                    if (false && (!false)) {
                        var v15 int32 = -425
                        v15 += v15
                        v16 := []int{}
                        _ = v15
                        _ = v16
                    }
                    if (!true) {
                        if false {
                            for fi17 := 0; fi17 < 16; fi17++ {
                                if (!false) {
                                    v18 := GenS2{F0: (v3 + v3), F1: 0}
                                    if ("foo bar" <= "foo bar") {
                                        v19 := v1["\n\t"]
                                        _ = v19
                                    } else {
                                        var v20 uint16 = 634
                                        _ = v20
                                    }
                                    fmt.Println(fi17, v6, fi17)
                                    _ = v18
                                } else {
                                    if (false || false) {
                                        if true {
                                            fmt.Println(v7)
                                            if (127 != 0) {
                                                for fi21 := 0; fi21 < 7; fi21++ {
                                                    var v22 string = v6
                                                    _ = fi21
                                                    _ = v22
                                                }
                                            }
                                        }
                                        v23 := GenS1{F0: v6, F1: 4294967295}
                                        for fi24 := 0; fi24 < 15; fi24++ {
                                            var v25 uint8 = 166
                                            var v26 int16 = 899
                                            _ = fi24
                                            _ = v25
                                            _ = v26
                                        }
                                        _ = v23
                                    } else {
                                        var v27 int64 = -224
                                        v2 = (v2 % ((v2) | 1))
                                        _ = v27
                                    }
                                }
                                _ = fi17
                            }
                            var v28 float32 = (1.5 + 0.0)
                            _ = v28
                        }
                        if (42 > 1) {
                            for fi29 := 0; fi29 < 14; fi29++ {
                                v4 -= 32767
                                _ = fi29
                            }
                            v7 = (v7 * (v7 & (v7 & (v7 >> ((307) & 63)))))
                        } else {
                            v3 = -1
                        }
                    }
                }
                for fi30 := 0; fi30 < 19; fi30++ {
                    for fi31 := 0; fi31 < 13; fi31++ {
                        fmt.Println(fi30)
                        v7 += -2147483648
                        v32 := GenS1{F0: v6, F1: 77}
                        _ = fi31
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 10; fi33++ {
                        var v34 string = v6
                        switch v34 {
                        case "a":
                            var v35 string = ("\n\t" + (v34 + v34))
                            _ = v35
                        case "\n\t":
                            delete(v1, (((((v6 + v6) + v6) + v34) + "a") + v6))
                        }
                        v36 := []int{(-(v4)), fi30, 2147483647}
                        _ = fi33
                        _ = v34
                        _ = v36
                    }
                    _ = fi30
                }
                _ = v13
            } else {
                switch v4 {
                case 127:
                    v8 = 1e10
                }
                var v37 int8 = 100
                var v38 int32 = (^(v7))
                _ = v37
                _ = v38
            }
        }
    } else {
        v39 := GenS2{F0: -771, F1: 1}
        _ = v39
    }
    for fi40 := 0; fi40 < 19; fi40++ {
        if (false || false) {
            v41 := []int{v4}
            if true {
                var v42 uint8 = 212
                v6 = v6
                _ = v42
            } else {
                fmt.Println(v3, v7)
                for fi43 := 0; fi43 < 5; fi43++ {
                    if (true && false) {
                        var v44 int8 = -110
                        var v45 int8 = -1
                        var v46 uint64 = 18446744073709551615
                        _ = v44
                        _ = v45
                        _ = v46
                    } else {
                        v47 := GenS2{F0: v3, F1: 217}
                        fmt.Println(v4, v6)
                        v48 := GenS2{F0: (v3 / ((0) | 1)), F1: 42}
                        _ = v47
                        _ = v48
                    }
                    fmt.Println(v8, v2, v4)
                    var v49 float32 = (3.14159 - float32((float64(float32((v8 + 3.14159))) / v8)))
                    _ = fi43
                    _ = v49
                }
            }
            var v50 uint64 = 18446744073709551615
            _ = v41
            _ = v50
        } else {
            v5 += 1
        }
        _ = fi40
    }
    var v51 uint16 = 42
    if false {
        _ = len(v1)
    }
    switch (^(v4)) {
    case 128:
        if true {
            var v52 uint16 = (v51 & (v51 * v51))
            switch v6 {
            case "a":
                if true {
                    if false {
                        var v53 uint = 18446744073709551615
                        var v54 bool = true
                        v7 = 179
                        _ = v53
                        _ = v54
                    } else {
                        _ = len(v1)
                    }
                    if (false && false) {
                        for fi55 := 0; fi55 < 12; fi55++ {
                            v52 -= 25
                            var v56 float32 = 0.0
                            fmt.Println(v56, v52)
                            _ = fi55
                            _ = v56
                        }
                        fmt.Println(v7)
                        if true {
                            fmt.Println(v52, v5)
                            v1[v6] = 616
                            for fi57 := 0; fi57 < 4; fi57++ {
                                v58 := GenS3{F0: 52}
                                v6 = v6
                                for fi59 := 0; fi59 < 12; fi59++ {
                                    v60 := []int{-1, (^(fi59)), v4, fi59}
                                    v6 = ((v6 + "a") + "\n\t")
                                    _ = fi59
                                    _ = v60
                                }
                                _ = fi57
                                _ = v58
                            }
                        } else {
                            v61 := GenS3{F0: 18446744073709551615}
                            v51 = v52
                            for fi62 := 0; fi62 < 12; fi62++ {
                                var v63 int32 = (v7 << ((255) & 63))
                                var v64 float64 = v8
                                _ = fi62
                                _ = v63
                                _ = v64
                            }
                            _ = v61
                        }
                    } else {
                        switch "hello" {
                        case "hello":
                            if (false && true) {
                                var v65 uint64 = 540
                                v51 += v52
                                _ = v65
                            } else {
                                v6 = v6
                                if true {
                                    v66 := []int{v4, v4, (v4 - -9223372036854775808), v4}
                                    if (v3 > v3) {
                                        var v67 int = (v4 + (v4 << ((1) & 63)))
                                        v68 := v1["foo bar"]
                                        var v69 int8 = -128
                                        _ = v67
                                        _ = v68
                                        _ = v69
                                    } else {
                                        v70 := GenS1{F0: "hello", F1: 4294967295}
                                        for fi71 := 0; fi71 < 5; fi71++ {
                                            if true {
                                                switch ("\n\t" + v6) {
                                                case "foo bar":
                                                    v72 := GenS1{F0: (v6 + v6), F1: 0}
                                                    _ = v72
                                                case "":
                                                    v73 := GenS4{F0: ((!(0 <= (v7 >> ((65535) & 63)))) && false), F1: 4294967295}
                                                    _ = v73
                                                }
                                                if true {
                                                    fmt.Println(v51)
                                                    if (!true) {
                                                        v52 -= v52
                                                        if len(v66) > 0 {
                                                            v66[0] = -9223372036854775808
                                                        }
                                                        var v74 uint = 0
                                                        _ = v74
                                                    } else {
                                                        v75 := GenS5{F0: (-(v3)), F1: 4294967295, F2: true}
                                                        fmt.Println(v6, v52, v7)
                                                        _ = v75
                                                    }
                                                }
                                            }
                                            if true {
                                                v66 = append(v66, (-(fi71)))
                                                for fi76 := 0; fi76 < 6; fi76++ {
                                                    v77 := GenS6{F0: fi71}
                                                    switch fi71 {
                                                    case 1:
                                                        if ((v5 * -32768) >= int16(v7)) {
                                                            fmt.Println(fi71, v51, fi71)
                                                            var v78 int64 = 9223372036854775807
                                                            fmt.Println(v8, v51)
                                                            _ = v78
                                                        }
                                                    case -694:
                                                        var v79 uint64 = 0
                                                        _ = v79
                                                    case 9223372036854775807:
                                                        if (true || true) {
                                                            switch ((("\n\t" + "\n\t") + v6) + v6) {
                                                            case "hello":
                                                                if false {
                                                                    var v80 int16 = 1000
                                                                    v81 := GenS1{F0: (v6 + "hello"), F1: 18446744073709551615}
                                                                    if ((!false) || ((4294967295 >= 1) || (127 > (fi76 / ((v4) | 1))))) {
                                                                        v82 := GenS2{F0: (v3 << ((4294967295) & 63)), F1: 0}
                                                                        switch "" {
                                                                        case "foo bar":
                                                                            var v83 int64 = v3
                                                                            _ = v83
                                                                        default:
                                                                            v84 := GenS5{F0: v3, F1: 255, F2: ((v3 & 9223372036854775807) <= v3)}
                                                                            _ = v84
                                                                        }
                                                                        _ = v82
                                                                    } else {
                                                                        var v85 uint = 255
                                                                        v85 -= v85
                                                                        var v86 uint = v85
                                                                        if len(v66) > 0 {
                                                                            v66[0] = fi71
                                                                        }
                                                                        _ = v85
                                                                        _ = v86
                                                                    }
                                                                    v1["a"] = 595
                                                                    _ = v80
                                                                    _ = v81
                                                                } else {
                                                                    v7 = 1
                                                                }
                                                            }
                                                            v87 := GenS3{F0: 42}
                                                            _ = v87
                                                        }
                                                    default:
                                                        v88 := GenS1{F0: "", F1: 255}
                                                        _ = v88
                                                    }
                                                    _ = fi76
                                                    _ = v77
                                                }
                                            } else {
                                                if (!true) {
                                                    switch -696 {
                                                    case -780:
                                                        if len(v66) >= 2 {
                                                            v66 = v66[1:len(v66)]
                                                        }
                                                    default:
                                                        var v89 bool = (-77 >= 100)
                                                        _ = v89
                                                    }
                                                    if (true && true) {
                                                        var v90 uint64 = 387
                                                        fmt.Println(v90)
                                                        v52 = (v51 * v51)
                                                        _ = v90
                                                    } else {
                                                        v6 = v6
                                                        if true {
                                                            v91 := GenS7{F0: (!true), F1: "a", F2: 0}
                                                            v51 = v51
                                                            _ = len(v1)
                                                            _ = v91
                                                        } else {
                                                            switch v4 {
                                                            case 1:
                                                                _ = len(v66)
                                                            default:
                                                                fmt.Println(v4, v8, v7)
                                                            }
                                                            switch v6 {
                                                            case "a":
                                                                if true {
                                                                    v92 := GenS8{F0: -1, F1: 128, F2: 2147483648}
                                                                    v93 := v1[v6]
                                                                    _ = v92
                                                                    _ = v93
                                                                } else {
                                                                    var v94 bool = (4294967295 != 4294967295)
                                                                    for fi95 := 0; fi95 < 15; fi95++ {
                                                                        v5 = -146
                                                                        _ = fi95
                                                                    }
                                                                    _ = v94
                                                                }
                                                            }
                                                            for fi96 := 0; fi96 < 1; fi96++ {
                                                                v8 = v8
                                                                _ = fi96
                                                            }
                                                        }
                                                        var v97 uint64 = 18446744073709551615
                                                        _ = v97
                                                    }
                                                }
                                                var v98 float32 = (1e10 / 1e10)
                                                _ = v98
                                            }
                                            if ((v8 * v8) >= ((v8 - (v8 - float64((float32(v8) / -2.25)))) + ((v8 + ((v8 * 0.1) - (v8 - v8))) - (-2.25 - (float64(3.14159) + (((v8 - float64(((3.14159 / 0.1) * 0.0))) - float64(0.1)) - v8)))))) {
                                                _ = len(v1)
                                            } else {
                                                if true {
                                                    var v99 int32 = 42
                                                    _ = len(v1)
                                                    for fi100 := 0; fi100 < 2; fi100++ {
                                                        if ((((v3 & 699) == (v3 % ((int64(v2)) | 1))) || (!true)) && (!(("foo bar" + v6) > v6))) {
                                                            if true {
                                                                var v101 uint64 = 4294967295
                                                                fmt.Println(v51, v101)
                                                                _ = v101
                                                            }
                                                            v51 -= 0
                                                            v4 -= fi100
                                                        } else {
                                                            v102 := GenS2{F0: (v3 + (v3 | 0)), F1: 748}
                                                            _ = v102
                                                        }
                                                        var v103 int = fi100
                                                        _ = fi100
                                                        _ = v103
                                                    }
                                                    _ = v99
                                                }
                                            }
                                            _ = fi71
                                        }
                                        _ = v70
                                    }
                                    _ = v66
                                } else {
                                    for fi104 := 0; fi104 < 12; fi104++ {
                                        v4 = v4
                                        _ = fi104
                                    }
                                }
                            }
                        }
                        var v105 int8 = 1
                        _ = v105
                    }
                } else {
                    _ = len(v1)
                    v5 += 32767
                    v8 = (1e10 + 1e10)
                }
            case "\n\t":
                for fi106 := 0; fi106 < 13; fi106++ {
                    for fi107 := 0; fi107 < 10; fi107++ {
                        switch v6 {
                        case "foo bar":
                            if (((false || (((v6 + (("" + v6) + (v6 + "hello"))) + v6) != v6)) && true) && ((v4 & 128) >= fi106)) {
                                switch v6 {
                                case "":
                                    v108 := []int{fi107, v4, (fi106 & fi107), (v4 + fi106)}
                                    _ = v108
                                case "a":
                                    var v109 uint = 1
                                    _ = v109
                                case "\n\t":
                                    switch 4 {
                                    case -1:
                                        var v110 int32 = 401
                                        switch ((v6 + "") + (("" + v6) + (v6 + v6))) {
                                        case "hello":
                                            v6 = v6
                                        default:
                                            v7 -= (^(v7))
                                        }
                                        if false {
                                            var v111 int64 = (v3 & v3)
                                            _ = v111
                                        } else {
                                            v112 := []int{(fi107 / (((v4 & 4294967295)) | 1)), v4, fi107, v4}
                                            var v113 int8 = 47
                                            _ = v112
                                            _ = v113
                                        }
                                        switch v6 {
                                        case "\n\t":
                                            switch v6 {
                                            case "hello":
                                                if (!false) {
                                                    var v114 uint8 = 128
                                                    switch (v4 * fi107) {
                                                    case 223:
                                                        v5 = v2
                                                    case -260:
                                                        v7 = 1
                                                    case 0:
                                                        v115 := GenS9{F0: (v114 >> ((255) & 63))}
                                                        _ = v115
                                                    default:
                                                        switch "a" {
                                                        case "\n\t":
                                                            v116 := GenS9{F0: (v114 + 2)}
                                                            _ = v116
                                                        }
                                                    }
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        fmt.Println(fi106, fi107)
                                                        var v117 string = (("foo bar" + "\n\t") + "a")
                                                        switch (fi107 >> ((42) & 63)) {
                                                        case 32767:
                                                            var v118 uint8 = 237
                                                            _ = v118
                                                        case 32768:
                                                            var v119 string = ((v117 + v117) + v117)
                                                            _ = v119
                                                        case 127:
                                                            switch (v4 ^ -9223372036854775808) {
                                                            case -942:
                                                                switch (fi106 << ((42) & 63)) {
                                                                case 342:
                                                                    v120 := []int{-2147483648, fi106, (v4 ^ v4)}
                                                                    _ = v120
                                                                case 2147483647:
                                                                    v6 = ("hello" + v6)
                                                                case 128:
                                                                    v2 = v5
                                                                default:
                                                                    var v121 int8 = -10
                                                                    switch fi107 {
                                                                    case 32768:
                                                                        if false {
                                                                            var v122 uint32 = 65535
                                                                            if (255 >= 0) {
                                                                                v123 := []int{v4, fi106, v4, (fi107 / ((163) | 1))}
                                                                                if true {
                                                                                    v124 := GenS1{F0: v6, F1: 18446744073709551615}
                                                                                    var v125 int = (-(fi107))
                                                                                    delete(v1, v6)
                                                                                    var v126 uint = 682
                                                                                    _ = v124
                                                                                    _ = v125
                                                                                    _ = v126
                                                                                } else {
                                                                                    _ = len(v123)
                                                                                }
                                                                                _ = v123
                                                                            }
                                                                            v121 -= v121
                                                                            _ = v122
                                                                        }
                                                                    case 0:
                                                                        v114 -= v114
                                                                    default:
                                                                        fmt.Println(v2)
                                                                    }
                                                                    fmt.Println(v6)
                                                                    _ = v121
                                                                }
                                                            case 42:
                                                                v127 := GenS5{F0: (v3 >> ((42) & 63)), F1: 199, F2: (true || true)}
                                                                _ = v127
                                                            case 2147483647:
                                                                var v128 int64 = 0
                                                                _ = v128
                                                            }
                                                        }
                                                        var v129 uint16 = v52
                                                        _ = v117
                                                        _ = v129
                                                    case "foo bar":
                                                        var v130 uint16 = v51
                                                        _ = v130
                                                    case "hello":
                                                        var v131 int = (-(v4))
                                                        _ = v131
                                                    default:
                                                        var v132 float32 = 0.1
                                                        _ = v132
                                                    }
                                                    _ = v114
                                                } else {
                                                    switch ("a" + "\n\t") {
                                                    case "a":
                                                        var v133 int16 = 410
                                                        _ = v133
                                                    case "":
                                                        var v134 uint32 = 553
                                                        _ = v134
                                                    case "hello":
                                                        v3 = (^(v3))
                                                    default:
                                                        if true {
                                                            if (4294967295 >= 2) {
                                                                v135 := GenS5{F0: (v3 - 9223372036854775807), F1: 2, F2: ((v110 / (((-(v110))) | 1)) >= 42)}
                                                                if (0 == 127) {
                                                                    v51 += (^(v51))
                                                                    if false {
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            if false {
                                                                                v3 = v3
                                                                                if ((0.0 * ((3.14159 + -2.25) - 0.1)) >= 1e10) {
                                                                                    delete(v1, v6)
                                                                                    if false {
                                                                                        switch 4294967295 {
                                                                                        case 329:
                                                                                            var v136 int = 957
                                                                                            _ = v136
                                                                                        case -260:
                                                                                            var v137 string = v6
                                                                                            _ = v137
                                                                                        default:
                                                                                            switch v6 {
                                                                                            case "a":
                                                                                                v138 := GenS10{F0: (v3 >> ((65535) & 63))}
                                                                                                var v139 int16 = 1
                                                                                                v140 := GenS11{F0: 406, F1: false}
                                                                                                _ = v138
                                                                                                _ = v139
                                                                                                _ = v140
                                                                                            case "foo bar":
                                                                                                var v141 int16 = (v5 ^ v2)
                                                                                                _ = v141
                                                                                            default:
                                                                                                v142 := GenS12{F0: 255, F1: 531}
                                                                                                _ = v142
                                                                                            }
                                                                                        }
                                                                                        v4 = 1
                                                                                    }
                                                                                    var v143 float64 = v8
                                                                                    var v144 string = v6
                                                                                    switch v4 {
                                                                                    case 1:
                                                                                        v110 -= v110
                                                                                        var v145 float32 = 1.5
                                                                                        var v146 bool = (!(!true))
                                                                                        _ = v145
                                                                                        _ = v146
                                                                                    case 127:
                                                                                        v5 = (-(v2))
                                                                                    case -979:
                                                                                        var v147 uint64 = uint64(v52)
                                                                                        _ = v147
                                                                                    default:
                                                                                        v148 := []int{}
                                                                                        _ = v148
                                                                                    }
                                                                                    _ = v143
                                                                                    _ = v144
                                                                                } else {
                                                                                    fmt.Println(fi107)
                                                                                }
                                                                            } else {
                                                                                v149 := GenS13{F0: v2, F1: 10, F2: 2}
                                                                                _ = v149
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v150 := []int{32768}
                                                                        v8 += 1e10
                                                                        _ = v150
                                                                    }
                                                                    switch 32768 {
                                                                    case 1:
                                                                        fmt.Println(v3, fi107, v5)
                                                                    }
                                                                }
                                                                fmt.Println(v3, v8)
                                                                _ = v135
                                                            } else {
                                                                var v151 int16 = (v2 >> ((65535) & 63))
                                                                _ = v151
                                                            }
                                                            var v152 int16 = v2
                                                            var v153 string = (v6 + v6)
                                                            _ = v152
                                                            _ = v153
                                                        }
                                                    }
                                                    if (!true) {
                                                        v6 = v6
                                                    } else {
                                                        v154 := GenS7{F0: (2 <= 252), F1: "foo bar", F2: 173}
                                                        _ = v154
                                                    }
                                                    v155 := v1[((("" + v6) + "a") + v6)]
                                                    _ = v155
                                                }
                                            case "":
                                                delete(v1, (v6 + ("" + (v6 + ((v6 + "hello") + v6)))))
                                            }
                                        case "a":
                                            v6 = v6
                                        case "hello":
                                            _ = len(v1)
                                        }
                                        _ = v110
                                    default:
                                        v2 += v5
                                    }
                                }
                                v2 = (-(v2))
                                v156 := v1[(((v6 + ("a" + (v6 + "foo bar"))) + v6) + v6)]
                                _ = v156
                            }
                        }
                        if false {
                            v52 += 1
                        } else {
                            fmt.Println(fi106)
                            v157 := GenS5{F0: v3, F1: 0, F2: (true && true)}
                            _ = v157
                        }
                        _ = fi107
                    }
                    v4 = 9223372036854775807
                    v8 -= v8
                    _ = fi106
                }
            case "":
                if (!(!true)) {
                    v6 = ("foo bar" + v6)
                    _ = len(v1)
                } else {
                    v8 -= (v8 * v8)
                    var v158 uint32 = 65535
                    fmt.Println(v2)
                    _ = v158
                }
            }
            if (0.1 > 1.5) {
                switch v6 {
                case "\n\t":
                    v159 := []int{v4, v4, v4, v4}
                    _ = v159
                case "hello":
                    if true {
                        v2 = (v2 ^ v5)
                    }
                default:
                    v7 += v7
                }
                var v160 int = (v4 + (v4 * (-(v4))))
                switch ("foo bar" + v6) {
                case "\n\t":
                    var v161 uint = 1
                    _ = v161
                default:
                    v162 := GenS9{F0: 1}
                    _ = v162
                }
                _ = v160
            }
            _ = v52
        }
    case -709:
        v7 -= v7
    case -9223372036854775808:
        for fi163 := 0; fi163 < 11; fi163++ {
            if false {
                v164 := GenS13{F0: int16(v7), F1: 42, F2: 623}
                _ = v164
            } else {
                v3 = v3
            }
            switch (v4 - -1) {
            case -2147483648:
                for fi165 := 0; fi165 < 15; fi165++ {
                    fmt.Println(v5, v2, v51)
                    v166 := GenS14{F0: v6, F1: 65535, F2: 1}
                    _ = fi165
                    _ = v166
                }
            case 139:
                var v167 int8 = 42
                _ = v167
            case 128:
                if (759 > 18446744073709551615) {
                    if (-80 > 0) {
                        switch v6 {
                        case "foo bar":
                            if (2 <= 234) {
                                var v168 uint32 = 42
                                _ = v168
                            } else {
                                fmt.Println(v3, v2)
                            }
                        case "\n\t":
                            switch v6 {
                            case "":
                                fmt.Println(v4, v8, v7)
                            case "foo bar":
                                var v169 uint64 = 837
                                _ = v169
                            }
                        case "a":
                            v170 := GenS15{F0: false, F1: false}
                            _ = v170
                        default:
                            for fi171 := 0; fi171 < 18; fi171++ {
                                v1["hello"] = 0
                                _ = fi171
                            }
                        }
                        switch (fi163 & int(v51)) {
                        case 1:
                            delete(v1, v6)
                        case -920:
                            v172 := GenS9{F0: 127}
                            _ = v172
                        case 128:
                            fmt.Println(v7)
                        default:
                            v7 = (-(v7))
                        }
                        v173 := GenS5{F0: (v3 - v3), F1: 4294967295, F2: false}
                        _ = v173
                    }
                }
            }
            _ = fi163
        }
    default:
        v8 = 1e10
    }
    v6 = ""
    var v174 float32 = (0.1 * float32(((float64(-2.25) * 0.0) / v8)))
    var v175 uint64 = 18446744073709551615
    if (float64(v174) != 0.1) {
        if true {
            v176 := v1[v6]
            v177 := GenS16{F0: (v5 | -32768)}
            _ = v176
            _ = v177
        }
        var v178 float64 = v8
        _ = v178
    } else {
        if false {
            var v179 string = "hello"
            _ = v179
        } else {
            v180 := GenS12{F0: 42, F1: 676}
            if true {
                switch (v4 / (((v4 ^ (v4 / ((v4) | 1)))) | 1)) {
                case -2147483648:
                    var v181 int8 = 100
                    _ = v181
                case 9223372036854775807:
                    v182 := GenS17{F0: 42, F1: 233}
                    _ = v182
                default:
                    var v183 uint = 819
                    _ = v183
                }
            } else {
                var v184 string = "foo bar"
                switch "a" {
                case "foo bar":
                    switch (("foo bar" + v6) + v184) {
                    case "\n\t":
                        v185 := GenS18{F0: true, F1: 0}
                        _ = v185
                    case "hello":
                        for fi186 := 0; fi186 < 9; fi186++ {
                            if true {
                                var v187 int64 = v3
                                v188 := GenS19{F0: false}
                                _ = v187
                                _ = v188
                            }
                            _ = fi186
                        }
                    case "":
                        if false {
                            switch v184 {
                            case "":
                                switch (v4 % ((v4) | 1)) {
                                case 245:
                                    var v189 int16 = v5
                                    _ = v189
                                case 9223372036854775807:
                                    v190 := GenS5{F0: (v3 & v3), F1: 65535, F2: false}
                                    _ = v190
                                }
                            case "foo bar":
                                for fi191 := 0; fi191 < 12; fi191++ {
                                    for fi192 := 0; fi192 < 13; fi192++ {
                                        if false {
                                            v193 := GenS8{F0: 127, F1: 42, F2: 13}
                                            if (!(!false)) {
                                                v194 := []int{fi192, (fi191 % (((fi192 + (-(fi191)))) | 1)), (v4 << ((1) & 63)), 598}
                                                switch v184 {
                                                case "hello":
                                                    var v195 int64 = 42
                                                    _ = v195
                                                case "\n\t":
                                                    if len(v194) >= 2 {
                                                        v194 = v194[1:len(v194)]
                                                    }
                                                    fmt.Println(v7, v4)
                                                case "foo bar":
                                                    switch (-(fi191)) {
                                                    case 2147483647:
                                                        if len(v194) >= 2 {
                                                            v194 = v194[1:len(v194)]
                                                        }
                                                    case 570:
                                                        switch fi191 {
                                                        case 280:
                                                            var v196 uint16 = v51
                                                            _ = v196
                                                        case -832:
                                                            var v197 int8 = int8(v51)
                                                            _ = v197
                                                        }
                                                    default:
                                                        var v198 uint = 42
                                                        _ = v198
                                                    }
                                                default:
                                                    v199 := GenS19{F0: false}
                                                    _ = v199
                                                }
                                                fmt.Println(v5, v5, v6)
                                                v200 := GenS9{F0: 42}
                                                _ = v194
                                                _ = v200
                                            }
                                            _ = v193
                                        } else {
                                            v6 = (v6 + "")
                                        }
                                        switch (fi191 * (v4 << ((65535) & 63))) {
                                        case 42:
                                            var v201 uint16 = 32767
                                            _ = v201
                                        default:
                                            v202 := []int{(-(v4)), (-(fi192))}
                                            v203 := GenS12{F0: 0, F1: 716}
                                            _ = v202
                                            _ = v203
                                        }
                                        var v204 int32 = v7
                                        _ = fi192
                                        _ = v204
                                    }
                                    _ = fi191
                                }
                            case "\n\t":
                                var v205 int64 = (v3 ^ v3)
                                _ = v205
                            default:
                                if (255 >= 255) {
                                    v206 := []int{-1, v4}
                                    for fi207 := 0; fi207 < 6; fi207++ {
                                        for fi208 := 0; fi208 < 19; fi208++ {
                                            v209 := GenS20{F0: v51, F1: 127}
                                            _ = fi208
                                            _ = v209
                                        }
                                        v5 = int16(v3)
                                        for fi210 := 0; fi210 < 12; fi210++ {
                                            v211 := v1[("\n\t" + "")]
                                            v4 -= 4294967295
                                            if true {
                                                if false {
                                                    var v212 int8 = int8(v4)
                                                    v206 = append(v206, (v4 + (fi210 << ((849) & 63))))
                                                    v213 := GenS21{F0: 1, F1: (v3 > -9223372036854775808)}
                                                    _ = v212
                                                    _ = v213
                                                }
                                                fmt.Println(v174, v175)
                                            } else {
                                                if (fi207 < (^(fi207))) {
                                                    v211 = v211
                                                    v211 = 1
                                                    fmt.Println(v2, v211)
                                                }
                                                if true {
                                                    var v214 int8 = -85
                                                    var v215 uint16 = 65535
                                                    switch fi207 {
                                                    case 32767:
                                                        if (0 >= (v211 * uint32(v175))) {
                                                            fmt.Println(v214, v8, v7)
                                                            v216 := GenS19{F0: (9223372036854775807 != -894)}
                                                            var v217 int8 = v214
                                                            _ = v216
                                                            _ = v217
                                                        } else {
                                                            if true {
                                                                v1[(v184 + ("\n\t" + "hello"))] = (v211 % (((v211 + (^(v211)))) | 1))
                                                                v175 += v175
                                                                var v218 uint16 = v51
                                                                var v219 uint8 = 36
                                                                _ = v218
                                                                _ = v219
                                                            } else {
                                                                switch 127 {
                                                                case -9223372036854775808:
                                                                    var v220 bool = (128 < 44)
                                                                    _ = v220
                                                                }
                                                                if false {
                                                                    if (int8(v211) >= (v214 | -87)) {
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            v206 = append(v206, fi207)
                                                                        case "foo bar":
                                                                            var v221 int8 = (v214 / ((42) | 1))
                                                                            _ = v221
                                                                        }
                                                                        v214 = (v214 / (((^(v214))) | 1))
                                                                    } else {
                                                                        v206 = append(v206, 1)
                                                                        var v222 uint64 = v175
                                                                        v175 = v175
                                                                        v223 := GenS22{F0: ("foo bar" + v6), F1: "foo bar", F2: ((((v184 + ("" + v6)) + v6) + v6) + ("" + "a"))}
                                                                        _ = v222
                                                                        _ = v223
                                                                    }
                                                                } else {
                                                                    delete(v1, (v184 + "\n\t"))
                                                                }
                                                            }
                                                            var v224 int16 = v2
                                                            _ = v224
                                                        }
                                                    case -1:
                                                        v225 := GenS23{F0: (v214 % (((v214 / ((v214) | 1))) | 1))}
                                                        _ = v225
                                                    default:
                                                        _ = len(v1)
                                                        v226 := GenS24{F0: (v214 % (((v214 / ((int8(v211)) | 1))) | 1))}
                                                        v5 = (v5 << ((18446744073709551615) & 63))
                                                        if false {
                                                            fmt.Println(v5, v214, v5)
                                                        }
                                                        if true {
                                                            switch ((v184 + v184) + "foo bar") {
                                                            case "a":
                                                                var v227 uint64 = v175
                                                                _ = v227
                                                            case "foo bar":
                                                                var v228 int64 = (v3 - (v3 & (-(v3))))
                                                                _ = v228
                                                            case "\n\t":
                                                                if (false || (-2147483648 > 42)) {
                                                                    var v229 int32 = -140
                                                                    fmt.Println(fi207, v3)
                                                                    v230 := GenS25{F0: (v3 & (v3 * -846))}
                                                                    _ = v229
                                                                    _ = v230
                                                                }
                                                            default:
                                                                var v231 int8 = (v214 / ((91) | 1))
                                                                _ = v231
                                                            }
                                                            v211 -= v211
                                                            v174 += v174
                                                        } else {
                                                            v214 -= (^(v214))
                                                            v175 = (v175 - v175)
                                                        }
                                                        var v232 int16 = (v2 - (v5 << ((255) & 63)))
                                                        switch (-(fi207)) {
                                                        case 4294967295:
                                                            switch v6 {
                                                            case "foo bar":
                                                                if false {
                                                                    if true {
                                                                        var v233 uint8 = 255
                                                                        v206 = append(v206, -9223372036854775808)
                                                                        switch 886 {
                                                                        case -234:
                                                                            switch (fi210 % ((v4) | 1)) {
                                                                            case 32768:
                                                                                v184 = (v184 + (v184 + ("" + v184)))
                                                                            }
                                                                        case -764:
                                                                            var v234 float64 = (v8 / v8)
                                                                            if len(v206) >= 2 {
                                                                                v206 = v206[1:len(v206)]
                                                                            }
                                                                            var v235 uint16 = (^(v215))
                                                                            _ = v234
                                                                            _ = v235
                                                                        default:
                                                                            var v236 uint16 = (^(v215))
                                                                            _ = v236
                                                                        }
                                                                        _ = v233
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        fmt.Println(v3, v232, v5)
                                                                    } else {
                                                                        if len(v206) > 0 {
                                                                            v206[0] = fi207
                                                                        }
                                                                    }
                                                                    v1["a"] = 2
                                                                    switch ((v184 + (v6 + ("" + "foo bar"))) + "\n\t") {
                                                                    case "hello":
                                                                        var v237 int32 = (-(v7))
                                                                        _ = v237
                                                                    case "foo bar":
                                                                        delete(v1, v6)
                                                                    case "\n\t":
                                                                        var v238 int32 = v7
                                                                        _ = v238
                                                                    default:
                                                                        fmt.Println(v174, fi207, v4)
                                                                    }
                                                                }
                                                            }
                                                        case 9223372036854775807:
                                                            switch (-(fi207)) {
                                                            case 917:
                                                                v239 := GenS26{F0: 42, F1: (^(v51))}
                                                                var v240 int16 = (v2 ^ 1000)
                                                                _ = v239
                                                                _ = v240
                                                            case 9223372036854775807:
                                                                var v241 float32 = v174
                                                                _ = v241
                                                            }
                                                        default:
                                                            v51 = v51
                                                        }
                                                        _ = v226
                                                        _ = v232
                                                    }
                                                    _ = v214
                                                    _ = v215
                                                } else {
                                                    if len(v206) > 0 {
                                                        v206[0] = 0
                                                    }
                                                    var v242 uint32 = v211
                                                    v6 = (v184 + ("\n\t" + ""))
                                                    _ = v242
                                                }
                                            }
                                            _ = fi210
                                            _ = v211
                                        }
                                        _ = fi207
                                    }
                                    _ = v206
                                }
                            }
                            var v243 float64 = float64(v174)
                            _ = v243
                        } else {
                            var v244 uint64 = 2
                            _ = v244
                        }
                    default:
                        for fi245 := 0; fi245 < 5; fi245++ {
                            for fi246 := 0; fi246 < 8; fi246++ {
                                if true {
                                    delete(v1, "hello")
                                    var v247 int = fi246
                                    v248 := GenS1{F0: v6, F1: 612}
                                    _ = v247
                                    _ = v248
                                } else {
                                    switch (fi245 - fi246) {
                                    case 735:
                                        var v249 uint16 = (v51 % ((uint16(v175)) | 1))
                                        _ = v249
                                    case 2147483647:
                                        v5 -= v5
                                    }
                                }
                                v250 := []int{(v4 ^ fi246)}
                                _ = fi246
                                _ = v250
                            }
                            switch v4 {
                            case 2147483647:
                                switch int(v175) {
                                case 4294967295:
                                    v251 := v1[v184]
                                    _ = v251
                                case -9223372036854775808:
                                    fmt.Println(fi245, v8, v7)
                                case 32:
                                    fmt.Println(v175)
                                }
                            case 45:
                                v252 := GenS27{F0: (v175 * v175), F1: int32(v3)}
                                _ = v252
                            case -9223372036854775808:
                                if (v174 < v174) {
                                    for fi253 := 0; fi253 < 12; fi253++ {
                                        v254 := []int{v4, v4}
                                        switch fi253 {
                                        case 178:
                                            var v255 uint16 = (v51 + v51)
                                            _ = v255
                                        default:
                                            v3 = (v3 & v3)
                                        }
                                        fmt.Println(v8, v8)
                                        _ = fi253
                                        _ = v254
                                    }
                                    fmt.Println(v6, v184)
                                    v5 = (v5 | v2)
                                }
                            }
                            v8 = float64(v174)
                            _ = fi245
                        }
                    }
                default:
                    fmt.Println(v174)
                }
                delete(v1, (v184 + v184))
                _ = v184
            }
            v256 := GenS28{F0: -1}
            _ = v180
            _ = v256
        }
        v174 -= 0.1
        var v257 uint64 = v175
        _ = v257
    }
    if (true || (v7 <= v7)) {
        v258 := GenS13{F0: -233, F1: 127, F2: 767}
        _ = v258
    }
    for fi259 := 0; fi259 < 18; fi259++ {
        var v260 float32 = 0.0
        _ = fi259
        _ = v260
    }
    v174 -= 1e10
    var v261 bool = (true && true)
    for fi262 := 0; fi262 < 5; fi262++ {
        v263 := GenS10{F0: 703}
        var v264 float32 = 1e10
        _ = fi262
        _ = v263
        _ = v264
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v174:", v174)
    fmt.Println("v175:", v175)
    fmt.Println("v2:", v2)
    fmt.Println("v261:", v261)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v174
    _ = v175
    _ = v2
    _ = v261
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v7
    _ = v8
}
