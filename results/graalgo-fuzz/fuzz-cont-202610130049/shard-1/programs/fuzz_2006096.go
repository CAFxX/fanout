package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint; F1 int64; F2 string }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int16; F1 bool; F2 int8 }
type GenS10 struct { F0 int8; F1 int16 }
type GenS11 struct { F0 uint64; F1 uint8; F2 bool }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint64 }
type GenS15 struct { F0 uint64; F1 string }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    var v3 uint = 2
    v4 := make(map[uint32]bool)
    if (true || false) {
        delete(v1, ("\n\t" + "\n\t"))
        var v5 float32 = (1.5 + 1e10)
        v6 := v2[961]
        _ = v5
        _ = v6
    } else {
        switch ("\n\t" + "\n\t") {
        case "hello":
            _ = len(v2)
        case "foo bar":
            var v7 int16 = -1
            _ = v7
        }
        var v8 uint64 = 644
        _ = v8
    }
    v3 += 65535
    v9 := GenS1{F0: "hello", F1: 32767, F2: 208}
    if false {
        v10 := v1[""]
        v11 := []int{128, -9223372036854775808}
        v12 := GenS1{F0: "foo bar", F1: 255, F2: 0}
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := []int{9223372036854775807, -23, 1, -534}
    if len(v13) > 0 {
        v13[0] = 4294967295
    }
    switch (("foo bar" + "\n\t") + "a") {
    case "":
        v3 += v3
    case "foo bar":
        for fi14 := 0; fi14 < 5; fi14++ {
            if (false || true) {
                if true {
                    v15 := GenS2{F0: 1, F1: int64(v3), F2: "a"}
                    _ = v15
                }
                for fi16 := 0; fi16 < 4; fi16++ {
                    var v17 uint = v3
                    if false {
                        fmt.Println(fi16)
                        if true {
                            if true {
                                var v18 uint8 = 2
                                v3 += 818
                                _ = v18
                            } else {
                                v3 = (v3 - 366)
                            }
                            if ((!(("\n\t" <= "foo bar") || ((((fi14 << ((v17) & 63)) <= (fi16 & fi14)) && (!(32767 > -32768))) || true))) && (true && (42 == 2147483647))) {
                                if true {
                                    var v19 uint = (v17 / ((v3) | 1))
                                    if false {
                                        switch "a" {
                                        case "":
                                            var v20 int64 = -1
                                            fmt.Println(v17, v3)
                                            v3 = (v19 >> ((v19) & 63))
                                            _ = v20
                                        }
                                        var v21 int64 = -1
                                        _ = v21
                                    } else {
                                        if false {
                                            v22 := GenS2{F0: v3, F1: 42, F2: "hello"}
                                            var v23 bool = (true || true)
                                            _ = v22
                                            _ = v23
                                        }
                                        v19 -= v19
                                        if (fi14 <= (fi16 & (fi16 % (((fi14 / ((fi16) | 1))) | 1)))) {
                                            v13 = append(v13, fi14)
                                        }
                                    }
                                    _ = v19
                                }
                                delete(v2, 1)
                                v13 = append(v13, (^(fi14)))
                            }
                        } else {
                            if true {
                                v3 -= 18446744073709551615
                            } else {
                                if len(v13) > 0 {
                                    v13[0] = 4294967295
                                }
                                fmt.Println(fi16, v17)
                                v24 := GenS1{F0: ("foo bar" + ""), F1: 42, F2: 99}
                                _ = v24
                            }
                            fmt.Println(fi14, fi16, v17)
                            delete(v1, "foo bar")
                            if false {
                                var v25 uint16 = 252
                                if false {
                                    v25 -= v25
                                    if (!false) {
                                        v26 := GenS1{F0: "foo bar", F1: 65535, F2: 2}
                                        switch fi16 {
                                        case 983:
                                            v17 -= v17
                                        case -326:
                                            var v27 string = ("a" + "hello")
                                            _ = v27
                                        }
                                        _ = v26
                                    } else {
                                        var v28 uint16 = 2
                                        v3 += v17
                                        _ = v28
                                    }
                                } else {
                                    var v29 uint64 = 1
                                    v30 := GenS2{F0: v3, F1: -1, F2: (("foo bar" + "foo bar") + "a")}
                                    v31 := GenS3{F0: 18446744073709551615, F1: -128}
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                                _ = v25
                            }
                            fmt.Println(v17, fi14)
                        }
                        v32 := GenS1{F0: ("foo bar" + "\n\t"), F1: 255, F2: 241}
                        _ = v32
                    }
                    _ = fi16
                    _ = v17
                }
                if false {
                    if true {
                        if (true || (true && (!(-9223372036854775808 == 9223372036854775807)))) {
                            if true {
                                for fi33 := 0; fi33 < 4; fi33++ {
                                    switch "a" {
                                    case "a":
                                        v34 := GenS4{F0: (^(v3))}
                                        _ = v34
                                    case "foo bar":
                                        switch fi33 {
                                        case 42:
                                            v3 = 919
                                        default:
                                            if false {
                                                fmt.Println(v3, fi33)
                                            } else {
                                                v35 := GenS1{F0: "hello", F1: 116, F2: 7}
                                                _ = v35
                                            }
                                        }
                                    default:
                                        v13 = append(v13, fi33)
                                    }
                                    v3 = v3
                                    switch 249 {
                                    case 712:
                                        switch ("a" + "\n\t") {
                                        case "foo bar":
                                            if false {
                                                var v36 int = fi33
                                                v2[1] = "foo bar"
                                                if (!(true || false)) {
                                                    switch ((("" + "foo bar") + ("a" + "hello")) + "a") {
                                                    case "a":
                                                        v37 := GenS5{F0: 0, F1: v3}
                                                        _ = v37
                                                    }
                                                }
                                                _ = v36
                                            } else {
                                                switch fi33 {
                                                case 1:
                                                    switch (("\n\t" + "a") + "foo bar") {
                                                    case "hello":
                                                        if len(v13) >= 2 {
                                                            v13 = v13[1:len(v13)]
                                                        }
                                                    case "":
                                                        v38 := GenS2{F0: 255, F1: -1, F2: "foo bar"}
                                                        _ = v38
                                                    case "a":
                                                        var v39 uint8 = 252
                                                        _ = v39
                                                    default:
                                                        var v40 int64 = -1
                                                        _ = v40
                                                    }
                                                    v41 := GenS3{F0: 2, F1: -1}
                                                    if (!(!false)) {
                                                        if (65535 != 255) {
                                                            var v42 uint32 = 42
                                                            v4[v42] = true
                                                            _ = v42
                                                        }
                                                        v43 := GenS3{F0: 1, F1: -128}
                                                        if (true && true) {
                                                            var v44 uint16 = 42
                                                            if false {
                                                                switch (("" + "a") + (("\n\t" + ("a" + "hello")) + "hello")) {
                                                                case "foo bar":
                                                                    if false {
                                                                        switch fi33 {
                                                                        case 32768:
                                                                            v3 += v3
                                                                            if false {
                                                                                if len(v13) >= 2 {
                                                                                    v13 = v13[1:len(v13)]
                                                                                }
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    switch fi14 {
                                                                                    case 600:
                                                                                        fmt.Println(v3, v3)
                                                                                    }
                                                                                }
                                                                                if (((-2.25 <= (((0.0 / ((0.1 / (1.5 - 0.0)) + -2.25)) * 0.0) - 1.5)) && false) && (!false)) {
                                                                                    v45 := GenS6{F0: (v3 | v3)}
                                                                                    if true {
                                                                                        var v46 string = "hello"
                                                                                        if true {
                                                                                            v47 := GenS4{F0: 255}
                                                                                            if false {
                                                                                                var v48 int = fi33
                                                                                                var v49 float64 = (float64(1e10) / 1.0)
                                                                                                _ = v48
                                                                                                _ = v49
                                                                                            } else {
                                                                                                var v50 uint16 = (v44 - (v44 - (v44 >> ((v3) & 63))))
                                                                                                v51 := GenS1{F0: (v46 + v46), F1: v44, F2: 128}
                                                                                                _ = v50
                                                                                                _ = v51
                                                                                            }
                                                                                            v52 := GenS2{F0: v3, F1: 9223372036854775807, F2: v46}
                                                                                            v46 = v46
                                                                                            v46 = v46
                                                                                            if len(v13) >= 2 {
                                                                                                v13 = v13[1:len(v13)]
                                                                                            }
                                                                                            _ = v47
                                                                                            _ = v52
                                                                                        }
                                                                                        v44 = 65535
                                                                                        _ = v46
                                                                                    } else {
                                                                                        if len(v13) >= 2 {
                                                                                            v13 = v13[1:len(v13)]
                                                                                        }
                                                                                        switch ("a" + (("hello" + "\n\t") + ("foo bar" + ("foo bar" + "foo bar")))) {
                                                                                        case "a":
                                                                                            var v53 uint64 = 1
                                                                                            _ = v53
                                                                                        case "":
                                                                                            if (!true) {
                                                                                                switch "\n\t" {
                                                                                                case "\n\t":
                                                                                                    v3 -= (v3 | v3)
                                                                                                    var v54 uint64 = 1
                                                                                                    switch (fi33 & 2147483647) {
                                                                                                    case 32767:
                                                                                                        fmt.Println(fi33)
                                                                                                    }
                                                                                                    delete(v1, "")
                                                                                                    _ = v54
                                                                                                }
                                                                                                var v55 int16 = 109
                                                                                                switch ("foo bar" + "\n\t") {
                                                                                                case "hello":
                                                                                                    v55 = 881
                                                                                                case "":
                                                                                                    v55 = v55
                                                                                                case "foo bar":
                                                                                                    if true {
                                                                                                        v44 = (v44 | v44)
                                                                                                        v56 := v2[2147483647]
                                                                                                        _ = v56
                                                                                                    }
                                                                                                default:
                                                                                                    v55 += (-(v55))
                                                                                                }
                                                                                                _ = v55
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v57 := GenS1{F0: ("foo bar" + "\n\t"), F1: (^(v44)), F2: 120}
                                                                                    _ = v45
                                                                                    _ = v57
                                                                                }
                                                                                v3 = v3
                                                                            } else {
                                                                                _ = len(v2)
                                                                            }
                                                                            switch fi14 {
                                                                            case 9223372036854775807:
                                                                                v3 -= v3
                                                                            case 2147483647:
                                                                                _ = len(v1)
                                                                            }
                                                                        case 575:
                                                                            if false {
                                                                                var v58 string = ""
                                                                                if ((v58 <= (v58 + (((v58 + v58) + "") + "a"))) && (true && false)) {
                                                                                    v59 := GenS5{F0: 2, F1: v3}
                                                                                    v60 := GenS7{F0: 1}
                                                                                    _ = v59
                                                                                    _ = v60
                                                                                }
                                                                                if (42 > 1) {
                                                                                    var v61 uint64 = 2
                                                                                    fmt.Println(fi14)
                                                                                    if true {
                                                                                        v62 := GenS3{F0: 0, F1: -54}
                                                                                        var v63 uint64 = (v61 | v61)
                                                                                        if (int8(fi14) >= -19) {
                                                                                            fmt.Println(v63, v61)
                                                                                        }
                                                                                        _ = v62
                                                                                        _ = v63
                                                                                    }
                                                                                    _ = v61
                                                                                } else {
                                                                                    var v64 float64 = 3.14159
                                                                                    if true {
                                                                                        v44 -= v44
                                                                                    } else {
                                                                                        v65 := v2[2147483648]
                                                                                        v66 := GenS4{F0: (v3 / ((v3) | 1))}
                                                                                        var v67 int32 = 2147483647
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                        _ = v67
                                                                                    }
                                                                                    v44 = v44
                                                                                    _ = v64
                                                                                }
                                                                                _ = v58
                                                                            }
                                                                        case 897:
                                                                            fmt.Println(fi14)
                                                                        }
                                                                    }
                                                                }
                                                                v68 := GenS5{F0: 255, F1: (v3 / ((1) | 1))}
                                                                _ = v68
                                                            }
                                                            _ = v44
                                                        }
                                                        _ = v43
                                                    }
                                                    _ = v41
                                                case 128:
                                                    var v69 int = (fi33 % ((fi33) | 1))
                                                    var v70 string = "foo bar"
                                                    if false {
                                                        var v71 bool = (true && (((!true) || (-32768 == -32768)) || true))
                                                        _ = v71
                                                    }
                                                    if ((!((v69 % ((fi33) | 1)) <= 540)) || (!false)) {
                                                        v3 += v3
                                                        v69 += -272
                                                        var v72 uint64 = 2
                                                        _ = v72
                                                    } else {
                                                        v70 = v70
                                                        _ = len(v13)
                                                        v73 := v1[((((v70 + "\n\t") + v70) + (v70 + "foo bar")) + v70)]
                                                        _ = v73
                                                    }
                                                    _ = v69
                                                    _ = v70
                                                }
                                                switch (fi33 % (((fi14 >> ((v3) & 63))) | 1)) {
                                                case -1:
                                                    v74 := GenS2{F0: (^(v3)), F1: 530, F2: (("\n\t" + ("" + "\n\t")) + "")}
                                                    _ = v74
                                                case 280:
                                                    v3 = (v3 >> ((v3) & 63))
                                                case 0:
                                                    _ = len(v13)
                                                default:
                                                    var v75 string = ("" + "hello")
                                                    _ = v75
                                                }
                                                v3 -= v3
                                                if true {
                                                    v3 += v3
                                                    v3 += (v3 / ((v3) | 1))
                                                }
                                            }
                                            var v76 uint16 = 255
                                            switch "\n\t" {
                                            case "\n\t":
                                                v76 += (v76 / ((2) | 1))
                                            case "a":
                                                var v77 uint8 = 1
                                                _ = v77
                                            case "foo bar":
                                                switch (fi33 % (((fi33 % ((fi14) | 1))) | 1)) {
                                                case 9223372036854775807:
                                                    var v78 uint32 = 4294967295
                                                    _ = v78
                                                default:
                                                    var v79 int32 = 0
                                                    _ = v79
                                                }
                                            }
                                            _ = v76
                                        case "\n\t":
                                            if false {
                                                fmt.Println(v3, fi14)
                                                v3 = v3
                                                var v80 uint = v3
                                                _ = v80
                                            } else {
                                                var v81 uint16 = 42
                                                _ = v81
                                            }
                                        case "":
                                            fmt.Println(v3, v3, v3)
                                        default:
                                            v13 = append(v13, fi33)
                                        }
                                    default:
                                        var v82 int32 = 2147483647
                                        _ = v82
                                    }
                                    _ = fi33
                                }
                                var v83 uint8 = 127
                                v84 := GenS8{F0: 42}
                                _ = v83
                                _ = v84
                            }
                        }
                        switch (("" + "\n\t") + "") {
                        case "hello":
                            var v85 int = 803
                            _ = v85
                        default:
                            var v86 uint32 = 2147483647
                            _ = v86
                        }
                    }
                }
            } else {
                var v87 int64 = 0
                var v88 uint8 = 42
                _ = v87
                _ = v88
            }
            v3 = 1
            _ = fi14
        }
    case "\n\t":
        v89 := GenS9{F0: -32768, F1: (-1 <= 100), F2: -128}
        _ = v89
    default:
        switch 32768 {
        case 127:
            var v90 string = "hello"
            _ = v90
        }
    }
    for fi91 := 0; fi91 < 7; fi91++ {
        var v92 int64 = 1
        _ = fi91
        _ = v92
    }
    v93 := GenS10{F0: 100, F1: 1000}
    for fi94 := 0; fi94 < 16; fi94++ {
        if (912 > 369) {
            var v95 string = ""
            switch (v95 + ("a" + (("" + v95) + ("hello" + "a")))) {
            case "\n\t":
                fmt.Println(fi94, v3, v3)
            case "foo bar":
                switch fi94 {
                case 4294967295:
                    fmt.Println(fi94, v3, v3)
                case -9223372036854775808:
                    for fi96 := 0; fi96 < 1; fi96++ {
                        switch (fi96 % ((fi96) | 1)) {
                        case -2147483648:
                            fmt.Println(v3, v95, fi96)
                        case 2147483647:
                            var v97 uint8 = 127
                            _ = v97
                        default:
                            switch fi96 {
                            case 343:
                                _ = len(v13)
                            default:
                                v95 = v95
                            }
                        }
                        var v98 int64 = -1
                        _ = fi96
                        _ = v98
                    }
                default:
                    var v99 int32 = 0
                    _ = v99
                }
            case "":
                if (-1 == 467) {
                    v100 := GenS10{F0: 95, F1: 1000}
                    v101 := GenS7{F0: 0}
                    if (0 > 0) {
                        if (!(true && true)) {
                            fmt.Println(v3, fi94)
                            fmt.Println(v95)
                        }
                        for fi102 := 0; fi102 < 7; fi102++ {
                            if (!true) {
                                v1["hello"] = true
                                if true {
                                    if true {
                                        fmt.Println(v95)
                                        if false {
                                            fmt.Println(v3)
                                            switch (fi102 ^ fi94) {
                                            case -800:
                                                v3 = 39
                                            case -2147483648:
                                                v13 = append(v13, fi102)
                                            }
                                            v3 += 42
                                        } else {
                                            switch "" {
                                            case "foo bar":
                                                if (495 == -1) {
                                                    var v103 uint16 = 188
                                                    switch v95 {
                                                    case "a":
                                                        _ = len(v13)
                                                    case "hello":
                                                        v95 = "hello"
                                                        if true {
                                                            v95 = v95
                                                        }
                                                    }
                                                    var v104 uint64 = 132
                                                    _ = v103
                                                    _ = v104
                                                } else {
                                                    if len(v13) >= 2 {
                                                        v13 = v13[1:len(v13)]
                                                    }
                                                    fmt.Println(fi102, fi102, fi94)
                                                    if len(v13) >= 2 {
                                                        v13 = v13[1:len(v13)]
                                                    }
                                                }
                                            case "hello":
                                                v105 := GenS5{F0: 33, F1: v3}
                                                _ = v105
                                            }
                                            if true {
                                                var v106 string = v95
                                                fmt.Println(v106, fi102, v106)
                                                var v107 int32 = -2147483648
                                                _ = v106
                                                _ = v107
                                            }
                                            fmt.Println(fi94, fi94)
                                            fmt.Println(v3)
                                        }
                                    } else {
                                        var v108 uint64 = 42
                                        if (128 != 118) {
                                            var v109 bool = false
                                            _ = v109
                                        } else {
                                            v108 = (v108 - (v108 / ((v108) | 1)))
                                            switch v95 {
                                            case "foo bar":
                                                if true {
                                                    v110 := GenS8{F0: 2147483647}
                                                    v111 := GenS11{F0: v108, F1: 237, F2: true}
                                                    v112 := GenS1{F0: ("a" + v95), F1: 255, F2: 1}
                                                    _ = v110
                                                    _ = v111
                                                    _ = v112
                                                } else {
                                                    var v113 int = fi102
                                                    v113 += 2147483647
                                                    _ = v113
                                                }
                                                var v114 int = 757
                                                if (713 < 0) {
                                                    var v115 int8 = 100
                                                    if len(v13) > 0 {
                                                        v13[0] = 90
                                                    }
                                                    v116 := GenS12{F0: (v108 + v108)}
                                                    _ = v115
                                                    _ = v116
                                                }
                                                _ = v114
                                            default:
                                                fmt.Println(v95, v3)
                                            }
                                        }
                                        _ = v108
                                    }
                                } else {
                                    fmt.Println(fi94, v95)
                                    var v117 uint16 = 371
                                    if (-2147483648 == fi94) {
                                        fmt.Println(v95, v117)
                                        if len(v13) > 0 {
                                            v13[0] = fi102
                                        }
                                    }
                                    _ = v117
                                }
                                switch "hello" {
                                case "":
                                    v118 := GenS6{F0: v3}
                                    _ = v118
                                case "hello":
                                    v95 = "\n\t"
                                case "a":
                                    var v119 int = (fi94 + fi94)
                                    _ = v119
                                default:
                                    switch v95 {
                                    case "a":
                                        var v120 string = ""
                                        _ = v120
                                    case "\n\t":
                                        v95 = v95
                                    case "foo bar":
                                        var v121 uint = (v3 >> ((v3) & 63))
                                        var v122 float64 = (1.5 - 1.5)
                                        _ = v121
                                        _ = v122
                                    }
                                }
                                var v123 int32 = 0
                                fmt.Println(fi94)
                                v13 = append(v13, fi94)
                                _ = v123
                            } else {
                                if ("a" == ("\n\t" + v95)) {
                                    switch (fi102 ^ 90) {
                                    case 128:
                                        var v124 float32 = 3.14159
                                        if (!false) {
                                            v125 := GenS13{F0: int64(fi94)}
                                            fmt.Println(v95, fi94, v95)
                                            _ = v125
                                        }
                                        _ = v124
                                    case 1:
                                        v3 = 1
                                    }
                                } else {
                                    var v126 bool = true
                                    var v127 int8 = -19
                                    v127 = (^(v127))
                                    _ = v126
                                    _ = v127
                                }
                                v13 = append(v13, 42)
                            }
                            _ = fi102
                        }
                        var v128 float32 = (0.0 + 0.1)
                        _ = v128
                    }
                    _ = v100
                    _ = v101
                } else {
                    v129 := GenS14{F0: 521}
                    _ = v129
                }
            default:
                v4[2147483647] = false
            }
            _ = v95
        }
        _ = fi94
    }
    if (!(!false)) {
        var v130 bool = true
        v130 = ((18446744073709551615 > 170) && v130)
        fmt.Println(v3, v3)
        _ = v130
    } else {
        var v131 string = "a"
        var v132 int8 = 45
        _ = v131
        _ = v132
    }
    switch "hello" {
    case "\n\t":
        v3 += 0
    case "":
        if false {
            fmt.Println(v3)
            fmt.Println(v3)
        }
    default:
        if false {
            v3 = v3
        }
    }
    v3 += (^(v3))
    v3 += (v3 * (v3 % (((v3 + v3)) | 1)))
    v133 := GenS15{F0: 687, F1: ""}
    if (((float64(-2.25) + float64(float32(3.14159))) - -2.25) <= (((-2.25 + (0.1 / 0.1)) - (3.14159 - 3.14159)) / -2.25)) {
        fmt.Println(v3)
        var v134 uint16 = 429
        _ = v134
    } else {
        v3 += v3
        v3 = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v133
    _ = v2
    _ = v3
    _ = v4
    _ = v9
    _ = v93
}
