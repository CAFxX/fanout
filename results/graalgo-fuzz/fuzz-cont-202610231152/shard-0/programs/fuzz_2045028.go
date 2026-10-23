package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 int32; F1 uint8; F2 string }
type GenS8 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int32; F1 uint; F2 uint8 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint64; F1 int8 }
type GenS14 struct { F0 string; F1 uint8; F2 uint32 }
func genFn0(p0 uint8, p1 int8, p2 int) int64 {
    if true {
        var v7 int16 = 449
        p0 = 0
        _ = v7
    } else {
        if (-2.25 != 3.14159) {
            if (1.5 != 0.1) {
                v8 := GenS1{F0: 43}
                v9 := []int{-825, p2, p2}
                _ = v8
                _ = v9
            }
            var v10 int8 = -50
            v11 := GenS2{F0: 1, F1: p1}
            _ = v10
            _ = v11
        } else {
            v12 := []int{2147483647, p2, p2, 258}
            var v13 int16 = -1
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 int32 = -1
            v16 := []int{(fi14 ^ fi14), fi14, (fi14 & (fi14 + p2))}
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 uint = 255
        _ = v17
    }
    v18 := GenS2{F0: 0, F1: p1}
    v19 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return 0
}
func genFn1(p0 string) int8 {
    for fi20 := 0; fi20 < 8; fi20++ {
        p0 = p0
        _ = fi20
    }
    v21 := []int{}
    for fi22 := 0; fi22 < 11; fi22++ {
        p0 = p0
        _ = fi22
    }
    _ = p0
    _ = v21
    return 127
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 0
    v3 := []string{"", ("foo bar" + "hello"), ("\n\t" + "foo bar")}
    v4 := make(map[int]uint32)
    v5 := make(map[string]string)
    v6 := make(map[string]string)
    if (-128 == int8(v2)) {
        if true {
            if ((0.0 + (float32(0.0) - 1.5)) == 1.5) {
                var v23 string = "foo bar"
                _ = v23
            }
            var v24 int = 588
            _ = v24
        }
        v25 := v1[565]
        _ = v25
    } else {
        switch "" {
        case "a":
            v5["hello"] = ("foo bar" + "hello")
        case "hello":
            v26 := GenS3{F0: 42, F1: 415, F2: 9}
            _ = v26
        case "foo bar":
            if true {
                fmt.Println(v2)
                fmt.Println(v2)
            } else {
                v2 = 255
                var v27 uint32 = 244
                _ = v27
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v2 += 128
    if false {
        v2 -= (v2 | 255)
    }
    v2 = v2
    switch 42 {
    case 9223372036854775807:
        var v28 uint8 = (v2 / ((42) | 1))
        _ = v28
    default:
        switch 1 {
        case -190:
            var v29 uint16 = 32767
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        if ((1.5 + 0.1) != -2.25) {
            v31 := GenS1{F0: 458}
            if (-2147483648 <= 42) {
                for fi32 := 0; fi32 < 3; fi32++ {
                    v2 += (v2 ^ (v2 / ((1) | 1)))
                    v2 += v2
                    v2 += (^(v2))
                    _ = fi32
                }
                v2 = (v2 >> ((65535) & 63))
            }
            _ = v31
        } else {
            if false {
                for fi33 := 0; fi33 < 5; fi33++ {
                    if false {
                        var v34 uint64 = 42
                        if true {
                            v2 -= (v2 * v2)
                            v35 := GenS3{F0: 2, F1: -2147483648, F2: v2}
                            if false {
                                var v36 uint64 = v34
                                _ = v36
                            } else {
                                v34 -= (^(v34))
                                switch fi30 {
                                case 4294967295:
                                    if (genFn0((v2 | v2), -128, 605) >= int64(v2)) {
                                        v34 -= (v34 & (v34 >> ((42) & 63)))
                                        v37 := GenS4{F0: ((("a" + ("hello" + "foo bar")) + "foo bar") + (("\n\t" + ("" + "\n\t")) + ("a" + (("a" + ("foo bar" + ("" + ""))) + ("a" + (("" + "\n\t") + ("hello" + "\n\t"))))))), F1: 42}
                                        var v38 bool = true
                                        v1[605] = "foo bar"
                                        var v39 uint = 767
                                        var v40 string = ("foo bar" + (("foo bar" + "") + ("" + "foo bar")))
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    }
                                case -2147483648:
                                    fmt.Println(v34)
                                case 127:
                                    v41 := v4[fi33]
                                    _ = v41
                                }
                            }
                            _ = v35
                        }
                        switch (fi33 ^ (^(fi30))) {
                        case 4294967295:
                            v42 := GenS5{F0: 127}
                            _ = v42
                        case -1:
                            if false {
                                if (genFn0(127, 3, (^(fi30))) == int64(v34)) {
                                    v2 = v2
                                    fmt.Println(fi33, fi30)
                                }
                                var v43 bool = ((((0.0 * float32(3.14159)) * 3.14159) - 0.1) <= 1e10)
                                _ = v43
                            }
                        }
                        var v44 int16 = -1
                        _ = v34
                        _ = v44
                    } else {
                        if (13 >= 1) {
                            v2 += v2
                            _ = len(v3)
                            var v45 int64 = -9223372036854775808
                            _ = v45
                        }
                        var v46 uint8 = (v2 << ((444) & 63))
                        var v47 float32 = 3.14159
                        _ = v46
                        _ = v47
                    }
                    fmt.Println(v2, fi33, v2)
                    _ = fi33
                }
                switch (fi30 / (((-(fi30))) | 1)) {
                case 127:
                    var v48 uint16 = 32767
                    _ = v48
                case 2147483647:
                    for fi49 := 0; fi49 < 2; fi49++ {
                        if (3.14159 != ((0.0 - ((1.5 / 1.0) + 1e10)) - ((1e10 - (-2.25 * 3.14159)) * 1e10))) {
                            var v50 float64 = 0.1
                            _ = v50
                        } else {
                            fmt.Println(fi30)
                            switch "\n\t" {
                            case "":
                                var v51 uint16 = 42
                                _ = v51
                            case "\n\t":
                                if (65535 <= 448) {
                                    if (false && (v2 != (v2 << ((0) & 63)))) {
                                        v52 := GenS4{F0: "foo bar", F1: -128}
                                        _ = v52
                                    }
                                    _ = len(v3)
                                } else {
                                    if (false || ((4294967295 >= 251) || false)) {
                                        if ((v2 ^ v2) >= v2) {
                                            switch "foo bar" {
                                            case "hello":
                                                var v53 uint16 = 0
                                                _ = v53
                                            case "\n\t":
                                                _ = len(v3)
                                            }
                                            var v54 int16 = 1
                                            _ = v54
                                        }
                                        v2 -= v2
                                    }
                                    v2 = 34
                                    v55 := GenS3{F0: 42, F1: -2147483648, F2: v2}
                                    _ = v55
                                }
                            case "hello":
                                _ = len(v5)
                            }
                        }
                        var v56 uint = 255
                        _ = fi49
                        _ = v56
                    }
                }
            }
            for fi57 := 0; fi57 < 2; fi57++ {
                if ((!false) && (255 == 255)) {
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                    switch fi30 {
                    case 9223372036854775807:
                        v2 -= (v2 ^ v2)
                    case 4294967295:
                        if (719 == -2147483648) {
                            v2 += v2
                            fmt.Println(fi57)
                        }
                    default:
                        v2 = 255
                    }
                    v2 = (v2 / (((v2 - (v2 * (v2 / ((2) | 1))))) | 1))
                } else {
                    v2 = v2
                    if true {
                        if (692 < 65535) {
                            v2 = v2
                            v2 = v2
                            var v58 float64 = (1e10 + 1.5)
                            _ = v58
                        }
                    } else {
                        var v59 uint8 = (v2 >> ((1) & 63))
                        if ((32767 >= 0) && (true && false)) {
                            if false {
                                switch "a" {
                                case "":
                                    if (!(1.5 < 0.0)) {
                                        fmt.Println(v2, fi30)
                                    } else {
                                        if (true && (!false)) {
                                            v60 := GenS6{F0: 65535, F1: v2}
                                            _ = v60
                                        } else {
                                            var v61 int64 = 42
                                            switch ("\n\t" + "") {
                                            case "a":
                                                if true {
                                                    if ("\n\t" == "foo bar") {
                                                        if false {
                                                            var v62 float64 = 1e10
                                                            _ = v62
                                                        }
                                                        v2 -= 246
                                                    } else {
                                                        var v63 uint8 = 0
                                                        if (255 > 255) {
                                                            var v64 uint = 499
                                                            _ = v64
                                                        }
                                                        _ = v63
                                                    }
                                                } else {
                                                    switch 247 {
                                                    case 0:
                                                        v65 := GenS7{F0: int32(v61), F1: (v59 * (v59 << ((1) & 63))), F2: "\n\t"}
                                                        _ = v65
                                                    case -516:
                                                        v61 = v61
                                                    case -1:
                                                        if false {
                                                            if true {
                                                                v59 -= (v59 & (^(v2)))
                                                                v3 = append(v3, ("\n\t" + "hello"))
                                                            } else {
                                                                var v66 int8 = -25
                                                                _ = v66
                                                            }
                                                            var v67 int32 = 2147483647
                                                            _ = v67
                                                        }
                                                    }
                                                }
                                                v2 += (v59 - (v2 | v59))
                                            default:
                                                var v68 uint16 = 32767
                                                _ = v68
                                            }
                                            _ = v61
                                        }
                                        v69 := GenS4{F0: "a", F1: 1}
                                        _ = v69
                                    }
                                default:
                                    if (1e10 < (0.0 - 1.5)) {
                                        if len(v3) > 0 {
                                            v3[0] = ("\n\t" + "hello")
                                        }
                                        var v70 uint64 = 548
                                        var v71 float32 = 1.5
                                        _ = v70
                                        _ = v71
                                    } else {
                                        switch "foo bar" {
                                        case "":
                                            var v72 uint64 = 0
                                            _ = v72
                                        case "hello":
                                            if true {
                                                v73 := GenS8{F0: 200, F1: 1, F2: 0}
                                                v2 -= 2
                                                var v74 int = 9223372036854775807
                                                _ = v73
                                                _ = v74
                                            }
                                        }
                                        fmt.Println(fi30, fi57)
                                        var v75 uint16 = 255
                                        _ = v75
                                    }
                                    var v76 int16 = 599
                                    switch (fi30 ^ int(v59)) {
                                    case 128:
                                        var v77 int8 = 6
                                        _ = v77
                                    case 0:
                                        if (101 > (v2 >> ((358) & 63))) {
                                            if false {
                                                var v78 uint = 874
                                                v3 = append(v3, "\n\t")
                                                _ = v78
                                            } else {
                                                var v79 uint16 = uint16(v59)
                                                if true {
                                                    fmt.Println(fi57)
                                                }
                                                _ = v79
                                            }
                                            _ = len(v3)
                                            fmt.Println(v76)
                                        }
                                    case 4294967295:
                                        v59 = v2
                                    default:
                                        var v80 bool = false
                                        _ = v80
                                    }
                                    _ = v76
                                }
                            }
                        }
                        switch "hello" {
                        case "hello":
                            fmt.Println(fi30, fi30, v2)
                        }
                        v3 = append(v3, ("hello" + "a"))
                        _ = v59
                    }
                    if true {
                        v2 = 111
                        v2 = (v2 % (((v2 << ((65) & 63))) | 1))
                        var v81 uint8 = 255
                        v2 = v81
                        fmt.Println(v81, v2)
                        _ = v81
                    } else {
                        var v82 uint8 = v2
                        fmt.Println(v2, v2, v82)
                        _ = v82
                    }
                }
                v6[""] = (("a" + (("\n\t" + "a") + "foo bar")) + (("foo bar" + "hello") + "\n\t"))
                switch ("hello" + "\n\t") {
                case "hello":
                    var v83 float32 = float32(0.0)
                    _ = v83
                }
                _ = fi57
            }
        }
        _ = len(v1)
        for fi84 := 0; fi84 < 20; fi84++ {
            var v85 uint64 = 240
            if (-9223372036854775808 != -791) {
                var v86 float32 = -2.25
                _ = len(v3)
                var v87 float32 = (v86 / 3.14159)
                if false {
                    if (917 != 65535) {
                        var v88 string = "\n\t"
                        v2 -= (^(v2))
                        _ = v88
                    } else {
                        v87 -= ((v86 + -2.25) * (v87 + v87))
                        v2 = v2
                    }
                    v89 := GenS9{F0: 617}
                    v86 = (v87 * v87)
                    if (false || true) {
                        v87 = float32(float64(0.0))
                        if true {
                            if false {
                                switch -356 {
                                case 9223372036854775807:
                                    v2 = v2
                                case -492:
                                    var v90 int = fi84
                                    _ = v90
                                default:
                                    v87 = 0.1
                                }
                                var v91 int = (fi30 * (fi84 >> ((2) & 63)))
                                if true {
                                    switch 42 {
                                    case 2147483647:
                                        v86 = v87
                                    case 1:
                                        v86 = (float32((-2.25 + (((0.0 * 3.14159) + 3.14159) - 3.14159))) / ((((float32((((1e10 - 3.14159) / (0.1 - 3.14159)) * (-2.25 + ((0.0 / 1.0) - 1e10)))) / v86) * v87) / (1.5 + (float32(float64(-2.25)) - v86))) / v87))
                                    case -9223372036854775808:
                                        if (2 < 32767) {
                                            v87 -= 3.14159
                                            _ = len(v3)
                                            fmt.Println(v86)
                                            if ((v85 + v85) <= 2) {
                                                v2 = (v2 ^ (^(v2)))
                                                v2 = v2
                                            } else {
                                                if true {
                                                    v92 := GenS10{F0: 0}
                                                    v91 -= (v91 % ((128) | 1))
                                                    fmt.Println(v87, v86, fi30)
                                                    _ = v92
                                                }
                                            }
                                            if false {
                                                v85 += v85
                                                v87 = (((-2.25 + float32((1.5 * 1.5))) + ((1.5 - (v87 / (1e10 / 1.5))) - (v86 * v86))) * v87)
                                                if (!true) {
                                                    v91 += (^(fi30))
                                                } else {
                                                    v87 += v87
                                                }
                                            } else {
                                                fmt.Println(v85)
                                            }
                                        } else {
                                            v2 -= 75
                                        }
                                    default:
                                        if (false || (!false)) {
                                            var v93 int8 = 127
                                            var v94 int = (v91 << ((uint(v85)) & 63))
                                            switch ("foo bar" + ("a" + (("\n\t" + (("foo bar" + "") + "a")) + ""))) {
                                            case "":
                                                if (((v85 | v85) != (v85 % (((v85 >> ((2) & 63))) | 1))) || false) {
                                                    if false {
                                                        if len(v3) > 0 {
                                                            v3[0] = "\n\t"
                                                        }
                                                        var v95 uint32 = 77
                                                        if (0 > genFn0(uint8(v95), -1, v91)) {
                                                            if (0 != 42) {
                                                                if false {
                                                                    var v96 uint64 = (v85 & (v85 ^ uint64(v95)))
                                                                    v97 := GenS8{F0: -9223372036854775808, F1: v96, F2: v95}
                                                                    if (uint32(v96) >= (v95 ^ v95)) {
                                                                        var v98 uint16 = 255
                                                                        _ = len(v3)
                                                                        _ = len(v4)
                                                                        if false {
                                                                            if (!(false || true)) {
                                                                                v3 = append(v3, "a")
                                                                                v94 -= (fi84 ^ v94)
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    v3 = append(v3, "\n\t")
                                                                                case "":
                                                                                    delete(v4, fi30)
                                                                                    v93 += (-(v93))
                                                                                default:
                                                                                    v99 := v6["a"]
                                                                                    _ = v99
                                                                                }
                                                                            } else {
                                                                                v95 = (v95 >> ((869) & 63))
                                                                                var v100 float64 = -2.25
                                                                                _ = len(v4)
                                                                                _ = v100
                                                                            }
                                                                        } else {
                                                                            v3 = append(v3, "hello")
                                                                        }
                                                                        _ = v98
                                                                    } else {
                                                                        var v101 float32 = ((v86 - (v87 / 1.0)) / 3.14159)
                                                                        var v102 uint16 = 65535
                                                                        if (42 != 2147483647) {
                                                                            var v103 bool = false
                                                                            fmt.Println(fi30, v85, v102)
                                                                            _ = v103
                                                                        } else {
                                                                            if ((0 == -1) && true) {
                                                                                var v104 uint16 = v102
                                                                                _ = v104
                                                                            } else {
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            }
                                                                            var v105 uint = 164
                                                                            _ = v105
                                                                        }
                                                                        _ = v101
                                                                        _ = v102
                                                                    }
                                                                    _ = v96
                                                                    _ = v97
                                                                }
                                                                switch (fi30 % ((v94) | 1)) {
                                                                case -9223372036854775808:
                                                                    var v106 string = (((("\n\t" + "foo bar") + ("a" + ("a" + "foo bar"))) + ("\n\t" + "")) + "")
                                                                    _ = v106
                                                                default:
                                                                    v107 := GenS2{F0: 42, F1: (-(v93))}
                                                                    _ = v107
                                                                }
                                                            } else {
                                                                var v108 uint16 = 267
                                                                _ = v108
                                                            }
                                                            var v109 uint16 = 255
                                                            _ = v109
                                                        }
                                                        _ = v95
                                                    }
                                                } else {
                                                    var v110 int32 = 1
                                                    _ = v110
                                                }
                                            case "\n\t":
                                                var v111 uint8 = (v2 * v2)
                                                _ = v111
                                            default:
                                                _ = len(v6)
                                            }
                                            _ = v93
                                            _ = v94
                                        } else {
                                            var v112 uint16 = 1
                                            if (int8(v85) >= -24) {
                                                v85 = 126
                                                if false {
                                                    switch "a" {
                                                    case "\n\t":
                                                        var v113 int64 = -1
                                                        _ = v113
                                                    }
                                                } else {
                                                    fmt.Println(v86, fi30, fi84)
                                                    fmt.Println(v91, fi30, fi30)
                                                }
                                                switch 2147483647 {
                                                case 1:
                                                    var v114 int32 = 42
                                                    _ = v114
                                                }
                                            } else {
                                                fmt.Println(fi30, v86, v85)
                                            }
                                            switch (-(fi84)) {
                                            case 42:
                                                if (true || false) {
                                                    var v115 int64 = 1
                                                    v116 := GenS6{F0: 524, F1: (v2 - v2)}
                                                    fmt.Println(v86, v112)
                                                    _ = v115
                                                    _ = v116
                                                }
                                            case -153:
                                                switch ("foo bar" + "foo bar") {
                                                case "foo bar":
                                                    if (float32(float64(-2.25)) == (1.5 * (v86 + float32(3.14159)))) {
                                                        v117 := GenS11{F0: -2147483648, F1: 2, F2: v2}
                                                        if true {
                                                            v86 = (v86 / float32((0.0 - -2.25)))
                                                            var v118 uint16 = 605
                                                            delete(v5, "a")
                                                            _ = v118
                                                        } else {
                                                            if true {
                                                                var v119 int8 = int8(v2)
                                                                _ = v119
                                                            }
                                                            if true {
                                                                var v120 int = (-(fi30))
                                                                if true {
                                                                    v121 := GenS11{F0: 0, F1: 0, F2: 42}
                                                                    _ = v121
                                                                }
                                                                v122 := GenS9{F0: 9223372036854775807}
                                                                _ = v120
                                                                _ = v122
                                                            } else {
                                                                var v123 uint32 = 65535
                                                                fmt.Println(v86, v86, v91)
                                                                _ = v123
                                                            }
                                                        }
                                                        delete(v6, "foo bar")
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = v117
                                                    } else {
                                                        v87 += 1.5
                                                        v87 -= (v86 / float32((3.14159 + 0.1)))
                                                    }
                                                case "hello":
                                                    var v124 uint8 = 128
                                                    _ = v124
                                                case "\n\t":
                                                    var v125 string = ("foo bar" + "")
                                                    _ = v125
                                                default:
                                                    if false {
                                                        v126 := GenS5{F0: -128}
                                                        _ = v126
                                                    } else {
                                                        v87 = v86
                                                        v112 -= 65535
                                                        v87 = v87
                                                    }
                                                }
                                            case 2147483647:
                                                if false {
                                                    var v127 uint64 = 2
                                                    _ = v127
                                                } else {
                                                    _ = len(v1)
                                                    if (0 <= -1) {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        var v128 float32 = v86
                                                        _ = v128
                                                    }
                                                }
                                            }
                                            _ = v112
                                        }
                                    }
                                    v86 -= ((v86 + 3.14159) - (-2.25 - -2.25))
                                }
                                v2 -= v2
                                _ = v91
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = ("a" + "\n\t")
                                }
                            }
                            v129 := GenS3{F0: 2, F1: -2147483648, F2: 45}
                            _ = v129
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        var v130 int8 = 100
                        _ = v130
                    } else {
                        v131 := v1[42]
                        v85 = 4294967295
                        v132 := GenS12{F0: (fi84 - (fi84 % (((fi84 % ((fi30) | 1))) | 1)))}
                        _ = v131
                        _ = v132
                    }
                    _ = v89
                } else {
                    v86 -= 0.0
                    v133 := GenS2{F0: 65535, F1: -36}
                    _ = v133
                }
                v2 = (v2 | (v2 / ((1) | 1)))
                _ = v86
                _ = v87
            } else {
                v85 += v85
                v2 += (v2 << ((255) & 63))
            }
            _ = fi84
            _ = v85
        }
        _ = fi30
    }
    v2 = v2
    v134 := GenS9{F0: 568}
    v6[("" + "hello")] = ""
    var v135 int = -2147483648
    v135 += (v135 % ((v135) | 1))
    v135 += (v135 >> ((0) & 63))
    fmt.Println(v2)
    var v136 int64 = -298
    var v137 float32 = 1.5
    switch "\n\t" {
    case "a":
        for fi138 := 0; fi138 < 7; fi138++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if true {
                var v139 string = "\n\t"
                _ = v139
            }
            fmt.Println(v137, v136, v2)
            _ = fi138
        }
    case "":
        var v140 string = ("hello" + ("\n\t" + "\n\t"))
        _ = v140
    case "\n\t":
        if (((v135 + v135) >= v135) || true) {
            v141 := GenS11{F0: 1, F1: 2, F2: (^(v2))}
            switch ("\n\t" + (("foo bar" + "foo bar") + ("foo bar" + "hello"))) {
            case "":
                v142 := GenS13{F0: 2, F1: -122}
                _ = v142
            case "\n\t":
                v143 := GenS14{F0: "hello", F1: (^(v2)), F2: 559}
                _ = v143
            case "foo bar":
                var v144 float64 = 0.1
                _ = v144
            }
            _ = v141
        } else {
            v2 += (v2 * v2)
        }
    default:
        _ = len(v3)
    }
    fmt.Println(v136)
    switch "" {
    case "\n\t":
        var v145 int64 = (^(v136))
        _ = v145
    case "":
        fmt.Println(v137, v137)
    case "a":
        delete(v4, -241)
    default:
        v146 := v6[""]
        _ = v146
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v135:", v135)
    fmt.Println("v136:", v136)
    fmt.Println("v137:", v137)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v134
    _ = v135
    _ = v136
    _ = v137
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
