package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 int16; F2 int32 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 int8; F1 uint32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 string; F1 int64 }
func genFn0() uint16 {
    var v7 string = "hello"
    v8 := GenS1{F0: 0}
    var v9 uint64 = 4294967295
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 uint, p1 uint) uint8 {
    var v10 int8 = -43
    if true {
        fmt.Println(p1)
    }
    v11 := GenS2{F0: 32767}
    v12 := []int{32767, 436}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := []int{1, 0, 433, -35}
    var v2 string = "a"
    var v3 uint8 = 13
    var v4 float32 = ((3.14159 - 0.0) * 3.14159)
    var v5 uint64 = uint64(v3)
    v6 := make(map[int]uint32)
    fmt.Println(v4, v4)
    if ((v2 < ("a" + "\n\t")) || (1e10 > float32(((((float64(v4) + 3.14159) / (3.14159 + -2.25)) - (0.1 / 1e10)) * -2.25)))) {
        v4 += 1e10
        if (177 == (v5 * (v5 >> ((18446744073709551615) & 63)))) {
            v2 = v2
        } else {
            _ = len(v6)
            v6[-1] = 0
            if (v5 < 1) {
                for fi13 := 0; fi13 < 12; fi13++ {
                    for fi14 := 0; fi14 < 6; fi14++ {
                        if (42 < 183) {
                            var v15 uint16 = 42
                            _ = v15
                        } else {
                            if true {
                                fmt.Println(v4)
                            }
                            v5 -= v5
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        _ = len(v6)
                        _ = fi14
                    }
                    _ = fi13
                }
                if (true || (-2147483648 > 442)) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v5 += v5
                    if true {
                        v5 = (v5 >> ((1) & 63))
                    } else {
                        for fi16 := 0; fi16 < 14; fi16++ {
                            v2 = "a"
                            v17 := v6[fi16]
                            _ = fi16
                            _ = v17
                        }
                        switch (v2 + v2) {
                        case "hello":
                            var v18 int64 = 42
                            _ = v18
                        case "foo bar":
                            fmt.Println(v3, v3, v3)
                        default:
                            var v19 uint8 = (v3 ^ v3)
                            _ = v19
                        }
                        for fi20 := 0; fi20 < 9; fi20++ {
                            for fi21 := 0; fi21 < 13; fi21++ {
                                v2 = v2
                                var v22 uint32 = 747
                                var v23 float32 = v4
                                var v24 int32 = -1
                                var v25 float32 = v4
                                _ = fi21
                                _ = v22
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            _ = fi20
                        }
                    }
                } else {
                    v2 = ("hello" + v2)
                    if true {
                        v26 := v6[-839]
                        if true {
                            var v27 int32 = -66
                            switch (v2 + v2) {
                            case "":
                                fmt.Println(v26, v26, v2)
                            default:
                                var v28 string = v2
                                _ = v28
                            }
                            _ = v27
                        }
                        _ = v26
                    } else {
                        if (false || false) {
                            fmt.Println(v2, v3, v2)
                            if (306 > -1) {
                                v29 := GenS1{F0: (v5 % ((42) | 1))}
                                _ = v29
                            }
                            var v30 float64 = ((1e10 + (0.0 - 0.0)) / (3.14159 / float64(3.14159)))
                            _ = v30
                        } else {
                            v3 -= (v3 << ((255) & 63))
                            fmt.Println(v5, v5, v3)
                        }
                    }
                    v1 = append(v1, 206)
                }
            }
        }
    } else {
        switch 1 {
        case 537:
            v5 = v5
        case -2147483648:
            v31 := GenS3{F0: 4294967295, F1: 32767, F2: -2147483648}
            _ = v31
        default:
            v32 := GenS4{F0: 42, F1: 0}
            _ = v32
        }
    }
    v33 := GenS5{F0: ("\n\t" + (v2 + ("\n\t" + v2))), F1: 2}
    var v34 string = ("hello" + (("a" + "") + v2))
    v3 = genFn1(400, 2)
    fmt.Println(v4)
    if false {
        switch v34 {
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "foo bar":
            v4 = -2.25
        }
        fmt.Println(v4)
        fmt.Println(v4, v34, v4)
    }
    switch 143 {
    case 953:
        for fi35 := 0; fi35 < 2; fi35++ {
            v6[(fi35 + fi35)] = uint32(v5)
            _ = fi35
        }
    case 384:
        for fi36 := 0; fi36 < 3; fi36++ {
            for fi37 := 0; fi37 < 15; fi37++ {
                v38 := GenS4{F0: 0, F1: 0}
                var v39 uint8 = (v3 ^ genFn1(921, 149))
                delete(v6, (fi36 & (-(fi37))))
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = fi36
        }
    }
    switch 32767 {
    case 2147483647:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 715:
        switch 32768 {
        case -9223372036854775808:
            for fi40 := 0; fi40 < 10; fi40++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi40
            }
        case -573:
            fmt.Println(v34, v5, v4)
        default:
            var v41 uint = 2
            _ = v41
        }
    case 4294967295:
        v4 = (v4 * v4)
    default:
        for fi42 := 0; fi42 < 12; fi42++ {
            var v43 uint32 = 323
            _ = fi42
            _ = v43
        }
    }
    switch -390 {
    case 32767:
        fmt.Println(v34, v34, v4)
    case 0:
        var v44 float32 = v4
        _ = v44
    case -1:
        switch 1 {
        case 1:
            var v45 int64 = 9223372036854775807
            _ = v45
        case 707:
            v46 := GenS4{F0: -295, F1: 1000}
            _ = v46
        }
    }
    for fi47 := 0; fi47 < 19; fi47++ {
        v3 = (v3 % ((genFn1(857, 4294967295)) | 1))
        if false {
            fmt.Println(v3)
            var v48 int64 = 316
            var v49 int = (fi47 - fi47)
            _ = v48
            _ = v49
        }
        _ = fi47
    }
    switch 126 {
    case 9223372036854775807:
        v50 := GenS6{F0: ("foo bar" + "\n\t"), F1: v5}
        _ = v50
    }
    var v51 int32 = 0
    fmt.Println(v2, v51, v3)
    var v52 int16 = 1
    v3 = (v3 & 42)
    if true {
        v53 := GenS4{F0: (v52 << ((42) & 63)), F1: (^(v52))}
        v51 = (v51 >> ((2) & 63))
        v52 += (v52 ^ (v52 ^ (v52 >> ((678) & 63))))
        _ = v53
    } else {
        fmt.Println(v5, v52, v3)
        if (true && (!(v51 >= v51))) {
            fmt.Println(v5)
            v4 = -2.25
            v6[0] = 65535
        } else {
            v51 = (v51 | -2147483648)
        }
    }
    switch 1 {
    case 473:
        for fi54 := 0; fi54 < 11; fi54++ {
            v55 := GenS7{F0: fi54, F1: -1}
            if (!(4294967295 < 4294967295)) {
                v34 = "a"
                var v56 uint8 = 128
                _ = v56
            } else {
                var v57 uint8 = genFn1(28, uint(v3))
                v34 = v34
                if true {
                    if ((v5 - v5) >= (^(v5))) {
                        switch v2 {
                        case "":
                            v58 := GenS4{F0: (v52 - -32768), F1: (v52 >> ((255) & 63))}
                            _ = v58
                        case "\n\t":
                            _ = len(v6)
                        }
                    } else {
                        v4 = v4
                        for fi59 := 0; fi59 < 4; fi59++ {
                            fmt.Println(v52)
                            v60 := GenS8{F0: 100, F1: 42}
                            _ = fi59
                            _ = v60
                        }
                        if false {
                            v4 -= v4
                            var v61 int64 = 9223372036854775807
                            var v62 int = (fi54 >> ((524) & 63))
                            _ = v61
                            _ = v62
                        } else {
                            v2 = ""
                            v2 = "\n\t"
                        }
                    }
                } else {
                    if false {
                        v63 := GenS9{F0: fi54}
                        v64 := GenS9{F0: (fi54 | fi54)}
                        _ = v63
                        _ = v64
                    }
                    var v65 uint64 = v5
                    _ = v65
                }
                _ = v57
            }
            _ = fi54
            _ = v55
        }
    case -2147483648:
        v66 := GenS10{F0: v2, F1: 42}
        _ = v66
    case -941:
        switch "foo bar" {
        case "":
            var v67 int16 = v52
            _ = v67
        case "\n\t":
            _ = len(v6)
        default:
            _ = len(v1)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
