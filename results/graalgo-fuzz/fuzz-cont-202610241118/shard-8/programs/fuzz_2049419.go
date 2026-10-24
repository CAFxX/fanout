package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int; F1 int; F2 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64; F1 uint32 }
type GenS10 struct { F0 int8; F1 int8; F2 uint64 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int16; F1 int; F2 int8 }
type GenS13 struct { F0 uint; F1 uint64 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint8; F1 int8 }
type GenS16 struct { F0 int16; F1 bool; F2 int }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS19 struct { F0 int16; F1 uint64; F2 int32 }
type GenS20 struct { F0 uint }
type GenS21 struct { F0 string; F1 uint64 }

func main() {
    var v1 uint = 728
    var v2 int16 = 32767
    var v3 int16 = (v2 >> ((1) & 63))
    v4 := []int{-2147483648, 128, -1, 2147483647}
    v3 -= v3
    switch "\n\t" {
    case "hello":
        var v5 int64 = 9223372036854775807
        _ = v5
    case "a":
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 int8 = 1
            _ = fi6
            _ = v7
        }
    case "":
        if (!(int64(v2) != 472)) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v2)
        } else {
            if false {
                switch "foo bar" {
                case "foo bar":
                    v3 = v3
                }
            } else {
                v8 := GenS1{F0: 2, F1: 2, F2: 2147483647}
                _ = v8
            }
        }
    default:
        var v9 int = 32767
        _ = v9
    }
    v2 += -721
    var v10 int32 = -1
    fmt.Println(v3, v2)
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 19; fi12++ {
            switch (-(fi11)) {
            case 128:
                if len(v4) > 0 {
                    v4[0] = -2147483648
                }
                v13 := GenS1{F0: 32767, F1: 2, F2: 563}
                _ = v13
            case 157:
                var v14 uint64 = 42
                _ = v14
            default:
                var v15 int64 = 9223372036854775807
                _ = v15
            }
            if ((v1 - 65535) > v1) {
                switch fi12 {
                case -1:
                    v2 += v2
                }
                v2 = (v2 >> ((v1) & 63))
            } else {
                var v16 int16 = v3
                var v17 int = fi11
                if false {
                    v18 := GenS2{F0: 9223372036854775807}
                    _ = v18
                } else {
                    switch "" {
                    case "hello":
                        _ = len(v4)
                    }
                }
                switch 127 {
                case 32768:
                    if (32768 > v17) {
                        v19 := GenS3{F0: v2, F1: "foo bar"}
                        var v20 uint = 65535
                        v21 := GenS4{F0: fi12, F1: 1, F2: int16(v1)}
                        v2 -= 32767
                        fmt.Println(v17, v10)
                        var v22 float64 = (0.0 * 1.5)
                        if (("foo bar" + "") != "\n\t") {
                            switch "foo bar" {
                            case "hello":
                                switch ("" + "hello") {
                                case "a":
                                    v20 -= (^(v1))
                                default:
                                    if false {
                                        v23 := GenS2{F0: 42}
                                        var v24 uint32 = 2
                                        v25 := GenS5{F0: (^(v1))}
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                }
                            }
                            fmt.Println(fi12, fi11, v17)
                            v16 = v3
                        } else {
                            v26 := GenS6{F0: false}
                            _ = v26
                        }
                        _ = v19
                        _ = v20
                        _ = v21
                        _ = v22
                    } else {
                        v27 := GenS6{F0: false}
                        v16 -= v3
                        switch (fi12 * 42) {
                        case 9223372036854775807:
                            var v28 bool = (42 == 1)
                            _ = v28
                        default:
                            v10 = 1
                        }
                        _ = v27
                    }
                case 32767:
                    if len(v4) > 0 {
                        v4[0] = (fi12 | 32768)
                    }
                }
                var v29 uint32 = 745
                _ = v16
                _ = v17
                _ = v29
            }
            _ = fi12
        }
        _ = fi11
    }
    fmt.Println(v10)
    switch ("hello" + "foo bar") {
    case "":
        for fi30 := 0; fi30 < 15; fi30++ {
            for fi31 := 0; fi31 < 10; fi31++ {
                v32 := GenS2{F0: 685}
                switch (fi30 + 32768) {
                case 4294967295:
                    switch "\n\t" {
                    case "\n\t":
                        v10 += (v10 / ((v10) | 1))
                    }
                case 0:
                    var v33 string = "hello"
                    _ = v33
                case 597:
                    switch (("a" + ((("\n\t" + (("hello" + "a") + "\n\t")) + "") + ("" + "hello"))) + "foo bar") {
                    case "a":
                        v34 := GenS7{F0: 2, F1: false}
                        _ = v34
                    default:
                        if true {
                            if (int32(v1) >= 1) {
                                fmt.Println(fi31)
                                if true {
                                    v2 += int16(v1)
                                    v35 := GenS8{F0: 0}
                                    if false {
                                        var v36 float64 = ((((3.14159 * 1e10) / 1.0) - float64(-2.25)) - float64(((0.1 + (1.5 / (-2.25 + 0.1))) + 3.14159)))
                                        if ((!(false || true)) && (1 != 969)) {
                                            v37 := GenS9{F0: 4294967295, F1: 2147483648}
                                            v10 += (v10 / (((v10 - 0)) | 1))
                                            _ = v37
                                        } else {
                                            fmt.Println(v1)
                                            v3 = 944
                                        }
                                        _ = v36
                                    } else {
                                        var v38 bool = (false || true)
                                        v10 -= (v10 % (((-(v10))) | 1))
                                        _ = v38
                                    }
                                    _ = v35
                                } else {
                                    v39 := GenS8{F0: 255}
                                    _ = v39
                                }
                                v40 := GenS1{F0: 255, F1: 0, F2: 1}
                                _ = v40
                            } else {
                                if false {
                                    v10 += v10
                                    v1 = (v1 - v1)
                                    fmt.Println(v10)
                                }
                                fmt.Println(v2, v1)
                                var v41 uint = v1
                                _ = v41
                            }
                            v10 = v10
                        }
                        fmt.Println(v1, fi31)
                        fmt.Println(fi31, v1, v3)
                    }
                default:
                    if (608 == 2147483647) {
                        v42 := GenS5{F0: v1}
                        _ = v42
                    } else {
                        fmt.Println(v2)
                        v43 := GenS10{F0: 1, F1: -1, F2: 4294967295}
                        fmt.Println(v2)
                        v2 += 42
                        if (65535 > 1) {
                            v3 = v3
                        } else {
                            if (4294967295 < 42) {
                                var v44 uint64 = 0
                                if false {
                                    switch "foo bar" {
                                    case "foo bar":
                                        if ((0 > 141) && (false && (454 >= 32767))) {
                                            v10 -= -1
                                        }
                                        if false {
                                            v3 -= (-(v2))
                                            if ((-2.25 - float32((0.0 / 3.14159))) < 1e10) {
                                                var v45 bool = false
                                                v45 = v45
                                                if ((fi30 != (-(fi30))) && true) {
                                                    v2 += 856
                                                    fmt.Println(v3)
                                                    v3 = v3
                                                }
                                                _ = v45
                                            }
                                            _ = len(v4)
                                        } else {
                                            if (!false) {
                                                v46 := GenS10{F0: 0, F1: 1, F2: (v44 * 4294967295)}
                                                _ = v46
                                            } else {
                                                fmt.Println(v44, v10, v2)
                                                var v47 int8 = 66
                                                var v48 uint64 = (v44 ^ v44)
                                                v47 += (v47 & v47)
                                                _ = v47
                                                _ = v48
                                            }
                                            v44 += (^(v44))
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                    case "\n\t":
                                        var v49 int16 = -32768
                                        _ = v49
                                    }
                                    switch "" {
                                    case "":
                                        v50 := GenS11{F0: (187 <= 58)}
                                        _ = v50
                                    default:
                                        var v51 uint64 = 163
                                        _ = v51
                                    }
                                }
                                _ = v44
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v52 := GenS1{F0: 0, F1: 2, F2: 65535}
                                v2 -= v2
                                _ = v52
                            }
                            var v53 bool = (!(!(!false)))
                            _ = v53
                        }
                        if len(v4) > 0 {
                            v4[0] = (fi30 / (((fi31 << ((65535) & 63))) | 1))
                        }
                        _ = v43
                    }
                }
                if true {
                    if false {
                        v54 := GenS8{F0: 0}
                        var v55 uint64 = 338
                        _ = v54
                        _ = v55
                    } else {
                        if ((32767 == 32767) || (((65535 > 32767) && (!(!(-56 > 1)))) || (int32(v3) >= v10))) {
                            v2 = (v3 >> ((18446744073709551615) & 63))
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        switch (-(fi30)) {
                        case 4294967295:
                            var v56 int = 2147483647
                            _ = v56
                        case 553:
                            var v57 float64 = (3.14159 / -2.25)
                            _ = v57
                        }
                    }
                    v4 = append(v4, fi30)
                } else {
                    var v58 uint32 = 870
                    if false {
                        v10 -= v10
                        fmt.Println(v10, v1)
                    }
                    v59 := GenS2{F0: -9223372036854775808}
                    _ = v58
                    _ = v59
                }
                _ = fi31
                _ = v32
            }
            var v60 int64 = 42
            _ = fi30
            _ = v60
        }
    case "\n\t":
        if len(v4) > 0 {
            v4[0] = 184
        }
    case "foo bar":
        var v61 uint = v1
        _ = v61
    default:
        if false {
            var v62 float32 = 1e10
            var v63 string = "hello"
            var v64 uint = v1
            _ = v62
            _ = v63
            _ = v64
        }
    }
    var v65 int32 = (v10 & (v10 | v10))
    _ = len(v4)
    if true {
        var v66 string = ""
        _ = v66
    } else {
        var v67 int16 = (v2 / ((908) | 1))
        _ = v67
    }
    if false {
        for fi68 := 0; fi68 < 5; fi68++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi68
        }
        fmt.Println(v2)
        for fi69 := 0; fi69 < 11; fi69++ {
            if true {
                v3 -= 62
                for fi70 := 0; fi70 < 12; fi70++ {
                    if (128 == 1) {
                        fmt.Println(v10, v1, v65)
                    }
                    v71 := GenS1{F0: 0, F1: 129, F2: 0}
                    switch ("hello" + "hello") {
                    case "a":
                        var v72 uint = v1
                        _ = v72
                    case "\n\t":
                        switch ((("" + "foo bar") + ("foo bar" + "\n\t")) + "\n\t") {
                        case "\n\t":
                            v73 := GenS12{F0: 1, F1: -446, F2: 1}
                            _ = v73
                        }
                    case "":
                        switch 1 {
                        case 9223372036854775807:
                            if (v1 > (v1 | v1)) {
                                if (0.1 == float64(-2.25)) {
                                    if (545 >= 255) {
                                        if (!true) {
                                            v65 += (v65 + -406)
                                            v2 = (v2 - v3)
                                            if ((313 == 32767) || (v1 <= (v1 % ((v1) | 1)))) {
                                                v65 = v65
                                            } else {
                                                v3 = -61
                                                var v74 float64 = 0.1
                                                _ = v74
                                            }
                                            if (2 >= 1) {
                                                v75 := GenS13{F0: 18446744073709551615, F1: 0}
                                                _ = v75
                                            }
                                        } else {
                                            v76 := GenS9{F0: 619, F1: 2147483648}
                                            _ = v76
                                        }
                                        if (false && (true || true)) {
                                            v1 -= (v1 * (^(v1)))
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v3 += 42
                                        }
                                    } else {
                                        fmt.Println(v1, v1, v65)
                                        if (2 >= 753) {
                                            if true {
                                                if (!true) {
                                                    v10 += -1
                                                    var v77 string = ("foo bar" + "\n\t")
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    _ = v77
                                                } else {
                                                    switch ("hello" + ("\n\t" + "")) {
                                                    case "":
                                                        v2 = (-(v3))
                                                    case "\n\t":
                                                        v65 = int32(fi69)
                                                    case "a":
                                                        if (!false) {
                                                            var v78 uint = (v1 - (v1 % (((v1 * (^(v1)))) | 1)))
                                                            _ = v78
                                                        } else {
                                                            v79 := GenS6{F0: true}
                                                            _ = v79
                                                        }
                                                    }
                                                    var v80 uint32 = 42
                                                    _ = v80
                                                }
                                                v65 = (v65 & (v10 / ((v65) | 1)))
                                            }
                                            if false {
                                                if true {
                                                    v81 := GenS4{F0: fi69, F1: 42, F2: (v3 & (v2 % ((v3) | 1)))}
                                                    _ = v81
                                                } else {
                                                    v2 += v2
                                                    fmt.Println(v2, fi70, v10)
                                                    if true {
                                                        if (!false) {
                                                            _ = len(v4)
                                                        }
                                                        v10 -= v65
                                                        if true {
                                                            var v82 int = fi69
                                                            _ = v82
                                                        } else {
                                                            v1 = (v1 % ((v1) | 1))
                                                            var v83 bool = true
                                                            _ = v83
                                                        }
                                                    } else {
                                                        v84 := GenS14{F0: ((-76 > 105) && true)}
                                                        var v85 int = (-(fi69))
                                                        if false {
                                                            var v86 bool = false
                                                            _ = v86
                                                        } else {
                                                            v3 = (v3 & v3)
                                                            v65 -= -2147483648
                                                            var v87 string = "a"
                                                            fmt.Println(v2)
                                                            if (470 >= 18446744073709551615) {
                                                                v2 = v2
                                                                v1 = v1
                                                                v65 = (v65 % (((^(v65))) | 1))
                                                            }
                                                            var v88 float64 = (1.5 - (((((-2.25 / -2.25) / 0.1) * 1e10) / -2.25) / float64(1e10)))
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        var v89 uint64 = 17
                                                        _ = v84
                                                        _ = v85
                                                        _ = v89
                                                    }
                                                }
                                                v65 = v10
                                                v90 := GenS2{F0: -1}
                                                _ = v90
                                            } else {
                                                fmt.Println(v1, v1, v1)
                                            }
                                            fmt.Println(fi70, v1)
                                        }
                                        if true {
                                            fmt.Println(v1)
                                            if ((("foo bar" + "foo bar") + "") > "hello") {
                                                var v91 int64 = 9223372036854775807
                                                _ = len(v4)
                                                _ = len(v4)
                                                _ = v91
                                            }
                                        } else {
                                            var v92 int8 = 0
                                            _ = v92
                                        }
                                        v10 += (v10 | int32(v1))
                                        var v93 int8 = 100
                                        if true {
                                            if false {
                                                v1 = (v1 ^ v1)
                                                v2 = (v3 - int16(v65))
                                                var v94 int64 = -9223372036854775808
                                                _ = len(v4)
                                                _ = v94
                                            } else {
                                                v1 = (v1 * v1)
                                            }
                                            fmt.Println(v65)
                                        }
                                        _ = v93
                                    }
                                    if (fi70 > -881) {
                                        if false {
                                            switch "" {
                                            case "a":
                                                if true {
                                                    var v95 float32 = 3.14159
                                                    v3 += 32767
                                                    v96 := GenS11{F0: (!true)}
                                                    if true {
                                                        v3 += int16(v1)
                                                    } else {
                                                        if (127 != -76) {
                                                            switch (fi69 | (fi70 >> ((v1) & 63))) {
                                                            case -304:
                                                                var v97 int = fi69
                                                                _ = v97
                                                            default:
                                                                var v98 int64 = -1
                                                                _ = v98
                                                            }
                                                            v99 := GenS12{F0: (v2 * v3), F1: (fi69 - fi69), F2: 1}
                                                            var v100 int64 = 1
                                                            _ = v100
                                                            _ = v99
                                                        }
                                                        var v101 uint16 = 2
                                                        _ = v101
                                                    }
                                                    _ = v95
                                                    _ = v96
                                                }
                                            case "\n\t":
                                                var v102 int32 = (v10 | (v10 | (v10 * 711)))
                                                _ = v102
                                            }
                                            v65 = (v10 * (v10 - -260))
                                            var v103 uint64 = 42
                                            _ = v103
                                        }
                                        switch "hello" {
                                        case "foo bar":
                                            var v104 uint8 = 0
                                            _ = v104
                                        case "":
                                            var v105 int32 = v65
                                            _ = v105
                                        default:
                                            if true {
                                                if true {
                                                    fmt.Println(fi69, fi69, v10)
                                                    v1 += (v1 >> ((2) & 63))
                                                    v106 := GenS15{F0: 0, F1: -1}
                                                    _ = v106
                                                } else {
                                                    v2 += (v3 / (((v2 / (((^(v3))) | 1))) | 1))
                                                    if (!(1 <= 1)) {
                                                        v107 := GenS15{F0: 1, F1: int8(v1)}
                                                        v108 := GenS16{F0: v2, F1: true, F2: (fi69 / ((fi69) | 1))}
                                                        _ = v107
                                                        _ = v108
                                                    } else {
                                                        v65 = 1
                                                    }
                                                }
                                                if (v10 <= (^(v10))) {
                                                    v10 = v65
                                                    if len(v4) > 0 {
                                                        v4[0] = fi69
                                                    }
                                                }
                                                switch "" {
                                                case "":
                                                    v109 := GenS2{F0: -9223372036854775808}
                                                    _ = v109
                                                case "foo bar":
                                                    v110 := GenS17{F0: 0}
                                                    _ = v110
                                                case "\n\t":
                                                    var v111 int = (fi70 % (((fi70 ^ fi69)) | 1))
                                                    _ = v111
                                                }
                                            } else {
                                                v112 := GenS3{F0: v2, F1: "hello"}
                                                _ = v112
                                            }
                                        }
                                        v10 = (v10 / (((v10 - v65)) | 1))
                                    }
                                    _ = len(v4)
                                    v113 := GenS18{F0: 0, F1: 2, F2: 2}
                                    _ = v113
                                }
                                v114 := GenS19{F0: (v3 ^ -1), F1: 0, F2: (v10 | v10)}
                                _ = v114
                            }
                        case -1:
                            v2 = v3
                        }
                        var v115 int = fi70
                        if true {
                            if false {
                                v3 = v2
                                var v116 uint32 = 2147483647
                                _ = v116
                            }
                            fmt.Println(v1)
                        } else {
                            v117 := GenS12{F0: v3, F1: v115, F2: 127}
                            var v118 uint16 = 42
                            fmt.Println(v118)
                            _ = v117
                            _ = v118
                        }
                        _ = v115
                    default:
                        switch (fi69 / ((fi70) | 1)) {
                        case 809:
                            var v119 int16 = v3
                            _ = v119
                        case 32768:
                            v4 = append(v4, (fi70 % ((fi69) | 1)))
                        case 127:
                            switch "foo bar" {
                            case "foo bar":
                                v65 += (-(v65))
                            }
                        }
                    }
                    _ = fi70
                    _ = v71
                }
                v2 = (v3 & 32767)
            }
            fmt.Println(v65, v1, v1)
            if true {
                v3 = int16(v65)
                if true {
                    v120 := GenS14{F0: false}
                    _ = v120
                }
            }
            _ = fi69
        }
    } else {
        switch (("a" + "a") + "\n\t") {
        case "":
            switch 32768 {
            case 42:
                v121 := GenS18{F0: 4294967295, F1: 2, F2: 0}
                _ = v121
            case -9223372036854775808:
                v1 = v1
            case 207:
                var v122 int8 = 127
                _ = v122
            default:
                _ = len(v4)
            }
        }
    }
    for fi123 := 0; fi123 < 11; fi123++ {
        v124 := GenS20{F0: 2}
        v65 = 42
        _ = fi123
        _ = v124
    }
    if (true || false) {
        v65 = -1
        v3 = (v3 - 104)
        var v125 int = 127
        _ = v125
    } else {
        switch 832 {
        case 567:
            for fi126 := 0; fi126 < 12; fi126++ {
                var v127 float64 = 1e10
                _ = fi126
                _ = v127
            }
        case 1:
            _ = len(v4)
        }
        if true {
            var v128 int32 = v65
            v10 = (v10 % ((v65) | 1))
            if (-107 > 73) {
                var v129 uint = v1
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v3 += v2
                _ = v129
            }
            _ = v128
        }
        v130 := GenS21{F0: "hello", F1: 18446744073709551615}
        _ = v130
    }
    if (true || false) {
        _ = len(v4)
        v131 := GenS2{F0: 858}
        v1 -= v1
        _ = v131
    }
    fmt.Println(v10, v1, v3)
    v65 -= (v10 - (v10 / (((v65 / ((0) | 1))) | 1)))
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v65
}
