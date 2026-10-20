package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 string; F1 bool; F2 int16 }
type GenS7 struct { F0 string; F1 uint32; F2 uint8 }
type GenS8 struct { F0 int64; F1 uint }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int64; F1 uint }
type GenS11 struct { F0 uint64; F1 int; F2 int32 }
type GenS12 struct { F0 uint64; F1 string; F2 int32 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int8; F1 uint }
type GenS16 struct { F0 int }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 uint8; F1 int; F2 int32 }
type GenS19 struct { F0 int; F1 int32; F2 uint16 }
type GenS20 struct { F0 int16; F1 int }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 uint64; F1 int32; F2 int32 }
type GenS23 struct { F0 int64 }
type GenS24 struct { F0 uint32 }

func main() {
    var v1 int64 = -995
    v2 := make(map[string]bool)
    var v3 uint8 = 255
    v3 += (v3 & v3)
    fmt.Println(v1)
    switch 0 {
    case 4294967295:
        switch -2147483648 {
        case 0:
            var v4 uint16 = 65535
            _ = v4
        }
    }
    v1 -= v1
    var v5 string = ""
    var v6 string = ("" + v5)
    if (!((v1 & v1) != v1)) {
        if true {
            v2[v5] = false
        } else {
            v7 := GenS1{F0: 32767}
            v8 := GenS2{F0: -1, F1: 42}
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: 32767}
        switch ("foo bar" + ((v6 + (v6 + v6)) + "hello")) {
        case "":
            var v10 string = ((((v6 + v5) + v5) + (v6 + (((((v5 + "hello") + (v6 + v6)) + "") + (("hello" + v5) + v6)) + v5))) + v5)
            _ = v10
        case "foo bar":
            if false {
                for fi11 := 0; fi11 < 12; fi11++ {
                    fmt.Println(v5, v5, v3)
                    v1 = v1
                    _ = fi11
                }
            } else {
                for fi12 := 0; fi12 < 1; fi12++ {
                    var v13 uint64 = 8
                    v14 := []int{(^(fi12)), fi12, int(v3), -954}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                var v15 int32 = -2147483648
                _ = v15
            }
        case "\n\t":
            if true {
                var v16 int16 = 32767
                for fi17 := 0; fi17 < 1; fi17++ {
                    v18 := GenS2{F0: 42, F1: 1}
                    v16 += v16
                    var v19 int = 127
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = v16
            } else {
                var v20 float64 = -2.25
                for fi21 := 0; fi21 < 5; fi21++ {
                    v3 += 42
                    for fi22 := 0; fi22 < 10; fi22++ {
                        v23 := []int{(-(fi22)), fi22}
                        var v24 float64 = -2.25
                        switch v6 {
                        case "\n\t":
                            var v25 uint16 = 42
                            _ = v25
                        case "foo bar":
                            switch v5 {
                            case "hello":
                                var v26 uint = 1
                                _ = v26
                            case "foo bar":
                                v5 = (v5 + ("a" + ((v6 + v5) + (((v5 + v5) + "hello") + "a"))))
                            default:
                                v27 := GenS2{F0: -1, F1: 4294967295}
                                _ = v27
                            }
                        default:
                            v28 := GenS1{F0: 255}
                            if (v3 > v3) {
                                var v29 uint8 = v3
                                v5 = v5
                                if (false || true) {
                                    v6 = v6
                                    var v30 int16 = 1000
                                    _ = v30
                                } else {
                                    switch "hello" {
                                    case "hello":
                                        v1 = (^(v1))
                                    case "":
                                        switch -2147483648 {
                                        case -9223372036854775808:
                                            var v31 int32 = -2147483648
                                            _ = v31
                                        default:
                                            _ = len(v2)
                                        }
                                    }
                                    v32 := GenS2{F0: 89, F1: (^(fi22))}
                                    _ = v32
                                }
                                _ = v29
                            } else {
                                fmt.Println(v20, v24)
                                if (float32(1e10) < 1.5) {
                                    var v33 float64 = (v24 * (float64(float32(v20)) * (float64(1.5) + (v20 - float64(((3.14159 - (((1e10 + 0.0) + -2.25) * float32(v24))) / 1.0))))))
                                    fmt.Println(v20, v1)
                                    fmt.Println(v1, fi22)
                                    var v34 uint64 = 0
                                    var v35 int16 = 32767
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                                fmt.Println(fi22)
                            }
                            if (true || (((true && true) || false) && false)) {
                                v5 = "foo bar"
                                if (42 < 0) {
                                    v24 = (((((((0.0 / (0.1 / 3.14159)) - v20) / 1.5) / (v24 + v24)) * (v24 - v24)) - 1e10) / (0.0 - (v20 * v24)))
                                    if len(v23) > 0 {
                                        v23[0] = (fi21 + fi21)
                                    }
                                    var v36 uint64 = 0
                                    _ = v36
                                }
                                var v37 int32 = -414
                                _ = v37
                            }
                            _ = v28
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    for fi38 := 0; fi38 < 8; fi38++ {
                        var v39 uint64 = 18446744073709551615
                        _ = fi38
                        _ = v39
                    }
                    _ = fi21
                }
                switch 42 {
                case -756:
                    v40 := []int{}
                    _ = v40
                }
                _ = v20
            }
        }
        _ = v9
    } else {
        if (808 == 236) {
            for fi41 := 0; fi41 < 13; fi41++ {
                for fi42 := 0; fi42 < 16; fi42++ {
                    if true {
                        var v43 float32 = -2.25
                        var v44 int8 = 42
                        _ = v43
                        _ = v44
                    } else {
                        var v45 int32 = 1
                        _ = v45
                    }
                    v5 = v5
                    _ = fi42
                }
                _ = fi41
            }
        }
        switch v5 {
        case "foo bar":
            for fi46 := 0; fi46 < 12; fi46++ {
                v47 := []int{(fi46 << ((18446744073709551615) & 63)), (-(fi46)), (fi46 | fi46)}
                if false {
                    var v48 uint32 = 696
                    var v49 bool = false
                    _ = v48
                    _ = v49
                } else {
                    fmt.Println(v5)
                    delete(v2, v6)
                    for fi50 := 0; fi50 < 18; fi50++ {
                        var v51 uint32 = 1
                        _ = fi50
                        _ = v51
                    }
                }
                for fi52 := 0; fi52 < 17; fi52++ {
                    v3 = v3
                    var v53 int32 = -123
                    delete(v2, v6)
                    var v54 uint = 382
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                _ = fi46
                _ = v47
            }
        }
        v55 := GenS2{F0: 70, F1: 1}
        _ = v55
    }
    v2[v6] = (415 > 0)
    for fi56 := 0; fi56 < 4; fi56++ {
        var v57 uint32 = 0
        v6 = v5
        var v58 bool = true
        _ = fi56
        _ = v57
        _ = v58
    }
    if ((v6 + (("" + v6) + ("a" + (("a" + ((v5 + ("\n\t" + v5)) + "\n\t")) + "\n\t")))) > v5) {
        if (32768 != -918) {
            for fi59 := 0; fi59 < 2; fi59++ {
                for fi60 := 0; fi60 < 20; fi60++ {
                    if true {
                        v6 = v5
                        var v61 int8 = -128
                        switch 4294967295 {
                        case 645:
                            var v62 uint16 = 677
                            _ = v62
                        case 32768:
                            var v63 uint64 = 340
                            _ = v63
                        case -379:
                            fmt.Println(v5, fi60, v3)
                        }
                        var v64 int = (fi59 ^ 1)
                        _ = v61
                        _ = v64
                    }
                    switch v5 {
                    case "":
                        var v65 int16 = 1
                        _ = v65
                    case "foo bar":
                        v66 := GenS1{F0: 42}
                        _ = v66
                    case "hello":
                        fmt.Println(v5, v3, v5)
                    default:
                        v67 := []int{fi59, fi60, -128}
                        _ = v67
                    }
                    v5 = (v6 + v6)
                    _ = fi60
                }
                _ = fi59
            }
            v68 := GenS3{F0: 93}
            switch "\n\t" {
            case "foo bar":
                v3 = v3
            case "a":
                for fi69 := 0; fi69 < 11; fi69++ {
                    if true {
                        for fi70 := 0; fi70 < 5; fi70++ {
                            if true {
                                if true {
                                    v71 := v2[(v6 + v6)]
                                    _ = v71
                                } else {
                                    v6 = v5
                                    v3 = v3
                                }
                            } else {
                                _ = len(v2)
                                if false {
                                    if false {
                                        if ((3.14159 - -2.25) < (-2.25 - 3.14159)) {
                                            switch "\n\t" {
                                            case "hello":
                                                if (fi70 >= fi69) {
                                                    v72 := GenS4{F0: v3, F1: 2147483648}
                                                    var v73 int = fi69
                                                    v73 = 32768
                                                    _ = v72
                                                    _ = v73
                                                }
                                            case "\n\t":
                                                v2[(v5 + ("" + v6))] = true
                                            default:
                                                v1 = 0
                                            }
                                        }
                                        var v74 int8 = -128
                                        v75 := GenS2{F0: (-(v74)), F1: (-(fi70))}
                                        _ = v74
                                        _ = v75
                                    }
                                }
                                v6 = (v6 + (v6 + "a"))
                                switch "" {
                                case "a":
                                    var v76 string = v6
                                    _ = v76
                                }
                            }
                            var v77 float32 = (3.14159 + 1.5)
                            switch 4294967295 {
                            case -343:
                                v2[((v6 + (v5 + (v5 + v6))) + v5)] = (18446744073709551615 > 817)
                            case -2147483648:
                                v77 += (float32((0.1 * -2.25)) - float32(1e10))
                            case 2147483647:
                                switch (v6 + v6) {
                                case "hello":
                                    var v78 bool = true
                                    _ = v78
                                case "":
                                    v79 := GenS5{F0: -1, F1: (-32768 != -32768)}
                                    _ = v79
                                case "\n\t":
                                    v80 := GenS6{F0: "\n\t", F1: true, F2: 1}
                                    _ = v80
                                default:
                                    v81 := []int{fi69}
                                    _ = v81
                                }
                            }
                            _ = fi70
                            _ = v77
                        }
                        v2["\n\t"] = true
                        switch "" {
                        case "hello":
                            for fi82 := 0; fi82 < 11; fi82++ {
                                v83 := []int{}
                                _ = fi82
                                _ = v83
                            }
                        case "":
                            var v84 int64 = v1
                            _ = v84
                        default:
                            v3 = 2
                        }
                    } else {
                        v1 -= 0
                        var v85 int16 = -1
                        _ = v85
                    }
                    _ = fi69
                }
            default:
                v6 = v6
            }
            _ = v68
        } else {
            v86 := []int{}
            v87 := GenS1{F0: 42}
            v6 = "foo bar"
            _ = v86
            _ = v87
        }
        if (4294967295 >= 650) {
            v88 := []int{}
            for fi89 := 0; fi89 < 5; fi89++ {
                if ((2 < 520) && true) {
                    _ = len(v88)
                    fmt.Println(v6, v6, v1)
                    _ = len(v2)
                }
                if (961 == 1) {
                    var v90 uint16 = 93
                    for fi91 := 0; fi91 < 18; fi91++ {
                        v5 = ("a" + (v6 + ((v5 + v5) + "\n\t")))
                        _ = fi91
                    }
                    fmt.Println(v6, fi89)
                    _ = v90
                }
                _ = len(v88)
                _ = fi89
            }
            if (!false) {
                if false {
                    v3 = 2
                    v92 := GenS5{F0: 32767, F1: false}
                    _ = v92
                } else {
                    var v93 int64 = v1
                    _ = v93
                }
                switch -2147483648 {
                case 210:
                    for fi94 := 0; fi94 < 8; fi94++ {
                        var v95 uint16 = 535
                        v96 := GenS2{F0: 1, F1: fi94}
                        _ = fi94
                        _ = v95
                        _ = v96
                    }
                case -1:
                    v97 := GenS7{F0: v5, F1: 342, F2: 0}
                    _ = v97
                default:
                    var v98 int64 = (v1 % ((v1) | 1))
                    _ = v98
                }
            } else {
                v3 += 127
                for fi99 := 0; fi99 < 1; fi99++ {
                    v88 = append(v88, fi99)
                    _ = fi99
                }
                var v100 float64 = (0.0 - 3.14159)
                _ = v100
            }
            _ = v88
        }
        fmt.Println(v6, v5, v5)
    }
    fmt.Println(v5, v5)
    for fi101 := 0; fi101 < 19; fi101++ {
        v102 := GenS8{F0: 42, F1: 0}
        fmt.Println(v3)
        var v103 uint32 = 65535
        _ = fi101
        _ = v102
        _ = v103
    }
    if false {
        v104 := GenS7{F0: "", F1: 2147483647, F2: (v3 | (v3 * (v3 >> ((279) & 63))))}
        switch v6 {
        case "\n\t":
            if true {
                for fi105 := 0; fi105 < 3; fi105++ {
                    v5 = ((v5 + v5) + v6)
                    switch fi105 {
                    case 127:
                        if (false && (0 < -560)) {
                            var v106 int32 = 2147483647
                            delete(v2, "foo bar")
                            var v107 string = v5
                            _ = v106
                            _ = v107
                        } else {
                            if false {
                                var v108 float64 = (-2.25 - ((1.5 - (0.1 - ((0.0 + (-2.25 - 3.14159)) * 0.0))) - 0.1))
                                if false {
                                    v6 = (v5 + v5)
                                }
                                _ = v108
                            } else {
                                if true {
                                    switch ("\n\t" + "hello") {
                                    case "a":
                                        v109 := []int{}
                                        _ = v109
                                    case "foo bar":
                                        for fi110 := 0; fi110 < 18; fi110++ {
                                            v111 := GenS7{F0: ("a" + v5), F1: 0, F2: v3}
                                            v3 = (^(v3))
                                            _ = fi110
                                            _ = v111
                                        }
                                    case "hello":
                                        if false {
                                            v112 := GenS9{F0: "a"}
                                            fmt.Println(v3)
                                            _ = v112
                                        } else {
                                            for fi113 := 0; fi113 < 12; fi113++ {
                                                if true {
                                                    v1 -= (v1 * (v1 | (v1 & v1)))
                                                    switch v5 {
                                                    case "hello":
                                                        var v114 int = fi105
                                                        v2[v6] = (334 <= 2147483648)
                                                        fmt.Println(fi113, v3)
                                                        _ = v114
                                                    }
                                                    var v115 float32 = -2.25
                                                    _ = v115
                                                }
                                                var v116 float32 = 1e10
                                                _ = fi113
                                                _ = v116
                                            }
                                            var v117 uint16 = 42
                                            _ = v117
                                        }
                                    default:
                                        switch (fi105 - 4294967295) {
                                        case 128:
                                            v1 = (-(v1))
                                        default:
                                            if (((!(true || (-467 != 0))) && (false && (!true))) && false) {
                                                for fi118 := 0; fi118 < 6; fi118++ {
                                                    fmt.Println(fi118)
                                                    _ = fi118
                                                }
                                                var v119 int = fi105
                                                if false {
                                                    if false {
                                                        var v120 string = ((v5 + "") + ("a" + "foo bar"))
                                                        var v121 uint16 = 380
                                                        for fi122 := 0; fi122 < 2; fi122++ {
                                                            v123 := GenS4{F0: (v3 >> ((2) & 63)), F1: 0}
                                                            if (0 != (^(v1))) {
                                                                v2[v5] = false
                                                                var v124 int8 = 1
                                                                _ = v124
                                                            }
                                                            _ = fi122
                                                            _ = v123
                                                        }
                                                        _ = v120
                                                        _ = v121
                                                    } else {
                                                        for fi125 := 0; fi125 < 10; fi125++ {
                                                            var v126 float64 = 0.0
                                                            v6 = v6
                                                            _ = fi125
                                                            _ = v126
                                                        }
                                                        v2[v5] = (!true)
                                                    }
                                                    if true {
                                                        v1 = (v1 * v1)
                                                        if false {
                                                            v2[v6] = true
                                                            var v127 int8 = 0
                                                            _ = v127
                                                        }
                                                        v119 = -2147483648
                                                    } else {
                                                        var v128 int16 = 32767
                                                        _ = v128
                                                    }
                                                    var v129 uint8 = (v3 << ((42) & 63))
                                                    _ = v129
                                                }
                                                _ = v119
                                            }
                                        }
                                    }
                                    if true {
                                        v130 := GenS2{F0: -60, F1: fi105}
                                        _ = v130
                                    }
                                    if (!(true && ("hello" >= v5))) {
                                        for fi131 := 0; fi131 < 2; fi131++ {
                                            if false {
                                                if false {
                                                    if false {
                                                        var v132 float32 = -2.25
                                                        var v133 float64 = 0.0
                                                        v134 := GenS10{F0: v1, F1: 2}
                                                        _ = v132
                                                        _ = v133
                                                        _ = v134
                                                    }
                                                    var v135 int = (fi105 + (fi131 + fi131))
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v5 = ("\n\t" + v5)
                                                    }
                                                    _ = v135
                                                } else {
                                                    v136 := []int{(fi105 / (((fi105 | fi131)) | 1))}
                                                    var v137 uint = 0
                                                    _ = v136
                                                    _ = v137
                                                }
                                                v5 = (v5 + v5)
                                                var v138 int = (fi131 - (-(fi105)))
                                                if false {
                                                    if false {
                                                        if (1 > 100) {
                                                            v1 -= v1
                                                            var v139 float64 = ((0.1 + ((((1.5 / (0.0 + ((0.1 * ((0.0 - 3.14159) + 1.5)) - 3.14159))) * 0.1) + (3.14159 * -2.25)) * 3.14159)) + float64(1.5))
                                                            _ = v139
                                                        } else {
                                                            _ = len(v2)
                                                            if (!true) {
                                                                fmt.Println(fi131)
                                                                if (-108 == 0) {
                                                                    var v140 int32 = -1
                                                                    if false {
                                                                        v141 := []int{9223372036854775807}
                                                                        v138 -= (fi131 << ((635) & 63))
                                                                        fmt.Println(fi131, fi131, v5)
                                                                        _ = v141
                                                                    }
                                                                    if ((^(v1)) <= (v1 * (v1 % ((v1) | 1)))) {
                                                                        var v142 bool = (!true)
                                                                        v143 := GenS2{F0: 18, F1: fi105}
                                                                        _ = v142
                                                                        _ = v143
                                                                    } else {
                                                                        v3 = v3
                                                                        _ = len(v2)
                                                                        var v144 uint = 42
                                                                        _ = v144
                                                                    }
                                                                    _ = v140
                                                                }
                                                                switch v6 {
                                                                case "a":
                                                                    v5 = "\n\t"
                                                                default:
                                                                    v6 = v6
                                                                }
                                                            } else {
                                                                switch 9223372036854775807 {
                                                                case 640:
                                                                    if true {
                                                                        if false {
                                                                            var v145 string = v6
                                                                            if (4294967295 == 0) {
                                                                                if (false && false) {
                                                                                    v146 := []int{v138, (^(v138)), -682, -1}
                                                                                    _ = v146
                                                                                }
                                                                                v3 += 1
                                                                                var v147 uint16 = 741
                                                                                _ = v147
                                                                            } else {
                                                                                switch v5 {
                                                                                case "a":
                                                                                    var v148 int = int(v3)
                                                                                    _ = v148
                                                                                }
                                                                                fmt.Println(fi131, v3, fi131)
                                                                                switch (fi131 % ((v138) | 1)) {
                                                                                case 521:
                                                                                    v149 := GenS1{F0: 0}
                                                                                    _ = v149
                                                                                }
                                                                            }
                                                                            _ = v145
                                                                        }
                                                                    }
                                                                case -2147483648:
                                                                    var v150 int16 = 1
                                                                    _ = v150
                                                                }
                                                                fmt.Println(v1)
                                                                v2[v6] = ((false && (2147483648 == 0)) || (("a" >= v6) && (!true)))
                                                            }
                                                        }
                                                        v151 := GenS5{F0: fi105, F1: false}
                                                        _ = v151
                                                    } else {
                                                        var v152 bool = (!(!(!(false || true))))
                                                        v153 := GenS11{F0: 2, F1: fi105, F2: 1}
                                                        var v154 float64 = (1e10 * (-2.25 + float64(-2.25)))
                                                        _ = v152
                                                        _ = v153
                                                        _ = v154
                                                    }
                                                    v2[v6] = ((1e10 + 0.0) > (((3.14159 + (3.14159 * 1e10)) - 0.1) / 1.0))
                                                } else {
                                                    if true {
                                                        if true {
                                                            var v155 uint16 = uint16(v3)
                                                            _ = v155
                                                        }
                                                        if (-128 >= -12) {
                                                            v156 := GenS7{F0: (v6 + ((((("a" + v6) + "") + v5) + (v5 + (("\n\t" + v6) + "\n\t"))) + ((v6 + (("\n\t" + (v5 + "foo bar")) + (v6 + (v6 + v5)))) + "\n\t"))), F1: 0, F2: (v3 + (v3 + (v3 * v3)))}
                                                            var v157 float64 = -2.25
                                                            delete(v2, "foo bar")
                                                            _ = v156
                                                            _ = v157
                                                        }
                                                        switch v5 {
                                                        case "a":
                                                            v6 = v5
                                                        case "":
                                                            switch v6 {
                                                            case "hello":
                                                                if true {
                                                                    v158 := GenS11{F0: 985, F1: v138, F2: 1}
                                                                    v6 = "hello"
                                                                    _ = v158
                                                                }
                                                            default:
                                                                var v159 int64 = v1
                                                                _ = v159
                                                            }
                                                        default:
                                                            switch fi105 {
                                                            case -566:
                                                                if false {
                                                                    _ = len(v2)
                                                                    var v160 uint = 58
                                                                    var v161 uint8 = v3
                                                                    _ = v160
                                                                    _ = v161
                                                                } else {
                                                                    v138 -= (v138 % ((fi131) | 1))
                                                                    switch fi105 {
                                                                    case -584:
                                                                        var v162 uint64 = 18446744073709551615
                                                                        _ = v162
                                                                    case 607:
                                                                        var v163 int8 = -1
                                                                        _ = v163
                                                                    case 147:
                                                                        switch (v6 + v6) {
                                                                        case "foo bar":
                                                                            fmt.Println(fi105)
                                                                        case "":
                                                                            v164 := GenS7{F0: ("" + v6), F1: 4294967295, F2: 1}
                                                                            _ = v164
                                                                        case "hello":
                                                                            v1 -= v1
                                                                        }
                                                                    default:
                                                                        v1 -= v1
                                                                    }
                                                                }
                                                            case -2147483648:
                                                                v1 = v1
                                                            default:
                                                                v165 := GenS10{F0: (v1 ^ v1), F1: 2}
                                                                fmt.Println(v5)
                                                                v166 := GenS12{F0: 42, F1: v6, F2: -1}
                                                                _ = v165
                                                                _ = v166
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = v138
                                            }
                                            var v167 int = (fi105 * (^(fi105)))
                                            _ = fi131
                                            _ = v167
                                        }
                                    } else {
                                        for fi168 := 0; fi168 < 3; fi168++ {
                                            switch v6 {
                                            case "hello":
                                                fmt.Println(fi168)
                                            default:
                                                var v169 uint64 = 18446744073709551615
                                                _ = v169
                                            }
                                            var v170 int = fi105
                                            var v171 string = (v6 + v6)
                                            _ = fi168
                                            _ = v170
                                            _ = v171
                                        }
                                        var v172 float32 = -2.25
                                        _ = v172
                                    }
                                }
                                v173 := GenS4{F0: v3, F1: 0}
                                if true {
                                    switch fi105 {
                                    case 713:
                                        v1 += v1
                                    default:
                                        fmt.Println(v3, fi105, fi105)
                                    }
                                } else {
                                    if false {
                                        v3 = 0
                                        if true {
                                            v174 := GenS7{F0: ((v5 + (v6 + v6)) + "hello"), F1: 0, F2: (v3 << ((483) & 63))}
                                            var v175 int16 = 0
                                            var v176 string = ("foo bar" + v5)
                                            _ = v174
                                            _ = v175
                                            _ = v176
                                        } else {
                                            v6 = v5
                                        }
                                    } else {
                                        var v177 int16 = 358
                                        _ = v177
                                    }
                                    var v178 int64 = v1
                                    if true {
                                        switch -784 {
                                        case 32768:
                                            var v179 int64 = v1
                                            _ = v179
                                        default:
                                            var v180 int64 = v1
                                            _ = v180
                                        }
                                    } else {
                                        v5 = ("\n\t" + (v6 + v5))
                                        v181 := GenS6{F0: "hello", F1: (((v178 * v1) == v1) && true), F2: -1}
                                        var v182 int = 128
                                        _ = v181
                                        _ = v182
                                    }
                                    _ = v178
                                }
                                _ = v173
                            }
                            fmt.Println(v3)
                            fmt.Println(v6)
                        }
                    case 42:
                        switch 30 {
                        case -9223372036854775808:
                            v5 = v6
                        }
                    default:
                        for fi183 := 0; fi183 < 8; fi183++ {
                            v3 = (v3 - (^(v3)))
                            _ = fi183
                        }
                    }
                    _ = fi105
                }
                var v184 bool = (false || false)
                switch (v5 + v6) {
                case "foo bar":
                    if (-355 > v1) {
                        v185 := []int{-278, int(v3), 616}
                        if (!(true && v184)) {
                            switch (v5 + v6) {
                            case "":
                                v5 = (v5 + ((("foo bar" + "hello") + (v5 + ("\n\t" + "foo bar"))) + ""))
                            case "hello":
                                if (true || (-326 >= -2147483648)) {
                                    v186 := GenS12{F0: 110, F1: "a", F2: -1}
                                    _ = v186
                                } else {
                                    fmt.Println(v184, v5)
                                    v184 = v184
                                    v184 = true
                                }
                            case "foo bar":
                                fmt.Println(v6, v6, v6)
                            }
                        }
                        for fi187 := 0; fi187 < 9; fi187++ {
                            switch (fi187 << ((449) & 63)) {
                            case 32768:
                                var v188 int = (fi187 - fi187)
                                _ = v188
                            case -2147483648:
                                var v189 uint64 = 277
                                _ = v189
                            case 42:
                                var v190 uint8 = v3
                                _ = v190
                            default:
                                fmt.Println(v6, fi187)
                            }
                            _ = fi187
                        }
                        _ = v185
                    }
                default:
                    var v191 int = 1
                    _ = v191
                }
                _ = v184
            } else {
                if true {
                    if false {
                        v192 := GenS11{F0: 0, F1: -2147483648, F2: 1}
                        _ = v192
                    } else {
                        switch 568 {
                        case 1:
                            if (!false) {
                                fmt.Println(v1)
                            } else {
                                fmt.Println(v6, v3, v1)
                            }
                        case 32768:
                            if (true || (4294967295 >= 65535)) {
                                switch 42 {
                                case 2147483647:
                                    switch (v6 + "a") {
                                    case "\n\t":
                                        var v193 float64 = ((1.5 + -2.25) / float64((-2.25 + -2.25)))
                                        _ = v193
                                    case "":
                                        if true {
                                            _ = len(v2)
                                        }
                                    default:
                                        v194 := GenS13{F0: -683}
                                        _ = v194
                                    }
                                case -1:
                                    for fi195 := 0; fi195 < 2; fi195++ {
                                        if false {
                                            v5 = "foo bar"
                                            if false {
                                                v196 := GenS5{F0: (^(fi195)), F1: true}
                                                _ = v196
                                            } else {
                                                for fi197 := 0; fi197 < 8; fi197++ {
                                                    var v198 uint16 = 2
                                                    if false {
                                                        var v199 int = (^(fi195))
                                                        _ = v199
                                                    }
                                                    var v200 int32 = -2147483648
                                                    _ = fi197
                                                    _ = v198
                                                    _ = v200
                                                }
                                            }
                                        } else {
                                            var v201 int8 = 42
                                            _ = v201
                                        }
                                        _ = fi195
                                    }
                                case 127:
                                    v3 += (v3 / ((v3) | 1))
                                }
                                switch -1 {
                                case 0:
                                    v6 = "foo bar"
                                case 19:
                                    var v202 int16 = -32768
                                    _ = v202
                                case -2147483648:
                                    _ = len(v2)
                                default:
                                    fmt.Println(v5, v5)
                                }
                                delete(v2, "foo bar")
                            }
                        case 4294967295:
                            if (true || ((0 == 42) || false)) {
                                var v203 int = 32767
                                var v204 uint64 = 2
                                v205 := []int{(v203 << ((uint(v3)) & 63)), 1, (^(v203)), (-(v203))}
                                _ = v203
                                _ = v204
                                _ = v205
                            }
                        }
                        var v206 uint8 = 127
                        _ = v206
                    }
                    var v207 int64 = 1
                    var v208 int = 127
                    _ = v207
                    _ = v208
                } else {
                    var v209 uint32 = 358
                    _ = len(v2)
                    switch 127 {
                    case 956:
                        var v210 int16 = -235
                        _ = v210
                    case 1:
                        v211 := []int{-1, 4294967295, -912, 128}
                        _ = v211
                    }
                    _ = v209
                }
                if ((65535 > 42) || false) {
                    v5 = v6
                    for fi212 := 0; fi212 < 5; fi212++ {
                        var v213 uint16 = 933
                        v214 := []int{fi212}
                        var v215 bool = (v5 <= "\n\t")
                        _ = fi212
                        _ = v213
                        _ = v214
                        _ = v215
                    }
                    switch -1 {
                    case 9223372036854775807:
                        fmt.Println(v6)
                    case 127:
                        fmt.Println(v3, v6)
                    case 32768:
                        var v216 int8 = 100
                        _ = v216
                    }
                }
                v217 := []int{}
                _ = v217
            }
        case "":
            switch int(v1) {
            case 2147483647:
                var v218 bool = false
                _ = v218
            case 32768:
                delete(v2, (v5 + "hello"))
            case 32767:
                fmt.Println(v3, v3, v3)
            }
        case "hello":
            var v219 uint8 = 0
            _ = v219
        default:
            if false {
                var v220 uint = 18446744073709551615
                switch ((v5 + v5) + "a") {
                case "\n\t":
                    for fi221 := 0; fi221 < 6; fi221++ {
                        v222 := []int{(fi221 >> ((v220) & 63))}
                        _ = fi221
                        _ = v222
                    }
                }
                _ = v220
            } else {
                v2[v5] = true
                switch -2147483648 {
                case -768:
                    v2[("" + v6)] = (-2.25 <= 1e10)
                case 821:
                    var v223 uint8 = v3
                    _ = v223
                }
                switch -293 {
                case 425:
                    if (127 == -121) {
                        if false {
                            for fi224 := 0; fi224 < 1; fi224++ {
                                v225 := GenS14{F0: 0}
                                v226 := GenS15{F0: 100, F1: 101}
                                _ = fi224
                                _ = v225
                                _ = v226
                            }
                            var v227 float32 = float32((((float64(-2.25) * (1.5 - -2.25)) + 1.5) - 0.0))
                            if (42 >= 293) {
                                var v228 uint64 = 42
                                for fi229 := 0; fi229 < 13; fi229++ {
                                    var v230 uint16 = 255
                                    _ = fi229
                                    _ = v230
                                }
                                _ = v228
                            } else {
                                if false {
                                    if true {
                                        var v231 uint64 = 2
                                        switch "hello" {
                                        case "\n\t":
                                            if (false || false) {
                                                switch v5 {
                                                case "":
                                                    switch -572 {
                                                    case 0:
                                                        if (!true) {
                                                            for fi232 := 0; fi232 < 9; fi232++ {
                                                                v231 = v231
                                                                var v233 uint = 65535
                                                                var v234 int = -1
                                                                _ = fi232
                                                                _ = v233
                                                                _ = v234
                                                            }
                                                        }
                                                    case 1:
                                                        if (false && false) {
                                                            v235 := []int{128, 2147483647, 4294967295}
                                                            fmt.Println(v6, v6)
                                                            for fi236 := 0; fi236 < 10; fi236++ {
                                                                _ = len(v235)
                                                                _ = fi236
                                                            }
                                                            _ = v235
                                                        } else {
                                                            if (((v231 & v231) <= 4294967295) && true) {
                                                                var v237 int16 = -657
                                                                v238 := GenS9{F0: "\n\t"}
                                                                _ = v237
                                                                _ = v238
                                                            } else {
                                                                var v239 uint16 = 255
                                                                switch -150 {
                                                                case -224:
                                                                    switch -1 {
                                                                    case -384:
                                                                        v240 := GenS1{F0: (v239 + (v239 | (v239 / ((1) | 1))))}
                                                                        _ = v240
                                                                    case 733:
                                                                        v241 := []int{-360, -847}
                                                                        _ = v241
                                                                    case 4294967295:
                                                                        var v242 uint = 65535
                                                                        _ = v242
                                                                    default:
                                                                        if false {
                                                                            for fi243 := 0; fi243 < 7; fi243++ {
                                                                                if (true || (534 <= v239)) {
                                                                                    var v244 uint32 = 2
                                                                                    for fi245 := 0; fi245 < 4; fi245++ {
                                                                                        v227 = v227
                                                                                        var v246 string = ""
                                                                                        if ((((-2.25 * float64(v227)) + (1.5 / (1.5 * 3.14159))) + 1e10) <= 0.1) {
                                                                                            var v247 int16 = 1000
                                                                                            var v248 uint16 = v239
                                                                                            _ = v247
                                                                                            _ = v248
                                                                                        }
                                                                                        _ = fi245
                                                                                        _ = v246
                                                                                    }
                                                                                    v6 = ("" + v5)
                                                                                    _ = v244
                                                                                } else {
                                                                                    v1 += v1
                                                                                }
                                                                                for fi249 := 0; fi249 < 5; fi249++ {
                                                                                    v250 := GenS13{F0: 9223372036854775807}
                                                                                    v251 := []int{fi243, -391, (fi243 + fi249)}
                                                                                    v5 = v6
                                                                                    _ = fi249
                                                                                    _ = v250
                                                                                    _ = v251
                                                                                }
                                                                                _ = fi243
                                                                            }
                                                                        } else {
                                                                            v252 := []int{-29, 127, 42, 4294967295}
                                                                            v253 := GenS6{F0: "\n\t", F1: (true || (887 != 933)), F2: 648}
                                                                            v254 := v2[("a" + v6)]
                                                                            _ = v252
                                                                            _ = v253
                                                                            _ = v254
                                                                        }
                                                                    }
                                                                case -38:
                                                                    var v255 int16 = -947
                                                                    _ = v255
                                                                case 42:
                                                                    fmt.Println(v239, v5)
                                                                default:
                                                                    var v256 uint8 = v3
                                                                    _ = v256
                                                                }
                                                                _ = v239
                                                            }
                                                        }
                                                    }
                                                case "a":
                                                    var v257 bool = false
                                                    _ = v257
                                                }
                                                var v258 string = v5
                                                _ = v258
                                            }
                                        case "foo bar":
                                            v6 = v5
                                        case "hello":
                                            v259 := GenS12{F0: (v231 / (((v231 - v231)) | 1)), F1: (v6 + v5), F2: -1}
                                            _ = v259
                                        default:
                                            fmt.Println(v6, v3)
                                        }
                                        _ = v231
                                    }
                                }
                                var v260 float32 = (v227 / v227)
                                v261 := []int{1}
                                _ = v260
                                _ = v261
                            }
                            _ = v227
                        } else {
                            v262 := GenS4{F0: 105, F1: 0}
                            _ = v262
                        }
                        var v263 float64 = 3.14159
                        _ = v263
                    } else {
                        if true {
                            fmt.Println(v5, v5)
                            v264 := GenS9{F0: v5}
                            _ = v264
                        }
                        for fi265 := 0; fi265 < 15; fi265++ {
                            v266 := []int{2147483647, fi265, (fi265 + (fi265 * (fi265 * fi265))), (fi265 << ((2) & 63))}
                            _ = fi265
                            _ = v266
                        }
                        for fi267 := 0; fi267 < 14; fi267++ {
                            if true {
                                var v268 bool = (false || true)
                                for fi269 := 0; fi269 < 8; fi269++ {
                                    switch "hello" {
                                    case "\n\t":
                                        var v270 string = "hello"
                                        var v271 int = (^(fi267))
                                        var v272 float64 = -2.25
                                        _ = v270
                                        _ = v271
                                        _ = v272
                                    case "a":
                                        fmt.Println(fi267, v6)
                                    case "":
                                        v273 := GenS2{F0: -57, F1: 938}
                                        _ = v273
                                    default:
                                        v5 = (v6 + (v6 + v6))
                                    }
                                    v1 = v1
                                    if true {
                                        switch 287 {
                                        case 4294967295:
                                            fmt.Println(v5, v3)
                                        case 42:
                                            var v274 int = 786
                                            _ = v274
                                        case 1:
                                            v2[v5] = v268
                                        }
                                    }
                                    _ = fi269
                                }
                                _ = v268
                            } else {
                                v275 := GenS11{F0: 0, F1: -603, F2: 42}
                                var v276 string = ""
                                var v277 string = (v6 + (v6 + (v6 + v276)))
                                _ = v275
                                _ = v276
                                _ = v277
                            }
                            for fi278 := 0; fi278 < 8; fi278++ {
                                v1 = (-(v1))
                                var v279 int64 = (^(v1))
                                v279 -= (v1 & v279)
                                v279 += v279
                                v279 = (v279 - v279)
                                if (v6 >= "foo bar") {
                                    v280 := v2[v5]
                                    _ = v280
                                }
                                _ = fi278
                                _ = v279
                            }
                            _ = fi267
                        }
                    }
                default:
                    var v281 bool = (42 >= 42)
                    _ = v281
                }
            }
        }
        _ = v104
    } else {
        if true {
            for fi282 := 0; fi282 < 13; fi282++ {
                v1 = v1
                _ = fi282
            }
            fmt.Println(v3, v3)
            var v283 int8 = 100
            _ = v283
        } else {
            var v284 bool = ((float32((float64(((1.5 / (-2.25 * ((-2.25 / float32(0.1)) / 0.1))) - 1.5)) + float64(-2.25))) - float32(0.1)) < 0.1)
            _ = v284
        }
        var v285 int16 = -1
        fmt.Println(v285)
        _ = v285
    }
    v286 := GenS16{F0: 495}
    switch (v5 + (v5 + ((("hello" + v6) + "") + "foo bar"))) {
    case "":
        if true {
            var v287 float32 = 0.0
            _ = v287
        } else {
            fmt.Println(v1)
        }
    case "a":
        v6 = v6
    case "foo bar":
        var v288 int16 = 1000
        _ = v288
    default:
        switch 2147483647 {
        case 42:
            fmt.Println(v1, v5, v6)
        case 32768:
            v289 := []int{-2147483648}
            _ = v289
        case 127:
            for fi290 := 0; fi290 < 11; fi290++ {
                var v291 uint16 = 127
                var v292 uint = 42
                _ = fi290
                _ = v291
                _ = v292
            }
        default:
            var v293 int64 = v1
            _ = v293
        }
    }
    switch 32768 {
    case 4294967295:
        v6 = v6
    case 127:
        var v294 uint = 0
        _ = v294
    default:
        if true {
            v3 -= v3
            var v295 int16 = -32768
            v296 := GenS14{F0: 42}
            _ = v295
            _ = v296
        } else {
            v5 = "hello"
            v6 = v5
            for fi297 := 0; fi297 < 9; fi297++ {
                if (false && true) {
                    if false {
                        if true {
                            for fi298 := 0; fi298 < 20; fi298++ {
                                var v299 int16 = -32768
                                if (!false) {
                                    fmt.Println(v299, v1, fi297)
                                    v6 = v6
                                }
                                var v300 float64 = -2.25
                                switch 0 {
                                case -843:
                                    if (0.1 == (1e10 - -2.25)) {
                                        var v301 uint64 = 0
                                        v1 -= -9223372036854775808
                                        _ = v301
                                    } else {
                                        switch ((("hello" + ("hello" + "hello")) + ("" + (v6 + "hello"))) + v5) {
                                        case "\n\t":
                                            fmt.Println(v1)
                                        case "a":
                                            v302 := GenS12{F0: 4294967295, F1: "foo bar", F2: -2147483648}
                                            _ = v302
                                        }
                                        if false {
                                            v3 += v3
                                            var v303 uint16 = 0
                                            v299 = 32767
                                            _ = v303
                                        } else {
                                            if false {
                                                switch 343 {
                                                case -1:
                                                    delete(v2, v6)
                                                }
                                                v304 := GenS7{F0: "a", F1: 115, F2: v3}
                                                v2[v6] = ((fi297 / ((fi297) | 1)) <= 42)
                                                _ = v304
                                            }
                                            fmt.Println(v6, fi298)
                                            if false {
                                                fmt.Println(fi298, v3)
                                                if (fi298 < (^(fi297))) {
                                                    if ((v3 + v3) == (^(v3))) {
                                                        var v305 float64 = 3.14159
                                                        _ = v305
                                                    } else {
                                                        v2["hello"] = (-9223372036854775808 <= -9223372036854775808)
                                                    }
                                                    switch v6 {
                                                    case "hello":
                                                        var v306 int16 = v299
                                                        _ = v306
                                                    case "a":
                                                        v307 := []int{-2147483648}
                                                        _ = v307
                                                    }
                                                    var v308 int64 = v1
                                                    _ = v308
                                                } else {
                                                    _ = len(v2)
                                                }
                                                switch (fi297 - fi297) {
                                                case 32768:
                                                    var v309 int = (fi298 - fi298)
                                                    _ = v309
                                                case 4294967295:
                                                    v310 := GenS17{F0: 1}
                                                    _ = v310
                                                case 128:
                                                    v300 = (float64(0.0) * (float64(3.14159) + v300))
                                                }
                                            } else {
                                                switch fi298 {
                                                case 4294967295:
                                                    switch -932 {
                                                    case -172:
                                                        var v311 int32 = -2147483648
                                                        _ = v311
                                                    case 9223372036854775807:
                                                        switch (fi298 << ((255) & 63)) {
                                                        case 425:
                                                            delete(v2, ((("hello" + "") + (v6 + "a")) + ""))
                                                        case 1:
                                                            var v312 uint16 = 42
                                                            if (true && true) {
                                                                if true {
                                                                    if true {
                                                                        v3 = (^(v3))
                                                                        if false {
                                                                            if (-1 == -540) {
                                                                                v313 := []int{fi298}
                                                                                v3 += 59
                                                                                _ = v313
                                                                            } else {
                                                                                v5 = ((v5 + v6) + v5)
                                                                                switch (("foo bar" + v6) + v6) {
                                                                                case "a":
                                                                                    v314 := GenS17{F0: -2147483648}
                                                                                    _ = v314
                                                                                case "\n\t":
                                                                                    switch (-(fi297)) {
                                                                                    case 2147483647:
                                                                                        v5 = v5
                                                                                    case -1:
                                                                                        if true {
                                                                                            var v315 float64 = v300
                                                                                            _ = v315
                                                                                        } else {
                                                                                            switch "hello" {
                                                                                            case "foo bar":
                                                                                                var v316 uint64 = 475
                                                                                                _ = v316
                                                                                            }
                                                                                            switch fi297 {
                                                                                            case 380:
                                                                                                var v317 int = fi298
                                                                                                _ = v317
                                                                                            }
                                                                                            v5 = ("a" + v5)
                                                                                        }
                                                                                    case -534:
                                                                                        if false {
                                                                                            fmt.Println(v312)
                                                                                        } else {
                                                                                            switch fi298 {
                                                                                            case 9223372036854775807:
                                                                                                fmt.Println(v6, v5, fi298)
                                                                                            default:
                                                                                                v318 := GenS18{F0: v3, F1: 0, F2: -2147483648}
                                                                                                _ = v318
                                                                                            }
                                                                                            if (uint32(v312) >= 0) {
                                                                                                if ("hello" > ((("" + ((v6 + "\n\t") + ("a" + (((v5 + (v5 + (("" + (v5 + "hello")) + "a"))) + ("\n\t" + "hello")) + v6)))) + v5) + "a")) {
                                                                                                    fmt.Println(fi297, fi297)
                                                                                                    var v319 bool = false
                                                                                                    v1 += v1
                                                                                                    v320 := GenS7{F0: v6, F1: 4294967295, F2: v3}
                                                                                                    _ = v319
                                                                                                    _ = v320
                                                                                                } else {
                                                                                                    v321 := GenS10{F0: (^(v1)), F1: 18446744073709551615}
                                                                                                    _ = len(v2)
                                                                                                    _ = len(v2)
                                                                                                    _ = v321
                                                                                                }
                                                                                            }
                                                                                            var v322 int32 = 142
                                                                                            _ = v322
                                                                                        }
                                                                                    }
                                                                                case "":
                                                                                    v323 := []int{128, fi298, fi297, (fi297 + -479)}
                                                                                    _ = v323
                                                                                default:
                                                                                    var v324 int = fi297
                                                                                    _ = v324
                                                                                }
                                                                                v1 = (v1 >> ((42) & 63))
                                                                            }
                                                                        }
                                                                        if false {
                                                                            if (!false) {
                                                                                var v325 int8 = 102
                                                                                switch (^(fi298)) {
                                                                                case 4294967295:
                                                                                    switch 1 {
                                                                                    case 755:
                                                                                        if (3.14159 >= 0.1) {
                                                                                            v1 = 9223372036854775807
                                                                                        } else {
                                                                                            fmt.Println(v299, v3, fi297)
                                                                                            v5 = "\n\t"
                                                                                            v299 -= v299
                                                                                            switch fi298 {
                                                                                            case 734:
                                                                                                fmt.Println(v300)
                                                                                                if false {
                                                                                                    var v326 bool = false
                                                                                                    var v327 int = fi298
                                                                                                    var v328 uint8 = 127
                                                                                                    _ = v326
                                                                                                    _ = v327
                                                                                                    _ = v328
                                                                                                }
                                                                                                v329 := []int{}
                                                                                                _ = v329
                                                                                            case 898:
                                                                                                var v330 bool = true
                                                                                                _ = v330
                                                                                            case 4294967295:
                                                                                                if false {
                                                                                                    if false {
                                                                                                        fmt.Println(v5, v312, v325)
                                                                                                        var v331 float64 = 1e10
                                                                                                        var v332 int32 = -1
                                                                                                        _ = v331
                                                                                                        _ = v332
                                                                                                    } else {
                                                                                                        v6 = "hello"
                                                                                                        var v333 uint8 = v3
                                                                                                        _ = v333
                                                                                                    }
                                                                                                    v5 = (("hello" + v5) + (v6 + v6))
                                                                                                }
                                                                                            default:
                                                                                                v334 := GenS4{F0: 255, F1: 42}
                                                                                                _ = v334
                                                                                            }
                                                                                        }
                                                                                        if false {
                                                                                            var v335 uint = 255
                                                                                            v336 := []int{fi297, -835, fi298}
                                                                                            _ = v335
                                                                                            _ = v336
                                                                                        } else {
                                                                                            var v337 int64 = (v1 + v1)
                                                                                            if (1.5 < ((0.0 * 3.14159) * 0.0)) {
                                                                                                var v338 int16 = -396
                                                                                                _ = v338
                                                                                            }
                                                                                            _ = v337
                                                                                        }
                                                                                        v5 = v6
                                                                                        v339 := GenS5{F0: (fi298 / ((128) | 1)), F1: false}
                                                                                        var v340 float64 = v300
                                                                                        fmt.Println(fi297, v1, v299)
                                                                                        _ = v339
                                                                                        _ = v340
                                                                                    case 42:
                                                                                        switch ((v6 + "") + v6) {
                                                                                        case "a":
                                                                                            v341 := GenS13{F0: (v1 | v1)}
                                                                                            fmt.Println(v3, v325)
                                                                                            _ = v341
                                                                                        }
                                                                                    case 4294967295:
                                                                                        var v342 int8 = v325
                                                                                        _ = v342
                                                                                    default:
                                                                                        var v343 int32 = 42
                                                                                        _ = v343
                                                                                    }
                                                                                case 127:
                                                                                    var v344 uint = 1
                                                                                    _ = v344
                                                                                case 32768:
                                                                                    switch (("" + "hello") + ((("a" + v5) + "foo bar") + "")) {
                                                                                    case "":
                                                                                        v345 := GenS17{F0: -120}
                                                                                        _ = v345
                                                                                    case "\n\t":
                                                                                        var v346 uint8 = v3
                                                                                        _ = v346
                                                                                    case "a":
                                                                                        if ((false || false) || false) {
                                                                                            v347 := []int{int(v1), int(v1)}
                                                                                            _ = v347
                                                                                        } else {
                                                                                            v299 -= v299
                                                                                            v348 := []int{(fi297 ^ (fi297 + (-(fi297)))), 42, (^(fi297))}
                                                                                            _ = v348
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v325
                                                                            } else {
                                                                                switch "a" {
                                                                                case "":
                                                                                    v3 = v3
                                                                                case "hello":
                                                                                    if false {
                                                                                        var v349 int8 = 42
                                                                                        _ = v349
                                                                                    }
                                                                                case "\n\t":
                                                                                    if ("a" >= v5) {
                                                                                        switch (fi298 >> ((42) & 63)) {
                                                                                        case 4294967295:
                                                                                            switch "hello" {
                                                                                            case "a":
                                                                                                var v350 int32 = 1
                                                                                                _ = v350
                                                                                            case "\n\t":
                                                                                                v351 := GenS19{F0: fi297, F1: 2147483647, F2: 0}
                                                                                                _ = v351
                                                                                            case "hello":
                                                                                                if (705 == 42) {
                                                                                                    switch ((v6 + v5) + v6) {
                                                                                                    case "foo bar":
                                                                                                        var v352 int32 = 2147483647
                                                                                                        _ = v352
                                                                                                    default:
                                                                                                        var v353 bool = false
                                                                                                        _ = v353
                                                                                                    }
                                                                                                    switch int(v312) {
                                                                                                    case -2147483648:
                                                                                                        if (!true) {
                                                                                                            var v354 uint8 = (v3 | 177)
                                                                                                            var v355 uint64 = 42
                                                                                                            v356 := []int{int(v312), (fi298 << ((970) & 63)), int(v312)}
                                                                                                            var v357 bool = false
                                                                                                            var v358 uint64 = (v355 - 2)
                                                                                                            _ = v354
                                                                                                            _ = v355
                                                                                                            _ = v356
                                                                                                            _ = v357
                                                                                                            _ = v358
                                                                                                        }
                                                                                                    case 32767:
                                                                                                        v359 := GenS9{F0: ((v5 + v6) + v6)}
                                                                                                        _ = v359
                                                                                                    case 32768:
                                                                                                        v1 = (v1 + -265)
                                                                                                    }
                                                                                                    var v360 uint32 = 42
                                                                                                    _ = v360
                                                                                                } else {
                                                                                                    v300 = 1e10
                                                                                                    if (!(false || ((uint32(v3) > 65535) || (255 > 76)))) {
                                                                                                        if true {
                                                                                                            if (false || true) {
                                                                                                                delete(v2, v6)
                                                                                                            } else {
                                                                                                                fmt.Println(v6)
                                                                                                                var v361 uint = 2
                                                                                                                _ = v361
                                                                                                            }
                                                                                                        }
                                                                                                        if ("hello" > ((((v5 + v5) + "") + (v5 + v5)) + (((v6 + ("hello" + v5)) + "foo bar") + v5))) {
                                                                                                            var v362 uint8 = v3
                                                                                                            v2[v5] = false
                                                                                                            switch (fi297 | (fi298 * (fi298 * (-(fi297))))) {
                                                                                                            case 2147483647:
                                                                                                                v363 := v2[((v5 + v5) + (("a" + (v6 + "")) + v6))]
                                                                                                                _ = v363
                                                                                                            case 676:
                                                                                                                if true {
                                                                                                                    var v364 string = v5
                                                                                                                    switch ((("\n\t" + v5) + v5) + v364) {
                                                                                                                    case "":
                                                                                                                        if (true && (!(!(1 <= -128)))) {
                                                                                                                            v5 = (("\n\t" + (v364 + (v5 + "foo bar"))) + "")
                                                                                                                            if ((v300 + v300) > v300) {
                                                                                                                                fmt.Println(v6, v312, fi298)
                                                                                                                                var v365 int64 = v1
                                                                                                                                _ = v365
                                                                                                                            } else {
                                                                                                                                v366 := GenS14{F0: v299}
                                                                                                                                var v367 int16 = (v299 & v299)
                                                                                                                                v368 := GenS20{F0: 42, F1: (-(fi297))}
                                                                                                                                _ = v366
                                                                                                                                _ = v367
                                                                                                                                _ = v368
                                                                                                                            }
                                                                                                                        }
                                                                                                                    case "\n\t":
                                                                                                                        switch fi298 {
                                                                                                                        case 2147483647:
                                                                                                                            v369 := GenS21{F0: 1}
                                                                                                                            _ = v369
                                                                                                                        default:
                                                                                                                            var v370 string = v6
                                                                                                                            _ = v370
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v371 float64 = 0.1
                                                                                                                        _ = v371
                                                                                                                    }
                                                                                                                    v3 = (v3 & (^(v362)))
                                                                                                                    _ = v364
                                                                                                                }
                                                                                                            default:
                                                                                                                v2["foo bar"] = false
                                                                                                            }
                                                                                                            var v372 bool = (true || false)
                                                                                                            v373 := GenS17{F0: -319}
                                                                                                            _ = v362
                                                                                                            _ = v372
                                                                                                            _ = v373
                                                                                                        } else {
                                                                                                            fmt.Println(v6, v5, v1)
                                                                                                            if true {
                                                                                                                if false {
                                                                                                                    switch "\n\t" {
                                                                                                                    case "":
                                                                                                                        var v374 uint64 = 18446744073709551615
                                                                                                                        _ = v374
                                                                                                                    case "\n\t":
                                                                                                                        v5 = (v6 + ((v5 + "") + "a"))
                                                                                                                    case "foo bar":
                                                                                                                        v1 += (-(v1))
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    var v375 float64 = ((float64(-2.25) + v300) / v300)
                                                                                                                    switch fi298 {
                                                                                                                    case 162:
                                                                                                                        _ = len(v2)
                                                                                                                    default:
                                                                                                                        var v376 float32 = float32(v300)
                                                                                                                        _ = v376
                                                                                                                    }
                                                                                                                    delete(v2, "")
                                                                                                                    _ = v375
                                                                                                                }
                                                                                                                _ = len(v2)
                                                                                                                switch fi297 {
                                                                                                                case 128:
                                                                                                                    switch fi298 {
                                                                                                                    case 32767:
                                                                                                                        if true {
                                                                                                                            var v377 float32 = (3.14159 / (0.0 - 1.5))
                                                                                                                            if true {
                                                                                                                                v377 -= v377
                                                                                                                                var v378 int64 = 1
                                                                                                                                v378 = v1
                                                                                                                                _ = v378
                                                                                                                            }
                                                                                                                            _ = v377
                                                                                                                        } else {
                                                                                                                            v379 := GenS14{F0: v299}
                                                                                                                            v300 -= 3.14159
                                                                                                                            v1 -= (v1 % (((v1 * v1)) | 1))
                                                                                                                            _ = v379
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v380 int8 = 42
                                                                                                                        _ = v380
                                                                                                                    }
                                                                                                                default:
                                                                                                                    v312 += (v312 * (^(v312)))
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        v300 = float64(-2.25)
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                switch v6 {
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v3)
                                                                                                case "\n\t":
                                                                                                    fmt.Println(v299, v300, v5)
                                                                                                case "":
                                                                                                    v5 = ("\n\t" + "foo bar")
                                                                                                }
                                                                                            }
                                                                                        case -399:
                                                                                            delete(v2, "")
                                                                                        case -2147483648:
                                                                                            v381 := v2["\n\t"]
                                                                                            v382 := GenS22{F0: 4294967295, F1: 797, F2: -496}
                                                                                            _ = v381
                                                                                            _ = v382
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v383 := []int{}
                                                                                    _ = v383
                                                                                }
                                                                            }
                                                                            if true {
                                                                                var v384 int16 = v299
                                                                                var v385 uint16 = v312
                                                                                var v386 int64 = (v1 | v1)
                                                                                _ = v384
                                                                                _ = v385
                                                                                _ = v386
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    v3 = v3
                                                                    v387 := GenS23{F0: v1}
                                                                    v6 = v6
                                                                    fmt.Println(v6, fi298)
                                                                    _ = v387
                                                                }
                                                                v388 := v2["\n\t"]
                                                                v389 := []int{}
                                                                _ = v388
                                                                _ = v389
                                                            } else {
                                                                var v390 uint16 = 541
                                                                _ = v390
                                                            }
                                                            v5 = "hello"
                                                            v391 := []int{}
                                                            _ = v312
                                                            _ = v391
                                                        case 127:
                                                            if true {
                                                                var v392 bool = (v1 >= 9223372036854775807)
                                                                v393 := GenS6{F0: "\n\t", F1: ((fi297 - (^(fi298))) > fi297), F2: (-(v299))}
                                                                var v394 string = (v6 + "\n\t")
                                                                v395 := v2[("\n\t" + v6)]
                                                                _ = v392
                                                                _ = v393
                                                                _ = v394
                                                                _ = v395
                                                            }
                                                        }
                                                    case 2147483647:
                                                        v396 := GenS12{F0: 303, F1: v5, F2: 42}
                                                        _ = v396
                                                    }
                                                case 32768:
                                                    var v397 uint8 = v3
                                                    _ = v397
                                                default:
                                                    if false {
                                                        v398 := []int{fi298, (fi298 >> ((0) & 63)), fi298}
                                                        _ = v398
                                                    } else {
                                                        v300 += (v300 / 1.5)
                                                    }
                                                }
                                            }
                                            if (false && (int8(v299) == 127)) {
                                                var v399 int = (fi298 - -1)
                                                v6 = v5
                                                _ = v399
                                            } else {
                                                var v400 int = (fi298 << ((695) & 63))
                                                v5 = ("hello" + (v5 + "foo bar"))
                                                v401 := GenS19{F0: fi297, F1: -2147483648, F2: 32767}
                                                v5 = (("a" + (v6 + ((v6 + "hello") + v5))) + v5)
                                                _ = v400
                                                _ = v401
                                            }
                                            var v402 string = v5
                                            var v403 uint16 = 191
                                            _ = v402
                                            _ = v403
                                        }
                                        if ((v299 * (v299 << ((61) & 63))) != v299) {
                                            var v404 float64 = 0.1
                                            _ = v404
                                        } else {
                                            v405 := GenS24{F0: 667}
                                            var v406 uint16 = 474
                                            if (2147483647 >= 2147483647) {
                                                fmt.Println(v5, v299)
                                                v407 := GenS6{F0: "", F1: (!false), F2: (v299 | 1000)}
                                                v408 := GenS12{F0: 2, F1: v5, F2: -1}
                                                _ = v407
                                                _ = v408
                                            } else {
                                                var v409 uint32 = 2147483648
                                                _ = v409
                                            }
                                            _ = v405
                                            _ = v406
                                        }
                                        fmt.Println(v3)
                                    }
                                case 128:
                                    switch 4294967295 {
                                    case 9223372036854775807:
                                        v5 = "\n\t"
                                    }
                                case -410:
                                    v410 := GenS10{F0: (-(v1)), F1: 668}
                                    _ = v410
                                }
                                _ = fi298
                                _ = v299
                                _ = v300
                            }
                            v5 = (v5 + (v6 + (v6 + v5)))
                            v411 := []int{(fi297 ^ 4294967295)}
                            _ = v411
                        }
                    }
                    v6 = v6
                } else {
                    v412 := GenS11{F0: 124, F1: fi297, F2: -2147483648}
                    for fi413 := 0; fi413 < 14; fi413++ {
                        var v414 float32 = float32((1.5 + -2.25))
                        v415 := v2[v5]
                        _ = fi413
                        _ = v414
                        _ = v415
                    }
                    _ = v412
                }
                for fi416 := 0; fi416 < 8; fi416++ {
                    if (true && ((v3 <= 1) || (-1 > 127))) {
                        var v417 uint64 = 2
                        v5 = "\n\t"
                        var v418 uint8 = 0
                        _ = v417
                        _ = v418
                    } else {
                        var v419 string = v6
                        _ = v419
                    }
                    if (0.1 <= (0.1 / 3.14159)) {
                        v3 = (v3 & (v3 >> ((201) & 63)))
                        if (v5 == (v5 + v5)) {
                            if false {
                                v420 := GenS13{F0: (v1 >> ((599) & 63))}
                                _ = v420
                            }
                        }
                        if (!(0 < 65535)) {
                            if false {
                                fmt.Println(v3, v1)
                                var v421 int16 = -930
                                _ = v421
                            } else {
                                v5 = v6
                                var v422 int64 = (v1 - v1)
                                v423 := []int{fi416}
                                _ = v422
                                _ = v423
                            }
                            v1 = v1
                        }
                        v3 += v3
                        v1 -= v1
                        var v424 float32 = ((1.5 - (0.1 * 0.1)) / 1.0)
                        _ = v424
                    } else {
                        var v425 float32 = -2.25
                        _ = v425
                    }
                    var v426 float64 = (float64((float32(((0.1 - -2.25) + ((-2.25 - (0.0 + (1e10 * 0.1))) / 1.5))) * (float32(((float64(1e10) / (1e10 - (0.0 - 0.1))) - -2.25)) / 1.0))) + 0.1)
                    _ = fi416
                    _ = v426
                }
                _ = fi297
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v286
    _ = v3
    _ = v5
    _ = v6
}
