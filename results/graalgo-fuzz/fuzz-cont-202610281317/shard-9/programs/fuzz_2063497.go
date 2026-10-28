package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 int; F2 string }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 int32; F1 int64; F2 bool }
type GenS8 struct { F0 int8; F1 uint64; F2 uint }
type GenS9 struct { F0 int32; F1 int; F2 int16 }
type GenS10 struct { F0 bool; F1 int8 }
type GenS11 struct { F0 uint8; F1 uint64 }
type GenS12 struct { F0 int32; F1 int32 }
func genFn0(p0 string, p1 uint16, p2 uint) uint8 {
    var v7 int64 = 345
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int64 = v7
        _ = fi8
        _ = v9
    }
    v10 := []int{-789, -1, 2147483647}
    if (!(0 == -2147483648)) {
        p2 = (p2 << (((^(p2))) & 63))
        v11 := GenS1{F0: 3, F1: -658}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 208
}
func genFn1(p0 string, p1 int) uint32 {
    v12 := []int{p1, 1, p1}
    v13 := GenS1{F0: genFn0("a", 798, 0), F1: -1}
    switch (p1 * p1) {
    case 4294967295:
        p0 = p0
    case 1:
        if (-2147483648 >= 1) {
            switch 742 {
            case 0:
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            case -829:
                var v14 bool = false
                _ = v14
            case 128:
                p1 = 9223372036854775807
            default:
                for fi15 := 0; fi15 < 2; fi15++ {
                    v16 := GenS1{F0: 42, F1: 0}
                    var v17 float64 = float64(3.14159)
                    switch "foo bar" {
                    case "a":
                        var v18 uint64 = 2
                        _ = v18
                    case "hello":
                        if (!(!false)) {
                            for fi19 := 0; fi19 < 9; fi19++ {
                                var v20 string = ""
                                p1 = (-(fi19))
                                if true {
                                    fmt.Println(v17)
                                    var v21 uint8 = genFn0("\n\t", 32767, 4294967295)
                                    _ = v21
                                } else {
                                    p0 = ("" + v20)
                                    fmt.Println(fi15, p0, p0)
                                }
                                _ = fi19
                                _ = v20
                            }
                            for fi22 := 0; fi22 < 19; fi22++ {
                                var v23 uint64 = 0
                                v12 = append(v12, -1)
                                _ = fi22
                                _ = v23
                            }
                            var v24 float64 = float64(-2.25)
                            _ = v24
                        }
                    case "foo bar":
                        if (1 > 65535) {
                            for fi25 := 0; fi25 < 5; fi25++ {
                                switch (p1 | (fi15 & fi15)) {
                                case 128:
                                    if true {
                                        if (!true) {
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            var v26 float32 = 1.5
                                            if ((p0 != ("\n\t" + ((("" + "\n\t") + p0) + p0))) && (false && ((false || true) || false))) {
                                                var v27 int64 = 1
                                                v28 := GenS1{F0: 128, F1: 102}
                                                _ = v27
                                                _ = v28
                                            } else {
                                                var v29 uint64 = 2
                                                fmt.Println(v26, fi15, p0)
                                                _ = v29
                                            }
                                            _ = v26
                                        } else {
                                            p1 = (p1 - (fi25 & (p1 & p1)))
                                        }
                                        p1 += fi25
                                        switch p1 {
                                        case 887:
                                            if len(v12) > 0 {
                                                v12[0] = (p1 << ((834) & 63))
                                            }
                                            _ = len(v12)
                                        case 42:
                                            _ = len(v12)
                                        case 783:
                                            if (!(0 != 65535)) {
                                                fmt.Println(p1, p0)
                                                v12 = append(v12, p1)
                                                v30 := GenS1{F0: genFn0("a", 489, 372), F1: 0}
                                                v17 = v17
                                                _ = v30
                                            } else {
                                                switch fi15 {
                                                case 423:
                                                    v31 := GenS2{F0: 65535}
                                                    _ = v31
                                                case 127:
                                                    p0 = p0
                                                }
                                            }
                                        }
                                        var v32 int8 = -128
                                        if ((3.14159 < 1e10) && (!true)) {
                                            if false {
                                                var v33 uint64 = 4294967295
                                                p0 = p0
                                                v17 = v17
                                                _ = v33
                                            }
                                        }
                                        v34 := GenS3{F0: (v32 | v32)}
                                        if (true && false) {
                                            v35 := GenS1{F0: genFn0(p0, 948, 139), F1: -1}
                                            _ = v35
                                        }
                                        switch (fi25 + fi25) {
                                        case 32767:
                                            p1 = 42
                                        }
                                        var v36 int64 = -9223372036854775808
                                        _ = v32
                                        _ = v34
                                        _ = v36
                                    }
                                default:
                                    p1 = -1
                                    p0 = (("foo bar" + p0) + "foo bar")
                                }
                                v37 := GenS2{F0: 65535}
                                var v38 uint64 = 1
                                _ = fi25
                                _ = v37
                                _ = v38
                            }
                            switch p1 {
                            case 517:
                                v39 := GenS2{F0: 65535}
                                _ = v39
                            case 2147483647:
                                var v40 int64 = 136
                                _ = v40
                            case 9223372036854775807:
                                for fi41 := 0; fi41 < 20; fi41++ {
                                    fmt.Println(v17, p1, fi15)
                                    _ = fi41
                                }
                            }
                            var v42 uint = 255
                            _ = v42
                        } else {
                            v17 -= 1e10
                        }
                    default:
                        switch "\n\t" {
                        case "hello":
                            if true {
                                if (1 == 1) {
                                    if len(v12) > 0 {
                                        v12[0] = (p1 - (p1 - 42))
                                    }
                                    if ((p1 << ((701) & 63)) >= p1) {
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        v43 := GenS4{F0: 65535}
                                        if false {
                                            _ = len(v12)
                                        }
                                        _ = v43
                                    } else {
                                        v44 := GenS5{F0: 43, F1: p1, F2: p0}
                                        _ = v44
                                    }
                                }
                                p1 -= p1
                                var v45 string = p0
                                _ = v45
                            } else {
                                if (genFn0("foo bar", 42, 255) > 255) {
                                    if (!false) {
                                        for fi46 := 0; fi46 < 18; fi46++ {
                                            v47 := GenS4{F0: 2147483647}
                                            var v48 string = p0
                                            v48 = p0
                                            _ = fi46
                                            _ = v47
                                            _ = v48
                                        }
                                    }
                                }
                            }
                        case "":
                            switch (p0 + (p0 + p0)) {
                            case "a":
                                fmt.Println(fi15, fi15)
                            default:
                                p0 = ("" + "\n\t")
                            }
                        case "\n\t":
                            for fi49 := 0; fi49 < 13; fi49++ {
                                var v50 uint8 = 42
                                if true {
                                    _ = len(v12)
                                } else {
                                    switch p0 {
                                    case "\n\t":
                                        var v51 int16 = 1
                                        _ = v51
                                    case "a":
                                        var v52 int16 = -32768
                                        _ = v52
                                    case "":
                                        switch -648 {
                                        case 142:
                                            var v53 uint = 938
                                            _ = v53
                                        case 9223372036854775807:
                                            v54 := GenS6{F0: p0, F1: 1}
                                            _ = v54
                                        }
                                    }
                                }
                                v17 = -2.25
                                v55 := GenS3{F0: -1}
                                _ = fi49
                                _ = v50
                                _ = v55
                            }
                        default:
                            v12 = append(v12, fi15)
                        }
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            if (false && (true && false)) {
                v56 := GenS2{F0: 2147483647}
                p0 = (p0 + p0)
                var v57 string = p0
                _ = v56
                _ = v57
            } else {
                var v58 float32 = 3.14159
                if true {
                    switch -2147483648 {
                    case 415:
                        var v59 int32 = 2147483647
                        _ = v59
                    }
                    var v60 bool = (p0 <= p0)
                    p0 = "a"
                    _ = v60
                }
                v58 -= (v58 - float32(0.1))
                _ = v58
            }
        } else {
            p0 = ""
            fmt.Println(p1, p1)
            if ((((p0 + p0) + p0) + p0) != "hello") {
                if false {
                    for fi61 := 0; fi61 < 15; fi61++ {
                        if (!true) {
                            fmt.Println(p0, p0, p1)
                            p1 = fi61
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                        } else {
                            p0 = ("foo bar" + "foo bar")
                        }
                        switch "\n\t" {
                        case "":
                            var v62 string = p0
                            _ = v62
                        case "\n\t":
                            var v63 uint8 = 109
                            _ = v63
                        default:
                            for fi64 := 0; fi64 < 18; fi64++ {
                                var v65 int64 = -9223372036854775808
                                v66 := GenS5{F0: 119, F1: 127, F2: p0}
                                var v67 int16 = 42
                                if false {
                                    var v68 int = fi64
                                    var v69 string = p0
                                    switch (v68 - 32768) {
                                    case 521:
                                        v12 = append(v12, (v68 ^ p1))
                                    case 1:
                                        v12 = append(v12, fi61)
                                    }
                                    _ = v68
                                    _ = v69
                                } else {
                                    switch (-(fi61)) {
                                    case 0:
                                        switch p0 {
                                        case "foo bar":
                                            var v70 uint64 = 42
                                            _ = v70
                                        case "\n\t":
                                            fmt.Println(v65, v67, fi61)
                                        }
                                    default:
                                        var v71 int = (^(fi64))
                                        _ = v71
                                    }
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    if ((112 == -2147483648) || (1 >= -125)) {
                                        var v72 uint64 = 639
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        v65 = (v65 % (((v65 * (-(v65)))) | 1))
                                        _ = v72
                                    } else {
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        var v73 float64 = 0.1
                                        var v74 int = (p1 * (fi61 << ((255) & 63)))
                                        _ = v73
                                        _ = v74
                                    }
                                    if (!(true || true)) {
                                        var v75 int32 = 190
                                        switch p0 {
                                        case "a":
                                            p1 += (fi61 ^ 32768)
                                        default:
                                            switch "foo bar" {
                                            case "a":
                                                if (float32((1.5 / 0.1)) == float32(1e10)) {
                                                    v76 := GenS3{F0: -128}
                                                    var v77 int8 = -1
                                                    _ = v76
                                                    _ = v77
                                                } else {
                                                    fmt.Println(v75)
                                                    switch p0 {
                                                    case "foo bar":
                                                        v78 := GenS1{F0: 0, F1: (v75 ^ (v75 * 1))}
                                                        _ = v78
                                                    case "\n\t":
                                                        switch int(v67) {
                                                        case 128:
                                                            v79 := GenS6{F0: ("" + "a"), F1: 549}
                                                            _ = v79
                                                        case -514:
                                                            fmt.Println(fi61)
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                p1 -= fi64
                                                            }
                                                        case 32767:
                                                            v75 = (^(v75))
                                                        }
                                                    case "hello":
                                                        var v80 uint8 = genFn0(p0, 2, 255)
                                                        _ = v80
                                                    }
                                                    v81 := GenS2{F0: 496}
                                                    _ = v81
                                                }
                                            default:
                                                v82 := GenS3{F0: 92}
                                                switch fi64 {
                                                case -969:
                                                    v75 += v75
                                                case -866:
                                                    v65 += v65
                                                }
                                                var v83 uint16 = 510
                                                var v84 uint16 = (v83 % ((981) | 1))
                                                _ = v82
                                                _ = v83
                                                _ = v84
                                            }
                                        }
                                        v12 = append(v12, (p1 << ((42) & 63)))
                                        fmt.Println(v75, v65)
                                        _ = v75
                                    }
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                }
                                _ = fi64
                                _ = v65
                                _ = v66
                                _ = v67
                            }
                        }
                        var v85 float64 = (-2.25 + ((-2.25 / -2.25) + float64(0.0)))
                        _ = fi61
                        _ = v85
                    }
                    var v86 uint16 = 1
                    _ = v86
                } else {
                    for fi87 := 0; fi87 < 10; fi87++ {
                        p0 = ((p0 + p0) + p0)
                        var v88 uint16 = 65535
                        _ = fi87
                        _ = v88
                    }
                }
                var v89 uint8 = 135
                var v90 int64 = 42
                _ = v89
                _ = v90
            } else {
                switch (("" + p0) + "\n\t") {
                case "a":
                    v91 := GenS4{F0: 444}
                    _ = v91
                case "":
                    p0 = p0
                default:
                    p1 = -2147483648
                }
                switch p0 {
                case "\n\t":
                    switch (((p0 + p0) + "a") + p0) {
                    case "":
                        if ((!true) || false) {
                            v92 := GenS3{F0: 116}
                            p0 = "hello"
                            _ = len(v12)
                            _ = v92
                        } else {
                            switch "" {
                            case "a":
                                if len(v12) > 0 {
                                    v12[0] = 2147483647
                                }
                            case "hello":
                                v93 := GenS5{F0: 127, F1: p1, F2: (p0 + p0)}
                                _ = v93
                            case "foo bar":
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            default:
                                p1 += p1
                            }
                            if (128 >= 128) {
                                var v94 float64 = (3.14159 + (-2.25 * (-2.25 - -2.25)))
                                _ = v94
                            } else {
                                for fi95 := 0; fi95 < 1; fi95++ {
                                    if (42 > p1) {
                                        v96 := GenS6{F0: (p0 + (("a" + (p0 + ("\n\t" + "a"))) + ("foo bar" + p0))), F1: 255}
                                        _ = v96
                                    } else {
                                        var v97 int = (fi95 << ((908) & 63))
                                        _ = v97
                                    }
                                    var v98 float32 = 1e10
                                    fmt.Println(p0)
                                    _ = fi95
                                    _ = v98
                                }
                            }
                            var v99 uint32 = 1
                            _ = v99
                        }
                    case "\n\t":
                        v12 = append(v12, (p1 * 32768))
                    case "foo bar":
                        if true {
                            if false {
                                var v100 uint8 = 42
                                fmt.Println(p0, v100)
                                var v101 int8 = -50
                                _ = v100
                                _ = v101
                            }
                            v102 := GenS7{F0: -393, F1: 42, F2: ((1.5 * float32(float64(1.5))) > (-2.25 - float32(1.5)))}
                            v103 := GenS8{F0: 0, F1: 2, F2: 83}
                            _ = v102
                            _ = v103
                        } else {
                            switch (p1 - (p1 / ((p1) | 1))) {
                            case 32767:
                                switch (p0 + p0) {
                                case "a":
                                    if true {
                                        if false {
                                            p1 -= 127
                                            fmt.Println(p1)
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                        } else {
                                            p0 = "foo bar"
                                        }
                                        v104 := GenS2{F0: 2}
                                        v105 := GenS8{F0: -1, F1: 900, F2: 236}
                                        _ = v104
                                        _ = v105
                                    }
                                case "hello":
                                    _ = len(v12)
                                case "foo bar":
                                    switch p1 {
                                    case 0:
                                        if (26 == 2147483647) {
                                            p0 = ((p0 + p0) + "a")
                                            switch (p1 << ((0) & 63)) {
                                            case 2147483647:
                                                p1 = (p1 - (-(p1)))
                                            case -624:
                                                var v106 string = p0
                                                _ = v106
                                            }
                                        } else {
                                            switch 4294967295 {
                                            case 0:
                                                var v107 uint8 = genFn0(p0, 0, 78)
                                                _ = v107
                                            case -492:
                                                var v108 int16 = 0
                                                _ = v108
                                            default:
                                                if len(v12) >= 2 {
                                                    v12 = v12[1:len(v12)]
                                                }
                                            }
                                            p0 = "foo bar"
                                        }
                                    case -300:
                                        switch -566 {
                                        case 128:
                                            var v109 uint = 65535
                                            _ = v109
                                        case -1:
                                            v110 := GenS3{F0: 0}
                                            _ = v110
                                        case 42:
                                            for fi111 := 0; fi111 < 10; fi111++ {
                                                if (!(!true)) {
                                                    for fi112 := 0; fi112 < 10; fi112++ {
                                                        if true {
                                                            p0 = ("\n\t" + ((("hello" + p0) + "a") + ("" + "\n\t")))
                                                            switch p0 {
                                                            case "":
                                                                p1 += 128
                                                            default:
                                                                var v113 uint16 = 65535
                                                                _ = v113
                                                            }
                                                            v114 := GenS2{F0: 65535}
                                                            _ = v114
                                                        }
                                                        v115 := GenS9{F0: 42, F1: (-(fi111)), F2: 0}
                                                        _ = fi112
                                                        _ = v115
                                                    }
                                                    for fi116 := 0; fi116 < 9; fi116++ {
                                                        var v117 int8 = -1
                                                        _ = fi116
                                                        _ = v117
                                                    }
                                                } else {
                                                    p0 = (((p0 + "foo bar") + "hello") + (("\n\t" + p0) + (p0 + "")))
                                                    if (genFn0((p0 + "a"), 42, 42) == 127) {
                                                        if len(v12) > 0 {
                                                            v12[0] = fi111
                                                        }
                                                    } else {
                                                        var v118 int16 = -200
                                                        for fi119 := 0; fi119 < 9; fi119++ {
                                                            if ((0.0 - ((1e10 - 1e10) / -2.25)) != 3.14159) {
                                                                p0 = (p0 + p0)
                                                                var v120 int32 = 2147483647
                                                                _ = v120
                                                            } else {
                                                                switch fi111 {
                                                                case -9223372036854775808:
                                                                    if len(v12) >= 2 {
                                                                        v12 = v12[1:len(v12)]
                                                                    }
                                                                }
                                                            }
                                                            _ = fi119
                                                        }
                                                        _ = v118
                                                    }
                                                }
                                                p1 = 127
                                                _ = fi111
                                            }
                                        default:
                                            p0 = p0
                                        }
                                    }
                                default:
                                    var v121 float32 = 3.14159
                                    _ = v121
                                }
                            default:
                                v122 := GenS8{F0: 19, F1: 42, F2: 255}
                                _ = v122
                            }
                        }
                    }
                case "a":
                    v123 := GenS9{F0: 1, F1: (p1 / ((p1) | 1)), F2: 1}
                    _ = v123
                case "hello":
                    fmt.Println(p1, p1)
                default:
                    for fi124 := 0; fi124 < 15; fi124++ {
                        for fi125 := 0; fi125 < 3; fi125++ {
                            v126 := GenS10{F0: true, F1: -38}
                            switch (p1 + 9223372036854775807) {
                            case 4294967295:
                                if len(v12) > 0 {
                                    v12[0] = fi124
                                }
                            case 1:
                                fmt.Println(fi124)
                            case 2147483647:
                                v127 := GenS6{F0: p0, F1: 505}
                                _ = v127
                            default:
                                fmt.Println(p0)
                            }
                            p1 += p1
                            if len(v12) > 0 {
                                v12[0] = fi125
                            }
                            switch (fi125 & (fi124 + 544)) {
                            case -1:
                                switch "" {
                                case "foo bar":
                                    if (682 != 568) {
                                        if (true || (32 <= 42)) {
                                            var v128 int = p1
                                            p1 = (p1 | v128)
                                            p1 += 32768
                                            var v129 int16 = -32768
                                            _ = v128
                                            _ = v129
                                        }
                                        switch fi125 {
                                        case 32767:
                                            p0 = "foo bar"
                                        case -483:
                                            var v130 uint = 65535
                                            _ = v130
                                        default:
                                            if true {
                                                var v131 uint64 = 42
                                                var v132 int32 = 452
                                                _ = v131
                                                _ = v132
                                            }
                                        }
                                        p1 += (p1 & 128)
                                    }
                                    v133 := GenS10{F0: true, F1: 1}
                                    fmt.Println(p1, fi124)
                                    _ = v133
                                case "hello":
                                    switch p0 {
                                    case "foo bar":
                                        p0 = ("hello" + "hello")
                                    case "\n\t":
                                        if len(v12) > 0 {
                                            v12[0] = -1
                                        }
                                    }
                                case "":
                                    var v134 uint16 = 190
                                    switch 32768 {
                                    case 32767:
                                        p0 = p0
                                    default:
                                        switch fi124 {
                                        case -967:
                                            if true {
                                                p1 = 633
                                            } else {
                                                v12 = append(v12, fi125)
                                            }
                                        case 32768:
                                            p1 -= (fi125 | int(v134))
                                        default:
                                            var v135 uint = 463
                                            _ = v135
                                        }
                                    }
                                    _ = v134
                                }
                                fmt.Println(p1, fi124, fi124)
                                p1 += fi124
                            default:
                                if (255 != 65535) {
                                    v136 := GenS5{F0: genFn0("", 1, 4294967295), F1: (fi124 % ((fi125) | 1)), F2: "a"}
                                    _ = v136
                                }
                            }
                            _ = fi125
                            _ = v126
                        }
                        if (float64(0.1) <= (1.5 * 1.5)) {
                            p0 = p0
                            p1 += -1
                        }
                        _ = fi124
                    }
                }
            }
        }
    case 32768:
        switch p0 {
        case "a":
            var v137 uint8 = 127
            _ = v137
        case "\n\t":
            switch (("\n\t" + "") + "\n\t") {
            case "\n\t":
                fmt.Println(p0, p1)
            case "foo bar":
                _ = len(v12)
            }
        case "foo bar":
            var v138 uint8 = 53
            _ = v138
        }
    }
    for fi139 := 0; fi139 < 8; fi139++ {
        var v140 float32 = 0.1
        _ = fi139
        _ = v140
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 627
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    var v3 uint32 = 759
    var v4 float32 = (((3.14159 + 0.1) * 1.5) * 1.5)
    v5 := make(map[string]uint32)
    v6 := make(map[int]uint32)
    _ = len(v5)
    fmt.Println(v3)
    _ = len(v1)
    for fi141 := 0; fi141 < 3; fi141++ {
        v142 := GenS9{F0: -1, F1: (fi141 >> ((120) & 63)), F2: -32768}
        v143 := []int{fi141}
        _ = fi141
        _ = v142
        _ = v143
    }
    switch 2147483647 {
    case 0:
        v144 := v2[128]
        _ = v144
    default:
        var v145 int64 = -1
        _ = v145
    }
    fmt.Println(v4, v4)
    v146 := GenS8{F0: 0, F1: 897, F2: 1}
    var v147 uint8 = 42
    delete(v1, "a")
    v3 += v3
    var v148 float64 = (3.14159 / -2.25)
    switch -848 {
    case 0:
        if ((v147 | (^(v147))) != v147) {
            v149 := []int{}
            _ = v149
        } else {
            v150 := GenS2{F0: v3}
            if (((v4 - 0.0) - v4) == v4) {
                v4 -= (v4 / (v4 / ((v4 * float32(1e10)) / -2.25)))
            }
            for fi151 := 0; fi151 < 12; fi151++ {
                if true {
                    for fi152 := 0; fi152 < 20; fi152++ {
                        v147 -= v147
                        fmt.Println(v4, v148, v3)
                        var v153 float32 = v4
                        _ = fi152
                        _ = v153
                    }
                    v154 := GenS11{F0: genFn0("foo bar", 42, 236), F1: 4294967295}
                    _ = v154
                } else {
                    var v155 uint64 = 0
                    v3 = genFn1((("" + "") + "hello"), fi151)
                    for fi156 := 0; fi156 < 2; fi156++ {
                        delete(v1, ("" + "hello"))
                        v157 := GenS12{F0: 0, F1: 42}
                        _ = fi156
                        _ = v157
                    }
                    _ = v155
                }
                _ = fi151
            }
            _ = v150
        }
    case 32768:
        v4 -= float32((v148 - (v148 * 0.1)))
    case 2147483647:
        if (!(!(!false))) {
            v148 -= v148
        }
    }
    v5["hello"] = (v3 % ((v3) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v147:", v147)
    fmt.Println("v148:", v148)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v146
    _ = v147
    _ = v148
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
