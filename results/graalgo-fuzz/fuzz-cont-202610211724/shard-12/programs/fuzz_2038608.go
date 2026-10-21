package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint8; F1 int16 }
type GenS11 struct { F0 int8; F1 uint8 }
type GenS12 struct { F0 uint16; F1 uint }
type GenS13 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS14 struct { F0 uint32; F1 uint16 }
type GenS15 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS16 struct { F0 uint64; F1 string }
func genFn0() uint {
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 float64 = (-2.25 + ((0.1 - (3.14159 - 3.14159)) / (((0.1 / (-2.25 * -2.25)) * float64(0.0)) + 1.5)))
        var v9 int = fi7
        v10 := []int{}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS1{F0: 1000, F1: 255, F2: 32767}
        if (49 != 128) {
            switch "hello" {
            case "foo bar":
                v13 := []int{127}
                _ = v13
            case "":
                var v14 uint8 = 0
                _ = v14
            }
        } else {
            fmt.Println(fi11)
            v15 := []int{(fi11 % ((fi11) | 1)), (fi11 & fi11), (-(fi11)), 1}
            _ = v15
        }
        var v16 uint = 75
        _ = fi11
        _ = v12
        _ = v16
    }
    var v17 float64 = (1e10 / (-2.25 + ((3.14159 * 1e10) / -2.25)))
    _ = v17
    return 4294967295
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int = -9223372036854775808
    v3 := []uint32{2147483647, 65535}
    v4 := make(map[int]bool)
    v5 := []int{}
    v6 := []uint32{42, 83, 92, 4294967295}
    var v18 int64 = 924
    for fi19 := 0; fi19 < 15; fi19++ {
        v1 -= v1
        v20 := GenS2{F0: -1}
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        if (("a" + "a") <= "hello") {
            _ = len(v4)
            v2 -= -697
        } else {
            if ((-240 > fi21) || false) {
                for fi22 := 0; fi22 < 14; fi22++ {
                    v23 := GenS3{F0: v1}
                    if false {
                        if (-2.25 >= 0.1) {
                            if len(v5) > 0 {
                                v5[0] = fi22
                            }
                            var v24 string = "hello"
                            _ = v24
                        } else {
                            var v25 uint = genFn0()
                            v26 := GenS4{F0: (fi21 + 32768), F1: (-398 < v18)}
                            v2 += (fi21 * (fi21 / ((int(v25)) | 1)))
                            _ = v25
                            _ = v26
                        }
                        if false {
                            switch ("\n\t" + "\n\t") {
                            case "foo bar":
                                v27 := v4[(fi21 & v2)]
                                _ = v27
                            case "a":
                                var v28 float64 = (float64(1.5) * float64((1e10 * (0.1 / -2.25))))
                                _ = v28
                            }
                            if (1 < v18) {
                                var v29 int8 = -61
                                _ = v29
                            } else {
                                var v30 bool = (0.0 >= 3.14159)
                                _ = v30
                            }
                            if (339 >= -34) {
                                v31 := GenS5{F0: 1000, F1: false}
                                var v32 uint = v1
                                var v33 string = "hello"
                                v32 -= 1
                                _ = v31
                                _ = v32
                                _ = v33
                            } else {
                                if (0.0 != 3.14159) {
                                    fmt.Println(v18, v2, fi21)
                                    if ((v18 >> (((^(v1))) & 63)) > 0) {
                                        v34 := GenS2{F0: 12}
                                        v2 += 42
                                        v35 := GenS1{F0: -1, F1: v1, F2: 2}
                                        v36 := GenS3{F0: v1}
                                        switch fi22 {
                                        case 4294967295:
                                            var v37 string = "\n\t"
                                            _ = v37
                                        default:
                                            var v38 uint = v1
                                            _ = v38
                                        }
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 643
                                    }
                                    if false {
                                        if false {
                                            var v39 int8 = -1
                                            fmt.Println(fi22, fi22)
                                            if ((!(!true)) || (0 <= 2)) {
                                                v1 -= v1
                                                switch "\n\t" {
                                                case "hello":
                                                    var v40 int8 = (^(v39))
                                                    _ = v40
                                                case "foo bar":
                                                    v18 += v18
                                                case "":
                                                    var v41 int32 = 1
                                                    _ = v41
                                                }
                                            } else {
                                                var v42 int16 = 42
                                                _ = v42
                                            }
                                            switch ("\n\t" + ("a" + "")) {
                                            case "hello":
                                                fmt.Println(v2, v1)
                                            case "":
                                                v39 -= (v39 % ((v39) | 1))
                                            }
                                            delete(v4, v2)
                                            _ = v39
                                        } else {
                                            if ((-(fi21)) != (^(fi22))) {
                                                v43 := GenS5{F0: -633, F1: false}
                                                var v44 int16 = 1
                                                v1 = v1
                                                _ = len(v3)
                                                fmt.Println(v1, v1)
                                                _ = v43
                                                _ = v44
                                            } else {
                                                v1 = (v1 & v1)
                                                fmt.Println(fi22)
                                                v1 = (v1 * (v1 * (v1 + (v1 ^ (v1 + v1)))))
                                            }
                                            v1 = (v1 - 948)
                                        }
                                        switch fi21 {
                                        case -2147483648:
                                            var v45 int8 = 0
                                            fmt.Println(v18, v18, v18)
                                            _ = v45
                                        case 32767:
                                            fmt.Println(v2, fi21)
                                        case 803:
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                        if false {
                                            var v46 float32 = 0.0
                                            var v47 float32 = 1.5
                                            var v48 int8 = 119
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        }
                                        v49 := GenS2{F0: 127}
                                        _ = len(v5)
                                        _ = v49
                                    } else {
                                        var v50 uint64 = 42
                                        _ = v50
                                    }
                                    if (32767 == 0) {
                                        switch fi21 {
                                        case -159:
                                            _ = len(v4)
                                        }
                                    } else {
                                        v18 -= (v18 / ((v18) | 1))
                                    }
                                    v51 := GenS1{F0: 581, F1: v1, F2: 2}
                                    _ = v51
                                } else {
                                    v52 := GenS3{F0: (^(v1))}
                                    var v53 uint8 = 27
                                    _ = v52
                                    _ = v53
                                }
                                var v54 string = (("" + ("\n\t" + (("hello" + "a") + "\n\t"))) + "foo bar")
                                _ = v54
                            }
                        }
                        _ = len(v4)
                    } else {
                        v18 -= (v18 >> ((42) & 63))
                        if true {
                            v18 = v18
                        } else {
                            _ = len(v4)
                            if (128 == (-(fi22))) {
                                if false {
                                    var v55 uint8 = 2
                                    _ = v55
                                }
                            }
                        }
                    }
                    if ((!(193 < uint8(v1))) && false) {
                        switch v2 {
                        case -746:
                            v1 -= 42
                        case -553:
                            v56 := GenS6{F0: 1, F1: true}
                            _ = v56
                        case 4294967295:
                            var v57 uint = genFn0()
                            _ = v57
                        }
                    }
                    _ = fi22
                    _ = v23
                }
                if true {
                    _ = len(v4)
                    if true {
                        v2 = v2
                        fmt.Println(fi21, fi21)
                    }
                } else {
                    fmt.Println(v1)
                }
            } else {
                var v58 int64 = v18
                fmt.Println(v58)
                _ = v58
            }
        }
        var v59 uint16 = 2
        for fi60 := 0; fi60 < 16; fi60++ {
            var v61 int64 = 9223372036854775807
            if true {
                v62 := GenS7{F0: -727}
                _ = v62
            }
            v4[v2] = false
            _ = fi60
            _ = v61
        }
        _ = fi21
        _ = v59
    }
    var v63 int8 = 42
    for fi64 := 0; fi64 < 19; fi64++ {
        v4[fi64] = true
        _ = fi64
    }
    v65 := GenS5{F0: 431, F1: (v2 == (v2 / (((v2 * (v2 & (^(v2))))) | 1)))}
    for fi66 := 0; fi66 < 6; fi66++ {
        switch ("hello" + "") {
        case "a":
            v2 += fi66
        default:
            v4[(^(fi66))] = (-2.25 == 1e10)
        }
        v63 = -128
        switch (fi66 | 4294967295) {
        case 9223372036854775807:
            var v67 bool = true
            _ = v67
        default:
            v5 = append(v5, (v2 ^ (v2 << ((v1) & 63))))
        }
        _ = fi66
    }
    var v68 int16 = 42
    if false {
        if (!(1.5 < (3.14159 / 1e10))) {
            var v69 float64 = (-2.25 / 1.0)
            _ = len(v4)
            v70 := GenS7{F0: -32768}
            _ = v69
            _ = v70
        } else {
            var v71 uint16 = 0
            var v72 uint = v1
            _ = v71
            _ = v72
        }
        fmt.Println(v18, v68, v18)
        v73 := GenS8{F0: ("a" + "hello")}
        _ = v73
    } else {
        switch "a" {
        case "hello":
            switch (("a" + ("foo bar" + ("foo bar" + "foo bar"))) + ((("" + "hello") + ("a" + "a")) + "")) {
            case "hello":
                if false {
                    var v74 float32 = 0.0
                    v63 += int8(v2)
                    fmt.Println(v2)
                    _ = v74
                } else {
                    v1 += 2
                    var v75 int16 = (v68 >> (((v1 * 18446744073709551615)) & 63))
                    v76 := GenS8{F0: ""}
                    _ = v75
                    _ = v76
                }
            default:
                v18 = (v18 - v18)
            }
        case "a":
            if true {
                fmt.Println(v1)
            } else {
                switch v2 {
                case 4294967295:
                    for fi77 := 0; fi77 < 13; fi77++ {
                        var v78 int = (v2 - v2)
                        v78 = (v2 | fi77)
                        v2 -= 128
                        _ = fi77
                        _ = v78
                    }
                }
            }
        default:
            for fi79 := 0; fi79 < 19; fi79++ {
                switch ("foo bar" + "") {
                case "foo bar":
                    switch ("hello" + "foo bar") {
                    case "":
                        if (!false) {
                            v80 := GenS4{F0: (^(v2)), F1: true}
                            v63 -= v63
                            _ = v80
                        }
                    case "a":
                        var v81 uint32 = 0
                        _ = v81
                    default:
                        for fi82 := 0; fi82 < 6; fi82++ {
                            var v83 int = (v2 << ((v1) & 63))
                            v84 := GenS9{F0: v18}
                            if (0 > 1) {
                                var v85 uint8 = 128
                                _ = v85
                            }
                            var v86 string = "hello"
                            if true {
                                switch v2 {
                                case 128:
                                    v87 := GenS8{F0: "a"}
                                    _ = v87
                                }
                            }
                            var v88 uint32 = 1
                            v89 := GenS10{F0: 2, F1: v68}
                            _ = fi82
                            _ = v83
                            _ = v84
                            _ = v86
                            _ = v88
                            _ = v89
                        }
                    }
                case "hello":
                    var v90 float64 = ((1.5 + (float64(1.5) / 3.14159)) * 1.5)
                    _ = v90
                }
                fmt.Println(v18, fi79, v63)
                _ = fi79
            }
        }
        if true {
            switch 32767 {
            case 777:
                v63 = v63
            case -2147483648:
                var v91 float64 = 0.1
                _ = v91
            default:
                var v92 float32 = 0.0
                _ = v92
            }
            _ = len(v4)
            v2 += 42
        } else {
            if true {
                fmt.Println(v18)
                v4[v2] = false
            } else {
                if true {
                    for fi93 := 0; fi93 < 18; fi93++ {
                        if false {
                            for fi94 := 0; fi94 < 6; fi94++ {
                                v4[2147483647] = (v68 == v68)
                                var v95 int16 = (v68 / (((v68 * v68)) | 1))
                                if (!false) {
                                    _ = len(v5)
                                } else {
                                    v96 := GenS6{F0: v63, F1: ((v68 * v68) != v68)}
                                    _ = v96
                                }
                                _ = fi94
                                _ = v95
                            }
                            if true {
                                switch ("foo bar" + "a") {
                                case "":
                                    switch int(v18) {
                                    case -463:
                                        v68 = v68
                                    case 42:
                                        if true {
                                            if false {
                                                if (("\n\t" == ("hello" + "\n\t")) || (2 < 4294967295)) {
                                                    v63 -= (v63 - 1)
                                                }
                                                v68 = -957
                                                var v97 string = ("" + "foo bar")
                                                _ = v97
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = 0
                                            }
                                            v18 = v18
                                        }
                                    case -970:
                                        v98 := GenS10{F0: 2, F1: v68}
                                        _ = v98
                                    }
                                default:
                                    switch (v2 >> (((v1 + v1)) & 63)) {
                                    case 74:
                                        if (42 != 216) {
                                            v99 := GenS5{F0: v68, F1: (false && ((0.0 / 1e10) != (-2.25 / 1.0)))}
                                            var v100 uint32 = 2
                                            v101 := GenS11{F0: int8(v1), F1: 228}
                                            _ = v100
                                            _ = v101
                                            _ = v99
                                        }
                                    case 565:
                                        delete(v4, (v2 / (((^(v2))) | 1)))
                                    case 9223372036854775807:
                                        if ((v18 * -9223372036854775808) != (-(v18))) {
                                            var v102 int8 = v63
                                            v103 := GenS3{F0: 1}
                                            v104 := GenS12{F0: 0, F1: (v1 % (((^(v1))) | 1))}
                                            _ = v102
                                            _ = v103
                                            _ = v104
                                        } else {
                                            fmt.Println(v1, fi93)
                                            var v105 int16 = v68
                                            _ = v105
                                        }
                                    default:
                                        v18 += (v18 ^ v18)
                                    }
                                }
                                var v106 uint8 = 255
                                if (1 >= 0) {
                                    var v107 bool = false
                                    for fi108 := 0; fi108 < 3; fi108++ {
                                        if (true || v107) {
                                            if v107 {
                                                var v109 int8 = v63
                                                _ = v109
                                            }
                                            v107 = (v18 == v18)
                                            var v110 float64 = 0.0
                                            _ = v110
                                        }
                                        if (((1 == 514) || (1 < 2147483648)) || v107) {
                                            var v111 uint8 = (v106 ^ v106)
                                            _ = v111
                                        } else {
                                            v107 = (769 != (v1 | v1))
                                        }
                                        _ = fi108
                                    }
                                    switch v2 {
                                    case 32767:
                                        var v112 int32 = 42
                                        _ = v112
                                    }
                                    _ = v107
                                } else {
                                    fmt.Println(v63)
                                }
                                _ = v106
                            } else {
                                var v113 uint16 = 42
                                _ = v113
                            }
                            var v114 int8 = (-(v63))
                            _ = v114
                        } else {
                            if false {
                                if len(v6) > 0 {
                                    v6[0] = 2147483648
                                }
                            }
                        }
                        if (!(2 == 1)) {
                            var v115 int = fi93
                            for fi116 := 0; fi116 < 9; fi116++ {
                                if true {
                                    var v117 int16 = v68
                                    var v118 float64 = float64((1e10 / 1.5))
                                    _ = v117
                                    _ = v118
                                }
                                switch (v2 >> ((v1) & 63)) {
                                case 42:
                                    var v119 uint16 = 131
                                    _ = v119
                                }
                                switch "hello" {
                                case "foo bar":
                                    switch "" {
                                    case "":
                                        if false {
                                            switch "\n\t" {
                                            case "":
                                                if ("hello" >= "a") {
                                                    v2 += 2147483647
                                                    _ = len(v3)
                                                    fmt.Println(v1)
                                                }
                                            case "foo bar":
                                                var v120 int16 = v68
                                                _ = v120
                                            case "\n\t":
                                                v68 = v68
                                            default:
                                                v63 = 100
                                            }
                                            _ = len(v4)
                                        }
                                    default:
                                        var v121 int16 = -786
                                        _ = v121
                                    }
                                }
                                _ = fi116
                            }
                            var v122 uint32 = 4294967295
                            _ = v115
                            _ = v122
                        }
                        switch (v2 >> (((v1 | (^(v1)))) & 63)) {
                        case 842:
                            v6 = append(v6, 952)
                        case 345:
                            for fi123 := 0; fi123 < 16; fi123++ {
                                v4[(fi123 * (fi93 << ((v1) & 63)))] = true
                                _ = fi123
                            }
                        case 773:
                            var v124 string = (("hello" + "\n\t") + "\n\t")
                            _ = v124
                        }
                        _ = fi93
                    }
                }
                fmt.Println(v63, v63)
                if ((0.0 * 0.1) > 1.5) {
                    var v125 int16 = v68
                    for fi126 := 0; fi126 < 8; fi126++ {
                        if ((-1 == v63) || false) {
                            for fi127 := 0; fi127 < 15; fi127++ {
                                v125 -= v125
                                var v128 string = "a"
                                _ = fi127
                                _ = v128
                            }
                        }
                        fmt.Println(v2, v63)
                        if true {
                            v129 := GenS9{F0: -9223372036854775808}
                            var v130 float64 = ((-2.25 * -2.25) / (1.5 + 0.1))
                            _ = v129
                            _ = v130
                        }
                        _ = fi126
                    }
                    v125 -= v68
                    _ = v125
                } else {
                    if (((1.5 / 3.14159) + ((1e10 + -2.25) / (3.14159 + 0.0))) > float64(1e10)) {
                        fmt.Println(v68, v1)
                    }
                    var v131 uint32 = 1
                    if (4294967295 > 42) {
                        if true {
                            var v132 uint8 = 143
                            _ = v132
                        } else {
                            v1 += genFn0()
                            v1 += genFn0()
                            if ((false || (0 < 0)) && (6 > 127)) {
                                var v133 int64 = (v18 * (v18 + (^(v18))))
                                for fi134 := 0; fi134 < 16; fi134++ {
                                    for fi135 := 0; fi135 < 16; fi135++ {
                                        var v136 uint8 = 243
                                        var v137 int64 = (v18 / ((9223372036854775807) | 1))
                                        _ = fi135
                                        _ = v136
                                        _ = v137
                                    }
                                    v138 := GenS13{F0: 851, F1: v63, F2: v131}
                                    for fi139 := 0; fi139 < 9; fi139++ {
                                        _ = len(v5)
                                        if true {
                                            v140 := GenS5{F0: v68, F1: ((255 != 113) || true)}
                                            switch "\n\t" {
                                            case "":
                                                fmt.Println(v133, v1)
                                            case "\n\t":
                                                switch (fi139 ^ (^(v2))) {
                                                case 1:
                                                    var v141 uint = v1
                                                    _ = v141
                                                case 128:
                                                    if true {
                                                        if false {
                                                            _ = len(v4)
                                                            switch (("hello" + (("foo bar" + "hello") + "hello")) + "hello") {
                                                            case "hello":
                                                                v142 := GenS14{F0: (v131 / (((v131 | (^(v131)))) | 1)), F1: 0}
                                                                _ = v142
                                                            default:
                                                                if false {
                                                                    if (false && false) {
                                                                        v68 -= (-(v68))
                                                                    } else {
                                                                        v68 -= v68
                                                                        v143 := GenS15{F0: (^(v68)), F1: 1, F2: 410}
                                                                        var v144 int64 = (v18 * v133)
                                                                        v145 := GenS11{F0: int8(v144), F1: 145}
                                                                        v1 += 255
                                                                        _ = v143
                                                                        _ = v144
                                                                        _ = v145
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            v1 = 42
                                                        }
                                                        if (float32(((-2.25 / -2.25) * 3.14159)) > 1.5) {
                                                            switch "a" {
                                                            case "hello":
                                                                fmt.Println(v68, v68)
                                                            case "":
                                                                if len(v5) > 0 {
                                                                    v5[0] = 97
                                                                }
                                                            case "a":
                                                                v68 = (-(v68))
                                                            }
                                                            if (true || (((!(true && (!((((^(v63)) < (v63 >> (((^(v1))) & 63))) || true) && false)))) || false) || true)) {
                                                                if len(v3) > 0 {
                                                                    v3[0] = (v131 << ((v1) & 63))
                                                                }
                                                                v2 -= fi134
                                                            }
                                                            if (true && (6 != 0)) {
                                                                fmt.Println(fi139)
                                                                v1 = (v1 & (^(v1)))
                                                            } else {
                                                                v146 := GenS11{F0: int8(v68), F1: 0}
                                                                _ = v146
                                                            }
                                                            var v147 uint8 = 255
                                                            _ = v147
                                                        }
                                                    }
                                                case 4294967295:
                                                    v63 -= -1
                                                }
                                            }
                                            var v148 bool = ((v63 >> ((0) & 63)) < v63)
                                            if (v148 || v148) {
                                                v149 := GenS13{F0: 4294967295, F1: -128, F2: v131}
                                                if ((!(!(-9223372036854775808 >= -394))) && v148) {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                } else {
                                                    if (int8(v133) < -128) {
                                                        v150 := GenS12{F0: 65535, F1: genFn0()}
                                                        _ = v150
                                                    }
                                                    var v151 int8 = v63
                                                    _ = v151
                                                }
                                                _ = v149
                                            }
                                            if ((v2 >> ((255) & 63)) > (fi134 ^ fi134)) {
                                                v152 := v4[(v2 << ((v1) & 63))]
                                                var v153 uint64 = 4294967295
                                                v68 = (v68 + (v68 | (v68 * v68)))
                                                _ = v152
                                                _ = v153
                                            } else {
                                                v131 -= v131
                                                fmt.Println(v18)
                                                if (!false) {
                                                    var v154 int32 = 0
                                                    if v148 {
                                                        v154 = (-(v154))
                                                        var v155 int32 = (v154 + 2147483647)
                                                        _ = v155
                                                    } else {
                                                        fmt.Println(fi134, v68, fi139)
                                                        fmt.Println(v154)
                                                    }
                                                    _ = v154
                                                }
                                            }
                                            _ = v140
                                            _ = v148
                                        } else {
                                            v131 = 796
                                        }
                                        var v156 int16 = (v68 + -32768)
                                        _ = fi139
                                        _ = v156
                                    }
                                    _ = fi134
                                    _ = v138
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v133
                            } else {
                                v18 = v18
                            }
                        }
                        v68 = (^(v68))
                        var v157 bool = false
                        _ = v157
                    } else {
                        v5 = append(v5, 9223372036854775807)
                        v158 := GenS8{F0: ""}
                        _ = v158
                    }
                    _ = v131
                }
            }
            for fi159 := 0; fi159 < 13; fi159++ {
                var v160 uint = v1
                switch "a" {
                case "\n\t":
                    var v161 int = (v2 | v2)
                    _ = v161
                case "a":
                    var v162 bool = (!true)
                    _ = v162
                case "hello":
                    var v163 uint = (v1 << ((v1) & 63))
                    _ = v163
                default:
                    var v164 float32 = float32(((0.0 * (0.0 + 1.5)) * 0.1))
                    _ = v164
                }
                _ = fi159
                _ = v160
            }
        }
        v2 += -2147483648
    }
    v165 := GenS16{F0: 128, F1: (("\n\t" + ((("foo bar" + "") + (("foo bar" + "foo bar") + "foo bar")) + ("a" + "a"))) + ("hello" + ("foo bar" + "a")))}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v63:", v63)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v165
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v65
    _ = v68
}
