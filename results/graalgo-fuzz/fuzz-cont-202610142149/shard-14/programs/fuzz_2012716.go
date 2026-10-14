package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int8; F1 uint32; F2 bool }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint16; F1 uint32 }
type GenS8 struct { F0 bool; F1 int8; F2 uint64 }
type GenS9 struct { F0 uint16; F1 int8 }
type GenS10 struct { F0 string; F1 int16; F2 uint8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int16; F1 int16; F2 int64 }
type GenS14 struct { F0 int64; F1 uint32 }

func main() {
    v1 := []int{}
    var v2 float64 = -2.25
    var v3 int8 = 49
    var v4 uint32 = 402
    v5 := []uint32{304}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (108 <= (v3 - (v3 << ((255) & 63)))) {
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 string = ""
            switch fi6 {
            case 42:
                switch (fi6 | (fi6 * fi6)) {
                case 9223372036854775807:
                    if true {
                        for fi8 := 0; fi8 < 9; fi8++ {
                            if (4294967295 != 1) {
                                if (-284 <= -32768) {
                                    switch (fi6 / ((fi6) | 1)) {
                                    case 438:
                                        var v9 int64 = -1
                                        _ = v9
                                    case -2147483648:
                                        if (!(v2 == v2)) {
                                            v2 += 1.5
                                            v10 := GenS1{F0: 127}
                                            v5 = append(v5, v4)
                                            _ = v10
                                        } else {
                                            v2 += (3.14159 * v2)
                                        }
                                    }
                                    var v11 uint64 = 350
                                    var v12 int8 = (v3 & 42)
                                    _ = v11
                                    _ = v12
                                } else {
                                    v2 = v2
                                    if (!true) {
                                        v1 = append(v1, (fi6 + fi6))
                                        v13 := GenS2{F0: 128, F1: 2147483647}
                                        _ = v13
                                    }
                                    if (0 < 70) {
                                        v14 := GenS3{F0: (v3 % (((v3 | (v3 + (v3 | v3)))) | 1)), F1: (^(v4)), F2: false}
                                        switch "\n\t" {
                                        case "\n\t":
                                            if true {
                                                if (!true) {
                                                    if true {
                                                        var v15 uint64 = 42
                                                        v7 = v7
                                                        var v16 uint32 = v4
                                                        v5 = append(v5, v16)
                                                        _ = v15
                                                        _ = v16
                                                    } else {
                                                        if (749 <= 42) {
                                                            v7 = v7
                                                            v2 = 1e10
                                                            if true {
                                                                switch v7 {
                                                                case "hello":
                                                                    fmt.Println(v3)
                                                                }
                                                            }
                                                            if false {
                                                                var v17 int16 = int16(v3)
                                                                v18 := GenS1{F0: 0}
                                                                v2 = v2
                                                                _ = v17
                                                                _ = v18
                                                            } else {
                                                                var v19 bool = (!((v7 + (((v7 + ((((v7 + (v7 + v7)) + v7) + v7) + v7)) + ((v7 + ("\n\t" + "hello")) + ((v7 + "\n\t") + v7))) + "foo bar")) < ((v7 + "") + ("hello" + ((v7 + v7) + (v7 + (v7 + v7)))))))
                                                                _ = v19
                                                            }
                                                        } else {
                                                            var v20 int = (fi6 / (((fi8 | fi8)) | 1))
                                                            v21 := GenS4{F0: 255, F1: -9223372036854775808}
                                                            _ = v20
                                                            _ = v21
                                                        }
                                                        if true {
                                                            switch v7 {
                                                            case "\n\t":
                                                                if false {
                                                                    fmt.Println(v3)
                                                                    if ((128 != 0) || true) {
                                                                        var v22 float32 = 1.5
                                                                        v23 := GenS5{F0: -2147483648}
                                                                        v1 = append(v1, 9223372036854775807)
                                                                        v22 = v22
                                                                        var v24 int64 = -519
                                                                        _ = v22
                                                                        _ = v23
                                                                        _ = v24
                                                                    } else {
                                                                        var v25 int64 = 9223372036854775807
                                                                        var v26 string = ((((((((((("" + "") + ("hello" + "hello")) + "hello") + v7) + v7) + (((v7 + (v7 + v7)) + v7) + ("a" + ("hello" + "\n\t")))) + "\n\t") + v7) + (v7 + (v7 + "hello"))) + (v7 + "\n\t")) + ((("foo bar" + (v7 + v7)) + v7) + "foo bar"))
                                                                        _ = v25
                                                                        _ = v26
                                                                    }
                                                                    var v27 uint64 = 1
                                                                    _ = v27
                                                                } else {
                                                                    v28 := GenS6{F0: (("hello" + ("foo bar" + "hello")) + (((v7 + "") + ("a" + v7)) + v7))}
                                                                    _ = v28
                                                                }
                                                            case "":
                                                                if (0 >= 18446744073709551615) {
                                                                    var v29 uint64 = 1
                                                                    var v30 int = fi8
                                                                    var v31 uint64 = v29
                                                                    _ = v29
                                                                    _ = v30
                                                                    _ = v31
                                                                }
                                                                var v32 bool = false
                                                                _ = v32
                                                            }
                                                            var v33 float32 = (0.1 - -2.25)
                                                            _ = v33
                                                        }
                                                        v2 = v2
                                                        v2 -= float64((1e10 + 3.14159))
                                                        v5 = append(v5, 854)
                                                    }
                                                    v34 := GenS7{F0: 32767, F1: 42}
                                                    _ = v34
                                                }
                                                v35 := GenS3{F0: (^(v3)), F1: (v4 << ((0) & 63)), F2: ((3.14159 + 1e10) <= 1e10)}
                                                v36 := GenS4{F0: uint16(v4), F1: 42}
                                                _ = v35
                                                _ = v36
                                            } else {
                                                if false {
                                                    _ = len(v5)
                                                    v4 = v4
                                                    fmt.Println(fi8, v3)
                                                    _ = len(v5)
                                                    v7 = v7
                                                }
                                                v37 := GenS8{F0: true, F1: (v3 | (v3 - (v3 + (^(v3))))), F2: 1}
                                                _ = v37
                                            }
                                        case "hello":
                                            v5 = append(v5, 65535)
                                        }
                                        var v38 uint8 = 159
                                        _ = v14
                                        _ = v38
                                    } else {
                                        v39 := GenS1{F0: 42}
                                        _ = v39
                                    }
                                }
                                switch v7 {
                                case "\n\t":
                                    v3 = (v3 % ((81) | 1))
                                case "hello":
                                    v2 -= v2
                                }
                            }
                            _ = fi8
                        }
                        for fi40 := 0; fi40 < 11; fi40++ {
                            v7 = ("\n\t" + "a")
                            if true {
                                var v41 float32 = 0.0
                                _ = v41
                            } else {
                                v42 := GenS9{F0: 65535, F1: v3}
                                fmt.Println(v4, v2)
                                switch (v7 + "a") {
                                case "a":
                                    if true {
                                        switch (fi40 | 82) {
                                        case -1:
                                            v43 := GenS10{F0: v7, F1: 0, F2: 20}
                                            _ = v43
                                        case 32767:
                                            var v44 int16 = -429
                                            _ = v44
                                        }
                                        v2 = (0.1 * (v2 - -2.25))
                                        v2 -= 3.14159
                                    }
                                case "":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                                _ = v42
                            }
                            _ = fi40
                        }
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        switch (v7 + (("foo bar" + "hello") + "")) {
                        case "hello":
                            if (true || (((^(v4)) != 4294967295) || ((v3 ^ (v3 % ((v3) | 1))) == 96))) {
                                if len(v5) > 0 {
                                    v5[0] = 0
                                }
                            }
                        case "\n\t":
                            v45 := GenS11{F0: 90}
                            _ = v45
                        case "":
                            switch 361 {
                            case -303:
                                v3 += v3
                            case 128:
                                v7 = v7
                            }
                        default:
                            if true {
                                var v46 int8 = -119
                                _ = v46
                            } else {
                                v47 := GenS3{F0: (v3 >> ((791) & 63)), F1: v4, F2: false}
                                v3 += (v3 & (v3 / (((^(v3))) | 1)))
                                switch fi6 {
                                case -9223372036854775808:
                                    for fi48 := 0; fi48 < 6; fi48++ {
                                        var v49 uint32 = (v4 | 4294967295)
                                        var v50 string = ("\n\t" + v7)
                                        if (0 <= 42) {
                                            var v51 float64 = 0.0
                                            _ = v51
                                        } else {
                                            if true {
                                                var v52 float32 = (0.0 + 1.5)
                                                _ = v52
                                            } else {
                                                v53 := GenS4{F0: 129, F1: 1}
                                                var v54 int = 127
                                                _ = v53
                                                _ = v54
                                            }
                                            v3 = (v3 * v3)
                                            var v55 bool = (1 == 9223372036854775807)
                                            _ = v55
                                        }
                                        _ = fi48
                                        _ = v49
                                        _ = v50
                                    }
                                case -589:
                                    _ = len(v5)
                                case 275:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                default:
                                    var v56 int16 = -149
                                    _ = v56
                                }
                                _ = v47
                            }
                        }
                        switch 128 {
                        case -631:
                            v7 = "\n\t"
                        case -9223372036854775808:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            for fi57 := 0; fi57 < 17; fi57++ {
                                var v58 int32 = 0
                                var v59 uint = 617
                                fmt.Println(v4, v59)
                                _ = fi57
                                _ = v58
                                _ = v59
                            }
                        }
                    }
                default:
                    v60 := GenS12{F0: -442}
                    _ = v60
                }
            case -2147483648:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                v61 := GenS7{F0: 255, F1: (v4 ^ v4)}
                _ = v61
            }
            _ = fi6
            _ = v7
        }
    }
    switch "\n\t" {
    case "":
        for fi62 := 0; fi62 < 19; fi62++ {
            v2 += v2
            v63 := GenS13{F0: -32768, F1: 742, F2: -9223372036854775808}
            _ = fi62
            _ = v63
        }
    default:
        fmt.Println(v4, v4)
    }
    if false {
        v64 := GenS6{F0: ("foo bar" + "foo bar")}
        _ = v64
    } else {
        v65 := GenS14{F0: 395, F1: v4}
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        fmt.Println(v4, v4, v2)
        _ = v65
    }
    var v66 uint16 = 0
    if ((!(v2 >= 1.5)) && false) {
        v66 += v66
        if true {
            fmt.Println(v66, v2)
        }
        if (18446744073709551615 < 4294967295) {
            fmt.Println(v3, v66, v4)
            for fi67 := 0; fi67 < 4; fi67++ {
                v2 -= v2
                _ = fi67
            }
            v3 = 0
        } else {
            if len(v5) > 0 {
                v5[0] = (v4 + 851)
            }
            if len(v1) > 0 {
                v1[0] = 329
            }
            v2 = 0.0
        }
    }
    if (42 >= -746) {
        fmt.Println(v3, v4, v3)
        for fi68 := 0; fi68 < 4; fi68++ {
            var v69 uint = 4294967295
            _ = fi68
            _ = v69
        }
        for fi70 := 0; fi70 < 18; fi70++ {
            var v71 int64 = 1
            fmt.Println(v66, v66)
            v1 = append(v1, fi70)
            _ = fi70
            _ = v71
        }
    } else {
        var v72 string = "\n\t"
        fmt.Println(v66, v72, v2)
        _ = v72
    }
    var v73 int64 = 1
    if false {
        v73 = (v73 << ((65535) & 63))
    }
    v73 = v73
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v66:", v66)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v73
}
