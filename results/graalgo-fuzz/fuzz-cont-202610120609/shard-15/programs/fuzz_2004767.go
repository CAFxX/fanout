package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint8; F1 string; F2 int32 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 1
    v3 := make(map[uint32]int)
    var v4 uint64 = 1
    var v5 int64 = 42
    v6 := []int{-1, 4294967295, -752}
    var v7 bool = false
    if v7 {
        if ((v2 | v2) > 1) {
            for fi8 := 0; fi8 < 10; fi8++ {
                fmt.Println(v4)
                _ = fi8
            }
        } else {
            if v7 {
                if ((9 >= -1) || v7) {
                    var v9 string = ("hello" + "foo bar")
                    for fi10 := 0; fi10 < 4; fi10++ {
                        v7 = ((true && v7) && (0.1 == 0.1))
                        var v11 uint16 = 32767
                        v12 := GenS1{F0: 484, F1: v7}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = v9
                } else {
                    v5 -= v5
                    var v13 string = "hello"
                    _ = v13
                }
                if true {
                    v4 = (v4 - 18446744073709551615)
                    for fi14 := 0; fi14 < 2; fi14++ {
                        v7 = (1e10 >= (0.1 / (float32(0.1) / 1e10)))
                        v1[65535] = (fi14 | 1)
                        _ = fi14
                    }
                    for fi15 := 0; fi15 < 5; fi15++ {
                        switch "\n\t" {
                        case "":
                            if v7 {
                                if (int32(fi15) < 42) {
                                    v16 := GenS1{F0: 2147483647, F1: true}
                                    _ = v16
                                } else {
                                    v4 += v4
                                    v17 := GenS2{F0: 42, F1: ("" + "foo bar"), F2: 42}
                                    _ = v17
                                }
                                v7 = ((2 < (v2 + 65535)) || false)
                                fmt.Println(v7, v2)
                            } else {
                                if false {
                                    var v18 uint8 = 255
                                    switch "hello" {
                                    case "foo bar":
                                        v19 := GenS3{F0: 2147483647, F1: "a"}
                                        _ = v19
                                    case "\n\t":
                                        v20 := v3[833]
                                        _ = v20
                                    case "hello":
                                        var v21 uint16 = 32767
                                        _ = v21
                                    }
                                    _ = v18
                                } else {
                                    for fi22 := 0; fi22 < 16; fi22++ {
                                        var v23 int16 = 1000
                                        fmt.Println(v7)
                                        _ = fi22
                                        _ = v23
                                    }
                                }
                                v24 := GenS2{F0: 96, F1: (("hello" + "") + "hello"), F2: -961}
                                _ = v24
                            }
                        case "hello":
                            if v7 {
                                if v7 {
                                    switch "\n\t" {
                                    case "\n\t":
                                        v2 += 18446744073709551615
                                    }
                                    for fi25 := 0; fi25 < 19; fi25++ {
                                        _ = len(v1)
                                        if (1.5 == 1.5) {
                                            v7 = v7
                                            v7 = (0.0 >= -2.25)
                                        } else {
                                            var v26 float32 = 0.1
                                            v3[493] = (fi25 + (^(fi15)))
                                            _ = v26
                                        }
                                        _ = fi25
                                    }
                                }
                                v6 = append(v6, 259)
                                var v27 int16 = 536
                                _ = v27
                            } else {
                                v2 += (^(v2))
                                fmt.Println(v7, v7)
                                fmt.Println(fi15)
                            }
                        case "foo bar":
                            switch fi15 {
                            case 127:
                                v4 = 0
                            case 42:
                                for fi28 := 0; fi28 < 19; fi28++ {
                                    v2 = (v2 * (^(v2)))
                                    fmt.Println(v2)
                                    _ = fi28
                                }
                            default:
                                if len(v6) > 0 {
                                    v6[0] = fi15
                                }
                            }
                        }
                        v29 := GenS3{F0: 42, F1: "foo bar"}
                        if ((v4 % (((v4 << ((v2) & 63))) | 1)) > (^(v4))) {
                            if (!v7) {
                                for fi30 := 0; fi30 < 18; fi30++ {
                                    var v31 float64 = (3.14159 + (-2.25 / 1e10))
                                    fmt.Println(v5)
                                    _ = fi30
                                    _ = v31
                                }
                                v32 := GenS4{F0: 127}
                                _ = v32
                            } else {
                                fmt.Println(v2)
                            }
                            if (!v7) {
                                var v33 int = fi15
                                delete(v1, 0)
                                _ = v33
                            }
                        } else {
                            v5 = (v5 << ((255) & 63))
                        }
                        _ = fi15
                        _ = v29
                    }
                } else {
                    var v34 string = "hello"
                    _ = v34
                }
            }
            v5 -= (v5 | v5)
            v6 = append(v6, 127)
        }
        v35 := GenS5{F0: 2147483647, F1: 750, F2: 1}
        v36 := GenS6{F0: 42, F1: 128}
        _ = v35
        _ = v36
    } else {
        v37 := GenS6{F0: 874, F1: -9223372036854775808}
        var v38 int32 = 111
        v39 := GenS6{F0: 42, F1: 0}
        _ = v37
        _ = v38
        _ = v39
    }
    if ((100 != 127) && v7) {
        delete(v1, 2147483648)
        for fi40 := 0; fi40 < 8; fi40++ {
            v6 = append(v6, 2147483647)
            var v41 float64 = float64(0.1)
            if v7 {
                v2 -= 39
                switch "" {
                case "\n\t":
                    v4 = v4
                case "a":
                    var v42 float64 = ((0.1 / v41) + v41)
                    _ = v42
                case "foo bar":
                    v5 -= v5
                default:
                    v2 = v2
                }
            } else {
                if true {
                    switch "hello" {
                    case "foo bar":
                        v4 -= v4
                    default:
                        var v43 string = ""
                        _ = v43
                    }
                    v3[42] = -1
                    v5 -= (v5 / ((v5) | 1))
                }
                for fi44 := 0; fi44 < 2; fi44++ {
                    var v45 bool = true
                    _ = fi44
                    _ = v45
                }
                v4 = v4
            }
            _ = fi40
            _ = v41
        }
    }
    v2 = 2
    fmt.Println(v4, v2, v2)
    if v7 {
        var v46 int8 = -110
        _ = v46
    } else {
        for fi47 := 0; fi47 < 11; fi47++ {
            if v7 {
                v7 = (!(("hello" + "foo bar") < ""))
            }
            _ = fi47
        }
    }
    if ((-1 >= 475) || (250 == 0)) {
        v48 := GenS3{F0: -2147483648, F1: "foo bar"}
        var v49 uint = (v2 % ((v2) | 1))
        if v7 {
            var v50 uint8 = 127
            v49 -= 445
            switch ("\n\t" + "\n\t") {
            case "":
                if (("a" + "\n\t") != "\n\t") {
                    var v51 float64 = 3.14159
                    if (true || true) {
                        _ = len(v1)
                        v52 := GenS7{F0: 42, F1: 42}
                        _ = v52
                    }
                    if v7 {
                        var v53 int64 = v5
                        switch "a" {
                        case "hello":
                            var v54 int32 = -983
                            _ = v54
                        case "\n\t":
                            switch 0 {
                            case 2147483647:
                                for fi55 := 0; fi55 < 16; fi55++ {
                                    if v7 {
                                        if v7 {
                                            var v56 uint = (v49 | (^(v49)))
                                            switch "" {
                                            case "hello":
                                                for fi57 := 0; fi57 < 12; fi57++ {
                                                    v58 := v1[571]
                                                    _ = fi57
                                                    _ = v58
                                                }
                                            case "\n\t":
                                                fmt.Println(v50)
                                            case "foo bar":
                                                var v59 uint64 = (v4 * (v4 - v4))
                                                _ = v59
                                            default:
                                                v56 = (v2 - (v49 - v2))
                                            }
                                            for fi60 := 0; fi60 < 17; fi60++ {
                                                fmt.Println(v49, v53)
                                                switch fi55 {
                                                case -9223372036854775808:
                                                    v7 = v7
                                                case 595:
                                                    v61 := GenS6{F0: -972, F1: (fi55 / (((fi55 ^ -2147483648)) | 1))}
                                                    _ = v61
                                                }
                                                v50 = v50
                                                _ = fi60
                                            }
                                            _ = v56
                                        }
                                        if v7 {
                                            v62 := GenS4{F0: v50}
                                            if v7 {
                                                v63 := GenS5{F0: 1, F1: 2147483647, F2: 42}
                                                v7 = false
                                                _ = v63
                                            }
                                            _ = v62
                                        } else {
                                            for fi64 := 0; fi64 < 14; fi64++ {
                                                if (32767 > 585) {
                                                    var v65 int64 = (-(v5))
                                                    var v66 string = "foo bar"
                                                    switch (fi64 << ((v49) & 63)) {
                                                    case 0:
                                                        if v7 {
                                                            if v7 {
                                                                v67 := GenS7{F0: 0, F1: 2147483647}
                                                                var v68 bool = v7
                                                                v66 = v66
                                                                _ = v67
                                                                _ = v68
                                                            }
                                                        }
                                                    case -9223372036854775808:
                                                        v69 := GenS4{F0: (v50 & 155)}
                                                        _ = v69
                                                    case 127:
                                                        _ = len(v6)
                                                    }
                                                    _ = v65
                                                    _ = v66
                                                }
                                                v70 := GenS8{F0: 541, F1: 775}
                                                var v71 float64 = v51
                                                _ = fi64
                                                _ = v70
                                                _ = v71
                                            }
                                            v72 := GenS9{F0: 42}
                                            v1[859] = -428
                                            _ = v72
                                        }
                                        v73 := GenS9{F0: 0}
                                        _ = v73
                                    } else {
                                        v74 := GenS7{F0: 2147483648, F1: 0}
                                        _ = v74
                                    }
                                    var v75 uint32 = 0
                                    if len(v6) > 0 {
                                        v6[0] = fi55
                                    }
                                    _ = fi55
                                    _ = v75
                                }
                            default:
                                var v76 string = "a"
                                _ = v76
                            }
                        }
                        fmt.Println(v4, v7, v50)
                        _ = v53
                    }
                    _ = v51
                }
            }
            _ = v50
        } else {
            var v77 uint8 = 127
            _ = v77
        }
        _ = v48
        _ = v49
    } else {
        if (true && (!(872 >= 859))) {
            var v78 float64 = (0.0 - (1.5 / (float64(1.5) + 3.14159)))
            if true {
                switch "a" {
                case "a":
                    switch "hello" {
                    case "hello":
                        v5 += (v5 / ((v5) | 1))
                    case "":
                        v79 := GenS2{F0: 1, F1: "foo bar", F2: -2147483648}
                        _ = v79
                    default:
                        switch "\n\t" {
                        case "a":
                            v80 := v3[415]
                            _ = v80
                        }
                    }
                default:
                    if (v7 || v7) {
                        v2 += v2
                    } else {
                        v81 := GenS7{F0: 2, F1: 33}
                        var v82 int = -1
                        _ = v81
                        _ = v82
                    }
                }
                if (18446744073709551615 != v2) {
                    v83 := GenS9{F0: v5}
                    switch "foo bar" {
                    case "foo bar":
                        v2 = (v2 << ((v2) & 63))
                    }
                    _ = v83
                } else {
                    if false {
                        for fi84 := 0; fi84 < 9; fi84++ {
                            var v85 float64 = v78
                            for fi86 := 0; fi86 < 12; fi86++ {
                                var v87 int8 = -28
                                _ = fi86
                                _ = v87
                            }
                            _ = fi84
                            _ = v85
                        }
                    }
                    v5 = v5
                    switch ("foo bar" + "a") {
                    case "\n\t":
                        _ = len(v1)
                    case "a":
                        var v88 int16 = 583
                        _ = v88
                    default:
                        for fi89 := 0; fi89 < 4; fi89++ {
                            v2 += (v2 << (((^(v2))) & 63))
                            fmt.Println(v78, v7)
                            for fi90 := 0; fi90 < 10; fi90++ {
                                if (((v2 * v2) < (^(v2))) && v7) {
                                    fmt.Println(v2, v2, v2)
                                }
                                v2 -= (v2 << ((v2) & 63))
                                var v91 uint64 = 0
                                _ = fi90
                                _ = v91
                            }
                            _ = fi89
                        }
                    }
                }
            }
            _ = v78
        } else {
            if (v7 || false) {
                v6 = append(v6, 1)
                _ = len(v1)
                var v92 int64 = 9223372036854775807
                _ = v92
            }
            var v93 uint = v2
            _ = v93
        }
        fmt.Println(v7)
        if (!(!false)) {
            v7 = (1.5 != (0.1 * (1.5 / (3.14159 * 1.5))))
        } else {
            v2 -= (v2 - (v2 >> ((v2) & 63)))
            var v94 string = (((("a" + "a") + ("\n\t" + "foo bar")) + "a") + "\n\t")
            delete(v3, 1)
            _ = v94
        }
    }
    var v95 bool = true
    fmt.Println(v7, v5, v7)
    var v96 int16 = -1
    for fi97 := 0; fi97 < 15; fi97++ {
        v98 := GenS7{F0: 569, F1: 42}
        if v95 {
            v99 := GenS8{F0: 1, F1: 32767}
            v100 := GenS9{F0: (v5 % ((-870) | 1))}
            _ = v100
            _ = v99
        } else {
            v2 -= 2
        }
        var v101 float64 = -2.25
        _ = fi97
        _ = v101
        _ = v98
    }
    v96 = (-(v96))
    v102 := GenS4{F0: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v95
    _ = v96
}
