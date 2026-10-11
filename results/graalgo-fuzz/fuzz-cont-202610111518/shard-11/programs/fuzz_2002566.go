package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int32 }
type GenS2 struct { F0 int; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 int8; F1 bool; F2 string }
func genFn0(p0 uint) uint64 {
    if ((!(128 == 255)) && false) {
        var v7 uint8 = 42
        _ = v7
    }
    p0 += p0
    v8 := GenS1{F0: false, F1: 181, F2: 2147483647}
    v9 := GenS1{F0: (-128 >= -124), F1: 211, F2: 42}
    v10 := []int{128, -940, -9223372036854775808, 32768}
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint = 255
    var v2 int8 = 100
    var v3 uint = 22
    var v4 uint32 = 2
    var v5 string = (("hello" + "a") + ("a" + ("\n\t" + ("a" + ("a" + "foo bar")))))
    v6 := make(map[uint32]string)
    delete(v6, uint32(v1))
    if true {
        var v11 int = 4294967295
        var v12 int32 = -1
        _ = v11
        _ = v12
    } else {
        v13 := GenS1{F0: (0.1 <= 1e10), F1: 0, F2: -492}
        if (391 >= 42) {
            v14 := GenS1{F0: true, F1: 127, F2: 606}
            _ = v14
        }
        if (!(32767 != -248)) {
            v1 -= (v3 + (v1 + 335))
            if (false && (!true)) {
                switch (("" + "foo bar") + v5) {
                case "\n\t":
                    v15 := GenS2{F0: -1, F1: 38, F2: 42}
                    _ = v15
                default:
                    v16 := v6[2147483647]
                    _ = v16
                }
            } else {
                switch v5 {
                case "foo bar":
                    var v17 float32 = ((float32((0.1 / 0.1)) * 0.1) + (((1e10 / 1.0) - (1e10 - 0.1)) + 1e10))
                    _ = v17
                case "":
                    v18 := GenS2{F0: 0, F1: 128, F2: 42}
                    _ = v18
                case "a":
                    _ = len(v6)
                default:
                    v6[v4] = v5
                }
                for fi19 := 0; fi19 < 12; fi19++ {
                    v20 := []int{(fi19 << (((v3 % (((v1 ^ (v3 ^ (^(v3))))) | 1))) & 63)), fi19, fi19}
                    if ((v4 * 4294967295) < 65535) {
                        fmt.Println(v1)
                    } else {
                        v21 := GenS2{F0: fi19, F1: 128, F2: 446}
                        _ = v21
                    }
                    if false {
                        var v22 bool = true
                        if true {
                            v23 := GenS2{F0: (fi19 & fi19), F1: uint8(v3), F2: 1}
                            _ = v23
                        } else {
                            var v24 int = (fi19 >> ((v3) & 63))
                            fmt.Println(v5)
                            fmt.Println(v4, v2, v24)
                            _ = v24
                        }
                        _ = v22
                    }
                    _ = fi19
                    _ = v20
                }
                v25 := []int{0}
                _ = v25
            }
            var v26 int64 = 192
            _ = v26
        }
        _ = v13
    }
    var v27 float32 = (1.5 - 1e10)
    for fi28 := 0; fi28 < 13; fi28++ {
        if (!true) {
            switch (-(fi28)) {
            case 4294967295:
                if (!(!(!false))) {
                    v27 += (v27 * -2.25)
                    _ = len(v6)
                } else {
                    var v29 string = "a"
                    _ = v29
                }
            case 32768:
                var v30 uint8 = 236
                _ = v30
            case -2147483648:
                var v31 int8 = (^(v2))
                _ = v31
            default:
                switch v5 {
                case "hello":
                    for fi32 := 0; fi32 < 2; fi32++ {
                        var v33 int = fi32
                        _ = fi32
                        _ = v33
                    }
                default:
                    if (v4 <= 2147483647) {
                        fmt.Println(v4)
                    }
                }
            }
            v34 := []int{(-(fi28)), fi28, (fi28 | (fi28 | (fi28 ^ (fi28 | fi28))))}
            _ = v34
        }
        _ = fi28
    }
    switch 4294967295 {
    case -9223372036854775808:
        v1 = v1
    case 32767:
        var v35 int16 = 1000
        _ = v35
    case 42:
        v36 := GenS3{F0: (v4 * v4)}
        _ = v36
    default:
        if (!false) {
            for fi37 := 0; fi37 < 11; fi37++ {
                var v38 float64 = -2.25
                var v39 int32 = -2147483648
                v40 := GenS4{F0: 292}
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
            }
        } else {
            for fi41 := 0; fi41 < 6; fi41++ {
                fmt.Println(v5, v5)
                if true {
                    if (!(false || (true || ((v27 != 1e10) || false)))) {
                        v3 = (v3 / ((v1) | 1))
                        v4 += 2
                        var v42 int8 = (v2 + v2)
                        _ = v42
                    }
                } else {
                    v43 := []int{880, 363}
                    if (false || (false && (false && false))) {
                        if false {
                            for fi44 := 0; fi44 < 17; fi44++ {
                                fmt.Println(v5, v27)
                                var v45 float32 = v27
                                _ = fi44
                                _ = v45
                            }
                            var v46 int16 = -420
                            v47 := GenS4{F0: 284}
                            _ = v46
                            _ = v47
                        } else {
                            var v48 uint64 = 18446744073709551615
                            switch ((v5 + v5) + "foo bar") {
                            case "a":
                                v49 := GenS2{F0: 32768, F1: 0, F2: 0}
                                _ = v49
                            default:
                                var v50 int = (fi41 - 42)
                                _ = v50
                            }
                            v48 -= (v48 * v48)
                            _ = v48
                        }
                    } else {
                        if (false || (true && (true || (1e10 == 1e10)))) {
                            var v51 int8 = v2
                            for fi52 := 0; fi52 < 8; fi52++ {
                                switch (fi41 << ((v1) & 63)) {
                                case 9223372036854775807:
                                    var v53 bool = (false || true)
                                    var v54 int16 = 0
                                    _ = v53
                                    _ = v54
                                case 42:
                                    var v55 float64 = 1.5
                                    _ = v55
                                default:
                                    var v56 float64 = float64(v27)
                                    _ = v56
                                }
                                switch v5 {
                                case "hello":
                                    switch v5 {
                                    case "\n\t":
                                        var v57 bool = false
                                        _ = v57
                                    default:
                                        _ = len(v43)
                                    }
                                case "foo bar":
                                    if true {
                                        var v58 uint32 = v4
                                        var v59 float64 = (((((float64(v27) + ((1e10 * 0.0) / (float64(v27) * (0.1 - float64(((((1e10 / (1.5 + ((v27 * v27) / v27))) * (v27 - (v27 * (v27 / 1.0)))) - (((v27 - v27) - v27) - 1.5)) * (v27 + v27))))))) / 0.1) + (0.0 - float64(v27))) - 3.14159) / (3.14159 / (0.1 * (1.5 - (3.14159 / 1e10)))))
                                        v60 := GenS4{F0: 263}
                                        if true {
                                            var v61 uint8 = 42
                                            _ = v61
                                        }
                                        var v62 int8 = 1
                                        _ = v58
                                        _ = v59
                                        _ = v60
                                        _ = v62
                                    }
                                case "":
                                    v5 = (v5 + v5)
                                }
                                if (2147483647 >= 1) {
                                    switch v5 {
                                    case "":
                                        v63 := GenS5{F0: (!true), F1: -533}
                                        _ = v63
                                    default:
                                        var v64 uint = (v1 & (^(v3)))
                                        _ = v64
                                    }
                                    fmt.Println(v27)
                                    fmt.Println(fi52)
                                    if true {
                                        v4 -= 4294967295
                                    }
                                    if (!(false || false)) {
                                        v2 -= (-(v51))
                                        fmt.Println(fi41)
                                        switch 658 {
                                        case 549:
                                            delete(v6, 2)
                                        case 127:
                                            v65 := GenS3{F0: v4}
                                            _ = v65
                                        default:
                                            var v66 int64 = -213
                                            _ = v66
                                        }
                                    } else {
                                        var v67 string = v5
                                        var v68 int8 = 0
                                        _ = v67
                                        _ = v68
                                    }
                                    var v69 uint32 = (v4 ^ v4)
                                    _ = v69
                                }
                                _ = fi52
                            }
                            _ = v51
                        } else {
                            var v70 uint64 = 4294967295
                            fmt.Println(v1, v27, v2)
                            _ = v70
                        }
                        var v71 uint = v1
                        if true {
                            v72 := v6[925]
                            _ = v72
                        } else {
                            v73 := GenS4{F0: 0}
                            for fi74 := 0; fi74 < 15; fi74++ {
                                v75 := GenS6{F0: 42}
                                v27 += v27
                                _ = fi74
                                _ = v75
                            }
                            _ = v73
                        }
                        _ = v71
                    }
                    _ = v43
                }
                _ = fi41
            }
            v1 += v1
            switch 4294967295 {
            case 32768:
                fmt.Println(v3)
            case 457:
                if (false && false) {
                    v4 = v4
                } else {
                    if true {
                        var v76 bool = ("" != v5)
                        if v76 {
                            if (-632 < int(v2)) {
                                var v77 uint = (^(v3))
                                _ = v77
                            }
                        } else {
                            v6[(v4 << (((^(v1))) & 63))] = v5
                            var v78 int32 = 1
                            _ = v78
                        }
                        for fi79 := 0; fi79 < 6; fi79++ {
                            v4 += v4
                            _ = fi79
                        }
                        _ = v76
                    } else {
                        v80 := v6[v4]
                        _ = v80
                    }
                    switch v5 {
                    case "\n\t":
                        v4 = v4
                    case "a":
                        fmt.Println(v27, v5, v27)
                    case "":
                        if ((!true) || false) {
                            v2 = v2
                        }
                    }
                    var v81 int8 = v2
                    _ = v81
                }
            case 32767:
                v27 = v27
            }
        }
    }
    var v82 int8 = v2
    delete(v6, (^(v4)))
    v1 += (v1 & v3)
    var v83 int32 = 1
    fmt.Println(v83, v5, v2)
    v82 -= 70
    for fi84 := 0; fi84 < 4; fi84++ {
        for fi85 := 0; fi85 < 3; fi85++ {
            var v86 uint16 = 0
            _ = fi85
            _ = v86
        }
        v87 := GenS3{F0: (v4 + (v4 >> ((18446744073709551615) & 63)))}
        _ = fi84
        _ = v87
    }
    var v88 bool = (false && (!("\n\t" == "hello")))
    v89 := []int{1, 128, 9223372036854775807}
    var v90 float64 = (1.5 - float64(0.0))
    v91 := GenS7{F0: v1, F1: (v2 | (v2 - (v82 & 100)))}
    v92 := GenS8{F0: (v2 * v2), F1: (v3 >= (v1 + (v3 ^ v1))), F2: v5}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    fmt.Println("v88:", v88)
    fmt.Println("slicelen v89:", len(v89), v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v82
    _ = v83
    _ = v88
    _ = v89
    _ = v90
    _ = v91
    _ = v92
}
