package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 uint; F2 int32 }
type GenS6 struct { F0 int; F1 int8; F2 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS10 struct { F0 uint8; F1 uint64 }
type GenS11 struct { F0 string; F1 uint; F2 uint }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int16; F1 uint8; F2 int64 }
type GenS15 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS16 struct { F0 uint8; F1 string }
type GenS17 struct { F0 int16; F1 int8 }
type GenS18 struct { F0 int64; F1 uint64 }
type GenS19 struct { F0 uint16; F1 int32 }
type GenS20 struct { F0 string; F1 bool }

func main() {
    var v1 float64 = (0.1 + 0.0)
    var v2 uint8 = 42
    var v3 int32 = -2147483648
    v4 := []string{}
    v2 = (v2 ^ v2)
    v3 = 2147483647
    v5 := GenS1{F0: 167}
    switch 9223372036854775807 {
    case -275:
        for fi6 := 0; fi6 < 10; fi6++ {
            v2 = 0
            for fi7 := 0; fi7 < 9; fi7++ {
                v3 += v3
                _ = fi7
            }
            _ = fi6
        }
    case 127:
        var v8 bool = false
        _ = v8
    default:
        if ((v2 * v2) < v2) {
            v1 -= v1
            if true {
                v3 += v3
                var v9 uint64 = 504
                v10 := GenS2{F0: 9223372036854775807, F1: 312, F2: v2}
                _ = v10
                _ = v9
            } else {
                v11 := GenS2{F0: -246, F1: 42, F2: v2}
                v1 = v1
                _ = len(v4)
                _ = v11
            }
        } else {
            if (!false) {
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
                v12 := GenS3{F0: 128, F1: 418}
                v1 += (0.0 / float64(1e10))
                _ = v12
            } else {
                for fi13 := 0; fi13 < 18; fi13++ {
                    fmt.Println(fi13)
                    for fi14 := 0; fi14 < 12; fi14++ {
                        if (1 <= 354) {
                            if (false || false) {
                                v15 := GenS4{F0: (v2 | v2)}
                                if (v2 != 2) {
                                    switch fi13 {
                                    case 9223372036854775807:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case 1:
                                        v4 = append(v4, ("hello" + "foo bar"))
                                    case 127:
                                        v3 = v3
                                    default:
                                        v16 := GenS1{F0: 592}
                                        v17 := GenS3{F0: v2, F1: uint64(v2)}
                                        var v18 int8 = 100
                                        _ = v16
                                        _ = v17
                                        _ = v18
                                    }
                                    v2 = v2
                                }
                                v19 := GenS1{F0: 263}
                                _ = v15
                                _ = v19
                            } else {
                                var v20 bool = ((!(0.1 <= v1)) || (3.14159 < 0.1))
                                _ = v20
                            }
                        }
                        switch (fi14 ^ 191) {
                        case 374:
                            var v21 int64 = -1
                            _ = v21
                        case 32767:
                            v2 += v2
                        case -949:
                            var v22 uint32 = 1
                            _ = v22
                        default:
                            fmt.Println(fi13, fi14, v2)
                        }
                        _ = fi14
                    }
                    v23 := GenS5{F0: v3, F1: 105, F2: 2147483647}
                    _ = fi13
                    _ = v23
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            for fi24 := 0; fi24 < 18; fi24++ {
                for fi25 := 0; fi25 < 8; fi25++ {
                    v26 := GenS1{F0: 1}
                    var v27 uint8 = (^(v2))
                    v28 := GenS4{F0: v2}
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = fi24
            }
        }
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        var v30 uint64 = 80
        for fi31 := 0; fi31 < 1; fi31++ {
            fmt.Println(v30)
            var v32 int16 = -979
            fmt.Println(v30, fi29, v2)
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 4; fi33++ {
            fmt.Println(v1, v3, v30)
            v2 = 128
            var v34 bool = (2 >= (v2 % ((v2) | 1)))
            _ = fi33
            _ = v34
        }
        _ = fi29
        _ = v30
    }
    var v35 float64 = v1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v36 := GenS1{F0: 84}
    if (3.14159 < 1e10) {
        v37 := GenS5{F0: 2147483647, F1: 349, F2: -2147483648}
        _ = v37
    }
    switch "foo bar" {
    case "foo bar":
        v38 := GenS2{F0: 4294967295, F1: 1, F2: 128}
        _ = v38
    default:
        if (469 > 65535) {
            if (18446744073709551615 > 65535) {
                switch -476 {
                case -150:
                    if len(v4) > 0 {
                        v4[0] = "a"
                    }
                }
                if true {
                    _ = len(v4)
                    if (42 != -51) {
                        if false {
                            var v39 int64 = 679
                            v40 := GenS4{F0: v2}
                            for fi41 := 0; fi41 < 1; fi41++ {
                                fmt.Println(v3, v35)
                                v3 = (v3 << ((2) & 63))
                                var v42 bool = (-439 != v3)
                                _ = fi41
                                _ = v42
                            }
                            _ = v39
                            _ = v40
                        }
                    }
                }
            }
            _ = len(v4)
            v1 = v35
        }
    }
    fmt.Println(v35)
    for fi43 := 0; fi43 < 4; fi43++ {
        var v44 int = fi43
        switch (^(v44)) {
        case 32768:
            for fi45 := 0; fi45 < 8; fi45++ {
                if ((true || true) || true) {
                    v46 := GenS3{F0: v2, F1: 446}
                    v47 := GenS3{F0: 255, F1: 2}
                    v48 := GenS3{F0: 170, F1: 2}
                    _ = v46
                    _ = v47
                    _ = v48
                } else {
                    if false {
                        var v49 uint16 = 0
                        _ = v49
                    }
                }
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    switch "hello" {
                    case "":
                        v2 += v2
                    case "\n\t":
                        v3 -= (v3 * v3)
                    default:
                        var v50 int64 = -1
                        _ = v50
                    }
                case "hello":
                    var v51 float64 = 3.14159
                    _ = v51
                default:
                    if true {
                        if (false || (!((!false) || (0 <= 9223372036854775807)))) {
                            v4 = append(v4, ("foo bar" + ("hello" + "")))
                            v1 += (float64(1e10) * v1)
                        }
                        v35 -= v1
                        switch "foo bar" {
                        case "\n\t":
                            var v52 bool = ("\n\t" >= "\n\t")
                            _ = v52
                        }
                        if (true && true) {
                            if (!true) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                fmt.Println(v35, v35, fi43)
                            } else {
                                switch "hello" {
                                case "":
                                    v35 = (float64(1e10) - (((float64(float32(v1)) * 3.14159) + float64((((1.5 / float32(v35)) + -2.25) / (3.14159 / ((1.5 / 3.14159) + float32((float64(0.0) * (v35 / float64((float32(((((v35 - v1) - (-2.25 * v35)) / v1) / (float64(1e10) - 0.1))) * 1e10)))))))))) / (v1 - (((float64(1e10) + float64((3.14159 - (-2.25 / 0.1)))) - 0.1) * ((v35 / (float64(1e10) / float64(3.14159))) * v35)))))
                                case "a":
                                    v2 -= (^(v2))
                                case "hello":
                                    switch 127 {
                                    case 4294967295:
                                        if len(v4) > 0 {
                                            v4[0] = (("\n\t" + ("hello" + (((("a" + ("a" + "hello")) + "") + "hello") + ("foo bar" + "foo bar")))) + (("" + "") + "a"))
                                        }
                                    }
                                default:
                                    var v53 float64 = v1
                                    _ = v53
                                }
                                v54 := GenS1{F0: 0}
                                _ = v54
                            }
                            fmt.Println(v2)
                            if (uint16(v2) <= 412) {
                                v35 = (0.1 * v35)
                                if (false && (!(!(469 >= 294)))) {
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(fi45, fi45)
                                    case "":
                                        v3 = (v3 | int32(v2))
                                    default:
                                        v55 := GenS6{F0: 32767, F1: 1, F2: 0}
                                        _ = v55
                                    }
                                    v56 := GenS4{F0: (v2 - v2)}
                                    fmt.Println(v35, v3, v3)
                                    _ = v56
                                } else {
                                    if false {
                                        if ((("" + "a") + ("a" + ("\n\t" + "a"))) >= "foo bar") {
                                            v4 = append(v4, (("foo bar" + ("\n\t" + "hello")) + ""))
                                        } else {
                                            switch v44 {
                                            case 32768:
                                                if true {
                                                    v57 := GenS6{F0: fi45, F1: 100, F2: 4294967295}
                                                    var v58 uint = 18446744073709551615
                                                    _ = v57
                                                    _ = v58
                                                } else {
                                                    v59 := GenS7{F0: v2}
                                                    switch (fi43 * (v44 + fi45)) {
                                                    case 551:
                                                        v60 := GenS8{F0: 132}
                                                        _ = v60
                                                    case 2147483647:
                                                        var v61 float32 = 0.0
                                                        if ("" != "a") {
                                                            switch "a" {
                                                            case "hello":
                                                                fmt.Println(fi43, fi45, v44)
                                                            case "a":
                                                                var v62 uint = 602
                                                                _ = v62
                                                            case "\n\t":
                                                                if true {
                                                                    var v63 int16 = 42
                                                                    _ = v63
                                                                } else {
                                                                    var v64 bool = true
                                                                    if v64 {
                                                                        fmt.Println(v64, fi43)
                                                                    } else {
                                                                        v44 += 938
                                                                    }
                                                                    _ = v64
                                                                }
                                                            }
                                                        }
                                                        if len(v4) > 0 {
                                                            v4[0] = "\n\t"
                                                        }
                                                        if false {
                                                            var v65 int64 = 80
                                                            if (-128 < 127) {
                                                                var v66 uint16 = 1
                                                                _ = v66
                                                            } else {
                                                                if false {
                                                                    v67 := GenS4{F0: v2}
                                                                    if (false && ((v35 + (float64((1e10 * 1e10)) + (v1 + float64(v61)))) >= v35)) {
                                                                        fmt.Println(v44, v3, v2)
                                                                        v68 := GenS6{F0: (fi45 * fi45), F1: -36, F2: 842}
                                                                        _ = v68
                                                                    }
                                                                    _ = v67
                                                                }
                                                            }
                                                            _ = v65
                                                        } else {
                                                            v61 = v61
                                                        }
                                                        var v69 float32 = -2.25
                                                        _ = v61
                                                        _ = v69
                                                    }
                                                    _ = v59
                                                }
                                            case 128:
                                                switch (fi45 ^ v44) {
                                                case 9223372036854775807:
                                                    v44 = fi45
                                                case 32767:
                                                    if (-1 >= 952) {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        fmt.Println(v1)
                                                    } else {
                                                        v4 = append(v4, ((("a" + (("" + "\n\t") + "\n\t")) + "") + "foo bar"))
                                                        var v70 int64 = -184
                                                        var v71 uint = 255
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                case 42:
                                                    v44 += 4294967295
                                                default:
                                                    v44 -= (fi45 >> ((1) & 63))
                                                }
                                            }
                                            v4 = append(v4, "")
                                        }
                                        var v72 uint32 = 0
                                        var v73 float32 = 0.1
                                        switch (fi45 ^ fi45) {
                                        case 128:
                                            var v74 bool = true
                                            _ = v74
                                        }
                                        _ = v72
                                        _ = v73
                                    }
                                    if (v3 > v3) {
                                        v35 += (float64(0.1) + (float64(((0.1 * (0.1 * -2.25)) * -2.25)) / v35))
                                        if true {
                                            fmt.Println(fi45, v35, v1)
                                            v75 := GenS8{F0: 859}
                                            _ = v75
                                        }
                                    } else {
                                        v44 += 42
                                        fmt.Println(fi43, v3, v1)
                                    }
                                }
                            }
                        }
                    } else {
                        var v76 int32 = (v3 % ((-478) | 1))
                        v77 := GenS8{F0: 65535}
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        _ = v76
                        _ = v77
                    }
                    v78 := GenS9{F0: 128, F1: 42, F2: 2}
                    if (1 > 357) {
                        var v79 int32 = v3
                        _ = v79
                    }
                    _ = v78
                }
                _ = fi45
            }
        }
        _ = fi43
        _ = v44
    }
    v80 := GenS10{F0: v2, F1: 42}
    for fi81 := 0; fi81 < 9; fi81++ {
        switch "hello" {
        case "\n\t":
            v2 = (v2 - 81)
        case "hello":
            v82 := GenS6{F0: (fi81 / ((fi81) | 1)), F1: 127, F2: 12}
            _ = v82
        case "foo bar":
            switch "hello" {
            case "":
                v83 := GenS11{F0: "\n\t", F1: 18446744073709551615, F2: 144}
                _ = v83
            case "hello":
                var v84 float32 = (0.1 + (1e10 + 3.14159))
                _ = v84
            case "\n\t":
                var v85 uint64 = 4294967295
                _ = v85
            }
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        v1 = v35
        var v86 int8 = 0
        _ = fi81
        _ = v86
    }
    fmt.Println(v3, v2, v2)
    switch "hello" {
    case "\n\t":
        var v87 int8 = 118
        _ = v87
    case "foo bar":
        if (v35 == v1) {
            v88 := GenS1{F0: 4294967295}
            if (false && true) {
                var v89 int32 = (v3 << ((275) & 63))
                _ = v89
            } else {
                fmt.Println(v35, v1)
                switch 32768 {
                case -309:
                    v90 := GenS12{F0: -128}
                    _ = v90
                case 426:
                    switch 2147483647 {
                    case 127:
                        var v91 uint8 = v2
                        _ = v91
                    case 32767:
                        v92 := GenS5{F0: 786, F1: 69, F2: v3}
                        _ = v92
                    case 538:
                        switch -330 {
                        case 9223372036854775807:
                            var v93 uint8 = (v2 ^ (v2 | v2))
                            _ = v93
                        case 388:
                            var v94 int = 900
                            _ = v94
                        case 2147483647:
                            if len(v4) > 0 {
                                v4[0] = ("foo bar" + "hello")
                            }
                        default:
                            if false {
                                v2 = v2
                            }
                        }
                    }
                default:
                    var v95 string = "hello"
                    _ = v95
                }
                var v96 int8 = 100
                _ = v96
            }
            _ = v88
        } else {
            for fi97 := 0; fi97 < 13; fi97++ {
                v98 := GenS13{F0: 4294967295}
                if (true || true) {
                    var v99 int16 = 32767
                    _ = v99
                }
                v2 += v2
                _ = fi97
                _ = v98
            }
            for fi100 := 0; fi100 < 1; fi100++ {
                v101 := GenS14{F0: 1000, F1: 1, F2: 1}
                for fi102 := 0; fi102 < 10; fi102++ {
                    var v103 int64 = 1
                    v3 += (v3 & 2147483647)
                    v35 = 1.5
                    _ = fi102
                    _ = v103
                }
                _ = fi100
                _ = v101
            }
            if (true || (("foo bar" < "hello") && true)) {
                var v104 uint = 4294967295
                for fi105 := 0; fi105 < 20; fi105++ {
                    if false {
                        var v106 uint8 = 127
                        var v107 uint8 = (v2 - v2)
                        v108 := GenS15{F0: 425, F1: 1, F2: v2}
                        _ = v106
                        _ = v107
                        _ = v108
                    } else {
                        if false {
                            if true {
                                v109 := GenS16{F0: v2, F1: "\n\t"}
                                v1 = 1e10
                                var v110 uint = v104
                                _ = v109
                                _ = v110
                            }
                            v2 += v2
                            if (true || true) {
                                if false {
                                    if (!true) {
                                        v2 = v2
                                        var v111 float32 = 0.0
                                        for fi112 := 0; fi112 < 14; fi112++ {
                                            v113 := GenS17{F0: int16(fi112), F1: -1}
                                            if len(v4) > 0 {
                                                v4[0] = ""
                                            }
                                            fmt.Println(v1, fi105, fi112)
                                            _ = fi112
                                            _ = v113
                                        }
                                        _ = v111
                                    }
                                    v114 := GenS18{F0: 28, F1: 150}
                                    if (!true) {
                                        var v115 float32 = 3.14159
                                        _ = v115
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = ""
                                        }
                                        var v116 uint16 = 32767
                                        _ = v116
                                    }
                                    _ = v114
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        switch (fi105 % ((int(v3)) | 1)) {
                                        case 32767:
                                            var v117 float32 = 3.14159
                                            _ = v117
                                        case 42:
                                            var v118 uint = 698
                                            _ = v118
                                        case 0:
                                            fmt.Println(v104, v35)
                                        }
                                    default:
                                        for fi119 := 0; fi119 < 16; fi119++ {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v1 -= v1
                                            _ = fi119
                                        }
                                    }
                                    switch ("\n\t" + "a") {
                                    case "hello":
                                        for fi120 := 0; fi120 < 1; fi120++ {
                                            v121 := GenS19{F0: 1, F1: 42}
                                            if (!(!false)) {
                                                var v122 uint8 = v2
                                                _ = v122
                                            } else {
                                                switch 127 {
                                                case 0:
                                                    if len(v4) > 0 {
                                                        v4[0] = "a"
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = "\n\t"
                                                    }
                                                }
                                                var v123 float64 = (v35 - v1)
                                                v124 := GenS8{F0: uint32(v104)}
                                                var v125 int = -1
                                                v125 = fi120
                                                v123 += (float64(3.14159) - v123)
                                                v126 := GenS20{F0: "foo bar", F1: ("foo bar" < "")}
                                                _ = v123
                                                _ = v124
                                                _ = v125
                                                _ = v126
                                            }
                                            _ = fi120
                                            _ = v121
                                        }
                                    case "\n\t":
                                        v2 -= (v2 | v2)
                                    default:
                                        v2 = (^(v2))
                                    }
                                }
                                v127 := GenS8{F0: 2147483647}
                                v3 -= v3
                                _ = v127
                            }
                        } else {
                            var v128 uint = (^(v104))
                            _ = len(v4)
                            _ = v128
                        }
                        for fi129 := 0; fi129 < 2; fi129++ {
                            v4 = append(v4, "foo bar")
                            var v130 bool = ((v104 ^ 0) != v104)
                            var v131 uint = 18446744073709551615
                            _ = fi129
                            _ = v130
                            _ = v131
                        }
                        v132 := GenS16{F0: v2, F1: ""}
                        _ = v132
                    }
                    _ = fi105
                }
                _ = v104
            }
        }
    }
    v133 := GenS17{F0: 1, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v133
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v80
}
