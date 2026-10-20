package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint32; F1 bool; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 bool; F1 uint16; F2 bool }
type GenS6 struct { F0 uint8; F1 uint32; F2 string }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint32; F1 int8; F2 int16 }
type GenS11 struct { F0 uint16; F1 string; F2 uint8 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint; F1 uint32 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 string }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 uint64; F1 string; F2 uint64 }
func genFn0(p0 int8, p1 uint64) uint8 {
    v5 := []int{988, -850, 685}
    var v6 int16 = 362
    fmt.Println(v6)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 127
}
func genFn1(p0 int64, p1 bool) bool {
    p0 = (-(p0))
    p1 = (!false)
    p1 = p1
    if true {
        v7 := GenS1{F0: -9223372036854775808, F1: 109}
        v8 := GenS2{F0: 2147483647, F1: (!p1), F2: 567}
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 uint = 702
            _ = fi9
            _ = v10
        }
        _ = v7
        _ = v8
    } else {
        switch 9223372036854775807 {
        case -843:
            if p1 {
                v11 := []int{}
                p1 = true
                _ = v11
            } else {
                v12 := GenS1{F0: 4294967295, F1: 667}
                var v13 bool = ((1 < 0) || true)
                switch "" {
                case "foo bar":
                    var v14 uint64 = 955
                    _ = v14
                case "hello":
                    var v15 int16 = -1
                    _ = v15
                case "a":
                    var v16 bool = (v13 || false)
                    _ = v16
                }
                _ = v12
                _ = v13
            }
        case 32768:
            fmt.Println(p1, p1)
        case -1:
            for fi17 := 0; fi17 < 11; fi17++ {
                if p1 {
                    var v18 uint32 = 2
                    p0 = (^(p0))
                    if p1 {
                        for fi19 := 0; fi19 < 14; fi19++ {
                            if p1 {
                                var v20 int = (-(fi19))
                                v21 := []int{(fi17 / (((fi19 + int(v18))) | 1)), 9223372036854775807, v20}
                                v18 = (^(v18))
                                fmt.Println(p0)
                                _ = v20
                                _ = v21
                            }
                            v22 := []int{fi19, fi17, (fi17 % (((fi19 >> ((65535) & 63))) | 1)), -359}
                            switch "" {
                            case "foo bar":
                                v23 := GenS1{F0: -9223372036854775808, F1: 18446744073709551615}
                                _ = v23
                            case "a":
                                if true {
                                    v24 := GenS2{F0: v18, F1: (p1 && (128 <= 128)), F2: -924}
                                    _ = v24
                                } else {
                                    if p1 {
                                        if (!false) {
                                            v18 = (^(v18))
                                            switch "foo bar" {
                                            case "a":
                                                p0 -= p0
                                            case "hello":
                                                var v25 int32 = 364
                                                _ = v25
                                            case "foo bar":
                                                var v26 int32 = -295
                                                _ = v26
                                            }
                                        }
                                    }
                                }
                            case "":
                                var v27 string = "a"
                                _ = v27
                            default:
                                var v28 int32 = -2147483648
                                _ = v28
                            }
                            _ = fi19
                            _ = v22
                        }
                        var v29 uint8 = 1
                        _ = v29
                    }
                    _ = v18
                } else {
                    p1 = (!(p1 && ((!false) || p1)))
                    for fi30 := 0; fi30 < 9; fi30++ {
                        var v31 uint64 = 4294967295
                        p1 = (1e10 != 0.1)
                        var v32 uint = 18446744073709551615
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    var v33 float64 = ((1.5 * 0.1) * ((((float64((0.0 - float32(0.1))) - 0.0) / 0.1) * ((float64(-2.25) + ((0.1 / 1.0) / ((1e10 - ((0.0 - float64(((1.5 - (0.1 - -2.25)) - 1.5))) - 0.0)) / float64(1.5)))) - 1e10)) * 0.0))
                    _ = v33
                }
                if p1 {
                    fmt.Println(p1, p1)
                    v34 := GenS1{F0: fi17, F1: 4294967295}
                    _ = v34
                } else {
                    fmt.Println(p1, p0, fi17)
                    var v35 int32 = 0
                    _ = v35
                }
                _ = fi17
            }
        }
    }
    if (32767 >= 127) {
        switch "hello" {
        case "":
            p0 = p0
        }
    }
    _ = p0
    _ = p1
    return (!p1)
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 127
    v3 := make(map[string]string)
    var v4 uint32 = 1
    var v36 string = "hello"
    v4 = (v4 / ((v4) | 1))
    fmt.Println(v4, v36, v36)
    v37 := GenS2{F0: 2, F1: false, F2: 1}
    delete(v3, (v36 + ("\n\t" + (v36 + v36))))
    fmt.Println(v4)
    var v38 int32 = 42
    if (!(v36 <= (v36 + "hello"))) {
        delete(v1, (v36 + ((v36 + "a") + v36)))
        v39 := GenS3{F0: 752}
        _ = v39
    } else {
        v40 := GenS3{F0: 32767}
        if false {
            if true {
                for fi41 := 0; fi41 < 15; fi41++ {
                    for fi42 := 0; fi42 < 15; fi42++ {
                        fmt.Println(fi41)
                        _ = fi42
                    }
                    _ = fi41
                }
                for fi43 := 0; fi43 < 4; fi43++ {
                    for fi44 := 0; fi44 < 11; fi44++ {
                        fmt.Println(fi43, v38, fi43)
                        fmt.Println(v4, fi44)
                        if true {
                            var v45 int = 128
                            _ = v45
                        }
                        _ = fi44
                    }
                    v46 := GenS1{F0: fi43, F1: 18446744073709551615}
                    v47 := []int{(fi43 & -1), (-(fi43))}
                    _ = fi43
                    _ = v46
                    _ = v47
                }
                v36 = ((((v36 + "a") + v36) + ((("\n\t" + (v36 + (("\n\t" + "\n\t") + "foo bar"))) + v36) + "foo bar")) + ("hello" + v36))
            }
            var v48 uint16 = 240
            v48 -= (v48 * (v48 * uint16(v4)))
            _ = v48
        }
        for fi49 := 0; fi49 < 12; fi49++ {
            var v50 int = (fi49 % (((fi49 | (fi49 << ((170) & 63)))) | 1))
            switch v36 {
            case "hello":
                fmt.Println(v36, v2)
            case "foo bar":
                if (897 != 255) {
                    switch (v36 + ("foo bar" + ("\n\t" + v36))) {
                    case "a":
                        if true {
                            var v51 int32 = (v38 / (((v38 & -1)) | 1))
                            _ = v51
                        }
                    case "hello":
                        fmt.Println(v2, v50)
                    case "":
                        if (899 > -1) {
                            fmt.Println(v38, fi49, v38)
                            if (((1 >= 615) || false) || (v2 == (v2 - (-(v2))))) {
                                v52 := GenS4{F0: fi49, F1: 65535}
                                switch 42 {
                                case 1:
                                    var v53 int64 = -1
                                    _ = v53
                                case 127:
                                    var v54 uint16 = 231
                                    _ = v54
                                case -490:
                                    v55 := GenS2{F0: (v4 - v4), F1: ((false || true) && true), F2: -17}
                                    _ = v55
                                default:
                                    fmt.Println(v50, v2, v50)
                                }
                                _ = v52
                            }
                        }
                    default:
                        for fi56 := 0; fi56 < 19; fi56++ {
                            v36 = (((("\n\t" + v36) + "a") + "hello") + ("a" + ("\n\t" + "hello")))
                            if (203 != 1) {
                                v57 := GenS5{F0: (((1 != 428) && ((v38 >> ((2) & 63)) < v38)) || false), F1: 65535, F2: (false || false)}
                                v58 := GenS3{F0: 1}
                                _ = v57
                                _ = v58
                            }
                            var v59 int8 = v2
                            switch -52 {
                            case 871:
                                if true {
                                    if (964 == 25) {
                                        switch (v50 | 9223372036854775807) {
                                        case 4294967295:
                                            if true {
                                                if (!true) {
                                                    if ((!false) && (899 == 1000)) {
                                                        var v60 int16 = int16(v59)
                                                        v61 := GenS6{F0: 2, F1: v4, F2: v36}
                                                        _ = v60
                                                        _ = v61
                                                    } else {
                                                        v36 = "foo bar"
                                                        v62 := GenS2{F0: 81, F1: ("" != (v36 + v36)), F2: 42}
                                                        fmt.Println(v2, v2)
                                                        _ = v62
                                                    }
                                                } else {
                                                    switch (^(fi56)) {
                                                    case -403:
                                                        v36 = (v36 + "a")
                                                    case 535:
                                                        v63 := []int{(fi49 + fi56), (fi56 + (fi49 & 32768)), 42, (^(v50))}
                                                        _ = v63
                                                    }
                                                    var v64 int32 = (^(v38))
                                                    switch -121 {
                                                    case 128:
                                                        if (("foo bar" < (((("" + v36) + (v36 + "")) + (v36 + v36)) + "a")) && false) {
                                                            fmt.Println(v4, v4, fi56)
                                                            if false {
                                                                v38 += 1
                                                            } else {
                                                                _ = len(v1)
                                                                v2 = v2
                                                            }
                                                        }
                                                    }
                                                    v36 = "foo bar"
                                                    v59 = v59
                                                    _ = v64
                                                }
                                                v65 := GenS7{F0: 42, F1: v59}
                                                fmt.Println(v38, fi49)
                                                _ = v65
                                            } else {
                                                v38 -= 42
                                                var v66 int32 = (v38 << ((0) & 63))
                                                var v67 string = (v36 + v36)
                                                _ = v66
                                                _ = v67
                                            }
                                        case 9223372036854775807:
                                            if (genFn0(-1, 18446744073709551615) != 42) {
                                                var v68 int8 = v2
                                                switch v36 {
                                                case "a":
                                                    v69 := []int{(-(v50)), (fi56 * -269), (v50 | int(v68))}
                                                    _ = v69
                                                case "":
                                                    if true {
                                                        var v70 int16 = int16(v38)
                                                        _ = v70
                                                    } else {
                                                        var v71 float32 = (-2.25 * (1e10 - -2.25))
                                                        v72 := GenS6{F0: 1, F1: 2147483647, F2: ""}
                                                        v73 := GenS3{F0: -32768}
                                                        _ = v71
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                default:
                                                    if ("" == v36) {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            fmt.Println(v59, v50)
                                                            var v74 string = ((v36 + v36) + v36)
                                                            v75 := []int{}
                                                            _ = v74
                                                            _ = v75
                                                        case "hello":
                                                            switch (v36 + v36) {
                                                            case "foo bar":
                                                                v76 := GenS8{F0: 42}
                                                                _ = v76
                                                            default:
                                                                v77 := GenS4{F0: (v50 ^ (fi49 >> ((255) & 63))), F1: 255}
                                                                _ = v77
                                                            }
                                                        }
                                                        if false {
                                                            v38 -= (v38 | (v38 + v38))
                                                            var v78 int64 = 0
                                                            _ = v78
                                                        } else {
                                                            if ((0 == 479) || false) {
                                                                fmt.Println(v36)
                                                                v68 -= v2
                                                                if false {
                                                                    var v79 bool = (2 == 4294967295)
                                                                    _ = v79
                                                                }
                                                                if (true || false) {
                                                                    v80 := []int{fi56}
                                                                    if true {
                                                                        switch (fi56 | (fi56 ^ 912)) {
                                                                        case 42:
                                                                            v80 = append(v80, (fi49 >> ((495) & 63)))
                                                                        case -29:
                                                                            fmt.Println(v4, v38, fi49)
                                                                        }
                                                                        v4 += v4
                                                                        var v81 uint16 = 0
                                                                        v81 -= v81
                                                                        switch v36 {
                                                                        case "\n\t":
                                                                            var v82 uint16 = (v81 >> ((2) & 63))
                                                                            _ = v82
                                                                        case "":
                                                                            switch "" {
                                                                            case "":
                                                                                if (uint64(v81) < 2) {
                                                                                    v59 = v68
                                                                                    v59 = v2
                                                                                    if true {
                                                                                        v50 = fi56
                                                                                        switch (fi56 / (((fi56 - 988)) | 1)) {
                                                                                        case 127:
                                                                                            var v83 float64 = 0.0
                                                                                            _ = v83
                                                                                        default:
                                                                                            switch v36 {
                                                                                            case "hello":
                                                                                                v50 += (fi56 - fi49)
                                                                                            default:
                                                                                                var v84 uint = 65535
                                                                                                _ = v84
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v68, v38)
                                                                                    } else {
                                                                                        v85 := GenS2{F0: 42, F1: true, F2: 191}
                                                                                        var v86 int64 = -9223372036854775808
                                                                                        fmt.Println(v50, v81)
                                                                                        v59 = v68
                                                                                        _ = v85
                                                                                        _ = v86
                                                                                    }
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        if (v4 <= (^(v4))) {
                                                                                            if (1 < -357) {
                                                                                                var v87 int16 = 553
                                                                                                fmt.Println(v81, v36)
                                                                                                _ = v87
                                                                                            } else {
                                                                                                var v88 int8 = 42
                                                                                                if (-2.25 < (1.5 / -2.25)) {
                                                                                                    var v89 uint8 = 42
                                                                                                    switch v50 {
                                                                                                    case 32767:
                                                                                                        v68 = int8(v4)
                                                                                                    case -1:
                                                                                                        v4 -= 0
                                                                                                    }
                                                                                                    if false {
                                                                                                        v90 := GenS8{F0: 255}
                                                                                                        fmt.Println(v38)
                                                                                                        _ = v90
                                                                                                    } else {
                                                                                                        if len(v80) >= 2 {
                                                                                                            v80 = v80[1:len(v80)]
                                                                                                        }
                                                                                                        switch (("hello" + (v36 + v36)) + (v36 + ("a" + "\n\t"))) {
                                                                                                        case "a":
                                                                                                            if false {
                                                                                                                fmt.Println(v36)
                                                                                                                if len(v80) >= 2 {
                                                                                                                    v80 = v80[1:len(v80)]
                                                                                                                }
                                                                                                                fmt.Println(v36, v4, fi56)
                                                                                                                fmt.Println(fi56, v36)
                                                                                                            }
                                                                                                        default:
                                                                                                            v88 -= v88
                                                                                                            var v91 uint32 = 59
                                                                                                            v92 := GenS9{F0: (^(v4))}
                                                                                                            if true {
                                                                                                                var v93 float32 = (3.14159 / 3.14159)
                                                                                                                v50 = 1
                                                                                                                var v94 int16 = 1000
                                                                                                                v95 := GenS10{F0: 2147483647, F1: v2, F2: v94}
                                                                                                                v36 = v36
                                                                                                                v59 -= (v68 | (-(v59)))
                                                                                                                _ = v93
                                                                                                                _ = v94
                                                                                                                _ = v95
                                                                                                            }
                                                                                                            _ = v91
                                                                                                            _ = v92
                                                                                                        }
                                                                                                        var v96 float64 = 1e10
                                                                                                        _ = v96
                                                                                                    }
                                                                                                    _ = v89
                                                                                                } else {
                                                                                                    v3[("a" + "hello")] = v36
                                                                                                    v36 = "hello"
                                                                                                }
                                                                                                switch fi56 {
                                                                                                case 1:
                                                                                                    var v97 bool = ((0 <= 65535) && false)
                                                                                                    _ = v97
                                                                                                }
                                                                                                _ = v88
                                                                                            }
                                                                                            v98 := GenS11{F0: (v81 ^ v81), F1: ("hello" + v36), F2: 128}
                                                                                            v4 -= 0
                                                                                            _ = v98
                                                                                        }
                                                                                    case "hello":
                                                                                        v81 = 798
                                                                                    }
                                                                                    v59 += v59
                                                                                } else {
                                                                                    if (0.1 != -2.25) {
                                                                                        var v99 uint8 = 2
                                                                                        fmt.Println(v99, v68, v68)
                                                                                        v100 := GenS7{F0: 4294967295, F1: (^(v68))}
                                                                                        _ = v100
                                                                                        _ = v99
                                                                                    }
                                                                                }
                                                                            case "hello":
                                                                                var v101 uint64 = 18446744073709551615
                                                                                _ = v101
                                                                            default:
                                                                                v81 -= (v81 >> ((65535) & 63))
                                                                            }
                                                                        case "foo bar":
                                                                            delete(v3, "\n\t")
                                                                        default:
                                                                            var v102 bool = (!true)
                                                                            _ = v102
                                                                        }
                                                                        _ = v81
                                                                    }
                                                                    _ = v80
                                                                } else {
                                                                    switch (fi56 << ((4294967295) & 63)) {
                                                                    case 4294967295:
                                                                        if (255 != 196) {
                                                                            v1[(("a" + (v36 + ((v36 + v36) + v36))) + v36)] = (v4 - v4)
                                                                            switch ("" + v36) {
                                                                            case "a":
                                                                                if (true || true) {
                                                                                    var v103 int8 = (^(v68))
                                                                                    _ = v103
                                                                                }
                                                                            case "foo bar":
                                                                                fmt.Println(v4)
                                                                            case "":
                                                                                v104 := GenS12{F0: 1}
                                                                                _ = v104
                                                                            default:
                                                                                var v105 uint8 = 2
                                                                                _ = v105
                                                                            }
                                                                            fmt.Println(v4, v4, v38)
                                                                        } else {
                                                                            v50 = (fi56 ^ v50)
                                                                            fmt.Println(v38, fi56, v50)
                                                                            v106 := []int{fi49, fi49, v50, (fi49 << ((65535) & 63))}
                                                                            _ = v106
                                                                        }
                                                                    case 32768:
                                                                        var v107 uint32 = (^(v4))
                                                                        _ = v107
                                                                    case 1:
                                                                        v108 := GenS13{F0: 4294967295, F1: (v4 % (((v4 / (((^(v4))) | 1))) | 1))}
                                                                        _ = v108
                                                                    }
                                                                    v36 = (v36 + "\n\t")
                                                                    if ("hello" != ("hello" + (("" + ((v36 + ((v36 + v36) + "hello")) + ("a" + ""))) + v36))) {
                                                                        v109 := []int{v50, v50, 9223372036854775807, 1}
                                                                        v110 := GenS11{F0: 255, F1: v36, F2: 1}
                                                                        switch (("hello" + "foo bar") + (v36 + v36)) {
                                                                        case "hello":
                                                                            v111 := GenS14{F0: ("hello" + "\n\t")}
                                                                            _ = v111
                                                                        case "foo bar":
                                                                            var v112 uint8 = uint8(v4)
                                                                            _ = v112
                                                                        }
                                                                        var v113 float64 = ((0.1 - float64(0.0)) * (3.14159 - (0.1 + (3.14159 / 1.0))))
                                                                        var v114 string = "foo bar"
                                                                        v115 := v3[(v36 + "a")]
                                                                        _ = v109
                                                                        _ = v110
                                                                        _ = v113
                                                                        _ = v114
                                                                        _ = v115
                                                                    }
                                                                    var v116 int64 = -16
                                                                    _ = v116
                                                                }
                                                            }
                                                            switch (v50 % ((fi49) | 1)) {
                                                            case -154:
                                                                var v117 int = (fi56 >> ((255) & 63))
                                                                _ = v117
                                                            case -55:
                                                                var v118 int32 = v38
                                                                _ = v118
                                                            case 237:
                                                                if true {
                                                                    v2 += -128
                                                                    fmt.Println(v50, v38, v4)
                                                                    var v119 uint8 = 42
                                                                    _ = v119
                                                                }
                                                            }
                                                            fmt.Println(v38)
                                                        }
                                                        if ((127 <= (v68 | v59)) || true) {
                                                            if (!(-2.25 == ((-2.25 * (float32(-2.25) * ((0.1 - 1.5) * float32((((1.5 - (1.5 * 1.5)) + 1e10) + 1e10))))) - 1.5))) {
                                                                v68 -= (v59 * (v59 << ((18446744073709551615) & 63)))
                                                                if (720 == 0) {
                                                                    if (!(!false)) {
                                                                        _ = len(v3)
                                                                        if (false && true) {
                                                                            var v120 int8 = (v68 | v59)
                                                                            var v121 uint32 = v4
                                                                            _ = v120
                                                                            _ = v121
                                                                        } else {
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                var v122 uint64 = 2
                                                                                _ = v122
                                                                            case "hello":
                                                                                v123 := []int{-796, v50}
                                                                                _ = v123
                                                                            case "\n\t":
                                                                                v124 := []int{}
                                                                                _ = v124
                                                                            default:
                                                                                var v125 uint8 = 5
                                                                                _ = v125
                                                                            }
                                                                            if true {
                                                                                fmt.Println(fi56, v68)
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    fmt.Println(v68, v2)
                                                                                }
                                                                            }
                                                                        }
                                                                        switch (v50 >> ((245) & 63)) {
                                                                        case -113:
                                                                            delete(v3, "a")
                                                                        case 889:
                                                                            switch fi56 {
                                                                            case 42:
                                                                                var v126 uint8 = 127
                                                                                _ = v126
                                                                            case 32768:
                                                                                var v127 uint32 = v4
                                                                                _ = v127
                                                                            case 863:
                                                                                v38 -= int32(v50)
                                                                            }
                                                                        }
                                                                    }
                                                                    switch 128 {
                                                                    case -2147483648:
                                                                        v128 := GenS15{F0: v36}
                                                                        v50 -= 0
                                                                        var v129 int16 = -32768
                                                                        v38 = v38
                                                                        _ = v128
                                                                        _ = v129
                                                                    case 0:
                                                                        v4 += 257
                                                                    }
                                                                    v130 := GenS16{F0: -968}
                                                                    _ = v130
                                                                }
                                                            } else {
                                                                if true {
                                                                    var v131 int64 = -407
                                                                    v132 := v1[((((v36 + ("a" + "a")) + ("foo bar" + ((v36 + (((("\n\t" + "a") + "hello") + v36) + ("\n\t" + ""))) + (v36 + "")))) + v36) + "\n\t")]
                                                                    v50 = (fi56 - (^(fi56)))
                                                                    var v133 string = "a"
                                                                    var v134 int8 = v59
                                                                    _ = v131
                                                                    _ = v132
                                                                    _ = v133
                                                                    _ = v134
                                                                } else {
                                                                    if ((1e10 * 1e10) <= 1e10) {
                                                                        var v135 int = fi56
                                                                        _ = v135
                                                                    }
                                                                    v136 := []int{62}
                                                                    _ = v136
                                                                }
                                                            }
                                                            v2 = (v68 * (v59 % ((127) | 1)))
                                                            var v137 int8 = -1
                                                            _ = v137
                                                        } else {
                                                            fmt.Println(v38, v50)
                                                            v4 = v4
                                                            v50 -= (fi56 | v50)
                                                        }
                                                    } else {
                                                        v4 -= v4
                                                        v138 := GenS11{F0: 42, F1: "foo bar", F2: 42}
                                                        _ = v138
                                                    }
                                                }
                                                v139 := GenS17{F0: 2, F1: (v36 + "hello"), F2: 18446744073709551615}
                                                _ = v139
                                                _ = v68
                                            } else {
                                                v140 := []int{32768, (fi56 ^ (fi56 + -1))}
                                                fmt.Println(v36)
                                                var v141 int32 = v38
                                                _ = v140
                                                _ = v141
                                            }
                                        case 564:
                                            v59 = v2
                                        }
                                        fmt.Println(v4, v38, v38)
                                    }
                                    v142 := []int{fi49, (^(fi56))}
                                    _ = v142
                                }
                            default:
                                v143 := []int{int(v59), int(v38), fi49, v50}
                                _ = v143
                            }
                            _ = fi56
                            _ = v59
                        }
                    }
                    var v144 uint16 = 2
                    var v145 uint = 0
                    _ = v144
                    _ = v145
                }
            case "a":
                if ((646 > 371) || (2 < 255)) {
                    v146 := []int{1, v50, -2147483648, fi49}
                    _ = v146
                }
            default:
                for fi147 := 0; fi147 < 15; fi147++ {
                    fmt.Println(v4, v4)
                    switch fi147 {
                    case 122:
                        v4 = v4
                    case -9223372036854775808:
                        var v148 bool = false
                        _ = v148
                    default:
                        var v149 int64 = -743
                        _ = v149
                    }
                    delete(v3, "foo bar")
                    _ = fi147
                }
            }
            _ = fi49
            _ = v50
        }
        _ = v40
    }
    switch 1 {
    case 791:
        switch v36 {
        case "":
            v150 := []int{128}
            _ = v150
        }
    case 32767:
        var v151 uint8 = genFn0(v2, 42)
        _ = v151
    case 0:
        var v152 float32 = 0.1
        _ = v152
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
