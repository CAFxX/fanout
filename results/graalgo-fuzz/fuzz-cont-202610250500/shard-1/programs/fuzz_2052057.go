package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint8; F1 int32; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 string; F1 string }
type GenS12 struct { F0 bool; F1 uint8; F2 bool }
func genFn0() int8 {
    switch "hello" {
    case "a":
        var v4 uint = 255
        _ = v4
    case "hello":
        var v5 int32 = 2147483647
        _ = v5
    case "foo bar":
        var v6 string = "\n\t"
        _ = v6
    default:
        switch "foo bar" {
        case "a":
            var v7 uint32 = 42
            _ = v7
        case "hello":
            var v8 uint8 = 42
            _ = v8
        case "foo bar":
            var v9 int64 = -9223372036854775808
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 uint64 = 0
        if (((false || false) && true) || false) {
            for fi12 := 0; fi12 < 16; fi12++ {
                fmt.Println(v11, v11)
                v11 -= (v11 | v11)
                var v13 int32 = -2147483648
                _ = fi12
                _ = v13
            }
            var v14 uint32 = 2147483648
            if ((-2.25 == (1.5 + 1e10)) || ((false && true) && false)) {
                if true {
                    var v15 int16 = 42
                    _ = v15
                } else {
                    v14 = 902
                }
                fmt.Println(v14)
                v16 := []int{(fi10 + fi10)}
                _ = v16
            }
            _ = v14
        }
        _ = fi10
        _ = v11
    }
    return 99
}
func genFn1(p0 int8, p1 bool) int16 {
    var v17 uint16 = 1
    v18 := GenS1{F0: "a", F1: 157, F2: 42}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return -1
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]uint32)
    var v3 string = v1
    v19 := []int{32767, -1}
    if (!false) {
        v3 = v1
    } else {
        v20 := GenS2{F0: 1, F1: 42, F2: -163}
        for fi21 := 0; fi21 < 13; fi21++ {
            if (!(((((float32(((((0.0 / 1.5) - 3.14159) - ((1.5 + (1.5 - 1.5)) / 1.0)) + 0.1)) - 3.14159) * (((float32(3.14159) * 1e10) + 1.5) - 3.14159)) / 0.1) <= 0.0) && false)) {
                v22 := GenS1{F0: (v3 + v1), F1: -479, F2: 757}
                _ = v22
            } else {
                for fi23 := 0; fi23 < 6; fi23++ {
                    v2[fi21] = 20
                    var v24 uint = 42
                    if true {
                        v1 = (v1 + ((v1 + v1) + v3))
                        v25 := v2[(fi23 / (((fi21 >> (((v24 & (v24 + 678))) & 63))) | 1))]
                        if false {
                            v24 = v24
                            v26 := GenS1{F0: v3, F1: 0, F2: (v25 >> ((2) & 63))}
                            _ = v26
                        }
                        _ = v25
                    } else {
                        if true {
                            switch (fi23 & 32768) {
                            case 42:
                                fmt.Println(v1, fi23)
                            case 4294967295:
                                delete(v2, (fi21 | (fi21 >> ((v24) & 63))))
                            }
                            var v27 int8 = genFn0()
                            var v28 float32 = 1e10
                            _ = v27
                            _ = v28
                        } else {
                            var v29 int8 = 0
                            v30 := GenS3{F0: 18446744073709551615}
                            v3 = ((("a" + v3) + v1) + v3)
                            _ = v29
                            _ = v30
                        }
                    }
                    v24 = v24
                    fmt.Println(v24, fi21)
                    _ = fi23
                    _ = v24
                }
                var v31 int16 = genFn1(100, true)
                _ = v31
            }
            var v32 uint32 = 4294967295
            if false {
                v1 = v1
                v33 := GenS4{F0: 65535, F1: 589}
                var v34 string = (v1 + v1)
                _ = v33
                _ = v34
            } else {
                var v35 uint64 = 94
                _ = v35
            }
            _ = fi21
            _ = v32
        }
        for fi36 := 0; fi36 < 2; fi36++ {
            var v37 float32 = 0.1
            v19 = append(v19, fi36)
            var v38 uint8 = 255
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = v20
    }
    var v39 uint = 2
    for fi40 := 0; fi40 < 2; fi40++ {
        for fi41 := 0; fi41 < 17; fi41++ {
            var v42 uint16 = 1
            var v43 float64 = 0.1
            _ = fi41
            _ = v42
            _ = v43
        }
        var v44 bool = false
        var v45 int64 = 328
        _ = fi40
        _ = v44
        _ = v45
    }
    v3 = ""
    v2[42] = 42
    fmt.Println(v1, v3)
    fmt.Println(v1, v39, v39)
    switch -430 {
    case -612:
        var v46 int8 = 0
        _ = v46
    case 2147483647:
        if len(v19) > 0 {
            v19[0] = -712
        }
    default:
        v1 = ("" + v1)
    }
    switch 264 {
    case 128:
        v1 = v1
    case -528:
        switch -9223372036854775808 {
        case -917:
            v47 := GenS5{F0: (v39 | v39)}
            _ = v47
        case 32768:
            v3 = v1
        }
    }
    if (-2147483648 <= 562) {
        switch ("" + "a") {
        case "foo bar":
            var v48 float64 = (0.0 + -2.25)
            _ = v48
        case "\n\t":
            if (791 == 4294967295) {
                switch 1 {
                case 1:
                    if true {
                        v49 := GenS4{F0: v39, F1: 0}
                        v3 = "hello"
                        if (-790 <= -411) {
                            v1 = ""
                        }
                        _ = v49
                    }
                case -705:
                    for fi50 := 0; fi50 < 3; fi50++ {
                        var v51 int8 = 0
                        if len(v19) > 0 {
                            v19[0] = 4294967295
                        }
                        if (v39 < 1) {
                            var v52 int8 = v51
                            _ = v52
                        }
                        _ = fi50
                        _ = v51
                    }
                }
                fmt.Println(v3, v39, v1)
            }
        default:
            v19 = append(v19, -1)
        }
        if true {
            switch 128 {
            case 128:
                if (true || (82 > 2)) {
                    switch ((v3 + (v1 + v3)) + v1) {
                    case "hello":
                        var v53 float32 = 1.5
                        _ = v53
                    default:
                        v3 = ""
                    }
                    for fi54 := 0; fi54 < 19; fi54++ {
                        if (805 != 0) {
                            if (9223372036854775807 != -691) {
                                var v55 uint32 = 48
                                if (!((true || (128 > 255)) && true)) {
                                    if false {
                                        if len(v19) >= 2 {
                                            v19 = v19[1:len(v19)]
                                        }
                                        v55 -= 65535
                                        v56 := GenS6{F0: 11}
                                        _ = v56
                                    } else {
                                        _ = len(v19)
                                        if (9223372036854775807 > -1) {
                                            v57 := GenS3{F0: 4294967295}
                                            var v58 int16 = genFn1(127, (!true))
                                            if false {
                                                if true {
                                                    v1 = v1
                                                    fmt.Println(fi54, fi54, v58)
                                                } else {
                                                    var v59 uint32 = 2147483647
                                                    var v60 int8 = -1
                                                    _ = v59
                                                    _ = v60
                                                }
                                                var v61 uint16 = 252
                                                if (false || (135 != 1)) {
                                                    v19 = append(v19, (-(fi54)))
                                                    for fi62 := 0; fi62 < 15; fi62++ {
                                                        v39 = (v39 << (((v39 >> ((v39) & 63))) & 63))
                                                        v63 := GenS7{F0: 2}
                                                        var v64 int = 9223372036854775807
                                                        _ = fi62
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                    var v65 bool = (v55 <= (v55 / (((v55 & v55)) | 1)))
                                                    _ = v65
                                                } else {
                                                    if (v61 != (^(v61))) {
                                                        var v66 float32 = -2.25
                                                        v55 += (v55 / (((v55 ^ uint32(v61))) | 1))
                                                        _ = v66
                                                    } else {
                                                        fmt.Println(v55)
                                                    }
                                                }
                                                _ = v61
                                            } else {
                                                v3 = v3
                                                fmt.Println(v3, v3)
                                                var v67 int8 = 127
                                                _ = v67
                                            }
                                            _ = v57
                                            _ = v58
                                        }
                                        for fi68 := 0; fi68 < 13; fi68++ {
                                            fmt.Println(v3, v39, v3)
                                            switch (fi68 * -1) {
                                            case 387:
                                                v2[(^(fi54))] = (^(v55))
                                            case -9223372036854775808:
                                                var v69 uint16 = 606
                                                _ = v69
                                            case 0:
                                                _ = len(v2)
                                            }
                                            v2[128] = (v55 & v55)
                                            v1 = v1
                                            if len(v19) >= 2 {
                                                v19 = v19[1:len(v19)]
                                            }
                                            var v70 int8 = 42
                                            v55 = (^(v55))
                                            _ = fi68
                                            _ = v70
                                        }
                                    }
                                }
                                var v71 uint16 = 271
                                _ = v55
                                _ = v71
                            } else {
                                if (!((!((((0.0 * (0.1 / (3.14159 / 0.1))) * 1e10) - -2.25) <= (((1e10 / (float32((1e10 + 0.0)) + float32(1.5))) + (-2.25 - (0.0 / 1.0))) + (1e10 * -2.25)))) || (42 != 0))) {
                                    if true {
                                        v3 = v3
                                        var v72 int32 = 494
                                        _ = v72
                                    } else {
                                        for fi73 := 0; fi73 < 10; fi73++ {
                                            switch -714 {
                                            case 2147483647:
                                                var v74 float64 = 0.1
                                                _ = v74
                                            }
                                            _ = fi73
                                        }
                                        for fi75 := 0; fi75 < 18; fi75++ {
                                            if (255 == v39) {
                                                v76 := GenS1{F0: "\n\t", F1: 42, F2: 2147483647}
                                                _ = v76
                                            }
                                            var v77 int64 = -1
                                            switch v1 {
                                            case "":
                                                fmt.Println(v3, v1)
                                            default:
                                                var v78 uint = 2
                                                _ = v78
                                            }
                                            v79 := v2[(fi54 / ((42) | 1))]
                                            _ = fi75
                                            _ = v77
                                            _ = v79
                                        }
                                    }
                                    var v80 uint64 = 42
                                    _ = v80
                                }
                            }
                            for fi81 := 0; fi81 < 17; fi81++ {
                                switch ((v1 + "foo bar") + "foo bar") {
                                case "":
                                    if false {
                                        switch (((v3 + v3) + v1) + v3) {
                                        case "foo bar":
                                            var v82 float64 = 3.14159
                                            _ = v82
                                        case "a":
                                            if ((fi81 | fi54) <= -1) {
                                                var v83 uint8 = 1
                                                v39 -= 38
                                                v84 := v2[-9223372036854775808]
                                                switch ("hello" + ("foo bar" + "\n\t")) {
                                                case "hello":
                                                    v1 = ((("a" + ("a" + (v1 + "foo bar"))) + "foo bar") + (v3 + (v3 + "a")))
                                                case "a":
                                                    switch "a" {
                                                    case "a":
                                                        var v85 uint64 = 142
                                                        _ = v85
                                                    case "hello":
                                                        switch v1 {
                                                        case "hello":
                                                            switch fi54 {
                                                            case -9223372036854775808:
                                                                v86 := GenS8{F0: 1}
                                                                _ = v86
                                                            }
                                                        case "":
                                                            v87 := GenS1{F0: v3, F1: 1, F2: (v84 * v84)}
                                                            _ = v87
                                                        }
                                                    case "\n\t":
                                                        if (65535 < uint16(v83)) {
                                                            if ((v84 != v84) && (("" != ("a" + ("hello" + ((v1 + v3) + ("a" + v3))))) || false)) {
                                                                if true {
                                                                    v88 := v2[(fi54 & 128)]
                                                                    _ = len(v19)
                                                                    var v89 int8 = 42
                                                                    _ = v88
                                                                    _ = v89
                                                                }
                                                            }
                                                            v90 := GenS5{F0: (^(v39))}
                                                            _ = len(v19)
                                                            v39 += v39
                                                            _ = v90
                                                        } else {
                                                            var v91 uint8 = v83
                                                            _ = v91
                                                        }
                                                    default:
                                                        v39 += (v39 << ((v39) & 63))
                                                    }
                                                }
                                                if (0 > genFn1(105, true)) {
                                                    var v92 float32 = 0.1
                                                    _ = v92
                                                }
                                                switch 127 {
                                                case 496:
                                                    fmt.Println(fi81, v39)
                                                    if (false || true) {
                                                        v3 = v3
                                                        var v93 int8 = 0
                                                        fmt.Println(v84, fi54, v39)
                                                        _ = v93
                                                    } else {
                                                        if true {
                                                            v94 := GenS5{F0: v39}
                                                            _ = v94
                                                        }
                                                    }
                                                    if len(v19) >= 2 {
                                                        v19 = v19[1:len(v19)]
                                                    }
                                                }
                                                if ((1.5 + ((((((1e10 + 3.14159) * 3.14159) * (1e10 / 1e10)) - ((3.14159 - (1e10 / (3.14159 / 0.1))) / 1.5)) + -2.25) - 0.1)) == 3.14159) {
                                                    if true {
                                                        v1 = v1
                                                        v19 = append(v19, (fi54 << ((uint(v83)) & 63)))
                                                        fmt.Println(v3, v3)
                                                    }
                                                    v95 := GenS9{F0: v83}
                                                    var v96 uint16 = 1
                                                    _ = v95
                                                    _ = v96
                                                }
                                                v84 += v84
                                                _ = v83
                                                _ = v84
                                            } else {
                                                v19 = append(v19, -1)
                                            }
                                        default:
                                            if len(v19) > 0 {
                                                v19[0] = fi81
                                            }
                                        }
                                        var v97 int64 = 42
                                        v98 := GenS5{F0: v39}
                                        _ = v97
                                        _ = v98
                                    }
                                default:
                                    v3 = v1
                                }
                                var v99 uint32 = 42
                                _ = fi81
                                _ = v99
                            }
                            var v100 string = ("foo bar" + "")
                            _ = v100
                        } else {
                            v3 = ("a" + ((("hello" + "\n\t") + "a") + ""))
                            v101 := GenS1{F0: v3, F1: -863, F2: 0}
                            var v102 float32 = 0.0
                            _ = v101
                            _ = v102
                        }
                        fmt.Println(v1, v1, fi54)
                        v39 -= (^(v39))
                        _ = fi54
                    }
                } else {
                    v103 := GenS10{F0: 0}
                    for fi104 := 0; fi104 < 8; fi104++ {
                        v3 = v1
                        var v105 bool = false
                        var v106 int8 = -76
                        _ = fi104
                        _ = v105
                        _ = v106
                    }
                    v107 := GenS11{F0: "\n\t", F1: v3}
                    _ = v103
                    _ = v107
                }
            case -2147483648:
                var v108 uint = v39
                _ = v108
            }
            var v109 float32 = ((-2.25 / (1e10 + 0.1)) * 1.5)
            _ = v109
        }
        v1 = "foo bar"
    } else {
        fmt.Println(v3, v1)
        if ((((0.0 * 1e10) + 1.5) > 0.1) && true) {
            if ((0 > 140) || (-318 == 9223372036854775807)) {
                fmt.Println(v3, v39, v39)
            } else {
                v3 = v1
            }
            fmt.Println(v39, v3, v1)
            for fi110 := 0; fi110 < 12; fi110++ {
                v3 = ("\n\t" + v1)
                _ = fi110
            }
        } else {
            v19 = append(v19, 127)
        }
    }
    var v111 float64 = 1.5
    for fi112 := 0; fi112 < 16; fi112++ {
        v3 = ""
        if (((true && false) && false) && true) {
            switch (fi112 - (fi112 | fi112)) {
            case -686:
                switch ((v1 + "hello") + ((v1 + v1) + "hello")) {
                case "\n\t":
                    if ((fi112 * (fi112 * (fi112 % ((fi112) | 1)))) >= fi112) {
                        switch (fi112 >> ((v39) & 63)) {
                        case 770:
                            v113 := GenS12{F0: false, F1: 128, F2: false}
                            _ = v113
                        case -731:
                            v3 = "hello"
                        case 1:
                            switch -2147483648 {
                            case -557:
                                if false {
                                    switch (fi112 / ((-621) | 1)) {
                                    case 1:
                                        if (false || (!false)) {
                                            for fi114 := 0; fi114 < 10; fi114++ {
                                                delete(v2, (fi114 & 32767))
                                                v111 += (1.5 * float64(float32(1.5)))
                                                _ = fi114
                                            }
                                            var v115 int8 = 1
                                            var v116 float64 = (0.1 + (1e10 + 1e10))
                                            _ = v115
                                            _ = v116
                                        } else {
                                            if (0 != 603) {
                                                _ = len(v2)
                                                var v117 uint = (v39 + v39)
                                                delete(v2, fi112)
                                                _ = v117
                                            }
                                            if true {
                                                var v118 uint = 1
                                                _ = v118
                                            }
                                            v119 := GenS11{F0: v3, F1: v1}
                                            _ = v119
                                        }
                                    case 2147483647:
                                        if false {
                                            fmt.Println(v3, v1)
                                            v1 = v3
                                            v2[32767] = 124
                                        }
                                    }
                                } else {
                                    if len(v19) > 0 {
                                        v19[0] = (fi112 | -57)
                                    }
                                    for fi120 := 0; fi120 < 19; fi120++ {
                                        var v121 uint = v39
                                        v3 = (v1 + "\n\t")
                                        _ = len(v19)
                                        _ = fi120
                                        _ = v121
                                    }
                                }
                            default:
                                v39 = v39
                            }
                        }
                        if len(v19) > 0 {
                            v19[0] = fi112
                        }
                        var v122 int32 = -500
                        _ = v122
                    }
                }
            case 4294967295:
                v39 += (v39 ^ v39)
            default:
                fmt.Println(fi112, fi112)
            }
            var v123 uint32 = 0
            _ = v123
        }
        v1 = "foo bar"
        _ = fi112
    }
    var v124 uint32 = 0
    var v125 float64 = v111
    v126 := GenS4{F0: (v39 * 1), F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v124:", v124)
    fmt.Println("v125:", v125)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v111
    _ = v124
    _ = v125
    _ = v126
    _ = v19
    _ = v2
    _ = v3
    _ = v39
}
