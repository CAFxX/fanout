package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint8; F1 int; F2 int16 }
type GenS5 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 string; F1 int64; F2 uint32 }

func main() {
    var v1 uint = 2
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    v4 := GenS1{F0: 1, F1: 127}
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            if (true && false) {
                fmt.Println(v1, fi5)
                v7 := GenS1{F0: 2, F1: -128}
                _ = v7
            }
            if true {
                var v8 int64 = 9223372036854775807
                _ = v8
            } else {
                switch ("" + "foo bar") {
                case "hello":
                    var v9 uint = v1
                    _ = v9
                default:
                    v10 := GenS1{F0: 4294967295, F1: 42}
                    _ = v10
                }
                v11 := GenS1{F0: v1, F1: 82}
                _ = v11
            }
            _ = fi6
        }
        _ = fi5
    }
    var v12 string = "\n\t"
    switch 127 {
    case -864:
        if (42 > -1) {
            var v13 int16 = 1
            v13 -= 42
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 int64 = 0
                for fi16 := 0; fi16 < 2; fi16++ {
                    v17 := []int{}
                    _ = fi16
                    _ = v17
                }
                _ = fi14
                _ = v15
            }
            _ = v13
        }
    case -1:
        var v18 bool = (993 < -304)
        _ = v18
    }
    v12 = v12
    fmt.Println(v12, v1)
    v19 := []int{}
    v3[v12] = 42
    if true {
        if false {
            v20 := GenS1{F0: v1, F1: 119}
            if (!(-637 <= 42)) {
                fmt.Println(v1)
                if (409 == 128) {
                    fmt.Println(v1, v12)
                    _ = len(v3)
                    var v21 uint64 = 725
                    _ = v21
                } else {
                    if false {
                        fmt.Println(v1)
                        switch -1 {
                        case 127:
                            for fi22 := 0; fi22 < 8; fi22++ {
                                fmt.Println(v12, fi22)
                                _ = fi22
                            }
                        default:
                            for fi23 := 0; fi23 < 16; fi23++ {
                                fmt.Println(fi23)
                                switch fi23 {
                                case -321:
                                    var v24 int8 = 0
                                    _ = v24
                                case -286:
                                    var v25 uint64 = 2
                                    _ = v25
                                case 1:
                                    var v26 uint8 = 127
                                    _ = v26
                                default:
                                    switch (-(fi23)) {
                                    case -856:
                                        var v27 int64 = -1
                                        _ = v27
                                    case 127:
                                        for fi28 := 0; fi28 < 11; fi28++ {
                                            fmt.Println(v12, v1, v12)
                                            var v29 uint = (^(v1))
                                            _ = fi28
                                            _ = v29
                                        }
                                    case 2147483647:
                                        if true {
                                            if false {
                                                if true {
                                                    v30 := v3[((v12 + (("" + ("hello" + v12)) + "")) + v12)]
                                                    v12 = v12
                                                    v12 = ("hello" + (v12 + v12))
                                                    _ = v30
                                                }
                                                switch 856 {
                                                case 32767:
                                                    if (!false) {
                                                        var v31 uint32 = 828
                                                        fmt.Println(fi23, fi23, fi23)
                                                        v12 = ("hello" + ((v12 + (v12 + "")) + v12))
                                                        _ = v31
                                                    }
                                                case 268:
                                                    if (-128 <= 100) {
                                                        if true {
                                                            v32 := GenS2{F0: (fi23 / ((32768) | 1))}
                                                            _ = v32
                                                        } else {
                                                            if (fi23 < fi23) {
                                                                v12 = v12
                                                                fmt.Println(v12)
                                                                var v33 int64 = 123
                                                                _ = v33
                                                            } else {
                                                                v19 = append(v19, 705)
                                                                var v34 uint64 = 2
                                                                _ = v34
                                                            }
                                                        }
                                                        switch fi23 {
                                                        case 32768:
                                                            v12 = (v12 + v12)
                                                        case 568:
                                                            for fi35 := 0; fi35 < 6; fi35++ {
                                                                if (32767 > 42) {
                                                                    v12 = v12
                                                                    if (201 != 1) {
                                                                        if true {
                                                                            _ = len(v19)
                                                                            var v36 int32 = 42
                                                                            _ = v36
                                                                        } else {
                                                                            var v37 uint64 = 4294967295
                                                                            var v38 float64 = -2.25
                                                                            v12 = ((v12 + (v12 + v12)) + "hello")
                                                                            fmt.Println(fi23, v1, v38)
                                                                            var v39 int16 = 630
                                                                            _ = v37
                                                                            _ = v38
                                                                            _ = v39
                                                                        }
                                                                    }
                                                                    v1 = 4294967295
                                                                } else {
                                                                    _ = len(v3)
                                                                    if true {
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            v40 := GenS1{F0: v1, F1: 101}
                                                                            _ = v40
                                                                        default:
                                                                            v12 = v12
                                                                        }
                                                                    }
                                                                    if (((fi35 & 1) >= (fi23 & 127)) || true) {
                                                                        switch -1 {
                                                                        case 952:
                                                                            v12 = (v12 + v12)
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                var v41 uint8 = 42
                                                                                _ = v41
                                                                            }
                                                                            v42 := GenS2{F0: fi35}
                                                                            _ = len(v19)
                                                                            if true {
                                                                                v12 = (v12 + v12)
                                                                                switch fi23 {
                                                                                case 52:
                                                                                    fmt.Println(v1, fi23)
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v1, v12, v1)
                                                                            }
                                                                            if true {
                                                                                v2[(fi35 >> ((0) & 63))] = 803
                                                                                v43 := GenS2{F0: -983}
                                                                                v44 := GenS1{F0: (v1 >> ((1) & 63)), F1: -37}
                                                                                v1 = (^(v1))
                                                                                var v45 int16 = 1
                                                                                _ = v43
                                                                                _ = v44
                                                                                _ = v45
                                                                            }
                                                                            switch fi23 {
                                                                            case 128:
                                                                                if false {
                                                                                    fmt.Println(v12, v1, fi23)
                                                                                    fmt.Println(fi35, fi23)
                                                                                    fmt.Println(fi35)
                                                                                } else {
                                                                                    fmt.Println(v12)
                                                                                    v12 = ((v12 + (("foo bar" + v12) + (v12 + v12))) + v12)
                                                                                    v1 -= 0
                                                                                }
                                                                            case -431:
                                                                                switch ("\n\t" + "hello") {
                                                                                case "\n\t":
                                                                                    fmt.Println(v12)
                                                                                case "":
                                                                                    v1 = (v1 << (((v1 * 65535)) & 63))
                                                                                default:
                                                                                    switch v12 {
                                                                                    case "\n\t":
                                                                                        var v46 int64 = 42
                                                                                        _ = v46
                                                                                    case "":
                                                                                        if false {
                                                                                            if false {
                                                                                                switch (-(fi35)) {
                                                                                                case 128:
                                                                                                    v12 = (v12 + (v12 + v12))
                                                                                                    v47 := GenS1{F0: v1, F1: 1}
                                                                                                    v1 -= v1
                                                                                                    _ = v47
                                                                                                default:
                                                                                                    if false {
                                                                                                        v12 = (v12 + v12)
                                                                                                        var v48 int = 128
                                                                                                        _ = v48
                                                                                                    } else {
                                                                                                        if (false || false) {
                                                                                                            v49 := v3["hello"]
                                                                                                            v1 -= (v1 + v1)
                                                                                                            v50 := GenS3{F0: (v49 % (((v49 - 404)) | 1)), F1: 42}
                                                                                                            _ = v49
                                                                                                            _ = v50
                                                                                                        }
                                                                                                        fmt.Println(fi35, fi23, v1)
                                                                                                        var v51 int64 = -860
                                                                                                        _ = v51
                                                                                                    }
                                                                                                }
                                                                                                switch "hello" {
                                                                                                case "a":
                                                                                                    var v52 uint32 = 464
                                                                                                    _ = v52
                                                                                                default:
                                                                                                    delete(v2, 599)
                                                                                                }
                                                                                            } else {
                                                                                                v12 = "foo bar"
                                                                                                if true {
                                                                                                    v12 = v12
                                                                                                    if true {
                                                                                                        fmt.Println(fi35)
                                                                                                    }
                                                                                                    v53 := GenS4{F0: 0, F1: (fi35 * fi35), F2: -87}
                                                                                                    _ = v53
                                                                                                } else {
                                                                                                    v12 = "\n\t"
                                                                                                    var v54 int = 305
                                                                                                    switch (("a" + v12) + (("\n\t" + "foo bar") + (v12 + v12))) {
                                                                                                    case "a":
                                                                                                        var v55 int32 = 726
                                                                                                        _ = v55
                                                                                                    case "foo bar":
                                                                                                        var v56 int32 = -1
                                                                                                        _ = v56
                                                                                                    }
                                                                                                    _ = v54
                                                                                                }
                                                                                                fmt.Println(fi23, fi23, fi23)
                                                                                                v12 = v12
                                                                                            }
                                                                                            switch fi35 {
                                                                                            case -1:
                                                                                                v1 -= v1
                                                                                            case 1:
                                                                                                var v57 int32 = 317
                                                                                                _ = v57
                                                                                            case -298:
                                                                                                if (42 >= 0) {
                                                                                                    switch 9223372036854775807 {
                                                                                                    case 9223372036854775807:
                                                                                                        switch v12 {
                                                                                                        case "":
                                                                                                            v58 := GenS5{F0: (-1 <= 42), F1: 42, F2: 2}
                                                                                                            _ = v58
                                                                                                        case "foo bar":
                                                                                                            fmt.Println(fi35)
                                                                                                            fmt.Println(v1)
                                                                                                        default:
                                                                                                            delete(v3, (v12 + (v12 + (((((v12 + v12) + v12) + "\n\t") + v12) + ((v12 + "a") + "hello")))))
                                                                                                        }
                                                                                                    case -705:
                                                                                                        var v59 bool = (!false)
                                                                                                        _ = v59
                                                                                                    default:
                                                                                                        var v60 int32 = 2147483647
                                                                                                        _ = v60
                                                                                                    }
                                                                                                    switch v12 {
                                                                                                    case "a":
                                                                                                        var v61 int64 = -1
                                                                                                        _ = v61
                                                                                                    case "foo bar":
                                                                                                        if false {
                                                                                                            switch "" {
                                                                                                            case "foo bar":
                                                                                                                v1 = v1
                                                                                                            case "hello":
                                                                                                                var v62 uint64 = 18446744073709551615
                                                                                                                if false {
                                                                                                                    _ = len(v19)
                                                                                                                    var v63 int64 = -9223372036854775808
                                                                                                                    v12 = "foo bar"
                                                                                                                    _ = v63
                                                                                                                }
                                                                                                                _ = v62
                                                                                                            default:
                                                                                                                fmt.Println(fi23, v1, fi23)
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v64 int = (fi35 << ((65535) & 63))
                                                                                                            fmt.Println(v12)
                                                                                                            if len(v19) >= 2 {
                                                                                                                v19 = v19[1:len(v19)]
                                                                                                            }
                                                                                                            switch v64 {
                                                                                                            case 0:
                                                                                                                fmt.Println(fi23)
                                                                                                            }
                                                                                                            _ = v64
                                                                                                        }
                                                                                                    }
                                                                                                    v19 = append(v19, (^(fi23)))
                                                                                                }
                                                                                            default:
                                                                                                v65 := GenS6{F0: 785}
                                                                                                _ = v65
                                                                                            }
                                                                                        } else {
                                                                                            v1 += v1
                                                                                            v66 := GenS7{F0: 4294967295, F1: 42}
                                                                                            var v67 uint32 = 42
                                                                                            _ = v66
                                                                                            _ = v67
                                                                                        }
                                                                                    default:
                                                                                        var v68 float64 = (-2.25 + 3.14159)
                                                                                        _ = v68
                                                                                    }
                                                                                }
                                                                            case 9223372036854775807:
                                                                                v69 := v2[0]
                                                                                _ = v69
                                                                            default:
                                                                                switch fi23 {
                                                                                case -804:
                                                                                    v70 := GenS2{F0: (fi23 * (fi35 / ((fi23) | 1)))}
                                                                                    _ = v70
                                                                                case -2147483648:
                                                                                    v71 := GenS2{F0: fi23}
                                                                                    _ = v71
                                                                                case 30:
                                                                                    v12 = v12
                                                                                default:
                                                                                    switch 42 {
                                                                                    case 128:
                                                                                        v12 = v12
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v42
                                                                        case 428:
                                                                            if ((((-2.25 / ((-2.25 + 1.5) + ((1e10 - 0.0) - ((1e10 - -2.25) * 1e10)))) / 1.5) * 3.14159) > (-2.25 + 3.14159)) {
                                                                                var v72 bool = true
                                                                                _ = len(v19)
                                                                                _ = v72
                                                                            } else {
                                                                                fmt.Println(fi35, v1, v12)
                                                                                var v73 uint64 = 4294967295
                                                                                v1 = 1
                                                                                if true {
                                                                                    var v74 int8 = 0
                                                                                    v1 -= v1
                                                                                    if (1e10 >= 1.5) {
                                                                                        var v75 int = 4294967295
                                                                                        _ = v75
                                                                                    }
                                                                                    if true {
                                                                                        var v76 uint64 = 18446744073709551615
                                                                                        if ((^(v1)) > 0) {
                                                                                            var v77 uint64 = (v73 << ((v1) & 63))
                                                                                            v78 := GenS7{F0: v77, F1: fi35}
                                                                                            _ = v77
                                                                                            _ = v78
                                                                                        } else {
                                                                                            var v79 uint8 = 191
                                                                                            var v80 float32 = (0.1 / 1.5)
                                                                                            _ = v79
                                                                                            _ = v80
                                                                                        }
                                                                                        var v81 bool = ((v73 + v76) == 42)
                                                                                        _ = v76
                                                                                        _ = v81
                                                                                    }
                                                                                    _ = v74
                                                                                } else {
                                                                                    fmt.Println(v1, v12, fi23)
                                                                                    if (0 > -193) {
                                                                                        fmt.Println(v12, v1, fi23)
                                                                                        v82 := GenS8{F0: v12, F1: 0, F2: 65535}
                                                                                        fmt.Println(v12)
                                                                                        _ = v82
                                                                                    }
                                                                                    v12 = v12
                                                                                }
                                                                                var v83 float64 = (1.5 - (-2.25 - -2.25))
                                                                                _ = v73
                                                                                _ = v83
                                                                            }
                                                                        case 1:
                                                                            var v84 uint64 = 0
                                                                            _ = v84
                                                                        }
                                                                        var v85 int8 = -96
                                                                        _ = v85
                                                                    }
                                                                }
                                                                _ = fi35
                                                            }
                                                        case -527:
                                                            var v86 uint16 = 65535
                                                            _ = v86
                                                        }
                                                        for fi87 := 0; fi87 < 4; fi87++ {
                                                            fmt.Println(fi23, v12)
                                                            _ = fi87
                                                        }
                                                    } else {
                                                        v1 -= (v1 & 535)
                                                        v88 := v2[fi23]
                                                        if true {
                                                            var v89 int = fi23
                                                            _ = v89
                                                        }
                                                        _ = v88
                                                    }
                                                case -1:
                                                    fmt.Println(fi23, v12, fi23)
                                                }
                                            } else {
                                                v90 := GenS6{F0: -2147483648}
                                                switch (fi23 / ((fi23) | 1)) {
                                                case 42:
                                                    var v91 int16 = 1000
                                                    _ = v91
                                                case -305:
                                                    for fi92 := 0; fi92 < 7; fi92++ {
                                                        if true {
                                                            v93 := GenS7{F0: 2, F1: (fi23 % ((fi92) | 1))}
                                                            _ = v93
                                                        }
                                                        v1 += v1
                                                        _ = fi92
                                                    }
                                                case -1:
                                                    if len(v19) > 0 {
                                                        v19[0] = -2147483648
                                                    }
                                                }
                                                _ = v90
                                            }
                                            v1 += 663
                                        }
                                    }
                                }
                                _ = fi23
                            }
                        }
                        var v94 int32 = -2147483648
                        _ = v94
                    }
                }
            } else {
                _ = len(v19)
            }
            switch ("foo bar" + "a") {
            case "":
                var v95 uint = (v1 - 18446744073709551615)
                _ = v95
            case "hello":
                v1 -= v1
            default:
                v12 = ""
            }
            _ = v20
        }
    } else {
        v1 += v1
    }
    fmt.Println(v12, v1)
    if len(v19) > 0 {
        v19[0] = 2147483647
    }
    var v96 float64 = float64(3.14159)
    var v97 uint32 = 65535
    fmt.Println(v1, v96, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v96
    _ = v97
}
