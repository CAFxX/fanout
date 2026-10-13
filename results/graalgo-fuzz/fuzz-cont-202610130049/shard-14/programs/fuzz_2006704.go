package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 int64; F1 int16; F2 uint64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int8; F1 uint8; F2 int32 }
type GenS10 struct { F0 uint16; F1 int8 }
type GenS11 struct { F0 int64; F1 uint8; F2 int8 }
type GenS12 struct { F0 int8; F1 uint }
type GenS13 struct { F0 int32; F1 int; F2 int64 }
type GenS14 struct { F0 int; F1 int }
type GenS15 struct { F0 bool; F1 bool; F2 int16 }
type GenS16 struct { F0 uint16; F1 int16; F2 string }
type GenS17 struct { F0 int; F1 string; F2 uint }
type GenS18 struct { F0 uint }
func genFn0(p0 int8) int32 {
    var v5 float32 = 0.1
    if true {
        switch -675 {
        case -9223372036854775808:
            var v6 float64 = (1e10 + 0.0)
            _ = v6
        case 2147483647:
            var v7 int16 = 381
            _ = v7
        case 0:
            var v8 string = "foo bar"
            _ = v8
        default:
            if true {
                v9 := []int{42, 32768, -2147483648, 76}
                var v10 float32 = (0.0 - (v5 / v5))
                v5 += (float32(3.14159) + v5)
                _ = v10
                _ = v9
            }
        }
        var v11 int64 = 70
        switch "hello" {
        case "\n\t":
            p0 = -128
        }
        _ = v11
    }
    _ = p0
    _ = v5
    return -2147483648
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 94
    v3 := []uint32{2147483648, 425}
    v4 := []string{}
    v2 = v2
    if (-1 < 1) {
        if ((false || ((1000 <= 840) && (65535 < 32767))) || true) {
            fmt.Println(v2)
            if true {
                fmt.Println(v2)
                for fi12 := 0; fi12 < 1; fi12++ {
                    switch fi12 {
                    case 32767:
                        v13 := GenS1{F0: 950}
                        _ = v13
                    default:
                        fmt.Println(fi12)
                    }
                    var v14 uint32 = 4294967295
                    _ = fi12
                    _ = v14
                }
                var v15 float64 = 1e10
                _ = v15
            } else {
                fmt.Println(v2)
            }
        } else {
            var v16 uint32 = 2147483648
            if (0 <= 0) {
                switch -2147483648 {
                case 32767:
                    for fi17 := 0; fi17 < 2; fi17++ {
                        if true {
                            var v18 uint8 = (v2 * (^(v2)))
                            if false {
                                fmt.Println(v16)
                                var v19 uint32 = v16
                                if (255 <= 0) {
                                    for fi20 := 0; fi20 < 16; fi20++ {
                                        var v21 int = -1
                                        v3 = append(v3, v19)
                                        v2 = 101
                                        _ = fi20
                                        _ = v21
                                    }
                                    var v22 uint16 = 42
                                    _ = v22
                                }
                                _ = v19
                            }
                            var v23 uint8 = (v2 * 29)
                            _ = v18
                            _ = v23
                        } else {
                            var v24 uint32 = 819
                            for fi25 := 0; fi25 < 6; fi25++ {
                                if true {
                                    v26 := GenS1{F0: 18446744073709551615}
                                    switch "hello" {
                                    case "foo bar":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v27 uint = 65535
                                        v2 -= (v2 & (^(v2)))
                                        v28 := GenS2{F0: 42}
                                        _ = v27
                                        _ = v28
                                    case "hello":
                                        v1[("\n\t" + "\n\t")] = (^(fi17))
                                    }
                                    if (2147483647 > -1) {
                                        if ((true && false) || (!true)) {
                                            v29 := GenS3{F0: 0, F1: -32768}
                                            _ = v29
                                        }
                                        if (((!(((1.5 + float64(((3.14159 / 1.5) / 1.0))) < -2.25) || (127 != (v2 % (((v2 - v2)) | 1))))) && (!true)) && (42 < v24)) {
                                            switch "hello" {
                                            case "a":
                                                v24 -= 2147483647
                                            case "hello":
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v30 := GenS3{F0: 1, F1: 1}
                                                    _ = v30
                                                case "foo bar":
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v31 := v1["\n\t"]
                                                        _ = v31
                                                    case "":
                                                        var v32 uint64 = 2
                                                        _ = v32
                                                    default:
                                                        v24 -= 4294967295
                                                    }
                                                }
                                            }
                                            v33 := GenS4{F0: 830, F1: -1}
                                            _ = v33
                                        } else {
                                            var v34 string = "\n\t"
                                            var v35 uint16 = 1
                                            var v36 int16 = 42
                                            if (42 < (v36 * v36)) {
                                                switch (fi25 - (fi17 / ((int(v36)) | 1))) {
                                                case 42:
                                                    if false {
                                                        v2 = v2
                                                    }
                                                default:
                                                    v16 -= (v24 / ((65535) | 1))
                                                }
                                                v37 := GenS2{F0: 427}
                                                switch int(v16) {
                                                case -1:
                                                    if ((42 <= 2) && true) {
                                                        fmt.Println(fi25, fi25)
                                                        v16 += (v24 + v16)
                                                    } else {
                                                        switch (fi25 * (^(fi17))) {
                                                        case 463:
                                                            v35 = (^(v35))
                                                            v24 = (v24 << ((154) & 63))
                                                        }
                                                        switch (v34 + v34) {
                                                        case "hello":
                                                            fmt.Println(v2, fi25, v35)
                                                        case "a":
                                                            var v38 uint16 = (v35 ^ (v35 << ((1) & 63)))
                                                            _ = v38
                                                        default:
                                                            var v39 uint = 473
                                                            _ = v39
                                                        }
                                                        _ = len(v3)
                                                    }
                                                case 1:
                                                    v40 := GenS3{F0: (v35 ^ v35), F1: (v36 & v36)}
                                                    _ = v40
                                                case 4294967295:
                                                    if false {
                                                        var v41 float64 = 3.14159
                                                        _ = v41
                                                    } else {
                                                        fmt.Println(v36)
                                                        v35 = v35
                                                        fmt.Println(v16, fi17)
                                                        if (false || false) {
                                                            var v42 float32 = float32(float64(float32(1.5)))
                                                            _ = v42
                                                        }
                                                    }
                                                    var v43 uint = 4294967295
                                                    _ = v43
                                                }
                                                _ = v37
                                            } else {
                                                v36 += 832
                                                v2 -= (^(v2))
                                                v16 -= 281
                                            }
                                            _ = v34
                                            _ = v35
                                            _ = v36
                                        }
                                        v1["foo bar"] = (fi17 + (fi25 << ((46) & 63)))
                                        var v44 string = "\n\t"
                                        _ = v44
                                    } else {
                                        switch "a" {
                                        case "foo bar":
                                            v16 += uint32(v2)
                                        case "":
                                            switch -329 {
                                            case 0:
                                                var v45 bool = ((0.0 < (1.5 - 0.1)) && (127 != 42))
                                                _ = v45
                                            case 668:
                                                v46 := GenS5{F0: 1, F1: -138}
                                                _ = v46
                                            case 127:
                                                v47 := GenS3{F0: 65535, F1: 0}
                                                _ = v47
                                            }
                                        default:
                                            switch "a" {
                                            case "a":
                                                v48 := GenS6{F0: (v2 ^ (v2 >> ((258) & 63))), F1: 747}
                                                _ = v48
                                            default:
                                                v49 := GenS3{F0: 255, F1: 0}
                                                _ = v49
                                            }
                                        }
                                    }
                                    _ = v26
                                } else {
                                    v1["a"] = -1
                                    if len(v3) > 0 {
                                        v3[0] = (v24 % ((v16) | 1))
                                    }
                                    fmt.Println(v2, fi25, v2)
                                }
                                _ = fi25
                            }
                            v2 = (^(v2))
                            _ = v24
                        }
                        v16 = v16
                        _ = fi17
                    }
                case -9223372036854775808:
                    switch 604 {
                    case 32767:
                        v2 += v2
                    case 1:
                        var v50 int16 = -1
                        _ = v50
                    }
                default:
                    if (-32768 < 0) {
                        var v51 float64 = 0.1
                        _ = v51
                    } else {
                        if false {
                            v16 -= 2147483647
                            v2 -= (v2 << ((65535) & 63))
                            v16 += (v16 ^ (v16 << ((2) & 63)))
                        } else {
                            if false {
                                fmt.Println(v16)
                                switch 32767 {
                                case 127:
                                    v3 = append(v3, (v16 % (((v16 % (((^(v16))) | 1))) | 1)))
                                case 816:
                                    switch 847 {
                                    case -500:
                                        if false {
                                            var v52 bool = true
                                            _ = v52
                                        }
                                    }
                                default:
                                    v16 = (v16 << ((255) & 63))
                                }
                                v53 := GenS7{F0: 4, F1: -519, F2: 18446744073709551615}
                                _ = v53
                            }
                            for fi54 := 0; fi54 < 19; fi54++ {
                                fmt.Println(fi54, fi54)
                                switch (fi54 << ((0) & 63)) {
                                case 32768:
                                    v1["foo bar"] = (-(fi54))
                                case -2147483648:
                                    if false {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if (-32768 == -32768) {
                                            for fi55 := 0; fi55 < 19; fi55++ {
                                                v56 := GenS1{F0: 65535}
                                                var v57 int64 = 0
                                                var v58 int = fi55
                                                _ = fi55
                                                _ = v56
                                                _ = v57
                                                _ = v58
                                            }
                                            if (32767 > 42) {
                                                if false {
                                                    if (!(65535 != 255)) {
                                                        _ = len(v3)
                                                    } else {
                                                        v59 := GenS8{F0: -1}
                                                        fmt.Println(fi54)
                                                        v2 = (v2 * 42)
                                                        _ = v59
                                                    }
                                                    if false {
                                                        fmt.Println(v2, v16, v16)
                                                    }
                                                    v60 := GenS2{F0: 255}
                                                    _ = v60
                                                }
                                                v61 := GenS4{F0: 1, F1: -201}
                                                _ = v61
                                            }
                                            if true {
                                                v1["a"] = fi54
                                                var v62 int = (fi54 | 1)
                                                for fi63 := 0; fi63 < 19; fi63++ {
                                                    fmt.Println(v62, v62, fi54)
                                                    v3 = append(v3, (v16 - (v16 ^ v16)))
                                                    v64 := v1["foo bar"]
                                                    var v65 uint16 = 42
                                                    _ = fi63
                                                    _ = v64
                                                    _ = v65
                                                }
                                                _ = v62
                                            } else {
                                                v2 -= (^(v2))
                                            }
                                        } else {
                                            v66 := GenS3{F0: 0, F1: 1}
                                            _ = v66
                                        }
                                        switch fi54 {
                                        case -2147483648:
                                            var v67 float64 = 1e10
                                            _ = v67
                                        case -562:
                                            if (1e10 <= ((((3.14159 + (float32(-2.25) - (1e10 * 3.14159))) - (0.1 + 0.1)) + 1.5) * 1e10)) {
                                                var v68 uint = 18446744073709551615
                                                v69 := GenS6{F0: 222, F1: uint64(v68)}
                                                var v70 float64 = 1.5
                                                _ = v68
                                                _ = v69
                                                _ = v70
                                            }
                                        case -372:
                                            switch (fi54 << ((2) & 63)) {
                                            case -2147483648:
                                                var v71 uint = 65535
                                                _ = v71
                                            case 32767:
                                                if false {
                                                    var v72 int8 = 1
                                                    if true {
                                                        v73 := GenS5{F0: 211, F1: -9223372036854775808}
                                                        v16 += v16
                                                        _ = v73
                                                    } else {
                                                        for fi74 := 0; fi74 < 7; fi74++ {
                                                            var v75 int16 = 0
                                                            var v76 uint64 = 401
                                                            _ = fi74
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                        v77 := GenS5{F0: 32767, F1: 1}
                                                        _ = v77
                                                    }
                                                    _ = v72
                                                }
                                            }
                                        }
                                    }
                                default:
                                    for fi78 := 0; fi78 < 17; fi78++ {
                                        fmt.Println(fi78)
                                        _ = fi78
                                    }
                                }
                                v79 := GenS9{F0: 127, F1: (v2 ^ 255), F2: 42}
                                _ = fi54
                                _ = v79
                            }
                            v2 -= v2
                        }
                    }
                }
                switch 1 {
                case 32767:
                    v80 := GenS4{F0: 32767, F1: 190}
                    _ = v80
                case 127:
                    fmt.Println(v2, v2)
                case 42:
                    for fi81 := 0; fi81 < 4; fi81++ {
                        fmt.Println(v16, v16, v2)
                        fmt.Println(fi81, v16, fi81)
                        v2 = 255
                        _ = fi81
                    }
                }
                v82 := GenS8{F0: 80}
                _ = v82
            } else {
                _ = len(v4)
            }
            _ = len(v3)
            _ = v16
        }
    } else {
        if (false || true) {
            for fi83 := 0; fi83 < 18; fi83++ {
                switch (("a" + "\n\t") + (("hello" + "a") + "a")) {
                case "":
                    var v84 string = ("a" + "a")
                    _ = v84
                case "\n\t":
                    v2 = 29
                }
                if (((("hello" + "hello") + "hello") + "a") > (("foo bar" + "") + "")) {
                    fmt.Println(fi83, v2)
                } else {
                    switch ("hello" + (("hello" + ("\n\t" + "foo bar")) + "")) {
                    case "":
                        _ = len(v1)
                    case "a":
                        var v85 int = fi83
                        _ = v85
                    }
                }
                v3 = append(v3, 18)
                _ = fi83
            }
        } else {
            for fi86 := 0; fi86 < 8; fi86++ {
                if true {
                    switch "" {
                    case "foo bar":
                        var v87 float32 = -2.25
                        _ = v87
                    case "":
                        for fi88 := 0; fi88 < 17; fi88++ {
                            v89 := GenS2{F0: 0}
                            _ = fi88
                            _ = v89
                        }
                    }
                    v1["a"] = fi86
                } else {
                    switch ("a" + "hello") {
                    case "\n\t":
                        v2 = 208
                    case "a":
                        v90 := GenS7{F0: -9223372036854775808, F1: 0, F2: 0}
                        _ = v90
                    case "foo bar":
                        v1[("foo bar" + ("\n\t" + (("foo bar" + ("" + ((("foo bar" + "") + ("foo bar" + ((("a" + "hello") + "\n\t") + (("foo bar" + "\n\t") + "a")))) + "foo bar"))) + "hello")))] = fi86
                    default:
                        var v91 uint16 = 65535
                        _ = v91
                    }
                    var v92 bool = false
                    switch fi86 {
                    case 32768:
                        for fi93 := 0; fi93 < 13; fi93++ {
                            v92 = true
                            fmt.Println(fi86)
                            var v94 bool = ((-(fi86)) != fi86)
                            _ = fi93
                            _ = v94
                        }
                    case 42:
                        switch (fi86 >> ((0) & 63)) {
                        case 1:
                            for fi95 := 0; fi95 < 8; fi95++ {
                                v96 := GenS6{F0: v2, F1: 271}
                                var v97 int8 = 57
                                var v98 uint16 = 42
                                _ = fi95
                                _ = v96
                                _ = v97
                                _ = v98
                            }
                        case -2147483648:
                            switch "\n\t" {
                            case "foo bar":
                                v99 := GenS5{F0: -1, F1: 42}
                                _ = v99
                            case "":
                                v100 := v1["hello"]
                                _ = v100
                            }
                        case 712:
                            var v101 uint64 = 0
                            _ = v101
                        }
                    }
                    _ = v92
                }
                switch "a" {
                case "hello":
                    if false {
                        v2 += (v2 % ((1) | 1))
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v2 -= v2
                        if false {
                            fmt.Println(fi86, fi86)
                            var v102 bool = (1.5 == 0.1)
                            for fi103 := 0; fi103 < 18; fi103++ {
                                var v104 int8 = -128
                                var v105 string = "hello"
                                _ = fi103
                                _ = v104
                                _ = v105
                            }
                            _ = v102
                        }
                    }
                }
                _ = fi86
            }
        }
        switch "foo bar" {
        case "\n\t":
            if (-1 >= 0) {
                for fi106 := 0; fi106 < 10; fi106++ {
                    fmt.Println(v2)
                    v2 -= 229
                    _ = fi106
                }
            }
        case "a":
            v2 = v2
        case "":
            if false {
                if false {
                    var v107 int64 = -9223372036854775808
                    v2 = (^(v2))
                    _ = v107
                } else {
                    _ = len(v4)
                }
            } else {
                fmt.Println(v2)
                var v108 uint = 219
                v2 = uint8(v108)
                _ = v108
            }
        default:
            v2 = 125
        }
    }
    for fi109 := 0; fi109 < 12; fi109++ {
        for fi110 := 0; fi110 < 3; fi110++ {
            var v111 uint = 924
            if true {
                var v112 uint32 = 65535
                switch ("" + "a") {
                case "a":
                    if ((0.1 > ((-2.25 + -2.25) + 0.0)) || (((0.0 * float32(((1e10 / 0.1) / 1.0))) / 1e10) < 3.14159)) {
                        var v113 int = (fi109 - fi110)
                        var v114 string = "\n\t"
                        _ = v113
                        _ = v114
                    }
                }
                v115 := GenS2{F0: 0}
                _ = v112
                _ = v115
            } else {
                v111 = 0
                var v116 float64 = 3.14159
                _ = v116
            }
            var v117 int8 = 0
            v111 = v111
            var v118 uint16 = 990
            _ = fi110
            _ = v111
            _ = v117
            _ = v118
        }
        v3 = append(v3, 1)
        if (v2 >= (^(v2))) {
            v2 = (v2 | v2)
            var v119 uint64 = 337
            if (514 < (^(v119))) {
                fmt.Println(v2, fi109, v119)
                if true {
                    var v120 bool = (!(genFn0(-1) >= 42))
                    var v121 int = (-(fi109))
                    _ = v120
                    _ = v121
                }
                switch 433 {
                case 32768:
                    v2 = 128
                case -1:
                    var v122 float64 = float64(1.5)
                    _ = v122
                case 128:
                    var v123 int16 = 1000
                    _ = v123
                default:
                    if true {
                        var v124 int64 = 623
                        if true {
                            v125 := GenS6{F0: v2, F1: v119}
                            v1[(("hello" + "\n\t") + (("\n\t" + ("a" + "hello")) + "a"))] = 2147483647
                            _ = v125
                        }
                        _ = v124
                    } else {
                        var v126 uint16 = 42
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v126
                    }
                }
            } else {
                for fi127 := 0; fi127 < 17; fi127++ {
                    var v128 int16 = int16(v2)
                    _ = fi127
                    _ = v128
                }
            }
            _ = v119
        } else {
            fmt.Println(v2, fi109)
            for fi129 := 0; fi129 < 18; fi129++ {
                var v130 float32 = (1.5 - 0.0)
                v2 -= (v2 + 1)
                _ = fi129
                _ = v130
            }
        }
        _ = fi109
    }
    v131 := GenS10{F0: 1, F1: 97}
    fmt.Println(v2)
    switch ("a" + "foo bar") {
    case "hello":
        if false {
            v2 = 102
            if (!false) {
                switch "a" {
                case "hello":
                    var v132 uint8 = (v2 & 255)
                    _ = v132
                case "foo bar":
                    for fi133 := 0; fi133 < 2; fi133++ {
                        v134 := GenS11{F0: 1, F1: 0, F2: 127}
                        v2 = v2
                        v135 := GenS7{F0: 1, F1: 42, F2: 243}
                        _ = fi133
                        _ = v134
                        _ = v135
                    }
                }
            }
            fmt.Println(v2)
        } else {
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
        }
    case "":
        delete(v1, "")
    case "a":
        var v136 int = 32767
        _ = v136
    }
    if false {
        var v137 uint32 = 1
        v2 = v2
        if (!false) {
            v137 -= (v137 << ((65535) & 63))
            var v138 float64 = (-2.25 - ((1e10 / 1e10) / (1.5 / 1.5)))
            if true {
                if ((("a" + ("hello" + "")) + ("a" + "hello")) < ("foo bar" + (("foo bar" + "\n\t") + "\n\t"))) {
                    var v139 uint8 = 0
                    _ = v139
                }
                for fi140 := 0; fi140 < 16; fi140++ {
                    if (!(v138 > float64(1.5))) {
                        switch (-(fi140)) {
                        case 1:
                            if ("foo bar" > ("a" + "")) {
                                v141 := GenS12{F0: 42, F1: uint(v137)}
                                v137 += v137
                                v142 := GenS13{F0: genFn0(100), F1: (fi140 / ((fi140) | 1)), F2: 206}
                                _ = v141
                                _ = v142
                            }
                        case -874:
                            switch (fi140 % (((fi140 % ((fi140) | 1))) | 1)) {
                            case -757:
                                var v143 uint8 = (v2 >> ((42) & 63))
                                _ = v143
                            }
                        case 9223372036854775807:
                            var v144 uint = 42
                            _ = v144
                        }
                        switch "" {
                        case "":
                            if (552 > -11) {
                                if ((-128 == 58) && true) {
                                    if true {
                                        var v145 uint32 = v137
                                        v2 = (^(v2))
                                        for fi146 := 0; fi146 < 10; fi146++ {
                                            v137 = (v145 ^ 4294967295)
                                            _ = len(v3)
                                            v137 -= (v137 << ((1) & 63))
                                            _ = fi146
                                        }
                                        _ = v145
                                    }
                                    for fi147 := 0; fi147 < 11; fi147++ {
                                        if (!false) {
                                            v2 -= v2
                                            fmt.Println(v2)
                                        }
                                        _ = fi147
                                    }
                                    fmt.Println(fi140)
                                }
                                v2 = (^(v2))
                                v148 := GenS14{F0: fi140, F1: (fi140 / ((fi140) | 1))}
                                _ = v148
                            }
                        case "\n\t":
                            v137 -= 1
                        default:
                            switch (^(fi140)) {
                            case -1:
                                if (4294967295 < 2) {
                                    for fi149 := 0; fi149 < 9; fi149++ {
                                        switch "hello" {
                                        case "foo bar":
                                            v150 := GenS4{F0: 32767, F1: 237}
                                            _ = v150
                                        default:
                                            v4 = append(v4, "foo bar")
                                        }
                                        if (-2147483648 == -2147483648) {
                                            v138 = (float64(0.1) / (v138 - v138))
                                            fmt.Println(v2)
                                            v1["a"] = 32768
                                        } else {
                                            if ((0.1 * (1e10 / (-2.25 + (float32(0.1) * (1.5 / (1.5 - 0.0)))))) >= (-2.25 * -2.25)) {
                                                switch fi140 {
                                                case -28:
                                                    v151 := GenS15{F0: false, F1: ((0.0 + (v138 - 3.14159)) > (v138 * v138)), F2: 588}
                                                    _ = v151
                                                case -553:
                                                    _ = len(v3)
                                                case -9223372036854775808:
                                                    var v152 uint = 2
                                                    _ = v152
                                                }
                                            }
                                            if true {
                                                switch "foo bar" {
                                                case "":
                                                    v153 := v1[""]
                                                    _ = v153
                                                default:
                                                    v2 += v2
                                                    switch fi140 {
                                                    case 4294967295:
                                                        if (255 <= uint16(v137)) {
                                                            if (("hello" + ("foo bar" + "hello")) <= "") {
                                                                v1["a"] = 0
                                                            }
                                                            var v154 int16 = -32768
                                                            _ = v154
                                                        }
                                                    }
                                                }
                                            } else {
                                                if (2 != 255) {
                                                    if false {
                                                        _ = len(v3)
                                                        switch fi149 {
                                                        case -6:
                                                            switch ((("foo bar" + "a") + "a") + ("" + "a")) {
                                                            case "hello":
                                                                fmt.Println(v138, fi149, v2)
                                                            case "\n\t":
                                                                if ((true || (!(true && (2147483647 == (v137 - v137))))) && false) {
                                                                    delete(v1, (("hello" + "hello") + (("foo bar" + ("hello" + "a")) + ("\n\t" + "\n\t"))))
                                                                    switch (fi149 >> ((4294967295) & 63)) {
                                                                    case 4294967295:
                                                                        v2 -= (v2 + v2)
                                                                    }
                                                                    var v155 string = ""
                                                                    _ = v155
                                                                } else {
                                                                    fmt.Println(fi149, v2, v137)
                                                                }
                                                            case "":
                                                                var v156 int8 = -55
                                                                _ = v156
                                                            default:
                                                                v157 := GenS8{F0: -1}
                                                                _ = v157
                                                            }
                                                        case 42:
                                                            if (v2 >= 241) {
                                                                var v158 string = ("a" + "\n\t")
                                                                _ = v158
                                                            } else {
                                                                var v159 int64 = -123
                                                                _ = v159
                                                            }
                                                            v137 = (v137 | 1)
                                                            v160 := GenS14{F0: fi140, F1: (^(fi140))}
                                                            _ = v160
                                                        default:
                                                            v2 = 128
                                                        }
                                                        var v161 uint32 = (v137 ^ (v137 % ((v137) | 1)))
                                                        _ = v161
                                                    }
                                                    var v162 bool = false
                                                    if v162 {
                                                        var v163 uint8 = (v2 % ((v2) | 1))
                                                        v164 := GenS6{F0: (^(v2)), F1: 801}
                                                        switch ("a" + "foo bar") {
                                                        case "":
                                                            if false {
                                                                fmt.Println(fi140)
                                                                v137 = (v137 * 4294967295)
                                                                fmt.Println(fi140)
                                                                v2 = 108
                                                            }
                                                        default:
                                                            v165 := GenS16{F0: 520, F1: 748, F2: ("a" + "a")}
                                                            _ = v165
                                                        }
                                                        _ = v163
                                                        _ = v164
                                                    } else {
                                                        v162 = v162
                                                    }
                                                    if (!(((float64(0.1) / ((float64(0.0) - (v138 * v138)) / v138)) >= 1e10) || (v162 && v162))) {
                                                        _ = len(v3)
                                                        fmt.Println(v137, v138)
                                                    }
                                                    var v166 uint16 = 741
                                                    switch ((("" + "") + "") + "") {
                                                    case "a":
                                                        v2 -= v2
                                                    default:
                                                        var v167 uint32 = (v137 >> ((255) & 63))
                                                        _ = v167
                                                    }
                                                    _ = v162
                                                    _ = v166
                                                }
                                            }
                                        }
                                        _ = fi149
                                    }
                                }
                            default:
                                if true {
                                    var v168 string = ("a" + ("foo bar" + "hello"))
                                    _ = v168
                                } else {
                                    v2 -= 98
                                    for fi169 := 0; fi169 < 12; fi169++ {
                                        v170 := GenS17{F0: (fi169 << ((18446744073709551615) & 63)), F1: "hello", F2: 65535}
                                        v1["a"] = (fi169 / (((fi140 % (((fi140 / ((fi140) | 1))) | 1))) | 1))
                                        var v171 bool = true
                                        _ = fi169
                                        _ = v170
                                        _ = v171
                                    }
                                }
                            }
                        }
                    } else {
                        v137 += (v137 | (v137 % (((v137 - v137)) | 1)))
                        var v172 uint8 = (v2 & v2)
                        _ = v172
                    }
                    for fi173 := 0; fi173 < 19; fi173++ {
                        var v174 float32 = ((0.1 / 1.5) + 1e10)
                        _ = fi173
                        _ = v174
                    }
                    _ = fi140
                }
                v2 = (v2 % (((v2 | (v2 % (((v2 ^ 153)) | 1)))) | 1))
            } else {
                v138 -= (v138 / (((float64(((((1e10 + (1.5 - -2.25)) / (1e10 * float32(1e10))) + (1e10 + (0.0 / (float32((v138 / 1e10)) + 1e10)))) - float32((-2.25 + (v138 * v138))))) - 0.0) / v138) * ((v138 / float64((1e10 / 1.5))) * 0.1)))
            }
            _ = v138
        } else {
            v2 = (v2 << ((585) & 63))
            v137 = 65535
        }
        _ = v137
    }
    v175 := GenS16{F0: 32767, F1: 0, F2: "foo bar"}
    delete(v1, "foo bar")
    var v176 int32 = -450
    v177 := GenS12{F0: 100, F1: 18446744073709551615}
    var v178 uint8 = v2
    var v179 uint32 = 2147483648
    for fi180 := 0; fi180 < 7; fi180++ {
        v181 := GenS12{F0: 100, F1: 4294967295}
        _ = fi180
        _ = v181
    }
    if ((v176 >> ((42) & 63)) != 957) {
        for fi182 := 0; fi182 < 19; fi182++ {
            v183 := GenS18{F0: 65535}
            var v184 uint64 = 2
            _ = fi182
            _ = v183
            _ = v184
        }
    } else {
        v2 = v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v176:", v176)
    fmt.Println("v178:", v178)
    fmt.Println("v179:", v179)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v131
    _ = v175
    _ = v176
    _ = v177
    _ = v178
    _ = v179
    _ = v2
    _ = v3
    _ = v4
}
