package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 string; F1 int16; F2 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 int; F1 int16 }
func genFn0() uint {
    for fi6 := 0; fi6 < 6; fi6++ {
        if true {
            v7 := GenS1{F0: 795}
            fmt.Println(fi6)
            for fi8 := 0; fi8 < 16; fi8++ {
                if (0 < 1) {
                    var v9 uint64 = 18446744073709551615
                    v9 -= v9
                    _ = v9
                }
                v10 := GenS2{F0: -305, F1: 20}
                if ((0.0 * -2.25) >= 0.1) {
                    fmt.Println(fi8, fi6)
                    switch ("a" + "foo bar") {
                    case "a":
                        v11 := []int{}
                        _ = v11
                    default:
                        var v12 int = (-(fi6))
                        _ = v12
                    }
                }
                _ = fi8
                _ = v10
            }
            _ = v7
        } else {
            v13 := GenS3{F0: ("" + ("a" + ("a" + "a"))), F1: 1000, F2: 9223372036854775807}
            var v14 int16 = -1
            _ = v13
            _ = v14
        }
        _ = fi6
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 float32 = 0.0
        switch (("hello" + "foo bar") + "\n\t") {
        case "hello":
            switch "foo bar" {
            case "":
                v16 = (((((v16 * (v16 / ((1.5 / v16) - v16))) + (float32((3.14159 / 3.14159)) * 0.1)) - ((-2.25 - 1e10) + v16)) - (v16 + 0.0)) * 1e10)
            case "a":
                switch ("foo bar" + ("a" + ("\n\t" + "hello"))) {
                case "\n\t":
                    v16 -= v16
                case "foo bar":
                    for fi17 := 0; fi17 < 9; fi17++ {
                        var v18 uint8 = 0
                        v18 = 0
                        _ = fi17
                        _ = v18
                    }
                case "hello":
                    if true {
                        switch ("foo bar" + ("" + "foo bar")) {
                        case "":
                            var v19 bool = (65535 != 104)
                            _ = v19
                        default:
                            fmt.Println(v16)
                        }
                        v16 = 1e10
                    }
                default:
                    for fi20 := 0; fi20 < 13; fi20++ {
                        var v21 float32 = (1.5 + v16)
                        v22 := GenS4{F0: 127}
                        v23 := []int{(-(fi20)), fi15, (^(fi15)), (fi20 / ((127) | 1))}
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                }
            default:
                var v24 string = "\n\t"
                _ = v24
            }
        case "\n\t":
            v16 = v16
        case "foo bar":
            v16 = v16
        default:
            for fi25 := 0; fi25 < 5; fi25++ {
                if true {
                    if true {
                        v26 := []int{-2147483648, (fi25 >> ((42) & 63)), fi25, fi25}
                        _ = v26
                    }
                    if (720 > 2) {
                        var v27 int8 = -128
                        if (!false) {
                            switch fi25 {
                            case 42:
                                var v28 uint = 572
                                _ = v28
                            case -9223372036854775808:
                                switch "\n\t" {
                                case "foo bar":
                                    v16 = v16
                                default:
                                    v27 += -1
                                }
                            case -672:
                                var v29 int = fi25
                                _ = v29
                            }
                            fmt.Println(v16, fi15)
                        }
                        _ = v27
                    } else {
                        var v30 int16 = 1000
                        var v31 uint = 1
                        var v32 uint16 = 255
                        fmt.Println(v16, v16, v31)
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    var v33 uint64 = 1
                    _ = v33
                }
                _ = fi25
            }
        }
        _ = fi15
        _ = v16
    }
    return 0
}
func genFn1(p0 int, p1 int64) int32 {
    fmt.Println(p0, p0)
    var v34 float32 = (-2.25 - ((0.1 + 0.1) * -2.25))
    switch (p0 ^ p0) {
    case -822:
        for fi35 := 0; fi35 < 19; fi35++ {
            fmt.Println(fi35)
            v36 := GenS4{F0: -128}
            p0 = 42
            _ = fi35
            _ = v36
        }
    case 4294967295:
        if true {
            fmt.Println(p1, p1, p1)
        } else {
            v37 := GenS5{F0: 4294967295, F1: 42}
            if true {
                p0 = 492
                for fi38 := 0; fi38 < 3; fi38++ {
                    v39 := GenS1{F0: 0}
                    _ = fi38
                    _ = v39
                }
            }
            var v40 int16 = 32767
            _ = v37
            _ = v40
        }
    default:
        var v41 float64 = 0.0
        _ = v41
    }
    for fi42 := 0; fi42 < 16; fi42++ {
        for fi43 := 0; fi43 < 16; fi43++ {
            var v44 uint64 = 546
            p0 = p0
            _ = fi43
            _ = v44
        }
        _ = fi42
    }
    _ = p0
    _ = p1
    _ = v34
    return 2147483647
}

func main() {
    v1 := []uint32{537, 484, 42}
    var v2 uint16 = 32767
    var v3 bool = (165 != 359)
    v4 := make(map[string]string)
    var v5 uint64 = 714
    v4[""] = "a"
    fmt.Println(v3, v3)
    if v3 {
        switch 996 {
        case 0:
            v45 := GenS2{F0: -2147483648, F1: -1}
            _ = v45
        case -1:
            for fi46 := 0; fi46 < 14; fi46++ {
                v47 := GenS5{F0: 2147483648, F1: (v5 * (v5 & v5))}
                _ = fi46
                _ = v47
            }
        case -746:
            v5 = (v5 >> ((18446744073709551615) & 63))
        default:
            if (((1.5 / ((0.0 - 1.5) - 1e10)) - 0.1) >= (0.1 / (0.0 - 3.14159))) {
                v5 = 365
                for fi48 := 0; fi48 < 6; fi48++ {
                    if (!v3) {
                        switch "foo bar" {
                        case "hello":
                            if (42 > 0) {
                                v2 -= 42
                            } else {
                                switch -995 {
                                case 0:
                                    _ = len(v4)
                                default:
                                    for fi49 := 0; fi49 < 17; fi49++ {
                                        v2 = 65535
                                        _ = fi49
                                    }
                                }
                                for fi50 := 0; fi50 < 12; fi50++ {
                                    v51 := GenS5{F0: 2, F1: uint64(v2)}
                                    _ = fi50
                                    _ = v51
                                }
                                var v52 int8 = -62
                                _ = v52
                            }
                        case "foo bar":
                            fmt.Println(fi48)
                        case "\n\t":
                            fmt.Println(v5, v3, v3)
                        default:
                            var v53 uint64 = (v5 | (v5 + 18446744073709551615))
                            _ = v53
                        }
                        switch 127 {
                        case 616:
                            var v54 float32 = float32((-2.25 + 1.5))
                            _ = v54
                        case -18:
                            var v55 int64 = 9223372036854775807
                            _ = v55
                        case 128:
                            if (127 < 2) {
                                var v56 uint64 = (v5 - (v5 ^ v5))
                                for fi57 := 0; fi57 < 14; fi57++ {
                                    var v58 int = fi48
                                    switch v58 {
                                    case 9223372036854775807:
                                        v59 := GenS6{F0: (v58 << ((4294967295) & 63)), F1: 1000}
                                        _ = v59
                                    }
                                    fmt.Println(v58, v5)
                                    _ = fi57
                                    _ = v58
                                }
                                _ = v56
                            } else {
                                v60 := GenS6{F0: 42, F1: -32768}
                                var v61 uint = 688
                                _ = v60
                                _ = v61
                            }
                        default:
                            for fi62 := 0; fi62 < 11; fi62++ {
                                var v63 uint64 = (^(v5))
                                switch fi48 {
                                case 4294967295:
                                    _ = len(v4)
                                case -115:
                                    v64 := GenS3{F0: "hello", F1: 42, F2: -1}
                                    _ = v64
                                case -647:
                                    v3 = ((fi48 ^ 127) > 766)
                                }
                                _ = fi62
                                _ = v63
                            }
                        }
                    } else {
                        fmt.Println(fi48)
                    }
                    _ = fi48
                }
            } else {
                v65 := GenS3{F0: "", F1: 42, F2: -1}
                var v66 int64 = 0
                var v67 int = -2147483648
                _ = v65
                _ = v66
                _ = v67
            }
        }
        var v68 int32 = 1
        _ = v68
    } else {
        _ = len(v1)
        for fi69 := 0; fi69 < 18; fi69++ {
            for fi70 := 0; fi70 < 2; fi70++ {
                var v71 bool = v3
                var v72 uint64 = 4294967295
                if v71 {
                    v73 := GenS3{F0: (("\n\t" + "hello") + "foo bar"), F1: 1, F2: 647}
                    var v74 int32 = 230
                    v75 := GenS5{F0: 42, F1: 2}
                    _ = v73
                    _ = v74
                    _ = v75
                }
                _ = fi70
                _ = v71
                _ = v72
            }
            fmt.Println(v2, v2, v5)
            _ = len(v4)
            _ = fi69
        }
    }
    if v3 {
        fmt.Println(v2, v5)
        fmt.Println(v5, v3)
    } else {
        var v76 uint16 = (v2 ^ (v2 * v2))
        _ = len(v1)
        _ = v76
    }
    for fi77 := 0; fi77 < 5; fi77++ {
        var v78 float32 = 1e10
        _ = fi77
        _ = v78
    }
    v3 = ((v5 << ((180) & 63)) < (v5 + (v5 % ((v5) | 1))))
    var v79 int64 = 511
    v80 := v4["hello"]
    var v81 float64 = ((0.0 + 1.5) - (float64((1e10 - -2.25)) / (0.1 + 1.5)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v79
    _ = v80
    _ = v81
}
