package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 int; F2 int64 }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS8 struct { F0 int64; F1 uint; F2 int64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int; F1 uint16; F2 uint8 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int64; F1 int; F2 int8 }
type GenS15 struct { F0 int64; F1 uint64 }
type GenS16 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 uint8; F1 int; F2 uint32 }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 string; F1 string }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 uint; F1 uint64; F2 int }
type GenS24 struct { F0 int64; F1 int8 }
type GenS25 struct { F0 uint; F1 uint8; F2 uint64 }
func genFn0(p0 int8) uint {
    fmt.Println(p0)
    if true {
        for fi6 := 0; fi6 < 19; fi6++ {
            fmt.Println(fi6)
            _ = fi6
        }
        var v7 uint8 = 255
        v8 := []int{32768, -2147483648, 32767}
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 39}
        var v10 uint32 = 1
        fmt.Println(v10)
        _ = v10
        _ = v9
    }
    v11 := []int{}
    _ = p0
    _ = v11
    return 2
}
func genFn1(p0 int, p1 int64) int {
    v12 := GenS1{F0: -32768}
    if ((0.1 * -2.25) > float32(0.0)) {
        p1 += p1
        var v13 int = 127
        _ = v13
    } else {
        if (42 < 45) {
            v14 := GenS2{F0: 0}
            var v15 int16 = 32767
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 6; fi16++ {
                p0 = (-(fi16))
                p0 = fi16
                fmt.Println(p1, fi16)
                _ = fi16
            }
            var v17 int = 42
            _ = v17
        }
        fmt.Println(p0)
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    _ = v12
    return (p0 % (((^(p0))) | 1))
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 255
    v3 := []uint32{}
    var v4 int8 = 91
    var v5 uint64 = 2
    fmt.Println(v2, v2, v4)
    for fi18 := 0; fi18 < 13; fi18++ {
        for fi19 := 0; fi19 < 6; fi19++ {
            fmt.Println(v5)
            _ = fi19
        }
        _ = fi18
    }
    switch "" {
    case "hello":
        if true {
            var v20 float32 = 3.14159
            if true {
                v2 -= 1
                for fi21 := 0; fi21 < 20; fi21++ {
                    fmt.Println(v20)
                    _ = fi21
                }
                switch -9223372036854775808 {
                case -2147483648:
                    if false {
                        v22 := v1[612]
                        _ = v22
                    }
                default:
                    for fi23 := 0; fi23 < 7; fi23++ {
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(v2, v4, v4)
                        case "a":
                            var v24 uint64 = 4294967295
                            _ = v24
                        case "":
                            v25 := GenS1{F0: 0}
                            _ = v25
                        default:
                            if (v20 >= v20) {
                                v20 = float32(-2.25)
                                var v26 int8 = (v4 ^ (v4 | -17))
                                if (v20 != float32(((1e10 * float64(0.1)) + 0.1))) {
                                    if false {
                                        for fi27 := 0; fi27 < 6; fi27++ {
                                            if (true || (!((^(v4)) < (v4 | (v4 << ((668) & 63)))))) {
                                                fmt.Println(fi23)
                                                var v28 string = "foo bar"
                                                v20 -= ((((v20 + v20) * (v20 / (v20 - (((v20 - ((v20 * 0.0) + (v20 - ((v20 * v20) * -2.25)))) / v20) + 3.14159)))) * (v20 - -2.25)) / v20)
                                                _ = v28
                                            } else {
                                                v29 := GenS2{F0: 791}
                                                _ = v29
                                            }
                                            var v30 float64 = (3.14159 - float64(v20))
                                            var v31 int = fi27
                                            if ((!(25 <= 106)) || (265 > 928)) {
                                                v32 := GenS1{F0: 32767}
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v33 := GenS2{F0: 4294967295}
                                                _ = v32
                                                _ = v33
                                            }
                                            _ = fi27
                                            _ = v30
                                            _ = v31
                                        }
                                    }
                                }
                                _ = v26
                            } else {
                                if (571 > 2147483648) {
                                    for fi34 := 0; fi34 < 2; fi34++ {
                                        v35 := GenS3{F0: true}
                                        _ = fi34
                                        _ = v35
                                    }
                                    switch ("foo bar" + (("hello" + ("" + "")) + "foo bar")) {
                                    case "a":
                                        var v36 int16 = 645
                                        _ = v36
                                    case "\n\t":
                                        var v37 uint64 = (^(v5))
                                        _ = v37
                                    case "hello":
                                        for fi38 := 0; fi38 < 2; fi38++ {
                                            if true {
                                                switch ("\n\t" + "a") {
                                                case "\n\t":
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                case "a":
                                                    if len(v3) > 0 {
                                                        v3[0] = 2
                                                    }
                                                default:
                                                    v39 := GenS1{F0: 42}
                                                    _ = v39
                                                }
                                                fmt.Println(v5, v4, v4)
                                                v4 -= 127
                                            } else {
                                                var v40 float32 = float32((0.0 - -2.25))
                                                _ = v40
                                            }
                                            v41 := v1[fi23]
                                            var v42 uint8 = 46
                                            _ = fi38
                                            _ = v41
                                            _ = v42
                                        }
                                    }
                                    v4 = v4
                                } else {
                                    v5 += v5
                                    if (true || false) {
                                        var v43 string = "hello"
                                        _ = v43
                                    } else {
                                        if true {
                                            v44 := GenS4{F0: v5}
                                            v5 -= v5
                                            var v45 uint32 = 1
                                            _ = v44
                                            _ = v45
                                        } else {
                                            switch "\n\t" {
                                            case "":
                                                v4 += v4
                                            case "a":
                                                switch int(v5) {
                                                case -1:
                                                    v20 += v20
                                                case -644:
                                                    var v46 int = fi23
                                                    _ = v46
                                                default:
                                                    if false {
                                                        v2 = 344
                                                        switch 789 {
                                                        case -863:
                                                            v4 += v4
                                                        case -9223372036854775808:
                                                            for fi47 := 0; fi47 < 6; fi47++ {
                                                                switch genFn1((fi47 ^ fi23), -9223372036854775808) {
                                                                case -229:
                                                                    v20 += 1.5
                                                                    v4 += (v4 * (v4 & (v4 & v4)))
                                                                case -684:
                                                                    fmt.Println(v2)
                                                                    fmt.Println(fi47, fi23)
                                                                    if (507 <= 2) {
                                                                        var v48 int8 = (v4 - (v4 - (v4 + (v4 << ((2) & 63)))))
                                                                        var v49 uint16 = 65535
                                                                        _ = v48
                                                                        _ = v49
                                                                    } else {
                                                                        if true {
                                                                            if true {
                                                                                if false {
                                                                                    v50 := GenS5{F0: 0, F1: (fi47 ^ fi47), F2: 832}
                                                                                    v51 := GenS5{F0: 628, F1: (fi47 / (((fi23 % (((fi47 / ((-43) | 1))) | 1))) | 1)), F2: 42}
                                                                                    v52 := GenS6{F0: v2, F1: false}
                                                                                    v53 := GenS7{F0: v4, F1: 0, F2: (v2 / ((v2) | 1))}
                                                                                    _ = v50
                                                                                    _ = v51
                                                                                    _ = v52
                                                                                    _ = v53
                                                                                }
                                                                            }
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                        } else {
                                                                            if len(v3) > 0 {
                                                                                v3[0] = 2
                                                                            }
                                                                            fmt.Println(fi47)
                                                                            v54 := GenS8{F0: int64(v2), F1: 813, F2: 716}
                                                                            fmt.Println(v4, v5, fi47)
                                                                            v5 = (v5 << ((2) & 63))
                                                                            _ = v54
                                                                        }
                                                                        _ = len(v3)
                                                                    }
                                                                case -1:
                                                                    var v55 uint64 = 0
                                                                    if (true && ("foo bar" > ("\n\t" + ""))) {
                                                                        v5 = v5
                                                                    }
                                                                    _ = v55
                                                                default:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                                v4 += v4
                                                                v4 += v4
                                                                _ = fi47
                                                            }
                                                        case -742:
                                                            v56 := GenS9{F0: 753}
                                                            _ = v56
                                                        }
                                                        switch ("" + "hello") {
                                                        case "hello":
                                                            if (!(true && (false || true))) {
                                                                v5 += v5
                                                                var v57 int32 = -875
                                                                fmt.Println(v4, v4, v5)
                                                                _ = v57
                                                            } else {
                                                                var v58 bool = (v4 == -90)
                                                                _ = v58
                                                            }
                                                        case "a":
                                                            for fi59 := 0; fi59 < 13; fi59++ {
                                                                if (42 >= 9223372036854775807) {
                                                                    fmt.Println(v5, fi59, v5)
                                                                    var v60 bool = (((!false) || (v2 <= (v2 ^ 2))) && true)
                                                                    _ = v60
                                                                } else {
                                                                    v61 := GenS3{F0: false}
                                                                    var v62 int8 = 19
                                                                    _ = v61
                                                                    _ = v62
                                                                }
                                                                _ = fi59
                                                            }
                                                        case "foo bar":
                                                            var v63 int16 = -311
                                                            _ = v63
                                                        default:
                                                            switch fi23 {
                                                            case 1:
                                                                var v64 int = (fi23 - fi23)
                                                                _ = v64
                                                            case 127:
                                                                var v65 uint8 = 79
                                                                _ = v65
                                                            case 2:
                                                                if false {
                                                                    fmt.Println(v2, v20, v5)
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                } else {
                                                                    _ = len(v1)
                                                                }
                                                            default:
                                                                if true {
                                                                    switch genFn1(fi23, 42) {
                                                                    case 32768:
                                                                        _ = len(v3)
                                                                    case 4294967295:
                                                                        v66 := GenS8{F0: 1, F1: 2, F2: -1}
                                                                        _ = v66
                                                                    default:
                                                                        v4 += v4
                                                                    }
                                                                    v67 := GenS5{F0: 849, F1: (fi23 << ((65535) & 63)), F2: 9223372036854775807}
                                                                    _ = v67
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        _ = len(v3)
                                                        for fi68 := 0; fi68 < 2; fi68++ {
                                                            switch "a" {
                                                            case "\n\t":
                                                                var v69 uint16 = v2
                                                                _ = v69
                                                            case "foo bar":
                                                                switch (fi23 + (fi68 << ((genFn0((v4 | (v4 ^ (^(v4)))))) & 63))) {
                                                                case 45:
                                                                    var v70 float64 = 1.5
                                                                    _ = v70
                                                                case 369:
                                                                    var v71 uint = 2
                                                                    _ = v71
                                                                case 0:
                                                                    var v72 string = "hello"
                                                                    _ = v72
                                                                default:
                                                                    if ((((!false) || ((!false) || false)) || true) || ((0.0 - v20) >= v20)) {
                                                                        switch genFn1(genFn1(fi23, -9223372036854775808), 42) {
                                                                        case 355:
                                                                            if true {
                                                                                v4 = v4
                                                                                v73 := GenS10{F0: -1}
                                                                                v74 := GenS4{F0: (^(v5))}
                                                                                _ = v73
                                                                                _ = v74
                                                                            }
                                                                        case 980:
                                                                            v75 := v1[fi68]
                                                                            fmt.Println(v75, v5, v4)
                                                                            var v76 uint32 = 203
                                                                            _ = v75
                                                                            _ = v76
                                                                        }
                                                                        v77 := GenS8{F0: -1, F1: 4294967295, F2: 1}
                                                                        v78 := GenS11{F0: fi23, F1: (v2 * (v2 / ((v2) | 1))), F2: 0}
                                                                        _ = v77
                                                                        _ = v78
                                                                    }
                                                                }
                                                            }
                                                            _ = len(v1)
                                                            var v79 int16 = int16(v4)
                                                            v4 -= v4
                                                            _ = fi68
                                                            _ = v79
                                                        }
                                                        switch "" {
                                                        case "a":
                                                            fmt.Println(v20, fi23)
                                                        case "hello":
                                                            v80 := GenS3{F0: false}
                                                            _ = v80
                                                        default:
                                                            var v81 int8 = 112
                                                            _ = v81
                                                        }
                                                    }
                                                }
                                            }
                                            if ((v5 << ((42) & 63)) >= 42) {
                                                v20 = v20
                                            } else {
                                                if false {
                                                    var v82 uint64 = (v5 ^ v5)
                                                    _ = v82
                                                } else {
                                                    fmt.Println(v2, v4, v5)
                                                    var v83 uint16 = (^(v2))
                                                    if true {
                                                        if true {
                                                            var v84 uint32 = 2
                                                            _ = v84
                                                        } else {
                                                            for fi85 := 0; fi85 < 15; fi85++ {
                                                                v20 -= v20
                                                                var v86 uint32 = 2147483647
                                                                if true {
                                                                    _ = len(v3)
                                                                    var v87 int64 = 42
                                                                    _ = v87
                                                                }
                                                                _ = fi85
                                                                _ = v86
                                                            }
                                                        }
                                                        for fi88 := 0; fi88 < 13; fi88++ {
                                                            switch "a" {
                                                            case "":
                                                                if (2 == 2) {
                                                                    v89 := GenS11{F0: fi88, F1: v83, F2: 255}
                                                                    _ = v89
                                                                }
                                                            default:
                                                                var v90 int32 = 42
                                                                _ = v90
                                                            }
                                                            var v91 float64 = 0.0
                                                            if false {
                                                                var v92 string = "a"
                                                                v93 := GenS12{F0: v5}
                                                                _ = v92
                                                                _ = v93
                                                            } else {
                                                                v91 = v91
                                                                var v94 uint16 = (^(v2))
                                                                _ = v94
                                                            }
                                                            var v95 float64 = v91
                                                            _ = fi88
                                                            _ = v91
                                                            _ = v95
                                                        }
                                                        for fi96 := 0; fi96 < 7; fi96++ {
                                                            v20 = 1e10
                                                            _ = fi96
                                                        }
                                                    }
                                                    _ = v83
                                                }
                                                if true {
                                                    switch 128 {
                                                    case 32767:
                                                        var v97 uint = 2
                                                        _ = v97
                                                    default:
                                                        _ = len(v1)
                                                    }
                                                    if false {
                                                        v98 := GenS5{F0: 18446744073709551615, F1: (fi23 - fi23), F2: 1}
                                                        v2 = 2
                                                        _ = v98
                                                    } else {
                                                        v99 := GenS13{F0: "\n\t"}
                                                        fmt.Println(v20, v5, v20)
                                                        for fi100 := 0; fi100 < 11; fi100++ {
                                                            var v101 uint = 667
                                                            var v102 int64 = 9223372036854775807
                                                            v102 -= v102
                                                            _ = fi100
                                                            _ = v101
                                                            _ = v102
                                                        }
                                                        _ = v99
                                                    }
                                                    v4 -= (v4 / (((v4 & v4)) | 1))
                                                } else {
                                                    for fi103 := 0; fi103 < 2; fi103++ {
                                                        v20 = 1e10
                                                        var v104 uint32 = 843
                                                        _ = fi103
                                                        _ = v104
                                                    }
                                                }
                                            }
                                            v105 := GenS14{F0: -1, F1: genFn1((fi23 | (fi23 * 127)), -9223372036854775808), F2: (^(v4))}
                                            _ = v105
                                        }
                                    }
                                    if ((!false) || false) {
                                        v106 := GenS9{F0: 2}
                                        _ = v106
                                    }
                                }
                            }
                        }
                        for fi107 := 0; fi107 < 20; fi107++ {
                            if (!true) {
                                v108 := GenS11{F0: (fi23 * 42), F1: (v2 | 42), F2: 1}
                                var v109 int32 = -472
                                v5 = 0
                                _ = v108
                                _ = v109
                            } else {
                                if true {
                                    var v110 int8 = -1
                                    _ = v110
                                } else {
                                    v20 = -2.25
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    fmt.Println(fi107, fi107)
                                    switch 127 {
                                    case 128:
                                        switch -2147483648 {
                                        case 9223372036854775807:
                                            var v111 int32 = -319
                                            _ = v111
                                        case -9223372036854775808:
                                            fmt.Println(fi107)
                                            _ = len(v3)
                                            var v112 int64 = 636
                                            _ = v112
                                        case -383:
                                            var v113 float32 = (v20 + v20)
                                            _ = v113
                                        default:
                                            v2 = (v2 & v2)
                                        }
                                    }
                                }
                                v2 -= 255
                                v114 := v1[(fi23 >> ((genFn0(-4)) & 63))]
                                v5 -= (v5 | 42)
                                switch (((("\n\t" + ("hello" + "\n\t")) + (("" + "a") + "")) + ("a" + "\n\t")) + "a") {
                                case "foo bar":
                                    fmt.Println(fi107)
                                case "a":
                                    var v115 bool = v114
                                    _ = v115
                                case "\n\t":
                                    var v116 int32 = 394
                                    _ = v116
                                default:
                                    v114 = (0 == 2147483647)
                                }
                                _ = v114
                            }
                            _ = fi107
                        }
                        v117 := GenS1{F0: -32768}
                        _ = fi23
                        _ = v117
                    }
                }
            } else {
                for fi118 := 0; fi118 < 2; fi118++ {
                    var v119 string = ((("\n\t" + "") + ((("a" + (("foo bar" + ("foo bar" + "")) + "\n\t")) + "hello") + "")) + "foo bar")
                    fmt.Println(v4, v5)
                    _ = fi118
                    _ = v119
                }
                var v120 int8 = 57
                fmt.Println(v5)
                _ = v120
            }
            switch "\n\t" {
            case "a":
                switch (("a" + ("a" + "hello")) + "a") {
                case "hello":
                    _ = len(v1)
                }
            }
            _ = v20
        } else {
            v4 -= 127
            var v121 int64 = 9223372036854775807
            _ = v121
        }
    case "foo bar":
        var v122 int64 = -1
        _ = v122
    case "":
        fmt.Println(v2, v4, v2)
    default:
        switch 0 {
        case 0:
            v123 := GenS4{F0: 79}
            _ = v123
        default:
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v2)
            case "hello":
                var v124 float32 = -2.25
                _ = v124
            case "\n\t":
                v5 += v5
            }
        }
    }
    v4 = -1
    if true {
        var v125 uint8 = 255
        _ = v125
    }
    var v126 string = "foo bar"
    if (!true) {
        var v127 int8 = (v4 ^ v4)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v127
    }
    v128 := GenS15{F0: 0, F1: (v5 % ((v5) | 1))}
    switch -9223372036854775808 {
    case 2147483647:
        fmt.Println(v2, v4, v5)
    default:
        for fi129 := 0; fi129 < 17; fi129++ {
            var v130 int64 = 1
            _ = fi129
            _ = v130
        }
    }
    var v131 float32 = 1.5
    if true {
        switch 0 {
        case 32767:
            if ((v5 ^ v5) != (v5 - (^(v5)))) {
                for fi132 := 0; fi132 < 1; fi132++ {
                    switch 1 {
                    case 268:
                        if false {
                            v126 = "hello"
                            _ = len(v1)
                            v131 = 0.0
                        }
                    case -1:
                        switch genFn1(int(v2), int64(v5)) {
                        case 666:
                            if true {
                                v133 := GenS16{F0: (v2 / ((938) | 1)), F1: 18446744073709551615, F2: 42}
                                fmt.Println(v131, v126)
                                switch (fi132 / ((fi132) | 1)) {
                                case 1:
                                    if true {
                                        v134 := GenS13{F0: ("\n\t" + ((v126 + v126) + (v126 + (v126 + v126))))}
                                        var v135 uint32 = 2147483647
                                        _ = v134
                                        _ = v135
                                    }
                                case 32768:
                                    for fi136 := 0; fi136 < 19; fi136++ {
                                        v137 := GenS17{F0: 1}
                                        v4 = v4
                                        _ = fi136
                                        _ = v137
                                    }
                                case 0:
                                    for fi138 := 0; fi138 < 20; fi138++ {
                                        fmt.Println(v126)
                                        var v139 uint32 = 32
                                        v4 += 127
                                        v140 := GenS18{F0: uint8(v2), F1: (fi138 + (fi138 & (fi132 / ((genFn1(genFn1((-(fi138)), -9223372036854775808), 194)) | 1)))), F2: (v139 << ((genFn0(v4)) & 63))}
                                        _ = fi138
                                        _ = v139
                                        _ = v140
                                    }
                                }
                                _ = v133
                            }
                        case -2147483648:
                            var v141 uint64 = 4294967295
                            _ = v141
                        case 2147483647:
                            var v142 float32 = v131
                            _ = v142
                        default:
                            fmt.Println(v126)
                        }
                    default:
                        var v143 bool = (!(!false))
                        _ = v143
                    }
                    if false {
                        if false {
                            v2 -= (v2 << ((4294967295) & 63))
                            v131 = 1e10
                        }
                        var v144 uint64 = v5
                        _ = v144
                    }
                    _ = fi132
                }
                if (!false) {
                    v126 = "hello"
                    fmt.Println(v2, v2)
                }
                v4 = (v4 - v4)
            } else {
                v2 += 32767
                v145 := GenS15{F0: 1, F1: 42}
                v2 -= 0
                _ = v145
            }
        case 42:
            if ((2147483648 <= 2) || (-849 <= 0)) {
                var v146 int32 = -2147483648
                for fi147 := 0; fi147 < 17; fi147++ {
                    var v148 float64 = (0.0 * 3.14159)
                    v146 = v146
                    fmt.Println(v131)
                    _ = fi147
                    _ = v148
                }
                _ = v146
            }
        }
    } else {
        if (true || (!true)) {
            fmt.Println(v5, v5)
            var v149 bool = true
            _ = v149
        }
        v3 = append(v3, 0)
    }
    var v150 int64 = 1
    v5 = (^(v5))
    v4 = v4
    fmt.Println(v5, v5)
    if (!((((v126 + v126) + v126) > ((("\n\t" + "hello") + "\n\t") + ("" + v126))) || (true || (v131 != v131)))) {
        for fi151 := 0; fi151 < 13; fi151++ {
            switch "a" {
            case "a":
                switch v126 {
                case "":
                    v5 -= 4294967295
                case "a":
                    if (!((-2.25 * 0.1) == 0.1)) {
                        var v152 float64 = -2.25
                        if len(v3) > 0 {
                            v3[0] = 65535
                        }
                        _ = v152
                    } else {
                        fmt.Println(v4, v5)
                    }
                case "\n\t":
                    v153 := v1[-9223372036854775808]
                    _ = v153
                default:
                    if true {
                        v1[fi151] = false
                        switch (fi151 << ((826) & 63)) {
                        case -1:
                            fmt.Println(v131, v5)
                        case 4294967295:
                            switch ((v126 + "foo bar") + v126) {
                            case "foo bar":
                                var v154 bool = (42 != -778)
                                _ = v154
                            case "\n\t":
                                v3 = append(v3, 65535)
                            }
                        case 159:
                            v5 -= (^(v5))
                        }
                        if (true || false) {
                            var v155 int16 = -32768
                            if (0.0 == 0.0) {
                                for fi156 := 0; fi156 < 13; fi156++ {
                                    var v157 int64 = (v150 + (v150 + v150))
                                    v155 = 347
                                    var v158 int16 = (v155 + v155)
                                    _ = fi156
                                    _ = v157
                                    _ = v158
                                }
                                var v159 string = "hello"
                                var v160 uint16 = (v2 | (v2 >> ((405) & 63)))
                                _ = v159
                                _ = v160
                            }
                            _ = v155
                        }
                    } else {
                        v161 := v1[fi151]
                        v5 -= (v5 >> ((0) & 63))
                        _ = v161
                    }
                }
            case "foo bar":
                v162 := GenS1{F0: -655}
                _ = v162
            case "hello":
                v150 = (v150 & (v150 | int64(v4)))
            }
            _ = fi151
        }
        switch int(v5) {
        case 806:
            if (2 < 2) {
                v163 := GenS4{F0: (v5 & (v5 << ((65535) & 63)))}
                v126 = ((v126 + (v126 + v126)) + ((v126 + v126) + "hello"))
                _ = v163
            } else {
                if false {
                    switch 42 {
                    case 467:
                        var v164 uint = 18446744073709551615
                        _ = v164
                    case -628:
                        switch 0 {
                        case 9223372036854775807:
                            switch v126 {
                            case "\n\t":
                                var v165 float64 = ((float64(v131) * 0.1) * -2.25)
                                _ = v165
                            }
                        default:
                            v166 := v1[128]
                            _ = v166
                        }
                    }
                    v2 = (v2 & 0)
                    v126 = v126
                }
                v167 := GenS19{F0: 213}
                _ = v167
            }
        }
        if true {
            switch v126 {
            case "a":
                var v168 uint16 = 151
                _ = v168
            case "":
                var v169 uint32 = 2147483648
                _ = v169
            }
            switch v126 {
            case "hello":
                v2 -= 833
            case "foo bar":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "":
                if true {
                    for fi170 := 0; fi170 < 20; fi170++ {
                        fmt.Println(v5, v5, v150)
                        v171 := GenS20{F0: (v126 + (((("hello" + (((("foo bar" + ((v126 + ((v126 + v126) + v126)) + (v126 + v126))) + v126) + "a") + v126)) + v126) + ((v126 + v126) + v126)) + v126)), F1: ((("" + (v126 + "")) + v126) + "a")}
                        v172 := GenS21{F0: (v4 % (((-(v4))) | 1))}
                        _ = fi170
                        _ = v171
                        _ = v172
                    }
                }
            default:
                switch (v126 + "") {
                case "\n\t":
                    var v173 uint8 = 1
                    _ = v173
                case "":
                    switch 127 {
                    case -607:
                        if true {
                            switch genFn1(32768, 9223372036854775807) {
                            case 9223372036854775807:
                                v4 += v4
                            case 1:
                                var v174 int64 = (v150 >> ((18446744073709551615) & 63))
                                _ = v174
                            case -9223372036854775808:
                                v175 := GenS12{F0: v5}
                                _ = v175
                            default:
                                var v176 uint16 = (v2 | v2)
                                _ = v176
                            }
                            if false {
                                v177 := GenS12{F0: (^(v5))}
                                _ = v177
                            }
                        }
                    case 127:
                        fmt.Println(v131, v126)
                    }
                case "hello":
                    for fi178 := 0; fi178 < 18; fi178++ {
                        for fi179 := 0; fi179 < 7; fi179++ {
                            v180 := GenS19{F0: 42}
                            _ = len(v1)
                            if (true || (1000 >= -940)) {
                                if false {
                                    var v181 uint = 4294967295
                                    v131 -= 1e10
                                    _ = v181
                                } else {
                                    v126 = (v126 + ("foo bar" + v126))
                                    v131 = v131
                                    if false {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        fmt.Println(fi178, v4)
                                        v1[4294967295] = true
                                        if (false && (!false)) {
                                            var v182 int64 = 9223372036854775807
                                            var v183 int8 = (v4 + v4)
                                            var v184 string = (v126 + ("foo bar" + "hello"))
                                            _ = v182
                                            _ = v183
                                            _ = v184
                                        } else {
                                            var v185 string = (v126 + v126)
                                            var v186 int8 = v4
                                            switch fi178 {
                                            case 852:
                                                v185 = (("a" + v126) + "hello")
                                            case 32767:
                                                v185 = "\n\t"
                                                var v187 uint = 0
                                                var v188 int64 = 42
                                                _ = v187
                                                _ = v188
                                            case 379:
                                                v4 = -1
                                            }
                                            _ = v185
                                            _ = v186
                                        }
                                    } else {
                                        if (v150 >= (v150 + v150)) {
                                            v189 := v1[fi178]
                                            var v190 string = v126
                                            var v191 uint32 = uint32(v2)
                                            _ = v189
                                            _ = v190
                                            _ = v191
                                        }
                                        var v192 int = 0
                                        v3 = append(v3, 65535)
                                        fmt.Println(v5, v5, v131)
                                        switch "foo bar" {
                                        case "":
                                            switch fi178 {
                                            case 276:
                                                var v193 int8 = (v4 / ((116) | 1))
                                                _ = v193
                                            default:
                                                fmt.Println(v4, v192)
                                            }
                                        case "a":
                                            if ((v5 >= (v5 + v5)) && (1e10 != 1e10)) {
                                                v194 := GenS22{F0: false}
                                                var v195 uint64 = v5
                                                switch "foo bar" {
                                                case "\n\t":
                                                    fmt.Println(v126, fi179)
                                                default:
                                                    var v196 uint = 115
                                                    _ = v196
                                                }
                                                _ = v194
                                                _ = v195
                                            } else {
                                                fmt.Println(fi179, fi179)
                                                v131 -= v131
                                            }
                                        case "hello":
                                            _ = len(v1)
                                        default:
                                            var v197 int = fi178
                                            _ = v197
                                        }
                                        switch 9223372036854775807 {
                                        case 1:
                                            v198 := GenS17{F0: 84}
                                            _ = v198
                                        default:
                                            var v199 int8 = v4
                                            _ = v199
                                        }
                                        _ = v192
                                    }
                                }
                            } else {
                                _ = len(v3)
                            }
                            var v200 float64 = (3.14159 * 0.1)
                            switch fi179 {
                            case -2147483648:
                                if (2147483647 == 2147483647) {
                                    _ = len(v3)
                                } else {
                                    if (v150 >= 42) {
                                        _ = len(v1)
                                        v2 -= (^(v2))
                                        v201 := GenS23{F0: genFn0((v4 << ((genFn0(v4)) & 63))), F1: 439, F2: 32767}
                                        _ = v201
                                    }
                                    var v202 bool = true
                                    _ = v202
                                }
                            case 32767:
                                var v203 int8 = (v4 >> ((4294967295) & 63))
                                _ = v203
                            }
                            _ = fi179
                            _ = v180
                            _ = v200
                        }
                        _ = fi178
                    }
                default:
                    var v204 int8 = v4
                    _ = v204
                }
            }
        }
    } else {
        if (809 < 42) {
            switch genFn1(127, v150) {
            case 886:
                for fi205 := 0; fi205 < 20; fi205++ {
                    fmt.Println(v2)
                    _ = fi205
                }
            default:
                for fi206 := 0; fi206 < 13; fi206++ {
                    if (!false) {
                        for fi207 := 0; fi207 < 8; fi207++ {
                            switch ("foo bar" + ("foo bar" + v126)) {
                            case "a":
                                fmt.Println(v126)
                            case "\n\t":
                                var v208 int64 = (v150 % ((v150) | 1))
                                switch fi206 {
                                case 42:
                                    v209 := GenS3{F0: (!(false && false))}
                                    _ = v209
                                }
                                var v210 int64 = (v150 + v150)
                                switch (v126 + v126) {
                                case "\n\t":
                                    if false {
                                        var v211 uint32 = 716
                                        _ = v211
                                    } else {
                                        switch 2147483647 {
                                        case 983:
                                            var v212 bool = true
                                            _ = v212
                                        case -2147483648:
                                            v4 += v4
                                        case 42:
                                            var v213 int8 = 71
                                            _ = v213
                                        default:
                                            var v214 int64 = v208
                                            _ = v214
                                        }
                                    }
                                case "a":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "hello":
                                    var v215 uint8 = 150
                                    _ = v215
                                }
                                v5 = v5
                                switch (("hello" + "") + (v126 + (v126 + "\n\t"))) {
                                case "":
                                    var v216 int32 = -2147483648
                                    _ = v216
                                case "foo bar":
                                    v217 := GenS17{F0: 1}
                                    _ = v217
                                default:
                                    fmt.Println(fi207, fi207)
                                }
                                _ = v208
                                _ = v210
                            }
                            if (927 < 18446744073709551615) {
                                var v218 uint64 = (v5 ^ v5)
                                v131 = ((float32(1e10) + (0.1 / v131)) - v131)
                                if false {
                                    v126 = v126
                                }
                                _ = v218
                            }
                            v219 := GenS24{F0: v150, F1: (v4 >> ((genFn0((v4 << ((18446744073709551615) & 63)))) & 63))}
                            var v220 uint = 255
                            _ = fi207
                            _ = v219
                            _ = v220
                        }
                        v221 := GenS25{F0: genFn0(55), F1: 127, F2: uint64(v2)}
                        _ = len(v1)
                        _ = v221
                    }
                    v5 += (v5 >> ((1) & 63))
                    _ = fi206
                }
            }
        } else {
            fmt.Println(v4)
            v222 := GenS21{F0: 1}
            _ = v222
        }
        var v223 float64 = 1.5
        _ = v223
    }
    delete(v1, 42)
    switch 1 {
    case -411:
        if (2 == 18446744073709551615) {
            _ = len(v1)
        } else {
            v224 := v1[1]
            _ = v224
        }
    case 2147483647:
        delete(v1, -9223372036854775808)
    default:
        fmt.Println(v126, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v126:", v126)
    fmt.Println("v131:", v131)
    fmt.Println("v150:", v150)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v126
    _ = v128
    _ = v131
    _ = v150
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
