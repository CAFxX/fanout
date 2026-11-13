package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int64 }
type GenS2 struct { F0 int64; F1 string; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 string; F2 uint64 }
type GenS6 struct { F0 uint32; F1 string; F2 uint16 }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16; F1 uint8 }
type GenS11 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint32; F1 uint }
type GenS14 struct { F0 int; F1 uint16 }
type GenS15 struct { F0 uint8; F1 string; F2 int }
type GenS16 struct { F0 uint; F1 uint16 }
type GenS17 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS18 struct { F0 uint8 }
type GenS19 struct { F0 uint16; F1 uint32 }
type GenS20 struct { F0 uint }
type GenS21 struct { F0 uint8; F1 uint16 }
type GenS22 struct { F0 uint; F1 bool }

func main() {
    var v1 int8 = -128
    var v2 uint8 = 42
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    if false {
        if (!true) {
            v5 := GenS1{F0: 2147483647, F1: 184, F2: -1}
            v6 := GenS1{F0: -2147483648, F1: v2, F2: 9223372036854775807}
            _ = v5
            _ = v6
        }
        if (0 <= 2147483648) {
            v1 -= v1
            v2 = (v2 << ((18446744073709551615) & 63))
            var v7 int32 = 836
            _ = v7
        } else {
            if (!(-441 < 2147483647)) {
                v8 := GenS1{F0: 0, F1: v2, F2: 42}
                _ = v8
            }
        }
        fmt.Println(v1, v2)
    } else {
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := []int{(fi9 ^ fi9)}
            var v11 uint16 = 409
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 uint = 367
                fmt.Println(v11, v2)
                _ = fi12
                _ = v13
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        switch -2147483648 {
        case 128:
            v14 := GenS1{F0: 14, F1: v2, F2: -275}
            _ = v14
        default:
            v15 := []int{-2147483648, -2147483648}
            _ = v15
        }
        v16 := GenS1{F0: 2147483647, F1: 127, F2: -1}
        _ = v16
    }
    v17 := GenS1{F0: -2147483648, F1: (^(v2)), F2: 42}
    var v18 uint32 = 42
    v2 -= v2
    v19 := GenS1{F0: 1, F1: v2, F2: 1}
    var v20 string = "foo bar"
    if false {
        v21 := []int{32768}
        switch v20 {
        case "hello":
            _ = len(v21)
        case "":
            if (!false) {
                v22 := GenS2{F0: -819, F1: v20, F2: 324}
                v18 += 1
                var v23 uint64 = uint64(v18)
                _ = v22
                _ = v23
            }
        }
        _ = v21
    } else {
        if (0.1 != ((1e10 / (-2.25 - (0.0 / 3.14159))) / float32(-2.25))) {
            v24 := []int{42, 106, 728, -304}
            for fi25 := 0; fi25 < 4; fi25++ {
                fmt.Println(v2, v18, v18)
                _ = fi25
            }
            var v26 uint32 = v18
            _ = v24
            _ = v26
        }
    }
    fmt.Println(v18, v20, v20)
    if false {
        var v27 int16 = 1000
        for fi28 := 0; fi28 < 6; fi28++ {
            v29 := GenS2{F0: 0, F1: "", F2: v18}
            _ = fi28
            _ = v29
        }
        _ = v27
    } else {
        delete(v4, 0)
        var v30 int = -9223372036854775808
        _ = v30
    }
    v31 := GenS3{F0: 32767, F1: 42}
    switch 563 {
    case -2147483648:
        for fi32 := 0; fi32 < 12; fi32++ {
            if (fi32 > int(v2)) {
                if true {
                    var v33 float64 = 0.1
                    if (0 > 0) {
                        v34 := GenS1{F0: -1, F1: (v2 & uint8(v18)), F2: 42}
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 20; fi35++ {
                            v36 := []int{fi32, fi32}
                            v2 += v2
                            _ = fi35
                            _ = v36
                        }
                    }
                    fmt.Println(v2, v20)
                    _ = v33
                } else {
                    v37 := GenS2{F0: 0, F1: (v20 + v20), F2: (v18 | (^(v18)))}
                    switch fi32 {
                    case 127:
                        var v38 int = (fi32 % ((-9223372036854775808) | 1))
                        _ = v38
                    default:
                        for fi39 := 0; fi39 < 3; fi39++ {
                            switch fi39 {
                            case 9223372036854775807:
                                v40 := GenS3{F0: -32768, F1: 2}
                                _ = v40
                            }
                            if true {
                                if (v20 <= (("\n\t" + "") + v20)) {
                                    if true {
                                        v1 -= (v1 | -75)
                                        switch 32768 {
                                        case -1:
                                            var v41 string = v20
                                            _ = v41
                                        case -672:
                                            switch fi32 {
                                            case 0:
                                                fmt.Println(v2)
                                                var v42 int = (fi32 - fi32)
                                                _ = v42
                                            case 9223372036854775807:
                                                if false {
                                                    delete(v3, (("" + (v20 + v20)) + v20))
                                                    v43 := []int{fi39, fi39, 534, fi39}
                                                    fmt.Println(v1)
                                                    switch (v20 + "hello") {
                                                    case "a":
                                                        switch ((v20 + ("\n\t" + (v20 + ""))) + ((("" + (v20 + v20)) + v20) + ("" + "hello"))) {
                                                        case "foo bar":
                                                            var v44 string = "foo bar"
                                                            _ = v44
                                                        }
                                                    case "foo bar":
                                                        switch fi32 {
                                                        case -1:
                                                            fmt.Println(v1, v18)
                                                        case 4294967295:
                                                            v45 := GenS4{F0: v20}
                                                            _ = v45
                                                        }
                                                    case "":
                                                        if len(v43) > 0 {
                                                            v43[0] = fi32
                                                        }
                                                        if false {
                                                            var v46 uint32 = (v18 / (((v18 + v18)) | 1))
                                                            _ = v46
                                                        }
                                                    }
                                                    var v47 float64 = 0.0
                                                    _ = v43
                                                    _ = v47
                                                }
                                            default:
                                                v48 := GenS4{F0: ((("" + ("" + ("" + v20))) + v20) + "\n\t")}
                                                _ = v48
                                            }
                                        case 42:
                                            var v49 int8 = (v1 ^ (v1 / ((v1) | 1)))
                                            _ = v49
                                        default:
                                            v3[((("" + v20) + v20) + "a")] = (v20 + v20)
                                        }
                                        v1 -= v1
                                    } else {
                                        var v50 uint32 = v18
                                        v18 += 2
                                        var v51 int64 = 380
                                        _ = v50
                                        _ = v51
                                    }
                                } else {
                                    var v52 uint32 = (^(v18))
                                    fmt.Println(v20)
                                    v53 := []int{-356, fi32, fi32, 503}
                                    _ = len(v4)
                                    var v54 uint16 = 506
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                }
                                fmt.Println(fi39, fi39)
                                var v55 uint = 1
                                _ = v55
                            } else {
                                switch (fi32 * 833) {
                                case 4294967295:
                                    var v56 string = "a"
                                    _ = v56
                                case 9223372036854775807:
                                    fmt.Println(v1, v18)
                                case -1:
                                    v57 := []int{(fi39 + (fi39 ^ fi32)), 0, fi32, (fi39 * (fi32 % (((fi39 | (fi39 >> ((107) & 63)))) | 1)))}
                                    _ = v57
                                }
                                var v58 float32 = 1.5
                                var v59 float32 = (((v58 / ((0.0 - float32((-2.25 / (-2.25 / 1.0)))) + (((1.5 * -2.25) + v58) + (0.1 / 1.5)))) / 1.5) + v58)
                                _ = v58
                                _ = v59
                            }
                            _ = fi39
                        }
                    }
                    fmt.Println(v1, v2, v18)
                    _ = v37
                }
            } else {
                v2 -= (v2 << ((65535) & 63))
            }
            switch ((("\n\t" + "") + (((v20 + (v20 + v20)) + (((v20 + (v20 + v20)) + (v20 + "foo bar")) + "")) + v20)) + v20) {
            case "a":
                fmt.Println(fi32)
            case "\n\t":
                for fi60 := 0; fi60 < 7; fi60++ {
                    v61 := v4[32768]
                    _ = fi60
                    _ = v61
                }
            default:
                for fi62 := 0; fi62 < 9; fi62++ {
                    v18 -= v18
                    v63 := GenS5{F0: v1, F1: (v20 + "\n\t"), F2: 18446744073709551615}
                    _ = fi62
                    _ = v63
                }
            }
            if true {
                if (v1 >= v1) {
                    if true {
                        if ((65535 <= 107) || true) {
                            if (true || false) {
                                switch v20 {
                                case "":
                                    for fi64 := 0; fi64 < 5; fi64++ {
                                        fmt.Println(v18)
                                        var v65 int16 = -1
                                        if ((true || true) || false) {
                                            v66 := GenS1{F0: 0, F1: 1, F2: 1}
                                            fmt.Println(v18)
                                            v67 := GenS3{F0: (v65 + v65), F1: 312}
                                            _ = v66
                                            _ = v67
                                        } else {
                                            v65 -= 1
                                            var v68 string = ("a" + v20)
                                            v69 := GenS6{F0: v18, F1: v20, F2: 65535}
                                            _ = v68
                                            _ = v69
                                        }
                                        _ = fi64
                                        _ = v65
                                    }
                                case "\n\t":
                                    var v70 int64 = 41
                                    _ = v70
                                }
                                if true {
                                    fmt.Println(v2, v1, v1)
                                    v2 -= v2
                                    if true {
                                        for fi71 := 0; fi71 < 10; fi71++ {
                                            var v72 uint16 = 32767
                                            v20 = ("a" + v20)
                                            _ = fi71
                                            _ = v72
                                        }
                                        for fi73 := 0; fi73 < 13; fi73++ {
                                            if true {
                                                var v74 string = v20
                                                _ = v74
                                            }
                                            switch (-(fi32)) {
                                            case 0:
                                                v2 += (v2 >> ((65535) & 63))
                                            case 9223372036854775807:
                                                v75 := GenS7{F0: 2, F1: true}
                                                _ = v75
                                            case 4294967295:
                                                v76 := GenS8{F0: (1 <= (v1 >> ((100) & 63)))}
                                                _ = v76
                                            }
                                            if false {
                                                switch (v20 + v20) {
                                                case "\n\t":
                                                    v77 := GenS5{F0: -128, F1: v20, F2: 4294967295}
                                                    if false {
                                                        var v78 uint16 = 355
                                                        v79 := GenS5{F0: (^(v1)), F1: ("hello" + v20), F2: 949}
                                                        var v80 bool = true
                                                        if (!v80) {
                                                            v81 := []int{(-(fi32)), fi32, 0}
                                                            switch fi32 {
                                                            case 2147483647:
                                                                var v82 int16 = 1
                                                                _ = v82
                                                            case -1:
                                                                v83 := GenS4{F0: ((v20 + v20) + (v20 + v20))}
                                                                _ = v83
                                                            case 1:
                                                                var v84 int32 = -46
                                                                _ = v84
                                                            }
                                                            _ = v81
                                                        }
                                                        var v85 int8 = (v1 - v1)
                                                        var v86 int16 = 42
                                                        if ((fi32 / ((fi32) | 1)) >= fi32) {
                                                            var v87 string = v20
                                                            v78 = (v78 >> ((1) & 63))
                                                            if (42 != 0) {
                                                                v88 := []int{fi32, fi73, fi73}
                                                                if false {
                                                                    if true {
                                                                        if true {
                                                                            if len(v88) >= 2 {
                                                                                v88 = v88[1:len(v88)]
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v2)
                                                                            var v89 float64 = ((1.5 / float64(3.14159)) * 1e10)
                                                                            _ = v89
                                                                        }
                                                                    }
                                                                    fmt.Println(v2)
                                                                } else {
                                                                    if len(v88) >= 2 {
                                                                        v88 = v88[1:len(v88)]
                                                                    }
                                                                    v88 = append(v88, (-(fi73)))
                                                                    v87 = ((v87 + v20) + "")
                                                                    switch fi73 {
                                                                    case 0:
                                                                        fmt.Println(fi73, v87)
                                                                    default:
                                                                        v4[(fi73 % (((fi73 / ((fi73) | 1))) | 1))] = (fi73 - 41)
                                                                    }
                                                                    switch fi32 {
                                                                    case -542:
                                                                        fmt.Println(fi32)
                                                                    default:
                                                                        var v90 int64 = 0
                                                                        _ = v90
                                                                    }
                                                                    switch (fi73 / ((-995) | 1)) {
                                                                    case 32767:
                                                                        v20 = (v20 + (v20 + v20))
                                                                    default:
                                                                        var v91 int32 = 1
                                                                        _ = v91
                                                                    }
                                                                }
                                                                v3[(("" + v87) + v20)] = v87
                                                                if len(v88) >= 2 {
                                                                    v88 = v88[1:len(v88)]
                                                                }
                                                                _ = v88
                                                            } else {
                                                                var v92 int = (fi32 ^ 4294967295)
                                                                if (v2 == 0) {
                                                                    var v93 uint64 = 42
                                                                    _ = v93
                                                                } else {
                                                                    v78 -= 814
                                                                    fmt.Println(v78, v80)
                                                                    var v94 uint8 = (^(v2))
                                                                    _ = v94
                                                                }
                                                                var v95 int32 = -2147483648
                                                                _ = v92
                                                                _ = v95
                                                            }
                                                            v80 = (9223372036854775807 < 42)
                                                            if (1 < 0) {
                                                                v96 := []int{}
                                                                if v80 {
                                                                    var v97 float32 = 1e10
                                                                    fmt.Println(v86)
                                                                    _ = v97
                                                                }
                                                                _ = v96
                                                            } else {
                                                                v18 -= (^(v18))
                                                                v98 := GenS8{F0: v80}
                                                                _ = v98
                                                            }
                                                            _ = v87
                                                        }
                                                        v78 = (v78 ^ (v78 - v78))
                                                        v99 := GenS1{F0: 1, F1: 47, F2: 9223372036854775807}
                                                        _ = v78
                                                        _ = v79
                                                        _ = v80
                                                        _ = v85
                                                        _ = v86
                                                        _ = v99
                                                    }
                                                    v2 = (v2 & v2)
                                                    var v100 int8 = int8(v18)
                                                    var v101 int32 = 350
                                                    var v102 int8 = (v1 + (v1 >> ((4294967295) & 63)))
                                                    _ = v100
                                                    _ = v101
                                                    _ = v102
                                                    _ = v77
                                                }
                                            } else {
                                                v103 := GenS7{F0: v2, F1: (42 < 899)}
                                                fmt.Println(v1)
                                                fmt.Println(v1, fi73, v18)
                                                _ = v103
                                            }
                                            _ = fi73
                                        }
                                    } else {
                                        switch fi32 {
                                        case 127:
                                            v104 := GenS8{F0: false}
                                            _ = v104
                                        case 4294967295:
                                            v3[v20] = "\n\t"
                                        case -529:
                                            v105 := v4[-147]
                                            _ = v105
                                        default:
                                            v106 := []int{fi32, (fi32 - fi32), fi32, fi32}
                                            _ = v106
                                        }
                                        v107 := GenS9{F0: (v20 + ((v20 + "a") + v20))}
                                        v108 := GenS10{F0: 42, F1: (^(v2))}
                                        _ = v107
                                        _ = v108
                                    }
                                }
                                if false {
                                    var v109 int64 = 568
                                    var v110 float64 = (1e10 / -2.25)
                                    v18 = 65535
                                    _ = v109
                                    _ = v110
                                } else {
                                    var v111 float64 = -2.25
                                    v112 := GenS6{F0: 4294967295, F1: "foo bar", F2: 32767}
                                    var v113 uint8 = (v2 | (v2 / ((v2) | 1)))
                                    _ = v111
                                    _ = v112
                                    _ = v113
                                }
                            }
                            for fi114 := 0; fi114 < 3; fi114++ {
                                switch v20 {
                                case "a":
                                    v18 += v18
                                }
                                _ = fi114
                            }
                            v20 = v20
                        } else {
                            v18 -= 2147483648
                            for fi115 := 0; fi115 < 19; fi115++ {
                                if (1 <= 18446744073709551615) {
                                    var v116 int = (fi115 * fi32)
                                    v117 := GenS11{F0: 0, F1: 0, F2: 0}
                                    v118 := GenS2{F0: -1, F1: "foo bar", F2: (v18 + v18)}
                                    switch (fi115 - fi115) {
                                    case -812:
                                        fmt.Println(v20)
                                    case -217:
                                        if (-9223372036854775808 < 42) {
                                            if (true && true) {
                                                v2 += (v2 & (v2 * (^(v2))))
                                                v1 += (v1 ^ (v1 + 123))
                                                var v119 int32 = 551
                                                v120 := []int{v116, fi32, fi115}
                                                if (1.5 <= (0.1 - 1e10)) {
                                                    switch 127 {
                                                    case 1:
                                                        if (42 <= 255) {
                                                            switch (^(v116)) {
                                                            case 127:
                                                                v121 := GenS12{F0: 42}
                                                                _ = v121
                                                            case 32768:
                                                                if false {
                                                                    var v122 uint8 = 57
                                                                    var v123 int64 = 752
                                                                    v124 := GenS13{F0: (v18 >> ((255) & 63)), F1: 255}
                                                                    _ = v122
                                                                    _ = v123
                                                                    _ = v124
                                                                }
                                                            case 32767:
                                                                v125 := GenS12{F0: 0}
                                                                v20 = v20
                                                                _ = v125
                                                            }
                                                            v126 := v4[1]
                                                            var v127 uint32 = 2
                                                            _ = v126
                                                            _ = v127
                                                        } else {
                                                            if false {
                                                                switch (v20 + "\n\t") {
                                                                case "":
                                                                    v18 = (v18 >> ((424) & 63))
                                                                }
                                                                var v128 string = v20
                                                                _ = v128
                                                            }
                                                            v129 := GenS14{F0: v116, F1: 0}
                                                            v20 = "a"
                                                            _ = v129
                                                        }
                                                    case -9223372036854775808:
                                                        var v130 string = (v20 + "hello")
                                                        _ = v130
                                                    }
                                                }
                                                _ = v119
                                                _ = v120
                                            }
                                            v1 -= int8(v2)
                                            var v131 uint32 = v18
                                            _ = v131
                                        } else {
                                            if true {
                                                if true {
                                                    v132 := GenS12{F0: -798}
                                                    _ = v132
                                                }
                                            } else {
                                                if false {
                                                    switch v20 {
                                                    case "":
                                                        v18 = v18
                                                    case "\n\t":
                                                        var v133 int16 = -32768
                                                        _ = v133
                                                    default:
                                                        if (int16(v18) == 1) {
                                                            switch (v116 >> ((965) & 63)) {
                                                            case 2147483647:
                                                                v134 := []int{(v116 & 0), fi32, 4294967295, fi32}
                                                                _ = v134
                                                            default:
                                                                var v135 int32 = 2147483647
                                                                v136 := GenS11{F0: 1, F1: 70, F2: (v2 / ((v2) | 1))}
                                                                switch ("a" + v20) {
                                                                case "a":
                                                                    fmt.Println(fi115)
                                                                default:
                                                                    v137 := GenS11{F0: (-(v135)), F1: 127, F2: v2}
                                                                    v138 := []int{-9223372036854775808, (fi115 - (fi32 * 42))}
                                                                    var v139 uint32 = 4294967295
                                                                    v140 := GenS13{F0: v18, F1: 42}
                                                                    _ = v137
                                                                    _ = v138
                                                                    _ = v139
                                                                    _ = v140
                                                                }
                                                                _ = v135
                                                                _ = v136
                                                            }
                                                            v141 := GenS1{F0: -1, F1: uint8(v18), F2: -190}
                                                            var v142 int16 = -890
                                                            _ = v141
                                                            _ = v142
                                                        } else {
                                                            var v143 int16 = 32767
                                                            _ = len(v3)
                                                            _ = v143
                                                        }
                                                    }
                                                    fmt.Println(v2, v1, v20)
                                                    v1 = 42
                                                }
                                                if (true && true) {
                                                    v144 := []int{(fi32 / (((v116 >> ((18446744073709551615) & 63))) | 1))}
                                                    var v145 uint = 225
                                                    if true {
                                                        _ = len(v4)
                                                        var v146 uint16 = 42
                                                        switch fi32 {
                                                        case 0:
                                                            if false {
                                                                if (false || false) {
                                                                    v147 := GenS15{F0: v2, F1: v20, F2: fi115}
                                                                    switch 127 {
                                                                    case 32767:
                                                                        v145 = (v145 & 0)
                                                                    default:
                                                                        var v148 int = -1
                                                                        _ = v148
                                                                    }
                                                                    v149 := GenS16{F0: v145, F1: v146}
                                                                    _ = v147
                                                                    _ = v149
                                                                } else {
                                                                    v145 = (v145 / ((v145) | 1))
                                                                    var v150 string = "\n\t"
                                                                    _ = v150
                                                                }
                                                                var v151 uint8 = (v2 ^ 1)
                                                                switch ((v20 + v20) + v20) {
                                                                case "hello":
                                                                    var v152 float32 = float32(0.0)
                                                                    _ = v152
                                                                case "":
                                                                    if len(v144) > 0 {
                                                                        v144[0] = 0
                                                                    }
                                                                case "foo bar":
                                                                    switch (v20 + v20) {
                                                                    case "foo bar":
                                                                        v145 = v145
                                                                    case "hello":
                                                                        if true {
                                                                            _ = len(v3)
                                                                            v153 := v4[(fi32 ^ v116)]
                                                                            v154 := GenS17{F0: 32767, F1: v2, F2: v18}
                                                                            _ = v153
                                                                            _ = v154
                                                                        } else {
                                                                            delete(v3, (v20 + v20))
                                                                            v155 := GenS6{F0: (v18 / (((v18 / (((v18 & 42)) | 1))) | 1)), F1: v20, F2: 1}
                                                                            _ = v155
                                                                        }
                                                                    case "a":
                                                                        v156 := GenS7{F0: v2, F1: false}
                                                                        _ = v156
                                                                    default:
                                                                        if (42 <= (fi115 << (((v145 * (v145 + (v145 >> ((v145) & 63))))) & 63))) {
                                                                            v116 += fi115
                                                                        } else {
                                                                            v157 := GenS18{F0: v151}
                                                                            _ = v157
                                                                        }
                                                                    }
                                                                }
                                                                v158 := GenS10{F0: -347, F1: v2}
                                                                _ = v151
                                                                _ = v158
                                                            }
                                                        case 32768:
                                                            var v159 uint8 = (v2 & (^(v2)))
                                                            _ = v159
                                                        default:
                                                            var v160 int = int(v2)
                                                            _ = v160
                                                        }
                                                        _ = v146
                                                    } else {
                                                        fmt.Println(v145)
                                                    }
                                                    _ = v144
                                                    _ = v145
                                                } else {
                                                    var v161 int32 = -18
                                                    fmt.Println(v161, v20, v1)
                                                    fmt.Println(v1, v161, v20)
                                                    _ = v161
                                                }
                                                v1 -= (v1 / ((v1) | 1))
                                            }
                                            var v162 uint32 = (^(v18))
                                            _ = v162
                                        }
                                    case 198:
                                        fmt.Println(v2, fi115, v2)
                                    default:
                                        var v163 string = ("" + v20)
                                        _ = v163
                                    }
                                    _ = v116
                                    _ = v117
                                    _ = v118
                                } else {
                                    switch v20 {
                                    case "hello":
                                        v1 = (v1 * (v1 & (v1 & 104)))
                                    case "a":
                                        v18 = v18
                                    }
                                }
                                var v164 int64 = -1
                                v2 = 2
                                _ = fi115
                                _ = v164
                            }
                        }
                    }
                    var v165 bool = (!true)
                    for fi166 := 0; fi166 < 20; fi166++ {
                        v167 := GenS19{F0: uint16(v2), F1: v18}
                        _ = fi166
                        _ = v167
                    }
                    _ = v165
                } else {
                    v168 := GenS5{F0: (^(v1)), F1: v20, F2: 432}
                    fmt.Println(v18, v1)
                    v20 = ""
                    _ = v168
                }
                switch "" {
                case "\n\t":
                    switch ((((v20 + (v20 + v20)) + "") + (v20 + "a")) + v20) {
                    case "":
                        if true {
                            var v169 bool = false
                            v170 := GenS20{F0: uint(v18)}
                            v169 = (v1 >= int8(v18))
                            _ = v169
                            _ = v170
                        } else {
                            for fi171 := 0; fi171 < 13; fi171++ {
                                if true {
                                    switch fi171 {
                                    case 949:
                                        v172 := GenS8{F0: true}
                                        _ = v172
                                    case 254:
                                        switch (fi171 | (fi171 ^ fi32)) {
                                        case -1:
                                            var v173 int32 = 42
                                            _ = v173
                                        default:
                                            var v174 int16 = -1
                                            _ = v174
                                        }
                                    default:
                                        v2 -= v2
                                    }
                                    switch v20 {
                                    case "":
                                        var v175 float64 = (3.14159 + 1e10)
                                        var v176 uint16 = 1
                                        _ = len(v4)
                                        var v177 float32 = 0.0
                                        _ = v175
                                        _ = v176
                                        _ = v177
                                    case "a":
                                        var v178 bool = true
                                        _ = v178
                                    }
                                }
                                _ = fi171
                            }
                            var v179 uint8 = 127
                            _ = v179
                        }
                    }
                default:
                    for fi180 := 0; fi180 < 9; fi180++ {
                        switch v20 {
                        case "hello":
                            v20 = "\n\t"
                        default:
                            v1 = (v1 & v1)
                        }
                        _ = fi180
                    }
                }
                for fi181 := 0; fi181 < 12; fi181++ {
                    var v182 bool = false
                    if (true && (539 <= 2)) {
                        var v183 uint64 = 72
                        v184 := []int{(fi32 % ((fi32) | 1)), fi181, fi32}
                        v18 = (^(v18))
                        _ = v183
                        _ = v184
                    }
                    v1 = -17
                    _ = fi181
                    _ = v182
                }
            } else {
                var v185 int8 = (v1 | (v1 - (v1 >> ((uint(v18)) & 63))))
                v18 += (^(v18))
                v186 := GenS18{F0: (v2 | v2)}
                _ = v185
                _ = v186
            }
            _ = fi32
        }
    case 32767:
        for fi187 := 0; fi187 < 7; fi187++ {
            v20 = v20
            var v188 int32 = 1
            v189 := GenS21{F0: v2, F1: 576}
            _ = fi187
            _ = v188
            _ = v189
        }
    case -1:
        v18 -= 68
    }
    for fi190 := 0; fi190 < 17; fi190++ {
        switch fi190 {
        case -259:
            for fi191 := 0; fi191 < 13; fi191++ {
                var v192 int32 = 2147483647
                _ = fi191
                _ = v192
            }
        case -831:
            v18 -= v18
        default:
            v193 := []int{fi190, (fi190 & (fi190 ^ (fi190 << ((507) & 63))))}
            _ = v193
        }
        _ = fi190
    }
    var v194 bool = (135 > 1000)
    delete(v4, -925)
    for fi195 := 0; fi195 < 13; fi195++ {
        for fi196 := 0; fi196 < 17; fi196++ {
            var v197 uint64 = 1
            var v198 float64 = 0.0
            v197 -= v197
            _ = fi196
            _ = v197
            _ = v198
        }
        v199 := GenS22{F0: 344, F1: v194}
        _ = fi195
        _ = v199
    }
    v18 = v18
    switch v20 {
    case "hello":
        v200 := []int{226, -1}
        _ = v200
    case "a":
        fmt.Println(v18, v20)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v194:", v194)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v194
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v4
}
