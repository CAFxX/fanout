package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 uint16; F1 string }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint64, p1 string) int {
    var v4 int = 32767
    if (!false) {
        switch p1 {
        case "":
            v5 := GenS1{F0: 65, F1: 2147483647, F2: 2}
            _ = v5
        case "a":
            fmt.Println(v4, v4, p0)
        }
        var v6 int8 = -128
        if true {
            v7 := []int{9223372036854775807}
            _ = v7
        }
        _ = v6
    }
    v8 := GenS1{F0: 1, F1: 0, F2: 1}
    _ = p0
    _ = p1
    _ = v4
    _ = v8
    return v4
}
func genFn1(p0 int16, p1 int16, p2 int64) int8 {
    if (2 != 2) {
        p1 = (p0 - p0)
        v9 := []int{42, 769, -2147483648, -95}
        var v10 bool = false
        _ = v10
        _ = v9
    }
    v11 := []int{-9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 0
}

func main() {
    v1 := []int{4294967295, -9223372036854775808, 473}
    var v2 uint64 = 0
    v3 := []uint32{2147483647}
    v2 = (v2 ^ (v2 - v2))
    var v12 int32 = -2147483648
    fmt.Println(v2, v12)
    v13 := GenS2{F0: false, F1: 1000}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v14 := GenS3{F0: 2147483647}
    fmt.Println(v12, v12)
    if true {
        var v15 int64 = 42
        _ = v15
    } else {
        v2 += v2
    }
    var v16 bool = false
    fmt.Println(v12)
    fmt.Println(v2)
    v1 = append(v1, -117)
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    if true {
        if true {
            for fi17 := 0; fi17 < 19; fi17++ {
                if ((fi17 * genFn0(61, "a")) == fi17) {
                    for fi18 := 0; fi18 < 5; fi18++ {
                        if ((0.0 < 0.0) && (1 < 42)) {
                            v19 := GenS4{F0: 2147483648, F1: 42}
                            if v16 {
                                var v20 uint = uint(v2)
                                _ = v20
                            } else {
                                fmt.Println(fi18, v16, fi17)
                                if v16 {
                                    v2 = 1
                                    var v21 int16 = -1
                                    _ = v21
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    fmt.Println(v2, fi18, v12)
                                    fmt.Println(v2, v12, fi17)
                                }
                                var v22 int8 = -1
                                _ = v22
                            }
                            if (42 <= 357) {
                                v23 := GenS3{F0: (v12 << ((65535) & 63))}
                                var v24 float32 = -2.25
                                v12 -= (v12 ^ v12)
                                _ = v23
                                _ = v24
                            }
                            switch 576 {
                            case 2147483647:
                                if (!((87 < 206) || ((v12 / ((-1) | 1)) <= (v12 << ((318) & 63))))) {
                                    v12 = (v12 + 2147483647)
                                }
                            case -511:
                                switch (fi18 * genFn0(v2, ("\n\t" + "foo bar"))) {
                                case 0:
                                    v25 := GenS1{F0: 100, F1: 1, F2: 1}
                                    _ = v25
                                case 32767:
                                    var v26 int8 = int8(v12)
                                    _ = v26
                                default:
                                    var v27 int = genFn0(v2, "\n\t")
                                    _ = v27
                                }
                            case -477:
                                _ = len(v3)
                            default:
                                switch 32768 {
                                case 9223372036854775807:
                                    v12 += v12
                                case 506:
                                    v2 = v2
                                default:
                                    switch (("\n\t" + ("foo bar" + ("\n\t" + ""))) + "a") {
                                    case "\n\t":
                                        v2 = v2
                                    case "hello":
                                        v16 = v16
                                    case "":
                                        if true {
                                            v12 = (v12 & (v12 << ((65535) & 63)))
                                        } else {
                                            v28 := GenS1{F0: 113, F1: 42, F2: 2}
                                            if (!((42 <= 0) || v16)) {
                                                fmt.Println(v2, v2)
                                                _ = len(v3)
                                            }
                                            var v29 int8 = -1
                                            _ = v28
                                            _ = v29
                                        }
                                    default:
                                        var v30 int32 = 42
                                        _ = v30
                                    }
                                }
                                v31 := GenS2{F0: true, F1: -1}
                                fmt.Println(v12)
                                _ = v31
                            }
                            _ = v19
                        }
                        var v32 int16 = 42
                        _ = fi18
                        _ = v32
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = 2147483647
                    }
                }
                _ = fi17
            }
            fmt.Println(v16, v2)
        }
        v33 := GenS2{F0: (-9223372036854775808 >= int64(v12)), F1: 1000}
        switch 32768 {
        case -9223372036854775808:
            var v34 float64 = 0.1
            _ = v34
        case 32:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            if (-717 > 0) {
                v12 -= -946
                v16 = (-2.25 < float32(((3.14159 - 0.0) * 3.14159)))
            }
        }
        _ = v33
    }
    var v35 int = 792
    for fi36 := 0; fi36 < 18; fi36++ {
        if v16 {
            if (((18446744073709551615 != 1) || true) || v16) {
                switch 125 {
                case 32768:
                    if (!(("hello" < "hello") || ((!v16) && v16))) {
                        v37 := GenS1{F0: genFn1(-845, -1, -553), F1: 2147483648, F2: 510}
                        _ = v37
                    } else {
                        var v38 uint32 = 691
                        if (!(1e10 <= 1e10)) {
                            for fi39 := 0; fi39 < 13; fi39++ {
                                var v40 uint64 = 0
                                if v16 {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    _ = len(v3)
                                }
                                _ = fi39
                                _ = v40
                            }
                        }
                        if (!(!v16)) {
                            switch 0 {
                            case 669:
                                switch "foo bar" {
                                case "":
                                    if len(v1) > 0 {
                                        v1[0] = 0
                                    }
                                case "hello":
                                    if v16 {
                                        v16 = (2147483647 >= (v12 / ((int32(v38)) | 1)))
                                        v38 += (v38 + (v38 * v38))
                                        for fi41 := 0; fi41 < 13; fi41++ {
                                            v42 := GenS5{F0: ("foo bar" + "\n\t"), F1: 1}
                                            _ = fi41
                                            _ = v42
                                        }
                                    }
                                case "\n\t":
                                    for fi43 := 0; fi43 < 6; fi43++ {
                                        var v44 int32 = -367
                                        _ = fi43
                                        _ = v44
                                    }
                                }
                            case 0:
                                v45 := GenS2{F0: (604 == 42), F1: 32767}
                                _ = v45
                            case 801:
                                var v46 float32 = 0.0
                                _ = v46
                            default:
                                switch (^(v35)) {
                                case 128:
                                    var v47 uint32 = 785
                                    _ = v47
                                case 599:
                                    for fi48 := 0; fi48 < 20; fi48++ {
                                        if (!false) {
                                            var v49 int64 = 9223372036854775807
                                            _ = v49
                                        } else {
                                            fmt.Println(v16)
                                            var v50 uint16 = 32767
                                            v51 := GenS5{F0: (("a" + "a") + "foo bar"), F1: 9223372036854775807}
                                            _ = v50
                                            _ = v51
                                        }
                                        _ = fi48
                                    }
                                case -452:
                                    switch ("hello" + "foo bar") {
                                    case "":
                                        if false {
                                            v38 -= 787
                                        } else {
                                            if (255 <= 255) {
                                                if len(v1) > 0 {
                                                    v1[0] = (v35 & fi36)
                                                }
                                                if len(v3) > 0 {
                                                    v3[0] = v38
                                                }
                                            }
                                            switch 2147483647 {
                                            case -312:
                                                for fi52 := 0; fi52 < 20; fi52++ {
                                                    v53 := GenS6{F0: 128, F1: 127}
                                                    _ = len(v1)
                                                    v54 := GenS6{F0: 1, F1: 127}
                                                    _ = fi52
                                                    _ = v53
                                                    _ = v54
                                                }
                                            }
                                            if false {
                                                if len(v3) > 0 {
                                                    v3[0] = (v38 * v38)
                                                }
                                                fmt.Println(v2)
                                            } else {
                                                v12 = (v12 & -163)
                                                if (!false) {
                                                    v55 := GenS7{F0: 527, F1: "foo bar"}
                                                    _ = v55
                                                } else {
                                                    if (((^(v38)) < 822) || false) {
                                                        for fi56 := 0; fi56 < 14; fi56++ {
                                                            var v57 int = (-(v35))
                                                            var v58 int32 = 0
                                                            _ = fi56
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                        v59 := GenS7{F0: 32767, F1: "hello"}
                                                        _ = v59
                                                    } else {
                                                        v38 = v38
                                                    }
                                                    v60 := GenS6{F0: 127, F1: 0}
                                                    for fi61 := 0; fi61 < 2; fi61++ {
                                                        v16 = true
                                                        _ = len(v3)
                                                        var v62 float32 = 3.14159
                                                        _ = fi61
                                                        _ = v62
                                                    }
                                                    _ = v60
                                                }
                                            }
                                        }
                                    case "\n\t":
                                        v16 = ((v2 / ((v2) | 1)) != (v2 % ((v2) | 1)))
                                    default:
                                        var v63 int32 = (v12 | (v12 << ((42) & 63)))
                                        _ = v63
                                    }
                                default:
                                    v64 := GenS7{F0: 2, F1: "\n\t"}
                                    _ = v64
                                }
                            }
                        } else {
                            var v65 string = "foo bar"
                            v16 = ((!(v16 && true)) || v16)
                            _ = v65
                        }
                        _ = v38
                    }
                case 9223372036854775807:
                    var v66 string = "a"
                    _ = v66
                }
            } else {
                v67 := GenS8{F0: v12}
                for fi68 := 0; fi68 < 11; fi68++ {
                    switch fi68 {
                    case -2147483648:
                        v35 -= v35
                    case 128:
                        fmt.Println(v16, fi68)
                    case 313:
                        var v69 uint64 = v2
                        _ = v69
                    default:
                        v3 = append(v3, 572)
                    }
                    v2 = v2
                    var v70 int8 = -39
                    _ = fi68
                    _ = v70
                }
                var v71 int32 = (v12 << ((4294967295) & 63))
                _ = v67
                _ = v71
            }
        }
        switch 522 {
        case 2147483647:
            switch "foo bar" {
            case "\n\t":
                switch v35 {
                case -1:
                    fmt.Println(fi36, v12, v12)
                case 668:
                    fmt.Println(v16, v2)
                }
            }
        case 521:
            v2 -= (v2 & v2)
        case -1:
            switch 128 {
            case 455:
                if ((0.0 - 0.0) == 1e10) {
                    if v16 {
                        fmt.Println(fi36)
                        var v72 int32 = v12
                        if len(v3) > 0 {
                            v3[0] = 0
                        }
                        _ = v72
                    } else {
                        v2 += (v2 / (((^(v2))) | 1))
                        var v73 float32 = float32(0.1)
                        _ = v73
                    }
                    fmt.Println(fi36)
                } else {
                    var v74 uint32 = 42
                    v75 := GenS9{F0: 32767, F1: 9223372036854775807}
                    v76 := GenS1{F0: 42, F1: 42, F2: 4294967295}
                    _ = v74
                    _ = v75
                    _ = v76
                }
            case -2147483648:
                v2 += v2
            }
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi36
    }
    var v77 uint16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v77
}
