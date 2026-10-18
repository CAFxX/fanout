package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint }
type GenS2 struct { F0 string; F1 int16; F2 int64 }
type GenS3 struct { F0 bool; F1 uint8; F2 int64 }
type GenS4 struct { F0 string; F1 uint32; F2 uint }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 string; F1 uint8; F2 string }
func genFn0(p0 uint8) int8 {
    v7 := []int{127, 32767}
    switch ("\n\t" + "foo bar") {
    case "a":
        if len(v7) > 0 {
            v7[0] = -9223372036854775808
        }
    }
    _ = p0
    _ = v7
    return 18
}

func main() {
    v1 := []string{"\n\t", ("" + "a"), ("hello" + ((("a" + ("\n\t" + (("" + "") + "\n\t"))) + "\n\t") + ""))}
    var v2 int = 127
    v3 := []uint32{2147483648, 2, 0}
    var v4 uint64 = 42
    var v5 uint16 = 101
    var v6 int64 = -451
    if (!true) {
        for fi8 := 0; fi8 < 17; fi8++ {
            v6 -= (v6 / (((v6 | v6)) | 1))
            var v9 bool = (!(-32768 < 32767))
            if (!v9) {
                if (42 > 963) {
                    _ = len(v1)
                }
                var v10 float64 = 1e10
                _ = v10
            } else {
                v5 = v5
                v4 = v4
                var v11 uint32 = 2147483647
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    }
    var v12 uint64 = 2
    switch v2 {
    case -214:
        for fi13 := 0; fi13 < 13; fi13++ {
            v14 := GenS1{F0: v2, F1: 65535, F2: 1}
            if true {
                var v15 int8 = -75
                if true {
                    v15 -= int8(v6)
                    for fi16 := 0; fi16 < 17; fi16++ {
                        v15 = (v15 >> ((255) & 63))
                        var v17 uint32 = 815
                        _ = fi16
                        _ = v17
                    }
                    fmt.Println(v12)
                } else {
                    v18 := GenS1{F0: fi13, F1: v5, F2: 255}
                    v19 := GenS1{F0: v2, F1: 2, F2: 18446744073709551615}
                    _ = v18
                    _ = v19
                }
                _ = v15
            } else {
                for fi20 := 0; fi20 < 18; fi20++ {
                    if ((false && (false || false)) || ((^(v12)) >= v4)) {
                        var v21 int32 = -2147483648
                        _ = v21
                    } else {
                        if false {
                            fmt.Println(v6, v4)
                        }
                        fmt.Println(v4, v4)
                    }
                    _ = fi20
                }
                v12 = 595
                if (!false) {
                    v22 := GenS1{F0: fi13, F1: v5, F2: 2}
                    _ = v22
                }
            }
            _ = fi13
            _ = v14
        }
    default:
        fmt.Println(v4)
    }
    var v23 uint8 = 0
    for fi24 := 0; fi24 < 5; fi24++ {
        v12 -= (v4 >> ((349) & 63))
        var v25 string = "\n\t"
        _ = fi24
        _ = v25
    }
    v26 := GenS1{F0: 32767, F1: uint16(v23), F2: 255}
    var v27 float32 = (1.5 / 3.14159)
    if false {
        _ = len(v3)
        var v28 bool = (v12 >= v4)
        if (v28 && false) {
            switch v2 {
            case 32768:
                v29 := GenS2{F0: "a", F1: int16(v23), F2: v6}
                _ = v29
            case 128:
                for fi30 := 0; fi30 < 12; fi30++ {
                    var v31 int32 = 2147483647
                    v32 := GenS2{F0: "a", F1: 0, F2: (^(v6))}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            case 1:
                v3 = append(v3, 42)
            default:
                for fi33 := 0; fi33 < 9; fi33++ {
                    for fi34 := 0; fi34 < 2; fi34++ {
                        if true {
                            v35 := GenS3{F0: v28, F1: (v23 * (v23 % ((v23) | 1))), F2: (-(v6))}
                            _ = v35
                        } else {
                            if (2147483647 > 2) {
                                switch "" {
                                case "\n\t":
                                    var v36 int = fi33
                                    _ = v36
                                case "":
                                    var v37 bool = v28
                                    _ = v37
                                default:
                                    var v38 int = v2
                                    _ = v38
                                }
                                switch -1 {
                                case 2147483647:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v28 = (((((float64(v27) + 1e10) - 0.0) - (1e10 * -2.25)) / -2.25) < float64(v27))
                                    v27 += float32(1e10)
                                case 0:
                                    var v39 float64 = 1.5
                                    _ = v39
                                case 42:
                                    fmt.Println(v6, v28, v27)
                                }
                                if (v28 && v28) {
                                    var v40 bool = true
                                    _ = v40
                                } else {
                                    if (255 != (v5 / ((v5) | 1))) {
                                        fmt.Println(fi33, v28)
                                    } else {
                                        if (42 == 2) {
                                            if (!("hello" == ("\n\t" + "hello"))) {
                                                fmt.Println(v2, v28)
                                            }
                                        }
                                        var v41 int64 = v6
                                        _ = v41
                                    }
                                    var v42 uint32 = 0
                                    var v43 uint16 = (v5 - v5)
                                    _ = v42
                                    _ = v43
                                }
                                v12 -= v12
                            }
                            var v44 float32 = (1e10 - (v27 / v27))
                            _ = v44
                        }
                        _ = fi34
                    }
                    v23 = v23
                    var v45 int16 = -796
                    _ = fi33
                    _ = v45
                }
            }
            _ = len(v3)
            if ("hello" < "a") {
                switch "hello" {
                case "":
                    v46 := GenS4{F0: "hello", F1: 2147483647, F2: uint(v12)}
                    _ = v46
                default:
                    fmt.Println(v6, v4, v5)
                }
            } else {
                if (2 <= 1) {
                    v12 += 4294967295
                    v47 := GenS4{F0: "hello", F1: 4, F2: 189}
                    _ = v47
                } else {
                    for fi48 := 0; fi48 < 11; fi48++ {
                        fmt.Println(v23)
                        fmt.Println(v4, v23)
                        v1 = append(v1, ("a" + "\n\t"))
                        _ = fi48
                    }
                }
                v1 = append(v1, ("hello" + "hello"))
                v1 = append(v1, "\n\t")
            }
        } else {
            if false {
                if (false || true) {
                    v49 := GenS4{F0: "foo bar", F1: 918, F2: uint(v23)}
                    var v50 uint8 = uint8(v5)
                    _ = v49
                    _ = v50
                }
                v4 -= v4
                for fi51 := 0; fi51 < 12; fi51++ {
                    v5 += 2
                    for fi52 := 0; fi52 < 20; fi52++ {
                        _ = len(v1)
                        if v28 {
                            var v53 uint32 = 806
                            _ = v53
                        } else {
                            fmt.Println(v6)
                        }
                        _ = fi52
                    }
                    _ = fi51
                }
            }
            v23 = (v23 % (((v23 ^ v23)) | 1))
        }
        _ = v28
    } else {
        if (!(true || (127 <= int8(v6)))) {
            fmt.Println(v23, v5, v4)
        }
    }
    if false {
        var v54 uint8 = (v23 - v23)
        var v55 int16 = 32767
        _ = v54
        _ = v55
    } else {
        var v56 uint = 212
        switch ("foo bar" + "hello") {
        case "":
            var v57 uint = (v56 ^ v56)
            _ = v57
        case "hello":
            var v58 int8 = -1
            _ = v58
        case "foo bar":
            _ = len(v3)
        }
        var v59 uint8 = v23
        _ = v56
        _ = v59
    }
    var v60 uint8 = (v23 % ((v23) | 1))
    for fi61 := 0; fi61 < 10; fi61++ {
        if false {
            v4 = (v4 >> ((301) & 63))
        } else {
            v2 = fi61
        }
        if true {
            _ = len(v1)
            if ((!(!true)) && true) {
                var v62 int16 = -32768
                v63 := GenS1{F0: v2, F1: (v5 - (v5 & v5)), F2: 65535}
                _ = v62
                _ = v63
            } else {
                _ = len(v1)
                if false {
                    if false {
                        var v64 int8 = -128
                        for fi65 := 0; fi65 < 10; fi65++ {
                            v66 := GenS4{F0: ("hello" + ("\n\t" + "hello")), F1: 2, F2: 1}
                            var v67 int32 = -70
                            _ = fi65
                            _ = v66
                            _ = v67
                        }
                        v3 = append(v3, 603)
                        _ = v64
                    } else {
                        v5 = v5
                        v2 += fi61
                    }
                    var v68 int16 = 868
                    var v69 int16 = (^(v68))
                    _ = v68
                    _ = v69
                } else {
                    if (-1 > 1000) {
                        if (!((v23 * (v60 << ((915) & 63))) == v23)) {
                            v3 = append(v3, 2147483648)
                            var v70 string = "\n\t"
                            fmt.Println(v70)
                            _ = v70
                        }
                        var v71 int32 = -1
                        fmt.Println(v60, v27)
                        _ = v71
                    } else {
                        for fi72 := 0; fi72 < 8; fi72++ {
                            switch "a" {
                            case "hello":
                                if true {
                                    var v73 uint64 = (v12 | v4)
                                    var v74 uint8 = v60
                                    _ = v73
                                    _ = v74
                                }
                                if (4294967295 < 854) {
                                    switch fi61 {
                                    case 128:
                                        fmt.Println(v2, v6)
                                    case 42:
                                        var v75 float32 = ((((v27 / 3.14159) * 1e10) * v27) + 1.5)
                                        _ = v75
                                    default:
                                        v2 -= -940
                                    }
                                    v2 = fi72
                                }
                                v5 -= 42
                                var v76 uint8 = v60
                                _ = v76
                            case "":
                                v6 -= 42
                            case "a":
                                var v77 int8 = 127
                                _ = v77
                            }
                            if (float64(v27) >= 3.14159) {
                                fmt.Println(v60, v6, v27)
                                switch fi72 {
                                case 4294967295:
                                    var v78 bool = (!false)
                                    _ = v78
                                case -465:
                                    v3 = append(v3, 590)
                                default:
                                    _ = len(v1)
                                }
                                var v79 uint16 = 488
                                _ = v79
                            }
                            _ = fi72
                        }
                        fmt.Println(v23)
                    }
                    fmt.Println(v23)
                }
                var v80 uint8 = uint8(v12)
                _ = v80
            }
            if false {
                v4 = v4
                for fi81 := 0; fi81 < 6; fi81++ {
                    v5 -= v5
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(v23)
                    _ = fi81
                }
                var v82 string = ("\n\t" + (("" + "foo bar") + "\n\t"))
                _ = v82
            } else {
                for fi83 := 0; fi83 < 3; fi83++ {
                    v12 -= v12
                    v12 = 4294967295
                    _ = fi83
                }
            }
        } else {
            fmt.Println(v27)
            v84 := GenS5{F0: v23, F1: 31}
            for fi85 := 0; fi85 < 5; fi85++ {
                if false {
                    _ = len(v3)
                }
                v86 := GenS6{F0: "", F1: 42, F2: "hello"}
                var v87 float32 = (v27 - 1.5)
                _ = fi85
                _ = v86
                _ = v87
            }
            _ = v84
        }
        _ = fi61
    }
    var v88 string = ""
    for fi89 := 0; fi89 < 9; fi89++ {
        fmt.Println(v5)
        v90 := GenS6{F0: ("a" + v88), F1: (v60 + (v60 - v23)), F2: (v88 + "foo bar")}
        if len(v3) > 0 {
            v3[0] = 948
        }
        _ = fi89
        _ = v90
    }
    for fi91 := 0; fi91 < 17; fi91++ {
        for fi92 := 0; fi92 < 7; fi92++ {
            switch "" {
            case "hello":
                fmt.Println(v88, v5)
            default:
                fmt.Println(v88)
            }
            if true {
                v2 = v2
            } else {
                fmt.Println(v88, v2, fi92)
            }
            if false {
                _ = len(v1)
                var v93 int16 = 1
                if (0.1 != 3.14159) {
                    var v94 uint16 = v5
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v94
                }
                if true {
                    if true {
                        v4 -= v12
                    }
                    v6 = (v6 * (v6 / ((v6) | 1)))
                    v12 = v12
                }
                fmt.Println(v6, v4)
                _ = v93
            } else {
                var v95 float64 = (-2.25 - (float64((1.5 / ((0.1 - 3.14159) / (1.5 / ((v27 - v27) - (0.0 - (v27 + (v27 / 1e10)))))))) * (float64(((v27 * v27) - float32(float64(v27)))) * 0.0)))
                if len(v3) > 0 {
                    v3[0] = 42
                }
                _ = v95
            }
            fmt.Println(v6, v6)
            if (-95 <= -1) {
                v96 := GenS3{F0: false, F1: 127, F2: 42}
                if (!false) {
                    fmt.Println(v6, fi92)
                }
                fmt.Println(v27, v88, v88)
                _ = v96
            }
            var v97 float64 = float64(-2.25)
            _ = fi92
            _ = v97
        }
        var v98 int = (^(v2))
        _ = fi91
        _ = v98
    }
    for fi99 := 0; fi99 < 11; fi99++ {
        var v100 int64 = (v6 * v6)
        var v101 uint = 42
        _ = fi99
        _ = v100
        _ = v101
    }
    fmt.Println(v4, v27, v23)
    switch v88 {
    case "\n\t":
        v2 = (v2 | (v2 >> ((1) & 63)))
    default:
        v88 = ""
    }
    v102 := GenS2{F0: ((v88 + v88) + v88), F1: int16(v2), F2: v6}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v102
    _ = v12
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v88
}
