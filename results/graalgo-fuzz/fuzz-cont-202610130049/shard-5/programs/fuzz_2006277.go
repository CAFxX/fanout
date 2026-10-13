package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int16; F1 int16; F2 int32 }
type GenS4 struct { F0 uint8; F1 uint; F2 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 uint64; F2 string }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int64; F1 int32; F2 uint64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 bool; F1 int }

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = -128
    v3 := []string{("a" + "foo bar"), (("" + "") + "\n\t")}
    var v4 int = 127
    v4 = v4
    fmt.Println(v2)
    if (true && true) {
        fmt.Println(v4, v4)
    }
    switch "" {
    case "foo bar":
        v5 := v1["foo bar"]
        _ = v5
    case "":
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := v1[""]
            _ = fi6
            _ = v7
        }
    case "hello":
        v4 = 0
    default:
        v2 = (v2 | (-(v2)))
    }
    if false {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 int32 = -849
            switch "\n\t" {
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "foo bar":
                if ((1e10 + 1e10) <= float64(3.14159)) {
                    var v10 int64 = 42
                    v11 := GenS1{F0: 1}
                    delete(v1, "hello")
                    _ = v10
                    _ = v11
                } else {
                    var v12 string = (("a" + ("" + "foo bar")) + "\n\t")
                    v13 := GenS2{F0: 18446744073709551615, F1: (2 == 18446744073709551615)}
                    if false {
                        var v14 float32 = 3.14159
                        switch ((v12 + ((v12 + (v12 + v12)) + v12)) + v12) {
                        case "foo bar":
                            switch (v4 - fi8) {
                            case 160:
                                for fi15 := 0; fi15 < 2; fi15++ {
                                    if true {
                                        delete(v1, "hello")
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v9 = v9
                                        v14 += float32(1.5)
                                    } else {
                                        var v16 bool = ((0.0 / (v14 / 3.14159)) != 3.14159)
                                        _ = v16
                                    }
                                    v14 += v14
                                    _ = fi15
                                }
                            default:
                                var v17 int32 = v9
                                _ = v17
                            }
                        case "":
                            v12 = (v12 + v12)
                        }
                        v1[("foo bar" + "\n\t")] = false
                        _ = v14
                    }
                    _ = v12
                    _ = v13
                }
            default:
                fmt.Println(v9, v2)
            }
            _ = fi8
            _ = v9
        }
        v2 = 105
    }
    if (!((v4 ^ (v4 - (v4 - -2147483648))) == v4)) {
        v18 := GenS1{F0: 0}
        v19 := GenS2{F0: 42, F1: true}
        var v20 int = (v4 >> ((1) & 63))
        _ = v18
        _ = v19
        _ = v20
    } else {
        _ = len(v3)
        var v21 int8 = v2
        v2 = (v2 ^ v2)
        _ = v21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 uint = 801
            var v25 uint64 = 0
            fmt.Println(v4, v2, v25)
            var v26 int64 = 560
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        v2 += v2
        v27 := GenS1{F0: 18446744073709551615}
        _ = fi22
        _ = v27
    }
    var v28 uint8 = 130
    var v29 int64 = -9223372036854775808
    if false {
        for fi30 := 0; fi30 < 16; fi30++ {
            for fi31 := 0; fi31 < 15; fi31++ {
                switch 568 {
                case 4294967295:
                    switch (fi31 & (fi31 * (-(v4)))) {
                    case 32768:
                        switch -2147483648 {
                        case 122:
                            v29 = (v29 | (v29 + 1))
                            v28 += v28
                            v2 += (v2 & v2)
                            if (true || (825 > 42)) {
                                v4 += v4
                            } else {
                                var v32 int8 = (^(v2))
                                _ = v32
                            }
                        case 152:
                            v33 := GenS3{F0: 32767, F1: -1, F2: -551}
                            _ = v33
                        }
                    case 423:
                        if (((-2.25 * float64(1e10)) <= 1.5) && false) {
                            if ((v2 ^ v2) < (v2 * (v2 * v2))) {
                                fmt.Println(v4, v4)
                                v34 := GenS4{F0: v28, F1: 65535, F2: "foo bar"}
                                v4 = (fi30 % (((fi30 | (v4 & (fi31 % (((^(fi30))) | 1))))) | 1))
                                v35 := GenS3{F0: 1, F1: 0, F2: int32(v4)}
                                switch fi31 {
                                case 4294967295:
                                    _ = len(v1)
                                    var v36 float32 = 1e10
                                    _ = v36
                                default:
                                    var v37 int32 = int32(v29)
                                    _ = v37
                                }
                                _ = v34
                                _ = v35
                            }
                        } else {
                            if (401 <= 1) {
                                var v38 float64 = float64(((3.14159 * ((float32(0.0) + 0.1) * float32(((float64((0.1 / (0.0 + float32(1e10)))) + 1.5) - -2.25)))) + ((-2.25 + 3.14159) / -2.25)))
                                switch fi30 {
                                case 2147483647:
                                    var v39 float64 = 1.5
                                    _ = v39
                                case 9223372036854775807:
                                    var v40 uint64 = 77
                                    _ = v40
                                case 127:
                                    _ = len(v1)
                                default:
                                    switch (("" + "") + "a") {
                                    case "a":
                                        switch "" {
                                        case "\n\t":
                                            if true {
                                                if true {
                                                    var v41 int64 = v29
                                                    _ = v41
                                                } else {
                                                    var v42 int8 = v2
                                                    if (false || ((!false) && (-128 <= (v42 / ((-75) | 1))))) {
                                                        v29 = (v29 * int64(v4))
                                                        var v43 uint8 = (v28 / ((v28) | 1))
                                                        _ = v43
                                                    } else {
                                                        if true {
                                                            switch fi31 {
                                                            case 32767:
                                                                var v44 int32 = 943
                                                                _ = v44
                                                            default:
                                                                if (!false) {
                                                                    v38 = 0.0
                                                                    var v45 int16 = 1000
                                                                    _ = v45
                                                                } else {
                                                                    v46 := v1["a"]
                                                                    var v47 uint16 = 32767
                                                                    _ = v46
                                                                    _ = v47
                                                                }
                                                            }
                                                            fmt.Println(v28)
                                                            if len(v3) > 0 {
                                                                v3[0] = "a"
                                                            }
                                                        } else {
                                                            switch "" {
                                                            case "hello":
                                                                var v48 int = v4
                                                                _ = v48
                                                            case "foo bar":
                                                                var v49 uint8 = 1
                                                                _ = v49
                                                            case "a":
                                                                if true {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    v29 += v29
                                                                    v42 = (v42 - (v42 / (((v42 >> ((519) & 63))) | 1)))
                                                                    if true {
                                                                        if (4294967295 != 2) {
                                                                            v50 := GenS1{F0: 42}
                                                                            _ = v50
                                                                        }
                                                                    } else {
                                                                        _ = len(v1)
                                                                    }
                                                                    var v51 uint8 = 127
                                                                    var v52 uint = 65535
                                                                    switch "" {
                                                                    case "":
                                                                        if len(v3) > 0 {
                                                                            v3[0] = "\n\t"
                                                                        }
                                                                    case "a":
                                                                        var v53 uint8 = v28
                                                                        _ = v53
                                                                    default:
                                                                        var v54 uint16 = 541
                                                                        _ = v54
                                                                    }
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                            }
                                                            var v55 string = ((("hello" + ((((("hello" + ("a" + "\n\t")) + ("foo bar" + ("a" + ("" + "foo bar")))) + "hello") + "hello") + "a")) + "a") + "\n\t")
                                                            _ = v55
                                                        }
                                                        if false {
                                                            switch fi30 {
                                                            case 9223372036854775807:
                                                                var v56 int64 = (-(v29))
                                                                _ = v56
                                                            case 2147483647:
                                                                fmt.Println(v28, v29, v2)
                                                            case -9223372036854775808:
                                                                switch fi30 {
                                                                case 9223372036854775807:
                                                                    if (0 == 42) {
                                                                        var v57 float32 = (0.0 + (3.14159 - 1.5))
                                                                        _ = v57
                                                                    }
                                                                }
                                                            default:
                                                                var v58 uint8 = 255
                                                                _ = v58
                                                            }
                                                        }
                                                    }
                                                    if (true || false) {
                                                        v2 -= (^(v42))
                                                        switch "" {
                                                        case "foo bar":
                                                            v38 = ((v38 / ((v38 * ((float64(1e10) / v38) * v38)) + 3.14159)) * (float64(0.1) * v38))
                                                        case "hello":
                                                            if (!(!(42 >= 1))) {
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    if ((false || true) && true) {
                                                                        v59 := GenS5{F0: v42}
                                                                        _ = v59
                                                                    }
                                                                case "\n\t":
                                                                    fmt.Println(v38, fi30)
                                                                case "":
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (("" + "foo bar") + "\n\t")
                                                                    }
                                                                default:
                                                                    if true {
                                                                        switch -338 {
                                                                        case 128:
                                                                            var v60 int16 = 1000
                                                                            _ = v60
                                                                        case 9223372036854775807:
                                                                            var v61 int = 376
                                                                            _ = v61
                                                                        }
                                                                        switch "" {
                                                                        case "foo bar":
                                                                            var v62 int = int(v2)
                                                                            _ = v62
                                                                        case "\n\t":
                                                                            var v63 float32 = 0.0
                                                                            _ = v63
                                                                        case "a":
                                                                            delete(v1, (("\n\t" + "\n\t") + (("\n\t" + "a") + "a")))
                                                                        default:
                                                                            v42 = 1
                                                                        }
                                                                        switch (fi30 - 128) {
                                                                        case -9223372036854775808:
                                                                            if true {
                                                                                if ((262 < -589) && (255 <= 32767)) {
                                                                                    v64 := v1[("\n\t" + ("hello" + ((("\n\t" + (("" + ("" + "hello")) + "\n\t")) + ("" + "foo bar")) + ((("" + "\n\t") + (("hello" + "foo bar") + "")) + ("foo bar" + "")))))]
                                                                                    fmt.Println(fi30, fi30)
                                                                                    fmt.Println(v64, fi31)
                                                                                    var v65 uint16 = 168
                                                                                    _ = v64
                                                                                    _ = v65
                                                                                } else {
                                                                                    var v66 bool = false
                                                                                    v29 = (v29 + v29)
                                                                                    if v66 {
                                                                                        v67 := GenS1{F0: 243}
                                                                                        _ = len(v3)
                                                                                        _ = v67
                                                                                    } else {
                                                                                        var v68 uint16 = 674
                                                                                        v2 = v2
                                                                                        _ = v68
                                                                                    }
                                                                                    _ = v66
                                                                                }
                                                                                delete(v1, "a")
                                                                                if false {
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = ("hello" + "hello")
                                                                                    }
                                                                                    switch fi30 {
                                                                                    case 1:
                                                                                        var v69 int32 = -89
                                                                                        _ = v69
                                                                                    case 4294967295:
                                                                                        v38 -= v38
                                                                                    case -156:
                                                                                        var v70 uint8 = v28
                                                                                        _ = v70
                                                                                    }
                                                                                    v2 -= v2
                                                                                }
                                                                            } else {
                                                                                var v71 uint8 = v28
                                                                                _ = v71
                                                                            }
                                                                        case -357:
                                                                            v72 := GenS6{F0: fi31, F1: 42, F2: ("hello" + "hello")}
                                                                            _ = v72
                                                                        }
                                                                    }
                                                                }
                                                                fmt.Println(v29, fi31, fi31)
                                                            } else {
                                                                delete(v1, (("a" + "") + ((("\n\t" + "hello") + (("a" + ("hello" + "")) + "\n\t")) + "foo bar")))
                                                                v73 := v1[("hello" + "\n\t")]
                                                                v73 = (!(v73 || true))
                                                                switch fi30 {
                                                                case 128:
                                                                    v74 := GenS3{F0: 1, F1: 0, F2: 2147483647}
                                                                    _ = v74
                                                                case 32768:
                                                                    v75 := GenS1{F0: 0}
                                                                    _ = v75
                                                                case -1:
                                                                    switch 128 {
                                                                    case 4294967295:
                                                                        if v73 {
                                                                            v76 := v1[("a" + (("" + "") + "foo bar"))]
                                                                            _ = v76
                                                                        } else {
                                                                            v28 = v28
                                                                            v29 = (v29 - v29)
                                                                            if true {
                                                                                delete(v1, "")
                                                                                v42 = v42
                                                                                var v77 uint = 255
                                                                                switch 195 {
                                                                                case 42:
                                                                                    v77 = (v77 >> ((v77) & 63))
                                                                                default:
                                                                                    fmt.Println(v42, v29, v28)
                                                                                }
                                                                                _ = v77
                                                                            }
                                                                            v78 := GenS7{F0: 0}
                                                                            _ = v78
                                                                        }
                                                                    default:
                                                                        v4 = (fi30 / ((v4) | 1))
                                                                    }
                                                                default:
                                                                    v28 -= v28
                                                                }
                                                                var v79 uint32 = 4294967295
                                                                _ = v73
                                                                _ = v79
                                                            }
                                                        }
                                                        if (!(!true)) {
                                                            var v80 int64 = v29
                                                            if (32767 <= fi31) {
                                                                if (!((v28 | v28) < v28)) {
                                                                    fmt.Println(v42)
                                                                }
                                                            } else {
                                                                var v81 string = "foo bar"
                                                                if (!(((("\n\t" + v81) + "foo bar") + ((("" + ((v81 + v81) + "a")) + "\n\t") + (v81 + (v81 + ("hello" + v81))))) >= v81)) {
                                                                    var v82 uint64 = 146
                                                                    v42 -= -1
                                                                    fmt.Println(v82, v80)
                                                                    _ = v82
                                                                }
                                                                var v83 float64 = (v38 - v38)
                                                                _ = v81
                                                                _ = v83
                                                            }
                                                            fmt.Println(fi30)
                                                            _ = v80
                                                        } else {
                                                            v29 += v29
                                                            v2 += v42
                                                            v29 = -781
                                                        }
                                                        v2 -= (v2 / (((v2 ^ (v2 & (^(v42))))) | 1))
                                                    }
                                                    _ = v42
                                                }
                                                var v84 uint16 = 788
                                                var v85 uint64 = 1
                                                _ = v84
                                                _ = v85
                                            }
                                        case "hello":
                                            v86 := GenS3{F0: -873, F1: -32768, F2: 1}
                                            _ = v86
                                        default:
                                            switch (-(fi31)) {
                                            case -36:
                                                delete(v1, "")
                                            default:
                                                v87 := GenS5{F0: (-(v2))}
                                                fmt.Println(v28, v29)
                                                _ = v87
                                            }
                                        }
                                    default:
                                        var v88 uint = 0
                                        _ = v88
                                    }
                                }
                                _ = v38
                            } else {
                                v29 = (v29 & (v29 % ((v29) | 1)))
                                var v89 uint32 = 65535
                                _ = v89
                            }
                        }
                    default:
                        if true {
                            if (360 == 2) {
                                if false {
                                    v28 = (v28 | (v28 / ((0) | 1)))
                                }
                                v2 += 70
                            }
                            v2 += (-(v2))
                            var v90 float32 = -2.25
                            _ = v90
                        } else {
                            v3 = append(v3, ("" + "a"))
                            _ = len(v1)
                            var v91 bool = ((!true) || true)
                            _ = v91
                        }
                    }
                case 128:
                    v28 += v28
                }
                v29 = 42
                _ = fi31
            }
            fmt.Println(v28)
            if false {
                v28 += v28
            } else {
                switch "hello" {
                case "foo bar":
                    if (996 < 42) {
                        if (true && true) {
                            var v92 int = v4
                            v93 := GenS8{F0: -631}
                            _ = len(v3)
                            _ = v92
                            _ = v93
                        } else {
                            v94 := GenS1{F0: 2}
                            var v95 int16 = -686
                            _ = v94
                            _ = v95
                        }
                        for fi96 := 0; fi96 < 19; fi96++ {
                            var v97 float64 = 3.14159
                            switch (v4 % (((fi30 - (fi30 + fi96))) | 1)) {
                            case 9223372036854775807:
                                fmt.Println(v29, v29, v29)
                            case -2147483648:
                                var v98 int64 = -417
                                _ = v98
                            case 128:
                                v28 = v28
                            default:
                                v99 := v1[("" + "")]
                                _ = v99
                            }
                            if (0 <= v29) {
                                v28 = v28
                                v2 += 127
                                if (0 >= 302) {
                                    if true {
                                        v100 := GenS6{F0: 42, F1: 2, F2: "a"}
                                        _ = v100
                                    }
                                }
                                var v101 uint = 753
                                _ = v101
                            } else {
                                v29 += v29
                                v102 := v1[""]
                                fmt.Println(v4, v97)
                                _ = v102
                            }
                            _ = fi96
                            _ = v97
                        }
                    }
                default:
                    switch ("\n\t" + "a") {
                    case "a":
                        fmt.Println(v29, fi30)
                    default:
                        if (2 > 18446744073709551615) {
                            v28 += v28
                        } else {
                            for fi103 := 0; fi103 < 17; fi103++ {
                                var v104 int16 = 42
                                delete(v1, (("foo bar" + "") + "hello"))
                                _ = fi103
                                _ = v104
                            }
                            v2 = 127
                        }
                    }
                }
                for fi105 := 0; fi105 < 9; fi105++ {
                    v106 := GenS9{F0: -1, F1: 1, F2: 343}
                    _ = fi105
                    _ = v106
                }
                var v107 bool = (255 == 2)
                _ = v107
            }
            _ = fi30
        }
        if (v28 > v28) {
            v108 := GenS5{F0: v2}
            fmt.Println(v29)
            if (42 > 2147483647) {
                if true {
                    if ((v4 << ((869) & 63)) > 32767) {
                        for fi109 := 0; fi109 < 20; fi109++ {
                            switch ("\n\t" + "\n\t") {
                            case "hello":
                                v110 := GenS10{F0: false}
                                _ = v110
                            case "a":
                                delete(v1, ("" + (("a" + "a") + "\n\t")))
                            case "":
                                fmt.Println(v29, v2, v28)
                            default:
                                switch "foo bar" {
                                case "\n\t":
                                    fmt.Println(fi109, v28, v28)
                                default:
                                    var v111 float64 = (3.14159 / 0.1)
                                    _ = v111
                                }
                            }
                            _ = fi109
                        }
                        var v112 uint = 65535
                        _ = v112
                    } else {
                        for fi113 := 0; fi113 < 12; fi113++ {
                            for fi114 := 0; fi114 < 5; fi114++ {
                                v29 -= v29
                                v29 += (^(v29))
                                v115 := GenS1{F0: 4294967295}
                                _ = fi114
                                _ = v115
                            }
                            fmt.Println(v4, v29, v28)
                            v116 := GenS8{F0: 1}
                            _ = fi113
                            _ = v116
                        }
                    }
                } else {
                    v28 += (v28 | (^(v28)))
                    var v117 int = v4
                    _ = v117
                }
                v118 := GenS11{F0: (true && ((304 != 0) || (true || (!(70 <= 259))))), F1: 140}
                v29 = v29
                _ = v118
            }
            _ = v108
        } else {
            fmt.Println(v4, v4)
            switch ("\n\t" + (("" + "\n\t") + "")) {
            case "\n\t":
                fmt.Println(v2, v4, v4)
            case "":
                _ = len(v1)
            case "hello":
                v119 := v1["hello"]
                _ = v119
            }
        }
    }
    for fi120 := 0; fi120 < 11; fi120++ {
        var v121 int64 = -1
        var v122 uint = 597
        v123 := GenS5{F0: v2}
        _ = fi120
        _ = v121
        _ = v122
        _ = v123
    }
    for fi124 := 0; fi124 < 9; fi124++ {
        for fi125 := 0; fi125 < 3; fi125++ {
            var v126 int64 = -1
            var v127 string = "\n\t"
            _ = fi125
            _ = v126
            _ = v127
        }
        if true {
            for fi128 := 0; fi128 < 3; fi128++ {
                v2 = (v2 / (((v2 + v2)) | 1))
                switch 42 {
                case 32767:
                    if (0 != 195) {
                        v28 -= v28
                        v2 = v2
                    }
                    v3 = append(v3, "foo bar")
                    fmt.Println(v29, v4, v28)
                }
                _ = fi128
            }
            switch (-(v4)) {
            case 9223372036854775807:
                var v129 string = ("hello" + ((((("\n\t" + "a") + ("\n\t" + "foo bar")) + (((("\n\t" + "a") + "a") + ("\n\t" + "\n\t")) + ("\n\t" + "a"))) + "a") + "\n\t"))
                _ = v129
            default:
                v2 = (-(v2))
            }
        } else {
            v29 -= (v29 * -9223372036854775808)
            for fi130 := 0; fi130 < 18; fi130++ {
                v28 += (v28 >> ((255) & 63))
                v3 = append(v3, "hello")
                _ = fi130
            }
        }
        v131 := v1["foo bar"]
        _ = fi124
        _ = v131
    }
    v29 = -340
    if false {
        v29 += v29
        if (1000 < -942) {
            var v132 int64 = v29
            var v133 float32 = 1e10
            fmt.Println(v132, v4)
            _ = v132
            _ = v133
        } else {
            if false {
                if true {
                    for fi134 := 0; fi134 < 9; fi134++ {
                        v29 -= 9223372036854775807
                        _ = fi134
                    }
                    _ = len(v3)
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v135 uint = 1
                    switch "\n\t" {
                    case "\n\t":
                        if true {
                            v136 := v1["foo bar"]
                            if ((-(v4)) <= 4294967295) {
                                var v137 int32 = 394
                                _ = v137
                            }
                            if (3.14159 > -2.25) {
                                var v138 float64 = -2.25
                                _ = v138
                            }
                            _ = v136
                        }
                    case "":
                        v4 += (v4 - (v4 + 1))
                    default:
                        fmt.Println(v4, v135, v4)
                    }
                    _ = v135
                }
                for fi139 := 0; fi139 < 7; fi139++ {
                    var v140 bool = (42 >= 2147483648)
                    _ = fi139
                    _ = v140
                }
            }
            var v141 string = ("" + ("\n\t" + "a"))
            var v142 int64 = v29
            _ = v141
            _ = v142
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
