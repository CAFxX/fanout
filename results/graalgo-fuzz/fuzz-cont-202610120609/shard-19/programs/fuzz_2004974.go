package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int; F1 bool; F2 uint64 }
type GenS5 struct { F0 int; F1 string; F2 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 int8; F2 int }
type GenS9 struct { F0 uint; F1 uint16; F2 uint }
type GenS10 struct { F0 int32; F1 int8 }
type GenS11 struct { F0 int; F1 uint32; F2 uint32 }
type GenS12 struct { F0 uint; F1 int64; F2 int }
type GenS13 struct { F0 string; F1 int }
type GenS14 struct { F0 int; F1 string; F2 uint }
type GenS15 struct { F0 uint32; F1 int8 }
type GenS16 struct { F0 uint32; F1 uint8 }
type GenS17 struct { F0 bool; F1 bool }
type GenS18 struct { F0 int32 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 int32; F1 string }
type GenS21 struct { F0 string }
func genFn0() uint {
    if (!false) {
        v5 := GenS1{F0: 0, F1: 0}
        var v6 bool = true
        _ = v5
        _ = v6
    }
    v7 := []int{-2147483648}
    _ = v7
    return 4294967295
}
func genFn1(p0 uint16) uint32 {
    switch 2147483647 {
    case 290:
        p0 += p0
    }
    switch "hello" {
    case "\n\t":
        v8 := []int{4294967295, 9223372036854775807, 680}
        _ = v8
    }
    fmt.Println(p0)
    var v9 uint = 603
    _ = p0
    _ = v9
    return 375
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 144
    var v3 int8 = -118
    var v4 uint8 = (v2 * v2)
    var v10 int64 = 0
    v4 += (^(v2))
    var v11 uint32 = 65535
    switch "a" {
    case "a":
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 bool = true
            v4 = (v4 | 255)
            fmt.Println(v10, v10, v13)
            _ = fi12
            _ = v13
        }
    case "\n\t":
        v14 := GenS1{F0: 0, F1: (v10 + 378)}
        _ = v14
    case "hello":
        if true {
            if false {
                v15 := []int{128}
                _ = v15
            }
        } else {
            switch 2147483647 {
            case -2147483648:
                v16 := GenS2{F0: (v4 / ((1) | 1)), F1: (v11 ^ 65535)}
                _ = v16
            case -9223372036854775808:
                v17 := []int{753, 190}
                _ = v17
            case 525:
                v18 := GenS3{F0: 2, F1: ("" + "\n\t")}
                _ = v18
            default:
                for fi19 := 0; fi19 < 13; fi19++ {
                    v20 := GenS1{F0: 42, F1: 0}
                    _ = fi19
                    _ = v20
                }
            }
            switch (((("hello" + "hello") + ("foo bar" + "\n\t")) + "foo bar") + ((("hello" + "a") + (("a" + "a") + ("hello" + ("hello" + "hello")))) + "")) {
            case "":
                v10 += -713
            case "a":
                fmt.Println(v4, v3, v10)
            case "foo bar":
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 uint16 = 12
                    _ = fi21
                    _ = v22
                }
            default:
                _ = len(v1)
            }
        }
    default:
        fmt.Println(v3)
    }
    v10 += -9223372036854775808
    v23 := GenS3{F0: 887, F1: "\n\t"}
    v24 := GenS4{F0: int(v10), F1: ((v11 % (((v11 / ((v11) | 1))) | 1)) == (v11 << ((255) & 63))), F2: 42}
    if false {
        switch -231 {
        case 80:
            v2 -= v4
        case 981:
            if false {
                v2 += v4
                v25 := GenS5{F0: 2147483647, F1: ("foo bar" + "\n\t"), F2: 42}
                _ = v25
            } else {
                v26 := GenS2{F0: (v4 | (^(v4))), F1: (v11 + (v11 | (v11 - (^(v11)))))}
                fmt.Println(v3, v4, v2)
                _ = v26
            }
        case -9223372036854775808:
            v27 := GenS4{F0: 0, F1: false, F2: 252}
            _ = v27
        }
        if (255 < 1) {
            switch 1 {
            case 128:
                var v28 int64 = int64(v3)
                _ = v28
            case 4294967295:
                v11 = (^(v11))
            }
            if (2147483647 <= 0) {
                if false {
                    var v29 int32 = 1
                    v29 += (v29 << ((uint(v11)) & 63))
                    _ = v29
                } else {
                    var v30 bool = (255 != 42)
                    switch "foo bar" {
                    case "foo bar":
                        switch ("foo bar" + "") {
                        case "foo bar":
                            if (42 != 18446744073709551615) {
                                for fi31 := 0; fi31 < 8; fi31++ {
                                    v32 := v1[("a" + "")]
                                    _ = fi31
                                    _ = v32
                                }
                            }
                        }
                    case "a":
                        var v33 uint16 = 65535
                        _ = v33
                    default:
                        if v30 {
                            delete(v1, "")
                            v30 = ((v11 > (v11 & v11)) || (v11 > v11))
                            delete(v1, "")
                        }
                    }
                    v4 -= (v2 ^ uint8(v11))
                    _ = v30
                }
            }
            v34 := GenS4{F0: -9223372036854775808, F1: (false || true), F2: 0}
            _ = v34
        }
        switch -621 {
        case 2147483647:
            v35 := GenS6{F0: 9223372036854775807}
            _ = v35
        case 128:
            v1[(("hello" + "\n\t") + "foo bar")] = ((v11 << ((18446744073709551615) & 63)) != (v11 + v11))
        default:
            var v36 int64 = (v10 << ((967) & 63))
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 12; fi37++ {
        var v38 uint64 = 1
        v39 := GenS5{F0: (fi37 << ((481) & 63)), F1: ("hello" + "foo bar"), F2: -752}
        var v40 float32 = float32((((0.0 + 1e10) * 1.5) + ((((float64(0.1) - (0.0 * 1e10)) - (1e10 * float64(1e10))) / 1.0) - 0.1)))
        _ = fi37
        _ = v38
        _ = v39
        _ = v40
    }
    var v41 int32 = 2147483647
    v4 += (v2 << ((42) & 63))
    if false {
        v42 := GenS4{F0: 32768, F1: ((v3 >> ((genFn0()) & 63)) > v3), F2: 2}
        fmt.Println(v41)
        _ = v42
    }
    v43 := []int{127}
    for fi44 := 0; fi44 < 7; fi44++ {
        for fi45 := 0; fi45 < 17; fi45++ {
            fmt.Println(v3, v10)
            _ = fi45
        }
        switch "\n\t" {
        case "a":
            v46 := GenS7{F0: 569}
            _ = v46
        case "hello":
            v47 := GenS6{F0: (fi44 % (((^(fi44))) | 1))}
            _ = v47
        default:
            v41 = v41
        }
        _ = fi44
    }
    if false {
        var v48 bool = (("foo bar" + "a") != "a")
        if (v48 && true) {
            if (!(v48 && (v48 || ((0.1 / float32(0.1)) <= 1.5)))) {
                v1["a"] = v48
                switch int(v3) {
                case 42:
                    v49 := GenS4{F0: -132, F1: v48, F2: 18446744073709551615}
                    _ = v49
                }
                v10 -= -9223372036854775808
            } else {
                if true {
                    if true {
                        var v50 bool = (-945 > -32768)
                        if v50 {
                            if (!v48) {
                                v51 := GenS8{F0: 42, F1: 42, F2: 2147483647}
                                fmt.Println(v11, v41, v50)
                                _ = v51
                            } else {
                                v52 := GenS6{F0: -9223372036854775808}
                                v53 := GenS9{F0: 255, F1: 1, F2: 4294967295}
                                switch 32767 {
                                case 32767:
                                    for fi54 := 0; fi54 < 12; fi54++ {
                                        for fi55 := 0; fi55 < 4; fi55++ {
                                            if v50 {
                                                v56 := GenS10{F0: (v41 % ((v41) | 1)), F1: (v3 + v3)}
                                                v48 = v48
                                                _ = v56
                                            } else {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v50 = (!((255 == 1) || v48))
                                                default:
                                                    var v57 int64 = (^(v10))
                                                    _ = v57
                                                }
                                            }
                                            _ = fi55
                                        }
                                        v58 := GenS11{F0: (fi54 | (fi54 >> ((uint(v11)) & 63))), F1: (v11 >> ((255) & 63)), F2: v11}
                                        _ = fi54
                                        _ = v58
                                    }
                                case 831:
                                    v3 = 0
                                case 4294967295:
                                    for fi59 := 0; fi59 < 1; fi59++ {
                                        _ = len(v43)
                                        _ = fi59
                                    }
                                default:
                                    v60 := GenS6{F0: -986}
                                    _ = v60
                                }
                                _ = v52
                                _ = v53
                            }
                            v10 = (v10 >> ((559) & 63))
                        } else {
                            var v61 uint = 18446744073709551615
                            _ = v61
                        }
                        _ = v50
                    }
                    for fi62 := 0; fi62 < 2; fi62++ {
                        switch "hello" {
                        case "a":
                            if (v10 >= (-(v10))) {
                                switch fi62 {
                                case 711:
                                    var v63 int64 = v10
                                    _ = v63
                                default:
                                    v64 := GenS11{F0: (fi62 >> ((1) & 63)), F1: 0, F2: (v11 - (v11 << ((genFn0()) & 63)))}
                                    _ = v64
                                }
                                v43 = append(v43, (fi62 << ((1) & 63)))
                                switch fi62 {
                                case 347:
                                    if (v48 || ((genFn0() == 0) && v48)) {
                                        switch (-(fi62)) {
                                        case -9223372036854775808:
                                            if v48 {
                                                switch (-(fi62)) {
                                                case -1:
                                                    for fi65 := 0; fi65 < 17; fi65++ {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v66 := GenS12{F0: 255, F1: v10, F2: (fi62 + (fi62 % (((fi62 / ((fi62) | 1))) | 1)))}
                                                            _ = v66
                                                        case "a":
                                                            _ = len(v43)
                                                        }
                                                        _ = fi65
                                                    }
                                                case 2147483647:
                                                    v67 := GenS2{F0: v4, F1: (v11 | v11)}
                                                    _ = v67
                                                case 9223372036854775807:
                                                    var v68 int = (fi62 + (fi62 - (-(fi62))))
                                                    _ = v68
                                                }
                                                _ = len(v1)
                                                _ = len(v43)
                                            } else {
                                                var v69 uint = 65535
                                                _ = v69
                                            }
                                        case 32767:
                                            switch 719 {
                                            case -2147483648:
                                                if false {
                                                    var v70 uint8 = (^(v2))
                                                    v70 = 23
                                                    _ = v70
                                                }
                                            case -326:
                                                delete(v1, ("hello" + (("a" + ("hello" + ((("foo bar" + ("a" + (("foo bar" + ("" + "foo bar")) + "a"))) + "") + ""))) + "")))
                                            }
                                        default:
                                            for fi71 := 0; fi71 < 20; fi71++ {
                                                v48 = (42 <= v41)
                                                _ = fi71
                                            }
                                        }
                                        switch (fi62 ^ (fi62 >> ((4294967295) & 63))) {
                                        case 763:
                                            v72 := GenS6{F0: fi62}
                                            _ = v72
                                        case 2147483647:
                                            var v73 int16 = 0
                                            _ = v73
                                        case 32767:
                                            v74 := GenS6{F0: 997}
                                            _ = v74
                                        default:
                                            var v75 bool = true
                                            _ = v75
                                        }
                                        var v76 uint = 0
                                        _ = v76
                                    }
                                case -15:
                                    if len(v43) > 0 {
                                        v43[0] = fi62
                                    }
                                case 328:
                                    if (!v48) {
                                        for fi77 := 0; fi77 < 5; fi77++ {
                                            if len(v43) >= 2 {
                                                v43 = v43[1:len(v43)]
                                            }
                                            if (!(!(!(43 != (^(v3)))))) {
                                                v2 = v4
                                            }
                                            _ = fi77
                                        }
                                    }
                                default:
                                    fmt.Println(v11)
                                }
                            } else {
                                var v78 int16 = -290
                                fmt.Println(v3)
                                switch 0 {
                                case -9223372036854775808:
                                    if ("foo bar" != ("" + ("\n\t" + ("foo bar" + "a")))) {
                                        _ = len(v43)
                                        fmt.Println(v4, v4, v10)
                                        switch "" {
                                        case "foo bar":
                                            v2 = v2
                                        case "":
                                            v79 := v1[("a" + "foo bar")]
                                            _ = v79
                                        case "hello":
                                            if (v10 >= 9223372036854775807) {
                                                for fi80 := 0; fi80 < 7; fi80++ {
                                                    v4 -= 175
                                                    var v81 float32 = 1e10
                                                    if ((fi80 * (fi80 & (fi62 << ((42) & 63)))) == fi80) {
                                                        var v82 int16 = v78
                                                        v83 := GenS2{F0: 0, F1: v11}
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v84 int64 = 1
                                                            _ = v84
                                                        case "a":
                                                            fmt.Println(v48, v2)
                                                            v11 += genFn1(1)
                                                            var v85 int64 = (v10 >> ((2) & 63))
                                                            if (!((!v48) && (!((3.14159 / (((v81 + v81) * v81) - (v81 + v81))) == v81)))) {
                                                                v10 = (v85 & v10)
                                                                switch ("hello" + "\n\t") {
                                                                case "foo bar":
                                                                    v86 := GenS13{F0: ("\n\t" + ""), F1: fi80}
                                                                    _ = v86
                                                                case "hello":
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        var v87 int32 = v41
                                                                        _ = v87
                                                                    case "hello":
                                                                        if v48 {
                                                                            if (true || ((2 != 409) || false)) {
                                                                                if len(v43) >= 2 {
                                                                                    v43 = v43[1:len(v43)]
                                                                                }
                                                                                var v88 uint8 = (v2 + 104)
                                                                                _ = v88
                                                                            } else {
                                                                                switch ("foo bar" + "a") {
                                                                                case "foo bar":
                                                                                    fmt.Println(v11)
                                                                                case "":
                                                                                    var v89 uint32 = 2
                                                                                    _ = v89
                                                                                case "\n\t":
                                                                                    switch ("hello" + "") {
                                                                                    case "a":
                                                                                        v85 = 403
                                                                                    }
                                                                                }
                                                                                v90 := GenS6{F0: fi80}
                                                                                if ((^(v11)) >= 2147483648) {
                                                                                    var v91 string = "hello"
                                                                                    _ = v91
                                                                                } else {
                                                                                    var v92 int8 = v3
                                                                                    v3 += v92
                                                                                    _ = v92
                                                                                }
                                                                                _ = v90
                                                                            }
                                                                        }
                                                                    case "\n\t":
                                                                        v48 = true
                                                                        v41 -= v41
                                                                    default:
                                                                        v11 = 0
                                                                    }
                                                                case "a":
                                                                    _ = len(v43)
                                                                }
                                                                v93 := GenS7{F0: 255}
                                                                _ = v93
                                                            } else {
                                                                v94 := GenS9{F0: 838, F1: 255, F2: 57}
                                                                _ = v94
                                                            }
                                                            _ = v85
                                                        }
                                                        _ = v82
                                                        _ = v83
                                                    } else {
                                                        v4 = (v2 | (v2 >> ((357) & 63)))
                                                    }
                                                    var v95 uint8 = v2
                                                    _ = fi80
                                                    _ = v81
                                                    _ = v95
                                                }
                                            }
                                        default:
                                            for fi96 := 0; fi96 < 19; fi96++ {
                                                var v97 uint = 36
                                                v3 += (v3 ^ int8(fi62))
                                                fmt.Println(v10)
                                                fmt.Println(v4, v10, v78)
                                                v98 := GenS9{F0: (v97 ^ 255), F1: 32767, F2: v97}
                                                v43 = append(v43, 42)
                                                _ = fi96
                                                _ = v97
                                                _ = v98
                                            }
                                        }
                                    }
                                case -2147483648:
                                    v2 -= 127
                                case -402:
                                    v2 -= (v2 % ((v2) | 1))
                                default:
                                    v99 := GenS10{F0: v41, F1: v3}
                                    _ = v99
                                }
                                _ = v78
                            }
                        case "\n\t":
                            if (2 >= 18446744073709551615) {
                                var v100 uint8 = (v2 << ((255) & 63))
                                _ = v100
                            }
                        default:
                            for fi101 := 0; fi101 < 9; fi101++ {
                                v1[("a" + "a")] = (448 >= v11)
                                _ = fi101
                            }
                        }
                        var v102 int16 = -1
                        _ = fi62
                        _ = v102
                    }
                    delete(v1, "")
                } else {
                    var v103 int64 = (v10 % ((v10) | 1))
                    _ = v103
                }
            }
        } else {
            for fi104 := 0; fi104 < 13; fi104++ {
                if v48 {
                    fmt.Println(v41)
                    for fi105 := 0; fi105 < 12; fi105++ {
                        v3 += 1
                        v2 -= v2
                        var v106 uint8 = (^(v2))
                        _ = fi105
                        _ = v106
                    }
                    if true {
                        switch "\n\t" {
                        case "a":
                            v4 = (v4 ^ v2)
                        case "":
                            v107 := GenS1{F0: 65535, F1: -348}
                            _ = v107
                        case "hello":
                            switch fi104 {
                            case 871:
                                var v108 uint64 = 1
                                _ = v108
                            case -121:
                                for fi109 := 0; fi109 < 15; fi109++ {
                                    if len(v43) > 0 {
                                        v43[0] = -558
                                    }
                                    v10 -= -1
                                    var v110 int = fi109
                                    _ = fi109
                                    _ = v110
                                }
                            }
                        default:
                            fmt.Println(fi104, v2)
                        }
                        switch "hello" {
                        case "\n\t":
                            v2 -= (^(v4))
                        }
                        if ((v10 ^ (v10 ^ 1)) == (-(v10))) {
                            if ((v2 / ((v2) | 1)) >= v4) {
                                v111 := GenS8{F0: 32767, F1: 1, F2: fi104}
                                switch (fi104 << ((2) & 63)) {
                                case 383:
                                    for fi112 := 0; fi112 < 19; fi112++ {
                                        fmt.Println(v2)
                                        _ = len(v43)
                                        _ = fi112
                                    }
                                case 2147483647:
                                    if false {
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v113 uint64 = 4294967295
                                            _ = v113
                                        default:
                                            v114 := GenS2{F0: v2, F1: 515}
                                            _ = v114
                                        }
                                        if v48 {
                                            if true {
                                                v115 := v1["hello"]
                                                _ = v115
                                            } else {
                                                if (!(false && v48)) {
                                                    if (!((v11 % ((v11) | 1)) < (v11 + genFn1(2)))) {
                                                        var v116 float32 = (1e10 - ((3.14159 / 1e10) / (3.14159 * -2.25)))
                                                        v3 = 82
                                                        v117 := GenS11{F0: fi104, F1: v11, F2: (v11 % (((v11 << ((65535) & 63))) | 1))}
                                                        _ = v116
                                                        _ = v117
                                                    } else {
                                                        var v118 int64 = (v10 << ((65535) & 63))
                                                        for fi119 := 0; fi119 < 1; fi119++ {
                                                            v1["\n\t"] = (0 <= 65535)
                                                            _ = fi119
                                                        }
                                                        var v120 uint = 255
                                                        _ = v118
                                                        _ = v120
                                                    }
                                                } else {
                                                    fmt.Println(v41, v41)
                                                }
                                            }
                                        } else {
                                            if (true && (true && true)) {
                                                if (71 != 2) {
                                                    var v121 string = ("a" + "hello")
                                                    for fi122 := 0; fi122 < 16; fi122++ {
                                                        if (!v48) {
                                                            v10 = (v10 >> ((genFn0()) & 63))
                                                            v121 = ("" + ("a" + ""))
                                                            _ = len(v43)
                                                        } else {
                                                            var v123 float64 = 0.0
                                                            var v124 uint32 = (v11 >> ((18446744073709551615) & 63))
                                                            switch v121 {
                                                            case "":
                                                                _ = len(v1)
                                                            case "foo bar":
                                                                if (0.0 > (1e10 * 0.1)) {
                                                                    var v125 uint16 = 65535
                                                                    v126 := GenS14{F0: fi122, F1: v121, F2: 2}
                                                                    _ = v125
                                                                    _ = v126
                                                                } else {
                                                                    v127 := GenS15{F0: genFn1(867), F1: (^(v3))}
                                                                    _ = v127
                                                                }
                                                                v2 = v4
                                                                v3 = -52
                                                                v124 = 485
                                                            }
                                                            var v128 uint64 = 2
                                                            _ = v123
                                                            _ = v124
                                                            _ = v128
                                                        }
                                                        if (1e10 != 3.14159) {
                                                            fmt.Println(v48, v11, v2)
                                                            var v129 uint8 = v4
                                                            v2 -= 1
                                                            _ = v129
                                                        } else {
                                                            if (v48 || ((fi122 - (fi104 & (fi104 << ((2) & 63)))) <= (fi104 + (fi104 | fi104)))) {
                                                                fmt.Println(fi104)
                                                            }
                                                            var v130 uint16 = 1
                                                            _ = v130
                                                        }
                                                        _ = len(v43)
                                                        v11 = genFn1(0)
                                                        _ = fi122
                                                    }
                                                    _ = v121
                                                }
                                                switch ("a" + "foo bar") {
                                                case "foo bar":
                                                    var v131 uint16 = 0
                                                    _ = v131
                                                case "a":
                                                    v3 += v3
                                                case "\n\t":
                                                    switch 32768 {
                                                    case 32768:
                                                        if ((!v48) || v48) {
                                                            var v132 uint32 = v11
                                                            v133 := GenS13{F0: (("foo bar" + ("hello" + ("hello" + ("" + "a")))) + "hello"), F1: (fi104 % ((-126) | 1))}
                                                            _ = v132
                                                            _ = v133
                                                        }
                                                    case 1:
                                                        v4 -= v2
                                                    }
                                                default:
                                                    for fi134 := 0; fi134 < 19; fi134++ {
                                                        v10 -= (v10 / (((^(v10))) | 1))
                                                        if v48 {
                                                            var v135 uint8 = (^(v2))
                                                            v2 -= (v135 << ((668) & 63))
                                                            _ = v135
                                                        }
                                                        _ = fi134
                                                    }
                                                }
                                                if (!v48) {
                                                    v136 := GenS4{F0: (fi104 | fi104), F1: (true || v48), F2: 1}
                                                    _ = v136
                                                }
                                            }
                                            v137 := GenS9{F0: genFn0(), F1: 720, F2: 233}
                                            _ = v137
                                        }
                                    }
                                default:
                                    for fi138 := 0; fi138 < 3; fi138++ {
                                        if v48 {
                                            v139 := GenS16{F0: (v11 * v11), F1: v2}
                                            fmt.Println(v10)
                                            _ = v139
                                        } else {
                                            var v140 uint64 = 18446744073709551615
                                            var v141 uint = 903
                                            _ = v140
                                            _ = v141
                                        }
                                        v142 := GenS17{F0: (v48 && v48), F1: (((!((v3 % ((v3) | 1)) <= (v3 - 0))) && ((int8(v11) < 1) && (((!v48) || ((v11 >> ((2) & 63)) > 0)) && (!true)))) || (65535 == 4294967295))}
                                        _ = fi138
                                        _ = v142
                                    }
                                }
                                _ = v111
                            }
                            var v143 uint64 = 1
                            for fi144 := 0; fi144 < 11; fi144++ {
                                v145 := GenS10{F0: 0, F1: (^(v3))}
                                v146 := v1["\n\t"]
                                v3 = v3
                                _ = fi144
                                _ = v145
                                _ = v146
                            }
                            _ = v143
                        } else {
                            v41 -= v41
                            var v147 uint = genFn0()
                            v148 := GenS14{F0: (^(fi104)), F1: (("a" + ("\n\t" + "\n\t")) + "hello"), F2: (v147 | (v147 << ((genFn0()) & 63)))}
                            _ = v147
                            _ = v148
                        }
                    } else {
                        for fi149 := 0; fi149 < 6; fi149++ {
                            v150 := GenS18{F0: v41}
                            if v48 {
                                switch (("hello" + "hello") + "foo bar") {
                                case "":
                                    v2 = 108
                                case "hello":
                                    switch (fi149 ^ (fi149 * fi104)) {
                                    case 757:
                                        if true {
                                            if true {
                                                if len(v43) > 0 {
                                                    v43[0] = -2147483648
                                                }
                                                var v151 int8 = -39
                                                var v152 uint64 = uint64(v11)
                                                _ = v151
                                                _ = v152
                                            }
                                            if v48 {
                                                var v153 float32 = 1.5
                                                fmt.Println(v41, v2, v3)
                                                fmt.Println(v4, v2)
                                                var v154 uint64 = 42
                                                if true {
                                                    v155 := GenS15{F0: (v11 & (v11 - v11)), F1: v3}
                                                    v41 = v41
                                                    switch "" {
                                                    case "a":
                                                        var v156 uint64 = 18446744073709551615
                                                        _ = v156
                                                    case "hello":
                                                        var v157 string = "foo bar"
                                                        _ = v157
                                                    case "\n\t":
                                                        if (!(!((-2.25 / 0.1) == float64((1e10 + v153))))) {
                                                            if (!(v48 || (42 > 1000))) {
                                                                if len(v43) > 0 {
                                                                    v43[0] = (^(fi149))
                                                                }
                                                            }
                                                            v10 = -9223372036854775808
                                                            fmt.Println(v41)
                                                        }
                                                    }
                                                    _ = v155
                                                }
                                                switch ((("\n\t" + ("\n\t" + "hello")) + "a") + "hello") {
                                                case "\n\t":
                                                    fmt.Println(v154, v2)
                                                case "":
                                                    if false {
                                                        var v158 uint = 42
                                                        _ = v158
                                                    }
                                                default:
                                                    v11 += genFn1(65535)
                                                }
                                                _ = v153
                                                _ = v154
                                            }
                                            v11 = v11
                                            if v48 {
                                                v159 := GenS13{F0: "a", F1: fi104}
                                                v1[""] = v48
                                                _ = v159
                                            } else {
                                                if v48 {
                                                    v160 := GenS19{F0: (-1 == 132)}
                                                    switch 8 {
                                                    case 0:
                                                        fmt.Println(v48)
                                                    case 29:
                                                        if (!(-2.25 <= 0.1)) {
                                                            var v161 uint64 = 2
                                                            fmt.Println(v2, v10, v2)
                                                            v41 -= (v41 >> ((51) & 63))
                                                            _ = v161
                                                        } else {
                                                            var v162 uint8 = v4
                                                            _ = v162
                                                        }
                                                    default:
                                                        v1[""] = (!true)
                                                    }
                                                    fmt.Println(v2, fi149)
                                                    v163 := GenS9{F0: 0, F1: 255, F2: 18446744073709551615}
                                                    if (!false) {
                                                        v11 -= v11
                                                        if false {
                                                            var v164 int = (fi149 & fi104)
                                                            fmt.Println(v11, v164, v164)
                                                            var v165 int32 = v41
                                                            _ = v164
                                                            _ = v165
                                                        } else {
                                                            v166 := GenS5{F0: (fi149 ^ fi104), F1: ("\n\t" + "\n\t"), F2: int32(v10)}
                                                            v167 := GenS9{F0: genFn0(), F1: 65535, F2: 18446744073709551615}
                                                            _ = v166
                                                            _ = v167
                                                        }
                                                        v41 -= v41
                                                    } else {
                                                        fmt.Println(v48)
                                                    }
                                                    var v168 float32 = 1e10
                                                    v10 += 42
                                                    _ = v160
                                                    _ = v163
                                                    _ = v168
                                                }
                                                switch "a" {
                                                case "":
                                                    v169 := v1["hello"]
                                                    _ = v169
                                                case "\n\t":
                                                    v170 := GenS7{F0: 1}
                                                    _ = v170
                                                case "hello":
                                                    var v171 int16 = 42
                                                    _ = v171
                                                }
                                                v2 = (v4 << ((65535) & 63))
                                                fmt.Println(v4, v10, fi149)
                                                if ("hello" != "hello") {
                                                    v4 = 0
                                                } else {
                                                    fmt.Println(v3, v4)
                                                    v43 = append(v43, (-(fi149)))
                                                    v172 := GenS20{F0: v41, F1: ("" + "foo bar")}
                                                    v2 -= (^(v4))
                                                    _ = v172
                                                }
                                            }
                                        } else {
                                            if v48 {
                                                fmt.Println(v48)
                                                var v173 int16 = 1
                                                delete(v1, ("foo bar" + "a"))
                                                _ = v173
                                            } else {
                                                _ = len(v43)
                                                var v174 bool = false
                                                if (555 < 2) {
                                                    switch "" {
                                                    case "":
                                                        v175 := GenS21{F0: ("\n\t" + "\n\t")}
                                                        _ = v175
                                                    case "\n\t":
                                                        fmt.Println(fi104, v48)
                                                        v41 += v41
                                                    case "foo bar":
                                                        v176 := GenS14{F0: 42, F1: "", F2: 255}
                                                        _ = v176
                                                    default:
                                                        var v177 uint8 = uint8(v11)
                                                        _ = v177
                                                    }
                                                }
                                                fmt.Println(v174)
                                                var v178 int32 = (v41 * (v41 * -848))
                                                v43 = append(v43, int(v3))
                                                _ = v174
                                                _ = v178
                                            }
                                            v179 := GenS3{F0: 42, F1: (("a" + "a") + ("\n\t" + "hello"))}
                                            var v180 bool = ("" > "\n\t")
                                            _ = v179
                                            _ = v180
                                        }
                                    }
                                }
                            } else {
                                if (18446744073709551615 > 1) {
                                    var v181 int16 = int16(v2)
                                    delete(v1, "a")
                                    _ = v181
                                } else {
                                    _ = len(v43)
                                }
                            }
                            v182 := GenS1{F0: 65535, F1: 9223372036854775807}
                            v11 = 65535
                            var v183 int64 = v10
                            _ = fi149
                            _ = v150
                            _ = v182
                            _ = v183
                        }
                        v43 = append(v43, fi104)
                        _ = len(v1)
                    }
                } else {
                    var v184 uint = 177
                    v185 := v1[((("hello" + ("\n\t" + ("hello" + "\n\t"))) + "\n\t") + "foo bar")]
                    v43 = append(v43, fi104)
                    _ = v184
                    _ = v185
                }
                _ = fi104
            }
            var v186 uint = 4294967295
            fmt.Println(v2, v4)
            _ = v186
        }
        _ = v48
    } else {
        v187 := GenS14{F0: 1, F1: "", F2: 255}
        if true {
            fmt.Println(v11, v11, v2)
            var v188 string = "\n\t"
            _ = v188
        } else {
            switch "foo bar" {
            case "foo bar":
                v4 -= (v4 - (^(v4)))
            }
        }
        _ = v187
    }
    switch ("\n\t" + "foo bar") {
    case "":
        v189 := GenS11{F0: -2147483648, F1: (^(v11)), F2: (v11 & 2147483648)}
        _ = v189
    default:
        var v190 string = "hello"
        _ = v190
    }
    var v191 uint16 = 2
    var v192 float32 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v191:", v191)
    fmt.Println("v192:", v192)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v10
    _ = v11
    _ = v191
    _ = v192
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v41
    _ = v43
}
