package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int32; F1 int16; F2 uint8 }
type GenS10 struct { F0 bool; F1 int16 }
type GenS11 struct { F0 int8; F1 int16 }

func main() {
    var v1 uint16 = 2
    var v2 uint8 = 2
    var v3 int64 = 0
    v4 := []uint32{65535, 4294967295, 2147483647}
    var v5 int = 940
    var v6 int8 = 42
    v4 = append(v4, 2)
    for fi7 := 0; fi7 < 17; fi7++ {
        v5 = v5
        if false {
            v2 = (v2 % ((128) | 1))
            if true {
                var v8 float64 = -2.25
                _ = v8
            } else {
                _ = len(v4)
                var v9 float32 = float32(3.14159)
                _ = v9
            }
            var v10 uint32 = 2147483647
            _ = v10
        } else {
            for fi11 := 0; fi11 < 13; fi11++ {
                v12 := GenS1{F0: 42, F1: v3}
                if (true && true) {
                    switch "a" {
                    case "":
                        v1 -= 888
                    case "\n\t":
                        v5 = (v5 >> ((4294967295) & 63))
                        v13 := GenS2{F0: 4294967295}
                        if true {
                            v4 = append(v4, 65535)
                        }
                        _ = v13
                    }
                    var v14 bool = ("foo bar" < ("hello" + ""))
                    fmt.Println(v1)
                    _ = v14
                } else {
                    var v15 uint16 = 827
                    _ = len(v4)
                    var v16 float64 = 0.1
                    v17 := GenS2{F0: 2}
                    _ = v15
                    _ = v16
                    _ = v17
                }
                _ = fi11
                _ = v12
            }
            var v18 uint32 = 580
            switch (("a" + "foo bar") + "\n\t") {
            case "foo bar":
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    v5 += (v5 / (((fi7 ^ v5)) | 1))
                case "hello":
                    switch fi7 {
                    case 32768:
                        _ = len(v4)
                    }
                case "foo bar":
                    var v19 int64 = 369
                    _ = v19
                default:
                    if false {
                        v18 = (v18 % ((297) | 1))
                        _ = len(v4)
                        if (9223372036854775807 != (v3 / (((v3 ^ 9223372036854775807)) | 1))) {
                            var v20 uint16 = v1
                            v21 := GenS2{F0: 4}
                            _ = v20
                            _ = v21
                        }
                    }
                }
            default:
                v22 := GenS2{F0: 4294967295}
                _ = v22
            }
            _ = v18
        }
        switch v5 {
        case 2147483647:
            var v23 int = 18
            _ = v23
        case 127:
            if (-1 >= 1) {
                var v24 uint64 = 1
                switch v5 {
                case 128:
                    if false {
                        for fi25 := 0; fi25 < 14; fi25++ {
                            v2 += v2
                            _ = fi25
                        }
                        for fi26 := 0; fi26 < 2; fi26++ {
                            switch "a" {
                            case "hello":
                                switch "hello" {
                                case "":
                                    var v27 string = "hello"
                                    _ = v27
                                default:
                                    v1 += 918
                                }
                            case "\n\t":
                                switch fi7 {
                                case -9223372036854775808:
                                    var v28 float64 = 0.1
                                    _ = v28
                                }
                                var v29 int64 = (v3 >> ((654) & 63))
                                _ = v29
                            case "":
                                v1 = v1
                                switch 0 {
                                case 4294967295:
                                    switch "hello" {
                                    case "":
                                        fmt.Println(v1)
                                    default:
                                        var v30 uint32 = 372
                                        if (!true) {
                                            fmt.Println(v6, v3)
                                        } else {
                                            if (true && true) {
                                                switch "\n\t" {
                                                case "hello":
                                                    var v31 uint64 = 42
                                                    _ = v31
                                                case "":
                                                    var v32 bool = (!false)
                                                    _ = v32
                                                default:
                                                    var v33 int16 = 1
                                                    _ = v33
                                                }
                                            } else {
                                                v34 := GenS2{F0: uint(v2)}
                                                _ = v34
                                            }
                                            _ = len(v4)
                                            v4 = append(v4, 42)
                                            v35 := GenS1{F0: 653, F1: v3}
                                            v36 := GenS1{F0: v30, F1: v3}
                                            _ = v35
                                            _ = v36
                                        }
                                        _ = v30
                                    }
                                default:
                                    if len(v4) > 0 {
                                        v4[0] = 854
                                    }
                                }
                                v37 := GenS3{F0: (!false)}
                                _ = v37
                            default:
                                var v38 uint8 = (v2 / ((127) | 1))
                                _ = v38
                            }
                            var v39 uint = 130
                            v40 := GenS1{F0: 394, F1: int64(v5)}
                            if (true || (!(255 != (v1 ^ v1)))) {
                                var v41 uint8 = v2
                                _ = v41
                            } else {
                                v42 := GenS4{F0: "\n\t"}
                                switch ((("\n\t" + ("foo bar" + (("" + "a") + ((("\n\t" + "hello") + "\n\t") + "a")))) + "hello") + "") {
                                case "hello":
                                    var v43 float64 = (-2.25 + 1.5)
                                    _ = v43
                                default:
                                    v5 = (^(fi7))
                                }
                                _ = v42
                            }
                            v4 = append(v4, 2147483647)
                            _ = fi26
                            _ = v39
                            _ = v40
                        }
                        if len(v4) > 0 {
                            v4[0] = 42
                        }
                    }
                }
                v5 += 2147483647
                _ = v24
            }
        case 601:
            var v44 uint = 4294967295
            _ = v44
        }
        _ = fi7
    }
    if true {
        for fi45 := 0; fi45 < 19; fi45++ {
            v4 = append(v4, 830)
            switch "a" {
            case "":
                if true {
                    v6 = 127
                    var v46 int = (^(v5))
                    var v47 uint32 = 4294967295
                    _ = v46
                    _ = v47
                } else {
                    v48 := GenS5{F0: 1}
                    _ = v48
                }
            }
            var v49 int = (-(fi45))
            _ = fi45
            _ = v49
        }
    }
    fmt.Println(v1)
    v50 := GenS6{F0: (v1 << ((42) & 63)), F1: (v2 * 2)}
    if (!true) {
        v51 := GenS2{F0: 4294967295}
        _ = v51
    } else {
        v52 := GenS1{F0: 1, F1: v3}
        fmt.Println(v2)
        v6 = v6
        _ = v52
    }
    v53 := GenS7{F0: 789, F1: (v5 >> ((967) & 63))}
    var v54 int = v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi55 := 0; fi55 < 18; fi55++ {
        switch ("foo bar" + "") {
        case "":
            var v56 float32 = (1e10 + (float32(float64(0.0)) / 1e10))
            _ = v56
        case "hello":
            v3 = (v3 / (((v3 / ((-1) | 1))) | 1))
        }
        if false {
            var v57 uint64 = 475
            var v58 uint64 = v57
            _ = v57
            _ = v58
        }
        if false {
            for fi59 := 0; fi59 < 3; fi59++ {
                if true {
                    var v60 uint32 = 1
                    _ = v60
                } else {
                    var v61 uint64 = 4294967295
                    fmt.Println(v3, v2)
                    v62 := GenS3{F0: (false || true)}
                    _ = v61
                    _ = v62
                }
                var v63 int16 = 32767
                v64 := GenS4{F0: "hello"}
                _ = fi59
                _ = v63
                _ = v64
            }
            if (255 <= uint16(v2)) {
                var v65 uint32 = 42
                if ((v2 | v2) < 128) {
                    if false {
                        if false {
                            if true {
                                v65 -= 2147483648
                                v2 -= (v2 & (v2 ^ 127))
                                v2 -= (v2 / (((^(v2))) | 1))
                            } else {
                                v5 -= (v5 & v5)
                                fmt.Println(v3, fi55, v65)
                                switch ("\n\t" + ("hello" + ("\n\t" + "a"))) {
                                case "foo bar":
                                    var v66 float64 = (((0.0 * 0.1) / -2.25) - 0.1)
                                    _ = v66
                                default:
                                    switch "foo bar" {
                                    case "a":
                                        for fi67 := 0; fi67 < 12; fi67++ {
                                            var v68 float64 = 1e10
                                            v6 = v6
                                            v3 = v3
                                            _ = fi67
                                            _ = v68
                                        }
                                    case "foo bar":
                                        fmt.Println(v1, v1, v65)
                                    }
                                }
                            }
                        }
                        v1 = v1
                    } else {
                        for fi69 := 0; fi69 < 19; fi69++ {
                            var v70 uint32 = (v65 * 118)
                            v1 += 2
                            _ = fi69
                            _ = v70
                        }
                    }
                    if true {
                        v2 = 255
                        var v71 uint32 = (v65 ^ 885)
                        for fi72 := 0; fi72 < 3; fi72++ {
                            if (!(false || ("foo bar" == ""))) {
                                var v73 int16 = 32767
                                _ = v73
                            }
                            _ = fi72
                        }
                        _ = v71
                    }
                } else {
                    v4 = append(v4, v65)
                    if len(v4) > 0 {
                        v4[0] = v65
                    }
                    v54 -= v5
                }
                var v74 float32 = (1e10 / 1.5)
                _ = v65
                _ = v74
            } else {
                v3 = 42
            }
        } else {
            if (-32768 == -32768) {
                fmt.Println(v5)
                if (false && true) {
                    v2 = (^(v2))
                    v75 := GenS4{F0: ("" + "\n\t")}
                    _ = v75
                }
                v1 -= (v1 << ((42) & 63))
            } else {
                if len(v4) > 0 {
                    v4[0] = 42
                }
                if false {
                    var v76 uint = 18446744073709551615
                    for fi77 := 0; fi77 < 17; fi77++ {
                        v76 -= (v76 | 458)
                        fmt.Println(v2, v3)
                        switch fi77 {
                        case 138:
                            v76 -= (v76 << (((v76 - 2)) & 63))
                            v78 := GenS8{F0: v3}
                            v2 -= (v2 ^ (v2 - v2))
                            _ = v78
                        case 128:
                            v79 := GenS2{F0: 2}
                            var v80 float32 = -2.25
                            var v81 uint64 = 0
                            _ = v79
                            _ = v80
                            _ = v81
                        }
                        _ = fi77
                    }
                    var v82 int8 = v6
                    _ = v76
                    _ = v82
                }
                for fi83 := 0; fi83 < 14; fi83++ {
                    switch ("a" + "\n\t") {
                    case "foo bar":
                        v2 += 4
                    case "\n\t":
                        v1 = v1
                    case "hello":
                        v84 := GenS1{F0: 580, F1: (v3 & v3)}
                        _ = v84
                    }
                    fmt.Println(v5, fi83)
                    _ = fi83
                }
            }
            if ((!(!(!(!true)))) || (((1 != 851) || true) || false)) {
                v85 := GenS7{F0: 65535, F1: (v54 | (v5 * fi55))}
                fmt.Println(v5, v3)
                _ = v85
            } else {
                var v86 uint64 = 42
                var v87 uint64 = v86
                if false {
                    var v88 uint64 = v86
                    var v89 uint32 = 2147483647
                    _ = v88
                    _ = v89
                } else {
                    v90 := GenS9{F0: -593, F1: -369, F2: v2}
                    _ = v90
                }
                _ = v86
                _ = v87
            }
            fmt.Println(v1)
        }
        _ = fi55
    }
    v5 -= (v54 / (((-(v54))) | 1))
    fmt.Println(v2, v2)
    switch (v5 / (((v5 * v54)) | 1)) {
    case -1:
        v2 += v2
    }
    v91 := GenS10{F0: false, F1: 1}
    _ = len(v4)
    for fi92 := 0; fi92 < 10; fi92++ {
        if (228 >= 0) {
            if len(v4) > 0 {
                v4[0] = 1
            }
            v93 := GenS11{F0: 1, F1: -266}
            v5 = (v54 ^ (v54 - -1))
            _ = v93
        }
        _ = fi92
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v53
    _ = v54
    _ = v6
    _ = v91
}
