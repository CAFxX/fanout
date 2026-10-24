package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint32; F2 string }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint8; F1 int; F2 int8 }
type GenS9 struct { F0 uint64; F1 int64; F2 string }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int32; F1 int64 }
func genFn0(p0 string, p1 int32) uint8 {
    for fi5 := 0; fi5 < 9; fi5++ {
        fmt.Println(fi5, p1)
        for fi6 := 0; fi6 < 11; fi6++ {
            p1 = (-(p1))
            p1 -= p1
            switch p0 {
            case "foo bar":
                fmt.Println(p1)
            default:
                v7 := []int{fi5, fi6}
                _ = v7
            }
            switch 364 {
            case 42:
                var v8 float32 = 0.0
                _ = v8
            case 1:
                p1 = (^(p1))
            default:
                p1 = p1
            }
            _ = fi6
        }
        _ = fi5
    }
    p1 -= (p1 & p1)
    _ = p0
    _ = p1
    return 161
}
func genFn1(p0 int32, p1 uint8, p2 uint64) bool {
    fmt.Println(p0, p0, p0)
    v9 := []int{-1}
    if len(v9) > 0 {
        v9[0] = 32768
    }
    switch "hello" {
    case "":
        var v10 uint8 = (p1 % (((^(p1))) | 1))
        _ = v10
    case "\n\t":
        var v11 uint8 = (p1 << ((65535) & 63))
        _ = v11
    case "hello":
        v9 = append(v9, -46)
    default:
        p0 += p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (32767 == 65535)
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]int)
    var v3 int64 = 9223372036854775807
    v4 := []string{"foo bar", (("hello" + "") + "a")}
    v3 += 0
    if ((100 > 127) && (-1 != (v3 | (v3 % (((^(v3))) | 1))))) {
        var v12 int16 = -32768
        if (-2147483648 < 42) {
            v1 += (v1 & v1)
            var v13 int = 590
            v3 = (^(v3))
            _ = v13
        }
        v14 := GenS1{F0: 18446744073709551615}
        _ = v12
        _ = v14
    }
    switch "hello" {
    case "foo bar":
        var v15 int8 = -128
        _ = v15
    case "\n\t":
        v16 := GenS2{F0: 2}
        _ = v16
    case "":
        _ = len(v4)
    default:
        for fi17 := 0; fi17 < 19; fi17++ {
            v1 += (v1 - v1)
            v3 += 9223372036854775807
            _ = fi17
        }
    }
    var v18 string = "\n\t"
    var v19 uint64 = (^(v1))
    v20 := GenS3{F0: 128, F1: 1, F2: 1}
    var v21 string = "a"
    if true {
        var v22 uint16 = 255
        switch -953 {
        case -505:
            switch v21 {
            case "foo bar":
                v23 := GenS4{F0: v3, F1: 42, F2: ((v21 + v21) + v21)}
                _ = v23
            }
        }
        _ = v22
    } else {
        switch v21 {
        case "a":
            fmt.Println(v3, v21)
        default:
            fmt.Println(v3, v18)
        }
        switch ((("" + ("a" + "\n\t")) + "a") + v18) {
        case "":
            if (false || (false || (!(42 > 32768)))) {
                v18 = "a"
                if (!true) {
                    var v24 uint32 = 2147483648
                    fmt.Println(v1, v21, v24)
                    _ = v24
                }
            }
        case "a":
            if (255 < 379) {
                fmt.Println(v18)
            }
        case "hello":
            for fi25 := 0; fi25 < 14; fi25++ {
                _ = len(v4)
                if (false || true) {
                    var v26 uint32 = 921
                    for fi27 := 0; fi27 < 11; fi27++ {
                        switch (((v18 + ("foo bar" + (v21 + ("" + v18)))) + "a") + (v21 + v21)) {
                        case "":
                            fmt.Println(v19)
                        case "a":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case "foo bar":
                            v19 = 42
                        }
                        if true {
                            var v28 int16 = -1
                            _ = v28
                        } else {
                            _ = len(v4)
                            var v29 int16 = 317
                            if ((v21 + (v18 + (("hello" + ((("foo bar" + v18) + v18) + "\n\t")) + "a"))) == v21) {
                                v18 = v21
                                v30 := GenS2{F0: 764}
                                v3 -= (v3 % ((int64(v19)) | 1))
                                _ = v30
                            } else {
                                var v31 uint16 = uint16(v26)
                                var v32 int16 = int16(fi25)
                                fmt.Println(fi25, v3)
                                _ = v31
                                _ = v32
                            }
                            _ = v29
                        }
                        _ = fi27
                    }
                    v33 := GenS4{F0: (v3 & v3), F1: 926, F2: v21}
                    _ = v26
                    _ = v33
                } else {
                    var v34 string = "a"
                    _ = v34
                }
                var v35 int16 = -714
                _ = fi25
                _ = v35
            }
        }
        var v36 int16 = 850
        _ = v36
    }
    for fi37 := 0; fi37 < 19; fi37++ {
        if (1e10 < 0.1) {
            var v38 uint64 = v1
            _ = v38
        } else {
            switch (fi37 % (((fi37 + 0)) | 1)) {
            case 535:
                for fi39 := 0; fi39 < 8; fi39++ {
                    var v40 string = (v21 + (v21 + v18))
                    _ = fi39
                    _ = v40
                }
            case 32768:
                fmt.Println(v3, v1)
            }
            var v41 uint32 = 0
            if false {
                v18 = "hello"
                fmt.Println(v3, v3)
                v21 = v21
            }
            _ = v41
        }
        var v42 int16 = 1
        for fi43 := 0; fi43 < 10; fi43++ {
            if (3.14159 < (1e10 + 0.1)) {
                var v44 int = fi43
                if false {
                    fmt.Println(v19)
                    if (2 > 1) {
                        var v45 int8 = 27
                        fmt.Println(v19)
                        _ = v45
                    } else {
                        var v46 float32 = (float32((float64(1.5) + ((0.0 + 1.5) / 1e10))) + (0.0 - (0.0 - 3.14159)))
                        fmt.Println(v46, fi37)
                        _ = v46
                    }
                    switch fi37 {
                    case 32768:
                        if true {
                            var v47 bool = (18446744073709551615 != 49)
                            var v48 int64 = v3
                            _ = v47
                            _ = v48
                        } else {
                            v49 := GenS5{F0: 127, F1: 0}
                            if (!true) {
                                v4 = append(v4, (v21 + v21))
                                v21 = ((v18 + (((v21 + "hello") + (((v18 + v18) + "foo bar") + v21)) + v18)) + (v18 + "foo bar"))
                            }
                            _ = v49
                        }
                    case 2147483647:
                        if true {
                            if false {
                                switch v18 {
                                case "\n\t":
                                    var v50 int16 = v42
                                    _ = v50
                                default:
                                    v18 = "\n\t"
                                }
                                fmt.Println(v44)
                                v44 -= (-(v44))
                            }
                            delete(v2, 443)
                        } else {
                            _ = len(v2)
                            var v51 uint = 777
                            var v52 int = (v44 ^ fi43)
                            _ = v51
                            _ = v52
                        }
                    default:
                        v21 = "a"
                    }
                }
                var v53 int16 = 0
                _ = v44
                _ = v53
            } else {
                switch "" {
                case "a":
                    v1 = (v19 * (^(v19)))
                case "hello":
                    var v54 float64 = (1e10 * (1.5 / 1e10))
                    _ = v54
                case "foo bar":
                    if ((v19 % ((18446744073709551615) | 1)) >= v1) {
                        if (-1 == -1) {
                            v3 = v3
                            v55 := GenS6{F0: v21, F1: 65535}
                            fmt.Println(v1)
                            _ = v55
                        }
                    } else {
                        if ("" >= (v18 + (("foo bar" + v21) + v21))) {
                            var v56 int16 = v42
                            if (v56 != v56) {
                                v18 = "foo bar"
                                v57 := GenS6{F0: ("" + v21), F1: 0}
                                v56 = 42
                                _ = v57
                            } else {
                                v4 = append(v4, "\n\t")
                                _ = len(v2)
                                var v58 float64 = ((0.0 + (3.14159 * 0.0)) - float64(1e10))
                                _ = v58
                            }
                            v3 -= (-(v3))
                            v21 = ("" + v21)
                            _ = v56
                        } else {
                            var v59 int8 = 127
                            _ = v59
                        }
                        v18 = (v18 + "\n\t")
                        switch fi43 {
                        case 162:
                            var v60 float64 = 1.5
                            if (!(2 < 1)) {
                                switch 127 {
                                case 127:
                                    v2[-910] = 98
                                case -9223372036854775808:
                                    v61 := GenS6{F0: (v18 + v18), F1: 239}
                                    _ = v61
                                case -116:
                                    if (9223372036854775807 < 274) {
                                        var v62 uint16 = 466
                                        v18 = v18
                                        v63 := GenS7{F0: 0}
                                        _ = v62
                                        _ = v63
                                    }
                                }
                                var v64 int = 4294967295
                                switch (fi43 % (((v64 & v64)) | 1)) {
                                case 32768:
                                    var v65 int32 = 42
                                    _ = v65
                                }
                                _ = v64
                            }
                            if (2 == 0) {
                                fmt.Println(v42, fi37)
                            } else {
                                switch fi37 {
                                case -2147483648:
                                    switch (^(fi37)) {
                                    case -555:
                                        var v66 uint64 = (v19 + (^(v19)))
                                        var v67 uint16 = 65535
                                        _ = v66
                                        _ = v67
                                    case -324:
                                        v2[(fi43 | 293)] = (fi37 >> ((18446744073709551615) & 63))
                                    case 213:
                                        var v68 uint32 = 2147483648
                                        _ = v68
                                    }
                                default:
                                    fmt.Println(fi37, v18)
                                }
                                if ((fi43 * 42) >= (fi37 << ((65535) & 63))) {
                                    v69 := v2[-9223372036854775808]
                                    _ = v69
                                }
                                var v70 uint32 = 2147483648
                                fmt.Println(v70)
                                fmt.Println(v42, v18, v19)
                                _ = v70
                            }
                            _ = v60
                        case -754:
                            switch (v18 + v21) {
                            case "a":
                                if (!(18446744073709551615 == 42)) {
                                    if (v3 >= v3) {
                                        if ((65535 <= 2147483647) && ((^(v1)) == v19)) {
                                            _ = len(v4)
                                            v42 = v42
                                            v71 := GenS8{F0: 128, F1: 4294967295, F2: -1}
                                            _ = v71
                                        }
                                        var v72 uint16 = 1
                                        v19 = (^(v1))
                                        _ = v72
                                    }
                                    switch fi43 {
                                    case 980:
                                        var v73 uint64 = (v19 - v19)
                                        _ = v73
                                    case 9223372036854775807:
                                        v3 = 0
                                        v19 = 18446744073709551615
                                        switch v21 {
                                        case "hello":
                                            if false {
                                                switch fi37 {
                                                case 9223372036854775807:
                                                    switch (fi37 / ((1) | 1)) {
                                                    case 1:
                                                        v2[128] = -2147483648
                                                    case -1:
                                                        switch v18 {
                                                        case "":
                                                            v3 += (v3 ^ (v3 | (v3 >> ((255) & 63))))
                                                        }
                                                    }
                                                    v74 := GenS7{F0: 683}
                                                    _ = v74
                                                case -1:
                                                    var v75 int16 = v42
                                                    _ = v75
                                                case 127:
                                                    fmt.Println(v3, v42)
                                                }
                                                var v76 uint16 = 32767
                                                _ = v76
                                            } else {
                                                if ((^(v3)) != v3) {
                                                    if (true && true) {
                                                        if false {
                                                            v42 += (v42 ^ v42)
                                                            if false {
                                                                fmt.Println(fi43, fi37)
                                                                v77 := GenS6{F0: v21, F1: 2}
                                                                _ = v77
                                                            } else {
                                                                switch fi43 {
                                                                case 2147483647:
                                                                    switch (v18 + (("a" + v21) + ("a" + "a"))) {
                                                                    case "\n\t":
                                                                        if (!(295 < 2147483647)) {
                                                                            fmt.Println(v21, v19)
                                                                            v78 := GenS9{F0: v1, F1: v3, F2: (v18 + ("\n\t" + v21))}
                                                                            v18 = ("" + (("hello" + "\n\t") + v21))
                                                                            var v79 string = "\n\t"
                                                                            _ = v78
                                                                            _ = v79
                                                                        }
                                                                    }
                                                                case 997:
                                                                    v1 += (v1 / ((v19) | 1))
                                                                default:
                                                                    var v80 uint = uint(v19)
                                                                    _ = v80
                                                                }
                                                                v3 -= -9223372036854775808
                                                            }
                                                        } else {
                                                            var v81 uint8 = genFn0(((("hello" + v21) + (v21 + v21)) + v18), -138)
                                                            if false {
                                                                v1 -= v19
                                                            }
                                                            fmt.Println(v19, v81, v21)
                                                            _ = v81
                                                        }
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    } else {
                                                        v19 = 0
                                                        _ = len(v2)
                                                        _ = len(v2)
                                                    }
                                                } else {
                                                    v82 := v2[127]
                                                    _ = v82
                                                }
                                                var v83 uint8 = 1
                                                var v84 int64 = (^(v3))
                                                _ = v83
                                                _ = v84
                                            }
                                        case "foo bar":
                                            v19 = (v19 | v19)
                                        case "a":
                                            v21 = "foo bar"
                                            v85 := GenS5{F0: 81, F1: 2}
                                            _ = v85
                                        }
                                    case -866:
                                        fmt.Println(v1)
                                        v86 := GenS6{F0: "hello", F1: 0}
                                        _ = v86
                                    }
                                } else {
                                    if (255 > 2) {
                                        fmt.Println(v42)
                                        _ = len(v4)
                                        var v87 string = "a"
                                        var v88 uint = 71
                                        _ = v87
                                        _ = v88
                                    }
                                    delete(v2, fi43)
                                }
                            default:
                                v18 = ""
                            }
                        case -588:
                            var v89 int8 = 127
                            v90 := v2[fi37]
                            _ = v89
                            _ = v90
                        }
                    }
                default:
                    v3 += -9223372036854775808
                }
                v42 += v42
                v4 = append(v4, ((((v21 + (((("" + (v18 + ((v21 + v21) + (((((v21 + "foo bar") + v18) + "\n\t") + (v21 + v18)) + ("a" + ((((v21 + "foo bar") + ("foo bar" + v18)) + v18) + v21)))))) + (("hello" + v18) + v18)) + "hello") + v18)) + v21) + "foo bar") + "a"))
            }
            var v91 int16 = (v42 & v42)
            if false {
                v19 -= (v1 + v1)
                var v92 uint32 = 65535
                var v93 string = "hello"
                fmt.Println(v93, v18)
                _ = v92
                _ = v93
            } else {
                v94 := GenS2{F0: 284}
                if (((fi43 > 42) || ((v1 & 298) > (v19 * 2))) && (!(!(725 >= (v91 * (v42 ^ 934)))))) {
                    v95 := GenS1{F0: 100}
                    v96 := GenS1{F0: 631}
                    _ = v95
                    _ = v96
                } else {
                    if false {
                        var v97 uint32 = 2
                        if (!true) {
                            v21 = "\n\t"
                            switch (fi43 - int(v97)) {
                            case 338:
                                fmt.Println(v97, v18, v1)
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            var v98 uint64 = (^(v1))
                            switch v21 {
                            case "\n\t":
                                v42 = (-(v91))
                            }
                            switch (v21 + ("\n\t" + "\n\t")) {
                            case "":
                                if (!(true || false)) {
                                    v99 := GenS3{F0: 2, F1: 2, F2: 2}
                                    v100 := GenS9{F0: v19, F1: (v3 / ((v3) | 1)), F2: (v18 + v21)}
                                    fmt.Println(v98, fi37, v97)
                                    _ = v100
                                    _ = v99
                                }
                            }
                            _ = v98
                        }
                        if ((false || (v97 == (v97 | 1))) && false) {
                            var v101 int16 = v42
                            _ = v101
                        }
                        v4 = append(v4, v21)
                        v102 := GenS10{F0: fi37}
                        v4 = append(v4, v18)
                        if false {
                            if false {
                                v1 -= v1
                                v2[int(v3)] = fi37
                            } else {
                                if false {
                                    if ((!true) || false) {
                                        v103 := GenS2{F0: v97}
                                        _ = v103
                                    }
                                }
                                switch (fi43 % ((1) | 1)) {
                                case 83:
                                    var v104 int16 = (v91 | v91)
                                    _ = v104
                                case 1:
                                    fmt.Println(v3, v42, v21)
                                default:
                                    fmt.Println(v18, v42)
                                }
                            }
                            var v105 int = fi37
                            _ = v105
                        }
                        _ = v102
                        _ = v97
                    } else {
                        v91 += v91
                    }
                    v106 := GenS2{F0: 4294967295}
                    _ = v106
                }
                switch (("" + v21) + v21) {
                case "foo bar":
                    var v107 uint = 0
                    _ = v107
                case "a":
                    v4 = append(v4, ("" + v21))
                case "\n\t":
                    switch (fi37 / ((-954) | 1)) {
                    case 925:
                        v19 = v19
                    case -871:
                        v3 = 42
                    default:
                        fmt.Println(fi37, fi43)
                    }
                    switch (^(fi37)) {
                    case 2147483647:
                        v42 += -574
                        v108 := GenS11{F0: 1, F1: v3}
                        _ = v108
                    case -575:
                        if false {
                            fmt.Println(v91)
                            switch (fi37 ^ -118) {
                            case 1:
                                var v109 int32 = 42
                                _ = v109
                            case -144:
                                v110 := GenS6{F0: ("foo bar" + v18), F1: 1}
                                _ = v110
                            case 3:
                                switch (v18 + v18) {
                                case "hello":
                                    if len(v4) > 0 {
                                        v4[0] = v18
                                    }
                                case "\n\t":
                                    v18 = "\n\t"
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v111 := GenS4{F0: (v3 & v3), F1: 42, F2: "\n\t"}
                                    _ = v111
                                }
                            default:
                                v1 = (v19 / ((0) | 1))
                            }
                            var v112 uint64 = 71
                            _ = v112
                        }
                    }
                    v19 = 18446744073709551615
                }
                _ = v94
            }
            fmt.Println(v18, v18, v1)
            _ = fi43
            _ = v91
        }
        _ = fi37
        _ = v42
    }
    var v113 uint8 = 255
    for fi114 := 0; fi114 < 5; fi114++ {
        if true {
            var v115 string = ("a" + (v21 + "hello"))
            var v116 int64 = -1
            v117 := GenS6{F0: "", F1: 42}
            _ = v115
            _ = v116
            _ = v117
        } else {
            v3 = v3
        }
        var v118 uint8 = uint8(v19)
        _ = fi114
        _ = v118
    }
    if false {
        if false {
            v18 = "foo bar"
        } else {
            for fi119 := 0; fi119 < 11; fi119++ {
                switch ("foo bar" + (v21 + ("a" + "a"))) {
                case "foo bar":
                    delete(v2, fi119)
                }
                _ = fi119
            }
        }
    }
    v3 += (v3 << ((42) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v113
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
