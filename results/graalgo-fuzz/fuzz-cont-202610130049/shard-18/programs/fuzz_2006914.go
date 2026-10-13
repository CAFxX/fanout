package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS3 struct { F0 int64; F1 int; F2 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 bool; F2 bool }
type GenS7 struct { F0 int; F1 int8 }
type GenS8 struct { F0 uint; F1 string; F2 string }
func genFn0(p0 string, p1 bool, p2 string) int64 {
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            fmt.Println(fi8, p2, fi7)
            switch (fi7 / ((fi7) | 1)) {
            case 128:
                p2 = (p0 + (p2 + p2))
            case 972:
                v9 := GenS1{F0: fi8}
                _ = v9
            }
            _ = fi8
        }
        p2 = "foo bar"
        v10 := []int{128, (fi7 >> ((885) & 63))}
        _ = fi7
        _ = v10
    }
    switch 32767 {
    case -2147483648:
        p2 = ("" + "")
    case 2147483647:
        switch 1 {
        case 4294967295:
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 float64 = 0.0
                fmt.Println(v12, p1, p2)
                p2 = (p2 + "hello")
                _ = fi11
                _ = v12
            }
        default:
            v13 := []int{}
            _ = v13
        }
    case 650:
        var v14 float32 = ((0.0 * 0.1) + 3.14159)
        _ = v14
    default:
        var v15 string = p0
        _ = v15
    }
    p2 = "foo bar"
    v16 := []int{}
    _ = len(v16)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return -1
}
func genFn1(p0 int8, p1 bool, p2 int64) uint {
    if (p0 > 0) {
        for fi17 := 0; fi17 < 17; fi17++ {
            fmt.Println(p0, p1)
            var v18 bool = p1
            if false {
                if (true && (1000 < 0)) {
                    p2 = p2
                    switch (fi17 & (^(fi17))) {
                    case -804:
                        for fi19 := 0; fi19 < 17; fi19++ {
                            var v20 uint16 = 194
                            switch "hello" {
                            case "foo bar":
                                var v21 uint16 = (v20 * (^(v20)))
                                _ = v21
                            }
                            _ = fi19
                            _ = v20
                        }
                    default:
                        if ((0.1 * 1e10) == 0.0) {
                            fmt.Println(p2)
                        }
                    }
                    v22 := []int{fi17, fi17, fi17}
                    _ = v22
                } else {
                    if p1 {
                        for fi23 := 0; fi23 < 18; fi23++ {
                            p0 = p0
                            _ = fi23
                        }
                        p2 += p2
                    } else {
                        if (-2147483648 < 2147483647) {
                            var v24 int = (fi17 | 0)
                            _ = v24
                        } else {
                            var v25 bool = (1.5 < -2.25)
                            fmt.Println(fi17, v25, p0)
                            _ = v25
                        }
                    }
                }
                for fi26 := 0; fi26 < 1; fi26++ {
                    switch fi26 {
                    case 196:
                        v27 := []int{fi17, -637}
                        _ = v27
                    case 32767:
                        if false {
                            if (p0 <= 127) {
                                v28 := []int{(fi26 | fi26), (^(fi26)), 42, 127}
                                _ = v28
                            } else {
                                var v29 float64 = (1.5 - 0.1)
                                v30 := GenS2{F0: 31, F1: 78, F2: 1}
                                v18 = ((int(p2) <= fi26) && false)
                                _ = v29
                                _ = v30
                            }
                        }
                    }
                    _ = fi26
                }
            } else {
                fmt.Println(p0, p2, p1)
                v31 := GenS3{F0: (p2 | p2), F1: 9223372036854775807, F2: genFn0("foo bar", ((4294967295 > 0) && false), "\n\t")}
                var v32 uint = 4294967295
                _ = v31
                _ = v32
            }
            _ = fi17
            _ = v18
        }
    } else {
        for fi33 := 0; fi33 < 14; fi33++ {
            var v34 uint = 42
            if (!((p0 > p0) || p1)) {
                v35 := GenS4{F0: (p2 << (((^(v34))) & 63))}
                fmt.Println(p0)
                if (!(p1 && (36 < -865))) {
                    for fi36 := 0; fi36 < 10; fi36++ {
                        v37 := []int{(fi33 % (((fi36 >> ((v34) & 63))) | 1)), 1, fi33}
                        if len(v37) > 0 {
                            v37[0] = (fi33 << (((v34 >> (((v34 * (v34 * v34))) & 63))) & 63))
                        }
                        _ = fi36
                        _ = v37
                    }
                    var v38 bool = (((0.1 / 1.0) < 0.0) || p1)
                    _ = v38
                }
                _ = v35
            }
            _ = fi33
            _ = v34
        }
        fmt.Println(p2)
    }
    var v39 string = "a"
    if p1 {
        for fi40 := 0; fi40 < 7; fi40++ {
            fmt.Println(p0, p0, fi40)
            _ = fi40
        }
        p1 = (false || p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v39
    return 691
}

func main() {
    v1 := []string{("" + "a"), "a"}
    v2 := []string{(("a" + "") + ""), ("hello" + (("foo bar" + "a") + "a")), "a", "\n\t"}
    v3 := make(map[uint32]string)
    var v4 uint8 = 127
    v5 := make(map[int]string)
    var v6 float64 = 3.14159
    fmt.Println(v4, v4)
    var v41 float64 = (v6 + v6)
    if (-18 >= -120) {
        var v42 uint64 = 863
        var v43 uint8 = (^(v4))
        v42 = (^(v42))
        _ = v42
        _ = v43
    } else {
        _ = len(v5)
        for fi44 := 0; fi44 < 7; fi44++ {
            var v45 uint16 = 2
            v46 := GenS4{F0: 178}
            _ = fi44
            _ = v45
            _ = v46
        }
    }
    switch ("a" + "\n\t") {
    case "a":
        v47 := GenS2{F0: 2, F1: 4294967295, F2: 0}
        _ = v47
    case "\n\t":
        var v48 uint64 = 60
        _ = v48
    default:
        v49 := GenS1{F0: -934}
        _ = v49
    }
    for fi50 := 0; fi50 < 8; fi50++ {
        switch (("" + "foo bar") + "a") {
        case "":
            if true {
                var v51 int = fi50
                fmt.Println(v51, v51, v51)
                _ = v51
            } else {
                v4 = (v4 ^ (^(v4)))
                v4 += v4
                if false {
                    var v52 bool = ((1 != 255) && (v4 <= 116))
                    var v53 int8 = 100
                    _ = v52
                    _ = v53
                }
            }
        case "a":
            v6 = (float64(0.1) * 0.0)
        default:
            v54 := GenS5{F0: 591}
            _ = v54
        }
        _ = fi50
    }
    v6 = v41
    switch 42 {
    case -9223372036854775808:
        v4 = (v4 ^ v4)
    }
    fmt.Println(v6)
    var v55 int16 = 1000
    for fi56 := 0; fi56 < 20; fi56++ {
        var v57 float32 = 0.1
        delete(v3, 2147483647)
        _ = fi56
        _ = v57
    }
    var v58 uint = genFn1(100, false, -1)
    v6 += v6
    v41 = v41
    v55 -= v55
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 128 {
    case 128:
        var v59 uint = v58
        _ = v59
    case 32767:
        switch 4294967295 {
        case 9223372036854775807:
            for fi60 := 0; fi60 < 13; fi60++ {
                v4 = 2
                for fi61 := 0; fi61 < 14; fi61++ {
                    v4 += (v4 / ((98) | 1))
                    _ = fi61
                }
                var v62 uint32 = 170
                _ = fi60
                _ = v62
            }
        }
    default:
        for fi63 := 0; fi63 < 8; fi63++ {
            switch "hello" {
            case "foo bar":
                v4 = (v4 << ((genFn1(-12, true, -9223372036854775808)) & 63))
            case "\n\t":
                if (1 == 703) {
                    var v64 int64 = -119
                    _ = v64
                } else {
                    if false {
                        v65 := GenS6{F0: 42, F1: (42 != 480), F2: (-1 > int8(fi63))}
                        _ = v65
                    } else {
                        v66 := GenS1{F0: (fi63 | (fi63 | (fi63 * (fi63 % (((fi63 / ((fi63) | 1))) | 1)))))}
                        _ = v66
                    }
                    fmt.Println(v55)
                    v4 += (v4 ^ 128)
                }
            case "a":
                v41 = v6
            default:
                for fi67 := 0; fi67 < 13; fi67++ {
                    fmt.Println(v55, fi63, v55)
                    v68 := GenS7{F0: 32768, F1: -1}
                    _ = fi67
                    _ = v68
                }
            }
            _ = fi63
        }
    }
    v6 += v6
    v69 := GenS8{F0: v58, F1: "", F2: ("" + "hello")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v55
    _ = v58
    _ = v6
    _ = v69
}
