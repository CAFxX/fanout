package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 string; F2 int32 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int32; F1 uint64; F2 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint8; F1 string; F2 string }
type GenS8 struct { F0 int16; F1 string; F2 string }
type GenS9 struct { F0 uint16; F1 uint32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int32; F1 int32 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint64; F1 uint32; F2 bool }
type GenS15 struct { F0 int; F1 int32 }
type GenS16 struct { F0 int8 }
func genFn0() string {
    var v6 int8 = -128
    v7 := []int{32768, 32767, -1, 0}
    var v8 int = 42
    if (!(true || false)) {
        v7 = append(v7, v8)
        var v9 int64 = -1
        _ = v9
    } else {
        v8 = (v8 << ((42) & 63))
        switch ((("a" + "hello") + "\n\t") + (("foo bar" + "\n\t") + "foo bar")) {
        case "\n\t":
            if len(v7) > 0 {
                v7[0] = 280
            }
        case "":
            switch int(v6) {
            case -818:
                for fi10 := 0; fi10 < 1; fi10++ {
                    v8 = 75
                    v11 := GenS1{F0: 255}
                    _ = fi10
                    _ = v11
                }
            case -2147483648:
                switch (v8 & 377) {
                case 124:
                    var v12 float32 = 0.1
                    _ = v12
                case 42:
                    var v13 int = 241
                    _ = v13
                case 127:
                    v14 := GenS2{F0: -290, F1: ("\n\t" + ("" + "hello")), F2: -2147483648}
                    _ = v14
                default:
                    v15 := GenS2{F0: -1, F1: "hello", F2: 2147483647}
                    _ = v15
                }
            case 42:
                switch 127 {
                case -1:
                    fmt.Println(v8, v6)
                case 9223372036854775807:
                    v8 = v8
                case 1:
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
            default:
                v16 := GenS1{F0: 2}
                _ = v16
            }
        default:
            for fi17 := 0; fi17 < 6; fi17++ {
                if (true && false) {
                    v6 -= (v6 | (v6 - (v6 & (v6 - v6))))
                    _ = len(v7)
                    var v18 int = (v8 << ((737) & 63))
                    _ = v18
                }
                _ = fi17
            }
        }
        fmt.Println(v6)
    }
    _ = v6
    _ = v7
    _ = v8
    return "hello"
}

func main() {
    var v1 bool = ("hello" < ("\n\t" + "foo bar"))
    var v2 bool = (!(((3.14159 - ((float32(0.0) / -2.25) + -2.25)) != 1e10) && v1))
    var v3 bool = v2
    v4 := make(map[uint32]string)
    var v5 uint64 = 0
    v19 := GenS2{F0: 301, F1: (("a" + "a") + "\n\t"), F2: -1}
    v1 = false
    if v1 {
        if (65535 < 18446744073709551615) {
            for fi20 := 0; fi20 < 14; fi20++ {
                for fi21 := 0; fi21 < 14; fi21++ {
                    v3 = ("hello" <= "")
                    v1 = (fi21 != (fi20 % ((fi21) | 1)))
                    fmt.Println(fi21)
                    v22 := GenS3{F0: ("" + ""), F1: -1}
                    switch fi20 {
                    case 128:
                        if v3 {
                            switch "a" {
                            case "foo bar":
                                var v23 float32 = 0.1
                                _ = v23
                            case "":
                                switch "foo bar" {
                                case "hello":
                                    v24 := GenS2{F0: 138, F1: ("a" + "a"), F2: 42}
                                    _ = v24
                                case "":
                                    var v25 uint32 = uint32(v5)
                                    _ = v25
                                case "\n\t":
                                    var v26 int64 = 1
                                    _ = v26
                                default:
                                    v1 = (false && (!true))
                                }
                            default:
                                if v1 {
                                    v27 := []int{-991}
                                    _ = v27
                                } else {
                                    v3 = (float64(float32(3.14159)) <= 0.0)
                                    switch 42 {
                                    case 9223372036854775807:
                                        fmt.Println(fi21, v5)
                                    case 42:
                                        switch 417 {
                                        case 1:
                                            var v28 bool = v3
                                            _ = v28
                                        case 0:
                                            v29 := []int{-9223372036854775808, (fi21 << ((2) & 63))}
                                            v1 = (!v2)
                                            switch "hello" {
                                            case "\n\t":
                                                var v30 uint32 = 696
                                                v2 = ((float32(-2.25) == (0.1 * ((0.0 + ((1e10 / float32((float64(((-2.25 * -2.25) + (0.0 / 3.14159))) / ((0.1 * (3.14159 / -2.25)) - float64((-2.25 + 0.1)))))) / 1.0)) / 3.14159))) && v3)
                                                v31 := GenS2{F0: 1, F1: "a", F2: 2147483647}
                                                if (255 != 42) {
                                                    var v32 float32 = 1.5
                                                    if v3 {
                                                        switch fi21 {
                                                        case 2147483647:
                                                            var v33 float32 = 1e10
                                                            _ = v33
                                                        case 128:
                                                            if (-551 <= -9223372036854775808) {
                                                                if (v2 || v3) {
                                                                    var v34 uint16 = 2
                                                                    v30 += (v30 | (v30 * (v30 & (v30 & (v30 & v30)))))
                                                                    _ = v34
                                                                } else {
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        var v35 bool = (((float32(0.0) + 1.5) > v32) && (!(751 < -641)))
                                                                        _ = v35
                                                                    }
                                                                    v2 = v1
                                                                }
                                                                var v36 uint16 = 2
                                                                fmt.Println(v36, v30, fi21)
                                                                v37 := GenS3{F0: "a", F1: 127}
                                                                _ = v36
                                                                _ = v37
                                                            }
                                                        case -1:
                                                            var v38 int16 = -785
                                                            fmt.Println(v32, v30, v3)
                                                            switch ("a" + "hello") {
                                                            case "a":
                                                                var v39 uint16 = 42
                                                                _ = v39
                                                            case "\n\t":
                                                                v40 := v4[42]
                                                                _ = v40
                                                            }
                                                            var v41 uint64 = (v5 % ((v5) | 1))
                                                            _ = v38
                                                            _ = v41
                                                        default:
                                                            fmt.Println(v3, v32, fi21)
                                                        }
                                                    } else {
                                                        v1 = false
                                                        v3 = false
                                                        if ((0.1 + v32) <= v32) {
                                                            if v2 {
                                                                fmt.Println(fi20, v3, v32)
                                                                switch fi20 {
                                                                case -742:
                                                                    v1 = (!((fi21 + 127) >= (fi21 << ((289) & 63))))
                                                                case 32768:
                                                                    delete(v4, (v30 + v30))
                                                                }
                                                            } else {
                                                                var v42 int16 = 534
                                                                if v3 {
                                                                    var v43 bool = (fi20 >= 2147483647)
                                                                    v44 := GenS4{F0: -221, F1: v5, F2: -1}
                                                                    _ = len(v4)
                                                                    if ((("foo bar" + ("" + "hello")) == ("hello" + "")) || (-2147483648 >= -2147483648)) {
                                                                        v3 = v1
                                                                        if v2 {
                                                                            v42 -= 0
                                                                            v45 := GenS3{F0: "hello", F1: -128}
                                                                            if (v30 >= v30) {
                                                                                if (v5 <= (v5 >> ((18446744073709551615) & 63))) {
                                                                                    var v46 bool = (v3 && true)
                                                                                    _ = v46
                                                                                } else {
                                                                                    fmt.Println(v5)
                                                                                    _ = len(v4)
                                                                                    v47 := v4[873]
                                                                                    if (!false) {
                                                                                        var v48 int32 = -665
                                                                                        v42 -= v42
                                                                                        var v49 string = v47
                                                                                        _ = v48
                                                                                        _ = v49
                                                                                    } else {
                                                                                        _ = len(v4)
                                                                                        fmt.Println(v3, v43, v3)
                                                                                        var v50 float32 = 1e10
                                                                                        _ = v50
                                                                                    }
                                                                                    _ = v47
                                                                                }
                                                                                v51 := GenS4{F0: 259, F1: 775, F2: (v42 ^ (v42 * (v42 + v42)))}
                                                                                if false {
                                                                                    var v52 float64 = float64(1e10)
                                                                                    _ = v52
                                                                                } else {
                                                                                    if ((0.1 - 1.5) == 3.14159) {
                                                                                        var v53 int8 = 1
                                                                                        _ = v53
                                                                                    }
                                                                                    var v54 uint32 = v30
                                                                                    var v55 int32 = 0
                                                                                    var v56 int8 = -22
                                                                                    var v57 float64 = ((float64(1e10) / float64(-2.25)) / 1.0)
                                                                                    _ = v54
                                                                                    _ = v55
                                                                                    _ = v56
                                                                                    _ = v57
                                                                                }
                                                                                _ = v51
                                                                            } else {
                                                                                var v58 uint64 = v5
                                                                                _ = v58
                                                                            }
                                                                            v29 = append(v29, -414)
                                                                            if v3 {
                                                                                v30 -= v30
                                                                                v59 := GenS4{F0: 1, F1: (v5 & v5), F2: (^(v42))}
                                                                                if v43 {
                                                                                    v60 := GenS5{F0: v42}
                                                                                    switch "" {
                                                                                    case "":
                                                                                        _ = len(v29)
                                                                                    default:
                                                                                        v61 := GenS5{F0: (v42 / ((v42) | 1))}
                                                                                        _ = v61
                                                                                    }
                                                                                    switch fi21 {
                                                                                    case -9223372036854775808:
                                                                                        if (-895 != 0) {
                                                                                            delete(v4, (v30 * 65535))
                                                                                        }
                                                                                    case -2147483648:
                                                                                        switch 1 {
                                                                                        case -62:
                                                                                            v32 -= (v32 * v32)
                                                                                        case 0:
                                                                                            if (255 < 165) {
                                                                                                var v62 int8 = int8(v30)
                                                                                                _ = v62
                                                                                            }
                                                                                            v63 := GenS5{F0: (v42 + (v42 ^ (v42 & v42)))}
                                                                                            _ = v63
                                                                                        default:
                                                                                            var v64 float64 = 0.1
                                                                                            _ = v64
                                                                                        }
                                                                                    case 9223372036854775807:
                                                                                        v5 += (^(v5))
                                                                                    default:
                                                                                        v5 += (v5 | (^(v5)))
                                                                                    }
                                                                                    v32 -= v32
                                                                                    var v65 int32 = -412
                                                                                    _ = v60
                                                                                    _ = v65
                                                                                } else {
                                                                                    if ((fi20 >> ((838) & 63)) <= -1) {
                                                                                        v29 = append(v29, fi20)
                                                                                    }
                                                                                    v66 := GenS6{F0: 343}
                                                                                    _ = v66
                                                                                }
                                                                                _ = v59
                                                                            } else {
                                                                                if v3 {
                                                                                    var v67 int16 = (v42 & -32768)
                                                                                    switch 0 {
                                                                                    case 32768:
                                                                                        var v68 uint16 = 122
                                                                                        _ = v68
                                                                                    case 32767:
                                                                                        if ((v42 >> ((2) & 63)) <= v42) {
                                                                                            var v69 uint16 = 0
                                                                                            _ = len(v4)
                                                                                            _ = len(v4)
                                                                                            _ = v69
                                                                                        } else {
                                                                                            var v70 float32 = (float32(((1.5 * 1.5) - (3.14159 + (1.5 / ((1e10 / 1.0) / float64(1e10)))))) - (float32(float64(v32)) - (1.5 - v32)))
                                                                                            if ((65535 < 252) || ((1 != 921) && (-9223372036854775808 <= (-(fi21))))) {
                                                                                                var v71 uint16 = 897
                                                                                                v71 += 2
                                                                                                v43 = v2
                                                                                                _ = v71
                                                                                            } else {
                                                                                                v43 = ((v42 ^ v67) > -32768)
                                                                                            }
                                                                                            v1 = v43
                                                                                            _ = v70
                                                                                        }
                                                                                    case 706:
                                                                                        v72 := v4[(v30 >> ((18446744073709551615) & 63))]
                                                                                        _ = v72
                                                                                    default:
                                                                                        v3 = (!(1e10 < ((-2.25 / 1.0) + (1e10 + 0.0))))
                                                                                    }
                                                                                    _ = v67
                                                                                } else {
                                                                                    var v73 int8 = 42
                                                                                    switch (("foo bar" + "") + "") {
                                                                                    case "":
                                                                                        v74 := GenS7{F0: 128, F1: ((("" + ("a" + ("\n\t" + "a"))) + "a") + "hello"), F2: "foo bar"}
                                                                                        var v75 uint32 = (^(v30))
                                                                                        var v76 int64 = -599
                                                                                        var v77 uint16 = 64
                                                                                        fmt.Println(fi21, v5)
                                                                                        _ = v74
                                                                                        _ = v75
                                                                                        _ = v76
                                                                                        _ = v77
                                                                                    case "hello":
                                                                                        v30 = v30
                                                                                    case "foo bar":
                                                                                        if (v43 && v43) {
                                                                                            v78 := GenS5{F0: -97}
                                                                                            v79 := GenS8{F0: -751, F1: ("a" + "a"), F2: "foo bar"}
                                                                                            if v1 {
                                                                                                var v80 float64 = 1.5
                                                                                                _ = v80
                                                                                            } else {
                                                                                                if len(v29) >= 2 {
                                                                                                    v29 = v29[1:len(v29)]
                                                                                                }
                                                                                            }
                                                                                            _ = v78
                                                                                            _ = v79
                                                                                        } else {
                                                                                            var v81 uint64 = v5
                                                                                            _ = v81
                                                                                        }
                                                                                    default:
                                                                                        _ = len(v4)
                                                                                    }
                                                                                    _ = v73
                                                                                }
                                                                                var v82 float64 = float64(v32)
                                                                                _ = v82
                                                                            }
                                                                            _ = v45
                                                                        } else {
                                                                            v42 += v42
                                                                        }
                                                                        var v83 uint16 = 255
                                                                        _ = v83
                                                                    } else {
                                                                        switch fi21 {
                                                                        case -2147483648:
                                                                            fmt.Println(v32, fi21)
                                                                            v4[680] = "a"
                                                                            if (9223372036854775807 == 42) {
                                                                                v4[v30] = ("" + "")
                                                                                switch (fi21 >> ((39) & 63)) {
                                                                                case 0:
                                                                                    v42 = -1
                                                                                default:
                                                                                    v1 = (4294967295 == 42)
                                                                                    switch fi21 {
                                                                                    case 754:
                                                                                        v84 := GenS9{F0: 2, F1: 42}
                                                                                        var v85 uint64 = v5
                                                                                        v1 = (v30 <= (v30 * (^(v30))))
                                                                                        v2 = ((v2 || (((((1e10 + (float64(v32) - 0.1)) + 1e10) + 3.14159) + 0.0) <= 3.14159)) && (-1 == 42))
                                                                                        _ = v84
                                                                                        _ = v85
                                                                                    case 278:
                                                                                        v86 := v4[v30]
                                                                                        _ = v86
                                                                                    }
                                                                                    var v87 float32 = v32
                                                                                    var v88 uint64 = v5
                                                                                    _ = v87
                                                                                    _ = v88
                                                                                }
                                                                                v89 := GenS5{F0: (-(v42))}
                                                                                _ = v89
                                                                            } else {
                                                                                v90 := GenS5{F0: (v42 * 1)}
                                                                                _ = v90
                                                                            }
                                                                        case 648:
                                                                            if len(v29) > 0 {
                                                                                v29[0] = fi21
                                                                            }
                                                                        default:
                                                                            if (994 == 917) {
                                                                                var v91 int8 = 94
                                                                                v92 := GenS9{F0: 1, F1: v30}
                                                                                _ = v91
                                                                                _ = v92
                                                                            }
                                                                        }
                                                                        if (int32(v42) >= -867) {
                                                                            var v93 uint64 = v5
                                                                            var v94 uint64 = 260
                                                                            _ = v93
                                                                            _ = v94
                                                                        } else {
                                                                            v42 = int16(v5)
                                                                            var v95 int = fi20
                                                                            var v96 int16 = (v42 ^ v42)
                                                                            fmt.Println(v32, v30, fi20)
                                                                            var v97 uint64 = (v5 >> ((0) & 63))
                                                                            _ = v95
                                                                            _ = v96
                                                                            _ = v97
                                                                        }
                                                                    }
                                                                    fmt.Println(v1, v3, v32)
                                                                    if true {
                                                                        _ = len(v29)
                                                                        v98 := GenS10{F0: v42}
                                                                        v30 += v30
                                                                        _ = v98
                                                                    }
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                                _ = v42
                                                            }
                                                        }
                                                        if ((141 == 255) || v1) {
                                                            var v99 string = "a"
                                                            _ = v99
                                                        }
                                                    }
                                                    v3 = (!true)
                                                    _ = v32
                                                } else {
                                                    _ = len(v4)
                                                    var v100 uint32 = (^(v30))
                                                    _ = len(v4)
                                                    _ = v100
                                                }
                                                _ = len(v4)
                                                _ = v30
                                                _ = v31
                                            case "":
                                                v3 = v1
                                            case "foo bar":
                                                v2 = true
                                                v101 := GenS4{F0: 2147483647, F1: 2, F2: 1000}
                                                var v102 uint16 = 32767
                                                _ = v101
                                                _ = v102
                                            default:
                                                v103 := GenS3{F0: "foo bar", F1: 0}
                                                _ = v103
                                            }
                                            _ = v29
                                        }
                                        var v104 float64 = (1.5 / (0.1 * 3.14159))
                                        v4[716] = ("\n\t" + "")
                                        v5 = v5
                                        _ = v104
                                    case 127:
                                        v3 = v1
                                    }
                                }
                                v2 = (127 < 2)
                            }
                            v5 += v5
                            if true {
                                var v105 uint32 = uint32(v5)
                                _ = v105
                            } else {
                                var v106 int16 = 0
                                v107 := GenS7{F0: uint8(v5), F1: "", F2: ("hello" + ("\n\t" + ""))}
                                _ = v106
                                _ = v107
                            }
                            v108 := GenS1{F0: 833}
                            v3 = v3
                            v3 = false
                            v109 := GenS11{F0: 0}
                            _ = v108
                            _ = v109
                        } else {
                            v1 = false
                            v1 = v3
                            fmt.Println(fi21, v3)
                        }
                    case 0:
                        switch fi20 {
                        case 32767:
                            var v110 uint8 = 1
                            _ = v110
                        case 42:
                            v1 = true
                        case 0:
                            v2 = v1
                        default:
                            v3 = (!true)
                        }
                    }
                    _ = fi21
                    _ = v22
                }
                var v111 int16 = -32768
                if true {
                    var v112 int8 = 127
                    _ = v112
                }
                _ = fi20
                _ = v111
            }
            v5 -= (^(v5))
        } else {
            if v2 {
                var v113 int16 = -985
                _ = v113
            }
        }
    }
    var v114 float64 = (-2.25 - ((0.1 * (0.1 / 3.14159)) - ((1.5 / (1e10 / (-2.25 / (1e10 + 1.5)))) * 0.1)))
    v4[2147483648] = ""
    switch 32768 {
    case 2147483647:
        var v115 float32 = 3.14159
        _ = v115
    }
    if (v3 && (v1 && ("a" >= ("hello" + "a")))) {
        var v116 uint8 = 127
        _ = v116
    } else {
        switch ("a" + "hello") {
        case "foo bar":
            if (!false) {
                v117 := []int{-1, 4294967295}
                _ = v117
            } else {
                var v118 uint8 = 122
                _ = v118
            }
        }
    }
    if v1 {
        if (1 == (v5 << ((447) & 63))) {
            v119 := GenS8{F0: 444, F1: "", F2: ("\n\t" + "foo bar")}
            if v3 {
                fmt.Println(v3, v3, v3)
            } else {
                v120 := GenS7{F0: 29, F1: (("hello" + "hello") + (("foo bar" + (("hello" + "") + "a")) + "")), F2: "hello"}
                _ = v120
            }
            if true {
                v121 := []int{127, 32768, 869, 42}
                if len(v121) >= 2 {
                    v121 = v121[1:len(v121)]
                }
                v5 = v5
                _ = v121
            }
            _ = v119
        } else {
            v4[1] = "foo bar"
            if v3 {
                v122 := GenS12{F0: 42, F1: 42}
                v1 = v3
                _ = v122
            }
            v123 := GenS13{F0: "\n\t"}
            _ = v123
        }
    }
    fmt.Println(v1, v114, v114)
    if v3 {
        v124 := GenS14{F0: 2, F1: 0, F2: true}
        for fi125 := 0; fi125 < 12; fi125++ {
            v126 := GenS14{F0: (^(v5)), F1: 4294967295, F2: v2}
            _ = fi125
            _ = v126
        }
        _ = v124
    } else {
        var v127 float32 = (float32(v114) / 1.5)
        if false {
            v128 := []int{9223372036854775807}
            var v129 uint = 4294967295
            var v130 int16 = -32768
            _ = v128
            _ = v129
            _ = v130
        }
        _ = v127
    }
    v131 := []int{32768}
    v3 = (!((true && true) && ((!v1) || true)))
    var v132 int = 42
    if (!true) {
        v3 = true
    } else {
        var v133 string = "\n\t"
        v134 := GenS13{F0: v133}
        if true {
            switch "a" {
            case "a":
                var v135 uint = 18446744073709551615
                _ = v135
            case "foo bar":
                switch v133 {
                case "":
                    if v3 {
                        var v136 string = v133
                        for fi137 := 0; fi137 < 1; fi137++ {
                            fmt.Println(v2, v3)
                            _ = fi137
                        }
                        v5 += v5
                        _ = v136
                    } else {
                        fmt.Println(v2, v3, v3)
                    }
                case "\n\t":
                    v1 = (!(v133 == (v133 + v133)))
                default:
                    v138 := GenS13{F0: ("" + "foo bar")}
                    _ = v138
                }
            case "hello":
                switch (v133 + v133) {
                case "hello":
                    var v139 uint16 = 0
                    _ = v139
                case "foo bar":
                    fmt.Println(v5, v132, v133)
                case "":
                    if true {
                        var v140 uint16 = 810
                        _ = v140
                    }
                }
            default:
                _ = len(v4)
            }
            if v3 {
                var v141 uint64 = (v5 - v5)
                _ = v141
            }
            for fi142 := 0; fi142 < 19; fi142++ {
                v1 = (v114 <= v114)
                v2 = v1
                _ = fi142
            }
        } else {
            if (32767 == 2) {
                v2 = (!(true || v3))
                var v143 string = "a"
                _ = v143
            }
            v144 := GenS4{F0: 2147483647, F1: 0, F2: 1}
            _ = v144
        }
        _ = v133
        _ = v134
    }
    var v145 float64 = v114
    v114 = -2.25
    v146 := GenS15{F0: 32767, F1: 574}
    if v1 {
        var v147 uint = 786
        switch "hello" {
        case "a":
            v147 += (v147 / ((v147) | 1))
        case "foo bar":
            v148 := GenS16{F0: 22}
            _ = v148
        }
        _ = v147
    }
    fmt.Println("v1:", v1)
    fmt.Println("v114:", v114)
    fmt.Println("slicelen v131:", len(v131), v131)
    fmt.Println("v132:", v132)
    fmt.Println("v145:", v145)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v114
    _ = v131
    _ = v132
    _ = v145
    _ = v146
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
