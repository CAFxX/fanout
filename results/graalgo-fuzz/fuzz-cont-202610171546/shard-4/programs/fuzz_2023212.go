package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint16; F2 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int; F1 uint32; F2 uint }
type GenS10 struct { F0 uint32; F1 string }
type GenS11 struct { F0 uint64; F1 int16 }

func main() {
    v1 := []uint32{0}
    var v2 uint64 = 42
    v3 := make(map[uint32]string)
    var v4 int = 616
    v5 := make(map[string]bool)
    if false {
        var v6 uint8 = 127
        for fi7 := 0; fi7 < 17; fi7++ {
            if false {
                fmt.Println(v6, v4)
                for fi8 := 0; fi8 < 2; fi8++ {
                    v6 = 134
                    var v9 float64 = 0.0
                    _ = fi8
                    _ = v9
                }
                var v10 bool = ("\n\t" != "hello")
                _ = v10
            } else {
                if (255 == 237) {
                    v11 := GenS1{F0: 503}
                    _ = v11
                }
                v12 := GenS2{F0: (116 < -44), F1: 267, F2: true}
                _ = v12
            }
            v13 := GenS3{F0: 42}
            _ = fi7
            _ = v13
        }
        switch v4 {
        case 0:
            v14 := GenS1{F0: 0}
            _ = v14
        case -1:
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 float32 = float32((0.0 / 3.14159))
                var v17 int8 = 0
                _ = fi15
                _ = v16
                _ = v17
            }
        default:
            fmt.Println(v4, v4, v2)
        }
        _ = v6
    } else {
        if false {
            var v18 uint64 = (v2 - 1)
            v18 = 598
            _ = v18
        } else {
            switch "foo bar" {
            case "a":
                v4 -= v4
            case "foo bar":
                var v19 int8 = -1
                _ = v19
            default:
                fmt.Println(v4, v2)
            }
            if false {
                var v20 float64 = 0.0
                for fi21 := 0; fi21 < 8; fi21++ {
                    var v22 uint64 = (v2 | v2)
                    _ = fi21
                    _ = v22
                }
                _ = v20
            } else {
                for fi23 := 0; fi23 < 11; fi23++ {
                    if true {
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                        switch (fi23 % (((-(v4))) | 1)) {
                        case -227:
                            delete(v5, ("hello" + "\n\t"))
                        }
                        if true {
                            switch fi23 {
                            case 872:
                                var v24 int = -2147483648
                                _ = v24
                            case 4294967295:
                                v25 := GenS4{F0: 35}
                                _ = v25
                            case -9223372036854775808:
                                v2 += (v2 * v2)
                            default:
                                if false {
                                    switch -254 {
                                    case -808:
                                        v4 += (v4 ^ (fi23 - (fi23 % ((-464) | 1))))
                                    case 42:
                                        v4 = (^(fi23))
                                    case -205:
                                        for fi26 := 0; fi26 < 13; fi26++ {
                                            v2 -= (v2 << ((42) & 63))
                                            switch ("hello" + (("hello" + "foo bar") + "foo bar")) {
                                            case "hello":
                                                v27 := GenS5{F0: 1}
                                                _ = v27
                                            }
                                            _ = fi26
                                        }
                                    default:
                                        if (false || (820 >= 0)) {
                                            var v28 bool = true
                                            _ = v28
                                        }
                                    }
                                    if false {
                                        var v29 int64 = 42
                                        switch 4294967295 {
                                        case 9223372036854775807:
                                            v29 = 42
                                        default:
                                            v4 -= (v4 * fi23)
                                        }
                                        _ = v29
                                    } else {
                                        for fi30 := 0; fi30 < 12; fi30++ {
                                            v2 = 42
                                            if true {
                                                var v31 float32 = ((0.0 + 0.0) * float32(1.5))
                                                fmt.Println(fi23, v31)
                                                v4 += (fi23 - v4)
                                                _ = v31
                                            } else {
                                                fmt.Println(v2, fi30)
                                                switch ("a" + "") {
                                                case "":
                                                    v32 := GenS3{F0: -128}
                                                    _ = v32
                                                case "foo bar":
                                                    _ = len(v1)
                                                case "hello":
                                                    v4 = 757
                                                default:
                                                    _ = len(v3)
                                                }
                                                switch "foo bar" {
                                                case "hello":
                                                    var v33 float64 = float64(0.0)
                                                    _ = v33
                                                case "\n\t":
                                                    switch (fi30 & fi30) {
                                                    case -1:
                                                        if (9223372036854775807 < -558) {
                                                            if true {
                                                                v2 = v2
                                                                var v34 bool = (0 > -583)
                                                                v35 := GenS5{F0: 42}
                                                                _ = v34
                                                                _ = v35
                                                            }
                                                            var v36 uint = 4294967295
                                                            _ = v36
                                                        }
                                                    case 478:
                                                        if (2147483647 <= 2147483647) {
                                                            switch "" {
                                                            case "hello":
                                                                fmt.Println(v4, v2, fi23)
                                                            case "":
                                                                if true {
                                                                    v37 := GenS4{F0: 1}
                                                                    if (false || (65535 > 393)) {
                                                                        var v38 int32 = -1
                                                                        _ = v38
                                                                    } else {
                                                                        if (true || true) {
                                                                            var v39 uint = 255
                                                                            if (1 == -547) {
                                                                                switch (v4 | fi23) {
                                                                                case 1:
                                                                                    switch "a" {
                                                                                    case "hello":
                                                                                        if ((230 >= 128) || (true && (217 != 127))) {
                                                                                            if false {
                                                                                                if (-32768 == -32768) {
                                                                                                    v39 -= (v39 - 1)
                                                                                                    var v40 int64 = -455
                                                                                                    _ = v40
                                                                                                }
                                                                                            } else {
                                                                                                v2 = v2
                                                                                                fmt.Println(v2, v4, v4)
                                                                                            }
                                                                                        }
                                                                                        if (0.1 >= 0.0) {
                                                                                            switch ("a" + "") {
                                                                                            case "":
                                                                                                if false {
                                                                                                    v39 += (v39 - v39)
                                                                                                    v39 = (v39 + v39)
                                                                                                    v4 = 9223372036854775807
                                                                                                }
                                                                                                fmt.Println(fi23, v4, fi30)
                                                                                                if (42 > 1) {
                                                                                                    v4 = (fi30 >> ((v39) & 63))
                                                                                                    v41 := GenS6{F0: "a"}
                                                                                                    v4 = fi23
                                                                                                    _ = v41
                                                                                                } else {
                                                                                                    var v42 int16 = -369
                                                                                                    var v43 int32 = 42
                                                                                                    var v44 int8 = 121
                                                                                                    _ = v42
                                                                                                    _ = v43
                                                                                                    _ = v44
                                                                                                }
                                                                                            case "\n\t":
                                                                                                switch "foo bar" {
                                                                                                case "\n\t":
                                                                                                    var v45 int64 = 114
                                                                                                    _ = v45
                                                                                                case "":
                                                                                                    var v46 uint = 0
                                                                                                    _ = v46
                                                                                                }
                                                                                                switch "hello" {
                                                                                                case "a":
                                                                                                    v39 += v39
                                                                                                }
                                                                                            default:
                                                                                                v2 = (^(v2))
                                                                                            }
                                                                                            var v47 int8 = 0
                                                                                            var v48 float32 = ((3.14159 * 3.14159) * -2.25)
                                                                                            _ = v47
                                                                                            _ = v48
                                                                                        } else {
                                                                                            var v49 uint16 = 32767
                                                                                            v39 = (v39 / ((v39) | 1))
                                                                                            switch ("\n\t" + (("" + "hello") + ("\n\t" + "a"))) {
                                                                                            case "":
                                                                                                var v50 int = fi30
                                                                                                _ = v50
                                                                                            case "foo bar":
                                                                                                var v51 bool = (1 != -2147483648)
                                                                                                _ = v51
                                                                                            default:
                                                                                                v1 = append(v1, 666)
                                                                                            }
                                                                                            _ = v49
                                                                                        }
                                                                                        switch "\n\t" {
                                                                                        case "hello":
                                                                                            var v52 int = 2147483647
                                                                                            _ = v52
                                                                                        case "a":
                                                                                            v53 := GenS6{F0: (("" + "foo bar") + (("hello" + ("" + ("foo bar" + ""))) + ("\n\t" + ("" + "\n\t"))))}
                                                                                            _ = v53
                                                                                        }
                                                                                    case "":
                                                                                        fmt.Println(fi30, v2)
                                                                                    }
                                                                                case -9223372036854775808:
                                                                                    v54 := GenS3{F0: -116}
                                                                                    _ = v54
                                                                                case 0:
                                                                                    var v55 int8 = -52
                                                                                    _ = v55
                                                                                }
                                                                                var v56 uint = (v39 >> ((v39) & 63))
                                                                                _ = v56
                                                                            } else {
                                                                                v2 += v2
                                                                            }
                                                                            _ = v39
                                                                        }
                                                                        if false {
                                                                            fmt.Println(fi23, fi30)
                                                                            var v57 int32 = 0
                                                                            _ = v57
                                                                        }
                                                                        if (true && (1 > 0)) {
                                                                            var v58 uint16 = 1
                                                                            _ = v58
                                                                        }
                                                                    }
                                                                    _ = v37
                                                                } else {
                                                                    if true {
                                                                        if true {
                                                                            var v59 int64 = -1
                                                                            fmt.Println(fi30)
                                                                            _ = v59
                                                                        } else {
                                                                            if false {
                                                                                fmt.Println(v2)
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            v4 = (-(fi30))
                                                            v60 := GenS5{F0: 2}
                                                            var v61 string = ("foo bar" + "")
                                                            if (1 <= v2) {
                                                                v2 = 215
                                                                v2 = 4294967295
                                                            }
                                                            switch "a" {
                                                            case "hello":
                                                                v62 := GenS1{F0: 1}
                                                                _ = v62
                                                            default:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                fmt.Println(fi23, fi30, fi30)
                                                                var v63 int8 = 64
                                                                _ = v63
                                                            }
                                                            _ = v60
                                                            _ = v61
                                                        }
                                                    case 2147483647:
                                                        if len(v1) > 0 {
                                                            v1[0] = 65535
                                                        }
                                                    default:
                                                        delete(v5, "\n\t")
                                                    }
                                                    var v64 uint8 = 0
                                                    var v65 float32 = 0.0
                                                    v65 = v65
                                                    v66 := GenS2{F0: false, F1: 0, F2: false}
                                                    switch (fi23 % ((fi30) | 1)) {
                                                    case 127:
                                                        if (fi23 > 1) {
                                                            v1 = append(v1, 1)
                                                            v64 += v64
                                                            if true {
                                                                var v67 int8 = -113
                                                                if ((!(349 != 0)) || ((2147483647 < 2147483648) || (false && true))) {
                                                                    switch ("\n\t" + "") {
                                                                    case "\n\t":
                                                                        var v68 bool = (!(((42 > 1) || ((v2 - 42) != uint64(v64))) || (((-128 == (v67 % (((v67 * v67)) | 1))) && false) || true)))
                                                                        _ = v68
                                                                    default:
                                                                        var v69 uint = 42
                                                                        _ = v69
                                                                    }
                                                                    _ = len(v1)
                                                                    v70 := v5[("hello" + "hello")]
                                                                    _ = v70
                                                                }
                                                                switch -341 {
                                                                case 1:
                                                                    v4 -= 32768
                                                                case 32768:
                                                                    v64 += (v64 - v64)
                                                                case 0:
                                                                    v2 = 804
                                                                default:
                                                                    if (!(((v64 + (v64 - v64)) != v64) && (898 != 322))) {
                                                                        var v71 int64 = 0
                                                                        var v72 float32 = 0.1
                                                                        _ = v71
                                                                        _ = v72
                                                                    } else {
                                                                        fmt.Println(v64, v4, v2)
                                                                        var v73 uint8 = 0
                                                                        var v74 int32 = -777
                                                                        _ = v73
                                                                        _ = v74
                                                                    }
                                                                }
                                                                _ = v67
                                                            }
                                                        }
                                                    case 32768:
                                                        switch (fi30 % ((fi30) | 1)) {
                                                        case 642:
                                                            var v75 float64 = (0.0 / (1e10 * 1e10))
                                                            v64 = v64
                                                            var v76 int32 = 42
                                                            _ = v75
                                                            _ = v76
                                                        default:
                                                            v4 = v4
                                                        }
                                                    case 42:
                                                        switch "foo bar" {
                                                        case "":
                                                            if (true && true) {
                                                                var v77 int = -692
                                                                if (true && true) {
                                                                    var v78 bool = (v65 == (((((v65 - (v65 * v65)) + v65) * v65) / v65) / (((v65 / (1e10 - (v65 * (1.5 * (((v65 + 0.1) + float32(-2.25)) - v65))))) - 0.0) - (0.1 + (v65 - 1.5)))))
                                                                    v79 := GenS4{F0: 1}
                                                                    _ = v78
                                                                    _ = v79
                                                                } else {
                                                                    v2 += (v2 | (v2 % (((^(v2))) | 1)))
                                                                    var v80 int32 = 1
                                                                    v4 = fi30
                                                                    _ = v80
                                                                }
                                                                switch 4294967295 {
                                                                case 128:
                                                                    switch v4 {
                                                                    case -385:
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                    case -9223372036854775808:
                                                                        var v81 float64 = (((0.0 - float64(3.14159)) * 0.0) / (0.1 / 1e10))
                                                                        _ = v81
                                                                    case 42:
                                                                        if true {
                                                                            var v82 float64 = (-2.25 / float64(((float32(0.1) + v65) * v65)))
                                                                            var v83 int = 4294967295
                                                                            if true {
                                                                                var v84 int8 = -6
                                                                                if (65535 != 4294967295) {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = 2147483647
                                                                                    }
                                                                                }
                                                                                if false {
                                                                                    var v85 float32 = 1.5
                                                                                    _ = v85
                                                                                } else {
                                                                                    switch fi23 {
                                                                                    case 1:
                                                                                        var v86 int64 = 114
                                                                                        v86 += (v86 | (^(v86)))
                                                                                        _ = v86
                                                                                    case -788:
                                                                                        if (1 == 2147483647) {
                                                                                            var v87 float32 = v65
                                                                                            if true {
                                                                                                switch int(v64) {
                                                                                                case -694:
                                                                                                    var v88 uint16 = 376
                                                                                                    _ = v88
                                                                                                case 488:
                                                                                                    v89 := GenS4{F0: v84}
                                                                                                    _ = v89
                                                                                                case 128:
                                                                                                    var v90 int16 = -1
                                                                                                    _ = v90
                                                                                                default:
                                                                                                    if true {
                                                                                                        v4 = (fi23 % ((v4) | 1))
                                                                                                        fmt.Println(v65, v87, fi30)
                                                                                                    } else {
                                                                                                        v91 := GenS4{F0: -1}
                                                                                                        _ = v91
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(v84, fi30)
                                                                                                switch ("foo bar" + ("" + "foo bar")) {
                                                                                                case "hello":
                                                                                                    v65 -= float32(v82)
                                                                                                case "a":
                                                                                                    if false {
                                                                                                        var v92 uint32 = 2147483648
                                                                                                        v87 = v65
                                                                                                        _ = v92
                                                                                                    } else {
                                                                                                        var v93 float64 = float64(v65)
                                                                                                        var v94 float32 = 0.0
                                                                                                        _ = v93
                                                                                                        _ = v94
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    v84 -= (v84 & (^(v84)))
                                                                                                }
                                                                                                fmt.Println(v2, fi23)
                                                                                                v87 -= float32(1.5)
                                                                                            } else {
                                                                                                var v95 int32 = 42
                                                                                                _ = v95
                                                                                            }
                                                                                            v65 = (0.1 / float32(1.5))
                                                                                            _ = v87
                                                                                        }
                                                                                    case 128:
                                                                                        fmt.Println(fi30, v2, v4)
                                                                                    default:
                                                                                        if true {
                                                                                            var v96 string = ("\n\t" + "\n\t")
                                                                                            _ = v96
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v77, v82, v83)
                                                                                    if false {
                                                                                        var v97 int32 = 1
                                                                                        v3[2147483647] = "\n\t"
                                                                                        v2 = v2
                                                                                        _ = v97
                                                                                    } else {
                                                                                        if (false || false) {
                                                                                            switch v83 {
                                                                                            case -239:
                                                                                                v98 := GenS1{F0: 809}
                                                                                                _ = v98
                                                                                            case 32767:
                                                                                                v99 := v5["foo bar"]
                                                                                                _ = v99
                                                                                            case 843:
                                                                                                v100 := GenS7{F0: (v84 + -1)}
                                                                                                _ = v100
                                                                                            }
                                                                                            v101 := GenS8{F0: 2}
                                                                                            _ = v101
                                                                                        } else {
                                                                                            if ((v64 / ((127) | 1)) > (v64 ^ (v64 / ((128) | 1)))) {
                                                                                                var v102 uint8 = v64
                                                                                                v103 := GenS6{F0: "hello"}
                                                                                                var v104 float64 = v82
                                                                                                _ = v102
                                                                                                _ = v103
                                                                                                _ = v104
                                                                                            } else {
                                                                                                v4 = v4
                                                                                                switch "\n\t" {
                                                                                                case "hello":
                                                                                                    v1 = append(v1, 2147483647)
                                                                                                case "foo bar":
                                                                                                    v105 := GenS9{F0: (fi30 ^ fi23), F1: 42, F2: 4294967295}
                                                                                                    _ = v105
                                                                                                }
                                                                                            }
                                                                                            var v106 uint32 = 266
                                                                                            _ = v106
                                                                                        }
                                                                                        v82 += v82
                                                                                    }
                                                                                }
                                                                                if false {
                                                                                    v65 -= v65
                                                                                    fmt.Println(v77, fi23)
                                                                                    v107 := GenS1{F0: 42}
                                                                                    _ = v107
                                                                                } else {
                                                                                    if ((!true) && false) {
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        fmt.Println(v84, v82, v83)
                                                                                    } else {
                                                                                        switch int(v64) {
                                                                                        case 9223372036854775807:
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                        case 0:
                                                                                            v2 = 42
                                                                                        default:
                                                                                            v2 = v2
                                                                                        }
                                                                                        var v108 float64 = v82
                                                                                        v5[""] = (true || (1 <= 232))
                                                                                        _ = v108
                                                                                    }
                                                                                }
                                                                                v64 = (v64 << ((42) & 63))
                                                                                _ = v84
                                                                            }
                                                                            _ = v82
                                                                            _ = v83
                                                                        } else {
                                                                            fmt.Println(v2)
                                                                        }
                                                                    default:
                                                                        var v109 int = (^(fi30))
                                                                        _ = v109
                                                                    }
                                                                }
                                                                _ = v77
                                                            }
                                                        case "hello":
                                                            v65 -= ((v65 - -2.25) / ((v65 / 0.1) / 3.14159))
                                                        }
                                                    }
                                                    _ = v64
                                                    _ = v65
                                                    _ = v66
                                                case "foo bar":
                                                    if (!false) {
                                                        if (3.14159 > 3.14159) {
                                                            v4 = 4294967295
                                                        }
                                                        var v110 float64 = (1.5 / 1.5)
                                                        _ = v110
                                                    } else {
                                                        var v111 uint64 = 849
                                                        var v112 float32 = -2.25
                                                        v112 += (v112 - 0.0)
                                                        var v113 uint32 = 2
                                                        _ = v111
                                                        _ = v112
                                                        _ = v113
                                                    }
                                                default:
                                                    v114 := GenS1{F0: 0}
                                                    var v115 float64 = (1.5 + float64(float32(0.1)))
                                                    v4 = v4
                                                    _ = v114
                                                    _ = v115
                                                }
                                            }
                                            _ = fi30
                                        }
                                        var v116 float32 = (1.5 - 1.5)
                                        var v117 int32 = 1
                                        _ = v116
                                        _ = v117
                                    }
                                    v4 = (-(v4))
                                }
                            }
                        }
                    } else {
                        switch v4 {
                        case 32768:
                            for fi118 := 0; fi118 < 14; fi118++ {
                                v119 := GenS1{F0: 42}
                                var v120 float64 = 1e10
                                _ = fi118
                                _ = v119
                                _ = v120
                            }
                        case -2147483648:
                            v4 -= (fi23 / ((472) | 1))
                        case 128:
                            fmt.Println(v2)
                        }
                        var v121 string = "foo bar"
                        for fi122 := 0; fi122 < 6; fi122++ {
                            if (0.1 <= -2.25) {
                                if (2147483647 >= 241) {
                                    v3[4294967295] = (v121 + (v121 + v121))
                                    v123 := GenS8{F0: 1}
                                    _ = v123
                                } else {
                                    var v124 uint = 4294967295
                                    v1 = append(v1, 427)
                                    if ((v124 | 883) <= (v124 & 431)) {
                                        var v125 int64 = -9223372036854775808
                                        fmt.Println(v124)
                                        var v126 uint64 = v2
                                        _ = v125
                                        _ = v126
                                    } else {
                                        var v127 uint16 = 997
                                        _ = v127
                                    }
                                    _ = v124
                                }
                            } else {
                                var v128 int8 = 42
                                _ = len(v1)
                                fmt.Println(v121, v121, v2)
                                v129 := GenS10{F0: 0, F1: v121}
                                _ = v128
                                _ = v129
                            }
                            _ = fi122
                        }
                        _ = v121
                    }
                    _ = fi23
                }
            }
            var v130 uint = 113
            _ = v130
        }
    }
    if (!(v2 < v2)) {
        var v131 float64 = 0.1
        v132 := GenS11{F0: (v2 & 0), F1: 520}
        _ = v131
        _ = v132
    }
    v133 := GenS10{F0: 42, F1: ("foo bar" + "\n\t")}
    if false {
        v134 := GenS2{F0: false, F1: 42, F2: (-32768 >= 0)}
        _ = v134
    } else {
        v4 = v4
        fmt.Println(v2, v4)
        if (!(2 == 32767)) {
            v135 := GenS11{F0: v2, F1: 570}
            _ = v135
        } else {
            for fi136 := 0; fi136 < 9; fi136++ {
                v4 = -811
                _ = fi136
            }
        }
    }
    delete(v3, 2147483648)
    v4 = 456
    if len(v1) > 0 {
        v1[0] = 955
    }
    for fi137 := 0; fi137 < 10; fi137++ {
        var v138 uint8 = 1
        fmt.Println(v4)
        _ = fi137
        _ = v138
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v133
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
