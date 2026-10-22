package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 string; F2 int16 }
type GenS7 struct { F0 string; F1 int16 }
type GenS8 struct { F0 uint32 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"foo bar"}
    var v3 uint16 = 1
    var v4 int64 = 42
    v5 := make(map[uint32]string)
    delete(v1, ("\n\t" + "foo bar"))
    if false {
        v6 := GenS1{F0: int8(v4)}
        v2 = append(v2, "")
        if (true || false) {
            v3 = v3
            v3 = v3
        } else {
            v4 += v4
        }
        _ = v6
    } else {
        fmt.Println(v3, v3)
    }
    v4 -= v4
    for fi7 := 0; fi7 < 17; fi7++ {
        if false {
            _ = len(v5)
        }
        var v8 uint32 = 0
        var v9 bool = (!true)
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        switch ("" + "a") {
        case "\n\t":
            fmt.Println(v3, v4)
        default:
            v11 := GenS1{F0: 0}
            _ = v11
        }
        if (2 < 2) {
            v3 += (v3 / (((v3 << ((0) & 63))) | 1))
            var v12 string = ("hello" + ("a" + (("\n\t" + "a") + "\n\t")))
            _ = v12
        }
        _ = fi10
    }
    switch 302 {
    case 606:
        v4 = v4
    case -349:
        v1[""] = ((1.5 / 1e10) != 0.0)
    }
    var v13 uint32 = 970
    if ((((v4 << ((563) & 63)) >= v4) || false) || (1.5 <= float32((1e10 - -2.25)))) {
        v4 = (v4 & -9223372036854775808)
    } else {
        v14 := GenS2{F0: (v13 * v13), F1: 2, F2: v3}
        var v15 uint8 = 0
        if true {
            if false {
                if false {
                    v16 := GenS3{F0: 128}
                    var v17 float32 = 0.0
                    fmt.Println(v17)
                    _ = v16
                    _ = v17
                } else {
                    v18 := v5[0]
                    _ = v18
                }
            }
        }
        _ = v14
        _ = v15
    }
    if (!(42 >= 1)) {
        v4 = 1
        v19 := GenS4{F0: (v4 & (v4 % ((int64(v3)) | 1)))}
        _ = v19
    }
    if (123 <= 255) {
        if (!true) {
            var v20 bool = (!false)
            _ = v20
        } else {
            if true {
                for fi21 := 0; fi21 < 17; fi21++ {
                    var v22 int = (^(fi21))
                    if ((false || (v13 != (v13 * v13))) && false) {
                        fmt.Println(v4, v13)
                        switch (("\n\t" + "hello") + ("a" + "foo bar")) {
                        case "hello":
                            if (0 >= -901) {
                                var v23 int32 = 0
                                delete(v1, "foo bar")
                                v24 := GenS3{F0: (^(fi21))}
                                _ = v23
                                _ = v24
                            } else {
                                switch "" {
                                case "hello":
                                    fmt.Println(v22, v4, fi21)
                                case "\n\t":
                                    for fi25 := 0; fi25 < 6; fi25++ {
                                        var v26 int64 = (-(v4))
                                        v27 := GenS3{F0: fi25}
                                        var v28 uint32 = 2
                                        _ = fi25
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                case "":
                                    var v29 int32 = 2147483647
                                    _ = v29
                                }
                                switch fi21 {
                                case 127:
                                    for fi30 := 0; fi30 < 6; fi30++ {
                                        var v31 int64 = 829
                                        _ = fi30
                                        _ = v31
                                    }
                                default:
                                    v32 := GenS2{F0: v13, F1: 1, F2: 42}
                                    _ = v32
                                }
                            }
                        case "a":
                            for fi33 := 0; fi33 < 20; fi33++ {
                                v22 = fi21
                                _ = fi33
                            }
                        case "foo bar":
                            if true {
                                v3 += (v3 % ((v3) | 1))
                                var v34 uint16 = 544
                                _ = v34
                            } else {
                                v35 := GenS5{F0: v13}
                                _ = v35
                            }
                        }
                        if false {
                            for fi36 := 0; fi36 < 7; fi36++ {
                                fmt.Println(fi21)
                                _ = fi36
                            }
                        } else {
                            for fi37 := 0; fi37 < 15; fi37++ {
                                switch 42 {
                                case 1:
                                    if (32768 <= fi21) {
                                        var v38 bool = ((!(!(false || true))) && false)
                                        v3 = 65535
                                        if true {
                                            var v39 int32 = -2147483648
                                            _ = v39
                                        } else {
                                            fmt.Println(fi37, v3, v4)
                                            var v40 uint = 1
                                            if v38 {
                                                v4 = (^(v4))
                                                var v41 bool = v38
                                                _ = v41
                                            }
                                            _ = v40
                                        }
                                        _ = v38
                                    }
                                default:
                                    var v42 uint32 = (^(v13))
                                    _ = v42
                                }
                                _ = fi37
                            }
                            v22 = (fi21 / (((fi21 >> ((765) & 63))) | 1))
                        }
                    }
                    v43 := GenS2{F0: v13, F1: 478, F2: (v3 + v3)}
                    _ = fi21
                    _ = v22
                    _ = v43
                }
                if (v3 >= (^(v3))) {
                    v44 := GenS3{F0: int(v3)}
                    if true {
                        fmt.Println(v13)
                        if (v4 > v4) {
                            switch "foo bar" {
                            case "\n\t":
                                for fi45 := 0; fi45 < 11; fi45++ {
                                    v1[("foo bar" + "a")] = true
                                    if false {
                                        fmt.Println(v13, v4)
                                        switch "foo bar" {
                                        case "":
                                            if len(v2) > 0 {
                                                v2[0] = "hello"
                                            }
                                        default:
                                            v2 = append(v2, ("hello" + "foo bar"))
                                        }
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v2 = append(v2, "foo bar")
                                    }
                                    fmt.Println(v13, v4)
                                    _ = fi45
                                }
                            case "hello":
                                var v46 int8 = -128
                                _ = v46
                            }
                        } else {
                            _ = len(v1)
                            var v47 int8 = 96
                            switch -219 {
                            case 128:
                                if ((false && (128 < 1)) && (18446744073709551615 < 24)) {
                                    for fi48 := 0; fi48 < 2; fi48++ {
                                        var v49 uint8 = 42
                                        var v50 bool = (1 != v13)
                                        _ = fi48
                                        _ = v49
                                        _ = v50
                                    }
                                }
                            case 4294967295:
                                v51 := GenS3{F0: 4294967295}
                                _ = v51
                            case -1:
                                for fi52 := 0; fi52 < 10; fi52++ {
                                    var v53 uint16 = 2
                                    for fi54 := 0; fi54 < 19; fi54++ {
                                        var v55 string = "a"
                                        _ = fi54
                                        _ = v55
                                    }
                                    var v56 int = (fi52 >> ((170) & 63))
                                    _ = fi52
                                    _ = v53
                                    _ = v56
                                }
                            }
                            _ = v47
                        }
                    } else {
                        if false {
                            for fi57 := 0; fi57 < 12; fi57++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v13 = v13
                                _ = fi57
                            }
                        } else {
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                            v58 := GenS5{F0: v13}
                            _ = v58
                        }
                        var v59 int8 = 127
                        _ = v59
                    }
                    for fi60 := 0; fi60 < 12; fi60++ {
                        var v61 uint8 = 128
                        _ = fi60
                        _ = v61
                    }
                    _ = v44
                }
            } else {
                for fi62 := 0; fi62 < 8; fi62++ {
                    var v63 int16 = -32768
                    for fi64 := 0; fi64 < 3; fi64++ {
                        var v65 uint32 = (v13 >> ((1) & 63))
                        v4 = (v4 ^ (^(v4)))
                        _ = fi64
                        _ = v65
                    }
                    _ = fi62
                    _ = v63
                }
            }
            if false {
                var v66 uint = 18446744073709551615
                var v67 int16 = 1
                v68 := v1[(("a" + "\n\t") + "hello")]
                _ = v66
                _ = v67
                _ = v68
            } else {
                switch "a" {
                case "hello":
                    fmt.Println(v3, v13, v4)
                case "foo bar":
                    switch "\n\t" {
                    case "hello":
                        var v69 string = "\n\t"
                        _ = v69
                    case "\n\t":
                        v70 := GenS6{F0: (v3 - 0), F1: "a", F2: 42}
                        _ = v70
                    case "":
                        var v71 uint8 = 2
                        _ = v71
                    default:
                        if (v13 < (v13 * (v13 / ((0) | 1)))) {
                            v3 = (v3 << ((32) & 63))
                        } else {
                            if len(v2) > 0 {
                                v2[0] = ((((("foo bar" + "a") + ("hello" + "foo bar")) + "foo bar") + "foo bar") + (("a" + (("a" + "hello") + "")) + "foo bar"))
                            }
                            fmt.Println(v13)
                            v4 += (v4 >> ((255) & 63))
                        }
                    }
                case "\n\t":
                    for fi72 := 0; fi72 < 15; fi72++ {
                        var v73 bool = (!(0.0 <= 3.14159))
                        v74 := GenS3{F0: fi72}
                        for fi75 := 0; fi75 < 15; fi75++ {
                            v4 += v4
                            fmt.Println(fi75, fi72, v3)
                            var v76 uint64 = 0
                            _ = fi75
                            _ = v76
                        }
                        _ = fi72
                        _ = v73
                        _ = v74
                    }
                }
            }
            var v77 uint32 = (v13 | 42)
            _ = v77
        }
        v78 := GenS7{F0: "", F1: -1}
        fmt.Println(v4, v4, v3)
        _ = v78
    } else {
        var v79 float32 = 3.14159
        v79 += v79
        switch ("\n\t" + "hello") {
        case "\n\t":
            v80 := GenS2{F0: (v13 ^ v13), F1: 2, F2: v3}
            _ = v80
        }
        _ = v79
    }
    if (!true) {
        if false {
            if true {
                fmt.Println(v4, v4, v13)
                v5[(v13 + v13)] = ("foo bar" + "a")
                var v81 uint32 = v13
                _ = v81
            }
        }
    }
    var v82 uint8 = 246
    var v83 string = "hello"
    for fi84 := 0; fi84 < 11; fi84++ {
        var v85 int32 = 0
        var v86 float64 = 1.5
        v83 = (("foo bar" + ("\n\t" + v83)) + ("a" + ("foo bar" + "foo bar")))
        _ = fi84
        _ = v85
        _ = v86
    }
    if (v82 != 0) {
        var v87 int32 = -1
        _ = v87
    }
    if (!false) {
        if true {
            for fi88 := 0; fi88 < 15; fi88++ {
                if true {
                    delete(v1, ("hello" + ("" + v83)))
                } else {
                    for fi89 := 0; fi89 < 4; fi89++ {
                        var v90 uint16 = (^(v3))
                        v91 := GenS7{F0: "hello", F1: 1}
                        v92 := GenS7{F0: v83, F1: 1}
                        v3 += v3
                        _ = fi89
                        _ = v90
                        _ = v91
                        _ = v92
                    }
                }
                _ = len(v2)
                if (-2.25 >= (float32(((0.1 + 1e10) - 1.5)) - 3.14159)) {
                    fmt.Println(v3, v82, v83)
                } else {
                    var v93 uint32 = uint32(v82)
                    var v94 uint = 18446744073709551615
                    _ = v93
                    _ = v94
                }
                _ = fi88
            }
        } else {
            for fi95 := 0; fi95 < 6; fi95++ {
                var v96 int = fi95
                _ = fi95
                _ = v96
            }
        }
    } else {
        for fi97 := 0; fi97 < 17; fi97++ {
            fmt.Println(v83, v82, v82)
            v98 := GenS5{F0: v13}
            delete(v1, v83)
            _ = fi97
            _ = v98
        }
        if (2147483647 != 127) {
            v4 = v4
            v99 := GenS8{F0: 2}
            _ = v99
        } else {
            _ = len(v2)
        }
    }
    switch "" {
    case "":
        v13 += v13
    case "foo bar":
        for fi100 := 0; fi100 < 3; fi100++ {
            var v101 uint8 = (v82 >> ((2) & 63))
            _ = fi100
            _ = v101
        }
    case "a":
        v3 += (v3 / ((v3) | 1))
    default:
        switch -880 {
        case -820:
            var v102 int16 = 32767
            _ = v102
        case 32768:
            if (!(-477 > 833)) {
                for fi103 := 0; fi103 < 18; fi103++ {
                    switch (v83 + ("a" + ("foo bar" + "hello"))) {
                    case "\n\t":
                        v82 = v82
                    }
                    _ = fi103
                }
                for fi104 := 0; fi104 < 4; fi104++ {
                    var v105 int32 = -1
                    _ = fi104
                    _ = v105
                }
            }
        case -157:
            fmt.Println(v13)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v82
    _ = v83
}
