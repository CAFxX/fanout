package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int64; F1 uint; F2 string }
type GenS6 struct { F0 int8; F1 int64; F2 uint16 }
type GenS7 struct { F0 bool; F1 int8; F2 uint }
type GenS8 struct { F0 int32; F1 int64; F2 int8 }
type GenS9 struct { F0 uint16; F1 uint64 }
type GenS10 struct { F0 int8; F1 int; F2 int }
type GenS11 struct { F0 int32; F1 int8; F2 string }
type GenS12 struct { F0 int }
type GenS13 struct { F0 bool; F1 int16 }
type GenS14 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS15 struct { F0 int16; F1 uint16 }
type GenS16 struct { F0 int16; F1 string }
type GenS17 struct { F0 uint64; F1 int }
type GenS18 struct { F0 int64; F1 int16 }
type GenS19 struct { F0 uint; F1 int }
type GenS20 struct { F0 bool; F1 int8; F2 bool }
type GenS21 struct { F0 int64; F1 int16; F2 int64 }
type GenS22 struct { F0 uint64; F1 int16; F2 int8 }
type GenS23 struct { F0 uint; F1 int }
type GenS24 struct { F0 int; F1 uint64 }
type GenS25 struct { F0 string }
type GenS26 struct { F0 int8 }
type GenS27 struct { F0 int8 }
type GenS28 struct { F0 int64 }
type GenS29 struct { F0 int16; F1 uint64 }
type GenS30 struct { F0 uint8; F1 bool; F2 uint8 }

func main() {
    v1 := make(map[string]string)
    v2 := []int{4294967295, 520}
    var v3 uint16 = 42
    v4 := make(map[uint32]uint32)
    var v5 int32 = 827
    v6 := make(map[int]string)
    var v7 int16 = int16(v5)
    var v8 uint32 = 497
    for fi9 := 0; fi9 < 5; fi9++ {
        switch "\n\t" {
        case "foo bar":
            if ((-59 < 1) && true) {
                v7 = v7
                v7 += v7
            } else {
                for fi10 := 0; fi10 < 11; fi10++ {
                    v11 := GenS1{F0: 4294967295, F1: -128, F2: (-(fi10))}
                    v7 += 0
                    var v12 int64 = 0
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = len(v2)
                v3 += (v3 + v3)
            }
        case "hello":
            for fi13 := 0; fi13 < 17; fi13++ {
                v14 := GenS2{F0: 253}
                _ = len(v2)
                var v15 string = ((("a" + (("" + "foo bar") + "a")) + "hello") + "hello")
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        fmt.Println(fi9, v5)
        _ = fi9
    }
    _ = len(v2)
    var v16 uint8 = 2
    v17 := GenS1{F0: 231, F1: -79, F2: 709}
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v18 float64 = (((3.14159 - (3.14159 + float64(3.14159))) + 1.5) * ((0.0 * 0.1) * 1.5))
        _ = v18
    } else {
        v3 += v3
        if false {
            if false {
                v7 = -200
                var v19 float64 = 3.14159
                _ = v19
            }
            if true {
                switch "\n\t" {
                case "hello":
                    for fi20 := 0; fi20 < 8; fi20++ {
                        var v21 int16 = (v7 * (^(v7)))
                        for fi22 := 0; fi22 < 2; fi22++ {
                            var v23 uint8 = 127
                            var v24 float64 = 3.14159
                            v7 = 32767
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        _ = fi20
                        _ = v21
                    }
                case "\n\t":
                    var v25 float64 = 1.5
                    _ = v25
                }
                if false {
                    switch -727 {
                    case 9223372036854775807:
                        if false {
                            v7 -= 42
                            if false {
                                v26 := GenS3{F0: ""}
                                _ = len(v1)
                                for fi27 := 0; fi27 < 5; fi27++ {
                                    for fi28 := 0; fi28 < 19; fi28++ {
                                        fmt.Println(v3)
                                        v8 = 42
                                        _ = fi28
                                    }
                                    for fi29 := 0; fi29 < 15; fi29++ {
                                        var v30 uint8 = v16
                                        _ = fi29
                                        _ = v30
                                    }
                                    for fi31 := 0; fi31 < 16; fi31++ {
                                        v32 := GenS4{F0: (v16 % ((v16) | 1)), F1: 42}
                                        fmt.Println(fi27)
                                        var v33 int32 = (v5 >> ((346) & 63))
                                        fmt.Println(v5, v3)
                                        switch fi31 {
                                        case 4294967295:
                                            var v34 uint16 = (^(v3))
                                            _ = v34
                                        }
                                        v35 := GenS5{F0: 1, F1: 2, F2: "hello"}
                                        var v36 int64 = -9223372036854775808
                                        _ = fi31
                                        _ = v32
                                        _ = v33
                                        _ = v35
                                        _ = v36
                                    }
                                    _ = fi27
                                }
                                _ = v26
                            }
                            v37 := GenS6{F0: 0, F1: 223, F2: 1}
                            _ = v37
                        }
                    case 128:
                        if false {
                            if ((v5 >> ((95) & 63)) > 0) {
                                var v38 int16 = v7
                                if true {
                                    switch 191 {
                                    case 127:
                                        if true {
                                            for fi39 := 0; fi39 < 13; fi39++ {
                                                var v40 int16 = v38
                                                if true {
                                                    var v41 uint8 = 1
                                                    var v42 float32 = 1.5
                                                    v43 := GenS3{F0: (("\n\t" + "foo bar") + ("hello" + ""))}
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                } else {
                                                    var v44 int16 = (v7 / ((-1) | 1))
                                                    _ = v44
                                                }
                                                v3 += 0
                                                _ = fi39
                                                _ = v40
                                            }
                                            for fi45 := 0; fi45 < 18; fi45++ {
                                                _ = len(v6)
                                                var v46 int64 = -9223372036854775808
                                                _ = fi45
                                                _ = v46
                                            }
                                        }
                                    case 956:
                                        v47 := GenS7{F0: false, F1: 127, F2: 1}
                                        _ = v47
                                    }
                                } else {
                                    v48 := v1["foo bar"]
                                    for fi49 := 0; fi49 < 1; fi49++ {
                                        var v50 uint64 = uint64(v8)
                                        _ = fi49
                                        _ = v50
                                    }
                                    v51 := GenS3{F0: ("" + v48)}
                                    _ = v48
                                    _ = v51
                                }
                                _ = v38
                            }
                            for fi52 := 0; fi52 < 9; fi52++ {
                                for fi53 := 0; fi53 < 9; fi53++ {
                                    v54 := GenS8{F0: v5, F1: -816, F2: 0}
                                    if false {
                                        var v55 uint8 = uint8(v8)
                                        _ = v55
                                    }
                                    _ = fi53
                                    _ = v54
                                }
                                if false {
                                    var v56 string = "\n\t"
                                    _ = v56
                                } else {
                                    v8 += 0
                                    fmt.Println(v8, v8)
                                    var v57 int32 = v5
                                    _ = v57
                                }
                                for fi58 := 0; fi58 < 15; fi58++ {
                                    fmt.Println(v5)
                                    _ = fi58
                                }
                                _ = fi52
                            }
                            for fi59 := 0; fi59 < 15; fi59++ {
                                fmt.Println(v8, v5, v7)
                                var v60 uint16 = v3
                                if ((2 < 65535) || (!true)) {
                                    v3 -= 1
                                    _ = len(v1)
                                }
                                _ = fi59
                                _ = v60
                            }
                        } else {
                            v3 = v3
                            if false {
                                var v61 int8 = -128
                                v3 = 933
                                _ = v61
                            } else {
                                _ = len(v4)
                                v62 := GenS7{F0: (("a" + "hello") == "hello"), F1: 6, F2: 255}
                                if ((!true) && ((true || true) && (!((v8 >= v8) && (true || ((995 != -2147483648) || (v5 != 823))))))) {
                                    v63 := GenS9{F0: (v3 << ((1) & 63)), F1: 343}
                                    _ = v63
                                } else {
                                    var v64 int8 = 1
                                    v64 += (v64 * 81)
                                    var v65 int8 = (v64 * (^(v64)))
                                    _ = v64
                                    _ = v65
                                }
                                _ = v62
                            }
                        }
                    case 127:
                        if true {
                            v66 := GenS5{F0: 9223372036854775807, F1: 42, F2: ""}
                            if false {
                                for fi67 := 0; fi67 < 9; fi67++ {
                                    v68 := GenS10{F0: -1, F1: fi67, F2: fi67}
                                    _ = fi67
                                    _ = v68
                                }
                                for fi69 := 0; fi69 < 5; fi69++ {
                                    var v70 uint64 = 441
                                    v71 := GenS11{F0: -2147483648, F1: 0, F2: "hello"}
                                    switch fi69 {
                                    case -2147483648:
                                        var v72 string = "foo bar"
                                        _ = v72
                                    case 584:
                                        v70 = v70
                                    default:
                                        _ = len(v4)
                                    }
                                    _ = fi69
                                    _ = v70
                                    _ = v71
                                }
                            }
                            for fi73 := 0; fi73 < 12; fi73++ {
                                if true {
                                    v74 := GenS12{F0: -501}
                                    var v75 string = ""
                                    var v76 float32 = (0.0 * float32(1.5))
                                    _ = v74
                                    _ = v75
                                    _ = v76
                                } else {
                                    v3 = 848
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        for fi77 := 0; fi77 < 12; fi77++ {
                                            v78 := GenS1{F0: 578, F1: 31, F2: 32767}
                                            if false {
                                                var v79 int32 = v5
                                                v80 := GenS13{F0: (1 == 4294967295), F1: v7}
                                                _ = v79
                                                _ = v80
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = (fi77 * (^(fi73)))
                                            }
                                            _ = fi77
                                            _ = v78
                                        }
                                        var v81 int64 = 1
                                        _ = v81
                                    } else {
                                        for fi82 := 0; fi82 < 10; fi82++ {
                                            if (3.14159 > 3.14159) {
                                                delete(v4, v8)
                                                v5 += v5
                                            } else {
                                                fmt.Println(fi73, v8, fi73)
                                                fmt.Println(v7, v8, v8)
                                            }
                                            switch -9223372036854775808 {
                                            case 32768:
                                                var v83 uint8 = (^(v16))
                                                _ = v83
                                            case 0:
                                                v8 = 342
                                            case 4294967295:
                                                switch (fi82 - fi82) {
                                                case 4294967295:
                                                    if ((-25 <= 1) || true) {
                                                        if (3.14159 > -2.25) {
                                                            if (!false) {
                                                                if (311 >= (v8 & (^(v8)))) {
                                                                    fmt.Println(v16, v8)
                                                                    v7 = 466
                                                                    if (!(int64(v8) <= 9223372036854775807)) {
                                                                        fmt.Println(v16, v7, fi82)
                                                                        if (!true) {
                                                                            var v84 int8 = 100
                                                                            var v85 string = ""
                                                                            switch v85 {
                                                                            case "hello":
                                                                                v86 := GenS9{F0: v3, F1: 42}
                                                                                _ = v86
                                                                            case "\n\t":
                                                                                var v87 string = ((v85 + (v85 + (v85 + "\n\t"))) + (v85 + "hello"))
                                                                                _ = v87
                                                                            }
                                                                            _ = v84
                                                                            _ = v85
                                                                        }
                                                                        var v88 int8 = -1
                                                                        _ = v88
                                                                    } else {
                                                                        v16 += v16
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        if true {
                                                                            var v89 bool = true
                                                                            v90 := GenS14{F0: v89, F1: v16, F2: 1}
                                                                            _ = v89
                                                                            _ = v90
                                                                        } else {
                                                                            if true {
                                                                                v16 -= 127
                                                                            } else {
                                                                                v7 -= 1000
                                                                                var v91 uint = 255
                                                                                _ = v91
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v3 += v3
                                                                        v92 := v1[("\n\t" + ("a" + "a"))]
                                                                        switch ("\n\t" + v92) {
                                                                        case "\n\t":
                                                                            if (false && true) {
                                                                                var v93 int32 = v5
                                                                                if true {
                                                                                    delete(v1, v92)
                                                                                    v7 -= v7
                                                                                }
                                                                                v2 = append(v2, fi73)
                                                                                if true {
                                                                                    switch v92 {
                                                                                    case "":
                                                                                        var v94 int16 = int16(v16)
                                                                                        _ = v94
                                                                                    case "a":
                                                                                        var v95 uint16 = (v3 * 32767)
                                                                                        _ = v95
                                                                                    }
                                                                                    var v96 uint64 = 4294967295
                                                                                    v5 += v5
                                                                                    _ = v96
                                                                                }
                                                                                switch (fi73 ^ (-(fi73))) {
                                                                                case -840:
                                                                                    v4[v8] = (v8 >> ((42) & 63))
                                                                                case 9223372036854775807:
                                                                                    fmt.Println(v7, v16, v3)
                                                                                default:
                                                                                    var v97 int32 = (^(v5))
                                                                                    _ = v97
                                                                                }
                                                                                _ = v93
                                                                            } else {
                                                                                var v98 int16 = 847
                                                                                fmt.Println(fi82, v8, v5)
                                                                                _ = v98
                                                                            }
                                                                        case "":
                                                                            v16 = 127
                                                                        default:
                                                                            var v99 uint = 838
                                                                            _ = v99
                                                                        }
                                                                        _ = v92
                                                                    }
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    v100 := GenS11{F0: v5, F1: 1, F2: "\n\t"}
                                                                    _ = v100
                                                                }
                                                                delete(v6, (fi82 | (fi82 | (fi73 - (fi82 / ((int(v3)) | 1))))))
                                                                var v101 bool = ((!(fi82 == fi82)) || false)
                                                                var v102 int8 = 42
                                                                var v103 bool = true
                                                                _ = v101
                                                                _ = v102
                                                                _ = v103
                                                            } else {
                                                                v7 = (v7 ^ 42)
                                                                v3 = (v3 | 1)
                                                                v3 -= (v3 - uint16(v16))
                                                            }
                                                            if true {
                                                                fmt.Println(v16)
                                                                switch int(v7) {
                                                                case -9223372036854775808:
                                                                    var v104 uint32 = v8
                                                                    _ = v104
                                                                case 128:
                                                                    v105 := GenS1{F0: 19, F1: 1, F2: (fi73 >> ((uint(v16)) & 63))}
                                                                    _ = v105
                                                                }
                                                                v106 := GenS5{F0: 186, F1: 330, F2: ("hello" + ((((("foo bar" + "") + "hello") + "") + "") + ((("" + ("" + "a")) + "\n\t") + "a")))}
                                                                _ = v106
                                                            } else {
                                                                _ = len(v1)
                                                                var v107 uint = 4294967295
                                                                _ = len(v4)
                                                                _ = v107
                                                            }
                                                        } else {
                                                            v3 += v3
                                                            v16 += (v16 & v16)
                                                            var v108 string = "\n\t"
                                                            _ = v108
                                                        }
                                                    } else {
                                                        var v109 int32 = (-(v5))
                                                        fmt.Println(v8, fi73, v5)
                                                        _ = v109
                                                    }
                                                default:
                                                    v110 := GenS4{F0: v16, F1: v8}
                                                    _ = v110
                                                }
                                            }
                                            if true {
                                                var v111 string = "a"
                                                _ = v111
                                            }
                                            var v112 int = fi73
                                            if false {
                                                fmt.Println(v8, fi73, fi73)
                                            }
                                            _ = fi82
                                            _ = v112
                                        }
                                    }
                                    v3 -= v3
                                }
                                _ = fi73
                            }
                            _ = v66
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                v5 -= -22
                            case "foo bar":
                                var v113 float32 = 3.14159
                                _ = v113
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                    var v114 uint = 3
                    _ = v114
                }
            } else {
                for fi115 := 0; fi115 < 15; fi115++ {
                    var v116 int16 = (v7 * v7)
                    for fi117 := 0; fi117 < 16; fi117++ {
                        if (true && ((fi115 * (fi117 ^ fi117)) != (fi115 ^ fi115))) {
                            v6[fi117] = "hello"
                            v5 = v5
                            v118 := GenS15{F0: 0, F1: 337}
                            _ = v118
                        } else {
                            var v119 string = "foo bar"
                            fmt.Println(v8)
                            _ = v119
                        }
                        _ = fi117
                    }
                    var v120 int16 = -307
                    _ = fi115
                    _ = v116
                    _ = v120
                }
                v121 := GenS16{F0: -706, F1: "a"}
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(v3)
                    switch -510 {
                    case 1:
                        switch 4294967295 {
                        case 0:
                            if false {
                                if false {
                                    var v122 uint = 18446744073709551615
                                    _ = v122
                                } else {
                                    v5 = v5
                                }
                                switch 269 {
                                case 672:
                                    fmt.Println(v3)
                                case 128:
                                    var v123 uint16 = v3
                                    _ = v123
                                case -9223372036854775808:
                                    _ = len(v2)
                                }
                            } else {
                                for fi124 := 0; fi124 < 9; fi124++ {
                                    var v125 uint32 = 4294967295
                                    _ = fi124
                                    _ = v125
                                }
                                for fi126 := 0; fi126 < 14; fi126++ {
                                    v127 := GenS10{F0: -43, F1: int(v7), F2: (fi126 % ((2147483647) | 1))}
                                    _ = len(v2)
                                    var v128 uint = 0
                                    _ = fi126
                                    _ = v127
                                    _ = v128
                                }
                            }
                        case -276:
                            fmt.Println(v3)
                        default:
                            for fi129 := 0; fi129 < 9; fi129++ {
                                switch "" {
                                case "foo bar":
                                    if (0.1 >= (((0.1 + ((((-2.25 + -2.25) * (1e10 - 0.0)) / 1.5) * 1.5)) + -2.25) + (0.1 + 0.1))) {
                                        if (("a" <= "\n\t") && true) {
                                            v16 += 0
                                            v130 := GenS17{F0: 42, F1: 32768}
                                            _ = v130
                                        }
                                    } else {
                                        var v131 uint16 = v3
                                        v8 += v8
                                        if ((fi129 & (fi129 << ((0) & 63))) > -9223372036854775808) {
                                            v5 = -2147483648
                                            var v132 int64 = 42
                                            switch "\n\t" {
                                            case "\n\t":
                                                v133 := GenS7{F0: false, F1: -38, F2: 4294967295}
                                                _ = v133
                                            case "hello":
                                                if false {
                                                    v134 := GenS8{F0: -2147483648, F1: (v132 + (v132 >> ((4294967295) & 63))), F2: 100}
                                                    var v135 int16 = 0
                                                    _ = v134
                                                    _ = v135
                                                } else {
                                                    if false {
                                                        delete(v4, v8)
                                                    } else {
                                                        if true {
                                                            switch fi129 {
                                                            case -352:
                                                                for fi136 := 0; fi136 < 7; fi136++ {
                                                                    v3 = 65535
                                                                    if (true && false) {
                                                                        v7 -= int16(fi129)
                                                                        if false {
                                                                            var v137 uint16 = v3
                                                                            v131 -= 32767
                                                                            var v138 float32 = (0.1 / 1.0)
                                                                            v7 = 1
                                                                            v139 := GenS10{F0: 127, F1: fi129, F2: 685}
                                                                            _ = v137
                                                                            _ = v138
                                                                            _ = v139
                                                                        }
                                                                    } else {
                                                                        v8 -= 65535
                                                                        if false {
                                                                            v131 = v3
                                                                            switch (fi129 | int(v5)) {
                                                                            case 42:
                                                                                var v140 int32 = 0
                                                                                _ = v140
                                                                            case 32767:
                                                                                if true {
                                                                                    var v141 bool = (!true)
                                                                                    fmt.Println(v132, v3, v131)
                                                                                    if (9223372036854775807 > fi129) {
                                                                                        _ = len(v6)
                                                                                        v132 = v132
                                                                                    }
                                                                                    v142 := GenS18{F0: (-(v132)), F1: v7}
                                                                                    var v143 uint16 = (v131 + 42)
                                                                                    _ = v141
                                                                                    _ = v142
                                                                                    _ = v143
                                                                                }
                                                                                var v144 uint64 = 812
                                                                                fmt.Println(fi136)
                                                                                _ = v144
                                                                            }
                                                                            var v145 bool = ((v3 <= v131) || (255 != 255))
                                                                            v5 -= (v5 & 42)
                                                                            _ = v145
                                                                        }
                                                                    }
                                                                    var v146 uint16 = v3
                                                                    _ = fi136
                                                                    _ = v146
                                                                }
                                                            case 32768:
                                                                v147 := GenS14{F0: false, F1: (v16 % ((v16) | 1)), F2: 18446744073709551615}
                                                                _ = v147
                                                            default:
                                                                v131 += 32767
                                                            }
                                                            v148 := v6[fi129]
                                                            _ = v148
                                                        }
                                                        var v149 uint32 = v8
                                                        if true {
                                                            var v150 float32 = float32(((((0.1 * 0.0) * (((0.1 / -2.25) - float64(float32(1.5))) * (-2.25 * -2.25))) * 0.1) + float64(3.14159)))
                                                            switch -9223372036854775808 {
                                                            case -241:
                                                                var v151 uint = 255
                                                                _ = v151
                                                            case -331:
                                                                fmt.Println(v16, v16, v132)
                                                            case -1:
                                                                var v152 float32 = v150
                                                                _ = v152
                                                            default:
                                                                v150 = ((v150 * (3.14159 + -2.25)) + v150)
                                                            }
                                                            _ = v150
                                                        } else {
                                                            var v153 uint32 = v149
                                                            v153 += v153
                                                            _ = v153
                                                        }
                                                        _ = v149
                                                    }
                                                    var v154 float32 = ((1e10 / (((3.14159 / (float32(3.14159) * float32(float64(((1e10 / 1e10) * 1.5))))) * float32(0.1)) * (((-2.25 / -2.25) + (1e10 * -2.25)) - 3.14159))) + (3.14159 + float32(-2.25)))
                                                    _ = v154
                                                }
                                            case "":
                                                fmt.Println(v8, v5)
                                            }
                                            _ = v132
                                        } else {
                                            v155 := GenS19{F0: 4294967295, F1: -1}
                                            _ = v155
                                        }
                                        _ = v131
                                    }
                                case "\n\t":
                                    var v156 int8 = -96
                                    _ = v156
                                }
                                _ = fi129
                            }
                        }
                    case -2147483648:
                        _ = len(v2)
                    case 539:
                        var v157 uint8 = (v16 | 2)
                        _ = v157
                    }
                } else {
                    for fi158 := 0; fi158 < 14; fi158++ {
                        if false {
                            if true {
                                v7 += (v7 - (v7 * (-(v7))))
                            } else {
                                for fi159 := 0; fi159 < 4; fi159++ {
                                    v160 := v4[uint32(v3)]
                                    if (127 >= 1) {
                                        _ = len(v2)
                                        if true {
                                            v161 := v6[(fi159 | (fi158 ^ 890))]
                                            fmt.Println(fi159, v8)
                                            _ = v161
                                        } else {
                                            fmt.Println(v160, v7, v7)
                                            v7 -= (^(v7))
                                            var v162 int = (fi159 & fi159)
                                            switch (fi158 * v162) {
                                            case 2147483647:
                                                fmt.Println(v5)
                                            case 128:
                                                v2 = append(v2, 9223372036854775807)
                                            case -9223372036854775808:
                                                v2 = append(v2, 2147483647)
                                            }
                                            _ = v162
                                        }
                                        _ = len(v2)
                                    }
                                    _ = fi159
                                    _ = v160
                                }
                                var v163 uint8 = 74
                                _ = v163
                            }
                            if (v3 == 409) {
                                for fi164 := 0; fi164 < 12; fi164++ {
                                    if true {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        switch (("a" + ((("foo bar" + "a") + ("a" + "a")) + "foo bar")) + "") {
                                        case "a":
                                            v3 -= v3
                                        case "foo bar":
                                            fmt.Println(fi158, v3)
                                            switch (^(fi164)) {
                                            case 32767:
                                                fmt.Println(v16, fi164)
                                            case -2147483648:
                                                delete(v4, (v8 % ((v8) | 1)))
                                            case 42:
                                                v8 = v8
                                            default:
                                                var v165 int = 792
                                                _ = v165
                                            }
                                            v5 = v5
                                        case "\n\t":
                                            var v166 int16 = (^(v7))
                                            _ = v166
                                        }
                                        fmt.Println(v5, v8)
                                    } else {
                                        if (1 <= 0) {
                                            v7 = (v7 << ((1) & 63))
                                            var v167 uint16 = v3
                                            var v168 float32 = 0.0
                                            _ = v167
                                            _ = v168
                                        } else {
                                            switch fi164 {
                                            case -551:
                                                var v169 int16 = (v7 * -1)
                                                _ = v169
                                            case -811:
                                                v170 := GenS15{F0: v7, F1: (v3 >> ((0) & 63))}
                                                _ = v170
                                            case 254:
                                                v8 = v8
                                            default:
                                                var v171 uint16 = v3
                                                _ = v171
                                            }
                                            if (-494 >= 42) {
                                                _ = len(v4)
                                                v16 = (v16 & v16)
                                            }
                                            var v172 int = (fi158 >> ((0) & 63))
                                            _ = v172
                                        }
                                        v7 += v7
                                        v173 := GenS4{F0: v16, F1: v8}
                                        _ = v173
                                    }
                                    var v174 uint32 = v8
                                    if (!true) {
                                        v5 += (v5 + 42)
                                        var v175 uint8 = v16
                                        var v176 string = ("hello" + "a")
                                        _ = v175
                                        _ = v176
                                    } else {
                                        var v177 uint8 = v16
                                        _ = v177
                                    }
                                    _ = fi164
                                    _ = v174
                                }
                            }
                        }
                        if (!false) {
                            if (true || false) {
                                v16 = (^(v16))
                            } else {
                                if false {
                                    for fi178 := 0; fi178 < 1; fi178++ {
                                        if false {
                                            var v179 int16 = -1
                                            v8 = (v8 ^ uint32(v16))
                                            _ = v179
                                        } else {
                                            v16 -= v16
                                            fmt.Println(v5)
                                            v3 = (^(v3))
                                        }
                                        if false {
                                            if false {
                                                v5 = (^(v5))
                                            } else {
                                                v3 = v3
                                            }
                                            v180 := GenS8{F0: v5, F1: -9223372036854775808, F2: int8(v5)}
                                            v181 := GenS14{F0: false, F1: 38, F2: 2}
                                            _ = v180
                                            _ = v181
                                        } else {
                                            v6[(fi158 % ((fi158) | 1))] = ("\n\t" + "hello")
                                            v182 := GenS20{F0: (true && ((^(v3)) >= (v3 % (((v3 ^ (v3 % ((255) | 1)))) | 1)))), F1: -1, F2: true}
                                            _ = v182
                                        }
                                        v183 := GenS3{F0: "\n\t"}
                                        _ = fi178
                                        _ = v183
                                    }
                                    v7 = 42
                                    for fi184 := 0; fi184 < 3; fi184++ {
                                        var v185 uint16 = 42
                                        if (v8 >= (v8 ^ v8)) {
                                            var v186 int = (fi158 % ((fi184) | 1))
                                            if (v8 != (v8 & (^(v8)))) {
                                                fmt.Println(v186, v16)
                                                switch 0 {
                                                case -1:
                                                    v8 = v8
                                                default:
                                                    fmt.Println(v3, v16, v8)
                                                }
                                                var v187 bool = (false || false)
                                                _ = v187
                                            } else {
                                                v188 := GenS11{F0: v5, F1: 42, F2: "a"}
                                                var v189 uint64 = 2
                                                v8 = 4294967295
                                                _ = v188
                                                _ = v189
                                            }
                                            var v190 int64 = 1
                                            _ = v186
                                            _ = v190
                                        } else {
                                            if true {
                                                var v191 uint = 42
                                                var v192 int16 = 0
                                                _ = v191
                                                _ = v192
                                            } else {
                                                var v193 string = ""
                                                switch v193 {
                                                case "\n\t":
                                                    var v194 uint64 = 42
                                                    _ = v194
                                                case "foo bar":
                                                    v16 += v16
                                                default:
                                                    var v195 uint32 = (^(v8))
                                                    _ = v195
                                                }
                                                var v196 string = ""
                                                v197 := GenS21{F0: 0, F1: (v7 & v7), F2: 1}
                                                if false {
                                                    var v198 float32 = 1e10
                                                    v199 := GenS22{F0: 77, F1: (v7 + -422), F2: 127}
                                                    _ = len(v2)
                                                    _ = v198
                                                    _ = v199
                                                } else {
                                                    v200 := v6[42]
                                                    if ((fi158 & fi184) < (fi158 % (((fi184 + fi184)) | 1))) {
                                                        var v201 int16 = 725
                                                        v16 = (v16 - 255)
                                                        _ = v201
                                                    } else {
                                                        var v202 bool = true
                                                        v8 -= v8
                                                        fmt.Println(fi158, fi184, v16)
                                                        v8 = v8
                                                        var v203 int32 = int32(v3)
                                                        var v204 float32 = 1e10
                                                        _ = v202
                                                        _ = v203
                                                        _ = v204
                                                    }
                                                    _ = v200
                                                }
                                                _ = v193
                                                _ = v196
                                                _ = v197
                                            }
                                        }
                                        v205 := v6[(fi184 | -25)]
                                        _ = fi184
                                        _ = v185
                                        _ = v205
                                    }
                                } else {
                                    v206 := GenS23{F0: uint(v3), F1: fi158}
                                    _ = v206
                                }
                            }
                            for fi207 := 0; fi207 < 2; fi207++ {
                                v7 = (^(v7))
                                switch fi158 {
                                case 252:
                                    if (false && (false || false)) {
                                        fmt.Println(v8, v16, v8)
                                    }
                                case 127:
                                    v208 := GenS24{F0: fi207, F1: 2}
                                    _ = v208
                                case 42:
                                    v5 = (v5 ^ -1)
                                }
                                _ = fi207
                            }
                            v209 := GenS25{F0: "a"}
                            _ = v209
                        } else {
                            v16 -= (^(v16))
                        }
                        v7 = (^(v7))
                        _ = fi158
                    }
                    var v210 int = 128
                    if ((1.5 < 1e10) || false) {
                        delete(v6, (v210 >> ((65535) & 63)))
                        var v211 uint = 421
                        for fi212 := 0; fi212 < 5; fi212++ {
                            v213 := GenS5{F0: 42, F1: 291, F2: "foo bar"}
                            v214 := GenS26{F0: -128}
                            _ = fi212
                            _ = v213
                            _ = v214
                        }
                        _ = v211
                    } else {
                        if false {
                            var v215 uint16 = 65535
                            _ = v215
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        switch -9223372036854775808 {
                        case 128:
                            v16 = 2
                        case -684:
                            if false {
                                v8 = (v8 << ((65535) & 63))
                                v216 := GenS27{F0: 100}
                                if true {
                                    var v217 uint16 = (v3 & (v3 * v3))
                                    switch v210 {
                                    case 42:
                                        var v218 uint64 = 2
                                        _ = v218
                                    case 1:
                                        fmt.Println(v3, v3, v7)
                                    case 0:
                                        for fi219 := 0; fi219 < 13; fi219++ {
                                            if true {
                                                v220 := GenS28{F0: -10}
                                                var v221 bool = false
                                                _ = v220
                                                _ = v221
                                            }
                                            v8 -= (v8 << ((0) & 63))
                                            _ = fi219
                                        }
                                    default:
                                        v8 += (v8 & v8)
                                    }
                                    _ = v217
                                } else {
                                    if false {
                                        switch "" {
                                        case "":
                                            fmt.Println(v16)
                                        case "a":
                                            delete(v1, "foo bar")
                                        }
                                        _ = len(v2)
                                        var v222 float64 = 3.14159
                                        _ = v222
                                    }
                                }
                                _ = v216
                            }
                        case -411:
                            if (v210 <= v210) {
                                var v223 bool = (v210 == 822)
                                var v224 int = v210
                                _ = v223
                                _ = v224
                            } else {
                                v8 = (v8 >> ((0) & 63))
                                v3 = (v3 % (((^(v3))) | 1))
                            }
                        }
                    }
                    _ = v210
                }
                _ = v121
            }
        }
    }
    if false {
        var v225 int64 = int64(v7)
        var v226 int32 = v5
        fmt.Println(v225, v3)
        _ = v225
        _ = v226
    } else {
        v5 = (v5 / ((v5) | 1))
        v16 = v16
    }
    v7 -= (v7 - (v7 & (-(v7))))
    var v227 int8 = 1
    switch "a" {
    case "a":
        for fi228 := 0; fi228 < 16; fi228++ {
            v227 += 42
            v2 = append(v2, -415)
            _ = fi228
        }
    }
    v8 = (v8 ^ (v8 >> ((255) & 63)))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        v229 := GenS29{F0: (v7 ^ (v7 % ((int16(v16)) | 1))), F1: 42}
        v5 += v5
        if false {
            v3 = v3
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v230 int8 = (v227 << ((441) & 63))
            _ = v230
        } else {
            var v231 float64 = 0.0
            for fi232 := 0; fi232 < 17; fi232++ {
                var v233 bool = (!(!false))
                _ = fi232
                _ = v233
            }
            v234 := GenS30{F0: (v16 | 1), F1: (v5 != 2147483647), F2: v16}
            _ = v231
            _ = v234
        }
        _ = v229
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v227:", v227)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v227
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
