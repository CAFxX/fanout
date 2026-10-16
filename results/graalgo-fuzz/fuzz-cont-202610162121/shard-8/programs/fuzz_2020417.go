package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 bool }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 bool; F2 uint }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 int; F1 int32; F2 int64 }
type GenS7 struct { F0 uint; F1 int; F2 uint8 }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 bool }

func main() {
    var v1 int32 = 942
    var v2 int = 32768
    v3 := []string{"hello"}
    var v4 uint32 = 18
    v5 := []uint32{(^(v4)), v4}
    v6 := make(map[string]uint32)
    if true {
        if false {
            var v7 bool = (!(false && false))
            for fi8 := 0; fi8 < 7; fi8++ {
                if (0 <= 0) {
                    var v9 uint = 255
                    for fi10 := 0; fi10 < 4; fi10++ {
                        var v11 int8 = -110
                        _ = fi10
                        _ = v11
                    }
                    _ = v9
                } else {
                    v12 := v6["\n\t"]
                    _ = v12
                }
                var v13 float64 = 0.1
                _ = fi8
                _ = v13
            }
            _ = v7
        } else {
            v14 := GenS1{F0: 893, F1: ("foo bar" + ("a" + ((((("a" + ("\n\t" + "")) + ("hello" + ("a" + "\n\t"))) + "") + "\n\t") + ("" + "hello")))), F2: (true && (v2 >= (v2 | v2)))}
            switch -368 {
            case -9223372036854775808:
                var v15 uint16 = 2
                _ = v15
            case 9223372036854775807:
                fmt.Println(v4, v4)
            case 0:
                for fi16 := 0; fi16 < 11; fi16++ {
                    fmt.Println(v4, fi16)
                    if false {
                        var v17 int16 = -695
                        v1 -= (v1 | (-(v1)))
                        if (!(4294967295 >= 2)) {
                            fmt.Println(v2)
                            v18 := GenS2{F0: (^(v2)), F1: ("\n\t" + "foo bar")}
                            _ = v18
                        }
                        _ = v17
                    }
                    v3 = append(v3, "foo bar")
                    _ = fi16
                }
            default:
                var v19 int64 = -751
                _ = v19
            }
            _ = v14
        }
        if (905 != -167) {
            delete(v6, "\n\t")
        }
        var v20 int32 = (v1 / ((v1) | 1))
        _ = v20
    } else {
        var v21 float64 = (float64(1.5) / ((0.1 / 0.1) * 1.5))
        if (v21 == 0.1) {
            if true {
                for fi22 := 0; fi22 < 18; fi22++ {
                    for fi23 := 0; fi23 < 1; fi23++ {
                        v2 -= 87
                        _ = fi23
                    }
                    _ = fi22
                }
                switch 127 {
                case 42:
                    var v24 uint64 = 42
                    _ = v24
                case 9223372036854775807:
                    for fi25 := 0; fi25 < 9; fi25++ {
                        var v26 float64 = float64(0.0)
                        v27 := GenS1{F0: (v4 | v4), F1: "foo bar", F2: true}
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
                switch (v2 % ((v2) | 1)) {
                case 758:
                    var v28 string = ("foo bar" + "a")
                    _ = v28
                default:
                    fmt.Println(v2)
                }
            } else {
                fmt.Println(v2)
                v2 = v2
            }
            for fi29 := 0; fi29 < 9; fi29++ {
                if true {
                    v1 = (v1 ^ (v1 ^ (v1 >> ((18446744073709551615) & 63))))
                    v4 = v4
                    switch "a" {
                    case "hello":
                        fmt.Println(fi29, v1)
                    case "foo bar":
                        if false {
                            if (272 == 1) {
                                _ = len(v6)
                            }
                        } else {
                            fmt.Println(v4, v4)
                            for fi30 := 0; fi30 < 1; fi30++ {
                                var v31 string = ""
                                switch (-(fi30)) {
                                case 32767:
                                    fmt.Println(v31)
                                    var v32 string = v31
                                    var v33 uint = uint(v4)
                                    _ = v32
                                    _ = v33
                                case 42:
                                    var v34 int = fi29
                                    _ = v34
                                case -9223372036854775808:
                                    v35 := GenS3{F0: 0}
                                    _ = v35
                                }
                                var v36 int16 = 32767
                                _ = fi30
                                _ = v31
                                _ = v36
                            }
                        }
                    case "":
                        for fi37 := 0; fi37 < 15; fi37++ {
                            switch "hello" {
                            case "hello":
                                v6["foo bar"] = v4
                            case "foo bar":
                                v38 := GenS3{F0: 378}
                                _ = v38
                            case "a":
                                switch (v2 - fi29) {
                                case 32767:
                                    v2 += fi29
                                }
                            default:
                                if false {
                                    switch ("\n\t" + "a") {
                                    case "\n\t":
                                        var v39 int32 = v1
                                        _ = v39
                                    case "":
                                        fmt.Println(fi29, v2, fi29)
                                    case "a":
                                        if false {
                                            var v40 int16 = 516
                                            _ = v40
                                        } else {
                                            v6["hello"] = (v4 << ((0) & 63))
                                            v4 = v4
                                            switch (fi37 / ((v2) | 1)) {
                                            case 32767:
                                                _ = len(v5)
                                            case -9223372036854775808:
                                                v41 := GenS3{F0: 4294967295}
                                                _ = v41
                                            case 566:
                                                if (!true) {
                                                    switch -648 {
                                                    case 2147483647:
                                                        v1 = v1
                                                    case -9223372036854775808:
                                                        switch (v2 & fi29) {
                                                        case -703:
                                                            v42 := GenS1{F0: (v4 & (v4 % ((1) | 1))), F1: ("\n\t" + ((("" + "foo bar") + "") + ("foo bar" + "\n\t"))), F2: false}
                                                            _ = v42
                                                        case -402:
                                                            v43 := v6[((("\n\t" + "\n\t") + "hello") + ("\n\t" + "hello"))]
                                                            _ = v43
                                                        case -9223372036854775808:
                                                            if (fi29 != (-(fi37))) {
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    fmt.Println(fi29)
                                                                case "foo bar":
                                                                    if false {
                                                                        var v44 uint32 = (v4 & (v4 % (((v4 / ((1) | 1))) | 1)))
                                                                        v3 = append(v3, "")
                                                                        var v45 uint8 = 42
                                                                        _ = v44
                                                                        _ = v45
                                                                    }
                                                                case "":
                                                                    v21 = ((1e10 / 3.14159) * (1.5 - float64(1e10)))
                                                                }
                                                                var v46 uint64 = 18446744073709551615
                                                                _ = v46
                                                            } else {
                                                                var v47 int8 = 42
                                                                switch fi37 {
                                                                case -9223372036854775808:
                                                                    var v48 bool = false
                                                                    _ = v48
                                                                case 127:
                                                                    var v49 uint64 = 438
                                                                    _ = v49
                                                                }
                                                                _ = v47
                                                            }
                                                        default:
                                                            var v50 uint16 = 32767
                                                            _ = v50
                                                        }
                                                    case 322:
                                                        if (true || (false || true)) {
                                                            v51 := GenS1{F0: v4, F1: ("hello" + ((("hello" + "foo bar") + "\n\t") + "a")), F2: false}
                                                            v1 = 0
                                                            _ = v51
                                                        }
                                                    default:
                                                        fmt.Println(v1)
                                                    }
                                                    v5 = append(v5, 377)
                                                    var v52 uint64 = 553
                                                    _ = v52
                                                }
                                            default:
                                                fmt.Println(v21, v4, fi29)
                                            }
                                        }
                                    }
                                    if (!(!(1 == v1))) {
                                        v53 := GenS2{F0: fi29, F1: ""}
                                        if false {
                                            _ = len(v6)
                                            var v54 int32 = v1
                                            _ = v54
                                        } else {
                                            if (float32(-2.25) >= (float32(v21) - (-2.25 / 1.0))) {
                                                v55 := GenS4{F0: -1, F1: (77 > 289), F2: 508}
                                                fmt.Println(fi37)
                                                if (true && false) {
                                                    switch (fi37 | fi37) {
                                                    case 4294967295:
                                                        if false {
                                                            switch fi37 {
                                                            case 9223372036854775807:
                                                                if ((v21 * (float64(3.14159) - float64(float32(v21)))) <= v21) {
                                                                    var v56 uint64 = 4294967295
                                                                    _ = v56
                                                                } else {
                                                                    v57 := v6[((("a" + "hello") + "") + "")]
                                                                    var v58 bool = (true && false)
                                                                    v2 += (fi37 & v2)
                                                                    _ = v57
                                                                    _ = v58
                                                                }
                                                            case 1:
                                                                var v59 int16 = 42
                                                                _ = v59
                                                            }
                                                            fmt.Println(fi37, v2)
                                                            v4 += (v4 << ((0) & 63))
                                                        }
                                                    case 128:
                                                        switch 127 {
                                                        case -9223372036854775808:
                                                            if ((!((true || false) || false)) && false) {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                                fmt.Println(v2)
                                                                fmt.Println(v2, v1, v2)
                                                            }
                                                        case 32768:
                                                            v60 := GenS3{F0: 18446744073709551615}
                                                            _ = v60
                                                        }
                                                    case 32768:
                                                        var v61 float64 = v21
                                                        _ = v61
                                                    default:
                                                        switch "hello" {
                                                        case "":
                                                            var v62 uint16 = 32767
                                                            _ = v62
                                                        case "a":
                                                            v5 = append(v5, (v4 * (^(v4))))
                                                            v63 := v6[((((("hello" + ("a" + (("foo bar" + "foo bar") + ("hello" + "\n\t")))) + ("hello" + "\n\t")) + "") + "foo bar") + "\n\t")]
                                                            if false {
                                                                var v64 int16 = 0
                                                                _ = v64
                                                            }
                                                            v65 := GenS4{F0: v1, F1: (((((float64(((float32(3.14159) + ((-2.25 / 1e10) + (3.14159 - -2.25))) * -2.25)) + (((v21 + ((float64(0.0) * float64(1.5)) - (v21 + (v21 * (v21 - 1.5))))) * float64(float32(v21))) * 1.5)) * (3.14159 / v21)) / (((v21 / float64(-2.25)) / v21) * v21)) - -2.25) <= v21), F2: 255}
                                                            _ = v63
                                                            _ = v65
                                                        case "\n\t":
                                                            if (-9223372036854775808 < 0) {
                                                                v4 -= v4
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            } else {
                                                                delete(v6, "hello")
                                                            }
                                                        }
                                                    }
                                                    var v66 uint16 = 669
                                                    if true {
                                                        var v67 bool = false
                                                        fmt.Println(v1)
                                                        _ = v67
                                                    } else {
                                                        v66 = v66
                                                        switch (fi29 >> ((1) & 63)) {
                                                        case 32768:
                                                            if (!true) {
                                                                v21 = (v21 - v21)
                                                            }
                                                        case 50:
                                                            switch (v2 ^ 0) {
                                                            case 128:
                                                                var v68 int = -2147483648
                                                                _ = v68
                                                            case -1:
                                                                _ = len(v6)
                                                            case -699:
                                                                if (9223372036854775807 != 9223372036854775807) {
                                                                    fmt.Println(fi29, v66, fi29)
                                                                    if (-128 == -128) {
                                                                        var v69 int8 = -1
                                                                        var v70 int64 = -362
                                                                        if ((!(v66 > v66)) && true) {
                                                                            fmt.Println(v2, v1)
                                                                        } else {
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                v71 := GenS4{F0: (v1 - v1), F1: (true || false), F2: uint(v4)}
                                                                                _ = v71
                                                                            case "a":
                                                                                fmt.Println(fi37, v2, v1)
                                                                            }
                                                                        }
                                                                        _ = v69
                                                                        _ = v70
                                                                    } else {
                                                                        v4 += (v4 >> ((146) & 63))
                                                                    }
                                                                    var v72 bool = (true && true)
                                                                    _ = v72
                                                                }
                                                            }
                                                        case 712:
                                                            fmt.Println(v21, v2, fi37)
                                                        }
                                                        v1 = v1
                                                    }
                                                    _ = v66
                                                }
                                                _ = v55
                                            }
                                            _ = len(v3)
                                        }
                                        v5 = append(v5, 42)
                                        v1 = -1
                                        _ = v53
                                    }
                                }
                            }
                            _ = fi37
                        }
                    default:
                        v4 -= 0
                    }
                }
                switch v2 {
                case -9223372036854775808:
                    var v73 float32 = (1e10 + -2.25)
                    _ = v73
                case 42:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                var v74 uint64 = 885
                _ = fi29
                _ = v74
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            var v75 string = ("a" + "foo bar")
            switch (-(v2)) {
            case -3:
                v21 -= float64(1.5)
            default:
                v4 -= v4
            }
            for fi76 := 0; fi76 < 12; fi76++ {
                fmt.Println(v21)
                switch -9223372036854775808 {
                case 1:
                    delete(v6, "hello")
                case 2147483647:
                    var v77 bool = false
                    _ = v77
                case 242:
                    for fi78 := 0; fi78 < 17; fi78++ {
                        if ((v2 >> ((42) & 63)) == -1) {
                            switch ("hello" + v75) {
                            case "foo bar":
                                var v79 float64 = 3.14159
                                _ = v79
                            case "\n\t":
                                var v80 string = (("a" + ("" + "\n\t")) + v75)
                                _ = v80
                            }
                            v81 := GenS5{F0: 602, F1: (v1 / ((v1) | 1))}
                            if true {
                                switch v75 {
                                case "foo bar":
                                    var v82 uint64 = 0
                                    _ = v82
                                case "":
                                    var v83 int8 = -1
                                    _ = v83
                                }
                                switch 9223372036854775807 {
                                case 654:
                                    fmt.Println(fi76, fi76)
                                case 2147483647:
                                    fmt.Println(fi76)
                                }
                                if ((v4 % ((v4) | 1)) <= 42) {
                                    switch v2 {
                                    case 89:
                                        v21 -= (v21 / float64(1e10))
                                    case -1:
                                        fmt.Println(fi76, v1)
                                    case -145:
                                        if false {
                                            fmt.Println(fi78, v75, v4)
                                        } else {
                                            var v84 int8 = 68
                                            v85 := v6[v75]
                                            _ = v84
                                            _ = v85
                                        }
                                    default:
                                        delete(v6, (v75 + v75))
                                    }
                                    switch ("hello" + ("foo bar" + v75)) {
                                    case "":
                                        v1 = (v1 % ((-832) | 1))
                                    case "\n\t":
                                        v1 = v1
                                    default:
                                        v2 = (^(v2))
                                    }
                                    var v86 bool = true
                                    v1 = (-(v1))
                                    _ = v86
                                }
                                var v87 float64 = (v21 / v21)
                                v87 = v21
                                var v88 int = (fi78 ^ fi78)
                                _ = v87
                                _ = v88
                            }
                            v89 := GenS4{F0: v1, F1: false, F2: 992}
                            if false {
                                delete(v6, v75)
                            } else {
                                var v90 uint16 = 1
                                v21 = 3.14159
                                _ = v90
                            }
                            _ = v81
                            _ = v89
                        } else {
                            var v91 float32 = float32(0.0)
                            if false {
                                var v92 bool = (!false)
                                v93 := GenS6{F0: fi78, F1: v1, F2: -1}
                                _ = v92
                                _ = v93
                            } else {
                                fmt.Println(v21, v75)
                                if false {
                                    fmt.Println(v91, fi78, fi76)
                                    v94 := GenS1{F0: v4, F1: (v75 + v75), F2: (true && (3.14159 >= v91))}
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v94
                                } else {
                                    var v95 uint16 = 1
                                    _ = v95
                                }
                            }
                            _ = v91
                        }
                        v2 -= (^(fi78))
                        if (0.0 <= ((1.5 - float32(((v21 / 0.1) + float64(1.5)))) * (((-2.25 + 1e10) * 3.14159) - float32((v21 / (v21 * 0.1)))))) {
                            var v96 int = fi76
                            var v97 int32 = 0
                            _ = v96
                            _ = v97
                        } else {
                            fmt.Println(v1, fi78)
                            v98 := v6[""]
                            var v99 int = (-(fi78))
                            if len(v3) > 0 {
                                v3[0] = v75
                            }
                            var v100 uint32 = (^(v4))
                            _ = v100
                            _ = v98
                            _ = v99
                        }
                        _ = fi78
                    }
                }
                _ = fi76
            }
            _ = v75
        }
        var v101 bool = true
        _ = v101
        _ = v21
    }
    if true {
        for fi102 := 0; fi102 < 2; fi102++ {
            fmt.Println(fi102, v2, fi102)
            var v103 bool = true
            _ = fi102
            _ = v103
        }
        v104 := GenS7{F0: 0, F1: v2, F2: 1}
        if true {
            var v105 int8 = 1
            _ = v105
        } else {
            var v106 float32 = ((float32(float64(1.5)) + 3.14159) / 1e10)
            _ = v106
        }
        _ = v104
    }
    var v107 int = v2
    v108 := GenS2{F0: v2, F1: "foo bar"}
    var v109 uint = 2
    _ = len(v6)
    v4 = v4
    if true {
        for fi110 := 0; fi110 < 18; fi110++ {
            fmt.Println(v107)
            if true {
                if ((3.14159 + 1e10) == ((float32(0.1) - 0.0) - (0.0 * (3.14159 + 1.5)))) {
                    var v111 int8 = 1
                    v112 := GenS8{F0: (int32(v111) < (v1 * v1)), F1: 65535}
                    _ = v111
                    _ = v112
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = ""
                }
                v113 := GenS9{F0: false}
                _ = v113
            }
            _ = fi110
        }
        v4 += (v4 & 0)
    } else {
        var v114 float32 = 0.0
        _ = v114
    }
    _ = len(v6)
    v107 = v2
    for fi115 := 0; fi115 < 8; fi115++ {
        var v116 float32 = 0.0
        v117 := GenS8{F0: false, F1: 2}
        _ = fi115
        _ = v116
        _ = v117
    }
    fmt.Println(v109)
    var v118 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v109:", v109)
    fmt.Println("v118:", v118)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v107
    _ = v108
    _ = v109
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
