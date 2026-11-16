package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint16; F1 int16; F2 int64 }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint }
type GenS6 struct { F0 uint64; F1 int32; F2 int8 }
type GenS7 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS8 struct { F0 int16; F1 uint16; F2 bool }
type GenS9 struct { F0 int32; F1 uint }
type GenS10 struct { F0 uint; F1 uint32; F2 int16 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int32; F1 int64 }
type GenS13 struct { F0 string; F1 uint8 }
type GenS14 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS15 struct { F0 uint; F1 uint }
type GenS16 struct { F0 bool; F1 uint8; F2 int16 }
type GenS17 struct { F0 uint64; F1 int64; F2 int8 }
type GenS18 struct { F0 int }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 uint32 }
type GenS22 struct { F0 int32; F1 string; F2 int64 }
type GenS23 struct { F0 string; F1 int64 }
type GenS24 struct { F0 int; F1 int; F2 int }
type GenS25 struct { F0 bool; F1 int }

func main() {
    var v1 int32 = -2147483648
    var v2 float32 = (-2.25 * 1e10)
    var v3 uint = 1
    v2 += 3.14159
    if false {
        fmt.Println(v1, v1, v3)
        switch "foo bar" {
        case "":
            for fi4 := 0; fi4 < 9; fi4++ {
                var v5 int16 = -32768
                _ = fi4
                _ = v5
            }
        case "foo bar":
            var v6 string = "hello"
            _ = v6
        case "hello":
            switch "foo bar" {
            case "hello":
                v7 := GenS1{F0: -1}
                _ = v7
            case "":
                var v8 uint32 = 988
                _ = v8
            default:
                v2 += (1e10 + v2)
            }
        default:
            var v9 bool = ((4294967295 >= 42) && true)
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 9; fi10++ {
            switch ("foo bar" + "foo bar") {
            case "\n\t":
                v1 -= (v1 - (v1 % (((v1 * -831)) | 1)))
            default:
                var v11 uint8 = 246
                _ = v11
            }
            var v12 float64 = 1.5
            var v13 float64 = (v12 + -2.25)
            _ = fi10
            _ = v12
            _ = v13
        }
        v1 = 894
    }
    v14 := []int{0, 32767, -2147483648, 2147483647}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    v15 := GenS1{F0: -1}
    var v16 int64 = 42
    for fi17 := 0; fi17 < 14; fi17++ {
        for fi18 := 0; fi18 < 1; fi18++ {
            switch "foo bar" {
            case "a":
                switch "a" {
                case "a":
                    fmt.Println(v2, v16)
                case "":
                    if false {
                        v16 -= v16
                        var v19 string = ("\n\t" + ("foo bar" + "a"))
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        _ = v19
                    }
                default:
                    var v20 uint8 = 127
                    _ = v20
                }
            case "":
                var v21 int64 = (v16 % ((0) | 1))
                _ = v21
            default:
                if (true && (1e10 < -2.25)) {
                    v16 -= (v16 >> (((^(v3))) & 63))
                    if len(v14) > 0 {
                        v14[0] = fi17
                    }
                    v14 = append(v14, 42)
                }
                switch fi17 {
                case 128:
                    switch "foo bar" {
                    case "foo bar":
                        v22 := GenS1{F0: v16}
                        switch fi18 {
                        case 42:
                            var v23 uint16 = 1
                            _ = v23
                        default:
                            v1 += v1
                            if (116 > 2) {
                                var v24 uint8 = 127
                                if false {
                                    var v25 int32 = (v1 >> ((370) & 63))
                                    if (3.14159 != 3.14159) {
                                        var v26 uint = v3
                                        var v27 uint8 = 179
                                        switch "a" {
                                        case "":
                                            switch "hello" {
                                            case "foo bar":
                                                var v28 int64 = -1
                                                _ = v28
                                            case "":
                                                v29 := GenS1{F0: (v16 | (v16 >> ((v26) & 63)))}
                                                _ = v29
                                            case "\n\t":
                                                var v30 float32 = (((float32(-2.25) * 0.1) + v2) * 0.1)
                                                _ = v30
                                            default:
                                                if true {
                                                    var v31 bool = (!false)
                                                    v14 = append(v14, fi18)
                                                    _ = v31
                                                }
                                            }
                                        default:
                                            var v32 float64 = (1e10 * 3.14159)
                                            _ = v32
                                        }
                                        var v33 int = fi17
                                        _ = v26
                                        _ = v27
                                        _ = v33
                                    }
                                    v2 -= ((0.0 - v2) * v2)
                                    _ = v25
                                } else {
                                    var v34 uint = v3
                                    var v35 float32 = v2
                                    var v36 int16 = 848
                                    var v37 bool = (42 != 901)
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                                v3 += (v3 ^ 33)
                                switch "foo bar" {
                                case "\n\t":
                                    v16 = (v16 + v16)
                                default:
                                    switch (-(fi18)) {
                                    case 42:
                                        v16 = (v16 << (((v3 / (((v3 | v3)) | 1))) & 63))
                                    case 604:
                                        switch "hello" {
                                        case "":
                                            var v38 int8 = -82
                                            _ = v38
                                        default:
                                            v3 += v3
                                        }
                                    default:
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                    }
                                    if (true && true) {
                                        switch "" {
                                        case "hello":
                                            if len(v14) >= 2 {
                                                v14 = v14[1:len(v14)]
                                            }
                                        case "foo bar":
                                            v2 = 3.14159
                                        case "\n\t":
                                            if (!false) {
                                                if (2 > 440) {
                                                    fmt.Println(fi17, v24, v16)
                                                    fmt.Println(v1, v3, v16)
                                                    if len(v14) > 0 {
                                                        v14[0] = 1
                                                    }
                                                    v39 := GenS2{F0: uint32(v3), F1: ((1e10 * 1e10) == v2)}
                                                    fmt.Println(v1, v1, v24)
                                                    if (true || false) {
                                                        if true {
                                                            var v40 int64 = v16
                                                            _ = v40
                                                        }
                                                        switch ("hello" + (("" + "hello") + "foo bar")) {
                                                        case "a":
                                                            v3 = (v3 + v3)
                                                        case "hello":
                                                            if (v16 >= v16) {
                                                                v14 = append(v14, (fi17 & (fi17 - fi17)))
                                                                var v41 int16 = 42
                                                                v42 := GenS3{F0: v24, F1: true}
                                                                _ = v41
                                                                _ = v42
                                                            } else {
                                                                switch "" {
                                                                case "\n\t":
                                                                    v24 = (^(v24))
                                                                    var v43 float32 = (-2.25 + 0.0)
                                                                    if (!(("a" > ("hello" + "foo bar")) && true)) {
                                                                        v3 -= 2
                                                                        v14 = append(v14, 2147483647)
                                                                        var v44 uint8 = (v24 ^ (v24 + (v24 & (v24 & v24))))
                                                                        _ = v44
                                                                    }
                                                                    if true {
                                                                        var v45 uint = (v3 | 44)
                                                                        var v46 uint = (v45 / ((717) | 1))
                                                                        if ((float32(1.5) > v43) || (-2.25 >= -2.25)) {
                                                                            switch "\n\t" {
                                                                            case "a":
                                                                                var v47 int16 = 42
                                                                                _ = v47
                                                                            case "\n\t":
                                                                                if len(v14) > 0 {
                                                                                    v14[0] = (-(fi18))
                                                                                }
                                                                            case "":
                                                                                if false {
                                                                                    fmt.Println(v16, fi17)
                                                                                    fmt.Println(v43, v16, v3)
                                                                                    v14 = append(v14, -9223372036854775808)
                                                                                } else {
                                                                                    v48 := GenS4{F0: 42, F1: -32768, F2: (v16 ^ -242)}
                                                                                    fmt.Println(v16, v24, fi18)
                                                                                    v46 = (^(v46))
                                                                                    _ = v48
                                                                                }
                                                                            default:
                                                                                fmt.Println(v46)
                                                                            }
                                                                        }
                                                                        _ = v45
                                                                        _ = v46
                                                                    } else {
                                                                        var v49 uint16 = 255
                                                                        if (true && ((0.1 / 0.1) == 3.14159)) {
                                                                            v50 := GenS3{F0: (v24 % ((127) | 1)), F1: true}
                                                                            if len(v14) >= 2 {
                                                                                v14 = v14[1:len(v14)]
                                                                            }
                                                                            _ = v50
                                                                        } else {
                                                                            v51 := GenS5{F0: 24, F1: 18446744073709551615, F2: v3}
                                                                            _ = v51
                                                                        }
                                                                        var v52 int = (fi18 << (((v3 >> ((42) & 63))) & 63))
                                                                        _ = v49
                                                                        _ = v52
                                                                    }
                                                                    fmt.Println(v24)
                                                                    _ = v43
                                                                case "":
                                                                    if (true || (894 <= 18446744073709551615)) {
                                                                        if (!((-(v1)) != -932)) {
                                                                            if (!(2 <= 481)) {
                                                                                v53 := GenS6{F0: 0, F1: 1, F2: -128}
                                                                                v16 -= v16
                                                                                _ = v53
                                                                            } else {
                                                                                var v54 uint64 = 1
                                                                                switch (fi17 << (((^(v3))) & 63)) {
                                                                                case 4294967295:
                                                                                    _ = len(v14)
                                                                                case 627:
                                                                                    fmt.Println(v3, v24, v2)
                                                                                    var v55 int = (fi17 >> ((v3) & 63))
                                                                                    v56 := GenS6{F0: (^(v54)), F1: (v1 & v1), F2: 127}
                                                                                    var v57 uint = 2
                                                                                    if len(v14) >= 2 {
                                                                                        v14 = v14[1:len(v14)]
                                                                                    }
                                                                                    _ = v55
                                                                                    _ = v56
                                                                                    _ = v57
                                                                                default:
                                                                                    v1 -= int32(v3)
                                                                                }
                                                                                fmt.Println(v1, fi17, v3)
                                                                                _ = v54
                                                                            }
                                                                        } else {
                                                                            var v58 uint32 = 2147483647
                                                                            _ = v58
                                                                        }
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            var v59 bool = ((v2 + ((v2 / v2) / v2)) <= v2)
                                                                            _ = v59
                                                                        case "\n\t":
                                                                            v60 := GenS6{F0: 669, F1: (v1 >> (((^(v3))) & 63)), F2: 100}
                                                                            _ = v60
                                                                        default:
                                                                            var v61 uint32 = 42
                                                                            _ = v61
                                                                        }
                                                                    } else {
                                                                        var v62 int64 = (v16 & v16)
                                                                        if false {
                                                                            var v63 int8 = 65
                                                                            _ = v63
                                                                        } else {
                                                                            var v64 uint32 = 2147483647
                                                                            if (547 == v62) {
                                                                                if false {
                                                                                    var v65 int64 = v16
                                                                                    _ = v65
                                                                                } else {
                                                                                    v1 = (v1 - v1)
                                                                                }
                                                                                var v66 int64 = 42
                                                                                if false {
                                                                                    if (false || true) {
                                                                                        v66 += 1
                                                                                        fmt.Println(v66)
                                                                                        fmt.Println(fi17, v3, fi18)
                                                                                        if true {
                                                                                            if ((float64((v2 / 3.14159)) < (float64(3.14159) + 0.0)) || true) {
                                                                                                if true {
                                                                                                    var v67 uint32 = v64
                                                                                                    if len(v14) > 0 {
                                                                                                        v14[0] = fi17
                                                                                                    }
                                                                                                    if len(v14) >= 2 {
                                                                                                        v14 = v14[1:len(v14)]
                                                                                                    }
                                                                                                    _ = len(v14)
                                                                                                    switch "" {
                                                                                                    case "foo bar":
                                                                                                        v68 := GenS4{F0: 255, F1: -1, F2: -9223372036854775808}
                                                                                                        var v69 uint16 = 209
                                                                                                        _ = v68
                                                                                                        _ = v69
                                                                                                    }
                                                                                                    _ = v67
                                                                                                }
                                                                                                v70 := GenS1{F0: (v62 + (v16 / (((v16 - 9223372036854775807)) | 1)))}
                                                                                                v71 := GenS7{F0: 680, F1: 309, F2: (v64 / (((^(v64))) | 1))}
                                                                                                _ = v70
                                                                                                _ = v71
                                                                                            }
                                                                                            switch "hello" {
                                                                                            case "foo bar":
                                                                                                fmt.Println(v24)
                                                                                            }
                                                                                            v24 -= v24
                                                                                        } else {
                                                                                            var v72 float64 = ((1e10 * 0.0) - float64(0.1))
                                                                                            var v73 float32 = 1.5
                                                                                            _ = v72
                                                                                            _ = v73
                                                                                        }
                                                                                    }
                                                                                    switch "" {
                                                                                    case "hello":
                                                                                        if (!(1.5 < 3.14159)) {
                                                                                            fmt.Println(v2, v66)
                                                                                            v3 = (v3 / (((^(v3))) | 1))
                                                                                            var v74 int = fi17
                                                                                            _ = v74
                                                                                        }
                                                                                    case "a":
                                                                                        switch -9223372036854775808 {
                                                                                        case 780:
                                                                                            fmt.Println(fi17, v62, v24)
                                                                                        default:
                                                                                            fmt.Println(v1)
                                                                                        }
                                                                                    case "foo bar":
                                                                                        var v75 float32 = v2
                                                                                        _ = v75
                                                                                    default:
                                                                                        v76 := GenS8{F0: -32768, F1: 32767, F2: true}
                                                                                        _ = v76
                                                                                    }
                                                                                }
                                                                                if (!false) {
                                                                                    if true {
                                                                                        fmt.Println(v3)
                                                                                        fmt.Println(v24)
                                                                                        v77 := GenS9{F0: (v1 % (((v1 * (v1 & (-(v1))))) | 1)), F1: (v3 ^ (v3 * v3))}
                                                                                        v66 -= (v66 >> (((v3 ^ (v3 & (^(v3))))) & 63))
                                                                                        _ = v77
                                                                                    } else {
                                                                                        v2 = (v2 + 1e10)
                                                                                        v24 += (v24 + v24)
                                                                                        v64 = (v64 << ((v3) & 63))
                                                                                        v78 := GenS7{F0: (^(v3)), F1: v64, F2: (v64 % (((v64 % ((4294967295) | 1))) | 1))}
                                                                                        _ = v78
                                                                                    }
                                                                                    if true {
                                                                                        v16 -= 1
                                                                                        v14 = append(v14, fi18)
                                                                                        v66 -= (v66 >> (((v3 / ((v3) | 1))) & 63))
                                                                                    } else {
                                                                                        var v79 string = ("hello" + (("a" + "\n\t") + "\n\t"))
                                                                                        var v80 int16 = 1
                                                                                        v80 = (v80 << ((255) & 63))
                                                                                        _ = v79
                                                                                        _ = v80
                                                                                    }
                                                                                    if false {
                                                                                        _ = len(v14)
                                                                                        _ = len(v14)
                                                                                    } else {
                                                                                        if (uint32(v3) < v64) {
                                                                                            if false {
                                                                                                var v81 float32 = float32(1e10)
                                                                                                var v82 uint64 = 18446744073709551615
                                                                                                _ = v81
                                                                                                _ = v82
                                                                                            }
                                                                                            if false {
                                                                                                v16 = (v62 | (v16 + v66))
                                                                                                _ = len(v14)
                                                                                            }
                                                                                        }
                                                                                        if len(v14) >= 2 {
                                                                                            v14 = v14[1:len(v14)]
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v66
                                                                            } else {
                                                                                if (!false) {
                                                                                    if false {
                                                                                        v83 := GenS10{F0: (^(v3)), F1: (v64 & v64), F2: 32767}
                                                                                        _ = v83
                                                                                    } else {
                                                                                        var v84 int8 = -24
                                                                                        _ = v84
                                                                                    }
                                                                                    var v85 string = ("\n\t" + "hello")
                                                                                    _ = v85
                                                                                }
                                                                                if false {
                                                                                    if true {
                                                                                        v86 := GenS11{F0: -1}
                                                                                        var v87 uint = v3
                                                                                        v24 -= v24
                                                                                        fmt.Println(v2)
                                                                                        _ = v86
                                                                                        _ = v87
                                                                                    } else {
                                                                                        v88 := GenS10{F0: (v3 + v3), F1: 2147483647, F2: -815}
                                                                                        v1 = v1
                                                                                        _ = v88
                                                                                    }
                                                                                    v14 = append(v14, -9223372036854775808)
                                                                                }
                                                                                v89 := GenS1{F0: v62}
                                                                                _ = v89
                                                                            }
                                                                            _ = v64
                                                                        }
                                                                        v14 = append(v14, (fi18 + fi17))
                                                                        _ = v62
                                                                    }
                                                                }
                                                                v16 += v16
                                                            }
                                                        default:
                                                            fmt.Println(fi17)
                                                        }
                                                        var v90 float64 = 0.0
                                                        _ = v90
                                                    }
                                                    _ = v39
                                                }
                                            }
                                        }
                                    } else {
                                        if (!true) {
                                            switch "" {
                                            case "\n\t":
                                                v16 = v16
                                            default:
                                                var v91 uint16 = 898
                                                _ = v91
                                            }
                                        } else {
                                            switch ("hello" + ("\n\t" + "")) {
                                            case "hello":
                                                var v92 uint64 = 2
                                                var v93 float64 = 1e10
                                                v94 := GenS9{F0: v1, F1: 0}
                                                var v95 int = 745
                                                _ = v92
                                                _ = v93
                                                _ = v94
                                                _ = v95
                                            case "foo bar":
                                                var v96 string = "foo bar"
                                                _ = v96
                                            default:
                                                fmt.Println(v1, fi17)
                                            }
                                            v3 = v3
                                            _ = len(v14)
                                        }
                                        if (("hello" + "hello") > "hello") {
                                            v97 := GenS12{F0: (v1 - v1), F1: (v16 << ((994) & 63))}
                                            v98 := GenS5{F0: 1, F1: 728, F2: v3}
                                            _ = v97
                                            _ = v98
                                        } else {
                                            if (2 != v24) {
                                                switch 32768 {
                                                case -152:
                                                    v2 = (-2.25 * 3.14159)
                                                    switch (fi17 % ((fi17) | 1)) {
                                                    case 2147483647:
                                                        switch fi18 {
                                                        case -506:
                                                            fmt.Println(v24, v24, v2)
                                                        case -873:
                                                            v99 := GenS11{F0: 648}
                                                            _ = v99
                                                        }
                                                    case 9223372036854775807:
                                                        switch fi18 {
                                                        case -2147483648:
                                                            v100 := GenS13{F0: "hello", F1: v24}
                                                            v101 := GenS14{F0: 0, F1: -2147483648, F2: v24}
                                                            _ = v100
                                                            _ = v101
                                                        default:
                                                            fmt.Println(v24, fi17, v1)
                                                        }
                                                    case -2147483648:
                                                        var v102 string = ""
                                                        _ = v102
                                                    default:
                                                        switch (fi17 % ((32768) | 1)) {
                                                        case 659:
                                                            var v103 uint = (v3 - v3)
                                                            if (!true) {
                                                                if ((v24 + (v24 >> (((v3 % ((v3) | 1))) & 63))) >= (v24 % ((1) | 1))) {
                                                                    if true {
                                                                        v104 := GenS2{F0: 827, F1: (!true)}
                                                                        _ = v104
                                                                    }
                                                                    fmt.Println(fi18, v3, v16)
                                                                    var v105 string = "a"
                                                                    _ = v105
                                                                }
                                                                fmt.Println(v1, v1, v24)
                                                                fmt.Println(fi18, fi17, v3)
                                                            }
                                                            var v106 string = ""
                                                            _ = v103
                                                            _ = v106
                                                        default:
                                                            if (v2 > v2) {
                                                                var v107 uint8 = 1
                                                                var v108 bool = false
                                                                _ = v107
                                                                _ = v108
                                                            }
                                                        }
                                                    }
                                                case 32768:
                                                    v16 = v16
                                                default:
                                                    switch (-(fi17)) {
                                                    case 221:
                                                        if true {
                                                            v109 := GenS15{F0: 4294967295, F1: (v3 + 119)}
                                                            v110 := GenS16{F0: (!true), F1: (v24 | (v24 & v24)), F2: 42}
                                                            if (true || false) {
                                                                if (true && (2 <= 2147483648)) {
                                                                    v16 = v16
                                                                } else {
                                                                    v24 = v24
                                                                    if true {
                                                                        v111 := GenS17{F0: 42, F1: (v16 + v16), F2: 0}
                                                                        _ = v111
                                                                    }
                                                                    v112 := GenS3{F0: v24, F1: (!false)}
                                                                    _ = v112
                                                                }
                                                                fmt.Println(fi18, fi17, fi18)
                                                            }
                                                            _ = v109
                                                            _ = v110
                                                        }
                                                    case 1:
                                                        v14 = append(v14, (fi18 << ((v3) & 63)))
                                                    case 117:
                                                        if true {
                                                            var v113 int16 = int16(v24)
                                                            if (128 < 211) {
                                                                switch "" {
                                                                case "hello":
                                                                    var v114 int8 = 103
                                                                    _ = v114
                                                                case "":
                                                                    var v115 int64 = (^(v16))
                                                                    _ = v115
                                                                case "foo bar":
                                                                    v14 = append(v14, fi17)
                                                                default:
                                                                    switch ("foo bar" + "foo bar") {
                                                                    case "a":
                                                                        var v116 int16 = -134
                                                                        _ = v116
                                                                    }
                                                                }
                                                                var v117 uint = (v3 >> ((v3) & 63))
                                                                switch "" {
                                                                case "foo bar":
                                                                    var v118 float64 = -2.25
                                                                    _ = v118
                                                                case "hello":
                                                                    if len(v14) > 0 {
                                                                        v14[0] = (fi17 | fi17)
                                                                    }
                                                                case "\n\t":
                                                                    var v119 uint = v117
                                                                    _ = v119
                                                                default:
                                                                    v113 = v113
                                                                }
                                                                _ = v117
                                                            }
                                                            var v120 int16 = (v113 % ((v113) | 1))
                                                            v121 := GenS8{F0: int16(v16), F1: 253, F2: (!(v24 < 0))}
                                                            _ = v113
                                                            _ = v120
                                                            _ = v121
                                                        } else {
                                                            fmt.Println(v16, v2, v24)
                                                        }
                                                    default:
                                                        v16 -= v16
                                                    }
                                                }
                                                v122 := GenS3{F0: (v24 | 128), F1: (3.14159 <= v2)}
                                                v16 += (v16 << ((4294967295) & 63))
                                                _ = v122
                                            } else {
                                                v24 = (v24 + 127)
                                            }
                                            switch "hello" {
                                            case "hello":
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v123 uint64 = 0
                                                    _ = v123
                                                case "a":
                                                    if true {
                                                        fmt.Println(v3, v2, v16)
                                                        if ((v2 == v2) || (42 > v1)) {
                                                            var v124 int64 = v16
                                                            _ = len(v14)
                                                            v125 := GenS18{F0: fi17}
                                                            _ = v124
                                                            _ = v125
                                                        }
                                                        switch "" {
                                                        case "foo bar":
                                                            v14 = append(v14, 128)
                                                        default:
                                                            v16 += v16
                                                        }
                                                        if false {
                                                            if true {
                                                                var v126 int8 = 100
                                                                var v127 uint32 = 42
                                                                _ = len(v14)
                                                                _ = v126
                                                                _ = v127
                                                            }
                                                            if true {
                                                                v16 += (v16 + (v16 / ((-9223372036854775808) | 1)))
                                                                if ((fi18 ^ (^(fi18))) == (fi18 & fi17)) {
                                                                    v128 := GenS19{F0: ((v1 / (((v1 % ((v1) | 1))) | 1)) >= v1)}
                                                                    _ = v128
                                                                } else {
                                                                    var v129 int = (fi17 ^ fi18)
                                                                    _ = v129
                                                                }
                                                                var v130 int64 = (v16 >> ((512) & 63))
                                                                _ = v130
                                                            }
                                                            if len(v14) > 0 {
                                                                v14[0] = (fi17 / ((fi18) | 1))
                                                            }
                                                        } else {
                                                            v2 = (1.5 - (-2.25 + (float32((-2.25 * 0.1)) + v2)))
                                                        }
                                                    } else {
                                                        var v131 uint32 = 1
                                                        v132 := GenS15{F0: 65535, F1: v3}
                                                        _ = v131
                                                        _ = v132
                                                    }
                                                case "hello":
                                                    var v133 uint16 = 0
                                                    _ = v133
                                                }
                                            case "\n\t":
                                                if true {
                                                    var v134 uint = (v3 - 788)
                                                    _ = v134
                                                }
                                            }
                                        }
                                        _ = len(v14)
                                        var v135 uint32 = 42
                                        _ = v135
                                    }
                                }
                                _ = v24
                            }
                            v16 = v16
                        }
                        v136 := GenS6{F0: 4294967295, F1: v1, F2: -1}
                        _ = v136
                        _ = v22
                    case "\n\t":
                        var v137 int32 = (-(v1))
                        _ = v137
                    case "a":
                        var v138 string = ""
                        _ = v138
                    }
                case 127:
                    fmt.Println(v1)
                case 509:
                    switch "foo bar" {
                    case "\n\t":
                        var v139 uint = (v3 * (v3 - 799))
                        _ = v139
                    case "a":
                        var v140 uint64 = 800
                        _ = v140
                    case "":
                        v2 -= v2
                    default:
                        if false {
                            if true {
                                v1 -= v1
                                v16 = v16
                                v2 += (v2 - ((float32(0.0) / -2.25) - v2))
                            }
                            fmt.Println(v2)
                            if len(v14) > 0 {
                                v14[0] = (fi18 ^ (fi18 % ((fi18) | 1)))
                            }
                        }
                    }
                default:
                    var v141 uint = v3
                    _ = v141
                }
                if (!true) {
                    v14 = append(v14, 2147483647)
                    var v142 uint8 = 209
                    v1 += v1
                    _ = v142
                }
            }
            switch 0 {
            case 1:
                if false {
                    fmt.Println(v2, v2)
                    if (((-(v16)) > (v16 ^ (v16 + 42))) && (2 == 2147483647)) {
                        switch fi17 {
                        case 552:
                            if len(v14) > 0 {
                                v14[0] = 127
                            }
                        case -723:
                            if ((687 != 0) || (!true)) {
                                var v143 uint8 = 72
                                v143 = v143
                                if (0 < 0) {
                                    _ = len(v14)
                                    v143 += (^(v143))
                                    v144 := GenS20{F0: true}
                                    _ = v144
                                } else {
                                    if (v16 <= (v16 * int64(v1))) {
                                        v16 = v16
                                        v143 -= (v143 / ((1) | 1))
                                    } else {
                                        fmt.Println(v143, v3, v16)
                                    }
                                }
                                v145 := GenS21{F0: 42}
                                _ = v143
                                _ = v145
                            }
                        }
                        v2 += v2
                        v146 := GenS22{F0: 2147483647, F1: "a", F2: (v16 - (v16 | v16))}
                        _ = v146
                    } else {
                        if false {
                            var v147 int16 = 1000
                            _ = v147
                        }
                        v3 -= (v3 * v3)
                        if (338 > 32767) {
                            v148 := GenS15{F0: (v3 >> (((v3 % ((v3) | 1))) & 63)), F1: 65535}
                            var v149 bool = true
                            v150 := GenS14{F0: 2, F1: 42, F2: uint8(v3)}
                            _ = v148
                            _ = v149
                            _ = v150
                        } else {
                            fmt.Println(v1)
                            var v151 int = (fi17 << ((v3) & 63))
                            v152 := GenS23{F0: "a", F1: v16}
                            _ = v151
                            _ = v152
                        }
                    }
                }
            case -657:
                v14 = append(v14, (fi18 << (((v3 / (((v3 | v3)) | 1))) & 63)))
            }
            _ = fi18
        }
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        var v153 uint = (v3 & (v3 ^ (v3 ^ (v3 + v3))))
        _ = fi17
        _ = v153
    }
    var v154 uint = (v3 >> ((v3) & 63))
    switch 32768 {
    case 32767:
        switch 530 {
        case 32768:
            if (true && (v16 < 682)) {
                fmt.Println(v1, v154)
                fmt.Println(v2, v1)
            } else {
                for fi155 := 0; fi155 < 15; fi155++ {
                    for fi156 := 0; fi156 < 20; fi156++ {
                        if true {
                            v157 := GenS24{F0: (-(fi156)), F1: 738, F2: (fi155 - (^(fi156)))}
                            _ = v157
                        }
                        var v158 bool = ((v2 * 1e10) != 3.14159)
                        _ = fi156
                        _ = v158
                    }
                    var v159 float64 = 1.5
                    _ = fi155
                    _ = v159
                }
            }
        case -2147483648:
            var v160 int8 = 42
            _ = v160
        case 185:
            var v161 int16 = 335
            _ = v161
        }
    case -319:
        for fi162 := 0; fi162 < 16; fi162++ {
            for fi163 := 0; fi163 < 2; fi163++ {
                v14 = append(v14, fi163)
                v164 := GenS25{F0: true, F1: -577}
                _ = fi163
                _ = v164
            }
            for fi165 := 0; fi165 < 12; fi165++ {
                var v166 float32 = 0.1
                var v167 int8 = int8(v1)
                _ = fi165
                _ = v166
                _ = v167
            }
            _ = fi162
        }
    default:
        v154 = v3
    }
    var v168 int32 = (v1 ^ 64)
    v169 := GenS24{F0: 116, F1: 387, F2: -9223372036854775808}
    v2 -= v2
    fmt.Println(v1, v1, v168)
    if false {
        if true {
            switch "a" {
            case "\n\t":
                v16 = v16
            case "hello":
                if len(v14) > 0 {
                    v14[0] = 127
                }
            case "":
                if (0 > 961) {
                    var v170 uint8 = 224
                    v171 := GenS24{F0: 2147483647, F1: 956, F2: 42}
                    _ = v170
                    _ = v171
                }
            default:
                if false {
                    for fi172 := 0; fi172 < 8; fi172++ {
                        switch fi172 {
                        case 0:
                            v154 = v3
                        default:
                            var v173 bool = (false || (32767 > 1000))
                            _ = v173
                        }
                        var v174 int8 = -128
                        if len(v14) > 0 {
                            v14[0] = -2147483648
                        }
                        _ = fi172
                        _ = v174
                    }
                }
            }
            var v175 int8 = 118
            _ = v175
        }
        v154 = 210
        v154 += v3
    }
    v176 := GenS14{F0: 693, F1: (-(v1)), F2: 255}
    v154 += (v3 / (((v3 / ((255) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v154:", v154)
    fmt.Println("v16:", v16)
    fmt.Println("v168:", v168)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v154
    _ = v16
    _ = v168
    _ = v169
    _ = v176
    _ = v2
    _ = v3
}
