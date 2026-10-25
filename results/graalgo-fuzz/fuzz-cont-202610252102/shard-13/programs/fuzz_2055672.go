package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint; F1 int8; F2 uint16 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 uint8; F1 uint }
func genFn0(p0 uint, p1 int64, p2 uint8) uint64 {
    p2 += (p2 / ((0) | 1))
    switch 0 {
    case 9223372036854775807:
        v7 := []int{-9223372036854775808, 1, 127, 738}
        _ = v7
    case 1:
        switch 4294967295 {
        case 32767:
            var v8 int = 798
            _ = v8
        default:
            var v9 int32 = 0
            _ = v9
        }
    case -704:
        var v10 float32 = ((3.14159 * 1e10) / 1.0)
        _ = v10
    default:
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 float64 = 0.0
            _ = fi11
            _ = v12
        }
    }
    if ((p0 << (((^(p0))) & 63)) != (p0 + p0)) {
        p1 -= p1
        if false {
            v13 := []int{}
            _ = v13
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            p0 -= p0
            var v15 uint8 = p2
            v16 := []int{fi14, fi14, fi14, fi14}
            _ = fi14
            _ = v15
            _ = v16
        }
    } else {
        p0 = 65535
    }
    v17 := GenS1{F0: -24, F1: (p0 - p0), F2: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 277
}
func genFn1() bool {
    var v18 string = ((("foo bar" + "foo bar") + "a") + "\n\t")
    v19 := GenS1{F0: 88, F1: 18446744073709551615, F2: 600}
    var v20 uint64 = 723
    var v21 bool = true
    v22 := GenS2{F0: 194, F1: "\n\t"}
    _ = v18
    _ = v19
    _ = v20
    _ = v21
    _ = v22
    return v21
}

func main() {
    var v1 string = ("hello" + "")
    v2 := []string{"foo bar"}
    var v3 int8 = -40
    var v4 float32 = (0.1 - 1.5)
    var v5 int16 = 1000
    var v6 int32 = -640
    var v23 uint8 = 128
    var v24 uint64 = genFn0(42, 77, v23)
    var v25 int16 = v5
    v26 := GenS3{F0: 42, F1: v3, F2: 436}
    v24 -= (^(v24))
    var v27 uint8 = 128
    if (v27 <= 133) {
        var v28 int16 = int16(v3)
        _ = v28
    } else {
        for fi29 := 0; fi29 < 6; fi29++ {
            switch ((("a" + v1) + "a") + v1) {
            case "hello":
                switch fi29 {
                case -661:
                    v5 -= -1
                case 127:
                    v2 = append(v2, "a")
                default:
                    v30 := GenS2{F0: v27, F1: v1}
                    _ = v30
                }
            }
            for fi31 := 0; fi31 < 14; fi31++ {
                v32 := GenS3{F0: 2, F1: (v3 ^ (v3 - (v3 & v3))), F2: 399}
                var v33 string = v1
                _ = fi31
                _ = v32
                _ = v33
            }
            v24 = v24
            _ = fi29
        }
        for fi34 := 0; fi34 < 3; fi34++ {
            if (false && false) {
                for fi35 := 0; fi35 < 16; fi35++ {
                    switch 1 {
                    case 42:
                        var v36 uint32 = 2147483648
                        _ = v36
                    case 1:
                        if (!(0 <= 65535)) {
                            var v37 int16 = v25
                            if (42 <= 32767) {
                                var v38 bool = true
                                _ = v38
                            } else {
                                v5 = v25
                                switch (fi34 * fi34) {
                                case 968:
                                    if false {
                                        v4 = ((0.0 / (v4 - (-2.25 - (v4 / (3.14159 + v4))))) * v4)
                                        v39 := GenS4{F0: "\n\t", F1: v1}
                                        _ = v39
                                    }
                                default:
                                    if (255 < 329) {
                                        if (true && (false && true)) {
                                            v2 = append(v2, "foo bar")
                                            switch v1 {
                                            case "foo bar":
                                                var v40 bool = (int(v5) < (fi35 / ((fi34) | 1)))
                                                _ = v40
                                            case "\n\t":
                                                v41 := GenS3{F0: 42, F1: v3, F2: 32767}
                                                _ = v41
                                            }
                                        } else {
                                            v42 := GenS2{F0: 1, F1: ("foo bar" + "foo bar")}
                                            _ = v42
                                        }
                                        if (0 <= 2) {
                                            v1 = v1
                                        }
                                        v23 += (v23 ^ v27)
                                    }
                                }
                            }
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                            _ = v37
                        }
                    case 32767:
                        _ = len(v2)
                    }
                    _ = fi35
                }
                for fi43 := 0; fi43 < 13; fi43++ {
                    v6 = (v6 - v6)
                    _ = fi43
                }
                switch (fi34 + fi34) {
                case 704:
                    var v44 uint64 = (^(v24))
                    _ = v44
                case -9223372036854775808:
                    v2 = append(v2, "hello")
                default:
                    for fi45 := 0; fi45 < 10; fi45++ {
                        v46 := GenS4{F0: v1, F1: v1}
                        _ = fi45
                        _ = v46
                    }
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v25 = v25
                for fi47 := 0; fi47 < 14; fi47++ {
                    v48 := GenS5{F0: (0.1 > 1e10), F1: 255}
                    v4 = 1e10
                    v27 = 42
                    v23 += 0
                    if false {
                        switch fi47 {
                        case -250:
                            var v49 int8 = (-(v3))
                            _ = v49
                        case -2147483648:
                            var v50 uint8 = 195
                            _ = v50
                        case 0:
                            if false {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v51 uint = 0
                                var v52 uint64 = genFn0(v51, 0, v23)
                                _ = v51
                                _ = v52
                            }
                        default:
                            v4 += v4
                        }
                        var v53 float32 = float32((-2.25 * 0.0))
                        _ = v53
                    }
                    fmt.Println(v24, v23, v6)
                    _ = fi47
                    _ = v48
                }
            }
            _ = fi34
        }
    }
    v54 := GenS2{F0: (v23 / ((0) | 1)), F1: (v1 + v1)}
    v55 := GenS3{F0: 65535, F1: 127, F2: 2}
    v27 -= 42
    var v56 uint16 = 42
    for fi57 := 0; fi57 < 3; fi57++ {
        if true {
            v4 += v4
        }
        for fi58 := 0; fi58 < 20; fi58++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi58
        }
        _ = fi57
    }
    fmt.Println(v1, v24, v23)
    for fi59 := 0; fi59 < 11; fi59++ {
        var v60 uint64 = v24
        for fi61 := 0; fi61 < 10; fi61++ {
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            switch v1 {
            case "foo bar":
                var v62 int16 = v25
                _ = v62
            default:
                if false {
                    fmt.Println(v5, fi59)
                    if (false && (v23 < (v27 - 0))) {
                        v5 += -1
                        v63 := GenS3{F0: 628, F1: (v3 | v3), F2: (v56 ^ v56)}
                        fmt.Println(v56, v27, v23)
                        var v64 int32 = 0
                        fmt.Println(v24, v24)
                        v1 = ((("\n\t" + ("" + "foo bar")) + v1) + v1)
                        fmt.Println(v6, fi59, v64)
                        _ = v63
                        _ = v64
                    }
                    v60 -= 42
                    v65 := GenS6{F0: v25, F1: 9223372036854775807}
                    v5 = v5
                    _ = v65
                } else {
                    fmt.Println(fi61, v5)
                    var v66 uint16 = (v56 & (v56 & v56))
                    v67 := GenS7{F0: v23, F1: 42}
                    _ = v66
                    _ = v67
                }
            }
            var v68 uint = 4294967295
            var v69 int32 = (v6 & int32(fi61))
            _ = fi61
            _ = v68
            _ = v69
        }
        _ = fi59
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v6
}
