package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS6 struct { F0 int64; F1 int16; F2 int32 }
type GenS7 struct { F0 int8; F1 int32; F2 uint64 }
func genFn0() int64 {
    var v5 float32 = 0.1
    switch "hello" {
    case "a":
        v5 = v5
    case "foo bar":
        var v6 int64 = -1
        _ = v6
    }
    if ((127 == 0) || ((false && (65535 >= 0)) && (!false))) {
        if false {
            v7 := []int{}
            v5 = -2.25
            if false {
                fmt.Println(v5)
            } else {
                v5 += 0.0
            }
            _ = v7
        } else {
            v5 -= v5
        }
    } else {
        if false {
            for fi8 := 0; fi8 < 12; fi8++ {
                v5 = (v5 - (v5 * v5))
                v5 += float32(3.14159)
                if (-1 >= 127) {
                    var v9 float64 = (0.0 * 0.1)
                    v5 = float32(0.0)
                    v9 += 1.5
                    _ = v9
                } else {
                    var v10 uint = 18446744073709551615
                    _ = v10
                }
                _ = fi8
            }
            var v11 float32 = v5
            if true {
                for fi12 := 0; fi12 < 18; fi12++ {
                    switch ("hello" + (((("hello" + "a") + "foo bar") + "foo bar") + ("hello" + "\n\t"))) {
                    case "hello":
                        v13 := GenS1{F0: ("\n\t" + ""), F1: 4294967295}
                        _ = v13
                    case "":
                        var v14 int16 = 32767
                        _ = v14
                    case "foo bar":
                        var v15 bool = false
                        _ = v15
                    default:
                        fmt.Println(v5, v11, v5)
                    }
                    var v16 uint32 = 395
                    var v17 int64 = -9223372036854775808
                    _ = fi12
                    _ = v16
                    _ = v17
                }
            }
            _ = v11
        } else {
            var v18 int = 32768
            v18 += v18
            _ = v18
        }
    }
    if (32767 <= 32767) {
        v5 = v5
    } else {
        fmt.Println(v5)
    }
    _ = v5
    return 42
}
func genFn1(p0 string, p1 uint64) uint {
    p1 -= p1
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 int64 = 42
            if (2 > -68) {
                if true {
                    v22 := []int{32767, (fi20 | 1), (fi20 + (-(fi19)))}
                    fmt.Println(p1, fi20, p1)
                    var v23 float64 = float64((1.5 + (1.5 - 1e10)))
                    _ = v22
                    _ = v23
                }
            }
            switch p0 {
            case "foo bar":
                v24 := GenS1{F0: p0, F1: p1}
                _ = v24
            case "\n\t":
                if (-32768 > 0) {
                    fmt.Println(p1, p0, p1)
                    switch p0 {
                    case "\n\t":
                        if true {
                            p0 = ((("" + (p0 + "foo bar")) + "foo bar") + "a")
                        }
                    case "":
                        var v25 int64 = (v21 & (v21 + v21))
                        _ = v25
                    }
                    v26 := GenS2{F0: 1}
                    switch (fi20 / (((fi20 / ((9223372036854775807) | 1))) | 1)) {
                    case -2147483648:
                        var v27 string = "hello"
                        _ = v27
                    default:
                        v28 := []int{(^(fi19))}
                        _ = v28
                    }
                    _ = v26
                } else {
                    if false {
                        p0 = p0
                        if ((!(true || (!true))) || true) {
                            p1 = p1
                            v21 += (v21 + 507)
                        } else {
                            var v29 uint16 = 32767
                            var v30 uint8 = 2
                            _ = v29
                            _ = v30
                        }
                        p1 = p1
                    }
                }
            default:
                v31 := []int{(fi19 % ((fi20) | 1)), (fi19 & fi20), (fi19 - fi20)}
                _ = v31
            }
            p0 = ((p0 + (p0 + "a")) + "a")
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    v32 := GenS2{F0: -9223372036854775808}
    v33 := GenS3{F0: 0}
    _ = p0
    _ = p1
    _ = v32
    _ = v33
    return 4294967295
}

func main() {
    v1 := []uint32{122, 414, 1, 975}
    var v2 uint32 = 395
    v3 := []uint32{(v2 * v2), (^(v2)), (v2 | v2), (v2 & (v2 / ((v2) | 1)))}
    v4 := make(map[int]int)
    switch 864 {
    case -1:
        var v34 uint32 = 0
        _ = v34
    default:
        var v35 float32 = 0.0
        _ = v35
    }
    if false {
        v2 = 1
        var v36 uint = 42
        _ = v36
    } else {
        var v37 uint32 = v2
        v2 += 1
        if true {
            _ = len(v4)
            var v38 bool = true
            _ = v38
        } else {
            for fi39 := 0; fi39 < 10; fi39++ {
                var v40 int = (fi39 % ((4294967295) | 1))
                _ = fi39
                _ = v40
            }
            fmt.Println(v37, v37)
            if len(v1) > 0 {
                v1[0] = (v2 >> ((genFn1("\n\t", 42)) & 63))
            }
        }
        _ = v37
    }
    if true {
        for fi41 := 0; fi41 < 6; fi41++ {
            if (-1 > -1) {
                switch (fi41 * 32767) {
                case 4294967295:
                    var v42 uint = genFn1("", 356)
                    _ = v42
                case -9223372036854775808:
                    var v43 int8 = 1
                    _ = v43
                }
                v44 := GenS4{F0: 377}
                _ = v44
            }
            for fi45 := 0; fi45 < 8; fi45++ {
                v2 = v2
                _ = fi45
            }
            if true {
                var v46 bool = (false && true)
                _ = v46
            }
            _ = fi41
        }
    }
    var v47 float64 = -2.25
    if false {
        v47 -= (0.0 + (float64((1e10 + ((3.14159 + -2.25) * float32(1.5)))) - ((0.1 / v47) - -2.25)))
        if false {
            switch "a" {
            case "":
                for fi48 := 0; fi48 < 16; fi48++ {
                    fmt.Println(fi48, v47)
                    var v49 float32 = ((1.5 / -2.25) / (float32(v47) / (((0.0 + float32(float64(float32((float64(-2.25) / v47))))) + ((0.1 * 0.1) * 0.1)) + 0.1)))
                    _ = fi48
                    _ = v49
                }
            case "foo bar":
                v47 = (v47 * v47)
            case "\n\t":
                for fi50 := 0; fi50 < 14; fi50++ {
                    switch "\n\t" {
                    case "hello":
                        switch "hello" {
                        case "foo bar":
                            if (true || (100 != -39)) {
                                v47 += 0.1
                            }
                        case "a":
                            _ = len(v4)
                        case "hello":
                            _ = len(v4)
                        }
                    case "\n\t":
                        var v51 uint16 = 1
                        _ = v51
                    }
                    _ = fi50
                }
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            if false {
                v52 := v4[246]
                if false {
                    var v53 uint8 = 0
                    _ = v53
                } else {
                    v2 = (v2 % ((v2) | 1))
                    var v54 int8 = 1
                    var v55 int8 = -49
                    _ = v54
                    _ = v55
                }
                v3 = append(v3, v2)
                _ = v52
            }
        }
    }
    if ((42 != 235) || true) {
        for fi56 := 0; fi56 < 1; fi56++ {
            fmt.Println(v2, v47)
            _ = fi56
        }
        var v57 uint8 = 255
        _ = v57
    }
    var v58 string = ""
    fmt.Println(v47, v47, v2)
    for fi59 := 0; fi59 < 16; fi59++ {
        if (!true) {
            if len(v1) > 0 {
                v1[0] = v2
            }
        }
        for fi60 := 0; fi60 < 14; fi60++ {
            delete(v4, 32767)
            if false {
                if true {
                    if (195 < 1) {
                        if false {
                            fmt.Println(fi60, fi60)
                            var v61 int32 = 1
                            v62 := GenS1{F0: v58, F1: 1}
                            _ = v61
                            _ = v62
                        } else {
                            if ((3.14159 < (-2.25 / 1.5)) && (!(fi59 == fi60))) {
                                var v63 uint32 = v2
                                v58 = v58
                                v64 := GenS5{F0: v2, F1: 984, F2: 18446744073709551615}
                                _ = v63
                                _ = v64
                            }
                            v65 := v4[(fi60 >> ((1) & 63))]
                            v66 := GenS6{F0: genFn0(), F1: 1000, F2: -2147483648}
                            _ = v65
                            _ = v66
                        }
                        v58 = (v58 + ("a" + "\n\t"))
                    }
                } else {
                    v58 = ((v58 + v58) + "\n\t")
                    if false {
                        if (true || false) {
                            if ((float64(3.14159) < (v47 - (v47 / (1e10 - 3.14159)))) && false) {
                                fmt.Println(fi60, fi60, fi60)
                            }
                            fmt.Println(fi59, fi59, fi60)
                            switch (v58 + v58) {
                            case "a":
                                if (!(245 >= 112)) {
                                    v4[fi60] = (fi60 % ((fi60) | 1))
                                    if true {
                                        var v67 uint16 = 65535
                                        _ = v67
                                    } else {
                                        v68 := GenS7{F0: 127, F1: 0, F2: 924}
                                        var v69 int = 714
                                        if false {
                                            v70 := GenS1{F0: ((v58 + "") + v58), F1: 623}
                                            if true {
                                                if len(v1) > 0 {
                                                    v1[0] = 2
                                                }
                                            }
                                            var v71 uint32 = (v2 / ((v2) | 1))
                                            var v72 bool = true
                                            _ = v70
                                            _ = v71
                                            _ = v72
                                        }
                                        _ = v68
                                        _ = v69
                                    }
                                }
                            default:
                                switch (v58 + (v58 + v58)) {
                                case "foo bar":
                                    var v73 int8 = -94
                                    _ = v73
                                }
                            }
                        }
                        v47 = 0.0
                    }
                }
                switch -840 {
                case 1:
                    v2 -= (v2 % ((2147483647) | 1))
                default:
                    var v74 float32 = 1.5
                    _ = v74
                }
                v47 = float64(-2.25)
            } else {
                v2 += 4294967295
                var v75 bool = true
                _ = v75
            }
            fmt.Println(fi59)
            _ = fi60
        }
        _ = fi59
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v58
}
