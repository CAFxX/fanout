package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool; F1 int64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int8; F1 uint }
type GenS10 struct { F0 uint16 }
func genFn0(p0 uint64) string {
    v7 := GenS1{F0: 4294967295}
    switch "a" {
    case "":
        p0 = (^(p0))
    case "hello":
        for fi8 := 0; fi8 < 12; fi8++ {
            switch "" {
            case "\n\t":
                var v9 string = "foo bar"
                _ = v9
            case "a":
                var v10 int64 = 463
                _ = v10
            case "hello":
                p0 -= 2
            }
            _ = fi8
        }
    case "a":
        p0 -= p0
    default:
        for fi11 := 0; fi11 < 19; fi11++ {
            v12 := GenS1{F0: (p0 - p0)}
            if (0 > (p0 * (p0 - p0))) {
                v13 := []int{}
                if (false || false) {
                    fmt.Println(fi11, p0)
                    v13 = append(v13, fi11)
                }
                _ = v13
            }
            var v14 uint64 = (^(p0))
            _ = fi11
            _ = v12
            _ = v14
        }
    }
    _ = p0
    _ = v7
    return ((((("\n\t" + "hello") + ("" + "foo bar")) + "a") + "") + ("a" + ("\n\t" + (("a" + "hello") + ("foo bar" + "")))))
}
func genFn1(p0 bool, p1 uint16) uint8 {
    p1 -= p1
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]bool)
    var v3 float32 = 1.5
    var v4 uint8 = 2
    v5 := make(map[uint32]int)
    v6 := []uint32{}
    v15 := GenS2{F0: 148}
    v16 := GenS1{F0: 360}
    switch "foo bar" {
    case "foo bar":
        v17 := GenS1{F0: 4294967295}
        _ = v17
    case "hello":
        var v18 float64 = 1e10
        _ = v18
    case "":
        if ("hello" <= "foo bar") {
            v4 = (v4 - (v4 << ((2) & 63)))
            v19 := GenS2{F0: -9223372036854775808}
            var v20 int32 = -1
            _ = v19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        delete(v2, fi21)
        v4 += 85
        _ = fi21
    }
    v3 += ((float32(((((((v1 * 1e10) * v1) / v1) + v1) / 1.0) / float64(-2.25))) - (v3 - (v3 * (1e10 + 1e10)))) + v3)
    if ((!false) || false) {
        if (("\n\t" <= (("" + "") + (("a" + "\n\t") + "hello"))) || true) {
            var v22 int8 = -1
            v23 := GenS3{F0: 0, F1: 42, F2: 788}
            _ = v22
            _ = v23
        }
        switch ("a" + ("foo bar" + "a")) {
        case "\n\t":
            var v24 uint8 = v4
            _ = v24
        case "a":
            for fi25 := 0; fi25 < 1; fi25++ {
                var v26 uint = 65535
                _ = fi25
                _ = v26
            }
        }
        var v27 uint = 2
        _ = v27
    } else {
        fmt.Println(v4, v1)
        _ = len(v6)
        var v28 uint64 = 304
        _ = v28
    }
    if false {
        if ("foo bar" >= "") {
            for fi29 := 0; fi29 < 4; fi29++ {
                var v30 float64 = 0.1
                if (true && (!(!(585 == 0)))) {
                    fmt.Println(fi29, fi29)
                    fmt.Println(v30, fi29, v1)
                } else {
                    if false {
                        fmt.Println(v1)
                        if true {
                            v31 := GenS1{F0: 676}
                            _ = v31
                        } else {
                            v4 += 71
                        }
                    } else {
                        v32 := GenS4{F0: fi29, F1: -99}
                        if (4294967295 != 427) {
                            var v33 float64 = v30
                            _ = v33
                        } else {
                            var v34 uint32 = 655
                            v35 := GenS2{F0: fi29}
                            _ = v34
                            _ = v35
                        }
                        _ = v32
                    }
                    if (710 >= -2147483648) {
                        v36 := v2[(fi29 % ((-9223372036854775808) | 1))]
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 17; fi37++ {
                        if false {
                            var v38 int32 = 986
                            if false {
                                if (1.5 <= (v3 * (v3 * v3))) {
                                    v39 := v5[1]
                                    v40 := GenS5{F0: (^(v4)), F1: 4294967295}
                                    v41 := GenS6{F0: -32768}
                                    switch ((("foo bar" + "\n\t") + "foo bar") + ("" + "a")) {
                                    case "\n\t":
                                        var v42 uint64 = 242
                                        _ = v42
                                    case "a":
                                        v43 := v2[v39]
                                        _ = v43
                                    case "foo bar":
                                        _ = len(v6)
                                    }
                                    var v44 uint = 18446744073709551615
                                    v30 += (1e10 * v1)
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                    _ = v44
                                }
                                if true {
                                    v45 := GenS7{F0: true, F1: -9223372036854775808}
                                    _ = v45
                                }
                                fmt.Println(fi37)
                            } else {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if (((v4 << ((173) & 63)) != (v4 ^ v4)) || (((false || (65535 > 255)) || false) && ((18446744073709551615 == 772) || (!true)))) {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v46 := GenS8{F0: 1}
                                        _ = v46
                                    case "foo bar":
                                        switch 2147483647 {
                                        case 127:
                                            fmt.Println(v3, fi29, v4)
                                        case 1:
                                            v47 := GenS6{F0: 0}
                                            _ = v47
                                        case 0:
                                            v48 := GenS4{F0: fi37, F1: -107}
                                            _ = v48
                                        default:
                                            v49 := GenS9{F0: -1, F1: 18446744073709551615}
                                            _ = v49
                                        }
                                    case "":
                                        if (false || (358 < 676)) {
                                            var v50 int64 = 9223372036854775807
                                            if ((100 <= 0) || false) {
                                                v1 = v1
                                            } else {
                                                fmt.Println(v1)
                                                if ((2 != 449) && (634 >= -32768)) {
                                                    var v51 int16 = 206
                                                    _ = v51
                                                }
                                            }
                                            _ = v50
                                        } else {
                                            if (!true) {
                                                var v52 uint16 = 32767
                                                v5[42] = 52
                                                _ = v52
                                            }
                                            v4 += (v4 - genFn1(("" < "hello"), 0))
                                            if true {
                                                var v53 int8 = -125
                                                switch "a" {
                                                case "a":
                                                    if false {
                                                        var v54 int8 = v53
                                                        var v55 uint8 = (v4 & v4)
                                                        var v56 uint = 584
                                                        _ = v54
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                }
                                                v57 := GenS6{F0: 306}
                                                _ = v53
                                                _ = v57
                                            }
                                        }
                                    }
                                } else {
                                    if (true && false) {
                                        switch "a" {
                                        case "a":
                                            v6 = append(v6, 1)
                                        case "\n\t":
                                            var v58 int16 = -444
                                            _ = v58
                                        default:
                                            v59 := GenS8{F0: 18446744073709551615}
                                            _ = v59
                                        }
                                    } else {
                                        var v60 int = (fi37 ^ fi29)
                                        var v61 uint = 42
                                        v2[-2147483648] = true
                                        _ = v60
                                        _ = v61
                                    }
                                    if (830 < 2) {
                                        v6 = append(v6, 0)
                                    }
                                    var v62 uint32 = 65535
                                    v2[fi29] = false
                                    v5[(v62 & 4294967295)] = (fi29 & fi37)
                                    _ = v62
                                }
                                v30 = (v1 * v1)
                            }
                            _ = v38
                        } else {
                            var v63 bool = true
                            _ = v63
                        }
                        _ = fi37
                    }
                }
                _ = fi29
                _ = v30
            }
            var v64 int8 = -1
            _ = v64
        }
        switch 128 {
        case 127:
            if ((false && false) || (!true)) {
                if false {
                    var v65 float32 = (v3 * v3)
                    var v66 float64 = ((0.0 / v1) / (v1 / 3.14159))
                    v66 = v66
                    _ = v65
                    _ = v66
                } else {
                    var v67 int8 = 1
                    switch "foo bar" {
                    case "hello":
                        v67 = (v67 | v67)
                    default:
                        fmt.Println(v4)
                    }
                    _ = v67
                }
                v68 := GenS10{F0: 42}
                _ = v68
            }
        default:
            v69 := GenS8{F0: 4294967295}
            _ = v69
        }
    }
    v70 := GenS7{F0: false, F1: 42}
    v4 = v4
    var v71 uint = uint(v4)
    for fi72 := 0; fi72 < 8; fi72++ {
        if (!((1e10 > v1) && (127 > -128))) {
            v73 := GenS6{F0: -849}
            var v74 int64 = 42
            fmt.Println(v4, fi72, v74)
            _ = v73
            _ = v74
        } else {
            var v75 uint16 = 32767
            var v76 int8 = 66
            v77 := GenS5{F0: 84, F1: (^(fi72))}
            _ = v75
            _ = v76
            _ = v77
        }
        v78 := GenS8{F0: 18446744073709551615}
        _ = fi72
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v71
}
