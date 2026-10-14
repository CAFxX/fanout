package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int64; F2 uint }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int32; F1 string }
type GenS7 struct { F0 uint64; F1 uint32 }
type GenS8 struct { F0 int16; F1 uint64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int64; F1 uint64; F2 int32 }
type GenS11 struct { F0 uint8; F1 int }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int32; F1 int64 }
type GenS14 struct { F0 int64; F1 uint16 }
type GenS15 struct { F0 int32; F1 string }
type GenS16 struct { F0 uint32; F1 uint32; F2 int64 }
func genFn0(p0 int64) string {
    var v6 int8 = 1
    var v7 uint8 = 2
    _ = p0
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    var v3 uint64 = 221
    var v4 int16 = 1
    var v5 bool = false
    switch "a" {
    case "a":
        for fi8 := 0; fi8 < 18; fi8++ {
            if v5 {
                if v5 {
                    var v9 float32 = float32(((1e10 / (1.5 * float64(float32(float64((0.1 / 1.5)))))) / (-2.25 * 0.1)))
                    _ = v9
                } else {
                    if ((!v5) && (fi8 != (fi8 % ((fi8) | 1)))) {
                        for fi10 := 0; fi10 < 11; fi10++ {
                            v11 := GenS1{F0: 128}
                            _ = fi10
                            _ = v11
                        }
                    } else {
                        v12 := GenS2{F0: ("a" + "a"), F1: int(v4)}
                        switch -1 {
                        case 32767:
                            v1[571] = "\n\t"
                        default:
                            v13 := []int{fi8}
                            _ = v13
                        }
                        _ = v12
                    }
                    switch (fi8 ^ -9223372036854775808) {
                    case 1:
                        var v14 uint8 = 1
                        _ = v14
                    case 4294967295:
                        v1[887] = "hello"
                    }
                    v15 := []int{(^(fi8)), (fi8 >> ((117) & 63)), (fi8 - fi8), fi8}
                    _ = v15
                }
                if v5 {
                    v16 := GenS1{F0: 1}
                    _ = v16
                } else {
                    v17 := GenS3{F0: v4}
                    _ = v17
                }
            } else {
                for fi18 := 0; fi18 < 19; fi18++ {
                    v19 := GenS4{F0: (v4 & 187), F1: 268, F2: 0}
                    if (!v5) {
                        v20 := []int{fi18}
                        if (int64(fi8) != -190) {
                            v3 = (v3 - v3)
                            fmt.Println(fi18, v3)
                        } else {
                            fmt.Println(v3, v4, v4)
                            var v21 int32 = -2147483648
                            if false {
                                var v22 int32 = (v21 >> ((542) & 63))
                                v23 := GenS4{F0: v4, F1: -662, F2: 894}
                                _ = v22
                                _ = v23
                            }
                            _ = v21
                        }
                        switch (fi18 / ((1) | 1)) {
                        case -589:
                            fmt.Println(fi18, v4)
                        case 272:
                            if v5 {
                                if len(v20) > 0 {
                                    v20[0] = (fi8 % ((-1) | 1))
                                }
                                switch (fi8 / ((fi18) | 1)) {
                                case -676:
                                    if (false && v5) {
                                        v2["\n\t"] = v5
                                    }
                                case 9223372036854775807:
                                    var v24 uint8 = 0
                                    _ = v24
                                }
                            }
                        case 0:
                            if (!(-2.25 == (float32(0.1) - 3.14159))) {
                                if ((!false) && (32767 > (fi18 ^ fi18))) {
                                    v4 = 0
                                    fmt.Println(fi8, fi8)
                                    if v5 {
                                        if v5 {
                                            v5 = (1 < 4294967295)
                                            var v25 int64 = 42
                                            fmt.Println(v3, v25)
                                            _ = v25
                                        }
                                    } else {
                                        v3 += (v3 * 0)
                                        var v26 int = fi8
                                        _ = v26
                                    }
                                } else {
                                    var v27 int16 = (v4 | v4)
                                    switch (fi8 | fi18) {
                                    case 42:
                                        v28 := GenS5{F0: 0, F1: 255}
                                        _ = v28
                                    case 1:
                                        fmt.Println(v3, v4, v4)
                                    default:
                                        if (int8(v27) <= 40) {
                                            switch ((((("foo bar" + ("" + "hello")) + "\n\t") + (("\n\t" + ("a" + "hello")) + ("\n\t" + "a"))) + ("" + ("foo bar" + ("a" + "foo bar")))) + "foo bar") {
                                            case "":
                                                var v29 uint8 = 0
                                                _ = v29
                                            case "foo bar":
                                                v27 = v4
                                            default:
                                                v4 = (v4 >> ((961) & 63))
                                            }
                                            switch -1 {
                                            case 32767:
                                                v27 = 350
                                            default:
                                                if (-88 > -128) {
                                                    v5 = (!v5)
                                                    v1[2] = "a"
                                                    if v5 {
                                                        var v30 int32 = 0
                                                        fmt.Println(v3)
                                                        fmt.Println(fi18, fi8, v27)
                                                        _ = v30
                                                    }
                                                }
                                            }
                                            v5 = v5
                                        }
                                    }
                                    _ = v27
                                }
                                if (2 <= 42) {
                                    if (4294967295 >= 402) {
                                        v3 -= v3
                                        if false {
                                            v31 := GenS4{F0: 0, F1: -1, F2: 2}
                                            var v32 uint8 = 1
                                            _ = v31
                                            _ = v32
                                        }
                                        if (v5 && (!v5)) {
                                            switch "\n\t" {
                                            case "":
                                                v4 += v4
                                            default:
                                                v4 += 879
                                            }
                                            switch fi18 {
                                            case 525:
                                                fmt.Println(fi8)
                                            }
                                        }
                                    } else {
                                        if v5 {
                                            if (!v5) {
                                                var v33 int = 128
                                                _ = v33
                                            } else {
                                                var v34 uint16 = 2
                                                var v35 float64 = 0.0
                                                if len(v20) > 0 {
                                                    v20[0] = fi8
                                                }
                                                var v36 bool = v5
                                                _ = len(v2)
                                                _ = v34
                                                _ = v35
                                                _ = v36
                                            }
                                            delete(v1, 1)
                                        }
                                        if v5 {
                                            var v37 int16 = (-(v4))
                                            var v38 int64 = 1
                                            v37 = v4
                                            var v39 uint64 = v3
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                        }
                                    }
                                    v5 = v5
                                }
                                var v40 int16 = v4
                                _ = v40
                            } else {
                                var v41 int64 = 911
                                _ = v41
                            }
                        default:
                            if ((true || (1 > 0)) && (1.5 == (3.14159 + 1e10))) {
                                v4 -= (-(v4))
                                v5 = true
                                if ("foo bar" > ("a" + "\n\t")) {
                                    v4 = (v4 | (v4 & (v4 + v4)))
                                }
                            } else {
                                v42 := GenS1{F0: 127}
                                var v43 string = "hello"
                                v44 := GenS1{F0: 188}
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                        }
                        _ = v20
                    } else {
                        if v5 {
                            v45 := []int{(fi8 ^ 32767), 1, fi8}
                            v46 := GenS3{F0: v4}
                            _ = v45
                            _ = v46
                        }
                        v3 += 42
                    }
                    _ = fi18
                    _ = v19
                }
                for fi47 := 0; fi47 < 11; fi47++ {
                    v4 = (v4 & v4)
                    var v48 string = "a"
                    fmt.Println(v3)
                    _ = fi47
                    _ = v48
                }
            }
            for fi49 := 0; fi49 < 8; fi49++ {
                if (!v5) {
                    v5 = v5
                    _ = len(v1)
                } else {
                    var v50 int16 = (-(v4))
                    var v51 uint32 = 65535
                    fmt.Println(v4, v4, v50)
                    _ = v50
                    _ = v51
                }
                _ = fi49
            }
            for fi52 := 0; fi52 < 17; fi52++ {
                if (-1 == 532) {
                    v53 := GenS6{F0: -62, F1: ("\n\t" + "\n\t")}
                    v3 += v3
                    _ = v53
                } else {
                    var v54 bool = true
                    fmt.Println(v4, v5, v54)
                    var v55 float64 = ((0.0 - 1e10) + (((1.5 * (1e10 / (1e10 * 0.1))) * -2.25) / -2.25))
                    if (0 > 18446744073709551615) {
                        var v56 int8 = 111
                        _ = v56
                    } else {
                        v54 = (!(65535 != 2))
                        if false {
                            if false {
                                if (553 <= 2147483647) {
                                    var v57 uint = 0
                                    _ = v57
                                } else {
                                    v58 := []int{fi52, (fi8 % ((fi52) | 1)), fi8}
                                    _ = v58
                                }
                                var v59 int32 = 336
                                var v60 int16 = v4
                                v61 := GenS7{F0: (v3 - (v3 >> ((987) & 63))), F1: 65535}
                                _ = v59
                                _ = v60
                                _ = v61
                            } else {
                                v54 = true
                                v62 := GenS8{F0: 994, F1: (v3 & (v3 & (v3 * 151)))}
                                var v63 int16 = (v4 & v4)
                                _ = v62
                                _ = v63
                            }
                            var v64 uint = 873
                            if ((((((0.0 / 1.0) - ((v55 - (v55 / (3.14159 - v55))) * -2.25)) / v55) - (float64(float32(float64(0.0))) * v55)) - (((0.0 - float64(float32(v55))) - v55) + (v55 + v55))) <= v55) {
                                var v65 uint8 = 127
                                _ = v65
                            } else {
                                if ((v5 && v5) && false) {
                                    v55 -= (1e10 / v55)
                                    v5 = false
                                    fmt.Println(v55, v4, v4)
                                } else {
                                    var v66 int = -9223372036854775808
                                    _ = v66
                                }
                                if (-771 == 1) {
                                    var v67 float64 = -2.25
                                    switch fi52 {
                                    case 32767:
                                        v54 = (22 >= 2)
                                    case -1:
                                        var v68 uint = (v64 % ((v64) | 1))
                                        _ = v68
                                    default:
                                        var v69 bool = false
                                        _ = v69
                                    }
                                    var v70 uint64 = 18446744073709551615
                                    _ = v67
                                    _ = v70
                                }
                                _ = len(v2)
                            }
                            _ = v64
                        }
                    }
                    _ = v54
                    _ = v55
                }
                _ = fi52
            }
            _ = fi8
        }
    default:
        for fi71 := 0; fi71 < 6; fi71++ {
            v3 -= (v3 << ((131) & 63))
            _ = fi71
        }
    }
    v4 = 991
    v72 := GenS9{F0: (v4 % (((v4 >> ((65535) & 63))) | 1))}
    for fi73 := 0; fi73 < 2; fi73++ {
        v74 := []int{(fi73 - fi73), (fi73 * (fi73 - fi73))}
        if true {
            for fi75 := 0; fi75 < 7; fi75++ {
                switch "a" {
                case "":
                    if false {
                        switch fi75 {
                        case 128:
                            v76 := GenS4{F0: (v4 << ((42) & 63)), F1: -1, F2: 568}
                            _ = v76
                        }
                        var v77 bool = ((!false) && v5)
                        v78 := GenS3{F0: v4}
                        _ = v77
                        _ = v78
                    } else {
                        v4 = 1
                        v79 := GenS10{F0: 332, F1: 4294967295, F2: -710}
                        delete(v1, 1)
                        _ = v79
                    }
                case "hello":
                    v3 -= 1
                default:
                    fmt.Println(fi75, v3)
                }
                fmt.Println(v4, fi73)
                _ = fi75
            }
            delete(v1, 0)
        } else {
            v4 = 0
            fmt.Println(fi73, v4)
        }
        _ = fi73
        _ = v74
    }
    var v80 string = "foo bar"
    var v81 float64 = float64(1e10)
    if ((!(-911 != 42)) || (1 > 2147483647)) {
        if true {
            v5 = (!v5)
            switch 42 {
            case 233:
                v1[4294967295] = v80
            case -9223372036854775808:
                v82 := []int{}
                _ = v82
            }
            v1[148] = ((v80 + v80) + v80)
        }
        fmt.Println(v3, v81)
    } else {
        if v5 {
            for fi83 := 0; fi83 < 2; fi83++ {
                for fi84 := 0; fi84 < 15; fi84++ {
                    var v85 uint16 = 0
                    v86 := GenS11{F0: 127, F1: fi83}
                    if v5 {
                        fmt.Println(v80, v4)
                        var v87 uint32 = 2147483647
                        v3 = uint64(v85)
                        _ = v87
                    } else {
                        fmt.Println(v4, v3, fi84)
                    }
                    _ = fi84
                    _ = v85
                    _ = v86
                }
                v88 := GenS12{F0: v5}
                _ = fi83
                _ = v88
            }
            switch -747 {
            case 329:
                var v89 float32 = 3.14159
                _ = v89
            }
            for fi90 := 0; fi90 < 10; fi90++ {
                v91 := GenS13{F0: -1, F1: 9223372036854775807}
                var v92 uint = 42
                _ = fi90
                _ = v91
                _ = v92
            }
        } else {
            var v93 float32 = 1e10
            _ = v93
        }
    }
    v2[v80] = (uint32(v3) <= 42)
    if (!(128 < 32767)) {
        for fi94 := 0; fi94 < 15; fi94++ {
            switch (v80 + v80) {
            case "hello":
                v5 = v5
            case "a":
                v95 := []int{fi94, fi94, (fi94 >> ((106) & 63))}
                _ = v95
            default:
                v5 = true
            }
            v80 = v80
            _ = fi94
        }
        v81 += 0.0
    } else {
        if v5 {
            var v96 uint8 = 136
            v97 := []int{}
            _ = v96
            _ = v97
        } else {
            for fi98 := 0; fi98 < 8; fi98++ {
                v99 := []int{}
                fmt.Println(v4, v80, v80)
                _ = fi98
                _ = v99
            }
            fmt.Println(v81)
        }
        fmt.Println(v4, v81)
        fmt.Println(v81, v5, v81)
    }
    for fi100 := 0; fi100 < 4; fi100++ {
        var v101 string = (v80 + (v80 + (v80 + "foo bar")))
        _ = fi100
        _ = v101
    }
    var v102 bool = v5
    var v103 uint = 18446744073709551615
    fmt.Println(v103, v81)
    v5 = (0 != 128)
    var v104 int = 32768
    for fi105 := 0; fi105 < 18; fi105++ {
        if v102 {
            for fi106 := 0; fi106 < 6; fi106++ {
                if (2 > 0) {
                    v103 += (^(v103))
                    var v107 int = (fi106 + 128)
                    _ = v107
                } else {
                    var v108 int64 = 0
                    if true {
                        fmt.Println(v104, v3)
                        v109 := []int{(v104 << ((v103) & 63)), (v104 >> (((v103 << ((uint(v3)) & 63))) & 63)), (fi106 & (v104 / ((fi105) | 1)))}
                        _ = v109
                    } else {
                        if (v3 < 188) {
                            var v110 uint8 = 101
                            if v5 {
                                delete(v1, 2147483647)
                                var v111 int16 = v4
                                _ = v111
                            }
                            delete(v1, 65535)
                            _ = v110
                        } else {
                            v112 := []int{(-(v104))}
                            v81 = v81
                            switch v80 {
                            case "\n\t":
                                switch ((("foo bar" + v80) + ((v80 + ((v80 + "hello") + ("a" + ""))) + ("foo bar" + ""))) + "hello") {
                                case "\n\t":
                                    v5 = ((-(v108)) != (v108 / ((int64(v4)) | 1)))
                                case "a":
                                    if ("hello" <= v80) {
                                        v113 := GenS12{F0: (v4 != v4)}
                                        _ = v113
                                    } else {
                                        var v114 uint = v103
                                        v114 -= (v103 - (v103 * (v103 >> ((v114) & 63))))
                                        v5 = (true || v5)
                                        _ = v114
                                    }
                                case "":
                                    var v115 int8 = 100
                                    _ = v115
                                default:
                                    v103 += 0
                                }
                            }
                            _ = v112
                        }
                    }
                    v3 = v3
                    _ = v108
                }
                if v102 {
                    if v5 {
                        switch v80 {
                        case "foo bar":
                            if (v102 || ((int64(v104) > -9223372036854775808) && (-2147483648 != 1))) {
                                v116 := GenS14{F0: 0, F1: 1}
                                if (!(v102 && ((v4 << (((^(v103))) & 63)) >= 1))) {
                                    var v117 float32 = 1.5
                                    v102 = (!false)
                                    var v118 int = (fi106 | 2147483647)
                                    _ = v117
                                    _ = v118
                                } else {
                                    switch ((((v80 + (v80 + "")) + (v80 + v80)) + "a") + v80) {
                                    case "hello":
                                        v119 := v2[""]
                                        _ = v119
                                    case "a":
                                        if false {
                                            var v120 int64 = -141
                                            v121 := GenS15{F0: 841, F1: v80}
                                            _ = v120
                                            _ = v121
                                        }
                                    case "foo bar":
                                        v80 = v80
                                    }
                                    v4 = v4
                                    var v122 bool = v5
                                    _ = v122
                                }
                                v123 := []int{(fi105 & 9223372036854775807), fi106, 0, int(v4)}
                                _ = v116
                                _ = v123
                            }
                        case "hello":
                            var v124 float32 = 0.1
                            _ = v124
                        case "a":
                            switch (fi105 >> ((2) & 63)) {
                            case -314:
                                switch v80 {
                                case "hello":
                                    v102 = (127 <= 1)
                                }
                            case -2147483648:
                                v3 += (v3 ^ 0)
                            }
                        default:
                            v125 := GenS2{F0: ("\n\t" + "foo bar"), F1: (^(fi105))}
                            _ = v125
                        }
                        v126 := []int{}
                        if (!((!v102) && (!v102))) {
                            if false {
                                var v127 uint8 = 173
                                if (!false) {
                                    delete(v1, 750)
                                }
                                var v128 float64 = (v81 / v81)
                                v127 = v127
                                var v129 uint = uint(v127)
                                _ = v127
                                _ = v128
                                _ = v129
                            } else {
                                v130 := GenS8{F0: -77, F1: 1}
                                _ = v130
                            }
                            v104 -= (fi106 ^ fi106)
                            v4 -= v4
                        }
                        _ = v126
                    }
                    var v131 uint64 = v3
                    _ = v131
                }
                var v132 int8 = 111
                _ = fi106
                _ = v132
            }
            delete(v1, 2147483648)
        }
        for fi133 := 0; fi133 < 13; fi133++ {
            var v134 int8 = 100
            _ = fi133
            _ = v134
        }
        var v135 uint8 = 255
        _ = fi105
        _ = v135
    }
    if true {
        v103 = (v103 - 0)
    }
    if v5 {
        v81 -= float64((float32(-2.25) * -2.25))
        v136 := GenS15{F0: 1, F1: "foo bar"}
        switch "\n\t" {
        case "\n\t":
            switch v80 {
            case "":
                if v5 {
                    var v137 float64 = ((v81 / 1e10) + v81)
                    if (!v102) {
                        fmt.Println(v137, v80)
                        _ = len(v2)
                        if (936 < v3) {
                            var v138 uint16 = 2
                            _ = v138
                        } else {
                            if v5 {
                                var v139 uint = (v103 - 937)
                                fmt.Println(v103, v80)
                                _ = v139
                            }
                        }
                    }
                    _ = v137
                } else {
                    v140 := GenS16{F0: 721, F1: 2147483647, F2: -620}
                    v141 := GenS4{F0: 909, F1: 1, F2: 792}
                    fmt.Println(v80, v80)
                    _ = v140
                    _ = v141
                }
            default:
                v80 = "hello"
            }
        case "foo bar":
            var v142 int32 = 0
            _ = v142
        }
        _ = v136
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v102
    _ = v103
    _ = v104
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v72
    _ = v80
    _ = v81
}
