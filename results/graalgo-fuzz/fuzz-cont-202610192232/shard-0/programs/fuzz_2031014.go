package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 uint16; F1 int32; F2 uint }
type GenS8 struct { F0 int64; F1 uint16; F2 int }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 uint16; F1 uint16; F2 string }
type GenS11 struct { F0 int16; F1 string }
type GenS12 struct { F0 uint16; F1 uint; F2 bool }
type GenS13 struct { F0 int; F1 uint8 }
type GenS14 struct { F0 int16; F1 uint16; F2 bool }
func genFn0(p0 string) uint16 {
    p0 = "a"
    fmt.Println(p0)
    switch ((p0 + ((("a" + ("hello" + "foo bar")) + p0) + (("" + p0) + "\n\t"))) + "hello") {
    case "foo bar":
        if false {
            switch (p0 + "foo bar") {
            case "hello":
                p0 = "a"
            case "":
                p0 = ""
            }
        }
    case "a":
        var v4 uint16 = 2
        _ = v4
    default:
        var v5 uint32 = 894
        _ = v5
    }
    _ = p0
    return 810
}
func genFn1(p0 int64) int8 {
    if (!true) {
        if (42 <= 134) {
            var v6 int32 = -1
            _ = v6
        } else {
            var v7 int16 = 411
            v8 := GenS1{F0: (p0 / ((p0) | 1))}
            _ = v7
            _ = v8
        }
    } else {
        if true {
            var v9 uint32 = 291
            if true {
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 int64 = p0
                    if (32767 < -635) {
                        var v12 uint64 = 2
                        _ = v12
                    }
                    if ((!true) && true) {
                        for fi13 := 0; fi13 < 3; fi13++ {
                            fmt.Println(p0, v11, p0)
                            _ = fi13
                        }
                        for fi14 := 0; fi14 < 18; fi14++ {
                            v9 = 1
                            _ = fi14
                        }
                    }
                    _ = fi10
                    _ = v11
                }
            }
            switch 12 {
            case 888:
                v15 := []int{1, 998, 173, -2147483648}
                _ = v15
            }
            _ = v9
        } else {
            p0 = p0
            switch 127 {
            case 106:
                var v16 uint = 42
                _ = v16
            default:
                var v17 uint16 = 1
                _ = v17
            }
        }
        v18 := []int{38, 99}
        _ = v18
    }
    fmt.Println(p0)
    if (255 == 2) {
        v19 := GenS2{F0: 255, F1: (1e10 <= 3.14159)}
        switch "hello" {
        case "foo bar":
            for fi20 := 0; fi20 < 9; fi20++ {
                p0 = p0
                fmt.Println(fi20)
                if true {
                    var v21 uint64 = 2
                    var v22 uint8 = 248
                    _ = v21
                    _ = v22
                } else {
                    if true {
                        p0 += (p0 - 9223372036854775807)
                        p0 -= p0
                        fmt.Println(p0, fi20)
                    }
                    for fi23 := 0; fi23 < 15; fi23++ {
                        if true {
                            v24 := GenS2{F0: 0, F1: (!(0 > 255))}
                            if ((false || (1.5 > float64(-2.25))) && true) {
                                if true {
                                    switch (("foo bar" + "\n\t") + "hello") {
                                    case "":
                                        if true {
                                            fmt.Println(fi20, p0)
                                        }
                                    case "a":
                                        switch "hello" {
                                        case "":
                                            if (("hello" + ("\n\t" + ("" + ("hello" + "\n\t")))) != "foo bar") {
                                                var v25 uint32 = 4294967295
                                                if ("hello" == ("foo bar" + "foo bar")) {
                                                    var v26 bool = true
                                                    _ = v26
                                                } else {
                                                    var v27 string = "\n\t"
                                                    var v28 bool = true
                                                    var v29 uint16 = genFn0(((("\n\t" + v27) + (v27 + ((("\n\t" + v27) + v27) + "foo bar"))) + v27))
                                                    _ = v27
                                                    _ = v28
                                                    _ = v29
                                                }
                                                v25 -= (v25 | v25)
                                                _ = v25
                                            } else {
                                                var v30 uint8 = 0
                                                switch "a" {
                                                case "foo bar":
                                                    v31 := GenS3{F0: "a"}
                                                    fmt.Println(fi20)
                                                    _ = v31
                                                case "a":
                                                    if false {
                                                        fmt.Println(fi23, p0, fi23)
                                                    }
                                                    var v32 uint = 42
                                                    _ = v32
                                                case "":
                                                    v33 := []int{(fi20 << ((2) & 63)), 1, (-(fi20))}
                                                    _ = v33
                                                default:
                                                    v34 := GenS3{F0: "hello"}
                                                    _ = v34
                                                }
                                                if false {
                                                    if ((3.14159 / -2.25) != 1e10) {
                                                        v35 := []int{}
                                                        v35 = append(v35, (fi20 | 2147483647))
                                                        p0 = p0
                                                        _ = v35
                                                    }
                                                    var v36 int32 = 42
                                                    var v37 bool = (!(true || false))
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    v30 = v30
                                                }
                                                _ = v30
                                            }
                                        case "foo bar":
                                            v38 := []int{1, fi23, fi20, fi20}
                                            _ = v38
                                        default:
                                            fmt.Println(p0, p0, p0)
                                        }
                                    }
                                }
                                var v39 uint64 = 833
                                _ = v39
                            } else {
                                v40 := []int{}
                                _ = v40
                            }
                            var v41 uint8 = 127
                            _ = v24
                            _ = v41
                        } else {
                            p0 = (p0 >> ((42) & 63))
                            v42 := GenS4{F0: (("\n\t" + "") + "\n\t"), F1: 42}
                            _ = v42
                        }
                        var v43 int64 = p0
                        var v44 string = "\n\t"
                        _ = fi23
                        _ = v43
                        _ = v44
                    }
                }
                _ = fi20
            }
        case "hello":
            v45 := GenS5{F0: "a", F1: false}
            _ = v45
        }
        v46 := []int{4294967295, 803, -9223372036854775808}
        _ = v19
        _ = v46
    }
    _ = p0
    return 104
}

func main() {
    var v1 int8 = 68
    var v2 uint32 = 116
    v3 := make(map[int]bool)
    var v47 bool = (990 >= 42)
    switch 9223372036854775807 {
    case 128:
        delete(v3, 32768)
    case 227:
        if (v47 || (18446744073709551615 <= 431)) {
            v48 := GenS6{F0: 13, F1: 0}
            for fi49 := 0; fi49 < 12; fi49++ {
                switch fi49 {
                case 2147483647:
                    if v47 {
                        v50 := GenS7{F0: 32767, F1: 2147483647, F2: 42}
                        var v51 int32 = 0
                        _ = v50
                        _ = v51
                    } else {
                        fmt.Println(v47)
                        var v52 uint = 65535
                        if v47 {
                            v2 += v2
                        }
                        _ = v52
                    }
                case 127:
                    switch -463 {
                    case -99:
                        if v47 {
                            for fi53 := 0; fi53 < 18; fi53++ {
                                if v47 {
                                    var v54 uint16 = 78
                                    var v55 string = "hello"
                                    if v47 {
                                        if false {
                                            if v47 {
                                                v47 = ((v1 / ((v1) | 1)) != (-(v1)))
                                            } else {
                                                var v56 string = v55
                                                _ = v56
                                            }
                                        } else {
                                            fmt.Println(v54, fi49)
                                            v47 = false
                                            _ = len(v3)
                                        }
                                    }
                                    _ = v54
                                    _ = v55
                                }
                                v57 := v3[0]
                                v2 = 0
                                v47 = (v2 <= (^(v2)))
                                _ = fi53
                                _ = v57
                            }
                        } else {
                            if (false && (false && v47)) {
                                if false {
                                    if (("foo bar" + "a") == "") {
                                        fmt.Println(v1, fi49)
                                        switch "\n\t" {
                                        case "hello":
                                            if (9223372036854775807 == 42) {
                                                for fi58 := 0; fi58 < 13; fi58++ {
                                                    if true {
                                                        v59 := []int{(fi58 / ((fi49) | 1)), -2147483648}
                                                        var v60 float32 = 0.0
                                                        v61 := GenS8{F0: 0, F1: 535, F2: (^(fi58))}
                                                        _ = v59
                                                        _ = v60
                                                        _ = v61
                                                    } else {
                                                        if true {
                                                            var v62 int = (fi58 ^ fi49)
                                                            v1 = 100
                                                            var v63 float32 = (-2.25 * 0.1)
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                    }
                                                    v64 := GenS9{F0: 255, F1: 0}
                                                    _ = fi58
                                                    _ = v64
                                                }
                                                fmt.Println(v1)
                                                switch (fi49 | (fi49 & fi49)) {
                                                case 767:
                                                    v47 = (18 > 1)
                                                case -2147483648:
                                                    if (v47 && ("\n\t" < (("hello" + "\n\t") + "hello"))) {
                                                        v65 := []int{-1}
                                                        for fi66 := 0; fi66 < 9; fi66++ {
                                                            _ = len(v3)
                                                            if false {
                                                                switch (fi66 << ((0) & 63)) {
                                                                case 4294967295:
                                                                    v1 += v1
                                                                case 128:
                                                                    fmt.Println(fi66, fi66, v1)
                                                                case 32767:
                                                                    v47 = (0.1 < -2.25)
                                                                }
                                                            } else {
                                                                fmt.Println(fi66, fi66)
                                                                if v47 {
                                                                    if false {
                                                                        var v67 uint = 244
                                                                        v2 -= (v2 | 2147483647)
                                                                        v1 += genFn1(9223372036854775807)
                                                                        _ = v67
                                                                    } else {
                                                                        if (v47 || v47) {
                                                                            var v68 uint32 = v2
                                                                            _ = v68
                                                                        } else {
                                                                            var v69 float64 = 0.1
                                                                            v2 = (v2 ^ v2)
                                                                            var v70 uint16 = genFn0("")
                                                                            _ = v69
                                                                            _ = v70
                                                                        }
                                                                        v47 = (true || v47)
                                                                        var v71 uint8 = 128
                                                                        _ = v71
                                                                    }
                                                                }
                                                                var v72 float64 = 3.14159
                                                                _ = v72
                                                            }
                                                            _ = fi66
                                                        }
                                                        _ = v65
                                                    } else {
                                                        v47 = (!(682 <= -9223372036854775808))
                                                    }
                                                case 42:
                                                    fmt.Println(v47, v2)
                                                default:
                                                    var v73 int = (fi49 * 616)
                                                    _ = v73
                                                }
                                            }
                                        }
                                    } else {
                                        v47 = (3.14159 <= 1.5)
                                    }
                                    v74 := GenS10{F0: 2, F1: 65535, F2: ""}
                                    v75 := GenS9{F0: 255, F1: -561}
                                    _ = v74
                                    _ = v75
                                }
                                switch "\n\t" {
                                case "":
                                    v76 := GenS4{F0: "", F1: 175}
                                    _ = v76
                                case "foo bar":
                                    var v77 int = -2147483648
                                    _ = v77
                                case "a":
                                    switch fi49 {
                                    case 32768:
                                        for fi78 := 0; fi78 < 9; fi78++ {
                                            fmt.Println(v1, v2)
                                            _ = fi78
                                        }
                                    }
                                default:
                                    fmt.Println(v1, fi49, v47)
                                }
                            }
                            v2 -= v2
                        }
                    default:
                        v79 := GenS11{F0: -1, F1: "\n\t"}
                        _ = v79
                    }
                case 1:
                    var v80 bool = v47
                    _ = v80
                }
                v2 = (v2 + 151)
                _ = fi49
            }
            for fi81 := 0; fi81 < 3; fi81++ {
                for fi82 := 0; fi82 < 18; fi82++ {
                    v83 := []int{(fi82 % ((fi82) | 1)), 1, -698}
                    v47 = v47
                    v83 = append(v83, fi82)
                    _ = fi82
                    _ = v83
                }
                for fi84 := 0; fi84 < 2; fi84++ {
                    v85 := GenS4{F0: ("\n\t" + ("hello" + "")), F1: 0}
                    v86 := GenS6{F0: 1, F1: fi81}
                    _ = fi84
                    _ = v85
                    _ = v86
                }
                for fi87 := 0; fi87 < 7; fi87++ {
                    var v88 uint16 = 65535
                    var v89 int8 = 42
                    _ = fi87
                    _ = v88
                    _ = v89
                }
                _ = fi81
            }
            _ = v48
        } else {
            v90 := GenS4{F0: "hello", F1: uint8(v2)}
            if v47 {
                if v47 {
                    v47 = (42 > 2)
                    if (388 < 218) {
                        v2 += (v2 << ((858) & 63))
                    }
                }
                switch 149 {
                case 9223372036854775807:
                    var v91 int8 = (v1 / ((genFn1(393)) | 1))
                    _ = v91
                case -228:
                    for fi92 := 0; fi92 < 8; fi92++ {
                        if v47 {
                            v2 += 42
                        }
                        v1 = v1
                        _ = fi92
                    }
                default:
                    v47 = (0 >= 128)
                }
            }
            v47 = false
            _ = v90
        }
    case -835:
        var v93 float64 = 1e10
        _ = v93
    }
    v2 -= 468
    var v94 uint32 = (v2 - (v2 << ((0) & 63)))
    v2 = v94
    var v95 int = 127
    for fi96 := 0; fi96 < 13; fi96++ {
        if ((!v47) || v47) {
            var v97 bool = v47
            _ = v97
        } else {
            var v98 int8 = (v1 + (v1 ^ (^(v1))))
            if (!v47) {
                if v47 {
                    var v99 int8 = (v98 << ((18446744073709551615) & 63))
                    for fi100 := 0; fi100 < 5; fi100++ {
                        var v101 int16 = 1000
                        var v102 bool = (((("a" + "hello") + ("" + (((("a" + "hello") + "a") + ("hello" + "foo bar")) + ("foo bar" + "a")))) + ("\n\t" + "\n\t")) == ("" + ("a" + "a")))
                        v98 -= 7
                        v103 := GenS12{F0: 411, F1: 426, F2: (2 == 334)}
                        _ = fi100
                        _ = v101
                        _ = v102
                        _ = v103
                    }
                    switch (fi96 + fi96) {
                    case 0:
                        if (((0.1 + (0.0 * 1.5)) - 0.0) <= ((3.14159 / (((float32(float64(0.0)) / 3.14159) + float32(0.1)) - 1.5)) + 1.5)) {
                            v99 = v1
                        }
                    case 9223372036854775807:
                        v104 := []int{(fi96 | v95), (v95 % (((v95 << ((1) & 63))) | 1))}
                        _ = v104
                    case 127:
                        v105 := []int{fi96, 1}
                        _ = v105
                    default:
                        v3[(v95 % (((v95 | fi96)) | 1))] = true
                    }
                    _ = v99
                }
                var v106 int32 = -2147483648
                _ = v106
            } else {
                v107 := GenS1{F0: -1}
                var v108 int32 = 2147483647
                for fi109 := 0; fi109 < 1; fi109++ {
                    var v110 uint32 = v94
                    if (2 == 127) {
                        var v111 string = "foo bar"
                        _ = v111
                    } else {
                        v112 := GenS8{F0: int64(v110), F1: 328, F2: fi109}
                        var v113 uint32 = (v110 + 42)
                        var v114 int8 = int8(v108)
                        _ = v112
                        _ = v113
                        _ = v114
                    }
                    _ = fi109
                    _ = v110
                }
                _ = v107
                _ = v108
            }
            var v115 string = "foo bar"
            _ = v115
            _ = v98
        }
        var v116 uint16 = genFn0("a")
        _ = fi96
        _ = v116
    }
    var v117 int64 = 9223372036854775807
    var v118 uint8 = 128
    if v47 {
        for fi119 := 0; fi119 < 12; fi119++ {
            switch 134 {
            case 32767:
                if (!("a" < "")) {
                    var v120 uint8 = 99
                    v121 := []int{(fi119 >> ((18446744073709551615) & 63)), fi119, fi119, fi119}
                    var v122 uint16 = 32767
                    _ = v120
                    _ = v121
                    _ = v122
                } else {
                    switch (fi119 - (^(v95))) {
                    case 2147483647:
                        v94 = v2
                    case 127:
                        for fi123 := 0; fi123 < 10; fi123++ {
                            v124 := []int{57, fi123, v95, fi123}
                            var v125 uint16 = 828
                            _ = fi123
                            _ = v124
                            _ = v125
                        }
                    case 333:
                        if (!v47) {
                            v126 := v3[(v95 ^ (fi119 / ((fi119) | 1)))]
                            var v127 uint16 = 42
                            _ = v126
                            _ = v127
                        }
                    default:
                        var v128 float32 = 1e10
                        _ = v128
                    }
                    var v129 string = ("\n\t" + "a")
                    if false {
                        v130 := []int{(fi119 << ((42) & 63)), (-(v95)), (v95 * 701)}
                        var v131 float32 = float32(((0.1 * 1.5) - 0.0))
                        _ = v130
                        _ = v131
                    }
                    _ = v129
                }
            case 144:
                v132 := GenS13{F0: v95, F1: v118}
                _ = v132
            case 1:
                if true {
                    v133 := GenS9{F0: 255, F1: (^(v117))}
                    var v134 float32 = -2.25
                    v135 := GenS14{F0: -668, F1: uint16(v118), F2: (210 > 135)}
                    _ = v133
                    _ = v134
                    _ = v135
                } else {
                    var v136 int32 = -766
                    v94 = 1
                    v117 += v117
                    _ = v136
                }
            default:
                if (-2.25 >= 0.0) {
                    var v137 int = fi119
                    var v138 float64 = 1.5
                    var v139 float32 = 1e10
                    _ = v137
                    _ = v138
                    _ = v139
                } else {
                    var v140 int = int(v94)
                    v141 := GenS9{F0: 741, F1: (v117 << ((307) & 63))}
                    var v142 int32 = 0
                    _ = v140
                    _ = v141
                    _ = v142
                }
            }
            v2 -= 541
            fmt.Println(v1, v47)
            _ = fi119
        }
    } else {
        v143 := GenS2{F0: 669, F1: v47}
        var v144 float64 = 1.5
        v144 = v144
        _ = v143
        _ = v144
    }
    var v145 int8 = -128
    var v146 int64 = v117
    var v147 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v118:", v118)
    fmt.Println("v145:", v145)
    fmt.Println("v146:", v146)
    fmt.Println("v147:", v147)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v47:", v47)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v117
    _ = v118
    _ = v145
    _ = v146
    _ = v147
    _ = v2
    _ = v3
    _ = v47
    _ = v94
    _ = v95
}
