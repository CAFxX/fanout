package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 bool; F2 int32 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32; F1 int8; F2 int8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int64; F1 bool; F2 int8 }
type GenS13 struct { F0 bool; F1 int32 }
type GenS14 struct { F0 string; F1 uint64; F2 int32 }
type GenS15 struct { F0 int; F1 string }
type GenS16 struct { F0 int64; F1 uint16; F2 uint }
type GenS17 struct { F0 string }
type GenS18 struct { F0 uint8; F1 int32 }
type GenS19 struct { F0 int16 }
type GenS20 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS21 struct { F0 uint16; F1 uint32 }
type GenS22 struct { F0 uint32; F1 int16; F2 int8 }
type GenS23 struct { F0 uint }
type GenS24 struct { F0 int8; F1 bool }
type GenS25 struct { F0 uint; F1 int }
type GenS26 struct { F0 int8; F1 uint64 }
type GenS27 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS28 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS29 struct { F0 int; F1 uint32; F2 int64 }
type GenS30 struct { F0 uint; F1 uint }
type GenS31 struct { F0 uint32 }
type GenS32 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS33 struct { F0 uint32 }
type GenS34 struct { F0 uint8; F1 uint8 }
type GenS35 struct { F0 string }
type GenS36 struct { F0 int; F1 int; F2 string }
type GenS37 struct { F0 uint64 }
type GenS38 struct { F0 int16 }
type GenS39 struct { F0 int16; F1 string; F2 uint16 }
type GenS40 struct { F0 bool; F1 int16; F2 int8 }
type GenS41 struct { F0 string }
type GenS42 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint32) int64 {
    p1 += (^(p1))
    var v7 uint64 = uint64(p1)
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 string = "foo bar"
        _ = fi8
        _ = v9
    }
    var v10 int32 = -515
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 32767
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    var v5 uint = 255
    var v6 uint = v5
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            v2 = 0
            var v13 bool = (fi12 == fi12)
            _ = fi12
            _ = v13
        }
        v6 = v6
        var v14 float32 = (0.0 / 1.0)
        _ = fi11
        _ = v14
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := GenS1{F0: 42, F1: "\n\t"}
        if (v5 <= 502) {
            v5 -= (v5 / ((688) | 1))
            var v17 uint64 = 4294967295
            _ = v17
        }
        var v18 int32 = 0
        _ = fi15
        _ = v16
        _ = v18
    }
    var v19 bool = true
    if (!v19) {
        v2 -= v2
        v5 -= v6
        for fi20 := 0; fi20 < 10; fi20++ {
            for fi21 := 0; fi21 < 12; fi21++ {
                v2 -= (v2 << ((v6) & 63))
                switch "\n\t" {
                case "":
                    if v19 {
                        var v22 string = "a"
                        _ = v22
                    } else {
                        switch (-(fi20)) {
                        case 42:
                            v23 := GenS2{F0: 0}
                            _ = v23
                        case -816:
                            var v24 uint8 = uint8(v2)
                            _ = v24
                        case 533:
                            v25 := []int{9223372036854775807, fi20, (fi21 * 32767), (^(fi21))}
                            _ = v25
                        }
                    }
                case "a":
                    v2 = (^(v2))
                case "hello":
                    v26 := []int{(fi20 ^ (fi21 & 265)), -654, 0}
                    _ = v26
                }
                _ = fi21
            }
            if true {
                var v27 string = "\n\t"
                if (!((v27 <= (v27 + (("a" + ("foo bar" + v27)) + v27))) || ((4294967295 >= uint(v2)) || (1 >= 42)))) {
                    switch (fi20 * (fi20 - fi20)) {
                    case 1:
                        for fi28 := 0; fi28 < 5; fi28++ {
                            v27 = v27
                            _ = fi28
                        }
                    case 128:
                        var v29 uint = v5
                        _ = v29
                    default:
                        var v30 uint16 = uint16(v5)
                        _ = v30
                    }
                    var v31 string = "a"
                    _ = v31
                } else {
                    fmt.Println(v27, v27, v19)
                }
                v2 = (v2 << ((510) & 63))
                _ = v27
            } else {
                fmt.Println(v19, fi20)
            }
            _ = fi20
        }
    } else {
        if ((0.1 * (1e10 - 1.5)) == ((0.1 / 1.0) / 1.0)) {
            v2 = v2
        }
        for fi32 := 0; fi32 < 3; fi32++ {
            v33 := GenS2{F0: 32767}
            if v19 {
                v34 := []int{2147483647, 4294967295, (^(fi32))}
                _ = v34
            } else {
                v35 := GenS1{F0: -63, F1: "\n\t"}
                if (v19 || (773 > v6)) {
                    v6 = v6
                    var v36 uint32 = 813
                    _ = v36
                }
                _ = v35
            }
            switch int(v2) {
            case 9223372036854775807:
                if (v19 && v19) {
                    v5 = 0
                    v2 = (^(v2))
                }
            default:
                v6 -= 852
            }
            _ = fi32
            _ = v33
        }
        if v19 {
            _ = len(v3)
            if (1 > -9223372036854775808) {
                for fi37 := 0; fi37 < 5; fi37++ {
                    v3["foo bar"] = (fi37 - -805)
                    if false {
                        if (true && false) {
                            v38 := []int{fi37, fi37, fi37, int(v6)}
                            if false {
                                v39 := GenS3{F0: 127}
                                var v40 int32 = -1
                                var v41 int16 = -244
                                _ = v39
                                _ = v40
                                _ = v41
                            } else {
                                if (!((1 <= 127) && (v19 && v19))) {
                                    v5 = 30
                                    switch (fi37 & (-(fi37))) {
                                    case 4294967295:
                                        v4[int(v2)] = (!(v19 && ((v2 % ((2) | 1)) == 783)))
                                    }
                                }
                            }
                            _ = v38
                        } else {
                            if false {
                                v19 = ("" != "")
                                v42 := GenS3{F0: 1}
                                v43 := GenS2{F0: 42}
                                _ = v42
                                _ = v43
                            } else {
                                for fi44 := 0; fi44 < 7; fi44++ {
                                    switch 42 {
                                    case -814:
                                        fmt.Println(v19)
                                    case 0:
                                        v45 := GenS4{F0: -48, F1: (true || true), F2: -1}
                                        _ = v45
                                    default:
                                        v5 = v6
                                    }
                                    _ = fi44
                                }
                                switch "foo bar" {
                                case "":
                                    v5 -= (v5 | 426)
                                case "foo bar":
                                    var v46 int8 = 100
                                    _ = v46
                                case "hello":
                                    for fi47 := 0; fi47 < 19; fi47++ {
                                        var v48 uint32 = 0
                                        fmt.Println(v19, fi47)
                                        if (127 == -128) {
                                            var v49 uint8 = 0
                                            var v50 int64 = -881
                                            switch (fi37 / ((fi37) | 1)) {
                                            case 9223372036854775807:
                                                var v51 float32 = float32(0.0)
                                                _ = v51
                                            case -1:
                                                if v19 {
                                                    v52 := GenS1{F0: -117, F1: "\n\t"}
                                                    v48 = v48
                                                    _ = v52
                                                } else {
                                                    v50 -= genFn0(v49, 2147483647)
                                                }
                                            default:
                                                v53 := v1["foo bar"]
                                                _ = v53
                                            }
                                            _ = v49
                                            _ = v50
                                        } else {
                                            if (255 >= 4294967295) {
                                                switch "hello" {
                                                case "":
                                                    fmt.Println(fi47, v2)
                                                case "\n\t":
                                                    switch ("" + "hello") {
                                                    case "a":
                                                        v54 := []int{(fi47 - fi37), int(v2), fi37, (fi37 % (((fi37 % (((fi47 << (((v6 + 1)) & 63))) | 1))) | 1))}
                                                        _ = v54
                                                    case "hello":
                                                        if (127 < 127) {
                                                            fmt.Println(v48)
                                                            var v55 int32 = 0
                                                            _ = v55
                                                        }
                                                    case "foo bar":
                                                        v56 := []int{int(v5), 32767, -766}
                                                        _ = v56
                                                    }
                                                }
                                            }
                                            var v57 bool = v19
                                            var v58 uint32 = v48
                                            _ = v57
                                            _ = v58
                                        }
                                        _ = fi47
                                        _ = v48
                                    }
                                default:
                                    if true {
                                        switch "foo bar" {
                                        case "hello":
                                            var v59 uint8 = 176
                                            _ = v59
                                        case "foo bar":
                                            var v60 int16 = 0
                                            _ = v60
                                        default:
                                            v61 := []int{fi37, fi37, (fi37 >> ((v6) & 63))}
                                            _ = v61
                                        }
                                    }
                                }
                            }
                            var v62 uint8 = uint8(v2)
                            for fi63 := 0; fi63 < 17; fi63++ {
                                v64 := GenS1{F0: 100, F1: ""}
                                switch "" {
                                case "\n\t":
                                    fmt.Println(v5, v19, v5)
                                default:
                                    v19 = (true || (v19 || v19))
                                }
                                switch fi37 {
                                case 1:
                                    v19 = v19
                                case 32767:
                                    if v19 {
                                        v65 := GenS5{F0: fi37, F1: v19}
                                        var v66 uint = v6
                                        switch fi63 {
                                        case 4294967295:
                                            switch fi37 {
                                            case 32767:
                                                v6 = 374
                                            case 1:
                                                v62 += v62
                                            case 9223372036854775807:
                                                var v67 string = "a"
                                                _ = v67
                                            default:
                                                v5 += (v66 << ((613) & 63))
                                            }
                                        }
                                        _ = v65
                                        _ = v66
                                    }
                                default:
                                    if false {
                                        fmt.Println(v5, v5, v19)
                                        if ((v5 & uint(v2)) <= (^(v6))) {
                                            switch (-(fi37)) {
                                            case 738:
                                                if false {
                                                    var v68 uint32 = 42
                                                    _ = v68
                                                } else {
                                                    v69 := []int{9223372036854775807}
                                                    if true {
                                                        fmt.Println(v2, v62)
                                                        if (!(42 <= 4294967295)) {
                                                            if (!v19) {
                                                                if (1.5 < (1e10 - (0.1 * ((float32((3.14159 / (0.1 - -2.25))) - 3.14159) + (1e10 - (0.1 - -2.25)))))) {
                                                                    v6 = 18446744073709551615
                                                                    var v70 uint8 = 128
                                                                    v62 = (v70 >> ((v6) & 63))
                                                                    _ = v70
                                                                } else {
                                                                    var v71 float32 = 0.0
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        v6 -= v5
                                                                    default:
                                                                        fmt.Println(v6)
                                                                    }
                                                                    _ = v71
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        v69 = append(v69, (fi63 / (((fi37 | fi37)) | 1)))
                                                        var v72 int32 = 797
                                                        _ = v72
                                                    }
                                                    _ = v69
                                                }
                                            case 653:
                                                v5 = uint(v2)
                                                v73 := GenS5{F0: 782, F1: true}
                                                var v74 float32 = -2.25
                                                _ = v73
                                                _ = v74
                                            default:
                                                if v19 {
                                                    if v19 {
                                                        v5 = uint(v62)
                                                        _ = len(v3)
                                                        var v75 uint32 = 517
                                                        switch "a" {
                                                        case "a":
                                                            fmt.Println(v62, v5)
                                                            var v76 float32 = 3.14159
                                                            var v77 float64 = 1e10
                                                            _ = v76
                                                            _ = v77
                                                        case "\n\t":
                                                            fmt.Println(v19, fi63)
                                                        case "hello":
                                                            var v78 float64 = 0.1
                                                            _ = v78
                                                        default:
                                                            var v79 int32 = 2147483647
                                                            _ = v79
                                                        }
                                                        var v80 uint32 = v75
                                                        _ = v75
                                                        _ = v80
                                                    } else {
                                                        v81 := GenS6{F0: int16(fi63)}
                                                        _ = v81
                                                    }
                                                }
                                            }
                                            v5 = (v5 / ((v5) | 1))
                                        }
                                        if false {
                                            if true {
                                                var v82 uint32 = 2147483647
                                                fmt.Println(v62, fi63, v19)
                                                fmt.Println(v6)
                                                _ = v82
                                            }
                                        }
                                        v6 -= (v6 ^ uint(v2))
                                    } else {
                                        v19 = (!(!false))
                                        _ = len(v1)
                                    }
                                }
                                _ = fi63
                                _ = v64
                            }
                            _ = v62
                        }
                    } else {
                        var v83 int16 = 1000
                        var v84 int64 = 42
                        _ = len(v1)
                        _ = v83
                        _ = v84
                    }
                    _ = fi37
                }
            }
            var v85 int = 2147483647
            _ = v85
        }
    }
    if ((0.0 >= 1.5) && v19) {
        v86 := GenS5{F0: 1, F1: (v19 || (!(!v19)))}
        var v87 uint8 = 0
        _ = v86
        _ = v87
    } else {
        for fi88 := 0; fi88 < 19; fi88++ {
            fmt.Println(v6, fi88, fi88)
            if (255 != 42) {
                var v89 float32 = (1e10 + 0.1)
                _ = v89
            } else {
                var v90 int = 830
                v91 := []int{(v90 * (fi88 | (fi88 >> (((v5 + v6)) & 63)))), v90, fi88, 42}
                var v92 uint64 = 683
                _ = v90
                _ = v91
                _ = v92
            }
            _ = fi88
        }
        v93 := GenS5{F0: -2147483648, F1: (896 != 42)}
        _ = v93
    }
    v6 = (v6 ^ v6)
    fmt.Println(v19, v5, v5)
    v5 = 4294967295
    for fi94 := 0; fi94 < 17; fi94++ {
        if ((2147483647 != 42) || (!(661 >= 2))) {
            switch "a" {
            case "\n\t":
                v19 = v19
            case "":
                v19 = true
            case "hello":
                if (!true) {
                    var v95 uint8 = 0
                    _ = v95
                } else {
                    var v96 int16 = -32768
                    _ = v96
                }
            default:
                for fi97 := 0; fi97 < 11; fi97++ {
                    var v98 uint = 65535
                    if (v5 != 0) {
                        fmt.Println(fi97)
                        if (32767 < 42) {
                            var v99 uint8 = 42
                            if (1.5 <= (0.1 * (-2.25 - float64(-2.25)))) {
                                if v19 {
                                    v19 = (808 <= 65535)
                                    if (fi97 == 812) {
                                        v100 := []int{}
                                        _ = v100
                                    }
                                }
                                var v101 uint64 = 18446744073709551615
                                _ = v101
                            }
                            _ = v99
                        } else {
                            if (((("\n\t" + "\n\t") + ("a" + "\n\t")) + ("a" + ("" + ("hello" + "a")))) != (("\n\t" + ("foo bar" + "foo bar")) + "foo bar")) {
                                v5 = 892
                                v5 -= (v6 | (v6 << (((v98 - (v5 ^ (v6 >> (((v6 | 876)) & 63))))) & 63)))
                                var v102 float64 = float64(3.14159)
                                v2 = v2
                                _ = v102
                            } else {
                                v19 = v19
                            }
                            var v103 int16 = 1000
                            v104 := GenS7{F0: "foo bar"}
                            v105 := []int{(fi94 % (((-(fi94))) | 1)), fi94, fi97}
                            _ = v103
                            _ = v104
                            _ = v105
                        }
                    }
                    _ = fi97
                    _ = v98
                }
            }
            if (v2 <= v2) {
                v5 += (^(v5))
                switch (fi94 + (fi94 & (fi94 + (fi94 % (((^(fi94))) | 1))))) {
                case 0:
                    for fi106 := 0; fi106 < 18; fi106++ {
                        if true {
                            fmt.Println(v2, v19)
                        }
                        v107 := GenS7{F0: ("\n\t" + "")}
                        v108 := GenS6{F0: 657}
                        _ = fi106
                        _ = v107
                        _ = v108
                    }
                case 4294967295:
                    if (!true) {
                        for fi109 := 0; fi109 < 12; fi109++ {
                            v110 := []int{(fi94 << (((v5 | (v6 >> (((v6 ^ 1)) & 63)))) & 63)), (fi109 / ((fi109) | 1)), fi109}
                            if true {
                                if (((3.14159 * float32(1e10)) > 0.0) && (-716 > -9223372036854775808)) {
                                    var v111 int8 = 1
                                    _ = v111
                                }
                                v112 := GenS6{F0: 245}
                                if false {
                                    if true {
                                        v113 := GenS1{F0: -128, F1: ""}
                                        var v114 bool = (128 > 0)
                                        var v115 uint16 = v2
                                        _ = v113
                                        _ = v114
                                        _ = v115
                                    } else {
                                        if true {
                                            var v116 string = ""
                                            _ = len(v110)
                                            v5 += 1
                                            _ = v116
                                        }
                                    }
                                }
                                _ = v112
                            } else {
                                v117 := GenS2{F0: 0}
                                if false {
                                    var v118 uint32 = 974
                                    var v119 int8 = 127
                                    v19 = (32767 < 32767)
                                    _ = v118
                                    _ = v119
                                }
                                if v19 {
                                    if (fi109 == 915) {
                                        if (-2147483648 <= 42) {
                                            if (((0.1 - (1.5 / 3.14159)) * float32(1e10)) > (1e10 * 0.1)) {
                                                v120 := GenS4{F0: 127, F1: v19, F2: 42}
                                                v121 := v3[("foo bar" + "")]
                                                _ = v120
                                                _ = v121
                                            }
                                            if v19 {
                                                switch (fi94 + 1) {
                                                case 640:
                                                    var v122 float64 = (0.1 - float64(0.0))
                                                    _ = v122
                                                }
                                                v4[fi94] = true
                                            } else {
                                                fmt.Println(v6, fi94)
                                                var v123 int16 = -1
                                                _ = v123
                                            }
                                        } else {
                                            if (v19 && (-804 > 42)) {
                                                var v124 uint32 = 257
                                                v110 = append(v110, fi109)
                                                if v19 {
                                                    if true {
                                                        var v125 int64 = 0
                                                        _ = v125
                                                    }
                                                    if false {
                                                        v5 += 42
                                                        if ((1.5 + (3.14159 / float32(0.1))) != (-2.25 + 0.1)) {
                                                            if ("\n\t" < ("a" + "foo bar")) {
                                                                v4[(fi109 * fi94)] = (-845 <= 2147483647)
                                                            }
                                                        }
                                                    } else {
                                                        v124 = v124
                                                        fmt.Println(v19)
                                                        if (false && (0 == v2)) {
                                                            v124 -= 622
                                                            var v126 uint64 = 1
                                                            var v127 uint16 = 486
                                                            v6 = (v6 + v6)
                                                            _ = v126
                                                            _ = v127
                                                        } else {
                                                            if false {
                                                                var v128 int32 = 0
                                                                var v129 uint = 495
                                                                _ = v128
                                                                _ = v129
                                                            }
                                                        }
                                                        v130 := GenS7{F0: "hello"}
                                                        _ = v130
                                                    }
                                                    if v19 {
                                                        if (!(!false)) {
                                                            if (-128 >= 100) {
                                                                if v19 {
                                                                    fmt.Println(v5, v6)
                                                                    var v131 float32 = float32((0.1 * 0.0))
                                                                    if (((false || false) && (!v19)) && (v5 > v5)) {
                                                                        v6 = uint(v2)
                                                                    } else {
                                                                        var v132 bool = v19
                                                                        _ = v132
                                                                    }
                                                                    _ = v131
                                                                } else {
                                                                    var v133 float32 = (0.1 + (0.1 / (1e10 - (1.5 / (0.0 + 3.14159)))))
                                                                    _ = v133
                                                                }
                                                                fmt.Println(v19)
                                                            } else {
                                                                switch ("" + "\n\t") {
                                                                case "hello":
                                                                    var v134 uint = 635
                                                                    _ = v134
                                                                case "a":
                                                                    if len(v110) > 0 {
                                                                        v110[0] = fi109
                                                                    }
                                                                    v2 = (v2 + v2)
                                                                }
                                                                v2 = 65535
                                                                v124 += v124
                                                            }
                                                            switch fi94 {
                                                            case -982:
                                                                v19 = (1 > 0)
                                                            case 32767:
                                                                v124 = v124
                                                            case 128:
                                                                var v135 int16 = 1000
                                                                _ = v135
                                                            default:
                                                                if ((false || v19) || (!(42 != 1000))) {
                                                                    var v136 bool = v19
                                                                    _ = v136
                                                                }
                                                            }
                                                        }
                                                        v19 = (v19 || v19)
                                                    }
                                                }
                                                v2 -= 42
                                                _ = len(v110)
                                                _ = v124
                                            } else {
                                                v137 := GenS8{F0: "\n\t"}
                                                fmt.Println(fi94, v19)
                                                _ = v137
                                            }
                                        }
                                        var v138 float64 = (((float64(((float32(1e10) * 0.1) / 1.0)) - (3.14159 / ((float64(float32(1e10)) + 1.5) - ((0.0 - -2.25) + 3.14159)))) * (1.5 + 3.14159)) * 0.0)
                                        switch "hello" {
                                        case "a":
                                            v139 := GenS6{F0: 0}
                                            _ = v139
                                        case "hello":
                                            if ((fi109 >> (((v5 & (v5 << (((v6 & v6)) & 63)))) & 63)) == 662) {
                                                if len(v110) >= 2 {
                                                    v110 = v110[1:len(v110)]
                                                }
                                                v19 = (1.5 < (float64(1.5) - 3.14159))
                                                switch 42 {
                                                case -702:
                                                    if true {
                                                        if v19 {
                                                            var v140 float32 = (-2.25 / (((0.0 * 1.5) + 0.1) + -2.25))
                                                            v2 = (v2 - v2)
                                                            _ = v140
                                                        } else {
                                                            var v141 float64 = ((0.1 - 0.1) - v138)
                                                            var v142 uint64 = 42
                                                            _ = v141
                                                            _ = v142
                                                        }
                                                        var v143 int16 = 1
                                                        _ = v143
                                                    }
                                                case 32768:
                                                    v5 = uint(v2)
                                                default:
                                                    if false {
                                                        if true {
                                                            v144 := GenS4{F0: 42, F1: (51 < 1), F2: 368}
                                                            _ = v144
                                                        }
                                                        var v145 float32 = 1e10
                                                        v6 += (v5 & uint(v2))
                                                        _ = v145
                                                    } else {
                                                        v1["a"] = ("\n\t" + "\n\t")
                                                        switch ("a" + ("" + (("" + "\n\t") + ""))) {
                                                        case "foo bar":
                                                            v146 := GenS9{F0: 36, F1: 42, F2: -109}
                                                            _ = v146
                                                        case "hello":
                                                            v2 += 0
                                                        case "":
                                                            v3[("a" + (("hello" + "hello") + "a"))] = (fi94 * (fi109 << ((255) & 63)))
                                                        default:
                                                            fmt.Println(v138, fi109)
                                                        }
                                                        v19 = ("hello" >= "\n\t")
                                                    }
                                                }
                                            } else {
                                                v147 := GenS6{F0: int16(fi109)}
                                                if (128 >= 255) {
                                                    var v148 float64 = ((v138 / (0.1 / (v138 - v138))) - v138)
                                                    if v19 {
                                                        var v149 string = ("foo bar" + "foo bar")
                                                        _ = v149
                                                    }
                                                    var v150 uint = v6
                                                    _ = v148
                                                    _ = v150
                                                }
                                                var v151 uint8 = 53
                                                _ = v147
                                                _ = v151
                                            }
                                        case "\n\t":
                                            if v19 {
                                                v152 := GenS9{F0: -1, F1: -128, F2: 0}
                                                v19 = (false && false)
                                                switch fi94 {
                                                case 127:
                                                    v138 = v138
                                                }
                                                if (2147483647 != -1) {
                                                    v19 = v19
                                                } else {
                                                    v153 := GenS5{F0: (fi109 + (fi94 / ((127) | 1))), F1: v19}
                                                    v110 = append(v110, int(v5))
                                                    switch "hello" {
                                                    case "a":
                                                        v6 -= (v5 * (v5 * v5))
                                                    default:
                                                        v138 = 0.0
                                                        if v19 {
                                                            v4[(fi109 / (((fi94 ^ (fi109 << (((v5 ^ v6)) & 63)))) | 1))] = v19
                                                            v154 := GenS10{F0: (v6 >> (((v5 + (v5 - (v6 | v5)))) & 63))}
                                                            var v155 int32 = 1
                                                            _ = v154
                                                            _ = v155
                                                        }
                                                    }
                                                    _ = v153
                                                }
                                                v6 += (v6 % ((1) | 1))
                                                _ = v152
                                            }
                                        default:
                                            if false {
                                                if v19 {
                                                    v156 := GenS11{F0: 2}
                                                    if (true && false) {
                                                        v19 = v19
                                                    } else {
                                                        v138 -= 1e10
                                                    }
                                                    _ = v156
                                                } else {
                                                    var v157 int32 = 2147483647
                                                    v157 = (v157 >> (((v6 + v5)) & 63))
                                                    v158 := GenS12{F0: -922, F1: v19, F2: 0}
                                                    switch (fi109 - 790) {
                                                    case 9223372036854775807:
                                                        switch ("a" + "") {
                                                        case "hello":
                                                            v159 := GenS8{F0: "a"}
                                                            _ = v159
                                                        default:
                                                            v160 := GenS13{F0: true, F1: (-(v157))}
                                                            _ = v160
                                                        }
                                                    case 27:
                                                        var v161 int = -787
                                                        _ = v161
                                                    case 0:
                                                        _ = len(v1)
                                                    default:
                                                        switch "" {
                                                        case "hello":
                                                            var v162 uint8 = 128
                                                            _ = v162
                                                        default:
                                                            var v163 bool = (false && false)
                                                            _ = v163
                                                        }
                                                    }
                                                    _ = v157
                                                    _ = v158
                                                }
                                                v2 += (v2 | 255)
                                                var v164 uint8 = 2
                                                _ = v164
                                            }
                                        }
                                        _ = v138
                                    } else {
                                        if false {
                                            v110 = append(v110, fi94)
                                            var v165 int32 = -2147483648
                                            _ = v165
                                        } else {
                                            if v19 {
                                                if false {
                                                    var v166 string = "foo bar"
                                                    v110 = append(v110, (fi109 * fi109))
                                                    v19 = false
                                                    _ = v166
                                                } else {
                                                    v5 = v5
                                                }
                                                v19 = true
                                            } else {
                                                switch -1 {
                                                case -846:
                                                    v5 = v6
                                                }
                                                var v167 string = ""
                                                _ = v167
                                            }
                                            fmt.Println(v19, v6, v6)
                                        }
                                    }
                                    v19 = (1000 > 42)
                                    if v19 {
                                        switch fi109 {
                                        case 32768:
                                            _ = len(v110)
                                        }
                                        fmt.Println(v2, v6, v6)
                                        switch "foo bar" {
                                        case "":
                                            if true {
                                                v168 := GenS10{F0: v5}
                                                _ = v168
                                            } else {
                                                var v169 float32 = 3.14159
                                                if len(v110) >= 2 {
                                                    v110 = v110[1:len(v110)]
                                                }
                                                v170 := GenS9{F0: 907, F1: -112, F2: 1}
                                                _ = v169
                                                _ = v170
                                            }
                                        case "hello":
                                            var v171 uint64 = 4294967295
                                            _ = v171
                                        }
                                    }
                                    var v172 uint64 = 18446744073709551615
                                    v172 = v172
                                    _ = v172
                                } else {
                                    v173 := v1["\n\t"]
                                    _ = v173
                                }
                                _ = v117
                            }
                            v110 = append(v110, int(v6))
                            _ = fi109
                            _ = v110
                        }
                        switch 128 {
                        case -131:
                            fmt.Println(v6)
                        case 32768:
                            fmt.Println(v19, v2)
                        case 2147483647:
                            for fi174 := 0; fi174 < 20; fi174++ {
                                v2 = (^(v2))
                                v19 = v19
                                fmt.Println(v6)
                                v175 := GenS2{F0: 1}
                                v2 += (^(v2))
                                _ = fi174
                                _ = v175
                            }
                        }
                        var v176 uint64 = 2
                        _ = v176
                    } else {
                        v2 = v2
                        var v177 int = (fi94 - (fi94 / (((fi94 % (((fi94 * (fi94 | fi94))) | 1))) | 1)))
                        _ = v177
                    }
                case 127:
                    var v178 uint64 = 4294967295
                    _ = v178
                }
            } else {
                switch -1 {
                case 0:
                    if v19 {
                        if ((42 == 42) && (3.14159 == float32(-2.25))) {
                            if v19 {
                                var v179 uint = v6
                                _ = v179
                            } else {
                                if false {
                                    v6 += 0
                                    fmt.Println(v5, fi94, v19)
                                    for fi180 := 0; fi180 < 4; fi180++ {
                                        v181 := GenS14{F0: "a", F1: 0, F2: 0}
                                        _ = fi180
                                        _ = v181
                                    }
                                } else {
                                    if v19 {
                                        switch (fi94 ^ -2147483648) {
                                        case 595:
                                            v182 := GenS15{F0: 9223372036854775807, F1: "a"}
                                            _ = v182
                                        case 127:
                                            for fi183 := 0; fi183 < 9; fi183++ {
                                                v5 -= 18446744073709551615
                                                v6 = 2
                                                fmt.Println(v2)
                                                _ = fi183
                                            }
                                        case 42:
                                            v2 = (v2 + 255)
                                        default:
                                            if ((v2 % (((v2 / ((v2) | 1))) | 1)) <= 614) {
                                                var v184 int = fi94
                                                _ = v184
                                            } else {
                                                for fi185 := 0; fi185 < 10; fi185++ {
                                                    var v186 uint16 = (v2 / ((32767) | 1))
                                                    if ((!false) || (v19 && false)) {
                                                        v187 := GenS1{F0: -1, F1: "a"}
                                                        _ = v187
                                                    }
                                                    if ("" < (("" + "\n\t") + "\n\t")) {
                                                        if v19 {
                                                            var v188 string = ("" + "\n\t")
                                                            var v189 int = int(v6)
                                                            v2 = v186
                                                            v190 := GenS8{F0: ("\n\t" + (v188 + (("hello" + v188) + (v188 + (v188 + v188)))))}
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                v191 := GenS16{F0: -1, F1: v186, F2: (v5 & 981)}
                                                                _ = v191
                                                            case "hello":
                                                                v192 := GenS17{F0: v188}
                                                                _ = v192
                                                            case "a":
                                                                var v193 uint8 = 1
                                                                v193 = v193
                                                                _ = v193
                                                            default:
                                                                v194 := GenS12{F0: -9223372036854775808, F1: v19, F2: 0}
                                                                _ = v194
                                                            }
                                                            v5 += v5
                                                            _ = v188
                                                            _ = v189
                                                            _ = v190
                                                        } else {
                                                            var v195 uint8 = 128
                                                            fmt.Println(v2, fi94, v19)
                                                            _ = v195
                                                        }
                                                        var v196 int64 = genFn0(128, 440)
                                                        _ = v196
                                                    } else {
                                                        var v197 int32 = -2147483648
                                                        if v19 {
                                                            v2 = 32767
                                                        } else {
                                                            v19 = false
                                                            var v198 uint8 = 127
                                                            _ = v198
                                                        }
                                                        v6 += (^(v6))
                                                        fmt.Println(fi185, v2)
                                                        v199 := []int{}
                                                        switch ((("foo bar" + "\n\t") + ("\n\t" + "foo bar")) + "foo bar") {
                                                        case "hello":
                                                            v2 = v186
                                                            v1[(("a" + "a") + "hello")] = "\n\t"
                                                            fmt.Println(fi185, v5)
                                                        default:
                                                            if len(v199) > 0 {
                                                                v199[0] = (fi185 % (((fi94 ^ fi185)) | 1))
                                                            }
                                                        }
                                                        fmt.Println(v2, fi185, fi185)
                                                        var v200 float32 = 0.1
                                                        var v201 int32 = 42
                                                        _ = v197
                                                        _ = v199
                                                        _ = v200
                                                        _ = v201
                                                    }
                                                    v202 := GenS7{F0: "\n\t"}
                                                    if (2 >= 255) {
                                                        v203 := []int{(fi94 - fi94)}
                                                        if ((!(((("hello" + "") + "a") + (((("foo bar" + ("foo bar" + ("hello" + (("foo bar" + (("a" + "") + "a")) + ("foo bar" + ((("a" + ("hello" + ("\n\t" + "hello"))) + "") + ("foo bar" + "a"))))))) + "") + "\n\t") + ("" + "a"))) <= "\n\t")) && (1 < 655)) {
                                                            switch fi185 {
                                                            case 42:
                                                                var v204 uint8 = 170
                                                                _ = v204
                                                            }
                                                            v205 := GenS18{F0: 111, F1: -1}
                                                            if (!false) {
                                                                switch (fi185 << ((2) & 63)) {
                                                                case 0:
                                                                    if v19 {
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            var v206 uint = v6
                                                                            _ = v206
                                                                        case "a":
                                                                            v19 = true
                                                                        default:
                                                                            v207 := GenS19{F0: -32768}
                                                                            _ = v207
                                                                        }
                                                                        v19 = ("" <= "foo bar")
                                                                        fmt.Println(fi185, v19)
                                                                        v186 = (v186 - (^(v186)))
                                                                    } else {
                                                                        var v208 uint64 = 236
                                                                        if len(v203) >= 2 {
                                                                            v203 = v203[1:len(v203)]
                                                                        }
                                                                        _ = v208
                                                                    }
                                                                case 9223372036854775807:
                                                                    var v209 uint = v6
                                                                    _ = v209
                                                                }
                                                            } else {
                                                                if (-1 >= 1) {
                                                                    fmt.Println(fi185, v6, v6)
                                                                    if true {
                                                                        var v210 uint64 = 141
                                                                        switch (fi185 >> (((v5 | (^(v5)))) & 63)) {
                                                                        case 0:
                                                                            if ((v186 | (v2 & (v2 >> ((v5) & 63)))) != v186) {
                                                                                v19 = true
                                                                                v211 := GenS20{F0: 1, F1: 1, F2: (v210 >> (((v5 % (((v6 & (v6 | 136))) | 1))) & 63))}
                                                                                var v212 uint64 = 833
                                                                                _ = v211
                                                                                _ = v212
                                                                            } else {
                                                                                fmt.Println(v19, v6, v210)
                                                                            }
                                                                        default:
                                                                            if len(v203) >= 2 {
                                                                                v203 = v203[1:len(v203)]
                                                                            }
                                                                        }
                                                                        _ = v210
                                                                    }
                                                                    v213 := GenS10{F0: (v5 % ((v5) | 1))}
                                                                    _ = v213
                                                                } else {
                                                                    if (v19 || (32767 != 0)) {
                                                                        var v214 uint = (^(v6))
                                                                        v214 += 4294967295
                                                                        var v215 uint16 = (v2 / ((525) | 1))
                                                                        _ = v214
                                                                        _ = v215
                                                                    }
                                                                    switch fi94 {
                                                                    case -566:
                                                                        v2 += (v186 & (v186 / ((0) | 1)))
                                                                    }
                                                                    v216 := GenS21{F0: 32767, F1: 2}
                                                                    _ = v216
                                                                }
                                                                switch (fi185 << ((0) & 63)) {
                                                                case 127:
                                                                    v217 := GenS20{F0: 100, F1: 255, F2: 725}
                                                                    _ = v217
                                                                default:
                                                                    v203 = append(v203, (fi94 - 564))
                                                                }
                                                                var v218 uint8 = 0
                                                                _ = v218
                                                            }
                                                            _ = v205
                                                        }
                                                        v203 = append(v203, fi185)
                                                        _ = v203
                                                    } else {
                                                        v6 = (v6 << ((v6) & 63))
                                                        v186 -= 2
                                                    }
                                                    switch fi185 {
                                                    case 32768:
                                                        if true {
                                                            v2 -= 65535
                                                            _ = len(v1)
                                                        } else {
                                                            v5 = 1
                                                            v219 := []int{}
                                                            var v220 int64 = -809
                                                            _ = v219
                                                            _ = v220
                                                        }
                                                    case -521:
                                                        v221 := GenS22{F0: 4294967295, F1: -244, F2: 1}
                                                        _ = v221
                                                    case 287:
                                                        var v222 int = (fi94 - fi185)
                                                        _ = v222
                                                    default:
                                                        var v223 int64 = 9223372036854775807
                                                        _ = v223
                                                    }
                                                    _ = fi185
                                                    _ = v186
                                                    _ = v202
                                                }
                                                v224 := GenS17{F0: ""}
                                                _ = v224
                                            }
                                        }
                                    }
                                    for fi225 := 0; fi225 < 12; fi225++ {
                                        v226 := GenS21{F0: v2, F1: 2147483648}
                                        _ = fi225
                                        _ = v226
                                    }
                                }
                                fmt.Println(v2)
                            }
                        }
                        var v227 int8 = -128
                        switch 9223372036854775807 {
                        case 470:
                            var v228 uint = v5
                            _ = v228
                        }
                        _ = v227
                    }
                default:
                    for fi229 := 0; fi229 < 1; fi229++ {
                        v230 := []int{fi94, -849}
                        v2 += 0
                        var v231 uint = v5
                        _ = fi229
                        _ = v230
                        _ = v231
                    }
                }
            }
            for fi232 := 0; fi232 < 3; fi232++ {
                v19 = v19
                v233 := GenS23{F0: 1}
                var v234 int = 141
                if true {
                    if v19 {
                        var v235 uint8 = 128
                        _ = v235
                    } else {
                        if true {
                            _ = len(v4)
                            fmt.Println(fi232)
                            v6 = (v5 - v6)
                            switch "hello" {
                            case "":
                                switch (("\n\t" + ("\n\t" + "foo bar")) + "a") {
                                case "\n\t":
                                    if (fi232 < fi232) {
                                        var v236 int32 = 0
                                        var v237 int8 = 1
                                        v1[("hello" + "\n\t")] = "a"
                                        _ = v236
                                        _ = v237
                                    } else {
                                        fmt.Println(fi94, fi94, v234)
                                        if v19 {
                                            var v238 int = (v234 * 127)
                                            switch "\n\t" {
                                            case "foo bar":
                                                v5 = v6
                                            case "":
                                                var v239 int64 = 64
                                                switch "" {
                                                case "foo bar":
                                                    var v240 int8 = -128
                                                    _ = v240
                                                case "":
                                                    v238 = v238
                                                }
                                                _ = v239
                                            case "a":
                                                v238 = 341
                                            default:
                                                var v241 int32 = 2147483647
                                                _ = v241
                                            }
                                            _ = v238
                                        }
                                        v242 := GenS3{F0: 76}
                                        _ = v242
                                    }
                                case "hello":
                                    v243 := GenS9{F0: 2147483647, F1: 0, F2: -22}
                                    _ = v243
                                }
                            case "\n\t":
                                switch fi94 {
                                case 42:
                                    var v244 uint64 = 819
                                    _ = v244
                                case 793:
                                    v245 := GenS3{F0: -128}
                                    _ = v245
                                default:
                                    if (32768 <= 32768) {
                                        delete(v1, "")
                                    } else {
                                        v246 := []int{(fi232 ^ 0), 332}
                                        v247 := GenS24{F0: 127, F1: v19}
                                        v234 += fi232
                                        if (1 < 9223372036854775807) {
                                            switch fi232 {
                                            case 34:
                                                if (42 != 42) {
                                                    switch "" {
                                                    case "\n\t":
                                                        v248 := GenS5{F0: v234, F1: (!(1000 != 32767))}
                                                        _ = v248
                                                    default:
                                                        v249 := GenS25{F0: (v6 * (v5 ^ v6)), F1: -2147483648}
                                                        _ = v249
                                                    }
                                                }
                                            default:
                                                fmt.Println(v19, v2)
                                            }
                                            fmt.Println(v234)
                                            var v250 float64 = 3.14159
                                            _ = v250
                                        } else {
                                            v251 := v4[fi94]
                                            switch "hello" {
                                            case "":
                                                v234 -= (^(fi94))
                                            case "foo bar":
                                                var v252 uint64 = 18446744073709551615
                                                _ = v252
                                            default:
                                                var v253 uint32 = 74
                                                if v19 {
                                                    if v19 {
                                                        fmt.Println(v2)
                                                        var v254 int = (^(fi94))
                                                        if true {
                                                            if v19 {
                                                                delete(v4, (fi94 << ((198) & 63)))
                                                            } else {
                                                                v255 := GenS10{F0: (^(v6))}
                                                                _ = v255
                                                            }
                                                            v2 += 0
                                                            v256 := GenS26{F0: 0, F1: 638}
                                                            if v251 {
                                                                var v257 uint = 2
                                                                var v258 float32 = (-2.25 * float32(0.0))
                                                                var v259 bool = v19
                                                                _ = v257
                                                                _ = v258
                                                                _ = v259
                                                            }
                                                            if (100 != 127) {
                                                                v251 = (uint8(v5) == 2)
                                                                v253 -= v253
                                                                var v260 int64 = genFn0(154, (v253 - v253))
                                                                var v261 int8 = 127
                                                                _ = v260
                                                                _ = v261
                                                            } else {
                                                                v253 = (v253 - 2)
                                                                if len(v246) >= 2 {
                                                                    v246 = v246[1:len(v246)]
                                                                }
                                                            }
                                                            _ = v256
                                                        } else {
                                                            switch 0 {
                                                            case 128:
                                                                fmt.Println(fi232, v19)
                                                                if ((v253 % ((165) | 1)) > (v253 ^ (v253 / (((v253 % ((447) | 1))) | 1)))) {
                                                                    fmt.Println(v6)
                                                                } else {
                                                                    if ((!((v251 && (0 != (v253 << ((v5) & 63)))) || (!v19))) && (v251 || ((v253 >> ((v5) & 63)) != (v253 + v253)))) {
                                                                        var v262 uint16 = 0
                                                                        var v263 uint16 = 32767
                                                                        v264 := GenS27{F0: uint64(v262), F1: 418, F2: uint32(v2)}
                                                                        _ = v262
                                                                        _ = v263
                                                                        _ = v264
                                                                    }
                                                                }
                                                                v265 := GenS28{F0: v2, F1: 0, F2: 128}
                                                                _ = v265
                                                            case 32768:
                                                                delete(v1, ((("\n\t" + "a") + ("hello" + "hello")) + "\n\t"))
                                                            default:
                                                                v266 := GenS29{F0: (v254 | 1), F1: (v253 >> ((v6) & 63)), F2: -9223372036854775808}
                                                                _ = v266
                                                            }
                                                        }
                                                        _ = v254
                                                    } else {
                                                        var v267 uint8 = 127
                                                        fmt.Println(v2, v251, fi232)
                                                        fmt.Println(v2, v234, v5)
                                                        var v268 uint8 = (v267 - 211)
                                                        switch "\n\t" {
                                                        case "":
                                                            if ((!((v267 != v268) || (0 != 127))) && false) {
                                                                v267 -= 1
                                                                var v269 int8 = 1
                                                                v19 = (((^(v269)) <= v269) && ((35 != 42) || (!v251)))
                                                                _ = v269
                                                            } else {
                                                                var v270 int = fi94
                                                                var v271 float64 = (0.1 + -2.25)
                                                                fmt.Println(v267)
                                                                var v272 int = 1
                                                                if len(v246) >= 2 {
                                                                    v246 = v246[1:len(v246)]
                                                                }
                                                                _ = v270
                                                                _ = v271
                                                                _ = v272
                                                            }
                                                        default:
                                                            var v273 int64 = -137
                                                            _ = v273
                                                        }
                                                        _ = v267
                                                        _ = v268
                                                    }
                                                } else {
                                                    if (true || false) {
                                                        var v274 int = v234
                                                        v253 += 2147483648
                                                        _ = v274
                                                    } else {
                                                        switch "" {
                                                        case "foo bar":
                                                            if len(v246) >= 2 {
                                                                v246 = v246[1:len(v246)]
                                                            }
                                                        case "\n\t":
                                                            var v275 string = ((("hello" + "hello") + "a") + ("\n\t" + "\n\t"))
                                                            _ = v275
                                                        case "a":
                                                            delete(v1, ("foo bar" + "foo bar"))
                                                        }
                                                        if v251 {
                                                            v276 := GenS10{F0: (v5 - v5)}
                                                            _ = v276
                                                        } else {
                                                            v253 = (v253 * (v253 ^ (v253 << ((v5) & 63))))
                                                            fmt.Println(v5, v251, v19)
                                                        }
                                                        var v277 uint64 = 18446744073709551615
                                                        _ = v277
                                                    }
                                                    fmt.Println(v234)
                                                    var v278 float64 = 0.1
                                                    _ = v278
                                                }
                                                v279 := GenS26{F0: 1, F1: uint64(v253)}
                                                switch "" {
                                                case "":
                                                    v280 := GenS22{F0: v253, F1: 32767, F2: 100}
                                                    _ = v280
                                                case "hello":
                                                    if true {
                                                        v2 = v2
                                                        var v281 string = (((("" + "hello") + "foo bar") + (("a" + "a") + ((("a" + "hello") + "") + ((("" + "hello") + (("hello" + "\n\t") + "a")) + ("\n\t" + "hello"))))) + "foo bar")
                                                        v19 = v251
                                                        _ = v281
                                                    } else {
                                                        var v282 float64 = 3.14159
                                                        _ = v282
                                                    }
                                                    var v283 int32 = 1
                                                    v5 = v6
                                                    _ = v283
                                                case "a":
                                                    switch (fi232 | (fi94 | v234)) {
                                                    case 42:
                                                        v284 := GenS30{F0: 0, F1: v6}
                                                        _ = v284
                                                    case 9223372036854775807:
                                                        fmt.Println(fi232, v19)
                                                    case 311:
                                                        v5 = v5
                                                    default:
                                                        switch "" {
                                                        case "foo bar":
                                                            var v285 uint8 = 28
                                                            v6 -= (v6 >> (((^(v6))) & 63))
                                                            _ = v285
                                                        case "hello":
                                                            switch 397 {
                                                            case 1:
                                                                v5 += v6
                                                            case -306:
                                                                v2 += uint16(v253)
                                                            case 0:
                                                                var v286 int32 = -1
                                                                _ = v286
                                                            }
                                                        }
                                                    }
                                                default:
                                                    if (v19 && v251) {
                                                        v287 := GenS31{F0: 2}
                                                        v234 = -206
                                                        _ = v287
                                                    } else {
                                                        v288 := GenS32{F0: uint64(v253), F1: 1, F2: 42}
                                                        fmt.Println(v2, v19, v5)
                                                        _ = v288
                                                    }
                                                }
                                                delete(v3, "\n\t")
                                                _ = v253
                                                _ = v279
                                            }
                                            if ((v251 && ((v19 && false) && v251)) && v251) {
                                                v5 = (v5 / ((20) | 1))
                                            }
                                            _ = v251
                                        }
                                        _ = v246
                                        _ = v247
                                    }
                                }
                            case "a":
                                v289 := GenS22{F0: 4294967295, F1: 1, F2: int8(v234)}
                                _ = v289
                            }
                            if (-2.25 < ((3.14159 + 3.14159) * 1e10)) {
                                v6 += v6
                                var v290 uint16 = (v2 / ((v2) | 1))
                                if v19 {
                                    fmt.Println(v6, v19, v6)
                                }
                                _ = v290
                            } else {
                                switch (v234 * (fi232 ^ (v234 + (^(fi94))))) {
                                case 42:
                                    v291 := []int{-9223372036854775808}
                                    _ = v291
                                default:
                                    fmt.Println(v234, v234)
                                }
                                var v292 uint = (v5 ^ v6)
                                if v19 {
                                    v293 := GenS33{F0: 2}
                                    if (!(146 <= -1)) {
                                        v19 = (205 > 1)
                                        fmt.Println(v6)
                                        var v294 int8 = 42
                                        _ = v294
                                    }
                                    switch fi94 {
                                    case 9223372036854775807:
                                        v295 := []int{42}
                                        _ = v295
                                    }
                                    _ = v293
                                }
                                if v19 {
                                    v296 := GenS33{F0: 979}
                                    _ = v296
                                } else {
                                    v297 := GenS27{F0: 58, F1: v292, F2: 0}
                                    _ = v297
                                }
                                _ = v292
                            }
                            if v19 {
                                if v19 {
                                    switch ("foo bar" + "foo bar") {
                                    case "\n\t":
                                        var v298 uint16 = uint16(v6)
                                        _ = v298
                                    case "foo bar":
                                        var v299 int32 = 2147483647
                                        _ = v299
                                    }
                                    v19 = false
                                } else {
                                    delete(v3, (("a" + "a") + "\n\t"))
                                    v5 -= v6
                                    switch "a" {
                                    case "\n\t":
                                        if v19 {
                                            fmt.Println(fi94)
                                        } else {
                                            v300 := []int{(-(fi232))}
                                            _ = v300
                                        }
                                    case "a":
                                        if v19 {
                                            var v301 int16 = 1
                                            v302 := GenS9{F0: -1, F1: -1, F2: 42}
                                            _ = v301
                                            _ = v302
                                        }
                                    }
                                }
                                switch "" {
                                case "foo bar":
                                    var v303 float32 = 1.5
                                    _ = v303
                                default:
                                    v304 := GenS30{F0: (v6 % ((v6) | 1)), F1: 330}
                                    _ = v304
                                }
                            } else {
                                switch (fi232 & fi94) {
                                case 9223372036854775807:
                                    switch "\n\t" {
                                    case "hello":
                                        v234 -= fi94
                                    case "\n\t":
                                        v305 := []int{}
                                        _ = v305
                                    default:
                                        v306 := GenS30{F0: (v6 << ((v6) & 63)), F1: v5}
                                        _ = v306
                                    }
                                    var v307 float64 = -2.25
                                    v308 := GenS14{F0: ("" + "a"), F1: 18446744073709551615, F2: -38}
                                    if (false && (fi94 == (v234 / ((4294967295) | 1)))) {
                                        switch fi94 {
                                        case -9223372036854775808:
                                            v309 := []int{}
                                            _ = v309
                                        default:
                                            delete(v4, v234)
                                        }
                                        var v310 int16 = 1
                                        switch "" {
                                        case "foo bar":
                                            if v19 {
                                                v307 += v307
                                                var v311 float32 = 1e10
                                                var v312 uint64 = 257
                                                _ = v311
                                                _ = v312
                                            }
                                        case "\n\t":
                                            var v313 int64 = 9223372036854775807
                                            _ = v313
                                        default:
                                            var v314 int64 = -595
                                            _ = v314
                                        }
                                        _ = v310
                                    }
                                    _ = v307
                                    _ = v308
                                case -2147483648:
                                    fmt.Println(v5, v6)
                                default:
                                    fmt.Println(fi94, v5, v2)
                                }
                                if v19 {
                                    var v315 uint16 = v2
                                    v316 := []int{fi94, (v234 << ((v5) & 63)), v234, fi94}
                                    _ = v315
                                    _ = v316
                                } else {
                                    var v317 int = 9223372036854775807
                                    v318 := GenS34{F0: 1, F1: 42}
                                    if (float32(-2.25) == 0.1) {
                                        fmt.Println(v2)
                                        v319 := GenS35{F0: "hello"}
                                        v320 := GenS28{F0: v2, F1: 1, F2: 1}
                                        _ = v319
                                        _ = v320
                                    } else {
                                        v2 += (v2 / (((v2 + (v2 / ((1) | 1)))) | 1))
                                        fmt.Println(v234, v19)
                                    }
                                    var v321 int16 = 331
                                    v317 -= (^(fi232))
                                    _ = v317
                                    _ = v318
                                    _ = v321
                                }
                                v234 -= v234
                                var v322 uint = 255
                                var v323 float32 = 1e10
                                _ = v322
                                _ = v323
                            }
                            switch fi232 {
                            case 4294967295:
                                fmt.Println(v234, v19, v2)
                            case -137:
                                v324 := []int{}
                                v234 = (^(fi232))
                                if v19 {
                                    if true {
                                        if (-1 == -1) {
                                            v234 = (v234 >> ((v6) & 63))
                                        }
                                        var v325 bool = (!v19)
                                        switch (v234 + fi94) {
                                        case 443:
                                            fmt.Println(v19, v19, fi232)
                                        case -920:
                                            switch 32768 {
                                            case 127:
                                                v326 := GenS36{F0: (fi94 ^ 32768), F1: v234, F2: "foo bar"}
                                                _ = v326
                                            case 42:
                                                v19 = v19
                                            default:
                                                if ("hello" == ("" + "")) {
                                                    fmt.Println(fi232)
                                                    if (false && ("a" >= ("foo bar" + "\n\t"))) {
                                                        if v325 {
                                                            var v327 int16 = 42
                                                            _ = v327
                                                        } else {
                                                            var v328 float32 = float32((1e10 / -2.25))
                                                            if true {
                                                                fmt.Println(fi94, v19, v6)
                                                                if (uint64(v6) >= 42) {
                                                                    v324 = append(v324, (fi94 >> ((667) & 63)))
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        v329 := v3["\n\t"]
                                                                        _ = v329
                                                                    case "foo bar":
                                                                        v330 := GenS21{F0: v2, F1: 42}
                                                                        var v331 uint32 = 980
                                                                        _ = v330
                                                                        _ = v331
                                                                    }
                                                                }
                                                            }
                                                            _ = v328
                                                        }
                                                        v6 = 1
                                                    }
                                                } else {
                                                    if v325 {
                                                        _ = len(v4)
                                                        fmt.Println(v19)
                                                        var v332 float32 = -2.25
                                                        v325 = (42 <= 726)
                                                        var v333 uint = v5
                                                        _ = v332
                                                        _ = v333
                                                    } else {
                                                        var v334 float32 = (-2.25 / -2.25)
                                                        v335 := GenS37{F0: 442}
                                                        _ = v334
                                                        _ = v335
                                                    }
                                                    var v336 uint32 = 770
                                                    _ = v336
                                                }
                                            }
                                        case 32767:
                                            var v337 float32 = -2.25
                                            _ = v337
                                        default:
                                            if false {
                                                var v338 int8 = int8(v234)
                                                switch (v234 << ((v5) & 63)) {
                                                case -198:
                                                    if v325 {
                                                        var v339 uint64 = 42
                                                        switch fi94 {
                                                        case 0:
                                                            var v340 int = v234
                                                            _ = v340
                                                        }
                                                        if (!v325) {
                                                            if false {
                                                                var v341 int16 = -32768
                                                                v5 = v5
                                                                _ = v341
                                                            } else {
                                                                v338 += (v338 / (((-(v338))) | 1))
                                                            }
                                                            v342 := GenS26{F0: 0, F1: (v339 & (v339 - 18446744073709551615))}
                                                            switch fi94 {
                                                            case 1:
                                                                v343 := GenS21{F0: v2, F1: 801}
                                                                _ = v343
                                                            case 411:
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    v5 = v5
                                                                case "\n\t":
                                                                    v324 = append(v324, (fi232 & v234))
                                                                case "a":
                                                                    fmt.Println(v339, v5, v5)
                                                                }
                                                            case 128:
                                                                var v344 int64 = 42
                                                                _ = v344
                                                            default:
                                                                var v345 uint = (v5 >> ((628) & 63))
                                                                _ = v345
                                                            }
                                                            _ = v342
                                                        } else {
                                                            fmt.Println(fi232)
                                                            v346 := GenS24{F0: v338, F1: v325}
                                                            if (0 < 183) {
                                                                var v347 float64 = float64((((0.0 * 1.5) - 1e10) + ((3.14159 * 0.1) + 1e10)))
                                                                var v348 string = ("a" + "\n\t")
                                                                var v349 uint16 = 734
                                                                _ = v347
                                                                _ = v348
                                                                _ = v349
                                                            } else {
                                                                switch fi232 {
                                                                case 4294967295:
                                                                    var v350 int = fi232
                                                                    _ = v350
                                                                case 127:
                                                                    v351 := GenS38{F0: 0}
                                                                    _ = v351
                                                                case -996:
                                                                    v6 -= (^(v5))
                                                                default:
                                                                    v5 += 42
                                                                }
                                                                if v325 {
                                                                    v339 = 664
                                                                    var v352 int64 = -210
                                                                    _ = v352
                                                                } else {
                                                                    var v353 int64 = -269
                                                                    _ = v353
                                                                }
                                                            }
                                                            _ = v346
                                                        }
                                                        _ = v339
                                                    }
                                                    var v354 uint64 = 0
                                                    v3[("\n\t" + "\n\t")] = (fi232 >> ((v6) & 63))
                                                    _ = v354
                                                case -220:
                                                    v6 = (v6 ^ (v5 - (v5 + v6)))
                                                default:
                                                    var v355 int16 = -32768
                                                    _ = v355
                                                }
                                                fmt.Println(v2, v2)
                                                _ = v338
                                            }
                                        }
                                        var v356 uint8 = 42
                                        _ = len(v324)
                                        fmt.Println(v325)
                                        _ = v325
                                        _ = v356
                                    }
                                } else {
                                    fmt.Println(v6, v2, v6)
                                    v357 := GenS13{F0: v19, F1: -2147483648}
                                    if (-417 != -586) {
                                        var v358 float32 = 1e10
                                        v359 := GenS39{F0: 692, F1: "a", F2: v2}
                                        v234 -= (fi232 | int(v6))
                                        delete(v3, ("foo bar" + "hello"))
                                        if (18446744073709551615 != uint64(v6)) {
                                            switch fi94 {
                                            case -697:
                                                v360 := GenS21{F0: v2, F1: 4294967295}
                                                var v361 int8 = -127
                                                var v362 bool = v19
                                                _ = v360
                                                _ = v361
                                                _ = v362
                                            case 128:
                                                v6 = v6
                                                var v363 float32 = -2.25
                                                _ = v363
                                            default:
                                                v19 = (!true)
                                            }
                                            v364 := v4[649]
                                            var v365 int16 = -278
                                            _ = v364
                                            _ = v365
                                        } else {
                                            switch v234 {
                                            case -673:
                                                var v366 float32 = (v358 + -2.25)
                                                _ = v366
                                            case 2147483647:
                                                v2 += v2
                                            default:
                                                v5 -= v5
                                            }
                                            v367 := GenS40{F0: v19, F1: 32767, F2: 127}
                                            _ = v367
                                        }
                                        _ = v358
                                        _ = v359
                                    } else {
                                        v368 := v4[(v234 + (fi232 / (((fi232 >> (((v5 & v5)) & 63))) | 1)))]
                                        var v369 int = fi232
                                        v370 := GenS19{F0: 0}
                                        var v371 bool = (v19 && (v19 && (!((v368 && (v368 || (!v19))) || true))))
                                        _ = v368
                                        _ = v369
                                        _ = v370
                                        _ = v371
                                    }
                                    _ = v357
                                }
                                v372 := GenS16{F0: -440, F1: 42, F2: v6}
                                fmt.Println(v2, v19, v234)
                                var v373 int64 = 9223372036854775807
                                _ = v324
                                _ = v372
                                _ = v373
                            }
                        }
                        var v374 int16 = 32767
                        var v375 uint16 = 1
                        _ = v374
                        _ = v375
                    }
                    v234 += (fi232 / ((fi232) | 1))
                    v2 -= v2
                }
                _ = fi232
                _ = v233
                _ = v234
            }
        } else {
            _ = len(v1)
        }
        for fi376 := 0; fi376 < 16; fi376++ {
            if v19 {
                v377 := []int{(fi376 ^ (^(fi94))), -525, fi94}
                if ((fi94 & fi94) >= fi376) {
                    if len(v377) > 0 {
                        v377[0] = -532
                    }
                } else {
                    v378 := GenS41{F0: ("\n\t" + ("\n\t" + "\n\t"))}
                    v6 = (v5 ^ 0)
                    _ = v378
                }
                switch (fi376 - (fi94 >> (((v5 & (v6 * v5))) & 63))) {
                case 32768:
                    v5 -= 18446744073709551615
                case 38:
                    v19 = (!false)
                case 2147483647:
                    var v379 uint = (v6 - v6)
                    _ = v379
                default:
                    var v380 float32 = (1e10 + (0.0 - ((1.5 / 1.0) - 0.1)))
                    _ = v380
                }
                if len(v377) > 0 {
                    v377[0] = -9223372036854775808
                }
                _ = v377
            }
            _ = fi376
        }
        _ = len(v1)
        _ = fi94
    }
    for fi381 := 0; fi381 < 15; fi381++ {
        fmt.Println(v6, v2, v5)
        switch "\n\t" {
        case "\n\t":
            v19 = (42 <= -780)
        default:
            v382 := GenS42{F0: 32767}
            _ = v382
        }
        _ = fi381
    }
    fmt.Println(v2, v6, v19)
    if true {
        v383 := []int{127, 368, 2147483647}
        var v384 uint16 = 989
        if len(v383) >= 2 {
            v383 = v383[1:len(v383)]
        }
        _ = v383
        _ = v384
    } else {
        v385 := []int{32767, 4294967295, 2147483647, 32768}
        v385 = append(v385, 1)
        v2 += (v2 % ((v2) | 1))
        _ = v385
    }
    _ = len(v1)
    v386 := []int{1, 9223372036854775807, -2147483648, 9223372036854775807}
    v6 = (v5 | (v6 | (^(v6))))
    var v387 int = 32767
    switch v387 {
    case 42:
        var v388 uint64 = 4294967295
        _ = v388
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v386:", len(v386), v386)
    fmt.Println("v387:", v387)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v386
    _ = v387
    _ = v4
    _ = v5
    _ = v6
}
