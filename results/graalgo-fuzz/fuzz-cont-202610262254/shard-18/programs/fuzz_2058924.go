package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 string; F1 int16; F2 int }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS8 struct { F0 uint8; F1 int64; F2 int16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint16; F1 int32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 string; F1 bool; F2 int16 }
func genFn0(p0 bool) uint16 {
    if ((p0 && false) && p0) {
        v7 := GenS1{F0: 765, F1: 0}
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 float32 = 3.14159
            var v10 uint8 = 251
            switch "hello" {
            case "a":
                v10 = (v10 * (v10 >> ((1) & 63)))
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    }
    v11 := GenS2{F0: "", F1: -1, F2: -9223372036854775808}
    var v12 uint64 = 334
    v12 -= v12
    fmt.Println(v12)
    _ = p0
    _ = v11
    _ = v12
    return 255
}
func genFn1(p0 uint16) uint16 {
    v13 := []int{-625, -9223372036854775808, 435}
    for fi14 := 0; fi14 < 20; fi14++ {
        p0 = 734
        _ = fi14
    }
    v15 := GenS1{F0: p0, F1: 1}
    if true {
        v16 := GenS3{F0: "", F1: 32768}
        if false {
            switch 127 {
            case 691:
                switch 9223372036854775807 {
                case -453:
                    var v17 int32 = -1
                    _ = v17
                }
            case 117:
                for fi18 := 0; fi18 < 7; fi18++ {
                    var v19 uint16 = p0
                    _ = fi18
                    _ = v19
                }
            case 32767:
                if len(v13) > 0 {
                    v13[0] = -940
                }
            default:
                v20 := GenS4{F0: 42}
                _ = v20
            }
            var v21 string = (("hello" + "hello") + ("\n\t" + ("hello" + ("foo bar" + "\n\t"))))
            switch 2147483647 {
            case 948:
                if (0 <= 1) {
                    for fi22 := 0; fi22 < 6; fi22++ {
                        v23 := GenS2{F0: "a", F1: -32768, F2: (fi22 + fi22)}
                        _ = fi22
                        _ = v23
                    }
                }
            case 2147483647:
                var v24 uint16 = (^(p0))
                _ = v24
            }
            _ = v21
        } else {
            var v25 int16 = 42
            _ = v25
        }
        v26 := GenS2{F0: "foo bar", F1: 0, F2: 127}
        _ = v16
        _ = v26
    } else {
        for fi27 := 0; fi27 < 6; fi27++ {
            p0 -= genFn0(true)
            _ = fi27
        }
        v13 = append(v13, 32768)
    }
    var v28 int64 = -9223372036854775808
    _ = p0
    _ = v13
    _ = v15
    _ = v28
    return p0
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 3.14159
    v3 := make(map[string]string)
    var v4 uint64 = 42
    var v5 bool = true
    v6 := []int{}
    if (-32768 > 42) {
        switch ("\n\t" + "") {
        case "foo bar":
            for fi29 := 0; fi29 < 20; fi29++ {
                fmt.Println(v5, v4)
                _ = fi29
            }
        case "a":
            var v30 uint8 = 148
            _ = v30
        case "hello":
            v31 := v3[("\n\t" + "foo bar")]
            _ = v31
        }
    }
    fmt.Println(v5, v4, v2)
    switch ("\n\t" + "foo bar") {
    case "a":
        v2 = (3.14159 / float64(-2.25))
    case "":
        if ((!v5) && false) {
            v5 = (false && v5)
            var v32 int8 = 112
            switch 164 {
            case 32768:
                v33 := GenS2{F0: (("\n\t" + "foo bar") + "foo bar"), F1: 0, F2: 127}
                _ = v33
            case 4294967295:
                v5 = v5
            }
            _ = v32
        } else {
            if v5 {
                for fi34 := 0; fi34 < 7; fi34++ {
                    for fi35 := 0; fi35 < 11; fi35++ {
                        if (-243 <= fi35) {
                            var v36 uint16 = genFn0((!false))
                            v6 = append(v6, int(v4))
                            v37 := GenS5{F0: 42, F1: 127, F2: (v36 | (^(v36)))}
                            _ = v36
                            _ = v37
                        }
                        var v38 float64 = v2
                        _ = fi35
                        _ = v38
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi34
                }
                var v39 int = 0
                _ = v39
            } else {
                for fi40 := 0; fi40 < 12; fi40++ {
                    var v41 float32 = 3.14159
                    _ = fi40
                    _ = v41
                }
                fmt.Println(v4, v5)
            }
            fmt.Println(v5)
        }
    default:
        for fi42 := 0; fi42 < 8; fi42++ {
            switch "" {
            case "a":
                switch "" {
                case "":
                    v3[("hello" + "hello")] = "\n\t"
                case "foo bar":
                    for fi43 := 0; fi43 < 19; fi43++ {
                        switch (fi43 - 32767) {
                        case 32768:
                            if v5 {
                                v2 = float64(1.5)
                            } else {
                                _ = len(v6)
                                v1[(fi43 & fi42)] = ""
                            }
                        default:
                            v4 -= v4
                        }
                        v5 = v5
                        fmt.Println(v5)
                        if (4294967295 <= 65535) {
                            if (261 < 1) {
                                if ((0.0 - -2.25) == 0.1) {
                                    if (fi43 < 128) {
                                        var v44 uint8 = 42
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        var v45 uint64 = 0
                                        _ = v44
                                        _ = v45
                                    } else {
                                        if (1e10 == 1.5) {
                                            v2 += v2
                                            if (51 < 2147483648) {
                                                v46 := GenS6{F0: 0}
                                                fmt.Println(fi43, fi42)
                                                _ = v46
                                            } else {
                                                var v47 float32 = 3.14159
                                                if true {
                                                    v48 := GenS2{F0: "\n\t", F1: 1000, F2: (fi42 << ((661) & 63))}
                                                    var v49 int64 = -9223372036854775808
                                                    v50 := GenS2{F0: "hello", F1: -318, F2: fi43}
                                                    _ = v48
                                                    _ = v49
                                                    _ = v50
                                                }
                                                _ = v47
                                            }
                                            v5 = v5
                                        } else {
                                            if (v5 && (-2147483648 > -782)) {
                                                v51 := GenS1{F0: 892, F1: 1}
                                                _ = v51
                                            } else {
                                                v5 = v5
                                            }
                                            switch fi43 {
                                            case 4294967295:
                                                v2 = float64(3.14159)
                                            }
                                            v4 = 890
                                        }
                                        v52 := GenS3{F0: ("hello" + "\n\t"), F1: fi43}
                                        _ = v52
                                    }
                                    switch (fi42 ^ (-(fi43))) {
                                    case 2147483647:
                                        switch 32768 {
                                        case 0:
                                            if v5 {
                                                switch 32767 {
                                                case -2147483648:
                                                    v5 = (false || v5)
                                                case 2147483647:
                                                    var v53 int = (-(fi43))
                                                    _ = v53
                                                default:
                                                    var v54 int64 = -1
                                                    _ = v54
                                                }
                                                v5 = (!true)
                                                v2 = 1.5
                                            }
                                        case 519:
                                            fmt.Println(v4, v2)
                                        default:
                                            v55 := GenS6{F0: -917}
                                            _ = v55
                                        }
                                    }
                                    fmt.Println(fi42, v4)
                                } else {
                                    v56 := GenS1{F0: genFn1(259), F1: 255}
                                    var v57 int64 = 9223372036854775807
                                    var v58 int = fi43
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                }
                                v59 := GenS5{F0: 128, F1: 127, F2: 705}
                                _ = v59
                            } else {
                                switch "hello" {
                                case "":
                                    if true {
                                        if (0 >= 42) {
                                            v5 = (42 <= 0)
                                            v4 = 4294967295
                                        }
                                        var v60 int32 = 42
                                        v5 = false
                                        _ = v60
                                    }
                                }
                                v4 += v4
                                var v61 int64 = 42
                                _ = v61
                            }
                            var v62 int8 = -1
                            if v5 {
                                v4 += v4
                            }
                            _ = v62
                        } else {
                            fmt.Println(v2, fi42, v5)
                        }
                        var v63 float32 = ((float32((float64(3.14159) - (3.14159 * 1e10))) - ((3.14159 + (0.0 + 1.5)) * 0.1)) / 1.0)
                        _ = fi43
                        _ = v63
                    }
                case "hello":
                    var v64 int = (fi42 ^ (fi42 | fi42))
                    _ = v64
                default:
                    v65 := GenS4{F0: 2}
                    _ = v65
                }
            }
            if ("" <= ("\n\t" + ("" + "a"))) {
                if ((v4 >> ((2) & 63)) > v4) {
                    var v66 uint = 42
                    _ = v66
                } else {
                    v4 += v4
                    switch (("" + "") + "hello") {
                    case "":
                        if v5 {
                            v67 := GenS5{F0: 127, F1: -128, F2: 255}
                            var v68 int8 = -17
                            if false {
                                if (32767 <= 152) {
                                    var v69 bool = (324 >= 0)
                                    v5 = (17 < 1)
                                    v5 = (true || v69)
                                    _ = v69
                                }
                            } else {
                                fmt.Println(v4, v4)
                            }
                            _ = v67
                            _ = v68
                        }
                    case "a":
                        var v70 uint32 = 203
                        _ = v70
                    case "foo bar":
                        var v71 float32 = float32((((-2.25 + v2) / -2.25) + (3.14159 + v2)))
                        _ = v71
                    default:
                        delete(v3, "a")
                    }
                    v5 = (v5 && true)
                }
                v72 := GenS7{F0: int8(v4), F1: 211, F2: 42}
                v2 = float64(((-2.25 + 0.1) - -2.25))
                _ = v72
            }
            for fi73 := 0; fi73 < 16; fi73++ {
                fmt.Println(v5, v4, v2)
                _ = fi73
            }
            _ = fi42
        }
    }
    if v5 {
        if v5 {
            v1[9223372036854775807] = ("foo bar" + "a")
            var v74 uint16 = 640
            _ = v74
        } else {
            switch "\n\t" {
            case "\n\t":
                for fi75 := 0; fi75 < 10; fi75++ {
                    v2 -= v2
                    v76 := GenS8{F0: 1, F1: -9223372036854775808, F2: 169}
                    _ = fi75
                    _ = v76
                }
            case "":
                if v5 {
                    switch 0 {
                    case -9223372036854775808:
                        var v77 uint8 = 255
                        _ = v77
                    case 2147483647:
                        var v78 uint16 = 32767
                        _ = v78
                    }
                }
            case "a":
                for fi79 := 0; fi79 < 14; fi79++ {
                    delete(v3, "foo bar")
                    _ = fi79
                }
            default:
                v5 = (0 <= 0)
            }
            fmt.Println(v5, v2, v2)
            v5 = (!false)
        }
        var v80 float32 = 1e10
        v81 := GenS9{F0: 0}
        _ = v80
        _ = v81
    }
    switch "" {
    case "\n\t":
        if true {
            v82 := GenS7{F0: -1, F1: 0, F2: 1}
            _ = v82
        }
    case "hello":
        v4 -= (v4 - 1)
    }
    if ("" == "\n\t") {
        switch -2147483648 {
        case -886:
            if false {
                _ = len(v3)
                var v83 uint16 = 560
                switch 127 {
                case 607:
                    _ = len(v3)
                case -246:
                    var v84 uint32 = 1
                    _ = v84
                }
                _ = v83
            } else {
                _ = len(v6)
                if v5 {
                    fmt.Println(v5)
                    for fi85 := 0; fi85 < 19; fi85++ {
                        switch 567 {
                        case 127:
                            v86 := GenS4{F0: 821}
                            _ = v86
                        case 0:
                            fmt.Println(v2)
                        case 42:
                            v87 := GenS10{F0: genFn0((65535 < 0)), F1: 1}
                            _ = v87
                        default:
                            v88 := GenS11{F0: 511}
                            _ = v88
                        }
                        v89 := GenS12{F0: ("foo bar" + "foo bar"), F1: true, F2: 699}
                        _ = fi85
                        _ = v89
                    }
                }
            }
        case 32767:
            v4 = v4
        }
        for fi90 := 0; fi90 < 13; fi90++ {
            fmt.Println(v4, v5)
            _ = fi90
        }
    } else {
        if len(v6) > 0 {
            v6[0] = -116
        }
        for fi91 := 0; fi91 < 14; fi91++ {
            var v92 int8 = 100
            v5 = (18446744073709551615 == v4)
            _ = fi91
            _ = v92
        }
    }
    v2 -= (float64(1.5) - float64(-2.25))
    v2 += 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
