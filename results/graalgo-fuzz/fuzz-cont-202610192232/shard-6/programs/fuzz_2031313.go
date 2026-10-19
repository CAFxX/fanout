package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int; F2 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 uint8; F2 string }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS8 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS9 struct { F0 string; F1 int; F2 int16 }
func genFn0() uint {
    if ((0.0 - 0.1) == 1e10) {
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 int8 = 18
            v8 += 4
            var v9 bool = true
            _ = fi7
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 int64 = -824
            var v12 float64 = 0.1
            _ = fi10
            _ = v11
            _ = v12
        }
        v13 := []int{4294967295, 32768, -94}
        _ = v13
    } else {
        v14 := GenS1{F0: 231, F1: 215, F2: 4294967295}
        _ = v14
    }
    v15 := []int{0}
    _ = v15
    return 651
}
func genFn1() int32 {
    if (106 != 59) {
        if false {
            for fi16 := 0; fi16 < 14; fi16++ {
                for fi17 := 0; fi17 < 14; fi17++ {
                    fmt.Println(fi16)
                    var v18 uint = 65535
                    switch fi16 {
                    case 2147483647:
                        v19 := GenS1{F0: 1, F1: 42, F2: (v18 / ((v18) | 1))}
                        _ = v19
                    }
                    _ = fi17
                    _ = v18
                }
                var v20 int32 = 2147483647
                var v21 uint16 = 1
                _ = fi16
                _ = v20
                _ = v21
            }
        } else {
            v22 := GenS1{F0: 255, F1: 255, F2: 170}
            var v23 int = -9223372036854775808
            v24 := GenS1{F0: 138, F1: 1, F2: 0}
            _ = v22
            _ = v23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        var v26 int32 = 1
        _ = fi25
        _ = v26
    }
    return -527
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = float32((1.5 + 1.5))
    v3 := []int{}
    v4 := make(map[int]string)
    v5 := make(map[uint32]string)
    var v6 uint64 = 18446744073709551615
    if (107 < 1) {
        var v27 bool = (!false)
        _ = v27
    }
    fmt.Println(v6, v2)
    fmt.Println(v6, v6)
    if true {
        var v28 int = 32767
        if true {
            switch "a" {
            case "hello":
                v3 = append(v3, v28)
            default:
                fmt.Println(v28, v28, v28)
            }
        } else {
            for fi29 := 0; fi29 < 4; fi29++ {
                for fi30 := 0; fi30 < 16; fi30++ {
                    _ = len(v4)
                    switch "foo bar" {
                    case "\n\t":
                        v28 = fi30
                    case "a":
                        v31 := v1[446]
                        _ = v31
                    default:
                        v4[fi30] = ("" + "\n\t")
                    }
                    _ = fi30
                }
                v3 = append(v3, (fi29 | v28))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi29
            }
            if false {
                v32 := GenS1{F0: 255, F1: 128, F2: 255}
                if (116 >= 255) {
                    for fi33 := 0; fi33 < 11; fi33++ {
                        v34 := GenS1{F0: 128, F1: 0, F2: 2}
                        delete(v1, 65535)
                        v2 += v2
                        _ = fi33
                        _ = v34
                    }
                    v4[(v28 | (v28 >> ((18446744073709551615) & 63)))] = ((("" + "hello") + ((("foo bar" + (("hello" + "") + "")) + "\n\t") + ("" + (("a" + ("\n\t" + "")) + "")))) + ("a" + "foo bar"))
                }
                switch (("" + "") + ("foo bar" + ("a" + "hello"))) {
                case "a":
                    v2 += (float32(((1e10 * 3.14159) + ((((0.1 / 3.14159) / 1e10) - 0.0) - ((-2.25 - float64(0.1)) + 1e10)))) / v2)
                case "hello":
                    if true {
                        var v35 uint32 = 4294967295
                        var v36 uint32 = v35
                        for fi37 := 0; fi37 < 14; fi37++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (("a" < "a") && false) {
                                for fi38 := 0; fi38 < 3; fi38++ {
                                    var v39 uint32 = (v35 % (((^(v35))) | 1))
                                    if false {
                                        if ((!(128 > uint8(v39))) && true) {
                                            var v40 float32 = v2
                                            v36 = (v39 % ((1) | 1))
                                            if ((828 > v36) && (100 <= 6)) {
                                                v41 := GenS2{F0: 32767}
                                                _ = v41
                                            }
                                            if (1e10 < 0.1) {
                                                v42 := GenS2{F0: 32767}
                                                switch ("\n\t" + "\n\t") {
                                                case "\n\t":
                                                    fmt.Println(v6, v6)
                                                case "hello":
                                                    if true {
                                                        var v43 float32 = (((1e10 * ((((v2 - v40) - float32(0.1)) * 0.1) * (v2 * ((-2.25 / -2.25) / v2)))) / v2) - float32(0.0))
                                                        var v44 int = (v28 | -9223372036854775808)
                                                        _ = v43
                                                        _ = v44
                                                    } else {
                                                        switch (fi37 % (((^(fi37))) | 1)) {
                                                        case -674:
                                                            v28 = fi37
                                                        }
                                                        v45 := GenS3{F0: 42, F1: (v28 | (v28 & (fi37 / ((fi37) | 1)))), F2: genFn0()}
                                                        _ = v45
                                                    }
                                                default:
                                                    var v46 int = 957
                                                    _ = v46
                                                }
                                                _ = v42
                                            }
                                            _ = v40
                                        }
                                    } else {
                                        if true {
                                            var v47 int64 = 441
                                            v4[32767] = "foo bar"
                                            _ = v47
                                        } else {
                                            if true {
                                                v48 := GenS1{F0: 0, F1: 42, F2: 423}
                                                var v49 string = "foo bar"
                                                switch 9223372036854775807 {
                                                case -9223372036854775808:
                                                    var v50 int16 = 42
                                                    var v51 string = ("" + ((v49 + "foo bar") + ("foo bar" + (v49 + "hello"))))
                                                    fmt.Println(v49, v50)
                                                    v52 := GenS2{F0: int16(v36)}
                                                    _ = v50
                                                    _ = v51
                                                    _ = v52
                                                case -1:
                                                    if (false && (!false)) {
                                                        if true {
                                                            fmt.Println(v2, v28)
                                                            switch v49 {
                                                            case "hello":
                                                                _ = len(v1)
                                                            case "foo bar":
                                                                _ = len(v3)
                                                            case "\n\t":
                                                                switch v49 {
                                                                case "hello":
                                                                    v53 := v1[296]
                                                                    _ = v53
                                                                case "":
                                                                    if true {
                                                                        v39 = 0
                                                                    } else {
                                                                        v36 = 2
                                                                        if (-41 < -1) {
                                                                            v54 := GenS3{F0: -123, F1: fi38, F2: 53}
                                                                            _ = v54
                                                                        } else {
                                                                            v49 = v49
                                                                            var v55 uint32 = (v39 >> ((65535) & 63))
                                                                            v55 -= v36
                                                                            _ = v55
                                                                        }
                                                                        v6 -= (^(v6))
                                                                    }
                                                                case "a":
                                                                    if true {
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            v56 := GenS4{F0: 795}
                                                                            _ = v56
                                                                        default:
                                                                            if (true || true) {
                                                                                v28 = fi38
                                                                            }
                                                                        }
                                                                        v57 := v4[(fi37 % ((v28) | 1))]
                                                                        _ = v57
                                                                    } else {
                                                                        if ("a" != (v49 + v49)) {
                                                                            if true {
                                                                                v3 = append(v3, 4294967295)
                                                                                v36 += (v35 >> ((4294967295) & 63))
                                                                                v58 := GenS4{F0: 1}
                                                                                _ = v58
                                                                            } else {
                                                                                v59 := GenS3{F0: -9223372036854775808, F1: (^(fi37)), F2: 638}
                                                                                switch ((v49 + ("" + ("a" + (("hello" + "hello") + "a")))) + ("\n\t" + v49)) {
                                                                                case "hello":
                                                                                    v36 = (v39 << ((genFn0()) & 63))
                                                                                default:
                                                                                    var v60 uint16 = 680
                                                                                    _ = v60
                                                                                }
                                                                                _ = v59
                                                                            }
                                                                            switch (^(v28)) {
                                                                            case 0:
                                                                                var v61 float32 = v2
                                                                                _ = v61
                                                                            case 1:
                                                                                if true {
                                                                                    if (!true) {
                                                                                        var v62 uint64 = v6
                                                                                        if (true && true) {
                                                                                            v4[v28] = "hello"
                                                                                            switch "\n\t" {
                                                                                            case "hello":
                                                                                                var v63 float64 = 0.1
                                                                                                _ = v63
                                                                                            case "a":
                                                                                                if false {
                                                                                                    if true {
                                                                                                        v36 += (v35 & (v36 % (((v35 + (v35 | (^(v35))))) | 1)))
                                                                                                        v6 = v6
                                                                                                        fmt.Println(fi38)
                                                                                                    }
                                                                                                    if (!false) {
                                                                                                        fmt.Println(fi38, v62, v39)
                                                                                                        fmt.Println(fi38)
                                                                                                    } else {
                                                                                                        var v64 float32 = v2
                                                                                                        v65 := GenS2{F0: 1000}
                                                                                                        _ = v64
                                                                                                        _ = v65
                                                                                                    }
                                                                                                } else {
                                                                                                    v66 := GenS5{F0: (("a" + "") + v49), F1: 1, F2: "foo bar"}
                                                                                                    _ = v66
                                                                                                }
                                                                                            default:
                                                                                                var v67 bool = (-18 > 1)
                                                                                                _ = v67
                                                                                            }
                                                                                            fmt.Println(v36)
                                                                                        } else {
                                                                                            var v68 float32 = 1.5
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            v49 = "hello"
                                                                                            _ = v68
                                                                                        }
                                                                                        _ = v62
                                                                                    } else {
                                                                                        var v69 int8 = -128
                                                                                        v70 := GenS1{F0: 199, F1: 1, F2: 478}
                                                                                        _ = v69
                                                                                        _ = v70
                                                                                    }
                                                                                    v49 = ""
                                                                                }
                                                                            case 128:
                                                                                var v71 int = (v28 | -1)
                                                                                _ = v71
                                                                            }
                                                                            v72 := GenS3{F0: 786, F1: 2147483647, F2: 370}
                                                                            _ = v72
                                                                        } else {
                                                                            v36 -= v39
                                                                            if false {
                                                                                var v73 uint16 = 32767
                                                                                var v74 float64 = (3.14159 + 1.5)
                                                                                if false {
                                                                                    var v75 bool = false
                                                                                    _ = v75
                                                                                } else {
                                                                                    if true {
                                                                                        if (-433 <= 700) {
                                                                                            _ = len(v3)
                                                                                            v28 = fi38
                                                                                        } else {
                                                                                            v76 := GenS3{F0: -9223372036854775808, F1: int(v35), F2: 2}
                                                                                            v77 := GenS6{F0: v39, F1: 331}
                                                                                            _ = v76
                                                                                            _ = v77
                                                                                        }
                                                                                    } else {
                                                                                        var v78 uint16 = (v73 % ((v73) | 1))
                                                                                        _ = v78
                                                                                    }
                                                                                }
                                                                                _ = v73
                                                                                _ = v74
                                                                            }
                                                                            v79 := GenS7{F0: 1000, F1: 1, F2: (v35 >> ((0) & 63))}
                                                                            _ = v79
                                                                        }
                                                                        v80 := GenS5{F0: v49, F1: 128, F2: v49}
                                                                        switch v49 {
                                                                        case "a":
                                                                            var v81 uint16 = 1
                                                                            _ = v81
                                                                        case "\n\t":
                                                                            if len(v3) > 0 {
                                                                                v3[0] = (v28 >> ((255) & 63))
                                                                            }
                                                                        case "hello":
                                                                            v49 = ("" + ("" + (v49 + "foo bar")))
                                                                        default:
                                                                            if false {
                                                                                if true {
                                                                                    if false {
                                                                                        if true {
                                                                                            v82 := v5[42]
                                                                                            var v83 int32 = 1
                                                                                            switch (-(v28)) {
                                                                                            case -966:
                                                                                                if (uint32(v6) == v35) {
                                                                                                    fmt.Println(v6)
                                                                                                }
                                                                                            case 0:
                                                                                                v3 = append(v3, (fi37 | 32768))
                                                                                            default:
                                                                                                switch v82 {
                                                                                                case "":
                                                                                                    if true {
                                                                                                        v3 = append(v3, 342)
                                                                                                        var v84 int = (-(fi37))
                                                                                                        if true {
                                                                                                            if true {
                                                                                                                var v85 uint16 = 42
                                                                                                                var v86 int32 = (-(v83))
                                                                                                                _ = v85
                                                                                                                _ = v86
                                                                                                            }
                                                                                                            switch ((v49 + ("hello" + v49)) + (((v49 + v82) + v49) + v82)) {
                                                                                                            case "\n\t":
                                                                                                                var v87 uint32 = (v36 & (v35 * v36))
                                                                                                                _ = v87
                                                                                                            default:
                                                                                                                var v88 float32 = (-2.25 / 0.1)
                                                                                                                _ = v88
                                                                                                            }
                                                                                                        } else {
                                                                                                            v36 -= 4294967295
                                                                                                            v3 = append(v3, (v28 * fi37))
                                                                                                        }
                                                                                                        fmt.Println(v36, v84, v6)
                                                                                                        v84 = 4294967295
                                                                                                        _ = v84
                                                                                                    } else {
                                                                                                        var v89 uint16 = 305
                                                                                                        _ = len(v3)
                                                                                                        v90 := GenS3{F0: 9223372036854775807, F1: 1, F2: 2}
                                                                                                        switch ("a" + "a") {
                                                                                                        case "\n\t":
                                                                                                            v91 := GenS6{F0: 568, F1: 307}
                                                                                                            _ = v91
                                                                                                        case "hello":
                                                                                                            if true {
                                                                                                                switch v28 {
                                                                                                                case 127:
                                                                                                                    if false {
                                                                                                                        var v92 uint8 = 255
                                                                                                                        delete(v4, -1)
                                                                                                                        _ = v92
                                                                                                                    }
                                                                                                                case 32767:
                                                                                                                    v49 = "hello"
                                                                                                                    var v93 int64 = 0
                                                                                                                    v83 += (v83 / (((v83 << ((0) & 63))) | 1))
                                                                                                                    var v94 int32 = (v83 | (v83 >> ((4294967295) & 63)))
                                                                                                                    _ = v93
                                                                                                                    _ = v94
                                                                                                                default:
                                                                                                                    fmt.Println(v36, v82, v2)
                                                                                                                }
                                                                                                                var v95 int = 718
                                                                                                                if true {
                                                                                                                    v82 = v49
                                                                                                                } else {
                                                                                                                    v49 = "a"
                                                                                                                    if len(v3) > 0 {
                                                                                                                        v3[0] = -64
                                                                                                                    }
                                                                                                                    v3 = append(v3, 167)
                                                                                                                    var v96 uint = 4294967295
                                                                                                                    _ = v96
                                                                                                                }
                                                                                                                _ = v95
                                                                                                            }
                                                                                                        }
                                                                                                        fmt.Println(v36, v2)
                                                                                                        _ = v89
                                                                                                        _ = v90
                                                                                                    }
                                                                                                case "hello":
                                                                                                    switch (v28 | (fi37 * fi38)) {
                                                                                                    case 2147483647:
                                                                                                        var v97 int16 = 1
                                                                                                        _ = v97
                                                                                                    case 32767:
                                                                                                        if true {
                                                                                                            switch fi37 {
                                                                                                            case -2147483648:
                                                                                                                fmt.Println(fi37, v6, v82)
                                                                                                            }
                                                                                                        } else {
                                                                                                            fmt.Println(v28)
                                                                                                            fmt.Println(v39, v28)
                                                                                                            _ = len(v3)
                                                                                                        }
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    switch (^(fi37)) {
                                                                                                    case 32767:
                                                                                                        var v98 uint16 = 546
                                                                                                        _ = v98
                                                                                                    default:
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    var v99 bool = (!(42 != 100))
                                                                                                    _ = v99
                                                                                                }
                                                                                            }
                                                                                            _ = v82
                                                                                            _ = v83
                                                                                        } else {
                                                                                            var v100 uint = 18446744073709551615
                                                                                            var v101 uint64 = 18446744073709551615
                                                                                            _ = v100
                                                                                            _ = v101
                                                                                        }
                                                                                        var v102 string = v49
                                                                                        if false {
                                                                                            v103 := GenS8{F0: 2147483647, F1: 472, F2: 84}
                                                                                            if true {
                                                                                                v49 = v49
                                                                                                v2 -= (3.14159 * (v2 / v2))
                                                                                                var v104 float32 = float32(1e10)
                                                                                                _ = v104
                                                                                            } else {
                                                                                                v36 += 42
                                                                                                v105 := GenS4{F0: (v6 + (v6 & uint64(v36)))}
                                                                                                _ = v105
                                                                                            }
                                                                                            v106 := GenS8{F0: 1, F1: (v39 >> ((750) & 63)), F2: 255}
                                                                                            _ = v103
                                                                                            _ = v106
                                                                                        }
                                                                                        _ = v102
                                                                                    }
                                                                                    v107 := GenS9{F0: v49, F1: fi37, F2: -95}
                                                                                    _ = v107
                                                                                } else {
                                                                                    var v108 int8 = 42
                                                                                    v109 := GenS9{F0: v49, F1: (v28 | (fi37 - (^(v28)))), F2: -176}
                                                                                    _ = v108
                                                                                    _ = v109
                                                                                }
                                                                                if (0.1 > float32(3.14159)) {
                                                                                    fmt.Println(v35, fi38)
                                                                                    v110 := GenS3{F0: -634, F1: (fi38 / (((fi38 ^ (fi37 + (fi37 >> ((genFn0()) & 63))))) | 1)), F2: 4294967295}
                                                                                    v111 := GenS3{F0: 9223372036854775807, F1: (fi37 & fi37), F2: 975}
                                                                                    _ = v110
                                                                                    _ = v111
                                                                                }
                                                                                v112 := GenS2{F0: -32768}
                                                                                _ = v112
                                                                            }
                                                                            v3 = append(v3, (v28 << ((2) & 63)))
                                                                        }
                                                                        v113 := GenS7{F0: -370, F1: v6, F2: 42}
                                                                        _ = v113
                                                                        _ = v80
                                                                    }
                                                                default:
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        fmt.Println(v28)
                                                                    }
                                                                }
                                                            default:
                                                                if false {
                                                                    var v114 int8 = 100
                                                                    _ = v114
                                                                }
                                                                var v115 int8 = 68
                                                                v115 = (v115 ^ (^(v115)))
                                                                _ = v115
                                                            }
                                                            var v116 uint64 = (v6 ^ v6)
                                                            var v117 uint8 = 0
                                                            switch v49 {
                                                            case "hello":
                                                                var v118 uint32 = (v39 + v39)
                                                                _ = v118
                                                            default:
                                                                v116 = 113
                                                            }
                                                            if true {
                                                                var v119 float64 = (3.14159 + float64(v2))
                                                                v5[(v36 | (v36 | v36))] = "foo bar"
                                                                var v120 float32 = float32(0.1)
                                                                var v121 uint8 = v117
                                                                var v122 bool = (512 < genFn0())
                                                                _ = v119
                                                                _ = v120
                                                                _ = v121
                                                                _ = v122
                                                            } else {
                                                                v39 += (v36 % (((v35 + v36)) | 1))
                                                            }
                                                            _ = v116
                                                            _ = v117
                                                        }
                                                    }
                                                default:
                                                    var v123 uint64 = v6
                                                    _ = v123
                                                }
                                                _ = v48
                                                _ = v49
                                            }
                                        }
                                    }
                                    var v124 string = ("" + "hello")
                                    _ = fi38
                                    _ = v124
                                    _ = v39
                                }
                            } else {
                                var v125 bool = false
                                var v126 uint16 = 902
                                _ = v125
                                _ = v126
                            }
                            _ = fi37
                        }
                        _ = v35
                        _ = v36
                    } else {
                        switch "\n\t" {
                        case "a":
                            v28 += v28
                        case "foo bar":
                            v127 := GenS7{F0: 42, F1: v6, F2: 392}
                            _ = v127
                        case "\n\t":
                            var v128 int = 9223372036854775807
                            _ = v128
                        }
                        v6 = (v6 << ((42) & 63))
                        for fi129 := 0; fi129 < 2; fi129++ {
                            var v130 int64 = 42
                            for fi131 := 0; fi131 < 19; fi131++ {
                                var v132 uint = 42
                                _ = fi131
                                _ = v132
                            }
                            _ = fi129
                            _ = v130
                        }
                    }
                }
                _ = v32
            } else {
                var v133 int16 = 422
                var v134 int8 = 0
                var v135 uint64 = v6
                _ = v133
                _ = v134
                _ = v135
            }
        }
        if (0.1 < 1e10) {
            v136 := v5[492]
            fmt.Println(v2)
            var v137 int8 = 1
            _ = v136
            _ = v137
        } else {
            v138 := GenS1{F0: 255, F1: 177, F2: 42}
            _ = v138
        }
        _ = v28
    } else {
        v139 := v4[4294967295]
        v139 = v139
        _ = v139
    }
    v6 += (^(v6))
    var v140 int64 = 0
    if (!true) {
        var v141 uint32 = 1
        var v142 int = 128
        _ = v141
        _ = v142
    }
    switch 42 {
    case -2147483648:
        for fi143 := 0; fi143 < 11; fi143++ {
            if false {
                var v144 int64 = v140
                fmt.Println(v2)
                var v145 int8 = -128
                _ = v144
                _ = v145
            } else {
                var v146 float64 = 0.1
                switch 9223372036854775807 {
                case 340:
                    var v147 uint8 = 127
                    _ = v147
                case 128:
                    var v148 uint = 4294967295
                    _ = v148
                case 179:
                    var v149 int16 = 0
                    _ = v149
                }
                if true {
                    fmt.Println(fi143, v146)
                    v146 -= (v146 - (v146 / (1e10 / ((v146 * 1e10) / -2.25))))
                    for fi150 := 0; fi150 < 2; fi150++ {
                        v140 = v140
                        v4[(fi143 | (-(fi150)))] = "a"
                        _ = fi150
                    }
                } else {
                    for fi151 := 0; fi151 < 10; fi151++ {
                        v140 = v140
                        _ = fi151
                    }
                    v140 -= 1
                }
                _ = v146
            }
            _ = fi143
        }
    case 1:
        if (0 == 1) {
            switch 9223372036854775807 {
            case 32768:
                for fi152 := 0; fi152 < 19; fi152++ {
                    if len(v3) > 0 {
                        v3[0] = fi152
                    }
                    var v153 string = ("hello" + "foo bar")
                    fmt.Println(v2)
                    _ = fi152
                    _ = v153
                }
            }
            if (false || (1 >= -1)) {
                if false {
                    var v154 float32 = 3.14159
                    v140 = (-(v140))
                    _ = v154
                } else {
                    v3 = append(v3, -2147483648)
                    var v155 int = -9223372036854775808
                    _ = v155
                }
                switch "a" {
                case "\n\t":
                    if len(v3) > 0 {
                        v3[0] = 9223372036854775807
                    }
                case "hello":
                    var v156 uint64 = (v6 >> ((18446744073709551615) & 63))
                    _ = v156
                }
                fmt.Println(v2)
            }
        }
    case 42:
        for fi157 := 0; fi157 < 19; fi157++ {
            switch "\n\t" {
            case "foo bar":
                delete(v1, 151)
            case "a":
                v6 = v6
            default:
                if false {
                    v6 = (^(v6))
                    for fi158 := 0; fi158 < 11; fi158++ {
                        if false {
                            switch (("a" + "a") + "foo bar") {
                            case "hello":
                                fmt.Println(v2, fi158)
                            }
                            delete(v1, 0)
                            var v159 int8 = 100
                            _ = v159
                        } else {
                            var v160 uint16 = 487
                            var v161 int8 = -1
                            v161 = v161
                            _ = v160
                            _ = v161
                        }
                        _ = fi158
                    }
                }
            }
            _ = fi157
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v140:", v140)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v140
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
