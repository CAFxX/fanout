package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint32; F2 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 int32; F1 uint16; F2 int64 }
func genFn0(p0 bool, p1 int32) int {
    p0 = true
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 128
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint64 = 75
    var v3 uint = 527
    var v4 int = -9223372036854775808
    v5 := make(map[uint32]bool)
    v4 = (v4 ^ (v4 - v4))
    v2 -= v2
    v1 = v1
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 uint = v3
            _ = len(v5)
            switch "" {
            case "foo bar":
                if true {
                    switch (((("" + ("foo bar" + "\n\t")) + ("foo bar" + "a")) + ("a" + "hello")) + "hello") {
                    case "\n\t":
                        v9 := GenS1{F0: 42}
                        _ = v9
                    case "foo bar":
                        v10 := GenS1{F0: 2}
                        _ = v10
                    case "":
                        var v11 float32 = 1e10
                        _ = v11
                    default:
                        var v12 int = (^(fi6))
                        _ = v12
                    }
                    v13 := []int{genFn0(true, 439), v4, (fi6 | v4), genFn0(true, 0)}
                    var v14 string = ("\n\t" + "foo bar")
                    fmt.Println(fi7, v14)
                    var v15 int64 = (v1 >> (((v8 * v3)) & 63))
                    var v16 uint16 = 42
                    v17 := GenS2{F0: uint(v16), F1: 2, F2: true}
                    _ = v13
                    _ = v14
                    _ = v15
                    _ = v16
                    _ = v17
                } else {
                    v18 := []int{}
                    var v19 int16 = -32768
                    var v20 uint8 = 42
                    _ = v18
                    _ = v19
                    _ = v20
                }
            case "\n\t":
                var v21 uint8 = 128
                _ = v21
            case "":
                if (true && (2 > v3)) {
                    delete(v5, 42)
                    if true {
                        fmt.Println(v2, v3)
                        fmt.Println(fi6, v8)
                        v22 := GenS2{F0: v8, F1: 904, F2: (0 <= 2147483647)}
                        _ = v22
                    }
                    _ = len(v5)
                }
            }
            _ = fi7
            _ = v8
        }
        _ = len(v5)
        for fi23 := 0; fi23 < 18; fi23++ {
            v4 = genFn0(true, 2147483647)
            v24 := GenS3{F0: ("" + "a")}
            _ = fi23
            _ = v24
        }
        _ = fi6
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        v26 := GenS4{F0: (-(v4)), F1: -32768}
        for fi27 := 0; fi27 < 17; fi27++ {
            if ((1.5 * (0.0 - (-2.25 - (0.1 + 0.1)))) > (1e10 * 1e10)) {
                _ = len(v5)
                if (65535 > 888) {
                    var v28 uint32 = 4294967295
                    _ = v28
                }
                v2 += v2
                v29 := []int{}
                fmt.Println(v2, v2, fi25)
                if true {
                    switch int(v1) {
                    case -381:
                        if true {
                            _ = len(v5)
                            v2 = 907
                        } else {
                            v4 += (fi25 - genFn0((false || (!true)), -566))
                            switch "foo bar" {
                            case "hello":
                                v3 += v3
                            case "\n\t":
                                var v30 uint = v3
                                _ = v30
                            }
                        }
                    }
                    v31 := GenS1{F0: 2147483647}
                    _ = v31
                } else {
                    switch 673 {
                    case -1:
                        v32 := GenS3{F0: "foo bar"}
                        _ = v32
                    case 626:
                        v33 := GenS5{F0: ((float64(-2.25) - -2.25) != float64((float32((-2.25 * 3.14159)) * 1e10)))}
                        _ = v33
                    default:
                        v34 := GenS4{F0: fi25, F1: 32767}
                        _ = v34
                    }
                    v1 -= (-(v1))
                }
                _ = v29
            } else {
                v35 := GenS2{F0: 4294967295, F1: 860, F2: true}
                _ = v35
            }
            var v36 int64 = (v1 | -9223372036854775808)
            _ = fi27
            _ = v36
        }
        _ = fi25
        _ = v26
    }
    var v37 bool = (-1 != 42)
    for fi38 := 0; fi38 < 2; fi38++ {
        v39 := GenS3{F0: ("" + "\n\t")}
        v40 := GenS6{F0: 2, F1: -1}
        _ = fi38
        _ = v39
        _ = v40
    }
    if false {
        switch genFn0(true, 42) {
        case 921:
            v41 := GenS7{F0: (true && (0 <= -128))}
            _ = v41
        }
        delete(v5, 869)
    }
    for fi42 := 0; fi42 < 3; fi42++ {
        fmt.Println(v1, v2, v1)
        for fi43 := 0; fi43 < 11; fi43++ {
            if false {
                if true {
                    v5[42] = (!((42 != 0) && (v4 == -2147483648)))
                    fmt.Println(v2, v1, fi42)
                    v5[672] = v37
                } else {
                    var v44 int8 = 21
                    v37 = (fi42 < fi42)
                    _ = v44
                }
                switch ("a" + "hello") {
                case "a":
                    v45 := GenS5{F0: false}
                    _ = v45
                case "foo bar":
                    v3 = 890
                case "\n\t":
                    var v46 float64 = 0.0
                    _ = v46
                default:
                    if ((!v37) && v37) {
                        if (-113 < 1000) {
                            var v47 int = fi42
                            v5[4294967295] = (!true)
                            _ = v47
                        } else {
                            switch (("" + "foo bar") + "") {
                            case "\n\t":
                                v48 := []int{(fi43 + 12), (fi43 + -595), genFn0((-1 > 1), 2147483647)}
                                _ = v48
                            case "foo bar":
                                v49 := []int{genFn0(false, -1), v4}
                                _ = v49
                            case "hello":
                                v50 := []int{(fi43 << (((^(v3))) & 63)), (fi43 << ((v3) & 63)), -9223372036854775808}
                                var v51 int8 = 127
                                _ = v50
                                _ = v51
                            }
                            v4 -= v4
                            if (3.14159 > 0.0) {
                                var v52 float64 = -2.25
                                var v53 uint = 876
                                var v54 int64 = (-(v1))
                                var v55 float32 = (0.1 + 3.14159)
                                v56 := GenS8{F0: (^(v53)), F1: 100}
                                _ = v52
                                _ = v53
                                _ = v54
                                _ = v55
                                _ = v56
                            } else {
                                var v57 uint = v3
                                var v58 float64 = 0.0
                                _ = v57
                                _ = v58
                            }
                        }
                    }
                }
                v1 = (v1 ^ 196)
            }
            _ = fi43
        }
        v2 += v2
        _ = fi42
    }
    fmt.Println(v37)
    switch ((("a" + ("a" + "a")) + "a") + "\n\t") {
    case "":
        switch "" {
        case "a":
            for fi59 := 0; fi59 < 4; fi59++ {
                switch genFn0((((v37 && v37) || (v37 || (v37 && false))) || v37), 12) {
                case -9223372036854775808:
                    fmt.Println(v1)
                case 0:
                    var v60 int16 = 0
                    _ = v60
                default:
                    v2 += v2
                }
                _ = fi59
            }
        default:
            var v61 uint = v3
            _ = v61
        }
    default:
        _ = len(v5)
    }
    fmt.Println(v4, v1)
    switch v4 {
    case 9223372036854775807:
        for fi62 := 0; fi62 < 14; fi62++ {
            for fi63 := 0; fi63 < 12; fi63++ {
                v2 = uint64(v3)
                v64 := GenS5{F0: (-128 <= -41)}
                _ = fi63
                _ = v64
            }
            for fi65 := 0; fi65 < 13; fi65++ {
                v66 := GenS4{F0: (fi65 & (fi65 - (fi62 >> (((v3 ^ v3)) & 63)))), F1: -1}
                _ = fi65
                _ = v66
            }
            _ = fi62
        }
    default:
        for fi67 := 0; fi67 < 19; fi67++ {
            for fi68 := 0; fi68 < 7; fi68++ {
                var v69 int8 = 1
                if v37 {
                    v70 := []int{genFn0((683 != 237), 42), genFn0(v37, 987), fi67, genFn0(v37, -1)}
                    fmt.Println(v3, v37)
                    if v37 {
                        var v71 int64 = 42
                        if false {
                            fmt.Println(v3, fi67)
                            var v72 int8 = (-(v69))
                            switch "foo bar" {
                            case "hello":
                                var v73 int16 = 1
                                _ = v73
                            default:
                                v74 := GenS7{F0: (!v37)}
                                _ = v74
                            }
                            _ = v72
                        } else {
                            v69 += (^(v69))
                        }
                        _ = v71
                    }
                    _ = v70
                } else {
                    v75 := []int{32768, (fi67 - -1), -596}
                    if v37 {
                        if v37 {
                            var v76 uint32 = 4294967295
                            fmt.Println(v4)
                            fmt.Println(fi68, v37, v69)
                            v3 -= (v3 / (((v3 % ((2) | 1))) | 1))
                            var v77 int16 = 32767
                            v78 := GenS9{F0: -2147483648, F1: 65535, F2: v1}
                            v77 -= (v77 / ((v77) | 1))
                            _ = v76
                            _ = v77
                            _ = v78
                        } else {
                            v4 = fi68
                            fmt.Println(fi68, fi67, v37)
                            var v79 int64 = v1
                            _ = v79
                        }
                        v80 := v5[0]
                        var v81 int16 = 295
                        var v82 uint = v3
                        if ((v81 | v81) <= 0) {
                            fmt.Println(v4)
                            v83 := GenS9{F0: 1, F1: 42, F2: v1}
                            v84 := GenS3{F0: "\n\t"}
                            _ = v83
                            _ = v84
                        } else {
                            if len(v75) > 0 {
                                v75[0] = (fi67 - (v4 / ((fi68) | 1)))
                            }
                            fmt.Println(v81, fi68)
                        }
                        _ = len(v75)
                        _ = v80
                        _ = v81
                        _ = v82
                    } else {
                        switch (("" + "\n\t") + ("a" + "foo bar")) {
                        case "":
                            if v37 {
                                v85 := v5[65535]
                                _ = v85
                            }
                        }
                        v1 += v1
                    }
                    _ = v75
                }
                v37 = (fi67 == (fi67 * 524))
                v86 := []int{(-(fi68)), (v4 * (v4 | (fi67 & (fi67 * v4)))), (-(v4))}
                _ = fi68
                _ = v69
                _ = v86
            }
            var v87 int8 = 42
            _ = fi67
            _ = v87
        }
    }
    v88 := []int{v4, (v4 / ((2147483647) | 1))}
    for fi89 := 0; fi89 < 16; fi89++ {
        v1 = 0
        fmt.Println(v37, v37, v2)
        _ = fi89
    }
    v90 := GenS8{F0: 18446744073709551615, F1: -128}
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v88:", len(v88), v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v88
    _ = v90
}
