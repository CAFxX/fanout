package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint8; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint16; F1 uint }
type GenS10 struct { F0 int16; F1 uint8 }
func genFn0(p0 string, p1 int, p2 uint) uint8 {
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 int16 = 0
        var v8 uint8 = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 float32 = -2.25
    var v4 string = ("hello" + ("foo bar" + "hello"))
    var v5 float64 = 1.5
    var v10 int32 = 601
    for fi11 := 0; fi11 < 3; fi11++ {
        v3 = v3
        v4 = "a"
        _ = fi11
    }
    var v12 float32 = v3
    for fi13 := 0; fi13 < 16; fi13++ {
        if false {
            if (32767 == 0) {
                if (!false) {
                    fmt.Println(v12, v10, fi13)
                    for fi14 := 0; fi14 < 18; fi14++ {
                        if ((false || true) && (false && (42 == 1))) {
                            v15 := []int{int(v10)}
                            if false {
                                var v16 uint32 = 4294967295
                                _ = v16
                            } else {
                                var v17 int32 = v10
                                fmt.Println(v10, v17)
                                _ = len(v2)
                                _ = v17
                            }
                            v18 := GenS1{F0: (fi14 * fi14)}
                            _ = v15
                            _ = v18
                        } else {
                            v19 := []int{(fi14 | (fi13 >> ((322) & 63))), 1, (fi14 - -966)}
                            _ = v19
                        }
                        var v20 int16 = -1
                        _ = fi14
                        _ = v20
                    }
                } else {
                    var v21 bool = false
                    _ = v21
                }
            } else {
                switch 0 {
                case -9223372036854775808:
                    v22 := []int{}
                    _ = v22
                default:
                    for fi23 := 0; fi23 < 14; fi23++ {
                        v24 := GenS1{F0: fi13}
                        v25 := GenS1{F0: (-(fi23))}
                        v5 = 1e10
                        v4 = ((((v4 + v4) + "foo bar") + "\n\t") + v4)
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
                v26 := GenS2{F0: -89}
                for fi27 := 0; fi27 < 13; fi27++ {
                    var v28 float32 = -2.25
                    if false {
                        var v29 int = (fi27 - fi13)
                        var v30 int = 32767
                        _ = v29
                        _ = v30
                    } else {
                        var v31 int32 = v10
                        _ = v31
                    }
                    _ = fi27
                    _ = v28
                }
                _ = v26
            }
            switch v4 {
            case "\n\t":
                for fi32 := 0; fi32 < 1; fi32++ {
                    if false {
                        switch 892 {
                        case 236:
                            v10 -= (v10 + v10)
                        }
                    }
                    delete(v1, (fi32 >> ((512) & 63)))
                    _ = fi32
                }
            case "":
                v12 += (-2.25 / (v12 / v12))
            default:
                for fi33 := 0; fi33 < 1; fi33++ {
                    v3 = (3.14159 + v3)
                    if false {
                        var v34 int16 = 1
                        _ = v34
                    }
                    _ = fi33
                }
            }
            v35 := []int{2147483647}
            _ = v35
        }
        _ = fi13
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        for fi37 := 0; fi37 < 20; fi37++ {
            fmt.Println(v4)
            var v38 int8 = 127
            fmt.Println(v38, v12, v38)
            _ = fi37
            _ = v38
        }
        v4 = "hello"
        if (2 < 2) {
            v39 := GenS3{F0: (fi36 >> ((26) & 63)), F1: 128, F2: 4294967295}
            switch fi36 {
            case 829:
                var v40 float64 = 3.14159
                _ = v40
            case 32767:
                for fi41 := 0; fi41 < 3; fi41++ {
                    v12 = (v3 * ((1.5 / v12) * v12))
                    switch v4 {
                    case "":
                        if true {
                            delete(v2, fi36)
                        } else {
                            v5 = v5
                        }
                    case "\n\t":
                        var v42 int = fi36
                        _ = v42
                    case "hello":
                        if true {
                            v5 -= (v5 + 1e10)
                            var v43 uint = 629
                            if false {
                                var v44 int16 = -442
                                if (fi36 >= -2147483648) {
                                    v45 := GenS1{F0: int(v43)}
                                    _ = v45
                                } else {
                                    v10 = v10
                                    switch "" {
                                    case "foo bar":
                                        var v46 uint32 = 815
                                        _ = v46
                                    }
                                    var v47 uint = 0
                                    _ = v47
                                }
                                if false {
                                    v48 := GenS3{F0: (fi36 | fi41), F1: 255, F2: v43}
                                    switch "\n\t" {
                                    case "a":
                                        v12 = (float32((v5 * (v5 * v5))) * v3)
                                    case "foo bar":
                                        var v49 float32 = ((-2.25 * float32((0.0 + 0.1))) * (((1.5 / 0.1) * (0.0 / 1e10)) + -2.25))
                                        _ = v49
                                    default:
                                        v5 += (0.1 - v5)
                                    }
                                    var v50 string = v4
                                    var v51 int32 = v10
                                    _ = v48
                                    _ = v50
                                    _ = v51
                                } else {
                                    switch fi36 {
                                    case 127:
                                        v10 += 42
                                    }
                                    v3 -= float32(((v5 * v5) - (((3.14159 - 1.5) * 0.1) * float64(float32(v5)))))
                                }
                                v52 := []int{}
                                _ = v44
                                _ = v52
                            } else {
                                switch fi41 {
                                case 32767:
                                    v3 += v3
                                case -85:
                                    v43 = v43
                                case 531:
                                    v53 := v1[(-(fi41))]
                                    _ = v53
                                default:
                                    if true {
                                        v54 := GenS1{F0: (^(fi41))}
                                        _ = v54
                                    } else {
                                        v1[(-(fi41))] = 655
                                        fmt.Println(v43)
                                        fmt.Println(v10, fi41, v43)
                                    }
                                }
                            }
                            if (false || true) {
                                v2[127] = 42
                            } else {
                                v55 := []int{(^(fi36)), fi36, 32768, fi36}
                                v56 := GenS4{F0: -128}
                                _ = v55
                                _ = v56
                            }
                            _ = v43
                        } else {
                            v4 = (v4 + (v4 + (("a" + v4) + (((v4 + v4) + (v4 + v4)) + ("foo bar" + ((("\n\t" + ("hello" + v4)) + ("" + v4)) + (v4 + "a")))))))
                            v57 := []int{}
                            var v58 int8 = 88
                            _ = v57
                            _ = v58
                        }
                    default:
                        v59 := GenS2{F0: fi36}
                        _ = v59
                    }
                    var v60 int64 = 42
                    _ = fi41
                    _ = v60
                }
            case 42:
                v61 := GenS2{F0: fi36}
                _ = v61
            }
            v62 := GenS4{F0: 63}
            _ = v39
            _ = v62
        }
        _ = fi36
    }
    for fi63 := 0; fi63 < 10; fi63++ {
        v5 = (((v5 / -2.25) - 1e10) / v5)
        var v64 bool = false
        switch ("" + v4) {
        case "foo bar":
            v5 = v5
        case "hello":
            v65 := GenS3{F0: fi63, F1: 210, F2: 4294967295}
            _ = v65
        default:
            v66 := GenS1{F0: fi63}
            _ = v66
        }
        _ = fi63
        _ = v64
    }
    v67 := []int{}
    switch v4 {
    case "hello":
        v12 -= ((v12 / ((v3 / v3) / v3)) / v3)
    default:
        fmt.Println(v4, v12)
    }
    for fi68 := 0; fi68 < 15; fi68++ {
        if false {
            v3 += float32(v5)
            var v69 string = ("" + v4)
            if true {
                if (781 >= 551) {
                    fmt.Println(fi68)
                    v5 = v5
                    if (-9223372036854775808 != 199) {
                        var v70 int = fi68
                        _ = v70
                    }
                } else {
                    if len(v67) >= 2 {
                        v67 = v67[1:len(v67)]
                    }
                    var v71 int = fi68
                    for fi72 := 0; fi72 < 20; fi72++ {
                        switch (fi72 % ((fi72) | 1)) {
                        case 128:
                            switch v4 {
                            case "\n\t":
                                v10 -= -2147483648
                            case "hello":
                                var v73 bool = ((v71 / (((v71 - fi68)) | 1)) >= (^(v71)))
                                _ = v73
                            }
                        default:
                            v12 = v3
                            v10 += (v10 / (((v10 * (v10 & (v10 & (v10 * (v10 - 42)))))) | 1))
                        }
                        v69 = v69
                        if true {
                            var v74 uint8 = 2
                            v75 := GenS5{F0: 1}
                            _ = v74
                            _ = v75
                        } else {
                            var v76 uint16 = 255
                            if true {
                                v77 := GenS6{F0: v76}
                                switch "" {
                                case "a":
                                    switch "foo bar" {
                                    case "\n\t":
                                        switch (v71 << ((4294967295) & 63)) {
                                        case -1:
                                            if len(v67) > 0 {
                                                v67[0] = 9223372036854775807
                                            }
                                        case -169:
                                            switch v69 {
                                            case "":
                                                var v78 int64 = 0
                                                _ = v78
                                            case "a":
                                                v5 += (v5 * v5)
                                            case "hello":
                                                if len(v67) >= 2 {
                                                    v67 = v67[1:len(v67)]
                                                }
                                            default:
                                                delete(v2, v71)
                                            }
                                        }
                                    default:
                                        var v79 int64 = 9223372036854775807
                                        var v80 uint = 790
                                        var v81 uint = (v80 - v80)
                                        if true {
                                            v10 += (v10 & v10)
                                            fmt.Println(v76)
                                            var v82 uint32 = 65535
                                            _ = v82
                                        }
                                        _ = v79
                                        _ = v80
                                        _ = v81
                                    }
                                case "\n\t":
                                    v83 := v1[1]
                                    _ = v83
                                }
                                v5 = v5
                                var v84 uint32 = 528
                                _ = v77
                                _ = v84
                            }
                            _ = v76
                        }
                        _ = fi72
                    }
                    _ = v71
                }
                for fi85 := 0; fi85 < 6; fi85++ {
                    v86 := GenS4{F0: -1}
                    switch 42 {
                    case -984:
                        var v87 float32 = v12
                        _ = v87
                    default:
                        var v88 uint32 = 4294967295
                        _ = v88
                    }
                    if true {
                        var v89 uint = 381
                        _ = v89
                    } else {
                        v90 := GenS7{F0: 2}
                        var v91 int16 = 32767
                        _ = v90
                        _ = v91
                    }
                    _ = fi85
                    _ = v86
                }
            } else {
                switch -1 {
                case -9223372036854775808:
                    for fi92 := 0; fi92 < 7; fi92++ {
                        if len(v67) > 0 {
                            v67[0] = (-(fi68))
                        }
                        switch v69 {
                        case "foo bar":
                            switch ((v4 + v69) + v69) {
                            case "\n\t":
                                if true {
                                    if (!(false && (0 >= 2))) {
                                        if ((fi92 * -186) >= fi68) {
                                            v93 := GenS3{F0: 127, F1: 128, F2: 2}
                                            v5 -= float64(v12)
                                            _ = v93
                                        }
                                    }
                                    switch fi92 {
                                    case 4294967295:
                                        v94 := GenS1{F0: (fi92 / ((fi68) | 1))}
                                        _ = v94
                                    case 104:
                                        if (-128 < 104) {
                                            delete(v2, fi92)
                                            var v95 uint64 = 736
                                            v95 = v95
                                            _ = v95
                                        } else {
                                            _ = len(v67)
                                            var v96 float64 = 1.5
                                            var v97 uint16 = 0
                                            _ = v96
                                            _ = v97
                                        }
                                    case 9223372036854775807:
                                        if true {
                                            switch fi92 {
                                            case -2147483648:
                                                _ = len(v1)
                                            default:
                                                switch fi68 {
                                                case 647:
                                                    v10 += v10
                                                case -2147483648:
                                                    switch v69 {
                                                    case "hello":
                                                        var v98 float32 = v3
                                                        _ = v98
                                                    }
                                                case -383:
                                                    fmt.Println(v10, fi92, v4)
                                                }
                                            }
                                        }
                                    default:
                                        v2[(fi68 * -570)] = 0
                                    }
                                    if true {
                                        var v99 uint16 = 0
                                        if ((!(173 >= v10)) || false) {
                                            switch (fi92 << ((0) & 63)) {
                                            case 9223372036854775807:
                                                if (996 <= uint(v99)) {
                                                    v3 -= (v3 - ((v3 / v12) * 1.5))
                                                    v69 = (v69 + v4)
                                                }
                                            case 4294967295:
                                                v3 -= v3
                                            }
                                            var v100 uint8 = 255
                                            v4 = (("hello" + ((v4 + v69) + v69)) + v69)
                                            v101 := GenS8{F0: true}
                                            switch fi92 {
                                            case -193:
                                                fmt.Println(v4)
                                            case 0:
                                                var v102 float32 = (0.1 + 0.1)
                                                _ = v102
                                            case 42:
                                                v5 += 1.5
                                            }
                                            if (!false) {
                                                v3 -= float32((1.5 * 3.14159))
                                            } else {
                                                if ((fi92 % ((fi92) | 1)) != -864) {
                                                    var v103 uint32 = 2147483648
                                                    if len(v67) > 0 {
                                                        v67[0] = (fi92 ^ fi92)
                                                    }
                                                    _ = v103
                                                } else {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if false {
                                                            v104 := GenS2{F0: 0}
                                                            _ = v104
                                                        }
                                                    case "":
                                                        fmt.Println(v100, v12)
                                                    default:
                                                        switch (-(fi92)) {
                                                        case 42:
                                                            _ = len(v2)
                                                            switch (fi68 / (((fi92 ^ fi92)) | 1)) {
                                                            case 127:
                                                                switch v69 {
                                                                case "":
                                                                    v1[4294967295] = 1
                                                                case "\n\t":
                                                                    v4 = v4
                                                                    v105 := GenS5{F0: 2}
                                                                    v12 = ((float32((0.0 / (((0.1 - 0.0) - ((3.14159 * 1e10) - (((v5 * float64(((v3 - 3.14159) - v12))) - 1e10) * 1.5))) - ((v5 / (v5 / (v5 - float64(-2.25)))) / (v5 / 1.5))))) + (v3 * ((float32(v5) * ((v3 / v3) * v3)) + (v12 - 0.0)))) * -2.25)
                                                                    _ = v105
                                                                }
                                                            case 42:
                                                                fmt.Println(v3, v99, v99)
                                                            case 329:
                                                                v106 := GenS8{F0: true}
                                                                _ = v106
                                                            default:
                                                                v107 := GenS7{F0: (v99 << ((445) & 63))}
                                                                _ = v107
                                                            }
                                                        case 611:
                                                            var v108 bool = (true || (-355 == int16(v10)))
                                                            _ = v108
                                                        }
                                                        fmt.Println(fi92, v4)
                                                        switch (fi92 * (fi92 * int(v99))) {
                                                        case 0:
                                                            var v109 int = (fi68 % ((fi92) | 1))
                                                            _ = v109
                                                        }
                                                        _ = len(v1)
                                                        var v110 bool = (-642 > 654)
                                                        _ = v110
                                                    }
                                                }
                                                v10 = (v10 - v10)
                                                var v111 int32 = v10
                                                _ = v111
                                            }
                                            switch fi68 {
                                            case 0:
                                                switch v69 {
                                                case "\n\t":
                                                    if len(v67) >= 2 {
                                                        v67 = v67[1:len(v67)]
                                                    }
                                                case "foo bar":
                                                    var v112 uint64 = 0
                                                    _ = v112
                                                case "a":
                                                    fmt.Println(v3, fi68, fi68)
                                                }
                                            default:
                                                if true {
                                                    var v113 bool = (!false)
                                                    _ = v113
                                                }
                                            }
                                            _ = v100
                                            _ = v101
                                        }
                                        v114 := GenS3{F0: -1, F1: 222, F2: 4294967295}
                                        v115 := GenS9{F0: v99, F1: 255}
                                        if len(v67) >= 2 {
                                            v67 = v67[1:len(v67)]
                                        }
                                        v116 := GenS10{F0: 1, F1: 0}
                                        _ = v114
                                        _ = v115
                                        _ = v116
                                        _ = v99
                                    } else {
                                        var v117 bool = ((-(fi68)) <= fi68)
                                        if v117 {
                                            switch (fi68 >> ((851) & 63)) {
                                            case 162:
                                                if true {
                                                    fmt.Println(v5, v4)
                                                    v3 = 1e10
                                                }
                                            case -1:
                                                _ = len(v2)
                                            case 0:
                                                v118 := GenS8{F0: v117}
                                                _ = v118
                                            }
                                            v3 = 0.0
                                            var v119 uint32 = 2
                                            _ = v119
                                        } else {
                                            v120 := GenS10{F0: 170, F1: 2}
                                            v3 -= v3
                                            _ = v120
                                        }
                                        if (!(!(42 != 18446744073709551615))) {
                                            fmt.Println(v3, fi68)
                                            v117 = true
                                        } else {
                                            fmt.Println(fi92)
                                            if true {
                                                var v121 int32 = v10
                                                v10 = (v10 / ((v121) | 1))
                                                var v122 int8 = 1
                                                _ = v121
                                                _ = v122
                                            } else {
                                                var v123 uint = 1
                                                if (519 == 0) {
                                                    v69 = "foo bar"
                                                }
                                                _ = v123
                                            }
                                        }
                                        _ = v117
                                    }
                                }
                            case "foo bar":
                                var v124 uint8 = 0
                                _ = v124
                            default:
                                fmt.Println(v69, v4, v4)
                            }
                        case "\n\t":
                            var v125 int = fi68
                            _ = v125
                        }
                        fmt.Println(fi68, v10, v12)
                        _ = fi92
                    }
                case 0:
                    var v126 uint32 = 0
                    _ = v126
                case 9223372036854775807:
                    v12 = ((float32(1e10) - v3) * v3)
                }
                var v127 int8 = -1
                _ = v127
            }
            _ = v69
        } else {
            if len(v67) > 0 {
                v67[0] = 128
            }
            var v128 string = ("a" + "foo bar")
            fmt.Println(v12)
            _ = v128
        }
        _ = fi68
    }
    fmt.Println(v10, v12, v3)
    fmt.Println(v3, v5)
    switch "" {
    case "\n\t":
        var v129 uint32 = 65535
        _ = v129
    }
    v130 := GenS7{F0: 719}
    fmt.Println(v10, v5)
    if (!true) {
        var v131 int8 = -128
        var v132 int64 = -697
        for fi133 := 0; fi133 < 13; fi133++ {
            delete(v1, (fi133 / ((2147483647) | 1)))
            if false {
                fmt.Println(v12, fi133)
            }
            var v134 uint8 = 2
            _ = fi133
            _ = v134
        }
        _ = v131
        _ = v132
    } else {
        var v135 uint32 = 703
        _ = v135
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v67:", len(v67), v67)
    _ = v1
    _ = v10
    _ = v12
    _ = v130
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v67
}
