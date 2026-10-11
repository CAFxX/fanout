package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint32; F1 int8; F2 int8 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 int8; F1 int64; F2 uint32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int64; F1 uint8; F2 int64 }
type GenS9 struct { F0 int64; F1 int8 }
type GenS10 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS11 struct { F0 uint; F1 uint; F2 int8 }
type GenS12 struct { F0 int64; F1 int }
type GenS13 struct { F0 string; F1 uint; F2 int8 }
func genFn0(p0 uint32, p1 int32, p2 string) uint8 {
    var v5 float64 = -2.25
    v6 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 bool, p1 int64, p2 uint) int {
    v7 := GenS1{F0: -1}
    if p0 {
        fmt.Println(p0, p1)
        for fi8 := 0; fi8 < 18; fi8++ {
            if (42 > 1) {
                for fi9 := 0; fi9 < 8; fi9++ {
                    var v10 uint16 = 65535
                    var v11 int8 = -128
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                var v12 string = "hello"
                var v13 uint64 = 18446744073709551615
                _ = v12
                _ = v13
            } else {
                p1 += p1
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := []int{fi14, fi8, (fi8 ^ (fi14 % ((fi8) | 1))), fi14}
                    if (18446744073709551615 > p2) {
                        p0 = true
                    }
                    p0 = p0
                    _ = fi14
                    _ = v15
                }
                for fi16 := 0; fi16 < 3; fi16++ {
                    p0 = (((fi16 << (((p2 & (p2 | 2))) & 63)) < (fi16 >> (((p2 - (p2 | (p2 | p2)))) & 63))) || p0)
                    var v17 uint32 = 531
                    fmt.Println(v17, p2)
                    v18 := GenS1{F0: (^(fi16))}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -1
}

func main() {
    var v1 string = "foo bar"
    var v2 int8 = -62
    var v3 int32 = -1
    var v4 uint8 = 16
    var v19 int64 = 42
    for fi20 := 0; fi20 < 17; fi20++ {
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 int64 = 979
            v23 := []int{-2147483648, genFn1(false, (v22 | (v19 * (-(v22)))), 2), -2147483648, fi20}
            _ = fi21
            _ = v22
            _ = v23
        }
        v24 := []int{}
        _ = fi20
        _ = v24
    }
    v3 = (v3 * (v3 * v3))
    switch 42 {
    case -502:
        for fi25 := 0; fi25 < 2; fi25++ {
            switch (fi25 ^ (-(fi25))) {
            case 4294967295:
                if true {
                    var v26 int = (fi25 ^ (fi25 + (fi25 / ((32767) | 1))))
                    if (2 >= 1) {
                        var v27 float32 = 0.1
                        _ = v27
                    }
                    _ = v26
                } else {
                    for fi28 := 0; fi28 < 2; fi28++ {
                        var v29 string = "\n\t"
                        _ = fi28
                        _ = v29
                    }
                    v30 := []int{(^(fi25)), genFn1((false || false), v19, 4294967295)}
                    _ = v30
                }
            default:
                v4 = 255
            }
            var v31 uint32 = 42
            _ = fi25
            _ = v31
        }
    case 9223372036854775807:
        v32 := []int{}
        _ = v32
    case 254:
        v1 = v1
    default:
        v33 := GenS2{F0: v2}
        _ = v33
    }
    fmt.Println(v1, v4)
    v19 = 880
    if (true && true) {
        v2 += (-(v2))
        switch 42 {
        case 0:
            v34 := GenS1{F0: -9223372036854775808}
            _ = v34
        case -20:
            v35 := GenS3{F0: -1, F1: v3}
            _ = v35
        case 42:
            switch genFn1((42 != 0), (v19 & v19), 560) {
            case 397:
                for fi36 := 0; fi36 < 20; fi36++ {
                    var v37 int32 = (v3 - 0)
                    switch v1 {
                    case "":
                        switch 0 {
                        case 127:
                            for fi38 := 0; fi38 < 12; fi38++ {
                                v39 := []int{}
                                v40 := GenS4{F0: 0, F1: (v2 + (v2 & (v2 & v2))), F2: (v2 % (((v2 % (((v2 ^ (v2 * (v2 * (v2 * (v2 % (((v2 ^ int8(v19))) | 1))))))) | 1))) | 1))}
                                _ = fi38
                                _ = v39
                                _ = v40
                            }
                        case 9223372036854775807:
                            if (v19 >= v19) {
                                switch (fi36 << ((1) & 63)) {
                                case -9223372036854775808:
                                    for fi41 := 0; fi41 < 3; fi41++ {
                                        v19 = -9223372036854775808
                                        v2 -= 1
                                        var v42 int64 = (v19 >> ((255) & 63))
                                        var v43 int = (fi36 | fi36)
                                        _ = fi41
                                        _ = v42
                                        _ = v43
                                    }
                                case -110:
                                    for fi44 := 0; fi44 < 17; fi44++ {
                                        fmt.Println(v37, v19)
                                        fmt.Println(v19, v3, v1)
                                        _ = fi44
                                    }
                                }
                                v37 += v37
                                v45 := GenS5{F0: v19, F1: 65535}
                                _ = v45
                            } else {
                                var v46 bool = true
                                if (0.0 < 1.5) {
                                    if (!(32767 >= (fi36 * genFn1(v46, (v19 - (v19 / ((int64(v3)) | 1))), 614)))) {
                                        switch (fi36 >> ((65535) & 63)) {
                                        case 2147483647:
                                            var v47 int64 = v19
                                            _ = v47
                                        case -2147483648:
                                            v48 := GenS6{F0: -112, F1: (v19 | (v19 * v19)), F2: 949}
                                            _ = v48
                                        case 42:
                                            var v49 uint16 = 42
                                            _ = v49
                                        }
                                        for fi50 := 0; fi50 < 9; fi50++ {
                                            v51 := GenS7{F0: (v1 + ("\n\t" + "hello"))}
                                            if (((1.5 / 3.14159) / 3.14159) >= 1e10) {
                                                fmt.Println(fi50)
                                                fmt.Println(fi50, v1)
                                            } else {
                                                if v46 {
                                                    v52 := []int{fi50}
                                                    _ = v52
                                                } else {
                                                    v3 += (v37 << ((319) & 63))
                                                    if (true && (((-1 == (v19 - int64(v2))) || (18446744073709551615 > 42)) && v46)) {
                                                        var v53 uint8 = (v4 % (((v4 << ((42) & 63))) | 1))
                                                        _ = v53
                                                    }
                                                }
                                                var v54 int = int(v3)
                                                v55 := []int{}
                                                _ = v54
                                                _ = v55
                                            }
                                            v56 := GenS7{F0: "a"}
                                            _ = fi50
                                            _ = v51
                                            _ = v56
                                        }
                                    }
                                    if v46 {
                                        v57 := []int{fi36, genFn1((!false), (^(v19)), 0), (fi36 ^ genFn1(v46, 126, 0))}
                                        var v58 int64 = -399
                                        fmt.Println(v37)
                                        _ = v57
                                        _ = v58
                                    } else {
                                        for fi59 := 0; fi59 < 3; fi59++ {
                                            v60 := []int{fi36}
                                            v60 = append(v60, genFn1(v46, v19, 2))
                                            v61 := GenS8{F0: v19, F1: (v4 % (((v4 % (((v4 >> ((2) & 63))) | 1))) | 1)), F2: (v19 ^ v19)}
                                            _ = fi59
                                            _ = v60
                                            _ = v61
                                        }
                                    }
                                    fmt.Println(v2)
                                } else {
                                    v62 := []int{(^(fi36)), -1, (fi36 % ((fi36) | 1))}
                                    var v63 int32 = (v3 * -2147483648)
                                    var v64 int32 = 42
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                }
                                _ = v46
                            }
                        case -534:
                            v37 += (v3 << ((255) & 63))
                        }
                    case "hello":
                        if true {
                            if true {
                                v65 := []int{(fi36 << ((688) & 63)), (fi36 * int(v4))}
                                _ = v65
                            }
                            for fi66 := 0; fi66 < 18; fi66++ {
                                var v67 uint8 = genFn0(1, v3, (("hello" + "foo bar") + ""))
                                switch int(v67) {
                                case -1:
                                    var v68 uint16 = 255
                                    _ = v68
                                case 0:
                                    if ((1 > 2) || false) {
                                        v69 := GenS9{F0: (-(v19)), F1: (v2 / ((v2) | 1))}
                                        _ = v69
                                    } else {
                                        v70 := []int{(^(fi66))}
                                        var v71 float32 = (0.1 * (1e10 + -2.25))
                                        fmt.Println(v4)
                                        v4 = (v4 & genFn0(104, (v3 + (v37 << ((40) & 63))), "hello"))
                                        var v72 string = ""
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                    }
                                    var v73 int8 = -38
                                    var v74 uint32 = 4294967295
                                    _ = v73
                                    _ = v74
                                default:
                                    v75 := []int{(^(fi36))}
                                    _ = v75
                                }
                                v76 := []int{(fi36 * genFn1((2 <= 18446744073709551615), (v19 & (v19 * v19)), 405)), fi66}
                                _ = fi66
                                _ = v67
                                _ = v76
                            }
                            for fi77 := 0; fi77 < 19; fi77++ {
                                v78 := []int{genFn1(false, v19, 792), genFn1((fi77 != 4294967295), 42, 4294967295)}
                                v4 += (v4 * (v4 >> ((801) & 63)))
                                fmt.Println(fi36, v37)
                                _ = fi77
                                _ = v78
                            }
                        }
                    case "foo bar":
                        if (212 != 1000) {
                            if true {
                                var v79 float32 = ((((-2.25 + (float32(1e10) - 0.1)) / 1.5) + 3.14159) + 0.0)
                                if (fi36 != (fi36 ^ fi36)) {
                                    if false {
                                        var v80 uint16 = 42
                                        v81 := GenS7{F0: "hello"}
                                        _ = v80
                                        _ = v81
                                    } else {
                                        var v82 uint = 650
                                        var v83 uint8 = 42
                                        var v84 uint = (v82 & 424)
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                    }
                                }
                                _ = v79
                            }
                            v85 := GenS3{F0: -204, F1: v3}
                            if (42 >= v4) {
                                var v86 bool = (0 == 0)
                                var v87 uint16 = 32767
                                v19 = v19
                                _ = v86
                                _ = v87
                            }
                            _ = v85
                        } else {
                            if false {
                                v2 += v2
                                var v88 bool = false
                                if v88 {
                                    v89 := GenS8{F0: (v19 + v19), F1: v4, F2: v19}
                                    var v90 int8 = v2
                                    var v91 int64 = v19
                                    _ = v89
                                    _ = v90
                                    _ = v91
                                }
                                _ = v88
                            } else {
                                if true {
                                    v92 := GenS5{F0: (-(v19)), F1: 706}
                                    v93 := GenS10{F0: 4294967295, F1: true, F2: 608}
                                    var v94 uint8 = v4
                                    _ = v92
                                    _ = v93
                                    _ = v94
                                } else {
                                    v95 := []int{fi36, (fi36 % ((genFn1(false, 42, 562)) | 1)), (fi36 * fi36)}
                                    _ = v95
                                }
                            }
                            var v96 uint = 0
                            var v97 uint64 = 168
                            _ = v96
                            _ = v97
                        }
                    default:
                        v98 := GenS2{F0: -128}
                        _ = v98
                    }
                    for fi99 := 0; fi99 < 3; fi99++ {
                        v2 = v2
                        v37 = (v3 ^ (v37 >> ((319) & 63)))
                        _ = fi99
                    }
                    _ = fi36
                    _ = v37
                }
            case 127:
                if false {
                    switch v1 {
                    case "\n\t":
                        switch -766 {
                        case 128:
                            switch -780 {
                            case 32768:
                                v100 := []int{}
                                _ = v100
                            case 127:
                                switch 2147483647 {
                                case 4294967295:
                                    for fi101 := 0; fi101 < 13; fi101++ {
                                        v102 := []int{genFn1(true, (-(v19)), 255), 42, -821}
                                        if len(v102) > 0 {
                                            v102[0] = fi101
                                        }
                                        for fi103 := 0; fi103 < 4; fi103++ {
                                            v19 = 42
                                            v4 += genFn0(2147483647, v3, v1)
                                            _ = fi103
                                        }
                                        _ = fi101
                                        _ = v102
                                    }
                                case -1:
                                    v104 := []int{32767, 32767, genFn1(true, 600, 4294967295), 0}
                                    _ = v104
                                case -562:
                                    switch (v1 + "hello") {
                                    case "hello":
                                        fmt.Println(v1)
                                    case "\n\t":
                                        var v105 float32 = 0.1
                                        _ = v105
                                    default:
                                        v3 -= v3
                                    }
                                }
                            default:
                                v106 := []int{4294967295, -746, 9223372036854775807}
                                _ = v106
                            }
                        case -1:
                            v107 := GenS11{F0: 263, F1: 916, F2: v2}
                            _ = v107
                        }
                    case "foo bar":
                        for fi108 := 0; fi108 < 6; fi108++ {
                            v109 := []int{fi108, (fi108 * (fi108 / (((fi108 % (((fi108 & fi108)) | 1))) | 1)))}
                            v19 += v19
                            _ = fi108
                            _ = v109
                        }
                    case "":
                        v110 := GenS12{F0: 963, F1: genFn1(true, -602, 65535)}
                        _ = v110
                    default:
                        v111 := GenS13{F0: ((("" + "hello") + (v1 + "foo bar")) + "hello"), F1: 2, F2: v2}
                        _ = v111
                    }
                    var v112 uint8 = v4
                    if false {
                        v1 = v1
                        var v113 int64 = 435
                        _ = v113
                    }
                    _ = v112
                }
            case 337:
                var v114 int8 = 100
                _ = v114
            default:
                v115 := []int{genFn1((2147483648 <= 2), v19, 1), 32767}
                _ = v115
            }
        default:
            v116 := GenS5{F0: v19, F1: 971}
            _ = v116
        }
    } else {
        v4 -= genFn0(2147483647, (^(v3)), "\n\t")
        fmt.Println(v19, v19)
        v19 -= (v19 | v19)
    }
    switch 32767 {
    case 42:
        switch "foo bar" {
        case "":
            var v117 int16 = 32767
            _ = v117
        case "a":
            for fi118 := 0; fi118 < 7; fi118++ {
                var v119 int8 = v2
                _ = fi118
                _ = v119
            }
        }
    }
    if (((false || (42 >= (v3 * (^(v3))))) && false) && (2 != 0)) {
        v4 -= v4
        v2 += (v2 << ((65535) & 63))
        fmt.Println(v4, v4, v19)
    } else {
        if (0.1 >= float64(1.5)) {
            var v120 int8 = -74
            switch "\n\t" {
            case "\n\t":
                if false {
                    for fi121 := 0; fi121 < 4; fi121++ {
                        for fi122 := 0; fi122 < 8; fi122++ {
                            v2 += v2
                            v120 -= v120
                            _ = fi122
                        }
                        v2 = v2
                        var v123 float64 = ((0.0 / (((1e10 * 0.0) + (1.5 / 0.1)) + (0.1 + 3.14159))) * 1.5)
                        _ = fi121
                        _ = v123
                    }
                    for fi124 := 0; fi124 < 16; fi124++ {
                        var v125 bool = false
                        fmt.Println(v2, v1)
                        for fi126 := 0; fi126 < 2; fi126++ {
                            v19 -= (-(v19))
                            var v127 int32 = (v3 + v3)
                            _ = fi126
                            _ = v127
                        }
                        _ = fi124
                        _ = v125
                    }
                } else {
                    var v128 bool = (!false)
                    v129 := []int{478}
                    _ = v128
                    _ = v129
                }
            }
            for fi130 := 0; fi130 < 13; fi130++ {
                fmt.Println(v4, v3)
                _ = fi130
            }
            _ = v120
        } else {
            for fi131 := 0; fi131 < 17; fi131++ {
                fmt.Println(v1, v2)
                var v132 uint8 = 1
                v3 = (v3 ^ 973)
                _ = fi131
                _ = v132
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
