package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint64; F2 bool }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint, p1 int16, p2 int8) int8 {
    p2 -= (p2 >> ((42) & 63))
    var v6 int8 = (p2 % ((100) | 1))
    fmt.Println(p2, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (v6 % ((0) | 1))
}

func main() {
    v1 := []int{127}
    var v2 int32 = -12
    var v3 uint64 = 838
    v4 := make(map[string]uint32)
    var v5 uint64 = v3
    v3 = v5
    v3 = 2
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: (fi7 & 811), F1: v2, F2: 0}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if true {
        fmt.Println(v3, v3)
        v10 := GenS2{F0: 424, F1: 719}
        _ = v10
    } else {
        v3 = 42
        if ((127 < 0) && false) {
            v2 -= -2147483648
            if len(v1) > 0 {
                v1[0] = -357
            }
            var v11 bool = (65535 == 2147483647)
            _ = v11
        }
    }
    if true {
        for fi12 := 0; fi12 < 17; fi12++ {
            switch (^(fi12)) {
            case -26:
                if true {
                    if len(v1) > 0 {
                        v1[0] = 32768
                    }
                    fmt.Println(v5, v3)
                    fmt.Println(fi12, v3, fi12)
                } else {
                    v2 += v2
                }
            case 1:
                fmt.Println(v2)
            case 32768:
                var v13 uint8 = uint8(v5)
                _ = v13
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi12
        }
        v14 := GenS1{F0: 9223372036854775807, F1: (v2 * v2), F2: -452}
        v5 = (^(v3))
        _ = v14
    } else {
        switch "a" {
        case "a":
            if true {
                switch 128 {
                case 0:
                    var v15 uint16 = 32767
                    _ = v15
                }
                if (((v3 < v5) || ((v3 >> ((4294967295) & 63)) > (v3 % ((18446744073709551615) | 1)))) || ((255 <= 42) && true)) {
                    var v16 uint64 = (v3 ^ (^(v3)))
                    switch "foo bar" {
                    case "foo bar":
                        var v17 float64 = 0.0
                        _ = v17
                    case "a":
                        var v18 int64 = -14
                        _ = v18
                    case "":
                        fmt.Println(v2, v2, v5)
                    }
                    v19 := GenS1{F0: -198, F1: 920, F2: 42}
                    _ = v16
                    _ = v19
                }
                v20 := v4["hello"]
                _ = v20
            } else {
                if (!true) {
                    for fi21 := 0; fi21 < 19; fi21++ {
                        fmt.Println(fi21)
                        var v22 uint8 = 127
                        var v23 int16 = 0
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    v5 -= v5
                    switch 2147483647 {
                    case 32767:
                        if (-2.25 <= (-2.25 * -2.25)) {
                            switch -9223372036854775808 {
                            case -1:
                                for fi24 := 0; fi24 < 4; fi24++ {
                                    if false {
                                        var v25 uint8 = 127
                                        _ = v25
                                    } else {
                                        v26 := v4["\n\t"]
                                        _ = v26
                                    }
                                    if ((fi24 & fi24) >= fi24) {
                                        for fi27 := 0; fi27 < 14; fi27++ {
                                            var v28 int16 = 32767
                                            var v29 int = (fi27 % ((int(v28)) | 1))
                                            switch (v29 % ((fi27) | 1)) {
                                            case 2147483647:
                                                v5 = (v3 % ((224) | 1))
                                            case 32768:
                                                if (!((fi24 & v29) > fi27)) {
                                                    var v30 bool = (0.0 <= 1e10)
                                                    if (0 != -119) {
                                                        var v31 int = fi27
                                                        delete(v4, "")
                                                        if false {
                                                            var v32 int16 = 32767
                                                            switch ("foo bar" + "\n\t") {
                                                            case "foo bar":
                                                                fmt.Println(v2)
                                                            case "a":
                                                                v31 += (fi24 / ((fi24) | 1))
                                                            }
                                                            if false {
                                                                v33 := GenS1{F0: fi27, F1: (v2 | v2), F2: v28}
                                                                v2 -= v2
                                                                _ = v33
                                                            } else {
                                                                switch ("hello" + "a") {
                                                                case "a":
                                                                    v34 := v4["a"]
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        v35 := GenS2{F0: (v34 % ((v34) | 1)), F1: v34}
                                                                        _ = v35
                                                                    }
                                                                    var v36 string = "foo bar"
                                                                    _ = v34
                                                                    _ = v36
                                                                case "":
                                                                    if (-9223372036854775808 == -1) {
                                                                        if (!true) {
                                                                            if (v30 || true) {
                                                                                _ = len(v4)
                                                                                var v37 uint64 = (^(v3))
                                                                                if (v31 == v31) {
                                                                                    switch fi24 {
                                                                                    case -337:
                                                                                        v38 := GenS2{F0: 1, F1: uint32(v5)}
                                                                                        _ = v38
                                                                                    }
                                                                                    v28 = (^(v32))
                                                                                    fmt.Println(v32, v37, v30)
                                                                                    v5 = v3
                                                                                }
                                                                                _ = v37
                                                                            }
                                                                            switch fi27 {
                                                                            case 128:
                                                                                var v39 string = (((("a" + "") + "\n\t") + ("" + (((("\n\t" + "a") + ("hello" + "foo bar")) + ("hello" + "foo bar")) + "a"))) + "a")
                                                                                _ = v39
                                                                            case 2147483647:
                                                                                var v40 int16 = v28
                                                                                _ = v40
                                                                            }
                                                                            if v30 {
                                                                                v41 := GenS2{F0: 1, F1: 4294967295}
                                                                                _ = v41
                                                                            } else {
                                                                                var v42 int8 = 0
                                                                                switch int(v32) {
                                                                                case 42:
                                                                                    v30 = false
                                                                                case -2147483648:
                                                                                    _ = len(v1)
                                                                                case 127:
                                                                                    var v43 bool = v30
                                                                                    _ = v43
                                                                                default:
                                                                                    delete(v4, (("\n\t" + ((("" + "") + "") + "")) + "hello"))
                                                                                }
                                                                                _ = v42
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v5, v2)
                                                                        }
                                                                        var v44 int64 = 0
                                                                        v29 = fi24
                                                                        _ = v44
                                                                    } else {
                                                                        var v45 float64 = 0.0
                                                                        _ = v45
                                                                    }
                                                                case "hello":
                                                                    var v46 float32 = 1e10
                                                                    _ = v46
                                                                default:
                                                                    v47 := GenS2{F0: 2, F1: 3}
                                                                    _ = v47
                                                                }
                                                                if false {
                                                                    v1 = append(v1, fi27)
                                                                    if v30 {
                                                                        v48 := GenS1{F0: v29, F1: v2, F2: (v28 - int16(v5))}
                                                                        v3 = v5
                                                                        _ = v48
                                                                    } else {
                                                                        var v49 float64 = 1.5
                                                                        var v50 string = "hello"
                                                                        _ = v49
                                                                        _ = v50
                                                                    }
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                } else {
                                                                    var v51 float64 = (1.5 / 1.0)
                                                                    var v52 float32 = 3.14159
                                                                    v53 := GenS2{F0: 509, F1: 42}
                                                                    if (-1 > int8(v2)) {
                                                                        v54 := GenS2{F0: 0, F1: 4294967295}
                                                                        if (!false) {
                                                                            switch 32768 {
                                                                            case -579:
                                                                                v4["hello"] = 42
                                                                            case 635:
                                                                                fmt.Println(v51)
                                                                                v55 := GenS3{F0: 65535}
                                                                                _ = v55
                                                                            }
                                                                            v56 := v4[("\n\t" + "hello")]
                                                                            var v57 uint32 = v56
                                                                            _ = v56
                                                                            _ = v57
                                                                        } else {
                                                                            v58 := GenS4{F0: 781, F1: v3, F2: ("" < ("\n\t" + "a"))}
                                                                            v59 := GenS5{F0: ("hello" + ("foo bar" + "foo bar")), F1: 32767}
                                                                            _ = v58
                                                                            _ = v59
                                                                        }
                                                                        v52 -= 0.0
                                                                        _ = v54
                                                                    } else {
                                                                        fmt.Println(v2, fi27)
                                                                        if false {
                                                                            v4[(("foo bar" + "a") + "hello")] = 632
                                                                            var v60 uint = 0
                                                                            fmt.Println(v31, fi27, v29)
                                                                            _ = v60
                                                                        }
                                                                        var v61 int8 = 42
                                                                        _ = v61
                                                                    }
                                                                    if true {
                                                                        if ((((v51 + 3.14159) + 0.0) == v51) && (-1 < 1)) {
                                                                            var v62 uint32 = 207
                                                                            _ = len(v1)
                                                                            v63 := GenS3{F0: 4294967295}
                                                                            _ = v62
                                                                            _ = v63
                                                                        } else {
                                                                            v64 := GenS6{F0: 0, F1: 848}
                                                                            switch v29 {
                                                                            case 292:
                                                                                v28 = v32
                                                                                if false {
                                                                                    v3 -= v5
                                                                                }
                                                                                var v65 int16 = v28
                                                                                _ = v65
                                                                            case 42:
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = fi27
                                                                                }
                                                                            case 4294967295:
                                                                                v66 := GenS5{F0: "a", F1: 0}
                                                                                _ = v66
                                                                            default:
                                                                                switch 4294967295 {
                                                                                case 1:
                                                                                    v28 -= (v28 & -32768)
                                                                                case 32767:
                                                                                    v67 := GenS3{F0: 65535}
                                                                                    _ = v67
                                                                                case -9223372036854775808:
                                                                                    var v68 float32 = ((v52 * -2.25) - v52)
                                                                                    _ = v68
                                                                                }
                                                                            }
                                                                            _ = v64
                                                                        }
                                                                    }
                                                                    var v69 int32 = v2
                                                                    var v70 uint8 = 127
                                                                    delete(v4, "\n\t")
                                                                    if (!(32767 == 65535)) {
                                                                        v70 -= (^(v70))
                                                                        v2 -= (v2 >> ((850) & 63))
                                                                        v71 := GenS5{F0: "", F1: 42}
                                                                        _ = v71
                                                                    } else {
                                                                        switch fi24 {
                                                                        case 127:
                                                                            switch -9223372036854775808 {
                                                                            case 1:
                                                                                if v30 {
                                                                                    v51 -= 1.5
                                                                                    switch 31 {
                                                                                    case -9223372036854775808:
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v28, v32, v30)
                                                                                    }
                                                                                    v72 := GenS7{F0: v70}
                                                                                    _ = v72
                                                                                } else {
                                                                                    v73 := GenS2{F0: 42, F1: 0}
                                                                                    _ = v73
                                                                                }
                                                                            case 27:
                                                                                fmt.Println(v70)
                                                                            case -784:
                                                                                v69 = v69
                                                                            }
                                                                        default:
                                                                            v2 -= 0
                                                                        }
                                                                    }
                                                                    v74 := GenS5{F0: "hello", F1: 42}
                                                                    _ = v51
                                                                    _ = v52
                                                                    _ = v53
                                                                    _ = v69
                                                                    _ = v70
                                                                    _ = v74
                                                                }
                                                                var v75 float64 = float64(-2.25)
                                                                _ = v75
                                                            }
                                                            var v76 float32 = 1e10
                                                            fmt.Println(v76, fi27, v5)
                                                            _ = v32
                                                            _ = v76
                                                        } else {
                                                            switch v31 {
                                                            case 1:
                                                                if v30 {
                                                                    fmt.Println(v29, v28)
                                                                    var v77 bool = v30
                                                                    _ = v77
                                                                }
                                                            case 4294967295:
                                                                switch (("hello" + "hello") + ((("foo bar" + ("\n\t" + ("" + (("\n\t" + ("foo bar" + ("foo bar" + "hello"))) + "hello")))) + "foo bar") + ("\n\t" + "\n\t"))) {
                                                                case "hello":
                                                                    v78 := v4[("hello" + "hello")]
                                                                    _ = v78
                                                                default:
                                                                    v30 = ("foo bar" <= "\n\t")
                                                                }
                                                            case 2147483647:
                                                                var v79 uint64 = v5
                                                                _ = v79
                                                            }
                                                            fmt.Println(fi24)
                                                            v80 := GenS4{F0: 4294967295, F1: (^(v5)), F2: true}
                                                            switch -168 {
                                                            case 918:
                                                                var v81 string = "hello"
                                                                _ = v81
                                                            case 32768:
                                                                v82 := GenS5{F0: ("foo bar" + "hello"), F1: 255}
                                                                _ = v82
                                                            default:
                                                                var v83 uint16 = 42
                                                                _ = v83
                                                            }
                                                            _ = v80
                                                        }
                                                        _ = v31
                                                    }
                                                    switch (v29 ^ 0) {
                                                    case 32768:
                                                        var v84 uint16 = 1
                                                        _ = v84
                                                    default:
                                                        var v85 string = "\n\t"
                                                        _ = v85
                                                    }
                                                    _ = v30
                                                } else {
                                                    fmt.Println(v5, v28, v29)
                                                    var v86 uint8 = 127
                                                    if true {
                                                        v28 += v28
                                                        if true {
                                                            fmt.Println(fi24, v2, v29)
                                                            v28 = (v28 | int16(v86))
                                                            delete(v4, "a")
                                                        }
                                                    }
                                                    _ = v86
                                                }
                                            case 4294967295:
                                                if true {
                                                    fmt.Println(v28, v5, v29)
                                                    if (-2.25 != ((-2.25 * -2.25) + 1.5)) {
                                                        v87 := GenS1{F0: fi24, F1: v2, F2: v28}
                                                        delete(v4, (((((("" + "") + "hello") + "hello") + "\n\t") + "a") + "a"))
                                                        v88 := GenS8{F0: 186}
                                                        _ = v87
                                                        _ = v88
                                                    } else {
                                                        fmt.Println(v28, v5)
                                                        if false {
                                                            if (127 != 1) {
                                                                v29 = v29
                                                                if (((fi27 - v29) <= fi24) && false) {
                                                                    v89 := GenS2{F0: 65535, F1: 182}
                                                                    _ = v89
                                                                } else {
                                                                    var v90 int8 = 1
                                                                    _ = v90
                                                                }
                                                                var v91 string = ("" + "foo bar")
                                                                _ = v91
                                                            }
                                                            var v92 int32 = -1
                                                            var v93 uint32 = 4294967295
                                                            var v94 float64 = 0.0
                                                            _ = v92
                                                            _ = v93
                                                            _ = v94
                                                        }
                                                        var v95 uint16 = 42
                                                        fmt.Println(v3, v3, v29)
                                                        if (((0.0 - 0.1) + (3.14159 * float32((0.1 / 1e10)))) < (0.0 / -2.25)) {
                                                            if (!false) {
                                                                if ((v28 < 1000) || false) {
                                                                    if false {
                                                                        if false {
                                                                            v95 = v95
                                                                            var v96 float32 = (0.0 / 3.14159)
                                                                            _ = v96
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v28, v29)
                                                                    }
                                                                    v1 = append(v1, (-(fi27)))
                                                                    switch (fi24 << ((734) & 63)) {
                                                                    case -423:
                                                                        v97 := GenS9{F0: 42, F1: uint16(v5)}
                                                                        _ = v97
                                                                    }
                                                                } else {
                                                                    _ = len(v1)
                                                                }
                                                                v98 := GenS6{F0: 135, F1: -759}
                                                                switch "foo bar" {
                                                                case "":
                                                                    if (false && (v95 >= 891)) {
                                                                        switch 32767 {
                                                                        case 4294967295:
                                                                            v28 -= v28
                                                                        case 931:
                                                                            v99 := v4[""]
                                                                            _ = v99
                                                                        case 346:
                                                                            var v100 uint = 1
                                                                            _ = v100
                                                                        }
                                                                        fmt.Println(v29)
                                                                    } else {
                                                                        fmt.Println(v5)
                                                                    }
                                                                default:
                                                                    v1 = append(v1, (v29 ^ -835))
                                                                }
                                                                _ = v98
                                                            }
                                                        }
                                                        _ = v95
                                                    }
                                                    v1 = append(v1, (fi24 % (((v29 | fi27)) | 1)))
                                                } else {
                                                    v1 = append(v1, (fi27 >> ((42) & 63)))
                                                    switch "" {
                                                    case "foo bar":
                                                        fmt.Println(v28, fi24)
                                                    case "":
                                                        fmt.Println(v3, fi24, v3)
                                                    }
                                                }
                                            }
                                            _ = fi27
                                            _ = v28
                                            _ = v29
                                        }
                                        v4[""] = 2
                                        v1 = append(v1, (fi24 * (fi24 >> ((2) & 63))))
                                    }
                                    for fi101 := 0; fi101 < 18; fi101++ {
                                        fmt.Println(fi101)
                                        fmt.Println(v2)
                                        _ = fi101
                                    }
                                    _ = fi24
                                }
                            case 0:
                                fmt.Println(v5, v2, v5)
                            }
                            for fi102 := 0; fi102 < 20; fi102++ {
                                v103 := GenS2{F0: 0, F1: 0}
                                fmt.Println(fi102, v5, fi102)
                                _ = fi102
                                _ = v103
                            }
                            v2 += (-(v2))
                        } else {
                            var v104 int = 0
                            switch v104 {
                            case -274:
                                if true {
                                    v104 = v104
                                    for fi105 := 0; fi105 < 4; fi105++ {
                                        var v106 string = "foo bar"
                                        switch int(v2) {
                                        case 549:
                                            var v107 float32 = -2.25
                                            _ = v107
                                        }
                                        _ = fi105
                                        _ = v106
                                    }
                                }
                            default:
                                v104 = v104
                            }
                            v2 = (v2 + v2)
                            _ = v104
                        }
                    case 4294967295:
                        v5 = (^(v5))
                    }
                } else {
                    for fi108 := 0; fi108 < 14; fi108++ {
                        var v109 float64 = (0.1 - 1.5)
                        _ = fi108
                        _ = v109
                    }
                }
                for fi110 := 0; fi110 < 6; fi110++ {
                    var v111 int16 = -177
                    v5 = v5
                    _ = fi110
                    _ = v111
                }
                v3 = 4294967295
            }
        case "foo bar":
            for fi112 := 0; fi112 < 3; fi112++ {
                v3 = v5
                if false {
                    v5 = 4294967295
                    switch "" {
                    case "foo bar":
                        if (255 != 65535) {
                            for fi113 := 0; fi113 < 13; fi113++ {
                                v2 -= (v2 | v2)
                                _ = fi113
                            }
                            if (true || (2 != (v3 * 214))) {
                                var v114 uint = 65535
                                if len(v1) > 0 {
                                    v1[0] = fi112
                                }
                                _ = v114
                            }
                        } else {
                            var v115 string = (("foo bar" + "hello") + ("\n\t" + "\n\t"))
                            v115 = "hello"
                            v116 := GenS7{F0: 42}
                            _ = v115
                            _ = v116
                        }
                    case "hello":
                        var v117 uint64 = (v5 / (((v5 % ((v3) | 1))) | 1))
                        _ = v117
                    default:
                        v2 += 0
                    }
                } else {
                    if true {
                        for fi118 := 0; fi118 < 6; fi118++ {
                            if true {
                                v1 = append(v1, -9223372036854775808)
                                v119 := v4[(("foo bar" + "") + "foo bar")]
                                _ = v119
                            } else {
                                var v120 bool = true
                                v5 -= v3
                                v2 = 442
                                _ = v120
                            }
                            fmt.Println(fi118, fi118, v2)
                            _ = fi118
                        }
                    }
                }
                for fi121 := 0; fi121 < 8; fi121++ {
                    var v122 int32 = v2
                    var v123 string = ("" + "a")
                    _ = fi121
                    _ = v122
                    _ = v123
                }
                _ = fi112
            }
        case "\n\t":
            var v124 uint32 = 1
            _ = v124
        default:
            fmt.Println(v3, v5)
        }
        v125 := GenS1{F0: 4294967295, F1: v2, F2: -1}
        v3 = (v5 % ((2) | 1))
        _ = v125
    }
    v126 := GenS3{F0: 255}
    var v127 int16 = -1
    switch -1 {
    case 1:
        if (255 <= 804) {
            v128 := GenS6{F0: 127, F1: 42}
            switch ("a" + ((("a" + "") + "a") + "a")) {
            case "hello":
                for fi129 := 0; fi129 < 9; fi129++ {
                    v130 := GenS6{F0: 0, F1: -202}
                    fmt.Println(v3)
                    _ = fi129
                    _ = v130
                }
            case "\n\t":
                v127 -= int16(v2)
            case "a":
                if false {
                    v5 -= v3
                } else {
                    v131 := GenS2{F0: 0, F1: 42}
                    v132 := GenS6{F0: 128, F1: 1}
                    v133 := GenS8{F0: 2}
                    _ = v131
                    _ = v132
                    _ = v133
                }
            default:
                v134 := GenS7{F0: 2}
                _ = v134
            }
            _ = v128
        }
    default:
        v127 = (v127 % (((v127 << ((18446744073709551615) & 63))) | 1))
    }
    fmt.Println(v127)
    v5 = v5
    var v135 uint8 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v127:", v127)
    fmt.Println("v135:", v135)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v126
    _ = v127
    _ = v135
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
