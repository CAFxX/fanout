package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int32; F1 int; F2 int32 }
type GenS3 struct { F0 bool; F1 int16; F2 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 uint16; F2 int64 }
type GenS6 struct { F0 int64; F1 int; F2 string }
type GenS7 struct { F0 int }
func genFn0(p0 bool) uint8 {
    v4 := GenS1{F0: p0, F1: 1}
    p0 = true
    _ = p0
    _ = v4
    return 42
}
func genFn1(p0 bool, p1 uint16, p2 int16) bool {
    p2 -= (^(p2))
    v5 := []int{32767, 2147483647, int(p2)}
    fmt.Println(p2, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p0
}

func main() {
    v1 := []uint32{42, 1}
    var v2 int64 = -9223372036854775808
    var v3 uint16 = 610
    var v6 int32 = 42
    var v7 uint = 444
    v2 += v2
    if true {
        if (0 > 4294967295) {
            v7 -= (v7 + (v7 | (v7 % (((v7 / ((v7) | 1))) | 1))))
            switch 127 {
            case 32767:
                switch -553 {
                case -842:
                    var v8 float32 = 1e10
                    _ = v8
                case -1:
                    if false {
                        for fi9 := 0; fi9 < 12; fi9++ {
                            _ = len(v1)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi9
                        }
                        fmt.Println(v2, v2, v2)
                        var v10 int64 = -9223372036854775808
                        _ = v10
                    }
                case -369:
                    for fi11 := 0; fi11 < 1; fi11++ {
                        if (v2 <= (^(v2))) {
                            fmt.Println(v6, v2)
                        }
                        for fi12 := 0; fi12 < 1; fi12++ {
                            if ((uint32(v7) > 2147483647) && ((^(v3)) >= v3)) {
                                if (!(-113 != -46)) {
                                    var v13 float32 = 1e10
                                    _ = v13
                                }
                            } else {
                                v6 -= (v6 / ((v6) | 1))
                            }
                            v14 := GenS2{F0: -2147483648, F1: fi12, F2: v6}
                            _ = fi12
                            _ = v14
                        }
                        _ = fi11
                    }
                default:
                    var v15 uint32 = 0
                    _ = v15
                }
            case -461:
                var v16 string = (("\n\t" + ("hello" + "foo bar")) + "foo bar")
                _ = v16
            case -1:
                v17 := GenS3{F0: (332 > (v3 / ((v3) | 1))), F1: 1, F2: 1}
                _ = v17
            default:
                for fi18 := 0; fi18 < 19; fi18++ {
                    if true {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = fi18
                }
            }
            var v19 uint = (v7 & (v7 + 515))
            _ = v19
        }
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 string = "a"
            _ = fi20
            _ = v21
        }
    } else {
        v2 = -9223372036854775808
        for fi22 := 0; fi22 < 5; fi22++ {
            v7 -= 217
            if ((57 < 127) || (112 > 78)) {
                v6 = (v6 - (v6 << ((uint(v3)) & 63)))
            } else {
                v23 := GenS4{F0: 0}
                v7 = v7
                v6 = v6
                _ = v23
            }
            _ = fi22
        }
        v2 = (v2 / ((9223372036854775807) | 1))
    }
    _ = len(v1)
    switch ("a" + ("a" + "\n\t")) {
    case "foo bar":
        if true {
            if false {
                if ((v6 ^ (-(v6))) == v6) {
                    v2 -= v2
                } else {
                    if ((!(1 > 255)) || false) {
                        if (genFn0(false) == 1) {
                            v1 = append(v1, 63)
                            var v24 int = 32767
                            switch ("hello" + ("hello" + "hello")) {
                            case "":
                                _ = len(v1)
                            case "\n\t":
                                v24 = (^(v24))
                            case "a":
                                for fi25 := 0; fi25 < 11; fi25++ {
                                    v26 := GenS5{F0: 38, F1: v3, F2: (v2 / ((v2) | 1))}
                                    _ = fi25
                                    _ = v26
                                }
                            }
                            _ = v24
                        }
                        var v27 int32 = -1
                        v28 := GenS4{F0: int8(v3)}
                        _ = v27
                        _ = v28
                    } else {
                        var v29 int16 = 42
                        var v30 int = -451
                        _ = v29
                        _ = v30
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v2 = 9223372036854775807
                }
                if len(v1) > 0 {
                    v1[0] = 2
                }
            }
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            var v31 int = -2147483648
            _ = v31
        } else {
            v32 := GenS1{F0: (32767 > 2147483647), F1: 1}
            _ = v32
        }
    default:
        v6 = (v6 & (v6 - (v6 >> ((v7) & 63))))
    }
    if (true && false) {
        switch 0 {
        case -839:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 2147483647:
            switch 1 {
            case 2147483647:
                fmt.Println(v7, v2)
            default:
                var v33 uint16 = (v3 ^ v3)
                _ = v33
            }
        }
        v34 := GenS6{F0: (v2 ^ 0), F1: -1, F2: ("a" + "")}
        _ = v34
    }
    v35 := GenS3{F0: false, F1: -32768, F2: 2}
    if (!((85 >= -597) || false)) {
        var v36 int = 9223372036854775807
        v37 := GenS6{F0: (v2 << ((4294967295) & 63)), F1: (v36 >> (((^(v7))) & 63)), F2: "\n\t"}
        switch v36 {
        case -1:
            var v38 float64 = 0.0
            _ = v38
        case 127:
            for fi39 := 0; fi39 < 4; fi39++ {
                fmt.Println(v36)
                _ = fi39
            }
        case 42:
            var v40 uint16 = (^(v3))
            _ = v40
        default:
            var v41 bool = ((1 <= 1) && false)
            _ = v41
        }
        _ = v36
        _ = v37
    }
    if ((0.0 + -2.25) == (1e10 - 1.5)) {
        v1 = append(v1, 65535)
        for fi42 := 0; fi42 < 16; fi42++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v1 = append(v1, 2147483648)
            var v43 int64 = v2
            _ = fi42
            _ = v43
        }
        var v44 float64 = 3.14159
        _ = v44
    }
    v1 = append(v1, 2)
    for fi45 := 0; fi45 < 7; fi45++ {
        v46 := GenS4{F0: 42}
        _ = fi45
        _ = v46
    }
    v47 := GenS6{F0: v2, F1: -987, F2: ("\n\t" + "foo bar")}
    if (!true) {
        var v48 int64 = 0
        v7 += (v7 & v7)
        var v49 int16 = -32768
        _ = v48
        _ = v49
    } else {
        var v50 uint = (v7 - 4294967295)
        _ = v50
    }
    for fi51 := 0; fi51 < 19; fi51++ {
        v2 -= 630
        var v52 float32 = float32(1e10)
        for fi53 := 0; fi53 < 9; fi53++ {
            var v54 uint = v7
            var v55 bool = (!false)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi53
            _ = v54
            _ = v55
        }
        _ = fi51
        _ = v52
    }
    var v56 uint32 = 2147483648
    switch -687 {
    case 820:
        fmt.Println(v3, v7, v7)
    case 42:
        if true {
            v57 := GenS7{F0: -9223372036854775808}
            var v58 uint16 = v3
            _ = v57
            _ = v58
        }
    case 134:
        v3 += v3
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v47
    _ = v56
    _ = v6
    _ = v7
}
