package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint64 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 uint; F2 uint8 }
type GenS6 struct { F0 uint64; F1 string; F2 int }
type GenS7 struct { F0 int8; F1 int16; F2 int8 }
type GenS8 struct { F0 uint; F1 uint64 }
func genFn0() bool {
    v5 := GenS1{F0: -14, F1: -9223372036854775808, F2: 18446744073709551615}
    var v6 bool = (-9223372036854775808 < 0)
    fmt.Println(v6)
    var v7 int32 = 935
    _ = v5
    _ = v6
    _ = v7
    return (112 < 592)
}
func genFn1() uint8 {
    if (65535 != 32767) {
        var v8 uint = 442
        v8 -= v8
        v9 := GenS1{F0: 42, F1: 0, F2: 640}
        _ = v8
        _ = v9
    } else {
        var v10 int16 = -1
        if (0 < 2147483647) {
            v10 -= (v10 * (-(v10)))
        }
        v11 := GenS1{F0: (^(v10)), F1: 9223372036854775807, F2: 18446744073709551615}
        _ = v10
        _ = v11
    }
    var v12 int16 = -40
    _ = v12
    return 127
}

func main() {
    var v1 uint = 1
    v2 := []uint32{uint32(v1)}
    var v3 float64 = 1e10
    v4 := make(map[uint32]bool)
    switch "foo bar" {
    case "":
        v3 -= v3
    case "a":
        fmt.Println(v1, v3)
    case "\n\t":
        switch "\n\t" {
        case "\n\t":
            var v13 int16 = 32767
            _ = v13
        }
    }
    v1 += v1
    var v14 uint64 = 4294967295
    var v15 uint8 = 1
    if false {
        _ = len(v4)
        if (-581 <= -1) {
            v14 += (v14 & (v14 >> ((v1) & 63)))
            if false {
                var v16 uint16 = 225
                _ = v16
            }
            v1 = 0
        } else {
            if ("\n\t" == "\n\t") {
                var v17 float32 = 1.5
                _ = v17
            }
            v18 := v4[4294967295]
            _ = v18
        }
    }
    fmt.Println(v14, v1, v3)
    for fi19 := 0; fi19 < 1; fi19++ {
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi19
    }
    switch "a" {
    case "hello":
        v4[557] = true
    case "foo bar":
        var v20 int64 = -9223372036854775808
        _ = v20
    }
    v21 := GenS2{F0: -1, F1: (("\n\t" + "") + (("hello" + "foo bar") + ("a" + "a")))}
    for fi22 := 0; fi22 < 20; fi22++ {
        if true {
            if false {
                v3 = 1.5
            }
            for fi23 := 0; fi23 < 15; fi23++ {
                var v24 float64 = (1.5 * (((float64(-2.25) - 0.1) + v3) * 1e10))
                v3 = 1e10
                fmt.Println(v3, v24)
                _ = fi23
                _ = v24
            }
            fmt.Println(v14)
        } else {
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 int64 = -1
                switch fi22 {
                case 42:
                    v27 := GenS1{F0: 706, F1: v26, F2: (v14 | (v14 & v14))}
                    _ = v27
                case -175:
                    v1 -= v1
                    switch (fi25 << ((1) & 63)) {
                    case 2147483647:
                        var v28 int64 = (v26 | 9223372036854775807)
                        _ = v28
                    }
                default:
                    if false {
                        switch fi25 {
                        case 32767:
                            if false {
                                var v29 uint = (v1 | (v1 | (v1 << ((v1) & 63))))
                                _ = v29
                            }
                            switch (fi22 | 128) {
                            case 0:
                                switch ("" + "foo bar") {
                                case "":
                                    var v30 string = ("\n\t" + (("hello" + "a") + "a"))
                                    _ = v30
                                case "\n\t":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch fi25 {
                                        case -1:
                                            v31 := GenS2{F0: -1, F1: "a"}
                                            _ = v31
                                        case 42:
                                            v4[0] = false
                                        }
                                    case "hello":
                                        if true {
                                            if ((!((-(v26)) <= int64(fi22))) || true) {
                                                var v32 uint8 = (^(v15))
                                                v15 = (v32 / (((^(v32))) | 1))
                                                _ = v32
                                            }
                                            if (("" > "hello") || (float32(((v3 * -2.25) + (v3 / float64(1e10)))) >= -2.25)) {
                                                v15 -= 192
                                            } else {
                                                if (-32768 != -1) {
                                                    var v33 uint16 = 255
                                                    v33 -= (v33 - 0)
                                                    if len(v2) > 0 {
                                                        v2[0] = 2147483647
                                                    }
                                                    if false {
                                                        _ = len(v2)
                                                        if (false || (42 <= (^(v33)))) {
                                                            if true {
                                                                v26 += 1
                                                                switch "" {
                                                                case "a":
                                                                    v34 := GenS3{F0: v14}
                                                                    _ = v34
                                                                }
                                                            }
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        if true {
                                                            fmt.Println(v15, v26, v14)
                                                            v35 := GenS4{F0: (v15 + (v15 % ((220) | 1)))}
                                                            _ = v35
                                                        }
                                                    }
                                                    _ = v33
                                                } else {
                                                    if true {
                                                        if (960 > 996) {
                                                            var v36 float64 = ((((((0.1 - 1.5) / 0.1) * ((((v3 + (-2.25 / v3)) + ((float64(((1.5 * (1.5 + 1.5)) / float32(((1.5 / -2.25) + v3)))) * v3) - v3)) - v3) / (v3 + v3))) - 1.5) * (1e10 * v3)) - v3)
                                                            var v37 float32 = 1.5
                                                            _ = v36
                                                            _ = v37
                                                        } else {
                                                            if true {
                                                                if ((!true) || (v26 < v26)) {
                                                                    delete(v4, 2147483647)
                                                                } else {
                                                                    if false {
                                                                        var v38 int = (fi22 & fi22)
                                                                        _ = v38
                                                                    }
                                                                }
                                                            }
                                                            if true {
                                                                v15 -= uint8(v1)
                                                                v39 := GenS4{F0: 143}
                                                                _ = v39
                                                            } else {
                                                                var v40 float64 = -2.25
                                                                if false {
                                                                    if ((!(v1 <= 255)) && false) {
                                                                        if true {
                                                                            fmt.Println(fi25, fi25, v3)
                                                                            var v41 uint8 = v15
                                                                            v1 -= v1
                                                                            v42 := GenS3{F0: 211}
                                                                            _ = v41
                                                                            _ = v42
                                                                        }
                                                                        if (v15 < genFn1()) {
                                                                            v3 -= (1e10 + (v40 / v3))
                                                                        }
                                                                        v43 := v4[65535]
                                                                        _ = v43
                                                                    }
                                                                    v44 := GenS3{F0: v14}
                                                                    var v45 float32 = 3.14159
                                                                    v46 := GenS1{F0: 1000, F1: (v26 | v26), F2: v14}
                                                                    _ = v44
                                                                    _ = v45
                                                                    _ = v46
                                                                }
                                                                _ = v40
                                                            }
                                                            if (1 >= -2147483648) {
                                                                v47 := v4[42]
                                                                _ = v47
                                                            } else {
                                                                v26 += (v26 * (v26 ^ 9223372036854775807))
                                                                _ = len(v4)
                                                                v1 += (v1 & (v1 / ((42) | 1)))
                                                            }
                                                        }
                                                        switch 42 {
                                                        case 345:
                                                            v26 = v26
                                                        case 36:
                                                            v14 = v14
                                                        case 0:
                                                            delete(v4, 65535)
                                                        default:
                                                            var v48 uint32 = 0
                                                            _ = v48
                                                        }
                                                    } else {
                                                        if (true || true) {
                                                            var v49 uint32 = 0
                                                            v50 := GenS3{F0: v14}
                                                            _ = v49
                                                            _ = v50
                                                        } else {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        fmt.Println(v1, fi25)
                                                    }
                                                }
                                            }
                                        } else {
                                            switch (-(fi22)) {
                                            case 128:
                                                v26 = 0
                                            case -9223372036854775808:
                                                _ = len(v2)
                                            case 0:
                                                var v51 float64 = v3
                                                _ = v51
                                            default:
                                                v52 := GenS5{F0: false, F1: (v1 - 45), F2: (^(v15))}
                                                _ = v52
                                            }
                                            fmt.Println(v15, v14)
                                            if false {
                                                v53 := GenS1{F0: 1000, F1: v26, F2: v14}
                                                _ = v53
                                            } else {
                                                if ((916 < -31) && (!(!false))) {
                                                    var v54 int32 = 42
                                                    delete(v4, 2)
                                                    switch (fi25 | -2147483648) {
                                                    case 42:
                                                        v55 := GenS6{F0: 0, F1: "foo bar", F2: fi25}
                                                        _ = v55
                                                    case 127:
                                                        if true {
                                                            fmt.Println(v15, fi22, v3)
                                                        } else {
                                                            var v56 int8 = 75
                                                            switch fi25 {
                                                            case 1:
                                                                switch fi25 {
                                                                case 1:
                                                                    v54 = 149
                                                                case 128:
                                                                    fmt.Println(fi22, fi25)
                                                                }
                                                            case 9223372036854775807:
                                                                if ((-134 != -383) || true) {
                                                                    if false {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        fmt.Println(fi22)
                                                                        var v57 float64 = float64(3.14159)
                                                                        _ = v57
                                                                    } else {
                                                                        var v58 float64 = v3
                                                                        _ = v58
                                                                    }
                                                                    v2 = append(v2, 481)
                                                                    v14 -= (^(v14))
                                                                }
                                                            default:
                                                                v54 = 2147483647
                                                            }
                                                            switch "a" {
                                                            case "\n\t":
                                                                v59 := GenS7{F0: (v56 | v56), F1: 0, F2: (v56 / ((123) | 1))}
                                                                _ = v59
                                                            case "foo bar":
                                                                v14 -= 541
                                                            case "a":
                                                                var v60 float64 = v3
                                                                _ = v60
                                                            default:
                                                                fmt.Println(fi22, v26)
                                                            }
                                                            _ = v56
                                                        }
                                                    }
                                                    _ = v54
                                                }
                                                var v61 int = (^(fi22))
                                                _ = v61
                                            }
                                        }
                                    }
                                }
                            }
                            var v62 string = (((("\n\t" + "a") + "hello") + ("hello" + "\n\t")) + (("" + "hello") + "\n\t"))
                            _ = v62
                        default:
                            v3 -= ((v3 * v3) + (((3.14159 / ((v3 / v3) / v3)) + v3) + 0.0))
                        }
                        v63 := GenS4{F0: 42}
                        v64 := GenS3{F0: v14}
                        _ = v63
                        _ = v64
                    }
                }
                v15 += 0
                _ = fi25
                _ = v26
            }
        }
        var v65 uint64 = v14
        if ((^(fi22)) < 9223372036854775807) {
            for fi66 := 0; fi66 < 16; fi66++ {
                if len(v2) > 0 {
                    v2[0] = 2
                }
                switch (fi66 + (fi66 / ((-104) | 1))) {
                case -21:
                    var v67 uint16 = 32767
                    _ = v67
                case 95:
                    var v68 int8 = 42
                    _ = v68
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v65 -= (v14 ^ 965)
                _ = fi66
            }
            v2 = append(v2, 42)
        } else {
            fmt.Println(v65, v3, v14)
            var v69 string = ("hello" + (("hello" + "hello") + "\n\t"))
            if true {
                v14 = (^(v14))
                var v70 uint = 0
                v70 -= v1
                _ = v70
            }
            _ = v69
        }
        _ = fi22
        _ = v65
    }
    if (-355 == 0) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi71 := 0; fi71 < 11; fi71++ {
        for fi72 := 0; fi72 < 16; fi72++ {
            fmt.Println(v14, fi71, v3)
            if true {
                v73 := GenS2{F0: -682, F1: "foo bar"}
                _ = v73
            }
            _ = fi72
        }
        if (42 != int(v1)) {
            _ = len(v2)
        } else {
            delete(v4, 2)
            _ = len(v2)
            var v74 int8 = 1
            _ = v74
        }
        _ = fi71
    }
    if true {
        if true {
            v15 += v15
            for fi75 := 0; fi75 < 11; fi75++ {
                v14 -= v14
                _ = fi75
            }
            var v76 uint16 = 2
            _ = v76
        } else {
            v77 := GenS7{F0: 6, F1: 1, F2: -90}
            for fi78 := 0; fi78 < 11; fi78++ {
                if (!true) {
                    for fi79 := 0; fi79 < 11; fi79++ {
                        fmt.Println(v15, v3, fi79)
                        _ = fi79
                    }
                    if true {
                        for fi80 := 0; fi80 < 15; fi80++ {
                            var v81 int64 = 0
                            _ = fi80
                            _ = v81
                        }
                        v1 -= 132
                        v2 = append(v2, 295)
                    } else {
                        for fi82 := 0; fi82 < 20; fi82++ {
                            v14 -= 1
                            var v83 float64 = 0.1
                            _ = fi82
                            _ = v83
                        }
                    }
                } else {
                    if ((v15 >> ((1) & 63)) != v15) {
                        var v84 string = "\n\t"
                        for fi85 := 0; fi85 < 20; fi85++ {
                            var v86 int8 = -49
                            v87 := GenS6{F0: v14, F1: ("foo bar" + v84), F2: (fi85 << ((18446744073709551615) & 63))}
                            _ = fi85
                            _ = v86
                            _ = v87
                        }
                        var v88 bool = true
                        _ = v84
                        _ = v88
                    } else {
                        if (-2147483648 != (-(fi78))) {
                            v89 := GenS5{F0: (!true), F1: (v1 + v1), F2: v15}
                            var v90 string = "hello"
                            var v91 int16 = 139
                            _ = v89
                            _ = v90
                            _ = v91
                        }
                        if false {
                            for fi92 := 0; fi92 < 17; fi92++ {
                                v15 = 247
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi92
                            }
                        } else {
                            switch fi78 {
                            case 42:
                                var v93 uint16 = 65535
                                _ = v93
                            default:
                                for fi94 := 0; fi94 < 7; fi94++ {
                                    fmt.Println(v15)
                                    _ = fi94
                                }
                            }
                            for fi95 := 0; fi95 < 11; fi95++ {
                                if true {
                                    v15 = 9
                                    fmt.Println(fi95, v14)
                                }
                                v96 := GenS8{F0: (v1 - (v1 | v1)), F1: v14}
                                var v97 float32 = float32(v3)
                                _ = fi95
                                _ = v96
                                _ = v97
                            }
                        }
                        var v98 string = (("foo bar" + "") + "hello")
                        _ = v98
                    }
                    for fi99 := 0; fi99 < 4; fi99++ {
                        var v100 uint8 = (^(v15))
                        if len(v2) > 0 {
                            v2[0] = 42
                        }
                        var v101 uint16 = 32767
                        _ = fi99
                        _ = v100
                        _ = v101
                    }
                }
                _ = fi78
            }
            v3 += 1.5
            _ = v77
        }
    }
    v102 := GenS5{F0: false, F1: v1, F2: (^(v15))}
    v103 := GenS3{F0: (v14 << ((322) & 63))}
    fmt.Println(v1)
    if true {
        var v104 int32 = 1
        v1 -= v1
        _ = v104
    }
    if (("" + ("\n\t" + "hello")) < "") {
        var v105 float32 = 1e10
        if true {
            var v106 int32 = 2147483647
            fmt.Println(v105, v106, v105)
            _ = v106
        } else {
            fmt.Println(v105, v3, v105)
            fmt.Println(v14, v1)
        }
        _ = v105
    } else {
        if true {
            if (-876 != 1000) {
                if (true && false) {
                    if (42 >= 42) {
                        fmt.Println(v3)
                        var v107 int16 = 32767
                        _ = v107
                    }
                    var v108 string = (("hello" + ("\n\t" + "")) + "\n\t")
                    fmt.Println(v1, v108)
                    _ = v108
                }
                if (((true && (32767 > 702)) || (990 != 1)) || (!(((1.5 * v3) - v3) > float64(-2.25)))) {
                    var v109 uint16 = 42
                    v14 -= 4294967295
                    if true {
                        var v110 uint64 = 606
                        var v111 int32 = int32(v1)
                        _ = v110
                        _ = v111
                    }
                    _ = v109
                }
                for fi112 := 0; fi112 < 5; fi112++ {
                    var v113 string = ("a" + ("foo bar" + "\n\t"))
                    v1 -= 18446744073709551615
                    _ = fi112
                    _ = v113
                }
            }
            v3 -= 3.14159
            var v114 string = ("a" + (("a" + ("hello" + "foo bar")) + "a"))
            _ = v114
        } else {
            fmt.Println(v15)
            fmt.Println(v1, v1)
        }
        var v115 int8 = int8(v1)
        _ = v115
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v102
    _ = v103
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
