package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 uint16; F2 string }
type GenS8 struct { F0 int; F1 uint64; F2 int32 }
type GenS9 struct { F0 string; F1 int64; F2 uint16 }
type GenS10 struct { F0 int64; F1 uint32 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 bool }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{}
    var v3 int16 = 0
    var v4 float64 = -2.25
    v5 := []uint32{2, 740}
    var v6 uint16 = 399
    delete(v1, 709)
    var v7 bool = false
    if false {
        if (v4 <= v4) {
            if len(v2) > 0 {
                v2[0] = (("a" + "\n\t") + "\n\t")
            }
            var v8 uint16 = 387
            _ = v8
        } else {
            v1[4294967295] = v7
            v6 -= v6
            var v9 uint = 255
            _ = v9
        }
        v10 := v1[176]
        if ("hello" >= ("a" + "hello")) {
            switch "hello" {
            case "":
                if true {
                    var v11 int8 = -29
                    switch "hello" {
                    case "a":
                        for fi12 := 0; fi12 < 18; fi12++ {
                            v6 = v6
                            v3 += v3
                            _ = fi12
                        }
                    default:
                        fmt.Println(v11, v4, v11)
                    }
                    _ = v11
                }
            case "hello":
                switch "foo bar" {
                case "hello":
                    v13 := GenS1{F0: 536, F1: -1}
                    _ = v13
                default:
                    fmt.Println(v3, v10)
                }
            default:
                switch "" {
                case "a":
                    var v14 int16 = (v3 & (^(v3)))
                    _ = v14
                case "foo bar":
                    _ = len(v2)
                }
            }
            switch "a" {
            case "hello":
                delete(v1, 896)
            case "foo bar":
                _ = len(v1)
            case "\n\t":
                if false {
                    switch 865 {
                    case 42:
                        var v15 bool = v7
                        _ = v15
                    case 1:
                        v10 = v10
                    default:
                        var v16 bool = v7
                        _ = v16
                    }
                    var v17 int64 = 9223372036854775807
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 9; fi18++ {
                        for fi19 := 0; fi19 < 16; fi19++ {
                            v20 := GenS1{F0: 1, F1: 127}
                            switch -2147483648 {
                            case -752:
                                _ = len(v2)
                            case 32767:
                                v21 := GenS2{F0: 881, F1: v10}
                                _ = v21
                            default:
                                v22 := GenS3{F0: true, F1: 124}
                                _ = v22
                            }
                            _ = fi19
                            _ = v20
                        }
                        var v23 uint16 = 185
                        v24 := GenS1{F0: 0, F1: 1}
                        _ = fi18
                        _ = v23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 14; fi25++ {
                        v7 = v7
                        for fi26 := 0; fi26 < 5; fi26++ {
                            v27 := GenS2{F0: -246, F1: v10}
                            v4 = 3.14159
                            _ = fi26
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 14; fi28++ {
                            if v7 {
                                switch (("hello" + (("\n\t" + "a") + ("\n\t" + "foo bar"))) + "") {
                                case "":
                                    var v29 uint16 = v6
                                    _ = v29
                                case "\n\t":
                                    var v30 int8 = 1
                                    _ = v30
                                case "a":
                                    if (!((fi25 % (((fi28 * fi28)) | 1)) >= (fi25 | fi28))) {
                                        v31 := GenS4{F0: 127, F1: 32767}
                                        v32 := GenS5{F0: 475}
                                        _ = v31
                                        _ = v32
                                    } else {
                                        var v33 int64 = 9223372036854775807
                                        if v10 {
                                            if true {
                                                v7 = v7
                                                if v10 {
                                                    v34 := GenS4{F0: 0, F1: 1}
                                                    _ = v34
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = 443
                                                }
                                                fmt.Println(v10)
                                                var v35 uint64 = 4294967295
                                                fmt.Println(v3, fi25)
                                                _ = v35
                                            } else {
                                                var v36 int8 = 1
                                                _ = v36
                                            }
                                            var v37 string = ""
                                            switch (v37 + "hello") {
                                            case "a":
                                                _ = len(v5)
                                                var v38 uint32 = 4294967295
                                                v39 := GenS6{F0: 1}
                                                _ = v38
                                                _ = v39
                                            }
                                            _ = v37
                                        } else {
                                            fmt.Println(fi25, fi25, v4)
                                            fmt.Println(v6, v4, v7)
                                            if (2147483647 == 65535) {
                                                if (!false) {
                                                    var v40 float32 = 3.14159
                                                    v41 := GenS4{F0: 1, F1: v3}
                                                    _ = v40
                                                    _ = v41
                                                }
                                                fmt.Println(v4, v10, v10)
                                                var v42 int = fi25
                                                v43 := GenS5{F0: 18446744073709551615}
                                                v44 := GenS2{F0: 2147483647, F1: v7}
                                                _ = v42
                                                _ = v43
                                                _ = v44
                                            } else {
                                                var v45 uint32 = 453
                                                v46 := GenS7{F0: "\n\t", F1: (v6 ^ uint16(v45)), F2: "foo bar"}
                                                if (!(901 != v6)) {
                                                    v3 = -458
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    var v47 int = (fi25 << ((668) & 63))
                                                    _ = v47
                                                } else {
                                                    v48 := GenS8{F0: fi25, F1: 42, F2: -235}
                                                    v6 = 881
                                                    var v49 float32 = 0.0
                                                    _ = v48
                                                    _ = v49
                                                }
                                                fmt.Println(v10, v33)
                                                v50 := GenS9{F0: "hello", F1: (v33 - (v33 - v33)), F2: (v6 % ((v6) | 1))}
                                                _ = v45
                                                _ = v46
                                                _ = v50
                                            }
                                        }
                                        _ = v33
                                    }
                                }
                                if v10 {
                                    var v51 bool = v10
                                    _ = v51
                                }
                            } else {
                                v7 = v10
                                delete(v1, 2147483648)
                            }
                            switch (fi25 >> ((1) & 63)) {
                            case 1:
                                switch fi25 {
                                case 4294967295:
                                    v4 += v4
                                case 32768:
                                    var v52 int32 = -1
                                    _ = v52
                                case 42:
                                    var v53 int32 = 42
                                    _ = v53
                                default:
                                    fmt.Println(v4, v7)
                                }
                            case 4294967295:
                                v7 = (!((false && v10) || (100 != 0)))
                            case 127:
                                var v54 uint16 = (^(v6))
                                _ = v54
                            default:
                                var v55 float32 = 0.1
                                _ = v55
                            }
                            var v56 uint16 = v6
                            v56 -= (v6 * 42)
                            _ = fi28
                            _ = v56
                        }
                        _ = fi25
                    }
                    if false {
                        v10 = (2147483647 >= -2147483648)
                        switch "" {
                        case "":
                            fmt.Println(v3)
                        case "foo bar":
                            v6 = v6
                        }
                        var v57 int = 9223372036854775807
                        _ = v57
                    } else {
                        var v58 string = (("hello" + "hello") + "")
                        _ = v58
                    }
                }
            default:
                if v7 {
                    v10 = v10
                    for fi59 := 0; fi59 < 10; fi59++ {
                        v6 = (v6 << ((65535) & 63))
                        _ = fi59
                    }
                } else {
                    var v60 int8 = -128
                    fmt.Println(v4)
                    if v7 {
                        switch ("hello" + ("" + ("" + ("a" + "foo bar")))) {
                        case "":
                            for fi61 := 0; fi61 < 13; fi61++ {
                                v3 -= -32768
                                _ = len(v1)
                                _ = fi61
                            }
                        case "a":
                            for fi62 := 0; fi62 < 3; fi62++ {
                                switch -897 {
                                case 4294967295:
                                    fmt.Println(v60, fi62)
                                case 128:
                                    for fi63 := 0; fi63 < 20; fi63++ {
                                        v4 = (((v4 / 0.1) * (v4 + 1e10)) + (1e10 - float64(0.0)))
                                        switch ("a" + "") {
                                        case "":
                                            var v64 string = "foo bar"
                                            _ = v64
                                        }
                                        v6 = v6
                                        _ = fi63
                                    }
                                case 2147483647:
                                    fmt.Println(v3, v6)
                                }
                                _ = fi62
                            }
                        default:
                            var v65 uint8 = 0
                            _ = v65
                        }
                        fmt.Println(v60)
                        var v66 int16 = -1
                        _ = v66
                    }
                    _ = v60
                }
            }
            if v10 {
                if (v10 || (!v7)) {
                    var v67 bool = (65535 > 0)
                    v68 := GenS10{F0: -476, F1: 2147483648}
                    if true {
                        for fi69 := 0; fi69 < 16; fi69++ {
                            v70 := GenS7{F0: "", F1: (^(v6)), F2: ("hello" + "\n\t")}
                            var v71 uint16 = v6
                            for fi72 := 0; fi72 < 3; fi72++ {
                                var v73 uint8 = 0
                                _ = fi72
                                _ = v73
                            }
                            _ = fi69
                            _ = v70
                            _ = v71
                        }
                        v2 = append(v2, "a")
                    }
                    _ = v67
                    _ = v68
                }
                var v74 float64 = (v4 - 3.14159)
                for fi75 := 0; fi75 < 8; fi75++ {
                    var v76 uint64 = 18446744073709551615
                    _ = fi75
                    _ = v76
                }
                _ = v74
            }
        } else {
            v10 = (-298 < 42)
            v10 = v10
            v77 := GenS5{F0: 42}
            _ = v77
        }
        _ = v10
    }
    if (!("\n\t" != ("\n\t" + ""))) {
        if v7 {
            v3 = (v3 + 42)
            v7 = (!(1 < 0))
        } else {
            v6 = v6
            for fi78 := 0; fi78 < 15; fi78++ {
                switch ("a" + ("foo bar" + (((("\n\t" + ((("\n\t" + "hello") + (((("hello" + ("" + ("foo bar" + "\n\t"))) + ("\n\t" + "\n\t")) + "\n\t") + "a")) + ("foo bar" + ("hello" + "foo bar")))) + "foo bar") + ("\n\t" + "\n\t")) + ("" + "")))) {
                case "hello":
                    for fi79 := 0; fi79 < 13; fi79++ {
                        _ = len(v1)
                        fmt.Println(v4, v3)
                        v2 = append(v2, "")
                        v80 := GenS11{F0: 42}
                        _ = fi79
                        _ = v80
                    }
                case "a":
                    switch (fi78 << ((631) & 63)) {
                    case 42:
                        var v81 int64 = 9223372036854775807
                        _ = v81
                    case 747:
                        v7 = (v6 == v6)
                    default:
                        var v82 uint64 = 809
                        _ = v82
                    }
                case "\n\t":
                    for fi83 := 0; fi83 < 4; fi83++ {
                        switch -651 {
                        case -9223372036854775808:
                            v6 = (v6 % ((v6) | 1))
                        default:
                            v7 = (float32(-2.25) >= 0.0)
                        }
                        v7 = (v7 && v7)
                        _ = fi83
                    }
                default:
                    v4 += 1e10
                }
                v6 = v6
                _ = fi78
            }
        }
        v1[0] = true
    }
    for fi84 := 0; fi84 < 7; fi84++ {
        switch "\n\t" {
        case "foo bar":
            var v85 uint32 = 266
            _ = v85
        case "":
            v4 += ((v4 + (v4 + v4)) * ((v4 / v4) + 3.14159))
        }
        var v86 int = -9223372036854775808
        _ = fi84
        _ = v86
    }
    var v87 bool = true
    var v88 string = ("hello" + "a")
    for fi89 := 0; fi89 < 20; fi89++ {
        v90 := GenS9{F0: v88, F1: 42, F2: (v6 << ((1) & 63))}
        var v91 uint16 = (v6 / ((1) | 1))
        _ = len(v5)
        _ = fi89
        _ = v90
        _ = v91
    }
    delete(v1, 2147483648)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3)
    switch 127 {
    case -1:
        for fi92 := 0; fi92 < 5; fi92++ {
            var v93 float32 = (0.1 - float32(v4))
            v94 := GenS5{F0: 4294967295}
            _ = fi92
            _ = v93
            _ = v94
        }
    }
    var v95 uint = 167
    var v96 int16 = 387
    if (526 == 0) {
        for fi97 := 0; fi97 < 2; fi97++ {
            v98 := GenS12{F0: int8(v3)}
            v99 := GenS13{F0: true}
            _ = fi97
            _ = v98
            _ = v99
        }
        switch 1 {
        case -84:
            v100 := v1[346]
            _ = v100
        }
    } else {
        v1[859] = v87
        if (!(!v7)) {
            fmt.Println(v6, v7)
        }
    }
    if (!v87) {
        var v101 float64 = ((3.14159 - ((v4 * (0.1 / (v4 / (v4 - ((float64((((float32(((v4 - (0.0 + 3.14159)) * float64(3.14159))) - (0.0 - ((float32(1.5) / 1e10) + 0.1))) / 1e10) * 1e10)) / (v4 + (0.1 - 1e10))) - 0.1))))) / v4)) + v4)
        _ = v101
    } else {
        v96 -= v3
        if (((v6 != (v6 & 42)) || ((v87 && v87) || (v6 != 330))) || ((v95 | v95) >= 255)) {
            var v102 uint64 = 18446744073709551615
            v3 = v3
            _ = v102
        } else {
            v3 += 1000
            if true {
                v4 -= float64((1.5 - 0.0))
                v5 = append(v5, 2)
            } else {
                var v103 uint8 = 255
                if v7 {
                    v104 := GenS5{F0: (v95 + (v95 + v95))}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v105 int16 = (v3 / ((v96) | 1))
                    _ = v104
                    _ = v105
                } else {
                    var v106 bool = (!true)
                    _ = v106
                }
                _ = v103
            }
            v107 := GenS12{F0: 100}
            _ = v107
        }
        v7 = v7
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v87
    _ = v88
    _ = v95
    _ = v96
}
