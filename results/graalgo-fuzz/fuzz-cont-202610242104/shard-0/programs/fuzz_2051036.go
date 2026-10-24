package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 bool; F2 int64 }
type GenS6 struct { F0 int; F1 bool; F2 uint64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 uint32; F1 bool }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int16; F1 uint16 }
type GenS13 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint32; F1 bool; F2 int16 }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 int64 }

func main() {
    var v1 int = 9223372036854775807
    v2 := []int{128, v1}
    var v3 int8 = -128
    v1 -= (v1 % ((v1) | 1))
    v2 = append(v2, 9223372036854775807)
    for fi4 := 0; fi4 < 5; fi4++ {
        var v5 uint64 = 900
        v5 -= 695
        _ = fi4
        _ = v5
    }
    v1 = v1
    var v6 int8 = v3
    if false {
        if false {
            v1 = (v1 ^ (v1 - (v1 + (^(v1)))))
        }
        var v7 uint32 = 46
        var v8 uint16 = 247
        _ = v7
        _ = v8
    } else {
        var v9 bool = (-17 < (v3 - v6))
        fmt.Println(v3)
        _ = v9
    }
    switch (-(v1)) {
    case 109:
        var v10 uint16 = 115
        _ = v10
    case 42:
        if false {
            v3 += (v3 >> ((42) & 63))
        }
    default:
        v11 := GenS1{F0: 9223372036854775807, F1: 18446744073709551615, F2: 18446744073709551615}
        _ = v11
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v12 := GenS1{F0: 9223372036854775807, F1: 4294967295, F2: 42}
    for fi13 := 0; fi13 < 20; fi13++ {
        switch (^(v1)) {
        case 32768:
            v14 := GenS1{F0: 9223372036854775807, F1: 35, F2: 688}
            _ = v14
        case 127:
            if false {
                switch "" {
                case "\n\t":
                    switch "foo bar" {
                    case "hello":
                        if (true && false) {
                            for fi15 := 0; fi15 < 6; fi15++ {
                                var v16 uint = 211
                                v17 := GenS1{F0: 0, F1: v16, F2: 256}
                                v18 := GenS1{F0: 1, F1: v16, F2: 42}
                                var v19 uint16 = 2
                                v16 = v16
                                v16 = (^(v16))
                                switch int(v19) {
                                case -9223372036854775808:
                                    var v20 int8 = (v3 >> (((^(v16))) & 63))
                                    _ = v20
                                case 42:
                                    v21 := GenS1{F0: -9223372036854775808, F1: v16, F2: 42}
                                    _ = v21
                                case 1:
                                    switch (("foo bar" + "foo bar") + ("" + "\n\t")) {
                                    case "hello":
                                        fmt.Println(fi13, v16, fi13)
                                    case "":
                                        v16 -= (v16 >> ((uint(v19)) & 63))
                                    case "foo bar":
                                        var v22 float32 = -2.25
                                        _ = v22
                                    default:
                                        switch "\n\t" {
                                        case "foo bar":
                                            fmt.Println(v3, v3, v6)
                                        case "a":
                                            v19 += (v19 / (((v19 / ((0) | 1))) | 1))
                                        }
                                    }
                                }
                                _ = fi15
                                _ = v16
                                _ = v17
                                _ = v18
                                _ = v19
                            }
                            var v23 int8 = 1
                            if (false && (2 == 0)) {
                                var v24 uint32 = 839
                                v25 := GenS1{F0: 1, F1: 0, F2: 1}
                                var v26 float64 = (((1e10 + 0.1) + (1e10 - 1.5)) - -2.25)
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                var v27 int32 = 0
                                _ = v27
                            }
                            _ = v23
                        } else {
                            v28 := GenS1{F0: -844, F1: 4294967295, F2: 4294967295}
                            fmt.Println(v3)
                            if true {
                                if false {
                                    _ = len(v2)
                                } else {
                                    if true {
                                        if len(v2) > 0 {
                                            v2[0] = v1
                                        }
                                        if false {
                                            if ("foo bar" <= "\n\t") {
                                                v1 += 4294967295
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v29 := GenS2{F0: 625}
                                                _ = v29
                                            }
                                            var v30 uint8 = 127
                                            for fi31 := 0; fi31 < 2; fi31++ {
                                                v32 := GenS3{F0: "a"}
                                                _ = fi31
                                                _ = v32
                                            }
                                            _ = v30
                                        } else {
                                            for fi33 := 0; fi33 < 14; fi33++ {
                                                if true {
                                                    var v34 string = ("\n\t" + "hello")
                                                    _ = len(v2)
                                                    _ = v34
                                                }
                                                var v35 int64 = 42
                                                v36 := GenS3{F0: ("" + "foo bar")}
                                                _ = fi33
                                                _ = v35
                                                _ = v36
                                            }
                                            var v37 int64 = 9223372036854775807
                                            switch "" {
                                            case "\n\t":
                                                for fi38 := 0; fi38 < 14; fi38++ {
                                                    _ = len(v2)
                                                    v39 := GenS1{F0: v37, F1: 4294967295, F2: 719}
                                                    _ = fi38
                                                    _ = v39
                                                }
                                            case "":
                                                var v40 float32 = 3.14159
                                                _ = v40
                                            default:
                                                v41 := GenS1{F0: v37, F1: 1, F2: 18446744073709551615}
                                                _ = v41
                                            }
                                            _ = v37
                                        }
                                    } else {
                                        fmt.Println(v3)
                                        v3 += v6
                                    }
                                }
                                if (360 <= 0) {
                                    if ((0.1 < ((float64(1e10) / 3.14159) / -2.25)) && (0 == 203)) {
                                        if (0 <= 416) {
                                            switch v1 {
                                            case 166:
                                                v42 := GenS4{F0: 18446744073709551615}
                                                _ = v42
                                            }
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if true {
                                                for fi43 := 0; fi43 < 10; fi43++ {
                                                    fmt.Println(fi43, v3)
                                                    var v44 uint = 18446744073709551615
                                                    _ = fi43
                                                    _ = v44
                                                }
                                                switch v1 {
                                                case 470:
                                                    var v45 int8 = (v3 ^ -10)
                                                    _ = v45
                                                }
                                            }
                                        }
                                        if true {
                                            var v46 uint32 = 65535
                                            v1 += 127
                                            fmt.Println(v6)
                                            _ = v46
                                        } else {
                                            _ = len(v2)
                                            if false {
                                                for fi47 := 0; fi47 < 6; fi47++ {
                                                    var v48 uint16 = 2
                                                    fmt.Println(fi13, fi47)
                                                    _ = fi47
                                                    _ = v48
                                                }
                                            }
                                            var v49 uint = 45
                                            _ = v49
                                        }
                                        fmt.Println(v6, v6, fi13)
                                    } else {
                                        v2 = append(v2, (fi13 + v1))
                                    }
                                    switch v1 {
                                    case -549:
                                        if false {
                                            var v50 uint = 42
                                            v3 -= (v6 * (v3 & -1))
                                            v51 := GenS3{F0: "foo bar"}
                                            _ = v50
                                            _ = v51
                                        }
                                    case 2147483647:
                                        v2 = append(v2, -2147483648)
                                    case 42:
                                        var v52 int8 = (v6 / (((v6 / (((-(v3))) | 1))) | 1))
                                        _ = v52
                                    default:
                                        _ = len(v2)
                                    }
                                } else {
                                    v53 := GenS5{F0: v1, F1: (2 <= 1), F2: 9223372036854775807}
                                    _ = v53
                                }
                                var v54 uint32 = 364
                                _ = v54
                            }
                            _ = v28
                        }
                    case "\n\t":
                        v55 := GenS2{F0: 2}
                        _ = v55
                    case "":
                        v6 = (v3 * 0)
                    default:
                        v1 = -474
                    }
                case "":
                    v56 := GenS3{F0: "hello"}
                    _ = v56
                case "foo bar":
                    if true {
                        fmt.Println(v6, v3)
                    }
                default:
                    var v57 uint64 = 202
                    _ = v57
                }
                fmt.Println(v1, v6)
                if true {
                    if true {
                        if (2147483647 < 4294967295) {
                            v6 -= (v6 | 75)
                            if false {
                                var v58 int32 = -551
                                _ = v58
                            } else {
                                if (0.0 >= 0.0) {
                                    switch "\n\t" {
                                    case "foo bar":
                                        v59 := GenS6{F0: (fi13 + (v1 + (fi13 - 487))), F1: false, F2: 42}
                                        _ = v59
                                    case "hello":
                                        for fi60 := 0; fi60 < 19; fi60++ {
                                            var v61 string = "foo bar"
                                            _ = fi60
                                            _ = v61
                                        }
                                    default:
                                        switch (fi13 | fi13) {
                                        case -9223372036854775808:
                                            v62 := GenS7{F0: 1000}
                                            _ = v62
                                        case 4294967295:
                                            v6 -= 127
                                        }
                                    }
                                    var v63 float64 = float64((-2.25 / 0.1))
                                    _ = v63
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = (fi13 ^ 221)
                                    }
                                    v1 = (v1 * 2147483647)
                                }
                                v1 = (fi13 | (v1 & v1))
                                switch (v1 ^ (v1 & 42)) {
                                case -709:
                                    for fi64 := 0; fi64 < 13; fi64++ {
                                        var v65 int8 = v6
                                        v66 := GenS2{F0: 18446744073709551615}
                                        v67 := GenS6{F0: (fi13 / ((v1) | 1)), F1: (-2.25 != float32(float64(((((float32((0.1 + 3.14159)) + 1.5) / 3.14159) / (-2.25 / (0.1 - 1e10))) - ((-2.25 * 1e10) / float32(((-2.25 - 3.14159) - 0.0))))))), F2: 1}
                                        var v68 bool = (!true)
                                        _ = fi64
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                        _ = v68
                                    }
                                default:
                                    switch 71 {
                                    case 32767:
                                        var v69 int64 = 1
                                        _ = v69
                                    default:
                                        v6 = 103
                                    }
                                }
                            }
                        }
                    }
                    for fi70 := 0; fi70 < 2; fi70++ {
                        var v71 float32 = ((0.1 - (1.5 * float32(0.1))) / 1.0)
                        v71 = v71
                        _ = fi70
                        _ = v71
                    }
                    _ = len(v2)
                } else {
                    fmt.Println(v3, fi13)
                    var v72 int64 = 30
                    for fi73 := 0; fi73 < 6; fi73++ {
                        v72 += v72
                        if (v72 >= (v72 + (^(v72)))) {
                            switch -515 {
                            case 721:
                                v74 := GenS3{F0: ("a" + ("foo bar" + "hello"))}
                                _ = v74
                            case -1:
                                v6 = (v6 << ((173) & 63))
                            default:
                                v2 = append(v2, fi13)
                            }
                            fmt.Println(v1)
                        }
                        _ = fi73
                    }
                    _ = v72
                }
            } else {
                switch ("" + ("" + "hello")) {
                case "hello":
                    v6 -= v6
                case "\n\t":
                    for fi75 := 0; fi75 < 14; fi75++ {
                        v76 := GenS6{F0: (v1 ^ (fi75 * (fi13 / (((-(fi75))) | 1)))), F1: true, F2: 0}
                        fmt.Println(v6, v3, v3)
                        v6 -= v3
                        _ = fi75
                        _ = v76
                    }
                case "a":
                    for fi77 := 0; fi77 < 9; fi77++ {
                        var v78 uint64 = 379
                        _ = fi77
                        _ = v78
                    }
                }
            }
        case -262:
            if (0.0 <= float64(((0.0 * -2.25) + float32(float64((-2.25 - (1.5 - 1e10))))))) {
                if (2 < 0) {
                    v79 := GenS2{F0: 608}
                    var v80 uint = 18446744073709551615
                    v6 = v6
                    _ = v79
                    _ = v80
                } else {
                    v1 = (v1 % (((fi13 * (v1 * -118))) | 1))
                    if (!true) {
                        if (false && true) {
                            for fi81 := 0; fi81 < 18; fi81++ {
                                fmt.Println(v3, fi81, fi13)
                                v82 := GenS3{F0: ("\n\t" + ("foo bar" + ""))}
                                _ = len(v2)
                                v83 := GenS5{F0: v1, F1: (!false), F2: 664}
                                _ = fi81
                                _ = v82
                                _ = v83
                            }
                            v1 += 128
                        } else {
                            v6 = (v3 / ((v6) | 1))
                        }
                        if true {
                            v84 := GenS4{F0: 0}
                            if false {
                                v6 = -1
                                v3 += (-(v3))
                            } else {
                                for fi85 := 0; fi85 < 3; fi85++ {
                                    if false {
                                        if true {
                                            var v86 uint32 = 4294967295
                                            _ = v86
                                        } else {
                                            v87 := GenS7{F0: 32767}
                                            v88 := GenS7{F0: 1000}
                                            _ = v87
                                            _ = v88
                                        }
                                    }
                                    fmt.Println(fi85)
                                    _ = fi85
                                }
                            }
                            for fi89 := 0; fi89 < 7; fi89++ {
                                switch -2147483648 {
                                case 128:
                                    v3 = v6
                                    v90 := GenS8{F0: fi13, F1: 32767}
                                    if false {
                                        v6 = v3
                                        switch "hello" {
                                        case "\n\t":
                                            v91 := GenS5{F0: 42, F1: (true && true), F2: 163}
                                            _ = v91
                                        case "a":
                                            var v92 string = "\n\t"
                                            _ = v92
                                        case "":
                                            v6 = -1
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = (fi89 - 4294967295)
                                        }
                                    } else {
                                        v1 -= fi13
                                        var v93 float64 = float64(3.14159)
                                        switch "hello" {
                                        case "hello":
                                            switch "hello" {
                                            case "foo bar":
                                                fmt.Println(v3)
                                            case "\n\t":
                                                if false {
                                                    if true {
                                                        if (true && true) {
                                                            var v94 int32 = 42
                                                            var v95 uint8 = 1
                                                            _ = v94
                                                            _ = v95
                                                        } else {
                                                            var v96 float32 = (1e10 + 1.5)
                                                            _ = v96
                                                        }
                                                        if false {
                                                            switch "" {
                                                            case "\n\t":
                                                                if (!(42 != 2147483647)) {
                                                                    v97 := GenS9{F0: 65535, F1: ((("a" >= "hello") || (false || true)) && (!((!true) || ("hello" > "a"))))}
                                                                    fmt.Println(v93, fi13, v3)
                                                                    v98 := GenS4{F0: 28}
                                                                    _ = v97
                                                                    _ = v98
                                                                }
                                                            case "a":
                                                                v93 = 0.1
                                                            case "hello":
                                                                if (true || (745 < 2)) {
                                                                    v93 += v93
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        if false {
                                                                            var v99 int16 = -32768
                                                                            var v100 int8 = v3
                                                                            _ = v100
                                                                            _ = v99
                                                                        } else {
                                                                            v101 := GenS10{F0: false}
                                                                            fmt.Println(v1)
                                                                            if false {
                                                                                v102 := GenS10{F0: false}
                                                                                _ = v102
                                                                            } else {
                                                                                v103 := GenS8{F0: (v1 * (v1 << ((855) & 63))), F1: 456}
                                                                                var v104 uint = 99
                                                                                switch fi89 {
                                                                                case 639:
                                                                                    if true {
                                                                                        v1 = (fi89 - fi89)
                                                                                        fmt.Println(fi13, fi89, v93)
                                                                                        if (1 == -1) {
                                                                                            v105 := GenS11{F0: v1}
                                                                                            var v106 int32 = 651
                                                                                            fmt.Println(fi13, v1)
                                                                                            _ = v105
                                                                                            _ = v106
                                                                                        }
                                                                                    }
                                                                                case -942:
                                                                                    v107 := GenS12{F0: -176, F1: 65535}
                                                                                    _ = v107
                                                                                }
                                                                                _ = v103
                                                                                _ = v104
                                                                            }
                                                                            _ = v101
                                                                        }
                                                                        v3 = v6
                                                                        if (!(!(true || false))) {
                                                                            v93 = float64(float32(1.5))
                                                                        } else {
                                                                            v108 := GenS6{F0: (-(fi89)), F1: (0 == 1), F2: 524}
                                                                            var v109 string = ""
                                                                            _ = v108
                                                                            _ = v109
                                                                        }
                                                                        v3 -= v3
                                                                    case "hello":
                                                                        if len(v2) > 0 {
                                                                            v2[0] = (fi13 ^ fi13)
                                                                        }
                                                                        v3 = v6
                                                                    }
                                                                    var v110 int8 = v3
                                                                    v6 += (v110 / ((v110) | 1))
                                                                    switch 1 {
                                                                    case 659:
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    case 581:
                                                                        v2 = append(v2, fi89)
                                                                    default:
                                                                        _ = len(v2)
                                                                    }
                                                                    v111 := GenS6{F0: (v1 % ((fi89) | 1)), F1: (1.5 < 0.1), F2: 25}
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 563
                                                                    }
                                                                    _ = v110
                                                                    _ = v111
                                                                }
                                                            default:
                                                                if ((fi89 * (fi89 / ((fi89) | 1))) < (v1 ^ 428)) {
                                                                    v112 := GenS9{F0: 42, F1: true}
                                                                    switch ("" + "") {
                                                                    case "\n\t":
                                                                        v6 = v6
                                                                    case "a":
                                                                        fmt.Println(fi13)
                                                                    default:
                                                                        v93 = 3.14159
                                                                    }
                                                                    _ = v112
                                                                } else {
                                                                    switch ("\n\t" + "") {
                                                                    case "\n\t":
                                                                        fmt.Println(v93, fi13)
                                                                        var v113 int64 = -9223372036854775808
                                                                        fmt.Println(v113)
                                                                        var v114 float64 = float64(((3.14159 / 1.5) - 1.5))
                                                                        var v115 int32 = -2147483648
                                                                        var v116 int = fi89
                                                                        if (v93 < -2.25) {
                                                                            if len(v2) > 0 {
                                                                                v2[0] = 127
                                                                            }
                                                                        }
                                                                        _ = v113
                                                                        _ = v114
                                                                        _ = v115
                                                                        _ = v116
                                                                    case "a":
                                                                        v1 = -63
                                                                    default:
                                                                        var v117 string = "hello"
                                                                        _ = v117
                                                                    }
                                                                    if (!true) {
                                                                        v118 := GenS13{F0: 236, F1: 259, F2: 77}
                                                                        v3 += v6
                                                                        var v119 uint8 = 1
                                                                        _ = v118
                                                                        _ = v119
                                                                    } else {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = (v1 | (fi13 % ((v1) | 1)))
                                                                        }
                                                                        _ = len(v2)
                                                                    }
                                                                }
                                                            }
                                                            if len(v2) > 0 {
                                                                v2[0] = fi89
                                                            }
                                                            v2 = append(v2, fi89)
                                                            switch ("foo bar" + "\n\t") {
                                                            case "foo bar":
                                                                switch (fi13 + -458) {
                                                                case 127:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = (v1 ^ (fi13 / ((32768) | 1)))
                                                                    }
                                                                case 128:
                                                                    if (767 <= 42) {
                                                                        v93 = v93
                                                                        v6 = (v3 * -106)
                                                                        var v120 int32 = 1
                                                                        _ = v120
                                                                    }
                                                                case 705:
                                                                    var v121 uint = 0
                                                                    _ = v121
                                                                default:
                                                                    v6 += 42
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        fmt.Println(v1, fi89)
                                                        var v122 bool = false
                                                        _ = v122
                                                    }
                                                }
                                            default:
                                                if (((-2.25 / 1e10) + 0.1) != float32(0.0)) {
                                                    var v123 int = (fi89 << ((255) & 63))
                                                    _ = v123
                                                }
                                            }
                                        case "":
                                            var v124 uint8 = 145
                                            _ = v124
                                        case "a":
                                            var v125 uint32 = 1
                                            _ = v125
                                        }
                                        switch "hello" {
                                        case "\n\t":
                                            if (true || false) {
                                                fmt.Println(fi89, fi89, fi13)
                                                if false {
                                                    v126 := GenS7{F0: -670}
                                                    v127 := GenS14{F0: 42}
                                                    _ = v126
                                                    _ = v127
                                                } else {
                                                    if true {
                                                        if false {
                                                            var v128 string = ""
                                                            v2 = append(v2, (^(fi13)))
                                                            _ = len(v2)
                                                            _ = v128
                                                        }
                                                        if true {
                                                            fmt.Println(v6, v3, fi13)
                                                        }
                                                    }
                                                }
                                                var v129 uint64 = 0
                                                _ = v129
                                            }
                                            var v130 int32 = -985
                                            _ = v130
                                        case "foo bar":
                                            v131 := GenS15{F0: 65535, F1: false, F2: 647}
                                            _ = v131
                                        }
                                        fmt.Println(v1)
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v132 uint32 = 763
                                            _ = v132
                                        }
                                        _ = v93
                                    }
                                    _ = v90
                                }
                                if false {
                                    fmt.Println(v1, v1)
                                } else {
                                    v6 -= (-(v6))
                                }
                                switch "\n\t" {
                                case "hello":
                                    v133 := GenS16{F0: 4294967295}
                                    _ = v133
                                case "":
                                    v134 := GenS2{F0: 274}
                                    var v135 int16 = 1000
                                    switch (("a" + "") + ("hello" + "\n\t")) {
                                    case "hello":
                                        if (false || true) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if false {
                                                var v136 float64 = -2.25
                                                if false {
                                                    v6 = v3
                                                    if len(v2) > 0 {
                                                        v2[0] = int(v135)
                                                    }
                                                    fmt.Println(fi13, v3)
                                                    var v137 bool = (fi89 > (-(fi89)))
                                                    _ = v137
                                                } else {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        fmt.Println(fi89)
                                                    case "\n\t":
                                                        fmt.Println(v1, v135)
                                                    default:
                                                        v135 += (v135 + (v135 + v135))
                                                    }
                                                    v135 = 1
                                                    _ = len(v2)
                                                }
                                                _ = v136
                                            }
                                            v138 := GenS8{F0: 9223372036854775807, F1: (v135 / ((v135) | 1))}
                                            v3 += (^(v6))
                                            var v139 uint8 = 145
                                            v140 := GenS2{F0: 147}
                                            _ = v138
                                            _ = v139
                                            _ = v140
                                        }
                                    case "a":
                                        v141 := GenS17{F0: 42}
                                        _ = v141
                                    case "":
                                        if false {
                                            v2 = append(v2, v1)
                                        }
                                    default:
                                        v135 = -32768
                                    }
                                    v142 := GenS8{F0: (^(fi89)), F1: int16(v6)}
                                    v143 := GenS16{F0: 2147483648}
                                    var v144 uint8 = 42
                                    var v145 int16 = (-(v135))
                                    _ = v134
                                    _ = v135
                                    _ = v142
                                    _ = v143
                                    _ = v144
                                    _ = v145
                                case "a":
                                    if (!("a" >= ("\n\t" + ""))) {
                                        v146 := GenS17{F0: 42}
                                        v147 := GenS15{F0: 4294967295, F1: false, F2: 1000}
                                        if false {
                                            v2 = append(v2, 0)
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        _ = v146
                                        _ = v147
                                    }
                                    var v148 uint = 0
                                    switch (-(v1)) {
                                    case 128:
                                        _ = len(v2)
                                    }
                                    _ = v148
                                default:
                                    var v149 uint32 = 1
                                    _ = v149
                                }
                                var v150 uint = 267
                                _ = fi89
                                _ = v150
                            }
                            _ = v84
                        } else {
                            var v151 bool = false
                            _ = v151
                        }
                    } else {
                        for fi152 := 0; fi152 < 4; fi152++ {
                            v3 += -128
                            var v153 int8 = 127
                            var v154 string = "hello"
                            _ = fi152
                            _ = v153
                            _ = v154
                        }
                        switch -2147483648 {
                        case -134:
                            v155 := GenS5{F0: -406, F1: (!true), F2: -819}
                            _ = v155
                        case -645:
                            switch 127 {
                            case 258:
                                if len(v2) > 0 {
                                    v2[0] = fi13
                                }
                            case -911:
                                var v156 uint32 = 2147483647
                                _ = v156
                            }
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    var v157 bool = true
                    _ = v157
                }
                var v158 uint8 = 232
                _ = v158
            } else {
                _ = len(v2)
                if (65535 == 2) {
                    var v159 int8 = v3
                    switch ("hello" + "hello") {
                    case "":
                        if false {
                            v160 := GenS4{F0: 4294967295}
                            v2 = append(v2, fi13)
                            v159 += v6
                            _ = v160
                        } else {
                            var v161 int32 = 1
                            _ = v161
                        }
                    case "\n\t":
                        v6 = (v3 + (v3 ^ (v3 | (v6 ^ v159))))
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = v159
                } else {
                    v1 = fi13
                    v1 += (fi13 % ((32768) | 1))
                    for fi162 := 0; fi162 < 17; fi162++ {
                        v3 += (v3 * (v3 / (((^(v3))) | 1)))
                        _ = fi162
                    }
                }
            }
        default:
            var v163 int16 = 67
            _ = v163
        }
        var v164 float64 = -2.25
        _ = fi13
        _ = v164
    }
    if false {
        fmt.Println(v6, v3, v1)
    } else {
        if true {
            var v165 int16 = -986
            for fi166 := 0; fi166 < 19; fi166++ {
                fmt.Println(fi166, v165)
                _ = fi166
            }
            var v167 uint = 47
            _ = v165
            _ = v167
        } else {
            switch (v1 % ((v1) | 1)) {
            case 42:
                for fi168 := 0; fi168 < 14; fi168++ {
                    if (4294967295 > 4294967295) {
                        v3 -= v6
                        var v169 float64 = 1.5
                        var v170 float64 = (1.5 - (v169 / v169))
                        _ = v169
                        _ = v170
                    }
                    _ = fi168
                }
            case 32768:
                var v171 int64 = -9223372036854775808
                _ = v171
            default:
                v1 += 32767
            }
            if false {
                for fi172 := 0; fi172 < 7; fi172++ {
                    if len(v2) > 0 {
                        v2[0] = v1
                    }
                    fmt.Println(fi172)
                    _ = fi172
                }
                if true {
                    v173 := GenS3{F0: ("foo bar" + "foo bar")}
                    _ = v173
                } else {
                    for fi174 := 0; fi174 < 1; fi174++ {
                        var v175 int = fi174
                        if len(v2) > 0 {
                            v2[0] = v1
                        }
                        _ = fi174
                        _ = v175
                    }
                }
            }
        }
    }
    var v176 uint64 = 88
    if (42 >= 650) {
        var v177 string = "hello"
        var v178 uint16 = 65535
        if (42 > -9223372036854775808) {
            v179 := GenS4{F0: 250}
            v6 = v6
            if (!((37 >= v6) || true)) {
                if len(v2) > 0 {
                    v2[0] = v1
                }
                var v180 int16 = -1
                _ = v180
            }
            _ = v179
        }
        _ = v177
        _ = v178
    } else {
        if true {
            var v181 string = "a"
            _ = v181
        } else {
            v2 = append(v2, v1)
            var v182 float32 = 1e10
            _ = v182
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v176:", v176)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v176
    _ = v2
    _ = v3
    _ = v6
}
