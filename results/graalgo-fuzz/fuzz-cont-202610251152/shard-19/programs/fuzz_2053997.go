package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint; F2 uint32 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 bool; F1 int64; F2 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 string; F1 int64 }
type GenS9 struct { F0 int8; F1 int64; F2 uint64 }
type GenS10 struct { F0 int; F1 int16 }
type GenS11 struct { F0 uint8; F1 int16 }
type GenS12 struct { F0 bool; F1 int32; F2 bool }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int32; F1 uint16; F2 int32 }
func genFn0(p0 uint, p1 bool) uint64 {
    var v4 int32 = 0
    v5 := GenS1{F0: 925, F1: "hello", F2: p0}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}
func genFn1() int64 {
    var v6 uint16 = 220
    if ((4294967295 <= 2147483647) || (0 != -2147483648)) {
        var v7 uint16 = 255
        _ = v7
    } else {
        v6 -= (^(v6))
        for fi8 := 0; fi8 < 14; fi8++ {
            v6 = (v6 + v6)
            var v9 uint64 = 18446744073709551615
            v6 = 42
            _ = fi8
            _ = v9
        }
        if (!false) {
            v6 = (v6 / ((v6) | 1))
        }
    }
    var v10 bool = (-1 == 0)
    if (-32768 <= 1000) {
        v6 -= 511
        v6 = v6
        switch 127 {
        case 32768:
            v11 := []int{1, 221, 335}
            _ = v11
        case -2147483648:
            v6 = v6
        default:
            v12 := GenS1{F0: 1, F1: (("a" + "") + "hello"), F2: 4294967295}
            _ = v12
        }
    }
    var v13 string = ""
    _ = v10
    _ = v13
    _ = v6
    return -9223372036854775808
}

func main() {
    var v1 float64 = 1e10
    v2 := []uint32{2147483648}
    v3 := make(map[uint32]bool)
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 int32 = 0
        _ = fi14
        _ = v15
    }
    var v16 int8 = 42
    fmt.Println(v1)
    v16 = v16
    _ = len(v2)
    var v17 uint32 = 227
    switch ("hello" + "\n\t") {
    case "\n\t":
        fmt.Println(v16)
    case "":
        switch "" {
        case "hello":
            fmt.Println(v1, v17, v16)
        }
    case "hello":
        switch 317 {
        case 0:
            switch 9223372036854775807 {
            case 210:
                var v18 string = ("hello" + "")
                _ = v18
            case -1:
                switch (("\n\t" + (("foo bar" + "a") + ("a" + ("\n\t" + "\n\t")))) + "") {
                case "":
                    if false {
                        if (0 >= genFn0(876, ((true || ((v1 - v1) >= (1e10 / (v1 * v1)))) || (("a" + "\n\t") < ("foo bar" + "a"))))) {
                            for fi19 := 0; fi19 < 16; fi19++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi19
                            }
                            var v20 float32 = (0.0 / 1e10)
                            _ = v20
                        } else {
                            var v21 int64 = -836
                            v1 -= (1e10 / 0.1)
                            _ = v21
                        }
                        if true {
                            if false {
                                fmt.Println(v16, v1)
                            } else {
                                v22 := GenS2{F0: v16}
                                var v23 int8 = v16
                                _ = v22
                                _ = v23
                            }
                        } else {
                            for fi24 := 0; fi24 < 2; fi24++ {
                                _ = len(v2)
                                for fi25 := 0; fi25 < 12; fi25++ {
                                    fmt.Println(fi25)
                                    var v26 int16 = 474
                                    _ = fi25
                                    _ = v26
                                }
                                v27 := GenS3{F0: (42 > 205), F1: 2, F2: (v17 / ((227) | 1))}
                                _ = fi24
                                _ = v27
                            }
                        }
                    }
                case "a":
                    v28 := GenS4{F0: "", F1: 2}
                    _ = v28
                case "\n\t":
                    switch "a" {
                    case "\n\t":
                        v17 = (v17 ^ (v17 % (((v17 << ((840) & 63))) | 1)))
                    case "foo bar":
                        if (1 != 42) {
                            fmt.Println(v17, v17, v1)
                        } else {
                            for fi29 := 0; fi29 < 17; fi29++ {
                                var v30 int8 = (v16 ^ (v16 * 91))
                                if ((1.5 * (float64((-2.25 - 3.14159)) + float64(float32(v1)))) == v1) {
                                    for fi31 := 0; fi31 < 19; fi31++ {
                                        v32 := GenS2{F0: -128}
                                        _ = fi31
                                        _ = v32
                                    }
                                    for fi33 := 0; fi33 < 10; fi33++ {
                                        v1 += (0.0 / v1)
                                        var v34 uint8 = 248
                                        v3[2147483648] = true
                                        _ = fi33
                                        _ = v34
                                    }
                                }
                                _ = fi29
                                _ = v30
                            }
                            if (1 >= 0) {
                                var v35 uint64 = 4294967295
                                _ = v35
                            } else {
                                _ = len(v2)
                            }
                        }
                    default:
                        _ = len(v2)
                    }
                }
            default:
                v17 += (v17 ^ v17)
            }
        case -9223372036854775808:
            if false {
                if len(v2) > 0 {
                    v2[0] = 452
                }
                for fi36 := 0; fi36 < 6; fi36++ {
                    var v37 string = "\n\t"
                    var v38 uint32 = 54
                    v39 := GenS5{F0: true, F1: 433, F2: genFn0(0, true)}
                    _ = fi36
                    _ = v37
                    _ = v38
                    _ = v39
                }
                _ = len(v2)
            }
        case 9223372036854775807:
            fmt.Println(v16, v1)
        }
    default:
        v1 += float64(0.1)
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v40 int8 = (^(v16))
        _ = v40
    } else {
        for fi41 := 0; fi41 < 3; fi41++ {
            if (!(!false)) {
                v17 -= v17
                for fi42 := 0; fi42 < 16; fi42++ {
                    v43 := GenS6{F0: fi41}
                    v16 = 100
                    v1 += v1
                    _ = fi42
                    _ = v43
                }
                var v44 uint16 = 0
                _ = v44
            }
            _ = fi41
        }
        v16 -= (^(v16))
        var v45 uint = 484
        _ = v45
    }
    v46 := GenS3{F0: (!(!(127 == 0))), F1: 698, F2: (v17 - (^(v17)))}
    if false {
        v17 = 947
        v47 := GenS3{F0: true, F1: 65535, F2: (^(v17))}
        if true {
            v16 += -103
            v48 := GenS7{F0: 837, F1: 4294967295}
            _ = v48
        }
        _ = v47
    } else {
        fmt.Println(v16)
        if false {
            fmt.Println(v1)
            if (v1 < (v1 * (v1 * -2.25))) {
                v49 := GenS6{F0: -396}
                fmt.Println(v16)
                _ = v49
            } else {
                if true {
                    if (v17 > v17) {
                        var v50 uint64 = genFn0(718, false)
                        _ = v50
                    } else {
                        v16 = v16
                    }
                    v51 := GenS8{F0: "", F1: -9223372036854775808}
                    _ = v51
                }
                if false {
                    v52 := GenS9{F0: (v16 - 83), F1: -470, F2: 2}
                    v1 = (3.14159 + v1)
                    _ = v52
                }
                for fi53 := 0; fi53 < 19; fi53++ {
                    var v54 int16 = -32768
                    if ((v16 << ((2) & 63)) >= v16) {
                        v16 -= v16
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v55 int = fi53
                        _ = v55
                    }
                    for fi56 := 0; fi56 < 3; fi56++ {
                        if ((false && (-1 == 1)) || true) {
                            v57 := GenS10{F0: 1, F1: 802}
                            v17 += v17
                            fmt.Println(v16, v17)
                            _ = v57
                        }
                        v54 -= (^(v54))
                        v1 = 0.0
                        _ = fi56
                    }
                    _ = fi53
                    _ = v54
                }
            }
            if true {
                var v58 int = 1
                _ = v58
            } else {
                if ((2 <= 127) || ((!((!(0 > v16)) || (!false))) || false)) {
                    if (((v1 + v1) * v1) > v1) {
                        fmt.Println(v16)
                        for fi59 := 0; fi59 < 6; fi59++ {
                            if false {
                                if true {
                                    var v60 bool = (1 == 661)
                                    switch fi59 {
                                    case 128:
                                        switch (-(fi59)) {
                                        case 42:
                                            v61 := GenS10{F0: (-(fi59)), F1: -1}
                                            _ = v61
                                        case 32768:
                                            for fi62 := 0; fi62 < 19; fi62++ {
                                                if (-417 == -825) {
                                                    switch (fi59 % ((-748) | 1)) {
                                                    case -2147483648:
                                                        var v63 int64 = -9223372036854775808
                                                        _ = v63
                                                    case 414:
                                                        v1 -= v1
                                                    case -937:
                                                        var v64 uint8 = 80
                                                        _ = v64
                                                    default:
                                                        var v65 float64 = 0.0
                                                        _ = v65
                                                    }
                                                    if ((v17 >> ((255) & 63)) <= 2147483648) {
                                                        var v66 bool = ((v17 & (^(v17))) <= 474)
                                                        v16 = (v16 | (v16 + (v16 % (((-(v16))) | 1))))
                                                        v67 := GenS11{F0: 128, F1: -478}
                                                        _ = v66
                                                        _ = v67
                                                    } else {
                                                        v16 = 127
                                                        v68 := GenS1{F0: 0, F1: ("" + "a"), F2: 1}
                                                        v16 = v16
                                                        _ = v68
                                                    }
                                                } else {
                                                    fmt.Println(v16)
                                                    v17 -= 710
                                                }
                                                _ = fi62
                                            }
                                        }
                                    case 9223372036854775807:
                                        v69 := GenS12{F0: v60, F1: 358, F2: false}
                                        _ = v69
                                    case 415:
                                        var v70 uint16 = 255
                                        _ = v70
                                    }
                                    _ = v60
                                } else {
                                    v1 = -2.25
                                }
                            }
                            var v71 uint8 = 2
                            _ = fi59
                            _ = v71
                        }
                        for fi72 := 0; fi72 < 20; fi72++ {
                            for fi73 := 0; fi73 < 13; fi73++ {
                                var v74 uint = 447
                                _ = fi73
                                _ = v74
                            }
                            _ = fi72
                        }
                    }
                }
            }
        } else {
            v1 = v1
            if (v17 == (v17 ^ (v17 << ((18446744073709551615) & 63)))) {
                switch "a" {
                case "\n\t":
                    var v75 bool = (!true)
                    _ = v75
                default:
                    v76 := GenS13{F0: 2147483647}
                    _ = v76
                }
                fmt.Println(v17, v17, v1)
            }
            var v77 int32 = -1
            _ = v77
        }
    }
    v1 += v1
    if true {
        if false {
            var v78 bool = true
            v79 := GenS8{F0: ("a" + "\n\t"), F1: 761}
            v80 := GenS5{F0: (v78 || (-473 >= 2147483647)), F1: -9223372036854775808, F2: 368}
            _ = v78
            _ = v79
            _ = v80
        }
        v81 := GenS14{F0: 0, F1: 42, F2: 123}
        switch 42 {
        case -670:
            for fi82 := 0; fi82 < 16; fi82++ {
                for fi83 := 0; fi83 < 9; fi83++ {
                    _ = len(v2)
                    _ = fi83
                }
                v17 = (v17 << ((469) & 63))
                var v84 uint32 = (v17 / ((v17) | 1))
                _ = fi82
                _ = v84
            }
        case 4294967295:
            var v85 float32 = 3.14159
            _ = v85
        case -2147483648:
            var v86 int8 = (v16 * v16)
            _ = v86
        default:
            var v87 uint64 = 131
            _ = v87
        }
        _ = v81
    }
    v16 = (v16 / ((v16) | 1))
    var v88 int64 = 49
    v89 := GenS12{F0: (32767 >= 9223372036854775807), F1: -2147483648, F2: (true || (true || true))}
    var v90 uint16 = 42
    _ = len(v3)
    if (true && ("" != "")) {
        v91 := GenS13{F0: 0}
        _ = v91
    } else {
        var v92 float32 = 1.5
        v88 += v88
        _ = v92
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v88:", v88)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v46
    _ = v88
    _ = v89
    _ = v90
}
