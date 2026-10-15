package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 string; F1 int; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 uint8; F1 int16; F2 int8 }
type GenS8 struct { F0 int8; F1 bool }
type GenS9 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS10 struct { F0 int; F1 int16 }
type GenS11 struct { F0 uint8; F1 string; F2 int }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int32; F1 uint8; F2 int }
type GenS14 struct { F0 int8; F1 int }
type GenS15 struct { F0 int; F1 uint8; F2 string }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 int16; F1 int16 }
type GenS18 struct { F0 int64; F1 uint32; F2 bool }
type GenS19 struct { F0 int32; F1 uint64 }
func genFn0() uint {
    v6 := GenS1{F0: 141, F1: 32767, F2: 0}
    switch (("a" + "foo bar") + "foo bar") {
    case "hello":
        for fi7 := 0; fi7 < 14; fi7++ {
            if true {
                var v8 uint16 = 688
                v9 := GenS2{F0: fi7, F1: 127}
                _ = v8
                _ = v9
            }
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 bool = false
                fmt.Println(v11, fi10, fi10)
                v12 := []int{(fi7 ^ fi7), fi7, 32767}
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi7
        }
    case "":
        var v13 float32 = (float32(0.1) + 0.1)
        _ = v13
    default:
        var v14 float32 = (1e10 - 1.5)
        _ = v14
    }
    _ = v6
    return 572
}
func genFn1(p0 bool, p1 int16, p2 int16) int16 {
    v15 := []int{128, 128, 128, 434}
    var v16 float32 = 1e10
    v17 := GenS3{F0: "\n\t", F1: -352, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v17
    return (p2 & (p1 ^ (p2 + p2)))
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]int)
    var v3 uint = 42
    v4 := make(map[string]string)
    var v5 int = 127
    v4["a"] = "hello"
    v3 += (^(v3))
    if true {
        for fi18 := 0; fi18 < 4; fi18++ {
            fmt.Println(v1, v1, fi18)
            switch "\n\t" {
            case "\n\t":
                for fi19 := 0; fi19 < 10; fi19++ {
                    v20 := GenS4{F0: 856}
                    var v21 uint8 = 1
                    v3 -= v3
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
            _ = fi18
        }
        switch (v5 - v5) {
        case 127:
            for fi22 := 0; fi22 < 8; fi22++ {
                if (false && true) {
                    v3 = (^(v3))
                }
                _ = fi22
            }
        case 0:
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 uint64 = 0
                v25 := GenS1{F0: 2, F1: (v1 / (((v1 % ((42) | 1))) | 1)), F2: (v5 >> (((v3 % ((2) | 1))) & 63))}
                for fi26 := 0; fi26 < 13; fi26++ {
                    var v27 string = "a"
                    _ = fi26
                    _ = v27
                }
                _ = fi23
                _ = v24
                _ = v25
            }
        case 9223372036854775807:
            for fi28 := 0; fi28 < 20; fi28++ {
                if (817 <= 924) {
                    if ((v1 > 0) && true) {
                        var v29 uint8 = 253
                        v30 := GenS5{F0: (("a" + "") + "foo bar")}
                        _ = v29
                        _ = v30
                    }
                    var v31 uint8 = 255
                    _ = v31
                } else {
                    var v32 string = (("hello" + "foo bar") + "hello")
                    v33 := GenS6{F0: -1, F1: 240}
                    if false {
                        v32 = ((("a" + ("\n\t" + v32)) + v32) + ("\n\t" + v32))
                        if false {
                            v34 := GenS1{F0: 171, F1: v1, F2: (^(v5))}
                            _ = v34
                        } else {
                            var v35 bool = false
                            _ = v35
                        }
                        _ = len(v2)
                    } else {
                        for fi36 := 0; fi36 < 8; fi36++ {
                            v1 = 32767
                            v37 := GenS7{F0: 128, F1: genFn1(((1.5 + 1.5) == 1e10), 42, -1), F2: -128}
                            _ = fi36
                            _ = v37
                        }
                        switch (("hello" + ("" + ("a" + ("hello" + v32)))) + v32) {
                        case "":
                            delete(v4, v32)
                        default:
                            for fi38 := 0; fi38 < 20; fi38++ {
                                var v39 uint16 = (v1 << ((4294967295) & 63))
                                _ = fi38
                                _ = v39
                            }
                        }
                        switch 4294967295 {
                        case -500:
                            var v40 uint64 = 2
                            _ = v40
                        case 146:
                            v32 = ((((v32 + "a") + "\n\t") + v32) + v32)
                        }
                    }
                    _ = v32
                    _ = v33
                }
                v1 = v1
                if false {
                    var v41 int16 = 32767
                    v42 := GenS8{F0: 58, F1: (!(!true))}
                    switch ("" + "a") {
                    case "\n\t":
                        fmt.Println(v3, fi28)
                    case "foo bar":
                        if (!true) {
                            if false {
                                fmt.Println(fi28)
                                var v43 uint8 = 33
                                var v44 float32 = (((-2.25 * 1.5) - (0.0 - ((1.5 * 0.1) - (1.5 - 1e10)))) - float32(0.0))
                                _ = v43
                                _ = v44
                            } else {
                                var v45 uint16 = v1
                                var v46 uint8 = 0
                                _ = v45
                                _ = v46
                            }
                            var v47 bool = false
                            v48 := GenS6{F0: 1000, F1: 666}
                            _ = v47
                            _ = v48
                        } else {
                            var v49 int32 = 2147483647
                            var v50 int64 = 42
                            _ = v49
                            _ = v50
                        }
                    }
                    _ = v41
                    _ = v42
                } else {
                    v1 = v1
                }
                _ = fi28
            }
        default:
            v5 += v5
        }
    } else {
        if (genFn1((828 >= 1), -318, -1) >= 18) {
            v5 -= (v5 / ((128) | 1))
        } else {
            for fi51 := 0; fi51 < 8; fi51++ {
                v3 += v3
                switch (("" + "a") + "") {
                case "\n\t":
                    v52 := GenS5{F0: "a"}
                    _ = v52
                case "hello":
                    if (-4 > 32767) {
                        for fi53 := 0; fi53 < 3; fi53++ {
                            delete(v2, (fi53 ^ 32767))
                            _ = fi53
                        }
                        v3 = v3
                        v3 = v3
                    } else {
                        fmt.Println(fi51, v3)
                        for fi54 := 0; fi54 < 2; fi54++ {
                            v55 := GenS9{F0: (v1 / (((^(v1))) | 1)), F1: 6, F2: 0}
                            _ = fi54
                            _ = v55
                        }
                    }
                default:
                    for fi56 := 0; fi56 < 5; fi56++ {
                        delete(v4, "a")
                        _ = fi56
                    }
                }
                var v57 uint = 2
                _ = fi51
                _ = v57
            }
            _ = len(v4)
            var v58 bool = false
            _ = v58
        }
        var v59 string = (("" + "foo bar") + "\n\t")
        _ = v59
    }
    v60 := GenS10{F0: v5, F1: 32767}
    fmt.Println(v5, v1, v5)
    for fi61 := 0; fi61 < 14; fi61++ {
        if (2 == (v1 | (v1 / ((uint16(v3)) | 1)))) {
            for fi62 := 0; fi62 < 20; fi62++ {
                var v63 float32 = ((-2.25 - 3.14159) * ((1e10 * -2.25) * (1e10 - (-2.25 - 3.14159))))
                var v64 uint = 1
                v5 = (v5 | (^(fi61)))
                v65 := GenS11{F0: 144, F1: "a", F2: -2147483648}
                _ = fi62
                _ = v63
                _ = v64
                _ = v65
            }
            v5 -= -793
            if (!false) {
                switch 127 {
                case 823:
                    if false {
                        v66 := []int{(fi61 ^ 0), (fi61 >> (((^(v3))) & 63))}
                        fmt.Println(v3)
                        v3 += 42
                        _ = v66
                    } else {
                        for fi67 := 0; fi67 < 1; fi67++ {
                            v68 := GenS12{F0: 4294967295}
                            _ = fi67
                            _ = v68
                        }
                        v3 = (^(v3))
                        switch fi61 {
                        case 989:
                            v69 := GenS8{F0: -128, F1: false}
                            _ = v69
                        default:
                            fmt.Println(v1, v5)
                        }
                    }
                }
                if ((((0.1 * 3.14159) + (3.14159 / 1e10)) < (0.1 - 1.5)) || true) {
                    var v70 string = ("foo bar" + "")
                    for fi71 := 0; fi71 < 11; fi71++ {
                        var v72 uint8 = 0
                        if true {
                            fmt.Println(v1, v3)
                        } else {
                            v73 := GenS13{F0: -170, F1: 128, F2: (v5 & v5)}
                            _ = v73
                        }
                        if (-1 == 42) {
                            if (4294967295 < 18446744073709551615) {
                                var v74 int16 = 1000
                                var v75 int8 = -93
                                var v76 float64 = -2.25
                                var v77 float32 = 0.1
                                switch v5 {
                                case -9223372036854775808:
                                    var v78 int16 = 42
                                    _ = v78
                                case -48:
                                    var v79 uint8 = v72
                                    _ = v79
                                default:
                                    fmt.Println(v3, v74)
                                }
                                if (!true) {
                                    v77 = float32(v76)
                                    switch (-(v5)) {
                                    case 127:
                                        _ = len(v2)
                                    case 554:
                                        switch (fi61 | -9223372036854775808) {
                                        case 1:
                                            var v80 uint32 = 2147483648
                                            _ = v80
                                        case -9223372036854775808:
                                            fmt.Println(v75, v74)
                                        default:
                                            v81 := GenS11{F0: 127, F1: (("\n\t" + (v70 + v70)) + v70), F2: 2147483647}
                                            _ = v81
                                        }
                                    default:
                                        if false {
                                            var v82 string = ""
                                            switch v82 {
                                            case "\n\t":
                                                var v83 uint16 = v1
                                                _ = v83
                                            default:
                                                var v84 uint16 = v1
                                                _ = v84
                                            }
                                            _ = v82
                                        }
                                    }
                                } else {
                                    switch fi61 {
                                    case 128:
                                        var v85 uint16 = (v1 & 65535)
                                        _ = v85
                                    case 42:
                                        v76 -= ((float64((float32(-2.25) / (v77 + v77))) / (((v76 - (float64(v77) / v76)) + ((((v76 / 1e10) - v76) * ((3.14159 - 1e10) * (0.0 * (1.5 * float64(((1e10 * (0.1 * v77)) - (0.1 + v77))))))) - (0.1 / 1.5))) - float64(v77))) * v76)
                                    case -2147483648:
                                        if true {
                                            v86 := GenS14{F0: v75, F1: (-(v5))}
                                            _ = v86
                                        } else {
                                            switch v70 {
                                            case "\n\t":
                                                var v87 uint = (v3 << (((v3 * 1)) & 63))
                                                _ = v87
                                            case "":
                                                v76 += float64(v77)
                                            case "foo bar":
                                                var v88 uint8 = uint8(v1)
                                                _ = v88
                                            }
                                        }
                                    default:
                                        var v89 int8 = (v75 * (v75 / (((v75 - v75)) | 1)))
                                        _ = v89
                                    }
                                    var v90 uint8 = 87
                                    _ = v90
                                }
                                fmt.Println(v5)
                                _ = v74
                                _ = v75
                                _ = v76
                                _ = v77
                            }
                        } else {
                            v3 += (^(v3))
                            fmt.Println(v72, v1, v3)
                            v70 = "hello"
                        }
                        _ = len(v4)
                        _ = fi71
                        _ = v72
                    }
                    _ = v70
                }
                v91 := []int{890, (v5 & (v5 | (fi61 | fi61))), (v5 + 32767)}
                _ = v91
            }
        } else {
            if (float32(-2.25) >= (-2.25 / 3.14159)) {
                v92 := GenS12{F0: 2}
                switch (v5 % (((^(v5))) | 1)) {
                case 42:
                    fmt.Println(fi61, v3)
                case 32767:
                    fmt.Println(v1, v1)
                default:
                    if (false && false) {
                        v1 = (^(v1))
                        if (-432 >= (fi61 | int(v1))) {
                            v93 := v2[(fi61 % ((v5) | 1))]
                            _ = v93
                        }
                    } else {
                        fmt.Println(v5, fi61, v5)
                    }
                }
                if (true || true) {
                    v1 = (v1 % (((v1 & (v1 ^ (v1 - v1)))) | 1))
                    var v94 int = (fi61 >> ((genFn0()) & 63))
                    _ = v94
                } else {
                    if false {
                        v3 = (v3 << (((v3 / ((v3) | 1))) & 63))
                        for fi95 := 0; fi95 < 2; fi95++ {
                            var v96 float32 = (((float32(-2.25) - 0.0) + 1.5) + -2.25)
                            switch "foo bar" {
                            case "":
                                v97 := GenS12{F0: 4294967295}
                                _ = v97
                            case "a":
                                switch (^(v5)) {
                                case 1:
                                    v5 = (-(fi61))
                                case 403:
                                    v98 := []int{-1, fi95, 4294967295, 2147483647}
                                    _ = v98
                                default:
                                    v99 := []int{}
                                    _ = v99
                                }
                            }
                            switch (("\n\t" + ("\n\t" + ((((("a" + "foo bar") + ("" + (("a" + ("foo bar" + "foo bar")) + "\n\t"))) + "foo bar") + "") + "hello"))) + "hello") {
                            case "hello":
                                v100 := []int{}
                                _ = v100
                            case "":
                                switch "foo bar" {
                                case "hello":
                                    fmt.Println(v5, v96, v96)
                                default:
                                    switch v5 {
                                    case 128:
                                        var v101 int64 = 1
                                        _ = v101
                                    case 4294967295:
                                        v3 += v3
                                    case 1:
                                        var v102 bool = (int8(v3) <= 1)
                                        _ = v102
                                    default:
                                        v103 := v2[(fi95 - v5)]
                                        _ = v103
                                    }
                                }
                            default:
                                v104 := GenS5{F0: "foo bar"}
                                _ = v104
                            }
                            if (v1 >= v1) {
                                var v105 uint32 = 42
                                fmt.Println(v96)
                                var v106 string = "foo bar"
                                var v107 int64 = 42
                                if (true || false) {
                                    var v108 uint = (^(v3))
                                    var v109 uint = v108
                                    v5 -= fi61
                                    _ = v108
                                    _ = v109
                                }
                                v110 := GenS15{F0: 32767, F1: 128, F2: v106}
                                switch "hello" {
                                case "a":
                                    var v111 bool = (!true)
                                    _ = v111
                                }
                                v112 := GenS12{F0: 22}
                                fmt.Println(v5, v106, v96)
                                switch -559 {
                                case -884:
                                    v5 = fi95
                                case -1:
                                    v113 := []int{(fi95 + 128), (fi61 | 9223372036854775807)}
                                    _ = v113
                                default:
                                    var v114 int = (^(v5))
                                    _ = v114
                                }
                                _ = v105
                                _ = v106
                                _ = v107
                                _ = v110
                                _ = v112
                            }
                            _ = fi95
                            _ = v96
                        }
                    }
                }
                _ = v92
            } else {
                for fi115 := 0; fi115 < 5; fi115++ {
                    v116 := GenS16{F0: 18446744073709551615}
                    delete(v4, ("hello" + ""))
                    v117 := GenS7{F0: 255, F1: 32767, F2: 40}
                    v5 += (v5 & (^(v5)))
                    v1 += (v1 >> (((v3 | v3)) & 63))
                    _ = fi115
                    _ = v116
                    _ = v117
                }
                var v118 int16 = 32767
                _ = v118
            }
            v119 := GenS10{F0: 107, F1: 42}
            _ = v119
        }
        _ = fi61
    }
    for fi120 := 0; fi120 < 16; fi120++ {
        var v121 float64 = 1.5
        _ = fi120
        _ = v121
    }
    for fi122 := 0; fi122 < 11; fi122++ {
        for fi123 := 0; fi123 < 13; fi123++ {
            switch "hello" {
            case "a":
                v5 = -1
            }
            v124 := []int{}
            _ = fi123
            _ = v124
        }
        for fi125 := 0; fi125 < 20; fi125++ {
            if false {
                v2[723] = int(v3)
            } else {
                _ = len(v2)
                switch "hello" {
                case "hello":
                    if (!true) {
                        if (!(0.1 != (-2.25 / float32(-2.25)))) {
                            v126 := []int{(fi125 + fi125), (fi122 * 2147483647), (^(v5)), fi125}
                            v1 -= 255
                            var v127 int16 = 1000
                            fmt.Println(v1, v1)
                            if true {
                                v3 = (^(v3))
                                var v128 uint = (v3 & (v3 - 255))
                                if (42 > 127) {
                                    var v129 int = fi122
                                    var v130 int16 = 32767
                                    v1 = v1
                                    _ = v129
                                    _ = v130
                                }
                                if true {
                                    v131 := GenS9{F0: v1, F1: 2, F2: 717}
                                    _ = v131
                                }
                                _ = v128
                            } else {
                                var v132 float32 = (0.1 * (1e10 / 1.5))
                                v133 := GenS17{F0: v127, F1: (v127 + 1)}
                                v3 = genFn0()
                                _ = v132
                                _ = v133
                            }
                            _ = v126
                            _ = v127
                        }
                        var v134 uint8 = 42
                        var v135 uint32 = 2147483647
                        _ = v134
                        _ = v135
                    }
                }
            }
            _ = fi125
        }
        for fi136 := 0; fi136 < 12; fi136++ {
            v137 := GenS18{F0: 1, F1: 0, F2: (1 < 42)}
            switch fi136 {
            case -770:
                var v138 float32 = ((0.0 + 3.14159) * ((((1.5 * -2.25) * 0.1) - 0.1) * 0.0))
                _ = v138
            }
            var v139 int64 = 42
            _ = fi136
            _ = v137
            _ = v139
        }
        _ = fi122
    }
    v1 += v1
    delete(v2, 2147483647)
    v3 += v3
    switch v5 {
    case -947:
        for fi140 := 0; fi140 < 2; fi140++ {
            fmt.Println(fi140)
            _ = fi140
        }
    case 0:
        switch "\n\t" {
        case "":
            for fi141 := 0; fi141 < 8; fi141++ {
                var v142 int16 = 0
                fmt.Println(fi141)
                v1 -= v1
                _ = fi141
                _ = v142
            }
        case "hello":
            v1 += 0
        default:
            v5 = (-(v5))
        }
    case -379:
        for fi143 := 0; fi143 < 5; fi143++ {
            if (42 != 0) {
                var v144 int64 = -362
                var v145 int = (fi143 - fi143)
                _ = v144
                _ = v145
            } else {
                v146 := v2[800]
                v3 -= (v3 / ((genFn0()) | 1))
                var v147 float32 = 1e10
                _ = v146
                _ = v147
            }
            _ = fi143
        }
    }
    if (int(v3) == v5) {
        v148 := []int{v5, v5, (v5 * (v5 - v5))}
        fmt.Println(v3, v5, v3)
        fmt.Println(v5, v5, v5)
        _ = v148
    } else {
        var v149 bool = (v5 > 1)
        fmt.Println(v3)
        _ = v149
    }
    for fi150 := 0; fi150 < 10; fi150++ {
        v151 := GenS19{F0: 42, F1: 764}
        fmt.Println(v5, v1, v5)
        v1 = 65535
        _ = fi150
        _ = v151
    }
    var v152 uint = genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v152:", v152)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v152
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
}
