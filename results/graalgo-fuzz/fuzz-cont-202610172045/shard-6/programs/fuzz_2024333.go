package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int64; F2 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 int32; F2 uint64 }
type GenS5 struct { F0 int32; F1 string; F2 uint64 }
type GenS6 struct { F0 int }

func main() {
    v1 := []uint32{0, 2, 2147483648}
    v2 := make(map[string]uint32)
    var v3 bool = false
    var v4 uint = 42
    v5 := make(map[uint32]string)
    v6 := GenS1{F0: 1}
    if (false || v3) {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: 42}
            for fi9 := 0; fi9 < 11; fi9++ {
                v10 := GenS1{F0: 124}
                _ = fi9
                _ = v10
            }
            switch fi7 {
            case 127:
                if (v3 || true) {
                    v11 := GenS1{F0: 251}
                    _ = v11
                } else {
                    var v12 uint8 = 101
                    v4 = 65535
                    _ = v12
                }
            case 128:
                v3 = true
            case -846:
                v13 := GenS1{F0: 2}
                _ = v13
            }
            _ = fi7
            _ = v8
        }
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 uint8 = 0
            if (-299 != -1) {
                v16 := v2["foo bar"]
                if v3 {
                    fmt.Println(v4, v15)
                    if v3 {
                        var v17 int = (-(fi14))
                        _ = v17
                    } else {
                        fmt.Println(v15, v4, v15)
                    }
                    var v18 int16 = int16(fi14)
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 1; fi19++ {
                        if v3 {
                            v20 := GenS2{F0: 127, F1: -181, F2: 0}
                            v16 += uint32(v4)
                            _ = v20
                        } else {
                            fmt.Println(v16, v16, fi14)
                            var v21 int16 = 0
                            _ = v21
                        }
                        var v22 uint = 253
                        _ = fi19
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 4; fi23++ {
                        v4 = 0
                        var v24 float64 = 3.14159
                        _ = fi23
                        _ = v24
                    }
                }
                switch "hello" {
                case "hello":
                    var v25 float64 = 1e10
                    _ = v25
                }
                _ = v16
            } else {
                v26 := v5[421]
                var v27 uint8 = v15
                fmt.Println(v26, v4)
                _ = v26
                _ = v27
            }
            _ = fi14
            _ = v15
        }
    } else {
        for fi28 := 0; fi28 < 18; fi28++ {
            var v29 float64 = 3.14159
            var v30 float32 = 0.0
            _ = fi28
            _ = v29
            _ = v30
        }
        v5[2147483648] = ("\n\t" + "hello")
        v31 := GenS2{F0: 42, F1: 202, F2: 42}
        _ = v31
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        switch "\n\t" {
        case "a":
            fmt.Println(fi32, v4)
        case "":
            switch (fi32 | -173) {
            case -933:
                fmt.Println(v4, v4, v4)
            case 332:
                var v33 bool = v3
                _ = v33
            case -1:
                for fi34 := 0; fi34 < 14; fi34++ {
                    fmt.Println(v3)
                    _ = fi34
                }
            default:
                v3 = (42 <= -683)
            }
        }
        if true {
            switch (^(fi32)) {
            case 32767:
                switch ("a" + ("a" + "foo bar")) {
                case "\n\t":
                    if (!v3) {
                        if (!false) {
                            var v35 float64 = (0.0 * 0.0)
                            fmt.Println(v3)
                            v4 -= 2
                            _ = v35
                        }
                        var v36 uint = (v4 / ((v4) | 1))
                        _ = v36
                    } else {
                        fmt.Println(fi32, v4)
                        for fi37 := 0; fi37 < 12; fi37++ {
                            fmt.Println(fi32, v3, v3)
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                            switch (fi37 ^ (^(fi37))) {
                            case 814:
                                var v38 int8 = -86
                                _ = v38
                            }
                            switch (-(fi32)) {
                            case -576:
                                v3 = ((("a" + "") + ("foo bar" + "")) <= "\n\t")
                            case 396:
                                var v39 float32 = 1.5
                                _ = v39
                            }
                            v4 = 766
                            v2[("foo bar" + "foo bar")] = 65535
                            _ = fi37
                        }
                        if false {
                            for fi40 := 0; fi40 < 19; fi40++ {
                                var v41 uint64 = 4294967295
                                var v42 uint16 = 498
                                _ = fi40
                                _ = v41
                                _ = v42
                            }
                            var v43 float64 = (0.1 + -2.25)
                            delete(v2, "\n\t")
                            _ = v43
                        }
                    }
                case "":
                    var v44 int64 = -343
                    _ = v44
                case "a":
                    v3 = (488 >= 530)
                default:
                    var v45 int = 607
                    _ = v45
                }
            case 128:
                var v46 bool = v3
                _ = v46
            }
        } else {
            v3 = (true || v3)
            v4 = (v4 & (v4 << ((v4) & 63)))
        }
        v3 = (127 != 1)
        _ = fi32
    }
    v4 += (v4 >> ((592) & 63))
    v5[630] = ("hello" + "")
    for fi47 := 0; fi47 < 18; fi47++ {
        fmt.Println(v4)
        var v48 int64 = -1
        if (((^(v4)) >= v4) && true) {
            switch "\n\t" {
            case "\n\t":
                v4 -= 255
            case "":
                fmt.Println(v48, fi47)
            case "hello":
                fmt.Println(v48, v48, v48)
            }
        } else {
            if v3 {
                if (v3 || false) {
                    fmt.Println(v48, v3, fi47)
                    v48 = v48
                    fmt.Println(fi47, v3)
                } else {
                    var v49 float32 = float32(1e10)
                    var v50 int8 = -1
                    _ = v49
                    _ = v50
                }
                fmt.Println(v48)
            }
            var v51 string = "foo bar"
            var v52 uint = v4
            _ = v51
            _ = v52
        }
        _ = fi47
        _ = v48
    }
    for fi53 := 0; fi53 < 14; fi53++ {
        v4 += v4
        fmt.Println(v3)
        var v54 int64 = 86
        _ = fi53
        _ = v54
    }
    v4 += (v4 % ((v4) | 1))
    v55 := GenS3{F0: 127}
    for fi56 := 0; fi56 < 7; fi56++ {
        for fi57 := 0; fi57 < 18; fi57++ {
            v3 = false
            var v58 int8 = 42
            _ = len(v5)
            _ = fi57
            _ = v58
        }
        _ = fi56
    }
    switch 42 {
    case 0:
        var v59 int8 = 42
        _ = v59
    case -133:
        v60 := GenS2{F0: 127, F1: -514, F2: 127}
        _ = v60
    case 42:
        v4 -= (v4 >> ((2) & 63))
    }
    switch -700 {
    case -558:
        switch "hello" {
        case "":
            for fi61 := 0; fi61 < 9; fi61++ {
                for fi62 := 0; fi62 < 7; fi62++ {
                    if v3 {
                        var v63 uint16 = 255
                        var v64 int8 = -125
                        _ = v63
                        _ = v64
                    }
                    if true {
                        var v65 float64 = float64(1e10)
                        _ = v65
                    } else {
                        switch "\n\t" {
                        case "a":
                            v5[206] = (("\n\t" + ("foo bar" + ("a" + ("a" + ("hello" + "foo bar"))))) + "\n\t")
                        case "\n\t":
                            var v66 int64 = 242
                            _ = v66
                        case "foo bar":
                            var v67 string = ("a" + "hello")
                            _ = v67
                        }
                        if (!v3) {
                            v1 = append(v1, 4294967295)
                            switch ("" + "a") {
                            case "\n\t":
                                _ = len(v2)
                            case "foo bar":
                                var v68 int64 = -9223372036854775808
                                _ = v68
                            }
                            if (1 <= 9223372036854775807) {
                                if (!(42 >= -9223372036854775808)) {
                                    if ("hello" == "hello") {
                                        var v69 uint32 = 4294967295
                                        fmt.Println(v4, v4)
                                        _ = v69
                                    }
                                } else {
                                    if v3 {
                                        v3 = true
                                        v70 := GenS2{F0: 48, F1: -1, F2: 42}
                                        _ = v70
                                    } else {
                                        fmt.Println(fi61, v4, v4)
                                        var v71 int16 = 0
                                        var v72 uint8 = 42
                                        _ = v71
                                        _ = v72
                                    }
                                }
                            } else {
                                var v73 uint = (^(v4))
                                switch fi62 {
                                case -2147483648:
                                    switch "\n\t" {
                                    case "hello":
                                        v74 := GenS3{F0: (-(fi62))}
                                        _ = v74
                                    }
                                case 0:
                                    fmt.Println(v3)
                                default:
                                    if (336 <= (^(v73))) {
                                        var v75 string = "\n\t"
                                        _ = v75
                                    } else {
                                        v3 = (float64((-2.25 * (1e10 / (((0.0 - 1.5) * (0.1 / ((3.14159 / 1e10) - (3.14159 + 1e10)))) + 1e10)))) >= ((float64(-2.25) + 1e10) * -2.25))
                                        v76 := GenS4{F0: -73, F1: 1, F2: 18446744073709551615}
                                        _ = v76
                                    }
                                }
                                v4 += v73
                                _ = v73
                            }
                        } else {
                            fmt.Println(fi62, fi62, v4)
                        }
                        v4 = v4
                    }
                    switch fi62 {
                    case -283:
                        var v77 uint32 = 492
                        var v78 uint32 = v77
                        v3 = v3
                        _ = v77
                        _ = v78
                    case 32767:
                        v79 := GenS4{F0: -1, F1: 0, F2: 638}
                        var v80 uint64 = 722
                        switch ("a" + "a") {
                        case "foo bar":
                            if true {
                                switch fi61 {
                                case 128:
                                    _ = len(v5)
                                default:
                                    if v3 {
                                        var v81 bool = false
                                        _ = v81
                                    } else {
                                        v3 = ((!v3) && v3)
                                        var v82 bool = v3
                                        v83 := GenS3{F0: fi62}
                                        v84 := GenS5{F0: 2147483647, F1: "foo bar", F2: 18446744073709551615}
                                        v4 = (v4 - (v4 >> (((v4 >> (((^(v4))) & 63))) & 63)))
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                    }
                                }
                                if (true || (!v3)) {
                                    var v85 string = "foo bar"
                                    var v86 int = (fi62 + 32767)
                                    _ = len(v1)
                                    _ = v85
                                    _ = v86
                                }
                            } else {
                                v80 -= (v80 << ((v4) & 63))
                                v87 := GenS6{F0: (fi61 & (fi61 + (-(fi61))))}
                                _ = v87
                            }
                        }
                        _ = v79
                        _ = v80
                    case 2147483647:
                        var v88 float64 = -2.25
                        _ = v88
                    default:
                        var v89 int16 = 42
                        _ = v89
                    }
                    _ = fi62
                }
                v90 := GenS1{F0: 122}
                if len(v1) > 0 {
                    v1[0] = 774
                }
                _ = fi61
                _ = v90
            }
        default:
            fmt.Println(v3, v3)
        }
    case 4294967295:
        if true {
            var v91 int8 = 83
            v91 = v91
            _ = v91
        } else {
            v92 := GenS4{F0: 57, F1: 42, F2: 42}
            _ = v92
        }
    case -32:
        delete(v2, ("a" + "foo bar"))
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v3 = (!(0 > -733))
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
