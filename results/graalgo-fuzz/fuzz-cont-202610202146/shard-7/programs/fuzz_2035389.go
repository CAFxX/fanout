package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 int16; F2 string }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 int32; F1 int64; F2 uint16 }
type GenS8 struct { F0 bool; F1 int64; F2 uint8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint8; F1 uint8 }
type GenS11 struct { F0 int32; F1 int32; F2 uint32 }
type GenS12 struct { F0 uint16; F1 int64 }
type GenS13 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 uint32; F1 int }
type GenS16 struct { F0 int64; F1 string; F2 int64 }
type GenS17 struct { F0 int }
type GenS18 struct { F0 uint; F1 int64 }
type GenS19 struct { F0 string; F1 string }
type GenS20 struct { F0 int8; F1 int32; F2 uint64 }
type GenS21 struct { F0 int64; F1 int32; F2 int32 }
func genFn0(p0 int32, p1 int8) uint16 {
    var v7 uint32 = 65535
    v8 := []int{143, -9223372036854775808, 1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 string, p1 int16, p2 int) uint {
    fmt.Println(p0)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 uint16 = 597
    v2 := make(map[uint32]string)
    var v3 int64 = 0
    var v4 string = ("foo bar" + "a")
    v5 := []string{}
    var v6 int64 = (v3 >> ((0) & 63))
    v9 := GenS1{F0: (-(v3)), F1: -1, F2: 9223372036854775807}
    var v10 int8 = -1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v3)
    switch (v4 + "hello") {
    case "a":
        if len(v5) > 0 {
            v5[0] = (v4 + v4)
        }
    case "foo bar":
        for fi11 := 0; fi11 < 19; fi11++ {
            for fi12 := 0; fi12 < 20; fi12++ {
                if ((((float64(((float32(1e10) / 1.0) + 0.0)) / 1.0) - (0.0 * 1e10)) + ((((float64(((1e10 / 1.0) + 0.1)) + (0.1 / 1.0)) - (-2.25 * -2.25)) * float64(-2.25)) - (0.1 / (((3.14159 + 1.5) + float64(float32(1e10))) + 0.1)))) < (float64(0.0) + -2.25)) {
                    v13 := GenS2{F0: 716, F1: 42}
                    var v14 uint16 = (v1 >> ((116) & 63))
                    _ = v13
                    _ = v14
                }
                var v15 uint8 = 164
                v16 := GenS3{F0: (v1 - (^(v1))), F1: (v10 * (v10 >> ((genFn1((v4 + "hello"), 42, (fi11 & (fi12 - (fi11 % ((fi12) | 1)))))) & 63))), F2: 604}
                fmt.Println(v3)
                v1 = (v1 >> ((4294967295) & 63))
                if len(v5) > 0 {
                    v5[0] = ("" + v4)
                }
                switch (("foo bar" + "foo bar") + "\n\t") {
                case "":
                    switch (v4 + "") {
                    case "hello":
                        if ((((^(v1)) < v1) && (fi12 <= 127)) && true) {
                            v4 = "a"
                            v3 -= 1
                            v17 := GenS2{F0: 847, F1: -32768}
                            _ = v17
                        }
                    case "foo bar":
                        switch v4 {
                        case "\n\t":
                            v4 = ("\n\t" + ("a" + v4))
                        case "hello":
                            v5 = append(v5, v4)
                        case "foo bar":
                            switch ("foo bar" + (v4 + ((((v4 + v4) + "a") + "a") + ""))) {
                            case "a":
                                var v18 int64 = v6
                                v19 := v2[65535]
                                if (271 <= -2147483648) {
                                    var v20 int64 = int64(v10)
                                    v20 = (v6 + int64(fi11))
                                    var v21 int = fi11
                                    _ = v20
                                    _ = v21
                                } else {
                                    if ((float64(0.0) + 1.5) < (0.1 - float64(float32(float64(float32((((0.1 * 0.0) + 0.1) - (1.5 / -2.25)))))))) {
                                        v22 := GenS4{F0: 2}
                                        v4 = v19
                                        _ = v22
                                    } else {
                                        fmt.Println(v4)
                                        v5 = append(v5, (v4 + "hello"))
                                    }
                                    v23 := GenS5{F0: (29 < -32768), F1: -1, F2: v19}
                                    fmt.Println(fi11)
                                    _ = v23
                                }
                                _ = v18
                                _ = v19
                            }
                        }
                    default:
                        fmt.Println(v10, v6)
                        v1 -= (v1 + v1)
                    }
                    if len(v5) > 0 {
                        v5[0] = ("" + (((v4 + ("hello" + "\n\t")) + ("foo bar" + ("a" + ("foo bar" + "hello")))) + v4))
                    }
                    if (1 <= 411) {
                        if true {
                            v24 := GenS4{F0: v15}
                            if (1 != 706) {
                                v6 = v6
                                if (4294967295 == 1) {
                                    v25 := GenS6{F0: int8(v1), F1: true}
                                    v1 -= (v1 << ((130) & 63))
                                    var v26 int = (fi11 % ((fi12) | 1))
                                    _ = v25
                                    _ = v26
                                }
                                v27 := GenS7{F0: 0, F1: (v6 - -9223372036854775808), F2: genFn0(-108, (v10 - v10))}
                                _ = v27
                            } else {
                                switch (fi12 % ((-1) | 1)) {
                                case 128:
                                    var v28 uint8 = v15
                                    if (42 != 0) {
                                        var v29 uint = 29
                                        var v30 string = (v4 + "foo bar")
                                        v31 := GenS5{F0: false, F1: 0, F2: ""}
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    } else {
                                        v28 += v15
                                        if (int8(v1) == 100) {
                                            if false {
                                                if true {
                                                    var v32 uint8 = (^(v28))
                                                    _ = v32
                                                }
                                                switch (-(fi11)) {
                                                case 128:
                                                    v10 = v10
                                                case 1:
                                                    v6 += (v6 + int64(v10))
                                                default:
                                                    var v33 uint64 = 1
                                                    _ = v33
                                                }
                                            } else {
                                                v5 = append(v5, (v4 + v4))
                                                v2[42] = v4
                                            }
                                        }
                                        if true {
                                            var v34 float64 = float64(3.14159)
                                            _ = v34
                                        } else {
                                            v35 := GenS5{F0: (782 != 21), F1: -878, F2: v4}
                                            _ = v35
                                        }
                                        if true {
                                            v2[0] = v4
                                            if len(v5) > 0 {
                                                v5[0] = v4
                                            }
                                        }
                                    }
                                    v36 := GenS4{F0: (v28 * 0)}
                                    _ = v28
                                    _ = v36
                                case 901:
                                    var v37 bool = true
                                    _ = v37
                                }
                                v4 = v4
                            }
                            if ((v6 & (v3 + (v3 * (v3 ^ (^(v6)))))) > v6) {
                                delete(v2, 42)
                                _ = len(v2)
                                v4 = v4
                                var v38 string = v4
                                _ = v38
                            }
                            _ = v24
                        }
                        switch 128 {
                        case -952:
                            if len(v5) > 0 {
                                v5[0] = v4
                            }
                        default:
                            v39 := GenS2{F0: 1, F1: -32768}
                            _ = v39
                        }
                    }
                case "\n\t":
                    var v40 uint = 358
                    _ = v40
                case "a":
                    v6 -= v3
                }
                switch 769 {
                case 4294967295:
                    v1 = 2
                case -2147483648:
                    if false {
                        v41 := GenS3{F0: 0, F1: (v10 & v10), F2: uint64(v1)}
                        v42 := v2[1]
                        switch fi11 {
                        case 32767:
                            var v43 uint16 = (v1 % (((v1 << ((0) & 63))) | 1))
                            _ = v43
                        }
                        _ = v41
                        _ = v42
                    }
                default:
                    switch -1 {
                    case -549:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case -540:
                        if (true || false) {
                            if (!true) {
                                v44 := GenS1{F0: 0, F1: 2147483647, F2: fi12}
                                var v45 float64 = 1e10
                                _ = v44
                                _ = v45
                            } else {
                                v46 := GenS4{F0: v15}
                                v15 = (v15 | v15)
                                var v47 uint8 = v15
                                v48 := GenS8{F0: true, F1: 0, F2: 255}
                                _ = v46
                                _ = v47
                                _ = v48
                            }
                            fmt.Println(fi12, v10, v3)
                        }
                    case -9223372036854775808:
                        if ((^(v3)) != (v3 / (((v3 / ((v6) | 1))) | 1))) {
                            v10 -= v10
                            if (42 > 403) {
                                v5 = append(v5, v4)
                                if true {
                                    var v49 uint8 = v15
                                    var v50 uint16 = v1
                                    _ = v49
                                    _ = v50
                                }
                            }
                            v51 := GenS6{F0: v10, F1: (355 > -928)}
                            _ = v51
                        }
                    default:
                        if ((false && (false && (0 < -1))) || false) {
                            v52 := GenS9{F0: 4294967295}
                            var v53 uint32 = 346
                            v54 := GenS10{F0: 1, F1: (v15 >> ((322) & 63))}
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                    }
                }
                _ = fi12
                _ = v15
                _ = v16
            }
            var v55 int64 = v6
            _ = fi11
            _ = v55
        }
    case "":
        var v56 float32 = 0.0
        _ = v56
    }
    v57 := GenS11{F0: 32, F1: -245, F2: 2}
    if true {
        v1 = 351
        var v58 uint = 18446744073709551615
        _ = v58
    } else {
        v10 += (-(v10))
        if ((v6 % (((v3 / ((42) | 1))) | 1)) == v6) {
            v6 = (v3 - (^(v3)))
        }
        fmt.Println(v6, v3, v1)
    }
    if false {
        for fi59 := 0; fi59 < 13; fi59++ {
            for fi60 := 0; fi60 < 8; fi60++ {
                v6 -= v6
                switch fi59 {
                case 44:
                    if false {
                        fmt.Println(v1, v1)
                        switch fi60 {
                        case -9223372036854775808:
                            var v61 uint8 = 84
                            _ = v61
                        }
                    }
                default:
                    if true {
                        var v62 int16 = -155
                        v63 := GenS12{F0: (v1 ^ v1), F1: 1}
                        _ = v62
                        _ = v63
                    } else {
                        var v64 int64 = v3
                        if false {
                            if (392 == 0) {
                                fmt.Println(v6)
                                v65 := GenS13{F0: 434, F1: v6, F2: (v1 / (((v1 - 346)) | 1))}
                                var v66 uint64 = 2
                                _ = v65
                                _ = v66
                            }
                        } else {
                            v6 = 9223372036854775807
                            var v67 uint64 = 42
                            v67 += v67
                            fmt.Println(v64, v10)
                            _ = v67
                        }
                        if false {
                            v64 -= v3
                        }
                        var v68 int8 = v10
                        v69 := GenS14{F0: ""}
                        fmt.Println(v10, v6)
                        v70 := GenS8{F0: (255 < 271), F1: v6, F2: 1}
                        if len(v5) > 0 {
                            v5[0] = v4
                        }
                        if true {
                            v4 = (((((v4 + "\n\t") + (v4 + (v4 + ("hello" + (((v4 + ((("hello" + v4) + v4) + v4)) + v4) + (v4 + ("hello" + v4))))))) + ("foo bar" + v4)) + (v4 + v4)) + (((v4 + (("\n\t" + (("\n\t" + v4) + v4)) + v4)) + v4) + ((("a" + ("\n\t" + v4)) + "") + v4)))
                        } else {
                            var v71 uint = 207
                            v72 := GenS14{F0: (("\n\t" + ((v4 + "foo bar") + "foo bar")) + (v4 + v4))}
                            var v73 uint16 = 1
                            _ = v71
                            _ = v72
                            _ = v73
                        }
                        v74 := v2[507]
                        _ = v64
                        _ = v68
                        _ = v69
                        _ = v70
                        _ = v74
                    }
                    if (-1 > -848) {
                        v10 -= (v10 & v10)
                        v75 := GenS15{F0: 107, F1: (fi59 - (-(fi59)))}
                        _ = v75
                    }
                }
                fmt.Println(v3, v1)
                if (false && true) {
                    switch v4 {
                    case "":
                        if len(v5) > 0 {
                            v5[0] = v4
                        }
                    case "foo bar":
                        var v76 uint32 = 0
                        _ = v76
                    default:
                        v4 = ""
                    }
                    var v77 int32 = 247
                    if false {
                        v78 := GenS16{F0: 1, F1: "", F2: 9223372036854775807}
                        var v79 uint16 = 32767
                        v3 = 400
                        _ = v78
                        _ = v79
                    } else {
                        v80 := v2[2147483647]
                        _ = v80
                    }
                    v10 -= (v10 % ((v10) | 1))
                    _ = v77
                } else {
                    if (true && (-245 < 42)) {
                        var v81 string = (v4 + "foo bar")
                        var v82 float32 = 1e10
                        _ = v81
                        _ = v82
                    }
                    var v83 string = v4
                    _ = v83
                }
                v3 += 914
                switch (("hello" + v4) + v4) {
                case "foo bar":
                    if (false || (fi60 >= fi59)) {
                        var v84 int8 = 100
                        fmt.Println(v84, v6)
                        if (4294967295 == 4294967295) {
                            fmt.Println(v6)
                            fmt.Println(v1, v1)
                        }
                        switch (v4 + ("a" + (v4 + "foo bar"))) {
                        case "\n\t":
                            v85 := GenS13{F0: 645, F1: 42, F2: genFn0(-2147483648, 1)}
                            _ = v85
                        case "hello":
                            if false {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            } else {
                                if true {
                                    v6 += -62
                                    v86 := v2[2]
                                    switch (v4 + v4) {
                                    case "a":
                                        var v87 uint32 = 2
                                        _ = v87
                                    case "\n\t":
                                        if true {
                                            var v88 float64 = 1e10
                                            switch fi60 {
                                            case 127:
                                                if false {
                                                    var v89 uint = 4294967295
                                                    _ = v89
                                                } else {
                                                    if (((!(-1 <= 1000)) || true) || false) {
                                                        var v90 uint32 = 0
                                                        v6 += 42
                                                        var v91 string = "hello"
                                                        _ = v90
                                                        _ = v91
                                                    } else {
                                                        if (("\n\t" + ("hello" + v4)) < v4) {
                                                            v10 = (v10 & v10)
                                                            if false {
                                                                var v92 string = v86
                                                                var v93 float32 = ((1.5 + (1e10 - 0.0)) / (((float32(v88) / (1e10 * 1.5)) * 0.1) + 0.0))
                                                                _ = v92
                                                                _ = v93
                                                            }
                                                        }
                                                        var v94 uint32 = 722
                                                        if false {
                                                            switch -424 {
                                                            case 32768:
                                                                if (!(v88 >= v88)) {
                                                                    var v95 float32 = 0.1
                                                                    if false {
                                                                        v2[(v94 - (v94 >> ((2) & 63)))] = v4
                                                                        switch (fi59 - (fi59 * (fi60 << ((0) & 63)))) {
                                                                        case -1:
                                                                            var v96 int16 = -32768
                                                                            _ = v96
                                                                        case -2147483648:
                                                                            if true {
                                                                                v95 += v95
                                                                                v84 -= (v84 & v84)
                                                                            }
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                            v97 := GenS14{F0: v4}
                                                                            _ = v97
                                                                        }
                                                                        var v98 uint64 = 2
                                                                        if ("\n\t" == v86) {
                                                                            v10 = 0
                                                                        } else {
                                                                            var v99 int8 = (v84 << ((1) & 63))
                                                                            if (2147483648 == v94) {
                                                                                switch int(v6) {
                                                                                case 708:
                                                                                    var v100 uint32 = (v94 + (v94 ^ (v94 / (((v94 << ((65535) & 63))) | 1))))
                                                                                    _ = v100
                                                                                case 1:
                                                                                    switch ((v4 + ("" + "")) + "hello") {
                                                                                    case "hello":
                                                                                        v95 += -2.25
                                                                                    case "":
                                                                                        var v101 int16 = 0
                                                                                        _ = v101
                                                                                    case "a":
                                                                                        fmt.Println(v6, v99, v84)
                                                                                    }
                                                                                    v2[637] = v86
                                                                                    v102 := GenS17{F0: -1}
                                                                                    _ = v102
                                                                                }
                                                                                v103 := GenS10{F0: 128, F1: 127}
                                                                                v104 := GenS7{F0: -2147483648, F1: -203, F2: v1}
                                                                                _ = v103
                                                                                _ = v104
                                                                            }
                                                                            fmt.Println(v84, v88)
                                                                            v105 := GenS18{F0: 0, F1: 0}
                                                                            v2[uint32(v98)] = v86
                                                                            _ = v105
                                                                            _ = v99
                                                                        }
                                                                        var v106 uint = 65535
                                                                        _ = v106
                                                                        _ = v98
                                                                    } else {
                                                                        switch (fi59 + fi59) {
                                                                        case 9223372036854775807:
                                                                            v95 -= (1e10 + (v95 * 3.14159))
                                                                            var v107 bool = (!(true || false))
                                                                            _ = v107
                                                                        case -1:
                                                                            v108 := GenS6{F0: (-(v84)), F1: true}
                                                                            _ = v108
                                                                        default:
                                                                            v84 = (v84 + 1)
                                                                        }
                                                                    }
                                                                    _ = len(v5)
                                                                    fmt.Println(v95, fi60)
                                                                    _ = v95
                                                                } else {
                                                                    v109 := GenS19{F0: "a", F1: v4}
                                                                    _ = v109
                                                                }
                                                            }
                                                            if len(v5) > 0 {
                                                                v5[0] = (v86 + v86)
                                                            }
                                                            var v110 int = (^(fi60))
                                                            _ = v110
                                                        }
                                                        _ = v94
                                                    }
                                                    var v111 bool = (!(false || true))
                                                    _ = v111
                                                }
                                            }
                                            _ = v88
                                        } else {
                                            if (int(v10) <= fi60) {
                                                v5 = append(v5, v86)
                                            } else {
                                                var v112 int64 = (-(v3))
                                                v113 := v2[580]
                                                _ = v112
                                                _ = v113
                                            }
                                            var v114 float32 = 1e10
                                            fmt.Println(v6, v3, v4)
                                            _ = v114
                                        }
                                    }
                                    _ = v86
                                } else {
                                    switch (fi59 ^ fi60) {
                                    case 834:
                                        if false {
                                            switch (v4 + v4) {
                                            case "":
                                                fmt.Println(v1, v6)
                                            case "a":
                                                var v115 uint8 = 5
                                                v84 += (v84 ^ 42)
                                                switch (fi59 - -2147483648) {
                                                case -1:
                                                    var v116 int32 = 16
                                                    _ = v116
                                                }
                                                _ = v115
                                            }
                                            if true {
                                                switch 32768 {
                                                case 32767:
                                                    _ = len(v5)
                                                case -573:
                                                    switch (fi59 - (fi60 << ((63) & 63))) {
                                                    case -9223372036854775808:
                                                        var v117 int16 = 42
                                                        _ = v117
                                                    case 9223372036854775807:
                                                        v118 := GenS15{F0: 892, F1: (fi59 % ((fi59) | 1))}
                                                        _ = v118
                                                    default:
                                                        fmt.Println(v84, v3)
                                                    }
                                                    v119 := GenS20{F0: v84, F1: -2147483648, F2: 334}
                                                    v5 = append(v5, "\n\t")
                                                    _ = v119
                                                default:
                                                    var v120 int64 = 1
                                                    _ = v120
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = (((v4 + ("" + v4)) + v4) + (v4 + "foo bar"))
                                                }
                                            } else {
                                                if false {
                                                    v4 = (v4 + ("a" + v4))
                                                    var v121 int8 = (v84 << ((664) & 63))
                                                    switch 32767 {
                                                    case 223:
                                                        if (!false) {
                                                            fmt.Println(v121, v10)
                                                        }
                                                    case 647:
                                                        fmt.Println(v3)
                                                    case -563:
                                                        if true {
                                                            fmt.Println(v10)
                                                            switch "" {
                                                            case "foo bar":
                                                                v4 = ("hello" + ("hello" + (v4 + v4)))
                                                                fmt.Println(v4, v6)
                                                                v1 = v1
                                                            case "hello":
                                                                var v122 int8 = (v84 + (v10 / ((v84) | 1)))
                                                                _ = v122
                                                            case "a":
                                                                var v123 float32 = float32(0.1)
                                                                _ = v123
                                                            }
                                                            var v124 uint64 = 0
                                                            _ = v124
                                                        }
                                                    }
                                                    _ = v121
                                                }
                                                fmt.Println(v3, fi60)
                                            }
                                            v125 := GenS19{F0: (v4 + (v4 + v4)), F1: v4}
                                            _ = v125
                                        } else {
                                            v4 = (v4 + ((("foo bar" + ("hello" + "a")) + "") + "foo bar"))
                                            fmt.Println(fi60, fi59)
                                        }
                                    }
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v126 int = fi60
                                _ = v126
                            }
                        }
                        _ = v84
                    } else {
                        var v127 int32 = int32(fi60)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v127
                    }
                case "hello":
                    fmt.Println(fi60)
                default:
                    v5 = append(v5, v4)
                }
                _ = fi60
            }
            for fi128 := 0; fi128 < 20; fi128++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v129 int8 = v10
                var v130 int8 = 1
                v6 = v6
                _ = fi128
                _ = v129
                _ = v130
            }
            _ = fi59
        }
        v131 := GenS21{F0: (v6 << ((721) & 63)), F1: -540, F2: 0}
        v1 = (v1 - v1)
        _ = v131
    } else {
        var v132 uint64 = 960
        fmt.Println(v3, v1)
        _ = v132
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v9
}
