package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int; F2 int64 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 int; F1 int64; F2 int16 }
type GenS7 struct { F0 int8; F1 uint; F2 uint16 }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 1
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    v5 := []uint32{}
    var v6 int64 = 1
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(fi7, v2, v6)
        var v8 int16 = -1
        v9 := GenS1{F0: 0}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        if true {
            switch fi10 {
            case 725:
                var v11 int16 = -1
                _ = v11
            case -204:
                if true {
                    var v12 string = (("foo bar" + "") + "a")
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 9; fi13++ {
                        v14 := GenS1{F0: 562}
                        v2 += (^(v2))
                        var v15 int32 = 2147483647
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS1{F0: 32767}
                    _ = v16
                }
            case 32768:
                var v17 uint16 = 1
                _ = v17
            default:
                for fi18 := 0; fi18 < 20; fi18++ {
                    var v19 bool = ((^(fi18)) >= 0)
                    _ = fi18
                    _ = v19
                }
            }
            for fi20 := 0; fi20 < 6; fi20++ {
                v5 = append(v5, 0)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v21 := GenS2{F0: (v6 ^ 9223372036854775807), F1: fi10, F2: 42}
                v22 := v1["\n\t"]
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        _ = fi10
    }
    var v23 uint8 = 127
    if ((!true) || ((v23 != 127) && false)) {
        if (-2.25 == 1e10) {
            for fi24 := 0; fi24 < 6; fi24++ {
                v6 = (-(v6))
                for fi25 := 0; fi25 < 14; fi25++ {
                    if false {
                        var v26 int8 = 42
                        if len(v5) > 0 {
                            v5[0] = 4294967295
                        }
                        _ = v26
                    } else {
                        var v27 string = (((("hello" + "") + "hello") + ("" + "foo bar")) + "")
                        _ = v27
                    }
                    var v28 int64 = -745
                    switch "hello" {
                    case "":
                        var v29 int32 = -894
                        _ = v29
                    case "a":
                        fmt.Println(v23, fi25, v23)
                    default:
                        _ = len(v5)
                    }
                    _ = fi25
                    _ = v28
                }
                _ = fi24
            }
            v30 := GenS3{F0: 1, F1: v6}
            _ = v30
        } else {
            fmt.Println(v6, v6)
            v2 -= 777
            for fi31 := 0; fi31 < 20; fi31++ {
                switch (((("\n\t" + "a") + "\n\t") + "foo bar") + "foo bar") {
                case "":
                    var v32 float64 = 1e10
                    _ = v32
                default:
                    v4[fi31] = fi31
                }
                _ = fi31
            }
        }
        v2 = v2
        fmt.Println(v23, v6)
    } else {
        v2 -= v2
        if (0 <= -32768) {
            var v33 int = 731
            _ = v33
        } else {
            if false {
                var v34 string = ("\n\t" + "\n\t")
                _ = v34
            } else {
                var v35 float64 = (0.0 - (3.14159 / 1.0))
                _ = v35
            }
            if false {
                for fi36 := 0; fi36 < 20; fi36++ {
                    for fi37 := 0; fi37 < 17; fi37++ {
                        var v38 int8 = 1
                        if len(v5) > 0 {
                            v5[0] = 42
                        }
                        _ = fi37
                        _ = v38
                    }
                    v6 += v6
                    _ = fi36
                }
                var v39 string = "\n\t"
                if false {
                    if len(v5) > 0 {
                        v5[0] = 2147483647
                    }
                    var v40 float64 = 0.1
                    switch "a" {
                    case "foo bar":
                        if ((v39 != v39) || ((v6 / ((-9223372036854775808) | 1)) != 0)) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if len(v5) > 0 {
                                v5[0] = 65535
                            }
                        } else {
                            v39 = "a"
                            var v41 float64 = (0.0 + v40)
                            for fi42 := 0; fi42 < 19; fi42++ {
                                v23 += (v23 >> ((42) & 63))
                                v6 = v6
                                _ = fi42
                            }
                            _ = v41
                        }
                    case "":
                        fmt.Println(v40, v40)
                    }
                    _ = v40
                }
                _ = v39
            } else {
                if (-2147483648 == 449) {
                    fmt.Println(v23, v6)
                }
                for fi43 := 0; fi43 < 8; fi43++ {
                    if false {
                        var v44 int8 = 0
                        for fi45 := 0; fi45 < 9; fi45++ {
                            var v46 bool = (false || ((v6 * 9223372036854775807) == (v6 & 9223372036854775807)))
                            var v47 float32 = (0.1 * ((float32((0.1 - ((1.5 - 1.5) + float64(1e10)))) + float32(1e10)) / 3.14159))
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                        v48 := GenS1{F0: 543}
                        _ = v44
                        _ = v48
                    }
                    if true {
                        var v49 uint32 = 2
                        _ = v49
                    } else {
                        v50 := GenS4{F0: 96}
                        for fi51 := 0; fi51 < 7; fi51++ {
                            v23 = v23
                            _ = fi51
                        }
                        v52 := GenS1{F0: 42}
                        _ = v50
                        _ = v52
                    }
                    _ = fi43
                }
                v2 -= 621
            }
        }
        _ = len(v4)
    }
    v6 = (v6 - v6)
    var v53 int = 42
    v2 -= v2
    v54 := GenS1{F0: 2}
    fmt.Println(v23, v53, v23)
    for fi55 := 0; fi55 < 4; fi55++ {
        if true {
            for fi56 := 0; fi56 < 17; fi56++ {
                fmt.Println(fi56, fi55, v2)
                _ = fi56
            }
            v2 = 703
        } else {
            v23 -= (v23 - 42)
        }
        switch (fi55 / (((fi55 % ((fi55) | 1))) | 1)) {
        case 0:
            v23 += v23
        case 42:
            if (!false) {
                var v57 uint16 = 994
                var v58 int64 = 9223372036854775807
                v58 -= (^(v58))
                _ = v57
                _ = v58
            } else {
                var v59 int64 = (^(v6))
                if (-36 > 189) {
                    v60 := GenS5{F0: (fi55 / (((^(v53))) | 1)), F1: v2}
                    for fi61 := 0; fi61 < 19; fi61++ {
                        v5 = append(v5, 4294967295)
                        if true {
                            v23 = (v23 % ((v23) | 1))
                            if (65535 != 193) {
                                if true {
                                    var v62 string = "foo bar"
                                    if true {
                                        v63 := GenS3{F0: 2147483647, F1: v6}
                                        var v64 uint64 = (v2 + v2)
                                        if (!(100 != -34)) {
                                            if (!((v59 + (v6 | v6)) >= (v6 - (v59 + int64(fi55))))) {
                                                var v65 uint16 = 2
                                                var v66 bool = false
                                                _ = v65
                                                _ = v66
                                            } else {
                                                fmt.Println(v59, v62, fi61)
                                                var v67 uint16 = 734
                                                _ = v67
                                            }
                                        }
                                        if (v59 != 77) {
                                            v59 = (v6 >> ((42) & 63))
                                            v68 := GenS6{F0: v53, F1: (^(v59)), F2: 1000}
                                            v62 = v62
                                            _ = v68
                                        } else {
                                            var v69 int32 = 591
                                            _ = v69
                                        }
                                        var v70 float32 = -2.25
                                        _ = v63
                                        _ = v64
                                        _ = v70
                                    } else {
                                        var v71 int8 = 46
                                        v72 := GenS5{F0: (v53 * 128), F1: 2}
                                        _ = v71
                                        _ = v72
                                    }
                                    fmt.Println(v62, v59, v6)
                                    var v73 int16 = 931
                                    _ = v62
                                    _ = v73
                                } else {
                                    switch 308 {
                                    case -650:
                                        if false {
                                            fmt.Println(fi61, v6)
                                            if false {
                                                fmt.Println(v53, fi61, fi61)
                                                v59 -= (^(v6))
                                                switch ((("foo bar" + ("\n\t" + "a")) + "hello") + "foo bar") {
                                                case "":
                                                    if (true && false) {
                                                        fmt.Println(v23)
                                                        fmt.Println(v53, v59)
                                                    }
                                                case "\n\t":
                                                    var v74 uint8 = v23
                                                    _ = v74
                                                case "hello":
                                                    var v75 uint8 = (^(v23))
                                                    _ = v75
                                                default:
                                                    var v76 int8 = 0
                                                    _ = v76
                                                }
                                            } else {
                                                v23 = (v23 & (v23 * (v23 - (v23 & (v23 / ((v23) | 1))))))
                                                v77 := GenS3{F0: 1, F1: -773}
                                                _ = v77
                                            }
                                            v78 := GenS2{F0: (v6 >> ((4294967295) & 63)), F1: fi61, F2: v6}
                                            v2 -= 758
                                            _ = v78
                                        } else {
                                            v79 := GenS3{F0: int32(v6), F1: -9223372036854775808}
                                            _ = v79
                                        }
                                    default:
                                        v80 := GenS7{F0: 127, F1: 529, F2: 65535}
                                        _ = v80
                                    }
                                    var v81 int64 = int64(v23)
                                    _ = v81
                                }
                                var v82 int64 = v59
                                _ = v82
                            } else {
                                if false {
                                    var v83 int16 = 874
                                    _ = v83
                                } else {
                                    var v84 uint16 = 283
                                    if false {
                                        switch fi61 {
                                        case 11:
                                            _ = len(v5)
                                        case 128:
                                            var v85 float64 = (((-2.25 - (1.5 / (0.1 - ((1.5 / 1.0) / (((((1.5 + ((float64((1.5 + 0.0)) - (1.5 + (0.1 * 1e10))) * 3.14159)) + 0.0) * ((((float64(((0.0 / (0.1 - -2.25)) * 3.14159)) + 3.14159) - 1e10) * 1.5) * (1.5 * -2.25))) - 0.1) - (0.0 - 3.14159)))))) + 1.5) - (-2.25 - (1e10 + float64((-2.25 * 0.1)))))
                                            _ = v85
                                        case 711:
                                            if len(v5) > 0 {
                                                v5[0] = 738
                                            }
                                        }
                                        v86 := GenS7{F0: -45, F1: 255, F2: 42}
                                        _ = v86
                                    } else {
                                        var v87 int64 = v6
                                        _ = v87
                                    }
                                    _ = v84
                                }
                                if (!true) {
                                    var v88 int16 = 1000
                                    _ = v88
                                } else {
                                    var v89 int16 = 42
                                    _ = v89
                                }
                            }
                        }
                        v6 = (v59 << ((4294967295) & 63))
                        v90 := GenS4{F0: -72}
                        _ = fi61
                        _ = v90
                    }
                    _ = v60
                }
                _ = v59
            }
        }
        _ = fi55
    }
    for fi91 := 0; fi91 < 12; fi91++ {
        fmt.Println(v23)
        var v92 uint32 = 302
        _ = fi91
        _ = v92
    }
    var v93 uint64 = (^(v2))
    if ((2147483648 >= 2147483647) && (1 <= 374)) {
        for fi94 := 0; fi94 < 13; fi94++ {
            v93 -= v93
            _ = fi94
        }
        var v95 uint32 = 2147483647
        var v96 int32 = -1
        _ = v95
        _ = v96
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
    _ = v93
}
