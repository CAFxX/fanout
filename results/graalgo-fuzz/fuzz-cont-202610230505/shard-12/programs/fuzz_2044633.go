package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int64; F1 int32; F2 int16 }
type GenS4 struct { F0 uint32; F1 int64; F2 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS8 struct { F0 uint64; F1 string; F2 uint16 }
type GenS9 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS10 struct { F0 uint64; F1 uint8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int16; F1 int16; F2 int8 }
type GenS13 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint) int8 {
    var v7 bool = true
    v8 := []int{4294967295}
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint16 = 725
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return -128
}
func genFn1(p0 uint8) int64 {
    var v11 int64 = 0
    v12 := GenS1{F0: v11, F1: 0}
    var v13 float32 = float32(0.0)
    if true {
        v14 := GenS1{F0: 1, F1: p0}
        _ = v14
    } else {
        v13 = (v13 + 0.1)
    }
    if true {
        if true {
            for fi15 := 0; fi15 < 9; fi15++ {
                if (105 > genFn0(18446744073709551615)) {
                    var v16 int16 = 874
                    var v17 bool = false
                    fmt.Println(v11)
                    _ = v16
                    _ = v17
                }
                var v18 int32 = 2147483647
                var v19 bool = (true && true)
                _ = fi15
                _ = v18
                _ = v19
            }
        } else {
            v13 += float32(0.0)
            v20 := GenS1{F0: (v11 % ((v11) | 1)), F1: (p0 / (((p0 | 13)) | 1))}
            for fi21 := 0; fi21 < 10; fi21++ {
                fmt.Println(fi21, p0)
                for fi22 := 0; fi22 < 15; fi22++ {
                    v23 := []int{(fi22 ^ (fi21 << ((456) & 63))), (^(fi21))}
                    _ = fi22
                    _ = v23
                }
                var v24 uint64 = 0
                _ = fi21
                _ = v24
            }
            _ = v20
        }
        var v25 uint64 = 680
        v26 := []int{32767, 2147483647}
        _ = v25
        _ = v26
    } else {
        if false {
            var v27 string = ""
            var v28 int16 = -1
            switch 9223372036854775807 {
            case 910:
                for fi29 := 0; fi29 < 17; fi29++ {
                    v30 := []int{}
                    _ = fi29
                    _ = v30
                }
            case 1:
                if false {
                    var v31 uint32 = 963
                    _ = v31
                } else {
                    v32 := []int{128}
                    v11 = v11
                    if false {
                        switch 161 {
                        case 42:
                            if (!(!true)) {
                                var v33 float64 = 1e10
                                if len(v32) > 0 {
                                    v32[0] = 839
                                }
                                _ = v33
                            }
                        case -518:
                            v11 = (v11 * -590)
                        case 32768:
                            switch 9223372036854775807 {
                            case 355:
                                for fi34 := 0; fi34 < 6; fi34++ {
                                    if false {
                                        fmt.Println(fi34)
                                    } else {
                                        v35 := GenS1{F0: (^(v11)), F1: 42}
                                        _ = v35
                                    }
                                    _ = fi34
                                }
                            }
                        }
                    } else {
                        for fi36 := 0; fi36 < 15; fi36++ {
                            for fi37 := 0; fi37 < 8; fi37++ {
                                var v38 uint8 = (p0 - (^(p0)))
                                v39 := GenS1{F0: v11, F1: 178}
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                            v40 := GenS1{F0: v11, F1: p0}
                            _ = fi36
                            _ = v40
                        }
                        for fi41 := 0; fi41 < 19; fi41++ {
                            fmt.Println(v13, v11)
                            if ((float32(0.1) + v13) == v13) {
                                var v42 int16 = 1000
                                _ = v42
                            } else {
                                if false {
                                    for fi43 := 0; fi43 < 12; fi43++ {
                                        var v44 int16 = v28
                                        var v45 uint32 = 65535
                                        _ = fi43
                                        _ = v44
                                        _ = v45
                                    }
                                    if len(v32) >= 2 {
                                        v32 = v32[1:len(v32)]
                                    }
                                }
                                switch (fi41 | -148) {
                                case 42:
                                    var v46 float32 = float32(3.14159)
                                    _ = v46
                                }
                                v47 := GenS2{F0: v28, F1: (v11 ^ (v11 - int64(p0)))}
                                _ = v47
                            }
                            v11 -= 453
                            _ = fi41
                        }
                    }
                    _ = v32
                }
            default:
                if false {
                    v48 := GenS1{F0: v11, F1: p0}
                    fmt.Println(v11)
                    v28 -= 42
                    _ = v48
                } else {
                    v49 := GenS3{F0: v11, F1: 0, F2: v28}
                    fmt.Println(p0, v13)
                    for fi50 := 0; fi50 < 14; fi50++ {
                        v11 += (v11 + (v11 + (^(v11))))
                        _ = fi50
                    }
                    _ = v49
                }
            }
            _ = v27
            _ = v28
        }
        v51 := GenS4{F0: 2, F1: v11, F2: -1}
        v52 := GenS4{F0: 2, F1: v11, F2: -7}
        _ = v51
        _ = v52
    }
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return (v11 & (v11 % ((v11) | 1)))
}

func main() {
    var v1 int64 = 42
    var v2 string = "hello"
    var v3 int16 = 0
    v4 := []int{128, 4294967295, -1}
    var v5 int8 = 55
    var v6 uint8 = 2
    if (genFn0(42) < 127) {
        var v53 uint8 = (v6 ^ 42)
        for fi54 := 0; fi54 < 2; fi54++ {
            if len(v4) > 0 {
                v4[0] = (fi54 + fi54)
            }
            switch (-(fi54)) {
            case 959:
                if (0 != 528) {
                    v4 = append(v4, fi54)
                    v3 += v3
                } else {
                    if (true || (!(uint16(v6) <= 869))) {
                        if true {
                            fmt.Println(v53, fi54, v53)
                        } else {
                            if ((("" + "") + (v2 + (v2 + (v2 + (v2 + "\n\t"))))) > "foo bar") {
                                fmt.Println(v53)
                                if (v2 < ((v2 + "hello") + "foo bar")) {
                                    if true {
                                        v55 := GenS5{F0: 2}
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v55
                                    } else {
                                        if (2147483647 > -1) {
                                            v56 := GenS4{F0: 42, F1: v1, F2: v5}
                                            _ = v56
                                        } else {
                                            var v57 float32 = (0.0 - 0.1)
                                            _ = v57
                                        }
                                    }
                                    v2 = v2
                                    for fi58 := 0; fi58 < 8; fi58++ {
                                        if (false || true) {
                                            _ = len(v4)
                                            v1 -= genFn1((v53 >> ((2) & 63)))
                                            v59 := GenS6{F0: 417}
                                            if (1e10 != 1.5) {
                                                v60 := GenS6{F0: -756}
                                                _ = v60
                                            } else {
                                                var v61 int8 = 42
                                                var v62 int = (fi54 & 118)
                                                fmt.Println(fi58, v3)
                                                _ = v61
                                                _ = v62
                                            }
                                            _ = v59
                                        }
                                        var v63 float32 = (3.14159 * 1.5)
                                        _ = fi58
                                        _ = v63
                                    }
                                } else {
                                    var v64 uint8 = v6
                                    if (65535 != 648) {
                                        switch (v2 + (v2 + v2)) {
                                        case "a":
                                            var v65 int8 = v5
                                            _ = v65
                                        default:
                                            if (42 != 350) {
                                                for fi66 := 0; fi66 < 15; fi66++ {
                                                    v2 = ("a" + "a")
                                                    _ = fi66
                                                }
                                                switch v2 {
                                                case "":
                                                    v53 = (v6 - 42)
                                                }
                                            } else {
                                                switch v2 {
                                                case "":
                                                    var v67 string = v2
                                                    _ = v67
                                                case "a":
                                                    _ = len(v4)
                                                case "\n\t":
                                                    switch (^(fi54)) {
                                                    case 2147483647:
                                                        v2 = v2
                                                    }
                                                }
                                                v68 := GenS3{F0: v1, F1: -1, F2: v3}
                                                _ = v68
                                            }
                                        }
                                        for fi69 := 0; fi69 < 18; fi69++ {
                                            fmt.Println(v3, v1)
                                            v3 = v3
                                            v70 := GenS7{F0: 0, F1: -1, F2: 1}
                                            switch "a" {
                                            case "hello":
                                                v5 -= v5
                                            case "a":
                                                fmt.Println(v3, v3)
                                            case "\n\t":
                                                switch v2 {
                                                case "foo bar":
                                                    var v71 int64 = 42
                                                    _ = v71
                                                default:
                                                    v53 += (v64 * (v53 << ((321) & 63)))
                                                }
                                            default:
                                                var v72 int64 = -9223372036854775808
                                                _ = v72
                                            }
                                            v73 := GenS8{F0: 18446744073709551615, F1: v2, F2: 1}
                                            _ = fi69
                                            _ = v70
                                            _ = v73
                                        }
                                        if ((v1 << ((2) & 63)) == (v1 >> ((229) & 63))) {
                                            var v74 float32 = (float32(float64(1e10)) / ((((float32(0.1) - (0.1 + (0.0 + 0.1))) / 1e10) * -2.25) + (3.14159 - 0.1)))
                                            v4 = append(v4, fi54)
                                            switch (fi54 * (fi54 & fi54)) {
                                            case -670:
                                                v75 := GenS3{F0: genFn1((v53 >> ((582) & 63))), F1: 1, F2: (-(v3))}
                                                _ = v75
                                            case 9223372036854775807:
                                                v76 := GenS4{F0: 2147483648, F1: (v1 >> ((4) & 63)), F2: genFn0(255)}
                                                _ = v76
                                            default:
                                                if false {
                                                    var v77 uint64 = uint64(v53)
                                                    var v78 uint8 = (v64 >> ((65535) & 63))
                                                    _ = v77
                                                    _ = v78
                                                } else {
                                                    if (42 <= 81) {
                                                        var v79 int16 = (v3 | 32767)
                                                        _ = v79
                                                    } else {
                                                        for fi80 := 0; fi80 < 15; fi80++ {
                                                            var v81 uint32 = 77
                                                            _ = fi80
                                                            _ = v81
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v74
                                        } else {
                                            if (!true) {
                                                for fi82 := 0; fi82 < 2; fi82++ {
                                                    var v83 int64 = v1
                                                    v4 = append(v4, fi82)
                                                    switch v2 {
                                                    case "hello":
                                                        switch (fi54 | (fi54 & fi54)) {
                                                        case 9223372036854775807:
                                                            v3 = (v3 % ((-896) | 1))
                                                        case -2147483648:
                                                            var v84 int64 = int64(v5)
                                                            _ = v84
                                                        case 42:
                                                            var v85 float32 = 1.5
                                                            var v86 bool = false
                                                            var v87 int = fi82
                                                            v88 := GenS1{F0: genFn1(v64), F1: (^(v64))}
                                                            _ = v85
                                                            _ = v86
                                                            _ = v87
                                                            _ = v88
                                                        default:
                                                            v89 := GenS3{F0: (v1 & int64(v5)), F1: 1, F2: v3}
                                                            _ = v89
                                                        }
                                                    case "":
                                                        v1 = (v83 * 42)
                                                    case "\n\t":
                                                        v90 := GenS9{F0: 2, F1: 737, F2: 1}
                                                        _ = v90
                                                    default:
                                                        v53 += 234
                                                    }
                                                    _ = fi82
                                                    _ = v83
                                                }
                                                if true {
                                                    switch (v2 + "foo bar") {
                                                    case "hello":
                                                        v6 = (v53 + 2)
                                                    default:
                                                        v91 := GenS10{F0: 4294967295, F1: v53}
                                                        _ = v91
                                                    }
                                                    if true {
                                                        var v92 int32 = 572
                                                        v6 = v64
                                                        for fi93 := 0; fi93 < 17; fi93++ {
                                                            _ = len(v4)
                                                            v94 := GenS6{F0: (-(v92))}
                                                            v1 = v1
                                                            var v95 float32 = float32(1.5)
                                                            v96 := GenS7{F0: 42, F1: (v92 << ((0) & 63)), F2: 1}
                                                            _ = fi93
                                                            _ = v94
                                                            _ = v95
                                                            _ = v96
                                                        }
                                                        _ = v92
                                                    } else {
                                                        fmt.Println(v6, v1)
                                                        v53 += v6
                                                    }
                                                } else {
                                                    var v97 uint32 = 2147483647
                                                    v3 -= v3
                                                    if false {
                                                        v53 += (v53 + v64)
                                                    } else {
                                                        for fi98 := 0; fi98 < 17; fi98++ {
                                                            if (18446744073709551615 != 0) {
                                                                var v99 int = fi98
                                                                var v100 uint = 276
                                                                switch v2 {
                                                                case "foo bar":
                                                                    v1 = (v1 & v1)
                                                                case "a":
                                                                    if len(v4) > 0 {
                                                                        v4[0] = (fi98 ^ v99)
                                                                    }
                                                                }
                                                                _ = v100
                                                                _ = v99
                                                            } else {
                                                                switch (fi98 ^ 0) {
                                                                case 4294967295:
                                                                    var v101 float64 = 3.14159
                                                                    _ = v101
                                                                case 42:
                                                                    var v102 int8 = v5
                                                                    _ = v102
                                                                }
                                                            }
                                                            fmt.Println(v3, v2)
                                                            _ = fi98
                                                        }
                                                    }
                                                    _ = v97
                                                }
                                            }
                                        }
                                    } else {
                                        v64 -= (^(v6))
                                        _ = len(v4)
                                        var v103 uint8 = (v53 * (v6 % ((v53) | 1)))
                                        _ = v103
                                    }
                                    fmt.Println(v53, v64)
                                    _ = v64
                                }
                            } else {
                                for fi104 := 0; fi104 < 16; fi104++ {
                                    var v105 uint32 = 0
                                    _ = fi104
                                    _ = v105
                                }
                            }
                            fmt.Println(v1)
                            v4 = append(v4, fi54)
                        }
                        var v106 bool = ((42 < 230) && true)
                        _ = v106
                    }
                }
            case 662:
                v107 := GenS11{F0: 2147483647}
                _ = v107
            case 0:
                var v108 int8 = v5
                _ = v108
            default:
                for fi109 := 0; fi109 < 2; fi109++ {
                    var v110 int8 = v5
                    if false {
                        var v111 uint16 = 1
                        var v112 bool = true
                        if v112 {
                            v5 -= genFn0(4294967295)
                            var v113 int8 = v110
                            _ = v113
                        }
                        _ = v111
                        _ = v112
                    }
                    _ = fi109
                    _ = v110
                }
            }
            v3 -= 0
            _ = fi54
        }
        if false {
            switch v2 {
            case "":
                var v114 uint32 = 65
                _ = v114
            case "hello":
                var v115 int8 = (v5 % ((genFn0(0)) | 1))
                _ = v115
            case "\n\t":
                var v116 uint8 = v53
                _ = v116
            }
            if false {
                v2 = ("foo bar" + ("" + "a"))
            }
            v1 = v1
        } else {
            fmt.Println(v3, v6)
        }
        _ = v53
    }
    var v117 uint8 = 128
    v118 := GenS12{F0: 0, F1: v3, F2: v5}
    var v119 string = "\n\t"
    if true {
        var v120 uint = 255
        v3 += 1000
        _ = v120
    } else {
        fmt.Println(v2, v117, v1)
        v4 = append(v4, 4294967295)
    }
    v119 = v119
    if true {
        if (true && true) {
            var v121 uint8 = v117
            v4 = append(v4, -403)
            _ = v121
        } else {
            var v122 string = ("foo bar" + "a")
            _ = v122
        }
    }
    if len(v4) > 0 {
        v4[0] = -1
    }
    switch "" {
    case "":
        if (!true) {
            v123 := GenS13{F0: v5, F1: 2}
            _ = v123
        } else {
            var v124 uint16 = 42
            _ = v124
        }
    case "foo bar":
        var v125 int32 = 42
        _ = v125
    case "\n\t":
        for fi126 := 0; fi126 < 20; fi126++ {
            v127 := GenS9{F0: 65535, F1: 2147483647, F2: 1}
            for fi128 := 0; fi128 < 10; fi128++ {
                var v129 float64 = (float64(1.5) * float64((-2.25 - (0.0 - 1.5))))
                _ = fi128
                _ = v129
            }
            v1 -= 915
            _ = fi126
            _ = v127
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v119:", v119)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v117
    _ = v118
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
