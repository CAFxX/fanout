package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int16; F1 bool; F2 string }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int64; F1 uint16; F2 string }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 bool; F1 bool; F2 uint }
type GenS11 struct { F0 int; F1 uint16; F2 uint32 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int8; F1 string }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint64; F1 string; F2 uint8 }

func main() {
    var v1 uint8 = 2
    var v2 bool = ("hello" > "\n\t")
    var v3 uint32 = 42
    var v4 float32 = 0.1
    v5 := make(map[uint32]string)
    v4 += ((v4 * (3.14159 + (((v4 * v4) - v4) - float32((0.0 * 1.5))))) / v4)
    if (v3 >= v3) {
        v6 := GenS1{F0: 67, F1: (v3 % ((904) | 1))}
        _ = v6
    }
    var v7 uint = 1
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            if (688 < 0) {
                delete(v5, (v3 & (v3 & v3)))
            } else {
                var v10 int8 = -1
                switch "" {
                case "foo bar":
                    fmt.Println(v10, v10, v10)
                case "a":
                    v11 := []int{fi9, (fi8 / ((-337) | 1)), (fi9 & (fi8 << ((v7) & 63)))}
                    _ = v11
                case "hello":
                    v2 = v2
                }
                var v12 float32 = (v4 - (v4 + v4))
                _ = v10
                _ = v12
            }
            if (719 != 42) {
                if (-32 != -54) {
                    if (0.0 > 3.14159) {
                        v13 := GenS2{F0: (v1 << ((255) & 63))}
                        if v2 {
                            v7 += uint(v1)
                            v14 := GenS2{F0: v1}
                            _ = v14
                        } else {
                            v15 := GenS1{F0: 14, F1: 4294967295}
                            v16 := GenS2{F0: v1}
                            _ = v15
                            _ = v16
                        }
                        v17 := GenS3{F0: 1, F1: fi8}
                        _ = v13
                        _ = v17
                    }
                }
                v2 = true
            }
            _ = fi9
        }
        _ = fi8
    }
    v18 := GenS2{F0: (v1 ^ 1)}
    var v19 uint16 = 166
    v7 = (v7 * v7)
    v20 := []int{32767}
    if (255 == (^(v7))) {
        fmt.Println(v19, v19)
        var v21 uint64 = 927
        for fi22 := 0; fi22 < 2; fi22++ {
            for fi23 := 0; fi23 < 11; fi23++ {
                var v24 uint32 = (v3 | v3)
                switch (("foo bar" + "a") + "") {
                case "foo bar":
                    v7 = (v7 ^ 18446744073709551615)
                case "a":
                    fmt.Println(v2, v21)
                case "\n\t":
                    if v2 {
                        switch ("" + "hello") {
                        case "\n\t":
                            switch 9223372036854775807 {
                            case 4294967295:
                                switch ("" + "\n\t") {
                                case "foo bar":
                                    v25 := v5[(v24 * (v3 / ((v3) | 1)))]
                                    _ = v25
                                case "hello":
                                    v4 = 1.5
                                case "":
                                    var v26 uint16 = 1
                                    _ = v26
                                }
                            case 32767:
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                            case -33:
                                v27 := GenS2{F0: v1}
                                _ = v27
                            default:
                                switch "foo bar" {
                                case "a":
                                    if v2 {
                                        v28 := v5[(^(v24))]
                                        _ = v28
                                    }
                                    v19 -= v19
                                    switch (fi23 | fi22) {
                                    case 32767:
                                        if (v2 || ("\n\t" != "hello")) {
                                            v19 += v19
                                            fmt.Println(fi22)
                                            var v29 int = fi23
                                            _ = v29
                                        }
                                    case 1:
                                        fmt.Println(v24, v2, v7)
                                    case -9223372036854775808:
                                        if true {
                                            if true {
                                                delete(v5, (v3 - 0))
                                                v7 -= 18446744073709551615
                                                fmt.Println(v7, v24)
                                                var v30 int64 = -9223372036854775808
                                                if len(v20) >= 2 {
                                                    v20 = v20[1:len(v20)]
                                                }
                                                _ = v30
                                            } else {
                                                v31 := v5[v24]
                                                var v32 bool = (!true)
                                                _ = v31
                                                _ = v32
                                            }
                                        }
                                    }
                                    fmt.Println(v19, v19)
                                    v33 := GenS2{F0: 198}
                                    _ = v33
                                }
                            }
                        case "hello":
                            var v34 string = "a"
                            _ = v34
                        case "a":
                            v19 += v19
                        }
                        var v35 bool = (0.0 >= v4)
                        _ = v35
                    } else {
                        switch "foo bar" {
                        case "":
                            v1 += v1
                        case "a":
                            var v36 uint8 = 2
                            _ = v36
                        case "hello":
                            var v37 int16 = 0
                            _ = v37
                        }
                    }
                default:
                    if true {
                        v3 -= (v24 * 2147483648)
                        v19 = v19
                        var v38 uint32 = (v24 / ((v3) | 1))
                        _ = v38
                    } else {
                        var v39 int8 = 9
                        _ = v39
                    }
                    var v40 int8 = 6
                    _ = v40
                }
                v24 = v24
                delete(v5, v3)
                _ = fi23
                _ = v24
            }
            v41 := GenS2{F0: v1}
            _ = fi22
            _ = v41
        }
        _ = v21
    }
    for fi42 := 0; fi42 < 20; fi42++ {
        if false {
            var v43 float64 = (1.5 - 0.0)
            _ = v43
        } else {
            for fi44 := 0; fi44 < 7; fi44++ {
                fmt.Println(v19)
                v3 = v3
                switch fi44 {
                case 2147483647:
                    var v45 uint = v7
                    _ = v45
                default:
                    switch "a" {
                    case "":
                        var v46 bool = v2
                        _ = v46
                    case "\n\t":
                        fmt.Println(v2, v7, v2)
                    }
                }
                _ = fi44
            }
            fmt.Println(v1, v7, fi42)
        }
        v47 := GenS4{F0: 42, F1: false, F2: ("\n\t" + "hello")}
        switch 180 {
        case 1:
            var v48 uint16 = v19
            _ = v48
        case -1:
            v49 := GenS5{F0: 375, F1: ("" + "")}
            _ = v49
        }
        _ = fi42
        _ = v47
    }
    fmt.Println(v7)
    var v50 float64 = (0.0 - float64(1.5))
    for fi51 := 0; fi51 < 19; fi51++ {
        for fi52 := 0; fi52 < 2; fi52++ {
            var v53 uint = v7
            if (v2 && (((0.1 + v50) * v50) >= v50)) {
                v50 = (float64(v4) / ((float64(((float32(v50) - 1e10) + v4)) + v50) * ((((1e10 - 0.0) - v50) + 0.0) - v50)))
            } else {
                switch 1 {
                case -355:
                    if (142 <= -1) {
                        if (127 > 100) {
                            switch (fi51 * (fi51 / ((-2147483648) | 1))) {
                            case 9223372036854775807:
                                v1 += (v1 & v1)
                            case -1:
                                var v54 bool = v2
                                _ = v54
                            default:
                                fmt.Println(v53, v19)
                            }
                            if (9223372036854775807 != -901) {
                                v55 := GenS6{F0: -535, F1: uint16(v1), F2: "a"}
                                _ = v55
                            }
                            v53 += (v7 << ((uint(v3)) & 63))
                            fmt.Println(v4)
                            fmt.Println(fi52)
                            switch fi51 {
                            case 128:
                                fmt.Println(fi51, v53, v50)
                            case 32767:
                                if (!v2) {
                                    switch (fi52 | fi52) {
                                    case -133:
                                        var v56 uint = 369
                                        _ = v56
                                    case 32768:
                                        fmt.Println(fi52, v4, v4)
                                    case -9223372036854775808:
                                        v57 := GenS7{F0: (^(v1))}
                                        _ = v57
                                    }
                                    if (v2 && (!v2)) {
                                        fmt.Println(v53, fi52)
                                        v58 := GenS8{F0: 0}
                                        if (true || (!v2)) {
                                            var v59 int32 = -724
                                            if ((v7 - v53) == v7) {
                                                fmt.Println(v4, v3)
                                            }
                                            v60 := GenS9{F0: false}
                                            _ = v59
                                            _ = v60
                                        } else {
                                            switch fi52 {
                                            case 94:
                                                if v2 {
                                                    var v61 bool = v2
                                                    if ((!true) || ((-32768 >= -32768) || (!(v19 < v19)))) {
                                                        if (!("\n\t" == "\n\t")) {
                                                            v53 += (^(v53))
                                                        } else {
                                                            if v61 {
                                                                if (false || v61) {
                                                                    v4 += v4
                                                                    _ = len(v5)
                                                                    v62 := GenS10{F0: (29 >= 2), F1: ((v19 % (((^(v19))) | 1)) < 674), F2: 1}
                                                                    _ = v62
                                                                }
                                                                delete(v5, v3)
                                                                if (fi52 <= -1) {
                                                                    v2 = v2
                                                                }
                                                                fmt.Println(v50, v61, v2)
                                                            }
                                                        }
                                                    } else {
                                                        v20 = append(v20, (-(fi51)))
                                                        var v63 int8 = 100
                                                        _ = v63
                                                    }
                                                    var v64 string = "a"
                                                    var v65 uint64 = 0
                                                    _ = v61
                                                    _ = v64
                                                    _ = v65
                                                } else {
                                                    var v66 float64 = v50
                                                    v4 -= (-2.25 - float32((v50 + (0.1 + (0.0 + 0.1)))))
                                                    switch "a" {
                                                    case "foo bar":
                                                        _ = len(v20)
                                                    case "hello":
                                                        if (v2 || v2) {
                                                            if true {
                                                                v50 += 1e10
                                                                v67 := v5[(v3 ^ 0)]
                                                                _ = v67
                                                            }
                                                            _ = len(v5)
                                                            var v68 uint16 = (v19 / ((v19) | 1))
                                                            _ = v68
                                                        } else {
                                                            switch (^(fi51)) {
                                                            case -9223372036854775808:
                                                                v69 := GenS11{F0: (fi52 | 378), F1: 32767, F2: (v3 << (((v7 ^ v53)) & 63))}
                                                                _ = v69
                                                            }
                                                        }
                                                    case "\n\t":
                                                        v70 := GenS12{F0: v3}
                                                        _ = v70
                                                    }
                                                    _ = v66
                                                }
                                            case 128:
                                                fmt.Println(fi52, v19)
                                                var v71 float64 = v50
                                                _ = v71
                                            case 32767:
                                                if v2 {
                                                    if false {
                                                        delete(v5, v3)
                                                    } else {
                                                        switch (fi51 | fi52) {
                                                        case 4294967295:
                                                            v72 := GenS1{F0: 42, F1: (v3 ^ v3)}
                                                            _ = v72
                                                        case 32768:
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                v73 := GenS13{F0: uint32(v19)}
                                                                v74 := GenS13{F0: (^(v3))}
                                                                var v75 int = -658
                                                                _ = v73
                                                                _ = v74
                                                                _ = v75
                                                            }
                                                        }
                                                    }
                                                    v76 := GenS13{F0: v3}
                                                    _ = v76
                                                } else {
                                                    if false {
                                                        v77 := GenS6{F0: -1, F1: (v19 / ((v19) | 1)), F2: "a"}
                                                        _ = v77
                                                    }
                                                }
                                            }
                                            v53 += v53
                                            v7 += v7
                                        }
                                        _ = v58
                                    } else {
                                        if len(v20) >= 2 {
                                            v20 = v20[1:len(v20)]
                                        }
                                        switch "\n\t" {
                                        case "hello":
                                            _ = len(v5)
                                        case "foo bar":
                                            v7 = (v7 - (v53 << ((v7) & 63)))
                                        case "a":
                                            if (!false) {
                                                _ = len(v20)
                                                v53 = v53
                                                v4 -= (((v4 / 3.14159) - (v4 / -2.25)) - (v4 + (3.14159 - float32(v50))))
                                            }
                                        }
                                        if v2 {
                                            fmt.Println(v3, fi52, fi52)
                                            v50 += (0.0 / v50)
                                        }
                                    }
                                    fmt.Println(v50)
                                }
                                var v78 int32 = 1
                                if (!(uint8(v3) > 1)) {
                                    switch (("\n\t" + "\n\t") + "") {
                                    case "\n\t":
                                        v79 := GenS1{F0: 61, F1: (v3 | (v3 ^ v3))}
                                        _ = v79
                                    default:
                                        switch "\n\t" {
                                        case "a":
                                            switch "\n\t" {
                                            case "hello":
                                                switch 612 {
                                                case -428:
                                                    v80 := GenS3{F0: v53, F1: 42}
                                                    _ = v80
                                                case 0:
                                                    switch ("hello" + "foo bar") {
                                                    case "a":
                                                        v81 := v5[4294967295]
                                                        _ = v81
                                                    case "":
                                                        v50 = v50
                                                    case "\n\t":
                                                        var v82 int = fi52
                                                        _ = v82
                                                    default:
                                                        v5[uint32(v19)] = "\n\t"
                                                    }
                                                case 128:
                                                    if len(v20) > 0 {
                                                        v20[0] = fi51
                                                    }
                                                }
                                            default:
                                                fmt.Println(v3)
                                            }
                                        default:
                                            var v83 uint16 = v19
                                            _ = v83
                                        }
                                    }
                                    var v84 uint16 = 32767
                                    _ = v84
                                }
                                _ = v78
                            }
                            var v85 float64 = v50
                            _ = v85
                        }
                        v20 = append(v20, (fi51 | (fi51 << ((v7) & 63))))
                    }
                case 127:
                    var v86 float64 = 0.0
                    _ = v86
                case -2147483648:
                    if false {
                        fmt.Println(v3)
                        var v87 int64 = 508
                        switch (^(fi52)) {
                        case -452:
                            var v88 int8 = 127
                            _ = v88
                        }
                        _ = v87
                    } else {
                        fmt.Println(v50, v2, v19)
                        v89 := GenS1{F0: 0, F1: v3}
                        v90 := v5[v3]
                        _ = v89
                        _ = v90
                    }
                }
                var v91 bool = (100 < -128)
                if v91 {
                    v20 = append(v20, fi52)
                    v92 := GenS2{F0: 1}
                    v93 := GenS14{F0: 1, F1: ""}
                    _ = v92
                    _ = v93
                }
                _ = v91
            }
            fmt.Println(fi51, v7)
            _ = fi52
            _ = v53
        }
        fmt.Println(v3, v4, v1)
        for fi94 := 0; fi94 < 3; fi94++ {
            v95 := GenS13{F0: v3}
            v96 := GenS15{F0: (v3 % ((uint32(v19)) | 1))}
            var v97 uint16 = v19
            _ = fi94
            _ = v95
            _ = v96
            _ = v97
        }
        _ = fi51
    }
    v4 -= (0.0 * v4)
    var v98 string = ""
    v7 = 42
    switch 128 {
    case -1:
        v98 = v98
    default:
        var v99 int64 = 733
        _ = v99
    }
    v100 := GenS16{F0: 4294967295, F1: "foo bar", F2: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v7:", v7)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v100
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v7
    _ = v98
}
