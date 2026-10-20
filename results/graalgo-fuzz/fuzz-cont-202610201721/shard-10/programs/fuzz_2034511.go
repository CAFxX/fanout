package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 string; F1 int16; F2 uint }
type GenS4 struct { F0 int16; F1 int8; F2 int32 }
type GenS5 struct { F0 int64; F1 int16; F2 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 int16; F2 uint32 }
type GenS8 struct { F0 uint8; F1 bool }
type GenS9 struct { F0 int8; F1 int; F2 uint8 }
type GenS10 struct { F0 uint; F1 uint8 }
type GenS11 struct { F0 int64; F1 uint; F2 int8 }
type GenS12 struct { F0 uint32; F1 bool }
type GenS13 struct { F0 uint8; F1 int16 }
type GenS14 struct { F0 int32; F1 uint64; F2 int32 }
type GenS15 struct { F0 uint32; F1 int }
type GenS16 struct { F0 int8; F1 bool; F2 uint16 }
type GenS17 struct { F0 int8; F1 int; F2 string }
type GenS18 struct { F0 int32; F1 uint16; F2 uint }
type GenS19 struct { F0 int32; F1 uint8 }
type GenS20 struct { F0 int16 }
func genFn0(p0 int32) int8 {
    if true {
        for fi4 := 0; fi4 < 16; fi4++ {
            if true {
                fmt.Println(p0, p0)
                var v5 int = 127
                switch fi4 {
                case -2147483648:
                    switch (fi4 * fi4) {
                    case 1:
                        var v6 int16 = 784
                        _ = v6
                    case 128:
                        v7 := []int{fi4, v5, v5, fi4}
                        _ = v7
                    }
                case 127:
                    for fi8 := 0; fi8 < 3; fi8++ {
                        var v9 uint32 = 1
                        if ("\n\t" != (("hello" + "") + ("foo bar" + "foo bar"))) {
                            p0 += (^(p0))
                        }
                        v10 := []int{fi4}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                case -691:
                    for fi11 := 0; fi11 < 6; fi11++ {
                        if (!true) {
                            if true {
                                var v12 uint = 255
                                _ = v12
                            } else {
                                v13 := GenS1{F0: v5, F1: fi4}
                                _ = v13
                            }
                            var v14 float64 = 1e10
                            var v15 uint64 = 570
                            _ = v14
                            _ = v15
                        }
                        p0 = (p0 % ((0) | 1))
                        var v16 float32 = (0.1 + ((float32(1.5) * -2.25) + 1.5))
                        _ = fi11
                        _ = v16
                    }
                }
                _ = v5
            }
            _ = fi4
        }
        v17 := GenS2{F0: p0, F1: "hello"}
        if true {
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := []int{fi18, 32768}
                _ = fi18
                _ = v19
            }
            var v20 int = -476
            v21 := GenS2{F0: p0, F1: "a"}
            _ = v20
            _ = v21
        }
        _ = v17
    } else {
        var v22 int = 622
        v23 := GenS2{F0: p0, F1: "foo bar"}
        v24 := GenS2{F0: p0, F1: (("hello" + "a") + "")}
        _ = v22
        _ = v23
        _ = v24
    }
    var v25 uint = 1
    var v26 int16 = -32768
    var v27 string = "\n\t"
    _ = p0
    _ = v25
    _ = v26
    _ = v27
    return -128
}

func main() {
    var v1 int = 0
    v2 := []int{}
    v3 := make(map[int]bool)
    fmt.Println(v1)
    if (false || (!(1 != 243))) {
        switch 128 {
        case 369:
            for fi28 := 0; fi28 < 3; fi28++ {
                v1 = (v1 ^ (v1 >> ((255) & 63)))
                var v29 int = (fi28 ^ v1)
                if true {
                    v29 = v29
                } else {
                    if false {
                        switch v1 {
                        case 9223372036854775807:
                            switch "foo bar" {
                            case "a":
                                var v30 int32 = int32(v29)
                                _ = v30
                            case "":
                                fmt.Println(v29)
                            case "foo bar":
                                for fi31 := 0; fi31 < 3; fi31++ {
                                    var v32 float64 = 0.0
                                    _ = fi31
                                    _ = v32
                                }
                            }
                        case -9223372036854775808:
                            v33 := GenS3{F0: "\n\t", F1: 270, F2: 255}
                            _ = v33
                        case 32768:
                            var v34 string = "foo bar"
                            _ = v34
                        default:
                            var v35 int32 = 1
                            _ = v35
                        }
                    } else {
                        var v36 int16 = 112
                        var v37 uint16 = 2
                        _ = v36
                        _ = v37
                    }
                    if (fi28 <= fi28) {
                        var v38 float32 = (((3.14159 - float32(3.14159)) + (1e10 * float32((1e10 - 1.5)))) * 1.5)
                        v39 := GenS4{F0: -1, F1: -1, F2: -952}
                        _ = v38
                        _ = v39
                    } else {
                        switch v1 {
                        case -1:
                            fmt.Println(fi28, v1, v29)
                        case -2147483648:
                            fmt.Println(v1, v1)
                        default:
                            v40 := GenS2{F0: 42, F1: "foo bar"}
                            _ = v40
                        }
                    }
                }
                _ = fi28
                _ = v29
            }
        case 9223372036854775807:
            v41 := GenS5{F0: 42, F1: 646, F2: 2}
            _ = v41
        case 614:
            fmt.Println(v1)
        }
        v2 = append(v2, v1)
    }
    var v42 float64 = 1.5
    v43 := GenS5{F0: 9223372036854775807, F1: 986, F2: 42}
    fmt.Println(v42, v1)
    if true {
        for fi44 := 0; fi44 < 6; fi44++ {
            var v45 int64 = -1
            var v46 uint8 = 128
            switch "hello" {
            case "\n\t":
                for fi47 := 0; fi47 < 11; fi47++ {
                    fmt.Println(v1)
                    var v48 uint64 = 2
                    switch ("a" + "a") {
                    case "\n\t":
                        fmt.Println(v42)
                    default:
                        if true {
                            var v49 int64 = (v45 ^ v45)
                            _ = v49
                        } else {
                            v1 -= v1
                            if false {
                                if (!true) {
                                    var v50 int = (v1 + 0)
                                    fmt.Println(v46)
                                    _ = v50
                                } else {
                                    if true {
                                        var v51 int64 = (^(v45))
                                        _ = v51
                                    }
                                    v52 := v3[(fi47 * 875)]
                                    _ = v52
                                }
                            } else {
                                if true {
                                    v46 += (^(v46))
                                } else {
                                    if (false || false) {
                                        fmt.Println(v46, v1, v46)
                                    }
                                    v1 -= (fi47 - -490)
                                    var v53 float64 = 1e10
                                    _ = v53
                                }
                                v45 = v45
                            }
                            if false {
                                v45 -= v45
                            }
                            if (625 <= 32767) {
                                v46 -= (^(v46))
                            }
                            v54 := GenS3{F0: "hello", F1: 1000, F2: 42}
                            v46 += (v46 / (((v46 * v46)) | 1))
                            var v55 float32 = (-2.25 + ((1e10 + 1.5) - (1e10 + 3.14159)))
                            if true {
                                v2 = append(v2, 15)
                                var v56 uint32 = 42
                                _ = v56
                            } else {
                                var v57 int = (v1 / ((-261) | 1))
                                _ = v57
                            }
                            v58 := GenS6{F0: ("foo bar" + ("" + "\n\t"))}
                            _ = v54
                            _ = v55
                            _ = v58
                        }
                    }
                    _ = fi47
                    _ = v48
                }
            default:
                v42 -= float64((float32((v42 - v42)) - (float32(v42) - 1e10)))
            }
            _ = fi44
            _ = v45
            _ = v46
        }
        switch (v1 - 9223372036854775807) {
        case 755:
            if (42 == 2147483647) {
                _ = len(v2)
                v59 := GenS4{F0: 1000, F1: 44, F2: 1}
                if (true || (false || (-455 == 42))) {
                    _ = len(v2)
                } else {
                    var v60 uint32 = 0
                    switch v1 {
                    case 859:
                        if false {
                            for fi61 := 0; fi61 < 12; fi61++ {
                                if (215 >= 17) {
                                    v42 -= v42
                                    fmt.Println(v42)
                                }
                                _ = fi61
                            }
                        }
                    case -2147483648:
                        switch int(v60) {
                        case -1:
                            if false {
                                for fi62 := 0; fi62 < 14; fi62++ {
                                    delete(v3, fi62)
                                    var v63 float64 = v42
                                    _ = fi62
                                    _ = v63
                                }
                                v60 -= (v60 & v60)
                            } else {
                                v64 := GenS3{F0: "", F1: 32767, F2: 255}
                                var v65 bool = false
                                _ = v64
                                _ = v65
                            }
                        default:
                            delete(v3, (-(v1)))
                        }
                    default:
                        _ = len(v2)
                    }
                    _ = v60
                }
                _ = v59
            }
        case 1:
            var v66 int32 = 836
            _ = v66
        case -2147483648:
            if (128 > 830) {
                var v67 uint32 = 2147483648
                _ = v67
            }
        }
        if false {
            if (!(float64(float32(float64(1.5))) < (v42 + v42))) {
                v1 -= v1
                var v68 bool = (42 <= -1)
                if v68 {
                    for fi69 := 0; fi69 < 16; fi69++ {
                        _ = len(v2)
                        var v70 uint = 2
                        _ = fi69
                        _ = v70
                    }
                    switch "" {
                    case "a":
                        fmt.Println(v1)
                    case "":
                        var v71 float64 = v42
                        _ = v71
                    case "\n\t":
                        fmt.Println(v1, v68, v1)
                    }
                } else {
                    if (((v1 / ((v1) | 1)) != (v1 * v1)) || false) {
                        if len(v2) > 0 {
                            v2[0] = (^(v1))
                        }
                        if (v68 && v68) {
                            v72 := GenS3{F0: "", F1: -759, F2: 0}
                            _ = v72
                        }
                    }
                    v1 = (v1 / ((v1) | 1))
                    v1 = v1
                }
                _ = v68
            } else {
                if false {
                    v73 := GenS3{F0: "\n\t", F1: 0, F2: 0}
                    switch ("\n\t" + "a") {
                    case "foo bar":
                        var v74 int32 = -317
                        _ = v74
                    case "a":
                        for fi75 := 0; fi75 < 10; fi75++ {
                            var v76 float64 = v42
                            _ = fi75
                            _ = v76
                        }
                    default:
                        var v77 uint64 = 1
                        _ = v77
                    }
                    var v78 bool = true
                    _ = v73
                    _ = v78
                } else {
                    v1 = (v1 * (v1 | -2147483648))
                    for fi79 := 0; fi79 < 8; fi79++ {
                        switch (("\n\t" + "\n\t") + "a") {
                        case "hello":
                            switch (fi79 - 1) {
                            case 932:
                                fmt.Println(fi79, fi79, v42)
                            case 32767:
                                v80 := GenS4{F0: -270, F1: 127, F2: 0}
                                _ = v80
                            default:
                                v81 := GenS3{F0: (("" + (("hello" + (((("hello" + "hello") + ("hello" + "a")) + "hello") + ("" + "\n\t"))) + "foo bar")) + "foo bar"), F1: 32767, F2: 771}
                                _ = v81
                            }
                        case "":
                            v42 = (v42 + v42)
                        }
                        fmt.Println(fi79, v42)
                        _ = fi79
                    }
                    var v82 int16 = -862
                    _ = v82
                }
                var v83 uint = 65535
                _ = v83
            }
            var v84 uint32 = 2147483648
            v84 -= 233
            _ = v84
        } else {
            switch "hello" {
            case "":
                switch ("foo bar" + "hello") {
                case "":
                    switch (v1 & v1) {
                    case 9223372036854775807:
                        if (!true) {
                            var v85 float32 = 0.0
                            var v86 int = v1
                            _ = v85
                            _ = v86
                        } else {
                            v42 = v42
                        }
                    case 127:
                        for fi87 := 0; fi87 < 13; fi87++ {
                            var v88 string = "hello"
                            _ = fi87
                            _ = v88
                        }
                    case 32767:
                        var v89 uint8 = 128
                        _ = v89
                    }
                default:
                    if (v1 < v1) {
                        fmt.Println(v1, v1)
                        var v90 bool = true
                        _ = v90
                    } else {
                        for fi91 := 0; fi91 < 4; fi91++ {
                            v42 += ((v42 - ((((v42 - v42) * v42) * float64(float32(3.14159))) / (1.5 / 1.0))) / -2.25)
                            for fi92 := 0; fi92 < 10; fi92++ {
                                v1 = (fi91 << ((65535) & 63))
                                _ = len(v2)
                                var v93 uint32 = 0
                                _ = fi92
                                _ = v93
                            }
                            _ = fi91
                        }
                    }
                }
            case "a":
                v94 := GenS5{F0: -9223372036854775808, F1: -32768, F2: 42}
                _ = v94
            case "hello":
                v1 -= 32768
            default:
                if ((136 > 42) && false) {
                    for fi95 := 0; fi95 < 6; fi95++ {
                        v1 -= v1
                        _ = fi95
                    }
                    var v96 string = "a"
                    if true {
                        if false {
                            if (32767 == v1) {
                                if (18446744073709551615 < 2) {
                                    for fi97 := 0; fi97 < 2; fi97++ {
                                        if (0.1 < (float32((float64(-2.25) * (float64(1.5) / v42))) + (-2.25 / -2.25))) {
                                            v98 := GenS7{F0: 1, F1: -1, F2: 2147483648}
                                            v42 -= v42
                                            _ = v98
                                        } else {
                                            if (42 == 1) {
                                                fmt.Println(v96, v1, v42)
                                                switch (v1 + v1) {
                                                case 42:
                                                    if ((v1 <= (^(v1))) && true) {
                                                        var v99 uint8 = 50
                                                        var v100 uint16 = 25
                                                        _ = v100
                                                        _ = v99
                                                    }
                                                case -582:
                                                    v42 = float64(1e10)
                                                case 0:
                                                    v1 = -9223372036854775808
                                                }
                                            } else {
                                                for fi101 := 0; fi101 < 7; fi101++ {
                                                    if (0 != 4294967295) {
                                                        switch fi97 {
                                                        case 4294967295:
                                                            var v102 int64 = 1
                                                            _ = v102
                                                        }
                                                        v2 = append(v2, fi101)
                                                        var v103 int16 = -32768
                                                        _ = v103
                                                    }
                                                    delete(v3, fi101)
                                                    v104 := GenS4{F0: 0, F1: 38, F2: -1}
                                                    if (!true) {
                                                        v42 -= (float64(1.5) / float64(((-2.25 + 3.14159) * (-2.25 * 3.14159))))
                                                        v2 = append(v2, 127)
                                                    } else {
                                                        fmt.Println(v96)
                                                        if false {
                                                            if false {
                                                                delete(v3, 0)
                                                                var v105 uint64 = 14
                                                                v2 = append(v2, 42)
                                                                _ = v105
                                                            }
                                                            switch (v96 + v96) {
                                                            case "hello":
                                                                v106 := GenS1{F0: (v1 + 1), F1: fi97}
                                                                _ = v106
                                                            default:
                                                                v107 := GenS8{F0: 0, F1: (915 >= 1000)}
                                                                _ = v107
                                                            }
                                                            var v108 uint = 42
                                                            _ = v108
                                                        } else {
                                                            var v109 uint64 = 42
                                                            var v110 uint = 1
                                                            var v111 int8 = 100
                                                            _ = v109
                                                            _ = v110
                                                            _ = v111
                                                        }
                                                    }
                                                    switch (("a" + "foo bar") + "foo bar") {
                                                    case "\n\t":
                                                        v112 := GenS9{F0: -127, F1: (fi97 * (fi101 + -779)), F2: 250}
                                                        var v113 uint16 = 24
                                                        _ = v112
                                                        _ = v113
                                                    case "foo bar":
                                                        v96 = (((v96 + ("a" + v96)) + ("" + "")) + ((v96 + v96) + "hello"))
                                                    }
                                                    _ = fi101
                                                    _ = v104
                                                }
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            for fi114 := 0; fi114 < 2; fi114++ {
                                                v115 := GenS8{F0: 255, F1: false}
                                                _ = fi114
                                                _ = v115
                                            }
                                        }
                                        if false {
                                            fmt.Println(v42)
                                            var v116 int32 = 628
                                            switch 8 {
                                            case -505:
                                                v117 := GenS10{F0: 255, F1: 19}
                                                _ = v117
                                            }
                                            _ = v116
                                        }
                                        _ = fi97
                                    }
                                }
                            } else {
                                v118 := v3[283]
                                for fi119 := 0; fi119 < 20; fi119++ {
                                    if (v118 || (!true)) {
                                        if true {
                                            v120 := GenS11{F0: 0, F1: 42, F2: -128}
                                            var v121 int32 = 2147483647
                                            _ = v120
                                            _ = v121
                                        }
                                        for fi122 := 0; fi122 < 3; fi122++ {
                                            v3[v1] = (v118 && v118)
                                            v118 = (-906 < 42)
                                            _ = fi122
                                        }
                                    }
                                    var v123 uint = 42
                                    _ = fi119
                                    _ = v123
                                }
                                _ = v118
                            }
                            v96 = ""
                            var v124 bool = false
                            _ = v124
                        }
                        v125 := GenS12{F0: 682, F1: false}
                        _ = v125
                    }
                    _ = v96
                } else {
                    for fi126 := 0; fi126 < 1; fi126++ {
                        v127 := GenS2{F0: -970, F1: ("\n\t" + "hello")}
                        var v128 float64 = 1.5
                        _ = fi126
                        _ = v127
                        _ = v128
                    }
                    fmt.Println(v1, v42)
                    var v129 uint16 = 0
                    _ = v129
                }
            }
            for fi130 := 0; fi130 < 19; fi130++ {
                v42 = v42
                v131 := v3[(fi130 * v1)]
                var v132 uint8 = 75
                _ = fi130
                _ = v131
                _ = v132
            }
        }
    } else {
        if (false || (v42 >= -2.25)) {
            v42 -= -2.25
            v42 = ((3.14159 - v42) * v42)
        } else {
            for fi133 := 0; fi133 < 9; fi133++ {
                v134 := GenS7{F0: -1, F1: -740, F2: 4294967295}
                if (42 < 18446744073709551615) {
                    v2 = append(v2, (fi133 % ((127) | 1)))
                    for fi135 := 0; fi135 < 20; fi135++ {
                        v136 := v3[-2147483648]
                        _ = fi135
                        _ = v136
                    }
                    var v137 uint = 0
                    _ = v137
                }
                for fi138 := 0; fi138 < 3; fi138++ {
                    v1 = (v1 / (((^(fi138))) | 1))
                    switch "" {
                    case "":
                        switch v1 {
                        case -9223372036854775808:
                            v139 := GenS2{F0: 0, F1: "foo bar"}
                            _ = v139
                        }
                    }
                    v140 := GenS13{F0: 2, F1: 0}
                    var v141 uint = 2
                    v1 -= 2147483647
                    _ = fi138
                    _ = v140
                    _ = v141
                }
                _ = fi133
                _ = v134
            }
        }
        if false {
            switch -127 {
            case 40:
                fmt.Println(v1, v1)
            case 4294967295:
                if true {
                    if (71 >= -114) {
                        for fi142 := 0; fi142 < 11; fi142++ {
                            for fi143 := 0; fi143 < 14; fi143++ {
                                v42 = 0.1
                                v1 = v1
                                var v144 float32 = (3.14159 / ((0.0 - ((1.5 - float32((v42 * -2.25))) / 1.5)) + (0.0 / 1.5)))
                                v42 += v42
                                _ = fi143
                                _ = v144
                            }
                            if false {
                                var v145 uint64 = 2
                                fmt.Println(v1, fi142)
                                _ = v145
                            } else {
                                for fi146 := 0; fi146 < 5; fi146++ {
                                    var v147 uint8 = 2
                                    if true {
                                        var v148 uint64 = 2
                                        v2 = append(v2, v1)
                                        _ = v148
                                    } else {
                                        var v149 bool = false
                                        v149 = (932 == 18446744073709551615)
                                        _ = v149
                                    }
                                    if (!(false || (!(3.14159 > 1.5)))) {
                                        var v150 float64 = v42
                                        v151 := GenS11{F0: -256, F1: 65535, F2: -128}
                                        v147 = v147
                                        _ = v150
                                        _ = v151
                                    }
                                    v42 = v42
                                    _ = fi146
                                    _ = v147
                                }
                            }
                            v42 = (float64(3.14159) / 1e10)
                            _ = fi142
                        }
                    }
                } else {
                    v1 -= -9223372036854775808
                    v152 := v3[(v1 | (v1 ^ (v1 % ((42) | 1))))]
                    var v153 float64 = v42
                    _ = v152
                    _ = v153
                }
            }
            var v154 uint8 = 1
            _ = v154
        } else {
            v42 -= ((0.0 / 1.5) * v42)
            v42 = ((float64((float32(v42) + float32(0.1))) + v42) - v42)
        }
        for fi155 := 0; fi155 < 7; fi155++ {
            for fi156 := 0; fi156 < 17; fi156++ {
                v42 = v42
                if ((v42 * v42) > ((v42 / float64(1e10)) * 1.5)) {
                    v157 := GenS14{F0: -2147483648, F1: 18446744073709551615, F2: -507}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if false {
                        var v158 int32 = 0
                        if true {
                            v159 := v3[(fi155 >> ((4294967295) & 63))]
                            v160 := GenS15{F0: 2147483648, F1: fi156}
                            fmt.Println(fi155, fi155)
                            _ = v159
                            _ = v160
                        } else {
                            if (true && (42 <= 2147483647)) {
                                switch fi155 {
                                case 0:
                                    v161 := GenS13{F0: 127, F1: -32768}
                                    _ = v161
                                case 127:
                                    var v162 uint16 = 42
                                    _ = v162
                                case 535:
                                    if false {
                                        v2 = append(v2, (fi156 | (fi156 >> ((1) & 63))))
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if false {
                                            if (255 <= 173) {
                                                if len(v2) > 0 {
                                                    v2[0] = (fi156 & fi155)
                                                }
                                                var v163 bool = true
                                                v158 -= (v158 ^ (v158 | -1))
                                                _ = v163
                                            }
                                            var v164 int32 = int32(fi156)
                                            v158 -= v164
                                            _ = v164
                                        } else {
                                            var v165 uint64 = 1
                                            if true {
                                                v3[-9223372036854775808] = false
                                                var v166 uint = 65535
                                                var v167 uint = uint(v165)
                                                if (true || (false && (2 >= 32767))) {
                                                    var v168 float32 = 3.14159
                                                    if (993 <= -9223372036854775808) {
                                                        if true {
                                                            var v169 int64 = 9223372036854775807
                                                            var v170 int8 = 100
                                                            var v171 uint32 = 318
                                                            var v172 float64 = (v42 * float64(3.14159))
                                                            _ = v169
                                                            _ = v170
                                                            _ = v171
                                                            _ = v172
                                                        }
                                                        v167 += v166
                                                    } else {
                                                        fmt.Println(v165)
                                                        v173 := GenS16{F0: -1, F1: false, F2: 2}
                                                        _ = v173
                                                    }
                                                    fmt.Println(v166, fi156)
                                                    _ = v168
                                                }
                                                if true {
                                                    if len(v2) > 0 {
                                                        v2[0] = -956
                                                    }
                                                }
                                                fmt.Println(v165)
                                                _ = v166
                                                _ = v167
                                            } else {
                                                var v174 int16 = -32768
                                                v175 := GenS16{F0: -1, F1: (1 >= 1), F2: 945}
                                                _ = v174
                                                _ = v175
                                            }
                                            if true {
                                                var v176 float64 = v42
                                                v1 -= 4294967295
                                                _ = v176
                                            } else {
                                                if (65535 < 795) {
                                                    var v177 int = fi156
                                                    v178 := GenS17{F0: 1, F1: v177, F2: "a"}
                                                    _ = v177
                                                    _ = v178
                                                } else {
                                                    switch (fi155 - fi156) {
                                                    case 46:
                                                        if false {
                                                            var v179 int64 = 108
                                                            var v180 int64 = v179
                                                            _ = v179
                                                            _ = v180
                                                        } else {
                                                            fmt.Println(v42, v1, fi155)
                                                            v181 := GenS18{F0: (v158 ^ v158), F1: 1, F2: 65535}
                                                            fmt.Println(fi155, fi155, fi156)
                                                            _ = v181
                                                        }
                                                    case 4294967295:
                                                        _ = len(v2)
                                                    case -2147483648:
                                                        v182 := GenS19{F0: (v158 | v158), F1: 127}
                                                        _ = v182
                                                    }
                                                    if (true || (true && (v158 == v158))) {
                                                        v183 := GenS9{F0: 1, F1: (fi155 + (v1 % (((fi155 % ((9223372036854775807) | 1))) | 1))), F2: uint8(v165)}
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v183
                                                    }
                                                }
                                                _ = len(v2)
                                            }
                                            _ = v165
                                        }
                                    }
                                default:
                                    v1 -= int(v158)
                                }
                            }
                        }
                        _ = v158
                    }
                    v184 := GenS1{F0: (fi155 / ((fi156) | 1)), F1: (v1 >> ((0) & 63))}
                    var v185 uint16 = 42
                    delete(v3, 32768)
                    if false {
                        var v186 int = (fi156 * v1)
                        v187 := GenS17{F0: 41, F1: v1, F2: "\n\t"}
                        v185 = 2
                        _ = v186
                        _ = v187
                    }
                    v185 = (v185 >> ((754) & 63))
                    _ = v157
                    _ = v184
                    _ = v185
                }
                var v188 int32 = 2147483647
                _ = fi156
                _ = v188
            }
            if (((!true) || true) && (643 < 42)) {
                var v189 uint16 = 255
                _ = v189
            } else {
                if (-1 < 1) {
                    switch fi155 {
                    case -863:
                        var v190 int8 = 42
                        _ = v190
                    case -9223372036854775808:
                        var v191 int8 = -1
                        _ = v191
                    case -261:
                        var v192 int8 = -1
                        _ = v192
                    }
                }
                v1 = v1
                switch (("\n\t" + "\n\t") + "hello") {
                case "foo bar":
                    v2 = append(v2, (fi155 ^ -9223372036854775808))
                }
            }
            _ = fi155
        }
    }
    var v193 int = (v1 / (((v1 << ((42) & 63))) | 1))
    v193 += v193
    v194 := GenS20{F0: 767}
    var v195 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v193:", v193)
    fmt.Println("v195:", v195)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v193
    _ = v194
    _ = v195
    _ = v2
    _ = v3
    _ = v42
    _ = v43
}
