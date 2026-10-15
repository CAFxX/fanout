package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32; F1 int16 }
type GenS9 struct { F0 uint64; F1 uint64; F2 uint }
type GenS10 struct { F0 uint16; F1 uint64 }
type GenS11 struct { F0 int; F1 uint; F2 uint64 }
type GenS12 struct { F0 string; F1 uint }
type GenS13 struct { F0 int; F1 uint16; F2 int16 }
type GenS14 struct { F0 uint64; F1 int8; F2 int64 }
type GenS15 struct { F0 uint64; F1 int32 }
type GenS16 struct { F0 uint; F1 uint; F2 uint32 }
type GenS17 struct { F0 int64; F1 uint }
type GenS18 struct { F0 bool }
func genFn0() uint8 {
    var v5 int32 = 0
    for fi6 := 0; fi6 < 4; fi6++ {
        v5 += (v5 | v5)
        _ = fi6
    }
    v7 := GenS1{F0: (2147483647 > 122), F1: -163}
    var v8 int = -1
    _ = v5
    _ = v7
    _ = v8
    return 155
}
func genFn1(p0 bool) int64 {
    v9 := GenS2{F0: -670, F1: 2}
    if false {
        p0 = (9223372036854775807 <= -9223372036854775808)
        switch 713 {
        case 4294967295:
            for fi10 := 0; fi10 < 18; fi10++ {
                switch ("\n\t" + "foo bar") {
                case "":
                    var v11 int64 = -9223372036854775808
                    _ = v11
                case "hello":
                    var v12 int32 = -1
                    _ = v12
                }
                _ = fi10
            }
        case -403:
            p0 = p0
        case 1:
            v13 := []int{1, -9223372036854775808, 842, -1}
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 1; fi14++ {
            for fi15 := 0; fi15 < 6; fi15++ {
                fmt.Println(p0, fi14, fi15)
                fmt.Println(fi14, fi14, fi15)
                _ = fi15
            }
            _ = fi14
        }
    }
    v16 := []int{-1, -736, 440, -9223372036854775808}
    p0 = (1 != 0)
    if p0 {
        fmt.Println(p0)
    } else {
        p0 = p0
    }
    _ = p0
    _ = v16
    _ = v9
    return 9223372036854775807
}

func main() {
    v1 := []string{("hello" + "a"), "hello", ("a" + "foo bar")}
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    v4 := []uint32{65535, 0, 597}
    switch ("a" + (("a" + ("a" + "foo bar")) + ("foo bar" + "hello"))) {
    case "\n\t":
        var v17 float64 = 1e10
        _ = v17
    case "foo bar":
        if true {
            var v18 int64 = genFn1((2 >= 607))
            switch "" {
            case "":
                v4 = append(v4, 2147483647)
            case "\n\t":
                v18 -= (v18 - v18)
            default:
                for fi19 := 0; fi19 < 4; fi19++ {
                    var v20 bool = true
                    for fi21 := 0; fi21 < 15; fi21++ {
                        v22 := GenS2{F0: 32767, F1: 246}
                        v20 = (1 <= -601)
                        v20 = v20
                        _ = fi21
                        _ = v22
                    }
                    _ = fi19
                    _ = v20
                }
            }
            _ = v18
        } else {
            var v23 int16 = 98
            _ = v23
        }
    case "":
        v24 := GenS3{F0: 9223372036854775807, F1: 1}
        _ = v24
    }
    _ = len(v4)
    if false {
        v25 := GenS2{F0: 42, F1: 910}
        var v26 float64 = (float64(3.14159) - 0.0)
        var v27 int8 = 127
        _ = v25
        _ = v26
        _ = v27
    }
    var v28 uint32 = 1
    for fi29 := 0; fi29 < 5; fi29++ {
        v28 -= 2147483648
        switch (fi29 * (fi29 & fi29)) {
        case 32767:
            fmt.Println(fi29, v28)
        case -267:
            if true {
                if ((!false) || false) {
                    if true {
                        for fi30 := 0; fi30 < 6; fi30++ {
                            var v31 int = fi29
                            if (-758 <= -680) {
                                if true {
                                    switch "foo bar" {
                                    case "a":
                                        v31 = fi30
                                        if false {
                                            var v32 int32 = 552
                                            _ = v32
                                        } else {
                                            var v33 float64 = (((3.14159 * 0.0) * float64(1e10)) * 1e10)
                                            _ = v33
                                        }
                                        fmt.Println(fi29, fi29, v28)
                                    }
                                }
                            }
                            _ = fi30
                            _ = v31
                        }
                        for fi34 := 0; fi34 < 10; fi34++ {
                            fmt.Println(fi29)
                            if (510 > 565) {
                                var v35 uint64 = 18446744073709551615
                                if true {
                                    v28 -= (v28 << ((65535) & 63))
                                    fmt.Println(fi29, fi34, fi34)
                                    v35 = 859
                                    var v36 int8 = -1
                                    if false {
                                        v2[(("\n\t" + "hello") + "hello")] = (v28 & (v28 ^ (v28 << ((4294967295) & 63))))
                                        switch ("a" + "foo bar") {
                                        case "":
                                            v37 := GenS3{F0: 17, F1: v35}
                                            _ = v37
                                        default:
                                            fmt.Println(v36, v36)
                                        }
                                    } else {
                                        v38 := GenS4{F0: 235, F1: (v28 & (v28 - v28)), F2: (v36 % (((v36 ^ -1)) | 1))}
                                        fmt.Println(v28, fi34, v35)
                                        _ = v38
                                    }
                                    _ = v36
                                }
                                v35 = v35
                                _ = v35
                            }
                            _ = fi34
                        }
                    } else {
                        if (!true) {
                            v39 := GenS5{F0: 1, F1: 1}
                            _ = v39
                        } else {
                            v28 += (^(v28))
                            var v40 float32 = float32(0.0)
                            v41 := GenS6{F0: (v28 * (v28 | 437)), F1: 436}
                            _ = v40
                            _ = v41
                        }
                        switch "\n\t" {
                        case "foo bar":
                            fmt.Println(v28, fi29)
                        case "a":
                            v42 := GenS2{F0: 347, F1: 879}
                            _ = v42
                        default:
                            var v43 bool = true
                            _ = v43
                        }
                    }
                    if (!true) {
                        if (2147483647 == -2147483648) {
                            v44 := GenS5{F0: 4294967295, F1: -476}
                            v28 = (v28 & v28)
                            _ = v44
                        } else {
                            delete(v3, 32768)
                            v28 = v28
                            for fi45 := 0; fi45 < 19; fi45++ {
                                switch (fi45 % (((fi29 - (fi29 | 1))) | 1)) {
                                case 710:
                                    switch "a" {
                                    case "\n\t":
                                        v46 := v2[("\n\t" + "a")]
                                        if true {
                                            v47 := GenS7{F0: (("\n\t" + "\n\t") + "")}
                                            var v48 bool = ("" <= "\n\t")
                                            _ = v47
                                            _ = v48
                                        }
                                        v28 += 0
                                        _ = v46
                                    case "a":
                                        var v49 bool = (!false)
                                        _ = v49
                                    default:
                                        if ((!(!false)) || (v28 >= 0)) {
                                            v50 := v2["\n\t"]
                                            fmt.Println(fi29, v50)
                                            if (1 <= 189) {
                                                fmt.Println(v28, fi45, v28)
                                            } else {
                                                var v51 string = ""
                                                _ = v51
                                            }
                                            _ = v50
                                        }
                                    }
                                case 1:
                                    v28 += (v28 / ((v28) | 1))
                                }
                                if true {
                                    switch (fi29 << ((2) & 63)) {
                                    case 42:
                                        var v52 bool = false
                                        _ = v52
                                    default:
                                        if (((0 != v28) || true) || (!true)) {
                                            switch 908 {
                                            case -1:
                                                _ = len(v2)
                                            case 42:
                                                var v53 uint8 = 1
                                                _ = v53
                                            case 9223372036854775807:
                                                delete(v2, "")
                                            default:
                                                v54 := GenS8{F0: (v28 >> ((18446744073709551615) & 63)), F1: -32768}
                                                _ = v54
                                            }
                                            switch "foo bar" {
                                            case "a":
                                                fmt.Println(v28)
                                            default:
                                                v28 = v28
                                            }
                                            var v55 int64 = 9223372036854775807
                                            _ = v55
                                        } else {
                                            v28 += 2
                                            if true {
                                                var v56 int = 2147483647
                                                _ = v56
                                            }
                                        }
                                    }
                                    var v57 string = "\n\t"
                                    v4 = append(v4, (v28 * (v28 - (v28 % ((v28) | 1)))))
                                    _ = v57
                                } else {
                                    switch ("foo bar" + (("hello" + "\n\t") + "")) {
                                    case "":
                                        var v58 uint = 2
                                        _ = v58
                                    case "a":
                                        v59 := GenS2{F0: -1, F1: 42}
                                        _ = v59
                                    }
                                    if (false || (!(((("a" + "foo bar") + "a") < ("\n\t" + "a")) && false))) {
                                        v60 := GenS5{F0: 18446744073709551615, F1: -9223372036854775808}
                                        v28 = v28
                                        v28 += (v28 - v28)
                                        _ = v60
                                    } else {
                                        var v61 int64 = 9223372036854775807
                                        _ = v61
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        switch ("hello" + (("\n\t" + "foo bar") + "hello")) {
                                        case "hello":
                                            v62 := GenS8{F0: v28, F1: 32767}
                                            _ = v62
                                        case "":
                                            if ((1e10 <= (float64(float32((0.1 * (3.14159 + (0.1 * 3.14159))))) + (1e10 - (1e10 / -2.25)))) && true) {
                                                v28 = v28
                                                _ = len(v1)
                                                var v63 uint16 = 0
                                                _ = v63
                                            } else {
                                                var v64 int32 = -2147483648
                                                _ = v64
                                            }
                                        }
                                    case "hello":
                                        if true {
                                            switch (fi29 * fi45) {
                                            case -1:
                                                var v65 uint8 = 1
                                                _ = v65
                                            case 4294967295:
                                                var v66 uint = 65535
                                                _ = v66
                                            }
                                        } else {
                                            v28 = (^(v28))
                                            v28 += v28
                                            switch ("foo bar" + "hello") {
                                            case "hello":
                                                fmt.Println(fi29, fi45, fi45)
                                            default:
                                                var v67 int8 = -100
                                                _ = v67
                                            }
                                            var v68 uint32 = v28
                                            _ = v68
                                        }
                                    case "\n\t":
                                        v69 := v2["a"]
                                        _ = v69
                                    }
                                }
                                v70 := GenS9{F0: 4294967295, F1: 875, F2: 4294967295}
                                _ = fi45
                                _ = v70
                            }
                        }
                        v28 += 758
                    } else {
                        var v71 uint64 = 449
                        for fi72 := 0; fi72 < 18; fi72++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            fmt.Println(v71, v71)
                            var v73 uint32 = v28
                            _ = fi72
                            _ = v73
                        }
                        _ = v71
                    }
                    v28 = (v28 - 205)
                } else {
                    fmt.Println(v28)
                    if (1e10 != -2.25) {
                        var v74 int64 = 0
                        _ = v74
                    } else {
                        switch ("a" + ("" + "a")) {
                        case "\n\t":
                            fmt.Println(v28, v28)
                        }
                        var v75 int16 = 32767
                        _ = v75
                    }
                }
                for fi76 := 0; fi76 < 14; fi76++ {
                    var v77 int = (^(fi76))
                    _ = fi76
                    _ = v77
                }
                var v78 uint16 = 549
                _ = v78
            } else {
                v28 = (v28 & v28)
            }
        case 127:
            v28 = (v28 << ((4294967295) & 63))
        }
        v79 := GenS3{F0: 9223372036854775807, F1: 2}
        _ = fi29
        _ = v79
    }
    switch 127 {
    case -989:
        var v80 int32 = 2147483647
        _ = v80
    case -1:
        if false {
            var v81 int64 = -9223372036854775808
            fmt.Println(v81, v28)
            _ = v81
        } else {
            if (!((0.1 < -2.25) || false)) {
                v28 -= 4294967295
                v28 = 0
            } else {
                if false {
                    v82 := GenS10{F0: 914, F1: 42}
                    for fi83 := 0; fi83 < 14; fi83++ {
                        v28 -= v28
                        var v84 uint = 42
                        _ = fi83
                        _ = v84
                    }
                    _ = v82
                }
                switch "foo bar" {
                case "hello":
                    v28 = v28
                case "foo bar":
                    v85 := GenS6{F0: v28, F1: 32767}
                    _ = v85
                case "a":
                    if false {
                        fmt.Println(v28)
                        v28 = (v28 | v28)
                    } else {
                        v28 = 660
                    }
                default:
                    v28 = (v28 % ((42) | 1))
                }
            }
            if (((("foo bar" + (("" + "foo bar") + (("a" + ("foo bar" + "hello")) + (("a" + ("hello" + ("hello" + ("\n\t" + ("\n\t" + ""))))) + ("\n\t" + "foo bar"))))) + "a") + ("foo bar" + "")) != "") {
                for fi86 := 0; fi86 < 2; fi86++ {
                    switch ("foo bar" + "foo bar") {
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        fmt.Println(fi86, v28)
                    }
                    if (-802 >= -1) {
                        var v87 string = "foo bar"
                        fmt.Println(v87, fi86, fi86)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v87
                    } else {
                        var v88 int16 = 1
                        _ = v88
                    }
                    _ = fi86
                }
                fmt.Println(v28)
                switch 4294967295 {
                case -367:
                    if true {
                        fmt.Println(v28)
                        v28 = (v28 | (v28 - (^(v28))))
                        var v89 uint32 = 1
                        _ = v89
                    }
                default:
                    v28 -= v28
                }
            } else {
                var v90 int64 = genFn1((true && false))
                switch -1 {
                case 32767:
                    if (732 >= 591) {
                        var v91 uint = 2
                        if false {
                            fmt.Println(v28, v90, v91)
                            if true {
                                for fi92 := 0; fi92 < 11; fi92++ {
                                    if true {
                                        switch fi92 {
                                        case 128:
                                            fmt.Println(fi92, v91, fi92)
                                        case -1:
                                            for fi93 := 0; fi93 < 9; fi93++ {
                                                v94 := GenS11{F0: fi92, F1: v91, F2: 18446744073709551615}
                                                if (genFn1((-375 == 42)) <= 0) {
                                                    _ = len(v3)
                                                    switch "\n\t" {
                                                    case "a":
                                                        v90 += genFn1(((v90 ^ (-(v90))) < (v90 + genFn1(false))))
                                                    case "hello":
                                                        v90 += genFn1(true)
                                                    default:
                                                        if true {
                                                            v95 := GenS10{F0: 728, F1: 1}
                                                            switch ("hello" + "hello") {
                                                            case "foo bar":
                                                                v90 = (v90 | v90)
                                                            case "a":
                                                                var v96 float32 = 1.5
                                                                _ = v96
                                                            }
                                                            _ = v95
                                                        }
                                                    }
                                                } else {
                                                    if false {
                                                        v97 := GenS12{F0: (((("hello" + "hello") + ("\n\t" + (("a" + ("a" + "")) + "\n\t"))) + "\n\t") + "\n\t"), F1: 18446744073709551615}
                                                        v98 := GenS13{F0: (fi92 / ((fi93) | 1)), F1: 2, F2: 32767}
                                                        v91 -= v91
                                                        _ = v97
                                                        _ = v98
                                                    } else {
                                                        fmt.Println(v91)
                                                        v28 = (v28 + 2147483647)
                                                        v2[("\n\t" + ("" + "foo bar"))] = v28
                                                    }
                                                    var v99 int8 = 42
                                                    v100 := GenS14{F0: 402, F1: 86, F2: (v90 & v90)}
                                                    var v101 bool = (true || true)
                                                    _ = v100
                                                    _ = v101
                                                    _ = v99
                                                }
                                                if (!(((v91 >> ((v91) & 63)) > v91) && (0 != 378))) {
                                                    v102 := GenS5{F0: 482, F1: 1}
                                                    v1 = append(v1, "foo bar")
                                                    v103 := GenS5{F0: uint64(v28), F1: v90}
                                                    _ = v102
                                                    _ = v103
                                                } else {
                                                    v91 = (v91 << ((255) & 63))
                                                }
                                                _ = fi93
                                                _ = v94
                                            }
                                        case -683:
                                            for fi104 := 0; fi104 < 10; fi104++ {
                                                var v105 int16 = 32767
                                                if true {
                                                    v106 := GenS12{F0: ("\n\t" + "a"), F1: (v91 / ((v91) | 1))}
                                                    _ = v106
                                                } else {
                                                    switch (fi104 / ((fi92) | 1)) {
                                                    case 105:
                                                        v107 := GenS9{F0: 42, F1: 965, F2: 2}
                                                        _ = v107
                                                    default:
                                                        fmt.Println(fi92, v105)
                                                    }
                                                    var v108 int16 = -1
                                                    _ = v108
                                                }
                                                if true {
                                                    var v109 int8 = -1
                                                    v110 := GenS6{F0: v28, F1: v105}
                                                    v90 = v90
                                                    _ = v109
                                                    _ = v110
                                                } else {
                                                    v111 := GenS15{F0: 2, F1: 0}
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v112 := GenS1{F0: (v91 != v91), F1: v105}
                                                    _ = v111
                                                    _ = v112
                                                }
                                                _ = fi104
                                                _ = v105
                                            }
                                        }
                                        if ("" <= "hello") {
                                            if (v91 > v91) {
                                                for fi113 := 0; fi113 < 12; fi113++ {
                                                    if (!false) {
                                                        v91 += 4294967295
                                                        var v114 int64 = (v90 & genFn1(false))
                                                        v114 += 1
                                                        _ = v114
                                                    }
                                                    var v115 bool = (9223372036854775807 < v90)
                                                    _ = fi113
                                                    _ = v115
                                                }
                                            }
                                            switch int(v91) {
                                            case 4294967295:
                                                v90 = (-(v90))
                                            case 32768:
                                                var v116 int16 = 1
                                                _ = v116
                                            default:
                                                if (688 < 42) {
                                                    v117 := v3[-360]
                                                    switch 32767 {
                                                    case -354:
                                                        delete(v3, 4294967295)
                                                    case 32767:
                                                        v90 = (v90 & genFn1((-1 > -84)))
                                                    case 59:
                                                        var v118 int64 = v90
                                                        _ = v118
                                                    }
                                                    v91 = v91
                                                    _ = v117
                                                } else {
                                                    for fi119 := 0; fi119 < 10; fi119++ {
                                                        v120 := GenS10{F0: 2, F1: 4294967295}
                                                        v28 = (v28 >> ((v91) & 63))
                                                        _ = fi119
                                                        _ = v120
                                                    }
                                                    v91 = v91
                                                }
                                            }
                                            var v121 uint64 = 1
                                            _ = v121
                                        } else {
                                            v1 = append(v1, ("\n\t" + ("a" + "")))
                                            v90 -= v90
                                            fmt.Println(fi92)
                                        }
                                    } else {
                                        v90 += -488
                                        v91 += (^(v91))
                                    }
                                    _ = fi92
                                }
                                v90 = genFn1((uint8(v28) > 42))
                            }
                            fmt.Println(v90, v28, v91)
                        } else {
                            fmt.Println(v91, v91, v90)
                            var v122 string = "hello"
                            _ = v122
                        }
                        v90 += -9223372036854775808
                        _ = v91
                    }
                default:
                    switch -177 {
                    case -171:
                        for fi123 := 0; fi123 < 13; fi123++ {
                            switch "\n\t" {
                            case "":
                                fmt.Println(v90, v90, v28)
                            }
                            _ = fi123
                        }
                    case 4294967295:
                        v28 = 42
                    default:
                        if true {
                            for fi124 := 0; fi124 < 7; fi124++ {
                                v28 += (v28 + v28)
                                v90 += (v90 | 0)
                                var v125 bool = (false && (0 > 70))
                                _ = fi124
                                _ = v125
                            }
                        } else {
                            if (127 != 1) {
                                if true {
                                    v126 := GenS5{F0: 1, F1: genFn1((1 <= -268))}
                                    if true {
                                        switch 32768 {
                                        case 170:
                                            fmt.Println(v90)
                                        }
                                        fmt.Println(v28, v28)
                                        v28 += (v28 + (v28 ^ v28))
                                    } else {
                                        v90 -= (v90 | v90)
                                    }
                                    _ = v126
                                } else {
                                    var v127 string = ("foo bar" + "hello")
                                    switch (v127 + v127) {
                                    case "foo bar":
                                        v3[2147483647] = (^(v28))
                                    }
                                    _ = v127
                                }
                                var v128 float32 = 0.0
                                _ = v128
                            } else {
                                var v129 uint = 365
                                var v130 uint8 = uint8(v129)
                                v130 = v130
                                _ = v129
                                _ = v130
                            }
                        }
                    }
                }
                v90 = v90
                _ = v90
            }
            for fi131 := 0; fi131 < 14; fi131++ {
                if (-1 > 42) {
                    if (false && false) {
                        for fi132 := 0; fi132 < 4; fi132++ {
                            var v133 int = -2147483648
                            v133 = fi131
                            v134 := GenS16{F0: 18446744073709551615, F1: 1, F2: 2147483647}
                            _ = fi132
                            _ = v133
                            _ = v134
                        }
                        v28 += (v28 - v28)
                        switch "foo bar" {
                        case "hello":
                            v135 := v2[("foo bar" + "hello")]
                            _ = v135
                        case "":
                            for fi136 := 0; fi136 < 20; fi136++ {
                                switch (-(fi136)) {
                                case -1:
                                    v1 = append(v1, "a")
                                case 127:
                                    fmt.Println(v28)
                                default:
                                    if true {
                                        v3[fi136] = 8
                                        v28 = (^(v28))
                                        v28 = v28
                                    }
                                }
                                _ = fi136
                            }
                        case "a":
                            v28 = v28
                        }
                    }
                } else {
                    for fi137 := 0; fi137 < 3; fi137++ {
                        fmt.Println(fi137, v28, v28)
                        var v138 int64 = 1
                        if true {
                            v139 := GenS17{F0: (-(v138)), F1: 42}
                            fmt.Println(fi137, fi137)
                            _ = v139
                        } else {
                            var v140 int32 = int32(v138)
                            var v141 int8 = 42
                            var v142 int8 = v141
                            _ = v140
                            _ = v141
                            _ = v142
                        }
                        _ = fi137
                        _ = v138
                    }
                }
                v28 = (v28 / ((v28) | 1))
                _ = fi131
            }
        }
    case 0:
        var v143 uint = 608
        _ = v143
    default:
        v144 := GenS18{F0: true}
        _ = v144
    }
    var v145 int64 = 0
    v146 := GenS12{F0: "hello", F1: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v145:", v145)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v145
    _ = v146
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
