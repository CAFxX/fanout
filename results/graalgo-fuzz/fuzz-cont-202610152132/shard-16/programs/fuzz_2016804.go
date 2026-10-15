package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 int8; F2 int }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 string; F1 int }
type GenS9 struct { F0 int32; F1 uint16 }
type GenS10 struct { F0 string; F1 uint32; F2 string }
type GenS11 struct { F0 uint; F1 bool }
type GenS12 struct { F0 uint16; F1 int64 }
func genFn0() int16 {
    v7 := GenS1{F0: 42, F1: 187}
    v8 := GenS2{F0: 100, F1: "a"}
    _ = v7
    _ = v8
    return -203
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 830
    v3 := make(map[string]int)
    var v4 int32 = 1
    var v5 float64 = 1.5
    var v6 int32 = (v4 * (v4 & v4))
    fmt.Println(v6, v6)
    var v9 uint64 = 18446744073709551615
    if (100 > 0) {
        switch "hello" {
        case "a":
            if true {
                var v10 uint = 87
                var v11 string = "hello"
                _ = v10
                _ = v11
            }
        case "hello":
            switch "\n\t" {
            case "\n\t":
                _ = len(v1)
            }
        default:
            switch 2147483647 {
            case -943:
                fmt.Println(v4)
            case 32768:
                v4 = (-(v6))
            default:
                var v12 uint = 85
                _ = v12
            }
        }
        if (-2.25 == float64(0.0)) {
            delete(v3, ("hello" + "foo bar"))
            fmt.Println(v5)
            for fi13 := 0; fi13 < 8; fi13++ {
                for fi14 := 0; fi14 < 5; fi14++ {
                    var v15 float64 = 1.5
                    v4 -= (v4 - (v6 >> ((729) & 63)))
                    _ = fi14
                    _ = v15
                }
                v4 = -632
                for fi16 := 0; fi16 < 15; fi16++ {
                    var v17 uint = 829
                    if true {
                        v2 = v2
                        _ = len(v3)
                        v3[""] = fi13
                        fmt.Println(v6, v17)
                        var v18 uint32 = 2147483647
                        v2 = (v2 / ((v2) | 1))
                        switch "hello" {
                        case "":
                            var v19 int32 = (v4 << ((v17) & 63))
                            _ = v19
                        case "\n\t":
                            v20 := GenS3{F0: 1000, F1: genFn0()}
                            _ = v20
                        case "hello":
                            v2 -= v2
                        }
                        v21 := GenS2{F0: -67, F1: "\n\t"}
                        switch fi16 {
                        case -434:
                            v22 := GenS4{F0: 0, F1: (^(v18)), F2: v18}
                            _ = v22
                        }
                        _ = v18
                        _ = v21
                    } else {
                        v5 = (float64((3.14159 + (3.14159 / (0.0 + 1e10)))) + float64(float32(-2.25)))
                        v23 := GenS5{F0: v9}
                        v2 = 2
                        _ = v23
                    }
                    if (3.14159 > 0.0) {
                        if (1 < v17) {
                            var v24 float32 = (1e10 - ((0.1 - ((((3.14159 - 0.1) - (0.0 + -2.25)) - 3.14159) * ((float32((3.14159 - v5)) + 0.0) / (1e10 * (((((0.0 - (3.14159 + (0.0 * (float32(1.5) + (1.5 * 3.14159))))) - -2.25) * ((0.0 - 1.5) / ((((0.0 * float32(1e10)) - float32(float64((-2.25 * 1e10)))) / (0.1 / 1.0)) - 1e10))) / 3.14159) / 0.1))))) - (0.1 - (-2.25 - ((3.14159 + (0.0 * (1e10 + 3.14159))) * (((float32(v5) / 1e10) + 1e10) + (((-2.25 / (1.5 - (0.1 - (((1.5 * 1.5) / 1.0) * (float32(((((v5 * float64((0.0 / 1.0))) + (1.5 + 1e10)) - v5) / (v5 / v5))) - (0.1 * (1e10 / (1e10 * -2.25)))))))) + 1e10) - 1.5)))))))
                            var v25 uint = v17
                            var v26 float32 = (v24 - v24)
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        if true {
                            var v27 uint = v17
                            if (!true) {
                                v17 += (v27 & 255)
                                var v28 uint8 = 2
                                v29 := []int{fi16, fi13}
                                v30 := GenS6{F0: 717, F1: 42, F2: fi16}
                                var v31 uint32 = 2
                                _ = v28
                                _ = v29
                                _ = v30
                                _ = v31
                            } else {
                                var v32 int16 = genFn0()
                                v17 = (v17 | uint(v9))
                                _ = v32
                            }
                            v17 = (^(v27))
                            _ = v27
                        } else {
                            fmt.Println(v17, v17, v6)
                        }
                    } else {
                        v2 -= 0
                        var v33 uint = (v17 % ((v17) | 1))
                        _ = v33
                    }
                    _ = fi16
                    _ = v17
                }
                _ = fi13
            }
        } else {
            if false {
                if true {
                    if (false || false) {
                        var v34 int = -803
                        _ = v34
                    } else {
                        var v35 uint32 = 4294967295
                        _ = v35
                    }
                    switch (("hello" + "") + "") {
                    case "a":
                        if true {
                            for fi36 := 0; fi36 < 15; fi36++ {
                                var v37 string = "\n\t"
                                _ = fi36
                                _ = v37
                            }
                        } else {
                            if true {
                                if (v2 >= (v2 & v2)) {
                                    for fi38 := 0; fi38 < 18; fi38++ {
                                        var v39 uint64 = (v9 & (v9 % ((v9) | 1)))
                                        v40 := GenS7{F0: -93, F1: 996}
                                        _ = fi38
                                        _ = v39
                                        _ = v40
                                    }
                                    v41 := GenS8{F0: ("a" + "a"), F1: 32768}
                                    _ = v41
                                } else {
                                    v42 := GenS2{F0: int8(v2), F1: "foo bar"}
                                    var v43 int32 = (v6 << ((169) & 63))
                                    _ = v42
                                    _ = v43
                                }
                                fmt.Println(v6, v6, v9)
                                if true {
                                    fmt.Println(v4, v4, v4)
                                }
                            } else {
                                var v44 uint32 = 0
                                v9 -= 18446744073709551615
                                fmt.Println(v5, v9)
                                _ = v44
                            }
                        }
                    case "foo bar":
                        for fi45 := 0; fi45 < 20; fi45++ {
                            v4 = v4
                            for fi46 := 0; fi46 < 6; fi46++ {
                                v6 += (v4 - 2147483647)
                                _ = fi46
                            }
                            _ = fi45
                        }
                    }
                } else {
                    for fi47 := 0; fi47 < 14; fi47++ {
                        v5 = ((((((v5 - v5) - 0.0) - -2.25) - float64((1e10 * ((1.5 / 1e10) * 0.0)))) / ((0.0 * float64(1.5)) - (float64(float32(v5)) / (v5 - 3.14159)))) * (3.14159 / v5))
                        var v48 uint32 = 42
                        for fi49 := 0; fi49 < 4; fi49++ {
                            v50 := GenS9{F0: 616, F1: v2}
                            delete(v1, -662)
                            var v51 bool = true
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                        _ = fi47
                        _ = v48
                    }
                    v9 -= v9
                }
                if (("foo bar" + "a") < "a") {
                    switch 4294967295 {
                    case -589:
                        if (v5 < 1.5) {
                            v52 := []int{2147483647, -1, 32767, 2147483647}
                            if len(v52) > 0 {
                                v52[0] = 32768
                            }
                            switch ("" + "hello") {
                            case "\n\t":
                                v6 += 42
                            case "hello":
                                fmt.Println(v4, v6)
                            case "":
                                v2 = (v2 % (((^(v2))) | 1))
                            }
                            _ = v52
                        }
                    default:
                        var v53 int = 4294967295
                        _ = v53
                    }
                    var v54 int16 = -1
                    _ = v54
                }
            } else {
                switch ((("" + ("" + ("\n\t" + "hello"))) + "hello") + "\n\t") {
                case "a":
                    fmt.Println(v5)
                case "hello":
                    if true {
                        var v55 float32 = (float32(0.1) - 0.1)
                        _ = v55
                    }
                case "":
                    if ("hello" != "a") {
                        switch "a" {
                        case "hello":
                            delete(v3, "hello")
                        case "\n\t":
                            for fi56 := 0; fi56 < 9; fi56++ {
                                var v57 bool = false
                                var v58 int8 = -128
                                var v59 bool = v57
                                _ = fi56
                                _ = v57
                                _ = v58
                                _ = v59
                            }
                        default:
                            var v60 uint = 462
                            _ = v60
                        }
                        switch "foo bar" {
                        case "a":
                            if (42 > 0) {
                                if (!true) {
                                    v5 += (-2.25 - v5)
                                    v61 := []int{-1, 45}
                                    v5 = float64(0.0)
                                    _ = v61
                                }
                            }
                        case "":
                            v2 = v2
                        default:
                            switch (("" + ("" + "foo bar")) + ((("" + ("\n\t" + "foo bar")) + "foo bar") + "a")) {
                            case "foo bar":
                                for fi62 := 0; fi62 < 2; fi62++ {
                                    if (true || false) {
                                        fmt.Println(v6, v4, v4)
                                    }
                                    _ = fi62
                                }
                            }
                        }
                    }
                }
                for fi63 := 0; fi63 < 17; fi63++ {
                    v6 = (-(v4))
                    _ = fi63
                }
            }
            if true {
                v64 := GenS5{F0: (v9 % ((712) | 1))}
                switch 2147483647 {
                case -9223372036854775808:
                    switch 863 {
                    case -1:
                        if false {
                            var v65 float32 = float32((v5 - v5))
                            v2 = v2
                            for fi66 := 0; fi66 < 17; fi66++ {
                                for fi67 := 0; fi67 < 15; fi67++ {
                                    if (705 >= 32767) {
                                        var v68 float64 = v5
                                        _ = v68
                                    }
                                    var v69 string = ((("" + "") + "a") + "\n\t")
                                    v70 := v3[(("foo bar" + "a") + v69)]
                                    _ = fi67
                                    _ = v69
                                    _ = v70
                                }
                                _ = fi66
                            }
                            _ = v65
                        }
                    case 32768:
                        fmt.Println(v6, v9, v6)
                    }
                default:
                    v2 = v2
                }
                if true {
                    for fi71 := 0; fi71 < 13; fi71++ {
                        switch "foo bar" {
                        case "\n\t":
                            v6 -= (v6 / ((-1) | 1))
                        case "hello":
                            var v72 string = ""
                            _ = v72
                        case "a":
                            v4 += 740
                        }
                        var v73 uint64 = (v9 >> ((1) & 63))
                        for fi74 := 0; fi74 < 5; fi74++ {
                            v9 -= (^(v73))
                            var v75 string = "hello"
                            var v76 bool = (((!true) && false) && false)
                            _ = fi74
                            _ = v75
                            _ = v76
                        }
                        _ = fi71
                        _ = v73
                    }
                    v5 = v5
                    switch 4294967295 {
                    case -9223372036854775808:
                        if false {
                            switch ("" + "foo bar") {
                            case "a":
                                if true {
                                    v2 = (v2 | v2)
                                    var v77 string = ("hello" + "\n\t")
                                    v2 = (v2 * (^(v2)))
                                    _ = v77
                                } else {
                                    if (-2147483648 <= v4) {
                                        v78 := GenS3{F0: 1, F1: -540}
                                        _ = v78
                                    } else {
                                        if (32768 <= -837) {
                                            var v79 float32 = (0.1 + 1.5)
                                            for fi80 := 0; fi80 < 2; fi80++ {
                                                var v81 float32 = v79
                                                _ = fi80
                                                _ = v81
                                            }
                                            v82 := []int{1, 2147483647, 9223372036854775807}
                                            _ = v79
                                            _ = v82
                                        } else {
                                            v2 += (^(v2))
                                        }
                                        if false {
                                            if (4294967295 != 2147483647) {
                                                var v83 int32 = -548
                                                _ = v83
                                            } else {
                                                if false {
                                                    var v84 int32 = (v4 - (v6 ^ v6))
                                                    var v85 int16 = genFn0()
                                                    fmt.Println(v9, v6, v5)
                                                    _ = v84
                                                    _ = v85
                                                }
                                                var v86 uint64 = 1
                                                _ = v86
                                            }
                                        } else {
                                            var v87 uint16 = v2
                                            v2 -= (^(v87))
                                            _ = v87
                                        }
                                        switch "a" {
                                        case "\n\t":
                                            if true {
                                                v88 := GenS2{F0: -114, F1: "a"}
                                                var v89 bool = false
                                                v4 -= v6
                                                _ = v88
                                                _ = v89
                                            }
                                        case "a":
                                            v90 := []int{128}
                                            _ = v90
                                        case "foo bar":
                                            if false {
                                                for fi91 := 0; fi91 < 19; fi91++ {
                                                    var v92 int32 = int32(v9)
                                                    _ = fi91
                                                    _ = v92
                                                }
                                                _ = len(v3)
                                            } else {
                                                v4 = v6
                                                v93 := GenS2{F0: 0, F1: ""}
                                                _ = v93
                                            }
                                        default:
                                            if ((((float32(1e10) * 1.5) + 1e10) * 1e10) != -2.25) {
                                                var v94 uint16 = (v2 << ((65535) & 63))
                                                _ = v94
                                            }
                                        }
                                    }
                                }
                            }
                        } else {
                            v95 := GenS2{F0: 14, F1: ""}
                            fmt.Println(v5)
                            _ = v95
                        }
                    case 2147483647:
                        fmt.Println(v9, v2)
                    case 4294967295:
                        if false {
                            switch 9223372036854775807 {
                            case 128:
                                if false {
                                    if (((((float32(float64(1e10)) + 0.1) - float32(v5)) + 0.0) - (0.0 / 1.0)) > 1e10) {
                                        var v96 int16 = 1000
                                        _ = v96
                                    } else {
                                        var v97 bool = (42 <= -1)
                                        _ = v97
                                    }
                                    switch "" {
                                    case "hello":
                                        v98 := GenS10{F0: "\n\t", F1: 42, F2: "foo bar"}
                                        _ = v98
                                    }
                                    v9 -= 1
                                } else {
                                    var v99 uint16 = 160
                                    if true {
                                        if (false && ((true && ((1 == 0) && false)) || false)) {
                                            var v100 string = "\n\t"
                                            _ = v100
                                        }
                                    }
                                    _ = v99
                                }
                            case 717:
                                switch 128 {
                                case 228:
                                    v101 := GenS8{F0: "a", F1: -450}
                                    _ = v101
                                case 309:
                                    for fi102 := 0; fi102 < 13; fi102++ {
                                        v5 -= v5
                                        var v103 float64 = (float64(1e10) - (v5 / 0.1))
                                        delete(v3, "foo bar")
                                        _ = fi102
                                        _ = v103
                                    }
                                case 4294967295:
                                    fmt.Println(v6)
                                default:
                                    if (v4 < 235) {
                                        fmt.Println(v2, v2, v2)
                                        v2 -= (v2 - 1)
                                    } else {
                                        v104 := GenS11{F0: 4294967295, F1: false}
                                        _ = v104
                                    }
                                }
                            }
                            v6 += -389
                        } else {
                            var v105 int32 = v6
                            _ = v105
                        }
                    default:
                        _ = len(v1)
                    }
                }
                _ = v64
            } else {
                var v106 uint8 = 127
                for fi107 := 0; fi107 < 20; fi107++ {
                    if (!false) {
                        for fi108 := 0; fi108 < 10; fi108++ {
                            if (!false) {
                                v109 := []int{32767, fi107, -963, (^(fi107))}
                                v1[fi108] = "foo bar"
                                fmt.Println(v5, fi108, v4)
                                _ = v109
                            } else {
                                var v110 int = fi107
                                if ((((v5 - (v5 + v5)) * 1.5) * (v5 / v5)) != 1.5) {
                                    v111 := []int{}
                                    if (739 < 18446744073709551615) {
                                        v106 += (v106 - v106)
                                    }
                                    _ = v111
                                }
                                if true {
                                    v112 := GenS9{F0: (v6 ^ v4), F1: (v2 ^ (v2 / ((v2) | 1)))}
                                    _ = v112
                                } else {
                                    var v113 int = fi107
                                    v6 += (v4 - -308)
                                    fmt.Println(v6)
                                    _ = v113
                                }
                                if ((v5 / v5) >= 1.5) {
                                    var v114 float32 = 0.0
                                    _ = v114
                                }
                                _ = v110
                            }
                            _ = fi108
                        }
                    }
                    v6 += v4
                    _ = fi107
                }
                _ = v106
            }
            for fi115 := 0; fi115 < 8; fi115++ {
                for fi116 := 0; fi116 < 14; fi116++ {
                    if true {
                        v117 := GenS6{F0: 9, F1: -85, F2: int(v2)}
                        if ("a" <= "a") {
                            v9 = v9
                        }
                        if (((1.5 * ((3.14159 * float64(0.0)) / v5)) / (3.14159 - v5)) >= (v5 * (float64(3.14159) + (((v5 - ((1e10 / v5) + 0.1)) + (0.0 * ((v5 * -2.25) * (-2.25 + (v5 / ((0.0 - (((float64(0.1) * 1.5) / float64(3.14159)) - v5)) + v5)))))) + float64(0.0))))) {
                            v118 := GenS9{F0: v6, F1: (^(v2))}
                            if (fi116 < fi116) {
                                var v119 int16 = -32768
                                v120 := GenS7{F0: 9223372036854775807, F1: 4294967295}
                                _ = v119
                                _ = v120
                            } else {
                                v9 = v9
                                v9 = 1
                                switch ("foo bar" + ("" + ("hello" + ""))) {
                                case "a":
                                    v121 := []int{(-(fi116)), (fi115 + (fi116 | int(v2))), (fi115 | (fi115 / ((fi115) | 1)))}
                                    _ = v121
                                case "foo bar":
                                    var v122 int32 = (v6 - v4)
                                    _ = v122
                                default:
                                    v123 := []int{fi115}
                                    _ = v123
                                }
                                switch fi115 {
                                case 128:
                                    if (!((false && (false || true)) && true)) {
                                        v124 := GenS8{F0: "hello", F1: (fi115 ^ 1)}
                                        v125 := []int{}
                                        _ = v124
                                        _ = v125
                                    }
                                case -2147483648:
                                    var v126 float32 = (3.14159 * -2.25)
                                    _ = v126
                                case 42:
                                    v127 := GenS3{F0: int16(v2), F1: -32768}
                                    _ = v127
                                }
                                fmt.Println(v2, v2, v5)
                            }
                            v2 = v2
                            _ = v118
                        }
                        _ = v117
                    }
                    var v128 int = (fi116 ^ fi115)
                    var v129 float64 = v5
                    v130 := []int{}
                    _ = fi116
                    _ = v128
                    _ = v129
                    _ = v130
                }
                var v131 uint32 = 2147483648
                var v132 int = -9223372036854775808
                _ = fi115
                _ = v131
                _ = v132
            }
        }
    } else {
        v133 := GenS9{F0: v6, F1: 255}
        var v134 uint16 = (^(v2))
        v6 = (v4 ^ (^(v6)))
        _ = v133
        _ = v134
    }
    for fi135 := 0; fi135 < 3; fi135++ {
        if (!(0 >= -90)) {
            switch fi135 {
            case 127:
                for fi136 := 0; fi136 < 15; fi136++ {
                    var v137 int64 = 42
                    if true {
                        if false {
                            _ = len(v3)
                            v2 -= 1
                            v6 -= v4
                        } else {
                            v138 := GenS12{F0: v2, F1: (v137 << ((346) & 63))}
                            _ = v138
                        }
                        v139 := v1[(fi136 ^ (fi136 + fi136))]
                        _ = v139
                    } else {
                        v140 := GenS8{F0: "hello", F1: (^(fi136))}
                        var v141 int16 = 399
                        var v142 int = (fi136 * fi136)
                        _ = v140
                        _ = v141
                        _ = v142
                    }
                    var v143 uint32 = 42
                    _ = fi136
                    _ = v137
                    _ = v143
                }
            case 1:
                v144 := GenS9{F0: 1, F1: (v2 - v2)}
                _ = v144
            case 9223372036854775807:
                v2 += (v2 >> ((4294967295) & 63))
            }
        }
        var v145 float32 = 1.5
        _ = fi135
        _ = v145
    }
    for fi146 := 0; fi146 < 2; fi146++ {
        var v147 uint32 = 4294967295
        fmt.Println(fi146, fi146, v9)
        _ = fi146
        _ = v147
    }
    delete(v3, "hello")
    _ = len(v1)
    for fi148 := 0; fi148 < 10; fi148++ {
        v149 := []int{fi148, fi148, fi148}
        _ = fi148
        _ = v149
    }
    for fi150 := 0; fi150 < 16; fi150++ {
        var v151 int32 = v6
        v2 += 65535
        _ = fi150
        _ = v151
    }
    for fi152 := 0; fi152 < 4; fi152++ {
        if false {
            v9 -= 0
            v153 := GenS3{F0: 1000, F1: 0}
            v6 -= v4
            _ = v153
        } else {
            if (true && false) {
                if false {
                    v154 := []int{}
                    _ = v154
                }
                for fi155 := 0; fi155 < 8; fi155++ {
                    fmt.Println(v4, v4, fi155)
                    fmt.Println(fi155)
                    _ = fi155
                }
                var v156 uint8 = 0
                _ = v156
            } else {
                var v157 uint32 = uint32(v9)
                var v158 float64 = (v5 * (v5 + v5))
                var v159 bool = true
                _ = v157
                _ = v158
                _ = v159
            }
        }
        _ = fi152
    }
    var v160 uint64 = (v9 | (v9 & v9))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v160:", v160)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v160
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
