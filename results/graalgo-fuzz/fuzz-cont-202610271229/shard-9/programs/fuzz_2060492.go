package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int8; F2 string }
type GenS3 struct { F0 int64; F1 int8; F2 uint32 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8; F1 int64; F2 int64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint16; F1 int }
type GenS13 struct { F0 uint16; F1 uint8 }
type GenS14 struct { F0 bool }

func main() {
    var v1 uint16 = 2
    var v2 int = 4294967295
    v3 := make(map[int]string)
    var v4 uint8 = 0
    var v5 int = -1
    var v6 int8 = -128
    for fi7 := 0; fi7 < 4; fi7++ {
        if false {
            var v8 int = (v5 + v2)
            var v9 string = "hello"
            var v10 int32 = -952
            _ = v10
            _ = v8
            _ = v9
        } else {
            for fi11 := 0; fi11 < 20; fi11++ {
                fmt.Println(v1, v2, v4)
                _ = fi11
            }
        }
        _ = fi7
    }
    v12 := []int{}
    v5 = (v2 << ((418) & 63))
    if true {
        for fi13 := 0; fi13 < 19; fi13++ {
            delete(v3, (v5 >> ((4294967295) & 63)))
            if true {
                delete(v3, (fi13 * (v2 << ((1) & 63))))
                var v14 uint8 = v4
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = v14
            } else {
                for fi15 := 0; fi15 < 19; fi15++ {
                    if (v1 > (v1 + v1)) {
                        if ((!((0 < (^(v1))) || false)) && (!true)) {
                            if (58 >= v4) {
                                var v16 string = "hello"
                                switch v16 {
                                case "foo bar":
                                    var v17 int32 = 1
                                    _ = v17
                                case "a":
                                    var v18 int8 = v6
                                    _ = v18
                                default:
                                    v1 -= 680
                                }
                                _ = len(v3)
                                _ = v16
                            } else {
                                var v19 uint16 = (v1 / ((323) | 1))
                                if false {
                                    var v20 uint = 0
                                    switch ("\n\t" + ("\n\t" + ("" + ""))) {
                                    case "hello":
                                        v4 += (v4 - v4)
                                    case "a":
                                        v5 = fi15
                                    default:
                                        if (int64(v1) <= 1) {
                                            v21 := GenS1{F0: v1}
                                            v22 := GenS1{F0: (^(v19))}
                                            if (false && false) {
                                                if ((false && false) && false) {
                                                    if true {
                                                        if false {
                                                            var v23 uint = v20
                                                            _ = v23
                                                        } else {
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                switch v2 {
                                                                case 127:
                                                                    if (-1 == 0) {
                                                                        v19 -= (v19 / (((v1 + (^(v1)))) | 1))
                                                                        v4 = (^(v4))
                                                                    }
                                                                default:
                                                                    v5 -= fi13
                                                                }
                                                            default:
                                                                fmt.Println(v4, v6)
                                                            }
                                                            v24 := GenS1{F0: v1}
                                                            _ = len(v3)
                                                            v2 += 42
                                                            v25 := GenS1{F0: 0}
                                                            _ = v24
                                                            _ = v25
                                                        }
                                                    } else {
                                                        if false {
                                                            var v26 string = ("hello" + (("\n\t" + "") + "hello"))
                                                            _ = v26
                                                        }
                                                    }
                                                    v27 := GenS2{F0: "\n\t", F1: v6, F2: "\n\t"}
                                                    v28 := GenS1{F0: (v19 + (v1 & v19))}
                                                    _ = v27
                                                    _ = v28
                                                } else {
                                                    var v29 uint16 = v1
                                                    var v30 uint = v20
                                                    _ = v29
                                                    _ = v30
                                                }
                                                switch (fi15 / ((127) | 1)) {
                                                case 988:
                                                    var v31 string = ("foo bar" + "hello")
                                                    _ = v31
                                                default:
                                                    v32 := GenS1{F0: (v1 << ((uint(v4)) & 63))}
                                                    _ = v32
                                                }
                                            } else {
                                                var v33 int64 = 147
                                                switch "foo bar" {
                                                case "a":
                                                    v34 := GenS3{F0: (-(v33)), F1: v6, F2: 4294967295}
                                                    _ = v34
                                                case "foo bar":
                                                    _ = len(v12)
                                                }
                                                if (!true) {
                                                    switch (("" + ("foo bar" + "foo bar")) + "a") {
                                                    case "hello":
                                                        v2 += 42
                                                    }
                                                    fmt.Println(v5)
                                                    fmt.Println(v4, v1, v1)
                                                    v35 := GenS4{F0: (v6 + 92), F1: -44}
                                                    v20 = v20
                                                    _ = v35
                                                } else {
                                                    fmt.Println(v6, v5)
                                                    if len(v12) >= 2 {
                                                        v12 = v12[1:len(v12)]
                                                    }
                                                    if true {
                                                        v5 = (fi13 | fi15)
                                                        v6 = v6
                                                    } else {
                                                        var v36 int64 = v33
                                                        var v37 float64 = 1e10
                                                        var v38 uint = v20
                                                        _ = v36
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                }
                                                _ = v33
                                            }
                                            _ = v21
                                            _ = v22
                                        } else {
                                            v19 = v19
                                        }
                                        if (55 > 160) {
                                            fmt.Println(v4)
                                        } else {
                                            fmt.Println(v19)
                                            switch "a" {
                                            case "hello":
                                                var v39 uint32 = 1
                                                _ = v39
                                            case "a":
                                                var v40 bool = (false && (-38 <= 86))
                                                _ = v40
                                            case "foo bar":
                                                delete(v3, -2147483648)
                                            }
                                        }
                                        if false {
                                            v5 = v5
                                        }
                                    }
                                    _ = v20
                                } else {
                                    if false {
                                        v41 := GenS2{F0: ("" + ((("foo bar" + "hello") + ((("\n\t" + "a") + "\n\t") + "\n\t")) + "foo bar")), F1: v6, F2: ("hello" + ("foo bar" + (("hello" + "a") + "\n\t")))}
                                        switch ("hello" + "foo bar") {
                                        case "hello":
                                            v4 -= (v4 * (^(v4)))
                                        case "":
                                            v19 += (v19 >> ((682) & 63))
                                        case "\n\t":
                                            var v42 uint = 42
                                            _ = v42
                                        default:
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                        }
                                        _ = len(v3)
                                        var v43 uint8 = (v4 | v4)
                                        var v44 float64 = 1.5
                                        var v45 int = v2
                                        if true {
                                            switch "" {
                                            case "\n\t":
                                                v43 = (v4 - v43)
                                            default:
                                                var v46 string = ("" + ("foo bar" + "foo bar"))
                                                _ = v46
                                            }
                                            v43 += v4
                                        }
                                        _ = v41
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    }
                                    if false {
                                        v47 := GenS5{F0: (false && (65535 >= 2))}
                                        var v48 int = -798
                                        var v49 int64 = 9223372036854775807
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                    }
                                }
                                _ = v19
                            }
                            if (-32768 != -1) {
                                _ = len(v3)
                                if (1 >= -2147483648) {
                                    v6 += (v6 / ((v6) | 1))
                                    switch ("" + "foo bar") {
                                    case "foo bar":
                                        if true {
                                            var v50 uint8 = (v4 / ((v4) | 1))
                                            fmt.Println(fi15, fi15)
                                            _ = v50
                                        } else {
                                            var v51 int16 = 968
                                            _ = len(v12)
                                            var v52 int8 = (v6 + (v6 / ((v6) | 1)))
                                            _ = v51
                                            _ = v52
                                        }
                                    }
                                    var v53 int = (v2 & fi13)
                                    _ = v53
                                } else {
                                    v1 = (v1 - (v1 - v1))
                                }
                            }
                        } else {
                            fmt.Println(v6)
                        }
                    }
                    switch (-(fi13)) {
                    case 192:
                        v54 := v3[(-(fi13))]
                        _ = v54
                    case -123:
                        v1 = v1
                    case 9223372036854775807:
                        if (0 == 2147483647) {
                            if true {
                                v55 := GenS5{F0: ((-2.25 * float64(((0.1 + (3.14159 - -2.25)) * 0.1))) >= 0.1)}
                                _ = v55
                            } else {
                                switch "hello" {
                                case "foo bar":
                                    v56 := GenS6{F0: v1}
                                    v12 = append(v12, v2)
                                    var v57 uint64 = 641
                                    _ = v56
                                    _ = v57
                                case "\n\t":
                                    if (!true) {
                                        if false {
                                            v1 += v1
                                            v6 = (v6 ^ (v6 | 0))
                                        }
                                        if (uint(v4) < 65535) {
                                            v58 := GenS2{F0: ("\n\t" + ("\n\t" + ("foo bar" + "foo bar"))), F1: v6, F2: "\n\t"}
                                            var v59 uint = 1
                                            _ = len(v3)
                                            _ = v58
                                            _ = v59
                                        } else {
                                            if true {
                                                var v60 uint8 = v4
                                                v6 += v6
                                                v1 += (^(v1))
                                                _ = v60
                                            }
                                        }
                                    } else {
                                        _ = len(v12)
                                    }
                                case "a":
                                    delete(v3, (fi15 ^ -1))
                                default:
                                    var v61 string = "a"
                                    _ = v61
                                }
                                v62 := GenS7{F0: 80, F1: -1, F2: 340}
                                if true {
                                    switch "a" {
                                    case "a":
                                        var v63 uint16 = 1
                                        _ = v63
                                    }
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    v64 := GenS8{F0: 0}
                                    _ = v64
                                } else {
                                    fmt.Println(fi13)
                                }
                                _ = v62
                            }
                            switch fi13 {
                            case 732:
                                if true {
                                    v4 -= (^(v4))
                                } else {
                                    v12 = append(v12, fi15)
                                    v1 = v1
                                    v65 := GenS5{F0: true}
                                    v12 = append(v12, -793)
                                    var v66 uint64 = 1
                                    fmt.Println(fi13, v5)
                                    fmt.Println(fi13, v66, fi15)
                                    _ = v65
                                    _ = v66
                                }
                            default:
                                v67 := GenS9{F0: ((v6 + 100) <= (^(v6))), F1: 0, F2: 0}
                                _ = v67
                            }
                            var v68 int16 = -1
                            _ = v68
                        } else {
                            v3[(v2 << ((4294967295) & 63))] = ("" + "hello")
                        }
                    default:
                        v5 += v2
                    }
                    _ = fi15
                }
            }
            fmt.Println(v2, v5)
            _ = fi13
        }
    } else {
        for fi69 := 0; fi69 < 9; fi69++ {
            var v70 int64 = -996
            if len(v12) > 0 {
                v12[0] = (v2 / ((int(v6)) | 1))
            }
            _ = len(v3)
            _ = fi69
            _ = v70
        }
        if true {
            for fi71 := 0; fi71 < 7; fi71++ {
                var v72 int16 = -615
                if ((true || (!false)) || false) {
                    var v73 int16 = (^(v72))
                    var v74 int32 = -1
                    _ = v73
                    _ = v74
                } else {
                    for fi75 := 0; fi75 < 16; fi75++ {
                        var v76 uint = 255
                        _ = fi75
                        _ = v76
                    }
                }
                if len(v12) > 0 {
                    v12[0] = v2
                }
                _ = fi71
                _ = v72
            }
        }
    }
    var v77 int32 = 2147483647
    v78 := GenS5{F0: ("\n\t" == "foo bar")}
    if true {
        var v79 bool = (0 >= 18446744073709551615)
        v1 = 1
        for fi80 := 0; fi80 < 10; fi80++ {
            if (false || (v77 >= -452)) {
                var v81 uint64 = 698
                var v82 float64 = (0.1 - -2.25)
                v83 := GenS5{F0: (!v79)}
                _ = v81
                _ = v82
                _ = v83
            }
            _ = fi80
        }
        _ = v79
    }
    if false {
        var v84 int = int(v1)
        _ = v84
    }
    var v85 int32 = (v77 & (v77 | -2147483648))
    var v86 float64 = -2.25
    switch "foo bar" {
    case "a":
        v4 += v4
    case "hello":
        var v87 uint = 467
        _ = v87
    case "\n\t":
        var v88 string = "foo bar"
        _ = v88
    }
    switch "" {
    case "a":
        for fi89 := 0; fi89 < 13; fi89++ {
            if false {
                var v90 uint32 = 4294967295
                _ = v90
            } else {
                delete(v3, v5)
                fmt.Println(v77, v2, v6)
                v77 = v77
            }
            _ = fi89
        }
    case "\n\t":
        if (3.14159 >= float32((v86 / v86))) {
            if (4294967295 != 2) {
                for fi91 := 0; fi91 < 2; fi91++ {
                    if false {
                        var v92 uint16 = 255
                        if (!false) {
                            for fi93 := 0; fi93 < 11; fi93++ {
                                fmt.Println(v6, v5, v2)
                                if (-44 >= (v6 >> ((251) & 63))) {
                                    v94 := GenS1{F0: v92}
                                    v1 = 65535
                                    var v95 string = "hello"
                                    _ = v94
                                    _ = v95
                                } else {
                                    switch "a" {
                                    case "hello":
                                        if false {
                                            v6 = -128
                                            v96 := GenS8{F0: 4294967295}
                                            fmt.Println(v1, v1, v5)
                                            _ = v96
                                        }
                                    case "foo bar":
                                        fmt.Println(v77, v4, v92)
                                    case "\n\t":
                                        var v97 uint8 = (^(v4))
                                        _ = v97
                                    default:
                                        if true {
                                            fmt.Println(v1, v6, v77)
                                            var v98 float32 = 1e10
                                            _ = v98
                                        }
                                    }
                                    var v99 string = ("hello" + ("foo bar" + ""))
                                    _ = v99
                                }
                                if ((v85 / ((v77) | 1)) >= (v85 / ((v85) | 1))) {
                                    fmt.Println(v5)
                                }
                                fmt.Println(v77, fi93, v5)
                                v5 += fi93
                                fmt.Println(v1, fi91, v2)
                                if false {
                                    v77 = v85
                                    v6 -= 28
                                    v100 := GenS1{F0: (v1 ^ 2)}
                                    _ = v100
                                }
                                v1 -= (v1 / ((v1) | 1))
                                _ = fi93
                            }
                            v101 := GenS4{F0: v6, F1: -1}
                            _ = v101
                        } else {
                            switch "hello" {
                            case "a":
                                v102 := GenS7{F0: 103, F1: 1, F2: -9223372036854775808}
                                _ = v102
                            case "\n\t":
                                var v103 int = v5
                                _ = v103
                            default:
                                var v104 int8 = (v6 & v6)
                                _ = v104
                            }
                            var v105 float64 = v86
                            _ = v105
                        }
                        _ = v92
                    }
                    v106 := GenS8{F0: 18446744073709551615}
                    if (v77 < v85) {
                        v12 = append(v12, int(v85))
                        for fi107 := 0; fi107 < 10; fi107++ {
                            v1 += v1
                            v108 := GenS10{F0: (v6 | (v6 ^ -1))}
                            var v109 uint = 671
                            _ = fi107
                            _ = v108
                            _ = v109
                        }
                    }
                    _ = fi91
                    _ = v106
                }
            }
        } else {
            if true {
                for fi110 := 0; fi110 < 19; fi110++ {
                    var v111 uint = 677
                    _ = len(v12)
                    _ = fi110
                    _ = v111
                }
                delete(v3, (-(v5)))
                if (v4 == (^(v4))) {
                    for fi112 := 0; fi112 < 13; fi112++ {
                        var v113 int32 = (v77 * (-(v77)))
                        _ = fi112
                        _ = v113
                    }
                    for fi114 := 0; fi114 < 3; fi114++ {
                        v3[-9223372036854775808] = ("foo bar" + ("hello" + (("a" + "foo bar") + "")))
                        _ = fi114
                    }
                    v115 := GenS3{F0: int64(v2), F1: (v6 / (((v6 >> ((0) & 63))) | 1)), F2: 534}
                    _ = v115
                }
            }
            v2 += v5
        }
    case "foo bar":
        var v116 string = ((("a" + "a") + "a") + ("\n\t" + ("hello" + ("" + ("\n\t" + "\n\t")))))
        _ = v116
    default:
        for fi117 := 0; fi117 < 5; fi117++ {
            v118 := GenS11{F0: -32768}
            _ = fi117
            _ = v118
        }
    }
    switch v5 {
    case 1:
        fmt.Println(v2, v77, v2)
    }
    var v119 int16 = -1
    if (v119 == (v119 - (v119 % ((v119) | 1)))) {
        var v120 int32 = v77
        v5 -= (v2 << ((1) & 63))
        v77 += int32(v119)
        _ = v120
    }
    switch 0 {
    case 4294967295:
        fmt.Println(v6, v119)
    case -9223372036854775808:
        if false {
            var v121 int64 = 1
            v119 = (v119 % (((v119 >> ((42) & 63))) | 1))
            var v122 int32 = (^(v85))
            _ = v121
            _ = v122
        } else {
            v123 := GenS12{F0: v1, F1: (v5 | v2)}
            switch v2 {
            case 9223372036854775807:
                if false {
                    var v124 float32 = 0.1
                    if ((v1 / ((v1) | 1)) == v1) {
                        v85 -= v85
                    }
                    _ = len(v3)
                    _ = v124
                }
            case -9223372036854775808:
                fmt.Println(v6, v1, v1)
            case 0:
                var v125 int8 = (v6 & 100)
                _ = v125
            default:
                if true {
                    fmt.Println(v5, v4)
                    if (2147483647 >= 789) {
                        v2 = 42
                        var v126 float32 = ((1e10 - 3.14159) + 3.14159)
                        _ = v126
                    } else {
                        var v127 int32 = (^(v77))
                        delete(v3, (-(v2)))
                        _ = v127
                    }
                    var v128 int = 2147483647
                    _ = v128
                } else {
                    v77 = v77
                    if (892 != 745) {
                        if false {
                            if (2 == (v4 << ((42) & 63))) {
                                if (9223372036854775807 > 1) {
                                    if true {
                                        if (false || true) {
                                            v5 = -1
                                            var v129 int8 = (^(v6))
                                            switch ("\n\t" + ("\n\t" + ("hello" + "\n\t"))) {
                                            case "a":
                                                var v130 float64 = v86
                                                _ = v130
                                            case "foo bar":
                                                switch "\n\t" {
                                                case "":
                                                    v131 := GenS6{F0: v1}
                                                    _ = v131
                                                default:
                                                    if (42 <= 411) {
                                                        _ = len(v3)
                                                        if ((v1 - 0) <= uint16(v4)) {
                                                            v132 := GenS11{F0: v119}
                                                            _ = v132
                                                        } else {
                                                            for fi133 := 0; fi133 < 20; fi133++ {
                                                                if ((v85 >= int32(v119)) || false) {
                                                                    var v134 uint32 = 65535
                                                                    v135 := GenS13{F0: v1, F1: 42}
                                                                    _ = v134
                                                                    _ = v135
                                                                }
                                                                v12 = append(v12, (v5 >> ((1) & 63)))
                                                                if false {
                                                                    if true {
                                                                        var v136 int16 = v119
                                                                        _ = v136
                                                                    }
                                                                    v137 := GenS7{F0: (v4 | v4), F1: 9223372036854775807, F2: -975}
                                                                    _ = v137
                                                                }
                                                                _ = fi133
                                                            }
                                                            for fi138 := 0; fi138 < 7; fi138++ {
                                                                var v139 uint16 = (v1 ^ (^(v1)))
                                                                if false {
                                                                    v77 += v85
                                                                } else {
                                                                    v129 += v6
                                                                    for fi140 := 0; fi140 < 17; fi140++ {
                                                                        var v141 uint8 = 127
                                                                        v139 = v1
                                                                        if len(v12) > 0 {
                                                                            v12[0] = v5
                                                                        }
                                                                        _ = fi140
                                                                        _ = v141
                                                                    }
                                                                }
                                                                _ = fi138
                                                                _ = v139
                                                            }
                                                        }
                                                        var v142 int = -9223372036854775808
                                                        _ = v142
                                                    } else {
                                                        v143 := GenS2{F0: "", F1: (v129 % (((^(v6))) | 1)), F2: ("a" + "")}
                                                        fmt.Println(v2)
                                                        _ = v143
                                                    }
                                                }
                                            }
                                            _ = v129
                                        }
                                        v85 = 42
                                        if false {
                                            v119 = (^(v119))
                                        }
                                    } else {
                                        v3[1] = "hello"
                                        v144 := GenS7{F0: (v4 * v4), F1: -1, F2: -1}
                                        _ = v144
                                    }
                                } else {
                                    if true {
                                        v4 += (v4 << ((1) & 63))
                                    }
                                }
                                if true {
                                    for fi145 := 0; fi145 < 15; fi145++ {
                                        v86 = 0.0
                                        fmt.Println(v119, v2, v5)
                                        _ = fi145
                                    }
                                    v1 = v1
                                }
                                fmt.Println(v6)
                            }
                            v86 = v86
                            var v146 uint = 65535
                            _ = v146
                        } else {
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            var v147 int16 = int16(v6)
                            switch "" {
                            case "a":
                                switch "" {
                                case "foo bar":
                                    v148 := GenS14{F0: true}
                                    _ = v148
                                case "\n\t":
                                    for fi149 := 0; fi149 < 13; fi149++ {
                                        var v150 int64 = 385
                                        _ = fi149
                                        _ = v150
                                    }
                                }
                            case "foo bar":
                                var v151 uint64 = 653
                                _ = v151
                            case "hello":
                                var v152 int8 = v6
                                _ = v152
                            default:
                                v77 -= (v77 >> ((2) & 63))
                            }
                            _ = v147
                        }
                        v153 := GenS1{F0: 255}
                        _ = v153
                    } else {
                        v5 = (v2 ^ v2)
                        var v154 float64 = float64(((3.14159 + -2.25) / ((1e10 / -2.25) + (1.5 + ((((3.14159 * -2.25) + (0.1 + (0.0 * 1e10))) / 1.5) + 3.14159)))))
                        v2 -= v5
                        _ = v154
                    }
                    for fi155 := 0; fi155 < 8; fi155++ {
                        v12 = append(v12, 127)
                        _ = fi155
                    }
                }
            }
            _ = v123
        }
    case -2147483648:
        v2 += (v5 >> ((65535) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v119
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
    _ = v78
    _ = v85
    _ = v86
}
