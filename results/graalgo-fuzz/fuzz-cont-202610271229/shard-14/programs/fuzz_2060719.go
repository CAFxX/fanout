package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int64; F1 int64; F2 uint32 }
type GenS11 struct { F0 int8; F1 int32 }
type GenS12 struct { F0 uint8 }

func main() {
    var v1 uint16 = 65535
    var v2 int16 = 32767
    var v3 string = "\n\t"
    var v4 uint32 = 274
    for fi5 := 0; fi5 < 17; fi5++ {
        fmt.Println(v4, v1, v2)
        v6 := []int{}
        switch (fi5 - fi5) {
        case 42:
            if (false || true) {
                var v7 int8 = -18
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v2 += v2
                _ = v7
            } else {
                var v8 int32 = 507
                _ = v8
            }
        }
        _ = fi5
        _ = v6
    }
    var v9 float32 = (1e10 / 1e10)
    if true {
        v10 := []int{32767, int(v4)}
        var v11 uint16 = (v1 ^ (v1 ^ v1))
        v12 := GenS1{F0: 1, F1: v2}
        _ = v10
        _ = v11
        _ = v12
    } else {
        v13 := GenS2{F0: 1}
        v14 := []int{32768, 128, 4294967295, 127}
        _ = v13
        _ = v14
    }
    switch "hello" {
    case "a":
        v9 = v9
    case "":
        if ((^(v4)) < v4) {
            var v15 string = v3
            var v16 uint = 255
            _ = v15
            _ = v16
        } else {
            v9 += (v9 * (v9 + (1.5 + 1.5)))
        }
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        if (false && (0 == 42)) {
            switch (-(fi17)) {
            case 910:
                v18 := GenS2{F0: -1}
                _ = v18
            case 127:
                switch (fi17 & (fi17 + fi17)) {
                case -594:
                    v19 := GenS1{F0: uint8(v4), F1: (v2 / ((v2) | 1))}
                    _ = v19
                default:
                    var v20 uint16 = v1
                    _ = v20
                }
            default:
                var v21 uint64 = 2
                _ = v21
            }
            if true {
                for fi22 := 0; fi22 < 1; fi22++ {
                    var v23 uint8 = 42
                    v24 := GenS2{F0: -81}
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                v25 := GenS2{F0: -178}
                v26 := []int{(fi17 + -379)}
                _ = v25
                _ = v26
            } else {
                var v27 int8 = -128
                if (-2147483648 <= -503) {
                    if ((v27 ^ (v27 ^ (^(v27)))) >= v27) {
                        v28 := []int{(fi17 * fi17)}
                        v28 = append(v28, -1)
                        _ = v28
                    } else {
                        var v29 int64 = -731
                        switch v3 {
                        case "\n\t":
                            switch fi17 {
                            case 448:
                                fmt.Println(v27)
                            default:
                                switch v3 {
                                case "hello":
                                    if ((v1 % ((32767) | 1)) == v1) {
                                        var v30 int32 = 0
                                        v31 := GenS2{F0: (v30 & v30)}
                                        for fi32 := 0; fi32 < 4; fi32++ {
                                            switch v3 {
                                            case "":
                                                v33 := []int{fi17, (fi17 >> ((uint(v4)) & 63)), (fi17 / (((fi32 >> ((247) & 63))) | 1))}
                                                _ = v33
                                            case "hello":
                                                v1 = 65535
                                            case "\n\t":
                                                var v34 int = (fi32 / (((fi32 >> ((65535) & 63))) | 1))
                                                _ = v34
                                            }
                                            _ = fi32
                                        }
                                        _ = v30
                                        _ = v31
                                    } else {
                                        for fi35 := 0; fi35 < 10; fi35++ {
                                            fmt.Println(v29)
                                            switch (^(fi35)) {
                                            case -9223372036854775808:
                                                v36 := GenS1{F0: 255, F1: v2}
                                                _ = v36
                                            default:
                                                var v37 bool = (-954 == 870)
                                                _ = v37
                                            }
                                            v38 := GenS3{F0: 255}
                                            switch (fi17 << ((809) & 63)) {
                                            case -529:
                                                switch fi35 {
                                                case 1:
                                                    var v39 uint64 = 0
                                                    _ = v39
                                                default:
                                                    v40 := []int{4294967295, int(v29), (fi35 | fi35), fi17}
                                                    _ = v40
                                                }
                                            }
                                            v9 -= 3.14159
                                            v3 = "hello"
                                            v4 += v4
                                            _ = fi35
                                            _ = v38
                                        }
                                        for fi41 := 0; fi41 < 9; fi41++ {
                                            v27 = (v27 & (v27 + (v27 * -1)))
                                            v42 := []int{}
                                            _ = fi41
                                            _ = v42
                                        }
                                        if true {
                                            fmt.Println(v1, v1)
                                        }
                                    }
                                case "a":
                                    var v43 uint8 = 255
                                    _ = v43
                                }
                            }
                        case "foo bar":
                            for fi44 := 0; fi44 < 4; fi44++ {
                                fmt.Println(v1, v3, v27)
                                v27 = 1
                                _ = fi44
                            }
                        case "":
                            if (128 <= 127) {
                                v45 := GenS3{F0: 127}
                                if false {
                                    v2 += (v2 - (v2 - v2))
                                }
                                _ = v45
                            } else {
                                v46 := []int{2147483647, (fi17 * int(v4)), 32768, (fi17 % ((1) | 1))}
                                _ = v46
                            }
                        default:
                            switch ("\n\t" + "\n\t") {
                            case "hello":
                                if true {
                                    v47 := GenS2{F0: 779}
                                    v27 -= (v27 - v27)
                                    var v48 int8 = v27
                                    _ = v47
                                    _ = v48
                                } else {
                                    if (v27 >= v27) {
                                        v27 = (-(v27))
                                    }
                                }
                            default:
                                v4 = v4
                            }
                        }
                        _ = v29
                    }
                    if (true && (42 <= 18446744073709551615)) {
                        v1 = v1
                    } else {
                        fmt.Println(v9, fi17, v9)
                        if false {
                            fmt.Println(v27, v27)
                        }
                    }
                }
                _ = v27
            }
        } else {
            for fi49 := 0; fi49 < 15; fi49++ {
                v4 += (v4 - (v4 & (v4 >> ((53) & 63))))
                var v50 uint = 65535
                _ = fi49
                _ = v50
            }
        }
        if (false || true) {
            v51 := GenS1{F0: 42, F1: v2}
            _ = v51
        }
        _ = fi17
    }
    if (v3 > v3) {
        for fi52 := 0; fi52 < 12; fi52++ {
            switch "foo bar" {
            case "hello":
                for fi53 := 0; fi53 < 17; fi53++ {
                    v54 := GenS3{F0: 1}
                    if (24 <= 160) {
                        v55 := GenS4{F0: -82}
                        v4 = (v4 + v4)
                        _ = v55
                    }
                    var v56 uint32 = v4
                    v1 -= (v1 % ((v1) | 1))
                    _ = fi53
                    _ = v54
                    _ = v56
                }
            case "":
                v57 := GenS2{F0: 2147483647}
                _ = v57
            case "a":
                var v58 int = fi52
                _ = v58
            }
            var v59 float32 = (v9 + v9)
            _ = fi52
            _ = v59
        }
        switch -412 {
        case 9223372036854775807:
            if (!false) {
                v60 := GenS5{F0: 12, F1: 127}
                v61 := GenS6{F0: "a"}
                switch 9223372036854775807 {
                case 0:
                    var v62 bool = (true && true)
                    _ = v62
                case 155:
                    v4 = v4
                case -26:
                    v63 := GenS4{F0: -91}
                    _ = v63
                }
                _ = v60
                _ = v61
            } else {
                if (!(true || (true && false))) {
                    v64 := GenS4{F0: -128}
                    v65 := GenS7{F0: -26, F1: 42}
                    _ = v64
                    _ = v65
                }
                fmt.Println(v2, v3)
                for fi66 := 0; fi66 < 8; fi66++ {
                    v1 -= 42
                    _ = fi66
                }
            }
        }
    } else {
        for fi67 := 0; fi67 < 15; fi67++ {
            for fi68 := 0; fi68 < 20; fi68++ {
                v1 = 1
                _ = fi68
            }
            for fi69 := 0; fi69 < 18; fi69++ {
                if (!true) {
                    v70 := GenS8{F0: 42}
                    fmt.Println(v9)
                    if true {
                        v71 := GenS8{F0: 1}
                        if true {
                            var v72 uint16 = (v1 - v1)
                            switch (-(fi69)) {
                            case -9223372036854775808:
                                v73 := []int{fi69}
                                _ = v73
                            case -143:
                                if ((!(fi69 >= (fi67 | 0))) || ((((-2.25 / (0.1 * -2.25)) * ((float64(float32(0.1)) - 1.5) / 0.1)) + 1.5) == (3.14159 + 0.1))) {
                                    if false {
                                        var v74 uint32 = 659
                                        _ = v74
                                    } else {
                                        var v75 int = fi69
                                        switch (fi67 & fi67) {
                                        case 892:
                                            var v76 float32 = 0.0
                                            _ = v76
                                        case 1:
                                            v77 := GenS6{F0: "a"}
                                            _ = v77
                                        case 127:
                                            fmt.Println(fi69, v1, v72)
                                        default:
                                            var v78 float64 = 0.1
                                            _ = v78
                                        }
                                        _ = v75
                                    }
                                    if (true && (!(42 != -144))) {
                                        v79 := []int{fi69}
                                        if len(v79) >= 2 {
                                            v79 = v79[1:len(v79)]
                                        }
                                        _ = v79
                                    }
                                    if (v9 >= -2.25) {
                                        v80 := GenS6{F0: v3}
                                        var v81 float64 = (1e10 * 1.5)
                                        if false {
                                            fmt.Println(v4)
                                            v82 := GenS4{F0: 100}
                                            var v83 uint16 = (v72 ^ (v72 % (((v1 | 173)) | 1)))
                                            _ = v82
                                            _ = v83
                                        } else {
                                            if (true && false) {
                                                var v84 string = "hello"
                                                v85 := []int{-9223372036854775808, -111, 2147483647, fi67}
                                                _ = v84
                                                _ = v85
                                            }
                                            var v86 int = (fi69 % ((4294967295) | 1))
                                            var v87 int32 = 902
                                            _ = v86
                                            _ = v87
                                        }
                                        _ = v80
                                        _ = v81
                                    } else {
                                        switch (fi69 | fi67) {
                                        case 4294967295:
                                            v88 := []int{-632, fi69, fi67}
                                            _ = v88
                                        case 32768:
                                            var v89 float32 = 1e10
                                            _ = v89
                                        }
                                        var v90 int16 = v2
                                        switch fi67 {
                                        case 372:
                                            v2 += (v90 + (v2 * v2))
                                        case 306:
                                            v91 := []int{-9223372036854775808, fi67, 710}
                                            _ = v91
                                        case 32768:
                                            v92 := GenS9{F0: ("a" + v3)}
                                            _ = v92
                                        }
                                        _ = v90
                                    }
                                } else {
                                    var v93 uint32 = 2147483647
                                    v3 = ((("\n\t" + (v3 + v3)) + ("foo bar" + (v3 + v3))) + "foo bar")
                                    v94 := []int{fi69, 1}
                                    fmt.Println(v3)
                                    _ = v93
                                    _ = v94
                                }
                            }
                            _ = v72
                        } else {
                            v95 := GenS1{F0: 42, F1: (v2 + v2)}
                            fmt.Println(v2, v4, v9)
                            if true {
                                v96 := []int{(fi67 & 127), fi69, -1, (fi69 + (fi69 / ((0) | 1)))}
                                v97 := GenS10{F0: 9223372036854775807, F1: -1, F2: 457}
                                v98 := GenS11{F0: 127, F1: -21}
                                v4 += 65535
                                var v99 int = (fi69 + 0)
                                _ = v96
                                _ = v97
                                _ = v98
                                _ = v99
                            }
                            _ = v95
                        }
                        _ = v71
                    }
                    _ = v70
                } else {
                    var v100 int = fi69
                    v101 := GenS5{F0: 42, F1: 100}
                    v1 += (v1 | v1)
                    _ = v100
                    _ = v101
                }
                fmt.Println(v3, v2, v1)
                var v102 uint32 = (^(v4))
                var v103 int32 = 1
                _ = fi69
                _ = v102
                _ = v103
            }
            _ = fi67
        }
        fmt.Println(v3)
    }
    var v104 uint64 = 4294967295
    var v105 float32 = v9
    if (v4 < (^(v4))) {
        for fi106 := 0; fi106 < 4; fi106++ {
            var v107 int = fi106
            switch (v3 + v3) {
            case "":
                v108 := GenS3{F0: 0}
                _ = v108
            case "\n\t":
                v109 := []int{(v107 * -927), fi106}
                _ = v109
            case "hello":
                v110 := []int{(fi106 + v107), int(v4)}
                _ = v110
            }
            _ = fi106
            _ = v107
        }
        v111 := []int{322, 0}
        v3 = ""
        _ = v111
    }
    v112 := []int{}
    for fi113 := 0; fi113 < 2; fi113++ {
        v114 := GenS12{F0: uint8(v1)}
        for fi115 := 0; fi115 < 1; fi115++ {
            var v116 int8 = -1
            _ = len(v112)
            v3 = v3
            _ = fi115
            _ = v116
        }
        _ = fi113
        _ = v114
    }
    if true {
        v105 -= v9
    }
    switch v3 {
    case "foo bar":
        var v117 uint32 = 2
        _ = v117
    case "hello":
        switch 599 {
        case -2147483648:
            v9 -= ((v105 / -2.25) * (1.5 - (((0.1 + 3.14159) / v105) * -2.25)))
        }
    default:
        switch v3 {
        case "":
            var v118 float64 = 3.14159
            _ = v118
        }
    }
    fmt.Println(v9, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("slicelen v112:", len(v112), v112)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v104
    _ = v105
    _ = v112
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
