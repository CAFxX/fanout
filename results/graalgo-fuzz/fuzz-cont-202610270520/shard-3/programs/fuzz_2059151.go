package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 bool; F1 bool; F2 uint }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int16; F1 int }
type GenS10 struct { F0 int8; F1 int64; F2 string }
type GenS11 struct { F0 int64; F1 bool; F2 uint8 }
func genFn0() int {
    if true {
        var v7 int32 = -2147483648
        _ = v7
    }
    var v8 float64 = -2.25
    switch "a" {
    case "foo bar":
        if (false && false) {
            v9 := GenS1{F0: -206, F1: 2}
            v10 := GenS1{F0: 895, F1: 32767}
            _ = v10
            _ = v9
        }
    case "":
        var v11 int16 = 1000
        _ = v11
    }
    _ = v8
    return -41
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]bool)
    var v3 string = ""
    v4 := make(map[int]bool)
    var v5 float64 = 1.5
    var v6 uint32 = 685
    if false {
        for fi12 := 0; fi12 < 17; fi12++ {
            if false {
                v6 += v6
                switch 123 {
                case -2147483648:
                    v13 := v1[(fi12 * fi12)]
                    _ = v13
                case 128:
                    if (65535 >= 65535) {
                        v14 := GenS2{F0: 378, F1: 127}
                        var v15 uint16 = uint16(v6)
                        if (!(1 <= 2147483647)) {
                            if (true && true) {
                                v16 := []int{(fi12 & fi12), (fi12 >> ((844) & 63)), (fi12 ^ (fi12 - fi12))}
                                var v17 int32 = -867
                                var v18 uint8 = 82
                                _ = v16
                                _ = v17
                                _ = v18
                            } else {
                                var v19 string = ("hello" + v3)
                                for fi20 := 0; fi20 < 7; fi20++ {
                                    switch v3 {
                                    case "hello":
                                        var v21 float64 = v5
                                        _ = v21
                                    case "foo bar":
                                        if true {
                                            if (v6 == 2147483648) {
                                                v6 += (v6 ^ 1)
                                                if (4294967295 < v6) {
                                                    var v22 uint32 = v6
                                                    _ = v22
                                                } else {
                                                    if (0.1 >= 1.5) {
                                                        v23 := GenS1{F0: -1, F1: v15}
                                                        _ = v23
                                                    }
                                                    v24 := GenS2{F0: 102, F1: uint8(v15)}
                                                    _ = v24
                                                }
                                                if true {
                                                    switch v3 {
                                                    case "":
                                                        v15 = 921
                                                    case "a":
                                                        fmt.Println(v15)
                                                    case "foo bar":
                                                        v19 = v19
                                                    }
                                                }
                                            }
                                            switch (("" + ("a" + "hello")) + "") {
                                            case "a":
                                                var v25 uint32 = (v6 / ((4294967295) | 1))
                                                _ = v25
                                            case "\n\t":
                                                v26 := GenS2{F0: 1000, F1: 2}
                                                _ = v26
                                            default:
                                                fmt.Println(v5, fi12, fi20)
                                            }
                                            fmt.Println(v3)
                                        } else {
                                            switch "foo bar" {
                                            case "hello":
                                                fmt.Println(v15, v5)
                                            case "":
                                                v27 := []int{(fi20 * (fi20 + (fi20 ^ (fi12 / (((^(fi12))) | 1))))), fi20, -586}
                                                _ = v27
                                            case "foo bar":
                                                switch (fi12 + (fi12 & (fi20 << ((18446744073709551615) & 63)))) {
                                                case 128:
                                                    fmt.Println(v19, v3, v5)
                                                case 36:
                                                    v28 := v2[(fi20 << ((255) & 63))]
                                                    _ = v28
                                                case -9223372036854775808:
                                                    v15 -= (^(v15))
                                                }
                                            default:
                                                var v29 int8 = 42
                                                _ = v29
                                            }
                                            switch (v19 + v3) {
                                            case "hello":
                                                v30 := []int{fi12, (fi20 >> ((4294967295) & 63))}
                                                _ = v30
                                            }
                                            var v31 float32 = 0.1
                                            _ = v31
                                        }
                                    default:
                                        v32 := GenS2{F0: -548, F1: 0}
                                        _ = v32
                                    }
                                    _ = fi20
                                }
                                for fi33 := 0; fi33 < 14; fi33++ {
                                    var v34 int16 = 42
                                    _ = fi33
                                    _ = v34
                                }
                                _ = v19
                            }
                            for fi35 := 0; fi35 < 15; fi35++ {
                                if (false && (1 >= 4294967295)) {
                                    v3 = "foo bar"
                                } else {
                                    v6 = 424
                                    fmt.Println(v6)
                                }
                                _ = fi35
                            }
                            v36 := []int{}
                            _ = v36
                        }
                        _ = v14
                        _ = v15
                    }
                case 9223372036854775807:
                    if (145 != -694) {
                        v37 := GenS3{F0: 1000}
                        _ = v37
                    }
                }
                if (855 == 508) {
                    var v38 string = "a"
                    v38 = v38
                    if true {
                        v39 := GenS4{F0: genFn0(), F1: 0}
                        for fi40 := 0; fi40 < 14; fi40++ {
                            v6 = v6
                            v41 := GenS1{F0: -9223372036854775808, F1: 636}
                            var v42 uint16 = 255
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                        _ = v39
                    }
                    _ = v38
                } else {
                    fmt.Println(v3)
                    v6 += (v6 | (v6 - (v6 >> ((4294967295) & 63))))
                }
            } else {
                if true {
                    fmt.Println(v3, fi12)
                    v6 = (v6 * (^(v6)))
                }
                v6 = (v6 >> ((65535) & 63))
                for fi43 := 0; fi43 < 10; fi43++ {
                    if ((!(!(false || true))) && ("\n\t" < "\n\t")) {
                        var v44 string = ""
                        v3 = v3
                        _ = v44
                    }
                    if (false || false) {
                        var v45 int32 = 42
                        v6 = v6
                        _ = v45
                    } else {
                        v46 := GenS4{F0: genFn0(), F1: 722}
                        if (!false) {
                            var v47 float64 = (v5 / v5)
                            v48 := GenS3{F0: -32768}
                            _ = v47
                            _ = v48
                        } else {
                            var v49 uint8 = 2
                            v50 := []int{}
                            _ = v49
                            _ = v50
                        }
                        switch fi12 {
                        case -1:
                            var v51 int64 = 438
                            _ = v51
                        case 42:
                            v3 = ""
                        default:
                            switch "foo bar" {
                            case "a":
                                v52 := []int{(fi12 - fi12), genFn0(), fi12}
                                _ = v52
                            case "hello":
                                if (561 < 0) {
                                    v53 := GenS5{F0: true, F1: (((v3 != ("hello" + v3)) || true) || false), F2: 255}
                                    v3 = ((v3 + v3) + ("hello" + (("a" + ("foo bar" + (v3 + (((v3 + "a") + (("" + "a") + v3)) + v3)))) + v3)))
                                    var v54 int64 = 0
                                    v5 -= v5
                                    v55 := v4[fi43]
                                    if false {
                                        if (false || (2147483647 >= int32(v54))) {
                                            v55 = v55
                                            v56 := []int{2147483647, genFn0(), 0, fi43}
                                            _ = v56
                                        }
                                        v6 = 2147483647
                                        v57 := []int{}
                                        _ = v57
                                    }
                                    v58 := GenS2{F0: 428, F1: 1}
                                    if true {
                                        var v59 uint32 = v6
                                        v5 -= v5
                                        switch "a" {
                                        case "a":
                                            var v60 float32 = float32((v5 * 0.1))
                                            _ = v60
                                        case "":
                                            var v61 int = -852
                                            _ = v61
                                        case "hello":
                                            var v62 uint8 = 128
                                            _ = v62
                                        }
                                        _ = v59
                                    } else {
                                        var v63 uint32 = v6
                                        var v64 uint64 = 558
                                        if v55 {
                                            if true {
                                                v6 = (v63 >> ((181) & 63))
                                                switch 492 {
                                                case 32767:
                                                    _ = len(v2)
                                                case 0:
                                                    v5 += v5
                                                default:
                                                    fmt.Println(fi12, v55, fi43)
                                                }
                                                if v55 {
                                                    if v55 {
                                                        fmt.Println(v63, v55)
                                                    } else {
                                                        v55 = (float32(((v5 - v5) * v5)) >= ((0.0 + (3.14159 + 3.14159)) - -2.25))
                                                    }
                                                    fmt.Println(v6, fi12, v3)
                                                    v65 := []int{}
                                                    _ = v65
                                                } else {
                                                    var v66 bool = v55
                                                    _ = v66
                                                }
                                            }
                                            switch "hello" {
                                            case "":
                                                fmt.Println(fi43)
                                            case "hello":
                                                var v67 bool = ((v54 + 918) != v54)
                                                _ = v67
                                            case "\n\t":
                                                switch v3 {
                                                case "a":
                                                    var v68 uint8 = 189
                                                    v55 = false
                                                    var v69 uint32 = v63
                                                    v69 = v69
                                                    if false {
                                                        var v70 int = genFn0()
                                                        _ = v70
                                                    }
                                                    _ = v68
                                                    _ = v69
                                                case "\n\t":
                                                    v71 := GenS5{F0: (false && (v54 < -9223372036854775808)), F1: false, F2: 255}
                                                    _ = v71
                                                }
                                            }
                                            var v72 string = ("\n\t" + "a")
                                            _ = v72
                                        } else {
                                            var v73 int8 = 72
                                            fmt.Println(fi12, v64)
                                            _ = v73
                                        }
                                        _ = v63
                                        _ = v64
                                    }
                                    v74 := []int{genFn0()}
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                    _ = v58
                                    _ = v74
                                }
                                switch (fi43 << ((2) & 63)) {
                                case 127:
                                    var v75 string = v3
                                    _ = v75
                                case -688:
                                    switch (fi43 / ((-624) | 1)) {
                                    case 0:
                                        if (2 != 2) {
                                            v5 -= (-2.25 * v5)
                                            v6 = 4294967295
                                        }
                                    default:
                                        if false {
                                            var v76 uint = 65535
                                            if (!true) {
                                                v77 := []int{1}
                                                if (int8(v76) < -1) {
                                                    if false {
                                                        if (!false) {
                                                            if (!true) {
                                                                v78 := v2[(fi12 - 595)]
                                                                var v79 uint = (v76 >> (((v76 % (((v76 - 570)) | 1))) & 63))
                                                                if false {
                                                                    if false {
                                                                        var v80 int32 = -2147483648
                                                                        v77 = append(v77, genFn0())
                                                                        _ = v80
                                                                    } else {
                                                                        var v81 uint = v76
                                                                        var v82 uint64 = 2
                                                                        var v83 uint32 = v6
                                                                        _ = v81
                                                                        _ = v82
                                                                        _ = v83
                                                                    }
                                                                    if false {
                                                                        fmt.Println(v6)
                                                                    } else {
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            var v84 int64 = 407
                                                                            _ = v84
                                                                        case "":
                                                                            var v85 int = (fi43 - (fi12 | (-(fi43))))
                                                                            _ = v85
                                                                        case "\n\t":
                                                                            var v86 bool = v78
                                                                            _ = v86
                                                                        }
                                                                        if (!false) {
                                                                            if true {
                                                                                delete(v1, genFn0())
                                                                            }
                                                                        } else {
                                                                            v6 = 65
                                                                            _ = len(v4)
                                                                            if (!v78) {
                                                                                v87 := GenS4{F0: (fi12 - (fi43 >> (((v79 | 42)) & 63))), F1: (^(v76))}
                                                                                fmt.Println(fi12, v79)
                                                                                if v78 {
                                                                                    var v88 int8 = 127
                                                                                    _ = v88
                                                                                } else {
                                                                                    if false {
                                                                                        v4[fi12] = v78
                                                                                        var v89 int64 = 103
                                                                                        _ = len(v1)
                                                                                        _ = v89
                                                                                    } else {
                                                                                        var v90 float64 = (3.14159 * 0.0)
                                                                                        var v91 uint64 = 0
                                                                                        _ = v90
                                                                                        _ = v91
                                                                                    }
                                                                                    fmt.Println(v78, fi12)
                                                                                    if (((fi43 + int(v6)) == (fi12 << ((v76) & 63))) && v78) {
                                                                                        var v92 uint16 = 32767
                                                                                        var v93 uint8 = 23
                                                                                        var v94 int32 = 0
                                                                                        _ = v92
                                                                                        _ = v93
                                                                                        _ = v94
                                                                                    } else {
                                                                                        fmt.Println(v79, fi12)
                                                                                        v6 = 0
                                                                                    }
                                                                                }
                                                                                _ = v87
                                                                            }
                                                                        }
                                                                        v3 = (("hello" + "foo bar") + "hello")
                                                                    }
                                                                } else {
                                                                    if v78 {
                                                                        switch (fi43 << ((v79) & 63)) {
                                                                        case 32768:
                                                                            var v95 int8 = 127
                                                                            _ = v95
                                                                        case -1:
                                                                            switch v3 {
                                                                            case "":
                                                                                var v96 uint16 = 2
                                                                                _ = v96
                                                                            case "a":
                                                                                v79 = v76
                                                                            }
                                                                            if ((3.14159 - 1e10) != float32(0.0)) {
                                                                                if ((fi12 << ((0) & 63)) == fi12) {
                                                                                    v97 := GenS6{F0: 32767, F1: 42}
                                                                                    v79 -= v76
                                                                                    _ = v97
                                                                                } else {
                                                                                    v98 := v2[(fi12 * (fi43 | fi12))]
                                                                                    v78 = v78
                                                                                    var v99 float32 = (1e10 * 0.1)
                                                                                    var v100 int32 = -2147483648
                                                                                    v101 := GenS3{F0: 1000}
                                                                                    _ = v100
                                                                                    _ = v101
                                                                                    _ = v98
                                                                                    _ = v99
                                                                                }
                                                                                v102 := GenS7{F0: -123, F1: 0}
                                                                                fmt.Println(v79, v3, v3)
                                                                                _ = v102
                                                                            } else {
                                                                                if (42 > 0) {
                                                                                    var v103 float64 = 0.0
                                                                                    v104 := GenS6{F0: -32768, F1: -32768}
                                                                                    _ = v103
                                                                                    _ = v104
                                                                                }
                                                                                switch v3 {
                                                                                case "":
                                                                                    switch v3 {
                                                                                    case "foo bar":
                                                                                        v79 += (^(v76))
                                                                                    case "":
                                                                                        if (v78 || v78) {
                                                                                            v76 += (v79 & v76)
                                                                                            v6 = 2147483647
                                                                                            var v105 int = fi43
                                                                                            _ = v105
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if (985 != -9223372036854775808) {
                                                                                    v77 = append(v77, fi43)
                                                                                    v106 := GenS7{F0: 1, F1: 95}
                                                                                    var v107 uint64 = 18446744073709551615
                                                                                    _ = v106
                                                                                    _ = v107
                                                                                }
                                                                            }
                                                                        default:
                                                                            v76 = (v79 + (v79 | 4294967295))
                                                                        }
                                                                        var v108 bool = (uint64(v6) < 478)
                                                                        _ = v108
                                                                    } else {
                                                                        v6 = v6
                                                                        v79 += v76
                                                                    }
                                                                    var v109 uint = uint(v6)
                                                                    var v110 uint8 = 128
                                                                    _ = v109
                                                                    _ = v110
                                                                }
                                                                if len(v77) > 0 {
                                                                    v77[0] = (fi43 | 9223372036854775807)
                                                                }
                                                                var v111 uint16 = 930
                                                                fmt.Println(v3, fi12)
                                                                var v112 uint = (v79 ^ 255)
                                                                _ = v111
                                                                _ = v112
                                                                _ = v78
                                                                _ = v79
                                                            }
                                                        } else {
                                                            v76 -= 2
                                                            v3 = v3
                                                            if false {
                                                                v113 := GenS8{F0: 1}
                                                                _ = v113
                                                            }
                                                        }
                                                        fmt.Println(v6)
                                                        v3 = ("" + "\n\t")
                                                    }
                                                    fmt.Println(fi43, fi43)
                                                }
                                                v76 = 42
                                                _ = v77
                                            }
                                            v6 = (v6 | (^(v6)))
                                            _ = v76
                                        }
                                    }
                                default:
                                    if (18446744073709551615 <= 387) {
                                        v114 := []int{(fi12 | fi12), 127, fi12, fi43}
                                        delete(v1, fi12)
                                        _ = v114
                                    } else {
                                        _ = len(v4)
                                        v3 = v3
                                    }
                                    v115 := GenS9{F0: -32768, F1: (fi12 & genFn0())}
                                    _ = v115
                                }
                                v6 = 534
                            case "\n\t":
                                v3 = v3
                            default:
                                if false {
                                    v116 := GenS3{F0: 0}
                                    _ = len(v4)
                                    var v117 bool = false
                                    _ = v116
                                    _ = v117
                                } else {
                                    var v118 int64 = 42
                                    v119 := []int{(fi12 % ((0) | 1)), genFn0()}
                                    if (!(!false)) {
                                        v5 += v5
                                        v120 := GenS10{F0: -1, F1: v118, F2: v3}
                                        _ = v120
                                    }
                                    _ = v118
                                    _ = v119
                                }
                            }
                        }
                        v121 := GenS11{F0: 0, F1: false, F2: 2}
                        _ = v121
                        _ = v46
                    }
                    v122 := []int{(-(fi43)), 32768, (fi43 >> ((4294967295) & 63))}
                    var v123 bool = ((((v5 / (((((0.0 - (((3.14159 * 0.1) + (0.1 - v5)) - v5)) + v5) / v5) - float64(-2.25)) + v5)) / (3.14159 / v5)) <= 1.5) && false)
                    var v124 int16 = 470
                    _ = fi43
                    _ = v122
                    _ = v123
                    _ = v124
                }
            }
            v3 = ""
            fmt.Println(fi12)
            _ = fi12
        }
        v125 := []int{32768}
        _ = v125
    }
    if (!(255 > 65535)) {
        var v126 uint8 = 127
        if ((1e10 <= -2.25) || (!true)) {
            var v127 int32 = 42
            _ = v127
        }
        for fi128 := 0; fi128 < 7; fi128++ {
            switch (fi128 << ((2) & 63)) {
            case 1:
                fmt.Println(v5, v6)
            case 32768:
                var v129 int8 = 127
                _ = v129
            case 243:
                v6 -= 42
            }
            _ = fi128
        }
        _ = v126
    } else {
        switch 915 {
        case -339:
            var v130 int64 = -9223372036854775808
            _ = v130
        case 0:
            v6 = 2
        }
        var v131 uint = 0
        _ = v131
    }
    for fi132 := 0; fi132 < 13; fi132++ {
        if (-1 > 42) {
            switch v3 {
            case "a":
                fmt.Println(v5)
            case "":
                var v133 uint16 = 753
                _ = v133
            case "foo bar":
                if false {
                    v1[(fi132 << ((2) & 63))] = genFn0()
                }
            }
            delete(v2, (fi132 / (((fi132 * (fi132 ^ -9223372036854775808))) | 1)))
        }
        _ = fi132
    }
    v6 = (v6 << ((692) & 63))
    v3 = ("foo bar" + ("foo bar" + ("hello" + "hello")))
    for fi134 := 0; fi134 < 19; fi134++ {
        v5 = (3.14159 * v5)
        if (522 <= 18446744073709551615) {
            v135 := []int{}
            _ = v135
        }
        v6 -= (^(v6))
        _ = fi134
    }
    fmt.Println(v6)
    var v136 uint32 = 2
    v5 = v5
    if ((v136 * v136) != (^(v6))) {
        for fi137 := 0; fi137 < 2; fi137++ {
            if false {
                var v138 uint = 42
                _ = v138
            }
            v136 = 1
            _ = fi137
        }
    }
    v6 -= (v136 - v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v136:", v136)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
