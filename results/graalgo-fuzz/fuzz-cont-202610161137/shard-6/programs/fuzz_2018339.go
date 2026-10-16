package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 string; F2 int16 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 uint32; F1 int; F2 int64 }
func genFn0(p0 uint8, p1 int64, p2 uint) uint64 {
    var v6 int8 = 42
    v7 := []int{127, 4294967295}
    if len(v7) > 0 {
        v7[0] = 9223372036854775807
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            fmt.Println(fi9)
            switch "hello" {
            case "\n\t":
                var v10 float64 = 0.0
                if len(v7) > 0 {
                    v7[0] = (fi9 * (fi8 << (((^(p2))) & 63)))
                }
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v10
            case "a":
                p1 -= (p1 * p1)
                p0 -= 18
                p1 -= p1
            case "hello":
                v6 = (^(v6))
            default:
                p2 += 4294967295
            }
            _ = fi9
        }
        v6 = 0
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 string, p1 int16) uint8 {
    if (!(42 == 670)) {
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 string = "\n\t"
            _ = fi11
            _ = v12
        }
        p1 = p1
    }
    if (502 < 2) {
        for fi13 := 0; fi13 < 11; fi13++ {
            if (!false) {
                switch (-(fi13)) {
                case -207:
                    p0 = p0
                case 32768:
                    var v14 int64 = 9223372036854775807
                    _ = v14
                case -1:
                    for fi15 := 0; fi15 < 17; fi15++ {
                        switch (fi15 - 9223372036854775807) {
                        case 172:
                            fmt.Println(fi15, fi15, p1)
                        default:
                            p1 -= (p1 % ((p1) | 1))
                        }
                        switch "" {
                        case "":
                            p1 -= p1
                        default:
                            v16 := GenS1{F0: int8(fi15)}
                            _ = v16
                        }
                        v17 := GenS1{F0: 127}
                        _ = fi15
                        _ = v17
                    }
                }
                switch fi13 {
                case 2147483647:
                    for fi18 := 0; fi18 < 17; fi18++ {
                        v19 := []int{}
                        _ = fi18
                        _ = v19
                    }
                case 4294967295:
                    for fi20 := 0; fi20 < 16; fi20++ {
                        p0 = ("a" + (p0 + p0))
                        var v21 string = (p0 + p0)
                        v21 = (((p0 + "foo bar") + "hello") + ((v21 + "a") + "hello"))
                        _ = fi20
                        _ = v21
                    }
                case -371:
                    if (57 <= -883) {
                        var v22 int16 = 1000
                        v23 := GenS1{F0: 1}
                        _ = v22
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 4; fi24++ {
                            var v25 uint16 = 0
                            _ = fi24
                            _ = v25
                        }
                        fmt.Println(p0, p0, p0)
                        for fi26 := 0; fi26 < 6; fi26++ {
                            v27 := []int{(fi26 % (((-(fi26))) | 1))}
                            var v28 uint32 = 2147483648
                            v29 := GenS1{F0: 1}
                            _ = fi26
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    }
                default:
                    fmt.Println(p0)
                }
            } else {
                p0 = ""
                p1 += p1
                fmt.Println(p1, p0)
            }
            var v30 uint16 = 1
            _ = fi13
            _ = v30
        }
    } else {
        v31 := []int{32768, 1}
        for fi32 := 0; fi32 < 7; fi32++ {
            p1 = (p1 ^ (^(p1)))
            _ = fi32
        }
        _ = v31
    }
    p0 = "\n\t"
    _ = p0
    _ = p1
    return 180
}

func main() {
    var v1 int16 = -1
    v2 := []int{-1, 42, -9223372036854775808}
    var v3 int8 = 42
    var v4 float32 = 1.5
    var v5 uint64 = 927
    _ = len(v2)
    var v33 uint8 = genFn1(("" + "\n\t"), -1)
    for fi34 := 0; fi34 < 11; fi34++ {
        _ = len(v2)
        _ = fi34
    }
    var v35 int32 = 2147483647
    if true {
        if ((v1 & (v1 + v1)) <= -27) {
            for fi36 := 0; fi36 < 5; fi36++ {
                for fi37 := 0; fi37 < 17; fi37++ {
                    v33 -= v33
                    _ = fi37
                }
                if ((!true) || false) {
                    var v38 uint16 = uint16(v5)
                    if true {
                        v4 = (float32(float64(0.0)) * v4)
                    }
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 4; fi39++ {
                        switch "hello" {
                        case "foo bar":
                            var v40 float64 = 0.1
                            if ((^(v1)) < v1) {
                                var v41 string = ("\n\t" + ("\n\t" + "a"))
                                _ = v41
                            }
                            fmt.Println(v4)
                            _ = v40
                        case "\n\t":
                            v2 = append(v2, -311)
                        case "":
                            v5 -= (^(v5))
                        }
                        v42 := GenS1{F0: v3}
                        if (true && false) {
                            if true {
                                v5 += 4294967295
                                v2 = append(v2, fi39)
                                fmt.Println(fi39, v4)
                                fmt.Println(fi39)
                            } else {
                                var v43 uint16 = 15
                                _ = v43
                            }
                            if (false && false) {
                                if len(v2) > 0 {
                                    v2[0] = fi39
                                }
                            }
                            v35 += 42
                        } else {
                            var v44 int = 0
                            fmt.Println(v3)
                            v33 -= v33
                            var v45 float64 = 3.14159
                            _ = v44
                            _ = v45
                        }
                        v46 := GenS1{F0: v3}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi39
                        _ = v42
                        _ = v46
                    }
                    _ = len(v2)
                    v5 -= (v5 & (v5 - (v5 + (^(v5)))))
                }
                _ = fi36
            }
            if (1 >= 770) {
                if true {
                    v1 = (v1 | v1)
                    for fi47 := 0; fi47 < 12; fi47++ {
                        if len(v2) > 0 {
                            v2[0] = fi47
                        }
                        if ((65535 == 527) || false) {
                            fmt.Println(v3)
                            fmt.Println(v35, v4, fi47)
                        } else {
                            var v48 float32 = v4
                            if len(v2) > 0 {
                                v2[0] = 695
                            }
                            var v49 int32 = v35
                            _ = v48
                            _ = v49
                        }
                        _ = fi47
                    }
                    fmt.Println(v5, v33)
                }
                var v50 bool = (!(!false))
                _ = v50
            } else {
                fmt.Println(v4, v1)
            }
        }
        switch "a" {
        case "a":
            var v51 int64 = -321
            _ = v51
        }
        var v52 int16 = 32767
        _ = v52
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v4, v35, v4)
        var v53 float64 = 0.0
        _ = v53
    }
    switch 32768 {
    case 128:
        switch (("a" + ("hello" + ("" + "hello"))) + ("hello" + ("\n\t" + "hello"))) {
        case "hello":
            v1 -= (v1 & v1)
        case "":
            v5 = v5
        }
    }
    v33 += v33
    for fi54 := 0; fi54 < 11; fi54++ {
        var v55 int32 = v35
        _ = fi54
        _ = v55
    }
    if true {
        if true {
            var v56 string = ("foo bar" + "a")
            _ = v56
        } else {
            for fi57 := 0; fi57 < 13; fi57++ {
                var v58 float32 = v4
                v59 := GenS2{F0: 42}
                if (1.5 > 1e10) {
                    v58 = ((v58 / (-2.25 * (v58 + (0.1 - float32(1e10))))) / (((v58 - v58) + (float32(float64(0.0)) + 3.14159)) / -2.25))
                }
                _ = fi57
                _ = v58
                _ = v59
            }
        }
    } else {
        if (65535 < 377) {
            if len(v2) > 0 {
                v2[0] = -439
            }
            v60 := GenS2{F0: 9223372036854775807}
            if false {
                var v61 int32 = v35
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v62 int8 = (v3 * -116)
                _ = v61
                _ = v62
            } else {
                v3 += v3
                if (true && ((v33 * (v33 ^ 255)) <= (v33 << ((411) & 63)))) {
                    if (568 == 18446744073709551615) {
                        var v63 bool = true
                        _ = v63
                    }
                    var v64 float64 = -2.25
                    v1 = v1
                    _ = v64
                } else {
                    for fi65 := 0; fi65 < 17; fi65++ {
                        switch "\n\t" {
                        case "\n\t":
                            v33 -= v33
                        case "hello":
                            if (false || (true && false)) {
                                if true {
                                    v4 += 0.1
                                }
                                if true {
                                    var v66 string = "foo bar"
                                    _ = v66
                                }
                                v67 := GenS3{F0: 4294967295, F1: ("" + (("hello" + "foo bar") + ("foo bar" + "\n\t"))), F2: int16(v5)}
                                _ = v67
                            } else {
                                v1 += v1
                                if false {
                                    for fi68 := 0; fi68 < 15; fi68++ {
                                        if len(v2) > 0 {
                                            v2[0] = (fi65 ^ (fi65 & 128))
                                        }
                                        v69 := GenS2{F0: 0}
                                        _ = fi68
                                        _ = v69
                                    }
                                }
                            }
                        case "foo bar":
                            v70 := GenS4{F0: 255, F1: (fi65 & fi65)}
                            _ = v70
                        }
                        v71 := GenS5{F0: 0, F1: (fi65 ^ fi65)}
                        fmt.Println(v4, v1)
                        _ = fi65
                        _ = v71
                    }
                    v3 -= v3
                }
                fmt.Println(v5, v3, v3)
            }
            _ = v60
        }
        v5 = (v5 << ((78) & 63))
    }
    v72 := GenS6{F0: 1, F1: 32768, F2: 331}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v4
    _ = v5
    _ = v72
}
