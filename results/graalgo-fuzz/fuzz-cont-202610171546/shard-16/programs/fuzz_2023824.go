package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS10 struct { F0 int16; F1 int32; F2 string }
type GenS11 struct { F0 uint32; F1 uint64 }
type GenS12 struct { F0 uint8; F1 int8 }
type GenS13 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS14 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint32) int16 {
    v4 := GenS1{F0: 42, F1: 225}
    p0 -= 65535
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    var v1 bool = true
    var v2 float64 = (1e10 / 1e10)
    v3 := []string{}
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 float64 = ((1e10 * (1e10 * v2)) / 1.5)
        switch 4294967295 {
        case 9:
            v6 = float64(float32(v2))
        case 2147483647:
            for fi7 := 0; fi7 < 13; fi7++ {
                switch (fi5 >> ((255) & 63)) {
                case -1:
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                default:
                    v8 := GenS1{F0: 42, F1: 2}
                    v9 := GenS1{F0: -42, F1: 251}
                    _ = v8
                    _ = v9
                }
                switch ("foo bar" + (("\n\t" + "hello") + "")) {
                case "foo bar":
                    if true {
                        v6 = ((v2 * v2) / v2)
                        if ((523 == 577) || (!true)) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = len(v3)
                            if (127 != 2) {
                                var v10 float64 = 1.5
                                var v11 uint16 = 32767
                                fmt.Println(v11, v11)
                                _ = v10
                                _ = v11
                            }
                            var v12 int32 = -1
                            if v1 {
                                if len(v3) > 0 {
                                    v3[0] = "hello"
                                }
                                fmt.Println(fi7)
                                var v13 uint16 = 32767
                                v3 = append(v3, ("hello" + (("foo bar" + ("foo bar" + (("\n\t" + "a") + (("\n\t" + "hello") + "\n\t")))) + "hello")))
                                _ = v13
                            } else {
                                var v14 bool = (fi7 > (^(fi7)))
                                _ = v14
                            }
                            switch 42 {
                            case -1:
                                fmt.Println(v1, v12)
                            default:
                                v15 := GenS2{F0: 100, F1: 458, F2: 42}
                                _ = v15
                            }
                            switch fi7 {
                            case 128:
                                v16 := GenS3{F0: 218}
                                _ = v16
                            case 1:
                                if v1 {
                                    var v17 uint64 = 42
                                    switch 854 {
                                    case 127:
                                        v3 = append(v3, ("hello" + "foo bar"))
                                    case 1:
                                        v6 = v2
                                    }
                                    _ = v17
                                } else {
                                    var v18 bool = true
                                    v1 = (0.1 == v6)
                                    if v1 {
                                        var v19 uint8 = 0
                                        _ = v19
                                    }
                                    _ = v18
                                }
                            }
                            _ = v12
                        }
                        if v1 {
                            if v1 {
                                var v20 int64 = 42
                                var v21 uint8 = 227
                                _ = v20
                                _ = v21
                            }
                            switch 0 {
                            case 32767:
                                var v22 int32 = 1
                                _ = v22
                            default:
                                v23 := GenS2{F0: -128, F1: 42, F2: 473}
                                if (!v1) {
                                    switch ("foo bar" + ("foo bar" + "hello")) {
                                    case "foo bar":
                                        var v24 int = (fi7 + fi5)
                                        switch -954 {
                                        case -1:
                                            v25 := GenS4{F0: 42, F1: "\n\t"}
                                            _ = v25
                                        case 128:
                                            if v1 {
                                                v24 = (fi5 + 9223372036854775807)
                                                v26 := GenS2{F0: -61, F1: 1, F2: -9223372036854775808}
                                                var v27 int8 = 42
                                                _ = v26
                                                _ = v27
                                            }
                                        }
                                        _ = v24
                                    case "hello":
                                        if (4294967295 != 0) {
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v28 uint64 = 0
                                                _ = v28
                                            case "hello":
                                                v29 := GenS5{F0: 2, F1: 1}
                                                _ = v29
                                            case "a":
                                                var v30 float64 = (v6 * v2)
                                                _ = v30
                                            }
                                        } else {
                                            v2 -= (((((((v2 - 1.5) - v2) - v6) * v6) * v6) / 1e10) + v2)
                                            var v31 float64 = (((((v6 * v6) / 0.1) / 1.5) / v6) + (v2 / v2))
                                            if v1 {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            } else {
                                                var v32 int8 = -1
                                                var v33 int16 = 32767
                                                fmt.Println(fi7, fi7)
                                                if true {
                                                    v34 := GenS4{F0: 42, F1: "hello"}
                                                    v33 = (^(v33))
                                                    _ = v34
                                                }
                                                _ = v32
                                                _ = v33
                                            }
                                            _ = v31
                                        }
                                    case "a":
                                        if (((!(1 <= (fi7 >> ((554) & 63)))) || (1000 <= 292)) && false) {
                                            if len(v3) > 0 {
                                                v3[0] = ""
                                            }
                                            var v35 uint = 42
                                            if (!(9223372036854775807 == 9223372036854775807)) {
                                                fmt.Println(fi7, v35)
                                                v36 := GenS6{F0: 578}
                                                v35 = (v35 * (v35 / ((4294967295) | 1)))
                                                _ = v36
                                            }
                                            _ = v35
                                        }
                                        v1 = (!(-32768 > -32768))
                                    default:
                                        v6 += float64(1.5)
                                    }
                                }
                                switch -2147483648 {
                                case 161:
                                    switch 972 {
                                    case 42:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case 0:
                                        v1 = true
                                    case 450:
                                        var v37 bool = v1
                                        _ = v37
                                    default:
                                        fmt.Println(v2, v6, fi5)
                                    }
                                case 2147483647:
                                    v38 := GenS7{F0: v1, F1: 2147483647}
                                    _ = v38
                                default:
                                    v39 := GenS6{F0: 946}
                                    _ = v39
                                }
                                _ = v23
                            }
                            v2 -= v6
                        } else {
                            var v40 int64 = -1
                            _ = v40
                        }
                    }
                case "hello":
                    v6 += (v6 * 3.14159)
                default:
                    switch (-(fi7)) {
                    case 9223372036854775807:
                        v41 := GenS8{F0: 777}
                        _ = v41
                    case 127:
                        if v1 {
                            v6 -= (v2 * (0.0 / v2))
                            var v42 uint = 993
                            fmt.Println(v6, fi7)
                            _ = v42
                        }
                    case 42:
                        v43 := GenS6{F0: 419}
                        _ = v43
                    default:
                        var v44 int = (fi7 ^ fi5)
                        _ = v44
                    }
                }
                var v45 uint16 = 460
                var v46 uint = 0
                _ = fi7
                _ = v45
                _ = v46
            }
        case 9223372036854775807:
            if false {
                var v47 uint64 = 526
                v48 := GenS9{F0: 0, F1: 2147483648, F2: 699}
                for fi49 := 0; fi49 < 12; fi49++ {
                    v50 := GenS4{F0: 501, F1: "hello"}
                    if false {
                        v51 := GenS5{F0: 100, F1: -367}
                        v6 += v6
                        v3 = append(v3, ("foo bar" + ("foo bar" + ((("foo bar" + ("foo bar" + ("a" + "a"))) + ("\n\t" + "foo bar")) + "foo bar"))))
                        _ = v51
                    }
                    switch ("foo bar" + ("foo bar" + ("\n\t" + "hello"))) {
                    case "foo bar":
                        var v52 float64 = (v2 / ((float64(1.5) + (float64(1.5) * v2)) / v2))
                        _ = v52
                    case "\n\t":
                        v6 = (3.14159 / (v6 + float64(float32(-2.25))))
                        switch fi49 {
                        case -1:
                            var v53 int8 = 1
                            _ = v53
                        case 897:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                    _ = fi49
                    _ = v50
                }
                _ = v47
                _ = v48
            } else {
                v6 = v6
                _ = len(v3)
            }
        }
        _ = fi5
        _ = v6
    }
    var v54 uint16 = 32767
    var v55 uint8 = 67
    v2 = v2
    _ = len(v3)
    v55 = 89
    for fi56 := 0; fi56 < 16; fi56++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(fi56, v54)
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        _ = fi56
    }
    v3 = append(v3, "a")
    var v57 uint32 = 2147483647
    for fi58 := 0; fi58 < 6; fi58++ {
        var v59 bool = (4294967295 < 4294967295)
        v60 := GenS5{F0: 42, F1: -566}
        _ = fi58
        _ = v59
        _ = v60
    }
    v54 = (v54 << ((18446744073709551615) & 63))
    for fi61 := 0; fi61 < 16; fi61++ {
        if v1 {
            v62 := GenS5{F0: v55, F1: 671}
            if true {
                var v63 int64 = 9223372036854775807
                v55 += v55
                if v1 {
                    v1 = v1
                    if (!false) {
                        v64 := GenS7{F0: v1, F1: 0}
                        v55 = 128
                        _ = v64
                    } else {
                        if len(v3) > 0 {
                            v3[0] = ("a" + "a")
                        }
                    }
                    if (2147483647 >= v57) {
                        v65 := GenS3{F0: 194}
                        fmt.Println(v54, v55)
                        fmt.Println(v1, fi61)
                        _ = v65
                    } else {
                        v66 := GenS10{F0: -203, F1: -2147483648, F2: ("a" + "\n\t")}
                        _ = v66
                    }
                } else {
                    var v67 bool = (!false)
                    _ = v67
                }
                _ = v63
            } else {
                v57 = (^(v57))
            }
            _ = v62
        } else {
            var v68 uint16 = (v54 | uint16(v55))
            switch "\n\t" {
            case "a":
                fmt.Println(v55, v57)
            case "foo bar":
                switch "hello" {
                case "":
                    if (-9223372036854775808 < -609) {
                        if (!v1) {
                            if (v2 > (v2 + ((((v2 / (v2 * v2)) / ((float64(3.14159) * float64(0.0)) - (float64((3.14159 - 1.5)) * (0.1 / v2)))) * (float64(0.0) / v2)) / v2))) {
                                v69 := GenS5{F0: v55, F1: 444}
                                _ = v69
                            }
                            var v70 uint = 1
                            _ = v70
                        }
                        v71 := GenS11{F0: v57, F1: 283}
                        switch fi61 {
                        case 128:
                            if false {
                                var v72 float64 = ((0.0 * (v2 * ((((3.14159 / (-2.25 * float64(3.14159))) - v2) / v2) / (((1e10 * (v2 / float64(3.14159))) / v2) + v2)))) / (3.14159 * float64(1e10)))
                                _ = v72
                            }
                        case -954:
                            if (v57 < (v57 / ((v57) | 1))) {
                                if (1 < 222) {
                                    v73 := GenS12{F0: uint8(v57), F1: 42}
                                    _ = v73
                                } else {
                                    if v1 {
                                        v74 := GenS13{F0: uint16(v57), F1: 334, F2: 1}
                                        v3 = append(v3, "")
                                        _ = v74
                                    }
                                }
                                v1 = true
                                fmt.Println(v68)
                            }
                        case 32767:
                            v1 = false
                        default:
                            var v75 bool = true
                            _ = v75
                        }
                        _ = v71
                    } else {
                        var v76 bool = v1
                        var v77 uint64 = 1
                        _ = v76
                        _ = v77
                    }
                case "\n\t":
                    v68 += 650
                default:
                    var v78 int16 = 0
                    _ = v78
                }
            case "\n\t":
                switch int(v57) {
                case 459:
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                case 9223372036854775807:
                    v57 += 2147483648
                case 2147483647:
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                default:
                    v79 := GenS9{F0: (v55 / ((v55) | 1)), F1: uint32(v54), F2: 2}
                    _ = v79
                }
            default:
                var v80 string = "foo bar"
                _ = v80
            }
            _ = v68
        }
        v81 := GenS8{F0: 65535}
        _ = fi61
        _ = v81
    }
    if v1 {
        for fi82 := 0; fi82 < 15; fi82++ {
            v55 = v55
            if false {
                for fi83 := 0; fi83 < 13; fi83++ {
                    var v84 uint32 = v57
                    var v85 int = fi83
                    if len(v3) > 0 {
                        v3[0] = ("" + (("" + "\n\t") + "foo bar"))
                    }
                    _ = fi83
                    _ = v84
                    _ = v85
                }
                v2 = v2
                _ = len(v3)
            }
            _ = fi82
        }
        v86 := GenS1{F0: 0, F1: v55}
        _ = v86
    } else {
        if (-402 != -1) {
            v87 := GenS14{F0: 18446744073709551615}
            _ = v87
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v55
    _ = v57
}
