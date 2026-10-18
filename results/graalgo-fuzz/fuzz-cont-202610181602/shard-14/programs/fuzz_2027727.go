package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 uint16; F1 uint32 }
type GenS9 struct { F0 string; F1 int32; F2 string }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint; F1 string }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int64; F1 int32 }
type GenS17 struct { F0 int32; F1 uint16 }
func genFn0() int32 {
    if (false || (0 < -94)) {
        var v7 int32 = 42
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint64 = 42
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        var v10 float32 = 3.14159
        _ = v10
    }
    var v11 string = ("foo bar" + "")
    v12 := GenS1{F0: 2, F1: 249, F2: 1}
    fmt.Println(v11)
    v13 := GenS2{F0: "\n\t", F1: -1}
    _ = v11
    _ = v12
    _ = v13
    return 2147483647
}
func genFn1(p0 uint) bool {
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 int64 = 0
        var v16 float32 = 0.0
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println(p0)
    var v17 int = -168
    fmt.Println(v17)
    var v18 int8 = -128
    _ = p0
    _ = v17
    _ = v18
    return (p0 == (^(p0)))
}

func main() {
    var v1 float32 = 0.0
    var v2 float64 = -2.25
    var v3 bool = false
    var v4 int = -975
    var v5 float32 = (0.0 * float32((1.5 * -2.25)))
    var v6 bool = (v3 || true)
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 uint32 = 2147483648
        _ = fi19
        _ = v20
    }
    if (float64(((((float32((1e10 - v2)) - 3.14159) / v1) - 0.0) * (0.1 + v1))) >= v2) {
        if false {
            for fi21 := 0; fi21 < 15; fi21++ {
                v22 := []int{fi21, 997, fi21}
                if (("a" + ("" + "\n\t")) != "") {
                    if ((v3 && (-36 >= -128)) && (v3 && (456 <= 65535))) {
                        v23 := GenS1{F0: 255, F1: 127, F2: 1}
                        var v24 uint8 = 128
                        if (((-1 <= 795) || v3) && v6) {
                            for fi25 := 0; fi25 < 4; fi25++ {
                                v26 := GenS3{F0: 438, F1: 318, F2: (v24 + v24)}
                                v24 = (v24 + 255)
                                _ = fi25
                                _ = v26
                            }
                        } else {
                            switch ("\n\t" + ("foo bar" + ("a" + "hello"))) {
                            case "foo bar":
                                var v27 bool = (false || (v24 > (v24 & v24)))
                                _ = v27
                            case "a":
                                v22 = append(v22, (fi21 - (-(v4))))
                            case "\n\t":
                                v1 = v1
                            }
                        }
                        _ = v23
                        _ = v24
                    } else {
                        var v28 float64 = (-2.25 - 1e10)
                        switch -9223372036854775808 {
                        case 1:
                            v1 = float32(v2)
                        case 32767:
                            if ((v28 * 1e10) != 0.1) {
                                var v29 int16 = 42
                                var v30 int64 = 42
                                v22 = append(v22, 0)
                                _ = v29
                                _ = v30
                            }
                        case -9223372036854775808:
                            for fi31 := 0; fi31 < 9; fi31++ {
                                switch ("foo bar" + "foo bar") {
                                case "hello":
                                    if v6 {
                                        if (127 > 1) {
                                            v6 = true
                                            if v6 {
                                                var v32 bool = (!(!v3))
                                                v3 = false
                                                _ = v32
                                            } else {
                                                _ = len(v22)
                                                switch (-(fi31)) {
                                                case -836:
                                                    _ = len(v22)
                                                case 657:
                                                    v33 := GenS3{F0: 18446744073709551615, F1: 42, F2: 36}
                                                    _ = v33
                                                case 4294967295:
                                                    var v34 int64 = -734
                                                    _ = v34
                                                default:
                                                    fmt.Println(v6, v28)
                                                }
                                                if false {
                                                    if len(v22) > 0 {
                                                        v22[0] = fi31
                                                    }
                                                } else {
                                                    v35 := GenS2{F0: "hello", F1: 483}
                                                    _ = len(v22)
                                                    if (false || v6) {
                                                        v1 = float32(v2)
                                                    }
                                                    _ = v35
                                                }
                                                v28 -= -2.25
                                                if len(v22) > 0 {
                                                    v22[0] = fi21
                                                }
                                                v28 -= (1.5 + v28)
                                                switch (fi31 ^ (fi31 ^ -746)) {
                                                case 952:
                                                    var v36 bool = ((v2 + v28) >= (v2 / (3.14159 / float64((((((((float32(((v2 - v2) - (1.5 * 3.14159))) * v1) / v5) - 0.1) / v1) - float32(3.14159)) + v5) - v5)))))
                                                    _ = v36
                                                case 128:
                                                    var v37 int32 = 176
                                                    _ = v37
                                                default:
                                                    switch (fi21 | fi31) {
                                                    case 32767:
                                                        if (!v6) {
                                                            v38 := GenS2{F0: "\n\t", F1: 32767}
                                                            var v39 uint16 = 737
                                                            v3 = (("" + "a") == ("a" + "hello"))
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                    case 42:
                                                        var v40 int8 = 0
                                                        _ = v40
                                                    }
                                                    if (1 < 65535) {
                                                        v3 = false
                                                        v22 = append(v22, (v4 & v4))
                                                    } else {
                                                        var v41 uint8 = 0
                                                        var v42 int32 = int32(v41)
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    if v3 {
                                                        if false {
                                                            if ("foo bar" < "\n\t") {
                                                                var v43 int16 = -944
                                                                _ = v43
                                                            } else {
                                                                if len(v22) > 0 {
                                                                    v22[0] = 127
                                                                }
                                                                v44 := GenS1{F0: 2, F1: 127, F2: 42}
                                                                _ = v44
                                                            }
                                                            var v45 uint8 = 0
                                                            _ = v45
                                                        }
                                                        if (v3 && (1 <= 42)) {
                                                            if false {
                                                                v1 = 1.5
                                                                v46 := GenS4{F0: 71}
                                                                _ = v46
                                                            }
                                                            var v47 int32 = -2147483648
                                                            _ = v47
                                                        } else {
                                                            v22 = append(v22, (fi31 + (fi31 | (v4 / ((-904) | 1)))))
                                                            if len(v22) >= 2 {
                                                                v22 = v22[1:len(v22)]
                                                            }
                                                        }
                                                        v48 := GenS2{F0: "\n\t", F1: 349}
                                                        _ = v48
                                                    }
                                                }
                                            }
                                            if len(v22) > 0 {
                                                v22[0] = fi21
                                            }
                                            if len(v22) > 0 {
                                                v22[0] = v4
                                            }
                                            var v49 float64 = -2.25
                                            _ = v49
                                        } else {
                                            fmt.Println(v28, v2, v4)
                                            switch fi21 {
                                            case 9223372036854775807:
                                                var v50 int16 = 519
                                                _ = v50
                                            case 127:
                                                v51 := GenS5{F0: 2, F1: 65535, F2: 0}
                                                _ = v51
                                            case -608:
                                                if v3 {
                                                    v3 = v6
                                                }
                                            default:
                                                if true {
                                                    if ((65535 > 1) || ((((3.14159 - v5) * v1) + 1e10) != v1)) {
                                                        v22 = append(v22, (v4 + fi21))
                                                    } else {
                                                        var v52 float64 = 3.14159
                                                        if len(v22) >= 2 {
                                                            v22 = v22[1:len(v22)]
                                                        }
                                                        _ = v52
                                                    }
                                                    fmt.Println(v28, v3, v4)
                                                    v28 = 0.1
                                                } else {
                                                    if len(v22) > 0 {
                                                        v22[0] = (v4 / (((fi31 - 128)) | 1))
                                                    }
                                                    if (!(127 > -1)) {
                                                        var v53 bool = (641 < 2)
                                                        var v54 int16 = 1000
                                                        _ = v53
                                                        _ = v54
                                                    } else {
                                                        v28 -= 1.5
                                                    }
                                                    fmt.Println(v3)
                                                }
                                            }
                                            var v55 uint16 = 32767
                                            _ = v55
                                        }
                                        v3 = ((!true) && (false && ((true && false) || v6)))
                                        v1 += (1e10 - (float32((float64(v5) * -2.25)) - (v5 - ((float32(0.0) / v1) * v1))))
                                    }
                                case "":
                                    v56 := GenS4{F0: 1}
                                    _ = v56
                                case "\n\t":
                                    switch (v4 >> ((255) & 63)) {
                                    case 32767:
                                        _ = len(v22)
                                    case 4294967295:
                                        var v57 int32 = 919
                                        _ = v57
                                    default:
                                        var v58 int64 = 42
                                        var v59 int64 = (v58 * v58)
                                        _ = v58
                                        _ = v59
                                    }
                                }
                                v5 = -2.25
                                _ = fi31
                            }
                        default:
                            for fi60 := 0; fi60 < 6; fi60++ {
                                switch fi60 {
                                case -1:
                                    var v61 int = fi21
                                    _ = v61
                                }
                                fmt.Println(v6, v4, v1)
                                _ = fi60
                            }
                        }
                        fmt.Println(v6, v2, v6)
                        _ = v28
                    }
                    var v62 string = ("a" + "foo bar")
                    _ = v62
                } else {
                    var v63 float64 = v2
                    if ((v3 && (true || (!(1000 == 360)))) || true) {
                        if len(v22) > 0 {
                            v22[0] = (fi21 | (-(fi21)))
                        }
                        var v64 int = (v4 << ((656) & 63))
                        _ = v64
                    }
                    _ = v63
                }
                _ = fi21
                _ = v22
            }
            if v6 {
                v65 := []int{}
                _ = v65
            } else {
                var v66 int8 = 1
                v67 := GenS6{F0: 247}
                for fi68 := 0; fi68 < 6; fi68++ {
                    v69 := GenS7{F0: 65535, F1: (!(-170 <= -629))}
                    if (42 <= 65535) {
                        if (2147483648 > 837) {
                            v70 := GenS2{F0: "foo bar", F1: -1}
                            switch "hello" {
                            case "a":
                                var v71 int8 = -71
                                _ = v71
                            case "hello":
                                v72 := GenS5{F0: 4294967295, F1: 972, F2: 42}
                                _ = v72
                            }
                            for fi73 := 0; fi73 < 9; fi73++ {
                                if v6 {
                                    v74 := []int{}
                                    if true {
                                        fmt.Println(v5)
                                        var v75 uint8 = 128
                                        _ = v75
                                    }
                                    v2 += (-2.25 - float64((0.1 / float32(((3.14159 / v2) - (float64(3.14159) + -2.25))))))
                                    if ("hello" >= "foo bar") {
                                        fmt.Println(v1, fi73, fi73)
                                        fmt.Println(v6, fi68)
                                    }
                                    fmt.Println(v3, v5)
                                    v76 := GenS8{F0: 255, F1: 446}
                                    _ = v74
                                    _ = v76
                                }
                                var v77 uint16 = 65535
                                _ = fi73
                                _ = v77
                            }
                            _ = v70
                        }
                    } else {
                        v1 -= 0.1
                        var v78 int16 = 0
                        switch "a" {
                        case "":
                            if ((-797 > 710) && true) {
                                for fi79 := 0; fi79 < 20; fi79++ {
                                    var v80 int16 = v78
                                    _ = fi79
                                    _ = v80
                                }
                                switch (fi68 << ((539) & 63)) {
                                case -406:
                                    switch "foo bar" {
                                    case "a":
                                        if v6 {
                                            for fi81 := 0; fi81 < 8; fi81++ {
                                                v2 -= 0.1
                                                if v6 {
                                                    v82 := GenS9{F0: "", F1: genFn0(), F2: "a"}
                                                    v83 := GenS10{F0: 42}
                                                    if v3 {
                                                        v2 += v2
                                                        var v84 uint32 = 924
                                                        _ = v84
                                                    }
                                                    _ = v82
                                                    _ = v83
                                                }
                                                var v85 int64 = -754
                                                _ = fi81
                                                _ = v85
                                            }
                                            for fi86 := 0; fi86 < 5; fi86++ {
                                                v4 += 2147483647
                                                _ = fi86
                                            }
                                        }
                                    default:
                                        var v87 uint64 = 1
                                        _ = v87
                                    }
                                case 1:
                                    if true {
                                        var v88 uint64 = 880
                                        v2 = 0.1
                                        if (v66 <= 100) {
                                            switch "hello" {
                                            case "\n\t":
                                                for fi89 := 0; fi89 < 1; fi89++ {
                                                    var v90 uint32 = 2147483648
                                                    var v91 uint16 = 0
                                                    fmt.Println(v91, v66, v90)
                                                    fmt.Println(v90)
                                                    _ = fi89
                                                    _ = v90
                                                    _ = v91
                                                }
                                            case "foo bar":
                                                switch ("foo bar" + "a") {
                                                case "\n\t":
                                                    for fi92 := 0; fi92 < 18; fi92++ {
                                                        switch -9223372036854775808 {
                                                        case -2147483648:
                                                            v88 = (v88 << ((544) & 63))
                                                        case 1:
                                                            v1 = v1
                                                        default:
                                                            var v93 int16 = v78
                                                            _ = v93
                                                        }
                                                        v94 := GenS11{F0: v78}
                                                        _ = fi92
                                                        _ = v94
                                                    }
                                                case "hello":
                                                    if false {
                                                        for fi95 := 0; fi95 < 8; fi95++ {
                                                            v96 := GenS5{F0: 4294967295, F1: 183, F2: 42}
                                                            var v97 int = fi68
                                                            _ = fi95
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                        for fi98 := 0; fi98 < 3; fi98++ {
                                                            var v99 float32 = 1e10
                                                            _ = fi98
                                                            _ = v99
                                                        }
                                                    }
                                                case "a":
                                                    for fi100 := 0; fi100 < 7; fi100++ {
                                                        v101 := GenS8{F0: 1, F1: 2}
                                                        v102 := []int{(fi100 | fi68), 9223372036854775807, fi68, (fi68 ^ fi100)}
                                                        var v103 float32 = (float32(float64((0.0 * v1))) * ((1e10 * v1) * 3.14159))
                                                        var v104 uint16 = 0
                                                        if len(v102) > 0 {
                                                            v102[0] = -999
                                                        }
                                                        if len(v102) >= 2 {
                                                            v102 = v102[1:len(v102)]
                                                        }
                                                        v105 := GenS12{F0: 1}
                                                        _ = fi100
                                                        _ = v101
                                                        _ = v102
                                                        _ = v103
                                                        _ = v104
                                                        _ = v105
                                                    }
                                                default:
                                                    v106 := []int{int(v78), -993, 4294967295}
                                                    _ = v106
                                                }
                                            default:
                                                v107 := []int{(v4 * fi68), (fi68 ^ v4), -651}
                                                _ = v107
                                            }
                                            switch "" {
                                            case "foo bar":
                                                for fi108 := 0; fi108 < 18; fi108++ {
                                                    v1 -= ((1e10 / v5) + (((float32(v2) - v5) + float32(0.0)) - ((1e10 * v1) - v1)))
                                                    v109 := GenS11{F0: (v78 % ((v78) | 1))}
                                                    v2 = v2
                                                    _ = fi108
                                                    _ = v109
                                                }
                                            default:
                                                v110 := []int{fi68, fi68}
                                                _ = v110
                                            }
                                            v3 = ((32767 < 0) && v3)
                                        }
                                        _ = v88
                                    } else {
                                        switch (v4 % ((v4) | 1)) {
                                        case 32767:
                                            var v111 int32 = -629
                                            _ = v111
                                        case -822:
                                            v6 = ("\n\t" <= "hello")
                                        }
                                    }
                                }
                                v78 = -1
                            }
                        default:
                            for fi112 := 0; fi112 < 10; fi112++ {
                                if v3 {
                                    v113 := GenS6{F0: 0}
                                    var v114 int32 = -690
                                    if ((^(v78)) == v78) {
                                        var v115 uint = 42
                                        if v3 {
                                            v116 := []int{}
                                            _ = v116
                                        }
                                        var v117 int = v4
                                        _ = v115
                                        _ = v117
                                    }
                                    _ = v113
                                    _ = v114
                                }
                                _ = fi112
                            }
                        }
                        _ = v78
                    }
                    _ = fi68
                    _ = v69
                }
                _ = v66
                _ = v67
            }
        } else {
            v1 -= 1e10
        }
        if (18446744073709551615 == 501) {
            v118 := GenS13{F0: 18446744073709551615, F1: (("foo bar" + "foo bar") + "a")}
            var v119 uint = 2
            v4 += (v4 & v4)
            _ = v118
            _ = v119
        } else {
            if v6 {
                v120 := GenS1{F0: 2, F1: 127, F2: 42}
                fmt.Println(v2)
                var v121 uint16 = 255
                _ = v120
                _ = v121
            } else {
                for fi122 := 0; fi122 < 4; fi122++ {
                    switch "foo bar" {
                    case "foo bar":
                        v123 := []int{(v4 / (((fi122 * fi122)) | 1)), -995, fi122}
                        _ = v123
                    case "a":
                        var v124 uint8 = 65
                        _ = v124
                    case "":
                        switch "a" {
                        case "foo bar":
                            var v125 int8 = -1
                            _ = v125
                        default:
                            v3 = v6
                        }
                    default:
                        var v126 uint64 = 0
                        _ = v126
                    }
                    _ = fi122
                }
                for fi127 := 0; fi127 < 7; fi127++ {
                    v5 = 3.14159
                    var v128 float32 = v5
                    _ = fi127
                    _ = v128
                }
            }
            v2 -= v2
        }
    }
    v129 := GenS12{F0: 2}
    var v130 uint8 = 255
    v131 := GenS14{F0: int16(v4)}
    var v132 uint8 = 15
    v133 := GenS15{F0: -97}
    v4 = 0
    switch v4 {
    case 128:
        if (!true) {
            var v134 uint64 = 416
            v2 = 0.1
            _ = v134
        } else {
            v132 += 104
            v135 := GenS16{F0: -338, F1: 1}
            fmt.Println(v4, v5)
            _ = v135
        }
    case 235:
        for fi136 := 0; fi136 < 14; fi136++ {
            v137 := GenS10{F0: v132}
            v4 -= -240
            v3 = (-128 != 81)
            _ = fi136
            _ = v137
        }
    default:
        fmt.Println(v4, v132)
    }
    if (-2147483648 == 873) {
        for fi138 := 0; fi138 < 17; fi138++ {
            fmt.Println(v5, fi138)
            _ = fi138
        }
        v139 := []int{v4}
        switch (^(v4)) {
        case 447:
            if (!(32767 != 0)) {
                var v140 int16 = -32768
                _ = len(v139)
                v3 = false
                _ = v140
            } else {
                v3 = v6
            }
        case 138:
            v6 = (v1 >= v5)
        case 4294967295:
            _ = len(v139)
        default:
            fmt.Println(v1, v4, v2)
        }
        _ = v139
    }
    switch ("\n\t" + "\n\t") {
    case "":
        v141 := []int{(v4 / ((32767) | 1)), v4, v4, -2147483648}
        _ = v141
    case "foo bar":
        v130 = (^(v130))
    case "hello":
        var v142 uint8 = v130
        _ = v142
    default:
        switch "a" {
        case "hello":
            for fi143 := 0; fi143 < 6; fi143++ {
                v144 := GenS17{F0: genFn0(), F1: 2}
                var v145 string = "a"
                v1 = (v1 * v5)
                _ = fi143
                _ = v144
                _ = v145
            }
        case "\n\t":
            if v3 {
                switch "\n\t" {
                case "a":
                    if (v5 != float32(float64(float32((0.0 + v2))))) {
                        v146 := []int{(v4 & (v4 - (-(v4)))), -594, (v4 - (v4 << ((0) & 63)))}
                        _ = v146
                    } else {
                        var v147 float64 = 0.0
                        _ = v147
                    }
                }
            } else {
                switch v4 {
                case 32768:
                    v148 := []int{(v4 & v4)}
                    _ = v148
                case -1:
                    var v149 uint8 = v132
                    _ = v149
                case 4294967295:
                    v150 := []int{(v4 << ((18446744073709551615) & 63)), v4, (-(v4))}
                    _ = v150
                }
            }
        case "foo bar":
            if false {
                v132 = (v130 & v130)
            }
        default:
            var v151 int32 = 277
            _ = v151
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v130:", v130)
    fmt.Println("v132:", v132)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v129
    _ = v130
    _ = v131
    _ = v132
    _ = v133
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
