package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int32; F1 bool; F2 uint }
type GenS4 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS5 struct { F0 int; F1 int64; F2 string }

func main() {
    var v1 int8 = -128
    v2 := []string{"hello"}
    v3 := []int{0, -163}
    v4 := []string{}
    var v5 uint8 = 89
    var v6 int8 = (v1 | 1)
    fmt.Println(v1)
    v7 := GenS1{F0: -151, F1: 1, F2: 42}
    v2 = append(v2, (("" + "foo bar") + "a"))
    v8 := GenS2{F0: 18446744073709551615, F1: -1}
    if ((42 >= (v1 | (v6 * v6))) && true) {
        v9 := GenS3{F0: -2147483648, F1: true, F2: uint(v5)}
        var v10 bool = ((!true) && true)
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 int16 = -32768
            v5 = (v5 >> ((0) & 63))
            v13 = (v13 & int16(fi12))
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v14 uint8 = (v5 << ((18446744073709551615) & 63))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        fmt.Println(v5, v5, v5)
        switch -966 {
        case 128:
            if true {
                if len(v3) > 0 {
                    v3[0] = 829
                }
                var v15 uint8 = v14
                var v16 int8 = 127
                _ = v15
                _ = v16
            } else {
                if (((v14 < (v14 | v5)) && (v5 == (v5 - 1))) || true) {
                    for fi17 := 0; fi17 < 18; fi17++ {
                        var v18 uint32 = 908
                        _ = fi17
                        _ = v18
                    }
                    var v19 int32 = 117
                    _ = v19
                }
                if false {
                    v5 = (v5 / (((v14 | (v14 - (v5 - v14)))) | 1))
                    switch 987 {
                    case 9223372036854775807:
                        fmt.Println(v1, v14, v6)
                    case 1:
                        if true {
                            fmt.Println(v1, v5)
                            v20 := GenS2{F0: 397, F1: -730}
                            var v21 int16 = 516
                            _ = v20
                            _ = v21
                        }
                    }
                } else {
                    for fi22 := 0; fi22 < 7; fi22++ {
                        var v23 int8 = v6
                        v5 -= 0
                        v6 -= v23
                        _ = fi22
                        _ = v23
                    }
                    var v24 string = (("hello" + "\n\t") + "hello")
                    for fi25 := 0; fi25 < 18; fi25++ {
                        if (!(-657 < -2147483648)) {
                            if true {
                                if false {
                                    var v26 int = (fi25 | fi25)
                                    v14 = 113
                                    v5 = v5
                                    _ = v26
                                } else {
                                    if ((!((1e10 * 3.14159) <= (-2.25 - (((0.0 + 1.5) + -2.25) + ((1.5 + 0.1) * 1e10))))) || (142 > 655)) {
                                        var v27 uint8 = v14
                                        v6 += v1
                                        _ = v27
                                    } else {
                                        var v28 uint64 = 0
                                        _ = v28
                                    }
                                }
                            }
                            v29 := GenS3{F0: -2147483648, F1: false, F2: 4294967295}
                            v30 := GenS1{F0: (-(fi25)), F1: 2147483647, F2: 2}
                            _ = v29
                            _ = v30
                        } else {
                            for fi31 := 0; fi31 < 3; fi31++ {
                                if (32767 >= 685) {
                                    v24 = v24
                                    v6 = 127
                                    v24 = "foo bar"
                                } else {
                                    fmt.Println(v1, fi25)
                                }
                                _ = fi31
                            }
                            v32 := GenS2{F0: 0, F1: 0}
                            _ = v32
                        }
                        v5 -= (v14 & v14)
                        _ = fi25
                    }
                    _ = v24
                }
                _ = len(v2)
            }
        default:
            switch "hello" {
            case "hello":
                v33 := GenS3{F0: -86, F1: false, F2: 6}
                _ = v33
            }
        }
        var v34 int8 = (^(v6))
        _ = v34
    } else {
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        v5 += 2
        if true {
            v1 = 68
            switch ("\n\t" + "\n\t") {
            case "hello":
                switch fi35 {
                case 4294967295:
                    v6 = v6
                case 32768:
                    for fi36 := 0; fi36 < 1; fi36++ {
                        v1 -= (v1 ^ v1)
                        switch 128 {
                        case 128:
                            if ((!(true || (18446744073709551615 != 1))) || (1 > 65535)) {
                                var v37 int8 = v1
                                _ = v37
                            } else {
                                v6 = 1
                            }
                        case 9223372036854775807:
                            fmt.Println(fi36, v6)
                        case -436:
                            v38 := GenS2{F0: 18446744073709551615, F1: 0}
                            _ = v38
                        default:
                            var v39 int64 = -1
                            _ = v39
                        }
                        _ = fi36
                    }
                case 32767:
                    if true {
                        var v40 uint8 = v14
                        fmt.Println(v14)
                        v41 := GenS2{F0: 42, F1: 191}
                        _ = v40
                        _ = v41
                    } else {
                        fmt.Println(v1, v14)
                        v5 += v14
                    }
                default:
                    for fi42 := 0; fi42 < 17; fi42++ {
                        fmt.Println(v6, fi35, fi35)
                        var v43 int32 = 42
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                        if ((v43 + v43) > v43) {
                            _ = len(v2)
                            fmt.Println(v14)
                            var v44 float64 = 1e10
                            _ = v44
                        } else {
                            var v45 int16 = -52
                            v1 += (v6 - v6)
                            _ = v45
                        }
                        _ = fi42
                        _ = v43
                    }
                }
            case "foo bar":
                var v46 bool = (v6 <= v1)
                _ = v46
            case "\n\t":
                v47 := GenS3{F0: 896, F1: false, F2: 2}
                _ = v47
            default:
                v5 += 2
            }
        } else {
            for fi48 := 0; fi48 < 9; fi48++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v6, v14, v1)
                v49 := GenS3{F0: int32(v14), F1: (false || true), F2: 4294967295}
                _ = fi48
                _ = v49
            }
            for fi50 := 0; fi50 < 7; fi50++ {
                var v51 uint16 = 2
                if len(v2) > 0 {
                    v2[0] = ("" + "foo bar")
                }
                _ = fi50
                _ = v51
            }
            if false {
                if false {
                    fmt.Println(fi35, v14, v1)
                    for fi52 := 0; fi52 < 14; fi52++ {
                        v4 = append(v4, "\n\t")
                        var v53 int = fi35
                        if (((65535 <= 601) && (!false)) && (2 < 641)) {
                            switch fi52 {
                            case 128:
                                _ = len(v2)
                                fmt.Println(fi52, v6)
                                v1 = v6
                            case -119:
                                var v54 uint8 = v14
                                if (!(!false)) {
                                    var v55 int64 = -9223372036854775808
                                    v56 := GenS1{F0: (fi35 | -958), F1: 4294967295, F2: 255}
                                    v53 += fi35
                                    v55 = v55
                                    v5 -= v54
                                    v14 += 35
                                    v14 = v5
                                    v57 := GenS4{F0: 1, F1: (v55 - v55), F2: 18446744073709551615}
                                    _ = len(v4)
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                }
                                _ = v54
                            default:
                                switch "" {
                                case "a":
                                    var v58 int32 = 278
                                    _ = v58
                                case "hello":
                                    _ = len(v3)
                                default:
                                    var v59 int64 = 9223372036854775807
                                    _ = v59
                                }
                            }
                        } else {
                            fmt.Println(v6, v14, v1)
                            v6 += v6
                        }
                        _ = fi52
                        _ = v53
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = fi35
                    }
                }
                var v60 int32 = 1
                if true {
                    if (!false) {
                        var v61 float32 = ((1e10 - ((0.1 / (-2.25 / 1.5)) - 1.5)) / ((1e10 - 3.14159) / float32(-2.25)))
                        v6 += (v6 | v6)
                        v4 = append(v4, "")
                        _ = v61
                    } else {
                        v62 := GenS5{F0: fi35, F1: 0, F2: ("foo bar" + "a")}
                        for fi63 := 0; fi63 < 17; fi63++ {
                            v64 := GenS4{F0: 65535, F1: 9223372036854775807, F2: 625}
                            var v65 uint8 = (v14 + 42)
                            switch (fi63 << ((18446744073709551615) & 63)) {
                            case 32767:
                                v65 = (v65 * v14)
                            case 128:
                                if len(v3) > 0 {
                                    v3[0] = fi63
                                }
                            case -9223372036854775808:
                                var v66 uint16 = 42
                                v67 := GenS4{F0: uint16(v14), F1: 1, F2: 18446744073709551615}
                                _ = v66
                                _ = v67
                            }
                            fmt.Println(fi35, v5)
                            fmt.Println(v1, v14)
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                        switch (fi35 & fi35) {
                        case -477:
                            var v68 int8 = (v6 >> ((0) & 63))
                            _ = v68
                        default:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        _ = v62
                    }
                }
                _ = v60
            }
        }
        _ = fi35
    }
    switch "\n\t" {
    case "\n\t":
        _ = len(v3)
    default:
        if false {
            if true {
                var v69 string = "foo bar"
                fmt.Println(v1, v14)
                _ = v69
            } else {
                switch 4294967295 {
                case 4294967295:
                    if len(v3) > 0 {
                        v3[0] = -1
                    }
                }
            }
        }
    }
    v2 = append(v2, "foo bar")
    if true {
        v1 += (v1 - v1)
    }
    v70 := GenS2{F0: 248, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v70
    _ = v8
}
