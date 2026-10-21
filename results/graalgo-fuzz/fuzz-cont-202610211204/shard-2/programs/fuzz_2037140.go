package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint8; F1 string; F2 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint; F1 int16; F2 int }
type GenS10 struct { F0 int64; F1 uint64 }
type GenS11 struct { F0 int; F1 uint32; F2 int8 }
type GenS12 struct { F0 uint16 }
func genFn0(p0 int8, p1 uint8, p2 int64) string {
    for fi7 := 0; fi7 < 11; fi7++ {
        if false {
            p1 = (p1 >> ((492) & 63))
        }
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 bool = false
            v10 := GenS1{F0: p1, F1: 213, F2: (p2 - (p2 % (((p2 % (((p2 << ((18446744073709551615) & 63))) | 1))) | 1)))}
            p2 = p2
            _ = fi8
            _ = v10
            _ = v9
        }
        switch fi7 {
        case 9223372036854775807:
            v11 := []int{-826, fi7}
            _ = v11
        }
        _ = fi7
    }
    var v12 float32 = (1.5 / 1.0)
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := []int{fi13}
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return ("a" + ("foo bar" + "a"))
}
func genFn1(p0 uint) int8 {
    v15 := []int{127}
    p0 = p0
    _ = p0
    _ = v15
    return -92
}

func main() {
    var v1 int = 32767
    var v2 uint16 = 255
    var v3 int32 = 115
    var v4 int64 = 0
    v5 := make(map[string]string)
    v6 := make(map[uint32]uint32)
    v16 := GenS2{F0: 2, F1: -468}
    for fi17 := 0; fi17 < 16; fi17++ {
        switch ((("" + ("\n\t" + (("" + "\n\t") + ("foo bar" + "a")))) + (("a" + "") + "foo bar")) + "foo bar") {
        case "\n\t":
            v18 := []int{}
            _ = v18
        case "a":
            var v19 float32 = (float32(0.0) - 0.0)
            _ = v19
        case "foo bar":
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v3)
            case "hello":
                var v20 int = (^(v1))
                _ = v20
            default:
                v4 = v4
            }
        default:
            for fi21 := 0; fi21 < 2; fi21++ {
                fmt.Println(v2)
                _ = fi21
            }
        }
        if (!(true && false)) {
            delete(v6, 616)
            if false {
                switch fi17 {
                case 711:
                    switch "" {
                    case "foo bar":
                        for fi22 := 0; fi22 < 13; fi22++ {
                            v2 -= v2
                            _ = fi22
                        }
                    case "hello":
                        var v23 uint64 = 665
                        _ = v23
                    case "\n\t":
                        v2 = v2
                    default:
                        v24 := v6[807]
                        _ = v24
                    }
                case 9223372036854775807:
                    for fi25 := 0; fi25 < 7; fi25++ {
                        fmt.Println(v2, fi25)
                        v26 := []int{(v1 / ((v1) | 1)), (fi17 & fi25)}
                        _ = fi25
                        _ = v26
                    }
                case -603:
                    fmt.Println(v2, v2, v1)
                default:
                    v6[4294967295] = 2147483647
                }
                switch ("hello" + "") {
                case "":
                    var v27 int32 = v3
                    _ = v27
                }
            } else {
                v28 := []int{}
                fmt.Println(v2, v4)
                v3 = v3
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 3; fi29++ {
            v1 = fi17
            var v30 uint16 = (v2 | (v2 | 884))
            _ = fi29
            _ = v30
        }
        _ = fi17
    }
    var v31 float32 = 1e10
    for fi32 := 0; fi32 < 18; fi32++ {
        fmt.Println(v31)
        v2 = 65535
        v31 = -2.25
        _ = fi32
    }
    if (478 <= 588) {
        if (int64(v2) > v4) {
            switch (v1 - v1) {
            case -265:
                v33 := GenS2{F0: 127, F1: 0}
                _ = v33
            case 32768:
                var v34 uint32 = 4294967295
                _ = v34
            default:
                v35 := []int{(v1 % ((9223372036854775807) | 1)), v1, (^(v1))}
                _ = v35
            }
            fmt.Println(v1, v2)
            v36 := GenS2{F0: 171, F1: 42}
            _ = v36
        }
        v37 := GenS3{F0: v1, F1: (v3 >> ((2) & 63))}
        _ = v37
    }
    switch ((("hello" + "") + "hello") + (("foo bar" + ((("foo bar" + "hello") + "\n\t") + ("foo bar" + ("" + ("a" + "\n\t"))))) + "hello")) {
    case "\n\t":
        v2 = 58
    }
    var v38 float32 = v31
    switch int(v4) {
    case 32768:
        for fi39 := 0; fi39 < 19; fi39++ {
            var v40 bool = (!true)
            var v41 int16 = 32767
            switch 2147483647 {
            case 1:
                v3 -= v3
            case 127:
                v42 := GenS4{F0: 128, F1: (("\n\t" + (((("foo bar" + ("\n\t" + "")) + "foo bar") + ("a" + "foo bar")) + ("hello" + (("" + "\n\t") + ((("\n\t" + "\n\t") + ("a" + "\n\t")) + (("\n\t" + "a") + ("foo bar" + "foo bar"))))))) + "foo bar"), F2: fi39}
                _ = v42
            }
            _ = fi39
            _ = v40
            _ = v41
        }
    case -297:
        v43 := []int{v1, -2147483648, v1, v1}
        _ = v43
    case 32767:
        fmt.Println(v3)
    }
    v44 := []int{(v1 + v1), (v1 << ((0) & 63)), -2147483648, (v1 - 2147483647)}
    v45 := v5[""]
    v46 := GenS3{F0: (v1 << ((553) & 63)), F1: v3}
    var v47 bool = (!((v1 & (v1 >> ((18446744073709551615) & 63))) >= v1))
    v48 := GenS5{F0: v1}
    for fi49 := 0; fi49 < 4; fi49++ {
        if (!v47) {
            switch "hello" {
            case "hello":
                v50 := GenS1{F0: 2, F1: (v3 | int32(v4)), F2: (v4 % (((v4 + v4)) | 1))}
                _ = v50
            case "\n\t":
                if (((42 != 0) || v47) && ((269 <= 1) && true)) {
                    _ = len(v44)
                }
            default:
                delete(v6, 4294967295)
            }
        } else {
            v38 -= float32(1.5)
        }
        v51 := GenS5{F0: fi49}
        switch -214 {
        case -400:
            if true {
                v4 = v4
            } else {
                v52 := GenS2{F0: 255, F1: 0}
                if (677 < 4294967295) {
                    var v53 uint32 = 832
                    for fi54 := 0; fi54 < 5; fi54++ {
                        var v55 uint64 = 2
                        _ = fi54
                        _ = v55
                    }
                    _ = v53
                }
                switch "\n\t" {
                case "foo bar":
                    var v56 int32 = -1
                    _ = v56
                case "":
                    if v47 {
                        for fi57 := 0; fi57 < 9; fi57++ {
                            v3 += 2147483647
                            fmt.Println(v1, v38)
                            _ = fi57
                        }
                        for fi58 := 0; fi58 < 6; fi58++ {
                            fmt.Println(v1, fi49)
                            _ = fi58
                        }
                        var v59 int = fi49
                        _ = v59
                    } else {
                        var v60 float64 = -2.25
                        _ = v60
                    }
                case "\n\t":
                    var v61 int32 = 0
                    _ = v61
                }
                _ = v52
            }
        default:
            fmt.Println(v47)
        }
        _ = fi49
        _ = v51
    }
    var v62 float64 = (((-2.25 - 0.0) - (3.14159 * (1e10 * (-2.25 / -2.25)))) / ((1.5 * -2.25) + (0.0 * -2.25)))
    var v63 uint16 = (v2 & (v2 + v2))
    if v47 {
        switch v45 {
        case "\n\t":
            switch v1 {
            case -2147483648:
                for fi64 := 0; fi64 < 18; fi64++ {
                    v4 += v4
                    _ = fi64
                }
            case 32768:
                if true {
                    var v65 int32 = (^(v3))
                    var v66 float32 = (0.0 / float32(v62))
                    _ = v65
                    _ = v66
                }
            default:
                v47 = v47
            }
        case "a":
            if (("foo bar" + v45) < "") {
                for fi67 := 0; fi67 < 7; fi67++ {
                    switch v45 {
                    case "foo bar":
                        var v68 int8 = 1
                        _ = v68
                    case "hello":
                        _ = len(v44)
                    }
                    if false {
                        v62 = v62
                        switch ((v45 + "a") + "") {
                        case "\n\t":
                            v69 := GenS3{F0: (v1 & v1), F1: -662}
                            _ = v69
                        case "a":
                            for fi70 := 0; fi70 < 12; fi70++ {
                                var v71 uint64 = 4294967295
                                switch (v1 / ((v1) | 1)) {
                                case 32768:
                                    switch fi70 {
                                    case -9223372036854775808:
                                        if true {
                                            v72 := GenS2{F0: 128, F1: 617}
                                            v38 = v31
                                            if len(v44) > 0 {
                                                v44[0] = 1
                                            }
                                            fmt.Println(v2, v31, fi67)
                                            var v73 uint = 980
                                            _ = v72
                                            _ = v73
                                        } else {
                                            v1 = (^(fi70))
                                        }
                                        if v47 {
                                            if v47 {
                                                if true {
                                                    v74 := GenS6{F0: (-(v4)), F1: (fi70 / ((int(v4)) | 1))}
                                                    switch 9223372036854775807 {
                                                    case 450:
                                                        if v47 {
                                                            v2 = v2
                                                            fmt.Println(v2)
                                                            switch v45 {
                                                            case "\n\t":
                                                                fmt.Println(v47, fi70)
                                                            case "a":
                                                                var v75 int = fi70
                                                                _ = v75
                                                            case "":
                                                                var v76 string = ("\n\t" + (v45 + ((v45 + (("" + "") + v45)) + ((((v45 + "foo bar") + (v45 + v45)) + v45) + ((v45 + "a") + v45)))))
                                                                _ = v76
                                                            }
                                                        }
                                                    case 32768:
                                                        _ = len(v6)
                                                    }
                                                    _ = v74
                                                } else {
                                                    if (!(!v47)) {
                                                        fmt.Println(v47, fi67, v71)
                                                        v77 := GenS7{F0: 127}
                                                        _ = v77
                                                    } else {
                                                        v44 = append(v44, v1)
                                                        v78 := GenS7{F0: 28}
                                                        v79 := GenS3{F0: v1, F1: v3}
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                    v80 := GenS5{F0: (fi67 + (fi67 - fi67))}
                                                    _ = v80
                                                }
                                                switch (("hello" + "hello") + v45) {
                                                case "foo bar":
                                                    if true {
                                                        v47 = (4294967295 <= v71)
                                                    }
                                                    v44 = append(v44, 32768)
                                                default:
                                                    if (127 < 1) {
                                                        var v81 uint16 = uint16(v71)
                                                        _ = v81
                                                    } else {
                                                        if (!false) {
                                                            v82 := GenS2{F0: 42, F1: 1}
                                                            v71 -= (v71 % (((v71 + v71)) | 1))
                                                            v2 = (v2 - uint16(v71))
                                                            _ = v82
                                                        } else {
                                                            switch 183 {
                                                            case -628:
                                                                if v47 {
                                                                    v83 := GenS8{F0: 2}
                                                                    if true {
                                                                        v5[(("hello" + v45) + (((v45 + "\n\t") + v45) + v45))] = v45
                                                                    }
                                                                    fmt.Println(v71, v1)
                                                                    _ = v83
                                                                } else {
                                                                    var v84 float64 = v62
                                                                    switch v45 {
                                                                    case "a":
                                                                        v45 = v45
                                                                    default:
                                                                        var v85 bool = v47
                                                                        _ = v85
                                                                    }
                                                                    v86 := GenS6{F0: (v4 | 42), F1: v1}
                                                                    _ = v84
                                                                    _ = v86
                                                                }
                                                            case -2147483648:
                                                                if (v47 && false) {
                                                                    v62 = 0.1
                                                                } else {
                                                                    var v87 int16 = -1
                                                                    _ = len(v44)
                                                                    var v88 uint = 2
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                            default:
                                                                v89 := GenS2{F0: 171, F1: -546}
                                                                _ = v89
                                                            }
                                                        }
                                                        var v90 int16 = 0
                                                        _ = v90
                                                    }
                                                }
                                                _ = len(v5)
                                            }
                                            v91 := GenS9{F0: 0, F1: -197, F2: 32768}
                                            var v92 int32 = (v3 + 1)
                                            _ = v91
                                            _ = v92
                                        }
                                    default:
                                        var v93 int64 = 130
                                        _ = v93
                                    }
                                    var v94 int16 = 962
                                    v63 -= (v63 + v63)
                                    _ = v94
                                case -9223372036854775808:
                                    fmt.Println(v45, v2)
                                }
                                _ = fi70
                                _ = v71
                            }
                        }
                        if v47 {
                            var v95 int64 = int64(v3)
                            var v96 string = v45
                            v2 = v2
                            _ = v95
                            _ = v96
                        } else {
                            var v97 int64 = (v4 - -1)
                            if len(v44) > 0 {
                                v44[0] = 396
                            }
                            _ = v97
                        }
                    } else {
                        if v47 {
                            for fi98 := 0; fi98 < 16; fi98++ {
                                v99 := GenS10{F0: v4, F1: 1}
                                v38 -= 0.0
                                _ = fi98
                                _ = v99
                            }
                            fmt.Println(v1, v62, v45)
                        }
                    }
                    _ = fi67
                }
                fmt.Println(v4, v63)
                for fi100 := 0; fi100 < 5; fi100++ {
                    switch "hello" {
                    case "hello":
                        for fi101 := 0; fi101 < 8; fi101++ {
                            v4 = v4
                            fmt.Println(v1, v2, v1)
                            _ = fi101
                        }
                    case "foo bar":
                        var v102 uint16 = v63
                        _ = v102
                    }
                    var v103 string = "foo bar"
                    v104 := GenS5{F0: v1}
                    _ = fi100
                    _ = v103
                    _ = v104
                }
            }
        case "hello":
            for fi105 := 0; fi105 < 10; fi105++ {
                v106 := GenS4{F0: 30, F1: "foo bar", F2: fi105}
                for fi107 := 0; fi107 < 18; fi107++ {
                    v5["\n\t"] = "\n\t"
                    switch (("hello" + v45) + "foo bar") {
                    case "\n\t":
                        fmt.Println(v38, v47)
                    case "a":
                        v108 := GenS11{F0: fi107, F1: 1, F2: 0}
                        if v47 {
                            v62 = ((1.5 - v62) - v62)
                            if (!(479 <= 0)) {
                                v2 = 65535
                            } else {
                                if true {
                                    if false {
                                        fmt.Println(fi107, v3, fi105)
                                    } else {
                                        v62 = 1.5
                                        var v109 int16 = 0
                                        switch fi107 {
                                        case -9223372036854775808:
                                            var v110 uint16 = 2
                                            _ = v110
                                        case 4294967295:
                                            var v111 uint32 = 1
                                            _ = v111
                                        default:
                                            v112 := GenS12{F0: v2}
                                            _ = v112
                                        }
                                        _ = v109
                                    }
                                    var v113 float32 = v38
                                    v113 += 1e10
                                    _ = v113
                                }
                                switch v1 {
                                case -2147483648:
                                    var v114 float32 = v31
                                    _ = v114
                                }
                                v2 += (v2 - 65535)
                            }
                        } else {
                            fmt.Println(v62, v45)
                            v4 += (v4 >> ((0) & 63))
                        }
                        _ = v108
                    default:
                        v115 := GenS1{F0: 186, F1: -2147483648, F2: v4}
                        _ = v115
                    }
                    _ = fi107
                }
                _ = fi105
                _ = v106
            }
        }
        if v47 {
            v45 = v45
        }
        v1 -= (v1 & (^(v1)))
    } else {
        v6[2147483648] = 2
        var v116 int = v1
        if v47 {
            v116 += v116
        }
        _ = v116
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v38
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v62
    _ = v63
}
