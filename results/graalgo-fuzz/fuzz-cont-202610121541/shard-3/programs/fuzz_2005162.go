package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS6 struct { F0 int64; F1 uint32 }
func genFn0(p0 string, p1 int) uint16 {
    var v4 bool = true
    v4 = ((p0 < "\n\t") && (-1 < 2147483647))
    for fi5 := 0; fi5 < 20; fi5++ {
        switch fi5 {
        case 128:
            if v4 {
                v4 = true
            } else {
                for fi6 := 0; fi6 < 15; fi6++ {
                    var v7 int32 = 42
                    _ = fi6
                    _ = v7
                }
                for fi8 := 0; fi8 < 5; fi8++ {
                    p1 = fi8
                    if (2147483647 < 42) {
                        fmt.Println(fi8)
                        v9 := []int{(fi8 >> ((882) & 63)), (fi5 % (((fi5 / ((-803) | 1))) | 1))}
                        var v10 uint = 775
                        if (p0 <= (("\n\t" + p0) + "a")) {
                            if v4 {
                                v4 = false
                                v11 := GenS1{F0: false}
                                var v12 int64 = 0
                                _ = v11
                                _ = v12
                            } else {
                                var v13 int = 128
                                var v14 bool = (true && v4)
                                v4 = (65535 >= 2)
                                _ = v13
                                _ = v14
                            }
                            v15 := GenS1{F0: (383 != 42)}
                            _ = len(v9)
                            var v16 bool = (v4 || v4)
                            _ = v15
                            _ = v16
                        }
                        if (v4 || false) {
                            var v17 int64 = 951
                            switch "hello" {
                            case "hello":
                                switch p0 {
                                case "a":
                                    var v18 int32 = 1
                                    _ = v18
                                case "\n\t":
                                    if (((("foo bar" + "\n\t") + p0) == p0) || (32767 >= 814)) {
                                        v10 = v10
                                        if len(v9) > 0 {
                                            v9[0] = (fi5 % ((-9223372036854775808) | 1))
                                        }
                                        var v19 int32 = 2147483647
                                        _ = v19
                                    }
                                default:
                                    v20 := GenS2{F0: (v17 & v17), F1: 42}
                                    var v21 int16 = 1000
                                    v10 = v10
                                    _ = v20
                                    _ = v21
                                }
                                var v22 int8 = 12
                                if ((v22 + 100) >= v22) {
                                    if (v4 && v4) {
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                        if true {
                                            var v23 float64 = (((0.1 - (float64(0.1) - -2.25)) * 0.1) * 0.1)
                                            var v24 uint64 = 1
                                            _ = v23
                                            _ = v24
                                        }
                                        var v25 float64 = (1e10 * 1.5)
                                        _ = v25
                                    }
                                    switch (fi8 << ((v10) & 63)) {
                                    case 0:
                                        _ = len(v9)
                                    case 496:
                                        switch p0 {
                                        case "":
                                            v26 := GenS1{F0: (false || v4)}
                                            _ = v26
                                        default:
                                            var v27 bool = v4
                                            _ = v27
                                        }
                                    case -1:
                                        var v28 float32 = float32(1.5)
                                        _ = v28
                                    }
                                }
                                if ((!true) && (-1 <= 1000)) {
                                    v22 += v22
                                    if len(v9) > 0 {
                                        v9[0] = 870
                                    }
                                }
                                _ = v22
                            }
                            fmt.Println(p0)
                            v9 = append(v9, (fi5 + -9223372036854775808))
                            v10 -= v10
                            var v29 bool = (((v4 && ((1.5 + (3.14159 - 0.0)) > 0.0)) && true) || (253 <= 32767))
                            p0 = ""
                            _ = v17
                            _ = v29
                        }
                        _ = v10
                        _ = v9
                    }
                    _ = fi8
                }
                if false {
                    v4 = (false || v4)
                } else {
                    switch p0 {
                    case "a":
                        v30 := GenS2{F0: 0, F1: 255}
                        _ = v30
                    case "foo bar":
                        switch p0 {
                        case "\n\t":
                            for fi31 := 0; fi31 < 16; fi31++ {
                                if v4 {
                                    v4 = true
                                    v32 := GenS2{F0: -281, F1: 65535}
                                    _ = v32
                                } else {
                                    switch ((p0 + "foo bar") + p0) {
                                    case "foo bar":
                                        var v33 uint16 = 255
                                        _ = v33
                                    default:
                                        p0 = p0
                                    }
                                    var v34 int32 = 42
                                    v4 = (491 <= 42)
                                    _ = v34
                                }
                                _ = fi31
                            }
                        default:
                            v35 := GenS1{F0: ((!v4) || (1.5 == (0.1 * (0.0 - 1e10))))}
                            _ = v35
                        }
                    default:
                        v36 := GenS1{F0: v4}
                        _ = v36
                    }
                    if true {
                        switch p1 {
                        case -1:
                            v4 = (1 < 48)
                        case 9223372036854775807:
                            p1 += 694
                        default:
                            var v37 uint = 384
                            _ = v37
                        }
                    }
                    fmt.Println(p0, p0, p1)
                }
            }
        case 1:
            fmt.Println(v4, v4)
        }
        _ = fi5
    }
    var v38 float64 = 1e10
    var v39 uint16 = 1
    _ = p0
    _ = p1
    _ = v38
    _ = v39
    _ = v4
    return (v39 + v39)
}
func genFn1(p0 uint32) int32 {
    switch "a" {
    case "foo bar":
        p0 -= p0
    case "":
        p0 = 1
    }
    if false {
        p0 += 2
        if true {
            var v40 int = 9223372036854775807
            if (1 > -128) {
                switch v40 {
                case 0:
                    fmt.Println(v40, p0)
                case 1:
                    p0 -= p0
                default:
                    v41 := GenS2{F0: 0, F1: 32767}
                    _ = v41
                }
                v40 = v40
                v42 := GenS3{F0: 824, F1: 109}
                _ = v42
            } else {
                for fi43 := 0; fi43 < 14; fi43++ {
                    if (-1 > -1) {
                        for fi44 := 0; fi44 < 4; fi44++ {
                            switch "hello" {
                            case "a":
                                fmt.Println(fi44, v40)
                            case "foo bar":
                                var v45 int64 = -9223372036854775808
                                _ = v45
                            default:
                                p0 = 4294967295
                            }
                            p0 = (p0 << ((4294967295) & 63))
                            var v46 uint = 42
                            fmt.Println(v46, fi44)
                            _ = fi44
                            _ = v46
                        }
                        v47 := GenS1{F0: true}
                        var v48 int8 = -128
                        _ = v47
                        _ = v48
                    }
                    var v49 int32 = -1
                    _ = fi43
                    _ = v49
                }
                if (177 == 128) {
                    for fi50 := 0; fi50 < 2; fi50++ {
                        var v51 uint16 = 0
                        for fi52 := 0; fi52 < 13; fi52++ {
                            var v53 uint32 = (p0 * p0)
                            v53 = (^(p0))
                            v54 := []int{v40, (fi50 << ((985) & 63)), fi52}
                            _ = fi52
                            _ = v53
                            _ = v54
                        }
                        v51 += 65535
                        _ = fi50
                        _ = v51
                    }
                    var v55 bool = (1.5 <= 3.14159)
                    _ = v55
                }
            }
            _ = v40
        } else {
            switch "a" {
            case "a":
                switch 4294967295 {
                case -2147483648:
                    switch ("a" + "foo bar") {
                    case "a":
                        p0 += 65535
                    case "foo bar":
                        if false {
                            var v56 bool = ((2147483647 > 103) || true)
                            if (!(2 != 42)) {
                                if (!v56) {
                                    v57 := GenS4{F0: -1}
                                    if (true && (1 != 1)) {
                                        var v58 uint8 = 42
                                        _ = v58
                                    }
                                    v56 = v56
                                    _ = v57
                                } else {
                                    var v59 uint = 0
                                    var v60 uint16 = 833
                                    switch 2147483647 {
                                    case 2147483647:
                                        var v61 uint = v59
                                        _ = v61
                                    }
                                    _ = v59
                                    _ = v60
                                }
                            } else {
                                v62 := []int{42, -9223372036854775808, -9223372036854775808}
                                v56 = (false || v56)
                                _ = v62
                            }
                            var v63 uint16 = 2
                            _ = v56
                            _ = v63
                        } else {
                            fmt.Println(p0)
                        }
                    case "hello":
                        fmt.Println(p0)
                    default:
                        var v64 uint = 18446744073709551615
                        _ = v64
                    }
                }
            case "":
                for fi65 := 0; fi65 < 1; fi65++ {
                    switch "foo bar" {
                    case "":
                        var v66 float32 = 3.14159
                        _ = v66
                    case "a":
                        v67 := GenS2{F0: 42, F1: genFn0("foo bar", fi65)}
                        _ = v67
                    case "hello":
                        fmt.Println(p0)
                    }
                    var v68 uint16 = 954
                    v69 := GenS3{F0: 4294967295, F1: -32768}
                    _ = fi65
                    _ = v68
                    _ = v69
                }
            case "hello":
                var v70 int = 185
                _ = v70
            }
            if (127 >= -9223372036854775808) {
                switch 4294967295 {
                case 32768:
                    for fi71 := 0; fi71 < 11; fi71++ {
                        switch "" {
                        case "a":
                            var v72 int8 = 127
                            _ = v72
                        }
                        var v73 int64 = -9223372036854775808
                        v74 := GenS3{F0: uint64(p0), F1: -363}
                        _ = fi71
                        _ = v73
                        _ = v74
                    }
                case 32767:
                    if false {
                        for fi75 := 0; fi75 < 15; fi75++ {
                            var v76 uint = 585
                            for fi77 := 0; fi77 < 2; fi77++ {
                                if (false || (!(206 < 23))) {
                                    fmt.Println(fi75)
                                }
                                v78 := GenS5{F0: 541, F1: 1, F2: 2147483647}
                                _ = fi77
                                _ = v78
                            }
                            v79 := []int{(^(fi75)), (fi75 - fi75), (fi75 / ((fi75) | 1))}
                            _ = fi75
                            _ = v76
                            _ = v79
                        }
                        fmt.Println(p0)
                    } else {
                        p0 = (p0 - (p0 % (((p0 ^ (p0 - (p0 & p0)))) | 1)))
                    }
                default:
                    p0 -= (p0 & p0)
                }
                for fi80 := 0; fi80 < 11; fi80++ {
                    var v81 float64 = 1e10
                    fmt.Println(fi80)
                    var v82 uint64 = 928
                    _ = fi80
                    _ = v81
                    _ = v82
                }
            } else {
                v83 := []int{42}
                if len(v83) > 0 {
                    v83[0] = -9223372036854775808
                }
                if len(v83) >= 2 {
                    v83 = v83[1:len(v83)]
                }
                _ = v83
            }
        }
        v84 := GenS2{F0: -100, F1: 42}
        _ = v84
    }
    var v85 int64 = 9223372036854775807
    _ = p0
    _ = v85
    return -503
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = -1
    var v3 float32 = 0.1
    var v86 uint32 = 622
    v2 += (-(v2))
    v3 -= 3.14159
    var v87 int16 = (-(v2))
    var v88 int8 = -1
    v89 := GenS6{F0: -1, F1: 2147483648}
    v86 -= (v86 << ((18446744073709551615) & 63))
    v90 := []int{}
    var v91 float32 = float32((0.0 - (((float64((v3 / 1.0)) * (1.5 * -2.25)) * (((((3.14159 - 3.14159) + float64(v3)) - (3.14159 + -2.25)) - -2.25) / 1.0)) + float64(0.1))))
    v1[0] = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("slicelen v90:", len(v90), v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v86
    _ = v87
    _ = v88
    _ = v89
    _ = v90
    _ = v91
}
