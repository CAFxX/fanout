package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int32; F1 string; F2 int32 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 uint; F1 uint; F2 int8 }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int8; F1 int16; F2 uint }
func genFn0(p0 uint, p1 bool, p2 uint8) int8 {
    if ((p2 / ((193) | 1)) >= (p2 - 42)) {
        if false {
            v5 := GenS1{F0: (p2 - p2), F1: (true || true)}
            v6 := GenS1{F0: p2, F1: p1}
            for fi7 := 0; fi7 < 6; fi7++ {
                if p1 {
                    switch ("foo bar" + "hello") {
                    case "foo bar":
                        v8 := []int{(fi7 << ((p0) & 63))}
                        _ = v8
                    case "a":
                        for fi9 := 0; fi9 < 7; fi9++ {
                            if false {
                                if ((float64((0.0 * 0.0)) - -2.25) >= ((0.0 - float64(((1.5 - 1e10) * -2.25))) * 3.14159)) {
                                    v10 := GenS1{F0: 42, F1: false}
                                    p2 = 255
                                    p0 -= (p0 << ((p0) & 63))
                                    v11 := GenS1{F0: 127, F1: ((!p1) || (p0 >= (p0 + (p0 % ((42) | 1)))))}
                                    if (-9223372036854775808 > 9223372036854775807) {
                                        fmt.Println(p1, p2)
                                        var v12 int = (fi9 | 941)
                                        var v13 int = (-(fi7))
                                        _ = v12
                                        _ = v13
                                    } else {
                                        v14 := []int{fi9, (fi9 << ((p0) & 63)), 0}
                                        p1 = p1
                                        _ = v14
                                    }
                                    _ = v10
                                    _ = v11
                                }
                                v15 := GenS1{F0: (p2 << (((p0 % ((940) | 1))) & 63)), F1: (42 <= 1)}
                                _ = v15
                            } else {
                                v16 := GenS1{F0: 0, F1: p1}
                                v17 := GenS1{F0: p2, F1: (65535 <= 65535)}
                                p2 += (p2 + (p2 ^ (p2 >> ((p0) & 63))))
                                v18 := GenS1{F0: (^(p2)), F1: (false && true)}
                                p0 -= (^(p0))
                                p0 = (p0 / ((p0) | 1))
                                if (!((32767 != 1) && (!false))) {
                                    v19 := GenS1{F0: (p2 & (^(p2))), F1: ((1 > p2) && (0 <= 0))}
                                    var v20 int8 = -1
                                    _ = v19
                                    _ = v20
                                }
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            _ = fi9
                        }
                    }
                    p0 = (p0 << (((p0 & 943)) & 63))
                }
                if true {
                    v21 := []int{}
                    if p1 {
                        if p1 {
                            if (-1 == -128) {
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                            }
                            fmt.Println(fi7)
                            v22 := GenS1{F0: p2, F1: (!p1)}
                            _ = v22
                        }
                        p0 = (p0 | p0)
                        if ((-2.25 / 0.1) < 1.5) {
                            v23 := GenS1{F0: (p2 * (p2 & (p2 / ((p2) | 1)))), F1: (!p1)}
                            _ = v23
                        }
                    } else {
                        switch fi7 {
                        case 42:
                            fmt.Println(fi7, p1)
                        case -218:
                            if true {
                                switch (("\n\t" + ("" + "hello")) + "") {
                                case "hello":
                                    v24 := GenS2{F0: 42, F1: "hello", F2: -917}
                                    _ = v24
                                default:
                                    fmt.Println(p0, p1, p2)
                                }
                            } else {
                                p1 = p1
                                v25 := GenS2{F0: -990, F1: "a", F2: 2147483647}
                                p2 = p2
                                _ = v25
                            }
                        }
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                        fmt.Println(p0, p0)
                    }
                    _ = v21
                }
                _ = fi7
            }
            _ = v5
            _ = v6
        }
    }
    v26 := GenS3{F0: (false && p1), F1: 2}
    for fi27 := 0; fi27 < 14; fi27++ {
        v28 := []int{(^(fi27)), fi27, 42, fi27}
        _ = fi27
        _ = v28
    }
    var v29 string = "\n\t"
    var v30 uint16 = 163
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v29
    _ = v30
    return 1
}
func genFn1(p0 string, p1 uint) int {
    v31 := []int{128, 0, 32767, 127}
    var v32 uint = (p1 % ((p1) | 1))
    _ = p0
    _ = p1
    _ = v31
    _ = v32
    return -2147483648
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = -128
    var v3 int64 = -1
    var v4 uint32 = 65535
    v1[-397] = true
    v2 -= 100
    for fi33 := 0; fi33 < 1; fi33++ {
        v4 += 0
        _ = fi33
    }
    var v34 int16 = 0
    if (!((v34 > (v34 % ((-756) | 1))) || false)) {
        v2 -= (v2 ^ 103)
        fmt.Println(v3, v34)
        var v35 int8 = genFn0(1, true, 2)
        _ = v35
    } else {
        v2 = genFn0(320, (2 >= 283), 36)
        for fi36 := 0; fi36 < 2; fi36++ {
            for fi37 := 0; fi37 < 8; fi37++ {
                fmt.Println(fi37)
                _ = fi37
            }
            if true {
                v1[fi36] = false
                v2 -= v2
                _ = len(v1)
            }
            _ = fi36
        }
    }
    v38 := GenS4{F0: true, F1: (-(v2))}
    if (!(-1 < 77)) {
        switch -510 {
        case -151:
            fmt.Println(v4, v4, v3)
        default:
            for fi39 := 0; fi39 < 11; fi39++ {
                v40 := GenS3{F0: true, F1: 0}
                v41 := GenS5{F0: 255, F1: 1, F2: 25}
                v34 = v34
                _ = fi39
                _ = v40
                _ = v41
            }
        }
        for fi42 := 0; fi42 < 8; fi42++ {
            v34 = (v34 % ((v34) | 1))
            if ((v34 + v34) > (v34 >> ((18446744073709551615) & 63))) {
                if (!true) {
                    var v43 int16 = 607
                    _ = v43
                } else {
                    if false {
                        v2 += 0
                        var v44 string = "hello"
                        _ = v44
                    }
                }
                delete(v1, fi42)
                for fi45 := 0; fi45 < 16; fi45++ {
                    switch "a" {
                    case "a":
                        v46 := GenS6{F0: v3, F1: (v4 * v4)}
                        _ = v46
                    case "\n\t":
                        v47 := GenS1{F0: 128, F1: (!(!true))}
                        _ = v47
                    case "hello":
                        if (true || (0 != 1)) {
                            v34 -= (v34 | -32768)
                            var v48 int = fi45
                            _ = v48
                        }
                    }
                    v49 := GenS7{F0: 128}
                    if ((^(v4)) <= v4) {
                        fmt.Println(fi42, fi45, fi45)
                        v3 += (^(v3))
                        if (0.1 <= (1.5 + 1.5)) {
                            v3 -= 0
                            var v50 uint = 0
                            _ = v50
                        } else {
                            var v51 int32 = -2147483648
                            if false {
                                v52 := []int{-180, genFn1("foo bar", 2), fi42, (fi45 + (fi42 - fi45))}
                                _ = v52
                            }
                            _ = v51
                        }
                        var v53 uint8 = 127
                        v53 += (v53 - 1)
                        var v54 bool = ((v53 << ((0) & 63)) > v53)
                        v2 += genFn0(2, v54, 127)
                        v2 = (v2 * genFn0(255, (714 == 0), uint8(v4)))
                        _ = v53
                        _ = v54
                    } else {
                        var v55 bool = false
                        switch "\n\t" {
                        case "":
                            fmt.Println(v55)
                        default:
                            var v56 uint64 = 18446744073709551615
                            _ = v56
                        }
                        _ = v55
                    }
                    _ = fi45
                    _ = v49
                }
            } else {
                if false {
                    if ((v3 >> ((uint(v4)) & 63)) != v3) {
                        switch (^(fi42)) {
                        case 4294967295:
                            for fi57 := 0; fi57 < 16; fi57++ {
                                switch "hello" {
                                case "hello":
                                    v58 := []int{(fi42 + fi57), (^(fi57)), 4294967295, genFn1((("hello" + "foo bar") + ("" + ("\n\t" + "foo bar"))), 2)}
                                    _ = v58
                                default:
                                    var v59 int16 = v34
                                    _ = v59
                                }
                                _ = fi57
                            }
                        case -2147483648:
                            v4 += 2
                        default:
                            v3 += int64(v4)
                        }
                        if true {
                            var v60 uint32 = (v4 & v4)
                            _ = v60
                        }
                        for fi61 := 0; fi61 < 6; fi61++ {
                            fmt.Println(fi42, v4)
                            var v62 uint64 = 2
                            v4 = v4
                            _ = fi61
                            _ = v62
                        }
                    }
                }
                if (genFn1("", 2) <= (^(fi42))) {
                    fmt.Println(v34, v4, v4)
                }
                var v63 uint16 = 926
                _ = v63
            }
            v64 := GenS5{F0: 255, F1: 65535, F2: (^(v2))}
            _ = fi42
            _ = v64
        }
        v65 := GenS4{F0: (805 < 0), F1: v2}
        _ = v65
    } else {
        for fi66 := 0; fi66 < 17; fi66++ {
            for fi67 := 0; fi67 < 16; fi67++ {
                v68 := GenS8{F0: (v34 + v34)}
                var v69 uint = 399
                var v70 uint16 = 17
                fmt.Println(fi66, v2)
                v70 = 2
                v34 = int16(v2)
                _ = fi67
                _ = v68
                _ = v69
                _ = v70
            }
            v71 := []int{fi66, fi66, (fi66 / ((fi66) | 1))}
            _ = fi66
            _ = v71
        }
        v72 := GenS2{F0: 2147483647, F1: ("hello" + "foo bar"), F2: 2147483647}
        for fi73 := 0; fi73 < 19; fi73++ {
            v74 := GenS9{F0: 125, F1: (^(v34)), F2: 255}
            _ = fi73
            _ = v74
        }
        _ = v72
    }
    var v75 int = genFn1("a", uint(v4))
    for fi76 := 0; fi76 < 4; fi76++ {
        v3 += 388
        v4 = 940
        v77 := []int{}
        _ = fi76
        _ = v77
    }
    var v78 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v75:", v75)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v75
    _ = v78
}
