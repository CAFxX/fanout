package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint }

func main() {
    var v1 int8 = 42
    var v2 int64 = -18
    v3 := []uint32{884, 574, 1, 2147483647}
    v4 := make(map[uint32]bool)
    _ = len(v4)
    if len(v3) > 0 {
        v3[0] = 0
    }
    fmt.Println(v1)
    v5 := GenS1{F0: 42}
    fmt.Println(v1)
    if true {
        fmt.Println(v1, v2)
        for fi6 := 0; fi6 < 13; fi6++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi6
        }
    } else {
        var v7 uint16 = 70
        var v8 int32 = -1
        _ = v7
        _ = v8
    }
    var v9 bool = (!false)
    var v10 string = "\n\t"
    var v11 uint8 = 0
    v12 := v4[0]
    var v13 int16 = -687
    switch 384 {
    case -219:
        if (-128 != v1) {
            for fi14 := 0; fi14 < 1; fi14++ {
                v10 = "a"
                _ = fi14
            }
            var v15 uint = 2
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v15
        } else {
            v4[2] = true
            if (114 > 4294967295) {
                var v16 int16 = 1000
                for fi17 := 0; fi17 < 2; fi17++ {
                    v1 = v1
                    _ = fi17
                }
                _ = v16
            } else {
                var v18 int32 = -1
                switch (v10 + v10) {
                case "\n\t":
                    v19 := GenS1{F0: (v18 / (((v18 + (v18 + (v18 ^ (v18 + (v18 % ((2147483647) | 1))))))) | 1))}
                    _ = v19
                case "":
                    v2 += (^(v2))
                case "foo bar":
                    switch 588 {
                    case 2147483647:
                        var v20 uint64 = 1
                        _ = v20
                    case 75:
                        for fi21 := 0; fi21 < 16; fi21++ {
                            var v22 int32 = v18
                            fmt.Println(v11)
                            _ = fi21
                            _ = v22
                        }
                    default:
                        if v12 {
                            var v23 float32 = 0.0
                            var v24 uint16 = 402
                            _ = v23
                            _ = v24
                        } else {
                            v10 = "\n\t"
                            if v12 {
                                fmt.Println(v11, v13)
                            } else {
                                _ = len(v3)
                                fmt.Println(v1, v12, v2)
                            }
                            var v25 int = 4294967295
                            _ = v25
                        }
                    }
                }
                for fi26 := 0; fi26 < 5; fi26++ {
                    var v27 int16 = v13
                    fmt.Println(fi26)
                    if len(v3) > 0 {
                        v3[0] = 677
                    }
                    _ = fi26
                    _ = v27
                }
                _ = v18
            }
        }
    case 896:
        if ((248 != 32767) && true) {
            var v28 string = "hello"
            _ = v28
        } else {
            var v29 float32 = (1.5 * float32(0.1))
            var v30 int64 = (v2 * (v2 % ((9) | 1)))
            if false {
                for fi31 := 0; fi31 < 13; fi31++ {
                    v2 -= 0
                    switch ((v10 + "\n\t") + v10) {
                    case "":
                        fmt.Println(fi31, fi31)
                    case "\n\t":
                        switch (v10 + v10) {
                        case "\n\t":
                            v10 = "\n\t"
                        case "hello":
                            for fi32 := 0; fi32 < 2; fi32++ {
                                var v33 float32 = 3.14159
                                v34 := GenS2{F0: 65535, F1: v11, F2: (v30 << ((244) & 63))}
                                v3 = append(v3, 665)
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                        default:
                            var v35 uint32 = 4294967295
                            _ = v35
                        }
                    default:
                        if (4294967295 == 42) {
                            var v36 uint8 = v11
                            if (v11 <= v36) {
                                var v37 uint16 = 1
                                var v38 int8 = (v1 * int8(v30))
                                _ = v37
                                _ = v38
                            } else {
                                if ((v1 & 0) == -1) {
                                    if v9 {
                                        v39 := GenS3{F0: v1, F1: 920}
                                        for fi40 := 0; fi40 < 13; fi40++ {
                                            if (!(v13 <= v13)) {
                                                if (2147483647 >= 65535) {
                                                    if true {
                                                        var v41 uint8 = 156
                                                        v13 = v13
                                                        v4[2147483648] = (1 >= 1)
                                                        _ = v41
                                                    } else {
                                                        v42 := GenS2{F0: 1, F1: v11, F2: 9223372036854775807}
                                                        switch fi31 {
                                                        case 128:
                                                            v43 := GenS4{F0: 1}
                                                            _ = v43
                                                        case 1:
                                                            v10 = v10
                                                        case 32767:
                                                            v13 = v13
                                                        default:
                                                            switch 42 {
                                                            case 2147483647:
                                                                if len(v3) > 0 {
                                                                    v3[0] = 4294967295
                                                                }
                                                            case -386:
                                                                v9 = true
                                                            default:
                                                                fmt.Println(v10)
                                                            }
                                                        }
                                                        _ = v42
                                                    }
                                                } else {
                                                    v44 := GenS5{F0: (fi40 << ((318) & 63))}
                                                    _ = len(v3)
                                                    var v45 uint16 = 65535
                                                    v13 = (v13 + v13)
                                                    if false {
                                                        v1 = v1
                                                        var v46 int8 = v1
                                                        v30 = -562
                                                        v30 = -9223372036854775808
                                                        _ = v46
                                                    }
                                                    var v47 bool = false
                                                    v36 = (v11 ^ 1)
                                                    v11 = v11
                                                    if v12 {
                                                        v29 += v29
                                                        var v48 int = int(v1)
                                                        _ = v48
                                                    }
                                                    _ = v44
                                                    _ = v45
                                                    _ = v47
                                                }
                                                var v49 float64 = 0.0
                                                v11 = v11
                                                _ = v49
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = 42
                                            }
                                            _ = fi40
                                        }
                                        _ = v39
                                    }
                                    var v50 uint8 = (v36 / (((v11 / (((^(v11))) | 1))) | 1))
                                    _ = v50
                                } else {
                                    v51 := GenS6{F0: 255, F1: 255}
                                    _ = v51
                                }
                                var v52 float64 = 1.5
                                switch (v10 + ("" + v10)) {
                                case "":
                                    switch v10 {
                                    case "":
                                        fmt.Println(v29, v12, v12)
                                    }
                                case "a":
                                    var v53 float64 = ((-2.25 * -2.25) * float64(0.0))
                                    _ = v53
                                case "\n\t":
                                    v52 -= (float64(v29) + (v52 / (1.5 - ((1.5 / ((0.0 - (v52 + (0.0 / v52))) / 1.0)) / v52))))
                                }
                                _ = v52
                            }
                            for fi54 := 0; fi54 < 10; fi54++ {
                                if (v29 != (v29 - v29)) {
                                    _ = len(v3)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    if (!v12) {
                                        if len(v3) > 0 {
                                            v3[0] = 1
                                        }
                                        if false {
                                            _ = len(v3)
                                            var v55 int64 = 1
                                            _ = v55
                                        } else {
                                            if v9 {
                                                v56 := GenS3{F0: (v1 << ((668) & 63)), F1: 2}
                                                _ = v56
                                            }
                                        }
                                    }
                                }
                                _ = fi54
                            }
                            _ = v36
                        }
                    }
                    v10 = v10
                    _ = fi31
                }
                v10 = v10
            } else {
                if v9 {
                    var v57 uint = 1
                    _ = v57
                } else {
                    var v58 int16 = (v13 - v13)
                    var v59 uint32 = 0
                    if true {
                        v60 := GenS7{F0: 852}
                        for fi61 := 0; fi61 < 1; fi61++ {
                            if (v9 && (uint(v11) == 0)) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v62 float32 = 1e10
                                _ = v62
                            } else {
                                var v63 uint16 = 369
                                fmt.Println(v58, v29)
                                _ = v63
                            }
                            _ = fi61
                        }
                        _ = v60
                    } else {
                        v64 := GenS8{F0: 2}
                        _ = v64
                    }
                    _ = v58
                    _ = v59
                }
                v30 += v30
            }
            _ = v29
            _ = v30
        }
    }
    var v65 string = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v65:", v65)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v9
}
