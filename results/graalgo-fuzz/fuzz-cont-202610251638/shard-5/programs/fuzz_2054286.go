package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int64; F2 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 int; F2 int8 }
type GenS8 struct { F0 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 729
    var v3 uint8 = 128
    var v4 uint32 = 254
    var v5 uint32 = (v4 + v4)
    var v6 uint64 = 0
    var v7 int64 = 0
    v8 := GenS1{F0: (-(v7)), F1: 0}
    for fi9 := 0; fi9 < 13; fi9++ {
        switch ("foo bar" + "hello") {
        case "hello":
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := GenS2{F0: (v7 - 9223372036854775807)}
                v12 := []int{fi10, fi9, 128, fi10}
                v6 = 4294967295
                fmt.Println(fi10, v2)
                v4 += v4
                _ = fi10
                _ = v11
                _ = v12
            }
        case "a":
            if false {
                v7 = v7
                v1["\n\t"] = (v5 / ((v4) | 1))
                var v13 uint8 = (v3 >> ((65535) & 63))
                _ = v13
            }
        }
        _ = fi9
    }
    if (true || false) {
        switch -2147483648 {
        case 0:
            v14 := GenS3{F0: 1, F1: 458, F2: (v3 + uint8(v4))}
            _ = v14
        case 1:
            for fi15 := 0; fi15 < 3; fi15++ {
                for fi16 := 0; fi16 < 3; fi16++ {
                    v6 = (v6 >> ((489) & 63))
                    _ = fi16
                }
                _ = fi15
            }
        }
        if true {
            v17 := GenS1{F0: (v7 | (v7 << ((0) & 63))), F1: 42}
            var v18 int8 = 96
            var v19 int16 = 42
            _ = v17
            _ = v18
            _ = v19
        } else {
            var v20 uint8 = (v3 - v3)
            _ = v20
        }
        for fi21 := 0; fi21 < 1; fi21++ {
            v6 = (v6 ^ v6)
            if false {
                if (false && (357 < 543)) {
                    var v22 int32 = 733
                    _ = v22
                } else {
                    fmt.Println(v2)
                }
                v4 -= v5
            } else {
                fmt.Println(v6, v2)
                if ((v4 % ((173) | 1)) <= 607) {
                    v23 := GenS3{F0: 255, F1: 9223372036854775807, F2: (^(v3))}
                    _ = v23
                }
                _ = len(v1)
            }
            var v24 int64 = v7
            _ = fi21
            _ = v24
        }
    }
    v2 = (v2 / ((v2) | 1))
    v25 := GenS2{F0: (v7 * -1)}
    v26 := []int{}
    if ((!false) || (!((v6 & (v6 | 1)) < v6))) {
        switch 128 {
        case 1:
            switch 729 {
            case 452:
                var v27 bool = (!(float64(1e10) <= 1.5))
                _ = v27
            default:
                var v28 int32 = -2147483648
                _ = v28
            }
        case 32768:
            if (!false) {
                var v29 int64 = -1
                for fi30 := 0; fi30 < 16; fi30++ {
                    switch "foo bar" {
                    case "hello":
                        switch 186 {
                        case 32768:
                            for fi31 := 0; fi31 < 13; fi31++ {
                                var v32 float64 = 3.14159
                                switch -2147483648 {
                                case 42:
                                    if len(v26) > 0 {
                                        v26[0] = fi30
                                    }
                                default:
                                    var v33 uint8 = (v3 >> ((187) & 63))
                                    _ = v33
                                }
                                v2 = (v2 | v2)
                                _ = fi31
                                _ = v32
                            }
                        case 251:
                            if false {
                                var v34 float64 = -2.25
                                _ = v34
                            }
                        }
                    case "":
                        v3 = (^(v3))
                    default:
                        var v35 float64 = 0.0
                        _ = v35
                    }
                    var v36 float64 = (0.0 + (((-2.25 - 1.5) + float64((0.1 / 1.0))) * 0.1))
                    v37 := GenS4{F0: 32767}
                    _ = fi30
                    _ = v36
                    _ = v37
                }
                switch "\n\t" {
                case "a":
                    for fi38 := 0; fi38 < 6; fi38++ {
                        for fi39 := 0; fi39 < 19; fi39++ {
                            v40 := GenS4{F0: 122}
                            v41 := GenS2{F0: (v29 * -100)}
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                        if (!false) {
                            for fi42 := 0; fi42 < 20; fi42++ {
                                if len(v26) > 0 {
                                    v26[0] = (fi38 * fi42)
                                }
                                _ = fi42
                            }
                            fmt.Println(v4)
                        }
                        fmt.Println(v5)
                        _ = fi38
                    }
                case "foo bar":
                    var v43 bool = false
                    _ = v43
                default:
                    delete(v1, ("" + ("a" + (((("a" + ("" + "hello")) + "foo bar") + ("hello" + "")) + "a"))))
                }
                _ = v29
            }
        default:
            v2 = (v2 ^ (v2 ^ (v2 % ((255) | 1))))
        }
        var v44 uint8 = v3
        _ = v44
    } else {
        if true {
            _ = len(v26)
        }
    }
    v7 = (v7 * v7)
    for fi45 := 0; fi45 < 18; fi45++ {
        switch 4294967295 {
        case 1:
            fmt.Println(v4, v7)
        case 2147483647:
            if (1.5 <= 1.5) {
                fmt.Println(fi45, v6)
            }
        case 9223372036854775807:
            var v46 int16 = 32767
            _ = v46
        default:
            fmt.Println(v2, v3)
        }
        _ = fi45
    }
    v6 = (v6 / (((^(v6))) | 1))
    if ("foo bar" < "a") {
        switch -9223372036854775808 {
        case -2147483648:
            v47 := GenS5{F0: (v6 / ((18446744073709551615) | 1)), F1: (v5 | (v5 / ((v4) | 1)))}
            _ = v47
        case 492:
            var v48 uint32 = (v5 << ((1) & 63))
            _ = v48
        }
    } else {
        if true {
            if false {
                v49 := GenS3{F0: 400, F1: (-(v7)), F2: 127}
                _ = v49
            } else {
                if (0 < 42) {
                    if false {
                        var v50 bool = (-128 != 80)
                        if v50 {
                            switch 165 {
                            case 4294967295:
                                v51 := GenS6{F0: (v7 * -9223372036854775808)}
                                _ = v51
                            default:
                                switch "foo bar" {
                                case "foo bar":
                                    fmt.Println(v4)
                                default:
                                    switch "hello" {
                                    case "foo bar":
                                        if v50 {
                                            var v52 int16 = 32767
                                            if v50 {
                                                switch -2147483648 {
                                                case -365:
                                                    if false {
                                                        for fi53 := 0; fi53 < 1; fi53++ {
                                                            if false {
                                                                v7 = (v7 << ((981) & 63))
                                                                v54 := GenS7{F0: (v7 - 1), F1: (fi53 | fi53), F2: -128}
                                                                _ = v54
                                                            }
                                                            var v55 uint32 = (v4 ^ v5)
                                                            _ = fi53
                                                            _ = v55
                                                        }
                                                    }
                                                default:
                                                    if (0 < v3) {
                                                        switch 0 {
                                                        case 42:
                                                            v52 += 1000
                                                        case 32768:
                                                            var v56 uint32 = v5
                                                            _ = v56
                                                        case 9223372036854775807:
                                                            var v57 uint16 = 834
                                                            _ = v57
                                                        default:
                                                            for fi58 := 0; fi58 < 6; fi58++ {
                                                                if false {
                                                                    var v59 int32 = -850
                                                                    v3 -= (v3 * v3)
                                                                    var v60 int = fi58
                                                                    _ = v59
                                                                    _ = v60
                                                                } else {
                                                                    v61 := GenS7{F0: v7, F1: (fi58 % (((fi58 * (^(fi58)))) | 1)), F2: -87}
                                                                    var v62 int64 = (v7 & int64(v52))
                                                                    _ = v61
                                                                    _ = v62
                                                                }
                                                                v63 := v1[("" + "\n\t")]
                                                                v64 := GenS6{F0: (v7 << ((uint(v6)) & 63))}
                                                                _ = fi58
                                                                _ = v63
                                                                _ = v64
                                                            }
                                                        }
                                                    } else {
                                                        var v65 int32 = 600
                                                        var v66 bool = v50
                                                        fmt.Println(v66, v5, v5)
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                }
                                            } else {
                                                var v67 int16 = v52
                                                var v68 int = -1
                                                _ = v67
                                                _ = v68
                                            }
                                            _ = v52
                                        }
                                    default:
                                        switch "\n\t" {
                                        case "hello":
                                            var v69 int32 = 42
                                            _ = v69
                                        case "\n\t":
                                            v6 = (v6 ^ (v6 + v6))
                                        case "":
                                            v5 = v4
                                        default:
                                            v70 := GenS4{F0: 42}
                                            _ = v70
                                        }
                                    }
                                }
                            }
                            var v71 uint = uint(v3)
                            _ = v71
                        }
                        var v72 float64 = 3.14159
                        _ = v50
                        _ = v72
                    }
                } else {
                    v73 := GenS5{F0: (^(v6)), F1: v5}
                    fmt.Println(v7, v3, v6)
                    v74 := GenS3{F0: 373, F1: v7, F2: 174}
                    _ = v73
                    _ = v74
                }
                fmt.Println(v6, v7)
            }
            fmt.Println(v4)
            if (true && true) {
                v3 -= (v3 << ((42) & 63))
                for fi75 := 0; fi75 < 14; fi75++ {
                    _ = len(v26)
                    v76 := GenS1{F0: (v7 >> ((0) & 63)), F1: 127}
                    _ = fi75
                    _ = v76
                }
                v77 := GenS5{F0: v6, F1: (^(v5))}
                _ = v77
            }
        }
    }
    v78 := GenS4{F0: 16}
    for fi79 := 0; fi79 < 10; fi79++ {
        for fi80 := 0; fi80 < 4; fi80++ {
            var v81 float64 = 1e10
            _ = fi80
            _ = v81
        }
        _ = fi79
    }
    v2 = v2
    v82 := GenS8{F0: "hello"}
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v78
    _ = v8
    _ = v82
}
