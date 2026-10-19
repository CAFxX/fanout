package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 uint8; F1 int16; F2 int }
type GenS7 struct { F0 string; F1 int16; F2 uint }
func genFn0(p0 bool, p1 int16, p2 int8) int32 {
    if p0 {
        for fi5 := 0; fi5 < 6; fi5++ {
            if ((^(fi5)) >= fi5) {
                p1 += -333
            }
            _ = fi5
        }
        if (128 <= 677) {
            if p0 {
                v6 := GenS1{F0: 1, F1: 255}
                for fi7 := 0; fi7 < 18; fi7++ {
                    for fi8 := 0; fi8 < 4; fi8++ {
                        var v9 int16 = p1
                        _ = fi8
                        _ = v9
                    }
                    p1 = p1
                    if true {
                        fmt.Println(fi7)
                        switch "" {
                        case "hello":
                            var v10 uint = 42
                            _ = v10
                        }
                        p2 -= p2
                    } else {
                        for fi11 := 0; fi11 < 2; fi11++ {
                            fmt.Println(p0)
                            _ = fi11
                        }
                        p2 = (-(p2))
                    }
                    _ = fi7
                }
                p1 += (p1 ^ p1)
                _ = v6
            } else {
                var v12 uint = 413
                _ = v12
            }
            switch ("a" + ("foo bar" + "hello")) {
            case "hello":
                v13 := []int{224, 32768, -1, 9223372036854775807}
                _ = v13
            case "\n\t":
                v14 := GenS1{F0: 32768, F1: 251}
                _ = v14
            case "":
                for fi15 := 0; fi15 < 17; fi15++ {
                    var v16 uint32 = 930
                    var v17 float64 = ((0.0 * 3.14159) * (float64(-2.25) - 1.5))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            default:
                var v18 uint64 = 1
                _ = v18
            }
        }
    }
    var v19 float64 = float64((0.1 * ((-2.25 / (3.14159 * -2.25)) / 1.0)))
    v19 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 172
}

func main() {
    var v1 int32 = 42
    var v2 uint = 42
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    switch 0 {
    case 1:
        if true {
            for fi20 := 0; fi20 < 5; fi20++ {
                v4["\n\t"] = "a"
                switch (fi20 / ((2147483647) | 1)) {
                case 42:
                    var v21 float32 = 0.1
                    _ = v21
                case 0:
                    v2 += v2
                case 32768:
                    fmt.Println(fi20, fi20)
                }
                v22 := GenS2{F0: 42, F1: 1000}
                _ = fi20
                _ = v22
            }
        }
    case 127:
        switch "a" {
        case "hello":
            for fi23 := 0; fi23 < 16; fi23++ {
                v1 = v1
                if ("\n\t" == "a") {
                    v1 += (v1 + 390)
                    v24 := []int{9223372036854775807, int(v2), (fi23 - (fi23 % ((-340) | 1)))}
                    _ = v24
                } else {
                    fmt.Println(v1, v2, fi23)
                    if true {
                        v25 := GenS1{F0: 890, F1: (v2 - (v2 >> (((v2 ^ (v2 + 4294967295))) & 63)))}
                        v26 := GenS3{F0: true}
                        if false {
                            var v27 int8 = -128
                            fmt.Println(fi23, fi23, v1)
                            var v28 string = ("" + "\n\t")
                            _ = v27
                            _ = v28
                        }
                        _ = v25
                        _ = v26
                    } else {
                        v2 += (v2 + 65535)
                        switch 42 {
                        case 127:
                            v1 += (^(v1))
                        case 0:
                            v29 := GenS3{F0: (false || false)}
                            _ = v29
                        case -220:
                            if false {
                                v30 := GenS4{F0: 1, F1: 922, F2: (v1 & (v1 * (v1 - (v1 + -1))))}
                                fmt.Println(fi23, fi23, v2)
                                _ = v30
                            } else {
                                fmt.Println(fi23, v1, fi23)
                                v31 := v4[(("a" + "foo bar") + ("a" + ("a" + "a")))]
                                _ = v31
                            }
                        default:
                            for fi32 := 0; fi32 < 19; fi32++ {
                                var v33 uint16 = 42
                                v34 := GenS5{F0: 42, F1: (fi32 ^ (fi32 & (fi23 >> ((v2) & 63))))}
                                if false {
                                    var v35 uint8 = 127
                                    var v36 uint8 = 2
                                    _ = v35
                                    _ = v36
                                } else {
                                    v37 := GenS6{F0: 127, F1: -783, F2: fi23}
                                    _ = v37
                                }
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                        }
                    }
                    v38 := GenS7{F0: "", F1: -196, F2: (^(v2))}
                    _ = v38
                }
                _ = fi23
            }
        case "foo bar":
            for fi39 := 0; fi39 < 19; fi39++ {
                var v40 int32 = int32(v2)
                var v41 float32 = (((float32(1e10) / -2.25) / 1.5) - 3.14159)
                if (true && true) {
                    v42 := GenS4{F0: -9, F1: 42, F2: v1}
                    for fi43 := 0; fi43 < 17; fi43++ {
                        if (float64(v41) == float64(((0.1 / v41) - -2.25))) {
                            if false {
                                fmt.Println(v1)
                                _ = len(v4)
                            }
                            if false {
                                switch fi43 {
                                case 127:
                                    v40 = (v1 ^ (v40 >> ((v2) & 63)))
                                case -497:
                                    fmt.Println(v2, v2, v40)
                                case 42:
                                    v1 += v1
                                }
                            }
                            switch ("\n\t" + "") {
                            case "":
                                var v44 uint = (v2 / ((v2) | 1))
                                _ = v44
                            default:
                                var v45 float32 = ((v41 / 0.1) / v41)
                                _ = v45
                            }
                        }
                        _ = fi43
                    }
                    var v46 uint32 = 2147483648
                    _ = v42
                    _ = v46
                }
                _ = fi39
                _ = v40
                _ = v41
            }
        case "\n\t":
            var v47 float32 = (3.14159 / (0.0 + 0.1))
            _ = v47
        }
    case 9223372036854775807:
        v2 = 2
    default:
        fmt.Println(v1, v2)
    }
    for fi48 := 0; fi48 < 6; fi48++ {
        v1 = genFn0((false && (0 != 9223372036854775807)), 1, 0)
        _ = fi48
    }
    if false {
        if (1 > 100) {
            if false {
                var v49 float32 = -2.25
                v50 := GenS4{F0: 120, F1: 32767, F2: genFn0((-287 != 42), -32768, 42)}
                _ = v49
                _ = v50
            } else {
                var v51 int32 = genFn0(false, 32767, -128)
                for fi52 := 0; fi52 < 12; fi52++ {
                    fmt.Println(v1, v2)
                    var v53 int16 = 1
                    _ = fi52
                    _ = v53
                }
                _ = v51
            }
            v54 := []int{127, int(v1), -1, -2147483648}
            _ = v54
        } else {
            var v55 int16 = 1
            _ = v55
        }
        var v56 uint16 = 255
        v2 = v2
        _ = v56
    } else {
        switch "\n\t" {
        case "foo bar":
            v3[0] = 4294967295
        case "":
            var v57 uint = 255
            _ = v57
        case "hello":
            v58 := []int{}
            _ = v58
        default:
            v59 := GenS3{F0: true}
            _ = v59
        }
        v60 := GenS4{F0: -24, F1: 299, F2: (v1 % ((genFn0((v1 >= (v1 % ((1) | 1))), -967, -1)) | 1))}
        _ = v60
    }
    fmt.Println(v1)
    var v61 uint32 = 2
    var v62 int64 = 42
    v1 -= genFn0((false || true), 42, -128)
    var v63 uint64 = 4294967295
    v64 := GenS3{F0: true}
    for fi65 := 0; fi65 < 6; fi65++ {
        v66 := GenS2{F0: 1, F1: -671}
        v63 = (v63 / (((v63 / ((2) | 1))) | 1))
        _ = fi65
        _ = v66
    }
    for fi67 := 0; fi67 < 8; fi67++ {
        var v68 int16 = 1000
        v69 := GenS5{F0: 2, F1: fi67}
        var v70 int8 = -1
        _ = fi67
        _ = v68
        _ = v69
        _ = v70
    }
    var v71 int8 = -76
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v71
}
