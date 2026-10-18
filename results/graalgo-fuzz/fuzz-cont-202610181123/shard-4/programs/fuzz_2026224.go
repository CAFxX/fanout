package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint16; F1 string; F2 string }
type GenS5 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int16; F1 int16; F2 int32 }
type GenS9 struct { F0 uint; F1 uint32; F2 int64 }
type GenS10 struct { F0 int64; F1 int }
type GenS11 struct { F0 int8; F1 bool; F2 uint }
type GenS12 struct { F0 uint; F1 int32 }
type GenS13 struct { F0 int16; F1 uint64 }
type GenS14 struct { F0 uint8; F1 uint16 }
type GenS15 struct { F0 uint8; F1 int32; F2 int8 }
type GenS16 struct { F0 uint; F1 uint8 }
type GenS17 struct { F0 int8; F1 string; F2 int32 }
type GenS18 struct { F0 int; F1 int16; F2 int16 }
type GenS19 struct { F0 int8; F1 uint8 }
type GenS20 struct { F0 int32 }
type GenS21 struct { F0 string; F1 int8 }
type GenS22 struct { F0 int64 }
type GenS23 struct { F0 uint32; F1 uint32 }
type GenS24 struct { F0 uint8; F1 string }
type GenS25 struct { F0 int32; F1 uint16; F2 int64 }
type GenS26 struct { F0 bool }
type GenS27 struct { F0 bool }
type GenS28 struct { F0 int8 }
type GenS29 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS30 struct { F0 int16 }
type GenS31 struct { F0 bool }
type GenS32 struct { F0 uint64; F1 uint16; F2 string }
type GenS33 struct { F0 string; F1 uint }
type GenS34 struct { F0 bool; F1 uint32; F2 uint }
type GenS35 struct { F0 int8; F1 uint64 }
type GenS36 struct { F0 int16; F1 uint }
type GenS37 struct { F0 int64 }
type GenS38 struct { F0 uint32 }
type GenS39 struct { F0 uint }
type GenS40 struct { F0 int8 }
type GenS41 struct { F0 uint32 }
type GenS42 struct { F0 int16; F1 int64; F2 uint32 }
type GenS43 struct { F0 uint8; F1 int32; F2 int8 }
type GenS44 struct { F0 int; F1 uint64 }
type GenS45 struct { F0 uint8 }

func main() {
    v1 := make(map[int]string)
    v2 := []int{-9223372036854775808}
    v3 := make(map[uint32]int)
    var v4 int32 = -883
    delete(v3, 65535)
    v5 := GenS1{F0: 1, F1: 2, F2: 32768}
    var v6 int8 = 76
    v7 := v1[915]
    if false {
        var v8 uint64 = 2
        switch v7 {
        case "hello":
            var v9 bool = true
            _ = v9
        case "\n\t":
            switch -9223372036854775808 {
            case 4294967295:
                var v10 int64 = 0
                _ = v10
            case -162:
                v6 = v6
            case 128:
                for fi11 := 0; fi11 < 6; fi11++ {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        if false {
                            if (!(float32(3.14159) >= 3.14159)) {
                                _ = len(v2)
                                delete(v1, (fi11 - fi12))
                                if len(v2) > 0 {
                                    v2[0] = fi11
                                }
                            } else {
                                var v13 float32 = 1.5
                                _ = v13
                            }
                            v7 = v7
                            v14 := GenS2{F0: fi12}
                            _ = v14
                        }
                        switch (fi12 ^ fi11) {
                        case -9223372036854775808:
                            fmt.Println(v7)
                        case -474:
                            var v15 int = fi12
                            _ = v15
                        }
                        if len(v2) > 0 {
                            v2[0] = 2147483647
                        }
                        _ = fi12
                    }
                    if true {
                        var v16 float64 = 1e10
                        if (true || (true || (v6 != v6))) {
                            for fi17 := 0; fi17 < 7; fi17++ {
                                if (true && (65535 > 1)) {
                                    if false {
                                        if len(v2) > 0 {
                                            v2[0] = (fi17 | (fi17 | (fi17 << ((1) & 63))))
                                        }
                                        v18 := GenS2{F0: fi11}
                                        _ = v18
                                    }
                                    if false {
                                        var v19 int = 0
                                        v20 := GenS3{F0: 0, F1: 1, F2: 1}
                                        _ = v19
                                        _ = v20
                                    } else {
                                        fmt.Println(fi17, v7)
                                        if true {
                                            var v21 int8 = (v6 >> ((65535) & 63))
                                            v16 = float64(0.1)
                                            _ = v21
                                        } else {
                                            if false {
                                                v22 := GenS3{F0: 202, F1: 1, F2: 1}
                                                fmt.Println(v7, v8, fi11)
                                                v16 = ((v16 * float64(0.0)) * v16)
                                                _ = v22
                                            }
                                        }
                                    }
                                } else {
                                    v23 := GenS2{F0: (fi17 % (((fi11 * (-(fi11)))) | 1))}
                                    var v24 int16 = 1
                                    _ = v23
                                    _ = v24
                                }
                                if (1e10 >= 3.14159) {
                                    if true {
                                        if false {
                                            if (65535 <= 1) {
                                                delete(v1, fi17)
                                                if false {
                                                    var v25 int32 = (v4 + 2147483647)
                                                    var v26 int = (^(fi11))
                                                    _ = v25
                                                    _ = v26
                                                } else {
                                                    v4 = -807
                                                    if false {
                                                        if false {
                                                            if true {
                                                                var v27 int8 = v6
                                                                _ = v27
                                                            }
                                                        } else {
                                                            var v28 uint32 = 37
                                                            _ = v28
                                                        }
                                                    } else {
                                                        var v29 float32 = ((-2.25 / 1.0) + 0.1)
                                                        _ = v29
                                                    }
                                                    var v30 uint16 = 1
                                                    _ = v30
                                                }
                                                fmt.Println(v4, v6)
                                            } else {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    switch v7 {
                                                    case "a":
                                                        v8 += (v8 - 518)
                                                    default:
                                                        v31 := GenS3{F0: 50, F1: 32767, F2: 0}
                                                        _ = v31
                                                    }
                                                case "":
                                                    var v32 int32 = v4
                                                    _ = v32
                                                case "a":
                                                    v16 = (((float64(1e10) - v16) / ((0.1 + v16) / (v16 - v16))) / -2.25)
                                                }
                                                v33 := GenS4{F0: 42, F1: (v7 + "a"), F2: "hello"}
                                                v4 = v4
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v34 := GenS5{F0: 42, F1: 0, F2: 0}
                                                _ = v33
                                                _ = v34
                                            }
                                            var v35 int16 = 0
                                            var v36 uint16 = 2
                                            if false {
                                                v7 = v7
                                                var v37 uint = 17
                                                v37 -= v37
                                                v38 := GenS6{F0: 128, F1: v8, F2: 92}
                                                var v39 uint = 2
                                                _ = v37
                                                _ = v38
                                                _ = v39
                                            } else {
                                                var v40 uint64 = v8
                                                v2 = append(v2, fi11)
                                                _ = v40
                                            }
                                            _ = v35
                                            _ = v36
                                        }
                                        switch fi17 {
                                        case -137:
                                            var v41 int8 = -128
                                            _ = v41
                                        case -9223372036854775808:
                                            v8 -= v8
                                        case 442:
                                            _ = len(v1)
                                        }
                                    } else {
                                        v2 = append(v2, fi11)
                                        if false {
                                            var v42 uint64 = (v8 - v8)
                                            _ = v42
                                        } else {
                                            switch v7 {
                                            case "":
                                                if len(v2) > 0 {
                                                    v2[0] = (^(fi17))
                                                }
                                            case "hello":
                                                var v43 int = (^(fi11))
                                                _ = v43
                                            case "\n\t":
                                                v16 = v16
                                            default:
                                                var v44 uint16 = 2
                                                _ = v44
                                            }
                                            fmt.Println(fi11, v8, v16)
                                        }
                                    }
                                } else {
                                    var v45 uint = 65535
                                    v46 := v3[941]
                                    _ = v45
                                    _ = v46
                                }
                                var v47 uint8 = 42
                                _ = fi17
                                _ = v47
                            }
                            var v48 int = -9223372036854775808
                            _ = len(v2)
                            _ = v48
                        }
                        _ = v16
                    }
                    _ = fi11
                }
            }
        case "foo bar":
            v49 := GenS1{F0: 0, F1: 2, F2: 42}
            _ = v49
        }
        v7 = v7
        _ = v8
    } else {
        if true {
            var v50 int8 = v6
            v50 += v50
            switch 42 {
            case 127:
                v51 := GenS7{F0: 4294967295}
                _ = v51
            case 128:
                if true {
                    if false {
                        _ = len(v1)
                        var v52 float32 = (0.0 - ((((((0.0 - 1.5) * (-2.25 + 0.0)) + (((0.0 + 1.5) / ((1e10 / 1.0) - (0.0 / float32(1e10)))) + 1.5)) / -2.25) / (1e10 + float32((0.0 / 0.1)))) - (-2.25 / (((0.0 * 1e10) * (-2.25 / 1.0)) - 3.14159))))
                        for fi53 := 0; fi53 < 9; fi53++ {
                            v54 := GenS8{F0: 1, F1: 0, F2: -2147483648}
                            _ = fi53
                            _ = v54
                        }
                        _ = v52
                    } else {
                        v55 := GenS9{F0: 42, F1: 2, F2: 42}
                        _ = v55
                    }
                    v56 := GenS2{F0: -1}
                    _ = v56
                } else {
                    var v57 uint32 = 2
                    _ = v57
                }
            case 32767:
                v58 := GenS10{F0: 0, F1: -9223372036854775808}
                _ = v58
            default:
                if true {
                    if (!(!false)) {
                        if false {
                            for fi59 := 0; fi59 < 17; fi59++ {
                                v7 = ("hello" + (v7 + (v7 + v7)))
                                if true {
                                    v2 = append(v2, (fi59 & fi59))
                                    if false {
                                        switch "foo bar" {
                                        case "hello":
                                            var v60 float32 = 3.14159
                                            _ = v60
                                        default:
                                            fmt.Println(v7, v4)
                                        }
                                    }
                                } else {
                                    _ = len(v1)
                                }
                                fmt.Println(v50)
                                _ = fi59
                            }
                            switch 32768 {
                            case -9223372036854775808:
                                v4 -= (-(v4))
                            case 187:
                                switch v7 {
                                case "foo bar":
                                    fmt.Println(v4)
                                case "hello":
                                    var v61 string = (v7 + (v7 + v7))
                                    _ = v61
                                }
                            default:
                                v62 := GenS1{F0: 0, F1: 4294967295, F2: 4294967295}
                                _ = v62
                            }
                        } else {
                            fmt.Println(v50, v4)
                            v1[-712] = v7
                        }
                        var v63 int32 = (v4 | (v4 * v4))
                        _ = v63
                    }
                }
            }
            _ = v50
        } else {
            for fi64 := 0; fi64 < 12; fi64++ {
                v65 := GenS1{F0: 255, F1: 4294967295, F2: -339}
                if true {
                    if ((v4 != (v4 - (-(v4)))) || true) {
                        for fi66 := 0; fi66 < 1; fi66++ {
                            v67 := v1[(fi64 / ((fi64) | 1))]
                            var v68 int32 = (v4 * (v4 * (v4 * (^(v4)))))
                            var v69 int16 = 787
                            _ = fi66
                            _ = v67
                            _ = v68
                            _ = v69
                        }
                        var v70 string = "foo bar"
                        for fi71 := 0; fi71 < 9; fi71++ {
                            v72 := GenS11{F0: -102, F1: (-128 < (-(v6))), F2: 46}
                            if true {
                                var v73 int8 = (v6 % ((v6) | 1))
                                switch fi64 {
                                case -126:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case -457:
                                    v74 := GenS12{F0: 680, F1: v4}
                                    _ = v74
                                }
                                _ = v73
                            }
                            var v75 int32 = (v4 % ((2147483647) | 1))
                            if (-2.25 > 0.1) {
                                switch (fi64 + (^(fi64))) {
                                case 364:
                                    _ = len(v2)
                                case -1:
                                    if (!(1.5 != -2.25)) {
                                        switch (("\n\t" + (v70 + "\n\t")) + v70) {
                                        case "":
                                            delete(v3, 42)
                                        case "hello":
                                            _ = len(v3)
                                            if true {
                                                fmt.Println(fi64, fi71, v6)
                                                v75 = v75
                                            } else {
                                                v70 = (v70 + (v7 + v70))
                                            }
                                            v76 := GenS13{F0: 32767, F1: 1}
                                            _ = v76
                                        default:
                                            v6 = v6
                                        }
                                        delete(v3, 2147483647)
                                        fmt.Println(fi71, v4)
                                        if true {
                                            v7 = "a"
                                            fmt.Println(v6, fi64, fi64)
                                            delete(v3, 4294967295)
                                        }
                                        var v77 uint16 = 584
                                        _ = v77
                                    } else {
                                        v7 = v70
                                        v4 += int32(v6)
                                    }
                                default:
                                    fmt.Println(fi64, v75)
                                    if (!true) {
                                        if true {
                                            v78 := GenS14{F0: 255, F1: 701}
                                            v79 := GenS15{F0: 127, F1: (v4 >> ((65535) & 63)), F2: (v6 % ((127) | 1))}
                                            _ = v78
                                            _ = v79
                                        }
                                    }
                                }
                                if (false && (true && (false || true))) {
                                    if (32767 >= 0) {
                                        var v80 uint8 = 161
                                        _ = v80
                                    } else {
                                        if false {
                                            fmt.Println(v75, v70, v75)
                                        }
                                        fmt.Println(v4, fi64, v7)
                                        var v81 bool = (v7 == v70)
                                        _ = v81
                                    }
                                    fmt.Println(fi71, v7)
                                    v82 := v1[fi71]
                                    _ = len(v3)
                                    v83 := GenS16{F0: 157, F1: 42}
                                    v84 := GenS1{F0: 1, F1: 919, F2: 0}
                                    if ((2147483647 != 2147483648) && true) {
                                        v7 = v7
                                        v85 := GenS4{F0: 2, F1: v82, F2: "a"}
                                        v86 := GenS14{F0: 128, F1: 96}
                                        _ = v85
                                        _ = v86
                                    } else {
                                        v82 = (v82 + "foo bar")
                                        v87 := GenS12{F0: 277, F1: v75}
                                        var v88 int16 = 32767
                                        _ = v87
                                        _ = v88
                                    }
                                    _ = v82
                                    _ = v83
                                    _ = v84
                                }
                                v70 = (v70 + ((v7 + v7) + v70))
                            } else {
                                v4 -= (v4 * -903)
                                v6 += v6
                                var v89 int64 = 9223372036854775807
                                var v90 string = v7
                                if false {
                                    fmt.Println(v70, fi71, v75)
                                    v91 := GenS17{F0: (v6 << ((923) & 63)), F1: "\n\t", F2: v75}
                                    v92 := GenS8{F0: 1, F1: 42, F2: -530}
                                    _ = v91
                                    _ = v92
                                } else {
                                    v93 := GenS4{F0: 0, F1: "a", F2: v90}
                                    v94 := GenS18{F0: fi71, F1: -84, F2: 281}
                                    _ = v93
                                    _ = v94
                                }
                                _ = v89
                                _ = v90
                            }
                            _ = fi71
                            _ = v72
                            _ = v75
                        }
                        _ = v70
                    }
                } else {
                    switch 42 {
                    case 42:
                        for fi95 := 0; fi95 < 20; fi95++ {
                            v3[0] = (fi95 << ((380) & 63))
                            _ = fi95
                        }
                    case 1:
                        var v96 float32 = 0.1
                        _ = v96
                    case 127:
                        if (42 >= 42) {
                            v97 := GenS19{F0: v6, F1: 208}
                            _ = v97
                        }
                    }
                    var v98 uint8 = 255
                    fmt.Println(fi64, fi64, v7)
                    _ = v98
                }
                var v99 string = ("a" + v7)
                _ = fi64
                _ = v65
                _ = v99
            }
            v100 := GenS20{F0: 1}
            for fi101 := 0; fi101 < 14; fi101++ {
                var v102 int16 = 32767
                for fi103 := 0; fi103 < 11; fi103++ {
                    var v104 int16 = (^(v102))
                    _ = fi103
                    _ = v104
                }
                var v105 uint32 = 2
                _ = fi101
                _ = v102
                _ = v105
            }
            _ = v100
        }
        if (!false) {
            for fi106 := 0; fi106 < 11; fi106++ {
                v7 = (v7 + (((v7 + (v7 + ((((v7 + ("hello" + v7)) + (((v7 + v7) + v7) + (v7 + v7))) + v7) + "a"))) + v7) + "foo bar"))
                for fi107 := 0; fi107 < 17; fi107++ {
                    fmt.Println(fi106, v7, v6)
                    _ = fi107
                }
                fmt.Println(v6, v6)
                _ = fi106
            }
            for fi108 := 0; fi108 < 2; fi108++ {
                if true {
                    var v109 uint32 = 1
                    _ = v109
                }
                var v110 bool = (!true)
                _ = fi108
                _ = v110
            }
            if ((false || true) && false) {
                v4 = (^(v4))
            }
        } else {
            for fi111 := 0; fi111 < 8; fi111++ {
                v112 := GenS21{F0: ("foo bar" + v7), F1: v6}
                v6 = (v6 >> ((328) & 63))
                _ = fi111
                _ = v112
            }
            for fi113 := 0; fi113 < 3; fi113++ {
                for fi114 := 0; fi114 < 1; fi114++ {
                    if (v6 < v6) {
                        var v115 uint16 = 32767
                        fmt.Println(fi114, v4, fi114)
                        var v116 float64 = 1.5
                        _ = v115
                        _ = v116
                    }
                    v117 := GenS22{F0: 9223372036854775807}
                    _ = fi114
                    _ = v117
                }
                _ = fi113
            }
        }
        v4 = v4
    }
    v118 := GenS23{F0: 73, F1: 4294967295}
    for fi119 := 0; fi119 < 2; fi119++ {
        switch "foo bar" {
        case "":
            v120 := GenS24{F0: 184, F1: ("\n\t" + "")}
            _ = v120
        case "\n\t":
            for fi121 := 0; fi121 < 5; fi121++ {
                switch 127 {
                case -1:
                    var v122 int16 = 0
                    _ = v122
                default:
                    delete(v1, fi121)
                }
                v6 = (^(v6))
                v123 := v3[2147483648]
                _ = fi121
                _ = v123
            }
        }
        for fi124 := 0; fi124 < 2; fi124++ {
            v7 = ((v7 + (v7 + v7)) + (v7 + (v7 + "a")))
            v6 = v6
            _ = fi124
        }
        _ = fi119
    }
    var v125 uint = 186
    v125 -= 4294967295
    v126 := GenS25{F0: (v4 << (((v125 * (^(v125)))) & 63)), F1: 553, F2: -9223372036854775808}
    v127 := GenS15{F0: 127, F1: -578, F2: (v6 | (v6 << (((v125 & v125)) & 63)))}
    if false {
        switch "\n\t" {
        case "":
            fmt.Println(v7, v4, v7)
        default:
            var v128 uint32 = 1
            _ = v128
        }
        v6 = v6
    } else {
        fmt.Println(v7, v6, v4)
        if (!true) {
            v7 = v7
            for fi129 := 0; fi129 < 10; fi129++ {
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v130 int32 = v4
                    _ = v130
                } else {
                    if true {
                        for fi131 := 0; fi131 < 15; fi131++ {
                            v6 = -1
                            v2 = append(v2, (fi129 * (fi131 * (fi129 << ((v125) & 63)))))
                            delete(v3, 2147483648)
                            _ = fi131
                        }
                        var v132 int16 = int16(v6)
                        _ = v132
                    } else {
                        fmt.Println(fi129, fi129, v7)
                        switch fi129 {
                        case 15:
                            for fi133 := 0; fi133 < 11; fi133++ {
                                if true {
                                    if false {
                                        var v134 string = v7
                                        v135 := GenS4{F0: 65535, F1: "\n\t", F2: (v134 + v134)}
                                        v136 := v3[0]
                                        _ = v134
                                        _ = v135
                                        _ = v136
                                    }
                                    v7 = v7
                                    switch v7 {
                                    case "hello":
                                        v6 = 127
                                        var v137 float32 = 0.0
                                        switch fi129 {
                                        case -1:
                                            var v138 float32 = 0.1
                                            _ = v138
                                        case -9223372036854775808:
                                            var v139 uint8 = 42
                                            _ = v139
                                        case -348:
                                            v140 := GenS26{F0: false}
                                            _ = v140
                                        }
                                        _ = v137
                                    case "a":
                                        var v141 uint32 = 2
                                        _ = v141
                                    }
                                    if true {
                                        switch (fi129 ^ fi129) {
                                        case 4294967295:
                                            delete(v1, 126)
                                        case 128:
                                            fmt.Println(fi133)
                                        case -1:
                                            var v142 uint64 = 624
                                            var v143 uint16 = 2
                                            _ = v142
                                            _ = v143
                                        default:
                                            if false {
                                                var v144 int16 = 743
                                                if (fi133 != int(v144)) {
                                                    v145 := GenS27{F0: false}
                                                    fmt.Println(v6)
                                                    var v146 int8 = v6
                                                    _ = v145
                                                    _ = v146
                                                } else {
                                                    if (3.14159 <= 0.1) {
                                                        if true {
                                                            v147 := GenS28{F0: -1}
                                                            var v148 string = "a"
                                                            _ = v147
                                                            _ = v148
                                                        }
                                                        v7 = (v7 + (v7 + "foo bar"))
                                                        fmt.Println(fi129)
                                                    }
                                                    v2 = append(v2, 1)
                                                    fmt.Println(v125, fi129)
                                                }
                                                _ = v144
                                            }
                                        }
                                        v149 := GenS13{F0: -278, F1: 42}
                                        if (2 >= 2147483647) {
                                            switch v7 {
                                            case "":
                                                if true {
                                                    if (false || false) {
                                                        var v150 int8 = (v6 >> (((v125 + (v125 & v125))) & 63))
                                                        fmt.Println(v4)
                                                        _ = v150
                                                    } else {
                                                        switch fi133 {
                                                        case 32768:
                                                            v151 := v1[-557]
                                                            _ = v151
                                                        case -833:
                                                            delete(v1, fi133)
                                                        case -480:
                                                            if len(v2) > 0 {
                                                                v2[0] = (fi133 - -2147483648)
                                                            }
                                                        default:
                                                            if (-1 > v6) {
                                                                switch fi133 {
                                                                case 127:
                                                                    switch (v7 + (("foo bar" + "") + "hello")) {
                                                                    case "foo bar":
                                                                        v7 = v7
                                                                    default:
                                                                        if len(v2) > 0 {
                                                                            v2[0] = (^(fi133))
                                                                        }
                                                                    }
                                                                case -693:
                                                                    _ = len(v1)
                                                                default:
                                                                    fmt.Println(fi129, v125)
                                                                }
                                                                fmt.Println(v125, v6)
                                                                switch (v7 + v7) {
                                                                case "":
                                                                    v152 := GenS4{F0: 65535, F1: (v7 + "hello"), F2: v7}
                                                                    _ = v152
                                                                default:
                                                                    var v153 uint64 = 2
                                                                    _ = v153
                                                                }
                                                            } else {
                                                                var v154 int = fi133
                                                                _ = v154
                                                            }
                                                            v7 = v7
                                                            fmt.Println(v125)
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        if ((float64(-2.25) > ((1e10 + -2.25) / 1e10)) || ((^(v125)) == (v125 / ((v125) | 1)))) {
                                                            fmt.Println(v6, v125)
                                                            var v155 uint8 = 24
                                                            _ = v155
                                                        } else {
                                                            var v156 float32 = 1.5
                                                            v157 := GenS3{F0: 61, F1: -822, F2: 1}
                                                            _ = v156
                                                            _ = v157
                                                        }
                                                        if (238 > 42) {
                                                            switch (v7 + "foo bar") {
                                                            case "":
                                                                if ((((float64(0.0) / 1.5) * 1e10) == float64(0.0)) || ((v6 * (-(v6))) > (v6 ^ v6))) {
                                                                    _ = len(v2)
                                                                }
                                                            case "\n\t":
                                                                if true {
                                                                    if ((1.5 + float32((1.5 / 1e10))) <= 1.5) {
                                                                        v158 := GenS22{F0: int64(v6)}
                                                                        v1[fi133] = ((v7 + v7) + v7)
                                                                        var v159 uint64 = 18446744073709551615
                                                                        _ = v158
                                                                        _ = v159
                                                                    } else {
                                                                        fmt.Println(v125)
                                                                    }
                                                                    switch fi129 {
                                                                    case 127:
                                                                        fmt.Println(fi129, v7, fi133)
                                                                        v7 = "hello"
                                                                    case 2147483647:
                                                                        if (0.1 < (((((float32(0.0) * 0.1) - 1.5) + (1e10 + (3.14159 + 1.5))) * 3.14159) + 1.5)) {
                                                                            var v160 uint16 = 0
                                                                            v6 -= v6
                                                                            fmt.Println(fi129, fi133, v4)
                                                                            _ = v160
                                                                        } else {
                                                                            v2 = append(v2, -9223372036854775808)
                                                                            if (1000 == -1) {
                                                                                fmt.Println(fi129, v7)
                                                                            }
                                                                        }
                                                                        var v161 int64 = -9223372036854775808
                                                                        _ = v161
                                                                    case -319:
                                                                        v6 = (v6 * (v6 % ((v6) | 1)))
                                                                    }
                                                                }
                                                            default:
                                                                var v162 uint = (v125 / ((v125) | 1))
                                                                _ = v162
                                                            }
                                                        }
                                                    }
                                                }
                                            case "\n\t":
                                                v4 += v4
                                            case "foo bar":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            default:
                                                if (4294967295 < 42) {
                                                    _ = len(v2)
                                                } else {
                                                    if false {
                                                        var v163 uint64 = 42
                                                        if (42 <= 446) {
                                                            fmt.Println(v125)
                                                        } else {
                                                            if (!(0 >= 32)) {
                                                                if (v163 != v163) {
                                                                    v164 := GenS29{F0: -680, F1: 2, F2: (^(v163))}
                                                                    _ = v164
                                                                } else {
                                                                    if (!true) {
                                                                        switch v7 {
                                                                        case "":
                                                                            v2 = append(v2, fi133)
                                                                        case "\n\t":
                                                                            v2 = append(v2, fi133)
                                                                        }
                                                                        if ((!(!false)) || false) {
                                                                            v6 -= (-(v6))
                                                                        } else {
                                                                            _ = len(v2)
                                                                            v125 = (v125 << (((v125 * (v125 & 799))) & 63))
                                                                            v165 := GenS30{F0: int16(v4)}
                                                                            _ = v165
                                                                        }
                                                                        v6 = 42
                                                                    } else {
                                                                        switch (fi129 ^ (fi129 + (-(fi129)))) {
                                                                        case 353:
                                                                            var v166 int64 = -9223372036854775808
                                                                            _ = v166
                                                                        case -1:
                                                                            if (4294967295 < (v163 * (v163 + 1))) {
                                                                                v167 := v1[(fi133 ^ fi133)]
                                                                                var v168 bool = true
                                                                                _ = v167
                                                                                _ = v168
                                                                            }
                                                                        case 0:
                                                                            var v169 uint16 = 481
                                                                            _ = v169
                                                                        }
                                                                        var v170 bool = (109 > 139)
                                                                        if false {
                                                                            switch fi129 {
                                                                            case 12:
                                                                                v4 -= v4
                                                                            default:
                                                                                var v171 uint32 = 0
                                                                                _ = v171
                                                                            }
                                                                            switch "a" {
                                                                            case "a":
                                                                                var v172 int16 = -1
                                                                                _ = v172
                                                                            case "":
                                                                                v4 = v4
                                                                            case "foo bar":
                                                                                v173 := v1[128]
                                                                                _ = v173
                                                                            default:
                                                                                var v174 float32 = 1.5
                                                                                _ = v174
                                                                            }
                                                                        }
                                                                        _ = v170
                                                                    }
                                                                    var v175 bool = false
                                                                    var v176 float64 = 3.14159
                                                                    _ = v175
                                                                    _ = v176
                                                                }
                                                                var v177 uint = v125
                                                                var v178 int = -1
                                                                _ = v177
                                                                _ = v178
                                                            } else {
                                                                _ = len(v2)
                                                            }
                                                            fmt.Println(v6, v125)
                                                            if len(v2) > 0 {
                                                                v2[0] = fi129
                                                            }
                                                        }
                                                        switch ((v7 + v7) + v7) {
                                                        case "hello":
                                                            v6 = (-(v6))
                                                        case "a":
                                                            var v179 int16 = -1
                                                            _ = v179
                                                        case "":
                                                            v180 := GenS31{F0: false}
                                                            _ = v180
                                                        default:
                                                            switch "" {
                                                            case "foo bar":
                                                                if (fi129 <= (fi133 * 127)) {
                                                                    _ = len(v2)
                                                                    switch "" {
                                                                    case "a":
                                                                        var v181 uint16 = 255
                                                                        _ = v181
                                                                    case "hello":
                                                                        v182 := GenS32{F0: (v163 - 863), F1: 211, F2: v7}
                                                                        _ = v182
                                                                    case "\n\t":
                                                                        delete(v1, fi129)
                                                                    default:
                                                                        if false {
                                                                            var v183 int = (fi133 + (fi129 >> (((v125 | (v125 * v125))) & 63)))
                                                                            var v184 uint32 = 1
                                                                            _ = v183
                                                                            _ = v184
                                                                        } else {
                                                                            v185 := GenS26{F0: false}
                                                                            if false {
                                                                                v6 = (v6 - v6)
                                                                                if true {
                                                                                    v2 = append(v2, fi133)
                                                                                    v2 = append(v2, fi133)
                                                                                    v4 = (-(v4))
                                                                                }
                                                                            }
                                                                            if (((v7 + v7) + v7) > ("hello" + (v7 + ("" + "\n\t")))) {
                                                                                if (65535 > 42) {
                                                                                    v125 = 184
                                                                                    v125 -= v125
                                                                                } else {
                                                                                    if (-778 <= -11) {
                                                                                        if (true || ((v6 / ((-1) | 1)) == (v6 + -117))) {
                                                                                            v186 := GenS33{F0: ("\n\t" + "foo bar"), F1: v125}
                                                                                            _ = v186
                                                                                        } else {
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                        }
                                                                                        switch (^(fi133)) {
                                                                                        case 127:
                                                                                            switch fi129 {
                                                                                            case 0:
                                                                                                switch v7 {
                                                                                                case "a":
                                                                                                    if len(v2) >= 2 {
                                                                                                        v2 = v2[1:len(v2)]
                                                                                                    }
                                                                                                case "":
                                                                                                    fmt.Println(fi133)
                                                                                                case "hello":
                                                                                                    v7 = "\n\t"
                                                                                                default:
                                                                                                    if (((1e10 * 1e10) - (0.0 / 1e10)) == 1.5) {
                                                                                                        var v187 int16 = -32768
                                                                                                        var v188 float64 = (0.1 * -2.25)
                                                                                                        var v189 bool = false
                                                                                                        _ = v187
                                                                                                        _ = v188
                                                                                                        _ = v189
                                                                                                    } else {
                                                                                                        var v190 string = (("a" + v7) + (((v7 + ((v7 + v7) + v7)) + "\n\t") + (v7 + v7)))
                                                                                                        var v191 uint16 = 0
                                                                                                        fmt.Println(fi129)
                                                                                                        _ = v190
                                                                                                        _ = v191
                                                                                                    }
                                                                                                }
                                                                                            case 9223372036854775807:
                                                                                                v192 := GenS34{F0: false, F1: 227, F2: v125}
                                                                                                _ = v192
                                                                                            case 127:
                                                                                                v7 = ("" + v7)
                                                                                            default:
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = fi129
                                                                                                }
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    v6 += 1
                                                                                                case "a":
                                                                                                    var v193 uint32 = 257
                                                                                                    _ = v193
                                                                                                case "\n\t":
                                                                                                    v163 = (v163 & v163)
                                                                                                default:
                                                                                                    v7 = "foo bar"
                                                                                                }
                                                                                            }
                                                                                        case -772:
                                                                                            v2 = append(v2, fi133)
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                        }
                                                                                        var v194 int8 = (v6 / ((v6) | 1))
                                                                                        _ = v194
                                                                                    }
                                                                                }
                                                                                switch v7 {
                                                                                case "a":
                                                                                    if true {
                                                                                        v2 = append(v2, fi129)
                                                                                        if false {
                                                                                            var v195 uint64 = v163
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = fi133
                                                                                            }
                                                                                            v196 := GenS18{F0: 32767, F1: 42, F2: 32767}
                                                                                            _ = v195
                                                                                            _ = v196
                                                                                        }
                                                                                    } else {
                                                                                        if false {
                                                                                            var v197 uint32 = 84
                                                                                            v198 := GenS26{F0: false}
                                                                                            v199 := GenS34{F0: (42 != 0), F1: (^(v197)), F2: (v125 + v125)}
                                                                                            if (v6 > (-(v6))) {
                                                                                                v200 := GenS4{F0: 255, F1: "\n\t", F2: ("" + ((("" + v7) + "\n\t") + "hello"))}
                                                                                                _ = v200
                                                                                            }
                                                                                            _ = v197
                                                                                            _ = v198
                                                                                            _ = v199
                                                                                        }
                                                                                    }
                                                                                case "":
                                                                                    fmt.Println(v4)
                                                                                default:
                                                                                    v201 := GenS35{F0: 1, F1: 0}
                                                                                    _ = v201
                                                                                }
                                                                                v202 := GenS6{F0: 42, F1: 0, F2: 255}
                                                                                _ = v202
                                                                            } else {
                                                                                v203 := GenS31{F0: false}
                                                                                if false {
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = fi133
                                                                                    }
                                                                                    if (int(v125) > (fi133 / (((fi129 >> (((^(v125))) & 63))) | 1))) {
                                                                                        _ = len(v3)
                                                                                        fmt.Println(v6, v7, fi133)
                                                                                        delete(v1, (-(fi129)))
                                                                                    } else {
                                                                                        switch ("hello" + ("hello" + "hello")) {
                                                                                        case "foo bar":
                                                                                            v204 := GenS29{F0: 205, F1: v163, F2: v163}
                                                                                            _ = v204
                                                                                        case "a":
                                                                                            var v205 uint = v125
                                                                                            if (-573 < -1) {
                                                                                                fmt.Println(v7, v6, v7)
                                                                                            } else {
                                                                                                var v206 uint64 = (^(v163))
                                                                                                v207 := v1[0]
                                                                                                _ = v206
                                                                                                _ = v207
                                                                                            }
                                                                                            _ = v205
                                                                                        case "":
                                                                                            var v208 bool = (!(true && false))
                                                                                            _ = v208
                                                                                        }
                                                                                        v209 := GenS26{F0: false}
                                                                                        _ = v209
                                                                                    }
                                                                                    switch (fi133 ^ fi133) {
                                                                                    case 9223372036854775807:
                                                                                        v163 += v163
                                                                                    default:
                                                                                        v4 += v4
                                                                                    }
                                                                                } else {
                                                                                    var v210 uint64 = v163
                                                                                    _ = v210
                                                                                }
                                                                                v6 = v6
                                                                                _ = v203
                                                                            }
                                                                            _ = v185
                                                                        }
                                                                        var v211 uint = v125
                                                                        _ = v211
                                                                    }
                                                                } else {
                                                                    fmt.Println(v163, v6, v4)
                                                                }
                                                            case "":
                                                                var v212 uint16 = 32767
                                                                _ = v212
                                                            case "a":
                                                                if (!(!true)) {
                                                                    fmt.Println(v6)
                                                                    v213 := GenS29{F0: 1, F1: v163, F2: (v163 * (v163 - 410))}
                                                                    if false {
                                                                        v4 = 488
                                                                        if false {
                                                                            if ((v163 << ((2) & 63)) < 42) {
                                                                                v214 := GenS36{F0: 42, F1: 1}
                                                                                v7 = ((v7 + v7) + v7)
                                                                                if true {
                                                                                    v4 = v4
                                                                                    v215 := GenS37{F0: -9223372036854775808}
                                                                                    delete(v1, 1)
                                                                                    _ = v215
                                                                                }
                                                                                _ = v214
                                                                            }
                                                                            v216 := GenS29{F0: -199, F1: (v163 / ((v163) | 1)), F2: v163}
                                                                            v7 = v7
                                                                            _ = v216
                                                                        }
                                                                        v7 = v7
                                                                    }
                                                                    _ = v213
                                                                }
                                                            default:
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(v125, v6, fi129)
                                                        _ = v163
                                                    }
                                                }
                                            }
                                            var v217 int32 = 0
                                            _ = v217
                                        } else {
                                            switch -1 {
                                            case 26:
                                                var v218 string = v7
                                                _ = v218
                                            case 149:
                                                switch fi129 {
                                                case -2147483648:
                                                    _ = len(v3)
                                                case 9223372036854775807:
                                                    var v219 int16 = 0
                                                    _ = v219
                                                case 2147483647:
                                                    v220 := GenS22{F0: -9223372036854775808}
                                                    _ = v220
                                                default:
                                                    v7 = ""
                                                }
                                            case -2147483648:
                                                if (fi129 > (fi133 * (fi133 ^ (fi133 >> ((v125) & 63))))) {
                                                    var v221 float32 = (1e10 * -2.25)
                                                    _ = v221
                                                }
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = (fi129 << ((130) & 63))
                                            }
                                        }
                                        _ = v149
                                    } else {
                                        v125 -= v125
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = (fi133 ^ (^(fi133)))
                                    }
                                } else {
                                    v6 += 100
                                    var v222 uint8 = 128
                                    if false {
                                        if false {
                                            var v223 int16 = 52
                                            if false {
                                                v224 := GenS38{F0: 4294967295}
                                                v225 := GenS27{F0: (!true)}
                                                _ = v224
                                                _ = v225
                                            }
                                            if true {
                                                switch "a" {
                                                case "":
                                                    switch fi129 {
                                                    case 32768:
                                                        var v226 int64 = 25
                                                        _ = v226
                                                    case 42:
                                                        fmt.Println(v7, v223, v4)
                                                    }
                                                default:
                                                    v227 := GenS31{F0: false}
                                                    _ = v227
                                                }
                                                var v228 uint16 = 42
                                                _ = v228
                                            } else {
                                                v229 := GenS35{F0: -128, F1: 42}
                                                v125 = (v125 << ((v125) & 63))
                                                _ = v229
                                            }
                                            switch "" {
                                            case "foo bar":
                                                switch (((v7 + "hello") + ((((v7 + v7) + v7) + (v7 + ("\n\t" + (("hello" + v7) + v7)))) + (("hello" + "a") + (v7 + "a")))) + ((("hello" + "\n\t") + v7) + ((v7 + v7) + "a"))) {
                                                case "":
                                                    var v230 uint32 = 0
                                                    _ = v230
                                                case "a":
                                                    switch v7 {
                                                    case "\n\t":
                                                        v222 = v222
                                                    case "":
                                                        v6 = 111
                                                    default:
                                                        if (v6 < int8(v223)) {
                                                            if len(v2) > 0 {
                                                                v2[0] = -800
                                                            }
                                                            var v231 uint = (v125 % (((v125 << (((v125 % (((^(v125))) | 1))) & 63))) | 1))
                                                            _ = v231
                                                        } else {
                                                            v125 = 18446744073709551615
                                                        }
                                                    }
                                                case "hello":
                                                    v6 = 100
                                                default:
                                                    if (-1 > 354) {
                                                        if false {
                                                            if (!true) {
                                                                if false {
                                                                    if true {
                                                                        v7 = v7
                                                                    } else {
                                                                        v4 = (^(v4))
                                                                        var v232 uint = (v125 - 0)
                                                                        if len(v2) > 0 {
                                                                            v2[0] = -9223372036854775808
                                                                        }
                                                                        v232 = v232
                                                                        var v233 int64 = -1
                                                                        _ = v232
                                                                        _ = v233
                                                                    }
                                                                    var v234 int = (fi129 & 415)
                                                                    if (!true) {
                                                                        switch v234 {
                                                                        case -9223372036854775808:
                                                                            v222 = (v222 >> ((610) & 63))
                                                                        case 927:
                                                                            var v235 uint = 18446744073709551615
                                                                            _ = v235
                                                                        }
                                                                        if (-608 == (-(v223))) {
                                                                            if (626 >= 579) {
                                                                                v223 = (v223 - 1000)
                                                                                if (v7 < (v7 + "foo bar")) {
                                                                                    var v236 uint8 = 250
                                                                                    v237 := GenS39{F0: (v125 & 4294967295)}
                                                                                    fmt.Println(v125, v4, v6)
                                                                                    _ = v236
                                                                                    _ = v237
                                                                                } else {
                                                                                    var v238 bool = true
                                                                                    delete(v3, uint32(v125))
                                                                                    _ = v238
                                                                                }
                                                                                v223 = 1000
                                                                            } else {
                                                                                if true {
                                                                                    var v239 int = fi129
                                                                                    _ = v239
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v234
                                                                } else {
                                                                    v240 := GenS4{F0: 255, F1: "\n\t", F2: v7}
                                                                    v2 = append(v2, fi129)
                                                                    _ = v240
                                                                }
                                                            } else {
                                                                var v241 float32 = 1.5
                                                                v242 := GenS40{F0: (v6 ^ int8(v4))}
                                                                _ = v241
                                                                _ = v242
                                                            }
                                                        }
                                                        if true {
                                                            if true {
                                                                switch v7 {
                                                                case "\n\t":
                                                                    switch (fi133 + fi133) {
                                                                    case -249:
                                                                        var v243 string = v7
                                                                        _ = v243
                                                                    case -2147483648:
                                                                        var v244 bool = (103 >= 65535)
                                                                        _ = v244
                                                                    case -990:
                                                                        v245 := GenS25{F0: (v4 & 736), F1: 0, F2: 42}
                                                                        _ = v245
                                                                    }
                                                                case "":
                                                                    v246 := GenS31{F0: true}
                                                                    _ = v246
                                                                case "hello":
                                                                    var v247 float32 = ((1.5 / ((float32(3.14159) + (float32(1.5) + 3.14159)) / ((0.0 - 0.1) + float32(float64(1.5))))) / float32(1e10))
                                                                    _ = v247
                                                                default:
                                                                    var v248 int64 = 9223372036854775807
                                                                    fmt.Println(v248)
                                                                    v4 = v4
                                                                    var v249 int = (fi133 >> ((v125) & 63))
                                                                    v222 += (v222 * (v222 + v222))
                                                                    _ = v248
                                                                    _ = v249
                                                                }
                                                                if true {
                                                                    var v250 string = "hello"
                                                                    v125 = v125
                                                                    var v251 uint = 1
                                                                    _ = v250
                                                                    _ = v251
                                                                }
                                                                v7 = v7
                                                            } else {
                                                                v252 := GenS37{F0: 534}
                                                                _ = v252
                                                            }
                                                            v125 = 392
                                                        } else {
                                                            var v253 int64 = -1
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            _ = len(v1)
                                                            _ = v253
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = (fi129 >> ((42) & 63))
                                                        }
                                                    } else {
                                                        if (285 == 595) {
                                                            var v254 int32 = (v4 * 2147483647)
                                                            _ = v254
                                                        }
                                                        switch 128 {
                                                        case 1:
                                                            if true {
                                                                v6 = 3
                                                                switch fi129 {
                                                                case 32768:
                                                                    v255 := v1[260]
                                                                    _ = v255
                                                                }
                                                            } else {
                                                                var v256 int64 = 136
                                                                _ = v256
                                                            }
                                                        case 2147483647:
                                                            _ = len(v2)
                                                        }
                                                        var v257 uint = 191
                                                        _ = v257
                                                    }
                                                }
                                            case "a":
                                                v125 -= 0
                                            }
                                            if (-1 != (v4 >> ((v125) & 63))) {
                                                v258 := GenS6{F0: (^(v222)), F1: 944, F2: 42}
                                                v7 = "hello"
                                                switch (v7 + v7) {
                                                case "":
                                                    var v259 int16 = v223
                                                    switch "" {
                                                    case "hello":
                                                        var v260 string = (v7 + ("foo bar" + v7))
                                                        _ = v260
                                                    }
                                                    _ = v259
                                                default:
                                                    fmt.Println(v223, v223)
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if ((0.0 + 3.14159) > 3.14159) {
                                                            fmt.Println(v6, v222)
                                                            switch (v7 + ("\n\t" + "hello")) {
                                                            case "":
                                                                v7 = v7
                                                            default:
                                                                v125 += v125
                                                            }
                                                            var v261 uint64 = 18446744073709551615
                                                            _ = v261
                                                        } else {
                                                            v262 := GenS41{F0: 0}
                                                            if false {
                                                                if (0 <= 405) {
                                                                    var v263 uint64 = 539
                                                                    fmt.Println(v125, v6, v223)
                                                                    if (!false) {
                                                                        if false {
                                                                            v264 := GenS42{F0: -1, F1: -51, F2: 259}
                                                                            switch v7 {
                                                                            case "foo bar":
                                                                                v265 := GenS30{F0: v223}
                                                                                _ = v265
                                                                            }
                                                                            fmt.Println(v7, v7, v263)
                                                                            _ = v264
                                                                        } else {
                                                                            v266 := v1[(fi133 + fi129)]
                                                                            _ = v266
                                                                        }
                                                                        v7 = v7
                                                                        fmt.Println(v4)
                                                                    } else {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    }
                                                                    _ = v263
                                                                } else {
                                                                    switch (^(fi133)) {
                                                                    case 127:
                                                                        if true {
                                                                            v4 -= v4
                                                                            v223 = (v223 - (v223 << (((v125 << ((v125) & 63))) & 63)))
                                                                        }
                                                                    case -9223372036854775808:
                                                                        v3[2] = (fi129 / ((fi129) | 1))
                                                                    case 4294967295:
                                                                        v267 := GenS43{F0: (v222 | v222), F1: (v4 | v4), F2: v6}
                                                                        _ = v267
                                                                    }
                                                                    v268 := GenS44{F0: (fi129 - fi133), F1: 18446744073709551615}
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 0
                                                                    }
                                                                    _ = v268
                                                                }
                                                                v269 := GenS24{F0: (v222 >> ((v125) & 63)), F1: "hello"}
                                                                _ = v269
                                                            }
                                                            v270 := v3[42]
                                                            if len(v2) > 0 {
                                                                v2[0] = (^(fi129))
                                                            }
                                                            fmt.Println(v223)
                                                            _ = v262
                                                            _ = v270
                                                        }
                                                    case "\n\t":
                                                        if true {
                                                            if len(v2) > 0 {
                                                                v2[0] = (fi133 ^ fi129)
                                                            }
                                                            v271 := GenS39{F0: v125}
                                                            switch v7 {
                                                            case "hello":
                                                                if (!(!("foo bar" == ""))) {
                                                                    v272 := v3[537]
                                                                    var v273 uint = (^(v125))
                                                                    _ = v272
                                                                    _ = v273
                                                                } else {
                                                                    v274 := v1[309]
                                                                    fmt.Println(v6, fi129, v223)
                                                                    var v275 string = v274
                                                                    v276 := GenS20{F0: (-(v4))}
                                                                    _ = v274
                                                                    _ = v275
                                                                    _ = v276
                                                                }
                                                            default:
                                                                v1[int(v222)] = v7
                                                                v277 := GenS14{F0: v222, F1: 32767}
                                                                _ = v277
                                                            }
                                                            _ = v271
                                                        }
                                                    default:
                                                        if (v7 < v7) {
                                                            if (61 <= v6) {
                                                                var v278 int32 = (^(v4))
                                                                v279 := GenS25{F0: -1, F1: 65535, F2: -9223372036854775808}
                                                                v280 := GenS12{F0: v125, F1: v278}
                                                                _ = v278
                                                                _ = v279
                                                                _ = v280
                                                            } else {
                                                                var v281 bool = (v6 <= v6)
                                                                fmt.Println(v7, v281)
                                                                _ = v281
                                                            }
                                                            var v282 string = (v7 + v7)
                                                            var v283 bool = (-802 > 0)
                                                            _ = v282
                                                            _ = v283
                                                        } else {
                                                            fmt.Println(v7)
                                                            v284 := v1[4294967295]
                                                            _ = v284
                                                        }
                                                    }
                                                }
                                                _ = v258
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            _ = len(v1)
                                            _ = v223
                                        }
                                        switch (^(fi133)) {
                                        case 1:
                                            fmt.Println(v6)
                                        case 0:
                                            var v285 uint16 = 0
                                            _ = v285
                                        case 4294967295:
                                            var v286 uint64 = 0
                                            _ = v286
                                        default:
                                            v7 = (v7 + "\n\t")
                                        }
                                        v287 := GenS18{F0: 127, F1: 34, F2: 0}
                                        v288 := GenS33{F0: "a", F1: 713}
                                        v289 := GenS8{F0: 0, F1: -708, F2: v4}
                                        _ = v287
                                        _ = v288
                                        _ = v289
                                    } else {
                                        var v290 int32 = (-(v4))
                                        v291 := v1[(fi129 ^ (fi129 & -377))]
                                        _ = v290
                                        _ = v291
                                    }
                                    _ = v222
                                }
                                v125 -= (^(v125))
                                v292 := GenS45{F0: 1}
                                v6 += 42
                                var v293 float32 = ((3.14159 + float32((0.0 + 0.1))) / -2.25)
                                _ = fi133
                                _ = v292
                                _ = v293
                            }
                        }
                        v7 = "foo bar"
                    }
                }
                _ = fi129
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v125:", v125)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v118
    _ = v125
    _ = v126
    _ = v127
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
