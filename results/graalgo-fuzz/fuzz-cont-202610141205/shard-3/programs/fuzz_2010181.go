package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint; F1 uint64; F2 int8 }
func genFn0(p0 int64) uint64 {
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: 1, F1: 2}
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: false, F1: -32768}
    if true {
        v10 := []int{42, -348, -9223372036854775808, -801}
        for fi11 := 0; fi11 < 6; fi11++ {
            if true {
                v12 := GenS2{F0: (0 < 0), F1: 1000}
                _ = v12
            }
            _ = fi11
        }
        _ = v10
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := []int{}
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = v9
    return 28
}

func main() {
    var v1 uint16 = 255
    v2 := []uint32{}
    var v3 int16 = 0
    var v4 int16 = v3
    v5 := make(map[int]uint32)
    var v6 uint8 = 0
    v6 -= 42
    fmt.Println(v6, v3)
    v15 := GenS1{F0: 100, F1: 42}
    if true {
        var v16 int64 = -1
        _ = v16
    }
    fmt.Println(v1, v6)
    var v17 int = -456
    var v18 int8 = 127
    fmt.Println(v4, v17)
    if true {
        if (("hello" + ("hello" + "a")) >= "") {
            var v19 int = v17
            switch v17 {
            case -734:
                if false {
                    fmt.Println(v6)
                    v20 := v5[v19]
                    if (!(v1 != v1)) {
                        if (true && (((^(v17)) == (v17 / ((2147483647) | 1))) && false)) {
                            v19 = (v19 ^ v17)
                            if (!((0.1 / -2.25) >= 3.14159)) {
                                fmt.Println(v4)
                                switch v19 {
                                case -849:
                                    var v21 uint32 = 829
                                    _ = v21
                                }
                            } else {
                                switch "a" {
                                case "a":
                                    v17 = (v17 % (((-(v17))) | 1))
                                case "hello":
                                    var v22 int16 = (v3 | (v4 << ((18446744073709551615) & 63)))
                                    _ = v22
                                default:
                                    if true {
                                        if false {
                                            _ = len(v5)
                                            fmt.Println(v19, v6)
                                        } else {
                                            v23 := GenS2{F0: false, F1: (v4 | v4)}
                                            _ = v23
                                        }
                                        fmt.Println(v3, v17, v18)
                                    } else {
                                        var v24 uint16 = (v1 * v1)
                                        fmt.Println(v1, v19, v17)
                                        var v25 int8 = (v18 | 0)
                                        _ = v24
                                        _ = v25
                                    }
                                }
                                for fi26 := 0; fi26 < 1; fi26++ {
                                    var v27 float32 = 1e10
                                    _ = fi26
                                    _ = v27
                                }
                            }
                            var v28 int = (v17 + (v17 * (v17 | v19)))
                            _ = v28
                        }
                    } else {
                        fmt.Println(v1, v19)
                        for fi29 := 0; fi29 < 1; fi29++ {
                            v5[128] = v20
                            var v30 int32 = 426
                            _ = fi29
                            _ = v30
                        }
                        switch (("" + ("" + "")) + "a") {
                        case "hello":
                            v31 := GenS1{F0: v18, F1: (v20 | (^(v20)))}
                            _ = v31
                        case "\n\t":
                            switch int(v20) {
                            case 1:
                                for fi32 := 0; fi32 < 10; fi32++ {
                                    var v33 int64 = 9223372036854775807
                                    v34 := GenS2{F0: (v33 > 1), F1: v3}
                                    _ = fi32
                                    _ = v33
                                    _ = v34
                                }
                            case 128:
                                if (!(255 != 752)) {
                                    var v35 uint8 = v6
                                    var v36 int8 = (-(v18))
                                    switch ("hello" + "") {
                                    case "foo bar":
                                        if false {
                                            v1 = v1
                                            var v37 int32 = 282
                                            for fi38 := 0; fi38 < 3; fi38++ {
                                                v17 -= -85
                                                var v39 float64 = (-2.25 / 1e10)
                                                _ = fi38
                                                _ = v39
                                            }
                                            _ = v37
                                        }
                                    case "a":
                                        for fi40 := 0; fi40 < 13; fi40++ {
                                            for fi41 := 0; fi41 < 3; fi41++ {
                                                if false {
                                                    if (2 == 81) {
                                                        v3 -= (v4 ^ (v4 | v3))
                                                    }
                                                } else {
                                                    if len(v2) > 0 {
                                                        v2[0] = v20
                                                    }
                                                }
                                                var v42 uint = 42
                                                _ = fi41
                                                _ = v42
                                            }
                                            _ = fi40
                                        }
                                    }
                                    _ = v35
                                    _ = v36
                                }
                            default:
                                var v43 uint8 = v6
                                _ = v43
                            }
                        }
                    }
                    _ = v20
                }
            case 713:
                v19 -= (v19 / (((v19 + 111)) | 1))
            case 42:
                fmt.Println(v19, v4)
            default:
                for fi44 := 0; fi44 < 12; fi44++ {
                    v19 += 496
                    v4 = v4
                    v3 = (v3 ^ (-(v4)))
                    _ = fi44
                }
            }
            v18 -= 127
            _ = v19
        } else {
            var v45 bool = (false && ("" <= ("foo bar" + "a")))
            delete(v5, (v17 + v17))
            for fi46 := 0; fi46 < 20; fi46++ {
                fmt.Println(fi46, v4)
                for fi47 := 0; fi47 < 11; fi47++ {
                    _ = len(v5)
                    switch (fi47 ^ -855) {
                    case 32767:
                        if (!(-251 <= 2147483647)) {
                            var v48 float32 = 1.5
                            v17 = fi47
                            v18 -= v18
                            v4 = (v4 >> ((426) & 63))
                            v49 := GenS3{F0: -1}
                            _ = v48
                            _ = v49
                        }
                    default:
                        if v45 {
                            v3 = v3
                            switch "a" {
                            case "hello":
                                v2 = append(v2, 2147483648)
                            case "":
                                switch v17 {
                                case 128:
                                    v17 = (-(v17))
                                }
                            case "foo bar":
                                switch "hello" {
                                case "a":
                                    fmt.Println(v17, fi46)
                                    _ = len(v5)
                                case "":
                                    var v50 string = "a"
                                    _ = v50
                                default:
                                    v4 = v3
                                }
                                if false {
                                    switch "a" {
                                    case "":
                                        if (((-2.25 / ((0.0 / 1.5) - -2.25)) + 3.14159) == float64(-2.25)) {
                                            switch ("hello" + "") {
                                            case "":
                                                fmt.Println(fi47, v17)
                                            case "hello":
                                                switch ("" + "a") {
                                                case "a":
                                                    var v51 uint32 = 0
                                                    _ = v51
                                                }
                                            default:
                                                _ = len(v5)
                                            }
                                            fmt.Println(v4, v45, v18)
                                            var v52 float32 = 0.1
                                            _ = v52
                                        }
                                    case "\n\t":
                                        fmt.Println(v18, v1)
                                    case "a":
                                        v4 = (v3 % (((v4 & (v3 ^ (v4 / ((v3) | 1))))) | 1))
                                    default:
                                        var v53 int8 = (v18 + (v18 % (((v18 - (-(v18)))) | 1)))
                                        _ = v53
                                    }
                                    switch ("\n\t" + "foo bar") {
                                    case "foo bar":
                                        _ = len(v2)
                                    case "":
                                        if (((-2.25 / ((1.5 / ((1.5 * 0.1) * -2.25)) + (float64(-2.25) / ((-2.25 + 0.1) * 3.14159)))) + -2.25) != (float64(((3.14159 / 1.5) * 1.5)) / 1.0)) {
                                            v54 := v5[-537]
                                            var v55 string = "a"
                                            switch 127 {
                                            case 127:
                                                if v45 {
                                                    var v56 float64 = 1e10
                                                    var v57 uint = 0
                                                    v58 := GenS4{F0: 0, F1: 2147483647, F2: v54}
                                                    switch ("foo bar" + v55) {
                                                    case "foo bar":
                                                        v6 -= 47
                                                    case "":
                                                        var v59 int32 = -540
                                                        _ = v59
                                                    case "\n\t":
                                                        var v60 int16 = (v4 | (v4 | v4))
                                                        v1 = (v1 << ((uint(v54)) & 63))
                                                        _ = v60
                                                    default:
                                                        fmt.Println(v45, v4)
                                                    }
                                                    _ = v56
                                                    _ = v57
                                                    _ = v58
                                                } else {
                                                    v61 := GenS4{F0: 1, F1: (v54 * (v54 * (v54 ^ 2))), F2: 461}
                                                    var v62 int = fi47
                                                    _ = v61
                                                    _ = v62
                                                }
                                            }
                                            _ = v54
                                            _ = v55
                                        }
                                    }
                                }
                            default:
                                var v63 uint = 255
                                _ = v63
                            }
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                var v64 uint32 = 89
                                _ = v64
                            case "hello":
                                v1 -= 42
                            case "a":
                                v18 -= v18
                            }
                        }
                    }
                    _ = fi47
                }
                _ = fi46
            }
            _ = v45
        }
        switch v17 {
        case 619:
            for fi65 := 0; fi65 < 11; fi65++ {
                v1 += (^(v1))
                _ = fi65
            }
        case -687:
            v3 -= (v3 & (-(v4)))
        case -1:
            v4 = v4
        default:
            v4 += -748
        }
        var v66 int = (v17 ^ v17)
        _ = v66
    } else {
        v3 += -534
        for fi67 := 0; fi67 < 17; fi67++ {
            v4 = (v3 + (v4 - (v4 & (v4 ^ 1))))
            _ = fi67
        }
        v1 = v1
    }
    v68 := GenS4{F0: 42, F1: 829, F2: 894}
    for fi69 := 0; fi69 < 5; fi69++ {
        v1 += v1
        fmt.Println(v17)
        v70 := GenS5{F0: 65535, F1: 2, F2: (^(v18))}
        _ = fi69
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
}
