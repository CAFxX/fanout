package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int32; F2 int32 }
type GenS3 struct { F0 uint8; F1 bool; F2 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 string; F1 string }
func genFn0(p0 uint8, p1 uint16, p2 int16) uint {
    p1 = p1
    switch "foo bar" {
    case "hello":
        if true {
            v4 := []int{32768, 594}
            var v5 float64 = 3.14159
            _ = v4
            _ = v5
        } else {
            var v6 uint = 1
            var v7 int16 = (p2 * p2)
            p0 = 255
            _ = v6
            _ = v7
        }
    case "":
        p0 = 42
    case "a":
        var v8 bool = true
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1() uint64 {
    if true {
        var v9 bool = true
        _ = v9
    }
    if (!true) {
        if false {
            switch ("" + "foo bar") {
            case "a":
                v10 := []int{}
                _ = v10
            case "hello":
                if (175 >= 255) {
                    var v11 uint16 = 65535
                    _ = v11
                }
            case "foo bar":
                if (1e10 != 0.0) {
                    for fi12 := 0; fi12 < 2; fi12++ {
                        var v13 float64 = (0.1 - 1e10)
                        switch "" {
                        case "foo bar":
                            for fi14 := 0; fi14 < 4; fi14++ {
                                v13 = v13
                                _ = fi14
                            }
                        case "":
                            for fi15 := 0; fi15 < 11; fi15++ {
                                fmt.Println(fi12)
                                _ = fi15
                            }
                        default:
                            var v16 int = (fi12 / ((127) | 1))
                            _ = v16
                        }
                        _ = fi12
                        _ = v13
                    }
                    var v17 uint8 = 86
                    _ = v17
                } else {
                    if (true || true) {
                        if false {
                            v18 := []int{128, 2147483647, 128}
                            switch "" {
                            case "":
                                v19 := GenS1{F0: 255}
                                _ = v19
                            case "hello":
                                switch "foo bar" {
                                case "":
                                    _ = len(v18)
                                }
                            default:
                                if len(v18) > 0 {
                                    v18[0] = 127
                                }
                            }
                            var v20 uint32 = 2
                            _ = v18
                            _ = v20
                        }
                    } else {
                        for fi21 := 0; fi21 < 18; fi21++ {
                            for fi22 := 0; fi22 < 20; fi22++ {
                                switch fi22 {
                                case 127:
                                    if true {
                                        v23 := []int{}
                                        if len(v23) > 0 {
                                            v23[0] = (fi22 + (fi21 | fi21))
                                        }
                                        _ = v23
                                    } else {
                                        fmt.Println(fi22)
                                        var v24 uint16 = 782
                                        _ = v24
                                    }
                                default:
                                    switch "" {
                                    case "\n\t":
                                        v25 := []int{(fi22 * (fi21 + (fi22 % (((^(fi22))) | 1))))}
                                        _ = v25
                                    }
                                }
                                fmt.Println(fi21)
                                _ = fi22
                            }
                            var v26 int16 = 1000
                            _ = fi21
                            _ = v26
                        }
                        v27 := []int{1, 32767}
                        _ = v27
                    }
                }
            default:
                for fi28 := 0; fi28 < 16; fi28++ {
                    var v29 int64 = 42
                    fmt.Println(v29, v29)
                    v30 := GenS2{F0: v29, F1: -837, F2: -2147483648}
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
        } else {
            for fi31 := 0; fi31 < 9; fi31++ {
                if false {
                    v32 := []int{2147483647}
                    _ = v32
                }
                var v33 int8 = 100
                _ = fi31
                _ = v33
            }
        }
        for fi34 := 0; fi34 < 18; fi34++ {
            var v35 string = (("hello" + (("a" + ("a" + "\n\t")) + "")) + "")
            v36 := GenS3{F0: 106, F1: ((1 <= 1) && true), F2: 127}
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        var v38 uint32 = 1
        if (-806 == 215) {
            for fi39 := 0; fi39 < 5; fi39++ {
                v40 := []int{(^(fi39)), fi39, fi39, fi37}
                _ = fi39
                _ = v40
            }
            v41 := GenS2{F0: -9223372036854775808, F1: 2147483647, F2: 1}
            for fi42 := 0; fi42 < 17; fi42++ {
                v43 := []int{-51, fi42}
                _ = fi42
                _ = v43
            }
            _ = v41
        }
        v44 := []int{fi37, 48, (^(fi37)), (fi37 & fi37)}
        _ = fi37
        _ = v38
        _ = v44
    }
    v45 := []int{4294967295}
    var v46 int = 32768
    _ = v45
    _ = v46
    return 432
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 42
    v3 := make(map[int]string)
    v2 = v2
    v47 := GenS3{F0: 2, F1: true, F2: 127}
    if (!((127 == 27) && (609 >= 2))) {
        v2 = (v2 ^ -1)
    } else {
        v2 = 42
        for fi48 := 0; fi48 < 20; fi48++ {
            if true {
                fmt.Println(v2, v2)
            } else {
                var v49 uint64 = genFn1()
                for fi50 := 0; fi50 < 11; fi50++ {
                    v2 += v2
                    _ = fi50
                }
                for fi51 := 0; fi51 < 5; fi51++ {
                    _ = len(v1)
                    v52 := []int{(fi51 & 4294967295), fi48, fi48}
                    if true {
                        switch (fi51 | 9223372036854775807) {
                        case -633:
                            fmt.Println(v2, v2, v49)
                        case 2147483647:
                            switch (fi48 ^ -257) {
                            case -829:
                                switch (("a" + "foo bar") + "hello") {
                                case "":
                                    v53 := GenS2{F0: -1, F1: -1, F2: 191}
                                    switch ((("\n\t" + "") + "hello") + "foo bar") {
                                    case "\n\t":
                                        var v54 int64 = -329
                                        _ = v54
                                    }
                                    v55 := GenS3{F0: 1, F1: (1 > 389), F2: v2}
                                    _ = v53
                                    _ = v55
                                }
                                switch fi48 {
                                case 42:
                                    if len(v52) >= 2 {
                                        v52 = v52[1:len(v52)]
                                    }
                                default:
                                    v49 = 18446744073709551615
                                }
                                v52 = append(v52, (fi48 & fi51))
                            case 127:
                                if false {
                                    var v56 int64 = 802
                                    fmt.Println(fi48, v49)
                                    if false {
                                        if false {
                                            if false {
                                                if len(v52) > 0 {
                                                    v52[0] = (fi48 | (fi48 << ((18446744073709551615) & 63)))
                                                }
                                                var v57 bool = (637 >= 427)
                                                v49 = v49
                                                v56 -= (v56 >> ((42) & 63))
                                                _ = v57
                                            } else {
                                                v3[int(v56)] = "\n\t"
                                                var v58 int = 9223372036854775807
                                                var v59 int8 = v2
                                                v3[fi48] = ("a" + ("a" + "foo bar"))
                                                _ = v58
                                                _ = v59
                                            }
                                            if (!(146 < 0)) {
                                                switch fi51 {
                                                case 491:
                                                    fmt.Println(v49, v2, fi51)
                                                }
                                                if (true || true) {
                                                    var v60 int64 = (v56 | v56)
                                                    _ = v60
                                                }
                                                if false {
                                                    v61 := GenS3{F0: 255, F1: false, F2: (v2 << ((42) & 63))}
                                                    _ = v61
                                                } else {
                                                    v62 := GenS2{F0: (v56 / ((1) | 1)), F1: 504, F2: -1}
                                                    if (!(819 < 0)) {
                                                        fmt.Println(fi51, fi48)
                                                    }
                                                    fmt.Println(fi51)
                                                    _ = v62
                                                }
                                            } else {
                                                if (!true) {
                                                    var v63 string = ("a" + "a")
                                                    _ = v63
                                                }
                                            }
                                        } else {
                                            v64 := GenS1{F0: 255}
                                            var v65 float32 = -2.25
                                            fmt.Println(fi51, v65, v65)
                                            _ = v64
                                            _ = v65
                                        }
                                    }
                                    _ = v56
                                }
                            }
                            v66 := GenS1{F0: uint16(v49)}
                            var v67 int16 = 42
                            if false {
                                v67 -= (v67 * 1000)
                                fmt.Println(fi51, v67, v2)
                                var v68 float32 = 1e10
                                _ = v68
                            }
                            var v69 string = (("foo bar" + ("a" + (("a" + "\n\t") + "hello"))) + "a")
                            _ = v66
                            _ = v67
                            _ = v69
                        case 32767:
                            v2 += -128
                        default:
                            var v70 int = 9223372036854775807
                            _ = v70
                        }
                        var v71 string = "hello"
                        switch 32768 {
                        case 128:
                            var v72 string = v71
                            _ = v72
                        case -591:
                            if false {
                                var v73 string = "foo bar"
                                var v74 uint8 = 12
                                v75 := GenS1{F0: 716}
                                v76 := GenS3{F0: v74, F1: (!(42 < -1)), F2: v2}
                                fmt.Println(fi48, v73)
                                v71 = v71
                                v52 = append(v52, fi51)
                                v77 := GenS2{F0: -1, F1: 1, F2: -2147483648}
                                _ = v73
                                _ = v74
                                _ = v75
                                _ = v76
                                _ = v77
                            } else {
                                var v78 uint16 = 688
                                fmt.Println(v71, v71)
                                if (v71 <= v71) {
                                    v78 += (v78 - v78)
                                    v78 -= 857
                                    fmt.Println(v71, fi51, v2)
                                }
                                _ = v78
                            }
                        default:
                            switch ("hello" + v71) {
                            case "":
                                var v79 uint = 4294967295
                                _ = v79
                            }
                        }
                        fmt.Println(fi51, v71, v2)
                        v80 := v3[fi51]
                        var v81 bool = true
                        var v82 uint32 = 0
                        v71 = "foo bar"
                        _ = v71
                        _ = v80
                        _ = v81
                        _ = v82
                    } else {
                        var v83 int16 = 847
                        fmt.Println(v2)
                        _ = v83
                    }
                    _ = fi51
                    _ = v52
                }
                _ = v49
            }
            v84 := GenS4{F0: -1}
            _ = fi48
            _ = v84
        }
    }
    switch -513 {
    case -818:
        switch 370 {
        case 619:
            _ = len(v1)
        }
    }
    for fi85 := 0; fi85 < 14; fi85++ {
        var v86 string = ""
        fmt.Println(v86, v86, fi85)
        _ = fi85
        _ = v86
    }
    if (32767 != 409) {
        var v87 float64 = 0.0
        if true {
            v88 := GenS4{F0: 109}
            _ = v88
        } else {
            fmt.Println(v2, v2)
            switch "\n\t" {
            case "foo bar":
                if true {
                    if true {
                        v2 += v2
                        v2 = v2
                    }
                    switch ("a" + "hello") {
                    case "":
                        if false {
                            v2 = 91
                        } else {
                            var v89 uint16 = 0
                            if (false || (-9223372036854775808 != 0)) {
                                switch 473 {
                                case 4294967295:
                                    var v90 uint8 = 2
                                    _ = v90
                                case 128:
                                    var v91 string = ""
                                    _ = v91
                                }
                                var v92 uint8 = uint8(v89)
                                _ = v92
                            }
                            for fi93 := 0; fi93 < 12; fi93++ {
                                for fi94 := 0; fi94 < 11; fi94++ {
                                    v89 -= v89
                                    v95 := GenS4{F0: -1}
                                    v87 -= v87
                                    _ = fi94
                                    _ = v95
                                }
                                for fi96 := 0; fi96 < 20; fi96++ {
                                    v89 = v89
                                    _ = fi96
                                }
                                _ = fi93
                            }
                            _ = v89
                        }
                    default:
                        if (0 != -1) {
                            fmt.Println(v2, v87)
                        }
                    }
                } else {
                    v87 += (v87 * (v87 - 0.1))
                    if (0 != 32767) {
                        v97 := []int{794, 42}
                        var v98 uint = 4294967295
                        _ = v97
                        _ = v98
                    } else {
                        if (!false) {
                            var v99 int16 = 32767
                            _ = v99
                        }
                        var v100 uint = 764
                        v101 := GenS3{F0: 0, F1: true, F2: (v2 & v2)}
                        _ = v100
                        _ = v101
                    }
                    for fi102 := 0; fi102 < 3; fi102++ {
                        v103 := []int{(fi102 % ((-349) | 1)), fi102, 688}
                        _ = len(v103)
                        v87 = v87
                        _ = fi102
                        _ = v103
                    }
                }
            }
            v2 -= v2
        }
        v87 = v87
        _ = v87
    } else {
        switch "\n\t" {
        case "":
            v2 += v2
        case "foo bar":
            v2 += v2
        case "a":
            for fi104 := 0; fi104 < 10; fi104++ {
                if false {
                    if false {
                        switch (fi104 + fi104) {
                        case 188:
                            var v105 uint16 = 32767
                            _ = v105
                        default:
                            var v106 uint32 = 346
                            _ = v106
                        }
                        var v107 int32 = -839
                        _ = v107
                    } else {
                        switch 42 {
                        case -1:
                            v2 = (v2 >> ((4294967295) & 63))
                        case -457:
                            for fi108 := 0; fi108 < 3; fi108++ {
                                v2 += v2
                                _ = fi108
                            }
                        case -925:
                            v3[fi104] = ("hello" + "hello")
                        default:
                            var v109 uint64 = 18446744073709551615
                            _ = v109
                        }
                    }
                    v2 = 120
                    switch (fi104 + (^(fi104))) {
                    case 127:
                        v2 -= v2
                    case 461:
                        var v110 float32 = 3.14159
                        _ = v110
                    case 289:
                        switch "hello" {
                        case "a":
                            v111 := []int{fi104, fi104, fi104}
                            _ = v111
                        case "hello":
                            v112 := []int{(fi104 | (fi104 | fi104))}
                            _ = v112
                        case "\n\t":
                            var v113 int32 = -2147483648
                            _ = v113
                        }
                    }
                } else {
                    for fi114 := 0; fi114 < 2; fi114++ {
                        switch fi104 {
                        case 0:
                            fmt.Println(v2, v2, fi104)
                        case 693:
                            v2 += v2
                        case 9223372036854775807:
                            if true {
                                var v115 uint64 = 1
                                _ = v115
                            }
                        }
                        v116 := GenS1{F0: 2}
                        _ = fi114
                        _ = v116
                    }
                    for fi117 := 0; fi117 < 19; fi117++ {
                        delete(v3, fi104)
                        v118 := GenS2{F0: 1, F1: -2147483648, F2: 1}
                        _ = fi117
                        _ = v118
                    }
                    for fi119 := 0; fi119 < 10; fi119++ {
                        fmt.Println(fi119, fi119)
                        var v120 int8 = (v2 * v2)
                        _ = fi119
                        _ = v120
                    }
                }
                var v121 int = 42
                if true {
                    switch fi104 {
                    case 2147483647:
                        v121 -= (^(v121))
                    case -2147483648:
                        var v122 uint8 = 16
                        _ = v122
                    default:
                        fmt.Println(v121, v121)
                    }
                }
                _ = fi104
                _ = v121
            }
        default:
            var v123 int = 0
            _ = v123
        }
        v124 := []int{}
        v125 := GenS5{F0: 42, F1: 34}
        _ = v124
        _ = v125
    }
    switch (("\n\t" + "") + "foo bar") {
    case "\n\t":
        if (v2 != v2) {
            var v126 uint64 = 4294967295
            _ = v126
        } else {
            for fi127 := 0; fi127 < 16; fi127++ {
                var v128 int32 = 42
                for fi129 := 0; fi129 < 9; fi129++ {
                    v130 := GenS4{F0: -1}
                    _ = fi129
                    _ = v130
                }
                var v131 uint64 = 1
                _ = fi127
                _ = v128
                _ = v131
            }
        }
    case "hello":
        for fi132 := 0; fi132 < 19; fi132++ {
            v2 = v2
            v133 := GenS2{F0: -9223372036854775808, F1: 42, F2: -22}
            _ = fi132
            _ = v133
        }
    case "":
        if (478 > genFn0(1, 255, 1)) {
            switch ("a" + ((("\n\t" + (((("a" + "\n\t") + "foo bar") + (("foo bar" + (("" + "hello") + "hello")) + "")) + "a")) + "hello") + "\n\t")) {
            case "foo bar":
                var v134 int64 = 595
                _ = v134
            case "\n\t":
                var v135 int32 = -717
                _ = v135
            default:
                switch "a" {
                case "foo bar":
                    switch 626 {
                    case 2147483647:
                        var v136 uint = genFn0(128, 65535, -990)
                        _ = v136
                    case -178:
                        v137 := GenS6{F0: "", F1: ("a" + "foo bar")}
                        _ = v137
                    case -9223372036854775808:
                        var v138 float64 = 0.0
                        _ = v138
                    }
                default:
                    for fi139 := 0; fi139 < 4; fi139++ {
                        if (!false) {
                            v2 = (^(v2))
                            v140 := GenS5{F0: -2147483648, F1: 2147483648}
                            var v141 int32 = 2147483647
                            _ = v140
                            _ = v141
                        }
                        for fi142 := 0; fi142 < 6; fi142++ {
                            v2 = v2
                            if (!false) {
                                v143 := GenS5{F0: -1, F1: 1}
                                fmt.Println(fi142, fi139)
                                _ = v143
                            }
                            v2 = v2
                            _ = fi142
                        }
                        _ = fi139
                    }
                }
            }
            v144 := []int{1, 32768, 504}
            _ = v144
        }
    default:
        v2 = v2
    }
    v145 := GenS6{F0: "\n\t", F1: "a"}
    v2 = v2
    switch "\n\t" {
    case "\n\t":
        v2 = (^(v2))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v145
    _ = v2
    _ = v3
    _ = v47
}
