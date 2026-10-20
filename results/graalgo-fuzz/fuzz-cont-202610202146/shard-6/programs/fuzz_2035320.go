package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 uint32; F1 uint16 }
type GenS8 struct { F0 int16 }

func main() {
    var v1 uint32 = 4294967295
    var v2 int32 = -1
    var v3 bool = (238 < 42)
    var v4 uint64 = 42
    v5 := []string{}
    v6 := make(map[int]string)
    if false {
        if ((v4 >= v4) && v3) {
            _ = len(v6)
            fmt.Println(v4, v4, v1)
            var v7 float64 = 1.5
            _ = v7
        }
        switch "hello" {
        case "\n\t":
            if v3 {
                var v8 int16 = 1
                switch "hello" {
                case "":
                    if false {
                        switch -929 {
                        case -2147483648:
                            switch ("foo bar" + "a") {
                            case "":
                                fmt.Println(v8, v4)
                            case "a":
                                v5 = append(v5, "a")
                            case "foo bar":
                                var v9 int32 = (v2 % ((-1) | 1))
                                _ = v9
                            default:
                                v8 -= (^(v8))
                            }
                        case 194:
                            if false {
                                _ = len(v5)
                            } else {
                                var v10 uint = 18446744073709551615
                                for fi11 := 0; fi11 < 1; fi11++ {
                                    if v3 {
                                        delete(v6, (fi11 / ((fi11) | 1)))
                                        var v12 float64 = -2.25
                                        _ = v12
                                    }
                                    _ = fi11
                                }
                                var v13 uint16 = uint16(v4)
                                _ = v10
                                _ = v13
                            }
                        }
                    }
                case "\n\t":
                    for fi14 := 0; fi14 < 8; fi14++ {
                        switch fi14 {
                        case 2147483647:
                            if (1.5 != -2.25) {
                                switch (fi14 % (((^(fi14))) | 1)) {
                                case 32768:
                                    for fi15 := 0; fi15 < 10; fi15++ {
                                        if v3 {
                                            var v16 float32 = 0.0
                                            v17 := GenS1{F0: 823, F1: 42}
                                            v18 := GenS1{F0: -1, F1: 4294967295}
                                            _ = v16
                                            _ = v17
                                            _ = v18
                                        }
                                        var v19 int = int(v8)
                                        _ = fi15
                                        _ = v19
                                    }
                                }
                            }
                        case 131:
                            v4 -= v4
                        case 32768:
                            v3 = v3
                        }
                        _ = fi14
                    }
                default:
                    for fi20 := 0; fi20 < 10; fi20++ {
                        var v21 string = ""
                        _ = fi20
                        _ = v21
                    }
                }
                _ = v8
            } else {
                for fi22 := 0; fi22 < 11; fi22++ {
                    var v23 float32 = float32(3.14159)
                    var v24 int32 = v2
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
        case "hello":
            if v3 {
                var v25 uint8 = 77
                _ = v25
            } else {
                switch "a" {
                case "":
                    _ = len(v5)
                default:
                    if true {
                        for fi26 := 0; fi26 < 5; fi26++ {
                            v1 += 4294967295
                            v27 := GenS1{F0: -1, F1: 65535}
                            _ = fi26
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 11; fi28++ {
                            v2 = v2
                            v29 := GenS2{F0: v1, F1: 65535}
                            v30 := v6[fi28]
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    }
                }
                for fi31 := 0; fi31 < 10; fi31++ {
                    v2 -= (v2 % ((-1) | 1))
                    for fi32 := 0; fi32 < 1; fi32++ {
                        fmt.Println(v3, v2, fi32)
                        if ("hello" != "") {
                            var v33 uint = 22
                            fmt.Println(fi31, v1, v3)
                            var v34 int32 = (v2 & v2)
                            _ = v33
                            _ = v34
                        } else {
                            var v35 float64 = (0.1 - -2.25)
                            _ = v35
                        }
                        var v36 int32 = v2
                        _ = fi32
                        _ = v36
                    }
                    _ = fi31
                }
                if (v3 && v3) {
                    var v37 uint8 = 42
                    for fi38 := 0; fi38 < 2; fi38++ {
                        if (4294967295 <= (v4 % ((0) | 1))) {
                            for fi39 := 0; fi39 < 8; fi39++ {
                                v4 -= (v4 % ((275) | 1))
                                var v40 float32 = 0.0
                                var v41 uint16 = 42
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                        } else {
                            var v42 uint64 = (^(v4))
                            var v43 uint64 = v42
                            if (2 == 65535) {
                                v44 := GenS3{F0: 0}
                                v2 -= (v2 + 2147483647)
                                _ = v44
                            } else {
                                v42 += v42
                                v43 = (v42 >> ((uint(v4)) & 63))
                                var v45 uint64 = (v4 * (v4 - (v43 % ((v43) | 1))))
                                _ = v45
                            }
                            _ = v42
                            _ = v43
                        }
                        _ = fi38
                    }
                    var v46 uint64 = v4
                    _ = v37
                    _ = v46
                } else {
                    v47 := v6[-607]
                    v3 = v3
                    for fi48 := 0; fi48 < 11; fi48++ {
                        for fi49 := 0; fi49 < 13; fi49++ {
                            fmt.Println(v4)
                            var v50 int8 = 81
                            _ = fi49
                            _ = v50
                        }
                        _ = fi48
                    }
                    _ = v47
                }
            }
        case "a":
            v51 := GenS2{F0: v1, F1: 42}
            _ = v51
        default:
            v4 += 1
        }
        for fi52 := 0; fi52 < 17; fi52++ {
            if v3 {
                for fi53 := 0; fi53 < 13; fi53++ {
                    var v54 int16 = 768
                    if true {
                        v1 += uint32(v4)
                        if (!false) {
                            if ((((3.14159 / float64(1e10)) * float64((-2.25 / 1.0))) + 0.1) > 1e10) {
                                if false {
                                    v55 := GenS4{F0: 42, F1: (v2 | -145)}
                                    v56 := GenS5{F0: 9223372036854775807, F1: (v54 ^ v54)}
                                    _ = v55
                                    _ = v56
                                } else {
                                    var v57 float64 = (-2.25 - -2.25)
                                    v4 = (v4 + (v4 - 2))
                                    _ = v57
                                }
                                switch (fi52 >> ((18446744073709551615) & 63)) {
                                case -2147483648:
                                    var v58 uint = 4294967295
                                    _ = v58
                                case -197:
                                    _ = len(v6)
                                case 95:
                                    v59 := GenS3{F0: 577}
                                    _ = v59
                                default:
                                    var v60 uint = 2
                                    _ = v60
                                }
                                v61 := GenS3{F0: 42}
                                fmt.Println(v3)
                                _ = v61
                            } else {
                                var v62 int8 = -22
                                if false {
                                    delete(v6, 555)
                                    v2 -= v2
                                    v63 := GenS2{F0: 2, F1: 970}
                                    _ = v63
                                }
                                v4 -= uint64(v1)
                                _ = v62
                            }
                            var v64 int = fi53
                            if ((fi52 / (((fi53 & (v64 & (fi52 * (v64 % ((fi52) | 1)))))) | 1)) != fi53) {
                                fmt.Println(v2)
                                switch ("foo bar" + "foo bar") {
                                case "":
                                    v65 := GenS6{F0: 65535, F1: ("" + "foo bar")}
                                    _ = v65
                                case "foo bar":
                                    _ = len(v5)
                                case "a":
                                    var v66 uint8 = 128
                                    _ = v66
                                default:
                                    v4 += 570
                                }
                                v1 -= v1
                            } else {
                                switch (("hello" + ("\n\t" + "hello")) + ("\n\t" + ("" + ("a" + "foo bar")))) {
                                case "a":
                                    var v67 int32 = 1
                                    _ = v67
                                }
                            }
                            _ = v64
                        } else {
                            fmt.Println(v54, v54, fi53)
                            var v68 int32 = (v2 & (v2 + v2))
                            _ = v68
                        }
                    } else {
                        fmt.Println(fi53)
                        if (false || false) {
                            v54 = 1
                        }
                        v1 -= (v1 / (((v1 % ((v1) | 1))) | 1))
                        var v69 bool = (!(18446744073709551615 >= (v4 & v4)))
                        _ = v69
                    }
                    _ = fi53
                    _ = v54
                }
            }
            for fi70 := 0; fi70 < 14; fi70++ {
                v4 = v4
                v1 -= v1
                _ = fi70
            }
            for fi71 := 0; fi71 < 12; fi71++ {
                v4 += 1
                _ = fi71
            }
            _ = fi52
        }
    } else {
        v72 := GenS3{F0: uint16(v1)}
        v2 = (v2 / (((^(v2))) | 1))
        _ = v72
    }
    for fi73 := 0; fi73 < 19; fi73++ {
        var v74 uint = 255
        fmt.Println(v74, v74)
        v5 = append(v5, ("hello" + ("a" + "\n\t")))
        _ = fi73
        _ = v74
    }
    switch 417 {
    case 42:
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
    default:
        v3 = v3
    }
    switch ("a" + ("" + "a")) {
    case "a":
        if (128 == -416) {
            var v75 int64 = 0
            _ = v75
        }
    case "hello":
        v3 = true
    case "foo bar":
        if true {
            var v76 bool = (!(707 == 32768))
            for fi77 := 0; fi77 < 1; fi77++ {
                for fi78 := 0; fi78 < 18; fi78++ {
                    var v79 float64 = 0.1
                    if (65535 != 1) {
                        var v80 float32 = ((((((1e10 + (0.0 + 0.0)) + 1e10) + 0.1) + 1.5) / (1e10 / 0.1)) / 1.0)
                        if (!true) {
                            v1 = 42
                            var v81 uint16 = 255
                            v82 := GenS7{F0: v1, F1: 804}
                            _ = v81
                            _ = v82
                        }
                        v83 := GenS7{F0: (^(v1)), F1: 2}
                        _ = v80
                        _ = v83
                    } else {
                        var v84 float64 = 1e10
                        _ = v84
                    }
                    var v85 int16 = 32767
                    var v86 float64 = (v79 * ((-2.25 / 1.5) / v79))
                    _ = fi78
                    _ = v79
                    _ = v85
                    _ = v86
                }
                if (!((float32(0.0) >= 0.1) || false)) {
                    v87 := GenS2{F0: v1, F1: 259}
                    for fi88 := 0; fi88 < 13; fi88++ {
                        v4 = v4
                        var v89 uint32 = (^(v1))
                        var v90 string = "foo bar"
                        fmt.Println(v3, v76, v90)
                        _ = fi88
                        _ = v89
                        _ = v90
                    }
                    _ = v87
                }
                _ = fi77
            }
            v5 = append(v5, (("hello" + "") + "a"))
            _ = v76
        } else {
            var v91 uint32 = (v1 << ((965) & 63))
            var v92 float64 = (0.0 - 0.1)
            _ = v91
            _ = v92
        }
    }
    v93 := GenS4{F0: 0, F1: v2}
    var v94 int16 = 1000
    v5 = append(v5, "\n\t")
    switch "foo bar" {
    case "\n\t":
        for fi95 := 0; fi95 < 19; fi95++ {
            var v96 int64 = int64(v4)
            var v97 string = "\n\t"
            _ = fi95
            _ = v96
            _ = v97
        }
    case "a":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "foo bar":
        v6[2147483647] = ""
    default:
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
    }
    if v3 {
        v98 := GenS8{F0: v94}
        _ = v98
    } else {
        v99 := GenS4{F0: 255, F1: v2}
        _ = v99
    }
    if ((true || v3) || v3) {
        switch ("" + "foo bar") {
        case "\n\t":
            if v3 {
                v100 := GenS3{F0: 1}
                _ = len(v6)
                fmt.Println(v4)
                _ = v100
            }
        case "hello":
            var v101 string = ""
            _ = v101
        }
        fmt.Println(v94)
    } else {
        if v3 {
            var v102 float64 = ((0.0 * 1e10) / (3.14159 * (0.0 - (1.5 / float64(-2.25)))))
            v1 = uint32(v4)
            _ = v102
        }
    }
    switch "a" {
    case "foo bar":
        for fi103 := 0; fi103 < 3; fi103++ {
            fmt.Println(v3)
            _ = fi103
        }
    }
    var v104 int = 4294967295
    for fi105 := 0; fi105 < 11; fi105++ {
        var v106 uint16 = 255
        v6[(v104 / ((-218) | 1))] = "\n\t"
        _ = fi105
        _ = v106
    }
    fmt.Println(v1, v104)
    var v107 uint = 630
    v2 -= v2
    v4 += 0
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v107:", v107)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v94:", v94)
    _ = v1
    _ = v104
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v93
    _ = v94
}
