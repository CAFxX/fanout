package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 int16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int32; F1 uint16 }
type GenS11 struct { F0 int8; F1 int64; F2 uint64 }
func genFn0() uint8 {
    v4 := GenS1{F0: 65535, F1: 128}
    var v5 uint16 = 1
    switch 0 {
    case 2147483647:
        var v6 int8 = 0
        _ = v6
    case 1:
        var v7 float64 = (0.0 * 1.5)
        _ = v7
    case -2147483648:
        var v8 int8 = -128
        _ = v8
    }
    if false {
        v9 := GenS2{F0: 42, F1: 0}
        fmt.Println(v5)
        for fi10 := 0; fi10 < 3; fi10++ {
            v5 = v5
            fmt.Println(v5, fi10)
            var v11 uint32 = 2
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    v12 := []int{-84, 127, 2147483647}
    _ = v12
    _ = v4
    _ = v5
    return 70
}
func genFn1() uint32 {
    var v13 int64 = 1
    var v14 int = 136
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := GenS1{F0: 42, F1: 190}
            var v18 float32 = ((-2.25 * 3.14159) + 0.0)
            if (false || false) {
                v19 := GenS1{F0: 2, F1: 1}
                _ = v19
            }
            if false {
                switch (fi16 >> ((213) & 63)) {
                case -864:
                    if false {
                        switch (fi15 - (fi16 >> ((65535) & 63))) {
                        case -1:
                            if (true || true) {
                                v20 := GenS1{F0: 0, F1: 2}
                                v13 -= -973
                                if true {
                                    fmt.Println(v14, v18)
                                }
                                _ = v20
                            } else {
                                var v21 uint32 = 42
                                _ = v21
                            }
                            v22 := []int{127, 2147483647, (fi15 - fi15), (-(fi15))}
                            var v23 float32 = -2.25
                            var v24 int16 = 42
                            _ = v22
                            _ = v23
                            _ = v24
                        case -355:
                            fmt.Println(v18, v18)
                        }
                        v14 -= fi16
                        var v25 string = "foo bar"
                        fmt.Println(v14)
                        v25 = ((v25 + (v25 + v25)) + v25)
                        _ = v25
                    } else {
                        switch fi15 {
                        case 32767:
                            switch "hello" {
                            case "\n\t":
                                var v26 uint8 = 128
                                _ = v26
                            case "a":
                                var v27 bool = (float32(1e10) <= 3.14159)
                                if v27 {
                                    switch fi16 {
                                    case 32768:
                                        switch "a" {
                                        case "":
                                            fmt.Println(v27, v14)
                                        case "foo bar":
                                            v27 = v27
                                        case "a":
                                            v18 += ((((v18 - v18) + v18) * v18) * (v18 * ((1.5 * v18) * float32(float64(1e10)))))
                                        }
                                    case -9223372036854775808:
                                        if v27 {
                                            switch "a" {
                                            case "":
                                                v28 := []int{(fi15 >> ((393) & 63)), fi16, fi15, (-(v14))}
                                                if (602 != 65535) {
                                                    v29 := GenS2{F0: 117, F1: 127}
                                                    _ = v29
                                                }
                                                _ = v28
                                            case "a":
                                                v30 := []int{}
                                                _ = v30
                                            default:
                                                if false {
                                                    v31 := GenS2{F0: 2147483647, F1: 1}
                                                    v13 = v13
                                                    v32 := GenS1{F0: 0, F1: 35}
                                                    _ = v31
                                                    _ = v32
                                                } else {
                                                    var v33 uint8 = 9
                                                    switch (fi15 << ((4294967295) & 63)) {
                                                    case 32768:
                                                        v34 := []int{v14, v14}
                                                        _ = v34
                                                    default:
                                                        var v35 uint8 = (^(v33))
                                                        _ = v35
                                                    }
                                                    _ = v33
                                                }
                                            }
                                        } else {
                                            fmt.Println(v27, v13)
                                            fmt.Println(v27, fi16, v14)
                                            switch "foo bar" {
                                            case "a":
                                                var v36 uint64 = 1
                                                _ = v36
                                            default:
                                                switch ("a" + "\n\t") {
                                                case "hello":
                                                    var v37 uint32 = 2
                                                    _ = v37
                                                case "foo bar":
                                                    v18 = v18
                                                case "":
                                                    switch "a" {
                                                    case "":
                                                        v38 := []int{32767}
                                                        _ = v38
                                                    case "foo bar":
                                                        if (65535 == 0) {
                                                            v27 = (false && v27)
                                                        } else {
                                                            switch ("" + "a") {
                                                            case "":
                                                                var v39 uint32 = 2
                                                                _ = v39
                                                            case "\n\t":
                                                                switch (fi16 ^ 42) {
                                                                case -2147483648:
                                                                    var v40 uint8 = 111
                                                                    _ = v40
                                                                case 42:
                                                                    if (v27 || false) {
                                                                        var v41 uint = 4294967295
                                                                        var v42 float32 = 1.5
                                                                        _ = v41
                                                                        _ = v42
                                                                    } else {
                                                                        v43 := GenS1{F0: 65535, F1: 45}
                                                                        v44 := []int{fi16, (fi15 ^ (-(v14))), -9223372036854775808}
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                case 536:
                                                                    var v45 bool = true
                                                                    _ = v45
                                                                }
                                                            }
                                                        }
                                                    case "\n\t":
                                                        if (int16(v13) > 948) {
                                                            var v46 float32 = (v18 * float32(0.0))
                                                            var v47 uint32 = 65535
                                                            _ = v46
                                                            _ = v47
                                                        } else {
                                                            v48 := []int{9223372036854775807, fi16}
                                                            v14 = v14
                                                            var v49 string = ("\n\t" + "hello")
                                                            if (133 <= 361) {
                                                                v50 := GenS1{F0: 917, F1: 1}
                                                                switch v49 {
                                                                case "hello":
                                                                    v51 := GenS3{F0: 609, F1: 65535}
                                                                    _ = v51
                                                                case "\n\t":
                                                                    var v52 float32 = (0.0 * v18)
                                                                    _ = v52
                                                                case "foo bar":
                                                                    v53 := GenS4{F0: (v13 % ((v13) | 1))}
                                                                    var v54 int8 = 0
                                                                    _ = v53
                                                                    _ = v54
                                                                }
                                                                if len(v48) >= 2 {
                                                                    v48 = v48[1:len(v48)]
                                                                }
                                                                if false {
                                                                    v55 := GenS4{F0: 9223372036854775807}
                                                                    v18 = v18
                                                                    _ = v55
                                                                } else {
                                                                    v27 = (255 > 1)
                                                                    v18 = 1.5
                                                                }
                                                                if false {
                                                                    var v56 uint8 = 127
                                                                    _ = v56
                                                                } else {
                                                                    v57 := GenS5{F0: 324}
                                                                    _ = v57
                                                                }
                                                                _ = v50
                                                            } else {
                                                                switch v49 {
                                                                case "hello":
                                                                    v58 := GenS6{F0: 0}
                                                                    _ = v58
                                                                case "a":
                                                                    var v59 bool = (2147483648 <= 0)
                                                                    _ = v59
                                                                }
                                                                switch fi15 {
                                                                case 4294967295:
                                                                    if len(v48) >= 2 {
                                                                        v48 = v48[1:len(v48)]
                                                                    }
                                                                    v18 += v18
                                                                    v18 -= (((v18 / -2.25) * v18) - ((v18 - ((v18 - (v18 + v18)) * ((0.1 - v18) + 1.5))) / v18))
                                                                    if (!(-797 != 486)) {
                                                                        var v60 uint8 = 128
                                                                        v13 -= -1
                                                                        if len(v48) > 0 {
                                                                            v48[0] = fi15
                                                                        }
                                                                        var v61 float32 = 1.5
                                                                        fmt.Println(fi15, v13)
                                                                        if v27 {
                                                                            v62 := GenS7{F0: 18446744073709551615, F1: 1000}
                                                                            if len(v48) >= 2 {
                                                                                v48 = v48[1:len(v48)]
                                                                            }
                                                                            var v63 int64 = 612
                                                                            _ = v62
                                                                            _ = v63
                                                                        } else {
                                                                            switch v49 {
                                                                            case "a":
                                                                                v27 = ((fi15 - v14) == 42)
                                                                            default:
                                                                                var v64 uint32 = 2147483647
                                                                                _ = v64
                                                                            }
                                                                            fmt.Println(v14)
                                                                            var v65 float64 = (1e10 * 0.1)
                                                                            _ = v65
                                                                        }
                                                                        _ = v60
                                                                        _ = v61
                                                                    } else {
                                                                        v13 -= 0
                                                                    }
                                                                    v66 := GenS6{F0: 567}
                                                                    switch (("foo bar" + ((v49 + v49) + v49)) + ((v49 + v49) + (v49 + (v49 + v49)))) {
                                                                    case "foo bar":
                                                                        if len(v48) >= 2 {
                                                                            v48 = v48[1:len(v48)]
                                                                        }
                                                                    default:
                                                                        v27 = (3 >= 1)
                                                                    }
                                                                    _ = v66
                                                                case -770:
                                                                    _ = len(v48)
                                                                case 588:
                                                                    v18 = 0.1
                                                                default:
                                                                    var v67 uint8 = 2
                                                                    _ = v67
                                                                }
                                                                var v68 uint8 = 128
                                                                _ = v68
                                                            }
                                                            v49 = v49
                                                            _ = v48
                                                            _ = v49
                                                        }
                                                    default:
                                                        fmt.Println(v27, v27, v27)
                                                    }
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    if v27 {
                                        v18 = v18
                                    }
                                }
                                _ = v27
                            default:
                                if true {
                                    if false {
                                        v69 := GenS6{F0: 32767}
                                        switch ("hello" + ("hello" + ("foo bar" + "hello"))) {
                                        case "\n\t":
                                            if false {
                                                v70 := GenS3{F0: 2147483648, F1: 255}
                                                v71 := GenS1{F0: 42, F1: genFn0()}
                                                var v72 uint = 0
                                                _ = v70
                                                _ = v71
                                                _ = v72
                                            } else {
                                                v18 += float32(3.14159)
                                                if (v13 == v13) {
                                                    fmt.Println(fi15, fi16)
                                                    v73 := GenS8{F0: (((("foo bar" + "\n\t") + (("" + "foo bar") + "hello")) + "hello") + ("\n\t" + "foo bar"))}
                                                    fmt.Println(v13, fi16)
                                                    _ = v73
                                                } else {
                                                    fmt.Println(v14)
                                                    var v74 int8 = 86
                                                    if (false || (127 > 127)) {
                                                        v75 := GenS8{F0: "a"}
                                                        v74 -= 121
                                                        _ = v75
                                                    } else {
                                                        var v76 uint64 = 583
                                                        _ = v76
                                                    }
                                                    _ = v74
                                                }
                                            }
                                        case "hello":
                                            v77 := GenS2{F0: 356, F1: 1}
                                            _ = v77
                                        default:
                                            v13 += v13
                                            var v78 int16 = -32768
                                            v13 += 219
                                            _ = v78
                                        }
                                        switch 1 {
                                        case 32767:
                                            if (345 >= 18446744073709551615) {
                                                v79 := GenS8{F0: "\n\t"}
                                                var v80 bool = (false || (8 > 680))
                                                _ = v79
                                                _ = v80
                                            }
                                        }
                                        _ = v69
                                    }
                                }
                            }
                        default:
                            var v81 int = (-(v14))
                            _ = v81
                        }
                        var v82 uint8 = 159
                        _ = v82
                    }
                case -210:
                    v14 = 127
                case 42:
                    v83 := []int{fi15, fi16, 32768}
                    _ = v83
                default:
                    var v84 int = -338
                    _ = v84
                }
            } else {
                v13 -= (v13 << ((42) & 63))
            }
            fmt.Println(fi16)
            _ = fi16
            _ = v17
            _ = v18
        }
        fmt.Println(v14, v13)
        var v85 uint64 = 109
        _ = fi15
        _ = v85
    }
    v86 := GenS7{F0: 1, F1: 920}
    v87 := []int{v14}
    _ = v13
    _ = v14
    _ = v86
    _ = v87
    return 4294967295
}

func main() {
    var v1 int8 = 17
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]string)
    v1 = v1
    fmt.Println(v2, v2)
    v1 = (-(v1))
    if (0 == 4294967295) {
        var v88 uint32 = 42
        for fi89 := 0; fi89 < 4; fi89++ {
            for fi90 := 0; fi90 < 4; fi90++ {
                v91 := []int{(fi89 & fi89), fi90, (fi89 ^ (fi89 % ((-2147483648) | 1))), fi89}
                _ = fi90
                _ = v91
            }
            _ = fi89
        }
        _ = v88
    } else {
        v92 := GenS9{F0: 42}
        v93 := []int{127}
        if false {
            for fi94 := 0; fi94 < 17; fi94++ {
                for fi95 := 0; fi95 < 1; fi95++ {
                    if true {
                        v96 := GenS3{F0: 2147483648, F1: 1}
                        v2 += v2
                        _ = v96
                    }
                    switch "" {
                    case "\n\t":
                        _ = len(v3)
                    case "foo bar":
                        v2 -= (v2 + v2)
                    }
                    if (0 >= genFn0()) {
                        fmt.Println(fi94, v2)
                        v1 = (v1 << ((4294967295) & 63))
                    } else {
                        var v97 float64 = float64(1.5)
                        v1 -= (v1 | -1)
                        _ = v97
                    }
                    _ = fi95
                }
                v1 = v1
                _ = fi94
            }
            var v98 uint64 = 0
            _ = v98
        } else {
            for fi99 := 0; fi99 < 6; fi99++ {
                for fi100 := 0; fi100 < 12; fi100++ {
                    if true {
                        fmt.Println(fi99, v2)
                    } else {
                        var v101 int64 = (v2 & 42)
                        v93 = append(v93, -117)
                        _ = v101
                    }
                    _ = fi100
                }
                _ = fi99
            }
        }
        _ = v92
        _ = v93
    }
    for fi102 := 0; fi102 < 10; fi102++ {
        var v103 int16 = 1
        if ("" <= "foo bar") {
            var v104 string = "\n\t"
            for fi105 := 0; fi105 < 6; fi105++ {
                if true {
                    fmt.Println(fi102, v104)
                    v106 := GenS8{F0: "hello"}
                    _ = v106
                } else {
                    switch (-(fi102)) {
                    case 331:
                        var v107 int64 = 9223372036854775807
                        _ = v107
                    case 645:
                        v103 += (v103 / (((v103 & v103)) | 1))
                    case 42:
                        fmt.Println(v104, v1, v103)
                    default:
                        v103 = (v103 << ((533) & 63))
                    }
                    v108 := GenS1{F0: 2, F1: 128}
                    _ = v108
                }
                v1 -= v1
                v2 = (v2 - -423)
                var v109 float64 = 0.1
                _ = fi105
                _ = v109
            }
            _ = v104
        } else {
            if false {
                v2 = (^(v2))
                for fi110 := 0; fi110 < 3; fi110++ {
                    fmt.Println(v1)
                    if false {
                        v111 := GenS6{F0: (v103 & 42)}
                        v112 := GenS10{F0: 722, F1: 42}
                        _ = v111
                        _ = v112
                    }
                    _ = fi110
                }
                v1 = v1
            }
            var v113 uint8 = 2
            v114 := GenS11{F0: 1, F1: v2, F2: 421}
            _ = v113
            _ = v114
        }
        switch fi102 {
        case 1:
            v3[fi102] = ("\n\t" + "")
        default:
            v115 := GenS11{F0: (v1 << ((4294967295) & 63)), F1: 42, F2: 18446744073709551615}
            _ = v115
        }
        _ = fi102
        _ = v103
    }
    var v116 float32 = -2.25
    if (-370 == 128) {
        var v117 uint = 0
        _ = v117
    } else {
        switch 42 {
        case 4294967295:
            for fi118 := 0; fi118 < 9; fi118++ {
                var v119 uint8 = 182
                v120 := []int{fi118, fi118, 2147483647, -194}
                _ = fi118
                _ = v119
                _ = v120
            }
        case 42:
            v2 = v2
        case 1:
            v1 += v1
        }
    }
    if true {
        v116 += v116
    }
    fmt.Println(v116, v116, v116)
    switch "hello" {
    case "a":
        switch "\n\t" {
        case "\n\t":
            if (!(!false)) {
                v116 = 0.1
                var v121 bool = true
                v122 := GenS4{F0: 42}
                _ = v121
                _ = v122
            }
        case "":
            var v123 bool = true
            _ = v123
        }
    }
    v1 -= (v1 | -15)
    fmt.Println(v116, v116, v2)
    var v124 int32 = -2147483648
    v125 := []int{-1, -266, -18}
    v116 += v116
    for fi126 := 0; fi126 < 20; fi126++ {
        v1 = -39
        v124 += 1
        _ = fi126
    }
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("v124:", v124)
    fmt.Println("slicelen v125:", len(v125), v125)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v116
    _ = v124
    _ = v125
    _ = v2
    _ = v3
}
