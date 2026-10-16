package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint64; F1 bool; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS9 struct { F0 uint64; F1 uint8 }
type GenS10 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS11 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS12 struct { F0 int16; F1 uint32; F2 int }
type GenS13 struct { F0 uint; F1 int32 }
type GenS14 struct { F0 uint64; F1 int16; F2 bool }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int16; F1 uint8; F2 int32 }
type GenS17 struct { F0 int16; F1 int8; F2 uint8 }
func genFn0(p0 int, p1 uint16, p2 bool) bool {
    var v7 uint = 0
    v8 := GenS1{F0: 255}
    var v9 uint32 = 159
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := []int{(^(p0)), (fi10 - (^(fi10))), 128}
        _ = fi10
        _ = v11
    }
    v12 := []int{p0, (p0 << ((42) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    _ = v8
    _ = v9
    return p2
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{518, 2147483648, 2147483648, 934}
    var v3 uint8 = 127
    var v4 uint32 = 2147483648
    v5 := []uint32{573, (^(v4)), (v4 >> ((0) & 63)), 4294967295}
    var v6 uint64 = 18446744073709551615
    switch 9223372036854775807 {
    case 413:
        switch 127 {
        case -1:
            var v13 int = -9223372036854775808
            _ = v13
        case 42:
            v1[v4] = uint32(v6)
        }
    case 972:
        v4 += v4
    case 694:
        var v14 string = "\n\t"
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS1{F0: 4294967295}
        var v17 int = fi15
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println(v4, v6)
    if ("\n\t" >= "a") {
        if false {
            for fi18 := 0; fi18 < 15; fi18++ {
                v4 = uint32(v3)
                v19 := GenS1{F0: 65535}
                var v20 uint64 = (v6 | v6)
                _ = fi18
                _ = v19
                _ = v20
            }
            switch "a" {
            case "":
                var v21 int8 = 100
                _ = v21
            case "\n\t":
                v22 := GenS1{F0: 0}
                _ = v22
            case "a":
                v3 += v3
            default:
                switch "hello" {
                case "a":
                    v3 = v3
                case "\n\t":
                    var v23 uint16 = 91
                    _ = v23
                }
            }
            if true {
                switch "a" {
                case "a":
                    fmt.Println(v6, v6)
                case "\n\t":
                    switch "foo bar" {
                    case "hello":
                        v24 := GenS1{F0: 850}
                        _ = v24
                    case "a":
                        v2 = append(v2, v4)
                    case "foo bar":
                        v25 := GenS2{F0: 114, F1: v6, F2: 854}
                        _ = v25
                    }
                case "":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    if true {
                        var v26 int = -559
                        var v27 int64 = 42
                        var v28 int8 = 18
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                }
            } else {
                var v29 int64 = -92
                var v30 bool = (false && true)
                v30 = (65535 > 559)
                _ = v29
                _ = v30
            }
        } else {
            for fi31 := 0; fi31 < 7; fi31++ {
                switch "foo bar" {
                case "":
                    var v32 int16 = 1000
                    _ = v32
                default:
                    switch "hello" {
                    case "\n\t":
                        v33 := v1[v4]
                        _ = v33
                    case "":
                        var v34 uint64 = uint64(v3)
                        _ = v34
                    default:
                        v4 += (v4 / (((v4 / ((v4) | 1))) | 1))
                    }
                }
                var v35 int = (fi31 & (-(fi31)))
                var v36 uint64 = (v6 | (v6 >> ((18446744073709551615) & 63)))
                _ = fi31
                _ = v35
                _ = v36
            }
        }
        var v37 string = "a"
        _ = v37
    } else {
        var v38 uint = 762
        v39 := GenS3{F0: (^(v6)), F1: (!(!("hello" == "hello"))), F2: "\n\t"}
        _ = v38
        _ = v39
    }
    if len(v2) > 0 {
        v2[0] = v4
    }
    v6 = (v6 & 262)
    _ = len(v2)
    for fi40 := 0; fi40 < 16; fi40++ {
        if false {
            fmt.Println(v6, v6)
        } else {
            v41 := GenS1{F0: 42}
            var v42 int32 = -1
            _ = v41
            _ = v42
        }
        if (false || (((1 >= 0) && (944 >= 2)) && true)) {
            v6 += v6
            if (255 < 2) {
                v1[(v4 + v4)] = 0
            }
            var v43 uint16 = 42
            _ = v43
        }
        _ = fi40
    }
    v44 := GenS3{F0: 249, F1: (-16 >= 227), F2: "foo bar"}
    fmt.Println(v4, v4, v3)
    if false {
        fmt.Println(v4, v4, v4)
    } else {
        v1[(^(v4))] = v4
        var v45 bool = false
        v46 := GenS2{F0: 32767, F1: 42, F2: 0}
        _ = v45
        _ = v46
    }
    fmt.Println(v6, v4)
    v47 := GenS2{F0: 42, F1: v6, F2: 42}
    if true {
        var v48 bool = false
        if true {
            for fi49 := 0; fi49 < 18; fi49++ {
                var v50 int16 = -676
                _ = fi49
                _ = v50
            }
            switch ("\n\t" + "") {
            case "foo bar":
                var v51 float64 = 1e10
                _ = v51
            case "hello":
                for fi52 := 0; fi52 < 14; fi52++ {
                    var v53 uint32 = v4
                    fmt.Println(v3, v3)
                    fmt.Println(v4, v53, v6)
                    _ = fi52
                    _ = v53
                }
            }
        } else {
            for fi54 := 0; fi54 < 1; fi54++ {
                switch "a" {
                case "foo bar":
                    fmt.Println(fi54, fi54)
                case "":
                    delete(v1, v4)
                case "hello":
                    v3 = (v3 >> ((940) & 63))
                }
                _ = fi54
            }
            v4 = v4
        }
        if ((!v48) || (1e10 > 0.0)) {
            switch "hello" {
            case "a":
                if v48 {
                    if (int(v3) > 127) {
                        _ = len(v1)
                        switch ((("a" + ("" + (("hello" + "") + "foo bar"))) + "hello") + "foo bar") {
                        case "\n\t":
                            for fi55 := 0; fi55 < 11; fi55++ {
                                for fi56 := 0; fi56 < 19; fi56++ {
                                    v48 = v48
                                    switch (("" + "\n\t") + "foo bar") {
                                    case "":
                                        if (644 != 18446744073709551615) {
                                            v4 -= v4
                                            fmt.Println(fi55, v48)
                                            v5 = append(v5, v4)
                                        } else {
                                            v4 = 42
                                            var v57 string = ("foo bar" + ("" + "hello"))
                                            _ = v57
                                        }
                                    case "\n\t":
                                        if false {
                                            if (!true) {
                                                _ = len(v5)
                                                fmt.Println(fi55, v4, v48)
                                                fmt.Println(v4, v6, fi56)
                                            } else {
                                                if true {
                                                    switch fi55 {
                                                    case 128:
                                                        if ((false && true) || true) {
                                                            if (551 >= 32767) {
                                                                if (127 < -128) {
                                                                    if (!(!v48)) {
                                                                        fmt.Println(v4, fi56)
                                                                        delete(v1, (v4 | v4))
                                                                        if ((false || v48) && (v48 && true)) {
                                                                            _ = len(v2)
                                                                            switch ("\n\t" + "foo bar") {
                                                                            case "a":
                                                                                var v58 uint32 = v4
                                                                                _ = v58
                                                                            case "":
                                                                                if (-1 <= 0) {
                                                                                    fmt.Println(fi56, v6, fi55)
                                                                                    v48 = v48
                                                                                    var v59 bool = ((v4 & v4) == (^(v4)))
                                                                                    v48 = (!v48)
                                                                                    _ = v59
                                                                                } else {
                                                                                    var v60 int = (fi56 | (fi56 * fi55))
                                                                                    _ = v60
                                                                                }
                                                                            case "\n\t":
                                                                                var v61 uint32 = v4
                                                                                _ = v61
                                                                            }
                                                                        }
                                                                        v4 = 56
                                                                        if v48 {
                                                                            fmt.Println(v48)
                                                                            v4 -= v4
                                                                        } else {
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                if (!((32767 != 739) || true)) {
                                                                                    v3 = v3
                                                                                    var v62 float64 = 3.14159
                                                                                    var v63 int64 = -1
                                                                                    _ = v62
                                                                                    _ = v63
                                                                                }
                                                                            default:
                                                                                v48 = v48
                                                                            }
                                                                            var v64 int8 = 42
                                                                            var v65 float32 = float32(-2.25)
                                                                            if (1.5 <= 0.1) {
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    var v66 int64 = -417
                                                                                    _ = v66
                                                                                case "foo bar":
                                                                                    var v67 int32 = 1
                                                                                    _ = v67
                                                                                default:
                                                                                    if false {
                                                                                        _ = len(v1)
                                                                                        v68 := GenS1{F0: 193}
                                                                                        v6 = (^(v6))
                                                                                        _ = v68
                                                                                    }
                                                                                }
                                                                                fmt.Println(v48, v64, v4)
                                                                                if (!(!true)) {
                                                                                    fmt.Println(fi56)
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    var v69 float32 = (v65 - (v65 / float32(1e10)))
                                                                                    _ = v69
                                                                                } else {
                                                                                    v70 := GenS4{F0: 2}
                                                                                    var v71 float32 = float32(-2.25)
                                                                                    _ = v70
                                                                                    _ = v71
                                                                                }
                                                                            }
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                    } else {
                                                                        switch 128 {
                                                                        case 42:
                                                                            v48 = false
                                                                        case -544:
                                                                            fmt.Println(v3, v3)
                                                                        case -387:
                                                                            if len(v2) > 0 {
                                                                                v2[0] = (v4 * (v4 * (v4 & (v4 - (v4 | 2147483648)))))
                                                                            }
                                                                        default:
                                                                            fmt.Println(fi55, v4, v3)
                                                                        }
                                                                        v6 += (v6 - 244)
                                                                    }
                                                                    v3 = (^(v3))
                                                                } else {
                                                                    var v72 uint64 = (v6 | (v6 >> ((18446744073709551615) & 63)))
                                                                    _ = v72
                                                                }
                                                            } else {
                                                                v73 := GenS2{F0: 255, F1: (v6 + (v6 | v6)), F2: -176}
                                                                _ = v73
                                                            }
                                                        }
                                                    case 0:
                                                        v1[v4] = (v4 - 922)
                                                    case 92:
                                                        v4 = uint32(v6)
                                                    }
                                                }
                                                fmt.Println(v6)
                                                v74 := GenS5{F0: (v4 & (v4 >> ((18446744073709551615) & 63))), F1: (^(v3))}
                                                _ = v74
                                            }
                                            v75 := GenS4{F0: 0}
                                            var v76 float64 = 1e10
                                            _ = v75
                                            _ = v76
                                        } else {
                                            var v77 int = (fi55 >> ((18446744073709551615) & 63))
                                            if (false || true) {
                                                v4 = v4
                                                v78 := GenS3{F0: (v6 ^ (v6 - v6)), F1: v48, F2: "\n\t"}
                                                if (!(3.14159 >= 1.5)) {
                                                    _ = len(v1)
                                                    if len(v2) > 0 {
                                                        v2[0] = v4
                                                    }
                                                    switch -9223372036854775808 {
                                                    case -2147483648:
                                                        switch ("\n\t" + "") {
                                                        case "":
                                                            if false {
                                                                v4 += v4
                                                                var v79 uint16 = 263
                                                                switch (fi55 & v77) {
                                                                case -1:
                                                                    v80 := GenS6{F0: -113}
                                                                    _ = v80
                                                                case 398:
                                                                    fmt.Println(v6, v48, v6)
                                                                case 127:
                                                                    if (!(1e10 != 0.1)) {
                                                                        var v81 uint16 = (v79 * (v79 / (((^(v79))) | 1)))
                                                                        var v82 int8 = 36
                                                                        switch (-(fi56)) {
                                                                        case 0:
                                                                            v77 -= 448
                                                                        case 32767:
                                                                            v83 := GenS5{F0: v4, F1: uint8(v6)}
                                                                            _ = v83
                                                                        case 913:
                                                                            v84 := GenS7{F0: ("a" + ("" + "foo bar"))}
                                                                            _ = v84
                                                                        default:
                                                                            if true {
                                                                                switch (("foo bar" + "") + (("\n\t" + ("\n\t" + "")) + "\n\t")) {
                                                                                case "":
                                                                                    v1[v4] = v4
                                                                                case "\n\t":
                                                                                    var v85 float32 = (-2.25 / 3.14159)
                                                                                    _ = v85
                                                                                case "a":
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        if v48 {
                                                                                            fmt.Println(v82)
                                                                                        } else {
                                                                                            fmt.Println(v77, fi55)
                                                                                            var v86 int16 = 0
                                                                                            _ = v86
                                                                                        }
                                                                                    }
                                                                                    switch v77 {
                                                                                    case 1:
                                                                                        var v87 uint = 344
                                                                                        _ = v87
                                                                                    case 213:
                                                                                        if v48 {
                                                                                            v6 = v6
                                                                                        }
                                                                                    default:
                                                                                        v4 = v4
                                                                                    }
                                                                                default:
                                                                                    if v48 {
                                                                                        if v48 {
                                                                                            _ = len(v5)
                                                                                        }
                                                                                        var v88 uint = 0
                                                                                        v88 = 18446744073709551615
                                                                                        _ = v88
                                                                                    } else {
                                                                                        v89 := GenS6{F0: 813}
                                                                                        v90 := GenS8{F0: ((v77 << ((uint(v4)) & 63)) != fi56), F1: uint16(v6), F2: v4}
                                                                                        switch 653 {
                                                                                        case 2147483647:
                                                                                            if v48 {
                                                                                                v91 := GenS9{F0: 335, F1: v3}
                                                                                                v92 := GenS1{F0: 1}
                                                                                                _ = v91
                                                                                                _ = v92
                                                                                            }
                                                                                        case 0:
                                                                                            if v48 {
                                                                                                v77 += (fi55 + 32768)
                                                                                                var v93 int = (v77 * (-(fi56)))
                                                                                                _ = v93
                                                                                            }
                                                                                        default:
                                                                                            if v48 {
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                if v48 {
                                                                                                    delete(v1, v4)
                                                                                                    var v94 float64 = float64(0.0)
                                                                                                    _ = v94
                                                                                                } else {
                                                                                                    var v95 uint8 = (v3 / ((v3) | 1))
                                                                                                    _ = v95
                                                                                                }
                                                                                                v4 -= (v4 ^ 1)
                                                                                                v81 = 701
                                                                                                var v96 float64 = (3.14159 * (0.1 / 1e10))
                                                                                                _ = v96
                                                                                            } else {
                                                                                                fmt.Println(v82, v77, v81)
                                                                                                _ = len(v2)
                                                                                            }
                                                                                        }
                                                                                        _ = v89
                                                                                        _ = v90
                                                                                    }
                                                                                }
                                                                                v48 = v48
                                                                            } else {
                                                                                fmt.Println(v77, v81, fi56)
                                                                            }
                                                                        }
                                                                        _ = v81
                                                                        _ = v82
                                                                    } else {
                                                                        _ = len(v1)
                                                                        _ = len(v5)
                                                                        fmt.Println(v3, v3)
                                                                    }
                                                                default:
                                                                    v5 = append(v5, v4)
                                                                }
                                                                _ = v79
                                                            }
                                                        case "foo bar":
                                                            if false {
                                                                if (v48 || v48) {
                                                                    v97 := GenS4{F0: 0}
                                                                    var v98 uint64 = uint64(v3)
                                                                    var v99 uint16 = 1
                                                                    _ = v97
                                                                    _ = v98
                                                                    _ = v99
                                                                } else {
                                                                    var v100 string = "foo bar"
                                                                    var v101 float32 = 1.5
                                                                    _ = v100
                                                                    _ = v101
                                                                }
                                                                v102 := GenS10{F0: v4, F1: -1, F2: (^(v4))}
                                                                var v103 uint32 = (v4 << ((4294967295) & 63))
                                                                v104 := GenS11{F0: v6, F1: (^(v3)), F2: v6}
                                                                _ = v102
                                                                _ = v103
                                                                _ = v104
                                                            }
                                                            v77 += (fi55 << ((674) & 63))
                                                            var v105 bool = (628 != 42)
                                                            _ = v105
                                                        default:
                                                            v106 := GenS12{F0: 42, F1: (v4 ^ (v4 * 1)), F2: 128}
                                                            _ = v106
                                                        }
                                                    default:
                                                        var v107 uint16 = 63
                                                        _ = v107
                                                    }
                                                }
                                                var v108 int16 = 185
                                                var v109 int32 = -1
                                                _ = v108
                                                _ = v109
                                                _ = v78
                                            } else {
                                                fmt.Println(fi56, v48)
                                                var v110 uint64 = uint64(v4)
                                                _ = v110
                                            }
                                            v77 = (-(fi55))
                                            _ = v77
                                        }
                                    default:
                                        v3 = v3
                                    }
                                    if true {
                                        v111 := GenS6{F0: -32768}
                                        _ = len(v1)
                                        _ = v111
                                    }
                                    _ = fi56
                                }
                                v6 -= (^(v6))
                                switch 9223372036854775807 {
                                case 32768:
                                    v48 = (!false)
                                case 63:
                                    v112 := GenS4{F0: 488}
                                    _ = v112
                                case 890:
                                    switch fi55 {
                                    case -2147483648:
                                        _ = len(v1)
                                    case -910:
                                        var v113 int8 = int8(v4)
                                        _ = v113
                                    case 32767:
                                        _ = len(v2)
                                    }
                                default:
                                    fmt.Println(v6)
                                }
                                _ = fi55
                            }
                        case "foo bar":
                            if (!v48) {
                                var v114 string = ""
                                for fi115 := 0; fi115 < 20; fi115++ {
                                    v116 := GenS2{F0: 2, F1: (v6 << ((0) & 63)), F2: 1}
                                    var v117 int32 = 1
                                    _ = fi115
                                    _ = v116
                                    _ = v117
                                }
                                _ = v114
                            }
                        }
                    } else {
                        var v118 int = 9223372036854775807
                        var v119 float32 = (0.1 * 1.5)
                        v118 += (v118 ^ (v118 / ((-984) | 1)))
                        _ = v118
                        _ = v119
                    }
                    if len(v5) > 0 {
                        v5[0] = v4
                    }
                } else {
                    var v120 string = ("" + "\n\t")
                    if (false && v48) {
                        if v48 {
                            var v121 uint64 = v6
                            switch 1 {
                            case -737:
                                var v122 float64 = 3.14159
                                _ = v122
                            case 9223372036854775807:
                                var v123 string = "foo bar"
                                _ = v123
                            default:
                                switch 32768 {
                                case -2147483648:
                                    v3 -= 255
                                case -606:
                                    switch 128 {
                                    case -254:
                                        if (!false) {
                                            for fi124 := 0; fi124 < 19; fi124++ {
                                                fmt.Println(v48)
                                                v125 := GenS13{F0: 42, F1: 0}
                                                if (32767 < 32767) {
                                                    if v48 {
                                                        if (32767 < 1000) {
                                                            fmt.Println(v121)
                                                            switch (fi124 >> ((0) & 63)) {
                                                            case -523:
                                                                var v126 uint16 = 255
                                                                _ = v126
                                                            }
                                                            fmt.Println(v120)
                                                        }
                                                        var v127 float64 = 0.1
                                                        _ = v127
                                                    }
                                                    v120 = v120
                                                } else {
                                                    v128 := v1[v4]
                                                    _ = v128
                                                }
                                                _ = fi124
                                                _ = v125
                                            }
                                            v129 := GenS7{F0: v120}
                                            if false {
                                                fmt.Println(v120, v4, v4)
                                                if (0.0 > ((-2.25 * (0.1 + 1e10)) + -2.25)) {
                                                    v3 = 197
                                                    v130 := GenS14{F0: (v6 / ((v6) | 1)), F1: -32768, F2: v48}
                                                    var v131 int8 = -128
                                                    _ = v130
                                                    _ = v131
                                                } else {
                                                    _ = len(v1)
                                                    if len(v2) > 0 {
                                                        v2[0] = v4
                                                    }
                                                }
                                            }
                                            _ = v129
                                        }
                                    }
                                case -1:
                                    _ = len(v2)
                                default:
                                    v121 = (^(v6))
                                }
                            }
                            _ = v121
                        } else {
                            for fi132 := 0; fi132 < 4; fi132++ {
                                for fi133 := 0; fi133 < 7; fi133++ {
                                    if v48 {
                                        if v48 {
                                            var v134 int16 = -973
                                            _ = len(v2)
                                            v48 = (-1 < -4)
                                            if (!v48) {
                                                switch "a" {
                                                case "\n\t":
                                                    var v135 int8 = 48
                                                    _ = v135
                                                case "":
                                                    switch fi133 {
                                                    case 183:
                                                        v3 += v3
                                                    case 520:
                                                        if false {
                                                            v120 = v120
                                                            if v48 {
                                                                if (int8(v6) < -1) {
                                                                    if (uint16(v3) == 32767) {
                                                                        var v136 int8 = 127
                                                                        _ = v136
                                                                    } else {
                                                                        v6 = v6
                                                                        if (18446744073709551615 <= 18446744073709551615) {
                                                                            var v137 uint = 2
                                                                            _ = v137
                                                                        } else {
                                                                            fmt.Println(v4, fi133, v134)
                                                                            v134 -= -208
                                                                            v138 := GenS8{F0: true, F1: 2, F2: uint32(v3)}
                                                                            _ = v138
                                                                        }
                                                                        switch 32767 {
                                                                        case 9223372036854775807:
                                                                            fmt.Println(fi133, fi133, v3)
                                                                        }
                                                                        switch v120 {
                                                                        case "":
                                                                            v3 = (v3 - (v3 + (^(v3))))
                                                                        case "a":
                                                                            var v139 uint64 = (v6 + (^(v6)))
                                                                            _ = v139
                                                                        }
                                                                        v134 = int16(v6)
                                                                        v120 = (v120 + "\n\t")
                                                                    }
                                                                }
                                                                var v140 uint32 = (v4 - uint32(v3))
                                                                v141 := GenS3{F0: 2, F1: ((v134 == v134) || false), F2: "a"}
                                                                _ = v140
                                                                _ = v141
                                                            } else {
                                                                fmt.Println(fi132, v120, v6)
                                                                if false {
                                                                    var v142 int8 = -128
                                                                    _ = len(v5)
                                                                    var v143 int8 = 81
                                                                    v144 := GenS10{F0: 2147483648, F1: int64(v6), F2: v4}
                                                                    _ = v142
                                                                    _ = v143
                                                                    _ = v144
                                                                } else {
                                                                    v48 = (v3 <= v3)
                                                                    var v145 string = ((((v120 + v120) + v120) + v120) + v120)
                                                                    if (((v145 + (v145 + v145)) > ("hello" + v145)) || v48) {
                                                                        v1[(^(v4))] = (v4 / ((v4) | 1))
                                                                        v3 -= 127
                                                                        if (65535 < 255) {
                                                                            v120 = "foo bar"
                                                                            v6 += (v6 * (v6 | v6))
                                                                            fmt.Println(v4, v120)
                                                                        }
                                                                    } else {
                                                                        v6 -= v6
                                                                        v146 := GenS3{F0: 1, F1: (v48 && false), F2: v145}
                                                                        v48 = v48
                                                                        _ = v146
                                                                    }
                                                                    v48 = ((^(v4)) <= (v4 % ((974) | 1)))
                                                                    var v147 int16 = (v134 - v134)
                                                                    _ = v145
                                                                    _ = v147
                                                                }
                                                                var v148 int64 = 701
                                                                _ = v148
                                                            }
                                                            v134 += (^(v134))
                                                        }
                                                    case -504:
                                                        if v48 {
                                                            if v48 {
                                                                var v149 int8 = 60
                                                                _ = len(v5)
                                                                var v150 int = fi132
                                                                var v151 uint32 = 415
                                                                _ = v149
                                                                _ = v150
                                                                _ = v151
                                                            } else {
                                                                v3 = (v3 - 128)
                                                                v4 += (v4 * (v4 * 42))
                                                            }
                                                            var v152 int = fi133
                                                            _ = v152
                                                        } else {
                                                            v3 = v3
                                                        }
                                                    }
                                                case "hello":
                                                    if len(v2) > 0 {
                                                        v2[0] = 65535
                                                    }
                                                default:
                                                    v6 -= v6
                                                }
                                                v153 := GenS6{F0: (v134 | 32767)}
                                                _ = v153
                                            } else {
                                                var v154 uint8 = uint8(v4)
                                                _ = v154
                                            }
                                            v155 := GenS3{F0: 0, F1: ((^(v3)) != v3), F2: v120}
                                            fmt.Println(v4)
                                            _ = v134
                                            _ = v155
                                        } else {
                                            if v48 {
                                                switch (v120 + (((((v120 + v120) + "foo bar") + "hello") + ("hello" + v120)) + v120)) {
                                                case "a":
                                                    var v156 int8 = 61
                                                    _ = v156
                                                case "foo bar":
                                                    v157 := GenS5{F0: (v4 / ((v4) | 1)), F1: 255}
                                                    _ = v157
                                                case "":
                                                    switch (fi132 * fi132) {
                                                    case 4294967295:
                                                        v158 := GenS15{F0: 2}
                                                        _ = v158
                                                    case 2147483647:
                                                        var v159 bool = v48
                                                        _ = v159
                                                    case 128:
                                                        v160 := GenS8{F0: (-32768 == 32), F1: 1, F2: (^(v4))}
                                                        fmt.Println(fi133)
                                                        _ = v160
                                                    default:
                                                        switch v120 {
                                                        case "":
                                                            if v48 {
                                                                v161 := GenS16{F0: -419, F1: (v3 + 139), F2: 2147483647}
                                                                v120 = v120
                                                                _ = v161
                                                            } else {
                                                                switch fi132 {
                                                                case 127:
                                                                    v120 = (v120 + ((("\n\t" + "a") + (("hello" + (((v120 + "a") + v120) + (((v120 + "hello") + "foo bar") + "hello"))) + ((("hello" + "hello") + (v120 + ((v120 + v120) + v120))) + (v120 + ((v120 + "a") + "hello"))))) + ("\n\t" + (v120 + v120))))
                                                                }
                                                                v3 = (^(v3))
                                                                _ = len(v1)
                                                            }
                                                        default:
                                                            if v48 {
                                                                var v162 uint16 = 65535
                                                                _ = v162
                                                            }
                                                        }
                                                    }
                                                }
                                                v163 := GenS17{F0: 1000, F1: 127, F2: 255}
                                                _ = len(v1)
                                                _ = v163
                                            }
                                        }
                                        var v164 uint = 65535
                                        v48 = (!v48)
                                        _ = v164
                                    }
                                    _ = len(v2)
                                    _ = fi133
                                }
                                if v48 {
                                    var v165 int64 = 1
                                    fmt.Println(v4, v3)
                                    _ = v165
                                }
                                _ = fi132
                            }
                        }
                        v166 := GenS14{F0: (v6 * (v6 / ((v6) | 1))), F1: 520, F2: true}
                        fmt.Println(v120)
                        _ = v166
                    } else {
                        var v167 int8 = -1
                        _ = v167
                    }
                    _ = v120
                }
            case "\n\t":
                v168 := GenS12{F0: 0, F1: (^(v4)), F2: 32768}
                _ = v168
            }
        }
        _ = v48
    }
    v169 := GenS13{F0: 458, F1: 0}
    v4 -= (v4 >> ((42) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v169
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v47
    _ = v5
    _ = v6
}
