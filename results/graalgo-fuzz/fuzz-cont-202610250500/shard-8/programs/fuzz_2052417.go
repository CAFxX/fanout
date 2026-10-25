package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int; F1 string; F2 bool }
type GenS5 struct { F0 int; F1 uint32; F2 int }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int64; F1 bool; F2 int }

func main() {
    v1 := []string{"foo bar"}
    v2 := []uint32{42}
    var v3 float64 = 0.0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = ("foo bar" + "")
    }
    if true {
        if ((-9223372036854775808 >= -9223372036854775808) && false) {
            fmt.Println(v3)
            v4 := GenS1{F0: 18446744073709551615}
            _ = v4
        }
        var v5 float64 = (v3 / float64(1e10))
        switch ("\n\t" + "a") {
        case "foo bar":
            var v6 int64 = 42
            _ = v6
        }
        _ = v5
    } else {
        v3 = v3
        var v7 int8 = 1
        v3 = (3.14159 + v3)
        _ = v7
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 += ((v3 - (float64((1e10 + ((3.14159 - 3.14159) * ((0.1 * 1.5) + 1.5)))) * v3)) + (v3 + 1.5))
    switch (("foo bar" + ("foo bar" + ((("\n\t" + ((("a" + (("a" + "hello") + ("a" + ("foo bar" + "a")))) + "a") + "")) + "") + "foo bar"))) + "") {
    case "hello":
        v3 += v3
    case "foo bar":
        if true {
            v8 := GenS1{F0: 0}
            _ = v8
        } else {
            fmt.Println(v3)
        }
    case "\n\t":
        var v9 int64 = -1
        _ = v9
    }
    switch -67 {
    case 9223372036854775807:
        v3 = 1.5
    }
    if (("foo bar" + "") >= (("" + "hello") + "hello")) {
        fmt.Println(v3)
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 int16 = 815
            var v12 int = (fi10 * (fi10 ^ (fi10 / (((fi10 * fi10)) | 1))))
            if true {
                for fi13 := 0; fi13 < 20; fi13++ {
                    v11 = (^(v11))
                    _ = fi13
                }
                for fi14 := 0; fi14 < 10; fi14++ {
                    v15 := GenS1{F0: 1}
                    var v16 int16 = (v11 % (((v11 - (v11 >> ((424) & 63)))) | 1))
                    fmt.Println(v16)
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            } else {
                v11 -= -32768
                v11 -= 32767
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    if false {
        if true {
            if (0 <= 42) {
                _ = len(v1)
                var v17 float32 = float32(v3)
                for fi18 := 0; fi18 < 13; fi18++ {
                    fmt.Println(v3, v17, v3)
                    v19 := GenS1{F0: 4294967295}
                    for fi20 := 0; fi20 < 12; fi20++ {
                        var v21 uint32 = 858
                        switch "hello" {
                        case "\n\t":
                            v3 = v3
                        case "":
                            v22 := GenS2{F0: -687}
                            _ = v22
                        }
                        _ = fi20
                        _ = v21
                    }
                    _ = fi18
                    _ = v19
                }
                _ = v17
            } else {
                v3 -= v3
                var v23 uint32 = 4294967295
                v24 := GenS1{F0: 302}
                _ = v23
                _ = v24
            }
            var v25 int = -691
            _ = v25
        }
        var v26 bool = (!true)
        v27 := GenS1{F0: 541}
        _ = v26
        _ = v27
    } else {
        switch 32767 {
        case -1:
            v3 = float64(1.5)
        case 2147483647:
            var v28 int16 = 2
            _ = v28
        default:
            for fi29 := 0; fi29 < 5; fi29++ {
                if true {
                    switch 42 {
                    case 42:
                        v30 := GenS3{F0: -188, F1: 255, F2: 42}
                        _ = v30
                    case 619:
                        if true {
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                            if (!true) {
                                if false {
                                    var v31 uint64 = 42
                                    _ = v31
                                } else {
                                    var v32 float32 = 1e10
                                    if false {
                                        fmt.Println(v32, v3, v3)
                                        if false {
                                            v33 := GenS4{F0: 127, F1: ("" + ("" + ("hello" + ""))), F2: true}
                                            v32 = 0.1
                                            _ = v33
                                        } else {
                                            for fi34 := 0; fi34 < 6; fi34++ {
                                                v32 -= v32
                                                fmt.Println(fi29)
                                                fmt.Println(fi34)
                                                v32 = -2.25
                                                v35 := GenS5{F0: fi29, F1: 956, F2: fi34}
                                                _ = fi34
                                                _ = v35
                                            }
                                        }
                                        v3 -= 0.0
                                    }
                                    _ = v32
                                }
                                fmt.Println(fi29)
                                if len(v2) > 0 {
                                    v2[0] = 153
                                }
                            } else {
                                v3 -= ((v3 * (v3 * ((-2.25 - 1e10) - v3))) + float64(1e10))
                            }
                        } else {
                            v3 = v3
                            var v36 uint8 = 255
                            var v37 bool = false
                            _ = v36
                            _ = v37
                        }
                    case 128:
                        switch ("foo bar" + "") {
                        case "a":
                            for fi38 := 0; fi38 < 7; fi38++ {
                                v3 -= 0.0
                                if len(v2) > 0 {
                                    v2[0] = 4294967295
                                }
                                v39 := GenS1{F0: 42}
                                _ = fi38
                                _ = v39
                            }
                        case "":
                            var v40 bool = (74 != 2)
                            _ = v40
                        }
                    }
                    if false {
                        v41 := GenS3{F0: 666, F1: 654, F2: 18446744073709551615}
                        _ = v41
                    } else {
                        var v42 float32 = ((1e10 * 3.14159) * float32(v3))
                        var v43 float64 = float64(1e10)
                        for fi44 := 0; fi44 < 9; fi44++ {
                            var v45 uint64 = 880
                            fmt.Println(fi29)
                            var v46 int16 = -724
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                        _ = v42
                        _ = v43
                    }
                    var v47 int32 = 1
                    _ = v47
                }
                v3 = -2.25
                _ = fi29
            }
        }
        v3 = v3
        for fi48 := 0; fi48 < 6; fi48++ {
            switch "a" {
            case "foo bar":
                v49 := GenS3{F0: -1, F1: 65535, F2: 884}
                _ = v49
            }
            v3 = 1.5
            for fi50 := 0; fi50 < 14; fi50++ {
                v3 = (0.1 / v3)
                switch fi48 {
                case 1:
                    v51 := GenS1{F0: 42}
                    _ = v51
                case -9223372036854775808:
                    v3 = 3.14159
                case 9223372036854775807:
                    var v52 float32 = 3.14159
                    _ = v52
                }
                switch "\n\t" {
                case "a":
                    v3 = 1.5
                case "foo bar":
                    fmt.Println(v3, v3, fi50)
                }
                _ = fi50
            }
            _ = fi48
        }
    }
    if (!(0 == 42)) {
        if (!(32767 < 504)) {
            _ = len(v1)
            var v53 float64 = 3.14159
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v53
        } else {
            if true {
                switch -2147483648 {
                case 746:
                    v3 -= v3
                case 2147483647:
                    v3 += 3.14159
                default:
                    v3 = v3
                }
                if true {
                    v54 := GenS3{F0: 1, F1: 2, F2: 457}
                    if (0 != 42) {
                        switch -1 {
                        case 387:
                            switch "a" {
                            case "foo bar":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                var v55 uint = 2
                                _ = v55
                            }
                        case 42:
                            v3 = (0.0 - float64(3.14159))
                        case 128:
                            for fi56 := 0; fi56 < 7; fi56++ {
                                switch (("\n\t" + ("hello" + "a")) + "") {
                                case "a":
                                    fmt.Println(v3)
                                case "foo bar":
                                    switch (^(fi56)) {
                                    case 128:
                                        switch fi56 {
                                        case 4294967295:
                                            for fi57 := 0; fi57 < 3; fi57++ {
                                                if false {
                                                    v58 := GenS6{F0: 487, F1: -393}
                                                    if (true && false) {
                                                        if len(v1) > 0 {
                                                            v1[0] = "hello"
                                                        }
                                                        var v59 int32 = 39
                                                        var v60 bool = false
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                    _ = v58
                                                } else {
                                                    var v61 int32 = -2147483648
                                                    _ = v61
                                                }
                                                _ = fi57
                                            }
                                        case -9223372036854775808:
                                            if ((!(("\n\t" != "") || (!false))) && (!((100 >= 0) && true))) {
                                                var v62 float64 = v3
                                                _ = v62
                                            }
                                        }
                                    case -734:
                                        v3 -= (1.5 * (v3 + (((1.5 / ((v3 + v3) - v3)) / 1e10) * v3)))
                                    }
                                default:
                                    _ = len(v2)
                                }
                                v63 := GenS7{F0: (fi56 % (((fi56 ^ fi56)) | 1))}
                                _ = fi56
                                _ = v63
                            }
                        default:
                            v3 -= v3
                        }
                    }
                    _ = v54
                }
            }
            if true {
                v3 = v3
            } else {
                for fi64 := 0; fi64 < 14; fi64++ {
                    v3 = (v3 / ((v3 / (v3 - 1.5)) - float64((1.5 + ((0.0 + 1e10) * 1e10)))))
                    _ = fi64
                }
                v3 += v3
            }
            var v65 string = ("\n\t" + "")
            _ = v65
        }
        if (0 == 1) {
            var v66 float32 = (0.1 / 1e10)
            v3 = v3
            _ = v66
        }
    } else {
        v3 = v3
    }
    for fi67 := 0; fi67 < 9; fi67++ {
        v68 := GenS8{F0: 42}
        _ = fi67
        _ = v68
    }
    switch -9223372036854775808 {
    case -485:
        var v69 bool = true
        _ = v69
    case 127:
        if true {
            fmt.Println(v3)
        } else {
            if false {
                v3 += v3
            } else {
                var v70 uint16 = 1
                switch -155 {
                case -144:
                    v71 := GenS5{F0: 128, F1: 42, F2: 478}
                    _ = v71
                }
                for fi72 := 0; fi72 < 11; fi72++ {
                    var v73 int32 = 42
                    v3 = v3
                    _ = fi72
                    _ = v73
                }
                _ = v70
            }
            switch "" {
            case "":
                var v74 uint32 = 42
                _ = v74
            }
        }
    case -983:
        switch 42 {
        case 2147483647:
            _ = len(v2)
        case -1:
            for fi75 := 0; fi75 < 6; fi75++ {
                fmt.Println(fi75)
                _ = fi75
            }
        case 32767:
            var v76 float64 = v3
            _ = v76
        default:
            fmt.Println(v3)
        }
    default:
        v77 := GenS9{F0: -1, F1: (358 > 2), F2: 881}
        _ = v77
    }
    var v78 int = -184
    var v79 float64 = (((0.1 * float64(3.14159)) + v3) - (((v3 / v3) - ((v3 / 1.5) - v3)) * ((float64(-2.25) * v3) + -2.25)))
    fmt.Println(v3)
    if false {
        v79 = v3
        v80 := GenS3{F0: 563, F1: 0, F2: 564}
        _ = v80
    }
    v81 := GenS2{F0: -1}
    for fi82 := 0; fi82 < 16; fi82++ {
        v78 = 128
        for fi83 := 0; fi83 < 11; fi83++ {
            v79 = (0.0 + (float64((3.14159 * 0.0)) - (float64(0.1) + 3.14159)))
            fmt.Println(v3, fi82)
            _ = fi83
        }
        _ = fi82
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v78
    _ = v79
    _ = v81
}
