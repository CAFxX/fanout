package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 bool; F2 int32 }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 int16; F1 uint32 }
type GenS8 struct { F0 uint16; F1 int16; F2 bool }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint; F1 int; F2 bool }
type GenS11 struct { F0 bool; F1 int8 }
type GenS12 struct { F0 int16; F1 uint8 }
type GenS13 struct { F0 uint8; F1 int8; F2 string }
type GenS14 struct { F0 int; F1 int }
type GenS15 struct { F0 int; F1 int; F2 uint8 }
func genFn0(p0 int16) uint16 {
    switch ((("" + "\n\t") + (("a" + ("\n\t" + (("" + "\n\t") + "hello"))) + "a")) + ("hello" + "\n\t")) {
    case "hello":
        var v7 int16 = p0
        _ = v7
    case "a":
        v8 := GenS1{F0: 42, F1: 1}
        _ = v8
    case "":
        fmt.Println(p0)
    default:
        var v9 int8 = int8(p0)
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        if (true && true) {
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 float64 = float64(-2.25)
                v13 := []int{(^(fi11))}
                _ = fi11
                _ = v12
                _ = v13
            }
            fmt.Println(fi10)
            fmt.Println(p0, p0)
        } else {
            p0 += (p0 - p0)
        }
        _ = fi10
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 string = "foo bar"
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := []int{fi16, fi16, 32767}
        p0 = 1
        _ = fi16
        _ = v17
    }
    _ = p0
    return 2
}
func genFn1(p0 uint, p1 uint8) int64 {
    if (100 < int8(p0)) {
        var v18 bool = false
        p0 = 535
        if false {
            fmt.Println(p0, p1)
        }
        _ = v18
    }
    v19 := GenS1{F0: 1, F1: 2147483647}
    var v20 uint = 4294967295
    var v21 int = 4294967295
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    _ = v21
    return -1
}

func main() {
    var v1 string = ((((("a" + "foo bar") + "foo bar") + "foo bar") + (("foo bar" + ("foo bar" + ("hello" + "\n\t"))) + "a")) + "a")
    var v2 int8 = -1
    var v3 string = v1
    v4 := make(map[int]int)
    v5 := make(map[string]string)
    var v6 uint64 = 2
    if false {
        v22 := GenS2{F0: 32767, F1: (v6 & (v6 >> ((1) & 63)))}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 15; fi23++ {
            v24 := []int{fi23}
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
            _ = fi23
            _ = v24
        }
        var v25 int8 = (v2 << ((54) & 63))
        _ = v25
    }
    v4[-485] = 112
    var v26 int8 = (v2 % ((100) | 1))
    var v27 bool = true
    v28 := []int{}
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v27, v1)
    }
    if ((v2 & (v26 >> ((2) & 63))) <= v26) {
        var v29 bool = ((v3 == v3) || v27)
        v28 = append(v28, 9223372036854775807)
        v29 = ((!((v6 ^ (v6 & v6)) == (v6 << ((534) & 63)))) || v27)
        _ = v29
    }
    v30 := GenS3{F0: v1, F1: 255}
    v6 -= (v6 << ((4294967295) & 63))
    v31 := GenS4{F0: true}
    var v32 uint8 = 42
    for fi33 := 0; fi33 < 10; fi33++ {
        if len(v28) > 0 {
            v28[0] = -2147483648
        }
        if ((v32 >> ((65535) & 63)) > v32) {
            for fi34 := 0; fi34 < 3; fi34++ {
                if v27 {
                    v32 += 1
                    _ = len(v5)
                    var v35 float32 = 0.1
                    _ = v35
                } else {
                    var v36 int8 = (^(v2))
                    switch ("foo bar" + "hello") {
                    case "a":
                        var v37 uint = 485
                        _ = v37
                    case "hello":
                        if v27 {
                            v38 := GenS5{F0: 1000, F1: true, F2: 42}
                            v27 = v27
                            _ = len(v4)
                            _ = v38
                        } else {
                            fmt.Println(v6, v3, v6)
                            var v39 uint = 214
                            _ = v39
                        }
                    }
                    fmt.Println(v2)
                    _ = v36
                }
                _ = fi34
            }
        } else {
            switch (fi33 % ((-9223372036854775808) | 1)) {
            case 1:
                var v40 int16 = 1000
                _ = v40
            case 127:
                v32 += v32
            default:
                v1 = (v1 + v1)
            }
            for fi41 := 0; fi41 < 10; fi41++ {
                var v42 uint16 = 1
                switch v3 {
                case "hello":
                    var v43 bool = ((v3 > v1) || v27)
                    _ = v43
                case "foo bar":
                    if v27 {
                        v44 := GenS2{F0: (^(fi41)), F1: v6}
                        fmt.Println(fi33, v26)
                        _ = v44
                    }
                case "a":
                    var v45 float64 = 0.0
                    _ = v45
                default:
                    v46 := GenS6{F0: v42, F1: 1}
                    _ = v46
                }
                var v47 uint = 255
                _ = fi41
                _ = v42
                _ = v47
            }
            var v48 int32 = 0
            _ = v48
        }
        _ = fi33
    }
    if v27 {
        v6 += (v6 % ((v6) | 1))
        v49 := GenS4{F0: ((!(float32(0.1) < (1e10 - float32(-2.25)))) || ("hello" > v3))}
        _ = v49
    }
    var v50 uint = 2
    v27 = (uint16(v50) != 1)
    if ((v50 + 345) == (^(v50))) {
        if len(v28) > 0 {
            v28[0] = 9223372036854775807
        }
        v2 -= (v2 % ((v2) | 1))
        for fi51 := 0; fi51 < 15; fi51++ {
            v27 = v27
            var v52 int64 = -1
            _ = fi51
            _ = v52
        }
    } else {
        for fi53 := 0; fi53 < 3; fi53++ {
            for fi54 := 0; fi54 < 9; fi54++ {
                switch "" {
                case "\n\t":
                    switch v1 {
                    case "":
                        v55 := GenS7{F0: 0, F1: uint32(v32)}
                        _ = v55
                    case "a":
                        if (65535 < 964) {
                            var v56 uint32 = 42
                            _ = v56
                        } else {
                            v50 = uint(v32)
                        }
                    }
                case "a":
                    v57 := GenS5{F0: 258, F1: v27, F2: 2147483647}
                    _ = v57
                default:
                    switch -876 {
                    case 32767:
                        if (!(0 < 358)) {
                            if false {
                                v2 -= v26
                                var v58 bool = ((^(v6)) < 613)
                                fmt.Println(v32, fi54, v32)
                                _ = v58
                            } else {
                                if (!(false && v27)) {
                                    v59 := GenS6{F0: 0, F1: 920}
                                    _ = v59
                                } else {
                                    v1 = v3
                                    if (!(v2 <= v26)) {
                                        v60 := GenS2{F0: 128, F1: 1}
                                        _ = v60
                                    } else {
                                        v4[-700] = fi54
                                        v61 := GenS6{F0: 349, F1: 1}
                                        var v62 uint = (v50 | v50)
                                        _ = v61
                                        _ = v62
                                    }
                                    var v63 int32 = 42
                                    _ = v63
                                }
                            }
                            switch 32767 {
                            case 9223372036854775807:
                                var v64 uint8 = (v32 ^ uint8(v50))
                                _ = v64
                            case 1:
                                switch v3 {
                                case "hello":
                                    var v65 uint16 = 2
                                    if (!((778 < 1000) || (v27 && ((v2 - v26) > v2)))) {
                                        v65 = v65
                                        if ((1 >= int16(v32)) && true) {
                                            if len(v28) >= 2 {
                                                v28 = v28[1:len(v28)]
                                            }
                                            switch (((v3 + v3) + (v3 + ((v3 + "") + ((v1 + v3) + ((((v1 + ("hello" + "hello")) + "a") + "\n\t") + ((v3 + v1) + v3)))))) + v3) {
                                            case "":
                                                v66 := GenS1{F0: uint16(v32), F1: 0}
                                                _ = v66
                                            }
                                            v65 = 65535
                                            fmt.Println(v26)
                                            v1 = "foo bar"
                                            var v67 uint8 = (v32 | v32)
                                            _ = v67
                                        } else {
                                            if (!(!(v2 >= v2))) {
                                                v68 := GenS7{F0: 1000, F1: 42}
                                                var v69 uint32 = 0
                                                _ = v68
                                                _ = v69
                                            } else {
                                                switch fi54 {
                                                case 32767:
                                                    delete(v4, (fi54 ^ fi53))
                                                default:
                                                    v3 = "\n\t"
                                                }
                                                if v27 {
                                                    v70 := GenS4{F0: v27}
                                                    _ = v70
                                                }
                                            }
                                            if ((!v27) || true) {
                                                if v27 {
                                                    fmt.Println(v26, v27, v50)
                                                    switch -112 {
                                                    case -1:
                                                        var v71 uint64 = (^(v6))
                                                        _ = v71
                                                    case 9223372036854775807:
                                                        v26 -= (v2 % (((v26 + int8(v6))) | 1))
                                                    case 2147483647:
                                                        v3 = (("\n\t" + v3) + (v1 + v3))
                                                    default:
                                                        v28 = append(v28, (fi54 & (fi54 * (-(fi54)))))
                                                    }
                                                }
                                                v1 = "hello"
                                            }
                                            if ((float32(0.1) <= (0.1 * 1e10)) && (uint16(v50) == uint16(v50))) {
                                                v27 = false
                                                var v72 float32 = (0.0 - 1.5)
                                                _ = v72
                                            }
                                            var v73 bool = v27
                                            if v73 {
                                                var v74 int = (fi53 ^ fi54)
                                                switch fi53 {
                                                case 0:
                                                    if ((1e10 / (-2.25 * (3.14159 / (((1.5 * 3.14159) + 0.1) / float64(float32(-2.25)))))) > (-2.25 / 0.1)) {
                                                        switch v3 {
                                                        case "a":
                                                            switch fi53 {
                                                            case 42:
                                                                if (!(v73 || v73)) {
                                                                    if v27 {
                                                                        v73 = (false && v73)
                                                                        fmt.Println(v50, v65, v32)
                                                                        v75 := GenS8{F0: v65, F1: 42, F2: v73}
                                                                        _ = v75
                                                                    } else {
                                                                        if false {
                                                                            fmt.Println(v26, v26)
                                                                            v76 := GenS9{F0: (v50 & v50)}
                                                                            _ = v76
                                                                        } else {
                                                                            if len(v28) >= 2 {
                                                                                v28 = v28[1:len(v28)]
                                                                            }
                                                                        }
                                                                        v73 = v27
                                                                    }
                                                                    var v77 uint = v50
                                                                    _ = v77
                                                                } else {
                                                                    v3 = ((v3 + v3) + v3)
                                                                    var v78 float32 = -2.25
                                                                    if v27 {
                                                                        v78 = v78
                                                                        v73 = (fi54 != (fi54 | 32768))
                                                                    } else {
                                                                        v79 := GenS8{F0: uint16(v32), F1: 0, F2: (!true)}
                                                                        if true {
                                                                            _ = len(v5)
                                                                        } else {
                                                                            var v80 uint16 = genFn0(32767)
                                                                            _ = v80
                                                                        }
                                                                        if len(v28) >= 2 {
                                                                            v28 = v28[1:len(v28)]
                                                                        }
                                                                        _ = v79
                                                                    }
                                                                    _ = v78
                                                                }
                                                            case -9223372036854775808:
                                                                switch v1 {
                                                                case "":
                                                                    var v81 int16 = 281
                                                                    _ = v81
                                                                case "\n\t":
                                                                    _ = len(v5)
                                                                case "hello":
                                                                    if (-1 >= 635) {
                                                                        switch (fi54 | 127) {
                                                                        case 32768:
                                                                            if len(v28) > 0 {
                                                                                v28[0] = -690
                                                                            }
                                                                        }
                                                                        v32 -= v32
                                                                    }
                                                                default:
                                                                    v50 -= 422
                                                                }
                                                            default:
                                                                var v82 uint16 = (v65 << (((v50 + (v50 * v50))) & 63))
                                                                _ = v82
                                                            }
                                                        case "\n\t":
                                                            v26 += v2
                                                        case "":
                                                            var v83 bool = (!true)
                                                            _ = v83
                                                        }
                                                        var v84 uint16 = genFn0(1)
                                                        var v85 uint16 = 722
                                                        _ = v84
                                                        _ = v85
                                                    } else {
                                                        v28 = append(v28, 4294967295)
                                                        v27 = (!false)
                                                        v86 := v4[(fi54 >> ((v50) & 63))]
                                                        _ = v86
                                                    }
                                                case -1:
                                                    v87 := GenS10{F0: v50, F1: (-(fi53)), F2: v73}
                                                    _ = v87
                                                case -984:
                                                    switch fi54 {
                                                    case -223:
                                                        switch "a" {
                                                        case "hello":
                                                            fmt.Println(v32)
                                                        }
                                                    }
                                                default:
                                                    v65 = v65
                                                }
                                                _ = v74
                                            } else {
                                                if false {
                                                    var v88 uint8 = v32
                                                    if false {
                                                        v89 := GenS8{F0: (^(v65)), F1: 1000, F2: true}
                                                        _ = v89
                                                    }
                                                    if ((1.5 / 0.1) == 1e10) {
                                                        v50 += 919
                                                        v90 := GenS2{F0: (-(fi53)), F1: uint64(v65)}
                                                        _ = v90
                                                    } else {
                                                        v3 = (((v1 + "hello") + (v1 + v1)) + (v3 + ("hello" + v3)))
                                                    }
                                                    _ = v88
                                                } else {
                                                    _ = len(v28)
                                                    v3 = v1
                                                }
                                                v50 -= 458
                                            }
                                            var v91 float32 = (0.0 - 0.1)
                                            _ = v73
                                            _ = v91
                                        }
                                    } else {
                                        var v92 int16 = 32767
                                        var v93 string = ((v3 + v1) + ("foo bar" + ("hello" + v1)))
                                        if ((v6 * 2) == (v6 << (((^(v50))) & 63))) {
                                            var v94 string = ""
                                            v94 = (v1 + v3)
                                            _ = v94
                                        }
                                        v93 = "foo bar"
                                        v95 := v5[v1]
                                        _ = v92
                                        _ = v93
                                        _ = v95
                                    }
                                    _ = v65
                                case "\n\t":
                                    v1 = v1
                                    var v96 int64 = 740
                                    switch v1 {
                                    case "":
                                        fmt.Println(v26)
                                    }
                                    _ = v96
                                }
                            default:
                                switch (v1 + v3) {
                                case "a":
                                    v32 = v32
                                    var v97 int16 = -1
                                    v98 := GenS11{F0: (65535 > genFn0((v97 | (-(v97))))), F1: (v2 << ((v50) & 63))}
                                    _ = v97
                                    _ = v98
                                case "foo bar":
                                    var v99 uint32 = 281
                                    _ = v99
                                }
                            }
                            v1 = ""
                        }
                    case 128:
                        var v100 string = v1
                        _ = v100
                    case -331:
                        var v101 uint16 = genFn0(42)
                        _ = v101
                    default:
                        var v102 int = (fi54 >> ((v50) & 63))
                        _ = v102
                    }
                }
                if ((0.1 != 0.1) && (uint64(v50) < 255)) {
                    var v103 bool = false
                    _ = len(v4)
                    _ = v103
                }
                if true {
                    var v104 uint64 = v6
                    _ = v104
                } else {
                    switch v1 {
                    case "":
                        v5[(v1 + "\n\t")] = "foo bar"
                    case "hello":
                        if v27 {
                            v26 = int8(v32)
                            if (int16(v50) == 32767) {
                                var v105 uint32 = 240
                                v6 -= v6
                                _ = v105
                            }
                            switch v1 {
                            case "\n\t":
                                v26 = -94
                            case "hello":
                                _ = len(v4)
                            default:
                                fmt.Println(v32, v27)
                            }
                        }
                    case "a":
                        if v27 {
                            switch v3 {
                            case "":
                                v5[v1] = "\n\t"
                            default:
                                v4[(fi54 ^ 2147483647)] = 1
                            }
                        } else {
                            switch (fi54 / ((-324) | 1)) {
                            case 0:
                                v106 := GenS10{F0: v50, F1: fi54, F2: false}
                                _ = v106
                            default:
                                v27 = false
                            }
                        }
                    default:
                        v107 := GenS12{F0: -969, F1: v32}
                        _ = v107
                    }
                    switch 42 {
                    case 1:
                        var v108 float64 = -2.25
                        _ = v108
                    case 91:
                        if (v27 && (v32 > 1)) {
                            v3 = "\n\t"
                            v27 = (-1 == 1)
                        } else {
                            _ = len(v5)
                            var v109 int64 = 1
                            _ = v109
                        }
                    case 9223372036854775807:
                        switch fi53 {
                        case -833:
                            fmt.Println(v50)
                        default:
                            if ((!v27) || false) {
                                fmt.Println(v1, v2, v6)
                                fmt.Println(v50)
                                v110 := GenS12{F0: -32768, F1: (^(v32))}
                                _ = v110
                            }
                        }
                        switch 32768 {
                        case 9223372036854775807:
                            var v111 string = "hello"
                            _ = v111
                        case -1:
                            var v112 string = v1
                            _ = v112
                        default:
                            var v113 uint16 = 65535
                            _ = v113
                        }
                        v114 := GenS13{F0: uint8(v6), F1: v2, F2: v1}
                        _ = v114
                    default:
                        if ((-1 > 32767) && v27) {
                            switch v3 {
                            case "":
                                v6 = v6
                            case "\n\t":
                                var v115 int = 9223372036854775807
                                _ = v115
                            case "a":
                                var v116 int8 = -1
                                switch (fi54 / (((fi53 & fi53)) | 1)) {
                                case 42:
                                    var v117 string = "foo bar"
                                    _ = v117
                                case -621:
                                    v28 = append(v28, fi54)
                                }
                                if (0.1 < (float64(-2.25) - (0.0 + 0.1))) {
                                    var v118 int8 = (v26 ^ v26)
                                    var v119 uint8 = v32
                                    _ = v118
                                    _ = v119
                                }
                                if (2 >= genFn0(642)) {
                                    var v120 float32 = -2.25
                                    var v121 uint32 = 42
                                    if false {
                                        v5[("\n\t" + "")] = (v1 + ((v3 + (v1 + "foo bar")) + v1))
                                        v122 := GenS7{F0: 1, F1: v121}
                                        v123 := v5[v1]
                                        v26 += (v116 + -128)
                                        _ = v122
                                        _ = v123
                                    } else {
                                        _ = len(v28)
                                    }
                                    _ = v120
                                    _ = v121
                                }
                                var v124 float64 = (0.1 * float64(0.1))
                                _ = v116
                                _ = v124
                            default:
                                switch ((("hello" + "\n\t") + (v1 + v1)) + ((v1 + v3) + v3)) {
                                case "hello":
                                    var v125 string = ""
                                    _ = v125
                                case "":
                                    v126 := GenS14{F0: (fi53 % (((fi53 % ((441) | 1))) | 1)), F1: -974}
                                    _ = v126
                                }
                            }
                        }
                        if (!false) {
                            var v127 int32 = -1
                            v32 = v32
                            _ = v127
                        } else {
                            v1 = v1
                        }
                        if (401 < 1) {
                            fmt.Println(v26)
                            v128 := GenS3{F0: "\n\t", F1: 237}
                            _ = v128
                        } else {
                            delete(v5, ("foo bar" + (v1 + "foo bar")))
                            if (true && v27) {
                                if v27 {
                                    var v129 int = 180
                                    v5[(("a" + "") + v1)] = "hello"
                                    _ = v129
                                } else {
                                    var v130 float64 = 0.0
                                    if len(v28) > 0 {
                                        v28[0] = 127
                                    }
                                    _ = v130
                                }
                                v26 = (v2 * (v2 ^ int8(v6)))
                                var v131 uint8 = (^(v32))
                                v50 -= 1
                                v27 = (1 < 0)
                                _ = v131
                            } else {
                                v6 -= (v6 % (((v6 << ((v50) & 63))) | 1))
                                if ((v2 * v26) > v2) {
                                    v50 = 0
                                    fmt.Println(v6, fi54, v6)
                                    fmt.Println(v1)
                                }
                                if true {
                                    _ = len(v4)
                                }
                                v6 = (v6 / ((61) | 1))
                            }
                        }
                    }
                }
                _ = fi54
            }
            v3 = "hello"
            v132 := GenS15{F0: int(v2), F1: int(v6), F2: v32}
            _ = fi53
            _ = v132
        }
        v27 = (!(42 > -1))
        if (0 >= 970) {
            if len(v28) > 0 {
                v28[0] = 32768
            }
            var v133 uint32 = 1
            switch 32768 {
            case 9223372036854775807:
                v134 := GenS4{F0: (true && true)}
                _ = v134
            case 587:
                if (0 < v50) {
                    fmt.Println(v26, v27)
                    var v135 uint32 = (v133 % (((v133 + uint32(v50))) | 1))
                    v136 := GenS11{F0: (true && v27), F1: (v26 | v2)}
                    _ = v135
                    _ = v136
                } else {
                    v1 = v1
                    v137 := GenS8{F0: 255, F1: 1, F2: (!v27)}
                    if (v32 < 2) {
                        v32 = (v32 / ((uint8(v133)) | 1))
                        var v138 uint16 = genFn0(-32768)
                        _ = v138
                    }
                    _ = v137
                }
            case -9223372036854775808:
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
            default:
                fmt.Println(v3, v50)
            }
            _ = v133
        }
    }
    var v139 int8 = (v26 & 43)
    var v140 float32 = float32(3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v139:", v139)
    fmt.Println("v140:", v140)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v139
    _ = v140
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
