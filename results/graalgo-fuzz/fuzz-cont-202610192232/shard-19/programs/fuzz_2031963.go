package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint; F1 uint8; F2 uint }
type GenS5 struct { F0 int16; F1 int32; F2 string }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 int16; F1 string }
type GenS8 struct { F0 uint32; F1 uint }
type GenS9 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS10 struct { F0 uint8 }

func main() {
    var v1 int32 = -1
    var v2 int64 = -164
    v3 := []string{""}
    v4 := GenS1{F0: 1}
    var v5 float32 = (-2.25 + ((3.14159 * 0.1) + 0.0))
    if (938 <= 120) {
        var v6 int64 = -9223372036854775808
        if false {
            for fi7 := 0; fi7 < 18; fi7++ {
                switch "foo bar" {
                case "\n\t":
                    var v8 float32 = float32(0.0)
                    _ = v8
                }
                v6 = v6
                for fi9 := 0; fi9 < 18; fi9++ {
                    var v10 float32 = v5
                    _ = fi9
                    _ = v10
                }
                _ = fi7
            }
        } else {
            v11 := GenS1{F0: 0}
            if (((float64((v5 + v5)) >= 0.1) && true) && true) {
                var v12 uint32 = 42
                if ((^(v12)) <= v12) {
                    switch (("foo bar" + "") + "foo bar") {
                    case "\n\t":
                        for fi13 := 0; fi13 < 1; fi13++ {
                            v2 = 42
                            if true {
                                v12 += v12
                                if true {
                                    v12 -= 4294967295
                                }
                                var v14 int16 = -730
                                _ = v14
                            } else {
                                var v15 uint16 = 329
                                _ = v15
                            }
                            switch "\n\t" {
                            case "foo bar":
                                fmt.Println(v12, fi13, fi13)
                            case "hello":
                                var v16 uint64 = 4294967295
                                _ = v16
                            }
                            _ = fi13
                        }
                    case "":
                        var v17 float64 = ((-2.25 / (1e10 * 1e10)) + float64((v5 - 3.14159)))
                        _ = v17
                    default:
                        if (((v6 * -1) != 0) || false) {
                            var v18 float64 = 0.1
                            _ = v18
                        } else {
                            fmt.Println(v6, v6, v12)
                        }
                    }
                    v5 += 1.5
                } else {
                    for fi19 := 0; fi19 < 20; fi19++ {
                        fmt.Println(v6, v5)
                        if false {
                            v1 -= 0
                            if len(v3) > 0 {
                                v3[0] = "\n\t"
                            }
                        }
                        _ = fi19
                    }
                    v2 = (v2 + v2)
                }
                _ = v12
            }
            _ = v11
        }
        v6 -= (^(v6))
        _ = v6
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        fmt.Println(v5, v5)
        _ = fi20
    }
    if true {
        for fi21 := 0; fi21 < 19; fi21++ {
            switch ("a" + ("foo bar" + ("" + "\n\t"))) {
            case "":
                for fi22 := 0; fi22 < 14; fi22++ {
                    v5 += (v5 + 3.14159)
                    _ = fi22
                }
            default:
                var v23 bool = true
                _ = v23
            }
            if len(v3) > 0 {
                v3[0] = ("\n\t" + "foo bar")
            }
            v5 += (v5 + v5)
            _ = fi21
        }
        _ = len(v3)
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 int = (fi24 * fi24)
            _ = fi24
            _ = v25
        }
        fmt.Println(v2, v1)
    }
    _ = len(v3)
    if (127 <= 100) {
        if len(v3) > 0 {
            v3[0] = ("\n\t" + "foo bar")
        }
    } else {
        var v26 uint16 = 255
        _ = v26
    }
    v27 := GenS1{F0: 127}
    var v28 uint16 = 370
    for fi29 := 0; fi29 < 10; fi29++ {
        switch "foo bar" {
        case "":
            switch (-(fi29)) {
            case 9223372036854775807:
                if false {
                    switch fi29 {
                    case -163:
                        if true {
                            v28 = 969
                            v30 := GenS2{F0: (^(v28)), F1: 108}
                            _ = v30
                        } else {
                            if false {
                                for fi31 := 0; fi31 < 7; fi31++ {
                                    var v32 uint16 = (v28 << ((18446744073709551615) & 63))
                                    _ = fi31
                                    _ = v32
                                }
                                var v33 uint32 = 2147483648
                                _ = v33
                            } else {
                                v3 = append(v3, ("foo bar" + "\n\t"))
                                var v34 int8 = 25
                                _ = v34
                            }
                            var v35 int64 = (^(v2))
                            _ = v35
                        }
                    case 32768:
                        v36 := GenS3{F0: 255, F1: ""}
                        _ = v36
                    case 420:
                        var v37 int16 = 42
                        _ = v37
                    }
                    switch (fi29 * (fi29 % ((-9223372036854775808) | 1))) {
                    case 32768:
                        switch "foo bar" {
                        case "\n\t":
                            if true {
                                var v38 int = fi29
                                _ = v38
                            } else {
                                if (!(!(true && (false && false)))) {
                                    v1 = 19
                                    fmt.Println(v28, v2, v2)
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                switch 690 {
                                case -661:
                                    if (!false) {
                                        for fi39 := 0; fi39 < 19; fi39++ {
                                            fmt.Println(v5, v28, v1)
                                            v28 -= 65535
                                            fmt.Println(v28)
                                            _ = fi39
                                        }
                                        v40 := GenS1{F0: 127}
                                        var v41 float32 = (v5 + ((float32(1.5) - v5) / ((v5 + float32(float64((v5 - (((1.5 / (float32(float64(float32(1e10))) - 1.5)) * (0.1 + (float32(3.14159) + 1.5))) * (v5 / v5)))))) * (float32(1.5) - 0.0))))
                                        _ = v40
                                        _ = v41
                                    }
                                case 0:
                                    var v42 bool = (false && (true && (("hello" > "foo bar") || true)))
                                    _ = v42
                                }
                            }
                        case "foo bar":
                            v3 = append(v3, "hello")
                        case "hello":
                            v1 = (v1 & (v1 ^ v1))
                        }
                    case 2147483647:
                        v3 = append(v3, ("" + ("\n\t" + "")))
                    case 127:
                        for fi43 := 0; fi43 < 10; fi43++ {
                            v5 = v5
                            if (-1 != v2) {
                                fmt.Println(v2, fi29, fi43)
                                if ((false || ((127 < 127) || (false || (42 == 203)))) || (!((v28 | (^(v28))) < 0))) {
                                    var v44 int = (fi43 - (fi43 * fi29))
                                    v45 := GenS4{F0: 0, F1: 0, F2: 18446744073709551615}
                                    fmt.Println(fi43)
                                    _ = v44
                                    _ = v45
                                }
                                v46 := GenS5{F0: -488, F1: v1, F2: "foo bar"}
                                _ = v46
                            } else {
                                v1 = 2147483647
                            }
                            switch (fi43 - fi43) {
                            case 32768:
                                if (((0 != 670) || false) || true) {
                                    v47 := GenS2{F0: (v28 + 2), F1: 4294967295}
                                    _ = v47
                                } else {
                                    switch (("a" + "hello") + "") {
                                    case "\n\t":
                                        var v48 uint32 = 391
                                        _ = v48
                                    case "hello":
                                        var v49 bool = true
                                        _ = v49
                                    }
                                    v1 = (v1 / ((v1) | 1))
                                }
                            case 419:
                                v2 = v2
                            }
                            _ = fi43
                        }
                    default:
                        if (33 != 255) {
                            fmt.Println(v2)
                        } else {
                            v2 += (^(v2))
                            for fi50 := 0; fi50 < 2; fi50++ {
                                _ = len(v3)
                                _ = fi50
                            }
                            v51 := GenS2{F0: v28, F1: 42}
                            _ = v51
                        }
                    }
                }
            default:
                for fi52 := 0; fi52 < 18; fi52++ {
                    if (42 > -278) {
                        var v53 int64 = 562
                        _ = v53
                    }
                    v54 := GenS2{F0: v28, F1: 1}
                    if false {
                        v55 := GenS4{F0: 42, F1: 128, F2: 149}
                        fmt.Println(fi52, v1, v28)
                        _ = v55
                    } else {
                        v56 := GenS2{F0: v28, F1: 369}
                        _ = v56
                    }
                    _ = fi52
                    _ = v54
                }
            }
        case "foo bar":
            for fi57 := 0; fi57 < 15; fi57++ {
                switch int(v2) {
                case -837:
                    fmt.Println(v1)
                case 128:
                    var v58 uint64 = 318
                    _ = v58
                case 42:
                    var v59 uint64 = 2
                    v2 = v2
                    var v60 string = ("hello" + "\n\t")
                    if (!false) {
                        var v61 uint = 2
                        _ = v61
                    } else {
                        fmt.Println(fi57)
                        v62 := GenS1{F0: 77}
                        v2 += (v2 - int64(fi29))
                        _ = v62
                    }
                    v1 = v1
                    v63 := GenS6{F0: (-(fi57)), F1: 1}
                    v60 = v60
                    _ = v59
                    _ = v60
                    _ = v63
                }
                if true {
                    fmt.Println(v5)
                    switch (fi57 / (((fi57 ^ (fi29 * fi57))) | 1)) {
                    case 216:
                        switch (("\n\t" + "\n\t") + "") {
                        case "foo bar":
                            fmt.Println(fi29, v1)
                        }
                    default:
                        var v64 uint8 = 255
                        _ = v64
                    }
                } else {
                    if false {
                        if (!false) {
                            var v65 uint8 = 130
                            _ = v65
                        } else {
                            v28 = v28
                            if (100 != -1) {
                                v66 := GenS7{F0: -1, F1: "\n\t"}
                                v1 = v1
                                _ = v66
                            } else {
                                fmt.Println(v28, v2, v28)
                                v3 = append(v3, "a")
                                if (!false) {
                                    switch fi57 {
                                    case 127:
                                        v67 := GenS5{F0: -81, F1: (^(v1)), F2: ""}
                                        _ = v67
                                    }
                                    _ = len(v3)
                                }
                            }
                        }
                        var v68 string = ("a" + "")
                        v69 := GenS2{F0: v28, F1: 4294967295}
                        _ = v68
                        _ = v69
                    } else {
                        v5 += (0.0 / float32((-2.25 * (1.5 + float64(v5)))))
                    }
                    switch (("foo bar" + "a") + (("a" + "foo bar") + "\n\t")) {
                    case "foo bar":
                        if false {
                            v2 += 1
                            if ("a" < ("a" + "\n\t")) {
                                if false {
                                    if false {
                                        v70 := GenS8{F0: 688, F1: 755}
                                        var v71 uint8 = 127
                                        _ = v70
                                        _ = v71
                                    } else {
                                        v2 -= (-(v2))
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = ("foo bar" + "hello")
                                        }
                                    }
                                    v72 := GenS2{F0: v28, F1: 581}
                                    var v73 string = "\n\t"
                                    _ = v72
                                    _ = v73
                                } else {
                                    v5 = 0.0
                                }
                                v5 = v5
                                var v74 int16 = -32768
                                if true {
                                    var v75 int8 = 127
                                    v76 := GenS9{F0: v28, F1: 0, F2: 42}
                                    _ = v75
                                    _ = v76
                                } else {
                                    if false {
                                        var v77 int32 = 2147483647
                                        _ = v77
                                    } else {
                                        var v78 int8 = 1
                                        if len(v3) > 0 {
                                            v3[0] = "\n\t"
                                        }
                                        var v79 int16 = v74
                                        _ = v78
                                        _ = v79
                                    }
                                    var v80 uint8 = 1
                                    var v81 bool = true
                                    _ = v80
                                    _ = v81
                                }
                                v82 := GenS6{F0: 408, F1: 127}
                                v74 -= (v74 >> ((65535) & 63))
                                _ = v74
                                _ = v82
                            } else {
                                if false {
                                    switch ("\n\t" + ("" + "a")) {
                                    case "\n\t":
                                        var v83 float64 = -2.25
                                        _ = v83
                                    case "foo bar":
                                        switch (fi29 * 625) {
                                        case 42:
                                            v1 += (-(v1))
                                        case 1:
                                            var v84 string = "\n\t"
                                            _ = v84
                                        case 446:
                                            var v85 uint8 = 204
                                            _ = v85
                                        default:
                                            v2 = (v2 | 1)
                                        }
                                    }
                                }
                                fmt.Println(v5, fi57, fi57)
                                v5 = 0.0
                            }
                            v1 = 859
                            var v86 uint = 1
                            _ = v86
                        } else {
                            var v87 int32 = v1
                            if false {
                                var v88 string = "hello"
                                _ = v88
                            } else {
                                v89 := GenS10{F0: 2}
                                _ = v89
                            }
                            var v90 uint8 = 60
                            _ = v87
                            _ = v90
                        }
                    case "\n\t":
                        fmt.Println(v2, v28)
                    case "a":
                        fmt.Println(v2, v2)
                    default:
                        v91 := GenS3{F0: 127, F1: "\n\t"}
                        _ = v91
                    }
                    fmt.Println(v2)
                }
                _ = fi57
            }
        }
        _ = fi29
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi92 := 0; fi92 < 15; fi92++ {
        v2 += v2
        v28 = (v28 & (v28 | v28))
        _ = fi92
    }
    v93 := GenS8{F0: 849, F1: 4294967295}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v94 int32 = v1
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v93
    _ = v94
}
