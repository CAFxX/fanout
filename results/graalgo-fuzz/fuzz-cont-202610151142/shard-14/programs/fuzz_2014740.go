package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int64; F2 int8 }
type GenS4 struct { F0 int8; F1 uint16; F2 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int; F1 string; F2 uint64 }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int32; F1 uint8 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 int16; F1 uint32; F2 int16 }
type GenS17 struct { F0 uint; F1 uint64 }
type GenS18 struct { F0 int16; F1 int16; F2 int16 }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 int16; F1 int16 }
type GenS21 struct { F0 uint; F1 uint64; F2 int8 }
func genFn0(p0 int16, p1 uint, p2 string) uint32 {
    p0 -= (p0 % ((32767) | 1))
    if (float64(1e10) == (float64(3.14159) / 1e10)) {
        var v6 uint32 = 1
        if true {
            for fi7 := 0; fi7 < 20; fi7++ {
                v6 = v6
                _ = fi7
            }
        } else {
            if false {
                for fi8 := 0; fi8 < 8; fi8++ {
                    v9 := []int{fi8, int(v6)}
                    for fi10 := 0; fi10 < 11; fi10++ {
                        p0 = 298
                        _ = fi10
                    }
                    p2 = ("hello" + (("hello" + (("a" + p2) + (p2 + "a"))) + "a"))
                    _ = fi8
                    _ = v9
                }
                var v11 uint16 = 743
                _ = v11
            } else {
                if (!false) {
                    fmt.Println(v6, p2)
                    v12 := []int{1, 4294967295, -1}
                    if (0 <= -128) {
                        var v13 int16 = (p0 / ((p0) | 1))
                        for fi14 := 0; fi14 < 18; fi14++ {
                            p2 = p2
                            p2 = "a"
                            if false {
                                fmt.Println(p1, p0, p2)
                                switch fi14 {
                                case 2147483647:
                                    var v15 string = p2
                                    _ = v15
                                default:
                                    v16 := GenS1{F0: 1, F1: (v6 * (^(v6)))}
                                    _ = v16
                                }
                            } else {
                                v17 := GenS2{F0: -25}
                                _ = len(v12)
                                if false {
                                    switch fi14 {
                                    case 4294967295:
                                        v18 := GenS1{F0: 65, F1: v6}
                                        _ = v18
                                    default:
                                        v19 := GenS3{F0: (^(v6)), F1: -1, F2: -89}
                                        _ = v19
                                    }
                                } else {
                                    if (float32(0.1) == 0.1) {
                                        v13 = (^(v13))
                                        fmt.Println(v6)
                                        fmt.Println(fi14)
                                    } else {
                                        v20 := GenS1{F0: 10, F1: (^(v6))}
                                        _ = v20
                                    }
                                    switch p2 {
                                    case "":
                                        v21 := GenS1{F0: 100, F1: 1}
                                        _ = v21
                                    case "foo bar":
                                        for fi22 := 0; fi22 < 7; fi22++ {
                                            var v23 float64 = 0.1
                                            v24 := GenS3{F0: v6, F1: -1, F2: 0}
                                            fmt.Println(p2, p1)
                                            _ = fi22
                                            _ = v23
                                            _ = v24
                                        }
                                    case "hello":
                                        v12 = append(v12, fi14)
                                    }
                                    if (2147483647 != 1) {
                                        var v25 int16 = p0
                                        var v26 int32 = -695
                                        _ = v25
                                        _ = v26
                                    } else {
                                        switch (fi14 - fi14) {
                                        case 32768:
                                            v6 = (v6 % (((v6 & 701)) | 1))
                                        case 42:
                                            v13 -= (v13 << (((p1 << ((p1) & 63))) & 63))
                                        default:
                                            var v27 string = "\n\t"
                                            _ = v27
                                        }
                                    }
                                }
                                _ = v17
                            }
                            _ = fi14
                        }
                        _ = v13
                    } else {
                        fmt.Println(p1, p2, v6)
                        for fi28 := 0; fi28 < 10; fi28++ {
                            fmt.Println(v6, p2, fi28)
                            if (1 != 519) {
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                v6 = v6
                            } else {
                                switch fi28 {
                                case 127:
                                    v29 := GenS3{F0: v6, F1: 9223372036854775807, F2: 42}
                                    _ = v29
                                case -158:
                                    if ((p1 ^ 42) >= (p1 >> ((p1) & 63))) {
                                        var v30 uint8 = 0
                                        fmt.Println(p1, v30, p1)
                                        _ = v30
                                    } else {
                                        v6 = (v6 + v6)
                                        p0 = (p0 >> ((p1) & 63))
                                        var v31 uint32 = 0
                                        _ = v31
                                    }
                                }
                            }
                            p1 = 1
                            _ = fi28
                        }
                        p0 = p0
                    }
                    _ = v12
                } else {
                    fmt.Println(v6, p1)
                    v32 := GenS4{F0: 1, F1: 0, F2: -2147483648}
                    if (27 < 99) {
                        v33 := GenS1{F0: 127, F1: (v6 * v6)}
                        _ = v33
                    } else {
                        p1 += p1
                    }
                    _ = v32
                }
            }
        }
        p0 = (p0 | 0)
        _ = v6
    } else {
        var v34 bool = (1 >= 0)
        var v35 bool = (!((!true) && true))
        _ = v34
        _ = v35
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1() int32 {
    if (-32768 != -32768) {
        if ("\n\t" == "hello") {
            for fi36 := 0; fi36 < 9; fi36++ {
                v37 := GenS5{F0: fi36}
                for fi38 := 0; fi38 < 9; fi38++ {
                    fmt.Println(fi38)
                    _ = fi38
                }
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(fi36)
                case "hello":
                    var v39 string = "foo bar"
                    _ = v39
                }
                _ = fi36
                _ = v37
            }
        }
        v40 := []int{1, 63, 4294967295}
        var v41 uint8 = 42
        _ = v40
        _ = v41
    }
    var v42 int = -331
    v42 = v42
    v42 = v42
    if (!false) {
        v42 = (v42 - -2147483648)
        v42 = v42
    }
    _ = v42
    return 42
}

func main() {
    var v1 string = "hello"
    var v2 uint = 0
    var v3 int32 = -583
    var v4 float64 = 1.5
    v5 := make(map[string]int)
    for fi43 := 0; fi43 < 6; fi43++ {
        for fi44 := 0; fi44 < 12; fi44++ {
            var v45 uint = 528
            switch "" {
            case "hello":
                v1 = v1
            case "":
                fmt.Println(fi44)
            }
            _ = fi44
            _ = v45
        }
        fmt.Println(v3)
        var v46 uint16 = 1
        _ = fi43
        _ = v46
    }
    if true {
        fmt.Println(v2)
        var v47 float64 = v4
        _ = v47
    } else {
        var v48 int = 4294967295
        for fi49 := 0; fi49 < 5; fi49++ {
            for fi50 := 0; fi50 < 4; fi50++ {
                switch (-(fi49)) {
                case 1:
                    switch v1 {
                    case "":
                        var v51 string = (v1 + "\n\t")
                        _ = v51
                    case "a":
                        v1 = v1
                    case "hello":
                        var v52 float64 = 1.5
                        _ = v52
                    }
                case -1:
                    fmt.Println(fi50, v4)
                    switch (fi49 << ((v2) & 63)) {
                    case -1:
                        v53 := []int{(v48 ^ (^(v48))), fi50, 1, (^(fi50))}
                        _ = v53
                    }
                case 127:
                    if false {
                        v4 -= (0.1 + ((0.0 * (0.1 / 3.14159)) * v4))
                        var v54 bool = false
                        v55 := v5[v1]
                        switch fi50 {
                        case 128:
                            var v56 float32 = -2.25
                            _ = v56
                        case -9223372036854775808:
                            v1 = "hello"
                        }
                        switch -87 {
                        case 42:
                            delete(v5, v1)
                        case 451:
                            _ = len(v5)
                        }
                        _ = v54
                        _ = v55
                    } else {
                        if (false || true) {
                            switch fi50 {
                            case 32767:
                                var v57 uint16 = 732
                                _ = v57
                            case 9223372036854775807:
                                v58 := GenS1{F0: -78, F1: 74}
                                _ = v58
                            case 973:
                                if (true && true) {
                                    var v59 int16 = -1
                                    _ = v59
                                }
                            default:
                                v60 := GenS6{F0: 65535, F1: 59, F2: 0}
                                _ = v60
                            }
                            switch fi49 {
                            case 810:
                                switch fi49 {
                                case 9223372036854775807:
                                    v61 := []int{fi50, (fi49 & (fi50 % ((fi50) | 1))), -9223372036854775808}
                                    _ = v61
                                case 32768:
                                    fmt.Println(fi50, fi49, v48)
                                case 127:
                                    v3 = (v3 / ((genFn1()) | 1))
                                default:
                                    v62 := GenS6{F0: 42, F1: 2, F2: 7}
                                    _ = v62
                                }
                            case -102:
                                if (32767 < 65535) {
                                    v63 := GenS7{F0: 915}
                                    if true {
                                        _ = len(v5)
                                    }
                                    _ = v63
                                } else {
                                    v1 = (("\n\t" + "foo bar") + v1)
                                    v64 := GenS8{F0: v48, F1: "foo bar", F2: 1}
                                    _ = v64
                                }
                            default:
                                v65 := []int{(-(fi49)), fi50, (v48 >> (((v2 >> ((0) & 63))) & 63)), (v48 << ((v2) & 63))}
                                _ = v65
                            }
                            v66 := GenS7{F0: uint32(v2)}
                            _ = v66
                        } else {
                            if false {
                                var v67 float64 = (v4 / v4)
                                v68 := GenS9{F0: v3, F1: 0}
                                v69 := GenS10{F0: 4294967295, F1: 1, F2: 2}
                                _ = v67
                                _ = v68
                                _ = v69
                            } else {
                                fmt.Println(v1, v1, fi50)
                                fmt.Println(v3, v4, v4)
                            }
                            if (!(!true)) {
                                var v70 float32 = float32(0.0)
                                _ = v70
                            } else {
                                switch -1 {
                                case -2147483648:
                                    var v71 string = v1
                                    _ = v71
                                case -307:
                                    fmt.Println(fi50)
                                case -663:
                                    var v72 uint16 = 1
                                    _ = v72
                                default:
                                    v1 = ((("a" + "\n\t") + "a") + "")
                                }
                                fmt.Println(v1)
                            }
                        }
                    }
                default:
                    v2 += 18446744073709551615
                }
                v73 := GenS11{F0: (-(v3))}
                v5[v1] = -1
                v4 = float64(0.0)
                _ = fi50
                _ = v73
            }
            var v74 string = "foo bar"
            v1 = (v1 + "foo bar")
            _ = fi49
            _ = v74
        }
        if false {
            if true {
                for fi75 := 0; fi75 < 6; fi75++ {
                    if (594 == 430) {
                        v76 := []int{-1}
                        var v77 float64 = (((((3.14159 / ((v4 + 3.14159) / (3.14159 + 1e10))) - (1.5 + float64(3.14159))) - 3.14159) - v4) * v4)
                        switch "foo bar" {
                        case "a":
                            if len(v76) > 0 {
                                v76[0] = -473
                            }
                        case "":
                            var v78 float32 = (0.1 + (1.5 - 3.14159))
                            _ = v78
                        case "hello":
                            v48 = (v48 | -711)
                        default:
                            var v79 int16 = -1
                            _ = v79
                        }
                        _ = v76
                        _ = v77
                    }
                    var v80 int64 = 0
                    var v81 int8 = 0
                    _ = fi75
                    _ = v80
                    _ = v81
                }
                v82 := []int{v48, 806, v48, (v48 << ((v2) & 63))}
                if (!(-137 < v48)) {
                    v82 = append(v82, v48)
                    if (v2 >= v2) {
                        var v83 bool = false
                        v84 := GenS10{F0: 542, F1: 231, F2: 2}
                        v48 -= v48
                        _ = v83
                        _ = v84
                    }
                }
                _ = v82
            } else {
                v85 := GenS2{F0: 1}
                v48 = (-(v48))
                _ = v85
            }
        }
        _ = v48
    }
    for fi86 := 0; fi86 < 12; fi86++ {
        v3 = (-(v3))
        _ = fi86
    }
    for fi87 := 0; fi87 < 1; fi87++ {
        switch ((v1 + (v1 + v1)) + (v1 + "")) {
        case "a":
            switch fi87 {
            case 4294967295:
                var v88 uint32 = 65535
                _ = v88
            case 148:
                v89 := GenS6{F0: 42, F1: 128, F2: 512}
                _ = v89
            default:
                v3 -= v3
            }
        default:
            switch fi87 {
            case -1:
                var v90 bool = true
                _ = v90
            case 0:
                v91 := []int{(fi87 >> (((v2 % ((v2) | 1))) & 63)), 128, fi87}
                _ = v91
            case 9223372036854775807:
                switch "a" {
                case "\n\t":
                    var v92 uint8 = 0
                    _ = v92
                case "a":
                    fmt.Println(v1, v3)
                default:
                    v93 := GenS12{F0: fi87}
                    _ = v93
                }
            }
        }
        fmt.Println(fi87, v4, fi87)
        for fi94 := 0; fi94 < 1; fi94++ {
            v95 := GenS10{F0: 65535, F1: 1, F2: 2}
            _ = len(v5)
            _ = fi94
            _ = v95
        }
        _ = fi87
    }
    for fi96 := 0; fi96 < 13; fi96++ {
        if ("\n\t" != v1) {
            fmt.Println(v1)
        } else {
            v97 := []int{1, 967, fi96}
            var v98 int = fi96
            for fi99 := 0; fi99 < 10; fi99++ {
                fmt.Println(v3)
                delete(v5, v1)
                _ = fi99
            }
            _ = v97
            _ = v98
        }
        v100 := GenS13{F0: v3}
        switch fi96 {
        case 49:
            if (true || ((v1 != ((("foo bar" + "foo bar") + ("foo bar" + v1)) + (v1 + v1))) || false)) {
                if (539 <= 351) {
                    if true {
                        v1 = (v1 + v1)
                        var v101 int8 = 0
                        _ = v101
                    }
                    if true {
                        switch "\n\t" {
                        case "a":
                            v102 := []int{(fi96 | (fi96 << (((v2 | v2)) & 63))), (fi96 >> ((v2) & 63)), fi96}
                            _ = v102
                        default:
                            _ = len(v5)
                        }
                    }
                    var v103 uint64 = 555
                    _ = v103
                }
                v1 = "a"
            } else {
                var v104 bool = (true && true)
                for fi105 := 0; fi105 < 10; fi105++ {
                    var v106 float32 = float32(v4)
                    if false {
                        v107 := GenS11{F0: 42}
                        _ = v107
                    } else {
                        switch fi96 {
                        case 4294967295:
                            var v108 int32 = (^(v3))
                            _ = v108
                        default:
                            var v109 uint8 = 132
                            v110 := GenS4{F0: 42, F1: 468, F2: fi105}
                            v111 := GenS7{F0: 605}
                            fmt.Println(v106, v1, v1)
                            _ = v109
                            _ = v110
                            _ = v111
                        }
                    }
                    _ = fi105
                    _ = v106
                }
                _ = v104
            }
        case 9223372036854775807:
            v2 = v2
        }
        _ = fi96
        _ = v100
    }
    v112 := []int{4294967295, 32767, 42}
    for fi113 := 0; fi113 < 10; fi113++ {
        fmt.Println(v4)
        switch v1 {
        case "":
            var v114 string = (v1 + "a")
            _ = v114
        default:
            var v115 uint32 = 2147483647
            _ = v115
        }
        if true {
            if ((fi113 % ((fi113) | 1)) >= fi113) {
                v2 = (^(v2))
                if (v1 != v1) {
                    var v116 int8 = 127
                    v1 = v1
                    _ = v116
                } else {
                    fmt.Println(v1)
                    if false {
                        v1 = ("hello" + "foo bar")
                        var v117 bool = true
                        _ = v117
                    } else {
                        fmt.Println(v2)
                        v118 := GenS14{F0: genFn1(), F1: 2}
                        _ = len(v5)
                        _ = v118
                    }
                    v2 = 18446744073709551615
                }
                fmt.Println(v1, v4)
            } else {
                if (!false) {
                    fmt.Println(v2)
                    v3 -= -242
                    for fi119 := 0; fi119 < 2; fi119++ {
                        switch v1 {
                        case "\n\t":
                            var v120 bool = true
                            _ = v120
                        default:
                            fmt.Println(fi119, fi113)
                        }
                        if (!(true || (680 > 255))) {
                            var v121 uint8 = 127
                            var v122 uint8 = 2
                            _ = v121
                            _ = v122
                        }
                        if true {
                            v2 = v2
                            v123 := GenS15{F0: 65535}
                            _ = v123
                        } else {
                            var v124 bool = false
                            var v125 string = v1
                            v126 := GenS16{F0: 463, F1: 2147483647, F2: -40}
                            var v127 bool = true
                            var v128 string = ("a" + v125)
                            var v129 int8 = 99
                            v125 = (v128 + "foo bar")
                            var v130 int32 = (v3 + genFn1())
                            _ = len(v5)
                            _ = v124
                            _ = v125
                            _ = v126
                            _ = v127
                            _ = v128
                            _ = v129
                            _ = v130
                        }
                        if (276 < 723) {
                            v2 = 42
                            v131 := GenS17{F0: v2, F1: 627}
                            _ = v131
                        } else {
                            var v132 bool = false
                            fmt.Println(v2, fi119, v4)
                            v2 = v2
                            _ = v132
                        }
                        v133 := GenS12{F0: (fi119 * fi113)}
                        v2 -= (^(v2))
                        _ = fi119
                        _ = v133
                    }
                }
            }
            var v134 int8 = -81
            var v135 uint16 = 32767
            _ = v134
            _ = v135
        }
        _ = fi113
    }
    if (true || (127 == 105)) {
        if (122 <= 42) {
            v4 -= v4
            if false {
                v3 = (v3 / (((v3 / ((v3) | 1))) | 1))
                var v136 float64 = (1e10 + ((((v4 * float64(1e10)) - 0.1) * (v4 * 0.1)) + (0.0 / 1e10)))
                var v137 uint16 = 730
                _ = v136
                _ = v137
            }
            fmt.Println(v4)
        }
        v138 := GenS18{F0: 1, F1: int16(v3), F2: -566}
        var v139 float64 = v4
        _ = v138
        _ = v139
    } else {
        if false {
            var v140 string = v1
            _ = v140
        } else {
            fmt.Println(v2, v1, v1)
            if (128 >= 298) {
                if (!(423 >= -311)) {
                    for fi141 := 0; fi141 < 18; fi141++ {
                        v2 += (v2 % (((^(v2))) | 1))
                        if true {
                            v1 = ((v1 + v1) + v1)
                            fmt.Println(v4, v2)
                        } else {
                            switch v1 {
                            case "hello":
                                v142 := GenS19{F0: 2}
                                _ = v142
                            }
                            for fi143 := 0; fi143 < 3; fi143++ {
                                v144 := GenS20{F0: 1000, F1: -686}
                                fmt.Println(v4)
                                _ = len(v5)
                                _ = fi143
                                _ = v144
                            }
                            for fi145 := 0; fi145 < 20; fi145++ {
                                var v146 uint32 = 65535
                                v4 = 0.1
                                v2 += 65535
                                v2 += 1
                                var v147 uint = 42
                                _ = fi145
                                _ = v146
                                _ = v147
                            }
                        }
                        _ = fi141
                    }
                    if false {
                        switch 9223372036854775807 {
                        case -1:
                            for fi148 := 0; fi148 < 20; fi148++ {
                                v112 = append(v112, 2147483647)
                                _ = fi148
                            }
                        default:
                            v149 := v5[("foo bar" + (v1 + ("a" + v1)))]
                            _ = v149
                        }
                        switch v1 {
                        case "\n\t":
                            var v150 uint = v2
                            _ = v150
                        case "a":
                            var v151 uint32 = 65535
                            _ = v151
                        default:
                            var v152 float64 = 1e10
                            _ = v152
                        }
                    } else {
                        v1 = "hello"
                        for fi153 := 0; fi153 < 14; fi153++ {
                            _ = len(v5)
                            fmt.Println(fi153, v4)
                            for fi154 := 0; fi154 < 11; fi154++ {
                                v4 += v4
                                v1 = ((v1 + (v1 + "foo bar")) + "foo bar")
                                _ = fi154
                            }
                            _ = fi153
                        }
                    }
                    if (2 < 255) {
                        var v155 uint64 = 1
                        _ = v155
                    }
                } else {
                    v156 := GenS21{F0: (^(v2)), F1: 175, F2: -128}
                    if true {
                        var v157 int = 22
                        _ = v157
                    } else {
                        fmt.Println(v4, v4)
                        v158 := GenS16{F0: int16(v3), F1: 1, F2: -501}
                        _ = v158
                    }
                    var v159 uint16 = 831
                    _ = v156
                    _ = v159
                }
                switch "hello" {
                case "foo bar":
                    v3 += (v3 - genFn1())
                }
            }
        }
        var v160 string = v1
        var v161 bool = (18446744073709551615 < 536)
        _ = v160
        _ = v161
    }
    switch 127 {
    case 2147483647:
        for fi162 := 0; fi162 < 4; fi162++ {
            v4 += 3.14159
            var v163 float32 = 0.1
            _ = fi162
            _ = v163
        }
    default:
        v164 := GenS18{F0: -32768, F1: 32767, F2: 0}
        _ = v164
    }
    var v165 int8 = 1
    v5[v1] = 97
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v112:", len(v112), v112)
    fmt.Println("v165:", v165)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v112
    _ = v165
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
