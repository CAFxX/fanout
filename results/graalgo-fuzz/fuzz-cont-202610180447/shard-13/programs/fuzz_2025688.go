package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 int; F1 int8; F2 uint64 }
type GenS7 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint64) int32 {
    if true {
        fmt.Println(p0)
        p0 += 379
    } else {
        v7 := GenS1{F0: p0, F1: 2147483647}
        for fi8 := 0; fi8 < 1; fi8++ {
            if (0 < 1) {
                v9 := GenS1{F0: p0, F1: -2147483648}
                v10 := []int{9223372036854775807}
                for fi11 := 0; fi11 < 18; fi11++ {
                    v12 := GenS1{F0: p0, F1: 2147483647}
                    p0 = (p0 * p0)
                    p0 = p0
                    _ = fi11
                    _ = v12
                }
                _ = v10
                _ = v9
            }
            var v13 uint8 = 0
            _ = fi8
            _ = v13
        }
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    return -1
}

func main() {
    var v1 uint64 = 541
    var v2 bool = ("hello" < ((((("foo bar" + "\n\t") + "a") + (("\n\t" + "a") + ("" + "foo bar"))) + (("hello" + ("\n\t" + "")) + "hello")) + ("foo bar" + ("a" + "hello"))))
    var v3 bool = (49 == 0)
    v4 := []int{4294967295, int(v1)}
    var v5 float64 = 3.14159
    var v6 int32 = 42
    v14 := GenS1{F0: (v1 >> ((65535) & 63)), F1: (v6 ^ 163)}
    if false {
        v2 = false
    } else {
        fmt.Println(v1)
        if v3 {
            if true {
                for fi15 := 0; fi15 < 5; fi15++ {
                    switch "\n\t" {
                    case "a":
                        var v16 uint64 = (v1 << ((2) & 63))
                        _ = v16
                    }
                    _ = len(v4)
                    var v17 uint = 42
                    _ = fi15
                    _ = v17
                }
                fmt.Println(v6, v1)
            }
            for fi18 := 0; fi18 < 9; fi18++ {
                for fi19 := 0; fi19 < 5; fi19++ {
                    v3 = true
                    if len(v4) > 0 {
                        v4[0] = (fi19 >> ((2) & 63))
                    }
                    _ = fi19
                }
                if (42 != -83) {
                    for fi20 := 0; fi20 < 12; fi20++ {
                        var v21 float64 = float64(0.0)
                        v22 := GenS1{F0: (v1 | (^(v1))), F1: (v6 - v6)}
                        var v23 int16 = -1
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    var v24 int = (fi18 * (^(fi18)))
                    v24 += (v24 + 32767)
                    _ = v24
                } else {
                    v6 -= genFn0((v1 >> ((111) & 63)))
                    if (!v2) {
                        for fi25 := 0; fi25 < 19; fi25++ {
                            var v26 float32 = 0.0
                            var v27 int16 = int16(v6)
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        v28 := GenS2{F0: genFn0(v1), F1: 255, F2: 4294967295}
                        fmt.Println(v1)
                        _ = v28
                    }
                }
                for fi29 := 0; fi29 < 17; fi29++ {
                    v30 := GenS3{F0: (v6 % ((0) | 1)), F1: 65535}
                    _ = fi29
                    _ = v30
                }
                _ = fi18
            }
            if false {
                v31 := GenS3{F0: v6, F1: 2147483648}
                _ = v31
            }
        }
        if false {
            for fi32 := 0; fi32 < 5; fi32++ {
                v33 := GenS3{F0: -1, F1: 1}
                v6 -= 2147483647
                var v34 int = (^(fi32))
                _ = fi32
                _ = v33
                _ = v34
            }
            var v35 bool = (0 <= 18446744073709551615)
            v36 := GenS2{F0: (v6 | (v6 << ((609) & 63))), F1: 1, F2: 646}
            _ = v35
            _ = v36
        } else {
            v37 := GenS4{F0: v1}
            _ = v37
        }
    }
    v1 += 517
    v38 := GenS2{F0: int32(v1), F1: 127, F2: 65535}
    fmt.Println(v1, v2)
    if (true && (!((!((-(v6)) <= (^(v6)))) && false))) {
        var v39 uint8 = 1
        for fi40 := 0; fi40 < 20; fi40++ {
            var v41 bool = (0.0 != 0.0)
            if (!v41) {
                var v42 uint8 = (v39 + v39)
                if (true && v3) {
                    fmt.Println(v5)
                    v2 = false
                    var v43 int = (fi40 & (fi40 & fi40))
                    _ = v43
                }
                _ = v42
            }
            if (255 == 4294967295) {
                for fi44 := 0; fi44 < 4; fi44++ {
                    var v45 uint64 = uint64(v39)
                    var v46 uint = 18446744073709551615
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            }
            _ = fi40
            _ = v41
        }
        v47 := GenS5{F0: 1, F1: 65535}
        _ = v39
        _ = v47
    }
    if true {
        if (!v2) {
            v48 := GenS1{F0: v1, F1: (v6 * (v6 / (((-(v6))) | 1)))}
            for fi49 := 0; fi49 < 9; fi49++ {
                switch ("" + ("foo bar" + "foo bar")) {
                case "":
                    v50 := GenS4{F0: (v1 | (v1 >> ((233) & 63)))}
                    _ = v50
                case "\n\t":
                    var v51 string = (("hello" + "a") + "a")
                    _ = v51
                case "hello":
                    for fi52 := 0; fi52 < 7; fi52++ {
                        var v53 uint = 743
                        _ = fi52
                        _ = v53
                    }
                }
                if (!true) {
                    if (1 >= -68) {
                        for fi54 := 0; fi54 < 6; fi54++ {
                            if ((2147483647 <= 392) && v2) {
                                v2 = false
                            } else {
                                v3 = v2
                            }
                            _ = fi54
                        }
                    }
                    switch (("\n\t" + "") + "") {
                    case "":
                        v3 = (!(!(true && v3)))
                    case "a":
                        v4 = append(v4, (fi49 ^ fi49))
                    default:
                        v2 = true
                    }
                }
                fmt.Println(fi49, v5)
                _ = fi49
            }
            v5 = 0.1
            _ = v48
        } else {
            if v2 {
                if v3 {
                    v55 := GenS1{F0: v1, F1: -1}
                    if (v3 && true) {
                        var v56 int8 = 100
                        if (133 == 255) {
                            var v57 int = 2147483647
                            _ = v57
                        }
                        _ = v56
                    } else {
                        if false {
                            var v58 string = ""
                            fmt.Println(v5, v6)
                            var v59 int32 = (v6 << ((7) & 63))
                            _ = v58
                            _ = v59
                        } else {
                            for fi60 := 0; fi60 < 20; fi60++ {
                                var v61 float64 = (v5 / 1.0)
                                for fi62 := 0; fi62 < 5; fi62++ {
                                    v63 := GenS1{F0: v1, F1: (-(v6))}
                                    _ = fi62
                                    _ = v63
                                }
                                fmt.Println(v5, v6, v2)
                                _ = fi60
                                _ = v61
                            }
                            for fi64 := 0; fi64 < 4; fi64++ {
                                for fi65 := 0; fi65 < 11; fi65++ {
                                    if len(v4) > 0 {
                                        v4[0] = (fi64 % ((fi64) | 1))
                                    }
                                    var v66 uint8 = 0
                                    _ = fi65
                                    _ = v66
                                }
                                _ = fi64
                            }
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v67 := GenS3{F0: (v6 + v6), F1: 2147483648}
                        _ = v67
                    }
                    _ = v55
                } else {
                    switch 32768 {
                    case -9223372036854775808:
                        var v68 bool = false
                        _ = v68
                    }
                    if (42 != 42) {
                        if len(v4) > 0 {
                            v4[0] = 42
                        }
                        v69 := GenS3{F0: v6, F1: 2147483647}
                        v1 = 0
                        _ = v69
                    } else {
                        var v70 int = 1
                        v71 := GenS3{F0: -944, F1: 1}
                        for fi72 := 0; fi72 < 17; fi72++ {
                            v3 = (v3 || (v3 || v3))
                            _ = fi72
                        }
                        _ = v70
                        _ = v71
                    }
                    var v73 int = 32768
                    _ = v73
                }
                switch "foo bar" {
                case "":
                    var v74 bool = (42 < 0)
                    _ = v74
                case "hello":
                    v5 += (((1e10 / v5) * v5) * ((0.1 - v5) * (v5 / v5)))
                case "foo bar":
                    if v3 {
                        if v2 {
                            var v75 float32 = 3.14159
                            var v76 uint32 = 65535
                            _ = v75
                            _ = v76
                        } else {
                            if ((v6 == v6) || (v3 && true)) {
                                switch (("hello" + "") + "foo bar") {
                                case "foo bar":
                                    v2 = true
                                }
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = 515
                                }
                                if (254 == 20) {
                                    if v3 {
                                        v4 = append(v4, 477)
                                        v5 = v5
                                        v6 = (v6 % (((v6 | genFn0((v1 ^ 822)))) | 1))
                                    } else {
                                        v77 := GenS6{F0: -789, F1: -11, F2: (v1 & (v1 * v1))}
                                        if v3 {
                                            var v78 uint32 = 2147483647
                                            _ = v78
                                        } else {
                                            fmt.Println(v6)
                                            v5 = v5
                                        }
                                        _ = v77
                                    }
                                    v2 = true
                                    fmt.Println(v5)
                                }
                            }
                            v4 = append(v4, 42)
                            _ = len(v4)
                        }
                        var v79 float64 = float64((((1.5 * (((0.1 - (0.0 * 1e10)) - 0.0) + float32(1.5))) + (((1.5 + 0.1) + (0.1 / 1.5)) + (((-2.25 - ((3.14159 + (((((0.0 * 0.1) * float32(v5)) + 1.5) - 0.0) + 3.14159)) - 1.5)) + (3.14159 - 1e10)) + ((0.0 / 3.14159) + 0.1)))) * (float32((3.14159 + (v5 * (-2.25 / v5)))) - float32(v5))))
                        v6 = (v6 - 2147483647)
                        _ = v79
                    }
                default:
                    for fi80 := 0; fi80 < 13; fi80++ {
                        for fi81 := 0; fi81 < 19; fi81++ {
                            if v2 {
                                v82 := GenS5{F0: 42, F1: 4294967295}
                                fmt.Println(v3, v6)
                                var v83 int64 = 0
                                if v3 {
                                    v4 = append(v4, fi81)
                                    if (!(255 <= 440)) {
                                        var v84 float32 = 1e10
                                        _ = v84
                                    }
                                    v83 = v83
                                } else {
                                    v85 := GenS5{F0: 471, F1: 255}
                                    _ = v85
                                }
                                v1 += v1
                                _ = v82
                                _ = v83
                            } else {
                                v6 = 1
                            }
                            v2 = v2
                            var v86 int16 = 1
                            var v87 uint8 = 42
                            _ = fi81
                            _ = v86
                            _ = v87
                        }
                        var v88 uint16 = 42
                        v1 += 629
                        _ = fi80
                        _ = v88
                    }
                }
            } else {
                fmt.Println(v2, v1, v3)
                var v89 int32 = (^(v6))
                _ = v89
            }
            var v90 uint32 = 42
            _ = v90
        }
        v4 = append(v4, -9223372036854775808)
        _ = len(v4)
    }
    v91 := GenS7{F0: 0, F1: 0}
    v6 -= (v6 >> ((2) & 63))
    v1 -= v1
    if v3 {
        var v92 int8 = 97
        _ = v92
    } else {
        v93 := GenS3{F0: (v6 % ((int32(v1)) | 1)), F1: 2147483648}
        _ = v93
    }
    for fi94 := 0; fi94 < 18; fi94++ {
        fmt.Println(v6, v6, v5)
        _ = fi94
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v91
}
