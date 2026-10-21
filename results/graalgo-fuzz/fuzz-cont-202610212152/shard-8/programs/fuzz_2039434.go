package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS3 struct { F0 int8; F1 uint8; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 int; F2 uint }
type GenS7 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS8 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS9 struct { F0 int64; F1 int8 }
type GenS10 struct { F0 int8; F1 uint8 }
type GenS11 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS12 struct { F0 bool; F1 bool }
type GenS13 struct { F0 string; F1 string }
type GenS14 struct { F0 uint64; F1 string; F2 uint8 }
type GenS15 struct { F0 string; F1 string; F2 uint8 }
type GenS16 struct { F0 uint64; F1 int8 }
type GenS17 struct { F0 uint32; F1 int }
type GenS18 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 int64 }
type GenS21 struct { F0 uint64; F1 int8 }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 int16 }
type GenS24 struct { F0 uint }
type GenS25 struct { F0 int64 }
type GenS26 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint16, p2 uint16) uint16 {
    p1 -= 42
    fmt.Println(p2)
    if ((("" + ("foo bar" + "foo bar")) + "\n\t") <= "hello") {
        switch "" {
        case "\n\t":
            var v5 string = "hello"
            _ = v5
        default:
            p1 -= 2
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}
func genFn1(p0 uint64, p1 string, p2 uint32) uint64 {
    p0 += 1
    if (0 < -2147483648) {
        var v6 uint16 = 1
        _ = v6
    }
    p2 = p2
    p1 = "hello"
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 2
    var v3 int8 = -67
    v4 := make(map[uint32]bool)
    v7 := GenS1{F0: 1000, F1: 4294967295}
    v2 = v2
    if false {
        v8 := GenS1{F0: -1, F1: 255}
        _ = v8
    } else {
        var v9 uint64 = 876
        delete(v4, 1)
        _ = v9
    }
    if (480 > 1) {
        for fi10 := 0; fi10 < 20; fi10++ {
            v2 += (v2 >> ((0) & 63))
            _ = fi10
        }
    }
    var v11 int32 = 2147483647
    var v12 uint16 = 255
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := GenS2{F0: 42, F1: 0, F2: 801}
        if false {
            var v15 bool = true
            _ = v15
        } else {
            if (!false) {
                if true {
                    var v16 uint = 42
                    if false {
                        for fi17 := 0; fi17 < 9; fi17++ {
                            if false {
                                if true {
                                    v12 -= (v12 ^ (v12 + genFn0(2, v12, (v12 & genFn0(127, v12, (v12 * 32767))))))
                                }
                            } else {
                                v12 += 255
                                var v18 int16 = -1
                                var v19 uint = 1
                                var v20 uint = v16
                                var v21 float64 = ((3.14159 / ((((0.0 / 1.5) + (float64(((1e10 + 3.14159) / -2.25)) * ((-2.25 - (3.14159 * 0.0)) * ((float64((((1.5 / 3.14159) / (float32(1e10) * float32(float64(3.14159)))) / 1e10)) - 3.14159) * (1e10 / float64(-2.25)))))) * 0.0) - (float64(((-2.25 + ((((1.5 + 1e10) * ((float32((0.0 * (0.0 - 0.0))) + 0.0) * 0.1)) / (1.5 + 0.0)) / float32(0.1))) + (float32((-2.25 - -2.25)) + 1.5))) * 1.5))) * 1.5)
                                v11 += (v11 << ((0) & 63))
                                switch ((("\n\t" + "a") + "a") + "hello") {
                                case "foo bar":
                                    var v22 uint32 = 288
                                    _ = v22
                                case "\n\t":
                                    var v23 uint16 = v12
                                    if true {
                                        _ = len(v4)
                                        var v24 bool = false
                                        _ = v24
                                    } else {
                                        var v25 uint16 = 255
                                        fmt.Println(v25)
                                        v26 := []int{9223372036854775807}
                                        _ = v25
                                        _ = v26
                                    }
                                    v27 := []int{(fi13 + 128)}
                                    _ = v23
                                    _ = v27
                                case "":
                                    if (!false) {
                                        if false {
                                            var v28 int = fi13
                                            v29 := GenS1{F0: 274, F1: v19}
                                            _ = v28
                                            _ = v29
                                        } else {
                                            var v30 uint8 = 169
                                            if ((!false) && (!(true || ("" == "hello")))) {
                                                v31 := GenS3{F0: v3, F1: v30, F2: 128}
                                                v32 := v4[4294967295]
                                                if (1e10 < 3.14159) {
                                                    if ((fi17 >> ((uint(v12)) & 63)) > (^(fi17))) {
                                                        v33 := GenS3{F0: (-(v3)), F1: (v30 & v30), F2: fi13}
                                                        v30 -= 103
                                                        delete(v1, fi13)
                                                        v18 += (v18 & (v18 + v18))
                                                        _ = v33
                                                    } else {
                                                        fmt.Println(fi13)
                                                        var v34 int64 = 1
                                                        v35 := GenS1{F0: v18, F1: v20}
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                    v36 := []int{fi13, -840}
                                                    delete(v1, -281)
                                                    _ = v36
                                                }
                                                _ = v31
                                                _ = v32
                                            } else {
                                                if true {
                                                    v37 := GenS4{F0: 496}
                                                    _ = v37
                                                } else {
                                                    v11 = v11
                                                    switch fi17 {
                                                    case 2147483647:
                                                        var v38 int32 = (^(v11))
                                                        _ = v38
                                                    case 4294967295:
                                                        var v39 int32 = (v11 % ((v11) | 1))
                                                        v40 := GenS5{F0: 634}
                                                        _ = v39
                                                        _ = v40
                                                    case 128:
                                                        var v41 uint16 = v12
                                                        var v42 float64 = v21
                                                        v3 -= int8(v12)
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    switch (fi17 % (((fi17 % (((-(fi17))) | 1))) | 1)) {
                                                    case 437:
                                                        v4[4294967295] = (int16(v20) == (v18 * (^(v18))))
                                                    case -705:
                                                        fmt.Println(fi13)
                                                    default:
                                                        switch fi13 {
                                                        case 9223372036854775807:
                                                            var v43 string = ""
                                                            _ = v43
                                                        }
                                                    }
                                                }
                                                v44 := GenS5{F0: (v18 >> ((v19) & 63))}
                                                var v45 float32 = 1.5
                                                switch fi13 {
                                                case 2147483647:
                                                    var v46 float64 = v21
                                                    _ = v46
                                                case 351:
                                                    if true {
                                                        if true {
                                                            var v47 string = "hello"
                                                            v3 = v3
                                                            _ = v47
                                                        }
                                                        var v48 int64 = 1
                                                        var v49 uint32 = 42
                                                        _ = v48
                                                        _ = v49
                                                    } else {
                                                        v50 := GenS6{F0: (v11 + (v11 | v11)), F1: fi13, F2: v20}
                                                        _ = v50
                                                    }
                                                case 127:
                                                    v11 = (v11 ^ int32(fi13))
                                                default:
                                                    fmt.Println(v30, v12, v30)
                                                }
                                                if (genFn1(genFn1((v2 * (^(v2))), ("" + ("foo bar" + "hello")), 0), ((("a" + ("hello" + "a")) + "") + (("hello" + "hello") + "foo bar")), 363) >= 0) {
                                                    v16 -= v16
                                                    if (v45 != v45) {
                                                        fmt.Println(v18)
                                                    } else {
                                                        v51 := GenS7{F0: (^(v30)), F1: v3, F2: 18446744073709551615}
                                                        var v52 int16 = 12
                                                        v53 := GenS7{F0: v30, F1: (v3 + (v3 << ((v20) & 63))), F2: uint64(v16)}
                                                        _ = v51
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                } else {
                                                    var v54 uint8 = 86
                                                    _ = v54
                                                }
                                                var v55 int16 = -1
                                                var v56 int = int(v30)
                                                v12 += v12
                                                if ((^(v3)) > (v3 ^ (-(v3)))) {
                                                    switch 1 {
                                                    case 42:
                                                        var v57 float32 = (v45 + (v45 * v45))
                                                        _ = v57
                                                    case 1:
                                                        v58 := GenS6{F0: v11, F1: (v56 + fi13), F2: 471}
                                                        _ = v58
                                                    case 2147483647:
                                                        v59 := []int{(v56 >> ((v19) & 63)), (v56 >> ((65535) & 63)), (v56 ^ 807)}
                                                        _ = v59
                                                    default:
                                                        if true {
                                                            v60 := GenS8{F0: 65535, F1: v2, F2: 0}
                                                            if false {
                                                                fmt.Println(v56)
                                                                var v61 float64 = v21
                                                                _ = v61
                                                            }
                                                            if true {
                                                                switch "a" {
                                                                case "a":
                                                                    v20 += v20
                                                                case "\n\t":
                                                                    if true {
                                                                        var v62 int16 = 42
                                                                        v12 += (^(v12))
                                                                        var v63 uint = v16
                                                                        _ = v62
                                                                        _ = v63
                                                                    }
                                                                }
                                                                switch int(v30) {
                                                                case 444:
                                                                    var v64 int8 = 42
                                                                    _ = v64
                                                                case -420:
                                                                    if true {
                                                                        v56 = -2147483648
                                                                    } else {
                                                                        fmt.Println(v16)
                                                                    }
                                                                case 247:
                                                                    v65 := []int{fi13}
                                                                    _ = v65
                                                                default:
                                                                    var v66 float64 = v21
                                                                    _ = v66
                                                                }
                                                            }
                                                            _ = v60
                                                        }
                                                    }
                                                }
                                                _ = v44
                                                _ = v45
                                                _ = v55
                                                _ = v56
                                            }
                                            v19 = (^(v20))
                                            _ = v30
                                        }
                                    } else {
                                        var v67 uint64 = 28
                                        v12 = v12
                                        _ = v67
                                    }
                                    switch ("" + "foo bar") {
                                    case "foo bar":
                                        v11 = (-(v11))
                                    case "hello":
                                        v68 := GenS9{F0: 42, F1: 99}
                                        _ = v68
                                    default:
                                        v69 := []int{127, fi13, -995, (fi17 % ((-1) | 1))}
                                        _ = v69
                                    }
                                    if false {
                                        v16 = (v16 | v20)
                                        v70 := GenS7{F0: 1, F1: v3, F2: (v2 / ((v2) | 1))}
                                        _ = v70
                                    } else {
                                        if (2147483648 < 65535) {
                                            var v71 uint16 = (v12 | uint16(v19))
                                            if true {
                                                v4[2147483648] = true
                                            }
                                            _ = v71
                                        } else {
                                            fmt.Println(v12)
                                            fmt.Println(v19, v20, fi17)
                                        }
                                        if false {
                                            fmt.Println(fi13, fi13, v18)
                                            var v72 int8 = v3
                                            v18 += -32768
                                            _ = v72
                                        } else {
                                            v73 := []int{}
                                            _ = v73
                                        }
                                    }
                                    var v74 int16 = (v18 % (((^(v18))) | 1))
                                    _ = v74
                                default:
                                    switch "hello" {
                                    case "foo bar":
                                        v18 -= v18
                                        if false {
                                            if (true && true) {
                                                switch "foo bar" {
                                                case "hello":
                                                    if false {
                                                        v75 := v4[65535]
                                                        switch ("" + "") {
                                                        case "hello":
                                                            switch (fi13 + 32767) {
                                                            case 172:
                                                                v12 = 32767
                                                            case -1:
                                                                v2 = 1
                                                            case 42:
                                                                switch fi17 {
                                                                case 720:
                                                                    v76 := GenS10{F0: (v3 - (^(v3))), F1: 0}
                                                                    _ = v76
                                                                case -621:
                                                                    var v77 float64 = v21
                                                                    _ = v77
                                                                case 127:
                                                                    var v78 int = 32767
                                                                    _ = v78
                                                                default:
                                                                    v79 := []int{fi17}
                                                                    _ = v79
                                                                }
                                                            default:
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    v80 := GenS6{F0: (v11 << ((836) & 63)), F1: 32767, F2: (v20 >> ((326) & 63))}
                                                                    _ = v80
                                                                case "a":
                                                                    var v81 int = (fi13 & 32768)
                                                                    _ = v81
                                                                case "foo bar":
                                                                    if v75 {
                                                                        var v82 int64 = 0
                                                                        v19 = 65535
                                                                        switch (fi17 & (fi17 | 127)) {
                                                                        case -853:
                                                                            var v83 uint8 = 34
                                                                            _ = v83
                                                                        case 128:
                                                                            v21 += (v21 * (v21 + v21))
                                                                            var v84 string = "foo bar"
                                                                            v18 = v18
                                                                            var v85 int32 = v11
                                                                            fmt.Println(v11)
                                                                            _ = v84
                                                                            _ = v85
                                                                        default:
                                                                            if (false && (fi13 == (-(fi13)))) {
                                                                                v82 -= v82
                                                                                v16 += (v20 & (v19 << ((801) & 63)))
                                                                                v86 := GenS6{F0: v11, F1: 127, F2: (v20 & 42)}
                                                                                _ = v86
                                                                            }
                                                                        }
                                                                        _ = v82
                                                                    }
                                                                    v87 := []int{}
                                                                    _ = v87
                                                                }
                                                            }
                                                        case "\n\t":
                                                            _ = len(v4)
                                                            v88 := []int{}
                                                            _ = v88
                                                        default:
                                                            fmt.Println(v21)
                                                        }
                                                        _ = v75
                                                    } else {
                                                        if (((v11 * (v11 % (((-(v11))) | 1))) != (-(v11))) || false) {
                                                            var v89 bool = ((!((("foo bar" + ("\n\t" + (("hello" + "") + "\n\t"))) + (("foo bar" + "hello") + "\n\t")) == ((("foo bar" + ("\n\t" + "a")) + ("foo bar" + "foo bar")) + "a"))) && true)
                                                            v90 := v1[(fi13 * (fi13 + (fi13 | (fi17 | 128))))]
                                                            _ = v89
                                                            _ = v90
                                                        } else {
                                                            v18 += (v18 - -136)
                                                        }
                                                        if (false || false) {
                                                            v19 = v19
                                                            v20 -= (v20 >> (((v19 & v19)) & 63))
                                                        } else {
                                                            if (888 < genFn0(128, v12, (v12 - (^(v12))))) {
                                                                var v91 bool = true
                                                                if true {
                                                                    v92 := v1[511]
                                                                    v93 := GenS11{F0: (v2 | genFn1((^(v2)), "", uint32(v12))), F1: 255, F2: int64(v2)}
                                                                    v94 := GenS12{F0: true, F1: v91}
                                                                    _ = v92
                                                                    _ = v93
                                                                    _ = v94
                                                                } else {
                                                                    v2 += v2
                                                                    if ((3.14159 + v21) == (v21 + (v21 + (float64((float32(v21) * (1.5 / (((3.14159 + float32(1.5)) + float32(v21)) * 1.5)))) / float64(((3.14159 + ((((float32(v21) / -2.25) / 0.1) / (-2.25 / ((-2.25 * 3.14159) * (3.14159 * (float32(v21) - 1.5))))) + 0.0)) + (-2.25 * 1e10))))))) {
                                                                        var v95 bool = (!(v91 || ((v2 ^ v2) > (v2 / ((genFn1(v2, "hello", 4294967295)) | 1)))))
                                                                        var v96 float64 = v21
                                                                        var v97 uint64 = 2
                                                                        if ((!false) || (0 >= -1)) {
                                                                            v16 -= v19
                                                                        } else {
                                                                            v98 := GenS11{F0: genFn1(v97, "a", 4294967295), F1: 0, F2: -167}
                                                                            fmt.Println(v11, v19, v2)
                                                                            v99 := GenS13{F0: "a", F1: ""}
                                                                            _ = v98
                                                                            _ = v99
                                                                        }
                                                                        _ = v95
                                                                        _ = v96
                                                                        _ = v97
                                                                    } else {
                                                                        switch fi13 {
                                                                        case 9223372036854775807:
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                var v100 int = (fi13 ^ 484)
                                                                                _ = v100
                                                                            case "foo bar":
                                                                                var v101 uint64 = (v2 - 1)
                                                                                _ = v101
                                                                            case "\n\t":
                                                                                v19 = v19
                                                                            }
                                                                        case 42:
                                                                            v102 := GenS3{F0: 127, F1: 42, F2: fi13}
                                                                            _ = v102
                                                                        case -137:
                                                                            v19 -= 367
                                                                        default:
                                                                            if true {
                                                                                var v103 uint16 = v12
                                                                                switch fi17 {
                                                                                case 128:
                                                                                    v3 -= v3
                                                                                case 32767:
                                                                                    v3 = v3
                                                                                case -295:
                                                                                    var v104 uint16 = uint16(v2)
                                                                                    v103 = (^(v104))
                                                                                    _ = v104
                                                                                }
                                                                                v105 := GenS8{F0: 2147483648, F1: 39, F2: 42}
                                                                                v19 = (v16 & 65535)
                                                                                switch -287 {
                                                                                case 2147483647:
                                                                                    if (v21 == v21) {
                                                                                        v106 := GenS2{F0: 4294967295, F1: 4294967295, F2: 42}
                                                                                        var v107 bool = v91
                                                                                        _ = len(v1)
                                                                                        _ = v106
                                                                                        _ = v107
                                                                                    } else {
                                                                                        if v91 {
                                                                                            v108 := GenS14{F0: v2, F1: "a", F2: 1}
                                                                                            _ = v108
                                                                                        } else {
                                                                                            switch fi17 {
                                                                                            case 0:
                                                                                                v109 := GenS15{F0: "\n\t", F1: "", F2: 35}
                                                                                                _ = v109
                                                                                            case 750:
                                                                                                var v110 uint = v20
                                                                                                _ = v110
                                                                                            case -1:
                                                                                                v111 := v1[-2147483648]
                                                                                                v12 = v103
                                                                                                var v112 int8 = (v3 + (v3 ^ (v3 / ((25) | 1))))
                                                                                                var v113 float64 = (-2.25 / ((((v21 / 1.0) / 1.5) * v21) / (v21 - 0.0)))
                                                                                                if v111 {
                                                                                                    v11 = (-(v11))
                                                                                                    v114 := GenS1{F0: 395, F1: 143}
                                                                                                    _ = v114
                                                                                                } else {
                                                                                                    v19 = 42
                                                                                                    if (uint32(v2) > 951) {
                                                                                                        v4[1] = (((v91 && true) && v91) && v91)
                                                                                                        var v115 int64 = 1
                                                                                                        if (((("hello" + "foo bar") + "foo bar") + "\n\t") <= ("\n\t" + "foo bar")) {
                                                                                                            v18 += int16(v2)
                                                                                                            var v116 int32 = v11
                                                                                                            v111 = false
                                                                                                            _ = v116
                                                                                                        }
                                                                                                        _ = v115
                                                                                                    } else {
                                                                                                        _ = len(v4)
                                                                                                    }
                                                                                                    fmt.Println(v111)
                                                                                                }
                                                                                                _ = v111
                                                                                                _ = v112
                                                                                                _ = v113
                                                                                            default:
                                                                                                v4[1] = true
                                                                                            }
                                                                                            v117 := v4[42]
                                                                                            _ = v117
                                                                                        }
                                                                                        switch ("" + "hello") {
                                                                                        case "":
                                                                                            switch (fi17 / ((-2147483648) | 1)) {
                                                                                            case 1:
                                                                                                v118 := GenS5{F0: (v18 * 0)}
                                                                                                _ = v118
                                                                                            case 852:
                                                                                                var v119 int64 = 596
                                                                                                v120 := []int{(fi17 % ((2147483647) | 1)), (fi17 % ((726) | 1))}
                                                                                                v16 += v20
                                                                                                if (247 <= 908) {
                                                                                                    v3 = (v3 & 33)
                                                                                                    var v121 uint32 = 0
                                                                                                    v121 = (v121 * 997)
                                                                                                    _ = v121
                                                                                                } else {
                                                                                                    switch (fi13 | (fi13 ^ (-(fi13)))) {
                                                                                                    case 0:
                                                                                                        v11 -= 2147483647
                                                                                                    case 32768:
                                                                                                        if (!(-2.25 == 1e10)) {
                                                                                                            v122 := GenS16{F0: genFn1((v2 << (((v20 ^ v20)) & 63)), ("foo bar" + "\n\t"), 65535), F1: v3}
                                                                                                            v19 -= v20
                                                                                                            _ = len(v1)
                                                                                                            fmt.Println(v91)
                                                                                                            fmt.Println(fi17, v103, v20)
                                                                                                            v103 -= (v12 * (v12 & (v103 >> ((v20) & 63))))
                                                                                                            var v123 uint = (v20 ^ v16)
                                                                                                            var v124 uint32 = 65535
                                                                                                            v125 := GenS17{F0: (v124 * 683), F1: fi17}
                                                                                                            _ = v122
                                                                                                            _ = v123
                                                                                                            _ = v124
                                                                                                            _ = v125
                                                                                                        } else {
                                                                                                            if true {
                                                                                                                v126 := GenS18{F0: uint8(v16), F1: 65535, F2: 613}
                                                                                                                _ = v126
                                                                                                            }
                                                                                                            var v127 uint64 = v2
                                                                                                            v128 := GenS19{F0: (v12 + (v12 * (v12 % ((uint16(v127)) | 1))))}
                                                                                                            _ = v127
                                                                                                            _ = v128
                                                                                                        }
                                                                                                    default:
                                                                                                        switch ("foo bar" + ("hello" + (("" + "") + "a"))) {
                                                                                                        case "":
                                                                                                            var v129 uint = (v19 >> (((v16 << ((uint(v103)) & 63))) & 63))
                                                                                                            _ = v129
                                                                                                        case "hello":
                                                                                                            fmt.Println(fi17)
                                                                                                        case "a":
                                                                                                            if (v18 <= v18) {
                                                                                                                delete(v1, (^(fi13)))
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    switch (-(fi13)) {
                                                                                                    case 0:
                                                                                                        v21 = (v21 * (v21 - (3.14159 - (v21 * v21))))
                                                                                                    case 1:
                                                                                                        if (3.14159 == 3.14159) {
                                                                                                            v130 := GenS10{F0: (-(v3)), F1: 64}
                                                                                                            var v131 uint16 = (v103 >> ((uint(v103)) & 63))
                                                                                                            _ = v130
                                                                                                            _ = v131
                                                                                                        } else {
                                                                                                            fmt.Println(v3, v12, v21)
                                                                                                            v11 = v11
                                                                                                        }
                                                                                                    case -2147483648:
                                                                                                        v132 := GenS18{F0: 42, F1: 65535, F2: (v2 << ((255) & 63))}
                                                                                                        _ = v132
                                                                                                    }
                                                                                                    v119 -= int64(v2)
                                                                                                    v133 := GenS20{F0: 1}
                                                                                                    _ = v133
                                                                                                }
                                                                                                var v134 int = fi17
                                                                                                _ = v119
                                                                                                _ = v120
                                                                                                _ = v134
                                                                                            case 9223372036854775807:
                                                                                                fmt.Println(v12)
                                                                                            }
                                                                                        case "hello":
                                                                                            v135 := []int{-9223372036854775808, (fi17 + (fi13 & (fi13 + (fi17 * fi13)))), 9223372036854775807}
                                                                                            var v136 string = "\n\t"
                                                                                            _ = v135
                                                                                            _ = v136
                                                                                        default:
                                                                                            v3 -= v3
                                                                                        }
                                                                                        var v137 uint = 830
                                                                                        _ = v137
                                                                                    }
                                                                                case 695:
                                                                                    v91 = (v91 || (uint32(v16) != 1))
                                                                                case -9223372036854775808:
                                                                                    v138 := GenS3{F0: v3, F1: 252, F2: 862}
                                                                                    _ = v138
                                                                                default:
                                                                                    if true {
                                                                                        switch "" {
                                                                                        case "":
                                                                                            fmt.Println(v20)
                                                                                        case "a":
                                                                                            v139 := GenS21{F0: v2, F1: v3}
                                                                                            _ = v139
                                                                                        case "\n\t":
                                                                                            if (v91 && ((v3 & -95) >= (v3 | (v3 << (((v20 * v16)) & 63))))) {
                                                                                                if (!(v11 <= -1)) {
                                                                                                    v140 := GenS22{F0: v2}
                                                                                                    _ = v140
                                                                                                }
                                                                                                v141 := v1[(fi17 + int(v3))]
                                                                                                fmt.Println(v141)
                                                                                                _ = v141
                                                                                            } else {
                                                                                                v142 := GenS12{F0: false, F1: (!v91)}
                                                                                                _ = v142
                                                                                            }
                                                                                        default:
                                                                                            if (((int(v2) == fi17) || (v18 != 631)) || v91) {
                                                                                                if v91 {
                                                                                                    if true {
                                                                                                        if v91 {
                                                                                                            fmt.Println(fi13, v91)
                                                                                                            v103 -= (v12 % (((v12 * (v12 / ((v12) | 1)))) | 1))
                                                                                                        }
                                                                                                    } else {
                                                                                                        fmt.Println(v103)
                                                                                                    }
                                                                                                    fmt.Println(v16, v3, v103)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v143 := []int{fi13, fi17, (fi17 * -9223372036854775808), fi13}
                                                                                        var v144 int16 = v18
                                                                                        _ = v143
                                                                                        _ = v144
                                                                                    } else {
                                                                                        if (((v21 - (v21 - ((0.0 / (0.0 + float64(-2.25))) - v21))) - v21) == ((0.1 / v21) * -2.25)) {
                                                                                            delete(v4, 4294967295)
                                                                                        }
                                                                                        fmt.Println(fi13, v2)
                                                                                    }
                                                                                }
                                                                                _ = v103
                                                                                _ = v105
                                                                            } else {
                                                                                var v145 float32 = -2.25
                                                                                var v146 uint64 = v2
                                                                                _ = v145
                                                                                _ = v146
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                var v147 int16 = v18
                                                                fmt.Println(fi17, v2)
                                                                _ = v147
                                                                _ = v91
                                                            } else {
                                                                var v148 uint16 = (v12 | 32767)
                                                                switch fi13 {
                                                                case -1:
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        v148 = v148
                                                                    case "a":
                                                                        v149 := GenS23{F0: v18}
                                                                        _ = v149
                                                                    default:
                                                                        if false {
                                                                            v12 += genFn0(97, (^(v12)), (^(v12)))
                                                                            v21 = v21
                                                                            v12 += 255
                                                                            if false {
                                                                                fmt.Println(fi13, fi13)
                                                                                var v150 int = fi17
                                                                                var v151 int8 = (v3 ^ v3)
                                                                                _ = v150
                                                                                _ = v151
                                                                            }
                                                                        }
                                                                    }
                                                                case 957:
                                                                    var v152 int32 = v11
                                                                    _ = v152
                                                                }
                                                                _ = v148
                                                            }
                                                            fmt.Println(v16, v12, v11)
                                                            switch fi13 {
                                                            case 72:
                                                                var v153 uint16 = (^(v12))
                                                                _ = v153
                                                            case 0:
                                                                var v154 uint32 = 1
                                                                v154 -= v154
                                                                _ = v154
                                                            }
                                                        }
                                                        v155 := GenS24{F0: v19}
                                                        var v156 float64 = 0.1
                                                        _ = v155
                                                        _ = v156
                                                    }
                                                case "foo bar":
                                                    var v157 int32 = (^(v11))
                                                    _ = v157
                                                case "a":
                                                    v158 := GenS15{F0: ("hello" + ("foo bar" + "hello")), F1: "foo bar", F2: 2}
                                                    _ = v158
                                                }
                                                var v159 int = fi13
                                                _ = v159
                                            }
                                            switch (fi17 >> ((v19) & 63)) {
                                            case 155:
                                                v160 := GenS19{F0: (v12 * (v12 & v12))}
                                                _ = v160
                                            default:
                                                v11 = v11
                                            }
                                            v16 -= (^(v19))
                                        }
                                        if false {
                                            v161 := GenS24{F0: (v20 ^ 42)}
                                            v162 := []int{int(v20), int(v19)}
                                            _ = v161
                                            _ = v162
                                        } else {
                                            if ("\n\t" > "foo bar") {
                                                v20 -= 65535
                                            }
                                            var v163 float64 = float64(((0.0 - 1.5) - 1.5))
                                            v16 = v20
                                            _ = v163
                                        }
                                        var v164 uint8 = 229
                                        _ = v164
                                    case "hello":
                                        _ = len(v4)
                                    }
                                }
                                _ = v18
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            v3 = v3
                            var v165 uint8 = 128
                            if (!(1.5 <= (float64(float32(((1.5 * -2.25) * 0.0))) - 1e10))) {
                                switch 32768 {
                                case 1:
                                    var v166 int8 = v3
                                    _ = v166
                                case 984:
                                    v3 += int8(v11)
                                }
                            } else {
                                v4[2147483647] = true
                                v165 = (^(v165))
                                v167 := GenS2{F0: 2147483647, F1: 849, F2: v11}
                                _ = v167
                            }
                            v168 := []int{}
                            _ = fi17
                            _ = v165
                            _ = v168
                        }
                        if ((!false) && (-2.25 >= (float32(float64((0.0 * float32(0.0)))) - 0.1))) {
                            switch fi13 {
                            case 32768:
                                fmt.Println(v2)
                            case 0:
                                v169 := []int{(fi13 * (fi13 << ((v16) & 63))), -314, fi13}
                                _ = v169
                            default:
                                fmt.Println(v12, v3)
                            }
                            switch ("a" + "") {
                            case "\n\t":
                                var v170 int8 = (v3 % ((int8(fi13)) | 1))
                                _ = v170
                            case "a":
                                for fi171 := 0; fi171 < 5; fi171++ {
                                    v172 := []int{}
                                    _ = fi171
                                    _ = v172
                                }
                            case "":
                                for fi173 := 0; fi173 < 7; fi173++ {
                                    v174 := GenS25{F0: -9223372036854775808}
                                    _ = fi173
                                    _ = v174
                                }
                            default:
                                fmt.Println(v11, v3)
                            }
                        }
                    }
                    var v175 uint64 = (v2 & 846)
                    _ = v16
                    _ = v175
                } else {
                    v2 = 1
                    for fi176 := 0; fi176 < 2; fi176++ {
                        if false {
                            fmt.Println(v11, v12)
                        }
                        _ = fi176
                    }
                    var v177 uint32 = 216
                    _ = v177
                }
                v4[65535] = false
                if (-32768 <= 42) {
                    if (((v3 << ((65535) & 63)) > int8(fi13)) && true) {
                        fmt.Println(v3, v2, fi13)
                        for fi178 := 0; fi178 < 15; fi178++ {
                            v12 = (v12 ^ (v12 ^ (v12 | 1)))
                            _ = fi178
                        }
                        var v179 uint32 = 493
                        _ = v179
                    }
                }
            }
        }
        _ = fi13
        _ = v14
    }
    v180 := GenS7{F0: 0, F1: (v3 + (v3 << ((uint(v2)) & 63))), F2: v2}
    v2 -= v2
    for fi181 := 0; fi181 < 13; fi181++ {
        switch (fi181 - -1) {
        case -2147483648:
            var v182 float64 = 0.1
            _ = v182
        case -48:
            if (0.1 < (-2.25 * 1e10)) {
                if false {
                    var v183 uint8 = 195
                    v184 := GenS26{F0: (v11 / ((-1) | 1))}
                    v185 := GenS1{F0: 1, F1: 18446744073709551615}
                    _ = v183
                    _ = v184
                    _ = v185
                } else {
                    if (!false) {
                        if true {
                            if (!false) {
                                v11 -= v11
                            } else {
                                v186 := GenS3{F0: v3, F1: 255, F2: fi181}
                                v187 := []int{}
                                var v188 uint = 1
                                _ = v186
                                _ = v187
                                _ = v188
                            }
                            var v189 uint32 = 2
                            delete(v1, (fi181 * fi181))
                            _ = v189
                        }
                    }
                    var v190 bool = true
                    var v191 uint64 = (^(v2))
                    _ = v190
                    _ = v191
                }
                for fi192 := 0; fi192 < 18; fi192++ {
                    fmt.Println(v2, fi192, v2)
                    var v193 uint8 = 128
                    _ = fi192
                    _ = v193
                }
                v4[2147483648] = true
            }
        case -371:
            v194 := []int{(fi181 - 42), (fi181 * fi181), fi181}
            _ = v194
        default:
            v2 += (v2 << ((65535) & 63))
        }
        _ = fi181
    }
    var v195 int8 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v195:", v195)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v180
    _ = v195
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
