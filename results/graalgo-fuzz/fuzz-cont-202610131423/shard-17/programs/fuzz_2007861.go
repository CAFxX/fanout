package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int32; F2 uint32 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 uint; F2 uint8 }
type GenS8 struct { F0 int16; F1 int8; F2 int16 }
type GenS9 struct { F0 int8; F1 int16; F2 uint8 }
type GenS10 struct { F0 int16; F1 int }
func genFn0() uint64 {
    var v6 uint16 = 2
    v7 := []int{}
    var v8 int32 = 2147483647
    var v9 int32 = (v8 | (v8 | 2147483647))
    for fi10 := 0; fi10 < 4; fi10++ {
        v9 += (v9 * (v8 - (v9 * -2147483648)))
        if (42 >= -128) {
            if false {
                var v11 int32 = (v8 ^ v8)
                _ = v11
            }
            var v12 int64 = 0
            v8 -= v9
            _ = v12
        } else {
            var v13 uint8 = 88
            v14 := GenS1{F0: 127, F1: 1000, F2: 41}
            if false {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                switch (fi10 / ((298) | 1)) {
                case -203:
                    for fi15 := 0; fi15 < 20; fi15++ {
                        v6 = v6
                        if true {
                            v16 := GenS2{F0: 42}
                            var v17 int64 = 9223372036854775807
                            var v18 int16 = 0
                            fmt.Println(fi15, v6)
                            _ = v16
                            _ = v17
                            _ = v18
                        } else {
                            if true {
                                if len(v7) > 0 {
                                    v7[0] = fi15
                                }
                                var v19 uint8 = uint8(v6)
                                _ = v19
                            }
                            if true {
                                v13 += 42
                                v8 = v9
                            }
                        }
                        switch "hello" {
                        case "a":
                            switch fi15 {
                            case -2147483648:
                                var v20 bool = (42 >= 1)
                                _ = v20
                            case 767:
                                v21 := GenS1{F0: 0, F1: 42, F2: 2}
                                _ = v21
                            case -876:
                                if (!false) {
                                    if true {
                                        if false {
                                            _ = len(v7)
                                        }
                                        v22 := GenS3{F0: v13}
                                        _ = v22
                                    }
                                    if (3.14159 > float64(((float32(3.14159) - 1.5) / 3.14159))) {
                                        if len(v7) > 0 {
                                            v7[0] = (-(fi15))
                                        }
                                        v9 -= v8
                                    } else {
                                        switch fi15 {
                                        case -793:
                                            if (18446744073709551615 > 42) {
                                                var v23 uint32 = 507
                                                var v24 int8 = 100
                                                _ = v23
                                                _ = v24
                                            } else {
                                                v13 = v13
                                                var v25 uint32 = 0
                                                v8 = (-(v9))
                                                _ = v25
                                            }
                                        default:
                                            var v26 float32 = (1.5 + ((0.0 / float32(float64(1.5))) / 0.1))
                                            _ = v26
                                        }
                                        if ((!true) || ((true || ((42 < 9223372036854775807) && false)) && ("hello" > ("a" + "")))) {
                                            var v27 bool = true
                                            _ = v27
                                        } else {
                                            v7 = append(v7, (-(fi10)))
                                            fmt.Println(fi15, v6)
                                        }
                                    }
                                } else {
                                    v28 := GenS1{F0: 127, F1: -32768, F2: 42}
                                    _ = v28
                                }
                            }
                            v29 := GenS2{F0: 177}
                            _ = v29
                        case "hello":
                            var v30 int8 = -1
                            _ = v30
                        case "":
                            v31 := GenS4{F0: (v9 ^ v9), F1: 2147483647, F2: 2147483647}
                            _ = v31
                        }
                        _ = fi15
                    }
                case -176:
                    fmt.Println(v6, v9)
                }
                for fi32 := 0; fi32 < 18; fi32++ {
                    v9 = (v8 & 2147483647)
                    var v33 int = fi32
                    var v34 uint32 = 619
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            } else {
                switch "hello" {
                case "foo bar":
                    var v35 string = ""
                    _ = v35
                case "\n\t":
                    v36 := GenS2{F0: 1}
                    _ = v36
                default:
                    v9 = (^(v8))
                }
                v13 = (v13 / ((v13) | 1))
            }
            _ = v13
            _ = v14
        }
        for fi37 := 0; fi37 < 6; fi37++ {
            v38 := GenS5{F0: 65535, F1: (((float64(float32(((float64(-2.25) - ((float64((0.0 + ((1.5 + 0.1) * float32(0.0)))) * 0.1) / 3.14159)) - 0.0))) * (0.1 + ((0.0 + 3.14159) - -2.25))) * 1e10) != 3.14159)}
            _ = fi37
            _ = v38
        }
        _ = fi10
    }
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 892
}

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]uint32)
    var v3 int16 = (v1 ^ (^(v1)))
    var v4 string = ("foo bar" + "hello")
    v5 := []uint32{65535, 0, 42}
    var v39 float32 = 0.0
    if false {
        if true {
            for fi40 := 0; fi40 < 7; fi40++ {
                switch 42 {
                case 166:
                    switch v4 {
                    case "foo bar":
                        var v41 float64 = 3.14159
                        _ = v41
                    case "\n\t":
                        var v42 int16 = (-(v3))
                        _ = v42
                    case "a":
                        v4 = (v4 + ("hello" + (((((("hello" + v4) + (v4 + "")) + "hello") + v4) + v4) + v4)))
                    }
                default:
                    for fi43 := 0; fi43 < 13; fi43++ {
                        v4 = "\n\t"
                        _ = fi43
                    }
                }
                for fi44 := 0; fi44 < 8; fi44++ {
                    fmt.Println(v3, v1)
                    if true {
                        if false {
                            v45 := GenS6{F0: -1}
                            v46 := GenS7{F0: -697, F1: 18446744073709551615, F2: 57}
                            _ = v45
                            _ = v46
                        } else {
                            var v47 uint64 = 560
                            _ = len(v2)
                            _ = v47
                        }
                        var v48 int = (fi40 % ((fi40) | 1))
                        var v49 uint32 = 1
                        _ = v48
                        _ = v49
                    }
                    var v50 uint64 = 866
                    _ = fi44
                    _ = v50
                }
                _ = fi40
            }
            v4 = v4
            fmt.Println(v39, v39)
        }
        var v51 uint = 48
        fmt.Println(v51)
        _ = v51
    } else {
        if len(v5) > 0 {
            v5[0] = 42
        }
        v39 = 3.14159
        fmt.Println(v4, v4, v3)
    }
    var v52 bool = true
    fmt.Println(v4)
    var v53 float32 = 3.14159
    for fi54 := 0; fi54 < 12; fi54++ {
        fmt.Println(v53, v3)
        fmt.Println(v53, v52)
        if (false && ((true && false) && (int32(v1) >= 0))) {
            v52 = ((-2.25 / v39) != v39)
            var v55 string = (v4 + v4)
            for fi56 := 0; fi56 < 11; fi56++ {
                var v57 int32 = int32(v3)
                if (v52 && v52) {
                    v58 := GenS6{F0: v57}
                    _ = v58
                } else {
                    v5 = append(v5, 2147483647)
                }
                fmt.Println(v39)
                if ((v39 != (v39 + (v53 / v53))) && true) {
                    switch fi56 {
                    case 1:
                        v3 = 0
                    case 127:
                        if v52 {
                            if false {
                                switch (fi54 & (fi56 << ((961) & 63))) {
                                case 377:
                                    v1 = 91
                                case 4294967295:
                                    v59 := GenS8{F0: v1, F1: -53, F2: (v1 % ((v3) | 1))}
                                    _ = v59
                                }
                                v5 = append(v5, 65535)
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                if (!(127 <= 42)) {
                                    fmt.Println(v52, v1, v52)
                                    var v60 int32 = -2147483648
                                    if v52 {
                                        if v52 {
                                            v39 -= (v53 - -2.25)
                                        } else {
                                            v3 = int16(v57)
                                            fmt.Println(v3, v3, v52)
                                            v61 := v2[4294967295]
                                            _ = v61
                                        }
                                        var v62 int32 = (v60 + v57)
                                        _ = v62
                                    }
                                    _ = v60
                                } else {
                                    var v63 uint = 2
                                    v64 := GenS9{F0: -128, F1: v3, F2: 248}
                                    _ = v63
                                    _ = v64
                                }
                                fmt.Println(v55, v3)
                                var v65 int = (fi54 * fi56)
                                _ = v65
                            } else {
                                v4 = (("foo bar" + "") + "foo bar")
                                var v66 float32 = float32(float64(((3.14159 / v39) * v39)))
                                v5 = append(v5, 1)
                                _ = v66
                            }
                            var v67 uint = 2
                            _ = v67
                        }
                    case -557:
                        v4 = "\n\t"
                    }
                } else {
                    var v68 int8 = 42
                    if len(v5) > 0 {
                        v5[0] = 1
                    }
                    v69 := GenS9{F0: v68, F1: v3, F2: 43}
                    _ = v68
                    _ = v69
                }
                _ = fi56
                _ = v57
            }
            _ = v55
        } else {
            for fi70 := 0; fi70 < 1; fi70++ {
                var v71 uint16 = 409
                if ("foo bar" <= "\n\t") {
                    v72 := GenS10{F0: 327, F1: (fi54 >> ((uint(v71)) & 63))}
                    _ = v72
                } else {
                    var v73 int = -9223372036854775808
                    delete(v2, uint32(v71))
                    var v74 float64 = 1e10
                    _ = v73
                    _ = v74
                }
                _ = fi70
                _ = v71
            }
        }
        _ = fi54
    }
    v3 = (v1 - (v1 / (((v3 ^ (v1 << ((18446744073709551615) & 63)))) | 1)))
    var v75 uint8 = 128
    var v76 uint32 = 561
    var v77 uint = 387
    for fi78 := 0; fi78 < 3; fi78++ {
        fmt.Println(v39, v1)
        v79 := GenS4{F0: 2147483647, F1: -205, F2: (v76 | 977)}
        if len(v5) > 0 {
            v5[0] = 688
        }
        _ = fi78
        _ = v79
    }
    if false {
        var v80 int32 = -1
        v81 := GenS7{F0: -9223372036854775808, F1: uint(v76), F2: v75}
        v80 -= 1
        _ = v80
        _ = v81
    } else {
        var v82 float64 = -2.25
        _ = v82
    }
    var v83 int16 = (v3 | (v1 << (((^(v77))) & 63)))
    switch -1 {
    case -239:
        var v84 int64 = -804
        _ = v84
    case -9223372036854775808:
        for fi85 := 0; fi85 < 17; fi85++ {
            v5 = append(v5, v76)
            _ = fi85
        }
    case 32768:
        v52 = v52
    }
    for fi86 := 0; fi86 < 12; fi86++ {
        fmt.Println(fi86, v76)
        v53 -= v53
        _ = fi86
    }
    switch "a" {
    case "":
        switch v4 {
        case "":
            if len(v5) > 0 {
                v5[0] = v76
            }
        }
    case "a":
        if ((v77 - (v77 / ((340) | 1))) > v77) {
            switch v4 {
            case "\n\t":
                v52 = v52
            case "a":
                var v87 int32 = -839
                _ = v87
            case "foo bar":
                var v88 uint16 = 0
                _ = v88
            default:
                switch int(v76) {
                case 9223372036854775807:
                    fmt.Println(v75)
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v75
    _ = v76
    _ = v77
    _ = v83
}
