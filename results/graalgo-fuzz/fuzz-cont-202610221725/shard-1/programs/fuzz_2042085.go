package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 string; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 string; F2 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint8; F1 int32 }
type GenS9 struct { F0 bool; F1 string; F2 int64 }
type GenS10 struct { F0 int; F1 uint64; F2 uint32 }
type GenS11 struct { F0 uint16; F1 uint64 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS14 struct { F0 int32; F1 uint }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint64, p1 int16, p2 int) int64 {
    p2 = (p2 << ((1) & 63))
    for fi5 := 0; fi5 < 10; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            if (100 >= 102) {
                if (fi6 >= (p2 + (fi5 / ((fi5) | 1)))) {
                    var v7 int64 = 42
                    _ = v7
                } else {
                    fmt.Println(fi6, p1)
                }
                v8 := GenS1{F0: (p1 & p1)}
                _ = v8
            }
            fmt.Println(fi6, p0, p1)
            _ = fi6
        }
        v9 := []int{(^(fi5))}
        _ = fi5
        _ = v9
    }
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    return 162
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int64 = -395
    v3 := make(map[uint32]int)
    var v4 uint64 = 2
    var v10 uint16 = 249
    if true {
        v11 := []int{-715, 9223372036854775807}
        if ("foo bar" >= "foo bar") {
            v2 -= (v2 ^ (v2 & v2))
            fmt.Println(v1)
        }
        switch ("foo bar" + "a") {
        case "":
            if true {
                var v12 float64 = 1.5
                var v13 string = (("foo bar" + "hello") + ((("foo bar" + "\n\t") + (("" + ("a" + "")) + "a")) + ("foo bar" + ("foo bar" + "\n\t"))))
                v13 = v13
                _ = v12
                _ = v13
            }
        case "a":
            var v14 int = 9223372036854775807
            _ = v14
        default:
            var v15 uint8 = 170
            _ = v15
        }
        _ = v11
    } else {
        var v16 int64 = (v2 & (v2 * 42))
        _ = v16
    }
    v4 += 18446744073709551615
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 int32 = 42
        switch (fi17 * 9223372036854775807) {
        case 4294967295:
            switch (fi17 | (fi17 << ((255) & 63))) {
            case 2147483647:
                var v19 int8 = 42
                _ = v19
            case 127:
                v20 := []int{fi17, (^(fi17)), (fi17 & fi17)}
                _ = v20
            case -9223372036854775808:
                v21 := GenS1{F0: 1000}
                _ = v21
            default:
                v3[v1] = fi17
            }
        case 128:
            for fi22 := 0; fi22 < 20; fi22++ {
                if (v10 != uint16(v1)) {
                    v23 := []int{(fi22 + (fi22 - (fi22 | (fi22 + fi22))))}
                    if (!(true || true)) {
                        v24 := v3[v1]
                        _ = v24
                    } else {
                        v25 := GenS2{F0: 2, F1: ("hello" + ""), F2: (("\n\t" + ("a" + "hello")) + "")}
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                        fmt.Println(v18, v2)
                        v26 := GenS3{F0: 685}
                        _ = v25
                        _ = v26
                    }
                    _ = v23
                }
                v18 = (v18 | (v18 / ((v18) | 1)))
                v1 -= v1
                _ = fi22
            }
        }
        _ = fi17
        _ = v18
    }
    if true {
        if ((^(v1)) >= v1) {
            var v27 int8 = 100
            if true {
                v28 := GenS1{F0: 42}
                _ = v28
            } else {
                switch "\n\t" {
                case "\n\t":
                    if false {
                        if (-1 > -32768) {
                            if true {
                                var v29 float32 = -2.25
                                _ = v29
                            } else {
                                v10 += v10
                            }
                        } else {
                            fmt.Println(v4)
                            var v30 bool = true
                            _ = v30
                        }
                        switch -2147483648 {
                        case 1:
                            var v31 int64 = genFn0((v4 & v4), -69, 12)
                            _ = v31
                        default:
                            switch "" {
                            case "hello":
                                var v32 uint = uint(v4)
                                _ = v32
                            case "":
                                var v33 bool = true
                                _ = v33
                            case "a":
                                switch -9223372036854775808 {
                                case 128:
                                    var v34 uint = 18446744073709551615
                                    _ = v34
                                }
                            }
                        }
                        v35 := v3[v1]
                        _ = v35
                    }
                case "hello":
                    switch 105 {
                    case -2147483648:
                        switch 0 {
                        case 0:
                            v36 := []int{9223372036854775807, 77, -9223372036854775808, 1}
                            _ = v36
                        }
                    default:
                        v37 := []int{}
                        _ = v37
                    }
                }
                for fi38 := 0; fi38 < 11; fi38++ {
                    for fi39 := 0; fi39 < 19; fi39++ {
                        v1 = uint32(v4)
                        var v40 int16 = -273
                        v41 := GenS4{F0: (v1 / (((v1 << ((uint(v10)) & 63))) | 1))}
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                    switch "\n\t" {
                    case "":
                        if false {
                            var v42 string = "\n\t"
                            v27 = v27
                            v27 = (v27 & v27)
                            _ = v42
                        }
                    case "hello":
                        if false {
                            v43 := GenS3{F0: 407}
                            delete(v3, v1)
                            v44 := GenS4{F0: v1}
                            _ = v43
                            _ = v44
                        }
                    case "foo bar":
                        if (!false) {
                            fmt.Println(v4, v27)
                            if true {
                                if true {
                                    v1 += 791
                                    if ((0.1 / -2.25) > (-2.25 - 1.5)) {
                                        var v45 int32 = 804
                                        _ = v45
                                    } else {
                                        switch (("a" + ("hello" + "\n\t")) + ("hello" + "")) {
                                        case "":
                                            var v46 uint8 = 127
                                            _ = v46
                                        default:
                                            if (2 >= (v1 & v1)) {
                                                v47 := GenS3{F0: 305}
                                                var v48 int64 = v2
                                                _ = v47
                                                _ = v48
                                            } else {
                                                var v49 float32 = -2.25
                                                _ = v49
                                            }
                                        }
                                        if false {
                                            v1 = v1
                                            switch (fi38 * fi38) {
                                            case 9223372036854775807:
                                                if (!((0.1 * (-2.25 + ((float64(3.14159) * float64(-2.25)) + (((1e10 + 0.1) - -2.25) * -2.25)))) >= (3.14159 / -2.25))) {
                                                    v50 := GenS3{F0: 1}
                                                    var v51 int = fi38
                                                    var v52 int16 = -1
                                                    _ = v50
                                                    _ = v51
                                                    _ = v52
                                                } else {
                                                    v2 += (-(v2))
                                                    v3[v1] = (fi38 & 0)
                                                    for fi53 := 0; fi53 < 12; fi53++ {
                                                        v4 = v4
                                                        var v54 uint = 0
                                                        var v55 int32 = int32(fi53)
                                                        _ = fi53
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                }
                                            case 917:
                                                if (v27 >= 127) {
                                                    for fi56 := 0; fi56 < 7; fi56++ {
                                                        v57 := GenS4{F0: 4294967295}
                                                        _ = fi56
                                                        _ = v57
                                                    }
                                                    if ((!((!false) || true)) && (1e10 > 3.14159)) {
                                                        var v58 int32 = 2147483647
                                                        if false {
                                                            fmt.Println(v2)
                                                        } else {
                                                            if false {
                                                                if ((v1 + (v1 - v1)) >= v1) {
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        if true {
                                                                            v59 := v3[(v1 * (v1 ^ (v1 & (v1 >> ((255) & 63)))))]
                                                                            v60 := []int{v59, v59}
                                                                            _ = v59
                                                                            _ = v60
                                                                        }
                                                                    }
                                                                } else {
                                                                    v58 -= (v58 + (v58 >> ((96) & 63)))
                                                                    fmt.Println(fi38, v2)
                                                                    v58 = int32(v10)
                                                                }
                                                            }
                                                        }
                                                        var v61 int8 = v27
                                                        _ = v58
                                                        _ = v61
                                                    } else {
                                                        for fi62 := 0; fi62 < 10; fi62++ {
                                                            var v63 uint32 = (v1 | 0)
                                                            var v64 int = -2147483648
                                                            v1 = 690
                                                            v3[2] = fi62
                                                            var v65 int32 = -2147483648
                                                            _ = fi62
                                                            _ = v63
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                    }
                                                    _ = len(v3)
                                                }
                                            }
                                            v66 := GenS5{F0: v1, F1: "a", F2: 1}
                                            _ = v66
                                        }
                                        v2 = (v2 << ((356) & 63))
                                    }
                                    switch fi38 {
                                    case 291:
                                        var v67 int64 = v2
                                        _ = v67
                                    case -2147483648:
                                        var v68 int8 = v27
                                        _ = v68
                                    case 9223372036854775807:
                                        if (float32(0.1) >= (1.5 + 1e10)) {
                                            var v69 uint32 = v1
                                            if false {
                                                fmt.Println(v1, v27, v1)
                                            } else {
                                                var v70 float64 = 1e10
                                                v70 = 1e10
                                                _ = v70
                                            }
                                            for fi71 := 0; fi71 < 20; fi71++ {
                                                v27 -= 1
                                                _ = fi71
                                            }
                                            _ = v69
                                        } else {
                                            var v72 float64 = 1e10
                                            if (0 <= 127) {
                                                v73 := GenS3{F0: (v4 * (v4 ^ (v4 | (v4 - v4))))}
                                                _ = v73
                                            }
                                            switch ("foo bar" + ((("foo bar" + "") + "hello") + "a")) {
                                            case "":
                                                for fi74 := 0; fi74 < 13; fi74++ {
                                                    if false {
                                                        v75 := v3[v1]
                                                        if true {
                                                            v76 := GenS1{F0: 159}
                                                            var v77 uint = 485
                                                            _ = v76
                                                            _ = v77
                                                        } else {
                                                            switch (fi74 * 32768) {
                                                            case -391:
                                                                v78 := []int{-242, v75, (fi74 % (((fi38 - 128)) | 1)), fi38}
                                                                _ = v78
                                                            case 32768:
                                                                v79 := []int{(v75 * fi38)}
                                                                _ = v79
                                                            case -240:
                                                                switch (v75 & v75) {
                                                                case -9223372036854775808:
                                                                    v75 = (^(fi74))
                                                                    var v80 string = ((("" + "") + ("\n\t" + "")) + "")
                                                                    _ = v80
                                                                case 2147483647:
                                                                    v1 = v1
                                                                case 127:
                                                                    v81 := GenS6{F0: uint(v4)}
                                                                    _ = v81
                                                                default:
                                                                    v72 += v72
                                                                }
                                                            default:
                                                                fmt.Println(fi74)
                                                            }
                                                            if (v10 >= (^(v10))) {
                                                                var v82 uint16 = v10
                                                                v83 := []int{42, (fi74 / (((fi38 - v75)) | 1)), (-(fi74))}
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                            switch v75 {
                                                            case 127:
                                                                v84 := GenS7{F0: "a"}
                                                                _ = v84
                                                            case 32767:
                                                                if (!true) {
                                                                    v85 := []int{}
                                                                    if false {
                                                                        v4 += (v4 >> ((247) & 63))
                                                                        _ = len(v85)
                                                                        if false {
                                                                            var v86 uint16 = v10
                                                                            var v87 int8 = 42
                                                                            _ = v86
                                                                            _ = v87
                                                                        } else {
                                                                            v85 = append(v85, (^(fi74)))
                                                                            var v88 int32 = -1
                                                                            switch ("\n\t" + "foo bar") {
                                                                            case "\n\t":
                                                                                v10 -= 214
                                                                            case "a":
                                                                                v72 = 1e10
                                                                            case "hello":
                                                                                switch "foo bar" {
                                                                                case "":
                                                                                    if true {
                                                                                        v89 := GenS8{F0: 161, F1: (v88 >> ((255) & 63))}
                                                                                        _ = v89
                                                                                    } else {
                                                                                        var v90 float64 = ((v72 / v72) * 0.0)
                                                                                        _ = v90
                                                                                    }
                                                                                case "a":
                                                                                    switch (fi74 - fi38) {
                                                                                    case 127:
                                                                                        if ((^(v88)) != (v88 + v88)) {
                                                                                            v72 = float64(1e10)
                                                                                        } else {
                                                                                            if len(v85) > 0 {
                                                                                                v85[0] = (v75 % ((32768) | 1))
                                                                                            }
                                                                                            _ = len(v3)
                                                                                            v91 := GenS7{F0: "a"}
                                                                                            _ = v91
                                                                                        }
                                                                                    case 128:
                                                                                        fmt.Println(v88)
                                                                                    case -632:
                                                                                        if (true && true) {
                                                                                            var v92 int8 = 127
                                                                                            _ = v92
                                                                                        } else {
                                                                                            if (v27 <= int8(v10)) {
                                                                                                var v93 string = "foo bar"
                                                                                                fmt.Println(fi38, v1, v10)
                                                                                                var v94 int64 = -19
                                                                                                v94 = v2
                                                                                                switch v93 {
                                                                                                case "a":
                                                                                                    switch "hello" {
                                                                                                    case "a":
                                                                                                        var v95 int = 9223372036854775807
                                                                                                        _ = v95
                                                                                                    case "foo bar":
                                                                                                        var v96 int32 = (^(v88))
                                                                                                        if false {
                                                                                                            _ = len(v3)
                                                                                                            v97 := GenS5{F0: uint32(v4), F1: v93, F2: 739}
                                                                                                            v75 = -900
                                                                                                            _ = v97
                                                                                                        } else {
                                                                                                            _ = len(v3)
                                                                                                        }
                                                                                                        v98 := GenS5{F0: (v1 - v1), F1: v93, F2: 18446744073709551615}
                                                                                                        _ = v96
                                                                                                        _ = v98
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    v99 := GenS4{F0: (v1 / (((^(v1))) | 1))}
                                                                                                    _ = v99
                                                                                                default:
                                                                                                    fmt.Println(v4, v94, v75)
                                                                                                }
                                                                                                _ = v93
                                                                                                _ = v94
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                case "hello":
                                                                                    fmt.Println(v75)
                                                                                default:
                                                                                    v100 := GenS7{F0: "\n\t"}
                                                                                    _ = v100
                                                                                }
                                                                            }
                                                                            _ = v88
                                                                        }
                                                                    }
                                                                    fmt.Println(fi74, v72, v27)
                                                                    _ = v85
                                                                } else {
                                                                    v101 := []int{(fi38 >> ((460) & 63)), (fi74 << ((280) & 63))}
                                                                    v102 := GenS9{F0: (genFn0(v4, 1000, -801) > genFn0(529, -1, (-(v75)))), F1: "", F2: v2}
                                                                    v103 := GenS4{F0: (v1 >> ((42) & 63))}
                                                                    _ = v101
                                                                    _ = v102
                                                                    _ = v103
                                                                }
                                                            case 1:
                                                                v10 += v10
                                                            default:
                                                                v104 := GenS6{F0: 4294967295}
                                                                _ = v104
                                                            }
                                                        }
                                                        var v105 int8 = 60
                                                        var v106 int64 = (v2 | v2)
                                                        _ = v105
                                                        _ = v106
                                                        _ = v75
                                                    } else {
                                                        fmt.Println(v10, fi74)
                                                    }
                                                    _ = fi74
                                                }
                                            }
                                            _ = v72
                                        }
                                    }
                                } else {
                                    v107 := []int{(-(fi38)), fi38, 2147483647}
                                    v4 = 42
                                    _ = v107
                                }
                                fmt.Println(v10)
                                for fi108 := 0; fi108 < 16; fi108++ {
                                    if ((!false) || true) {
                                        v10 = v10
                                        v109 := GenS10{F0: (^(fi108)), F1: (v4 % (((^(v4))) | 1)), F2: (v1 * v1)}
                                        _ = v109
                                    } else {
                                        delete(v3, (v1 & 4294967295))
                                        v110 := []int{-756}
                                        switch "foo bar" {
                                        case "":
                                            v1 -= v1
                                        default:
                                            fmt.Println(v2, v4)
                                        }
                                        switch (((("a" + "") + "foo bar") + ("" + "a")) + "\n\t") {
                                        case "a":
                                            if ((fi38 - fi38) >= fi38) {
                                                var v111 uint = 538
                                                if true {
                                                    var v112 uint8 = 230
                                                    var v113 bool = true
                                                    _ = v112
                                                    _ = v113
                                                } else {
                                                    switch 519 {
                                                    case -1:
                                                        if len(v110) >= 2 {
                                                            v110 = v110[1:len(v110)]
                                                        }
                                                    case -478:
                                                        switch (("hello" + "foo bar") + "") {
                                                        case "":
                                                            var v114 string = ("" + "")
                                                            _ = v114
                                                        case "a":
                                                            if (0.0 < float64(1.5)) {
                                                                var v115 int64 = v2
                                                                _ = v115
                                                            }
                                                        }
                                                    case 568:
                                                        fmt.Println(v10)
                                                        v116 := GenS7{F0: "a"}
                                                        switch (fi38 / (((fi108 / (((fi108 | 127)) | 1))) | 1)) {
                                                        case 472:
                                                            if (!false) {
                                                                v3[(^(v1))] = 128
                                                            }
                                                        case 2147483647:
                                                            if false {
                                                                if false {
                                                                    if true {
                                                                        v111 = (v111 * 42)
                                                                    }
                                                                    switch fi108 {
                                                                    case 1:
                                                                        _ = len(v3)
                                                                    case 4294967295:
                                                                        var v117 uint8 = uint8(v1)
                                                                        _ = v117
                                                                    case 9223372036854775807:
                                                                        fmt.Println(fi108)
                                                                    }
                                                                } else {
                                                                    if (!false) {
                                                                        v110 = append(v110, -1)
                                                                        v110 = append(v110, fi38)
                                                                        if (132 != 0) {
                                                                            v4 -= v4
                                                                            _ = len(v3)
                                                                            v10 -= 32767
                                                                        } else {
                                                                            v118 := GenS11{F0: v10, F1: v4}
                                                                            var v119 int32 = 1
                                                                            _ = v118
                                                                            _ = v119
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v4)
                                                                        v120 := GenS7{F0: ""}
                                                                        _ = v120
                                                                    }
                                                                }
                                                                var v121 uint8 = 176
                                                                _ = v121
                                                            }
                                                            if true {
                                                                if (-2147483648 != 42) {
                                                                    _ = len(v110)
                                                                    var v122 int64 = (^(v2))
                                                                    if false {
                                                                        var v123 uint32 = (v1 << ((v111) & 63))
                                                                        _ = v123
                                                                    } else {
                                                                        var v124 uint64 = (^(v4))
                                                                        var v125 uint64 = v124
                                                                        _ = len(v3)
                                                                        _ = v124
                                                                        _ = v125
                                                                    }
                                                                    _ = v122
                                                                }
                                                            } else {
                                                                fmt.Println(v111)
                                                                var v126 int8 = (^(v27))
                                                                v1 = 2147483648
                                                                _ = v126
                                                            }
                                                        case 4294967295:
                                                            if true {
                                                                v10 = v10
                                                                var v127 float32 = (0.1 / 1.5)
                                                                v128 := GenS12{F0: fi38}
                                                                switch fi108 {
                                                                case -9223372036854775808:
                                                                    fmt.Println(v4, v10, fi38)
                                                                default:
                                                                    v4 = v4
                                                                }
                                                                var v129 float32 = (0.0 * -2.25)
                                                                _ = v127
                                                                _ = v128
                                                                _ = v129
                                                            }
                                                        default:
                                                            var v130 bool = (v10 == (v10 * uint16(v4)))
                                                            _ = v130
                                                        }
                                                        var v131 int32 = 0
                                                        _ = v116
                                                        _ = v131
                                                    default:
                                                        fmt.Println(fi38, v27, fi38)
                                                        v2 -= v2
                                                        _ = len(v110)
                                                    }
                                                    switch "hello" {
                                                    case "hello":
                                                        var v132 string = ""
                                                        _ = v132
                                                    default:
                                                        fmt.Println(v1)
                                                    }
                                                    var v133 int = fi108
                                                    var v134 int8 = int8(v4)
                                                    switch (fi108 << (((v111 - v111)) & 63)) {
                                                    case 2147483647:
                                                        v135 := GenS13{F0: 679, F1: v27, F2: (v1 % (((v1 & v1)) | 1))}
                                                        _ = v135
                                                    case 4294967295:
                                                        if (false || true) {
                                                            if (!(v2 == (v2 << ((65535) & 63)))) {
                                                                switch (("" + "a") + "hello") {
                                                                case "a":
                                                                    var v136 int64 = (v2 | int64(v10))
                                                                    _ = v136
                                                                case "hello":
                                                                    fmt.Println(v1)
                                                                case "\n\t":
                                                                    var v137 bool = (true && ("foo bar" == "foo bar"))
                                                                    _ = v137
                                                                default:
                                                                    var v138 int = fi38
                                                                    _ = v138
                                                                }
                                                            }
                                                            var v139 uint8 = 217
                                                            _ = v139
                                                        } else {
                                                            v27 -= -38
                                                        }
                                                    }
                                                    _ = v133
                                                    _ = v134
                                                }
                                                var v140 int32 = 925
                                                _ = v111
                                                _ = v140
                                            }
                                        case "\n\t":
                                            v141 := v3[(v1 - v1)]
                                            _ = v141
                                        case "hello":
                                            var v142 bool = (!((720 == 0) || true))
                                            _ = v142
                                        }
                                        _ = v110
                                    }
                                    v10 -= 1
                                    _ = fi108
                                }
                            } else {
                                v143 := []int{(fi38 >> ((92) & 63)), (fi38 & fi38), fi38}
                                fmt.Println(v27, v10, v1)
                                v143 = append(v143, fi38)
                                _ = v143
                            }
                            if (v1 <= (v1 / ((388) | 1))) {
                                v2 = v2
                            } else {
                                _ = len(v3)
                                v144 := GenS10{F0: (^(fi38)), F1: 0, F2: v1}
                                _ = v144
                            }
                        } else {
                            var v145 int = (-(fi38))
                            for fi146 := 0; fi146 < 8; fi146++ {
                                switch fi146 {
                                case 0:
                                    var v147 uint = 590
                                    _ = v147
                                }
                                _ = fi146
                            }
                            v148 := []int{fi38, 127, fi38}
                            _ = v145
                            _ = v148
                        }
                    }
                    _ = fi38
                }
                var v149 int8 = 105
                _ = v149
            }
            v3[v1] = -1
            _ = v27
        }
        v150 := GenS14{F0: -977, F1: 708}
        for fi151 := 0; fi151 < 15; fi151++ {
            v152 := GenS15{F0: v4}
            _ = fi151
            _ = v152
        }
        _ = v150
    }
    v3[v1] = 9223372036854775807
    var v153 int32 = -2147483648
    v154 := GenS16{F0: 57}
    switch 32767 {
    case 2147483647:
        v4 += v4
    case 127:
        v10 = (v10 ^ 159)
    case -1:
        for fi155 := 0; fi155 < 3; fi155++ {
            fmt.Println(v10, v153, fi155)
            v156 := GenS17{F0: (v10 << ((4294967295) & 63))}
            _ = fi155
            _ = v156
        }
    default:
        switch "\n\t" {
        case "foo bar":
            v157 := []int{972, 4294967295}
            _ = v157
        }
    }
    switch "foo bar" {
    case "\n\t":
        v158 := GenS13{F0: 32767, F1: -122, F2: v1}
        _ = v158
    case "":
        v2 += v2
    case "a":
        v159 := []int{2147483647, -1, 575}
        _ = v159
    }
    var v160 int8 = -32
    for fi161 := 0; fi161 < 10; fi161++ {
        var v162 float32 = 1.5
        v1 = (v1 * (^(v1)))
        if (0.0 < v162) {
            var v163 int64 = 568
            v164 := GenS18{F0: 24, F1: genFn0(v4, 992, fi161)}
            _ = v163
            _ = v164
        }
        _ = fi161
        _ = v162
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v153:", v153)
    fmt.Println("v160:", v160)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v153
    _ = v154
    _ = v160
    _ = v2
    _ = v3
    _ = v4
}
