package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int; F2 int64 }
type GenS4 struct { F0 bool; F1 uint32; F2 int64 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 string; F1 string; F2 int32 }
type GenS9 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS10 struct { F0 int; F1 bool; F2 bool }
func genFn0(p0 int16, p1 int, p2 int16) int32 {
    fmt.Println(p2)
    p2 -= (p2 / ((-32768) | 1))
    p1 -= 32768
    v7 := []int{1, (p1 - 42), p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 uint64, p1 int64) bool {
    p0 = (^(p0))
    v8 := []int{-30, -921}
    var v9 float64 = 1.5
    v10 := GenS1{F0: 9223372036854775807}
    var v11 uint64 = (^(p0))
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 int16 = 32767
    v2 := []uint32{1, 410}
    v3 := make(map[string]bool)
    v4 := []uint32{2147483647, 2, 2, 4294967295}
    var v5 uint8 = 128
    var v6 int32 = 669
    if (-2147483648 <= 375) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        v5 += (v5 << ((255) & 63))
        v12 := GenS1{F0: 18}
        var v13 int16 = (-(v1))
        _ = v12
        _ = v13
    }
    v6 -= 42
    if (-9223372036854775808 > -2147483648) {
        fmt.Println(v1)
    }
    switch ("\n\t" + "a") {
    case "hello":
        v14 := GenS2{F0: false}
        _ = v14
    case "foo bar":
        v15 := GenS1{F0: 519}
        _ = v15
    case "\n\t":
        fmt.Println(v6)
    default:
        if true {
            if (42 > 4294967295) {
                v16 := GenS2{F0: (true || false)}
                var v17 uint8 = v5
                switch 42 {
                case 32768:
                    fmt.Println(v5, v5, v17)
                default:
                    fmt.Println(v5, v5)
                }
                _ = v16
                _ = v17
            } else {
                var v18 int16 = v1
                v6 -= v6
                switch "" {
                case "hello":
                    if true {
                        if (588 < (^(v6))) {
                            if true {
                                var v19 uint = 909
                                v1 += v1
                                switch (("\n\t" + "\n\t") + ("\n\t" + "\n\t")) {
                                case "\n\t":
                                    for fi20 := 0; fi20 < 3; fi20++ {
                                        for fi21 := 0; fi21 < 13; fi21++ {
                                            var v22 float32 = 0.0
                                            fmt.Println(v19, v6, v19)
                                            delete(v3, "foo bar")
                                            var v23 int8 = 42
                                            _ = len(v3)
                                            var v24 float32 = (((3.14159 * -2.25) + v22) * float32(1e10))
                                            _ = fi21
                                            _ = v22
                                            _ = v23
                                            _ = v24
                                        }
                                        v25 := v3["a"]
                                        _ = fi20
                                        _ = v25
                                    }
                                }
                                _ = v19
                            } else {
                                v26 := GenS1{F0: -1}
                                _ = v26
                            }
                            if ((^(v5)) > v5) {
                                var v27 int32 = v6
                                _ = v27
                            } else {
                                v5 = 127
                                if true {
                                    for fi28 := 0; fi28 < 10; fi28++ {
                                        var v29 uint = 510
                                        _ = fi28
                                        _ = v29
                                    }
                                } else {
                                    _ = len(v4)
                                    if false {
                                        v30 := v3["\n\t"]
                                        fmt.Println(v1, v30)
                                        if (!true) {
                                            if (0.0 >= 3.14159) {
                                                v18 -= int16(v5)
                                            } else {
                                                var v31 int32 = v6
                                                _ = v31
                                            }
                                        }
                                        _ = v30
                                    } else {
                                        if (v18 < 0) {
                                            v1 = v18
                                            if (false || true) {
                                                v5 += v5
                                            } else {
                                                if true {
                                                    v32 := GenS1{F0: -9223372036854775808}
                                                    _ = v32
                                                } else {
                                                    fmt.Println(v1, v18)
                                                }
                                                switch ("hello" + (("hello" + "hello") + "foo bar")) {
                                                case "foo bar":
                                                    for fi33 := 0; fi33 < 13; fi33++ {
                                                        v34 := v3["a"]
                                                        _ = fi33
                                                        _ = v34
                                                    }
                                                default:
                                                    var v35 uint = 0
                                                    _ = v35
                                                }
                                                v5 = (v5 - 163)
                                            }
                                        }
                                    }
                                    for fi36 := 0; fi36 < 19; fi36++ {
                                        for fi37 := 0; fi37 < 19; fi37++ {
                                            v6 = (^(v6))
                                            if false {
                                                fmt.Println(v18, fi37, v1)
                                            }
                                            v38 := GenS3{F0: -2147483648, F1: 946, F2: 0}
                                            _ = fi37
                                            _ = v38
                                        }
                                        _ = fi36
                                    }
                                }
                                for fi39 := 0; fi39 < 12; fi39++ {
                                    for fi40 := 0; fi40 < 19; fi40++ {
                                        v1 -= v18
                                        if false {
                                            var v41 uint64 = 939
                                            v6 = 42
                                            fmt.Println(v41)
                                            _ = v41
                                        }
                                        _ = fi40
                                    }
                                    _ = fi39
                                }
                            }
                        }
                    } else {
                        v3["foo bar"] = ((3.14159 + ((0.1 - (0.0 + (float32(3.14159) * -2.25))) - (1.5 / 3.14159))) != 3.14159)
                        if false {
                            var v42 uint32 = uint32(v5)
                            for fi43 := 0; fi43 < 18; fi43++ {
                                for fi44 := 0; fi44 < 17; fi44++ {
                                    var v45 uint = 461
                                    var v46 float32 = -2.25
                                    _ = fi44
                                    _ = v45
                                    _ = v46
                                }
                                v42 = (v42 / ((v42) | 1))
                                _ = fi43
                            }
                            _ = v42
                        }
                        var v47 float32 = (1.5 - 0.0)
                        _ = v47
                    }
                case "foo bar":
                    v48 := GenS4{F0: (true && (0 == 317)), F1: 238, F2: -1}
                    _ = v48
                case "a":
                    if (4294967295 != 2) {
                        if ((42 != 65535) && (1 == (^(v6)))) {
                            fmt.Println(v6, v18)
                            var v49 uint64 = 1
                            delete(v3, (("\n\t" + ("foo bar" + "a")) + ("a" + "a")))
                            _ = v49
                        } else {
                            for fi50 := 0; fi50 < 17; fi50++ {
                                v51 := GenS4{F0: (!true), F1: 73, F2: 48}
                                _ = fi50
                                _ = v51
                            }
                            for fi52 := 0; fi52 < 8; fi52++ {
                                v18 = (v18 ^ -1)
                                fmt.Println(v6, fi52, v5)
                                _ = fi52
                            }
                            v1 -= 1
                        }
                        for fi53 := 0; fi53 < 5; fi53++ {
                            switch int(v5) {
                            case 32767:
                                v3["hello"] = false
                            case -544:
                                v54 := v3["a"]
                                _ = v54
                            case 966:
                                var v55 int16 = (v18 / ((32767) | 1))
                                _ = v55
                            }
                            v5 = v5
                            _ = fi53
                        }
                    } else {
                        if true {
                            v56 := GenS4{F0: false, F1: 2147483647, F2: 42}
                            v57 := GenS5{F0: 997, F1: 1}
                            _ = v56
                            _ = v57
                        }
                        var v58 string = (("foo bar" + "foo bar") + "a")
                        switch 128 {
                        case 22:
                            v59 := GenS2{F0: true}
                            _ = v59
                        }
                        _ = v58
                    }
                }
                _ = v18
            }
            var v60 int = 22
            v61 := v3["a"]
            _ = v60
            _ = v61
        } else {
            var v62 uint8 = v5
            for fi63 := 0; fi63 < 7; fi63++ {
                var v64 bool = false
                _ = fi63
                _ = v64
            }
            _ = v62
        }
    }
    var v65 int8 = 25
    var v66 uint64 = 0
    v67 := GenS5{F0: 106, F1: -1}
    if (127 != v65) {
        var v68 uint = 0
        fmt.Println(v6, v1, v5)
        for fi69 := 0; fi69 < 2; fi69++ {
            v70 := GenS4{F0: false, F1: 1, F2: -9223372036854775808}
            _ = fi69
            _ = v70
        }
        _ = v68
    } else {
        for fi71 := 0; fi71 < 9; fi71++ {
            for fi72 := 0; fi72 < 15; fi72++ {
                if (true && (!((true || ((!((v6 + v6) <= v6)) && true)) && true))) {
                    switch ("a" + ("hello" + "")) {
                    case "hello":
                        var v73 uint = 670
                        _ = v73
                    case "\n\t":
                        switch (fi72 ^ (fi72 % ((fi71) | 1))) {
                        case 32767:
                            if (!(("" >= "\n\t") && false)) {
                                switch fi71 {
                                case 235:
                                    v74 := GenS6{F0: v6}
                                    _ = v74
                                default:
                                    var v75 string = ""
                                    _ = v75
                                }
                                var v76 uint8 = (^(v5))
                                _ = v76
                            }
                        case -235:
                            v1 += v1
                        case 233:
                            var v77 int16 = v1
                            _ = v77
                        }
                    default:
                        if false {
                            var v78 int8 = (v65 >> ((4294967295) & 63))
                            var v79 uint16 = 42
                            _ = v78
                            _ = v79
                        } else {
                            v3["\n\t"] = ((true && ((v66 ^ (v66 * v66)) <= v66)) || false)
                            switch (fi72 & 948) {
                            case 4294967295:
                                var v80 float64 = float64(0.0)
                                _ = v80
                            default:
                                switch "a" {
                                case "\n\t":
                                    v5 = (v5 ^ 127)
                                case "":
                                    var v81 float32 = -2.25
                                    _ = v81
                                case "a":
                                    switch fi71 {
                                    case 113:
                                        if (18446744073709551615 <= v66) {
                                            var v82 uint16 = 598
                                            switch (fi71 & 1) {
                                            case -9223372036854775808:
                                                switch -1 {
                                                case -1:
                                                    if false {
                                                        if false {
                                                            if false {
                                                                v66 -= v66
                                                                var v83 string = ("\n\t" + "a")
                                                                _ = v83
                                                            }
                                                            if (true && true) {
                                                                v6 = -2147483648
                                                            }
                                                            _ = len(v2)
                                                        }
                                                        var v84 uint32 = 4294967295
                                                        v85 := GenS6{F0: (^(v6))}
                                                        _ = v84
                                                        _ = v85
                                                    } else {
                                                        var v86 uint32 = 754
                                                        switch fi71 {
                                                        case 42:
                                                            var v87 int8 = v65
                                                            _ = v87
                                                        case 0:
                                                            var v88 int64 = -1
                                                            _ = v88
                                                        case 32768:
                                                            v89 := GenS2{F0: false}
                                                            _ = v89
                                                        }
                                                        v90 := GenS2{F0: true}
                                                        _ = v86
                                                        _ = v90
                                                    }
                                                case -349:
                                                    fmt.Println(v82)
                                                case 127:
                                                    if (v66 == 4294967295) {
                                                        v6 = v6
                                                        if true {
                                                            v82 -= (v82 >> ((255) & 63))
                                                            if (-35 < -142) {
                                                                v91 := GenS5{F0: v82, F1: 42}
                                                                v82 = 0
                                                                _ = v91
                                                            } else {
                                                                if false {
                                                                    v92 := GenS7{F0: ((v82 * v82) == 32767)}
                                                                    _ = v92
                                                                }
                                                                if len(v2) > 0 {
                                                                    v2[0] = 2147483647
                                                                }
                                                                var v93 int64 = -66
                                                                _ = v93
                                                            }
                                                        } else {
                                                            switch (^(fi72)) {
                                                            case 128:
                                                                if true {
                                                                    fmt.Println(v6)
                                                                    if ((1000 <= v1) && true) {
                                                                        if (!((-(v6)) != v6)) {
                                                                            var v94 uint8 = v5
                                                                            var v95 int64 = 194
                                                                            v66 += (v66 >> ((255) & 63))
                                                                            _ = v94
                                                                            _ = v95
                                                                        } else {
                                                                            switch fi71 {
                                                                            case 1:
                                                                                fmt.Println(v82, v6, v5)
                                                                            default:
                                                                                switch ("foo bar" + ((((("" + "\n\t") + "foo bar") + ("foo bar" + "hello")) + "a") + "a")) {
                                                                                case "hello":
                                                                                    _ = len(v3)
                                                                                case "foo bar":
                                                                                    var v96 float32 = 3.14159
                                                                                    _ = v96
                                                                                case "\n\t":
                                                                                    v6 += int32(v1)
                                                                                default:
                                                                                    v97 := GenS1{F0: (fi71 / ((-277) | 1))}
                                                                                    _ = v97
                                                                                }
                                                                            }
                                                                            v98 := GenS5{F0: v82, F1: 0}
                                                                            _ = v98
                                                                        }
                                                                    } else {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                        delete(v3, ("foo bar" + "a"))
                                                                        if (true || (0 != 65535)) {
                                                                            if len(v4) > 0 {
                                                                                v4[0] = 2147483647
                                                                            }
                                                                            v82 = 65535
                                                                        } else {
                                                                            v1 -= v1
                                                                            switch 704 {
                                                                            case -9223372036854775808:
                                                                                _ = len(v3)
                                                                            }
                                                                            switch fi71 {
                                                                            case 603:
                                                                                _ = len(v3)
                                                                                if false {
                                                                                    var v99 uint32 = 4294967295
                                                                                    v100 := GenS2{F0: true}
                                                                                    switch "foo bar" {
                                                                                    case "\n\t":
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = v99
                                                                                        }
                                                                                    case "foo bar":
                                                                                        if ((v99 > uint32(v82)) && (0 < 1)) {
                                                                                            switch (fi71 << ((65535) & 63)) {
                                                                                            case 32767:
                                                                                                switch (^(fi71)) {
                                                                                                case 127:
                                                                                                    fmt.Println(v1, v66, v6)
                                                                                                case 32768:
                                                                                                    v82 += v82
                                                                                                case -1:
                                                                                                    if (!true) {
                                                                                                        _ = len(v3)
                                                                                                        fmt.Println(v65, v1)
                                                                                                        v66 -= (^(v66))
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                switch "\n\t" {
                                                                                                case "\n\t":
                                                                                                    var v101 uint64 = (v66 + v66)
                                                                                                    _ = v101
                                                                                                default:
                                                                                                    var v102 uint8 = 255
                                                                                                    _ = v102
                                                                                                }
                                                                                            }
                                                                                            var v103 int16 = (v1 | v1)
                                                                                            _ = v103
                                                                                        }
                                                                                    default:
                                                                                        if false {
                                                                                            fmt.Println(v65)
                                                                                            var v104 uint16 = 32767
                                                                                            v3[(("a" + "foo bar") + "hello")] = false
                                                                                            _ = v104
                                                                                        } else {
                                                                                            var v105 float32 = float32(1e10)
                                                                                            v6 += -1
                                                                                            var v106 uint16 = v82
                                                                                            if (!((-199 == 116) && (fi72 == -382))) {
                                                                                                if ("a" < (("hello" + ("\n\t" + ("a" + ("a" + "\n\t")))) + ("" + "foo bar"))) {
                                                                                                    v6 = (v6 % ((2147483647) | 1))
                                                                                                    var v107 int16 = -32768
                                                                                                    switch "foo bar" {
                                                                                                    case "hello":
                                                                                                        var v108 string = ("\n\t" + "a")
                                                                                                        _ = v108
                                                                                                    default:
                                                                                                        v82 -= uint16(v99)
                                                                                                    }
                                                                                                    _ = v107
                                                                                                } else {
                                                                                                    var v109 uint = 2
                                                                                                    _ = v109
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(v6, v1, v65)
                                                                                                fmt.Println(v66, v5)
                                                                                                v110 := v3["foo bar"]
                                                                                                _ = v110
                                                                                            }
                                                                                            _ = v105
                                                                                            _ = v106
                                                                                        }
                                                                                    }
                                                                                    if (0.0 < 1e10) {
                                                                                        if true {
                                                                                            v1 = 878
                                                                                            if (false && false) {
                                                                                                var v111 int16 = v1
                                                                                                _ = v111
                                                                                            } else {
                                                                                                delete(v3, ("hello" + "foo bar"))
                                                                                                if ((uint32(v82) >= v99) && (((uint8(v99) < v5) && (!((!false) && false))) && (42 <= (v1 + (^(v1)))))) {
                                                                                                    var v112 uint32 = v99
                                                                                                    _ = v112
                                                                                                } else {
                                                                                                    v3[("\n\t" + ("" + "a"))] = true
                                                                                                }
                                                                                                var v113 int64 = -871
                                                                                                var v114 int = 2147483647
                                                                                                _ = v113
                                                                                                _ = v114
                                                                                            }
                                                                                            switch "" {
                                                                                            case "":
                                                                                                var v115 uint8 = v5
                                                                                                _ = v115
                                                                                            }
                                                                                        } else {
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = 42
                                                                                            }
                                                                                            var v116 float32 = ((0.1 + 0.0) * 1e10)
                                                                                            v6 = (v6 + (v6 >> ((4294967295) & 63)))
                                                                                            v6 -= (v6 % ((v6) | 1))
                                                                                            v117 := GenS8{F0: "", F1: "foo bar", F2: v6}
                                                                                            _ = v116
                                                                                            _ = v117
                                                                                        }
                                                                                    }
                                                                                    _ = v100
                                                                                    _ = v99
                                                                                }
                                                                            }
                                                                            _ = len(v4)
                                                                            if (!true) {
                                                                                switch ("\n\t" + (("" + "") + ("\n\t" + ""))) {
                                                                                case "":
                                                                                    var v118 uint32 = 2
                                                                                    _ = v118
                                                                                }
                                                                                fmt.Println(v82, v5, v1)
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            }
                                                                        }
                                                                        v66 -= v66
                                                                        delete(v3, "\n\t")
                                                                        var v119 int8 = v65
                                                                        _ = v119
                                                                    }
                                                                }
                                                            case -259:
                                                                v1 = (v1 % (((v1 / ((v1) | 1))) | 1))
                                                            case 9223372036854775807:
                                                                var v120 int32 = v6
                                                                _ = v120
                                                            }
                                                            switch "hello" {
                                                            case "hello":
                                                                if false {
                                                                    fmt.Println(v6, fi71, v1)
                                                                } else {
                                                                    delete(v3, ("" + "foo bar"))
                                                                }
                                                            default:
                                                                v121 := GenS3{F0: fi71, F1: fi71, F2: int64(v66)}
                                                                v1 -= (v1 * int16(v65))
                                                                v122 := GenS9{F0: (v5 | (^(v5))), F1: v5, F2: 42}
                                                                _ = v121
                                                                _ = v122
                                                            }
                                                            if ((fi71 % ((-9223372036854775808) | 1)) > 1) {
                                                                if (!(!("\n\t" > ((("foo bar" + "\n\t") + "a") + "a")))) {
                                                                    var v123 uint16 = (v82 | v82)
                                                                    v2 = append(v2, 2147483648)
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 1
                                                                    }
                                                                    _ = v123
                                                                }
                                                                if ("hello" > ("\n\t" + ("foo bar" + "foo bar"))) {
                                                                    var v124 uint32 = 2147483648
                                                                    var v125 int64 = -860
                                                                    _ = v124
                                                                    _ = v125
                                                                }
                                                            }
                                                        }
                                                        v4 = append(v4, uint32(v66))
                                                        v82 = 2
                                                    }
                                                }
                                            case 772:
                                                switch (^(fi71)) {
                                                case -1:
                                                    v126 := GenS7{F0: true}
                                                    _ = v126
                                                }
                                            }
                                            v66 -= (v66 * (v66 / (((v66 - v66)) | 1)))
                                            _ = v82
                                        }
                                    }
                                }
                                var v127 int32 = v6
                                if (4294967295 == (v66 + v66)) {
                                    switch -2147483648 {
                                    case 390:
                                        fmt.Println(v1)
                                    case 42:
                                        fmt.Println(v5, v5)
                                    }
                                    fmt.Println(v127)
                                    if false {
                                        if (39 < (v65 ^ v65)) {
                                            if true {
                                                var v128 int16 = 1
                                                var v129 int64 = 42
                                                if true {
                                                    v5 = 0
                                                    var v130 uint16 = 42
                                                    _ = v130
                                                }
                                                if true {
                                                    fmt.Println(fi72)
                                                    switch (fi71 & (^(fi71))) {
                                                    case -770:
                                                        var v131 float64 = 3.14159
                                                        _ = len(v3)
                                                        _ = v131
                                                    case 32768:
                                                        if false {
                                                            v132 := GenS1{F0: fi72}
                                                            var v133 int8 = (v65 / (((v65 % (((v65 / (((v65 * 127)) | 1))) | 1))) | 1))
                                                            var v134 uint64 = v66
                                                            v65 -= v133
                                                            switch fi71 {
                                                            case 2147483647:
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    if (false && (!false)) {
                                                                        v135 := GenS1{F0: fi72}
                                                                        _ = v135
                                                                    } else {
                                                                        if true {
                                                                            var v136 float32 = 1e10
                                                                            var v137 uint8 = (^(v5))
                                                                            _ = v136
                                                                            _ = v137
                                                                        }
                                                                        var v138 int32 = v127
                                                                        if false {
                                                                            var v139 uint16 = 32767
                                                                            _ = v139
                                                                        }
                                                                        _ = v138
                                                                    }
                                                                case "hello":
                                                                    var v140 bool = true
                                                                    _ = v140
                                                                case "a":
                                                                    if true {
                                                                        var v141 uint32 = 51
                                                                        var v142 int8 = 42
                                                                        _ = v141
                                                                        _ = v142
                                                                    }
                                                                }
                                                            }
                                                            switch (fi72 >> ((333) & 63)) {
                                                            case 128:
                                                                v133 = (v133 + (v65 * 127))
                                                            default:
                                                                v65 -= int8(v129)
                                                            }
                                                            var v143 uint64 = 2
                                                            v144 := GenS3{F0: (fi71 & fi72), F1: fi71, F2: (v129 * (v129 / ((v129) | 1)))}
                                                            if (!(true || true)) {
                                                                v1 = v1
                                                                v65 += (v65 - v133)
                                                                switch "" {
                                                                case "":
                                                                    v134 += (v143 << ((65535) & 63))
                                                                }
                                                            } else {
                                                                v145 := GenS2{F0: false}
                                                                var v146 uint32 = 2147483648
                                                                _ = v145
                                                                _ = v146
                                                            }
                                                            _ = v132
                                                            _ = v133
                                                            _ = v134
                                                            _ = v143
                                                            _ = v144
                                                        }
                                                    default:
                                                        if (true && (v127 > genFn0((v128 ^ v1), int(v129), (^(v128))))) {
                                                            v147 := GenS5{F0: 255, F1: -1}
                                                            var v148 int8 = (v65 ^ (-(v65)))
                                                            _ = v147
                                                            _ = v148
                                                        } else {
                                                            var v149 int8 = 42
                                                            fmt.Println(v5)
                                                            if false {
                                                                v5 -= (v5 * (v5 + v5))
                                                            }
                                                            _ = v149
                                                        }
                                                    }
                                                    var v150 uint64 = 4294967295
                                                    _ = v150
                                                } else {
                                                    var v151 uint16 = 42
                                                    _ = v151
                                                }
                                                var v152 int8 = -61
                                                _ = v128
                                                _ = v129
                                                _ = v152
                                            } else {
                                                var v153 uint = 255
                                                fmt.Println(v5)
                                                if true {
                                                    _ = len(v2)
                                                    v5 = uint8(v66)
                                                    var v154 int8 = (v65 << ((v153) & 63))
                                                    _ = v154
                                                } else {
                                                    var v155 int64 = -9223372036854775808
                                                    if len(v4) > 0 {
                                                        v4[0] = 4294967295
                                                    }
                                                    _ = v155
                                                }
                                                v1 = (v1 / (((v1 ^ v1)) | 1))
                                                _ = v153
                                            }
                                            v6 += v6
                                            v6 = (v6 % (((-(v127))) | 1))
                                            var v156 uint16 = 2
                                            _ = v156
                                        } else {
                                            v1 = (-(v1))
                                        }
                                        if (2 < 2147483648) {
                                            var v157 int = fi71
                                            switch "" {
                                            case "hello":
                                                v127 = v6
                                            default:
                                                v5 = 128
                                            }
                                            _ = v157
                                        }
                                        var v158 string = "a"
                                        fmt.Println(v127, v158, v66)
                                        _ = v158
                                    } else {
                                        _ = len(v3)
                                        v66 = v66
                                        var v159 float32 = float32((1e10 / -2.25))
                                        if true {
                                            var v160 uint64 = 42
                                            _ = v160
                                        }
                                        _ = v159
                                    }
                                }
                                _ = v127
                            }
                        }
                    }
                }
                v161 := GenS10{F0: -986, F1: ((^(v66)) >= (v66 - 0)), F2: (!(fi71 > fi71))}
                _ = fi72
                _ = v161
            }
            _ = fi71
        }
        var v162 int32 = (v6 + genFn0(0, 42, 899))
        v1 -= 546
        _ = v162
    }
    for fi163 := 0; fi163 < 14; fi163++ {
        fmt.Println(v66, v5)
        _ = fi163
    }
    v164 := GenS2{F0: (float64(1.5) == 0.0)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v164
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v67
}
