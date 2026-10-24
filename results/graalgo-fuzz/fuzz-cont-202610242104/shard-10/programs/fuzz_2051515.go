package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 string; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS6 struct { F0 string; F1 int64; F2 string }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int8; F1 uint; F2 int }
type GenS10 struct { F0 bool; F1 uint32 }
type GenS11 struct { F0 uint8; F1 bool; F2 int64 }
type GenS12 struct { F0 string; F1 int32; F2 string }
type GenS13 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS14 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int32) uint32 {
    switch 128 {
    case 1:
        for fi6 := 0; fi6 < 4; fi6++ {
            p1 = p1
            v7 := GenS1{F0: 2, F1: fi6}
            _ = fi6
            _ = v7
        }
    case 127:
        var v8 int8 = 127
        _ = v8
    case 482:
        v9 := []int{127, 110, 1}
        _ = v9
    default:
        v10 := []int{9223372036854775807, -9223372036854775808, 42, 1}
        _ = v10
    }
    v11 := GenS2{F0: (p1 + -1), F1: 4294967295}
    fmt.Println(p1, p1)
    v12 := []int{}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 4294967295
}
func genFn1(p0 uint32, p1 uint8) int8 {
    fmt.Println(p0)
    var v13 uint64 = 2
    v14 := GenS3{F0: ("hello" + "foo bar"), F1: 65535, F2: 83}
    p0 += 2147483648
    if false {
        p1 -= 42
        switch "hello" {
        case "a":
            var v15 bool = (1e10 >= -2.25)
            _ = v15
        case "foo bar":
            v16 := []int{}
            _ = v16
        case "":
            var v17 uint64 = (v13 - v13)
            _ = v17
        default:
            p1 -= p1
        }
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 1
}

func main() {
    var v1 int32 = -383
    v2 := make(map[uint32]string)
    var v3 float32 = float32(3.14159)
    var v4 string = "\n\t"
    v5 := make(map[uint32]bool)
    var v18 uint64 = 498
    v1 -= -1
    if false {
        for fi19 := 0; fi19 < 20; fi19++ {
            v18 = 42
            for fi20 := 0; fi20 < 15; fi20++ {
                v21 := GenS2{F0: v1, F1: 943}
                switch v4 {
                case "hello":
                    v4 = (v4 + ("hello" + "hello"))
                case "a":
                    var v22 string = "a"
                    _ = v22
                }
                v3 = v3
                _ = fi20
                _ = v21
            }
            v1 += v1
            _ = fi19
        }
    } else {
        switch 305 {
        case 128:
            for fi23 := 0; fi23 < 6; fi23++ {
                for fi24 := 0; fi24 < 14; fi24++ {
                    var v25 bool = true
                    var v26 string = (v4 + v4)
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi23
            }
        case -1:
            var v27 uint = 890
            _ = v27
        case 52:
            v3 += v3
        default:
            var v28 float32 = v3
            _ = v28
        }
        var v29 int = -9223372036854775808
        if (128 <= 128) {
            for fi30 := 0; fi30 < 8; fi30++ {
                fmt.Println(fi30)
                v4 = v4
                if (!((-(fi30)) == 127)) {
                    switch ("hello" + v4) {
                    case "\n\t":
                        v18 += v18
                    case "hello":
                        v4 = "hello"
                    case "foo bar":
                        var v31 float64 = 0.1
                        _ = v31
                    default:
                        fmt.Println(v18, fi30)
                    }
                } else {
                    fmt.Println(v1, v4)
                    var v32 bool = false
                    _ = v32
                }
                _ = fi30
            }
            v4 = v4
            for fi33 := 0; fi33 < 4; fi33++ {
                for fi34 := 0; fi34 < 17; fi34++ {
                    v35 := []int{fi33}
                    var v36 float64 = 0.1
                    v18 -= (v18 ^ (v18 + (v18 - 4294967295)))
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                if true {
                    v2[42] = (v4 + "foo bar")
                } else {
                    if (!((true && false) || (-1 == -128))) {
                        v1 += v1
                        var v37 uint16 = 42
                        _ = v37
                    } else {
                        v4 = v4
                        if (18446744073709551615 == 255) {
                            var v38 float64 = 3.14159
                            var v39 int = 9223372036854775807
                            if (18446744073709551615 == 18446744073709551615) {
                                v40 := GenS4{F0: 0, F1: 42}
                                var v41 uint32 = 2147483648
                                if true {
                                    switch ("\n\t" + (v4 + "hello")) {
                                    case "\n\t":
                                        var v42 float64 = float64(v3)
                                        _ = v42
                                    case "":
                                        for fi43 := 0; fi43 < 6; fi43++ {
                                            v39 -= (-(v29))
                                            v2[(^(v41))] = ""
                                            fmt.Println(v39)
                                            var v44 bool = (4294967295 >= 18446744073709551615)
                                            v45 := []int{}
                                            _ = fi43
                                            _ = v44
                                            _ = v45
                                        }
                                    case "a":
                                        v46 := GenS5{F0: -1, F1: 705, F2: v18}
                                        _ = v46
                                    }
                                } else {
                                    for fi47 := 0; fi47 < 20; fi47++ {
                                        v48 := GenS4{F0: 363, F1: -972}
                                        _ = fi47
                                        _ = v48
                                    }
                                    v49 := GenS6{F0: v4, F1: int64(v1), F2: "a"}
                                    var v50 int = -9223372036854775808
                                    _ = v49
                                    _ = v50
                                }
                                _ = v40
                                _ = v41
                            }
                            _ = v38
                            _ = v39
                        }
                    }
                    v51 := GenS7{F0: ("foo bar" + "foo bar"), F1: 0}
                    if (42 != 2) {
                        var v52 float64 = -2.25
                        v53 := v5[623]
                        _ = v52
                        _ = v53
                    }
                    _ = v51
                }
                _ = fi33
            }
        } else {
            var v54 int16 = 32767
            switch v4 {
            case "":
                var v55 float32 = v3
                _ = v55
            case "foo bar":
                v54 = v54
            }
            _ = v54
        }
        _ = v29
    }
    var v56 uint16 = 831
    var v57 string = (v4 + v4)
    v58 := GenS8{F0: 952}
    v18 = (v18 ^ 42)
    v59 := GenS9{F0: genFn1(genFn0(2, (^(v1))), 2), F1: 255, F2: 32768}
    if (-2.25 <= -2.25) {
        v60 := GenS6{F0: "a", F1: -699, F2: v57}
        for fi61 := 0; fi61 < 14; fi61++ {
            var v62 uint16 = 230
            switch (fi61 & int(v56)) {
            case 32767:
                var v63 uint32 = 65535
                _ = v63
            case 1:
                var v64 int = fi61
                _ = v64
            }
            _ = fi61
            _ = v62
        }
        _ = v60
    } else {
        if false {
            for fi65 := 0; fi65 < 8; fi65++ {
                fmt.Println(v4, v18, v57)
                if (false && (false || true)) {
                    v4 = ""
                }
                _ = fi65
            }
            var v66 bool = true
            _ = v66
        }
        switch v4 {
        case "hello":
            var v67 int8 = 0
            _ = v67
        case "\n\t":
            if (!(v4 < (v57 + (v57 + ((v57 + v4) + "hello"))))) {
                delete(v2, 717)
                for fi68 := 0; fi68 < 7; fi68++ {
                    if (!(-9223372036854775808 != 1)) {
                        var v69 int16 = 55
                        if false {
                            fmt.Println(fi68, v57, v57)
                            var v70 uint16 = 822
                            _ = v70
                        } else {
                            v57 = "\n\t"
                            var v71 uint64 = 1
                            v57 = v57
                            _ = v71
                        }
                        _ = v69
                    }
                    _ = fi68
                }
                v4 = v57
            } else {
                var v72 int8 = -47
                _ = v72
            }
        case "a":
            var v73 int64 = -1
            _ = v73
        default:
            fmt.Println(v18)
        }
        if (65535 <= 0) {
            switch 127 {
            case 128:
                for fi74 := 0; fi74 < 3; fi74++ {
                    v56 -= 65535
                    _ = fi74
                }
            }
            fmt.Println(v1)
            v3 = float32(3.14159)
        } else {
            if false {
                var v75 int64 = 9223372036854775807
                var v76 bool = (3.14159 == (v3 + -2.25))
                var v77 uint64 = 0
                _ = v75
                _ = v76
                _ = v77
            } else {
                var v78 int8 = 0
                for fi79 := 0; fi79 < 7; fi79++ {
                    switch v57 {
                    case "hello":
                        if true {
                            for fi80 := 0; fi80 < 5; fi80++ {
                                switch ("foo bar" + v4) {
                                case "\n\t":
                                    v81 := GenS10{F0: (!((!true) || false)), F1: genFn0(v56, v1)}
                                    _ = v81
                                default:
                                    if true {
                                        v82 := []int{(fi80 << ((255) & 63)), 128, fi79, int(v18)}
                                        _ = v82
                                    }
                                }
                                v83 := GenS11{F0: 255, F1: false, F2: -9223372036854775808}
                                v84 := []int{9223372036854775807, (fi80 & (-(fi80)))}
                                _ = fi80
                                _ = v83
                                _ = v84
                            }
                            var v85 float32 = (v3 - v3)
                            _ = v85
                        }
                    }
                    v86 := GenS1{F0: 21, F1: 169}
                    if (!false) {
                        v87 := GenS6{F0: (v4 + v57), F1: 1, F2: ""}
                        fmt.Println(fi79)
                        for fi88 := 0; fi88 < 2; fi88++ {
                            v1 -= v1
                            _ = fi88
                        }
                        _ = v87
                    } else {
                        v3 -= (v3 - v3)
                        fmt.Println(fi79)
                        var v89 uint32 = 606
                        _ = v89
                    }
                    _ = fi79
                    _ = v86
                }
                _ = v78
            }
        }
    }
    if ((((v3 - v3) * v3) - (v3 - 1e10)) == float32(0.0)) {
        for fi90 := 0; fi90 < 12; fi90++ {
            v91 := GenS3{F0: (v57 + ""), F1: v56, F2: 254}
            delete(v2, 2147483648)
            _ = fi90
            _ = v91
        }
        for fi92 := 0; fi92 < 18; fi92++ {
            for fi93 := 0; fi93 < 8; fi93++ {
                var v94 uint64 = 2
                v95 := GenS9{F0: 1, F1: 4294967295, F2: fi93}
                _ = fi93
                _ = v94
                _ = v95
            }
            var v96 int64 = 42
            if (v3 != ((1.5 + 1e10) - (-2.25 / 0.1))) {
                fmt.Println(v18, v3)
                v97 := GenS10{F0: (!false), F1: 4294967295}
                v98 := GenS11{F0: 192, F1: false, F2: (v96 & v96)}
                _ = v97
                _ = v98
            } else {
                fmt.Println(v18, fi92)
                var v99 uint8 = 181
                v99 -= (v99 - v99)
                _ = v99
            }
            _ = fi92
            _ = v96
        }
    } else {
        v56 += v56
    }
    var v100 uint16 = (v56 + v56)
    if false {
        var v101 uint32 = 2147483648
        v3 += float32(((-2.25 - (0.0 - (0.0 - (-2.25 / (((1.5 / 0.1) / 3.14159) + 1e10))))) * 1.5))
        var v102 uint32 = v101
        _ = v101
        _ = v102
    } else {
        v103 := v2[550]
        switch 42 {
        case 127:
            switch 480 {
            case 530:
                for fi104 := 0; fi104 < 4; fi104++ {
                    v105 := v2[2147483647]
                    _ = fi104
                    _ = v105
                }
            case -805:
                if false {
                    switch 127 {
                    case 42:
                        v106 := GenS12{F0: (("a" + ("foo bar" + (v57 + "foo bar"))) + v57), F1: v1, F2: ("foo bar" + "")}
                        _ = v106
                    case -9223372036854775808:
                        var v107 float64 = (-2.25 - 1e10)
                        _ = v107
                    case 4294967295:
                        switch v4 {
                        case "\n\t":
                            if (65535 >= 512) {
                                v18 = 2
                            }
                        case "hello":
                            var v108 int64 = -1
                            _ = v108
                        default:
                            for fi109 := 0; fi109 < 11; fi109++ {
                                if (2 < 118) {
                                    for fi110 := 0; fi110 < 15; fi110++ {
                                        if (v56 != v56) {
                                            var v111 bool = false
                                            var v112 uint8 = 42
                                            var v113 float64 = -2.25
                                            _ = v111
                                            _ = v112
                                            _ = v113
                                        } else {
                                            _ = len(v5)
                                            switch fi109 {
                                            case -922:
                                                v114 := GenS13{F0: 2147483648, F1: (^(v18)), F2: -32768}
                                                _ = v114
                                            case 32767:
                                                switch (fi109 / ((fi109) | 1)) {
                                                case 149:
                                                    v115 := GenS7{F0: v57, F1: (v100 + v56)}
                                                    _ = v115
                                                }
                                            }
                                            var v116 uint = 279
                                            _ = v116
                                        }
                                        var v117 int32 = v1
                                        if (255 <= uint16(v18)) {
                                            if false {
                                                v4 = ((v103 + ((("\n\t" + ((("hello" + "\n\t") + v57) + "hello")) + (("a" + (v4 + "")) + v4)) + "hello")) + "a")
                                                var v118 float32 = float32(1.5)
                                                var v119 float32 = (((v3 * float32(0.1)) * v118) - v118)
                                                _ = v118
                                                _ = v119
                                            } else {
                                                var v120 int = fi110
                                                v1 += (v117 * (v1 / ((int32(v56)) | 1)))
                                                v3 -= (v3 - v3)
                                                switch 42 {
                                                case 32767:
                                                    var v121 int16 = 0
                                                    _ = v121
                                                }
                                                _ = v120
                                            }
                                            var v122 int8 = -128
                                            v123 := GenS12{F0: "", F1: 28, F2: v57}
                                            _ = v122
                                            _ = v123
                                        }
                                        _ = fi110
                                        _ = v117
                                    }
                                    v1 += -778
                                    v56 += (v56 % (((^(v100))) | 1))
                                } else {
                                    var v124 float64 = -2.25
                                    v2[219] = v4
                                    v103 = (v103 + v57)
                                    _ = v124
                                }
                                v125 := GenS14{F0: 710}
                                var v126 bool = (((42 == 1) && (true && (v3 >= 1e10))) && (!false))
                                _ = fi109
                                _ = v125
                                _ = v126
                            }
                        }
                    default:
                        v103 = v4
                    }
                    var v127 float64 = 3.14159
                    _ = v127
                }
            case 4294967295:
                fmt.Println(v4, v1, v57)
            default:
                v128 := GenS13{F0: 142, F1: v18, F2: -592}
                _ = v128
            }
        case -1:
            v100 -= (v56 % ((v56) | 1))
        }
        _ = v103
    }
    v129 := []int{-67, -367}
    v130 := v2[1]
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("slicelen v129:", len(v129), v129)
    fmt.Println("v130:", v130)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v100
    _ = v129
    _ = v130
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}
