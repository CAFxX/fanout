package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int32; F2 int32 }
type GenS4 struct { F0 uint32; F1 string; F2 int32 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 uint8; F2 bool }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int; F1 int; F2 int8 }
type GenS11 struct { F0 uint16; F1 int16; F2 int16 }
type GenS12 struct { F0 int16; F1 int8; F2 int }
type GenS13 struct { F0 string; F1 uint32; F2 int }
type GenS14 struct { F0 uint8; F1 uint8 }
type GenS15 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint64, p1 uint, p2 int) int64 {
    if (p0 == (p0 << ((p1) & 63))) {
        v5 := GenS1{F0: "foo bar", F1: p1, F2: -32768}
        _ = v5
    }
    if (!false) {
        if (true || true) {
            switch "a" {
            case "hello":
                v6 := []int{(p2 | p2), 42, p2, (p2 ^ (^(p2)))}
                _ = v6
            case "foo bar":
                for fi7 := 0; fi7 < 14; fi7++ {
                    var v8 float64 = 1e10
                    if (!(846 < 2)) {
                        p2 += (fi7 & p2)
                        fmt.Println(v8, p0, p0)
                        p1 -= 255
                    }
                    _ = fi7
                    _ = v8
                }
            case "":
                v9 := GenS2{F0: 50}
                _ = v9
            }
        } else {
            v10 := GenS3{F0: ("a" + ("" + "foo bar")), F1: 582, F2: -300}
            if true {
                var v11 bool = true
                if (-1 != 1) {
                    v11 = (0 < 100)
                    p0 -= p0
                    v11 = v11
                } else {
                    if (-128 <= -1) {
                        var v12 int64 = 813
                        if ((p2 + (p2 + (p2 << ((315) & 63)))) >= -566) {
                            if true {
                                fmt.Println(p1, v11)
                                for fi13 := 0; fi13 < 6; fi13++ {
                                    if v11 {
                                        v14 := GenS3{F0: "\n\t", F1: 1, F2: 0}
                                        switch ("a" + ("foo bar" + ((("" + "hello") + ("hello" + ("a" + "\n\t"))) + "\n\t"))) {
                                        case "hello":
                                            var v15 int64 = (-(v12))
                                            _ = v15
                                        }
                                        _ = v14
                                    } else {
                                        if (!v11) {
                                            var v16 int8 = 1
                                            for fi17 := 0; fi17 < 6; fi17++ {
                                                p0 = 42
                                                p1 -= p1
                                                var v18 bool = v11
                                                _ = fi17
                                                _ = v18
                                            }
                                            v19 := GenS4{F0: 2, F1: "a", F2: -1}
                                            _ = v16
                                            _ = v19
                                        }
                                    }
                                    _ = fi13
                                }
                                var v20 bool = v11
                                _ = v20
                            } else {
                                var v21 uint32 = 2147483647
                                v12 = (-(v12))
                                v22 := []int{p2}
                                _ = v21
                                _ = v22
                            }
                            p2 = (p2 * 694)
                            var v23 int32 = int32(v12)
                            _ = v23
                        } else {
                            fmt.Println(v11, v11, v11)
                            for fi24 := 0; fi24 < 3; fi24++ {
                                p2 -= p2
                                _ = fi24
                            }
                        }
                        _ = v12
                    } else {
                        v25 := []int{}
                        _ = v25
                    }
                    if true {
                        var v26 uint8 = 255
                        switch "foo bar" {
                        case "\n\t":
                            if false {
                                fmt.Println(v11)
                            }
                        default:
                            if (192 != 1) {
                                if v11 {
                                    var v27 uint64 = p0
                                    _ = v27
                                }
                                v28 := GenS5{F0: 42, F1: 42}
                                fmt.Println(p2, p1)
                                _ = v28
                            }
                        }
                        _ = v26
                    }
                }
                var v29 bool = v11
                _ = v11
                _ = v29
            } else {
                v30 := GenS5{F0: 338, F1: -924}
                v31 := GenS3{F0: ((("\n\t" + "a") + "a") + ("" + "foo bar")), F1: -1, F2: 42}
                p0 -= uint64(p1)
                _ = v30
                _ = v31
            }
            var v32 bool = false
            _ = v10
            _ = v32
        }
        if (true || false) {
            v33 := []int{}
            _ = v33
        } else {
            var v34 int8 = 127
            for fi35 := 0; fi35 < 14; fi35++ {
                p2 += (fi35 / ((fi35) | 1))
                v36 := GenS6{F0: 1}
                p1 += (^(p1))
                _ = fi35
                _ = v36
            }
            _ = v34
        }
    }
    switch (^(p2)) {
    case -155:
        fmt.Println(p2, p2)
    case 32767:
        v37 := GenS4{F0: 4294967295, F1: ("foo bar" + "a"), F2: 2147483647}
        _ = v37
    default:
        p0 += (p0 / (((^(p0))) | 1))
    }
    v38 := GenS1{F0: "", F1: (p1 << ((p1) & 63)), F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v38
    return -1
}
func genFn1(p0 bool) bool {
    var v39 uint64 = 532
    p0 = (!((!(2147483647 != 0)) && p0))
    _ = p0
    _ = v39
    return (-2147483648 >= 32767)
}

func main() {
    var v1 int = -9223372036854775808
    var v2 string = ((("\n\t" + "foo bar") + ("foo bar" + "foo bar")) + ((("foo bar" + ("foo bar" + "")) + (("a" + "a") + "foo bar")) + ""))
    v3 := []string{v2, "foo bar", v2}
    v4 := make(map[uint32]int)
    v40 := GenS2{F0: 2}
    v1 += 32767
    if true {
        v41 := GenS6{F0: 615}
        _ = v41
    } else {
        delete(v4, 563)
        var v42 uint = 4294967295
        _ = v42
    }
    fmt.Println(v1, v1)
    v4[126] = 2147483647
    var v43 int8 = 42
    v44 := GenS7{F0: 255, F1: 255, F2: (true && ((3.14159 + 1e10) < -2.25))}
    v1 -= v1
    v43 += (^(v43))
    for fi45 := 0; fi45 < 20; fi45++ {
        v46 := GenS3{F0: ("\n\t" + (v2 + v2)), F1: 42, F2: 42}
        _ = fi45
        _ = v46
    }
    switch (v1 << ((2) & 63)) {
    case 354:
        if true {
            if len(v3) > 0 {
                v3[0] = "a"
            }
            var v47 bool = false
            _ = v47
        } else {
            if true {
                if false {
                    v48 := GenS6{F0: 32767}
                    _ = v48
                } else {
                    var v49 int8 = 92
                    _ = v49
                }
                var v50 int16 = 4
                _ = v50
            } else {
                for fi51 := 0; fi51 < 14; fi51++ {
                    if (!(42 <= 782)) {
                        for fi52 := 0; fi52 < 4; fi52++ {
                            v1 -= fi52
                            switch (v1 | (fi51 ^ v1)) {
                            case 0:
                                var v53 string = v2
                                _ = v53
                            case -930:
                                v54 := GenS4{F0: 1, F1: v2, F2: 1}
                                var v55 string = ""
                                if len(v3) > 0 {
                                    v3[0] = v2
                                }
                                _ = v54
                                _ = v55
                            case 42:
                                if false {
                                    fmt.Println(v43)
                                    if (586 >= 4294967295) {
                                        if ((true || (true && ((v2 >= "hello") || false))) || false) {
                                            if false {
                                                if false {
                                                    _ = len(v3)
                                                } else {
                                                    if true {
                                                        var v56 uint32 = 2147483648
                                                        v2 = "a"
                                                        _ = v56
                                                    }
                                                    delete(v4, 42)
                                                }
                                                v1 -= (v1 >> ((18446744073709551615) & 63))
                                                switch fi51 {
                                                case -130:
                                                    v57 := GenS8{F0: 0}
                                                    _ = v57
                                                }
                                            } else {
                                                if true {
                                                    switch "a" {
                                                    case "a":
                                                        var v58 int64 = 0
                                                        _ = v58
                                                    case "foo bar":
                                                        v43 = -105
                                                    }
                                                    var v59 float64 = 0.1
                                                    var v60 uint = 436
                                                    switch (-(fi51)) {
                                                    case 148:
                                                        v60 = v60
                                                    case 1:
                                                        if (1 > 4294967295) {
                                                            v43 = (^(v43))
                                                        } else {
                                                            if ((v43 & -1) == v43) {
                                                                fmt.Println(fi52)
                                                                fmt.Println(fi52, v1)
                                                                var v61 int64 = 0
                                                                _ = v61
                                                            } else {
                                                                v62 := GenS7{F0: 255, F1: 2, F2: true}
                                                                _ = v62
                                                            }
                                                        }
                                                    case 2147483647:
                                                        if (-1 > -192) {
                                                            v60 -= (v60 * v60)
                                                            if (genFn0(18446744073709551615, 375, fi51) == 42) {
                                                                v60 = 65535
                                                                fmt.Println(v59, fi51)
                                                            }
                                                            if true {
                                                                var v63 uint64 = 0
                                                                if false {
                                                                    delete(v4, 140)
                                                                    v63 = v63
                                                                    _ = len(v4)
                                                                } else {
                                                                    switch v2 {
                                                                    case "a":
                                                                        v64 := GenS9{F0: int8(fi52)}
                                                                        _ = v64
                                                                    default:
                                                                        v59 -= v59
                                                                        var v65 int64 = -1
                                                                        if ((v65 - v65) == int64(v43)) {
                                                                            if len(v3) > 0 {
                                                                                v3[0] = (v2 + "")
                                                                            }
                                                                        }
                                                                        _ = v65
                                                                    }
                                                                    switch -50 {
                                                                    case 9223372036854775807:
                                                                        v60 = v60
                                                                    }
                                                                    v66 := GenS1{F0: v2, F1: v60, F2: -32768}
                                                                    _ = v66
                                                                }
                                                                v63 -= v63
                                                                delete(v4, 753)
                                                                switch ("hello" + v2) {
                                                                case "hello":
                                                                    var v67 int = (fi51 ^ int(v63))
                                                                    _ = v67
                                                                case "a":
                                                                    if (-2147483648 == (v1 >> ((18446744073709551615) & 63))) {
                                                                        _ = len(v3)
                                                                    } else {
                                                                        v68 := GenS10{F0: (v1 << ((v60) & 63)), F1: 9223372036854775807, F2: v43}
                                                                        var v69 uint64 = v63
                                                                        v2 = "hello"
                                                                        v69 -= v69
                                                                        v2 = v2
                                                                        fmt.Println(v2)
                                                                        _ = v68
                                                                        _ = v69
                                                                    }
                                                                case "foo bar":
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        if true {
                                                                            if ((false || (fi52 < (^(fi51)))) && (1000 != 1000)) {
                                                                                var v70 uint = v60
                                                                                v71 := GenS11{F0: 973, F1: -32768, F2: 782}
                                                                                _ = v70
                                                                                _ = v71
                                                                            } else {
                                                                                switch v2 {
                                                                                case "a":
                                                                                    var v72 string = v2
                                                                                    _ = v72
                                                                                case "hello":
                                                                                    v73 := GenS7{F0: 1, F1: 42, F2: false}
                                                                                    _ = v73
                                                                                }
                                                                                switch v1 {
                                                                                case 578:
                                                                                    var v74 float64 = 0.0
                                                                                    _ = v74
                                                                                case 9223372036854775807:
                                                                                    v75 := GenS2{F0: 255}
                                                                                    _ = v75
                                                                                }
                                                                            }
                                                                        }
                                                                    case "":
                                                                        v2 = "\n\t"
                                                                    case "foo bar":
                                                                        var v76 int16 = 778
                                                                        _ = v76
                                                                    }
                                                                    v4[4294967295] = (v1 << ((v60) & 63))
                                                                    var v77 uint32 = 819
                                                                    _ = v77
                                                                default:
                                                                    if (40 > 127) {
                                                                        v78 := GenS10{F0: v1, F1: v1, F2: 127}
                                                                        v43 -= v43
                                                                        _ = v78
                                                                    }
                                                                }
                                                                var v79 float32 = 0.1
                                                                _ = v63
                                                                _ = v79
                                                            }
                                                        }
                                                    }
                                                    v80 := GenS12{F0: 0, F1: v43, F2: (-(v1))}
                                                    _ = v59
                                                    _ = v60
                                                    _ = v80
                                                }
                                                if true {
                                                    switch v2 {
                                                    case "\n\t":
                                                        if (true && (!false)) {
                                                            var v81 float32 = 1e10
                                                            _ = v81
                                                        } else {
                                                            var v82 bool = (!(!true))
                                                            v2 = "a"
                                                            fmt.Println(fi52, v82, v82)
                                                            var v83 uint32 = 42
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                        if (-2147483648 >= 2147483647) {
                                                            if false {
                                                                fmt.Println(fi52, fi52, v2)
                                                                var v84 uint16 = 255
                                                                v85 := GenS3{F0: "hello", F1: -1, F2: 2147483647}
                                                                _ = v84
                                                                _ = v85
                                                            }
                                                            _ = len(v4)
                                                        } else {
                                                            fmt.Println(v2)
                                                            switch v1 {
                                                            case 32767:
                                                                v86 := GenS3{F0: "hello", F1: 1, F2: 2147483647}
                                                                _ = v86
                                                            case -9223372036854775808:
                                                                fmt.Println(fi52, v2, v2)
                                                            case -821:
                                                                var v87 int32 = 2147483647
                                                                _ = v87
                                                            default:
                                                                if false {
                                                                    v88 := GenS9{F0: int8(fi52)}
                                                                    fmt.Println(fi52)
                                                                    switch -2147483648 {
                                                                    case 2147483647:
                                                                        switch (v2 + ((((v2 + ("" + (("" + (v2 + ("a" + v2))) + v2))) + "hello") + ("a" + "")) + v2)) {
                                                                        case "\n\t":
                                                                            v89 := GenS8{F0: 1}
                                                                            _ = v89
                                                                        }
                                                                    default:
                                                                        var v90 bool = false
                                                                        _ = v90
                                                                    }
                                                                    _ = v88
                                                                } else {
                                                                    var v91 int16 = 1
                                                                    _ = v91
                                                                }
                                                            }
                                                            v43 = (v43 | v43)
                                                        }
                                                    default:
                                                        fmt.Println(fi51, v43)
                                                    }
                                                    if true {
                                                        var v92 int8 = 100
                                                        var v93 uint64 = 4294967295
                                                        v94 := GenS3{F0: ("\n\t" + (v2 + v2)), F1: -94, F2: -106}
                                                        if true {
                                                            if ((v2 + (((v2 + v2) + v2) + v2)) == "a") {
                                                                var v95 uint = 0
                                                                var v96 int = (-(v1))
                                                                _ = v95
                                                                _ = v96
                                                            }
                                                            var v97 float64 = 3.14159
                                                            fmt.Println(v97)
                                                            _ = v97
                                                        }
                                                        _ = v92
                                                        _ = v93
                                                        _ = v94
                                                    } else {
                                                        v98 := GenS1{F0: v2, F1: 1, F2: -32768}
                                                        _ = v98
                                                    }
                                                } else {
                                                    var v99 bool = true
                                                    if len(v3) > 0 {
                                                        v3[0] = ((("foo bar" + v2) + v2) + v2)
                                                    }
                                                    _ = v99
                                                }
                                            }
                                        } else {
                                            var v100 uint16 = 42
                                            v101 := GenS13{F0: v2, F1: 174, F2: 9223372036854775807}
                                            v4[1] = (v1 | (^(v1)))
                                            switch 4294967295 {
                                            case 128:
                                                switch fi52 {
                                                case 1:
                                                    v102 := GenS1{F0: (v2 + ((v2 + "hello") + v2)), F1: 18446744073709551615, F2: 1000}
                                                    _ = v102
                                                case 32767:
                                                    v103 := v4[4294967295]
                                                    _ = v103
                                                case 2147483647:
                                                    v104 := v4[2]
                                                    _ = v104
                                                }
                                            case 4294967295:
                                                v105 := GenS4{F0: 0, F1: (v2 + "a"), F2: 0}
                                                _ = v105
                                            default:
                                                v2 = v2
                                            }
                                            var v106 int16 = 755
                                            switch 127 {
                                            case 9223372036854775807:
                                                var v107 uint8 = 128
                                                _ = v107
                                            default:
                                                fmt.Println(v106, v2, v43)
                                                v1 -= fi51
                                            }
                                            _ = v100
                                            _ = v101
                                            _ = v106
                                        }
                                    }
                                    if true {
                                        var v108 float64 = 1e10
                                        v43 += (v43 % ((v43) | 1))
                                        _ = v108
                                    }
                                    v2 = "foo bar"
                                } else {
                                    var v109 int16 = 0
                                    if len(v3) > 0 {
                                        v3[0] = ((v2 + v2) + v2)
                                    }
                                    _ = v109
                                }
                            default:
                                v43 = (v43 << ((18446744073709551615) & 63))
                            }
                            _ = fi52
                        }
                        v110 := GenS6{F0: 2}
                        _ = v110
                    } else {
                        var v111 int64 = 42
                        if true {
                            v112 := v4[42]
                            _ = v112
                        } else {
                            for fi113 := 0; fi113 < 12; fi113++ {
                                if false {
                                    v114 := GenS12{F0: -1, F1: v43, F2: 0}
                                    _ = v114
                                }
                                _ = fi113
                            }
                            fmt.Println(v2, v2, v43)
                            v4[2147483648] = fi51
                        }
                        v115 := v4[0]
                        _ = v111
                        _ = v115
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi51
                }
                fmt.Println(v2)
            }
            for fi116 := 0; fi116 < 9; fi116++ {
                var v117 int16 = 393
                switch (("a" + "\n\t") + ("" + v2)) {
                case "a":
                    for fi118 := 0; fi118 < 7; fi118++ {
                        switch ("a" + v2) {
                        case "":
                            var v119 uint = 279
                            _ = v119
                        default:
                            v120 := GenS14{F0: 128, F1: 11}
                            _ = v120
                        }
                        var v121 float64 = (-2.25 / 1e10)
                        _ = fi118
                        _ = v121
                    }
                case "\n\t":
                    switch "hello" {
                    case "\n\t":
                        for fi122 := 0; fi122 < 18; fi122++ {
                            var v123 int16 = (v117 * (v117 | (^(v117))))
                            v124 := GenS15{F0: (v123 | -776), F1: 43}
                            v123 -= (^(v123))
                            if false {
                                v125 := v4[741]
                                var v126 int16 = -668
                                _ = v125
                                _ = v126
                            }
                            _ = fi122
                            _ = v123
                            _ = v124
                        }
                    }
                case "foo bar":
                    switch v1 {
                    case -2147483648:
                        v2 = ((v2 + v2) + ((v2 + "foo bar") + (("a" + v2) + v2)))
                    default:
                        if false {
                            v2 = (v2 + "hello")
                            fmt.Println(fi116, v43)
                            if ((!true) && (4294967295 != 612)) {
                                for fi127 := 0; fi127 < 14; fi127++ {
                                    v128 := GenS13{F0: "foo bar", F1: 2147483647, F2: (fi116 | v1)}
                                    _ = fi127
                                    _ = v128
                                }
                                v117 = (v117 + 933)
                            } else {
                                for fi129 := 0; fi129 < 12; fi129++ {
                                    var v130 uint64 = 642
                                    _ = fi129
                                    _ = v130
                                }
                                _ = len(v4)
                                fmt.Println(v117, v117, v1)
                            }
                        }
                    }
                }
                _ = fi116
                _ = v117
            }
            for fi131 := 0; fi131 < 5; fi131++ {
                if (-2147483648 > 1) {
                    var v132 uint64 = 1
                    _ = v132
                } else {
                    switch v2 {
                    case "a":
                        var v133 uint = 65535
                        _ = v133
                    case "\n\t":
                        switch "foo bar" {
                        case "\n\t":
                            for fi134 := 0; fi134 < 4; fi134++ {
                                v135 := v4[2147483647]
                                _ = fi134
                                _ = v135
                            }
                        case "hello":
                            if true {
                                var v136 uint32 = 2147483647
                                _ = v136
                            } else {
                                if false {
                                    v43 += v43
                                }
                            }
                        case "":
                            _ = len(v4)
                        }
                    case "foo bar":
                        v2 = v2
                    default:
                        for fi137 := 0; fi137 < 14; fi137++ {
                            fmt.Println(fi131)
                            v43 += (-(v43))
                            var v138 float64 = 0.1
                            fmt.Println(v138)
                            _ = fi137
                            _ = v138
                        }
                    }
                }
                _ = fi131
            }
        }
    default:
        if (float64(0.0) < 3.14159) {
            var v139 uint8 = 34
            v43 -= 100
            _ = v139
        }
    }
    for fi140 := 0; fi140 < 17; fi140++ {
        v141 := v4[845]
        v2 = v2
        _ = fi140
        _ = v141
    }
    var v142 uint32 = 1
    fmt.Println(v142, v2)
    v3 = append(v3, v2)
    for fi143 := 0; fi143 < 19; fi143++ {
        for fi144 := 0; fi144 < 18; fi144++ {
            var v145 float32 = 1.5
            var v146 float64 = (3.14159 * 3.14159)
            if (787 <= 0) {
                v147 := GenS15{F0: 42, F1: 255}
                _ = v147
            }
            _ = fi144
            _ = v145
            _ = v146
        }
        _ = fi143
    }
    fmt.Println("v1:", v1)
    fmt.Println("v142:", v142)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v142
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v44
}
