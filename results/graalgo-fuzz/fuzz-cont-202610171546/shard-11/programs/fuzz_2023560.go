package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 bool }
type GenS2 struct { F0 int16; F1 bool; F2 uint64 }
type GenS3 struct { F0 string; F1 bool; F2 int }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 uint8; F2 int }
type GenS8 struct { F0 int8; F1 uint8 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int; F1 int32; F2 int64 }
type GenS12 struct { F0 int16; F1 uint }
type GenS13 struct { F0 uint8; F1 int16 }
type GenS14 struct { F0 int16; F1 string; F2 uint8 }
type GenS15 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS16 struct { F0 int16; F1 bool }

func main() {
    var v1 uint16 = 2
    var v2 int32 = 0
    var v3 float32 = (((0.1 - 0.1) + -2.25) - (3.14159 * 1e10))
    var v4 int64 = -205
    v5 := []int{-1000}
    var v6 float32 = 0.1
    v1 = 0
    var v7 uint16 = 95
    v2 = -74
    var v8 int32 = (v2 ^ 42)
    var v9 uint16 = (v1 >> ((0) & 63))
    var v10 uint16 = 1
    switch -2147483648 {
    case -1:
        var v11 uint32 = 65535
        _ = v11
    default:
        fmt.Println(v9, v1, v7)
    }
    v12 := GenS1{F0: v2, F1: 42, F2: (true || false)}
    var v13 float32 = ((float32((float64(1.5) + 1.5)) / v3) / v3)
    v14 := GenS2{F0: 1000, F1: (v8 == (v8 + (^(v8)))), F2: 18446744073709551615}
    var v15 bool = (216 != 253)
    if (255 != v7) {
        var v16 uint64 = 1
        if (1.5 == 0.0) {
            var v17 int = 42
            var v18 float32 = (v13 - float32(3.14159))
            for fi19 := 0; fi19 < 2; fi19++ {
                if v15 {
                    if len(v5) > 0 {
                        v5[0] = fi19
                    }
                } else {
                    v5 = append(v5, v17)
                    if v15 {
                        switch "hello" {
                        case "hello":
                            var v20 string = ("" + "foo bar")
                            _ = v20
                        case "foo bar":
                            v15 = (!true)
                        default:
                            fmt.Println(v8, v8, v9)
                        }
                        var v21 uint64 = v16
                        var v22 uint = 4294967295
                        _ = v21
                        _ = v22
                    } else {
                        var v23 int8 = 100
                        if len(v5) > 0 {
                            v5[0] = int(v4)
                        }
                        if v15 {
                            if (!true) {
                                for fi24 := 0; fi24 < 12; fi24++ {
                                    fmt.Println(v6, v16)
                                    _ = fi24
                                }
                            } else {
                                if ((true || false) || v15) {
                                    v4 = (v4 - (v4 << ((2) & 63)))
                                    for fi25 := 0; fi25 < 18; fi25++ {
                                        var v26 uint = 42
                                        _ = fi25
                                        _ = v26
                                    }
                                    v10 = (v7 ^ v7)
                                } else {
                                    fmt.Println(v17, v10, v3)
                                    v4 = (-(v4))
                                    switch 32767 {
                                    case -2147483648:
                                        if len(v5) > 0 {
                                            v5[0] = fi19
                                        }
                                    case 32767:
                                        if false {
                                            v27 := GenS1{F0: (v2 & (v8 + (v2 ^ (v2 * v2)))), F1: v16, F2: true}
                                            _ = v27
                                        }
                                    default:
                                        var v28 string = (("" + ("\n\t" + "a")) + "hello")
                                        _ = v28
                                    }
                                }
                                var v29 float64 = 3.14159
                                _ = v29
                            }
                            var v30 uint8 = 0
                            var v31 float64 = 1e10
                            _ = v30
                            _ = v31
                        }
                        _ = v23
                    }
                }
                var v32 uint32 = 2
                var v33 int = v17
                _ = fi19
                _ = v32
                _ = v33
            }
            _ = v17
            _ = v18
        } else {
            var v34 string = "hello"
            switch "hello" {
            case "hello":
                v4 = -130
            }
            _ = v34
        }
        _ = v16
    } else {
        for fi35 := 0; fi35 < 8; fi35++ {
            if (1.5 < 0.1) {
                if (v15 && v15) {
                    if v15 {
                        v6 += float32(float64(v13))
                    } else {
                        for fi36 := 0; fi36 < 13; fi36++ {
                            var v37 uint64 = 18446744073709551615
                            v38 := GenS2{F0: 32767, F1: (-32768 <= 1000), F2: 604}
                            if v15 {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v3 = float32(1.5)
                                v7 += v9
                            }
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        v39 := GenS3{F0: "foo bar", F1: (v2 == -507), F2: fi35}
                        _ = v39
                    }
                    v6 -= 1e10
                    v7 += (v7 & v9)
                }
            } else {
                v5 = append(v5, fi35)
                var v40 uint64 = 18446744073709551615
                switch 2147483647 {
                case 4294967295:
                    for fi41 := 0; fi41 < 16; fi41++ {
                        if len(v5) > 0 {
                            v5[0] = -2147483648
                        }
                        switch "foo bar" {
                        case "hello":
                            if false {
                                fmt.Println(fi41)
                                var v42 int8 = 127
                                v43 := GenS4{F0: (v2 << ((966) & 63)), F1: 0}
                                if len(v5) > 0 {
                                    v5[0] = int(v40)
                                }
                                if false {
                                    fmt.Println(v8, v13, v6)
                                    var v44 int = (fi35 * (fi35 % ((fi35) | 1)))
                                    _ = v44
                                }
                                _ = v42
                                _ = v43
                            } else {
                                if v15 {
                                    if v15 {
                                        if v15 {
                                            v45 := GenS1{F0: v8, F1: 0, F2: false}
                                            v9 = uint16(v40)
                                            v46 := GenS2{F0: 304, F1: (v15 || v15), F2: v40}
                                            var v47 string = "a"
                                            _ = len(v5)
                                            var v48 uint16 = v7
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        } else {
                                            v8 += (v2 / ((42) | 1))
                                            var v49 string = "\n\t"
                                            _ = v49
                                        }
                                        v6 += float32(1e10)
                                        var v50 uint32 = 4294967295
                                        var v51 bool = v15
                                        _ = v50
                                        _ = v51
                                    } else {
                                        switch (((("hello" + "\n\t") + "a") + "") + ("foo bar" + "hello")) {
                                        case "\n\t":
                                            v6 = v6
                                        case "foo bar":
                                            v52 := GenS5{F0: v4, F1: v7}
                                            var v53 uint16 = v1
                                            _ = v52
                                            _ = v53
                                        case "hello":
                                            fmt.Println(fi35, v13)
                                        }
                                        if (v8 < (v2 >> ((255) & 63))) {
                                            if v15 {
                                                var v54 uint32 = 2
                                                if (1 >= 1) {
                                                    var v55 uint = 551
                                                    _ = v55
                                                }
                                                _ = v54
                                            }
                                        } else {
                                            if v15 {
                                                var v56 uint = 2
                                                if len(v5) > 0 {
                                                    v5[0] = fi35
                                                }
                                                _ = v56
                                            }
                                            if v15 {
                                                if v15 {
                                                    var v57 int32 = v2
                                                    fmt.Println(v6, v10, v7)
                                                    if (!(fi35 > -42)) {
                                                        v58 := GenS4{F0: (v8 | v8), F1: (v10 >> ((18446744073709551615) & 63))}
                                                        var v59 uint64 = uint64(v10)
                                                        if (2 >= 255) {
                                                            v60 := GenS1{F0: v8, F1: 0, F2: (v13 != 1.5)}
                                                            _ = v60
                                                        } else {
                                                            if false {
                                                                v7 = v10
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                if (18446744073709551615 >= 4294967295) {
                                                                    var v61 uint8 = 2
                                                                    _ = v61
                                                                } else {
                                                                    var v62 int16 = 0
                                                                    _ = v62
                                                                }
                                                            }
                                                            if (true && (v6 != 0.0)) {
                                                                v5 = append(v5, fi41)
                                                                v63 := GenS3{F0: ((("hello" + "foo bar") + ("hello" + (("" + "hello") + ("hello" + (("\n\t" + "a") + "\n\t"))))) + ("\n\t" + "a")), F1: ((^(v59)) > (^(v40))), F2: (fi35 << ((65535) & 63))}
                                                                if v15 {
                                                                    var v64 int32 = 2147483647
                                                                    _ = v64
                                                                }
                                                                var v65 float32 = v3
                                                                fmt.Println(v4, v3, v2)
                                                                _ = v63
                                                                _ = v65
                                                            } else {
                                                                fmt.Println(v9, v7)
                                                                if false {
                                                                    var v66 int64 = (-(v4))
                                                                    switch (fi35 + int(v10)) {
                                                                    case 963:
                                                                        v67 := GenS1{F0: 2147483647, F1: (^(v40)), F2: ((((0.1 - 3.14159) * 1e10) + float64(v3)) != 1e10)}
                                                                        _ = v67
                                                                    case 127:
                                                                        _ = len(v5)
                                                                    default:
                                                                        switch "hello" {
                                                                        case "":
                                                                            var v68 uint16 = (^(v10))
                                                                            _ = v68
                                                                        case "hello":
                                                                            v2 -= -485
                                                                        }
                                                                    }
                                                                    _ = v66
                                                                } else {
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    switch 42 {
                                                                    case 42:
                                                                        _ = len(v5)
                                                                    case 4294967295:
                                                                        v69 := GenS4{F0: v57, F1: v7}
                                                                        _ = v69
                                                                    }
                                                                    var v70 uint16 = (v1 | v9)
                                                                    var v71 uint16 = 2
                                                                    _ = v70
                                                                    _ = v71
                                                                }
                                                                var v72 uint32 = 2147483648
                                                                if v15 {
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    var v73 uint16 = v7
                                                                    _ = v73
                                                                } else {
                                                                    if v15 {
                                                                        v6 -= (((((0.0 - v13) * 0.0) * v13) + ((0.1 - (float32(0.1) + 3.14159)) + v3)) - (float32((0.1 + float64(v3))) + 1e10))
                                                                    } else {
                                                                        v74 := GenS1{F0: (v2 & (v57 % (((v57 | (v8 * (v8 % (((v8 ^ v8)) | 1))))) | 1))), F1: (v40 << ((2) & 63)), F2: (v15 && v15)}
                                                                        v4 = (v4 / (((v4 & v4)) | 1))
                                                                        v75 := GenS6{F0: (^(v1))}
                                                                        if true {
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                            v8 += (v2 ^ (v2 & (v8 >> ((2) & 63))))
                                                                        } else {
                                                                            switch (^(fi41)) {
                                                                            case -139:
                                                                                v8 = int32(v72)
                                                                            case 9223372036854775807:
                                                                                v15 = ((128 != 116) && v15)
                                                                                fmt.Println(v7, v6, v3)
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = int(v4)
                                                                                }
                                                                                var v76 int = fi35
                                                                                v7 -= 32767
                                                                                if (!true) {
                                                                                    if v15 {
                                                                                        var v77 int32 = (v8 ^ 2147483647)
                                                                                        switch 127 {
                                                                                        case 42:
                                                                                            v78 := GenS7{F0: true, F1: 0, F2: v76}
                                                                                            _ = v78
                                                                                        case 32767:
                                                                                            var v79 uint16 = 292
                                                                                            _ = v79
                                                                                        }
                                                                                        var v80 int32 = (v57 + v8)
                                                                                        var v81 int = v76
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = (v76 | (v76 / (((^(v76))) | 1)))
                                                                                        }
                                                                                        v82 := GenS7{F0: v15, F1: 42, F2: (fi41 | (v76 >> ((2) & 63)))}
                                                                                        var v83 uint8 = 2
                                                                                        _ = v77
                                                                                        _ = v80
                                                                                        _ = v81
                                                                                        _ = v82
                                                                                        _ = v83
                                                                                    } else {
                                                                                        if true {
                                                                                            if (false && (fi41 < fi41)) {
                                                                                                if ("" >= ("hello" + (("foo bar" + "foo bar") + "hello"))) {
                                                                                                    v84 := GenS8{F0: 0, F1: 255}
                                                                                                    fmt.Println(fi41)
                                                                                                    v5 = append(v5, (fi35 % (((fi41 * fi41)) | 1)))
                                                                                                    v5 = append(v5, 32767)
                                                                                                    v13 -= v3
                                                                                                    _ = v84
                                                                                                }
                                                                                                var v85 uint32 = v72
                                                                                                _ = v85
                                                                                            } else {
                                                                                                v86 := GenS9{F0: v2}
                                                                                                _ = v86
                                                                                            }
                                                                                            var v87 int64 = v4
                                                                                            _ = v87
                                                                                        } else {
                                                                                            v88 := GenS7{F0: (v15 && v15), F1: 29, F2: -9223372036854775808}
                                                                                            v9 = v7
                                                                                            _ = v88
                                                                                        }
                                                                                        v76 -= fi35
                                                                                        v89 := GenS3{F0: "", F1: v15, F2: fi41}
                                                                                        _ = v89
                                                                                    }
                                                                                }
                                                                                _ = v76
                                                                            default:
                                                                                v9 -= (v10 * uint16(v59))
                                                                            }
                                                                            var v90 string = "hello"
                                                                            _ = v90
                                                                        }
                                                                        _ = v74
                                                                        _ = v75
                                                                    }
                                                                    switch 4294967295 {
                                                                    case 127:
                                                                        var v91 float32 = (v3 * v3)
                                                                        _ = v91
                                                                    case 1:
                                                                        var v92 float32 = v3
                                                                        _ = v92
                                                                    case -224:
                                                                        fmt.Println(v7, v59, v72)
                                                                    }
                                                                }
                                                                _ = v72
                                                            }
                                                        }
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                    v13 += v6
                                                    v40 += 750
                                                    if true {
                                                        var v93 float32 = ((float32(1e10) - 0.0) - float32(3.14159))
                                                        if v15 {
                                                            if v15 {
                                                                var v94 uint = 966
                                                                if false {
                                                                    v95 := GenS9{F0: (v8 & (^(v57)))}
                                                                    var v96 uint = v94
                                                                    var v97 float32 = float32(((((0.0 - -2.25) + 1.5) - 0.1) * 1e10))
                                                                    _ = v95
                                                                    _ = v96
                                                                    _ = v97
                                                                } else {
                                                                    var v98 int8 = -128
                                                                    switch fi41 {
                                                                    case 32767:
                                                                        if (-2147483648 == v2) {
                                                                            _ = len(v5)
                                                                            v5 = append(v5, fi35)
                                                                        }
                                                                    default:
                                                                        var v99 string = "hello"
                                                                        _ = v99
                                                                    }
                                                                    if v15 {
                                                                        v93 = (v3 * v93)
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                    }
                                                                    _ = v98
                                                                }
                                                                _ = v94
                                                            } else {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            }
                                                            v100 := GenS10{F0: "foo bar"}
                                                            _ = v100
                                                        } else {
                                                            if (v6 > v3) {
                                                                var v101 int64 = (v4 ^ int64(v2))
                                                                if ((v8 | 0) <= (v57 % ((0) | 1))) {
                                                                    if len(v5) > 0 {
                                                                        v5[0] = 2147483647
                                                                    }
                                                                } else {
                                                                    v15 = true
                                                                }
                                                                _ = v101
                                                            } else {
                                                                switch "hello" {
                                                                case "":
                                                                    if (((0.1 - ((0.1 * 1.5) * ((-2.25 / (float64((v3 + v3)) * -2.25)) * -2.25))) + 0.1) != (0.0 / 1.5)) {
                                                                        fmt.Println(v93, v1, v10)
                                                                        v8 -= (^(v2))
                                                                        if (0 == -1) {
                                                                            switch "hello" {
                                                                            case "a":
                                                                                fmt.Println(v8, v6)
                                                                            case "hello":
                                                                                var v102 bool = (709 <= 359)
                                                                                _ = v102
                                                                            case "foo bar":
                                                                                _ = len(v5)
                                                                            default:
                                                                                _ = len(v5)
                                                                            }
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                        } else {
                                                                            v103 := GenS11{F0: (fi41 << ((694) & 63)), F1: v8, F2: v4}
                                                                            _ = v103
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    var v104 float32 = (v3 / v6)
                                                                    _ = v104
                                                                case "foo bar":
                                                                    v13 = (((0.0 + (float32(1e10) - v6)) - (((3.14159 * v3) / 1.5) * (v13 * float32((0.0 * (float64(v6) + -2.25)))))) * 1e10)
                                                                default:
                                                                    v15 = (false && (v15 || v15))
                                                                }
                                                                switch "" {
                                                                case "\n\t":
                                                                    if len(v5) > 0 {
                                                                        v5[0] = fi35
                                                                    }
                                                                }
                                                                var v105 uint = 309
                                                                _ = v105
                                                            }
                                                            var v106 float32 = v13
                                                            var v107 float64 = (1e10 / 1.0)
                                                            v9 = (^(v1))
                                                            _ = v106
                                                            _ = v107
                                                        }
                                                        if ((uint32(v40) > 2147483648) || false) {
                                                            v8 = 42
                                                            if (!v15) {
                                                                fmt.Println(fi35)
                                                                v108 := GenS4{F0: v2, F1: 2}
                                                                if false {
                                                                    if true {
                                                                        switch (("" + ("hello" + ("foo bar" + "hello"))) + ("hello" + "\n\t")) {
                                                                        case "\n\t":
                                                                            v109 := GenS10{F0: "foo bar"}
                                                                            _ = v109
                                                                        case "":
                                                                            v110 := GenS5{F0: v4, F1: (v10 - v9)}
                                                                            _ = v110
                                                                        case "hello":
                                                                            v111 := GenS12{F0: -1, F1: 831}
                                                                            _ = v111
                                                                        }
                                                                        _ = len(v5)
                                                                        v112 := GenS8{F0: 100, F1: 0}
                                                                        _ = v112
                                                                    }
                                                                    v93 = (v6 - 1e10)
                                                                    if len(v5) > 0 {
                                                                        v5[0] = (fi35 / ((int(v57)) | 1))
                                                                    }
                                                                } else {
                                                                    v40 += v40
                                                                    _ = len(v5)
                                                                    v9 = (v1 % (((v10 & uint16(v40))) | 1))
                                                                }
                                                                _ = v108
                                                            }
                                                        } else {
                                                            fmt.Println(v1)
                                                            if true {
                                                                var v113 int16 = 1000
                                                                v6 -= v6
                                                                _ = v113
                                                            }
                                                            v5 = append(v5, (fi35 | (fi41 - (fi41 + (fi41 << ((18446744073709551615) & 63))))))
                                                        }
                                                        _ = v93
                                                    }
                                                    _ = v57
                                                }
                                                var v114 float32 = (v13 / (v13 + ((v3 + (v13 - ((v3 - (v3 - -2.25)) / (v13 - float32(float64(v13)))))) - (v3 * 3.14159))))
                                                var v115 uint8 = 1
                                                _ = v114
                                                _ = v115
                                            } else {
                                                if v15 {
                                                    v5 = append(v5, int(v1))
                                                    var v116 uint = 695
                                                    _ = v116
                                                } else {
                                                    var v117 uint8 = 0
                                                    v118 := GenS13{F0: 236, F1: 32767}
                                                    _ = v117
                                                    _ = v118
                                                }
                                                v119 := GenS14{F0: 42, F1: "", F2: 95}
                                                v3 = 1e10
                                                _ = v119
                                            }
                                            var v120 uint = 2
                                            _ = v120
                                        }
                                        v121 := GenS15{F0: v40, F1: 220, F2: 0}
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v121
                                    }
                                    fmt.Println(v7)
                                }
                                var v122 uint16 = (v1 << ((4294967295) & 63))
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = v122
                            }
                        case "":
                            v5 = append(v5, 274)
                        case "a":
                            if (4294967295 >= v40) {
                                if true {
                                    if (304 <= 1) {
                                        if len(v5) > 0 {
                                            v5[0] = 32768
                                        }
                                        if (v15 && (v15 && true)) {
                                            var v123 int8 = 0
                                            v124 := GenS8{F0: v123, F1: 210}
                                            _ = v123
                                            _ = v124
                                        }
                                        v6 -= 0.1
                                        v125 := GenS10{F0: "a"}
                                        _ = v125
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                v5 = append(v5, (fi41 & fi35))
                                v1 = v9
                            }
                            v10 = v9
                        default:
                            _ = len(v5)
                        }
                        v9 = (v1 ^ (v10 / (((v9 - (v1 / (((v9 * (v1 >> ((18446744073709551615) & 63)))) | 1)))) | 1)))
                        var v126 float32 = v3
                        _ = fi41
                        _ = v126
                    }
                default:
                    if (!true) {
                        var v127 int = 626
                        _ = v127
                    } else {
                        switch ("foo bar" + "a") {
                        case "":
                            v13 = 0.1
                        case "foo bar":
                            for fi128 := 0; fi128 < 6; fi128++ {
                                switch fi128 {
                                case 2147483647:
                                    v10 = 1
                                default:
                                    if (v13 == float32((1e10 - float64(v13)))) {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v2 -= -228
                                    }
                                }
                                var v129 float64 = 0.1
                                _ = fi128
                                _ = v129
                            }
                        case "hello":
                            if (fi35 <= fi35) {
                                switch ("foo bar" + ("foo bar" + "foo bar")) {
                                case "":
                                    v130 := GenS7{F0: v15, F1: 255, F2: (^(fi35))}
                                    _ = v130
                                case "foo bar":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                switch "a" {
                                case "":
                                    for fi131 := 0; fi131 < 7; fi131++ {
                                        var v132 bool = v15
                                        _ = fi131
                                        _ = v132
                                    }
                                case "a":
                                    var v133 int64 = (v4 / ((-804) | 1))
                                    _ = v133
                                case "hello":
                                    if false {
                                        if v15 {
                                            var v134 uint8 = 255
                                            _ = v134
                                        }
                                    }
                                }
                                v2 = v2
                            } else {
                                if ((2 > 255) && (!((float32(0.0) * v13) != (v13 / v13)))) {
                                    fmt.Println(v7)
                                    for fi135 := 0; fi135 < 14; fi135++ {
                                        v136 := GenS2{F0: 1000, F1: v15, F2: v40}
                                        _ = fi135
                                        _ = v136
                                    }
                                } else {
                                    v137 := GenS9{F0: int32(fi35)}
                                    v138 := GenS12{F0: 32767, F1: 65535}
                                    _ = v137
                                    _ = v138
                                }
                            }
                        default:
                            var v139 uint16 = (v9 << ((2) & 63))
                            _ = v139
                        }
                        var v140 float64 = -2.25
                        if (127 == -125) {
                            v141 := GenS16{F0: 613, F1: (v2 <= v2)}
                            _ = v141
                        } else {
                            if (true || v15) {
                                v3 -= 3.14159
                                v140 -= 0.0
                            } else {
                                v10 = v1
                            }
                        }
                        _ = v140
                    }
                }
                _ = v40
            }
            _ = fi35
        }
        var v142 uint32 = 800
        _ = v142
    }
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v9, v13)
    case "\n\t":
        v13 = v13
    }
    switch "foo bar" {
    case "":
        switch 538 {
        case -2147483648:
            switch 101 {
            case 32767:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case 619:
                var v143 int64 = (v4 | (-(v4)))
                _ = v143
            case -385:
                v4 += (v4 * 9223372036854775807)
            }
        default:
            if (false && (!v15)) {
                v144 := GenS4{F0: v2, F1: v1}
                _ = v144
            }
        }
    case "foo bar":
        v10 += (v9 + (v9 + (^(v9))))
    default:
        fmt.Println(v10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
