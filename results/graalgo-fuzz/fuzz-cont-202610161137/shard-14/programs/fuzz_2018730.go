package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int64; F2 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 bool; F1 int16; F2 uint64 }
type GenS9 struct { F0 int64; F1 bool }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint32; F1 bool; F2 int16 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int64; F1 int32; F2 int16 }

func main() {
    var v1 int16 = 1
    var v2 uint64 = 42
    v3 := make(map[int]bool)
    var v4 float64 = 3.14159
    v5 := []string{}
    var v6 uint64 = 0
    switch 0 {
    case 32768:
        v7 := v3[-2147483648]
        _ = v7
    }
    if true {
        v8 := GenS1{F0: ("foo bar" + (("a" + "foo bar") + "a"))}
        var v9 uint32 = 490
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: ("a" + "hello")}
    if false {
        v1 = (v1 % ((v1) | 1))
        if false {
            v4 = v4
            var v11 string = (("foo bar" + "") + "")
            var v12 bool = true
            _ = v11
            _ = v12
        } else {
            v13 := GenS2{F0: 1}
            var v14 uint64 = v6
            switch (("hello" + "a") + (("hello" + ("foo bar" + ("a" + ("foo bar" + "foo bar")))) + "a")) {
            case "hello":
                var v15 uint8 = 124
                _ = v15
            case "a":
                for fi16 := 0; fi16 < 9; fi16++ {
                    v4 = 3.14159
                    v17 := GenS3{F0: (^(v1))}
                    v4 += 0.0
                    _ = fi16
                    _ = v17
                }
            default:
                var v18 int16 = 1
                _ = v18
            }
            _ = v13
            _ = v14
        }
    }
    if (413 < 65535) {
        for fi19 := 0; fi19 < 12; fi19++ {
            fmt.Println(v2, v2, fi19)
            v6 = 1
            for fi20 := 0; fi20 < 14; fi20++ {
                switch "foo bar" {
                case "\n\t":
                    var v21 int8 = 0
                    _ = v21
                default:
                    fmt.Println(v1)
                }
                v1 = 42
                if true {
                    if ((!(65535 < 135)) || false) {
                        v4 = (v4 - (((v4 * -2.25) + v4) + -2.25))
                        var v22 uint32 = 42
                        if (1.5 >= (((v4 - float64((0.1 - -2.25))) / v4) + v4)) {
                            v22 -= (v22 + (^(v22)))
                        } else {
                            v6 -= (v2 & v2)
                            if (v6 != uint64(v22)) {
                                delete(v3, fi20)
                                var v23 uint = 1
                                switch fi19 {
                                case -9223372036854775808:
                                    if false {
                                        var v24 float32 = (0.1 - float32((v4 + (v4 + v4))))
                                        var v25 float64 = v4
                                        v4 = v25
                                        _ = v24
                                        _ = v25
                                    }
                                case -6:
                                    fmt.Println(v6)
                                default:
                                    _ = len(v3)
                                }
                                _ = v23
                            }
                        }
                        v26 := GenS1{F0: "foo bar"}
                        _ = v22
                        _ = v26
                    } else {
                        if (false && false) {
                            if (true || (0.1 == 1.5)) {
                                v27 := GenS4{F0: 1000, F1: 9223372036854775807, F2: -128}
                                fmt.Println(v6)
                                v28 := GenS4{F0: 78, F1: 0, F2: 0}
                                var v29 int32 = -740
                                v30 := GenS5{F0: "hello"}
                                if (!true) {
                                    if (255 < 138) {
                                        v29 += 42
                                        v2 -= v6
                                        var v31 int32 = (v29 << ((18446744073709551615) & 63))
                                        _ = v31
                                    } else {
                                        v29 = v29
                                    }
                                } else {
                                    fmt.Println(v4)
                                }
                                _ = v27
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                            switch fi20 {
                            case 71:
                                v6 = v2
                            case 32767:
                                if false {
                                    delete(v3, -9223372036854775808)
                                    if (((0 != 0) && true) && (100 >= 42)) {
                                        v4 = (((float64(1e10) / (v4 * 1.5)) - v4) / v4)
                                    }
                                    v4 += 1.5
                                } else {
                                    v32 := GenS6{F0: 65535, F1: (v6 << ((4294967295) & 63))}
                                    if (!(((true && true) || false) || (42 != 1))) {
                                        v4 -= float64(3.14159)
                                    }
                                    v2 -= v2
                                    _ = v32
                                }
                            }
                            v2 = (^(v2))
                            v1 = v1
                        } else {
                            v33 := GenS1{F0: "a"}
                            if (32767 == 2) {
                                v34 := GenS6{F0: 546, F1: 2}
                                v4 += 1.5
                                _ = v34
                            }
                            if (((((float64((0.1 - 0.0)) + ((v4 + (v4 + 3.14159)) * (1e10 * 3.14159))) - 0.1) - v4) - (float64(0.1) / (3.14159 - ((1e10 / (v4 * -2.25)) - 3.14159)))) > v4) {
                                v2 = 77
                                var v35 int8 = -1
                                switch (("\n\t" + "a") + ("\n\t" + ("" + (("\n\t" + (("\n\t" + "hello") + ("foo bar" + "\n\t"))) + (((("foo bar" + ("" + "")) + "") + "") + ("" + "foo bar")))))) {
                                case "":
                                    v2 = (v6 >> ((65535) & 63))
                                    if (("" >= "foo bar") && true) {
                                        var v36 int8 = (v35 % ((v35) | 1))
                                        _ = v36
                                    } else {
                                        var v37 int = fi19
                                        _ = v37
                                    }
                                    v38 := GenS7{F0: 127, F1: 2}
                                    _ = v38
                                case "hello":
                                    var v39 float32 = 3.14159
                                    _ = v39
                                case "foo bar":
                                    fmt.Println(v6)
                                }
                                _ = v35
                            }
                            v2 -= v2
                            var v40 string = ("" + "foo bar")
                            switch "foo bar" {
                            case "":
                                v41 := GenS8{F0: (42 < 2), F1: v1, F2: v6}
                                _ = v41
                            case "a":
                                var v42 int = -2147483648
                                _ = v42
                            }
                            _ = v33
                            _ = v40
                        }
                        v1 = v1
                        v2 -= v6
                    }
                    v43 := GenS5{F0: (("foo bar" + "a") + ((((((("foo bar" + "\n\t") + "a") + ("foo bar" + "a")) + "foo bar") + "a") + "") + "foo bar"))}
                    _ = v43
                }
                _ = fi20
            }
            _ = fi19
        }
        v2 += (v2 % (((v6 - (v6 << ((1) & 63)))) | 1))
        v1 += v1
    } else {
        v44 := GenS5{F0: ""}
        _ = v44
    }
    var v45 uint16 = 741
    switch "\n\t" {
    case "":
        for fi46 := 0; fi46 < 12; fi46++ {
            v45 = (v45 ^ (v45 | (v45 * v45)))
            v47 := GenS9{F0: 1, F1: false}
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
            _ = fi46
            _ = v47
        }
    }
    if true {
        var v48 int64 = 42
        delete(v3, 128)
        _ = v48
    } else {
        v5 = append(v5, "\n\t")
        for fi49 := 0; fi49 < 18; fi49++ {
            v50 := GenS5{F0: "a"}
            v51 := GenS10{F0: fi49}
            v6 = 42
            _ = fi49
            _ = v50
            _ = v51
        }
        if false {
            switch "\n\t" {
            case "hello":
                switch -2147483648 {
                case -488:
                    var v52 bool = true
                    _ = v52
                }
            case "\n\t":
                switch "foo bar" {
                case "":
                    for fi53 := 0; fi53 < 17; fi53++ {
                        v6 = v6
                        _ = fi53
                    }
                case "\n\t":
                    if (true && (255 > 1)) {
                        v54 := GenS11{F0: 1, F1: ((2 == 360) && false), F2: (v1 << ((4294967295) & 63))}
                        if false {
                            switch 0 {
                            case -2147483648:
                                for fi55 := 0; fi55 < 17; fi55++ {
                                    fmt.Println(v4)
                                    if (1 <= 18446744073709551615) {
                                        var v56 bool = ("hello" >= ("a" + "hello"))
                                        v45 += v45
                                        _ = v56
                                    } else {
                                        fmt.Println(v45, v45, v2)
                                        for fi57 := 0; fi57 < 6; fi57++ {
                                            var v58 int16 = v1
                                            v1 += (v58 >> ((2) & 63))
                                            v4 -= (0.1 / v4)
                                            _ = fi57
                                            _ = v58
                                        }
                                    }
                                    _ = fi55
                                }
                            case -159:
                                for fi59 := 0; fi59 < 20; fi59++ {
                                    v6 = v6
                                    v45 -= v45
                                    _ = fi59
                                }
                            }
                        }
                        switch ("" + "a") {
                        case "foo bar":
                            if false {
                                var v60 int = 2
                                if true {
                                    v2 += (v2 / (((^(v2))) | 1))
                                    var v61 uint8 = 2
                                    _ = v61
                                } else {
                                    switch (v60 << ((1) & 63)) {
                                    case -2147483648:
                                        v62 := GenS3{F0: -32768}
                                        _ = v62
                                    case 460:
                                        if false {
                                            v63 := GenS12{F0: 0}
                                            var v64 int8 = 0
                                            _ = v63
                                            _ = v64
                                        }
                                    case 9223372036854775807:
                                        if (127 > 1) {
                                            v65 := GenS2{F0: 276}
                                            _ = len(v5)
                                            for fi66 := 0; fi66 < 5; fi66++ {
                                                if (!false) {
                                                    if (!(false || false)) {
                                                        var v67 uint16 = (v45 ^ v45)
                                                        _ = v67
                                                    }
                                                    for fi68 := 0; fi68 < 3; fi68++ {
                                                        var v69 int32 = 2147483647
                                                        var v70 float64 = v4
                                                        _ = fi68
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                    v4 += 0.1
                                                } else {
                                                    v71 := GenS9{F0: 42, F1: false}
                                                    v2 += 555
                                                    _ = v71
                                                }
                                                switch int(v2) {
                                                case 32768:
                                                    v72 := GenS5{F0: ("foo bar" + "hello")}
                                                    _ = v72
                                                case 173:
                                                    fmt.Println(v2)
                                                default:
                                                    switch 9223372036854775807 {
                                                    case 9223372036854775807:
                                                        v2 = (^(v2))
                                                    case -687:
                                                        v45 = (v45 + 359)
                                                    }
                                                }
                                                v73 := GenS8{F0: (42 < 127), F1: -32768, F2: (v6 | (v2 - v2))}
                                                _ = fi66
                                                _ = v73
                                            }
                                            _ = v65
                                        }
                                    }
                                }
                                fmt.Println(v60, v2, v60)
                                _ = v60
                            }
                        case "a":
                            var v74 int = 453
                            _ = v74
                        case "hello":
                            v75 := GenS13{F0: -81, F1: 1, F2: 32767}
                            _ = v75
                        }
                        _ = v54
                    }
                }
            }
            for fi76 := 0; fi76 < 10; fi76++ {
                fmt.Println(v1)
                v1 = v1
                _ = fi76
            }
            fmt.Println(v2, v45, v4)
        } else {
            switch -614 {
            case 127:
                for fi77 := 0; fi77 < 20; fi77++ {
                    if false {
                        for fi78 := 0; fi78 < 20; fi78++ {
                            var v79 bool = true
                            fmt.Println(fi77, v1)
                            var v80 uint32 = 433
                            _ = fi78
                            _ = v79
                            _ = v80
                        }
                    } else {
                        _ = len(v5)
                        for fi81 := 0; fi81 < 7; fi81++ {
                            var v82 float64 = v4
                            v83 := GenS9{F0: -9223372036854775808, F1: true}
                            var v84 int = -2147483648
                            _ = fi81
                            _ = v82
                            _ = v83
                            _ = v84
                        }
                        v5 = append(v5, ("\n\t" + "\n\t"))
                    }
                    var v85 bool = false
                    if ((v85 || (!v85)) && (!true)) {
                        if len(v5) > 0 {
                            v5[0] = ("\n\t" + "")
                        }
                        if v85 {
                            var v86 uint32 = 42
                            v86 += 1
                            for fi87 := 0; fi87 < 4; fi87++ {
                                _ = len(v5)
                                _ = fi87
                            }
                            _ = v86
                        }
                    } else {
                        for fi88 := 0; fi88 < 14; fi88++ {
                            if (fi88 < 128) {
                                v1 = (v1 % ((v1) | 1))
                                var v89 int = fi77
                                fmt.Println(v45)
                                _ = v89
                            }
                            var v90 float32 = 0.1
                            _ = fi88
                            _ = v90
                        }
                        var v91 uint16 = v45
                        var v92 uint8 = 127
                        _ = v91
                        _ = v92
                    }
                    _ = fi77
                    _ = v85
                }
            default:
                if (v2 == v6) {
                    fmt.Println(v4)
                    fmt.Println(v4)
                    v93 := GenS12{F0: -108}
                    _ = v93
                }
            }
        }
    }
    fmt.Println(v4, v4)
    for fi94 := 0; fi94 < 3; fi94++ {
        v4 += float64((1.5 * 1.5))
        v95 := v3[(fi94 % (((fi94 << ((42) & 63))) | 1))]
        _ = fi94
        _ = v95
    }
    if true {
        var v96 float32 = (3.14159 * 0.0)
        _ = v96
    } else {
        if false {
            var v97 string = "a"
            fmt.Println(v45, v4, v97)
            v98 := GenS12{F0: -91}
            _ = v97
            _ = v98
        }
        if true {
            fmt.Println(v45, v6)
            var v99 uint64 = (v6 / (((v6 & (v2 + v6))) | 1))
            v100 := v3[-575]
            _ = v100
            _ = v99
        } else {
            delete(v3, 674)
        }
    }
    var v101 uint16 = (v45 | (v45 % ((1) | 1)))
    fmt.Println(v6)
    for fi102 := 0; fi102 < 11; fi102++ {
        if true {
            if len(v5) > 0 {
                v5[0] = ""
            }
            if (1 <= 65) {
                v1 = -1
            } else {
                v103 := GenS10{F0: (fi102 / ((fi102) | 1))}
                fmt.Println(v101, v101, v4)
                _ = v103
            }
        }
        switch "hello" {
        case "hello":
            var v104 uint8 = 42
            _ = v104
        case "":
            v105 := GenS1{F0: "foo bar"}
            _ = v105
        default:
            fmt.Println(v6, v45, v4)
        }
        _ = fi102
    }
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
