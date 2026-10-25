package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int8; F1 int32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int64; F1 uint32; F2 bool }
type GenS10 struct { F0 uint }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]uint32)
    var v3 bool = ((^(v1)) != (v1 >> ((v1) & 63)))
    var v4 bool = (((((0.1 * ((1e10 * 3.14159) - 3.14159)) - (float64(1.5) - (3.14159 * 0.1))) + 0.1) + float64(-2.25)) > 3.14159)
    v5 := []int{128, -9223372036854775808, 0, -895}
    var v6 float32 = 0.1
    delete(v2, -2147483648)
    if v3 {
        var v7 float32 = 1.5
        if (2 > 42) {
            v7 += (v6 - v7)
            if false {
                if v3 {
                    var v8 int32 = -155
                    _ = v8
                }
                v6 = (v6 - v7)
                var v9 uint32 = 320
                _ = v9
            }
            v6 += v7
        } else {
            if v3 {
                var v10 float64 = (-2.25 / float64(v7))
                _ = v10
            } else {
                var v11 float32 = (float32(0.0) / 1.5)
                if (false && v4) {
                    var v12 string = "hello"
                    var v13 bool = (297 > 32767)
                    _ = v12
                    _ = v13
                } else {
                    switch ("a" + ("hello" + (("" + "hello") + ("a" + "foo bar")))) {
                    case "":
                        v4 = false
                    case "hello":
                        v6 -= (0.0 * float32(0.0))
                    case "\n\t":
                        switch 9223372036854775807 {
                        case 128:
                            for fi14 := 0; fi14 < 19; fi14++ {
                                var v15 float32 = 1.5
                                for fi16 := 0; fi16 < 14; fi16++ {
                                    v17 := GenS1{F0: 9223372036854775807, F1: uint16(v1), F2: -624}
                                    _ = fi16
                                    _ = v17
                                }
                                for fi18 := 0; fi18 < 13; fi18++ {
                                    _ = len(v5)
                                    _ = fi18
                                }
                                _ = fi14
                                _ = v15
                            }
                        case 0:
                            for fi19 := 0; fi19 < 15; fi19++ {
                                v20 := GenS1{F0: 487, F1: 65535, F2: -164}
                                v1 = (^(v1))
                                _ = fi19
                                _ = v20
                            }
                        }
                    default:
                        v4 = ((0 != -643) && v3)
                    }
                }
                _ = v11
            }
            if (v4 && false) {
                for fi21 := 0; fi21 < 8; fi21++ {
                    if v4 {
                        var v22 int8 = -1
                        for fi23 := 0; fi23 < 11; fi23++ {
                            var v24 int64 = 9223372036854775807
                            _ = fi23
                            _ = v24
                        }
                        _ = v22
                    }
                    v25 := GenS2{F0: 1}
                    if v4 {
                        for fi26 := 0; fi26 < 11; fi26++ {
                            if (!v3) {
                                v4 = v3
                                var v27 bool = (v4 && (fi26 == fi26))
                                _ = v27
                            } else {
                                switch (^(fi26)) {
                                case 32768:
                                    v6 = float32((3.14159 + (1e10 / (1.5 - 1e10))))
                                case 32767:
                                    var v28 uint = (v1 << (((v1 * (v1 << (((v1 | (v1 % (((v1 * (^(v1)))) | 1)))) & 63)))) & 63))
                                    _ = v28
                                default:
                                    var v29 uint8 = 1
                                    _ = v29
                                }
                                var v30 uint32 = 2147483647
                                _ = v30
                            }
                            var v31 uint = 4294967295
                            var v32 float32 = v6
                            _ = fi26
                            _ = v31
                            _ = v32
                        }
                    }
                    _ = fi21
                    _ = v25
                }
                var v33 int16 = -32768
                _ = v33
            }
            v6 += (0.0 + float32((0.1 * 0.0)))
        }
        _ = v7
    }
    for fi34 := 0; fi34 < 3; fi34++ {
        v6 = v6
        _ = fi34
    }
    fmt.Println(v3, v4)
    switch -545 {
    case -745:
        var v35 int64 = -416
        _ = v35
    case -444:
        v4 = (v1 <= v1)
    case 187:
        for fi36 := 0; fi36 < 11; fi36++ {
            v4 = v4
            v6 -= -2.25
            _ = fi36
        }
    default:
        v6 -= v6
    }
    v1 -= v1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch (("\n\t" + "hello") + "foo bar") {
    case "foo bar":
        if (v4 && v3) {
            var v37 float32 = 0.1
            v6 = v6
            for fi38 := 0; fi38 < 18; fi38++ {
                for fi39 := 0; fi39 < 11; fi39++ {
                    if len(v5) > 0 {
                        v5[0] = (fi39 | fi38)
                    }
                    var v40 uint16 = 22
                    var v41 int16 = 1
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                switch "\n\t" {
                case "a":
                    var v42 int = 4294967295
                    _ = v42
                case "hello":
                    for fi43 := 0; fi43 < 19; fi43++ {
                        if v4 {
                            fmt.Println(v4)
                            var v44 int64 = 9223372036854775807
                            v45 := GenS1{F0: v44, F1: 77, F2: 2147483647}
                            _ = v44
                            _ = v45
                        }
                        v4 = v4
                        var v46 float64 = (1.5 / ((float64(v37) * (3.14159 * (1.5 * 3.14159))) * -2.25))
                        delete(v2, fi38)
                        v4 = v3
                        if v3 {
                            v47 := GenS3{F0: 342, F1: -9223372036854775808}
                            delete(v2, fi43)
                            _ = v47
                        }
                        _ = fi43
                        _ = v46
                    }
                case "":
                    var v48 float64 = 0.1
                    _ = v48
                }
                var v49 float32 = v6
                _ = fi38
                _ = v49
            }
            _ = v37
        } else {
            var v50 uint8 = 0
            _ = v50
        }
    default:
        v1 = v1
    }
    delete(v2, 32768)
    v1 = v1
    if v4 {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if ((2 != 18446744073709551615) || false) {
            if (true && v3) {
                if (false && false) {
                    if false {
                        _ = len(v2)
                    } else {
                        v3 = (48 != 42)
                    }
                    if v3 {
                        v3 = (!false)
                    } else {
                        v4 = ((v1 / ((18446744073709551615) | 1)) == 255)
                        v3 = true
                        var v51 float32 = v6
                        _ = v51
                    }
                    if v4 {
                        if v3 {
                            v4 = v3
                            for fi52 := 0; fi52 < 12; fi52++ {
                                for fi53 := 0; fi53 < 20; fi53++ {
                                    if len(v5) > 0 {
                                        v5[0] = (fi52 | fi53)
                                    }
                                    _ = fi53
                                }
                                v54 := GenS3{F0: 671, F1: 0}
                                _ = fi52
                                _ = v54
                            }
                        }
                        switch (("hello" + ("hello" + "foo bar")) + "hello") {
                        case "hello":
                            if v4 {
                                for fi55 := 0; fi55 < 3; fi55++ {
                                    v56 := GenS4{F0: v4}
                                    v3 = v3
                                    var v57 uint64 = 278
                                    _ = fi55
                                    _ = v56
                                    _ = v57
                                }
                                v5 = append(v5, 9223372036854775807)
                            } else {
                                v58 := GenS5{F0: 426, F1: 542}
                                v59 := GenS6{F0: 528}
                                v3 = false
                                _ = v58
                                _ = v59
                            }
                        default:
                            v1 -= 4294967295
                        }
                        for fi60 := 0; fi60 < 11; fi60++ {
                            fmt.Println(v3, v1, v6)
                            _ = fi60
                        }
                    } else {
                        if v4 {
                            for fi61 := 0; fi61 < 4; fi61++ {
                                var v62 int64 = 42
                                v6 -= float32(float64(3.14159))
                                v6 = ((v6 * 1e10) / (float32(((0.1 - 3.14159) * (float64(v6) * 0.1))) / v6))
                                _ = fi61
                                _ = v62
                            }
                            var v63 int32 = -1
                            v2[4294967295] = 518
                            _ = v63
                        }
                        switch 42 {
                        case -957:
                            var v64 int = 2147483647
                            _ = v64
                        case 127:
                            var v65 bool = false
                            _ = v65
                        }
                        v2[405] = 2147483647
                    }
                } else {
                    v4 = v3
                    v66 := GenS1{F0: -1, F1: 545, F2: int32(v1)}
                    _ = v66
                }
            }
            v67 := GenS7{F0: 18, F1: 1}
            switch 849 {
            case -496:
                fmt.Println(v6, v3)
            }
            _ = v67
        } else {
            v4 = (1 != 0)
        }
        var v68 string = ("\n\t" + (("\n\t" + "") + "hello"))
        _ = v68
    } else {
        fmt.Println(v3, v4)
        for fi69 := 0; fi69 < 9; fi69++ {
            var v70 int64 = 857
            _ = fi69
            _ = v70
        }
        if (0 >= 0) {
            var v71 int = 0
            v5 = append(v5, 2147483647)
            if v3 {
                switch ((("\n\t" + "hello") + "") + "") {
                case "hello":
                    if (float32(0.0) != ((v6 + v6) / v6)) {
                        v72 := GenS1{F0: -295, F1: 32767, F2: int32(v1)}
                        switch ("" + ("a" + "a")) {
                        case "hello":
                            fmt.Println(v6, v4)
                        }
                        if (!(!v3)) {
                            v4 = (!(!(true && ("\n\t" >= "foo bar"))))
                            switch "hello" {
                            case "foo bar":
                                switch "\n\t" {
                                case "a":
                                    if v3 {
                                        for fi73 := 0; fi73 < 14; fi73++ {
                                            if v3 {
                                                v71 = fi73
                                            }
                                            var v74 int = -209
                                            _ = fi73
                                            _ = v74
                                        }
                                        v71 += 550
                                        fmt.Println(v4, v6, v3)
                                    }
                                case "hello":
                                    if (v3 || (v4 && v3)) {
                                        if v4 {
                                            v5 = append(v5, (v71 - v71))
                                        }
                                    }
                                }
                            }
                        }
                        _ = v72
                    } else {
                        v75 := GenS3{F0: 2147483647, F1: 9223372036854775807}
                        var v76 bool = true
                        if false {
                            if (false && false) {
                                v76 = (127 > 49)
                                if true {
                                    if v3 {
                                        v76 = false
                                        if v3 {
                                            v6 -= ((0.1 * (0.0 * (v6 + v6))) * v6)
                                            for fi77 := 0; fi77 < 8; fi77++ {
                                                switch v71 {
                                                case 128:
                                                    v3 = v3
                                                }
                                                _ = fi77
                                            }
                                            for fi78 := 0; fi78 < 20; fi78++ {
                                                fmt.Println(v4, v6, v3)
                                                _ = fi78
                                            }
                                        } else {
                                            for fi79 := 0; fi79 < 20; fi79++ {
                                                for fi80 := 0; fi80 < 14; fi80++ {
                                                    var v81 uint64 = 609
                                                    v82 := GenS8{F0: 110}
                                                    _ = len(v2)
                                                    _ = fi80
                                                    _ = v81
                                                    _ = v82
                                                }
                                                var v83 int8 = -99
                                                v84 := GenS3{F0: 1, F1: -1}
                                                _ = fi79
                                                _ = v83
                                                _ = v84
                                            }
                                            if (!(!v4)) {
                                                fmt.Println(v76, v3, v4)
                                                switch (^(v71)) {
                                                case -1:
                                                    if false {
                                                        v76 = false
                                                        switch ("hello" + "") {
                                                        case "foo bar":
                                                            var v85 int = v71
                                                            _ = v85
                                                        case "a":
                                                            switch "hello" {
                                                            case "hello":
                                                                v71 = -2147483648
                                                            default:
                                                                if false {
                                                                    if true {
                                                                        v86 := v2[v71]
                                                                        var v87 int8 = 127
                                                                        _ = v86
                                                                        _ = v87
                                                                    }
                                                                }
                                                            }
                                                        case "":
                                                            v88 := GenS9{F0: -9223372036854775808, F1: 749, F2: (295 >= -32768)}
                                                            _ = v88
                                                        }
                                                    } else {
                                                        var v89 int32 = 42
                                                        v90 := GenS3{F0: 2147483647, F1: 42}
                                                        v6 += float32((3.14159 / float64((0.1 + (((0.1 * ((v6 * (float32((((0.1 - (0.1 + -2.25)) + (0.0 / 3.14159)) / (0.0 - (float64(0.1) / 3.14159)))) / ((float32(0.1) * 1.5) + 0.0))) - v6)) - 1e10) + v6)))))
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                default:
                                                    var v91 uint64 = 1
                                                    _ = v91
                                                }
                                                v92 := GenS10{F0: (v1 / ((55) | 1))}
                                                _ = v92
                                            }
                                        }
                                    }
                                    v93 := GenS5{F0: 0, F1: 499}
                                    if v3 {
                                        _ = len(v2)
                                    }
                                    _ = v93
                                } else {
                                    if v3 {
                                        fmt.Println(v3, v6)
                                        v6 = float32(-2.25)
                                        fmt.Println(v3, v1)
                                    }
                                    fmt.Println(v71, v76, v76)
                                }
                            } else {
                                for fi94 := 0; fi94 < 11; fi94++ {
                                    for fi95 := 0; fi95 < 7; fi95++ {
                                        v4 = v76
                                        v4 = (696 < 37)
                                        fmt.Println(v71, v1, fi95)
                                        _ = fi95
                                    }
                                    v4 = v4
                                    _ = fi94
                                }
                            }
                        } else {
                            _ = len(v2)
                            var v96 int32 = -1
                            _ = v96
                        }
                        _ = v75
                        _ = v76
                    }
                default:
                    for fi97 := 0; fi97 < 5; fi97++ {
                        var v98 uint8 = 2
                        _ = fi97
                        _ = v98
                    }
                }
                for fi99 := 0; fi99 < 5; fi99++ {
                    var v100 int8 = 1
                    _ = fi99
                    _ = v100
                }
            }
            _ = v71
        } else {
            for fi101 := 0; fi101 < 8; fi101++ {
                var v102 int16 = -32768
                var v103 uint8 = 1
                if len(v5) > 0 {
                    v5[0] = 32768
                }
                _ = fi101
                _ = v102
                _ = v103
            }
        }
    }
    if (!(0.0 != 0.0)) {
        var v104 uint16 = 2
        switch 4294967295 {
        case 817:
            _ = len(v5)
        default:
            switch "hello" {
            case "foo bar":
                v105 := GenS10{F0: v1}
                _ = v105
            default:
                switch 127 {
                case -326:
                    v4 = (1000 <= 32767)
                case -9223372036854775808:
                    v106 := GenS6{F0: 602}
                    _ = v106
                default:
                    for fi107 := 0; fi107 < 20; fi107++ {
                        var v108 int8 = -125
                        _ = len(v5)
                        v109 := v2[(fi107 + (fi107 & 9223372036854775807))]
                        _ = fi107
                        _ = v108
                        _ = v109
                    }
                }
            }
        }
        var v110 string = ("\n\t" + "hello")
        _ = v104
        _ = v110
    }
    v111 := GenS3{F0: 0, F1: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v111
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
