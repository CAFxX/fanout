package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int; F1 uint32; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint64; F1 string; F2 int16 }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS8 struct { F0 int16; F1 string }
type GenS9 struct { F0 int; F1 string }
type GenS10 struct { F0 int64; F1 string; F2 bool }
type GenS11 struct { F0 uint8; F1 bool; F2 int }
type GenS12 struct { F0 bool; F1 string }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 int; F1 uint32; F2 int8 }
func genFn0() int {
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            switch (fi6 * fi5) {
            case 1:
                var v7 float64 = 0.0
                _ = v7
            case -59:
                var v8 int32 = -150
                _ = v8
            case 127:
                switch 32768 {
                case 42:
                    v9 := GenS1{F0: 0, F1: (fi6 - (fi6 - 4294967295))}
                    _ = v9
                default:
                    var v10 int8 = 1
                    _ = v10
                }
            default:
                var v11 int32 = 214
                _ = v11
            }
            _ = fi6
        }
        var v12 uint8 = 127
        for fi13 := 0; fi13 < 20; fi13++ {
            v14 := GenS2{F0: -1, F1: 2147483648, F2: 987}
            if true {
                v15 := GenS3{F0: 2}
                v16 := []int{(-(fi13)), fi5}
                _ = v15
                _ = v16
            }
            var v17 uint = 2
            _ = fi13
            _ = v14
            _ = v17
        }
        _ = fi5
        _ = v12
    }
    v18 := []int{2147483647}
    var v19 int16 = 1
    switch 127 {
    case 32768:
        var v20 float32 = 0.1
        _ = v20
    case 128:
        if true {
            var v21 float64 = float64((3.14159 + 1.5))
            v22 := GenS2{F0: -608, F1: 65535, F2: 42}
            v18 = append(v18, -1)
            _ = v21
            _ = v22
        }
    default:
        var v23 int16 = (v19 | v19)
        _ = v23
    }
    _ = v18
    _ = v19
    return -9223372036854775808
}
func genFn1(p0 int16, p1 int, p2 uint64) string {
    var v24 int64 = 1
    v25 := GenS4{F0: false, F1: 489, F2: 42}
    switch 934 {
    case -1:
        var v26 int64 = (v24 >> ((18) & 63))
        _ = v26
    default:
        var v27 int = 912
        _ = v27
    }
    var v28 string = ("foo bar" + "a")
    var v29 int = (p1 ^ (p1 >> ((2) & 63)))
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    return (v28 + v28)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"", "foo bar"}
    var v3 string = "a"
    v4 := make(map[int]uint32)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = (v3 + v3)
    switch genFn0() {
    case 298:
        for fi30 := 0; fi30 < 8; fi30++ {
            fmt.Println(v3)
            _ = fi30
        }
    case -1:
        switch -2147483648 {
        case 127:
            if (v3 < v3) {
                if (2147483647 == -2147483648) {
                    for fi31 := 0; fi31 < 7; fi31++ {
                        if (((((3.14159 - 0.0) == (0.1 - (3.14159 / 1.5))) || false) && false) || (!(45 > 4294967295))) {
                            switch v3 {
                            case "foo bar":
                                var v32 float32 = float32((3.14159 - 1e10))
                                _ = v32
                            case "a":
                                fmt.Println(fi31)
                            case "\n\t":
                                var v33 uint16 = 255
                                _ = v33
                            default:
                                var v34 bool = (!(false || false))
                                _ = v34
                            }
                        }
                        v3 = v3
                        _ = fi31
                    }
                    if (18446744073709551615 <= 4294967295) {
                        var v35 int16 = 42
                        _ = v35
                    } else {
                        v3 = (v3 + ("a" + (v3 + (v3 + v3))))
                        v36 := GenS3{F0: 2}
                        for fi37 := 0; fi37 < 11; fi37++ {
                            switch "a" {
                            case "\n\t":
                                v38 := GenS4{F0: true, F1: 2147483647, F2: -717}
                                _ = v38
                            default:
                                _ = len(v2)
                            }
                            var v39 int = genFn0()
                            _ = fi37
                            _ = v39
                        }
                        _ = v36
                    }
                }
            }
        case 825:
            if (4294967295 == 939) {
                v3 = v3
                for fi40 := 0; fi40 < 8; fi40++ {
                    if ("foo bar" <= (v3 + v3)) {
                        var v41 bool = (!true)
                        v3 = v3
                        fmt.Println(v41, v41, v3)
                        _ = v41
                    }
                    v1["foo bar"] = 1
                    _ = fi40
                }
                switch 1 {
                case 2147483647:
                    fmt.Println(v3)
                }
            }
        }
    default:
        v3 = v3
    }
    var v42 int16 = -416
    if true {
        if false {
            switch (v3 + "hello") {
            case "foo bar":
                v43 := GenS2{F0: 4294967295, F1: 2147483648, F2: 255}
                _ = v43
            default:
                var v44 int64 = 42
                _ = v44
            }
            if false {
                switch 9223372036854775807 {
                case 4294967295:
                    switch -304 {
                    case -2147483648:
                        var v45 uint64 = 4294967295
                        _ = v45
                    case 4294967295:
                        var v46 int32 = -856
                        _ = v46
                    case 32767:
                        if true {
                            v47 := GenS5{F0: 42, F1: v3, F2: (v42 + (v42 >> ((1) & 63)))}
                            if false {
                                _ = len(v1)
                                var v48 int32 = -456
                                var v49 uint = 0
                                _ = v48
                                _ = v49
                            }
                            _ = v47
                        } else {
                            v4[-2147483648] = 0
                        }
                    default:
                        _ = len(v2)
                    }
                case 1:
                    if true {
                        v3 = v3
                        v42 = (v42 << ((18446744073709551615) & 63))
                    } else {
                        v50 := GenS6{F0: 9223372036854775807, F1: v3}
                        switch 127 {
                        case 127:
                            if (42 == -367) {
                                for fi51 := 0; fi51 < 19; fi51++ {
                                    var v52 float64 = (-2.25 / 1.0)
                                    var v53 int32 = 0
                                    var v54 uint32 = 42
                                    _ = fi51
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                }
                                v3 = ((v3 + v3) + "a")
                            } else {
                                v3 = "a"
                                if false {
                                    if (-2.25 >= ((1.5 * 1.5) - ((0.1 - -2.25) + (-2.25 + 0.1)))) {
                                        fmt.Println(v42, v3)
                                        for fi55 := 0; fi55 < 4; fi55++ {
                                            v56 := GenS5{F0: 42, F1: "", F2: -407}
                                            fmt.Println(v3)
                                            v3 = ("" + ("hello" + v3))
                                            _ = fi55
                                            _ = v56
                                        }
                                        v42 = (v42 * v42)
                                    } else {
                                        switch v3 {
                                        case "a":
                                            switch -840 {
                                            case -137:
                                                v3 = v3
                                            case 32768:
                                                if false {
                                                    v57 := GenS7{F0: 32767, F1: 2, F2: 100}
                                                    _ = v57
                                                } else {
                                                    v58 := GenS8{F0: (v42 - 0), F1: ("\n\t" + v3)}
                                                    _ = v58
                                                }
                                            default:
                                                v1[v3] = 1
                                            }
                                        default:
                                            v3 = (((v3 + v3) + "\n\t") + (v3 + v3))
                                        }
                                        v42 += 0
                                        var v59 float32 = 0.0
                                        _ = v59
                                    }
                                    v3 = "a"
                                    if (float32(3.14159) < 3.14159) {
                                        v60 := v4[738]
                                        v3 = v3
                                        for fi61 := 0; fi61 < 14; fi61++ {
                                            fmt.Println(v42, v3)
                                            _ = fi61
                                        }
                                        _ = v60
                                    }
                                }
                            }
                        case -9223372036854775808:
                            fmt.Println(v3)
                        default:
                            if (!(!true)) {
                                v3 = v3
                            }
                        }
                        _ = v50
                    }
                case 32767:
                    if true {
                        for fi62 := 0; fi62 < 14; fi62++ {
                            if false {
                                v3 = (v3 + v3)
                            }
                            _ = fi62
                        }
                        _ = len(v2)
                        var v63 uint = 489
                        _ = v63
                    } else {
                        v3 = v3
                        v42 += (v42 / (((v42 | -246)) | 1))
                    }
                default:
                    var v64 int8 = 20
                    _ = v64
                }
                v65 := GenS5{F0: 18446744073709551615, F1: "\n\t", F2: (^(v42))}
                _ = v65
            } else {
                var v66 int16 = 610
                switch v3 {
                case "\n\t":
                    v67 := v1[((v3 + v3) + ("hello" + "hello"))]
                    _ = v67
                default:
                    v68 := GenS9{F0: 0, F1: v3}
                    _ = v68
                }
                fmt.Println(v42, v3, v42)
                _ = v66
            }
        }
        v69 := GenS5{F0: 618, F1: (v3 + "\n\t"), F2: v42}
        _ = v69
    }
    delete(v4, -775)
    _ = len(v2)
    v42 += v42
    for fi70 := 0; fi70 < 13; fi70++ {
        v2 = append(v2, v3)
        fmt.Println(v42, v3)
        _ = fi70
    }
    v71 := GenS10{F0: -1, F1: "", F2: (((v3 + v3) <= ("foo bar" + v3)) || ((272 <= 2) || true))}
    switch "" {
    case "hello":
        if true {
            v72 := GenS9{F0: 1, F1: (v3 + (v3 + ("a" + "foo bar")))}
            if false {
                fmt.Println(v3)
                var v73 uint64 = 0
                v74 := v4[-649]
                _ = v73
                _ = v74
            }
            _ = v72
        } else {
            var v75 int16 = 1000
            var v76 uint64 = 0
            for fi77 := 0; fi77 < 16; fi77++ {
                v78 := GenS1{F0: 1, F1: fi77}
                for fi79 := 0; fi79 < 9; fi79++ {
                    v80 := v1["foo bar"]
                    var v81 uint16 = 0
                    _ = fi79
                    _ = v80
                    _ = v81
                }
                _ = fi77
                _ = v78
            }
            _ = v75
            _ = v76
        }
    }
    for fi82 := 0; fi82 < 16; fi82++ {
        v2 = append(v2, v3)
        _ = fi82
    }
    v3 = v3
    if (0 > 42) {
        v83 := GenS10{F0: 42, F1: ((("foo bar" + v3) + (v3 + v3)) + "a"), F2: (9 >= 127)}
        _ = v83
    } else {
        v84 := GenS11{F0: 1, F1: true, F2: genFn0()}
        if true {
            v3 = v3
            if false {
                v85 := GenS12{F0: false, F1: v3}
                _ = v85
            } else {
                fmt.Println(v42, v3)
                v86 := GenS13{F0: 255}
                if (128 >= 0) {
                    v3 = v3
                    if (false || true) {
                        v87 := GenS14{F0: 0}
                        _ = v87
                    } else {
                        var v88 float32 = -2.25
                        _ = v88
                    }
                    if true {
                        for fi89 := 0; fi89 < 16; fi89++ {
                            for fi90 := 0; fi90 < 7; fi90++ {
                                if (0 <= 4294967295) {
                                    var v91 uint16 = 65535
                                    if len(v2) > 0 {
                                        v2[0] = (((v3 + "foo bar") + "") + v3)
                                    }
                                    var v92 int16 = v42
                                    _ = v91
                                    _ = v92
                                } else {
                                    var v93 int = (fi89 ^ (fi89 - 32767))
                                    v94 := GenS15{F0: 4294967295}
                                    _ = v93
                                    _ = v94
                                }
                                if (((v42 & -943) <= (v42 << ((0) & 63))) && true) {
                                    if len(v2) > 0 {
                                        v2[0] = "a"
                                    }
                                }
                                switch (v3 + v3) {
                                case "\n\t":
                                    v95 := GenS15{F0: 85}
                                    _ = v95
                                case "foo bar":
                                    fmt.Println(fi90, v3, fi90)
                                case "a":
                                    if true {
                                        var v96 int16 = (-(v42))
                                        switch fi89 {
                                        case 0:
                                            var v97 int8 = -1
                                            _ = v97
                                        case 9223372036854775807:
                                            v42 = v96
                                        }
                                        if (785 == 444) {
                                            v42 = (^(v42))
                                        }
                                        _ = v96
                                    }
                                    v3 = ""
                                    var v98 uint = 42
                                    _ = v98
                                default:
                                    fmt.Println(fi90)
                                }
                                _ = fi90
                            }
                            v99 := GenS13{F0: 65535}
                            v100 := GenS16{F0: 42}
                            _ = fi89
                            _ = v100
                            _ = v99
                        }
                    }
                } else {
                    switch v3 {
                    case "hello":
                        fmt.Println(v42, v3)
                    default:
                        for fi101 := 0; fi101 < 16; fi101++ {
                            switch (((v3 + ("a" + v3)) + "hello") + v3) {
                            case "a":
                                for fi102 := 0; fi102 < 14; fi102++ {
                                    var v103 uint32 = 109
                                    var v104 uint64 = 1
                                    _ = fi102
                                    _ = v103
                                    _ = v104
                                }
                            case "hello":
                                var v105 float64 = (float64(3.14159) * (float64(0.1) - (1.5 / 0.1)))
                                _ = v105
                            }
                            _ = fi101
                        }
                    }
                }
                _ = v86
            }
            var v106 bool = true
            _ = v106
        }
        v42 -= v42
        _ = v84
    }
    v107 := GenS17{F0: -242, F1: 2147483648, F2: 100}
    switch "" {
    case "\n\t":
        var v108 float32 = -2.25
        _ = v108
    case "a":
        if false {
            switch "\n\t" {
            case "a":
                switch 32767 {
                case 32767:
                    switch v3 {
                    case "":
                        v3 = (((("" + v3) + v3) + v3) + v3)
                    default:
                        var v109 float32 = 0.1
                        _ = v109
                    }
                case -308:
                    var v110 string = v3
                    _ = v110
                case -507:
                    v111 := GenS12{F0: true, F1: ""}
                    _ = v111
                }
            default:
                v3 = ((v3 + "\n\t") + v3)
            }
            var v112 int8 = 0
            switch "foo bar" {
            case "\n\t":
                var v113 uint = 42
                _ = v113
            }
            _ = v112
        }
    case "hello":
        var v114 int8 = 100
        _ = v114
    }
    delete(v1, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v71
}
