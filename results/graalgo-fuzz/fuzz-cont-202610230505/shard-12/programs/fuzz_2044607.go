package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int16 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 bool; F2 int64 }
type GenS7 struct { F0 int16; F1 string; F2 int }
type GenS8 struct { F0 int8; F1 uint }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int32; F1 uint16 }
type GenS14 struct { F0 int; F1 uint16; F2 uint64 }
type GenS15 struct { F0 int8; F1 string; F2 uint32 }
type GenS16 struct { F0 uint16; F1 string }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 int32; F1 int32; F2 int32 }
type GenS21 struct { F0 uint32; F1 int32; F2 int }
type GenS22 struct { F0 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = false
    var v3 float32 = -2.25
    v4 := []uint32{2147483648, 42, 856}
    var v5 float32 = (v3 + (v3 / v3))
    var v6 uint8 = 128
    if (95 != 127) {
        if v2 {
            v3 -= v3
        } else {
            switch "foo bar" {
            case "\n\t":
                v7 := GenS1{F0: 42, F1: 0, F2: 538}
                _ = v7
            default:
                for fi8 := 0; fi8 < 2; fi8++ {
                    v9 := v1[(-(fi8))]
                    for fi10 := 0; fi10 < 11; fi10++ {
                        var v11 float64 = (1.5 - float64(float32(0.0)))
                        v6 = (v6 * 128)
                        var v12 string = ""
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = fi8
                    _ = v9
                }
            }
        }
        if v2 {
            if (127 != v6) {
                var v13 int8 = -128
                var v14 float32 = (v5 - v3)
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 5; fi15++ {
                v5 = ((0.0 - (v5 * 3.14159)) - v5)
                v5 -= (v3 - v5)
                if true {
                    for fi16 := 0; fi16 < 18; fi16++ {
                        switch "a" {
                        case "foo bar":
                            switch (-(fi16)) {
                            case 1:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case 980:
                                if (v6 > (^(v6))) {
                                    v5 = 1e10
                                    if len(v4) > 0 {
                                        v4[0] = 2
                                    }
                                    switch fi15 {
                                    case 803:
                                        v6 = v6
                                    default:
                                        if (v2 && (127 == (^(fi15)))) {
                                            delete(v1, 127)
                                            v17 := GenS1{F0: -1, F1: 339, F2: 1}
                                            var v18 int16 = -649
                                            _ = v17
                                            _ = v18
                                        }
                                    }
                                    v5 = v5
                                }
                            case 32767:
                                var v19 int64 = -118
                                _ = v19
                            default:
                                fmt.Println(fi16, v5, v3)
                                var v20 int8 = -53
                                if true {
                                    fmt.Println(v6, v3)
                                }
                                var v21 bool = (1 != 2)
                                _ = v20
                                _ = v21
                            }
                        case "hello":
                            if (1 != 1) {
                                v22 := GenS2{F0: 114, F1: 150}
                                fmt.Println(v2, v5)
                                v23 := GenS3{F0: 255, F1: 0}
                                _ = v22
                                _ = v23
                            }
                        default:
                            var v24 uint64 = 2
                            _ = v24
                        }
                        v4 = append(v4, 2147483647)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi16
                    }
                    switch "foo bar" {
                    case "hello":
                        var v25 string = ("hello" + "")
                        _ = v25
                    case "\n\t":
                        for fi26 := 0; fi26 < 8; fi26++ {
                            if (12 <= 0) {
                                v5 = (-2.25 + (((((v3 / (v3 * 1.5)) * v5) - v3) - float32(3.14159)) / (v3 - (float32(3.14159) / (v3 - -2.25)))))
                            }
                            var v27 int64 = 1
                            _ = fi26
                            _ = v27
                        }
                    }
                }
                _ = fi15
            }
        } else {
            var v28 int8 = 92
            v29 := GenS4{F0: 0}
            _ = v28
            _ = v29
        }
        v2 = ((v6 | 208) >= v6)
    } else {
        var v30 uint32 = 674
        _ = v30
    }
    if v2 {
        for fi31 := 0; fi31 < 4; fi31++ {
            if v2 {
                if false {
                    var v32 float32 = float32(1e10)
                    for fi33 := 0; fi33 < 6; fi33++ {
                        if v2 {
                            fmt.Println(fi33, v5, v2)
                            switch "hello" {
                            case "\n\t":
                                var v34 int64 = 0
                                _ = v34
                            default:
                                v5 = v3
                            }
                        }
                        v35 := GenS1{F0: 0, F1: 255, F2: -840}
                        _ = fi33
                        _ = v35
                    }
                    if true {
                        v36 := GenS5{F0: 1}
                        v4 = append(v4, uint32(v6))
                        _ = v36
                    } else {
                        var v37 uint16 = 255
                        switch "" {
                        case "":
                            v37 += 1
                        }
                        _ = v37
                    }
                    _ = v32
                } else {
                    fmt.Println(v6, v2)
                }
                if true {
                    for fi38 := 0; fi38 < 18; fi38++ {
                        var v39 int8 = -128
                        var v40 int = (fi31 ^ fi38)
                        fmt.Println(fi31, v6)
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                } else {
                    v3 = v3
                    v5 = float32(1e10)
                    fmt.Println(v2, v3)
                }
            } else {
                var v41 uint64 = 976
                if v2 {
                    v41 = v41
                    if (fi31 <= fi31) {
                        v41 = (v41 & (v41 * v41))
                    }
                    v41 -= v41
                }
                var v42 bool = v2
                _ = v41
                _ = v42
            }
            v3 -= v3
            v2 = (true && v2)
            _ = fi31
        }
        if true {
            for fi43 := 0; fi43 < 10; fi43++ {
                v6 = v6
                var v44 uint64 = 1
                var v45 string = "hello"
                _ = fi43
                _ = v44
                _ = v45
            }
            v46 := GenS3{F0: 32767, F1: 333}
            var v47 float64 = 3.14159
            _ = v46
            _ = v47
        } else {
            v4 = append(v4, 2147483647)
            if v2 {
                if (-722 == 0) {
                    v6 += (v6 & (v6 + 64))
                }
                v48 := GenS6{F0: 1000, F1: (!v2), F2: -9223372036854775808}
                var v49 int = -1
                _ = v48
                _ = v49
            } else {
                if false {
                    switch "" {
                    case "foo bar":
                        if false {
                            var v50 int = -601
                            if (9223372036854775807 == 1) {
                                var v51 int16 = 32767
                                if (v2 || (-84 < 127)) {
                                    fmt.Println(v6, v3)
                                    v5 += 0.1
                                } else {
                                    v4 = append(v4, 119)
                                    for fi52 := 0; fi52 < 11; fi52++ {
                                        v50 = v50
                                        fmt.Println(fi52, fi52, v3)
                                        if (127 < v6) {
                                            fmt.Println(v51, fi52)
                                            for fi53 := 0; fi53 < 14; fi53++ {
                                                if (584 <= 382) {
                                                    v50 = 127
                                                    var v54 int16 = v51
                                                    _ = v54
                                                } else {
                                                    v6 += (v6 - v6)
                                                }
                                                v1[-271] = v2
                                                _ = fi53
                                            }
                                            var v55 uint16 = 402
                                            _ = v55
                                        }
                                        _ = fi52
                                    }
                                }
                                _ = v51
                            } else {
                                v50 -= v50
                                switch "" {
                                case "\n\t":
                                    var v56 uint64 = 18446744073709551615
                                    _ = v56
                                case "hello":
                                    fmt.Println(v3, v50)
                                case "foo bar":
                                    for fi57 := 0; fi57 < 2; fi57++ {
                                        var v58 int = (fi57 ^ fi57)
                                        v5 = (v5 + (v3 + v5))
                                        delete(v1, (fi57 ^ 824))
                                        _ = fi57
                                        _ = v58
                                    }
                                default:
                                    switch v50 {
                                    case 9223372036854775807:
                                        if (18446744073709551615 < 0) {
                                            for fi59 := 0; fi59 < 20; fi59++ {
                                                fmt.Println(v2, fi59)
                                                var v60 bool = (((((float64(1e10) - 1.5) + float64(float32(-2.25))) / 3.14159) / -2.25) < 3.14159)
                                                v61 := GenS4{F0: 1}
                                                _ = fi59
                                                _ = v60
                                                _ = v61
                                            }
                                            for fi62 := 0; fi62 < 13; fi62++ {
                                                v2 = (65535 > 506)
                                                if v2 {
                                                    v63 := GenS3{F0: 32767, F1: 355}
                                                    if v2 {
                                                        for fi64 := 0; fi64 < 17; fi64++ {
                                                            var v65 float64 = (1.5 / 0.1)
                                                            _ = fi64
                                                            _ = v65
                                                        }
                                                        _ = len(v4)
                                                    } else {
                                                        v66 := GenS1{F0: -32768, F1: 2, F2: 881}
                                                        _ = v66
                                                    }
                                                    _ = v63
                                                }
                                                var v67 int8 = 42
                                                _ = fi62
                                                _ = v67
                                            }
                                            v1[(v50 * (v50 / (((v50 | 981)) | 1)))] = true
                                        } else {
                                            for fi68 := 0; fi68 < 17; fi68++ {
                                                for fi69 := 0; fi69 < 6; fi69++ {
                                                    var v70 uint16 = 42
                                                    switch -1 {
                                                    case -926:
                                                        var v71 uint8 = v6
                                                        var v72 int = (v50 & (^(fi68)))
                                                        if ((1e10 - v3) == (((1e10 - v3) / v5) + v3)) {
                                                            var v73 float64 = (3.14159 / 3.14159)
                                                            _ = v73
                                                        } else {
                                                            v6 = v71
                                                            if (!v2) {
                                                                if false {
                                                                    var v74 float64 = 3.14159
                                                                    v72 += v50
                                                                    _ = v74
                                                                } else {
                                                                    _ = len(v1)
                                                                    if (!false) {
                                                                        v75 := v1[2147483647]
                                                                        if len(v4) > 0 {
                                                                            v4[0] = 4294967295
                                                                        }
                                                                        if false {
                                                                            var v76 int8 = 0
                                                                            var v77 uint8 = (v6 / ((v6) | 1))
                                                                            fmt.Println(v71, v70)
                                                                            _ = v76
                                                                            _ = v77
                                                                        } else {
                                                                            switch (^(v72)) {
                                                                            case -794:
                                                                                v4 = append(v4, 1)
                                                                            case 260:
                                                                                v2 = true
                                                                            }
                                                                            var v78 float64 = 3.14159
                                                                            v5 -= (v3 / -2.25)
                                                                            if v75 {
                                                                                if (true || v2) {
                                                                                    v71 = 1
                                                                                    v50 -= fi68
                                                                                } else {
                                                                                    v79 := GenS7{F0: 42, F1: ("foo bar" + ("a" + "a")), F2: v72}
                                                                                    _ = v79
                                                                                }
                                                                            }
                                                                            _ = v78
                                                                        }
                                                                        _ = v75
                                                                    } else {
                                                                        var v80 bool = true
                                                                        _ = v80
                                                                    }
                                                                    v81 := GenS4{F0: 42}
                                                                    _ = v81
                                                                }
                                                                v82 := GenS4{F0: -1}
                                                                var v83 int64 = 192
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                        }
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                    v6 = v6
                                                    _ = fi69
                                                    _ = v70
                                                }
                                                _ = fi68
                                            }
                                            var v84 int = v50
                                            v84 = (^(v84))
                                            _ = v84
                                        }
                                    case 32768:
                                        switch v50 {
                                        case -33:
                                            var v85 string = "\n\t"
                                            _ = v85
                                        case 1:
                                            if (71 == 255) {
                                                v50 -= v50
                                            } else {
                                                v6 = v6
                                                if v2 {
                                                    switch 9223372036854775807 {
                                                    case 128:
                                                        switch 128 {
                                                        case 335:
                                                            v86 := GenS8{F0: -1, F1: 2}
                                                            _ = v86
                                                        }
                                                    case 1:
                                                        v87 := GenS5{F0: 65535}
                                                        _ = v87
                                                    }
                                                    switch "foo bar" {
                                                    case "a":
                                                        fmt.Println(v3, v2)
                                                    case "":
                                                        for fi88 := 0; fi88 < 1; fi88++ {
                                                            v89 := GenS6{F0: -394, F1: true, F2: 9223372036854775807}
                                                            if ((v3 > v5) || v2) {
                                                                v3 = ((0.0 - v5) + (((float32(1e10) * v3) * (v3 / (float32(0.1) / (0.1 * float32(1e10))))) * (1.5 + (float32(3.14159) - v3))))
                                                                v5 -= 0.1
                                                            } else {
                                                                v3 = v5
                                                                v90 := GenS9{F0: 0}
                                                                _ = v90
                                                            }
                                                            _ = fi88
                                                            _ = v89
                                                        }
                                                    }
                                                    if (255 > 825) {
                                                        fmt.Println(v6)
                                                        v2 = false
                                                        v6 = 128
                                                    }
                                                }
                                            }
                                        }
                                    case 127:
                                        switch "foo bar" {
                                        case "foo bar":
                                            v91 := GenS5{F0: 2}
                                            _ = v91
                                        case "hello":
                                            var v92 int = v50
                                            _ = v92
                                        }
                                    }
                                }
                                if (!true) {
                                    v93 := GenS5{F0: 932}
                                    if (!v2) {
                                        for fi94 := 0; fi94 < 15; fi94++ {
                                            for fi95 := 0; fi95 < 9; fi95++ {
                                                if true {
                                                    var v96 int8 = -1
                                                    _ = v96
                                                } else {
                                                    v6 += (^(v6))
                                                    var v97 int8 = 127
                                                    _ = v97
                                                }
                                                _ = fi95
                                            }
                                            _ = fi94
                                        }
                                        if (127 < 75) {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = len(v1)
                                        }
                                        v50 = (v50 ^ 2147483647)
                                    }
                                    v98 := v1[-558]
                                    _ = v93
                                    _ = v98
                                } else {
                                    var v99 string = "\n\t"
                                    _ = v99
                                }
                            }
                            for fi100 := 0; fi100 < 19; fi100++ {
                                if ("a" > (("hello" + "") + "\n\t")) {
                                    v101 := GenS10{F0: 2147483647}
                                    _ = v101
                                }
                                v2 = false
                                v6 += v6
                                _ = fi100
                            }
                            _ = v50
                        }
                    case "a":
                        for fi102 := 0; fi102 < 8; fi102++ {
                            if v2 {
                                if v2 {
                                    var v103 bool = v2
                                    if v103 {
                                        v104 := v1[fi102]
                                        if ("\n\t" != "a") {
                                            if false {
                                                switch 9223372036854775807 {
                                                case 127:
                                                    var v105 float64 = 1.5
                                                    _ = v105
                                                case -247:
                                                    v106 := GenS11{F0: 32767}
                                                    _ = v106
                                                case 960:
                                                    switch "a" {
                                                    case "foo bar":
                                                        v103 = ((!((v3 >= 0.0) || false)) && (474 != 255))
                                                    case "hello":
                                                        switch (fi102 * fi102) {
                                                        case -2147483648:
                                                            _ = len(v4)
                                                        case -1:
                                                            v107 := GenS10{F0: -236}
                                                            _ = v107
                                                        default:
                                                            switch ("\n\t" + "") {
                                                            case "\n\t":
                                                                switch "a" {
                                                                case "":
                                                                    for fi108 := 0; fi108 < 10; fi108++ {
                                                                        v109 := GenS6{F0: -1, F1: v2, F2: 9223372036854775807}
                                                                        _ = fi108
                                                                        _ = v109
                                                                    }
                                                                case "foo bar":
                                                                    if v104 {
                                                                        var v110 uint = 18446744073709551615
                                                                        var v111 float64 = float64(v5)
                                                                        _ = v110
                                                                        _ = v111
                                                                    }
                                                                default:
                                                                    var v112 uint64 = 2
                                                                    _ = v112
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                fmt.Println(v103, v5, v3)
                                                v5 += (1e10 / (1e10 * 3.14159))
                                            } else {
                                                var v113 uint64 = 336
                                                _ = len(v1)
                                                _ = v113
                                            }
                                            if (v103 || v104) {
                                                for fi114 := 0; fi114 < 3; fi114++ {
                                                    if v2 {
                                                        v115 := GenS6{F0: -176, F1: v104, F2: 1}
                                                        v6 = v6
                                                        _ = v115
                                                    }
                                                    v116 := GenS3{F0: 65535, F1: 0}
                                                    switch ("hello" + "hello") {
                                                    case "\n\t":
                                                        switch fi102 {
                                                        case 298:
                                                            switch fi114 {
                                                            case 891:
                                                                v3 += float32(1e10)
                                                            case 1:
                                                                switch 127 {
                                                                case 9223372036854775807:
                                                                    v3 = v3
                                                                }
                                                            default:
                                                                v6 += (^(v6))
                                                            }
                                                        case 42:
                                                            v117 := GenS12{F0: -530}
                                                            _ = v117
                                                        case -2147483648:
                                                            var v118 uint64 = 2
                                                            _ = v118
                                                        default:
                                                            var v119 float32 = (3.14159 - 3.14159)
                                                            _ = v119
                                                        }
                                                    default:
                                                        var v120 string = "a"
                                                        _ = v120
                                                    }
                                                    _ = fi114
                                                    _ = v116
                                                }
                                                var v121 uint64 = 42
                                                switch "hello" {
                                                case "":
                                                    if (!v103) {
                                                        var v122 int8 = 1
                                                        var v123 uint = 825
                                                        _ = v122
                                                        _ = v123
                                                    } else {
                                                        v121 = (v121 & v121)
                                                    }
                                                case "a":
                                                    v124 := GenS9{F0: 1}
                                                    _ = v124
                                                }
                                                _ = v121
                                            }
                                            switch 723 {
                                            case -1:
                                                for fi125 := 0; fi125 < 20; fi125++ {
                                                    v104 = v104
                                                    v126 := GenS13{F0: 42, F1: 2}
                                                    var v127 int32 = 0
                                                    if v103 {
                                                        fmt.Println(v3)
                                                    } else {
                                                        if (!v104) {
                                                            if true {
                                                                v127 = v127
                                                                v1[(-(fi125))] = false
                                                                v3 -= (v5 - v5)
                                                            } else {
                                                                var v128 uint64 = 4294967295
                                                                var v129 uint8 = (v6 | (v6 & (v6 * (^(v6)))))
                                                                fmt.Println(fi102)
                                                                var v130 uint16 = 255
                                                                if (v6 != 0) {
                                                                    if (v130 < v130) {
                                                                        if ((0 != 1) && true) {
                                                                            v131 := GenS14{F0: fi125, F1: v130, F2: (v128 | (^(v128)))}
                                                                            _ = v131
                                                                        }
                                                                    }
                                                                    v132 := GenS15{F0: 127, F1: "a", F2: 905}
                                                                    switch "hello" {
                                                                    case "":
                                                                        switch (^(fi102)) {
                                                                        case 699:
                                                                            v133 := GenS16{F0: (v130 - v130), F1: (("foo bar" + ("a" + "foo bar")) + "\n\t")}
                                                                            _ = v133
                                                                        }
                                                                    }
                                                                    _ = v132
                                                                } else {
                                                                    if (47 >= 0) {
                                                                        var v134 uint16 = v130
                                                                        var v135 uint64 = 1
                                                                        fmt.Println(v103)
                                                                        v136 := GenS12{F0: -1}
                                                                        fmt.Println(fi125)
                                                                        _ = v134
                                                                        _ = v135
                                                                        _ = v136
                                                                    }
                                                                    v129 = (v129 / ((2) | 1))
                                                                }
                                                                _ = v128
                                                                _ = v129
                                                                _ = v130
                                                            }
                                                            if v104 {
                                                                var v137 uint16 = 371
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                var v138 uint = 1
                                                                _ = v137
                                                                _ = v138
                                                            }
                                                        } else {
                                                            if false {
                                                                v104 = v103
                                                                if len(v4) > 0 {
                                                                    v4[0] = 75
                                                                }
                                                            } else {
                                                                var v139 int64 = -186
                                                                var v140 int16 = 32767
                                                                var v141 uint32 = 4294967295
                                                                v142 := GenS3{F0: 861, F1: 9223372036854775807}
                                                                _ = len(v1)
                                                                if v103 {
                                                                    v143 := GenS1{F0: 0, F1: 200, F2: (v140 % ((v140) | 1))}
                                                                    v144 := GenS17{F0: v2}
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        v103 = v2
                                                                    case "\n\t":
                                                                        fmt.Println(v6, v127)
                                                                        if len(v4) > 0 {
                                                                            v4[0] = 42
                                                                        }
                                                                    default:
                                                                        switch ("foo bar" + "") {
                                                                        case "a":
                                                                            fmt.Println(v2, v103, v6)
                                                                        case "\n\t":
                                                                            v145 := GenS9{F0: 923}
                                                                            _ = v145
                                                                        }
                                                                    }
                                                                    _ = v143
                                                                    _ = v144
                                                                }
                                                                _ = v139
                                                                _ = v140
                                                                _ = v141
                                                                _ = v142
                                                            }
                                                        }
                                                        v146 := GenS18{F0: 41}
                                                        _ = v146
                                                    }
                                                    _ = fi125
                                                    _ = v126
                                                    _ = v127
                                                }
                                            }
                                        }
                                        _ = v104
                                    } else {
                                        var v147 bool = v103
                                        var v148 int64 = -1
                                        _ = v147
                                        _ = v148
                                    }
                                    _ = v103
                                } else {
                                    for fi149 := 0; fi149 < 1; fi149++ {
                                        v4 = append(v4, 65535)
                                        if v2 {
                                            v150 := v1[(fi102 | fi149)]
                                            _ = v150
                                        } else {
                                            if (v3 == 1.5) {
                                                v5 -= (1e10 * v5)
                                                var v151 uint = 4294967295
                                                v5 = v5
                                                _ = v151
                                            }
                                            if (((^(v6)) != (v6 << ((1) & 63))) && v2) {
                                                switch "" {
                                                case "\n\t":
                                                    var v152 uint16 = 0
                                                    _ = v152
                                                case "foo bar":
                                                    _ = len(v1)
                                                }
                                                v3 = 0.0
                                                switch fi102 {
                                                case -2147483648:
                                                    v4 = append(v4, 691)
                                                case -144:
                                                    var v153 uint32 = 0
                                                    _ = v153
                                                case -9223372036854775808:
                                                    if v2 {
                                                        v4 = append(v4, 2147483647)
                                                    } else {
                                                        v5 -= (0.0 / ((1.5 * ((3.14159 + v3) - v5)) / 3.14159))
                                                    }
                                                    switch ("\n\t" + ("a" + "hello")) {
                                                    case "hello":
                                                        var v154 uint8 = v6
                                                        _ = v154
                                                    case "\n\t":
                                                        var v155 int = 610
                                                        _ = v155
                                                    default:
                                                        fmt.Println(v6)
                                                    }
                                                }
                                                v156 := GenS3{F0: 42, F1: 42}
                                                if (!false) {
                                                    var v157 int32 = 0
                                                    v3 = v5
                                                    v2 = v2
                                                    _ = v157
                                                } else {
                                                    fmt.Println(v5)
                                                    if (0 > 65535) {
                                                        v158 := GenS19{F0: -57}
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if v2 {
                                                                v3 = float32(-2.25)
                                                                v3 = v3
                                                            }
                                                        case "a":
                                                            fmt.Println(fi102, fi149, v3)
                                                        default:
                                                            fmt.Println(v2, v6, fi102)
                                                        }
                                                        v3 = 0.0
                                                        var v159 uint32 = 2
                                                        _ = v158
                                                        _ = v159
                                                    } else {
                                                        v160 := GenS17{F0: (v6 != (v6 + v6))}
                                                        _ = len(v4)
                                                        v2 = v2
                                                        _ = v160
                                                    }
                                                }
                                                var v161 uint = 4294967295
                                                fmt.Println(v3)
                                                _ = v156
                                                _ = v161
                                            } else {
                                                v6 = 1
                                                if (!v2) {
                                                    v6 = (^(v6))
                                                }
                                            }
                                            var v162 bool = (143 < 65535)
                                            _ = v162
                                        }
                                        _ = fi149
                                    }
                                }
                            } else {
                                v163 := GenS20{F0: -2147483648, F1: -1, F2: 42}
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v163
                            }
                            v4 = append(v4, 1)
                            v164 := GenS21{F0: 2147483647, F1: -694, F2: -1}
                            _ = fi102
                            _ = v164
                        }
                    case "\n\t":
                        var v165 int32 = 0
                        _ = v165
                    }
                }
                if (!(!(true || v2))) {
                    fmt.Println(v3)
                    var v166 uint8 = 42
                    fmt.Println(v166)
                    _ = v166
                }
                v6 -= 178
            }
        }
        fmt.Println(v2)
    }
    if ((32768 < 358) || v2) {
        if (!v2) {
            if v2 {
                var v167 int = -401
                _ = v167
            }
            var v168 uint64 = 331
            _ = v168
        }
        var v169 int = 1
        _ = v169
    } else {
        if (127 == 1) {
            v6 = (v6 >> ((65535) & 63))
        } else {
            for fi170 := 0; fi170 < 4; fi170++ {
                v171 := GenS3{F0: 65535, F1: 9223372036854775807}
                var v172 uint = 925
                _ = len(v1)
                _ = fi170
                _ = v171
                _ = v172
            }
        }
        switch 1 {
        case -354:
            var v173 int32 = -1
            _ = v173
        case 42:
            v2 = (-9223372036854775808 > -329)
        case -9223372036854775808:
            if (32767 > -1) {
                fmt.Println(v6, v6, v2)
                v3 += (0.1 - ((-2.25 / 0.1) / ((1.5 / v5) + (0.0 / v5))))
            } else {
                fmt.Println(v6, v3, v3)
                fmt.Println(v3, v6)
            }
        }
        switch -31 {
        case -338:
            var v174 bool = v2
            _ = v174
        case 1:
            if v2 {
                for fi175 := 0; fi175 < 4; fi175++ {
                    var v176 int16 = 1000
                    _ = fi175
                    _ = v176
                }
            }
        default:
            if (v2 && v2) {
                switch -444 {
                case -1:
                    for fi177 := 0; fi177 < 6; fi177++ {
                        fmt.Println(v6)
                        delete(v1, (fi177 >> ((411) & 63)))
                        _ = fi177
                    }
                case 2147483647:
                    if true {
                        var v178 int8 = 51
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        switch 0 {
                        case 32767:
                            if (!((!((v178 < -57) && v2)) || (true || (false || ((!true) || v2))))) {
                                if v2 {
                                    if v2 {
                                        var v179 float64 = float64(v3)
                                        _ = v179
                                    }
                                    var v180 float64 = 0.0
                                    switch "\n\t" {
                                    case "a":
                                        _ = len(v1)
                                    case "foo bar":
                                        var v181 uint64 = 0
                                        _ = v181
                                    case "\n\t":
                                        var v182 string = "a"
                                        _ = v182
                                    default:
                                        switch "hello" {
                                        case "a":
                                            v183 := GenS22{F0: (v6 * (v6 + (v6 + 128)))}
                                            _ = v183
                                        default:
                                            fmt.Println(v5, v2)
                                        }
                                    }
                                    _ = v180
                                }
                            }
                        case -2147483648:
                            for fi184 := 0; fi184 < 1; fi184++ {
                                var v185 int = (fi184 % (((fi184 * -2147483648)) | 1))
                                v185 += (fi184 % ((2147483647) | 1))
                                fmt.Println(v5, v6, v2)
                                _ = fi184
                                _ = v185
                            }
                        case 737:
                            v6 -= (^(v6))
                        default:
                            var v186 int16 = 0
                            _ = v186
                        }
                        _ = v178
                    } else {
                        for fi187 := 0; fi187 < 3; fi187++ {
                            v6 -= 2
                            var v188 int16 = 1000
                            _ = fi187
                            _ = v188
                        }
                    }
                case 32767:
                    var v189 int32 = -1
                    _ = v189
                default:
                    fmt.Println(v3, v5)
                }
                v3 = -2.25
            } else {
                v190 := GenS7{F0: -555, F1: "", F2: 2147483647}
                v5 = 3.14159
                _ = v190
            }
        }
    }
    fmt.Println(v5, v3)
    v191 := v1[4294967295]
    delete(v1, 32767)
    if (0.0 != (0.0 - (1e10 - -2.25))) {
        v192 := GenS2{F0: 788, F1: 65535}
        _ = v192
    }
    fmt.Println(v6, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v191:", v191)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v191
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
