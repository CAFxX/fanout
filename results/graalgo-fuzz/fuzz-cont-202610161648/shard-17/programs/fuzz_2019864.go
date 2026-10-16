package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32; F1 uint8; F2 bool }
type GenS8 struct { F0 int64; F1 uint8; F2 string }
type GenS9 struct { F0 uint8; F1 uint64; F2 string }
type GenS10 struct { F0 uint16; F1 uint }
type GenS11 struct { F0 uint32; F1 int; F2 int64 }
type GenS12 struct { F0 int16; F1 uint }
type GenS13 struct { F0 uint32; F1 int32 }
type GenS14 struct { F0 int16; F1 int64; F2 int16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{42, 378, 1, 1}
    v3 := make(map[string]uint32)
    var v4 int = 1
    v4 = v4
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS1{F0: 0, F1: 2}
        _ = fi5
        _ = v6
    }
    if (244 < 42) {
        v7 := GenS1{F0: 42, F1: 4294967295}
        fmt.Println(v4)
        var v8 int = (v4 | 32767)
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            v4 = (fi9 >> ((65) & 63))
            v2 = append(v2, 65535)
            _ = fi9
        }
        fmt.Println(v4)
        switch v4 {
        case 1:
            var v10 bool = true
            _ = v10
        case 4294967295:
            for fi11 := 0; fi11 < 18; fi11++ {
                fmt.Println(v4, v4)
                var v12 float32 = float32((1.5 * float64(-2.25)))
                _ = fi11
                _ = v12
            }
        default:
            switch (v4 >> ((187) & 63)) {
            case 42:
                switch ("" + "a") {
                case "\n\t":
                    v13 := GenS1{F0: -685, F1: 2147483648}
                    _ = v13
                case "hello":
                    switch 1 {
                    case 2147483647:
                        if true {
                            var v14 uint64 = 42
                            _ = v14
                        }
                    default:
                        var v15 int = v4
                        _ = v15
                    }
                default:
                    v16 := v3[("hello" + ("a" + "a"))]
                    _ = v16
                }
            case 32768:
                delete(v1, 1)
            case -9223372036854775808:
                v4 = v4
            }
        }
    }
    var v17 int64 = int64(v4)
    if true {
        v18 := GenS1{F0: v17, F1: 2147483648}
        _ = v18
    }
    fmt.Println(v4, v4)
    v19 := GenS2{F0: 0}
    switch 32768 {
    case 42:
        v17 -= v17
    case 394:
        v4 = (^(v4))
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        if false {
            if false {
                switch (fi20 ^ 32767) {
                case 522:
                    fmt.Println(v4, fi20, v17)
                case 922:
                    v2 = append(v2, 1)
                case 4294967295:
                    fmt.Println(fi20, v4)
                }
            }
        }
        if true {
            var v21 uint64 = 42
            for fi22 := 0; fi22 < 15; fi22++ {
                if ((65535 == 65535) || false) {
                    fmt.Println(v4, v17, v17)
                } else {
                    if (3.14159 < (-2.25 + float32(3.14159))) {
                        fmt.Println(v17, fi22)
                        if (false || true) {
                            switch 872 {
                            case 128:
                                v1[1] = false
                            }
                            var v23 int8 = 100
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if ((127 > (v23 ^ -1)) || false) {
                                if false {
                                    var v24 float64 = (((1.5 * float64((0.1 / 1.0))) / float64(3.14159)) - float64(0.1))
                                    if ((!(((float32(1e10) * 3.14159) + -2.25) <= 1.5)) && ((uint16(v21) != 42) && false)) {
                                        v17 += (-(v17))
                                    } else {
                                        _ = len(v2)
                                    }
                                    fmt.Println(v4, fi22)
                                    _ = v24
                                }
                                if true {
                                    var v25 int32 = 0
                                    if false {
                                        var v26 int16 = 0
                                        var v27 uint = 814
                                        switch "hello" {
                                        case "foo bar":
                                            var v28 int16 = v26
                                            _ = v28
                                        case "a":
                                            if false {
                                                fmt.Println(v26, v4)
                                                v29 := GenS1{F0: v17, F1: 65535}
                                                switch -52 {
                                                case 127:
                                                    if false {
                                                        var v30 int32 = v25
                                                        var v31 int64 = (v17 + int64(v4))
                                                        v4 -= -834
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                default:
                                                    if (!(2147483647 != 4294967295)) {
                                                        var v32 int32 = v25
                                                        fmt.Println(v4)
                                                        _ = v32
                                                    }
                                                }
                                                if false {
                                                    v33 := GenS1{F0: (v17 / ((v17) | 1)), F1: 814}
                                                    var v34 string = (("a" + "a") + "")
                                                    v26 -= int16(v27)
                                                    _ = v33
                                                    _ = v34
                                                } else {
                                                    var v35 int32 = -1
                                                    v36 := v1[2]
                                                    _ = v35
                                                    _ = v36
                                                }
                                                v37 := GenS3{F0: v25}
                                                var v38 int64 = -1
                                                switch (fi20 << ((42) & 63)) {
                                                case 128:
                                                    if (224 <= 0) {
                                                        v39 := GenS4{F0: uint32(v21)}
                                                        v21 += (v21 + 18446744073709551615)
                                                        v23 = v23
                                                        var v40 bool = (0 != 141)
                                                        _ = v39
                                                        _ = v40
                                                    } else {
                                                        var v41 uint32 = 783
                                                        if true {
                                                            _ = len(v3)
                                                            v4 -= (^(fi22))
                                                            if (("a" + "hello") != ("\n\t" + "")) {
                                                                _ = len(v2)
                                                                var v42 int16 = -215
                                                                if (!true) {
                                                                    v43 := v3[""]
                                                                    v38 = 9223372036854775807
                                                                    v23 -= (^(v23))
                                                                    _ = v43
                                                                } else {
                                                                    if ((100 < v23) && (true && false)) {
                                                                        var v44 float64 = 1.5
                                                                        v44 += (v44 + v44)
                                                                        v42 = (v42 % ((-260) | 1))
                                                                        _ = v44
                                                                    }
                                                                    v23 = v23
                                                                }
                                                                _ = v42
                                                            }
                                                            v21 = (^(v21))
                                                            v25 = (v25 / ((v25) | 1))
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            delete(v3, "foo bar")
                                                        } else {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            if (("a" <= "") && (float32(0.1) > float32(3.14159))) {
                                                                var v45 uint32 = v41
                                                                fmt.Println(v17)
                                                                _ = v45
                                                            }
                                                        }
                                                        _ = v41
                                                    }
                                                case -2147483648:
                                                    if (!false) {
                                                        if (113 <= v23) {
                                                            if true {
                                                                if true {
                                                                    v25 = (v25 - (v25 >> ((18446744073709551615) & 63)))
                                                                    if false {
                                                                        switch ("foo bar" + "\n\t") {
                                                                        case "foo bar":
                                                                            switch (v4 * v4) {
                                                                            case -2147483648:
                                                                                if false {
                                                                                    var v46 int32 = v25
                                                                                    var v47 float64 = 0.1
                                                                                    v21 -= v21
                                                                                    _ = v46
                                                                                    _ = v47
                                                                                } else {
                                                                                    fmt.Println(v21)
                                                                                }
                                                                            case 0:
                                                                                fmt.Println(v26, v25)
                                                                            }
                                                                        default:
                                                                            fmt.Println(v25, v38)
                                                                            v2 = append(v2, 419)
                                                                            v4 -= -502
                                                                        }
                                                                    } else {
                                                                        var v48 float32 = (((float32(-2.25) / float32(0.1)) * 1.5) / (3.14159 + (0.0 / 1.5)))
                                                                        if (false && true) {
                                                                            var v49 int64 = -9223372036854775808
                                                                            _ = v49
                                                                        } else {
                                                                            v50 := GenS2{F0: (v21 / (((v21 << ((v27) & 63))) | 1))}
                                                                            fmt.Println(v48, v21)
                                                                            fmt.Println(v27, v26, fi22)
                                                                            if (((v48 * (float32(1.5) + v48)) > (0.0 + v48)) || false) {
                                                                                var v51 int32 = (-(v25))
                                                                                fmt.Println(v51, v26)
                                                                                _ = v51
                                                                            } else {
                                                                                v25 -= (v25 - -257)
                                                                                if ((v21 + 195) == v21) {
                                                                                    if true {
                                                                                        if false {
                                                                                            var v52 string = ("" + ("" + (("hello" + "foo bar") + ("a" + "\n\t"))))
                                                                                            fmt.Println(fi22, v52)
                                                                                            var v53 uint32 = 65535
                                                                                            _ = v52
                                                                                            _ = v53
                                                                                        } else {
                                                                                            var v54 int8 = v23
                                                                                            if false {
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                if false {
                                                                                                    var v55 int8 = (v54 - 1)
                                                                                                    _ = v55
                                                                                                } else {
                                                                                                    v38 -= (v17 | 445)
                                                                                                    fmt.Println(v4, v54, fi22)
                                                                                                    switch "" {
                                                                                                    case "":
                                                                                                        fmt.Println(v17)
                                                                                                    case "foo bar":
                                                                                                        var v56 uint = (v27 + (^(v27)))
                                                                                                        _ = v56
                                                                                                    }
                                                                                                }
                                                                                                var v57 int16 = v26
                                                                                                _ = v57
                                                                                            } else {
                                                                                                v58 := GenS2{F0: 1}
                                                                                                _ = v58
                                                                                            }
                                                                                            _ = v54
                                                                                        }
                                                                                        fmt.Println(v21, fi20, v25)
                                                                                        v59 := GenS3{F0: v25}
                                                                                        _ = v59
                                                                                    }
                                                                                    if true {
                                                                                        v4 += (fi22 + int(v21))
                                                                                        var v60 string = "\n\t"
                                                                                        _ = v60
                                                                                    } else {
                                                                                        v38 -= v17
                                                                                        fmt.Println(v23)
                                                                                        var v61 uint8 = 255
                                                                                        _ = v61
                                                                                    }
                                                                                    fmt.Println(v21, v27, v17)
                                                                                }
                                                                                var v62 int32 = (v25 | (v25 & v25))
                                                                                v63 := GenS5{F0: v21}
                                                                                _ = v62
                                                                                _ = v63
                                                                            }
                                                                            _ = v50
                                                                        }
                                                                        _ = v48
                                                                    }
                                                                    v64 := GenS3{F0: v25}
                                                                    _ = v64
                                                                } else {
                                                                    var v65 string = "\n\t"
                                                                    v66 := GenS6{F0: 255}
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    _ = v65
                                                                    _ = v66
                                                                }
                                                                if false {
                                                                    fmt.Println(v23)
                                                                    if false {
                                                                        switch (v4 & -933) {
                                                                        case 128:
                                                                            var v67 int = v4
                                                                            _ = v67
                                                                        case -845:
                                                                            if false {
                                                                                v68 := GenS2{F0: (^(v21))}
                                                                                v2 = append(v2, 4294967295)
                                                                                if (v17 <= (v38 << (((v27 | 255)) & 63))) {
                                                                                    v25 -= (v25 + 755)
                                                                                    v27 = 0
                                                                                    var v69 uint8 = 42
                                                                                    _ = v69
                                                                                }
                                                                                _ = v68
                                                                            } else {
                                                                                if true {
                                                                                    var v70 bool = (211 >= 65535)
                                                                                    _ = len(v1)
                                                                                    _ = v70
                                                                                }
                                                                                v71 := GenS7{F0: 2147483647, F1: 217, F2: true}
                                                                                _ = v71
                                                                            }
                                                                        default:
                                                                            v17 -= v38
                                                                        }
                                                                    }
                                                                    if true {
                                                                        var v72 uint64 = 1
                                                                        v38 -= (v38 / (((v38 + (v38 >> (((^(v27))) & 63)))) | 1))
                                                                        v27 = (v27 + (v27 << ((42) & 63)))
                                                                        _ = v72
                                                                    } else {
                                                                        fmt.Println(v26, v17, fi22)
                                                                    }
                                                                    var v73 int16 = int16(v27)
                                                                    if true {
                                                                        fmt.Println(fi22, v27)
                                                                    }
                                                                    _ = v73
                                                                }
                                                                v74 := GenS8{F0: v38, F1: 202, F2: "hello"}
                                                                _ = v74
                                                            } else {
                                                                v25 += 1
                                                            }
                                                        }
                                                        switch "\n\t" {
                                                        case "hello":
                                                            var v75 int16 = v26
                                                            _ = v75
                                                        }
                                                        var v76 float32 = 1e10
                                                        _ = v76
                                                    } else {
                                                        var v77 bool = true
                                                        v78 := GenS9{F0: 255, F1: v21, F2: ("hello" + "foo bar")}
                                                        switch (^(fi20)) {
                                                        case -88:
                                                            v21 = (v21 << (((^(v27))) & 63))
                                                        case 87:
                                                            v23 -= 100
                                                            switch fi20 {
                                                            case -1:
                                                                var v79 int16 = -760
                                                                _ = v79
                                                            default:
                                                                if (!v77) {
                                                                    if v77 {
                                                                        v80 := GenS10{F0: 1, F1: v27}
                                                                        _ = v80
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        v26 -= (v26 >> (((^(v27))) & 63))
                                                                        v27 -= (v27 << ((v27) & 63))
                                                                        v26 = (v26 | v26)
                                                                        v77 = v77
                                                                    case "a":
                                                                        var v81 int32 = (v25 % ((805) | 1))
                                                                        _ = v81
                                                                    case "hello":
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    default:
                                                                        v4 += fi20
                                                                    }
                                                                    v38 += v38
                                                                } else {
                                                                    fmt.Println(v4, fi20)
                                                                    v82 := GenS1{F0: (-(v38)), F1: 4294967295}
                                                                    _ = v82
                                                                }
                                                            }
                                                        case 32767:
                                                            if true {
                                                                v83 := GenS7{F0: v25, F1: 53, F2: ((v17 < (v17 % (((^(v38))) | 1))) || v77)}
                                                                v38 = (v17 & v17)
                                                                v25 = v25
                                                                _ = v83
                                                            } else {
                                                                delete(v3, "hello")
                                                                switch "" {
                                                                case "a":
                                                                    v84 := GenS11{F0: 0, F1: 9223372036854775807, F2: (v17 << ((v27) & 63))}
                                                                    _ = v84
                                                                case "":
                                                                    v85 := GenS12{F0: 32767, F1: v27}
                                                                    _ = v85
                                                                case "\n\t":
                                                                    if (v77 && ((1.5 / (float64(1e10) * 1e10)) == 1.5)) {
                                                                        fmt.Println(v25, v77)
                                                                        v86 := GenS10{F0: 32767, F1: 18446744073709551615}
                                                                        _ = v86
                                                                    } else {
                                                                        switch (fi20 % ((v4) | 1)) {
                                                                        case -905:
                                                                            if false {
                                                                                delete(v3, "foo bar")
                                                                                switch (fi20 * (fi22 | (fi22 >> ((v27) & 63)))) {
                                                                                case 2147483647:
                                                                                    var v87 int16 = v26
                                                                                    _ = v87
                                                                                case 128:
                                                                                    var v88 uint = 0
                                                                                    _ = v88
                                                                                case 559:
                                                                                    v17 -= int64(fi20)
                                                                                default:
                                                                                    v23 -= 67
                                                                                }
                                                                            }
                                                                        case -181:
                                                                            v27 += 255
                                                                        case -9223372036854775808:
                                                                            if (!(v26 > (v26 | 735))) {
                                                                                v3[("hello" + "hello")] = uint32(v21)
                                                                                var v89 int64 = (v17 % (((^(v38))) | 1))
                                                                                if (v77 || true) {
                                                                                    v23 = (v23 / ((v23) | 1))
                                                                                }
                                                                                var v90 float64 = (3.14159 - ((1.5 - 1e10) + 1.5))
                                                                                _ = v89
                                                                                _ = v90
                                                                            } else {
                                                                                var v91 string = "foo bar"
                                                                                _ = v91
                                                                            }
                                                                        default:
                                                                            v21 += (v21 + v21)
                                                                        }
                                                                        var v92 float32 = (3.14159 / 1.0)
                                                                        v38 = (v38 / (((v17 << ((v27) & 63))) | 1))
                                                                        _ = v92
                                                                    }
                                                                default:
                                                                    var v93 uint32 = 2
                                                                    _ = v93
                                                                }
                                                                if (v77 && (v77 || v77)) {
                                                                    var v94 int16 = (-(v26))
                                                                    _ = v94
                                                                }
                                                            }
                                                        }
                                                        var v95 uint32 = 870
                                                        _ = v77
                                                        _ = v78
                                                        _ = v95
                                                    }
                                                case -431:
                                                    var v96 uint = v27
                                                    _ = v96
                                                }
                                                _ = v29
                                                _ = v37
                                                _ = v38
                                            } else {
                                                if true {
                                                    fmt.Println(v27, v27)
                                                    var v97 uint64 = v21
                                                    _ = v97
                                                } else {
                                                    fmt.Println(v25, v4)
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                var v98 uint16 = 525
                                                _ = v98
                                            }
                                        default:
                                            var v99 string = (("foo bar" + "") + "hello")
                                            _ = v99
                                        }
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = v25
                                }
                                var v100 bool = false
                                var v101 int = v4
                                _ = v100
                                _ = v101
                            }
                            v102 := GenS11{F0: 42, F1: (v4 & 0), F2: v17}
                            _ = v102
                            _ = v23
                        } else {
                            var v103 int16 = -724
                            _ = v103
                        }
                        v104 := GenS13{F0: 1, F1: 2147483647}
                        _ = v104
                    } else {
                        fmt.Println(v21, fi20, v17)
                        var v105 int16 = 266
                        _ = v105
                    }
                    _ = len(v3)
                    v4 -= v4
                    v17 -= v17
                }
                var v106 int16 = -1
                v107 := GenS10{F0: 32767, F1: 18446744073709551615}
                _ = fi22
                _ = v106
                _ = v107
            }
            _ = len(v3)
            _ = v21
        } else {
            var v108 uint16 = 997
            if true {
                v17 = -1
            } else {
                switch "\n\t" {
                case "hello":
                    for fi109 := 0; fi109 < 19; fi109++ {
                        v108 = 32767
                        var v110 int8 = 100
                        v4 -= (fi109 >> ((874) & 63))
                        _ = fi109
                        _ = v110
                    }
                default:
                    v108 = 65535
                }
                var v111 float32 = (-2.25 * 0.0)
                _ = v111
            }
            _ = v108
        }
        var v112 int16 = 42
        _ = fi20
        _ = v112
    }
    fmt.Println(v17, v4)
    for fi113 := 0; fi113 < 7; fi113++ {
        v17 = (v17 | v17)
        v1[2147483648] = (false || true)
        _ = fi113
    }
    switch v4 {
    case 128:
        for fi114 := 0; fi114 < 2; fi114++ {
            v115 := GenS1{F0: (-(v17)), F1: 42}
            _ = fi114
            _ = v115
        }
    case 59:
        switch "a" {
        case "foo bar":
            v4 -= v4
        }
    default:
        var v116 int = (v4 ^ v4)
        _ = v116
    }
    var v117 float64 = ((0.1 + ((0.0 / 1e10) - ((3.14159 + 0.0) / 0.1))) - float64((0.1 + 1e10)))
    if true {
        v118 := GenS14{F0: 0, F1: (^(v17)), F2: -178}
        v4 = v4
        fmt.Println(v17, v117)
        _ = v118
    }
    for fi119 := 0; fi119 < 15; fi119++ {
        fmt.Println(fi119, fi119, v4)
        var v120 uint8 = 2
        _ = fi119
        _ = v120
    }
    var v121 int = (v4 % ((1) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v117:", v117)
    fmt.Println("v121:", v121)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v117
    _ = v121
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
