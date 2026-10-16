package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 bool = (43 <= 255)
    v2 := make(map[uint32]string)
    var v3 uint64 = 0
    v4 := make(map[int]bool)
    v5 := []int{901, 360}
    switch "\n\t" {
    case "":
        var v6 bool = v1
        _ = v6
    default:
        if v1 {
            v1 = ((v3 != (v3 | (^(v3)))) && v1)
        } else {
            for fi7 := 0; fi7 < 16; fi7++ {
                for fi8 := 0; fi8 < 4; fi8++ {
                    _ = len(v2)
                    var v9 uint32 = 65535
                    _ = fi8
                    _ = v9
                }
                if (!(730 < 65535)) {
                    _ = len(v2)
                    for fi10 := 0; fi10 < 17; fi10++ {
                        v3 += (v3 % ((v3) | 1))
                        var v11 float64 = 3.14159
                        _ = fi10
                        _ = v11
                    }
                    v12 := v2[2147483647]
                    _ = v12
                }
                _ = fi7
            }
        }
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        v1 = (!false)
        if len(v5) > 0 {
            v5[0] = fi13
        }
        switch "" {
        case "foo bar":
            var v14 int64 = 773
            _ = v14
        case "\n\t":
            fmt.Println(v1)
        case "a":
            v3 -= 0
        default:
            fmt.Println(v1)
        }
        _ = fi13
    }
    if (2 <= 65535) {
        v15 := GenS1{F0: v3, F1: 255, F2: 2}
        _ = v15
    } else {
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 int = fi16
            var v18 uint64 = v3
            _ = fi16
            _ = v17
            _ = v18
        }
        v1 = false
    }
    var v19 uint = 18446744073709551615
    if (v1 || (true && (uint64(v19) != (v3 ^ v3)))) {
        for fi20 := 0; fi20 < 8; fi20++ {
            if len(v5) > 0 {
                v5[0] = (-(fi20))
            }
            _ = fi20
        }
    } else {
        var v21 float32 = (1.5 - (3.14159 + 1.5))
        v5 = append(v5, 31)
        _ = v21
    }
    var v22 uint16 = 1
    switch ("\n\t" + "\n\t") {
    case "hello":
        if ((!(30 < 140)) && (42 > 1)) {
            v23 := GenS2{F0: 65535}
            switch (("" + "foo bar") + ("" + (("\n\t" + "") + "foo bar"))) {
            case "a":
                switch "hello" {
                case "foo bar":
                    v24 := GenS3{F0: (!(-2.25 >= 0.1)), F1: 18, F2: 2}
                    _ = v24
                case "":
                    v25 := GenS4{F0: (v22 | uint16(v3)), F1: -842}
                    _ = v25
                default:
                    _ = len(v5)
                }
            }
            _ = v23
        } else {
            if v1 {
                v26 := GenS5{F0: (!((-2.25 != ((0.0 + (1.5 * (0.0 / 1.0))) + 0.0)) || (true || v1))), F1: (v22 + (v22 - 32767)), F2: 1}
                _ = v26
            } else {
                for fi27 := 0; fi27 < 5; fi27++ {
                    if (!v1) {
                        v28 := GenS1{F0: 34, F1: 32767, F2: 0}
                        _ = v28
                    }
                    if (fi27 > (fi27 << (((^(v19))) & 63))) {
                        var v29 int8 = 127
                        _ = v29
                    } else {
                        v30 := GenS2{F0: v19}
                        _ = v30
                    }
                    v22 = (v22 / ((v22) | 1))
                    _ = fi27
                }
            }
            v31 := v2[1]
            _ = v31
        }
    case "":
        if v1 {
            for fi32 := 0; fi32 < 19; fi32++ {
                var v33 int64 = -9223372036854775808
                if ((^(v3)) >= v3) {
                    v33 += (v33 + (^(v33)))
                    for fi34 := 0; fi34 < 9; fi34++ {
                        var v35 int32 = 2147483647
                        _ = len(v2)
                        v36 := GenS1{F0: v3, F1: v22, F2: 678}
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                } else {
                    for fi37 := 0; fi37 < 18; fi37++ {
                        _ = len(v4)
                        var v38 int8 = 1
                        switch ("foo bar" + "foo bar") {
                        case "a":
                            if ((((-2.25 * 3.14159) * 3.14159) + 0.1) <= float64(-2.25)) {
                                if (v22 <= (v22 ^ (v22 % ((160) | 1)))) {
                                    if ((v22 * v22) > v22) {
                                        _ = len(v2)
                                        var v39 float32 = 1.5
                                        _ = v39
                                    }
                                    if (float64(1e10) != ((((float64((float32(3.14159) / 0.1)) + ((1e10 / (1e10 * 1.5)) / float64(1.5))) - float64(3.14159)) * (1e10 / 1.0)) * 0.0)) {
                                        _ = len(v2)
                                        if (!((uint8(v3) >= 128) || ((float64(0.1) * 0.0) == 1e10))) {
                                            fmt.Println(fi37, v1)
                                            if v1 {
                                                v22 -= v22
                                                v1 = (0.1 < 1e10)
                                                switch (fi37 / (((^(fi32))) | 1)) {
                                                case 688:
                                                    var v40 uint16 = 628
                                                    _ = v40
                                                default:
                                                    var v41 uint64 = v3
                                                    _ = v41
                                                }
                                            } else {
                                                fmt.Println(v3, fi37)
                                            }
                                        } else {
                                            fmt.Println(v3, v3)
                                            v3 -= (v3 | v3)
                                        }
                                        switch (("foo bar" + "a") + ("" + ("hello" + "hello"))) {
                                        case "\n\t":
                                            v38 -= (^(v38))
                                        default:
                                            v33 += (v33 >> ((v19) & 63))
                                        }
                                    }
                                }
                                var v42 bool = true
                                _ = v42
                            }
                        default:
                            var v43 int32 = -210
                            var v44 int32 = -282
                            v1 = (false || (106 > v38))
                            _ = v43
                            _ = v44
                        }
                        _ = fi37
                        _ = v38
                    }
                }
                _ = fi32
                _ = v33
            }
        } else {
            for fi45 := 0; fi45 < 7; fi45++ {
                v3 = v3
                fmt.Println(v3)
                _ = fi45
            }
            if len(v5) > 0 {
                v5[0] = 32768
            }
            for fi46 := 0; fi46 < 20; fi46++ {
                for fi47 := 0; fi47 < 12; fi47++ {
                    var v48 uint32 = 2147483647
                    var v49 uint64 = v3
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                for fi50 := 0; fi50 < 19; fi50++ {
                    v51 := GenS6{F0: 1, F1: v1}
                    v22 = (^(v22))
                    if ((true || v1) || true) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if v1 {
                            v52 := GenS7{F0: (^(v19)), F1: (fi46 % ((fi50) | 1))}
                            _ = v52
                        }
                    }
                    _ = fi50
                    _ = v51
                }
                v53 := GenS1{F0: v3, F1: (v22 & (v22 >> (((v19 / ((v19) | 1))) & 63))), F2: 4294967295}
                _ = fi46
                _ = v53
            }
        }
    }
    v22 -= (^(v22))
    for fi54 := 0; fi54 < 20; fi54++ {
        v19 = (v19 % (((v19 + v19)) | 1))
        _ = fi54
    }
    switch "foo bar" {
    case "foo bar":
        v22 = v22
    case "":
        var v55 uint32 = 556
        _ = v55
    default:
        v3 = v3
    }
    for fi56 := 0; fi56 < 12; fi56++ {
        v57 := GenS8{F0: -453}
        v58 := GenS7{F0: (v19 - (v19 & v19)), F1: fi56}
        var v59 int32 = 633
        _ = fi56
        _ = v57
        _ = v58
        _ = v59
    }
    for fi60 := 0; fi60 < 12; fi60++ {
        switch "\n\t" {
        case "\n\t":
            v61 := GenS9{F0: (v3 - (^(v3))), F1: 42}
            _ = v61
        }
        _ = fi60
    }
    for fi62 := 0; fi62 < 20; fi62++ {
        v19 = v19
        var v63 int = fi62
        _ = fi62
        _ = v63
    }
    for fi64 := 0; fi64 < 17; fi64++ {
        if v1 {
            for fi65 := 0; fi65 < 13; fi65++ {
                if v1 {
                    fmt.Println(v22, v19)
                }
                _ = fi65
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v66 := GenS3{F0: ((((v22 | 552) < (v22 * (v22 - (v22 & 737)))) && false) || false), F1: 42, F2: 105}
            _ = v66
        }
        var v67 uint32 = 65535
        fmt.Println(v1)
        _ = fi64
        _ = v67
    }
    var v68 int16 = -222
    var v69 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v69
}
