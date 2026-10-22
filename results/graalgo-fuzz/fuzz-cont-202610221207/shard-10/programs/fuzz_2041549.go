package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int; F1 uint32; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 int16; F1 int16; F2 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint; F1 int32 }
type GenS9 struct { F0 int }

func main() {
    v1 := []string{"a", "a", "foo bar", "foo bar"}
    v2 := []uint32{914, 0, 2147483647}
    v3 := make(map[uint32]string)
    var v4 float32 = -2.25
    v5 := make(map[int]int)
    switch 460 {
    case 127:
        switch ("foo bar" + "\n\t") {
        case "hello":
            v4 = (((((float32(1.5) * 0.1) + 1.5) + (v4 - 0.0)) + -2.25) - 0.1)
        case "a":
            if (-18 != -1) {
                v6 := GenS1{F0: 1}
                v4 = (1.5 / (v4 + 0.1))
                var v7 bool = false
                _ = v6
                _ = v7
            } else {
                for fi8 := 0; fi8 < 16; fi8++ {
                    var v9 int32 = 1
                    if true {
                        v10 := GenS1{F0: 2}
                        if false {
                            if true {
                                v9 += -1
                                switch (fi8 | (fi8 >> ((255) & 63))) {
                                case 42:
                                    v4 += 3.14159
                                case 0:
                                    for fi11 := 0; fi11 < 9; fi11++ {
                                        fmt.Println(fi11, v9)
                                        _ = fi11
                                    }
                                default:
                                    for fi12 := 0; fi12 < 10; fi12++ {
                                        v4 += float32(((-2.25 / 1e10) * 3.14159))
                                        _ = fi12
                                    }
                                }
                                if false {
                                    fmt.Println(fi8, v4)
                                    var v13 uint = 2
                                    _ = v13
                                }
                            } else {
                                v9 -= 257
                            }
                            if false {
                                v14 := GenS2{F0: 2, F1: 42}
                                if false {
                                    v9 = v9
                                    _ = len(v3)
                                    var v15 bool = (77 <= 42)
                                    _ = v15
                                } else {
                                    fmt.Println(v4, v9, v9)
                                    v9 = (v9 & v9)
                                }
                                for fi16 := 0; fi16 < 7; fi16++ {
                                    if true {
                                        if true {
                                            v4 = ((float32(-2.25) / float32((3.14159 * 0.1))) + v4)
                                            delete(v3, 65535)
                                        } else {
                                            switch "foo bar" {
                                            case "hello":
                                                _ = len(v2)
                                            case "a":
                                                var v17 uint = 255
                                                _ = v17
                                            case "foo bar":
                                                v9 -= -1
                                            }
                                        }
                                        if (((v4 + -2.25) < v4) && (true || ((2147483647 >= 861) && ((-(v9)) < (v9 * (-(v9))))))) {
                                            v4 -= v4
                                            var v18 bool = (18446744073709551615 < 274)
                                            _ = v18
                                        }
                                    } else {
                                        if (true || true) {
                                            var v19 int16 = 1
                                            _ = v19
                                        } else {
                                            if (!true) {
                                                var v20 int16 = 42
                                                var v21 float32 = (1.5 * (((v4 - (v4 / v4)) * ((float32((0.1 * (((3.14159 / -2.25) / float64((1.5 - -2.25))) * (0.0 + 0.1)))) / 1.0) / v4)) / v4))
                                                _ = v20
                                                _ = v21
                                            } else {
                                                v22 := GenS3{F0: 32767, F1: 145, F2: ""}
                                                fmt.Println(fi8)
                                                v23 := GenS1{F0: 2}
                                                _ = v22
                                                _ = v23
                                            }
                                            switch fi16 {
                                            case 9223372036854775807:
                                                switch ("hello" + "hello") {
                                                case "":
                                                    v24 := GenS4{F0: -926}
                                                    _ = v24
                                                default:
                                                    _ = len(v2)
                                                }
                                            case -262:
                                                if (-168 <= fi16) {
                                                    if (255 > 0) {
                                                        var v25 uint64 = 18446744073709551615
                                                        _ = len(v2)
                                                        var v26 int8 = 105
                                                        if (0 >= 2) {
                                                            v26 = v26
                                                            fmt.Println(v9)
                                                            v4 = -2.25
                                                        }
                                                        _ = v25
                                                        _ = v26
                                                    } else {
                                                        if (fi8 <= fi16) {
                                                            var v27 uint64 = 330
                                                            _ = v27
                                                        }
                                                    }
                                                } else {
                                                    v9 -= (v9 / ((v9) | 1))
                                                }
                                            }
                                            if (!true) {
                                                var v28 int32 = v9
                                                v9 += (v9 ^ (^(v28)))
                                                _ = v28
                                            } else {
                                                var v29 int16 = -1
                                                _ = v29
                                            }
                                        }
                                        var v30 int16 = -89
                                        var v31 int64 = -877
                                        v9 += (v9 + 2147483647)
                                        _ = v30
                                        _ = v31
                                    }
                                    switch ("a" + ("a" + "")) {
                                    case "foo bar":
                                        v4 += (((0.1 / v4) / ((v4 / (v4 * 1e10)) + (float32((3.14159 * 1.5)) + v4))) / 1.0)
                                    default:
                                        var v32 int16 = 32767
                                        _ = v32
                                    }
                                    _ = fi16
                                }
                                _ = v14
                            }
                            v33 := GenS5{F0: -411, F1: 65535}
                            _ = v33
                        }
                        _ = v10
                    }
                    _ = fi8
                    _ = v9
                }
                for fi34 := 0; fi34 < 9; fi34++ {
                    fmt.Println(v4)
                    for fi35 := 0; fi35 < 1; fi35++ {
                        v4 -= 1e10
                        var v36 float64 = float64(0.1)
                        _ = fi35
                        _ = v36
                    }
                    _ = fi34
                }
            }
        case "foo bar":
            v37 := GenS2{F0: 42, F1: -55}
            _ = v37
        }
    case -1:
        if true {
            v4 = 1.5
        } else {
            v4 = 1e10
            for fi38 := 0; fi38 < 19; fi38++ {
                v3[84] = "hello"
                v39 := v3[4294967295]
                v4 += v4
                _ = fi38
                _ = v39
            }
        }
    default:
        var v40 uint = 255
        _ = v40
    }
    var v41 uint16 = 455
    if true {
        var v42 int = 0
        _ = v42
    }
    for fi43 := 0; fi43 < 15; fi43++ {
        for fi44 := 0; fi44 < 5; fi44++ {
            if true {
                v45 := GenS2{F0: (v41 | v41), F1: -128}
                v3[0] = "\n\t"
                _ = v45
            } else {
                switch (fi44 << ((0) & 63)) {
                case -387:
                    v4 = v4
                case -9223372036854775808:
                    switch fi43 {
                    case 32768:
                        fmt.Println(fi44, fi43, fi44)
                    case 0:
                        v41 -= v41
                    default:
                        v4 -= 0.0
                    }
                case -191:
                    var v46 int = (fi44 - (fi44 / (((fi44 * (fi44 * fi44))) | 1)))
                    _ = v46
                default:
                    v4 -= ((float32(1e10) - 3.14159) + float32(float64(v4)))
                }
                var v47 int16 = 0
                switch fi44 {
                case 32767:
                    var v48 float32 = v4
                    _ = v48
                case -340:
                    switch ("a" + (((("" + "a") + "\n\t") + ("a" + "\n\t")) + ("" + "hello"))) {
                    case "a":
                        _ = len(v3)
                    case "\n\t":
                        fmt.Println(v47, fi43)
                    case "hello":
                        v49 := v3[2]
                        _ = v49
                    }
                default:
                    switch ((("" + "\n\t") + "hello") + "a") {
                    case "hello":
                        if (9223372036854775807 != -79) {
                            switch (fi44 % ((fi44) | 1)) {
                            case 557:
                                v50 := GenS6{F0: v47, F1: v47, F2: ("a" + "hello")}
                                _ = v50
                            }
                            var v51 string = "hello"
                            _ = v51
                        } else {
                            var v52 float32 = ((float32(1e10) + (1.5 - v4)) - v4)
                            v53 := GenS4{F0: -569}
                            _ = v52
                            _ = v53
                        }
                    default:
                        var v54 uint16 = v41
                        _ = v54
                    }
                }
                var v55 int = (fi43 / (((fi44 & fi44)) | 1))
                if (!true) {
                    var v56 bool = (0 >= 263)
                    _ = v56
                }
                _ = v47
                _ = v55
            }
            v57 := GenS7{F0: 100}
            _ = fi44
            _ = v57
        }
        switch "foo bar" {
        case "a":
            for fi58 := 0; fi58 < 4; fi58++ {
                var v59 uint32 = 0
                if (true || false) {
                    v4 += (v4 + (float32(0.1) / 1.0))
                }
                _ = fi58
                _ = v59
            }
        default:
            var v60 int = fi43
            _ = v60
        }
        fmt.Println(v4, v4, v4)
        _ = fi43
    }
    v41 = v41
    var v61 uint8 = 0
    var v62 uint16 = v41
    v63 := GenS8{F0: 766, F1: 0}
    switch "foo bar" {
    case "\n\t":
        if false {
            var v64 uint64 = 42
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            if ((v61 >> ((510) & 63)) > v61) {
                for fi65 := 0; fi65 < 19; fi65++ {
                    v5[(fi65 / (((fi65 / ((fi65) | 1))) | 1))] = (fi65 << ((955) & 63))
                    _ = fi65
                }
                var v66 uint64 = v64
                if (!(!("\n\t" == ""))) {
                    for fi67 := 0; fi67 < 9; fi67++ {
                        var v68 bool = true
                        _ = len(v5)
                        _ = fi67
                        _ = v68
                    }
                    fmt.Println(v62, v61)
                    if ("\n\t" == (("hello" + "foo bar") + ("\n\t" + "a"))) {
                        for fi69 := 0; fi69 < 2; fi69++ {
                            var v70 int16 = -223
                            switch 0 {
                            case 4294967295:
                                fmt.Println(v66)
                            default:
                                v61 = 127
                            }
                            for fi71 := 0; fi71 < 8; fi71++ {
                                switch int(v70) {
                                case 0:
                                    v66 = 731
                                case -240:
                                    _ = len(v2)
                                default:
                                    fmt.Println(v62, v4, v41)
                                    v1 = append(v1, "\n\t")
                                }
                                var v72 int = (fi71 << ((0) & 63))
                                _ = fi71
                                _ = v72
                            }
                            _ = fi69
                            _ = v70
                        }
                        fmt.Println(v41, v61)
                    } else {
                        for fi73 := 0; fi73 < 13; fi73++ {
                            var v74 int16 = 32767
                            for fi75 := 0; fi75 < 8; fi75++ {
                                switch fi75 {
                                case 1:
                                    v61 -= (v61 & 42)
                                }
                                switch "foo bar" {
                                case "a":
                                    switch (fi75 % ((fi75) | 1)) {
                                    case 562:
                                        if false {
                                            fmt.Println(v74, v74)
                                        }
                                    case 0:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    default:
                                        if (1.5 != 0.1) {
                                            fmt.Println(fi75, v66)
                                        }
                                    }
                                case "\n\t":
                                    switch "hello" {
                                    case "":
                                        var v76 int64 = -9223372036854775808
                                        _ = v76
                                    case "\n\t":
                                        var v77 string = "\n\t"
                                        _ = v77
                                    case "foo bar":
                                        v41 += v62
                                    default:
                                        if true {
                                            fmt.Println(v74, v61, v41)
                                            var v78 uint8 = 1
                                            v5[-1] = fi75
                                            _ = v78
                                        } else {
                                            v79 := GenS5{F0: (v74 - -32768), F1: (v62 + (^(v41)))}
                                            v62 -= 1
                                            fmt.Println(v74, v41)
                                            v80 := GenS1{F0: 4294967295}
                                            var v81 uint32 = 477
                                            _ = v79
                                            _ = v80
                                            _ = v81
                                        }
                                    }
                                }
                                v66 += v66
                                _ = fi75
                            }
                            _ = fi73
                            _ = v74
                        }
                    }
                }
                _ = v66
            } else {
                v3[40] = ("" + "hello")
                v82 := GenS9{F0: 510}
                _ = v82
            }
            _ = v64
        } else {
            v83 := GenS9{F0: 774}
            for fi84 := 0; fi84 < 1; fi84++ {
                v61 -= v61
                if true {
                    var v85 int64 = 0
                    _ = v85
                }
                _ = fi84
            }
            _ = v83
        }
    case "foo bar":
        for fi86 := 0; fi86 < 14; fi86++ {
            for fi87 := 0; fi87 < 12; fi87++ {
                if ((9223372036854775807 < -778) || true) {
                    var v88 int64 = 0
                    switch ("" + "a") {
                    case "hello":
                        if (true && true) {
                            v2 = append(v2, 65535)
                        } else {
                            fmt.Println(v61, fi87, v61)
                            var v89 int = (fi87 % ((-2147483648) | 1))
                            var v90 string = ("" + "a")
                            _ = v89
                            _ = v90
                        }
                    case "foo bar":
                        v91 := v5[42]
                        _ = v91
                    default:
                        v4 -= -2.25
                    }
                    v92 := GenS2{F0: (v62 * v62), F1: -109}
                    _ = v88
                    _ = v92
                }
                v62 -= 860
                _ = fi87
            }
            v62 = 837
            v61 += 0
            _ = fi86
        }
    default:
        var v93 uint8 = v61
        _ = v93
    }
    fmt.Println(v62)
    var v94 int16 = -1
    for fi95 := 0; fi95 < 5; fi95++ {
        for fi96 := 0; fi96 < 17; fi96++ {
            var v97 uint64 = 42
            _ = fi96
            _ = v97
        }
        for fi98 := 0; fi98 < 12; fi98++ {
            _ = len(v3)
            _ = fi98
        }
        _ = fi95
    }
    v99 := GenS5{F0: v94, F1: uint16(v61)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v61
    _ = v62
    _ = v63
    _ = v94
    _ = v99
}
