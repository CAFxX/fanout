package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8; F1 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int; F1 int }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int64; F1 int16; F2 int8 }
type GenS12 struct { F0 uint16; F1 uint64 }
type GenS13 struct { F0 int8; F1 int16; F2 uint8 }
type GenS14 struct { F0 uint32; F1 int }
func genFn0(p0 uint8, p1 uint32, p2 uint32) string {
    if (!((p1 ^ p2) == p1)) {
        fmt.Println(p1)
    }
    var v7 uint = 997
    var v8 uint = (v7 / ((v7) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 1
    var v3 uint16 = 255
    v4 := make(map[string]int)
    var v5 int32 = -2147483648
    var v6 int16 = 32767
    var v9 float64 = float64(1e10)
    v2 = (^(v2))
    v10 := GenS1{F0: 18446744073709551615, F1: 128}
    v11 := v1[9223372036854775807]
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 16; fi13++ {
            v6 = 1000
            if ((v6 % ((v6) | 1)) <= (v6 << ((2) & 63))) {
                fmt.Println(v5, v3, v2)
                v5 = (v5 | v5)
                var v14 uint64 = 2
                _ = v14
            } else {
                switch "foo bar" {
                case "hello":
                    if false {
                        if v11 {
                            v15 := GenS2{F0: v3, F1: -2147483648}
                            _ = v15
                        } else {
                            v11 = v11
                        }
                        if true {
                            var v16 float32 = (1e10 * (3.14159 * ((1.5 - 1e10) - 3.14159)))
                            if v11 {
                                var v17 int = fi12
                                v5 = (^(v5))
                                var v18 uint32 = uint32(v2)
                                _ = v17
                                _ = v18
                            } else {
                                var v19 uint8 = v2
                                v11 = true
                                if ((1000 > 1) && (v11 && v11)) {
                                    if v11 {
                                        switch (fi12 * 2147483647) {
                                        case -836:
                                            if false {
                                                v20 := GenS2{F0: v3, F1: (v5 % ((int32(v19)) | 1))}
                                                var v21 uint8 = v2
                                                var v22 uint64 = 2
                                                _ = v20
                                                _ = v21
                                                _ = v22
                                            }
                                        case -83:
                                            var v23 bool = true
                                            _ = v23
                                        default:
                                            switch (fi13 << ((743) & 63)) {
                                            case -931:
                                                var v24 int8 = 42
                                                v24 = v24
                                                _ = v24
                                            case 2147483647:
                                                v9 += -2.25
                                            }
                                        }
                                    } else {
                                        var v25 float64 = v9
                                        _ = v25
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        v16 -= 0.1
                                        switch "foo bar" {
                                        case "a":
                                            switch fi13 {
                                            case 741:
                                                switch "a" {
                                                case "foo bar":
                                                    if (v11 || false) {
                                                        switch "" {
                                                        case "\n\t":
                                                            switch (fi13 % (((fi12 / (((fi13 & (fi13 << ((65535) & 63)))) | 1))) | 1)) {
                                                            case 127:
                                                                v26 := []int{fi13}
                                                                _ = v26
                                                            case 850:
                                                                _ = len(v4)
                                                            case 128:
                                                                var v27 string = ((("\n\t" + "hello") + (("a" + ("a" + "a")) + "foo bar")) + (("\n\t" + ("\n\t" + "\n\t")) + "\n\t"))
                                                                _ = v27
                                                            }
                                                        case "a":
                                                            v28 := GenS3{F0: (v5 | v5)}
                                                            _ = v28
                                                        case "foo bar":
                                                            var v29 bool = (!true)
                                                            _ = v29
                                                        default:
                                                            if v11 {
                                                                v30 := GenS4{F0: (v11 || (-1 < -596))}
                                                                v5 = v5
                                                                var v31 uint64 = 794
                                                                _ = v30
                                                                _ = v31
                                                            }
                                                        }
                                                        var v32 uint8 = (v19 | 1)
                                                        _ = v32
                                                    } else {
                                                        var v33 string = ("hello" + ((("a" + ("hello" + ((("hello" + ("\n\t" + ("a" + ("a" + "\n\t")))) + ("\n\t" + "\n\t")) + "foo bar"))) + "a") + "foo bar"))
                                                        _ = v33
                                                    }
                                                case "hello":
                                                    fmt.Println(v9, v11)
                                                case "\n\t":
                                                    v6 = -1
                                                default:
                                                    if (!v11) {
                                                        if (v9 != 1e10) {
                                                            fmt.Println(v11)
                                                        }
                                                        if (!(!(545 != 42))) {
                                                            if v11 {
                                                                v34 := GenS3{F0: (v5 * v5)}
                                                                var v35 float64 = 0.0
                                                                fmt.Println(v16, v2, v11)
                                                                _ = v34
                                                                _ = v35
                                                            }
                                                        }
                                                    } else {
                                                        fmt.Println(v11, v19, v11)
                                                        var v36 uint16 = 0
                                                        _ = v36
                                                    }
                                                }
                                            }
                                        case "hello":
                                            var v37 bool = v11
                                            _ = v37
                                        case "\n\t":
                                            var v38 float64 = v9
                                            _ = v38
                                        }
                                    case "":
                                        v6 = (v6 / (((^(v6))) | 1))
                                    case "a":
                                        switch "\n\t" {
                                        case "foo bar":
                                            fmt.Println(v5, v3, v6)
                                        case "hello":
                                            var v39 uint16 = 351
                                            _ = v39
                                        case "a":
                                            v40 := []int{}
                                            _ = v40
                                        default:
                                            v41 := GenS3{F0: v5}
                                            _ = v41
                                        }
                                    }
                                } else {
                                    var v42 string = "a"
                                    _ = v42
                                }
                                _ = v19
                            }
                            switch fi13 {
                            case -9223372036854775808:
                                var v43 int16 = (v6 << ((65535) & 63))
                                _ = v43
                            case 128:
                                var v44 bool = true
                                _ = v44
                            }
                            _ = v16
                        }
                        v2 += 41
                        v6 = v6
                    } else {
                        if ((v2 ^ v2) < v2) {
                            if v11 {
                                var v45 string = ("" + "foo bar")
                                delete(v1, fi13)
                                _ = v45
                            }
                            var v46 float32 = float32(float64((float32((v9 + v9)) * float32((v9 / 3.14159)))))
                            _ = v46
                        }
                        if (-1 == (v6 | (-(v6)))) {
                            v6 = (v6 / (((^(v6))) | 1))
                            fmt.Println(v5, v11)
                        }
                        if (v11 || true) {
                            fmt.Println(v2, v9)
                        } else {
                            v47 := []int{(fi13 / (((fi13 & fi12)) | 1)), (fi13 >> ((uint(v3)) & 63))}
                            switch ("a" + (("foo bar" + "") + "a")) {
                            case "a":
                                v48 := GenS5{F0: fi12, F1: -10}
                                _ = v48
                            case "foo bar":
                                fmt.Println(v11, v9, v3)
                            }
                            switch ("\n\t" + "hello") {
                            case "\n\t":
                                var v49 int = 4294967295
                                _ = v49
                            case "a":
                                v50 := GenS6{F0: (^(v2))}
                                _ = v50
                            }
                            _ = v47
                        }
                    }
                case "":
                    if ((v6 - int16(v2)) <= 42) {
                        var v51 uint = 0
                        if (100 != -99) {
                            fmt.Println(v5, v6, v2)
                        }
                        switch (fi13 - 32767) {
                        case -1:
                            v52 := GenS7{F0: v2, F1: (v6 & 32767)}
                            _ = v52
                        default:
                            v3 = (v3 ^ 255)
                        }
                        _ = v51
                    }
                default:
                    if (162 >= 452) {
                        var v53 int = -9223372036854775808
                        _ = v53
                    } else {
                        if v11 {
                            var v54 int16 = 609
                            v55 := []int{}
                            if v11 {
                                v5 -= (v5 - v5)
                            }
                            if (-1 > -9223372036854775808) {
                                v2 = 245
                                var v56 uint32 = 645
                                _ = v56
                            } else {
                                var v57 uint32 = 1
                                _ = v57
                            }
                            switch "a" {
                            case "hello":
                                var v58 int8 = 100
                                _ = v58
                            default:
                                v59 := GenS7{F0: 194, F1: (v6 + (v54 >> ((255) & 63)))}
                                _ = v59
                            }
                            v5 -= int32(v2)
                            _ = v54
                            _ = v55
                        } else {
                            var v60 uint8 = v2
                            var v61 int8 = -63
                            if (2147483647 < v5) {
                                v62 := GenS8{F0: (^(v60))}
                                _ = v62
                            }
                            _ = v60
                            _ = v61
                        }
                    }
                }
            }
            if false {
                var v63 float64 = (-2.25 * 1.5)
                v64 := v1[-697]
                v65 := GenS1{F0: 4294967295, F1: -337}
                _ = v63
                _ = v64
                _ = v65
            }
            _ = fi13
        }
        switch -1 {
        case -89:
            var v66 string = ("hello" + "a")
            _ = v66
        case 1:
            var v67 int8 = -1
            _ = v67
        default:
            v3 = v3
        }
        _ = fi12
    }
    var v68 int8 = int8(v3)
    var v69 uint32 = 65535
    v68 = (v68 | v68)
    if true {
        var v70 int8 = 100
        if v11 {
            fmt.Println(v68, v6, v68)
        }
        _ = v70
    } else {
        v9 = v9
        var v71 uint = 43
        if v11 {
            switch -2147483648 {
            case 1:
                for fi72 := 0; fi72 < 18; fi72++ {
                    var v73 int64 = -1
                    _ = fi72
                    _ = v73
                }
            case -192:
                switch -158 {
                case 32768:
                    if ("\n\t" < "a") {
                        v74 := GenS4{F0: false}
                        var v75 uint = v71
                        _ = v74
                        _ = v75
                    } else {
                        v76 := GenS5{F0: -336, F1: (v68 * (-(v68)))}
                        v77 := GenS9{F0: 127, F1: 128}
                        _ = v76
                        _ = v77
                    }
                case -9223372036854775808:
                    v68 = -17
                case 128:
                    if v11 {
                        if v11 {
                            v78 := GenS10{F0: (v6 | v6)}
                            var v79 int8 = v68
                            _ = v78
                            _ = v79
                        } else {
                            v1[128] = (4294967295 < 952)
                            v2 -= v2
                            var v80 int8 = (v68 ^ (-(v68)))
                            _ = v80
                        }
                        if (32768 < 4294967295) {
                            var v81 uint = 18446744073709551615
                            v82 := GenS8{F0: v2}
                            _ = v81
                            _ = v82
                        } else {
                            if v11 {
                                var v83 uint8 = (v2 & 127)
                                _ = v83
                            }
                            v84 := GenS11{F0: 9223372036854775807, F1: v6, F2: int8(v3)}
                            fmt.Println(v2, v71)
                            _ = v84
                        }
                        fmt.Println(v3, v6, v6)
                    }
                default:
                    v85 := v4["\n\t"]
                    _ = v85
                }
            case 32767:
                if (!v11) {
                    if v11 {
                        var v86 bool = false
                        switch (((("a" + ("\n\t" + "hello")) + ("\n\t" + "hello")) + "a") + "\n\t") {
                        case "a":
                            var v87 uint64 = 1
                            _ = v87
                        case "hello":
                            v88 := GenS10{F0: -32768}
                            _ = v88
                        }
                        _ = v86
                    }
                    fmt.Println(v3, v5, v5)
                    if ((1.5 * -2.25) <= -2.25) {
                        for fi89 := 0; fi89 < 20; fi89++ {
                            if v11 {
                                var v90 float64 = v9
                                if v11 {
                                    fmt.Println(fi89, v69)
                                } else {
                                    for fi91 := 0; fi91 < 1; fi91++ {
                                        switch fi89 {
                                        case 128:
                                            if v11 {
                                                v71 = v71
                                                fmt.Println(v11, v2, v68)
                                            }
                                        default:
                                            v90 -= float64(-2.25)
                                        }
                                        if (v11 || v11) {
                                            _ = len(v1)
                                            switch "" {
                                            case "foo bar":
                                                switch (fi89 % ((int(v5)) | 1)) {
                                                case 32767:
                                                    v6 += v6
                                                case 42:
                                                    var v92 uint = (^(v71))
                                                    _ = v92
                                                case -685:
                                                    if false {
                                                        switch fi91 {
                                                        case 463:
                                                            if v11 {
                                                                v93 := []int{(^(fi89)), fi89, (fi89 + 377)}
                                                                if v11 {
                                                                    v69 = (v69 << (((v71 * 226)) & 63))
                                                                    fmt.Println(fi89, v71)
                                                                    var v94 float32 = 1.5
                                                                    _ = v94
                                                                } else {
                                                                    _ = len(v93)
                                                                    v95 := v1[(fi91 ^ fi91)]
                                                                    v5 = (v5 ^ v5)
                                                                    _ = v95
                                                                }
                                                                _ = v93
                                                            } else {
                                                                v96 := []int{4294967295, fi91}
                                                                _ = v96
                                                            }
                                                        case 255:
                                                            v97 := GenS6{F0: (v2 >> ((v71) & 63))}
                                                            _ = v97
                                                        case 42:
                                                            fmt.Println(fi91)
                                                        }
                                                        switch ("a" + "\n\t") {
                                                        case "a":
                                                            fmt.Println(fi91, fi91, v2)
                                                        case "\n\t":
                                                            var v98 uint64 = 0
                                                            _ = v98
                                                        case "hello":
                                                            v99 := v1[(fi89 - 128)]
                                                            _ = v99
                                                        default:
                                                            if (!v11) {
                                                                var v100 int32 = (v5 << ((v71) & 63))
                                                                v101 := GenS9{F0: fi91, F1: fi91}
                                                                _ = v100
                                                                _ = v101
                                                            } else {
                                                                var v102 uint32 = v69
                                                                v69 += v69
                                                                _ = v102
                                                            }
                                                        }
                                                        var v103 uint64 = 42
                                                        _ = v103
                                                    }
                                                default:
                                                    var v104 int = 0
                                                    _ = v104
                                                }
                                            case "a":
                                                if v11 {
                                                    var v105 uint8 = 128
                                                    var v106 float32 = 3.14159
                                                    v71 -= v71
                                                    _ = v105
                                                    _ = v106
                                                } else {
                                                    if ((int64(v69) == 0) || v11) {
                                                        fmt.Println(fi91, v69, v5)
                                                        var v107 bool = v11
                                                        var v108 string = "foo bar"
                                                        _ = v107
                                                        _ = v108
                                                    } else {
                                                        if (uint64(v3) == 0) {
                                                            v2 = v2
                                                        } else {
                                                            v6 = -32768
                                                        }
                                                        v109 := []int{(fi89 & int(v6))}
                                                        v71 = 725
                                                        var v110 int = (fi91 + -452)
                                                        v3 = 796
                                                        _ = v109
                                                        _ = v110
                                                    }
                                                    var v111 int64 = int64(v71)
                                                    _ = v111
                                                }
                                            case "hello":
                                                v112 := v4["\n\t"]
                                                var v113 uint16 = v3
                                                _ = v112
                                                _ = v113
                                            default:
                                                v114 := GenS12{F0: v3, F1: uint64(v3)}
                                                _ = v114
                                            }
                                            v115 := []int{fi89, (fi91 + int(v71)), (fi89 >> ((346) & 63)), 315}
                                            switch fi89 {
                                            case 32768:
                                                var v116 uint8 = v2
                                                _ = v116
                                            }
                                            _ = v115
                                        }
                                        if v11 {
                                            if (0.0 > 0.1) {
                                                var v117 int16 = (v6 >> (((v71 ^ v71)) & 63))
                                                switch fi89 {
                                                case 706:
                                                    fmt.Println(v11, v6, v69)
                                                default:
                                                    if (!v11) {
                                                        var v118 int64 = 0
                                                        var v119 int16 = 22
                                                        var v120 string = "a"
                                                        _ = v118
                                                        _ = v119
                                                        _ = v120
                                                    } else {
                                                        switch "hello" {
                                                        case "a":
                                                            v117 = (v6 << (((v71 - 65535)) & 63))
                                                        case "hello":
                                                            var v121 int32 = (v5 >> ((v71) & 63))
                                                            _ = v121
                                                        case "foo bar":
                                                            v5 -= (-(v5))
                                                        }
                                                    }
                                                }
                                                switch (fi89 & fi91) {
                                                case -9223372036854775808:
                                                    v122 := GenS12{F0: (v3 & uint16(v71)), F1: uint64(v71)}
                                                    _ = v122
                                                case 128:
                                                    v123 := v4["foo bar"]
                                                    _ = v123
                                                case -2147483648:
                                                    var v124 int32 = (v5 | v5)
                                                    _ = v124
                                                }
                                                _ = v117
                                            }
                                            var v125 uint8 = v2
                                            var v126 float32 = 3.14159
                                            _ = v125
                                            _ = v126
                                        } else {
                                            v5 = v5
                                        }
                                        _ = fi91
                                    }
                                    v127 := []int{}
                                    _ = v127
                                }
                                if true {
                                    for fi128 := 0; fi128 < 20; fi128++ {
                                        var v129 int8 = int8(v3)
                                        var v130 float32 = 1.5
                                        fmt.Println(v11, fi89, v68)
                                        _ = fi128
                                        _ = v129
                                        _ = v130
                                    }
                                }
                                _ = v90
                            }
                            for fi131 := 0; fi131 < 6; fi131++ {
                                v5 = 0
                                fmt.Println(v6, v6)
                                _ = fi131
                            }
                            _ = fi89
                        }
                        var v132 int16 = (v6 | -1)
                        if true {
                            fmt.Println(v132)
                        }
                        _ = v132
                    } else {
                        v68 = v68
                    }
                }
            default:
                switch -9223372036854775808 {
                case -1:
                    for fi133 := 0; fi133 < 11; fi133++ {
                        var v134 int32 = v5
                        _ = fi133
                        _ = v134
                    }
                case 42:
                    var v135 bool = v11
                    _ = v135
                case 4294967295:
                    fmt.Println(v11)
                }
            }
        }
        _ = v71
    }
    v136 := GenS13{F0: int8(v3), F1: (v6 ^ v6), F2: (v2 / (((^(v2))) | 1))}
    var v137 uint8 = v2
    switch "hello" {
    case "foo bar":
        v2 = (v137 * v137)
    case "":
        if v11 {
            v6 = v6
            v9 = 3.14159
            if v11 {
                var v138 int8 = v68
                _ = v138
            } else {
                for fi139 := 0; fi139 < 18; fi139++ {
                    if v11 {
                        v5 = 42
                    } else {
                        v6 -= (v6 - (v6 + 681))
                        v2 += v137
                    }
                    var v140 string = ("hello" + ("hello" + ""))
                    _ = fi139
                    _ = v140
                }
                v141 := v4[("\n\t" + "a")]
                v68 -= v68
                _ = v141
            }
        }
    case "hello":
        for fi142 := 0; fi142 < 14; fi142++ {
            var v143 int64 = int64(fi142)
            for fi144 := 0; fi144 < 10; fi144++ {
                v145 := []int{(fi142 + fi142)}
                _ = fi144
                _ = v145
            }
            _ = fi142
            _ = v143
        }
    }
    v146 := GenS14{F0: v69, F1: 9223372036854775807}
    v147 := v4[""]
    v9 = v9
    v69 -= (^(v69))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v137:", v137)
    fmt.Println("v147:", v147)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v136
    _ = v137
    _ = v146
    _ = v147
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v69
    _ = v9
}
