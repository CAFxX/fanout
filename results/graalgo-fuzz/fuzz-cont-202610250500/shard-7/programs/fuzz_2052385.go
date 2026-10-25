package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint16; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 int16; F1 uint8; F2 int64 }
type GenS7 struct { F0 uint; F1 uint }
type GenS8 struct { F0 string; F1 int8; F2 int16 }
type GenS9 struct { F0 int8; F1 uint; F2 uint }
type GenS10 struct { F0 int64; F1 uint; F2 bool }
type GenS11 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS12 struct { F0 uint }
func genFn0(p0 uint16, p1 bool, p2 int) uint64 {
    if p1 {
        for fi7 := 0; fi7 < 7; fi7++ {
            p1 = (-112 > 0)
            _ = fi7
        }
        p1 = (!(("a" > "hello") && (32767 <= -20)))
        switch (("foo bar" + "hello") + "hello") {
        case "foo bar":
            if p1 {
                var v8 int32 = -1
                _ = v8
            } else {
                if (p0 < 32767) {
                    switch "" {
                    case "":
                        v9 := []int{}
                        _ = v9
                    case "foo bar":
                        if p1 {
                            p2 += (p2 + -9223372036854775808)
                            for fi10 := 0; fi10 < 4; fi10++ {
                                if true {
                                    var v11 string = ("foo bar" + "foo bar")
                                    p0 = 65535
                                    switch v11 {
                                    case "a":
                                        fmt.Println(p0)
                                    case "hello":
                                        for fi12 := 0; fi12 < 4; fi12++ {
                                            v13 := GenS1{F0: v11}
                                            _ = fi12
                                            _ = v13
                                        }
                                    case "":
                                        fmt.Println(p0)
                                    }
                                    _ = v11
                                }
                                p2 = (p2 + (fi10 & p2))
                                _ = fi10
                            }
                            for fi14 := 0; fi14 < 3; fi14++ {
                                for fi15 := 0; fi15 < 15; fi15++ {
                                    fmt.Println(fi15)
                                    switch ("hello" + "a") {
                                    case "":
                                        var v16 uint16 = 42
                                        _ = v16
                                    case "\n\t":
                                        p1 = (4294967295 <= 679)
                                    case "a":
                                        p2 += fi15
                                    }
                                    p1 = p1
                                    _ = fi15
                                }
                                _ = fi14
                            }
                        } else {
                            var v17 string = "a"
                            for fi18 := 0; fi18 < 19; fi18++ {
                                v19 := GenS1{F0: v17}
                                for fi20 := 0; fi20 < 11; fi20++ {
                                    p1 = p1
                                    v21 := []int{-499, (p2 & (fi20 + p2)), fi20, (-(p2))}
                                    p1 = (((9223372036854775807 < 614) || (485 == 0)) && (!true))
                                    _ = fi20
                                    _ = v21
                                }
                                switch v17 {
                                case "hello":
                                    switch (v17 + v17) {
                                    case "":
                                        switch "a" {
                                        case "hello":
                                            p2 -= (p2 % ((p2) | 1))
                                        default:
                                            fmt.Println(v17, v17)
                                        }
                                    }
                                }
                                _ = fi18
                                _ = v19
                            }
                            p2 = (^(p2))
                            _ = v17
                        }
                    }
                    v22 := GenS2{F0: (!(!(true && p1))), F1: (p0 * p0), F2: -1}
                    _ = v22
                } else {
                    switch p2 {
                    case -2147483648:
                        p2 = p2
                    case 128:
                        var v23 int64 = -1
                        _ = v23
                    }
                }
                for fi24 := 0; fi24 < 8; fi24++ {
                    var v25 int8 = -128
                    if ((false || (!(2 < 255))) || (0 < 2147483647)) {
                        var v26 float64 = (0.1 + (0.0 - 0.1))
                        _ = v26
                    } else {
                        p2 -= p2
                        if false {
                            v27 := []int{(p2 | fi24), 42, (p2 >> ((uint(p0)) & 63)), (p2 % ((15) | 1))}
                            p0 -= (^(p0))
                            _ = v27
                        }
                        p1 = ((!(!(2147483647 >= -730))) || true)
                    }
                    var v28 int32 = 1
                    _ = fi24
                    _ = v25
                    _ = v28
                }
                var v29 bool = p1
                _ = v29
            }
        case "a":
            var v30 float64 = 3.14159
            _ = v30
        case "hello":
            switch p2 {
            case 4294967295:
                for fi31 := 0; fi31 < 14; fi31++ {
                    var v32 uint32 = 2147483648
                    _ = fi31
                    _ = v32
                }
            case 32768:
                var v33 int32 = 1
                _ = v33
            default:
                v34 := []int{p2}
                _ = v34
            }
        }
    } else {
        fmt.Println(p2, p2)
        var v35 string = ("a" + "hello")
        var v36 uint64 = 24
        _ = v35
        _ = v36
    }
    v37 := GenS3{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v37
    return 0
}
func genFn1() string {
    v38 := []int{-23, 4294967295, 127, -9223372036854775808}
    v39 := GenS3{F0: -124}
    var v40 uint64 = genFn0(543, false, 9223372036854775807)
    var v41 int = 0
    _ = v38
    _ = v39
    _ = v40
    _ = v41
    return "foo bar"
}

func main() {
    var v1 float32 = (3.14159 * 3.14159)
    v2 := make(map[uint32]string)
    var v3 int8 = 27
    v4 := make(map[string]uint32)
    v5 := make(map[int]string)
    var v6 int16 = -524
    if true {
        v42 := []int{-213, 32768, 42, 0}
        _ = v42
    }
    v43 := GenS1{F0: ("\n\t" + "")}
    if ("foo bar" != ("a" + "")) {
        var v44 int64 = -1
        v44 += v44
        _ = v44
    } else {
        var v45 uint32 = 0
        var v46 string = ""
        _ = v45
        _ = v46
    }
    switch 2147483647 {
    case 680:
        var v47 uint16 = 65535
        _ = v47
    case 42:
        var v48 uint32 = 800
        _ = v48
    }
    fmt.Println(v1, v3)
    v3 = 99
    v2[2147483647] = ((("foo bar" + (("foo bar" + "hello") + "hello")) + "a") + "")
    v49 := GenS1{F0: (("a" + "hello") + "hello")}
    for fi50 := 0; fi50 < 20; fi50++ {
        v6 = (-(v6))
        if (0 != 42) {
            fmt.Println(v3, fi50, v1)
            var v51 int32 = 42
            var v52 string = "hello"
            _ = v51
            _ = v52
        } else {
            v1 = v1
            if ((!((130 == 42) || ((0 < 42) || (((((-2.25 - (float32((float64(((v1 - v1) * (v1 + 1e10))) / 1.5)) - 1.5)) * float32((((3.14159 + -2.25) - -2.25) + -2.25))) / float32(float64(v1))) + 0.0) > 1.5)))) && false) {
                for fi53 := 0; fi53 < 9; fi53++ {
                    if (!false) {
                        v6 = v6
                        delete(v2, 65535)
                        v54 := GenS4{F0: -128}
                        _ = v54
                    }
                    if true {
                        if (0 > 18446744073709551615) {
                            switch "hello" {
                            case "a":
                                switch (("foo bar" + "a") + ((("" + "hello") + "a") + "foo bar")) {
                                case "":
                                    var v55 uint32 = 2147483648
                                    _ = v55
                                default:
                                    fmt.Println(fi50)
                                }
                            default:
                                fmt.Println(v6)
                            }
                        } else {
                            switch -449 {
                            case -545:
                                var v56 int64 = int64(fi50)
                                _ = v56
                            default:
                                if (false && false) {
                                    v57 := []int{2147483647}
                                    if true {
                                        var v58 int32 = -814
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v59 float32 = v1
                                            _ = v59
                                        default:
                                            switch ("foo bar" + (("" + "hello") + "\n\t")) {
                                            case "foo bar":
                                                var v60 float64 = 1.5
                                                _ = v60
                                            default:
                                                fmt.Println(v58, v1, fi50)
                                            }
                                            var v61 uint16 = 1
                                            _ = v61
                                        }
                                        _ = v58
                                    }
                                    v62 := GenS5{F0: genFn0(626, false, fi53), F1: v6}
                                    _ = v57
                                    _ = v62
                                } else {
                                    switch ((("" + "") + "hello") + ("a" + "a")) {
                                    case "":
                                        v1 = v1
                                    case "\n\t":
                                        fmt.Println(v3, v6)
                                    default:
                                        v6 = (v6 + 42)
                                    }
                                    v3 += (v3 + (v3 % (((v3 * (v3 >> ((425) & 63)))) | 1)))
                                    if true {
                                        switch "a" {
                                        case "foo bar":
                                            var v63 int32 = -2147483648
                                            _ = v63
                                        default:
                                            switch "" {
                                            case "":
                                                if false {
                                                    if true {
                                                        var v64 int8 = 100
                                                        var v65 uint16 = 1
                                                        if false {
                                                            fmt.Println(v64, v1)
                                                        }
                                                        v66 := GenS6{F0: -32768, F1: 42, F2: -1}
                                                        _ = v64
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                    fmt.Println(fi53, v1)
                                                    var v67 int32 = -2147483648
                                                    _ = v67
                                                }
                                            case "hello":
                                                var v68 float64 = ((-2.25 - (1e10 + 1.5)) + 1e10)
                                                _ = v68
                                            }
                                        }
                                        switch "foo bar" {
                                        case "hello":
                                            v3 += (v3 & v3)
                                            if true {
                                                if (-1 <= 1) {
                                                    var v69 string = "\n\t"
                                                    var v70 int32 = 1
                                                    _ = v69
                                                    _ = v70
                                                }
                                                v71 := []int{(fi50 * 128), fi50, -2147483648}
                                                if len(v71) > 0 {
                                                    v71[0] = (-(fi50))
                                                }
                                                _ = v71
                                            }
                                            var v72 int16 = 42
                                            _ = v72
                                        }
                                        v6 += (v6 * (v6 * -32768))
                                    } else {
                                        v6 += v6
                                        if true {
                                            if true {
                                                var v73 float32 = v1
                                                var v74 uint = 159
                                                _ = v73
                                                _ = v74
                                            } else {
                                                if (!(818 >= 65535)) {
                                                    v1 += v1
                                                    v75 := []int{fi53, fi53, (fi53 - 32767)}
                                                    _ = v75
                                                } else {
                                                    switch 32767 {
                                                    case -2147483648:
                                                        switch fi50 {
                                                        case 4294967295:
                                                            v76 := GenS7{F0: 42, F1: 644}
                                                            var v77 uint16 = 0
                                                            _ = v76
                                                            _ = v77
                                                        }
                                                    case 4294967295:
                                                        switch "foo bar" {
                                                        case "hello":
                                                            if (!(!(true && false))) {
                                                                v78 := GenS3{F0: (v3 - 48)}
                                                                if true {
                                                                    v79 := GenS6{F0: -32768, F1: 255, F2: 445}
                                                                    _ = v79
                                                                } else {
                                                                    var v80 uint32 = 856
                                                                    _ = v80
                                                                }
                                                                var v81 int16 = 1000
                                                                _ = v78
                                                                _ = v81
                                                            }
                                                        default:
                                                            v3 += 42
                                                            var v82 int = 1
                                                            _ = v82
                                                        }
                                                    case -516:
                                                        v1 -= -2.25
                                                    }
                                                    v4[("\n\t" + "a")] = 65535
                                                }
                                            }
                                        } else {
                                            v6 += -32768
                                            var v83 bool = (324 < 18446744073709551615)
                                            var v84 float32 = (float32(float64((((3.14159 * v1) / v1) + (-2.25 - float32(0.0))))) * (-2.25 / (v1 * float32(-2.25))))
                                            _ = v83
                                            _ = v84
                                        }
                                        v6 += (v6 >> ((255) & 63))
                                    }
                                    var v85 uint64 = 2
                                    _ = v85
                                }
                            }
                            var v86 float64 = float64((v1 - 3.14159))
                            if true {
                                fmt.Println(fi53, v1, v1)
                                v86 += (v86 * float64((v1 + 0.0)))
                            } else {
                                v87 := []int{}
                                fmt.Println(v86, fi53)
                                v3 -= (v3 + (v3 * v3))
                                _ = v87
                            }
                            if false {
                                var v88 string = "foo bar"
                                v89 := []int{}
                                v1 -= v1
                                _ = v88
                                _ = v89
                            } else {
                                v90 := GenS7{F0: 255, F1: 201}
                                switch fi50 {
                                case 0:
                                    if (1 == 430) {
                                        v1 -= 3.14159
                                        var v91 uint32 = 4294967295
                                        _ = v91
                                    } else {
                                        var v92 int8 = 4
                                        _ = v92
                                    }
                                case 128:
                                    switch (fi50 * 2147483647) {
                                    case 584:
                                        fmt.Println(v1, fi50)
                                    case -1:
                                        if false {
                                            delete(v5, (fi53 >> ((2) & 63)))
                                        }
                                        v6 = v6
                                        switch "foo bar" {
                                        case "\n\t":
                                            v86 += 0.0
                                        case "a":
                                            var v93 float64 = ((1.5 * 0.0) + float64(v1))
                                            _ = v93
                                        case "":
                                            switch (^(fi53)) {
                                            case 42:
                                                if false {
                                                    v1 -= 0.0
                                                    v94 := GenS5{F0: 885, F1: (^(v6))}
                                                    switch ("a" + "") {
                                                    case "\n\t":
                                                        var v95 string = "a"
                                                        _ = v95
                                                    case "hello":
                                                        if (32767 == fi50) {
                                                            var v96 int = (fi50 >> ((42) & 63))
                                                            var v97 uint32 = 1
                                                            v2[v97] = ("" + "\n\t")
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                    }
                                                    _ = v94
                                                } else {
                                                    v98 := GenS5{F0: genFn0(255, false, fi50), F1: (v6 % ((v6) | 1))}
                                                    fmt.Println(v1)
                                                    _ = v98
                                                }
                                            case 1:
                                                var v99 int32 = 2147483647
                                                _ = v99
                                            case 127:
                                                v6 -= 42
                                            default:
                                                var v100 uint = 255
                                                _ = v100
                                            }
                                            fmt.Println(fi53, v1, fi53)
                                        default:
                                            v6 += 1
                                        }
                                    case 2147483647:
                                        v101 := GenS1{F0: ("a" + "a")}
                                        _ = v101
                                    default:
                                        switch "\n\t" {
                                        case "foo bar":
                                            if ((true || false) || false) {
                                                if (9223372036854775807 <= 9223372036854775807) {
                                                    v102 := GenS8{F0: ((("hello" + "hello") + "foo bar") + "foo bar"), F1: int8(v6), F2: (v6 << ((881) & 63))}
                                                    fmt.Println(fi53, fi50, v6)
                                                    delete(v5, (fi50 * (fi53 >> ((422) & 63))))
                                                    _ = v102
                                                }
                                                switch (("a" + "foo bar") + ("hello" + "\n\t")) {
                                                case "hello":
                                                    v103 := GenS8{F0: ("a" + (("a" + (("a" + "a") + "a")) + (("\n\t" + (("hello" + "hello") + ("foo bar" + "hello"))) + ("foo bar" + "foo bar")))), F1: -128, F2: v6}
                                                    _ = v103
                                                case "a":
                                                    switch (-(fi50)) {
                                                    case 2147483647:
                                                        var v104 float32 = v1
                                                        _ = v104
                                                    }
                                                case "foo bar":
                                                    fmt.Println(fi53, v1, v86)
                                                }
                                            }
                                        case "\n\t":
                                            v105 := []int{128, (fi53 * fi50), (fi50 * fi53)}
                                            _ = v105
                                        case "a":
                                            switch ("foo bar" + "\n\t") {
                                            case "hello":
                                                var v106 int = fi50
                                                _ = v106
                                            default:
                                                fmt.Println(v1)
                                                fmt.Println(v3)
                                            }
                                        default:
                                            v107 := GenS2{F0: true, F1: 29, F2: -1}
                                            _ = v107
                                        }
                                    }
                                case 4294967295:
                                    var v108 uint8 = 1
                                    _ = v108
                                }
                                v86 = (float64(v1) / ((float64(float32(0.0)) + (v86 / ((((float64(v1) + v86) + v86) * v86) * v86))) - v86))
                                v6 -= -789
                                v109 := GenS8{F0: "", F1: v3, F2: v6}
                                _ = v109
                                _ = v90
                            }
                            _ = v86
                        }
                    }
                    var v110 string = "a"
                    v111 := GenS8{F0: v110, F1: (v3 / (((v3 / ((0) | 1))) | 1)), F2: v6}
                    _ = fi53
                    _ = v110
                    _ = v111
                }
                v1 += (float32(1e10) - -2.25)
            } else {
                if (!false) {
                    v112 := []int{(^(fi50)), (fi50 | fi50), (fi50 ^ (fi50 ^ (-(fi50))))}
                    if true {
                        for fi113 := 0; fi113 < 4; fi113++ {
                            v3 += 65
                            _ = fi113
                        }
                        var v114 int32 = -1
                        _ = v114
                    } else {
                        v115 := GenS5{F0: 18446744073709551615, F1: 32767}
                        _ = v115
                    }
                    _ = v112
                }
                var v116 int32 = 0
                _ = v116
            }
            var v117 float64 = float64(float32(0.0))
            _ = v117
        }
        v1 = (v1 * (v1 * (-2.25 * v1)))
        _ = fi50
    }
    if (!(true && ("foo bar" > "hello"))) {
        delete(v4, "\n\t")
        switch ("" + "hello") {
        case "foo bar":
            for fi118 := 0; fi118 < 7; fi118++ {
                if (fi118 == fi118) {
                    switch (("\n\t" + "a") + ("a" + ("a" + "hello"))) {
                    case "\n\t":
                        for fi119 := 0; fi119 < 13; fi119++ {
                            if (0 <= 255) {
                                if (false && (-1 <= 1)) {
                                    v3 += (v3 * (v3 >> ((266) & 63)))
                                }
                                fmt.Println(v1, fi118, v3)
                            }
                            _ = fi119
                        }
                    case "hello":
                        v3 = 42
                    case "foo bar":
                        v120 := GenS5{F0: genFn0(32767, ("hello" <= "\n\t"), fi118), F1: (v6 * v6)}
                        _ = v120
                    default:
                        v121 := GenS6{F0: 32767, F1: 1, F2: 0}
                        _ = v121
                    }
                }
                for fi122 := 0; fi122 < 14; fi122++ {
                    var v123 float64 = 1e10
                    if (!(int16(fi122) == -27)) {
                        v123 = 0.0
                        if false {
                            if ((1 != 0) || (94 > 1)) {
                                if true {
                                    var v124 uint64 = 926
                                    _ = v124
                                } else {
                                    switch fi122 {
                                    case -2147483648:
                                        fmt.Println(fi118, v3, v6)
                                    }
                                    if false {
                                        v125 := GenS6{F0: v6, F1: 176, F2: 1}
                                        v1 = (v1 - ((0.1 * (0.1 * 1.5)) - v1))
                                        _ = v125
                                    }
                                    v126 := []int{fi122}
                                    _ = v126
                                }
                                if ((((("\n\t" + "\n\t") + "") + "a") == ("\n\t" + "hello")) && false) {
                                    v127 := []int{4294967295}
                                    _ = v127
                                } else {
                                    if (v3 <= 0) {
                                        v6 -= v6
                                        switch fi118 {
                                        case -195:
                                            v123 = ((v123 / 0.1) * v123)
                                        default:
                                            v123 = 1.5
                                        }
                                    }
                                    var v128 int16 = 1
                                    var v129 float64 = 1.5
                                    v129 = (float64(float32((((v129 * v129) * float64(-2.25)) - 1.5))) + float64((v1 - v1)))
                                    _ = v128
                                    _ = v129
                                }
                                v3 = v3
                                var v130 uint8 = 107
                                _ = v130
                            } else {
                                var v131 uint8 = 1
                                _ = v131
                            }
                            v132 := GenS1{F0: "\n\t"}
                            _ = v132
                        }
                        var v133 uint = 65535
                        _ = v133
                    }
                    v134 := GenS3{F0: 42}
                    var v135 float32 = v1
                    v135 = 1e10
                    v3 = -1
                    if ("a" == "a") {
                        if (128 == 142) {
                            v135 += (v135 * 1.5)
                            if (65535 > 65535) {
                                v135 -= float32((float64(float32(v123)) / (v123 * v123)))
                                var v136 int32 = 0
                                _ = v136
                            }
                            v137 := []int{fi118, (fi118 ^ fi118), (fi122 >> ((228) & 63)), (-(fi122))}
                            _ = v137
                        } else {
                            v138 := []int{(-(fi118)), fi118}
                            _ = v138
                        }
                    } else {
                        fmt.Println(v135)
                        var v139 uint64 = 2
                        fmt.Println(v6)
                        _ = v139
                    }
                    _ = fi122
                    _ = v123
                    _ = v134
                    _ = v135
                }
                for fi140 := 0; fi140 < 20; fi140++ {
                    if (1.5 != (3.14159 + 0.1)) {
                        v5[128] = "hello"
                        var v141 float32 = 0.1
                        _ = v141
                    }
                    v3 = (v3 - -128)
                    _ = fi140
                }
                _ = fi118
            }
        case "":
            v142 := GenS1{F0: ""}
            _ = v142
        }
        for fi143 := 0; fi143 < 5; fi143++ {
            var v144 int32 = 0
            switch (fi143 | (^(fi143))) {
            case -630:
                fmt.Println(v3, v3, v6)
            case -1:
                v145 := GenS8{F0: (("a" + "hello") + "a"), F1: (-(v3)), F2: (v6 & 0)}
                _ = v145
            }
            _ = fi143
            _ = v144
        }
    }
    for fi146 := 0; fi146 < 8; fi146++ {
        if (!(!(-769 <= (-(v6))))) {
            var v147 float64 = -2.25
            if false {
                v3 += 127
            }
            v1 = float32(0.0)
            _ = v147
        }
        if (164 == 255) {
            for fi148 := 0; fi148 < 13; fi148++ {
                switch fi148 {
                case 4294967295:
                    v149 := GenS4{F0: v3}
                    _ = v149
                case -1:
                    switch "hello" {
                    case "foo bar":
                        v3 -= v3
                    case "a":
                        v150 := GenS1{F0: "\n\t"}
                        _ = v150
                    case "\n\t":
                        var v151 float32 = -2.25
                        if false {
                            v6 -= (^(v6))
                            var v152 uint16 = 42
                            _ = v152
                        } else {
                            switch ("a" + "a") {
                            case "foo bar":
                                v153 := []int{fi148}
                                _ = v153
                            }
                            var v154 float32 = v1
                            _ = v154
                        }
                        _ = v151
                    default:
                        switch (fi148 ^ -133) {
                        case -2147483648:
                            switch ("hello" + "a") {
                            case "foo bar":
                                v155 := GenS3{F0: 1}
                                _ = v155
                            case "hello":
                                delete(v4, "")
                                switch (fi148 * fi148) {
                                case -1:
                                    v156 := []int{fi146}
                                    _ = v156
                                case 418:
                                    v157 := GenS6{F0: 0, F1: 33, F2: 9223372036854775807}
                                    v3 = 127
                                    var v158 uint = 1
                                    _ = v157
                                    _ = v158
                                case 1:
                                    v3 += -128
                                }
                                if (!(!false)) {
                                    var v159 float32 = v1
                                    _ = v159
                                } else {
                                    if ((42 <= -600) && false) {
                                        if (42 >= 878) {
                                            if (((fi146 | 962) != -758) && (2 == 42)) {
                                                var v160 uint = 435
                                                v161 := GenS9{F0: -128, F1: (v160 / (((v160 & 506)) | 1)), F2: (^(v160))}
                                                v3 = 42
                                                _ = v160
                                                _ = v161
                                            } else {
                                                var v162 int = (fi146 ^ fi148)
                                                v163 := GenS10{F0: -833, F1: 785, F2: (-249 != 1)}
                                                var v164 int64 = 1
                                                var v165 int32 = 2147483647
                                                if (18446744073709551615 > 539) {
                                                    v5[v162] = "a"
                                                    v166 := GenS2{F0: true, F1: 0, F2: (v164 ^ int64(v165))}
                                                    fmt.Println(v3, v162)
                                                    _ = v166
                                                } else {
                                                    v162 = (fi148 >> ((745) & 63))
                                                    var v167 float64 = float64(v1)
                                                    v3 = v3
                                                    var v168 int = (fi148 >> ((0) & 63))
                                                    _ = v167
                                                    _ = v168
                                                }
                                                if (v165 > (^(v165))) {
                                                    if false {
                                                        if true {
                                                            var v169 float32 = v1
                                                            v170 := GenS11{F0: (v164 >> ((490) & 63)), F1: 799, F2: 128}
                                                            v171 := []int{42, (fi146 ^ fi146)}
                                                            _ = v169
                                                            _ = v170
                                                            _ = v171
                                                        } else {
                                                            v172 := GenS9{F0: 0, F1: 2, F2: 2}
                                                            var v173 uint32 = 573
                                                            _ = v172
                                                            _ = v173
                                                        }
                                                    } else {
                                                        v165 += v165
                                                        v1 = v1
                                                        v1 = ((((v1 / -2.25) - v1) / (1.5 - float32((1.5 / 3.14159)))) + v1)
                                                    }
                                                    var v174 uint8 = 42
                                                    _ = v174
                                                } else {
                                                    v175 := GenS12{F0: 239}
                                                    _ = v175
                                                }
                                                _ = v162
                                                _ = v163
                                                _ = v164
                                                _ = v165
                                            }
                                            v1 = ((v1 * float32(1.5)) + (v1 - 0.1))
                                        } else {
                                            var v176 int = (fi146 | fi148)
                                            _ = v176
                                        }
                                        v177 := []int{}
                                        _ = v177
                                    }
                                    v5[(fi148 | (fi146 + fi148))] = "foo bar"
                                    v178 := GenS8{F0: "a", F1: v3, F2: v6}
                                    _ = v178
                                }
                            case "\n\t":
                                if (!false) {
                                    var v179 string = ""
                                    _ = v179
                                } else {
                                    var v180 int32 = 0
                                    fmt.Println(v180, fi148)
                                    if (v1 <= float32(float64(v1))) {
                                        if (false && true) {
                                            var v181 int32 = (v180 * (v180 << ((4294967295) & 63)))
                                            v181 = (v181 ^ 1)
                                            delete(v5, (^(fi146)))
                                            var v182 int16 = int16(v180)
                                            var v183 uint = 0
                                            _ = v181
                                            _ = v182
                                            _ = v183
                                        }
                                        fmt.Println(v3, v3, v180)
                                    }
                                    _ = v180
                                }
                            default:
                                switch (fi146 | 176) {
                                case 42:
                                    var v184 bool = true
                                    _ = v184
                                case -758:
                                    var v185 int8 = (^(v3))
                                    _ = v185
                                case 2147483647:
                                    v186 := []int{}
                                    _ = v186
                                }
                            }
                        default:
                            var v187 uint8 = 64
                            _ = v187
                        }
                    }
                default:
                    v3 -= v3
                }
                switch fi146 {
                case 4294967295:
                    fmt.Println(fi146, v6, fi148)
                case -2147483648:
                    var v188 uint64 = 848
                    _ = v188
                }
                _ = fi148
            }
            fmt.Println(v6, v6)
            v1 += -2.25
        } else {
            for fi189 := 0; fi189 < 20; fi189++ {
                v190 := GenS10{F0: 1, F1: 920, F2: (false || true)}
                _ = fi189
                _ = v190
            }
            switch (("hello" + "hello") + "hello") {
            case "\n\t":
                var v191 uint = 18446744073709551615
                _ = v191
            default:
                v6 = 32767
            }
        }
        _ = fi146
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v49
    _ = v5
    _ = v6
}
