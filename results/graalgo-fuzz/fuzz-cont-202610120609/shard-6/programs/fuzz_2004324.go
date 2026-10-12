package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int16; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint; F1 string; F2 int64 }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 uint64; F1 int16; F2 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int64; F1 bool; F2 uint64 }
type GenS11 struct { F0 uint32; F1 int16 }
type GenS12 struct { F0 uint16; F1 int8; F2 int }
type GenS13 struct { F0 int16; F1 string; F2 int16 }
type GenS14 struct { F0 int16; F1 uint32 }
type GenS15 struct { F0 int16; F1 uint32 }
func genFn0() uint16 {
    var v4 int16 = 42
    var v5 int16 = (v4 / ((v4) | 1))
    if (true || false) {
        fmt.Println(v5, v4)
        v5 += v5
    } else {
        fmt.Println(v4, v5)
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    switch (("a" + ("" + ("a" + ((("foo bar" + ("\n\t" + "hello")) + "hello") + "\n\t")))) + ("\n\t" + (("" + ("foo bar" + "hello")) + "a"))) {
    case "foo bar":
        var v8 uint8 = 2
        _ = v8
    case "hello":
        var v9 int16 = v5
        _ = v9
    case "":
        if (255 <= 905) {
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := GenS1{F0: v4, F1: -362, F2: 307}
                v12 := []int{fi10, (^(fi10)), (fi10 + (fi10 & -824))}
                if (!true) {
                    v13 := GenS1{F0: (v4 + (v4 | v5)), F1: -2147483648, F2: 65535}
                    _ = v13
                }
                _ = fi10
                _ = v11
                _ = v12
            }
            if (("\n\t" >= "hello") || (32767 <= 1)) {
                v14 := []int{4294967295, 128, 1, 32767}
                fmt.Println(v4, v4)
                var v15 uint = 65535
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 6; fi16++ {
                v17 := GenS1{F0: (v5 - v5), F1: 1, F2: 183}
                switch ("hello" + "a") {
                case "foo bar":
                    v4 += (v5 | (v4 / ((v4) | 1)))
                case "a":
                    var v18 int64 = 597
                    _ = v18
                case "hello":
                    for fi19 := 0; fi19 < 19; fi19++ {
                        if false {
                            if true {
                                var v20 float64 = (float64(1.5) - -2.25)
                                _ = v20
                            } else {
                                if true {
                                    v4 = 1
                                    var v21 int32 = 596
                                    v4 += v5
                                    var v22 int16 = (v5 >> ((1) & 63))
                                    v23 := GenS1{F0: (v5 / (((v4 + (v4 & (v22 & 32767)))) | 1)), F1: (v21 >> ((346) & 63)), F2: 65535}
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                                v24 := []int{}
                                var v25 float32 = (float32(0.0) - -2.25)
                                _ = v24
                                _ = v25
                            }
                        }
                        var v26 float64 = 0.1
                        var v27 uint8 = 17
                        _ = fi19
                        _ = v26
                        _ = v27
                    }
                default:
                    for fi28 := 0; fi28 < 6; fi28++ {
                        var v29 bool = true
                        if (((!v29) && v29) && (v29 && ((98 >= 229) && (255 <= 255)))) {
                            v4 -= v4
                        }
                        _ = fi28
                        _ = v29
                    }
                }
                for fi30 := 0; fi30 < 14; fi30++ {
                    v31 := []int{fi16, (fi16 - (fi30 >> ((367) & 63))), 1}
                    v31 = append(v31, -920)
                    v32 := GenS2{F0: 1, F1: 0}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = fi16
                _ = v17
            }
        } else {
            var v33 int8 = 100
            v34 := GenS3{F0: (v4 >> ((152) & 63)), F1: -402, F2: 2}
            _ = v33
            _ = v34
        }
    default:
        v35 := GenS3{F0: (v4 | 779), F1: 0, F2: 18446744073709551615}
        _ = v35
    }
    _ = v4
    _ = v5
    return 32767
}
func genFn1(p0 int16, p1 int32) int {
    p0 += 1000
    var v36 uint64 = 2
    var v37 bool = (128 > 2)
    var v38 bool = (true || true)
    for fi39 := 0; fi39 < 8; fi39++ {
        v40 := []int{(fi39 & 32768), (fi39 / ((fi39) | 1)), (^(fi39)), fi39}
        _ = fi39
        _ = v40
    }
    _ = p0
    _ = p1
    _ = v36
    _ = v37
    _ = v38
    return 127
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{166, 785, 504, 684}
    var v3 uint32 = 65535
    var v41 string = "a"
    switch -401 {
    case 42:
        var v42 string = v41
        _ = v42
    case 2147483647:
        if false {
            switch ((v41 + "") + v41) {
            case "\n\t":
                if (!(!false)) {
                    var v43 uint32 = (v3 % ((v3) | 1))
                    var v44 uint8 = 2
                    _ = v43
                    _ = v44
                } else {
                    var v45 int8 = -128
                    v41 = (("\n\t" + ((("\n\t" + "foo bar") + (("" + (v41 + ("\n\t" + (v41 + v41)))) + "")) + (v41 + v41))) + v41)
                    v41 = ""
                    _ = v45
                }
            }
            switch 9223372036854775807 {
            case 0:
                v41 = (("" + v41) + ("a" + (v41 + (((v41 + v41) + "") + ("a" + "")))))
            case 649:
                for fi46 := 0; fi46 < 11; fi46++ {
                    fmt.Println(v41)
                    for fi47 := 0; fi47 < 19; fi47++ {
                        var v48 uint = 255
                        _ = fi47
                        _ = v48
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi46
                }
            default:
                if false {
                    v49 := GenS4{F0: 65535, F1: v41, F2: -9223372036854775808}
                    v41 = (v41 + (v41 + "\n\t"))
                    _ = v49
                }
            }
        } else {
            for fi50 := 0; fi50 < 6; fi50++ {
                if true {
                    v51 := GenS5{F0: -1, F1: uint16(v3)}
                    _ = v51
                } else {
                    switch fi50 {
                    case 232:
                        fmt.Println(v41, v41, fi50)
                    case -189:
                        if false {
                            v52 := GenS6{F0: 4294967295, F1: (!(32767 < 255))}
                            for fi53 := 0; fi53 < 12; fi53++ {
                                fmt.Println(fi53, v41)
                                _ = fi53
                            }
                            if (false || false) {
                                v2 = append(v2, (v3 >> ((639) & 63)))
                                if false {
                                    var v54 int8 = 61
                                    fmt.Println(fi50)
                                    var v55 int64 = 324
                                    _ = v54
                                    _ = v55
                                } else {
                                    switch (fi50 >> ((1) & 63)) {
                                    case -2147483648:
                                        v56 := GenS1{F0: 0, F1: 598, F2: 371}
                                        _ = v56
                                    }
                                }
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = (v3 % ((1) | 1))
                                }
                                v3 = 2147483648
                                switch fi50 {
                                case 0:
                                    for fi57 := 0; fi57 < 20; fi57++ {
                                        switch v41 {
                                        case "":
                                            if false {
                                                var v58 int32 = 42
                                                _ = v58
                                            } else {
                                                switch fi50 {
                                                case 2147483647:
                                                    if true {
                                                        var v59 int64 = 507
                                                        _ = v59
                                                    }
                                                case 128:
                                                    v60 := GenS6{F0: uint64(v3), F1: (-2.25 <= float32(3.14159))}
                                                    _ = v60
                                                case 4294967295:
                                                    if ((float64(float32((((1.5 * (3.14159 / (float64(float32(((-2.25 - (float64(0.1) - 0.0)) / 0.1))) + 0.0))) * (-2.25 - (0.0 + -2.25))) * 1.5))) * (1.5 * 0.0)) == 1e10) {
                                                        v41 = v41
                                                        v41 = "foo bar"
                                                    } else {
                                                        var v61 float64 = float64(float32((((1.5 - float64((float32((1.5 + -2.25)) - -2.25))) * 0.1) - 3.14159)))
                                                        _ = v61
                                                    }
                                                }
                                                var v62 int64 = 0
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                delete(v1, 1)
                                                if (!false) {
                                                    var v63 uint16 = genFn0()
                                                    _ = v63
                                                } else {
                                                    if (v62 == -417) {
                                                        v62 = (v62 >> ((42) & 63))
                                                        switch v41 {
                                                        case "foo bar":
                                                            v64 := GenS7{F0: 18446744073709551615, F1: -32768, F2: -492}
                                                            _ = v64
                                                        case "hello":
                                                            var v65 uint32 = v3
                                                            _ = v65
                                                        case "":
                                                            switch (fi57 - (fi57 ^ int(v62))) {
                                                            case 32768:
                                                                v66 := GenS2{F0: 720, F1: 32767}
                                                                _ = v66
                                                            default:
                                                                if false {
                                                                    if (v62 != v62) {
                                                                        var v67 uint8 = 162
                                                                        v2 = append(v2, 510)
                                                                        _ = v67
                                                                    } else {
                                                                        switch ("a" + v41) {
                                                                        case "hello":
                                                                            if (false || (false || (!true))) {
                                                                                var v68 int64 = (-(v62))
                                                                                _ = v68
                                                                            } else {
                                                                                delete(v1, v3)
                                                                                var v69 string = "foo bar"
                                                                                var v70 uint8 = 128
                                                                                _ = v69
                                                                                _ = v70
                                                                            }
                                                                        case "":
                                                                            v71 := GenS5{F0: 0, F1: 1}
                                                                            _ = v71
                                                                        case "foo bar":
                                                                            v72 := GenS5{F0: 1, F1: 255}
                                                                            _ = v72
                                                                        }
                                                                    }
                                                                    var v73 uint8 = 128
                                                                    var v74 uint32 = 65535
                                                                    var v75 int64 = -914
                                                                    _ = v73
                                                                    _ = v74
                                                                    _ = v75
                                                                }
                                                            }
                                                        }
                                                    }
                                                    var v76 int8 = 100
                                                    fmt.Println(v3)
                                                    _ = v76
                                                }
                                                _ = v62
                                            }
                                        }
                                        v41 = ("\n\t" + "")
                                        if false {
                                            v77 := GenS8{F0: v41}
                                            v41 = v41
                                            var v78 int32 = 0
                                            _ = v77
                                            _ = v78
                                        } else {
                                            v79 := GenS9{F0: false}
                                            if (!true) {
                                                v41 = ((v41 + "hello") + "")
                                            } else {
                                                var v80 bool = false
                                                var v81 bool = ((18446744073709551615 == 1) || v80)
                                                _ = v80
                                                _ = v81
                                            }
                                            switch "foo bar" {
                                            case "":
                                                v82 := GenS10{F0: 42, F1: false, F2: 18446744073709551615}
                                                _ = v82
                                            case "\n\t":
                                                switch 32767 {
                                                case 0:
                                                    if true {
                                                        fmt.Println(fi57)
                                                    } else {
                                                        v2 = append(v2, 2147483648)
                                                        v41 = ((v41 + "") + (("\n\t" + "") + v41))
                                                        fmt.Println(fi57)
                                                        var v83 uint16 = 241
                                                        _ = v83
                                                    }
                                                case 73:
                                                    if len(v2) > 0 {
                                                        v2[0] = (v3 - v3)
                                                    }
                                                }
                                            default:
                                                fmt.Println(fi50)
                                            }
                                            _ = v79
                                        }
                                        _ = fi57
                                    }
                                case -565:
                                    v2 = append(v2, (v3 & (v3 - v3)))
                                }
                            }
                            _ = v52
                        } else {
                            switch (^(fi50)) {
                            case 0:
                                var v84 uint = 394
                                _ = v84
                            case 9223372036854775807:
                                if false {
                                    for fi85 := 0; fi85 < 12; fi85++ {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v86 := GenS11{F0: v3, F1: 234}
                                        if true {
                                            if len(v2) > 0 {
                                                v2[0] = v3
                                            }
                                            var v87 int8 = 0
                                            delete(v1, (^(v3)))
                                            _ = v87
                                        }
                                        _ = fi85
                                        _ = v86
                                    }
                                    switch (v41 + (v41 + v41)) {
                                    case "\n\t":
                                        v3 += v3
                                    case "a":
                                        v3 = 0
                                    case "":
                                        if true {
                                            switch fi50 {
                                            case 32767:
                                                for fi88 := 0; fi88 < 18; fi88++ {
                                                    var v89 string = "\n\t"
                                                    v90 := GenS12{F0: 255, F1: 127, F2: (fi88 & genFn1(1000, 42))}
                                                    v41 = v89
                                                    _ = fi88
                                                    _ = v89
                                                    _ = v90
                                                }
                                            case 0:
                                                v91 := GenS11{F0: (v3 | v3), F1: -875}
                                                _ = v91
                                            case 9223372036854775807:
                                                var v92 uint32 = 355
                                                _ = v92
                                            }
                                            if (!(100 < -1)) {
                                                fmt.Println(v41, fi50)
                                                v3 = v3
                                            }
                                            if false {
                                                switch v41 {
                                                case "":
                                                    for fi93 := 0; fi93 < 2; fi93++ {
                                                        var v94 int8 = 100
                                                        fmt.Println(fi93, v41)
                                                        _ = fi93
                                                        _ = v94
                                                    }
                                                case "\n\t":
                                                    if (1 < 58) {
                                                        fmt.Println(v41, v41)
                                                    }
                                                }
                                            }
                                        } else {
                                            v95 := GenS7{F0: 198, F1: 617, F2: -2147483648}
                                            _ = v95
                                        }
                                    default:
                                        v96 := GenS12{F0: genFn0(), F1: -1, F2: (^(fi50))}
                                        _ = v96
                                    }
                                }
                            case -2147483648:
                                v3 = v3
                            default:
                                var v97 uint = 4294967295
                                _ = v97
                            }
                            v41 = (v41 + v41)
                        }
                    case -678:
                        if true {
                            v98 := GenS13{F0: 513, F1: v41, F2: 42}
                            v41 = v41
                            _ = v98
                        }
                    default:
                        fmt.Println(v41, v3)
                    }
                    var v99 uint16 = 768
                    for fi100 := 0; fi100 < 20; fi100++ {
                        var v101 bool = false
                        if len(v2) > 0 {
                            v2[0] = 0
                        }
                        _ = fi100
                        _ = v101
                    }
                    _ = v99
                }
                var v102 float32 = 3.14159
                _ = fi50
                _ = v102
            }
            v41 = v41
        }
    }
    fmt.Println(v41)
    if ((-2.25 != 0.1) && ((1e10 + 3.14159) > 0.0)) {
        if (-32768 <= 948) {
            for fi103 := 0; fi103 < 6; fi103++ {
                v3 = (v3 - (v3 + 280))
                var v104 float32 = 1e10
                _ = fi103
                _ = v104
            }
        }
        var v105 int32 = -985
        _ = v105
    }
    fmt.Println(v3, v3)
    fmt.Println(v41, v41)
    var v106 uint64 = 18446744073709551615
    if false {
        fmt.Println(v41, v106, v41)
        v107 := GenS14{F0: 32767, F1: v3}
        v106 = v106
        _ = v107
    } else {
        var v108 uint64 = (v106 % ((v106) | 1))
        fmt.Println(v41)
        _ = v108
    }
    if (0.0 == (3.14159 / 0.1)) {
        fmt.Println(v106, v41, v3)
        var v109 uint32 = (v3 | v3)
        fmt.Println(v3)
        _ = v109
    } else {
        v3 -= uint32(v106)
        if (!false) {
            var v110 int = -2147483648
            var v111 float32 = float32(1.5)
            _ = v110
            _ = v111
        } else {
            if false {
                fmt.Println(v41)
                if false {
                    var v112 int16 = 0
                    _ = v112
                } else {
                    if (false && (1 < -2147483648)) {
                        v3 += (v3 | v3)
                    } else {
                        delete(v1, v3)
                        for fi113 := 0; fi113 < 10; fi113++ {
                            v41 = "\n\t"
                            delete(v1, v3)
                            _ = fi113
                        }
                    }
                    for fi114 := 0; fi114 < 10; fi114++ {
                        var v115 int32 = 519
                        for fi116 := 0; fi116 < 5; fi116++ {
                            v117 := GenS15{F0: -32768, F1: 148}
                            v3 = v3
                            _ = fi116
                            _ = v117
                        }
                        var v118 uint16 = 901
                        _ = fi114
                        _ = v115
                        _ = v118
                    }
                }
            }
            v119 := GenS12{F0: 42, F1: 1, F2: 9223372036854775807}
            var v120 uint = 0
            _ = v119
            _ = v120
        }
    }
    delete(v1, v3)
    for fi121 := 0; fi121 < 4; fi121++ {
        var v122 uint8 = uint8(v3)
        var v123 uint64 = v106
        _ = fi121
        _ = v122
        _ = v123
    }
    if true {
        var v124 float32 = 0.1
        v2 = append(v2, (v3 * (v3 % ((v3) | 1))))
        _ = v124
    }
    fmt.Println(v41, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v106:", v106)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v106
    _ = v2
    _ = v3
    _ = v41
}
