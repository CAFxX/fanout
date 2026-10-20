package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool; F1 uint32; F2 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint16; F1 uint16 }
type GenS10 struct { F0 uint8; F1 uint }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 bool; F1 int8 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS16 struct { F0 int32; F1 uint64 }
func genFn0() uint32 {
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int = (fi5 ^ fi5)
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := []int{}
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    var v9 bool = true
    _ = v9
    return 42
}
func genFn1(p0 uint8, p1 uint32, p2 bool) uint64 {
    if p2 {
        var v10 int8 = 127
        switch "\n\t" {
        case "a":
            p2 = p2
        default:
            v11 := []int{209, 9223372036854775807, -544, 64}
            _ = v11
        }
        _ = v10
    }
    var v12 uint32 = genFn0()
    if p2 {
        if (false || (p2 && (!true))) {
            p1 = p1
            fmt.Println(p0, v12)
            v13 := []int{2147483647}
            _ = v13
        } else {
            if (true || (1 <= 32767)) {
                switch 42 {
                case 414:
                    fmt.Println(v12)
                case 4294967295:
                    v14 := []int{32767, -137, 0}
                    _ = v14
                }
                v12 += 798
                v12 = (v12 >> ((2) & 63))
            } else {
                var v15 int32 = 42
                for fi16 := 0; fi16 < 6; fi16++ {
                    if p2 {
                        v17 := GenS1{F0: -1, F1: 1}
                        v18 := GenS2{F0: (p1 & (v12 >> ((0) & 63))), F1: p2, F2: 47}
                        v12 = (v12 + (v12 >> ((688) & 63)))
                        _ = v17
                        _ = v18
                    }
                    fmt.Println(p0, fi16, fi16)
                    _ = fi16
                }
                _ = v15
            }
            v19 := GenS3{F0: 785}
            if p2 {
                p2 = false
                for fi20 := 0; fi20 < 15; fi20++ {
                    for fi21 := 0; fi21 < 16; fi21++ {
                        var v22 int8 = 1
                        v23 := GenS2{F0: (v12 ^ (p1 - genFn0())), F1: p2, F2: 530}
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    p0 -= (p0 & p0)
                    _ = fi20
                }
            } else {
                for fi24 := 0; fi24 < 10; fi24++ {
                    var v25 float64 = (3.14159 + float64(1e10))
                    p2 = true
                    _ = fi24
                    _ = v25
                }
                p1 = 200
                fmt.Println(v12, v12)
            }
            _ = v19
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (2147483647 < 76)
    var v3 float32 = (0.0 / 1e10)
    v4 := make(map[string]uint32)
    var v26 int = 4294967295
    fmt.Println(v2, v2, v26)
    v1[2147483648] = v26
    if false {
        fmt.Println(v26, v26, v3)
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        switch "\n\t" {
        case "hello":
            if (true && ("\n\t" < ((("a" + "foo bar") + (("foo bar" + "\n\t") + "hello")) + "foo bar"))) {
                var v28 uint16 = 65535
                v29 := GenS3{F0: 255}
                _ = v28
                _ = v29
            }
        case "a":
            v26 -= v26
        default:
            for fi30 := 0; fi30 < 14; fi30++ {
                v31 := GenS2{F0: genFn0(), F1: ((false || (0 != 797)) && ((genFn0() < genFn0()) && (2147483648 == 762))), F2: 165}
                var v32 uint = 255
                _ = fi30
                _ = v31
                _ = v32
            }
        }
        switch 26 {
        case -355:
            if false {
                v33 := []int{}
                var v34 int32 = 1
                _ = v33
                _ = v34
            } else {
                for fi35 := 0; fi35 < 4; fi35++ {
                    if (v3 <= float32(1.5)) {
                        v36 := GenS3{F0: 65535}
                        var v37 int64 = 9223372036854775807
                        v38 := GenS4{F0: fi35}
                        _ = v36
                        _ = v37
                        _ = v38
                    } else {
                        v39 := GenS5{F0: (fi35 | 0)}
                        v40 := []int{}
                        v2 = (fi27 < (v26 % ((v26) | 1)))
                        v41 := GenS5{F0: (fi27 - (v26 * (fi27 + fi27)))}
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                    _ = fi35
                }
                var v42 int32 = -755
                _ = v42
            }
        }
        _ = fi27
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        v2 = (18446744073709551615 > 1)
        _ = fi43
    }
    v44 := GenS6{F0: (!((("\n\t" + "\n\t") <= "a") || (false && (!v2)))), F1: 65535, F2: 42}
    var v45 uint64 = 2
    switch v26 {
    case 42:
        switch "foo bar" {
        case "":
            v46 := []int{}
            _ = v46
        case "hello":
            var v47 int8 = 0
            _ = v47
        }
    case 1:
        fmt.Println(v45)
    case 610:
        _ = len(v1)
    }
    switch (v26 + -2147483648) {
    case 128:
        var v48 uint64 = 886
        _ = v48
    }
    if v2 {
        if v2 {
            v3 = (v3 * ((((3.14159 * 3.14159) * (v3 + v3)) * float32((3.14159 / ((1.5 * 1e10) / 1.5)))) * v3))
            v49 := GenS7{F0: -443}
            var v50 int = (v26 | v26)
            _ = v49
            _ = v50
        }
        v51 := GenS8{F0: 0}
        _ = v51
    }
    if v2 {
        v26 -= (v26 / ((v26) | 1))
        v52 := v4[""]
        _ = v52
    } else {
        if (v2 || v2) {
            var v53 int8 = 0
            var v54 uint16 = 2
            var v55 uint32 = 2
            _ = v53
            _ = v54
            _ = v55
        }
    }
    v56 := GenS5{F0: v26}
    switch "a" {
    case "a":
        if (2 >= 86) {
            if true {
                var v57 float32 = 0.1
                fmt.Println(v57, v26)
                if (true || false) {
                    fmt.Println(v57, v3)
                    switch "a" {
                    case "\n\t":
                        v58 := []int{(v26 * -354), v26, (-(v26))}
                        _ = v58
                    }
                    switch "\n\t" {
                    case "":
                        fmt.Println(v57, v3, v2)
                    case "\n\t":
                        for fi59 := 0; fi59 < 20; fi59++ {
                            var v60 int32 = -1
                            if v2 {
                                var v61 uint16 = 255
                                if (("\n\t" + "") > "a") {
                                    if v2 {
                                        var v62 int8 = -1
                                        _ = v62
                                    } else {
                                        v2 = true
                                        v26 += 127
                                        v61 += (v61 / (((v61 ^ v61)) | 1))
                                    }
                                } else {
                                    fmt.Println(v45, fi59)
                                    v63 := GenS9{F0: (v61 % (((v61 & (v61 & (v61 % ((v61) | 1))))) | 1)), F1: v61}
                                    if ((v61 * 545) <= v61) {
                                        switch "hello" {
                                        case "\n\t":
                                            if ("foo bar" != ("\n\t" + "a")) {
                                                var v64 uint16 = v61
                                                if (-32768 != 1) {
                                                    if v2 {
                                                        fmt.Println(v64, v61)
                                                        v65 := GenS10{F0: 0, F1: 1}
                                                        var v66 int8 = -11
                                                        _ = v65
                                                        _ = v66
                                                    } else {
                                                        if v2 {
                                                            var v67 int32 = v60
                                                            if v2 {
                                                                v68 := GenS11{F0: 0}
                                                                fmt.Println(v64, v45, v26)
                                                                _ = v68
                                                            }
                                                            switch ("\n\t" + "\n\t") {
                                                            case "hello":
                                                                if (v61 < (v61 * (v61 * 2))) {
                                                                    for fi69 := 0; fi69 < 4; fi69++ {
                                                                        v70 := GenS3{F0: 65535}
                                                                        _ = fi69
                                                                        _ = v70
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                for fi71 := 0; fi71 < 11; fi71++ {
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        v2 = v2
                                                                    case "":
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            v72 := GenS12{F0: uint(v45)}
                                                                            _ = v72
                                                                        case "":
                                                                            v57 -= float32(((0.0 * 3.14159) + -2.25))
                                                                        case "\n\t":
                                                                            if (255 == 5) {
                                                                                var v73 uint = 0
                                                                                v74 := GenS5{F0: 775}
                                                                                var v75 uint = (v73 | 1)
                                                                                _ = v73
                                                                                _ = v74
                                                                                _ = v75
                                                                            } else {
                                                                                v26 -= 32767
                                                                                if v2 {
                                                                                    v76 := GenS3{F0: 42}
                                                                                    v3 += -2.25
                                                                                    if false {
                                                                                        var v77 float32 = 1.5
                                                                                        _ = v77
                                                                                    } else {
                                                                                        var v78 float64 = (0.1 / -2.25)
                                                                                        v26 = v26
                                                                                        switch "foo bar" {
                                                                                        case "\n\t":
                                                                                            v79 := GenS13{F0: false, F1: 0}
                                                                                            _ = v79
                                                                                        }
                                                                                        _ = v78
                                                                                    }
                                                                                    _ = v76
                                                                                } else {
                                                                                    v1[797] = (fi71 % ((-9223372036854775808) | 1))
                                                                                }
                                                                                switch "" {
                                                                                case "a":
                                                                                    v3 = ((v57 + v57) - v3)
                                                                                case "hello":
                                                                                    v80 := []int{-794, (v26 % ((fi59) | 1)), fi59}
                                                                                    _ = v80
                                                                                default:
                                                                                    if ((fi59 << ((4294967295) & 63)) == fi59) {
                                                                                        v81 := GenS3{F0: 255}
                                                                                        v2 = ((255 == v64) && (!(1 > 2147483648)))
                                                                                        v82 := []int{(^(v26))}
                                                                                        var v83 int64 = 1
                                                                                        _ = v81
                                                                                        _ = v82
                                                                                        _ = v83
                                                                                    } else {
                                                                                        if (v2 || ((127 >= 214) && v2)) {
                                                                                            var v84 int = fi59
                                                                                            _ = v84
                                                                                        }
                                                                                        v85 := GenS7{F0: (v60 + v60)}
                                                                                        _ = v85
                                                                                    }
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v86 int32 = v67
                                                                            _ = v86
                                                                        }
                                                                        switch -740 {
                                                                        case 1:
                                                                            if v2 {
                                                                                switch "" {
                                                                                case "hello":
                                                                                    v87 := []int{fi59}
                                                                                    _ = v87
                                                                                case "foo bar":
                                                                                    v60 = v60
                                                                                case "":
                                                                                    v88 := GenS3{F0: 0}
                                                                                    _ = v88
                                                                                }
                                                                                v89 := []int{(v26 + (fi59 << ((754) & 63))), (fi71 - (fi59 * fi59)), v26, (v26 * (fi71 - 197))}
                                                                                _ = v89
                                                                            } else {
                                                                                switch ((("a" + "\n\t") + "a") + "hello") {
                                                                                case "\n\t":
                                                                                    v90 := []int{(fi71 | (^(v26)))}
                                                                                    _ = v90
                                                                                default:
                                                                                    switch "\n\t" {
                                                                                    case "hello":
                                                                                        fmt.Println(v60)
                                                                                    default:
                                                                                        v67 = v67
                                                                                    }
                                                                                    v3 = v3
                                                                                    v91 := GenS14{F0: v67}
                                                                                    _ = v91
                                                                                }
                                                                                v61 += uint16(v45)
                                                                            }
                                                                        case -681:
                                                                            switch int(v45) {
                                                                            case 203:
                                                                                v92 := []int{32768}
                                                                                _ = v92
                                                                            }
                                                                        case -2147483648:
                                                                            v45 += (v45 ^ v45)
                                                                        }
                                                                        var v93 string = "hello"
                                                                        switch v93 {
                                                                        case "hello":
                                                                            var v94 bool = v2
                                                                            _ = v94
                                                                        case "a":
                                                                            v95 := []int{fi59, -2147483648}
                                                                            _ = v95
                                                                        default:
                                                                            v2 = ((-1 >= -643) && v2)
                                                                        }
                                                                        v4[(v93 + (v93 + ("" + "a")))] = 2147483648
                                                                        if v2 {
                                                                            var v96 uint8 = 0
                                                                            if (v2 && v2) {
                                                                                var v97 int8 = -1
                                                                                var v98 int16 = 0
                                                                                var v99 bool = true
                                                                                _ = v97
                                                                                _ = v98
                                                                                _ = v99
                                                                            }
                                                                            switch v26 {
                                                                            case -2147483648:
                                                                                v100 := GenS3{F0: 867}
                                                                                _ = v100
                                                                            default:
                                                                                var v101 bool = v2
                                                                                _ = v101
                                                                            }
                                                                            _ = v96
                                                                        } else {
                                                                            v102 := GenS13{F0: (42 == 32767), F1: 42}
                                                                            fmt.Println(v2, v3)
                                                                            var v103 bool = v2
                                                                            _ = v102
                                                                            _ = v103
                                                                        }
                                                                        _ = v93
                                                                    case "a":
                                                                        var v104 int8 = 25
                                                                        _ = v104
                                                                    default:
                                                                        var v105 string = ("hello" + (((("\n\t" + "") + "") + "\n\t") + ""))
                                                                        _ = v105
                                                                    }
                                                                    var v106 uint16 = v64
                                                                    v107 := v4[(("" + "") + "")]
                                                                    _ = fi71
                                                                    _ = v106
                                                                    _ = v107
                                                                }
                                                            }
                                                            _ = v67
                                                        } else {
                                                            for fi108 := 0; fi108 < 19; fi108++ {
                                                                var v109 int64 = -9223372036854775808
                                                                v109 = -320
                                                                v61 -= (v61 >> ((2) & 63))
                                                                _ = fi108
                                                                _ = v109
                                                            }
                                                            for fi110 := 0; fi110 < 1; fi110++ {
                                                                var v111 int64 = 836
                                                                var v112 int16 = 42
                                                                _ = fi110
                                                                _ = v111
                                                                _ = v112
                                                            }
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                v26 -= (v26 % ((fi59) | 1))
                                                            case "":
                                                                v113 := GenS15{F0: 42, F1: 1, F2: (v61 - 1)}
                                                                _ = v113
                                                            default:
                                                                fmt.Println(v60)
                                                            }
                                                        }
                                                    }
                                                    v114 := GenS14{F0: (v60 >> ((uint(v61)) & 63))}
                                                    _ = v114
                                                } else {
                                                    var v115 uint64 = (v45 & v45)
                                                    v2 = ((!(((fi59 << ((65535) & 63)) > -575) || true)) && false)
                                                    _ = v115
                                                }
                                                _ = v64
                                            }
                                        default:
                                            switch "\n\t" {
                                            case "\n\t":
                                                v61 = (v61 ^ 32767)
                                            case "foo bar":
                                                var v116 int32 = v60
                                                _ = v116
                                            case "a":
                                                fmt.Println(v26)
                                            }
                                        }
                                        var v117 uint64 = 0
                                        for fi118 := 0; fi118 < 2; fi118++ {
                                            v119 := GenS14{F0: (v60 - (v60 >> ((18446744073709551615) & 63)))}
                                            v61 += 32767
                                            v120 := GenS5{F0: fi59}
                                            v4["foo bar"] = 2
                                            v121 := GenS6{F0: (uint8(v61) < 42), F1: 0, F2: 65535}
                                            _ = fi118
                                            _ = v119
                                            _ = v120
                                            _ = v121
                                        }
                                        _ = v117
                                    }
                                    _ = v63
                                }
                                if false {
                                    v45 += genFn1(128, 1, ((^(v61)) < v61))
                                }
                                _ = v61
                            } else {
                                v60 += (v60 & 1)
                                for fi122 := 0; fi122 < 9; fi122++ {
                                    fmt.Println(v60, v60)
                                    _ = fi122
                                }
                            }
                            _ = fi59
                            _ = v60
                        }
                    case "foo bar":
                        v45 += (v45 | v45)
                    default:
                        v45 -= (v45 + 1)
                    }
                }
                _ = v57
            } else {
                _ = len(v1)
                if v2 {
                    if true {
                        v2 = false
                        for fi123 := 0; fi123 < 14; fi123++ {
                            var v124 uint32 = 1
                            v125 := GenS5{F0: int(v124)}
                            for fi126 := 0; fi126 < 20; fi126++ {
                                v127 := GenS16{F0: -1, F1: (v45 & (v45 * v45))}
                                if v2 {
                                    switch fi126 {
                                    case 128:
                                        fmt.Println(v124)
                                    case -952:
                                        fmt.Println(v3, v45)
                                    default:
                                        var v128 int32 = -2147483648
                                        _ = v128
                                    }
                                    v26 -= (-(fi126))
                                }
                                var v129 float64 = -2.25
                                _ = fi126
                                _ = v127
                                _ = v129
                            }
                            _ = fi123
                            _ = v124
                            _ = v125
                        }
                        var v130 uint = 0
                        _ = v130
                    }
                    var v131 uint8 = 128
                    _ = v131
                }
            }
            var v132 bool = (v2 || v2)
            _ = v132
        } else {
            var v133 float64 = (0.1 + -2.25)
            v134 := GenS8{F0: 2147483647}
            _ = v133
            _ = v134
        }
    case "\n\t":
        v2 = ((((!v2) || (v45 == v45)) || (32767 >= 0)) && (((0 == -248) && v2) || (!false)))
    default:
        var v135 uint32 = 1
        _ = v135
    }
    v45 += (v45 | 34)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v56
}
