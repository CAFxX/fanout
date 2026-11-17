package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 string; F2 int8 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 int8; F2 int8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16; F1 uint16; F2 int16 }
type GenS9 struct { F0 uint8; F1 int }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 string; F1 uint64 }
func genFn0(p0 int16, p1 uint16, p2 int16) uint32 {
    p1 = p1
    fmt.Println(p1, p1)
    var v7 int8 = -1
    var v8 int8 = (^(v7))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 915
}
func genFn1() int64 {
    var v9 string = "foo bar"
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 int = fi10
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 115, F1: 642}
    fmt.Println(v9)
    if true {
        if false {
            if false {
                fmt.Println(v9)
                for fi13 := 0; fi13 < 9; fi13++ {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        fmt.Println(v9, v9, fi14)
                        _ = fi14
                    }
                    _ = fi13
                }
            }
            v9 = (v9 + v9)
            switch "hello" {
            case "":
                v9 = "foo bar"
            case "foo bar":
                switch v9 {
                case "a":
                    var v15 int64 = 1
                    _ = v15
                default:
                    if (false || true) {
                        fmt.Println(v9)
                        switch ((v9 + v9) + ("hello" + ("\n\t" + ("hello" + (v9 + v9))))) {
                        case "hello":
                            var v16 float64 = 0.0
                            _ = v16
                        case "a":
                            var v17 int = 0
                            _ = v17
                        }
                        v9 = ((v9 + v9) + v9)
                    } else {
                        switch 42 {
                        case 1:
                            var v18 uint64 = 0
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 16; fi19++ {
                            v20 := []int{(fi19 / ((-721) | 1))}
                            _ = fi19
                            _ = v20
                        }
                    }
                }
            case "hello":
                fmt.Println(v9)
            }
        }
        v9 = ("a" + ("foo bar" + ""))
    } else {
        v9 = v9
    }
    _ = v12
    _ = v9
    return 1
}

func main() {
    var v1 float32 = -2.25
    var v2 uint64 = 482
    var v3 int = 9223372036854775807
    v4 := make(map[int]int)
    v5 := make(map[int]uint32)
    var v6 string = ""
    v1 = v1
    switch v6 {
    case "":
        var v21 bool = true
        _ = v21
    }
    switch v6 {
    case "hello":
        for fi22 := 0; fi22 < 11; fi22++ {
            if (2 < 2) {
                var v23 uint16 = 670
                _ = v23
            }
            _ = fi22
        }
    default:
        switch (v3 / ((v3) | 1)) {
        case 127:
            if false {
                fmt.Println(v3, v3)
                for fi24 := 0; fi24 < 14; fi24++ {
                    switch (-(v3)) {
                    case -9223372036854775808:
                        var v25 float32 = (v1 / (((v1 * v1) * (v1 * (((v1 / (0.0 - (v1 - v1))) * (((float32((3.14159 + (1.5 - 1.5))) - 0.0) * 1e10) * float32(0.1))) - v1))) - v1))
                        _ = v25
                    case 848:
                        var v26 uint8 = 205
                        _ = v26
                    default:
                        v1 = float32(-2.25)
                    }
                    v27 := GenS1{F0: uint8(v2), F1: -345}
                    _ = fi24
                    _ = v27
                }
            }
        default:
            if (!true) {
                if (genFn0(539, 514, -527) < 2) {
                    var v28 int8 = 79
                    for fi29 := 0; fi29 < 9; fi29++ {
                        if (0 > 22) {
                            fmt.Println(v1)
                            if true {
                                v30 := []int{}
                                if false {
                                    v2 = v2
                                    if (true || false) {
                                        for fi31 := 0; fi31 < 16; fi31++ {
                                            v28 += (^(v28))
                                            switch v6 {
                                            case "":
                                                v30 = append(v30, v3)
                                            case "hello":
                                                v32 := GenS2{F0: 42}
                                                _ = v32
                                            case "foo bar":
                                                v28 += v28
                                            }
                                            _ = fi31
                                        }
                                        switch v6 {
                                        case "foo bar":
                                            var v33 float64 = 0.0
                                            _ = v33
                                        case "hello":
                                            for fi34 := 0; fi34 < 8; fi34++ {
                                                var v35 uint = 255
                                                var v36 int64 = genFn1()
                                                _ = fi34
                                                _ = v35
                                                _ = v36
                                            }
                                        case "":
                                            var v37 uint32 = 4294967295
                                            _ = v37
                                        default:
                                            fmt.Println(fi29, v2)
                                        }
                                    }
                                } else {
                                    v2 += 666
                                    var v38 int32 = -279
                                    _ = v38
                                }
                                v28 -= v28
                                _ = v30
                            }
                            var v39 float64 = 1.5
                            _ = v39
                        } else {
                            if false {
                                _ = len(v5)
                            } else {
                                _ = len(v5)
                                if false {
                                    var v40 int = 42
                                    _ = v40
                                } else {
                                    for fi41 := 0; fi41 < 7; fi41++ {
                                        if (true || false) {
                                            if true {
                                                switch v6 {
                                                case "hello":
                                                    v42 := []int{fi41, 128, fi29, -2147483648}
                                                    _ = v42
                                                default:
                                                    if true {
                                                        if (true || true) {
                                                            if (v28 == 100) {
                                                                var v43 int64 = genFn1()
                                                                fmt.Println(fi29, v2)
                                                                v1 += v1
                                                                _ = v43
                                                            }
                                                            switch "a" {
                                                            case "":
                                                                if (!false) {
                                                                    v2 -= (v2 % ((4294967295) | 1))
                                                                } else {
                                                                    if false {
                                                                        switch fi41 {
                                                                        case 32768:
                                                                            if true {
                                                                                var v44 float64 = 3.14159
                                                                                v45 := GenS1{F0: 97, F1: 50}
                                                                                _ = v44
                                                                                _ = v45
                                                                            }
                                                                        default:
                                                                            switch v6 {
                                                                            case "":
                                                                                switch fi29 {
                                                                                case -926:
                                                                                    switch (fi41 ^ v3) {
                                                                                    case -235:
                                                                                        if (!(127 == 255)) {
                                                                                            switch v3 {
                                                                                            case 634:
                                                                                                if (v3 >= (^(fi41))) {
                                                                                                    v46 := []int{(^(v3)), 0, 32768}
                                                                                                    v47 := GenS2{F0: 223}
                                                                                                    _ = v46
                                                                                                    _ = v47
                                                                                                }
                                                                                            }
                                                                                            v48 := GenS2{F0: 153}
                                                                                            _ = v48
                                                                                        } else {
                                                                                            v2 += v2
                                                                                            if true {
                                                                                                v49 := GenS3{F0: 32767, F1: ("hello" + v6), F2: v28}
                                                                                                fmt.Println(v6, v1)
                                                                                                if (true && (!((v2 ^ v2) >= (v2 * 18446744073709551615)))) {
                                                                                                    v50 := []int{128, fi41, fi29, 381}
                                                                                                    v1 += v1
                                                                                                    v51 := GenS4{F0: 65535, F1: -1}
                                                                                                    _ = v50
                                                                                                    _ = v51
                                                                                                }
                                                                                                v4[42] = 127
                                                                                                _ = v49
                                                                                            }
                                                                                            if ((0.0 / 1.0) != (v1 / v1)) {
                                                                                                var v52 uint32 = 0
                                                                                                var v53 int64 = 9223372036854775807
                                                                                                _ = v52
                                                                                                _ = v53
                                                                                            } else {
                                                                                                fmt.Println(v6, v1)
                                                                                                v1 -= ((v1 / v1) - 3.14159)
                                                                                                if (0 >= 103) {
                                                                                                    fmt.Println(fi29)
                                                                                                    fmt.Println(v1, v6)
                                                                                                    var v54 int8 = (v28 / (((v28 >> ((255) & 63))) | 1))
                                                                                                    fmt.Println(fi29, fi41)
                                                                                                    v1 = (((0.1 * v1) - (((v1 * v1) + v1) + ((0.0 + v1) - v1))) / v1)
                                                                                                    _ = v54
                                                                                                } else {
                                                                                                    if (false || (42 <= -2147483648)) {
                                                                                                        v55 := GenS5{F0: (v2 | 42)}
                                                                                                        _ = v55
                                                                                                    }
                                                                                                    fmt.Println(v2, v6, v1)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case 128:
                                                                                        fmt.Println(v3, fi41, v3)
                                                                                    case 840:
                                                                                        var v56 int32 = int32(v28)
                                                                                        _ = v56
                                                                                    }
                                                                                case 127:
                                                                                    v57 := []int{v3, (v3 / (((-(fi29))) | 1)), 9223372036854775807}
                                                                                    _ = v57
                                                                                }
                                                                            case "\n\t":
                                                                                var v58 int32 = 42
                                                                                _ = v58
                                                                            }
                                                                            fmt.Println(fi41, v28)
                                                                        }
                                                                        var v59 int32 = -2147483648
                                                                        switch (fi41 & v3) {
                                                                        case 871:
                                                                            if (!(!(true || (true && true)))) {
                                                                                switch -729 {
                                                                                case -344:
                                                                                    v3 = (v3 | v3)
                                                                                }
                                                                            }
                                                                        case 9223372036854775807:
                                                                            var v60 uint8 = 255
                                                                            _ = v60
                                                                        case 0:
                                                                            var v61 bool = true
                                                                            v61 = true
                                                                            _ = v61
                                                                        }
                                                                        _ = v59
                                                                    }
                                                                }
                                                            case "a":
                                                                fmt.Println(fi41, v28)
                                                            case "hello":
                                                                var v62 uint64 = v2
                                                                _ = v62
                                                            default:
                                                                v2 += v2
                                                            }
                                                            if false {
                                                                var v63 string = "hello"
                                                                v1 = v1
                                                                v64 := GenS5{F0: 4294967295}
                                                                _ = v63
                                                                _ = v64
                                                            } else {
                                                                v65 := GenS1{F0: 0, F1: 1}
                                                                v2 += (^(v2))
                                                                v28 += (-(v28))
                                                                _ = v65
                                                            }
                                                        } else {
                                                            v6 = (("foo bar" + ("foo bar" + "hello")) + v6)
                                                            fmt.Println(v6, v28, v28)
                                                            fmt.Println(fi29, v6)
                                                        }
                                                        v1 -= (v1 / (((v1 - float32(0.1)) + v1) - v1))
                                                    }
                                                    fmt.Println(fi41, v6)
                                                    v66 := GenS6{F0: false, F1: (v28 >> ((255) & 63)), F2: 42}
                                                    _ = v66
                                                }
                                                v67 := GenS4{F0: 2147483648, F1: 9223372036854775807}
                                                _ = v67
                                            } else {
                                                v1 -= v1
                                                v68 := GenS2{F0: 4294967295}
                                                _ = v68
                                            }
                                            v69 := GenS7{F0: 1}
                                            var v70 float64 = 3.14159
                                            _ = v69
                                            _ = v70
                                        } else {
                                            fmt.Println(v3)
                                        }
                                        var v71 uint32 = 1
                                        var v72 uint64 = 0
                                        _ = fi41
                                        _ = v71
                                        _ = v72
                                    }
                                    fmt.Println(v2)
                                }
                            }
                            var v73 uint32 = 2
                            if false {
                                var v74 uint8 = 209
                                _ = v74
                            }
                            _ = v73
                        }
                        var v75 uint32 = 4294967295
                        fmt.Println(v75, v6, v3)
                        _ = fi29
                        _ = v75
                    }
                    v6 = v6
                    _ = v28
                }
                var v76 string = (((v6 + v6) + v6) + v6)
                _ = v76
            } else {
                v3 += v3
                var v77 uint = 874
                v77 -= 881
                _ = v77
            }
        }
    }
    delete(v5, v3)
    if false {
        v1 += (float32(1e10) / 1e10)
        if false {
            for fi78 := 0; fi78 < 12; fi78++ {
                switch (-(fi78)) {
                case 32768:
                    var v79 int = fi78
                    _ = v79
                }
                _ = fi78
            }
            v6 = ("a" + (((((v6 + v6) + "hello") + "hello") + v6) + v6))
        }
        v80 := []int{-581}
        _ = v80
    }
    var v81 uint16 = 32767
    v82 := GenS8{F0: 1, F1: (v81 & 1), F2: 42}
    v83 := GenS4{F0: 2147483648, F1: 9223372036854775807}
    if (v6 < "a") {
        v1 -= v1
    } else {
        if false {
            var v84 int32 = -702
            if (v84 > (v84 + v84)) {
                _ = len(v4)
                if (-1 < 127) {
                    var v85 uint8 = 2
                    fmt.Println(v1, v2)
                    var v86 uint64 = (v2 | (^(v2)))
                    _ = v85
                    _ = v86
                }
            } else {
                v87 := GenS9{F0: 127, F1: (v3 / (((v3 * v3)) | 1))}
                _ = v87
            }
            v6 = ((("foo bar" + v6) + (v6 + v6)) + v6)
            _ = v84
        } else {
            if false {
                var v88 float32 = ((v1 * (v1 * (v1 / v1))) + float32(3.14159))
                for fi89 := 0; fi89 < 7; fi89++ {
                    if (1 == v81) {
                        var v90 uint = 65535
                        switch fi89 {
                        case -278:
                            v91 := []int{(v3 / (((-(fi89))) | 1)), v3, v3, int(v81)}
                            _ = v91
                        case 0:
                            fmt.Println(v6, v88)
                        }
                        _ = v90
                    } else {
                        v2 += (^(v2))
                        v6 = v6
                        v1 += 1e10
                    }
                    for fi92 := 0; fi92 < 11; fi92++ {
                        v93 := []int{fi89}
                        delete(v4, (v3 / ((9223372036854775807) | 1)))
                        delete(v5, v3)
                        _ = fi92
                        _ = v93
                    }
                    v1 = 0.1
                    _ = fi89
                }
                for fi94 := 0; fi94 < 12; fi94++ {
                    var v95 int8 = 42
                    _ = fi94
                    _ = v95
                }
                _ = v88
            }
            v81 -= v81
            for fi96 := 0; fi96 < 19; fi96++ {
                if (!(!((genFn1() <= 9223372036854775807) && (65535 <= 2)))) {
                    fmt.Println(v1)
                } else {
                    switch fi96 {
                    case -378:
                        for fi97 := 0; fi97 < 20; fi97++ {
                            var v98 int64 = -311
                            v99 := GenS10{F0: v98}
                            _ = fi97
                            _ = v98
                            _ = v99
                        }
                    }
                    fmt.Println(v81, v6)
                }
                _ = fi96
            }
        }
        v100 := GenS1{F0: 255, F1: 1000}
        switch v3 {
        case 1:
            v101 := v5[(v3 | v3)]
            _ = v101
        case -1:
            v102 := []int{v3, (v3 << ((1) & 63)), v3}
            _ = v102
        case 128:
            v3 -= 9223372036854775807
        }
        _ = v100
    }
    switch v6 {
    case "hello":
        _ = len(v5)
    default:
        for fi103 := 0; fi103 < 8; fi103++ {
            var v104 int = v3
            switch fi103 {
            case 1:
                switch 615 {
                case 127:
                    if (42 > 203) {
                        v105 := []int{(v104 ^ fi103), v3, 2147483647}
                        fmt.Println(v104, v3)
                        if (true && true) {
                            var v106 float32 = (1e10 - (((v1 - 0.0) + float32(((1e10 - 3.14159) + 0.1))) / 1.5))
                            var v107 int8 = 0
                            _ = v106
                            _ = v107
                        } else {
                            var v108 uint16 = (v81 << ((65535) & 63))
                            if (true && true) {
                                var v109 int16 = int16(fi103)
                                v110 := GenS10{F0: 1}
                                v111 := GenS7{F0: 1}
                                _ = v109
                                _ = v110
                                _ = v111
                            }
                            var v112 int16 = -1
                            _ = v108
                            _ = v112
                        }
                        _ = v105
                    }
                case 828:
                    if false {
                        var v113 bool = (!true)
                        _ = v113
                    }
                case 431:
                    for fi114 := 0; fi114 < 6; fi114++ {
                        var v115 uint16 = (v81 * v81)
                        _ = fi114
                        _ = v115
                    }
                default:
                    v116 := GenS11{F0: v6, F1: 0}
                    _ = v116
                }
            case 0:
                v117 := []int{(fi103 >> ((255) & 63)), 32768}
                _ = v117
            case 128:
                v2 = (^(v2))
            default:
                v6 = "hello"
            }
            _ = fi103
            _ = v104
        }
    }
    var v118 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v118:", v118)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v82
    _ = v83
}
