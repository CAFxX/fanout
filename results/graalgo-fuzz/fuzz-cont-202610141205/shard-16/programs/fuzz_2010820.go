package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32 }

func main() {
    v1 := []string{"\n\t"}
    var v2 bool = (true && true)
    var v3 uint32 = 351
    var v4 int32 = 2147483647
    for fi5 := 0; fi5 < 1; fi5++ {
        v1 = append(v1, "")
        _ = fi5
    }
    switch "" {
    case "hello":
        fmt.Println(v3)
    case "a":
        if true {
            var v6 uint16 = 255
            _ = v6
        } else {
            var v7 string = "\n\t"
            _ = v7
        }
    case "foo bar":
        if (v2 || (false || false)) {
            if false {
                fmt.Println(v4)
            }
            v8 := GenS1{F0: 18446744073709551615, F1: 42}
            for fi9 := 0; fi9 < 10; fi9++ {
                v3 = (v3 / ((v3) | 1))
                _ = fi9
            }
            _ = v8
        } else {
            v3 -= v3
            for fi10 := 0; fi10 < 14; fi10++ {
                for fi11 := 0; fi11 < 3; fi11++ {
                    v4 += 1
                    _ = fi11
                }
                var v12 int64 = 9223372036854775807
                var v13 float64 = (-2.25 * 1e10)
                _ = fi10
                _ = v12
                _ = v13
            }
            _ = len(v1)
        }
    default:
        v14 := GenS2{F0: (v2 || (!v2)), F1: 42}
        _ = v14
    }
    if v2 {
        var v15 uint32 = 2147483648
        var v16 float32 = ((1e10 + (3.14159 * (0.1 - float32(1.5)))) / -2.25)
        if (((0 < 540) || (-2.25 < 1.5)) && (v2 || true)) {
            v17 := GenS3{F0: 100}
            _ = v17
        }
        _ = v15
        _ = v16
    } else {
        if (493 > 0) {
            if (-1 < 9223372036854775807) {
                switch ("\n\t" + (("hello" + ("foo bar" + "hello")) + "a")) {
                case "\n\t":
                    var v18 uint16 = 42
                    _ = v18
                case "a":
                    switch 9223372036854775807 {
                    case 32768:
                        v19 := GenS4{F0: 127, F1: (v3 & v3), F2: 2147483647}
                        _ = v19
                    }
                case "":
                    var v20 uint = 18446744073709551615
                    _ = v20
                }
                var v21 int8 = 112
                _ = v21
            }
            var v22 int8 = -2
            for fi23 := 0; fi23 < 19; fi23++ {
                switch "a" {
                case "":
                    for fi24 := 0; fi24 < 16; fi24++ {
                        v25 := GenS1{F0: 416, F1: 375}
                        v3 = (v3 - v3)
                        if false {
                            var v26 uint64 = 0
                            _ = v26
                        } else {
                            switch "a" {
                            case "\n\t":
                                v27 := GenS1{F0: 4294967295, F1: 999}
                                _ = v27
                            default:
                                v28 := GenS5{F0: 124, F1: -9223372036854775808}
                                _ = v28
                            }
                        }
                        _ = fi24
                        _ = v25
                    }
                default:
                    v22 = v22
                }
                v3 -= 2147483648
                v29 := GenS6{F0: 2}
                _ = fi23
                _ = v29
            }
            _ = v22
        } else {
            for fi30 := 0; fi30 < 4; fi30++ {
                for fi31 := 0; fi31 < 6; fi31++ {
                    v32 := GenS4{F0: -128, F1: (v3 * (^(v3))), F2: v4}
                    v3 = v3
                    _ = fi31
                    _ = v32
                }
                fmt.Println(v4, fi30, fi30)
                _ = fi30
            }
            v33 := GenS3{F0: 100}
            fmt.Println(v4)
            _ = v33
        }
        for fi34 := 0; fi34 < 9; fi34++ {
            if ((1 == 0) && (597 < 42)) {
                var v35 int64 = 1
                v1 = append(v1, "\n\t")
                if ((fi34 >> ((65535) & 63)) == fi34) {
                    fmt.Println(v3, v4)
                    for fi36 := 0; fi36 < 12; fi36++ {
                        fmt.Println(v4, v3)
                        if (v3 <= 0) {
                            if v2 {
                                if v2 {
                                    var v37 float32 = 1.5
                                    v38 := GenS4{F0: -128, F1: 2147483648, F2: (v4 / ((v4) | 1))}
                                    var v39 bool = true
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                                v3 -= (v3 << ((255) & 63))
                            } else {
                                if (!((128 < 1) && v2)) {
                                    var v40 int8 = 0
                                    v4 += v4
                                    v2 = true
                                    _ = v40
                                } else {
                                    v4 = (-(v4))
                                    v41 := GenS6{F0: 841}
                                    if (407 >= 920) {
                                        switch (^(fi36)) {
                                        case 127:
                                            if ((!false) && (v2 && false)) {
                                                if len(v1) > 0 {
                                                    v1[0] = "hello"
                                                }
                                            } else {
                                                var v42 uint32 = v3
                                                _ = v42
                                            }
                                        case 4294967295:
                                            var v43 bool = v2
                                            _ = v43
                                        case -9223372036854775808:
                                            if len(v1) > 0 {
                                                v1[0] = ""
                                            }
                                        default:
                                            var v44 float32 = 3.14159
                                            _ = v44
                                        }
                                    } else {
                                        v45 := GenS4{F0: 0, F1: (v3 - (v3 << ((42) & 63))), F2: v4}
                                        var v46 int8 = 100
                                        _ = v45
                                        _ = v46
                                    }
                                    _ = v41
                                }
                                v4 += 1
                                fmt.Println(v3, v4, v35)
                                var v47 int8 = 100
                                _ = v47
                            }
                            if (!false) {
                                fmt.Println(v4, v3)
                            } else {
                                _ = len(v1)
                            }
                            v3 -= 4294967295
                        }
                        var v48 string = "hello"
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi36
                        _ = v48
                    }
                } else {
                    v35 += v35
                    var v49 uint64 = 628
                    _ = v49
                }
                _ = v35
            }
            switch -784 {
            case 2147483647:
                var v50 float64 = 3.14159
                _ = v50
            default:
                v3 += (v3 % (((v3 ^ (^(v3)))) | 1))
            }
            _ = fi34
        }
        for fi51 := 0; fi51 < 18; fi51++ {
            var v52 uint16 = 578
            _ = fi51
            _ = v52
        }
    }
    switch 127 {
    case 42:
        if (690 != -726) {
            v4 -= v4
            if v2 {
                switch -838 {
                case -1:
                    var v53 float32 = 0.0
                    _ = v53
                }
                if true {
                    switch -169 {
                    case 4294967295:
                        var v54 string = "a"
                        _ = v54
                    }
                    for fi55 := 0; fi55 < 13; fi55++ {
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        if (2 >= 2) {
                            for fi56 := 0; fi56 < 17; fi56++ {
                                v4 = 2147483647
                                switch (fi56 ^ (fi55 % (((fi55 << ((18446744073709551615) & 63))) | 1))) {
                                case 1:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                default:
                                    v3 -= (v3 + v3)
                                }
                                v57 := GenS4{F0: 100, F1: 535, F2: (^(v4))}
                                _ = fi56
                                _ = v57
                            }
                            if v2 {
                                v3 -= 65535
                                v4 -= (v4 - v4)
                            } else {
                                v1 = append(v1, ("foo bar" + ""))
                                v2 = ((-76 < -128) || v2)
                            }
                            fmt.Println(v2, v3, v3)
                        }
                        _ = fi55
                    }
                    var v58 int8 = 36
                    _ = v58
                } else {
                    if v2 {
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                    } else {
                        for fi59 := 0; fi59 < 11; fi59++ {
                            if false {
                                v1 = append(v1, "a")
                                for fi60 := 0; fi60 < 13; fi60++ {
                                    v2 = v2
                                    if len(v1) > 0 {
                                        v1[0] = ("" + ("\n\t" + ""))
                                    }
                                    v61 := GenS4{F0: 116, F1: 908, F2: (-(v4))}
                                    _ = fi60
                                    _ = v61
                                }
                                v62 := GenS7{F0: 127}
                                _ = v62
                            }
                            v2 = (!((fi59 | fi59) == fi59))
                            _ = fi59
                        }
                        v63 := GenS5{F0: 523, F1: 0}
                        var v64 int64 = -9223372036854775808
                        _ = v63
                        _ = v64
                    }
                    v4 = v4
                }
                var v65 uint32 = (v3 ^ (v3 << ((65535) & 63)))
                _ = v65
            }
        } else {
            for fi66 := 0; fi66 < 15; fi66++ {
                switch "\n\t" {
                case "":
                    v67 := GenS8{F0: v4}
                    _ = v67
                case "a":
                    var v68 int16 = 1000
                    _ = v68
                case "hello":
                    var v69 float64 = (float64(float32(-2.25)) + (0.1 + (3.14159 + float64(float32(0.0)))))
                    _ = v69
                }
                v70 := GenS8{F0: (v4 & (v4 + -1))}
                for fi71 := 0; fi71 < 10; fi71++ {
                    if v2 {
                        fmt.Println(v2, v3)
                    } else {
                        var v72 int16 = 1
                        fmt.Println(v2)
                        if (!((!false) && (!(1 == 635)))) {
                            var v73 int64 = -9223372036854775808
                            _ = v73
                        }
                        _ = v72
                    }
                    _ = fi71
                }
                _ = fi66
                _ = v70
            }
        }
    default:
        v4 += (v4 << ((18446744073709551615) & 63))
    }
    v3 += v3
    v2 = v2
    var v74 bool = v2
    var v75 uint16 = 2
    fmt.Println(v3)
    var v76 uint8 = 0
    if v74 {
        v77 := GenS7{F0: uint64(v3)}
        switch "" {
        case "foo bar":
            var v78 uint16 = v75
            _ = v78
        default:
            var v79 int = -1
            _ = v79
        }
        _ = v77
    } else {
        var v80 float32 = -2.25
        if (9223372036854775807 > 605) {
            var v81 float64 = 1e10
            switch int(v75) {
            case 32767:
                var v82 bool = v2
                _ = v82
            }
            var v83 uint32 = v3
            _ = v81
            _ = v83
        }
        _ = v80
    }
    if v74 {
        switch 1 {
        case 128:
            for fi84 := 0; fi84 < 10; fi84++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi84
            }
        case 1:
            var v85 uint = 398
            _ = v85
        case 2147483647:
            fmt.Println(v2, v76)
        default:
            if ((^(v3)) <= v3) {
                v86 := GenS4{F0: 42, F1: (v3 << ((123) & 63)), F2: int32(v75)}
                _ = v86
            } else {
                for fi87 := 0; fi87 < 7; fi87++ {
                    var v88 int8 = 13
                    v89 := GenS1{F0: 18446744073709551615, F1: uint64(v76)}
                    _ = fi87
                    _ = v88
                    _ = v89
                }
                if v74 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if (v2 && (v75 >= (v75 ^ v75))) {
                        var v90 int64 = 0
                        _ = v90
                    } else {
                        for fi91 := 0; fi91 < 20; fi91++ {
                            var v92 uint32 = (^(v3))
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi91
                            _ = v92
                        }
                        v74 = false
                    }
                    var v93 float64 = 1.5
                    _ = v93
                }
                v2 = v2
            }
        }
    }
    v3 = 774
    v3 -= v3
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v74
    _ = v75
    _ = v76
}
