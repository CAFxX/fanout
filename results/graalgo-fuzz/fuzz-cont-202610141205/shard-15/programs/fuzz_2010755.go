package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint64; F1 uint64; F2 int }
type GenS5 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 int; F1 int16; F2 uint }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 bool; F1 int32 }
type GenS11 struct { F0 int64; F1 bool }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int32; F1 string; F2 int32 }
type GenS15 struct { F0 uint32; F1 uint; F2 string }
type GenS16 struct { F0 bool; F1 uint; F2 int }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 bool }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint64 = 18446744073709551615
    v3 := []uint32{417}
    v4 := []string{"foo bar", (("hello" + "\n\t") + "a")}
    if false {
        v5 := GenS1{F0: 1, F1: 42, F2: 42}
        for fi6 := 0; fi6 < 4; fi6++ {
            if true {
                v1 = v1
                v2 = v2
                v7 := GenS2{F0: 1}
                _ = v7
            }
            _ = fi6
        }
        var v8 uint32 = 988
        _ = v5
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 float32 = 1e10
            _ = fi10
            _ = v11
        }
        v12 := GenS3{F0: -1, F1: 169}
        v13 := GenS3{F0: 42, F1: 32767}
        _ = fi9
        _ = v12
        _ = v13
    }
    if (true || false) {
        fmt.Println(v1, v2)
    } else {
        fmt.Println(v2, v1)
        if true {
            switch "hello" {
            case "hello":
                for fi14 := 0; fi14 < 8; fi14++ {
                    if true {
                        fmt.Println(v2)
                        var v15 uint = 343
                        for fi16 := 0; fi16 < 19; fi16++ {
                            _ = len(v3)
                            _ = fi16
                        }
                        _ = v15
                    } else {
                        for fi17 := 0; fi17 < 14; fi17++ {
                            if false {
                                if ((true || ((1 != 65535) && (true || true))) && true) {
                                    v18 := GenS4{F0: v2, F1: v2, F2: fi14}
                                    _ = v18
                                }
                                switch fi17 {
                                case 127:
                                    v1 += 822
                                case -1:
                                    v19 := GenS5{F0: 4294967295, F1: 42, F2: -2147483648}
                                    _ = v19
                                case 384:
                                    var v20 uint32 = 617
                                    _ = v20
                                default:
                                    var v21 uint8 = 147
                                    var v22 string = ("foo bar" + "\n\t")
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = v21
                                    _ = v22
                                }
                            }
                            _ = fi17
                        }
                        for fi23 := 0; fi23 < 16; fi23++ {
                            v2 -= v2
                            var v24 bool = true
                            if v24 {
                                v2 += v2
                                v25 := GenS4{F0: (v2 << ((87) & 63)), F1: (v2 % (((v2 ^ v2)) | 1)), F2: -709}
                                _ = v25
                            }
                            _ = fi23
                            _ = v24
                        }
                        switch fi14 {
                        case 9223372036854775807:
                            v3 = append(v3, 43)
                        case 1:
                            if false {
                                v2 -= 18446744073709551615
                                switch ("" + "hello") {
                                case "a":
                                    v1 = (-(v1))
                                case "hello":
                                    v26 := GenS6{F0: v2}
                                    _ = v26
                                case "":
                                    fmt.Println(v2, v1)
                                }
                                v3 = append(v3, 384)
                            } else {
                                if true {
                                    var v27 uint16 = 1
                                    _ = v27
                                }
                            }
                        }
                    }
                    _ = fi14
                }
            case "\n\t":
                var v28 int = 32767
                _ = v28
            case "foo bar":
                for fi29 := 0; fi29 < 17; fi29++ {
                    v30 := GenS4{F0: 0, F1: v2, F2: 0}
                    var v31 uint = 65535
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
            fmt.Println(v1, v1)
            var v32 int16 = 159
            _ = v32
        } else {
            v33 := GenS4{F0: (v2 * 340), F1: 0, F2: 32767}
            _ = v33
        }
        fmt.Println(v2, v2)
    }
    v2 += (v2 >> ((42) & 63))
    _ = len(v3)
    var v34 bool = ((!((float32((float64(-2.25) - 0.0)) * (1.5 * 0.1)) > 3.14159)) && (-128 != 1))
    if v34 {
        v1 = -9223372036854775808
        v1 = 1
        v3 = append(v3, 2147483648)
    } else {
        if v34 {
            var v35 uint8 = 0
            _ = v35
        }
    }
    var v36 float64 = 1e10
    v36 = 3.14159
    if (!(287 > 2147483648)) {
        var v37 uint16 = 453
        _ = v37
    }
    if len(v4) > 0 {
        v4[0] = ""
    }
    switch "\n\t" {
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    case "hello":
        fmt.Println(v1)
    default:
        if (0 < -2147483648) {
            if (0 >= 0) {
                var v38 int = 680
                v39 := GenS2{F0: 179}
                fmt.Println(v2, v1, v38)
                _ = v38
                _ = v39
            }
        } else {
            for fi40 := 0; fi40 < 12; fi40++ {
                v1 = v1
                switch 2147483647 {
                case -194:
                    for fi41 := 0; fi41 < 14; fi41++ {
                        if v34 {
                            fmt.Println(v34, v36, v34)
                            switch ("\n\t" + "foo bar") {
                            case "a":
                                fmt.Println(v34, v1, fi40)
                            default:
                                var v42 int8 = -1
                                _ = v42
                            }
                        }
                        _ = fi41
                    }
                case 2147483647:
                    if v34 {
                        for fi43 := 0; fi43 < 7; fi43++ {
                            if v34 {
                                switch ("\n\t" + "\n\t") {
                                case "a":
                                    switch fi40 {
                                    case 406:
                                        var v44 int16 = 177
                                        _ = v44
                                    default:
                                        v36 = v36
                                    }
                                case "\n\t":
                                    v34 = v34
                                case "hello":
                                    v34 = v34
                                }
                            }
                            var v45 int8 = 127
                            if (v34 || v34) {
                                if (-2.25 >= -2.25) {
                                    v45 += v45
                                    v45 += (v45 / ((v45) | 1))
                                    switch fi43 {
                                    case 1:
                                        var v46 uint64 = 2
                                        _ = v46
                                    }
                                }
                                if (!((v1 | v1) > 9223372036854775807)) {
                                    var v47 int = (^(fi40))
                                    var v48 int = (v47 << ((255) & 63))
                                    var v49 uint32 = 65535
                                    if true {
                                        v50 := GenS7{F0: (^(v49)), F1: 2147483648}
                                        var v51 uint16 = 32767
                                        _ = v50
                                        _ = v51
                                    }
                                    _ = v47
                                    _ = v48
                                    _ = v49
                                }
                            } else {
                                if true {
                                    v52 := GenS4{F0: v2, F1: v2, F2: (fi40 * (-(fi43)))}
                                    v53 := GenS5{F0: (^(v2)), F1: 0, F2: 42}
                                    _ = v52
                                    _ = v53
                                } else {
                                    var v54 bool = (-2147483648 != -2147483648)
                                    var v55 int32 = -246
                                    var v56 uint64 = v2
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                }
                                if (4294967295 >= 4294967295) {
                                    if (!(!((fi40 + 128) <= 2147483647))) {
                                        fmt.Println(fi43)
                                        var v57 string = "a"
                                        if len(v3) > 0 {
                                            v3[0] = 681
                                        }
                                        var v58 uint64 = (v2 & v2)
                                        if false {
                                            switch v57 {
                                            case "a":
                                                var v59 int = fi43
                                                _ = v59
                                            }
                                            v60 := GenS4{F0: 193, F1: (v58 << ((0) & 63)), F2: (fi40 & -1)}
                                            if (!v34) {
                                                switch ("hello" + v57) {
                                                case "":
                                                    v1 -= v1
                                                default:
                                                    fmt.Println(v2, v45, v45)
                                                }
                                                var v61 int = (fi40 * fi43)
                                                _ = v61
                                            } else {
                                                v34 = true
                                            }
                                            _ = v60
                                        } else {
                                            var v62 int64 = 0
                                            if (v1 < (v62 & v1)) {
                                                if (v36 == 1.5) {
                                                    v63 := GenS8{F0: (fi43 << ((2) & 63)), F1: -1, F2: 42}
                                                    var v64 bool = (v1 > v62)
                                                    _ = v63
                                                    _ = v64
                                                } else {
                                                    v45 = (v45 >> ((18446744073709551615) & 63))
                                                }
                                                if v34 {
                                                    v57 = ""
                                                    v65 := GenS9{F0: 0, F1: 255}
                                                    _ = v65
                                                } else {
                                                    v66 := GenS10{F0: true, F1: 445}
                                                    if (764 != 109) {
                                                        fmt.Println(fi40)
                                                        v45 += (v45 ^ (v45 - (v45 - (v45 * (v45 >> ((4294967295) & 63))))))
                                                        var v67 float64 = (1.5 * ((0.0 + v36) + v36))
                                                        _ = v67
                                                    } else {
                                                        switch (fi43 << ((2) & 63)) {
                                                        case 32767:
                                                            var v68 float32 = 1.5
                                                            _ = v68
                                                        case 4294967295:
                                                            v69 := GenS4{F0: v2, F1: (v58 | 933), F2: fi40}
                                                            _ = v69
                                                        case 32768:
                                                            if true {
                                                                v34 = v34
                                                            }
                                                        }
                                                        if (v34 && true) {
                                                            var v70 uint32 = 42
                                                            _ = v70
                                                        }
                                                        _ = len(v3)
                                                    }
                                                    _ = v66
                                                }
                                                if (2147483647 < 0) {
                                                    switch ((("foo bar" + ("\n\t" + v57)) + v57) + v57) {
                                                    case "hello":
                                                        switch (^(fi40)) {
                                                        case 32768:
                                                            switch ((v57 + v57) + v57) {
                                                            case "\n\t":
                                                                var v71 string = v57
                                                                _ = v71
                                                            }
                                                        case -1:
                                                            v3 = append(v3, 4294967295)
                                                        default:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(v45, v45)
                                                    switch 1 {
                                                    case 0:
                                                        switch fi43 {
                                                        case -1:
                                                            fmt.Println(v45, v1)
                                                        case 127:
                                                            v57 = (v57 + v57)
                                                        case -333:
                                                            if false {
                                                                v62 += -9223372036854775808
                                                            } else {
                                                                var v72 uint = 2
                                                                var v73 string = v57
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                        default:
                                                            var v74 uint64 = 18446744073709551615
                                                            _ = v74
                                                        }
                                                    case -538:
                                                        v45 = v45
                                                    case 32768:
                                                        fmt.Println(v62, v57)
                                                    }
                                                } else {
                                                    switch v57 {
                                                    case "":
                                                        v36 = 1.5
                                                    case "\n\t":
                                                        var v75 int32 = -1
                                                        _ = v75
                                                    default:
                                                        if (3.14159 == -2.25) {
                                                            fmt.Println(v36, v62)
                                                            if v34 {
                                                                fmt.Println(v62, v2, v2)
                                                                if len(v4) > 0 {
                                                                    v4[0] = ((v57 + (v57 + v57)) + v57)
                                                                }
                                                            }
                                                        }
                                                        if ((v34 || true) && v34) {
                                                            switch (((v57 + v57) + v57) + (("" + v57) + (((v57 + v57) + "") + (((v57 + v57) + v57) + ("a" + (v57 + ("\n\t" + "hello"))))))) {
                                                            case "\n\t":
                                                                v45 = v45
                                                            default:
                                                                var v76 uint32 = 65535
                                                                _ = v76
                                                            }
                                                            fmt.Println(v57, v57, fi40)
                                                        }
                                                    }
                                                    v77 := GenS2{F0: 18446744073709551615}
                                                    v57 = ("\n\t" + (v57 + (("\n\t" + ("foo bar" + v57)) + "foo bar")))
                                                    _ = v77
                                                }
                                                var v78 uint32 = 2
                                                fmt.Println(v34, v62, fi40)
                                                _ = v78
                                            }
                                            _ = v62
                                        }
                                        var v79 int = int(v2)
                                        _ = v57
                                        _ = v58
                                        _ = v79
                                    }
                                    _ = len(v3)
                                    var v80 uint32 = 0
                                    _ = v80
                                }
                                if (false || (!v34)) {
                                    v45 += -60
                                    var v81 int8 = (-(v45))
                                    _ = v81
                                } else {
                                    if v34 {
                                        switch -2147483648 {
                                        case 32767:
                                            if v34 {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v34 = v34
                                                default:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                if v34 {
                                                    var v82 float64 = (1e10 * v36)
                                                    if (((-2.25 / 0.1) >= -2.25) || true) {
                                                        if (1.5 != 0.1) {
                                                            v82 += (1.5 / v36)
                                                            switch fi43 {
                                                            case 128:
                                                                v83 := GenS6{F0: v2}
                                                                _ = v83
                                                            default:
                                                                fmt.Println(fi40, fi43, v45)
                                                            }
                                                            v84 := GenS11{F0: -9223372036854775808, F1: false}
                                                            _ = v84
                                                        }
                                                        switch ("\n\t" + "foo bar") {
                                                        case "hello":
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            if (2 == 710) {
                                                                if len(v3) > 0 {
                                                                    v3[0] = 65535
                                                                }
                                                            } else {
                                                                var v85 int8 = v45
                                                                var v86 int8 = v45
                                                                if v34 {
                                                                    v86 -= -128
                                                                    var v87 uint16 = 167
                                                                    _ = v87
                                                                }
                                                                switch (fi40 << ((42) & 63)) {
                                                                case 42:
                                                                    if v34 {
                                                                        var v88 int = fi40
                                                                        if v34 {
                                                                            v34 = false
                                                                            v85 = v85
                                                                            var v89 uint64 = v2
                                                                            _ = v89
                                                                        }
                                                                        switch fi43 {
                                                                        case -872:
                                                                            if v34 {
                                                                                v86 = 17
                                                                                var v90 uint = 0
                                                                                _ = v90
                                                                            } else {
                                                                                var v91 int32 = 2147483647
                                                                                _ = v91
                                                                            }
                                                                        }
                                                                        v36 += v82
                                                                        if len(v4) > 0 {
                                                                            v4[0] = "hello"
                                                                        }
                                                                        _ = v88
                                                                    }
                                                                }
                                                                v86 = (v45 + -128)
                                                                _ = v85
                                                                _ = v86
                                                            }
                                                        case "foo bar":
                                                            v3 = append(v3, 2)
                                                        case "a":
                                                            v34 = true
                                                        }
                                                        if (v2 != v2) {
                                                            v92 := GenS12{F0: (65535 <= 772)}
                                                            if false {
                                                                if (!true) {
                                                                    v2 = v2
                                                                } else {
                                                                    var v93 string = "\n\t"
                                                                    var v94 int = (^(fi40))
                                                                    if (326 < 4294967295) {
                                                                        fmt.Println(v94, v2)
                                                                        switch (fi40 - -9223372036854775808) {
                                                                        case 1:
                                                                            if (!v34) {
                                                                                if (18446744073709551615 > 18446744073709551615) {
                                                                                    switch fi40 {
                                                                                    case 953:
                                                                                        v82 = v82
                                                                                    }
                                                                                }
                                                                                var v95 uint32 = 4294967295
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = (v93 + (v93 + ""))
                                                                                }
                                                                                _ = v95
                                                                            }
                                                                        case 0:
                                                                            _ = len(v3)
                                                                        case 128:
                                                                            v96 := GenS13{F0: 42}
                                                                            _ = v96
                                                                        }
                                                                    } else {
                                                                        v97 := GenS5{F0: v2, F1: 2, F2: -1}
                                                                        _ = v97
                                                                    }
                                                                    var v98 uint64 = (v2 - (v2 * (v2 + 4294967295)))
                                                                    if (!(!(-1 > -32768))) {
                                                                        if true {
                                                                            if (18446744073709551615 == 18446744073709551615) {
                                                                                var v99 uint32 = 525
                                                                                var v100 uint64 = 0
                                                                                _ = v100
                                                                                _ = v99
                                                                            } else {
                                                                                var v101 bool = true
                                                                                var v102 int16 = -32768
                                                                                v103 := GenS9{F0: 0, F1: 42}
                                                                                _ = v101
                                                                                _ = v102
                                                                                _ = v103
                                                                            }
                                                                            if v34 {
                                                                                v104 := GenS7{F0: 883, F1: 4294967295}
                                                                                var v105 float64 = (1e10 / v82)
                                                                                var v106 int16 = 1
                                                                                _ = v104
                                                                                _ = v105
                                                                                _ = v106
                                                                            }
                                                                            if (3.14159 == ((0.0 * 1e10) - 1e10)) {
                                                                                if (!true) {
                                                                                    var v107 bool = (v36 >= 1e10)
                                                                                    _ = v107
                                                                                } else {
                                                                                    var v108 uint32 = 143
                                                                                    if v34 {
                                                                                        if v34 {
                                                                                            if v34 {
                                                                                                fmt.Println(fi43, v98)
                                                                                                v3 = append(v3, v108)
                                                                                                switch "foo bar" {
                                                                                                case "":
                                                                                                    var v109 int = -1
                                                                                                    _ = v109
                                                                                                case "hello":
                                                                                                    var v110 int8 = v45
                                                                                                    _ = v110
                                                                                                case "a":
                                                                                                    if ((!(226 == v2)) && false) {
                                                                                                        fmt.Println(v94)
                                                                                                    } else {
                                                                                                        v36 = v36
                                                                                                        v108 = (^(v108))
                                                                                                        v111 := GenS3{F0: -2147483648, F1: 999}
                                                                                                        if false {
                                                                                                            v112 := GenS2{F0: uint(v2)}
                                                                                                            if ((1 == -136) || v34) {
                                                                                                                if len(v3) > 0 {
                                                                                                                    v3[0] = 65535
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v112
                                                                                                        }
                                                                                                        v113 := GenS14{F0: -340, F1: v93, F2: -569}
                                                                                                        v34 = (!v34)
                                                                                                        _ = v111
                                                                                                        _ = v113
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            if v34 {
                                                                                                var v114 int64 = (v1 >> ((65) & 63))
                                                                                                v1 -= v114
                                                                                                var v115 int = v94
                                                                                                var v116 int32 = 1
                                                                                                _ = v114
                                                                                                _ = v115
                                                                                                _ = v116
                                                                                            }
                                                                                            fmt.Println(v94, v82)
                                                                                            var v117 uint8 = 128
                                                                                            if len(v4) >= 2 {
                                                                                                v4 = v4[1:len(v4)]
                                                                                            }
                                                                                            v118 := GenS4{F0: (^(v98)), F1: v2, F2: (-(fi40))}
                                                                                            _ = v117
                                                                                            _ = v118
                                                                                        }
                                                                                    } else {
                                                                                        switch fi40 {
                                                                                        case 9223372036854775807:
                                                                                            var v119 int8 = 0
                                                                                            _ = v119
                                                                                        case 32767:
                                                                                            if v34 {
                                                                                                v94 += fi43
                                                                                                if len(v4) > 0 {
                                                                                                    v4[0] = v93
                                                                                                }
                                                                                                v3 = append(v3, v108)
                                                                                            }
                                                                                        default:
                                                                                            v82 = (-2.25 * (v36 * (v36 * float64(0.1))))
                                                                                        }
                                                                                        v45 = -1
                                                                                        v120 := GenS14{F0: 42, F1: "a", F2: 0}
                                                                                        v1 = int64(v108)
                                                                                        var v121 float64 = (v36 * v82)
                                                                                        _ = v120
                                                                                        _ = v121
                                                                                    }
                                                                                    _ = v108
                                                                                }
                                                                                if v34 {
                                                                                    var v122 float64 = 1e10
                                                                                    _ = v122
                                                                                }
                                                                            }
                                                                            if ((73 > 42) && (v34 || (!true))) {
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v94)
                                                                                var v123 uint64 = v2
                                                                                var v124 int = 2147483647
                                                                                _ = v123
                                                                                _ = v124
                                                                            }
                                                                            v1 = (v1 << ((0) & 63))
                                                                            var v125 uint8 = 2
                                                                            _ = v125
                                                                        }
                                                                        fmt.Println(v1)
                                                                        if len(v4) > 0 {
                                                                            v4[0] = v93
                                                                        }
                                                                    }
                                                                    fmt.Println(v98, v98)
                                                                    _ = v93
                                                                    _ = v94
                                                                    _ = v98
                                                                }
                                                            } else {
                                                                if v34 {
                                                                    if true {
                                                                        v34 = (v34 || ((v45 ^ (v45 >> ((1) & 63))) != v45))
                                                                        switch (fi43 & 2147483647) {
                                                                        case 2147483647:
                                                                            _ = len(v4)
                                                                        case 707:
                                                                            var v126 bool = v34
                                                                            _ = v126
                                                                        case 0:
                                                                            var v127 float32 = float32(v36)
                                                                            _ = v127
                                                                        }
                                                                        fmt.Println(v36, v2, v1)
                                                                    }
                                                                } else {
                                                                    var v128 int8 = 84
                                                                    var v129 uint32 = 797
                                                                    _ = v128
                                                                    _ = v129
                                                                }
                                                            }
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v130 := GenS13{F0: 775}
                                                            if false {
                                                                var v131 uint64 = (v2 ^ v2)
                                                                v2 -= 18446744073709551615
                                                                _ = v131
                                                            } else {
                                                                if (761 > 427) {
                                                                    fmt.Println(v2)
                                                                    v132 := GenS9{F0: 396, F1: 255}
                                                                    _ = v132
                                                                }
                                                            }
                                                            v1 -= v1
                                                            if v34 {
                                                                v45 -= (v45 >> ((883) & 63))
                                                            }
                                                            var v133 bool = v34
                                                            var v134 uint = 247
                                                            var v135 uint8 = 206
                                                            _ = v130
                                                            _ = v133
                                                            _ = v134
                                                            _ = v135
                                                            _ = v92
                                                        } else {
                                                            v34 = v34
                                                        }
                                                    }
                                                    _ = v82
                                                }
                                                var v136 float32 = (1e10 - 0.1)
                                                _ = v136
                                            } else {
                                                v2 += v2
                                                if false {
                                                    v45 -= (v45 % ((v45) | 1))
                                                    if true {
                                                        v1 = v1
                                                        if true {
                                                            var v137 int16 = 32767
                                                            _ = v137
                                                        } else {
                                                            v138 := GenS15{F0: 274, F1: 840, F2: "hello"}
                                                            _ = v138
                                                        }
                                                    } else {
                                                        var v139 int = (fi43 - fi43)
                                                        v1 += int64(v2)
                                                        v1 = v1
                                                        var v140 string = ("a" + ("\n\t" + ("" + (("\n\t" + "") + "hello"))))
                                                        _ = v139
                                                        _ = v140
                                                    }
                                                    v1 = 9223372036854775807
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = ("\n\t" + "a")
                                                }
                                            }
                                        default:
                                            v141 := GenS16{F0: v34, F1: 255, F2: (^(fi43))}
                                            _ = v141
                                        }
                                    } else {
                                        var v142 int = 42
                                        _ = v142
                                    }
                                }
                            }
                            v143 := GenS4{F0: 1, F1: (v2 + v2), F2: fi43}
                            var v144 int32 = 2147483647
                            _ = fi43
                            _ = v143
                            _ = v144
                            _ = v45
                        }
                    }
                case 127:
                    v145 := GenS17{F0: 65535}
                    _ = v145
                default:
                    if len(v4) > 0 {
                        v4[0] = "hello"
                    }
                }
                for fi146 := 0; fi146 < 16; fi146++ {
                    var v147 uint16 = 65535
                    v34 = (!(42 < 127))
                    v147 += v147
                    _ = fi146
                    _ = v147
                }
                _ = fi40
            }
        }
    }
    for fi148 := 0; fi148 < 11; fi148++ {
        var v149 int64 = (-(v1))
        if ((^(fi148)) == fi148) {
            v149 = v149
            v1 += v149
        } else {
            var v150 bool = ((fi148 + fi148) != (fi148 + (^(fi148))))
            for fi151 := 0; fi151 < 19; fi151++ {
                v150 = true
                _ = fi151
            }
            _ = v150
        }
        v149 = (v1 * 912)
        _ = fi148
        _ = v149
    }
    switch 894 {
    case 1:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 42:
        for fi152 := 0; fi152 < 15; fi152++ {
            var v153 string = "a"
            _ = fi152
            _ = v153
        }
    case 831:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v154 float32 = 1e10
        _ = v154
    }
    v155 := GenS2{F0: 689}
    for fi156 := 0; fi156 < 14; fi156++ {
        var v157 uint16 = 65535
        var v158 float32 = 0.1
        if (v34 || (!(((("hello" + "a") + "a") + ("a" + "hello")) != "foo bar"))) {
            switch fi156 {
            case 4294967295:
                if (275 <= 65535) {
                    for fi159 := 0; fi159 < 18; fi159++ {
                        switch "\n\t" {
                        case "a":
                            fmt.Println(v2, v2)
                        case "hello":
                            fmt.Println(v34)
                        default:
                            if v34 {
                                var v160 string = ""
                                var v161 string = v160
                                v162 := GenS18{F0: (false || ((v161 + v160) > "hello"))}
                                var v163 bool = v34
                                _ = v160
                                _ = v161
                                _ = v162
                                _ = v163
                            }
                        }
                        if true {
                            var v164 uint64 = 923
                            _ = v164
                        }
                        var v165 uint32 = 42
                        _ = fi159
                        _ = v165
                    }
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(v34, v34, fi156)
                }
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            v34 = true
        }
        _ = fi156
        _ = v157
        _ = v158
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v155
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v4
}
