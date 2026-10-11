package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint32; F2 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 bool; F2 uint16 }
type GenS5 struct { F0 int8; F1 string; F2 bool }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8; F1 int16; F2 int16 }
type GenS9 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS10 struct { F0 uint16; F1 string; F2 uint16 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 bool; F1 int16 }
type GenS13 struct { F0 uint; F1 uint; F2 string }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 string; F1 string; F2 int64 }
type GenS16 struct { F0 uint8; F1 string; F2 uint16 }
type GenS17 struct { F0 uint8; F1 string }
type GenS18 struct { F0 uint8 }
type GenS19 struct { F0 string; F1 bool }
type GenS20 struct { F0 int64; F1 string }
type GenS21 struct { F0 uint32; F1 int8 }
type GenS22 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint64) int32 {
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := []int{fi5, fi5, (fi5 % (((fi5 - (-(fi5)))) | 1))}
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 uint = 42
            var v9 int32 = -1
            if false {
                var v10 int64 = 0
                v11 := GenS1{F0: v9}
                if (2 != 42) {
                    switch (fi5 | fi7) {
                    case -9223372036854775808:
                        var v12 uint64 = (p0 / (((p0 % ((18446744073709551615) | 1))) | 1))
                        _ = v12
                    default:
                        v13 := GenS1{F0: (-(v9))}
                        _ = v13
                    }
                }
                if (100 >= int8(v10)) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                fmt.Println(fi7, v9)
                _ = v10
                _ = v11
            } else {
                if false {
                    if false {
                        v14 := GenS2{F0: 226, F1: 65535, F2: -14}
                        switch (("a" + "") + "\n\t") {
                        case "a":
                            p0 += p0
                        case "hello":
                            _ = len(v6)
                        }
                        _ = v14
                    } else {
                        v8 = v8
                        fmt.Println(fi7, fi5)
                    }
                    var v15 uint8 = 2
                    _ = v15
                }
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    p0 = (p0 << ((0) & 63))
    var v16 int64 = -9223372036854775808
    _ = p0
    _ = v16
    return -2147483648
}
func genFn1(p0 int32, p1 int) string {
    v17 := []int{128, -163, -825, (p1 + p1)}
    if (!false) {
        p1 = p1
        var v18 float32 = 0.1
        _ = v18
    } else {
        var v19 float32 = -2.25
        _ = v19
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v17
    return "foo bar"
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 446
    v3 := make(map[string]bool)
    var v4 int8 = -122
    var v20 uint64 = 0
    if (798 >= int(v4)) {
        switch 127 {
        case 9223372036854775807:
            fmt.Println(v20)
        case 4294967295:
            v20 += v20
        default:
            v21 := GenS1{F0: -10}
            _ = v21
        }
    } else {
        if ((v20 | (^(v20))) >= (^(v20))) {
            if false {
                var v22 uint64 = v20
                switch "foo bar" {
                case "a":
                    if false {
                        var v23 int16 = 42
                        v24 := []int{32768, 1, 128}
                        _ = v23
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 11; fi25++ {
                            v26 := GenS3{F0: (-(fi25))}
                            v27 := []int{}
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        fmt.Println(v22)
                        v22 = v20
                    }
                case "\n\t":
                    var v28 uint = 255
                    _ = v28
                default:
                    for fi29 := 0; fi29 < 14; fi29++ {
                        v30 := []int{fi29}
                        fmt.Println(v2, v4)
                        _ = fi29
                        _ = v30
                    }
                }
                _ = v22
            } else {
                v31 := GenS2{F0: -32768, F1: 708, F2: 42}
                v32 := GenS2{F0: 42, F1: 934, F2: 0}
                _ = v31
                _ = v32
            }
            var v33 uint = 79
            _ = v33
        } else {
            switch "\n\t" {
            case "hello":
                for fi34 := 0; fi34 < 3; fi34++ {
                    v35 := GenS3{F0: fi34}
                    if true {
                        v20 += 508
                        var v36 string = ""
                        if false {
                            v36 = v36
                            if (!false) {
                                v2 += (v2 / (((v2 | (^(v2)))) | 1))
                                var v37 int64 = 865
                                _ = v37
                            } else {
                                for fi38 := 0; fi38 < 3; fi38++ {
                                    v39 := GenS4{F0: 1000, F1: true, F2: v2}
                                    switch v36 {
                                    case "hello":
                                        var v40 uint16 = 65535
                                        _ = v40
                                    default:
                                        var v41 int64 = 547
                                        _ = v41
                                    }
                                    v42 := GenS5{F0: v4, F1: (("foo bar" + v36) + v36), F2: true}
                                    v36 = (v36 + v36)
                                    if true {
                                        switch fi34 {
                                        case 0:
                                            if false {
                                                v43 := []int{fi38, fi34}
                                                var v44 uint64 = (v20 ^ 4294967295)
                                                if false {
                                                    switch 402 {
                                                    case 42:
                                                        var v45 uint16 = (v2 - 823)
                                                        _ = v45
                                                    case 4294967295:
                                                        var v46 bool = true
                                                        _ = v46
                                                    }
                                                    var v47 string = v36
                                                    if true {
                                                        if true {
                                                            v4 += (v4 / ((v4) | 1))
                                                        }
                                                        if ((v44 - v20) < (^(v20))) {
                                                            if true {
                                                                var v48 int = (fi34 ^ (fi34 & fi38))
                                                                v49 := GenS4{F0: -1, F1: ((v4 + (v4 >> ((42) & 63))) <= 0), F2: 559}
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                        } else {
                                                            v50 := GenS6{F0: v47, F1: 42}
                                                            v2 = (v2 % ((v2) | 1))
                                                            fmt.Println(v4, v2, fi34)
                                                            _ = v50
                                                        }
                                                    }
                                                    _ = v47
                                                } else {
                                                    v36 = v36
                                                    v1[0] = fi34
                                                }
                                                _ = v43
                                                _ = v44
                                            } else {
                                                switch fi34 {
                                                case 9223372036854775807:
                                                    fmt.Println(v2, v20)
                                                default:
                                                    switch v36 {
                                                    case "hello":
                                                        switch "" {
                                                        case "\n\t":
                                                            if (v2 <= v2) {
                                                                switch v36 {
                                                                case "a":
                                                                    var v51 uint8 = 2
                                                                    _ = v51
                                                                case "foo bar":
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        if (2147483647 >= int32(v2)) {
                                                                            var v52 uint16 = v2
                                                                            _ = v52
                                                                        }
                                                                    default:
                                                                        var v53 int8 = v4
                                                                        _ = v53
                                                                    }
                                                                case "\n\t":
                                                                    v36 = (("a" + "\n\t") + "hello")
                                                                }
                                                                var v54 bool = (false && false)
                                                                _ = v54
                                                            } else {
                                                                var v55 float64 = float64(1.5)
                                                                v56 := v1[fi38]
                                                                v36 = "a"
                                                                _ = v55
                                                                _ = v56
                                                            }
                                                        default:
                                                            fmt.Println(v4)
                                                        }
                                                    default:
                                                        if true {
                                                            if (!(!true)) {
                                                                if false {
                                                                    v4 += (v4 / ((-102) | 1))
                                                                    switch (((v36 + (v36 + "a")) + ("foo bar" + (v36 + "hello"))) + "\n\t") {
                                                                    case "hello":
                                                                        v57 := []int{}
                                                                        _ = v57
                                                                    case "":
                                                                        v1[fi34] = -775
                                                                    default:
                                                                        v4 = (v4 - v4)
                                                                    }
                                                                    v58 := GenS6{F0: "foo bar", F1: 128}
                                                                    _ = v58
                                                                } else {
                                                                    v36 = v36
                                                                    v4 = v4
                                                                }
                                                            }
                                                        } else {
                                                            fmt.Println(fi38, fi38)
                                                            var v59 float32 = -2.25
                                                            v20 = (^(v20))
                                                            _ = v59
                                                        }
                                                    }
                                                }
                                                var v60 uint64 = (^(v20))
                                                if true {
                                                    switch v36 {
                                                    case "\n\t":
                                                        v60 -= v60
                                                    }
                                                } else {
                                                    v4 = (v4 << ((65535) & 63))
                                                    if true {
                                                        var v61 float64 = (float64(1.5) / 1.0)
                                                        var v62 int64 = -9223372036854775808
                                                        if false {
                                                            if false {
                                                                var v63 uint16 = 1
                                                                v61 = v61
                                                                _ = v63
                                                            }
                                                        } else {
                                                            v64 := GenS7{F0: 65535}
                                                            _ = v64
                                                        }
                                                        switch "foo bar" {
                                                        case "hello":
                                                            if false {
                                                                var v65 int8 = (^(v4))
                                                                var v66 int8 = (v4 << ((0) & 63))
                                                                switch (fi34 | fi34) {
                                                                case 128:
                                                                    var v67 int32 = -2147483648
                                                                    _ = v67
                                                                case 2:
                                                                    switch v36 {
                                                                    case "hello":
                                                                        if false {
                                                                            if false {
                                                                                v4 += (v66 ^ (v66 * (v4 + v4)))
                                                                                var v68 uint8 = 0
                                                                                _ = v68
                                                                            }
                                                                            v60 = 2
                                                                            if false {
                                                                                if true {
                                                                                    _ = len(v3)
                                                                                    var v69 uint32 = 293
                                                                                    fmt.Println(v60, fi38, v36)
                                                                                    _ = v69
                                                                                }
                                                                                fmt.Println(v36, fi38, v2)
                                                                                var v70 uint64 = v20
                                                                                fmt.Println(v62, v61)
                                                                                switch v36 {
                                                                                case "foo bar":
                                                                                    if ((918 <= 4294967295) || false) {
                                                                                        var v71 uint16 = v2
                                                                                        _ = v71
                                                                                    }
                                                                                }
                                                                                _ = len(v3)
                                                                                _ = v70
                                                                            } else {
                                                                                v66 = (v4 * -38)
                                                                            }
                                                                        } else {
                                                                            v4 = 127
                                                                            if true {
                                                                                var v72 float32 = 0.1
                                                                                fmt.Println(fi38, v72)
                                                                                _ = v72
                                                                            } else {
                                                                                v73 := GenS8{F0: v66, F1: 42, F2: -32768}
                                                                                _ = v73
                                                                            }
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                v74 := GenS9{F0: (true && (("foo bar" + "foo bar") > v36)), F1: (^(v20)), F2: 1}
                                                                                _ = v74
                                                                            case "hello":
                                                                                switch 42 {
                                                                                case 116:
                                                                                    var v75 bool = (0 >= 0)
                                                                                    _ = v75
                                                                                case 9223372036854775807:
                                                                                    switch v36 {
                                                                                    case "hello":
                                                                                        fmt.Println(v20, v62)
                                                                                    case "foo bar":
                                                                                        if true {
                                                                                            v61 = 3.14159
                                                                                        }
                                                                                    default:
                                                                                        var v76 uint64 = 983
                                                                                        _ = v76
                                                                                    }
                                                                                    var v77 float32 = 1e10
                                                                                    _ = len(v3)
                                                                                    fmt.Println(fi38, v65)
                                                                                    var v78 string = v36
                                                                                    _ = v77
                                                                                    _ = v78
                                                                                default:
                                                                                    fmt.Println(v66, v61)
                                                                                }
                                                                            case "\n\t":
                                                                                var v79 bool = true
                                                                                _ = v79
                                                                            }
                                                                            fmt.Println(fi34, v4, v66)
                                                                            var v80 string = "foo bar"
                                                                            _ = v80
                                                                        }
                                                                    case "a":
                                                                        v20 -= v60
                                                                    }
                                                                case -2147483648:
                                                                    v81 := GenS6{F0: "foo bar", F1: 125}
                                                                    _ = v81
                                                                default:
                                                                    v82 := GenS10{F0: (v2 % ((0) | 1)), F1: "foo bar", F2: v2}
                                                                    if ((3.14159 + -2.25) == 1.5) {
                                                                        switch (fi34 % ((fi38) | 1)) {
                                                                        case -9223372036854775808:
                                                                            var v83 int8 = (v4 + (v66 << ((4294967295) & 63)))
                                                                            _ = v83
                                                                        case 32767:
                                                                            var v84 uint16 = (v2 * (v2 >> ((4294967295) & 63)))
                                                                            _ = v84
                                                                        case -2147483648:
                                                                            switch (fi34 - fi38) {
                                                                            case 4294967295:
                                                                                delete(v3, "a")
                                                                            case 966:
                                                                                v85 := []int{fi38, (^(fi34))}
                                                                                _ = v85
                                                                            case 888:
                                                                                var v86 int32 = int32(v20)
                                                                                _ = v86
                                                                            default:
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    _ = len(v1)
                                                                                case "":
                                                                                    v20 = (v60 | (v20 >> ((2) & 63)))
                                                                                }
                                                                                _ = len(v1)
                                                                            }
                                                                        default:
                                                                            v20 = 2
                                                                        }
                                                                    }
                                                                    if (!(true && false)) {
                                                                        if ((((fi38 >= (-(fi34))) && true) && ((v62 % ((0) | 1)) == v62)) || false) {
                                                                            if false {
                                                                                v36 = "\n\t"
                                                                            } else {
                                                                                v3["foo bar"] = (true && true)
                                                                            }
                                                                            v66 -= v4
                                                                            v3["hello"] = (620 != 84)
                                                                        }
                                                                        if false {
                                                                            v3["\n\t"] = (!true)
                                                                        } else {
                                                                            if true {
                                                                                fmt.Println(v4, v61, v4)
                                                                                v66 = v4
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            v87 := []int{}
                                                                            _ = v87
                                                                        } else {
                                                                            var v88 int = fi34
                                                                            v65 -= (v4 & v4)
                                                                            if ((false && (2147483647 >= 113)) && true) {
                                                                                v20 -= (v20 & v60)
                                                                                v89 := []int{-536, -2147483648, 0}
                                                                                _ = v89
                                                                            } else {
                                                                                if false {
                                                                                    switch (fi34 / ((-499) | 1)) {
                                                                                    case -951:
                                                                                        v90 := GenS11{F0: (1 > 1)}
                                                                                        _ = v90
                                                                                    case -9223372036854775808:
                                                                                        if true {
                                                                                            if false {
                                                                                                switch v36 {
                                                                                                case "foo bar":
                                                                                                    var v91 uint64 = 18446744073709551615
                                                                                                    _ = v91
                                                                                                case "\n\t":
                                                                                                    var v92 uint = 18446744073709551615
                                                                                                    v36 = ((v36 + v36) + v36)
                                                                                                    _ = v92
                                                                                                default:
                                                                                                    fmt.Println(v2)
                                                                                                }
                                                                                                if ((v60 | (v20 << ((966) & 63))) <= v60) {
                                                                                                    if false {
                                                                                                        fmt.Println(v88, v60, fi38)
                                                                                                        var v93 int64 = (v62 ^ -9223372036854775808)
                                                                                                        _ = v93
                                                                                                    }
                                                                                                    v94 := []int{fi38, fi38}
                                                                                                    v62 += v62
                                                                                                    var v95 uint16 = v2
                                                                                                    _ = v94
                                                                                                    _ = v95
                                                                                                }
                                                                                                v96 := GenS4{F0: 32767, F1: (!(!((-(fi38)) == (fi34 >> ((2) & 63))))), F2: v2}
                                                                                                _ = v96
                                                                                            } else {
                                                                                                v60 = v20
                                                                                                switch (v88 * fi38) {
                                                                                                case 4294967295:
                                                                                                    fmt.Println(v20, v62, v20)
                                                                                                case 32768:
                                                                                                    fmt.Println(v62, v20, v2)
                                                                                                }
                                                                                            }
                                                                                            var v97 uint8 = 42
                                                                                            v60 = (v20 ^ (v20 - v20))
                                                                                            _ = v97
                                                                                        } else {
                                                                                            if true {
                                                                                                var v98 uint8 = 127
                                                                                                _ = v98
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        var v99 uint16 = 2
                                                                                        fmt.Println(v62, v65)
                                                                                        _ = v99
                                                                                    }
                                                                                    v100 := GenS12{F0: true, F1: 32767}
                                                                                    v101 := GenS13{F0: 4294967295, F1: 0, F2: (v36 + v36)}
                                                                                    _ = v100
                                                                                    _ = v101
                                                                                } else {
                                                                                    var v102 float32 = float32((((v61 + v61) - 1.5) - (v61 * (v61 / 1.5))))
                                                                                    _ = v102
                                                                                }
                                                                            }
                                                                            _ = v88
                                                                        }
                                                                    }
                                                                    fmt.Println(v66)
                                                                    _ = v82
                                                                }
                                                                v36 = "a"
                                                                _ = v65
                                                                _ = v66
                                                            } else {
                                                                if false {
                                                                    v103 := GenS10{F0: v2, F1: v36, F2: v2}
                                                                    _ = v103
                                                                }
                                                                v61 -= (float64((-2.25 - float32(v61))) - 1.5)
                                                                v104 := GenS12{F0: (true && false), F1: -1}
                                                                _ = v104
                                                            }
                                                        case "a":
                                                            var v105 bool = (70 == 128)
                                                            _ = v105
                                                        case "foo bar":
                                                            _ = len(v1)
                                                            v106 := []int{(fi34 % ((fi34) | 1)), 2147483647, (-(fi34))}
                                                            _ = v106
                                                        }
                                                        var v107 float32 = 1.5
                                                        var v108 uint16 = v2
                                                        _ = v107
                                                        _ = v108
                                                        _ = v61
                                                        _ = v62
                                                    } else {
                                                        fmt.Println(fi34, v60, v4)
                                                        v20 = (v20 >> ((4294967295) & 63))
                                                    }
                                                    v20 += (^(v20))
                                                    switch (fi34 / ((fi34) | 1)) {
                                                    case -1:
                                                        fmt.Println(fi38, v60)
                                                    case -2147483648:
                                                        v2 = 32767
                                                    case 898:
                                                        v60 -= 904
                                                    default:
                                                        fmt.Println(v60, v20)
                                                    }
                                                }
                                                _ = v60
                                            }
                                        case 42:
                                            if ((((float64((0.0 - 0.0)) - -2.25) + (float64(1.5) / ((0.0 * float64(3.14159)) - (0.0 - -2.25)))) - 1.5) >= float64(0.0)) {
                                                var v109 float32 = (0.0 - (-2.25 - (1e10 - (0.0 - (float32(0.0) - -2.25)))))
                                                var v110 uint = 2
                                                v4 += (v4 - (^(v4)))
                                                _ = v109
                                                _ = v110
                                            } else {
                                                v3["a"] = false
                                                switch (((((("a" + (((v36 + ((v36 + "\n\t") + "a")) + (((v36 + ((v36 + "hello") + v36)) + v36) + (v36 + ((v36 + (v36 + v36)) + v36)))) + "\n\t")) + v36) + v36) + "foo bar") + ("" + (v36 + ((v36 + ("hello" + v36)) + "foo bar")))) + "a") {
                                                case "foo bar":
                                                    var v111 int = (^(fi38))
                                                    _ = v111
                                                case "a":
                                                    v20 = v20
                                                default:
                                                    var v112 uint64 = v20
                                                    _ = v112
                                                }
                                                var v113 float64 = ((0.0 - (0.1 + 0.1)) * 1.5)
                                                v114 := GenS12{F0: (!false), F1: 712}
                                                v115 := []int{}
                                                v116 := v1[fi34]
                                                if (0.1 < 3.14159) {
                                                    v117 := GenS7{F0: 390}
                                                    _ = v117
                                                } else {
                                                    if false {
                                                        var v118 float64 = v113
                                                        v113 = 0.0
                                                        _ = v118
                                                    }
                                                    v20 = v20
                                                }
                                                _ = v113
                                                _ = v114
                                                _ = v115
                                                _ = v116
                                            }
                                        default:
                                            switch v36 {
                                            case "hello":
                                                v4 -= -5
                                                v2 -= 255
                                                var v119 int64 = 150
                                                _ = v119
                                            case "\n\t":
                                                v120 := GenS14{F0: 1}
                                                _ = v120
                                            case "":
                                                fmt.Println(fi34, v36)
                                            }
                                        }
                                    } else {
                                        v20 += v20
                                        v36 = ("foo bar" + v36)
                                        var v121 float64 = (3.14159 - ((0.0 + 0.0) + (1.5 / 0.1)))
                                        _ = v121
                                    }
                                    fmt.Println(fi34)
                                    switch fi38 {
                                    case -631:
                                        var v122 int = -974
                                        _ = v122
                                    default:
                                        switch (v36 + ((v36 + v36) + v36)) {
                                        case "foo bar":
                                            var v123 uint32 = uint32(v2)
                                            _ = v123
                                        case "\n\t":
                                            var v124 bool = false
                                            _ = v124
                                        }
                                    }
                                    _ = fi38
                                    _ = v39
                                    _ = v42
                                }
                            }
                        }
                        _ = v36
                    }
                    _ = fi34
                    _ = v35
                }
            }
            var v125 uint16 = 32767
            switch "foo bar" {
            case "hello":
                v2 = (v2 % ((v125) | 1))
            default:
                for fi126 := 0; fi126 < 14; fi126++ {
                    v127 := []int{}
                    _ = fi126
                    _ = v127
                }
            }
            _ = v125
        }
        if (-937 == 430) {
            if true {
                var v128 float64 = (((float64((1.5 + 0.0)) * 0.0) - ((1.5 + 1.5) + float64((0.1 / 1.0)))) - 0.0)
                v129 := GenS15{F0: "", F1: "", F2: -1}
                _ = v128
                _ = v129
            } else {
                var v130 uint = 2
                fmt.Println(v130, v20)
                v131 := GenS16{F0: 134, F1: "a", F2: (v2 - uint16(v20))}
                _ = v130
                _ = v131
            }
            for fi132 := 0; fi132 < 20; fi132++ {
                v133 := []int{(fi132 << ((255) & 63)), 549, 128, (^(fi132))}
                for fi134 := 0; fi134 < 7; fi134++ {
                    if len(v133) >= 2 {
                        v133 = v133[1:len(v133)]
                    }
                    v133 = append(v133, 127)
                    fmt.Println(fi134, fi134, v4)
                    _ = fi134
                }
                var v135 bool = (true && false)
                _ = fi132
                _ = v133
                _ = v135
            }
        } else {
            if false {
                v136 := []int{}
                for fi137 := 0; fi137 < 13; fi137++ {
                    for fi138 := 0; fi138 < 3; fi138++ {
                        v4 -= (v4 >> ((767) & 63))
                        _ = fi138
                    }
                    var v139 uint64 = 186
                    _ = fi137
                    _ = v139
                }
                v4 = v4
                _ = v136
            }
        }
        v4 += (v4 - (-(v4)))
    }
    v140 := GenS17{F0: 0, F1: ""}
    _ = len(v3)
    v141 := GenS6{F0: ("" + ""), F1: 127}
    v4 += (^(v4))
    var v142 float64 = float64(float32(1e10))
    if true {
        v2 -= v2
    } else {
        var v143 int64 = -1
        _ = v143
    }
    v144 := GenS14{F0: 42}
    v145 := GenS8{F0: int8(v20), F1: -30, F2: 42}
    if (v20 == v20) {
        v20 += 2
        v146 := []int{469, 28, 760}
        _ = v146
    }
    _ = len(v1)
    var v147 uint32 = 42
    v20 -= v20
    if (-124 != v4) {
        switch ((("a" + "\n\t") + ("hello" + "hello")) + ("a" + ("" + "a"))) {
        case "hello":
            for fi148 := 0; fi148 < 12; fi148++ {
                if (false && (true && (!true))) {
                    if (true || false) {
                        switch ((("" + ("hello" + "foo bar")) + "foo bar") + ("a" + ("hello" + ((("\n\t" + ("hello" + ("\n\t" + "foo bar"))) + "\n\t") + "\n\t")))) {
                        case "a":
                            var v149 uint16 = 255
                            _ = v149
                        default:
                            delete(v1, fi148)
                        }
                        switch -739 {
                        case 360:
                            var v150 uint = 4294967295
                            _ = v150
                        case 583:
                            fmt.Println(fi148)
                        default:
                            if false {
                                v151 := GenS18{F0: 127}
                                v152 := []int{(fi148 ^ fi148), 937, (fi148 / (((fi148 * fi148)) | 1)), 953}
                                _ = v151
                                _ = v152
                            } else {
                                v4 += -1
                            }
                        }
                        v153 := GenS1{F0: 793}
                        _ = v153
                    }
                    v154 := GenS11{F0: (0.0 <= (((-2.25 + 1e10) + -2.25) / 1.5))}
                    for fi155 := 0; fi155 < 5; fi155++ {
                        var v156 float64 = v142
                        v157 := GenS19{F0: ("hello" + ""), F1: false}
                        var v158 string = ("a" + ("hello" + "foo bar"))
                        _ = fi155
                        _ = v156
                        _ = v157
                        _ = v158
                    }
                    _ = v154
                } else {
                    var v159 string = ""
                    _ = v159
                }
                if true {
                    if false {
                        fmt.Println(v20)
                        fmt.Println(v147, v20, fi148)
                    } else {
                        for fi160 := 0; fi160 < 16; fi160++ {
                            var v161 uint = 348
                            var v162 uint64 = (v20 ^ (^(v20)))
                            _ = fi160
                            _ = v161
                            _ = v162
                        }
                        switch 4294967295 {
                        case 0:
                            v163 := GenS16{F0: uint8(v2), F1: "foo bar", F2: v2}
                            _ = v163
                        default:
                            v164 := GenS20{F0: 172, F1: "hello"}
                            _ = v164
                        }
                    }
                    for fi165 := 0; fi165 < 5; fi165++ {
                        if false {
                            var v166 float32 = 0.0
                            fmt.Println(v142, v2)
                            if true {
                                v2 -= (^(v2))
                            } else {
                                switch fi148 {
                                case 9223372036854775807:
                                    if (-2147483648 > -916) {
                                        fmt.Println(v2, v147, v2)
                                        v167 := GenS21{F0: 0, F1: (v4 ^ (v4 << ((470) & 63)))}
                                        if true {
                                            v168 := []int{}
                                            _ = len(v3)
                                            var v169 uint32 = 2147483648
                                            _ = v168
                                            _ = v169
                                        }
                                        _ = v167
                                    } else {
                                        v147 += 925
                                    }
                                default:
                                    v170 := GenS11{F0: (4294967295 != 2)}
                                    _ = v170
                                }
                                v2 = v2
                            }
                            var v171 uint64 = v20
                            _ = v166
                            _ = v171
                        }
                        var v172 uint32 = (v147 % ((2147483648) | 1))
                        _ = fi165
                        _ = v172
                    }
                } else {
                    switch (fi148 << ((65535) & 63)) {
                    case 32768:
                        var v173 int8 = (^(v4))
                        _ = v173
                    default:
                        if true {
                            if (true && true) {
                                v20 -= v20
                                if false {
                                    var v174 float64 = 0.0
                                    for fi175 := 0; fi175 < 4; fi175++ {
                                        v176 := []int{(fi175 + (fi148 - (^(fi148))))}
                                        _ = fi175
                                        _ = v176
                                    }
                                    v177 := v1[fi148]
                                    _ = v174
                                    _ = v177
                                }
                                v178 := []int{(fi148 - 703), (fi148 | 2147483647), fi148, 112}
                                _ = v178
                            } else {
                                v179 := GenS9{F0: true, F1: (v20 & v20), F2: 18446744073709551615}
                                _ = len(v3)
                                _ = v179
                            }
                            for fi180 := 0; fi180 < 12; fi180++ {
                                delete(v3, "")
                                _ = fi180
                            }
                            v181 := GenS9{F0: (v142 < v142), F1: v20, F2: (v20 ^ (^(v20)))}
                            _ = v181
                        } else {
                            v182 := GenS22{F0: 42, F1: 620}
                            var v183 uint16 = (^(v2))
                            _ = v182
                            _ = v183
                        }
                    }
                }
                _ = fi148
            }
        case "\n\t":
            for fi184 := 0; fi184 < 5; fi184++ {
                fmt.Println(v2, v147, v147)
                for fi185 := 0; fi185 < 1; fi185++ {
                    v186 := []int{26, fi184, (fi184 + 4294967295)}
                    var v187 uint64 = v20
                    fmt.Println(v187)
                    v4 = 100
                    _ = fi185
                    _ = v186
                    _ = v187
                }
                _ = fi184
            }
        default:
            switch -9223372036854775808 {
            case 1:
                for fi188 := 0; fi188 < 15; fi188++ {
                    v2 = 1
                    _ = fi188
                }
            case -1:
                v2 = v2
            case 32768:
                for fi189 := 0; fi189 < 3; fi189++ {
                    var v190 int16 = 1
                    v191 := []int{(fi189 * (fi189 - 2147483647)), 42, (fi189 & (fi189 % ((fi189) | 1)))}
                    var v192 int8 = (v4 - (v4 >> ((0) & 63)))
                    _ = fi189
                    _ = v190
                    _ = v191
                    _ = v192
                }
            }
        }
        var v193 uint = 792
        _ = v193
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v142:", v142)
    fmt.Println("v147:", v147)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v140
    _ = v141
    _ = v142
    _ = v144
    _ = v145
    _ = v147
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
