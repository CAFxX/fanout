package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16; F1 int32; F2 uint16 }

func main() {
    var v1 bool = (!true)
    var v2 uint32 = 42
    var v3 uint = 0
    var v4 bool = v1
    v3 = (v3 | uint(v2))
    if (0.0 >= ((float64(((1e10 * 1e10) - 1.5)) - 1.5) + -2.25)) {
        if v1 {
            v5 := GenS1{F0: 255, F1: 2, F2: v2}
            _ = v5
        } else {
            var v6 uint16 = uint16(v3)
            v7 := []int{}
            _ = v6
            _ = v7
        }
        if (1 >= -69) {
            v3 = (v3 - (v3 * v3))
        }
        if v1 {
            v8 := []int{-1, 863, 90, -9223372036854775808}
            v2 = (^(v2))
            v4 = v4
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 uint16 = 2
            if (42 <= -128) {
                if v1 {
                    if true {
                        v11 := GenS2{F0: 127}
                        _ = v11
                    }
                } else {
                    if true {
                        v12 := []int{fi9}
                        var v13 uint = (^(v3))
                        for fi14 := 0; fi14 < 12; fi14++ {
                            v15 := GenS2{F0: (fi9 + fi14)}
                            if (!(!(0 > -966))) {
                                fmt.Println(v2, v2)
                                v1 = (-328 != 1)
                                if ((("hello" + "\n\t") + "\n\t") < ("" + "hello")) {
                                    var v16 float64 = (-2.25 + ((3.14159 - (3.14159 + 0.1)) - (0.0 / ((1e10 - ((1.5 * (-2.25 * 1e10)) + ((0.1 / 3.14159) - (-2.25 - 1e10)))) + 0.0))))
                                    v4 = false
                                    _ = v16
                                } else {
                                    fmt.Println(v10, fi14)
                                    v12 = append(v12, (fi14 ^ -2147483648))
                                    if v4 {
                                        v1 = true
                                        _ = len(v12)
                                    }
                                }
                            }
                            v17 := GenS1{F0: v3, F1: 244, F2: (v2 >> (((^(v3))) & 63))}
                            _ = fi14
                            _ = v15
                            _ = v17
                        }
                        _ = v12
                        _ = v13
                    } else {
                        v10 = 42
                        v18 := []int{(fi9 ^ fi9), (fi9 + (^(fi9))), 0}
                        _ = v18
                    }
                    v19 := []int{fi9}
                    _ = v19
                }
                v20 := []int{}
                if (-1 == 466) {
                    for fi21 := 0; fi21 < 3; fi21++ {
                        switch fi21 {
                        case -63:
                            var v22 float32 = 1e10
                            _ = v22
                        case 1:
                            v3 += 2
                        }
                        v23 := GenS3{F0: 42, F1: 128, F2: uint(v10)}
                        _ = fi21
                        _ = v23
                    }
                }
                _ = v20
            }
            if true {
                for fi24 := 0; fi24 < 13; fi24++ {
                    if v1 {
                        switch -990 {
                        case 2147483647:
                            if (776 < (^(v3))) {
                                var v25 uint = 18446744073709551615
                                _ = v25
                            } else {
                                v26 := []int{fi24, fi9, (fi24 | fi24), int(v3)}
                                v1 = true
                                v1 = (int(v2) > 2147483647)
                                _ = v26
                            }
                        case 42:
                            switch ("\n\t" + "foo bar") {
                            case "\n\t":
                                v3 = (v3 >> ((v3) & 63))
                            case "a":
                                var v27 uint32 = (v2 * (^(v2)))
                                _ = v27
                            }
                        default:
                            switch "a" {
                            case "":
                                var v28 uint32 = v2
                                _ = v28
                            }
                        }
                        var v29 string = ("hello" + "a")
                        _ = v29
                    } else {
                        v30 := []int{}
                        _ = v30
                    }
                    fmt.Println(fi9, v2, fi24)
                    fmt.Println(fi24)
                    if true {
                        if ((((1.5 * 1e10) * 3.14159) / ((1e10 - 1.5) - 0.0)) != 0.0) {
                            v31 := []int{(fi24 ^ 2147483647), fi24}
                            v10 -= 65535
                            v31 = append(v31, -9223372036854775808)
                            _ = v31
                        }
                        v32 := []int{(fi9 % ((-2147483648) | 1))}
                        var v33 int16 = 32767
                        var v34 int = (fi9 << ((42) & 63))
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    _ = fi24
                }
                v35 := GenS4{F0: 1, F1: 255}
                _ = v35
            }
            _ = fi9
            _ = v10
        }
        v36 := []int{}
        _ = v36
    }
    for fi37 := 0; fi37 < 9; fi37++ {
        v2 = (v2 << ((v3) & 63))
        v38 := []int{fi37}
        _ = fi37
        _ = v38
    }
    v1 = (2147483647 > 2147483647)
    for fi39 := 0; fi39 < 14; fi39++ {
        var v40 int64 = 9223372036854775807
        var v41 uint16 = 414
        fmt.Println(v2, fi39, v2)
        _ = fi39
        _ = v40
        _ = v41
    }
    switch -48 {
    case -1:
        if false {
            for fi42 := 0; fi42 < 14; fi42++ {
                if (v3 == v3) {
                    v43 := []int{(fi42 ^ fi42), 4294967295, fi42}
                    _ = v43
                } else {
                    switch fi42 {
                    case 1:
                        if (v1 && ((v1 || false) && (-1 > 510))) {
                            if (v4 || (!(1.5 != (float64((0.0 - 1.5)) + -2.25)))) {
                                v44 := GenS4{F0: 0, F1: 0}
                                v45 := GenS2{F0: (fi42 * 2147483647)}
                                var v46 uint16 = 1
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                            v2 -= (v2 << ((v3) & 63))
                        }
                    case 128:
                        var v47 uint16 = 65535
                        _ = v47
                    case -2147483648:
                        fmt.Println(v2, fi42)
                    default:
                        fmt.Println(fi42, v2)
                    }
                    v48 := []int{fi42, 128, (^(fi42))}
                    _ = v48
                }
                _ = fi42
            }
            v4 = ((^(v3)) != 430)
            for fi49 := 0; fi49 < 16; fi49++ {
                v50 := []int{32768}
                var v51 int32 = 0
                _ = fi49
                _ = v50
                _ = v51
            }
        }
    case 127:
        var v52 string = "a"
        _ = v52
    }
    v53 := GenS1{F0: uint(v2), F1: 87, F2: (v2 & v2)}
    if v1 {
        v54 := GenS5{F0: uint16(v2), F1: 0}
        _ = v54
    }
    if (2147483647 < -2147483648) {
        if true {
            if v1 {
                fmt.Println(v1, v2)
                switch 32768 {
                case 4294967295:
                    if (!v1) {
                        v1 = false
                        for fi55 := 0; fi55 < 11; fi55++ {
                            var v56 uint = v3
                            var v57 uint16 = 58
                            var v58 uint64 = 1
                            _ = fi55
                            _ = v56
                            _ = v57
                            _ = v58
                        }
                        v59 := GenS1{F0: 65535, F1: 127, F2: (v2 | (^(v2)))}
                        _ = v59
                    }
                default:
                    switch -49 {
                    case 910:
                        switch -175 {
                        case -484:
                            fmt.Println(v1, v3)
                        }
                    default:
                        v4 = (!v1)
                    }
                }
                v60 := []int{}
                _ = v60
            }
        }
    } else {
        fmt.Println(v2, v4, v1)
    }
    var v61 uint = v3
    var v62 int16 = 833
    var v63 uint64 = 784
    if (!(!(1e10 >= float64((0.0 * 1.5))))) {
        v64 := GenS6{F0: -2147483648}
        _ = v64
    }
    for fi65 := 0; fi65 < 13; fi65++ {
        v66 := []int{(fi65 % ((9223372036854775807) | 1)), 0, fi65}
        if len(v66) > 0 {
            v66[0] = 128
        }
        _ = fi65
        _ = v66
    }
    for fi67 := 0; fi67 < 2; fi67++ {
        if v1 {
            v68 := []int{fi67, (fi67 / ((32768) | 1))}
            _ = v68
        } else {
            v3 = 696
            if v1 {
                v61 = v61
                for fi69 := 0; fi69 < 18; fi69++ {
                    fmt.Println(v62, v4, v61)
                    _ = fi69
                }
            }
            fmt.Println(v1)
        }
        var v70 uint8 = 223
        switch "foo bar" {
        case "foo bar":
            var v71 int32 = -2147483648
            _ = v71
        default:
            if (-923 < 42) {
                v3 = 2
                v62 = (v62 >> ((v61) & 63))
                for fi72 := 0; fi72 < 12; fi72++ {
                    v1 = (42 != 1)
                    v73 := GenS5{F0: 2, F1: 42}
                    _ = fi72
                    _ = v73
                }
            }
        }
        _ = fi67
        _ = v70
    }
    v74 := GenS7{F0: v62, F1: 702, F2: 1}
    v75 := []int{1, 625, 805, 870}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("slicelen v75:", len(v75), v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v61
    _ = v62
    _ = v63
    _ = v74
    _ = v75
}
