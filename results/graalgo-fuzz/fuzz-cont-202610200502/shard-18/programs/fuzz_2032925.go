package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 string; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 int64; F1 bool; F2 uint32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 string; F1 uint8 }
type GenS12 struct { F0 int32; F1 int32 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int64; F1 int16; F2 uint32 }
type GenS15 struct { F0 int8; F1 uint32 }
type GenS16 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS17 struct { F0 int32; F1 uint32; F2 int8 }
type GenS18 struct { F0 int32; F1 uint64; F2 int16 }
type GenS19 struct { F0 int32 }
func genFn0(p0 uint8) int {
    v5 := GenS1{F0: ((509 != 2) || true), F1: "foo bar"}
    if (3.14159 <= 3.14159) {
        fmt.Println(p0)
    }
    p0 -= (p0 & (p0 ^ p0))
    p0 = p0
    _ = p0
    _ = v5
    return -160
}
func genFn1(p0 string, p1 uint16) uint32 {
    v6 := GenS2{F0: (("\n\t" + p0) + ("foo bar" + "a")), F1: 507, F2: 4294967295}
    fmt.Println(p1)
    v7 := []int{951, genFn0(149), 128, 2147483647}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 739
}

func main() {
    var v1 int16 = -435
    var v2 uint32 = 2147483647
    var v3 float32 = 1e10
    var v4 uint64 = 353
    if false {
        var v8 int8 = 82
        for fi9 := 0; fi9 < 16; fi9++ {
            fmt.Println(v1, fi9, v4)
            v10 := []int{fi9, 665, fi9, int(v4)}
            _ = fi9
            _ = v10
        }
        var v11 int8 = v8
        _ = v11
        _ = v8
    }
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        v12 := GenS3{F0: 2}
        _ = v12
    }
    switch "foo bar" {
    case "":
        v1 -= v1
    case "foo bar":
        v13 := []int{}
        _ = v13
    case "hello":
        v14 := []int{1}
        _ = v14
    }
    switch -9223372036854775808 {
    case -2147483648:
        if (float64(v3) < (-2.25 / 1.0)) {
            if (v4 != v4) {
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := GenS3{F0: 0}
                    v17 := GenS4{F0: v2}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            } else {
                v18 := []int{}
                if false {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        v20 := GenS3{F0: 2}
                        _ = fi19
                        _ = v20
                    }
                    fmt.Println(v3)
                    var v21 string = ""
                    _ = v21
                }
                _ = v18
            }
            switch "\n\t" {
            case "a":
                for fi22 := 0; fi22 < 14; fi22++ {
                    v4 = (^(v4))
                    if false {
                        for fi23 := 0; fi23 < 15; fi23++ {
                            switch genFn0(130) {
                            case 9223372036854775807:
                                v24 := GenS5{F0: v1, F1: (("hello" + "a") + "hello")}
                                _ = v24
                            }
                            _ = fi23
                        }
                        for fi25 := 0; fi25 < 6; fi25++ {
                            if (!(!true)) {
                                switch "\n\t" {
                                case "a":
                                    v1 += (v1 << ((686) & 63))
                                case "":
                                    v2 += uint32(v4)
                                    if (false && true) {
                                        if (true || true) {
                                            fmt.Println(v3, v3)
                                            fmt.Println(v4)
                                            switch 0 {
                                            case 42:
                                                switch -2147483648 {
                                                case 127:
                                                    v1 = v1
                                                    v4 += v4
                                                    v26 := GenS5{F0: (-(v1)), F1: "a"}
                                                    if true {
                                                        var v27 float32 = (((v3 + (-2.25 + float32((float64((float32((float64((v3 + v3)) * 0.0)) - (float32(1.5) + (float32((0.1 - (0.1 * ((1.5 * 1.5) * 3.14159)))) / 1.5)))) + -2.25)))) + (1.5 - v3)) - v3)
                                                        v4 += 18446744073709551615
                                                        _ = v27
                                                    }
                                                    if false {
                                                        var v28 uint32 = 42
                                                        var v29 float32 = 3.14159
                                                        var v30 int16 = (v1 | v1)
                                                        v3 -= float32(0.0)
                                                        v2 += (v2 * v28)
                                                        _ = v28
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                    fmt.Println(fi25)
                                                    v31 := GenS6{F0: -894, F1: true, F2: (v2 + (v2 >> ((623) & 63)))}
                                                    _ = v26
                                                    _ = v31
                                                case 32767:
                                                    v4 += v4
                                                    if (false && false) {
                                                        var v32 uint16 = 65535
                                                        _ = v32
                                                    } else {
                                                        v33 := []int{-2147483648}
                                                        v33 = append(v33, fi22)
                                                        v2 = v2
                                                        _ = v33
                                                    }
                                                case 4294967295:
                                                    var v34 int32 = 958
                                                    _ = v34
                                                default:
                                                    v1 -= v1
                                                }
                                            case -882:
                                                if (65535 == 2) {
                                                    var v35 string = "hello"
                                                    var v36 uint64 = v4
                                                    switch (-(fi22)) {
                                                    case 127:
                                                        if (false && false) {
                                                            var v37 uint = 0
                                                            _ = v37
                                                        } else {
                                                            v2 = genFn1(v35, 255)
                                                            if (206 != 0) {
                                                                var v38 float64 = (-2.25 / 0.1)
                                                                v38 -= (1.5 * v38)
                                                                if false {
                                                                    if (-1 == 1) {
                                                                        if false {
                                                                            if ((v38 / v38) == 0.0) {
                                                                                switch (fi22 ^ int(v2)) {
                                                                                case 9223372036854775807:
                                                                                    var v39 uint16 = 255
                                                                                    _ = v39
                                                                                case 128:
                                                                                    var v40 float32 = float32(1e10)
                                                                                    _ = v40
                                                                                case -532:
                                                                                    switch v35 {
                                                                                    case "\n\t":
                                                                                        if false {
                                                                                            v41 := []int{}
                                                                                            _ = v41
                                                                                        }
                                                                                        switch v35 {
                                                                                        case "\n\t":
                                                                                            var v42 int64 = 810
                                                                                            _ = v42
                                                                                        case "foo bar":
                                                                                            v43 := GenS3{F0: 128}
                                                                                            _ = v43
                                                                                        case "":
                                                                                            var v44 float32 = 1e10
                                                                                            _ = v44
                                                                                        }
                                                                                    case "foo bar":
                                                                                        if true {
                                                                                            v45 := GenS5{F0: (v1 - 42), F1: (v35 + (("foo bar" + v35) + "foo bar"))}
                                                                                            fmt.Println(v38, v38)
                                                                                            var v46 int = fi22
                                                                                            _ = v45
                                                                                            _ = v46
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v2, v36)
                                                                                    }
                                                                                }
                                                                                v47 := []int{(fi25 + (fi22 * genFn0(2)))}
                                                                                var v48 int8 = 0
                                                                                _ = v47
                                                                                _ = v48
                                                                            }
                                                                        }
                                                                        v35 = (v35 + v35)
                                                                        v35 = v35
                                                                    } else {
                                                                        v3 = (1e10 + ((float32(1e10) / 1.0) + (-2.25 - v3)))
                                                                        var v49 uint64 = v4
                                                                        v3 = 1e10
                                                                        _ = v49
                                                                    }
                                                                    var v50 float32 = v3
                                                                    v3 -= float32(0.0)
                                                                    switch v35 {
                                                                    case "foo bar":
                                                                        v35 = "\n\t"
                                                                    case "\n\t":
                                                                        v51 := GenS7{F0: (fi25 + (fi25 >> ((966) & 63)))}
                                                                        _ = v51
                                                                    case "hello":
                                                                        if true {
                                                                            var v52 int8 = 127
                                                                            v53 := []int{(fi25 + 2147483647), int(v52), -897}
                                                                            fmt.Println(v1, v2)
                                                                            _ = v52
                                                                            _ = v53
                                                                        } else {
                                                                            v54 := GenS2{F0: v35, F1: v36, F2: (v4 | v4)}
                                                                            v1 -= v1
                                                                            _ = v54
                                                                        }
                                                                    }
                                                                    _ = v50
                                                                }
                                                                var v55 int64 = -9223372036854775808
                                                                _ = v38
                                                                _ = v55
                                                            }
                                                        }
                                                    case 4294967295:
                                                        switch v35 {
                                                        case "hello":
                                                            switch fi25 {
                                                            case 4294967295:
                                                                v56 := []int{(-(fi25))}
                                                                _ = v56
                                                            }
                                                            v57 := GenS8{F0: 4294967295}
                                                            var v58 uint = 2
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                    case 331:
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            if ((((("hello" + ("foo bar" + v35)) + (v35 + v35)) + v35) + ((v35 + v35) + v35)) != ("" + "hello")) {
                                                                v59 := []int{(fi25 & (-(fi22))), (fi22 % ((genFn0(127)) | 1))}
                                                                _ = v59
                                                            } else {
                                                                var v60 int32 = -1
                                                                v36 = v36
                                                                v61 := GenS9{F0: 1}
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                        case "a":
                                                            if (1 != 999) {
                                                                var v62 int64 = 42
                                                                _ = v62
                                                            } else {
                                                                v2 -= (v2 | uint32(v4))
                                                            }
                                                        default:
                                                            if ((((((1.5 - float32(-2.25)) + float32((0.1 - (1.5 - -2.25)))) + (v3 / -2.25)) * v3) / (float32(((1.5 * float64(v3)) / float64(-2.25))) - v3)) < v3) {
                                                                fmt.Println(fi25, v3)
                                                                var v63 int8 = 127
                                                                _ = v63
                                                            }
                                                        }
                                                    default:
                                                        var v64 string = v35
                                                        _ = v64
                                                    }
                                                    _ = v35
                                                    _ = v36
                                                } else {
                                                    v65 := []int{(^(fi25))}
                                                    v1 = (v1 - 1000)
                                                    _ = v65
                                                }
                                            default:
                                                if true {
                                                    var v66 int8 = 65
                                                    _ = v66
                                                } else {
                                                    switch ("foo bar" + ("" + "foo bar")) {
                                                    case "a":
                                                        v2 = v2
                                                    case "foo bar":
                                                        switch "" {
                                                        case "":
                                                            var v67 bool = (!((42 > 0) && (!false)))
                                                            _ = v67
                                                        default:
                                                            v68 := []int{}
                                                            _ = v68
                                                        }
                                                    default:
                                                        v2 += genFn1(("foo bar" + "\n\t"), 0)
                                                    }
                                                    v69 := []int{-2147483648, -9223372036854775808, 9223372036854775807, (fi22 * (-(fi22)))}
                                                    if len(v69) >= 2 {
                                                        v69 = v69[1:len(v69)]
                                                    }
                                                    _ = v69
                                                }
                                                var v70 bool = false
                                                _ = v70
                                            }
                                        } else {
                                            v71 := GenS10{F0: false}
                                            var v72 float32 = v3
                                            v4 += v4
                                            _ = v71
                                            _ = v72
                                        }
                                        v3 = (3.14159 / (v3 * v3))
                                        if true {
                                            v73 := GenS11{F0: "a", F1: 79}
                                            _ = v73
                                        }
                                    }
                                    var v74 uint8 = 1
                                    _ = v74
                                case "\n\t":
                                    switch genFn0(1) {
                                    case 0:
                                        v2 -= 70
                                    case -1:
                                        var v75 uint = 65535
                                        _ = v75
                                    case 32767:
                                        v76 := GenS12{F0: -1, F1: 2147483647}
                                        if (!(0 > 2147483647)) {
                                            v3 -= 0.1
                                            v3 -= v3
                                            v77 := GenS10{F0: false}
                                            _ = v77
                                        } else {
                                            var v78 uint16 = 42
                                            var v79 int16 = v1
                                            _ = v78
                                            _ = v79
                                        }
                                        v80 := GenS13{F0: 1}
                                        _ = v76
                                        _ = v80
                                    }
                                default:
                                    v81 := []int{fi25, fi25, (^(fi25))}
                                    _ = v81
                                }
                                switch (fi25 ^ genFn0(255)) {
                                case -2147483648:
                                    v82 := []int{(^(fi22)), genFn0(1)}
                                    _ = v82
                                case -897:
                                    v1 += v1
                                case 4294967295:
                                    var v83 uint8 = 1
                                    _ = v83
                                default:
                                    if (42 < 71) {
                                        fmt.Println(v2)
                                        v2 += v2
                                        v1 -= (v1 - (v1 * -166))
                                    }
                                }
                                var v84 int16 = int16(fi25)
                                _ = v84
                            } else {
                                v85 := []int{}
                                var v86 uint16 = 2
                                v1 += (v1 ^ -1)
                                v2 += v2
                                var v87 int64 = -1
                                _ = v85
                                _ = v86
                                _ = v87
                            }
                            _ = fi25
                        }
                        v88 := []int{(fi22 + fi22)}
                        _ = v88
                    } else {
                        var v89 string = "a"
                        v90 := []int{}
                        _ = v89
                        _ = v90
                    }
                    _ = fi22
                }
            case "":
                v2 = (v2 + v2)
            default:
                if (42 > 497) {
                    v3 = (0.1 * 1e10)
                    for fi91 := 0; fi91 < 4; fi91++ {
                        v92 := GenS14{F0: -771, F1: 32767, F2: v2}
                        if false {
                            if false {
                                switch fi91 {
                                case -70:
                                    v4 += (v4 % ((372) | 1))
                                case -2147483648:
                                    v2 = genFn1("\n\t", uint16(v4))
                                case 2147483647:
                                    v93 := GenS14{F0: -1, F1: (v1 << ((4294967295) & 63)), F2: 2}
                                    _ = v93
                                default:
                                    v94 := []int{fi91, fi91, fi91}
                                    _ = v94
                                }
                                v95 := GenS15{F0: 0, F1: genFn1("hello", 255)}
                                v2 += (v2 ^ 2)
                                _ = v95
                            }
                            fmt.Println(fi91, v1)
                            switch "a" {
                            case "hello":
                                fmt.Println(fi91)
                            case "":
                                v96 := GenS4{F0: (^(v2))}
                                _ = v96
                            }
                        }
                        v1 -= 489
                        _ = fi91
                        _ = v92
                    }
                    v97 := GenS16{F0: 48, F1: 127, F2: 2}
                    _ = v97
                }
            }
            v98 := GenS17{F0: 0, F1: (v2 & v2), F2: -52}
            _ = v98
        }
    }
    var v99 int32 = 0
    v4 = v4
    v100 := GenS2{F0: (("foo bar" + ("foo bar" + "\n\t")) + ""), F1: (v4 + v4), F2: 42}
    switch "" {
    case "\n\t":
        for fi101 := 0; fi101 < 17; fi101++ {
            v2 = genFn1("a", 2)
            _ = fi101
        }
    case "foo bar":
        var v102 uint16 = 509
        _ = v102
    case "a":
        if (false || false) {
            if true {
                for fi103 := 0; fi103 < 8; fi103++ {
                    v3 = (v3 * 1.5)
                    _ = fi103
                }
                v104 := GenS18{F0: 2147483647, F1: (v4 * (v4 - uint64(v2))), F2: v1}
                for fi105 := 0; fi105 < 18; fi105++ {
                    var v106 uint8 = 50
                    var v107 int8 = 42
                    var v108 int16 = 0
                    _ = fi105
                    _ = v106
                    _ = v107
                    _ = v108
                }
                _ = v104
            }
            v99 = v99
            v109 := GenS13{F0: -479}
            _ = v109
        } else {
            var v110 bool = true
            v110 = v110
            if (v110 || v110) {
                v111 := []int{-281}
                if len(v111) > 0 {
                    v111[0] = -9223372036854775808
                }
                switch genFn0(222) {
                case -1:
                    if false {
                        var v112 int64 = -1
                        var v113 uint32 = (v2 & (v2 + genFn1("", 895)))
                        _ = v112
                        _ = v113
                    }
                }
                _ = v111
            }
            _ = v110
        }
    default:
        for fi114 := 0; fi114 < 3; fi114++ {
            v115 := []int{(-(fi114)), (fi114 / ((int(v2)) | 1)), fi114, 417}
            var v116 float64 = 3.14159
            _ = fi114
            _ = v115
            _ = v116
        }
    }
    if true {
        fmt.Println(v99, v3, v3)
    }
    v2 = genFn1("foo bar", 0)
    v117 := GenS19{F0: v99}
    v118 := GenS13{F0: 9223372036854775807}
    var v119 int8 = 100
    for fi120 := 0; fi120 < 15; fi120++ {
        fmt.Println(v2)
        _ = fi120
    }
    var v121 int64 = 42
    v122 := []int{128, -131, -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("v121:", v121)
    fmt.Println("slicelen v122:", len(v122), v122)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v117
    _ = v118
    _ = v119
    _ = v121
    _ = v122
    _ = v2
    _ = v3
    _ = v4
    _ = v99
}
