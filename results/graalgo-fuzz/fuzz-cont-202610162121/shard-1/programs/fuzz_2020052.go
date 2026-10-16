package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS8 struct { F0 int8; F1 int64; F2 int32 }
type GenS9 struct { F0 int; F1 string; F2 uint8 }
type GenS10 struct { F0 int32; F1 int64; F2 int8 }
func genFn0(p0 uint) uint32 {
    var v7 uint64 = 481
    var v8 string = "a"
    v8 = ("hello" + "hello")
    v8 = v8
    for fi9 := 0; fi9 < 13; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 uint8 = 127
            var v12 int = (fi9 | -1)
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return uint32(v7)
}
func genFn1(p0 int, p1 int16, p2 string) uint32 {
    for fi13 := 0; fi13 < 7; fi13++ {
        if false {
            p1 = (p1 | 32767)
            if false {
                v14 := GenS1{F0: 192}
                if ((1 >= 0) || (0 > 1)) {
                    for fi15 := 0; fi15 < 20; fi15++ {
                        fmt.Println(p2, fi13, fi13)
                        _ = fi15
                    }
                } else {
                    for fi16 := 0; fi16 < 19; fi16++ {
                        v17 := GenS2{F0: 0, F1: 38}
                        _ = fi16
                        _ = v17
                    }
                }
                fmt.Println(p2, p2)
                _ = v14
            } else {
                if (230 < 1) {
                    v18 := []int{(p0 | p0), (fi13 | (p0 >> ((42) & 63)))}
                    var v19 uint32 = genFn0(296)
                    p2 = p2
                    _ = v18
                    _ = v19
                }
                p2 = p2
                p1 -= (p1 ^ (p1 * p1))
            }
        }
        var v20 string = p2
        _ = fi13
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 float64 = (-2.25 - float64(3.14159))
            _ = fi22
            _ = v23
        }
        var v24 string = ("a" + "hello")
        _ = fi21
        _ = v24
    }
    fmt.Println(p1, p0, p0)
    if false {
        p1 = (p1 >> ((427) & 63))
        switch (p0 >> ((250) & 63)) {
        case 0:
            fmt.Println(p2, p2, p2)
        case 299:
            var v25 int = 128
            _ = v25
        case 2147483647:
            fmt.Println(p0, p0, p2)
        default:
            var v26 int8 = 0
            _ = v26
        }
        for fi27 := 0; fi27 < 19; fi27++ {
            fmt.Println(p1, p1)
            if true {
                var v28 int8 = -80
                switch 2147483647 {
                case -2147483648:
                    p0 = 128
                }
                v29 := []int{fi27, p0, -1, fi27}
                _ = v28
                _ = v29
            }
            switch p2 {
            case "a":
                if false {
                    v30 := GenS2{F0: 100, F1: 131}
                    _ = v30
                }
            default:
                if (false || false) {
                    for fi31 := 0; fi31 < 2; fi31++ {
                        v32 := []int{-2147483648}
                        switch (fi31 & (fi27 >> ((65535) & 63))) {
                        case 0:
                            if (4294967295 != 65535) {
                                var v33 uint = 4294967295
                                var v34 int = fi31
                                _ = v33
                                _ = v34
                            }
                            var v35 int8 = 72
                            v32 = append(v32, (p0 - (p0 & fi27)))
                            _ = v35
                        case 293:
                            v36 := GenS1{F0: 0}
                            var v37 float32 = float32((1.5 + 0.0))
                            p2 = p2
                            _ = v36
                            _ = v37
                        case 9223372036854775807:
                            var v38 uint8 = 94
                            _ = v38
                        default:
                            switch 637 {
                            case 0:
                                if (1.5 >= 0.0) {
                                    var v39 uint32 = 2147483647
                                    _ = v39
                                } else {
                                    var v40 float64 = float64(1.5)
                                    if false {
                                        var v41 uint32 = 42
                                        if false {
                                            v40 = v40
                                        }
                                        var v42 uint16 = 228
                                        _ = v41
                                        _ = v42
                                    } else {
                                        if (127 != int8(fi31)) {
                                            var v43 float64 = 1e10
                                            _ = v43
                                        }
                                        if (!((fi27 / (((p0 % ((fi31) | 1))) | 1)) < (p0 ^ (fi27 * (p0 | fi27))))) {
                                            var v44 uint32 = 42
                                            v45 := GenS1{F0: -1}
                                            _ = v44
                                            _ = v45
                                        }
                                    }
                                    _ = v40
                                }
                            case 128:
                                v46 := GenS1{F0: -9223372036854775808}
                                _ = v46
                            default:
                                if false {
                                    p1 -= -436
                                    fmt.Println(p1)
                                }
                            }
                        }
                        fmt.Println(p1)
                        var v47 uint = 4
                        if len(v32) >= 2 {
                            v32 = v32[1:len(v32)]
                        }
                        _ = fi31
                        _ = v32
                        _ = v47
                    }
                    fmt.Println(p2, p2, p0)
                    p0 += 0
                }
            }
            _ = fi27
        }
    } else {
        v48 := []int{(p0 % ((4294967295) | 1)), (p0 % (((^(p0))) | 1)), p0, -335}
        _ = v48
    }
    if ((false || ((!(p2 > (("" + "\n\t") + p2))) || false)) || (1 > -2147483648)) {
        var v49 int64 = 931
        v50 := GenS3{F0: p1}
        switch (p2 + p2) {
        case "foo bar":
            v51 := []int{0, (p0 | (p0 * 42))}
            _ = v51
        case "":
            for fi52 := 0; fi52 < 8; fi52++ {
                var v53 string = p2
                v53 = p2
                var v54 uint64 = 2
                _ = fi52
                _ = v53
                _ = v54
            }
        case "a":
            if ((v49 * (^(v49))) < (v49 + 0)) {
                v55 := GenS3{F0: 0}
                _ = v55
            } else {
                for fi56 := 0; fi56 < 10; fi56++ {
                    v57 := []int{p0}
                    switch p2 {
                    case "":
                        v58 := GenS2{F0: 1, F1: 2}
                        _ = v58
                    }
                    v59 := GenS4{F0: 18446744073709551615}
                    _ = fi56
                    _ = v57
                    _ = v59
                }
            }
        default:
            var v60 float32 = 3.14159
            _ = v60
        }
        _ = v49
        _ = v50
    }
    _ = p0
    _ = p1
    _ = p2
    return genFn0(2)
}

func main() {
    var v1 uint16 = 2
    v2 := []string{}
    var v3 int = 1
    var v4 int = v3
    v5 := make(map[string]int)
    var v6 uint8 = 4
    var v61 uint = 341
    switch 1 {
    case 34:
        if true {
            var v62 int = -658
            _ = v62
        }
    default:
        v63 := GenS3{F0: 906}
        _ = v63
    }
    if (-32768 > -32768) {
        if (-2147483648 < -767) {
            for fi64 := 0; fi64 < 2; fi64++ {
                var v65 int32 = 42
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi64
                _ = v65
            }
        }
        if len(v2) > 0 {
            v2[0] = ((("" + "\n\t") + ("a" + (("hello" + "hello") + ""))) + "foo bar")
        }
        switch "hello" {
        case "foo bar":
            v61 = v61
        case "":
            if false {
                var v66 float32 = float32((0.1 + 1.5))
                var v67 bool = (42 <= 42)
                _ = v66
                _ = v67
            }
        }
    } else {
        v68 := GenS4{F0: v61}
        if ((true && (!false)) && true) {
            var v69 int8 = 42
            if true {
                v5[("\n\t" + "hello")] = (v4 % ((2147483647) | 1))
                v70 := GenS5{F0: (v6 / (((v6 << ((65535) & 63))) | 1)), F1: v69}
                fmt.Println(v6, v69)
                _ = v70
            }
            _ = v69
        } else {
            fmt.Println(v61, v61, v6)
            for fi71 := 0; fi71 < 6; fi71++ {
                if ("hello" == "foo bar") {
                    var v72 float32 = (3.14159 / 3.14159)
                    _ = v72
                } else {
                    var v73 uint32 = 2147483648
                    var v74 float32 = 0.0
                    v4 -= fi71
                    _ = v73
                    _ = v74
                }
                var v75 uint64 = 233
                v1 += 42
                _ = fi71
                _ = v75
            }
            switch 1 {
            case 2147483647:
                switch ("" + ("foo bar" + "\n\t")) {
                case "hello":
                    if ((v3 <= 912) || true) {
                        for fi76 := 0; fi76 < 2; fi76++ {
                            if (true && (!true)) {
                                switch ("hello" + "foo bar") {
                                case "a":
                                    fmt.Println(v3)
                                case "foo bar":
                                    if true {
                                        if ((v1 <= (v1 | uint16(v61))) || false) {
                                            var v77 string = "\n\t"
                                            _ = v77
                                        }
                                    } else {
                                        if false {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            fmt.Println(v3)
                                            v3 -= (v4 << ((v61) & 63))
                                        } else {
                                            delete(v5, ("" + "foo bar"))
                                            v78 := v5["foo bar"]
                                            _ = v78
                                        }
                                        var v79 uint32 = 2147483648
                                        _ = v79
                                    }
                                default:
                                    switch ("hello" + ("a" + "a")) {
                                    case "\n\t":
                                        if true {
                                            v80 := GenS6{F0: 0}
                                            v5["hello"] = (v3 | (^(v4)))
                                            var v81 bool = true
                                            _ = v80
                                            _ = v81
                                        } else {
                                            switch v3 {
                                            case 4294967295:
                                                for fi82 := 0; fi82 < 7; fi82++ {
                                                    v83 := GenS3{F0: 988}
                                                    _ = fi82
                                                    _ = v83
                                                }
                                            case 42:
                                                v3 = fi76
                                            default:
                                                if (2 <= 475) {
                                                    v84 := GenS3{F0: 42}
                                                    if false {
                                                        for fi85 := 0; fi85 < 8; fi85++ {
                                                            var v86 uint32 = 323
                                                            _ = fi85
                                                            _ = v86
                                                        }
                                                        if ((false || false) && true) {
                                                            v87 := GenS4{F0: 614}
                                                            if false {
                                                                for fi88 := 0; fi88 < 10; fi88++ {
                                                                    switch (fi88 + (fi88 | (-(fi76)))) {
                                                                    case 4294967295:
                                                                        fmt.Println(v1, v6)
                                                                    case 2147483647:
                                                                        var v89 string = "foo bar"
                                                                        v90 := GenS6{F0: 0}
                                                                        v91 := GenS4{F0: (v61 % (((v61 % (((v61 & v61)) | 1))) | 1))}
                                                                        _ = v89
                                                                        _ = v90
                                                                        _ = v91
                                                                    case 0:
                                                                        var v92 int64 = -506
                                                                        _ = v92
                                                                    }
                                                                    fmt.Println(v61, v4)
                                                                    _ = len(v5)
                                                                    _ = fi88
                                                                }
                                                            } else {
                                                                var v93 int8 = -1
                                                                _ = v93
                                                            }
                                                            v61 -= v61
                                                            _ = v87
                                                        }
                                                    } else {
                                                        fmt.Println(v61, v4, v3)
                                                    }
                                                    if (((v6 - (v6 % (((v6 >> (((^(v61))) & 63))) | 1))) < 255) && (258 < (v3 & (v3 >> ((v61) & 63))))) {
                                                        if false {
                                                            if true {
                                                                fmt.Println(v4, fi76)
                                                                v6 += 42
                                                            } else {
                                                                v3 = (-(fi76))
                                                            }
                                                            switch v3 {
                                                            case 42:
                                                                var v94 uint64 = 950
                                                                _ = v94
                                                            }
                                                            fmt.Println(v61, fi76, fi76)
                                                        } else {
                                                            var v95 int64 = 1
                                                            var v96 string = "hello"
                                                            var v97 float32 = float32((3.14159 - (float64(3.14159) - 1e10)))
                                                            _ = v95
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                    }
                                                    _ = v84
                                                }
                                            }
                                        }
                                    case "foo bar":
                                        _ = len(v2)
                                    default:
                                        if (false || true) {
                                            for fi98 := 0; fi98 < 7; fi98++ {
                                                delete(v5, "")
                                                if true {
                                                    v99 := GenS3{F0: 1000}
                                                    if len(v2) > 0 {
                                                        v2[0] = ("\n\t" + ("\n\t" + ("" + "foo bar")))
                                                    }
                                                    _ = v99
                                                }
                                                _ = fi98
                                            }
                                            v100 := GenS5{F0: (v6 / ((255) | 1)), F1: -1}
                                            var v101 int = (v3 / ((32767) | 1))
                                            _ = v100
                                            _ = v101
                                        } else {
                                            if (-2.25 < 0.0) {
                                                v102 := GenS7{F0: 2, F1: 18446744073709551615, F2: 304}
                                                _ = v102
                                            }
                                            v103 := GenS8{F0: 127, F1: 9223372036854775807, F2: 2147483647}
                                            _ = v103
                                        }
                                    }
                                }
                            } else {
                                fmt.Println(v61)
                                if (!false) {
                                    v1 = 495
                                    v104 := GenS3{F0: 1000}
                                    delete(v5, "foo bar")
                                    _ = v104
                                }
                            }
                            switch -1 {
                            case -9223372036854775808:
                                var v105 int16 = 1
                                _ = v105
                            case 32767:
                                switch 0 {
                                case 2147483647:
                                    var v106 float64 = float64((3.14159 / 0.1))
                                    _ = v106
                                case 42:
                                    v4 += (fi76 * (v4 & v4))
                                default:
                                    v107 := GenS9{F0: 4294967295, F1: (((("foo bar" + "") + (("\n\t" + "") + "hello")) + "a") + ("\n\t" + ("" + ((("foo bar" + ("foo bar" + "a")) + "") + ("\n\t" + "a"))))), F2: (^(v6))}
                                    _ = v107
                                }
                            }
                            _ = len(v2)
                            _ = fi76
                        }
                    }
                case "\n\t":
                    for fi108 := 0; fi108 < 16; fi108++ {
                        for fi109 := 0; fi109 < 5; fi109++ {
                            fmt.Println(fi108)
                            _ = fi109
                        }
                        var v110 uint = v61
                        if false {
                            fmt.Println(v6)
                            fmt.Println(v110)
                            for fi111 := 0; fi111 < 16; fi111++ {
                                fmt.Println(v1)
                                _ = fi111
                            }
                        } else {
                            v5[(("\n\t" + "hello") + "a")] = fi108
                            if (-9223372036854775808 <= -9223372036854775808) {
                                fmt.Println(v110, v4, v61)
                                var v112 float64 = (0.1 - 0.0)
                                _ = v112
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = ("" + "\n\t")
                                }
                                v2 = append(v2, "")
                            }
                        }
                        _ = fi108
                        _ = v110
                    }
                case "foo bar":
                    switch v3 {
                    case 127:
                        v6 -= v6
                    case -1:
                        v4 = v3
                    case 872:
                        v4 = v4
                    default:
                        if false {
                            fmt.Println(v61, v6, v4)
                            v113 := GenS3{F0: 0}
                            v3 -= (v3 >> (((v61 >> (((v61 * (v61 - 42))) & 63))) & 63))
                            _ = v113
                        }
                    }
                default:
                    v6 = (^(v6))
                }
            case 32768:
                fmt.Println(v3, v4)
            }
        }
        _ = v68
    }
    fmt.Println(v1)
    if false {
        switch "hello" {
        case "a":
            if (v61 > (v61 + (^(v61)))) {
                v2 = append(v2, "\n\t")
            }
        default:
            if (uint(v1) == v61) {
                v114 := GenS9{F0: v3, F1: "\n\t", F2: v6}
                if true {
                    v1 -= 1
                    switch v3 {
                    case 32767:
                        v2 = append(v2, "\n\t")
                    case 4294967295:
                        v1 = (v1 % ((v1) | 1))
                    }
                }
                for fi115 := 0; fi115 < 16; fi115++ {
                    v116 := v5[(("\n\t" + "foo bar") + "\n\t")]
                    v1 -= v1
                    for fi117 := 0; fi117 < 1; fi117++ {
                        v1 += v1
                        if true {
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                            v3 -= v4
                            if true {
                                switch ("\n\t" + ("\n\t" + "hello")) {
                                case "foo bar":
                                    v5[""] = (v116 ^ (v4 * 128))
                                case "a":
                                    v61 = v61
                                default:
                                    v61 += v61
                                }
                                v4 = 0
                                v118 := GenS7{F0: 65535, F1: v61, F2: 42}
                                if true {
                                    _ = len(v2)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                } else {
                                    var v119 int64 = 741
                                    v3 = v4
                                    v116 += fi117
                                    v116 = fi117
                                    _ = v119
                                }
                                _ = v118
                            }
                        }
                        _ = fi117
                    }
                    _ = fi115
                    _ = v116
                }
                _ = v114
            }
        }
        v4 -= (v3 % ((v4) | 1))
        switch v3 {
        case 893:
            for fi120 := 0; fi120 < 18; fi120++ {
                v6 = (^(v6))
                _ = fi120
            }
        case 128:
            v121 := GenS2{F0: -128, F1: v6}
            _ = v121
        }
    } else {
        switch (v3 >> ((255) & 63)) {
        case 32767:
            v6 = v6
        }
    }
    var v122 uint64 = 42
    if (!false) {
        if false {
            v123 := GenS5{F0: (^(v6)), F1: 1}
            v124 := GenS10{F0: -39, F1: -1, F2: -128}
            _ = v123
            _ = v124
        }
        delete(v5, ("a" + "\n\t"))
    } else {
        v125 := GenS4{F0: v61}
        v6 -= v6
        _ = v125
    }
    v126 := GenS10{F0: 2147483647, F1: -484, F2: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v122:", v122)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v122
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
}
