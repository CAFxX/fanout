package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 bool; F2 bool }
type GenS8 struct { F0 uint; F1 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint8; F1 uint8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS13 struct { F0 int8; F1 uint64 }
type GenS14 struct { F0 int16; F1 string; F2 int16 }
func genFn0(p0 uint, p1 int64, p2 uint64) string {
    v6 := GenS1{F0: (p1 ^ (p1 | (p1 << ((18446744073709551615) & 63))))}
    if ((p2 - 906) > 1) {
        p0 -= p0
        v7 := GenS1{F0: (p1 / ((p1) | 1))}
        p2 = 18446744073709551615
        _ = v7
    }
    p2 += p2
    p2 = 948
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ("foo bar" + "foo bar")
}

func main() {
    var v1 uint32 = 264
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]bool)
    var v5 int16 = -443
    switch 32767 {
    case 664:
        v8 := GenS2{F0: (v1 | 322), F1: 1}
        _ = v8
    case 552:
        switch "a" {
        case "a":
            var v9 string = ((("a" + "foo bar") + ("" + "foo bar")) + "a")
            _ = v9
        case "hello":
            v5 = 42
        case "foo bar":
            var v10 uint16 = 0
            _ = v10
        }
    case 651:
        var v11 uint32 = v1
        _ = v11
    default:
        v12 := GenS2{F0: 65535, F1: v1}
        _ = v12
    }
    switch 32767 {
    case 9223372036854775807:
        var v13 float32 = (3.14159 + 3.14159)
        _ = v13
    case 128:
        var v14 float64 = 3.14159
        _ = v14
    case 127:
        v15 := GenS3{F0: (v1 ^ v1)}
        _ = v15
    }
    var v16 string = "a"
    switch int(v1) {
    case -708:
        fmt.Println(v5, v1, v1)
    case 32768:
        switch v16 {
        case "":
            var v17 float64 = (1e10 + 1e10)
            _ = v17
        case "\n\t":
            if false {
                v5 = (v5 >> ((795) & 63))
                if true {
                    fmt.Println(v16)
                    fmt.Println(v5)
                }
                v18 := GenS1{F0: 1}
                _ = v18
            } else {
                var v19 uint32 = (^(v1))
                _ = v19
            }
        }
    default:
        v20 := []int{-987}
        _ = v20
    }
    var v21 uint = 255
    v16 = ("a" + v16)
    v3[v1] = 567
    var v22 int16 = (v5 >> ((42) & 63))
    var v23 int8 = -128
    var v24 int = 2147483647
    switch (v16 + v16) {
    case "a":
        v25 := GenS1{F0: -1}
        _ = v25
    case "":
        v26 := GenS4{F0: (v24 * v24), F1: -958}
        _ = v26
    case "hello":
        var v27 uint = v21
        _ = v27
    default:
        for fi28 := 0; fi28 < 8; fi28++ {
            v29 := GenS4{F0: (v24 % (((v24 ^ (fi28 / ((-14) | 1)))) | 1)), F1: fi28}
            _ = fi28
            _ = v29
        }
    }
    v30 := GenS5{F0: 0}
    var v31 float64 = 1e10
    for fi32 := 0; fi32 < 7; fi32++ {
        for fi33 := 0; fi33 < 4; fi33++ {
            var v34 float32 = 0.1
            v35 := []int{42, v24, -772, -77}
            v5 -= 1
            _ = fi33
            _ = v34
            _ = v35
        }
        if (!false) {
            if true {
                switch v24 {
                case -2147483648:
                    fmt.Println(v1)
                default:
                    var v36 uint16 = 197
                    _ = v36
                }
            } else {
                if true {
                    switch 2147483647 {
                    case 127:
                        var v37 uint8 = 42
                        _ = v37
                    case -1:
                        if (!true) {
                            if (1 != -1) {
                                v2[(fi32 - (v24 / ((fi32) | 1)))] = v16
                            } else {
                                var v38 float64 = v31
                                _ = v38
                            }
                            v39 := GenS2{F0: (v1 + v1), F1: v1}
                            v40 := GenS5{F0: -1}
                            _ = v39
                            _ = v40
                        }
                    case 32767:
                        switch v24 {
                        case 535:
                            v41 := GenS6{F0: 198}
                            _ = v41
                        case 42:
                            var v42 uint32 = (v1 >> (((^(v21))) & 63))
                            _ = v42
                        case 1:
                            if (v21 < (v21 / ((v21) | 1))) {
                                for fi43 := 0; fi43 < 17; fi43++ {
                                    v23 -= v23
                                    _ = fi43
                                }
                                if (-1 != -2147483648) {
                                    var v44 int8 = v23
                                    var v45 bool = true
                                    _ = v44
                                    _ = v45
                                } else {
                                    v23 += (v23 >> (((v21 ^ 2)) & 63))
                                }
                                switch fi32 {
                                case 1:
                                    v1 = uint32(v21)
                                default:
                                    if true {
                                        if false {
                                            fmt.Println(v5, v5, v31)
                                            v46 := GenS6{F0: uint8(v1)}
                                            _ = v46
                                        } else {
                                            var v47 float64 = 1.5
                                            switch "a" {
                                            case "foo bar":
                                                var v48 uint8 = 127
                                                _ = v48
                                            case "":
                                                _ = len(v4)
                                            default:
                                                var v49 uint64 = 0
                                                _ = v49
                                            }
                                            if true {
                                                fmt.Println(fi32, v47)
                                                var v50 uint8 = 1
                                                _ = v50
                                            }
                                            _ = v47
                                        }
                                        var v51 int8 = int8(v21)
                                        v31 = v31
                                        _ = v51
                                    } else {
                                        for fi52 := 0; fi52 < 17; fi52++ {
                                            var v53 int8 = 42
                                            v54 := GenS4{F0: (^(fi32)), F1: v24}
                                            v55 := GenS2{F0: 187, F1: v1}
                                            switch -209 {
                                            case 32768:
                                                var v56 uint = (v21 >> (((v21 % (((v21 / ((2) | 1))) | 1))) & 63))
                                                _ = v56
                                            case 32767:
                                                if true {
                                                    var v57 int32 = 95
                                                    var v58 uint16 = 255
                                                    switch (v24 << (((v21 & v21)) & 63)) {
                                                    case 2147483647:
                                                        var v59 uint = (v21 + 1)
                                                        _ = v59
                                                    case -2147483648:
                                                        var v60 int32 = (v57 ^ -2147483648)
                                                        _ = v60
                                                    case 228:
                                                        var v61 int32 = 42
                                                        _ = v61
                                                    }
                                                    _ = v57
                                                    _ = v58
                                                }
                                            case 0:
                                                v31 = (((v31 + (v31 / 1.5)) + (v31 * (3.14159 + v31))) * -2.25)
                                            }
                                            _ = fi52
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        }
                                        if false {
                                            var v62 int = (-(v24))
                                            if ((v31 / float64(0.1)) <= float64(0.0)) {
                                                v63 := []int{}
                                                v23 -= (v23 % ((v23) | 1))
                                                if false {
                                                    if true {
                                                        _ = len(v63)
                                                        fmt.Println(v24, v1)
                                                    }
                                                    var v64 uint16 = 0
                                                    _ = v64
                                                }
                                                _ = v63
                                            } else {
                                                if (v21 > 4294967295) {
                                                    v65 := GenS2{F0: v1, F1: v1}
                                                    _ = v65
                                                } else {
                                                    var v66 bool = (((v21 | v21) <= (v21 >> ((v21) & 63))) && true)
                                                    v67 := GenS7{F0: v16, F1: (42 > 18446744073709551615), F2: true}
                                                    v68 := GenS8{F0: (v21 + 255), F1: v66}
                                                    _ = v66
                                                    _ = v67
                                                    _ = v68
                                                }
                                            }
                                            switch (((("foo bar" + "") + "hello") + v16) + v16) {
                                            case "\n\t":
                                                v22 += 0
                                            default:
                                                for fi69 := 0; fi69 < 18; fi69++ {
                                                    switch 4294967295 {
                                                    case 128:
                                                        var v70 float32 = -2.25
                                                        _ = v70
                                                    case 9223372036854775807:
                                                        v71 := GenS5{F0: -873}
                                                        _ = v71
                                                    case 1:
                                                        var v72 string = v16
                                                        _ = v72
                                                    default:
                                                        if (32767 == 319) {
                                                            _ = len(v3)
                                                            v31 = 0.0
                                                            v62 += fi69
                                                        } else {
                                                            if true {
                                                                var v73 uint16 = 2
                                                                v1 = 1
                                                                _ = v73
                                                            }
                                                            var v74 int16 = v5
                                                            switch fi32 {
                                                            case 9223372036854775807:
                                                                fmt.Println(v21, v16, v31)
                                                            case 32767:
                                                                if (1 >= 0) {
                                                                    switch ((v16 + (v16 + v16)) + v16) {
                                                                    case "foo bar":
                                                                        var v75 int16 = (v22 - 0)
                                                                        _ = v75
                                                                    case "a":
                                                                        var v76 string = (v16 + ((("hello" + v16) + v16) + "foo bar"))
                                                                        _ = v76
                                                                    case "":
                                                                        var v77 float32 = 0.0
                                                                        _ = v77
                                                                    }
                                                                    v23 += 100
                                                                    v78 := []int{}
                                                                    _ = v78
                                                                }
                                                            case 2147483647:
                                                                v79 := []int{v62}
                                                                switch -1 {
                                                                case 32767:
                                                                    var v80 int8 = (v23 ^ v23)
                                                                    _ = v80
                                                                default:
                                                                    var v81 string = v16
                                                                    _ = v81
                                                                }
                                                                _ = v79
                                                            }
                                                            _ = v74
                                                        }
                                                    }
                                                    _ = fi69
                                                }
                                            }
                                            _ = v62
                                        }
                                        v21 += (v21 ^ v21)
                                    }
                                }
                            } else {
                                for fi82 := 0; fi82 < 3; fi82++ {
                                    var v83 int8 = v23
                                    v22 = 0
                                    if ("" != (v16 + v16)) {
                                        if (float64(-2.25) < float64(((float32((0.1 + 3.14159)) / 1.0) + float32((0.1 / (v31 + v31)))))) {
                                            var v84 uint64 = 1
                                            _ = v84
                                        } else {
                                            v85 := GenS9{F0: ((("\n\t" + v16) + (("" + v16) + v16)) + v16)}
                                            v2[2147483647] = v16
                                            switch (fi82 | v24) {
                                            case -284:
                                                v86 := v4[v16]
                                                _ = v86
                                            default:
                                                v87 := GenS10{F0: 139, F1: 163}
                                                _ = v87
                                            }
                                            _ = v85
                                        }
                                    } else {
                                        if true {
                                            var v88 string = v16
                                            _ = v88
                                        } else {
                                            fmt.Println(v83, v83)
                                            v89 := []int{}
                                            v90 := GenS11{F0: v23}
                                            v91 := GenS12{F0: false, F1: 284, F2: 255}
                                            _ = v89
                                            _ = v90
                                            _ = v91
                                        }
                                        var v92 int64 = -121
                                        if false {
                                            switch (v24 << (((^(v21))) & 63)) {
                                            case 127:
                                                if (true || ((v23 >> ((48) & 63)) >= 0)) {
                                                    v16 = v16
                                                    if ((((-1 != (v83 ^ v83)) && (100 != (v83 | (v83 - -1)))) || true) && true) {
                                                        var v93 uint8 = 239
                                                        _ = v93
                                                    }
                                                    v92 = v92
                                                }
                                            case 0:
                                                v31 = 0.1
                                            case 32767:
                                                v94 := GenS8{F0: v21, F1: false}
                                                _ = v94
                                            }
                                            var v95 uint64 = 210
                                            var v96 int = (fi32 ^ int(v92))
                                            var v97 string = "hello"
                                            if (686 < 571) {
                                                v95 = v95
                                            } else {
                                                switch v97 {
                                                case "foo bar":
                                                    _ = len(v3)
                                                case "a":
                                                    v3[v1] = 1
                                                default:
                                                    fmt.Println(v24, v16, v96)
                                                }
                                                var v98 int64 = int64(v1)
                                                fmt.Println(v21, fi82)
                                                v99 := v2[fi32]
                                                if true {
                                                    v100 := []int{42, 592, 32768}
                                                    _ = v100
                                                }
                                                _ = v98
                                                _ = v99
                                            }
                                            v101 := []int{}
                                            if ((3.14159 >= (-2.25 + (1.5 - 1e10))) && false) {
                                                if ((v1 & v1) <= 2147483647) {
                                                    if (1e10 <= ((0.1 * 1.5) * 0.0)) {
                                                        var v102 uint = v21
                                                        v22 += (v22 * (v5 | v22))
                                                        var v103 bool = ((v21 % (((v102 & (v21 ^ 18446744073709551615))) | 1)) > v102)
                                                        _ = v102
                                                        _ = v103
                                                    } else {
                                                        var v104 uint8 = 26
                                                        if ((true || (true || true)) && true) {
                                                            var v105 float32 = 3.14159
                                                            _ = v105
                                                        } else {
                                                            v97 = v97
                                                        }
                                                        v106 := GenS3{F0: v1}
                                                        _ = v104
                                                        _ = v106
                                                    }
                                                    var v107 uint = 2
                                                    var v108 uint64 = (v95 / (((v95 << (((v107 % (((v107 & (v21 & 255))) | 1))) & 63))) | 1))
                                                    if (!(!false)) {
                                                        switch v16 {
                                                        case "a":
                                                            v109 := GenS5{F0: v92}
                                                            _ = v109
                                                        default:
                                                            _ = len(v101)
                                                        }
                                                        v110 := GenS13{F0: (-(v83)), F1: (v108 >> ((v21) & 63))}
                                                        var v111 uint8 = 211
                                                        _ = v110
                                                        _ = v111
                                                    } else {
                                                        v112 := v2[(fi32 | (fi82 - 0))]
                                                        _ = v112
                                                    }
                                                    _ = v107
                                                    _ = v108
                                                } else {
                                                    v1 -= (v1 | 4294967295)
                                                    var v113 uint = v21
                                                    _ = len(v4)
                                                    _ = v113
                                                }
                                                switch ("hello" + v97) {
                                                case "foo bar":
                                                    v114 := GenS13{F0: (v83 & -1), F1: v95}
                                                    _ = v114
                                                case "\n\t":
                                                    fmt.Println(v22, v21)
                                                    fmt.Println(v92)
                                                    v115 := GenS6{F0: 139}
                                                    _ = v115
                                                case "":
                                                    v96 = int(v23)
                                                default:
                                                    v96 = 4294967295
                                                }
                                                v116 := v2[(fi82 - fi32)]
                                                _ = v116
                                            }
                                            v4[(((v16 + v16) + v97) + v16)] = (true || true)
                                            var v117 string = (("foo bar" + (v16 + v97)) + (v16 + (v97 + ("" + v97))))
                                            _ = v101
                                            _ = v117
                                            _ = v95
                                            _ = v96
                                            _ = v97
                                        } else {
                                            v118 := GenS14{F0: (v5 - v5), F1: (v16 + v16), F2: (v22 ^ (v22 - v5))}
                                            fmt.Println(v22, v21)
                                            if ((v16 + "") > "") {
                                                v16 = (("foo bar" + v16) + (v16 + ("a" + "\n\t")))
                                            } else {
                                                fmt.Println(v92, v22, v24)
                                            }
                                            _ = v118
                                        }
                                        _ = v92
                                    }
                                    v5 = (v22 + v22)
                                    _ = fi82
                                    _ = v83
                                }
                                v22 = v22
                                var v119 bool = false
                                _ = v119
                            }
                        }
                    default:
                        if false {
                            if (1 >= 4294967295) {
                                v5 = v5
                                _ = len(v2)
                            } else {
                                v120 := GenS11{F0: 0}
                                v121 := GenS3{F0: 369}
                                for fi122 := 0; fi122 < 3; fi122++ {
                                    v24 += fi32
                                    switch ((v16 + ((v16 + "hello") + "a")) + "foo bar") {
                                    case "foo bar":
                                        var v123 uint16 = uint16(v1)
                                        _ = v123
                                    case "hello":
                                        v23 = v23
                                    }
                                    if false {
                                        var v124 uint = uint(v1)
                                        var v125 int32 = 0
                                        _ = v124
                                        _ = v125
                                    }
                                    var v126 bool = (!false)
                                    switch v24 {
                                    case -9223372036854775808:
                                        v16 = (v16 + "")
                                    }
                                    var v127 int64 = -770
                                    v23 -= (v23 >> ((v21) & 63))
                                    _ = fi122
                                    _ = v126
                                    _ = v127
                                }
                                _ = v120
                                _ = v121
                            }
                            v128 := []int{fi32, v24}
                            _ = v128
                        } else {
                            v129 := GenS3{F0: v1}
                            for fi130 := 0; fi130 < 13; fi130++ {
                                var v131 uint = 2
                                _ = fi130
                                _ = v131
                            }
                            _ = v129
                        }
                    }
                }
            }
        } else {
            v22 = v5
            if ((^(v1)) <= (v1 - (v1 >> ((255) & 63)))) {
                v16 = v16
                var v132 uint64 = 232
                _ = v132
            } else {
                switch (v24 * fi32) {
                case 128:
                    v133 := []int{v24}
                    _ = v133
                case 2147483647:
                    var v134 uint16 = 0
                    _ = v134
                }
                fmt.Println(v24, v1)
                v135 := v3[(v1 * (^(v1)))]
                _ = v135
            }
        }
        for fi136 := 0; fi136 < 3; fi136++ {
            if true {
                delete(v2, (v24 / (((fi136 | fi32)) | 1)))
                v23 -= 42
                fmt.Println(v5)
            }
            _ = fi136
        }
        _ = fi32
    }
    v31 = -2.25
    if true {
        if (!(42 < 42)) {
            var v137 uint32 = 65535
            _ = v137
        }
        v2[32767] = v16
        var v138 float64 = v31
        _ = v138
    }
    v22 = (v22 >> ((v21) & 63))
    v22 += v22
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
