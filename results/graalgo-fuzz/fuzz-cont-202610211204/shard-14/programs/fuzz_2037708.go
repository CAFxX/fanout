package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 int8; F1 int32; F2 string }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 uint64; F1 int; F2 int64 }
type GenS8 struct { F0 uint32; F1 int32 }
type GenS9 struct { F0 uint; F1 uint }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int8; F1 int }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int8; F1 string; F2 bool }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 int64; F1 uint8; F2 int64 }
type GenS19 struct { F0 int64; F1 int64 }
type GenS20 struct { F0 uint32; F1 int32 }

func main() {
    var v1 int64 = 771
    var v2 uint64 = 18446744073709551615
    v3 := []uint32{2147483647, 476, 2147483647}
    var v4 float32 = 1.5
    switch 127 {
    case -2147483648:
        v1 = 1
    }
    switch -963 {
    case 1:
        v5 := GenS1{F0: 345}
        _ = v5
    case 42:
        if true {
            switch "" {
            case "hello":
                v2 += v2
            }
            var v6 bool = false
            var v7 int32 = 0
            _ = v6
            _ = v7
        } else {
            var v8 uint = 362
            _ = v8
        }
    case -781:
        switch "foo bar" {
        case "hello":
            if ((-128 == -19) && false) {
                for fi9 := 0; fi9 < 6; fi9++ {
                    v2 = v2
                    if (933 <= 2147483647) {
                        var v10 int8 = 100
                        _ = v10
                    }
                    _ = fi9
                }
                var v11 uint8 = 127
                _ = v11
            }
        case "":
            v12 := GenS1{F0: 186}
            _ = v12
        case "\n\t":
            var v13 uint = 18446744073709551615
            _ = v13
        default:
            switch 533 {
            case 1:
                if false {
                    if (0 < 1) {
                        v14 := GenS1{F0: 835}
                        _ = v14
                    }
                    if false {
                        for fi15 := 0; fi15 < 19; fi15++ {
                            if false {
                                var v16 int16 = -151
                                _ = v16
                            }
                            for fi17 := 0; fi17 < 14; fi17++ {
                                v2 -= 0
                                _ = fi17
                            }
                            var v18 string = "\n\t"
                            _ = fi15
                            _ = v18
                        }
                        if (300 < 881) {
                            switch 226 {
                            case 491:
                                switch "" {
                                case "":
                                    if false {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if ((false || true) || (true && true)) {
                                            v1 = (v1 % ((v1) | 1))
                                            v3 = append(v3, 620)
                                        } else {
                                            switch ((("hello" + "\n\t") + ("a" + "")) + "") {
                                            case "\n\t":
                                                for fi19 := 0; fi19 < 19; fi19++ {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v20 := GenS2{F0: false}
                                                    if len(v3) > 0 {
                                                        v3[0] = 98
                                                    }
                                                    _ = fi19
                                                    _ = v20
                                                }
                                            default:
                                                v1 -= v1
                                            }
                                            v3 = append(v3, 4294967295)
                                        }
                                    } else {
                                        var v21 float64 = -2.25
                                        switch "a" {
                                        case "hello":
                                            v3 = append(v3, 2147483648)
                                        case "foo bar":
                                            v22 := GenS1{F0: 42}
                                            _ = v22
                                        }
                                        _ = v21
                                    }
                                case "a":
                                    v1 = (v1 + v1)
                                }
                            default:
                                v4 -= v4
                            }
                            var v23 uint8 = 127
                            _ = v23
                        } else {
                            if false {
                                v24 := GenS2{F0: true}
                                _ = v24
                            } else {
                                var v25 string = (("foo bar" + "") + ((("foo bar" + (("" + "\n\t") + ("a" + ("" + "a")))) + "\n\t") + "a"))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v26 int32 = 242
                                _ = v25
                                _ = v26
                            }
                            for fi27 := 0; fi27 < 6; fi27++ {
                                for fi28 := 0; fi28 < 6; fi28++ {
                                    v4 = (1e10 + float32(1e10))
                                    _ = fi28
                                }
                                _ = fi27
                            }
                            v1 = 1
                        }
                        var v29 uint = 65535
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 8; fi30++ {
                        v2 -= v2
                        v31 := GenS1{F0: 370}
                        for fi32 := 0; fi32 < 4; fi32++ {
                            var v33 uint = 65535
                            v34 := GenS2{F0: (false || (false || false))}
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        _ = fi30
                        _ = v31
                    }
                } else {
                    _ = len(v3)
                }
            case -9223372036854775808:
                for fi35 := 0; fi35 < 9; fi35++ {
                    v36 := GenS1{F0: 720}
                    _ = fi35
                    _ = v36
                }
            case -1:
                for fi37 := 0; fi37 < 5; fi37++ {
                    switch (fi37 & fi37) {
                    case 32767:
                        var v38 float64 = 3.14159
                        _ = v38
                    }
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                    _ = fi37
                }
            default:
                for fi39 := 0; fi39 < 14; fi39++ {
                    switch "a" {
                    case "a":
                        if true {
                            _ = len(v3)
                        } else {
                            v2 = (v2 >> ((748) & 63))
                            for fi40 := 0; fi40 < 17; fi40++ {
                                switch fi39 {
                                case 0:
                                    v3 = append(v3, 4294967295)
                                }
                                _ = fi40
                            }
                        }
                    }
                    _ = fi39
                }
            }
        }
    default:
        if len(v3) > 0 {
            v3[0] = 2
        }
    }
    var v41 int64 = (-(v1))
    var v42 int8 = -1
    v42 += v42
    v43 := GenS3{F0: 42, F1: (-(v42))}
    if true {
        var v44 bool = false
        switch 127 {
        case 904:
            v2 = (v2 | 0)
        }
        var v45 uint64 = (v2 | (^(v2)))
        _ = v44
        _ = v45
    } else {
        var v46 uint = 255
        switch 1 {
        case -9223372036854775808:
            if true {
                for fi47 := 0; fi47 < 5; fi47++ {
                    var v48 int32 = 42
                    _ = fi47
                    _ = v48
                }
                var v49 uint = v46
                _ = v49
            }
        case -129:
            switch 1 {
            case 2147483647:
                var v50 int32 = 1
                _ = v50
            case 0:
                if true {
                    v51 := GenS4{F0: 2147483647, F1: 766}
                    for fi52 := 0; fi52 < 14; fi52++ {
                        _ = len(v3)
                        _ = fi52
                    }
                    _ = v51
                }
            case 32768:
                fmt.Println(v1, v4, v46)
            default:
                v4 = v4
            }
        case 704:
            fmt.Println(v2)
        }
        if (128 < 128) {
            v53 := GenS5{F0: (v42 | (v42 - v42)), F1: 42, F2: "\n\t"}
            v4 += (1e10 * ((0.0 / (v4 / v4)) + v4))
            _ = v53
        }
        _ = v46
    }
    var v54 uint32 = 2
    v54 = v54
    switch "hello" {
    case "\n\t":
        switch "hello" {
        case "":
            v55 := GenS6{F0: (v54 & v54), F1: v1}
            _ = v55
        case "hello":
            v56 := GenS6{F0: uint32(v2), F1: v1}
            _ = v56
        case "\n\t":
            v3 = append(v3, (^(v54)))
        default:
            fmt.Println(v41, v41, v4)
        }
    default:
        for fi57 := 0; fi57 < 1; fi57++ {
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                switch (("a" + "") + "") {
                case "a":
                    var v58 string = ("hello" + "foo bar")
                    _ = v58
                case "":
                    _ = len(v3)
                case "hello":
                    if true {
                        switch (-(fi57)) {
                        case -879:
                            var v59 int16 = 1000
                            _ = v59
                        }
                        fmt.Println(v41, v42, v4)
                    } else {
                        v60 := GenS4{F0: (v54 << ((18446744073709551615) & 63)), F1: uint16(v2)}
                        v1 = (^(v1))
                        if false {
                            if ((^(v54)) != 65535) {
                                v61 := GenS7{F0: 1, F1: fi57, F2: v1}
                                v1 -= v1
                                _ = v61
                            }
                            fmt.Println(v42, v54, v54)
                        } else {
                            var v62 float64 = (1e10 * 3.14159)
                            var v63 float64 = (v62 * v62)
                            _ = v62
                            _ = v63
                        }
                        _ = v60
                    }
                }
            }
            _ = fi57
        }
    }
    v64 := GenS8{F0: (v54 + v54), F1: 40}
    for fi65 := 0; fi65 < 6; fi65++ {
        if true {
            v66 := GenS9{F0: 1, F1: 771}
            _ = v66
        } else {
            if false {
                var v67 uint32 = 2
                if true {
                    if (true && true) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v67 = 208
                        if len(v3) > 0 {
                            v3[0] = (v67 / ((2) | 1))
                        }
                    } else {
                        v68 := GenS10{F0: true}
                        var v69 float64 = 0.1
                        _ = v68
                        _ = v69
                    }
                } else {
                    for fi70 := 0; fi70 < 20; fi70++ {
                        switch fi65 {
                        case -986:
                            if (("foo bar" + "a") > ("" + (("a" + (("a" + ((("hello" + "hello") + "") + "a")) + "hello")) + ("\n\t" + "hello")))) {
                                var v71 int8 = 16
                                v1 -= (v1 / ((v41) | 1))
                                _ = v71
                            }
                        }
                        if (210 > -32768) {
                            if false {
                                if len(v3) > 0 {
                                    v3[0] = (v67 << ((uint(v2)) & 63))
                                }
                                switch (^(fi70)) {
                                case -969:
                                    if ((^(v42)) <= (v42 ^ (v42 + v42))) {
                                        if (v41 == (v41 - v41)) {
                                            v72 := GenS11{F0: (((("foo bar" + (("a" + "") + "hello")) + "a") + "hello") + (((("" + ("a" + "\n\t")) + "") + (("\n\t" + "hello") + "hello")) + "hello"))}
                                            v42 = (v42 >> ((18446744073709551615) & 63))
                                            _ = v72
                                        } else {
                                            v54 = 65535
                                            v42 -= 42
                                            v73 := GenS8{F0: v67, F1: 1}
                                            fmt.Println(v4, v2, v54)
                                            var v74 int32 = -100
                                            if false {
                                                var v75 uint = 0
                                                v76 := GenS2{F0: true}
                                                var v77 uint64 = (^(v2))
                                                _ = v75
                                                _ = v76
                                                _ = v77
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = 2147483648
                                                }
                                            }
                                            _ = v73
                                            _ = v74
                                        }
                                        var v78 int64 = v1
                                        v42 += (-(v42))
                                        if false {
                                            switch "hello" {
                                            case "a":
                                                fmt.Println(v42, fi65, v67)
                                            case "\n\t":
                                                v54 = 243
                                            case "hello":
                                                if len(v3) > 0 {
                                                    v3[0] = (v54 * (v67 | (v67 - (^(v67)))))
                                                }
                                            default:
                                                v79 := GenS4{F0: v54, F1: 2}
                                                _ = v79
                                            }
                                            fmt.Println(v2, fi70, fi70)
                                            var v80 uint16 = 255
                                            switch "foo bar" {
                                            case "a":
                                                v3 = append(v3, v54)
                                            }
                                            v81 := GenS5{F0: -128, F1: -830, F2: (((("a" + "\n\t") + "foo bar") + ("a" + "hello")) + "hello")}
                                            fmt.Println(fi70, v78)
                                            v2 = uint64(v80)
                                            v82 := GenS3{F0: fi70, F1: v42}
                                            if (65535 != 980) {
                                                v4 += ((3.14159 + float32(3.14159)) + ((v4 + v4) * v4))
                                            }
                                            v4 = ((3.14159 + float32(0.1)) * (((v4 / 0.1) / (v4 + 1e10)) / v4))
                                            _ = v80
                                            _ = v81
                                            _ = v82
                                        }
                                        _ = v78
                                    }
                                case 889:
                                    fmt.Println(v54)
                                }
                                switch ("\n\t" + "hello") {
                                case "foo bar":
                                    _ = len(v3)
                                }
                            } else {
                                if false {
                                    var v83 int32 = -344
                                    var v84 uint64 = 120
                                    fmt.Println(v1)
                                    _ = v83
                                    _ = v84
                                } else {
                                    var v85 int32 = 1
                                    if len(v3) > 0 {
                                        v3[0] = (v54 ^ v54)
                                    }
                                    _ = v85
                                }
                                v42 = (v42 | (v42 | -128))
                            }
                            var v86 float32 = v4
                            _ = v86
                        }
                        _ = fi70
                    }
                    if ((^(v54)) > v54) {
                        var v87 int32 = -73
                        v88 := GenS11{F0: ""}
                        var v89 int8 = (v42 - (v42 + v42))
                        _ = v87
                        _ = v88
                        _ = v89
                    } else {
                        var v90 string = "hello"
                        var v91 int32 = -2147483648
                        var v92 float64 = 0.1
                        _ = v90
                        _ = v91
                        _ = v92
                    }
                }
                _ = v67
            } else {
                v42 -= v42
                if (true && false) {
                    fmt.Println(fi65)
                    if true {
                        if true {
                            var v93 int64 = v1
                            fmt.Println(v41, v93, v93)
                            _ = v93
                        } else {
                            fmt.Println(v2, v54)
                            if ((fi65 % ((fi65) | 1)) == (fi65 % ((fi65) | 1))) {
                                var v94 bool = false
                                if ((("foo bar" == "foo bar") && (-1 == -1)) && false) {
                                    v95 := GenS9{F0: 42, F1: 933}
                                    var v96 int16 = -32768
                                    _ = v95
                                    _ = v96
                                } else {
                                    switch ("\n\t" + "foo bar") {
                                    case "":
                                        v42 = -29
                                    case "foo bar":
                                        v4 += (-2.25 / v4)
                                    case "hello":
                                        var v97 uint32 = v54
                                        _ = v97
                                    }
                                }
                                _ = v94
                            }
                        }
                        for fi98 := 0; fi98 < 11; fi98++ {
                            v4 = float32(1.5)
                            v41 -= 9223372036854775807
                            if (!(int64(v54) >= (v1 >> ((0) & 63)))) {
                                var v99 int = (fi65 >> ((18446744073709551615) & 63))
                                fmt.Println(v41, v41, fi65)
                                _ = v99
                            } else {
                                if (v42 == v42) {
                                    var v100 float32 = v4
                                    if true {
                                        v101 := GenS9{F0: 612, F1: 1}
                                        v41 += -9223372036854775808
                                        _ = v101
                                    } else {
                                        var v102 float64 = -2.25
                                        _ = v102
                                    }
                                    _ = v100
                                }
                            }
                            v4 += v4
                            if ("hello" > "foo bar") {
                                fmt.Println(fi98, v42, v4)
                                var v103 int16 = 877
                                _ = v103
                            }
                            var v104 uint8 = 1
                            _ = fi98
                            _ = v104
                        }
                    } else {
                        v105 := GenS12{F0: (v42 / ((127) | 1)), F1: -9223372036854775808}
                        if (true && ((fi65 - fi65) <= fi65)) {
                            if (65535 == 2) {
                                for fi106 := 0; fi106 < 17; fi106++ {
                                    switch "\n\t" {
                                    case "a":
                                        v4 = v4
                                    default:
                                        v54 += v54
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = v54
                                    }
                                    switch "" {
                                    case "hello":
                                        switch (fi106 | (fi65 + (fi65 / ((4294967295) | 1)))) {
                                        case -845:
                                            var v107 uint32 = 0
                                            _ = v107
                                        }
                                    case "":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    default:
                                        switch 2147483647 {
                                        case -2147483648:
                                            v42 = v42
                                        case 32768:
                                            if true {
                                                if ((v54 | 2147483648) != v54) {
                                                    v108 := GenS13{F0: 128}
                                                    _ = v108
                                                } else {
                                                    fmt.Println(v54, v1)
                                                    switch fi65 {
                                                    case -9223372036854775808:
                                                        var v109 uint64 = 2
                                                        if false {
                                                            if true {
                                                                var v110 float64 = (0.1 * float64(((float32(1e10) - v4) - -2.25)))
                                                                _ = v110
                                                            } else {
                                                                if len(v3) > 0 {
                                                                    v3[0] = (v54 * (v54 * v54))
                                                                }
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                var v111 bool = (true && true)
                                                                fmt.Println(v41, fi106, fi106)
                                                                _ = v111
                                                            }
                                                            v54 = v54
                                                            var v112 float32 = v4
                                                            _ = v112
                                                        } else {
                                                            v41 -= 821
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                        if (507 == fi106) {
                                                            if false {
                                                                v42 = int8(v54)
                                                            } else {
                                                                v42 -= -50
                                                            }
                                                            v54 = v54
                                                        }
                                                        _ = v109
                                                    }
                                                }
                                                v42 -= v42
                                            } else {
                                                fmt.Println(v4)
                                                _ = len(v3)
                                                var v113 uint32 = v54
                                                _ = v113
                                            }
                                            var v114 int8 = (v42 + (^(v42)))
                                            _ = v114
                                        }
                                    }
                                    if true {
                                        v3 = append(v3, (v54 & (v54 | (v54 + v54))))
                                        v42 -= (v42 | 127)
                                        v41 -= v1
                                        if len(v3) > 0 {
                                            v3[0] = v54
                                        }
                                    } else {
                                        v115 := GenS4{F0: (v54 << ((18446744073709551615) & 63)), F1: 255}
                                        _ = v115
                                    }
                                    _ = fi106
                                }
                                v42 += 94
                                v116 := GenS13{F0: 0}
                                _ = v116
                            }
                        }
                        _ = v105
                    }
                    var v117 float64 = 1.5
                    _ = v117
                }
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch int(v42) {
        case -592:
            if (1 == 4294967295) {
                v1 -= (v1 << ((42) & 63))
            }
        case -9223372036854775808:
            v1 -= v41
        case 128:
            for fi118 := 0; fi118 < 19; fi118++ {
                v119 := GenS3{F0: (fi118 ^ (fi65 / ((fi118) | 1))), F1: 42}
                switch fi118 {
                case 128:
                    if false {
                        _ = len(v3)
                        v120 := GenS14{F0: (v42 >> ((2) & 63)), F1: "", F2: false}
                        if len(v3) > 0 {
                            v3[0] = v54
                        }
                        _ = v120
                    } else {
                        v121 := GenS1{F0: uint16(v54)}
                        switch "\n\t" {
                        case "\n\t":
                            v122 := GenS14{F0: (v42 | v42), F1: "\n\t", F2: (0 > 47)}
                            _ = v122
                        case "foo bar":
                            v41 = v41
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (false || (false && ("foo bar" < "foo bar"))) {
                            var v123 int16 = 42
                            v41 = 0
                            if true {
                                if (224 != 128) {
                                    v41 -= v41
                                    fmt.Println(v1, fi118, v41)
                                    fmt.Println(v2)
                                    v3 = append(v3, (v54 << ((268) & 63)))
                                } else {
                                    fmt.Println(fi65, v54)
                                    v54 -= v54
                                }
                            } else {
                                v124 := GenS15{F0: 65535}
                                if len(v3) > 0 {
                                    v3[0] = v54
                                }
                                if true {
                                    var v125 int64 = v41
                                    if len(v3) > 0 {
                                        v3[0] = (v54 & v54)
                                    }
                                    _ = v125
                                }
                                var v126 int8 = 42
                                var v127 float64 = ((3.14159 / 1e10) - 3.14159)
                                switch (fi118 - 0) {
                                case 2147483647:
                                    var v128 uint8 = 216
                                    _ = v128
                                case 4294967295:
                                    v3 = append(v3, v54)
                                default:
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v129 int32 = 1
                                        _ = v129
                                    case "":
                                        var v130 uint32 = 65535
                                        _ = v130
                                    default:
                                        v42 = (v42 + int8(v2))
                                    }
                                }
                                _ = v124
                                _ = v126
                                _ = v127
                            }
                            var v131 int16 = v123
                            _ = v123
                            _ = v131
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v1 = v1
                        }
                        var v132 int = -881
                        _ = v121
                        _ = v132
                    }
                }
                _ = fi118
                _ = v119
            }
        default:
            if (fi65 >= (^(fi65))) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v133 int64 = v41
                _ = v133
            }
        }
        _ = fi65
    }
    if len(v3) > 0 {
        v3[0] = (^(v54))
    }
    v3 = append(v3, 2147483648)
    for fi134 := 0; fi134 < 8; fi134++ {
        v4 -= (1.5 + v4)
        if (!false) {
            var v135 int16 = 0
            _ = v135
        } else {
            if false {
                v136 := GenS16{F0: v2, F1: (!(false || ((^(v2)) == v2))), F2: 2}
                var v137 uint8 = 42
                _ = v136
                _ = v137
            }
            v1 = v41
        }
        var v138 bool = ((("hello" + "") + "\n\t") > "\n\t")
        _ = fi134
        _ = v138
    }
    if (false && false) {
        if (!(0 == v42)) {
            v1 = (v1 - v41)
            var v139 float32 = v4
            _ = v139
        } else {
            for fi140 := 0; fi140 < 7; fi140++ {
                var v141 bool = (!true)
                fmt.Println(fi140)
                _ = fi140
                _ = v141
            }
            switch ("hello" + "hello") {
            case "a":
                switch 127 {
                case 0:
                    fmt.Println(v2, v42)
                case 42:
                    var v142 int8 = (v42 & 0)
                    _ = v142
                case -2147483648:
                    _ = len(v3)
                default:
                    var v143 float64 = 0.1
                    _ = v143
                }
            case "":
                if true {
                    v54 = v54
                    if false {
                        var v144 uint = 2
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v144
                    }
                }
            default:
                var v145 int8 = v42
                _ = v145
            }
            switch "\n\t" {
            case "hello":
                var v146 uint8 = 15
                _ = v146
            case "":
                if ((^(v42)) == (v42 % (((v42 + 82)) | 1))) {
                    v41 = (v1 << ((2) & 63))
                    fmt.Println(v41)
                }
            default:
                switch 42 {
                case -1:
                    var v147 float64 = 1e10
                    _ = v147
                case -2147483648:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 955:
                    var v148 float32 = float32((3.14159 * 1.5))
                    _ = v148
                default:
                    switch "\n\t" {
                    case "":
                        switch -743 {
                        case 32767:
                            v54 -= (^(v54))
                        case 4294967295:
                            var v149 uint32 = (v54 * 587)
                            _ = v149
                        default:
                            for fi150 := 0; fi150 < 2; fi150++ {
                                var v151 uint8 = 186
                                _ = fi150
                                _ = v151
                            }
                        }
                    case "\n\t":
                        for fi152 := 0; fi152 < 14; fi152++ {
                            if true {
                                var v153 int16 = 32767
                                switch fi152 {
                                case 32768:
                                    v4 = v4
                                default:
                                    for fi154 := 0; fi154 < 7; fi154++ {
                                        if len(v3) > 0 {
                                            v3[0] = 42
                                        }
                                        v41 = v41
                                        if true {
                                            var v155 bool = (false && true)
                                            var v156 float64 = -2.25
                                            _ = v155
                                            _ = v156
                                        } else {
                                            v157 := GenS17{F0: -149}
                                            var v158 uint8 = 42
                                            var v159 int8 = (^(v42))
                                            _ = v157
                                            _ = v158
                                            _ = v159
                                        }
                                        _ = fi154
                                    }
                                }
                                if true {
                                    var v160 uint32 = v54
                                    _ = v160
                                }
                                _ = v153
                            }
                            v161 := GenS18{F0: (v41 * (v1 - (v41 & (v41 - 1)))), F1: 42, F2: 1}
                            _ = fi152
                            _ = v161
                        }
                    case "foo bar":
                        var v162 float64 = 1.5
                        _ = v162
                    }
                }
            }
        }
        var v163 uint64 = (v2 ^ 18446744073709551615)
        var v164 int64 = v41
        _ = v163
        _ = v164
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v165 := GenS16{F0: 325, F1: (false || (false && ("a" >= "\n\t"))), F2: 213}
        _ = v165
    }
    v166 := GenS8{F0: 42, F1: -2147483648}
    if (-76 != -269) {
        v167 := GenS19{F0: (v41 ^ v1), F1: v41}
        _ = v167
    } else {
        var v168 int32 = 42
        v169 := GenS20{F0: (v54 & v54), F1: (v168 | int32(v2))}
        _ = v168
        _ = v169
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v166
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v54
    _ = v64
}
