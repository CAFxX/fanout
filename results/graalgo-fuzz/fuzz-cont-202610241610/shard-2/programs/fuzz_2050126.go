package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 bool; F1 int8; F2 string }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int8; F1 int16 }
type GenS10 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS11 struct { F0 int64; F1 string }
func genFn0() string {
    for fi4 := 0; fi4 < 10; fi4++ {
        for fi5 := 0; fi5 < 16; fi5++ {
            switch (("hello" + ((("foo bar" + "foo bar") + "hello") + "")) + "\n\t") {
            case "\n\t":
                var v6 float64 = ((float64(0.0) / 1e10) + (0.1 + 1.5))
                _ = v6
            case "foo bar":
                v7 := GenS1{F0: 42}
                _ = v7
            default:
                v8 := []int{42, (fi5 ^ (fi5 >> ((2) & 63))), fi5}
                _ = v8
            }
            if true {
                switch 4294967295 {
                case 202:
                    v9 := GenS1{F0: 32}
                    _ = v9
                case -20:
                    var v10 float32 = 1.5
                    _ = v10
                }
                v11 := GenS2{F0: 95}
                _ = v11
            }
            var v12 int8 = -31
            var v13 bool = false
            _ = fi5
            _ = v12
            _ = v13
        }
        v14 := []int{(^(fi4))}
        if true {
            if true {
                v14 = append(v14, fi4)
                if len(v14) > 0 {
                    v14[0] = 42
                }
            }
            v15 := GenS2{F0: 74}
            _ = v15
        } else {
            var v16 int8 = 107
            fmt.Println(v16, v16)
            _ = v16
        }
        _ = fi4
        _ = v14
    }
    if true {
        var v17 int8 = -128
        v18 := []int{32767}
        fmt.Println(v17)
        _ = v17
        _ = v18
    } else {
        var v19 uint = 18446744073709551615
        switch ("" + "a") {
        case "hello":
            v19 -= 1
        default:
            var v20 int16 = -32768
            _ = v20
        }
        if true {
            v21 := []int{32, 127, 127, 127}
            if (!(true || true)) {
                var v22 int64 = -1
                _ = v22
            }
            for fi23 := 0; fi23 < 2; fi23++ {
                v19 += 42
                _ = fi23
            }
            _ = v21
        }
        _ = v19
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        v25 := []int{fi24}
        _ = fi24
        _ = v25
    }
    if (4294967295 == -2147483648) {
        var v26 bool = (false || (1 <= 766))
        _ = v26
    } else {
        v27 := []int{86, 127, 9223372036854775807, -1}
        _ = v27
    }
    return ""
}

func main() {
    var v1 uint = 1
    var v2 uint8 = 172
    v3 := make(map[uint32]uint32)
    if false {
        for fi28 := 0; fi28 < 3; fi28++ {
            if true {
                if (237 > 4294967295) {
                    if false {
                        v1 = (v1 ^ uint(v2))
                    } else {
                        v1 -= v1
                    }
                    for fi29 := 0; fi29 < 18; fi29++ {
                        var v30 int32 = 42
                        fmt.Println(v30, fi28)
                        if (("foo bar" + (("\n\t" + "hello") + "a")) != "\n\t") {
                            v30 = v30
                            fmt.Println(fi28, v1, v30)
                            v3[1] = 0
                            v31 := GenS1{F0: 1}
                            var v32 uint = (v1 >> ((2) & 63))
                            fmt.Println(v32, fi29, v30)
                            v3[766] = 0
                            fmt.Println(v32)
                            _ = v31
                            _ = v32
                        }
                        _ = fi29
                        _ = v30
                    }
                    for fi33 := 0; fi33 < 19; fi33++ {
                        var v34 bool = (!true)
                        var v35 int = fi28
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                } else {
                    for fi36 := 0; fi36 < 17; fi36++ {
                        fmt.Println(fi36, v1)
                        switch ("" + (("" + "") + "\n\t")) {
                        case "a":
                            var v37 uint = 643
                            _ = v37
                        default:
                            var v38 int32 = 260
                            _ = v38
                        }
                        _ = fi36
                    }
                    switch 4294967295 {
                    case -1:
                        v39 := []int{(-(fi28)), (fi28 ^ fi28)}
                        _ = v39
                    case -9223372036854775808:
                        fmt.Println(v1, v2, v2)
                    default:
                        if (32767 != 0) {
                            for fi40 := 0; fi40 < 3; fi40++ {
                                v2 += v2
                                var v41 uint8 = v2
                                var v42 uint = v1
                                v43 := GenS2{F0: 0}
                                _ = fi40
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            for fi44 := 0; fi44 < 3; fi44++ {
                                switch ("a" + "") {
                                case "a":
                                    var v45 uint = v1
                                    _ = v45
                                default:
                                    if (false && (!(!true))) {
                                        if (32767 > 32767) {
                                            v2 -= (v2 & v2)
                                            v2 -= 255
                                            v46 := GenS2{F0: 117}
                                            var v47 uint = 42
                                            _ = v46
                                            _ = v47
                                        }
                                        v2 += 0
                                    } else {
                                        fmt.Println(v2, fi28)
                                        v48 := []int{1}
                                        _ = v48
                                    }
                                }
                                if false {
                                    var v49 uint64 = 0
                                    _ = v49
                                }
                                _ = fi44
                            }
                            var v50 int16 = -32768
                            _ = v50
                        } else {
                            for fi51 := 0; fi51 < 12; fi51++ {
                                var v52 float32 = 1e10
                                v53 := []int{}
                                _ = fi51
                                _ = v52
                                _ = v53
                            }
                            v54 := GenS3{F0: v2}
                            for fi55 := 0; fi55 < 12; fi55++ {
                                fmt.Println(v2, v2)
                                v1 -= (^(v1))
                                v2 -= (v2 % (((v2 % (((v2 - 0)) | 1))) | 1))
                                if true {
                                    var v56 uint16 = 42
                                    _ = v56
                                } else {
                                    var v57 uint16 = 1
                                    _ = v57
                                }
                                var v58 int16 = -630
                                _ = fi55
                                _ = v58
                            }
                            _ = v54
                        }
                    }
                    fmt.Println(fi28, fi28, v2)
                }
                v59 := GenS4{F0: "", F1: 632}
                v2 = 127
                _ = v59
            } else {
                v1 = v1
                var v60 int32 = -2147483648
                fmt.Println(fi28, fi28, v1)
                _ = v60
            }
            var v61 int32 = -1
            switch "foo bar" {
            case "hello":
                v61 -= v61
            case "\n\t":
                if true {
                    v62 := []int{}
                    _ = v62
                } else {
                    v1 = v1
                }
            }
            _ = fi28
            _ = v61
        }
        switch "hello" {
        case "hello":
            if (!true) {
                var v63 uint16 = 1
                if true {
                    var v64 int32 = 0
                    if (466 != 9223372036854775807) {
                        v65 := []int{9223372036854775807}
                        _ = v65
                    }
                    var v66 float32 = 3.14159
                    _ = v64
                    _ = v66
                } else {
                    for fi67 := 0; fi67 < 3; fi67++ {
                        v2 += (v2 ^ v2)
                        switch ("hello" + "") {
                        case "":
                            fmt.Println(v2, v2)
                        case "foo bar":
                            var v68 float32 = ((3.14159 / 1.0) * (0.1 - 1.5))
                            _ = v68
                        case "a":
                            v69 := []int{32767}
                            _ = v69
                        }
                        v70 := v3[4294967295]
                        _ = fi67
                        _ = v70
                    }
                    switch ("" + "foo bar") {
                    case "a":
                        v63 = (v63 & v63)
                    default:
                        v71 := GenS2{F0: 0}
                        _ = v71
                    }
                }
                _ = v63
            } else {
                var v72 float32 = 1.5
                switch "hello" {
                case "hello":
                    var v73 int64 = -9223372036854775808
                    _ = v73
                }
                switch "hello" {
                case "foo bar":
                    _ = len(v3)
                case "hello":
                    v2 = (v2 & v2)
                case "a":
                    v2 = v2
                }
                _ = v72
            }
        case "foo bar":
            if (1 != -107) {
                if (true && true) {
                    v74 := []int{4294967295, 571, -719, 128}
                    _ = v74
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        if false {
                            if false {
                                var v75 bool = (!(!true))
                                _ = v75
                            } else {
                                var v76 int16 = 1
                                if (42 != 42) {
                                    v77 := []int{-649, -9223372036854775808, 4294967295, 506}
                                    var v78 int16 = (v76 << (((v1 & 18446744073709551615)) & 63))
                                    fmt.Println(v78, v78)
                                    _ = v77
                                    _ = v78
                                }
                                _ = v76
                            }
                            fmt.Println(v2, v1)
                        } else {
                            fmt.Println(v1)
                            switch "foo bar" {
                            case "\n\t":
                                var v79 uint32 = 4294967295
                                _ = v79
                            case "a":
                                delete(v3, 65535)
                            default:
                                fmt.Println(v1, v1)
                            }
                            v80 := GenS1{F0: 255}
                            _ = v80
                        }
                    default:
                        v3[2] = 2147483647
                    }
                }
                var v81 uint8 = 233
                _ = v81
            }
        case "\n\t":
            switch -9223372036854775808 {
            case 42:
                var v82 int = 127
                _ = v82
            default:
                switch ("hello" + "foo bar") {
                case "":
                    var v83 string = ("hello" + "\n\t")
                    _ = v83
                case "a":
                    v84 := GenS3{F0: (v2 << ((v1) & 63))}
                    _ = v84
                case "\n\t":
                    var v85 float64 = 1.5
                    _ = v85
                }
            }
        }
        if (1 > 2147483647) {
            v86 := []int{-2147483648, 32768, 1, 4294967295}
            v87 := GenS2{F0: -38}
            if (!false) {
                for fi88 := 0; fi88 < 14; fi88++ {
                    var v89 uint32 = 2147483647
                    _ = fi88
                    _ = v89
                }
                var v90 float64 = 0.0
                _ = v90
            }
            _ = v86
            _ = v87
        } else {
            if (v2 == (v2 * v2)) {
                delete(v3, 62)
                v3[4294967295] = 199
            } else {
                v2 = v2
            }
            v1 -= uint(v2)
            if true {
                var v91 string = ("foo bar" + "")
                var v92 int = 128
                v92 = v92
                _ = v91
                _ = v92
            }
        }
    }
    var v93 uint = (v1 << ((255) & 63))
    for fi94 := 0; fi94 < 4; fi94++ {
        var v95 uint32 = 740
        v2 -= v2
        var v96 uint = (v93 | (v93 >> ((v1) & 63)))
        _ = fi94
        _ = v95
        _ = v96
    }
    delete(v3, 255)
    v2 = 52
    switch ("hello" + "hello") {
    case "a":
        v97 := GenS2{F0: -1}
        _ = v97
    default:
        if ((!true) || true) {
            fmt.Println(v1, v1, v93)
            v98 := GenS3{F0: (^(v2))}
            var v99 uint16 = 548
            _ = v98
            _ = v99
        } else {
            var v100 uint8 = v2
            fmt.Println(v100)
            _ = v100
        }
    }
    if false {
        v2 = (v2 + (v2 ^ v2))
    } else {
        if ((v1 << ((v1) & 63)) <= v93) {
            v101 := v3[2147483648]
            v101 += (v101 - 927)
            switch "\n\t" {
            case "\n\t":
                var v102 uint16 = uint16(v2)
                _ = v102
            case "":
                for fi103 := 0; fi103 < 13; fi103++ {
                    var v104 bool = false
                    v105 := GenS3{F0: v2}
                    if (false && (v104 || (v104 && false))) {
                        var v106 float32 = ((float32(3.14159) * 1.5) / (1.5 - -2.25))
                        var v107 int64 = 0
                        _ = v106
                        _ = v107
                    } else {
                        v108 := GenS5{F0: (42 == 0), F1: 42, F2: "a"}
                        v109 := []int{fi103, (fi103 ^ (^(fi103)))}
                        if (127 == 16) {
                            for fi110 := 0; fi110 < 18; fi110++ {
                                var v111 uint32 = 0
                                if true {
                                    fmt.Println(v111, v111)
                                } else {
                                    var v112 float64 = ((1.5 + 1.5) / (3.14159 / float64(0.1)))
                                    if true {
                                        if false {
                                            if v104 {
                                                v113 := GenS6{F0: v104}
                                                var v114 uint = (v93 >> ((42) & 63))
                                                v101 = v111
                                                _ = v113
                                                _ = v114
                                            }
                                        } else {
                                            v112 = 0.1
                                        }
                                        v2 = (v2 << (((v93 & (^(v93)))) & 63))
                                        v3[(v101 + v101)] = v101
                                    } else {
                                        switch 42 {
                                        case 2147483647:
                                            var v115 uint8 = v2
                                            _ = v115
                                        case 9223372036854775807:
                                            v109 = append(v109, (-(fi103)))
                                        case -271:
                                            v116 := GenS7{F0: -1, F1: -1}
                                            _ = v116
                                        }
                                        if false {
                                            var v117 int8 = -71
                                            _ = len(v109)
                                            if ((!(0 > 4294967295)) && v104) {
                                                _ = len(v3)
                                                v93 = 42
                                            } else {
                                                var v118 float64 = v112
                                                var v119 uint32 = v111
                                                v120 := GenS4{F0: ("\n\t" + ("hello" + "hello")), F1: (v111 + (v101 % ((0) | 1)))}
                                                _ = v118
                                                _ = v119
                                                _ = v120
                                            }
                                            _ = v117
                                        } else {
                                            v112 += 0.1
                                        }
                                    }
                                    v104 = false
                                    _ = v112
                                }
                                _ = fi110
                                _ = v111
                            }
                            for fi121 := 0; fi121 < 3; fi121++ {
                                if v104 {
                                    v122 := GenS8{F0: -1}
                                    if (!(2147483647 <= -1)) {
                                        fmt.Println(v2)
                                        fmt.Println(fi121)
                                        v101 += v101
                                    } else {
                                        var v123 int = (fi103 % ((fi103) | 1))
                                        _ = v123
                                    }
                                    _ = v122
                                }
                                var v124 bool = (v104 || (!v104))
                                fmt.Println(v93, v104)
                                _ = fi121
                                _ = v124
                            }
                        } else {
                            for fi125 := 0; fi125 < 18; fi125++ {
                                v126 := GenS9{F0: -2, F1: -908}
                                var v127 uint32 = v101
                                v93 += (v93 / ((v1) | 1))
                                _ = fi125
                                _ = v126
                                _ = v127
                            }
                            v128 := GenS10{F0: 763, F1: 137, F2: 2147483647}
                            if v104 {
                                v129 := GenS2{F0: -128}
                                for fi130 := 0; fi130 < 3; fi130++ {
                                    v2 -= (^(v2))
                                    v3[v101] = (^(v101))
                                    _ = fi130
                                }
                                if (v104 || v104) {
                                    v2 = (^(v2))
                                    fmt.Println(v104, v93)
                                    if (v2 <= 255) {
                                        _ = len(v109)
                                        var v131 string = ("a" + "foo bar")
                                        v93 = v1
                                        _ = v131
                                    } else {
                                        for fi132 := 0; fi132 < 10; fi132++ {
                                            v101 -= v101
                                            fmt.Println(v101)
                                            var v133 uint8 = (v2 >> (((v1 & 811)) & 63))
                                            _ = fi132
                                            _ = v133
                                        }
                                        v2 = (v2 << ((v1) & 63))
                                        if v104 {
                                            _ = len(v3)
                                            var v134 bool = (42 < -9223372036854775808)
                                            _ = v134
                                        } else {
                                            var v135 int64 = 0
                                            v1 -= (v1 % (((v93 ^ 0)) | 1))
                                            _ = v135
                                        }
                                    }
                                } else {
                                    var v136 float64 = 0.1
                                    _ = v136
                                }
                                _ = v129
                            } else {
                                var v137 uint64 = 2
                                if (!((0.1 + (float32((3.14159 - 1.5)) - (1e10 - 3.14159))) <= float32(1.5))) {
                                    v138 := v3[0]
                                    var v139 uint = v1
                                    _ = v138
                                    _ = v139
                                }
                                _ = v137
                            }
                            _ = v128
                        }
                        _ = v108
                        _ = v109
                    }
                    _ = fi103
                    _ = v104
                    _ = v105
                }
            case "hello":
                fmt.Println(v93, v93, v1)
            }
            _ = v101
        }
        if false {
            if (!(!(v93 > (v1 / (((^(v1))) | 1))))) {
                v140 := GenS5{F0: (false || (1 < 42)), F1: -128, F2: "foo bar"}
                _ = v140
            }
            fmt.Println(v2)
            v141 := GenS9{F0: 100, F1: 1}
            _ = v141
        } else {
            v93 = v93
            v1 = v93
        }
        var v142 uint8 = (v2 - v2)
        _ = v142
    }
    var v143 uint16 = 568
    for fi144 := 0; fi144 < 8; fi144++ {
        v93 = uint(v2)
        fmt.Println(fi144, fi144)
        v145 := GenS6{F0: false}
        _ = fi144
        _ = v145
    }
    v146 := GenS2{F0: 100}
    var v147 uint32 = 1
    if false {
        _ = len(v3)
    } else {
        v148 := v3[uint32(v2)]
        _ = v148
    }
    v93 = (v93 + (v1 | (v93 / ((v93) | 1))))
    v149 := GenS11{F0: 0, F1: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v143:", v143)
    fmt.Println("v147:", v147)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v93:", v93)
    _ = v1
    _ = v143
    _ = v146
    _ = v147
    _ = v149
    _ = v2
    _ = v3
    _ = v93
}
