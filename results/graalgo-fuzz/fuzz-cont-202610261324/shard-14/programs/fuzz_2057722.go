package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 string }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 bool; F1 bool; F2 int64 }
type GenS4 struct { F0 uint; F1 uint32; F2 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 uint32; F1 uint32; F2 string }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int32; F1 int16 }
type GenS11 struct { F0 uint32; F1 int64; F2 string }
type GenS12 struct { F0 int32; F1 string; F2 uint64 }
type GenS13 struct { F0 uint64; F1 int8; F2 uint }

func main() {
    v1 := []string{"\n\t", "", "hello"}
    var v2 uint16 = 951
    var v3 int64 = 42
    v4 := []string{("foo bar" + ""), "a", "foo bar", "foo bar"}
    var v5 float64 = 3.14159
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    if false {
        fmt.Println(v3)
        v6 := GenS1{F0: -583, F1: false, F2: ""}
        _ = v6
    } else {
        var v7 bool = false
        _ = len(v1)
        var v8 float64 = v5
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := GenS1{F0: 1, F1: (uint(v2) != 171), F2: "foo bar"}
        v11 := GenS1{F0: 304, F1: true, F2: "\n\t"}
        _ = fi9
        _ = v10
        _ = v11
    }
    if (!(972 < 0)) {
        if true {
            switch 252 {
            case 42:
                var v12 uint32 = 2
                _ = v12
            case -9223372036854775808:
                var v13 int32 = -17
                _ = v13
            default:
                switch -9223372036854775808 {
                case 2147483647:
                    var v14 int64 = 9223372036854775807
                    _ = v14
                default:
                    if (18446744073709551615 <= 42) {
                        switch 32767 {
                        case -2147483648:
                            var v15 int = 42
                            _ = v15
                        case 42:
                            if false {
                                v2 = 32767
                                var v16 int64 = v3
                                _ = v16
                            }
                        case 0:
                            var v17 int16 = -504
                            _ = v17
                        }
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        v2 = 0
                    }
                }
            }
        } else {
            if len(v4) > 0 {
                v4[0] = "hello"
            }
            var v18 int64 = (v3 & v3)
            fmt.Println(v2, v3, v3)
            _ = v18
        }
    }
    v1 = append(v1, "\n\t")
    for fi19 := 0; fi19 < 15; fi19++ {
        if (1 <= 4294967295) {
            switch fi19 {
            case 455:
                fmt.Println(v2)
            default:
                var v20 uint = 263
                _ = v20
            }
            var v21 int16 = -856
            v1 = append(v1, (("" + "foo bar") + "foo bar"))
            _ = v21
        } else {
            v22 := GenS2{F0: 1, F1: 42}
            _ = v22
        }
        v23 := GenS3{F0: false, F1: (65535 > 2), F2: v3}
        _ = fi19
        _ = v23
    }
    switch 1 {
    case 42:
        var v24 float32 = (float32(3.14159) * float32(float64(1e10)))
        _ = v24
    default:
        fmt.Println(v3)
    }
    v25 := GenS2{F0: 524, F1: 4294967295}
    var v26 float32 = 0.0
    v27 := GenS4{F0: 195, F1: 2, F2: 1}
    switch "hello" {
    case "foo bar":
        v28 := GenS5{F0: 0}
        _ = v28
    case "a":
        v29 := GenS4{F0: 2, F1: 644, F2: (v3 + (v3 >> ((65535) & 63)))}
        _ = v29
    case "":
        if true {
            for fi30 := 0; fi30 < 5; fi30++ {
                var v31 string = (("" + "foo bar") + "\n\t")
                for fi32 := 0; fi32 < 16; fi32++ {
                    if false {
                        if false {
                            v31 = v31
                            v5 = v5
                        }
                    } else {
                        var v33 uint = 4294967295
                        _ = v33
                    }
                    _ = fi32
                }
                _ = fi30
                _ = v31
            }
            fmt.Println(v26)
        }
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if (1 < 61) {
        var v34 int32 = 2147483647
        fmt.Println(v26, v2)
        _ = v34
    } else {
        if (!false) {
            var v35 int = 1
            for fi36 := 0; fi36 < 15; fi36++ {
                if true {
                    switch "hello" {
                    case "hello":
                        _ = len(v4)
                    case "":
                        if len(v4) > 0 {
                            v4[0] = "foo bar"
                        }
                    default:
                        v37 := GenS3{F0: false, F1: false, F2: (v3 | v3)}
                        _ = v37
                    }
                    switch (("\n\t" + ("hello" + ("\n\t" + "a"))) + "foo bar") {
                    case "a":
                        v2 -= (v2 & (v2 & v2))
                    case "hello":
                        for fi38 := 0; fi38 < 15; fi38++ {
                            if (42 == 217) {
                                if false {
                                    if false {
                                        v35 = v35
                                        var v39 int = (fi36 + (fi36 + (v35 - fi36)))
                                        _ = v39
                                    } else {
                                        var v40 int64 = (v3 | (v3 * v3))
                                        v41 := GenS6{F0: ("\n\t" + ""), F1: 738}
                                        v26 = v26
                                        _ = v40
                                        _ = v41
                                    }
                                    v3 -= -9223372036854775808
                                }
                            } else {
                                var v42 uint16 = 32767
                                _ = v42
                            }
                            _ = fi38
                        }
                    }
                } else {
                    var v43 int16 = -1
                    if (!(!(true && false))) {
                        var v44 int16 = v43
                        for fi45 := 0; fi45 < 13; fi45++ {
                            var v46 int = (fi45 / ((32767) | 1))
                            var v47 uint32 = 2147483648
                            if len(v1) > 0 {
                                v1[0] = ("foo bar" + "hello")
                            }
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                        _ = v44
                    } else {
                        var v48 uint64 = 18446744073709551615
                        for fi49 := 0; fi49 < 4; fi49++ {
                            switch ("\n\t" + ((("" + "a") + "hello") + "foo bar")) {
                            case "hello":
                                v50 := GenS7{F0: 2147483647, F1: 2147483648, F2: ("hello" + "")}
                                _ = v50
                            case "foo bar":
                                v26 = 0.1
                            case "\n\t":
                                v26 = float32(v5)
                            default:
                                v35 -= (v35 + (^(fi49)))
                            }
                            var v51 uint8 = 255
                            if true {
                                var v52 int32 = 1
                                _ = v52
                            }
                            _ = fi49
                            _ = v51
                        }
                        var v53 int32 = -1
                        _ = v48
                        _ = v53
                    }
                    _ = v43
                }
                _ = fi36
            }
            if (581 == 255) {
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
            }
            _ = v35
        } else {
            _ = len(v1)
            switch (("hello" + "hello") + "hello") {
            case "hello":
                for fi54 := 0; fi54 < 9; fi54++ {
                    var v55 int8 = 42
                    for fi56 := 0; fi56 < 3; fi56++ {
                        v26 = (v26 - ((v26 * float32(v5)) / v26))
                        v2 = (v2 / (((v2 / ((32767) | 1))) | 1))
                        _ = fi56
                    }
                    _ = fi54
                    _ = v55
                }
            case "\n\t":
                for fi57 := 0; fi57 < 20; fi57++ {
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                    var v58 float32 = 0.0
                    _ = fi57
                    _ = v58
                }
            default:
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(v2)
                }
            }
            if (!(!true)) {
                for fi59 := 0; fi59 < 10; fi59++ {
                    if false {
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        v3 = 32
                    }
                    switch (("\n\t" + "foo bar") + "hello") {
                    case "foo bar":
                        switch (fi59 % ((32767) | 1)) {
                        case 1:
                            for fi60 := 0; fi60 < 13; fi60++ {
                                switch (fi59 + fi60) {
                                case -2147483648:
                                    switch "hello" {
                                    case "hello":
                                        var v61 float32 = 0.1
                                        _ = v61
                                    default:
                                        v62 := GenS8{F0: 1}
                                        _ = v62
                                    }
                                case 411:
                                    if (-127 != 100) {
                                        var v63 uint8 = 0
                                        _ = v63
                                    }
                                case 2147483647:
                                    v26 -= v26
                                default:
                                    v64 := GenS9{F0: (v3 ^ v3)}
                                    _ = v64
                                }
                                if false {
                                    v26 += 1e10
                                    if true {
                                        if false {
                                            fmt.Println(v5, fi60, v3)
                                            var v65 uint32 = 2147483647
                                            if (!(2 < 0)) {
                                                var v66 bool = true
                                                var v67 uint16 = (v2 ^ v2)
                                                v26 += v26
                                                _ = v66
                                                _ = v67
                                            } else {
                                                if (0 != 738) {
                                                    var v68 uint = 255
                                                    _ = v68
                                                } else {
                                                    if false {
                                                        v3 -= (^(v3))
                                                    } else {
                                                        v3 = 1
                                                        var v69 int16 = 1
                                                        if true {
                                                            v70 := GenS3{F0: false, F1: true, F2: (v3 & (-(v3)))}
                                                            v71 := GenS10{F0: 42, F1: (v69 | v69)}
                                                            var v72 string = ""
                                                            v73 := GenS2{F0: 0, F1: 4294967295}
                                                            _ = v70
                                                            _ = v71
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                        _ = v69
                                                    }
                                                    var v74 float64 = ((float64((float32(v5) / (v26 / 1.0))) - v5) - float64(((v26 * v26) - (v26 + v26))))
                                                    _ = v74
                                                }
                                                v26 += float32(-2.25)
                                                var v75 int64 = (v3 >> ((4294967295) & 63))
                                                _ = v75
                                            }
                                            _ = v65
                                        } else {
                                            fmt.Println(v26, v5, v3)
                                            v3 = v3
                                            var v76 int32 = 0
                                            if false {
                                                var v77 float32 = float32(3.14159)
                                                _ = v77
                                            }
                                            var v78 float32 = ((-2.25 + -2.25) / v26)
                                            _ = v76
                                            _ = v78
                                        }
                                        v3 -= v3
                                        v3 += (v3 << ((1) & 63))
                                    } else {
                                        if ((v5 - (v5 * (((((0.1 * (((v5 + v5) - float64(float32(v5))) / 1e10)) * ((v5 * (v5 / float64((-2.25 - (-2.25 - v26))))) / 1.0)) / 0.1) + 0.1) - v5))) == float64((v26 + -2.25))) {
                                            var v79 string = "a"
                                            _ = v79
                                        }
                                        v3 -= (v3 - 9223372036854775807)
                                        switch (-(fi60)) {
                                        case 127:
                                            v26 = (v26 + v26)
                                        default:
                                            if len(v4) > 0 {
                                                v4[0] = "hello"
                                            }
                                        }
                                        if false {
                                            v80 := GenS3{F0: (!true), F1: (!false), F2: 959}
                                            _ = v80
                                        } else {
                                            if (616 < 2147483648) {
                                                switch fi60 {
                                                case 9223372036854775807:
                                                    var v81 uint64 = 473
                                                    _ = v81
                                                default:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                            } else {
                                                v82 := GenS11{F0: 67, F1: v3, F2: "\n\t"}
                                                _ = v82
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v2 = 32767
                                        }
                                        _ = len(v1)
                                        if len(v4) > 0 {
                                            v4[0] = ("\n\t" + "foo bar")
                                        }
                                    }
                                    var v83 string = "hello"
                                    _ = v83
                                }
                                v26 -= (v26 + 0.0)
                                var v84 int8 = 88
                                v26 += v26
                                switch -895 {
                                case -890:
                                    switch (fi60 | 32768) {
                                    case 42:
                                        _ = len(v1)
                                    }
                                default:
                                    switch "a" {
                                    case "":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case "foo bar":
                                        if (-32768 != 616) {
                                            v85 := GenS8{F0: 26}
                                            v86 := GenS3{F0: (false && (41 > 255)), F1: false, F2: (v3 << ((378) & 63))}
                                            _ = v85
                                            _ = v86
                                        } else {
                                            v87 := GenS5{F0: 42}
                                            switch (fi59 >> ((18446744073709551615) & 63)) {
                                            case -9223372036854775808:
                                                var v88 int = fi60
                                                _ = v88
                                            case 32768:
                                                switch (("hello" + "foo bar") + ((("\n\t" + "") + ((("hello" + "hello") + "foo bar") + "\n\t")) + "")) {
                                                case "\n\t":
                                                    v2 = (v2 % (((v2 | v2)) | 1))
                                                case "foo bar":
                                                    var v89 uint32 = 225
                                                    _ = v89
                                                }
                                            case 966:
                                                var v90 int32 = 283
                                                _ = v90
                                            default:
                                                _ = len(v1)
                                            }
                                            var v91 int64 = 9223372036854775807
                                            _ = v87
                                            _ = v91
                                        }
                                    case "\n\t":
                                        v92 := GenS12{F0: -2147483648, F1: ("a" + "\n\t"), F2: 1}
                                        _ = v92
                                    default:
                                        if false {
                                            if (!true) {
                                                var v93 string = "\n\t"
                                                v26 += (v26 + 0.0)
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v93
                                            } else {
                                                v84 -= (^(v84))
                                            }
                                            fmt.Println(fi59, fi59, v2)
                                            fmt.Println(v2)
                                            v94 := GenS13{F0: 1, F1: (v84 & v84), F2: 0}
                                            v95 := GenS11{F0: 1, F1: (v3 | 42), F2: ("a" + "hello")}
                                            _ = v94
                                            _ = v95
                                        }
                                    }
                                }
                                _ = fi60
                                _ = v84
                            }
                        default:
                            v2 -= (^(v2))
                        }
                    case "hello":
                        for fi96 := 0; fi96 < 12; fi96++ {
                            v97 := GenS3{F0: true, F1: (65535 < (v2 * 225)), F2: (v3 % ((v3) | 1))}
                            var v98 int32 = 42
                            _ = fi96
                            _ = v97
                            _ = v98
                        }
                    case "":
                        v3 = 9223372036854775807
                    default:
                        fmt.Println(fi59, fi59, fi59)
                    }
                    _ = fi59
                }
                var v99 uint16 = v2
                fmt.Println(v5, v2, v99)
                _ = v99
            } else {
                for fi100 := 0; fi100 < 1; fi100++ {
                    fmt.Println(v2, v3, fi100)
                    _ = fi100
                }
            }
        }
        v1 = append(v1, "hello")
    }
    var v101 string = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v101
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
