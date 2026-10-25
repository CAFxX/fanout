package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint32; F1 int16; F2 string }
type GenS4 struct { F0 uint32; F1 int8; F2 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16; F1 int8; F2 int32 }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS12 struct { F0 uint; F1 int16; F2 int64 }
type GenS13 struct { F0 int64; F1 uint8; F2 int32 }
type GenS14 struct { F0 uint32; F1 string }
type GenS15 struct { F0 int8; F1 int16; F2 bool }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 int; F1 uint16; F2 uint8 }
func genFn0(p0 uint16, p1 int, p2 string) uint {
    p0 -= 343
    p2 = "hello"
    switch p1 {
    case -39:
        var v6 uint64 = 42
        _ = v6
    default:
        if true {
            v7 := []int{}
            fmt.Println(p0, p2, p2)
            p2 = p2
            _ = v7
        }
    }
    switch (p1 >> ((1) & 63)) {
    case 127:
        var v8 bool = (42 == 915)
        _ = v8
    case 42:
        if (false || (true && ((3.14159 * 0.0) < 0.0))) {
            p1 -= 0
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 uint64, p1 uint32) uint {
    var v9 int16 = 0
    if (42 <= -9223372036854775808) {
        for fi10 := 0; fi10 < 14; fi10++ {
            switch (fi10 * (fi10 ^ fi10)) {
            case 128:
                v11 := GenS1{F0: v9, F1: true}
                _ = v11
            case -381:
                v12 := GenS2{F0: 176, F1: "\n\t"}
                _ = v12
            case -1:
                for fi13 := 0; fi13 < 7; fi13++ {
                    v14 := []int{}
                    fmt.Println(fi10)
                    if (4294967295 >= 0) {
                        switch "hello" {
                        case "hello":
                            var v15 int64 = 988
                            _ = v15
                        case "a":
                            fmt.Println(v9)
                        case "":
                            if ((-(fi13)) != (^(fi10))) {
                                v16 := GenS3{F0: p1, F1: (-(v9)), F2: "hello"}
                                _ = v16
                            } else {
                                var v17 uint32 = (p1 ^ p1)
                                switch int(v17) {
                                case -49:
                                    var v18 int64 = 0
                                    _ = v18
                                }
                                _ = v17
                            }
                        default:
                            var v19 int32 = 1
                            _ = v19
                        }
                    }
                    _ = fi13
                    _ = v14
                }
            }
            var v20 uint8 = 1
            switch fi10 {
            case 578:
                for fi21 := 0; fi21 < 13; fi21++ {
                    v22 := GenS2{F0: -2147483648, F1: ("hello" + "foo bar")}
                    _ = fi21
                    _ = v22
                }
            case 823:
                if true {
                    p1 -= p1
                    var v23 bool = ("foo bar" < "")
                    v24 := []int{}
                    _ = v23
                    _ = v24
                }
            case 32767:
                v9 = (v9 / (((v9 << ((42) & 63))) | 1))
            default:
                for fi25 := 0; fi25 < 8; fi25++ {
                    fmt.Println(p1)
                    _ = fi25
                }
            }
            _ = fi10
            _ = v20
        }
        v26 := GenS4{F0: (p1 << ((genFn0(32767, -838, "foo bar")) & 63)), F1: 127, F2: (true || true)}
        _ = v26
    } else {
        var v27 int32 = -512
        v9 = v9
        _ = v27
    }
    _ = p0
    _ = p1
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 224
    var v2 int8 = 100
    var v3 int32 = -1
    v4 := []int{1, 9223372036854775807, -658, 32767}
    var v5 int8 = (v2 << ((uint(v1)) & 63))
    var v28 int16 = 32767
    switch 128 {
    case 4294967295:
        fmt.Println(v28)
    case 789:
        v29 := GenS2{F0: 2147483647, F1: "hello"}
        _ = v29
    case 32767:
        v30 := GenS4{F0: 0, F1: 1, F2: ((!true) && false)}
        _ = v30
    }
    v1 = (v1 << ((genFn0(1, 128, "hello")) & 63))
    v5 += (v2 | 100)
    _ = len(v4)
    v31 := GenS3{F0: uint32(v1), F1: v28, F2: "foo bar"}
    _ = len(v4)
    v2 = v2
    var v32 uint16 = 118
    for fi33 := 0; fi33 < 15; fi33++ {
        v32 -= 1
        switch (fi33 % (((fi33 * (fi33 | fi33))) | 1)) {
        case -1:
            var v34 int64 = 0
            _ = v34
        case 9223372036854775807:
            if false {
                switch "hello" {
                case "foo bar":
                    if ((v32 - (v32 * v32)) != 1) {
                        v32 = (v32 - 892)
                        fmt.Println(v32, fi33, v2)
                        for fi35 := 0; fi35 < 5; fi35++ {
                            var v36 uint64 = uint64(v32)
                            _ = fi35
                            _ = v36
                        }
                    }
                default:
                    if false {
                        v32 += v32
                    } else {
                        if (-2147483648 <= v3) {
                            fmt.Println(fi33)
                        } else {
                            v37 := GenS1{F0: v28, F1: (v1 == (v1 * v1))}
                            fmt.Println(v3, v32, v2)
                            v3 = (-(v3))
                            _ = v37
                        }
                    }
                }
                switch "foo bar" {
                case "foo bar":
                    for fi38 := 0; fi38 < 2; fi38++ {
                        switch (fi33 ^ fi33) {
                        case 1:
                            var v39 int16 = v28
                            _ = v39
                        case 32768:
                            if (uint16(v1) < 65535) {
                                var v40 uint32 = 870
                                v4 = append(v4, -160)
                                switch "\n\t" {
                                case "":
                                    switch "hello" {
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "":
                                            v5 = int8(v3)
                                        case "foo bar":
                                            v41 := GenS4{F0: v40, F1: 0, F2: ((((!(true && true)) && (2 > (v40 >> ((42) & 63)))) && (42 > v32)) || ((v2 - v2) <= v2))}
                                            _ = v41
                                        }
                                    case "a":
                                        var v42 float64 = 1.5
                                        _ = v42
                                    default:
                                        v43 := GenS5{F0: v1}
                                        _ = v43
                                    }
                                }
                                _ = v40
                            }
                        }
                        switch fi33 {
                        case 127:
                            var v44 int32 = 2147483647
                            _ = v44
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v5 -= v2
                        _ = fi38
                    }
                case "hello":
                    v28 -= 42
                case "a":
                    switch -412 {
                    case 994:
                        var v45 int8 = (v5 >> ((genFn1(605, 362)) & 63))
                        _ = v45
                    }
                }
            } else {
                if (1 > (v3 & (v3 * (^(v3))))) {
                    var v46 uint32 = 65535
                    _ = v46
                }
                var v47 uint16 = 124
                _ = v47
            }
        }
        _ = fi33
    }
    if false {
        switch -707 {
        case 128:
            v28 += (v28 / ((v28) | 1))
        case 893:
            v2 = v2
        case -306:
            if true {
                if true {
                    v5 = v5
                    var v48 uint16 = (v32 / (((v32 / ((0) | 1))) | 1))
                    _ = len(v4)
                    _ = v48
                }
                switch -38 {
                case 85:
                    for fi49 := 0; fi49 < 4; fi49++ {
                        v50 := GenS5{F0: v1}
                        for fi51 := 0; fi51 < 15; fi51++ {
                            if true {
                                v52 := GenS6{F0: 666}
                                switch (((("" + (("hello" + "") + ((("\n\t" + "hello") + "") + "\n\t"))) + ("hello" + "foo bar")) + ("hello" + "\n\t")) + "hello") {
                                case "":
                                    if false {
                                        var v53 bool = false
                                        var v54 string = ("" + "hello")
                                        _ = v53
                                        _ = v54
                                    } else {
                                        v55 := GenS6{F0: 4294967295}
                                        if len(v4) > 0 {
                                            v4[0] = (fi51 * (-(fi51)))
                                        }
                                        _ = v55
                                    }
                                }
                                var v56 uint = genFn1(228, 406)
                                v1 = (v1 % ((0) | 1))
                                v57 := GenS4{F0: 42, F1: v5, F2: true}
                                fmt.Println(v32)
                                v32 = v32
                                _ = v52
                                _ = v56
                                _ = v57
                            }
                            if (false && (!true)) {
                                v4 = append(v4, 86)
                                fmt.Println(fi49, v2)
                            } else {
                                v28 -= v28
                                if false {
                                    v58 := GenS2{F0: int32(v32), F1: ("" + "foo bar")}
                                    v32 += v32
                                    _ = v58
                                }
                                var v59 uint16 = v32
                                fmt.Println(v3, v32)
                                _ = v59
                            }
                            switch ("" + "\n\t") {
                            case "\n\t":
                                var v60 string = "\n\t"
                                _ = v60
                            case "":
                                v28 += (v28 * -13)
                            case "foo bar":
                                if false {
                                    if ("" <= "foo bar") {
                                        v61 := GenS6{F0: uint64(v1)}
                                        v32 = 65535
                                        _ = v61
                                    } else {
                                        v3 += int32(v1)
                                        if false {
                                            if false {
                                                if (!false) {
                                                    switch 128 {
                                                    case 2147483647:
                                                        if (fi51 > 4294967295) {
                                                            fmt.Println(v3, v32, v2)
                                                        }
                                                    case -785:
                                                        _ = len(v4)
                                                    case -340:
                                                        switch (fi49 % ((9223372036854775807) | 1)) {
                                                        case -1:
                                                            v62 := GenS7{F0: 713, F1: (v5 % ((-7) | 1)), F2: int32(fi51)}
                                                            _ = v62
                                                        case 42:
                                                            v3 -= v3
                                                        case 32768:
                                                            if ((0.0 == ((0.1 - 1e10) - 0.0)) && false) {
                                                                if (2147483647 >= 2147483647) {
                                                                    var v63 uint = 18446744073709551615
                                                                    fmt.Println(v2, v32, fi49)
                                                                    var v64 uint8 = (v1 + v1)
                                                                    var v65 bool = (2 < 4294967295)
                                                                    _ = v63
                                                                    _ = v64
                                                                    _ = v65
                                                                } else {
                                                                    switch (fi51 << ((4294967295) & 63)) {
                                                                    case 9223372036854775807:
                                                                        v2 = v5
                                                                    default:
                                                                        if len(v4) > 0 {
                                                                            v4[0] = fi49
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            _ = len(v4)
                                                        }
                                                    }
                                                }
                                                v32 -= 65535
                                                v66 := GenS8{F0: false, F1: 390}
                                                _ = v66
                                            }
                                            switch "hello" {
                                            case "foo bar":
                                                v5 -= (v5 + 100)
                                            case "":
                                                switch "a" {
                                                case "foo bar":
                                                    fmt.Println(v28, v2)
                                                default:
                                                    switch (((("" + ("" + "hello")) + "\n\t") + (("\n\t" + "foo bar") + "foo bar")) + "hello") {
                                                    case "foo bar":
                                                        fmt.Println(v3, v5, v1)
                                                    }
                                                }
                                            default:
                                                fmt.Println(v32, fi49, fi49)
                                            }
                                        }
                                        var v67 int = 1
                                        _ = v67
                                    }
                                    _ = len(v4)
                                    if false {
                                        switch fi51 {
                                        case 231:
                                            v2 += -1
                                        default:
                                            if false {
                                                switch 127 {
                                                case 42:
                                                    v4 = append(v4, fi51)
                                                }
                                                v32 -= v32
                                                if ((!(!(v1 >= v1))) || false) {
                                                    v1 += uint8(v32)
                                                    fmt.Println(v1, v1, v1)
                                                }
                                            }
                                        }
                                        var v68 float32 = 3.14159
                                        switch -297 {
                                        case -2147483648:
                                            switch (fi49 - fi51) {
                                            case 4294967295:
                                                v69 := GenS3{F0: 65535, F1: v28, F2: "hello"}
                                                _ = v69
                                            case 128:
                                                var v70 int64 = 42
                                                _ = v70
                                            default:
                                                if (((!false) || (220 < (v3 & v3))) && (274 <= 1000)) {
                                                    v3 = (v3 % (((v3 % ((v3) | 1))) | 1))
                                                }
                                            }
                                        case 32768:
                                            fmt.Println(v68, v1)
                                        default:
                                            fmt.Println(v1, v3, v1)
                                            if false {
                                                if (1 <= 255) {
                                                    v3 += 0
                                                } else {
                                                    fmt.Println(fi49, v5, v28)
                                                }
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                if ((!true) || false) {
                                                    if (!(("" + "hello") > "foo bar")) {
                                                        switch "a" {
                                                        case "":
                                                            v3 += (v3 % ((-244) | 1))
                                                        case "foo bar":
                                                            fmt.Println(v1)
                                                        case "a":
                                                            if true {
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                var v71 int8 = v2
                                                                _ = v71
                                                            }
                                                        default:
                                                            if false {
                                                                if (false || (1e10 > ((1.5 / 1.5) / 3.14159))) {
                                                                    if (237 <= 2) {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        fmt.Println(v68, fi51, v68)
                                                                        var v72 int = -2147483648
                                                                        _ = v72
                                                                    } else {
                                                                        switch fi51 {
                                                                        case -9223372036854775808:
                                                                            var v73 string = "hello"
                                                                            _ = v73
                                                                        case -529:
                                                                            var v74 uint8 = v1
                                                                            var v75 uint8 = uint8(v32)
                                                                            _ = v74
                                                                            _ = v75
                                                                        }
                                                                        fmt.Println(v2)
                                                                    }
                                                                    v68 += float32(0.1)
                                                                }
                                                            }
                                                        }
                                                        var v76 uint64 = 2
                                                        v77 := GenS3{F0: 2147483648, F1: (v28 / (((v28 ^ v28)) | 1)), F2: "a"}
                                                        _ = v76
                                                        _ = v77
                                                    } else {
                                                        if false {
                                                            fmt.Println(v68)
                                                        }
                                                        v78 := GenS2{F0: 0, F1: "hello"}
                                                        _ = v78
                                                    }
                                                    switch (fi51 * fi51) {
                                                    case 127:
                                                        var v79 float64 = ((3.14159 + 0.1) - -2.25)
                                                        _ = v79
                                                    case 42:
                                                        v68 = float32(1.5)
                                                    }
                                                    _ = len(v4)
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v32, v68, fi51)
                                                v1 = v1
                                                var v80 uint32 = 2147483648
                                                var v81 uint = 18446744073709551615
                                                _ = v80
                                                _ = v81
                                            }
                                        }
                                        v82 := GenS6{F0: 33}
                                        v3 -= v3
                                        _ = v68
                                        _ = v82
                                    } else {
                                        v83 := GenS8{F0: false, F1: 1}
                                        var v84 int16 = (v28 * v28)
                                        _ = v83
                                        _ = v84
                                    }
                                }
                            }
                            v4 = append(v4, 127)
                            _ = fi51
                        }
                        switch "" {
                        case "foo bar":
                            v85 := GenS9{F0: 2147483648}
                            _ = v85
                        default:
                            for fi86 := 0; fi86 < 5; fi86++ {
                                _ = len(v4)
                                v1 = v1
                                var v87 int64 = 42
                                _ = fi86
                                _ = v87
                            }
                        }
                        _ = fi49
                        _ = v50
                    }
                }
            }
        }
        var v88 int64 = 406
        _ = v88
    } else {
        switch "foo bar" {
        case "a":
            if len(v4) > 0 {
                v4[0] = 886
            }
        case "":
            v3 = (v3 * 199)
        }
        v32 -= v32
    }
    for fi89 := 0; fi89 < 11; fi89++ {
        var v90 uint32 = 579
        fmt.Println(v28)
        if (true && true) {
            for fi91 := 0; fi91 < 7; fi91++ {
                if ((42 == 538) && false) {
                    switch "foo bar" {
                    case "":
                        v5 -= v2
                        var v92 float32 = 0.0
                        _ = v92
                    case "hello":
                        switch (fi89 / ((fi91) | 1)) {
                        case 128:
                            v4 = append(v4, int(v5))
                        case -4:
                            v93 := GenS10{F0: "foo bar"}
                            _ = v93
                        case 42:
                            var v94 int16 = v28
                            _ = v94
                        }
                    case "a":
                        switch "\n\t" {
                        case "a":
                            v28 = v28
                        case "":
                            v28 -= (v28 + (^(v28)))
                        case "\n\t":
                            switch "hello" {
                            case "a":
                                if (0.1 >= 1e10) {
                                    v3 = v3
                                    if true {
                                        var v95 float64 = float64((3.14159 * 1.5))
                                        if true {
                                            if true {
                                                v4 = append(v4, fi91)
                                            }
                                            switch "a" {
                                            case "\n\t":
                                                switch ("foo bar" + "\n\t") {
                                                case "\n\t":
                                                    v96 := GenS3{F0: 42, F1: v28, F2: "a"}
                                                    _ = v96
                                                case "hello":
                                                    var v97 int64 = -413
                                                    _ = v97
                                                case "foo bar":
                                                    if true {
                                                        var v98 float32 = (-2.25 / 1.5)
                                                        _ = v98
                                                    }
                                                }
                                            case "a":
                                                fmt.Println(v2, v3)
                                            case "hello":
                                                v99 := GenS11{F0: 0, F1: ("a" == ""), F2: 406}
                                                _ = v99
                                            default:
                                                switch (fi89 / ((1) | 1)) {
                                                case 1:
                                                    v100 := GenS7{F0: (v32 - 255), F1: (v2 >> ((42) & 63)), F2: (-(v3))}
                                                    _ = v100
                                                case 4294967295:
                                                    if ((v32 ^ (v32 / ((820) | 1))) >= v32) {
                                                        var v101 uint32 = v90
                                                        v3 -= 2147483647
                                                        var v102 int32 = 0
                                                        _ = v101
                                                        _ = v102
                                                    } else {
                                                        var v103 float64 = (v95 / 1.5)
                                                        var v104 uint64 = 42
                                                        _ = v103
                                                        _ = v104
                                                    }
                                                }
                                            }
                                            v28 = v28
                                        } else {
                                            var v105 bool = true
                                            v106 := GenS8{F0: ("\n\t" < (("" + "") + "\n\t")), F1: 42}
                                            _ = v105
                                            _ = v106
                                        }
                                        var v107 uint64 = 849
                                        _ = v107
                                        _ = v95
                                    }
                                    if (18446744073709551615 > 1) {
                                        v108 := GenS12{F0: 42, F1: 781, F2: int64(v1)}
                                        _ = v108
                                    } else {
                                        var v109 int = fi89
                                        var v110 float32 = 1.5
                                        _ = v109
                                        _ = v110
                                    }
                                } else {
                                    var v111 float64 = 1.5
                                    v112 := GenS1{F0: v28, F1: (!true)}
                                    _ = v111
                                    _ = v112
                                }
                            case "foo bar":
                                switch fi89 {
                                case 2147483647:
                                    var v113 int = int(v3)
                                    _ = v113
                                }
                            case "hello":
                                v4 = append(v4, fi89)
                            default:
                                var v114 int = (fi89 & (fi91 >> ((0) & 63)))
                                _ = v114
                            }
                        }
                        var v115 bool = ("a" > (("" + "foo bar") + ("a" + "")))
                        var v116 int32 = v3
                        _ = v115
                        _ = v116
                    default:
                        switch (-(fi89)) {
                        case 32768:
                            var v117 uint64 = 0
                            _ = v117
                        case 1:
                            var v118 uint64 = 18446744073709551615
                            _ = v118
                        case 2147483647:
                            v90 = 4294967295
                        }
                    }
                } else {
                    if false {
                        fmt.Println(v28, v5, v5)
                        var v119 bool = true
                        _ = v119
                    }
                }
                v1 = v1
                _ = fi91
            }
            for fi120 := 0; fi120 < 5; fi120++ {
                switch (^(fi120)) {
                case 6:
                    v121 := GenS13{F0: -1, F1: 128, F2: (v3 ^ v3)}
                    _ = v121
                case 4294967295:
                    var v122 int64 = 9223372036854775807
                    _ = v122
                case 127:
                    v123 := GenS14{F0: v90, F1: ("\n\t" + "hello")}
                    _ = v123
                }
                _ = fi120
            }
        }
        _ = fi89
        _ = v90
    }
    for fi124 := 0; fi124 < 7; fi124++ {
        fmt.Println(fi124, v32, v1)
        switch (fi124 & fi124) {
        case 2147483647:
            v3 += -2147483648
        case -2147483648:
            v125 := GenS2{F0: -2147483648, F1: "foo bar"}
            _ = v125
        case -9223372036854775808:
            v28 = (v28 & -947)
        }
        fmt.Println(v1, v28)
        _ = fi124
    }
    switch (("foo bar" + ("" + "\n\t")) + "a") {
    case "":
        var v126 float64 = float64(0.1)
        _ = v126
    }
    var v127 float64 = 3.14159
    v1 += v1
    var v128 float64 = (float64(-2.25) / 0.1)
    if (2147483648 != 4294967295) {
        for fi129 := 0; fi129 < 16; fi129++ {
            if (!("a" != "hello")) {
                if (1 >= 42) {
                    for fi130 := 0; fi130 < 13; fi130++ {
                        v3 -= v3
                        var v131 float32 = 0.1
                        v1 += v1
                        _ = fi130
                        _ = v131
                    }
                    var v132 bool = (v1 < (v1 % ((v1) | 1)))
                    _ = v132
                }
                if true {
                    var v133 int8 = (v2 & (v2 / (((-(v2))) | 1)))
                    if (!((v1 | v1) >= 0)) {
                        v134 := GenS15{F0: (v5 << ((2) & 63)), F1: v28, F2: true}
                        v3 -= v3
                        var v135 int64 = 1
                        _ = v134
                        _ = v135
                    } else {
                        var v136 bool = false
                        switch (fi129 + -9223372036854775808) {
                        case 0:
                            for fi137 := 0; fi137 < 13; fi137++ {
                                v1 = (v1 * (v1 >> ((genFn1(4294967295, 2)) & 63)))
                                v138 := GenS16{F0: 127}
                                _ = fi137
                                _ = v138
                            }
                        }
                        fmt.Println(v32, v32, v136)
                        _ = v136
                    }
                    _ = v133
                } else {
                    v139 := GenS6{F0: 42}
                    v128 = (v128 + float64(1.5))
                    var v140 int32 = 42
                    _ = v139
                    _ = v140
                }
                if true {
                    v4 = append(v4, fi129)
                }
            } else {
                v128 += v128
                v4 = append(v4, 32767)
            }
            _ = fi129
        }
        if false {
            var v141 int8 = int8(v1)
            var v142 uint64 = 2
            v127 = -2.25
            _ = v141
            _ = v142
        } else {
            if false {
                var v143 float64 = (0.1 - (v128 * v127))
                _ = v143
            } else {
                for fi144 := 0; fi144 < 9; fi144++ {
                    for fi145 := 0; fi145 < 8; fi145++ {
                        v2 = 127
                        v3 -= 130
                        if false {
                            switch (fi145 | (-(fi145))) {
                            case 0:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case -9223372036854775808:
                                v128 = float64((3.14159 - 0.1))
                            case 9223372036854775807:
                                var v146 int64 = 9223372036854775807
                                _ = v146
                            }
                        }
                        if true {
                            var v147 uint32 = 65535
                            var v148 string = "\n\t"
                            _ = v147
                            _ = v148
                        } else {
                            var v149 float32 = 0.0
                            var v150 uint32 = 71
                            _ = v149
                            _ = v150
                        }
                        var v151 int64 = 42
                        switch -2147483648 {
                        case 128:
                            var v152 uint32 = 2
                            _ = v152
                        case 469:
                            if false {
                                v32 = v32
                            } else {
                                v151 = (v151 & 42)
                            }
                        default:
                            var v153 int8 = (v2 + 0)
                            _ = v153
                        }
                        var v154 uint64 = uint64(v1)
                        v154 = v154
                        v154 = v154
                        fmt.Println(v32)
                        var v155 float64 = (v127 * (((0.1 / (v128 + 3.14159)) * 3.14159) * ((((v127 - (((v128 + 3.14159) - v127) * (float64(0.0) / (v128 - v128)))) + v128) * (1.5 * 3.14159)) - float64(1.5))))
                        if len(v4) > 0 {
                            v4[0] = (fi144 << ((430) & 63))
                        }
                        _ = fi145
                        _ = v151
                        _ = v154
                        _ = v155
                    }
                    v28 -= -1
                    if true {
                        v5 = v5
                    } else {
                        for fi156 := 0; fi156 < 10; fi156++ {
                            switch (fi144 ^ (fi144 ^ (fi156 & fi156))) {
                            case 472:
                                _ = len(v4)
                            default:
                                v1 = (v1 | v1)
                            }
                            fmt.Println(fi156)
                            _ = fi156
                        }
                    }
                    _ = fi144
                }
                v127 = 0.1
            }
            if (true || true) {
                var v157 float32 = (float32(v127) - 0.0)
                _ = v157
            }
        }
    } else {
        v158 := GenS17{F0: -608, F1: v32, F2: v1}
        var v159 uint = 134
        switch "a" {
        case "a":
            switch 0 {
            case 929:
                var v160 int8 = v2
                _ = v160
            }
        case "foo bar":
            var v161 int8 = v5
            _ = v161
        default:
            v128 += ((v127 * 0.0) + v128)
        }
        _ = v158
        _ = v159
    }
    fmt.Println("v1:", v1)
    fmt.Println("v127:", v127)
    fmt.Println("v128:", v128)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v127
    _ = v128
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
