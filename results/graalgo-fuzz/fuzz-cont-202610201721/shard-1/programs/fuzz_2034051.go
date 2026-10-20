package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int64; F2 int }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 bool; F1 int; F2 uint16 }
type GenS7 struct { F0 int; F1 string }
func genFn0(p0 uint, p1 int8) int16 {
    var v5 uint8 = 21
    var v6 int64 = -514
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int64, p1 string, p2 string) int8 {
    for fi7 := 0; fi7 < 17; fi7++ {
        switch p1 {
        case "a":
            switch fi7 {
            case 42:
                p0 += (p0 % ((-1) | 1))
            default:
                if (-97 == 125) {
                    v8 := []int{-2147483648, -1, 2147483647}
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 17; fi9++ {
                        p0 += p0
                        v10 := []int{fi7, 128, fi7}
                        v11 := GenS1{F0: (p0 - -9223372036854775808)}
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    var v12 float64 = ((((float64((3.14159 * 1e10)) * 0.0) - (float64(-2.25) / -2.25)) * 1.5) / 1e10)
                    for fi13 := 0; fi13 < 11; fi13++ {
                        v14 := []int{(fi7 | fi7), 1, fi7, (fi7 ^ fi7)}
                        _ = fi13
                        _ = v14
                    }
                    _ = v12
                }
            }
        }
        var v15 int32 = -1
        v16 := GenS1{F0: p0}
        _ = fi7
        _ = v15
        _ = v16
    }
    switch p2 {
    case "":
        switch p2 {
        case "hello":
            fmt.Println(p0, p1)
        case "foo bar":
            var v17 string = ("\n\t" + "\n\t")
            _ = v17
        case "":
            var v18 uint = 4294967295
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 16; fi19++ {
            if (((823 <= 4294967295) || (!(!(false || (1 == 32767))))) && (p0 != -1)) {
                for fi20 := 0; fi20 < 10; fi20++ {
                    p1 = "hello"
                    _ = fi20
                }
                v21 := GenS2{F0: (fi19 >> ((4294967295) & 63)), F1: (p0 + p0), F2: fi19}
                for fi22 := 0; fi22 < 18; fi22++ {
                    v23 := GenS2{F0: 32767, F1: 685, F2: -2147483648}
                    var v24 uint = 2
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                _ = v21
            } else {
                var v25 float64 = float64((0.1 / 1.5))
                if (-128 <= -1) {
                    var v26 int = fi19
                    p1 = "a"
                    var v27 string = "\n\t"
                    _ = v26
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 14; fi28++ {
                        fmt.Println(fi28)
                        v29 := []int{fi19}
                        _ = fi28
                        _ = v29
                    }
                    if true {
                        p1 = "\n\t"
                        v30 := GenS3{F0: 2, F1: -117}
                        var v31 int16 = -968
                        _ = v30
                        _ = v31
                    } else {
                        if true {
                            v25 = ((v25 / (v25 + (1.5 / (0.0 * v25)))) - ((v25 - 3.14159) + float64(1e10)))
                            for fi32 := 0; fi32 < 17; fi32++ {
                                if false {
                                    v25 -= v25
                                } else {
                                    v33 := GenS3{F0: 65535, F1: 0}
                                    p1 = (p1 + (p1 + ((((((p2 + ("foo bar" + "a")) + ("\n\t" + ((((("a" + "foo bar") + p2) + "a") + p1) + p1))) + "") + p1) + "a") + ("foo bar" + p1))))
                                    _ = v33
                                }
                                v34 := []int{fi32, (fi19 >> ((4294967295) & 63)), (^(fi19))}
                                _ = fi32
                                _ = v34
                            }
                        }
                        if true {
                            if true {
                                v35 := GenS2{F0: fi19, F1: (p0 ^ p0), F2: 2147483647}
                                var v36 uint = 0
                                if true {
                                    p0 = p0
                                }
                                _ = v35
                                _ = v36
                            }
                            var v37 int8 = -3
                            _ = v37
                        } else {
                            if false {
                                p2 = p1
                                p1 = (p2 + ("hello" + p2))
                            } else {
                                var v38 int64 = (p0 >> ((920) & 63))
                                _ = v38
                            }
                            var v39 float32 = (-2.25 * (float32(v25) - (float32((0.1 / v25)) + -2.25)))
                            var v40 uint8 = 8
                            _ = v39
                            _ = v40
                        }
                        p2 = p1
                    }
                }
                v41 := GenS4{F0: 4294967295, F1: 42}
                _ = v25
                _ = v41
            }
            v42 := []int{(fi19 % (((fi19 & (fi19 / ((fi19) | 1)))) | 1)), (fi19 << ((937) & 63)), (-(fi19))}
            _ = fi19
            _ = v42
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -59
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    var v4 int = -382
    fmt.Println(v4, v1)
    switch "hello" {
    case "foo bar":
        switch -935 {
        case -1:
            if (true && true) {
                if ("hello" != "hello") {
                    for fi43 := 0; fi43 < 13; fi43++ {
                        var v44 bool = (0 == -41)
                        _ = fi43
                        _ = v44
                    }
                    fmt.Println(v4, v1)
                    v4 += -198
                }
            }
        case 1:
            v45 := GenS2{F0: v4, F1: 9223372036854775807, F2: -156}
            _ = v45
        case 128:
            v1 += (v1 * (((1.5 - (3.14159 - float64((((float32(v1) / float32(v1)) / -2.25) / 1.5)))) - (((1.5 * 0.0) / (v1 - v1)) - v1)) - float64(0.1)))
        }
    case "":
        for fi46 := 0; fi46 < 20; fi46++ {
            switch 571 {
            case -656:
                v1 = ((((v1 * float64(-2.25)) + -2.25) + (float64(0.0) * 3.14159)) / v1)
            default:
                var v47 int = -9223372036854775808
                _ = v47
            }
            v48 := GenS4{F0: 858, F1: 956}
            v4 = fi46
            _ = fi46
            _ = v48
        }
    default:
        switch v4 {
        case 127:
            var v49 uint64 = 4294967295
            _ = v49
        case 1:
            v1 -= v1
        default:
            v50 := []int{v4}
            _ = v50
        }
    }
    if (-2.25 != (((float32(float64(-2.25)) + 3.14159) * (3.14159 / 3.14159)) * (1e10 * 0.1))) {
        for fi51 := 0; fi51 < 7; fi51++ {
            fmt.Println(fi51)
            var v52 int32 = 1
            _ = fi51
            _ = v52
        }
        for fi53 := 0; fi53 < 12; fi53++ {
            fmt.Println(v1, v4, fi53)
            var v54 uint = 4294967295
            _ = fi53
            _ = v54
        }
    }
    var v55 int = 459
    switch "hello" {
    case "foo bar":
        if false {
            if (false && (0 >= 1)) {
                for fi56 := 0; fi56 < 16; fi56++ {
                    v57 := []int{(v55 / ((fi56) | 1)), v55, (v55 ^ (v4 * fi56)), v55}
                    _ = fi56
                    _ = v57
                }
                v1 = v1
            } else {
                v58 := GenS5{F0: 42, F1: "\n\t"}
                var v59 uint8 = 88
                _ = v58
                _ = v59
            }
        }
    case "hello":
        if true {
            var v60 uint = 762
            fmt.Println(v1)
            var v61 int16 = -1
            _ = v60
            _ = v61
        }
    }
    fmt.Println(v4)
    v62 := []int{-2147483648, v55}
    fmt.Println(v55, v55, v1)
    switch "hello" {
    case "hello":
        for fi63 := 0; fi63 < 10; fi63++ {
            var v64 int = v55
            _ = fi63
            _ = v64
        }
    case "foo bar":
        for fi65 := 0; fi65 < 12; fi65++ {
            v66 := GenS3{F0: 1, F1: genFn1(-1, "hello", "\n\t")}
            switch "hello" {
            case "":
                for fi67 := 0; fi67 < 11; fi67++ {
                    var v68 string = "foo bar"
                    fmt.Println(v1, v68, fi65)
                    _ = fi67
                    _ = v68
                }
            default:
                var v69 float64 = float64(1e10)
                _ = v69
            }
            _ = fi65
            _ = v66
        }
    default:
        v55 = v4
    }
    if true {
        v55 = 0
        for fi70 := 0; fi70 < 10; fi70++ {
            v1 -= 1e10
            for fi71 := 0; fi71 < 20; fi71++ {
                v72 := GenS6{F0: false, F1: (v4 >> ((255) & 63)), F2: 255}
                _ = fi71
                _ = v72
            }
            _ = fi70
        }
        if false {
            v62 = append(v62, v4)
            v73 := GenS7{F0: v4, F1: ("a" + "hello")}
            _ = v73
        } else {
            fmt.Println(v55, v4)
            var v74 uint64 = 2
            _ = v74
        }
    }
    v55 -= v55
    switch (v55 * v4) {
    case -118:
        if (4294967295 != 936) {
            v1 += ((3.14159 - (v1 + v1)) * v1)
            var v75 uint32 = 0
            v2[v55] = -865
            _ = v75
        } else {
            v4 += v55
            _ = len(v62)
        }
    }
    v76 := GenS7{F0: (v4 ^ (^(v55))), F1: ""}
    if (!(!(false || ((genFn1(9223372036854775807, "", "hello") > 68) || true)))) {
        if false {
            v55 -= (v55 & (v4 - 0))
        } else {
            v4 = v4
            v77 := GenS3{F0: 2, F1: 7}
            _ = v77
        }
    } else {
        switch "foo bar" {
        case "foo bar":
            var v78 uint = 365
            _ = v78
        }
    }
    v1 += 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v62
    _ = v76
}
