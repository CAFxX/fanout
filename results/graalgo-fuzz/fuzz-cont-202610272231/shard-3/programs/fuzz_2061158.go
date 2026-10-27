package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 string; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 string; F2 uint }
type GenS6 struct { F0 uint32; F1 uint32; F2 int }
type GenS7 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS8 struct { F0 int32; F1 int8 }
type GenS9 struct { F0 uint64; F1 uint16; F2 bool }
type GenS10 struct { F0 uint32 }

func main() {
    v1 := []uint32{65535, 4294967295}
    v2 := []string{"\n\t", "\n\t"}
    var v3 float32 = 3.14159
    if false {
        v4 := GenS1{F0: true}
        if false {
            switch "" {
            case "hello":
                var v5 int8 = 0
                _ = v5
            default:
                v6 := GenS1{F0: (true || (!(0 >= 1)))}
                _ = v6
            }
        }
        _ = v4
    }
    if (0 <= -550) {
        if false {
            var v7 uint8 = 0
            if len(v1) > 0 {
                v1[0] = 738
            }
            _ = v7
        } else {
            for fi8 := 0; fi8 < 7; fi8++ {
                switch fi8 {
                case 0:
                    for fi9 := 0; fi9 < 4; fi9++ {
                        v10 := GenS2{F0: 17}
                        _ = fi9
                        _ = v10
                    }
                }
                var v11 int = 9223372036854775807
                v12 := GenS3{F0: 1, F1: "foo bar", F2: 0}
                _ = fi8
                _ = v11
                _ = v12
            }
            if true {
                v13 := GenS2{F0: 42}
                for fi14 := 0; fi14 < 6; fi14++ {
                    var v15 uint = 1
                    var v16 int16 = -578
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                _ = v13
            } else {
                v17 := GenS4{F0: -699}
                for fi18 := 0; fi18 < 12; fi18++ {
                    if (18446744073709551615 != 0) {
                        switch (fi18 >> ((100) & 63)) {
                        case 32768:
                            if false {
                                for fi19 := 0; fi19 < 2; fi19++ {
                                    if len(v2) > 0 {
                                        v2[0] = ("hello" + ("\n\t" + "foo bar"))
                                    }
                                    v20 := GenS2{F0: 0}
                                    if false {
                                        fmt.Println(fi18)
                                        v3 -= -2.25
                                        if (0 <= 128) {
                                            var v21 uint16 = 850
                                            _ = len(v1)
                                            fmt.Println(fi19, fi19)
                                            if len(v2) > 0 {
                                                v2[0] = (("foo bar" + "foo bar") + "\n\t")
                                            }
                                            if (1 != uint8(v21)) {
                                                var v22 int64 = -1
                                                _ = v22
                                            } else {
                                                var v23 uint32 = 330
                                                var v24 uint = 65535
                                                var v25 int64 = 42
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                            }
                                            switch (fi19 << ((532) & 63)) {
                                            case -2147483648:
                                                var v26 uint = 51
                                                _ = v26
                                            default:
                                                v21 += (v21 * (v21 & 746))
                                            }
                                            var v27 uint64 = 18446744073709551615
                                            v21 += 65535
                                            _ = v21
                                            _ = v27
                                        } else {
                                            fmt.Println(fi19, fi19)
                                        }
                                    } else {
                                        var v28 uint16 = 32767
                                        _ = v28
                                    }
                                    _ = fi19
                                    _ = v20
                                }
                                var v29 int = fi18
                                fmt.Println(v3, fi18)
                                _ = v29
                            } else {
                                _ = len(v1)
                            }
                        case -9223372036854775808:
                            var v30 uint64 = 251
                            _ = v30
                        case -2147483648:
                            for fi31 := 0; fi31 < 13; fi31++ {
                                v3 -= (v3 + (float32(3.14159) + (3.14159 - 3.14159)))
                                var v32 float32 = float32(((1.5 - 1.5) * (float64(float32(0.1)) - -2.25)))
                                _ = fi31
                                _ = v32
                            }
                        }
                        var v33 int8 = -1
                        _ = v33
                    }
                    if true {
                        v1 = append(v1, 65535)
                        fmt.Println(v3)
                        v34 := GenS4{F0: 639}
                        _ = v34
                    } else {
                        var v35 uint64 = 18446744073709551615
                        v3 = float32((0.1 - ((float64((((float32(1e10) * v3) + ((0.1 - 3.14159) * v3)) / v3)) + (3.14159 / 1.5)) - 1.5)))
                        switch ((("hello" + "") + (("foo bar" + ("a" + "")) + "")) + ("" + "a")) {
                        case "foo bar":
                            switch fi18 {
                            case 127:
                                for fi36 := 0; fi36 < 10; fi36++ {
                                    if true {
                                        v37 := GenS5{F0: fi18, F1: "\n\t", F2: 18446744073709551615}
                                        if (!true) {
                                            fmt.Println(v35, v3)
                                            v3 += 0.0
                                        } else {
                                            v3 = (1.5 + 1.5)
                                        }
                                        v35 += v35
                                        _ = v37
                                    }
                                    v38 := GenS6{F0: 1, F1: 1, F2: -637}
                                    _ = fi36
                                    _ = v38
                                }
                            }
                        default:
                            switch "" {
                            case "hello":
                                for fi39 := 0; fi39 < 19; fi39++ {
                                    v40 := GenS7{F0: 0, F1: -1, F2: 0}
                                    _ = fi39
                                    _ = v40
                                }
                            case "foo bar":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                v35 = 414
                            }
                        }
                        _ = v35
                    }
                    if false {
                        switch ("foo bar" + "") {
                        case "foo bar":
                            switch "\n\t" {
                            case "\n\t":
                                var v41 float64 = float64((0.1 / ((0.0 - 0.1) - (v3 - v3))))
                                _ = v41
                            }
                        case "a":
                            switch "\n\t" {
                            case "":
                                if true {
                                    switch "a" {
                                    case "hello":
                                        switch fi18 {
                                        case 2147483647:
                                            var v42 uint64 = 2
                                            _ = v42
                                        case 1:
                                            for fi43 := 0; fi43 < 2; fi43++ {
                                                var v44 int64 = 9223372036854775807
                                                var v45 uint64 = 4294967295
                                                v45 = (v45 | v45)
                                                _ = fi43
                                                _ = v44
                                                _ = v45
                                            }
                                        case -152:
                                            _ = len(v1)
                                        }
                                    case "":
                                        v3 = v3
                                    default:
                                        var v46 int = (fi18 ^ -674)
                                        _ = v46
                                    }
                                }
                            }
                        default:
                            for fi47 := 0; fi47 < 18; fi47++ {
                                v3 += v3
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = len(v1)
                                fmt.Println(fi18, fi18, fi47)
                                _ = fi47
                            }
                        }
                    } else {
                        if (!(!(65535 <= 882))) {
                            var v48 string = "hello"
                            _ = v48
                        } else {
                            if true {
                                var v49 int16 = -32768
                                v50 := GenS4{F0: 513}
                                if true {
                                    if ((0.1 / 1.0) < (1.5 * (0.1 + 1e10))) {
                                        var v51 uint8 = 128
                                        for fi52 := 0; fi52 < 4; fi52++ {
                                            v49 = (-(v49))
                                            _ = fi52
                                        }
                                        _ = v51
                                    } else {
                                        switch ("a" + ("a" + "hello")) {
                                        case "":
                                            if (((float64((1.5 * (v3 - float32(-2.25)))) + 1.5) >= 1.5) && (421 <= 2147483647)) {
                                                v49 = v49
                                                if true {
                                                    if true {
                                                        v49 = v49
                                                        var v53 string = ""
                                                        v54 := GenS8{F0: -1, F1: 127}
                                                        _ = v53
                                                        _ = v54
                                                    } else {
                                                        if false {
                                                            v3 += v3
                                                        }
                                                    }
                                                    for fi55 := 0; fi55 < 7; fi55++ {
                                                        v3 = 1.5
                                                        var v56 int = (-(fi18))
                                                        if ((2147483647 != 1) || (!false)) {
                                                            v57 := GenS9{F0: 2, F1: 2, F2: ((false && (true || (!false))) || false)}
                                                            var v58 bool = false
                                                            fmt.Println(fi55)
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                        v59 := GenS1{F0: ((2147483647 <= 2147483647) || ("foo bar" != ("\n\t" + "")))}
                                                        v49 -= v49
                                                        _ = fi55
                                                        _ = v56
                                                        _ = v59
                                                    }
                                                    v3 += float32(0.1)
                                                }
                                                switch 32767 {
                                                case 278:
                                                    var v60 float32 = (v3 + (0.1 + v3))
                                                    _ = v60
                                                case 9223372036854775807:
                                                    v49 += (v49 | v49)
                                                case 42:
                                                    v2 = append(v2, "\n\t")
                                                default:
                                                    var v61 int64 = -9223372036854775808
                                                    _ = v61
                                                }
                                            } else {
                                                v49 -= v49
                                            }
                                        case "foo bar":
                                            for fi62 := 0; fi62 < 10; fi62++ {
                                                v49 = (-(v49))
                                                _ = fi62
                                            }
                                        default:
                                            v49 = v49
                                        }
                                        switch "a" {
                                        case "hello":
                                            var v63 uint = 0
                                            _ = v63
                                        case "a":
                                            var v64 int32 = 0
                                            _ = v64
                                        case "":
                                            switch (("foo bar" + ("\n\t" + "foo bar")) + "a") {
                                            case "foo bar":
                                                var v65 float64 = 1.5
                                                _ = v65
                                            case "hello":
                                                if true {
                                                    var v66 float32 = float32((1.5 / -2.25))
                                                    v3 = v66
                                                    fmt.Println(fi18)
                                                    _ = v66
                                                }
                                            case "a":
                                                if false {
                                                    v49 = (v49 | (v49 ^ v49))
                                                    if (4294967295 >= 2) {
                                                        v49 -= (v49 << ((4294967295) & 63))
                                                        var v67 uint8 = 46
                                                        var v68 float32 = 1.5
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                } else {
                                                    switch ("" + "a") {
                                                    case "a":
                                                        fmt.Println(v49)
                                                    case "":
                                                        var v69 int16 = v49
                                                        _ = v69
                                                    default:
                                                        for fi70 := 0; fi70 < 19; fi70++ {
                                                            var v71 uint8 = 248
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            _ = fi70
                                                            _ = v71
                                                        }
                                                    }
                                                    for fi72 := 0; fi72 < 4; fi72++ {
                                                        v73 := GenS3{F0: 43, F1: "foo bar", F2: 65535}
                                                        _ = fi72
                                                        _ = v73
                                                    }
                                                    var v74 uint = 4294967295
                                                    _ = v74
                                                }
                                            }
                                        default:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    }
                                    if false {
                                        fmt.Println(v49, fi18)
                                        var v75 int16 = -714
                                        _ = v75
                                    } else {
                                        var v76 int64 = 42
                                        _ = v76
                                    }
                                } else {
                                    fmt.Println(v49)
                                    if true {
                                        v77 := GenS2{F0: -71}
                                        _ = v77
                                    } else {
                                        var v78 int16 = v49
                                        v78 += (^(v49))
                                        v79 := GenS2{F0: -128}
                                        _ = v78
                                        _ = v79
                                    }
                                }
                                _ = v49
                                _ = v50
                            } else {
                                if false {
                                    var v80 uint8 = 1
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v80
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = 383
                                    }
                                    for fi81 := 0; fi81 < 2; fi81++ {
                                        var v82 uint8 = 2
                                        _ = fi81
                                        _ = v82
                                    }
                                }
                            }
                            var v83 uint16 = 42
                            v83 = (v83 | 490)
                            _ = v83
                        }
                    }
                    _ = fi18
                }
                _ = v17
            }
        }
        v3 -= (v3 - (v3 - (v3 + (((1.5 - 0.0) - v3) - 3.14159))))
    } else {
        for fi84 := 0; fi84 < 12; fi84++ {
            if (!(1 <= -32768)) {
                var v85 uint = 1
                var v86 uint16 = 42
                _ = v85
                _ = v86
            } else {
                var v87 uint = 2
                if ((v87 != (v87 + (^(v87)))) && (-9223372036854775808 >= 9223372036854775807)) {
                    if len(v2) > 0 {
                        v2[0] = (("a" + "\n\t") + "a")
                    }
                    for fi88 := 0; fi88 < 6; fi88++ {
                        v3 -= 0.0
                        var v89 int16 = -1
                        switch fi84 {
                        case 9223372036854775807:
                            v89 = (v89 << ((v87) & 63))
                        case -463:
                            v90 := GenS8{F0: -1, F1: 42}
                            _ = v90
                        case -9223372036854775808:
                            if (0 > -128) {
                                fmt.Println(v89, fi88, v89)
                            } else {
                                switch "" {
                                case "":
                                    v1 = append(v1, 826)
                                case "\n\t":
                                    var v91 float64 = float64(v3)
                                    if true {
                                        v3 -= -2.25
                                    } else {
                                        _ = len(v1)
                                    }
                                    v92 := GenS10{F0: 2147483647}
                                    _ = v91
                                    _ = v92
                                case "foo bar":
                                    var v93 uint64 = 705
                                    _ = v93
                                default:
                                    fmt.Println(fi88)
                                }
                                v94 := GenS7{F0: 54, F1: 1, F2: 83}
                                var v95 float64 = -2.25
                                _ = v94
                                _ = v95
                            }
                        default:
                            var v96 int64 = -1
                            _ = v96
                        }
                        _ = fi88
                        _ = v89
                    }
                }
                _ = v87
            }
            _ = fi84
        }
        v97 := GenS3{F0: 106, F1: "a", F2: 32767}
        _ = v97
    }
    for fi98 := 0; fi98 < 20; fi98++ {
        var v99 bool = (-1 <= 32767)
        fmt.Println(fi98, v99)
        _ = fi98
        _ = v99
    }
    v3 = (v3 - float32(1e10))
    var v100 bool = (1e10 != 0.0)
    v3 += 1.5
    var v101 int32 = 0
    for fi102 := 0; fi102 < 3; fi102++ {
        var v103 int = fi102
        var v104 int = (v103 << ((18446744073709551615) & 63))
        _ = fi102
        _ = v103
        _ = v104
    }
    for fi105 := 0; fi105 < 17; fi105++ {
        fmt.Println(v3)
        switch fi105 {
        case 128:
            fmt.Println(v100)
        default:
            v101 += -366
        }
        _ = fi105
    }
    fmt.Println(v101, v3)
    for fi106 := 0; fi106 < 14; fi106++ {
        v100 = true
        v101 = -2147483648
        _ = fi106
    }
    switch ("a" + "hello") {
    case "a":
        fmt.Println(v100, v100, v100)
    case "foo bar":
        if v100 {
            fmt.Println(v3)
            for fi107 := 0; fi107 < 15; fi107++ {
                if true {
                    if true {
                        v1 = append(v1, 2147483647)
                    }
                    fmt.Println(v100, v100)
                    var v108 uint64 = 2
                    _ = v108
                } else {
                    v100 = (4294967295 != 2)
                    var v109 int64 = 746
                    _ = v109
                }
                v110 := GenS7{F0: 562, F1: (v101 + v101), F2: 183}
                _ = fi107
                _ = v110
            }
            for fi111 := 0; fi111 < 2; fi111++ {
                v2 = append(v2, "hello")
                _ = fi111
            }
        } else {
            var v112 string = ""
            v3 -= (((v3 / v3) + v3) / (v3 + (0.0 / (v3 / v3))))
            fmt.Println(v100)
            _ = v112
        }
    case "\n\t":
        v3 -= (v3 + (v3 * v3))
    }
    v1 = append(v1, 309)
    var v113 float32 = v3
    var v114 int8 = 1
    var v115 uint8 = 127
    var v116 uint = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v113:", v113)
    fmt.Println("v114:", v114)
    fmt.Println("v115:", v115)
    fmt.Println("v116:", v116)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v100
    _ = v101
    _ = v113
    _ = v114
    _ = v115
    _ = v116
    _ = v2
    _ = v3
}
