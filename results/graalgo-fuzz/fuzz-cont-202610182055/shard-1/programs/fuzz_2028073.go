package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS6 struct { F0 int; F1 int; F2 bool }
type GenS7 struct { F0 uint; F1 bool; F2 int }
type GenS8 struct { F0 int8; F1 uint32 }
type GenS9 struct { F0 uint8; F1 int64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int32; F1 int16; F2 int }

func main() {
    v1 := []int{127, 155}
    v2 := make(map[uint32]int)
    var v3 float32 = 1e10
    var v4 uint32 = 2147483647
    v4 -= v4
    v2[1] = 128
    v5 := GenS1{F0: -1, F1: 571}
    var v6 bool = true
    if v6 {
        if (!(28 != 511)) {
            for fi7 := 0; fi7 < 11; fi7++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi8 := 0; fi8 < 11; fi8++ {
                    v6 = true
                    v9 := GenS1{F0: -1, F1: 2147483647}
                    v10 := GenS1{F0: 246, F1: 763}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = fi7
            }
            v11 := GenS1{F0: -32768, F1: -1}
            v3 -= ((v3 + 3.14159) * 3.14159)
            _ = v11
        } else {
            v4 += v4
            switch -782 {
            case -1:
                var v12 float32 = 0.1
                _ = v12
            case -104:
                delete(v2, v4)
            }
        }
        v13 := GenS2{F0: (0 <= 100)}
        switch "\n\t" {
        case "hello":
            v3 -= v3
        case "a":
            var v14 float64 = 1e10
            _ = v14
        }
        _ = v13
    }
    v1 = append(v1, 128)
    var v15 string = ""
    v15 = "hello"
    if true {
        fmt.Println(v6, v15)
        v16 := GenS2{F0: false}
        _ = v16
    }
    if (!true) {
        fmt.Println(v6)
        var v17 string = v15
        _ = v17
    } else {
        if (267 >= 436) {
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 int64 = 1
                v2[(v4 ^ v4)] = int(v19)
                if v6 {
                    if v6 {
                        v20 := GenS3{F0: v4, F1: v4, F2: 79}
                        _ = v20
                    } else {
                        if ((("" + (v15 + ((v15 + v15) + v15))) + (v15 + "foo bar")) != ("\n\t" + ("" + ("" + v15)))) {
                            if (661 <= 2) {
                                if true {
                                    v3 = (v3 - float32(0.0))
                                    var v21 int32 = 1
                                    var v22 string = v15
                                    _ = v21
                                    _ = v22
                                }
                                v3 = float32(0.0)
                                v15 = (((v15 + v15) + "hello") + v15)
                            } else {
                                var v23 int16 = 42
                                _ = v23
                            }
                            for fi24 := 0; fi24 < 7; fi24++ {
                                var v25 int = fi18
                                var v26 float32 = (-2.25 - v3)
                                switch "hello" {
                                case "a":
                                    if len(v1) > 0 {
                                        v1[0] = (fi24 * (v25 * (fi18 >> ((65535) & 63))))
                                    }
                                case "\n\t":
                                    var v27 uint32 = (v4 % ((0) | 1))
                                    _ = v27
                                case "foo bar":
                                    if (127 == 183) {
                                        switch -564 {
                                        case -9223372036854775808:
                                            fmt.Println(v15)
                                            v19 += -9223372036854775808
                                            v26 += (v3 / (v26 / v3))
                                            var v28 uint8 = 2
                                            if (!true) {
                                                v29 := GenS1{F0: 1000, F1: 42}
                                                fmt.Println(v15, fi18)
                                                if ((v19 < (v19 * v19)) || v6) {
                                                    var v30 bool = v6
                                                    v26 += v26
                                                    var v31 int = fi24
                                                    _ = v30
                                                    _ = v31
                                                }
                                                v26 = 0.0
                                                _ = v29
                                            } else {
                                                switch v15 {
                                                case "hello":
                                                    fmt.Println(v15, fi24, v26)
                                                case "foo bar":
                                                    if (v6 || true) {
                                                        v4 = v4
                                                        var v32 float32 = ((-2.25 + -2.25) / float32(-2.25))
                                                        var v33 int64 = (-(v19))
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                    var v34 int8 = -117
                                                    v35 := GenS2{F0: v6}
                                                    _ = v34
                                                    _ = v35
                                                default:
                                                    var v36 int64 = v19
                                                    _ = v36
                                                }
                                            }
                                            _ = v28
                                        default:
                                            v2[v4] = v25
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = (v25 * fi18)
                                        }
                                        v4 = v4
                                        switch fi24 {
                                        case -627:
                                            v1 = append(v1, v25)
                                        case 42:
                                            v6 = v6
                                            v37 := GenS2{F0: v6}
                                            v19 = (v19 | v19)
                                            _ = v37
                                        }
                                        v6 = v6
                                    } else {
                                        v25 = fi18
                                    }
                                }
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                        } else {
                            for fi38 := 0; fi38 < 14; fi38++ {
                                v3 -= float32(1e10)
                                v39 := GenS3{F0: (^(v4)), F1: v4, F2: 909}
                                if (!(false && v6)) {
                                    var v40 string = (v15 + v15)
                                    if (v3 != float32(float64((3.14159 / float32(1.5))))) {
                                        if v6 {
                                            fmt.Println(v19)
                                            v1 = append(v1, fi38)
                                            v2[(v4 >> ((1) & 63))] = fi38
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        fmt.Println(v15, fi18)
                                    } else {
                                        delete(v2, v4)
                                        switch (fi18 & (fi18 % ((int(v4)) | 1))) {
                                        case 9223372036854775807:
                                            var v41 int64 = v19
                                            _ = v41
                                        case 127:
                                            var v42 int64 = v19
                                            _ = v42
                                        case -9223372036854775808:
                                            if (762 < 999) {
                                                if (v6 && v6) {
                                                    v15 = "\n\t"
                                                }
                                            }
                                        }
                                    }
                                    var v43 int32 = 3
                                    _ = v40
                                    _ = v43
                                }
                                _ = fi38
                                _ = v39
                            }
                            v44 := GenS4{F0: fi18, F1: 305}
                            for fi45 := 0; fi45 < 9; fi45++ {
                                var v46 int = fi18
                                if len(v1) > 0 {
                                    v1[0] = (^(v46))
                                }
                                switch 128 {
                                case 2147483647:
                                    if (557 >= 1) {
                                        var v47 uint64 = 0
                                        _ = v47
                                    } else {
                                        var v48 float64 = (1e10 - 1.5)
                                        var v49 uint8 = 13
                                        _ = v48
                                        _ = v49
                                    }
                                }
                                _ = fi45
                                _ = v46
                            }
                            _ = v44
                        }
                        fmt.Println(v3, v15, v19)
                        var v50 uint8 = 0
                        _ = v50
                    }
                    for fi51 := 0; fi51 < 10; fi51++ {
                        switch (fi51 + fi18) {
                        case 32768:
                            var v52 uint32 = 42
                            _ = v52
                        case 4294967295:
                            v53 := GenS5{F0: 67, F1: 1, F2: 2}
                            v19 -= -9223372036854775808
                            v19 += 1
                            _ = v53
                        case -2147483648:
                            v3 = (-2.25 * ((v3 - v3) / 1.0))
                        default:
                            if (!v6) {
                                var v54 int = (fi51 & 128)
                                _ = v54
                            }
                        }
                        var v55 float64 = 1.5
                        _ = fi51
                        _ = v55
                    }
                }
                _ = fi18
                _ = v19
            }
            var v56 int32 = -1
            _ = v56
        }
        v6 = (0 <= 127)
        v4 += (v4 & v4)
    }
    switch -847 {
    case -906:
        if (int8(v4) == 1) {
            for fi57 := 0; fi57 < 7; fi57++ {
                if (!((685 == 350) || (v6 && true))) {
                    switch 685 {
                    case -430:
                        fmt.Println(fi57, fi57)
                    case 32768:
                        v6 = (!v6)
                    default:
                        for fi58 := 0; fi58 < 19; fi58++ {
                            v3 = (float32(3.14159) / 0.1)
                            switch (("" + v15) + v15) {
                            case "\n\t":
                                var v59 int32 = -548
                                _ = v59
                            case "hello":
                                var v60 float32 = 1e10
                                _ = v60
                            case "a":
                                fmt.Println(fi58, v4)
                            }
                            fmt.Println(fi58, fi57, v4)
                            var v61 int32 = -1
                            v61 -= -2147483648
                            delete(v2, 0)
                            var v62 uint32 = v4
                            _ = fi58
                            _ = v61
                            _ = v62
                        }
                    }
                    _ = len(v2)
                    var v63 uint64 = 0
                    _ = v63
                }
                _ = fi57
            }
            fmt.Println(v3, v4, v6)
            v4 = 42
        } else {
            v64 := GenS6{F0: 127, F1: -217, F2: true}
            var v65 uint = 2
            _ = v64
            _ = v65
        }
    case 32767:
        if ((0.0 - 3.14159) <= float64(v3)) {
            var v66 uint32 = v4
            v3 = 1.5
            _ = v66
        } else {
            var v67 uint16 = 1
            switch 9223372036854775807 {
            case 704:
                fmt.Println(v15, v15, v4)
            default:
                v3 = (v3 / 3.14159)
            }
            _ = v67
        }
    }
    switch -725 {
    case 1:
        var v68 float32 = v3
        _ = v68
    }
    fmt.Println(v15, v6)
    v6 = v6
    if v6 {
        switch "a" {
        case "hello":
            v69 := GenS6{F0: -402, F1: 9223372036854775807, F2: v6}
            _ = v69
        case "a":
            switch (v15 + (("a" + ((v15 + (v15 + ((v15 + (((v15 + (("\n\t" + "") + v15)) + v15) + "hello")) + v15))) + v15)) + "\n\t")) {
            case "foo bar":
                var v70 uint16 = 255
                _ = v70
            case "a":
                v2[2147483648] = -219
            default:
                fmt.Println(v15, v4)
            }
        }
        var v71 uint = 18446744073709551615
        var v72 int8 = 100
        _ = v71
        _ = v72
    } else {
        fmt.Println(v6)
        _ = len(v2)
    }
    if (4294967295 != 4294967295) {
        if (255 != 108) {
            for fi73 := 0; fi73 < 20; fi73++ {
                fmt.Println(v6, v3, v15)
                var v74 string = v15
                if v6 {
                    for fi75 := 0; fi75 < 14; fi75++ {
                        var v76 uint16 = 32767
                        if ((^(fi73)) > -362) {
                            switch (fi75 % ((4294967295) | 1)) {
                            case 0:
                                switch fi73 {
                                case -9223372036854775808:
                                    fmt.Println(fi73)
                                }
                            case 127:
                                v77 := GenS7{F0: 424, F1: v6, F2: fi75}
                                _ = v77
                            case 32767:
                                v78 := GenS8{F0: -1, F1: v4}
                                _ = v78
                            default:
                                v79 := GenS6{F0: fi75, F1: fi73, F2: (1 > 2)}
                                _ = v79
                            }
                            var v80 int64 = 1
                            v81 := GenS1{F0: 42, F1: 42}
                            _ = v80
                            _ = v81
                        }
                        if v6 {
                            if (!(!v6)) {
                                var v82 bool = false
                                v83 := GenS5{F0: 47, F1: 42, F2: v4}
                                _ = v82
                                _ = v83
                            } else {
                                v2[(v4 / (((v4 & 42)) | 1))] = (fi73 << ((4294967295) & 63))
                            }
                            _ = len(v1)
                            v74 = v15
                        } else {
                            if v6 {
                                v2[v4] = fi75
                                if (4294967295 <= 4294967295) {
                                    v74 = ""
                                    if v6 {
                                        if v6 {
                                            fmt.Println(fi73, v76)
                                        } else {
                                            _ = len(v2)
                                        }
                                        var v84 int16 = 213
                                        _ = v84
                                    } else {
                                        switch v15 {
                                        case "foo bar":
                                            var v85 uint64 = 0
                                            _ = v85
                                        case "a":
                                            fmt.Println(v15, v3)
                                        case "\n\t":
                                            var v86 uint = 65535
                                            _ = v86
                                        }
                                        var v87 bool = v6
                                        _ = v87
                                    }
                                    var v88 uint64 = 1
                                    _ = v88
                                } else {
                                    v1 = append(v1, (fi75 << ((555) & 63)))
                                }
                                v89 := GenS2{F0: (v74 != v15)}
                                _ = v89
                            }
                        }
                        _ = fi75
                        _ = v76
                    }
                    if v6 {
                        v15 = "\n\t"
                        if len(v1) > 0 {
                            v1[0] = (fi73 + (fi73 - fi73))
                        }
                        v90 := GenS9{F0: 215, F1: -1}
                        _ = v90
                    } else {
                        for fi91 := 0; fi91 < 3; fi91++ {
                            if v6 {
                                var v92 float64 = 0.1
                                var v93 float32 = (v3 - 0.0)
                                _ = v92
                                _ = v93
                            }
                            v94 := GenS9{F0: 127, F1: 42}
                            fmt.Println(v3)
                            _ = fi91
                            _ = v94
                        }
                        v2[(v4 | (v4 % ((v4) | 1)))] = (fi73 >> ((42) & 63))
                    }
                } else {
                    v3 -= (v3 + ((1.5 - v3) / -2.25))
                    var v95 int16 = -917
                    _ = v95
                }
                _ = fi73
                _ = v74
            }
            if true {
                for fi96 := 0; fi96 < 3; fi96++ {
                    switch fi96 {
                    case -24:
                        fmt.Println(v15)
                    case -697:
                        v97 := GenS4{F0: (fi96 << ((2) & 63)), F1: (v4 / ((v4) | 1))}
                        _ = v97
                    default:
                        v98 := GenS1{F0: -915, F1: -1}
                        _ = v98
                    }
                    fmt.Println(v15)
                    _ = fi96
                }
                switch 837 {
                case 4294967295:
                    fmt.Println(v6, v4, v15)
                case 106:
                    var v99 uint32 = (v4 * v4)
                    _ = v99
                default:
                    for fi100 := 0; fi100 < 15; fi100++ {
                        v15 = "foo bar"
                        switch fi100 {
                        case 661:
                            var v101 int32 = -1
                            _ = v101
                        case 42:
                            v102 := GenS4{F0: -1, F1: v4}
                            _ = v102
                        default:
                            var v103 float32 = v3
                            _ = v103
                        }
                        _ = fi100
                    }
                }
            }
        }
        fmt.Println(v6, v6, v3)
        var v104 uint16 = 978
        _ = v104
    } else {
        var v105 int8 = -128
        fmt.Println(v105)
        _ = v105
    }
    switch 9223372036854775807 {
    case 32767:
        v106 := GenS3{F0: (v4 % ((v4) | 1)), F1: v4, F2: 900}
        _ = v106
    case 128:
        if v6 {
            v3 = v3
            if (!v6) {
                if v6 {
                    fmt.Println(v15, v4, v6)
                    var v107 float64 = ((1e10 - (0.1 - 3.14159)) / float64(1e10))
                    fmt.Println(v15)
                    _ = v107
                } else {
                    switch 278 {
                    case -315:
                        v108 := GenS10{F0: -1}
                        _ = v108
                    case -55:
                        var v109 int = 32768
                        _ = v109
                    }
                    for fi110 := 0; fi110 < 9; fi110++ {
                        delete(v2, v4)
                        v15 = ((v15 + v15) + v15)
                        switch "hello" {
                        case "":
                            v15 = "foo bar"
                        case "foo bar":
                            for fi111 := 0; fi111 < 3; fi111++ {
                                v6 = (true && (-1 <= 68))
                                switch fi110 {
                                case 127:
                                    if len(v1) > 0 {
                                        v1[0] = fi110
                                    }
                                case 42:
                                    if v6 {
                                        if ((v4 << ((42) & 63)) != (v4 & 609)) {
                                            v15 = ((v15 + "\n\t") + v15)
                                        }
                                    } else {
                                        if v6 {
                                            v112 := GenS11{F0: 42, F1: -32768, F2: (fi111 >> ((731) & 63))}
                                            var v113 uint8 = 1
                                            v113 = v113
                                            _ = v112
                                            _ = v113
                                        } else {
                                            var v114 float32 = -2.25
                                            v115 := GenS9{F0: 128, F1: 42}
                                            _ = v114
                                            _ = v115
                                        }
                                        v116 := GenS7{F0: 727, F1: true, F2: (fi110 % ((fi110) | 1))}
                                        _ = v116
                                    }
                                }
                                var v117 uint = 2
                                _ = fi111
                                _ = v117
                            }
                        }
                        _ = fi110
                    }
                }
            } else {
                var v118 string = ("a" + "hello")
                fmt.Println(v118, v118)
                _ = v118
            }
            fmt.Println(v3)
        } else {
            delete(v2, v4)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = len(v2)
        }
    case -2147483648:
        delete(v2, (v4 << ((1) & 63)))
    default:
        if len(v1) > 0 {
            v1[0] = 599
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
