package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint; F2 uint }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 uint; F1 int16; F2 string }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32; F1 int16 }

func main() {
    var v1 uint8 = 231
    var v2 int = -9223372036854775808
    v3 := []string{("a" + "\n\t")}
    v4 := []uint32{65535, 2147483648, 1}
    switch 32768 {
    case 32767:
        v1 += (v1 & v1)
    case 2147483647:
        v1 += (v1 ^ v1)
    }
    v2 = (v2 << ((65535) & 63))
    fmt.Println(v1)
    v1 = (^(v1))
    var v5 int32 = int32(v1)
    v2 = v2
    if (99 <= 42) {
        v5 += v5
    } else {
        for fi6 := 0; fi6 < 10; fi6++ {
            fmt.Println(fi6, v1, v1)
            var v7 int8 = -1
            if true {
                v4 = append(v4, 42)
                v7 = v7
            } else {
                for fi8 := 0; fi8 < 15; fi8++ {
                    switch (fi8 + fi6) {
                    case -2147483648:
                        switch ("" + "a") {
                        case "":
                            if (false || false) {
                                if (3.14159 != 1e10) {
                                    var v9 uint16 = 2
                                    _ = v9
                                } else {
                                    var v10 uint64 = 234
                                    fmt.Println(fi6, fi8)
                                    v2 -= (fi8 & -404)
                                    _ = v10
                                }
                            }
                        case "hello":
                            v7 += (v7 - 127)
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case 42:
                        var v11 bool = (!(true || false))
                        _ = v11
                    case -840:
                        if true {
                            v12 := GenS1{F0: int32(v2)}
                            switch 32768 {
                            case 128:
                                switch "a" {
                                case "foo bar":
                                    switch "a" {
                                    case "\n\t":
                                        if (-1 >= 0) {
                                            v2 -= 4294967295
                                            v4 = append(v4, uint32(v1))
                                        } else {
                                            var v13 uint16 = 547
                                            switch (v2 | -949) {
                                            case 4294967295:
                                                var v14 uint16 = v13
                                                _ = v14
                                            case 38:
                                                var v15 uint = 2
                                                _ = v15
                                            }
                                            if (int(v5) == v2) {
                                                var v16 uint32 = 42
                                                fmt.Println(fi6, v2)
                                                _ = v16
                                            } else {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v17 := GenS2{F0: 9223372036854775807, F1: 65535, F2: 65535}
                                                v4 = append(v4, 2147483647)
                                                _ = len(v3)
                                                if false {
                                                    v18 := GenS1{F0: (v5 * 292)}
                                                    v2 -= fi6
                                                    _ = v18
                                                }
                                                _ = v17
                                            }
                                            _ = v13
                                        }
                                    case "a":
                                        if (844 >= 42) {
                                            var v19 float32 = ((1e10 - 3.14159) * 3.14159)
                                            _ = v19
                                        } else {
                                            var v20 uint32 = 2147483647
                                            if false {
                                                v7 = v7
                                            } else {
                                                var v21 int16 = 42
                                                if ((v20 <= (v20 >> ((39) & 63))) && true) {
                                                    var v22 int64 = 0
                                                    var v23 uint32 = 2147483647
                                                    fmt.Println(v7, v5, v5)
                                                    _ = v22
                                                    _ = v23
                                                } else {
                                                    v7 = v7
                                                }
                                                if ((v20 % ((v20) | 1)) != (v20 % ((v20) | 1))) {
                                                    var v24 int16 = (v21 ^ -1)
                                                    v25 := GenS3{F0: 2, F1: v5}
                                                    _ = v24
                                                    _ = v25
                                                } else {
                                                    var v26 int32 = v5
                                                    _ = v26
                                                }
                                                v27 := GenS3{F0: 18446744073709551615, F1: -100}
                                                _ = v21
                                                _ = v27
                                            }
                                            v20 = v20
                                            _ = v20
                                        }
                                    default:
                                        v7 -= v7
                                    }
                                default:
                                    var v28 string = "hello"
                                    _ = v28
                                }
                            case 349:
                                var v29 int16 = -32768
                                _ = v29
                            case 263:
                                if (v7 <= v7) {
                                    fmt.Println(v5)
                                }
                            default:
                                v2 += (v2 ^ (-(fi6)))
                            }
                            _ = v12
                        } else {
                            v4 = append(v4, 42)
                        }
                    default:
                        v30 := GenS3{F0: 255, F1: -890}
                        _ = v30
                    }
                    _ = fi8
                }
                v31 := GenS4{F0: "\n\t", F1: 577}
                _ = v31
            }
            _ = fi6
            _ = v7
        }
        var v32 float64 = 1e10
        v5 += (v5 + (v5 >> ((42) & 63)))
        _ = v32
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        if (205 > 271) {
            v2 += fi33
            if true {
                var v34 int64 = -1
                v35 := GenS5{F0: 113, F1: 1000, F2: (("\n\t" + "hello") + "")}
                for fi36 := 0; fi36 < 10; fi36++ {
                    v1 -= 128
                    switch fi33 {
                    case 32768:
                        switch "a" {
                        case "foo bar":
                            if ((v5 << ((686) & 63)) <= (-(v5))) {
                                var v37 int8 = 0
                                _ = v37
                            }
                        case "hello":
                            if len(v3) > 0 {
                                v3[0] = ("a" + "\n\t")
                            }
                        default:
                            switch (v2 * 544) {
                            case 350:
                                var v38 int = fi33
                                var v39 uint8 = v1
                                v40 := GenS5{F0: 877, F1: 0, F2: "\n\t"}
                                _ = v38
                                _ = v39
                                _ = v40
                            default:
                                if true {
                                    fmt.Println(fi36)
                                    fmt.Println(v2, v5, fi33)
                                    if false {
                                        v2 -= 910
                                    } else {
                                        var v41 int = fi36
                                        _ = v41
                                    }
                                }
                            }
                        }
                    case 392:
                        v34 = 42
                    default:
                        fmt.Println(v1, fi36, fi36)
                    }
                    fmt.Println(v2, v5)
                    _ = fi36
                }
                _ = v34
                _ = v35
            } else {
                for fi42 := 0; fi42 < 18; fi42++ {
                    v43 := GenS1{F0: v5}
                    v5 -= (^(v5))
                    _ = fi42
                    _ = v43
                }
            }
            _ = len(v3)
        }
        for fi44 := 0; fi44 < 20; fi44++ {
            fmt.Println(v5)
            var v45 int = v2
            _ = fi44
            _ = v45
        }
        _ = fi33
    }
    fmt.Println(v1, v5, v1)
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    if (((-2.25 * float64((0.0 + 0.1))) * -2.25) == (-2.25 + 0.0)) {
        for fi46 := 0; fi46 < 2; fi46++ {
            if true {
                switch "foo bar" {
                case "a":
                    if true {
                        for fi47 := 0; fi47 < 11; fi47++ {
                            v2 = 4294967295
                            _ = fi47
                        }
                        v1 += (v1 << ((1) & 63))
                        var v48 int32 = -526
                        _ = v48
                    } else {
                        for fi49 := 0; fi49 < 9; fi49++ {
                            if false {
                                var v50 uint = 2
                                _ = v50
                            }
                            _ = len(v3)
                            _ = fi49
                        }
                        v51 := GenS6{F0: 9223372036854775807}
                        for fi52 := 0; fi52 < 12; fi52++ {
                            var v53 int8 = 62
                            _ = fi52
                            _ = v53
                        }
                        _ = v51
                    }
                default:
                    fmt.Println(v1, fi46, v5)
                }
                v5 += v5
            }
            v5 -= (^(v5))
            _ = fi46
        }
        switch "a" {
        case "":
            for fi54 := 0; fi54 < 15; fi54++ {
                var v55 uint8 = (v1 + (v1 / ((v1) | 1)))
                switch v2 {
                case 9223372036854775807:
                    for fi56 := 0; fi56 < 9; fi56++ {
                        v55 = (v1 - (v55 / (((v1 / ((2) | 1))) | 1)))
                        var v57 float32 = (1e10 * 0.1)
                        _ = fi56
                        _ = v57
                    }
                case 2147483647:
                    fmt.Println(v55)
                default:
                    switch ("foo bar" + (("" + ("" + "")) + (("hello" + "a") + "\n\t"))) {
                    case "a":
                        for fi58 := 0; fi58 < 16; fi58++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = len(v4)
                            fmt.Println(v55, v1, v1)
                            _ = fi58
                        }
                    default:
                        if len(v3) > 0 {
                            v3[0] = (("foo bar" + ("" + "hello")) + "")
                        }
                    }
                }
                _ = fi54
                _ = v55
            }
        case "a":
            if false {
                switch (("" + ("a" + "hello")) + (("hello" + (((("" + (("a" + (("" + "\n\t") + "foo bar")) + ("" + ((("hello" + "a") + "a") + "")))) + (("" + (("foo bar" + ("" + (("hello" + ("hello" + ("" + (((("foo bar" + ("a" + "foo bar")) + "hello") + "") + "foo bar")))) + "hello"))) + "foo bar")) + "\n\t")) + "") + ("\n\t" + "hello"))) + ("" + ("foo bar" + (("hello" + "hello") + "hello"))))) {
                case "":
                    var v59 uint = 0
                    _ = v59
                }
                var v60 uint64 = 4294967295
                _ = v60
            } else {
                var v61 uint8 = (v1 + v1)
                _ = v61
            }
        case "\n\t":
            switch "a" {
            case "foo bar":
                v3 = append(v3, "")
            case "a":
                var v62 int8 = -55
                _ = v62
            }
        }
    } else {
        v1 -= 0
        var v63 uint = 982
        _ = v63
    }
    for fi64 := 0; fi64 < 4; fi64++ {
        v1 = v1
        _ = fi64
    }
    if (v5 <= v5) {
        var v65 int8 = 88
        var v66 uint32 = 4294967295
        var v67 uint16 = uint16(v66)
        _ = v65
        _ = v66
        _ = v67
    }
    for fi68 := 0; fi68 < 3; fi68++ {
        var v69 int = (^(fi68))
        v69 += fi68
        _ = fi68
        _ = v69
    }
    v70 := GenS2{F0: -3, F1: 562, F2: 47}
    if len(v3) > 0 {
        v3[0] = ("a" + ("a" + "a"))
    }
    _ = len(v3)
    for fi71 := 0; fi71 < 9; fi71++ {
        if (881 < 0) {
            _ = len(v4)
        } else {
            v72 := GenS6{F0: 0}
            switch "" {
            case "":
                var v73 uint16 = 255
                _ = v73
            case "hello":
                fmt.Println(fi71)
            case "\n\t":
                for fi74 := 0; fi74 < 11; fi74++ {
                    switch "hello" {
                    case "":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (uint32(v1) <= 2147483647) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        fmt.Println(v5)
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = ""
                        }
                    }
                    _ = fi74
                }
            }
            for fi75 := 0; fi75 < 16; fi75++ {
                var v76 float64 = 1e10
                if false {
                    var v77 uint64 = 4294967295
                    _ = v77
                } else {
                    var v78 int64 = -76
                    _ = v78
                }
                _ = fi75
                _ = v76
            }
            _ = v72
        }
        v79 := GenS7{F0: 2147483647, F1: 0}
        v80 := GenS6{F0: -400}
        _ = fi71
        _ = v79
        _ = v80
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v70
}
