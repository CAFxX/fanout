package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 string; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 int16; F1 uint8; F2 int64 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 bool; F1 int16; F2 int16 }
type GenS10 struct { F0 int16; F1 int16; F2 uint }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 int32 = 0
    var v2 bool = (true || false)
    v3 := make(map[int]uint32)
    var v4 uint64 = 264
    v5 := []string{"\n\t", ("foo bar" + ("" + ("a" + "a"))), ((("" + (("foo bar" + "a") + "hello")) + "foo bar") + "")}
    var v6 int32 = (v1 / (((v1 / (((v1 & v1)) | 1))) | 1))
    v1 = v6
    if v2 {
        switch -933 {
        case 42:
            v1 -= v6
        case 2147483647:
            fmt.Println(v4)
        }
        v4 += (v4 + v4)
        v7 := GenS1{F0: 217}
        _ = v7
    } else {
        switch 319 {
        case 42:
            v6 -= -2147483648
        case 32768:
            v2 = (!v2)
        case 9223372036854775807:
            v4 -= 1
        default:
            if true {
                v4 -= 18446744073709551615
                v4 = v4
            } else {
                v5 = append(v5, ("foo bar" + ""))
                switch "" {
                case "a":
                    var v8 uint16 = 2
                    _ = v8
                case "hello":
                    var v9 float32 = -2.25
                    _ = v9
                default:
                    var v10 int8 = 127
                    _ = v10
                }
                var v11 bool = v2
                _ = v11
            }
        }
    }
    v6 = (v1 | v1)
    v2 = v2
    v12 := v3[-335]
    v13 := v3[-953]
    for fi14 := 0; fi14 < 1; fi14++ {
        fmt.Println(v6)
        delete(v3, 9223372036854775807)
        _ = fi14
    }
    var v15 float64 = 0.0
    _ = len(v3)
    for fi16 := 0; fi16 < 13; fi16++ {
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int32 = 2147483647
            if true {
                if false {
                    v12 += (v12 & v13)
                } else {
                    if (v2 || true) {
                        v18 += (v18 + v18)
                        switch "" {
                        case "":
                            if (!(((v6 + v6) != (v1 % (((v18 * v1)) | 1))) && v2)) {
                                v19 := GenS1{F0: 65535}
                                var v20 int64 = 42
                                _ = v19
                                _ = v20
                            } else {
                                if (2147483647 > (v13 | 908)) {
                                    v21 := GenS2{F0: 78}
                                    var v22 uint8 = 204
                                    if v2 {
                                        v23 := GenS3{F0: 9223372036854775807, F1: ((("hello" + "hello") + "") + "\n\t"), F2: fi16}
                                        _ = v23
                                    }
                                    _ = v21
                                    _ = v22
                                } else {
                                    var v24 int = fi17
                                    _ = v24
                                }
                                v12 = (v13 + v13)
                                if (65535 >= 65535) {
                                    var v25 uint32 = (^(v13))
                                    _ = v25
                                }
                                if (v2 || v2) {
                                    var v26 int16 = 1
                                    _ = len(v3)
                                    v6 -= -710
                                    if true {
                                        v6 += (^(v6))
                                    }
                                    fmt.Println(v13)
                                    _ = v26
                                } else {
                                    v5 = append(v5, "foo bar")
                                    v27 := v3[(fi16 % ((-175) | 1))]
                                    _ = v27
                                }
                                switch fi17 {
                                case -182:
                                    v18 = v1
                                }
                            }
                            var v28 int64 = 42
                            _ = v28
                        case "\n\t":
                            var v29 uint32 = v13
                            _ = v29
                        }
                        switch "\n\t" {
                        case "foo bar":
                            if false {
                                var v30 int32 = v1
                                _ = v30
                            } else {
                                var v31 uint = 255
                                _ = v31
                            }
                        case "a":
                            var v32 int16 = 1
                            v13 = v13
                            var v33 int8 = -81
                            v34 := GenS4{F0: (^(v12))}
                            fmt.Println(v4)
                            var v35 uint32 = uint32(v4)
                            var v36 uint = 255
                            _ = v32
                            _ = v33
                            _ = v34
                            _ = v35
                            _ = v36
                        case "\n\t":
                            v6 -= v6
                            if false {
                                v37 := v3[(fi16 << ((255) & 63))]
                                _ = v37
                            }
                            switch -1 {
                            case 32768:
                                v2 = (v15 >= v15)
                            }
                        default:
                            var v38 int = fi16
                            var v39 string = ""
                            v6 = (v6 ^ (^(v6)))
                            v40 := v3[1]
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                    } else {
                        var v41 int16 = 1
                        var v42 uint = 2
                        if len(v5) > 0 {
                            v5[0] = "foo bar"
                        }
                        switch fi17 {
                        case -9223372036854775808:
                            if false {
                                var v43 uint64 = v4
                                v44 := GenS5{F0: -916}
                                _ = v43
                                _ = v44
                            }
                        case 128:
                            _ = len(v3)
                        default:
                            if true {
                                if (v2 || (42 < 42)) {
                                    fmt.Println(fi17, fi16, fi16)
                                    _ = len(v3)
                                } else {
                                    fmt.Println(v2, v18)
                                    switch (fi16 << ((65535) & 63)) {
                                    case 942:
                                        var v45 int8 = 1
                                        _ = v45
                                    case 32768:
                                        var v46 bool = true
                                        _ = v46
                                    case 32767:
                                        var v47 int32 = 42
                                        _ = v47
                                    }
                                }
                                v48 := GenS4{F0: (v12 ^ 2147483647)}
                                _ = v48
                            }
                        }
                        v13 += v12
                        var v49 int = (fi17 << ((v42) & 63))
                        _ = v41
                        _ = v42
                        _ = v49
                    }
                }
                v50 := GenS6{F0: 128, F1: 1}
                switch 42 {
                case -941:
                    var v51 uint32 = (v13 | (^(v12)))
                    _ = v51
                }
                var v52 int8 = 45
                _ = v50
                _ = v52
            } else {
                v53 := GenS5{F0: 1}
                var v54 string = ((("\n\t" + "foo bar") + "a") + ("a" + "foo bar"))
                _ = v53
                _ = v54
            }
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    fmt.Println(v1, v12, v2)
    if (v2 || false) {
        _ = len(v5)
    } else {
        if (2147483647 != (v1 - (v1 | 1))) {
            for fi55 := 0; fi55 < 8; fi55++ {
                fmt.Println(v12)
                var v56 string = "hello"
                if v2 {
                    v15 = 0.1
                } else {
                    v57 := GenS5{F0: 1}
                    _ = v57
                }
                _ = fi55
                _ = v56
            }
            v4 = v4
        } else {
            if (2 != 1) {
                fmt.Println(v6, v15)
                v6 = 42
                switch 0 {
                case 2147483647:
                    var v58 uint = 2
                    _ = v58
                default:
                    for fi59 := 0; fi59 < 4; fi59++ {
                        var v60 string = ""
                        v15 = ((v15 * v15) * 1.5)
                        for fi61 := 0; fi61 < 10; fi61++ {
                            var v62 bool = v2
                            fmt.Println(v62)
                            _ = fi61
                            _ = v62
                        }
                        _ = fi59
                        _ = v60
                    }
                }
            } else {
                var v63 uint32 = 158
                var v64 bool = ((0 != 42) || (!(float64((-2.25 * 3.14159)) != (v15 - ((1e10 + v15) + v15)))))
                v65 := GenS7{F0: 589, F1: 127, F2: 42}
                _ = v63
                _ = v64
                _ = v65
            }
            if true {
                if (745 <= -2147483648) {
                    v1 += (v1 / ((-2147483648) | 1))
                    for fi66 := 0; fi66 < 1; fi66++ {
                        for fi67 := 0; fi67 < 4; fi67++ {
                            switch "a" {
                            case "":
                                if v2 {
                                    if v2 {
                                        v5 = append(v5, "a")
                                    } else {
                                        v68 := GenS6{F0: 0, F1: 1}
                                        _ = v68
                                    }
                                    var v69 int8 = 31
                                    _ = v69
                                } else {
                                    _ = len(v5)
                                    v15 = (1.5 + v15)
                                    if v2 {
                                        v70 := v3[fi67]
                                        v3[(-(fi66))] = (v13 >> ((18446744073709551615) & 63))
                                        var v71 uint8 = 0
                                        v3[9223372036854775807] = v70
                                        var v72 int = (-(fi67))
                                        switch "foo bar" {
                                        case "\n\t":
                                            switch "foo bar" {
                                            case "foo bar":
                                                fmt.Println(v4, fi67, v71)
                                                v73 := v3[(fi67 ^ 9223372036854775807)]
                                                _ = v73
                                            case "":
                                                _ = len(v5)
                                            case "a":
                                                var v74 int = fi66
                                                _ = v74
                                            }
                                        case "hello":
                                            switch (("hello" + "a") + "foo bar") {
                                            case "":
                                                var v75 int = 42
                                                _ = v75
                                            case "hello":
                                                v76 := GenS8{F0: uint32(v71)}
                                                _ = v76
                                            case "a":
                                                if (1e10 >= ((((((float32((v15 + -2.25)) - 0.1) + -2.25) * 1e10) / float32((0.0 + v15))) / -2.25) + -2.25)) {
                                                    v77 := GenS6{F0: v71, F1: 0}
                                                    _ = v77
                                                }
                                                if (65535 == 7) {
                                                    v78 := GenS9{F0: false, F1: -1, F2: 1}
                                                    switch (fi67 + (^(fi66))) {
                                                    case -9223372036854775808:
                                                        var v79 float32 = (1e10 + (1e10 / 0.1))
                                                        _ = v79
                                                    case 32768:
                                                        var v80 int16 = 1
                                                        _ = v80
                                                    case 32767:
                                                        switch fi66 {
                                                        case 32768:
                                                            v81 := GenS9{F0: ((!v2) && false), F1: 367, F2: -872}
                                                            _ = v81
                                                        case -790:
                                                            v2 = (v1 <= (v1 >> ((678) & 63)))
                                                        default:
                                                            var v82 float64 = 3.14159
                                                            _ = v82
                                                        }
                                                        var v83 int64 = 42
                                                        _ = v83
                                                    default:
                                                        v84 := GenS5{F0: 1}
                                                        _ = v84
                                                    }
                                                    _ = v78
                                                } else {
                                                    var v85 uint16 = 1
                                                    _ = v85
                                                }
                                            default:
                                                fmt.Println(v72)
                                            }
                                        case "a":
                                            switch (v72 >> ((0) & 63)) {
                                            case 1:
                                                v12 -= 50
                                            case 0:
                                                switch fi67 {
                                                case 966:
                                                    if false {
                                                        var v86 uint16 = 65535
                                                        var v87 float32 = 3.14159
                                                        _ = v86
                                                        _ = v87
                                                    }
                                                case 32767:
                                                    switch (((("foo bar" + "hello") + "foo bar") + (("a" + "hello") + ("hello" + ""))) + "") {
                                                    case "":
                                                        var v88 float32 = ((1.5 - 0.0) / (-2.25 / (float32((v15 / (float64((float32((v15 / (3.14159 / v15))) * 1.5)) - float64(float32(v15))))) * (1e10 * ((float32(v15) / 1.0) / 1e10)))))
                                                        _ = v88
                                                    default:
                                                        fmt.Println(v71, v1)
                                                    }
                                                case 293:
                                                    fmt.Println(v13, v6, fi67)
                                                default:
                                                    var v89 uint8 = v71
                                                    _ = v89
                                                }
                                            case 127:
                                                v90 := GenS9{F0: false, F1: 0, F2: -407}
                                                _ = v90
                                            }
                                            var v91 uint16 = 651
                                            _ = len(v5)
                                            _ = v91
                                        default:
                                            v71 = 67
                                        }
                                        _ = len(v3)
                                        fmt.Println(fi67, v15)
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                    } else {
                                        v1 = (v6 % ((v1) | 1))
                                    }
                                    if ((1.5 > 1.5) && ((v12 << ((18446744073709551615) & 63)) != 1)) {
                                        if ((("" + "foo bar") + (("\n\t" + "\n\t") + "hello")) != "") {
                                            var v92 bool = v2
                                            _ = v92
                                        } else {
                                            var v93 uint = 0
                                            _ = v93
                                        }
                                        v13 += v12
                                        fmt.Println(v13, v1)
                                    }
                                }
                            case "foo bar":
                                if true {
                                    var v94 string = ("foo bar" + "\n\t")
                                    switch fi66 {
                                    case 32768:
                                        switch ("foo bar" + "\n\t") {
                                        case "\n\t":
                                            v94 = (v94 + v94)
                                        }
                                    case 127:
                                        v4 -= v4
                                    case -230:
                                        var v95 uint16 = 88
                                        _ = v95
                                    }
                                    _ = v94
                                } else {
                                    v12 -= (v13 % (((v13 >> ((255) & 63))) | 1))
                                    v96 := GenS3{F0: 42, F1: "a", F2: (fi67 & 85)}
                                    _ = v96
                                }
                            }
                            v5 = append(v5, "hello")
                            if (!v2) {
                                if (("foo bar" + "foo bar") != (("hello" + "\n\t") + ((("a" + "") + "a") + ("" + "\n\t")))) {
                                    v3[fi67] = uint32(v4)
                                    _ = len(v5)
                                    v15 -= (((v15 * (v15 * v15)) / 0.1) - v15)
                                } else {
                                    if (v12 != (v12 / ((v12) | 1))) {
                                        var v97 uint32 = (v13 % (((v13 >> ((0) & 63))) | 1))
                                        v3[fi67] = v97
                                        fmt.Println(v1)
                                        _ = v97
                                    } else {
                                        fmt.Println(v2, fi67, v4)
                                        if v2 {
                                            if (!(4294967295 == (v4 + v4))) {
                                                fmt.Println(v13, v12, v12)
                                            } else {
                                                v98 := GenS10{F0: 309, F1: 232, F2: 4294967295}
                                                var v99 string = ""
                                                v6 = v1
                                                if (false && true) {
                                                    v100 := GenS4{F0: v12}
                                                    _ = v100
                                                } else {
                                                    v6 -= -2147483648
                                                }
                                                _ = v98
                                                _ = v99
                                            }
                                        }
                                        _ = len(v5)
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        fmt.Println(v6, fi67, v2)
                                    }
                                    switch (fi67 / (((fi66 - 42)) | 1)) {
                                    case -452:
                                        if v2 {
                                            _ = len(v3)
                                        } else {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            v1 = v1
                                            _ = len(v5)
                                            if v2 {
                                                fmt.Println(fi66)
                                            }
                                        }
                                    default:
                                        v13 += (v13 & 519)
                                    }
                                }
                                var v101 int8 = 127
                                _ = v101
                            }
                            _ = fi67
                        }
                        var v102 uint8 = 119
                        for fi103 := 0; fi103 < 1; fi103++ {
                            _ = len(v5)
                            _ = len(v3)
                            var v104 int = (-(fi66))
                            _ = fi103
                            _ = v104
                        }
                        _ = fi66
                        _ = v102
                    }
                }
                v105 := GenS11{F0: (v4 & 2)}
                _ = v105
            }
        }
        var v106 float32 = 1e10
        _ = v106
    }
    var v107 int8 = 100
    v12 -= v12
    v108 := GenS2{F0: 2}
    var v109 float32 = 0.0
    var v110 uint16 = 741
    v111 := GenS12{F0: v4, F1: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v107
    _ = v108
    _ = v109
    _ = v110
    _ = v111
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
