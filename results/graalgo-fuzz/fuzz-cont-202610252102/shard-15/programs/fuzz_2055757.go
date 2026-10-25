package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 int; F2 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 string; F2 int32 }
type GenS7 struct { F0 int32; F1 int8; F2 uint }
type GenS8 struct { F0 uint; F1 uint }
type GenS9 struct { F0 uint16; F1 bool }
type GenS10 struct { F0 uint64; F1 int64; F2 int }
type GenS11 struct { F0 uint16; F1 uint }
type GenS12 struct { F0 uint; F1 int32; F2 int }
type GenS13 struct { F0 int64; F1 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 255
    var v3 int64 = 42
    var v4 uint8 = 42
    var v5 int8 = 100
    var v6 uint32 = 15
    fmt.Println(v4)
    v4 = 1
    v7 := GenS1{F0: v4, F1: v2, F2: (95 != 2)}
    switch 128 {
    case 731:
        v8 := GenS2{F0: true}
        _ = v8
    case -220:
        if (!false) {
            switch 2147483647 {
            case 42:
                fmt.Println(v2, v4)
            case -2147483648:
                var v9 uint8 = v4
                _ = v9
            case 1:
                switch -970 {
                case -1:
                    if (-1 == 42) {
                        v5 = 127
                        for fi10 := 0; fi10 < 13; fi10++ {
                            v6 += v6
                            fmt.Println(v5, fi10, v4)
                            _ = fi10
                        }
                    }
                case 128:
                    if ((v2 | v2) == v2) {
                        if (75 < 32767) {
                            for fi11 := 0; fi11 < 16; fi11++ {
                                var v12 int32 = 42
                                var v13 int = fi11
                                var v14 float64 = 0.1
                                _ = fi11
                                _ = v12
                                _ = v13
                                _ = v14
                            }
                            var v15 bool = (3.14159 == (((1.5 - 3.14159) * 1e10) / 1.5))
                            fmt.Println(v5)
                            _ = v15
                        } else {
                            v6 = (v6 / ((0) | 1))
                            var v16 uint = v2
                            v17 := GenS3{F0: v3}
                            _ = v16
                            _ = v17
                        }
                        switch ((("foo bar" + "\n\t") + "a") + "") {
                        case "hello":
                            v6 -= 0
                        case "a":
                            var v18 float64 = 3.14159
                            _ = v18
                        }
                        v19 := GenS3{F0: (v3 & 1)}
                        _ = v19
                    } else {
                        var v20 int = -167
                        switch "" {
                        case "hello":
                            var v21 uint16 = uint16(v4)
                            _ = v21
                        default:
                            if (1e10 <= 0.0) {
                                fmt.Println(v5)
                                switch "\n\t" {
                                case "hello":
                                    var v22 float32 = -2.25
                                    _ = v22
                                case "foo bar":
                                    v2 += (^(v2))
                                default:
                                    for fi23 := 0; fi23 < 13; fi23++ {
                                        fmt.Println(v6, v6, v5)
                                        _ = fi23
                                    }
                                }
                                switch "\n\t" {
                                case "\n\t":
                                    v3 = (v3 ^ v3)
                                default:
                                    v3 = (v3 - -9223372036854775808)
                                }
                            }
                        }
                        var v24 uint16 = 343
                        _ = v20
                        _ = v24
                    }
                case 1:
                    if true {
                        v25 := []int{-20, int(v5), 0, 9223372036854775807}
                        var v26 int = -705
                        _ = v25
                        _ = v26
                    }
                default:
                    var v27 float64 = (3.14159 - 3.14159)
                    _ = v27
                }
            default:
                v4 = 128
            }
        } else {
            v28 := []int{9223372036854775807, 2147483647, -1}
            _ = v28
        }
    case 127:
        var v29 uint64 = 1
        _ = v29
    default:
        v30 := GenS4{F0: 18446744073709551615, F1: 32767, F2: ("foo bar" + (("foo bar" + "a") + "\n\t"))}
        _ = v30
    }
    if false {
        var v31 string = ("\n\t" + "\n\t")
        _ = v31
    } else {
        switch "a" {
        case "hello":
            v1[(^(v6))] = 127
        default:
            v6 = v6
        }
        v1[(v6 << ((v2) & 63))] = 32768
        for fi32 := 0; fi32 < 12; fi32++ {
            switch (fi32 * fi32) {
            case 32768:
                for fi33 := 0; fi33 < 2; fi33++ {
                    if false {
                        if true {
                            v34 := GenS5{F0: 1000}
                            v5 = int8(v6)
                            _ = v34
                        } else {
                            delete(v1, 653)
                            var v35 int64 = v3
                            switch ((("hello" + ("foo bar" + (("\n\t" + ("hello" + ("foo bar" + ("\n\t" + "")))) + ("foo bar" + "a")))) + "hello") + ("\n\t" + "a")) {
                            case "foo bar":
                                switch ("" + "hello") {
                                case "\n\t":
                                    var v36 uint32 = uint32(v2)
                                    _ = v36
                                }
                            case "\n\t":
                                if true {
                                    var v37 uint8 = v4
                                    _ = v37
                                }
                            default:
                                if (v6 <= v6) {
                                    v2 -= 145
                                } else {
                                    v35 -= v3
                                    fmt.Println(v5, v2, v3)
                                }
                            }
                            _ = v35
                        }
                        fmt.Println(fi32)
                    }
                    fmt.Println(fi32)
                    delete(v1, 2)
                    var v38 uint16 = 32767
                    _ = fi33
                    _ = v38
                }
            case 127:
                if true {
                    v4 -= 1
                    var v39 int16 = 32767
                    v40 := []int{fi32}
                    _ = v39
                    _ = v40
                }
            case 42:
                fmt.Println(v6)
            default:
                var v41 bool = false
                _ = v41
            }
            switch (-(fi32)) {
            case 128:
                for fi42 := 0; fi42 < 14; fi42++ {
                    switch 32767 {
                    case -169:
                        switch ("\n\t" + "") {
                        case "\n\t":
                            switch "foo bar" {
                            case "":
                                v6 = v6
                            case "foo bar":
                                v4 = (v4 + (v4 >> ((v2) & 63)))
                            default:
                                switch fi32 {
                                case -2147483648:
                                    v43 := GenS2{F0: (-1 > 32767)}
                                    _ = v43
                                default:
                                    if true {
                                        var v44 float32 = 0.1
                                        _ = v44
                                    } else {
                                        var v45 float32 = 0.1
                                        switch "foo bar" {
                                        case "foo bar":
                                            fmt.Println(fi32)
                                        }
                                        _ = v45
                                    }
                                }
                            }
                        }
                    case -2147483648:
                        v5 = (v5 - v5)
                    case 127:
                        var v46 int = fi32
                        _ = v46
                    default:
                        v6 = (v6 * v6)
                    }
                    _ = fi42
                }
            default:
                delete(v1, (v6 >> ((v2) & 63)))
            }
            delete(v1, (v6 << ((v2) & 63)))
            _ = fi32
        }
    }
    var v47 int64 = -1
    var v48 int16 = 1000
    var v49 uint16 = 65535
    v3 = (v47 | (v47 >> ((v2) & 63)))
    switch (("\n\t" + "\n\t") + "\n\t") {
    case "a":
        switch "" {
        case "hello":
            if false {
                for fi50 := 0; fi50 < 16; fi50++ {
                    var v51 string = ("\n\t" + "")
                    var v52 float64 = 0.0
                    _ = fi50
                    _ = v51
                    _ = v52
                }
                fmt.Println(v47, v6, v48)
            } else {
                if false {
                    v53 := GenS6{F0: (v48 % (((v48 >> (((v2 + v2)) & 63))) | 1)), F1: "foo bar", F2: 1}
                    _ = v53
                } else {
                    var v54 int = -1
                    if true {
                        v5 = v5
                        fmt.Println(v3)
                    }
                    _ = v54
                }
                var v55 int = 0
                _ = v55
            }
        case "foo bar":
            var v56 int = 127
            _ = v56
        default:
            var v57 string = ("hello" + ("hello" + ("foo bar" + ("a" + "a"))))
            _ = v57
        }
    case "foo bar":
        for fi58 := 0; fi58 < 17; fi58++ {
            if false {
                var v59 int16 = v48
                var v60 int64 = (^(v3))
                for fi61 := 0; fi61 < 1; fi61++ {
                    fmt.Println(v49)
                    v62 := GenS6{F0: v48, F1: "a", F2: 42}
                    _ = fi61
                    _ = v62
                }
                _ = v59
                _ = v60
            } else {
                if ("foo bar" <= "a") {
                    switch "" {
                    case "\n\t":
                        fmt.Println(v2, v47, fi58)
                    }
                }
            }
            fmt.Println(v48)
            _ = fi58
        }
    case "":
        if (v49 != (v49 * (v49 ^ (v49 & 954)))) {
            if (4294967295 <= 127) {
                for fi63 := 0; fi63 < 6; fi63++ {
                    for fi64 := 0; fi64 < 14; fi64++ {
                        fmt.Println(v4, v6)
                        switch fi64 {
                        case 1:
                            fmt.Println(v48, v48)
                        }
                        _ = fi64
                    }
                    _ = fi63
                }
                if (((!false) && ((v49 - uint16(v4)) >= v49)) || (!false)) {
                    v48 = v48
                    for fi65 := 0; fi65 < 17; fi65++ {
                        if false {
                            v66 := v1[(^(v6))]
                            v67 := []int{-704}
                            _ = v66
                            _ = v67
                        }
                        for fi68 := 0; fi68 < 3; fi68++ {
                            v2 = (v2 & v2)
                            _ = fi68
                        }
                        switch (((("foo bar" + ("\n\t" + "\n\t")) + "hello") + "\n\t") + "\n\t") {
                        case "":
                            var v69 uint16 = v49
                            _ = v69
                        case "hello":
                            var v70 uint32 = (v6 ^ v6)
                            _ = v70
                        default:
                            var v71 int16 = int16(v6)
                            _ = v71
                        }
                        _ = fi65
                    }
                    v2 = v2
                } else {
                    v72 := GenS7{F0: -1, F1: (v5 - (-(v5))), F2: v2}
                    fmt.Println(v4, v4, v3)
                    switch ("foo bar" + ("\n\t" + "\n\t")) {
                    case "hello":
                        if false {
                            for fi73 := 0; fi73 < 15; fi73++ {
                                v3 -= 1
                                for fi74 := 0; fi74 < 15; fi74++ {
                                    if true {
                                        switch fi74 {
                                        case 2147483647:
                                            v75 := []int{}
                                            _ = v75
                                        }
                                    }
                                    _ = fi74
                                }
                                v76 := GenS8{F0: v2, F1: uint(v49)}
                                _ = fi73
                                _ = v76
                            }
                        } else {
                            switch 0 {
                            case 42:
                                if ("" > "a") {
                                    for fi77 := 0; fi77 < 10; fi77++ {
                                        if false {
                                            fmt.Println(v4)
                                            for fi78 := 0; fi78 < 18; fi78++ {
                                                if (false || true) {
                                                    switch fi77 {
                                                    case 0:
                                                        if true {
                                                            var v79 float32 = 1.5
                                                            if (!(true || false)) {
                                                                var v80 uint64 = 902
                                                                var v81 uint = (v2 % (((v2 + 4294967295)) | 1))
                                                                _ = v80
                                                                _ = v81
                                                            } else {
                                                                var v82 int16 = 0
                                                                _ = v82
                                                            }
                                                            _ = v79
                                                        } else {
                                                            if (("\n\t" >= "\n\t") && true) {
                                                                var v83 uint16 = v49
                                                                _ = v83
                                                            }
                                                            var v84 int16 = -51
                                                            v85 := GenS4{F0: 751, F1: (fi77 >> ((v2) & 63)), F2: "\n\t"}
                                                            _ = v84
                                                            _ = v85
                                                        }
                                                    case 4294967295:
                                                        v6 = 2147483648
                                                    }
                                                    if false {
                                                        v86 := GenS9{F0: v49, F1: false}
                                                        _ = v86
                                                    }
                                                } else {
                                                    v4 = (v4 & 0)
                                                }
                                                _ = fi78
                                            }
                                            v87 := v1[42]
                                            _ = v87
                                        }
                                        v49 = uint16(v4)
                                        _ = fi77
                                    }
                                    fmt.Println(v3, v4, v5)
                                    v88 := GenS9{F0: 668, F1: (!(255 != 255))}
                                    _ = v88
                                }
                            case -9223372036854775808:
                                if (1.5 <= 1e10) {
                                    var v89 uint32 = v6
                                    _ = v89
                                } else {
                                    var v90 int = 2147483647
                                    var v91 string = "hello"
                                    switch v91 {
                                    case "foo bar":
                                        if false {
                                            _ = len(v1)
                                            fmt.Println(v5)
                                        }
                                    }
                                    _ = v90
                                    _ = v91
                                }
                            case 9223372036854775807:
                                for fi92 := 0; fi92 < 8; fi92++ {
                                    v5 = (v5 ^ v5)
                                    _ = fi92
                                }
                            }
                            v3 = (-(v47))
                        }
                    case "\n\t":
                        var v93 int8 = v5
                        _ = v93
                    }
                    _ = v72
                }
            } else {
                if true {
                    v94 := GenS4{F0: 18446744073709551615, F1: 1, F2: "\n\t"}
                    v95 := GenS10{F0: 749, F1: v3, F2: -672}
                    if ((v2 | (^(v2))) > v2) {
                        fmt.Println(v4, v5)
                        var v96 uint32 = (v6 % ((v6) | 1))
                        _ = v96
                    } else {
                        v97 := GenS11{F0: v49, F1: 286}
                        v47 -= -447
                        if (v4 < (^(v4))) {
                            v47 -= (-(v47))
                            if (!false) {
                                v4 = (v4 << ((v2) & 63))
                                var v98 int32 = 0
                                var v99 int32 = v98
                                _ = v98
                                _ = v99
                            }
                        } else {
                            v3 -= v3
                            v100 := GenS2{F0: true}
                            _ = v100
                        }
                        _ = v97
                    }
                    _ = v94
                    _ = v95
                } else {
                    v5 += v5
                    var v101 float64 = 3.14159
                    var v102 uint16 = (v49 * v49)
                    _ = v101
                    _ = v102
                }
            }
            v1[(v6 * v6)] = 128
        } else {
            if (-2.25 < ((1e10 - (0.1 * (((0.0 - float64(1e10)) - 1.5) - float64(-2.25)))) * ((0.1 + 3.14159) * -2.25))) {
                v6 -= (v6 - v6)
                for fi103 := 0; fi103 < 10; fi103++ {
                    fmt.Println(v4, v5)
                    _ = fi103
                }
                for fi104 := 0; fi104 < 15; fi104++ {
                    fmt.Println(v49, v48, v48)
                    v48 -= (v48 / ((-1) | 1))
                    _ = fi104
                }
            } else {
                if true {
                    v105 := GenS12{F0: uint(v4), F1: 0, F2: 2147483647}
                    _ = v105
                }
                if ((2147483647 == 42) || ("hello" > (("hello" + "foo bar") + "hello"))) {
                    v3 = (v3 * 0)
                } else {
                    var v106 uint16 = (v49 & (v49 / ((v49) | 1)))
                    var v107 int16 = -79
                    _ = v106
                    _ = v107
                }
            }
            fmt.Println(v3)
            if true {
                var v108 int16 = v48
                v109 := GenS13{F0: v3, F1: 32767}
                var v110 uint64 = 18446744073709551615
                _ = v108
                _ = v109
                _ = v110
            }
        }
    default:
        switch ("hello" + "") {
        case "foo bar":
            var v111 uint8 = 1
            _ = v111
        default:
            var v112 string = "foo bar"
            _ = v112
        }
    }
    v113 := GenS9{F0: (^(v49)), F1: (v6 >= (v6 ^ 938))}
    for fi114 := 0; fi114 < 15; fi114++ {
        var v115 uint16 = uint16(v6)
        v116 := GenS8{F0: 4294967295, F1: v2}
        v117 := GenS11{F0: (^(v49)), F1: v2}
        _ = fi114
        _ = v115
        _ = v116
        _ = v117
    }
    v118 := v1[509]
    fmt.Println(v5, v118)
    fmt.Println(v47)
    for fi119 := 0; fi119 < 7; fi119++ {
        v120 := []int{int(v47), (-(v118)), fi119, (-(v118))}
        v3 = (v47 ^ 42)
        _ = fi119
        _ = v120
    }
    v6 = v6
    v6 -= (v6 * v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v118:", v118)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v113
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
