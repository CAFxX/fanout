package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 uint; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string }

func main() {
    var v1 bool = true
    v2 := []uint32{2147483648, 902, 42}
    var v3 uint8 = 1
    var v4 int8 = -128
    v5 := []uint32{62}
    var v6 uint32 = 356
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int8 = 42
        switch fi7 {
        case -1:
            switch fi7 {
            case -364:
                var v9 int16 = 32767
                _ = v9
            case 0:
                if (18446744073709551615 < 0) {
                    if (0.0 == 1e10) {
                        switch (fi7 << ((65535) & 63)) {
                        case -1:
                            var v10 int16 = -929
                            _ = v10
                        case -625:
                            var v11 uint32 = v6
                            _ = v11
                        case 9223372036854775807:
                            if (true || (false && false)) {
                                for fi12 := 0; fi12 < 12; fi12++ {
                                    switch "hello" {
                                    case "":
                                        var v13 int16 = 1
                                        _ = v13
                                    case "\n\t":
                                        v6 = (v6 ^ 65535)
                                    default:
                                        v14 := GenS1{F0: ((2147483647 < 634) || (((-32768 != 0) && v1) || v1)), F1: 128}
                                        _ = v14
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = fi12
                                }
                            } else {
                                for fi15 := 0; fi15 < 2; fi15++ {
                                    v3 = (^(v3))
                                    _ = fi15
                                }
                                for fi16 := 0; fi16 < 12; fi16++ {
                                    v6 = v6
                                    var v17 int8 = v4
                                    _ = len(v5)
                                    _ = fi16
                                    _ = v17
                                }
                                v1 = (0.1 == (1.5 - 1.5))
                            }
                        default:
                            fmt.Println(v3, v1, v1)
                        }
                        var v18 uint64 = 42
                        v18 += v18
                        _ = v18
                    } else {
                        if (false && (!v1)) {
                            if ((v1 && v1) && (-903 > 328)) {
                                if ((^(v3)) == 2) {
                                    var v19 uint = 2
                                    _ = v19
                                } else {
                                    for fi20 := 0; fi20 < 12; fi20++ {
                                        switch "hello" {
                                        case "foo bar":
                                            if (0 < 42) {
                                                var v21 float64 = 3.14159
                                                _ = v21
                                            } else {
                                                _ = len(v2)
                                            }
                                        default:
                                            v22 := GenS2{F0: 2147483647, F1: 2}
                                            _ = v22
                                        }
                                        var v23 bool = v1
                                        if true {
                                            v24 := GenS3{F0: (v6 - (^(v6))), F1: v23}
                                            _ = v24
                                        }
                                        _ = fi20
                                        _ = v23
                                    }
                                    fmt.Println(v8, v6)
                                    var v25 int32 = 2147483647
                                    _ = v25
                                }
                                if false {
                                    fmt.Println(v4, v1)
                                    var v26 uint64 = 129
                                    _ = v26
                                }
                            } else {
                                v27 := GenS4{F0: 255, F1: v3, F2: 1}
                                v28 := GenS1{F0: true, F1: (v3 % ((42) | 1))}
                                v29 := GenS5{F0: 42, F1: 65535, F2: 233}
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            v8 += (v4 >> ((42) & 63))
                            v4 = 50
                        } else {
                            fmt.Println(v4)
                            v3 = 128
                        }
                        var v30 int16 = 0
                        for fi31 := 0; fi31 < 1; fi31++ {
                            fmt.Println(v3, fi31)
                            _ = fi31
                        }
                        _ = v30
                    }
                } else {
                    var v32 uint64 = 479
                    _ = v32
                }
            case -200:
                var v33 float32 = 0.1
                _ = v33
            default:
                if v1 {
                    v8 = v8
                }
            }
        case 0:
            fmt.Println(v3, v3)
        case 614:
            var v34 string = ("" + "foo bar")
            _ = v34
        }
        var v35 uint = 255
        _ = fi7
        _ = v35
        _ = v8
    }
    var v36 uint16 = 65535
    v36 -= (^(v36))
    for fi37 := 0; fi37 < 18; fi37++ {
        v1 = v1
        _ = fi37
    }
    var v38 int32 = 1
    switch -60 {
    case 127:
        if v1 {
            for fi39 := 0; fi39 < 13; fi39++ {
                v36 = v36
                _ = fi39
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v40 := GenS6{F0: 32767}
            _ = v40
        } else {
            var v41 float64 = (float64(0.0) / -2.25)
            if true {
                var v42 int64 = 232
                v36 += (v36 * 2)
                _ = v42
            } else {
                v36 -= (v36 << ((1) & 63))
            }
            switch -9223372036854775808 {
            case 128:
                if v1 {
                    var v43 uint8 = v3
                    _ = v43
                } else {
                    v5 = append(v5, 65535)
                    for fi44 := 0; fi44 < 10; fi44++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v45 := GenS5{F0: v3, F1: 1, F2: 183}
                        _ = fi44
                        _ = v45
                    }
                    v2 = append(v2, v6)
                }
            case 32768:
                var v46 int = -895
                _ = v46
            case 4294967295:
                v47 := GenS7{F0: 40}
                _ = v47
            }
            _ = v41
        }
    case 32768:
        v1 = v1
    case -237:
        for fi48 := 0; fi48 < 18; fi48++ {
            _ = len(v2)
            v4 = (v4 | 127)
            if (32767 <= 81) {
                for fi49 := 0; fi49 < 14; fi49++ {
                    var v50 int16 = 1
                    if (!((v4 <= v4) && v1)) {
                        v51 := GenS4{F0: 18446744073709551615, F1: v3, F2: 42}
                        if true {
                            v52 := GenS8{F0: (("a" + "") + "\n\t")}
                            _ = v52
                        } else {
                            var v53 uint64 = uint64(v3)
                            fmt.Println(v1, v6)
                            _ = v53
                        }
                        v1 = (!(v1 || v1))
                        _ = v51
                    } else {
                        _ = len(v5)
                        var v54 string = "hello"
                        var v55 int = fi48
                        _ = v54
                        _ = v55
                    }
                    var v56 int32 = v38
                    _ = fi49
                    _ = v50
                    _ = v56
                }
                v3 = v3
                var v57 float32 = 3.14159
                _ = v57
            }
            _ = fi48
        }
    default:
        var v58 float32 = ((0.1 / ((1.5 - (((3.14159 * -2.25) * ((1.5 - float32((-2.25 * -2.25))) * ((1.5 / (1e10 - -2.25)) - 1.5))) - -2.25)) - -2.25)) - float32((1e10 - 0.1)))
        _ = v58
    }
    for fi59 := 0; fi59 < 1; fi59++ {
        v60 := GenS3{F0: v6, F1: (0.0 != -2.25)}
        fmt.Println(v38, v4, v1)
        v61 := GenS1{F0: ((3.14159 * 0.1) >= ((0.1 * (3.14159 * (float32(1e10) / (3.14159 / 3.14159)))) * 1.5)), F1: 1}
        _ = fi59
        _ = v60
        _ = v61
    }
    v62 := GenS1{F0: false, F1: (v3 + 2)}
    v3 = 210
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v62
}
