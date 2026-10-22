package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 string }
type GenS2 struct { F0 uint32; F1 string; F2 uint16 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 bool; F1 int16; F2 int16 }
type GenS5 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 int64 }
type GenS8 struct { F0 int32; F1 int32; F2 int32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int32; F1 int32; F2 uint }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int; F1 int }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS15 struct { F0 uint32; F1 uint32 }
type GenS16 struct { F0 uint32; F1 int16 }
type GenS17 struct { F0 bool; F1 bool }
type GenS18 struct { F0 uint32; F1 bool }
type GenS19 struct { F0 string; F1 int32 }
type GenS20 struct { F0 uint16; F1 uint32 }
type GenS21 struct { F0 uint32; F1 bool }
type GenS22 struct { F0 int64 }
type GenS23 struct { F0 int32; F1 uint16; F2 int32 }
type GenS24 struct { F0 int32; F1 int32; F2 int64 }

func main() {
    var v1 int16 = 0
    var v2 int = 198
    v3 := make(map[uint32]string)
    var v4 uint32 = 2
    if (2 > 1) {
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 int = v2
            v6 = (v2 ^ (^(v6)))
            v6 -= fi5
            _ = fi5
            _ = v6
        }
        v7 := GenS1{F0: 255, F1: 0, F2: "hello"}
        _ = v7
    } else {
        delete(v3, 740)
    }
    fmt.Println(v1, v1)
    v8 := GenS1{F0: 2, F1: -1, F2: ("a" + "a")}
    var v9 uint64 = 18446744073709551615
    if (((false || true) && (-90 < 100)) && (v4 == (v4 << ((968) & 63)))) {
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := []int{(v2 - 1)}
            if ((!("foo bar" != "foo bar")) && (v1 <= (v1 - (v1 + (^(v1)))))) {
                v12 := GenS2{F0: v4, F1: (("a" + (("\n\t" + (((("a" + "a") + "foo bar") + (("" + "hello") + "")) + "a")) + "hello")) + "a"), F2: 0}
                if (98 == v4) {
                    v11 = append(v11, fi10)
                    v1 -= v1
                } else {
                    v4 = v4
                    switch (v2 >> ((18446744073709551615) & 63)) {
                    case 2147483647:
                        for fi13 := 0; fi13 < 3; fi13++ {
                            v9 -= v9
                            _ = fi13
                        }
                    case 32767:
                        switch fi10 {
                        case 0:
                            for fi14 := 0; fi14 < 14; fi14++ {
                                v15 := GenS3{F0: "foo bar", F1: "hello"}
                                v4 -= 670
                                switch "" {
                                case "foo bar":
                                    v16 := GenS2{F0: (v4 * (v4 & 65535)), F1: "a", F2: 810}
                                    v17 := GenS4{F0: (-189 <= -1), F1: v1, F2: (v1 / (((v1 >> ((65535) & 63))) | 1))}
                                    _ = v16
                                    _ = v17
                                case "hello":
                                    fmt.Println(v9, v9)
                                case "\n\t":
                                    fmt.Println(fi14, v1, v1)
                                default:
                                    v18 := v3[v4]
                                    _ = v18
                                }
                                _ = fi14
                                _ = v15
                            }
                        default:
                            v4 = (v4 & 2147483648)
                        }
                    case 554:
                        for fi19 := 0; fi19 < 7; fi19++ {
                            var v20 int32 = int32(fi19)
                            _ = fi19
                            _ = v20
                        }
                    }
                    if ((("foo bar" > "a") || true) && (false && (-1 != 127))) {
                        v1 -= (-(v1))
                        for fi21 := 0; fi21 < 9; fi21++ {
                            switch v2 {
                            case 32767:
                                switch (fi21 + (v2 << ((446) & 63))) {
                                case -591:
                                    v22 := GenS1{F0: 0, F1: 127, F2: "hello"}
                                    _ = v22
                                case -1:
                                    var v23 uint64 = 42
                                    _ = v23
                                default:
                                    v4 = v4
                                }
                            case 32768:
                                v24 := GenS5{F0: (v9 / ((v9) | 1)), F1: v9, F2: 100}
                                v25 := GenS2{F0: v4, F1: "a", F2: 2}
                                _ = v24
                                _ = v25
                            case 9223372036854775807:
                                var v26 int16 = (v1 | (v1 | v1))
                                _ = v26
                            default:
                                if (false && (2147483647 <= int32(v2))) {
                                    fmt.Println(fi10, v2)
                                    v9 = v9
                                    var v27 int32 = -795
                                    var v28 uint32 = 2147483647
                                    _ = v27
                                    _ = v28
                                }
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                v4 = (v4 >> ((0) & 63))
                                v29 := GenS5{F0: (v9 << ((255) & 63)), F1: (v9 ^ v9), F2: 1}
                                _ = v29
                            }
                            _ = fi21
                        }
                        var v30 float32 = 1e10
                        _ = v30
                    } else {
                        var v31 string = ((("\n\t" + "hello") + ("a" + "hello")) + ("a" + "a"))
                        _ = v31
                    }
                }
                _ = v12
            }
            _ = fi10
            _ = v11
        }
    } else {
        var v32 int8 = 100
        v33 := GenS5{F0: v9, F1: 832, F2: v32}
        if (true || (v4 > v4)) {
            var v34 uint64 = 42
            _ = v34
        } else {
            v32 -= v32
            var v35 float64 = 1.5
            if (42 <= 255) {
                v36 := GenS2{F0: v4, F1: "", F2: 0}
                var v37 uint8 = 128
                v38 := []int{32767, v2}
                _ = v36
                _ = v37
                _ = v38
            } else {
                v32 += v32
            }
            _ = v35
        }
        _ = v32
        _ = v33
    }
    var v39 uint64 = (v9 ^ (v9 | 4294967295))
    for fi40 := 0; fi40 < 2; fi40++ {
        fmt.Println(v4, v4)
        v39 += 630
        switch ("a" + "\n\t") {
        case "hello":
            for fi41 := 0; fi41 < 1; fi41++ {
                if false {
                    if ((111 < 1) && false) {
                        v2 = (fi40 - 0)
                        v2 += 9223372036854775807
                    }
                }
                v42 := GenS4{F0: ((!true) || (v4 != 4294967295)), F1: 0, F2: v1}
                v4 = (v4 | (^(v4)))
                _ = fi41
                _ = v42
            }
        case "\n\t":
            v43 := GenS6{F0: 42}
            _ = v43
        case "a":
            v44 := GenS7{F0: (v4 - 728), F1: -9223372036854775808}
            _ = v44
        }
        _ = fi40
    }
    for fi45 := 0; fi45 < 1; fi45++ {
        if (!(0.0 == ((float64(0.0) + 3.14159) / 1e10))) {
            v9 = (v9 + v39)
            if (!false) {
                if true {
                    v9 = 42
                    var v46 uint16 = uint16(v39)
                    _ = v46
                }
                v4 -= v4
                v47 := GenS4{F0: false, F1: -1, F2: (v1 >> ((0) & 63))}
                _ = v47
            }
        }
        if false {
            v4 -= v4
            var v48 uint32 = uint32(v9)
            switch (-(v2)) {
            case 2147483647:
                var v49 int8 = -1
                _ = v49
            case 0:
                v50 := []int{-865, 42}
                _ = v50
            default:
                if (false && (true || false)) {
                    for fi51 := 0; fi51 < 9; fi51++ {
                        v9 -= uint64(v4)
                        _ = fi51
                    }
                }
            }
            _ = v48
        } else {
            v4 = v4
            delete(v3, 4294967295)
            for fi52 := 0; fi52 < 9; fi52++ {
                if false {
                    v9 = v9
                    v53 := GenS7{F0: 4294967295, F1: -1}
                    _ = v53
                }
                v54 := []int{(fi52 + fi45), (v2 >> ((65535) & 63))}
                _ = fi52
                _ = v54
            }
        }
        v55 := []int{v2, v2, 42}
        _ = fi45
        _ = v55
    }
    v56 := GenS8{F0: 2147483647, F1: 911, F2: 0}
    if true {
        switch v2 {
        case 128:
            if (-9223372036854775808 != 241) {
                if false {
                    for fi57 := 0; fi57 < 5; fi57++ {
                        v58 := GenS9{F0: 443}
                        if true {
                            v59 := []int{830, (fi57 + fi57), 4294967295}
                            for fi60 := 0; fi60 < 13; fi60++ {
                                _ = len(v3)
                                _ = fi60
                            }
                            switch ("hello" + ("hello" + "")) {
                            case "foo bar":
                                if true {
                                    switch (((("foo bar" + "") + ("hello" + ("a" + ("" + "a")))) + "") + "a") {
                                    case "hello":
                                        fmt.Println(v2, v1)
                                    case "a":
                                        _ = len(v3)
                                    case "\n\t":
                                        if ((0.1 + (3.14159 + (0.1 - (float32(1e10) - ((0.0 - 3.14159) + (1.5 * (1e10 + (0.1 - 1.5)))))))) == (0.0 - float32(3.14159))) {
                                            v2 = fi57
                                            var v61 float64 = (float64(-2.25) * -2.25)
                                            _ = v61
                                        }
                                    }
                                } else {
                                    v39 = (v39 ^ 1)
                                    var v62 uint16 = 539
                                    if true {
                                        v2 -= v2
                                    }
                                    _ = v62
                                }
                            }
                            _ = v59
                        }
                        switch fi57 {
                        case -9223372036854775808:
                            fmt.Println(v4)
                        default:
                            v63 := []int{}
                            _ = v63
                        }
                        _ = fi57
                        _ = v58
                    }
                    if true {
                        v64 := GenS7{F0: 323, F1: -785}
                        if (!(18446744073709551615 >= 42)) {
                            v4 = (v4 / ((v4) | 1))
                            v65 := GenS10{F0: 744, F1: 2147483647, F2: 42}
                            v4 = 42
                            _ = v65
                        } else {
                            var v66 int = v2
                            _ = v66
                        }
                        for fi67 := 0; fi67 < 17; fi67++ {
                            switch ("" + "hello") {
                            case "hello":
                                v9 = 0
                            case "foo bar":
                                var v68 int64 = 539
                                _ = v68
                            case "\n\t":
                                v9 -= v39
                            default:
                                for fi69 := 0; fi69 < 1; fi69++ {
                                    v70 := GenS7{F0: v4, F1: 1}
                                    _ = fi69
                                    _ = v70
                                }
                            }
                            if false {
                                for fi71 := 0; fi71 < 9; fi71++ {
                                    var v72 uint16 = 483
                                    v73 := v3[v4]
                                    if false {
                                        switch fi67 {
                                        case 9223372036854775807:
                                            fmt.Println(v2, v2)
                                        case 128:
                                            fmt.Println(v1, v39)
                                        case 32768:
                                            v74 := GenS11{F0: (v1 >> ((1) & 63))}
                                            _ = v74
                                        }
                                        v39 += v9
                                        var v75 bool = true
                                        fmt.Println(v73, fi71)
                                        _ = v75
                                    } else {
                                        if (-1 > v1) {
                                            v76 := GenS12{F0: 517, F1: fi71}
                                            _ = v76
                                        }
                                        switch ((v73 + "a") + (v73 + v73)) {
                                        case "a":
                                            switch (fi71 * fi67) {
                                            case 32768:
                                                switch fi67 {
                                                case -456:
                                                    v73 = (v73 + v73)
                                                case -717:
                                                    switch (v73 + (v73 + v73)) {
                                                    case "":
                                                        fmt.Println(fi67, fi71, fi71)
                                                    default:
                                                        if false {
                                                            if false {
                                                                v2 = -9223372036854775808
                                                            } else {
                                                                var v77 string = (("foo bar" + "a") + "a")
                                                                _ = v77
                                                            }
                                                        }
                                                    }
                                                case 32767:
                                                    v1 = v1
                                                default:
                                                    switch v73 {
                                                    case "hello":
                                                        switch (fi67 % ((fi71) | 1)) {
                                                        case 0:
                                                            if (false || false) {
                                                                v78 := GenS9{F0: 9223372036854775807}
                                                                var v79 int32 = 42
                                                                _ = v78
                                                                _ = v79
                                                            } else {
                                                                if false {
                                                                    if true {
                                                                        if (444 <= v72) {
                                                                            var v80 uint8 = 0
                                                                            var v81 uint = 1
                                                                            if (!true) {
                                                                                switch v73 {
                                                                                case "hello":
                                                                                    switch (fi67 % ((fi71) | 1)) {
                                                                                    case 4294967295:
                                                                                        v39 = v39
                                                                                    case -479:
                                                                                        if false {
                                                                                            if (!true) {
                                                                                                if ("a" < v73) {
                                                                                                    var v82 uint64 = (v9 | (v39 + (v9 & (v9 - (v39 + (v39 & v39))))))
                                                                                                    var v83 float32 = 1.5
                                                                                                    _ = v82
                                                                                                    _ = v83
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            var v84 int32 = 1
                                                                                            v1 = v1
                                                                                            v9 -= 585
                                                                                            switch ("hello" + (("hello" + (v73 + (("" + (((v73 + (("\n\t" + ("" + v73)) + "\n\t")) + ((v73 + v73) + ((v73 + v73) + ((v73 + "hello") + "\n\t")))) + v73)) + (v73 + v73)))) + v73)) {
                                                                                            case "foo bar":
                                                                                                var v85 uint64 = 18446744073709551615
                                                                                                _ = v85
                                                                                            case "":
                                                                                                var v86 uint64 = v39
                                                                                                _ = v86
                                                                                            }
                                                                                            _ = v84
                                                                                        }
                                                                                    }
                                                                                case "":
                                                                                    fmt.Println(v1, v39, v81)
                                                                                case "a":
                                                                                    v87 := GenS1{F0: 128, F1: -1, F2: "foo bar"}
                                                                                    v4 += v4
                                                                                    var v88 uint8 = v80
                                                                                    var v89 int16 = v1
                                                                                    v90 := GenS13{F0: 4}
                                                                                    _ = v87
                                                                                    _ = v88
                                                                                    _ = v89
                                                                                    _ = v90
                                                                                default:
                                                                                    var v91 float64 = (3.14159 * (3.14159 + 1.5))
                                                                                    _ = v91
                                                                                }
                                                                                var v92 string = v73
                                                                                if ((true || (true && (0 < -32))) || (v9 < (v9 | v9))) {
                                                                                    if true {
                                                                                        _ = len(v3)
                                                                                        v72 = v72
                                                                                        v9 += (v39 & v9)
                                                                                    } else {
                                                                                        fmt.Println(v39)
                                                                                        var v93 float64 = 1.5
                                                                                        v94 := GenS14{F0: v4, F1: 223, F2: v9}
                                                                                        _ = v93
                                                                                        _ = v94
                                                                                    }
                                                                                } else {
                                                                                    switch int(v80) {
                                                                                    case 2147483647:
                                                                                        v3[(v4 | (v4 ^ v4))] = ""
                                                                                    case 9223372036854775807:
                                                                                        v1 = (-(v1))
                                                                                    case -101:
                                                                                        v95 := GenS10{F0: 2147483647, F1: -172, F2: (^(v81))}
                                                                                        _ = v95
                                                                                    }
                                                                                }
                                                                                v92 = (v92 + (((("hello" + "") + v92) + v73) + v92))
                                                                                if true {
                                                                                    v96 := []int{(^(fi71)), (v2 >> (((v81 % ((v81) | 1))) & 63))}
                                                                                    var v97 uint = (v81 - uint(v80))
                                                                                    switch fi71 {
                                                                                    case 32767:
                                                                                        v9 += (^(v39))
                                                                                    case -2147483648:
                                                                                        if (-1 < -1) {
                                                                                            v97 = v81
                                                                                            v72 += uint16(v81)
                                                                                            fmt.Println(v81)
                                                                                        } else {
                                                                                            var v98 uint8 = (v80 % (((^(v80))) | 1))
                                                                                            v9 += v39
                                                                                            _ = v98
                                                                                        }
                                                                                    case 598:
                                                                                        var v99 uint32 = (v4 | (v4 & (^(v4))))
                                                                                        _ = v99
                                                                                    }
                                                                                    _ = v96
                                                                                    _ = v97
                                                                                }
                                                                                if false {
                                                                                    v100 := []int{}
                                                                                    _ = v100
                                                                                } else {
                                                                                    v101 := GenS9{F0: 128}
                                                                                    _ = v101
                                                                                }
                                                                                _ = v92
                                                                            }
                                                                            _ = v80
                                                                            _ = v81
                                                                        } else {
                                                                            v1 -= 32767
                                                                            fmt.Println(v4)
                                                                            if (!(4294967295 >= 65535)) {
                                                                                if ((v39 << ((uint(v9)) & 63)) >= (^(v39))) {
                                                                                    switch (("hello" + v73) + (v73 + v73)) {
                                                                                    case "a":
                                                                                        v9 -= v9
                                                                                    case "":
                                                                                        fmt.Println(v1, v1, fi67)
                                                                                    case "foo bar":
                                                                                        var v102 uint8 = 64
                                                                                        _ = v102
                                                                                    }
                                                                                    var v103 string = v73
                                                                                    if (true || ((v9 ^ v39) < 18446744073709551615)) {
                                                                                        switch (^(v2)) {
                                                                                        case 690:
                                                                                            v104 := GenS15{F0: 726, F1: 706}
                                                                                            _ = v104
                                                                                        case 9223372036854775807:
                                                                                            v73 = v73
                                                                                        case -2147483648:
                                                                                            var v105 uint64 = v39
                                                                                            _ = v105
                                                                                        default:
                                                                                            var v106 uint = 18446744073709551615
                                                                                            var v107 uint32 = v4
                                                                                            var v108 int64 = 1
                                                                                            _ = v106
                                                                                            _ = v107
                                                                                            _ = v108
                                                                                        }
                                                                                    }
                                                                                    _ = v103
                                                                                } else {
                                                                                    v73 = v73
                                                                                    var v109 uint64 = v9
                                                                                    v110 := GenS16{F0: 2147483648, F1: 32767}
                                                                                    _ = v109
                                                                                    _ = v110
                                                                                }
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    var v111 int64 = -1
                                                                                    _ = v111
                                                                                case "hello":
                                                                                    v112 := GenS17{F0: false, F1: ((208 != -77) && true)}
                                                                                    _ = v112
                                                                                case "foo bar":
                                                                                    if false {
                                                                                        v113 := GenS18{F0: 236, F1: false}
                                                                                        _ = v113
                                                                                    } else {
                                                                                        v114 := GenS10{F0: 974, F1: -215, F2: 255}
                                                                                        _ = v114
                                                                                    }
                                                                                default:
                                                                                    var v115 float32 = ((1e10 / 1.5) * 1e10)
                                                                                    _ = v115
                                                                                }
                                                                                var v116 int = -1
                                                                                v39 = (v9 % ((v9) | 1))
                                                                                _ = len(v3)
                                                                                _ = v116
                                                                            } else {
                                                                                switch (v73 + v73) {
                                                                                case "a":
                                                                                    var v117 int = -2147483648
                                                                                    _ = v117
                                                                                case "hello":
                                                                                    switch ((v73 + ((v73 + (v73 + v73)) + ("a" + (v73 + v73)))) + ("foo bar" + v73)) {
                                                                                    case "":
                                                                                        var v118 uint16 = 111
                                                                                        _ = v118
                                                                                    case "a":
                                                                                        if false {
                                                                                            var v119 float32 = float32(1.5)
                                                                                            v1 = (v1 ^ v1)
                                                                                            if true {
                                                                                                v73 = (v73 + "foo bar")
                                                                                                var v120 uint = 42
                                                                                                v121 := []int{v2}
                                                                                                _ = v120
                                                                                                _ = v121
                                                                                            } else {
                                                                                                if (!(!("\n\t" < v73))) {
                                                                                                    var v122 string = v73
                                                                                                    _ = v122
                                                                                                } else {
                                                                                                    if (v4 >= v4) {
                                                                                                        v72 -= v72
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            _ = v119
                                                                                        } else {
                                                                                            var v123 float32 = -2.25
                                                                                            switch (fi71 + fi67) {
                                                                                            case -276:
                                                                                                switch v73 {
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v73, v2)
                                                                                                case "":
                                                                                                    fmt.Println(v72, v73, v72)
                                                                                                case "\n\t":
                                                                                                    if (!false) {
                                                                                                        v124 := GenS19{F0: "\n\t", F1: -1}
                                                                                                        _ = v124
                                                                                                    }
                                                                                                }
                                                                                            case 9223372036854775807:
                                                                                                var v125 float64 = ((0.0 / -2.25) * 0.1)
                                                                                                _ = v125
                                                                                            default:
                                                                                                v2 = fi67
                                                                                            }
                                                                                            v123 -= v123
                                                                                            _ = v123
                                                                                        }
                                                                                    }
                                                                                case "\n\t":
                                                                                    fmt.Println(v4, fi67, fi67)
                                                                                    v126 := v3[(v4 + v4)]
                                                                                    var v127 uint32 = v4
                                                                                    if ("foo bar" > (v73 + v73)) {
                                                                                        var v128 int32 = 2147483647
                                                                                        fmt.Println(v1)
                                                                                        switch (v126 + v73) {
                                                                                        case "foo bar":
                                                                                            var v129 float32 = (3.14159 - float32(1.5))
                                                                                            var v130 int = (fi67 + fi71)
                                                                                            v72 -= v72
                                                                                            _ = v129
                                                                                            _ = v130
                                                                                        }
                                                                                        _ = v128
                                                                                    }
                                                                                    if (3.14159 < (1.5 / 1e10)) {
                                                                                        v131 := GenS4{F0: true, F1: 845, F2: (v1 * 42)}
                                                                                        var v132 int64 = 0
                                                                                        _ = v131
                                                                                        _ = v132
                                                                                    }
                                                                                    v133 := []int{v2, fi67, (^(fi71)), (v2 % (((fi71 - fi67)) | 1))}
                                                                                    _ = v126
                                                                                    _ = v127
                                                                                    _ = v133
                                                                                default:
                                                                                    var v134 float64 = float64(((1.5 * (1.5 - float32(3.14159))) * (-2.25 * 3.14159)))
                                                                                    _ = v134
                                                                                }
                                                                                v72 = (v72 & v72)
                                                                                if (v9 > v9) {
                                                                                    v73 = v73
                                                                                    var v135 float64 = (1.5 + (1e10 - 0.1))
                                                                                    v4 += (^(v4))
                                                                                    _ = v135
                                                                                } else {
                                                                                    v136 := GenS20{F0: v72, F1: (v4 + (v4 >> ((626) & 63)))}
                                                                                    _ = v136
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v137 := GenS21{F0: 4294967295, F1: true}
                                                                        _ = v137
                                                                    }
                                                                    fmt.Println(fi67, v2, v1)
                                                                    var v138 string = v73
                                                                    _ = v138
                                                                }
                                                                var v139 uint32 = (v4 | v4)
                                                                v140 := []int{-135}
                                                                _ = v139
                                                                _ = v140
                                                            }
                                                        }
                                                    default:
                                                        var v141 bool = (true && (int(v1) < fi67))
                                                        _ = v141
                                                    }
                                                }
                                            case -198:
                                                switch v73 {
                                                case "foo bar":
                                                    fmt.Println(v4)
                                                case "a":
                                                    v142 := GenS5{F0: 1, F1: v39, F2: -1}
                                                    _ = v142
                                                case "hello":
                                                    v143 := GenS22{F0: int64(v1)}
                                                    _ = v143
                                                default:
                                                    var v144 uint16 = v72
                                                    _ = v144
                                                }
                                            case 128:
                                                if true {
                                                    v1 = v1
                                                    v145 := []int{}
                                                    if true {
                                                        var v146 float64 = float64(float32(((0.1 + 1.5) + -2.25)))
                                                        v39 -= v9
                                                        v147 := GenS23{F0: -2147483648, F1: (v72 << ((132) & 63)), F2: 0}
                                                        _ = v146
                                                        _ = v147
                                                    }
                                                    _ = v145
                                                } else {
                                                    v148 := []int{-1}
                                                    switch (fi67 % (((-(fi71))) | 1)) {
                                                    case 32767:
                                                        v1 = -32768
                                                    }
                                                    var v149 bool = true
                                                    _ = v148
                                                    _ = v149
                                                }
                                            default:
                                                v150 := []int{fi67, v2, (fi67 + (fi67 >> ((18446744073709551615) & 63))), (fi71 | (^(fi67)))}
                                                _ = v150
                                            }
                                        }
                                        var v151 bool = true
                                        _ = v151
                                    }
                                    _ = fi71
                                    _ = v72
                                    _ = v73
                                }
                                if (59 != 42) {
                                    fmt.Println(fi67)
                                    v152 := GenS7{F0: v4, F1: -1}
                                    for fi153 := 0; fi153 < 2; fi153++ {
                                        v3[v4] = "\n\t"
                                        _ = fi153
                                    }
                                    _ = v152
                                } else {
                                    for fi154 := 0; fi154 < 9; fi154++ {
                                        v9 -= (^(v9))
                                        v2 = (^(v2))
                                        _ = fi154
                                    }
                                }
                                v155 := []int{}
                                _ = v155
                            } else {
                                fmt.Println(v9)
                                v3[(v4 / ((v4) | 1))] = "foo bar"
                            }
                            _ = fi67
                        }
                        _ = v64
                    }
                }
            } else {
                var v156 float64 = 0.1
                switch (("" + "hello") + "\n\t") {
                case "\n\t":
                    fmt.Println(v39)
                case "hello":
                    v156 += v156
                case "a":
                    if false {
                        var v157 string = "\n\t"
                        var v158 uint8 = 128
                        switch (v157 + v157) {
                        case "":
                            v159 := GenS14{F0: uint32(v158), F1: 65535, F2: (^(v9))}
                            _ = v159
                        case "hello":
                            var v160 float32 = -2.25
                            _ = v160
                        default:
                            var v161 int32 = -2147483648
                            _ = v161
                        }
                        _ = v157
                        _ = v158
                    }
                }
                _ = v156
            }
        case -349:
            v2 -= (v2 & v2)
        default:
            v2 = -605
        }
    }
    v162 := GenS24{F0: 1, F1: -2147483648, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v162
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v56
    _ = v8
    _ = v9
}
