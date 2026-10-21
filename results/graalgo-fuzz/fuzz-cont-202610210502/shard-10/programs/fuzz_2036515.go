package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 string; F1 uint; F2 uint8 }
type GenS9 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint16 }
func genFn0(p0 uint64, p1 int, p2 uint16) int {
    if true {
        if (42 <= 2) {
            if false {
                fmt.Println(p2)
                p2 -= 42
                var v5 int16 = 1
                _ = v5
            }
            v6 := GenS1{F0: 4294967295, F1: p2}
            _ = v6
        } else {
            if true {
                p1 = -452
            } else {
                for fi7 := 0; fi7 < 11; fi7++ {
                    p1 = 127
                    _ = fi7
                }
                for fi8 := 0; fi8 < 12; fi8++ {
                    if (!true) {
                        fmt.Println(fi8, p1, p0)
                    }
                    switch 435 {
                    case 876:
                        if false {
                            v9 := GenS2{F0: 2147483648}
                            p1 += (fi8 * 0)
                            var v10 uint64 = 476
                            _ = v10
                            _ = v9
                        }
                    case 32767:
                        var v11 int8 = -128
                        _ = v11
                    default:
                        for fi12 := 0; fi12 < 14; fi12++ {
                            var v13 float64 = (0.0 * (1e10 / 1.0))
                            _ = fi12
                            _ = v13
                        }
                    }
                    p2 -= 745
                    _ = fi8
                }
                v14 := []int{}
                _ = v14
            }
            p0 -= (p0 % ((p0) | 1))
        }
    } else {
        switch (("a" + "hello") + ((("hello" + "\n\t") + ("\n\t" + "a")) + "\n\t")) {
        case "hello":
            if true {
                if true {
                    p1 -= p1
                    if ((-1 > 1) || false) {
                        var v15 int8 = 97
                        var v16 uint16 = (p2 >> ((65535) & 63))
                        var v17 string = ("\n\t" + ("a" + "\n\t"))
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    var v18 int = p1
                    _ = v18
                } else {
                    v19 := GenS1{F0: 4294967295, F1: (^(p2))}
                    if (!false) {
                        if true {
                            switch ((((("a" + "foo bar") + "hello") + "\n\t") + "hello") + "") {
                            case "":
                                p2 += 42
                            case "hello":
                                p1 = p1
                            default:
                                var v20 int16 = -1
                                _ = v20
                            }
                            var v21 string = "\n\t"
                            var v22 bool = (!(!((p2 * (p2 / ((p2) | 1))) <= 841)))
                            _ = v21
                            _ = v22
                        } else {
                            p0 = (p0 ^ p0)
                        }
                        fmt.Println(p2)
                    }
                    _ = v19
                }
            }
        case "\n\t":
            fmt.Println(p0, p0, p1)
        default:
            v23 := []int{p1, p1, (p1 & (p1 * (^(p1)))), 1}
            _ = v23
        }
        fmt.Println(p0, p0)
        p2 = 1000
    }
    var v24 int8 = 0
    for fi25 := 0; fi25 < 16; fi25++ {
        for fi26 := 0; fi26 < 17; fi26++ {
            switch "\n\t" {
            case "a":
                switch "a" {
                case "a":
                    var v27 uint = 951
                    _ = v27
                case "":
                    fmt.Println(p0, p2, p2)
                }
            case "foo bar":
                var v28 float32 = 1e10
                _ = v28
            case "":
                if (true || true) {
                    v29 := GenS2{F0: 65535}
                    v24 -= -128
                    if true {
                        switch fi25 {
                        case 9223372036854775807:
                            var v30 int = p1
                            _ = v30
                        case 1:
                            v24 -= v24
                        case -478:
                            var v31 float32 = (1e10 - ((1.5 - -2.25) + -2.25))
                            _ = v31
                        }
                        switch "" {
                        case "foo bar":
                            v32 := GenS1{F0: 255, F1: p2}
                            _ = v32
                        case "a":
                            var v33 string = "foo bar"
                            _ = v33
                        default:
                            var v34 int8 = 0
                            _ = v34
                        }
                    }
                    _ = v29
                } else {
                    p0 = (p0 % (((p0 << ((255) & 63))) | 1))
                    var v35 float32 = 1.5
                    _ = v35
                }
            default:
                v36 := []int{(fi26 & (fi26 - (p1 | (fi25 + fi26))))}
                _ = v36
            }
            v24 = v24
            _ = fi26
        }
        for fi37 := 0; fi37 < 18; fi37++ {
            if (!false) {
                var v38 string = "a"
                v24 -= v24
                p2 += 2
                v39 := GenS2{F0: 0}
                _ = v38
                _ = v39
            } else {
                p2 = p2
            }
            _ = fi37
        }
        _ = fi25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return (-(p1))
}
func genFn1(p0 int8, p1 uint8) uint32 {
    v40 := []int{9223372036854775807, 32768, 294}
    p1 = (p1 * p1)
    var v41 uint8 = (p1 ^ 127)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v40
    _ = v41
    return 65535
}

func main() {
    var v1 string = "\n\t"
    var v2 int8 = 100
    var v3 float32 = 0.0
    v4 := []uint32{4294967295, 4294967295, 4294967295, 4294967295}
    for fi42 := 0; fi42 < 11; fi42++ {
        v43 := GenS3{F0: 4294967295, F1: 503, F2: 0}
        v44 := GenS4{F0: -9223372036854775808, F1: genFn1((v2 - 1), 2)}
        v4 = append(v4, 2)
        _ = fi42
        _ = v43
        _ = v44
    }
    if true {
        v45 := GenS5{F0: 128}
        _ = v45
    } else {
        fmt.Println(v3)
    }
    for fi46 := 0; fi46 < 18; fi46++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi46
    }
    for fi47 := 0; fi47 < 2; fi47++ {
        var v48 string = v1
        if true {
            var v49 int32 = -595
            v50 := GenS6{F0: 1}
            for fi51 := 0; fi51 < 19; fi51++ {
                var v52 uint = 65535
                if len(v4) > 0 {
                    v4[0] = 376
                }
                _ = fi51
                _ = v52
            }
            _ = v49
            _ = v50
        }
        _ = fi47
        _ = v48
    }
    switch v1 {
    case "hello":
        fmt.Println(v2)
    case "":
        v2 = v2
    case "a":
        for fi53 := 0; fi53 < 13; fi53++ {
            v1 = v1
            _ = fi53
        }
    }
    var v54 float32 = (v3 - 0.0)
    v55 := GenS5{F0: 61}
    var v56 uint8 = 127
    var v57 uint16 = 0
    v58 := GenS7{F0: 9223372036854775807, F1: 0}
    if len(v4) > 0 {
        v4[0] = genFn1((-(v2)), v56)
    }
    if false {
        if len(v4) > 0 {
            v4[0] = 1
        }
    }
    var v59 uint = 65535
    for fi60 := 0; fi60 < 7; fi60++ {
        for fi61 := 0; fi61 < 15; fi61++ {
            var v62 uint32 = 65535
            v2 -= -128
            _ = fi61
            _ = v62
        }
        _ = fi60
    }
    switch 1 {
    case 128:
        var v63 string = v1
        _ = v63
    case -2147483648:
        var v64 int32 = 919
        _ = v64
    case -1:
        for fi65 := 0; fi65 < 5; fi65++ {
            v54 += v3
            v2 = int8(fi65)
            for fi66 := 0; fi66 < 3; fi66++ {
                v4 = append(v4, genFn1(127, 0))
                _ = fi66
            }
            _ = fi65
        }
    }
    if (!((v2 % ((1) | 1)) > (v2 % ((v2) | 1)))) {
        v57 += 0
        _ = len(v4)
        for fi67 := 0; fi67 < 19; fi67++ {
            if ((v57 | (v57 ^ 32767)) > v57) {
                switch "" {
                case "":
                    for fi68 := 0; fi68 < 3; fi68++ {
                        if (84 > 292) {
                            var v69 int64 = 0
                            if false {
                                fmt.Println(v56, v57, v3)
                            } else {
                                if true {
                                    v70 := GenS6{F0: v69}
                                    switch (("foo bar" + "") + "foo bar") {
                                    case "":
                                        var v71 float64 = 1e10
                                        _ = v71
                                    case "foo bar":
                                        v56 -= 127
                                    default:
                                        v3 = v54
                                    }
                                    v4 = append(v4, 360)
                                    _ = v70
                                }
                                var v72 string = v1
                                _ = v72
                            }
                            switch 4294967295 {
                            case -9223372036854775808:
                                var v73 float64 = 1.5
                                _ = v73
                            case 93:
                                fmt.Println(v1, v3, fi67)
                            default:
                                v2 -= (v2 | v2)
                            }
                            fmt.Println(v69, v3, v3)
                            _ = len(v4)
                            _ = v69
                        } else {
                            _ = len(v4)
                        }
                        _ = fi68
                    }
                }
                fmt.Println(v3, v57)
                if false {
                    for fi74 := 0; fi74 < 15; fi74++ {
                        v56 = (v56 >> ((uint(v56)) & 63))
                        v75 := GenS8{F0: (((v1 + v1) + "") + v1), F1: (v59 >> ((v59) & 63)), F2: v56}
                        _ = fi74
                        _ = v75
                    }
                }
            }
            switch fi67 {
            case 32768:
                if true {
                    v76 := GenS6{F0: 42}
                    v56 -= 0
                    var v77 uint32 = 1
                    _ = v76
                    _ = v77
                }
            case 737:
                var v78 int16 = -755
                _ = v78
            case 4294967295:
                switch (fi67 - (fi67 ^ (fi67 % (((fi67 | (fi67 - (-(fi67))))) | 1)))) {
                case 233:
                    if (228 <= 42) {
                        var v79 bool = true
                        _ = v79
                    }
                }
            default:
                fmt.Println(fi67, v59, v54)
            }
            _ = fi67
        }
    } else {
        v56 -= v56
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if false {
            fmt.Println(v1, v57)
        } else {
            if true {
                var v80 uint64 = 2
                switch 32767 {
                case 1:
                    v81 := GenS9{F0: (v57 | v57), F1: -19, F2: 2}
                    _ = v81
                case -2147483648:
                    var v82 string = ""
                    _ = v82
                default:
                    for fi83 := 0; fi83 < 4; fi83++ {
                        if true {
                            if false {
                                v59 += (^(v59))
                                var v84 int32 = 0
                                _ = v84
                            } else {
                                for fi85 := 0; fi85 < 2; fi85++ {
                                    v80 += 1
                                    v2 = -128
                                    v80 = v80
                                    _ = fi85
                                }
                                v86 := GenS10{F0: 2}
                                _ = v86
                            }
                            var v87 uint64 = (v80 & (v80 & v80))
                            if (1.5 == float64(1.5)) {
                                for fi88 := 0; fi88 < 12; fi88++ {
                                    var v89 int = 32768
                                    var v90 int = v89
                                    v91 := GenS11{F0: 42}
                                    if false {
                                        v92 := GenS12{F0: 42}
                                        _ = len(v4)
                                        _ = v92
                                    }
                                    _ = fi88
                                    _ = v89
                                    _ = v90
                                    _ = v91
                                }
                                var v93 int32 = 42
                                fmt.Println(v1, v80)
                                _ = v93
                            } else {
                                var v94 uint8 = (v56 / ((v56) | 1))
                                v2 += 42
                                v54 = ((-2.25 / 1e10) - (1e10 + v3))
                                _ = v94
                            }
                            _ = v87
                        } else {
                            var v95 uint64 = 0
                            _ = v95
                        }
                        fmt.Println(v1)
                        _ = fi83
                    }
                }
                var v96 int16 = -32768
                _ = v80
                _ = v96
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}
