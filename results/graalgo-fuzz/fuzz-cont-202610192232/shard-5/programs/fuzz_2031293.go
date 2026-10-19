package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 bool; F1 int64; F2 int32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint16; F1 uint8 }
type GenS10 struct { F0 int; F1 bool }
type GenS11 struct { F0 string; F1 uint8 }
type GenS12 struct { F0 int64; F1 int32 }
type GenS13 struct { F0 uint64 }

func main() {
    v1 := []int{32768}
    var v2 string = (("foo bar" + "hello") + "")
    var v3 uint8 = 42
    v4 := []string{v2, ("hello" + v2), "a"}
    var v5 int16 = 42
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi7
        }
        v2 = v2
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int32 = 42
            switch ((((v2 + "foo bar") + ((v2 + "") + v2)) + v2) + ((v2 + v2) + (v2 + "foo bar"))) {
            case "":
                var v10 float32 = 1e10
                _ = v10
            case "\n\t":
                var v11 uint8 = (v3 - 253)
                _ = v11
            case "foo bar":
                fmt.Println(fi8, v5, v2)
            default:
                var v12 int32 = (v9 - -2147483648)
                _ = v12
            }
            var v13 int8 = 0
            var v14 bool = (160 >= 246)
            v2 = ""
            _ = fi8
            _ = v13
            _ = v14
            _ = v9
        }
        _ = fi6
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        if ((0 != 0) && true) {
            if (false && (true || (!(2 != 0)))) {
                var v16 uint32 = 761
                var v17 bool = (4294967295 >= v16)
                switch ((v2 + (v2 + v2)) + v2) {
                case "\n\t":
                    v2 = ""
                case "hello":
                    for fi18 := 0; fi18 < 2; fi18++ {
                        if v17 {
                            var v19 int32 = 1
                            fmt.Println(v5, v2, fi18)
                            fmt.Println(fi18)
                            _ = v19
                        } else {
                            switch v2 {
                            case "\n\t":
                                fmt.Println(fi15)
                            case "hello":
                                v17 = false
                            case "a":
                                var v20 int32 = 405
                                _ = v20
                            default:
                                var v21 int16 = (v5 % ((v5) | 1))
                                _ = v21
                            }
                            v1 = append(v1, fi15)
                            v5 = (v5 * 42)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v22 uint64 = 54
                            _ = v22
                        }
                        _ = fi18
                    }
                case "foo bar":
                    v4 = append(v4, (v2 + "a"))
                default:
                    if v17 {
                        v2 = v2
                    }
                }
                _ = v16
                _ = v17
            }
            var v23 bool = false
            _ = v23
        } else {
            var v24 uint32 = 2147483647
            if (false && false) {
                var v25 float32 = 0.0
                _ = v25
            }
            if ((v24 + (v24 ^ 65535)) == (^(v24))) {
                var v26 string = v2
                v27 := GenS1{F0: 100, F1: v5}
                _ = v26
                _ = v27
            }
            _ = v24
        }
        if ((v5 == v5) || false) {
            for fi28 := 0; fi28 < 14; fi28++ {
                v4 = append(v4, ((v2 + v2) + (("\n\t" + "\n\t") + (v2 + "\n\t"))))
                var v29 uint64 = 1
                _ = fi28
                _ = v29
            }
            switch (fi15 % (((-(fi15))) | 1)) {
            case 2147483647:
                switch "\n\t" {
                case "hello":
                    switch "hello" {
                    case "a":
                        if false {
                            for fi30 := 0; fi30 < 11; fi30++ {
                                v5 = (v5 << ((65535) & 63))
                                v31 := GenS1{F0: -1, F1: 32767}
                                var v32 int32 = 1
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                            fmt.Println(v2, fi15)
                            fmt.Println(v3, fi15, v5)
                        } else {
                            v3 = v3
                            v3 = 39
                            for fi33 := 0; fi33 < 12; fi33++ {
                                if ((4294967295 > 0) && false) {
                                    var v34 uint8 = 42
                                    _ = v34
                                }
                                if (!false) {
                                    v4 = append(v4, "")
                                    v35 := GenS1{F0: 103, F1: v5}
                                    var v36 uint32 = 4294967295
                                    switch v2 {
                                    case "":
                                        if false {
                                            if (!true) {
                                                fmt.Println(v2, v5)
                                                v36 = v36
                                            } else {
                                                v2 = ((v2 + "a") + v2)
                                                var v37 int8 = 1
                                                if true {
                                                    if ((v36 + (v36 | v36)) == (v36 % (((^(v36))) | 1))) {
                                                        switch v2 {
                                                        case "":
                                                            v38 := GenS1{F0: -1, F1: -64}
                                                            _ = v38
                                                        }
                                                        var v39 float64 = -2.25
                                                        _ = v39
                                                    } else {
                                                        v40 := GenS2{F0: uint64(v3)}
                                                        _ = v40
                                                    }
                                                }
                                                _ = v37
                                            }
                                            var v41 int8 = int8(v5)
                                            var v42 int8 = v41
                                            _ = v41
                                            _ = v42
                                        } else {
                                            v43 := GenS2{F0: 0}
                                            _ = v43
                                        }
                                    case "a":
                                        if len(v1) > 0 {
                                            v1[0] = (fi33 % ((fi33) | 1))
                                        }
                                    default:
                                        switch (v2 + "\n\t") {
                                        case "foo bar":
                                            fmt.Println(v2, v5, v3)
                                        case "hello":
                                            var v44 string = v2
                                            _ = v44
                                        case "\n\t":
                                            var v45 int = (fi15 + (fi15 * (fi15 - fi33)))
                                            _ = v45
                                        }
                                    }
                                    _ = v35
                                    _ = v36
                                }
                                var v46 int32 = 250
                                v5 += v5
                                v46 -= (v46 >> ((556) & 63))
                                _ = fi33
                                _ = v46
                            }
                        }
                    default:
                        var v47 uint = 0
                        _ = v47
                    }
                default:
                    if (((42 >= -854) && ((!false) || (514 < 397))) && (uint64(v3) >= 71)) {
                        v48 := GenS1{F0: -1, F1: v5}
                        _ = v48
                    }
                }
            case 783:
                for fi49 := 0; fi49 < 9; fi49++ {
                    var v50 uint64 = uint64(v3)
                    fmt.Println(v5)
                    var v51 int16 = v5
                    var v52 int8 = 127
                    _ = fi49
                    _ = v50
                    _ = v51
                    _ = v52
                }
            case 127:
                for fi53 := 0; fi53 < 5; fi53++ {
                    if true {
                        v54 := GenS3{F0: 1}
                        _ = len(v4)
                        _ = v54
                    } else {
                        fmt.Println(fi53)
                    }
                    if (4294967295 < 3) {
                        var v55 int32 = 2
                        _ = v55
                    }
                    _ = fi53
                }
            default:
                for fi56 := 0; fi56 < 12; fi56++ {
                    _ = len(v4)
                    switch 32767 {
                    case -9223372036854775808:
                        v57 := GenS1{F0: -128, F1: (v5 >> ((424) & 63))}
                        _ = v57
                    case 4294967295:
                        switch fi15 {
                        case 2147483647:
                            var v58 bool = (65535 <= 892)
                            switch (fi56 | fi15) {
                            case 42:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case -642:
                                v59 := GenS3{F0: 100}
                                _ = v59
                            case 2147483647:
                                v2 = v2
                            }
                            v2 = v2
                            _ = v58
                        case -873:
                            switch "\n\t" {
                            case "foo bar":
                                v2 = "hello"
                            case "a":
                                v3 = v3
                            case "":
                                switch v2 {
                                case "foo bar":
                                    v2 = ("\n\t" + v2)
                                case "\n\t":
                                    fmt.Println(fi15, v5)
                                case "hello":
                                    if false {
                                        if (fi56 != -1) {
                                            var v60 int16 = v5
                                            v3 -= v3
                                            fmt.Println(v60, fi15)
                                            _ = v60
                                        }
                                        if (false && (false && (!(2147483647 <= int32(fi56))))) {
                                            var v61 uint16 = 255
                                            _ = v61
                                        }
                                    }
                                    v3 = (^(v3))
                                default:
                                    var v62 uint32 = 630
                                    _ = v62
                                }
                            }
                            v63 := GenS4{F0: v3}
                            if (!true) {
                                var v64 float32 = 0.1
                                v64 = (v64 + (v64 - 1e10))
                                _ = v64
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v63
                        }
                    case 127:
                        v65 := GenS3{F0: -73}
                        _ = v65
                    default:
                        switch (v2 + v2) {
                        case "foo bar":
                            v2 = ((v2 + v2) + "\n\t")
                        case "":
                            if (v3 == (v3 - v3)) {
                                var v66 string = v2
                                var v67 uint = 72
                                _ = v66
                                _ = v67
                            }
                        default:
                            v4 = append(v4, v2)
                        }
                    }
                    _ = fi56
                }
            }
            v5 = (v5 & -1)
        } else {
            v5 += (v5 >> ((0) & 63))
            v5 -= (v5 | 1)
        }
        if false {
            var v68 int32 = -2147483648
            v69 := GenS5{F0: ((65535 > 849) && (!true))}
            _ = v68
            _ = v69
        }
        _ = fi15
    }
    switch 816 {
    case 127:
        if true {
            var v70 int8 = 100
            for fi71 := 0; fi71 < 12; fi71++ {
                v72 := GenS6{F0: (^(fi71)), F1: (true || (0.0 >= 1e10))}
                _ = fi71
                _ = v72
            }
            switch 127 {
            case 2147483647:
                _ = len(v1)
            case 32768:
                fmt.Println(v2, v70, v2)
            case 4294967295:
                fmt.Println(v3)
            }
            _ = v70
        } else {
            if false {
                v73 := GenS5{F0: (9223372036854775807 <= -1)}
                _ = v73
            } else {
                switch 32767 {
                case 42:
                    var v74 uint64 = 0
                    _ = v74
                case 818:
                    for fi75 := 0; fi75 < 8; fi75++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        for fi76 := 0; fi76 < 13; fi76++ {
                            v2 = ((v2 + "a") + v2)
                            v5 = v5
                            _ = fi76
                        }
                        _ = fi75
                    }
                }
            }
            v3 -= 128
            fmt.Println(v2)
        }
    default:
        var v77 int = 127
        _ = v77
    }
    v3 -= 93
    v3 += (v3 & 139)
    switch "" {
    case "a":
        var v78 int32 = 200
        _ = v78
    }
    var v79 int8 = 0
    v3 -= (^(v3))
    if (1 < 255) {
        for fi80 := 0; fi80 < 5; fi80++ {
            for fi81 := 0; fi81 < 5; fi81++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if (((3.14159 * (1e10 - (0.1 * (3.14159 + 1.5)))) - 1e10) == (1.5 + (1.5 + -2.25))) {
                    if true {
                        fmt.Println(v79, v79, fi81)
                        var v82 uint32 = 2147483647
                        if (true && ((-544 <= 0) && ((-2147483648 < -1) || false))) {
                            v83 := GenS4{F0: v3}
                            v82 = v82
                            v82 -= (^(v82))
                            _ = v83
                        } else {
                            var v84 float64 = 1e10
                            v3 -= 255
                            _ = v84
                        }
                        v85 := GenS5{F0: false}
                        _ = v82
                        _ = v85
                    }
                } else {
                    v86 := GenS2{F0: 4294967295}
                    _ = v86
                }
                if true {
                    v2 = v2
                    v5 += (v5 << ((614) & 63))
                    if false {
                        v2 = "foo bar"
                        if len(v1) > 0 {
                            v1[0] = fi80
                        }
                        if len(v1) > 0 {
                            v1[0] = -274
                        }
                        var v87 uint64 = 0
                        _ = v87
                    } else {
                        v5 += (v5 * (v5 << ((156) & 63)))
                        v88 := GenS7{F0: false, F1: -89, F2: -622}
                        if (707 < 718) {
                            v3 = v3
                        }
                        _ = v88
                    }
                    switch (v2 + v2) {
                    case "\n\t":
                        var v89 uint64 = 976
                        _ = v89
                    case "foo bar":
                        if (false && true) {
                            var v90 float64 = 0.1
                            _ = v90
                        }
                    default:
                        var v91 int = (fi80 | fi81)
                        _ = v91
                    }
                    v1 = append(v1, -666)
                }
                _ = fi81
            }
            v92 := GenS7{F0: ((fi80 * fi80) != 535), F1: 1, F2: int32(v3)}
            var v93 float64 = -2.25
            _ = fi80
            _ = v92
            _ = v93
        }
    }
    switch 9223372036854775807 {
    case -2147483648:
        v79 += int8(v3)
    case 32768:
        v79 = v79
    case 435:
        switch -2147483648 {
        case -1:
            for fi94 := 0; fi94 < 3; fi94++ {
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v2 = "\n\t"
                    v5 = 787
                }
                if true {
                    if false {
                        var v95 uint16 = 0
                        _ = v95
                    }
                } else {
                    var v96 int = (fi94 - fi94)
                    _ = v96
                }
                _ = fi94
            }
        case 4294967295:
            if (!true) {
                for fi97 := 0; fi97 < 13; fi97++ {
                    v98 := GenS4{F0: v3}
                    _ = fi97
                    _ = v98
                }
                if false {
                    v3 = 2
                }
            } else {
                for fi99 := 0; fi99 < 1; fi99++ {
                    if (fi99 < 1) {
                        var v100 float32 = (3.14159 / 1.0)
                        if (1 != uint16(v3)) {
                            if ((!false) && true) {
                                v101 := GenS8{F0: 129}
                                fmt.Println(v3)
                                _ = v101
                            } else {
                                if ((!false) || (true || ((v3 ^ v3) > 83))) {
                                    switch (fi99 % (((fi99 ^ (fi99 ^ fi99))) | 1)) {
                                    case 670:
                                        if (!((!(((false && (!(!false))) || false) && (fi99 > 32768))) || (563 <= 1))) {
                                            for fi102 := 0; fi102 < 20; fi102++ {
                                                v79 = v79
                                                _ = fi102
                                            }
                                            if (!true) {
                                                if (0.1 < -2.25) {
                                                    for fi103 := 0; fi103 < 6; fi103++ {
                                                        v104 := GenS8{F0: 42}
                                                        var v105 uint64 = 170
                                                        v106 := GenS9{F0: 200, F1: v3}
                                                        _ = fi103
                                                        _ = v104
                                                        _ = v105
                                                        _ = v106
                                                    }
                                                } else {
                                                    v2 = "foo bar"
                                                    for fi107 := 0; fi107 < 16; fi107++ {
                                                        var v108 int64 = -1
                                                        _ = fi107
                                                        _ = v108
                                                    }
                                                    fmt.Println(v5)
                                                }
                                            }
                                            var v109 uint16 = 2
                                            _ = v109
                                        } else {
                                            if (42 != 288) {
                                                fmt.Println(v2)
                                                v79 += -128
                                            } else {
                                                var v110 uint16 = 132
                                                if (-331 < 0) {
                                                    if true {
                                                        for fi111 := 0; fi111 < 15; fi111++ {
                                                            if false {
                                                                if false {
                                                                    v110 = 0
                                                                    if (!false) {
                                                                        v112 := GenS1{F0: (-(v79)), F1: 803}
                                                                        v1 = append(v1, 32768)
                                                                        var v113 uint16 = (v110 >> ((65535) & 63))
                                                                        var v114 float64 = 3.14159
                                                                        _ = v112
                                                                        _ = v113
                                                                        _ = v114
                                                                    } else {
                                                                        var v115 uint64 = 279
                                                                        _ = v115
                                                                    }
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    fmt.Println(fi99)
                                                                }
                                                                v116 := GenS10{F0: (^(fi111)), F1: false}
                                                                _ = v116
                                                            }
                                                            _ = fi111
                                                        }
                                                        if false {
                                                            v110 += (^(v110))
                                                            v117 := GenS6{F0: (^(fi99)), F1: (1 >= 0)}
                                                            fmt.Println(fi99)
                                                            _ = v117
                                                        } else {
                                                            var v118 uint16 = v110
                                                            _ = v118
                                                        }
                                                    } else {
                                                        var v119 uint32 = 0
                                                        _ = v119
                                                    }
                                                } else {
                                                    for fi120 := 0; fi120 < 2; fi120++ {
                                                        if ((^(v3)) != 0) {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        } else {
                                                            var v121 uint16 = 1
                                                            switch 2147483647 {
                                                            case 2147483647:
                                                                v122 := GenS1{F0: (^(v79)), F1: v5}
                                                                _ = v122
                                                            default:
                                                                v2 = (v2 + v2)
                                                            }
                                                            switch (v2 + "") {
                                                            case "\n\t":
                                                                fmt.Println(v100, fi120)
                                                            case "a":
                                                                var v123 uint = 18446744073709551615
                                                                _ = v123
                                                            }
                                                            _ = v121
                                                        }
                                                        var v124 uint = 479
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        _ = fi120
                                                        _ = v124
                                                    }
                                                }
                                                v100 -= v100
                                                _ = v110
                                            }
                                            switch fi99 {
                                            case 286:
                                                var v125 int16 = 1
                                                _ = v125
                                            default:
                                                var v126 uint8 = 0
                                                _ = v126
                                            }
                                        }
                                    default:
                                        for fi127 := 0; fi127 < 8; fi127++ {
                                            v5 -= (v5 ^ -1)
                                            v79 = (v79 >> ((0) & 63))
                                            _ = fi127
                                        }
                                    }
                                } else {
                                    switch (fi99 | int(v3)) {
                                    case 895:
                                        if (false || ((v5 & (v5 / (((-(v5))) | 1))) < v5)) {
                                            v2 = ((v2 + "hello") + ((v2 + "hello") + ""))
                                            fmt.Println(v79, v79, v5)
                                        } else {
                                            v1 = append(v1, 2147483647)
                                        }
                                    case 9223372036854775807:
                                        v5 = v5
                                    case -2147483648:
                                        if len(v4) > 0 {
                                            v4[0] = "\n\t"
                                        }
                                    }
                                    if true {
                                        v128 := GenS8{F0: 995}
                                        switch 127 {
                                        case 507:
                                            _ = len(v4)
                                        }
                                        _ = v128
                                    } else {
                                        for fi129 := 0; fi129 < 1; fi129++ {
                                            fmt.Println(v100, v5, v3)
                                            var v130 uint8 = v3
                                            _ = fi129
                                            _ = v130
                                        }
                                    }
                                }
                                v100 += (v100 / (v100 / v100))
                                v3 = 127
                            }
                        } else {
                            for fi131 := 0; fi131 < 19; fi131++ {
                                if (!false) {
                                    v132 := GenS2{F0: 4294967295}
                                    var v133 uint16 = 65535
                                    if true {
                                        var v134 int8 = 42
                                        if (false || false) {
                                            v4 = append(v4, v2)
                                            var v135 float32 = (0.1 / v100)
                                            if true {
                                                var v136 float64 = 0.1
                                                if false {
                                                    if true {
                                                        v5 -= (v5 | (v5 - -32768))
                                                    }
                                                    fmt.Println(v134)
                                                    _ = len(v4)
                                                } else {
                                                    v136 += v136
                                                }
                                                _ = v136
                                            }
                                            v134 += (v79 & (v79 / ((100) | 1)))
                                            _ = v135
                                        } else {
                                            var v137 int = fi131
                                            _ = v137
                                        }
                                        if (42 >= -9223372036854775808) {
                                            v5 += 1
                                            v2 = ("a" + ("" + "hello"))
                                            if len(v4) > 0 {
                                                v4[0] = "foo bar"
                                            }
                                        }
                                        _ = v134
                                    }
                                    v138 := GenS10{F0: fi131, F1: (!false)}
                                    _ = v132
                                    _ = v133
                                    _ = v138
                                } else {
                                    v79 = 127
                                    v139 := GenS9{F0: 1, F1: (v3 - (^(v3)))}
                                    _ = v139
                                }
                                v140 := GenS4{F0: v3}
                                _ = fi131
                                _ = v140
                            }
                            for fi141 := 0; fi141 < 6; fi141++ {
                                if true {
                                    fmt.Println(v79, v2, v100)
                                    fmt.Println(v5, v79)
                                    switch v2 {
                                    case "foo bar":
                                        var v142 int16 = v5
                                        _ = v142
                                    case "hello":
                                        var v143 uint8 = (^(v3))
                                        _ = v143
                                    default:
                                        if true {
                                            v100 = v100
                                            fmt.Println(v100, v2)
                                            if (-1 > 42) {
                                                if (177 >= 116) {
                                                    switch fi141 {
                                                    case 42:
                                                        fmt.Println(fi99)
                                                    case 9223372036854775807:
                                                        if (647 <= 2) {
                                                            fmt.Println(v3)
                                                            v144 := GenS11{F0: v2, F1: (^(v3))}
                                                            _ = v144
                                                        }
                                                    }
                                                } else {
                                                    v145 := GenS12{F0: -406, F1: 42}
                                                    fmt.Println(v79, v2)
                                                    if (635 <= 2) {
                                                        v5 += 32767
                                                    } else {
                                                        v5 -= v5
                                                        v2 = ("\n\t" + "\n\t")
                                                    }
                                                    _ = v145
                                                }
                                                var v146 int8 = -1
                                                var v147 uint8 = 1
                                                _ = v146
                                                _ = v147
                                            }
                                            var v148 string = "foo bar"
                                            switch v148 {
                                            case "":
                                                fmt.Println(fi141, v3)
                                            default:
                                                var v149 int = (-(fi141))
                                                _ = v149
                                            }
                                            v5 -= (v5 * v5)
                                            _ = v148
                                        }
                                    }
                                }
                                var v150 bool = (((!(int(v3) == fi141)) || false) || false)
                                fmt.Println(v79, v5)
                                v151 := GenS13{F0: 834}
                                _ = fi141
                                _ = v150
                                _ = v151
                            }
                            switch (fi99 - (fi99 + fi99)) {
                            case -131:
                                v152 := GenS2{F0: 1}
                                _ = v152
                            case 32768:
                                v5 -= 32767
                            default:
                                for fi153 := 0; fi153 < 14; fi153++ {
                                    fmt.Println(v79, v5)
                                    v79 = v79
                                    var v154 int16 = int16(v3)
                                    _ = fi153
                                    _ = v154
                                }
                            }
                        }
                        if (2 == 2147483648) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v155 := GenS13{F0: 2}
                            _ = v155
                        } else {
                            if false {
                                v3 += v3
                                if (v2 <= v2) {
                                    var v156 int8 = 42
                                    _ = v156
                                }
                            }
                            if (-644 > -861) {
                                v100 = ((v100 * ((3.14159 * v100) * (v100 * (float32(-2.25) * ((0.0 / float32(-2.25)) / (0.1 + v100)))))) * ((3.14159 - 0.1) / v100))
                                v100 -= (-2.25 - 1e10)
                            }
                        }
                        _ = v100
                    }
                    v79 = (v79 / ((0) | 1))
                    _ = fi99
                }
                v79 = v79
            }
        }
    default:
        v157 := GenS9{F0: 42, F1: v3}
        _ = v157
    }
    var v158 uint = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v158:", v158)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v158
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v79
}
