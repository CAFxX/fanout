package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 int32; F2 bool }
type GenS5 struct { F0 int64; F1 uint8; F2 string }
type GenS6 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 string; F1 int64; F2 uint32 }
type GenS11 struct { F0 int8; F1 uint8 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint16; F1 int16; F2 uint64 }
func genFn0() uint16 {
    var v4 int64 = 9223372036854775807
    var v5 uint64 = 629
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 string = "hello"
        var v8 float64 = 0.1
        _ = fi6
        _ = v7
        _ = v8
    }
    if (!false) {
        v9 := GenS1{F0: 56, F1: v4}
        var v10 uint16 = 1
        v10 += (v10 * v10)
        _ = v10
        _ = v9
    }
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 int8 = -44
    var v2 int16 = 1000
    var v3 int8 = (v1 | v1)
    v3 -= (v3 * (^(v3)))
    switch "a" {
    case "hello":
        var v11 uint32 = 219
        _ = v11
    case "foo bar":
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 int32 = 2147483647
            v14 := GenS1{F0: v1, F1: 437}
            v15 := []int{(fi12 | (^(fi12))), 9223372036854775807, fi12}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    case "a":
        if (!true) {
            v2 = v2
            if (255 > 18446744073709551615) {
                v1 = (v1 >> ((4294967295) & 63))
                for fi16 := 0; fi16 < 6; fi16++ {
                    v2 = v2
                    fmt.Println(fi16, v2)
                    for fi17 := 0; fi17 < 19; fi17++ {
                        v18 := GenS1{F0: (v3 & 82), F1: 1}
                        _ = fi17
                        _ = v18
                    }
                    _ = fi16
                }
                v1 = 42
            } else {
                for fi19 := 0; fi19 < 8; fi19++ {
                    v3 = 121
                    v2 += 0
                    _ = fi19
                }
                for fi20 := 0; fi20 < 10; fi20++ {
                    fmt.Println(v2, v2)
                    v3 = 0
                    _ = fi20
                }
                v21 := GenS1{F0: v3, F1: 1}
                _ = v21
            }
            fmt.Println(v2)
        } else {
            v22 := []int{0}
            _ = v22
        }
    default:
        v3 = v3
    }
    switch -14 {
    case 127:
        v23 := []int{}
        _ = v23
    case 1:
        v3 = (v1 | v3)
    }
    if (int32(v2) < -2147483648) {
        v24 := []int{32767}
        if true {
            var v25 string = ("a" + "hello")
            v26 := GenS1{F0: (v3 / (((v3 / ((42) | 1))) | 1)), F1: -9223372036854775808}
            _ = v25
            _ = v26
        }
        switch "foo bar" {
        case "hello":
            var v27 int = 337
            _ = v27
        case "\n\t":
            var v28 uint = 18446744073709551615
            _ = v28
        case "":
            v1 = -1
        default:
            for fi29 := 0; fi29 < 19; fi29++ {
                if true {
                    var v30 uint32 = 42
                    v3 -= (-(v1))
                    _ = v30
                }
                if len(v24) > 0 {
                    v24[0] = (fi29 & (fi29 / ((1) | 1)))
                }
                var v31 int = (-(fi29))
                _ = fi29
                _ = v31
            }
        }
        _ = v24
    } else {
        if true {
            v1 -= v3
            if (!false) {
                if (!(!(false || (-2.25 != float64((0.1 * 3.14159)))))) {
                    switch "foo bar" {
                    case "hello":
                        v3 = (v1 % (((v3 >> ((0) & 63))) | 1))
                    }
                    for fi32 := 0; fi32 < 4; fi32++ {
                        v33 := GenS2{F0: 42, F1: (v3 | v3)}
                        switch ("a" + "") {
                        case "a":
                            for fi34 := 0; fi34 < 7; fi34++ {
                                v35 := []int{fi34, (fi34 & fi34), fi34}
                                _ = fi34
                                _ = v35
                            }
                        case "\n\t":
                            v3 -= (v1 ^ 1)
                        case "":
                            v3 += v1
                        }
                        _ = fi32
                        _ = v33
                    }
                }
                if (false || (644 > -939)) {
                    var v36 float32 = 0.0
                    _ = v36
                }
            } else {
                v1 = v1
                v3 = (v1 / (((v3 << ((0) & 63))) | 1))
                for fi37 := 0; fi37 < 2; fi37++ {
                    v38 := GenS2{F0: 42, F1: v3}
                    v1 += v3
                    var v39 int8 = (v3 & (^(v3)))
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
            var v40 int32 = -1
            _ = v40
        } else {
            v41 := []int{}
            if false {
                var v42 uint16 = 32767
                for fi43 := 0; fi43 < 18; fi43++ {
                    var v44 uint32 = 2147483647
                    var v45 int32 = 766
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                _ = v42
            } else {
                v46 := GenS3{F0: genFn0()}
                v3 = (v3 | v3)
                if (42 <= 445) {
                    v47 := GenS4{F0: true, F1: 1, F2: (!true)}
                    var v48 bool = true
                    for fi49 := 0; fi49 < 18; fi49++ {
                        var v50 int8 = 100
                        switch ("foo bar" + "foo bar") {
                        case "hello":
                            var v51 uint8 = 255
                            _ = v51
                        }
                        _ = fi49
                        _ = v50
                    }
                    _ = v47
                    _ = v48
                } else {
                    if true {
                        v52 := GenS3{F0: 0}
                        v1 += (v3 ^ v3)
                        _ = v52
                    } else {
                        if (65535 >= 2) {
                            if (false || (!(!(!true)))) {
                                switch 373 {
                                case -2147483648:
                                    v2 = (v2 / (((v2 * v2)) | 1))
                                }
                                for fi53 := 0; fi53 < 8; fi53++ {
                                    v3 -= 0
                                    fmt.Println(v2, fi53)
                                    _ = fi53
                                }
                                if ((-2.25 >= 0.0) || true) {
                                    for fi54 := 0; fi54 < 12; fi54++ {
                                        if ("hello" != "\n\t") {
                                            var v55 float64 = 0.1
                                            _ = v55
                                        } else {
                                            fmt.Println(v2, fi54, v3)
                                            if len(v41) >= 2 {
                                                v41 = v41[1:len(v41)]
                                            }
                                        }
                                        if false {
                                            v3 = (v3 | v1)
                                            if true {
                                                v1 -= 1
                                            }
                                            var v56 uint16 = 40
                                            _ = v56
                                        } else {
                                            if false {
                                                if (!("hello" >= ("hello" + ("hello" + ("foo bar" + "foo bar"))))) {
                                                    fmt.Println(fi54, v2)
                                                } else {
                                                    v1 += 61
                                                }
                                                var v57 float64 = float64(3.14159)
                                                _ = v57
                                            }
                                            fmt.Println(v1, fi54, v2)
                                        }
                                        _ = fi54
                                    }
                                    v41 = append(v41, 42)
                                    fmt.Println(v3, v2, v1)
                                } else {
                                    for fi58 := 0; fi58 < 2; fi58++ {
                                        for fi59 := 0; fi59 < 16; fi59++ {
                                            if ((v2 - v2) <= (-(v2))) {
                                                if (true && ((!false) || false)) {
                                                    var v60 float32 = 0.1
                                                    var v61 bool = true
                                                    if v61 {
                                                        var v62 int16 = (v2 % (((-(v2))) | 1))
                                                        var v63 int8 = v1
                                                        var v64 float64 = 0.1
                                                        v65 := GenS3{F0: 195}
                                                        switch fi58 {
                                                        case 707:
                                                            var v66 bool = v61
                                                            _ = v66
                                                        }
                                                        v1 -= 1
                                                        fmt.Println(fi58)
                                                        switch "\n\t" {
                                                        case "":
                                                            if (1e10 <= v64) {
                                                                var v67 float64 = (-2.25 + float64(v60))
                                                                var v68 int64 = 0
                                                                switch 127 {
                                                                case -9223372036854775808:
                                                                    var v69 float64 = 3.14159
                                                                    _ = v69
                                                                case 128:
                                                                    v70 := GenS5{F0: -9223372036854775808, F1: 127, F2: "hello"}
                                                                    _ = v70
                                                                case 751:
                                                                    v71 := GenS2{F0: int32(fi59), F1: v63}
                                                                    _ = v71
                                                                }
                                                                var v72 string = ""
                                                                _ = v67
                                                                _ = v68
                                                                _ = v72
                                                            }
                                                        case "a":
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v63 = v3
                                                            case "":
                                                                fmt.Println(v1, v3)
                                                            case "hello":
                                                                v64 -= v64
                                                                var v73 int8 = 42
                                                                _ = v73
                                                            }
                                                        }
                                                        var v74 int8 = v1
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                        _ = v65
                                                        _ = v74
                                                    } else {
                                                        var v75 int64 = 9223372036854775807
                                                        v60 = -2.25
                                                        fmt.Println(v61, v75, v61)
                                                        _ = v75
                                                    }
                                                    _ = v60
                                                    _ = v61
                                                } else {
                                                    var v76 int = (fi59 + (fi59 % (((fi59 >> ((291) & 63))) | 1)))
                                                    _ = v76
                                                }
                                            }
                                            fmt.Println(v2)
                                            var v77 int = -892
                                            var v78 uint16 = 0
                                            if (!true) {
                                                if (true || false) {
                                                    var v79 uint16 = (v78 << ((820) & 63))
                                                    v1 += 91
                                                    _ = v79
                                                } else {
                                                    if len(v41) > 0 {
                                                        v41[0] = fi59
                                                    }
                                                }
                                                switch (("foo bar" + "") + ("a" + "\n\t")) {
                                                case "hello":
                                                    var v80 uint64 = 2
                                                    _ = v80
                                                case "foo bar":
                                                    v3 = (^(v3))
                                                }
                                                if ((^(v78)) >= genFn0()) {
                                                    switch fi58 {
                                                    case -356:
                                                        v81 := GenS6{F0: 42, F1: 187, F2: 255}
                                                        _ = v81
                                                    case 509:
                                                        fmt.Println(fi59)
                                                    }
                                                    var v82 float64 = 3.14159
                                                    _ = v82
                                                } else {
                                                    var v83 int64 = -959
                                                    var v84 bool = false
                                                    _ = v83
                                                    _ = v84
                                                }
                                            }
                                            var v85 uint32 = 544
                                            var v86 int = (-(fi58))
                                            _ = fi59
                                            _ = v77
                                            _ = v78
                                            _ = v85
                                            _ = v86
                                        }
                                        _ = fi58
                                    }
                                    if false {
                                        for fi87 := 0; fi87 < 2; fi87++ {
                                            var v88 float32 = (0.1 + -2.25)
                                            var v89 int16 = 42
                                            if true {
                                                v88 = (v88 + ((v88 / 3.14159) / ((float32(0.1) * (1.5 - -2.25)) - 0.0)))
                                            } else {
                                                if true {
                                                    var v90 int = -1
                                                    v91 := GenS1{F0: (-(v3)), F1: -1}
                                                    _ = v90
                                                    _ = v91
                                                } else {
                                                    for fi92 := 0; fi92 < 2; fi92++ {
                                                        v93 := GenS5{F0: -9223372036854775808, F1: 127, F2: "\n\t"}
                                                        v2 -= v2
                                                        v3 = (v1 - int8(fi92))
                                                        _ = fi92
                                                        _ = v93
                                                    }
                                                    switch (("\n\t" + "a") + "\n\t") {
                                                    case "a":
                                                        v41 = append(v41, fi87)
                                                    case "hello":
                                                        v94 := GenS4{F0: (false || (20 > 639)), F1: 580, F2: (42 != -951)}
                                                        _ = v94
                                                    case "\n\t":
                                                        v1 = v3
                                                    }
                                                    if (!(2 != 581)) {
                                                        if false {
                                                            v95 := GenS7{F0: 3, F1: v89}
                                                            switch "a" {
                                                            case "":
                                                                v89 = (v2 / ((v2) | 1))
                                                            case "foo bar":
                                                                var v96 uint64 = 4294967295
                                                                _ = v96
                                                            case "a":
                                                                v88 = (v88 + float32(3.14159))
                                                            default:
                                                                switch "hello" {
                                                                case "":
                                                                    v41 = append(v41, (fi87 % ((fi87) | 1)))
                                                                case "\n\t":
                                                                    v97 := GenS8{F0: 489}
                                                                    _ = v97
                                                                case "foo bar":
                                                                    for fi98 := 0; fi98 < 18; fi98++ {
                                                                        var v99 int = (fi87 - fi87)
                                                                        switch "a" {
                                                                        case "hello":
                                                                            var v100 bool = (true && (0.1 <= (v88 + v88)))
                                                                            _ = v100
                                                                        case "\n\t":
                                                                            _ = len(v41)
                                                                        case "a":
                                                                            v3 = v3
                                                                            v99 += (fi87 / (((v99 % (((fi98 - 4294967295)) | 1))) | 1))
                                                                        }
                                                                        if (!((1 < 391) || true)) {
                                                                            v101 := GenS5{F0: -9223372036854775808, F1: 255, F2: "foo bar"}
                                                                            v102 := GenS9{F0: 1}
                                                                            v2 += v89
                                                                            _ = v101
                                                                            _ = v102
                                                                        } else {
                                                                            if false {
                                                                                var v103 bool = false
                                                                                v104 := GenS1{F0: (^(v1)), F1: 9223372036854775807}
                                                                                var v105 int32 = -437
                                                                                _ = v103
                                                                                _ = v104
                                                                                _ = v105
                                                                            }
                                                                            if false {
                                                                                _ = len(v41)
                                                                                var v106 uint32 = 2147483648
                                                                                if (32767 <= 544) {
                                                                                    v1 += -23
                                                                                    v107 := GenS10{F0: "\n\t", F1: 42, F2: v106}
                                                                                    if false {
                                                                                        if len(v41) > 0 {
                                                                                            v41[0] = (fi98 | (fi98 - v99))
                                                                                        }
                                                                                    }
                                                                                    _ = v107
                                                                                }
                                                                                _ = v106
                                                                            } else {
                                                                                v1 -= v3
                                                                            }
                                                                        }
                                                                        _ = fi98
                                                                        _ = v99
                                                                    }
                                                                default:
                                                                    var v108 uint8 = 127
                                                                    _ = v108
                                                                }
                                                            }
                                                            v41 = append(v41, (fi87 >> ((1) & 63)))
                                                            _ = v95
                                                        } else {
                                                            fmt.Println(v2, fi87, v2)
                                                        }
                                                        v2 = (^(v2))
                                                        for fi109 := 0; fi109 < 3; fi109++ {
                                                            var v110 int = (fi109 * (fi87 + fi109))
                                                            _ = fi109
                                                            _ = v110
                                                        }
                                                    }
                                                }
                                                var v111 uint32 = 1
                                                _ = v111
                                            }
                                            _ = fi87
                                            _ = v88
                                            _ = v89
                                        }
                                        v112 := GenS11{F0: 100, F1: 2}
                                        var v113 int16 = v2
                                        _ = v112
                                        _ = v113
                                    }
                                }
                            } else {
                                fmt.Println(v1, v2)
                            }
                            if len(v41) > 0 {
                                v41[0] = -349
                            }
                            if len(v41) > 0 {
                                v41[0] = 0
                            }
                        } else {
                            v114 := GenS10{F0: "hello", F1: 0, F2: 654}
                            v115 := GenS6{F0: 0, F1: 668, F2: 42}
                            if (1 >= 148) {
                                v41 = append(v41, 281)
                            }
                            _ = v114
                            _ = v115
                        }
                        var v116 int64 = -358
                        _ = v116
                    }
                    v117 := GenS12{F0: 61}
                    _ = v117
                }
                _ = v46
            }
            _ = v41
        }
    }
    switch -9223372036854775808 {
    case 0:
        for fi118 := 0; fi118 < 12; fi118++ {
            if (4294967295 != 341) {
                if true {
                    v119 := []int{(fi118 * fi118)}
                    _ = v119
                }
            } else {
                for fi120 := 0; fi120 < 11; fi120++ {
                    switch fi120 {
                    case 1:
                        v121 := GenS5{F0: -657, F1: 221, F2: "foo bar"}
                        _ = v121
                    case -9223372036854775808:
                        v2 = 32767
                    case 0:
                        fmt.Println(fi118)
                        switch "foo bar" {
                        case "hello":
                            v122 := []int{(^(fi120)), (fi118 + fi120), (fi120 * -393)}
                            _ = v122
                        default:
                            v123 := []int{(fi120 ^ fi120)}
                            _ = v123
                        }
                        fmt.Println(v1, v3, fi118)
                        fmt.Println(v1)
                    }
                    v124 := []int{fi120, (fi118 & (fi120 % ((fi118) | 1))), (fi120 - fi120), (fi120 & fi118)}
                    _ = fi120
                    _ = v124
                }
                v125 := GenS3{F0: 806}
                _ = v125
            }
            v126 := []int{(fi118 | (fi118 % ((2147483647) | 1)))}
            _ = fi118
            _ = v126
        }
    case 32768:
        v2 += v2
    case -2147483648:
        v127 := GenS8{F0: 863}
        _ = v127
    }
    var v128 uint = 255
    var v129 uint8 = 42
    v130 := GenS13{F0: 2, F1: v2, F2: 42}
    if true {
        v131 := []int{95, -1, -889, -1}
        _ = v131
    } else {
        v129 += (v129 % ((v129) | 1))
        var v132 uint = (^(v128))
        _ = v132
    }
    v1 -= v1
    var v133 uint32 = 65535
    var v134 int32 = -523
    v133 -= v133
    var v135 int = 393
    fmt.Println("v1:", v1)
    fmt.Println("v128:", v128)
    fmt.Println("v129:", v129)
    fmt.Println("v133:", v133)
    fmt.Println("v134:", v134)
    fmt.Println("v135:", v135)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v128
    _ = v129
    _ = v130
    _ = v133
    _ = v134
    _ = v135
    _ = v2
    _ = v3
}
