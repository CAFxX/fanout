package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint8; F2 bool }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64; F1 uint; F2 int }
type GenS9 struct { F0 int; F1 string }
type GenS10 struct { F0 uint32; F1 int8 }
type GenS11 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS12 struct { F0 uint32; F1 string; F2 int32 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = -78
    var v3 int = 0
    var v4 uint8 = 1
    for fi5 := 0; fi5 < 11; fi5++ {
        _ = len(v1)
        v2 += v2
        if (!(0 <= 340)) {
            v6 := GenS1{F0: "a", F1: -1}
            fmt.Println(v3, v2)
            _ = v6
        } else {
            if (115 >= 4294967295) {
                var v7 uint32 = 2
                v7 += (^(v7))
                _ = v7
            } else {
                var v8 float32 = float32((0.1 - 3.14159))
                if true {
                    if (true && (!(false && (false || false)))) {
                        var v9 uint = 0
                        _ = v9
                    } else {
                        var v10 uint64 = 0
                        _ = len(v1)
                        v4 += v4
                        _ = v10
                    }
                    v3 -= 32768
                    var v11 float64 = float64(v8)
                    _ = v11
                }
                _ = v8
            }
            v12 := GenS1{F0: "", F1: 42}
            switch "" {
            case "a":
                var v13 uint32 = 65535
                _ = v13
            case "hello":
                v14 := GenS1{F0: "a", F1: int16(fi5)}
                _ = v14
            default:
                fmt.Println(v3)
            }
            _ = v12
        }
        _ = fi5
    }
    fmt.Println(v3, v2, v2)
    if (!(0 >= 255)) {
        if false {
            v3 = (v3 >> ((198) & 63))
            v15 := []int{v3, (v3 ^ v3)}
            _ = v15
        }
        var v16 bool = true
        var v17 int64 = -195
        _ = v16
        _ = v17
    } else {
        v2 = 42
        v4 -= (v4 - 42)
        var v18 int16 = -32768
        _ = v18
    }
    switch "hello" {
    case "foo bar":
        delete(v1, "foo bar")
    case "":
        v19 := GenS1{F0: "a", F1: 0}
        _ = v19
    default:
        if true {
            v4 -= v4
            v20 := GenS1{F0: "foo bar", F1: 1000}
            _ = v20
        } else {
            if true {
                for fi21 := 0; fi21 < 6; fi21++ {
                    v2 -= v2
                    for fi22 := 0; fi22 < 15; fi22++ {
                        switch "\n\t" {
                        case "foo bar":
                            var v23 int32 = int32(v2)
                            _ = v23
                        case "hello":
                            fmt.Println(v2, v2, fi21)
                        case "a":
                            if true {
                                switch "" {
                                case "hello":
                                    v3 = 1
                                case "foo bar":
                                    switch (v3 | fi21) {
                                    case 1:
                                        fmt.Println(fi22, v3, fi21)
                                    default:
                                        v24 := []int{(^(fi21)), fi21, 42}
                                        _ = v24
                                    }
                                    if true {
                                        _ = len(v1)
                                        var v25 float64 = (-2.25 + float64(0.1))
                                        _ = v25
                                    } else {
                                        if (!false) {
                                            v26 := GenS1{F0: ("\n\t" + ""), F1: 117}
                                            v3 = (v3 ^ v3)
                                            switch ("" + "a") {
                                            case "":
                                                fmt.Println(v4, fi22)
                                            case "hello":
                                                _ = len(v1)
                                            default:
                                                switch (v3 | int(v4)) {
                                                case -1:
                                                    switch "\n\t" {
                                                    case "a":
                                                        v3 = (fi21 * fi22)
                                                        var v27 uint = 255
                                                        switch "foo bar" {
                                                        case "":
                                                            v28 := GenS1{F0: "hello", F1: 1}
                                                            _ = v28
                                                        case "foo bar":
                                                            v2 -= (v2 - v2)
                                                        default:
                                                            v3 = fi22
                                                        }
                                                        v3 = int(v4)
                                                        var v29 uint8 = 53
                                                        switch (("" + "hello") + ("" + ("\n\t" + ("\n\t" + ("foo bar" + "\n\t"))))) {
                                                        case "":
                                                            v30 := GenS2{F0: 1, F1: 127}
                                                            _ = v30
                                                        default:
                                                            v4 = uint8(v27)
                                                        }
                                                        v31 := GenS2{F0: -1, F1: v4}
                                                        var v32 int16 = -1
                                                        _ = v27
                                                        _ = v29
                                                        _ = v31
                                                        _ = v32
                                                    case "\n\t":
                                                        var v33 float64 = 1.5
                                                        _ = v33
                                                    case "foo bar":
                                                        v34 := GenS2{F0: 1, F1: (v4 | v4)}
                                                        _ = v34
                                                    }
                                                    v35 := GenS1{F0: "a", F1: -445}
                                                    switch ("foo bar" + (("a" + "") + "a")) {
                                                    case "\n\t":
                                                        fmt.Println(v2, v3)
                                                    case "foo bar":
                                                        fmt.Println(v4, v2, v2)
                                                    }
                                                    _ = v35
                                                case 2147483647:
                                                    var v36 int16 = 0
                                                    _ = v36
                                                default:
                                                    var v37 uint8 = 1
                                                    var v38 uint32 = 89
                                                    _ = v37
                                                    _ = v38
                                                }
                                            }
                                            v39 := []int{fi21}
                                            _ = v26
                                            _ = v39
                                        } else {
                                            v3 -= v3
                                            var v40 uint64 = 1
                                            _ = v40
                                        }
                                        switch "\n\t" {
                                        case "":
                                            fmt.Println(v3)
                                        default:
                                            v41 := v1["a"]
                                            _ = v41
                                        }
                                    }
                                    v3 = (^(v3))
                                case "\n\t":
                                    v3 = (fi22 % ((9223372036854775807) | 1))
                                default:
                                    if (737 < 848) {
                                        var v42 uint64 = 4294967295
                                        var v43 int8 = (v2 ^ 42)
                                        _ = v42
                                        _ = v43
                                    } else {
                                        if false {
                                            var v44 float64 = 0.1
                                            _ = v44
                                        }
                                        v45 := GenS2{F0: -824, F1: v4}
                                        v46 := GenS3{F0: 42, F1: (v4 * v4), F2: false}
                                        v2 -= (v2 / ((88) | 1))
                                        _ = v45
                                        _ = v46
                                    }
                                }
                                if true {
                                    var v47 float32 = 0.0
                                    v48 := []int{fi22, (fi21 - fi21), (v3 | v3), 943}
                                    _ = v47
                                    _ = v48
                                } else {
                                    if (-85 != 32767) {
                                        switch ("foo bar" + "hello") {
                                        case "foo bar":
                                            var v49 bool = ((-(v3)) < fi21)
                                            var v50 int32 = -736
                                            _ = v49
                                            _ = v50
                                        default:
                                            switch "foo bar" {
                                            case "\n\t":
                                                v51 := GenS3{F0: -1, F1: (v4 % (((^(v4))) | 1)), F2: (true || (0 <= 1))}
                                                _ = v51
                                            case "hello":
                                                if false {
                                                    if false {
                                                        switch (fi22 ^ (^(fi21))) {
                                                        case -1:
                                                            if true {
                                                                v52 := []int{(fi22 ^ 671)}
                                                                fmt.Println(fi22)
                                                                _ = v52
                                                            }
                                                        case 42:
                                                            v4 -= v4
                                                        case 63:
                                                            v53 := GenS1{F0: "\n\t", F1: 0}
                                                            _ = v53
                                                        }
                                                        var v54 uint16 = 100
                                                        if true {
                                                            var v55 uint16 = v54
                                                            switch ("foo bar" + "foo bar") {
                                                            case "a":
                                                                v1[(("hello" + ((("foo bar" + "foo bar") + ("a" + "a")) + "hello")) + "hello")] = fi22
                                                            default:
                                                                _ = len(v1)
                                                            }
                                                            if false {
                                                                var v56 uint64 = 1
                                                                var v57 uint = 0
                                                                _ = v56
                                                                _ = v57
                                                            } else {
                                                                if true {
                                                                    fmt.Println(v55)
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        if (!false) {
                                                                            v2 = v2
                                                                            v55 += v55
                                                                            switch v3 {
                                                                            case -1:
                                                                                switch "a" {
                                                                                case "a":
                                                                                    v58 := GenS4{F0: (v3 != -172), F1: 272}
                                                                                    _ = v58
                                                                                default:
                                                                                    v3 -= (v3 << ((2) & 63))
                                                                                }
                                                                            default:
                                                                                v59 := GenS4{F0: false, F1: 2147483647}
                                                                                _ = v59
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v3)
                                                                        }
                                                                    case "":
                                                                        v60 := GenS3{F0: -1, F1: v4, F2: (float64(0.1) > -2.25)}
                                                                        _ = v60
                                                                    case "hello":
                                                                        v61 := GenS5{F0: v55, F1: true}
                                                                        _ = v61
                                                                    }
                                                                } else {
                                                                    var v62 int64 = -1
                                                                    _ = v62
                                                                }
                                                                v63 := []int{4294967295}
                                                                if len(v63) >= 2 {
                                                                    v63 = v63[1:len(v63)]
                                                                }
                                                                _ = v63
                                                            }
                                                            _ = v55
                                                        } else {
                                                            if false {
                                                                var v64 bool = true
                                                                v65 := GenS5{F0: 32767, F1: v64}
                                                                var v66 int8 = -1
                                                                _ = v64
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                        }
                                                        _ = v54
                                                    } else {
                                                        switch -857 {
                                                        case -2147483648:
                                                            switch "a" {
                                                            case "a":
                                                                if (true || (697 == 0)) {
                                                                    v3 += fi21
                                                                    if false {
                                                                        v3 = (fi21 << ((42) & 63))
                                                                    } else {
                                                                        if true {
                                                                            var v67 string = ""
                                                                            if false {
                                                                                if (v2 <= v2) {
                                                                                    v68 := GenS5{F0: 394, F1: false}
                                                                                    _ = v68
                                                                                }
                                                                                v67 = v67
                                                                            }
                                                                            _ = v67
                                                                        } else {
                                                                            if (!(0 < 0)) {
                                                                                v69 := []int{4294967295, (-(v3)), fi21}
                                                                                _ = v69
                                                                            } else {
                                                                                delete(v1, ("foo bar" + "\n\t"))
                                                                            }
                                                                            v4 += (v4 & (v4 << ((131) & 63)))
                                                                        }
                                                                    }
                                                                    v2 -= v2
                                                                }
                                                                v70 := GenS6{F0: 654, F1: v4}
                                                                _ = v70
                                                            case "hello":
                                                                var v71 int = -9223372036854775808
                                                                _ = v71
                                                            case "foo bar":
                                                                _ = len(v1)
                                                            default:
                                                                v72 := GenS7{F0: (fi21 >> ((42) & 63))}
                                                                _ = v72
                                                            }
                                                        default:
                                                            v73 := v1["foo bar"]
                                                            _ = v73
                                                        }
                                                        switch (fi21 % ((fi21) | 1)) {
                                                        case -1:
                                                            var v74 int16 = 42
                                                            _ = v74
                                                        case 4294967295:
                                                            var v75 uint64 = 18446744073709551615
                                                            _ = v75
                                                        case 0:
                                                            v3 += v3
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    } else {
                                        v3 += 32767
                                        v4 -= 2
                                        if (true || false) {
                                            v76 := GenS8{F0: 980, F1: 4294967295, F2: fi21}
                                            _ = v76
                                        } else {
                                            var v77 bool = (1.5 > 3.14159)
                                            fmt.Println(v77, v4)
                                            _ = v77
                                        }
                                    }
                                }
                            }
                        }
                        _ = fi22
                    }
                    _ = fi21
                }
                for fi78 := 0; fi78 < 14; fi78++ {
                    fmt.Println(fi78, fi78)
                    if true {
                        if (32767 <= 1000) {
                            switch (v3 & v3) {
                            case -2147483648:
                                _ = len(v1)
                            case 32767:
                                if (false || false) {
                                    if (42 > 42) {
                                        v2 -= (v2 ^ 0)
                                        var v79 float32 = -2.25
                                        _ = v79
                                    }
                                } else {
                                    v80 := GenS9{F0: 159, F1: (("" + ("a" + "\n\t")) + ("hello" + ("foo bar" + ("\n\t" + ("\n\t" + "")))))}
                                    if false {
                                        switch fi78 {
                                        case 4294967295:
                                            if true {
                                                var v81 int8 = v2
                                                v82 := GenS10{F0: 2147483648, F1: v81}
                                                _ = v81
                                                _ = v82
                                            }
                                        case 32767:
                                            v2 -= 30
                                        }
                                    }
                                    var v83 int64 = 9223372036854775807
                                    _ = v80
                                    _ = v83
                                }
                            case 32768:
                                fmt.Println(v2)
                            }
                            if (4294967295 < 65535) {
                                fmt.Println(v4, v4, v4)
                                v2 -= v2
                            } else {
                                v3 = fi78
                                fmt.Println(fi78, v3)
                                switch "\n\t" {
                                case "a":
                                    v3 = (v3 & fi78)
                                case "\n\t":
                                    var v84 int64 = 9223372036854775807
                                    _ = v84
                                case "hello":
                                    _ = len(v1)
                                }
                            }
                            switch "" {
                            case "foo bar":
                                v2 = v2
                            }
                        } else {
                            for fi85 := 0; fi85 < 11; fi85++ {
                                if false {
                                    v86 := GenS1{F0: (("" + "hello") + ""), F1: 1000}
                                    switch ("foo bar" + "a") {
                                    case "hello":
                                        var v87 bool = ("foo bar" >= "hello")
                                        _ = v87
                                    case "":
                                        switch 32767 {
                                        case -2147483648:
                                            fmt.Println(v2)
                                        case 0:
                                            var v88 int = 2147483647
                                            var v89 int64 = -1
                                            var v90 uint16 = 2
                                            var v91 int8 = v2
                                            _ = v88
                                            _ = v89
                                            _ = v90
                                            _ = v91
                                        }
                                    default:
                                        if (false || false) {
                                            var v92 int16 = -32768
                                            fmt.Println(v2)
                                            _ = v92
                                        } else {
                                            fmt.Println(v3, v2, v3)
                                            if ((0 >= 1) && true) {
                                                v2 += int8(v4)
                                                fmt.Println(fi78, v3, v3)
                                            }
                                            fmt.Println(fi85, fi78)
                                        }
                                    }
                                    v4 = 191
                                    _ = v86
                                } else {
                                    switch (v3 | (fi78 + (fi78 >> ((425) & 63)))) {
                                    case -483:
                                        if (0 >= 0) {
                                            v3 += v3
                                            if false {
                                                var v93 string = ("foo bar" + "a")
                                                fmt.Println(fi78)
                                                var v94 uint = 330
                                                _ = v93
                                                _ = v94
                                            } else {
                                                if ((42 < -1) && (1e10 >= -2.25)) {
                                                    if true {
                                                        fmt.Println(v2, fi78)
                                                    } else {
                                                        var v95 float32 = (0.0 * 1e10)
                                                        _ = v95
                                                    }
                                                } else {
                                                    v3 += (^(v3))
                                                    v4 = (v4 / (((v4 & (v4 & (^(v4))))) | 1))
                                                }
                                                switch 1 {
                                                case -442:
                                                    var v96 int16 = 1000
                                                    _ = v96
                                                case 1:
                                                    v97 := GenS5{F0: 42, F1: (!true)}
                                                    _ = v97
                                                default:
                                                    var v98 uint16 = 255
                                                    _ = v98
                                                }
                                            }
                                        }
                                    case 199:
                                        v99 := GenS4{F0: (!true), F1: -33}
                                        _ = v99
                                    default:
                                        if false {
                                            v4 -= v4
                                            v3 = fi85
                                        }
                                    }
                                    delete(v1, "")
                                }
                                _ = fi85
                            }
                        }
                        v100 := GenS11{F0: 1, F1: 960, F2: 395}
                        _ = v100
                    }
                    v101 := []int{(fi78 + v3), 9223372036854775807, -1, -736}
                    _ = fi78
                    _ = v101
                }
            }
            switch "\n\t" {
            case "foo bar":
                v3 = 32767
            case "a":
                for fi102 := 0; fi102 < 12; fi102++ {
                    var v103 float32 = 1e10
                    v104 := []int{(fi102 % (((v3 * fi102)) | 1)), (fi102 << ((4294967295) & 63))}
                    fmt.Println(v4, fi102, v103)
                    _ = fi102
                    _ = v103
                    _ = v104
                }
            case "":
                for fi105 := 0; fi105 < 12; fi105++ {
                    v106 := v1[("foo bar" + "foo bar")]
                    _ = len(v1)
                    for fi107 := 0; fi107 < 14; fi107++ {
                        v4 = 255
                        delete(v1, "hello")
                        _ = fi107
                    }
                    _ = fi105
                    _ = v106
                }
            default:
                v1[""] = v3
            }
        }
    }
    var v108 uint16 = 65535
    v3 += (v3 % (((v3 + v3)) | 1))
    fmt.Println(v2, v108)
    v1["\n\t"] = v3
    switch "a" {
    case "hello":
        v2 -= 100
    default:
        fmt.Println(v3)
    }
    v109 := GenS12{F0: 0, F1: ("\n\t" + ""), F2: 0}
    v4 -= v4
    v2 += (v2 << ((0) & 63))
    if false {
        v108 -= v108
        v110 := []int{(v3 ^ int(v4)), v3}
        _ = v110
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v108
    _ = v109
    _ = v2
    _ = v3
    _ = v4
}
