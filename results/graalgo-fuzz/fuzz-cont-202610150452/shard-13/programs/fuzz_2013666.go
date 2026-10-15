package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 int64; F1 int16; F2 int16 }
type GenS8 struct { F0 int64; F1 string }
type GenS9 struct { F0 int; F1 int64 }
type GenS10 struct { F0 uint32; F1 int }
type GenS11 struct { F0 int64; F1 string; F2 int32 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 string; F1 int32; F2 uint32 }
type GenS14 struct { F0 int32; F1 int16; F2 int8 }
type GenS15 struct { F0 uint64; F1 int8; F2 int64 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 2
    var v3 int64 = 387
    v4 := make(map[uint32]uint32)
    var v5 int64 = (v3 & v3)
    var v6 int8 = -38
    v5 += (v5 | (v3 - v3))
    if true {
        if (127 < 42) {
            var v7 uint = 1
            fmt.Println(v2, v5)
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 float64 = (0.0 + -2.25)
                _ = fi8
                _ = v9
            }
            _ = v7
        } else {
            _ = len(v4)
            fmt.Println(v3)
            v10 := GenS1{F0: 18446744073709551615, F1: 1000}
            _ = v10
        }
    } else {
        for fi11 := 0; fi11 < 13; fi11++ {
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 uint16 = 2
                var v14 int = (fi11 ^ fi11)
                _ = fi12
                _ = v13
                _ = v14
            }
            switch (fi11 % ((128) | 1)) {
            case -415:
                v5 = (v5 << ((1) & 63))
            default:
                v15 := []int{}
                _ = v15
            }
            v16 := GenS1{F0: 2, F1: 0}
            _ = fi11
            _ = v16
        }
    }
    var v17 string = "foo bar"
    if (1e10 <= float64(0.1)) {
        var v18 bool = true
        fmt.Println(v17)
        _ = v18
    }
    if (1 > 42) {
        for fi19 := 0; fi19 < 15; fi19++ {
            switch v17 {
            case "foo bar":
                v2 -= (^(v2))
            case "hello":
                for fi20 := 0; fi20 < 6; fi20++ {
                    if (0 != 73) {
                        if (42 >= -2147483648) {
                            fmt.Println(v2)
                            v21 := v1[42]
                            _ = v21
                        } else {
                            if (true || ("hello" <= v17)) {
                                v5 = 9223372036854775807
                                var v22 uint64 = 4294967295
                                switch ((v17 + v17) + v17) {
                                case "a":
                                    var v23 int = (fi19 % ((fi20) | 1))
                                    _ = v23
                                case "":
                                    if (true || true) {
                                        v24 := GenS2{F0: 1000, F1: 244, F2: 645}
                                        v25 := GenS3{F0: -128, F1: fi20}
                                        fmt.Println(fi20)
                                        if (true || (!(!(true && false)))) {
                                            var v26 uint64 = v22
                                            _ = v26
                                        } else {
                                            if (3.14159 != 0.0) {
                                                var v27 int16 = int16(fi20)
                                                if (v22 >= v22) {
                                                    v28 := GenS4{F0: (286 == 255)}
                                                    fmt.Println(v6, v3)
                                                    fmt.Println(v17, v22)
                                                    _ = v28
                                                }
                                                var v29 string = (v17 + "\n\t")
                                                _ = v27
                                                _ = v29
                                            }
                                        }
                                        v2 -= v2
                                        _ = v24
                                        _ = v25
                                    } else {
                                        v6 -= v6
                                        v3 = (v5 + (v5 & (-(v5))))
                                        var v30 string = ("\n\t" + (v17 + ((v17 + v17) + ("foo bar" + v17))))
                                        _ = v30
                                    }
                                default:
                                    var v31 uint32 = 652
                                    _ = v31
                                }
                                _ = v22
                            } else {
                                v5 = v5
                                v32 := []int{fi20, -239, (fi20 & fi20)}
                                if len(v32) >= 2 {
                                    v32 = v32[1:len(v32)]
                                }
                                var v33 int32 = -2147483648
                                var v34 uint = 0
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                            v17 = "hello"
                            if (true && true) {
                                v35 := GenS4{F0: false}
                                fmt.Println(v3, v3, v5)
                                _ = v35
                            } else {
                                var v36 float32 = 1e10
                                fmt.Println(v6, v6, v6)
                                var v37 int32 = 42
                                _ = v36
                                _ = v37
                            }
                            if true {
                                delete(v1, 112)
                                if (-2147483648 > 2147483647) {
                                    fmt.Println(v17, v6, fi20)
                                } else {
                                    if (-437 <= -1) {
                                        v38 := v4[335]
                                        _ = v38
                                    }
                                    v39 := GenS3{F0: (v6 << ((1) & 63)), F1: fi20}
                                    if false {
                                        if false {
                                            v6 -= 22
                                        } else {
                                            v17 = "\n\t"
                                            fmt.Println(fi19, v3)
                                            if true {
                                                var v40 uint32 = 65535
                                                _ = v40
                                            } else {
                                                switch "a" {
                                                case "\n\t":
                                                    v17 = v17
                                                case "a":
                                                    if (true || false) {
                                                        var v41 float64 = (3.14159 - (-2.25 / 3.14159))
                                                        var v42 int16 = -1
                                                        var v43 string = v17
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                    } else {
                                                        switch "" {
                                                        case "":
                                                            if ((false || true) || (755 > 126)) {
                                                                var v44 int16 = 1
                                                                fmt.Println(v44, fi20)
                                                                switch (fi19 * 32767) {
                                                                case 4294967295:
                                                                    _ = len(v1)
                                                                case 2147483647:
                                                                    var v45 uint16 = 2
                                                                    _ = v45
                                                                }
                                                                _ = v44
                                                            } else {
                                                                if (!true) {
                                                                    var v46 float64 = (1.5 + 3.14159)
                                                                    _ = v46
                                                                }
                                                                if (!("a" <= v17)) {
                                                                    fmt.Println(v2, v5)
                                                                    v5 -= v5
                                                                    v47 := GenS2{F0: -516, F1: 42, F2: 65535}
                                                                    _ = v47
                                                                } else {
                                                                    v3 -= v5
                                                                    var v48 int16 = 329
                                                                    fmt.Println(fi20)
                                                                    if true {
                                                                        var v49 uint = 1
                                                                        if false {
                                                                            v50 := []int{}
                                                                            _ = v50
                                                                        }
                                                                        _ = v49
                                                                    }
                                                                    v51 := GenS4{F0: (v17 != ((v17 + v17) + v17))}
                                                                    _ = v48
                                                                    _ = v51
                                                                }
                                                                delete(v1, 2147483648)
                                                            }
                                                            var v52 uint8 = (v2 | v2)
                                                            v17 = (v17 + (v17 + ""))
                                                            fmt.Println(v17, v6, v3)
                                                            _ = v52
                                                        case "a":
                                                            v5 = (v5 / ((v5) | 1))
                                                        case "foo bar":
                                                            switch (("\n\t" + "hello") + (("" + (v17 + v17)) + "foo bar")) {
                                                            case "foo bar":
                                                                v4[0] = 832
                                                            case "":
                                                                v6 = v6
                                                            default:
                                                                if true {
                                                                    v53 := []int{fi19}
                                                                    if (v17 == (v17 + v17)) {
                                                                        v3 += 1
                                                                        v5 = v5
                                                                        var v54 int32 = -767
                                                                        _ = v54
                                                                    } else {
                                                                        var v55 int8 = v6
                                                                        if (!false) {
                                                                            var v56 uint8 = v2
                                                                            _ = v56
                                                                        }
                                                                        v57 := GenS1{F0: 1, F1: 0}
                                                                        _ = v55
                                                                        _ = v57
                                                                    }
                                                                    fmt.Println(fi19, v17, v5)
                                                                    _ = v53
                                                                }
                                                            }
                                                        }
                                                        v58 := GenS3{F0: (-(v6)), F1: fi19}
                                                        _ = v58
                                                    }
                                                case "foo bar":
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v59 int32 = -2147483648
                                                        v60 := GenS5{F0: 420}
                                                        _ = v59
                                                        _ = v60
                                                    case "\n\t":
                                                        v17 = (v17 + (v17 + "hello"))
                                                        if false {
                                                            var v61 int = fi20
                                                            _ = v61
                                                        } else {
                                                            var v62 int64 = 1
                                                            var v63 int8 = 100
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                        fmt.Println(v5, v5, v6)
                                                        v2 = v2
                                                        if (-2.25 <= (0.0 / 1.0)) {
                                                            var v64 int64 = (v5 * v5)
                                                            switch "\n\t" {
                                                            case "a":
                                                                var v65 uint8 = 2
                                                                _ = v65
                                                            case "foo bar":
                                                                v5 -= int64(fi19)
                                                            default:
                                                                v66 := v4[0]
                                                                _ = v66
                                                            }
                                                            if true {
                                                                v64 = int64(fi20)
                                                                if true {
                                                                    if (0 == v2) {
                                                                        if ((v6 * v6) < v6) {
                                                                            v67 := []int{800, (fi19 - (fi19 | (fi19 >> ((1) & 63)))), (fi20 * (fi20 << ((0) & 63))), fi19}
                                                                            var v68 int16 = -361
                                                                            _ = len(v67)
                                                                            v69 := v1[4294967295]
                                                                            if (152 >= int64(v68)) {
                                                                                v64 += (v64 + (v5 + v3))
                                                                            }
                                                                            _ = v67
                                                                            _ = v68
                                                                            _ = v69
                                                                        }
                                                                        var v70 uint32 = 1
                                                                        var v71 uint64 = 2
                                                                        _ = v70
                                                                        _ = v71
                                                                    }
                                                                } else {
                                                                    v72 := GenS4{F0: false}
                                                                    _ = len(v1)
                                                                    v73 := []int{}
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                                if false {
                                                                    v74 := GenS6{F0: 0, F1: true}
                                                                    if false {
                                                                        if (65535 > 2147483648) {
                                                                            v75 := []int{(fi19 >> ((2) & 63)), (fi20 / ((fi20) | 1))}
                                                                            _ = v75
                                                                        } else {
                                                                            if (!(false || (0 >= 42))) {
                                                                                fmt.Println(v64, fi20)
                                                                                fmt.Println(v64, v3)
                                                                                v4[968] = 42
                                                                            } else {
                                                                                v2 = (v2 * 2)
                                                                                v3 = v3
                                                                            }
                                                                        }
                                                                        var v76 bool = false
                                                                        _ = v76
                                                                    } else {
                                                                        v77 := GenS2{F0: -1, F1: uint32(v2), F2: 32767}
                                                                        fmt.Println(v17)
                                                                        if (-410 == (fi20 % ((276) | 1))) {
                                                                            v78 := GenS7{F0: -9223372036854775808, F1: -32768, F2: 32767}
                                                                            if true {
                                                                                var v79 int64 = 1
                                                                                var v80 uint8 = (v2 >> ((2) & 63))
                                                                                var v81 int32 = 1
                                                                                v79 += v5
                                                                                v2 += v2
                                                                                v5 = v5
                                                                                if (!false) {
                                                                                    fmt.Println(v6)
                                                                                    _ = len(v1)
                                                                                    v82 := GenS8{F0: v64, F1: (v17 + ((v17 + "hello") + v17))}
                                                                                    _ = v82
                                                                                } else {
                                                                                    switch fi19 {
                                                                                    case 4294967295:
                                                                                        v83 := v4[905]
                                                                                        _ = v83
                                                                                    case -1:
                                                                                        if (((1.5 * 0.0) + float64(1.5)) == (1.5 * 3.14159)) {
                                                                                            v84 := []int{int(v6), (fi19 * (fi20 & (fi19 - (fi20 >> ((65535) & 63))))), (fi20 / (((-(fi19))) | 1))}
                                                                                            if (!((-2.25 + -2.25) == 0.0)) {
                                                                                                fmt.Println(v2, v17, fi20)
                                                                                                v85 := GenS2{F0: 32767, F1: 1, F2: 0}
                                                                                                fmt.Println(v80, v17)
                                                                                                _ = v85
                                                                                            } else {
                                                                                                var v86 uint64 = 212
                                                                                                var v87 int64 = (^(v79))
                                                                                                fmt.Println(fi20, v79, v80)
                                                                                                _ = v86
                                                                                                _ = v87
                                                                                            }
                                                                                            var v88 bool = ((v5 / ((v79) | 1)) < (-(v79)))
                                                                                            _ = v84
                                                                                            _ = v88
                                                                                        } else {
                                                                                            fmt.Println(v64, v3, v2)
                                                                                            var v89 uint8 = v2
                                                                                            _ = v89
                                                                                        }
                                                                                    }
                                                                                    if true {
                                                                                        var v90 float32 = 0.1
                                                                                        var v91 string = "foo bar"
                                                                                        v81 -= int32(v64)
                                                                                        delete(v1, 65535)
                                                                                        v17 = ("foo bar" + v91)
                                                                                        _ = v90
                                                                                        _ = v91
                                                                                    } else {
                                                                                        v80 = (v80 + 127)
                                                                                        v92 := GenS9{F0: fi20, F1: (v5 >> ((398) & 63))}
                                                                                        v79 -= v5
                                                                                        var v93 int16 = 1
                                                                                        _ = v92
                                                                                        _ = v93
                                                                                    }
                                                                                    var v94 bool = (false || false)
                                                                                    _ = v94
                                                                                }
                                                                                v2 = (v80 % (((v2 % ((v2) | 1))) | 1))
                                                                                _ = v79
                                                                                _ = v80
                                                                                _ = v81
                                                                            }
                                                                            var v95 int64 = v5
                                                                            var v96 uint = 42
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                var v97 uint16 = 32767
                                                                                _ = v97
                                                                            case "\n\t":
                                                                                if false {
                                                                                    if (true || (true && ((uint32(v96) >= 205) || (!true)))) {
                                                                                        v98 := GenS10{F0: 42, F1: fi19}
                                                                                        _ = v98
                                                                                    } else {
                                                                                        v2 -= 96
                                                                                        v64 = (v95 % (((v64 | (v95 % ((v95) | 1)))) | 1))
                                                                                    }
                                                                                    var v99 string = (v17 + "foo bar")
                                                                                    _ = v99
                                                                                } else {
                                                                                    v100 := GenS11{F0: v3, F1: "", F2: 1}
                                                                                    _ = v100
                                                                                }
                                                                            case "a":
                                                                                v101 := GenS7{F0: (v3 - (v64 << ((v96) & 63))), F1: -32768, F2: 913}
                                                                                _ = v101
                                                                            }
                                                                            v102 := GenS2{F0: -1, F1: 2147483647, F2: 32767}
                                                                            _ = v102
                                                                            _ = v78
                                                                            _ = v95
                                                                            _ = v96
                                                                        }
                                                                        _ = v77
                                                                    }
                                                                    _ = v74
                                                                }
                                                                fmt.Println(fi19, fi19)
                                                            } else {
                                                                v103 := GenS5{F0: 2}
                                                                v104 := GenS12{F0: (v17 + "")}
                                                                v105 := v1[252]
                                                                switch fi19 {
                                                                case -2147483648:
                                                                    v106 := GenS8{F0: (v3 << ((717) & 63)), F1: v17}
                                                                    _ = v106
                                                                case -484:
                                                                    v6 = -44
                                                                case 186:
                                                                    v107 := GenS5{F0: 394}
                                                                    _ = v107
                                                                }
                                                                var v108 bool = true
                                                                _ = v103
                                                                _ = v104
                                                                _ = v105
                                                                _ = v108
                                                            }
                                                            _ = v64
                                                        } else {
                                                            v109 := GenS3{F0: v6, F1: fi19}
                                                            var v110 int16 = -1
                                                            var v111 float32 = 1e10
                                                            _ = v109
                                                            _ = v110
                                                            _ = v111
                                                        }
                                                    }
                                                default:
                                                    v6 += v6
                                                }
                                                v5 = v3
                                            }
                                        }
                                    } else {
                                        fmt.Println(v6, fi20, fi20)
                                        v6 = v6
                                    }
                                    _ = v39
                                }
                                if true {
                                    var v112 float64 = 1.5
                                    fmt.Println(v112, fi19, fi19)
                                    v112 = float64((float32((1e10 * ((1.5 + float64((-2.25 - 3.14159))) * float64(0.1)))) - 0.0))
                                    v113 := GenS13{F0: "\n\t", F1: 0, F2: 1}
                                    switch v17 {
                                    case "hello":
                                        delete(v4, 345)
                                    case "a":
                                        v114 := GenS14{F0: 42, F1: -1, F2: -74}
                                        _ = v114
                                    default:
                                        var v115 bool = (42 > 2147483647)
                                        _ = v115
                                    }
                                    var v116 uint64 = 1
                                    _ = v112
                                    _ = v113
                                    _ = v116
                                }
                            } else {
                                v117 := v1[0]
                                _ = len(v1)
                                _ = v117
                            }
                            if true {
                                v5 = -9223372036854775808
                                var v118 uint = uint(v2)
                                v2 += v2
                                _ = v118
                            }
                            if (18446744073709551615 >= 659) {
                                if (711 > 2147483648) {
                                    v2 = v2
                                } else {
                                    if false {
                                        fmt.Println(v5)
                                        if false {
                                            var v119 int64 = v3
                                            _ = v119
                                        } else {
                                            v17 = ((v17 + ((v17 + v17) + (("hello" + v17) + "a"))) + v17)
                                            if ((!true) && ((v6 + v6) > v6)) {
                                                fmt.Println(v17, v3, v6)
                                            }
                                        }
                                        var v120 uint = 433
                                        v121 := []int{(fi19 * (fi19 << ((v120) & 63)))}
                                        _ = v120
                                        _ = v121
                                    }
                                    var v122 string = v17
                                    _ = v122
                                }
                                v123 := GenS15{F0: 0, F1: (v6 - 1), F2: v3}
                                _ = v123
                            }
                        }
                        switch v17 {
                        case "":
                            if false {
                                var v124 int32 = 994
                                v124 = v124
                                _ = v124
                            } else {
                                switch v17 {
                                case "hello":
                                    v3 = v3
                                }
                                switch v17 {
                                case "foo bar":
                                    var v125 string = v17
                                    _ = v125
                                }
                            }
                        default:
                            if (uint32(v2) < 497) {
                                var v126 bool = (!(!true))
                                var v127 uint32 = 835
                                _ = v126
                                _ = v127
                            }
                        }
                    } else {
                        if ((("hello" + v17) + (v17 + "a")) <= v17) {
                            var v128 float64 = ((0.0 / (float64(0.1) + 0.0)) / 1.5)
                            _ = v128
                        }
                        var v129 uint32 = 2147483647
                        _ = v129
                    }
                    switch (fi20 % (((fi20 >> ((1) & 63))) | 1)) {
                    case 667:
                        v17 = v17
                    default:
                        var v130 uint8 = (v2 << ((255) & 63))
                        _ = v130
                    }
                    var v131 int32 = 2147483647
                    _ = fi20
                    _ = v131
                }
            case "":
                var v132 int = fi19
                _ = v132
            }
            _ = fi19
        }
        if (255 < 505) {
            fmt.Println(v2, v6, v3)
        }
        switch v17 {
        case "foo bar":
            v6 = v6
        case "hello":
            switch -84 {
            case 186:
                var v133 uint8 = 255
                _ = v133
            case -386:
                v134 := GenS1{F0: 413, F1: 1000}
                _ = v134
            case 2147483647:
                v135 := GenS1{F0: 1, F1: -32768}
                _ = v135
            }
        case "a":
            var v136 bool = false
            _ = v136
        default:
            v137 := GenS3{F0: 0, F1: 2147483647}
            _ = v137
        }
    } else {
        v138 := GenS4{F0: (!false)}
        _ = v138
    }
    switch -1 {
    case 32768:
        var v139 uint8 = v2
        _ = v139
    default:
        var v140 uint = 1
        _ = v140
    }
    v141 := GenS16{F0: (v3 & (v5 - int64(v6)))}
    if ((^(v3)) >= 928) {
        for fi142 := 0; fi142 < 6; fi142++ {
            for fi143 := 0; fi143 < 19; fi143++ {
                switch v17 {
                case "a":
                    switch (^(fi142)) {
                    case 681:
                        v6 = v6
                    case -2147483648:
                        v2 = (v2 * v2)
                    default:
                        var v144 float64 = (-2.25 / 1e10)
                        _ = v144
                    }
                default:
                    v145 := GenS1{F0: 2, F1: 1000}
                    _ = v145
                }
                _ = fi143
            }
            v3 = v5
            v3 -= v3
            _ = fi142
        }
        switch 32768 {
        case 9223372036854775807:
            fmt.Println(v17)
        default:
            v146 := GenS17{F0: 18446744073709551615}
            _ = v146
        }
    }
    v6 += v6
    v147 := GenS17{F0: 190}
    v6 = (-(v6))
    if ((("a" + v17) + "foo bar") < "a") {
        var v148 float64 = 1e10
        fmt.Println(v5, v3)
        for fi149 := 0; fi149 < 20; fi149++ {
            for fi150 := 0; fi150 < 16; fi150++ {
                var v151 uint8 = 255
                fmt.Println(v3)
                _ = fi150
                _ = v151
            }
            if (175 >= 1) {
                if (42 >= 4294967295) {
                    switch v17 {
                    case "\n\t":
                        var v152 uint8 = 128
                        _ = v152
                    default:
                        v153 := GenS4{F0: false}
                        _ = v153
                    }
                    if true {
                        var v154 uint16 = 1
                        if (4294967295 <= 707) {
                            switch (-(fi149)) {
                            case 4294967295:
                                var v155 int16 = 42
                                _ = v155
                            case 798:
                                v5 = v5
                            }
                        } else {
                            v6 = 42
                        }
                        _ = v154
                    }
                    var v156 int8 = (v6 / ((v6) | 1))
                    _ = v156
                }
            }
            _ = fi149
        }
        _ = v148
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v141
    _ = v147
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
