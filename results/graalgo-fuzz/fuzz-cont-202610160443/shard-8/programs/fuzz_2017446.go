package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 uint64; F1 uint64; F2 bool }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint; F1 int64 }
type GenS9 struct { F0 uint64; F1 uint32; F2 int }
type GenS10 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS11 struct { F0 int; F1 uint64 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int16; F1 int32 }
type GenS16 struct { F0 uint; F1 uint }
type GenS17 struct { F0 uint64; F1 int8 }
type GenS18 struct { F0 bool; F1 bool }
type GenS19 struct { F0 int8; F1 uint32; F2 int8 }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 uint32; F1 int16 }
type GenS22 struct { F0 string }
type GenS23 struct { F0 int32; F1 uint8 }
type GenS24 struct { F0 int32; F1 int8; F2 uint }
func genFn0(p0 uint16, p1 bool) int32 {
    if true {
        p1 = (!p1)
    } else {
        if ((p0 - p0) >= 65535) {
            var v7 int = 813
            v8 := GenS1{F0: 1, F1: 392}
            _ = v7
            _ = v8
        }
        var v9 float64 = (3.14159 - float64(float32((1.5 / 1.5))))
        var v10 uint32 = 2147483647
        _ = v10
        _ = v9
    }
    if p1 {
        if p1 {
            v11 := []int{-1}
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 10; fi12++ {
            switch "hello" {
            case "\n\t":
                for fi13 := 0; fi13 < 3; fi13++ {
                    if p1 {
                        v14 := GenS1{F0: 127, F1: 343}
                        _ = v14
                    }
                    p0 = p0
                    _ = fi13
                }
            case "hello":
                v15 := GenS1{F0: 127, F1: 106}
                _ = v15
            case "foo bar":
                switch 839 {
                case 32768:
                    p0 = (^(p0))
                case 213:
                    if p1 {
                        v16 := []int{(fi12 % ((fi12) | 1)), (fi12 & fi12)}
                        v16 = append(v16, 0)
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 11; fi17++ {
                            switch (fi17 % ((fi17) | 1)) {
                            case 4294967295:
                                if ((0 < 1) || (true || (42 >= 1000))) {
                                    v18 := []int{}
                                    if true {
                                        var v19 bool = p1
                                        var v20 int = fi17
                                        switch "hello" {
                                        case "a":
                                            switch ((("a" + "hello") + "foo bar") + ("\n\t" + "")) {
                                            case "":
                                                var v21 uint = 509
                                                _ = v21
                                            default:
                                                var v22 uint8 = 1
                                                _ = v22
                                            }
                                        default:
                                            p0 = p0
                                        }
                                        _ = v19
                                        _ = v20
                                    } else {
                                        var v23 int = fi12
                                        _ = len(v18)
                                        _ = v23
                                    }
                                    _ = v18
                                }
                            }
                            p1 = p1
                            v24 := []int{-957, fi17}
                            _ = fi17
                            _ = v24
                        }
                        var v25 int = (fi12 / (((fi12 * fi12)) | 1))
                        _ = v25
                    }
                case 4294967295:
                    for fi26 := 0; fi26 < 2; fi26++ {
                        v27 := []int{32767}
                        _ = fi26
                        _ = v27
                    }
                default:
                    fmt.Println(p0, fi12)
                }
            }
            p0 = 42
            _ = fi12
        }
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    v1 := []int{}
    v2 := make(map[int]int)
    var v3 uint8 = 42
    v4 := make(map[uint32]string)
    var v5 uint32 = 2
    var v6 int16 = 914
    for fi28 := 0; fi28 < 11; fi28++ {
        for fi29 := 0; fi29 < 18; fi29++ {
            if len(v1) > 0 {
                v1[0] = (fi28 & (fi29 - fi29))
            }
            if true {
                var v30 int64 = -511
                _ = v30
            }
            _ = fi29
        }
        if (1 <= 2) {
            if len(v1) > 0 {
                v1[0] = (^(fi28))
            }
            if ((3.14159 <= 1e10) || ((v5 - (v5 % (((v5 & 298)) | 1))) > uint32(v3))) {
                var v31 float64 = 1e10
                _ = v31
            } else {
                for fi32 := 0; fi32 < 20; fi32++ {
                    fmt.Println(fi28, v3)
                    v33 := GenS2{F0: v6}
                    _ = fi32
                    _ = v33
                }
            }
            if (716 == 2) {
                v34 := v4[v5]
                v6 = (-(v6))
                _ = v34
            }
        }
        _ = fi28
    }
    if true {
        if (!(531 > 533)) {
            var v35 uint = 731
            v1 = append(v1, 32768)
            _ = v35
        } else {
            if (!(true || (true && false))) {
                var v36 uint32 = v5
                var v37 uint64 = 1
                _ = v36
                _ = v37
            }
        }
        if (127 < 0) {
            if true {
                var v38 uint32 = v5
                _ = v38
            } else {
                v39 := GenS3{F0: 0, F1: 0, F2: 18446744073709551615}
                var v40 int = -519
                if true {
                    if false {
                        v41 := GenS4{F0: 9223372036854775807, F1: -1}
                        _ = v41
                    }
                    if false {
                        v42 := v4[(v5 >> ((18446744073709551615) & 63))]
                        _ = v42
                    }
                    if (((float32(-2.25) + (-2.25 * 1e10)) - (3.14159 * 1.5)) != 0.1) {
                        var v43 int16 = v6
                        _ = v43
                    } else {
                        switch (v40 ^ v40) {
                        case 4294967295:
                            v1 = append(v1, v40)
                        case 42:
                            fmt.Println(v40, v5, v5)
                        }
                    }
                } else {
                    v5 -= (v5 + 32)
                    for fi44 := 0; fi44 < 8; fi44++ {
                        if ((v5 % ((90) | 1)) < v5) {
                            var v45 int8 = -1
                            _ = v45
                        }
                        _ = fi44
                    }
                }
                _ = v39
                _ = v40
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    } else {
        fmt.Println(v5, v5, v6)
    }
    if (-1 > 0) {
        v46 := GenS2{F0: v6}
        v47 := GenS5{F0: 42, F1: (^(v5))}
        v48 := GenS6{F0: 2, F1: 4294967295, F2: false}
        _ = v46
        _ = v47
        _ = v48
    }
    switch "foo bar" {
    case "":
        v49 := GenS4{F0: -1, F1: 52}
        _ = v49
    case "\n\t":
        v50 := GenS7{F0: (v3 * uint8(v5)), F1: 1}
        _ = v50
    case "a":
        if false {
            var v51 uint8 = v3
            for fi52 := 0; fi52 < 15; fi52++ {
                v51 -= v51
                _ = len(v4)
                _ = fi52
            }
            _ = v51
        } else {
            v53 := GenS3{F0: 188, F1: 1, F2: 718}
            switch ("\n\t" + "hello") {
            case "hello":
                switch "a" {
                case "":
                    fmt.Println(v3, v3)
                case "hello":
                    v3 -= 128
                case "\n\t":
                    delete(v2, 127)
                }
            case "":
                switch ("hello" + "a") {
                case "":
                    switch -968 {
                    case 128:
                        var v54 bool = (1 < v6)
                        _ = v54
                    case 42:
                        if true {
                            if (2 < 42) {
                                fmt.Println(v6, v5, v6)
                                if ((float64((1.5 - (0.1 * float32(((float64((0.1 / 1e10)) - (float64(0.0) - 0.1)) - -2.25))))) != (1.5 * (float64((0.1 + 3.14159)) * float64(-2.25)))) || true) {
                                    for fi55 := 0; fi55 < 6; fi55++ {
                                        if (42 == -1) {
                                            _ = len(v2)
                                        } else {
                                            v6 = 1
                                        }
                                        v2[fi55] = (fi55 - (fi55 << ((0) & 63)))
                                        if (!false) {
                                            v3 = 182
                                            if len(v1) > 0 {
                                                v1[0] = fi55
                                            }
                                        } else {
                                            var v56 float32 = 0.0
                                            _ = v56
                                        }
                                        _ = fi55
                                    }
                                }
                            }
                            fmt.Println(v6)
                            for fi57 := 0; fi57 < 16; fi57++ {
                                v5 = (^(v5))
                                _ = fi57
                            }
                        }
                    case 32768:
                        if true {
                            var v58 uint = 873
                            v5 += (^(v5))
                            v5 = (v5 >> (((v58 - 4294967295)) & 63))
                            _ = v58
                        } else {
                            v59 := GenS2{F0: (v6 + v6)}
                            _ = v59
                        }
                    }
                case "foo bar":
                    v6 = -1
                case "a":
                    var v60 uint32 = (v5 + (v5 & 0))
                    _ = v60
                default:
                    var v61 int64 = 1
                    _ = v61
                }
            case "\n\t":
                v3 = 42
            default:
                switch "hello" {
                case "\n\t":
                    if true {
                        var v62 uint64 = 18446744073709551615
                        for fi63 := 0; fi63 < 11; fi63++ {
                            v64 := GenS6{F0: (v62 - 42), F1: v62, F2: (0 >= 42)}
                            if true {
                                v65 := GenS8{F0: uint(v62), F1: -9223372036854775808}
                                _ = v65
                            } else {
                                _ = len(v4)
                            }
                            _ = fi63
                            _ = v64
                        }
                        var v66 uint = 42
                        _ = v62
                        _ = v66
                    } else {
                        v6 = -1
                    }
                case "":
                    v3 -= (v3 % (((^(v3))) | 1))
                case "hello":
                    for fi67 := 0; fi67 < 14; fi67++ {
                        v68 := GenS9{F0: 18446744073709551615, F1: 312, F2: (fi67 % (((fi67 + (fi67 >> ((0) & 63)))) | 1))}
                        v69 := GenS8{F0: 65535, F1: -9223372036854775808}
                        _ = fi67
                        _ = v68
                        _ = v69
                    }
                default:
                    fmt.Println(v6)
                }
            }
            if ((v6 % ((v6) | 1)) < -32768) {
                for fi70 := 0; fi70 < 18; fi70++ {
                    var v71 int8 = 19
                    _ = fi70
                    _ = v71
                }
                var v72 uint16 = 255
                v3 -= v3
                _ = v72
            } else {
                v3 += 255
                switch ("" + "") {
                case "foo bar":
                    if (65535 != 859) {
                        v73 := GenS2{F0: (v6 - (v6 * (v6 & 1)))}
                        v3 -= 53
                        _ = v73
                    }
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                fmt.Println(v6, v3)
            }
            _ = v53
        }
    }
    for fi74 := 0; fi74 < 20; fi74++ {
        v75 := GenS2{F0: 1000}
        _ = fi74
        _ = v75
    }
    if ((((3.14159 + ((float64(0.0) / -2.25) / 1.5)) > 0.1) || true) && ((-575 > -734) && (!(-404 == 9223372036854775807)))) {
        var v76 bool = true
        _ = v76
    } else {
        var v77 int = 6
        v78 := GenS4{F0: -681, F1: -1}
        _ = v77
        _ = v78
    }
    if (1 != v5) {
        var v79 int16 = 1000
        switch 42 {
        case 0:
            fmt.Println(v5, v5)
        case 4294967295:
            for fi80 := 0; fi80 < 20; fi80++ {
                var v81 uint = 65535
                for fi82 := 0; fi82 < 7; fi82++ {
                    var v83 int64 = 42
                    v84 := GenS8{F0: (v81 | v81), F1: v83}
                    _ = fi82
                    _ = v83
                    _ = v84
                }
                _ = fi80
                _ = v81
            }
        default:
            v79 -= (v79 + v79)
        }
        if (true && false) {
            v6 = v6
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            _ = len(v4)
            switch -1 {
            case 4294967295:
                var v85 uint = 65535
                _ = v85
            case 434:
                v86 := GenS5{F0: 1, F1: (^(v5))}
                _ = v86
            case 489:
                for fi87 := 0; fi87 < 5; fi87++ {
                    delete(v4, 645)
                    _ = fi87
                }
            default:
                for fi88 := 0; fi88 < 20; fi88++ {
                    if (0 <= 18446744073709551615) {
                        v89 := GenS8{F0: 4294967295, F1: -9223372036854775808}
                        v3 += uint8(v5)
                        _ = v89
                    }
                    v79 = 0
                    v90 := GenS8{F0: 1, F1: -701}
                    _ = fi88
                    _ = v90
                }
            }
            if len(v1) > 0 {
                v1[0] = 0
            }
        }
        _ = v79
    }
    for fi91 := 0; fi91 < 6; fi91++ {
        var v92 float32 = (1.5 - 1e10)
        if (2 != 18446744073709551615) {
            switch fi91 {
            case 4294967295:
                fmt.Println(fi91)
            default:
                var v93 uint16 = 65535
                _ = v93
            }
            v94 := GenS4{F0: 9223372036854775807, F1: -776}
            _ = v94
        }
        _ = fi91
        _ = v92
    }
    for fi95 := 0; fi95 < 15; fi95++ {
        switch fi95 {
        case 4294967295:
            for fi96 := 0; fi96 < 7; fi96++ {
                var v97 bool = ((fi96 * fi96) >= fi95)
                if v97 {
                    v97 = (!v97)
                }
                v97 = (1.5 == 3.14159)
                v6 = -32768
                v98 := GenS5{F0: 126, F1: (v5 + v5)}
                _ = fi96
                _ = v97
                _ = v98
            }
        case -747:
            switch -802 {
            case 9223372036854775807:
                switch (^(fi95)) {
                case 603:
                    if false {
                        switch "\n\t" {
                        case "\n\t":
                            var v99 uint8 = (v3 ^ 49)
                            _ = v99
                        case "":
                            v100 := GenS4{F0: 1, F1: -22}
                            _ = v100
                        case "a":
                            var v101 float64 = 3.14159
                            _ = v101
                        }
                    } else {
                        for fi102 := 0; fi102 < 20; fi102++ {
                            if (9223372036854775807 <= -325) {
                                _ = len(v1)
                                v103 := GenS10{F0: 0, F1: 255, F2: v3}
                                _ = v103
                            }
                            _ = len(v4)
                            _ = fi102
                        }
                        var v104 int16 = v6
                        _ = v104
                    }
                case 9223372036854775807:
                    if (fi95 > (fi95 << ((42) & 63))) {
                        switch "" {
                        case "hello":
                            if false {
                                fmt.Println(v5, v3, v6)
                            }
                        }
                    } else {
                        switch ("hello" + "foo bar") {
                        case "\n\t":
                            switch (fi95 * (fi95 / (((fi95 >> ((65535) & 63))) | 1))) {
                            case 32768:
                                switch "hello" {
                                case "":
                                    var v105 uint16 = 0
                                    _ = v105
                                default:
                                    switch 66 {
                                    case -9223372036854775808:
                                        for fi106 := 0; fi106 < 20; fi106++ {
                                            switch "\n\t" {
                                            case "":
                                                v107 := GenS9{F0: 42, F1: v5, F2: fi95}
                                                _ = v107
                                            case "a":
                                                v6 = v6
                                            default:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            v108 := GenS11{F0: -377, F1: 42}
                                            v6 = v6
                                            if (!(!((v6 + (-(v6))) <= v6))) {
                                                v109 := GenS11{F0: 1, F1: 629}
                                                _ = v109
                                            }
                                            v110 := GenS12{F0: 42}
                                            _ = fi106
                                            _ = v108
                                            _ = v110
                                        }
                                    case 1:
                                        for fi111 := 0; fi111 < 13; fi111++ {
                                            _ = len(v2)
                                            v5 += (^(v5))
                                            _ = fi111
                                        }
                                    }
                                }
                            default:
                                for fi112 := 0; fi112 < 1; fi112++ {
                                    switch fi95 {
                                    case 32767:
                                        var v113 uint32 = v5
                                        _ = v113
                                    }
                                    _ = fi112
                                }
                            }
                        case "":
                            v3 -= (v3 / (((v3 | (^(v3)))) | 1))
                        default:
                            v114 := GenS13{F0: v5, F1: v3, F2: v3}
                            _ = v114
                        }
                    }
                case 42:
                    switch (-(fi95)) {
                    case -199:
                        v115 := v2[fi95]
                        _ = v115
                    case 2147483647:
                        v6 = v6
                    case -2147483648:
                        for fi116 := 0; fi116 < 9; fi116++ {
                            fmt.Println(fi95)
                            _ = fi116
                        }
                    }
                default:
                    v6 -= (-(v6))
                }
            }
        default:
            v1 = append(v1, (fi95 & (fi95 >> ((238) & 63))))
        }
        _ = fi95
    }
    var v117 uint = 512
    for fi118 := 0; fi118 < 8; fi118++ {
        v5 = (^(v5))
        _ = len(v2)
        _ = fi118
    }
    for fi119 := 0; fi119 < 7; fi119++ {
        v3 = v3
        if false {
            _ = len(v2)
            for fi120 := 0; fi120 < 14; fi120++ {
                v3 = v3
                _ = fi120
            }
        } else {
            switch "" {
            case "foo bar":
                if (42 <= 1) {
                    if len(v1) > 0 {
                        v1[0] = fi119
                    }
                }
            }
            if true {
                v121 := GenS9{F0: 42, F1: 4294967295, F2: 0}
                _ = v121
            } else {
                fmt.Println(v117)
                switch ("a" + "hello") {
                case "hello":
                    var v122 string = ("\n\t" + "")
                    _ = v122
                default:
                    v117 -= (^(v117))
                }
            }
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (56 > int8(v6)) {
                    switch "foo bar" {
                    case "foo bar":
                        v117 = v117
                    default:
                        v3 = 42
                    }
                    for fi123 := 0; fi123 < 19; fi123++ {
                        if true {
                            var v124 uint64 = 1
                            delete(v4, v5)
                            v4[(v5 & 294)] = "\n\t"
                            _ = v124
                        } else {
                            switch "\n\t" {
                            case "":
                                fmt.Println(fi119, v6, v5)
                            case "a":
                                switch 75 {
                                case -9223372036854775808:
                                    var v125 int32 = genFn0(0, ((^(v117)) < 18446744073709551615))
                                    _ = v125
                                case 127:
                                    if ((v5 / ((uint32(v3)) | 1)) >= uint32(v117)) {
                                        if (1 == 42) {
                                            if false {
                                                v126 := GenS1{F0: 42, F1: 586}
                                                fmt.Println(v6, fi123, v3)
                                                _ = v126
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v6 = (v6 * (v6 ^ 1000))
                                            switch "\n\t" {
                                            case "foo bar":
                                                delete(v2, 764)
                                                var v127 int8 = -128
                                                switch "\n\t" {
                                                case "foo bar":
                                                    switch fi123 {
                                                    case 545:
                                                        v117 -= v117
                                                    case 986:
                                                        var v128 int8 = (v127 >> ((v117) & 63))
                                                        _ = v128
                                                    default:
                                                        v117 += (v117 + (v117 >> (((^(v117))) & 63)))
                                                    }
                                                    v129 := GenS14{F0: (fi119 - 9223372036854775807)}
                                                    _ = v129
                                                }
                                                _ = v127
                                            }
                                            fmt.Println(v6)
                                        } else {
                                            var v130 int = 42
                                            var v131 uint32 = v5
                                            _ = v130
                                            _ = v131
                                        }
                                        var v132 uint8 = v3
                                        _ = v132
                                    } else {
                                        _ = len(v2)
                                        v133 := GenS15{F0: v6, F1: -209}
                                        _ = v133
                                    }
                                    switch ((((("\n\t" + "\n\t") + "hello") + "a") + ("foo bar" + "hello")) + "a") {
                                    case "foo bar":
                                        v6 = int16(v5)
                                    case "":
                                        v6 = 32767
                                    case "hello":
                                        fmt.Println(v6)
                                    default:
                                        v5 += uint32(v117)
                                    }
                                    var v134 float64 = (0.0 - 3.14159)
                                    if (1 < (v117 << (((v117 / (((v117 & v117)) | 1))) & 63))) {
                                        v134 += float64((((0.0 / (1e10 * ((((float32(1e10) - float32(1e10)) - 0.1) / 1.0) * 0.1))) / (0.0 - ((3.14159 * 0.1) + 1e10))) + 3.14159))
                                    } else {
                                        var v135 int8 = 0
                                        v136 := v2[32767]
                                        _ = v135
                                        _ = v136
                                    }
                                    _ = v134
                                default:
                                    var v137 uint8 = v3
                                    _ = v137
                                }
                            }
                            switch (fi123 - (^(fi123))) {
                            case 381:
                                var v138 uint64 = 427
                                _ = v138
                            case 32767:
                                var v139 uint64 = uint64(v117)
                                _ = v139
                            case 128:
                                v6 = v6
                            }
                            v140 := GenS7{F0: v3, F1: genFn0(0, (!(uint64(v3) >= 42)))}
                            _ = v140
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v3 = v3
                        _ = fi123
                    }
                    if true {
                        switch "" {
                        case "hello":
                            var v141 int16 = v6
                            _ = v141
                        case "foo bar":
                            if false {
                                v142 := GenS15{F0: (v6 - (^(v6))), F1: 0}
                                v143 := GenS3{F0: 140, F1: -120, F2: 4294967295}
                                _ = v142
                                _ = v143
                            } else {
                                var v144 bool = (true || false)
                                switch fi119 {
                                case 0:
                                    v3 = (v3 - uint8(v5))
                                case 815:
                                    if ((42 != 69) && true) {
                                        _ = len(v1)
                                        v144 = (!(-1 >= 42))
                                    } else {
                                        v117 = 2
                                    }
                                case -277:
                                    switch "\n\t" {
                                    case "hello":
                                        v145 := GenS3{F0: 255, F1: -1, F2: 1}
                                        _ = v145
                                    case "a":
                                        v146 := GenS16{F0: (v117 << (((v117 | (^(v117)))) & 63)), F1: 178}
                                        _ = v146
                                    default:
                                        var v147 int = fi119
                                        _ = v147
                                    }
                                default:
                                    v1 = append(v1, fi119)
                                }
                                if v144 {
                                    v148 := GenS17{F0: 1, F1: 80}
                                    v3 -= v3
                                    var v149 string = (("foo bar" + "hello") + "hello")
                                    _ = v148
                                    _ = v149
                                }
                                _ = v144
                            }
                        case "a":
                            v150 := GenS17{F0: uint64(v117), F1: 100}
                            _ = v150
                        default:
                            if ((^(v5)) == v5) {
                                for fi151 := 0; fi151 < 17; fi151++ {
                                    var v152 int8 = 0
                                    switch "" {
                                    case "":
                                        var v153 bool = true
                                        _ = v153
                                    case "hello":
                                        var v154 bool = ((float64(1e10) < 3.14159) && (!(255 >= 862)))
                                        _ = v154
                                    case "foo bar":
                                        var v155 string = "a"
                                        _ = v155
                                    default:
                                        delete(v4, (v5 & (v5 | (v5 >> ((v117) & 63)))))
                                    }
                                    if (v117 == 529) {
                                        v156 := GenS6{F0: 543, F1: uint64(v117), F2: true}
                                        var v157 uint = v117
                                        delete(v4, uint32(v117))
                                        _ = v156
                                        _ = v157
                                    } else {
                                        if (false && ((^(v117)) <= 721)) {
                                            v6 = (v6 % ((1) | 1))
                                        } else {
                                            v158 := GenS18{F0: (-481 != 9223372036854775807), F1: ((!false) && false)}
                                            if true {
                                                v5 += 65535
                                                var v159 bool = (1 > 2147483647)
                                                _ = v159
                                            } else {
                                                var v160 int = (^(fi119))
                                                fmt.Println(v5)
                                                _ = v160
                                            }
                                            var v161 int32 = -2147483648
                                            _ = v158
                                            _ = v161
                                        }
                                        var v162 float64 = 1.5
                                        v163 := GenS10{F0: (v117 | v117), F1: 32767, F2: 1}
                                        v164 := GenS19{F0: (v152 + v152), F1: (v5 + (v5 % ((552) | 1))), F2: (v152 >> ((v117) & 63))}
                                        if false {
                                            if (true || (32767 >= v6)) {
                                                _ = len(v2)
                                                v117 += 18446744073709551615
                                                if true {
                                                    if true {
                                                        _ = len(v2)
                                                        v3 = (v3 | 78)
                                                    } else {
                                                        v165 := GenS16{F0: 4294967295, F1: (^(v117))}
                                                        _ = v165
                                                    }
                                                    if false {
                                                        if ((65535 < uint(v5)) || (!(false && ((^(v152)) == (^(v152)))))) {
                                                            var v166 uint = (v117 / ((v117) | 1))
                                                            switch "foo bar" {
                                                            case "hello":
                                                                fmt.Println(v166, v6)
                                                            case "foo bar":
                                                                fmt.Println(v6)
                                                            case "a":
                                                                var v167 int16 = (v6 >> (((^(v166))) & 63))
                                                                _ = v167
                                                            default:
                                                                fmt.Println(v162, v152, v5)
                                                            }
                                                            var v168 uint64 = 4294967295
                                                            _ = v166
                                                            _ = v168
                                                        }
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = fi119
                                                    }
                                                    var v169 uint = (v117 & (v117 >> ((v117) & 63)))
                                                    if false {
                                                        if true {
                                                            switch "\n\t" {
                                                            case "a":
                                                                var v170 uint = (^(v169))
                                                                _ = v170
                                                            case "foo bar":
                                                                v171 := GenS20{F0: true}
                                                                _ = v171
                                                            default:
                                                                v117 += (v117 | (v169 ^ (v117 & (^(v169)))))
                                                            }
                                                            var v172 float64 = (float64(0.0) / 3.14159)
                                                            _ = v172
                                                        } else {
                                                            v162 += v162
                                                        }
                                                        v173 := GenS18{F0: (32767 <= 2), F1: true}
                                                        v162 = ((v162 * 1e10) * float64(0.0))
                                                        _ = v173
                                                    } else {
                                                        if len(v1) > 0 {
                                                            v1[0] = 32767
                                                        }
                                                        v174 := GenS4{F0: -882, F1: 42}
                                                        _ = v174
                                                    }
                                                    _ = v169
                                                } else {
                                                    switch "a" {
                                                    case "hello":
                                                        var v175 uint64 = 18446744073709551615
                                                        _ = v175
                                                    case "a":
                                                        if true {
                                                            v152 -= 13
                                                        } else {
                                                            v176 := v4[v5]
                                                            _ = v176
                                                        }
                                                    case "foo bar":
                                                        if false {
                                                            v5 = (^(v5))
                                                            v3 -= 2
                                                            v5 = uint32(v117)
                                                        } else {
                                                            if ((v6 << (((^(v117))) & 63)) >= (^(v6))) {
                                                                if true {
                                                                    var v177 int64 = -1
                                                                    var v178 float64 = 1e10
                                                                    _ = v177
                                                                    _ = v178
                                                                } else {
                                                                    v179 := GenS6{F0: 0, F1: 1, F2: false}
                                                                    var v180 float32 = 1e10
                                                                    _ = v179
                                                                    _ = v180
                                                                }
                                                                switch ("\n\t" + "a") {
                                                                case "foo bar":
                                                                    var v181 int = (fi119 / (((fi151 / ((0) | 1))) | 1))
                                                                    _ = v181
                                                                case "\n\t":
                                                                    if true {
                                                                        switch "" {
                                                                        case "a":
                                                                            var v182 bool = false
                                                                            _ = v182
                                                                        case "hello":
                                                                            v183 := GenS12{F0: 2147483647}
                                                                            _ = v183
                                                                        case "\n\t":
                                                                            switch (fi119 ^ (fi119 * fi119)) {
                                                                            case 2147483647:
                                                                                _ = len(v4)
                                                                                var v184 float32 = float32(((1.5 + 0.1) + v162))
                                                                                _ = v184
                                                                            case 32768:
                                                                                var v185 int8 = (-(v152))
                                                                                _ = v185
                                                                            default:
                                                                                switch "\n\t" {
                                                                                case "\n\t":
                                                                                    switch fi119 {
                                                                                    case -1:
                                                                                        switch (-(fi119)) {
                                                                                        case -9223372036854775808:
                                                                                            if (v3 < (v3 - v3)) {
                                                                                                var v186 int64 = 874
                                                                                                fmt.Println(fi151, v186, v5)
                                                                                                _ = v186
                                                                                            } else {
                                                                                                var v187 uint8 = v3
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                                _ = v187
                                                                                            }
                                                                                        default:
                                                                                            var v188 bool = ((v162 + 0.1) >= (v162 - 1e10))
                                                                                            _ = v188
                                                                                        }
                                                                                    case 128:
                                                                                        var v189 int8 = v152
                                                                                        _ = v189
                                                                                    case 673:
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = fi151
                                                                                        }
                                                                                    default:
                                                                                        delete(v2, (fi151 * (^(fi119))))
                                                                                    }
                                                                                case "a":
                                                                                    if true {
                                                                                        switch "foo bar" {
                                                                                        case "hello":
                                                                                            v190 := GenS1{F0: (v152 >> ((v117) & 63)), F1: 788}
                                                                                            _ = v190
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        default:
                                                                            v117 += 255
                                                                        }
                                                                        var v191 int64 = int64(v5)
                                                                        var v192 string = "foo bar"
                                                                        _ = v191
                                                                        _ = v192
                                                                    } else {
                                                                        var v193 int64 = -1
                                                                        var v194 uint8 = v3
                                                                        _ = v193
                                                                        _ = v194
                                                                    }
                                                                case "a":
                                                                    switch int(v152) {
                                                                    case 32768:
                                                                        var v195 int8 = 73
                                                                        _ = v195
                                                                    case 42:
                                                                        v5 -= v5
                                                                    case 128:
                                                                        v5 += v5
                                                                        switch (fi151 / (((fi151 << ((v117) & 63))) | 1)) {
                                                                        case 0:
                                                                            v196 := v4[v5]
                                                                            _ = v196
                                                                        default:
                                                                            fmt.Println(v6, v162, fi119)
                                                                        }
                                                                    default:
                                                                        switch -818 {
                                                                        case -1:
                                                                            delete(v4, uint32(v117))
                                                                        }
                                                                    }
                                                                default:
                                                                    var v197 int16 = v6
                                                                    _ = v197
                                                                }
                                                                var v198 uint32 = 0
                                                                _ = v198
                                                            } else {
                                                                v152 -= int8(fi119)
                                                            }
                                                        }
                                                    default:
                                                        v199 := GenS13{F0: (^(v5)), F1: 239, F2: (^(v3))}
                                                        _ = v199
                                                    }
                                                    v3 += 0
                                                }
                                            }
                                            var v200 int32 = genFn0(32767, (true || (-336 <= -1)))
                                            _ = v200
                                        } else {
                                            if false {
                                                switch "" {
                                                case "foo bar":
                                                    v201 := GenS20{F0: (float32(v162) == -2.25)}
                                                    _ = v201
                                                }
                                                fmt.Println(v6, fi151, v152)
                                                if true {
                                                    v2[9223372036854775807] = (fi151 >> ((uint(v5)) & 63))
                                                    v202 := GenS12{F0: -488}
                                                    var v203 uint = v117
                                                    v6 += v6
                                                    _ = v202
                                                    _ = v203
                                                } else {
                                                    v5 = 643
                                                    switch fi119 {
                                                    case -9223372036854775808:
                                                        switch (("" + "foo bar") + "foo bar") {
                                                        case "a":
                                                            var v204 int = (fi151 + -2147483648)
                                                            _ = v204
                                                        case "\n\t":
                                                            var v205 string = "a"
                                                            var v206 float32 = (0.0 / float32(v162))
                                                            _ = v205
                                                            _ = v206
                                                        case "hello":
                                                            _ = len(v4)
                                                        }
                                                        if true {
                                                            v6 = (v6 >> ((65535) & 63))
                                                            v3 += (v3 >> ((173) & 63))
                                                            if false {
                                                                var v207 int = fi151
                                                                _ = v207
                                                            } else {
                                                                var v208 int = (fi119 - (fi119 / ((int(v152)) | 1)))
                                                                fmt.Println(v208, v208, v208)
                                                                _ = v208
                                                            }
                                                        }
                                                    case 1:
                                                        var v209 float32 = 0.0
                                                        _ = v209
                                                    default:
                                                        if ((true || ((v5 & v5) >= uint32(v3))) && (false && false)) {
                                                            v162 = (0.0 * v162)
                                                            v3 = v3
                                                            switch "" {
                                                            case "a":
                                                                v1 = append(v1, 42)
                                                            case "\n\t":
                                                                var v210 int8 = v152
                                                                _ = v210
                                                            default:
                                                                if len(v1) > 0 {
                                                                    v1[0] = (fi151 * 4294967295)
                                                                }
                                                            }
                                                            var v211 int64 = -1
                                                            v117 = (v117 & (^(v117)))
                                                            _ = v211
                                                        } else {
                                                            fmt.Println(v6, fi119)
                                                            var v212 float32 = (3.14159 - 0.1)
                                                            _ = v212
                                                        }
                                                    }
                                                }
                                                switch 1 {
                                                case -483:
                                                    switch "foo bar" {
                                                    case "a":
                                                        v162 = v162
                                                    default:
                                                        var v213 int = -2147483648
                                                        _ = v213
                                                    }
                                                case 32767:
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if (false && false) {
                                                            var v214 int = (fi119 << ((v117) & 63))
                                                            v215 := GenS21{F0: v5, F1: (v6 - 1)}
                                                            v216 := GenS19{F0: (v152 << ((v117) & 63)), F1: 2, F2: (v152 / (((-(v152))) | 1))}
                                                            _ = v214
                                                            _ = v215
                                                            _ = v216
                                                        }
                                                    case "":
                                                        var v217 uint = v117
                                                        if ((^(v5)) > 2147483647) {
                                                            var v218 uint = (v217 | (v117 ^ (^(v117))))
                                                            if false {
                                                                var v219 int64 = 1
                                                                v2[fi151] = fi151
                                                                _ = v219
                                                            }
                                                            _ = v218
                                                        }
                                                        _ = v217
                                                    }
                                                case -1:
                                                    switch (("\n\t" + (("\n\t" + "") + "")) + ("\n\t" + ("foo bar" + ("a" + "a")))) {
                                                    case "hello":
                                                        v117 = (v117 ^ v117)
                                                    default:
                                                        fmt.Println(v162, v152, v3)
                                                    }
                                                default:
                                                    v152 = -1
                                                }
                                            }
                                            fmt.Println(v6, v3)
                                        }
                                        _ = v162
                                        _ = v163
                                        _ = v164
                                    }
                                    _ = fi151
                                    _ = v152
                                }
                            }
                        }
                        v220 := GenS3{F0: 255, F1: -21, F2: 0}
                        if len(v1) > 0 {
                            v1[0] = (fi119 * (fi119 + fi119))
                        }
                        _ = v220
                    }
                } else {
                    switch "hello" {
                    case "foo bar":
                        switch int(v6) {
                        case 854:
                            for fi221 := 0; fi221 < 2; fi221++ {
                                if false {
                                    v5 += v5
                                    var v222 uint = uint(v5)
                                    _ = v222
                                }
                                v117 += (v117 >> ((v117) & 63))
                                switch 42 {
                                case 42:
                                    var v223 uint64 = 2
                                    _ = v223
                                case -215:
                                    if true {
                                        v6 += (v6 - 200)
                                        v6 -= (v6 + v6)
                                    } else {
                                        _ = len(v1)
                                        v224 := GenS22{F0: ""}
                                        var v225 int64 = 703
                                        _ = v224
                                        _ = v225
                                    }
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                switch (-(fi221)) {
                                case 127:
                                    var v226 bool = (v6 <= v6)
                                    _ = v226
                                default:
                                    if true {
                                        v117 -= 296
                                    } else {
                                        if true {
                                            fmt.Println(fi221, fi221)
                                        }
                                        v227 := GenS21{F0: v5, F1: 1}
                                        var v228 int64 = -651
                                        _ = v227
                                        _ = v228
                                    }
                                }
                                _ = fi221
                            }
                        case 32767:
                            v117 = v117
                        case -967:
                            v4[42] = ((("" + "a") + "foo bar") + "a")
                        default:
                            var v229 uint64 = 0
                            _ = v229
                        }
                    case "a":
                        for fi230 := 0; fi230 < 6; fi230++ {
                            v231 := GenS23{F0: 1, F1: (^(v3))}
                            var v232 uint32 = (v5 / ((2) | 1))
                            var v233 float64 = 1.5
                            var v234 float64 = -2.25
                            _ = fi230
                            _ = v231
                            _ = v232
                            _ = v233
                            _ = v234
                        }
                    case "hello":
                        switch ("hello" + "\n\t") {
                        case "a":
                            v6 = (v6 ^ v6)
                        case "hello":
                            v117 += (v117 - (v117 >> (((^(v117))) & 63)))
                        case "\n\t":
                            switch (("a" + "hello") + ("\n\t" + ("\n\t" + "a"))) {
                            case "":
                                var v235 uint = v117
                                _ = v235
                            case "\n\t":
                                if true {
                                    var v236 uint64 = 582
                                    _ = v236
                                }
                            case "foo bar":
                                _ = len(v1)
                            default:
                                for fi237 := 0; fi237 < 11; fi237++ {
                                    v6 = v6
                                    if true {
                                        v5 = (v5 / ((v5) | 1))
                                        var v238 int64 = -1
                                        v239 := GenS23{F0: 170, F1: v3}
                                        var v240 string = ("foo bar" + ("" + "hello"))
                                        v241 := GenS20{F0: true}
                                        _ = v238
                                        _ = v239
                                        _ = v240
                                        _ = v241
                                    }
                                    v242 := GenS14{F0: 127}
                                    _ = fi237
                                    _ = v242
                                }
                            }
                        default:
                            _ = len(v1)
                        }
                    }
                    for fi243 := 0; fi243 < 5; fi243++ {
                        v244 := GenS24{F0: 2147483647, F1: 127, F2: v117}
                        v6 -= 1000
                        _ = fi243
                        _ = v244
                    }
                }
            }
        }
        var v245 float32 = ((3.14159 / 1.0) + (0.1 - 1.5))
        _ = fi119
        _ = v245
    }
    v2[127] = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v117:", v117)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
