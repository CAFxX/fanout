package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 int8; F2 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32; F1 int16; F2 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 string; F1 uint32; F2 int32 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int8; F1 int16 }
type GenS12 struct { F0 bool; F1 int16 }
type GenS13 struct { F0 string; F1 uint }
type GenS14 struct { F0 int32; F1 uint }

func main() {
    var v1 uint = 4294967295
    var v2 float32 = 0.1
    v3 := []string{"", "", "\n\t", "\n\t"}
    fmt.Println(v2, v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 0 {
    case 468:
        for fi4 := 0; fi4 < 16; fi4++ {
            switch fi4 {
            case 18:
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v5 uint8 = 128
                    if true {
                        for fi6 := 0; fi6 < 12; fi6++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v7 := GenS1{F0: (v5 / ((1) | 1)), F1: "hello", F2: (fi6 | (fi6 * (fi6 * fi4)))}
                            _ = fi6
                            _ = v7
                        }
                    }
                    _ = v5
                }
            case -9223372036854775808:
                for fi8 := 0; fi8 < 3; fi8++ {
                    switch ("foo bar" + "") {
                    case "foo bar":
                        switch -9223372036854775808 {
                        case 549:
                            v2 = (float32(-2.25) / v2)
                        case 0:
                            v9 := GenS2{F0: 0, F1: 2}
                            _ = v9
                        case -2147483648:
                            var v10 int64 = -267
                            _ = v10
                        default:
                            var v11 bool = (-181 >= 2147483647)
                            _ = v11
                        }
                    case "\n\t":
                        v2 += v2
                    }
                    _ = fi8
                }
            case -363:
                v12 := GenS1{F0: 1, F1: "hello", F2: (fi4 - fi4)}
                _ = v12
            }
            switch "hello" {
            case "hello":
                var v13 int8 = 42
                _ = v13
            }
            _ = fi4
        }
    default:
        v1 = v1
    }
    var v14 int16 = 0
    var v15 float64 = 0.1
    v16 := GenS1{F0: 31, F1: "a", F2: -698}
    if false {
        if (!true) {
            for fi17 := 0; fi17 < 5; fi17++ {
                _ = len(v3)
                for fi18 := 0; fi18 < 10; fi18++ {
                    fmt.Println(fi17)
                    v19 := GenS2{F0: 0, F1: 492}
                    _ = fi18
                    _ = v19
                }
                v14 += v14
                _ = fi17
            }
        } else {
            var v20 string = "a"
            var v21 uint = v1
            if (true && (611 < v1)) {
                v21 += (v21 + 238)
                var v22 int32 = int32(v21)
                var v23 int64 = -15
                _ = v22
                _ = v23
            }
            _ = v20
            _ = v21
        }
    }
    if true {
        if false {
            var v24 uint64 = 186
            v25 := GenS3{F0: (("a" != "") || (!true))}
            _ = v24
            _ = v25
        } else {
            for fi26 := 0; fi26 < 11; fi26++ {
                v27 := GenS1{F0: 44, F1: "foo bar", F2: fi26}
                var v28 int32 = 954
                _ = fi26
                _ = v27
                _ = v28
            }
            for fi29 := 0; fi29 < 11; fi29++ {
                v2 = v2
                switch (("foo bar" + (("a" + "foo bar") + (("a" + ("\n\t" + "\n\t")) + "hello"))) + "foo bar") {
                case "a":
                    v15 += ((0.1 + 3.14159) / v15)
                case "\n\t":
                    if len(v3) > 0 {
                        v3[0] = ("\n\t" + "")
                    }
                case "":
                    switch (fi29 ^ (-(fi29))) {
                    case 32767:
                        var v30 uint32 = 2
                        _ = v30
                    }
                default:
                    v31 := GenS4{F0: v14, F1: -37, F2: "a"}
                    _ = v31
                }
                v32 := GenS4{F0: (v14 - (v14 >> ((392) & 63))), F1: 100, F2: ("a" + "foo bar")}
                _ = fi29
                _ = v32
            }
        }
        _ = len(v3)
        var v33 uint64 = 814
        _ = v33
    } else {
        for fi34 := 0; fi34 < 7; fi34++ {
            v15 += 0.1
            fmt.Println(v2, v2, v1)
            _ = fi34
        }
        for fi35 := 0; fi35 < 19; fi35++ {
            v36 := GenS5{F0: 42}
            _ = fi35
            _ = v36
        }
    }
    fmt.Println(v15, v15, v1)
    for fi37 := 0; fi37 < 16; fi37++ {
        if true {
            if (596 < 2147483647) {
                var v38 int64 = -2
                _ = v38
            } else {
                switch -9223372036854775808 {
                case 128:
                    for fi39 := 0; fi39 < 4; fi39++ {
                        var v40 int64 = -646
                        v41 := GenS6{F0: 2147483647, F1: (v14 >> ((18446744073709551615) & 63)), F2: 1}
                        v42 := GenS1{F0: 0, F1: ("" + "a"), F2: (fi39 - fi37)}
                        _ = fi39
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                }
                for fi43 := 0; fi43 < 11; fi43++ {
                    v44 := GenS2{F0: 0, F1: 2}
                    if (false || (-128 == 42)) {
                        if false {
                            var v45 uint8 = 0
                            switch "hello" {
                            case "foo bar":
                                v14 = (^(v14))
                            case "\n\t":
                                v46 := GenS3{F0: false}
                                _ = v46
                            case "hello":
                                var v47 int64 = 42
                                if false {
                                    v2 += v2
                                    var v48 uint64 = 545
                                    var v49 int64 = 129
                                    _ = v48
                                    _ = v49
                                } else {
                                    var v50 int32 = 42
                                    fmt.Println(v1)
                                    _ = len(v3)
                                    _ = v50
                                }
                                v51 := GenS3{F0: true}
                                v2 = (v2 + -2.25)
                                _ = v47
                                _ = v51
                            default:
                                v1 = v1
                            }
                            switch fi43 {
                            case 0:
                                if false {
                                    v45 = 0
                                    v45 -= v45
                                    var v52 int32 = -122
                                    _ = v52
                                }
                            case 784:
                                v2 = (v2 - 0.0)
                            case 1:
                                v53 := GenS7{F0: 42}
                                _ = v53
                            }
                            _ = v45
                        } else {
                            var v54 float32 = 0.1
                            v55 := GenS8{F0: ("foo bar" + (("" + "hello") + (("\n\t" + "") + "foo bar"))), F1: 315, F2: -1}
                            _ = len(v3)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v2 = 3.14159
                            _ = v54
                            _ = v55
                        }
                    } else {
                        if ((!false) || (!(!((1 <= v1) || ((v1 | (v1 | (v1 ^ v1))) <= 255))))) {
                            fmt.Println(v15, v2, v1)
                            var v56 int = fi37
                            _ = len(v3)
                            _ = v56
                        } else {
                            fmt.Println(v15, v1, fi43)
                            var v57 uint32 = 4294967295
                            if ((fi43 - fi43) != 4294967295) {
                                _ = len(v3)
                                v58 := GenS2{F0: 183, F1: 1}
                                v57 = v57
                                fmt.Println(v14, fi37, fi43)
                                _ = v58
                            } else {
                                var v59 string = "foo bar"
                                v59 = "\n\t"
                                var v60 float64 = v15
                                _ = v59
                                _ = v60
                            }
                            _ = v57
                        }
                        switch "hello" {
                        case "a":
                            switch "a" {
                            case "foo bar":
                                if true {
                                    v3 = append(v3, "hello")
                                } else {
                                    v1 = 719
                                    v14 = -1
                                }
                            case "a":
                                if true {
                                    if true {
                                        v2 -= (0.1 / v2)
                                        v14 += v14
                                        switch fi43 {
                                        case 4294967295:
                                            v61 := GenS2{F0: 75, F1: 4294967295}
                                            if true {
                                                if (false && (0 < 1)) {
                                                    fmt.Println(v1, fi43)
                                                    if true {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        v62 := GenS7{F0: int64(v14)}
                                                        fmt.Println(fi37, v15, fi37)
                                                        _ = v62
                                                    } else {
                                                        switch "" {
                                                        case "":
                                                            fmt.Println(v2, fi37, fi43)
                                                        case "a":
                                                            if (!(false && (!false))) {
                                                                var v63 int = 9223372036854775807
                                                                _ = v63
                                                            }
                                                            v1 -= v1
                                                        default:
                                                            if false {
                                                                var v64 int32 = -1
                                                                _ = v64
                                                            } else {
                                                                var v65 float64 = (v15 / 1.5)
                                                                var v66 uint16 = 32767
                                                                _ = v65
                                                                _ = v66
                                                            }
                                                        }
                                                    }
                                                }
                                                v67 := GenS9{F0: 327}
                                                v68 := GenS10{F0: (!true)}
                                                _ = v67
                                                _ = v68
                                            } else {
                                                v69 := GenS3{F0: ((384 >= 2) && (212 <= 128))}
                                                v15 = v15
                                                var v70 float32 = ((float32(v15) + (-2.25 * 0.0)) / float32(1.5))
                                                _ = v69
                                                _ = v70
                                            }
                                            if true {
                                                v2 += v2
                                                var v71 float64 = float64((v2 * 0.0))
                                                var v72 uint32 = 2
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v71
                                                _ = v72
                                            }
                                            _ = v61
                                        }
                                    }
                                    v15 += float64(v2)
                                    if (201 < 127) {
                                        var v73 int32 = -918
                                        _ = v73
                                    }
                                } else {
                                    var v74 int = fi43
                                    v15 = 1e10
                                    _ = v74
                                }
                            case "":
                                switch ("\n\t" + ("" + ("" + "\n\t"))) {
                                case "a":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "foo bar":
                                    var v75 uint8 = 42
                                    _ = v75
                                case "hello":
                                    if (v14 >= 21) {
                                        var v76 int32 = -1
                                        _ = v76
                                    }
                                    v15 -= 1.5
                                }
                                fmt.Println(fi37, v14)
                                switch "\n\t" {
                                case "\n\t":
                                    switch "" {
                                    case "hello":
                                        fmt.Println(fi43)
                                    case "":
                                        v77 := GenS11{F0: 42, F1: v14}
                                        _ = v77
                                    }
                                case "foo bar":
                                    v2 = (float32(-2.25) / ((float32(1e10) + 1.5) + ((float32(1.5) / v2) / v2)))
                                }
                            default:
                                if (0.0 <= v2) {
                                    fmt.Println(v14, v2, v2)
                                }
                            }
                        case "hello":
                            var v78 int = fi43
                            _ = v78
                        }
                        var v79 int8 = -25
                        _ = v79
                    }
                    fmt.Println(fi37)
                    _ = fi43
                    _ = v44
                }
            }
            v80 := GenS12{F0: (fi37 >= (fi37 % (((fi37 / ((fi37) | 1))) | 1))), F1: v14}
            v2 -= v2
            _ = v80
        }
        _ = fi37
    }
    switch 128 {
    case -4:
        switch "foo bar" {
        case "a":
            var v81 int16 = 950
            _ = v81
        default:
            var v82 float32 = float32(0.1)
            _ = v82
        }
    default:
        var v83 string = "a"
        _ = v83
    }
    v2 += v2
    fmt.Println(v15, v14)
    if (-9223372036854775808 > 0) {
        if (v2 < (v2 / v2)) {
            var v84 int32 = int32(v14)
            fmt.Println(v2, v2, v14)
            fmt.Println(v84, v2, v1)
            _ = v84
        }
    } else {
        v3 = append(v3, ("" + (("a" + ("a" + "hello")) + ("foo bar" + ("foo bar" + "")))))
        switch -892 {
        case -1:
            var v85 uint32 = 1
            _ = v85
        case 127:
            if (-535 < 9223372036854775807) {
                v86 := GenS13{F0: "a", F1: (v1 / (((^(v1))) | 1))}
                for fi87 := 0; fi87 < 20; fi87++ {
                    switch (fi87 + -2147483648) {
                    case -201:
                        for fi88 := 0; fi88 < 17; fi88++ {
                            v14 -= v14
                            v15 -= (0.1 * 0.0)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi88
                        }
                    case -876:
                        var v89 string = ("a" + (((("a" + "hello") + "hello") + "") + "foo bar"))
                        _ = v89
                    case 9223372036854775807:
                        var v90 float32 = -2.25
                        _ = v90
                    default:
                        v1 -= v1
                    }
                    _ = fi87
                }
                _ = v86
            }
        case 817:
            v91 := GenS5{F0: 0}
            _ = v91
        }
    }
    var v92 int32 = 0
    var v93 int16 = v14
    if (v2 > v2) {
        var v94 int = 0
        if (42 < 0) {
            v95 := GenS14{F0: v92, F1: 65535}
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
            fmt.Println(v93)
            _ = v95
        } else {
            v96 := GenS5{F0: 89}
            _ = v96
        }
        _ = v94
    } else {
        v92 = v92
    }
    v97 := GenS6{F0: -2147483648, F1: -1, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v92
    _ = v93
    _ = v97
}
