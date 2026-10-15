package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int16; F1 int64; F2 uint }
type GenS9 struct { F0 int64; F1 uint64 }
type GenS10 struct { F0 int64; F1 uint64; F2 int64 }
func genFn0(p0 bool, p1 int64, p2 uint32) uint {
    var v7 int8 = 42
    if p0 {
        v8 := []int{1, 2147483647}
        v9 := GenS1{F0: v7}
        _ = v8
        _ = v9
    } else {
        v10 := []int{128}
        _ = v10
    }
    if (p2 != p2) {
        v11 := []int{0, -299, -1, -730}
        var v12 int64 = (p1 + p1)
        _ = v11
        _ = v12
    } else {
        fmt.Println(p1, p1)
        switch 855 {
        case 9223372036854775807:
            p2 += p2
        case 42:
            fmt.Println(v7)
        case 836:
            fmt.Println(p1, p2, p0)
        default:
            switch -2147483648 {
            case 127:
                v13 := []int{128, -911, -2147483648}
                _ = v13
            case -9223372036854775808:
                v14 := []int{}
                _ = v14
            case 34:
                var v15 bool = p0
                _ = v15
            default:
                p1 = -884
            }
        }
    }
    switch ("" + (((("a" + ("hello" + "a")) + ("" + ((("" + ("" + "foo bar")) + ("\n\t" + "foo bar")) + "a"))) + "a") + ("" + "\n\t"))) {
    case "a":
        v16 := []int{273, 9223372036854775807}
        _ = v16
    case "\n\t":
        switch -387 {
        case 2147483647:
            p1 -= p1
        case 363:
            var v17 uint32 = p2
            _ = v17
        default:
            if (p2 > 4294967295) {
                var v18 uint32 = p2
                _ = v18
            } else {
                v7 = (-(v7))
                var v19 int32 = 42
                var v20 int16 = -955
                _ = v19
                _ = v20
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float64 = float64(0.0)
    var v3 uint = 255
    var v4 bool = false
    var v5 uint = v3
    var v6 int64 = 325
    v21 := []int{0, 0, 2147483647}
    switch (("\n\t" + "a") + "") {
    case "\n\t":
        if true {
            var v22 int = 2147483647
            _ = v22
        }
    case "foo bar":
        if (!(!(v2 > ((v2 + float64(1e10)) + v2)))) {
            fmt.Println(v5, v5)
            v23 := GenS2{F0: v6}
            var v24 bool = (0.1 == ((1e10 + 1e10) + float32(((v2 - 0.0) * v2))))
            _ = v23
            _ = v24
        }
    case "":
        if (v6 >= v6) {
            for fi25 := 0; fi25 < 4; fi25++ {
                if (241 >= 1) {
                    v26 := GenS1{F0: -38}
                    switch ("hello" + "\n\t") {
                    case "a":
                        v5 = 0
                    case "foo bar":
                        v27 := GenS2{F0: 1}
                        _ = v27
                    case "":
                        for fi28 := 0; fi28 < 8; fi28++ {
                            v5 += (^(v3))
                            _ = fi28
                        }
                    default:
                        switch 4294967295 {
                        case 32767:
                            if ((fi25 >> ((v5) & 63)) != (fi25 >> ((v3) & 63))) {
                                if ((!(v3 != 42)) || (!v4)) {
                                    v2 = (float64(-2.25) + (0.1 / v2))
                                    if v4 {
                                        v3 -= v5
                                    }
                                } else {
                                    var v29 int = fi25
                                    switch (("foo bar" + "\n\t") + ("\n\t" + ("" + ""))) {
                                    case "hello":
                                        var v30 int8 = 0
                                        _ = v30
                                    case "":
                                        if len(v21) >= 2 {
                                            v21 = v21[1:len(v21)]
                                        }
                                    case "\n\t":
                                        fmt.Println(v6, v5)
                                    default:
                                        var v31 float64 = 0.0
                                        _ = v31
                                    }
                                    if false {
                                        if (!v4) {
                                            switch (((("" + "a") + ((("" + "hello") + ("a" + "\n\t")) + ("\n\t" + ""))) + "\n\t") + ("hello" + "foo bar")) {
                                            case "hello":
                                                if len(v21) > 0 {
                                                    v21[0] = (v29 | fi25)
                                                }
                                            default:
                                                var v32 uint16 = 296
                                                _ = v32
                                            }
                                            v33 := GenS3{F0: 1, F1: 1, F2: 2147483647}
                                            if len(v21) > 0 {
                                                v21[0] = (v29 & fi25)
                                            }
                                            _ = v33
                                        } else {
                                            var v34 uint = v5
                                            var v35 float64 = 0.1
                                            _ = v34
                                            _ = v35
                                        }
                                        if v4 {
                                            v36 := GenS4{F0: v3, F1: (^(v29))}
                                            var v37 int16 = 32767
                                            v21 = append(v21, (fi25 >> ((v3) & 63)))
                                            _ = v36
                                            _ = v37
                                        } else {
                                            _ = len(v21)
                                        }
                                    } else {
                                        for fi38 := 0; fi38 < 1; fi38++ {
                                            fmt.Println(v4, v3)
                                            fmt.Println(v6, v3, v5)
                                            _ = fi38
                                        }
                                    }
                                    _ = v29
                                }
                            }
                        default:
                            var v39 uint32 = 2
                            _ = v39
                        }
                    }
                    var v40 uint = v3
                    _ = v26
                    _ = v40
                } else {
                    v6 += int64(v5)
                }
                var v41 string = ("\n\t" + "foo bar")
                fmt.Println(v4, v1, v1)
                _ = fi25
                _ = v41
            }
        }
    default:
        var v42 uint8 = 42
        _ = v42
    }
    var v43 int64 = v6
    v44 := GenS2{F0: v6}
    var v45 int64 = v43
    v3 += uint(v1)
    fmt.Println(v1, v1, v3)
    var v46 uint32 = 145
    if ((!false) && v4) {
        var v47 uint64 = 18446744073709551615
        switch "hello" {
        case "hello":
            v3 = (^(v5))
        default:
            if v4 {
                var v48 string = ""
                var v49 uint8 = 127
                _ = v48
                _ = v49
            } else {
                for fi50 := 0; fi50 < 7; fi50++ {
                    var v51 int8 = 100
                    fmt.Println(v2)
                    v52 := GenS2{F0: int64(v51)}
                    _ = fi50
                    _ = v51
                    _ = v52
                }
            }
        }
        v21 = append(v21, 1)
        _ = v47
    } else {
        v45 = 327
        var v53 int = 4294967295
        switch v53 {
        case 127:
            var v54 uint64 = 18446744073709551615
            _ = v54
        default:
            v55 := GenS2{F0: (v6 % ((-9223372036854775808) | 1))}
            _ = v55
        }
        _ = v53
    }
    switch ((("\n\t" + "a") + "a") + "hello") {
    case "\n\t":
        for fi56 := 0; fi56 < 9; fi56++ {
            switch fi56 {
            case 32768:
                v6 -= 9223372036854775807
            case 127:
                v6 = (v45 + 9223372036854775807)
            }
            _ = fi56
        }
    default:
        switch "hello" {
        case "a":
            fmt.Println(v3, v2, v46)
        case "hello":
            switch -695 {
            case 4294967295:
                for fi57 := 0; fi57 < 1; fi57++ {
                    var v58 int8 = 30
                    _ = fi57
                    _ = v58
                }
            case -1:
                v21 = append(v21, 32768)
            case -264:
                var v59 int = 655
                _ = v59
            }
        case "foo bar":
            fmt.Println(v4, v2)
        }
    }
    var v60 int64 = 517
    switch ("\n\t" + ("\n\t" + "a")) {
    case "foo bar":
        if (!v4) {
            _ = len(v21)
            for fi61 := 0; fi61 < 15; fi61++ {
                if v4 {
                    var v62 bool = (false || (-128 != -1))
                    _ = v62
                } else {
                    if len(v21) > 0 {
                        v21[0] = fi61
                    }
                    var v63 int = 42
                    _ = v63
                }
                v64 := GenS5{F0: v46, F1: int32(v46)}
                _ = fi61
                _ = v64
            }
        }
    case "":
        for fi65 := 0; fi65 < 4; fi65++ {
            v21 = append(v21, 0)
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
            _ = fi65
        }
    default:
        v5 = 2
    }
    v45 = v43
    if true {
        for fi66 := 0; fi66 < 5; fi66++ {
            v6 = 42
            switch "foo bar" {
            case "\n\t":
                for fi67 := 0; fi67 < 4; fi67++ {
                    v68 := GenS2{F0: v6}
                    _ = fi67
                    _ = v68
                }
            case "hello":
                if (v4 && (276 <= 472)) {
                    v69 := GenS5{F0: (v46 >> (((v3 >> (((v3 & 1)) & 63))) & 63)), F1: -2147483648}
                    v1 -= (v1 ^ (v1 % ((v1) | 1)))
                    _ = v69
                } else {
                    if ((true && ("hello" >= "foo bar")) && v4) {
                        switch "a" {
                        case "":
                            fmt.Println(v2, v1, v60)
                        case "a":
                            v45 += (^(v60))
                        default:
                            for fi70 := 0; fi70 < 11; fi70++ {
                                _ = len(v21)
                                v43 -= (^(v6))
                                _ = fi70
                            }
                        }
                    }
                    if true {
                        if (!v4) {
                            fmt.Println(v2, v6)
                            v71 := GenS5{F0: v46, F1: -635}
                            _ = v71
                        }
                        v3 = uint(v46)
                    } else {
                        if true {
                            v72 := GenS5{F0: v46, F1: -193}
                            for fi73 := 0; fi73 < 2; fi73++ {
                                var v74 float32 = 1e10
                                if (-38 >= -1) {
                                    switch fi73 {
                                    case -1:
                                        if true {
                                            v2 = v2
                                            if v4 {
                                                switch ("foo bar" + "foo bar") {
                                                case "foo bar":
                                                    v43 -= (v6 * (v45 % (((v6 * v60)) | 1)))
                                                case "hello":
                                                    if (false && (((0.0 + -2.25) / v74) == v74)) {
                                                        if (v2 >= v2) {
                                                            var v75 float32 = (v74 / v74)
                                                            v5 += v3
                                                            _ = v75
                                                        }
                                                        v76 := GenS6{F0: (^(v1))}
                                                        v43 = v60
                                                        _ = v76
                                                    }
                                                case "\n\t":
                                                    v5 = 255
                                                default:
                                                    var v77 uint64 = (v1 & v1)
                                                    _ = v77
                                                }
                                                var v78 int16 = 1
                                                _ = v78
                                            }
                                            switch fi66 {
                                            case 4294967295:
                                                switch 9223372036854775807 {
                                                case 2:
                                                    switch "\n\t" {
                                                    case "hello":
                                                        fmt.Println(v4, fi73)
                                                        fmt.Println(v2, v60)
                                                    case "\n\t":
                                                        v2 = float64(v74)
                                                    case "":
                                                        fmt.Println(v60)
                                                    }
                                                }
                                            case -2147483648:
                                                v4 = v4
                                            }
                                        } else {
                                            _ = len(v21)
                                            if (v5 >= (v3 - genFn0(false, v60, (v46 + v46)))) {
                                                var v79 uint64 = (v1 - (^(v1)))
                                                _ = v79
                                            } else {
                                                fmt.Println(v60, v1, v2)
                                                v3 = v5
                                            }
                                            var v80 float64 = 0.0
                                            _ = v80
                                        }
                                    case 474:
                                        if len(v21) >= 2 {
                                            v21 = v21[1:len(v21)]
                                        }
                                    default:
                                        var v81 int8 = -110
                                        v46 += 42
                                        _ = v81
                                    }
                                }
                                _ = fi73
                                _ = v74
                            }
                            _ = v72
                        } else {
                            _ = len(v21)
                        }
                        if v4 {
                            v82 := GenS7{F0: fi66}
                            if ((18446744073709551615 == v1) || (42 <= -32768)) {
                                for fi83 := 0; fi83 < 16; fi83++ {
                                    v46 = 2147483648
                                    if ((v45 - (-(v6))) == v43) {
                                        var v84 uint64 = 0
                                        var v85 int16 = 42
                                        _ = v84
                                        _ = v85
                                    } else {
                                        var v86 int = fi83
                                        var v87 uint8 = 42
                                        v5 += (v3 * uint(v87))
                                        _ = v86
                                        _ = v87
                                    }
                                    var v88 uint = v3
                                    _ = fi83
                                    _ = v88
                                }
                                switch 2147483647 {
                                case -464:
                                    v89 := GenS8{F0: -466, F1: -1, F2: v3}
                                    _ = v89
                                case 128:
                                    fmt.Println(v45, v2, v45)
                                case 2147483647:
                                    if len(v21) > 0 {
                                        v21[0] = fi66
                                    }
                                }
                                var v90 string = "foo bar"
                                _ = v90
                            }
                            if false {
                                var v91 uint64 = v1
                                fmt.Println(v3, v3)
                                _ = v91
                            }
                            _ = v82
                        } else {
                            for fi92 := 0; fi92 < 4; fi92++ {
                                if (false && (-2.25 > (1.5 * (((0.1 * 0.1) + ((float32((float64(((3.14159 + -2.25) / 3.14159)) + float64(1.5))) / (float32(v2) / 1.5)) * 0.1)) * (1.5 + (((3.14159 + 1e10) / 1.0) * (0.0 + (3.14159 / float32(float64(1.5)))))))))) {
                                    v93 := GenS5{F0: v46, F1: -1}
                                    var v94 int = (fi66 / ((fi92) | 1))
                                    if true {
                                        v95 := GenS7{F0: -701}
                                        v96 := GenS9{F0: v45, F1: v1}
                                        if (42 != 32767) {
                                            var v97 bool = v4
                                            var v98 int = v94
                                            _ = v97
                                            _ = v98
                                        } else {
                                            v99 := GenS9{F0: v43, F1: 42}
                                            switch (^(fi92)) {
                                            case 4294967295:
                                                v100 := GenS10{F0: (v45 / ((v45) | 1)), F1: v1, F2: v6}
                                                _ = v100
                                            case 110:
                                                v43 = (v43 + v43)
                                            case -2147483648:
                                                v101 := GenS8{F0: 1000, F1: (v60 * v6), F2: v5}
                                                _ = v101
                                            }
                                            var v102 float64 = v2
                                            fmt.Println(v6)
                                            _ = v102
                                            _ = v99
                                        }
                                        var v103 uint = uint(v1)
                                        _ = v103
                                        _ = v95
                                        _ = v96
                                    }
                                    v104 := GenS6{F0: 2}
                                    if v4 {
                                        _ = len(v21)
                                        if (-1 == -128) {
                                            v60 = -110
                                        } else {
                                            var v105 string = ""
                                            _ = v105
                                        }
                                    }
                                    _ = v104
                                    _ = v93
                                    _ = v94
                                }
                                v43 -= int64(v3)
                                var v106 uint16 = 255
                                _ = fi92
                                _ = v106
                            }
                            var v107 int32 = 1
                            _ = v107
                        }
                    }
                    var v108 int32 = 1
                    _ = v108
                }
            case "foo bar":
                if v4 {
                    var v109 uint64 = 225
                    _ = len(v21)
                    _ = v109
                }
            }
            _ = fi66
        }
        v60 += (v60 + (v60 * v60))
        var v110 float64 = float64(-2.25)
        _ = v110
    }
    v43 -= (v43 + v6)
    if len(v21) > 0 {
        v21[0] = 2147483647
    }
    var v111 int16 = 796
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v111
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v60
}
