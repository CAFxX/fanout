package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint8; F2 string }
type GenS5 struct { F0 uint; F1 int32; F2 uint64 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 uint64; F1 string; F2 uint64 }
type GenS8 struct { F0 uint64; F1 int }
type GenS9 struct { F0 int16; F1 uint }

func main() {
    var v1 string = "hello"
    var v2 int64 = 42
    var v3 float32 = 1e10
    var v4 uint32 = 137
    v5 := []int{-1, 32768, -9223372036854775808}
    if (!(!(-2147483648 != 2147483647))) {
        var v6 int16 = 0
        _ = v6
    }
    v4 += (v4 << ((2) & 63))
    fmt.Println(v3, v3, v2)
    v2 -= -585
    var v7 uint64 = 2
    if false {
        for fi8 := 0; fi8 < 19; fi8++ {
            if true {
                fmt.Println(v7, v4, v1)
                v9 := GenS1{F0: 128}
                _ = v9
            }
            _ = fi8
        }
        v1 = "a"
        v5 = append(v5, 127)
    }
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    switch (v1 + v1) {
    case "":
        if false {
            var v10 int8 = 127
            v1 = v1
            _ = v10
        } else {
            for fi11 := 0; fi11 < 9; fi11++ {
                if false {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v2 += v2
                    var v12 int = fi11
                    _ = v12
                } else {
                    v2 -= (^(v2))
                    fmt.Println(v4, v2, fi11)
                    if (v3 < 0.1) {
                        if (-243 < (v2 << ((4294967295) & 63))) {
                            switch v1 {
                            case "hello":
                                v13 := GenS1{F0: 255}
                                _ = v13
                            case "a":
                                if true {
                                    v14 := GenS1{F0: 231}
                                    if (0 <= -722) {
                                        fmt.Println(fi11)
                                    } else {
                                        var v15 float32 = 1e10
                                        for fi16 := 0; fi16 < 17; fi16++ {
                                            v17 := GenS1{F0: 0}
                                            v1 = v1
                                            v18 := GenS1{F0: 128}
                                            _ = fi16
                                            _ = v17
                                            _ = v18
                                        }
                                        _ = v15
                                    }
                                    fmt.Println(v2, v3)
                                    _ = v14
                                } else {
                                    fmt.Println(v7, v7, v3)
                                    _ = len(v5)
                                    var v19 int8 = -128
                                    _ = v19
                                }
                            case "foo bar":
                                if (!false) {
                                    if (v1 != v1) {
                                        v20 := GenS1{F0: 42}
                                        var v21 int32 = 1
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v20
                                        _ = v21
                                    } else {
                                        var v22 uint8 = 182
                                        _ = v22
                                    }
                                } else {
                                    for fi23 := 0; fi23 < 12; fi23++ {
                                        v2 -= (v2 | (v2 << ((2) & 63)))
                                        v24 := GenS2{F0: 933, F1: -1}
                                        _ = fi23
                                        _ = v24
                                    }
                                }
                            default:
                                var v25 int64 = v2
                                _ = v25
                            }
                            var v26 string = "foo bar"
                            if (115 <= (v7 / (((^(v7))) | 1))) {
                                if (3.14159 <= ((v3 - 1e10) * float32((1.5 / 1.5)))) {
                                    var v27 int32 = -2147483648
                                    var v28 uint = 642
                                    v28 += 42
                                    _ = v27
                                    _ = v28
                                } else {
                                    var v29 float32 = v3
                                    for fi30 := 0; fi30 < 11; fi30++ {
                                        var v31 bool = ((v4 | 2147483648) <= v4)
                                        if v31 {
                                            v32 := GenS3{F0: v2}
                                            if v31 {
                                                v33 := GenS4{F0: 42, F1: 127, F2: ((v26 + (v26 + (((((v26 + v1) + v26) + v26) + v1) + v26))) + v1)}
                                                _ = v33
                                            }
                                            if (32767 < 1000) {
                                                switch (fi11 / ((fi30) | 1)) {
                                                case 32767:
                                                    switch fi30 {
                                                    case 9223372036854775807:
                                                        var v34 float64 = float64(0.1)
                                                        _ = v34
                                                    default:
                                                        v35 := GenS4{F0: 226, F1: 42, F2: ("a" + v26)}
                                                        _ = v35
                                                    }
                                                case 199:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                case 148:
                                                    var v36 string = "a"
                                                    _ = v36
                                                default:
                                                    v1 = ("\n\t" + v26)
                                                }
                                            } else {
                                                v37 := GenS4{F0: 2, F1: 44, F2: v1}
                                                v38 := GenS2{F0: 0, F1: -1}
                                                _ = v37
                                                _ = v38
                                            }
                                            v29 += float32(1.5)
                                            if (100 < -128) {
                                                v39 := GenS3{F0: -479}
                                                var v40 float64 = 1e10
                                                _ = v39
                                                _ = v40
                                            } else {
                                                if (float32(((3.14159 + 1e10) - (3.14159 + 1e10))) < (float32((1e10 * (0.1 - -2.25))) / (0.0 - ((float32(((1e10 - 1e10) / 0.1)) - 1.5) / (v3 / 1e10))))) {
                                                    v29 += v29
                                                    var v41 int16 = -32768
                                                    v42 := GenS4{F0: 1, F1: 2, F2: ((v1 + "\n\t") + "")}
                                                    _ = v41
                                                    _ = v42
                                                } else {
                                                    v43 := GenS5{F0: 255, F1: -882, F2: v7}
                                                    _ = v43
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v3 += v3
                                            }
                                            if v31 {
                                                var v44 uint8 = 42
                                                v26 = "a"
                                                fmt.Println(v31, v2, fi11)
                                                v45 := GenS6{F0: (v7 + v7), F1: 0}
                                                v4 = v4
                                                v5 = append(v5, fi30)
                                                var v46 int = int(v4)
                                                _ = v44
                                                _ = v45
                                                _ = v46
                                            }
                                            _ = v32
                                        }
                                        if v31 {
                                            v47 := GenS3{F0: (^(v2))}
                                            fmt.Println(fi30, fi11)
                                            _ = v47
                                        } else {
                                            fmt.Println(v4, v29)
                                            v3 = float32((1.5 - (3.14159 * -2.25)))
                                        }
                                        _ = fi30
                                        _ = v31
                                    }
                                    _ = v29
                                }
                            } else {
                                v48 := GenS7{F0: (^(v7)), F1: v1, F2: v7}
                                v49 := GenS8{F0: (^(v7)), F1: 4294967295}
                                if true {
                                    if ((v1 + v26) <= v1) {
                                        v26 = v26
                                        switch v26 {
                                        case "foo bar":
                                            var v50 int8 = 127
                                            _ = v50
                                        case "hello":
                                            var v51 uint32 = (^(v4))
                                            _ = v51
                                        case "":
                                            v2 = v2
                                        default:
                                            fmt.Println(v4, v4, fi11)
                                        }
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = (fi11 + (fi11 % ((fi11) | 1)))
                                    }
                                    if (v26 > ((v26 + (v26 + "")) + v26)) {
                                        switch v26 {
                                        case "":
                                            if len(v5) > 0 {
                                                v5[0] = fi11
                                            }
                                        }
                                        var v52 int8 = -128
                                        switch v26 {
                                        case "hello":
                                            var v53 int8 = (v52 % ((127) | 1))
                                            _ = v53
                                        case "a":
                                            fmt.Println(v1, v4)
                                        case "foo bar":
                                            var v54 int = (fi11 | (fi11 | -1))
                                            _ = v54
                                        }
                                        _ = v52
                                    } else {
                                        var v55 uint = 65535
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        v56 := GenS8{F0: 1, F1: fi11}
                                        _ = v55
                                        _ = v56
                                    }
                                } else {
                                    if false {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    } else {
                                        if (v3 != v3) {
                                            for fi57 := 0; fi57 < 17; fi57++ {
                                                switch (fi57 >> ((18446744073709551615) & 63)) {
                                                case 9223372036854775807:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    v7 -= v7
                                                }
                                                _ = fi57
                                            }
                                            v58 := GenS6{F0: (v7 >> ((4294967295) & 63)), F1: 1}
                                            switch "hello" {
                                            case "foo bar":
                                                var v59 int = (fi11 | fi11)
                                                _ = v59
                                            default:
                                                var v60 int64 = (v2 ^ (v2 ^ (v2 ^ v2)))
                                                _ = v60
                                            }
                                            _ = v58
                                        } else {
                                            for fi61 := 0; fi61 < 5; fi61++ {
                                                v62 := GenS6{F0: v7, F1: 65535}
                                                v63 := GenS8{F0: v7, F1: (fi61 / ((-680) | 1))}
                                                _ = fi61
                                                _ = v62
                                                _ = v63
                                            }
                                            var v64 int32 = -2147483648
                                            _ = v64
                                        }
                                        v3 += v3
                                    }
                                    v3 += (0.1 * 1.5)
                                }
                                _ = v48
                                _ = v49
                            }
                            _ = v26
                        } else {
                            for fi65 := 0; fi65 < 10; fi65++ {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = fi65
                            }
                            switch (fi11 >> ((0) & 63)) {
                            case 32768:
                                fmt.Println(v4)
                            case -970:
                                for fi66 := 0; fi66 < 12; fi66++ {
                                    v67 := GenS7{F0: v7, F1: (v1 + v1), F2: (v7 + v7)}
                                    _ = fi66
                                    _ = v67
                                }
                            case -605:
                                v2 = v2
                            }
                        }
                        if (-676 >= -451) {
                            v68 := GenS9{F0: -753, F1: uint(v4)}
                            _ = v68
                        }
                    }
                }
                _ = fi11
            }
            v1 = "\n\t"
        }
    default:
        if true {
            v69 := GenS5{F0: 0, F1: -2147483648, F2: (v7 * v7)}
            v7 = v7
            _ = v69
        }
    }
    fmt.Println(v3, v2)
    for fi70 := 0; fi70 < 3; fi70++ {
        v4 = (v4 / (((^(v4))) | 1))
        var v71 uint8 = 165
        _ = fi70
        _ = v71
    }
    v1 = v1
    for fi72 := 0; fi72 < 18; fi72++ {
        fmt.Println(v4, v7)
        var v73 uint8 = 2
        v3 = v3
        _ = fi72
        _ = v73
    }
    fmt.Println(v1)
    fmt.Println(v7)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
