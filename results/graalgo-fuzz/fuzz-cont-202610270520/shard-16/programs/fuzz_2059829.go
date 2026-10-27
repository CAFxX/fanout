package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 uint; F1 int8; F2 string }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 int8; F1 int; F2 uint32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int32; F1 uint32 }
type GenS11 struct { F0 int32; F1 int; F2 int64 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint8; F1 uint8 }
type GenS14 struct { F0 int64; F1 uint32 }
type GenS15 struct { F0 uint32; F1 int64; F2 int8 }
func genFn0(p0 uint32, p1 uint32) bool {
    v6 := GenS1{F0: 1}
    v7 := []int{-9223372036854775808, 9223372036854775807, 0}
    var v8 float32 = 1e10
    var v9 uint32 = (^(p0))
    var v10 bool = (65535 == 18446744073709551615)
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (1 == 32768)
}
func genFn1(p0 uint, p1 bool, p2 uint) uint64 {
    v11 := []int{0, 127, 4294967295, 42}
    var v12 uint16 = 255
    fmt.Println(p0, p2)
    switch "\n\t" {
    case "hello":
        if (p1 && p1) {
            var v13 uint = p0
            var v14 int16 = 0
            _ = v13
            _ = v14
        } else {
            switch "hello" {
            case "foo bar":
                if false {
                    var v15 uint64 = 42
                    v16 := GenS2{F0: 1}
                    _ = v15
                    _ = v16
                }
            case "a":
                switch 9223372036854775807 {
                case 9223372036854775807:
                    var v17 int = 32767
                    _ = v17
                case 0:
                    var v18 float64 = float64(float32((-2.25 - 1e10)))
                    _ = v18
                }
            case "":
                p2 = p2
            }
            fmt.Println(p0, v12)
            if p1 {
                fmt.Println(p2, p0, p1)
                v12 -= (v12 ^ 42)
            }
        }
    case "a":
        p2 = p2
    }
    switch -125 {
    case -786:
        for fi19 := 0; fi19 < 9; fi19++ {
            fmt.Println(v12, p1)
            _ = fi19
        }
    case -2147483648:
        if p1 {
            v12 -= (v12 & (v12 / (((^(v12))) | 1)))
            v20 := GenS3{F0: 127, F1: uint64(p2)}
            switch int(p2) {
            case 32767:
                p2 += 801
            case -197:
                var v21 int8 = 42
                _ = v21
            case 4294967295:
                v22 := GenS4{F0: 18446744073709551615, F1: (v12 % ((436) | 1)), F2: 128}
                _ = v22
            default:
                for fi23 := 0; fi23 < 15; fi23++ {
                    var v24 uint = 65535
                    _ = fi23
                    _ = v24
                }
            }
            _ = v20
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    var v3 int8 = 0
    var v4 float32 = 1.5
    v5 := make(map[string]int)
    fmt.Println(v4)
    if false {
        v3 = -128
        for fi25 := 0; fi25 < 5; fi25++ {
            if false {
                for fi26 := 0; fi26 < 9; fi26++ {
                    var v27 int64 = 379
                    switch -2147483648 {
                    case 2147483647:
                        switch "foo bar" {
                        case "hello":
                            if (!(2147483648 < 42)) {
                                if false {
                                    if (!false) {
                                        v3 -= (v3 & (^(v3)))
                                    } else {
                                        fmt.Println(v3, fi26, v27)
                                    }
                                    if (v27 == 9223372036854775807) {
                                        v28 := GenS2{F0: 24}
                                        v29 := []int{fi25, (-(fi26)), fi26}
                                        _ = v28
                                        _ = v29
                                    } else {
                                        _ = len(v1)
                                        v30 := []int{fi25}
                                        v3 = v3
                                        if true {
                                            _ = len(v5)
                                            v3 -= (v3 % (((v3 ^ (v3 % (((^(v3))) | 1)))) | 1))
                                        } else {
                                            v3 -= v3
                                            fmt.Println(v27)
                                            var v31 bool = (true || (v27 < v27))
                                            _ = v31
                                        }
                                        _ = v30
                                    }
                                    var v32 uint = 255
                                    _ = v32
                                } else {
                                    if ((v27 << ((42) & 63)) <= v27) {
                                        fmt.Println(v3, v3)
                                    } else {
                                        var v33 int16 = -32768
                                        var v34 int16 = v33
                                        var v35 float32 = v4
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                    if (float64(((0.0 * ((1.5 + v4) * v4)) + -2.25)) == 0.1) {
                                        v4 += v4
                                        switch (fi26 | (fi25 & (-(fi26)))) {
                                        case -1:
                                            if false {
                                                var v36 uint64 = 293
                                                _ = v36
                                            }
                                        case 330:
                                            v5["hello"] = fi25
                                        default:
                                            var v37 bool = (!(0.0 == 1.5))
                                            _ = v37
                                        }
                                        var v38 float64 = (1.5 - 0.0)
                                        fmt.Println(fi26, v4, fi26)
                                        if false {
                                            v39 := GenS4{F0: 1, F1: 255, F2: (fi25 / ((fi25) | 1))}
                                            _ = v39
                                        } else {
                                            var v40 int8 = v3
                                            var v41 uint64 = genFn1(651, ((-(fi25)) >= fi26), 2)
                                            _ = v40
                                            _ = v41
                                        }
                                        _ = v38
                                    }
                                    if false {
                                        v42 := []int{(fi26 | fi25), (fi26 / ((9223372036854775807) | 1)), fi26}
                                        _ = v42
                                    } else {
                                        v3 = v3
                                        var v43 bool = (v27 >= v27)
                                        _ = v43
                                    }
                                    switch fi25 {
                                    case 207:
                                        v44 := GenS4{F0: 2, F1: 255, F2: (-(fi25))}
                                        _ = v44
                                    case -270:
                                        v45 := []int{-1, (fi25 & fi25), fi25, (fi26 | (fi26 * fi25))}
                                        _ = v45
                                    }
                                    v4 += 3.14159
                                    if (v3 < (v3 - int8(fi25))) {
                                        if false {
                                            var v46 uint32 = 2147483648
                                            v47 := GenS3{F0: 0, F1: 1}
                                            _ = v46
                                            _ = v47
                                        }
                                        _ = len(v1)
                                        if (0 == 1) {
                                            v48 := []int{fi25, (fi25 | fi26), (fi25 - (fi25 & (fi26 % ((-185) | 1))))}
                                            v49 := GenS5{F0: 255, F1: false}
                                            if len(v48) > 0 {
                                                v48[0] = 451
                                            }
                                            _ = v48
                                            _ = v49
                                        } else {
                                            var v50 int = -2147483648
                                            _ = v50
                                        }
                                        v51 := []int{(fi26 & (fi26 & (fi26 & 32767))), fi25, 9223372036854775807, fi25}
                                        _ = v51
                                    } else {
                                        v4 = v4
                                        var v52 float64 = -2.25
                                        v53 := GenS6{F0: 18446744073709551615, F1: v3, F2: "hello"}
                                        _ = v52
                                        _ = v53
                                    }
                                    _ = len(v1)
                                    if ((42 >= 74) && false) {
                                        fmt.Println(v3, v4, v27)
                                    }
                                    fmt.Println(v3)
                                }
                                v4 += float32((float64(v4) + 1.5))
                            } else {
                                var v54 float32 = (((1.5 - v4) - (float32(1.5) + (((0.0 - 0.1) - v4) / 1.5))) / float32(((0.0 + 1.5) - 1e10)))
                                fmt.Println(fi26, v54, v3)
                                switch ("foo bar" + ("hello" + "\n\t")) {
                                case "hello":
                                    v55 := v2["a"]
                                    var v56 bool = true
                                    var v57 string = v55
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                case "\n\t":
                                    var v58 uint16 = 0
                                    _ = v58
                                case "a":
                                    if true {
                                        switch 0 {
                                        case 127:
                                            var v59 uint8 = 128
                                            _ = v59
                                        case 313:
                                            switch (-(fi25)) {
                                            case 0:
                                                v3 += v3
                                            case 32767:
                                                if true {
                                                    v60 := GenS7{F0: 42, F1: 2}
                                                    v61 := []int{fi25}
                                                    _ = v60
                                                    _ = v61
                                                } else {
                                                    v62 := GenS3{F0: 0, F1: 18446744073709551615}
                                                    var v63 uint = 1
                                                    v3 = 1
                                                    _ = v62
                                                    _ = v63
                                                }
                                            case 1:
                                                v3 = v3
                                            default:
                                                fmt.Println(v54)
                                            }
                                        case 4294967295:
                                            v54 -= float32((3.14159 * -2.25))
                                        default:
                                            v64 := []int{(fi26 * int(v3)), fi25, (fi26 - fi25)}
                                            _ = v64
                                        }
                                        v54 = v4
                                    } else {
                                        v65 := []int{1, fi25, fi26, 4294967295}
                                        switch "foo bar" {
                                        case "foo bar":
                                            if (1 < 0) {
                                                var v66 uint = 343
                                                v67 := GenS8{F0: (v3 - v3), F1: (fi25 / ((-1) | 1)), F2: 65535}
                                                v68 := GenS6{F0: (v66 % ((v66) | 1)), F1: (v3 ^ v3), F2: "\n\t"}
                                                _ = v66
                                                _ = v67
                                                _ = v68
                                            }
                                        case "":
                                            var v69 string = "foo bar"
                                            _ = v69
                                        }
                                        _ = v65
                                    }
                                }
                                if true {
                                    v3 -= v3
                                    var v70 float32 = 1.5
                                    v71 := GenS9{F0: v27}
                                    _ = v70
                                    _ = v71
                                } else {
                                    var v72 uint16 = 2
                                    v54 = v54
                                    _ = v72
                                }
                                var v73 int16 = 0
                                _ = v54
                                _ = v73
                            }
                        case "foo bar":
                            v74 := GenS10{F0: -801, F1: 292}
                            _ = v74
                        case "":
                            var v75 float64 = 1e10
                            _ = v75
                        default:
                            switch "a" {
                            case "foo bar":
                                v4 -= float32(1e10)
                                v76 := []int{(fi25 << ((2) & 63)), (fi26 & fi25)}
                                if true {
                                    var v77 uint64 = 4294967295
                                    _ = v77
                                }
                                var v78 int16 = -1
                                if (!(fi25 > (fi26 / ((fi26) | 1)))) {
                                    var v79 uint8 = 255
                                    _ = v79
                                } else {
                                    var v80 int = fi25
                                    switch "foo bar" {
                                    case "foo bar":
                                        if len(v76) >= 2 {
                                            v76 = v76[1:len(v76)]
                                        }
                                    case "a":
                                        v4 = v4
                                    }
                                    _ = v80
                                }
                                _ = v76
                                _ = v78
                            case "":
                                v81 := []int{fi26}
                                _ = v81
                            default:
                                if false {
                                    var v82 float32 = v4
                                    var v83 uint16 = 32767
                                    _ = v82
                                    _ = v83
                                }
                            }
                        }
                    default:
                        var v84 float32 = float32(0.0)
                        _ = v84
                    }
                    v4 = v4
                    v3 -= v3
                    v85 := []int{(fi25 >> ((430) & 63)), 32768, (fi25 * fi25)}
                    _ = fi26
                    _ = v27
                    _ = v85
                }
            }
            if true {
                var v86 uint = 921
                switch 32767 {
                case 790:
                    var v87 uint32 = uint32(v86)
                    _ = v87
                case 4294967295:
                    if true {
                        var v88 float64 = ((0.1 - -2.25) + -2.25)
                        v86 += 85
                        v3 = (v3 << (((v86 ^ v86)) & 63))
                        _ = v88
                    }
                }
                _ = v86
            }
            switch (^(fi25)) {
            case 714:
                v1[fi25] = false
            case 42:
                if false {
                    switch 2147483647 {
                    case -765:
                        v89 := GenS1{F0: -1}
                        _ = v89
                    case -2147483648:
                        fmt.Println(v4, v3)
                    default:
                        v90 := GenS4{F0: genFn1(4294967295, false, 802), F1: 42, F2: fi25}
                        _ = v90
                    }
                    fmt.Println(v3, fi25, v4)
                } else {
                    var v91 uint32 = 293
                    v92 := []int{492}
                    var v93 uint8 = 128
                    _ = v91
                    _ = v92
                    _ = v93
                }
            case -565:
                v94 := GenS11{F0: 504, F1: fi25, F2: -1}
                _ = v94
            default:
                if true {
                    for fi95 := 0; fi95 < 9; fi95++ {
                        if (0 != 18446744073709551615) {
                            _ = len(v1)
                            v4 -= v4
                        } else {
                            switch "\n\t" {
                            case "":
                                v96 := GenS7{F0: 2147483647, F1: 2}
                                _ = v96
                            case "foo bar":
                                switch "hello" {
                                case "":
                                    v97 := v2[("foo bar" + "foo bar")]
                                    _ = v97
                                case "hello":
                                    v98 := []int{-636, (fi95 ^ 32767), (fi95 ^ 0), fi95}
                                    _ = v98
                                case "\n\t":
                                    switch fi25 {
                                    case 0:
                                        v99 := GenS8{F0: v3, F1: (fi25 >> ((2) & 63)), F2: 2147483648}
                                        _ = v99
                                    case -969:
                                        var v100 float32 = (-2.25 / (v4 - (v4 / v4)))
                                        _ = v100
                                    default:
                                        if (!((422 < 2) || (0 == 1))) {
                                            v4 = (v4 + 3.14159)
                                        }
                                    }
                                default:
                                    v4 += v4
                                }
                            case "a":
                                var v101 float32 = v4
                                _ = v101
                            }
                        }
                        v4 = v4
                        if (65535 == 2147483648) {
                            switch (fi95 - (-(fi95))) {
                            case 952:
                                if false {
                                    _ = len(v1)
                                    v102 := GenS4{F0: 18446744073709551615, F1: 255, F2: int(v3)}
                                    v103 := []int{(fi95 % ((128) | 1)), (fi25 + (fi25 % ((fi95) | 1))), (-(fi25))}
                                    v4 -= 0.1
                                    v104 := GenS7{F0: 2147483647, F1: 161}
                                    _ = v102
                                    _ = v103
                                    _ = v104
                                } else {
                                    v4 = (float32(0.1) * ((v4 - v4) + v4))
                                    var v105 float32 = (3.14159 / v4)
                                    _ = v105
                                }
                            case 128:
                                fmt.Println(fi25, fi95, v3)
                            }
                            if (!(0 == 65535)) {
                                var v106 int8 = v3
                                v107 := GenS11{F0: -2147483648, F1: (fi25 >> ((121) & 63)), F2: 1}
                                _ = v106
                                _ = v107
                            }
                        } else {
                            v4 -= (1.5 * -2.25)
                            if true {
                                fmt.Println(v4, fi95)
                                if false {
                                    var v108 uint16 = 0
                                    var v109 int64 = 42
                                    var v110 uint = 1
                                    _ = v108
                                    _ = v109
                                    _ = v110
                                } else {
                                    v111 := []int{fi95, (-(fi25))}
                                    var v112 uint = 2
                                    v113 := v5[""]
                                    _ = v111
                                    _ = v112
                                    _ = v113
                                }
                            } else {
                                v3 += (v3 + (v3 << ((65535) & 63)))
                                v114 := []int{fi95}
                                fmt.Println(fi25, fi95, v4)
                                _ = v114
                            }
                        }
                        _ = fi95
                    }
                } else {
                    v115 := GenS12{F0: v3}
                    var v116 int64 = 42
                    fmt.Println(v4, v116, v116)
                    _ = v115
                    _ = v116
                }
            }
            _ = fi25
        }
        for fi117 := 0; fi117 < 5; fi117++ {
            var v118 uint16 = 2
            v4 = v4
            _ = fi117
            _ = v118
        }
    } else {
        if (!(!true)) {
            v3 -= v3
        } else {
            v4 -= v4
            v119 := GenS11{F0: 741, F1: -1, F2: 1}
            _ = v119
        }
        v120 := GenS13{F0: 0, F1: 1}
        _ = v120
    }
    for fi121 := 0; fi121 < 10; fi121++ {
        var v122 float32 = (v4 + v4)
        v123 := []int{(fi121 << ((4294967295) & 63)), 119, fi121}
        _ = fi121
        _ = v122
        _ = v123
    }
    if true {
        if false {
            for fi124 := 0; fi124 < 16; fi124++ {
                var v125 string = ""
                v126 := []int{(fi124 << ((468) & 63))}
                _ = fi124
                _ = v125
                _ = v126
            }
            for fi127 := 0; fi127 < 2; fi127++ {
                if (-2.25 > (3.14159 * 0.1)) {
                    v128 := []int{(fi127 ^ fi127)}
                    var v129 int16 = 1
                    var v130 string = ""
                    _ = v128
                    _ = v129
                    _ = v130
                } else {
                    var v131 uint8 = 2
                    v4 += v4
                    delete(v1, fi127)
                    _ = v131
                }
                v3 += v3
                _ = fi127
            }
            if true {
                switch "a" {
                case "a":
                    fmt.Println(v3)
                case "foo bar":
                    v3 = -128
                default:
                    var v132 uint32 = 0
                    _ = v132
                }
                if true {
                    for fi133 := 0; fi133 < 2; fi133++ {
                        fmt.Println(fi133, v4)
                        _ = fi133
                    }
                    if false {
                        switch 230 {
                        case 32768:
                            var v134 float32 = (((0.0 - v4) - -2.25) * float32(1e10))
                            _ = v134
                        case -1:
                            v4 -= v4
                        case 127:
                            v135 := GenS14{F0: 880, F1: 2}
                            _ = v135
                        default:
                            v136 := []int{1, 4294967295, -1}
                            _ = v136
                        }
                        var v137 float64 = 1e10
                        _ = len(v1)
                        _ = v137
                    }
                } else {
                    var v138 bool = (127 != v3)
                    v139 := v1[-1]
                    _ = v138
                    _ = v139
                }
                _ = len(v5)
            }
        } else {
            v140 := v1[-530]
            _ = v140
        }
    } else {
        v4 -= 0.0
        v141 := []int{-105, 42}
        switch 1 {
        case 0:
            var v142 uint32 = 42
            _ = v142
        case 32767:
            _ = len(v141)
        default:
            var v143 string = "a"
            _ = v143
        }
        _ = v141
    }
    var v144 int = 127
    var v145 int16 = -554
    if true {
        v146 := GenS15{F0: 538, F1: 9223372036854775807, F2: v3}
        v144 -= (v144 >> ((522) & 63))
        _ = v146
    }
    v145 -= (v145 ^ v145)
    v144 -= (v144 & v144)
    fmt.Println(v4, v3, v144)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v144:", v144)
    fmt.Println("v145:", v145)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v144
    _ = v145
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
