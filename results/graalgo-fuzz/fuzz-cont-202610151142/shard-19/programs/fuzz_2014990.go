package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 int16; F2 int }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 uint8; F1 int16 }
type GenS9 struct { F0 uint16 }
func genFn0(p0 uint32, p1 int) uint64 {
    var v6 uint16 = 2
    if true {
        for fi7 := 0; fi7 < 19; fi7++ {
            v6 += v6
            v8 := GenS1{F0: 42}
            _ = fi7
            _ = v8
        }
        switch ("foo bar" + ("a" + "a")) {
        case "foo bar":
            var v9 uint = 18446744073709551615
            _ = v9
        default:
            if false {
                switch p1 {
                case 128:
                    fmt.Println(v6, v6)
                default:
                    switch "\n\t" {
                    case "\n\t":
                        v6 = v6
                    case "":
                        v10 := GenS1{F0: 1}
                        _ = v10
                    }
                }
                fmt.Println(p0, v6)
            }
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            fmt.Println(p0)
            switch fi11 {
            case 128:
                switch ("" + "\n\t") {
                case "a":
                    fmt.Println(p0, p0, p1)
                case "\n\t":
                    var v12 uint = 255
                    _ = v12
                default:
                    fmt.Println(v6, v6, fi11)
                }
            case 4294967295:
                var v13 int16 = 42
                _ = v13
            case 1:
                if (!false) {
                    if false {
                        for fi14 := 0; fi14 < 18; fi14++ {
                            var v15 float32 = 0.1
                            _ = fi14
                            _ = v15
                        }
                    } else {
                        switch (fi11 ^ (p1 >> ((523) & 63))) {
                        case 32768:
                            var v16 uint64 = 0
                            _ = v16
                        }
                        v17 := GenS1{F0: 42}
                        _ = v17
                    }
                    v18 := GenS1{F0: -9223372036854775808}
                    if (false && true) {
                        v19 := GenS1{F0: -227}
                        _ = v19
                    }
                    _ = v18
                }
            }
            _ = fi11
        }
    } else {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 int16 = 1
            fmt.Println(p0)
            switch "hello" {
            case "\n\t":
                switch (fi20 / (((^(p1))) | 1)) {
                case 32767:
                    var v22 int64 = -1
                    _ = v22
                case -2147483648:
                    var v23 string = "\n\t"
                    _ = v23
                case 4294967295:
                    p0 = p0
                default:
                    v21 = 1
                }
            default:
                for fi24 := 0; fi24 < 11; fi24++ {
                    var v25 string = (("hello" + ("a" + "")) + "a")
                    if true {
                        v26 := GenS2{F0: 2, F1: v25}
                        v27 := GenS1{F0: 42}
                        v28 := []int{(p1 >> ((42) & 63)), (fi24 - (-(p1))), 726, -1}
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    v29 := []int{1}
                    _ = fi24
                    _ = v25
                    _ = v29
                }
            }
            _ = fi20
            _ = v21
        }
        for fi30 := 0; fi30 < 9; fi30++ {
            fmt.Println(fi30, fi30, p0)
            var v31 int64 = 42
            switch "hello" {
            case "hello":
                for fi32 := 0; fi32 < 2; fi32++ {
                    var v33 uint = uint(p0)
                    fmt.Println(v31, fi32, fi32)
                    _ = fi32
                    _ = v33
                }
            case "a":
                if false {
                    v34 := GenS1{F0: int64(p1)}
                    var v35 int8 = 100
                    _ = v34
                    _ = v35
                } else {
                    v6 = (v6 << ((0) & 63))
                }
            case "\n\t":
                if true {
                    var v36 float64 = (1.5 + ((1e10 - 0.0) - 1e10))
                    for fi37 := 0; fi37 < 2; fi37++ {
                        switch "" {
                        case "\n\t":
                            var v38 int64 = v31
                            _ = v38
                        case "a":
                            v39 := []int{p1}
                            _ = v39
                        }
                        switch fi30 {
                        case 42:
                            fmt.Println(v6)
                        case 9223372036854775807:
                            fmt.Println(v6)
                            p0 = (p0 + p0)
                            var v40 string = "hello"
                            _ = v40
                        }
                        var v41 uint16 = (v6 / ((0) | 1))
                        _ = fi37
                        _ = v41
                    }
                    var v42 uint8 = 0
                    _ = v36
                    _ = v42
                } else {
                    for fi43 := 0; fi43 < 5; fi43++ {
                        v31 = v31
                        switch fi43 {
                        case 128:
                            v6 -= 32767
                        default:
                            fmt.Println(p1, fi43)
                        }
                        fmt.Println(v31)
                        var v44 uint = 1
                        _ = fi43
                        _ = v44
                    }
                    v45 := []int{-635, 2147483647}
                    _ = v45
                }
            }
            _ = fi30
            _ = v31
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1() uint {
    if (!(!false)) {
        var v46 uint64 = genFn0(2, 1)
        if true {
            for fi47 := 0; fi47 < 17; fi47++ {
                var v48 float64 = 3.14159
                v49 := GenS3{F0: -1}
                _ = fi47
                _ = v48
                _ = v49
            }
            switch "foo bar" {
            case "a":
                v50 := []int{42, -1, -2147483648}
                _ = v50
            case "foo bar":
                if false {
                    v51 := []int{32767, 385}
                    for fi52 := 0; fi52 < 18; fi52++ {
                        _ = len(v51)
                        var v53 int32 = 1
                        _ = fi52
                        _ = v53
                    }
                    _ = v51
                } else {
                    v54 := []int{9223372036854775807, 1, 42, 32767}
                    _ = v54
                }
            default:
                v55 := GenS4{F0: "\n\t", F1: 2}
                _ = v55
            }
        } else {
            v56 := GenS5{F0: -2147483648}
            if (true || false) {
                v46 = (^(v46))
            } else {
                v57 := []int{32767}
                if (true && false) {
                    fmt.Println(v46)
                    var v58 uint8 = 255
                    _ = v58
                } else {
                    v59 := GenS4{F0: "hello", F1: 255}
                    _ = v59
                }
                _ = v57
            }
            v46 += (v46 / ((1) | 1))
            _ = v56
        }
        _ = v46
    }
    if false {
        var v60 float32 = (1e10 / -2.25)
        var v61 string = "\n\t"
        _ = v60
        _ = v61
    } else {
        var v62 float32 = -2.25
        _ = v62
    }
    switch -9223372036854775808 {
    case 4294967295:
        v63 := []int{9223372036854775807, -195, -134, 595}
        _ = v63
    }
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 255
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]uint32)
    var v4 int8 = -128
    v5 := []uint32{2147483647, 65535}
    if ((0 != 279) && false) {
        if (false && ((1 < 0) && (-2147483648 <= 975))) {
            for fi64 := 0; fi64 < 14; fi64++ {
                for fi65 := 0; fi65 < 15; fi65++ {
                    v66 := GenS6{F0: -962, F1: 1, F2: fi65}
                    var v67 uint8 = 128
                    _ = fi65
                    _ = v66
                    _ = v67
                }
                fmt.Println(v4, v1)
                _ = fi64
            }
            v4 = v4
            v5 = append(v5, 4294967295)
        } else {
            for fi68 := 0; fi68 < 15; fi68++ {
                for fi69 := 0; fi69 < 10; fi69++ {
                    v2 = (v2 % (((v2 << ((42) & 63))) | 1))
                    _ = fi69
                }
                if (127 == 50) {
                    fmt.Println(v1)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v70 := GenS6{F0: int16(v2), F1: 0, F2: (fi68 - 1)}
                    _ = v70
                }
                for fi71 := 0; fi71 < 2; fi71++ {
                    var v72 uint16 = (v1 | (v1 >> ((355) & 63)))
                    fmt.Println(v2)
                    v3[65535] = 2
                    _ = fi71
                    _ = v72
                }
                _ = fi68
            }
            switch "a" {
            case "\n\t":
                var v73 uint32 = 171
                _ = v73
            default:
                var v74 uint16 = (v1 / ((v1) | 1))
                _ = v74
            }
            v75 := GenS7{F0: (v4 | (^(v4))), F1: 879}
            _ = v75
        }
    }
    for fi76 := 0; fi76 < 8; fi76++ {
        if (1.5 >= 3.14159) {
            v77 := GenS2{F0: 42, F1: ("foo bar" + "a")}
            if (255 >= 115) {
                switch 2147483647 {
                case 128:
                    v78 := GenS4{F0: "", F1: v1}
                    _ = v78
                default:
                    switch "" {
                    case "":
                        v1 = v1
                    case "foo bar":
                        var v79 int8 = -128
                        _ = v79
                    default:
                        var v80 int8 = -108
                        _ = v80
                    }
                }
                for fi81 := 0; fi81 < 8; fi81++ {
                    fmt.Println(v1, fi76, v1)
                    switch fi81 {
                    case 37:
                        var v82 uint64 = 42
                        _ = v82
                    case 32768:
                        if false {
                            var v83 int16 = 411
                            v83 += v83
                            switch "\n\t" {
                            case "":
                                if (fi76 < fi81) {
                                    v5 = append(v5, 0)
                                    v2 -= (v2 << ((707) & 63))
                                    var v84 int32 = -842
                                    if ((v1 * (v1 * (v1 << ((0) & 63)))) > (v1 % (((v1 & 0)) | 1))) {
                                        var v85 float64 = 0.0
                                        v86 := GenS4{F0: "a", F1: v1}
                                        if ((genFn1() < 218) && true) {
                                            var v87 uint8 = 255
                                            v84 = v84
                                            v84 = v84
                                            if ("" != "") {
                                                v84 = 1
                                                var v88 int32 = 1
                                                _ = v88
                                            }
                                            _ = v87
                                        } else {
                                            var v89 int32 = (v84 - v84)
                                            var v90 int16 = 1000
                                            _ = v89
                                            _ = v90
                                        }
                                        switch (-(fi76)) {
                                        case 4294967295:
                                            fmt.Println(v83, v83, v1)
                                        case -2147483648:
                                            if len(v5) > 0 {
                                                v5[0] = 102
                                            }
                                        case -938:
                                            switch "a" {
                                            case "a":
                                                v1 -= v1
                                            }
                                        }
                                        v3[1] = 345
                                        _ = v85
                                        _ = v86
                                    } else {
                                        var v91 int32 = (v84 | v84)
                                        v1 = 2
                                        switch "a" {
                                        case "a":
                                            var v92 float32 = -2.25
                                            _ = v92
                                        default:
                                            var v93 string = "foo bar"
                                            _ = v93
                                        }
                                        _ = v91
                                    }
                                    _ = v84
                                }
                            case "hello":
                                if true {
                                    switch "hello" {
                                    case "":
                                        fmt.Println(v83, v4)
                                    case "foo bar":
                                        v4 = v4
                                    case "\n\t":
                                        var v94 float32 = 1e10
                                        _ = v94
                                    default:
                                        var v95 int32 = 562
                                        _ = v95
                                    }
                                    v4 -= (-(v4))
                                }
                            }
                            fmt.Println(fi76)
                            _ = v83
                        } else {
                            v1 = 255
                            v2 = (v2 << ((2) & 63))
                        }
                        if (-2147483648 < -303) {
                            var v96 int = fi76
                            switch "" {
                            case "foo bar":
                                fmt.Println(fi76, fi81, fi76)
                                _ = len(v3)
                            case "":
                                if ((false || true) || true) {
                                    switch 327 {
                                    case -784:
                                        v97 := GenS1{F0: -9223372036854775808}
                                        _ = v97
                                    default:
                                        var v98 string = ("hello" + "foo bar")
                                        var v99 uint64 = 1
                                        fmt.Println(fi81, fi76, v96)
                                        _ = v98
                                        _ = v99
                                    }
                                    var v100 int32 = 786
                                    v101 := GenS8{F0: 128, F1: 1000}
                                    var v102 uint64 = 4294967295
                                    _ = v100
                                    _ = v101
                                    _ = v102
                                } else {
                                    var v103 uint = 65535
                                    _ = v103
                                }
                            default:
                                v104 := GenS9{F0: 77}
                                _ = v104
                            }
                            v3[2147483648] = 534
                            _ = v96
                        }
                        var v105 uint8 = 128
                        _ = v105
                    case 1:
                        v3[2] = 1
                    }
                    _ = fi81
                }
                fmt.Println(v1, fi76, v4)
            } else {
                fmt.Println(v2, v2, fi76)
                var v106 uint32 = 42
                var v107 int32 = 42
                _ = v106
                _ = v107
            }
            _ = v77
        } else {
            v2 -= int64(fi76)
            var v108 int16 = 1
            var v109 int32 = -1
            _ = v108
            _ = v109
        }
        v3[1] = 65535
        _ = fi76
    }
    fmt.Println(v4)
    v110 := GenS8{F0: 150, F1: 134}
    v4 = (v4 * (v4 >> ((409) & 63)))
    _ = len(v3)
    _ = len(v3)
    fmt.Println(v4)
    v1 = v1
    v1 += v1
    v2 -= (-(v2))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v111 := v3[65535]
    var v112 int32 = 1
    if (!((18446744073709551615 >= 4294967295) && (331 != 523))) {
        v113 := GenS7{F0: int8(v1), F1: -32768}
        _ = v113
    }
    var v114 int64 = (v2 / ((v2) | 1))
    v112 = (v112 >> ((genFn1()) & 63))
    var v115 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v114:", v114)
    fmt.Println("v115:", v115)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v110
    _ = v111
    _ = v112
    _ = v114
    _ = v115
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
