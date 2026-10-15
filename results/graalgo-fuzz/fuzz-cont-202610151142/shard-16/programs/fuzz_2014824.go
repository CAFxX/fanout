package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS5 struct { F0 bool; F1 uint8; F2 uint }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 int8; F1 string }
type GenS8 struct { F0 int; F1 string }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint16; F1 int64 }
type GenS11 struct { F0 int8; F1 string }
type GenS12 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS13 struct { F0 string; F1 uint8; F2 int16 }
type GenS14 struct { F0 uint; F1 uint8; F2 int8 }
type GenS15 struct { F0 int32; F1 uint; F2 uint }
type GenS16 struct { F0 bool; F1 uint8 }

func main() {
    var v1 uint8 = 2
    var v2 int16 = 32767
    var v3 int = 4294967295
    var v4 int16 = v2
    v5 := GenS1{F0: true}
    v6 := GenS1{F0: (!false)}
    v7 := []int{v3}
    v8 := GenS1{F0: (!(v1 < 255))}
    switch -590 {
    case -2147483648:
        v9 := GenS2{F0: "hello", F1: -114}
        _ = v9
    default:
        var v10 uint8 = 2
        _ = v10
    }
    fmt.Println(v4, v4)
    if true {
        var v11 float32 = 0.0
        if true {
            if len(v7) > 0 {
                v7[0] = v3
            }
        }
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        if ("\n\t" == "a") {
            v2 = (v4 ^ (v2 / ((v4) | 1)))
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 float64 = 0.0
                v3 -= v3
                if false {
                    if (!true) {
                        fmt.Println(v2, v2, v4)
                    } else {
                        if true {
                            v14 += (0.0 * (v14 * v14))
                            v14 = 1.5
                        }
                    }
                    if false {
                        v15 := GenS1{F0: (!(true && true))}
                        fmt.Println(v1)
                        _ = v15
                    } else {
                        if (true && false) {
                            switch "" {
                            case "a":
                                if true {
                                    v4 += v2
                                }
                            case "\n\t":
                                if false {
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch fi12 {
                                        case -172:
                                            v16 := GenS2{F0: "foo bar", F1: -38}
                                            _ = v16
                                        case 127:
                                            if true {
                                                v2 += 1
                                            }
                                        case -9223372036854775808:
                                            v4 = v4
                                        default:
                                            if false {
                                                var v17 int32 = 0
                                                _ = v17
                                            } else {
                                                v7 = append(v7, (fi12 + fi12))
                                                v7 = append(v7, fi13)
                                                v18 := GenS3{F0: 722, F1: v4}
                                                _ = v18
                                            }
                                        }
                                    case "a":
                                        fmt.Println(fi12, v14, v4)
                                    }
                                }
                            }
                            if false {
                                if false {
                                    if len(v7) > 0 {
                                        v7[0] = (v3 & v3)
                                    }
                                    var v19 float32 = float32(((v14 + 1.5) + v14))
                                    v7 = append(v7, (^(fi13)))
                                    _ = v19
                                }
                            }
                            if false {
                                if true {
                                    if (42 <= 4294967295) {
                                        var v20 int16 = (v4 | 501)
                                        _ = v20
                                    } else {
                                        if true {
                                            switch 216 {
                                            case -1:
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                            case 127:
                                                v3 += (fi12 ^ (v3 >> ((18446744073709551615) & 63)))
                                            default:
                                                switch ("a" + ((("" + "") + "foo bar") + "foo bar")) {
                                                case "foo bar":
                                                    if true {
                                                        v2 = (-(v2))
                                                        if false {
                                                            v7 = append(v7, (v3 + 767))
                                                            fmt.Println(v4, v14, fi12)
                                                        }
                                                        var v21 uint16 = 0
                                                        _ = v21
                                                    }
                                                case "\n\t":
                                                    fmt.Println(v4, v14)
                                                default:
                                                    switch (^(fi13)) {
                                                    case 1:
                                                        v1 += v1
                                                        if false {
                                                            v4 = 0
                                                            if (2147483648 == 2) {
                                                                var v22 uint32 = 65535
                                                                fmt.Println(v1, fi13)
                                                                v1 += uint8(v22)
                                                                _ = v22
                                                            } else {
                                                                v3 -= v3
                                                                switch fi12 {
                                                                case 32768:
                                                                    fmt.Println(v4, v1, v3)
                                                                case -2147483648:
                                                                    if ((^(v1)) != v1) {
                                                                        v3 = (-(fi12))
                                                                        var v23 float32 = 0.1
                                                                        var v24 uint64 = 415
                                                                        v25 := GenS1{F0: false}
                                                                        _ = v23
                                                                        _ = v24
                                                                        _ = v25
                                                                    }
                                                                case 1:
                                                                    var v26 uint8 = v1
                                                                    _ = v26
                                                                }
                                                                var v27 uint16 = 1
                                                                _ = v27
                                                            }
                                                            v28 := GenS4{F0: true, F1: 2, F2: 1}
                                                            _ = v28
                                                        }
                                                        fmt.Println(fi12, v3)
                                                    case 127:
                                                        if false {
                                                            v3 += v3
                                                            v3 += (-(fi12))
                                                            if (true && true) {
                                                                v4 = (v2 | int16(fi12))
                                                            } else {
                                                                var v29 string = "a"
                                                                v14 += v14
                                                                v2 = 0
                                                                _ = v29
                                                            }
                                                            switch ("a" + "foo bar") {
                                                            case "a":
                                                                switch (v3 - -70) {
                                                                case 1:
                                                                    v7 = append(v7, 414)
                                                                case 307:
                                                                    fmt.Println(v2)
                                                                default:
                                                                    if (((-9223372036854775808 == 899) && (170 < v1)) || false) {
                                                                        switch fi12 {
                                                                        case 71:
                                                                            v14 = v14
                                                                        }
                                                                    } else {
                                                                        var v30 int32 = 1
                                                                        v31 := GenS1{F0: false}
                                                                        if false {
                                                                            _ = len(v7)
                                                                            v32 := GenS5{F0: true, F1: (v1 % ((1) | 1)), F2: 732}
                                                                            _ = v32
                                                                        }
                                                                        _ = v30
                                                                        _ = v31
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                v2 += (v2 / (((v4 / ((-1) | 1))) | 1))
                                                            }
                                                        } else {
                                                            fmt.Println(v3, v3)
                                                            v33 := GenS4{F0: (1.5 < -2.25), F1: uint64(v1), F2: 32767}
                                                            _ = v33
                                                        }
                                                    }
                                                }
                                            }
                                        } else {
                                            if true {
                                                if true {
                                                    if (((fi12 < (fi13 * -1)) && ((-989 != (-(v2))) || true)) || true) {
                                                        if true {
                                                            if false {
                                                                switch "" {
                                                                case "":
                                                                    var v34 float64 = float64(-2.25)
                                                                    _ = v34
                                                                case "hello":
                                                                    v14 = ((v14 / float64(-2.25)) - v14)
                                                                default:
                                                                    v35 := GenS4{F0: (!(("a" + "") < ("foo bar" + ("" + "hello")))), F1: 16, F2: 2}
                                                                    _ = v35
                                                                }
                                                                var v36 uint = 255
                                                                _ = v36
                                                            }
                                                        } else {
                                                            var v37 int16 = v2
                                                            var v38 float32 = 0.0
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                        var v39 string = ("" + ("hello" + "\n\t"))
                                                        _ = v39
                                                    }
                                                } else {
                                                    var v40 uint16 = 65535
                                                    _ = v40
                                                }
                                                if false {
                                                    if true {
                                                        v41 := GenS2{F0: ("foo bar" + ("foo bar" + "foo bar")), F1: -1}
                                                        if ("" != ("a" + "foo bar")) {
                                                            v42 := GenS6{F0: "hello", F1: 232}
                                                            if false {
                                                                v43 := GenS7{F0: -128, F1: "hello"}
                                                                switch (v3 / (((v3 - (^(v3)))) | 1)) {
                                                                case -400:
                                                                    v2 -= (v4 & (-(v4)))
                                                                case 404:
                                                                    v14 = v14
                                                                default:
                                                                    if (!false) {
                                                                        var v44 uint8 = (^(v1))
                                                                        if len(v7) > 0 {
                                                                            v7[0] = (v3 % ((93) | 1))
                                                                        }
                                                                        v4 = (-(v4))
                                                                        _ = v44
                                                                    } else {
                                                                        var v45 float32 = ((((1e10 + (((-2.25 + 1e10) - 0.1) - 0.1)) * 3.14159) - (0.1 - -2.25)) * (0.0 / -2.25))
                                                                        fmt.Println(v45, v2)
                                                                        _ = v45
                                                                    }
                                                                }
                                                                _ = v43
                                                            } else {
                                                                v46 := GenS5{F0: false, F1: (v1 & 194), F2: 1}
                                                                if false {
                                                                    v14 = v14
                                                                } else {
                                                                    fmt.Println(v3, v14)
                                                                    if len(v7) >= 2 {
                                                                        v7 = v7[1:len(v7)]
                                                                    }
                                                                    var v47 int = (^(fi12))
                                                                    v3 -= 53
                                                                    var v48 bool = false
                                                                    _ = v47
                                                                    _ = v48
                                                                }
                                                                if len(v7) >= 2 {
                                                                    v7 = v7[1:len(v7)]
                                                                }
                                                                _ = v46
                                                            }
                                                            var v49 float64 = 0.0
                                                            _ = v42
                                                            _ = v49
                                                        }
                                                        _ = v41
                                                    }
                                                    var v50 int64 = 645
                                                    _ = v50
                                                } else {
                                                    var v51 int32 = 2147483647
                                                    v3 += fi13
                                                    if len(v7) > 0 {
                                                        v7[0] = (fi13 >> ((1) & 63))
                                                    }
                                                    if (1.5 < float32((3.14159 + (((1.5 - 0.1) / v14) / (v14 - (float64(0.1) - ((((v14 / (((((v14 / float64(1e10)) - float64(0.0)) - 0.0) + float64(float32((v14 - float64((0.0 + float32((float64((((((float32(0.0) * ((((1e10 - -2.25) + 1.5) * (-2.25 - 1.5)) - 1e10)) * 1.5) * 0.1) - 0.0) * (float32((((((float64(-2.25) / (v14 + (1e10 / 1e10))) - (float64(1.5) - (float64((1.5 / 1.0)) + (((1e10 - v14) * v14) / (v14 + v14))))) + 1e10) / v14) * -2.25)) + 0.0))) - (v14 * v14))))))))) * v14)) - float64(1e10)) * float64(-2.25)) - (v14 - v14)))))))) {
                                                        var v52 int16 = v2
                                                        if (255 > 42) {
                                                            v53 := GenS6{F0: "hello", F1: 255}
                                                            var v54 float64 = (v14 / (0.1 - (v14 + (v14 / 1e10))))
                                                            _ = v53
                                                            _ = v54
                                                        }
                                                        _ = v52
                                                    } else {
                                                        fmt.Println(v14, v3)
                                                        if false {
                                                            fmt.Println(fi12)
                                                            if (3.14159 <= (1.5 + float32(1e10))) {
                                                                var v55 uint16 = 440
                                                                if (true || true) {
                                                                    if len(v7) >= 2 {
                                                                        v7 = v7[1:len(v7)]
                                                                    }
                                                                    switch (v3 >> ((1) & 63)) {
                                                                    case 4294967295:
                                                                        v56 := GenS5{F0: ((^(v51)) <= (v51 % ((v51) | 1))), F1: (v1 ^ (^(v1))), F2: 4294967295}
                                                                        _ = v56
                                                                    case -390:
                                                                        v14 += v14
                                                                    }
                                                                    fmt.Println(v3, v3, v3)
                                                                } else {
                                                                    fmt.Println(v2, v2)
                                                                }
                                                                if false {
                                                                    _ = len(v7)
                                                                }
                                                                _ = v55
                                                            } else {
                                                                fmt.Println(v2)
                                                            }
                                                        }
                                                        var v57 uint32 = 123
                                                        _ = v57
                                                    }
                                                    var v58 uint64 = 4294967295
                                                    var v59 uint8 = 60
                                                    fmt.Println(v2, fi13, v2)
                                                    _ = v51
                                                    _ = v58
                                                    _ = v59
                                                }
                                                switch (fi13 | 128) {
                                                case 978:
                                                    v14 -= v14
                                                case -2147483648:
                                                    if (926 <= 65535) {
                                                        fmt.Println(v2)
                                                        v4 = v4
                                                    }
                                                default:
                                                    switch fi13 {
                                                    case 1:
                                                        var v60 string = "hello"
                                                        _ = v60
                                                    }
                                                }
                                                var v61 uint32 = 2
                                                _ = v61
                                            } else {
                                                var v62 int = (v3 + fi12)
                                                _ = v62
                                            }
                                            v2 = (^(v2))
                                        }
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                    }
                                    var v63 int64 = 0
                                    v64 := GenS4{F0: true, F1: 836, F2: 255}
                                    fmt.Println(v4)
                                    _ = len(v7)
                                    var v65 uint32 = 65535
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                }
                                switch (fi13 >> ((255) & 63)) {
                                case 1:
                                    var v66 bool = false
                                    _ = v66
                                case 32767:
                                    switch "hello" {
                                    case "hello":
                                        if (2147483647 > (fi12 + fi13)) {
                                            if len(v7) > 0 {
                                                v7[0] = (fi13 / (((v3 * v3)) | 1))
                                            }
                                            fmt.Println(v2, fi13, v1)
                                        } else {
                                            if (v1 >= (v1 << ((1) & 63))) {
                                                if true {
                                                    switch 1 {
                                                    case 32768:
                                                        switch "foo bar" {
                                                        case "":
                                                            var v67 float64 = float64(-2.25)
                                                            _ = v67
                                                        case "a":
                                                            v3 = 4294967295
                                                        case "\n\t":
                                                            var v68 float64 = v14
                                                            _ = v68
                                                        }
                                                    case 819:
                                                        v69 := GenS6{F0: "a", F1: 55}
                                                        _ = v69
                                                    case 2147483647:
                                                        v1 = 2
                                                    default:
                                                        if true {
                                                            if false {
                                                                v70 := GenS1{F0: (3.14159 < (float32(v14) * 0.1))}
                                                                v4 = (v4 ^ v2)
                                                                var v71 uint32 = 2147483648
                                                                var v72 float64 = (v14 - v14)
                                                                var v73 uint16 = 42
                                                                _ = v70
                                                                _ = v71
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                            var v74 float32 = 1e10
                                                            _ = v74
                                                        } else {
                                                            fmt.Println(fi13)
                                                        }
                                                    }
                                                    switch ("a" + ("foo bar" + "\n\t")) {
                                                    case "":
                                                        var v75 int8 = 100
                                                        _ = v75
                                                    case "a":
                                                        v1 -= 42
                                                    }
                                                }
                                                var v76 bool = (!false)
                                                var v77 float64 = v14
                                                fmt.Println(fi12, fi13, v1)
                                                if false {
                                                    fmt.Println(v4)
                                                } else {
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        _ = len(v7)
                                                    case "a":
                                                        var v78 int = (v3 >> ((255) & 63))
                                                        _ = v78
                                                    case "hello":
                                                        v79 := GenS8{F0: v3, F1: ("" + ("\n\t" + ""))}
                                                        _ = v79
                                                    }
                                                    v1 = v1
                                                }
                                                v1 += (v1 & (v1 % (((^(v1))) | 1)))
                                                var v80 int16 = 1000
                                                _ = v76
                                                _ = v77
                                                _ = v80
                                            } else {
                                                var v81 uint16 = 1
                                                switch "" {
                                                case "hello":
                                                    var v82 bool = true
                                                    _ = v82
                                                }
                                                if false {
                                                    v14 = v14
                                                    v83 := GenS9{F0: ""}
                                                    switch (-(fi13)) {
                                                    case 1:
                                                        v84 := GenS10{F0: (^(v81)), F1: -730}
                                                        _ = v84
                                                    }
                                                    _ = v83
                                                } else {
                                                    fmt.Println(v4, v81, v2)
                                                    v85 := GenS10{F0: v81, F1: -1000}
                                                    _ = v85
                                                }
                                                _ = v81
                                            }
                                            switch fi12 {
                                            case 32767:
                                                v1 += v1
                                            }
                                            switch "hello" {
                                            case "":
                                                v86 := GenS1{F0: (v3 > (fi13 + 1))}
                                                _ = v86
                                            case "foo bar":
                                                v87 := GenS11{F0: 42, F1: ("a" + "hello")}
                                                var v88 uint8 = (^(v1))
                                                if true {
                                                    v2 -= v4
                                                    var v89 bool = true
                                                    if len(v7) > 0 {
                                                        v7[0] = (v3 << ((846) & 63))
                                                    }
                                                    _ = v89
                                                }
                                                _ = v87
                                                _ = v88
                                            default:
                                                fmt.Println(v14, fi12)
                                            }
                                        }
                                    }
                                }
                            } else {
                                var v90 float64 = ((-2.25 * 1.5) + v14)
                                var v91 float32 = 3.14159
                                var v92 int64 = 823
                                _ = v90
                                _ = v91
                                _ = v92
                            }
                        } else {
                            var v93 int16 = v4
                            _ = v93
                        }
                        v2 += (v4 * (v2 << ((138) & 63)))
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                } else {
                    fmt.Println(v2, v1)
                }
                _ = fi13
                _ = v14
            }
            v94 := GenS4{F0: false, F1: 183, F2: 42}
            _ = v94
        }
        _ = fi12
    }
    v95 := GenS4{F0: ((v3 | v3) >= v3), F1: 242, F2: 2}
    fmt.Println(v1)
    var v96 float64 = float64(0.1)
    v7 = append(v7, (v3 / (((v3 ^ v3)) | 1)))
    v3 += v3
    v97 := GenS12{F0: 47, F1: 0, F2: 1}
    v3 -= (-(v3))
    for fi98 := 0; fi98 < 6; fi98++ {
        if true {
            v1 += v1
            switch ("hello" + ("" + "a")) {
            case "\n\t":
                if (2 < 0) {
                    v2 += v2
                    fmt.Println(v4)
                }
            case "":
                if (((1.5 * (1.5 - 3.14159)) / ((-2.25 - (1e10 - 1e10)) / 1.5)) <= 1e10) {
                    v99 := GenS2{F0: "", F1: 100}
                    if true {
                        var v100 uint16 = 828
                        _ = v100
                    } else {
                        v101 := GenS11{F0: -111, F1: "hello"}
                        v4 -= v2
                        switch 479 {
                        case -1:
                            v3 = v3
                        case -648:
                            fmt.Println(v96)
                        case 0:
                            v96 = v96
                        }
                        _ = v101
                    }
                    _ = len(v7)
                    _ = v99
                } else {
                    v4 = (v4 << ((0) & 63))
                    if false {
                        if false {
                            v7 = append(v7, fi98)
                            if len(v7) > 0 {
                                v7[0] = 4294967295
                            }
                            var v102 int64 = 682
                            _ = v102
                        }
                        var v103 float64 = (float64(0.0) * 0.1)
                        _ = v103
                    } else {
                        for fi104 := 0; fi104 < 1; fi104++ {
                            if false {
                                v105 := GenS2{F0: "", F1: -44}
                                _ = v105
                            } else {
                                switch v3 {
                                case -2147483648:
                                    var v106 int = (-(v3))
                                    _ = v106
                                case 486:
                                    if (!false) {
                                        _ = len(v7)
                                        var v107 int8 = 100
                                        _ = v107
                                    } else {
                                        v1 -= (^(v1))
                                        if ((1 >= 1) && (18446744073709551615 < 0)) {
                                            fmt.Println(v3, v1)
                                            var v108 uint64 = 4294967295
                                            v109 := GenS2{F0: "a", F1: 42}
                                            _ = v108
                                            _ = v109
                                        }
                                    }
                                case 798:
                                    var v110 uint = 585
                                    v111 := GenS13{F0: "\n\t", F1: (v1 + (^(v1))), F2: v4}
                                    v4 -= (v2 >> ((v110) & 63))
                                    _ = v110
                                    _ = v111
                                default:
                                    var v112 float32 = 0.0
                                    _ = v112
                                }
                            }
                            var v113 uint16 = 42
                            v114 := GenS10{F0: (v113 + uint16(v1)), F1: 1}
                            _ = fi104
                            _ = v113
                            _ = v114
                        }
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        v3 += (fi98 >> ((255) & 63))
                    }
                    for fi115 := 0; fi115 < 8; fi115++ {
                        v7 = append(v7, fi98)
                        switch fi98 {
                        case -1:
                            v2 = v2
                        case 128:
                            var v116 int32 = 2147483647
                            _ = v116
                        default:
                            v117 := GenS14{F0: 911, F1: 0, F2: -3}
                            _ = v117
                        }
                        var v118 uint8 = 128
                        switch "a" {
                        case "hello":
                            switch ("a" + (("hello" + ("" + ("hello" + ""))) + "foo bar")) {
                            case "hello":
                                if true {
                                    if false {
                                        if true {
                                            if (731 > 42) {
                                                var v119 int16 = 42
                                                var v120 int64 = 9223372036854775807
                                                _ = v119
                                                _ = v120
                                            }
                                            v1 = v1
                                        } else {
                                            v3 -= v3
                                        }
                                        _ = len(v7)
                                        v1 = v1
                                    }
                                } else {
                                    v121 := GenS15{F0: -1, F1: 353, F2: 2}
                                    var v122 int8 = 35
                                    _ = v121
                                    _ = v122
                                }
                            case "a":
                                v96 = (float64((0.1 + 1e10)) - -2.25)
                                var v123 uint = 1
                                _ = v123
                            case "\n\t":
                                v124 := GenS16{F0: true, F1: 128}
                                _ = v124
                            default:
                                v2 = (^(v2))
                            }
                        case "a":
                            var v125 int = (^(fi98))
                            _ = v125
                        default:
                            v2 = v4
                        }
                        _ = fi115
                        _ = v118
                    }
                }
            }
            var v126 uint = 1
            _ = v126
        }
        _ = fi98
    }
    v4 = (v4 | -515)
    switch 32768 {
    case 1:
        if false {
            v96 -= (1e10 * float64((1.5 - 0.1)))
            fmt.Println(v3)
        }
    case 42:
        v7 = append(v7, (-(v3)))
    default:
        for fi127 := 0; fi127 < 10; fi127++ {
            v128 := GenS7{F0: 127, F1: "foo bar"}
            if (false && (1 < 42)) {
                for fi129 := 0; fi129 < 3; fi129++ {
                    fmt.Println(fi129, v1, v4)
                    v1 = v1
                    _ = fi129
                }
                v130 := GenS7{F0: 42, F1: "foo bar"}
                _ = v130
            }
            _ = fi127
            _ = v128
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v95
    _ = v96
    _ = v97
}
