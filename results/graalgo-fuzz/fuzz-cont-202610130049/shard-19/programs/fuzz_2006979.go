package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int8; F2 string }
type GenS4 struct { F0 int8; F1 string; F2 int }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 uint16; F1 uint32; F2 bool }
type GenS7 struct { F0 string; F1 uint32 }
type GenS8 struct { F0 int; F1 uint }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint8; F1 int16 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint8; F1 uint64 }
type GenS13 struct { F0 int8; F1 uint16; F2 int64 }
type GenS14 struct { F0 uint8; F1 int; F2 uint }
type GenS15 struct { F0 int16; F1 int8; F2 int64 }
func genFn0() uint8 {
    switch "a" {
    case "":
        var v4 uint64 = 18446744073709551615
        _ = v4
    case "foo bar":
        v5 := []int{-326, 928, -641}
        _ = v5
    case "a":
        var v6 uint16 = 2
        _ = v6
    }
    var v7 float32 = -2.25
    v8 := []int{}
    _ = v7
    _ = v8
    return 72
}

func main() {
    var v1 float32 = (1e10 / (((((3.14159 / 1e10) * float32(0.1)) - ((1.5 + 0.1) - (3.14159 * (float32(1.5) + 3.14159)))) - (0.1 * ((1e10 / 1.0) - (-2.25 + 1.5)))) * (0.1 * 3.14159)))
    var v2 int32 = -131
    v3 := []string{"foo bar", ("hello" + "foo bar")}
    if false {
        v2 = -1
        if false {
            v9 := GenS1{F0: 2, F1: 4294967295, F2: -256}
            _ = v9
        }
        fmt.Println(v1, v1)
    } else {
        switch 0 {
        case 0:
            if true {
                if (0 <= 1) {
                    v10 := GenS1{F0: 381, F1: 2, F2: -1}
                    v11 := GenS1{F0: 762, F1: 18446744073709551615, F2: -32768}
                    _ = v10
                    _ = v11
                } else {
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v1 += 1e10
                    }
                    switch ("hello" + "a") {
                    case "foo bar":
                        v1 = (v1 + float32(-2.25))
                    case "":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "\n\t":
                        v2 = v2
                    default:
                        switch 4294967295 {
                        case -790:
                            if (326 <= 2147483648) {
                                if (!(128 == 1)) {
                                    v1 -= -2.25
                                    v2 -= (v2 & (^(v2)))
                                    v1 += -2.25
                                } else {
                                    fmt.Println(v1)
                                }
                            } else {
                                var v12 float64 = -2.25
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = ""
                                    }
                                } else {
                                    var v13 int = 850
                                    _ = v13
                                }
                                _ = v12
                            }
                        case 32767:
                            v1 += 3.14159
                        case 391:
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v2, v2)
                            case "\n\t":
                                if len(v3) > 0 {
                                    v3[0] = ("\n\t" + "foo bar")
                                }
                            }
                        }
                    }
                    for fi14 := 0; fi14 < 11; fi14++ {
                        switch "a" {
                        case "\n\t":
                            if (943 == 175) {
                                var v15 int = fi14
                                _ = v15
                            }
                        default:
                            v3 = append(v3, ("a" + (("foo bar" + "\n\t") + "hello")))
                        }
                        if (!(255 >= 1)) {
                            if len(v3) > 0 {
                                v3[0] = ("hello" + "foo bar")
                            }
                            if false {
                                v1 = v1
                                for fi16 := 0; fi16 < 14; fi16++ {
                                    var v17 int16 = 32767
                                    var v18 uint64 = 4294967295
                                    if ((fi16 >> ((42) & 63)) < -9223372036854775808) {
                                        if (false && false) {
                                            if len(v3) > 0 {
                                                v3[0] = ("foo bar" + "hello")
                                            }
                                        } else {
                                            fmt.Println(v17, fi14)
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v19 int8 = 1
                                                _ = v19
                                            }
                                            fmt.Println(fi14)
                                        }
                                        var v20 int32 = v2
                                        v20 -= v20
                                        _ = v20
                                    }
                                    var v21 uint = 0
                                    fmt.Println(v1, v2)
                                    v1 -= v1
                                    switch "hello" {
                                    case "foo bar":
                                        v22 := GenS2{F0: v18}
                                        _ = len(v3)
                                        _ = v22
                                    default:
                                        fmt.Println(v2, v21, v17)
                                    }
                                    _ = fi16
                                    _ = v17
                                    _ = v18
                                    _ = v21
                                }
                            }
                            v1 = (3.14159 * (-2.25 + v1))
                        } else {
                            v23 := GenS3{F0: 100, F1: 0, F2: "\n\t"}
                            if true {
                                v24 := GenS4{F0: -18, F1: "hello", F2: -832}
                                for fi25 := 0; fi25 < 17; fi25++ {
                                    if (v2 <= 42) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        fmt.Println(fi25, fi14)
                                        fmt.Println(v2, v1)
                                        v2 += (v2 << ((832) & 63))
                                    }
                                    v26 := GenS1{F0: 639, F1: 472, F2: 1000}
                                    v3 = append(v3, "foo bar")
                                    _ = fi25
                                    _ = v26
                                }
                                _ = v24
                            }
                            _ = v23
                        }
                        var v27 uint32 = 633
                        _ = fi14
                        _ = v27
                    }
                }
                v2 += (v2 >> ((733) & 63))
                fmt.Println(v1)
            } else {
                v28 := GenS5{F0: 215, F1: true}
                v2 += v2
                for fi29 := 0; fi29 < 18; fi29++ {
                    v2 += (v2 * -2147483648)
                    v1 = (v1 + float32(1e10))
                    _ = fi29
                }
                _ = v28
            }
        case 42:
            v30 := GenS3{F0: -128, F1: 127, F2: "hello"}
            _ = v30
        default:
            v31 := GenS4{F0: 42, F1: "\n\t", F2: int(v2)}
            _ = v31
        }
        var v32 int16 = -32768
        switch 4294967295 {
        case -210:
            switch 953 {
            case -945:
                fmt.Println(v1, v32)
            case 673:
                v32 = 1000
            case 646:
                v1 = v1
            default:
                if ((2 >= 1) && false) {
                    var v33 uint64 = 0
                    switch 42 {
                    case 32768:
                        var v34 uint16 = 2
                        _ = v34
                    }
                    _ = v33
                } else {
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                    if false {
                        if (42 > -25) {
                            var v35 float64 = 3.14159
                            _ = v35
                        }
                    }
                }
            }
        case 1:
            if false {
                for fi36 := 0; fi36 < 17; fi36++ {
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    v37 := GenS6{F0: 88, F1: 2147483647, F2: true}
                    var v38 int64 = 9223372036854775807
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                v39 := GenS7{F0: "", F1: 42}
                v40 := GenS1{F0: 4294967295, F1: 733, F2: (v32 >> ((1) & 63))}
                _ = v39
                _ = v40
            }
        }
        _ = v32
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        v3 = append(v3, "foo bar")
        v42 := GenS8{F0: 1, F1: 4294967295}
        if false {
            v43 := GenS5{F0: 391, F1: false}
            _ = v43
        }
        _ = fi41
        _ = v42
    }
    if (18446744073709551615 <= 18446744073709551615) {
        v44 := GenS9{F0: 859}
        v2 = 42
        v2 = v2
        _ = v44
    }
    if (9223372036854775807 != 9223372036854775807) {
        if (false || true) {
            v2 = v2
            if false {
                for fi45 := 0; fi45 < 9; fi45++ {
                    if ((true && (18446744073709551615 >= 42)) || (1 == 100)) {
                        var v46 int8 = 100
                        switch (fi45 << ((42) & 63)) {
                        case 128:
                            var v47 int64 = 9223372036854775807
                            _ = v47
                        case -278:
                            var v48 string = "hello"
                            _ = v48
                        default:
                            var v49 uint64 = 18446744073709551615
                            _ = v49
                        }
                        _ = v46
                    }
                    _ = fi45
                }
                var v50 int32 = (v2 - (v2 & (v2 >> ((329) & 63))))
                _ = v50
            } else {
                v51 := GenS9{F0: 42}
                for fi52 := 0; fi52 < 13; fi52++ {
                    var v53 uint64 = 18446744073709551615
                    v53 -= (v53 - 837)
                    _ = fi52
                    _ = v53
                }
                if ((714 <= 32767) || true) {
                    v1 -= v1
                    var v54 int64 = 813
                    _ = v54
                } else {
                    _ = len(v3)
                }
                _ = v51
            }
        } else {
            v1 -= v1
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            for fi55 := 0; fi55 < 5; fi55++ {
                if (491 >= 978) {
                    _ = len(v3)
                    v56 := GenS5{F0: 1, F1: (false && (331 != 65535))}
                    _ = v56
                } else {
                    if false {
                        v2 = v2
                        switch "hello" {
                        case "hello":
                            var v57 uint8 = 255
                            _ = v57
                        case "a":
                            var v58 bool = (false && true)
                            _ = v58
                        case "foo bar":
                            fmt.Println(v2, fi55, fi55)
                        }
                    } else {
                        if (!true) {
                            fmt.Println(fi55)
                            v59 := GenS9{F0: 4294967295}
                            _ = v59
                        } else {
                            if (!true) {
                                v1 += v1
                            } else {
                                fmt.Println(v1, fi55, fi55)
                            }
                        }
                    }
                    v3 = append(v3, (("" + ((("hello" + ("" + "hello")) + "foo bar") + ("\n\t" + (("\n\t" + "hello") + "hello")))) + "a"))
                }
                var v60 uint8 = 113
                _ = fi55
                _ = v60
            }
        }
    }
    switch ("" + "hello") {
    case "":
        switch "hello" {
        case "hello":
            for fi61 := 0; fi61 < 3; fi61++ {
                v2 = (v2 * v2)
                _ = fi61
            }
        case "a":
            fmt.Println(v2)
        default:
            var v62 uint16 = 255
            _ = v62
        }
    case "\n\t":
        v2 -= v2
    }
    if true {
        v63 := GenS7{F0: "", F1: 507}
        _ = v63
    }
    for fi64 := 0; fi64 < 6; fi64++ {
        switch (fi64 - 127) {
        case -9223372036854775808:
            if (9 > 127) {
                _ = len(v3)
                v3 = append(v3, "\n\t")
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                var v65 string = ("hello" + "a")
                switch ((((((v65 + "") + v65) + "a") + v65) + ("a" + ((((v65 + ("a" + (("\n\t" + (("\n\t" + v65) + v65)) + ("\n\t" + (("foo bar" + (("\n\t" + v65) + v65)) + ""))))) + "foo bar") + (v65 + (v65 + "a"))) + (v65 + v65)))) + v65) {
                case "foo bar":
                    v66 := GenS6{F0: 0, F1: 1, F2: false}
                    _ = v66
                case "hello":
                    for fi67 := 0; fi67 < 10; fi67++ {
                        switch (v65 + "foo bar") {
                        case "":
                            v1 += (float32(1.5) - 3.14159)
                        case "a":
                            var v68 int = (^(fi67))
                            _ = v68
                        case "\n\t":
                            fmt.Println(v2, v65, v65)
                        default:
                            var v69 string = v65
                            _ = v69
                        }
                        if true {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v70 := GenS4{F0: 1, F1: v65, F2: 1000}
                            _ = v70
                        }
                        _ = fi67
                    }
                default:
                    for fi71 := 0; fi71 < 1; fi71++ {
                        v72 := GenS10{F0: 128, F1: 1000}
                        fmt.Println(fi71, fi64, v65)
                        if true {
                            if false {
                                switch fi71 {
                                case 1:
                                    if (((!(620 != 0)) || false) && (!(352 <= 1))) {
                                        v65 = (v65 + v65)
                                        switch "hello" {
                                        case "a":
                                            fmt.Println(v65, v1, fi64)
                                        case "":
                                            var v73 uint16 = 694
                                            _ = v73
                                        default:
                                            var v74 int16 = 32767
                                            _ = v74
                                        }
                                    }
                                case -2147483648:
                                    var v75 uint16 = 0
                                    _ = v75
                                default:
                                    switch 32768 {
                                    case -208:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case 0:
                                        v65 = ""
                                    }
                                }
                            }
                        } else {
                            if (742 <= 42) {
                                var v76 int16 = -32768
                                if false {
                                    var v77 int16 = (v76 << ((599) & 63))
                                    if len(v3) > 0 {
                                        v3[0] = (v65 + v65)
                                    }
                                    _ = v77
                                } else {
                                    v65 = "hello"
                                }
                                v78 := GenS7{F0: v65, F1: 4294967295}
                                _ = v76
                                _ = v78
                            }
                        }
                        _ = fi71
                        _ = v72
                    }
                }
                _ = len(v3)
                _ = v65
            }
        case 731:
            v79 := GenS5{F0: 2147483647, F1: (true && true)}
            _ = v79
        default:
            if ((v2 & v2) < (v2 + v2)) {
                _ = len(v3)
                v1 += float32(-2.25)
            } else {
                v2 -= (v2 % (((v2 - v2)) | 1))
                for fi80 := 0; fi80 < 16; fi80++ {
                    var v81 uint8 = 255
                    _ = fi80
                    _ = v81
                }
            }
        }
        fmt.Println(fi64)
        v2 = (v2 % (((v2 / (((v2 ^ v2)) | 1))) | 1))
        _ = fi64
    }
    if (false && (42 == 42)) {
        fmt.Println(v1)
        var v82 int32 = -834
        _ = v82
    } else {
        var v83 int32 = 42
        v84 := GenS11{F0: 4294967295}
        _ = v83
        _ = v84
    }
    _ = len(v3)
    switch -2147483648 {
    case 9223372036854775807:
        v85 := GenS7{F0: ("a" + "hello"), F1: 207}
        _ = v85
    case 32767:
        var v86 int32 = -237
        _ = v86
    case -694:
        for fi87 := 0; fi87 < 5; fi87++ {
            v2 -= v2
            v2 = v2
            v88 := GenS11{F0: 2}
            _ = fi87
            _ = v88
        }
    }
    v89 := GenS11{F0: 790}
    v90 := GenS12{F0: 2, F1: 2}
    var v91 uint16 = 2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if ((v2 << ((2) & 63)) == v2) {
        var v92 string = ""
        _ = v92
    } else {
        v93 := GenS4{F0: -1, F1: "\n\t", F2: 2147483647}
        if true {
            switch (("\n\t" + ("foo bar" + ((("hello" + ("foo bar" + "a")) + ("hello" + "a")) + "foo bar"))) + "") {
            case "hello":
                for fi94 := 0; fi94 < 1; fi94++ {
                    var v95 bool = (-43 > 18)
                    _ = fi94
                    _ = v95
                }
            case "":
                if len(v3) > 0 {
                    v3[0] = ("hello" + ("hello" + ("" + (((("foo bar" + ("" + "foo bar")) + "") + "") + ("hello" + "")))))
                }
            case "foo bar":
                switch 2147483647 {
                case 155:
                    for fi96 := 0; fi96 < 7; fi96++ {
                        var v97 uint16 = 65535
                        _ = fi96
                        _ = v97
                    }
                default:
                    for fi98 := 0; fi98 < 19; fi98++ {
                        v99 := GenS13{F0: 0, F1: v91, F2: 42}
                        switch (("foo bar" + "a") + "") {
                        case "":
                            if false {
                                switch (-(fi98)) {
                                case -9223372036854775808:
                                    _ = len(v3)
                                default:
                                    for fi100 := 0; fi100 < 9; fi100++ {
                                        var v101 int = int(v2)
                                        _ = fi100
                                        _ = v101
                                    }
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        case "a":
                            for fi102 := 0; fi102 < 4; fi102++ {
                                v1 += (v1 * 0.1)
                                v1 -= v1
                                fmt.Println(fi102, v2, v2)
                                _ = fi102
                            }
                        case "foo bar":
                            v103 := GenS12{F0: genFn0(), F1: 50}
                            _ = v103
                        }
                        _ = fi98
                        _ = v99
                    }
                }
            default:
                v2 = 1
            }
        } else {
            if true {
                v104 := GenS8{F0: 311, F1: 923}
                fmt.Println(v91)
                _ = v104
            }
            var v105 float32 = float32(-2.25)
            _ = v105
        }
        var v106 uint = 2
        _ = v106
        _ = v93
    }
    v107 := GenS6{F0: (v91 >> ((255) & 63)), F1: 0, F2: ((v91 * v91) < 352)}
    if ((-830 != 62) || (-58 != 41)) {
        fmt.Println(v91, v1)
    } else {
        switch "a" {
        case "\n\t":
            for fi108 := 0; fi108 < 18; fi108++ {
                switch fi108 {
                case 1:
                    if (fi108 == (fi108 / ((fi108) | 1))) {
                        if len(v3) > 0 {
                            v3[0] = ("a" + ("a" + (("a" + "\n\t") + "")))
                        }
                        v1 += v1
                        v1 -= ((0.0 * (v1 - v1)) / v1)
                    } else {
                        _ = len(v3)
                        var v109 uint32 = 197
                        v2 = v2
                        _ = v109
                    }
                case 4294967295:
                    v110 := GenS14{F0: 29, F1: 764, F2: 1}
                    _ = v110
                case -2147483648:
                    v91 += v91
                }
                v1 = v1
                _ = fi108
            }
        }
        for fi111 := 0; fi111 < 7; fi111++ {
            if false {
                v3 = append(v3, ("\n\t" + "a"))
                var v112 int32 = int32(fi111)
                if false {
                    for fi113 := 0; fi113 < 12; fi113++ {
                        var v114 uint32 = 654
                        var v115 uint64 = 18446744073709551615
                        _ = fi113
                        _ = v114
                        _ = v115
                    }
                } else {
                    v116 := GenS9{F0: 704}
                    for fi117 := 0; fi117 < 2; fi117++ {
                        v118 := GenS11{F0: 42}
                        switch "a" {
                        case "":
                            var v119 float32 = v1
                            _ = v119
                        default:
                            var v120 int8 = -1
                            _ = v120
                        }
                        _ = fi117
                        _ = v118
                    }
                    v2 = v2
                    _ = v116
                }
                _ = v112
            } else {
                var v121 float64 = 0.0
                v122 := GenS2{F0: 635}
                switch fi111 {
                case -9223372036854775808:
                    v3 = append(v3, "hello")
                case -2147483648:
                    var v123 uint64 = 58
                    _ = v123
                case 9223372036854775807:
                    if true {
                        if true {
                            v2 -= 42
                            var v124 bool = (42 == 9223372036854775807)
                            _ = v124
                        } else {
                            switch fi111 {
                            case 42:
                                v125 := GenS15{F0: -1, F1: -78, F2: 9223372036854775807}
                                _ = v125
                            }
                        }
                        for fi126 := 0; fi126 < 14; fi126++ {
                            v91 += v91
                            _ = fi126
                        }
                    }
                default:
                    switch ("\n\t" + "hello") {
                    case "foo bar":
                        var v127 bool = true
                        _ = v127
                    case "":
                        fmt.Println(fi111, fi111, fi111)
                    case "a":
                        v91 += 255
                    }
                }
                _ = v121
                _ = v122
            }
            fmt.Println(v91, v1)
            _ = fi111
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v107
    _ = v2
    _ = v3
    _ = v89
    _ = v90
    _ = v91
}
