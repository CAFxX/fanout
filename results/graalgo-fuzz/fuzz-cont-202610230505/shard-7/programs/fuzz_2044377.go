package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint32; F1 int16; F2 string }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string; F1 int }
type GenS9 struct { F0 uint32; F1 int8 }
type GenS10 struct { F0 uint64 }
func genFn0(p0 int32, p1 bool, p2 int) int8 {
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            fmt.Println(fi6, p1, fi5)
            v7 := []int{}
            p1 = true
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    v8 := []int{p2, (p2 * (p2 << ((1) & 63)))}
    var v9 int8 = -121
    v10 := GenS1{F0: 65535, F1: p2}
    var v11 uint16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return (v9 ^ v9)
}
func genFn1(p0 uint8, p1 uint) uint {
    p1 = p1
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint = (p1 << (((p1 ^ (p1 - (p1 | (p1 - (p1 & 18446744073709551615)))))) & 63))
        fmt.Println(p0, v13, v13)
        _ = fi12
        _ = v13
    }
    if false {
        switch "\n\t" {
        case "":
            fmt.Println(p1, p1)
        case "\n\t":
            v14 := []int{}
            _ = v14
        case "foo bar":
            if true {
                v15 := []int{1, -1, 295, -328}
                _ = v15
            }
        }
        p1 = p1
    } else {
        p1 -= p1
        if (!false) {
            v16 := []int{9223372036854775807, -9223372036854775808, -1, 32768}
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            if true {
                v17 := GenS1{F0: 4294967295, F1: -231}
                v18 := GenS2{F0: 826, F1: 9223372036854775807}
                if true {
                    for fi19 := 0; fi19 < 15; fi19++ {
                        fmt.Println(fi19, p1, p1)
                        switch (fi19 / (((fi19 >> (((p1 * (p1 | p1))) & 63))) | 1)) {
                        case -751:
                            p0 += 255
                        case -761:
                            p1 -= p1
                        }
                        _ = fi19
                    }
                    for fi20 := 0; fi20 < 6; fi20++ {
                        fmt.Println(fi20, fi20, p1)
                        _ = fi20
                    }
                } else {
                    v21 := GenS1{F0: 65535, F1: 53}
                    _ = v21
                }
                _ = v17
                _ = v18
            }
            _ = v16
        } else {
            p1 = 255
        }
        v22 := []int{-1}
        _ = v22
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        if (!(!false)) {
            v24 := GenS3{F0: p1}
            var v25 int32 = -1
            _ = v24
            _ = v25
        } else {
            switch 0 {
            case -339:
                if true {
                    p1 = (p1 - 255)
                    for fi26 := 0; fi26 < 11; fi26++ {
                        var v27 uint = p1
                        v28 := GenS2{F0: 787, F1: 0}
                        v29 := []int{4294967295, fi26, (fi23 / ((128) | 1)), fi23}
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    var v30 bool = (!true)
                    _ = v30
                } else {
                    if true {
                        p0 = p0
                        var v31 uint8 = 128
                        if ((^(v31)) != 42) {
                            var v32 int8 = 42
                            var v33 int = -9223372036854775808
                            if false {
                                var v34 int8 = 42
                                for fi35 := 0; fi35 < 5; fi35++ {
                                    v36 := GenS4{F0: (p1 * p1), F1: 9223372036854775807, F2: 0}
                                    _ = fi35
                                    _ = v36
                                }
                                _ = v34
                            }
                            _ = v32
                            _ = v33
                        }
                        _ = v31
                    } else {
                        if true {
                            v37 := []int{fi23, (-(fi23))}
                            _ = v37
                        }
                        p0 = 214
                    }
                    var v38 float64 = 1.5
                    _ = v38
                }
            case -9223372036854775808:
                for fi39 := 0; fi39 < 17; fi39++ {
                    fmt.Println(p1, fi23)
                    _ = fi39
                }
            }
            v40 := GenS4{F0: (^(p1)), F1: 9223372036854775807, F2: 983}
            v41 := GenS5{F0: 2147483647, F1: int16(fi23), F2: "a"}
            _ = v40
            _ = v41
        }
        var v42 float32 = 3.14159
        _ = fi23
        _ = v42
    }
    v43 := []int{853, 1, -1}
    _ = p0
    _ = p1
    _ = v43
    return p1
}

func main() {
    v1 := make(map[int]int)
    var v2 float32 = (((3.14159 / float32((float64(0.1) / float64(-2.25)))) * 3.14159) - 1e10)
    var v3 int8 = 107
    var v4 string = "foo bar"
    var v44 float64 = 0.0
    v45 := GenS6{F0: v4, F1: 42}
    if false {
        v3 -= genFn0(650, (true && (-9223372036854775808 == 42)), 548)
        var v46 uint = genFn1(0, genFn1(1, 65535))
        _ = v46
    }
    if true {
        v44 = v44
        v47 := []int{993}
        if (((float64((0.1 / v2)) + v44) * 0.1) == float64(v2)) {
            var v48 int = 42
            v2 = v2
            var v49 int32 = 2147483647
            _ = v48
            _ = v49
        } else {
            switch 543 {
            case -548:
                switch -1 {
                case -9223372036854775808:
                    var v50 int32 = 50
                    _ = v50
                default:
                    if ((18446744073709551615 < 341) || (4294967295 > 403)) {
                        if (557 != 674) {
                            v3 -= 127
                            v51 := GenS4{F0: genFn1(2, 914), F1: -304, F2: 465}
                            _ = v51
                        } else {
                            v52 := GenS7{F0: 0}
                            _ = v52
                        }
                    } else {
                        v53 := GenS5{F0: 65535, F1: -1, F2: v4}
                        v54 := GenS8{F0: v4, F1: 2147483647}
                        _ = v53
                        _ = v54
                    }
                }
            case -959:
                if len(v47) > 0 {
                    v47[0] = -9223372036854775808
                }
            default:
                for fi55 := 0; fi55 < 1; fi55++ {
                    if (42 != -308) {
                        if false {
                            v56 := v1[(fi55 >> ((42) & 63))]
                            for fi57 := 0; fi57 < 20; fi57++ {
                                if (false && (!(v3 > (v3 & (v3 % (((v3 ^ -70)) | 1)))))) {
                                    if true {
                                        fmt.Println(fi55, v2)
                                    } else {
                                        switch "a" {
                                        case "foo bar":
                                            var v58 int64 = 1
                                            _ = v58
                                        case "a":
                                            var v59 int16 = -32768
                                            _ = v59
                                        case "hello":
                                            if (("foo bar" > ("foo bar" + "foo bar")) || true) {
                                                if false {
                                                    fmt.Println(v3, v4, fi57)
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    if (!false) {
                                        if false {
                                            v56 += v56
                                            v60 := GenS9{F0: 0, F1: 1}
                                            _ = v60
                                        } else {
                                            var v61 float32 = v2
                                            switch (("\n\t" + ("" + v4)) + "foo bar") {
                                            case "\n\t":
                                                var v62 int8 = -128
                                                _ = v62
                                            default:
                                                v2 = 1.5
                                                if (-9223372036854775808 > -977) {
                                                    var v63 int64 = -1
                                                    v61 += v2
                                                    fmt.Println(v2, v63)
                                                    _ = v63
                                                } else {
                                                    switch v4 {
                                                    case "foo bar":
                                                        v3 = (v3 << ((42) & 63))
                                                    case "a":
                                                        switch v4 {
                                                        case "":
                                                            switch (fi57 % ((v56) | 1)) {
                                                            case 2147483647:
                                                                fmt.Println(v4)
                                                            case 4294967295:
                                                                v3 -= v3
                                                            case 778:
                                                                v4 = (v4 + v4)
                                                            default:
                                                                v61 = (-2.25 * 0.0)
                                                                delete(v1, fi57)
                                                                v64 := GenS9{F0: 2147483647, F1: (v3 >> ((1) & 63))}
                                                                _ = v64
                                                            }
                                                        case "\n\t":
                                                            var v65 uint64 = 42
                                                            _ = v65
                                                        default:
                                                            v66 := GenS8{F0: ((v4 + (v4 + v4)) + ""), F1: (fi57 / ((v56) | 1))}
                                                            _ = v66
                                                        }
                                                    case "":
                                                        var v67 bool = false
                                                        _ = v67
                                                    default:
                                                        var v68 uint8 = 0
                                                        _ = v68
                                                    }
                                                    var v69 uint = 780
                                                    var v70 bool = (!(true && true))
                                                    _ = v69
                                                    _ = v70
                                                }
                                                v71 := GenS8{F0: v4, F1: (v56 * (fi57 ^ (v56 - (v56 >> ((genFn1(42, genFn1(128, genFn1(252, 42)))) & 63)))))}
                                                _ = v71
                                            }
                                            _ = v61
                                        }
                                        fmt.Println(fi55)
                                        v72 := GenS1{F0: 2, F1: 127}
                                        _ = v72
                                    }
                                    v73 := GenS6{F0: "foo bar", F1: -1}
                                    _ = v73
                                }
                                var v74 int64 = -1
                                v74 += (v74 - (v74 - 167))
                                v75 := GenS6{F0: (v4 + v4), F1: (v74 ^ (v74 << ((262) & 63)))}
                                v47 = append(v47, fi55)
                                _ = fi57
                                _ = v74
                                _ = v75
                            }
                            _ = v56
                        }
                        var v76 uint64 = 4294967295
                        _ = v76
                    }
                    for fi77 := 0; fi77 < 15; fi77++ {
                        if true {
                            v78 := GenS10{F0: 1}
                            _ = v78
                        }
                        var v79 float64 = float64((-2.25 / (3.14159 + v2)))
                        _ = fi77
                        _ = v79
                    }
                    v4 = (("\n\t" + v4) + v4)
                    _ = fi55
                }
            }
            v3 += (-(v3))
        }
        _ = v47
    } else {
        fmt.Println(v3, v3, v2)
        for fi80 := 0; fi80 < 5; fi80++ {
            switch (fi80 * (fi80 | fi80)) {
            case 128:
                fmt.Println(fi80, v44)
            case 4294967295:
                switch (((((v4 + (("\n\t" + v4) + "")) + "") + "a") + ((v4 + "a") + (v4 + v4))) + "foo bar") {
                case "\n\t":
                    fmt.Println(fi80, v4, v4)
                case "a":
                    var v81 float64 = (3.14159 / v44)
                    _ = v81
                case "hello":
                    v82 := []int{-1, 42, 278, (fi80 * fi80)}
                    _ = v82
                default:
                    v83 := []int{fi80, (fi80 * 614), fi80}
                    _ = v83
                }
            case -2147483648:
                var v84 uint16 = 42
                _ = v84
            default:
                v44 = ((v44 + v44) / 3.14159)
            }
            _ = fi80
        }
        var v85 float32 = float32(-2.25)
        _ = v85
    }
    for fi86 := 0; fi86 < 16; fi86++ {
        v3 -= (v3 % ((genFn0(-2147483648, ((fi86 ^ fi86) != (fi86 + (-(fi86)))), fi86)) | 1))
        var v87 int64 = 1
        _ = fi86
        _ = v87
    }
    v4 = (v4 + (v4 + "hello"))
    v3 -= -1
    var v88 float64 = 3.14159
    var v89 uint64 = 10
    var v90 int8 = (v3 >> ((2) & 63))
    fmt.Println(v4, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v88
    _ = v89
    _ = v90
}
