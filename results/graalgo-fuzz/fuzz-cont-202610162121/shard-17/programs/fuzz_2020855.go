package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int8; F1 int64; F2 int }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 int64 }
func genFn0(p0 int, p1 uint, p2 int32) uint16 {
    switch (^(p0)) {
    case 590:
        p0 -= p0
    case -1:
        var v4 int8 = -28
        _ = v4
    case 127:
        p1 += (p1 / (((p1 << (((p1 / ((749) | 1))) & 63))) | 1))
    default:
        for fi5 := 0; fi5 < 5; fi5++ {
            p1 += (^(p1))
            p2 -= p2
            _ = fi5
        }
    }
    switch p0 {
    case -236:
        v6 := GenS1{F0: 61}
        _ = v6
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return 32767
}
func genFn1(p0 uint8, p1 uint16) uint {
    var v7 float64 = ((1.5 * float64((-2.25 / -2.25))) / 1.0)
    p0 += 2
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    var v1 uint16 = 2
    var v2 int32 = 0
    v3 := []string{("foo bar" + (("foo bar" + "hello") + ("foo bar" + "hello"))), "hello", "hello", ("\n\t" + "hello")}
    v8 := GenS1{F0: 2}
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(v1, v1)
        v10 := GenS2{F0: 1}
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 uint64 = 18446744073709551615
            if false {
                v1 -= genFn0((^(fi11)), genFn1(128, v1), v2)
                var v13 bool = true
                fmt.Println(v12)
                _ = v13
            } else {
                v1 = genFn0(4294967295, uint(v1), v2)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if false {
                    v1 = genFn0(fi9, 65535, (v2 | (v2 * 1)))
                    v14 := GenS3{F0: (v12 % ((v12) | 1)), F1: v12, F2: 445}
                    switch (fi11 * fi9) {
                    case 127:
                        v3 = append(v3, "a")
                    case -1:
                        v2 += v2
                    default:
                        if ((42 == genFn1(175, v1)) || (true || (42 != -32768))) {
                            v12 += v12
                            v15 := GenS1{F0: 165}
                            v1 += (v1 + (v1 + uint16(v12)))
                            _ = v15
                        }
                    }
                    _ = v14
                }
            }
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    v2 += (-(v2))
    if (true && true) {
        fmt.Println(v1, v1)
        var v16 uint8 = 119
        var v17 uint64 = 21
        _ = v16
        _ = v17
    }
    v2 = (^(v2))
    v18 := GenS4{F0: 127, F1: -1, F2: 196}
    var v19 uint = 675
    if (295 <= 42) {
        v20 := GenS1{F0: 255}
        _ = v20
    } else {
        if false {
            v2 += 1
            if (int32(v19) < (v2 | 0)) {
                if true {
                    if (true && true) {
                        v1 += (v1 % ((v1) | 1))
                        if (-1 <= 0) {
                            v21 := GenS2{F0: -1}
                            var v22 int = 9223372036854775807
                            var v23 int8 = 100
                            _ = v21
                            _ = v22
                            _ = v23
                        } else {
                            for fi24 := 0; fi24 < 4; fi24++ {
                                var v25 uint32 = 1
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi24
                                _ = v25
                            }
                            if false {
                                var v26 string = "a"
                                _ = v26
                            } else {
                                v19 = 1
                                var v27 string = ("foo bar" + ((((((("" + (("" + ("a" + "a")) + "foo bar")) + "foo bar") + (("foo bar" + ("\n\t" + "")) + "")) + "") + "hello") + (("" + ("hello" + "\n\t")) + "hello")) + ""))
                                var v28 uint16 = genFn0(58, v19, 42)
                                _ = v27
                                _ = v28
                            }
                        }
                    }
                    if (("hello" + ("\n\t" + "a")) < (("" + "foo bar") + "a")) {
                        v19 -= v19
                        for fi29 := 0; fi29 < 10; fi29++ {
                            var v30 int16 = 1
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi29
                            _ = v30
                        }
                        switch "a" {
                        case "":
                            if false {
                                var v31 uint32 = 884
                                for fi32 := 0; fi32 < 12; fi32++ {
                                    v1 = v1
                                    _ = fi32
                                }
                                _ = v31
                            }
                        default:
                            for fi33 := 0; fi33 < 7; fi33++ {
                                fmt.Println(v1, v19, fi33)
                                var v34 uint64 = 42
                                v35 := GenS3{F0: 1, F1: v34, F2: 0}
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                        }
                    } else {
                        var v36 int = 0
                        _ = v36
                    }
                    v37 := GenS5{F0: int8(v19), F1: 42}
                    _ = v37
                } else {
                    var v38 bool = false
                    if v38 {
                        for fi39 := 0; fi39 < 7; fi39++ {
                            v40 := GenS2{F0: int16(fi39)}
                            v1 -= genFn0(2147483647, 459, v2)
                            v38 = v38
                            _ = fi39
                            _ = v40
                        }
                        fmt.Println(v38, v1, v19)
                        if true {
                            fmt.Println(v38, v2, v1)
                            var v41 int = 128
                            var v42 float64 = 1e10
                            _ = v41
                            _ = v42
                        }
                    } else {
                        v38 = true
                        switch ("foo bar" + "\n\t") {
                        case "hello":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "a":
                            if false {
                                v2 -= -2147483648
                            }
                        case "":
                            if v38 {
                                fmt.Println(v1)
                                v2 = (v2 * (v2 | (v2 >> (((v19 | v19)) & 63))))
                            } else {
                                switch ("foo bar" + "\n\t") {
                                case "foo bar":
                                    v19 -= v19
                                }
                                v43 := GenS6{F0: 42, F1: "\n\t"}
                                var v44 uint64 = 18446744073709551615
                                _ = v43
                                _ = v44
                            }
                        default:
                            for fi45 := 0; fi45 < 13; fi45++ {
                                if (9223372036854775807 != -724) {
                                    switch ("a" + "hello") {
                                    case "foo bar":
                                        if v38 {
                                            switch "\n\t" {
                                            case "a":
                                                _ = len(v3)
                                            }
                                        } else {
                                            var v46 int = (-(fi45))
                                            var v47 int = (v46 >> ((v19) & 63))
                                            var v48 float32 = float32(1.5)
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        }
                                    case "":
                                        v1 += (v1 ^ v1)
                                    }
                                } else {
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    if v38 {
                                        for fi49 := 0; fi49 < 4; fi49++ {
                                            v2 -= (^(v2))
                                            v1 = (v1 >> ((2) & 63))
                                            v1 += genFn0(1, v19, (-(v2)))
                                            _ = fi49
                                        }
                                    } else {
                                        v38 = v38
                                    }
                                }
                                _ = fi45
                            }
                        }
                    }
                    if len(v3) > 0 {
                        v3[0] = (("a" + "") + "foo bar")
                    }
                    _ = v38
                }
                for fi50 := 0; fi50 < 9; fi50++ {
                    for fi51 := 0; fi51 < 9; fi51++ {
                        switch ("\n\t" + "\n\t") {
                        case "":
                            v3 = append(v3, "foo bar")
                        case "\n\t":
                            var v52 int8 = 1
                            _ = v52
                        case "hello":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            _ = len(v3)
                        }
                        _ = fi51
                    }
                    v19 -= v19
                    _ = fi50
                }
                if ((!(0 != 2147483648)) && (-1 < 127)) {
                    for fi53 := 0; fi53 < 8; fi53++ {
                        v3 = append(v3, "a")
                        var v54 int32 = (v2 & v2)
                        var v55 uint16 = 1
                        _ = fi53
                        _ = v54
                        _ = v55
                    }
                    var v56 int = 0
                    for fi57 := 0; fi57 < 4; fi57++ {
                        for fi58 := 0; fi58 < 15; fi58++ {
                            v56 = (fi58 >> (((^(v19))) & 63))
                            var v59 string = "hello"
                            switch (fi57 | -9223372036854775808) {
                            case -1:
                                if (0 != 9223372036854775807) {
                                    var v60 int64 = 76
                                    var v61 uint8 = 128
                                    _ = v60
                                    _ = v61
                                } else {
                                    fmt.Println(v59, v19)
                                    var v62 int8 = -128
                                    _ = v62
                                }
                            case -9223372036854775808:
                                var v63 float64 = -2.25
                                _ = v63
                            }
                            _ = fi58
                            _ = v59
                        }
                        if true {
                            _ = len(v3)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v64 int8 = 100
                            _ = v64
                        }
                        if (-128 < -1) {
                            _ = len(v3)
                        } else {
                            if false {
                                v65 := GenS3{F0: 2, F1: 4294967295, F2: 668}
                                _ = v65
                            }
                            if ((^(v19)) <= v19) {
                                if ((!(-2.25 != 0.0)) && false) {
                                    var v66 float64 = 1e10
                                    v1 += (v1 ^ 42)
                                    _ = v66
                                }
                            } else {
                                for fi67 := 0; fi67 < 20; fi67++ {
                                    var v68 int32 = (-(v2))
                                    _ = fi67
                                    _ = v68
                                }
                                v56 += (v56 >> (((v19 - v19)) & 63))
                            }
                            if len(v3) > 0 {
                                v3[0] = "a"
                            }
                        }
                        _ = fi57
                    }
                    _ = v56
                }
            } else {
                fmt.Println(v1)
                var v69 float64 = 1.5
                var v70 uint8 = 127
                _ = v69
                _ = v70
            }
            if ((v2 / ((v2) | 1)) >= -2147483648) {
                var v71 uint = (v19 - (v19 + v19))
                for fi72 := 0; fi72 < 17; fi72++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v71 = (^(v71))
                    _ = fi72
                }
                v73 := GenS7{F0: 1, F1: 1}
                _ = v71
                _ = v73
            } else {
                fmt.Println(v19)
                v19 -= v19
                v1 = uint16(v19)
            }
        } else {
            if true {
                v74 := GenS7{F0: 4294967295, F1: 4294967295}
                var v75 uint32 = 2147483647
                _ = v74
                _ = v75
            } else {
                if ((3.14159 - -2.25) == 0.1) {
                    if true {
                        v76 := GenS8{F0: -1}
                        _ = v76
                    } else {
                        if (487 != 1) {
                            var v77 int8 = -128
                            _ = v77
                        }
                    }
                }
                if (true || true) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v78 uint32 = 43
                    _ = v78
                } else {
                    var v79 bool = ((v1 | 32767) != 32767)
                    v3 = append(v3, "hello")
                    _ = v79
                }
                v80 := GenS1{F0: 2}
                _ = v80
            }
            for fi81 := 0; fi81 < 17; fi81++ {
                v2 += -2147483648
                var v82 int16 = 1000
                switch fi81 {
                case -2147483648:
                    if (!((42 <= 589) && true)) {
                        v83 := GenS5{F0: 0, F1: v1}
                        _ = v83
                    } else {
                        fmt.Println(v82, v82, fi81)
                        if ((^(v1)) != 0) {
                            v82 = v82
                            var v84 uint32 = 1
                            _ = v84
                        }
                        v82 -= (v82 / ((v82) | 1))
                    }
                case -522:
                    v85 := GenS7{F0: 685, F1: 18446744073709551615}
                    _ = v85
                default:
                    switch -2147483648 {
                    case 9223372036854775807:
                        for fi86 := 0; fi86 < 14; fi86++ {
                            _ = len(v3)
                            switch "hello" {
                            case "foo bar":
                                if len(v3) > 0 {
                                    v3[0] = "foo bar"
                                }
                            }
                            _ = fi86
                        }
                    case 128:
                        var v87 int32 = v2
                        _ = v87
                    default:
                        var v88 uint8 = 1
                        _ = v88
                    }
                }
                _ = fi81
                _ = v82
            }
            var v89 uint64 = 42
            _ = v89
        }
    }
    v1 += genFn0(1, 64, v2)
    v1 = 663
    v2 += (v2 / (((v2 | 2147483647)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
}
