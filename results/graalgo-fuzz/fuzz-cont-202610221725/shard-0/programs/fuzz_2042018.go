package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int16; F1 uint; F2 bool }
type GenS7 struct { F0 int8; F1 uint8 }
type GenS8 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS9 struct { F0 uint8 }

func main() {
    var v1 int64 = 42
    var v2 uint = 42
    v3 := []int{32767, -974, 2147483647}
    v4 := []int{514, 736}
    v5 := []string{}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 -= v2
    if true {
        if (!false) {
            v4 = append(v4, 127)
        } else {
            var v6 uint32 = 2
            v2 = 18446744073709551615
            fmt.Println(v2, v6, v2)
            _ = v6
        }
        var v7 int = 42
        _ = v7
    } else {
        if false {
            v3 = append(v3, 32767)
            switch -475 {
            case 408:
                switch "hello" {
                case "foo bar":
                    if (false || (42 >= v1)) {
                        if (454 >= 2) {
                            var v8 int8 = 0
                            for fi9 := 0; fi9 < 7; fi9++ {
                                v1 = -81
                                _ = fi9
                            }
                            v10 := GenS1{F0: "hello", F1: 228, F2: -5}
                            _ = v10
                            _ = v8
                        }
                        v11 := GenS1{F0: ((((("foo bar" + (("a" + "a") + "foo bar")) + "a") + (("hello" + "foo bar") + "a")) + ("" + ("" + "hello"))) + ("hello" + (("" + ("" + "\n\t")) + "\n\t"))), F1: 0, F2: 0}
                        _ = v11
                    } else {
                        var v12 string = "\n\t"
                        v1 = (v1 >> (((^(v2))) & 63))
                        if (false || false) {
                            var v13 uint16 = 0
                            if true {
                                v12 = v12
                                if (552 < -133) {
                                    var v14 int = 0
                                    v15 := GenS2{F0: 0, F1: 1}
                                    _ = v14
                                    _ = v15
                                } else {
                                    switch 394 {
                                    case -776:
                                        for fi16 := 0; fi16 < 17; fi16++ {
                                            v2 -= 0
                                            if ((1e10 + ((1.5 - 0.0) + 1e10)) == ((0.1 / 0.1) * (1.5 / 3.14159))) {
                                                fmt.Println(v2, v1)
                                            } else {
                                                _ = len(v5)
                                                v12 = v12
                                            }
                                            v13 = v13
                                            _ = fi16
                                        }
                                    case -1:
                                        var v17 float32 = (1e10 + (1e10 - 0.1))
                                        _ = v17
                                    case 32768:
                                        v13 = 1
                                    }
                                }
                            }
                            if false {
                                v18 := GenS3{F0: v13, F1: 1, F2: v2}
                                v1 -= 42
                                v2 += (^(v2))
                                _ = v18
                            }
                            _ = v13
                        } else {
                            var v19 string = v12
                            switch v12 {
                            case "hello":
                                if len(v3) > 0 {
                                    v3[0] = 685
                                }
                            case "\n\t":
                                for fi20 := 0; fi20 < 13; fi20++ {
                                    var v21 uint32 = 503
                                    v12 = ""
                                    var v22 float64 = ((1.5 - 0.0) - (-2.25 * 0.0))
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                            }
                            _ = v19
                        }
                        _ = v12
                    }
                case "\n\t":
                    for fi23 := 0; fi23 < 11; fi23++ {
                        for fi24 := 0; fi24 < 8; fi24++ {
                            v25 := GenS2{F0: 334, F1: 2147483647}
                            v2 += 152
                            _ = len(v5)
                            _ = fi24
                            _ = v25
                        }
                        v2 += v2
                        _ = fi23
                    }
                default:
                    fmt.Println(v2, v1)
                }
            case 128:
                v5 = append(v5, ("" + "a"))
            case -2147483648:
                v1 = (^(v1))
            }
            fmt.Println(v2, v1)
        } else {
            if true {
                _ = len(v3)
                var v26 uint64 = 695
                if false {
                    v1 = (-(v1))
                    switch 9223372036854775807 {
                    case 811:
                        switch 128 {
                        case 379:
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                        case 1:
                            if true {
                                v3 = append(v3, 128)
                                v26 = v26
                                v27 := GenS1{F0: "hello", F1: 2, F2: 2147483647}
                                _ = v27
                            }
                        case 0:
                            if (true || true) {
                                v2 += (v2 | v2)
                            }
                        default:
                            switch "foo bar" {
                            case "":
                                if (v2 < (v2 >> ((v2) & 63))) {
                                    var v28 uint64 = (^(v26))
                                    for fi29 := 0; fi29 < 11; fi29++ {
                                        var v30 uint64 = 0
                                        for fi31 := 0; fi31 < 1; fi31++ {
                                            var v32 int8 = 0
                                            switch fi29 {
                                            case 32768:
                                                if (-260 >= int32(v2)) {
                                                    if (v30 < v26) {
                                                        fmt.Println(v2, v26)
                                                        v26 -= (^(v26))
                                                        var v33 uint32 = 2
                                                        _ = v33
                                                    } else {
                                                        if true {
                                                            var v34 int = -9223372036854775808
                                                            v1 = (v1 >> ((65535) & 63))
                                                            var v35 uint64 = (^(v28))
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                        fmt.Println(fi31, v30, v2)
                                                    }
                                                } else {
                                                    var v36 uint8 = 0
                                                    var v37 uint = (^(v2))
                                                    _ = v36
                                                    _ = v37
                                                }
                                                fmt.Println(v1, fi31, v30)
                                            case 9223372036854775807:
                                                v38 := GenS1{F0: "foo bar", F1: 4294967295, F2: -2147483648}
                                                _ = v38
                                            default:
                                                var v39 uint = (v2 | 2)
                                                _ = v39
                                            }
                                            _ = fi31
                                            _ = v32
                                        }
                                        var v40 bool = true
                                        _ = fi29
                                        _ = v30
                                        _ = v40
                                    }
                                    _ = v28
                                } else {
                                    var v41 int8 = 127
                                    _ = v41
                                }
                            case "a":
                                var v42 float32 = (float32(1.5) + 0.0)
                                _ = v42
                            case "foo bar":
                                v2 += v2
                            default:
                                if (1e10 <= 0.0) {
                                    v4 = append(v4, 42)
                                } else {
                                    var v43 uint8 = 91
                                    _ = v43
                                }
                            }
                        }
                    case 812:
                        switch ("" + ("a" + ("" + "a"))) {
                        case "":
                            v1 = -262
                        case "a":
                            v2 = 636
                        }
                    }
                }
                _ = v26
            }
        }
    }
    fmt.Println(v2, v2)
    if false {
        for fi44 := 0; fi44 < 14; fi44++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if (!(!(608 != 2))) {
                for fi45 := 0; fi45 < 2; fi45++ {
                    var v46 uint32 = 65535
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi45
                    _ = v46
                }
                var v47 float64 = -2.25
                v48 := GenS4{F0: 100}
                _ = v47
                _ = v48
            }
            _ = fi44
        }
        for fi49 := 0; fi49 < 17; fi49++ {
            fmt.Println(v1, fi49, v2)
            _ = fi49
        }
    } else {
        switch 0 {
        case -545:
            for fi50 := 0; fi50 < 7; fi50++ {
                v2 = v2
                var v51 float32 = (1e10 / (0.1 * 3.14159))
                _ = fi50
                _ = v51
            }
        case 9223372036854775807:
            v1 = -1
        case 32767:
            fmt.Println(v1)
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    }
    v1 = v1
    for fi52 := 0; fi52 < 13; fi52++ {
        var v53 uint8 = 42
        switch (fi52 & (fi52 >> ((18446744073709551615) & 63))) {
        case -887:
            if false {
                var v54 string = "a"
                _ = v54
            } else {
                for fi55 := 0; fi55 < 12; fi55++ {
                    v53 = (v53 << (((v2 << ((255) & 63))) & 63))
                    var v56 float64 = 3.14159
                    v1 -= (v1 & (-(v1)))
                    v56 = (float64(1.5) - 3.14159)
                    _ = fi55
                    _ = v56
                }
                v57 := GenS1{F0: (("\n\t" + "\n\t") + "a"), F1: 1, F2: 2147483647}
                _ = v57
            }
        case 696:
            v53 = v53
        }
        _ = fi52
        _ = v53
    }
    v2 = v2
    if true {
        v1 = v1
        v2 = (v2 + (v2 << ((0) & 63)))
    } else {
        for fi58 := 0; fi58 < 16; fi58++ {
            for fi59 := 0; fi59 < 4; fi59++ {
                switch "hello" {
                case "a":
                    var v60 uint = (v2 >> ((v2) & 63))
                    _ = v60
                case "foo bar":
                    switch 105 {
                    case 0:
                        v61 := GenS3{F0: 513, F1: 2, F2: (v2 >> ((74) & 63))}
                        _ = v61
                    case 32767:
                        v62 := GenS5{F0: (^(v1))}
                        _ = v62
                    case 32768:
                        if (!(fi59 >= (fi58 >> ((v2) & 63)))) {
                            v1 += (v1 >> ((v2) & 63))
                            var v63 float64 = (1.5 - ((0.1 - (1e10 / (0.0 + (float64(((-2.25 - float32(1.5)) / 1.0)) * (1.5 / 1.5))))) + (-2.25 - float64(1e10))))
                            _ = v63
                        }
                    default:
                        if true {
                            v1 += v1
                        }
                    }
                case "\n\t":
                    var v64 string = "hello"
                    _ = v64
                default:
                    if (2 > 283) {
                        if (!((v2 << (((v2 * v2)) & 63)) < (v2 ^ (^(v2))))) {
                            if (v2 < 65535) {
                                switch (fi58 * (fi59 ^ (fi58 * fi59))) {
                                case -997:
                                    var v65 uint16 = 261
                                    _ = v65
                                default:
                                    v2 = 948
                                }
                                var v66 uint32 = 42
                                _ = len(v3)
                                var v67 bool = (1 < 32767)
                                _ = v66
                                _ = v67
                            }
                            var v68 int64 = (v1 >> ((v2) & 63))
                            var v69 uint8 = 112
                            _ = v68
                            _ = v69
                        } else {
                            var v70 uint8 = 42
                            _ = v70
                        }
                        var v71 uint = v2
                        _ = v71
                    }
                }
                var v72 uint32 = 2147483648
                var v73 uint64 = 42
                var v74 uint16 = 2
                var v75 float64 = (-2.25 - (0.1 * (0.0 + (1.5 + (((0.0 + 0.0) * 0.0) - 1e10)))))
                _ = fi59
                _ = v72
                _ = v73
                _ = v74
                _ = v75
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if (int8(fi58) <= 10) {
                v1 += v1
                var v76 uint64 = 2
                var v77 int = fi58
                _ = v76
                _ = v77
            } else {
                fmt.Println(v2)
                v2 -= 2
                v2 = 65535
            }
            _ = fi58
        }
        v2 = 410
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi78 := 0; fi78 < 14; fi78++ {
        if (288 > 279) {
            switch (fi78 & (fi78 << ((922) & 63))) {
            case -9223372036854775808:
                _ = len(v3)
            case 42:
                for fi79 := 0; fi79 < 18; fi79++ {
                    var v80 int16 = 0
                    _ = fi79
                    _ = v80
                }
            default:
                v2 = (v2 >> ((v2) & 63))
            }
            fmt.Println(fi78)
            var v81 int = fi78
            _ = v81
        } else {
            if (2 <= (v2 ^ 324)) {
                v82 := GenS2{F0: 954, F1: 0}
                var v83 string = "hello"
                if len(v3) > 0 {
                    v3[0] = fi78
                }
                _ = v82
                _ = v83
            } else {
                if (2147483647 <= 4294967295) {
                    var v84 uint = (v2 * 18446744073709551615)
                    _ = v84
                }
            }
            for fi85 := 0; fi85 < 9; fi85++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi85
            }
            var v86 uint = v2
            _ = v86
        }
        _ = fi78
    }
    switch -239 {
    case 2147483647:
        if (("foo bar" + "a") <= "") {
            fmt.Println(v1, v1)
            v4 = append(v4, 675)
            v87 := GenS4{F0: -1}
            _ = v87
        } else {
            fmt.Println(v1, v1)
            fmt.Println(v2, v1)
            if false {
                var v88 uint8 = 128
                var v89 float64 = 1.5
                var v90 uint8 = (v88 * 127)
                _ = v88
                _ = v89
                _ = v90
            } else {
                fmt.Println(v2, v1)
                v91 := GenS6{F0: -601, F1: v2, F2: true}
                _ = v91
            }
        }
    case -9223372036854775808:
        if true {
            var v92 bool = (false && (42 > 0))
            _ = v92
        } else {
            for fi93 := 0; fi93 < 15; fi93++ {
                fmt.Println(v1, v2)
                _ = fi93
            }
            if false {
                var v94 uint8 = 0
                if (4294967295 > 165) {
                    var v95 uint8 = v94
                    for fi96 := 0; fi96 < 11; fi96++ {
                        var v97 float32 = ((0.1 * 1e10) + -2.25)
                        for fi98 := 0; fi98 < 15; fi98++ {
                            if (false || ((1 == 0) && false)) {
                                v99 := GenS7{F0: -70, F1: v95}
                                if (54 > 1000) {
                                    if false {
                                        var v100 float64 = 0.0
                                        _ = v100
                                    } else {
                                        fmt.Println(v2, v97)
                                        if false {
                                            var v101 int8 = 1
                                            fmt.Println(v1, fi96, fi98)
                                            _ = v101
                                        }
                                    }
                                    if (-2147483648 <= -2147483648) {
                                        switch "foo bar" {
                                        case "hello":
                                            v102 := GenS7{F0: -61, F1: 1}
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            if true {
                                                if len(v3) > 0 {
                                                    v3[0] = fi98
                                                }
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                            _ = v102
                                        case "foo bar":
                                            v103 := GenS4{F0: 100}
                                            _ = v103
                                        case "\n\t":
                                            fmt.Println(v1, v94, v95)
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        var v104 int8 = 0
                                        v95 -= v94
                                        _ = v104
                                    } else {
                                        v2 = v2
                                        v5 = append(v5, ("" + ((("hello" + "") + ("" + "")) + "a")))
                                        var v105 bool = (true && ((float32(1e10) - (3.14159 + v97)) < float32(1.5)))
                                        v94 -= 42
                                        _ = v105
                                    }
                                } else {
                                    switch 32767 {
                                    case 116:
                                        var v106 float32 = 0.1
                                        switch ("\n\t" + "foo bar") {
                                        case "a":
                                            if true {
                                                v107 := GenS6{F0: 947, F1: 1, F2: false}
                                                v108 := GenS8{F0: 42, F1: int64(v94), F2: 1}
                                                _ = v107
                                                _ = v108
                                            }
                                        case "\n\t":
                                            var v109 bool = true
                                            _ = v109
                                        case "":
                                            var v110 uint = 471
                                            _ = v110
                                        }
                                        _ = v106
                                    }
                                    switch fi96 {
                                    case -2147483648:
                                        switch 2147483647 {
                                        case -475:
                                            v3 = append(v3, (fi98 ^ fi96))
                                        case -289:
                                            switch 2147483647 {
                                            case -9223372036854775808:
                                                v111 := GenS5{F0: v1}
                                                _ = v111
                                            case 128:
                                                var v112 float64 = -2.25
                                                var v113 int64 = (^(v1))
                                                _ = v112
                                                _ = v113
                                            }
                                        default:
                                            var v114 uint = v2
                                            _ = v114
                                        }
                                    }
                                    v94 -= 1
                                }
                                _ = v99
                            }
                            var v115 int64 = v1
                            _ = fi98
                            _ = v115
                        }
                        var v116 float32 = v97
                        _ = fi96
                        _ = v116
                        _ = v97
                    }
                    _ = v95
                } else {
                    if true {
                        if (!false) {
                            fmt.Println(v2, v2)
                        }
                    }
                    v117 := GenS7{F0: 1, F1: 127}
                    _ = v117
                }
                _ = v94
            } else {
                v118 := GenS9{F0: 1}
                v1 = -9223372036854775808
                _ = v118
            }
            if (42 < 65535) {
                var v119 uint8 = 255
                _ = v119
            }
        }
    case 127:
        v2 -= v2
    default:
        var v120 int = -9223372036854775808
        _ = v120
    }
    v1 = v1
    var v121 int16 = -68
    fmt.Println("v1:", v1)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
