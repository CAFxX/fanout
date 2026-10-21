package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 int16; F2 int32 }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 uint64; F1 int; F2 string }
type GenS9 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint64; F1 uint32; F2 int64 }
func genFn0(p0 uint16) int32 {
    p0 = (p0 >> ((2) & 63))
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 2}
        var v7 uint64 = 2
        v8 := GenS1{F0: 127}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    _ = p0
    return -2147483648
}

func main() {
    var v1 float64 = ((0.0 * 1e10) * 1e10)
    var v2 bool = (!(false || (0 == 128)))
    var v3 uint8 = 51
    v4 := []int{128}
    v1 -= (v1 * v1)
    v9 := GenS1{F0: 2}
    if true {
        v1 = v1
        if v2 {
            for fi10 := 0; fi10 < 2; fi10++ {
                if true {
                    var v11 uint8 = (v3 % (((v3 << ((745) & 63))) | 1))
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 13; fi12++ {
                        v3 += (v3 / ((v3) | 1))
                        v13 := GenS2{F0: (v3 + (v3 ^ 2))}
                        v4 = append(v4, (^(fi10)))
                        _ = fi12
                        _ = v13
                    }
                    v14 := GenS2{F0: v3}
                    _ = v14
                }
                fmt.Println(v3, fi10, fi10)
                _ = fi10
            }
            if len(v4) > 0 {
                v4[0] = 9223372036854775807
            }
        } else {
            if (true || v2) {
                v15 := GenS3{F0: 0}
                v2 = (9223372036854775807 <= 451)
                v16 := GenS3{F0: 255}
                _ = v15
                _ = v16
            } else {
                var v17 uint = 255
                var v18 bool = v2
                _ = v17
                _ = v18
            }
            fmt.Println(v1, v2, v1)
        }
    } else {
        var v19 int8 = -1
        v2 = ((v1 - v1) == 0.0)
        v2 = v2
        _ = v19
    }
    v4 = append(v4, -1)
    v3 = (v3 | v3)
    v3 += v3
    for fi20 := 0; fi20 < 13; fi20++ {
        if v2 {
            v21 := GenS4{F0: 4294967295, F1: -32768, F2: genFn0(32767)}
            _ = v21
        } else {
            v3 = v3
            v22 := GenS5{F0: fi20, F1: uint64(v3)}
            _ = v22
        }
        v23 := GenS3{F0: 1}
        for fi24 := 0; fi24 < 15; fi24++ {
            switch (fi24 - fi20) {
            case -240:
                v3 += (v3 << ((0) & 63))
            }
            v25 := GenS4{F0: 4294967295, F1: 1, F2: 846}
            v3 -= (v3 % ((v3) | 1))
            _ = fi24
            _ = v25
        }
        _ = fi20
        _ = v23
    }
    var v26 uint = 4294967295
    if (1 != 0) {
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 int8 = 127
            _ = fi27
            _ = v28
        }
    } else {
        _ = len(v4)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        for fi30 := 0; fi30 < 19; fi30++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch (fi29 >> (((v26 + v26)) & 63)) {
            case 4294967295:
                v31 := GenS1{F0: v3}
                _ = v31
            case 128:
                v4 = append(v4, 490)
            }
            switch ("hello" + "\n\t") {
            case "":
                var v32 uint64 = uint64(v26)
                _ = v32
            default:
                fmt.Println(v26, v1)
            }
            v1 = 0.0
            _ = fi30
        }
        v33 := GenS6{F0: 140, F1: 42}
        v1 -= ((float64(1e10) / ((v1 + (0.1 - ((1e10 - float64((1.5 * 1e10))) + (float64(3.14159) + -2.25)))) - 0.1)) * (0.0 - 0.1))
        _ = fi29
        _ = v33
    }
    v26 += v26
    if v2 {
        if v2 {
            for fi34 := 0; fi34 < 17; fi34++ {
                if (754 == 9223372036854775807) {
                    v1 = v1
                } else {
                    switch "hello" {
                    case "hello":
                        if ("foo bar" != "foo bar") {
                            switch "hello" {
                            case "":
                                fmt.Println(v2)
                            default:
                                var v35 string = ""
                                _ = v35
                            }
                            for fi36 := 0; fi36 < 17; fi36++ {
                                if v2 {
                                    v37 := GenS7{F0: 125, F1: 396}
                                    v1 = (v1 * v1)
                                    v1 += v1
                                    _ = v37
                                }
                                _ = fi36
                            }
                        } else {
                            v1 += 1.5
                            for fi38 := 0; fi38 < 4; fi38++ {
                                v39 := GenS1{F0: (v3 * 87)}
                                fmt.Println(fi38, v26)
                                _ = fi38
                                _ = v39
                            }
                            for fi40 := 0; fi40 < 15; fi40++ {
                                v4 = append(v4, 2147483647)
                                fmt.Println(fi40, v3)
                                _ = fi40
                            }
                        }
                    case "\n\t":
                        if ((2 == 2) && v2) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v26 = 965
                        } else {
                            v1 = (-2.25 * v1)
                        }
                    case "a":
                        if (42 >= 728) {
                            if ((genFn0(383) == 42) || v2) {
                                for fi41 := 0; fi41 < 12; fi41++ {
                                    var v42 uint32 = 42
                                    if len(v4) > 0 {
                                        v4[0] = int(v42)
                                    }
                                    var v43 int8 = 0
                                    switch "foo bar" {
                                    case "\n\t":
                                        switch ("foo bar" + ("\n\t" + (("a" + "foo bar") + ("\n\t" + "a")))) {
                                        case "foo bar":
                                            if len(v4) > 0 {
                                                v4[0] = 526
                                            }
                                        default:
                                            var v44 int = fi41
                                            _ = v44
                                        }
                                    default:
                                        if v2 {
                                            fmt.Println(v3, v42, v2)
                                            if (v1 == -2.25) {
                                                var v45 int32 = -1
                                                var v46 int64 = 959
                                                _ = v45
                                                _ = v46
                                            } else {
                                                var v47 int16 = 32767
                                                _ = v47
                                            }
                                        }
                                    }
                                    if true {
                                        v43 += (v43 - v43)
                                        v48 := GenS7{F0: v3, F1: uint64(v42)}
                                        v49 := GenS7{F0: 42, F1: 2}
                                        _ = v48
                                        _ = v49
                                    } else {
                                        fmt.Println(fi34, v43, v1)
                                    }
                                    _ = len(v4)
                                    var v50 int8 = 127
                                    v50 += v50
                                    v51 := GenS8{F0: 2, F1: (fi41 - 99), F2: "foo bar"}
                                    _ = fi41
                                    _ = v42
                                    _ = v43
                                    _ = v50
                                    _ = v51
                                }
                                v26 = 42
                                v52 := GenS9{F0: 1, F1: 1, F2: 2}
                                _ = v52
                            }
                            for fi53 := 0; fi53 < 2; fi53++ {
                                v54 := GenS10{F0: 255}
                                v55 := GenS3{F0: 255}
                                _ = len(v4)
                                _ = fi53
                                _ = v54
                                _ = v55
                            }
                            if (((v26 + v26) != (^(v26))) && v2) {
                                v26 = v26
                                switch (fi34 % ((fi34) | 1)) {
                                case 0:
                                    var v56 int8 = 42
                                    _ = v56
                                default:
                                    var v57 uint8 = v3
                                    _ = v57
                                }
                            } else {
                                if (1 > 2147483648) {
                                    v58 := GenS7{F0: v3, F1: 954}
                                    _ = v58
                                }
                                for fi59 := 0; fi59 < 15; fi59++ {
                                    if (219 == (^(v3))) {
                                        switch (^(fi34)) {
                                        case 0:
                                            v26 = (v26 & 759)
                                        case -69:
                                            switch ("a" + "a") {
                                            case "a":
                                                var v60 int32 = -1
                                                _ = v60
                                            case "":
                                                switch fi59 {
                                                case 329:
                                                    var v61 int16 = 1000
                                                    _ = v61
                                                }
                                            default:
                                                fmt.Println(v1)
                                            }
                                        case 4294967295:
                                            var v62 uint32 = 802
                                            _ = v62
                                        default:
                                            var v63 uint8 = v3
                                            _ = v63
                                        }
                                        if (!true) {
                                            if v2 {
                                                fmt.Println(v2, v1)
                                            }
                                            v64 := GenS3{F0: 32767}
                                            _ = v64
                                        }
                                        fmt.Println(fi59)
                                        switch ("hello" + "a") {
                                        case "a":
                                            switch (fi59 % (((fi34 + (fi34 / ((fi34) | 1)))) | 1)) {
                                            case 0:
                                                if len(v4) > 0 {
                                                    v4[0] = fi34
                                                }
                                            case 1:
                                                if v2 {
                                                    if v2 {
                                                        var v65 int = fi59
                                                        v3 = 14
                                                        _ = v65
                                                    }
                                                } else {
                                                    var v66 string = ("hello" + "")
                                                    _ = v66
                                                }
                                            case -2147483648:
                                                switch "foo bar" {
                                                case "foo bar":
                                                    if v2 {
                                                        if (false && false) {
                                                            v67 := GenS11{F0: 32767}
                                                            _ = v67
                                                        }
                                                    }
                                                }
                                                switch "foo bar" {
                                                case "":
                                                    var v68 int32 = 408
                                                    _ = v68
                                                case "foo bar":
                                                    var v69 float32 = float32(v1)
                                                    _ = v69
                                                case "a":
                                                    var v70 uint16 = 0
                                                    _ = v70
                                                default:
                                                    v2 = (((((-2.25 + 1.5) + 0.1) >= 1.5) && v2) || ((0.1 + 1e10) >= -2.25))
                                                }
                                                if (-2147483648 < genFn0(42)) {
                                                    var v71 uint32 = 897
                                                    _ = v71
                                                }
                                            }
                                        case "\n\t":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        case "":
                                            fmt.Println(v3)
                                        default:
                                            v72 := GenS7{F0: (v3 % ((1) | 1)), F1: 1}
                                            _ = v72
                                        }
                                    } else {
                                        v3 += (v3 | v3)
                                        v3 += 6
                                    }
                                    _ = fi59
                                }
                                for fi73 := 0; fi73 < 2; fi73++ {
                                    fmt.Println(v3, v1, v1)
                                    var v74 int16 = 0
                                    _ = fi73
                                    _ = v74
                                }
                            }
                        }
                    default:
                        switch (("a" + ("\n\t" + "\n\t")) + ("foo bar" + "foo bar")) {
                        case "hello":
                            switch "\n\t" {
                            case "a":
                                switch fi34 {
                                case 42:
                                    var v75 float64 = 1.5
                                    _ = v75
                                case 127:
                                    var v76 int16 = 32767
                                    _ = v76
                                }
                            case "foo bar":
                                _ = len(v4)
                            }
                        case "":
                            fmt.Println(v26, v2)
                        }
                    }
                    for fi77 := 0; fi77 < 6; fi77++ {
                        var v78 uint16 = 32767
                        var v79 uint16 = v78
                        _ = fi77
                        _ = v78
                        _ = v79
                    }
                }
                fmt.Println(fi34)
                _ = fi34
            }
            if (895 == 147) {
                var v80 int32 = genFn0(434)
                _ = v80
            }
        } else {
            v81 := GenS11{F0: 808}
            switch 9223372036854775807 {
            case 127:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 9223372036854775807:
                v2 = (630 > 561)
            default:
                v82 := GenS12{F0: 4294967295, F1: 925, F2: -873}
                _ = v82
            }
            _ = v81
        }
        var v83 uint32 = 329
        _ = v83
    }
    var v84 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v84
    _ = v9
}
