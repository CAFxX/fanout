package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int64; F2 uint32 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint16; F1 uint32 }
type GenS11 struct { F0 int32; F1 uint }
type GenS12 struct { F0 int; F1 uint64 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int16; F1 int16; F2 uint }
type GenS17 struct { F0 int16; F1 uint32 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 int64; F1 int32 }
type GenS20 struct { F0 string }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 int; F1 int }
type GenS23 struct { F0 int16; F1 int16 }
type GenS24 struct { F0 uint8 }
func genFn0(p0 bool, p1 uint8) int16 {
    v4 := []int{4294967295, -2147483648, -1}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1(p0 uint16) uint16 {
    switch "foo bar" {
    case "":
        v5 := GenS1{F0: 42}
        _ = v5
    }
    var v6 float32 = 3.14159
    switch (("\n\t" + "hello") + "foo bar") {
    case "":
        var v7 float32 = (((1e10 / (1.5 - 0.1)) + (v6 * ((1e10 + (float32((0.1 / 1.0)) * (1.5 - (float32((1.5 - (0.0 / 1.5))) / v6)))) * (0.1 - ((v6 / v6) / (v6 / v6)))))) + 1.5)
        _ = v7
    case "a":
        if (true && ((1000 == -512) && ((false && true) && (!(true && (503 < 32767)))))) {
            for fi8 := 0; fi8 < 4; fi8++ {
                v9 := GenS1{F0: 2}
                if (("foo bar" + "hello") != "\n\t") {
                    for fi10 := 0; fi10 < 8; fi10++ {
                        if (0 != -719) {
                            v11 := GenS2{F0: (fi10 << ((255) & 63))}
                            v12 := []int{-1}
                            p0 = p0
                            var v13 int8 = -1
                            if true {
                                v12 = append(v12, int(v13))
                                var v14 uint8 = 255
                                _ = v14
                            }
                            if false {
                                if false {
                                    fmt.Println(v13)
                                    if (2147483647 > 2147483648) {
                                        v6 = v6
                                        var v15 float64 = -2.25
                                        _ = v15
                                    } else {
                                        var v16 int16 = -32768
                                        _ = v16
                                    }
                                    if true {
                                        if false {
                                            if true {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if (!true) {
                                                        var v17 uint64 = 18446744073709551615
                                                        var v18 uint64 = 623
                                                        v13 += -1
                                                        var v19 string = "\n\t"
                                                        v12 = append(v12, int(v13))
                                                        _ = v17
                                                        _ = v18
                                                        _ = v19
                                                    } else {
                                                        switch (("" + (((("" + "") + ("\n\t" + "\n\t")) + (("a" + "hello") + ("\n\t" + "hello"))) + ("a" + ("foo bar" + "hello")))) + "a") {
                                                        case "foo bar":
                                                            switch fi10 {
                                                            case -9223372036854775808:
                                                                switch fi10 {
                                                                case -695:
                                                                    switch (fi8 & fi8) {
                                                                    case 42:
                                                                        var v20 int32 = -370
                                                                        v12 = append(v12, 32767)
                                                                        v20 -= v20
                                                                        _ = v20
                                                                    case 1:
                                                                        v21 := GenS3{F0: 42, F1: 784, F2: 2}
                                                                        _ = v21
                                                                    default:
                                                                        p0 += 583
                                                                    }
                                                                    v13 = (v13 & 127)
                                                                }
                                                            case 42:
                                                                fmt.Println(p0, v6, fi8)
                                                            default:
                                                                v13 -= (v13 / ((v13) | 1))
                                                            }
                                                        case "hello":
                                                            var v22 uint8 = 255
                                                            _ = v22
                                                        }
                                                        var v23 string = ""
                                                        var v24 int = fi8
                                                        _ = v23
                                                        _ = v24
                                                    }
                                                case "a":
                                                    if len(v12) >= 2 {
                                                        v12 = v12[1:len(v12)]
                                                    }
                                                case "hello":
                                                    v25 := GenS4{F0: -1, F1: 2, F2: 255}
                                                    _ = v25
                                                default:
                                                    if true {
                                                        fmt.Println(v13, fi8)
                                                        v26 := GenS1{F0: 2147483647}
                                                        fmt.Println(v13)
                                                        _ = v26
                                                    }
                                                }
                                            } else {
                                                v12 = append(v12, fi10)
                                                v27 := GenS5{F0: 2}
                                                var v28 int = fi8
                                                _ = v27
                                                _ = v28
                                            }
                                            if ("foo bar" != ("" + "")) {
                                                fmt.Println(v6, p0, fi10)
                                            } else {
                                                v29 := GenS2{F0: fi10}
                                                if len(v12) > 0 {
                                                    v12[0] = 9223372036854775807
                                                }
                                                _ = v29
                                            }
                                            var v30 bool = (-128 < -1)
                                            _ = v30
                                        } else {
                                            switch "hello" {
                                            case "foo bar":
                                                if false {
                                                    var v31 string = ("" + "foo bar")
                                                    v31 = (v31 + (v31 + "a"))
                                                    fmt.Println(v13, v31)
                                                    _ = v31
                                                }
                                            }
                                            if false {
                                                var v32 int8 = (v13 % (((v13 + v13)) | 1))
                                                var v33 uint16 = (p0 ^ 0)
                                                switch "" {
                                                case "a":
                                                    if true {
                                                        switch 32767 {
                                                        case -383:
                                                            switch "hello" {
                                                            case "hello":
                                                                if len(v12) > 0 {
                                                                    v12[0] = (fi8 % ((fi8) | 1))
                                                                }
                                                            default:
                                                                v13 -= v32
                                                            }
                                                        case 159:
                                                            v33 += (v33 >> ((2) & 63))
                                                        case 9223372036854775807:
                                                            v34 := GenS6{F0: 0, F1: 42, F2: 4294967295}
                                                            _ = v34
                                                        default:
                                                            var v35 int32 = -368
                                                            _ = v35
                                                        }
                                                    }
                                                case "":
                                                    var v36 uint32 = 42
                                                    _ = v36
                                                }
                                                v33 = (v33 ^ v33)
                                                _ = v32
                                                _ = v33
                                            }
                                        }
                                    }
                                } else {
                                    v37 := GenS1{F0: 975}
                                    p0 += (^(p0))
                                    var v38 string = ((("" + "\n\t") + "\n\t") + ("foo bar" + "\n\t"))
                                    if (1 <= 1) {
                                        if false {
                                            v39 := GenS1{F0: 948}
                                            _ = v39
                                        }
                                        v40 := GenS7{F0: 0}
                                        _ = v40
                                    } else {
                                        var v41 string = v38
                                        _ = v41
                                    }
                                    _ = v37
                                    _ = v38
                                }
                            } else {
                                fmt.Println(v6)
                                var v42 int16 = 1
                                _ = v42
                            }
                            if false {
                                var v43 string = ""
                                _ = v43
                            } else {
                                _ = len(v12)
                            }
                            if len(v12) > 0 {
                                v12[0] = (fi8 * fi8)
                            }
                            v13 += 0
                            var v44 float32 = v6
                            v45 := GenS1{F0: 2147483647}
                            v13 = v13
                            v44 += v44
                            var v46 uint16 = p0
                            _ = v11
                            _ = v12
                            _ = v13
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        v47 := []int{(fi8 << ((1) & 63)), -2147483648, (fi10 << ((255) & 63)), fi8}
                        var v48 uint32 = 42
                        v49 := GenS8{F0: 127}
                        _ = fi10
                        _ = v47
                        _ = v48
                        _ = v49
                    }
                    v50 := GenS4{F0: 42, F1: 254, F2: 0}
                    _ = v50
                }
                _ = fi8
                _ = v9
            }
            fmt.Println(v6, v6)
        }
    default:
        var v51 string = ((("" + ("hello" + "")) + "\n\t") + "\n\t")
        _ = v51
    }
    _ = p0
    _ = v6
    return 255
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    var v52 uint64 = 0
    if (true || true) {
        if false {
            switch 42 {
            case -792:
                if true {
                    var v53 int16 = genFn0((!false), 57)
                    _ = v53
                }
            case 138:
                v52 += 32
            default:
                v54 := GenS7{F0: 1}
                _ = v54
            }
            var v55 int32 = 42
            _ = v55
        }
        v52 += v52
        fmt.Println(v52)
    } else {
        var v56 int64 = 1
        _ = v56
    }
    v57 := GenS9{F0: 32767}
    var v58 uint16 = 0
    v58 += v58
    var v59 uint32 = 573
    v60 := GenS10{F0: 69, F1: v59}
    v58 -= (v58 * v58)
    var v61 uint64 = v52
    v59 -= 4294967295
    if true {
        if (false && (0 < 42)) {
            v62 := GenS10{F0: v58, F1: (v59 / (((v59 ^ 2)) | 1))}
            var v63 bool = (9223372036854775807 >= -842)
            if true {
                switch "foo bar" {
                case "\n\t":
                    var v64 uint8 = 128
                    _ = v64
                case "hello":
                    v65 := []int{}
                    _ = v65
                default:
                    var v66 int = -9223372036854775808
                    _ = v66
                }
                fmt.Println(v58)
                v52 -= v52
            }
            _ = v62
            _ = v63
        } else {
            fmt.Println(v58, v58)
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v61, v58, v61)
            }
        }
    } else {
        if true {
            for fi67 := 0; fi67 < 14; fi67++ {
                switch ("" + "\n\t") {
                case "\n\t":
                    fmt.Println(v58)
                default:
                    v58 = v58
                }
                v68 := GenS7{F0: 215}
                v58 += v58
                _ = fi67
                _ = v68
            }
            v69 := GenS11{F0: -2147483648, F1: 18446744073709551615}
            var v70 int8 = -1
            _ = v69
            _ = v70
        }
    }
    if (genFn0(false, 1) <= -525) {
        v71 := GenS12{F0: 9223372036854775807, F1: (v52 / ((v61) | 1))}
        for fi72 := 0; fi72 < 6; fi72++ {
            if false {
                for fi73 := 0; fi73 < 9; fi73++ {
                    v74 := []int{}
                    v59 = (v59 | 65535)
                    _ = fi73
                    _ = v74
                }
            } else {
                if (0 == 113) {
                    for fi75 := 0; fi75 < 1; fi75++ {
                        v76 := GenS13{F0: v59}
                        delete(v2, fi72)
                        _ = fi75
                        _ = v76
                    }
                    v77 := []int{}
                    v78 := GenS14{F0: false}
                    _ = v77
                    _ = v78
                }
                for fi79 := 0; fi79 < 17; fi79++ {
                    v58 += v58
                    _ = fi79
                }
                for fi80 := 0; fi80 < 9; fi80++ {
                    if (uint8(v59) <= uint8(v58)) {
                        switch "foo bar" {
                        case "\n\t":
                            switch (("\n\t" + ("a" + "foo bar")) + "a") {
                            case "a":
                                var v81 int8 = 100
                                _ = v81
                            case "\n\t":
                                v82 := GenS6{F0: 1, F1: 127, F2: (^(v61))}
                                _ = v82
                            case "":
                                if false {
                                    var v83 int = (fi72 << ((uint(v61)) & 63))
                                    var v84 float32 = float32((float64(0.0) * (3.14159 + (float64(((3.14159 * 1.5) + 1.5)) + 1e10))))
                                    var v85 uint64 = v52
                                    _ = v83
                                    _ = v84
                                    _ = v85
                                } else {
                                    if false {
                                        v58 -= 729
                                        v86 := GenS15{F0: false}
                                        if false {
                                            v58 += uint16(v59)
                                            var v87 uint32 = v59
                                            _ = v87
                                        } else {
                                            if false {
                                                v88 := GenS11{F0: 42, F1: 0}
                                                v89 := GenS7{F0: 2}
                                                if ((0.1 - float32(float64(1e10))) != 0.0) {
                                                    _ = len(v3)
                                                } else {
                                                    v90 := GenS2{F0: (^(fi72))}
                                                    v1["\n\t"] = "foo bar"
                                                    _ = v90
                                                }
                                                _ = v88
                                                _ = v89
                                            } else {
                                                var v91 uint16 = 0
                                                if true {
                                                    v58 -= (v58 >> ((745) & 63))
                                                    var v92 int32 = 12
                                                    fmt.Println(v52, v58, v61)
                                                    _ = v92
                                                } else {
                                                    var v93 float64 = 1e10
                                                    if false {
                                                        if true {
                                                            v2[666] = true
                                                        }
                                                        v52 -= (v52 ^ v52)
                                                        fmt.Println(v91)
                                                    } else {
                                                        var v94 uint = 534
                                                        var v95 float32 = 0.0
                                                        switch "hello" {
                                                        case "foo bar":
                                                            switch "a" {
                                                            case "foo bar":
                                                                v58 += 156
                                                            case "hello":
                                                                var v96 int32 = 2147483647
                                                                _ = v96
                                                            default:
                                                                v97 := []int{(-(fi72)), fi80}
                                                                _ = v97
                                                            }
                                                        }
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                    if (!true) {
                                                        var v98 uint64 = uint64(v91)
                                                        v99 := GenS16{F0: 1000, F1: 42, F2: 1}
                                                        _ = v98
                                                        _ = v99
                                                    } else {
                                                        v91 += (v91 - (v58 / (((^(v58))) | 1)))
                                                    }
                                                    _ = v93
                                                }
                                                v100 := GenS4{F0: 248, F1: 1, F2: uint(v52)}
                                                _ = v100
                                                _ = v91
                                            }
                                            var v101 int32 = -291
                                            _ = v101
                                        }
                                        _ = v86
                                    } else {
                                        v59 += 587
                                        switch "\n\t" {
                                        case "hello":
                                            if true {
                                                fmt.Println(fi72, v52)
                                            } else {
                                                if true {
                                                    var v102 uint64 = (v61 << ((255) & 63))
                                                    _ = v102
                                                } else {
                                                    v59 = (^(v59))
                                                    v103 := GenS17{F0: 42, F1: v59}
                                                    var v104 uint16 = genFn1(42)
                                                    _ = v103
                                                    _ = v104
                                                }
                                                var v105 int16 = 0
                                                _ = v105
                                            }
                                        case "foo bar":
                                            if (856 > 927) {
                                                if (882 <= v59) {
                                                    var v106 int64 = 0
                                                    v107 := []int{(-(fi72)), -2147483648}
                                                    if len(v107) > 0 {
                                                        v107[0] = (fi72 - (-(fi72)))
                                                    }
                                                    _ = v106
                                                    _ = v107
                                                }
                                                v52 -= 18446744073709551615
                                                var v108 int16 = 32767
                                                _ = v108
                                            } else {
                                                switch "\n\t" {
                                                case "a":
                                                    delete(v2, (-(fi72)))
                                                default:
                                                    v109 := GenS10{F0: 0, F1: (v59 ^ (v59 ^ v59))}
                                                    fmt.Println(v61, fi72)
                                                    v52 = (v61 ^ (v52 ^ (v61 - v61)))
                                                    _ = v109
                                                }
                                            }
                                        case "\n\t":
                                            fmt.Println(v52, v52, v61)
                                        default:
                                            v110 := GenS18{F0: 0}
                                            _ = v110
                                        }
                                        v111 := v2[32767]
                                        _ = v111
                                    }
                                }
                            }
                        case "hello":
                            if ((0.0 + float64((float32((0.0 + (3.14159 - 1e10))) / float32(3.14159)))) != ((3.14159 * 0.0) / 1e10)) {
                                fmt.Println(v59)
                                fmt.Println(v59)
                            } else {
                                v112 := GenS19{F0: 0, F1: 0}
                                v113 := []int{}
                                _ = v112
                                _ = v113
                            }
                        case "":
                            var v114 float64 = -2.25
                            _ = v114
                        }
                        v115 := GenS20{F0: ("a" + "\n\t")}
                        _ = v115
                    } else {
                        var v116 uint8 = 42
                        v117 := GenS19{F0: 42, F1: 0}
                        _ = v116
                        _ = v117
                    }
                    _ = fi80
                }
            }
            for fi118 := 0; fi118 < 11; fi118++ {
                v61 = v52
                _ = fi118
            }
            _ = fi72
        }
        for fi119 := 0; fi119 < 1; fi119++ {
            var v120 int8 = 100
            for fi121 := 0; fi121 < 16; fi121++ {
                v122 := GenS5{F0: v52}
                if true {
                    fmt.Println(v58, fi119)
                    v58 = 814
                    switch ("foo bar" + "") {
                    case "\n\t":
                        if ((0.0 / float32(3.14159)) >= float32(3.14159)) {
                            var v123 bool = true
                            v124 := GenS21{F0: (2 >= v61)}
                            v125 := GenS22{F0: int(v58), F1: (fi119 + (fi119 << ((42) & 63)))}
                            v126 := GenS8{F0: v120}
                            v127 := GenS23{F0: 691, F1: 42}
                            _ = v123
                            _ = v124
                            _ = v125
                            _ = v126
                            _ = v127
                        }
                    case "foo bar":
                        var v128 float64 = 0.0
                        _ = v128
                    }
                }
                switch "" {
                case "a":
                    v59 = 2147483648
                case "hello":
                    if false {
                        switch "" {
                        case "a":
                            var v129 uint = 18446744073709551615
                            _ = v129
                        default:
                            v130 := v3[""]
                            _ = v130
                        }
                        var v131 uint16 = v58
                        _ = v131
                    } else {
                        v120 = v120
                        fmt.Println(v61)
                        if (945 < 1) {
                            v132 := GenS24{F0: 42}
                            fmt.Println(v58)
                            _ = v132
                        } else {
                            var v133 int8 = (-(v120))
                            v134 := []int{(-(fi119)), fi121, (fi121 & int(v133)), 766}
                            _ = v133
                            _ = v134
                        }
                    }
                case "\n\t":
                    v135 := GenS4{F0: -2147483648, F1: 255, F2: 255}
                    _ = v135
                default:
                    v59 += v59
                }
                _ = fi121
                _ = v122
            }
            _ = fi119
            _ = v120
        }
        _ = v71
    } else {
        v1["a"] = (("foo bar" + ("foo bar" + ("foo bar" + ""))) + "foo bar")
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v61
}
