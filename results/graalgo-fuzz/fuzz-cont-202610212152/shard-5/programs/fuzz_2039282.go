package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 uint64; F1 int32; F2 uint }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16; F1 uint64; F2 int16 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint16; F1 bool; F2 int64 }

func main() {
    var v1 string = "\n\t"
    var v2 float64 = 0.1
    var v3 int32 = 0
    v4 := []int{2147483647, 42, 886, -664}
    v5 := []string{}
    var v6 uint8 = 139
    var v7 int16 = -32768
    switch v1 {
    case "a":
        var v8 float32 = float32(0.1)
        _ = v8
    case "hello":
        v3 = 0
    case "foo bar":
        if false {
            v7 += (-(v7))
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := GenS1{F0: 742}
                _ = fi9
                _ = v10
            }
            _ = len(v4)
        }
    }
    switch v1 {
    case "\n\t":
        switch "" {
        case "\n\t":
            v3 += v3
        case "foo bar":
            switch "a" {
            case "hello":
                if false {
                    if false {
                        for fi11 := 0; fi11 < 17; fi11++ {
                            v12 := GenS2{F0: (700 < 18446744073709551615), F1: 255, F2: 0}
                            if (243 != 174) {
                                var v13 int32 = (-(v3))
                                v4 = append(v4, (fi11 - fi11))
                                fmt.Println(v2)
                                _ = v13
                            } else {
                                v14 := GenS3{F0: 2}
                                if len(v5) > 0 {
                                    v5[0] = v1
                                }
                                _ = v14
                            }
                            _ = fi11
                            _ = v12
                        }
                        if false {
                            if len(v5) > 0 {
                                v5[0] = ((((((v1 + "hello") + "hello") + ((v1 + "") + v1)) + v1) + ((v1 + ((v1 + "\n\t") + ((v1 + "\n\t") + v1))) + v1)) + v1)
                            }
                            v15 := GenS4{F0: 9223372036854775807, F1: 2}
                            _ = v15
                        } else {
                            _ = len(v5)
                            var v16 int8 = 1
                            for fi17 := 0; fi17 < 20; fi17++ {
                                for fi18 := 0; fi18 < 1; fi18++ {
                                    v19 := GenS4{F0: int64(v16), F1: 0}
                                    fmt.Println(v6, v7)
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = fi18
                                    _ = v19
                                }
                                _ = fi17
                            }
                            _ = v16
                        }
                        switch v1 {
                        case "a":
                            if (685 == 928) {
                                switch "hello" {
                                case "foo bar":
                                    if ((0.1 / -2.25) < ((1e10 + (1.5 * float32((0.0 * v2)))) - 0.1)) {
                                        v6 -= (v6 / ((v6) | 1))
                                    }
                                }
                                var v20 float32 = (0.0 + 0.0)
                                _ = v20
                            } else {
                                fmt.Println(v3, v7)
                            }
                        case "foo bar":
                            if (!(-2.25 < 1.5)) {
                                switch 32768 {
                                case 754:
                                    var v21 int = -9223372036854775808
                                    _ = v21
                                }
                                var v22 int8 = 127
                                _ = v22
                            }
                        }
                    }
                    fmt.Println(v3)
                }
            case "foo bar":
                var v23 string = v1
                _ = v23
            case "":
                if ((-1 == -128) && true) {
                    var v24 int = 127
                    var v25 int = 1
                    fmt.Println(v24, v1, v25)
                    _ = v24
                    _ = v25
                } else {
                    v6 += v6
                    var v26 int64 = -9223372036854775808
                    var v27 int8 = 100
                    _ = v26
                    _ = v27
                }
            }
        }
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = -1
        }
    case "a":
        if false {
            var v28 int = 133
            for fi29 := 0; fi29 < 16; fi29++ {
                v30 := GenS3{F0: 4294967295}
                _ = fi29
                _ = v30
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v28
        }
    }
    if (18446744073709551615 != 2) {
        var v31 int = -860
        if len(v4) > 0 {
            v4[0] = (v31 >> ((255) & 63))
        }
        _ = v31
    } else {
        fmt.Println(v2, v3)
        if true {
            var v32 int8 = 100
            v33 := GenS5{F0: v7, F1: 42}
            _ = v32
            _ = v33
        }
        if true {
            var v34 float32 = ((((1e10 - 0.0) / 1.5) + (3.14159 * float32(0.1))) * (-2.25 * float32(v2)))
            if (((v1 + (v1 + (((v1 + ((v1 + v1) + ("foo bar" + ("\n\t" + "")))) + "a") + "foo bar"))) == (((v1 + ("\n\t" + "a")) + (v1 + (v1 + ((v1 + (v1 + v1)) + ("\n\t" + ("\n\t" + (v1 + v1))))))) + v1)) && false) {
                fmt.Println(v2, v2)
                v3 = v3
                fmt.Println(v1, v1)
            } else {
                if false {
                    v35 := GenS6{F0: 1, F1: 944, F2: 988}
                    _ = len(v4)
                    _ = v35
                } else {
                    if len(v4) > 0 {
                        v4[0] = 0
                    }
                    _ = len(v5)
                    fmt.Println(v7, v34)
                }
                fmt.Println(v6, v34)
            }
            var v36 int64 = -1
            _ = v34
            _ = v36
        } else {
            v5 = append(v5, v1)
            v2 = v2
        }
    }
    if true {
        fmt.Println(v7, v1)
        switch 2147483647 {
        case 789:
            v1 = "foo bar"
        case 32767:
            for fi37 := 0; fi37 < 14; fi37++ {
                if len(v4) > 0 {
                    v4[0] = 0
                }
                fmt.Println(v6, v6)
                _ = fi37
            }
        default:
            switch -9223372036854775808 {
            case 42:
                var v38 bool = (!false)
                _ = v38
            case -1:
                switch 32767 {
                case 200:
                    v3 += (v3 ^ v3)
                default:
                    if ((18446744073709551615 != 541) && false) {
                        for fi39 := 0; fi39 < 6; fi39++ {
                            switch (fi39 ^ (fi39 / (((-(fi39))) | 1))) {
                            case -2147483648:
                                var v40 float64 = v2
                                _ = v40
                            default:
                                v1 = (v1 + ("hello" + v1))
                            }
                            for fi41 := 0; fi41 < 4; fi41++ {
                                fmt.Println(fi41, v3)
                                var v42 uint32 = 4294967295
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = fi41
                                _ = v42
                            }
                            _ = fi39
                        }
                        if len(v4) > 0 {
                            v4[0] = 128
                        }
                    }
                }
            case 2147483647:
                v6 += v6
            default:
                v3 -= (v3 | (v3 >> ((2) & 63)))
            }
        }
    } else {
        v43 := GenS3{F0: 484}
        v7 -= v7
        _ = v43
    }
    if (192 >= 18446744073709551615) {
        var v44 float32 = 3.14159
        _ = v44
    } else {
        var v45 float64 = (v2 / 1.0)
        if (189 < 129) {
            for fi46 := 0; fi46 < 8; fi46++ {
                var v47 int16 = (v7 - 42)
                v6 += 127
                _ = len(v4)
                _ = fi46
                _ = v47
            }
            v1 = "foo bar"
        } else {
            switch "foo bar" {
            case "\n\t":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "":
                v48 := GenS7{F0: -9223372036854775808}
                _ = v48
            default:
                v1 = ((v1 + ("" + (("" + v1) + ((v1 + v1) + v1)))) + "hello")
            }
            switch 42 {
            case 2147483647:
                v45 += v2
            case 9223372036854775807:
                for fi49 := 0; fi49 < 4; fi49++ {
                    v2 = float64(1.5)
                    _ = fi49
                }
            default:
                var v50 uint = 4294967295
                _ = v50
            }
            if (v2 == 3.14159) {
                var v51 uint32 = 2147483647
                _ = v51
            } else {
                v5 = append(v5, "")
                for fi52 := 0; fi52 < 4; fi52++ {
                    fmt.Println(v1)
                    fmt.Println(v1)
                    v53 := GenS3{F0: 1}
                    _ = fi52
                    _ = v53
                }
            }
        }
        _ = v45
    }
    if false {
        v7 -= v7
        v2 += v2
    }
    fmt.Println(v6)
    fmt.Println(v1, v3)
    fmt.Println(v6, v1)
    if true {
        if (!true) {
            var v54 float64 = 1e10
            v55 := GenS8{F0: 1000, F1: 2, F2: (^(v7))}
            switch "foo bar" {
            case "\n\t":
                v6 = (v6 | v6)
            case "foo bar":
                if (!false) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v56 int = 2147483647
                    v1 = ((v1 + (v1 + (v1 + (v1 + v1)))) + v1)
                    _ = v56
                } else {
                    if (42 <= 32767) {
                        switch -2147483648 {
                        case -343:
                            var v57 float64 = -2.25
                            _ = v57
                        case 42:
                            var v58 int8 = -128
                            _ = v58
                        case 2147483647:
                            fmt.Println(v54, v7, v6)
                        default:
                            _ = len(v5)
                        }
                        var v59 int = -1
                        switch 0 {
                        case -956:
                            if false {
                                if (660 < 42) {
                                    var v60 uint64 = 4294967295
                                    _ = v60
                                }
                                v61 := GenS9{F0: v3}
                                v2 = -2.25
                                _ = v61
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v59 += 2147483647
                            }
                        case -9223372036854775808:
                            v3 = -2147483648
                        case 1:
                            var v62 int = v59
                            _ = v62
                        }
                        _ = v59
                    }
                    v3 = 42
                }
            case "":
                fmt.Println(v54)
            }
            _ = v54
            _ = v55
        }
    } else {
        v2 = v2
        var v63 int64 = 42
        if len(v5) > 0 {
            v5[0] = v1
        }
        _ = v63
    }
    v64 := GenS4{F0: 9223372036854775807, F1: 32}
    switch "foo bar" {
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "hello":
        var v65 uint16 = 1
        _ = v65
    case "foo bar":
        if false {
            v66 := GenS10{F0: 65535, F1: (0 <= -1), F2: -275}
            _ = v66
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v7
}
