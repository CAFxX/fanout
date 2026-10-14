package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 bool; F1 int8; F2 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 string; F1 int64; F2 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32; F1 int16 }
type GenS9 struct { F0 int32 }

func main() {
    var v1 int16 = 42
    var v2 uint = 18446744073709551615
    var v3 bool = (!true)
    v4 := []uint32{}
    v5 := make(map[uint32]string)
    var v6 int16 = v1
    var v7 uint8 = 2
    var v8 int16 = -1
    v7 -= (^(v7))
    if v3 {
        fmt.Println(v7, v7, v1)
        for fi9 := 0; fi9 < 20; fi9++ {
            for fi10 := 0; fi10 < 3; fi10++ {
                v6 += (v8 >> ((274) & 63))
                v11 := GenS1{F0: 0, F1: -1}
                var v12 uint16 = 65535
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
        if len(v4) > 0 {
            v4[0] = 2
        }
    }
    var v13 uint64 = uint64(v2)
    v4 = append(v4, 4294967295)
    v7 = v7
    if (65535 >= 2147483647) {
        if (128 != 1) {
            var v14 int = int(v2)
            _ = v14
        } else {
            var v15 int16 = int16(v13)
            var v16 int32 = -1
            var v17 uint = (v2 >> (((^(v2))) & 63))
            _ = v15
            _ = v16
            _ = v17
        }
    } else {
        if v3 {
            for fi18 := 0; fi18 < 2; fi18++ {
                if v3 {
                    for fi19 := 0; fi19 < 4; fi19++ {
                        var v20 uint16 = 32767
                        if false {
                            var v21 uint64 = 42
                            _ = v21
                        }
                        if false {
                            fmt.Println(v13, v20, v20)
                        } else {
                            if (0 < -128) {
                                fmt.Println(fi19, fi18, fi19)
                                switch "a" {
                                case "hello":
                                    if v3 {
                                        v22 := v5[42]
                                        _ = v22
                                    }
                                case "":
                                    fmt.Println(v3)
                                case "foo bar":
                                    v5[42] = "\n\t"
                                }
                            } else {
                                v2 = v2
                            }
                            fmt.Println(v8, v2)
                        }
                        _ = fi19
                        _ = v20
                    }
                    for fi23 := 0; fi23 < 2; fi23++ {
                        v5[2147483647] = "\n\t"
                        switch fi18 {
                        case 0:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            fmt.Println(v3)
                            v24 := GenS2{F0: v3, F1: -113, F2: int32(v7)}
                            _ = v24
                        case 330:
                            fmt.Println(v1, fi23)
                        }
                        switch fi23 {
                        case 1:
                            v25 := GenS1{F0: 719, F1: 1}
                            _ = v25
                        case 32767:
                            var v26 float32 = 0.1
                            _ = v26
                        }
                        _ = fi23
                    }
                }
                _ = fi18
            }
            var v27 uint = v2
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v27
        } else {
            if (1 != 0) {
                switch "" {
                case "\n\t":
                    v28 := GenS1{F0: -1, F1: 977}
                    _ = v28
                case "foo bar":
                    v29 := GenS2{F0: (!v3), F1: -128, F2: 2147483647}
                    _ = v29
                }
                for fi30 := 0; fi30 < 14; fi30++ {
                    var v31 uint64 = 18446744073709551615
                    for fi32 := 0; fi32 < 1; fi32++ {
                        _ = len(v5)
                        if v3 {
                            v1 -= v8
                            var v33 uint8 = v7
                            _ = v33
                        }
                        _ = fi32
                    }
                    v6 += v8
                    _ = fi30
                    _ = v31
                }
                v1 = 1000
            }
            v2 = v2
        }
        if v3 {
            if v3 {
                v5[4294967295] = "\n\t"
                for fi34 := 0; fi34 < 7; fi34++ {
                    delete(v5, 374)
                    v8 -= v8
                    var v35 bool = (-1 == -107)
                    _ = fi34
                    _ = v35
                }
                for fi36 := 0; fi36 < 7; fi36++ {
                    fmt.Println(v6, v7, v13)
                    v13 = v13
                    _ = fi36
                }
            } else {
                var v37 uint = (v2 | 613)
                v38 := GenS3{F0: 2}
                _ = v37
                _ = v38
            }
        } else {
            v2 = (v2 % ((190) | 1))
            for fi39 := 0; fi39 < 15; fi39++ {
                v3 = ((v3 || (v3 && v3)) || v3)
                var v40 string = "hello"
                for fi41 := 0; fi41 < 17; fi41++ {
                    var v42 int = (fi41 * (fi39 & 32767))
                    v43 := GenS1{F0: 0, F1: -99}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                _ = fi39
                _ = v40
            }
            v44 := GenS1{F0: -1, F1: 519}
            _ = v44
        }
        v8 = (v6 + v1)
    }
    _ = len(v5)
    v45 := GenS4{F0: -326, F1: 917}
    fmt.Println(v1)
    if (v3 || v3) {
        for fi46 := 0; fi46 < 11; fi46++ {
            var v47 uint8 = 42
            v13 = 568
            _ = fi46
            _ = v47
        }
    } else {
        if v3 {
            var v48 int16 = (v8 >> (((v2 << ((v2) & 63))) & 63))
            fmt.Println(v48)
            _ = v48
        } else {
            v4 = append(v4, 4294967295)
            if ((v2 << ((uint(v7)) & 63)) == (v2 ^ (v2 / ((255) | 1)))) {
                v13 = uint64(v7)
                if v3 {
                    if v3 {
                        v49 := GenS5{F0: "a", F1: -1, F2: (v8 >> ((v2) & 63))}
                        for fi50 := 0; fi50 < 7; fi50++ {
                            if false {
                                if (float32(0.0) > 3.14159) {
                                    v7 = v7
                                    v8 = (v6 % ((v8) | 1))
                                } else {
                                    fmt.Println(v3)
                                }
                                switch ("a" + (("hello" + "hello") + (("foo bar" + "hello") + ("\n\t" + "hello")))) {
                                case "foo bar":
                                    for fi51 := 0; fi51 < 4; fi51++ {
                                        var v52 float32 = (1e10 + (3.14159 - (float32(-2.25) * float32(0.0))))
                                        v1 += v8
                                        _ = fi51
                                        _ = v52
                                    }
                                case "a":
                                    switch fi50 {
                                    case -661:
                                        for fi53 := 0; fi53 < 8; fi53++ {
                                            v2 += (^(v2))
                                            _ = fi53
                                        }
                                    default:
                                        var v54 uint = v2
                                        _ = v54
                                    }
                                default:
                                    var v55 uint8 = v7
                                    _ = v55
                                }
                                switch fi50 {
                                case 890:
                                    if true {
                                        v56 := GenS1{F0: -2147483648, F1: -502}
                                        _ = v56
                                    } else {
                                        if v3 {
                                            for fi57 := 0; fi57 < 14; fi57++ {
                                                fmt.Println(v13)
                                                var v58 bool = ((fi57 / ((160) | 1)) > (fi57 >> ((v2) & 63)))
                                                _ = fi57
                                                _ = v58
                                            }
                                            v59 := GenS6{F0: (^(fi50))}
                                            _ = v59
                                        } else {
                                            if (!v3) {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                if true {
                                                    var v60 uint16 = 32767
                                                    _ = v60
                                                } else {
                                                    var v61 int16 = (^(v8))
                                                    fmt.Println(v61, v61, v7)
                                                    fmt.Println(fi50, v6, v3)
                                                    _ = v61
                                                }
                                            }
                                            v62 := GenS2{F0: v3, F1: -1, F2: -101}
                                            _ = v62
                                        }
                                    }
                                default:
                                    if true {
                                        v63 := GenS6{F0: (fi50 | int(v13))}
                                        var v64 float32 = -2.25
                                        if (-1 > 42) {
                                            v8 -= v6
                                            var v65 int32 = -1
                                            _ = v65
                                        } else {
                                            v8 = 42
                                        }
                                        _ = v63
                                        _ = v64
                                    }
                                }
                            } else {
                                var v66 int = (fi50 / (((fi50 & (^(fi50)))) | 1))
                                v67 := GenS4{F0: 127, F1: (fi50 ^ 1)}
                                _ = v66
                                _ = v67
                            }
                            fmt.Println(v8, v7)
                            v6 = (^(v6))
                            _ = fi50
                        }
                        v68 := GenS7{F0: 18446744073709551615}
                        _ = v49
                        _ = v68
                    } else {
                        v69 := v5[955]
                        v2 += v2
                        _ = v69
                    }
                    v70 := GenS5{F0: "", F1: -835, F2: v8}
                    v3 = v3
                    _ = v70
                }
                _ = len(v4)
            }
            var v71 int64 = -9223372036854775808
            _ = v71
        }
        for fi72 := 0; fi72 < 18; fi72++ {
            for fi73 := 0; fi73 < 2; fi73++ {
                fmt.Println(v7, v8, v7)
                v6 += (v1 + v1)
                fmt.Println(v2, fi73)
                _ = fi73
            }
            var v74 float64 = (1.5 * ((((1e10 + -2.25) + ((0.1 * 1e10) - 1e10)) - 3.14159) + ((0.0 * 0.0) / 1.0)))
            v8 = v6
            _ = fi72
            _ = v74
        }
        switch int(v7) {
        case 32767:
            v75 := GenS1{F0: -719, F1: -9223372036854775808}
            _ = v75
        case 566:
            v2 = 259
        default:
            v2 += v2
        }
    }
    switch "foo bar" {
    case "a":
        var v76 uint8 = v7
        _ = v76
    case "hello":
        var v77 uint = v2
        _ = v77
    case "\n\t":
        if (true && true) {
            v1 = 42
        }
    }
    for fi78 := 0; fi78 < 17; fi78++ {
        v79 := GenS8{F0: 42, F1: v6}
        fmt.Println(v13, fi78, v13)
        v80 := v5[42]
        _ = fi78
        _ = v79
        _ = v80
    }
    v8 += (-(v8))
    if false {
        if ("a" > "foo bar") {
            if ((!((((v7 | v7) > (^(v7))) || v3) && true)) && v3) {
                if (("foo bar" + ("a" + ("" + "a"))) > (("hello" + "foo bar") + "")) {
                    v2 += (v2 << ((v2) & 63))
                    if (0 != (v2 | v2)) {
                        v81 := GenS9{F0: 0}
                        _ = v81
                    } else {
                        var v82 float32 = 0.1
                        v3 = v3
                        var v83 float64 = ((-2.25 * 0.0) / 1e10)
                        _ = v82
                        _ = v83
                    }
                } else {
                    v13 = 4294967295
                }
                v7 = (v7 % (((v7 >> ((797) & 63))) | 1))
            } else {
                v7 += 82
                var v84 float64 = 1.5
                v85 := GenS7{F0: (v2 ^ (v2 >> (((v2 << ((v2) & 63))) & 63)))}
                _ = v84
                _ = v85
            }
        }
    } else {
        var v86 uint64 = 682
        v3 = false
        var v87 int64 = 42
        _ = v86
        _ = v87
    }
    v88 := GenS3{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v88
}
