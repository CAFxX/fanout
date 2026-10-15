package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint; F1 uint16; F2 bool }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 bool; F1 uint; F2 int16 }
type GenS8 struct { F0 int8; F1 string; F2 uint }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint32) bool {
    var v4 uint64 = 1
    if ((((p0 <= p0) || true) || false) || true) {
        var v5 float64 = -2.25
        if (!true) {
            v5 -= (float64((3.14159 - (0.1 + -2.25))) * 1e10)
            var v6 bool = (((0 == 255) || true) || (9223372036854775807 >= -95))
            _ = v6
        }
        _ = v5
    }
    switch 798 {
    case 42:
        if (42 <= 2147483647) {
            v7 := GenS1{F0: p0}
            var v8 uint = 2
            _ = v7
            _ = v8
        }
    case 424:
        v4 += (v4 >> ((0) & 63))
    case 0:
        if false {
            for fi9 := 0; fi9 < 16; fi9++ {
                v10 := GenS2{F0: 42}
                v11 := []int{751, fi9}
                _ = fi9
                _ = v10
                _ = v11
            }
            var v12 uint = 4294967295
            v13 := []int{}
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 11; fi14++ {
                fmt.Println(v4, fi14, p0)
                _ = fi14
            }
        }
    default:
        for fi15 := 0; fi15 < 12; fi15++ {
            fmt.Println(fi15)
            _ = fi15
        }
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        v17 := GenS1{F0: 4294967295}
        v4 -= 4294967295
        if false {
            var v18 bool = true
            _ = v18
        } else {
            p0 -= (p0 / (((^(p0))) | 1))
        }
        _ = fi16
        _ = v17
    }
    v19 := GenS3{F0: 30, F1: (^(p0))}
    _ = p0
    _ = v19
    _ = v4
    return true
}
func genFn1() uint16 {
    switch ("\n\t" + (("" + (("\n\t" + "") + "foo bar")) + "a")) {
    case "\n\t":
        switch 127 {
        case 4294967295:
            var v20 uint8 = 185
            _ = v20
        case 166:
            for fi21 := 0; fi21 < 11; fi21++ {
                fmt.Println(fi21)
                _ = fi21
            }
        default:
            var v22 uint16 = 688
            _ = v22
        }
    }
    v23 := []int{1, 1, 109}
    _ = v23
    return 691
}

func main() {
    v1 := []int{2147483647}
    v2 := []string{"a", ("hello" + ("a" + "foo bar")), "\n\t"}
    v3 := make(map[uint32]int)
    var v24 uint64 = 1
    v3[65535] = -9223372036854775808
    fmt.Println(v24)
    var v25 int8 = 0
    v24 += v24
    var v26 int8 = (-(v25))
    if true {
        switch "a" {
        case "hello":
            for fi27 := 0; fi27 < 5; fi27++ {
                _ = len(v3)
                _ = fi27
            }
        case "a":
            v2 = append(v2, "\n\t")
        default:
            _ = len(v1)
        }
        if (!false) {
            var v28 uint64 = 823
            v29 := GenS3{F0: 1, F1: uint32(v28)}
            _ = v28
            _ = v29
        }
        for fi30 := 0; fi30 < 2; fi30++ {
            for fi31 := 0; fi31 < 4; fi31++ {
                switch (fi30 ^ -1) {
                case -1:
                    if false {
                        v32 := GenS1{F0: 42}
                        _ = v32
                    } else {
                        var v33 int16 = 1
                        v33 = v33
                        switch "a" {
                        case "a":
                            var v34 uint64 = (v24 ^ v24)
                            _ = v34
                        case "hello":
                            var v35 bool = false
                            _ = v35
                        case "foo bar":
                            v36 := GenS2{F0: 420}
                            _ = v36
                        }
                        _ = v33
                    }
                case 128:
                    v37 := GenS4{F0: true, F1: 0}
                    _ = v37
                default:
                    var v38 int64 = -9223372036854775808
                    _ = v38
                }
                _ = len(v3)
                _ = fi31
            }
            v25 = -128
            v39 := GenS5{F0: 65535, F1: 255, F2: ((false || false) || true)}
            _ = fi30
            _ = v39
        }
    } else {
        v40 := GenS4{F0: true, F1: 4294967295}
        for fi41 := 0; fi41 < 20; fi41++ {
            switch 0 {
            case 0:
                for fi42 := 0; fi42 < 19; fi42++ {
                    if (1 == -9223372036854775808) {
                        _ = len(v3)
                    }
                    var v43 uint8 = 0
                    var v44 int32 = -1
                    var v45 string = "a"
                    _ = fi42
                    _ = v43
                    _ = v44
                    _ = v45
                }
            case 4294967295:
                switch fi41 {
                case 0:
                    v46 := GenS6{F0: ("hello" + "foo bar"), F1: 874}
                    _ = v46
                case -794:
                    if (!false) {
                        if (false || (65535 < 255)) {
                            v25 = int8(v24)
                            v47 := GenS7{F0: true, F1: 2, F2: 151}
                            _ = v47
                        }
                        delete(v3, 38)
                        if len(v1) > 0 {
                            v1[0] = -990
                        }
                    } else {
                        fmt.Println(fi41, fi41)
                        for fi48 := 0; fi48 < 5; fi48++ {
                            fmt.Println(fi48)
                            if (255 <= 0) {
                                var v49 int32 = int32(v24)
                                if false {
                                    var v50 string = "hello"
                                    _ = v50
                                } else {
                                    switch "" {
                                    case "a":
                                        v51 := GenS1{F0: 4294967295}
                                        if false {
                                            v52 := GenS2{F0: 128}
                                            _ = len(v1)
                                            if len(v1) > 0 {
                                                v1[0] = int(v25)
                                            }
                                            _ = v52
                                        }
                                        _ = v51
                                    default:
                                        v24 += (v24 * v24)
                                    }
                                    var v53 uint32 = 836
                                    v54 := GenS8{F0: (v25 * (v26 - v25)), F1: ("a" + (("\n\t" + "") + "\n\t")), F2: 4294967295}
                                    _ = v53
                                    _ = v54
                                }
                                var v55 uint32 = 2147483648
                                _ = v49
                                _ = v55
                            }
                            _ = fi48
                        }
                        switch "" {
                        case "\n\t":
                            fmt.Println(v26, v25, fi41)
                        case "foo bar":
                            var v56 uint64 = v24
                            _ = v56
                        }
                    }
                case 9223372036854775807:
                    if len(v1) > 0 {
                        v1[0] = (fi41 * (fi41 + (^(fi41))))
                    }
                }
            case 128:
                v57 := GenS7{F0: (-1 <= 1), F1: 0, F2: -1}
                _ = v57
            }
            var v58 int8 = v25
            _ = fi41
            _ = v58
        }
        v59 := GenS5{F0: 2, F1: 252, F2: ((-(v25)) <= 42)}
        _ = v40
        _ = v59
    }
    switch "" {
    case "foo bar":
        if (!((true && false) || (!false))) {
            v26 = (v26 | (v26 % ((-128) | 1)))
            var v60 float64 = 0.1
            _ = v60
        } else {
            if (0 != 497) {
                v25 = 127
                v61 := GenS9{F0: 357, F1: 255}
                switch 83 {
                case 1:
                    for fi62 := 0; fi62 < 2; fi62++ {
                        fmt.Println(fi62, v24, v25)
                        for fi63 := 0; fi63 < 12; fi63++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            fmt.Println(v25, v26, v26)
                            fmt.Println(v24, fi62, fi62)
                            if (true && (genFn1() < 0)) {
                                var v64 int16 = -32768
                                if true {
                                    v26 += -1
                                    if true {
                                        fmt.Println(v26)
                                    }
                                    v65 := GenS10{F0: -1}
                                    _ = v65
                                }
                                var v66 int64 = 0
                                _ = len(v3)
                                if false {
                                    v66 -= (v66 >> ((1) & 63))
                                    if true {
                                        v67 := GenS6{F0: ("foo bar" + ("\n\t" + "foo bar")), F1: 970}
                                        switch "" {
                                        case "foo bar":
                                            var v68 uint8 = 48
                                            _ = v68
                                        case "hello":
                                            switch (fi62 | (fi62 ^ int(v25))) {
                                            case 1:
                                                switch ((("hello" + (("\n\t" + "foo bar") + "foo bar")) + "hello") + "hello") {
                                                case "\n\t":
                                                    var v69 float32 = 0.0
                                                    _ = v69
                                                case "":
                                                    if false {
                                                        v24 -= (v24 & v24)
                                                        v64 += (v64 ^ (^(v64)))
                                                    } else {
                                                        v1 = append(v1, fi63)
                                                    }
                                                default:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            case 32768:
                                                switch "foo bar" {
                                                case "":
                                                    if true {
                                                        var v70 float64 = 1.5
                                                        var v71 uint = 2
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                default:
                                                    if true {
                                                        if false {
                                                            v25 = v25
                                                        } else {
                                                            var v72 int32 = -144
                                                            v73 := GenS5{F0: 92, F1: 34, F2: true}
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                        var v74 int16 = (v64 >> ((4294967295) & 63))
                                                        var v75 int = 32768
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                }
                                            default:
                                                if true {
                                                    var v76 float32 = 3.14159
                                                    var v77 uint32 = 2147483648
                                                    var v78 uint = 1
                                                    _ = v76
                                                    _ = v77
                                                    _ = v78
                                                }
                                            }
                                        default:
                                            v64 -= v64
                                        }
                                        v79 := GenS9{F0: 1, F1: 0}
                                        _ = v67
                                        _ = v79
                                    }
                                }
                                _ = v64
                                _ = v66
                            } else {
                                if ((42 >= 43) && (!(!true))) {
                                    v25 -= (v25 | 127)
                                } else {
                                    var v80 uint8 = 1
                                    v26 = v25
                                    _ = v80
                                }
                                var v81 int8 = v25
                                _ = v81
                            }
                            _ = fi63
                        }
                        _ = fi62
                    }
                default:
                    for fi82 := 0; fi82 < 1; fi82++ {
                        switch fi82 {
                        case 1:
                            if false {
                                v26 += int8(v24)
                                v24 = (v24 << ((42) & 63))
                                fmt.Println(fi82, v26)
                            }
                        case 2147483647:
                            v26 = (v25 + (v25 | (v25 * v25)))
                        }
                        _ = fi82
                    }
                }
                _ = v61
            }
        }
    default:
        var v83 uint64 = (v24 >> ((255) & 63))
        _ = v83
    }
    if (-587 >= 589) {
        if false {
            if ((!(!(487 > 2147483647))) && false) {
                var v84 string = ("\n\t" + (("" + "hello") + (("" + ("hello" + "foo bar")) + "foo bar")))
                fmt.Println(v26, v84)
                _ = v84
            } else {
                if ((0 < 1) && ((!(v26 < (v25 | (-(v26))))) && (18446744073709551615 < v24))) {
                    var v85 string = ("hello" + "foo bar")
                    _ = v85
                }
            }
        }
        if (!((2147483647 >= 32768) || (((true && false) || (691 > 0)) || true))) {
            for fi86 := 0; fi86 < 11; fi86++ {
                v87 := v3[4294967295]
                fmt.Println(fi86)
                _ = fi86
                _ = v87
            }
        }
        v88 := GenS11{F0: 32767, F1: -722}
        _ = v88
    } else {
        v25 = (v26 >> ((0) & 63))
        _ = len(v3)
    }
    var v89 uint = 533
    var v90 int8 = (v26 | (v25 & v25))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v89
    _ = v90
}
