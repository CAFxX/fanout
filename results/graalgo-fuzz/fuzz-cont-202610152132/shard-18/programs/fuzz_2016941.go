package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 0.1
    v3 := []string{("a" + ("\n\t" + "a"))}
    var v4 uint64 = 579
    v5 := make(map[int]string)
    var v6 string = (((("" + "") + "\n\t") + "a") + "foo bar")
    if ((!(false || true)) && (!true)) {
        v6 = (("hello" + ("a" + v6)) + v6)
        switch "" {
        case "":
            for fi7 := 0; fi7 < 3; fi7++ {
                if (128 < 127) {
                    var v8 int8 = -41
                    _ = v8
                } else {
                    switch fi7 {
                    case 198:
                        var v9 int = (fi7 + -293)
                        _ = v9
                    case 4294967295:
                        var v10 uint = 920
                        _ = v10
                    }
                    var v11 uint16 = 65535
                    _ = v11
                }
                var v12 float32 = 0.1
                _ = fi7
                _ = v12
            }
        }
        switch "\n\t" {
        case "":
            var v13 bool = (37 == 127)
            _ = v13
        case "a":
            fmt.Println(v2, v4)
        case "hello":
            fmt.Println(v2)
        default:
            var v14 string = "foo bar"
            _ = v14
        }
    }
    switch -711 {
    case -9223372036854775808:
        v1["hello"] = v6
    }
    if true {
        var v15 uint64 = v4
        var v16 int8 = 127
        _ = v15
        _ = v16
    }
    if true {
        v2 += v2
    } else {
        switch 4294967295 {
        case 42:
            for fi17 := 0; fi17 < 5; fi17++ {
                for fi18 := 0; fi18 < 3; fi18++ {
                    fmt.Println(v6, v2, fi17)
                    var v19 int8 = -55
                    delete(v1, v6)
                    var v20 uint8 = 46
                    if (true || (0 == 83)) {
                        var v21 uint8 = v20
                        _ = v21
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                v4 += (v4 * (v4 - (v4 % ((v4) | 1))))
                fmt.Println(v4)
                _ = fi17
            }
        case -446:
            v2 -= float64(0.1)
        case 2:
            if (!(v6 != (v6 + (v6 + ("" + ("foo bar" + "a")))))) {
                v2 += v2
                if true {
                    if true {
                        if false {
                            v22 := GenS1{F0: 26, F1: 33, F2: 1}
                            _ = v22
                        } else {
                            fmt.Println(v6, v2)
                            v2 += v2
                        }
                        switch ("hello" + (v6 + v6)) {
                        case "":
                            switch 32768 {
                            case 128:
                                var v23 uint = 1
                                _ = v23
                            case 0:
                                var v24 int64 = -577
                                _ = v24
                            case -627:
                                v2 += 0.1
                            }
                        }
                        for fi25 := 0; fi25 < 11; fi25++ {
                            fmt.Println(v2, v6)
                            _ = fi25
                        }
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            switch v6 {
                            case "foo bar":
                                var v26 int16 = 1
                                _ = v26
                            case "a":
                                v4 = (^(v4))
                            }
                        case "foo bar":
                            v2 -= ((v2 - float64((0.1 * 0.1))) / (((((((3.14159 - v2) * ((v2 + ((v2 - (v2 * v2)) * float64(((-2.25 - ((3.14159 + (-2.25 + (3.14159 * 0.0))) + (1.5 / 3.14159))) * (float32(1e10) / float32(1.5)))))) / float64((-2.25 - 1e10)))) + -2.25) * v2) - (v2 * v2)) + v2) / v2))
                        case "hello":
                            if ((v6 + v6) == v6) {
                                v6 = ("foo bar" + v6)
                            } else {
                                switch ((("a" + v6) + "") + (v6 + v6)) {
                                case "\n\t":
                                    var v27 uint = 1
                                    _ = v27
                                case "a":
                                    var v28 uint32 = 2
                                    _ = v28
                                default:
                                    switch "\n\t" {
                                    case "hello":
                                        fmt.Println(v4, v4, v4)
                                    case "\n\t":
                                        _ = len(v3)
                                    }
                                }
                                switch v6 {
                                case "foo bar":
                                    v6 = v6
                                default:
                                    v4 = (^(v4))
                                }
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                            }
                        default:
                            for fi29 := 0; fi29 < 4; fi29++ {
                                if true {
                                    if false {
                                        if true {
                                            v30 := GenS1{F0: 65535, F1: 42, F2: 65535}
                                            v2 = 3.14159
                                            _ = v30
                                        } else {
                                            for fi31 := 0; fi31 < 6; fi31++ {
                                                var v32 uint8 = 113
                                                switch fi31 {
                                                case 4294967295:
                                                    switch fi31 {
                                                    case -9223372036854775808:
                                                        var v33 uint64 = v4
                                                        _ = v33
                                                    case 4294967295:
                                                        var v34 float64 = 3.14159
                                                        _ = v34
                                                    case 0:
                                                        if true {
                                                            var v35 int8 = 127
                                                            var v36 uint16 = 444
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v35
                                                            _ = v36
                                                        } else {
                                                            var v37 int = fi29
                                                            if true {
                                                                v38 := GenS1{F0: 1, F1: -2147483648, F2: 2}
                                                                delete(v1, (v6 + "\n\t"))
                                                                if false {
                                                                    v32 = 255
                                                                    if (!true) {
                                                                        if false {
                                                                            var v39 bool = false
                                                                            _ = v39
                                                                        } else {
                                                                            var v40 uint64 = 1
                                                                            v41 := GenS1{F0: 68, F1: -385, F2: 2147483648}
                                                                            fmt.Println(v4, fi31)
                                                                            _ = v40
                                                                            _ = v41
                                                                        }
                                                                        v4 = uint64(v32)
                                                                        v42 := GenS1{F0: 375, F1: 42, F2: 652}
                                                                        _ = v42
                                                                    }
                                                                    if (652 >= 4294967295) {
                                                                        var v43 float64 = v2
                                                                        if false {
                                                                            var v44 uint = 0
                                                                            v45 := GenS2{F0: (!true)}
                                                                            fmt.Println(v4)
                                                                            v4 = 1
                                                                            var v46 uint64 = 933
                                                                            _ = v44
                                                                            _ = v45
                                                                            _ = v46
                                                                        } else {
                                                                            if true {
                                                                                v2 = (((v2 + (-2.25 - v43)) - v43) * 3.14159)
                                                                                v4 += (v4 ^ v4)
                                                                                fmt.Println(v43, v4)
                                                                            }
                                                                            v4 = (v4 | v4)
                                                                        }
                                                                        _ = len(v5)
                                                                        _ = v43
                                                                    }
                                                                }
                                                                _ = v38
                                                            } else {
                                                                var v47 float32 = 0.1
                                                                _ = v47
                                                            }
                                                            var v48 string = v6
                                                            _ = v37
                                                            _ = v48
                                                        }
                                                    }
                                                case 292:
                                                    v2 = float64(0.0)
                                                    v32 = 0
                                                case -9223372036854775808:
                                                    v4 -= 0
                                                }
                                                _ = fi31
                                                _ = v32
                                            }
                                            if true {
                                                switch (v6 + ("hello" + v6)) {
                                                case "\n\t":
                                                    v2 = float64(-2.25)
                                                default:
                                                    for fi49 := 0; fi49 < 7; fi49++ {
                                                        v4 -= (v4 - v4)
                                                        v6 = (v6 + (v6 + "a"))
                                                        var v50 float32 = (3.14159 + 1.5)
                                                        switch (^(fi49)) {
                                                        case 32768:
                                                            if false {
                                                                var v51 uint16 = 65535
                                                                if false {
                                                                    fmt.Println(v4, v51, fi29)
                                                                    var v52 uint16 = (v51 >> ((18446744073709551615) & 63))
                                                                    _ = v52
                                                                }
                                                                _ = len(v3)
                                                                v50 = 0.0
                                                                _ = v51
                                                            }
                                                        case 838:
                                                            var v53 float64 = (3.14159 + v2)
                                                            _ = v53
                                                        case 4294967295:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v54 := GenS1{F0: 384, F1: 1, F2: 2}
                                                            if false {
                                                                if false {
                                                                    v6 = v6
                                                                    var v55 uint32 = 65535
                                                                    var v56 uint = 65535
                                                                    _ = v55
                                                                    _ = v56
                                                                } else {
                                                                    var v57 int16 = 1000
                                                                    fmt.Println(v4, fi49)
                                                                    if (((-66 == -1) || false) && true) {
                                                                        switch v6 {
                                                                        case "foo bar":
                                                                            if true {
                                                                                var v58 bool = (!false)
                                                                                fmt.Println(v58, v57, v6)
                                                                                v2 = 1.5
                                                                                _ = v58
                                                                            }
                                                                        case "\n\t":
                                                                            v57 = v57
                                                                        default:
                                                                            if false {
                                                                                v59 := GenS1{F0: 2147483647, F1: -658, F2: 530}
                                                                                var v60 uint32 = 915
                                                                                if (532 >= 4294967295) {
                                                                                    var v61 int16 = (v57 | (-(v57)))
                                                                                    _ = v61
                                                                                }
                                                                                _ = v59
                                                                                _ = v60
                                                                            } else {
                                                                                if false {
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = v6
                                                                                    }
                                                                                    v2 -= v2
                                                                                    var v62 uint64 = 0
                                                                                    _ = v62
                                                                                }
                                                                                var v63 string = v6
                                                                                _ = v63
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            v50 -= v50
                                                                            if true {
                                                                                var v64 int16 = 1
                                                                                if false {
                                                                                    if false {
                                                                                        var v65 uint32 = 2147483647
                                                                                        _ = v65
                                                                                    } else {
                                                                                        var v66 int32 = 344
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = "hello"
                                                                                        }
                                                                                        var v67 int32 = v66
                                                                                        _ = v66
                                                                                        _ = v67
                                                                                    }
                                                                                }
                                                                                _ = v64
                                                                            } else {
                                                                                v4 -= 0
                                                                                switch (v6 + "a") {
                                                                                case "":
                                                                                    switch (("hello" + v6) + "a") {
                                                                                    case "\n\t":
                                                                                        v6 = v6
                                                                                    case "":
                                                                                        if false {
                                                                                            v2 = float64(v50)
                                                                                            var v68 int64 = 0
                                                                                            if (!(1 > 1)) {
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = v6
                                                                                                }
                                                                                                fmt.Println(fi29, fi29, v4)
                                                                                                _ = len(v5)
                                                                                            } else {
                                                                                                v57 += v57
                                                                                                v4 -= v4
                                                                                            }
                                                                                            _ = v68
                                                                                        }
                                                                                    default:
                                                                                        v50 -= v50
                                                                                    }
                                                                                case "foo bar":
                                                                                    v2 = -2.25
                                                                                }
                                                                                var v69 float64 = 3.14159
                                                                                _ = v69
                                                                            }
                                                                        }
                                                                        if true {
                                                                            var v70 int8 = -128
                                                                            v6 = v6
                                                                            v71 := GenS2{F0: (2 > 42)}
                                                                            var v72 uint8 = 42
                                                                            fmt.Println(v4, v57, v50)
                                                                            _ = v70
                                                                            _ = v71
                                                                            _ = v72
                                                                        } else {
                                                                            var v73 int32 = -996
                                                                            _ = v73
                                                                        }
                                                                        if false {
                                                                            v74 := GenS2{F0: (!true)}
                                                                            var v75 uint16 = 2
                                                                            _ = v74
                                                                            _ = v75
                                                                        }
                                                                        v50 -= 1.5
                                                                    }
                                                                    _ = v57
                                                                }
                                                            } else {
                                                                fmt.Println(fi49, v50, v4)
                                                                v76 := GenS2{F0: true}
                                                                _ = v76
                                                            }
                                                            switch (fi49 / ((fi29) | 1)) {
                                                            case -162:
                                                                if (!((((!true) && (!true)) || true) || true)) {
                                                                    if false {
                                                                        var v77 float32 = float32(v2)
                                                                        _ = v77
                                                                    } else {
                                                                        v2 = v2
                                                                        if (0 < 0) {
                                                                            switch fi29 {
                                                                            case 128:
                                                                                v4 += (^(v4))
                                                                            case 72:
                                                                                var v78 int8 = -127
                                                                                _ = v78
                                                                            case 0:
                                                                                if false {
                                                                                    switch v6 {
                                                                                    case "a":
                                                                                        v4 -= (v4 >> ((0) & 63))
                                                                                    case "":
                                                                                        v79 := GenS1{F0: 65535, F1: 2147483647, F2: 2}
                                                                                        _ = v79
                                                                                    }
                                                                                    var v80 float64 = ((float64(v50) / (v2 / ((v2 + (v2 - v2)) / v2))) / v2)
                                                                                    fmt.Println(v50, fi29, v50)
                                                                                    _ = v80
                                                                                } else {
                                                                                    fmt.Println(v4, v4, v6)
                                                                                }
                                                                            }
                                                                            fmt.Println(v4, v50)
                                                                            if (9 <= 100) {
                                                                                switch v6 {
                                                                                case "":
                                                                                    switch (fi49 / ((127) | 1)) {
                                                                                    case 42:
                                                                                        if true {
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = v6
                                                                                            }
                                                                                            var v81 uint = 997
                                                                                            var v82 bool = false
                                                                                            _ = v81
                                                                                            _ = v82
                                                                                        }
                                                                                    case 118:
                                                                                        v83 := GenS1{F0: 42, F1: 42, F2: 2147483648}
                                                                                        _ = v83
                                                                                    default:
                                                                                        if false {
                                                                                            v50 = (v50 / (v50 - (v50 + v50)))
                                                                                            var v84 bool = (fi49 < 128)
                                                                                            var v85 string = ((v6 + ("" + "")) + v6)
                                                                                            _ = v84
                                                                                            _ = v85
                                                                                        }
                                                                                    }
                                                                                case "\n\t":
                                                                                    var v86 uint32 = 42
                                                                                    _ = v86
                                                                                case "hello":
                                                                                    var v87 int16 = 1000
                                                                                    _ = v87
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    v88 := v1[("foo bar" + v6)]
                                                                    if false {
                                                                        switch (fi29 - fi49) {
                                                                        case -934:
                                                                            v88 = ("" + v88)
                                                                        case 960:
                                                                            var v89 uint8 = 68
                                                                            _ = v89
                                                                        }
                                                                        if ((v4 & v4) == v4) {
                                                                            fmt.Println(v6, fi49, v88)
                                                                            switch fi29 {
                                                                            case 127:
                                                                                v3 = append(v3, "hello")
                                                                            case 513:
                                                                                var v90 float32 = v50
                                                                                _ = v90
                                                                            case 32767:
                                                                                v6 = ((((v6 + v88) + ("\n\t" + (((v88 + "foo bar") + ((((v88 + ("foo bar" + "a")) + v6) + v6) + (v6 + v6))) + ((("hello" + v6) + "hello") + v6)))) + ("\n\t" + (v88 + "\n\t"))) + v88)
                                                                            }
                                                                            var v91 int = 811
                                                                            v92 := v1["foo bar"]
                                                                            _ = v91
                                                                            _ = v92
                                                                        } else {
                                                                            var v93 int8 = 101
                                                                            _ = v93
                                                                        }
                                                                        var v94 int = 42
                                                                        var v95 uint64 = v4
                                                                        v96 := GenS1{F0: 2147483648, F1: 42, F2: 4294967295}
                                                                        _ = v94
                                                                        _ = v95
                                                                        _ = v96
                                                                    } else {
                                                                        v97 := GenS1{F0: 1, F1: -1, F2: 384}
                                                                        var v98 int64 = -827
                                                                        if (-32768 != 668) {
                                                                            var v99 uint = 42
                                                                            _ = v99
                                                                        } else {
                                                                            switch 0 {
                                                                            case 4294967295:
                                                                                v3 = append(v3, "\n\t")
                                                                            case -2147483648:
                                                                                if (uint8(v4) > 127) {
                                                                                    v100 := GenS1{F0: 0, F1: 2147483647, F2: 2147483648}
                                                                                    var v101 int32 = 42
                                                                                    v4 += (^(v4))
                                                                                    v102 := GenS1{F0: 888, F1: (v101 + 2147483647), F2: uint32(v4)}
                                                                                    var v103 float32 = 3.14159
                                                                                    var v104 int32 = (v101 << ((0) & 63))
                                                                                    _ = v100
                                                                                    _ = v101
                                                                                    _ = v102
                                                                                    _ = v103
                                                                                    _ = v104
                                                                                } else {
                                                                                    if (!true) {
                                                                                        v105 := GenS3{F0: v98, F1: 42}
                                                                                        _ = v105
                                                                                    }
                                                                                    fmt.Println(fi49)
                                                                                    var v106 int64 = v98
                                                                                    _ = v106
                                                                                }
                                                                            default:
                                                                                v98 += (-(v98))
                                                                            }
                                                                        }
                                                                        var v107 string = "hello"
                                                                        switch fi49 {
                                                                        case 2147483647:
                                                                            if true {
                                                                                fmt.Println(fi29, fi49)
                                                                                v88 = ("\n\t" + v88)
                                                                                v108 := GenS1{F0: 2147483648, F1: -1, F2: 0}
                                                                                _ = v108
                                                                            }
                                                                        case -481:
                                                                            v88 = v6
                                                                        case 285:
                                                                            switch (fi29 + fi49) {
                                                                            case 42:
                                                                                v50 = 1.5
                                                                            case 4294967295:
                                                                                v88 = v88
                                                                            default:
                                                                                v1[v88] = v88
                                                                                if ((42 > 2147483648) || false) {
                                                                                    v109 := GenS1{F0: 2147483647, F1: 2147483647, F2: 4294967295}
                                                                                    _ = v109
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v107
                                                                        _ = v97
                                                                        _ = v98
                                                                    }
                                                                    var v110 int64 = 1
                                                                    _ = v110
                                                                    _ = v88
                                                                }
                                                            case 32768:
                                                                var v111 int = fi29
                                                                _ = v111
                                                            default:
                                                                var v112 uint8 = 127
                                                                _ = v112
                                                            }
                                                            _ = v54
                                                        }
                                                        var v113 int64 = 9223372036854775807
                                                        _ = fi49
                                                        _ = v113
                                                        _ = v50
                                                    }
                                                }
                                                if (-9223372036854775808 < 950) {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                var v114 int64 = 0
                                                _ = v114
                                            } else {
                                                fmt.Println(v2, v4)
                                                var v115 uint16 = 32767
                                                if true {
                                                    var v116 bool = false
                                                    v2 = (0.1 / v2)
                                                    _ = v116
                                                } else {
                                                    for fi117 := 0; fi117 < 1; fi117++ {
                                                        if len(v3) > 0 {
                                                            v3[0] = ((((("hello" + ("a" + ("\n\t" + v6))) + (("" + "") + ("" + v6))) + (v6 + "a")) + "\n\t") + v6)
                                                        }
                                                        fmt.Println(fi117)
                                                        v118 := GenS3{F0: 1, F1: 100}
                                                        var v119 int16 = -32768
                                                        v120 := v5[(fi117 << ((4294967295) & 63))]
                                                        var v121 string = v120
                                                        v122 := GenS4{F0: 127, F1: false}
                                                        v4 -= (^(v4))
                                                        _ = fi117
                                                        _ = v118
                                                        _ = v119
                                                        _ = v120
                                                        _ = v121
                                                        _ = v122
                                                    }
                                                    v6 = v6
                                                    v2 -= v2
                                                }
                                                _ = v115
                                            }
                                        }
                                        delete(v5, (fi29 + -172))
                                    } else {
                                        if false {
                                            if true {
                                                v6 = v6
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            } else {
                                                v6 = ""
                                                if true {
                                                    v123 := GenS1{F0: 0, F1: -2147483648, F2: 2147483648}
                                                    v124 := GenS3{F0: -894, F1: 0}
                                                    _ = v123
                                                    _ = v124
                                                } else {
                                                    delete(v1, v6)
                                                    v2 -= 1e10
                                                    v6 = v6
                                                }
                                            }
                                            var v125 uint16 = 2
                                            _ = v125
                                        } else {
                                            if false {
                                                v6 = "a"
                                                v126 := GenS4{F0: 11, F1: true}
                                                var v127 uint32 = 137
                                                _ = v126
                                                _ = v127
                                            }
                                            var v128 uint64 = (v4 << ((65535) & 63))
                                            _ = len(v3)
                                            _ = v128
                                        }
                                    }
                                    var v129 uint32 = 2147483648
                                    _ = v129
                                } else {
                                    v6 = v6
                                }
                                _ = fi29
                            }
                        }
                    }
                    v2 = 1e10
                    fmt.Println(v4, v4, v6)
                } else {
                    v4 = 42
                }
                switch 34 {
                case -797:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 246:
                    var v130 uint8 = 2
                    _ = v130
                case 579:
                    fmt.Println(v6)
                }
            } else {
                fmt.Println(v4, v4)
                v4 += (v4 * v4)
                delete(v1, v6)
            }
        default:
            switch -524 {
            case 32767:
                switch "a" {
                case "hello":
                    for fi131 := 0; fi131 < 13; fi131++ {
                        v4 += v4
                        switch (fi131 & (fi131 % ((fi131) | 1))) {
                        case -474:
                            var v132 float64 = v2
                            _ = v132
                        default:
                            var v133 uint64 = v4
                            _ = v133
                        }
                        var v134 int16 = 1
                        _ = fi131
                        _ = v134
                    }
                }
            case 0:
                var v135 float64 = (float64(1.5) / v2)
                _ = v135
            }
        }
        v6 = v6
    }
    var v136 int64 = -553
    for fi137 := 0; fi137 < 18; fi137++ {
        var v138 int32 = 822
        if true {
            if false {
                var v139 uint64 = 397
                v140 := GenS3{F0: v136, F1: 1}
                v138 += (v138 >> ((18446744073709551615) & 63))
                _ = v139
                _ = v140
            }
            for fi141 := 0; fi141 < 3; fi141++ {
                var v142 uint8 = 1
                v143 := v1[(("\n\t" + v6) + v6)]
                fmt.Println(v143, v143)
                _ = fi141
                _ = v142
                _ = v143
            }
            v144 := GenS3{F0: (^(v136)), F1: 42}
            _ = v144
        } else {
            for fi145 := 0; fi145 < 15; fi145++ {
                var v146 int64 = (v136 >> ((844) & 63))
                fmt.Println(v138, fi145, v138)
                v2 = float64(float32(float64((-2.25 * (3.14159 + ((float32(0.1) * 0.0) * (-2.25 + (float32(float64(((1.5 * 3.14159) - 1.5))) / 1.5))))))))
                _ = fi145
                _ = v146
            }
            switch fi137 {
            case 42:
                var v147 float32 = -2.25
                _ = v147
            case 1:
                if false {
                    var v148 uint = 255
                    v4 -= 369
                    _ = v148
                }
            }
        }
        _ = fi137
        _ = v138
    }
    for fi149 := 0; fi149 < 8; fi149++ {
        switch v6 {
        case "hello":
            v150 := GenS5{F0: 155}
            _ = v150
        default:
            delete(v1, "foo bar")
        }
        for fi151 := 0; fi151 < 16; fi151++ {
            var v152 uint32 = 2
            v1["hello"] = "foo bar"
            _ = fi151
            _ = v152
        }
        _ = fi149
    }
    switch 767 {
    case 9223372036854775807:
        v5[-1] = ("hello" + "")
    case 42:
        for fi153 := 0; fi153 < 18; fi153++ {
            v6 = ""
            v154 := GenS6{F0: (true || (469 >= 42)), F1: 1}
            _ = fi153
            _ = v154
        }
    case -1:
        v155 := GenS7{F0: 42}
        _ = v155
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
