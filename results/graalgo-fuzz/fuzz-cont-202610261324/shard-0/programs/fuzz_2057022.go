package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8; F1 bool; F2 int16 }
type GenS9 struct { F0 uint64; F1 uint8; F2 int }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 string; F1 string }
type GenS12 struct { F0 int16; F1 int64 }
type GenS13 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS14 struct { F0 int16; F1 int; F2 int }
type GenS15 struct { F0 bool; F1 uint16 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint32; F1 int16; F2 int64 }
type GenS18 struct { F0 int32; F1 bool }
type GenS19 struct { F0 uint16; F1 uint8 }
func genFn0() bool {
    var v7 int16 = 42
    var v8 int64 = -1
    v9 := []int{-902, -1, -2147483648}
    v7 += (v7 >> ((4294967295) & 63))
    fmt.Println(v8, v8)
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    v1 := []int{-2147483648, 4294967295}
    var v2 float64 = 1.5
    var v3 int = -317
    var v4 int = 3
    v5 := make(map[int]uint32)
    var v6 uint8 = 0
    switch v3 {
    case 4294967295:
        for fi10 := 0; fi10 < 8; fi10++ {
            for fi11 := 0; fi11 < 2; fi11++ {
                if (42 != 1) {
                    var v12 float64 = v2
                    _ = v12
                }
                v13 := GenS1{F0: 0, F1: 255, F2: v4}
                fmt.Println(v2, v2)
                _ = fi11
                _ = v13
            }
            _ = fi10
        }
    case 128:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -9223372036854775808:
        switch v3 {
        case 9223372036854775807:
            v14 := GenS2{F0: 4294967295, F1: 42}
            _ = v14
        case -2147483648:
            switch (v4 << ((2) & 63)) {
            case 32768:
                switch "a" {
                case "a":
                    if false {
                        var v15 uint8 = v6
                        var v16 float64 = ((float64((0.1 - 0.0)) + v2) - float64(float32(v2)))
                        _ = v15
                        _ = v16
                    } else {
                        delete(v5, (v4 >> ((570) & 63)))
                        var v17 uint16 = 2
                        var v18 uint = 146
                        _ = v17
                        _ = v18
                    }
                default:
                    fmt.Println(v3, v6)
                }
            }
        default:
            v4 = (v3 + v4)
        }
    default:
        var v19 uint = 782
        _ = v19
    }
    var v20 string = "foo bar"
    v21 := GenS2{F0: 2147483647, F1: 424}
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 int64 = -1
        for fi24 := 0; fi24 < 13; fi24++ {
            switch fi24 {
            case 42:
                var v25 uint32 = 2
                _ = v25
            default:
                if (1 <= 0) {
                    switch (v20 + v20) {
                    case "\n\t":
                        v6 -= (v6 >> ((4294967295) & 63))
                    default:
                        if true {
                            switch "\n\t" {
                            case "":
                                v26 := GenS2{F0: 1, F1: 18446744073709551615}
                                _ = v26
                            case "\n\t":
                                v27 := GenS3{F0: 547}
                                _ = v27
                            case "a":
                                _ = len(v1)
                            default:
                                v23 = 42
                            }
                            var v28 float64 = (v2 * v2)
                            _ = v28
                        }
                    }
                }
            }
            _ = fi24
        }
        _ = fi22
        _ = v23
    }
    var v29 int8 = int8(v6)
    v6 = v6
    if len(v1) > 0 {
        v1[0] = (v4 << ((1) & 63))
    }
    var v30 uint16 = 27
    for fi31 := 0; fi31 < 6; fi31++ {
        switch (((v20 + "") + ("\n\t" + ("\n\t" + v20))) + "\n\t") {
        case "":
            var v32 float32 = 1e10
            _ = v32
        case "\n\t":
            switch v3 {
            case -261:
                var v33 int8 = int8(v30)
                _ = v33
            case 183:
                if ((v6 >= (v6 - 1)) && (false || (v6 <= v6))) {
                    switch ((v20 + (v20 + v20)) + v20) {
                    case "\n\t":
                        if false {
                            var v34 uint16 = (v30 >> ((27) & 63))
                            for fi35 := 0; fi35 < 1; fi35++ {
                                v1 = append(v1, (v4 % ((fi31) | 1)))
                                _ = fi35
                            }
                            var v36 float64 = v2
                            _ = v34
                            _ = v36
                        } else {
                            var v37 uint64 = 1
                            switch (fi31 * (v4 - -1)) {
                            case 32767:
                                v37 += (v37 >> ((18446744073709551615) & 63))
                            case -228:
                                if true {
                                    v38 := GenS4{F0: 1}
                                    switch ("hello" + "") {
                                    case "hello":
                                        v39 := GenS5{F0: 0, F1: v20}
                                        _ = v39
                                    default:
                                        v2 -= v2
                                    }
                                    _ = v38
                                } else {
                                    for fi40 := 0; fi40 < 10; fi40++ {
                                        v6 += v6
                                        _ = len(v1)
                                        _ = fi40
                                    }
                                }
                            case 42:
                                var v41 string = (v20 + "a")
                                _ = v41
                            }
                            _ = v37
                        }
                    }
                    if (!false) {
                        v4 = (-(fi31))
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if true {
                                for fi42 := 0; fi42 < 9; fi42++ {
                                    if (127 != 100) {
                                        v43 := GenS6{F0: 1}
                                        v30 += 875
                                        v44 := GenS7{F0: 904}
                                        if (true && (!true)) {
                                            v1 = append(v1, (v3 & -94))
                                            delete(v5, (v3 + fi42))
                                            if (false || (4294967295 <= 4294967295)) {
                                                var v45 uint8 = 55
                                                fmt.Println(v45)
                                                v46 := GenS7{F0: (v30 - 1)}
                                                v45 = 42
                                                fmt.Println(v3, v45)
                                                _ = v45
                                                _ = v46
                                            } else {
                                                var v47 int8 = (v29 << ((736) & 63))
                                                v20 = ""
                                                v48 := GenS8{F0: v47, F1: false, F2: 32767}
                                                _ = v47
                                                _ = v48
                                            }
                                        } else {
                                            if false {
                                                fmt.Println(v29)
                                                if true {
                                                    var v49 uint8 = (v6 >> ((969) & 63))
                                                    _ = v49
                                                } else {
                                                    var v50 uint8 = v6
                                                    if ((!((v50 << ((65535) & 63)) >= (v6 ^ v50))) && ((614 >= -2147483648) && (!(42 <= 0)))) {
                                                        v3 += v4
                                                    }
                                                    var v51 int8 = (v29 / ((48) | 1))
                                                    _ = v50
                                                    _ = v51
                                                }
                                                fmt.Println(v3, fi42)
                                            }
                                            v2 = (((v2 - v2) / ((((v2 / 1.0) * float64(((float32((v2 * 1e10)) / (-2.25 * 0.1)) + (1.5 / 1e10)))) * float64((float32((float64(1e10) * (v2 / (v2 - v2)))) - (-2.25 * 0.0)))) / 1e10)) - ((v2 + v2) + 3.14159))
                                        }
                                        _ = v43
                                        _ = v44
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        v52 := GenS4{F0: (v6 & v6)}
                                        _ = v52
                                    default:
                                        v53 := v5[(fi31 & (v4 * fi42))]
                                        _ = v53
                                    }
                                    switch 368 {
                                    case 127:
                                        v54 := GenS5{F0: 0, F1: v20}
                                        _ = v54
                                    case 4294967295:
                                        var v55 bool = (true || true)
                                        _ = v55
                                    default:
                                        switch v20 {
                                        case "\n\t":
                                            switch ("foo bar" + v20) {
                                            case "":
                                                var v56 float64 = 1.5
                                                _ = v56
                                            case "\n\t":
                                                var v57 int8 = v29
                                                _ = v57
                                            default:
                                                fmt.Println(v4)
                                            }
                                        case "hello":
                                            if true {
                                                v29 -= (v29 * 100)
                                            }
                                        case "foo bar":
                                            switch v20 {
                                            case "":
                                                v20 = v20
                                            case "\n\t":
                                                v4 += (^(v3))
                                            }
                                        default:
                                            var v58 uint8 = 147
                                            _ = v58
                                        }
                                    }
                                    v29 = -128
                                    _ = fi42
                                }
                            }
                        } else {
                            v59 := GenS1{F0: v30, F1: 18446744073709551615, F2: (v3 * (v4 + fi31))}
                            switch v20 {
                            case "foo bar":
                                if (-1 > 1) {
                                    v2 += 0.0
                                    fmt.Println(v20, v6, v2)
                                    fmt.Println(v4, v4)
                                } else {
                                    if (v20 == (("hello" + v20) + v20)) {
                                        if ((!false) || (false && (1 <= 994))) {
                                            var v60 uint = 255
                                            _ = v60
                                        }
                                        for fi61 := 0; fi61 < 1; fi61++ {
                                            var v62 uint = 743
                                            var v63 int32 = 0
                                            var v64 uint8 = (^(v6))
                                            _ = fi61
                                            _ = v62
                                            _ = v63
                                            _ = v64
                                        }
                                    }
                                    if false {
                                        var v65 int8 = v29
                                        _ = v65
                                    } else {
                                        v66 := GenS9{F0: 18446744073709551615, F1: (v6 >> ((255) & 63)), F2: v4}
                                        v30 = (v30 % ((v30) | 1))
                                        _ = v66
                                    }
                                }
                            case "a":
                                v29 = (v29 ^ v29)
                            default:
                                v2 += v2
                            }
                            _ = v59
                        }
                    }
                } else {
                    for fi67 := 0; fi67 < 6; fi67++ {
                        var v68 int64 = -349
                        _ = fi67
                        _ = v68
                    }
                    for fi69 := 0; fi69 < 7; fi69++ {
                        if (!(!true)) {
                            v3 = v3
                        } else {
                            v4 = (fi31 | 127)
                            v70 := GenS10{F0: 2}
                            v6 = (v6 % ((42) | 1))
                            if (118 == v30) {
                                if (!(-485 < -1)) {
                                    if true {
                                        switch fi31 {
                                        case 32768:
                                            v3 = fi31
                                        case 92:
                                            v20 = ((((v20 + "") + "") + v20) + (v20 + ((("a" + ("a" + v20)) + (("\n\t" + (v20 + v20)) + v20)) + v20)))
                                        default:
                                            if (42 > 538) {
                                                if (((v30 / ((v30) | 1)) > uint16(v6)) || false) {
                                                    switch v20 {
                                                    case "\n\t":
                                                        var v71 int16 = 1000
                                                        _ = v71
                                                    case "hello":
                                                        if true {
                                                            var v72 int16 = 1000
                                                            switch fi31 {
                                                            case 0:
                                                                v30 = 120
                                                            default:
                                                                var v73 bool = (2 <= 834)
                                                                _ = v73
                                                            }
                                                            _ = v72
                                                        } else {
                                                            fmt.Println(fi31)
                                                            var v74 uint = 219
                                                            var v75 int = fi31
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                    default:
                                                        fmt.Println(v2, v6)
                                                        if (v4 < -2147483648) {
                                                            if (0 > 65535) {
                                                                var v76 float32 = ((((3.14159 - 3.14159) * ((0.0 - ((0.0 / 1e10) * 0.1)) * 1e10)) - (((0.0 / 1e10) * float32(1.5)) + -2.25)) / -2.25)
                                                                var v77 uint16 = v30
                                                                if ("foo bar" != v20) {
                                                                    v78 := GenS10{F0: 45}
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    _ = v78
                                                                } else {
                                                                    v4 -= (^(v3))
                                                                    if (!((!(!true)) || (!(true || ("hello" >= ("\n\t" + v20)))))) {
                                                                        switch (v3 & -2147483648) {
                                                                        case 4294967295:
                                                                            if ((v6 % ((v6) | 1)) == v6) {
                                                                                var v79 int8 = (^(v29))
                                                                                v80 := v5[v4]
                                                                                var v81 uint32 = v80
                                                                                _ = v79
                                                                                _ = v80
                                                                                _ = v81
                                                                            } else {
                                                                                if (true && ((!false) || (1 > 4294967295))) {
                                                                                    if ((!(139 <= 0)) && true) {
                                                                                        switch v20 {
                                                                                        case "a":
                                                                                            var v82 int16 = -1
                                                                                            _ = v82
                                                                                        case "":
                                                                                            if (!(65535 > 4294967295)) {
                                                                                                switch ("hello" + (("hello" + v20) + v20)) {
                                                                                                case "":
                                                                                                    v6 += (^(v6))
                                                                                                case "a":
                                                                                                    v83 := GenS7{F0: v77}
                                                                                                    _ = v83
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v4, v2, v6)
                                                                                                default:
                                                                                                    v1 = append(v1, fi31)
                                                                                                }
                                                                                                var v84 int = v4
                                                                                                delete(v5, v84)
                                                                                                _ = v84
                                                                                            }
                                                                                        case "\n\t":
                                                                                            v85 := GenS7{F0: v30}
                                                                                            _ = v85
                                                                                        default:
                                                                                            v2 -= v2
                                                                                        }
                                                                                        v86 := GenS10{F0: 2147483647}
                                                                                        if (454 <= 42) {
                                                                                            var v87 bool = (!(53 >= 1))
                                                                                            var v88 float32 = (1.5 - v76)
                                                                                            fmt.Println(v88, v2, fi69)
                                                                                            _ = v87
                                                                                            _ = v88
                                                                                        }
                                                                                        if true {
                                                                                            v29 = (v29 << ((198) & 63))
                                                                                        }
                                                                                        switch 0 {
                                                                                        case 639:
                                                                                            var v89 uint = 4294967295
                                                                                            _ = v89
                                                                                        case 785:
                                                                                            v29 = v29
                                                                                        case 230:
                                                                                            v29 = 23
                                                                                        }
                                                                                        _ = v86
                                                                                    }
                                                                                }
                                                                                var v90 int32 = 1
                                                                                var v91 uint = 676
                                                                                _ = v90
                                                                                _ = v91
                                                                            }
                                                                        case 1:
                                                                            fmt.Println(v20)
                                                                        case 0:
                                                                            var v92 int64 = 1
                                                                            switch fi69 {
                                                                            case -590:
                                                                                v77 = (^(v77))
                                                                            case 127:
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            default:
                                                                                var v93 uint64 = 1
                                                                                _ = v93
                                                                            }
                                                                            v94 := GenS11{F0: "\n\t", F1: "hello"}
                                                                            _ = v92
                                                                            _ = v94
                                                                        }
                                                                    } else {
                                                                        v76 = (v76 + v76)
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                    }
                                                                    switch ((v20 + v20) + "hello") {
                                                                    case "hello":
                                                                        v30 += (v30 >> ((4294967295) & 63))
                                                                        var v95 uint64 = 1
                                                                        _ = v95
                                                                    case "foo bar":
                                                                        if (1 != 0) {
                                                                            v76 = (v76 / 3.14159)
                                                                            var v96 int32 = -1
                                                                            _ = v96
                                                                        } else {
                                                                            if false {
                                                                                v97 := GenS12{F0: -1, F1: -80}
                                                                                var v98 int = 9223372036854775807
                                                                                fmt.Println(v3, fi69, fi31)
                                                                                _ = v97
                                                                                _ = v98
                                                                            } else {
                                                                                v29 -= (-(v29))
                                                                                v99 := GenS3{F0: (v30 >> ((18446744073709551615) & 63))}
                                                                                _ = v99
                                                                            }
                                                                            var v100 bool = false
                                                                            fmt.Println(fi31, v100, v2)
                                                                            switch v20 {
                                                                            case "foo bar":
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = int(v29)
                                                                                }
                                                                            }
                                                                            v101 := v5[v4]
                                                                            var v102 float32 = ((v76 + v76) * 1.5)
                                                                            _ = v100
                                                                            _ = v101
                                                                            _ = v102
                                                                        }
                                                                    case "":
                                                                        fmt.Println(v6)
                                                                    }
                                                                    v4 = (fi31 << ((0) & 63))
                                                                }
                                                                if (-705 >= -9223372036854775808) {
                                                                    v103 := GenS13{F0: v6, F1: -1, F2: v6}
                                                                    fmt.Println(v29, v77, v4)
                                                                    switch fi69 {
                                                                    case 0:
                                                                        var v104 uint = 960
                                                                        _ = v104
                                                                    case 2147483647:
                                                                        switch fi69 {
                                                                        case 178:
                                                                            if (int32(v30) <= -1) {
                                                                                v6 = v6
                                                                                var v105 int16 = 1000
                                                                                _ = v105
                                                                            } else {
                                                                                v20 = ((("" + (v20 + v20)) + v20) + "")
                                                                            }
                                                                        case 0:
                                                                            v106 := GenS14{F0: 1, F1: fi69, F2: (-(fi31))}
                                                                            _ = v106
                                                                        }
                                                                    case -373:
                                                                        fmt.Println(v2)
                                                                    }
                                                                    _ = v103
                                                                }
                                                                v107 := GenS14{F0: 1000, F1: (^(v4)), F2: v4}
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    if (!false) {
                                                                        v77 -= v77
                                                                        var v108 uint = 255
                                                                        _ = len(v1)
                                                                        _ = v108
                                                                    } else {
                                                                        var v109 uint64 = 18446744073709551615
                                                                        v110 := GenS6{F0: v6}
                                                                        var v111 uint16 = uint16(v109)
                                                                        _ = v109
                                                                        _ = v110
                                                                        _ = v111
                                                                    }
                                                                }
                                                                _ = v107
                                                                _ = v76
                                                                _ = v77
                                                            }
                                                            fmt.Println(v6)
                                                        } else {
                                                            var v112 float64 = (v2 / -2.25)
                                                            if false {
                                                                if (!true) {
                                                                    _ = len(v1)
                                                                } else {
                                                                    var v113 float32 = 0.1
                                                                    v114 := GenS10{F0: 2147483647}
                                                                    _ = v113
                                                                    _ = v114
                                                                }
                                                            } else {
                                                                v115 := GenS15{F0: (9223372036854775807 != -864), F1: v30}
                                                                _ = v115
                                                            }
                                                            _ = v112
                                                        }
                                                    }
                                                    v116 := GenS7{F0: v30}
                                                    v29 += v29
                                                    _ = v116
                                                } else {
                                                    v117 := GenS16{F0: 482}
                                                    _ = v117
                                                }
                                            }
                                        }
                                        v29 = v29
                                        switch "\n\t" {
                                        case "foo bar":
                                            v30 = (v30 + (v30 / ((0) | 1)))
                                            v118 := GenS9{F0: 42, F1: 42, F2: (fi69 - fi69)}
                                            v20 = "foo bar"
                                            _ = v118
                                        }
                                    }
                                    v30 -= v30
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                var v119 uint = 65535
                                _ = v119
                            } else {
                                v120 := GenS17{F0: 0, F1: -49, F2: -1}
                                _ = v120
                            }
                            _ = v70
                        }
                        v121 := GenS18{F0: 2147483647, F1: true}
                        var v122 int = fi69
                        v2 = 0.0
                        _ = fi69
                        _ = v121
                        _ = v122
                    }
                    fmt.Println(v4)
                }
            }
        case "hello":
            v123 := GenS19{F0: 2, F1: v6}
            _ = v123
        default:
            fmt.Println(fi31, fi31)
        }
        _ = fi31
    }
    if (((!(((!(("a" + v20) < "")) && (false || ((-579 <= 0) || ("" == v20)))) || true)) || false) && (910 <= 2147483647)) {
        switch (^(v4)) {
        case 480:
            var v124 int64 = -9223372036854775808
            _ = v124
        default:
            fmt.Println(v6, v6)
        }
        if false {
            var v125 float32 = 3.14159
            _ = v125
        } else {
            v126 := GenS2{F0: 2147483647, F1: 18446744073709551615}
            v5[(v3 % ((v4) | 1))] = 403
            _ = v126
        }
    }
    if true {
        v127 := GenS9{F0: 4294967295, F1: (^(v6)), F2: (v3 | (v3 ^ v4))}
        var v128 int64 = 9223372036854775807
        for fi129 := 0; fi129 < 3; fi129++ {
            if len(v1) > 0 {
                v1[0] = v4
            }
            fmt.Println(v128)
            _ = fi129
        }
        _ = v127
        _ = v128
    }
    fmt.Println(v29, v4)
    if false {
        fmt.Println(v2)
        var v130 uint8 = 42
        _ = v130
    }
    switch "\n\t" {
    case "a":
        v29 = int8(v6)
    case "":
        v3 = (v3 % (((v3 * (v3 | v3))) | 1))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
