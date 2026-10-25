package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 uint16; F2 int64 }
func genFn0() string {
    var v7 float64 = (1e10 / 3.14159)
    v7 = (((v7 - 0.0) - (float64(3.14159) + 3.14159)) + (v7 + (0.0 - float64(float32(v7)))))
    if (1 < -3) {
        var v8 float64 = 0.1
        _ = v8
    } else {
        var v9 bool = (false || (!(!(false || false))))
        _ = v9
    }
    _ = v7
    return "hello"
}
func genFn1(p0 bool, p1 uint16, p2 int16) uint32 {
    if ((p2 + p2) <= p2) {
        p1 = p1
        var v10 uint32 = 799
        _ = v10
    } else {
        v11 := []int{2147483647, -227, 664}
        _ = v11
    }
    var v12 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 42
}

func main() {
    v1 := []int{}
    var v2 int64 = 522
    v3 := []string{"\n\t", ("a" + "\n\t"), "hello", (((("" + "a") + "hello") + "a") + "")}
    var v4 int16 = 1
    var v5 int16 = (v4 & (-(v4)))
    var v6 uint = 2
    switch "\n\t" {
    case "":
        v13 := GenS1{F0: 0, F1: 65535}
        _ = v13
    case "foo bar":
        var v14 uint16 = 253
        _ = v14
    }
    if len(v3) > 0 {
        v3[0] = "a"
    }
    fmt.Println(v2)
    if false {
        switch 32768 {
        case 128:
            _ = len(v1)
        case -1:
            for fi15 := 0; fi15 < 16; fi15++ {
                v6 = 609
                _ = fi15
            }
        case -2147483648:
            var v16 int32 = 0
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        v1 = append(v1, (fi17 | fi17))
        fmt.Println(fi17, v2)
        var v18 float32 = 3.14159
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: -123, F1: 414}
    if true {
        for fi20 := 0; fi20 < 10; fi20++ {
            if false {
                fmt.Println(v6, fi20, v4)
                v6 -= v6
            } else {
                if false {
                    fmt.Println(v4)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch 128 {
                    case 32768:
                        v21 := GenS2{F0: 235}
                        _ = v21
                    default:
                        if false {
                            v2 = -1
                            v6 -= 0
                        }
                    }
                } else {
                    for fi22 := 0; fi22 < 2; fi22++ {
                        switch "" {
                        case "hello":
                            fmt.Println(fi22, fi22)
                        case "":
                            v4 -= 0
                        default:
                            if true {
                                v5 -= (v5 & v5)
                                var v23 string = "hello"
                                v5 = v5
                                _ = v23
                            }
                        }
                        var v24 int = fi22
                        _ = fi22
                        _ = v24
                    }
                }
                for fi25 := 0; fi25 < 17; fi25++ {
                    var v26 int32 = 611
                    _ = fi25
                    _ = v26
                }
            }
            var v27 uint8 = 42
            _ = fi20
            _ = v27
        }
        fmt.Println(v2)
    } else {
        v6 += 909
        for fi28 := 0; fi28 < 17; fi28++ {
            if (!false) {
                for fi29 := 0; fi29 < 19; fi29++ {
                    switch "a" {
                    case "hello":
                        v30 := GenS2{F0: 42}
                        _ = v30
                    }
                    v2 = v2
                    _ = fi29
                }
                v4 = (v5 * v4)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v31 := GenS2{F0: 42}
            _ = fi28
            _ = v31
        }
        switch 2147483647 {
        case 2147483647:
            var v32 int = 42
            _ = v32
        }
    }
    v5 += v4
    for fi33 := 0; fi33 < 4; fi33++ {
        switch "a" {
        case "foo bar":
            var v34 uint = (v6 * v6)
            _ = v34
        case "\n\t":
            for fi35 := 0; fi35 < 9; fi35++ {
                if true {
                    var v36 float32 = (3.14159 - 0.1)
                    if (2 < v6) {
                        v5 = v4
                        v36 = (((((0.1 / v36) / v36) - v36) - v36) / v36)
                    } else {
                        var v37 uint16 = 2
                        _ = v37
                    }
                    v1 = append(v1, fi33)
                    _ = v36
                } else {
                    switch "foo bar" {
                    case "":
                        if (3.14159 >= 0.1) {
                            if true {
                                v38 := GenS2{F0: 263}
                                _ = v38
                            }
                            switch ("\n\t" + "foo bar") {
                            case "a":
                                var v39 uint8 = 128
                                _ = v39
                            default:
                                var v40 uint8 = 128
                                _ = v40
                            }
                            v41 := GenS2{F0: 189}
                            if true {
                                if (!true) {
                                    var v42 uint = v6
                                    switch ("foo bar" + "hello") {
                                    case "hello":
                                        v43 := GenS2{F0: 42}
                                        _ = v43
                                    case "a":
                                        switch (^(fi35)) {
                                        case -1:
                                            fmt.Println(v2, v5, v6)
                                            if true {
                                                v44 := GenS2{F0: 1}
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v45 := GenS1{F0: 42, F1: 1}
                                                if false {
                                                    _ = len(v3)
                                                }
                                                var v46 uint64 = 42
                                                _ = v44
                                                _ = v45
                                                _ = v46
                                            } else {
                                                var v47 int64 = v2
                                                switch ("foo bar" + ("" + (("hello" + "") + (("\n\t" + ("foo bar" + (("a" + (((((("foo bar" + "foo bar") + "hello") + ("\n\t" + ("hello" + "hello"))) + "") + "hello") + "hello")) + ""))) + "")))) {
                                                case "hello":
                                                    v6 = v42
                                                default:
                                                    fmt.Println(fi35)
                                                }
                                                var v48 uint = 65535
                                                _ = v47
                                                _ = v48
                                            }
                                        default:
                                            v6 += v6
                                        }
                                    case "\n\t":
                                        if (((((3.14159 / 1.5) < -2.25) || false) && false) || (!true)) {
                                            switch "hello" {
                                            case "":
                                                v42 = v42
                                            case "foo bar":
                                                v6 = (^(v6))
                                            }
                                        } else {
                                            v42 = (v6 | v6)
                                        }
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = (^(fi33))
                                    }
                                    v6 -= (v42 ^ (v42 % ((v6) | 1)))
                                    v49 := GenS1{F0: -65, F1: 422}
                                    var v50 int16 = -1
                                    v42 = v6
                                    _ = v42
                                    _ = v49
                                    _ = v50
                                } else {
                                    v51 := GenS1{F0: 1, F1: 32767}
                                    v4 = (v5 * v5)
                                    v2 -= (^(v2))
                                    _ = len(v1)
                                    v4 = 0
                                    var v52 int = 9223372036854775807
                                    _ = v51
                                    _ = v52
                                }
                                fmt.Println(v4)
                            } else {
                                var v53 uint32 = 4294967295
                                switch fi35 {
                                case 0:
                                    v2 -= -1
                                case 778:
                                    fmt.Println(v4, v2)
                                case -341:
                                    var v54 int = (fi33 * fi35)
                                    _ = v54
                                default:
                                    fmt.Println(v5, v53, fi35)
                                }
                                _ = v53
                            }
                            fmt.Println(fi35, v6)
                            if true {
                                v2 += (v2 >> (((v6 << ((1) & 63))) & 63))
                                if (-48 > 117) {
                                    v6 = (v6 / (((v6 | v6)) | 1))
                                } else {
                                    fmt.Println(fi33, v2, v2)
                                    var v55 uint8 = 255
                                    _ = v55
                                }
                                v1 = append(v1, (fi35 % ((fi33) | 1)))
                            } else {
                                v56 := GenS1{F0: 55, F1: 2}
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(v2)
                                case "hello":
                                    if false {
                                        var v57 uint32 = 237
                                        fmt.Println(fi33, v2)
                                        _ = v57
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                default:
                                    if false {
                                        var v58 int8 = 1
                                        v59 := GenS2{F0: 515}
                                        if false {
                                            var v60 uint = 42
                                            _ = v60
                                        }
                                        v61 := GenS1{F0: (-(v58)), F1: 317}
                                        _ = v58
                                        _ = v59
                                        _ = v61
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                }
                                v62 := GenS3{F0: false}
                                _ = v56
                                _ = v62
                            }
                            _ = len(v1)
                            v5 -= v4
                            _ = v41
                        } else {
                            var v63 uint8 = 127
                            var v64 uint8 = 255
                            _ = v63
                            _ = v64
                        }
                    case "foo bar":
                        v2 = v2
                    default:
                        switch fi33 {
                        case 0:
                            switch (fi33 & 0) {
                            case -9223372036854775808:
                                if true {
                                    var v65 uint8 = 2
                                    var v66 int16 = -432
                                    v66 = -449
                                    v4 = 42
                                    _ = v65
                                    _ = v66
                                } else {
                                    _ = len(v3)
                                    v2 = (v2 - (v2 % ((v2) | 1)))
                                    v3 = append(v3, "foo bar")
                                }
                            }
                        }
                        v2 -= (-(v2))
                        switch fi35 {
                        case -9223372036854775808:
                            var v67 bool = false
                            _ = v67
                        case 0:
                            v2 = -1
                        case 1:
                            var v68 int8 = 62
                            _ = v68
                        }
                    }
                    var v69 float32 = 1.5
                    v3 = append(v3, "a")
                    _ = v69
                }
                v1 = append(v1, (fi33 >> ((0) & 63)))
                v70 := GenS1{F0: -128, F1: 0}
                _ = fi35
                _ = v70
            }
        case "a":
            for fi71 := 0; fi71 < 12; fi71++ {
                v72 := GenS2{F0: 681}
                _ = fi71
                _ = v72
            }
        }
        v2 -= (v2 & (v2 + v2))
        _ = fi33
    }
    var v73 uint = v6
    var v74 uint = (v6 + v73)
    v75 := GenS4{F0: 364}
    for fi76 := 0; fi76 < 3; fi76++ {
        if (42 != 255) {
            for fi77 := 0; fi77 < 19; fi77++ {
                switch (fi77 << ((v6) & 63)) {
                case 952:
                    if (float32(1.5) < 0.1) {
                        v78 := GenS4{F0: 0}
                        switch fi76 {
                        case 4294967295:
                            v6 -= 2
                        case -2147483648:
                            if false {
                                v79 := GenS5{F0: -912}
                                v80 := GenS4{F0: 993}
                                _ = v79
                                _ = v80
                            }
                        case -1:
                            fmt.Println(v4)
                        }
                        switch fi76 {
                        case -673:
                            v81 := GenS5{F0: 1}
                            _ = v81
                        case 4294967295:
                            var v82 float64 = float64(1e10)
                            _ = v82
                        default:
                            v2 += (v2 / ((v2) | 1))
                        }
                        _ = v78
                    } else {
                        v73 = (v73 % (((^(v73))) | 1))
                        v83 := GenS3{F0: true}
                        var v84 float64 = 0.1
                        _ = v83
                        _ = v84
                    }
                case -181:
                    if ((190 == v73) && (true && (false && (v2 != v2)))) {
                        var v85 float32 = (((0.1 + -2.25) * 0.1) / 3.14159)
                        _ = v85
                    }
                case -9223372036854775808:
                    fmt.Println(fi77, fi77)
                default:
                    v86 := GenS2{F0: 4294967295}
                    _ = v86
                }
                v87 := GenS4{F0: 0}
                _ = fi77
                _ = v87
            }
            for fi88 := 0; fi88 < 5; fi88++ {
                v89 := GenS6{F0: (v2 << (((v6 - v73)) & 63)), F1: 32767, F2: v2}
                var v90 int = fi88
                v91 := GenS2{F0: 18446744073709551615}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
                _ = fi88
                _ = v89
                _ = v90
                _ = v91
            }
            v92 := GenS3{F0: false}
            _ = v92
        }
        if ("" != "foo bar") {
            var v93 uint32 = 1
            _ = len(v3)
            _ = v93
        } else {
            v1 = append(v1, -987)
        }
        var v94 uint32 = 42
        _ = fi76
        _ = v94
    }
    var v95 int64 = 9223372036854775807
    v6 = v6
    v74 += v6
    var v96 uint32 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
    _ = v75
    _ = v95
    _ = v96
}
