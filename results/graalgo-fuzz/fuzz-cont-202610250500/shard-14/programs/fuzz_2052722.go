package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int16; F1 string; F2 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint; F2 string }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint8; F1 int8 }
type GenS9 struct { F0 int8; F1 uint8; F2 string }
func genFn0(p0 uint16) int8 {
    p0 -= p0
    var v6 int8 = 13
    var v7 uint16 = p0
    _ = p0
    _ = v6
    _ = v7
    return (v6 + (^(v6)))
}

func main() {
    var v1 uint64 = 148
    var v2 string = "\n\t"
    var v3 float64 = 0.1
    v4 := make(map[string]bool)
    v5 := make(map[uint32]bool)
    for fi8 := 0; fi8 < 5; fi8++ {
        if false {
            v1 += (v1 | (v1 | v1))
            if false {
                v9 := GenS1{F0: v2, F1: 632, F2: 192}
                var v10 int = (fi8 >> ((891) & 63))
                _ = v10
                _ = v9
            } else {
                v11 := v4[v2]
                var v12 string = "\n\t"
                switch (v2 + (((((v2 + (v2 + ((v2 + v12) + v2))) + ("a" + "foo bar")) + v12) + v12) + (((v12 + "a") + (v12 + (v2 + (v12 + "a")))) + ("foo bar" + "foo bar")))) {
                case "\n\t":
                    v13 := GenS1{F0: (((v12 + ("hello" + ((("foo bar" + "a") + "hello") + v2))) + v12) + ("a" + "")), F1: 684, F2: 4294967295}
                    _ = v13
                case "a":
                    switch (v2 + v12) {
                    case "a":
                        if ((fi8 | -2147483648) > 4294967295) {
                            v14 := []int{fi8, (fi8 & fi8)}
                            _ = v14
                        }
                    case "":
                        _ = len(v4)
                    default:
                        if true {
                            var v15 uint = 65535
                            _ = v15
                        } else {
                            v16 := v4[(v2 + v2)]
                            for fi17 := 0; fi17 < 15; fi17++ {
                                v18 := GenS1{F0: "foo bar", F1: 27, F2: 0}
                                v19 := GenS1{F0: ((((v2 + ("a" + v2)) + "") + ((v12 + v2) + "")) + ""), F1: 0, F2: 0}
                                var v20 int64 = 9223372036854775807
                                _ = fi17
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                            _ = v16
                        }
                    }
                }
                _ = v11
                _ = v12
            }
        } else {
            v21 := v5[65535]
            var v22 int8 = 100
            for fi23 := 0; fi23 < 16; fi23++ {
                v21 = true
                _ = fi23
            }
            _ = v21
            _ = v22
        }
        for fi24 := 0; fi24 < 10; fi24++ {
            fmt.Println(v1)
            _ = fi24
        }
        switch 32767 {
        case 831:
            var v25 int32 = 618
            _ = v25
        default:
            for fi26 := 0; fi26 < 13; fi26++ {
                v27 := GenS2{F0: -408, F1: (v2 + v2), F2: 255}
                _ = fi26
                _ = v27
            }
        }
        _ = fi8
    }
    _ = len(v5)
    var v28 uint32 = 259
    var v29 bool = true
    fmt.Println(v29, v1, v3)
    v28 = (v28 ^ v28)
    for fi30 := 0; fi30 < 6; fi30++ {
        var v31 int = fi30
        v32 := GenS3{F0: v2}
        if ((v29 || ((65535 > v28) || (1 == 100))) || v29) {
            var v33 int16 = -576
            _ = v33
        }
        _ = fi30
        _ = v31
        _ = v32
    }
    var v34 uint64 = 2
    var v35 int32 = 42
    for fi36 := 0; fi36 < 5; fi36++ {
        for fi37 := 0; fi37 < 7; fi37++ {
            switch (("a" + "") + "") {
            case "":
                var v38 float32 = (float32(v3) / (1e10 + 3.14159))
                _ = v38
            case "\n\t":
                if (v29 || ((-2.25 / float32(v3)) < 1e10)) {
                    if v29 {
                        switch (("a" + v2) + "") {
                        case "a":
                            if (!v29) {
                                delete(v5, (^(v28)))
                                var v39 float64 = float64(0.1)
                                var v40 int8 = -107
                                _ = v39
                                _ = v40
                            } else {
                                v41 := GenS2{F0: -32768, F1: ((("\n\t" + "") + (v2 + (((v2 + (v2 + ((v2 + v2) + v2))) + v2) + v2))) + v2), F2: 255}
                                v35 += -2147483648
                                v42 := GenS2{F0: 42, F1: v2, F2: 255}
                                _ = v41
                                _ = v42
                            }
                        case "foo bar":
                            if v29 {
                                v43 := GenS2{F0: -1, F1: v2, F2: 270}
                                _ = v43
                            }
                        case "":
                            v44 := GenS4{F0: fi37, F1: 42, F2: "\n\t"}
                            _ = v44
                        }
                    }
                    v3 += (v3 * v3)
                    switch 635 {
                    case 2147483647:
                        if (fi37 > fi37) {
                            fmt.Println(fi37, fi36)
                        }
                    case 42:
                        fmt.Println(v3)
                    default:
                        if (v35 > v35) {
                            switch (fi37 % ((int(v34)) | 1)) {
                            case -2147483648:
                                if false {
                                    v34 += 0
                                    switch (fi36 - 4294967295) {
                                    case 9223372036854775807:
                                        if true {
                                            if ((v3 - v3) == float64(0.0)) {
                                                fmt.Println(v1)
                                                fmt.Println(v3, v2, v35)
                                            }
                                            v5[v28] = true
                                            v45 := GenS2{F0: 1, F1: ("a" + "a"), F2: 1}
                                            switch (fi37 << ((1) & 63)) {
                                            case 2147483647:
                                                switch (^(fi37)) {
                                                case -2147483648:
                                                    if (!((v34 & 1) < (v1 + v1))) {
                                                        switch fi37 {
                                                        case -49:
                                                            v46 := GenS3{F0: "\n\t"}
                                                            _ = v46
                                                        case -1:
                                                            _ = len(v5)
                                                        case -2147483648:
                                                            var v47 int = (fi37 * fi37)
                                                            var v48 float64 = -2.25
                                                            v49 := GenS2{F0: -21, F1: ("" + v2), F2: 311}
                                                            _ = v47
                                                            _ = v48
                                                            _ = v49
                                                        default:
                                                            v34 -= (^(v34))
                                                        }
                                                        v50 := GenS2{F0: -1, F1: v2, F2: 0}
                                                        _ = v50
                                                    } else {
                                                        switch (v2 + v2) {
                                                        case "hello":
                                                            fmt.Println(v34)
                                                        case "\n\t":
                                                            v2 = v2
                                                        case "a":
                                                            v35 -= v35
                                                        default:
                                                            var v51 string = "hello"
                                                            _ = v51
                                                        }
                                                        if true {
                                                            v1 = uint64(v28)
                                                            v2 = "hello"
                                                            switch (v2 + ((v2 + v2) + v2)) {
                                                            case "\n\t":
                                                                if (!v29) {
                                                                    v52 := []int{(fi36 % ((fi37) | 1)), fi37, fi36}
                                                                    fmt.Println(fi37, v28, v35)
                                                                    _ = v52
                                                                }
                                                            case "a":
                                                                var v53 uint8 = 127
                                                                _ = v53
                                                            default:
                                                                var v54 float32 = (1e10 - 3.14159)
                                                                _ = v54
                                                            }
                                                        }
                                                        delete(v5, v28)
                                                        var v55 int32 = (v35 & v35)
                                                        _ = v55
                                                    }
                                                default:
                                                    var v56 int16 = 1
                                                    _ = v56
                                                }
                                            case -370:
                                                var v57 string = (v2 + v2)
                                                _ = v57
                                            }
                                            _ = v45
                                        } else {
                                            if false {
                                                var v58 uint8 = 1
                                                _ = v58
                                            } else {
                                                fmt.Println(v1)
                                            }
                                            var v59 uint16 = 65535
                                            _ = v59
                                        }
                                    default:
                                        v60 := []int{fi36, -9223372036854775808}
                                        _ = v60
                                    }
                                    fmt.Println(v35, v28, v34)
                                } else {
                                    if (true && (!true)) {
                                        v35 = (v35 >> ((26) & 63))
                                        var v61 bool = true
                                        if (0 >= 42) {
                                            switch v2 {
                                            case "a":
                                                var v62 bool = v29
                                                _ = v62
                                            default:
                                                var v63 bool = true
                                                _ = v63
                                            }
                                            v2 = (v2 + "")
                                            v1 = 342
                                        }
                                        var v64 uint8 = 1
                                        _ = v61
                                        _ = v64
                                    } else {
                                        fmt.Println(v35)
                                    }
                                    fmt.Println(v3, fi36)
                                }
                            case 32767:
                                switch (^(fi36)) {
                                case 127:
                                    var v65 int16 = 1000
                                    _ = v65
                                case -200:
                                    delete(v5, (v28 / ((2147483648) | 1)))
                                case 9223372036854775807:
                                    if (v29 || (v2 >= (v2 + (v2 + "")))) {
                                        if (((false && (!v29)) && v29) && (v29 && true)) {
                                            v66 := GenS3{F0: v2}
                                            var v67 uint = 2
                                            _ = v66
                                            _ = v67
                                        } else {
                                            switch "\n\t" {
                                            case "":
                                                v3 = v3
                                            default:
                                                fmt.Println(v29, fi36)
                                            }
                                            v68 := []int{(-(fi36)), fi37, -2147483648}
                                            _ = v68
                                        }
                                        var v69 int = (fi37 >> ((4294967295) & 63))
                                        var v70 bool = true
                                        _ = v69
                                        _ = v70
                                    }
                                default:
                                    var v71 int16 = -32768
                                    _ = v71
                                }
                            }
                            v72 := []int{(fi36 & (fi36 & fi36)), (fi36 * (-(fi37)))}
                            _ = v72
                        } else {
                            v73 := []int{305, fi37, (fi36 ^ fi37)}
                            var v74 uint64 = (v1 / ((555) | 1))
                            v2 = (("\n\t" + (v2 + (("a" + ("\n\t" + ((v2 + v2) + v2))) + v2))) + v2)
                            v1 = 4294967295
                            _ = v73
                            _ = v74
                        }
                    }
                    v34 = (v34 / ((v1) | 1))
                    var v75 uint32 = (v28 * v28)
                    switch ((v2 + (((v2 + "hello") + "a") + (("a" + "\n\t") + ("a" + (v2 + v2))))) + (v2 + ((v2 + v2) + v2))) {
                    case "\n\t":
                        var v76 bool = false
                        _ = v76
                    case "hello":
                        fmt.Println(v1)
                    case "foo bar":
                        if (42 != 255) {
                            if (v35 < v35) {
                                fmt.Println(v2)
                                switch ("" + (v2 + (v2 + v2))) {
                                case "a":
                                    v1 -= v34
                                    v77 := GenS1{F0: v2, F1: 32767, F2: v28}
                                    _ = v77
                                case "foo bar":
                                    v78 := []int{fi37, fi37, fi36}
                                    _ = v78
                                default:
                                    switch -858 {
                                    case -652:
                                        v35 = -2147483648
                                    case 852:
                                        fmt.Println(v75, v29)
                                    }
                                }
                                if v29 {
                                    var v79 uint32 = 42
                                    v35 = (v35 / (((v35 - v35)) | 1))
                                    _ = v79
                                } else {
                                    var v80 int = (-(fi37))
                                    v29 = ((v35 << ((233) & 63)) != v35)
                                    var v81 uint64 = (v1 << ((42) & 63))
                                    var v82 uint64 = (v81 * 4294967295)
                                    v4[("a" + (v2 + (v2 + v2)))] = v29
                                    switch v2 {
                                    case "\n\t":
                                        var v83 string = ("" + v2)
                                        v75 += v75
                                        if true {
                                            fmt.Println(v28)
                                            var v84 uint64 = (v1 - v81)
                                            _ = v84
                                        } else {
                                            if ((fi36 % (((v80 & v80)) | 1)) >= 4294967295) {
                                                v2 = v83
                                                v85 := GenS4{F0: fi36, F1: 4294967295, F2: ""}
                                                _ = v85
                                            }
                                            _ = len(v4)
                                        }
                                        _ = v83
                                    default:
                                        v86 := GenS5{F0: 32767, F1: (!v29)}
                                        _ = v86
                                    }
                                    _ = v80
                                    _ = v81
                                    _ = v82
                                }
                                v87 := GenS6{F0: (v2 + "a"), F1: 169}
                                _ = v87
                            } else {
                                v88 := []int{(fi37 * (-(fi36)))}
                                v88 = append(v88, (fi36 - (fi37 / ((-558) | 1))))
                                _ = v88
                            }
                            _ = len(v4)
                            if ((v35 << ((4294967295) & 63)) != 297) {
                                v75 = (v28 | v75)
                                if ((^(v75)) >= (v75 % ((937) | 1))) {
                                    var v89 float32 = (1e10 + (-2.25 * 1.5))
                                    if v29 {
                                        v35 = (v35 >> ((65535) & 63))
                                        switch v2 {
                                        case "\n\t":
                                            var v90 uint8 = 0
                                            _ = v90
                                        case "hello":
                                            var v91 uint = 743
                                            _ = v91
                                        case "a":
                                            var v92 string = (v2 + v2)
                                            var v93 int32 = v35
                                            fmt.Println(v35, v29)
                                            _ = v92
                                            _ = v93
                                        default:
                                            if (!true) {
                                                switch v2 {
                                                case "a":
                                                    var v94 string = "hello"
                                                    _ = v94
                                                case "hello":
                                                    switch fi36 {
                                                    case 1:
                                                        v95 := v4[(v2 + v2)]
                                                        _ = v95
                                                    case 4294967295:
                                                        fmt.Println(v3, v3)
                                                    default:
                                                        v4[(v2 + ("foo bar" + "a"))] = v29
                                                    }
                                                }
                                                v96 := []int{fi36, 9223372036854775807, (fi36 - fi37), (^(fi37))}
                                                v1 = v34
                                                _ = v96
                                            }
                                        }
                                        v97 := GenS3{F0: "a"}
                                        _ = v97
                                    }
                                    var v98 uint64 = 328
                                    _ = v89
                                    _ = v98
                                } else {
                                    v2 = (v2 + ("a" + (v2 + ("a" + ""))))
                                }
                                v3 += ((float64(0.0) + 0.0) + 3.14159)
                            } else {
                                var v99 float64 = v3
                                fmt.Println(v3)
                                _ = v99
                            }
                            var v100 uint16 = 2
                            v29 = false
                            _ = v100
                        } else {
                            v75 = (v75 + v75)
                            v101 := GenS1{F0: v2, F1: 255, F2: v75}
                            v3 += v3
                            delete(v5, v75)
                            _ = v101
                        }
                    }
                    _ = v75
                } else {
                    v102 := []int{(fi37 | fi36), (fi36 ^ fi36), fi37, (-(fi37))}
                    v102 = append(v102, 713)
                    _ = v102
                }
            }
            var v103 int64 = 9223372036854775807
            _ = fi37
            _ = v103
        }
        v28 = v28
        v104 := []int{}
        _ = fi36
        _ = v104
    }
    for fi105 := 0; fi105 < 7; fi105++ {
        var v106 int16 = 42
        switch (v2 + (v2 + "a")) {
        case "\n\t":
            for fi107 := 0; fi107 < 19; fi107++ {
                var v108 uint8 = 128
                v109 := []int{fi105, fi107, -166, (fi107 % ((127) | 1))}
                v2 = v2
                _ = fi107
                _ = v108
                _ = v109
            }
        }
        var v110 uint8 = 138
        _ = fi105
        _ = v106
        _ = v110
    }
    fmt.Println(v28, v2, v35)
    v3 = v3
    delete(v5, v28)
    v111 := GenS5{F0: 324, F1: (v29 || (!(0 < 45)))}
    switch v2 {
    case "hello":
        for fi112 := 0; fi112 < 8; fi112++ {
            for fi113 := 0; fi113 < 12; fi113++ {
                v34 = 807
                if v29 {
                    v28 = uint32(v1)
                } else {
                    if (v29 || ((fi112 & fi113) <= (fi113 * (-(fi113))))) {
                        v114 := GenS1{F0: v2, F1: 42, F2: 913}
                        var v115 int8 = 0
                        _ = v114
                        _ = v115
                    }
                    if v29 {
                        v116 := GenS5{F0: -1, F1: true}
                        var v117 uint8 = 127
                        if (true || v29) {
                            var v118 uint32 = v28
                            fmt.Println(fi112, v34, v117)
                            if ((-9223372036854775808 != 0) || (!(!v29))) {
                                v119 := GenS5{F0: 5, F1: (v29 || true)}
                                switch v2 {
                                case "":
                                    if (-32768 > -32768) {
                                        v120 := GenS6{F0: v2, F1: 18446744073709551615}
                                        if true {
                                            var v121 uint32 = v118
                                            v117 = 255
                                            fmt.Println(fi112)
                                            _ = v121
                                        } else {
                                            v5[(v28 + (v28 + v28))] = v29
                                        }
                                        v122 := GenS7{F0: -1}
                                        _ = v120
                                        _ = v122
                                    }
                                default:
                                    var v123 int8 = -128
                                    _ = v123
                                }
                                v124 := GenS4{F0: (fi112 % ((-1) | 1)), F1: 0, F2: v2}
                                _ = v119
                                _ = v124
                            } else {
                                v1 = uint64(v118)
                                v29 = v29
                                if ((-(v35)) < (v35 >> ((255) & 63))) {
                                    var v125 uint = 699
                                    _ = v125
                                }
                            }
                            var v126 uint64 = v34
                            var v127 int8 = -77
                            _ = v118
                            _ = v126
                            _ = v127
                        } else {
                            v29 = v29
                            v128 := GenS1{F0: v2, F1: 65535, F2: (v28 << ((539) & 63))}
                            v28 = v28
                            _ = v128
                        }
                        _ = v116
                        _ = v117
                    } else {
                        v129 := GenS8{F0: 2, F1: 100}
                        v1 += 2
                        var v130 float32 = 0.0
                        _ = v129
                        _ = v130
                    }
                    switch (^(fi113)) {
                    case -760:
                        v131 := GenS9{F0: -1, F1: 42, F2: "foo bar"}
                        _ = v131
                    case 9223372036854775807:
                        v34 -= v1
                    default:
                        switch (fi112 % ((-9223372036854775808) | 1)) {
                        case -819:
                            if (-128 == 1) {
                                v132 := []int{-9223372036854775808, fi113, fi113}
                                _ = len(v132)
                                _ = v132
                            }
                        case 2147483647:
                            if (18446744073709551615 != 4294967295) {
                                var v133 int64 = 9223372036854775807
                                _ = v133
                            }
                        default:
                            switch (^(fi112)) {
                            case 127:
                                v134 := GenS9{F0: 1, F1: 127, F2: v2}
                                _ = v134
                            case 4294967295:
                                v34 = (v1 % ((2) | 1))
                            }
                        }
                    }
                }
                v135 := []int{(-(fi113)), 128, fi112, fi113}
                var v136 uint16 = 255
                if (889 < -592) {
                    switch (-(fi113)) {
                    case -359:
                        if ((v136 | v136) <= (v136 >> ((287) & 63))) {
                            var v137 uint = 2
                            switch "hello" {
                            case "\n\t":
                                if ((v28 * 694) < v28) {
                                    var v138 string = v2
                                    fmt.Println(fi112, v3, v137)
                                    _ = v138
                                }
                            }
                            _ = v137
                        } else {
                            _ = len(v135)
                        }
                    case 127:
                        _ = len(v5)
                    case -2147483648:
                        fmt.Println(v34, v1)
                    default:
                        v136 += v136
                    }
                    v2 = ""
                    if true {
                        var v139 uint32 = v28
                        switch v2 {
                        case "a":
                            v140 := GenS5{F0: 1, F1: v29}
                            _ = v140
                        case "\n\t":
                            v2 = (v2 + v2)
                            fmt.Println(v1, v35, v35)
                        }
                        _ = v139
                    } else {
                        switch (fi113 / ((2147483647) | 1)) {
                        case 753:
                            v2 = v2
                        case 32767:
                            v28 -= uint32(v1)
                        case -9223372036854775808:
                            var v141 uint16 = 32767
                            _ = v141
                        }
                    }
                } else {
                    switch 961 {
                    case -901:
                        if len(v135) > 0 {
                            v135[0] = 2147483647
                        }
                    default:
                        var v142 int = 2147483647
                        _ = v142
                    }
                }
                _ = fi113
                _ = v135
                _ = v136
            }
            _ = fi112
        }
    }
    for fi143 := 0; fi143 < 2; fi143++ {
        v144 := []int{}
        v28 = (^(v28))
        _ = fi143
        _ = v144
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v111
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
