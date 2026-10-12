package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint16; F2 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 uint; F2 int8 }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool; F1 uint64; F2 int }
type GenS9 struct { F0 string; F1 bool; F2 int8 }
type GenS10 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS11 struct { F0 int64; F1 uint16; F2 bool }
type GenS12 struct { F0 uint8; F1 uint }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int8; F1 int64 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 uint32; F1 int }
type GenS18 struct { F0 int }
func genFn0(p0 uint) int {
    if (!false) {
        for fi4 := 0; fi4 < 18; fi4++ {
            var v5 float32 = 0.1
            if (0 >= 9223372036854775807) {
                v6 := GenS1{F0: 1, F1: 1}
                fmt.Println(v5, v5)
                v7 := []int{217, (fi4 / ((-2147483648) | 1))}
                _ = v6
                _ = v7
            } else {
                var v8 uint8 = 127
                v9 := []int{(fi4 - -9223372036854775808), fi4, -9223372036854775808}
                _ = v8
                _ = v9
            }
            p0 += 2
            _ = fi4
            _ = v5
        }
    } else {
        var v10 uint8 = 221
        v11 := []int{127}
        _ = v10
        _ = v11
    }
    fmt.Println(p0)
    if ((p0 >> ((18446744073709551615) & 63)) < (p0 ^ (p0 % ((2) | 1)))) {
        v12 := GenS1{F0: 65535, F1: 670}
        _ = v12
    } else {
        v13 := GenS1{F0: 42, F1: 233}
        v14 := GenS1{F0: 851, F1: 776}
        _ = v13
        _ = v14
    }
    _ = p0
    return 32768
}
func genFn1(p0 int16, p1 int) uint {
    v15 := GenS2{F0: p1}
    for fi16 := 0; fi16 < 12; fi16++ {
        p0 = p0
        v17 := []int{fi16, 32768}
        p1 = fi16
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 uint8 = 1
        v20 := []int{fi18}
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := []int{p1, (p1 + genFn0(42)), 2147483647}
    if true {
        v22 := GenS2{F0: p1}
        if ((((0.0 * float32(1e10)) * 3.14159) / (1e10 - (1.5 - 1e10))) == (0.1 + float32(0.0))) {
            var v23 uint = 255
            var v24 uint64 = 18446744073709551615
            _ = v23
            _ = v24
        } else {
            var v25 uint64 = 4294967295
            v25 += 747
            var v26 uint32 = 2147483647
            _ = v25
            _ = v26
        }
        switch (p1 / ((p1) | 1)) {
        case -487:
            p1 += p1
        case 42:
            var v27 uint = 42
            _ = v27
        }
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v21
    return 218
}

func main() {
    var v1 int8 = -128
    v2 := make(map[int]string)
    var v3 int16 = -1
    switch ("\n\t" + "a") {
    case "foo bar":
        if true {
            v1 = v1
        }
    case "a":
        var v28 int64 = 565
        _ = v28
    case "":
        if true {
            v3 = (v3 + -541)
            v3 -= v3
            v1 = -14
        }
    }
    v29 := GenS2{F0: -418}
    var v30 uint64 = 2
    v31 := GenS3{F0: 419, F1: 32767, F2: 4294967295}
    if (32767 != 42) {
        v32 := GenS4{F0: 9223372036854775807}
        var v33 int8 = 42
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        var v35 float64 = (3.14159 + 0.0)
        var v36 int = genFn0(uint(v30))
        var v37 uint8 = 2
        _ = fi34
        _ = v35
        _ = v36
        _ = v37
    }
    switch "" {
    case "":
        switch -2147483648 {
        case -164:
            var v38 float32 = (0.1 - (0.1 - 3.14159))
            _ = v38
        case -978:
            var v39 bool = false
            _ = v39
        case 1:
            for fi40 := 0; fi40 < 3; fi40++ {
                v30 = (v30 / ((v30) | 1))
                fmt.Println(fi40)
                _ = len(v2)
                _ = fi40
            }
        }
    case "\n\t":
        fmt.Println(v1, v30, v3)
    case "hello":
        if (1000 == v3) {
            v41 := []int{127, 127}
            _ = v41
        } else {
            switch "\n\t" {
            case "a":
                switch 32768 {
                case 0:
                    if false {
                        v42 := GenS1{F0: 42, F1: 251}
                        _ = v42
                    }
                case 127:
                    for fi43 := 0; fi43 < 10; fi43++ {
                        if false {
                            if (32767 <= 768) {
                                if (!false) {
                                    var v44 int16 = 32767
                                    _ = v44
                                }
                            } else {
                                var v45 uint16 = 280
                                _ = v45
                            }
                            for fi46 := 0; fi46 < 9; fi46++ {
                                var v47 int8 = -128
                                v1 += v47
                                _ = fi46
                                _ = v47
                            }
                        }
                        v48 := GenS2{F0: (fi43 ^ fi43)}
                        if false {
                            for fi49 := 0; fi49 < 7; fi49++ {
                                if (v3 <= int16(v30)) {
                                    if (1e10 > 1.5) {
                                        v30 += v30
                                        v50 := GenS3{F0: 42, F1: 1, F2: 255}
                                        _ = v50
                                    }
                                }
                                v51 := GenS5{F0: "foo bar", F1: 1, F2: v1}
                                _ = fi49
                                _ = v51
                            }
                            for fi52 := 0; fi52 < 9; fi52++ {
                                v3 -= v3
                                v30 -= (v30 >> ((genFn1(-32768, fi43)) & 63))
                                _ = fi52
                            }
                            v1 = (v1 / (((v1 << ((255) & 63))) | 1))
                        }
                        _ = fi43
                        _ = v48
                    }
                default:
                    v53 := GenS6{F0: 4294967295, F1: ""}
                    _ = v53
                }
            case "":
                if false {
                    var v54 uint64 = v30
                    _ = v54
                }
            }
        }
    }
    for fi55 := 0; fi55 < 2; fi55++ {
        if false {
            v30 -= 38
        } else {
            v1 = v1
        }
        _ = fi55
    }
    fmt.Println(v30)
    if true {
        v3 += v3
    } else {
        v56 := []int{9223372036854775807, 9223372036854775807, genFn0(520)}
        fmt.Println(v3)
        _ = v56
    }
    if true {
        var v57 int64 = -58
        for fi58 := 0; fi58 < 8; fi58++ {
            switch fi58 {
            case 0:
                for fi59 := 0; fi59 < 12; fi59++ {
                    v60 := GenS7{F0: -1}
                    switch fi59 {
                    case 127:
                        v2[(fi59 & genFn0(genFn1(32767, genFn0(4294967295))))] = ""
                    case -1:
                        switch fi58 {
                        case 9223372036854775807:
                            v61 := GenS1{F0: 255, F1: 0}
                            _ = v61
                        case 128:
                            v62 := GenS2{F0: fi58}
                            _ = v62
                        }
                    }
                    switch "" {
                    case "\n\t":
                        var v63 int16 = (v3 + (v3 / (((v3 ^ v3)) | 1)))
                        _ = v63
                    case "a":
                        if (0.0 <= 0.1) {
                            var v64 float32 = (0.1 - 0.0)
                            _ = v64
                        } else {
                            var v65 uint = 1
                            v66 := v2[(fi59 / ((-1) | 1))]
                            var v67 string = ("a" + "a")
                            _ = v65
                            _ = v66
                            _ = v67
                        }
                    case "hello":
                        v57 = (v57 >> ((genFn1((v3 - (v3 >> ((65535) & 63))), (-(fi58)))) & 63))
                    }
                    switch "" {
                    case "foo bar":
                        v68 := []int{9223372036854775807, (fi59 - (fi58 - (fi58 / (((fi58 % (((fi59 | fi59)) | 1))) | 1))))}
                        _ = v68
                    }
                    v30 += v30
                    _ = fi59
                    _ = v60
                }
            case 42:
                var v69 int = fi58
                _ = v69
            case 271:
                fmt.Println(fi58, v3, v3)
            default:
                if true {
                    if false {
                        var v70 int64 = v57
                        for fi71 := 0; fi71 < 19; fi71++ {
                            v72 := GenS8{F0: true, F1: (^(v30)), F2: fi58}
                            if (255 == 827) {
                                v70 -= 9223372036854775807
                            } else {
                                var v73 uint16 = 65535
                                if (false || ("" <= "a")) {
                                    v74 := GenS5{F0: "a", F1: genFn1(-32768, genFn0(195)), F2: (v1 ^ v1)}
                                    v2[fi71] = (("\n\t" + "foo bar") + "\n\t")
                                    _ = v74
                                }
                                v3 += 32767
                                v75 := GenS7{F0: (fi58 & fi58)}
                                _ = v73
                                _ = v75
                            }
                            fmt.Println(fi58, fi71)
                            _ = fi71
                            _ = v72
                        }
                        _ = v70
                    } else {
                        v1 -= (v1 / (((v1 / (((^(v1))) | 1))) | 1))
                        if (!false) {
                            switch "foo bar" {
                            case "a":
                                v76 := GenS4{F0: 493}
                                _ = v76
                            case "":
                                if (42 < -1) {
                                    for fi77 := 0; fi77 < 19; fi77++ {
                                        if false {
                                            switch ("\n\t" + ("hello" + ("hello" + ("hello" + (("\n\t" + "a") + ("hello" + "\n\t")))))) {
                                            case "hello":
                                                switch ("\n\t" + ("a" + ("hello" + ("hello" + "hello")))) {
                                                case "\n\t":
                                                    v30 = v30
                                                default:
                                                    v78 := GenS1{F0: 651, F1: 703}
                                                    _ = v78
                                                }
                                            }
                                            v57 = 0
                                            switch ((("\n\t" + (("hello" + "\n\t") + "a")) + "") + "") {
                                            case "foo bar":
                                                _ = len(v2)
                                            default:
                                                var v79 uint64 = (v30 + 42)
                                                _ = v79
                                            }
                                            switch "hello" {
                                            case "":
                                                v3 = v3
                                            case "hello":
                                                v80 := GenS9{F0: "a", F1: false, F2: v1}
                                                _ = v80
                                            case "a":
                                                var v81 int32 = -556
                                                _ = v81
                                            }
                                            delete(v2, fi58)
                                            v82 := v2[genFn0(18446744073709551615)]
                                            _ = v82
                                        } else {
                                            var v83 float32 = float32(0.1)
                                            _ = v83
                                        }
                                        v84 := []int{588}
                                        _ = fi77
                                        _ = v84
                                    }
                                    switch fi58 {
                                    case 32768:
                                        var v85 uint32 = 2
                                        _ = v85
                                    }
                                    for fi86 := 0; fi86 < 1; fi86++ {
                                        if ((v30 % (((^(v30))) | 1)) == v30) {
                                            var v87 uint8 = 131
                                            v88 := GenS5{F0: "a", F1: 65535, F2: 42}
                                            _ = v87
                                            _ = v88
                                        } else {
                                            v89 := GenS6{F0: 173, F1: ("hello" + ("a" + "foo bar"))}
                                            switch genFn0(380) {
                                            case -816:
                                                v1 -= 1
                                            }
                                            _ = v89
                                        }
                                        v30 -= 18446744073709551615
                                        _ = fi86
                                    }
                                } else {
                                    v57 -= v57
                                    for fi90 := 0; fi90 < 13; fi90++ {
                                        v30 -= (v30 - (v30 - (^(v30))))
                                        if false {
                                            v91 := v2[127]
                                            _ = len(v2)
                                            if false {
                                                var v92 float64 = (0.1 * 3.14159)
                                                _ = v92
                                            } else {
                                                if false {
                                                    v93 := GenS6{F0: 65535, F1: "\n\t"}
                                                    _ = v93
                                                }
                                                if false {
                                                    var v94 uint32 = 1
                                                    v3 = 1000
                                                    switch (^(fi90)) {
                                                    case 32767:
                                                        var v95 uint32 = (v94 - 962)
                                                        _ = v95
                                                    default:
                                                        var v96 int = genFn0(65535)
                                                        _ = v96
                                                    }
                                                    var v97 int8 = (v1 / (((v1 / (((v1 ^ (v1 - v1))) | 1))) | 1))
                                                    if true {
                                                        v98 := GenS10{F0: 189, F1: (v94 * v94), F2: (-(v57))}
                                                        v94 += 2147483648
                                                        switch v91 {
                                                        case "foo bar":
                                                            var v99 uint8 = 0
                                                            _ = v99
                                                        case "a":
                                                            v100 := GenS11{F0: v57, F1: 2, F2: (650 == 2147483647)}
                                                            _ = v100
                                                        default:
                                                            var v101 uint8 = 42
                                                            _ = v101
                                                        }
                                                        _ = v98
                                                    } else {
                                                        fmt.Println(v57, v91, v97)
                                                        v91 = "\n\t"
                                                        switch (v91 + v91) {
                                                        case "a":
                                                            fmt.Println(v1, v1)
                                                        default:
                                                            var v102 bool = true
                                                            v103 := GenS10{F0: 0, F1: (v94 - (v94 + v94)), F2: v57}
                                                            var v104 uint32 = (v94 ^ 42)
                                                            _ = v102
                                                            _ = v103
                                                            _ = v104
                                                        }
                                                    }
                                                    _ = v94
                                                    _ = v97
                                                } else {
                                                    v91 = v91
                                                    fmt.Println(fi90, fi58, fi90)
                                                    var v105 int = (fi58 * 32767)
                                                    _ = v105
                                                }
                                                if (((v3 / (((^(v3))) | 1)) > v3) && (v91 < v91)) {
                                                    var v106 int64 = v57
                                                    _ = v106
                                                }
                                            }
                                            _ = v91
                                        } else {
                                            var v107 string = ""
                                            _ = v107
                                        }
                                        var v108 int64 = (v57 % (((v57 & -238)) | 1))
                                        _ = fi90
                                        _ = v108
                                    }
                                    v109 := GenS12{F0: 1, F1: 42}
                                    _ = v109
                                }
                            case "\n\t":
                                var v110 int16 = (v3 - int16(v30))
                                _ = v110
                            default:
                                for fi111 := 0; fi111 < 19; fi111++ {
                                    fmt.Println(v30, v3, fi111)
                                    switch ("foo bar" + ("foo bar" + "foo bar")) {
                                    case "":
                                        v3 = v3
                                    default:
                                        switch "a" {
                                        case "\n\t":
                                            if (uint8(v30) < uint8(v30)) {
                                                switch fi111 {
                                                case -2147483648:
                                                    v112 := GenS5{F0: "", F1: 255, F2: v1}
                                                    _ = v112
                                                case 799:
                                                    v113 := []int{-2147483648}
                                                    _ = v113
                                                default:
                                                    var v114 int32 = 2147483647
                                                    _ = v114
                                                }
                                                v115 := GenS13{F0: false}
                                                _ = v115
                                            }
                                        case "foo bar":
                                            v116 := []int{genFn0(61)}
                                            _ = v116
                                        case "hello":
                                            v57 = (v57 - v57)
                                        default:
                                            v57 += (^(v57))
                                            v117 := GenS4{F0: v57}
                                            v118 := []int{int(v30), fi58, int(v1)}
                                            v30 = 0
                                            v119 := GenS13{F0: (-2.25 <= 0.0)}
                                            _ = v117
                                            _ = v118
                                            _ = v119
                                        }
                                    }
                                    _ = fi111
                                }
                            }
                            switch fi58 {
                            case 0:
                                v120 := GenS14{F0: v1, F1: (-(v57))}
                                _ = v120
                            default:
                                if true {
                                    v121 := GenS15{F0: (("hello" + "hello") + ("a" + ("a" + "")))}
                                    _ = v121
                                } else {
                                    for fi122 := 0; fi122 < 9; fi122++ {
                                        fmt.Println(fi58, v3, fi58)
                                        fmt.Println(v30)
                                        v1 = v1
                                        _ = fi122
                                    }
                                    v123 := []int{genFn0(2), (fi58 + fi58), fi58}
                                    _ = v123
                                }
                            }
                            var v124 bool = true
                            _ = v124
                        }
                    }
                    fmt.Println(v57)
                    v2[-40] = "hello"
                } else {
                    for fi125 := 0; fi125 < 20; fi125++ {
                        _ = len(v2)
                        switch ("" + "") {
                        case "\n\t":
                            v126 := GenS16{F0: (!true)}
                            _ = v126
                        case "a":
                            var v127 uint64 = (v30 * (^(v30)))
                            var v128 int16 = v3
                            _ = v127
                            _ = v128
                        case "":
                            v3 = (v3 - (v3 ^ v3))
                        default:
                            switch "hello" {
                            case "":
                                v3 = 32767
                            case "a":
                                v30 = (v30 * v30)
                                v129 := GenS8{F0: false, F1: v30, F2: -2147483648}
                                v130 := GenS10{F0: 1, F1: 0, F2: v57}
                                _ = v129
                                _ = v130
                            }
                        }
                        var v131 float64 = ((-2.25 - 3.14159) * (0.0 * 1e10))
                        _ = fi125
                        _ = v131
                    }
                    switch (fi58 % ((genFn0(0)) | 1)) {
                    case -1:
                        for fi132 := 0; fi132 < 2; fi132++ {
                            var v133 float64 = (float64(float32(-2.25)) * 0.0)
                            _ = fi132
                            _ = v133
                        }
                    case 127:
                        var v134 uint8 = 2
                        _ = v134
                    default:
                        if (!true) {
                            for fi135 := 0; fi135 < 13; fi135++ {
                                v136 := GenS6{F0: 13, F1: "hello"}
                                _ = fi135
                                _ = v136
                            }
                        } else {
                            v3 = -32768
                        }
                    }
                    _ = len(v2)
                }
            }
            v137 := []int{-9223372036854775808, fi58, fi58, (fi58 & -1)}
            _ = fi58
            _ = v137
        }
        v57 = v57
        _ = v57
    } else {
        var v138 uint8 = 42
        _ = v138
    }
    if false {
        v139 := []int{-457, 127, -9223372036854775808}
        _ = v139
    } else {
        v140 := GenS17{F0: 65535, F1: -951}
        v1 = v1
        v141 := []int{}
        _ = v140
        _ = v141
    }
    v142 := v2[genFn0(136)]
    for fi143 := 0; fi143 < 1; fi143++ {
        var v144 float64 = 0.0
        fmt.Println(v142)
        v145 := GenS18{F0: genFn0(1)}
        _ = fi143
        _ = v144
        _ = v145
    }
    v146 := GenS7{F0: 128}
    v147 := GenS7{F0: 2147483647}
    fmt.Println(v1)
    if false {
        if (true && (true && (-383 >= 480))) {
            fmt.Println(v30, v30, v3)
            v142 = (v142 + v142)
            v3 += (v3 / ((42) | 1))
        } else {
            var v148 float32 = 0.0
            v1 = 127
            for fi149 := 0; fi149 < 9; fi149++ {
                v148 += float32((3.14159 - 0.0))
                if ((!true) || (-597 != 844)) {
                    var v150 uint = 22
                    v3 = (v3 * (v3 | v3))
                    for fi151 := 0; fi151 < 15; fi151++ {
                        var v152 int = genFn0(990)
                        fmt.Println(v1, v30)
                        var v153 uint = (v150 | v150)
                        v154 := []int{}
                        _ = fi151
                        _ = v152
                        _ = v153
                        _ = v154
                    }
                    _ = v150
                } else {
                    v155 := []int{fi149}
                    _ = v155
                }
                _ = fi149
            }
            _ = v148
        }
        var v156 int16 = -863
        _ = v156
    }
    fmt.Println("v1:", v1)
    fmt.Println("v142:", v142)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v142
    _ = v146
    _ = v147
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
