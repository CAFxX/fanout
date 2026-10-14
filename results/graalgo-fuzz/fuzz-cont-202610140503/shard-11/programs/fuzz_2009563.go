package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int8 }
type GenS2 struct { F0 string; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint; F1 int32; F2 string }
type GenS4 struct { F0 int64; F1 string; F2 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int32; F1 int8; F2 bool }
type GenS8 struct { F0 uint16; F1 string; F2 int32 }
type GenS9 struct { F0 uint32; F1 uint }
type GenS10 struct { F0 uint; F1 int32 }
func genFn0(p0 bool, p1 uint, p2 string) uint8 {
    switch "hello" {
    case "a":
        var v7 int16 = 42
        _ = v7
    case "":
        fmt.Println(p0)
    case "hello":
        if p0 {
            fmt.Println(p2, p0, p2)
            if p0 {
                if p0 {
                    var v8 int8 = -32
                    _ = v8
                }
                fmt.Println(p0)
            }
        }
    default:
        p1 += (p1 << (((p1 + (p1 + (p1 / ((601) | 1))))) & 63))
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := []int{(fi9 - fi9), fi9, (fi9 + -1), fi9}
        _ = fi9
        _ = v10
    }
    var v11 bool = false
    var v12 float64 = 1e10
    v12 -= 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 83
}
func genFn1(p0 uint32) uint8 {
    if false {
        var v13 uint8 = 42
        for fi14 := 0; fi14 < 10; fi14++ {
            for fi15 := 0; fi15 < 4; fi15++ {
                fmt.Println(v13, fi14, p0)
                v13 = (v13 << ((1) & 63))
                var v16 uint16 = uint16(p0)
                v16 = 32767
                var v17 uint32 = 1
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi14
        }
        if (127 < 32768) {
            fmt.Println(v13, v13)
            var v18 float64 = (0.1 * 3.14159)
            _ = v18
        } else {
            switch 316 {
            case 127:
                var v19 string = "hello"
                _ = v19
            case 557:
                var v20 uint = 0
                _ = v20
            case 128:
                var v21 float64 = (0.0 - 0.1)
                _ = v21
            }
            fmt.Println(v13, p0)
        }
        _ = v13
    }
    p0 = (p0 + p0)
    var v22 uint16 = 72
    switch 488 {
    case 32768:
        fmt.Println(p0, p0)
    case 128:
        fmt.Println(v22, v22)
    }
    _ = p0
    _ = v22
    return uint8(p0)
}

func main() {
    var v1 int32 = 0
    var v2 int = 1
    v3 := make(map[string]bool)
    v4 := make(map[uint32]uint32)
    v5 := []uint32{42, 519}
    var v6 float32 = 0.0
    var v23 uint16 = 818
    if (true || ((v1 - int32(v23)) < v1)) {
        switch v2 {
        case 0:
            var v24 float64 = 0.1
            _ = v24
        case 466:
            var v25 float32 = float32(((3.14159 - 3.14159) * 1.5))
            _ = v25
        case 519:
            fmt.Println(v6)
        default:
            var v26 int64 = -894
            _ = v26
        }
        var v27 float64 = (1.5 * 1.5)
        v2 += (v2 - 42)
        _ = v27
    }
    v23 += v23
    for fi28 := 0; fi28 < 10; fi28++ {
        v29 := GenS1{F0: 127, F1: 0, F2: -1}
        for fi30 := 0; fi30 < 6; fi30++ {
            v2 = fi30
            var v31 uint = 0
            v32 := GenS2{F0: "a", F1: 9223372036854775807, F2: (v23 | v23)}
            _ = fi30
            _ = v31
            _ = v32
        }
        v6 -= (v6 / (v6 + v6))
        _ = fi28
        _ = v29
    }
    var v33 int = v2
    v34 := GenS3{F0: 18446744073709551615, F1: v1, F2: "\n\t"}
    var v35 int8 = 100
    fmt.Println(v33, v6, v33)
    switch (v2 & (v2 % ((432) | 1))) {
    case 128:
        v36 := GenS4{F0: -9223372036854775808, F1: "hello", F2: 408}
        _ = v36
    case -9223372036854775808:
        v33 = (^(v2))
    case 1:
        switch v33 {
        case -9223372036854775808:
            if (42 <= 958) {
                switch "" {
                case "a":
                    if true {
                        fmt.Println(v2, v23)
                        for fi37 := 0; fi37 < 10; fi37++ {
                            if true {
                                v33 = fi37
                                var v38 float64 = 0.1
                                v2 += -2147483648
                                _ = v38
                            } else {
                                v2 = (fi37 + (fi37 | 128))
                            }
                            fmt.Println(v35)
                            v23 = (v23 - v23)
                            _ = fi37
                        }
                    } else {
                        if true {
                            fmt.Println(v1, v1, v33)
                            if (!true) {
                                for fi39 := 0; fi39 < 18; fi39++ {
                                    if false {
                                        var v40 float64 = float64(float32(0.1))
                                        _ = v40
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        if (((fi39 - (fi39 | -214)) < (v33 % (((fi39 << ((2) & 63))) | 1))) && true) {
                                            var v41 int8 = (v35 ^ (v35 << ((4294967295) & 63)))
                                            _ = v41
                                        } else {
                                            if ("" >= "foo bar") {
                                                if false {
                                                    fmt.Println(v35)
                                                    switch (fi39 ^ fi39) {
                                                    case 9223372036854775807:
                                                        for fi42 := 0; fi42 < 19; fi42++ {
                                                            switch ("hello" + "") {
                                                            case "foo bar":
                                                                v43 := GenS5{F0: 1}
                                                                _ = v43
                                                            case "":
                                                                if len(v5) > 0 {
                                                                    v5[0] = 904
                                                                }
                                                                v6 = 0.0
                                                                if (true || (2 != 102)) {
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                }
                                                            case "a":
                                                                v44 := GenS2{F0: "", F1: 9223372036854775807, F2: (v23 >> ((65535) & 63))}
                                                                _ = v44
                                                            }
                                                            _ = fi42
                                                        }
                                                    default:
                                                        if false {
                                                            fmt.Println(v6, v23, fi39)
                                                            for fi45 := 0; fi45 < 9; fi45++ {
                                                                v6 += 0.1
                                                                v46 := GenS1{F0: (v35 & 42), F1: 141, F2: 67}
                                                                _ = fi45
                                                                _ = v46
                                                            }
                                                            var v47 float64 = 1e10
                                                            _ = v47
                                                        } else {
                                                            v48 := GenS6{F0: "hello"}
                                                            v49 := GenS7{F0: v1, F1: v35, F2: false}
                                                            _ = v48
                                                            _ = v49
                                                        }
                                                    }
                                                } else {
                                                    if len(v5) > 0 {
                                                        v5[0] = 1
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    v6 += v6
                                    _ = fi39
                                }
                                if true {
                                    var v50 uint64 = 0
                                    var v51 uint16 = 1
                                    v33 += 32768
                                    _ = v50
                                    _ = v51
                                } else {
                                    if ((((0.0 * 0.0) < 1.5) || true) || (true || false)) {
                                        fmt.Println(v23, v2, v1)
                                        var v52 uint = 0
                                        _ = v52
                                    } else {
                                        delete(v4, 4294967295)
                                        if len(v5) > 0 {
                                            v5[0] = 65535
                                        }
                                    }
                                    var v53 float64 = 1.5
                                    _ = v53
                                }
                                if (false && (((0.0 + 1e10) / (1.5 / 1.0)) == 3.14159)) {
                                    var v54 int = v33
                                    _ = v54
                                } else {
                                    for fi55 := 0; fi55 < 20; fi55++ {
                                        var v56 float64 = float64(float32((0.0 / 0.1)))
                                        v23 = v23
                                        for fi57 := 0; fi57 < 5; fi57++ {
                                            v1 = (v1 | (-(v1)))
                                            v58 := GenS8{F0: (v23 + 130), F1: "", F2: (v1 << ((601) & 63))}
                                            _ = fi57
                                            _ = v58
                                        }
                                        _ = fi55
                                        _ = v56
                                    }
                                }
                            } else {
                                var v59 float32 = (v6 / 1.0)
                                v6 = ((v59 - 0.1) * (v59 * ((v59 - (v59 * (v6 * v6))) / v6)))
                                _ = v59
                            }
                        } else {
                            switch ("\n\t" + (("\n\t" + "hello") + ("foo bar" + ("hello" + "a")))) {
                            case "hello":
                                var v60 uint8 = genFn1(2147483648)
                                _ = v60
                            default:
                                v1 -= v1
                            }
                        }
                        v23 = v23
                    }
                }
                v61 := GenS9{F0: 2147483648, F1: 531}
                _ = v61
            }
        case 9223372036854775807:
            if false {
                v23 = 131
                var v62 uint8 = 0
                var v63 uint8 = genFn1(2147483647)
                _ = v62
                _ = v63
            }
        case -2147483648:
            v64 := GenS3{F0: 2, F1: (v1 + v1), F2: ("a" + "")}
            _ = v64
        }
    }
    if true {
        var v65 uint16 = (^(v23))
        _ = v65
    }
    for fi66 := 0; fi66 < 2; fi66++ {
        v6 -= (v6 - (v6 * -2.25))
        var v67 float32 = ((3.14159 * v6) * v6)
        _ = fi66
        _ = v67
    }
    v68 := GenS2{F0: ("foo bar" + "\n\t"), F1: -9223372036854775808, F2: 2}
    delete(v4, 1)
    var v69 int16 = 497
    var v70 int32 = v1
    switch v33 {
    case 980:
        v71 := v4[2147483648]
        _ = v71
    case 1:
        v72 := GenS10{F0: 1, F1: (-(v70))}
        _ = v72
    case 127:
        if (true && false) {
            if (65535 != v23) {
                if true {
                    for fi73 := 0; fi73 < 4; fi73++ {
                        if false {
                            var v74 int = fi73
                            _ = v74
                        }
                        var v75 uint16 = v23
                        _ = fi73
                        _ = v75
                    }
                    v76 := GenS6{F0: (("a" + "hello") + "hello")}
                    _ = v76
                }
            }
            v35 += v35
            if (545 < -689) {
                fmt.Println(v23, v6)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                for fi77 := 0; fi77 < 11; fi77++ {
                    v78 := GenS2{F0: ("\n\t" + ""), F1: 891, F2: v23}
                    _ = fi77
                    _ = v78
                }
            }
        }
    }
    var v79 int64 = -805
    v79 = (v79 * -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v69
    _ = v70
    _ = v79
}
