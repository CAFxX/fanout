package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 int16; F1 uint; F2 int64 }
type GenS6 struct { F0 int; F1 int32; F2 uint32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS9 struct { F0 uint; F1 uint16 }
type GenS10 struct { F0 int8; F1 uint16; F2 int16 }
type GenS11 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS12 struct { F0 uint; F1 string }
type GenS13 struct { F0 int64; F1 int64; F2 bool }
type GenS14 struct { F0 int64; F1 string }
type GenS15 struct { F0 uint8; F1 uint16 }
type GenS16 struct { F0 int8; F1 uint16; F2 string }
type GenS17 struct { F0 uint16; F1 bool; F2 int16 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    v3 := []uint32{512, 674, 2, 570}
    v4 := make(map[uint32]bool)
    switch "\n\t" {
    case "foo bar":
        var v5 uint64 = 2
        _ = v5
    case "a":
        switch -131 {
        case -1:
            if (!false) {
                var v6 int32 = 2147483647
                var v7 uint = 18446744073709551615
                _ = v6
                _ = v7
            } else {
                for fi8 := 0; fi8 < 20; fi8++ {
                    var v9 int64 = -520
                    v9 = -1
                    _ = len(v4)
                    _ = fi8
                    _ = v9
                }
                switch "a" {
                case "a":
                    var v10 int32 = -502
                    _ = v10
                default:
                    v11 := GenS1{F0: 682}
                    _ = v11
                }
                var v12 int8 = -121
                _ = v12
            }
        case -78:
            var v13 float64 = (float64(1.5) * (0.1 - 1.5))
            _ = v13
        case 9223372036854775807:
            if true {
                _ = len(v4)
            }
        default:
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 int = fi14
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi14
                _ = v15
            }
        }
    case "hello":
        v16 := GenS2{F0: 18446744073709551615}
        _ = v16
    }
    var v17 uint16 = 32767
    v18 := GenS3{F0: 42, F1: 0}
    v17 = v17
    v17 = (^(v17))
    if (244 != 1) {
        switch (("\n\t" + "hello") + ((("a" + "a") + "foo bar") + ((("hello" + ("foo bar" + "hello")) + "hello") + "foo bar"))) {
        case "hello":
            var v19 string = ("a" + "")
            _ = v19
        case "a":
            switch 668 {
            case 127:
                if true {
                    _ = len(v2)
                    v20 := GenS4{F0: 658, F1: v17}
                    _ = v20
                } else {
                    if true {
                        fmt.Println(v17)
                        for fi21 := 0; fi21 < 6; fi21++ {
                            var v22 uint = 0
                            if ((^(v17)) >= v17) {
                                for fi23 := 0; fi23 < 12; fi23++ {
                                    v22 += v22
                                    var v24 int16 = 992
                                    delete(v2, "a")
                                    _ = fi23
                                    _ = v24
                                }
                            } else {
                                v25 := GenS3{F0: 1, F1: 42}
                                v26 := GenS5{F0: -32768, F1: 0, F2: -9223372036854775808}
                                _ = v25
                                _ = v26
                            }
                            _ = fi21
                            _ = v22
                        }
                    }
                    for fi27 := 0; fi27 < 6; fi27++ {
                        for fi28 := 0; fi28 < 19; fi28++ {
                            switch "a" {
                            case "":
                                var v29 int64 = 9223372036854775807
                                _ = v29
                            case "foo bar":
                                var v30 int32 = 1
                                _ = v30
                            case "a":
                                v31 := GenS4{F0: 42, F1: (^(v17))}
                                _ = v31
                            default:
                                v32 := GenS4{F0: -759, F1: 0}
                                _ = v32
                            }
                            _ = fi28
                        }
                        _ = fi27
                    }
                    if true {
                        v33 := GenS6{F0: -764, F1: 0, F2: 2147483648}
                        _ = v33
                    }
                }
            default:
                if false {
                    v34 := GenS6{F0: 9223372036854775807, F1: -1, F2: 0}
                    _ = v34
                } else {
                    if (42 != 255) {
                        var v35 int8 = 45
                        _ = v35
                    }
                    var v36 int64 = 926
                    v37 := GenS4{F0: 42, F1: (v17 - 42)}
                    _ = v36
                    _ = v37
                }
            }
        case "foo bar":
            fmt.Println(v17)
        }
    } else {
        for fi38 := 0; fi38 < 13; fi38++ {
            for fi39 := 0; fi39 < 10; fi39++ {
                var v40 int8 = 0
                if false {
                    fmt.Println(fi38)
                    v17 = 173
                } else {
                    v41 := GenS5{F0: 42, F1: 1, F2: 1}
                    v42 := GenS7{F0: 93}
                    _ = v41
                    _ = v42
                }
                var v43 uint8 = 2
                var v44 int = (fi38 << ((950) & 63))
                _ = fi39
                _ = v40
                _ = v43
                _ = v44
            }
            _ = fi38
        }
        switch -398 {
        case 724:
            switch 128 {
            case 32768:
                v3 = append(v3, 1)
            case 2147483647:
                v17 += v17
            case -790:
                if (-32768 > -266) {
                    var v45 int16 = 42
                    _ = v45
                } else {
                    fmt.Println(v17)
                    var v46 float64 = 3.14159
                    for fi47 := 0; fi47 < 11; fi47++ {
                        for fi48 := 0; fi48 < 16; fi48++ {
                            v49 := GenS8{F0: 1, F1: 7, F2: (v17 / (((v17 << ((0) & 63))) | 1))}
                            var v50 float32 = 3.14159
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                        _ = fi47
                    }
                    _ = v46
                }
            default:
                var v51 uint32 = 0
                _ = v51
            }
        default:
            v17 += (v17 + (v17 - (^(v17))))
        }
        v52 := GenS4{F0: 2147483647, F1: v17}
        _ = v52
    }
    if true {
        v17 -= v17
        v53 := GenS8{F0: 223, F1: 128, F2: v17}
        _ = v53
    }
    v54 := v2[(("" + "a") + "hello")]
    for fi55 := 0; fi55 < 1; fi55++ {
        for fi56 := 0; fi56 < 17; fi56++ {
            v54 = v54
            _ = fi56
        }
        _ = fi55
    }
    if false {
        if ((0.0 + (1.5 - 3.14159)) >= 1.5) {
            v54 = (v54 && v54)
            var v57 int32 = -1
            v57 = -1
            _ = v57
        }
    } else {
        switch 4294967295 {
        case 263:
            v17 += 612
        case -1:
            v17 = (^(v17))
        default:
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if (-130 != 246) {
                    switch 127 {
                    case -385:
                        fmt.Println(v54, v54)
                    case 722:
                        var v58 uint = 2
                        _ = v58
                    default:
                        v59 := GenS9{F0: 723, F1: v17}
                        _ = v59
                    }
                    v17 = 0
                    v17 -= 74
                } else {
                    var v60 float32 = 1.5
                    _ = v60
                }
            } else {
                if (81 < 0) {
                    for fi61 := 0; fi61 < 9; fi61++ {
                        fmt.Println(v17)
                        fmt.Println(v17)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi61
                    }
                } else {
                    v17 = v17
                }
            }
        }
        if v54 {
            if len(v3) > 0 {
                v3[0] = 42
            }
        } else {
            var v62 uint16 = v17
            fmt.Println(v54, v62)
            _ = v62
        }
    }
    v54 = ("hello" >= ("a" + ("" + "\n\t")))
    switch "hello" {
    case "\n\t":
        switch "hello" {
        case "hello":
            for fi63 := 0; fi63 < 5; fi63++ {
                for fi64 := 0; fi64 < 19; fi64++ {
                    v65 := GenS10{F0: 1, F1: (v17 << ((42) & 63)), F2: 1}
                    v66 := GenS11{F0: 0, F1: 411, F2: 42}
                    if (2147483648 < 2147483647) {
                        var v67 int8 = 1
                        var v68 int16 = 964
                        if (65535 >= 1) {
                            switch "hello" {
                            case "a":
                                if v54 {
                                    var v69 uint16 = (^(v17))
                                    _ = v69
                                } else {
                                    v70 := GenS12{F0: 65535, F1: "\n\t"}
                                    if (!v54) {
                                        var v71 int64 = -132
                                        _ = v71
                                    }
                                    v54 = true
                                    _ = v70
                                }
                            case "foo bar":
                                v72 := GenS6{F0: 32768, F1: -2147483648, F2: 4294967295}
                                _ = v72
                            case "hello":
                                v73 := GenS13{F0: 9223372036854775807, F1: -1, F2: (false || (1 < (v68 | v68)))}
                                _ = v73
                            }
                            switch ("\n\t" + ("hello" + "a")) {
                            case "":
                                v3 = append(v3, 341)
                            case "foo bar":
                                v74 := GenS14{F0: 365, F1: ""}
                                _ = v74
                            default:
                                var v75 float32 = float32((1.5 / (-2.25 * 1e10)))
                                _ = v75
                            }
                        }
                        _ = v67
                        _ = v68
                    } else {
                        v76 := GenS15{F0: 2, F1: 65535}
                        _ = v76
                    }
                    _ = fi64
                    _ = v65
                    _ = v66
                }
                _ = fi63
            }
        }
    case "":
        if v54 {
            v77 := GenS16{F0: -128, F1: (^(v17)), F2: ""}
            if (float64((float32((0.0 * 3.14159)) - (0.1 - -2.25))) > 0.1) {
                for fi78 := 0; fi78 < 1; fi78++ {
                    delete(v2, "hello")
                    delete(v4, 2147483647)
                    if true {
                        v3 = append(v3, 1)
                        var v79 uint = 42
                        _ = v79
                    } else {
                        fmt.Println(v17)
                        for fi80 := 0; fi80 < 12; fi80++ {
                            v81 := GenS11{F0: 828, F1: 2147483647, F2: 0}
                            _ = fi80
                            _ = v81
                        }
                        if v54 {
                            var v82 string = ""
                            v1[244] = v82
                            _ = v82
                        }
                    }
                    _ = fi78
                }
            }
            for fi83 := 0; fi83 < 4; fi83++ {
                fmt.Println(fi83)
                v84 := GenS17{F0: v17, F1: (!v54), F2: 1}
                var v85 uint = 1
                _ = fi83
                _ = v84
                _ = v85
            }
            _ = v77
        } else {
            if len(v3) > 0 {
                v3[0] = 177
            }
            if v54 {
                for fi86 := 0; fi86 < 14; fi86++ {
                    for fi87 := 0; fi87 < 10; fi87++ {
                        switch "hello" {
                        case "hello":
                            var v88 uint32 = 2147483647
                            _ = v88
                        default:
                            delete(v2, "\n\t")
                        }
                        _ = fi87
                    }
                    _ = fi86
                }
                v17 -= v17
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    case "a":
        for fi89 := 0; fi89 < 12; fi89++ {
            var v90 int64 = 677
            var v91 int32 = 1
            fmt.Println(v54)
            _ = fi89
            _ = v90
            _ = v91
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v54:", v54)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v54
}
