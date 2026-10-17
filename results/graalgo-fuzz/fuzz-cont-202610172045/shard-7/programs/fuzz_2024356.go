package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int64 }
type GenS2 struct { F0 int64; F1 uint; F2 uint32 }
type GenS3 struct { F0 int64; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 int; F1 bool; F2 uint }
type GenS9 struct { F0 uint8; F1 string; F2 uint8 }
func genFn0(p0 int64) uint8 {
    v7 := GenS1{F0: 255, F1: true, F2: p0}
    v8 := GenS1{F0: 65535, F1: false, F2: p0}
    _ = p0
    _ = v7
    _ = v8
    return 128
}
func genFn1(p0 uint32, p1 uint) string {
    switch 128 {
    case 4294967295:
        v9 := GenS1{F0: 629, F1: (true && false), F2: 0}
        _ = v9
    case -1:
        fmt.Println(p1)
    }
    v10 := []int{749, 2147483647, 42}
    switch "foo bar" {
    case "\n\t":
        if true {
            v10 = append(v10, int(p0))
            p0 = 0
        } else {
            for fi11 := 0; fi11 < 13; fi11++ {
                fmt.Println(fi11, p0)
                var v12 uint32 = (p0 - 42)
                v12 = p0
                _ = fi11
                _ = v12
            }
            v10 = append(v10, 32767)
            for fi13 := 0; fi13 < 5; fi13++ {
                for fi14 := 0; fi14 < 19; fi14++ {
                    p0 -= (^(p0))
                    fmt.Println(p1, fi14)
                    _ = fi14
                }
                for fi15 := 0; fi15 < 16; fi15++ {
                    var v16 uint8 = 128
                    v16 = 212
                    var v17 float32 = ((0.0 - (1.5 + 1e10)) - 0.1)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                if len(v10) > 0 {
                    v10[0] = fi13
                }
                _ = fi13
            }
        }
    default:
        for fi18 := 0; fi18 < 3; fi18++ {
            fmt.Println(p0, p0, p0)
            v19 := GenS1{F0: 0, F1: (-1 <= -1), F2: -1}
            _ = fi18
            _ = v19
        }
    }
    v20 := GenS1{F0: 42, F1: (p1 >= p1), F2: 0}
    if true {
        switch int(p0) {
        case 0:
            v21 := GenS2{F0: 9223372036854775807, F1: (p1 / ((p1) | 1)), F2: (^(p0))}
            _ = v21
        default:
            p1 = (p1 * p1)
        }
        if false {
            fmt.Println(p1)
            switch "a" {
            case "\n\t":
                if true {
                    v22 := GenS3{F0: -528, F1: 100, F2: (p0 ^ p0)}
                    fmt.Println(p1)
                    var v23 int32 = 0
                    _ = v22
                    _ = v23
                } else {
                    p1 += 42
                    var v24 int8 = 110
                    var v25 bool = false
                    _ = v24
                    _ = v25
                }
            case "a":
                if false {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    if (42 >= 127) {
                        for fi26 := 0; fi26 < 19; fi26++ {
                            p1 -= 255
                            v27 := GenS4{F0: 255}
                            var v28 int32 = -2147483648
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        for fi29 := 0; fi29 < 14; fi29++ {
                            v30 := GenS5{F0: 213, F1: (^(fi29))}
                            var v31 int16 = -1
                            for fi32 := 0; fi32 < 19; fi32++ {
                                v33 := GenS6{F0: -128}
                                _ = fi32
                                _ = v33
                            }
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        p1 += (p1 * p1)
                    } else {
                        if true {
                            v34 := GenS5{F0: 128, F1: 42}
                            fmt.Println(p0, p0)
                            _ = v34
                        }
                    }
                } else {
                    switch 506 {
                    case -9223372036854775808:
                        p1 += (p1 - (^(p1)))
                    default:
                        var v35 uint64 = 0
                        _ = v35
                    }
                }
            default:
                for fi36 := 0; fi36 < 15; fi36++ {
                    p0 -= (p0 | (p0 | (p0 * (p0 - p0))))
                    for fi37 := 0; fi37 < 8; fi37++ {
                        v38 := GenS7{F0: (!false), F1: (fi36 + (^(fi36)))}
                        _ = fi37
                        _ = v38
                    }
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = fi36
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v20
    return ""
}

func main() {
    v1 := []int{}
    v2 := make(map[int]string)
    var v3 uint32 = 337
    var v4 int32 = 1
    var v5 uint16 = 65535
    var v6 uint64 = 4294967295
    for fi39 := 0; fi39 < 8; fi39++ {
        fmt.Println(v3, v4, v4)
        v4 += (v4 + (v4 << ((18446744073709551615) & 63)))
        v3 = uint32(v6)
        _ = fi39
    }
    v1 = append(v1, -901)
    delete(v2, 139)
    for fi40 := 0; fi40 < 18; fi40++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if false {
            var v41 float32 = 1e10
            v41 += (v41 / 1.5)
            v5 = (v5 << ((384) & 63))
            _ = v41
        } else {
            v5 -= v5
            v3 = 4294967295
            if true {
                v42 := GenS4{F0: v5}
                var v43 int16 = 1000
                fmt.Println(v43, v3, v3)
                _ = v42
                _ = v43
            } else {
                v5 = 2
            }
        }
        _ = fi40
    }
    if false {
        _ = len(v2)
        v1 = append(v1, -407)
        for fi44 := 0; fi44 < 12; fi44++ {
            if ("a" < "") {
                var v45 int = (fi44 | (fi44 << ((801) & 63)))
                _ = v45
            }
            _ = fi44
        }
    } else {
        for fi46 := 0; fi46 < 15; fi46++ {
            var v47 int32 = v4
            v48 := v2[(fi46 >> ((604) & 63))]
            var v49 uint8 = 255
            _ = fi46
            _ = v47
            _ = v48
            _ = v49
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v50 int32 = v4
        _ = v50
    }
    for fi51 := 0; fi51 < 19; fi51++ {
        if (true || (100 >= 63)) {
            if ((true || (!true)) && (!((763 > 4294967295) && true))) {
                for fi52 := 0; fi52 < 9; fi52++ {
                    v5 = v5
                    _ = fi52
                }
            } else {
                for fi53 := 0; fi53 < 5; fi53++ {
                    _ = len(v1)
                    v6 -= v6
                    v54 := GenS8{F0: -9223372036854775808, F1: (false || ((1 <= 1) || (((float64(1.5) / ((1e10 / ((float64(1e10) / 0.1) / 1.5)) - 1.5)) + -2.25) < ((3.14159 / 1.5) - (0.0 / (1.5 * -2.25)))))), F2: 4294967295}
                    _ = fi53
                    _ = v54
                }
                v55 := GenS8{F0: (-(fi51)), F1: false, F2: 2}
                _ = v55
            }
            v56 := GenS5{F0: genFn0(-401), F1: fi51}
            v57 := v2[42]
            _ = v56
            _ = v57
        } else {
            for fi58 := 0; fi58 < 4; fi58++ {
                switch (("\n\t" + "foo bar") + "\n\t") {
                case "\n\t":
                    switch "" {
                    case "\n\t":
                        _ = len(v1)
                    case "foo bar":
                        var v59 uint8 = 2
                        _ = v59
                    case "hello":
                        v6 -= 42
                    default:
                        var v60 uint64 = (^(v6))
                        _ = v60
                    }
                case "a":
                    v5 += 1
                case "":
                    v5 = v5
                }
                fmt.Println(v6, v6, fi51)
                var v61 uint64 = (v6 << ((468) & 63))
                _ = fi58
                _ = v61
            }
        }
        v1 = append(v1, 32767)
        v3 += 127
        _ = fi51
    }
    if (42 != -9223372036854775808) {
        var v62 int16 = 32767
        fmt.Println(v3)
        if (126 < 1) {
            for fi63 := 0; fi63 < 3; fi63++ {
                if true {
                    fmt.Println(v5, v6)
                } else {
                    var v64 bool = ("a" <= "\n\t")
                    if (!true) {
                        for fi65 := 0; fi65 < 19; fi65++ {
                            v62 = (v62 & (-(v62)))
                            v66 := GenS8{F0: fi65, F1: (!true), F2: 2}
                            _ = fi65
                            _ = v66
                        }
                    }
                    var v67 float32 = (float32((-2.25 - float64(-2.25))) * ((0.1 * ((0.0 / (-2.25 * (1e10 + (float32((0.1 * 3.14159)) / -2.25)))) + 0.1)) * 0.1))
                    _ = v64
                    _ = v67
                }
                v5 += uint16(v6)
                switch (-(fi63)) {
                case 1:
                    for fi68 := 0; fi68 < 19; fi68++ {
                        if false {
                            switch fi68 {
                            case 516:
                                var v69 uint64 = 393
                                v6 = 721
                                v69 += (^(v69))
                                v69 = 1
                                _ = v69
                            default:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            var v70 uint16 = 541
                            v5 = 1
                            _ = v70
                        }
                        if false {
                            v3 += (v3 / (((^(v3))) | 1))
                        }
                        fmt.Println(v4, v3)
                        _ = fi68
                    }
                case 128:
                    for fi71 := 0; fi71 < 16; fi71++ {
                        var v72 uint8 = 73
                        switch "\n\t" {
                        case "hello":
                            v73 := GenS9{F0: (v72 - (v72 * (^(v72)))), F1: ("" + (("" + ("" + (((("foo bar" + ((("\n\t" + ("foo bar" + "foo bar")) + "foo bar") + ("" + "foo bar"))) + (("a" + "\n\t") + ("" + "foo bar"))) + "hello") + ("" + "foo bar")))) + "")), F2: (v72 % (((v72 * 127)) | 1))}
                            _ = v73
                        case "a":
                            var v74 uint64 = (v6 ^ v6)
                            _ = v74
                        }
                        _ = fi71
                        _ = v72
                    }
                case -1:
                    fmt.Println(v62)
                }
                _ = fi63
            }
            v75 := v2[662]
            var v76 uint = 1
            _ = v75
            _ = v76
        } else {
            switch "a" {
            case "hello":
                v4 += v4
            default:
                v6 += (^(v6))
            }
            switch "hello" {
            case "hello":
                fmt.Println(v5, v5, v4)
            default:
                for fi77 := 0; fi77 < 13; fi77++ {
                    v4 -= (^(v4))
                    _ = fi77
                }
            }
            var v78 uint32 = 2
            _ = v78
        }
        _ = v62
    } else {
        switch -1 {
        case 128:
            v5 += (v5 ^ (v5 / (((v5 >> ((559) & 63))) | 1)))
        case 4294967295:
            if true {
                v79 := GenS2{F0: 1, F1: 18446744073709551615, F2: v3}
                _ = v79
            }
        }
        fmt.Println(v5, v4)
    }
    v6 = v6
    if (uint8(v6) > 2) {
        var v80 uint16 = v5
        for fi81 := 0; fi81 < 11; fi81++ {
            if true {
                v2[fi81] = ("foo bar" + "\n\t")
                if false {
                    var v82 int16 = -1
                    _ = v82
                } else {
                    var v83 int8 = int8(v5)
                    v80 -= v5
                    _ = v83
                }
            } else {
                var v84 uint64 = 42
                if false {
                    v3 = v3
                } else {
                    v85 := GenS9{F0: 0, F1: "a", F2: 128}
                    var v86 bool = true
                    if ((v80 | 255) >= (v80 & (^(v80)))) {
                        v4 -= (v4 & -837)
                        v80 = 65535
                    } else {
                        fmt.Println(v80, v6, v4)
                        var v87 uint64 = (v6 - uint64(v3))
                        _ = v87
                    }
                    _ = v85
                    _ = v86
                }
                v6 = (v84 * (v84 - v84))
                _ = v84
            }
            _ = fi81
        }
        _ = v80
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
