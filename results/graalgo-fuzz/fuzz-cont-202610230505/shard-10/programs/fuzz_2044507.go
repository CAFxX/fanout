package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint; F1 int; F2 uint }
type GenS8 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 uint16; F1 int8 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 string; F1 uint8 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint64; F1 int8 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int32; F1 bool }
type GenS17 struct { F0 int }
type GenS18 struct { F0 int16 }
func genFn0(p0 uint64, p1 uint, p2 uint32) int64 {
    fmt.Println(p1)
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 uint16 = 1
        v5 = (v5 & (^(v5)))
        v6 := GenS1{F0: (fi4 << (((p1 / (((p1 << ((p1) & 63))) | 1))) & 63))}
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: -183}
    if false {
        if true {
            var v8 uint8 = 16
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int64 = 0
    v3 := []string{("foo bar" + "hello"), "\n\t", ""}
    switch 42 {
    case -224:
        var v9 int32 = 1
        _ = v9
    }
    v3 = append(v3, (("foo bar" + "hello") + "hello"))
    v2 += -586
    for fi10 := 0; fi10 < 6; fi10++ {
        v3 = append(v3, ("" + ("a" + "a")))
        if (("foo bar" + "a") != "") {
            if true {
                if false {
                    var v11 int16 = 42
                    v12 := GenS2{F0: 0}
                    _ = v11
                    _ = v12
                } else {
                    switch "\n\t" {
                    case "hello":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
                switch 2147483647 {
                case 0:
                    v1 = 65535
                case -1:
                    var v13 uint = 907
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 7; fi14++ {
                        v2 = v2
                        switch 13 {
                        case 127:
                            v15 := GenS2{F0: 2}
                            _ = v15
                        case 1:
                            var v16 int64 = v2
                            _ = v16
                        case 489:
                            v17 := GenS3{F0: 42}
                            _ = v17
                        default:
                            var v18 int = 32768
                            var v19 uint16 = 0
                            fmt.Println(fi10, v18)
                            _ = v18
                            _ = v19
                        }
                        v2 += (v2 & (v2 + (^(v2))))
                        _ = fi14
                    }
                }
            }
        } else {
            var v20 int16 = int16(v2)
            switch (fi10 >> ((18446744073709551615) & 63)) {
            case -225:
                _ = len(v3)
            case 1:
                v20 = (v20 / (((v20 * v20)) | 1))
            case -419:
                if (true && false) {
                    for fi21 := 0; fi21 < 12; fi21++ {
                        v20 = (v20 * (v20 * (v20 / ((v20) | 1))))
                        var v22 bool = (!(-123 < 0))
                        if true {
                            var v23 int8 = -128
                            v23 += (v23 >> ((2) & 63))
                            var v24 bool = false
                            _ = v23
                            _ = v24
                        }
                        _ = fi21
                        _ = v22
                    }
                    for fi25 := 0; fi25 < 5; fi25++ {
                        var v26 uint8 = 128
                        _ = fi25
                        _ = v26
                    }
                }
            }
            _ = v20
        }
        _ = fi10
    }
    if true {
        var v27 string = (("foo bar" + "foo bar") + ("foo bar" + (("hello" + ("\n\t" + ("" + ""))) + "")))
        _ = v27
    } else {
        _ = len(v3)
        var v28 uint8 = 127
        _ = v28
    }
    v29 := GenS1{F0: 65}
    v2 = int64(v1)
    v2 += v2
    fmt.Println(v1, v2)
    var v30 int16 = 1
    var v31 uint8 = 127
    for fi32 := 0; fi32 < 15; fi32++ {
        for fi33 := 0; fi33 < 1; fi33++ {
            var v34 int = (fi32 | int(v1))
            switch (fi33 % ((fi32) | 1)) {
            case 2147483647:
                v34 += -244
            case -9223372036854775808:
                var v35 uint64 = 18446744073709551615
                _ = v35
            default:
                if (v31 == v31) {
                    _ = len(v3)
                    switch "\n\t" {
                    case "foo bar":
                        v36 := GenS4{F0: false, F1: (fi32 % ((1) | 1))}
                        _ = v36
                    case "\n\t":
                        v2 -= genFn0(982, 656, v1)
                    }
                } else {
                    var v37 int16 = (v30 + (v30 % ((v30) | 1)))
                    switch (v34 - fi33) {
                    case -981:
                        fmt.Println(v31, v1)
                        v38 := GenS1{F0: (fi32 | -925)}
                        v39 := GenS5{F0: -128}
                        _ = v38
                        _ = v39
                    case 1:
                        if (!(!(!(-123 < 63)))) {
                            var v40 uint64 = 18446744073709551615
                            var v41 int64 = (v2 & genFn0((v40 - 828), 18446744073709551615, (v1 - v1)))
                            var v42 bool = false
                            v2 -= v2
                            v34 += (fi33 >> ((42) & 63))
                            _ = v40
                            _ = v41
                            _ = v42
                        } else {
                            v31 = v31
                            switch (v34 / ((fi32) | 1)) {
                            case 9223372036854775807:
                                _ = len(v3)
                            case -9223372036854775808:
                                if false {
                                    var v43 string = ""
                                    v34 -= (fi32 << ((2) & 63))
                                    v34 += 1
                                    _ = v43
                                } else {
                                    v30 -= (^(v37))
                                    if ((float32((1e10 + (0.0 - -2.25))) * 0.0) != (0.1 + -2.25)) {
                                        v3 = append(v3, ("hello" + (("foo bar" + ("foo bar" + "\n\t")) + "")))
                                        var v44 uint32 = 925
                                        _ = v44
                                    }
                                }
                            case 4294967295:
                                var v45 uint64 = 547
                                _ = v45
                            default:
                                var v46 uint16 = 0
                                _ = v46
                            }
                            v3 = append(v3, ((("" + ("" + "\n\t")) + "hello") + "hello"))
                        }
                    case -225:
                        v47 := GenS3{F0: 191}
                        var v48 string = (("\n\t" + ("" + ("foo bar" + ("foo bar" + "foo bar")))) + "")
                        var v49 uint = 4294967295
                        _ = v47
                        _ = v48
                        _ = v49
                    }
                    switch (fi33 & (-(v34))) {
                    case 32767:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    default:
                        var v50 uint8 = 1
                        _ = v50
                    }
                    _ = v37
                }
            }
            if true {
                switch -513 {
                case 2147483647:
                    switch fi32 {
                    case 4294967295:
                        v34 = fi33
                    case 32767:
                        if false {
                            switch (fi33 / ((9223372036854775807) | 1)) {
                            case 9223372036854775807:
                                switch v34 {
                                case 818:
                                    fmt.Println(fi33, v1, v30)
                                case -464:
                                    if (((v31 - (v31 * (v31 ^ 119))) >= v31) || true) {
                                        v30 = int16(v2)
                                    }
                                case -1:
                                    if true {
                                        var v51 float64 = (0.0 + 1e10)
                                        if (1 <= 1) {
                                            v34 = (fi33 * 2147483647)
                                            switch v34 {
                                            case 609:
                                                fmt.Println(fi32)
                                            case 128:
                                                v1 -= v1
                                            default:
                                                v31 -= (v31 ^ (^(v31)))
                                            }
                                            v51 += v51
                                            v52 := GenS6{F0: 627}
                                            _ = v52
                                        } else {
                                            var v53 int64 = v2
                                            var v54 bool = (!true)
                                            _ = v53
                                            _ = v54
                                        }
                                        _ = v51
                                    }
                                default:
                                    v30 = v30
                                }
                            case 42:
                                v31 = 128
                            default:
                                switch ("foo bar" + "\n\t") {
                                case "":
                                    var v55 uint = 225
                                    _ = v55
                                case "a":
                                    switch (("foo bar" + "") + ("foo bar" + "foo bar")) {
                                    case "hello":
                                        v3 = append(v3, "foo bar")
                                        var v56 uint64 = 757
                                        _ = v56
                                    }
                                default:
                                    fmt.Println(v34)
                                }
                                if false {
                                    v34 = (fi32 - fi32)
                                    var v57 int16 = (v30 + 1000)
                                    _ = v57
                                }
                                switch "\n\t" {
                                case "":
                                    if true {
                                        v2 = (^(v2))
                                        switch fi32 {
                                        case -2147483648:
                                            var v58 int16 = v30
                                            _ = v58
                                        case 127:
                                            var v59 uint8 = v31
                                            _ = v59
                                        case 9223372036854775807:
                                            _ = len(v3)
                                        default:
                                            switch (^(v34)) {
                                            case 128:
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v1 += v1
                                                default:
                                                    v1 -= v1
                                                }
                                            case -9223372036854775808:
                                                var v60 int8 = 100
                                                _ = v60
                                            case -1:
                                                var v61 float32 = float32((((3.14159 / -2.25) / -2.25) - 3.14159))
                                                _ = v61
                                            }
                                        }
                                    } else {
                                        v1 = 671
                                        v34 = (fi33 >> ((43) & 63))
                                        if (uint8(v1) != (v31 & v31)) {
                                            if len(v3) > 0 {
                                                v3[0] = "hello"
                                            }
                                            fmt.Println(v34, v30)
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        v34 = (-(fi32))
                                    }
                                case "a":
                                    if false {
                                        v62 := GenS7{F0: 791, F1: (^(fi33)), F2: 65535}
                                        _ = v62
                                    } else {
                                        if (true && true) {
                                            var v63 uint = 18446744073709551615
                                            var v64 int = (v34 / ((fi33) | 1))
                                            v34 -= (^(v64))
                                            var v65 int32 = 42
                                            if (!true) {
                                                v1 = v1
                                                v65 -= 0
                                            } else {
                                                var v66 int8 = 0
                                                _ = v66
                                            }
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        } else {
                                            var v67 uint = 0
                                            _ = v67
                                        }
                                    }
                                case "hello":
                                    if ((1.5 - float64(((((-2.25 + (3.14159 / 1.5)) - (0.1 + 3.14159)) / 3.14159) * (3.14159 - (1.5 - (((0.0 / 1.5) * -2.25) - 0.0)))))) <= -2.25) {
                                        v68 := GenS8{F0: 0, F1: 276, F2: 68}
                                        _ = v68
                                    } else {
                                        fmt.Println(v31, v31, v34)
                                        if false {
                                            if true {
                                                v69 := GenS7{F0: 255, F1: fi33, F2: 18446744073709551615}
                                                if false {
                                                    v34 = (fi32 | v34)
                                                    v30 += v30
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v70 := GenS9{F0: 42, F1: -885}
                                                        _ = v70
                                                    }
                                                }
                                                var v71 string = "foo bar"
                                                _ = v69
                                                _ = v71
                                            } else {
                                                v31 = uint8(v1)
                                                var v72 uint = 255
                                                if (32767 <= -1) {
                                                    fmt.Println(v31, v34, v2)
                                                    v73 := GenS5{F0: -128}
                                                    switch "a" {
                                                    case "a":
                                                        var v74 uint = 4294967295
                                                        _ = v74
                                                    case "foo bar":
                                                        v3 = append(v3, "")
                                                    case "":
                                                        var v75 int8 = 0
                                                        _ = v75
                                                    default:
                                                        var v76 uint = v72
                                                        _ = v76
                                                    }
                                                    v34 = v34
                                                    if len(v3) > 0 {
                                                        v3[0] = ("\n\t" + ("foo bar" + "hello"))
                                                    }
                                                    _ = v73
                                                } else {
                                                    var v77 uint32 = (^(v1))
                                                    _ = v77
                                                }
                                                _ = v72
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = "foo bar"
                                            }
                                            switch "a" {
                                            case "\n\t":
                                                v78 := GenS10{F0: 2, F1: 127}
                                                _ = v78
                                            case "":
                                                _ = len(v3)
                                                var v79 int8 = -92
                                                _ = v79
                                            case "foo bar":
                                                fmt.Println(fi33, v2)
                                            default:
                                                v3 = append(v3, ("\n\t" + "a"))
                                            }
                                            var v80 bool = false
                                            switch -118 {
                                            case 1:
                                                if false {
                                                    var v81 int8 = 127
                                                    v3 = append(v3, "hello")
                                                    _ = v81
                                                }
                                            case 0:
                                                if (42 >= (v30 / (((v30 * v30)) | 1))) {
                                                    if true {
                                                        if false {
                                                            if (((!(0.1 >= ((((((0.1 / 0.1) - 0.0) + ((3.14159 * 1.5) / float64(3.14159))) + 3.14159) / 1e10) - 3.14159))) && (486 == 255)) || v80) {
                                                                var v82 int64 = (v2 >> ((255) & 63))
                                                                fmt.Println(v80, v30)
                                                                v30 += v30
                                                                _ = v82
                                                            }
                                                        }
                                                    } else {
                                                        v83 := GenS2{F0: v31}
                                                        _ = v83
                                                    }
                                                } else {
                                                    v34 = 2147483647
                                                }
                                            }
                                            _ = v80
                                        } else {
                                            v84 := GenS11{F0: -995}
                                            v30 = (-(v30))
                                            fmt.Println(v31, v34)
                                            _ = v84
                                        }
                                    }
                                default:
                                    v2 -= (v2 >> ((65535) & 63))
                                }
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                }
            } else {
                if true {
                    v85 := GenS5{F0: -128}
                    if (-479 < -2147483648) {
                        var v86 int32 = 0
                        if false {
                            fmt.Println(fi32, v1, v31)
                        }
                        if true {
                            switch v34 {
                            case 607:
                                var v87 int16 = v30
                                _ = v87
                            case -9223372036854775808:
                                switch fi32 {
                                case 557:
                                    fmt.Println(v30, v1)
                                    var v88 int8 = 35
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    _ = v88
                                case 1:
                                    v89 := GenS12{F0: "hello", F1: (v31 / (((v31 >> ((870) & 63))) | 1))}
                                    _ = v89
                                case 32768:
                                    v3 = append(v3, ("a" + "hello"))
                                }
                            case 9223372036854775807:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                v90 := GenS11{F0: v86}
                                _ = v90
                            }
                            var v91 uint16 = 42
                            _ = v91
                        } else {
                            var v92 int32 = (v86 << ((18446744073709551615) & 63))
                            fmt.Println(fi32, v31)
                            _ = v92
                        }
                        _ = v86
                    }
                    _ = v85
                } else {
                    v34 += 154
                    var v93 uint32 = 0
                    v34 += (v34 >> ((4294967295) & 63))
                    _ = v93
                }
                if true {
                    v2 += 0
                    v94 := GenS10{F0: 65535, F1: 127}
                    switch (v34 ^ fi33) {
                    case 2147483647:
                        if (!true) {
                            var v95 int64 = 9223372036854775807
                            var v96 uint16 = 42
                            switch (fi33 & (-(fi32))) {
                            case 4294967295:
                                fmt.Println(v30)
                            case -2147483648:
                                if false {
                                    switch "\n\t" {
                                    case "a":
                                        v34 -= fi32
                                    case "":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    default:
                                        var v97 int32 = -1
                                        v98 := GenS13{F0: v2}
                                        var v99 uint32 = 2
                                        var v100 int8 = 100
                                        _ = v100
                                        _ = v97
                                        _ = v98
                                        _ = v99
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        if true {
                                            var v101 uint32 = 4294967295
                                            _ = v101
                                        } else {
                                            var v102 float32 = float32(-2.25)
                                            if false {
                                                var v103 uint8 = (v31 & v31)
                                                _ = v103
                                            }
                                            switch "\n\t" {
                                            case "hello":
                                                fmt.Println(v1)
                                            }
                                            _ = v102
                                        }
                                    default:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if false {
                                            switch "a" {
                                            case "a":
                                                v104 := GenS1{F0: fi32}
                                                v31 = (v31 - v31)
                                                v31 -= (v31 | (^(v31)))
                                                _ = v104
                                            case "foo bar":
                                                if (!false) {
                                                    switch (("foo bar" + "\n\t") + (("" + (("" + "") + "")) + ("\n\t" + "a"))) {
                                                    case "a":
                                                        _ = len(v3)
                                                        if len(v3) > 0 {
                                                            v3[0] = "\n\t"
                                                        }
                                                    case "hello":
                                                        if len(v3) > 0 {
                                                            v3[0] = ("" + "hello")
                                                        }
                                                    }
                                                } else {
                                                    if true {
                                                        var v105 bool = false
                                                        v95 += int64(v1)
                                                        v106 := GenS14{F0: 0, F1: 100}
                                                        _ = v105
                                                        _ = v106
                                                    } else {
                                                        v30 += 32767
                                                        if (!(("hello" + "\n\t") == ("hello" + "a"))) {
                                                            if (v31 > v31) {
                                                                v95 = int64(v30)
                                                                var v107 int8 = 100
                                                                if (false || (2 <= uint32(v31))) {
                                                                    var v108 int16 = v30
                                                                    if (v30 <= (^(v108))) {
                                                                        _ = len(v3)
                                                                    } else {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = "a"
                                                                        }
                                                                        if len(v3) > 0 {
                                                                            v3[0] = "\n\t"
                                                                        }
                                                                        var v109 int8 = int8(v1)
                                                                        _ = v109
                                                                    }
                                                                    _ = v108
                                                                } else {
                                                                    var v110 int8 = 127
                                                                    _ = v110
                                                                }
                                                                _ = v107
                                                            }
                                                            if true {
                                                                v111 := GenS2{F0: 42}
                                                                switch ("hello" + "") {
                                                                case "hello":
                                                                    var v112 uint16 = (v96 >> ((2) & 63))
                                                                    _ = v112
                                                                case "a":
                                                                    if (true || true) {
                                                                        fmt.Println(fi32)
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (("" + ("" + "hello")) + "foo bar")
                                                                        }
                                                                        v31 += (v31 % ((v31) | 1))
                                                                    } else {
                                                                        v30 -= (-(v30))
                                                                    }
                                                                case "\n\t":
                                                                    var v113 uint16 = (^(v96))
                                                                    _ = v113
                                                                default:
                                                                    if len(v3) > 0 {
                                                                        v3[0] = "foo bar"
                                                                    }
                                                                }
                                                                _ = v111
                                                            }
                                                            var v114 bool = true
                                                            _ = v114
                                                        }
                                                        if false {
                                                            v2 -= 964
                                                            if len(v3) > 0 {
                                                                v3[0] = ("foo bar" + "\n\t")
                                                            }
                                                            v115 := GenS15{F0: v30}
                                                            _ = v115
                                                        }
                                                    }
                                                    fmt.Println(v96, v96, v30)
                                                    v95 += 1
                                                    v3 = append(v3, "hello")
                                                }
                                            }
                                            v31 = (v31 << ((42) & 63))
                                            if len(v3) > 0 {
                                                v3[0] = "a"
                                            }
                                            v34 = (fi33 | fi32)
                                            var v116 uint = 42
                                            _ = v116
                                        }
                                        v3 = append(v3, "")
                                        v117 := GenS15{F0: (v30 + (v30 ^ v30))}
                                        if true {
                                            fmt.Println(v2, fi32)
                                            switch fi32 {
                                            case -2147483648:
                                                switch v34 {
                                                case 9223372036854775807:
                                                    var v118 int64 = (v95 % (((v2 * v2)) | 1))
                                                    _ = v118
                                                default:
                                                    var v119 uint64 = 220
                                                    _ = v119
                                                }
                                            case 9223372036854775807:
                                                var v120 uint16 = (v96 / ((1) | 1))
                                                _ = v120
                                            default:
                                                v121 := GenS2{F0: (v31 % ((1) | 1))}
                                                v34 -= (fi33 | v34)
                                                _ = v121
                                            }
                                        } else {
                                            fmt.Println(v2, v30, v31)
                                            var v122 float32 = 1e10
                                            v123 := GenS5{F0: 1}
                                            v2 += v2
                                            v3 = append(v3, "a")
                                            _ = v122
                                            _ = v123
                                        }
                                        _ = v117
                                    }
                                } else {
                                    v3 = append(v3, ("a" + "foo bar"))
                                    v124 := GenS16{F0: 752, F1: (v1 >= 65535)}
                                    _ = v124
                                }
                            default:
                                v34 -= fi32
                            }
                            _ = v95
                            _ = v96
                        }
                    case -2147483648:
                        fmt.Println(v2, v31)
                    case 32767:
                        v125 := GenS17{F0: (v34 ^ (fi32 % ((fi33) | 1)))}
                        _ = v125
                    default:
                        v1 = v1
                        var v126 int32 = 382
                        v127 := GenS18{F0: 1}
                        v126 = (v126 << ((uint(v31)) & 63))
                        _ = v126
                        _ = v127
                    }
                    v128 := GenS17{F0: (fi33 % (((fi32 - fi33)) | 1))}
                    _ = v128
                    _ = v94
                } else {
                    var v129 uint16 = 2
                    v2 = 0
                    v129 -= (^(v129))
                    _ = v129
                }
            }
            _ = fi33
            _ = v34
        }
        var v130 int16 = v30
        _ = fi32
        _ = v130
    }
    v3 = append(v3, ("" + ("\n\t" + ("\n\t" + ("" + "hello")))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
