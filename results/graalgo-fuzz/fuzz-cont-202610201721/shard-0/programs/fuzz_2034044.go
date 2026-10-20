package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 uint8; F1 int16; F2 int32 }
type GenS7 struct { F0 uint32; F1 int; F2 uint64 }
type GenS8 struct { F0 int64; F1 int; F2 uint32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int16; F1 uint32; F2 int64 }
type GenS11 struct { F0 string; F1 uint8 }
type GenS12 struct { F0 int64; F1 int8 }
type GenS13 struct { F0 string; F1 int16; F2 int }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint16; F1 string }
type GenS18 struct { F0 string }
type GenS19 struct { F0 int16; F1 int16; F2 uint8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 1
    var v3 uint64 = (v2 << ((215) & 63))
    v4 := []int{42, 19, -1}
    var v5 int16 = 1000
    v6 := []uint32{42, 2147483648}
    v5 = (-(v5))
    v7 := GenS1{F0: 169}
    for fi8 := 0; fi8 < 4; fi8++ {
        _ = len(v1)
        v9 := GenS1{F0: (fi8 | fi8)}
        switch "" {
        case "\n\t":
            var v10 bool = ((true && true) || false)
            _ = v10
        case "foo bar":
            if ((100 < -111) || (65535 != 65535)) {
                fmt.Println(v2, v5, v3)
                var v11 uint = 2
                for fi12 := 0; fi12 < 16; fi12++ {
                    if ((fi8 != 9223372036854775807) && true) {
                        if false {
                            v13 := GenS1{F0: (^(fi12))}
                            _ = v13
                        }
                        if ((0 != 95) && true) {
                            fmt.Println(fi8)
                            v14 := GenS1{F0: (fi12 ^ 9223372036854775807)}
                            _ = v14
                        }
                        var v15 float64 = 3.14159
                        fmt.Println(fi8, fi8)
                        _ = v15
                    } else {
                        v5 = (v5 | 1)
                        var v16 uint8 = 127
                        delete(v1, "foo bar")
                        _ = v16
                    }
                    if (-867 <= 42) {
                        if ("\n\t" != "") {
                            if ((!true) && false) {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                        } else {
                            if true {
                                var v17 int64 = 42
                                _ = v17
                            } else {
                                if (-115 <= -26) {
                                    if false {
                                        var v18 uint64 = v2
                                        var v19 uint8 = 255
                                        _ = v18
                                        _ = v19
                                    } else {
                                        v5 = (^(v5))
                                        switch (fi8 >> ((v11) & 63)) {
                                        case 9223372036854775807:
                                            if len(v4) > 0 {
                                                v4[0] = fi12
                                            }
                                        case 0:
                                            var v20 int = 42
                                            _ = v20
                                        case -9223372036854775808:
                                            v3 -= (v3 ^ 42)
                                        }
                                        v3 -= (v3 | v2)
                                    }
                                }
                                if (!true) {
                                    var v21 int32 = -2147483648
                                    fmt.Println(v5, v11, v5)
                                    v1["a"] = ("hello" + "hello")
                                    _ = v21
                                }
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            fmt.Println(v5)
                            if false {
                                v11 -= v11
                            } else {
                                switch fi12 {
                                case 594:
                                    v2 += 2
                                case -976:
                                    var v22 int = fi12
                                    _ = v22
                                case -418:
                                    delete(v1, "a")
                                }
                                switch (fi12 << ((v11) & 63)) {
                                case 481:
                                    switch ("a" + "") {
                                    case "foo bar":
                                        fmt.Println(fi12, v5, fi12)
                                    }
                                case 2147483647:
                                    v11 = (v11 << ((v11) & 63))
                                case 33:
                                    var v23 string = (("" + "hello") + "hello")
                                    _ = v23
                                default:
                                    var v24 int64 = -878
                                    _ = v24
                                }
                            }
                            v25 := GenS2{F0: -2147483648}
                            _ = v25
                        }
                        fmt.Println(v3)
                        var v26 bool = (!(true || ((((0.0 + float32(3.14159)) + float32((3.14159 - -2.25))) / 3.14159) > 3.14159)))
                        _ = v26
                    } else {
                        v3 += v2
                        var v27 uint64 = (v3 ^ (v3 << ((1) & 63)))
                        _ = v27
                    }
                    v28 := GenS1{F0: (fi8 / ((fi12) | 1))}
                    _ = fi12
                    _ = v28
                }
                _ = v11
            }
        case "":
            v29 := GenS3{F0: 0, F1: -357}
            _ = v29
        }
        _ = fi8
        _ = v9
    }
    v5 -= v5
    switch "\n\t" {
    case "\n\t":
        var v30 uint16 = 65535
        _ = v30
    case "hello":
        fmt.Println(v3)
    case "":
        v6 = append(v6, 2)
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v3)
    case "":
        v31 := GenS4{F0: v5}
        _ = v31
    case "\n\t":
        var v32 uint16 = 42
        _ = v32
    }
    v3 = (v2 + (v3 & (v2 | (^(v3)))))
    switch 450 {
    case 0:
        fmt.Println(v3, v5, v5)
    case 507:
        switch 786 {
        case 1:
            if false {
                var v33 bool = true
                _ = v33
            }
        case -2147483648:
            for fi34 := 0; fi34 < 18; fi34++ {
                if (3.14159 >= 3.14159) {
                    for fi35 := 0; fi35 < 9; fi35++ {
                        v2 = (v3 | v2)
                        switch 32767 {
                        case 4294967295:
                            var v36 int32 = 0
                            _ = v36
                        case 271:
                            switch "foo bar" {
                            case "a":
                                if false {
                                    v37 := GenS3{F0: uint(v2), F1: fi35}
                                    _ = v37
                                } else {
                                    switch "" {
                                    case "\n\t":
                                        v5 = (v5 % (((-(v5))) | 1))
                                    case "":
                                        var v38 int32 = 1
                                        _ = v38
                                    default:
                                        fmt.Println(v5)
                                        fmt.Println(v2, v2, fi35)
                                        var v39 int = -172
                                        _ = v39
                                    }
                                }
                            case "hello":
                                var v40 uint32 = 42
                                _ = v40
                            case "":
                                if false {
                                    if true {
                                        var v41 int8 = 42
                                        switch (fi35 + (fi34 + 1)) {
                                        case 1:
                                            if true {
                                                v3 = (^(v2))
                                                v3 -= v2
                                            }
                                        case 32767:
                                            var v42 int16 = 539
                                            _ = v42
                                        default:
                                            if ((-(v41)) == 42) {
                                                v43 := GenS2{F0: -1}
                                                _ = v43
                                            } else {
                                                var v44 int = (fi34 >> ((18446744073709551615) & 63))
                                                _ = v44
                                            }
                                        }
                                        switch fi34 {
                                        case 1:
                                            fmt.Println(fi34, fi35, v2)
                                        case -1:
                                            v45 := GenS5{F0: (v41 >> ((33) & 63)), F1: fi34}
                                            _ = v45
                                        case 4294967295:
                                            v5 = v5
                                        }
                                        _ = v41
                                    }
                                } else {
                                    var v46 bool = ((0 > -9223372036854775808) || true)
                                    if v46 {
                                        v46 = v46
                                        var v47 uint64 = (v2 | v3)
                                        _ = v47
                                    } else {
                                        switch fi34 {
                                        case 32767:
                                            v2 += 259
                                            var v48 uint = 4294967295
                                            _ = v48
                                        case -2147483648:
                                            if (2147483647 < 2147483648) {
                                                var v49 int64 = 425
                                                if len(v4) > 0 {
                                                    v4[0] = (fi35 << ((65535) & 63))
                                                }
                                                _ = len(v6)
                                                _ = v49
                                            }
                                        case 1:
                                            switch ("foo bar" + "foo bar") {
                                            case "a":
                                                var v50 float64 = 0.0
                                                _ = v50
                                            case "foo bar":
                                                v51 := GenS1{F0: fi34}
                                                _ = v51
                                            case "hello":
                                                v3 = (v2 & 911)
                                            default:
                                                switch ("a" + "") {
                                                case "":
                                                    var v52 uint64 = 495
                                                    _ = v52
                                                default:
                                                    var v53 uint = 2
                                                    _ = v53
                                                }
                                                v5 = 0
                                                var v54 int = 1
                                                fmt.Println(v54, fi34)
                                                _ = v54
                                            }
                                        }
                                        if true {
                                            if v46 {
                                                fmt.Println(v5, v3, fi34)
                                                v46 = v46
                                            } else {
                                                if v46 {
                                                    switch (fi35 & fi34) {
                                                    case 344:
                                                        if false {
                                                            var v55 int = fi35
                                                            var v56 float32 = 1e10
                                                            _ = v55
                                                            _ = v56
                                                        } else {
                                                            v3 += v2
                                                            var v57 uint16 = 255
                                                            if true {
                                                                fmt.Println(fi34)
                                                                var v58 int32 = -1
                                                                fmt.Println(v3)
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                _ = v58
                                                            }
                                                            _ = len(v4)
                                                            _ = v57
                                                        }
                                                    case -712:
                                                        if (!false) {
                                                            v59 := v1["\n\t"]
                                                            var v60 int = (fi35 % (((fi34 << ((18446744073709551615) & 63))) | 1))
                                                            v59 = v59
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                    default:
                                                        var v61 int = 987
                                                        var v62 uint16 = 42
                                                        var v63 int8 = 127
                                                        if ((v62 >> ((4294967295) & 63)) == (v62 ^ (v62 << ((0) & 63)))) {
                                                            _ = len(v4)
                                                        }
                                                        _ = v61
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                    if v46 {
                                                        v64 := GenS6{F0: 0, F1: v5, F2: -1}
                                                        var v65 float64 = (3.14159 + 0.0)
                                                        switch "foo bar" {
                                                        case "a":
                                                            var v66 bool = true
                                                            _ = v66
                                                        case "":
                                                            switch "a" {
                                                            case "":
                                                                var v67 uint16 = 32767
                                                                _ = v67
                                                            default:
                                                                v5 -= (v5 - (v5 / (((^(v5))) | 1)))
                                                            }
                                                        case "\n\t":
                                                            switch "a" {
                                                            case "\n\t":
                                                                if true {
                                                                    var v68 float64 = v65
                                                                    var v69 uint64 = 231
                                                                    if (v46 && ("hello" < "")) {
                                                                        if true {
                                                                            if true {
                                                                                v70 := GenS5{F0: 79, F1: (fi35 & -1)}
                                                                                _ = v70
                                                                            }
                                                                            if (!(v5 == v5)) {
                                                                                v71 := GenS1{F0: fi35}
                                                                                _ = v71
                                                                            } else {
                                                                                v68 = 1e10
                                                                                v5 -= -32768
                                                                            }
                                                                        }
                                                                        v5 -= (v5 + (v5 << ((4294967295) & 63)))
                                                                        v5 -= v5
                                                                    } else {
                                                                        v3 += v2
                                                                        v72 := GenS6{F0: 193, F1: 1000, F2: 0}
                                                                        v73 := v1[("foo bar" + ("" + "hello"))]
                                                                        if (((v3 >= (v2 ^ (^(v69)))) && (true && (!v46))) || v46) {
                                                                            v65 = (1e10 - -2.25)
                                                                            if true {
                                                                                var v74 string = v73
                                                                                v75 := GenS1{F0: (fi34 | fi34)}
                                                                                v76 := GenS7{F0: 2147483648, F1: fi35, F2: v2}
                                                                                v77 := GenS2{F0: -904}
                                                                                _ = v74
                                                                                _ = v75
                                                                                _ = v76
                                                                                _ = v77
                                                                            } else {
                                                                                if ((v73 + ((v73 + ("a" + v73)) + "\n\t")) > "") {
                                                                                    if (128 >= int(v3)) {
                                                                                        switch fi34 {
                                                                                        case 4294967295:
                                                                                            v78 := GenS8{F0: -9223372036854775808, F1: 127, F2: 944}
                                                                                            _ = v78
                                                                                        case 9223372036854775807:
                                                                                            var v79 int8 = 0
                                                                                            _ = v79
                                                                                        }
                                                                                        _ = len(v4)
                                                                                        switch (fi34 & fi34) {
                                                                                        case -2147483648:
                                                                                            var v80 int8 = -1
                                                                                            _ = v80
                                                                                        case 127:
                                                                                            switch ("foo bar" + ((v73 + v73) + v73)) {
                                                                                            case "":
                                                                                                var v81 uint8 = 125
                                                                                                var v82 uint = 0
                                                                                                _ = v81
                                                                                                _ = v82
                                                                                            }
                                                                                        case 32768:
                                                                                            var v83 string = v73
                                                                                            _ = v83
                                                                                        default:
                                                                                            switch ("a" + v73) {
                                                                                            case "a":
                                                                                                v2 = 42
                                                                                            case "foo bar":
                                                                                                if v46 {
                                                                                                    if len(v4) > 0 {
                                                                                                        v4[0] = (fi35 | fi34)
                                                                                                    }
                                                                                                    v5 -= (-(v5))
                                                                                                }
                                                                                            case "hello":
                                                                                                v46 = false
                                                                                            }
                                                                                            v84 := GenS8{F0: -9223372036854775808, F1: fi34, F2: 1}
                                                                                            v85 := GenS5{F0: -128, F1: fi35}
                                                                                            _ = v84
                                                                                            _ = v85
                                                                                        }
                                                                                    } else {
                                                                                        v5 = v5
                                                                                    }
                                                                                    if v46 {
                                                                                        var v86 uint8 = 0
                                                                                        _ = v86
                                                                                    }
                                                                                    var v87 string = v73
                                                                                    fmt.Println(v69)
                                                                                    switch fi35 {
                                                                                    case 54:
                                                                                        v65 = (v65 + v65)
                                                                                    case 32767:
                                                                                        if (fi35 <= 42) {
                                                                                            var v88 float64 = 1e10
                                                                                            v89 := GenS9{F0: (fi35 & -657)}
                                                                                            _ = len(v6)
                                                                                            _ = v88
                                                                                            _ = v89
                                                                                        }
                                                                                    default:
                                                                                        if v46 {
                                                                                            _ = len(v1)
                                                                                        } else {
                                                                                            var v90 float32 = -2.25
                                                                                            v87 = v73
                                                                                            if v46 {
                                                                                                var v91 bool = false
                                                                                                _ = v91
                                                                                            } else {
                                                                                                fmt.Println(v5, v5)
                                                                                                v46 = true
                                                                                                v92 := v1[(((v73 + v87) + v87) + v87)]
                                                                                                _ = v92
                                                                                            }
                                                                                            _ = v90
                                                                                        }
                                                                                        switch v87 {
                                                                                        case "\n\t":
                                                                                            v73 = v87
                                                                                        case "hello":
                                                                                            v6 = append(v6, 2)
                                                                                        }
                                                                                    }
                                                                                    var v93 uint = 18446744073709551615
                                                                                    _ = v87
                                                                                    _ = v93
                                                                                }
                                                                                var v94 uint = 292
                                                                                switch (fi35 % ((fi34) | 1)) {
                                                                                case 32767:
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                default:
                                                                                    v46 = (!v46)
                                                                                }
                                                                                _ = v94
                                                                            }
                                                                            fmt.Println(v73, v73, v3)
                                                                        }
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                    _ = v68
                                                                    _ = v69
                                                                }
                                                                v46 = v46
                                                                var v95 uint64 = (v2 & 927)
                                                                v46 = v46
                                                                _ = len(v6)
                                                                fmt.Println(v46, v2)
                                                                v96 := GenS8{F0: -1, F1: 32768, F2: 2}
                                                                _ = v95
                                                                _ = v96
                                                            case "foo bar":
                                                                var v97 int16 = v5
                                                                _ = v97
                                                            case "":
                                                                delete(v1, "hello")
                                                            }
                                                        default:
                                                            var v98 uint = 524
                                                            _ = v98
                                                        }
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    var v99 int = fi35
                                                    _ = v99
                                                } else {
                                                    var v100 int = 579
                                                    _ = v100
                                                }
                                                if v46 {
                                                    switch ("\n\t" + "foo bar") {
                                                    case "foo bar":
                                                        var v101 uint64 = (v2 >> ((42) & 63))
                                                        _ = v101
                                                    case "\n\t":
                                                        v102 := GenS6{F0: 42, F1: 1, F2: 2147483647}
                                                        _ = v102
                                                    default:
                                                        v3 = (v3 << ((586) & 63))
                                                    }
                                                }
                                            }
                                            switch fi34 {
                                            case 127:
                                                v3 = 18446744073709551615
                                            case 1:
                                                if (true && true) {
                                                    v103 := GenS5{F0: 42, F1: 32767}
                                                    var v104 int32 = 0
                                                    var v105 uint32 = 65535
                                                    var v106 float32 = -2.25
                                                    var v107 string = (("" + "hello") + ("" + "\n\t"))
                                                    v2 -= (v3 * 691)
                                                    v105 -= (v105 + (v105 + v105))
                                                    v108 := GenS8{F0: -9223372036854775808, F1: (fi34 & (fi35 % (((fi34 * 9223372036854775807)) | 1))), F2: v105}
                                                    _ = v103
                                                    _ = v104
                                                    _ = v105
                                                    _ = v106
                                                    _ = v107
                                                    _ = v108
                                                }
                                            case 223:
                                                var v109 float32 = 1e10
                                                _ = v109
                                            }
                                        } else {
                                            switch (("foo bar" + "\n\t") + ((("hello" + "hello") + ("foo bar" + ("hello" + "\n\t"))) + "hello")) {
                                            case "a":
                                                v110 := GenS10{F0: v5, F1: 1, F2: 42}
                                                _ = v110
                                            default:
                                                var v111 float32 = 0.0
                                                _ = v111
                                            }
                                            var v112 uint16 = 1
                                            _ = v112
                                        }
                                    }
                                    _ = v46
                                }
                            default:
                                v3 -= v2
                            }
                        default:
                            fmt.Println(v2)
                        }
                        fmt.Println(fi34, fi35)
                        _ = fi35
                    }
                    var v113 string = ("a" + "foo bar")
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v113
                } else {
                    var v114 string = ("hello" + (("foo bar" + "a") + "foo bar"))
                    v115 := GenS10{F0: (v5 & v5), F1: 0, F2: -565}
                    if true {
                        _ = len(v4)
                        switch (fi34 * (fi34 - fi34)) {
                        case 374:
                            for fi116 := 0; fi116 < 12; fi116++ {
                                var v117 string = v114
                                _ = fi116
                                _ = v117
                            }
                        case 32767:
                            v3 = v3
                        case 127:
                            var v118 int64 = -1
                            _ = v118
                        default:
                            if (65535 > 0) {
                                v5 += (-(v5))
                                switch v114 {
                                case "\n\t":
                                    v5 = v5
                                }
                            }
                        }
                        if true {
                            v119 := GenS6{F0: 255, F1: v5, F2: -490}
                            if (!(!(157 > 32767))) {
                                if (false || true) {
                                    var v120 int8 = -121
                                    _ = v120
                                }
                                v3 += v3
                            } else {
                                var v121 uint32 = 0
                                _ = v121
                            }
                            if ((227 == -1) || ((fi34 >> ((4294967295) & 63)) != (fi34 << ((255) & 63)))) {
                                var v122 int16 = (v5 + 32767)
                                for fi123 := 0; fi123 < 20; fi123++ {
                                    var v124 float64 = 0.1
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    v122 -= (v122 >> ((2) & 63))
                                    if ((v124 * (0.0 * -2.25)) <= v124) {
                                        var v125 float32 = -2.25
                                        _ = v125
                                    }
                                    fmt.Println(fi123, v124)
                                    v126 := GenS7{F0: 884, F1: fi123, F2: v3}
                                    _ = fi123
                                    _ = v124
                                    _ = v126
                                }
                                v3 = (v3 >> ((2) & 63))
                                _ = v122
                            }
                            _ = v119
                        } else {
                            var v127 int64 = 71
                            v5 = (v5 * 277)
                            switch "" {
                            case "":
                                v128 := GenS11{F0: (("foo bar" + "foo bar") + "\n\t"), F1: 2}
                                _ = v128
                            case "hello":
                                for fi129 := 0; fi129 < 11; fi129++ {
                                    switch (fi129 * (-(fi129))) {
                                    case 127:
                                        delete(v1, (v114 + (v114 + v114)))
                                    case 1:
                                        v127 -= v127
                                    case -895:
                                        fmt.Println(v127)
                                    default:
                                        fmt.Println(fi129, v2)
                                    }
                                    _ = fi129
                                }
                            case "\n\t":
                                v130 := GenS12{F0: v127, F1: -128}
                                _ = v130
                            default:
                                for fi131 := 0; fi131 < 10; fi131++ {
                                    if false {
                                        v132 := GenS9{F0: fi131}
                                        _ = v132
                                    }
                                    var v133 float32 = 0.1
                                    _ = fi131
                                    _ = v133
                                }
                            }
                            _ = v127
                        }
                    }
                    _ = v114
                    _ = v115
                }
                var v134 bool = true
                v3 -= (v2 & 0)
                _ = fi34
                _ = v134
            }
        case 127:
            if true {
                if (v5 == (v5 / ((v5) | 1))) {
                    for fi135 := 0; fi135 < 1; fi135++ {
                        for fi136 := 0; fi136 < 19; fi136++ {
                            v2 -= v3
                            _ = fi136
                        }
                        v3 = 42
                        _ = fi135
                    }
                    for fi137 := 0; fi137 < 6; fi137++ {
                        if false {
                            v138 := GenS1{F0: (fi137 ^ 1)}
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v3 = (v2 / ((v3) | 1))
                            _ = v138
                        }
                        fmt.Println(fi137, fi137, v3)
                        if true {
                            if true {
                                v139 := GenS13{F0: "a", F1: (v5 - (^(v5))), F2: (fi137 + fi137)}
                                _ = v139
                            } else {
                                v1[""] = (((("" + "hello") + "\n\t") + "foo bar") + ("a" + ("\n\t" + "")))
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v140 := GenS4{F0: v5}
                                _ = v140
                            }
                        } else {
                            switch fi137 {
                            case -418:
                                if false {
                                    if (false && (!(false || false))) {
                                        v141 := GenS14{F0: fi137}
                                        for fi142 := 0; fi142 < 6; fi142++ {
                                            delete(v1, "foo bar")
                                            fmt.Println(v5)
                                            var v143 float32 = float32(0.1)
                                            switch 32767 {
                                            case -795:
                                                fmt.Println(v143)
                                                if true {
                                                    v1["hello"] = ""
                                                } else {
                                                    switch "foo bar" {
                                                    case "a":
                                                        switch (^(fi137)) {
                                                        case 32768:
                                                            switch fi137 {
                                                            case 127:
                                                                v5 = v5
                                                            }
                                                        case 1:
                                                            v144 := GenS15{F0: 9223372036854775807}
                                                            _ = v144
                                                        }
                                                    }
                                                    v145 := GenS12{F0: -100, F1: -1}
                                                    switch int(v2) {
                                                    case -351:
                                                        fmt.Println(fi137, v3, v143)
                                                    case 242:
                                                        var v146 int16 = v5
                                                        _ = v146
                                                    case 42:
                                                        v1["a"] = "a"
                                                    default:
                                                        if false {
                                                            switch "\n\t" {
                                                            case "a":
                                                                switch (fi142 & fi137) {
                                                                case -658:
                                                                    var v147 float32 = v143
                                                                    _ = v147
                                                                case 127:
                                                                    var v148 uint = 18446744073709551615
                                                                    _ = v148
                                                                case -1:
                                                                    v2 += v3
                                                                }
                                                            case "":
                                                                fmt.Println(v143, fi137, fi137)
                                                            case "hello":
                                                                fmt.Println(v2, fi137)
                                                            }
                                                        } else {
                                                            var v149 int32 = 1
                                                            switch "foo bar" {
                                                            case "":
                                                                switch "foo bar" {
                                                                case "a":
                                                                    switch 42 {
                                                                    case 32767:
                                                                        v150 := GenS16{F0: (fi137 ^ 0)}
                                                                        _ = v150
                                                                    case -611:
                                                                        var v151 int32 = (v149 << ((4294967295) & 63))
                                                                        v143 = float32(1e10)
                                                                        if false {
                                                                            var v152 float32 = float32((-2.25 - 1.5))
                                                                            if true {
                                                                                switch "hello" {
                                                                                case "\n\t":
                                                                                    var v153 uint64 = (^(v2))
                                                                                    _ = v153
                                                                                }
                                                                            } else {
                                                                                v154 := GenS1{F0: 32768}
                                                                                _ = v154
                                                                            }
                                                                            _ = v152
                                                                        } else {
                                                                            v3 = (v2 + (v2 % (((v3 + v2)) | 1)))
                                                                        }
                                                                        _ = v151
                                                                    default:
                                                                        _ = len(v4)
                                                                    }
                                                                case "\n\t":
                                                                    v155 := v1[""]
                                                                    _ = v155
                                                                }
                                                            default:
                                                                var v156 uint8 = 128
                                                                _ = v156
                                                            }
                                                            _ = v149
                                                        }
                                                    }
                                                    v157 := GenS17{F0: 1, F1: "foo bar"}
                                                    _ = v145
                                                    _ = v157
                                                }
                                                v158 := v1[("\n\t" + ("" + "hello"))]
                                                v159 := v1[v158]
                                                _ = v158
                                                _ = v159
                                            }
                                            var v160 int64 = 42
                                            _ = fi142
                                            _ = v143
                                            _ = v160
                                        }
                                        fmt.Println(fi137)
                                        _ = v141
                                    } else {
                                        var v161 float64 = float64(1e10)
                                        _ = v161
                                    }
                                    v3 += v3
                                    v1[""] = ((("hello" + "a") + ("a" + "\n\t")) + "a")
                                } else {
                                    var v162 uint16 = 32767
                                    _ = v162
                                }
                            case 811:
                                var v163 string = "\n\t"
                                _ = v163
                            case 515:
                                for fi164 := 0; fi164 < 17; fi164++ {
                                    var v165 int64 = 42
                                    _ = fi164
                                    _ = v165
                                }
                            default:
                                if len(v4) > 0 {
                                    v4[0] = fi137
                                }
                            }
                        }
                        _ = fi137
                    }
                }
                if ((v3 == v3) || true) {
                    v166 := GenS18{F0: ""}
                    v5 += (v5 ^ v5)
                    _ = v166
                } else {
                    for fi167 := 0; fi167 < 7; fi167++ {
                        fmt.Println(v3, fi167)
                        var v168 int8 = 0
                        v168 = (-(v168))
                        _ = fi167
                        _ = v168
                    }
                }
            } else {
                if false {
                    switch ("\n\t" + (("hello" + "a") + "foo bar")) {
                    case "hello":
                        v5 = (^(v5))
                    default:
                        if true {
                            var v169 uint32 = 0
                            _ = v169
                        } else {
                            v170 := GenS19{F0: 0, F1: v5, F2: 255}
                            _ = v170
                        }
                    }
                    v5 += (v5 / ((32767) | 1))
                } else {
                    _ = len(v4)
                    var v171 uint64 = (v2 ^ (v3 + (v3 + v2)))
                    var v172 int = int(v5)
                    _ = v171
                    _ = v172
                }
                _ = len(v1)
            }
        }
    default:
        for fi173 := 0; fi173 < 16; fi173++ {
            fmt.Println(v2, v2)
            var v174 float32 = 0.1
            _ = fi173
            _ = v174
        }
    }
    if (549 <= 2147483648) {
        for fi175 := 0; fi175 < 15; fi175++ {
            v2 = (v3 * v2)
            v3 += v2
            _ = fi175
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
