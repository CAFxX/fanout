package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 uint64; F1 int; F2 int16 }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 uint16; F1 uint16; F2 uint }
type GenS8 struct { F0 string; F1 uint8 }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint; F1 int32 }
func genFn0(p0 uint32, p1 string) int {
    for fi7 := 0; fi7 < 12; fi7++ {
        p0 = 381
        p1 = p1
        p0 = p0
        _ = fi7
    }
    p0 = (^(p0))
    v8 := GenS1{F0: 65535}
    if (!false) {
        p0 = (p0 - p0)
    } else {
        var v9 float32 = -2.25
        v10 := GenS2{F0: -1, F1: 14}
        v11 := []int{}
        _ = v10
        _ = v11
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v8
    return -1
}

func main() {
    v1 := []uint32{4294967295, 2147483647, 760, 42}
    v2 := make(map[uint32]int)
    var v3 float64 = 1e10
    var v4 float32 = float32(0.1)
    var v5 string = "hello"
    var v6 int64 = -263
    v12 := GenS2{F0: 0, F1: (v6 << ((4294967295) & 63))}
    var v13 float64 = 0.0
    var v14 uint = 134
    v6 = (v6 >> ((424) & 63))
    if (!(-2147483648 >= 127)) {
        v6 = v6
        v4 = ((float32(0.1) * 3.14159) / (v4 - 1e10))
    }
    switch -174 {
    case -2147483648:
        var v15 bool = true
        _ = v15
    case 252:
        var v16 float32 = 3.14159
        _ = v16
    }
    var v17 uint16 = 65535
    v13 = v13
    switch ((v5 + v5) + (("\n\t" + v5) + "\n\t")) {
    case "\n\t":
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := GenS3{F0: int16(fi18)}
            _ = fi18
            _ = v19
        }
    case "hello":
        v20 := GenS4{F0: ((^(v6)) == -1), F1: uint8(v17)}
        _ = v20
    default:
        if (!(((v3 - (v3 + v13)) / v3) >= 1.5)) {
            if (!(0 != -128)) {
                v5 = ("a" + v5)
                v13 += float64(v4)
                v5 = ("\n\t" + v5)
            }
            for fi21 := 0; fi21 < 19; fi21++ {
                if false {
                    if len(v1) > 0 {
                        v1[0] = 0
                    }
                    v1 = append(v1, 4294967295)
                    var v22 uint64 = 963
                    _ = v22
                } else {
                    v23 := GenS5{F0: 669, F1: 1, F2: 0}
                    for fi24 := 0; fi24 < 6; fi24++ {
                        v25 := GenS3{F0: -32768}
                        fmt.Println(v6)
                        _ = fi24
                        _ = v25
                    }
                    _ = v23
                }
                _ = fi21
            }
            v14 += v14
        }
    }
    v26 := GenS4{F0: (!true), F1: 208}
    switch 130 {
    case 0:
        if (false && false) {
            if true {
                if (!(true && (-1 <= 127))) {
                    if (!false) {
                        v17 -= 1
                        var v27 int64 = v6
                        _ = v27
                    }
                    if true {
                        v13 -= v3
                    }
                } else {
                    switch genFn0(641, v5) {
                    case -1:
                        v28 := GenS6{F0: 0, F1: 128}
                        _ = v28
                    case 2147483647:
                        for fi29 := 0; fi29 < 3; fi29++ {
                            fmt.Println(v14, v3, v6)
                            switch v5 {
                            case "\n\t":
                                v30 := GenS5{F0: 783, F1: genFn0(2, "hello"), F2: 0}
                                _ = v30
                            case "hello":
                                for fi31 := 0; fi31 < 3; fi31++ {
                                    switch (("" + v5) + "") {
                                    case "a":
                                        switch (fi29 << ((18446744073709551615) & 63)) {
                                        case 32768:
                                            var v32 float64 = 3.14159
                                            _ = v32
                                        }
                                    case "foo bar":
                                        var v33 float64 = 1.5
                                        _ = v33
                                    }
                                    _ = fi31
                                }
                            case "a":
                                v4 = v4
                            default:
                                var v34 int = (fi29 & fi29)
                                _ = v34
                            }
                            _ = fi29
                        }
                    case -9223372036854775808:
                        for fi35 := 0; fi35 < 2; fi35++ {
                            v36 := GenS6{F0: 2147483647, F1: 1}
                            _ = fi35
                            _ = v36
                        }
                    default:
                        v4 = v4
                    }
                }
            } else {
                _ = len(v2)
                if (!false) {
                    for fi37 := 0; fi37 < 2; fi37++ {
                        v38 := GenS7{F0: (v17 % (((v17 << ((4294967295) & 63))) | 1)), F1: v17, F2: v14}
                        fmt.Println(v5)
                        var v39 uint64 = 42
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                } else {
                    for fi40 := 0; fi40 < 19; fi40++ {
                        v6 = v6
                        switch (v5 + (v5 + v5)) {
                        case "hello":
                            v14 -= v14
                        case "a":
                            var v41 float32 = float32(-2.25)
                            _ = v41
                        case "\n\t":
                            for fi42 := 0; fi42 < 12; fi42++ {
                                var v43 int64 = (v6 * (v6 - v6))
                                if (v43 != -835) {
                                    fmt.Println(v43, v13, v4)
                                    fmt.Println(v13)
                                    v5 = (v5 + ("a" + ""))
                                    if false {
                                        v44 := GenS2{F0: 683, F1: (v43 - (v43 | (v6 + (v6 / ((v43) | 1)))))}
                                        v45 := GenS8{F0: (v5 + ((v5 + v5) + ((v5 + "hello") + v5))), F1: 127}
                                        _ = v44
                                        _ = v45
                                    } else {
                                        if (1 != (fi42 - genFn0(714, "foo bar"))) {
                                            v1 = append(v1, 65535)
                                            var v46 int64 = (-(v43))
                                            _ = v46
                                        } else {
                                            var v47 uint = 509
                                            _ = v47
                                        }
                                        v5 = (("hello" + v5) + v5)
                                        var v48 uint64 = 1
                                        _ = v48
                                    }
                                    v3 -= float64(1e10)
                                    var v49 uint64 = 4294967295
                                    _ = len(v2)
                                    _ = v49
                                } else {
                                    _ = len(v2)
                                    var v50 int8 = -1
                                    if (0 <= int16(v43)) {
                                        var v51 bool = true
                                        var v52 uint16 = v17
                                        v53 := GenS3{F0: 193}
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    } else {
                                        var v54 string = v5
                                        _ = v54
                                    }
                                    v17 += 1
                                    _ = v50
                                }
                                switch (fi42 | genFn0(0, v5)) {
                                case 0:
                                    var v55 uint64 = 992
                                    _ = v55
                                case -993:
                                    var v56 int16 = 1
                                    _ = v56
                                default:
                                    if (!(v6 <= (v43 + (v6 << (((v14 + (v14 % ((v14) | 1)))) & 63))))) {
                                        if (753 <= (v14 | (v14 - (v14 / (((^(v14))) | 1))))) {
                                            v57 := GenS9{F0: 42, F1: 0}
                                            if true {
                                                var v58 int8 = 111
                                                _ = v58
                                            }
                                            _ = v57
                                        } else {
                                            if false {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = len(v1)
                                                fmt.Println(v17, fi40)
                                            }
                                            switch ((v5 + (("foo bar" + ((("\n\t" + "") + v5) + v5)) + "")) + v5) {
                                            case "hello":
                                                v3 -= v13
                                            default:
                                                if (127 >= 255) {
                                                    if len(v1) > 0 {
                                                        v1[0] = 481
                                                    }
                                                    v6 += v6
                                                    _ = len(v1)
                                                    switch (fi42 / (((fi42 / (((-(fi40))) | 1))) | 1)) {
                                                    case -2147483648:
                                                        var v59 float32 = (v4 - ((-2.25 - (v4 * ((1e10 + v4) / (float32(1e10) / 3.14159)))) + v4))
                                                        _ = v59
                                                    }
                                                    if (650 > 756) {
                                                        v13 -= 1.5
                                                    } else {
                                                        if false {
                                                            var v60 bool = true
                                                            fmt.Println(v4, v13, v43)
                                                            fmt.Println(v5)
                                                            var v61 int = -259
                                                            v62 := v2[1]
                                                            var v63 uint64 = 525
                                                            _ = v60
                                                            _ = v61
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                        v64 := GenS10{F0: -2147483648}
                                                        _ = v64
                                                    }
                                                    var v65 float64 = float64(float32((3.14159 + (-2.25 - v3))))
                                                    _ = v65
                                                }
                                            }
                                        }
                                        v13 -= (1e10 / (v3 * v3))
                                        var v66 int8 = 1
                                        _ = v66
                                    } else {
                                        v4 = v4
                                        fmt.Println(v43)
                                        var v67 int32 = 2147483647
                                        _ = v67
                                    }
                                    var v68 bool = (!(v3 < ((v3 * 0.1) / (v13 * 1.5))))
                                    v1 = append(v1, 87)
                                    _ = v68
                                }
                                _ = fi42
                                _ = v43
                            }
                        default:
                            v69 := GenS11{F0: v14, F1: -1}
                            _ = v69
                        }
                        v70 := v2[uint32(v14)]
                        _ = fi40
                        _ = v70
                    }
                }
                if (255 != (v14 + 663)) {
                    if false {
                        _ = len(v2)
                    } else {
                        var v71 uint16 = v17
                        _ = v71
                    }
                    delete(v2, 4294967295)
                }
            }
            v2[65535] = 123
        }
    case -1:
        var v72 uint16 = (v17 % ((v17) | 1))
        _ = v72
    case -760:
        v73 := v2[293]
        _ = v73
    }
    var v74 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
}
