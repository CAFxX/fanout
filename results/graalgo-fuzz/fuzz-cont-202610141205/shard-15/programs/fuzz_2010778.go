package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 string; F1 bool; F2 int32 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 int64; F2 int32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint; F1 int32; F2 int8 }
type GenS11 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS12 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS13 struct { F0 uint32; F1 uint64 }
type GenS14 struct { F0 int32; F1 uint32 }
type GenS15 struct { F0 bool; F1 uint16 }
type GenS16 struct { F0 uint; F1 int32 }
type GenS17 struct { F0 uint16; F1 int16 }
type GenS18 struct { F0 bool; F1 uint8; F2 int32 }
type GenS19 struct { F0 int8 }

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]int)
    v3 := []int{-2147483648, 32768}
    var v4 int32 = -2147483648
    var v5 int8 = -69
    v6 := make(map[int]bool)
    if true {
        v7 := GenS1{F0: 4294967295}
        _ = v7
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        v1 -= v1
        _ = fi8
    }
    var v9 int8 = (v5 % ((v5) | 1))
    v4 -= -79
    var v10 int16 = 42
    if (181 != 18446744073709551615) {
        switch 32768 {
        case 127:
            fmt.Println(v4, v10)
        }
        switch 42 {
        case -1:
            fmt.Println(v5, v4, v5)
        case -828:
            v3 = append(v3, 601)
        }
        if (!(42 >= (v9 - 67))) {
            v11 := GenS1{F0: -64}
            v10 -= v10
            v5 = (v5 - (v9 & v5))
            _ = v11
        } else {
            var v12 float64 = ((((float64(float32((-2.25 - -2.25))) * 0.0) * v1) / v1) + v1)
            _ = v12
        }
    }
    switch 1 {
    case -508:
        switch ("hello" + (("hello" + (((("a" + "a") + "foo bar") + "foo bar") + ("foo bar" + ("foo bar" + "a")))) + "foo bar")) {
        case "hello":
            v10 = (v10 % ((v10) | 1))
        case "foo bar":
            v10 = (v10 * -71)
        }
    case 127:
        v4 -= (v4 >> ((0) & 63))
    case 336:
        if false {
            switch "\n\t" {
            case "\n\t":
                for fi13 := 0; fi13 < 12; fi13++ {
                    v10 -= (v10 << ((65535) & 63))
                    v14 := GenS1{F0: 900}
                    _ = fi13
                    _ = v14
                }
            case "a":
                fmt.Println(v1)
            default:
                fmt.Println(v1, v9, v5)
            }
            v10 += (v10 | (v10 | 32767))
        }
    }
    switch "foo bar" {
    case "foo bar":
        v5 += 42
    case "":
        fmt.Println(v4)
    default:
        if false {
            v15 := GenS2{F0: 127, F1: 86}
            _ = v15
        }
    }
    v5 = (v9 - (v5 % ((v5) | 1)))
    v16 := v6[32768]
    v17 := v6[507]
    v4 -= v4
    v9 = 1
    if (1000 >= (^(v10))) {
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 uint32 = 1
            if v17 {
                if v16 {
                    for fi20 := 0; fi20 < 4; fi20++ {
                        v10 += v10
                        v5 = (v5 + (v5 % (((v9 % (((v9 | v9)) | 1))) | 1)))
                        if v16 {
                            v21 := GenS1{F0: -504}
                            v22 := GenS1{F0: (fi18 + (fi18 + 32767))}
                            _ = v21
                            _ = v22
                        } else {
                            if v17 {
                                if (!v17) {
                                    var v23 int = fi18
                                    var v24 uint8 = 0
                                    v25 := v6[v23]
                                    if (!(!false)) {
                                        if v16 {
                                            if true {
                                                v24 = (v24 ^ (v24 >> ((1) & 63)))
                                                fmt.Println(fi20)
                                                v3 = append(v3, fi20)
                                            } else {
                                                var v26 uint16 = 308
                                                v9 -= (v9 & (^(v5)))
                                                v27 := GenS1{F0: v23}
                                                _ = v26
                                                _ = v27
                                            }
                                            var v28 uint32 = (^(v19))
                                            if (v28 > v19) {
                                                if len(v3) > 0 {
                                                    v3[0] = v23
                                                }
                                                _ = len(v2)
                                                fmt.Println(v24, v1)
                                            } else {
                                                v25 = v17
                                                if (v25 && v25) {
                                                    v29 := GenS3{F0: (v19 & v28), F1: 255}
                                                    v5 += v9
                                                    _ = len(v3)
                                                    v17 = (!false)
                                                    _ = v29
                                                }
                                                switch (-(v23)) {
                                                case 128:
                                                    v30 := v6[v23]
                                                    _ = v30
                                                case 32768:
                                                    switch "a" {
                                                    case "":
                                                        v17 = true
                                                    default:
                                                        var v31 uint8 = (^(v24))
                                                        _ = v31
                                                    }
                                                case 127:
                                                    v32 := GenS3{F0: v19, F1: 255}
                                                    _ = v32
                                                default:
                                                    var v33 int32 = (v4 % (((v4 | v4)) | 1))
                                                    _ = v33
                                                }
                                            }
                                            var v34 uint32 = (v28 << ((255) & 63))
                                            v2[(^(v19))] = (^(fi20))
                                            _ = v28
                                            _ = v34
                                        }
                                    } else {
                                        var v35 uint32 = v19
                                        if (!((v10 % ((v10) | 1)) == v10)) {
                                            if v16 {
                                                if (v10 < v10) {
                                                    v36 := GenS4{F0: "", F1: v25, F2: (v4 * 42)}
                                                    fmt.Println(v17, v4, v17)
                                                    if v16 {
                                                        var v37 float32 = 0.1
                                                        v5 -= -128
                                                        v38 := GenS3{F0: v35, F1: 0}
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                    _ = v36
                                                }
                                                if v17 {
                                                    fmt.Println(fi18)
                                                    var v39 int8 = v5
                                                    var v40 int64 = 286
                                                    _ = v39
                                                    _ = v40
                                                } else {
                                                    v41 := GenS5{F0: (!((v23 + -1) > fi20)), F1: v24}
                                                    _ = v41
                                                }
                                                v35 -= (^(v35))
                                            } else {
                                                var v42 int16 = 32767
                                                fmt.Println(fi20, fi20, v17)
                                                v43 := GenS6{F0: 65535}
                                                var v44 uint16 = 42
                                                var v45 int = (fi20 * (-(v23)))
                                                switch fi18 {
                                                case 0:
                                                    v10 += (v42 + (v42 >> ((1) & 63)))
                                                case -1:
                                                    var v46 uint8 = (v24 * v24)
                                                    _ = v46
                                                case 278:
                                                    switch "hello" {
                                                    case "foo bar":
                                                        fmt.Println(v44, fi18)
                                                    case "":
                                                        var v47 int32 = (^(v4))
                                                        _ = v47
                                                    }
                                                default:
                                                    v44 += v44
                                                }
                                                _ = v42
                                                _ = v43
                                                _ = v44
                                                _ = v45
                                            }
                                        }
                                        _ = v35
                                    }
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                                fmt.Println(v4)
                            }
                        }
                        _ = fi20
                    }
                    v48 := GenS7{F0: 0}
                    var v49 uint16 = 1
                    _ = v48
                    _ = v49
                }
                if v16 {
                    if v16 {
                        if v17 {
                            fmt.Println(fi18, v4, fi18)
                            var v50 int64 = -1
                            _ = v50
                        } else {
                            for fi51 := 0; fi51 < 7; fi51++ {
                                if true {
                                    switch fi51 {
                                    case 588:
                                        fmt.Println(v17)
                                    case -1:
                                        switch (("hello" + ("a" + (("foo bar" + "hello") + "hello"))) + "") {
                                        case "":
                                            switch "\n\t" {
                                            case "hello":
                                                v17 = v17
                                            case "\n\t":
                                                var v52 int8 = v5
                                                _ = v52
                                            }
                                        case "foo bar":
                                            fmt.Println(v9)
                                        }
                                    case -2147483648:
                                        v53 := GenS8{F0: v4, F1: 1, F2: (v4 + v4)}
                                        _ = v53
                                    default:
                                        var v54 int8 = v5
                                        _ = v54
                                    }
                                    v55 := v6[fi18]
                                    _ = v55
                                }
                                v56 := GenS9{F0: ""}
                                v57 := GenS10{F0: 65535, F1: v4, F2: (v5 / (((v9 + v9)) | 1))}
                                _ = fi51
                                _ = v56
                                _ = v57
                            }
                            if (v16 || (true && (42 != v10))) {
                                v9 += 112
                            } else {
                                var v58 int32 = v4
                                var v59 int = -2147483648
                                _ = len(v3)
                                _ = v58
                                _ = v59
                            }
                            var v60 uint64 = 18
                            _ = v60
                        }
                    } else {
                        var v61 int32 = 264
                        if v17 {
                            if len(v3) > 0 {
                                v3[0] = (fi18 | (-(fi18)))
                            }
                            fmt.Println(v61, fi18)
                            v6[42] = ((0.0 + (1.5 * (float32(float64(float32(v1))) + (1.5 + (0.0 - (((1e10 * 1.5) * 0.1) * (-2.25 / 1e10))))))) <= float32(v1))
                        }
                        v5 -= v5
                        _ = v61
                    }
                } else {
                    if (239 < 18446744073709551615) {
                        var v62 uint64 = 0
                        switch fi18 {
                        case 100:
                            var v63 uint16 = 42
                            _ = v63
                        default:
                            for fi64 := 0; fi64 < 13; fi64++ {
                                var v65 float32 = 1e10
                                var v66 uint = 42
                                v5 -= v9
                                _ = fi64
                                _ = v65
                                _ = v66
                            }
                        }
                        _ = v62
                    } else {
                        v17 = v16
                        switch "hello" {
                        case "":
                            v67 := GenS3{F0: v19, F1: 42}
                            _ = v67
                        }
                    }
                    v68 := GenS9{F0: "hello"}
                    for fi69 := 0; fi69 < 7; fi69++ {
                        v1 = 1e10
                        if v17 {
                            v3 = append(v3, fi18)
                            fmt.Println(v16, v16)
                        }
                        fmt.Println(v1, v19, v19)
                        _ = fi69
                    }
                    _ = v68
                }
                _ = len(v3)
            }
            _ = fi18
            _ = v19
        }
        var v70 uint32 = 2147483648
        v5 -= (v9 >> ((4294967295) & 63))
        _ = v70
    }
    switch "a" {
    case "":
        if v17 {
            for fi71 := 0; fi71 < 17; fi71++ {
                for fi72 := 0; fi72 < 4; fi72++ {
                    var v73 int16 = v10
                    var v74 uint = 214
                    _ = fi72
                    _ = v73
                    _ = v74
                }
                _ = fi71
            }
            v16 = false
        } else {
            for fi75 := 0; fi75 < 2; fi75++ {
                if v16 {
                    v76 := GenS11{F0: v4, F1: 977, F2: 391}
                    _ = v76
                }
                v77 := GenS4{F0: "a", F1: (1 >= 0), F2: 944}
                _ = fi75
                _ = v77
            }
            var v78 int8 = 25
            v16 = false
            _ = v78
        }
    case "a":
        for fi79 := 0; fi79 < 13; fi79++ {
            if true {
                var v80 int32 = 42
                if (!v16) {
                    if v16 {
                        v81 := GenS11{F0: v80, F1: 18446744073709551615, F2: 860}
                        switch (fi79 % ((fi79) | 1)) {
                        case 9223372036854775807:
                            v9 += (v5 << ((18446744073709551615) & 63))
                        case -323:
                            for fi82 := 0; fi82 < 6; fi82++ {
                                fmt.Println(v17, v5)
                                var v83 bool = (true && true)
                                if true {
                                    fmt.Println(fi79, v9)
                                    var v84 int16 = v10
                                    _ = v84
                                }
                                switch -949 {
                                case 700:
                                    v2[42] = (^(fi79))
                                }
                                fmt.Println(v10, fi79)
                                if (v16 && v17) {
                                    switch "\n\t" {
                                    case "":
                                        var v85 int64 = 539
                                        _ = v85
                                    }
                                    v86 := v6[(fi79 - (fi79 - (fi82 + fi82)))]
                                    v87 := GenS7{F0: 0}
                                    var v88 uint16 = 89
                                    _ = v86
                                    _ = v87
                                    _ = v88
                                }
                                v16 = (65535 < 32767)
                                _ = fi82
                                _ = v83
                            }
                        default:
                            switch "\n\t" {
                            case "\n\t":
                                v89 := GenS2{F0: 115, F1: (^(v5))}
                                _ = v89
                            }
                        }
                        v80 += (^(v4))
                        _ = v81
                    }
                    v90 := v2[660]
                    var v91 uint32 = 4294967295
                    _ = v90
                    _ = v91
                } else {
                    for fi92 := 0; fi92 < 13; fi92++ {
                        if (fi92 >= (fi79 >> ((108) & 63))) {
                            v17 = (0 <= 510)
                            v93 := GenS12{F0: 2, F1: 255, F2: 42}
                            _ = v93
                        }
                        v80 -= (^(v80))
                        v2[2] = (fi79 + fi92)
                        if true {
                            if (!(!(1000 < (^(v10))))) {
                                switch (fi79 ^ fi92) {
                                case 9223372036854775807:
                                    v94 := GenS1{F0: 127}
                                    _ = v94
                                default:
                                    v16 = (0 != 42)
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                v5 = (v9 - (v5 + v9))
                                var v95 int = (fi92 >> ((1) & 63))
                                var v96 int8 = v9
                                _ = v95
                                _ = v96
                            }
                        } else {
                            v97 := v2[2147483647]
                            var v98 int32 = v4
                            _ = v97
                            _ = v98
                        }
                        _ = fi92
                    }
                    if v17 {
                        for fi99 := 0; fi99 < 19; fi99++ {
                            v9 = -1
                            v80 = (v80 >> ((18446744073709551615) & 63))
                            _ = fi99
                        }
                        var v100 int = (fi79 >> ((18446744073709551615) & 63))
                        v101 := GenS2{F0: 128, F1: v5}
                        _ = v100
                        _ = v101
                    } else {
                        var v102 int8 = 27
                        switch "a" {
                        case "a":
                            v5 = v102
                        case "hello":
                            if v17 {
                                delete(v2, 2)
                                if v16 {
                                    v4 = (v80 | v80)
                                } else {
                                    for fi103 := 0; fi103 < 7; fi103++ {
                                        if len(v3) > 0 {
                                            v3[0] = fi103
                                        }
                                        var v104 int8 = (v9 % ((-91) | 1))
                                        _ = fi103
                                        _ = v104
                                    }
                                }
                            }
                        }
                        _ = v102
                    }
                    for fi105 := 0; fi105 < 11; fi105++ {
                        var v106 float64 = v1
                        switch "" {
                        case "foo bar":
                            v1 += v1
                        case "a":
                            var v107 uint16 = 255
                            _ = v107
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        if v17 {
                            v5 -= 1
                            v4 += v80
                            switch fi79 {
                            case -2147483648:
                                delete(v6, (fi105 << ((474) & 63)))
                            case 2147483647:
                                if v17 {
                                    var v108 int64 = -347
                                    v3 = append(v3, fi105)
                                    _ = v108
                                }
                            case 9223372036854775807:
                                switch "a" {
                                case "hello":
                                    fmt.Println(v80, v106)
                                case "":
                                    fmt.Println(v5, v10)
                                default:
                                    switch ("" + (("foo bar" + "hello") + "")) {
                                    case "":
                                        var v109 uint64 = 1
                                        _ = v109
                                    }
                                }
                            default:
                                if v16 {
                                    v1 = v1
                                    if (!true) {
                                        fmt.Println(v5, fi79)
                                    }
                                }
                            }
                        }
                        _ = fi105
                        _ = v106
                    }
                }
                _ = v80
            }
            for fi110 := 0; fi110 < 17; fi110++ {
                if false {
                    v6[32767] = true
                } else {
                    switch 1 {
                    case -536:
                        fmt.Println(v17)
                    case -757:
                        fmt.Println(v10)
                    case 1:
                        fmt.Println(v16, v5)
                    }
                    if v17 {
                        var v111 uint32 = 249
                        _ = v111
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        switch ((("" + "a") + (("\n\t" + "") + "")) + "a") {
                        case "foo bar":
                            var v112 uint64 = 18446744073709551615
                            _ = v112
                        case "a":
                            if (0 != (^(v5))) {
                                fmt.Println(v16, v5, fi79)
                            }
                        }
                        v4 = v4
                        fmt.Println(v17, v10)
                    }
                    if (!v16) {
                        var v113 uint16 = 509
                        var v114 uint8 = 2
                        _ = v113
                        _ = v114
                    } else {
                        v9 += v9
                        var v115 uint = 2
                        if v17 {
                            v116 := v6[fi110]
                            _ = v116
                        } else {
                            v6[(fi79 - (fi79 / ((821) | 1)))] = ((v4 - v4) > int32(v9))
                            if (float32(-2.25) < 0.0) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                fmt.Println(fi79)
                            } else {
                                if (4294967295 != fi110) {
                                    v115 -= v115
                                } else {
                                    v10 = v10
                                    v117 := GenS5{F0: ((-2147483648 < (v4 + (v4 & (v4 / ((v4) | 1))))) && (2 > v115)), F1: 2}
                                    fmt.Println(v16, v5)
                                    _ = v117
                                }
                                var v118 float32 = float32(v1)
                                _ = v118
                            }
                        }
                        _ = v115
                    }
                }
                switch ("\n\t" + ("hello" + ("hello" + "\n\t"))) {
                case "a":
                    v10 -= (v10 >> ((76) & 63))
                case "hello":
                    switch "foo bar" {
                    case "a":
                        switch ("foo bar" + "\n\t") {
                        case "":
                            v16 = (65535 != 42)
                        case "\n\t":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case "":
                        switch (fi110 % ((fi79) | 1)) {
                        case 42:
                            switch (fi110 & (^(fi110))) {
                            case 42:
                                var v119 int = (fi110 & (-(fi110)))
                                _ = v119
                            case 127:
                                if (131 == -306) {
                                    v9 = 100
                                } else {
                                    _ = len(v6)
                                    if v17 {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        switch fi110 {
                                        case 4294967295:
                                            v3 = append(v3, fi79)
                                        case 128:
                                            var v120 int32 = v4
                                            _ = v120
                                        case -156:
                                            v1 += (v1 - float64(-2.25))
                                        default:
                                            if true {
                                                if v17 {
                                                    v9 += v9
                                                    var v121 int32 = int32(fi110)
                                                    _ = v121
                                                }
                                                var v122 uint16 = 2
                                                if true {
                                                    if (!false) {
                                                        if v17 {
                                                            v123 := GenS5{F0: (fi79 != (fi110 + fi79)), F1: 127}
                                                            switch "a" {
                                                            case "":
                                                                switch "" {
                                                                case "":
                                                                    v4 = (v4 / ((int32(v9)) | 1))
                                                                case "a":
                                                                    switch (fi79 | fi110) {
                                                                    case 695:
                                                                        switch fi79 {
                                                                        case 1:
                                                                            var v124 int32 = v4
                                                                            _ = v124
                                                                        }
                                                                    default:
                                                                        fmt.Println(v9, v10, v4)
                                                                    }
                                                                default:
                                                                    v5 -= v9
                                                                }
                                                            case "foo bar":
                                                                switch fi79 {
                                                                case 931:
                                                                    delete(v2, 4294967295)
                                                                default:
                                                                    v122 -= 32767
                                                                    var v125 uint64 = 0
                                                                    _ = v125
                                                                }
                                                                var v126 int16 = v10
                                                                _ = v126
                                                            }
                                                            fmt.Println(v9, v122, v10)
                                                            var v127 uint = 4294967295
                                                            _ = v123
                                                            _ = v127
                                                        }
                                                        delete(v6, (fi110 ^ fi110))
                                                        switch (^(fi79)) {
                                                        case 127:
                                                            var v128 uint16 = v122
                                                            _ = v128
                                                        case 633:
                                                            switch (^(fi79)) {
                                                            case -1:
                                                                v129 := GenS13{F0: 2147483647, F1: 18446744073709551615}
                                                                _ = v129
                                                            }
                                                        default:
                                                            v130 := GenS11{F0: v4, F1: 0, F2: (v122 + (v122 ^ v122))}
                                                            _ = v130
                                                        }
                                                        v16 = ((v16 && v16) && v16)
                                                    }
                                                } else {
                                                    fmt.Println(v9)
                                                    v4 -= (v4 >> ((502) & 63))
                                                    v122 += (v122 | 439)
                                                }
                                                switch "hello" {
                                                case "foo bar":
                                                    if v16 {
                                                        var v131 int32 = (v4 << ((65535) & 63))
                                                        _ = v131
                                                    } else {
                                                        if true {
                                                            v17 = ((v10 - int16(v122)) == (v10 | 1))
                                                            var v132 string = (("\n\t" + "a") + "\n\t")
                                                            if (942 == 4294967295) {
                                                                if v17 {
                                                                    v16 = (v17 || (523 < fi79))
                                                                } else {
                                                                    fmt.Println(v9, v17, fi79)
                                                                    var v133 string = v132
                                                                    if ((v10 + 1) > (v10 * (v10 % ((v10) | 1)))) {
                                                                        v1 = (-2.25 * v1)
                                                                        v6[812] = (!(v122 == (v122 ^ v122)))
                                                                    } else {
                                                                        var v134 uint32 = 2
                                                                        v3 = append(v3, fi110)
                                                                        if (((1.5 / 1e10) + (-2.25 / ((float32((1e10 - 0.0)) * 0.0) + -2.25))) != 1e10) {
                                                                            v122 += v122
                                                                        } else {
                                                                            switch ("" + ("hello" + (v132 + "a"))) {
                                                                            case "foo bar":
                                                                                if v17 {
                                                                                    v135 := GenS14{F0: 42, F1: v134}
                                                                                    v136 := GenS13{F0: (^(v134)), F1: 1}
                                                                                    _ = v135
                                                                                    _ = v136
                                                                                } else {
                                                                                    var v137 uint64 = 1
                                                                                    _ = len(v3)
                                                                                    if v16 {
                                                                                        v138 := GenS7{F0: 9223372036854775807}
                                                                                        v139 := GenS15{F0: false, F1: v122}
                                                                                        v4 -= -2147483648
                                                                                        _ = v138
                                                                                        _ = v139
                                                                                    }
                                                                                    _ = v137
                                                                                }
                                                                            }
                                                                            if (86 >= (v134 >> ((42) & 63))) {
                                                                                fmt.Println(v10, fi110, v1)
                                                                                v140 := GenS16{F0: 323, F1: v4}
                                                                                v17 = v17
                                                                                _ = v140
                                                                            }
                                                                        }
                                                                        _ = v134
                                                                    }
                                                                    _ = v133
                                                                }
                                                            }
                                                            if v16 {
                                                                var v141 int64 = 0
                                                                v16 = v16
                                                                var v142 string = "foo bar"
                                                                _ = v141
                                                                _ = v142
                                                            } else {
                                                                var v143 float32 = 1.5
                                                                v144 := GenS4{F0: "\n\t", F1: v17, F2: v4}
                                                                _ = v143
                                                                _ = v144
                                                            }
                                                            _ = v132
                                                        }
                                                        v3 = append(v3, fi110)
                                                    }
                                                case "a":
                                                    if len(v3) > 0 {
                                                        v3[0] = -434
                                                    }
                                                }
                                                _ = v122
                                            } else {
                                                if (!v17) {
                                                    v145 := GenS17{F0: 2, F1: v10}
                                                    _ = v145
                                                }
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v146 := GenS5{F0: (v17 && (!v16)), F1: 0}
                                                    _ = v146
                                                case "":
                                                    var v147 int16 = (^(v10))
                                                    _ = v147
                                                case "hello":
                                                    fmt.Println(v9, v16)
                                                }
                                                var v148 string = ("\n\t" + "\n\t")
                                                _ = v148
                                            }
                                        }
                                        var v149 int32 = (v4 << ((2) & 63))
                                        var v150 uint32 = 963
                                        _ = v149
                                        _ = v150
                                    }
                                    if v17 {
                                        fmt.Println(v1, v16)
                                        v1 += v1
                                    } else {
                                        delete(v6, (^(fi110)))
                                        switch fi79 {
                                        case -980:
                                            v151 := GenS18{F0: (!(255 > 42)), F1: 181, F2: v4}
                                            _ = v151
                                        }
                                        v9 = -85
                                    }
                                    var v152 uint = 18446744073709551615
                                    var v153 float32 = 1.5
                                    if (v17 && v16) {
                                        v154 := GenS6{F0: v152}
                                        fmt.Println(v4, v153, v4)
                                        _ = v154
                                    } else {
                                        var v155 bool = ((v10 != (v10 % ((v10) | 1))) || (("a" < "\n\t") && v16))
                                        _ = v155
                                    }
                                    v156 := GenS11{F0: -1, F1: 1, F2: 2}
                                    _ = v152
                                    _ = v153
                                    _ = v156
                                }
                            case 385:
                                var v157 uint64 = 23
                                _ = v157
                            }
                        case -474:
                            fmt.Println(v5)
                        }
                    case "\n\t":
                        if false {
                            var v158 float32 = -2.25
                            v1 = (3.14159 - v1)
                            if (!true) {
                                fmt.Println(v1, fi79, v16)
                                var v159 bool = (!v16)
                                _ = v159
                            } else {
                                v5 += v9
                                var v160 int16 = (v10 << ((707) & 63))
                                _ = v160
                            }
                            _ = v158
                        }
                    default:
                        v16 = v16
                    }
                case "\n\t":
                    var v161 uint64 = 2
                    _ = v161
                }
                var v162 int8 = v9
                v163 := v6[fi110]
                if false {
                    v163 = v16
                    v164 := GenS19{F0: v162}
                    if (v5 > v162) {
                        v9 -= v162
                    }
                    _ = v164
                } else {
                    var v165 uint16 = 255
                    _ = v165
                }
                _ = fi110
                _ = v162
                _ = v163
            }
            v10 = v10
            _ = fi79
        }
    default:
        for fi166 := 0; fi166 < 3; fi166++ {
            v167 := GenS2{F0: 237, F1: (^(v5))}
            delete(v6, 42)
            for fi168 := 0; fi168 < 15; fi168++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                delete(v6, fi168)
                var v169 float32 = (float32(0.1) * 3.14159)
                _ = fi168
                _ = v169
            }
            _ = fi166
            _ = v167
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
