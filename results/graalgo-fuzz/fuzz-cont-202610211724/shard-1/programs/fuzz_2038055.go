package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint; F1 int16; F2 int64 }
type GenS4 struct { F0 int64; F1 string; F2 uint }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS7 struct { F0 int64; F1 int16; F2 uint }
type GenS8 struct { F0 uint32 }
func genFn0(p0 int8, p1 int64, p2 uint32) int32 {
    v5 := []int{127, 42}
    var v6 int32 = -252
    var v7 int = 32768
    switch "hello" {
    case "":
        fmt.Println(p1, v6)
    case "\n\t":
        fmt.Println(p2, v6, v6)
    case "foo bar":
        for fi8 := 0; fi8 < 17; fi8++ {
            fmt.Println(p0)
            switch (v7 * (fi8 | 350)) {
            case 32767:
                if len(v5) > 0 {
                    v5[0] = fi8
                }
            case 727:
                for fi9 := 0; fi9 < 10; fi9++ {
                    var v10 float64 = (1.5 * (3.14159 * float64(1.5)))
                    _ = fi9
                    _ = v10
                }
            case 42:
                for fi11 := 0; fi11 < 2; fi11++ {
                    var v12 uint = 255
                    if false {
                        v13 := GenS1{F0: p1}
                        _ = v13
                    } else {
                        if (!((1e10 + float32((3.14159 - ((float64(-2.25) / 1.0) + -2.25)))) > 0.1)) {
                            v12 -= (v12 % (((v12 ^ 4)) | 1))
                        } else {
                            fmt.Println(p1, p1, p0)
                            var v14 int = fi11
                            _ = v14
                        }
                        if len(v5) > 0 {
                            v5[0] = (fi11 * (fi11 | fi11))
                        }
                    }
                    var v15 uint16 = 637
                    v16 := GenS1{F0: p1}
                    var v17 int8 = p0
                    var v18 int16 = 42
                    fmt.Println(p1, p1)
                    _ = fi11
                    _ = v12
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v18
                }
            }
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (-(v6))
}
func genFn1(p0 string, p1 int64, p2 int32) uint16 {
    var v19 int32 = -1
    var v20 float64 = 0.0
    switch -1 {
    case 128:
        v21 := GenS2{F0: 65535, F1: (p1 << ((2) & 63))}
        _ = v21
    default:
        var v22 uint8 = 1
        _ = v22
    }
    switch (("" + "") + "\n\t") {
    case "":
        switch 2147483647 {
        case 2147483647:
            if true {
                v23 := GenS2{F0: 243, F1: p1}
                v24 := []int{0}
                _ = v23
                _ = v24
            } else {
                v25 := GenS2{F0: 570, F1: -343}
                for fi26 := 0; fi26 < 7; fi26++ {
                    v27 := GenS1{F0: p1}
                    v28 := GenS1{F0: (p1 & (p1 | p1))}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = v25
            }
        default:
            for fi29 := 0; fi29 < 13; fi29++ {
                for fi30 := 0; fi30 < 1; fi30++ {
                    fmt.Println(fi30)
                    v31 := GenS3{F0: 0, F1: 0, F2: -1}
                    fmt.Println(v20, fi30, v19)
                    v32 := []int{4294967295, (fi29 ^ 2147483647), (-(fi30))}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                var v33 uint32 = 65535
                _ = fi29
                _ = v33
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return 65535
}

func main() {
    var v1 bool = true
    var v2 int = 9223372036854775807
    var v3 int16 = -973
    v4 := make(map[uint32]uint32)
    switch v2 {
    case -750:
        if false {
            var v34 int16 = 0
            _ = v34
        } else {
            for fi35 := 0; fi35 < 16; fi35++ {
                v2 -= (fi35 + (fi35 - 1))
                _ = fi35
            }
            var v36 float32 = 0.0
            _ = v36
        }
    case 83:
        var v37 int = (v2 / ((v2) | 1))
        _ = v37
    case 1:
        v2 = v2
    }
    v2 = 32768
    v3 -= v3
    v38 := GenS4{F0: 9223372036854775807, F1: "\n\t", F2: 255}
    if v1 {
        if v1 {
            var v39 float32 = 1e10
            _ = v39
        }
    } else {
        switch 1 {
        case -2147483648:
            var v40 uint32 = 65535
            _ = v40
        }
        var v41 uint = 2
        for fi42 := 0; fi42 < 17; fi42++ {
            v43 := GenS2{F0: 0, F1: -9223372036854775808}
            var v44 string = "foo bar"
            for fi45 := 0; fi45 < 5; fi45++ {
                v4[uint32(v41)] = 2
                var v46 int64 = 104
                v4[1] = 0
                v47 := GenS5{F0: 2, F1: genFn1(v44, -365, 42)}
                _ = fi45
                _ = v46
                _ = v47
            }
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = v41
    }
    switch ("\n\t" + ("" + "foo bar")) {
    case "a":
        var v48 string = "hello"
        _ = v48
    case "\n\t":
        switch "hello" {
        case "foo bar":
            v2 = v2
        case "a":
            if (((1e10 / 3.14159) == 0.1) || false) {
                v49 := GenS6{F0: 42, F1: (v3 >> ((533) & 63)), F2: 570}
                _ = v49
            }
        default:
            if v1 {
                if (false && v1) {
                    v2 = v2
                }
                var v50 int8 = 1
                if ((1.5 == float64((((3.14159 * 1.5) / (3.14159 + float32(3.14159))) + 1.5))) && (float64(1e10) < 0.1)) {
                    if (-437 == 0) {
                        if (3.14159 != (0.0 / ((-2.25 * 0.1) + 0.1))) {
                            v51 := []int{v2, v2}
                            _ = v51
                        } else {
                            if (-507 == v2) {
                                v3 += -32768
                                switch "a" {
                                case "\n\t":
                                    v52 := GenS7{F0: 898, F1: v3, F2: 255}
                                    _ = v52
                                case "foo bar":
                                    for fi53 := 0; fi53 < 18; fi53++ {
                                        v1 = (v1 || false)
                                        _ = fi53
                                    }
                                default:
                                    v1 = (!(true || true))
                                }
                                v2 = (^(v2))
                            }
                        }
                        var v54 uint64 = 2
                        if v1 {
                            delete(v4, 675)
                        }
                        _ = v54
                    } else {
                        switch (-(v2)) {
                        case 4294967295:
                            for fi55 := 0; fi55 < 2; fi55++ {
                                for fi56 := 0; fi56 < 19; fi56++ {
                                    v3 += v3
                                    v2 += (v2 | v2)
                                    _ = fi56
                                }
                                _ = fi55
                            }
                        default:
                            if v1 {
                                for fi57 := 0; fi57 < 11; fi57++ {
                                    delete(v4, 1)
                                    var v58 uint32 = 4294967295
                                    _ = fi57
                                    _ = v58
                                }
                            } else {
                                if ((128 > 8) && false) {
                                    fmt.Println(v1)
                                    for fi59 := 0; fi59 < 1; fi59++ {
                                        var v60 float32 = (-2.25 * (float32(1.5) * 0.0))
                                        v60 -= ((-2.25 + v60) / v60)
                                        _ = fi59
                                        _ = v60
                                    }
                                    v2 = v2
                                } else {
                                    v1 = v1
                                    v2 = (v2 ^ v2)
                                }
                                var v61 float32 = 0.0
                                _ = v61
                            }
                        }
                    }
                    if v1 {
                        for fi62 := 0; fi62 < 5; fi62++ {
                            v50 = v50
                            _ = fi62
                        }
                        if v1 {
                            v3 -= -1
                            fmt.Println(v2, v3, v1)
                            v63 := GenS1{F0: -33}
                            _ = v63
                        }
                    } else {
                        v64 := GenS1{F0: 642}
                        _ = v64
                    }
                    var v65 uint32 = 42
                    _ = v65
                }
                _ = v50
            } else {
                var v66 uint64 = 1
                delete(v4, 42)
                var v67 bool = ((-31 <= 60) || v1)
                _ = v66
                _ = v67
            }
        }
    default:
        for fi68 := 0; fi68 < 3; fi68++ {
            v69 := v4[2147483648]
            v2 = fi68
            _ = fi68
            _ = v69
        }
    }
    if (!true) {
        var v70 uint8 = 42
        var v71 uint = 4294967295
        v1 = v1
        _ = v70
        _ = v71
    } else {
        var v72 int16 = (v3 - v3)
        _ = v72
    }
    v2 += v2
    v73 := GenS8{F0: 2147483647}
    var v74 uint64 = 593
    for fi75 := 0; fi75 < 9; fi75++ {
        if (!(false && ((!v1) && v1))) {
            for fi76 := 0; fi76 < 5; fi76++ {
                v3 -= (^(v3))
                v77 := []int{fi76, (v2 >> ((757) & 63)), 248}
                var v78 float32 = (-2.25 * 0.0)
                _ = fi76
                _ = v77
                _ = v78
            }
            if (255 >= 83) {
                var v79 float32 = 1.5
                var v80 int64 = -1
                for fi81 := 0; fi81 < 17; fi81++ {
                    var v82 float32 = 0.0
                    var v83 int32 = genFn0(-1, v80, 65535)
                    var v84 float64 = 0.0
                    var v85 int16 = v3
                    var v86 int16 = 414
                    _ = fi81
                    _ = v82
                    _ = v83
                    _ = v84
                    _ = v85
                    _ = v86
                }
                _ = v79
                _ = v80
            }
        }
        var v87 uint64 = v74
        _ = fi75
        _ = v87
    }
    fmt.Println(v2)
    switch ("foo bar" + "a") {
    case "\n\t":
        for fi88 := 0; fi88 < 17; fi88++ {
            switch (v2 | 1) {
            case 0:
                v89 := []int{fi88}
                _ = v89
            }
            v90 := GenS4{F0: -642, F1: "a", F2: 334}
            v74 -= v74
            _ = fi88
            _ = v90
        }
    case "foo bar":
        v91 := GenS7{F0: 42, F1: (v3 | (v3 ^ (v3 | v3))), F2: 721}
        _ = v91
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v73
    _ = v74
}
