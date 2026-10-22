package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int32; F1 uint32 }
type GenS14 struct { F0 string; F1 uint8 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = -2147483648
    var v3 int32 = v2
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := []int{128, fi4}
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 int8 = 0
            switch "foo bar" {
            case "foo bar":
                var v8 uint8 = 150
                _ = v8
            case "":
                if (false && true) {
                    if false {
                        if false {
                            v7 = 127
                        } else {
                            v9 := GenS1{F0: 0, F1: "hello", F2: -253}
                            _ = v9
                        }
                        if false {
                            v10 := GenS2{F0: 65535}
                            if false {
                                var v11 uint32 = 2
                                _ = v11
                            }
                            if (2147483647 > 1) {
                                v12 := v1["foo bar"]
                                var v13 uint8 = 243
                                switch v12 {
                                case "\n\t":
                                    var v14 uint = 255
                                    var v15 int = -2147483648
                                    v15 = (fi4 % (((fi6 | (v15 / ((v15) | 1)))) | 1))
                                    _ = v14
                                    _ = v15
                                case "foo bar":
                                    if (65535 != 32767) {
                                        var v16 int = (^(fi4))
                                        v12 = ("\n\t" + (v12 + (v12 + v12)))
                                        _ = v16
                                    }
                                }
                                _ = v12
                                _ = v13
                            } else {
                                var v17 int64 = 9223372036854775807
                                _ = v17
                            }
                            _ = v10
                        } else {
                            v7 -= 31
                        }
                    }
                    if false {
                        switch fi6 {
                        case -81:
                            v3 += 2147483647
                        case 127:
                            switch -2147483648 {
                            case 127:
                                v7 = v7
                            }
                        case 2147483647:
                            if true {
                                v18 := GenS3{F0: (-(fi4)), F1: 18446744073709551615}
                                fmt.Println(fi6)
                                _ = v18
                            }
                        }
                        if false {
                            if true {
                                v19 := GenS4{F0: 42}
                                _ = v19
                            }
                            switch (-(fi6)) {
                            case -9223372036854775808:
                                v20 := GenS5{F0: (^(v2)), F1: 42}
                                _ = v20
                            case -926:
                                if len(v5) > 0 {
                                    v5[0] = (fi4 % (((fi6 - fi4)) | 1))
                                }
                            }
                            v7 = v7
                        } else {
                            if (fi4 > fi6) {
                                fmt.Println(fi4, v3)
                                var v21 uint64 = 545
                                v2 += -140
                                if ((("hello" + "") < "\n\t") && true) {
                                    v21 += (v21 | (v21 % ((v21) | 1)))
                                    v2 += 42
                                } else {
                                    var v22 int64 = 42
                                    v2 = (^(v3))
                                    _ = v22
                                }
                                v23 := GenS2{F0: 893}
                                _ = v21
                                _ = v23
                            } else {
                                var v24 uint = 585
                                _ = v24
                            }
                            var v25 uint32 = 744
                            _ = len(v5)
                            var v26 int = fi4
                            _ = len(v1)
                            _ = v25
                            _ = v26
                        }
                    }
                    var v27 uint64 = 302
                    if true {
                        if false {
                            var v28 int16 = -79
                            _ = v28
                        } else {
                            v3 -= (v3 - (v2 / ((0) | 1)))
                            v7 += (v7 % ((1) | 1))
                        }
                    } else {
                        var v29 int64 = -9223372036854775808
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v27 += 904
                        var v30 bool = ((!false) && (!true))
                        _ = v29
                        _ = v30
                    }
                    if (-847 >= 0) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    _ = v27
                } else {
                    fmt.Println(fi4)
                    v2 = (-(v2))
                    var v31 string = "hello"
                    _ = v31
                }
            }
            _ = fi6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    _ = len(v1)
    v32 := GenS4{F0: 536}
    if false {
        v1[(("" + ("a" + "a")) + "hello")] = ""
        var v33 uint16 = 682
        v33 += (^(v33))
        _ = v33
    } else {
        if false {
            var v34 int = 4294967295
            for fi35 := 0; fi35 < 13; fi35++ {
                for fi36 := 0; fi36 < 2; fi36++ {
                    v37 := []int{(v34 - fi36), 128, fi35}
                    var v38 int16 = -281
                    v34 += (v34 % ((9223372036854775807) | 1))
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                v39 := GenS3{F0: v34, F1: 0}
                _ = fi35
                _ = v39
            }
            _ = v34
        } else {
            for fi40 := 0; fi40 < 10; fi40++ {
                var v41 uint32 = 0
                _ = fi40
                _ = v41
            }
            if true {
                v3 = v2
            }
        }
        fmt.Println(v3)
        v3 -= v2
    }
    var v42 int = 0
    v3 = int32(v42)
    v42 += (v42 + (v42 % (((v42 % ((127) | 1))) | 1)))
    v3 = (v3 * v2)
    for fi43 := 0; fi43 < 4; fi43++ {
        v44 := GenS6{F0: 255}
        var v45 int64 = 1
        _ = fi43
        _ = v44
        _ = v45
    }
    var v46 uint64 = 2
    fmt.Println(v46, v46, v46)
    var v47 int16 = 1000
    v47 = 304
    var v48 int32 = (v3 % (((v3 + v2)) | 1))
    var v49 int64 = 9223372036854775807
    v50 := GenS5{F0: v2, F1: 1}
    switch (v42 >> ((18446744073709551615) & 63)) {
    case -9223372036854775808:
        if true {
            v51 := GenS7{F0: 881}
            var v52 int = -2147483648
            _ = v51
            _ = v52
        }
    case 0:
        v53 := GenS8{F0: v46, F1: (v47 + 1000), F2: 1}
        _ = v53
    case 779:
        switch (v42 / ((2147483647) | 1)) {
        case -1:
            var v54 float64 = 0.1
            _ = v54
        case -273:
            v47 += v47
        default:
            switch "\n\t" {
            case "a":
                var v55 uint64 = v46
                _ = v55
            default:
                switch v42 {
                case 9223372036854775807:
                    for fi56 := 0; fi56 < 10; fi56++ {
                        var v57 float32 = ((float32(((-2.25 + 1.5) * -2.25)) / 3.14159) / 0.1)
                        _ = fi56
                        _ = v57
                    }
                default:
                    v3 -= int32(v42)
                }
            }
        }
    }
    if (!true) {
        for fi58 := 0; fi58 < 19; fi58++ {
            if (42 == 1) {
                fmt.Println(fi58, v48, v3)
                for fi59 := 0; fi59 < 7; fi59++ {
                    _ = len(v1)
                    _ = fi59
                }
            } else {
                switch (("a" + "hello") + "a") {
                case "hello":
                    if (2147483647 != 65535) {
                        var v60 int32 = (v48 % ((-1) | 1))
                        switch (^(fi58)) {
                        case 4294967295:
                            v42 = v42
                        case 0:
                            if true {
                                if true {
                                    v60 = (v60 + 0)
                                    switch "" {
                                    case "hello":
                                        v1[("" + ((("\n\t" + "") + (("" + ("\n\t" + "foo bar")) + ("" + ("a" + (((("" + "hello") + (((("hello" + "a") + "") + "\n\t") + "\n\t")) + ("" + ("" + (("a" + ("foo bar" + "")) + "")))) + ("hello" + ("" + "foo bar"))))))) + "\n\t"))] = "hello"
                                    case "a":
                                        var v61 int = (-(v42))
                                        _ = v61
                                    case "foo bar":
                                        v62 := GenS1{F0: v46, F1: "\n\t", F2: (v49 >> ((444) & 63))}
                                        _ = v62
                                    }
                                } else {
                                    v63 := GenS4{F0: 0}
                                    v2 += v48
                                    v64 := GenS9{F0: 128}
                                    _ = v63
                                    _ = v64
                                }
                                for fi65 := 0; fi65 < 19; fi65++ {
                                    fmt.Println(fi58)
                                    v66 := GenS10{F0: (-(v49))}
                                    _ = fi65
                                    _ = v66
                                }
                            } else {
                                v67 := GenS11{F0: (^(v3))}
                                switch "" {
                                case "":
                                    var v68 string = "\n\t"
                                    _ = v68
                                }
                                switch ("foo bar" + ("\n\t" + "a")) {
                                case "hello":
                                    v47 -= v47
                                }
                                _ = v67
                            }
                        case 1:
                            var v69 int16 = (v47 ^ (v47 % ((204) | 1)))
                            _ = v69
                        }
                        _ = v60
                    }
                case "":
                    if false {
                        v46 = (v46 | v46)
                        for fi70 := 0; fi70 < 8; fi70++ {
                            if (144 != 359) {
                                switch "\n\t" {
                                case "":
                                    v71 := GenS12{F0: "foo bar"}
                                    v72 := []int{fi70, (fi58 + fi58), (v42 % ((v42) | 1)), v42}
                                    _ = v71
                                    _ = v72
                                case "\n\t":
                                    var v73 int8 = -1
                                    _ = v73
                                case "foo bar":
                                    var v74 int32 = v48
                                    _ = v74
                                default:
                                    if false {
                                        var v75 uint16 = 131
                                        v76 := GenS13{F0: v48, F1: 991}
                                        if ((v75 * v75) != v75) {
                                            v77 := GenS6{F0: 4294967295}
                                            switch "hello" {
                                            case "hello":
                                                var v78 int32 = int32(v49)
                                                _ = v78
                                            case "":
                                                switch 0 {
                                                case -336:
                                                    v46 = v46
                                                }
                                            case "a":
                                                v79 := []int{fi70, (^(v42))}
                                                _ = v79
                                            }
                                            fmt.Println(v49, fi70, v49)
                                            v80 := GenS14{F0: "", F1: 42}
                                            _ = v77
                                            _ = v80
                                        } else {
                                            v75 = 2
                                            _ = len(v1)
                                            v75 = (v75 & 142)
                                        }
                                        _ = v75
                                        _ = v76
                                    }
                                }
                                var v81 uint16 = 1
                                var v82 int64 = v49
                                fmt.Println(v2)
                                _ = v81
                                _ = v82
                            } else {
                                switch ("a" + ("foo bar" + "foo bar")) {
                                case "":
                                    v2 = v3
                                case "hello":
                                    switch (("a" + "") + ("a" + "a")) {
                                    case "hello":
                                        v83 := GenS7{F0: 0}
                                        _ = v83
                                    case "":
                                        if (true || (((0.1 / (3.14159 / float32(1e10))) - 0.1) > 1.5)) {
                                            v3 += (v48 >> ((558) & 63))
                                            if (32767 != 186) {
                                                switch fi58 {
                                                case -9223372036854775808:
                                                    var v84 uint16 = 514
                                                    _ = v84
                                                }
                                            } else {
                                                if false {
                                                    var v85 int8 = 127
                                                    switch (fi58 % (((fi58 | fi58)) | 1)) {
                                                    case -9223372036854775808:
                                                        v86 := GenS14{F0: ("hello" + "foo bar"), F1: 0}
                                                        _ = v86
                                                    default:
                                                        v3 += 640
                                                    }
                                                    _ = len(v1)
                                                    _ = v85
                                                }
                                                var v87 float64 = 0.0
                                                delete(v1, "")
                                                _ = v87
                                            }
                                            v88 := GenS12{F0: "foo bar"}
                                            _ = v88
                                        } else {
                                            v89 := GenS3{F0: fi58, F1: (v46 - v46)}
                                            _ = v89
                                        }
                                    default:
                                        switch ((("" + ("hello" + "a")) + "hello") + ("hello" + "\n\t")) {
                                        case "":
                                            if true {
                                                var v90 bool = ((float32(((1e10 - float64(3.14159)) - 1e10)) + 1e10) != 0.0)
                                                v3 -= (v3 - (-(v3)))
                                                v46 -= v46
                                                _ = v90
                                            }
                                        case "foo bar":
                                            delete(v1, "a")
                                        default:
                                            fmt.Println(v47)
                                        }
                                    }
                                }
                                fmt.Println(v3, fi70, v46)
                            }
                            v47 = (^(v47))
                            _ = fi70
                        }
                    }
                }
            }
            _ = fi58
        }
        var v91 int64 = (v49 / (((v49 * (v49 << ((472) & 63)))) | 1))
        _ = v91
    } else {
        fmt.Println(v48)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v42
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
