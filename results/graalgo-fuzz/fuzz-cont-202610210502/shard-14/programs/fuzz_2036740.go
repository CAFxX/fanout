package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int16 }
type GenS2 struct { F0 int64; F1 string; F2 int32 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint16; F1 int32; F2 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 int64; F2 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32; F1 uint }
type GenS10 struct { F0 int; F1 uint16 }
type GenS11 struct { F0 uint64; F1 int16 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int16 }

func main() {
    v1 := []uint32{514, 2147483647}
    var v2 float64 = 1.5
    var v3 int16 = 0
    v4 := make(map[int]int)
    var v5 bool = (true || false)
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 int = fi6
        _ = fi6
        _ = v7
    }
    switch (("foo bar" + "foo bar") + "a") {
    case "\n\t":
        var v8 bool = ((v3 / ((657) | 1)) > v3)
        _ = v8
    case "a":
        _ = len(v1)
    case "foo bar":
        v9 := GenS1{F0: true, F1: 18446744073709551615, F2: 299}
        _ = v9
    default:
        var v10 int32 = -2147483648
        _ = v10
    }
    var v11 int64 = 752
    if v5 {
        var v12 uint32 = 2147483648
        var v13 int = 42
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            v5 = (1 > 987)
            _ = fi15
        }
        var v16 int64 = v11
        fmt.Println(fi14, fi14, fi14)
        _ = fi14
        _ = v16
    }
    var v17 uint64 = 18446744073709551615
    fmt.Println(v2)
    if (0.1 > ((-2.25 + 0.0) + float32(v2))) {
        if ("\n\t" > ("foo bar" + ("a" + "\n\t"))) {
            if true {
                var v18 uint16 = 496
                for fi19 := 0; fi19 < 17; fi19++ {
                    if true {
                        if false {
                            if (v5 || v5) {
                                if v5 {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v20 uint16 = 0
                                    _ = v20
                                }
                            } else {
                                v21 := GenS1{F0: v5, F1: 0, F2: (v3 - (^(v3)))}
                                if (((fi19 % ((fi19) | 1)) >= fi19) || (int64(fi19) == (v11 & v11))) {
                                    if v5 {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        switch (fi19 ^ 0) {
                                        case 32768:
                                            var v22 uint32 = 0
                                            _ = v22
                                        case 830:
                                            v23 := GenS1{F0: (!v5), F1: 0, F2: v3}
                                            _ = v23
                                        default:
                                            v2 -= (v2 + v2)
                                        }
                                    } else {
                                        v2 = float64(-2.25)
                                        v5 = (835 < -158)
                                        fmt.Println(v5, fi19, v11)
                                    }
                                    switch ("a" + "\n\t") {
                                    case "":
                                        v1 = append(v1, 580)
                                    default:
                                        var v24 uint16 = v18
                                        _ = v24
                                    }
                                    v25 := GenS1{F0: ((!v5) || true), F1: 121, F2: -978}
                                    _ = v25
                                } else {
                                    for fi26 := 0; fi26 < 3; fi26++ {
                                        var v27 uint = 255
                                        _ = fi26
                                        _ = v27
                                    }
                                }
                                _ = v21
                            }
                        } else {
                            if (-1 >= 42) {
                                if (1000 <= v3) {
                                    v28 := GenS2{F0: -705, F1: ("foo bar" + "hello"), F2: 265}
                                    if ((v17 * (v17 << ((472) & 63))) != v17) {
                                        var v29 int8 = 111
                                        _ = v29
                                    }
                                    _ = v28
                                } else {
                                    var v30 uint32 = 836
                                    v3 += -480
                                    for fi31 := 0; fi31 < 8; fi31++ {
                                        v11 = v11
                                        var v32 int = fi19
                                        _ = fi31
                                        _ = v32
                                    }
                                    _ = v30
                                }
                            }
                        }
                        for fi33 := 0; fi33 < 20; fi33++ {
                            var v34 uint64 = v17
                            v11 = (v11 & -830)
                            if (((^(v34)) == 879) && true) {
                                v35 := GenS3{F0: -300, F1: 65535}
                                _ = v35
                            } else {
                                switch 2147483647 {
                                case 503:
                                    fmt.Println(v2)
                                }
                                v18 -= v18
                            }
                            _ = fi33
                            _ = v34
                        }
                    } else {
                        switch fi19 {
                        case 128:
                            v2 -= float64(0.1)
                        }
                        v11 += (-(v11))
                        var v36 int8 = -1
                        _ = v36
                    }
                    v17 = 42
                    v1 = append(v1, 42)
                    _ = fi19
                }
                _ = v18
            } else {
                if false {
                    fmt.Println(v17)
                } else {
                    var v37 int32 = 0
                    _ = v37
                }
                delete(v4, -462)
            }
        }
    } else {
        if true {
            var v38 uint = 65535
            _ = v38
        } else {
            var v39 bool = v5
            if (!v39) {
                switch "" {
                case "\n\t":
                    if (v17 == (v17 + (v17 / (((v17 * 0)) | 1)))) {
                        for fi40 := 0; fi40 < 1; fi40++ {
                            for fi41 := 0; fi41 < 10; fi41++ {
                                switch (("\n\t" + "foo bar") + "hello") {
                                case "\n\t":
                                    switch fi41 {
                                    case 9223372036854775807:
                                        switch fi40 {
                                        case 128:
                                            if (1e10 <= 1e10) {
                                                switch (-(fi41)) {
                                                case -2147483648:
                                                    if len(v1) > 0 {
                                                        v1[0] = 2
                                                    }
                                                case 191:
                                                    fmt.Println(v17, v11, v2)
                                                case 42:
                                                    var v42 uint = 42
                                                    _ = v42
                                                default:
                                                    v39 = v39
                                                }
                                                if true {
                                                    fmt.Println(v11)
                                                    switch fi40 {
                                                    case 1:
                                                        var v43 uint = 267
                                                        _ = v43
                                                    }
                                                }
                                                v17 -= v17
                                            }
                                        case 32767:
                                            var v44 uint = 18446744073709551615
                                            _ = v44
                                        case -387:
                                            v5 = v5
                                        }
                                    }
                                case "hello":
                                    v45 := GenS4{F0: 65535, F1: 1, F2: (v17 % ((v17) | 1))}
                                    _ = v45
                                }
                                var v46 float64 = 0.0
                                switch fi41 {
                                case -967:
                                    v46 = v2
                                case 128:
                                    v4[(fi40 | (fi40 & fi41))] = (fi41 >> ((232) & 63))
                                    var v47 uint64 = v17
                                    switch fi41 {
                                    case 32768:
                                        v1 = append(v1, 2147483648)
                                    default:
                                        v17 += v47
                                    }
                                    v3 -= -32768
                                    v48 := GenS5{F0: 435, F1: -1, F2: 974}
                                    if ((42 != 1) && (!v5)) {
                                        switch "foo bar" {
                                        case "hello":
                                            v47 -= 42
                                        case "":
                                            _ = len(v4)
                                        case "\n\t":
                                            v49 := GenS2{F0: v11, F1: "hello", F2: -2147483648}
                                            _ = v49
                                        default:
                                            v47 = 131
                                        }
                                        var v50 float32 = (3.14159 + 1.5)
                                        v51 := GenS6{F0: 68}
                                        _ = v50
                                        _ = v51
                                    }
                                    _ = v47
                                    _ = v48
                                case 4294967295:
                                    v46 += (((1.5 - 1.5) - ((float64((0.1 * (float32((v2 / float64(-2.25))) - -2.25))) * (v46 + v2)) - v46)) / v2)
                                default:
                                    var v52 int = fi40
                                    _ = v52
                                }
                                _ = fi41
                                _ = v46
                            }
                            switch (fi40 >> ((18446744073709551615) & 63)) {
                            case 4294967295:
                                v3 -= 42
                            case 306:
                                fmt.Println(v5, v3)
                            case 32768:
                                delete(v4, (fi40 >> ((255) & 63)))
                            default:
                                if true {
                                    v53 := v4[fi40]
                                    var v54 uint8 = 255
                                    _ = v53
                                    _ = v54
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                            _ = fi40
                        }
                    } else {
                        fmt.Println(v2, v39)
                    }
                default:
                    for fi55 := 0; fi55 < 10; fi55++ {
                        if (!(2147483647 < 1)) {
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                        } else {
                            v2 -= v2
                        }
                        if true {
                            v56 := GenS5{F0: 209, F1: -360, F2: -20}
                            var v57 float32 = -2.25
                            _ = v56
                            _ = v57
                        } else {
                            var v58 int32 = -579
                            _ = v58
                        }
                        for fi59 := 0; fi59 < 5; fi59++ {
                            fmt.Println(fi55)
                            var v60 uint16 = 763
                            var v61 string = ((((("foo bar" + ("foo bar" + "hello")) + "a") + "foo bar") + ("\n\t" + "hello")) + "a")
                            _ = fi59
                            _ = v60
                            _ = v61
                        }
                        _ = fi55
                    }
                }
            } else {
                v3 = -32768
                if ((((1.5 > float32(v2)) || (v5 || (v39 || v39))) || (!v39)) && (!(-2147483648 == 2147483647))) {
                    switch 42 {
                    case 128:
                        for fi62 := 0; fi62 < 5; fi62++ {
                            fmt.Println(fi62, v5)
                            _ = len(v4)
                            _ = fi62
                        }
                    case 719:
                        for fi63 := 0; fi63 < 16; fi63++ {
                            fmt.Println(v11, v5, v2)
                            _ = fi63
                        }
                    case 810:
                        if true {
                            var v64 uint32 = 65535
                            _ = v64
                        }
                    default:
                        var v65 int = 9223372036854775807
                        _ = v65
                    }
                    if false {
                        if (v3 <= v3) {
                            switch "hello" {
                            case "":
                                v66 := GenS5{F0: 255, F1: 2147483647, F2: 203}
                                _ = v66
                            case "\n\t":
                                v2 += (-2.25 + 3.14159)
                            case "a":
                                for fi67 := 0; fi67 < 6; fi67++ {
                                    fmt.Println(v5)
                                    _ = fi67
                                }
                            default:
                                var v68 int8 = 1
                                _ = v68
                            }
                            if v5 {
                                v3 += v3
                                v69 := GenS5{F0: 65535, F1: 750, F2: -540}
                                if v5 {
                                    v70 := GenS1{F0: (v39 || (1 > 32767)), F1: 2, F2: (v3 - (v3 / ((1000) | 1)))}
                                    var v71 int = 4294967295
                                    _ = v70
                                    _ = v71
                                }
                                _ = v69
                            }
                            if (42 >= 991) {
                                switch "a" {
                                case "hello":
                                    var v72 int64 = (v11 << ((889) & 63))
                                    _ = v72
                                case "foo bar":
                                    v73 := GenS2{F0: (v11 << ((2) & 63)), F1: ("foo bar" + "hello"), F2: 586}
                                    _ = v73
                                case "a":
                                    for fi74 := 0; fi74 < 20; fi74++ {
                                        fmt.Println(v3)
                                        if v39 {
                                            if v5 {
                                                v75 := GenS4{F0: 989, F1: 1, F2: 2}
                                                switch fi74 {
                                                case 128:
                                                    var v76 string = (("hello" + (("foo bar" + "foo bar") + "a")) + ("\n\t" + "foo bar"))
                                                    _ = v76
                                                case 298:
                                                    v77 := GenS7{F0: (fi74 % ((fi74) | 1)), F1: v11, F2: (!(v5 && (42 >= 127)))}
                                                    _ = v77
                                                case 1:
                                                    for fi78 := 0; fi78 < 3; fi78++ {
                                                        v3 = v3
                                                        _ = fi78
                                                    }
                                                }
                                                fmt.Println(fi74, v39)
                                                _ = v75
                                            }
                                            var v79 float32 = (((-2.25 + -2.25) - 0.1) - float32(v2))
                                            v80 := GenS8{F0: ("foo bar" + "a")}
                                            _ = v79
                                            _ = v80
                                        }
                                        v81 := GenS5{F0: 65535, F1: 1, F2: 667}
                                        _ = fi74
                                        _ = v81
                                    }
                                default:
                                    v82 := GenS9{F0: -185, F1: 65535}
                                    _ = v82
                                }
                                var v83 uint = 479
                                _ = v83
                            } else {
                                v17 -= v17
                                var v84 uint32 = 65535
                                var v85 bool = (0 >= 0)
                                _ = v84
                                _ = v85
                            }
                        }
                        v3 = (v3 * (-(v3)))
                    } else {
                        if false {
                            v86 := GenS10{F0: 2147483647, F1: 23}
                            _ = v86
                        }
                        for fi87 := 0; fi87 < 10; fi87++ {
                            v3 += 192
                            v88 := GenS11{F0: v17, F1: -1}
                            for fi89 := 0; fi89 < 4; fi89++ {
                                fmt.Println(v2, fi87, v11)
                                var v90 int8 = 0
                                _ = fi89
                                _ = v90
                            }
                            _ = fi87
                            _ = v88
                        }
                    }
                    for fi91 := 0; fi91 < 16; fi91++ {
                        if (51 == -83) {
                            var v92 uint = 2
                            _ = v92
                        }
                        _ = fi91
                    }
                } else {
                    var v93 int8 = -128
                    _ = v93
                }
            }
            v17 += (^(v17))
            _ = v39
        }
        v3 = (-(v3))
        var v94 bool = false
        _ = v94
    }
    for fi95 := 0; fi95 < 1; fi95++ {
        v1 = append(v1, 0)
        var v96 uint64 = (v17 ^ (v17 | 18446744073709551615))
        v97 := GenS12{F0: 42}
        _ = fi95
        _ = v96
        _ = v97
    }
    v98 := GenS13{F0: v3}
    v3 = -32768
    v3 -= v3
    v3 = int16(v17)
    for fi99 := 0; fi99 < 2; fi99++ {
        v5 = (v17 > (v17 % ((v17) | 1)))
        v5 = (0 == 2)
        v100 := GenS1{F0: v5, F1: 1, F2: v3}
        _ = fi99
        _ = v100
    }
    var v101 int = int(v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v101:", v101)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v101
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v98
}
