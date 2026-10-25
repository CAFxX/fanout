package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 int8; F1 uint16 }
type GenS10 struct { F0 bool; F1 int16; F2 uint64 }
type GenS11 struct { F0 string; F1 string; F2 int }
type GenS12 struct { F0 uint64 }
func genFn0() uint32 {
    var v4 bool = false
    v5 := GenS1{F0: 105, F1: 127}
    fmt.Println(v4)
    switch ("hello" + "") {
    case "hello":
        v6 := []int{4294967295, -652, 1, -709}
        _ = v6
    case "foo bar":
        var v7 uint8 = 1
        _ = v7
    case "a":
        v4 = v4
    default:
        v8 := []int{128, -473, 294}
        _ = v8
    }
    if (v4 || (42 != 1)) {
        v9 := GenS1{F0: 710, F1: 32767}
        _ = v9
    } else {
        v4 = (255 <= 42)
        if true {
            var v10 int64 = 627
            _ = v10
        }
        var v11 int = 4294967295
        _ = v11
    }
    _ = v4
    _ = v5
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (700 < 128)
    v3 := make(map[int]bool)
    var v12 uint64 = 4294967295
    switch "" {
    case "hello":
        if false {
            fmt.Println(v12, v12)
            if (!(127 < 42)) {
                if true {
                    var v13 uint32 = 2147483647
                    _ = v13
                }
                v12 = 0
                if false {
                    if (-1 < -2147483648) {
                        v14 := GenS1{F0: 65535, F1: 127}
                        if (v2 && v2) {
                            var v15 uint16 = 65535
                            for fi16 := 0; fi16 < 15; fi16++ {
                                if v2 {
                                    v2 = v2
                                }
                                _ = fi16
                            }
                            v17 := []int{-940, -598, 4294967295}
                            _ = v15
                            _ = v17
                        } else {
                            fmt.Println(v12, v2)
                        }
                        _ = v14
                    }
                    if v2 {
                        v18 := []int{4294967295, 4294967295, 4294967295}
                        v12 = 42
                        _ = v18
                    } else {
                        if true {
                            fmt.Println(v12, v2)
                            v12 += v12
                        }
                        switch ("a" + "hello") {
                        case "a":
                            v19 := GenS1{F0: 728, F1: 42}
                            _ = v19
                        case "foo bar":
                            if (false || (v2 || ((1.5 + -2.25) != (-2.25 - (1.5 - ((0.1 / 0.1) + ((0.0 / (1e10 - (0.0 + float32(1.5)))) / (0.1 - (0.1 - 0.1))))))))) {
                                switch 4294967295 {
                                case -976:
                                    fmt.Println(v2)
                                case 51:
                                    switch 952 {
                                    case -1:
                                        v12 += (v12 ^ v12)
                                    case 492:
                                        v20 := GenS1{F0: 2, F1: 1}
                                        _ = v20
                                    case 128:
                                        for fi21 := 0; fi21 < 18; fi21++ {
                                            if v2 {
                                                if ("foo bar" == "") {
                                                    fmt.Println(fi21, v12)
                                                    v22 := []int{}
                                                    v2 = v2
                                                    _ = v22
                                                } else {
                                                    v23 := GenS2{F0: 127}
                                                    var v24 float32 = 1e10
                                                    if v2 {
                                                        v25 := GenS3{F0: 42, F1: (v12 >> ((18446744073709551615) & 63))}
                                                        var v26 int8 = -36
                                                        _ = v25
                                                        _ = v26
                                                    }
                                                    _ = v23
                                                    _ = v24
                                                }
                                                for fi27 := 0; fi27 < 19; fi27++ {
                                                    v12 = (^(v12))
                                                    _ = fi27
                                                }
                                            } else {
                                                if true {
                                                    var v28 int8 = 127
                                                    v29 := GenS2{F0: 255}
                                                    if (195 < 609) {
                                                        _ = len(v3)
                                                    } else {
                                                        fmt.Println(v12)
                                                    }
                                                    _ = v28
                                                    _ = v29
                                                } else {
                                                    v2 = v2
                                                }
                                                var v30 float64 = 1e10
                                                fmt.Println(v2)
                                                _ = v30
                                            }
                                            for fi31 := 0; fi31 < 1; fi31++ {
                                                var v32 bool = ((false || v2) && (-639 > -70))
                                                if v32 {
                                                    var v33 uint32 = 394
                                                    if ((!v2) && false) {
                                                        var v34 float64 = 0.1
                                                        if (-1 < -1) {
                                                            var v35 string = "hello"
                                                            if (v32 && (v33 != (^(v33)))) {
                                                                v33 = v33
                                                            } else {
                                                                v36 := GenS2{F0: uint8(v33)}
                                                                var v37 uint = 0
                                                                v35 = (v35 + "foo bar")
                                                                _ = v36
                                                                _ = v37
                                                            }
                                                            v38 := v3[fi21]
                                                            _ = v35
                                                            _ = v38
                                                        }
                                                        switch fi31 {
                                                        case -9223372036854775808:
                                                            delete(v3, 1)
                                                        default:
                                                            fmt.Println(v32, v12)
                                                        }
                                                        v33 -= v33
                                                        var v39 uint64 = (v12 | (^(v12)))
                                                        _ = v34
                                                        _ = v39
                                                    }
                                                    _ = v33
                                                } else {
                                                    var v40 uint64 = 4294967295
                                                    v41 := GenS4{F0: -128, F1: 307}
                                                    v12 += 489
                                                    _ = v40
                                                    _ = v41
                                                }
                                                v12 = v12
                                                v42 := GenS1{F0: uint(v12), F1: fi21}
                                                _ = fi31
                                                _ = v32
                                                _ = v42
                                            }
                                            v43 := GenS5{F0: 65535, F1: -327}
                                            _ = fi21
                                            _ = v43
                                        }
                                    default:
                                        v12 -= 286
                                    }
                                case 2147483647:
                                    v44 := GenS6{F0: 255, F1: 42}
                                    _ = v44
                                default:
                                    var v45 int = -398
                                    _ = v45
                                }
                                v12 = 4294967295
                                v46 := GenS7{F0: 318, F1: 18446744073709551615, F2: 127}
                                _ = v46
                            }
                        default:
                            for fi47 := 0; fi47 < 9; fi47++ {
                                var v48 int16 = -557
                                v2 = (fi47 <= 4294967295)
                                _ = fi47
                                _ = v48
                            }
                        }
                    }
                    v12 = (v12 * 0)
                } else {
                    if v2 {
                        v2 = v2
                    } else {
                        v49 := v3[0]
                        _ = v49
                    }
                    switch -866 {
                    case 32767:
                        fmt.Println(v2, v12)
                    }
                    var v50 string = "\n\t"
                    _ = v50
                }
            } else {
                v51 := GenS8{F0: v2, F1: 255}
                if v2 {
                    var v52 uint = 1
                    switch "hello" {
                    case "a":
                        v53 := []int{4294967295, 1, 388, 1}
                        _ = v53
                    }
                    _ = v52
                }
                v54 := GenS9{F0: 46, F1: 2}
                _ = v51
                _ = v54
            }
        }
    case "":
        for fi55 := 0; fi55 < 6; fi55++ {
            fmt.Println(v12, v2, fi55)
            var v56 uint = 91
            _ = len(v1)
            _ = fi55
            _ = v56
        }
    }
    v12 = v12
    v12 = (v12 >> ((255) & 63))
    v12 = (v12 & (v12 ^ (v12 * v12)))
    var v57 uint16 = 312
    if (true || v2) {
        if (!v2) {
            var v58 uint64 = v12
            if v2 {
                if v2 {
                    v59 := GenS3{F0: (^(v12)), F1: v12}
                    fmt.Println(v58)
                    var v60 int16 = -868
                    _ = v59
                    _ = v60
                }
                if v2 {
                    var v61 int = 32767
                    if v2 {
                        v61 += (v61 % (((v61 >> ((2) & 63))) | 1))
                    }
                    fmt.Println(v12, v58)
                    _ = v61
                } else {
                    for fi62 := 0; fi62 < 3; fi62++ {
                        var v63 uint64 = (^(v58))
                        _ = fi62
                        _ = v63
                    }
                    switch ("foo bar" + "") {
                    case "foo bar":
                        v12 -= 0
                    case "\n\t":
                        v64 := []int{4294967295, -1}
                        _ = v64
                    case "hello":
                        v58 += v12
                    default:
                        v65 := GenS2{F0: uint8(v58)}
                        _ = v65
                    }
                    switch (("a" + "") + "a") {
                    case "hello":
                        var v66 int32 = -1
                        _ = v66
                    case "foo bar":
                        v2 = v2
                    default:
                        for fi67 := 0; fi67 < 18; fi67++ {
                            switch "" {
                            case "foo bar":
                                if v2 {
                                    for fi68 := 0; fi68 < 20; fi68++ {
                                        var v69 float32 = 0.1
                                        _ = fi68
                                        _ = v69
                                    }
                                    if v2 {
                                        v57 = 255
                                    } else {
                                        v70 := GenS9{F0: 125, F1: v57}
                                        _ = v70
                                    }
                                }
                            case "":
                                for fi71 := 0; fi71 < 15; fi71++ {
                                    fmt.Println(fi71, fi67, v58)
                                    switch 1 {
                                    case -2147483648:
                                        v72 := GenS10{F0: ("\n\t" <= "a"), F1: 0, F2: (v12 % (((v58 & (v58 << ((1) & 63)))) | 1))}
                                        _ = v72
                                    case 127:
                                        var v73 uint64 = 988
                                        _ = v73
                                    case 2147483647:
                                        var v74 float64 = 0.0
                                        _ = v74
                                    default:
                                        if false {
                                            switch ("\n\t" + ("foo bar" + "")) {
                                            case "\n\t":
                                                v3[(-(fi71))] = v2
                                            default:
                                                v75 := GenS4{F0: 100, F1: (v57 << ((1) & 63))}
                                                _ = v75
                                            }
                                        }
                                    }
                                    _ = fi71
                                }
                            default:
                                var v76 float64 = 0.1
                                _ = v76
                            }
                            _ = fi67
                        }
                    }
                }
            } else {
                var v77 float64 = 1e10
                _ = v77
            }
            _ = v58
        } else {
            delete(v1, "\n\t")
            v78 := []int{}
            if (!true) {
                var v79 uint8 = 227
                for fi80 := 0; fi80 < 12; fi80++ {
                    fmt.Println(v57, fi80, fi80)
                    _ = fi80
                }
                if false {
                    if v2 {
                        var v81 float32 = (-2.25 * float32((((1.5 - 3.14159) + 1e10) - (((1.5 * 1e10) / 1.0) / ((1e10 * ((-2.25 * (1.5 / 1e10)) - 1.5)) - (1.5 + (1e10 / 3.14159)))))))
                        for fi82 := 0; fi82 < 13; fi82++ {
                            fmt.Println(v81, v81)
                            _ = fi82
                        }
                        for fi83 := 0; fi83 < 6; fi83++ {
                            v84 := GenS11{F0: ("\n\t" + ""), F1: "\n\t", F2: (fi83 % (((-(fi83))) | 1))}
                            _ = fi83
                            _ = v84
                        }
                        _ = v81
                    }
                } else {
                    var v85 uint16 = 32767
                    _ = v85
                }
                _ = v79
            } else {
                delete(v3, 755)
                fmt.Println(v12, v57)
            }
            _ = v78
        }
        if (960 > 1) {
            fmt.Println(v2, v57, v2)
        } else {
            for fi86 := 0; fi86 < 18; fi86++ {
                var v87 bool = ((57 == 42) || v2)
                _ = fi86
                _ = v87
            }
            delete(v3, -722)
            v2 = v2
        }
    }
    for fi88 := 0; fi88 < 1; fi88++ {
        var v89 int8 = 79
        if (-1 < 947) {
            v90 := GenS12{F0: 18446744073709551615}
            v2 = v2
            _ = v90
        }
        _ = fi88
        _ = v89
    }
    var v91 bool = (v12 >= 4294967295)
    var v92 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v57:", v57)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v57
    _ = v91
    _ = v92
}
