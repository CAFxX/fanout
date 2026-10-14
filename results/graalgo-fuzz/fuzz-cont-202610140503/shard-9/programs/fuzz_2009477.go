package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int64; F1 uint8; F2 int64 }
type GenS4 struct { F0 bool; F1 int; F2 uint16 }
type GenS5 struct { F0 uint; F1 uint; F2 bool }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64; F1 uint; F2 bool }
type GenS9 struct { F0 uint16; F1 string; F2 int16 }

func main() {
    v1 := []int{0, 9223372036854775807, 1}
    var v2 uint8 = 127
    var v3 uint32 = 0
    var v4 int64 = 797
    var v5 string = "foo bar"
    var v6 int = 54
    var v7 uint32 = (v3 + (v3 ^ (v3 % ((2147483648) | 1))))
    fmt.Println(v3, v4, v3)
    var v8 uint8 = (v2 - (v2 / (((v2 << ((18446744073709551615) & 63))) | 1)))
    v9 := GenS1{F0: v5, F1: -598}
    if (v5 != (v5 + v5)) {
        var v10 string = ("" + ("foo bar" + v5))
        if len(v1) > 0 {
            v1[0] = (v6 | v6)
        }
        v3 = v7
        _ = v10
    } else {
        switch v5 {
        case "hello":
            if (3.14159 != 3.14159) {
                switch 128 {
                case -1:
                    var v11 int8 = -128
                    _ = v11
                case 9223372036854775807:
                    if true {
                        for fi12 := 0; fi12 < 11; fi12++ {
                            fmt.Println(v3, v3)
                            _ = fi12
                        }
                        for fi13 := 0; fi13 < 3; fi13++ {
                            if true {
                                var v14 float32 = float32(0.0)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v15 uint64 = 4294967295
                                _ = v14
                                _ = v15
                            }
                            for fi16 := 0; fi16 < 18; fi16++ {
                                fmt.Println(fi16, fi16)
                                v7 += 42
                                fmt.Println(v2, v7)
                                v17 := GenS1{F0: v5, F1: 0}
                                switch (v5 + ("" + ("hello" + "\n\t"))) {
                                case "hello":
                                    if true {
                                        _ = len(v1)
                                        v1 = append(v1, fi13)
                                    }
                                    v8 -= (v8 + (v2 / (((^(v2))) | 1)))
                                    if (-32768 < 619) {
                                        var v18 uint32 = v7
                                        fmt.Println(v8, fi13, v7)
                                        _ = v18
                                    }
                                    var v19 int16 = 104
                                    _ = v19
                                default:
                                    switch 127 {
                                    case 692:
                                        var v20 int = (fi13 >> ((2) & 63))
                                        _ = v20
                                    case 1:
                                        if true {
                                            if (false && false) {
                                                switch (fi13 * (v6 * fi16)) {
                                                case 0:
                                                    var v21 float32 = 3.14159
                                                    _ = v21
                                                case 32767:
                                                    v4 = v4
                                                }
                                                var v22 int16 = -818
                                                _ = v22
                                            } else {
                                                if false {
                                                    switch -1 {
                                                    case -258:
                                                        if (!(1 <= 18446744073709551615)) {
                                                            _ = len(v1)
                                                            if (375 != -2147483648) {
                                                                fmt.Println(v7, v3)
                                                                v3 -= (v7 % (((v7 << ((4294967295) & 63))) | 1))
                                                            }
                                                        }
                                                    default:
                                                        switch (fi16 % ((v6) | 1)) {
                                                        case 4294967295:
                                                            v23 := GenS2{F0: 0, F1: (18446744073709551615 != 586)}
                                                            _ = v23
                                                        case -979:
                                                            var v24 int = fi16
                                                            _ = v24
                                                        }
                                                    }
                                                } else {
                                                    v2 -= v2
                                                    var v25 int64 = -865
                                                    v6 += -1
                                                    _ = v25
                                                }
                                                v26 := GenS3{F0: (v4 >> ((42) & 63)), F1: v8, F2: 1}
                                                _ = v26
                                            }
                                            var v27 uint32 = v7
                                            v8 -= 211
                                            _ = v27
                                        }
                                    case 32767:
                                        if (1 <= v2) {
                                            fmt.Println(fi13, fi13, v3)
                                            var v28 float64 = (-2.25 + 1e10)
                                            _ = v28
                                        }
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = fi16
                                _ = v17
                            }
                            if true {
                                for fi29 := 0; fi29 < 20; fi29++ {
                                    switch int(v4) {
                                    case 128:
                                        if len(v1) > 0 {
                                            v1[0] = (fi13 << ((523) & 63))
                                        }
                                    case 79:
                                        var v30 bool = true
                                        _ = v30
                                    default:
                                        v2 -= v8
                                    }
                                    var v31 string = v5
                                    v3 += v3
                                    _ = fi29
                                    _ = v31
                                }
                                if false {
                                    if ((v7 >> ((65535) & 63)) < v7) {
                                        for fi32 := 0; fi32 < 17; fi32++ {
                                            v4 = (v4 * (v4 | v4))
                                            _ = fi32
                                        }
                                        v33 := GenS4{F0: false, F1: 32768, F2: 488}
                                        _ = v33
                                    }
                                    var v34 int16 = -604
                                    v35 := GenS5{F0: 1, F1: 255, F2: (!(false || true))}
                                    _ = v34
                                    _ = v35
                                } else {
                                    v36 := GenS6{F0: v4, F1: -1}
                                    fmt.Println(v3, v7)
                                    _ = v36
                                }
                                var v37 uint32 = v7
                                _ = v37
                            } else {
                                var v38 int64 = (v4 % ((v4) | 1))
                                fmt.Println(v7)
                                switch v6 {
                                case 127:
                                    var v39 int64 = v4
                                    _ = v39
                                }
                                _ = v38
                            }
                            _ = fi13
                        }
                        var v40 float64 = float64(float32(0.1))
                        _ = v40
                    }
                case -909:
                    for fi41 := 0; fi41 < 8; fi41++ {
                        var v42 uint32 = 0
                        var v43 int16 = 430
                        v5 = (("a" + v5) + (v5 + ((v5 + v5) + v5)))
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                default:
                    if len(v1) > 0 {
                        v1[0] = (v6 & 127)
                    }
                }
                if false {
                    v44 := GenS7{F0: -1}
                    for fi45 := 0; fi45 < 20; fi45++ {
                        v5 = (v5 + v5)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if (65535 == 60) {
                            var v46 uint32 = (v3 * 2147483647)
                            _ = v46
                        }
                        _ = fi45
                    }
                    v4 = (-(v4))
                    _ = v44
                }
                var v47 bool = ((v4 ^ (v4 * v4)) == (v4 >> ((18446744073709551615) & 63)))
                _ = v47
            } else {
                if true {
                    _ = len(v1)
                }
            }
        case "\n\t":
            if ((1e10 != 0.1) && true) {
                _ = len(v1)
                for fi48 := 0; fi48 < 13; fi48++ {
                    for fi49 := 0; fi49 < 18; fi49++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v5 = (("\n\t" + "\n\t") + "a")
                        _ = fi49
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi50 := 0; fi50 < 12; fi50++ {
                        var v51 uint8 = (v2 >> ((18446744073709551615) & 63))
                        _ = fi50
                        _ = v51
                    }
                    _ = fi48
                }
            } else {
                for fi52 := 0; fi52 < 15; fi52++ {
                    if (!(true || false)) {
                        v2 -= (v8 % ((255) | 1))
                        v53 := GenS1{F0: "foo bar", F1: 1}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v53
                    } else {
                        if len(v1) > 0 {
                            v1[0] = (-(v6))
                        }
                    }
                    v3 = v3
                    v54 := GenS4{F0: (-286 > 2147483647), F1: v6, F2: uint16(v8)}
                    _ = fi52
                    _ = v54
                }
                if (1 >= 2) {
                    for fi55 := 0; fi55 < 4; fi55++ {
                        v8 -= v2
                        v7 = 42
                        var v56 int32 = 828
                        _ = fi55
                        _ = v56
                    }
                } else {
                    v57 := GenS5{F0: 1, F1: 365, F2: (1 <= 1)}
                    if ((830 == 907) || (!true)) {
                        var v58 uint8 = (v2 | v8)
                        v6 -= 4294967295
                        _ = v58
                    } else {
                        if (true && true) {
                            var v59 uint64 = 18446744073709551615
                            _ = v59
                        } else {
                            v6 = v6
                            var v60 int64 = 0
                            _ = v60
                        }
                        fmt.Println(v4, v7)
                    }
                    switch -1 {
                    case 1:
                        v7 = 65535
                    case -947:
                        v7 = (v7 | (v7 - (v7 / ((v7) | 1))))
                    case 4294967295:
                        var v61 int32 = -942
                        _ = v61
                    }
                    _ = v57
                }
            }
        }
        var v62 int64 = -833
        _ = v62
    }
    v6 += (v6 ^ v6)
    var v63 uint16 = 2
    if len(v1) > 0 {
        v1[0] = v6
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v64 uint8 = (^(v2))
    if (0.1 > ((3.14159 + 3.14159) + 0.0)) {
        v65 := GenS8{F0: int64(v6), F1: 518, F2: (!false)}
        if (!false) {
            if true {
                v8 = v64
            }
            var v66 string = ((v5 + ("foo bar" + (v5 + "\n\t"))) + ("hello" + ("foo bar" + "a")))
            switch v66 {
            case "a":
                v67 := GenS9{F0: v63, F1: v5, F2: -273}
                _ = v67
            case "hello":
                v8 = (v64 << ((uint(v2)) & 63))
            default:
                fmt.Println(v3)
            }
            _ = v66
        }
        _ = v65
    }
    if false {
        fmt.Println(v2, v64, v3)
        for fi68 := 0; fi68 < 8; fi68++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi68
        }
    }
    var v69 int32 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v69:", v69)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v69
    _ = v7
    _ = v8
    _ = v9
}
