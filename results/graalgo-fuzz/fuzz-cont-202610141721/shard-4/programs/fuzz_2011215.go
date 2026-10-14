package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int16 }
type GenS2 struct { F0 int32; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int32; F2 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 uint; F2 string }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string; F1 uint8; F2 uint16 }
type GenS9 struct { F0 int32; F1 bool; F2 int }
type GenS10 struct { F0 uint; F1 bool; F2 int }
type GenS11 struct { F0 uint8; F1 int8; F2 int }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint8; F1 int64; F2 int64 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint16; F1 int16; F2 bool }
type GenS16 struct { F0 int; F1 uint16; F2 uint }
type GenS17 struct { F0 int32; F1 uint16 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 uint64 }
func genFn0(p0 int64, p1 int64) bool {
    for fi4 := 0; fi4 < 17; fi4++ {
        for fi5 := 0; fi5 < 14; fi5++ {
            switch "foo bar" {
            case "foo bar":
                v6 := GenS1{F0: 65535, F1: "", F2: -846}
                _ = v6
            case "":
                switch fi4 {
                case 32767:
                    v7 := []int{(fi5 + 127), (^(fi5)), 32767, (^(fi4))}
                    _ = v7
                }
                p0 = (p0 & (p0 - (p0 & p1)))
                p1 = p0
            default:
                if false {
                    var v8 uint8 = 225
                    if true {
                        if (41 != v8) {
                            p1 = p0
                            if ((-128 != -1) || (false || false)) {
                                p1 -= 0
                                if (21 >= 2) {
                                    v8 -= 0
                                }
                            }
                            v8 = 2
                            if ((2 > 0) || true) {
                                v9 := GenS2{F0: 968, F1: (p1 - (p1 - 9223372036854775807)), F2: 32767}
                                if (255 < 65535) {
                                    v10 := GenS3{F0: (v8 + v8), F1: -541, F2: "a"}
                                    v11 := GenS1{F0: 0, F1: "a", F2: 271}
                                    fmt.Println(p0)
                                    _ = v10
                                    _ = v11
                                } else {
                                    var v12 int = fi5
                                    v13 := []int{v12, fi5}
                                    if len(v13) > 0 {
                                        v13[0] = (fi5 ^ fi5)
                                    }
                                    _ = v12
                                    _ = v13
                                }
                                p1 += p0
                                _ = v9
                            }
                            if true {
                                p0 += (-(p0))
                            }
                            if false {
                                v8 += 0
                                if true {
                                    if (p1 < (p0 % (((p0 ^ (-(p1)))) | 1))) {
                                        switch (fi5 & 1) {
                                        case 32767:
                                            v8 -= 1
                                        case 999:
                                            v14 := GenS4{F0: 2147483647}
                                            _ = v14
                                        case 361:
                                            if (447 != 94) {
                                                v8 -= v8
                                            } else {
                                                v15 := GenS5{F0: (fi4 >> ((307) & 63)), F1: 18446744073709551615, F2: ""}
                                                _ = v15
                                            }
                                        default:
                                            var v16 uint16 = 2
                                            _ = v16
                                        }
                                    } else {
                                        v8 = 127
                                        var v17 uint16 = 2
                                        _ = v17
                                    }
                                    p1 = -1
                                    v18 := GenS4{F0: 0}
                                    if (-1 <= 295) {
                                        var v19 uint8 = 0
                                        switch fi4 {
                                        case -77:
                                            switch 127 {
                                            case 0:
                                                var v20 uint32 = 319
                                                _ = v20
                                            case -344:
                                                fmt.Println(p1)
                                            }
                                        case -394:
                                            p0 = 0
                                        case 4294967295:
                                            v19 = (v19 & v8)
                                        }
                                        var v21 string = "\n\t"
                                        _ = v19
                                        _ = v21
                                    } else {
                                        v22 := GenS1{F0: 65535, F1: "foo bar", F2: 32767}
                                        v8 -= 0
                                        _ = v22
                                    }
                                    _ = v18
                                }
                                fmt.Println(p0, fi5, p1)
                            } else {
                                var v23 int = fi4
                                v23 += (fi5 / ((int(v8)) | 1))
                                v24 := GenS6{F0: 0, F1: 252}
                                _ = v23
                                _ = v24
                            }
                        } else {
                            switch (("" + "\n\t") + "a") {
                            case "foo bar":
                                fmt.Println(p0, v8, p1)
                            case "":
                                var v25 string = "\n\t"
                                _ = v25
                            }
                        }
                        var v26 float64 = 1.5
                        _ = v26
                    }
                    var v27 int16 = -32768
                    _ = v27
                    _ = v8
                } else {
                    p0 -= -782
                    if false {
                        v28 := GenS3{F0: 0, F1: 956, F2: (((((((((("a" + (("" + "a") + "a")) + "\n\t") + "hello") + "hello") + "\n\t") + "foo bar") + "a") + "a") + (("a" + ("hello" + "a")) + (("\n\t" + "hello") + ("hello" + "")))) + ((("foo bar" + "a") + "hello") + "foo bar"))}
                        p0 = (-(p1))
                        p1 = p1
                        v29 := GenS4{F0: 1}
                        _ = v28
                        _ = v29
                    } else {
                        if (1 < 65535) {
                            p0 -= p0
                            switch "foo bar" {
                            case "\n\t":
                                var v30 float32 = 0.1
                                _ = v30
                            }
                            p1 = -600
                        } else {
                            var v31 int16 = 185
                            if (fi4 > (-(fi4))) {
                                switch "a" {
                                case "":
                                    p1 = p1
                                case "foo bar":
                                    var v32 uint32 = 65535
                                    _ = v32
                                case "hello":
                                    if (4294967295 <= 2) {
                                        v33 := []int{fi5}
                                        v31 += v31
                                        var v34 float32 = 1.5
                                        _ = v33
                                        _ = v34
                                    }
                                }
                                fmt.Println(fi4)
                            } else {
                                switch ("" + "foo bar") {
                                case "\n\t":
                                    var v35 uint32 = 451
                                    _ = v35
                                case "a":
                                    var v36 uint64 = 569
                                    _ = v36
                                }
                                if false {
                                    v37 := GenS2{F0: 2147483647, F1: (p1 & 9223372036854775807), F2: 0}
                                    if (!true) {
                                        if false {
                                            p0 += (p0 << ((906) & 63))
                                            switch "" {
                                            case "":
                                                var v38 int = (fi5 % (((fi4 * 128)) | 1))
                                                _ = v38
                                            case "hello":
                                                if true {
                                                    var v39 float32 = 0.0
                                                    _ = v39
                                                } else {
                                                    v40 := []int{}
                                                    _ = v40
                                                }
                                            case "\n\t":
                                                var v41 uint16 = 733
                                                _ = v41
                                            default:
                                                p0 -= (p1 + -238)
                                            }
                                            v42 := []int{(fi5 * 128), (fi4 ^ -413), 32768, fi4}
                                            v43 := GenS7{F0: -8}
                                            p1 += p1
                                            _ = v42
                                            _ = v43
                                        } else {
                                            switch "a" {
                                            case "foo bar":
                                                var v44 uint16 = 1
                                                _ = v44
                                            default:
                                                var v45 int16 = v31
                                                _ = v45
                                            }
                                            switch "hello" {
                                            case "\n\t":
                                                var v46 uint16 = 2
                                                _ = v46
                                            default:
                                                v47 := []int{(fi5 << ((711) & 63)), 1}
                                                _ = v47
                                            }
                                        }
                                        var v48 int8 = 0
                                        v49 := GenS8{F0: ("" + "hello"), F1: 0, F2: 42}
                                        _ = v48
                                        _ = v49
                                    } else {
                                        v50 := []int{}
                                        if false {
                                            var v51 uint = 255
                                            _ = v51
                                        }
                                        var v52 uint64 = 4294967295
                                        _ = v50
                                        _ = v52
                                    }
                                    var v53 int8 = 100
                                    _ = v37
                                    _ = v53
                                }
                            }
                            v54 := GenS9{F0: 1, F1: false, F2: 948}
                            _ = v31
                            _ = v54
                        }
                        if (0.1 < (-2.25 / (((1.5 / float64(1e10)) - float64((0.1 / (0.1 / 1.5)))) * 3.14159))) {
                            if (!true) {
                                v55 := []int{int(p0), 521}
                                if false {
                                    var v56 uint32 = 1
                                    switch fi5 {
                                    case 32768:
                                        switch (-(fi4)) {
                                        case 127:
                                            fmt.Println(p1, fi5, p0)
                                            if true {
                                                p0 = (p1 * (p1 << ((846) & 63)))
                                                v57 := GenS4{F0: -1}
                                                _ = v57
                                            }
                                            fmt.Println(fi5)
                                        case -94:
                                            if false {
                                                if true {
                                                    p1 += 119
                                                }
                                                var v58 int64 = (^(p0))
                                                if (true || false) {
                                                    if (-1 < 32767) {
                                                        if false {
                                                            v59 := GenS9{F0: -1, F1: true, F2: fi5}
                                                            v60 := GenS2{F0: 119, F1: p0, F2: 255}
                                                            var v61 bool = (0 <= -1)
                                                            _ = v59
                                                            _ = v60
                                                            _ = v61
                                                        } else {
                                                            p0 += -1
                                                        }
                                                        if (!(!true)) {
                                                            switch 32767 {
                                                            case 676:
                                                                p1 = (p1 * (p1 - p1))
                                                            case 4294967295:
                                                                v56 = (v56 / ((968) | 1))
                                                            case 805:
                                                                switch ("" + "a") {
                                                                case "a":
                                                                    v62 := GenS10{F0: 765, F1: false, F2: fi4}
                                                                    _ = v62
                                                                case "":
                                                                    switch fi4 {
                                                                    case 128:
                                                                        if (1 >= 34) {
                                                                            var v63 string = ("foo bar" + "\n\t")
                                                                            var v64 int32 = -2147483648
                                                                            _ = v63
                                                                            _ = v64
                                                                        }
                                                                        switch "" {
                                                                        case "a":
                                                                            p1 = (p1 & p1)
                                                                        default:
                                                                            switch "hello" {
                                                                            case "a":
                                                                                switch -870 {
                                                                                case 869:
                                                                                    var v65 uint8 = 64
                                                                                    _ = v65
                                                                                }
                                                                            case "hello":
                                                                                if false {
                                                                                    fmt.Println(fi4, p0, fi4)
                                                                                    p0 += p1
                                                                                } else {
                                                                                    switch "" {
                                                                                    case "a":
                                                                                        v66 := GenS3{F0: uint8(v56), F1: 0, F2: ("hello" + "\n\t")}
                                                                                        _ = v66
                                                                                    default:
                                                                                        v55 = append(v55, (fi4 + 32768))
                                                                                        if (32767 == 32767) {
                                                                                            fmt.Println(v56)
                                                                                        } else {
                                                                                            v67 := GenS3{F0: 255, F1: 2147483647, F2: (("\n\t" + "") + "\n\t")}
                                                                                            var v68 int8 = 116
                                                                                            switch "" {
                                                                                            case "a":
                                                                                                var v69 int = 32767
                                                                                                fmt.Println(p1)
                                                                                                var v70 uint32 = v56
                                                                                                fmt.Println(v58)
                                                                                                _ = v69
                                                                                                _ = v70
                                                                                            case "\n\t":
                                                                                                v71 := GenS11{F0: 2, F1: (v68 >> ((531) & 63)), F2: fi5}
                                                                                                _ = v71
                                                                                            default:
                                                                                                fmt.Println(fi5)
                                                                                            }
                                                                                            _ = v67
                                                                                            _ = v68
                                                                                        }
                                                                                    }
                                                                                    v72 := GenS12{F0: -1}
                                                                                    _ = v72
                                                                                }
                                                                            case "":
                                                                                fmt.Println(v56, v56, p0)
                                                                            default:
                                                                                if len(v55) >= 2 {
                                                                                    v55 = v55[1:len(v55)]
                                                                                }
                                                                            }
                                                                        }
                                                                    case 2147483647:
                                                                        v73 := GenS7{F0: 42}
                                                                        _ = v73
                                                                    default:
                                                                        var v74 int16 = 42
                                                                        _ = v74
                                                                    }
                                                                case "hello":
                                                                    v75 := GenS13{F0: 127, F1: 80, F2: 0}
                                                                    _ = len(v55)
                                                                    _ = v75
                                                                }
                                                            }
                                                        }
                                                        if false {
                                                            v58 = (-(p0))
                                                        }
                                                    } else {
                                                        _ = len(v55)
                                                        v76 := GenS1{F0: v56, F1: "a", F2: -1}
                                                        v77 := GenS14{F0: 148}
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                } else {
                                                    if (true || (false || true)) {
                                                        if true {
                                                            if true {
                                                                if len(v55) >= 2 {
                                                                    v55 = v55[1:len(v55)]
                                                                }
                                                                if (1e10 < 3.14159) {
                                                                    var v78 int16 = 581
                                                                    fmt.Println(v78)
                                                                    var v79 uint64 = 0
                                                                    v80 := GenS9{F0: 557, F1: false, F2: fi5}
                                                                    _ = v78
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                            }
                                                            fmt.Println(p0, v56, p0)
                                                        } else {
                                                            v55 = append(v55, (fi5 ^ fi5))
                                                            p1 += p1
                                                        }
                                                        switch (fi4 >> ((4294967295) & 63)) {
                                                        case 782:
                                                            v58 -= (p1 * 9223372036854775807)
                                                        case -1:
                                                            var v81 uint = 1
                                                            _ = v81
                                                        case -601:
                                                            p1 += -9223372036854775808
                                                        default:
                                                            fmt.Println(p1)
                                                        }
                                                    }
                                                }
                                                if (0 != -2147483648) {
                                                    switch ("foo bar" + (("\n\t" + "\n\t") + "foo bar")) {
                                                    case "\n\t":
                                                        if false {
                                                            var v82 float32 = 0.1
                                                            v83 := GenS4{F0: 324}
                                                            _ = v82
                                                            _ = v83
                                                        } else {
                                                            if true {
                                                                v56 += v56
                                                            }
                                                        }
                                                    case "foo bar":
                                                        var v84 int32 = 1
                                                        _ = v84
                                                    case "a":
                                                        v85 := GenS11{F0: 42, F1: -120, F2: (fi5 | fi5)}
                                                        _ = v85
                                                    }
                                                } else {
                                                    var v86 int16 = 0
                                                    v56 -= 65535
                                                    var v87 uint8 = 127
                                                    _ = v86
                                                    _ = v87
                                                }
                                                _ = v58
                                            } else {
                                                if (-1 >= 0) {
                                                    var v88 int16 = 1
                                                    var v89 uint = 73
                                                    fmt.Println(fi4)
                                                    _ = v88
                                                    _ = v89
                                                }
                                            }
                                        case -2147483648:
                                            switch "\n\t" {
                                            case "\n\t":
                                                p0 = p0
                                            default:
                                                v56 -= v56
                                            }
                                        default:
                                            fmt.Println(p1, fi5, fi4)
                                        }
                                    }
                                    if len(v55) >= 2 {
                                        v55 = v55[1:len(v55)]
                                    }
                                    _ = v56
                                }
                                var v90 string = "a"
                                _ = v55
                                _ = v90
                            }
                        }
                    }
                    fmt.Println(p0, fi4)
                }
            }
            var v91 float32 = (0.1 / float32(((-2.25 / 3.14159) * 0.1)))
            _ = fi5
            _ = v91
        }
        p1 = 885
        if ("foo bar" <= ("foo bar" + ("\n\t" + "foo bar"))) {
            p1 = p0
            var v92 uint = 255
            _ = v92
        }
        _ = fi4
    }
    v93 := GenS15{F0: 255, F1: -32768, F2: ((((0.1 - -2.25) / (float32(-2.25) / 0.1)) / (0.1 + 0.0)) < 3.14159)}
    for fi94 := 0; fi94 < 2; fi94++ {
        var v95 int = (fi94 & fi94)
        if false {
            for fi96 := 0; fi96 < 16; fi96++ {
                if (!(false || (32767 < (fi96 % (((fi96 << ((181) & 63))) | 1))))) {
                    var v97 string = "a"
                    _ = v97
                }
                var v98 uint = 255
                v99 := GenS16{F0: (fi96 % (((fi94 << ((v98) & 63))) | 1)), F1: 1, F2: v98}
                switch (fi94 - int(p1)) {
                case 128:
                    if true {
                        var v100 uint16 = 588
                        v101 := GenS5{F0: fi94, F1: 4294967295, F2: ""}
                        _ = v100
                        _ = v101
                    } else {
                        if (!(3.14159 > (0.1 / ((0.0 - -2.25) - (0.1 / 1.0))))) {
                            v98 += (^(v98))
                            var v102 bool = true
                            var v103 bool = v102
                            _ = v102
                            _ = v103
                        } else {
                            if (1 <= 0) {
                                if false {
                                    var v104 string = "\n\t"
                                    fmt.Println(fi96, v98, p0)
                                    _ = v104
                                } else {
                                    var v105 int8 = 119
                                    var v106 string = ("\n\t" + "hello")
                                    switch (fi96 ^ fi94) {
                                    case -1:
                                        v107 := []int{}
                                        _ = v107
                                    case 182:
                                        v108 := []int{}
                                        _ = v108
                                    case 2147483647:
                                        v109 := GenS12{F0: (-(v105))}
                                        _ = v109
                                    default:
                                        var v110 uint64 = 4294967295
                                        _ = v110
                                    }
                                    _ = v105
                                    _ = v106
                                }
                                switch fi96 {
                                case 1:
                                    v111 := GenS17{F0: -55, F1: 877}
                                    _ = v111
                                case 2147483647:
                                    var v112 bool = true
                                    _ = v112
                                case 127:
                                    v98 += v98
                                }
                            }
                            switch (("hello" + "\n\t") + "hello") {
                            case "a":
                                v113 := GenS5{F0: (fi96 / ((1) | 1)), F1: 595, F2: "\n\t"}
                                _ = v113
                            case "":
                                fmt.Println(fi96, v98, v98)
                            case "foo bar":
                                p0 = (p1 & 619)
                            }
                            v114 := []int{}
                            switch ("\n\t" + "a") {
                            case "":
                                var v115 int8 = -90
                                if ((false || (!false)) && (-2.25 < (3.14159 / ((float32(float64(1.5)) - (((((1e10 + 0.0) - (3.14159 / -2.25)) + 1e10) / (0.1 / 3.14159)) - float32(-2.25))) - 3.14159)))) {
                                    v114 = append(v114, (v95 ^ int(v115)))
                                    var v116 int64 = int64(fi96)
                                    if (false && true) {
                                        v95 = (fi96 + fi96)
                                        if len(v114) > 0 {
                                            v114[0] = fi94
                                        }
                                    } else {
                                        v117 := GenS2{F0: -1, F1: v116, F2: 535}
                                        _ = v117
                                    }
                                    v95 = (fi94 << ((v98) & 63))
                                    _ = v116
                                } else {
                                    var v118 int8 = v115
                                    var v119 int16 = -124
                                    _ = v118
                                    _ = v119
                                }
                                var v120 uint16 = 836
                                _ = v115
                                _ = v120
                            case "foo bar":
                                fmt.Println(fi94, v95)
                            default:
                                if true {
                                    var v121 uint8 = 255
                                    var v122 uint64 = 2
                                    _ = v121
                                    _ = v122
                                }
                            }
                            _ = v114
                        }
                        p0 = p1
                    }
                case -1:
                    if ((0 >= 398) || false) {
                        fmt.Println(p1)
                        if (!true) {
                            if (!(-128 == 100)) {
                                switch (fi96 >> ((v98) & 63)) {
                                case -9223372036854775808:
                                    v98 -= v98
                                default:
                                    v95 = fi94
                                }
                            }
                        } else {
                            var v123 int8 = -1
                            _ = v123
                        }
                    }
                default:
                    switch "" {
                    case "":
                        switch (-(fi96)) {
                        case 4294967295:
                            v124 := GenS13{F0: 71, F1: -703, F2: (p0 * int64(v98))}
                            _ = v124
                        }
                    default:
                        var v125 uint64 = 18446744073709551615
                        _ = v125
                    }
                }
                _ = fi96
                _ = v98
                _ = v99
            }
            if true {
                var v126 uint16 = 65535
                for fi127 := 0; fi127 < 6; fi127++ {
                    p1 = (p0 >> ((747) & 63))
                    if true {
                        var v128 float32 = float32(((-2.25 - (3.14159 / 1.0)) + -2.25))
                        _ = v128
                    }
                    _ = fi127
                }
                switch "foo bar" {
                case "":
                    if (v126 != (v126 + (v126 / (((v126 / ((v126) | 1))) | 1)))) {
                        for fi129 := 0; fi129 < 9; fi129++ {
                            var v130 int64 = p1
                            v131 := GenS14{F0: 0}
                            fmt.Println(v126, p1, fi129)
                            _ = fi129
                            _ = v130
                            _ = v131
                        }
                    } else {
                        v126 -= 734
                    }
                default:
                    for fi132 := 0; fi132 < 6; fi132++ {
                        if (18446744073709551615 < 70) {
                            var v133 uint32 = 544
                            v134 := []int{fi132, -2147483648, 127, (-(fi94))}
                            _ = v133
                            _ = v134
                        } else {
                            v95 += -759
                            var v135 int16 = 326
                            if (100 == -114) {
                                if true {
                                    fmt.Println(p1, v95)
                                    var v136 float32 = 1.5
                                    v137 := GenS1{F0: 667, F1: ("a" + ("" + "")), F2: 32767}
                                    _ = v136
                                    _ = v137
                                }
                            }
                            _ = v135
                        }
                        var v138 int8 = 127
                        var v139 float64 = 1e10
                        fmt.Println(v126)
                        _ = fi132
                        _ = v138
                        _ = v139
                    }
                }
                _ = v126
            } else {
                var v140 uint = 255
                _ = v140
            }
        }
        _ = fi94
        _ = v95
    }
    v141 := GenS8{F0: "hello", F1: 127, F2: 1}
    v142 := GenS18{F0: true}
    _ = p0
    _ = p1
    _ = v141
    _ = v142
    _ = v93
    return false
}
func genFn1() int16 {
    v143 := []int{-9223372036854775808, -2147483648, 128, -210}
    var v144 uint64 = 0
    fmt.Println(v144)
    if false {
        var v145 int64 = 1
        for fi146 := 0; fi146 < 15; fi146++ {
            fmt.Println(fi146, v144)
            var v147 bool = (695 < 0)
            _ = fi146
            _ = v147
        }
        _ = v145
    }
    _ = v143
    _ = v144
    return -32768
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{("" + "foo bar")}
    v3 := make(map[string]uint32)
    switch "" {
    case "a":
        var v148 uint = 459
        _ = v148
    default:
        switch "" {
        case "":
            for fi149 := 0; fi149 < 20; fi149++ {
                fmt.Println(fi149)
                var v150 uint16 = 244
                v150 = v150
                _ = fi149
                _ = v150
            }
        case "a":
            switch 128 {
            case 0:
                var v151 uint32 = 4294967295
                _ = v151
            }
        }
    }
    var v152 uint32 = 706
    if false {
        var v153 uint = 999
        var v154 uint8 = 128
        v155 := GenS19{F0: 42}
        _ = v153
        _ = v154
        _ = v155
    }
    fmt.Println(v152)
    for fi156 := 0; fi156 < 19; fi156++ {
        v157 := GenS7{F0: -98}
        _ = fi156
        _ = v157
    }
    var v158 bool = false
    v159 := GenS7{F0: -25}
    for fi160 := 0; fi160 < 16; fi160++ {
        v161 := GenS6{F0: 0, F1: 0}
        switch "" {
        case "hello":
            v158 = true
        case "a":
            var v162 int = fi160
            _ = v162
        }
        _ = fi160
        _ = v161
    }
    switch 451 {
    case -131:
        var v163 uint64 = uint64(v152)
        _ = v163
    case -9223372036854775808:
        var v164 float32 = float32(((1e10 * 1.5) + ((((((1.5 + float64((((-2.25 - ((1e10 * (3.14159 / (((1.5 * (1.5 / 3.14159)) - 3.14159) + 0.0))) + 3.14159)) + 0.0) / 1e10))) + (3.14159 + float64(-2.25))) + 0.1) + 0.1) + 1e10) - 1e10)))
        _ = v164
    }
    var v165 uint = 65535
    v158 = v158
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v152:", v152)
    fmt.Println("v158:", v158)
    fmt.Println("v165:", v165)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v152
    _ = v158
    _ = v159
    _ = v165
    _ = v2
    _ = v3
}
