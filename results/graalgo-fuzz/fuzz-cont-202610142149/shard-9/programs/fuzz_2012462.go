package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint32; F2 int }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 uint; F1 int; F2 int64 }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 uint8; F1 int64 }
type GenS8 struct { F0 uint32; F1 int; F2 bool }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint64; F1 uint32 }
type GenS12 struct { F0 int8; F1 uint16 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int64; F1 string; F2 string }
type GenS16 struct { F0 string }

func main() {
    var v1 int16 = -361
    v2 := []int{}
    var v3 uint8 = 26
    var v4 int = 127
    var v5 int32 = 2147483647
    var v6 int8 = -1
    _ = len(v2)
    var v7 uint = 18446744073709551615
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint = 18446744073709551615
        switch (fi8 % ((v4) | 1)) {
        case 0:
            if true {
                var v10 bool = (!true)
                _ = v10
            } else {
                v4 -= 128
                if ((!("" > "\n\t")) || true) {
                    v11 := GenS1{F0: 896}
                    _ = v11
                }
            }
        case -583:
            v3 += (v3 + v3)
        }
        _ = fi8
        _ = v9
    }
    v5 = v5
    var v12 uint32 = 2147483648
    var v13 uint64 = 18446744073709551615
    fmt.Println(v12, v4)
    for fi14 := 0; fi14 < 13; fi14++ {
        if ((true && ("" == "a")) && false) {
            _ = len(v2)
        } else {
            for fi15 := 0; fi15 < 5; fi15++ {
                var v16 uint64 = v13
                var v17 uint16 = 32767
                v2 = append(v2, (fi15 ^ (v4 | (fi14 * fi14))))
                _ = len(v2)
                var v18 uint16 = 255
                fmt.Println(v3)
                if false {
                    if len(v2) > 0 {
                        v2[0] = -9223372036854775808
                    }
                    v5 += (v5 >> (((v7 & (v7 ^ v7))) & 63))
                } else {
                    var v19 float64 = 1e10
                    v20 := GenS2{F0: -1, F1: (v7 >> ((v7) & 63)), F2: v3}
                    if false {
                        var v21 int = fi15
                        switch int(v7) {
                        case 9223372036854775807:
                            v5 -= (v5 & 42)
                            var v22 uint16 = uint16(v3)
                            v23 := GenS3{F0: (v5 * (v5 ^ v5)), F1: (v12 >> (((v7 << (((^(v7))) & 63))) & 63)), F2: (-(v4))}
                            _ = v22
                            _ = v23
                        }
                        var v24 int32 = int32(v6)
                        if true {
                            switch "a" {
                            case "hello":
                                fmt.Println(fi14)
                            case "\n\t":
                                var v25 int8 = int8(v1)
                                _ = v25
                            case "":
                                v26 := GenS4{F0: int(v5), F1: (v3 << (((v7 ^ v7)) & 63))}
                                _ = v26
                            default:
                                if len(v2) > 0 {
                                    v2[0] = v21
                                }
                            }
                            v27 := GenS3{F0: 428, F1: (v12 >> ((255) & 63)), F2: (fi14 / (((v21 * int(v5))) | 1))}
                            _ = v27
                        }
                        v28 := GenS5{F0: v7, F1: (-(fi15)), F2: int64(v17)}
                        v17 += (^(v18))
                        _ = len(v2)
                        switch "a" {
                        case "a":
                            v29 := GenS3{F0: (v5 << (((^(v7))) & 63)), F1: v12, F2: v21}
                            _ = v29
                        case "\n\t":
                            switch "" {
                            case "":
                                if true {
                                    var v30 uint32 = (v12 >> (((v7 - (v7 + (v7 << (((v7 % (((v7 << (((v7 << ((2) & 63))) & 63))) | 1))) & 63))))) & 63))
                                    var v31 float64 = ((v19 - 0.0) / ((0.1 * 0.0) + v19))
                                    fmt.Println(v13, v12, fi15)
                                    _ = v30
                                    _ = v31
                                } else {
                                    switch 32768 {
                                    case 245:
                                        v32 := GenS5{F0: (v7 / ((v7) | 1)), F1: int(v5), F2: 1}
                                        _ = v32
                                    default:
                                        v21 -= (v4 + v4)
                                    }
                                    var v33 string = "foo bar"
                                    var v34 int8 = 1
                                    _ = v33
                                    _ = v34
                                }
                            default:
                                var v35 int32 = (v24 >> (((v7 % (((v7 << ((v7) & 63))) | 1))) & 63))
                                _ = v35
                            }
                        default:
                            if ((v1 / (((v1 ^ v1)) | 1)) <= (v1 << (((v7 + 670)) & 63))) {
                                if (-1 >= -610) {
                                    v24 = 2147483647
                                    v18 = v18
                                    v2 = append(v2, (v21 >> (((v7 / ((439) | 1))) & 63)))
                                }
                                var v36 int = (v21 % ((127) | 1))
                                if false {
                                    if ((v24 + v5) >= (v24 ^ v5)) {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        v37 := GenS6{F0: -9223372036854775808, F1: (^(fi14))}
                                        v2 = append(v2, fi14)
                                        _ = v37
                                    }
                                    if true {
                                        fmt.Println(v4, v1, fi14)
                                        fmt.Println(v3, v6, fi14)
                                    }
                                }
                                _ = v36
                            } else {
                                if (0 == v17) {
                                    _ = len(v2)
                                    if false {
                                        if ((v17 ^ 65535) != 42) {
                                            switch ("" + ((("\n\t" + "") + "foo bar") + "hello")) {
                                            case "a":
                                                v38 := GenS3{F0: (-(v24)), F1: (^(v12)), F2: (fi15 & int(v12))}
                                                _ = v38
                                            }
                                            if (!(((!((0.1 < 0.1) || (4294967295 > (v12 ^ v12)))) && true) || (v1 == (-(v1))))) {
                                                var v39 int = fi14
                                                var v40 bool = (!false)
                                                v2 = append(v2, (^(fi15)))
                                                _ = v39
                                                _ = v40
                                            }
                                        } else {
                                            var v41 uint16 = uint16(v13)
                                            v42 := GenS6{F0: -682, F1: int(v1)}
                                            fmt.Println(fi14)
                                            _ = v41
                                            _ = v42
                                        }
                                        if (false && (18446744073709551615 >= (v13 * (v13 ^ 569)))) {
                                            if (1 > -9223372036854775808) {
                                                var v43 uint16 = (v17 - (v17 / (((v17 + (v17 % ((v18) | 1)))) | 1)))
                                                if (42 < uint16(v16)) {
                                                    _ = len(v2)
                                                }
                                                _ = v43
                                            }
                                        }
                                        var v44 uint16 = v17
                                        _ = v44
                                    }
                                    fmt.Println(v5, v21, fi15)
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                var v45 uint = (v7 * (v7 % ((v7) | 1)))
                                v19 += (-2.25 - -2.25)
                                _ = v45
                            }
                        }
                        v46 := GenS6{F0: -855, F1: v4}
                        v47 := GenS5{F0: v7, F1: (^(fi14)), F2: 943}
                        _ = v21
                        _ = v24
                        _ = v28
                        _ = v46
                        _ = v47
                    } else {
                        switch (v4 / ((v4) | 1)) {
                        case 9223372036854775807:
                            v48 := GenS4{F0: 32768, F1: (v3 + v3)}
                            _ = v48
                        default:
                            switch 9223372036854775807 {
                            case -9:
                                v1 -= -32768
                                var v49 int = (fi15 >> ((v7) & 63))
                                v6 = v6
                                _ = v49
                            case -78:
                                if (v17 != (v18 << (((v7 - 5)) & 63))) {
                                    v6 -= v6
                                    v50 := GenS3{F0: 42, F1: 42, F2: v4}
                                    var v51 int8 = int8(v1)
                                    _ = v50
                                    _ = v51
                                } else {
                                    var v52 uint = v7
                                    var v53 uint = 2
                                    v2 = append(v2, (v4 / (((fi14 | 342)) | 1)))
                                    var v54 int64 = -9223372036854775808
                                    v1 -= (v1 % ((v1) | 1))
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                }
                            }
                        }
                        v55 := GenS7{F0: 29, F1: 9223372036854775807}
                        switch (fi15 % ((fi15) | 1)) {
                        case -512:
                            if true {
                                if true {
                                    var v56 string = ("foo bar" + ("a" + ("" + "a")))
                                    _ = v56
                                } else {
                                    v13 = v16
                                }
                                if (-9223372036854775808 > 32768) {
                                    if len(v2) > 0 {
                                        v2[0] = -161
                                    }
                                }
                            }
                        case 32768:
                            var v57 bool = (0.1 != -2.25)
                            _ = v57
                        }
                        _ = v55
                    }
                    if false {
                        var v58 int32 = 1
                        _ = v58
                    }
                    switch (-(fi15)) {
                    case 687:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v59 := GenS8{F0: 954, F1: (fi15 << (((v7 >> (((v7 * (v7 ^ 989))) & 63))) & 63)), F2: (v12 <= (v12 | v12))}
                        v16 += 0
                        var v60 float64 = ((v19 + 1.5) / v19)
                        _ = v59
                        _ = v60
                    case 9223372036854775807:
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if len(v2) > 0 {
                                v2[0] = (v4 * (fi15 >> ((65535) & 63)))
                            }
                            fmt.Println(v6)
                        }
                    }
                    _ = v19
                    _ = v20
                }
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            if ((v1 | (v1 * v1)) >= v1) {
                var v61 uint8 = 1
                if ("foo bar" >= "a") {
                    var v62 uint16 = 65535
                    switch "\n\t" {
                    case "\n\t":
                        for fi63 := 0; fi63 < 6; fi63++ {
                            if len(v2) > 0 {
                                v2[0] = fi63
                            }
                            _ = fi63
                        }
                    case "foo bar":
                        v5 -= (v5 << ((v7) & 63))
                    default:
                        if false {
                            if false {
                                switch (("\n\t" + "") + "a") {
                                case "":
                                    if false {
                                        for fi64 := 0; fi64 < 13; fi64++ {
                                            var v65 bool = ((v12 << (((v7 & 1)) & 63)) == (v12 - v12))
                                            fmt.Println(fi64)
                                            switch "" {
                                            case "\n\t":
                                                v6 += 42
                                            }
                                            _ = fi64
                                            _ = v65
                                        }
                                        for fi66 := 0; fi66 < 18; fi66++ {
                                            v13 += (v13 + (v13 / ((v13) | 1)))
                                            switch ("\n\t" + "foo bar") {
                                            case "hello":
                                                fmt.Println(v12, fi14, v1)
                                            default:
                                                if false {
                                                    v12 -= (v12 << ((42) & 63))
                                                }
                                            }
                                            _ = fi66
                                        }
                                    } else {
                                        if (0.1 > 1.5) {
                                            var v67 int16 = (v1 * -124)
                                            var v68 float64 = 0.0
                                            _ = v67
                                            _ = v68
                                        } else {
                                            var v69 int16 = (v1 - -949)
                                            _ = v69
                                        }
                                    }
                                case "\n\t":
                                    fmt.Println(v62, v62, v6)
                                }
                                v6 = v6
                                if true {
                                    if false {
                                        for fi70 := 0; fi70 < 19; fi70++ {
                                            v71 := GenS9{F0: (v12 % (((v12 >> ((607) & 63))) | 1))}
                                            v72 := GenS8{F0: (v12 - (v12 % (((v12 % ((uint32(v3)) | 1))) | 1))), F1: fi70, F2: (!(int64(fi14) < 9223372036854775807))}
                                            _ = fi70
                                            _ = v71
                                            _ = v72
                                        }
                                        v73 := GenS3{F0: int32(v7), F1: (v12 - v12), F2: -457}
                                        _ = v73
                                    } else {
                                        var v74 uint16 = (v62 | v62)
                                        v7 = (^(v7))
                                        _ = v74
                                    }
                                } else {
                                    _ = len(v2)
                                }
                            } else {
                                v75 := GenS10{F0: (v1 >> ((947) & 63))}
                                fmt.Println(v1)
                                _ = v75
                            }
                            if (1e10 > 1.5) {
                                for fi76 := 0; fi76 < 13; fi76++ {
                                    v4 -= (fi76 / ((-799) | 1))
                                    switch (-(fi14)) {
                                    case 128:
                                        if true {
                                            var v77 float64 = (0.0 / float64(3.14159))
                                            switch "hello" {
                                            case "":
                                                if true {
                                                    var v78 uint8 = (v61 & v61)
                                                    var v79 uint = (v7 << (((v7 * 18446744073709551615)) & 63))
                                                    _ = v78
                                                    _ = v79
                                                }
                                            case "a":
                                                v80 := GenS1{F0: v62}
                                                _ = v80
                                            case "foo bar":
                                                v77 = v77
                                                if (true && (13 >= (v6 | (v6 - int8(v62))))) {
                                                    v61 += (v61 & v61)
                                                    var v81 float64 = 3.14159
                                                    v82 := GenS7{F0: (v3 ^ 128), F1: 0}
                                                    if (0 < (v5 >> ((v7) & 63))) {
                                                        fmt.Println(v61)
                                                        var v83 int = 127
                                                        _ = len(v2)
                                                        _ = v83
                                                    } else {
                                                        if (false && (!true)) {
                                                            fmt.Println(v62)
                                                        } else {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        switch ("a" + ("\n\t" + "\n\t")) {
                                                        case "a":
                                                            if ((v7 >> ((570) & 63)) == (^(v7))) {
                                                                v61 = 127
                                                                if true {
                                                                    if (((v4 * -2147483648) <= (fi14 & v4)) && true) {
                                                                        if true {
                                                                            fmt.Println(fi14, fi76)
                                                                            v84 := GenS1{F0: (v62 >> (((v7 + (v7 & (v7 - 1)))) & 63))}
                                                                            _ = v84
                                                                        } else {
                                                                            v85 := GenS11{F0: v13, F1: 2147483647}
                                                                            var v86 float32 = (float32((v77 + 3.14159)) / 0.1)
                                                                            v87 := GenS7{F0: (v3 | 31), F1: -849}
                                                                            _ = v85
                                                                            _ = v86
                                                                            _ = v87
                                                                        }
                                                                    }
                                                                    if true {
                                                                        var v88 int = 711
                                                                        v89 := GenS10{F0: (v1 ^ -32768)}
                                                                        var v90 int32 = v5
                                                                        v77 = v77
                                                                        _ = v88
                                                                        _ = v89
                                                                        _ = v90
                                                                    } else {
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            switch (^(v4)) {
                                                                            case -396:
                                                                                v1 = v1
                                                                            default:
                                                                                if true {
                                                                                    v6 = (v6 / ((int8(v61)) | 1))
                                                                                    var v91 int16 = -21
                                                                                    _ = v91
                                                                                } else {
                                                                                    var v92 bool = false
                                                                                    _ = v92
                                                                                }
                                                                                v7 += 718
                                                                                fmt.Println(v12, v81)
                                                                            }
                                                                        case "\n\t":
                                                                            _ = len(v2)
                                                                        default:
                                                                            v93 := GenS10{F0: (v1 % ((801) | 1))}
                                                                            _ = v93
                                                                        }
                                                                        if false {
                                                                            v7 = (v7 >> ((uint(v12)) & 63))
                                                                            fmt.Println(v5, v62)
                                                                        } else {
                                                                            var v94 int = (fi76 + (-(fi14)))
                                                                            v77 = v81
                                                                            if len(v2) > 0 {
                                                                                v2[0] = fi14
                                                                            }
                                                                            _ = v94
                                                                        }
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        v95 := GenS4{F0: (fi76 % ((v4) | 1)), F1: (v61 & v3)}
                                                                        _ = v95
                                                                    }
                                                                } else {
                                                                    v96 := GenS8{F0: uint32(v3), F1: 4294967295, F2: true}
                                                                    if (!false) {
                                                                        fmt.Println(fi14)
                                                                        switch (^(fi76)) {
                                                                        case -208:
                                                                            v97 := GenS12{F0: v6, F1: (v62 - v62)}
                                                                            _ = v97
                                                                        default:
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                if ("hello" <= "a") {
                                                                                    v98 := GenS6{F0: 9223372036854775807, F1: fi14}
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        if ((v62 % ((uint16(v13)) | 1)) >= v62) {
                                                                                            switch (v4 | v4) {
                                                                                            case -237:
                                                                                                var v99 float32 = 0.0
                                                                                                _ = v99
                                                                                            case 32767:
                                                                                                var v100 uint = (^(v7))
                                                                                                _ = v100
                                                                                            case -2147483648:
                                                                                                var v101 uint8 = (v61 / (((v61 & uint8(v13))) | 1))
                                                                                                _ = v101
                                                                                            default:
                                                                                                v102 := GenS11{F0: 2, F1: 65535}
                                                                                                var v103 uint = v7
                                                                                                v4 = (fi14 ^ int(v62))
                                                                                                _ = v102
                                                                                                _ = v103
                                                                                            }
                                                                                            _ = len(v2)
                                                                                        }
                                                                                    case "":
                                                                                        var v104 uint8 = (v61 + 1)
                                                                                        _ = v104
                                                                                    case "a":
                                                                                        fmt.Println(fi14)
                                                                                    default:
                                                                                        var v105 float32 = 1e10
                                                                                        _ = v105
                                                                                    }
                                                                                    _ = v98
                                                                                }
                                                                            case "\n\t":
                                                                                v77 -= (3.14159 + 0.1)
                                                                            }
                                                                        }
                                                                        if (true || (true || (!(!false)))) {
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            v6 = -74
                                                                            v3 = (v61 | (v3 % ((v61) | 1)))
                                                                        }
                                                                    }
                                                                    v106 := GenS7{F0: v61, F1: 9223372036854775807}
                                                                    _ = v106
                                                                    _ = v96
                                                                }
                                                            } else {
                                                                switch "a" {
                                                                case "a":
                                                                    v6 = v6
                                                                case "\n\t":
                                                                    var v107 uint8 = 128
                                                                    _ = v107
                                                                case "":
                                                                    switch ("a" + "") {
                                                                    case "hello":
                                                                        var v108 uint32 = v12
                                                                        _ = v108
                                                                    case "\n\t":
                                                                        v109 := GenS13{F0: v12}
                                                                        _ = v109
                                                                    case "a":
                                                                        var v110 int16 = 428
                                                                        _ = v110
                                                                    }
                                                                }
                                                            }
                                                            v12 = (v12 + v12)
                                                            if (!true) {
                                                                if true {
                                                                    fmt.Println(v12, v61, v13)
                                                                    _ = len(v2)
                                                                } else {
                                                                    fmt.Println(v62)
                                                                }
                                                                if (false && (-510 >= 9223372036854775807)) {
                                                                    if true {
                                                                        switch (^(fi14)) {
                                                                        case -2147483648:
                                                                            if true {
                                                                                v12 = v12
                                                                                if true {
                                                                                    fmt.Println(v5, v13)
                                                                                    if ((0.0 - 0.0) > (float32(v77) - 3.14159)) {
                                                                                        var v111 int = int(v12)
                                                                                        _ = v111
                                                                                    } else {
                                                                                        if false {
                                                                                            if ((v7 << ((v7) & 63)) <= uint(v13)) {
                                                                                                _ = len(v2)
                                                                                                var v112 float64 = 1e10
                                                                                                var v113 uint = (^(v7))
                                                                                                _ = v112
                                                                                                _ = v113
                                                                                            }
                                                                                            v114 := GenS14{F0: ""}
                                                                                            var v115 uint16 = 340
                                                                                            _ = v114
                                                                                            _ = v115
                                                                                        } else {
                                                                                            var v116 uint = (^(v7))
                                                                                            _ = v116
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v117 bool = ((uint8(v62) != (v3 | v61)) && false)
                                                                                _ = v117
                                                                            } else {
                                                                                switch ("\n\t" + "hello") {
                                                                                case "":
                                                                                    v118 := GenS14{F0: "\n\t"}
                                                                                    _ = v118
                                                                                }
                                                                                var v119 int32 = (v5 / ((v5) | 1))
                                                                                var v120 float32 = 3.14159
                                                                                _ = v119
                                                                                _ = v120
                                                                            }
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            var v121 uint32 = (v12 / ((v12) | 1))
                                                                            _ = v121
                                                                        case "":
                                                                            v13 += (v13 | v13)
                                                                        default:
                                                                            if len(v2) > 0 {
                                                                                v2[0] = (v4 * v4)
                                                                            }
                                                                            v1 -= v1
                                                                            v12 += (v12 >> (((v7 / ((v7) | 1))) & 63))
                                                                        }
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            fmt.Println(v13)
                                                                            v62 = (v62 + (v62 / ((uint16(v61)) | 1)))
                                                                            v122 := GenS10{F0: (v1 / (((v1 + (-(v1)))) | 1))}
                                                                            _ = v122
                                                                        case "hello":
                                                                            if len(v2) > 0 {
                                                                                v2[0] = fi14
                                                                            }
                                                                            switch (fi14 + fi76) {
                                                                            case 42:
                                                                                _ = len(v2)
                                                                            case -333:
                                                                                if ((572 < (v13 - (^(v13)))) && true) {
                                                                                    if true {
                                                                                        if ((-(v6)) <= (^(v6))) {
                                                                                            fmt.Println(v62, v5)
                                                                                            var v123 string = "\n\t"
                                                                                            _ = v123
                                                                                        }
                                                                                        v61 -= v3
                                                                                    } else {
                                                                                        v124 := GenS5{F0: v7, F1: -516, F2: -1}
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = 128
                                                                                        }
                                                                                        var v125 int32 = v5
                                                                                        _ = v124
                                                                                        _ = v125
                                                                                    }
                                                                                    v7 = (^(v7))
                                                                                }
                                                                            case -491:
                                                                                fmt.Println(v81, v81, v4)
                                                                            default:
                                                                                v126 := GenS12{F0: (v6 * v6), F1: (v62 * uint16(v61))}
                                                                                _ = v126
                                                                            }
                                                                        default:
                                                                            switch 9223372036854775807 {
                                                                            case 818:
                                                                                if (true || (!false)) {
                                                                                    v3 = (^(v3))
                                                                                } else {
                                                                                    if true {
                                                                                        fmt.Println(v62)
                                                                                        switch v4 {
                                                                                        case 42:
                                                                                            v1 -= -1
                                                                                        }
                                                                                        var v127 int64 = -318
                                                                                        _ = v127
                                                                                    }
                                                                                    var v128 int8 = 127
                                                                                    _ = v128
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        switch "a" {
                                                                        case "a":
                                                                            fmt.Println(v62)
                                                                        case "":
                                                                            fmt.Println(v81, v12, v3)
                                                                        case "hello":
                                                                            v62 -= v62
                                                                            switch (^(fi76)) {
                                                                            case 0:
                                                                                var v129 uint16 = 42
                                                                                _ = v129
                                                                            case 153:
                                                                                _ = len(v2)
                                                                            }
                                                                        default:
                                                                            v61 += uint8(v7)
                                                                        }
                                                                        v62 = uint16(v12)
                                                                    }
                                                                    switch int(v6) {
                                                                    case -65:
                                                                        var v130 int16 = -733
                                                                        _ = v130
                                                                    case 127:
                                                                        fmt.Println(v3, v81, v1)
                                                                    }
                                                                    v131 := GenS15{F0: int64(v6), F1: ("\n\t" + "foo bar"), F2: "foo bar"}
                                                                    if ((v13 ^ v13) < (v13 & 0)) {
                                                                        v132 := GenS2{F0: 1, F1: (^(v7)), F2: 184}
                                                                        if false {
                                                                            v133 := GenS5{F0: (^(v7)), F1: (fi76 - (fi14 & v4)), F2: 639}
                                                                            var v134 uint32 = (v12 | 681)
                                                                            _ = v133
                                                                            _ = v134
                                                                        }
                                                                        v135 := GenS8{F0: 0, F1: -2147483648, F2: (false || true)}
                                                                        _ = v132
                                                                        _ = v135
                                                                    }
                                                                    if len(v2) > 0 {
                                                                        v2[0] = fi76
                                                                    }
                                                                    v136 := GenS6{F0: -9223372036854775808, F1: fi14}
                                                                    _ = v131
                                                                    _ = v136
                                                                } else {
                                                                    v7 = v7
                                                                    v1 = (v1 ^ int16(v62))
                                                                    v81 = float64(0.0)
                                                                }
                                                                var v137 uint64 = 554
                                                                _ = v137
                                                            } else {
                                                                v4 += fi76
                                                            }
                                                        default:
                                                            var v138 int16 = (v1 << ((v7) & 63))
                                                            _ = v138
                                                        }
                                                    }
                                                    _ = v81
                                                    _ = v82
                                                } else {
                                                    v1 = -32768
                                                    var v139 int64 = -9223372036854775808
                                                    v2 = append(v2, (fi14 * (v4 - fi76)))
                                                    if (v3 < (v61 - v61)) {
                                                        var v140 string = "hello"
                                                        v5 -= v5
                                                        v141 := GenS2{F0: (v139 % ((int64(v6)) | 1)), F1: (v7 << ((uint(v12)) & 63)), F2: (^(v3))}
                                                        fmt.Println(v140, v5)
                                                        _ = v140
                                                        _ = v141
                                                    }
                                                    var v142 int16 = -32768
                                                    _ = v139
                                                    _ = v142
                                                }
                                            default:
                                                v143 := GenS14{F0: ("hello" + "a")}
                                                _ = v143
                                            }
                                            v2 = append(v2, 0)
                                            _ = v77
                                        }
                                    case 9223372036854775807:
                                        switch fi76 {
                                        case 4294967295:
                                            _ = len(v2)
                                        case 879:
                                            fmt.Println(fi76, fi14)
                                        case -9223372036854775808:
                                            v144 := GenS2{F0: 9223372036854775807, F1: v7, F2: (v61 & v3)}
                                            _ = v144
                                        }
                                    case 364:
                                        _ = len(v2)
                                    }
                                    v145 := GenS4{F0: (fi76 / (((fi14 * fi14)) | 1)), F1: 42}
                                    _ = fi76
                                    _ = v145
                                }
                                var v146 bool = false
                                if len(v2) > 0 {
                                    v2[0] = v4
                                }
                                _ = v146
                            } else {
                                v13 += uint64(v7)
                                if false {
                                    if false {
                                        v2 = append(v2, 128)
                                        v61 = (v3 << (((v7 + (^(v7)))) & 63))
                                        v13 = 2
                                    } else {
                                        v2 = append(v2, (fi14 << ((uint(v62)) & 63)))
                                        v5 -= (v5 % (((-(v5))) | 1))
                                        v147 := GenS16{F0: "\n\t"}
                                        _ = v147
                                    }
                                }
                            }
                        }
                    }
                    var v148 int16 = 1000
                    _ = v148
                    _ = v62
                }
                _ = v61
            }
        }
        _ = fi14
    }
    v2 = append(v2, int(v12))
    if ((-(v1)) != (^(v1))) {
        var v149 bool = true
        _ = v149
    }
    if true {
        if (-168 > -110) {
            if len(v2) > 0 {
                v2[0] = (-(v4))
            }
        }
    } else {
        v4 = v4
        for fi150 := 0; fi150 < 5; fi150++ {
            var v151 int16 = 766
            _ = fi150
            _ = v151
        }
    }
    v3 = v3
    if false {
        v152 := GenS16{F0: "foo bar"}
        _ = v152
    } else {
        for fi153 := 0; fi153 < 7; fi153++ {
            var v154 uint16 = 255
            var v155 int64 = 533
            var v156 uint8 = (^(v3))
            _ = fi153
            _ = v154
            _ = v155
            _ = v156
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
