package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint; F2 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 uint; F2 uint }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 int16; F1 int16 }
func genFn0() int32 {
    var v7 uint32 = 150
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int16 = -32768
        switch fi8 {
        case -860:
            v9 = (v9 ^ v9)
        case 128:
            for fi10 := 0; fi10 < 11; fi10++ {
                fmt.Println(fi10)
                _ = fi10
            }
        case 247:
            v9 = v9
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            if (3.14159 <= ((0.1 / -2.25) - 3.14159)) {
                switch (fi8 + fi11) {
                case -2147483648:
                    fmt.Println(v9, v9, fi11)
                }
                switch "hello" {
                case "foo bar":
                    v12 := []int{fi8, (fi8 | (fi11 / (((fi8 % ((272) | 1))) | 1)))}
                    _ = v12
                default:
                    var v13 int8 = 91
                    _ = v13
                }
            }
            v14 := []int{fi11, (fi8 ^ (fi8 - fi8)), fi8}
            var v15 int = fi8
            _ = fi11
            _ = v14
            _ = v15
        }
        _ = fi8
        _ = v9
    }
    if (true && (v7 != (^(v7)))) {
        var v16 int16 = -776
        _ = v16
    }
    if false {
        if true {
            switch 32768 {
            case 32768:
                v17 := []int{840, -966}
                _ = v17
            case 128:
                var v18 int = 2147483647
                _ = v18
            default:
                var v19 int16 = 42
                _ = v19
            }
            var v20 int = 691
            fmt.Println(v20, v20)
            _ = v20
        }
        fmt.Println(v7)
        v21 := []int{-131, -1}
        _ = v21
    } else {
        switch "hello" {
        case "foo bar":
            fmt.Println(v7)
        }
    }
    _ = v7
    return 1
}
func genFn1() uint {
    var v22 int16 = 0
    v22 -= 1000
    for fi23 := 0; fi23 < 15; fi23++ {
        for fi24 := 0; fi24 < 15; fi24++ {
            fmt.Println(v22, v22, v22)
            _ = fi24
        }
        v25 := []int{fi23, fi23, fi23, fi23}
        _ = fi23
        _ = v25
    }
    v22 = (v22 + (-(v22)))
    v26 := []int{-2147483648, 0}
    _ = v22
    _ = v26
    return 65535
}

func main() {
    var v1 int8 = 127
    v2 := []string{}
    var v3 float32 = (3.14159 - ((float32(((1.5 + 1e10) * 1e10)) / 1e10) / 1.5))
    v4 := []uint32{1, 288, 2147483647, 4294967295}
    v5 := make(map[string]string)
    var v6 uint64 = 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v27 float64 = 1.5
    if true {
        if false {
            v1 -= v1
        }
        var v28 uint16 = 386
        v28 += v28
        _ = v28
    } else {
        switch 2147483647 {
        case -2147483648:
            v29 := GenS1{F0: 155}
            _ = v29
        case 1:
            if len(v2) > 0 {
                v2[0] = ("hello" + "hello")
            }
        default:
            if true {
                v30 := GenS1{F0: 127}
                _ = v30
            }
        }
        v6 -= (v6 - (v6 | v6))
        switch "\n\t" {
        case "a":
            for fi31 := 0; fi31 < 6; fi31++ {
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    if ((float32(v27) - v3) != 1.5) {
                        if true {
                            for fi32 := 0; fi32 < 15; fi32++ {
                                fmt.Println(v6)
                                _ = fi32
                            }
                            v3 = float32(1.5)
                        } else {
                            switch "a" {
                            case "":
                                for fi33 := 0; fi33 < 2; fi33++ {
                                    v3 += (1.5 - (v3 / (0.0 / v3)))
                                    v3 -= 0.0
                                    v27 -= ((3.14159 + 1.5) + 1e10)
                                    _ = fi33
                                }
                            default:
                                if true {
                                    _ = len(v5)
                                    v3 += float32(1.5)
                                    fmt.Println(v27, fi31, fi31)
                                }
                            }
                            _ = len(v4)
                            for fi34 := 0; fi34 < 7; fi34++ {
                                var v35 bool = false
                                v36 := GenS2{F0: 42}
                                v4 = append(v4, 148)
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                        }
                        for fi37 := 0; fi37 < 14; fi37++ {
                            if ((v27 / (v27 / (((v27 - 1e10) - float64(float32(0.0))) - (3.14159 / (0.0 / ((1.5 - v27) + (v27 / ((0.1 / (-2.25 / ((v27 - 0.0) - (v27 / v27)))) - (((v27 / (1.5 - v27)) * -2.25) - (float64(v3) * (float64(0.0) * v27))))))))))) == (((v27 - (3.14159 / 1.0)) + ((v27 + v27) - ((v27 + float64(v3)) - v27))) * 3.14159)) {
                                v27 += v27
                                if (-1 < 1) {
                                    switch ("\n\t" + ("foo bar" + "hello")) {
                                    case "a":
                                        v38 := GenS1{F0: 144}
                                        _ = v38
                                    case "foo bar":
                                        v39 := GenS2{F0: 2147483647}
                                        _ = v39
                                    case "\n\t":
                                        v6 = v6
                                        v6 += v6
                                        if false {
                                            var v40 uint = genFn1()
                                            if false {
                                                v41 := GenS2{F0: 2147483647}
                                                v42 := GenS3{F0: v1, F1: 2, F2: v1}
                                                var v43 int16 = 472
                                                v44 := GenS1{F0: 2}
                                                v3 -= v3
                                                _ = v41
                                                _ = v42
                                                _ = v43
                                                _ = v44
                                            }
                                            _ = v40
                                        } else {
                                            v45 := GenS2{F0: 307}
                                            _ = v45
                                        }
                                        _ = len(v5)
                                    default:
                                        v46 := GenS3{F0: 1, F1: 0, F2: (v1 >> ((2) & 63))}
                                        _ = v46
                                    }
                                    v27 = (((v27 * (1.5 - (v27 * v27))) + (3.14159 - float64((((v3 * -2.25) - 1e10) * 3.14159)))) / 1.0)
                                } else {
                                    if false {
                                        v3 -= float32(v27)
                                        v3 += v3
                                        switch (-(fi37)) {
                                        case -9223372036854775808:
                                            delete(v5, "a")
                                        case 1:
                                            v6 = v6
                                        case 851:
                                            v47 := GenS1{F0: 42}
                                            _ = v47
                                        default:
                                            v3 = 0.1
                                        }
                                    } else {
                                        var v48 bool = (134 == uint8(v6))
                                        _ = v48
                                    }
                                    if true {
                                        v49 := GenS3{F0: (v1 << ((4294967295) & 63)), F1: 18446744073709551615, F2: v1}
                                        if len(v4) > 0 {
                                            v4[0] = 562
                                        }
                                        _ = v49
                                    }
                                    v1 += 100
                                }
                                var v50 uint = 350
                                _ = v50
                            }
                            _ = fi37
                        }
                        if false {
                            var v51 uint8 = 0
                            if (42 < genFn0()) {
                                v1 += (v1 + int8(v51))
                            } else {
                                switch "hello" {
                                case "a":
                                    var v52 uint = 2
                                    _ = v52
                                default:
                                    var v53 int = -2147483648
                                    _ = v53
                                }
                            }
                            for fi54 := 0; fi54 < 3; fi54++ {
                                fmt.Println(v51)
                                if false {
                                    v1 += v1
                                    if (-703 != -1) {
                                        var v55 int = 127
                                        if (((v6 >> ((0) & 63)) < v6) || ((1.5 / v3) >= v3)) {
                                            switch "\n\t" {
                                            case "a":
                                                if true {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v56 := GenS1{F0: 2}
                                                        _ = v56
                                                    case "a":
                                                        v1 -= int8(fi31)
                                                    default:
                                                        switch "hello" {
                                                        case "hello":
                                                            var v57 uint32 = 2147483647
                                                            _ = v57
                                                        }
                                                    }
                                                } else {
                                                    v58 := GenS1{F0: v51}
                                                    v1 = (-(v1))
                                                    var v59 int8 = v1
                                                    _ = v58
                                                    _ = v59
                                                }
                                            case "\n\t":
                                                if (42 < 0) {
                                                    v60 := v5["hello"]
                                                    var v61 int = fi31
                                                    v55 -= 306
                                                    _ = v60
                                                    _ = v61
                                                } else {
                                                    v3 = v3
                                                    var v62 int32 = -305
                                                    _ = v62
                                                }
                                            }
                                        }
                                        v55 = v55
                                        _ = v55
                                    }
                                    v63 := GenS3{F0: v1, F1: 4294967295, F2: (v1 | v1)}
                                    if ((2 > genFn1()) || true) {
                                        if false {
                                            fmt.Println(fi31)
                                            if ((v6 - uint64(v51)) != v6) {
                                                if true {
                                                    if (2147483647 <= fi54) {
                                                        v3 -= float32((float64(v3) * v27))
                                                        var v64 int32 = -1
                                                        v6 = (v6 % ((v6) | 1))
                                                        _ = v64
                                                    }
                                                } else {
                                                    v65 := GenS4{F0: true}
                                                    _ = v65
                                                }
                                            } else {
                                                v66 := GenS5{F0: 32767, F1: 889, F2: genFn1()}
                                                _ = v66
                                            }
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = "\n\t"
                                            }
                                            v6 += (^(v6))
                                            if false {
                                                v27 = ((float64((float32((((1.5 + (v27 * v27)) * float64(float32(((((v27 - float64((1.5 / 1e10))) + v27) / float64(((-2.25 + float32(v27)) * float32(v27)))) * (0.1 + 1e10))))) - 1.5)) + v3)) - float64(1.5)) + 1.5)
                                            } else {
                                                v2 = append(v2, ("a" + ""))
                                                if false {
                                                    if true {
                                                        var v67 uint64 = (v6 / ((v6) | 1))
                                                        _ = len(v5)
                                                        if (!true) {
                                                            _ = len(v5)
                                                            switch fi54 {
                                                            case 128:
                                                                _ = len(v5)
                                                            default:
                                                                var v68 float32 = 1.5
                                                                _ = v68
                                                            }
                                                        }
                                                        _ = v67
                                                    } else {
                                                        v1 += 8
                                                        switch (^(fi54)) {
                                                        case 32768:
                                                            if (!true) {
                                                                switch 4294967295 {
                                                                case -2147483648:
                                                                    var v69 uint8 = (v51 >> ((1) & 63))
                                                                    var v70 uint = 1
                                                                    fmt.Println(v27, v27, v70)
                                                                    _ = v69
                                                                    _ = v70
                                                                case 1:
                                                                    var v71 float64 = (((v27 * (((float64(1.5) * float64(v3)) / (0.1 / 1e10)) / (v27 - (-2.25 - (v27 / v27))))) / 3.14159) / v27)
                                                                    _ = len(v2)
                                                                    v72 := v5["foo bar"]
                                                                    _ = v71
                                                                    _ = v72
                                                                }
                                                                var v73 float64 = v27
                                                                _ = v73
                                                            } else {
                                                                v3 += (v3 - (v3 * v3))
                                                                var v74 string = "foo bar"
                                                                _ = v74
                                                            }
                                                        case 127:
                                                            var v75 string = "\n\t"
                                                            _ = v75
                                                        }
                                                        if true {
                                                            var v76 bool = true
                                                            var v77 float64 = 3.14159
                                                            switch (("foo bar" + "") + "hello") {
                                                            case "hello":
                                                                v78 := GenS6{F0: 32767, F1: -196}
                                                                var v79 uint8 = v51
                                                                _ = v78
                                                                _ = v79
                                                            }
                                                            var v80 uint64 = (^(v6))
                                                            _ = v76
                                                            _ = v77
                                                            _ = v80
                                                        } else {
                                                            var v81 int8 = v1
                                                            _ = v81
                                                        }
                                                    }
                                                    var v82 uint8 = (v51 * 255)
                                                    if ((float64(3.14159) != 1.5) || (1 >= 175)) {
                                                        fmt.Println(v6)
                                                        switch (("" + ("a" + "")) + "\n\t") {
                                                        case "hello":
                                                            var v83 int64 = 1
                                                            _ = v83
                                                        }
                                                    } else {
                                                        v5["\n\t"] = ""
                                                        v82 = 128
                                                    }
                                                    if (int(v51) <= fi31) {
                                                        if true {
                                                            var v84 int16 = 42
                                                            var v85 string = ("a" + "hello")
                                                            switch "" {
                                                            case "\n\t":
                                                                var v86 uint64 = (^(v6))
                                                                _ = v86
                                                            case "foo bar":
                                                                v1 = 100
                                                                var v87 int8 = (-(v1))
                                                                _ = v87
                                                            }
                                                            _ = v84
                                                            _ = v85
                                                        } else {
                                                            v88 := GenS2{F0: genFn0()}
                                                            fmt.Println(v1, v1)
                                                            switch "" {
                                                            case "\n\t":
                                                                v2 = append(v2, "")
                                                            default:
                                                                fmt.Println(fi31, v6, v6)
                                                            }
                                                            _ = v88
                                                        }
                                                        if (float32(v27) <= v3) {
                                                            fmt.Println(v27)
                                                            v89 := GenS6{F0: 42, F1: 1}
                                                            var v90 float32 = ((1.5 - 1e10) * (v3 + v3))
                                                            _ = v89
                                                            _ = v90
                                                        }
                                                    }
                                                    v1 -= 127
                                                    v27 = float64(0.0)
                                                    _ = v82
                                                }
                                            }
                                            var v91 uint64 = (^(v6))
                                            if (2 < 438) {
                                                v3 = (1e10 - ((v3 / 1e10) + (v3 * v3)))
                                                if false {
                                                    v92 := GenS7{F0: 42, F1: "foo bar"}
                                                    v91 = (v6 ^ (v91 ^ v6))
                                                    v93 := GenS8{F0: -9223372036854775808, F1: v51}
                                                    _ = v92
                                                    _ = v93
                                                } else {
                                                    switch "" {
                                                    case "a":
                                                        var v94 uint8 = 127
                                                        _ = v94
                                                    case "":
                                                        if (genFn1() >= 42) {
                                                            var v95 int = 32768
                                                            fmt.Println(v1)
                                                            _ = v95
                                                        } else {
                                                            _ = len(v4)
                                                        }
                                                    case "foo bar":
                                                        v3 += 0.1
                                                    }
                                                    v96 := GenS9{F0: 32767, F1: 42}
                                                    if (572 < 454) {
                                                        _ = len(v2)
                                                    } else {
                                                        fmt.Println(v51, v1)
                                                        var v97 bool = false
                                                        var v98 int64 = 42
                                                        _ = v97
                                                        _ = v98
                                                    }
                                                    var v99 int16 = 1
                                                    var v100 int = 32768
                                                    _ = v100
                                                    _ = v96
                                                    _ = v99
                                                }
                                            }
                                            _ = v91
                                        }
                                        var v101 string = "hello"
                                        switch v101 {
                                        case "\n\t":
                                            var v102 int8 = (^(v1))
                                            _ = v102
                                        case "hello":
                                            var v103 uint16 = 255
                                            _ = v103
                                        default:
                                            switch "a" {
                                            case "a":
                                                var v104 float64 = (((v27 - v27) * ((v27 - float64((3.14159 * 1.5))) * (-2.25 / v27))) / (v27 - ((v27 - float64(v3)) - v27)))
                                                _ = v104
                                            case "\n\t":
                                                fmt.Println(v101)
                                            }
                                        }
                                        _ = v101
                                    }
                                    _ = v63
                                }
                                fmt.Println(v3, fi31, fi31)
                                _ = fi54
                            }
                            _ = v51
                        } else {
                            v1 = (v1 * 127)
                            var v105 int = (fi31 >> ((2) & 63))
                            var v106 float64 = (float64(1e10) / (1.5 - -2.25))
                            _ = v105
                            _ = v106
                        }
                    } else {
                        if (-673 >= 1) {
                            fmt.Println(v1, v27, v6)
                            var v107 uint = 206
                            _ = v107
                        } else {
                            delete(v5, "hello")
                            if (2 == 42) {
                                v27 += v27
                            }
                            var v108 uint8 = 227
                            _ = v108
                        }
                        var v109 float32 = v3
                        _ = v109
                    }
                }
                _ = fi31
            }
        case "":
            for fi110 := 0; fi110 < 7; fi110++ {
                if (4294967295 < 164) {
                    v111 := GenS6{F0: 485, F1: -9223372036854775808}
                    v1 = v1
                    for fi112 := 0; fi112 < 1; fi112++ {
                        if true {
                            v1 = v1
                            v3 -= (float32(float64(-2.25)) - v3)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            var v113 int = fi112
                            var v114 uint16 = 0
                            _ = v113
                            _ = v114
                        }
                        var v115 int8 = -112
                        delete(v5, "hello")
                        _ = fi112
                        _ = v115
                    }
                    _ = v111
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi116 := 0; fi116 < 20; fi116++ {
                        fmt.Println(fi110)
                        _ = fi116
                    }
                    var v117 uint = 0
                    _ = v117
                }
                _ = fi110
            }
        case "foo bar":
            fmt.Println(v27, v1, v6)
        default:
            fmt.Println(v6, v6, v1)
        }
    }
    var v118 uint16 = 1
    var v119 uint = uint(v118)
    v120 := GenS9{F0: 0, F1: -1}
    v1 = -1
    var v121 uint32 = 4294967295
    v6 -= 1
    v122 := GenS8{F0: 1, F1: 244}
    _ = len(v2)
    v6 = v6
    var v123 int16 = 687
    v2 = append(v2, "foo bar")
    v6 += v6
    v124 := GenS8{F0: 347, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("v121:", v121)
    fmt.Println("v123:", v123)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v118
    _ = v119
    _ = v120
    _ = v121
    _ = v122
    _ = v123
    _ = v124
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
