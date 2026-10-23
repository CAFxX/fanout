package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 string; F1 bool; F2 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8; F1 uint }
type GenS8 struct { F0 int8; F1 uint16; F2 int }

func main() {
    v1 := []uint32{2147483648, 42, 21}
    v2 := []int{204, 4294967295, 4294967295, 128}
    var v3 uint64 = 979
    v4 := make(map[uint32]string)
    if (!true) {
        v5 := GenS1{F0: 42}
        _ = v5
    }
    v3 += (v3 - v3)
    if true {
        v3 -= (v3 & v3)
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        delete(v4, 0)
        _ = fi6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 float32 = 1.5
        if ((1e10 - float64(((v8 + v8) / 1e10))) >= -2.25) {
            switch ("hello" + "foo bar") {
            case "a":
                if false {
                    v8 = (float32((float64(v8) / (((1.5 * ((3.14159 * ((1e10 - 1.5) / (1e10 - 1.5))) + 1e10)) + 1e10) / ((3.14159 - ((0.1 - (1.5 * 0.1)) - float64(-2.25))) + 1.5)))) / float32(1e10))
                    v9 := GenS1{F0: -1}
                    _ = v9
                } else {
                    v8 += (0.0 / v8)
                }
            case "foo bar":
                var v10 uint32 = 1
                _ = v10
            }
            if false {
                v8 += v8
                if true {
                    v8 += 3.14159
                    v11 := GenS1{F0: -1}
                    _ = v11
                }
            } else {
                var v12 int16 = 32767
                _ = v12
            }
        }
        _ = fi7
        _ = v8
    }
    if false {
        if false {
            v3 = (v3 & v3)
            var v13 bool = false
            _ = v13
        }
    } else {
        if true {
            v3 = (v3 >> ((2) & 63))
        }
        if false {
            var v14 int8 = 100
            fmt.Println(v3, v14)
            v15 := v4[4294967295]
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 1; fi16++ {
                if true {
                    v17 := GenS1{F0: int8(v3)}
                    v18 := GenS1{F0: -128}
                    _ = v17
                    _ = v18
                } else {
                    v2 = append(v2, -255)
                }
                v19 := GenS1{F0: 42}
                _ = fi16
                _ = v19
            }
            for fi20 := 0; fi20 < 20; fi20++ {
                var v21 bool = (false || ("a" > "hello"))
                if v21 {
                    v3 += v3
                }
                _ = fi20
                _ = v21
            }
            fmt.Println(v3)
        }
    }
    v3 -= (v3 & 9)
    for fi22 := 0; fi22 < 9; fi22++ {
        switch "a" {
        case "\n\t":
            v3 = 159
        case "":
            switch fi22 {
            case 9223372036854775807:
                if (-924 <= 530) {
                    for fi23 := 0; fi23 < 20; fi23++ {
                        fmt.Println(fi23)
                        switch "" {
                        case "\n\t":
                            var v24 uint = 42
                            _ = v24
                        case "":
                            if true {
                                if len(v2) > 0 {
                                    v2[0] = -963
                                }
                                fmt.Println(fi22, v3)
                            } else {
                                var v25 bool = (!(42 < 46))
                                _ = v25
                            }
                        case "foo bar":
                            v4[2147483648] = "hello"
                            if (398 >= 2147483647) {
                                if (128 != 255) {
                                    v3 += 2
                                } else {
                                    var v26 float32 = 1e10
                                    var v27 uint = 2
                                    v2 = append(v2, fi22)
                                    _ = v26
                                    _ = v27
                                }
                                switch ("hello" + "a") {
                                case "\n\t":
                                    v28 := GenS1{F0: -128}
                                    _ = v28
                                }
                                if false {
                                    v29 := GenS2{F0: 9223372036854775807, F1: -128}
                                    _ = v29
                                } else {
                                    v30 := v4[2]
                                    v31 := v4[83]
                                    v1 = append(v1, 2)
                                    _ = v30
                                    _ = v31
                                }
                                var v32 int = fi23
                                _ = v32
                            }
                            v33 := GenS3{F0: 128, F1: 4294967295}
                            _ = v33
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi23
                    }
                    var v34 bool = false
                    if false {
                        var v35 uint64 = 531
                        for fi36 := 0; fi36 < 18; fi36++ {
                            v34 = v34
                            if ("foo bar" == "") {
                                var v37 uint = 42
                                _ = v37
                            } else {
                                var v38 uint16 = 990
                                var v39 uint = 2
                                _ = len(v1)
                                _ = v38
                                _ = v39
                            }
                            v4[2] = ("a" + "a")
                            _ = fi36
                        }
                        v40 := GenS3{F0: 128, F1: 2147483647}
                        _ = v35
                        _ = v40
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            for fi41 := 0; fi41 < 17; fi41++ {
                                fmt.Println(v34, v34, fi41)
                                var v42 string = "\n\t"
                                _ = fi41
                                _ = v42
                            }
                        case "hello":
                            switch (("\n\t" + "a") + "") {
                            case "a":
                                for fi43 := 0; fi43 < 11; fi43++ {
                                    v44 := GenS1{F0: -128}
                                    _ = fi43
                                    _ = v44
                                }
                            default:
                                switch -2147483648 {
                                case 0:
                                    if (v34 || (-283 < -40)) {
                                        var v45 int16 = -939
                                        _ = v45
                                    } else {
                                        switch 2147483647 {
                                        case 9223372036854775807:
                                            if (100 > 0) {
                                                for fi46 := 0; fi46 < 17; fi46++ {
                                                    v3 += (v3 + (^(v3)))
                                                    v47 := GenS2{F0: 1, F1: 91}
                                                    var v48 float64 = (0.1 + (3.14159 / 3.14159))
                                                    _ = fi46
                                                    _ = v47
                                                    _ = v48
                                                }
                                                v49 := v4[2147483647]
                                                _ = v49
                                            }
                                        case 42:
                                            if v34 {
                                                v50 := GenS2{F0: -1, F1: 57}
                                                switch "hello" {
                                                case "\n\t":
                                                    _ = len(v4)
                                                case "a":
                                                    v34 = ("a" == "")
                                                }
                                                var v51 uint32 = 42
                                                _ = v50
                                                _ = v51
                                            } else {
                                                if (42 != 1) {
                                                    for fi52 := 0; fi52 < 14; fi52++ {
                                                        v3 -= v3
                                                        delete(v4, 2147483648)
                                                        _ = fi52
                                                    }
                                                    if (!(-2147483648 != -703)) {
                                                        for fi53 := 0; fi53 < 4; fi53++ {
                                                            v34 = v34
                                                            delete(v4, 107)
                                                            _ = fi53
                                                        }
                                                        v54 := GenS4{F0: 42, F1: 0}
                                                        _ = v54
                                                    } else {
                                                        switch (fi22 + (fi22 * (fi22 << ((1) & 63)))) {
                                                        case 2147483647:
                                                            fmt.Println(v3, fi22)
                                                        case -9223372036854775808:
                                                            if (!(true && true)) {
                                                                fmt.Println(v3, v34, v34)
                                                            } else {
                                                                var v55 float32 = (((1e10 / 1e10) - (1.5 + 3.14159)) + ((1e10 - 0.1) / 0.1))
                                                                for fi56 := 0; fi56 < 2; fi56++ {
                                                                    fmt.Println(fi22, v55)
                                                                    fmt.Println(fi56, v3)
                                                                    v57 := GenS1{F0: 42}
                                                                    _ = fi56
                                                                    _ = v57
                                                                }
                                                                if v34 {
                                                                    for fi58 := 0; fi58 < 19; fi58++ {
                                                                        switch ("\n\t" + "\n\t") {
                                                                        case "a":
                                                                            v3 = (v3 << ((440) & 63))
                                                                        case "hello":
                                                                            var v59 uint8 = 2
                                                                            _ = v59
                                                                        default:
                                                                            v60 := GenS2{F0: -9223372036854775808, F1: -1}
                                                                            _ = v60
                                                                        }
                                                                        v55 = v55
                                                                        if v34 {
                                                                            v61 := v4[uint32(v3)]
                                                                            var v62 uint8 = 1
                                                                            fmt.Println(fi58)
                                                                            v61 = (v61 + "\n\t")
                                                                            _ = v61
                                                                            _ = v62
                                                                        } else {
                                                                            _ = len(v1)
                                                                            v3 = (v3 * v3)
                                                                        }
                                                                        _ = fi58
                                                                    }
                                                                    var v63 float32 = (-2.25 - v55)
                                                                    for fi64 := 0; fi64 < 13; fi64++ {
                                                                        v55 += (1e10 / float32(-2.25))
                                                                        v3 += v3
                                                                        _ = fi64
                                                                    }
                                                                    _ = v63
                                                                } else {
                                                                    if (!v34) {
                                                                        fmt.Println(v3, v34)
                                                                        fmt.Println(v34)
                                                                        fmt.Println(v55, v3, fi22)
                                                                    } else {
                                                                        if true {
                                                                            var v65 uint = 4294967295
                                                                            _ = v65
                                                                        } else {
                                                                            var v66 bool = v34
                                                                            _ = len(v4)
                                                                            _ = v66
                                                                        }
                                                                        if v34 {
                                                                            v55 += (v55 * v55)
                                                                        } else {
                                                                            var v67 int = (^(fi22))
                                                                            _ = len(v4)
                                                                            _ = v67
                                                                        }
                                                                        switch fi22 {
                                                                        case 32767:
                                                                            v68 := GenS5{F0: "a", F1: v34, F2: ((("a" + "foo bar") + "foo bar") + "\n\t")}
                                                                            _ = v68
                                                                        case 128:
                                                                            v34 = (32767 == 1)
                                                                        case 32768:
                                                                            fmt.Println(v3, fi22, v3)
                                                                        default:
                                                                            v69 := GenS4{F0: 2, F1: 6}
                                                                            _ = v69
                                                                        }
                                                                    }
                                                                    for fi70 := 0; fi70 < 16; fi70++ {
                                                                        v34 = true
                                                                        _ = fi70
                                                                    }
                                                                    var v71 string = "a"
                                                                    _ = v71
                                                                }
                                                                _ = v55
                                                            }
                                                        case 1:
                                                            if (128 == 63) {
                                                                for fi72 := 0; fi72 < 10; fi72++ {
                                                                    fmt.Println(fi22)
                                                                    _ = fi72
                                                                }
                                                                switch -947 {
                                                                case -743:
                                                                    if (!v34) {
                                                                        if (-2.25 > (0.1 / -2.25)) {
                                                                            if v34 {
                                                                                fmt.Println(v3, v3, fi22)
                                                                                fmt.Println(fi22, v3)
                                                                            }
                                                                            var v73 int64 = 9223372036854775807
                                                                            _ = v73
                                                                        }
                                                                        var v74 int16 = 32767
                                                                        var v75 uint16 = 351
                                                                        _ = v74
                                                                        _ = v75
                                                                    }
                                                                default:
                                                                    delete(v4, uint32(v3))
                                                                }
                                                                v34 = false
                                                            } else {
                                                                var v76 int64 = 42
                                                                var v77 uint = 2
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                _ = v76
                                                                _ = v77
                                                            }
                                                        default:
                                                            v34 = (!v34)
                                                        }
                                                        v3 = v3
                                                        v34 = v34
                                                    }
                                                }
                                            }
                                        case 32767:
                                            if v34 {
                                                v34 = v34
                                            } else {
                                                v34 = (0 != 366)
                                                _ = len(v4)
                                                if v34 {
                                                    var v78 uint = 18446744073709551615
                                                    _ = v78
                                                }
                                            }
                                        default:
                                            fmt.Println(fi22)
                                        }
                                    }
                                case 9223372036854775807:
                                    if (true && (1 != 1)) {
                                        var v79 float32 = (float32(1e10) / -2.25)
                                        _ = v79
                                    }
                                case 4294967295:
                                    v34 = ((fi22 + fi22) <= 220)
                                default:
                                    switch (fi22 * 32768) {
                                    case -2147483648:
                                        v34 = ((fi22 != (^(fi22))) && (!true))
                                    case 127:
                                        for fi80 := 0; fi80 < 17; fi80++ {
                                            fmt.Println(fi80, fi22, v3)
                                            v34 = (-2.25 > 0.0)
                                            var v81 int = 648
                                            v81 -= (fi80 / ((fi22) | 1))
                                            var v82 float64 = 1e10
                                            _ = fi80
                                            _ = v81
                                            _ = v82
                                        }
                                    case -42:
                                        var v83 float32 = (((1.5 + ((float32(-2.25) + 1e10) + 0.0)) - 3.14159) - -2.25)
                                        _ = v83
                                    }
                                }
                            }
                        case "a":
                            switch fi22 {
                            case 42:
                                v84 := v4[2147483647]
                                _ = v84
                            case 559:
                                v85 := GenS5{F0: "\n\t", F1: (65535 > 337), F2: "hello"}
                                _ = v85
                            }
                        }
                        for fi86 := 0; fi86 < 4; fi86++ {
                            v3 = (v3 & 0)
                            if v34 {
                                var v87 uint32 = 42
                                var v88 bool = v34
                                _ = v87
                                _ = v88
                            } else {
                                if v34 {
                                    v89 := GenS3{F0: 1, F1: 845}
                                    v34 = v34
                                    if (true || v34) {
                                        fmt.Println(fi22)
                                        if (((!v34) || (false && (1 != 623))) || (0 < 104)) {
                                            v3 = v3
                                        }
                                        fmt.Println(fi86, v34, fi22)
                                        var v90 uint64 = (v3 & 1)
                                        _ = v90
                                    } else {
                                        if v34 {
                                            v3 -= 2
                                            var v91 int16 = 225
                                            switch fi86 {
                                            case 42:
                                                v4[2] = "hello"
                                                _ = len(v4)
                                            default:
                                                switch ("\n\t" + "foo bar") {
                                                case "":
                                                    var v92 uint32 = 4294967295
                                                    _ = v92
                                                default:
                                                    switch fi86 {
                                                    case 96:
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if len(v1) > 0 {
                                                                v1[0] = 2147483647
                                                            }
                                                            switch fi22 {
                                                            case 343:
                                                                _ = len(v1)
                                                            case 128:
                                                                if len(v1) > 0 {
                                                                    v1[0] = 274
                                                                }
                                                            }
                                                        default:
                                                            var v93 string = ("hello" + "a")
                                                            _ = v93
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v91
                                        }
                                        v3 = v3
                                        if v34 {
                                            v34 = ((!v34) || v34)
                                            _ = len(v4)
                                            v1 = append(v1, 65535)
                                            v3 = (^(v3))
                                        }
                                    }
                                    _ = v89
                                }
                                v3 = (v3 - v3)
                                var v94 uint16 = 498
                                _ = v94
                            }
                            _ = fi86
                        }
                    }
                    _ = v34
                }
            case 128:
                v95 := GenS1{F0: 42}
                _ = v95
            case 397:
                var v96 uint8 = 127
                _ = v96
            default:
                v97 := GenS5{F0: "", F1: ("foo bar" == "a"), F2: "\n\t"}
                _ = v97
            }
        case "hello":
            v98 := GenS1{F0: 42}
            _ = v98
        default:
            for fi99 := 0; fi99 < 14; fi99++ {
                var v100 uint64 = v3
                fmt.Println(v100, v3, fi99)
                var v101 string = ("hello" + (("" + ((("" + "\n\t") + ("" + (("\n\t" + ("hello" + "foo bar")) + ""))) + ("" + "\n\t"))) + ("\n\t" + "")))
                _ = fi99
                _ = v100
                _ = v101
            }
        }
        _ = fi22
    }
    for fi102 := 0; fi102 < 4; fi102++ {
        for fi103 := 0; fi103 < 19; fi103++ {
            if (!true) {
                v3 -= v3
            }
            _ = len(v4)
            _ = fi103
        }
        _ = fi102
    }
    switch "hello" {
    case "hello":
        var v104 int = -2147483648
        _ = v104
    case "":
        for fi105 := 0; fi105 < 16; fi105++ {
            if false {
                v106 := GenS1{F0: 102}
                if false {
                    fmt.Println(fi105, fi105)
                    var v107 int = (fi105 + 127)
                    _ = v107
                }
                _ = v106
            } else {
                v3 = (v3 / ((v3) | 1))
            }
            var v108 int = fi105
            _ = fi105
            _ = v108
        }
    }
    var v109 bool = true
    for fi110 := 0; fi110 < 14; fi110++ {
        var v111 uint32 = 170
        _ = len(v4)
        _ = fi110
        _ = v111
    }
    v3 -= v3
    if (192 < 1) {
        if (-2.25 >= 1e10) {
            for fi112 := 0; fi112 < 3; fi112++ {
                _ = len(v1)
                for fi113 := 0; fi113 < 4; fi113++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v114 uint32 = 2
                    var v115 int = fi112
                    _ = fi113
                    _ = v114
                    _ = v115
                }
                for fi116 := 0; fi116 < 6; fi116++ {
                    if (true || true) {
                        switch "\n\t" {
                        case "\n\t":
                            v3 -= 42
                        }
                        _ = len(v4)
                    } else {
                        v3 = (v3 - (v3 << ((18446744073709551615) & 63)))
                        if v109 {
                            var v117 float64 = ((1e10 - (0.0 + 1.5)) / 1e10)
                            var v118 int8 = 100
                            _ = v117
                            _ = v118
                        }
                        fmt.Println(v3, fi116)
                    }
                    var v119 int = 1
                    if true {
                        var v120 int8 = 1
                        _ = v120
                    }
                    _ = fi116
                    _ = v119
                }
                _ = fi112
            }
        }
    }
    for fi121 := 0; fi121 < 18; fi121++ {
        switch 2147483647 {
        case 0:
            var v122 uint8 = 1
            _ = v122
        case 128:
            if (true || ("a" == "foo bar")) {
                var v123 int64 = -494
                for fi124 := 0; fi124 < 3; fi124++ {
                    var v125 bool = v109
                    if v109 {
                        v109 = false
                    } else {
                        switch (fi121 / (((^(fi121))) | 1)) {
                        case 2147483647:
                            fmt.Println(v109)
                        }
                        var v126 uint8 = 0
                        var v127 uint8 = (v126 >> ((425) & 63))
                        _ = v126
                        _ = v127
                    }
                    fmt.Println(fi124, v3)
                    fmt.Println(fi121, v109, fi121)
                    _ = fi124
                    _ = v125
                }
                v109 = true
                _ = v123
            }
        case 249:
            switch fi121 {
            case -2147483648:
                if v109 {
                    if len(v2) > 0 {
                        v2[0] = (fi121 % (((fi121 ^ (^(fi121)))) | 1))
                    }
                    var v128 uint64 = 1
                    var v129 uint = 0
                    _ = v128
                    _ = v129
                }
            case 2147483647:
                for fi130 := 0; fi130 < 12; fi130++ {
                    if false {
                        var v131 int32 = -701
                        var v132 uint64 = 2
                        _ = v131
                        _ = v132
                    }
                    var v133 int64 = 778
                    var v134 string = ""
                    _ = fi130
                    _ = v133
                    _ = v134
                }
            case 128:
                switch (fi121 * 4294967295) {
                case -9223372036854775808:
                    v109 = v109
                case 127:
                    for fi135 := 0; fi135 < 11; fi135++ {
                        var v136 string = "a"
                        if (0 <= 2) {
                            var v137 int = fi135
                            _ = v137
                        } else {
                            v109 = v109
                            fmt.Println(fi121, fi135)
                        }
                        if v109 {
                            v1 = append(v1, 0)
                            v4[42] = (v136 + v136)
                            if v109 {
                                switch fi121 {
                                case -868:
                                    v138 := GenS1{F0: 127}
                                    _ = v138
                                case 42:
                                    v3 = (v3 ^ v3)
                                }
                                var v139 bool = (0 == 255)
                                _ = v139
                            } else {
                                v136 = v136
                            }
                        } else {
                            var v140 uint64 = 69
                            _ = v140
                        }
                        _ = fi135
                        _ = v136
                    }
                case 9223372036854775807:
                    v109 = (((float32(3.14159) - float32(0.0)) > 1.5) || v109)
                }
            default:
                if true {
                    v3 += (^(v3))
                    if true {
                        switch "a" {
                        case "\n\t":
                            v141 := GenS5{F0: ("hello" + "a"), F1: v109, F2: "a"}
                            _ = v141
                        case "":
                            v4[1] = ("foo bar" + "")
                        }
                        v142 := GenS6{F0: 127}
                        _ = v142
                    }
                    v143 := GenS7{F0: 1, F1: 291}
                    _ = v143
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi144 := 0; fi144 < 5; fi144++ {
                        switch (-(fi121)) {
                        case 127:
                            var v145 int16 = -440
                            _ = v145
                        }
                        var v146 uint8 = 52
                        if (0 < -1) {
                            var v147 uint = 18446744073709551615
                            v3 = 42
                            v1 = append(v1, 720)
                            var v148 uint = v147
                            _ = v147
                            _ = v148
                        }
                        _ = fi144
                        _ = v146
                    }
                    v149 := GenS1{F0: 42}
                    _ = v149
                }
            }
        }
        v150 := v4[42]
        v3 -= v3
        _ = fi121
        _ = v150
    }
    fmt.Println(v3)
    fmt.Println(v109, v109)
    if (false || (v109 && v109)) {
        var v151 int64 = -9223372036854775808
        v152 := GenS8{F0: 1, F1: 0, F2: -2147483648}
        _ = v151
        _ = v152
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v109:", v109)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v109
    _ = v2
    _ = v3
    _ = v4
}
