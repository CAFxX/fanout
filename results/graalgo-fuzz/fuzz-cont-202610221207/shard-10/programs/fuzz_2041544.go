package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 uint8; F2 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int; F1 uint64 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint8; F1 int16; F2 bool }
type GenS14 struct { F0 bool; F1 int64; F2 bool }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int32; F1 int32 }
func genFn0(p0 uint8, p1 string, p2 int) int16 {
    v5 := []int{(p2 - (p2 * p2)), (p2 | p2), p2, p2}
    v5 = append(v5, (p2 % ((p2) | 1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 722
}
func genFn1(p0 uint8) uint {
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            if (!true) {
                if ((!false) && false) {
                    p0 -= p0
                    p0 = p0
                    var v8 int32 = 42
                    v9 := GenS1{F0: -128}
                    if (42 > 2) {
                        p0 = (p0 - p0)
                    }
                    _ = v8
                    _ = v9
                } else {
                    if true {
                        if (p0 >= 42) {
                            fmt.Println(fi7, fi7, p0)
                            switch "foo bar" {
                            case "hello":
                                v10 := []int{(fi7 << ((65535) & 63)), fi7, fi6, 478}
                                _ = v10
                            case "\n\t":
                                var v11 bool = false
                                _ = v11
                            case "a":
                                var v12 string = "a"
                                _ = v12
                            }
                            var v13 float64 = (float64(((0.0 * float32(0.0)) * (1.5 / 1.5))) * 1e10)
                            fmt.Println(fi7, fi7)
                            var v14 int = -1
                            v15 := GenS2{F0: 2}
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                        fmt.Println(fi7, p0)
                        fmt.Println(fi7)
                    } else {
                        v16 := []int{fi7, fi6, fi7, -9223372036854775808}
                        if true {
                            fmt.Println(p0, p0)
                        }
                        if (128 > (p0 ^ p0)) {
                            fmt.Println(fi7, fi6)
                            switch ("" + "foo bar") {
                            case "":
                                p0 = p0
                            case "foo bar":
                                var v17 uint8 = p0
                                v18 := GenS1{F0: -8}
                                v17 += v17
                                _ = v17
                                _ = v18
                            case "hello":
                                fmt.Println(p0, fi6)
                            default:
                                if len(v16) > 0 {
                                    v16[0] = fi6
                                }
                            }
                            var v19 uint = 1
                            _ = v19
                        }
                        _ = v16
                    }
                    var v20 uint = 18446744073709551615
                    _ = v20
                }
            } else {
                switch "foo bar" {
                case "a":
                    p0 = 0
                case "foo bar":
                    var v21 uint32 = 2
                    _ = v21
                case "\n\t":
                    fmt.Println(fi6, fi7, fi6)
                default:
                    var v22 int8 = 1
                    v22 = (-(v22))
                    if true {
                        switch ("\n\t" + ("\n\t" + "foo bar")) {
                        case "":
                            var v23 uint16 = 970
                            _ = v23
                        case "hello":
                            var v24 int8 = v22
                            _ = v24
                        case "\n\t":
                            p0 = (p0 + (p0 | p0))
                        }
                    } else {
                        fmt.Println(fi7)
                        p0 = (p0 & 2)
                    }
                    _ = v22
                }
            }
            v25 := GenS1{F0: 100}
            switch ("foo bar" + ("" + "\n\t")) {
            case "\n\t":
                v26 := []int{}
                _ = v26
            case "hello":
                var v27 uint8 = 77
                _ = v27
            }
            _ = fi7
            _ = v25
        }
        p0 -= p0
        _ = fi6
    }
    var v28 int64 = 0
    v28 -= (^(v28))
    p0 -= p0
    var v29 uint64 = 42
    _ = p0
    _ = v28
    _ = v29
    return 698
}

func main() {
    var v1 int16 = -32768
    var v2 uint = 42
    var v3 bool = false
    v4 := []uint32{}
    if v3 {
        v3 = true
    }
    if ("foo bar" <= "\n\t") {
        if false {
            var v30 bool = v3
            var v31 uint8 = 42
            _ = v30
            _ = v31
        }
        v3 = false
    } else {
        v2 = v2
        v2 -= genFn1(25)
        switch "a" {
        case "a":
            var v32 int = 4294967295
            _ = v32
        case "\n\t":
            if v3 {
                _ = len(v4)
                if v3 {
                    fmt.Println(v3, v3, v1)
                }
                if (!false) {
                    fmt.Println(v3, v2)
                    var v33 float64 = -2.25
                    _ = v33
                }
            } else {
                fmt.Println(v2, v3, v3)
                var v34 string = "\n\t"
                _ = v34
            }
        }
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        v36 := GenS1{F0: 1}
        switch fi35 {
        case 9223372036854775807:
            switch ("foo bar" + (("a" + "a") + ("a" + (("" + "") + "\n\t")))) {
            case "a":
                v37 := GenS1{F0: 42}
                _ = v37
            }
        default:
            switch "a" {
            case "hello":
                fmt.Println(v2, fi35)
            default:
                v38 := GenS2{F0: 560}
                _ = v38
            }
        }
        v39 := GenS2{F0: 32767}
        _ = fi35
        _ = v36
        _ = v39
    }
    v3 = (false || (127 == 100))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 416 {
    case -170:
        v3 = v3
    case -2147483648:
        var v40 uint8 = 128
        _ = v40
    case 0:
        if (v3 && (42 != 4294967295)) {
            if v3 {
                var v41 float32 = 0.1
                if ((float32(0.1) / 1e10) < 1.5) {
                    var v42 int16 = (v1 + v1)
                    _ = v42
                } else {
                    if (!v3) {
                        var v43 float64 = 1e10
                        _ = v43
                    } else {
                        var v44 uint8 = 19
                        v45 := GenS2{F0: 912}
                        if (478 >= 977) {
                            if (v3 || (!(v44 < 43))) {
                                v1 += (v1 ^ v1)
                                if len(v4) > 0 {
                                    v4[0] = 0
                                }
                            } else {
                                var v46 uint16 = 255
                                _ = len(v4)
                                _ = v46
                            }
                            v41 -= v41
                            v47 := GenS2{F0: 32767}
                            _ = v47
                        }
                        _ = v44
                        _ = v45
                    }
                }
                v4 = append(v4, 2)
                _ = v41
            }
            v48 := GenS3{F0: 2147483647}
            var v49 uint16 = 65535
            _ = v48
            _ = v49
        } else {
            if (1 >= 42) {
                var v50 uint = (v2 * (^(v2)))
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v50
            }
            v1 -= (v1 & v1)
            v51 := GenS4{F0: 127}
            _ = v51
        }
    }
    v52 := GenS4{F0: 127}
    v2 = v2
    if v3 {
        for fi53 := 0; fi53 < 20; fi53++ {
            switch "hello" {
            case "foo bar":
                v3 = v3
            default:
                v2 = 1
            }
            _ = fi53
        }
    } else {
        var v54 uint64 = 328
        var v55 int = -962
        if v3 {
            v1 -= 32767
            if (float32((float64(float32(0.1)) - float64(-2.25))) >= 0.0) {
                if (v3 || v3) {
                    v56 := GenS3{F0: 1}
                    switch v55 {
                    case 2147483647:
                        _ = len(v4)
                    case -1:
                        var v57 float64 = 0.1
                        _ = v57
                    case 787:
                        v58 := GenS1{F0: 100}
                        _ = v58
                    }
                    _ = v56
                } else {
                    v1 += v1
                }
                if v3 {
                    fmt.Println(v3)
                    if (32767 > 1) {
                        v4 = append(v4, 992)
                        var v59 int8 = -127
                        _ = v59
                    } else {
                        v60 := GenS1{F0: 8}
                        _ = v60
                    }
                } else {
                    if len(v4) > 0 {
                        v4[0] = 2147483647
                    }
                }
            } else {
                for fi61 := 0; fi61 < 8; fi61++ {
                    v54 -= (v54 << ((v2) & 63))
                    v4 = append(v4, 42)
                    _ = fi61
                }
                var v62 uint16 = 538
                _ = v62
            }
            fmt.Println(v3, v55, v54)
        } else {
            v63 := GenS1{F0: 50}
            v64 := GenS5{F0: v55, F1: 42, F2: 1}
            _ = v63
            _ = v64
        }
        _ = v54
        _ = v55
    }
    var v65 uint = v2
    var v66 int8 = -1
    v1 = (v1 & (v1 >> ((genFn1(246)) & 63)))
    for fi67 := 0; fi67 < 1; fi67++ {
        v68 := GenS6{F0: 63}
        switch ("a" + "") {
        case "a":
            for fi69 := 0; fi69 < 12; fi69++ {
                v2 = 18446744073709551615
                _ = fi69
            }
        }
        v70 := GenS7{F0: 72, F1: 1, F2: 32767}
        _ = fi67
        _ = v68
        _ = v70
    }
    if v3 {
        v65 = (^(v65))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = len(v4)
    }
    for fi71 := 0; fi71 < 11; fi71++ {
        switch fi71 {
        case 0:
            if (genFn0(114, "foo bar", -829) <= 173) {
                var v72 int32 = 1
                var v73 uint64 = 0
                v3 = (v3 || (uint8(v65) >= 1))
                _ = v72
                _ = v73
            } else {
                for fi74 := 0; fi74 < 11; fi74++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v1 -= (v1 - v1)
                    v75 := GenS8{F0: 255, F1: (v66 * (v66 - -128))}
                    _ = fi74
                    _ = v75
                }
                if v3 {
                    for fi76 := 0; fi76 < 8; fi76++ {
                        if v3 {
                            var v77 float64 = (0.1 - 0.1)
                            switch "foo bar" {
                            case "\n\t":
                                v3 = (1 != 1)
                            case "foo bar":
                                _ = len(v4)
                            case "hello":
                                var v78 uint16 = 32767
                                _ = v78
                            default:
                                _ = len(v4)
                            }
                            v77 -= (float64(0.1) / -2.25)
                            _ = v77
                        }
                        var v79 uint32 = 42
                        var v80 float64 = 1.5
                        _ = fi76
                        _ = v79
                        _ = v80
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    var v81 uint32 = 2147483648
                    fmt.Println(v65, v66)
                    _ = v81
                }
                switch fi71 {
                case 127:
                    switch "a" {
                    case "hello":
                        v3 = (!true)
                    case "\n\t":
                        v82 := GenS4{F0: 1}
                        _ = v82
                    }
                case 42:
                    v3 = ((373 > 32767) || v3)
                case 0:
                    fmt.Println(v65, v65)
                }
            }
        case -168:
            v66 = int8(v1)
        case 4294967295:
            switch (-(fi71)) {
            case 32767:
                var v83 uint16 = 65535
                _ = v83
            }
        }
        _ = len(v4)
        _ = fi71
    }
    switch "\n\t" {
    case "foo bar":
        if (!false) {
            _ = len(v4)
            for fi84 := 0; fi84 < 3; fi84++ {
                v1 = -75
                if v3 {
                    if v3 {
                        fmt.Println(v66)
                        if false {
                            v85 := GenS4{F0: 204}
                            v86 := GenS1{F0: (v66 ^ v66)}
                            var v87 int32 = 2147483647
                            _ = v85
                            _ = v86
                            _ = v87
                        } else {
                            switch (fi84 % ((fi84) | 1)) {
                            case 9223372036854775807:
                                switch ("a" + "hello") {
                                case "":
                                    switch (fi84 / ((42) | 1)) {
                                    case 1:
                                        switch (fi84 >> (((v65 >> (((v65 >> ((v2) & 63))) & 63))) & 63)) {
                                        case 127:
                                            if ((!false) && true) {
                                                var v88 int32 = 0
                                                _ = v88
                                            } else {
                                                var v89 bool = ("hello" >= "hello")
                                                for fi90 := 0; fi90 < 6; fi90++ {
                                                    switch 32768 {
                                                    case 128:
                                                        v4 = append(v4, 65535)
                                                    case -1:
                                                        if v89 {
                                                            if (506 != 42) {
                                                                fmt.Println(v65)
                                                                fmt.Println(fi90, v89)
                                                            }
                                                            var v91 string = "foo bar"
                                                            switch ("\n\t" + v91) {
                                                            case "hello":
                                                                v65 -= (^(v65))
                                                            case "\n\t":
                                                                v4 = append(v4, 688)
                                                            case "":
                                                                var v92 uint8 = 42
                                                                _ = v92
                                                            default:
                                                                if v3 {
                                                                    if (v3 && (v66 >= -1)) {
                                                                        if (v2 <= (^(v65))) {
                                                                            v1 += genFn0(255, "hello", (fi90 + (fi84 << ((42) & 63))))
                                                                        }
                                                                    } else {
                                                                        if (145 != 0) {
                                                                            v4 = append(v4, 2)
                                                                        }
                                                                        var v93 int = (fi90 >> (((v2 * (v2 ^ (v65 << ((v2) & 63))))) & 63))
                                                                        fmt.Println(v66, v1, v91)
                                                                        _ = v93
                                                                    }
                                                                    if (v3 && true) {
                                                                        var v94 float32 = float32(0.1)
                                                                        _ = v94
                                                                    }
                                                                    if true {
                                                                        fmt.Println(v2)
                                                                        switch 0 {
                                                                        case 42:
                                                                            v95 := GenS5{F0: (fi84 - -700), F1: 1, F2: 869}
                                                                            _ = v95
                                                                        default:
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                        }
                                                                        switch (^(fi84)) {
                                                                        case 915:
                                                                            var v96 uint64 = 18446744073709551615
                                                                            _ = v96
                                                                        case 127:
                                                                            v1 += v1
                                                                        default:
                                                                            switch (^(fi90)) {
                                                                            case -2147483648:
                                                                                var v97 uint16 = 237
                                                                                _ = v97
                                                                            case 128:
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            case -919:
                                                                                v98 := GenS9{F0: 65535}
                                                                                _ = v98
                                                                            default:
                                                                                var v99 uint8 = 249
                                                                                _ = v99
                                                                            }
                                                                            v1 += v1
                                                                            var v100 float32 = float32(1e10)
                                                                            _ = v100
                                                                        }
                                                                        v66 = -128
                                                                    }
                                                                } else {
                                                                    v66 = (v66 / ((100) | 1))
                                                                }
                                                            }
                                                            _ = v91
                                                        } else {
                                                            if len(v4) > 0 {
                                                                v4[0] = 1
                                                            }
                                                            v101 := GenS10{F0: 1}
                                                            if v3 {
                                                                v102 := GenS10{F0: 2}
                                                                _ = v102
                                                            }
                                                            _ = v101
                                                        }
                                                    default:
                                                        v103 := GenS11{F0: 42, F1: 2}
                                                        _ = v103
                                                    }
                                                    _ = fi90
                                                }
                                                _ = v89
                                            }
                                        case 0:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                    case 128:
                                        v104 := GenS3{F0: -221}
                                        _ = v104
                                    default:
                                        var v105 int8 = (v66 ^ 0)
                                        _ = v105
                                    }
                                case "foo bar":
                                    fmt.Println(v65)
                                case "\n\t":
                                    v4 = append(v4, 42)
                                default:
                                    if ((0 < 2) || v3) {
                                        if v3 {
                                            if ((v1 <= (v1 & (v1 / (((v1 << ((v2) & 63))) | 1)))) || (0.0 == 3.14159)) {
                                                for fi106 := 0; fi106 < 13; fi106++ {
                                                    var v107 float64 = ((0.1 + -2.25) * float64(((float32(1e10) - (((-2.25 * -2.25) - float32(0.0)) / 3.14159)) + 3.14159)))
                                                    _ = fi106
                                                    _ = v107
                                                }
                                            } else {
                                                var v108 int8 = (v66 | v66)
                                                v4 = append(v4, 65535)
                                                _ = v108
                                            }
                                        }
                                    }
                                }
                            default:
                                for fi109 := 0; fi109 < 19; fi109++ {
                                    switch (fi84 / ((fi109) | 1)) {
                                    case 95:
                                        v65 -= v65
                                    case -1:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    default:
                                        v66 += (v66 << (((v2 << ((v65) & 63))) & 63))
                                    }
                                    switch (fi109 | fi109) {
                                    case -106:
                                        if (775 != 18446744073709551615) {
                                            var v110 int32 = -162
                                            if ("hello" >= ("a" + "hello")) {
                                                if ((2147483647 >= 2147483648) || true) {
                                                    v111 := GenS12{F0: -816}
                                                    var v112 int16 = v1
                                                    var v113 int64 = -394
                                                    _ = v111
                                                    _ = v112
                                                    _ = v113
                                                }
                                                if false {
                                                    switch fi109 {
                                                    case 2147483647:
                                                        if len(v4) > 0 {
                                                            v4[0] = 4294967295
                                                        }
                                                    }
                                                    fmt.Println(v1, fi109, v3)
                                                    if v3 {
                                                        var v114 uint = 4294967295
                                                        var v115 int8 = -86
                                                        if (1.5 > 1.5) {
                                                            var v116 int8 = v66
                                                            _ = v116
                                                        }
                                                        v66 = 1
                                                        v117 := GenS11{F0: (fi84 & (fi84 + 32768)), F1: 1}
                                                        _ = v114
                                                        _ = v115
                                                        _ = v117
                                                    } else {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        if v3 {
                                                            v2 = genFn1(128)
                                                            fmt.Println(v2)
                                                            fmt.Println(v66)
                                                            var v118 bool = v3
                                                            _ = v118
                                                        } else {
                                                            if (776 <= uint16(v2)) {
                                                                v2 = (v2 | (v65 ^ 135))
                                                            }
                                                        }
                                                    }
                                                }
                                                var v119 float32 = 0.0
                                                _ = v119
                                            } else {
                                                var v120 string = ("\n\t" + "foo bar")
                                                switch (fi84 >> ((v2) & 63)) {
                                                case 9223372036854775807:
                                                    switch v120 {
                                                    case "hello":
                                                        var v121 int32 = (v110 >> ((v2) & 63))
                                                        _ = v121
                                                    case "foo bar":
                                                        v1 += (v1 ^ v1)
                                                    }
                                                default:
                                                    v3 = v3
                                                }
                                                var v122 uint16 = 358
                                                var v123 uint64 = 1
                                                _ = v120
                                                _ = v122
                                                _ = v123
                                            }
                                            _ = v110
                                        } else {
                                            if v3 {
                                                v65 += genFn1(114)
                                            }
                                        }
                                        var v124 uint = (v2 + (v65 & 0))
                                        v2 = v2
                                        _ = v124
                                    case -642:
                                        var v125 int = 32767
                                        switch (("" + ("hello" + "a")) + "") {
                                        case "\n\t":
                                            fmt.Println(v3)
                                        case "foo bar":
                                            switch "" {
                                            case "a":
                                                v2 = (v65 ^ v2)
                                            case "foo bar":
                                                v1 -= (v1 ^ v1)
                                                v125 = (v125 % ((-1) | 1))
                                                v126 := GenS7{F0: 18446744073709551615, F1: 42, F2: genFn0(42, "a", fi84)}
                                                _ = v126
                                            case "":
                                                switch ("a" + "\n\t") {
                                                case "\n\t":
                                                    v3 = ((560 == 977) || (fi84 != fi109))
                                                    var v127 int32 = -90
                                                    if v3 {
                                                        var v128 int64 = 0
                                                        var v129 int32 = (v127 - (v127 / (((v127 << (((v2 ^ genFn1(128))) & 63))) | 1)))
                                                        _ = v128
                                                        _ = v129
                                                    } else {
                                                        v130 := GenS13{F0: 79, F1: v1, F2: false}
                                                        v2 += genFn1(128)
                                                        _ = v130
                                                    }
                                                    _ = v127
                                                default:
                                                    if v3 {
                                                        v131 := GenS14{F0: (!v3), F1: 42, F2: false}
                                                        v132 := GenS5{F0: (fi109 % ((fi109) | 1)), F1: 2, F2: 2}
                                                        v65 = (v65 % ((genFn1(73)) | 1))
                                                        switch (^(v125)) {
                                                        case -644:
                                                            v3 = v3
                                                            _ = len(v4)
                                                        case 210:
                                                            var v133 float64 = 1.5
                                                            _ = v133
                                                        case 42:
                                                            v2 = 1
                                                        }
                                                        _ = v131
                                                        _ = v132
                                                    }
                                                }
                                            }
                                        case "hello":
                                            _ = len(v4)
                                        }
                                        _ = v125
                                    case 128:
                                        v134 := GenS1{F0: -1}
                                        _ = v134
                                    default:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    if (v3 && (!v3)) {
                                        var v135 uint16 = 65535
                                        v2 = genFn1(128)
                                        var v136 int = 29
                                        _ = v135
                                        _ = v136
                                    }
                                    switch "a" {
                                    case "a":
                                        fmt.Println(v65, v65)
                                    case "hello":
                                        _ = len(v4)
                                        v4 = append(v4, 1)
                                        switch (-(fi109)) {
                                        case 593:
                                            if (!v3) {
                                                var v137 string = "hello"
                                                var v138 uint = v2
                                                var v139 float32 = 1e10
                                                _ = v137
                                                _ = v138
                                                _ = v139
                                            } else {
                                                v1 -= int16(v2)
                                                v140 := GenS5{F0: (fi109 >> ((v2) & 63)), F1: 2, F2: 42}
                                                _ = v140
                                            }
                                            fmt.Println(fi109, v66, fi84)
                                        case -434:
                                            _ = len(v4)
                                        case 42:
                                            var v141 int32 = 1
                                            _ = v141
                                        default:
                                            var v142 uint32 = 0
                                            _ = v142
                                        }
                                    case "foo bar":
                                        switch (fi109 * (fi109 | fi109)) {
                                        case 42:
                                            fmt.Println(fi109)
                                        }
                                    default:
                                        if (0 != 127) {
                                            v143 := GenS12{F0: -952}
                                            switch "a" {
                                            case "":
                                                fmt.Println(v2, v3)
                                            }
                                            _ = v143
                                        } else {
                                            v2 = v2
                                            var v144 bool = (255 == 255)
                                            if len(v4) > 0 {
                                                v4[0] = 42
                                            }
                                            _ = v144
                                        }
                                    }
                                    fmt.Println(v3)
                                    _ = fi109
                                }
                            }
                            for fi145 := 0; fi145 < 1; fi145++ {
                                v146 := GenS15{F0: 3}
                                _ = fi145
                                _ = v146
                            }
                        }
                        if (-1 < 0) {
                            var v147 int32 = 42
                            for fi148 := 0; fi148 < 11; fi148++ {
                                var v149 uint32 = 486
                                v147 = (v147 * -957)
                                _ = fi148
                                _ = v149
                            }
                            _ = v147
                        }
                    } else {
                        for fi150 := 0; fi150 < 6; fi150++ {
                            v65 -= v2
                            _ = fi150
                        }
                    }
                    var v151 uint = 1
                    v152 := GenS16{F0: -1, F1: 1}
                    _ = v151
                    _ = v152
                }
                for fi153 := 0; fi153 < 18; fi153++ {
                    var v154 uint32 = 126
                    _ = fi153
                    _ = v154
                }
                _ = fi84
            }
            v65 = (v2 ^ (v65 % ((v2) | 1)))
        }
    default:
        if ((-506 <= -684) && (!(v65 < (v2 & v2)))) {
            v2 = (v65 - (v2 % ((1) | 1)))
            var v155 string = "foo bar"
            if (-1 >= 0) {
                fmt.Println(v3, v1, v2)
                v4 = append(v4, 2147483647)
            }
            _ = v155
        } else {
            var v156 uint8 = 2
            v157 := GenS13{F0: 255, F1: (v1 * (v1 >> ((0) & 63))), F2: v3}
            _ = v156
            _ = v157
        }
    }
    v3 = (128 != 128)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v65
    _ = v66
}
