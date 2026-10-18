package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int8; F1 int16; F2 uint }
type GenS5 struct { F0 int; F1 uint16; F2 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 int8; F2 int64 }
type GenS8 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int64 }
func genFn0(p0 int16, p1 int8) uint {
    switch 42 {
    case -906:
        if false {
            v5 := GenS1{F0: false, F1: 255, F2: 719}
            if (true && ("hello" >= "\n\t")) {
                var v6 float32 = -2.25
                _ = v6
            }
            _ = v5
        } else {
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 string = ""
                v8 = v8
                _ = fi7
                _ = v8
            }
            p1 -= p1
            var v9 int8 = p1
            _ = v9
        }
    case 1:
        if true {
            var v10 bool = true
            _ = v10
        }
    case -696:
        v11 := []int{-2147483648, 32767, 32767, 9223372036854775807}
        _ = v11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        p0 += (-(p0))
        _ = fi12
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 uint32 = 1
    v2 := []int{32768, 9223372036854775807, 688, 32767}
    var v3 uint32 = v1
    var v4 bool = (-749 != 205)
    v13 := GenS1{F0: true, F1: 255, F2: 2147483647}
    v4 = v4
    v4 = v4
    var v14 bool = v4
    v14 = v14
    if false {
        for fi15 := 0; fi15 < 8; fi15++ {
            v1 = 502
            for fi16 := 0; fi16 < 12; fi16++ {
                if len(v2) > 0 {
                    v2[0] = (fi16 ^ (fi15 / (((fi16 * fi16)) | 1)))
                }
                if ((127 != 158) && (!(v14 || (626 <= 359)))) {
                    v4 = v4
                }
                v17 := GenS2{F0: 42, F1: 126}
                _ = fi16
                _ = v17
            }
            if (!("" != "hello")) {
                v18 := GenS1{F0: (!(-605 <= -246)), F1: 17, F2: 2147483647}
                switch "hello" {
                case "foo bar":
                    switch "" {
                    case "\n\t":
                        var v19 uint = genFn0(-522, -8)
                        _ = v19
                    case "hello":
                        v20 := GenS1{F0: (65535 > 1), F1: 1, F2: -2147483648}
                        _ = v20
                    case "foo bar":
                        if (v14 && (4294967295 >= 1)) {
                            v1 -= v1
                        } else {
                            switch (^(fi15)) {
                            case 42:
                                var v21 bool = true
                                _ = v21
                            case -275:
                                v2 = append(v2, (fi15 & fi15))
                            case 0:
                                if (v14 && true) {
                                    v1 = (v3 ^ v1)
                                    if (1 >= 1) {
                                        for fi22 := 0; fi22 < 5; fi22++ {
                                            if len(v2) > 0 {
                                                v2[0] = (fi22 * (fi22 | fi22))
                                            }
                                            v23 := GenS2{F0: 529, F1: 255}
                                            _ = fi22
                                            _ = v23
                                        }
                                        var v24 int = 128
                                        _ = v24
                                    }
                                }
                            default:
                                var v25 float64 = (float64((1e10 + 1e10)) / float64(float32(1e10)))
                                _ = v25
                            }
                            var v26 uint = 42
                            _ = v26
                        }
                    }
                case "\n\t":
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v4)
                    case "\n\t":
                        var v27 bool = v14
                        _ = v27
                    }
                case "":
                    var v28 int16 = 0
                    _ = v28
                default:
                    var v29 bool = (0 >= 1)
                    _ = v29
                }
                _ = v18
            }
            _ = fi15
        }
        v3 += (v3 & (v1 & 65535))
    } else {
        var v30 float64 = ((1e10 / 0.1) * 1e10)
        fmt.Println(v30, v30)
        var v31 float32 = (0.0 - 0.0)
        _ = v30
        _ = v31
    }
    var v32 float64 = float64(0.0)
    switch 0 {
    case 1:
        _ = len(v2)
    case 32768:
        for fi33 := 0; fi33 < 19; fi33++ {
            if v14 {
                if v4 {
                    fmt.Println(v4, v32, v32)
                }
                var v34 int32 = 0
                v34 = v34
                _ = v34
            }
            _ = fi33
        }
    default:
        v35 := GenS3{F0: "\n\t", F1: 1}
        _ = v35
    }
    if v4 {
        if v14 {
            for fi36 := 0; fi36 < 17; fi36++ {
                if false {
                    fmt.Println(fi36, v1)
                    v1 -= (v1 & v1)
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v37 := GenS3{F0: ("foo bar" + ("\n\t" + "")), F1: 42}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v37
                }
                if true {
                    for fi38 := 0; fi38 < 8; fi38++ {
                        v39 := GenS4{F0: -74, F1: 1, F2: genFn0(42, 127)}
                        v3 -= v1
                        _ = fi38
                        _ = v39
                    }
                    var v40 uint32 = v1
                    _ = v40
                }
                _ = fi36
            }
            for fi41 := 0; fi41 < 3; fi41++ {
                if (v14 || v4) {
                    v3 += (v3 >> ((2) & 63))
                    v42 := GenS2{F0: 12, F1: 4}
                    _ = v42
                } else {
                    if v4 {
                        if (785 > -1) {
                            var v43 bool = v4
                            if v14 {
                                if false {
                                    if (!true) {
                                        v14 = true
                                        if v14 {
                                            v32 = (0.0 - (((v32 + v32) + float64(0.0)) - (v32 + 0.0)))
                                        } else {
                                            var v44 uint8 = 42
                                            v1 = (v1 ^ 978)
                                            _ = v44
                                        }
                                    } else {
                                        switch "a" {
                                        case "":
                                            if v43 {
                                                v45 := GenS5{F0: 1, F1: 916, F2: -93}
                                                for fi46 := 0; fi46 < 12; fi46++ {
                                                    switch "hello" {
                                                    case "a":
                                                        var v47 bool = ((-128 < 13) && ((1 < 338) || (uint64(v3) != 42)))
                                                        _ = v47
                                                    case "hello":
                                                        switch (fi41 & (fi46 >> ((1) & 63))) {
                                                        case 1:
                                                            if (65535 < uint16(v1)) {
                                                                var v48 bool = v4
                                                                v49 := GenS1{F0: v48, F1: 1, F2: (fi41 % (((fi46 & fi41)) | 1))}
                                                                fmt.Println(v48, v4, v43)
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                        default:
                                                            var v50 int16 = -32768
                                                            if ((false || (v50 >= -823)) || v14) {
                                                                _ = len(v2)
                                                            }
                                                            if (fi46 >= fi46) {
                                                                var v51 int = (fi46 & 662)
                                                                _ = v51
                                                            } else {
                                                                var v52 int64 = -9223372036854775808
                                                                v4 = v14
                                                                _ = v52
                                                            }
                                                            _ = v50
                                                        }
                                                    case "\n\t":
                                                        if len(v2) > 0 {
                                                            v2[0] = fi41
                                                        }
                                                        switch -839 {
                                                        case 2147483647:
                                                            v4 = (!v43)
                                                        case 32768:
                                                            v53 := GenS2{F0: -2147483648, F1: 2}
                                                            _ = v53
                                                        case -9223372036854775808:
                                                            var v54 uint64 = 18446744073709551615
                                                            _ = v54
                                                        }
                                                        var v55 uint8 = 128
                                                        _ = v55
                                                    }
                                                    switch "hello" {
                                                    case "hello":
                                                        fmt.Println(fi41, v32, fi41)
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = fi46
                                                }
                                                _ = v45
                                            } else {
                                                fmt.Println(v4, v14)
                                            }
                                        }
                                    }
                                } else {
                                    var v56 uint16 = 42
                                    if (1 != -46) {
                                        if v4 {
                                            v57 := GenS2{F0: -2147483648, F1: 127}
                                            _ = v57
                                        } else {
                                            switch 474 {
                                            case -9223372036854775808:
                                                var v58 int = -714
                                                _ = v58
                                            default:
                                                var v59 bool = (1e10 == v32)
                                                _ = v59
                                            }
                                        }
                                    }
                                    _ = v56
                                }
                                switch "hello" {
                                case "a":
                                    v60 := GenS6{F0: v3}
                                    _ = v60
                                case "":
                                    for fi61 := 0; fi61 < 17; fi61++ {
                                        var v62 uint8 = 2
                                        _ = fi61
                                        _ = v62
                                    }
                                case "\n\t":
                                    v63 := GenS7{F0: 2147483647, F1: 100, F2: -1}
                                    _ = v63
                                }
                                if v4 {
                                    v32 += (v32 / float64(3.14159))
                                    if len(v2) > 0 {
                                        v2[0] = fi41
                                    }
                                } else {
                                    var v64 int16 = 42
                                    _ = v64
                                }
                            } else {
                                v32 = float64((0.0 + 0.1))
                                fmt.Println(v1, v32)
                                var v65 uint = 4294967295
                                _ = v65
                            }
                            v3 = 0
                            _ = v43
                        } else {
                            fmt.Println(fi41, v3, v14)
                        }
                        var v66 float64 = 0.0
                        if (true || true) {
                            v66 = v32
                        } else {
                            for fi67 := 0; fi67 < 3; fi67++ {
                                if v4 {
                                    fmt.Println(v66, v66, v66)
                                    if true {
                                        v4 = (v4 && (127 == 0))
                                        fmt.Println(v4, v3)
                                        fmt.Println(v14, v14)
                                    } else {
                                        v68 := GenS2{F0: 698, F1: 2}
                                        if v14 {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v4 = v4
                                        } else {
                                            _ = len(v2)
                                            var v69 uint16 = 4
                                            v2 = append(v2, fi41)
                                            _ = v69
                                        }
                                        var v70 float32 = 0.0
                                        _ = v68
                                        _ = v70
                                    }
                                } else {
                                    v14 = false
                                    switch "hello" {
                                    case "foo bar":
                                        v32 += v66
                                    default:
                                        var v71 uint16 = 2
                                        _ = v71
                                    }
                                }
                                if v14 {
                                    var v72 int64 = -895
                                    switch (fi41 % (((fi67 & 2147483647)) | 1)) {
                                    case -980:
                                        v73 := GenS4{F0: 100, F1: 1, F2: 4294967295}
                                        _ = v73
                                    case -871:
                                        v66 += 0.0
                                    case 128:
                                        if (-2.25 > (((1.5 + (float32(0.0) / 1e10)) / (float32(float64(3.14159)) / ((1.5 - 0.1) * 1.5))) * (0.1 + 1.5))) {
                                            v66 -= (((v32 + v32) - 0.1) * v32)
                                        } else {
                                            fmt.Println(v1, v3, v72)
                                            var v74 int = 128
                                            v75 := GenS8{F0: 127, F1: 462, F2: 77}
                                            _ = v74
                                            _ = v75
                                        }
                                    }
                                    _ = v72
                                } else {
                                    if v14 {
                                        if ((v3 + 2147483647) == (v3 + v1)) {
                                            var v76 int16 = 42
                                            _ = v76
                                        } else {
                                            switch "" {
                                            case "a":
                                                if v14 {
                                                    if len(v2) > 0 {
                                                        v2[0] = 850
                                                    }
                                                    var v77 int16 = -891
                                                    if (0.1 < float64(-2.25)) {
                                                        if (1 > 2) {
                                                            v14 = (!(!(4294967295 > fi67)))
                                                        } else {
                                                            fmt.Println(fi67, v66)
                                                        }
                                                        var v78 uint64 = 18446744073709551615
                                                        _ = v78
                                                    }
                                                    _ = v77
                                                }
                                            case "":
                                                var v79 float64 = -2.25
                                                _ = v79
                                            }
                                            var v80 int = (fi41 << ((65535) & 63))
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = v80
                                        }
                                        if ((42 == 0) || v14) {
                                            var v81 int = 16
                                            _ = v81
                                        }
                                        if (-1 <= int8(v1)) {
                                            fmt.Println(v14, v1)
                                            v14 = (-2147483648 <= 63)
                                            switch (fi67 >> ((255) & 63)) {
                                            case -910:
                                                v82 := GenS1{F0: v14, F1: 127, F2: fi41}
                                                _ = v82
                                            }
                                        } else {
                                            fmt.Println(v1, v1, v32)
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            switch fi41 {
                                            case -2147483648:
                                                var v83 int64 = 0
                                                _ = v83
                                            }
                                        }
                                    } else {
                                        v14 = (-9223372036854775808 <= -9223372036854775808)
                                    }
                                }
                                v84 := GenS3{F0: "foo bar", F1: 1000}
                                _ = fi67
                                _ = v84
                            }
                            fmt.Println(v32)
                            v3 += v3
                        }
                        _ = v66
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v3 -= (v1 % (((v3 | v1)) | 1))
                        var v85 int = (fi41 >> ((65535) & 63))
                        _ = v85
                    }
                    fmt.Println(v4, v32)
                    v3 = v3
                }
                var v86 string = "a"
                v87 := GenS9{F0: 230}
                _ = fi41
                _ = v86
                _ = v87
            }
            if v14 {
                var v88 float32 = 0.0
                var v89 uint8 = 45
                for fi90 := 0; fi90 < 10; fi90++ {
                    if false {
                        var v91 uint64 = 143
                        _ = v91
                    }
                    v89 = v89
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi90
                }
                _ = v88
                _ = v89
            } else {
                for fi92 := 0; fi92 < 5; fi92++ {
                    var v93 int16 = 32767
                    _ = fi92
                    _ = v93
                }
            }
        }
        switch ("hello" + ("hello" + ("\n\t" + "a"))) {
        case "":
            if (-1 != 9223372036854775807) {
                v94 := GenS10{F0: genFn0(1, 48), F1: 430, F2: 65535}
                var v95 float64 = 0.0
                switch "\n\t" {
                case "\n\t":
                    if false {
                        if len(v2) > 0 {
                            v2[0] = 42
                        }
                        switch 2147483647 {
                        case 42:
                            fmt.Println(v95, v95, v32)
                        case -2147483648:
                            v4 = true
                        case 2147483647:
                            v2 = append(v2, -663)
                        default:
                            var v96 int32 = -2147483648
                            _ = v96
                        }
                        var v97 string = "foo bar"
                        _ = v97
                    } else {
                        v4 = false
                        var v98 uint8 = 42
                        _ = v98
                    }
                case "hello":
                    v99 := GenS4{F0: -128, F1: -1, F2: 751}
                    _ = v99
                default:
                    switch "a" {
                    case "\n\t":
                        v95 = ((((0.1 / (0.1 + (float64(0.0) * (v32 / (v95 / v95))))) / (v95 + v95)) * v95) - (((v32 + v95) / ((-2.25 - 0.1) + 1e10)) / 1e10))
                    }
                }
                _ = v94
                _ = v95
            } else {
                fmt.Println(v4, v4, v3)
                var v100 uint32 = (v3 & v3)
                _ = v100
            }
        case "hello":
            var v101 bool = (128 == 2147483647)
            _ = v101
        }
        v102 := GenS11{F0: 2147483647}
        _ = v102
    }
    var v103 float32 = 1e10
    fmt.Println(v14, v32)
    v103 = v103
    v104 := GenS12{F0: 9223372036854775807}
    fmt.Println(v3, v1, v14)
    switch "a" {
    case "a":
        v105 := GenS7{F0: -1, F1: 42, F2: 197}
        _ = v105
    case "":
        var v106 int = 141
        _ = v106
    }
    switch "a" {
    case "hello":
        var v107 int32 = 1
        _ = v107
    case "":
        var v108 uint = 18446744073709551615
        _ = v108
    }
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v103
    _ = v104
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
