package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 int16; F2 int16 }
type GenS7 struct { F0 bool; F1 uint8 }
type GenS8 struct { F0 uint8; F1 int16; F2 uint }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint64; F1 int16 }
type GenS11 struct { F0 uint; F1 uint32; F2 int16 }
type GenS12 struct { F0 uint8 }
func genFn0(p0 int32, p1 int16) int32 {
    fmt.Println(p1)
    var v5 int8 = 7
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v5
    return p0
}
func genFn1() int64 {
    var v6 bool = (2 != 65535)
    v7 := []int{1, 1, -2147483648, 1}
    for fi8 := 0; fi8 < 5; fi8++ {
        if v6 {
            var v9 int16 = 32767
            _ = v9
        }
        v6 = v6
        for fi10 := 0; fi10 < 1; fi10++ {
            if len(v7) > 0 {
                v7[0] = (fi10 % ((637) | 1))
            }
            _ = fi10
        }
        _ = fi8
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        if v6 {
            _ = len(v7)
            v12 := GenS1{F0: 718}
            _ = v12
        }
        var v13 uint = 260
        v14 := GenS1{F0: 264}
        _ = fi11
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v6 = false
        var v16 uint8 = 127
        var v17 uint64 = 4294967295
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = v6
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]int)
    var v3 uint32 = 65535
    v4 := make(map[string]string)
    v18 := GenS1{F0: 1}
    _ = len(v2)
    v19 := GenS1{F0: genFn1()}
    v1 = "\n\t"
    var v20 bool = false
    switch 2147483647 {
    case -1:
        v21 := []int{127, 128}
        _ = v21
    default:
        v20 = false
    }
    if v20 {
        var v22 bool = false
        delete(v2, 4294967295)
        for fi23 := 0; fi23 < 5; fi23++ {
            switch ("\n\t" + "\n\t") {
            case "a":
                v3 += v3
            case "\n\t":
                for fi24 := 0; fi24 < 15; fi24++ {
                    v1 = "\n\t"
                    fmt.Println(v1)
                    var v25 float32 = -2.25
                    if (true || ((false || (9223372036854775807 >= fi23)) || v20)) {
                        fmt.Println(fi23, v3)
                        fmt.Println(fi24, fi23)
                        delete(v2, 2147483647)
                    } else {
                        v26 := GenS2{F0: -1, F1: 0}
                        if (2147483647 == 286) {
                            if (0.0 == float64(float32(-2.25))) {
                                var v27 int32 = -1
                                if (v22 && (v22 && (1 != 1))) {
                                    v3 += 160
                                    if v22 {
                                        v1 = ("" + v1)
                                    } else {
                                        if ((!false) || (127 < (fi24 | fi24))) {
                                            if (876 == 4294967295) {
                                                if true {
                                                    var v28 bool = (1 == -1)
                                                    switch fi23 {
                                                    case -1:
                                                        v22 = v20
                                                    case 128:
                                                        var v29 uint8 = 128
                                                        _ = v29
                                                    default:
                                                        v30 := v4[v1]
                                                        _ = v30
                                                    }
                                                    v31 := []int{32768, (fi23 ^ fi24), (fi24 & (fi23 & 1)), (-(fi24))}
                                                    var v32 int32 = v27
                                                    v20 = (false && (v22 || (127 <= 690)))
                                                    _ = v28
                                                    _ = v31
                                                    _ = v32
                                                } else {
                                                    v22 = true
                                                    v33 := GenS3{F0: int16(v3)}
                                                    v1 = v1
                                                    var v34 uint8 = 42
                                                    var v35 uint64 = 42
                                                    var v36 uint8 = (v34 / ((v34) | 1))
                                                    v34 = (v34 - (v36 / ((127) | 1)))
                                                    var v37 int64 = 42
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                    _ = v36
                                                    _ = v37
                                                }
                                                v38 := []int{}
                                                switch (fi24 & (fi24 | 919)) {
                                                case -1:
                                                    v39 := GenS2{F0: 0, F1: 306}
                                                    _ = v39
                                                case 1:
                                                    var v40 float64 = (((3.14159 * 1e10) + float64(3.14159)) * -2.25)
                                                    _ = v40
                                                case 32767:
                                                    if true {
                                                        var v41 uint = 255
                                                        var v42 int8 = 0
                                                        var v43 uint8 = 128
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                default:
                                                    v3 = (v3 % (((v3 + (^(v3)))) | 1))
                                                }
                                                v27 = (v27 << ((255) & 63))
                                                _ = v38
                                            } else {
                                                v44 := []int{fi23, (fi24 ^ (fi23 ^ (fi24 | (fi23 * fi24)))), (fi23 / ((260) | 1))}
                                                var v45 float64 = (0.0 * 1.5)
                                                var v46 int8 = -128
                                                _ = v44
                                                _ = v45
                                                _ = v46
                                            }
                                        } else {
                                            fmt.Println(v27, fi24, fi24)
                                            var v47 int8 = 86
                                            _ = v47
                                        }
                                    }
                                    v20 = (v3 <= v3)
                                } else {
                                    var v48 uint32 = 2147483647
                                    v49 := GenS3{F0: 560}
                                    _ = v48
                                    _ = v49
                                }
                                v3 -= v3
                                _ = v27
                            } else {
                                v22 = false
                                if false {
                                    switch "a" {
                                    case "\n\t":
                                        v22 = false
                                    case "a":
                                        fmt.Println(v25, v1, v22)
                                    case "hello":
                                        v50 := v4[(v1 + v1)]
                                        _ = v50
                                    default:
                                        var v51 int = (fi24 - (fi23 - 32767))
                                        _ = v51
                                    }
                                    if (1 < -96) {
                                        fmt.Println(v25)
                                        v52 := GenS3{F0: 42}
                                        _ = len(v4)
                                        _ = v52
                                    } else {
                                        if v22 {
                                            switch v1 {
                                            case "hello":
                                                fmt.Println(v3)
                                            case "foo bar":
                                                var v53 int8 = 12
                                                _ = v53
                                            }
                                            if v22 {
                                                var v54 float64 = float64(0.0)
                                                fmt.Println(v1, v3)
                                                var v55 int64 = -1
                                                _ = v54
                                                _ = v55
                                            }
                                            fmt.Println(fi23, v22)
                                        }
                                        v56 := GenS1{F0: -328}
                                        v1 = v1
                                        _ = v56
                                    }
                                    v22 = v20
                                } else {
                                    v3 = (^(v3))
                                    if (genFn0(genFn0(42, -665), 1) <= genFn0(genFn0(0, 32767), -72)) {
                                        var v57 int32 = -220
                                        if (!v22) {
                                            if true {
                                                var v58 int16 = 35
                                                switch ("hello" + "") {
                                                case "":
                                                    if (1 >= 2) {
                                                        fmt.Println(v3)
                                                    }
                                                case "a":
                                                    fmt.Println(fi23, v1, v1)
                                                }
                                                if (!(!(-1 < 1))) {
                                                    v59 := GenS4{F0: fi23}
                                                    v58 = v58
                                                    _ = v59
                                                } else {
                                                    v22 = (true && (((v1 < ("" + ("hello" + v1))) || (89 > 42)) && v22))
                                                }
                                                _ = v58
                                            }
                                        } else {
                                            switch ((("\n\t" + "") + v1) + v1) {
                                            case "\n\t":
                                                if v20 {
                                                    v60 := []int{fi23, 965}
                                                    fmt.Println(fi23)
                                                    if len(v60) > 0 {
                                                        v60[0] = (-(fi23))
                                                    }
                                                    var v61 uint = 42
                                                    _ = v60
                                                    _ = v61
                                                } else {
                                                    v62 := v4[(v1 + "")]
                                                    var v63 int64 = 0
                                                    fmt.Println(v57, v22)
                                                    var v64 float32 = 0.1
                                                    switch -2147483648 {
                                                    case 0:
                                                        v63 = v63
                                                    case -912:
                                                        fmt.Println(v64, v25)
                                                    default:
                                                        v65 := GenS5{F0: (^(v57))}
                                                        _ = v65
                                                    }
                                                    _ = v62
                                                    _ = v63
                                                    _ = v64
                                                }
                                            default:
                                                v3 = (v3 >> ((0) & 63))
                                            }
                                            fmt.Println(v57, v25, fi23)
                                        }
                                        switch v1 {
                                        case "hello":
                                            v66 := []int{}
                                            _ = v66
                                        case "foo bar":
                                            if v22 {
                                                var v67 int32 = genFn0(v57, 1000)
                                                _ = v67
                                            } else {
                                                fmt.Println(v57, fi23)
                                                if false {
                                                    fmt.Println(v1, fi24)
                                                } else {
                                                    switch fi23 {
                                                    case 127:
                                                        delete(v2, (v3 * (v3 - 1)))
                                                    }
                                                }
                                                v68 := GenS1{F0: 1}
                                                _ = v68
                                            }
                                        case "\n\t":
                                            var v69 int32 = genFn0((-(v57)), 1)
                                            var v70 uint32 = (v3 ^ v3)
                                            _ = v69
                                            _ = v70
                                        default:
                                            var v71 int64 = 389
                                            _ = v71
                                        }
                                        v72 := GenS2{F0: int64(fi23), F1: 4294967295}
                                        _ = v57
                                        _ = v72
                                    } else {
                                        var v73 int = (fi24 >> ((4294967295) & 63))
                                        var v74 int = 0
                                        fmt.Println(v22)
                                        _ = v73
                                        _ = v74
                                    }
                                    var v75 uint8 = 255
                                    _ = v75
                                }
                                if true {
                                    if v20 {
                                        if (4294967295 >= 2) {
                                            var v76 bool = (!v22)
                                            var v77 bool = (!false)
                                            fmt.Println(v25, v3)
                                            _ = v76
                                            _ = v77
                                        }
                                        if v20 {
                                            var v78 int8 = 42
                                            v79 := GenS1{F0: 1}
                                            if (!(v25 >= 1.5)) {
                                                var v80 uint64 = 0
                                                _ = v80
                                            } else {
                                                switch v1 {
                                                case "":
                                                    if (65535 <= 994) {
                                                        switch ("" + "") {
                                                        case "\n\t":
                                                            var v81 uint16 = 371
                                                            _ = v81
                                                        case "hello":
                                                            v78 = 42
                                                        }
                                                        var v82 int64 = -1
                                                        v83 := GenS6{F0: genFn1(), F1: -91, F2: -75}
                                                        v84 := GenS6{F0: genFn1(), F1: 1000, F2: 1000}
                                                        _ = v82
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                case "hello":
                                                    if true {
                                                        if v20 {
                                                            var v85 uint = 0
                                                            _ = v85
                                                        } else {
                                                            if (true || (1 > 255)) {
                                                                var v86 int8 = v78
                                                                _ = v86
                                                            } else {
                                                                v87 := GenS2{F0: -1, F1: 18446744073709551615}
                                                                _ = v87
                                                            }
                                                            var v88 float64 = (-2.25 - (0.0 + 0.1))
                                                            _ = v88
                                                        }
                                                        v20 = true
                                                        v89 := v2[(v3 * 4294967295)]
                                                        _ = v89
                                                    } else {
                                                        var v90 uint8 = 255
                                                        if (fi24 <= (fi24 & 32768)) {
                                                            fmt.Println(v1, v90)
                                                            if v20 {
                                                                var v91 string = ("" + (((((v1 + v1) + v1) + v1) + v1) + (v1 + (v1 + v1))))
                                                                v92 := GenS1{F0: 42}
                                                                _ = v91
                                                                _ = v92
                                                            }
                                                            v93 := []int{fi23, (fi24 + 368)}
                                                            _ = v93
                                                        } else {
                                                            v94 := GenS5{F0: 2147483647}
                                                            _ = v94
                                                        }
                                                        v95 := GenS7{F0: (!true), F1: 2}
                                                        switch v1 {
                                                        case "":
                                                            _ = len(v4)
                                                        case "foo bar":
                                                            v96 := GenS8{F0: v90, F1: 0, F2: 255}
                                                            v78 = v78
                                                            if v22 {
                                                                v20 = (-32768 != -130)
                                                                if (!v20) {
                                                                    v97 := GenS9{F0: (v3 % ((v3) | 1))}
                                                                    v98 := GenS9{F0: (v3 - v3)}
                                                                    _ = v97
                                                                    _ = v98
                                                                }
                                                            }
                                                            _ = v96
                                                        case "hello":
                                                            v99 := GenS10{F0: 1, F1: -32768}
                                                            _ = v99
                                                        default:
                                                            var v100 int8 = -1
                                                            _ = v100
                                                        }
                                                        _ = v90
                                                        _ = v95
                                                    }
                                                default:
                                                    v78 += (v78 / (((v78 << ((0) & 63))) | 1))
                                                }
                                                v101 := GenS1{F0: 42}
                                                v102 := GenS11{F0: 2, F1: (^(v3)), F2: -403}
                                                v3 = v3
                                                if false {
                                                    if v22 {
                                                        var v103 int64 = -1
                                                        _ = v103
                                                    }
                                                    var v104 bool = (float64(0.0) >= -2.25)
                                                    v25 -= float32(0.1)
                                                    delete(v4, (((v1 + (v1 + v1)) + v1) + v1))
                                                    _ = v104
                                                }
                                                _ = v101
                                                _ = v102
                                            }
                                            _ = v78
                                            _ = v79
                                        } else {
                                            if (v20 || (false && (42 <= 255))) {
                                                v105 := GenS7{F0: (v20 || v20), F1: 128}
                                                v1 = "hello"
                                                v106 := []int{}
                                                _ = v105
                                                _ = v106
                                            }
                                        }
                                    } else {
                                        if (fi24 <= fi23) {
                                            switch v1 {
                                            case "hello":
                                                delete(v4, ((v1 + "") + (v1 + ("hello" + "foo bar"))))
                                            case "\n\t":
                                                if (((v25 - v25) + v25) == float32((0.1 + float64((0.0 - v25))))) {
                                                    v107 := []int{fi24}
                                                    _ = v107
                                                } else {
                                                    var v108 int = fi23
                                                    var v109 int = 1
                                                    v22 = (((v1 + "\n\t") + ("foo bar" + v1)) > (v1 + v1))
                                                    _ = v108
                                                    _ = v109
                                                }
                                            }
                                            v2[2147483647] = fi24
                                            if v22 {
                                                fmt.Println(v3)
                                                v4[(v1 + v1)] = "hello"
                                                if ((("hello" + ("foo bar" + (v1 + (v1 + v1)))) >= "foo bar") && v20) {
                                                    if (0 <= -106) {
                                                        if (128 == 82) {
                                                            delete(v2, (^(v3)))
                                                        }
                                                        v25 -= 1e10
                                                        var v110 int8 = 42
                                                        _ = v110
                                                    } else {
                                                        v25 = (1.5 + v25)
                                                        switch (fi24 + (-(fi24))) {
                                                        case 32767:
                                                            fmt.Println(v1)
                                                        case -2147483648:
                                                            var v111 int8 = 0
                                                            _ = v111
                                                        case 339:
                                                            var v112 int16 = 982
                                                            _ = v112
                                                        }
                                                    }
                                                    v1 = v1
                                                    var v113 uint32 = v3
                                                    _ = v113
                                                }
                                            } else {
                                                var v114 int = -143
                                                v22 = true
                                                fmt.Println(v3, v22, fi24)
                                                _ = v114
                                            }
                                        } else {
                                            var v115 int32 = -2147483648
                                            v22 = (fi23 != (fi23 & 32767))
                                            v25 = v25
                                            _ = v115
                                        }
                                        v116 := GenS12{F0: 76}
                                        _ = v116
                                    }
                                }
                            }
                        }
                        _ = v26
                    }
                    _ = fi24
                    _ = v25
                }
            }
            v3 -= 0
            fmt.Println(v1)
            _ = fi23
        }
        _ = v22
    }
    for fi117 := 0; fi117 < 20; fi117++ {
        var v118 int = fi117
        _ = fi117
        _ = v118
    }
    v1 = "hello"
    v20 = (1 >= 2)
    if (189 >= 127) {
        var v119 uint8 = 1
        v119 = (^(v119))
        _ = v119
    }
    for fi120 := 0; fi120 < 8; fi120++ {
        _ = len(v2)
        var v121 int = fi120
        _ = fi120
        _ = v121
    }
    fmt.Println(v20, v1, v20)
    fmt.Println(v20, v20, v1)
    var v122 int64 = -556
    v122 -= (^(v122))
    fmt.Println("v1:", v1)
    fmt.Println("v122:", v122)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v122
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
