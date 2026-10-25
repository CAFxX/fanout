package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 uint8 }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int16; F1 uint64; F2 string }
type GenS10 struct { F0 bool; F1 int64 }
type GenS11 struct { F0 int64; F1 uint32; F2 int16 }
type GenS12 struct { F0 int64; F1 int32 }
type GenS13 struct { F0 int; F1 int32; F2 int32 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 int8; F1 uint64; F2 int32 }
type GenS18 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 uint64; F1 uint64; F2 string }
func genFn0(p0 string, p1 int16, p2 int32) uint16 {
    fmt.Println(p2, p1)
    if false {
        v6 := GenS1{F0: 641}
        if (true || true) {
            for fi7 := 0; fi7 < 9; fi7++ {
                var v8 uint16 = 42
                switch fi7 {
                case 32768:
                    p2 += (p2 << ((221) & 63))
                }
                var v9 float64 = (1.5 * float64(3.14159))
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        _ = v6
    } else {
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := GenS2{F0: 18446744073709551615, F1: 251}
            p2 = 2147483647
            p0 = "foo bar"
            _ = fi10
            _ = v11
        }
        p0 = p0
    }
    v12 := GenS1{F0: 255}
    v13 := []int{32767}
    v14 := GenS2{F0: 0, F1: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    return 575
}
func genFn1(p0 uint, p1 uint16) uint8 {
    v15 := []int{int(p0), -9223372036854775808, 2147483647, 2147483647}
    var v16 uint16 = 470
    if true {
        v16 = p1
    }
    var v17 bool = (!(1.5 != 1e10))
    for fi18 := 0; fi18 < 5; fi18++ {
        switch "" {
        case "hello":
            if ("foo bar" > "a") {
                fmt.Println(fi18, v16, p0)
            }
        case "foo bar":
            var v19 int64 = -886
            _ = v19
        case "":
            v20 := GenS2{F0: 344, F1: (p1 | p1)}
            _ = v20
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    _ = v17
    return 44
}

func main() {
    v1 := []string{"a", "\n\t", "foo bar", "a"}
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    var v4 uint32 = 4294967295
    var v5 string = "a"
    for fi21 := 0; fi21 < 4; fi21++ {
        switch (v5 + (v5 + (v5 + ("foo bar" + (v5 + v5))))) {
        case "\n\t":
            if false {
                for fi22 := 0; fi22 < 20; fi22++ {
                    var v23 uint64 = 4294967295
                    _ = len(v3)
                    _ = fi22
                    _ = v23
                }
            }
        case "hello":
            if (!true) {
                for fi24 := 0; fi24 < 16; fi24++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v25 := v2[fi24]
                    _ = fi24
                    _ = v25
                }
                v1 = append(v1, v5)
            } else {
                fmt.Println(v4, v4, v5)
                for fi26 := 0; fi26 < 8; fi26++ {
                    if false {
                        if false {
                            if (false || true) {
                                v5 = v5
                            } else {
                                if false {
                                    var v27 int = 9223372036854775807
                                    switch fi26 {
                                    case -2147483648:
                                        fmt.Println(v4, v27)
                                    case -1:
                                        v27 += (fi21 | (v27 | (v27 / (((v27 & (v27 | fi21))) | 1))))
                                    case 721:
                                        v27 -= -2147483648
                                    default:
                                        v28 := GenS3{F0: -122, F1: int(v4)}
                                        _ = v28
                                    }
                                    _ = v27
                                }
                            }
                        } else {
                            if (!true) {
                                if (82 > 1) {
                                    v5 = "hello"
                                    v29 := GenS1{F0: 334}
                                    var v30 int32 = 571
                                    _ = v29
                                    _ = v30
                                } else {
                                    v4 = v4
                                    v31 := GenS1{F0: 0}
                                    _ = v31
                                }
                                var v32 float64 = (-2.25 * (0.0 - 1e10))
                                _ = v32
                            }
                            if (0 >= 0) {
                                fmt.Println(v5, v4)
                                v33 := GenS1{F0: 951}
                                if false {
                                    var v34 int = 128
                                    var v35 int = (v34 ^ v34)
                                    if len(v1) > 0 {
                                        v1[0] = v5
                                    }
                                    _ = v34
                                    _ = v35
                                } else {
                                    v4 = 861
                                }
                                v5 = "foo bar"
                                var v36 uint32 = 42
                                if (!true) {
                                    var v37 int16 = 1
                                    switch (v5 + v5) {
                                    case "hello":
                                        v36 -= (v4 << ((1) & 63))
                                    }
                                    _ = v37
                                }
                                var v38 int64 = 9223372036854775807
                                v39 := GenS1{F0: 65535}
                                _ = v33
                                _ = v36
                                _ = v38
                                _ = v39
                            }
                        }
                        if false {
                            var v40 uint16 = 65535
                            v4 = 4294967295
                            var v41 bool = (true && (-1 > 61))
                            var v42 int32 = 1
                            _ = v40
                            _ = v41
                            _ = v42
                        } else {
                            fmt.Println(fi21, v4)
                            v5 = "hello"
                            v43 := GenS4{F0: -17, F1: 127}
                            _ = v43
                        }
                        var v44 uint32 = 2147483647
                        _ = v44
                    } else {
                        var v45 uint64 = 1
                        if (!false) {
                            if true {
                                var v46 bool = (-9223372036854775808 <= 42)
                                _ = v46
                            } else {
                                switch v5 {
                                case "a":
                                    v47 := GenS5{F0: fi26}
                                    _ = v47
                                }
                                v4 -= (v4 << ((846) & 63))
                            }
                            v4 += v4
                        }
                        _ = v45
                    }
                    if (true || (802 <= 709)) {
                        _ = len(v1)
                        if false {
                            v48 := GenS2{F0: 4294967295, F1: 32767}
                            _ = v48
                        }
                        if false {
                            var v49 float32 = 0.1
                            switch (fi21 | fi26) {
                            case 9:
                                v4 = v4
                            case -1:
                                v50 := GenS6{F0: 89, F1: genFn1(0, 255)}
                                _ = v50
                            case 127:
                                v4 -= (^(v4))
                            }
                            _ = v49
                        } else {
                            fmt.Println(v4, v4)
                            v4 += v4
                        }
                        fmt.Println(v4)
                    } else {
                        switch 80 {
                        case 2147483647:
                            switch "hello" {
                            case "foo bar":
                                v4 = 65535
                            default:
                                if (uint(v4) > 2) {
                                    fmt.Println(v4, v5, fi21)
                                    v5 = v5
                                    v4 -= 597
                                    v51 := GenS7{F0: v4, F1: -239}
                                    _ = v51
                                } else {
                                    v52 := GenS8{F0: 1}
                                    _ = v52
                                }
                            }
                        case 9223372036854775807:
                            v5 = (((v5 + ((v5 + v5) + "hello")) + "hello") + "hello")
                            var v53 uint16 = 33
                            delete(v2, int(v4))
                            _ = v53
                        case 32767:
                            fmt.Println(fi21)
                        }
                    }
                    _ = fi26
                }
            }
        default:
            for fi54 := 0; fi54 < 20; fi54++ {
                if ((^(v4)) != v4) {
                    var v55 uint32 = (v4 ^ (^(v4)))
                    v56 := GenS1{F0: 36}
                    v5 = v5
                    _ = v55
                    _ = v56
                } else {
                    if (-2.25 < 0.0) {
                        var v57 int32 = 42
                        _ = v57
                    } else {
                        switch ((v5 + v5) + (v5 + (((v5 + "foo bar") + v5) + ((v5 + (v5 + (("foo bar" + (((("hello" + "hello") + v5) + "a") + (v5 + (((v5 + "foo bar") + "") + "hello")))) + ""))) + v5)))) {
                        case "hello":
                            if (uint64(v4) <= 42) {
                                var v58 string = (v5 + v5)
                                _ = v58
                            }
                        case "\n\t":
                            fmt.Println(fi21)
                        default:
                            var v59 bool = ("foo bar" >= "a")
                            _ = v59
                        }
                    }
                }
                if false {
                    var v60 uint = 65535
                    var v61 bool = ((false || (!false)) || false)
                    _ = v60
                    _ = v61
                } else {
                    var v62 uint8 = 128
                    var v63 uint16 = 221
                    _ = v62
                    _ = v63
                }
                _ = fi54
            }
        }
        _ = fi21
    }
    for fi64 := 0; fi64 < 14; fi64++ {
        if (!(false || true)) {
            v65 := GenS3{F0: 50, F1: 791}
            v5 = ("a" + "")
            _ = v65
        }
        _ = fi64
    }
    v5 = v5
    v4 = v4
    var v66 uint16 = 0
    fmt.Println(v4)
    if true {
        v67 := GenS5{F0: 1}
        fmt.Println(v4, v66)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v67
    }
    v68 := GenS9{F0: -102, F1: 2, F2: "foo bar"}
    switch -1 {
    case 127:
        v4 -= 257
    case 42:
        v69 := GenS10{F0: false, F1: 1}
        _ = v69
    default:
        var v70 int16 = -565
        _ = v70
    }
    _ = len(v3)
    for fi71 := 0; fi71 < 15; fi71++ {
        for fi72 := 0; fi72 < 5; fi72++ {
            switch (fi72 + 128) {
            case 32768:
                v73 := GenS11{F0: -601, F1: v4, F2: -620}
                _ = v73
            case 4294967295:
                var v74 uint = 18446744073709551615
                _ = v74
            case 42:
                if false {
                    if false {
                        if true {
                            v75 := GenS12{F0: 0, F1: -1}
                            var v76 string = v5
                            _ = v75
                            _ = v76
                        } else {
                            if (0 < 2) {
                                switch v5 {
                                case "a":
                                    delete(v2, -1)
                                    v77 := GenS4{F0: 0, F1: 127}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v5 = v5
                                    v66 -= v66
                                    _ = v77
                                case "":
                                    var v78 uint32 = v4
                                    _ = v78
                                case "\n\t":
                                    v5 = "foo bar"
                                    var v79 int16 = 42
                                    var v80 int64 = 9223372036854775807
                                    if true {
                                        v81 := v3["hello"]
                                        switch ((((v5 + (v5 + "foo bar")) + v5) + v5) + v5) {
                                        case "hello":
                                            if true {
                                                var v82 int = fi71
                                                if ((v80 - 9223372036854775807) >= 0) {
                                                    fmt.Println(v81, v80, v5)
                                                    if true {
                                                        v80 += (^(v80))
                                                        v5 = (("a" + v5) + v5)
                                                        v4 += v81
                                                    } else {
                                                        v82 = (fi71 - fi71)
                                                        if (217 < 0) {
                                                            var v83 int8 = 127
                                                            var v84 int32 = 2147483647
                                                            var v85 float32 = 0.1
                                                            _ = v83
                                                            _ = v84
                                                            _ = v85
                                                        } else {
                                                            fmt.Println(v79, v80, v82)
                                                            switch v5 {
                                                            case "foo bar":
                                                                v82 += fi71
                                                            case "hello":
                                                                fmt.Println(v80)
                                                            }
                                                            if (1 != -128) {
                                                                var v86 uint = 4294967295
                                                                _ = v86
                                                            } else {
                                                                var v87 float64 = 0.1
                                                                _ = v87
                                                            }
                                                            var v88 uint32 = (^(v81))
                                                            v80 = v80
                                                            _ = v88
                                                        }
                                                        v4 = (^(v4))
                                                    }
                                                    if true {
                                                        switch (fi71 & fi71) {
                                                        case 128:
                                                            var v89 uint = 0
                                                            _ = v89
                                                        case -9223372036854775808:
                                                            var v90 bool = false
                                                            _ = v90
                                                        case 1:
                                                            switch (fi71 | 896) {
                                                            case 128:
                                                                v91 := GenS13{F0: v82, F1: -867, F2: -289}
                                                                _ = v91
                                                            case -762:
                                                                fmt.Println(v66, v66)
                                                            default:
                                                                if false {
                                                                    var v92 int16 = 0
                                                                    var v93 uint64 = 4294967295
                                                                    fmt.Println(fi72)
                                                                    _ = v92
                                                                    _ = v93
                                                                } else {
                                                                    var v94 int8 = 0
                                                                    v66 += 818
                                                                    v95 := GenS6{F0: 85, F1: 42}
                                                                    _ = v94
                                                                    _ = v95
                                                                }
                                                            }
                                                        }
                                                        if (1.5 <= -2.25) {
                                                            if true {
                                                                v80 = 0
                                                                var v96 int32 = 1
                                                                v1 = append(v1, ((v5 + "hello") + v5))
                                                                if (v79 >= v79) {
                                                                    _ = len(v1)
                                                                    v97 := GenS9{F0: 42, F1: 4294967295, F2: ((("foo bar" + v5) + v5) + v5)}
                                                                    _ = v97
                                                                }
                                                                fmt.Println(fi72)
                                                                var v98 uint64 = 584
                                                                var v99 int32 = (v96 << ((255) & 63))
                                                                _ = v96
                                                                _ = v98
                                                                _ = v99
                                                            } else {
                                                                var v100 int8 = 0
                                                                _ = v100
                                                            }
                                                            fmt.Println(fi72, v5, fi72)
                                                            v82 = (v82 + 32768)
                                                        } else {
                                                            v79 += (v79 << ((65535) & 63))
                                                            v101 := GenS14{F0: 42}
                                                            _ = v101
                                                        }
                                                    } else {
                                                        v4 -= (v4 ^ 0)
                                                        switch v5 {
                                                        case "\n\t":
                                                            if (genFn1(671, v66) >= 255) {
                                                                v4 += 42
                                                                var v102 uint16 = (v66 - uint16(v4))
                                                                v103 := GenS15{F0: true}
                                                                _ = v102
                                                                _ = v103
                                                            } else {
                                                                if true {
                                                                    var v104 uint32 = 42
                                                                    v105 := GenS11{F0: (v80 * 108), F1: 42, F2: v79}
                                                                    v104 = 2
                                                                    var v106 int64 = 649
                                                                    v107 := GenS16{F0: 0}
                                                                    v82 = fi72
                                                                    var v108 int64 = (v80 << ((0) & 63))
                                                                    _ = v104
                                                                    _ = v105
                                                                    _ = v106
                                                                    _ = v107
                                                                    _ = v108
                                                                } else {
                                                                    v109 := v3["a"]
                                                                    v81 = (v81 + (v109 * v81))
                                                                    if true {
                                                                        v4 -= (^(v109))
                                                                    }
                                                                    var v110 int16 = (-(v79))
                                                                    _ = v109
                                                                    _ = v110
                                                                }
                                                                v111 := GenS11{F0: (v80 % ((v80) | 1)), F1: (v4 / (((v4 >> ((166) & 63))) | 1)), F2: v79}
                                                                _ = v111
                                                            }
                                                        case "foo bar":
                                                            if len(v1) > 0 {
                                                                v1[0] = "\n\t"
                                                            }
                                                        case "hello":
                                                            fmt.Println(fi72, v81)
                                                        default:
                                                            fmt.Println(v81)
                                                        }
                                                    }
                                                }
                                                v112 := GenS1{F0: (v66 | (v66 - v66))}
                                                if false {
                                                    if true {
                                                        if false {
                                                            _ = len(v2)
                                                            var v113 uint32 = (v4 % ((687) | 1))
                                                            var v114 float64 = -2.25
                                                            fmt.Println(v80, v114)
                                                            _ = v113
                                                            _ = v114
                                                        }
                                                        v115 := GenS10{F0: true, F1: (v80 / (((v80 - v80)) | 1))}
                                                        var v116 int8 = -1
                                                        switch (^(v82)) {
                                                        case 2147483647:
                                                            var v117 int16 = v79
                                                            _ = v117
                                                        }
                                                        _ = v115
                                                        _ = v116
                                                    } else {
                                                        delete(v2, int(v79))
                                                        delete(v2, -2147483648)
                                                        var v118 uint16 = 255
                                                        _ = v118
                                                    }
                                                    var v119 int32 = -541
                                                    v120 := GenS17{F0: 1, F1: 1, F2: v119}
                                                    _ = v119
                                                    _ = v120
                                                }
                                                switch fi72 {
                                                case 812:
                                                    v2[(fi72 << ((977) & 63))] = (v81 % ((v4) | 1))
                                                default:
                                                    var v121 uint16 = (^(v66))
                                                    _ = v121
                                                }
                                                _ = v112
                                                _ = v82
                                            } else {
                                                var v122 string = (v5 + v5)
                                                fmt.Println(v79, v5, v79)
                                                _ = v122
                                            }
                                            var v123 int = fi72
                                            _ = v123
                                        case "foo bar":
                                            if false {
                                                _ = len(v1)
                                            } else {
                                                if true {
                                                    v81 -= v4
                                                }
                                                var v124 uint32 = v81
                                                _ = v124
                                            }
                                        case "a":
                                            if ("" >= v5) {
                                                v66 = 227
                                                var v125 int8 = -9
                                                _ = v125
                                            }
                                        }
                                        var v126 int16 = (v79 >> ((2) & 63))
                                        _ = v126
                                        _ = v81
                                    }
                                    _ = v79
                                    _ = v80
                                }
                                v127 := GenS18{F0: 255, F1: 2, F2: genFn0((v5 + v5), 32767, 0)}
                                _ = v127
                            }
                        }
                        if true {
                            v128 := GenS12{F0: -244, F1: 596}
                            var v129 int32 = 54
                            _ = v128
                            _ = v129
                        } else {
                            v4 = v4
                        }
                        _ = len(v1)
                    } else {
                        v130 := GenS5{F0: fi71}
                        _ = v130
                    }
                    v66 -= v66
                    if (0 < -128) {
                        v5 = v5
                    } else {
                        v131 := v2[(fi71 >> ((7) & 63))]
                        v132 := GenS19{F0: v4}
                        v4 -= v4
                        _ = v131
                        _ = v132
                    }
                    v133 := GenS7{F0: (v4 % (((v4 ^ v4)) | 1)), F1: -182}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v133
                } else {
                    if (!(!false)) {
                        v134 := GenS1{F0: v66}
                        v66 = v66
                        var v135 int64 = 42
                        _ = len(v2)
                        _ = v134
                        _ = v135
                    }
                    switch v5 {
                    case "\n\t":
                        switch 32768 {
                        case -346:
                            v136 := GenS4{F0: 1, F1: 182}
                            var v137 float32 = float32(((0.1 - 1.5) + (0.0 / 0.1)))
                            v138 := v3[v5]
                            if true {
                                switch v5 {
                                case "hello":
                                    var v139 float64 = 1.5
                                    _ = v139
                                case "foo bar":
                                    var v140 float64 = 1.5
                                    _ = v140
                                case "a":
                                    v138 = v138
                                default:
                                    fmt.Println(v138, fi72)
                                }
                                delete(v2, fi72)
                            }
                            v1 = append(v1, "hello")
                            _ = v136
                            _ = v137
                            _ = v138
                        default:
                            var v141 uint8 = 255
                            _ = v141
                        }
                    default:
                        fmt.Println(v5, fi72)
                    }
                }
            }
            _ = fi72
        }
        v142 := GenS13{F0: (-(fi71)), F1: -1, F2: 290}
        if ((!false) || true) {
            v4 = v4
            for fi143 := 0; fi143 < 3; fi143++ {
                if true {
                    v66 -= (v66 << ((2) & 63))
                } else {
                    v4 -= (v4 >> ((1) & 63))
                    var v144 uint64 = 737
                    v66 = 311
                    _ = v144
                }
                v145 := GenS9{F0: 1000, F1: 1, F2: v5}
                _ = fi143
                _ = v145
            }
            if true {
                v66 = v66
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v66)
            }
        } else {
            var v146 bool = (312 <= 1)
            v147 := GenS12{F0: 33, F1: 2147483647}
            _ = v146
            _ = v147
        }
        _ = fi71
        _ = v142
    }
    switch -9223372036854775808 {
    case 261:
        var v148 float64 = 0.0
        _ = v148
    case 720:
        fmt.Println(v66, v4, v4)
    }
    v149 := GenS20{F0: 4294967295, F1: 4294967295, F2: v5}
    var v150 int = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v150:", v150)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v149
    _ = v150
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v68
}
