package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 int; F1 uint16; F2 uint32 }
type GenS5 struct { F0 uint32; F1 int16; F2 int64 }
type GenS6 struct { F0 bool; F1 int32; F2 int8 }
type GenS7 struct { F0 uint16; F1 int8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int64; F1 int }
type GenS10 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int8; F1 uint; F2 int64 }
type GenS13 struct { F0 uint16; F1 int }

func main() {
    v1 := []int{-992, -421}
    var v2 int16 = -32768
    v3 := []uint32{2}
    v4 := make(map[string]uint32)
    fmt.Println(v2)
    v5 := GenS1{F0: 707, F1: 42}
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: 2, F1: 518}
            var v9 string = (("" + "hello") + "")
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: 2, F1: -1}
        if true {
            for fi11 := 0; fi11 < 5; fi11++ {
                v12 := GenS1{F0: 2, F1: 0}
                var v13 uint16 = 2
                v14 := GenS2{F0: 2, F1: v2}
                v13 = 1
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        }
        _ = fi6
        _ = v10
    }
    if false {
        for fi15 := 0; fi15 < 16; fi15++ {
            if false {
                v16 := GenS3{F0: false, F1: 255}
                _ = v16
            } else {
                if true {
                    v17 := v4[""]
                    switch ((("" + "hello") + "") + "") {
                    case "foo bar":
                        switch -13 {
                        case 127:
                            var v18 int = (fi15 / ((fi15) | 1))
                            _ = v18
                        case 9223372036854775807:
                            v19 := GenS2{F0: 4294967295, F1: (v2 / ((1) | 1))}
                            _ = v19
                        }
                    }
                    _ = v17
                }
                var v20 float32 = -2.25
                _ = v20
            }
            for fi21 := 0; fi21 < 6; fi21++ {
                fmt.Println(fi15)
                v2 -= (v2 / ((v2) | 1))
                var v22 int8 = 0
                switch "hello" {
                case "a":
                    v2 += (v2 | 1)
                case "hello":
                    switch (fi15 & (fi15 & (^(fi21)))) {
                    case 127:
                        v2 = (v2 * (v2 ^ v2))
                    case 9223372036854775807:
                        v23 := GenS4{F0: (fi15 / (((fi15 << ((255) & 63))) | 1)), F1: 36, F2: 2}
                        _ = v23
                    case 2147483647:
                        if true {
                            var v24 int8 = (-(v22))
                            v22 = (-(v22))
                            _ = v24
                        }
                    default:
                        v22 = -21
                    }
                case "":
                    if (42 > 253) {
                        _ = len(v1)
                        v25 := GenS5{F0: 792, F1: v2, F2: 0}
                        var v26 uint = 255
                        if (2 >= 1) {
                            switch (fi15 % ((fi15) | 1)) {
                            case -1:
                                _ = len(v4)
                            default:
                                if (true || (0 > -2147483648)) {
                                    var v27 uint = (v26 / (((^(v26))) | 1))
                                    if (((228 > 190) || (0 < 32767)) || true) {
                                        _ = len(v3)
                                        var v28 uint16 = 65535
                                        if len(v1) > 0 {
                                            v1[0] = 604
                                        }
                                        _ = v28
                                    }
                                    if (0 != 959) {
                                        fmt.Println(v27, v22, v26)
                                    } else {
                                        var v29 bool = true
                                        _ = v29
                                    }
                                    _ = v27
                                } else {
                                    switch ("\n\t" + ("a" + ("foo bar" + "foo bar"))) {
                                    case "foo bar":
                                        v30 := GenS1{F0: 0, F1: -132}
                                        _ = v30
                                    case "a":
                                        v22 = v22
                                    }
                                    var v31 uint8 = 91
                                    var v32 uint64 = 42
                                    _ = v31
                                    _ = v32
                                }
                            }
                            switch (fi15 * fi21) {
                            case 32768:
                                switch 128 {
                                case -551:
                                    var v33 float64 = (-2.25 * 1.5)
                                    _ = v33
                                case 32768:
                                    var v34 uint = (^(v26))
                                    _ = v34
                                default:
                                    var v35 uint16 = 530
                                    _ = v35
                                }
                            default:
                                var v36 bool = true
                                _ = v36
                            }
                        } else {
                            v37 := GenS6{F0: (!(2 < 532)), F1: 866, F2: 0}
                            _ = v37
                        }
                        var v38 float32 = (((0.1 * 0.1) + ((float32(0.1) / (1e10 + -2.25)) / (-2.25 * float32((3.14159 * ((3.14159 * 1e10) / 3.14159)))))) / float32(1.5))
                        if len(v1) > 0 {
                            v1[0] = fi15
                        }
                        _ = v25
                        _ = v26
                        _ = v38
                    } else {
                        var v39 string = "\n\t"
                        fmt.Println(v39, fi15, fi15)
                        _ = v39
                    }
                default:
                    var v40 int32 = 42
                    _ = v40
                }
                v22 -= int8(v2)
                v22 -= (-(v22))
                v22 -= (v22 << ((2) & 63))
                if true {
                    if false {
                        var v41 bool = true
                        _ = v41
                    } else {
                        fmt.Println(v2, v2, fi15)
                        v2 = v2
                    }
                    v42 := GenS6{F0: true, F1: -48, F2: v22}
                    if (109 >= 42) {
                        v22 = (-(v22))
                        switch fi15 {
                        case 32767:
                            var v43 int64 = 42
                            _ = v43
                        case 569:
                            var v44 uint64 = 708
                            _ = v44
                        default:
                            v2 = 42
                        }
                        var v45 uint64 = 42
                        switch "a" {
                        case "":
                            v46 := GenS7{F0: 362, F1: 100}
                            _ = v46
                        case "foo bar":
                            v47 := GenS5{F0: 272, F1: (v2 << ((1) & 63)), F2: -606}
                            _ = v47
                        case "hello":
                            var v48 float32 = (-2.25 / 1.5)
                            _ = v48
                        default:
                            v1 = append(v1, fi15)
                        }
                        var v49 int8 = v22
                        _ = v45
                        _ = v49
                    } else {
                        switch "a" {
                        case "":
                            var v50 bool = (true || false)
                            switch ("hello" + "foo bar") {
                            case "\n\t":
                                v51 := v4[("foo bar" + "")]
                                _ = v51
                            default:
                                fmt.Println(v2, v2)
                            }
                            var v52 uint32 = 42
                            _ = v50
                            _ = v52
                        case "\n\t":
                            v53 := GenS1{F0: 1, F1: 2147483647}
                            _ = v53
                        case "foo bar":
                            v22 += -128
                        }
                        switch "hello" {
                        case "":
                            fmt.Println(fi21, fi21)
                        case "a":
                            var v54 int = (fi15 - (fi15 ^ fi15))
                            _ = len(v4)
                            _ = v54
                        default:
                            fmt.Println(v2)
                            delete(v4, (("\n\t" + "hello") + "foo bar"))
                        }
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v22, fi15, fi15)
                        }
                        v2 = (v2 & (v2 << ((1) & 63)))
                    }
                    _ = v42
                } else {
                    if false {
                        v2 += 205
                        v22 -= (v22 | (^(v22)))
                    }
                }
                v22 -= (-(v22))
                _ = fi21
                _ = v22
            }
            _ = fi15
        }
        if (42 < -1) {
            for fi55 := 0; fi55 < 20; fi55++ {
                v56 := GenS4{F0: (fi55 * (-(fi55))), F1: 0, F2: 2}
                _ = fi55
                _ = v56
            }
        } else {
            var v57 string = "foo bar"
            if true {
                if false {
                    v58 := GenS8{F0: 255}
                    _ = v58
                } else {
                    v59 := GenS9{F0: -1, F1: 4294967295}
                    _ = v59
                }
                for fi60 := 0; fi60 < 4; fi60++ {
                    v2 -= v2
                    _ = fi60
                }
                fmt.Println(v57)
            }
            _ = v57
        }
    }
    if false {
        var v61 int = 32768
        _ = v61
    }
    v1 = append(v1, 4294967295)
    v2 = (v2 % ((374) | 1))
    if (1 <= 87) {
        delete(v4, (("a" + "hello") + "hello"))
        var v62 bool = true
        var v63 int64 = -9223372036854775808
        _ = v62
        _ = v63
    } else {
        fmt.Println(v2)
        if (!(0 == v2)) {
            fmt.Println(v2)
            fmt.Println(v2)
            if (0.1 > (1e10 / 3.14159)) {
                for fi64 := 0; fi64 < 6; fi64++ {
                    var v65 int32 = -327
                    var v66 uint64 = 507
                    if (fi64 != fi64) {
                        switch "foo bar" {
                        case "foo bar":
                            v65 = v65
                        case "a":
                            v65 -= v65
                        default:
                            fmt.Println(fi64, v66, v2)
                        }
                        switch 9223372036854775807 {
                        case 394:
                            for fi67 := 0; fi67 < 8; fi67++ {
                                v68 := GenS10{F0: (!false), F1: 2147483648, F2: 65535}
                                switch fi67 {
                                case -316:
                                    fmt.Println(v65)
                                }
                                _ = fi67
                                _ = v68
                            }
                        case 127:
                            var v69 uint64 = 947
                            _ = v69
                        case 2147483647:
                            var v70 uint = 18446744073709551615
                            _ = v70
                        default:
                            switch -1 {
                            case 567:
                                v2 = v2
                            case 0:
                                var v71 int64 = int64(fi64)
                                _ = v71
                            case 2:
                                var v72 int32 = (v65 & -378)
                                _ = v72
                            }
                        }
                        for fi73 := 0; fi73 < 13; fi73++ {
                            v65 = v65
                            switch "hello" {
                            case "":
                                if false {
                                    _ = len(v1)
                                    var v74 uint64 = (^(v66))
                                    switch fi64 {
                                    case 127:
                                        var v75 float64 = (-2.25 / 0.1)
                                        _ = v75
                                    case -208:
                                        switch -9223372036854775808 {
                                        case 764:
                                            switch 0 {
                                            case 128:
                                                v2 += v2
                                                var v76 bool = true
                                                v74 = v74
                                                _ = v76
                                            case -1:
                                                switch "\n\t" {
                                                case "":
                                                    v77 := v4[((("foo bar" + "foo bar") + (("\n\t" + (("" + ("hello" + ("" + ("\n\t" + ("a" + (("a" + "foo bar") + ("\n\t" + (("hello" + "hello") + "foo bar")))))))) + ("hello" + "foo bar"))) + "a")) + "\n\t")]
                                                    _ = v77
                                                default:
                                                    var v78 float32 = -2.25
                                                    _ = v78
                                                }
                                            default:
                                                var v79 uint32 = 671
                                                _ = v79
                                            }
                                        case 42:
                                            var v80 int8 = -128
                                            _ = v80
                                        case 4294967295:
                                            fmt.Println(v74, fi64, fi64)
                                        default:
                                            v81 := GenS6{F0: false, F1: v65, F2: 100}
                                            _ = v81
                                        }
                                    default:
                                        if (!false) {
                                            var v82 int = (fi73 >> ((0) & 63))
                                            _ = v82
                                        } else {
                                            v66 -= v66
                                            var v83 uint64 = v66
                                            _ = v83
                                        }
                                    }
                                    _ = v74
                                }
                            default:
                                v84 := GenS2{F0: 18446744073709551615, F1: v2}
                                _ = v84
                            }
                            _ = fi73
                        }
                    } else {
                        if ((1.5 != (((float32((float64(((1e10 + 3.14159) - (0.0 - ((float32((1.5 * 1e10)) - 3.14159) - (0.1 + 1e10))))) * 1.5)) - 3.14159) - -2.25) + 1e10)) || true) {
                            v85 := GenS11{F0: ((("hello" + "") + "foo bar") + "")}
                            if true {
                                if true {
                                    v65 = v65
                                    v2 -= (v2 * (v2 ^ (v2 & (v2 - v2))))
                                }
                            }
                            switch "hello" {
                            case "foo bar":
                                var v86 uint8 = 0
                                _ = v86
                            default:
                                v2 += 377
                            }
                            _ = v85
                        }
                    }
                    _ = fi64
                    _ = v65
                    _ = v66
                }
                v87 := GenS11{F0: "a"}
                for fi88 := 0; fi88 < 1; fi88++ {
                    switch "hello" {
                    case "\n\t":
                        v89 := GenS4{F0: 128, F1: 255, F2: 65535}
                        _ = v89
                    default:
                        var v90 int = fi88
                        _ = v90
                    }
                    v2 -= v2
                    _ = fi88
                }
                _ = v87
            } else {
                v91 := GenS12{F0: -1, F1: 968, F2: 1}
                fmt.Println(v2)
                _ = v91
            }
        } else {
            for fi92 := 0; fi92 < 5; fi92++ {
                v93 := GenS13{F0: 1, F1: fi92}
                fmt.Println(v2)
                _ = fi92
                _ = v93
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
