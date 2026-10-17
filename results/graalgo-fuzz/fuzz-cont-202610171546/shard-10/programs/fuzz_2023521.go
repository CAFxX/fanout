package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 uint16; F2 int8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS9 struct { F0 int32; F1 string }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int32; F1 uint8 }
type GenS13 struct { F0 int64; F1 bool; F2 uint }
type GenS14 struct { F0 string }
type GenS15 struct { F0 uint64; F1 int64 }

func main() {
    var v1 bool = (true || (!false))
    v2 := []string{"hello"}
    var v3 int16 = 32767
    var v4 uint = 90
    var v5 uint32 = 804
    if true {
        v2 = append(v2, ("a" + ("" + "a")))
    }
    var v6 uint64 = 0
    if (884 >= 255) {
        v4 -= (v4 | v4)
        for fi7 := 0; fi7 < 2; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                switch fi8 {
                case -9223372036854775808:
                    v4 = (^(v4))
                case 1:
                    var v9 int32 = 42
                    _ = v9
                case 42:
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                }
                _ = fi8
            }
            for fi10 := 0; fi10 < 3; fi10++ {
                v5 = (v5 * (v5 * v5))
                _ = fi10
            }
            _ = fi7
        }
    }
    if v1 {
        if true {
            for fi11 := 0; fi11 < 2; fi11++ {
                fmt.Println(v6, v4, v5)
                v2 = append(v2, "foo bar")
                _ = fi11
            }
            for fi12 := 0; fi12 < 3; fi12++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi12
            }
        } else {
            var v13 int = 4294967295
            _ = v13
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS1{F0: -70}
            switch "foo bar" {
            case "a":
                v6 = 4294967295
            case "":
                for fi16 := 0; fi16 < 3; fi16++ {
                    v17 := GenS1{F0: int8(v5)}
                    _ = fi16
                    _ = v17
                }
            case "hello":
                fmt.Println(v5)
            default:
                v18 := GenS2{F0: (v3 % ((42) | 1))}
                _ = v18
            }
            for fi19 := 0; fi19 < 15; fi19++ {
                fmt.Println(v6, fi19)
                v20 := GenS3{F0: 42}
                _ = fi19
                _ = v20
            }
            _ = fi14
            _ = v15
        }
    }
    if v1 {
        var v21 bool = true
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if v1 {
            for fi22 := 0; fi22 < 12; fi22++ {
                if ((!(v1 && ((1e10 != -2.25) || v1))) || false) {
                    switch "a" {
                    case "a":
                        if true {
                            switch "hello" {
                            case "":
                                v23 := GenS1{F0: -1}
                                _ = v23
                            default:
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                            }
                            v3 += 32767
                            fmt.Println(v4)
                        }
                    case "foo bar":
                        v5 = v5
                    default:
                        var v24 uint16 = 32767
                        _ = v24
                    }
                    switch fi22 {
                    case 9223372036854775807:
                        if v21 {
                            var v25 string = "a"
                            var v26 int = -646
                            _ = v25
                            _ = v26
                        } else {
                            for fi27 := 0; fi27 < 15; fi27++ {
                                v3 -= v3
                                switch (fi22 | (-(fi22))) {
                                case 127:
                                    if v1 {
                                        fmt.Println(v3, v6, v5)
                                        if true {
                                            v1 = v21
                                        } else {
                                            if (false && (!(v21 && true))) {
                                                v28 := GenS1{F0: 1}
                                                var v29 int8 = -76
                                                v6 -= v6
                                                _ = v28
                                                _ = v29
                                            } else {
                                                v21 = ("a" > ("hello" + "\n\t"))
                                            }
                                            if ("" < ("hello" + (("a" + ("" + ("foo bar" + "hello"))) + ""))) {
                                                fmt.Println(v4, fi27, v3)
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                if false {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    if ((436 > v6) || v1) {
                                                        v30 := GenS4{F0: 999, F1: v3}
                                                        var v31 int8 = 100
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                } else {
                                                    if len(v2) > 0 {
                                                        v2[0] = ("foo bar" + ("\n\t" + "foo bar"))
                                                    }
                                                    v1 = (-2147483648 != (-(fi27)))
                                                    var v32 uint32 = v5
                                                    if (v3 < v3) {
                                                        var v33 string = "foo bar"
                                                        v6 = uint64(v4)
                                                        if len(v2) > 0 {
                                                            v2[0] = ("" + (v33 + v33))
                                                        }
                                                        _ = v33
                                                    }
                                                    if (v32 <= 65535) {
                                                        v34 := GenS5{F0: v6}
                                                        v35 := GenS1{F0: -128}
                                                        _ = v34
                                                        _ = v35
                                                    } else {
                                                        switch "a" {
                                                        case "hello":
                                                            var v36 string = ("a" + "hello")
                                                            _ = v36
                                                        case "foo bar":
                                                            var v37 bool = v21
                                                            _ = v37
                                                        }
                                                        fmt.Println(fi22, v21, fi27)
                                                        if ((v4 + 733) <= 2) {
                                                            v38 := GenS6{F0: -9223372036854775808, F1: 941, F2: 33}
                                                            var v39 string = ""
                                                            if (!false) {
                                                                if (128 >= 41) {
                                                                    if v1 {
                                                                        v1 = false
                                                                        if (!(-2.25 != -2.25)) {
                                                                            if (-822 == 243) {
                                                                                v40 := GenS2{F0: v3}
                                                                                _ = v40
                                                                            }
                                                                        }
                                                                        v41 := GenS1{F0: 42}
                                                                        _ = v41
                                                                    } else {
                                                                        v42 := GenS5{F0: 2}
                                                                        v4 = uint(v6)
                                                                        v39 = v39
                                                                        fmt.Println(v39, v21)
                                                                        _ = v42
                                                                    }
                                                                }
                                                                var v43 int32 = 0
                                                                _ = v43
                                                            }
                                                            _ = v38
                                                            _ = v39
                                                        } else {
                                                            switch "foo bar" {
                                                            case "hello":
                                                                switch "\n\t" {
                                                                case "":
                                                                    fmt.Println(v32)
                                                                case "hello":
                                                                    v44 := GenS3{F0: 2}
                                                                    _ = v44
                                                                }
                                                            case "\n\t":
                                                                fmt.Println(v6, v32, v6)
                                                            default:
                                                                v45 := GenS7{F0: -1}
                                                                _ = v45
                                                            }
                                                            v32 = v32
                                                            switch "" {
                                                            case "":
                                                                _ = len(v2)
                                                            case "\n\t":
                                                                v2 = append(v2, "hello")
                                                            default:
                                                                v46 := GenS1{F0: 0}
                                                                _ = v46
                                                            }
                                                            v1 = false
                                                        }
                                                        var v47 string = "hello"
                                                        var v48 float64 = (3.14159 * ((1.5 - (1.5 * (1e10 * (1.5 * (((1e10 / (float64(0.1) / (((1.5 / (float64(((1.5 * (1e10 / 0.1)) * (0.0 / 1.0))) + ((float64(3.14159) + (0.0 - -2.25)) - (0.1 - 1e10)))) * float64((0.1 - ((-2.25 / -2.25) / float32((1.5 + (1.5 + (0.0 + -2.25)))))))) + 0.1))) + (0.1 / 3.14159)) - -2.25))))) - (float64((((((float32(0.1) + 0.1) - 0.0) + (((3.14159 / 1.5) + 1.5) / 1.5)) + ((1e10 + 0.0) / (((float32(float64((((1.5 - -2.25) / 1e10) - (1e10 + 1.5)))) * (0.1 - ((3.14159 * ((0.0 - 3.14159) / 3.14159)) - (3.14159 + -2.25)))) + 0.1) * 0.1))) * 0.0)) / 1.5)))
                                                        v21 = false
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                    v6 = (^(v6))
                                                    _ = v32
                                                }
                                                v1 = (v1 && (v21 || v21))
                                            }
                                            fmt.Println(v21, v5, v6)
                                        }
                                    } else {
                                        v49 := GenS1{F0: int8(v5)}
                                        switch fi27 {
                                        case 4294967295:
                                            if false {
                                                v4 = (^(v4))
                                                fmt.Println(fi27, v6, fi27)
                                                if true {
                                                    var v50 float64 = 1e10
                                                    fmt.Println(v21, v4)
                                                    if (!(!(!(1.5 <= 1.5)))) {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if (1 != -1) {
                                                                var v51 string = ("" + "a")
                                                                var v52 uint = (v4 % (((v4 >> ((5) & 63))) | 1))
                                                                _ = v51
                                                                _ = v52
                                                            } else {
                                                                var v53 bool = (false && true)
                                                                _ = v53
                                                            }
                                                        case "\n\t":
                                                            var v54 float32 = (1e10 / (0.1 - 0.0))
                                                            _ = v54
                                                        case "":
                                                            if (true && (!((!v1) || v21))) {
                                                                v55 := GenS8{F0: (v5 * v5), F1: 485, F2: v6}
                                                                if (!v21) {
                                                                    v56 := GenS8{F0: v5, F1: 124, F2: v6}
                                                                    var v57 int = (fi22 + int(v6))
                                                                    v58 := GenS9{F0: 0, F1: "foo bar"}
                                                                    _ = v56
                                                                    _ = v57
                                                                    _ = v58
                                                                } else {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    if v21 {
                                                                        var v59 uint64 = (v6 + v6)
                                                                        var v60 float32 = -2.25
                                                                        switch 32768 {
                                                                        case 4294967295:
                                                                            var v61 int8 = 1
                                                                            _ = v61
                                                                        default:
                                                                            if len(v2) > 0 {
                                                                                v2[0] = "foo bar"
                                                                            }
                                                                        }
                                                                        _ = v59
                                                                        _ = v60
                                                                    } else {
                                                                        fmt.Println(v1, v3, fi22)
                                                                    }
                                                                }
                                                                var v62 float64 = (v50 * v50)
                                                                var v63 string = ("foo bar" + "hello")
                                                                _ = v55
                                                                _ = v62
                                                                _ = v63
                                                            }
                                                        default:
                                                            var v64 uint64 = v6
                                                            _ = v64
                                                        }
                                                        var v65 uint = uint(v6)
                                                        _ = v65
                                                    }
                                                    _ = v50
                                                } else {
                                                    v21 = (-2147483648 > 0)
                                                    fmt.Println(v5)
                                                }
                                            } else {
                                                var v66 uint8 = 2
                                                if ((1.5 <= -2.25) || (v66 <= 1)) {
                                                    _ = len(v2)
                                                }
                                                if (!v1) {
                                                    if (169 < -990) {
                                                        var v67 int = (fi27 / ((fi22) | 1))
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            fmt.Println(fi27, v66, fi22)
                                                        case "hello":
                                                            v2 = append(v2, "hello")
                                                        default:
                                                            if (-9223372036854775808 > 9223372036854775807) {
                                                                var v68 float32 = float32(((1.5 * (3.14159 / float64(-2.25))) * (0.1 + -2.25)))
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    if true {
                                                                        v1 = (("a" + (("a" + "") + "hello")) == "\n\t")
                                                                        if true {
                                                                            if len(v2) > 0 {
                                                                                v2[0] = ""
                                                                            }
                                                                            _ = len(v2)
                                                                        }
                                                                        if (127 == 100) {
                                                                            var v69 string = "a"
                                                                            v67 = 128
                                                                            if ((v68 - v68) == v68) {
                                                                                var v70 uint64 = v6
                                                                                _ = v70
                                                                            }
                                                                            _ = v69
                                                                        }
                                                                    } else {
                                                                        switch 1 {
                                                                        case 2147483647:
                                                                            v68 = v68
                                                                        case -2147483648:
                                                                            v2 = append(v2, "a")
                                                                            var v71 uint8 = 2
                                                                            _ = v71
                                                                        default:
                                                                            v5 -= 551
                                                                        }
                                                                        if (110 == 255) {
                                                                            fmt.Println(v3, v6, fi22)
                                                                            var v72 uint8 = v66
                                                                            _ = v72
                                                                        } else {
                                                                            var v73 int64 = 208
                                                                            var v74 uint8 = v66
                                                                            _ = v73
                                                                            _ = v74
                                                                        }
                                                                        _ = len(v2)
                                                                    }
                                                                default:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = ("hello" + "a")
                                                                    }
                                                                }
                                                                var v75 int16 = v3
                                                                _ = v68
                                                                _ = v75
                                                            }
                                                        }
                                                        _ = v67
                                                    }
                                                    fmt.Println(v6)
                                                } else {
                                                    v76 := GenS10{F0: -48}
                                                    if len(v2) > 0 {
                                                        v2[0] = ""
                                                    }
                                                    if (v4 == v4) {
                                                        v3 -= v3
                                                    } else {
                                                        if (-2147483648 < 0) {
                                                            v66 += (v66 >> (((v4 - v4)) & 63))
                                                            var v77 uint32 = v5
                                                            switch "" {
                                                            case "foo bar":
                                                                v78 := GenS2{F0: (v3 % (((v3 - (^(v3)))) | 1))}
                                                                _ = v78
                                                            case "\n\t":
                                                                v5 = 2
                                                            }
                                                            switch 0 {
                                                            case 2147483647:
                                                                _ = len(v2)
                                                            }
                                                            fmt.Println(v6, v5)
                                                            _ = len(v2)
                                                            _ = v77
                                                        }
                                                    }
                                                    _ = v76
                                                }
                                                v79 := GenS10{F0: -1}
                                                var v80 int = fi22
                                                _ = v66
                                                _ = v79
                                                _ = v80
                                            }
                                        case 1:
                                            v2 = append(v2, "hello")
                                        }
                                        _ = v49
                                    }
                                case 32767:
                                    v3 = (-(v3))
                                    fmt.Println(v21, v4)
                                    if (int32(v4) == -2147483648) {
                                        var v81 int64 = 1
                                        switch 42 {
                                        case 0:
                                            switch ("hello" + ("foo bar" + "a")) {
                                            case "":
                                                v1 = (v21 || (!(1 <= 42)))
                                            case "a":
                                                var v82 uint = (v4 / ((18446744073709551615) | 1))
                                                _ = v82
                                            case "\n\t":
                                                switch fi22 {
                                                case 857:
                                                    v81 = v81
                                                }
                                            }
                                        default:
                                            var v83 string = "\n\t"
                                            _ = v83
                                        }
                                        v84 := GenS7{F0: 9223372036854775807}
                                        var v85 uint32 = (v5 % ((v5) | 1))
                                        _ = v81
                                        _ = v84
                                        _ = v85
                                    }
                                case 32768:
                                    v86 := GenS1{F0: 127}
                                    _ = v86
                                }
                                v1 = (!(!v21))
                                var v87 uint64 = 2
                                _ = fi27
                                _ = v87
                            }
                        }
                    case 32767:
                        var v88 uint16 = 255
                        _ = v88
                    case 127:
                        v6 += (^(v6))
                    }
                    if true {
                        var v89 uint16 = 55
                        var v90 int16 = 1
                        _ = v89
                        _ = v90
                    }
                }
                _ = fi22
            }
        } else {
            v1 = false
        }
        _ = v21
    } else {
        for fi91 := 0; fi91 < 16; fi91++ {
            if false {
                _ = len(v2)
                for fi92 := 0; fi92 < 11; fi92++ {
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    _ = fi92
                }
                fmt.Println(v6, fi91)
            }
            _ = len(v2)
            if v1 {
                var v93 uint = 1
                if len(v2) > 0 {
                    v2[0] = ("" + (("" + "") + "foo bar"))
                }
                var v94 int32 = -2147483648
                _ = v93
                _ = v94
            }
            _ = fi91
        }
        switch 42 {
        case 966:
            fmt.Println(v6, v6)
        case -779:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 128:
            v95 := GenS11{F0: 42}
            _ = v95
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    for fi96 := 0; fi96 < 9; fi96++ {
        switch "foo bar" {
        case "":
            for fi97 := 0; fi97 < 9; fi97++ {
                if (1 <= 42) {
                    switch "" {
                    case "":
                        if (42 != 1) {
                            var v98 int = fi96
                            switch 1 {
                            case 0:
                                var v99 uint32 = v5
                                _ = v99
                            default:
                                fmt.Println(fi96, v1, v1)
                            }
                            _ = v98
                        } else {
                            if true {
                                fmt.Println(v3, v5, v5)
                            } else {
                                var v100 uint = uint(v5)
                                fmt.Println(v4, v3, v4)
                                _ = v100
                            }
                        }
                    case "\n\t":
                        var v101 int = (fi97 | (fi97 * fi97))
                        _ = v101
                    }
                    v6 -= (v6 % ((v6) | 1))
                    if false {
                        if ((v3 / ((v3) | 1)) != v3) {
                            v3 = 245
                        }
                    } else {
                        var v102 int32 = 42
                        _ = v102
                    }
                    var v103 int16 = (v3 >> ((v4) & 63))
                    v104 := GenS9{F0: -1, F1: "a"}
                    _ = v103
                    _ = v104
                }
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                v105 := GenS9{F0: 2147483647, F1: ("\n\t" + "a")}
                _ = fi97
                _ = v105
            }
        default:
            var v106 int8 = 47
            _ = v106
        }
        _ = fi96
    }
    switch "" {
    case "foo bar":
        var v107 uint8 = 128
        _ = v107
    case "":
        if true {
            _ = len(v2)
            if ((!(42 < 128)) && (!((v5 - 65535) != v5))) {
                if (-62 >= 100) {
                    switch ("hello" + "a") {
                    case "hello":
                        v3 -= (v3 | v3)
                    case "foo bar":
                        if v1 {
                            _ = len(v2)
                            if false {
                                _ = len(v2)
                            } else {
                                var v108 uint8 = 127
                                if (!(-2.25 > -2.25)) {
                                    switch 9223372036854775807 {
                                    case 4294967295:
                                        var v109 int32 = -2147483648
                                        _ = v109
                                    case -2147483648:
                                        switch -2147483648 {
                                        case 4294967295:
                                            v5 = 2147483648
                                        case 42:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case 0:
                                            if v1 {
                                                var v110 bool = v1
                                                _ = v110
                                            }
                                        default:
                                            _ = len(v2)
                                        }
                                    default:
                                        if (!(!(!false))) {
                                            v111 := GenS12{F0: 0, F1: v108}
                                            switch -2147483648 {
                                            case 2147483647:
                                                for fi112 := 0; fi112 < 12; fi112++ {
                                                    v6 = (v6 << ((v4) & 63))
                                                    var v113 uint64 = 1
                                                    var v114 int8 = 123
                                                    _ = fi112
                                                    _ = v113
                                                    _ = v114
                                                }
                                            case 696:
                                                var v115 int = -234
                                                _ = v115
                                            case 1:
                                                v5 = 0
                                            default:
                                                fmt.Println(v3)
                                            }
                                            fmt.Println(v108, v108)
                                            _ = v111
                                        } else {
                                            v116 := GenS13{F0: int64(v5), F1: v1, F2: uint(v5)}
                                            _ = v116
                                        }
                                    }
                                    for fi117 := 0; fi117 < 12; fi117++ {
                                        var v118 uint8 = v108
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = fi117
                                        _ = v118
                                    }
                                }
                                _ = v108
                            }
                        } else {
                            if (!(true && (2 >= 2))) {
                                for fi119 := 0; fi119 < 15; fi119++ {
                                    v3 = int16(v4)
                                    var v120 bool = ((v6 & uint64(v4)) == 18446744073709551615)
                                    if len(v2) > 0 {
                                        v2[0] = "\n\t"
                                    }
                                    _ = fi119
                                    _ = v120
                                }
                            }
                            fmt.Println(v3)
                            v121 := GenS14{F0: "foo bar"}
                            _ = v121
                        }
                    default:
                        var v122 float64 = float64(0.0)
                        _ = v122
                    }
                }
                v6 = (v6 + v6)
                if v1 {
                    v4 += 4294967295
                    v6 += 42
                    var v123 int64 = -294
                    _ = v123
                }
            }
            for fi124 := 0; fi124 < 16; fi124++ {
                v5 = v5
                v125 := GenS15{F0: 553, F1: -297}
                _ = fi124
                _ = v125
            }
        }
    case "\n\t":
        var v126 int8 = 0
        _ = v126
    }
    if (v5 == 461) {
        v1 = v1
    } else {
        fmt.Println(v1, v1, v6)
        switch ("\n\t" + "a") {
        case "foo bar":
            v127 := GenS13{F0: 0, F1: v1, F2: (v4 / (((v4 | (v4 - 609))) | 1))}
            _ = v127
        }
    }
    if len(v2) > 0 {
        v2[0] = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
