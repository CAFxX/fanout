package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int8 }
type GenS2 struct { F0 bool; F1 uint16; F2 int }
type GenS3 struct { F0 int16; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 int64; F1 uint16 }
func genFn0(p0 int, p1 int16) int8 {
    p0 = p0
    p1 += -10
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 uint = 2
            fmt.Println(fi8, v9)
            var v10 int64 = 9223372036854775807
            _ = fi8
            _ = v10
            _ = v9
        }
        if false {
            var v11 float64 = (0.1 - 1.5)
            _ = v11
        } else {
            v12 := []int{fi7, (p0 + fi7), (fi7 >> ((18446744073709551615) & 63)), (p0 % ((p0) | 1))}
            v13 := GenS1{F0: ("foo bar" <= ""), F1: 42, F2: -81}
            _ = v12
            _ = v13
        }
        _ = fi7
    }
    v14 := []int{528, (p0 + (^(p0))), p0}
    _ = p0
    _ = p1
    _ = v14
    return int8(p0)
}
func genFn1(p0 uint64, p1 int8, p2 uint16) int16 {
    if ((-301 >= -2147483648) && true) {
        p2 = p2
        p2 -= (p2 >> ((2) & 63))
    } else {
        var v15 uint16 = 2
        if false {
            if true {
                v16 := []int{-9223372036854775808, 2147483647, -755}
                if len(v16) > 0 {
                    v16[0] = 32767
                }
                _ = v16
            }
        } else {
            if (!false) {
                if (-32768 <= 0) {
                    p2 = (v15 | (v15 + 163))
                } else {
                    for fi17 := 0; fi17 < 2; fi17++ {
                        if false {
                            var v18 int32 = int32(p2)
                            _ = v18
                        }
                        var v19 int64 = 9223372036854775807
                        var v20 int = fi17
                        _ = fi17
                        _ = v19
                        _ = v20
                    }
                    fmt.Println(p2)
                    fmt.Println(p1, p2)
                }
                v21 := GenS2{F0: (p0 != 18446744073709551615), F1: p2, F2: -1}
                v15 -= (^(p2))
                _ = v21
            }
            v22 := []int{128, 127, 128, 774}
            if (false || true) {
                var v23 uint8 = 128
                _ = v23
            } else {
                var v24 int = -9223372036854775808
                v22 = append(v22, 42)
                _ = v24
            }
            _ = v22
        }
        if (((p0 & 1) > (p0 | (^(p0)))) || false) {
            v15 = p2
        } else {
            if false {
                if true {
                    v25 := GenS1{F0: true, F1: v15, F2: p1}
                    if (false || false) {
                        for fi26 := 0; fi26 < 17; fi26++ {
                            for fi27 := 0; fi27 < 19; fi27++ {
                                if false {
                                    var v28 float64 = (float64((-2.25 * 1.5)) - 0.1)
                                    v29 := GenS1{F0: (false || ((-(p1)) >= p1)), F1: (p2 << ((65535) & 63)), F2: int8(p0)}
                                    _ = v28
                                    _ = v29
                                } else {
                                    fmt.Println(p0, p0)
                                }
                                switch fi26 {
                                case -9223372036854775808:
                                    var v30 float32 = float32((float64((0.0 / (float32(1e10) * 1.5))) * (((0.1 + 1e10) / 1.0) - (0.1 - 3.14159))))
                                    _ = v30
                                case 32768:
                                    p1 -= 1
                                case 9223372036854775807:
                                    v31 := []int{}
                                    _ = v31
                                default:
                                    fmt.Println(fi27, fi26, p0)
                                }
                                p0 = p0
                                _ = fi27
                            }
                            _ = fi26
                        }
                        for fi32 := 0; fi32 < 7; fi32++ {
                            v33 := GenS3{F0: 1, F1: (p0 ^ p0), F2: p1}
                            _ = fi32
                            _ = v33
                        }
                    }
                    _ = v25
                }
                var v34 uint16 = (^(v15))
                _ = v34
            }
        }
        _ = v15
    }
    v35 := []int{-2147483648, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v35
    return 1000
}

func main() {
    var v1 int8 = 126
    v2 := make(map[uint32]bool)
    var v3 bool = (!(v1 != (v1 % (((v1 % (((^(v1))) | 1))) | 1))))
    var v4 uint = 0
    v5 := []uint32{42, 47, 2147483648}
    v6 := []string{("\n\t" + "hello")}
    fmt.Println(v1, v1)
    if len(v6) > 0 {
        v6[0] = "hello"
    }
    var v36 bool = (v3 || v3)
    fmt.Println(v4, v3)
    v37 := GenS4{F0: (v4 - v4)}
    var v38 uint16 = 42
    v4 -= v4
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    if true {
        for fi39 := 0; fi39 < 16; fi39++ {
            fmt.Println(fi39, v38)
            _ = fi39
        }
    } else {
        v2[2147483648] = false
    }
    v40 := GenS2{F0: v36, F1: v38, F2: 997}
    if (true && (2147483647 <= 1)) {
        switch 128 {
        case 543:
            for fi41 := 0; fi41 < 13; fi41++ {
                v36 = v36
                _ = fi41
            }
        case 32768:
            for fi42 := 0; fi42 < 20; fi42++ {
                var v43 bool = v36
                var v44 uint = 18446744073709551615
                _ = fi42
                _ = v43
                _ = v44
            }
        }
        if (4294967295 >= 0) {
            var v45 uint = 65535
            var v46 int16 = 1000
            _ = v45
            _ = v46
        } else {
            v47 := GenS5{F0: -1}
            if (!(!(true && false))) {
                _ = len(v6)
            } else {
                v1 -= genFn0(-999, 0)
                for fi48 := 0; fi48 < 17; fi48++ {
                    var v49 int8 = 1
                    v1 -= v1
                    _ = fi48
                    _ = v49
                }
                switch "foo bar" {
                case "a":
                    if v3 {
                        for fi50 := 0; fi50 < 14; fi50++ {
                            fmt.Println(v36, v1, v1)
                            _ = fi50
                        }
                        for fi51 := 0; fi51 < 2; fi51++ {
                            var v52 float32 = (1.5 - (-2.25 / (float32(float64(3.14159)) * (1e10 + ((0.0 - 1.5) + 1.5)))))
                            for fi53 := 0; fi53 < 16; fi53++ {
                                if (!v3) {
                                    if false {
                                        if (true || (1 == 599)) {
                                            v38 += (v38 / (((^(v38))) | 1))
                                        }
                                        v54 := GenS6{F0: 1, F1: 0}
                                        v55 := GenS7{F0: (v1 % (((v1 % (((v1 << ((v4) & 63))) | 1))) | 1)), F1: 2147483647}
                                        _ = v54
                                        _ = v55
                                    } else {
                                        var v56 bool = (2147483647 >= fi53)
                                        fmt.Println(v52)
                                        _ = v56
                                    }
                                    var v57 int16 = -1
                                    switch "a" {
                                    case "hello":
                                        var v58 uint64 = 566
                                        _ = v58
                                    case "foo bar":
                                        v59 := GenS8{F0: 9223372036854775807, F1: (v38 + v38)}
                                        _ = v59
                                    }
                                    _ = v57
                                }
                                var v60 int64 = -9223372036854775808
                                _ = fi53
                                _ = v60
                            }
                            if false {
                                var v61 int = (fi51 + (fi51 << ((2) & 63)))
                                _ = v61
                            }
                            _ = fi51
                            _ = v52
                        }
                    }
                }
            }
            _ = v47
        }
    }
    v2[308] = v3
    var v62 uint64 = 473
    var v63 int16 = 0
    v4 = v4
    fmt.Println(v62, v62, v62)
    for fi64 := 0; fi64 < 13; fi64++ {
        v62 += v62
        var v65 bool = (v62 < 1)
        switch "hello" {
        case "a":
            var v66 int64 = 0
            _ = v66
        }
        _ = fi64
        _ = v65
    }
    if (!((-369 >= 42) || v3)) {
        v67 := GenS5{F0: -1}
        for fi68 := 0; fi68 < 18; fi68++ {
            v69 := GenS6{F0: -775, F1: -2147483648}
            var v70 uint64 = v62
            _ = fi68
            _ = v69
            _ = v70
        }
        var v71 string = "foo bar"
        _ = v67
        _ = v71
    } else {
        if false {
            var v72 int16 = v63
            v2[800] = v3
            _ = v72
        } else {
            switch "" {
            case "\n\t":
                if len(v6) > 0 {
                    v6[0] = ("foo bar" + ("" + ""))
                }
            default:
                var v73 int32 = 697
                _ = v73
            }
            var v74 bool = v3
            switch 42 {
            case -224:
                var v75 bool = v3
                _ = v75
            case 128:
                var v76 int8 = genFn0(int(v1), v63)
                _ = v76
            case 1:
                var v77 int64 = -698
                _ = v77
            }
            _ = v74
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v62
    _ = v63
}
