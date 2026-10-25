package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 int; F2 uint32 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16; F1 int }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint64; F1 int64 }
type GenS11 struct { F0 int64; F1 uint16 }
type GenS12 struct { F0 int; F1 string }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 uint; F1 string }
type GenS15 struct { F0 int32; F1 string; F2 int64 }
type GenS16 struct { F0 int32; F1 uint16 }
type GenS17 struct { F0 int16; F1 uint64 }
type GenS18 struct { F0 int64; F1 uint; F2 uint }
type GenS19 struct { F0 bool; F1 bool; F2 int }
func genFn0() int {
    if true {
        var v6 float64 = 0.1
        v7 := []int{2147483647}
        fmt.Println(v6)
        _ = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(fi8)
            var v9 float32 = 1.5
            for fi10 := 0; fi10 < 19; fi10++ {
                fmt.Println(v9, fi10, fi8)
                switch "foo bar" {
                case "hello":
                    v11 := []int{(fi8 * (^(fi8))), fi8, (fi8 & fi10), (fi10 ^ (fi10 * fi8))}
                    _ = v11
                case "\n\t":
                    v9 -= float32(0.0)
                default:
                    v12 := GenS1{F0: 0}
                    _ = v12
                }
                v13 := GenS1{F0: -2147483648}
                _ = fi10
                _ = v13
            }
            _ = fi8
            _ = v9
        }
    }
    v14 := []int{-9223372036854775808, 128, -2147483648, 1}
    _ = v14
    return -684
}

func main() {
    var v1 bool = true
    var v2 int16 = 735
    var v3 int16 = v2
    var v4 uint32 = 542
    var v5 int16 = 32767
    v1 = v1
    if v1 {
        v5 = (v5 % (((^(v3))) | 1))
        v3 = (v3 / ((v2) | 1))
        v15 := []int{-369, 2147483647, 0, 42}
        _ = v15
    } else {
        v16 := []int{9223372036854775807}
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 int8 = 97
            var v19 uint64 = 1
            v20 := GenS2{F0: (255 <= 32767), F1: fi17, F2: v4}
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 1; fi21++ {
            if ((1.5 == float32(1e10)) && true) {
                for fi22 := 0; fi22 < 8; fi22++ {
                    fmt.Println(v5, v3, fi21)
                    v3 += 42
                    _ = fi22
                }
                v3 = (v3 >> ((87) & 63))
            }
            if true {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                v2 = v2
                v23 := GenS1{F0: 0}
                _ = v23
            }
            v3 = (v3 - v3)
            _ = fi21
        }
        _ = v16
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        for fi25 := 0; fi25 < 18; fi25++ {
            if (!((3.14159 * 1e10) <= (0.0 / 1.5))) {
                if v1 {
                    if false {
                        v26 := GenS2{F0: true, F1: -2147483648, F2: 11}
                        v27 := GenS3{F0: 567, F1: (^(v4))}
                        if false {
                            v3 -= (v3 ^ (v5 >> ((830) & 63)))
                            v2 = 1
                            fmt.Println(fi25)
                        } else {
                            v28 := []int{}
                            switch (-(fi24)) {
                            case 932:
                                fmt.Println(v1, fi24, v3)
                            case 190:
                                switch "hello" {
                                case "":
                                    v1 = ((v2 ^ -1) < v5)
                                case "hello":
                                    var v29 int16 = v3
                                    _ = v29
                                case "a":
                                    var v30 float32 = 1.5
                                    _ = v30
                                default:
                                    v1 = (v5 == -1)
                                }
                            case -1:
                                fmt.Println(v4, fi25)
                            default:
                                v31 := GenS4{F0: 42, F1: 867}
                                _ = v31
                            }
                            fmt.Println(v2, fi24)
                            if (v1 && v1) {
                                v2 = (v2 & (^(v2)))
                                fmt.Println(fi24)
                            }
                            switch (^(fi25)) {
                            case 1:
                                if v1 {
                                    var v32 float64 = -2.25
                                    var v33 int = (fi25 | (fi25 + genFn0()))
                                    _ = v32
                                    _ = v33
                                } else {
                                    v34 := GenS5{F0: ""}
                                    if len(v28) > 0 {
                                        v28[0] = -1
                                    }
                                    switch genFn0() {
                                    case -537:
                                        var v35 uint8 = uint8(v4)
                                        _ = v35
                                    case 4294967295:
                                        v36 := GenS6{F0: 255}
                                        _ = v36
                                    case 9223372036854775807:
                                        v37 := GenS7{F0: 42, F1: (fi24 >> ((429) & 63))}
                                        _ = v37
                                    default:
                                        v4 += (^(v4))
                                    }
                                    _ = v34
                                }
                            }
                            _ = v28
                        }
                        _ = v26
                        _ = v27
                    } else {
                        fmt.Println(v3, fi25)
                        v3 = (v2 % ((v5) | 1))
                        var v38 uint8 = 1
                        _ = v38
                    }
                    if v1 {
                        v39 := []int{}
                        v2 = 1
                        fmt.Println(v4, fi24)
                        _ = v39
                    } else {
                        v2 -= -761
                        v3 += v5
                        var v40 uint8 = 0
                        v1 = ((true || (1.5 <= -2.25)) || v1)
                        _ = v40
                    }
                    v2 -= v2
                }
                if true {
                    var v41 int64 = 1
                    _ = v41
                } else {
                    var v42 uint32 = v4
                    v43 := GenS8{F0: 139}
                    _ = v42
                    _ = v43
                }
                switch fi24 {
                case -491:
                    switch "\n\t" {
                    case "":
                        v1 = (116 != 100)
                    case "foo bar":
                        if ((42 <= 0) && (42 > 4294967295)) {
                            v44 := GenS9{F0: (float64((1e10 * (1.5 - 0.1))) == (1e10 - 3.14159))}
                            var v45 float32 = 3.14159
                            switch "" {
                            case "foo bar":
                                fmt.Println(fi25, v45)
                            }
                            _ = v44
                            _ = v45
                        } else {
                            if v1 {
                                switch "foo bar" {
                                case "hello":
                                    if v1 {
                                        v4 = (v4 & (^(v4)))
                                        var v46 int64 = -9223372036854775808
                                        var v47 bool = ((0.1 - 0.0) < 1e10)
                                        _ = v46
                                        _ = v47
                                    } else {
                                        if (127 <= 42) {
                                            if (!false) {
                                                if (385 < 542) {
                                                    v48 := []int{(-(fi25))}
                                                    v2 = (v3 % ((v3) | 1))
                                                    var v49 int = fi24
                                                    var v50 int16 = (v2 >> ((1) & 63))
                                                    var v51 uint8 = 113
                                                    v48 = append(v48, (^(fi24)))
                                                    var v52 uint = 255
                                                    _ = len(v48)
                                                    switch 0 {
                                                    case 32768:
                                                        var v53 int64 = 0
                                                        _ = v53
                                                    case 0:
                                                        v48 = append(v48, (v49 << ((v52) & 63)))
                                                    }
                                                    _ = v48
                                                    _ = v49
                                                    _ = v50
                                                    _ = v51
                                                    _ = v52
                                                } else {
                                                    var v54 uint = 742
                                                    fmt.Println(fi24)
                                                    var v55 uint32 = v4
                                                    _ = v54
                                                    _ = v55
                                                }
                                                v3 += v2
                                            } else {
                                                if v1 {
                                                    v3 += (v5 - (^(v5)))
                                                    if (!true) {
                                                        v56 := []int{}
                                                        if true {
                                                            v1 = ("\n\t" != "\n\t")
                                                        } else {
                                                            fmt.Println(v4, v3)
                                                        }
                                                        _ = v56
                                                    }
                                                    v57 := GenS5{F0: ((("foo bar" + "hello") + "\n\t") + "hello")}
                                                    var v58 uint64 = 975
                                                    _ = v57
                                                    _ = v58
                                                } else {
                                                    if v1 {
                                                        var v59 int32 = 749
                                                        var v60 uint8 = 42
                                                        switch -9223372036854775808 {
                                                        case 128:
                                                            var v61 string = "a"
                                                            _ = v61
                                                        default:
                                                            if (!((v1 || (0 == -128)) || (!(((true && v1) && (972 <= 42)) && (18446744073709551615 <= uint64(v60)))))) {
                                                                v3 += -192
                                                                v5 -= (v3 >> ((784) & 63))
                                                            } else {
                                                                v62 := []int{fi24, (^(fi25)), 923, genFn0()}
                                                                var v63 uint8 = v60
                                                                var v64 uint8 = v60
                                                                v5 += v2
                                                                var v65 uint16 = 1
                                                                _ = v62
                                                                _ = v63
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                        }
                                                        _ = v59
                                                        _ = v60
                                                    } else {
                                                        if false {
                                                            v66 := GenS9{F0: (v1 || v1)}
                                                            fmt.Println(v4)
                                                            _ = v66
                                                        }
                                                        v67 := GenS10{F0: 42, F1: -843}
                                                        switch (("a" + "\n\t") + ("foo bar" + "\n\t")) {
                                                        case "foo bar":
                                                            if v1 {
                                                                switch fi25 {
                                                                case 9223372036854775807:
                                                                    v68 := []int{fi25, fi25}
                                                                    _ = v68
                                                                case 128:
                                                                    v5 = -96
                                                                case 2147483647:
                                                                    var v69 uint32 = 2
                                                                    _ = v69
                                                                default:
                                                                    v3 = (v3 & 1)
                                                                }
                                                                v70 := GenS11{F0: 42, F1: 844}
                                                                var v71 float64 = 0.1
                                                                _ = v70
                                                                _ = v71
                                                            }
                                                        case "hello":
                                                            if (!(!(v4 <= v4))) {
                                                                switch fi25 {
                                                                case 4294967295:
                                                                    var v72 string = ("" + ("a" + ""))
                                                                    _ = v72
                                                                case -9223372036854775808:
                                                                    var v73 string = (("hello" + "\n\t") + (("hello" + "\n\t") + ("hello" + "a")))
                                                                    _ = v73
                                                                }
                                                                switch fi24 {
                                                                case 2147483647:
                                                                    fmt.Println(v3, fi24, fi24)
                                                                case 84:
                                                                    v74 := []int{(-(fi24)), (fi24 ^ fi25)}
                                                                    _ = v74
                                                                default:
                                                                    v75 := GenS8{F0: 2}
                                                                    _ = v75
                                                                }
                                                                var v76 uint16 = 908
                                                                _ = v76
                                                            }
                                                        case "":
                                                            v1 = (9223372036854775807 >= -158)
                                                        default:
                                                            if v1 {
                                                                v4 -= (v4 % ((767) | 1))
                                                            } else {
                                                                v3 = 42
                                                            }
                                                        }
                                                        _ = v67
                                                    }
                                                    var v77 uint16 = 1
                                                    _ = v77
                                                }
                                                v78 := []int{fi25, 42, genFn0(), -2147483648}
                                                _ = v78
                                            }
                                            if true {
                                                switch ("hello" + "a") {
                                                case "hello":
                                                    v5 = v5
                                                case "a":
                                                    v2 = v3
                                                case "foo bar":
                                                    if v1 {
                                                        fmt.Println(v1)
                                                    } else {
                                                        if true {
                                                            var v79 int64 = 769
                                                            if v1 {
                                                                v80 := GenS12{F0: (fi24 >> ((152) & 63)), F1: ("" + "")}
                                                                var v81 int32 = -1
                                                                v82 := []int{fi24, (fi24 & fi24)}
                                                                _ = v80
                                                                _ = v81
                                                                _ = v82
                                                            } else {
                                                                if false {
                                                                    var v83 int64 = v79
                                                                    fmt.Println(v83, fi24)
                                                                    _ = v83
                                                                } else {
                                                                    if v1 {
                                                                        v79 = (v79 ^ (v79 % ((v79) | 1)))
                                                                        if v1 {
                                                                            if v1 {
                                                                                v4 -= 162
                                                                            } else {
                                                                                v84 := GenS7{F0: 2, F1: (fi25 & 32768)}
                                                                                v1 = true
                                                                                _ = v84
                                                                            }
                                                                            v85 := GenS13{F0: 42}
                                                                            _ = v85
                                                                        }
                                                                        switch ("" + "foo bar") {
                                                                        case "":
                                                                            v86 := []int{}
                                                                            _ = v86
                                                                        }
                                                                        v5 = v2
                                                                        var v87 float64 = 0.1
                                                                        if v1 {
                                                                            if false {
                                                                                var v88 int32 = -151
                                                                                _ = v88
                                                                            } else {
                                                                                var v89 uint = 65535
                                                                                v79 = (-(v79))
                                                                                _ = v89
                                                                            }
                                                                        } else {
                                                                            var v90 uint = 417
                                                                            var v91 float32 = 3.14159
                                                                            _ = v90
                                                                            _ = v91
                                                                        }
                                                                        _ = v87
                                                                    }
                                                                }
                                                                v92 := []int{(fi24 + fi25), genFn0(), fi24}
                                                                v4 = (^(v4))
                                                                _ = v92
                                                            }
                                                            v5 += (-(v2))
                                                            _ = v79
                                                        } else {
                                                            v93 := []int{0, -789, (fi24 ^ genFn0()), genFn0()}
                                                            v5 = v2
                                                            _ = v93
                                                        }
                                                        v4 += (v4 / ((0) | 1))
                                                        fmt.Println(v4, v1, v3)
                                                    }
                                                default:
                                                    var v94 float32 = 0.0
                                                    _ = v94
                                                }
                                                v95 := []int{}
                                                v4 -= (v4 << ((65535) & 63))
                                                _ = v95
                                            }
                                            var v96 uint = 255
                                            _ = v96
                                        }
                                        var v97 int64 = -1
                                        v2 = 0
                                        _ = v97
                                    }
                                case "a":
                                    switch (fi24 ^ fi25) {
                                    case -879:
                                        switch fi24 {
                                        case 699:
                                            var v98 bool = (127 == 2)
                                            _ = v98
                                        default:
                                            var v99 int8 = 41
                                            _ = v99
                                        }
                                    }
                                case "\n\t":
                                    if (v1 || (!v1)) {
                                        var v100 string = ("foo bar" + "")
                                        _ = v100
                                    } else {
                                        if true {
                                            var v101 string = "foo bar"
                                            switch 127 {
                                            case 32768:
                                                var v102 int = 4294967295
                                                _ = v102
                                            case 1:
                                                var v103 uint32 = (^(v4))
                                                var v104 bool = (!(2 != 65535))
                                                var v105 float32 = (-2.25 + 0.1)
                                                switch "foo bar" {
                                                case "\n\t":
                                                    var v106 bool = true
                                                    _ = v106
                                                case "foo bar":
                                                    if false {
                                                        v107 := []int{128, (-(fi25)), (fi24 ^ (fi24 >> ((255) & 63)))}
                                                        v108 := GenS14{F0: 214, F1: (v101 + "\n\t")}
                                                        _ = v107
                                                        _ = v108
                                                    } else {
                                                        v109 := []int{}
                                                        v101 = v101
                                                        _ = v109
                                                    }
                                                case "":
                                                    v110 := GenS8{F0: 326}
                                                    _ = v110
                                                default:
                                                    v104 = false
                                                }
                                                _ = v103
                                                _ = v104
                                                _ = v105
                                            }
                                            var v111 float32 = -2.25
                                            _ = v101
                                            _ = v111
                                        } else {
                                            v112 := []int{fi25, fi25, 1}
                                            var v113 bool = (true && (!(0.0 == 3.14159)))
                                            _ = v112
                                            _ = v113
                                        }
                                        v3 -= v5
                                    }
                                }
                                switch fi25 {
                                case 1:
                                    var v114 uint16 = 0
                                    _ = v114
                                case -137:
                                    if v1 {
                                        fmt.Println(fi24, v4, v1)
                                    }
                                default:
                                    v3 -= (v3 * (v5 >> ((65535) & 63)))
                                    var v115 uint8 = 127
                                    _ = v115
                                }
                            } else {
                                v3 += (v5 * (^(v5)))
                                if (v1 && (-598 < -691)) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v3 = (^(v5))
                                    case "hello":
                                        fmt.Println(v2, fi25, v2)
                                    case "":
                                        v116 := []int{fi25, (fi24 & (fi25 * fi24)), 0, (fi24 ^ (^(fi24)))}
                                        _ = v116
                                    default:
                                        var v117 float32 = (-2.25 - 3.14159)
                                        _ = v117
                                    }
                                    fmt.Println(v3)
                                    var v118 int8 = -52
                                    v119 := GenS14{F0: 255, F1: ""}
                                    var v120 int16 = -1
                                    _ = v118
                                    _ = v119
                                    _ = v120
                                }
                            }
                            if (-1 > 1) {
                                fmt.Println(fi24, fi25)
                            }
                        }
                    default:
                        var v121 uint8 = 173
                        _ = v121
                    }
                    v122 := GenS10{F0: 0, F1: -26}
                    fmt.Println(fi24, v1, v5)
                    _ = v122
                case -2147483648:
                    var v123 int = fi25
                    _ = v123
                case 37:
                    if false {
                        v2 = v3
                        fmt.Println(v3, v1)
                    } else {
                        v124 := GenS10{F0: 0, F1: -9223372036854775808}
                        v4 -= v4
                        _ = v124
                    }
                }
                var v125 uint8 = 127
                _ = v125
            } else {
                fmt.Println(v5, fi24, v3)
                fmt.Println(v2)
            }
            switch 2147483647 {
            case 1:
                if (!((fi25 & genFn0()) >= (fi25 & 9223372036854775807))) {
                    var v126 string = ("\n\t" + ("foo bar" + "a"))
                    v2 -= v5
                    _ = v126
                }
            case 937:
                var v127 int8 = 127
                _ = v127
            }
            if (255 < 2) {
                v128 := []int{genFn0()}
                _ = v128
            } else {
                v4 += v4
                if v1 {
                    var v129 int8 = 27
                    switch "hello" {
                    case "":
                        var v130 uint64 = 742
                        v129 -= (v129 ^ (v129 << ((1) & 63)))
                        if (true || v1) {
                            var v131 int16 = (v3 % ((1000) | 1))
                            _ = v131
                        } else {
                            var v132 uint32 = 42
                            v133 := []int{(fi24 * 42)}
                            v134 := GenS15{F0: int32(v132), F1: "foo bar", F2: 42}
                            _ = v132
                            _ = v133
                            _ = v134
                        }
                        v135 := GenS16{F0: -420, F1: 2}
                        v136 := GenS14{F0: 2, F1: ((("" + ("a" + "")) + ("" + "foo bar")) + "hello")}
                        _ = v130
                        _ = v135
                        _ = v136
                    }
                    var v137 uint8 = 42
                    v129 = v129
                    switch "a" {
                    case "hello":
                        v138 := GenS2{F0: ("hello" <= "hello"), F1: (fi25 - fi25), F2: v4}
                        _ = v138
                    case "a":
                        v2 -= v5
                    default:
                        if (!("a" != "a")) {
                            v139 := GenS6{F0: (v137 / (((v137 - 1)) | 1))}
                            _ = v139
                        }
                    }
                    _ = v129
                    _ = v137
                }
                v5 -= 0
            }
            _ = fi25
        }
        for fi140 := 0; fi140 < 8; fi140++ {
            v141 := GenS13{F0: 0}
            if ((v4 >= 924) || v1) {
                if true {
                    var v142 uint32 = v4
                    var v143 string = "hello"
                    _ = v142
                    _ = v143
                } else {
                    v144 := GenS17{F0: (v3 >> ((726) & 63)), F1: 4294967295}
                    _ = v144
                }
                if (!true) {
                    var v145 int32 = 2147483647
                    v146 := []int{genFn0()}
                    if (v2 > 43) {
                        v145 = -198
                        var v147 int16 = (v2 * (v3 % (((v3 | 32767)) | 1)))
                        fmt.Println(v145, fi24, v3)
                        _ = v147
                    } else {
                        var v148 int32 = v145
                        if true {
                            v2 = v2
                            v2 += v5
                        }
                        _ = v148
                    }
                    _ = v145
                    _ = v146
                }
                switch ("hello" + "foo bar") {
                case "\n\t":
                    v149 := []int{}
                    _ = v149
                }
            } else {
                fmt.Println(v3, v3)
            }
            v5 -= (-(v3))
            _ = fi140
            _ = v141
        }
        _ = fi24
    }
    for fi150 := 0; fi150 < 14; fi150++ {
        for fi151 := 0; fi151 < 3; fi151++ {
            v5 = (v5 | (v2 + (v3 >> ((0) & 63))))
            v152 := []int{}
            _ = len(v152)
            _ = fi151
            _ = v152
        }
        v153 := GenS18{F0: 42, F1: 18446744073709551615, F2: uint(v4)}
        _ = fi150
        _ = v153
    }
    v154 := GenS9{F0: (((((25 > 2) && false) && (0 >= 42)) && (v1 || (1 == (-(v2))))) && (!false))}
    v155 := GenS17{F0: (v2 % (((v2 - (^(v2)))) | 1)), F1: 450}
    v156 := []int{0, 2147483647}
    for fi157 := 0; fi157 < 13; fi157++ {
        if (!v1) {
            if v1 {
                if v1 {
                    if len(v156) > 0 {
                        v156[0] = genFn0()
                    }
                    fmt.Println(v3, v3)
                }
                v158 := GenS19{F0: v1, F1: false, F2: 2147483647}
                _ = v158
            } else {
                fmt.Println(fi157, v1, v1)
                v159 := GenS5{F0: "foo bar"}
                _ = v159
            }
            for fi160 := 0; fi160 < 5; fi160++ {
                var v161 float32 = float32(1.5)
                _ = fi160
                _ = v161
            }
        }
        if len(v156) > 0 {
            v156[0] = fi157
        }
        _ = fi157
    }
    _ = len(v156)
    v2 -= 1
    var v162 uint32 = v4
    switch 415 {
    case 42:
        var v163 uint32 = v162
        _ = v163
    case 2147483647:
        var v164 uint64 = 18446744073709551615
        _ = v164
    default:
        _ = len(v156)
    }
    v4 -= (v162 * v162)
    switch ((("a" + "") + "") + "foo bar") {
    case "hello":
        var v165 uint16 = 32767
        _ = v165
    case "":
        var v166 uint8 = 127
        _ = v166
    default:
        var v167 uint32 = (v162 / ((4294967295) | 1))
        _ = v167
    }
    if (!(v2 > 42)) {
        var v168 int32 = 146
        if false {
            var v169 uint8 = 1
            _ = v169
        }
        _ = v168
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v156:", len(v156), v156)
    fmt.Println("v162:", v162)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v154
    _ = v155
    _ = v156
    _ = v162
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
