package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 bool }
type GenS2 struct { F0 int8; F1 uint32; F2 int }
type GenS3 struct { F0 int; F1 uint; F2 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool; F1 bool }
type GenS9 struct { F0 int; F1 int64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int; F1 int16; F2 bool }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint32; F1 int }
type GenS16 struct { F0 int64; F1 uint16; F2 int64 }
type GenS17 struct { F0 int32; F1 int32 }
type GenS18 struct { F0 string; F1 uint32; F2 uint }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 bool }

func main() {
    v1 := []int{9223372036854775807}
    v2 := make(map[uint32]int)
    var v3 string = ("\n\t" + "hello")
    var v4 uint32 = 79
    if (true && false) {
        v5 := GenS1{F0: 4294967295, F1: 1, F2: true}
        v4 += v4
        _ = v5
    } else {
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS1{F0: (v4 >> ((65535) & 63)), F1: 389, F2: true}
            _ = fi6
            _ = v7
        }
        v4 = (v4 / (((v4 / (((v4 - 358)) | 1))) | 1))
        var v8 int32 = -2147483648
        _ = v8
    }
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    v9 := GenS1{F0: v4, F1: 525, F2: true}
    for fi10 := 0; fi10 < 9; fi10++ {
        fmt.Println(v3)
        _ = fi10
    }
    if false {
        if (!true) {
            switch "" {
            case "foo bar":
                fmt.Println(v4)
            case "hello":
                for fi11 := 0; fi11 < 14; fi11++ {
                    var v12 int8 = 17
                    if (("a" >= "hello") || true) {
                        if false {
                            v3 = (("" + "\n\t") + "hello")
                        }
                        for fi13 := 0; fi13 < 4; fi13++ {
                            v12 = v12
                            switch (fi13 & 32767) {
                            case -884:
                                _ = len(v1)
                            case 32767:
                                switch 42 {
                                case 127:
                                    var v14 uint64 = 1
                                    _ = v14
                                case 0:
                                    var v15 int32 = -1
                                    _ = v15
                                default:
                                    var v16 int16 = 124
                                    _ = v16
                                }
                            default:
                                switch v3 {
                                case "":
                                    v3 = v3
                                case "foo bar":
                                    _ = len(v2)
                                case "a":
                                    var v17 int = (^(fi13))
                                    _ = v17
                                default:
                                    v18 := GenS1{F0: v4, F1: -2147483648, F2: (-1 >= 321)}
                                    _ = v18
                                }
                            }
                            fmt.Println(fi13)
                            v12 = (v12 - 42)
                            if true {
                                fmt.Println(fi11)
                            } else {
                                switch "foo bar" {
                                case "a":
                                    switch 42 {
                                    case 9223372036854775807:
                                        var v19 int8 = (-(v12))
                                        _ = v19
                                    default:
                                        var v20 int64 = 42
                                        _ = v20
                                    }
                                case "hello":
                                    if ((fi11 | 128) >= (fi11 | fi13)) {
                                        var v21 int64 = -1
                                        if true {
                                            if (!false) {
                                                if true {
                                                    if false {
                                                        v12 = (v12 >> ((944) & 63))
                                                    } else {
                                                        var v22 int64 = v21
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        _ = v22
                                                    }
                                                    fmt.Println(fi13, v12, fi13)
                                                }
                                            } else {
                                                fmt.Println(v21, v21, fi13)
                                            }
                                        } else {
                                            if (v4 >= v4) {
                                                fmt.Println(fi13, v21)
                                                var v23 float32 = 0.1
                                                if false {
                                                    var v24 uint8 = 127
                                                    v25 := GenS1{F0: v4, F1: 42, F2: ((v4 - v4) < v4)}
                                                    if (!((v23 / (0.1 / (-2.25 * 3.14159))) >= 3.14159)) {
                                                        var v26 float32 = float32((0.1 / 1.5))
                                                        v27 := GenS1{F0: 2147483648, F1: -987, F2: true}
                                                        switch fi11 {
                                                        case -1:
                                                            switch 371 {
                                                            case 1:
                                                                var v28 string = (v3 + "a")
                                                                if (1000 >= 1) {
                                                                    if (1 <= 694) {
                                                                        delete(v2, v4)
                                                                        var v29 uint64 = 263
                                                                        _ = v29
                                                                    }
                                                                    v4 -= (v4 * 359)
                                                                }
                                                                v3 = (v28 + ("foo bar" + (((v3 + "a") + (("hello" + v3) + v3)) + (v3 + v3))))
                                                                _ = v28
                                                            case -641:
                                                                v30 := GenS2{F0: v12, F1: uint32(v24), F2: -65}
                                                                var v31 float32 = ((v26 - (0.1 + v23)) - (v26 + 1e10))
                                                                _ = v30
                                                                _ = v31
                                                            case 952:
                                                                v21 -= 9223372036854775807
                                                            default:
                                                                var v32 bool = ((v4 + v4) >= 42)
                                                                _ = v32
                                                            }
                                                        case 2147483647:
                                                            v2[4294967295] = (-(fi13))
                                                            var v33 float64 = 1.5
                                                            _ = v33
                                                        case 411:
                                                            v23 = (3.14159 * (1.5 * (v26 * v26)))
                                                        }
                                                        _ = v26
                                                        _ = v27
                                                    } else {
                                                        if ((^(v24)) > v24) {
                                                            switch -9223372036854775808 {
                                                            case 4294967295:
                                                                _ = len(v1)
                                                            default:
                                                                fmt.Println(v12, v12)
                                                            }
                                                            v23 = 0.0
                                                            fmt.Println(v4)
                                                            var v34 uint = 48
                                                            _ = v34
                                                        }
                                                        var v35 uint16 = 833
                                                        _ = v35
                                                    }
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v36 uint32 = (^(v4))
                                                        _ = v36
                                                    case "hello":
                                                        v4 = (v4 & (v4 << ((1) & 63)))
                                                    default:
                                                        fmt.Println(v3, v3, fi13)
                                                    }
                                                    if true {
                                                        v12 += (v12 + (v12 / ((int8(v4)) | 1)))
                                                        var v37 bool = true
                                                        v12 += v12
                                                        _ = v37
                                                    }
                                                    var v38 uint8 = 1
                                                    _ = v24
                                                    _ = v25
                                                    _ = v38
                                                }
                                                _ = v23
                                            } else {
                                                if true {
                                                    v12 = 1
                                                }
                                                v4 += v4
                                                var v39 float64 = -2.25
                                                switch v3 {
                                                case "a":
                                                    v2[v4] = (-(fi11))
                                                case "foo bar":
                                                    if (((fi13 ^ (fi13 << ((2) & 63))) >= 42) && (4294967295 >= fi11)) {
                                                        delete(v2, v4)
                                                        v40 := GenS1{F0: 4294967295, F1: 42, F2: ((false || false) && false)}
                                                        if (((v4 > v4) || false) || true) {
                                                            if (833 >= 1) {
                                                                v3 = "\n\t"
                                                                var v41 int64 = v21
                                                                _ = v41
                                                            } else {
                                                                if false {
                                                                    v42 := GenS2{F0: v12, F1: v4, F2: fi13}
                                                                    v39 += v39
                                                                    _ = v42
                                                                } else {
                                                                    v43 := GenS2{F0: 0, F1: v4, F2: fi13}
                                                                    _ = v43
                                                                }
                                                                v12 -= (v12 & 100)
                                                                if false {
                                                                    fmt.Println(v21)
                                                                    if false {
                                                                        var v44 uint32 = v4
                                                                        v2[4294967295] = fi13
                                                                        switch "" {
                                                                        case "hello":
                                                                            var v45 int16 = -32768
                                                                            _ = v45
                                                                        case "\n\t":
                                                                            fmt.Println(v3, v39, v4)
                                                                        case "a":
                                                                            if true {
                                                                                _ = len(v2)
                                                                            }
                                                                        }
                                                                        if false {
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 1
                                                                            }
                                                                            v12 = (v12 % ((v12) | 1))
                                                                            v46 := GenS1{F0: (v44 | v44), F1: -1, F2: (true && true)}
                                                                            _ = v46
                                                                        }
                                                                        _ = v44
                                                                    }
                                                                    var v47 int64 = v21
                                                                    _ = v47
                                                                } else {
                                                                    v21 = v21
                                                                    var v48 int = 2147483647
                                                                    if (true && (true || (true || false))) {
                                                                        var v49 int = fi11
                                                                        _ = v49
                                                                    }
                                                                    switch v3 {
                                                                    case "\n\t":
                                                                        var v50 float64 = 0.1
                                                                        _ = v50
                                                                    }
                                                                    var v51 int = 1
                                                                    var v52 float64 = 0.0
                                                                    _ = v48
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                            }
                                                        } else {
                                                            v21 -= v21
                                                            var v53 uint16 = 571
                                                            v3 = ((v3 + (v3 + "a")) + v3)
                                                            _ = v53
                                                        }
                                                        _ = v40
                                                    } else {
                                                        v4 -= 2
                                                    }
                                                default:
                                                    v54 := v2[2147483648]
                                                    _ = v54
                                                }
                                                if true {
                                                    var v55 float64 = v39
                                                    if true {
                                                        if (!false) {
                                                            var v56 float32 = (((3.14159 / float32(3.14159)) / 3.14159) / 1e10)
                                                            if (true || false) {
                                                                if (false || true) {
                                                                    v57 := GenS3{F0: fi13, F1: 663, F2: ((v3 + v3) + v3)}
                                                                    switch v3 {
                                                                    case "":
                                                                        v4 = v4
                                                                    case "foo bar":
                                                                        if true {
                                                                            var v58 bool = (fi13 <= (-(fi11)))
                                                                            _ = v58
                                                                        }
                                                                    case "hello":
                                                                        var v59 uint = 1
                                                                        switch -2147483648 {
                                                                        case 32768:
                                                                            var v60 uint8 = 128
                                                                            _ = v60
                                                                        case 32767:
                                                                            v21 -= v21
                                                                        default:
                                                                            fmt.Println(v12, v39)
                                                                        }
                                                                        _ = len(v1)
                                                                        v61 := GenS4{F0: ((v3 + "a") + "\n\t")}
                                                                        _ = v59
                                                                        _ = v61
                                                                    }
                                                                    if false {
                                                                        switch 0 {
                                                                        case 4294967295:
                                                                            var v62 int8 = (v12 * v12)
                                                                            _ = v62
                                                                        case 42:
                                                                            switch (fi11 & fi11) {
                                                                            case -2147483648:
                                                                                v1 = append(v1, fi13)
                                                                            }
                                                                        default:
                                                                            v63 := GenS1{F0: (v4 * v4), F1: 1, F2: true}
                                                                            _ = v63
                                                                        }
                                                                        if (!(894 == 2)) {
                                                                            var v64 int = (fi11 | (fi13 - (fi13 - fi11)))
                                                                            v65 := GenS5{F0: v3, F1: v21}
                                                                            delete(v2, (v4 / ((v4) | 1)))
                                                                            var v66 int32 = 1
                                                                            var v67 uint32 = (v4 ^ v4)
                                                                            _ = v64
                                                                            _ = v65
                                                                            _ = v66
                                                                            _ = v67
                                                                        }
                                                                        if true {
                                                                            if false {
                                                                                delete(v2, v4)
                                                                                v68 := GenS6{F0: (2 < 465)}
                                                                                _ = v68
                                                                            }
                                                                            v69 := GenS6{F0: true}
                                                                            _ = v69
                                                                        }
                                                                    }
                                                                    _ = v57
                                                                }
                                                                v70 := GenS4{F0: (v3 + v3)}
                                                                var v71 bool = ((fi13 << ((2) & 63)) > (fi11 | (fi11 % (((^(fi13))) | 1))))
                                                                _ = v70
                                                                _ = v71
                                                            } else {
                                                                v3 = v3
                                                                switch fi11 {
                                                                case -2147483648:
                                                                    v72 := GenS7{F0: 1}
                                                                    _ = v72
                                                                case 4294967295:
                                                                    v39 = (1e10 + 1.5)
                                                                case 127:
                                                                    if false {
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        v39 -= float64((v56 * v56))
                                                                    } else {
                                                                        fmt.Println(fi11, fi13)
                                                                        v4 -= 826
                                                                    }
                                                                default:
                                                                    v73 := GenS3{F0: 9223372036854775807, F1: 42, F2: v3}
                                                                    _ = v73
                                                                }
                                                                var v74 float32 = -2.25
                                                                _ = v74
                                                            }
                                                            v75 := GenS8{F0: true, F1: true}
                                                            v1 = append(v1, fi11)
                                                            _ = v56
                                                            _ = v75
                                                        }
                                                        var v76 uint32 = (v4 | v4)
                                                        switch (v3 + "foo bar") {
                                                        case "hello":
                                                            if len(v1) > 0 {
                                                                v1[0] = 551
                                                            }
                                                        case "\n\t":
                                                            if false {
                                                                var v77 bool = false
                                                                v12 -= (v12 ^ (v12 >> ((821) & 63)))
                                                                var v78 bool = v77
                                                                _ = v77
                                                                _ = v78
                                                            }
                                                        }
                                                        _ = v76
                                                    }
                                                    _ = v55
                                                } else {
                                                    var v79 int16 = -1
                                                    v12 = (v12 / ((-128) | 1))
                                                    v79 -= (v79 + 42)
                                                    switch (v3 + ("foo bar" + v3)) {
                                                    case "foo bar":
                                                        v80 := v2[(^(v4))]
                                                        _ = v80
                                                    case "":
                                                        switch "hello" {
                                                        case "hello":
                                                            v81 := GenS2{F0: v12, F1: v4, F2: (fi11 << ((2) & 63))}
                                                            _ = v81
                                                        }
                                                    case "hello":
                                                        v21 = v21
                                                    }
                                                    if (906 <= 28) {
                                                        var v82 uint64 = 97
                                                        _ = v82
                                                    } else {
                                                        v79 -= (v79 ^ (v79 % ((int16(v21)) | 1)))
                                                        v3 = v3
                                                    }
                                                    _ = v79
                                                }
                                                _ = v39
                                            }
                                            var v83 int64 = 1
                                            _ = v83
                                        }
                                        delete(v2, v4)
                                        _ = v21
                                    } else {
                                        fmt.Println(v12, v4, v3)
                                        if (255 <= 0) {
                                            _ = len(v2)
                                            var v84 string = "a"
                                            v85 := v2[v4]
                                            _ = v84
                                            _ = v85
                                        }
                                    }
                                }
                            }
                            _ = fi13
                        }
                    } else {
                        v86 := GenS5{F0: "", F1: -1}
                        if true {
                            v3 = ((v3 + ("" + "\n\t")) + ("\n\t" + v3))
                        } else {
                            v4 -= v4
                            var v87 int32 = 50
                            _ = v87
                        }
                        if (32767 != -1) {
                            var v88 int64 = 42
                            if ((!(!false)) || ((42 != 1000) && (false && true))) {
                                v12 = v12
                                var v89 string = v3
                                _ = v89
                            } else {
                                v4 += v4
                                v88 -= (^(v88))
                            }
                            delete(v2, 1)
                            _ = v88
                        }
                        _ = v86
                    }
                    _ = fi11
                    _ = v12
                }
            default:
                var v90 string = v3
                _ = v90
            }
        } else {
            var v91 uint = 4294967295
            _ = v91
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi92 := 0; fi92 < 17; fi92++ {
        if true {
            for fi93 := 0; fi93 < 1; fi93++ {
                var v94 string = v3
                if true {
                    v4 += (^(v4))
                }
                _ = fi93
                _ = v94
            }
            var v95 uint32 = v4
            for fi96 := 0; fi96 < 14; fi96++ {
                v3 = ("\n\t" + v3)
                _ = fi96
            }
            _ = v95
        }
        v97 := GenS9{F0: fi92, F1: -197}
        var v98 int = fi92
        _ = fi92
        _ = v97
        _ = v98
    }
    switch "\n\t" {
    case "hello":
        v4 = v4
    }
    fmt.Println(v4, v4)
    v4 -= v4
    switch v3 {
    case "a":
        delete(v2, v4)
    case "\n\t":
        if ("\n\t" != (v3 + v3)) {
            var v99 uint = 65535
            v100 := GenS9{F0: 32767, F1: 9223372036854775807}
            v4 -= v4
            _ = v100
            _ = v99
        }
    case "hello":
        delete(v2, (^(v4)))
    }
    switch 370 {
    case 392:
        var v101 int64 = 516
        _ = v101
    default:
        if (!true) {
            var v102 int16 = -206
            for fi103 := 0; fi103 < 6; fi103++ {
                if false {
                    if (!((-(fi103)) <= (fi103 * 1))) {
                        switch fi103 {
                        case 32767:
                            if ((((((0.1 / -2.25) - (1.5 / (-2.25 - 0.0))) + 1.5) * -2.25) / 0.1) == -2.25) {
                                switch (("a" + ((("a" + (((("hello" + (v3 + "\n\t")) + v3) + v3) + "\n\t")) + "foo bar") + ("" + "hello"))) + (v3 + "hello")) {
                                case "\n\t":
                                    delete(v2, 0)
                                case "hello":
                                    if (!(true && (!false))) {
                                        for fi104 := 0; fi104 < 16; fi104++ {
                                            v105 := GenS3{F0: fi103, F1: 0, F2: v3}
                                            var v106 uint64 = 0
                                            _ = fi104
                                            _ = v105
                                            _ = v106
                                        }
                                        v4 = v4
                                        var v107 int32 = 0
                                        _ = v107
                                    }
                                case "foo bar":
                                    for fi108 := 0; fi108 < 8; fi108++ {
                                        v109 := GenS9{F0: (fi103 << ((1) & 63)), F1: -1}
                                        _ = fi108
                                        _ = v109
                                    }
                                default:
                                    var v110 uint32 = (^(v4))
                                    _ = v110
                                }
                                if true {
                                    var v111 uint16 = 32767
                                    v4 = v4
                                    _ = v111
                                } else {
                                    v102 += (^(v102))
                                    var v112 int = 0
                                    _ = v112
                                }
                                var v113 uint32 = (v4 ^ v4)
                                _ = v113
                            }
                        default:
                            v3 = "hello"
                        }
                        for fi114 := 0; fi114 < 3; fi114++ {
                            v4 -= (v4 ^ v4)
                            if len(v1) > 0 {
                                v1[0] = (fi103 >> ((2) & 63))
                            }
                            _ = fi114
                        }
                        if false {
                            if false {
                                v115 := GenS5{F0: v3, F1: -1}
                                _ = len(v1)
                                _ = v115
                            }
                            v3 = ("hello" + v3)
                        } else {
                            _ = len(v1)
                        }
                    } else {
                        v116 := GenS10{F0: 2}
                        if (1 == 1) {
                            for fi117 := 0; fi117 < 1; fi117++ {
                                switch ((v3 + v3) + ("a" + ("" + "a"))) {
                                case "a":
                                    v4 = v4
                                case "hello":
                                    if false {
                                        v2[v4] = (^(fi103))
                                        var v118 uint = 484
                                        _ = v118
                                    } else {
                                        delete(v2, 2147483647)
                                        _ = len(v1)
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = 9223372036854775807
                                    }
                                default:
                                    v119 := GenS11{F0: 2}
                                    _ = v119
                                }
                                _ = fi117
                            }
                            fmt.Println(v3, v102, fi103)
                            for fi120 := 0; fi120 < 17; fi120++ {
                                v4 = 2147483648
                                var v121 string = v3
                                var v122 int64 = 0
                                if true {
                                    v122 = (v122 ^ v122)
                                } else {
                                    v123 := GenS12{F0: 127}
                                    if (v122 < (v122 % ((v122) | 1))) {
                                        v1 = append(v1, fi120)
                                    } else {
                                        fmt.Println(v3, v102)
                                    }
                                    switch v3 {
                                    case "foo bar":
                                        v122 = v122
                                    case "hello":
                                        fmt.Println(v121)
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = (^(fi103))
                                    }
                                    fmt.Println(v122, v102, v121)
                                    v4 = (v4 - (^(v4)))
                                    _ = v123
                                }
                                v124 := GenS13{F0: fi120, F1: (v102 << ((373) & 63)), F2: (!(((((float32(3.14159) / (1.5 - float32(0.1))) >= 3.14159) && (2 > 42)) && true) && false))}
                                v125 := GenS9{F0: fi120, F1: (v122 / (((v122 | v122)) | 1))}
                                v4 -= (v4 - (v4 >> ((750) & 63)))
                                if (1.5 <= 0.0) {
                                    if (((false || (1 >= 0)) && true) && true) {
                                        v126 := GenS11{F0: 0}
                                        _ = v126
                                    } else {
                                        var v127 uint = 2
                                        var v128 uint8 = 255
                                        v129 := GenS14{F0: 42}
                                        if true {
                                            var v130 uint8 = v128
                                            fmt.Println(v127, v102, v3)
                                            _ = v130
                                        }
                                        if true {
                                            v127 += v127
                                            v4 += v4
                                            var v131 int = (fi120 * (fi120 % ((int(v122)) | 1)))
                                            _ = v131
                                        } else {
                                            v132 := GenS5{F0: "foo bar", F1: v122}
                                            var v133 float32 = 1.5
                                            _ = v132
                                            _ = v133
                                        }
                                        v134 := GenS12{F0: (^(v128))}
                                        var v135 uint8 = 158
                                        _ = v127
                                        _ = v128
                                        _ = v129
                                        _ = v134
                                        _ = v135
                                    }
                                    var v136 float64 = 3.14159
                                    v122 = (v122 - (v122 ^ -1))
                                    _ = v136
                                } else {
                                    v137 := GenS15{F0: (^(v4)), F1: fi103}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if true {
                                        v138 := GenS16{F0: v122, F1: 1, F2: v122}
                                        _ = v138
                                    }
                                    _ = v137
                                }
                                _ = fi120
                                _ = v121
                                _ = v122
                                _ = v124
                                _ = v125
                            }
                        } else {
                            if false {
                                if (100 > 0) {
                                    v4 += (v4 * v4)
                                    if (!true) {
                                        var v139 uint32 = 1
                                        _ = v139
                                    } else {
                                        for fi140 := 0; fi140 < 6; fi140++ {
                                            if len(v1) > 0 {
                                                v1[0] = fi140
                                            }
                                            fmt.Println(v4)
                                            _ = fi140
                                        }
                                        v4 += 991
                                        for fi141 := 0; fi141 < 18; fi141++ {
                                            if false {
                                                v3 = (v3 + v3)
                                                var v142 uint16 = 255
                                                v143 := GenS17{F0: 3, F1: 55}
                                                _ = v142
                                                _ = v143
                                            }
                                            _ = fi141
                                        }
                                    }
                                }
                                var v144 int32 = 0
                                fmt.Println(v3, v102, fi103)
                                _ = v144
                            }
                        }
                        _ = v116
                    }
                }
                for fi145 := 0; fi145 < 2; fi145++ {
                    v146 := GenS18{F0: (((("a" + v3) + (v3 + v3)) + "hello") + (v3 + (v3 + v3))), F1: 65535, F2: 18446744073709551615}
                    _ = fi145
                    _ = v146
                }
                _ = fi103
            }
            _ = v102
        }
    }
    var v147 uint32 = v4
    for fi148 := 0; fi148 < 10; fi148++ {
        switch v3 {
        case "a":
            if (((255 == 42) || true) && false) {
                for fi149 := 0; fi149 < 18; fi149++ {
                    _ = len(v1)
                    var v150 uint = 65535
                    _ = fi149
                    _ = v150
                }
            } else {
                if (255 >= 2) {
                    v3 = ("hello" + (v3 + (("\n\t" + v3) + "\n\t")))
                }
            }
        case "\n\t":
            _ = len(v1)
        case "foo bar":
            var v151 uint8 = 148
            _ = v151
        default:
            v4 += 4294967295
        }
        v152 := GenS19{F0: 127}
        if len(v1) > 0 {
            v1[0] = (fi148 ^ 647)
        }
        _ = fi148
        _ = v152
    }
    switch v3 {
    case "":
        var v153 uint16 = 1
        _ = v153
    case "hello":
        v154 := GenS2{F0: 42, F1: (v147 >> ((42) & 63)), F2: -149}
        _ = v154
    }
    v147 -= 2147483647
    v155 := GenS20{F0: (((1e10 / 0.1) >= 1.5) || true)}
    switch 852 {
    case 4294967295:
        switch ((v3 + ("a" + v3)) + "hello") {
        case "a":
            v3 = v3
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v147:", v147)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v147
    _ = v155
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
