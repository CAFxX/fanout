package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint; F1 int; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int; F1 string; F2 uint32 }
type GenS6 struct { F0 string; F1 uint8; F2 string }
type GenS7 struct { F0 int; F1 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 string; F1 uint32 }
func genFn0() uint8 {
    var v7 string = (("a" + ((("hello" + (((("" + (("" + "foo bar") + "a")) + "hello") + ("\n\t" + "hello")) + "hello")) + (("\n\t" + (((("hello" + ("" + "")) + "") + ("a" + "a")) + "a")) + "hello")) + ("" + "foo bar"))) + "")
    var v8 uint64 = 710
    v9 := GenS1{F0: -341}
    _ = v7
    _ = v8
    _ = v9
    return 128
}
func genFn1(p0 uint8) int {
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint16 = 42
        v12 := GenS2{F0: (v11 * v11), F1: 65535}
        _ = fi10
        _ = v11
        _ = v12
    }
    switch -2147483648 {
    case 881:
        switch -2147483648 {
        case 32767:
            switch "hello" {
            case "a":
                if true {
                    var v13 float64 = 3.14159
                    _ = v13
                }
            }
        case 709:
            if true {
                v14 := []int{-9223372036854775808, 127, 4294967295}
                _ = v14
            }
        case -824:
            for fi15 := 0; fi15 < 7; fi15++ {
                if (!((852 >= 791) && true)) {
                    var v16 bool = (!(2 <= 255))
                    switch fi15 {
                    case -676:
                        var v17 float32 = (-2.25 - 0.1)
                        _ = v17
                    case 171:
                        if false {
                            fmt.Println(fi15)
                            var v18 int8 = 1
                            fmt.Println(fi15, v16)
                            _ = v18
                        } else {
                            for fi19 := 0; fi19 < 16; fi19++ {
                                var v20 float32 = 3.14159
                                _ = fi19
                                _ = v20
                            }
                        }
                    }
                    _ = v16
                }
                _ = fi15
            }
        default:
            for fi21 := 0; fi21 < 11; fi21++ {
                if false {
                    switch ("a" + ("foo bar" + "\n\t")) {
                    case "\n\t":
                        if false {
                            for fi22 := 0; fi22 < 13; fi22++ {
                                p0 = p0
                                _ = fi22
                            }
                            v23 := GenS2{F0: 336, F1: 4294967295}
                            var v24 float64 = (0.1 - float64(0.0))
                            _ = v23
                            _ = v24
                        } else {
                            var v25 int16 = -1
                            _ = v25
                        }
                    case "a":
                        fmt.Println(p0, fi21)
                    }
                } else {
                    v26 := []int{fi21, fi21}
                    var v27 float64 = -2.25
                    v28 := GenS3{F0: 402, F1: 4294967295, F2: 42}
                    _ = v26
                    _ = v27
                    _ = v28
                }
                v29 := GenS4{F0: 100, F1: 325}
                var v30 int32 = -1
                _ = fi21
                _ = v29
                _ = v30
            }
        }
    case 0:
        switch 0 {
        case 845:
            var v31 uint8 = genFn0()
            _ = v31
        default:
            switch 0 {
            case 778:
                var v32 float64 = 0.1
                _ = v32
            case 32768:
                var v33 uint32 = 839
                _ = v33
            case -9223372036854775808:
                var v34 uint8 = p0
                _ = v34
            default:
                v35 := []int{-297}
                _ = v35
            }
        }
    default:
        switch -315 {
        case -192:
            v36 := GenS5{F0: 0, F1: (("hello" + (("" + (("" + "foo bar") + "hello")) + "foo bar")) + "hello"), F2: 114}
            _ = v36
        case 826:
            var v37 uint32 = 65535
            _ = v37
        case -1:
            switch "foo bar" {
            case "hello":
                for fi38 := 0; fi38 < 6; fi38++ {
                    fmt.Println(fi38, p0)
                    v39 := GenS3{F0: 779, F1: fi38, F2: 1}
                    _ = fi38
                    _ = v39
                }
            case "\n\t":
                if true {
                    var v40 float64 = 3.14159
                    var v41 int = 0
                    v41 -= v41
                    _ = v40
                    _ = v41
                } else {
                    if false {
                        switch (("a" + "hello") + ("" + "a")) {
                        case "a":
                            v42 := GenS6{F0: "foo bar", F1: 127, F2: ("foo bar" + "hello")}
                            _ = v42
                        default:
                            p0 = p0
                        }
                    }
                }
            case "a":
                v43 := []int{42, 2147483647}
                _ = v43
            default:
                v44 := []int{-9223372036854775808, -720, 0}
                _ = v44
            }
        default:
            fmt.Println(p0)
        }
    }
    p0 += 29
    var v45 uint = 238
    _ = p0
    _ = v45
    return 127
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 0
    v3 := make(map[uint32]int)
    var v4 uint8 = 42
    v5 := []uint32{1, 2, 4294967295, 949}
    var v6 int64 = 1
    if false {
        var v46 uint32 = uint32(v4)
        for fi47 := 0; fi47 < 14; fi47++ {
            var v48 float32 = 1.5
            if true {
                if len(v5) > 0 {
                    v5[0] = v46
                }
                v5 = append(v5, 652)
            } else {
                switch ("foo bar" + "\n\t") {
                case "\n\t":
                    fmt.Println(v2, v48)
                case "foo bar":
                    var v49 float32 = v48
                    _ = v49
                case "":
                    delete(v3, v46)
                default:
                    _ = len(v5)
                }
            }
            _ = fi47
            _ = v48
        }
        _ = v46
    }
    v6 = v6
    _ = len(v1)
    fmt.Println(v6, v2, v6)
    if (((889 <= 18446744073709551615) || false) || (0 > 1)) {
        v50 := v3[0]
        var v51 float32 = 0.0
        v4 = (v4 & 1)
        _ = v50
        _ = v51
    } else {
        v6 = (v6 & (v6 >> ((462) & 63)))
        if true {
            for fi52 := 0; fi52 < 11; fi52++ {
                _ = len(v1)
                switch "\n\t" {
                case "foo bar":
                    for fi53 := 0; fi53 < 8; fi53++ {
                        if true {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = len(v5)
                        } else {
                            switch "" {
                            case "foo bar":
                                if false {
                                    var v54 int16 = 32767
                                    if true {
                                        v54 -= 0
                                        var v55 uint8 = (v4 | (v4 * 42))
                                        _ = v55
                                    } else {
                                        var v56 int64 = (-(v6))
                                        switch "foo bar" {
                                        case "":
                                            if true {
                                                var v57 float64 = -2.25
                                                v4 -= (v4 >> ((4294967295) & 63))
                                                fmt.Println(v56)
                                                _ = v57
                                            } else {
                                                v56 = v56
                                                var v58 int32 = -1
                                                _ = v58
                                            }
                                        case "\n\t":
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        case "foo bar":
                                            v4 = genFn0()
                                        default:
                                            var v59 bool = true
                                            _ = v59
                                        }
                                        var v60 uint = 312
                                        _ = v56
                                        _ = v60
                                    }
                                    _ = v54
                                } else {
                                    _ = len(v1)
                                }
                            }
                            var v61 uint64 = 4294967295
                            _ = v61
                        }
                        v62 := GenS1{F0: 862}
                        if ("" >= ("foo bar" + ("hello" + "foo bar"))) {
                            var v63 int = fi53
                            v64 := GenS5{F0: (fi52 / ((fi52) | 1)), F1: "foo bar", F2: 2147483648}
                            _ = v63
                            _ = v64
                        }
                        _ = fi53
                        _ = v62
                    }
                default:
                    v6 += v6
                }
                _ = fi52
            }
            switch "\n\t" {
            case "\n\t":
                for fi65 := 0; fi65 < 16; fi65++ {
                    v66 := GenS7{F0: fi65, F1: 0}
                    _ = fi65
                    _ = v66
                }
            case "":
                var v67 int = 1
                _ = v67
            case "a":
                v68 := GenS5{F0: -763, F1: "a", F2: 333}
                _ = v68
            default:
                for fi69 := 0; fi69 < 2; fi69++ {
                    for fi70 := 0; fi70 < 20; fi70++ {
                        v2 = v2
                        _ = fi70
                    }
                    _ = len(v5)
                    _ = fi69
                }
            }
        }
        var v71 uint8 = (^(v4))
        _ = v71
    }
    var v72 int = 0
    if (42 != 2) {
        var v73 uint8 = (v4 * (v4 ^ genFn0()))
        _ = v73
    }
    v4 -= genFn0()
    if false {
        _ = len(v5)
        switch "\n\t" {
        case "a":
            if (1 < 801) {
                v74 := GenS3{F0: 42, F1: genFn1(204), F2: 934}
                if false {
                    v2 += -128
                    v75 := GenS1{F0: -2147483648}
                    var v76 int = v72
                    _ = v75
                    _ = v76
                }
                var v77 float32 = float32(float64(1.5))
                _ = v74
                _ = v77
            } else {
                fmt.Println(v6)
                var v78 uint64 = 570
                v72 += int(v78)
                _ = v78
            }
        }
        if true {
            v4 = v4
        }
    } else {
        switch -9223372036854775808 {
        case 954:
            var v79 uint32 = 2147483647
            _ = v79
        case -791:
            v80 := GenS1{F0: -2147483648}
            _ = v80
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    }
    var v81 int32 = -1
    switch ("foo bar" + ("" + ("hello" + "hello"))) {
    case "\n\t":
        switch "\n\t" {
        case "\n\t":
            if true {
                var v82 bool = true
                _ = v82
            } else {
                for fi83 := 0; fi83 < 2; fi83++ {
                    for fi84 := 0; fi84 < 20; fi84++ {
                        switch (^(v72)) {
                        case 9223372036854775807:
                            switch ("a" + "foo bar") {
                            case "":
                                var v85 uint32 = 0
                                _ = v85
                            case "hello":
                                v2 = v2
                            case "foo bar":
                                var v86 uint8 = v4
                                _ = v86
                            default:
                                if true {
                                    v6 += 42
                                }
                            }
                            fmt.Println(fi84)
                        default:
                            if (954 > 413) {
                                v4 = v4
                            } else {
                                switch int(v2) {
                                case -2147483648:
                                    fmt.Println(fi84, v81)
                                default:
                                    var v87 float32 = (((3.14159 * 1e10) - (0.0 - 0.0)) - 1e10)
                                    v5 = append(v5, 2147483647)
                                    _ = v87
                                }
                                v5 = append(v5, 2147483648)
                                fmt.Println(v4)
                            }
                        }
                        _ = fi84
                    }
                    for fi88 := 0; fi88 < 8; fi88++ {
                        if (v2 > (^(v2))) {
                            if false {
                                v89 := v1[(v72 | v72)]
                                var v90 int = v72
                                fmt.Println(fi83, v2)
                                _ = v89
                                _ = v90
                            }
                            if false {
                                v4 = genFn0()
                                switch 127 {
                                case 128:
                                    v2 += (v2 * 87)
                                default:
                                    if len(v5) > 0 {
                                        v5[0] = 4294967295
                                    }
                                }
                                v91 := GenS3{F0: 18446744073709551615, F1: 1, F2: 255}
                                _ = v91
                            } else {
                                var v92 uint16 = 42
                                v72 = (fi88 ^ genFn1(127))
                                var v93 uint16 = (v92 ^ (v92 | (v92 << ((343) & 63))))
                                var v94 bool = false
                                v95 := GenS4{F0: (v2 % (((v2 | 107)) | 1)), F1: 460}
                                _ = v92
                                _ = v93
                                _ = v94
                                _ = v95
                            }
                            v81 = -1
                        } else {
                            if false {
                                if (32767 == 1000) {
                                    switch "" {
                                    case "":
                                        fmt.Println(fi88, v6)
                                    case "\n\t":
                                        if true {
                                            v96 := GenS8{F0: -9223372036854775808}
                                            _ = v96
                                        } else {
                                            fmt.Println(v6)
                                            _ = len(v1)
                                        }
                                    default:
                                        if (!(!true)) {
                                            switch "" {
                                            case "foo bar":
                                                v72 += (fi83 >> ((18446744073709551615) & 63))
                                            case "\n\t":
                                                _ = len(v3)
                                            case "":
                                                v2 = v2
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v3[42] = (fi83 / (((-(fi88))) | 1))
                                                switch (v72 % ((fi88) | 1)) {
                                                case -9223372036854775808:
                                                    _ = len(v3)
                                                }
                                                v4 = 0
                                                var v97 uint = 83
                                                fmt.Println(v97, v6)
                                                _ = v97
                                            }
                                            v81 = (v81 | (v81 | (v81 / (((v81 * (v81 * (v81 >> ((295) & 63))))) | 1))))
                                            v98 := v1[(fi83 >> ((242) & 63))]
                                            _ = v98
                                        } else {
                                            switch ("foo bar" + "") {
                                            case "a":
                                                switch ("" + "hello") {
                                                case "foo bar":
                                                    v81 += v81
                                                }
                                            default:
                                                v4 = v4
                                            }
                                            if true {
                                                v81 -= (v81 << ((255) & 63))
                                            } else {
                                                if false {
                                                    var v99 uint = 2
                                                    var v100 int8 = (v2 * v2)
                                                    _ = v100
                                                    _ = v99
                                                }
                                                if (503 >= 689) {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    if true {
                                                        v72 -= fi83
                                                        var v101 uint64 = 65
                                                        _ = v101
                                                    }
                                                } else {
                                                    var v102 int32 = (v81 << ((255) & 63))
                                                    var v103 uint = 2
                                                    var v104 string = "foo bar"
                                                    _ = v102
                                                    _ = v103
                                                    _ = v104
                                                }
                                            }
                                        }
                                    }
                                }
                                var v105 uint = 65535
                                _ = v105
                            } else {
                                v106 := GenS9{F0: genFn0()}
                                var v107 string = "a"
                                _ = v106
                                _ = v107
                            }
                        }
                        fmt.Println(v81, fi88)
                        if (v4 <= v4) {
                            v108 := GenS10{F0: 0}
                            _ = v108
                        }
                        switch "\n\t" {
                        case "\n\t":
                            switch ("a" + ("" + "hello")) {
                            case "":
                                switch "" {
                                case "":
                                    switch (("hello" + "") + (("\n\t" + "a") + ("foo bar" + (("" + "hello") + ("foo bar" + "foo bar"))))) {
                                    case "hello":
                                        var v109 uint32 = 65535
                                        _ = v109
                                    }
                                case "\n\t":
                                    if (true || (663 == 42)) {
                                        v6 = -9223372036854775808
                                    } else {
                                        switch (("" + ("foo bar" + ("foo bar" + "\n\t"))) + "a") {
                                        case "hello":
                                            var v110 int16 = -32768
                                            _ = v110
                                        }
                                        v111 := GenS4{F0: 118, F1: 622}
                                        if len(v5) > 0 {
                                            v5[0] = 2147483648
                                        }
                                        _ = v111
                                    }
                                    delete(v3, 2147483647)
                                }
                            case "a":
                                if false {
                                    switch "a" {
                                    case "foo bar":
                                        v81 = 42
                                    case "a":
                                        var v112 int64 = v6
                                        _ = v112
                                    }
                                }
                            }
                            v113 := GenS7{F0: (v72 << ((18446744073709551615) & 63)), F1: 293}
                            _ = v113
                        }
                        v72 -= -2147483648
                        _ = fi88
                    }
                    fmt.Println(v81, fi83)
                    _ = fi83
                }
                v6 = v6
                v72 -= (v72 | 42)
            }
        }
    case "":
        fmt.Println(v4, v6)
    case "hello":
        var v114 uint = 4294967295
        _ = v114
    }
    for fi115 := 0; fi115 < 14; fi115++ {
        switch fi115 {
        case 176:
            if (!(!(!((v4 ^ v4) < v4)))) {
                v2 += (v2 + v2)
            } else {
                var v116 uint = 1
                if ((((("" + "a") + "") + "") + "a") == "hello") {
                    var v117 int16 = 373
                    _ = v117
                }
                switch "\n\t" {
                case "a":
                    if false {
                        fmt.Println(v6, v2)
                        v116 = (v116 / ((v116) | 1))
                    } else {
                        v118 := GenS10{F0: 32767}
                        _ = v118
                    }
                }
                _ = v116
            }
        case 127:
            if true {
                for fi119 := 0; fi119 < 18; fi119++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = len(v1)
                    _ = fi119
                }
                v2 = -128
                for fi120 := 0; fi120 < 20; fi120++ {
                    if (42 == 488) {
                        fmt.Println(v81)
                    }
                    if (((v81 >> ((796) & 63)) <= v81) && true) {
                        v4 -= (v4 % ((v4) | 1))
                        if (false || (int(v2) != (fi115 & -802))) {
                            v6 -= (-(v6))
                        } else {
                            var v121 uint8 = v4
                            _ = v121
                        }
                        var v122 uint64 = 1
                        _ = v122
                    } else {
                        v2 -= (^(v2))
                        if false {
                            if false {
                                var v123 uint32 = 821
                                v124 := GenS10{F0: 2}
                                _ = v123
                                _ = v124
                            }
                        }
                        if (-315 == -72) {
                            if (!(false && true)) {
                                var v125 bool = (!((1e10 <= (((0.0 - 0.1) + ((3.14159 * 1.5) * 0.0)) * ((1e10 * float32(1e10)) / 0.1))) && (((!false) && false) && true)))
                                v126 := GenS8{F0: (v72 | int(v81))}
                                switch "foo bar" {
                                case "hello":
                                    v2 = (v2 / ((int8(fi115)) | 1))
                                case "":
                                    var v127 bool = false
                                    _ = v127
                                }
                                _ = v125
                                _ = v126
                            } else {
                                var v128 int32 = v81
                                if len(v5) > 0 {
                                    v5[0] = 721
                                }
                                _ = v128
                            }
                            fmt.Println(v81, fi120, v81)
                        }
                    }
                    _ = fi120
                }
            } else {
                var v129 uint16 = 327
                delete(v3, 2147483647)
                _ = v129
            }
        default:
            fmt.Println(v2)
        }
        _ = fi115
    }
    switch genFn1(255) {
    case 32768:
        var v130 float32 = 0.0
        _ = v130
    }
    v81 = (v81 >> ((51) & 63))
    v131 := GenS11{F0: "a", F1: 1}
    for fi132 := 0; fi132 < 1; fi132++ {
        switch int(v81) {
        case 1:
            var v133 bool = ((fi132 * (v72 * fi132)) >= genFn1(v4))
            _ = v133
        }
        v6 = v6
        var v134 bool = true
        _ = fi132
        _ = v134
    }
    switch 32767 {
    case 2147483647:
        if (110 >= 1000) {
            v4 -= v4
            var v135 uint32 = 2147483648
            _ = v135
        }
    default:
        var v136 string = ""
        _ = v136
    }
    switch ("" + (("foo bar" + "hello") + "\n\t")) {
    case "\n\t":
        var v137 int64 = (-(v6))
        _ = v137
    case "a":
        if (!((-32768 <= 1000) || false)) {
            fmt.Println(v4)
        } else {
            var v138 int8 = 100
            _ = v138
        }
    case "hello":
        var v139 string = ""
        _ = v139
    default:
        v2 = 92
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v131
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
    _ = v81
}
