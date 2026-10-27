package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 string; F1 uint16 }
type GenS10 struct { F0 uint; F1 int16; F2 string }
type GenS11 struct { F0 uint8; F1 int }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 bool; F1 string; F2 int }

func main() {
    var v1 bool = (!(false || (!(true && (42 == 32767)))))
    v2 := []int{128, 2147483647, 669}
    var v3 uint = 18446744073709551615
    v4 := []int{-60, 1}
    var v5 uint8 = 1
    var v6 uint = v3
    v7 := GenS1{F0: (v5 >> ((1) & 63))}
    var v8 bool = v1
    for fi9 := 0; fi9 < 18; fi9++ {
        if v8 {
            fmt.Println(v6, v8, v6)
            v10 := GenS1{F0: (v5 / (((^(v5))) | 1))}
            fmt.Println(v5, v1)
            _ = v10
        } else {
            for fi11 := 0; fi11 < 7; fi11++ {
                v12 := GenS2{F0: 0}
                var v13 int16 = 1000
                v14 := GenS1{F0: v5}
                v3 += v3
                var v15 int32 = 0
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
                _ = v15
            }
        }
        _ = fi9
    }
    v16 := GenS3{F0: 664}
    v17 := GenS1{F0: 52}
    switch 794 {
    case 42:
        for fi18 := 0; fi18 < 3; fi18++ {
            switch ("" + "") {
            case "foo bar":
                for fi19 := 0; fi19 < 1; fi19++ {
                    var v20 float32 = -2.25
                    _ = fi19
                    _ = v20
                }
            case "a":
                if true {
                    var v21 uint8 = v5
                    _ = v21
                }
            case "":
                v3 = (v6 << ((2) & 63))
            }
            _ = len(v2)
            if ((fi18 + 2147483647) <= (fi18 & -1)) {
                for fi22 := 0; fi22 < 1; fi22++ {
                    switch "a" {
                    case "a":
                        switch fi22 {
                        case 2147483647:
                            if len(v4) > 0 {
                                v4[0] = (fi22 & -2147483648)
                            }
                            if v1 {
                                var v23 uint64 = 42
                                if (!v8) {
                                    switch (fi18 * 9223372036854775807) {
                                    case 9223372036854775807:
                                        var v24 int32 = -2147483648
                                        _ = v24
                                    case -744:
                                        switch fi18 {
                                        case 127:
                                            if v8 {
                                                fmt.Println(fi18, fi18, v3)
                                                fmt.Println(v3)
                                                fmt.Println(v8, v6, v3)
                                            } else {
                                                v25 := GenS2{F0: 0}
                                                v1 = (42 > int64(v6))
                                                _ = v25
                                            }
                                        case 0:
                                            if (true && v1) {
                                                var v26 int64 = 0
                                                _ = v26
                                            } else {
                                                if ("a" != (("hello" + "") + ("hello" + "foo bar"))) {
                                                    v27 := GenS2{F0: 151}
                                                    fmt.Println(v5)
                                                    v5 = v5
                                                    _ = v27
                                                }
                                                if ((float32(-2.25) <= ((1e10 + 3.14159) - ((0.1 * (0.0 - (1e10 / float32((0.1 - 0.0))))) / 0.1))) && (-1 != 9223372036854775807)) {
                                                    switch ("hello" + ("a" + "foo bar")) {
                                                    case "":
                                                        var v28 float64 = 0.1
                                                        _ = v28
                                                    case "foo bar":
                                                        var v29 float64 = 3.14159
                                                        _ = v29
                                                    case "\n\t":
                                                        v30 := GenS4{F0: true, F1: 42}
                                                        _ = v30
                                                    }
                                                } else {
                                                    if ("\n\t" <= "") {
                                                        switch (("hello" + "\n\t") + "\n\t") {
                                                        case "hello":
                                                            var v31 float64 = -2.25
                                                            _ = v31
                                                        case "\n\t":
                                                            v23 += (v23 % ((uint64(v6)) | 1))
                                                        default:
                                                            v6 += v6
                                                        }
                                                        v5 += (^(v5))
                                                    } else {
                                                        fmt.Println(v6, v5)
                                                        var v32 float32 = ((((1e10 / 1.5) + (float32(((0.0 / 1.0) * ((1e10 / 0.1) + (3.14159 * ((((float64(0.0) - 0.1) / (1.5 - 0.0)) + 1e10) / 1.0))))) + 1.5)) + 1.5) / 3.14159)
                                                        _ = v32
                                                    }
                                                    fmt.Println(fi18)
                                                    v1 = (1000 < -149)
                                                }
                                            }
                                        }
                                    }
                                    v33 := GenS5{F0: 42, F1: "foo bar"}
                                    fmt.Println(fi22, fi18)
                                    _ = v33
                                }
                                if len(v2) > 0 {
                                    v2[0] = (fi18 | 32768)
                                }
                                var v34 uint = (v3 % ((1) | 1))
                                v35 := GenS3{F0: -32768}
                                v2 = append(v2, fi22)
                                _ = v23
                                _ = v34
                                _ = v35
                            }
                            v36 := GenS3{F0: -507}
                            _ = v36
                        case 909:
                            if true {
                                v1 = (v8 && v1)
                                v1 = v8
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            switch int(v5) {
                            case 2147483647:
                                var v37 uint64 = 106
                                _ = v37
                            case 32767:
                                var v38 uint8 = v5
                                if v8 {
                                    v2 = append(v2, (fi18 % (((fi22 / (((-(fi22))) | 1))) | 1)))
                                    if false {
                                        var v39 bool = true
                                        _ = v39
                                    } else {
                                        v1 = ((0.0 * 1.5) < 0.0)
                                    }
                                    var v40 string = ("" + "foo bar")
                                    _ = v40
                                } else {
                                    if v8 {
                                        var v41 uint16 = 236
                                        switch (("" + "hello") + "a") {
                                        case "a":
                                            v42 := GenS6{F0: "\n\t"}
                                            _ = v42
                                        default:
                                            v6 -= v6
                                        }
                                        _ = v41
                                    } else {
                                        if v1 {
                                            switch -694 {
                                            case 9223372036854775807:
                                                v8 = v8
                                            case 2147483647:
                                                switch "a" {
                                                case "\n\t":
                                                    if len(v4) > 0 {
                                                        v4[0] = (fi18 >> ((v3) & 63))
                                                    }
                                                case "hello":
                                                    var v43 uint16 = 2
                                                    _ = v43
                                                case "foo bar":
                                                    if (!(!((1e10 != float64(0.0)) && (false && ((v5 != v5) || v8))))) {
                                                        v5 -= (^(v38))
                                                    } else {
                                                        var v44 uint32 = 2
                                                        _ = v44
                                                    }
                                                default:
                                                    v38 -= 0
                                                }
                                            case -1:
                                                var v45 bool = (true && (((float32(1e10) == (-2.25 + 1.5)) || (int8(fi22) == 0)) || v1))
                                                if (!v8) {
                                                    fmt.Println(v3)
                                                } else {
                                                    fmt.Println(v45, fi22, v1)
                                                    if len(v2) > 0 {
                                                        v2[0] = (^(fi22))
                                                    }
                                                    if v8 {
                                                        _ = len(v4)
                                                        var v46 int = 32767
                                                        _ = v46
                                                    } else {
                                                        if ((2 != 2) || ((0.1 > 1.5) || false)) {
                                                            if (!v45) {
                                                                v47 := GenS3{F0: 0}
                                                                _ = v47
                                                            } else {
                                                                v48 := GenS5{F0: -1, F1: ("\n\t" + "hello")}
                                                                v49 := GenS2{F0: -209}
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                        }
                                                        switch fi22 {
                                                        case 1:
                                                            var v50 bool = (v8 || (1 <= 32767))
                                                            _ = v50
                                                        case 906:
                                                            if (false && false) {
                                                                if (2147483648 > 4294967295) {
                                                                    if true {
                                                                        fmt.Println(fi22, fi18, v38)
                                                                    }
                                                                    v45 = (v5 == (^(v5)))
                                                                    v4 = append(v4, fi18)
                                                                    switch (((((((("" + ("foo bar" + ("a" + ""))) + ("hello" + "a")) + "hello") + (("a" + "") + (((("" + "a") + ("\n\t" + ("hello" + "a"))) + "") + ("foo bar" + (("foo bar" + ("a" + "foo bar")) + ("" + "foo bar")))))) + ("hello" + ("hello" + "\n\t"))) + "hello") + "\n\t") + "a") {
                                                                    case "hello":
                                                                        var v51 bool = (!((!v1) && (0 != 21)))
                                                                        _ = v51
                                                                    case "":
                                                                        switch ("foo bar" + "\n\t") {
                                                                        case "":
                                                                            v38 -= (v38 % (((v38 << (((v6 - v3)) & 63))) | 1))
                                                                        case "\n\t":
                                                                            if true {
                                                                                v5 -= v38
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = (fi22 | fi18)
                                                                                }
                                                                                if v45 {
                                                                                    var v52 bool = (!(((!v8) && (v45 && (42 > 127))) || (v45 && v1)))
                                                                                    fmt.Println(v52, fi22, v38)
                                                                                    v4 = append(v4, (fi18 | 32767))
                                                                                    var v53 int = -324
                                                                                    _ = v52
                                                                                    _ = v53
                                                                                } else {
                                                                                    switch ("a" + "\n\t") {
                                                                                    case "\n\t":
                                                                                        if (v3 > 4294967295) {
                                                                                            fmt.Println(v1, v1, fi22)
                                                                                            var v54 float64 = -2.25
                                                                                            _ = v54
                                                                                        } else {
                                                                                            if false {
                                                                                                var v55 int32 = -2147483648
                                                                                                _ = v55
                                                                                            }
                                                                                            v5 += (^(v38))
                                                                                            if (v8 || true) {
                                                                                                var v56 int8 = 0
                                                                                                if len(v4) > 0 {
                                                                                                    v4[0] = (fi18 & fi22)
                                                                                                }
                                                                                                _ = v56
                                                                                            } else {
                                                                                                var v57 int32 = -908
                                                                                                var v58 uint8 = (^(v38))
                                                                                                v58 -= 27
                                                                                                _ = v57
                                                                                                _ = v58
                                                                                            }
                                                                                            switch "hello" {
                                                                                            case "":
                                                                                                var v59 int16 = 1
                                                                                                _ = v59
                                                                                            default:
                                                                                                var v60 int64 = 844
                                                                                                _ = v60
                                                                                            }
                                                                                            v3 = (v6 ^ 756)
                                                                                        }
                                                                                    default:
                                                                                        v38 -= 128
                                                                                    }
                                                                                    if (v1 || (!v1)) {
                                                                                        if (0.1 == ((float64((float32(0.1) - 1.5)) / float64((float32(1e10) + 3.14159))) / float64(1e10))) {
                                                                                            v61 := GenS6{F0: ("hello" + "hello")}
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = fi22
                                                                                            }
                                                                                            fmt.Println(v8, v45)
                                                                                            var v62 int8 = 0
                                                                                            _ = v61
                                                                                            _ = v62
                                                                                        }
                                                                                        if (("a" > "\n\t") && (61 < 42)) {
                                                                                            if v8 {
                                                                                                v3 -= 0
                                                                                            } else {
                                                                                                if v45 {
                                                                                                    if v45 {
                                                                                                        var v63 uint32 = 2
                                                                                                        var v64 int = (fi18 - 32767)
                                                                                                        if len(v4) >= 2 {
                                                                                                            v4 = v4[1:len(v4)]
                                                                                                        }
                                                                                                        _ = v63
                                                                                                        _ = v64
                                                                                                    }
                                                                                                }
                                                                                                v65 := GenS6{F0: ("a" + "foo bar")}
                                                                                                _ = v65
                                                                                            }
                                                                                            var v66 uint32 = 379
                                                                                            v67 := GenS7{F0: uint64(v38)}
                                                                                            var v68 float64 = 0.0
                                                                                            fmt.Println(v5, v1)
                                                                                            v45 = (!v45)
                                                                                            var v69 uint8 = 2
                                                                                            _ = v66
                                                                                            _ = v67
                                                                                            _ = v68
                                                                                            _ = v69
                                                                                        } else {
                                                                                            var v70 uint32 = 375
                                                                                            _ = v70
                                                                                        }
                                                                                        switch ("" + "a") {
                                                                                        case "":
                                                                                            v5 += (^(v38))
                                                                                        default:
                                                                                            var v71 uint32 = uint32(v38)
                                                                                            _ = v71
                                                                                        }
                                                                                        v6 += (v6 ^ (v3 % (((^(v3))) | 1)))
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                _ = len(v4)
                                                                                _ = len(v4)
                                                                            }
                                                                        default:
                                                                            switch ("foo bar" + "hello") {
                                                                            case "":
                                                                                var v72 uint8 = (v5 - 127)
                                                                                _ = v72
                                                                            case "a":
                                                                                var v73 int64 = 0
                                                                                _ = v73
                                                                            case "foo bar":
                                                                                var v74 uint8 = (v38 * 255)
                                                                                _ = v74
                                                                            default:
                                                                                var v75 uint32 = 2
                                                                                _ = v75
                                                                            }
                                                                        }
                                                                    default:
                                                                        var v76 float32 = 0.1
                                                                        _ = v76
                                                                    }
                                                                    v45 = (122 <= (v6 + v3))
                                                                } else {
                                                                    switch "a" {
                                                                    case "a":
                                                                        var v77 bool = v1
                                                                        _ = v77
                                                                    default:
                                                                        v3 = (v6 << (((v3 * v6)) & 63))
                                                                        fmt.Println(v6, v45)
                                                                    }
                                                                    if true {
                                                                        var v78 uint64 = 42
                                                                        _ = v78
                                                                    } else {
                                                                        v1 = v1
                                                                        if (!v45) {
                                                                            var v79 uint32 = 4294967295
                                                                            var v80 int = (fi22 ^ (fi18 + int(v79)))
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                        if false {
                                                                            v81 := GenS8{F0: fi22, F1: 1}
                                                                            switch (("" + "a") + "foo bar") {
                                                                            case "":
                                                                                fmt.Println(v5, v1, v1)
                                                                            case "foo bar":
                                                                                v8 = (0.1 < (-2.25 - 1e10))
                                                                            case "\n\t":
                                                                                fmt.Println(v38)
                                                                            }
                                                                            _ = v81
                                                                        }
                                                                    }
                                                                }
                                                                var v82 int = fi22
                                                                _ = v82
                                                            } else {
                                                                v8 = v1
                                                                v83 := GenS9{F0: "\n\t", F1: 32767}
                                                                _ = v83
                                                            }
                                                        default:
                                                            v1 = (float64(((0.1 * (0.1 / 3.14159)) * 3.14159)) != 1e10)
                                                            fmt.Println(v45)
                                                        }
                                                        var v84 float32 = 1e10
                                                        if len(v4) > 0 {
                                                            v4[0] = fi18
                                                        }
                                                        _ = v84
                                                    }
                                                }
                                                _ = v45
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v85 := GenS10{F0: v6, F1: 42, F2: ""}
                                            _ = v85
                                        } else {
                                            if ((false || v1) || v1) {
                                                v4 = append(v4, (fi18 + fi22))
                                                if true {
                                                    fmt.Println(v5, fi18)
                                                }
                                                switch "\n\t" {
                                                case "\n\t":
                                                    switch (fi18 * -1) {
                                                    case -9223372036854775808:
                                                        switch "foo bar" {
                                                        case "hello":
                                                            switch fi18 {
                                                            case 9223372036854775807:
                                                                var v86 uint16 = 936
                                                                _ = v86
                                                            case 422:
                                                                switch "foo bar" {
                                                                case "a":
                                                                    if (128 != v38) {
                                                                        v8 = (false && (0 < 1))
                                                                        v3 = v3
                                                                    }
                                                                case "hello":
                                                                    v87 := GenS5{F0: -1, F1: "foo bar"}
                                                                    _ = v87
                                                                }
                                                            default:
                                                                switch 127 {
                                                                case 42:
                                                                    fmt.Println(v38, v1, v1)
                                                                default:
                                                                    var v88 uint = v6
                                                                    _ = v88
                                                                }
                                                            }
                                                        case "\n\t":
                                                            var v89 int64 = -1
                                                            _ = v89
                                                        case "":
                                                            var v90 uint16 = 65535
                                                            _ = v90
                                                        }
                                                    case 128:
                                                        v91 := GenS11{F0: (^(v38)), F1: (fi22 << (((v3 ^ v6)) & 63))}
                                                        _ = v91
                                                    default:
                                                        v4 = append(v4, (-(fi22)))
                                                    }
                                                    if v8 {
                                                        v5 -= v5
                                                        var v92 int8 = 12
                                                        _ = v92
                                                    } else {
                                                        var v93 uint = v6
                                                        var v94 float32 = (((0.1 + 0.1) / 1.5) + (3.14159 + 0.1))
                                                        var v95 string = "foo bar"
                                                        _ = v93
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                case "":
                                                    switch "a" {
                                                    case "":
                                                        var v96 uint16 = 1
                                                        _ = v96
                                                    case "foo bar":
                                                        if (-2.25 <= 1.5) {
                                                            var v97 float32 = float32(((3.14159 + 0.0) + float64(-2.25)))
                                                            v4 = append(v4, (-(fi18)))
                                                            v8 = (false && (v8 || (!false)))
                                                            var v98 float32 = v97
                                                            v99 := GenS12{F0: v5}
                                                            _ = v97
                                                            _ = v98
                                                            _ = v99
                                                        }
                                                    case "hello":
                                                        switch "foo bar" {
                                                        case "a":
                                                            switch fi22 {
                                                            case 1:
                                                                _ = len(v2)
                                                            case 42:
                                                                v2 = append(v2, fi18)
                                                            default:
                                                                fmt.Println(v3, v5)
                                                            }
                                                        }
                                                    }
                                                }
                                                var v100 uint16 = 2
                                                fmt.Println(v6, v3, v38)
                                                _ = v100
                                            } else {
                                                v3 = 18446744073709551615
                                            }
                                            var v101 int16 = -1
                                            var v102 int16 = v101
                                            _ = v101
                                            _ = v102
                                        }
                                        var v103 uint8 = v38
                                        _ = v103
                                    }
                                    var v104 int32 = 2147483647
                                    if v1 {
                                        if (v1 || (-98 != -16)) {
                                            if (2 > 0) {
                                                var v105 uint64 = 42
                                                _ = v105
                                            }
                                            v106 := GenS10{F0: v6, F1: -941, F2: "hello"}
                                            _ = v106
                                        }
                                    } else {
                                        v107 := GenS13{F0: (1.5 > 3.14159)}
                                        _ = v107
                                    }
                                    _ = v104
                                }
                                v108 := GenS6{F0: ""}
                                _ = len(v2)
                                _ = v108
                                _ = v38
                            }
                            var v109 string = ("" + "a")
                            v110 := GenS8{F0: (fi22 % (((-(fi22))) | 1)), F1: -32768}
                            _ = v109
                            _ = v110
                        default:
                            v4 = append(v4, (fi22 + (fi18 & (fi22 >> ((511) & 63)))))
                        }
                    case "foo bar":
                        v111 := GenS3{F0: 1}
                        _ = v111
                    default:
                        v5 = 255
                    }
                    _ = fi22
                }
            } else {
                if v8 {
                    var v112 uint = (v3 * v3)
                    _ = v112
                }
                fmt.Println(v8)
            }
            _ = fi18
        }
    case -665:
        _ = len(v4)
    default:
        v113 := GenS10{F0: (^(v3)), F1: 1, F2: ("foo bar" + "hello")}
        _ = v113
    }
    v114 := GenS14{F0: v5}
    for fi115 := 0; fi115 < 18; fi115++ {
        var v116 int32 = -2147483648
        if (!(v1 || (1e10 == ((1.5 + 1e10) + 1.5)))) {
            _ = len(v2)
        }
        _ = fi115
        _ = v116
    }
    _ = len(v4)
    v8 = (42 > 2147483647)
    switch 127 {
    case -838:
        v117 := GenS14{F0: (v5 & (v5 - uint8(v6)))}
        _ = v117
    case 32767:
        v118 := GenS15{F0: v1, F1: (("a" + "") + ""), F2: 4294967295}
        _ = v118
    case -104:
        v4 = append(v4, 127)
    default:
        v3 += v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v114
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
