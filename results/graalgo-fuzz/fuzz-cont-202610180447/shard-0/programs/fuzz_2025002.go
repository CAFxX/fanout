package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 int; F2 int16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int16; F1 int8; F2 int16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 bool; F1 uint64; F2 uint32 }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]int)
    var v3 int = -9223372036854775808
    v4 := GenS1{F0: (("hello" + ("" + "hello")) + "hello")}
    if false {
        switch (^(v3)) {
        case -581:
            var v5 string = ("a" + "foo bar")
            _ = v5
        case 32767:
            if true {
                if (9223372036854775807 != 9223372036854775807) {
                    fmt.Println(v3)
                }
                v6 := GenS1{F0: ((("" + "hello") + ("hello" + ((("" + "") + "") + "hello"))) + "\n\t")}
                if false {
                    for fi7 := 0; fi7 < 17; fi7++ {
                        for fi8 := 0; fi8 < 5; fi8++ {
                            v2[0] = fi8
                            _ = fi8
                        }
                        _ = fi7
                    }
                    for fi9 := 0; fi9 < 20; fi9++ {
                        v3 += v3
                        v10 := GenS1{F0: "hello"}
                        var v11 float32 = (-2.25 - 1.5)
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                } else {
                    for fi12 := 0; fi12 < 12; fi12++ {
                        v13 := v2[431]
                        for fi14 := 0; fi14 < 16; fi14++ {
                            v3 -= 291
                            v13 -= (v13 + 9223372036854775807)
                            var v15 uint32 = 945
                            _ = fi14
                            _ = v15
                        }
                        _ = fi12
                        _ = v13
                    }
                    v2[65535] = 127
                }
                _ = v6
            } else {
                v3 = 2147483647
                fmt.Println(v3)
            }
        default:
            v3 += (v3 % ((2147483647) | 1))
        }
        if false {
            v3 -= (^(v3))
            v16 := GenS2{F0: 0}
            _ = v16
        } else {
            switch "a" {
            case "foo bar":
                v17 := v2[42]
                _ = v17
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                v3 = (v3 % (((v3 / (((v3 | v3)) | 1))) | 1))
            }
        }
    }
    if true {
        if (1000 > 42) {
            var v18 uint = 255
            fmt.Println(v3, v3)
            _ = v18
        }
    }
    if (!false) {
        switch (v3 / ((v3) | 1)) {
        case 677:
            var v19 int = (v3 >> ((115) & 63))
            _ = v19
        case -872:
            for fi20 := 0; fi20 < 3; fi20++ {
                fmt.Println(v3)
                _ = fi20
            }
        case 2147483647:
            if (2 >= 146) {
                if true {
                    v21 := GenS3{F0: true, F1: (^(v3))}
                    v22 := v2[766]
                    _ = v21
                    _ = v22
                }
                switch (-(v3)) {
                case 2147483647:
                    if (42 != -2147483648) {
                        var v23 uint32 = 15
                        v23 -= 712
                        v23 = v23
                        _ = v23
                    }
                case 32768:
                    if false {
                        if (964 <= 42) {
                            v24 := GenS1{F0: "\n\t"}
                            v2[2] = -2147483648
                            v3 += 71
                            _ = v24
                        } else {
                            var v25 uint = 1
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v25
                        }
                        v26 := GenS3{F0: false, F1: 0}
                        v27 := GenS4{F0: 525}
                        _ = v26
                        _ = v27
                    }
                case -506:
                    var v28 int8 = -1
                    _ = v28
                default:
                    if (-1 != -780) {
                        v29 := GenS5{F0: 2}
                        _ = v29
                    } else {
                        var v30 bool = (!true)
                        _ = v30
                    }
                }
            }
        default:
            fmt.Println(v3)
        }
        v3 = (v3 * v3)
    }
    v3 += v3
    fmt.Println(v3)
    for fi31 := 0; fi31 < 7; fi31++ {
        if false {
            if (3.14159 < -2.25) {
                switch v3 {
                case -9223372036854775808:
                    _ = len(v1)
                case 128:
                    v3 = 2147483647
                case -2147483648:
                    v3 = 1
                default:
                    var v32 int64 = -9223372036854775808
                    _ = v32
                }
                if true {
                    if (!true) {
                        var v33 uint16 = 17
                        _ = v33
                    }
                    v34 := GenS4{F0: 4294967295}
                    if (true || false) {
                        var v35 int16 = -1
                        v35 = 42
                        v36 := GenS5{F0: 65535}
                        _ = v35
                        _ = v36
                    } else {
                        switch "\n\t" {
                        case "":
                            for fi37 := 0; fi37 < 4; fi37++ {
                                var v38 bool = (!true)
                                _ = fi37
                                _ = v38
                            }
                        case "hello":
                            v39 := GenS6{F0: 354, F1: fi31, F2: 363}
                            _ = v39
                        case "a":
                            v40 := GenS6{F0: 4294967295, F1: fi31, F2: -32768}
                            _ = v40
                        default:
                            if false {
                                if false {
                                    var v41 int = (fi31 | fi31)
                                    _ = v41
                                } else {
                                    for fi42 := 0; fi42 < 5; fi42++ {
                                        var v43 bool = (100 >= 128)
                                        fmt.Println(v3)
                                        v43 = v43
                                        _ = fi42
                                        _ = v43
                                    }
                                    switch v3 {
                                    case -565:
                                        if false {
                                            v3 = fi31
                                            switch v3 {
                                            case 128:
                                                v3 += 42
                                            default:
                                                v44 := GenS5{F0: 1}
                                                _ = v44
                                            }
                                        }
                                    case 128:
                                        if len(v1) > 0 {
                                            v1[0] = "hello"
                                        }
                                    default:
                                        _ = len(v2)
                                    }
                                }
                                for fi45 := 0; fi45 < 2; fi45++ {
                                    var v46 string = ("foo bar" + ("a" + "\n\t"))
                                    fmt.Println(v3)
                                    v1 = append(v1, (v46 + (v46 + v46)))
                                    _ = fi45
                                    _ = v46
                                }
                                switch "hello" {
                                case "":
                                    if false {
                                        switch (fi31 & 2147483647) {
                                        case -206:
                                            _ = len(v2)
                                        case -482:
                                            v47 := GenS7{F0: 42}
                                            _ = v47
                                        }
                                        var v48 int = (v3 - (v3 << ((42) & 63)))
                                        _ = v48
                                    }
                                }
                            } else {
                                var v49 uint8 = 31
                                for fi50 := 0; fi50 < 14; fi50++ {
                                    if true {
                                        fmt.Println(fi50, v3, fi31)
                                    }
                                    _ = fi50
                                }
                                _ = v49
                            }
                        }
                    }
                    _ = v34
                }
            }
        } else {
            var v51 int = fi31
            switch "foo bar" {
            case "":
                v51 += fi31
            case "\n\t":
                for fi52 := 0; fi52 < 2; fi52++ {
                    v2[42] = fi31
                    var v53 int = v3
                    v2[1] = 2147483647
                    v51 += fi31
                    v1 = append(v1, ("" + "hello"))
                    _ = fi52
                    _ = v53
                }
            default:
                for fi54 := 0; fi54 < 7; fi54++ {
                    fmt.Println(fi31, v51)
                    switch fi54 {
                    case -393:
                        v55 := GenS3{F0: false, F1: -581}
                        _ = v55
                    case -291:
                        v3 += (v51 - (-(v3)))
                    }
                    if true {
                        if (!true) {
                            if false {
                                var v56 int8 = 0
                                v51 = v51
                                var v57 uint64 = 817
                                if false {
                                    var v58 uint8 = 2
                                    v56 = v56
                                    _ = v58
                                } else {
                                    if false {
                                        var v59 float32 = 1.5
                                        _ = v59
                                    } else {
                                        v60 := GenS8{F0: 42, F1: v56, F2: 1}
                                        v61 := GenS2{F0: 581}
                                        var v62 bool = (int(v57) < fi31)
                                        _ = v60
                                        _ = v61
                                        _ = v62
                                    }
                                    if (true || true) {
                                        v63 := GenS9{F0: -324}
                                        var v64 int32 = 2147483647
                                        switch fi31 {
                                        case 2147483647:
                                            delete(v2, 1)
                                        }
                                        _ = v63
                                        _ = v64
                                    }
                                }
                                v65 := GenS7{F0: (v56 >> ((106) & 63))}
                                if (143 <= 1) {
                                    switch ("foo bar" + (((("" + (("foo bar" + "\n\t") + ("foo bar" + ("foo bar" + "hello")))) + "foo bar") + ("a" + ("" + ("foo bar" + "foo bar")))) + "hello")) {
                                    case "a":
                                        var v66 uint32 = 4294967295
                                        _ = v66
                                    default:
                                        if ((-2.25 != 0.1) || (true || false)) {
                                            v3 = (v51 / ((int(v56)) | 1))
                                            switch "foo bar" {
                                            case "foo bar":
                                                v67 := GenS7{F0: (v56 - v56)}
                                                _ = v67
                                            case "hello":
                                                if false {
                                                    var v68 int8 = 1
                                                    if (124 > 42) {
                                                        if false {
                                                            if (!false) {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                var v69 float64 = -2.25
                                                                _ = v69
                                                            } else {
                                                                v56 = (^(v68))
                                                                var v70 uint = 2
                                                                v51 += (^(v51))
                                                                _ = v70
                                                            }
                                                            switch fi31 {
                                                            case 9223372036854775807:
                                                                switch v51 {
                                                                case 2147483647:
                                                                    v3 = fi54
                                                                case 941:
                                                                    v68 = v68
                                                                }
                                                            case 690:
                                                                var v71 int32 = -2147483648
                                                                _ = v71
                                                            default:
                                                                fmt.Println(v3)
                                                            }
                                                            _ = len(v2)
                                                        } else {
                                                            if false {
                                                                if (!true) {
                                                                    if (true && true) {
                                                                        v72 := v2[2147483647]
                                                                        var v73 uint = 0
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                    if true {
                                                                        v68 -= v68
                                                                        v51 += (fi54 | 127)
                                                                        v51 = -9223372036854775808
                                                                    }
                                                                    var v74 int16 = 0
                                                                    _ = v74
                                                                }
                                                            }
                                                            fmt.Println(v57)
                                                        }
                                                    }
                                                    _ = v68
                                                }
                                                if true {
                                                    _ = len(v1)
                                                    var v75 bool = false
                                                    _ = v75
                                                } else {
                                                    v1 = append(v1, ("" + ((("a" + (("foo bar" + "hello") + "\n\t")) + "a") + ("hello" + "hello"))))
                                                }
                                                v3 = (-(v51))
                                            default:
                                                fmt.Println(v51, v56, fi54)
                                            }
                                            v76 := GenS10{F0: ((v56 << ((42) & 63)) < v56), F1: v57, F2: 2147483648}
                                            _ = v76
                                        }
                                    }
                                } else {
                                    var v77 int16 = 42
                                    _ = v77
                                }
                                _ = v56
                                _ = v57
                                _ = v65
                            }
                        }
                    }
                    _ = fi54
                }
            }
            if false {
                if (((v3 * (^(v51))) == fi31) || false) {
                    var v78 int8 = 100
                    _ = v78
                }
                switch "a" {
                case "foo bar":
                    fmt.Println(fi31)
                }
            } else {
                v51 = -585
                if false {
                    for fi79 := 0; fi79 < 3; fi79++ {
                        if ("\n\t" > "a") {
                            if (((false || (((v3 - (v3 >> ((18446744073709551615) & 63))) < (v3 ^ (-(v3)))) && ((346 < -1) || true))) && true) || (-771 > 1)) {
                                if true {
                                    v51 += (v3 % ((fi79) | 1))
                                }
                                var v80 uint64 = 600
                                v80 -= (v80 + (v80 | (v80 / (((v80 * (v80 - (v80 & (v80 ^ 817))))) | 1))))
                                _ = v80
                            } else {
                                v51 = 2147483647
                            }
                        }
                        var v81 uint = 752
                        _ = fi79
                        _ = v81
                    }
                    v82 := GenS3{F0: (!(18446744073709551615 <= 4294967295)), F1: v3}
                    _ = v82
                }
            }
            _ = v51
        }
        var v83 uint32 = 2147483647
        _ = fi31
        _ = v83
    }
    if (18446744073709551615 < 777) {
        var v84 float32 = (1.5 * ((1.5 * 1.5) / 1.5))
        if true {
            var v85 int32 = 2147483647
            var v86 int16 = 1
            _ = v85
            _ = v86
        }
        _ = v84
    }
    switch v3 {
    case -2147483648:
        v2[65535] = 127
    }
    v87 := GenS4{F0: 1}
    v3 = (v3 | (v3 / (((v3 * (v3 - v3))) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v87
}
