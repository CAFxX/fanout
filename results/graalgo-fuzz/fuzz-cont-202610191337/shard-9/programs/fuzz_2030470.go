package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int32 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int16; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 bool; F2 uint }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int64; F1 int16 }
type GenS11 struct { F0 uint64; F1 int8; F2 int64 }
type GenS12 struct { F0 int64; F1 int8 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint32; F1 uint64; F2 string }

func main() {
    var v1 uint64 = 912
    var v2 uint64 = v1
    v3 := make(map[int]string)
    var v4 int32 = -991
    var v5 bool = true
    v6 := []uint32{}
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(v5, v1, v2)
        switch "a" {
        case "hello":
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS1{F0: 796, F1: 0, F2: v4}
                var v10 uint8 = 255
                var v11 int8 = 42
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        default:
            for fi12 := 0; fi12 < 11; fi12++ {
                if false {
                    if (2 > 297) {
                        v4 -= (v4 - (v4 & (v4 / ((2147483647) | 1))))
                        if ((-(v4)) > (v4 * (v4 * 794))) {
                            v5 = (!v5)
                        } else {
                            fmt.Println(fi12, v5)
                            var v13 int32 = -1
                            _ = v13
                        }
                        v14 := GenS1{F0: 2, F1: -864, F2: 134}
                        v15 := GenS1{F0: 264, F1: 599, F2: 1}
                        _ = v14
                        _ = v15
                    } else {
                        if ((true && v5) && v5) {
                            var v16 float64 = (1.5 + 3.14159)
                            _ = v16
                        }
                        var v17 int8 = -1
                        _ = v17
                    }
                } else {
                    if (("hello" > ("\n\t" + "foo bar")) && false) {
                        if ((v1 - v1) != 23) {
                            var v18 int16 = int16(v1)
                            if (!false) {
                                delete(v3, fi12)
                            } else {
                                if v5 {
                                    v1 -= 212
                                    v19 := GenS1{F0: 540, F1: 1, F2: v4}
                                    v18 += (v18 - v18)
                                    var v20 string = "a"
                                    _ = v19
                                    _ = v20
                                } else {
                                    var v21 string = ("" + "")
                                    if ((v5 || (((v18 & (v18 | 42)) >= (v18 % ((v18) | 1))) || (!(-755 > (v4 & 0))))) || false) {
                                        fmt.Println(v21, fi12)
                                        var v22 int32 = (^(v4))
                                        _ = v22
                                    }
                                    _ = v21
                                }
                                v2 += (v1 + 0)
                                var v23 uint = 4294967295
                                _ = v23
                            }
                            _ = v18
                        }
                        var v24 int32 = int32(fi12)
                        if v5 {
                            v24 -= v24
                            switch "a" {
                            case "":
                                v25 := GenS1{F0: 1, F1: -539, F2: -276}
                                _ = v25
                            case "foo bar":
                                _ = len(v6)
                            case "hello":
                                v4 = v24
                            default:
                                v5 = v5
                            }
                        } else {
                            v26 := v3[1]
                            _ = v26
                        }
                        _ = v24
                    } else {
                        switch "a" {
                        case "foo bar":
                            switch (fi12 * -1) {
                            case -379:
                                if (v5 && false) {
                                    if v5 {
                                        v6 = append(v6, 923)
                                        var v27 uint16 = 32767
                                        if len(v6) > 0 {
                                            v6[0] = 65535
                                        }
                                        _ = v27
                                    } else {
                                        delete(v3, fi7)
                                        var v28 uint16 = 0
                                        v2 -= (v2 & (v2 ^ v2))
                                        if (v5 && (-32768 > -988)) {
                                            v29 := GenS1{F0: 1, F1: -1, F2: v4}
                                            _ = v29
                                        } else {
                                            v2 -= (v2 - (v1 & (v1 / ((uint64(v28)) | 1))))
                                            delete(v3, int(v2))
                                            v30 := GenS2{F0: (fi12 | fi7), F1: (^(v4))}
                                            _ = v30
                                        }
                                        _ = v28
                                    }
                                    delete(v3, fi12)
                                    v31 := GenS1{F0: 18446744073709551615, F1: -604, F2: 42}
                                    _ = v31
                                }
                            case -9223372036854775808:
                                var v32 int = (fi7 % ((fi12) | 1))
                                _ = v32
                            case 1:
                                var v33 int8 = 100
                                _ = v33
                            default:
                                var v34 float32 = 1e10
                                _ = v34
                            }
                        case "":
                            if (v5 || (true || true)) {
                                switch ("\n\t" + ("" + "")) {
                                case "hello":
                                    switch fi12 {
                                    case 127:
                                        if (0.1 <= 1.5) {
                                            delete(v3, fi7)
                                            v5 = v5
                                        } else {
                                            if v5 {
                                                v35 := GenS3{F0: 32767, F1: 65535, F2: 0}
                                                v4 = v4
                                                var v36 int = int(v1)
                                                if v5 {
                                                    delete(v3, -1)
                                                    v5 = v5
                                                }
                                                v36 = v36
                                                var v37 uint16 = 2
                                                v5 = v5
                                                _ = v35
                                                _ = v36
                                                _ = v37
                                            }
                                            var v38 uint32 = 2
                                            fmt.Println(v38, v5)
                                            _ = v38
                                        }
                                    case -688:
                                        switch (fi12 >> ((545) & 63)) {
                                        case 0:
                                            if v5 {
                                                if (100 < 53) {
                                                    var v39 uint16 = 2
                                                    var v40 string = "foo bar"
                                                    var v41 uint8 = 255
                                                    _ = v39
                                                    _ = v40
                                                    _ = v41
                                                }
                                            } else {
                                                v5 = false
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                v4 = (v4 - 42)
                                                _ = len(v3)
                                            }
                                            v5 = (v5 && (false || v5))
                                        case 4294967295:
                                            v6 = append(v6, 2)
                                        default:
                                            v42 := GenS2{F0: (fi7 >> ((18446744073709551615) & 63)), F1: v4}
                                            _ = v42
                                        }
                                    case 9223372036854775807:
                                        var v43 int8 = 127
                                        _ = v43
                                    }
                                case "\n\t":
                                    v4 -= (v4 * (v4 >> ((880) & 63)))
                                }
                                var v44 int32 = v4
                                switch (fi7 + (fi7 % ((fi7) | 1))) {
                                case 9223372036854775807:
                                    if true {
                                        v1 += 0
                                        fmt.Println(v1, v4, v4)
                                        var v45 uint = 0
                                        _ = v45
                                    } else {
                                        v1 = 145
                                        if v5 {
                                            switch (fi12 * fi7) {
                                            case -776:
                                                if len(v6) > 0 {
                                                    v6[0] = 2147483647
                                                }
                                            case 211:
                                                v46 := GenS3{F0: int16(v1), F1: 252, F2: 110}
                                                _ = v46
                                            }
                                            v2 -= 18446744073709551615
                                            v5 = true
                                        }
                                        v44 += v44
                                        if ((float64(0.0) != 3.14159) || v5) {
                                            v3[-456] = "a"
                                            v4 = v44
                                        } else {
                                            if v5 {
                                                var v47 bool = (-831 >= 65)
                                                v3[fi7] = "a"
                                                _ = v47
                                            } else {
                                                _ = len(v3)
                                                v48 := GenS2{F0: fi7, F1: (v4 * v4)}
                                                _ = v48
                                            }
                                            v44 += 42
                                            v6 = append(v6, 42)
                                        }
                                    }
                                case 4294967295:
                                    var v49 bool = v5
                                    _ = v49
                                case 2147483647:
                                    v50 := GenS4{F0: 65535}
                                    _ = v50
                                default:
                                    v51 := GenS1{F0: 1, F1: 554, F2: v4}
                                    _ = v51
                                }
                                _ = v44
                            }
                        default:
                            var v52 string = ("a" + "\n\t")
                            _ = v52
                        }
                        var v53 string = ""
                        _ = v53
                    }
                }
                _ = fi12
            }
        }
        switch "" {
        case "a":
            for fi54 := 0; fi54 < 12; fi54++ {
                var v55 int64 = -9223372036854775808
                v56 := GenS5{F0: (127 != 2), F1: 2147483647}
                switch "foo bar" {
                case "hello":
                    fmt.Println(fi54)
                case "foo bar":
                    if v5 {
                        var v57 int32 = (v4 + -795)
                        _ = v57
                    } else {
                        if len(v6) > 0 {
                            v6[0] = 1
                        }
                        if (!false) {
                            v58 := v3[(fi7 * (fi54 / (((-(fi54))) | 1)))]
                            _ = v58
                        } else {
                            v5 = (-1 >= v55)
                            switch "" {
                            case "":
                                if len(v6) > 0 {
                                    v6[0] = 270
                                }
                            }
                            v5 = false
                        }
                    }
                }
                v1 -= (^(v1))
                var v59 int64 = 0
                _ = fi54
                _ = v55
                _ = v56
                _ = v59
            }
        default:
            for fi60 := 0; fi60 < 19; fi60++ {
                if (4294967295 == 0) {
                    var v61 uint8 = 185
                    v61 = v61
                    _ = v61
                } else {
                    if v5 {
                        v1 -= v1
                        v62 := GenS6{F0: false, F1: (-(v4))}
                        v4 = -1
                        _ = v62
                    }
                    var v63 uint8 = 42
                    if false {
                        v64 := GenS7{F0: -684}
                        var v65 uint8 = 2
                        v66 := GenS8{F0: int32(fi7), F1: false, F2: 4294967295}
                        _ = v64
                        _ = v65
                        _ = v66
                    }
                    switch fi60 {
                    case -343:
                        var v67 uint32 = uint32(v2)
                        _ = v67
                    case 32768:
                        if (!true) {
                            switch (fi7 & fi60) {
                            case -734:
                                fmt.Println(fi60, fi7, v1)
                                v68 := GenS2{F0: (fi7 * (^(fi7))), F1: (v4 * (-(v4)))}
                                v5 = (false || ((0.0 * 0.0) <= 3.14159))
                                _ = v68
                            }
                            v4 -= -23
                            v6 = append(v6, 233)
                        } else {
                            v69 := GenS5{F0: false, F1: 42}
                            _ = v69
                        }
                    case 128:
                        delete(v3, -780)
                    }
                    _ = v63
                }
                var v70 int32 = (^(v4))
                _ = fi60
                _ = v70
            }
        }
        _ = fi7
    }
    for fi71 := 0; fi71 < 19; fi71++ {
        v72 := GenS4{F0: 42}
        var v73 uint8 = 2
        var v74 int = (fi71 % ((-1) | 1))
        _ = fi71
        _ = v72
        _ = v73
        _ = v74
    }
    if (("a" <= "") && (0 < (v4 << ((255) & 63)))) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v6 = append(v6, 2)
    }
    if true {
        if (0 != 752) {
            switch 9223372036854775807 {
            case -996:
                v5 = (-89 == 42)
            default:
                v5 = ((!v5) && false)
            }
            for fi75 := 0; fi75 < 17; fi75++ {
                var v76 int64 = 1
                if (v4 <= (v4 << ((167) & 63))) {
                    if false {
                        delete(v3, (fi75 & -2147483648))
                        v1 += v1
                        if ((fi75 ^ fi75) > (fi75 ^ -2147483648)) {
                            switch 371 {
                            case -9223372036854775808:
                                var v77 uint64 = (v1 << ((65535) & 63))
                                _ = v77
                            default:
                                v1 += v1
                            }
                        }
                    }
                    if v5 {
                        switch "a" {
                        case "foo bar":
                            fmt.Println(v1, v1)
                        case "hello":
                            var v78 float32 = ((0.1 + 1e10) * ((1.5 * (1.5 * ((0.1 / -2.25) + (3.14159 - float32((-2.25 + float64(((1e10 / 1.0) - (((0.1 / -2.25) + -2.25) * 1e10))))))))) - 3.14159))
                            _ = v78
                        case "a":
                            fmt.Println(v5)
                        }
                        if (!(0 != 996)) {
                            var v79 uint = 65535
                            var v80 uint32 = 0
                            for fi81 := 0; fi81 < 16; fi81++ {
                                v82 := GenS5{F0: true, F1: -389}
                                var v83 int = 32767
                                switch "hello" {
                                case "":
                                    v84 := GenS9{F0: 65535}
                                    _ = v84
                                }
                                _ = fi81
                                _ = v82
                                _ = v83
                            }
                            _ = v79
                            _ = v80
                        } else {
                            switch (fi75 & (fi75 % ((1) | 1))) {
                            case 127:
                                if (!((^(v1)) == v1)) {
                                    v3[int(v4)] = ("foo bar" + "a")
                                }
                            }
                            v85 := GenS4{F0: 42}
                            if v5 {
                                v2 += (v2 << ((65535) & 63))
                                for fi86 := 0; fi86 < 7; fi86++ {
                                    switch "hello" {
                                    case "a":
                                        switch (fi86 | 4294967295) {
                                        case 4294967295:
                                            var v87 uint8 = 255
                                            _ = v87
                                        case -9223372036854775808:
                                            v5 = v5
                                        case 898:
                                            var v88 int32 = (^(v4))
                                            var v89 float64 = 0.0
                                            _ = v88
                                            _ = v89
                                        }
                                    default:
                                        if (-32768 == -32768) {
                                            v90 := GenS10{F0: (^(v76)), F1: 42}
                                            _ = v90
                                        }
                                    }
                                    v91 := v3[1]
                                    _ = fi86
                                    _ = v91
                                }
                            } else {
                                var v92 uint32 = 65535
                                _ = v92
                            }
                            _ = v85
                        }
                    }
                } else {
                    if v5 {
                        if false {
                            v93 := v3[fi75]
                            _ = v93
                        } else {
                            _ = len(v3)
                        }
                        fmt.Println(fi75, v76)
                    } else {
                        for fi94 := 0; fi94 < 17; fi94++ {
                            _ = len(v3)
                            _ = fi94
                        }
                        v95 := GenS11{F0: 4294967295, F1: -1, F2: (-(v76))}
                        _ = v95
                    }
                    if v5 {
                        switch 127 {
                        case -2147483648:
                            v4 = 1
                        case 9223372036854775807:
                            for fi96 := 0; fi96 < 13; fi96++ {
                                if (2 <= 0) {
                                    _ = len(v3)
                                    if (false || (!(2 <= 1))) {
                                        if (v5 && (!(!v5))) {
                                            v1 -= 4294967295
                                            v76 -= (v76 & (v76 | (-(v76))))
                                        }
                                        v1 -= (v1 ^ 157)
                                    }
                                    v76 = 0
                                }
                                switch "\n\t" {
                                case "a":
                                    v2 = v1
                                default:
                                    v4 += (-(v4))
                                }
                                var v97 bool = v5
                                v98 := GenS4{F0: 761}
                                var v99 string = "a"
                                _ = fi96
                                _ = v97
                                _ = v98
                                _ = v99
                            }
                        }
                        for fi100 := 0; fi100 < 17; fi100++ {
                            if v5 {
                                v3[288] = (("\n\t" + "a") + "a")
                            } else {
                                v2 += v2
                                switch fi75 {
                                case -9223372036854775808:
                                    switch ("foo bar" + ("\n\t" + ("foo bar" + ""))) {
                                    case "\n\t":
                                        if v5 {
                                            var v101 int32 = -1
                                            _ = v101
                                        } else {
                                            if len(v6) > 0 {
                                                v6[0] = 42
                                            }
                                        }
                                    case "":
                                        if len(v6) > 0 {
                                            v6[0] = 2
                                        }
                                    default:
                                        fmt.Println(v5)
                                    }
                                default:
                                    var v102 uint32 = 4294967295
                                    _ = v102
                                }
                            }
                            if ((v76 + (v76 * -9223372036854775808)) == -9223372036854775808) {
                                switch "hello" {
                                case "foo bar":
                                    var v103 int = (fi75 + fi100)
                                    _ = v103
                                case "hello":
                                    if v5 {
                                        if (18446744073709551615 > 0) {
                                            if (32767 > 941) {
                                                _ = len(v6)
                                            } else {
                                                if (-3 == 127) {
                                                    v6 = append(v6, 4294967295)
                                                    var v104 float32 = ((0.1 + 0.1) - 0.1)
                                                    _ = v104
                                                }
                                            }
                                        } else {
                                            var v105 int32 = v4
                                            _ = v105
                                        }
                                        if ((fi75 >> ((2) & 63)) == (-(fi75))) {
                                            var v106 uint8 = 0
                                            v5 = v5
                                            _ = v106
                                        } else {
                                            v107 := GenS12{F0: (v76 ^ v76), F1: 100}
                                            var v108 string = "\n\t"
                                            switch (fi100 >> ((989) & 63)) {
                                            case 207:
                                                v6 = append(v6, 2)
                                            case 127:
                                                var v109 float32 = (3.14159 + 0.0)
                                                _ = v109
                                            }
                                            _ = v107
                                            _ = v108
                                        }
                                    }
                                }
                            }
                            var v110 int64 = int64(fi75)
                            _ = fi100
                            _ = v110
                        }
                        switch "" {
                        case "hello":
                            if false {
                                for fi111 := 0; fi111 < 16; fi111++ {
                                    fmt.Println(v4)
                                    var v112 int16 = -461
                                    var v113 uint8 = 1
                                    _ = fi111
                                    _ = v112
                                    _ = v113
                                }
                                if v5 {
                                    switch "foo bar" {
                                    case "a":
                                        v3[(fi75 << ((1) & 63))] = "foo bar"
                                    default:
                                        fmt.Println(v5, v2, v76)
                                    }
                                } else {
                                    v114 := GenS13{F0: (v76 | (-(v76)))}
                                    v5 = (0.0 > 1e10)
                                    _ = v114
                                }
                            } else {
                                var v115 uint8 = 255
                                if false {
                                    switch ((("hello" + "a") + ("\n\t" + "")) + "\n\t") {
                                    case "":
                                        v1 -= v2
                                    case "\n\t":
                                        v116 := GenS12{F0: v76, F1: -14}
                                        _ = v116
                                    case "foo bar":
                                        for fi117 := 0; fi117 < 18; fi117++ {
                                            v118 := GenS11{F0: (v2 << ((18446744073709551615) & 63)), F1: 127, F2: v76}
                                            v119 := GenS14{F0: v4}
                                            switch 42 {
                                            case -2147483648:
                                                switch "hello" {
                                                case "\n\t":
                                                    var v120 float32 = 1.5
                                                    _ = v120
                                                }
                                            case 382:
                                                if (190 <= 2) {
                                                    switch (("a" + ("" + "\n\t")) + ("foo bar" + "foo bar")) {
                                                    case "a":
                                                        var v121 uint16 = 315
                                                        _ = v121
                                                    }
                                                    var v122 float32 = float32(-2.25)
                                                    _ = len(v6)
                                                    _ = v122
                                                } else {
                                                    v115 -= (v115 << ((4294967295) & 63))
                                                    if true {
                                                        v76 = (v76 + (v76 ^ v76))
                                                        switch (fi117 * (fi117 - int(v4))) {
                                                        case 32768:
                                                            v123 := v3[(^(fi117))]
                                                            _ = v123
                                                        case 838:
                                                            var v124 int = (fi117 / ((-2147483648) | 1))
                                                            _ = v124
                                                        }
                                                    }
                                                }
                                            case 32768:
                                                fmt.Println(v76, v115)
                                            }
                                            _ = fi117
                                            _ = v118
                                            _ = v119
                                        }
                                    }
                                    v4 = v4
                                }
                                switch fi75 {
                                case -1:
                                    switch "" {
                                    case "\n\t":
                                        for fi125 := 0; fi125 < 3; fi125++ {
                                            v2 -= 0
                                            v4 = int32(v115)
                                            _ = fi125
                                        }
                                    case "":
                                        switch "" {
                                        case "\n\t":
                                            var v126 int16 = 0
                                            _ = v126
                                        case "a":
                                            if v5 {
                                                v127 := GenS11{F0: (v1 & 73), F1: 100, F2: v76}
                                                _ = v127
                                            }
                                        default:
                                            v1 = v1
                                        }
                                    case "foo bar":
                                        if len(v6) > 0 {
                                            v6[0] = 1
                                        }
                                    }
                                default:
                                    fmt.Println(v115, v5, v4)
                                }
                                _ = v115
                            }
                        case "\n\t":
                            v6 = append(v6, 4294967295)
                        }
                    }
                    var v128 uint64 = (v1 | 2)
                    _ = v128
                }
                v129 := GenS14{F0: -2147483648}
                _ = fi75
                _ = v129
                _ = v76
            }
        } else {
            for fi130 := 0; fi130 < 18; fi130++ {
                for fi131 := 0; fi131 < 1; fi131++ {
                    if (float32(1e10) > 0.0) {
                        v132 := GenS15{F0: 2147483647, F1: (^(v2)), F2: "\n\t"}
                        v1 += 0
                        fmt.Println(fi131, v1, fi130)
                        _ = v132
                    }
                    _ = len(v6)
                    v133 := v3[(^(fi130))]
                    _ = fi131
                    _ = v133
                }
                _ = fi130
            }
            _ = len(v6)
        }
    } else {
        for fi134 := 0; fi134 < 8; fi134++ {
            v5 = ((-2.25 + float64(1e10)) == (float64(0.1) / float64((3.14159 - 1e10))))
            v135 := v3[fi134]
            _ = fi134
            _ = v135
        }
    }
    v5 = v5
    fmt.Println(v5, v2, v1)
    if v5 {
        var v136 uint64 = v1
        _ = v136
    }
    _ = len(v3)
    v4 -= -2147483648
    v5 = v5
    for fi137 := 0; fi137 < 1; fi137++ {
        if (!(!((130 <= 65535) && (47 < 750)))) {
            v138 := GenS13{F0: -1}
            if false {
                var v139 uint = 255
                var v140 string = "a"
                fmt.Println(v2, v140)
                _ = v139
                _ = v140
            } else {
                for fi141 := 0; fi141 < 8; fi141++ {
                    fmt.Println(v2, v1, v5)
                    _ = fi141
                }
            }
            if ((!(32767 > 325)) && v5) {
                fmt.Println(v5, v4)
                var v142 float32 = 3.14159
                _ = v142
            }
            _ = v138
        }
        if true {
            v143 := GenS13{F0: -487}
            _ = v143
        }
        _ = fi137
    }
    fmt.Println(v4, v5)
    if true {
        var v144 string = (((("hello" + "a") + "hello") + ((((("" + "hello") + "foo bar") + ("a" + ("\n\t" + ("foo bar" + (((("foo bar" + (((("foo bar" + ("\n\t" + "foo bar")) + (("\n\t" + ("\n\t" + "a")) + "foo bar")) + "foo bar") + "foo bar")) + "hello") + "foo bar") + ((((("hello" + "") + "\n\t") + "") + (("\n\t" + ("\n\t" + "\n\t")) + "")) + ("a" + ""))))))) + "") + "a")) + "")
        if v5 {
            v3[9223372036854775807] = v144
            for fi145 := 0; fi145 < 16; fi145++ {
                var v146 uint16 = 939
                _ = fi145
                _ = v146
            }
        } else {
            _ = len(v3)
            v144 = "a"
            switch -9223372036854775808 {
            case -647:
                v2 -= 50
            case -2147483648:
                delete(v3, 128)
            default:
                v4 -= -263
            }
        }
        _ = v144
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
