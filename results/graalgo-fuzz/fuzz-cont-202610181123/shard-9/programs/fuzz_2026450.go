package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int64; F1 int8; F2 string }
type GenS3 struct { F0 int8; F1 string; F2 int16 }
type GenS4 struct { F0 int; F1 uint; F2 uint64 }
type GenS5 struct { F0 int64; F1 bool; F2 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 int16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 uint16; F1 bool }
type GenS11 struct { F0 int; F1 int8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int8 }
func genFn0() int {
    v4 := GenS1{F0: 499, F1: false}
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := []int{fi5, 9223372036854775807, fi5}
        var v7 bool = (float64(1.5) == -2.25)
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint = 4294967295
    v9 := GenS1{F0: 462, F1: (((-2.25 / ((-2.25 - float32(0.0)) / 0.1)) + (1.5 * 1.5)) <= ((1.5 / -2.25) - 0.1))}
    _ = v4
    _ = v8
    _ = v9
    return 32767
}
func genFn1(p0 uint64, p1 int16) int32 {
    if (true || ((65535 <= 632) || (2147483647 != 42))) {
        var v10 int16 = 42
        var v11 float32 = 3.14159
        _ = v10
        _ = v11
    } else {
        var v12 uint8 = 2
        p0 += (^(p0))
        var v13 uint16 = 200
        _ = v12
        _ = v13
    }
    var v14 int32 = -2147483648
    _ = p0
    _ = p1
    _ = v14
    return v14
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    v3 := []string{"\n\t", (("a" + "\n\t") + (("a" + "\n\t") + "foo bar")), "\n\t", "a"}
    if (0 > 42) {
        var v15 int32 = -1
        _ = v15
    } else {
        if (0 >= -2147483648) {
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 float32 = -2.25
                if ((true && (true && (true && (!true)))) || true) {
                    if false {
                        _ = len(v3)
                        v17 = v17
                        v17 += v17
                    }
                    switch "\n\t" {
                    case "\n\t":
                        var v18 int8 = 42
                        _ = v18
                    case "a":
                        var v19 int = 414
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 2; fi20++ {
                        var v21 uint = 838
                        _ = fi20
                        _ = v21
                    }
                } else {
                    switch fi16 {
                    case 9223372036854775807:
                        var v22 int8 = 0
                        _ = v22
                    case -43:
                        v17 -= (3.14159 * 1e10)
                    case -89:
                        fmt.Println(fi16)
                    }
                    var v23 uint16 = 42
                    for fi24 := 0; fi24 < 18; fi24++ {
                        fmt.Println(fi16, fi16, v17)
                        if (65535 > 40) {
                            fmt.Println(v17)
                        }
                        var v25 int16 = -1
                        _ = fi24
                        _ = v25
                    }
                    _ = v23
                }
                _ = fi16
                _ = v17
            }
            if (50 >= -128) {
                if false {
                    for fi26 := 0; fi26 < 12; fi26++ {
                        fmt.Println(fi26)
                        var v27 float32 = 3.14159
                        if true {
                            if len(v3) > 0 {
                                v3[0] = ""
                            }
                            if ((fi26 - fi26) != fi26) {
                                v3 = append(v3, "foo bar")
                            } else {
                                var v28 int = (fi26 / ((fi26) | 1))
                                if (!false) {
                                    v29 := GenS1{F0: 785, F1: true}
                                    _ = v29
                                }
                                _ = v28
                            }
                            if (false && false) {
                                for fi30 := 0; fi30 < 11; fi30++ {
                                    v31 := GenS1{F0: 746, F1: false}
                                    _ = fi30
                                    _ = v31
                                }
                            } else {
                                v32 := GenS1{F0: 32767, F1: (((v27 - v27) * (v27 - (3.14159 + float32(((((0.0 / 1e10) * 0.1) / (0.1 * 0.1)) / 1.5))))) <= 0.1)}
                                var v33 float64 = (0.1 + (1e10 + 3.14159))
                                _ = v32
                                _ = v33
                            }
                        }
                        _ = fi26
                        _ = v27
                    }
                }
                var v34 uint64 = 2
                v35 := GenS2{F0: 1, F1: 42, F2: "foo bar"}
                _ = v34
                _ = v35
            }
            var v36 uint8 = 155
            _ = v36
        } else {
            for fi37 := 0; fi37 < 16; fi37++ {
                switch genFn0() {
                case 171:
                    switch "" {
                    case "a":
                        fmt.Println(fi37)
                    }
                }
                var v38 uint16 = 32767
                _ = fi37
                _ = v38
            }
            var v39 int = 42
            v40 := GenS3{F0: -1, F1: "\n\t", F2: -32768}
            _ = v39
            _ = v40
        }
        var v41 string = ""
        _ = v41
    }
    var v42 uint64 = 4294967295
    for fi43 := 0; fi43 < 10; fi43++ {
        var v44 uint = 1
        v45 := v1["hello"]
        _ = fi43
        _ = v44
        _ = v45
    }
    if false {
        v46 := GenS4{F0: 2147483647, F1: 588, F2: v42}
        v47 := GenS5{F0: -528, F1: false, F2: -885}
        _ = v46
        _ = v47
    } else {
        fmt.Println(v42)
        var v48 uint64 = (v42 - v42)
        v49 := GenS3{F0: 100, F1: "foo bar", F2: -486}
        _ = v48
        _ = v49
    }
    fmt.Println(v42)
    v42 -= (v42 * v42)
    v42 = v42
    fmt.Println(v42)
    v42 = 881
    var v50 float64 = 0.1
    if (0 < 42) {
        if (((-277 >= -9223372036854775808) || false) && ((^(v42)) > v42)) {
            for fi51 := 0; fi51 < 10; fi51++ {
                v42 = (^(v42))
                _ = fi51
            }
        } else {
            v42 = (v42 / ((v42) | 1))
        }
    }
    v42 = v42
    if false {
        var v52 string = "\n\t"
        v1[v52] = v52
        var v53 bool = true
        _ = v52
        _ = v53
    } else {
        if len(v3) > 0 {
            v3[0] = (("hello" + "a") + "a")
        }
    }
    switch "\n\t" {
    case "":
        fmt.Println(v50, v42)
    case "\n\t":
        v42 += 42
    case "a":
        v54 := v2["\n\t"]
        _ = v54
    default:
        if (true && ((float32(v50) * float32(float64((-2.25 * (1.5 / 3.14159))))) > 1.5)) {
            for fi55 := 0; fi55 < 10; fi55++ {
                var v56 int32 = 2147483647
                fmt.Println(v42, v56)
                _ = fi55
                _ = v56
            }
            for fi57 := 0; fi57 < 3; fi57++ {
                if true {
                    if true {
                        _ = len(v3)
                    }
                    var v58 int16 = 791
                    v59 := GenS6{F0: 890}
                    _ = v58
                    _ = v59
                }
                if true {
                    v42 += v42
                    var v60 bool = (!false)
                    _ = v60
                } else {
                    var v61 int32 = 1
                    var v62 uint32 = 2
                    v63 := v2["hello"]
                    _ = v61
                    _ = v62
                    _ = v63
                }
                _ = fi57
            }
        }
    }
    switch -2147483648 {
    case 32768:
        for fi64 := 0; fi64 < 9; fi64++ {
            switch (fi64 - (fi64 - fi64)) {
            case 128:
                var v65 uint8 = 255
                _ = v65
            case 1:
                if true {
                    for fi66 := 0; fi66 < 2; fi66++ {
                        v42 = v42
                        v67 := GenS7{F0: v42, F1: -559}
                        _ = fi66
                        _ = v67
                    }
                    var v68 float64 = (0.1 + 3.14159)
                    _ = v68
                } else {
                    var v69 uint16 = 2
                    switch (fi64 >> ((0) & 63)) {
                    case 32767:
                        var v70 uint = 987
                        _ = v70
                    case 9223372036854775807:
                        v2["a"] = 42
                    case 2147483647:
                        v71 := GenS4{F0: (fi64 ^ (fi64 << ((59) & 63))), F1: 422, F2: 1}
                        _ = v71
                    }
                    _ = v69
                }
            case -675:
                for fi72 := 0; fi72 < 6; fi72++ {
                    var v73 uint16 = 255
                    v42 -= (v42 >> ((1) & 63))
                    v42 -= v42
                    _ = fi72
                    _ = v73
                }
            default:
                if false {
                    switch fi64 {
                    case -1:
                        var v74 bool = true
                        _ = v74
                    case 2147483647:
                        var v75 uint16 = uint16(v42)
                        _ = v75
                    case 1:
                        if (65535 > 255) {
                            switch "a" {
                            case "hello":
                                v50 -= v50
                            default:
                                if len(v3) > 0 {
                                    v3[0] = ""
                                }
                            }
                            var v76 int16 = 32
                            v77 := GenS8{F0: "hello"}
                            _ = v76
                            _ = v77
                        }
                    }
                    if false {
                        var v78 float32 = -2.25
                        _ = v78
                    }
                } else {
                    var v79 bool = false
                    _ = v79
                }
            }
            var v80 uint64 = v42
            _ = fi64
            _ = v80
        }
    }
    if true {
        v50 -= v50
    } else {
        for fi81 := 0; fi81 < 13; fi81++ {
            switch "" {
            case "a":
                if true {
                    var v82 int8 = 127
                    _ = v82
                } else {
                    var v83 uint = 0
                    v42 = v42
                    for fi84 := 0; fi84 < 11; fi84++ {
                        switch (-(fi84)) {
                        case -9223372036854775808:
                            v83 = 65535
                        case 689:
                            switch fi81 {
                            case 42:
                                var v85 int32 = genFn1((^(v42)), -144)
                                _ = v85
                            case -2147483648:
                                fmt.Println(v83)
                                v86 := GenS8{F0: "foo bar"}
                                var v87 float64 = ((v50 + 3.14159) * float64(3.14159))
                                _ = v86
                                _ = v87
                            default:
                                if (!(!true)) {
                                    var v88 float64 = float64(3.14159)
                                    var v89 bool = ((-299 <= fi81) && false)
                                    _ = v88
                                    _ = v89
                                } else {
                                    v50 -= v50
                                    v90 := GenS8{F0: ""}
                                    _ = v90
                                }
                            }
                        }
                        var v91 uint8 = 215
                        switch (fi81 ^ (-(fi84))) {
                        case -135:
                            switch 0 {
                            case -9223372036854775808:
                                if (0.1 < float32(float64(((1e10 + float32(v50)) + 1.5)))) {
                                    v92 := GenS2{F0: 1, F1: 67, F2: (("hello" + "a") + "hello")}
                                    _ = v92
                                }
                            case -940:
                                v91 = (^(v91))
                            }
                        case -245:
                            var v93 int16 = -32768
                            _ = v93
                        }
                        v83 = v83
                        _ = fi84
                        _ = v91
                    }
                    _ = v83
                }
            }
            for fi94 := 0; fi94 < 4; fi94++ {
                v42 = (v42 | 558)
                v95 := GenS9{F0: 199, F1: 42}
                if false {
                    var v96 int8 = -51
                    v42 = v42
                    switch "\n\t" {
                    case "":
                        switch "a" {
                        case "\n\t":
                            v50 = (v50 - 1.5)
                        }
                        fmt.Println(v50, fi94, v50)
                        v96 += v96
                        var v97 uint16 = 891
                        if (870 < -65) {
                            v50 = (v50 - -2.25)
                            var v98 bool = false
                            _ = v98
                        }
                        _ = v97
                    case "a":
                        fmt.Println(v96)
                    case "foo bar":
                        if true {
                            if false {
                                if false {
                                    v99 := GenS10{F0: 757, F1: false}
                                    switch fi81 {
                                    case 284:
                                        _ = len(v3)
                                    case 225:
                                        v3 = append(v3, (("foo bar" + ("foo bar" + "foo bar")) + "a"))
                                    case 128:
                                        switch -575 {
                                        case 9223372036854775807:
                                            v50 -= (1e10 + v50)
                                        case -1:
                                            switch (fi94 % ((fi94) | 1)) {
                                            case -1:
                                                switch (^(fi81)) {
                                                case 9223372036854775807:
                                                    var v100 float64 = float64(3.14159)
                                                    _ = v100
                                                case 121:
                                                    var v101 uint = 1
                                                    _ = v101
                                                case -9223372036854775808:
                                                    v102 := GenS11{F0: (fi94 | fi81), F1: 20}
                                                    _ = v102
                                                }
                                            case 128:
                                                var v103 int = (fi94 | (fi94 - (fi94 * (fi94 / ((fi81) | 1)))))
                                                _ = v103
                                            case -2147483648:
                                                v50 = v50
                                            default:
                                                v50 = -2.25
                                            }
                                        case 128:
                                            v50 = 0.1
                                        }
                                    default:
                                        switch 478 {
                                        case -9223372036854775808:
                                            v96 = (v96 / (((v96 - 1)) | 1))
                                        case 4294967295:
                                            v104 := GenS8{F0: "hello"}
                                            _ = v104
                                        default:
                                            v50 -= 1.5
                                        }
                                    }
                                    var v105 uint16 = 2
                                    _ = v105
                                    _ = v99
                                }
                                fmt.Println(v50)
                                fmt.Println(fi81, v42, fi81)
                            } else {
                                fmt.Println(fi94)
                            }
                            v106 := GenS3{F0: (v96 + (v96 >> ((947) & 63))), F1: "hello", F2: 0}
                            v50 += float64(((0.0 * 1e10) - float32(((v50 * 3.14159) + v50))))
                            switch (fi81 | fi81) {
                            case 128:
                                var v107 float64 = (-2.25 / v50)
                                _ = v107
                            case -858:
                                fmt.Println(v50)
                            default:
                                var v108 int64 = -9223372036854775808
                                _ = v108
                            }
                            v3 = append(v3, "")
                            _ = v106
                        } else {
                            v109 := GenS11{F0: (fi94 / (((fi81 | (fi94 | 127))) | 1)), F1: 68}
                            _ = v109
                        }
                    default:
                        if true {
                            v110 := GenS12{F0: "foo bar"}
                            var v111 uint8 = 255
                            v111 = uint8(v42)
                            if true {
                                switch "" {
                                case "foo bar":
                                    var v112 int32 = 42
                                    _ = v112
                                }
                                var v113 uint64 = (v42 & (v42 << ((uint(v111)) & 63)))
                                v42 += (v113 & uint64(v111))
                                _ = v113
                            }
                            fmt.Println(v50)
                            v114 := GenS3{F0: (v96 ^ (v96 << ((1) & 63))), F1: "foo bar", F2: 158}
                            fmt.Println(v42, v111)
                            _ = v110
                            _ = v111
                            _ = v114
                        } else {
                            v96 += -128
                            v115 := GenS13{F0: v96}
                            _ = v115
                        }
                    }
                    _ = v96
                }
                _ = fi94
                _ = v95
            }
            _ = fi81
        }
    }
    var v116 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v116:", v116)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v42:", v42)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v116
    _ = v2
    _ = v3
    _ = v42
    _ = v50
}
