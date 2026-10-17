package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int; F1 int; F2 uint }
type GenS4 struct { F0 int64; F1 int32; F2 int }
type GenS5 struct { F0 int32; F1 int16; F2 uint }
type GenS6 struct { F0 bool; F1 int; F2 uint }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 int; F1 uint64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 bool; F1 int32; F2 uint32 }
type GenS11 struct { F0 bool; F1 uint32 }
type GenS12 struct { F0 int32; F1 uint64 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint8; F1 uint16 }
type GenS15 struct { F0 int32; F1 int16 }
type GenS16 struct { F0 uint16; F1 uint8 }
type GenS17 struct { F0 int64; F1 uint; F2 int8 }
func genFn0(p0 string, p1 uint32, p2 bool) int32 {
    v4 := GenS1{F0: 1, F1: 543}
    var v5 int8 = 100
    switch p0 {
    case "\n\t":
        v6 := GenS1{F0: 42, F1: 1}
        _ = v6
    case "":
        v7 := GenS2{F0: 1, F1: (p0 + (p0 + ("a" + (p0 + ""))))}
        _ = v7
    case "hello":
        v8 := GenS2{F0: 4294967295, F1: (((p0 + p0) + "foo bar") + (p0 + ("foo bar" + ("hello" + "a"))))}
        _ = v8
    default:
        p2 = (9223372036854775807 <= 1)
    }
    var v9 uint64 = 2
    v10 := []int{2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v4
    _ = v5
    _ = v9
    return 42
}
func genFn1(p0 uint32) string {
    v11 := GenS3{F0: 32767, F1: -410, F2: 316}
    for fi12 := 0; fi12 < 8; fi12++ {
        if false {
            var v13 int = 1
            _ = v13
        }
        var v14 int32 = 1
        switch fi12 {
        case 712:
            if (2 == 2) {
                v15 := GenS2{F0: fi12, F1: "\n\t"}
                _ = v15
            } else {
                fmt.Println(v14)
                for fi16 := 0; fi16 < 10; fi16++ {
                    v14 += (v14 * (v14 >> ((18446744073709551615) & 63)))
                    v17 := []int{42, (fi12 >> ((42) & 63)), fi16}
                    var v18 float64 = 1.5
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
        default:
            for fi19 := 0; fi19 < 13; fi19++ {
                v20 := GenS2{F0: -1, F1: ("" + (("hello" + "\n\t") + ("" + "foo bar")))}
                if (42 >= 1) {
                    var v21 int16 = -343
                    if true {
                        var v22 uint = 97
                        v23 := []int{}
                        _ = v22
                        _ = v23
                    } else {
                        switch fi19 {
                        case -2147483648:
                            switch "" {
                            case "\n\t":
                                fmt.Println(fi12, v14, fi19)
                                v14 += v14
                                v24 := []int{}
                                _ = v24
                            case "foo bar":
                                v25 := []int{}
                                _ = v25
                            case "a":
                                v21 -= (^(v21))
                            default:
                                if (true && true) {
                                    if false {
                                        fmt.Println(fi19, p0)
                                        p0 = 4294967295
                                        v26 := GenS4{F0: -713, F1: (v14 << ((4294967295) & 63)), F2: fi19}
                                        _ = v26
                                    } else {
                                        v27 := []int{(fi12 - 32768)}
                                        if ((!true) || (-128 < -128)) {
                                            v28 := GenS3{F0: -852, F1: int(p0), F2: 807}
                                            v29 := GenS5{F0: (-(v14)), F1: (v21 << ((65535) & 63)), F2: 42}
                                            if len(v27) >= 2 {
                                                v27 = v27[1:len(v27)]
                                            }
                                            _ = v28
                                            _ = v29
                                        } else {
                                            var v30 bool = true
                                            _ = v30
                                        }
                                        if len(v27) >= 2 {
                                            v27 = v27[1:len(v27)]
                                        }
                                        _ = v27
                                    }
                                    fmt.Println(p0, v21, fi12)
                                }
                            }
                        case 1:
                            v31 := []int{(-(fi19))}
                            _ = v31
                        case 2147483647:
                            v14 = v14
                        }
                    }
                    var v32 bool = true
                    _ = v21
                    _ = v32
                }
                _ = fi19
                _ = v20
            }
        }
        _ = fi12
        _ = v14
    }
    _ = p0
    _ = v11
    return "a"
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    var v3 int64 = 9223372036854775807
    v3 = (^(v3))
    for fi33 := 0; fi33 < 2; fi33++ {
        v34 := []int{-143}
        v3 -= (v3 % ((v3) | 1))
        var v35 int8 = 101
        _ = fi33
        _ = v34
        _ = v35
    }
    v3 = v3
    var v36 float32 = -2.25
    if (2147483647 <= 32767) {
        v3 = 0
        if (255 < 1) {
            var v37 uint64 = 91
            for fi38 := 0; fi38 < 6; fi38++ {
                switch ("a" + "foo bar") {
                case "hello":
                    if true {
                        fmt.Println(v3, v36, v37)
                        switch (^(fi38)) {
                        case 127:
                            switch ("hello" + "foo bar") {
                            case "":
                                v39 := GenS4{F0: v3, F1: 42, F2: fi38}
                                _ = v39
                            case "foo bar":
                                var v40 uint32 = 42
                                _ = v40
                            default:
                                var v41 uint16 = 1
                                _ = v41
                            }
                        }
                    }
                case "a":
                    switch 127 {
                    case 1:
                        v3 = v3
                    case 9223372036854775807:
                        var v42 uint8 = 42
                        _ = v42
                    case 252:
                        switch fi38 {
                        case -172:
                            switch "" {
                            case "a":
                                v3 -= (v3 & (v3 % ((-1) | 1)))
                            case "":
                                var v43 bool = (v3 >= 42)
                                _ = v43
                            default:
                                for fi44 := 0; fi44 < 1; fi44++ {
                                    if (false && true) {
                                        var v45 int64 = v3
                                        fmt.Println(fi44, fi44, v37)
                                        _ = v45
                                    } else {
                                        if (true || true) {
                                            if false {
                                                var v46 int16 = 0
                                                var v47 float64 = 3.14159
                                                _ = v46
                                                _ = v47
                                            } else {
                                                switch 115 {
                                                case -1:
                                                    switch -308 {
                                                    case 32768:
                                                        v3 += (v3 / (((v3 - 9223372036854775807)) | 1))
                                                    }
                                                case 32768:
                                                    var v48 uint64 = (v37 ^ v37)
                                                    _ = v48
                                                case -459:
                                                    fmt.Println(v36)
                                                    var v49 uint = 440
                                                    _ = v49
                                                default:
                                                    v50 := GenS1{F0: v37, F1: v37}
                                                    _ = v50
                                                }
                                                switch "\n\t" {
                                                case "hello":
                                                    if true {
                                                        var v51 int = fi44
                                                        _ = v51
                                                    } else {
                                                        v3 += (^(v3))
                                                    }
                                                case "":
                                                    var v52 uint32 = 65535
                                                    _ = v52
                                                case "\n\t":
                                                    v36 = (((((1.5 + -2.25) * v36) / v36) * ((v36 * (((v36 + (v36 / 3.14159)) + float32((-2.25 * float64(0.1)))) - v36)) - float32(-2.25))) - (float32(((-2.25 * 1e10) / float64((((1.5 + (-2.25 - v36)) - (v36 / ((v36 - v36) + (0.1 + (v36 / (float32(3.14159) - float32(0.0))))))) + v36)))) - v36))
                                                    var v53 uint = 297
                                                    v37 -= 18446744073709551615
                                                    _ = v53
                                                default:
                                                    if (false && (65535 != 802)) {
                                                        var v54 uint = 1
                                                        v55 := GenS3{F0: 42, F1: (fi44 | -866), F2: (^(v54))}
                                                        switch fi38 {
                                                        case -9223372036854775808:
                                                            v56 := GenS5{F0: 309, F1: -1, F2: v54}
                                                            _ = v56
                                                        }
                                                        _ = v54
                                                        _ = v55
                                                    } else {
                                                        var v57 int8 = int8(v3)
                                                        _ = v57
                                                    }
                                                }
                                                if true {
                                                    if (fi38 >= fi38) {
                                                        var v58 int64 = (v3 % ((v3) | 1))
                                                        _ = v58
                                                    }
                                                }
                                            }
                                            switch (^(fi38)) {
                                            case 4294967295:
                                                fmt.Println(v37)
                                            default:
                                                delete(v1, fi38)
                                            }
                                            if false {
                                                v37 = v37
                                                v59 := GenS5{F0: genFn0("", 934, true), F1: -32768, F2: 65535}
                                                if (v3 <= v3) {
                                                    v60 := GenS6{F0: (true && false), F1: (fi44 & fi38), F2: 255}
                                                    _ = v60
                                                }
                                                _ = v59
                                            }
                                        } else {
                                            var v61 int = (fi44 % (((fi38 * fi44)) | 1))
                                            var v62 int = v61
                                            if (v37 == 1) {
                                                v63 := GenS7{F0: fi38, F1: (fi44 ^ v61)}
                                                _ = v63
                                            } else {
                                                var v64 uint16 = 1
                                                if false {
                                                    switch ("" + "foo bar") {
                                                    case "a":
                                                        fmt.Println(v64, v64)
                                                    case "hello":
                                                        v65 := GenS6{F0: true, F1: 9223372036854775807, F2: 42}
                                                        _ = v65
                                                    default:
                                                        switch ("a" + "\n\t") {
                                                        case "a":
                                                            v66 := GenS8{F0: -2147483648, F1: 1}
                                                            _ = v66
                                                        case "":
                                                            v64 += (v64 << ((1) & 63))
                                                        case "\n\t":
                                                            v62 += v61
                                                        default:
                                                            v67 := GenS2{F0: 0, F1: ("a" + "foo bar")}
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v68 float32 = v36
                                                                _ = v68
                                                            case "\n\t":
                                                                if (false || false) {
                                                                    if true {
                                                                        var v69 int16 = 0
                                                                        v61 += v61
                                                                        _ = v69
                                                                    } else {
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            if false {
                                                                                fmt.Println(fi44)
                                                                                var v70 string = ("hello" + (("a" + "hello") + "a"))
                                                                                _ = v70
                                                                            }
                                                                        case "foo bar":
                                                                            fmt.Println(fi38, v36)
                                                                        default:
                                                                            v64 += v64
                                                                        }
                                                                        v37 += (v37 | (v37 ^ 559))
                                                                        v64 = v64
                                                                    }
                                                                    fmt.Println(v36)
                                                                } else {
                                                                    var v71 string = "a"
                                                                    switch ((((("" + v71) + ((("a" + v71) + (v71 + (("" + ("a" + "\n\t")) + ("a" + v71)))) + v71)) + v71) + "foo bar") + v71) {
                                                                    case "":
                                                                        var v72 int16 = 568
                                                                        _ = v72
                                                                    case "hello":
                                                                        fmt.Println(v3)
                                                                    case "\n\t":
                                                                        v37 -= 100
                                                                    default:
                                                                        var v73 int = 404
                                                                        _ = v73
                                                                    }
                                                                    var v74 uint8 = 2
                                                                    _ = v71
                                                                    _ = v74
                                                                }
                                                            case "":
                                                                if ((0 == 1) || true) {
                                                                    if true {
                                                                        _ = len(v2)
                                                                        if true {
                                                                            v64 = 1
                                                                            v75 := GenS9{F0: ""}
                                                                            v37 = 1
                                                                            v76 := v2[2147483647]
                                                                            _ = v75
                                                                            _ = v76
                                                                        } else {
                                                                            v77 := GenS10{F0: (false || true), F1: 251, F2: 671}
                                                                            v64 = 293
                                                                            v37 -= 4294967295
                                                                            v61 = (fi38 + (^(v61)))
                                                                            _ = v77
                                                                        }
                                                                    }
                                                                } else {
                                                                    delete(v1, -47)
                                                                    if (1 == 136) {
                                                                        v61 = 32768
                                                                    }
                                                                    v78 := GenS11{F0: true, F1: 2147483648}
                                                                    v37 += (v37 + 73)
                                                                    _ = v78
                                                                }
                                                            default:
                                                                v79 := GenS8{F0: fi38, F1: v37}
                                                                _ = v79
                                                            }
                                                            v80 := GenS4{F0: 9223372036854775807, F1: 1, F2: (^(fi38))}
                                                            _ = v67
                                                            _ = v80
                                                        }
                                                    }
                                                }
                                                v81 := []int{-9223372036854775808, (fi44 % ((-51) | 1)), (^(fi44)), (fi44 ^ 4294967295)}
                                                _ = v64
                                                _ = v81
                                            }
                                            _ = v61
                                            _ = v62
                                        }
                                        var v82 uint = 4294967295
                                        _ = v82
                                    }
                                    if ((false && (true || (false || false))) || false) {
                                        v83 := []int{(fi38 >> ((2) & 63)), fi38, 0, -386}
                                        var v84 float32 = 1e10
                                        _ = v83
                                        _ = v84
                                    } else {
                                        v37 -= (v37 << ((2) & 63))
                                        var v85 uint32 = 551
                                        _ = v85
                                    }
                                    _ = fi44
                                }
                            }
                        }
                    }
                case "\n\t":
                    var v86 uint8 = 226
                    _ = v86
                default:
                    v37 = v37
                }
                fmt.Println(v36, v37, v37)
                fmt.Println(v37, fi38, v3)
                _ = fi38
            }
            switch "hello" {
            case "a":
                v87 := []int{-325, 42}
                _ = v87
            default:
                v88 := GenS12{F0: -2147483648, F1: v37}
                _ = v88
            }
            _ = v37
        }
        v36 = v36
    } else {
        fmt.Println(v3, v3)
        var v89 bool = true
        var v90 string = ""
        _ = v89
        _ = v90
    }
    fmt.Println(v3)
    for fi91 := 0; fi91 < 5; fi91++ {
        v2[65535] = (!true)
        v36 += v36
        v3 += 1
        _ = fi91
    }
    delete(v2, 0)
    for fi92 := 0; fi92 < 4; fi92++ {
        switch -9223372036854775808 {
        case 0:
            v93 := GenS13{F0: (1 >= 1)}
            _ = v93
        }
        v36 -= v36
        for fi94 := 0; fi94 < 20; fi94++ {
            if (102 < int8(v3)) {
                var v95 string = "foo bar"
                v2[360] = false
                _ = v95
            } else {
                switch (("\n\t" + ("" + "\n\t")) + ("foo bar" + "\n\t")) {
                case "\n\t":
                    if (!true) {
                        switch "a" {
                        case "a":
                            var v96 int8 = 127
                            _ = v96
                        case "foo bar":
                            v36 -= float32(float64(-2.25))
                        case "":
                            var v97 int64 = (v3 ^ v3)
                            _ = v97
                        }
                    } else {
                        var v98 int16 = 1
                        if (1 <= 118) {
                            v99 := GenS5{F0: 42, F1: v98, F2: 0}
                            if ((1 != -1) || false) {
                                var v100 bool = true
                                v36 += v36
                                if true {
                                    var v101 string = ("" + "hello")
                                    v102 := GenS8{F0: fi94, F1: 189}
                                    var v103 uint16 = 42
                                    fmt.Println(v3, v36)
                                    _ = v101
                                    _ = v102
                                    _ = v103
                                } else {
                                    v100 = v100
                                    var v104 bool = v100
                                    v105 := v1[4294967295]
                                    _ = v104
                                    _ = v105
                                }
                                _ = v100
                            }
                            if false {
                                v98 -= -550
                                v98 = v98
                                if false {
                                    fmt.Println(fi94)
                                    switch (((("hello" + "foo bar") + "\n\t") + ("hello" + ("a" + "a"))) + "a") {
                                    case "":
                                        var v106 int8 = -120
                                        _ = v106
                                    default:
                                        v36 -= v36
                                    }
                                    var v107 int = (-(fi92))
                                    _ = v107
                                } else {
                                    if (42 == 65535) {
                                        v108 := GenS2{F0: (fi94 * -83), F1: ""}
                                        v109 := GenS13{F0: ((0.1 * (v36 * v36)) != v36)}
                                        _ = v108
                                        _ = v109
                                    } else {
                                        var v110 int16 = v98
                                        _ = v110
                                    }
                                    v111 := GenS6{F0: false, F1: 300, F2: 171}
                                    fmt.Println(v98, v3, v36)
                                    v98 += (^(v98))
                                    _ = v111
                                }
                            }
                            _ = v99
                        } else {
                            if false {
                                v112 := []int{fi92, fi92}
                                v98 += 495
                                var v113 float32 = v36
                                _ = v112
                                _ = v113
                            }
                        }
                        switch "hello" {
                        case "":
                            v114 := []int{}
                            _ = v114
                        default:
                            var v115 uint = 792
                            _ = v115
                        }
                        _ = v98
                    }
                case "":
                    if (!(1 > 1)) {
                        if true {
                            v36 += (v36 + (v36 + 3.14159))
                            v3 = (v3 + 0)
                            v3 = (v3 ^ -1)
                        }
                        fmt.Println(v3, v3, fi92)
                        var v116 int32 = genFn0((("" + ("hello" + "foo bar")) + "\n\t"), 65535, (0 >= 0))
                        _ = v116
                    } else {
                        switch "hello" {
                        case "a":
                            v3 = (v3 * -9223372036854775808)
                        default:
                            v3 += (v3 | (v3 ^ v3))
                        }
                    }
                }
                var v117 int64 = (v3 >> ((535) & 63))
                _ = v117
            }
            v36 -= (v36 - v36)
            var v118 float64 = 0.1
            _ = fi94
            _ = v118
        }
        _ = fi92
    }
    var v119 string = "foo bar"
    var v120 uint8 = 99
    var v121 uint16 = 1
    for fi122 := 0; fi122 < 1; fi122++ {
        switch fi122 {
        case 1:
            v123 := GenS14{F0: 1, F1: (v121 / (((v121 ^ (v121 - 32767))) | 1))}
            _ = v123
        case 127:
            v124 := GenS5{F0: 2147483647, F1: 0, F2: 255}
            _ = v124
        case -477:
            v36 = (v36 + (1.5 - v36))
        default:
            v125 := GenS15{F0: -153, F1: 0}
            _ = v125
        }
        fmt.Println(v3, v121, fi122)
        _ = fi122
    }
    v126 := GenS5{F0: 2147483647, F1: 1, F2: 18446744073709551615}
    var v127 bool = (1 == 0)
    var v128 uint8 = v120
    v121 = v121
    if v127 {
        v129 := GenS16{F0: 116, F1: v128}
        v130 := GenS17{F0: v3, F1: 42, F2: -128}
        _ = v129
        _ = v130
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("v127:", v127)
    fmt.Println("v128:", v128)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v119
    _ = v120
    _ = v121
    _ = v126
    _ = v127
    _ = v128
    _ = v2
    _ = v3
    _ = v36
}
