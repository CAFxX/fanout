package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 uint64; F1 bool; F2 int64 }
type GenS11 struct { F0 int64; F1 bool; F2 bool }
type GenS12 struct { F0 int32 }
func genFn0() uint32 {
    var v6 uint64 = 4294967295
    if false {
        v6 -= (^(v6))
    } else {
        for fi7 := 0; fi7 < 5; fi7++ {
            switch "hello" {
            case "\n\t":
                v8 := GenS1{F0: 255}
                _ = v8
            case "a":
                v9 := GenS1{F0: 65535}
                _ = v9
            default:
                for fi10 := 0; fi10 < 6; fi10++ {
                    v11 := []int{fi10}
                    v6 = 2
                    _ = fi10
                    _ = v11
                }
            }
            var v12 float32 = ((float32(3.14159) + float32(-2.25)) * ((3.14159 * 1e10) - ((0.0 - -2.25) * 1.5)))
            _ = fi7
            _ = v12
        }
        var v13 uint8 = 128
        v13 = 255
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := []int{(fi14 - (fi14 ^ -951)), (fi14 - (fi14 / ((fi14) | 1))), (fi14 + 689)}
        fmt.Println(fi14, fi14)
        _ = fi14
        _ = v15
    }
    switch -2147483648 {
    case 127:
        v6 = v6
    }
    v6 -= (v6 & (v6 | v6))
    _ = v6
    return 0
}
func genFn1() int8 {
    var v16 int32 = 42
    switch "hello" {
    case "hello":
        var v17 int = 450
        _ = v17
    case "\n\t":
        if false {
            var v18 int = 128
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := GenS2{F0: (42 <= 65535)}
            v16 = (v16 | v16)
            if true {
                var v21 int16 = 120
                if ("foo bar" <= "foo bar") {
                    fmt.Println(v21, v16, v16)
                    v16 += (v16 + -2147483648)
                }
                if false {
                    var v22 string = "\n\t"
                    var v23 uint32 = 1
                    _ = v22
                    _ = v23
                } else {
                    v24 := GenS3{F0: -9223372036854775808, F1: 32767}
                    for fi25 := 0; fi25 < 18; fi25++ {
                        v26 := []int{fi25, fi19}
                        _ = fi25
                        _ = v26
                    }
                    _ = v24
                }
                _ = v21
            }
            _ = fi19
            _ = v20
        }
    }
    _ = v16
    return 105
}

func main() {
    var v1 bool = false
    var v2 uint = 4294967295
    v3 := make(map[string]int)
    var v4 float32 = -2.25
    v5 := make(map[uint32]int)
    var v27 int64 = 951
    v2 -= v2
    v28 := []int{1, -2147483648, -1}
    v2 = v2
    if v1 {
        v29 := GenS1{F0: v2}
        for fi30 := 0; fi30 < 19; fi30++ {
            var v31 uint64 = 1
            var v32 float64 = -2.25
            var v33 int64 = -9223372036854775808
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
        for fi34 := 0; fi34 < 20; fi34++ {
            fmt.Println(fi34, v4, fi34)
            _ = fi34
        }
        _ = v29
    }
    if true {
        for fi35 := 0; fi35 < 15; fi35++ {
            var v36 uint64 = 516
            v28 = append(v28, fi35)
            _ = fi35
            _ = v36
        }
        v37 := GenS3{F0: (v27 * -79), F1: 149}
        _ = v37
    } else {
        v38 := GenS2{F0: (243 == 0)}
        v27 = (v27 - v27)
        switch -2147483648 {
        case 127:
            if ((v27 % (((v27 >> (((^(v2))) & 63))) | 1)) > (v27 | v27)) {
                for fi39 := 0; fi39 < 4; fi39++ {
                    fmt.Println(fi39, v27, v27)
                    if false {
                        v40 := GenS4{F0: 0, F1: ("a" + "a")}
                        v1 = true
                        for fi41 := 0; fi41 < 15; fi41++ {
                            var v42 int16 = 32767
                            v2 = v2
                            if (!(v42 <= v42)) {
                                v27 = -9223372036854775808
                            } else {
                                var v43 uint64 = 42
                                var v44 uint32 = genFn0()
                                v45 := GenS5{F0: 31}
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                            v27 = v27
                            _ = fi41
                            _ = v42
                        }
                        _ = v40
                    } else {
                        var v46 int32 = -2147483648
                        var v47 int8 = -1
                        _ = v46
                        _ = v47
                    }
                    v48 := GenS6{F0: 0, F1: -806}
                    _ = fi39
                    _ = v48
                }
                if true {
                    if true {
                        v4 -= ((1.5 - float32(1e10)) * 1e10)
                    }
                    for fi49 := 0; fi49 < 15; fi49++ {
                        if (-396 < 0) {
                            v2 = 18446744073709551615
                            if true {
                                fmt.Println(v2, v27, v27)
                                for fi50 := 0; fi50 < 10; fi50++ {
                                    fmt.Println(v4, v2)
                                    if v1 {
                                        v4 = float32((0.0 - 0.1))
                                        var v51 int16 = 42
                                        switch fi49 {
                                        case 640:
                                            _ = len(v28)
                                        case -2147483648:
                                            var v52 float32 = 3.14159
                                            _ = v52
                                        }
                                        _ = v51
                                    } else {
                                        v53 := GenS5{F0: -96}
                                        _ = v53
                                    }
                                    _ = fi50
                                }
                                if false {
                                    v54 := GenS7{F0: 42}
                                    _ = v54
                                } else {
                                    switch "" {
                                    case "a":
                                        v4 = v4
                                    case "\n\t":
                                        for fi55 := 0; fi55 < 9; fi55++ {
                                            v56 := GenS7{F0: -358}
                                            if len(v28) >= 2 {
                                                v28 = v28[1:len(v28)]
                                            }
                                            _ = fi55
                                            _ = v56
                                        }
                                    case "foo bar":
                                        var v57 int = 127
                                        _ = v57
                                    }
                                    var v58 int8 = 127
                                    var v59 uint = v2
                                    _ = v58
                                    _ = v59
                                }
                            } else {
                                for fi60 := 0; fi60 < 20; fi60++ {
                                    var v61 uint = v2
                                    var v62 float64 = float64(v4)
                                    v27 = (v27 >> ((v61) & 63))
                                    _ = fi60
                                    _ = v61
                                    _ = v62
                                }
                            }
                        }
                        var v63 uint8 = 0
                        _ = fi49
                        _ = v63
                    }
                } else {
                    if (2147483647 != 0) {
                        for fi64 := 0; fi64 < 17; fi64++ {
                            fmt.Println(v2, v1)
                            fmt.Println(v4)
                            _ = fi64
                        }
                    }
                    v4 += -2.25
                }
                if (v2 > (v2 * (v2 | 1))) {
                    var v65 int8 = int8(v2)
                    _ = v65
                }
            } else {
                _ = len(v5)
            }
        }
        _ = v38
    }
    switch -1 {
    case 32768:
        for fi66 := 0; fi66 < 17; fi66++ {
            if v1 {
                switch "hello" {
                case "":
                    var v67 bool = ((false || ((0.1 - 0.1) < -2.25)) && v1)
                    _ = v67
                case "foo bar":
                    switch fi66 {
                    case -998:
                        switch fi66 {
                        case -130:
                            if v1 {
                                var v68 uint32 = 1
                                if (!(true && (true || false))) {
                                    for fi69 := 0; fi69 < 12; fi69++ {
                                        v4 += (v4 - v4)
                                        v68 = (^(v68))
                                        var v70 uint32 = 2147483647
                                        v4 -= v4
                                        fmt.Println(fi69, fi69, v27)
                                        if (false || v1) {
                                            if (v1 && ((("a" + "\n\t") + "hello") == "")) {
                                                v1 = v1
                                                v70 = (v68 % ((v68) | 1))
                                                var v71 float32 = ((3.14159 - float32(0.0)) - (3.14159 - (v4 / 3.14159)))
                                                _ = v71
                                            }
                                        }
                                        var v72 bool = (-172 <= int16(v27))
                                        fmt.Println(v68, v2, v70)
                                        switch 42 {
                                        case 2147483647:
                                            if false {
                                                if true {
                                                    _ = len(v5)
                                                    v73 := GenS6{F0: 127, F1: 304}
                                                    var v74 int8 = 127
                                                    _ = v73
                                                    _ = v74
                                                } else {
                                                    if true {
                                                        var v75 int8 = 117
                                                        _ = v75
                                                    } else {
                                                        if false {
                                                            v3["\n\t"] = (fi66 >> (((v2 * v2)) & 63))
                                                            var v76 float64 = (((1.5 - (((1.5 + ((((1.5 + 0.1) - float64(1.5)) + (float64(v4) / 1.0)) - (0.0 / 1.5))) / ((0.1 * 0.0) + 0.1)) - (-2.25 - 1e10))) / (-2.25 / 1.0)) / 1.0)
                                                            var v77 int8 = genFn1()
                                                            v68 += v70
                                                            _ = v76
                                                            _ = v77
                                                        } else {
                                                            var v78 uint64 = 0
                                                            v28 = append(v28, (fi69 + fi69))
                                                            _ = v78
                                                        }
                                                        var v79 uint32 = (v68 | 767)
                                                        var v80 uint32 = genFn0()
                                                        _ = v79
                                                        _ = v80
                                                    }
                                                    if v72 {
                                                        var v81 uint8 = 93
                                                        if true {
                                                            var v82 string = "\n\t"
                                                            if len(v28) >= 2 {
                                                                v28 = v28[1:len(v28)]
                                                            }
                                                            _ = v82
                                                        } else {
                                                            fmt.Println(v27, v1, v2)
                                                        }
                                                        if true {
                                                            if len(v28) >= 2 {
                                                                v28 = v28[1:len(v28)]
                                                            }
                                                            if (!v72) {
                                                                v68 += genFn0()
                                                                fmt.Println(v1)
                                                            } else {
                                                                v83 := GenS8{F0: 9223372036854775807}
                                                                if len(v28) > 0 {
                                                                    v28[0] = (fi69 % (((fi69 - (fi69 - fi66))) | 1))
                                                                }
                                                                v84 := GenS9{F0: 978, F1: 0}
                                                                var v85 int64 = -1
                                                                v86 := GenS6{F0: 100, F1: 42}
                                                                if (1 == (v2 << ((42) & 63))) {
                                                                    if v72 {
                                                                        v70 += v68
                                                                        var v87 bool = (0.1 > 1.5)
                                                                        _ = v87
                                                                    }
                                                                } else {
                                                                    var v88 uint16 = 622
                                                                    _ = v88
                                                                }
                                                                var v89 int8 = int8(v70)
                                                                var v90 uint32 = (v68 / (((^(v70))) | 1))
                                                                _ = v83
                                                                _ = v84
                                                                _ = v85
                                                                _ = v86
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                        }
                                                        _ = v81
                                                    } else {
                                                        v68 -= (v68 << (((v2 % (((v2 >> (((v2 << (((v2 - v2)) & 63))) & 63))) | 1))) & 63))
                                                        switch "foo bar" {
                                                        case "a":
                                                            v28 = append(v28, (fi69 | 42))
                                                        case "foo bar":
                                                            if true {
                                                                fmt.Println(v2, v2)
                                                                var v91 uint64 = 412
                                                                var v92 uint32 = genFn0()
                                                                _ = v91
                                                                _ = v92
                                                            } else {
                                                                v4 += v4
                                                                var v93 bool = v1
                                                                _ = v93
                                                            }
                                                        case "":
                                                            _ = len(v28)
                                                            v72 = (382 > 861)
                                                            var v94 int64 = (v27 % ((v27) | 1))
                                                            _ = v94
                                                        default:
                                                            var v95 string = ("a" + "")
                                                            _ = v95
                                                        }
                                                    }
                                                    v4 = (v4 * v4)
                                                }
                                                v2 = v2
                                                var v96 uint64 = 42
                                                _ = v96
                                            } else {
                                                switch "foo bar" {
                                                case "a":
                                                    if true {
                                                        if false {
                                                            var v97 string = "hello"
                                                            _ = len(v28)
                                                            v98 := GenS5{F0: -128}
                                                            var v99 float32 = (0.1 - ((float32((1e10 / ((((3.14159 * float64(v4)) - 3.14159) * float64(float32(float64(float32(3.14159))))) + (float64(v4) - -2.25)))) - 0.1) + v4))
                                                            _ = v97
                                                            _ = v98
                                                            _ = v99
                                                        } else {
                                                            var v100 int16 = 216
                                                            _ = v100
                                                        }
                                                        var v101 uint = (v2 * v2)
                                                        if true {
                                                            v2 = (^(v101))
                                                            v102 := GenS1{F0: v101}
                                                            v103 := GenS10{F0: 42, F1: v72, F2: (^(v27))}
                                                            _ = v102
                                                            _ = v103
                                                        } else {
                                                            v1 = v1
                                                        }
                                                        _ = v101
                                                    }
                                                case "foo bar":
                                                    if v1 {
                                                        var v104 float64 = -2.25
                                                        var v105 string = (("a" + "a") + (("" + "a") + ((("foo bar" + "hello") + ("hello" + "")) + ((("\n\t" + "foo bar") + "\n\t") + (("foo bar" + "hello") + "foo bar")))))
                                                        _ = v104
                                                        _ = v105
                                                    }
                                                default:
                                                    v106 := GenS2{F0: v72}
                                                    _ = v106
                                                }
                                                _ = len(v28)
                                            }
                                        case 464:
                                            v107 := v5[(v70 - v68)]
                                            _ = v107
                                        case 128:
                                            _ = len(v5)
                                        }
                                        _ = fi69
                                        _ = v70
                                        _ = v72
                                    }
                                    for fi108 := 0; fi108 < 12; fi108++ {
                                        var v109 int = (fi108 ^ int(v27))
                                        if (235 < 49) {
                                            v110 := v3["hello"]
                                            _ = v110
                                        } else {
                                            if len(v28) >= 2 {
                                                v28 = v28[1:len(v28)]
                                            }
                                            v1 = true
                                        }
                                        var v111 uint = (v2 - v2)
                                        if (224 > 42) {
                                            v68 = v68
                                            if v1 {
                                                v4 = float32(0.0)
                                            } else {
                                                fmt.Println(v111, v4, v68)
                                                v111 += 18446744073709551615
                                                var v112 bool = (v68 >= (v68 >> ((v2) & 63)))
                                                _ = v112
                                            }
                                            if true {
                                                var v113 int8 = 119
                                                v27 += v27
                                                _ = v113
                                            }
                                        } else {
                                            _ = len(v28)
                                        }
                                        var v114 uint16 = uint16(v68)
                                        if (v2 != v111) {
                                            if ((v68 / ((v68) | 1)) <= (^(v68))) {
                                                var v115 bool = (0 >= 0)
                                                if v115 {
                                                    v5[v68] = fi108
                                                }
                                                v116 := GenS5{F0: 0}
                                                var v117 uint16 = v114
                                                switch 42 {
                                                case 0:
                                                    v68 = (^(v68))
                                                case -2147483648:
                                                    v109 += (-(fi66))
                                                    v114 += 2
                                                    var v118 int16 = -1
                                                    _ = v118
                                                case -9223372036854775808:
                                                    if len(v28) > 0 {
                                                        v28[0] = fi108
                                                    }
                                                }
                                                v119 := GenS11{F0: v27, F1: (1 <= -32768), F2: v1}
                                                _ = v115
                                                _ = v116
                                                _ = v117
                                                _ = v119
                                            }
                                        } else {
                                            if len(v28) >= 2 {
                                                v28 = v28[1:len(v28)]
                                            }
                                            if len(v28) > 0 {
                                                v28[0] = (^(v109))
                                            }
                                        }
                                        _ = fi108
                                        _ = v109
                                        _ = v111
                                        _ = v114
                                    }
                                }
                                _ = v68
                            } else {
                                v4 -= (1.5 * v4)
                            }
                        }
                    case 127:
                        if len(v28) > 0 {
                            v28[0] = fi66
                        }
                    case -2147483648:
                        for fi120 := 0; fi120 < 15; fi120++ {
                            v121 := GenS12{F0: -586}
                            _ = fi120
                            _ = v121
                        }
                    }
                case "a":
                    for fi122 := 0; fi122 < 3; fi122++ {
                        if (genFn0() == 354) {
                            switch fi66 {
                            case -1:
                                fmt.Println(v1, v4)
                            }
                            var v123 int8 = -128
                            if ((((!v1) && (0 < 227)) && false) || (v1 && (255 > 255))) {
                                var v124 uint = v2
                                var v125 uint64 = 142
                                var v126 uint32 = 262
                                delete(v3, "a")
                                _ = v124
                                _ = v125
                                _ = v126
                            }
                            _ = v123
                        } else {
                            var v127 float64 = 0.1
                            v127 += v127
                            fmt.Println(fi122, v27, v1)
                            _ = v127
                        }
                        _ = fi122
                    }
                }
                var v128 uint16 = 42
                var v129 float32 = (((float32(-2.25) / v4) / (v4 - -2.25)) / v4)
                _ = v128
                _ = v129
            }
            var v130 int8 = -128
            v4 = (((v4 / ((v4 * float32(3.14159)) * (-2.25 * v4))) - 0.1) + (v4 + v4))
            _ = fi66
            _ = v130
        }
    case -9223372036854775808:
        if ((!(v2 >= v2)) && v1) {
            v131 := GenS7{F0: 83}
            _ = v131
        }
    case 2147483647:
        fmt.Println(v2, v4)
    }
    var v132 float32 = (((v4 + ((((v4 - v4) + -2.25) / ((v4 - v4) + 3.14159)) - (v4 / v4))) + 1.5) * v4)
    fmt.Println(v4)
    var v133 uint32 = 65535
    _ = len(v5)
    if len(v28) > 0 {
        v28[0] = 42
    }
    for fi134 := 0; fi134 < 3; fi134++ {
        v135 := GenS1{F0: (v2 << ((v2) & 63))}
        _ = len(v5)
        _ = fi134
        _ = v135
    }
    v136 := GenS2{F0: false}
    v28 = append(v28, int(v133))
    _ = len(v28)
    v28 = append(v28, 715)
    fmt.Println("v1:", v1)
    fmt.Println("v132:", v132)
    fmt.Println("v133:", v133)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v132
    _ = v133
    _ = v136
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
