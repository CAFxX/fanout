package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 uint64; F1 string; F2 bool }
type GenS6 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int32; F1 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int32; F1 int8; F2 uint64 }
type GenS11 struct { F0 uint64; F1 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 0
    v3 := make(map[string]int)
    var v4 int32 = -8
    v5 := []string{}
    switch "foo bar" {
    case "\n\t":
        var v6 float64 = 0.1
        _ = v6
    case "foo bar":
        for fi7 := 0; fi7 < 17; fi7++ {
            fmt.Println(v4, v4, v2)
            _ = fi7
        }
    case "":
        var v8 string = ""
        _ = v8
    default:
        v4 -= v4
    }
    var v9 int32 = -958
    fmt.Println(v9, v4, v4)
    if true {
        if true {
            if (!((^(v4)) == (v4 / ((v9) | 1)))) {
                var v10 uint16 = 2
                if false {
                    if false {
                        for fi11 := 0; fi11 < 5; fi11++ {
                            var v12 uint32 = 2147483647
                            _ = fi11
                            _ = v12
                        }
                    } else {
                        var v13 int16 = 42
                        _ = v13
                    }
                    v9 = (^(v9))
                    var v14 int64 = 42
                    _ = v14
                } else {
                    fmt.Println(v2, v4)
                    switch 42 {
                    case 1:
                        var v15 uint64 = 810
                        _ = v15
                    case 2147483647:
                        var v16 int32 = v9
                        _ = v16
                    }
                }
                for fi17 := 0; fi17 < 20; fi17++ {
                    _ = len(v3)
                    switch (fi17 / ((fi17) | 1)) {
                    case 1:
                        var v18 float64 = (1e10 + 1e10)
                        _ = v18
                    case -1:
                        var v19 uint64 = 0
                        _ = v19
                    case 128:
                        v20 := GenS1{F0: fi17}
                        _ = v20
                    default:
                        v2 += 800
                    }
                    _ = fi17
                }
                _ = v10
            }
            var v21 uint32 = 65535
            _ = v21
        } else {
            if false {
                if (true || ((18446744073709551615 < 4294967295) && false)) {
                    v3[""] = 127
                }
                fmt.Println(v9)
                if (((3.14159 * (((1.5 / ((float32((0.0 - 0.1)) * float32((1e10 - 1.5))) + (1.5 + (-2.25 - 1e10)))) - -2.25) - 0.1)) - (float32((3.14159 * 1e10)) * (((3.14159 - 0.0) / 1e10) - (0.1 / 1.5)))) == 0.1) {
                    v2 += (^(v2))
                } else {
                    var v22 uint32 = 119
                    _ = v22
                }
            } else {
                if true {
                    if false {
                        switch "" {
                        case "hello":
                            fmt.Println(v2, v4)
                        case "foo bar":
                            var v23 uint16 = v2
                            _ = v23
                        default:
                            _ = len(v1)
                        }
                    }
                    v9 += (v4 & -101)
                    for fi24 := 0; fi24 < 7; fi24++ {
                        v9 += (^(v9))
                        fmt.Println(fi24, v9, v4)
                        _ = fi24
                    }
                } else {
                    if false {
                        for fi25 := 0; fi25 < 3; fi25++ {
                            for fi26 := 0; fi26 < 9; fi26++ {
                                var v27 uint16 = (v2 * v2)
                                switch "\n\t" {
                                case "\n\t":
                                    v28 := v1[4294967295]
                                    _ = v28
                                case "a":
                                    var v29 float64 = 3.14159
                                    _ = v29
                                case "":
                                    if true {
                                        if len(v5) > 0 {
                                            v5[0] = ("a" + (("a" + (((("" + "a") + "") + ("" + "a")) + "foo bar")) + "foo bar"))
                                        }
                                        v2 = 255
                                    } else {
                                        switch ("hello" + "hello") {
                                        case "\n\t":
                                            if (4294967295 < 0) {
                                                v9 -= (v9 - v4)
                                                v30 := GenS1{F0: (fi26 + -526)}
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                _ = v30
                                            }
                                        case "hello":
                                            v2 -= v2
                                        case "foo bar":
                                            v4 += 567
                                        default:
                                            v31 := GenS2{F0: 255}
                                            fmt.Println(v9, fi25, fi26)
                                            switch 2147483647 {
                                            case 32767:
                                                fmt.Println(v2, v2, fi25)
                                            case 42:
                                                if true {
                                                    fmt.Println(fi25, v9, v2)
                                                } else {
                                                    if false {
                                                        v4 -= v4
                                                        v2 = v2
                                                        delete(v1, 1)
                                                    }
                                                    if len(v5) > 0 {
                                                        v5[0] = ("a" + (("hello" + "") + ("" + ((("\n\t" + "hello") + ((("" + "foo bar") + "\n\t") + ("hello" + ("a" + "hello")))) + "hello"))))
                                                    }
                                                }
                                            case 903:
                                                v32 := GenS2{F0: 18446744073709551615}
                                                _ = v32
                                            default:
                                                v33 := GenS3{F0: 42}
                                                _ = v33
                                            }
                                            var v34 int16 = 135
                                            _ = v31
                                            _ = v34
                                        }
                                    }
                                default:
                                    var v35 int32 = (v4 / ((v4) | 1))
                                    _ = v35
                                }
                                _ = len(v5)
                                _ = fi26
                                _ = v27
                            }
                            if true {
                                var v36 int32 = (v4 + v9)
                                _ = v36
                            } else {
                                var v37 uint32 = 2
                                _ = v37
                            }
                            _ = fi25
                        }
                        for fi38 := 0; fi38 < 14; fi38++ {
                            var v39 uint8 = 127
                            switch (((("foo bar" + ("foo bar" + ("hello" + "foo bar"))) + "\n\t") + "") + "") {
                            case "\n\t":
                                v2 += (v2 + 2)
                            case "a":
                                switch ((("\n\t" + "") + ("foo bar" + "a")) + "") {
                                case "":
                                    for fi40 := 0; fi40 < 1; fi40++ {
                                        if (!(!false)) {
                                            v4 = (v4 ^ v4)
                                        } else {
                                            var v41 uint16 = 0
                                            v9 = (v4 * (^(v9)))
                                            if false {
                                                if false {
                                                    v9 = (v4 & (v9 & (v4 / (((v4 | (v9 - v9))) | 1))))
                                                    _ = len(v3)
                                                    v42 := GenS4{F0: (v2 * (v2 % ((344) | 1))), F1: (fi38 * fi40)}
                                                    _ = v42
                                                } else {
                                                    fmt.Println(fi38)
                                                    if ((-(v4)) == v9) {
                                                        v43 := GenS5{F0: 810, F1: "hello", F2: false}
                                                        _ = v43
                                                    } else {
                                                        if (!(-1 != -1)) {
                                                            var v44 uint16 = 2
                                                            _ = v44
                                                        }
                                                        switch 9223372036854775807 {
                                                        case -586:
                                                            switch (fi38 & fi40) {
                                                            case -24:
                                                                if (9 != 255) {
                                                                    var v45 int64 = -1
                                                                    _ = v45
                                                                }
                                                            default:
                                                                if true {
                                                                    if false {
                                                                        var v46 string = ""
                                                                        switch (-(fi40)) {
                                                                        case -2147483648:
                                                                            v47 := GenS6{F0: 42, F1: 0, F2: v4}
                                                                            v46 = v46
                                                                            _ = v47
                                                                        case -990:
                                                                            v48 := GenS2{F0: 2}
                                                                            var v49 string = ("\n\t" + "hello")
                                                                            _ = v48
                                                                            _ = v49
                                                                        case 2147483647:
                                                                            if (fi38 > (fi40 + (fi38 * fi38))) {
                                                                                fmt.Println(fi40)
                                                                            }
                                                                            v50 := GenS7{F0: (-2.25 >= 3.14159)}
                                                                            _ = v50
                                                                        default:
                                                                            v1[655] = (!true)
                                                                        }
                                                                        var v51 int64 = -9223372036854775808
                                                                        var v52 int16 = -32768
                                                                        var v53 uint32 = 536
                                                                        _ = v46
                                                                        _ = v51
                                                                        _ = v52
                                                                        _ = v53
                                                                    } else {
                                                                        var v54 uint32 = 246
                                                                        fmt.Println(v41, v4)
                                                                        v9 -= (v4 + v4)
                                                                        var v55 uint = 1
                                                                        var v56 uint32 = v54
                                                                        _ = v54
                                                                        _ = v55
                                                                        _ = v56
                                                                    }
                                                                    v57 := GenS8{F0: (v9 * v4), F1: 100}
                                                                    if true {
                                                                        fmt.Println(fi38, fi38, v9)
                                                                    } else {
                                                                        fmt.Println(v9, v41)
                                                                    }
                                                                    _ = v57
                                                                } else {
                                                                    if true {
                                                                        v58 := GenS9{F0: 255}
                                                                        _ = len(v3)
                                                                        if false {
                                                                            _ = len(v3)
                                                                        }
                                                                        var v59 uint64 = 42
                                                                        _ = v58
                                                                        _ = v59
                                                                    }
                                                                    if true {
                                                                        v9 = v9
                                                                        var v60 uint = 2
                                                                        _ = v60
                                                                    } else {
                                                                        var v61 uint = 131
                                                                        var v62 uint64 = 4294967295
                                                                        fmt.Println(v62, v2)
                                                                        _ = v61
                                                                        _ = v62
                                                                    }
                                                                    fmt.Println(v41)
                                                                }
                                                            }
                                                        case 1:
                                                            var v63 int = (fi40 << ((255) & 63))
                                                            _ = v63
                                                        }
                                                        var v64 string = "hello"
                                                        _ = v64
                                                    }
                                                    v4 = v9
                                                    var v65 int = (^(fi38))
                                                    _ = v65
                                                }
                                                fmt.Println(v4, v41)
                                            }
                                            _ = v41
                                        }
                                        var v66 int8 = 127
                                        switch "foo bar" {
                                        case "\n\t":
                                            if ((^(v39)) > 128) {
                                                _ = len(v1)
                                            } else {
                                                v67 := GenS5{F0: 4294967295, F1: ("hello" + "foo bar"), F2: false}
                                                fmt.Println(fi40, v39)
                                                _ = v67
                                            }
                                        }
                                        switch (fi38 % ((4294967295) | 1)) {
                                        case -1:
                                            fmt.Println(fi40)
                                        default:
                                            switch 274 {
                                            case 190:
                                                fmt.Println(v39, v66)
                                            default:
                                                v39 += 111
                                            }
                                        }
                                        v66 += v66
                                        _ = fi40
                                        _ = v66
                                    }
                                default:
                                    v9 = -2147483648
                                }
                            default:
                                v4 = (v4 * v4)
                            }
                            _ = fi38
                            _ = v39
                        }
                    } else {
                        switch 4294967295 {
                        case -599:
                            for fi68 := 0; fi68 < 16; fi68++ {
                                if true {
                                    v69 := GenS4{F0: 32767, F1: fi68}
                                    v4 = v4
                                    var v70 int16 = -1
                                    _ = v69
                                    _ = v70
                                } else {
                                    _ = len(v3)
                                }
                                switch 32767 {
                                case -9223372036854775808:
                                    if (!false) {
                                        for fi71 := 0; fi71 < 3; fi71++ {
                                            fmt.Println(v9, fi68)
                                            if true {
                                                v1[0] = true
                                                var v72 int = fi71
                                                v73 := GenS4{F0: (v2 & 32767), F1: fi71}
                                                if (2147483647 > v4) {
                                                    if false {
                                                        if (false && false) {
                                                            v2 -= (^(v2))
                                                            var v74 float32 = 1.5
                                                            v3["foo bar"] = (fi71 | 42)
                                                            _ = v74
                                                        }
                                                        var v75 uint64 = 843
                                                        var v76 uint64 = (v75 / ((v75) | 1))
                                                        v4 += (v9 * int32(fi68))
                                                        var v77 int = 127
                                                        _ = v75
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    var v78 uint64 = 4294967295
                                                    var v79 string = ("foo bar" + ("a" + (("a" + "\n\t") + "hello")))
                                                    _ = v78
                                                    _ = v79
                                                }
                                                var v80 int64 = -1
                                                v9 = 2147483647
                                                var v81 float32 = 0.0
                                                v82 := GenS10{F0: -1, F1: 42, F2: 4294967295}
                                                _ = v72
                                                _ = v73
                                                _ = v80
                                                _ = v81
                                                _ = v82
                                            } else {
                                                v4 = (v9 ^ (^(v4)))
                                            }
                                            _ = fi71
                                        }
                                    }
                                case 9223372036854775807:
                                    v83 := GenS8{F0: (-(v9)), F1: 100}
                                    _ = v83
                                }
                                if false {
                                    switch "\n\t" {
                                    case "":
                                        var v84 uint32 = 2147483648
                                        _ = v84
                                    default:
                                        fmt.Println(v4, v2)
                                    }
                                    var v85 uint64 = 18446744073709551615
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v85
                                }
                                _ = fi68
                            }
                        case -2147483648:
                            var v86 int32 = (v4 & (v9 / ((42) | 1)))
                            _ = v86
                        case -1:
                            v2 += (v2 - 42)
                        }
                    }
                    if false {
                        if true {
                            var v87 int16 = 1
                            v88 := GenS11{F0: 398, F1: (v2 % ((v2) | 1))}
                            _ = v87
                            _ = v88
                        } else {
                            var v89 uint32 = 906
                            v90 := v3["hello"]
                            _ = v89
                            _ = v90
                        }
                        var v91 int64 = 9223372036854775807
                        _ = v91
                    }
                }
                fmt.Println(v9, v9, v2)
                var v92 bool = true
                _ = v92
            }
            v4 -= 42
        }
    } else {
        v9 -= (v9 >> ((65535) & 63))
    }
    switch 128 {
    case 0:
        v2 = (v2 & (v2 >> ((255) & 63)))
    }
    v9 = 1
    var v93 uint32 = 0
    var v94 float32 = 1.5
    for fi95 := 0; fi95 < 12; fi95++ {
        var v96 int32 = (v9 & (v4 + -2147483648))
        _ = fi95
        _ = v96
    }
    var v97 int16 = 32767
    if (!(127 != 128)) {
        v5 = append(v5, "\n\t")
    }
    var v98 string = ("a" + "a")
    v99 := GenS6{F0: (^(v93)), F1: (^(v93)), F2: (-(v4))}
    for fi100 := 0; fi100 < 20; fi100++ {
        var v101 string = v98
        fmt.Println(v98)
        _ = fi100
        _ = v101
    }
    var v102 float64 = float64(v94)
    v93 = (^(v93))
    v98 = (v98 + "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
    _ = v93
    _ = v94
    _ = v97
    _ = v98
    _ = v99
}
