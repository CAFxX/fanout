package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint }
type GenS2 struct { F0 int64; F1 int64; F2 int }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint16; F1 uint64; F2 string }
type GenS6 struct { F0 int; F1 bool; F2 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16; F1 int64; F2 string }

func main() {
    var v1 int16 = 32767
    v2 := []uint32{}
    v3 := []string{"a", "a"}
    v4 := []string{"foo bar", ("a" + "a")}
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS1{F0: "\n\t", F1: 0, F2: 19}
        fmt.Println(v1, v1)
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            fmt.Println(fi7, fi7)
            if false {
                var v9 int8 = 127
                _ = v9
            } else {
                v1 -= v1
                v10 := GenS1{F0: ((("hello" + ("foo bar" + "hello")) + "foo bar") + ("\n\t" + ("" + (("a" + (("foo bar" + ((("\n\t" + ("foo bar" + ("hello" + "a"))) + "foo bar") + (("hello" + "foo bar") + "hello"))) + "\n\t")) + "\n\t")))), F1: 42, F2: 255}
                _ = v10
            }
            _ = fi8
        }
        for fi11 := 0; fi11 < 3; fi11++ {
            v12 := GenS1{F0: ("hello" + "a"), F1: 659, F2: 818}
            var v13 int16 = (v1 >> ((18446744073709551615) & 63))
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi7
    }
    v1 = v1
    for fi14 := 0; fi14 < 16; fi14++ {
        v1 += (v1 ^ 32767)
        switch "a" {
        case "":
            var v15 uint64 = 2
            _ = v15
        case "hello":
            switch "foo bar" {
            case "":
                v1 += v1
            }
        case "a":
            var v16 float32 = (0.1 * 1e10)
            _ = v16
        default:
            fmt.Println(fi14, v1)
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 int64 = 9223372036854775807
            _ = fi17
            _ = v18
        }
        _ = fi14
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 float64 = float64((((0.1 - float32(-2.25)) - ((-2.25 * float32((1e10 * (-2.25 / 1.5)))) + 1e10)) / ((0.0 - (-2.25 * 1.5)) / -2.25)))
        _ = fi19
        _ = v20
    }
    v1 = v1
    var v21 float32 = ((1e10 + 0.1) + 1.5)
    switch 4294967295 {
    case 0:
        if (9223372036854775807 >= 9223372036854775807) {
            for fi22 := 0; fi22 < 8; fi22++ {
                for fi23 := 0; fi23 < 19; fi23++ {
                    fmt.Println(v21, v21, fi23)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v2 = append(v2, 2)
                    v1 = v1
                    _ = fi23
                }
                var v24 float64 = (3.14159 / ((1.5 / 1e10) + ((1.5 - 1e10) / -2.25)))
                v25 := GenS1{F0: ((("a" + "a") + ("a" + (("" + ((("\n\t" + "") + "foo bar") + "\n\t")) + "a"))) + ("hello" + "\n\t")), F1: 42, F2: 4294967295}
                _ = fi22
                _ = v24
                _ = v25
            }
            v26 := GenS2{F0: -9223372036854775808, F1: 9223372036854775807, F2: 1}
            _ = v26
        }
    default:
        var v27 int16 = 1000
        _ = v27
    }
    v21 = ((1.5 * v21) / v21)
    var v28 uint8 = 55
    v4 = append(v4, "a")
    if true {
        if (2 == v28) {
            switch 127 {
            case 128:
                if (false && true) {
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    if (240 < 606) {
                        var v29 int8 = 78
                        fmt.Println(v1, v21)
                        v30 := GenS3{F0: (v1 + v1), F1: 115}
                        _ = v29
                        _ = v30
                    } else {
                        if ((1 <= v28) || (!true)) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            for fi31 := 0; fi31 < 19; fi31++ {
                                if false {
                                    if false {
                                        var v32 string = "\n\t"
                                        _ = v32
                                    } else {
                                        if (0 > 2147483647) {
                                            if false {
                                                switch fi31 {
                                                case -2147483648:
                                                    v28 -= v28
                                                case 39:
                                                    switch 4294967295 {
                                                    case 758:
                                                        v21 = float32(((3.14159 + 1e10) / 1.0))
                                                    case 598:
                                                        v28 = 0
                                                    default:
                                                        var v33 int32 = 920
                                                        _ = v33
                                                    }
                                                default:
                                                    switch fi31 {
                                                    case 32768:
                                                        fmt.Println(v21)
                                                    case -9223372036854775808:
                                                        for fi34 := 0; fi34 < 7; fi34++ {
                                                            v3 = append(v3, "\n\t")
                                                            v35 := GenS4{F0: 4294967295, F1: true, F2: 65535}
                                                            v36 := GenS2{F0: 42, F1: 1, F2: -2147483648}
                                                            _ = fi34
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                    }
                                                }
                                                var v37 float64 = float64(v21)
                                                _ = v37
                                            }
                                            switch ("a" + "\n\t") {
                                            case "a":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            v21 += ((v21 / ((v21 + (((float32(-2.25) + v21) / ((v21 + (((((0.1 + 0.1) - v21) + (v21 / 1.5)) + 1.5) + v21)) * v21)) * (v21 + 1.5))) * (v21 - v21))) / (-2.25 - (v21 - 1.5)))
                                        } else {
                                            fmt.Println(v28)
                                        }
                                        for fi38 := 0; fi38 < 15; fi38++ {
                                            var v39 string = ""
                                            if true {
                                                v39 = v39
                                                if true {
                                                    var v40 uint16 = 2
                                                    _ = v40
                                                } else {
                                                    v41 := GenS5{F0: 823, F1: 4294967295, F2: v39}
                                                    if (float64((3.14159 * 1e10)) != 1.5) {
                                                        var v42 float32 = v21
                                                        var v43 int8 = -128
                                                        _ = v42
                                                        _ = v43
                                                    } else {
                                                        fmt.Println(fi38)
                                                        v44 := GenS5{F0: 32767, F1: 2, F2: v39}
                                                        v45 := GenS1{F0: "", F1: 179, F2: 18446744073709551615}
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                    _ = v41
                                                }
                                                fmt.Println(v21)
                                                _ = len(v4)
                                            }
                                            switch v39 {
                                            case "foo bar":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            default:
                                                var v46 int = 486
                                                _ = v46
                                            }
                                            _ = fi38
                                            _ = v39
                                        }
                                        var v47 int8 = 42
                                        _ = v47
                                    }
                                } else {
                                    for fi48 := 0; fi48 < 4; fi48++ {
                                        var v49 int = (fi31 ^ -465)
                                        switch (v49 & v49) {
                                        case 2147483647:
                                            if true {
                                                v1 = v1
                                                var v50 float32 = 0.0
                                                v49 += (v49 ^ 128)
                                                v51 := GenS3{F0: 1000, F1: 89}
                                                _ = v50
                                                _ = v51
                                            } else {
                                                v21 = 1.5
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v52 string = ("hello" + ("\n\t" + "foo bar"))
                                                var v53 bool = (true || (!(127 <= 1)))
                                                if true {
                                                    _ = len(v4)
                                                } else {
                                                    var v54 uint16 = 65535
                                                    var v55 uint32 = 1
                                                    _ = v54
                                                    _ = v55
                                                }
                                                if (42 < 0) {
                                                    v28 -= v28
                                                    if false {
                                                        v28 = (^(v28))
                                                        v3 = append(v3, ((((v52 + v52) + v52) + v52) + "foo bar"))
                                                    } else {
                                                        switch v52 {
                                                        case "\n\t":
                                                            v28 -= v28
                                                        case "a":
                                                            v28 = (v28 + 38)
                                                        default:
                                                            fmt.Println(v52, fi31, v53)
                                                        }
                                                        var v56 int32 = int32(v1)
                                                        _ = v56
                                                    }
                                                }
                                                _ = v52
                                                _ = v53
                                            }
                                        case -9223372036854775808:
                                            var v57 uint16 = 1
                                            _ = v57
                                        }
                                        _ = fi48
                                        _ = v49
                                    }
                                    var v58 int16 = v1
                                    if (!true) {
                                        if false {
                                            fmt.Println(v28, v21, fi31)
                                            v21 += 1.5
                                            var v59 int8 = -57
                                            _ = v59
                                        } else {
                                            var v60 int8 = -1
                                            _ = v60
                                        }
                                    }
                                    _ = v58
                                }
                                _ = fi31
                            }
                            v61 := GenS6{F0: -1, F1: false, F2: -138}
                            _ = v61
                        } else {
                            if ((false || true) && (4294967295 >= 127)) {
                                _ = len(v4)
                                v28 -= v28
                            }
                            var v62 float64 = (0.0 * (((-2.25 / 1.0) + (0.1 * ((-2.25 / (0.1 / 1.0)) / 1e10))) + 1e10))
                            _ = v62
                        }
                        fmt.Println(v21, v1)
                    }
                    v28 = (v28 - (v28 % ((255) | 1)))
                } else {
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                }
            default:
                var v63 int16 = v1
                _ = v63
            }
            var v64 int8 = 102
            v65 := GenS7{F0: (v1 & (v1 * (v1 + (v1 | -867))))}
            _ = v64
            _ = v65
        }
    } else {
        var v66 bool = false
        _ = v66
    }
    v67 := GenS5{F0: 32767, F1: 18446744073709551615, F2: ""}
    var v68 float64 = (1.5 * 1e10)
    v2 = append(v2, 2147483648)
    switch 127 {
    case 127:
        v69 := GenS8{F0: 255, F1: -70, F2: ""}
        _ = v69
    case 128:
        for fi70 := 0; fi70 < 2; fi70++ {
            _ = len(v2)
            v71 := GenS3{F0: (-(v1)), F1: 613}
            for fi72 := 0; fi72 < 19; fi72++ {
                var v73 string = "\n\t"
                _ = fi72
                _ = v73
            }
            _ = fi70
            _ = v71
        }
    case 92:
        switch 0 {
        case 1:
            fmt.Println(v1)
        default:
            var v74 float64 = v68
            _ = v74
        }
    }
    v28 = (^(v28))
    var v75 uint16 = 813
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v68:", v68)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v67
    _ = v68
    _ = v75
}
