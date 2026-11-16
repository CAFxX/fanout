package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int32 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32; F1 string; F2 int32 }
type GenS9 struct { F0 int; F1 uint32; F2 uint8 }
type GenS10 struct { F0 uint64 }
func genFn0(p0 uint32, p1 int64, p2 int) string {
    var v6 string = "a"
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: -1, F1: 42, F2: 2147483647}
            _ = fi8
            _ = v9
        }
        switch int(p0) {
        case 211:
            if false {
                v10 := GenS1{F0: 127, F1: (^(p2)), F2: 1}
                _ = v10
            }
        case 362:
            var v11 uint64 = 546
            _ = v11
        case 9223372036854775807:
            if ((^(p0)) > 457) {
                v12 := []int{(p2 * (^(fi7)))}
                v13 := GenS2{F0: 100, F1: 127, F2: 0}
                p2 = (p2 | (fi7 / ((fi7) | 1)))
                _ = v12
                _ = v13
            } else {
                switch fi7 {
                case -2147483648:
                    p2 = fi7
                case 1:
                    if (!true) {
                        p0 += (p0 % ((p0) | 1))
                        var v14 uint = 1
                        _ = v14
                    }
                case -627:
                    var v15 string = v6
                    _ = v15
                }
            }
        }
        p1 += 0
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "\n\t"
}
func genFn1(p0 string, p1 int32) uint64 {
    for fi16 := 0; fi16 < 12; fi16++ {
        switch p0 {
        case "hello":
            switch (fi16 + -1) {
            case -1:
                var v17 uint16 = 0
                _ = v17
            default:
                fmt.Println(p0, p0, fi16)
            }
        default:
            if (255 < 65535) {
                p1 -= (p1 >> ((429) & 63))
            }
        }
        var v18 int64 = -316
        _ = fi16
        _ = v18
    }
    var v19 string = ""
    p0 = (("\n\t" + "\n\t") + ("a" + ((p0 + "\n\t") + "\n\t")))
    if ((!(-1 < 1)) || true) {
        v20 := GenS1{F0: 127, F1: -510, F2: -890}
        _ = v20
    } else {
        switch -1 {
        case 2147483647:
            fmt.Println(v19)
        case -621:
            var v21 uint64 = 42
            _ = v21
        case -400:
            p1 = p1
        default:
            if false {
                v22 := []int{}
                if (false || (18446744073709551615 > 461)) {
                    v23 := GenS3{F0: 42}
                    p0 = "hello"
                    switch -223 {
                    case -2147483648:
                        for fi24 := 0; fi24 < 16; fi24++ {
                            p0 = (v19 + "")
                            v25 := GenS3{F0: 32767}
                            var v26 uint64 = 42
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    case 2:
                        fmt.Println(p1, p0)
                    case -9223372036854775808:
                        var v27 float64 = (((0.1 * float64(0.0)) - 0.0) / 0.1)
                        _ = v27
                    default:
                        for fi28 := 0; fi28 < 16; fi28++ {
                            if false {
                                for fi29 := 0; fi29 < 5; fi29++ {
                                    if (((("foo bar" + v19) + "") + p0) != ("a" + v19)) {
                                        var v30 uint16 = 683
                                        v30 = v30
                                        _ = v30
                                    }
                                    _ = fi29
                                }
                                fmt.Println(p1)
                                v31 := GenS3{F0: 371}
                                _ = v31
                            }
                            _ = fi28
                        }
                    }
                    _ = v23
                }
                for fi32 := 0; fi32 < 18; fi32++ {
                    v19 = p0
                    fmt.Println(v19, v19)
                    _ = fi32
                }
                _ = v22
            }
        }
        for fi33 := 0; fi33 < 15; fi33++ {
            var v34 string = p0
            _ = fi33
            _ = v34
        }
    }
    _ = p0
    _ = p1
    _ = v19
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -602
    v3 := []uint32{4294967295, 0, 453}
    v4 := make(map[int]string)
    v5 := make(map[uint32]string)
    var v35 int8 = 42
    if (false && false) {
        switch 42 {
        case 42:
            v35 += 127
        case 32767:
            fmt.Println(v35, v2)
        default:
            for fi36 := 0; fi36 < 10; fi36++ {
                if false {
                    v37 := GenS1{F0: v35, F1: (fi36 << ((738) & 63)), F2: 1}
                    if (!((32767 <= -32768) && false)) {
                        if ((false || false) && (!(573 != 32767))) {
                            for fi38 := 0; fi38 < 15; fi38++ {
                                switch "\n\t" {
                                case "\n\t":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                default:
                                    if false {
                                        switch "foo bar" {
                                        case "a":
                                            if true {
                                                v35 = v35
                                            } else {
                                                fmt.Println(fi36)
                                                var v39 uint64 = 1
                                                _ = v39
                                            }
                                        default:
                                            v2 -= 9223372036854775807
                                        }
                                    } else {
                                        switch fi36 {
                                        case 0:
                                            v2 = v2
                                            if true {
                                                var v40 float32 = -2.25
                                                if true {
                                                    switch ("\n\t" + "\n\t") {
                                                    case "foo bar":
                                                        switch "" {
                                                        case "hello":
                                                            if false {
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    v1[2147483647] = 895
                                                                }
                                                                if ((1 != 32767) || false) {
                                                                    switch fi36 {
                                                                    case 1:
                                                                        v41 := GenS4{F0: v2}
                                                                        _ = v41
                                                                    }
                                                                    v35 -= (v35 / ((42) | 1))
                                                                } else {
                                                                    switch fi38 {
                                                                    case -9223372036854775808:
                                                                        v42 := GenS5{F0: 1000}
                                                                        _ = v42
                                                                    case -2147483648:
                                                                        v43 := v5[151]
                                                                        _ = v43
                                                                    default:
                                                                        var v44 uint8 = 162
                                                                        _ = v44
                                                                    }
                                                                    v2 = (^(v2))
                                                                    v2 -= 577
                                                                }
                                                            }
                                                        default:
                                                            v40 = 0.0
                                                        }
                                                    case "a":
                                                        v2 -= 474
                                                    case "\n\t":
                                                        if true {
                                                            var v45 uint16 = 2
                                                            v45 -= (v45 & v45)
                                                            switch (((("\n\t" + "a") + "hello") + ("\n\t" + "")) + (("foo bar" + (("a" + ("a" + "a")) + "foo bar")) + ("\n\t" + "\n\t"))) {
                                                            case "hello":
                                                                v46 := GenS2{F0: (v35 % (((v35 ^ v35)) | 1)), F1: 179, F2: (v45 + (^(v45)))}
                                                                _ = v46
                                                            case "a":
                                                                if true {
                                                                    v3 = append(v3, 42)
                                                                    switch (fi36 * 1) {
                                                                    case 4294967295:
                                                                        var v47 uint8 = 222
                                                                        _ = v47
                                                                    default:
                                                                        v1[42] = 1
                                                                    }
                                                                    if false {
                                                                        v48 := GenS6{F0: (false && ((127 < 1) || false))}
                                                                        _ = len(v1)
                                                                        var v49 int8 = -81
                                                                        _ = v48
                                                                        _ = v49
                                                                    } else {
                                                                        var v50 uint = uint(v45)
                                                                        if true {
                                                                            var v51 int16 = 42
                                                                            var v52 uint8 = 2
                                                                            _ = v51
                                                                            _ = v52
                                                                        }
                                                                        _ = v50
                                                                    }
                                                                }
                                                            case "":
                                                                if len(v3) > 0 {
                                                                    v3[0] = 65535
                                                                }
                                                            default:
                                                                var v53 uint8 = 0
                                                                _ = v53
                                                            }
                                                            _ = v45
                                                        }
                                                    }
                                                    switch 32768 {
                                                    case 127:
                                                        if true {
                                                            v35 -= (v35 * (v35 % (((^(v35))) | 1)))
                                                            fmt.Println(v2)
                                                        } else {
                                                            v2 -= v2
                                                        }
                                                    case 366:
                                                        switch fi38 {
                                                        case 0:
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v3 = append(v3, 2147483648)
                                                            default:
                                                                var v54 int8 = 127
                                                                _ = v54
                                                            }
                                                        case -805:
                                                            fmt.Println(fi38, fi38, fi36)
                                                        case 4294967295:
                                                            v55 := GenS4{F0: 0}
                                                            _ = v55
                                                        default:
                                                            switch "hello" {
                                                            case "foo bar":
                                                                v2 = (v2 | int64(fi36))
                                                            case "hello":
                                                                v2 -= (v2 - -9223372036854775808)
                                                            }
                                                        }
                                                    case -1:
                                                        var v56 int = (fi36 & fi38)
                                                        _ = v56
                                                    default:
                                                        if (v35 == 45) {
                                                            v57 := GenS4{F0: (v2 / ((v2) | 1))}
                                                            _ = v57
                                                        }
                                                    }
                                                    v35 += (v35 & v35)
                                                } else {
                                                    if true {
                                                        var v58 bool = true
                                                        fmt.Println(fi36)
                                                        if (true && false) {
                                                            var v59 bool = (4294967295 == 2147483648)
                                                            _ = v59
                                                        } else {
                                                            v58 = (867 >= 1)
                                                            v4[32767] = "\n\t"
                                                            if (1 == 1) {
                                                                v58 = (fi36 >= fi38)
                                                                fmt.Println(fi38, v40)
                                                                v60 := GenS7{F0: 0}
                                                                v2 = (^(v2))
                                                                _ = v60
                                                            }
                                                        }
                                                        _ = v58
                                                    }
                                                    var v61 int = (fi38 % (((fi36 >> ((42) & 63))) | 1))
                                                    var v62 int32 = -355
                                                    _ = v61
                                                    _ = v62
                                                }
                                                _ = len(v3)
                                                _ = v40
                                            } else {
                                                if ((0.1 < float64(((-2.25 / float32(0.1)) - 3.14159))) || ((((1e10 + (-2.25 - float32(((3.14159 + 3.14159) / (1.5 / (1.5 + 3.14159)))))) / float32(-2.25)) * 1.5) > 3.14159)) {
                                                    fmt.Println(v35, fi38)
                                                    v35 = (v35 >> ((19) & 63))
                                                } else {
                                                    if true {
                                                        if (((1.5 - 1e10) / 1.5) == 1e10) {
                                                            var v63 uint = 902
                                                            var v64 int32 = 431
                                                            v65 := GenS4{F0: v2}
                                                            _ = v63
                                                            _ = v64
                                                            _ = v65
                                                        } else {
                                                            v35 = (v35 + v35)
                                                            var v66 float64 = 0.1
                                                            var v67 bool = true
                                                            var v68 string = "\n\t"
                                                            var v69 int8 = (v35 & v35)
                                                            _ = v66
                                                            _ = v67
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                        _ = len(v3)
                                                    }
                                                }
                                                v70 := v1[2147483648]
                                                fmt.Println(v70)
                                                _ = v70
                                            }
                                            fmt.Println(fi36)
                                        default:
                                            var v71 float64 = float64(1.5)
                                            _ = v71
                                        }
                                        v72 := v5[2]
                                        v73 := GenS3{F0: -32768}
                                        _ = v72
                                        _ = v73
                                    }
                                }
                                v35 = v35
                                var v74 uint16 = 729
                                _ = fi38
                                _ = v74
                            }
                        } else {
                            v35 -= v35
                            fmt.Println(fi36, fi36, fi36)
                        }
                    }
                    if ((((fi36 * fi36) >= fi36) && false) && true) {
                        for fi75 := 0; fi75 < 17; fi75++ {
                            var v76 uint64 = 1
                            _ = fi75
                            _ = v76
                        }
                        v5[2147483648] = ("hello" + "")
                        var v77 uint8 = 216
                        _ = v77
                    }
                    _ = v37
                }
                v35 = 1
                _ = fi36
            }
        }
    } else {
        v78 := GenS8{F0: 2147483647, F1: ((("hello" + "") + ((("foo bar" + "hello") + "hello") + (((("" + "foo bar") + "\n\t") + ((("hello" + "") + (("a" + "\n\t") + ("hello" + "\n\t"))) + "a")) + ("\n\t" + ("hello" + "\n\t"))))) + ("" + (((("a" + ("" + ("a" + "a"))) + "foo bar") + "") + (("foo bar" + "foo bar") + "a")))), F2: -1}
        var v79 uint32 = 580
        for fi80 := 0; fi80 < 12; fi80++ {
            for fi81 := 0; fi81 < 3; fi81++ {
                v2 -= v2
                v4[fi80] = (("hello" + "\n\t") + (((("" + "hello") + "a") + "\n\t") + "hello"))
                var v82 int32 = 775
                var v83 int8 = (v35 | (v35 * v35))
                v82 += v82
                fmt.Println(v2)
                v4[fi80] = "\n\t"
                if (!true) {
                    var v84 int64 = v2
                    if (990 >= (v2 - (v2 + v2))) {
                        var v85 uint16 = 631
                        _ = len(v5)
                        if (!true) {
                            v86 := GenS5{F0: 1}
                            v82 = 1
                            _ = v86
                        } else {
                            switch (fi81 / ((fi81) | 1)) {
                            case 669:
                                var v87 bool = false
                                _ = v87
                            case -2147483648:
                                if len(v3) > 0 {
                                    v3[0] = v79
                                }
                            default:
                                var v88 int64 = 1
                                _ = v88
                            }
                            if (!(((v35 / ((v35) | 1)) == 80) || (v83 != -125))) {
                                var v89 float64 = 1e10
                                _ = v89
                            } else {
                                v90 := GenS2{F0: (v35 - 1), F1: 1, F2: v85}
                                v91 := GenS6{F0: (((false && true) && false) && true)}
                                var v92 int64 = v84
                                v92 -= (v92 / ((v84) | 1))
                                fmt.Println(v79, v79)
                                var v93 uint = 2
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                switch (fi81 & -1) {
                                case -9223372036854775808:
                                    v84 += v2
                                case -542:
                                    if (!false) {
                                        v35 = (^(v83))
                                        v94 := GenS9{F0: int(v82), F1: 2147483647, F2: 109}
                                        var v95 float64 = -2.25
                                        _ = v94
                                        _ = v95
                                    } else {
                                        var v96 int32 = int32(v2)
                                        _ = v96
                                    }
                                case 42:
                                    var v97 int32 = (v82 * v82)
                                    _ = v97
                                }
                                _ = v90
                                _ = v91
                                _ = v92
                                _ = v93
                            }
                        }
                        _ = v85
                    } else {
                        v98 := GenS10{F0: 765}
                        v83 -= (v35 % (((v35 | 60)) | 1))
                        _ = v98
                    }
                    _ = v84
                } else {
                    v82 = 113
                    var v99 int = (fi81 % (((-(fi81))) | 1))
                    v99 = (v99 * (fi80 + (v99 >> ((0) & 63))))
                    _ = v99
                }
                v79 -= v79
                var v100 uint32 = v79
                _ = fi81
                _ = v100
                _ = v82
                _ = v83
            }
            v2 += (v2 / (((v2 | v2)) | 1))
            _ = fi80
        }
        _ = v78
        _ = v79
    }
    var v101 int32 = 106
    var v102 uint32 = 42
    for fi103 := 0; fi103 < 19; fi103++ {
        var v104 int32 = v101
        _ = fi103
        _ = v104
    }
    v2 = 393
    fmt.Println(v35, v101, v35)
    switch -9223372036854775808 {
    case 32768:
        fmt.Println(v102)
    default:
        if (true || (0 > 42)) {
            _ = len(v1)
        }
    }
    fmt.Println(v101, v101)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v101
    _ = v102
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
